// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
package ecdsa

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
import (
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
)
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
import (
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:5
)

import (
	"crypto/elliptic"
	"errors"
	"io"
	"math/big"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:20
func generateLegacy(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:20
	_go_fuzz_dep_.CoverTab[9077]++
								k, err := randFieldElement(c, rand)
								if err != nil {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:22
		_go_fuzz_dep_.CoverTab[9079]++
									return nil, err
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:23
		// _ = "end of CoverTab[9079]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:24
		_go_fuzz_dep_.CoverTab[9080]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:24
		// _ = "end of CoverTab[9080]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:24
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:24
	// _ = "end of CoverTab[9077]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:24
	_go_fuzz_dep_.CoverTab[9078]++

								priv := new(PrivateKey)
								priv.PublicKey.Curve = c
								priv.D = k
								priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
								return priv, nil
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:30
	// _ = "end of CoverTab[9078]"
}

// hashToInt converts a hash value to an integer. Per FIPS 186-4, Section 6.4,
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:33
// we use the left-most bits of the hash to match the bit-length of the order of
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:33
// the curve. This also performs Step 5 of SEC 1, Version 2.0, Section 4.1.3.
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:36
func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:36
	_go_fuzz_dep_.CoverTab[9081]++
								orderBits := c.Params().N.BitLen()
								orderBytes := (orderBits + 7) / 8
								if len(hash) > orderBytes {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:39
		_go_fuzz_dep_.CoverTab[9084]++
									hash = hash[:orderBytes]
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:40
		// _ = "end of CoverTab[9084]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:41
		_go_fuzz_dep_.CoverTab[9085]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:41
		// _ = "end of CoverTab[9085]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:41
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:41
	// _ = "end of CoverTab[9081]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:41
	_go_fuzz_dep_.CoverTab[9082]++

								ret := new(big.Int).SetBytes(hash)
								excess := len(hash)*8 - orderBits
								if excess > 0 {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:45
		_go_fuzz_dep_.CoverTab[9086]++
									ret.Rsh(ret, uint(excess))
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:46
		// _ = "end of CoverTab[9086]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:47
		_go_fuzz_dep_.CoverTab[9087]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:47
		// _ = "end of CoverTab[9087]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:47
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:47
	// _ = "end of CoverTab[9082]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:47
	_go_fuzz_dep_.CoverTab[9083]++
								return ret
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:48
	// _ = "end of CoverTab[9083]"
}

var errZeroParam = errors.New("zero parameter")

// Sign signs a hash (which should be the result of hashing a larger message)
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:53
// using the private key, priv. If the hash is longer than the bit-length of the
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:53
// private key's curve order, the hash will be truncated to that length. It
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:53
// returns the signature as a pair of integers. Most applications should use
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:53
// SignASN1 instead of dealing directly with r, s.
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:58
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error) {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:58
	_go_fuzz_dep_.CoverTab[9088]++
								sig, err := SignASN1(rand, priv, hash)
								if err != nil {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:60
		_go_fuzz_dep_.CoverTab[9091]++
									return nil, nil, err
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:61
		// _ = "end of CoverTab[9091]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:62
		_go_fuzz_dep_.CoverTab[9092]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:62
		// _ = "end of CoverTab[9092]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:62
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:62
	// _ = "end of CoverTab[9088]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:62
	_go_fuzz_dep_.CoverTab[9089]++

								r, s = new(big.Int), new(big.Int)
								var inner cryptobyte.String
								input := cryptobyte.String(sig)
								if !input.ReadASN1(&inner, asn1.SEQUENCE) || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:67
		_go_fuzz_dep_.CoverTab[9093]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:67
		return !input.Empty()
									// _ = "end of CoverTab[9093]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:68
	}() || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:68
		_go_fuzz_dep_.CoverTab[9094]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:68
		return !inner.ReadASN1Integer(r)
									// _ = "end of CoverTab[9094]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:69
	}() || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:69
		_go_fuzz_dep_.CoverTab[9095]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:69
		return !inner.ReadASN1Integer(s)
									// _ = "end of CoverTab[9095]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:70
	}() || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:70
		_go_fuzz_dep_.CoverTab[9096]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:70
		return !inner.Empty()
									// _ = "end of CoverTab[9096]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:71
	}() {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:71
		_go_fuzz_dep_.CoverTab[9097]++
									return nil, nil, errors.New("invalid ASN.1 from SignASN1")
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:72
		// _ = "end of CoverTab[9097]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:73
		_go_fuzz_dep_.CoverTab[9098]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:73
		// _ = "end of CoverTab[9098]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:73
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:73
	// _ = "end of CoverTab[9089]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:73
	_go_fuzz_dep_.CoverTab[9090]++
								return r, s, nil
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:74
	// _ = "end of CoverTab[9090]"
}

