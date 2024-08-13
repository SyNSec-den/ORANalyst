// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/elliptic/elliptic.go:5
// Package elliptic implements the standard NIST P-224, P-256, P-384, and P-521
//line /usr/local/go/src/crypto/elliptic/elliptic.go:5
// elliptic curves over prime fields.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:5
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:5
// The P224(), P256(), P384() and P521() values are necessary to use the crypto/ecdsa package.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:5
// Most other uses should migrate to the more efficient and safer crypto/ecdh package.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
package elliptic

//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
import (
//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
)
//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
import (
//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:10
)

import (
	"io"
	"math/big"
	"sync"
)

// A Curve represents a short-form Weierstrass curve with a=-3.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
// The behavior of Add, Double, and ScalarMult when the input is not a point on
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
// the curve is undefined.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
// Note that the conventional point at infinity (0, 0) is not considered on the
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
// curve, although it can be returned by Add, Double, ScalarMult, or
//line /usr/local/go/src/crypto/elliptic/elliptic.go:18
// ScalarBaseMult (but not the Unmarshal or UnmarshalCompressed functions).
//line /usr/local/go/src/crypto/elliptic/elliptic.go:26
type Curve interface {
	// Params returns the parameters for the curve.
	Params() *CurveParams

	// IsOnCurve reports whether the given (x,y) lies on the curve.
	//
	// Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. The NewPublicKey methods of NIST curves in crypto/ecdh accept
	// the same encoding as the Unmarshal function, and perform on-curve checks.
	IsOnCurve(x, y *big.Int) bool

	// Add returns the sum of (x1,y1) and (x2,y2).
	//
	// Note: this is a low-level unsafe API.
	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

	// Double returns 2*(x,y).
	//
	// Note: this is a low-level unsafe API.
	Double(x1, y1 *big.Int) (x, y *big.Int)

	// ScalarMult returns k*(x,y) where k is an integer in big-endian form.
	//
	// Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. Most uses of ScalarMult can be replaced by a call to the ECDH
	// methods of NIST curves in crypto/ecdh.
	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

	// ScalarBaseMult returns k*G, where G is the base point of the group
	// and k is an integer in big-endian form.
	//
	// Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. Most uses of ScalarBaseMult can be replaced by a call to the
	// PrivateKey.PublicKey method in crypto/ecdh.
	ScalarBaseMult(k []byte) (x, y *big.Int)
}

var mask = []byte{0xff, 0x1, 0x3, 0x7, 0xf, 0x1f, 0x3f, 0x7f}

// GenerateKey returns a public/private key pair. The private key is
//line /usr/local/go/src/crypto/elliptic/elliptic.go:65
// generated using the given reader, which must return random data.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:65
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:65
// Note: for ECDH, use the GenerateKey methods of the crypto/ecdh package;
//line /usr/local/go/src/crypto/elliptic/elliptic.go:65
// for ECDSA, use the GenerateKey function of the crypto/ecdsa package.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:70
func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:70
	_go_fuzz_dep_.CoverTab[6939]++
								N := curve.Params().N
								bitSize := N.BitLen()
								byteLen := (bitSize + 7) / 8
								priv = make([]byte, byteLen)

								for x == nil {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:76
		_go_fuzz_dep_.CoverTab[6941]++
									_, err = io.ReadFull(rand, priv)
									if err != nil {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:78
			_go_fuzz_dep_.CoverTab[6944]++
										return
//line /usr/local/go/src/crypto/elliptic/elliptic.go:79
			// _ = "end of CoverTab[6944]"
		} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:80
			_go_fuzz_dep_.CoverTab[6945]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:80
			// _ = "end of CoverTab[6945]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:80
		}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:80
		// _ = "end of CoverTab[6941]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:80
		_go_fuzz_dep_.CoverTab[6942]++

//line /usr/local/go/src/crypto/elliptic/elliptic.go:83
		priv[0] &= mask[bitSize%8]

//line /usr/local/go/src/crypto/elliptic/elliptic.go:86
		priv[1] ^= 0x42

