// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/ecdh/nist.go:5
package ecdh

//line /usr/local/go/src/crypto/ecdh/nist.go:5
import (
//line /usr/local/go/src/crypto/ecdh/nist.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/ecdh/nist.go:5
)
//line /usr/local/go/src/crypto/ecdh/nist.go:5
import (
//line /usr/local/go/src/crypto/ecdh/nist.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/ecdh/nist.go:5
)

import (
	"crypto/internal/boring"
	"crypto/internal/nistec"
	"crypto/internal/randutil"
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
)

type nistCurve[Point nistPoint[Point]] struct {
	name		string
	newPoint	func() Point
	scalarOrder	[]byte
}

// nistPoint is a generic constraint for the nistec Point types.
type nistPoint[T any] interface {
	Bytes() []byte
	BytesX() ([]byte, error)
	SetBytes([]byte) (T, error)
	ScalarMult(T, []byte) (T, error)
	ScalarBaseMult([]byte) (T, error)
}

func (c *nistCurve[Point]) String() string {
//line /usr/local/go/src/crypto/ecdh/nist.go:32
	_go_fuzz_dep_.CoverTab[2803]++
							return c.name
//line /usr/local/go/src/crypto/ecdh/nist.go:33
	// _ = "end of CoverTab[2803]"
}

var errInvalidPrivateKey = errors.New("crypto/ecdh: invalid private key")