func signLegacy(priv *PrivateKey, csprng io.Reader, hash []byte) (sig []byte, err error) {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:77
	_go_fuzz_dep_.CoverTab[9099]++
								c := priv.Curve

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:81
	N := c.Params().N
	if N.Sign() == 0 {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:82
		_go_fuzz_dep_.CoverTab[9102]++
									return nil, errZeroParam
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:83
		// _ = "end of CoverTab[9102]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:84
		_go_fuzz_dep_.CoverTab[9103]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:84
		// _ = "end of CoverTab[9103]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:84
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:84
	// _ = "end of CoverTab[9099]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:84
	_go_fuzz_dep_.CoverTab[9100]++
								var k, kInv, r, s *big.Int
								for {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:86
		_go_fuzz_dep_.CoverTab[9104]++
									for {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:87
			_go_fuzz_dep_.CoverTab[9106]++
										k, err = randFieldElement(c, csprng)
										if err != nil {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:89
				_go_fuzz_dep_.CoverTab[9108]++
											return nil, err
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:90
				// _ = "end of CoverTab[9108]"
			} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:91
				_go_fuzz_dep_.CoverTab[9109]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:91
				// _ = "end of CoverTab[9109]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:91
			}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:91
			// _ = "end of CoverTab[9106]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:91
			_go_fuzz_dep_.CoverTab[9107]++

										kInv = new(big.Int).ModInverse(k, N)

										r, _ = c.ScalarBaseMult(k.Bytes())
										r.Mod(r, N)
										if r.Sign() != 0 {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:97
				_go_fuzz_dep_.CoverTab[9110]++
											break
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:98
				// _ = "end of CoverTab[9110]"
			} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:99
				_go_fuzz_dep_.CoverTab[9111]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:99
				// _ = "end of CoverTab[9111]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:99
			}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:99
			// _ = "end of CoverTab[9107]"
		}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:100
		// _ = "end of CoverTab[9104]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:100
		_go_fuzz_dep_.CoverTab[9105]++

									e := hashToInt(hash, c)
									s = new(big.Int).Mul(priv.D, r)
									s.Add(s, e)
									s.Mul(s, kInv)
									s.Mod(s, N)
									if s.Sign() != 0 {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:107
			_go_fuzz_dep_.CoverTab[9112]++
										break
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:108
			// _ = "end of CoverTab[9112]"
		} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:109
			_go_fuzz_dep_.CoverTab[9113]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:109
			// _ = "end of CoverTab[9113]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:109
		}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:109
		// _ = "end of CoverTab[9105]"
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:110
	// _ = "end of CoverTab[9100]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:110
	_go_fuzz_dep_.CoverTab[9101]++

								return encodeSignature(r.Bytes(), s.Bytes())
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:112
	// _ = "end of CoverTab[9101]"
}