//line /usr/local/go/src/crypto/elliptic/elliptic.go:89
		if new(big.Int).SetBytes(priv).Cmp(N) >= 0 {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:89
			_go_fuzz_dep_.CoverTab[6946]++
										continue
//line /usr/local/go/src/crypto/elliptic/elliptic.go:90
			// _ = "end of CoverTab[6946]"
		} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:91
			_go_fuzz_dep_.CoverTab[6947]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:91
			// _ = "end of CoverTab[6947]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:91
		}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:91
		// _ = "end of CoverTab[6942]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:91
		_go_fuzz_dep_.CoverTab[6943]++

									x, y = curve.ScalarBaseMult(priv)
//line /usr/local/go/src/crypto/elliptic/elliptic.go:93
		// _ = "end of CoverTab[6943]"
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:94
	// _ = "end of CoverTab[6939]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:94
	_go_fuzz_dep_.CoverTab[6940]++
								return
//line /usr/local/go/src/crypto/elliptic/elliptic.go:95
	// _ = "end of CoverTab[6940]"
}

// Marshal converts a point on the curve into the uncompressed form specified in
//line /usr/local/go/src/crypto/elliptic/elliptic.go:98
// SEC 1, Version 2.0, Section 2.3.3. If the point is not on the curve (or is
//line /usr/local/go/src/crypto/elliptic/elliptic.go:98
// the conventional point at infinity), the behavior is undefined.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:98
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:98
// Note: for ECDH, use the crypto/ecdh package. This function returns an
//line /usr/local/go/src/crypto/elliptic/elliptic.go:98
// encoding equivalent to that of PublicKey.Bytes in crypto/ecdh.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:104
func Marshal(curve Curve, x, y *big.Int) []byte {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:104
	_go_fuzz_dep_.CoverTab[6948]++
								panicIfNotOnCurve(curve, x, y)

								byteLen := (curve.Params().BitSize + 7) / 8

								ret := make([]byte, 1+2*byteLen)
								ret[0] = 4

								x.FillBytes(ret[1 : 1+byteLen])
								y.FillBytes(ret[1+byteLen : 1+2*byteLen])

								return ret
//line /usr/local/go/src/crypto/elliptic/elliptic.go:115
	// _ = "end of CoverTab[6948]"
}

// MarshalCompressed converts a point on the curve into the compressed form
//line /usr/local/go/src/crypto/elliptic/elliptic.go:118
// specified in SEC 1, Version 2.0, Section 2.3.3. If the point is not on the
//line /usr/local/go/src/crypto/elliptic/elliptic.go:118
// curve (or is the conventional point at infinity), the behavior is undefined.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:121
func MarshalCompressed(curve Curve, x, y *big.Int) []byte {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:121
	_go_fuzz_dep_.CoverTab[6949]++
								panicIfNotOnCurve(curve, x, y)
								byteLen := (curve.Params().BitSize + 7) / 8
								compressed := make([]byte, 1+byteLen)
								compressed[0] = byte(y.Bit(0)) | 2
								x.FillBytes(compressed[1:])
								return compressed
//line /usr/local/go/src/crypto/elliptic/elliptic.go:127
	// _ = "end of CoverTab[6949]"
}

// unmarshaler is implemented by curves with their own constant-time Unmarshal.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:130
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:130
// There isn't an equivalent interface for Marshal/MarshalCompressed because
//line /usr/local/go/src/crypto/elliptic/elliptic.go:130
// that doesn't involve any mathematical operations, only FillBytes and Bit.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:134
type unmarshaler interface {
	Unmarshal([]byte) (x, y *big.Int)
	UnmarshalCompressed([]byte) (x, y *big.Int)
}

// Assert that the known curves implement unmarshaler.
var _ = []unmarshaler{p224, p256, p384, p521}

