// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/ecdh/x25519.go:5
package ecdh

//line /usr/local/go/src/crypto/ecdh/x25519.go:5
import (
//line /usr/local/go/src/crypto/ecdh/x25519.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/ecdh/x25519.go:5
)
//line /usr/local/go/src/crypto/ecdh/x25519.go:5
import (
//line /usr/local/go/src/crypto/ecdh/x25519.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/ecdh/x25519.go:5
)

import (
	"crypto/internal/edwards25519/field"
	"crypto/internal/randutil"
	"errors"
	"io"
)

var (
	x25519PublicKeySize	= 32
	x25519PrivateKeySize	= 32
	x25519SharedSecretSize	= 32
)

// X25519 returns a Curve which implements the X25519 function over Curve25519
//line /usr/local/go/src/crypto/ecdh/x25519.go:20
// (RFC 7748, Section 5).
//line /usr/local/go/src/crypto/ecdh/x25519.go:20
//
//line /usr/local/go/src/crypto/ecdh/x25519.go:20
// Multiple invocations of this function will return the same value, so it can
//line /usr/local/go/src/crypto/ecdh/x25519.go:20
// be used for equality checks and switch statements.
//line /usr/local/go/src/crypto/ecdh/x25519.go:25
func X25519() Curve	{ _go_fuzz_dep_.CoverTab[2887]++; return x25519; // _ = "end of CoverTab[2887]" }

var x25519 = &x25519Curve{}

type x25519Curve struct{}

func (c *x25519Curve) String() string {
//line /usr/local/go/src/crypto/ecdh/x25519.go:31
	_go_fuzz_dep_.CoverTab[2888]++
							return "X25519"
//line /usr/local/go/src/crypto/ecdh/x25519.go:32
	// _ = "end of CoverTab[2888]"
}

func (c *x25519Curve) GenerateKey(rand io.Reader) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/ecdh/x25519.go:35
	_go_fuzz_dep_.CoverTab[2889]++
							key := make([]byte, x25519PrivateKeySize)
							randutil.MaybeReadByte(rand)
							if _, err := io.ReadFull(rand, key); err != nil {
//line /usr/local/go/src/crypto/ecdh/x25519.go:38
		_go_fuzz_dep_.CoverTab[2891]++
								return nil, err
//line /usr/local/go/src/crypto/ecdh/x25519.go:39
		// _ = "end of CoverTab[2891]"
	} else {
//line /usr/local/go/src/crypto/ecdh/x25519.go:40
		_go_fuzz_dep_.CoverTab[2892]++
//line /usr/local/go/src/crypto/ecdh/x25519.go:40
		// _ = "end of CoverTab[2892]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:40
	}
//line /usr/local/go/src/crypto/ecdh/x25519.go:40
	// _ = "end of CoverTab[2889]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:40
	_go_fuzz_dep_.CoverTab[2890]++
							return c.NewPrivateKey(key)
//line /usr/local/go/src/crypto/ecdh/x25519.go:41
	// _ = "end of CoverTab[2890]"
}

func (c *x25519Curve) NewPrivateKey(key []byte) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/ecdh/x25519.go:44
	_go_fuzz_dep_.CoverTab[2893]++
							if len(key) != x25519PrivateKeySize {
//line /usr/local/go/src/crypto/ecdh/x25519.go:45
		_go_fuzz_dep_.CoverTab[2895]++
								return nil, errors.New("crypto/ecdh: invalid private key size")
//line /usr/local/go/src/crypto/ecdh/x25519.go:46
		// _ = "end of CoverTab[2895]"
	} else {
//line /usr/local/go/src/crypto/ecdh/x25519.go:47
		_go_fuzz_dep_.CoverTab[2896]++
//line /usr/local/go/src/crypto/ecdh/x25519.go:47
		// _ = "end of CoverTab[2896]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:47
	}
//line /usr/local/go/src/crypto/ecdh/x25519.go:47
	// _ = "end of CoverTab[2893]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:47
	_go_fuzz_dep_.CoverTab[2894]++
							return &PrivateKey{
		curve:		c,
		privateKey:	append([]byte{}, key...),
	}, nil
//line /usr/local/go/src/crypto/ecdh/x25519.go:51
	// _ = "end of CoverTab[2894]"
}

func (c *x25519Curve) privateKeyToPublicKey(key *PrivateKey) *PublicKey {
//line /usr/local/go/src/crypto/ecdh/x25519.go:54
	_go_fuzz_dep_.CoverTab[2897]++
							if key.curve != c {
//line /usr/local/go/src/crypto/ecdh/x25519.go:55
		_go_fuzz_dep_.CoverTab[2899]++
								panic("crypto/ecdh: internal error: converting the wrong key type")
//line /usr/local/go/src/crypto/ecdh/x25519.go:56
		// _ = "end of CoverTab[2899]"
	} else {
//line /usr/local/go/src/crypto/ecdh/x25519.go:57
		_go_fuzz_dep_.CoverTab[2900]++
//line /usr/local/go/src/crypto/ecdh/x25519.go:57
		// _ = "end of CoverTab[2900]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:57
	}
//line /usr/local/go/src/crypto/ecdh/x25519.go:57
	// _ = "end of CoverTab[2897]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:57
	_go_fuzz_dep_.CoverTab[2898]++
							k := &PublicKey{
		curve:		key.curve,
		publicKey:	make([]byte, x25519PublicKeySize),
	}
							x25519Basepoint := [32]byte{9}
							x25519ScalarMult(k.publicKey, key.privateKey, x25519Basepoint[:])
							return k
//line /usr/local/go/src/crypto/ecdh/x25519.go:64
	// _ = "end of CoverTab[2898]"
}