// Verify verifies the signature in r, s of hash using the public key, pub. Its
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:115
// return value records whether the signature is valid. Most applications should
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:115
// use VerifyASN1 instead of dealing directly with r, s.
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:118
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:118
	_go_fuzz_dep_.CoverTab[9114]++
								if r.Sign() <= 0 || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:119
		_go_fuzz_dep_.CoverTab[9117]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:119
		return s.Sign() <= 0
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:119
		// _ = "end of CoverTab[9117]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:119
	}() {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:119
		_go_fuzz_dep_.CoverTab[9118]++
									return false
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:120
		// _ = "end of CoverTab[9118]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:121
		_go_fuzz_dep_.CoverTab[9119]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:121
		// _ = "end of CoverTab[9119]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:121
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:121
	// _ = "end of CoverTab[9114]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:121
	_go_fuzz_dep_.CoverTab[9115]++
								sig, err := encodeSignature(r.Bytes(), s.Bytes())
								if err != nil {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:123
		_go_fuzz_dep_.CoverTab[9120]++
									return false
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:124
		// _ = "end of CoverTab[9120]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:125
		_go_fuzz_dep_.CoverTab[9121]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:125
		// _ = "end of CoverTab[9121]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:125
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:125
	// _ = "end of CoverTab[9115]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:125
	_go_fuzz_dep_.CoverTab[9116]++
								return VerifyASN1(pub, hash, sig)
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:126
	// _ = "end of CoverTab[9116]"
}

func verifyLegacy(pub *PublicKey, hash []byte, sig []byte) bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:129
	_go_fuzz_dep_.CoverTab[9122]++
								rBytes, sBytes, err := parseSignature(sig)
								if err != nil {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:131
		_go_fuzz_dep_.CoverTab[9127]++
									return false
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:132
		// _ = "end of CoverTab[9127]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:133
		_go_fuzz_dep_.CoverTab[9128]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:133
		// _ = "end of CoverTab[9128]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:133
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:133
	// _ = "end of CoverTab[9122]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:133
	_go_fuzz_dep_.CoverTab[9123]++
								r, s := new(big.Int).SetBytes(rBytes), new(big.Int).SetBytes(sBytes)

								c := pub.Curve
								N := c.Params().N

								if r.Sign() <= 0 || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:139
		_go_fuzz_dep_.CoverTab[9129]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:139
		return s.Sign() <= 0
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:139
		// _ = "end of CoverTab[9129]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:139
	}() {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:139
		_go_fuzz_dep_.CoverTab[9130]++
									return false
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:140
		// _ = "end of CoverTab[9130]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:141
		_go_fuzz_dep_.CoverTab[9131]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:141
		// _ = "end of CoverTab[9131]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:141
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:141
	// _ = "end of CoverTab[9123]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:141
	_go_fuzz_dep_.CoverTab[9124]++
								if r.Cmp(N) >= 0 || func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:142
		_go_fuzz_dep_.CoverTab[9132]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:142
		return s.Cmp(N) >= 0
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:142
		// _ = "end of CoverTab[9132]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:142
	}() {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:142
		_go_fuzz_dep_.CoverTab[9133]++
									return false
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:143
		// _ = "end of CoverTab[9133]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:144
		_go_fuzz_dep_.CoverTab[9134]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:144
		// _ = "end of CoverTab[9134]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:144
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:144
	// _ = "end of CoverTab[9124]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:144
	_go_fuzz_dep_.CoverTab[9125]++

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:147
	e := hashToInt(hash, c)
	w := new(big.Int).ModInverse(s, N)

	u1 := e.Mul(e, w)
	u1.Mod(u1, N)
	u2 := w.Mul(r, w)
	u2.Mod(u2, N)

	x1, y1 := c.ScalarBaseMult(u1.Bytes())
	x2, y2 := c.ScalarMult(pub.X, pub.Y, u2.Bytes())
	x, y := c.Add(x1, y1, x2, y2)

	if x.Sign() == 0 && func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:159
		_go_fuzz_dep_.CoverTab[9135]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:159
		return y.Sign() == 0
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:159
		// _ = "end of CoverTab[9135]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:159
	}() {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:159
		_go_fuzz_dep_.CoverTab[9136]++
									return false
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:160
		// _ = "end of CoverTab[9136]"
	} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:161
		_go_fuzz_dep_.CoverTab[9137]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:161
		// _ = "end of CoverTab[9137]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:161
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:161
	// _ = "end of CoverTab[9125]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:161
	_go_fuzz_dep_.CoverTab[9126]++
								x.Mod(x, N)
								return x.Cmp(r) == 0
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:163
	// _ = "end of CoverTab[9126]"
}

var one = new(big.Int).SetInt64(1)

// randFieldElement returns a random element of the order of the given
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:168
// curve using the procedure given in FIPS 186-4, Appendix B.5.2.
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:170
func randFieldElement(c elliptic.Curve, rand io.Reader) (k *big.Int, err error) {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:170
	_go_fuzz_dep_.CoverTab[9138]++

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:174
	for {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:174
		_go_fuzz_dep_.CoverTab[9139]++
									N := c.Params().N
									b := make([]byte, (N.BitLen()+7)/8)
									if _, err = io.ReadFull(rand, b); err != nil {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:177
			_go_fuzz_dep_.CoverTab[9142]++
										return
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:178
			// _ = "end of CoverTab[9142]"
		} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:179
			_go_fuzz_dep_.CoverTab[9143]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:179
			// _ = "end of CoverTab[9143]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:179
		}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:179
		// _ = "end of CoverTab[9139]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:179
		_go_fuzz_dep_.CoverTab[9140]++
									if excess := len(b)*8 - N.BitLen(); excess > 0 {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:180
			_go_fuzz_dep_.CoverTab[9144]++
										b[0] >>= excess
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:181
			// _ = "end of CoverTab[9144]"
		} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:182
			_go_fuzz_dep_.CoverTab[9145]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:182
			// _ = "end of CoverTab[9145]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:182
		}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:182
		// _ = "end of CoverTab[9140]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:182
		_go_fuzz_dep_.CoverTab[9141]++
									k = new(big.Int).SetBytes(b)
									if k.Sign() != 0 && func() bool {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:184
			_go_fuzz_dep_.CoverTab[9146]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:184
			return k.Cmp(N) < 0
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:184
			// _ = "end of CoverTab[9146]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:184
		}() {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:184
			_go_fuzz_dep_.CoverTab[9147]++
										return
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:185
			// _ = "end of CoverTab[9147]"
		} else {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:186
			_go_fuzz_dep_.CoverTab[9148]++
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:186
			// _ = "end of CoverTab[9148]"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:186
		}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:186
		// _ = "end of CoverTab[9141]"
	}
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:187
	// _ = "end of CoverTab[9138]"
}

//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:188
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/ecdsa/ecdsa_legacy.go:188
var _ = _go_fuzz_dep_.CoverTab