// Unmarshal converts a point, serialized by Marshal, into an x, y pair. It is
//line /usr/local/go/src/crypto/elliptic/elliptic.go:142
// an error if the point is not in uncompressed form, is not on the curve, or is
//line /usr/local/go/src/crypto/elliptic/elliptic.go:142
// the point at infinity. On error, x = nil.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:142
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:142
// Note: for ECDH, use the crypto/ecdh package. This function accepts an
//line /usr/local/go/src/crypto/elliptic/elliptic.go:142
// encoding equivalent to that of the NewPublicKey methods in crypto/ecdh.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:148
func Unmarshal(curve Curve, data []byte) (x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:148
	_go_fuzz_dep_.CoverTab[6950]++
								if c, ok := curve.(unmarshaler); ok {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:149
		_go_fuzz_dep_.CoverTab[6956]++
									return c.Unmarshal(data)
//line /usr/local/go/src/crypto/elliptic/elliptic.go:150
		// _ = "end of CoverTab[6956]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:151
		_go_fuzz_dep_.CoverTab[6957]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:151
		// _ = "end of CoverTab[6957]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:151
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:151
	// _ = "end of CoverTab[6950]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:151
	_go_fuzz_dep_.CoverTab[6951]++

								byteLen := (curve.Params().BitSize + 7) / 8
								if len(data) != 1+2*byteLen {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:154
		_go_fuzz_dep_.CoverTab[6958]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:155
		// _ = "end of CoverTab[6958]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:156
		_go_fuzz_dep_.CoverTab[6959]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:156
		// _ = "end of CoverTab[6959]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:156
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:156
	// _ = "end of CoverTab[6951]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:156
	_go_fuzz_dep_.CoverTab[6952]++
								if data[0] != 4 {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:157
		_go_fuzz_dep_.CoverTab[6960]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:158
		// _ = "end of CoverTab[6960]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:159
		_go_fuzz_dep_.CoverTab[6961]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:159
		// _ = "end of CoverTab[6961]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:159
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:159
	// _ = "end of CoverTab[6952]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:159
	_go_fuzz_dep_.CoverTab[6953]++
								p := curve.Params().P
								x = new(big.Int).SetBytes(data[1 : 1+byteLen])
								y = new(big.Int).SetBytes(data[1+byteLen:])
								if x.Cmp(p) >= 0 || func() bool {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:163
		_go_fuzz_dep_.CoverTab[6962]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:163
		return y.Cmp(p) >= 0
//line /usr/local/go/src/crypto/elliptic/elliptic.go:163
		// _ = "end of CoverTab[6962]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:163
	}() {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:163
		_go_fuzz_dep_.CoverTab[6963]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:164
		// _ = "end of CoverTab[6963]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:165
		_go_fuzz_dep_.CoverTab[6964]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:165
		// _ = "end of CoverTab[6964]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:165
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:165
	// _ = "end of CoverTab[6953]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:165
	_go_fuzz_dep_.CoverTab[6954]++
								if !curve.IsOnCurve(x, y) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:166
		_go_fuzz_dep_.CoverTab[6965]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:167
		// _ = "end of CoverTab[6965]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:168
		_go_fuzz_dep_.CoverTab[6966]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:168
		// _ = "end of CoverTab[6966]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:168
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:168
	// _ = "end of CoverTab[6954]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:168
	_go_fuzz_dep_.CoverTab[6955]++
								return
//line /usr/local/go/src/crypto/elliptic/elliptic.go:169
	// _ = "end of CoverTab[6955]"
}

// UnmarshalCompressed converts a point, serialized by MarshalCompressed, into
//line /usr/local/go/src/crypto/elliptic/elliptic.go:172
// an x, y pair. It is an error if the point is not in compressed form, is not
//line /usr/local/go/src/crypto/elliptic/elliptic.go:172
// on the curve, or is the point at infinity. On error, x = nil.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:175
func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:175
	_go_fuzz_dep_.CoverTab[6967]++
								if c, ok := curve.(unmarshaler); ok {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:176
		_go_fuzz_dep_.CoverTab[6975]++
									return c.UnmarshalCompressed(data)
//line /usr/local/go/src/crypto/elliptic/elliptic.go:177
		// _ = "end of CoverTab[6975]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:178
		_go_fuzz_dep_.CoverTab[6976]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:178
		// _ = "end of CoverTab[6976]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:178
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:178
	// _ = "end of CoverTab[6967]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:178
	_go_fuzz_dep_.CoverTab[6968]++

								byteLen := (curve.Params().BitSize + 7) / 8
								if len(data) != 1+byteLen {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:181
		_go_fuzz_dep_.CoverTab[6977]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:182
		// _ = "end of CoverTab[6977]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:183
		_go_fuzz_dep_.CoverTab[6978]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:183
		// _ = "end of CoverTab[6978]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:183
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:183
	// _ = "end of CoverTab[6968]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:183
	_go_fuzz_dep_.CoverTab[6969]++
								if data[0] != 2 && func() bool {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:184
		_go_fuzz_dep_.CoverTab[6979]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:184
		return data[0] != 3
//line /usr/local/go/src/crypto/elliptic/elliptic.go:184
		// _ = "end of CoverTab[6979]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:184
	}() {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:184
		_go_fuzz_dep_.CoverTab[6980]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:185
		// _ = "end of CoverTab[6980]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:186
		_go_fuzz_dep_.CoverTab[6981]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:186
		// _ = "end of CoverTab[6981]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:186
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:186
	// _ = "end of CoverTab[6969]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:186
	_go_fuzz_dep_.CoverTab[6970]++
								p := curve.Params().P
								x = new(big.Int).SetBytes(data[1:])
								if x.Cmp(p) >= 0 {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:189
		_go_fuzz_dep_.CoverTab[6982]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:190
		// _ = "end of CoverTab[6982]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:191
		_go_fuzz_dep_.CoverTab[6983]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:191
		// _ = "end of CoverTab[6983]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:191
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:191
	// _ = "end of CoverTab[6970]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:191
	_go_fuzz_dep_.CoverTab[6971]++

								y = curve.Params().polynomial(x)
								y = y.ModSqrt(y, p)
								if y == nil {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:195
		_go_fuzz_dep_.CoverTab[6984]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:196
		// _ = "end of CoverTab[6984]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:197
		_go_fuzz_dep_.CoverTab[6985]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:197
		// _ = "end of CoverTab[6985]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:197
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:197
	// _ = "end of CoverTab[6971]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:197
	_go_fuzz_dep_.CoverTab[6972]++
								if byte(y.Bit(0)) != data[0]&1 {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:198
		_go_fuzz_dep_.CoverTab[6986]++
									y.Neg(y).Mod(y, p)
//line /usr/local/go/src/crypto/elliptic/elliptic.go:199
		// _ = "end of CoverTab[6986]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:200
		_go_fuzz_dep_.CoverTab[6987]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:200
		// _ = "end of CoverTab[6987]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:200
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:200
	// _ = "end of CoverTab[6972]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:200
	_go_fuzz_dep_.CoverTab[6973]++
								if !curve.IsOnCurve(x, y) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:201
		_go_fuzz_dep_.CoverTab[6988]++
									return nil, nil
//line /usr/local/go/src/crypto/elliptic/elliptic.go:202
		// _ = "end of CoverTab[6988]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:203
		_go_fuzz_dep_.CoverTab[6989]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:203
		// _ = "end of CoverTab[6989]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:203
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:203
	// _ = "end of CoverTab[6973]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:203
	_go_fuzz_dep_.CoverTab[6974]++
								return
//line /usr/local/go/src/crypto/elliptic/elliptic.go:204
	// _ = "end of CoverTab[6974]"
}

func panicIfNotOnCurve(curve Curve, x, y *big.Int) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:207
	_go_fuzz_dep_.CoverTab[6990]++

//line /usr/local/go/src/crypto/elliptic/elliptic.go:210
	if x.Sign() == 0 && func() bool {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:210
		_go_fuzz_dep_.CoverTab[6992]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:210
		return y.Sign() == 0
//line /usr/local/go/src/crypto/elliptic/elliptic.go:210
		// _ = "end of CoverTab[6992]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:210
	}() {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:210
		_go_fuzz_dep_.CoverTab[6993]++
									return
//line /usr/local/go/src/crypto/elliptic/elliptic.go:211
		// _ = "end of CoverTab[6993]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:212
		_go_fuzz_dep_.CoverTab[6994]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:212
		// _ = "end of CoverTab[6994]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:212
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:212
	// _ = "end of CoverTab[6990]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:212
	_go_fuzz_dep_.CoverTab[6991]++

								if !curve.IsOnCurve(x, y) {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:214
		_go_fuzz_dep_.CoverTab[6995]++
									panic("crypto/elliptic: attempted operation on invalid point")
//line /usr/local/go/src/crypto/elliptic/elliptic.go:215
		// _ = "end of CoverTab[6995]"
	} else {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:216
		_go_fuzz_dep_.CoverTab[6996]++
//line /usr/local/go/src/crypto/elliptic/elliptic.go:216
		// _ = "end of CoverTab[6996]"
//line /usr/local/go/src/crypto/elliptic/elliptic.go:216
	}
//line /usr/local/go/src/crypto/elliptic/elliptic.go:216
	// _ = "end of CoverTab[6991]"
}

var initonce sync.Once

func initAll() {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:221
	_go_fuzz_dep_.CoverTab[6997]++
								initP224()
								initP256()
								initP384()
								initP521()
//line /usr/local/go/src/crypto/elliptic/elliptic.go:225
	// _ = "end of CoverTab[6997]"
}

// P224 returns a Curve which implements NIST P-224 (FIPS 186-3, section D.2.2),
//line /usr/local/go/src/crypto/elliptic/elliptic.go:228
// also known as secp224r1. The CurveParams.Name of this Curve is "P-224".
//line /usr/local/go/src/crypto/elliptic/elliptic.go:228
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:228
// Multiple invocations of this function will return the same value, so it can
//line /usr/local/go/src/crypto/elliptic/elliptic.go:228
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:228
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:228
// The cryptographic operations are implemented using constant-time algorithms.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:235
func P224() Curve {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:235
	_go_fuzz_dep_.CoverTab[6998]++
								initonce.Do(initAll)
								return p224
//line /usr/local/go/src/crypto/elliptic/elliptic.go:237
	// _ = "end of CoverTab[6998]"
}

// P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3),
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
// also known as secp256r1 or prime256v1. The CurveParams.Name of this Curve is
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
// "P-256".
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
// Multiple invocations of this function will return the same value, so it can
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:240
// The cryptographic operations are implemented using constant-time algorithms.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:248
func P256() Curve {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:248
	_go_fuzz_dep_.CoverTab[6999]++
								initonce.Do(initAll)
								return p256
//line /usr/local/go/src/crypto/elliptic/elliptic.go:250
	// _ = "end of CoverTab[6999]"
}

// P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4),
//line /usr/local/go/src/crypto/elliptic/elliptic.go:253
// also known as secp384r1. The CurveParams.Name of this Curve is "P-384".
//line /usr/local/go/src/crypto/elliptic/elliptic.go:253
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:253
// Multiple invocations of this function will return the same value, so it can
//line /usr/local/go/src/crypto/elliptic/elliptic.go:253
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:253
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:253
// The cryptographic operations are implemented using constant-time algorithms.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:260
func P384() Curve {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:260
	_go_fuzz_dep_.CoverTab[7000]++
								initonce.Do(initAll)
								return p384
//line /usr/local/go/src/crypto/elliptic/elliptic.go:262
	// _ = "end of CoverTab[7000]"
}

// P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5),
//line /usr/local/go/src/crypto/elliptic/elliptic.go:265
// also known as secp521r1. The CurveParams.Name of this Curve is "P-521".
//line /usr/local/go/src/crypto/elliptic/elliptic.go:265
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:265
// Multiple invocations of this function will return the same value, so it can
//line /usr/local/go/src/crypto/elliptic/elliptic.go:265
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:265
//
//line /usr/local/go/src/crypto/elliptic/elliptic.go:265
// The cryptographic operations are implemented using constant-time algorithms.
//line /usr/local/go/src/crypto/elliptic/elliptic.go:272
func P521() Curve {
//line /usr/local/go/src/crypto/elliptic/elliptic.go:272
	_go_fuzz_dep_.CoverTab[7001]++
								initonce.Do(initAll)
								return p521
//line /usr/local/go/src/crypto/elliptic/elliptic.go:274
	// _ = "end of CoverTab[7001]"
}

//line /usr/local/go/src/crypto/elliptic/elliptic.go:275
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/elliptic/elliptic.go:275
var _ = _go_fuzz_dep_.CoverTab