func (c *x25519Curve) NewPublicKey(key []byte) (*PublicKey, error) {
//line /usr/local/go/src/crypto/ecdh/x25519.go:67
	_go_fuzz_dep_.CoverTab[2901]++
							if len(key) != x25519PublicKeySize {
//line /usr/local/go/src/crypto/ecdh/x25519.go:68
		_go_fuzz_dep_.CoverTab[2903]++
								return nil, errors.New("crypto/ecdh: invalid public key")
//line /usr/local/go/src/crypto/ecdh/x25519.go:69
		// _ = "end of CoverTab[2903]"
	} else {
//line /usr/local/go/src/crypto/ecdh/x25519.go:70
		_go_fuzz_dep_.CoverTab[2904]++
//line /usr/local/go/src/crypto/ecdh/x25519.go:70
		// _ = "end of CoverTab[2904]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:70
	}
//line /usr/local/go/src/crypto/ecdh/x25519.go:70
	// _ = "end of CoverTab[2901]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:70
	_go_fuzz_dep_.CoverTab[2902]++
							return &PublicKey{
		curve:		c,
		publicKey:	append([]byte{}, key...),
	}, nil
//line /usr/local/go/src/crypto/ecdh/x25519.go:74
	// _ = "end of CoverTab[2902]"
}

func (c *x25519Curve) ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error) {
//line /usr/local/go/src/crypto/ecdh/x25519.go:77
	_go_fuzz_dep_.CoverTab[2905]++
							out := make([]byte, x25519SharedSecretSize)
							x25519ScalarMult(out, local.privateKey, remote.publicKey)
							if isZero(out) {
//line /usr/local/go/src/crypto/ecdh/x25519.go:80
		_go_fuzz_dep_.CoverTab[2907]++
								return nil, errors.New("crypto/ecdh: bad X25519 remote ECDH input: low order point")
//line /usr/local/go/src/crypto/ecdh/x25519.go:81
		// _ = "end of CoverTab[2907]"
	} else {
//line /usr/local/go/src/crypto/ecdh/x25519.go:82
		_go_fuzz_dep_.CoverTab[2908]++
//line /usr/local/go/src/crypto/ecdh/x25519.go:82
		// _ = "end of CoverTab[2908]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:82
	}
//line /usr/local/go/src/crypto/ecdh/x25519.go:82
	// _ = "end of CoverTab[2905]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:82
	_go_fuzz_dep_.CoverTab[2906]++
							return out, nil
//line /usr/local/go/src/crypto/ecdh/x25519.go:83
	// _ = "end of CoverTab[2906]"
}

func x25519ScalarMult(dst, scalar, point []byte) {
//line /usr/local/go/src/crypto/ecdh/x25519.go:86
	_go_fuzz_dep_.CoverTab[2909]++
							var e [32]byte

							copy(e[:], scalar[:])
							e[0] &= 248
							e[31] &= 127
							e[31] |= 64

							var x1, x2, z2, x3, z3, tmp0, tmp1 field.Element
							x1.SetBytes(point[:])
							x2.One()
							x3.Set(&x1)
							z3.One()

							swap := 0
							for pos := 254; pos >= 0; pos-- {
//line /usr/local/go/src/crypto/ecdh/x25519.go:101
		_go_fuzz_dep_.CoverTab[2911]++
								b := e[pos/8] >> uint(pos&7)
								b &= 1
								swap ^= int(b)
								x2.Swap(&x3, swap)
								z2.Swap(&z3, swap)
								swap = int(b)

								tmp0.Subtract(&x3, &z3)
								tmp1.Subtract(&x2, &z2)
								x2.Add(&x2, &z2)
								z2.Add(&x3, &z3)
								z3.Multiply(&tmp0, &x2)
								z2.Multiply(&z2, &tmp1)
								tmp0.Square(&tmp1)
								tmp1.Square(&x2)
								x3.Add(&z3, &z2)
								z2.Subtract(&z3, &z2)
								x2.Multiply(&tmp1, &tmp0)
								tmp1.Subtract(&tmp1, &tmp0)
								z2.Square(&z2)

								z3.Mult32(&tmp1, 121666)
								x3.Square(&x3)
								tmp0.Add(&tmp0, &z3)
								z3.Multiply(&x1, &z2)
								z2.Multiply(&tmp1, &tmp0)
//line /usr/local/go/src/crypto/ecdh/x25519.go:127
		// _ = "end of CoverTab[2911]"
	}
//line /usr/local/go/src/crypto/ecdh/x25519.go:128
	// _ = "end of CoverTab[2909]"
//line /usr/local/go/src/crypto/ecdh/x25519.go:128
	_go_fuzz_dep_.CoverTab[2910]++

							x2.Swap(&x3, swap)
							z2.Swap(&z3, swap)

							z2.Invert(&z2)
							x2.Multiply(&x2, &z2)
							copy(dst[:], x2.Bytes())
//line /usr/local/go/src/crypto/ecdh/x25519.go:135
	// _ = "end of CoverTab[2910]"
}

//line /usr/local/go/src/crypto/ecdh/x25519.go:136
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/ecdh/x25519.go:136
var _ = _go_fuzz_dep_.CoverTab