func (c *nistCurve[Point]) GenerateKey(rand io.Reader) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/ecdh/nist.go:38
	_go_fuzz_dep_.CoverTab[2804]++
							if boring.Enabled && func() bool {
//line /usr/local/go/src/crypto/ecdh/nist.go:39
		_go_fuzz_dep_.CoverTab[2806]++
//line /usr/local/go/src/crypto/ecdh/nist.go:39
		return rand == boring.RandReader
//line /usr/local/go/src/crypto/ecdh/nist.go:39
		// _ = "end of CoverTab[2806]"
//line /usr/local/go/src/crypto/ecdh/nist.go:39
	}() {
//line /usr/local/go/src/crypto/ecdh/nist.go:39
		_go_fuzz_dep_.CoverTab[2807]++
								key, bytes, err := boring.GenerateKeyECDH(c.name)
								if err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:41
			_go_fuzz_dep_.CoverTab[2809]++
									return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:42
			// _ = "end of CoverTab[2809]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:43
			_go_fuzz_dep_.CoverTab[2810]++
//line /usr/local/go/src/crypto/ecdh/nist.go:43
			// _ = "end of CoverTab[2810]"
//line /usr/local/go/src/crypto/ecdh/nist.go:43
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:43
		// _ = "end of CoverTab[2807]"
//line /usr/local/go/src/crypto/ecdh/nist.go:43
		_go_fuzz_dep_.CoverTab[2808]++
								return newBoringPrivateKey(c, key, bytes)
//line /usr/local/go/src/crypto/ecdh/nist.go:44
		// _ = "end of CoverTab[2808]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:45
		_go_fuzz_dep_.CoverTab[2811]++
//line /usr/local/go/src/crypto/ecdh/nist.go:45
		// _ = "end of CoverTab[2811]"
//line /usr/local/go/src/crypto/ecdh/nist.go:45
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:45
	// _ = "end of CoverTab[2804]"
//line /usr/local/go/src/crypto/ecdh/nist.go:45
	_go_fuzz_dep_.CoverTab[2805]++

							key := make([]byte, len(c.scalarOrder))
							randutil.MaybeReadByte(rand)
							for {
//line /usr/local/go/src/crypto/ecdh/nist.go:49
		_go_fuzz_dep_.CoverTab[2812]++
								if _, err := io.ReadFull(rand, key); err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:50
			_go_fuzz_dep_.CoverTab[2816]++
									return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:51
			// _ = "end of CoverTab[2816]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:52
			_go_fuzz_dep_.CoverTab[2817]++
//line /usr/local/go/src/crypto/ecdh/nist.go:52
			// _ = "end of CoverTab[2817]"
//line /usr/local/go/src/crypto/ecdh/nist.go:52
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:52
		// _ = "end of CoverTab[2812]"
//line /usr/local/go/src/crypto/ecdh/nist.go:52
		_go_fuzz_dep_.CoverTab[2813]++

//line /usr/local/go/src/crypto/ecdh/nist.go:58
		if &c.scalarOrder[0] == &p521Order[0] {
//line /usr/local/go/src/crypto/ecdh/nist.go:58
			_go_fuzz_dep_.CoverTab[2818]++
									key[0] &= 0b0000_0001
//line /usr/local/go/src/crypto/ecdh/nist.go:59
			// _ = "end of CoverTab[2818]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:60
			_go_fuzz_dep_.CoverTab[2819]++
//line /usr/local/go/src/crypto/ecdh/nist.go:60
			// _ = "end of CoverTab[2819]"
//line /usr/local/go/src/crypto/ecdh/nist.go:60
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:60
		// _ = "end of CoverTab[2813]"
//line /usr/local/go/src/crypto/ecdh/nist.go:60
		_go_fuzz_dep_.CoverTab[2814]++

//line /usr/local/go/src/crypto/ecdh/nist.go:65
		key[1] ^= 0x42

		k, err := c.NewPrivateKey(key)
		if err == errInvalidPrivateKey {
//line /usr/local/go/src/crypto/ecdh/nist.go:68
			_go_fuzz_dep_.CoverTab[2820]++
									continue
//line /usr/local/go/src/crypto/ecdh/nist.go:69
			// _ = "end of CoverTab[2820]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:70
			_go_fuzz_dep_.CoverTab[2821]++
//line /usr/local/go/src/crypto/ecdh/nist.go:70
			// _ = "end of CoverTab[2821]"
//line /usr/local/go/src/crypto/ecdh/nist.go:70
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:70
		// _ = "end of CoverTab[2814]"
//line /usr/local/go/src/crypto/ecdh/nist.go:70
		_go_fuzz_dep_.CoverTab[2815]++
								return k, err
//line /usr/local/go/src/crypto/ecdh/nist.go:71
		// _ = "end of CoverTab[2815]"
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:72
	// _ = "end of CoverTab[2805]"
}

func (c *nistCurve[Point]) NewPrivateKey(key []byte) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/ecdh/nist.go:75
	_go_fuzz_dep_.CoverTab[2822]++
							if len(key) != len(c.scalarOrder) {
//line /usr/local/go/src/crypto/ecdh/nist.go:76
		_go_fuzz_dep_.CoverTab[2826]++
								return nil, errors.New("crypto/ecdh: invalid private key size")
//line /usr/local/go/src/crypto/ecdh/nist.go:77
		// _ = "end of CoverTab[2826]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:78
		_go_fuzz_dep_.CoverTab[2827]++
//line /usr/local/go/src/crypto/ecdh/nist.go:78
		// _ = "end of CoverTab[2827]"
//line /usr/local/go/src/crypto/ecdh/nist.go:78
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:78
	// _ = "end of CoverTab[2822]"
//line /usr/local/go/src/crypto/ecdh/nist.go:78
	_go_fuzz_dep_.CoverTab[2823]++
							if isZero(key) || func() bool {
//line /usr/local/go/src/crypto/ecdh/nist.go:79
		_go_fuzz_dep_.CoverTab[2828]++
//line /usr/local/go/src/crypto/ecdh/nist.go:79
		return !isLess(key, c.scalarOrder)
//line /usr/local/go/src/crypto/ecdh/nist.go:79
		// _ = "end of CoverTab[2828]"
//line /usr/local/go/src/crypto/ecdh/nist.go:79
	}() {
//line /usr/local/go/src/crypto/ecdh/nist.go:79
		_go_fuzz_dep_.CoverTab[2829]++
								return nil, errInvalidPrivateKey
//line /usr/local/go/src/crypto/ecdh/nist.go:80
		// _ = "end of CoverTab[2829]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:81
		_go_fuzz_dep_.CoverTab[2830]++
//line /usr/local/go/src/crypto/ecdh/nist.go:81
		// _ = "end of CoverTab[2830]"
//line /usr/local/go/src/crypto/ecdh/nist.go:81
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:81
	// _ = "end of CoverTab[2823]"
//line /usr/local/go/src/crypto/ecdh/nist.go:81
	_go_fuzz_dep_.CoverTab[2824]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/ecdh/nist.go:82
		_go_fuzz_dep_.CoverTab[2831]++
								bk, err := boring.NewPrivateKeyECDH(c.name, key)
								if err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:84
			_go_fuzz_dep_.CoverTab[2833]++
									return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:85
			// _ = "end of CoverTab[2833]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:86
			_go_fuzz_dep_.CoverTab[2834]++
//line /usr/local/go/src/crypto/ecdh/nist.go:86
			// _ = "end of CoverTab[2834]"
//line /usr/local/go/src/crypto/ecdh/nist.go:86
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:86
		// _ = "end of CoverTab[2831]"
//line /usr/local/go/src/crypto/ecdh/nist.go:86
		_go_fuzz_dep_.CoverTab[2832]++
								return newBoringPrivateKey(c, bk, key)
//line /usr/local/go/src/crypto/ecdh/nist.go:87
		// _ = "end of CoverTab[2832]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:88
		_go_fuzz_dep_.CoverTab[2835]++
//line /usr/local/go/src/crypto/ecdh/nist.go:88
		// _ = "end of CoverTab[2835]"
//line /usr/local/go/src/crypto/ecdh/nist.go:88
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:88
	// _ = "end of CoverTab[2824]"
//line /usr/local/go/src/crypto/ecdh/nist.go:88
	_go_fuzz_dep_.CoverTab[2825]++
							k := &PrivateKey{
		curve:		c,
		privateKey:	append([]byte{}, key...),
	}
							return k, nil
//line /usr/local/go/src/crypto/ecdh/nist.go:93
	// _ = "end of CoverTab[2825]"
}

func newBoringPrivateKey(c Curve, bk *boring.PrivateKeyECDH, privateKey []byte) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/ecdh/nist.go:96
	_go_fuzz_dep_.CoverTab[2836]++
							k := &PrivateKey{
		curve:		c,
		boring:		bk,
		privateKey:	append([]byte(nil), privateKey...),
	}
							return k, nil
//line /usr/local/go/src/crypto/ecdh/nist.go:102
	// _ = "end of CoverTab[2836]"
}

func (c *nistCurve[Point]) privateKeyToPublicKey(key *PrivateKey) *PublicKey {
//line /usr/local/go/src/crypto/ecdh/nist.go:105
	_go_fuzz_dep_.CoverTab[2837]++
							boring.Unreachable()
							if key.curve != c {
//line /usr/local/go/src/crypto/ecdh/nist.go:107
		_go_fuzz_dep_.CoverTab[2841]++
								panic("crypto/ecdh: internal error: converting the wrong key type")
//line /usr/local/go/src/crypto/ecdh/nist.go:108
		// _ = "end of CoverTab[2841]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:109
		_go_fuzz_dep_.CoverTab[2842]++
//line /usr/local/go/src/crypto/ecdh/nist.go:109
		// _ = "end of CoverTab[2842]"
//line /usr/local/go/src/crypto/ecdh/nist.go:109
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:109
	// _ = "end of CoverTab[2837]"
//line /usr/local/go/src/crypto/ecdh/nist.go:109
	_go_fuzz_dep_.CoverTab[2838]++
							p, err := c.newPoint().ScalarBaseMult(key.privateKey)
							if err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:111
		_go_fuzz_dep_.CoverTab[2843]++

//line /usr/local/go/src/crypto/ecdh/nist.go:114
		panic("crypto/ecdh: internal error: nistec ScalarBaseMult failed for a fixed-size input")
//line /usr/local/go/src/crypto/ecdh/nist.go:114
		// _ = "end of CoverTab[2843]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:115
		_go_fuzz_dep_.CoverTab[2844]++
//line /usr/local/go/src/crypto/ecdh/nist.go:115
		// _ = "end of CoverTab[2844]"
//line /usr/local/go/src/crypto/ecdh/nist.go:115
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:115
	// _ = "end of CoverTab[2838]"
//line /usr/local/go/src/crypto/ecdh/nist.go:115
	_go_fuzz_dep_.CoverTab[2839]++
							publicKey := p.Bytes()
							if len(publicKey) == 1 {
//line /usr/local/go/src/crypto/ecdh/nist.go:117
		_go_fuzz_dep_.CoverTab[2845]++

//line /usr/local/go/src/crypto/ecdh/nist.go:121
		panic("crypto/ecdh: internal error: nistec ScalarBaseMult returned the identity")
//line /usr/local/go/src/crypto/ecdh/nist.go:121
		// _ = "end of CoverTab[2845]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:122
		_go_fuzz_dep_.CoverTab[2846]++
//line /usr/local/go/src/crypto/ecdh/nist.go:122
		// _ = "end of CoverTab[2846]"
//line /usr/local/go/src/crypto/ecdh/nist.go:122
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:122
	// _ = "end of CoverTab[2839]"
//line /usr/local/go/src/crypto/ecdh/nist.go:122
	_go_fuzz_dep_.CoverTab[2840]++
							return &PublicKey{
		curve:		key.curve,
		publicKey:	publicKey,
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:126
	// _ = "end of CoverTab[2840]"
}

// isZero returns whether a is all zeroes in constant time.
func isZero(a []byte) bool {
//line /usr/local/go/src/crypto/ecdh/nist.go:130
	_go_fuzz_dep_.CoverTab[2847]++
							var acc byte
							for _, b := range a {
//line /usr/local/go/src/crypto/ecdh/nist.go:132
		_go_fuzz_dep_.CoverTab[2849]++
								acc |= b
//line /usr/local/go/src/crypto/ecdh/nist.go:133
		// _ = "end of CoverTab[2849]"
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:134
	// _ = "end of CoverTab[2847]"
//line /usr/local/go/src/crypto/ecdh/nist.go:134
	_go_fuzz_dep_.CoverTab[2848]++
							return acc == 0
//line /usr/local/go/src/crypto/ecdh/nist.go:135
	// _ = "end of CoverTab[2848]"
}

// isLess returns whether a < b, where a and b are big-endian buffers of the
//line /usr/local/go/src/crypto/ecdh/nist.go:138
// same length and shorter than 72 bytes.
//line /usr/local/go/src/crypto/ecdh/nist.go:140
func isLess(a, b []byte) bool {
//line /usr/local/go/src/crypto/ecdh/nist.go:140
	_go_fuzz_dep_.CoverTab[2850]++
							if len(a) != len(b) {
//line /usr/local/go/src/crypto/ecdh/nist.go:141
		_go_fuzz_dep_.CoverTab[2855]++
								panic("crypto/ecdh: internal error: mismatched isLess inputs")
//line /usr/local/go/src/crypto/ecdh/nist.go:142
		// _ = "end of CoverTab[2855]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:143
		_go_fuzz_dep_.CoverTab[2856]++
//line /usr/local/go/src/crypto/ecdh/nist.go:143
		// _ = "end of CoverTab[2856]"
//line /usr/local/go/src/crypto/ecdh/nist.go:143
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:143
	// _ = "end of CoverTab[2850]"
//line /usr/local/go/src/crypto/ecdh/nist.go:143
	_go_fuzz_dep_.CoverTab[2851]++

//line /usr/local/go/src/crypto/ecdh/nist.go:148
	if len(a) > 72 {
//line /usr/local/go/src/crypto/ecdh/nist.go:148
		_go_fuzz_dep_.CoverTab[2857]++
								panic("crypto/ecdh: internal error: isLess input too large")
//line /usr/local/go/src/crypto/ecdh/nist.go:149
		// _ = "end of CoverTab[2857]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:150
		_go_fuzz_dep_.CoverTab[2858]++
//line /usr/local/go/src/crypto/ecdh/nist.go:150
		// _ = "end of CoverTab[2858]"
//line /usr/local/go/src/crypto/ecdh/nist.go:150
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:150
	// _ = "end of CoverTab[2851]"
//line /usr/local/go/src/crypto/ecdh/nist.go:150
	_go_fuzz_dep_.CoverTab[2852]++
							bufA, bufB := make([]byte, 72), make([]byte, 72)
							for i := range a {
//line /usr/local/go/src/crypto/ecdh/nist.go:152
		_go_fuzz_dep_.CoverTab[2859]++
								bufA[i], bufB[i] = a[len(a)-i-1], b[len(b)-i-1]
//line /usr/local/go/src/crypto/ecdh/nist.go:153
		// _ = "end of CoverTab[2859]"
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:154
	// _ = "end of CoverTab[2852]"
//line /usr/local/go/src/crypto/ecdh/nist.go:154
	_go_fuzz_dep_.CoverTab[2853]++

	// Perform a subtraction with borrow.
	var borrow uint64
	for i := 0; i < len(bufA); i += 8 {
//line /usr/local/go/src/crypto/ecdh/nist.go:158
		_go_fuzz_dep_.CoverTab[2860]++
								limbA, limbB := binary.LittleEndian.Uint64(bufA[i:]), binary.LittleEndian.Uint64(bufB[i:])
								_, borrow = bits.Sub64(limbA, limbB, borrow)
//line /usr/local/go/src/crypto/ecdh/nist.go:160
		// _ = "end of CoverTab[2860]"
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:161
	// _ = "end of CoverTab[2853]"
//line /usr/local/go/src/crypto/ecdh/nist.go:161
	_go_fuzz_dep_.CoverTab[2854]++

//line /usr/local/go/src/crypto/ecdh/nist.go:164
	return borrow == 1
//line /usr/local/go/src/crypto/ecdh/nist.go:164
	// _ = "end of CoverTab[2854]"
}

func (c *nistCurve[Point]) NewPublicKey(key []byte) (*PublicKey, error) {
//line /usr/local/go/src/crypto/ecdh/nist.go:167
	_go_fuzz_dep_.CoverTab[2861]++

							if len(key) == 0 || func() bool {
//line /usr/local/go/src/crypto/ecdh/nist.go:169
		_go_fuzz_dep_.CoverTab[2864]++
//line /usr/local/go/src/crypto/ecdh/nist.go:169
		return key[0] != 4
//line /usr/local/go/src/crypto/ecdh/nist.go:169
		// _ = "end of CoverTab[2864]"
//line /usr/local/go/src/crypto/ecdh/nist.go:169
	}() {
//line /usr/local/go/src/crypto/ecdh/nist.go:169
		_go_fuzz_dep_.CoverTab[2865]++
								return nil, errors.New("crypto/ecdh: invalid public key")
//line /usr/local/go/src/crypto/ecdh/nist.go:170
		// _ = "end of CoverTab[2865]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:171
		_go_fuzz_dep_.CoverTab[2866]++
//line /usr/local/go/src/crypto/ecdh/nist.go:171
		// _ = "end of CoverTab[2866]"
//line /usr/local/go/src/crypto/ecdh/nist.go:171
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:171
	// _ = "end of CoverTab[2861]"
//line /usr/local/go/src/crypto/ecdh/nist.go:171
	_go_fuzz_dep_.CoverTab[2862]++
							k := &PublicKey{
		curve:		c,
		publicKey:	append([]byte{}, key...),
	}
	if boring.Enabled {
//line /usr/local/go/src/crypto/ecdh/nist.go:176
		_go_fuzz_dep_.CoverTab[2867]++
								bk, err := boring.NewPublicKeyECDH(c.name, k.publicKey)
								if err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:178
			_go_fuzz_dep_.CoverTab[2869]++
									return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:179
			// _ = "end of CoverTab[2869]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:180
			_go_fuzz_dep_.CoverTab[2870]++
//line /usr/local/go/src/crypto/ecdh/nist.go:180
			// _ = "end of CoverTab[2870]"
//line /usr/local/go/src/crypto/ecdh/nist.go:180
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:180
		// _ = "end of CoverTab[2867]"
//line /usr/local/go/src/crypto/ecdh/nist.go:180
		_go_fuzz_dep_.CoverTab[2868]++
								k.boring = bk
//line /usr/local/go/src/crypto/ecdh/nist.go:181
		// _ = "end of CoverTab[2868]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:182
		_go_fuzz_dep_.CoverTab[2871]++

								if _, err := c.newPoint().SetBytes(key); err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:184
			_go_fuzz_dep_.CoverTab[2872]++
									return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:185
			// _ = "end of CoverTab[2872]"
		} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:186
			_go_fuzz_dep_.CoverTab[2873]++
//line /usr/local/go/src/crypto/ecdh/nist.go:186
			// _ = "end of CoverTab[2873]"
//line /usr/local/go/src/crypto/ecdh/nist.go:186
		}
//line /usr/local/go/src/crypto/ecdh/nist.go:186
		// _ = "end of CoverTab[2871]"
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:187
	// _ = "end of CoverTab[2862]"
//line /usr/local/go/src/crypto/ecdh/nist.go:187
	_go_fuzz_dep_.CoverTab[2863]++
							return k, nil
//line /usr/local/go/src/crypto/ecdh/nist.go:188
	// _ = "end of CoverTab[2863]"
}

func (c *nistCurve[Point]) ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error) {
//line /usr/local/go/src/crypto/ecdh/nist.go:191
	_go_fuzz_dep_.CoverTab[2874]++

//line /usr/local/go/src/crypto/ecdh/nist.go:199
	if boring.Enabled {
//line /usr/local/go/src/crypto/ecdh/nist.go:199
		_go_fuzz_dep_.CoverTab[2878]++
								return boring.ECDH(local.boring, remote.boring)
//line /usr/local/go/src/crypto/ecdh/nist.go:200
		// _ = "end of CoverTab[2878]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:201
		_go_fuzz_dep_.CoverTab[2879]++
//line /usr/local/go/src/crypto/ecdh/nist.go:201
		// _ = "end of CoverTab[2879]"
//line /usr/local/go/src/crypto/ecdh/nist.go:201
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:201
	// _ = "end of CoverTab[2874]"
//line /usr/local/go/src/crypto/ecdh/nist.go:201
	_go_fuzz_dep_.CoverTab[2875]++

							boring.Unreachable()
							p, err := c.newPoint().SetBytes(remote.publicKey)
							if err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:205
		_go_fuzz_dep_.CoverTab[2880]++
								return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:206
		// _ = "end of CoverTab[2880]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:207
		_go_fuzz_dep_.CoverTab[2881]++
//line /usr/local/go/src/crypto/ecdh/nist.go:207
		// _ = "end of CoverTab[2881]"
//line /usr/local/go/src/crypto/ecdh/nist.go:207
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:207
	// _ = "end of CoverTab[2875]"
//line /usr/local/go/src/crypto/ecdh/nist.go:207
	_go_fuzz_dep_.CoverTab[2876]++
							if _, err := p.ScalarMult(p, local.privateKey); err != nil {
//line /usr/local/go/src/crypto/ecdh/nist.go:208
		_go_fuzz_dep_.CoverTab[2882]++
								return nil, err
//line /usr/local/go/src/crypto/ecdh/nist.go:209
		// _ = "end of CoverTab[2882]"
	} else {
//line /usr/local/go/src/crypto/ecdh/nist.go:210
		_go_fuzz_dep_.CoverTab[2883]++
//line /usr/local/go/src/crypto/ecdh/nist.go:210
		// _ = "end of CoverTab[2883]"
//line /usr/local/go/src/crypto/ecdh/nist.go:210
	}
//line /usr/local/go/src/crypto/ecdh/nist.go:210
	// _ = "end of CoverTab[2876]"
//line /usr/local/go/src/crypto/ecdh/nist.go:210
	_go_fuzz_dep_.CoverTab[2877]++
							return p.BytesX()
//line /usr/local/go/src/crypto/ecdh/nist.go:211
	// _ = "end of CoverTab[2877]"
}

// P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3),
//line /usr/local/go/src/crypto/ecdh/nist.go:214
// also known as secp256r1 or prime256v1.
//line /usr/local/go/src/crypto/ecdh/nist.go:214
//
//line /usr/local/go/src/crypto/ecdh/nist.go:214
// Multiple invocations of this function will return the same value, which can
//line /usr/local/go/src/crypto/ecdh/nist.go:214
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/ecdh/nist.go:219
func P256() Curve	{ _go_fuzz_dep_.CoverTab[2884]++; return p256; // _ = "end of CoverTab[2884]" }

var p256 = &nistCurve[*nistec.P256Point]{
	name:		"P-256",
	newPoint:	nistec.NewP256Point,
	scalarOrder:	p256Order,
}

var p256Order = []byte{
	0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xbc, 0xe6, 0xfa, 0xad, 0xa7, 0x17, 0x9e, 0x84,
	0xf3, 0xb9, 0xca, 0xc2, 0xfc, 0x63, 0x25, 0x51}

// P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4),
//line /usr/local/go/src/crypto/ecdh/nist.go:233
// also known as secp384r1.
//line /usr/local/go/src/crypto/ecdh/nist.go:233
//
//line /usr/local/go/src/crypto/ecdh/nist.go:233
// Multiple invocations of this function will return the same value, which can
//line /usr/local/go/src/crypto/ecdh/nist.go:233
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/ecdh/nist.go:238
func P384() Curve	{ _go_fuzz_dep_.CoverTab[2885]++; return p384; // _ = "end of CoverTab[2885]" }

var p384 = &nistCurve[*nistec.P384Point]{
	name:		"P-384",
	newPoint:	nistec.NewP384Point,
	scalarOrder:	p384Order,
}

var p384Order = []byte{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xc7, 0x63, 0x4d, 0x81, 0xf4, 0x37, 0x2d, 0xdf,
	0x58, 0x1a, 0x0d, 0xb2, 0x48, 0xb0, 0xa7, 0x7a,
	0xec, 0xec, 0x19, 0x6a, 0xcc, 0xc5, 0x29, 0x73}

// P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5),
//line /usr/local/go/src/crypto/ecdh/nist.go:254
// also known as secp521r1.
//line /usr/local/go/src/crypto/ecdh/nist.go:254
//
//line /usr/local/go/src/crypto/ecdh/nist.go:254
// Multiple invocations of this function will return the same value, which can
//line /usr/local/go/src/crypto/ecdh/nist.go:254
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/ecdh/nist.go:259
func P521() Curve	{ _go_fuzz_dep_.CoverTab[2886]++; return p521; // _ = "end of CoverTab[2886]" }

var p521 = &nistCurve[*nistec.P521Point]{
	name:		"P-521",
	newPoint:	nistec.NewP521Point,
	scalarOrder:	p521Order,
}

var p521Order = []byte{0x01, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfa,
	0x51, 0x86, 0x87, 0x83, 0xbf, 0x2f, 0x96, 0x6b,
	0x7f, 0xcc, 0x01, 0x48, 0xf7, 0x09, 0xa5, 0xd0,
	0x3b, 0xb5, 0xc9, 0xb8, 0x89, 0x9c, 0x47, 0xae,
	0xbb, 0x6f, 0xb7, 0x1e, 0x91, 0x38, 0x64, 0x09}
//line /usr/local/go/src/crypto/ecdh/nist.go:275
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/ecdh/nist.go:275
var _ = _go_fuzz_dep_.CoverTab
