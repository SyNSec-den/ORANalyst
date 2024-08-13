// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/ecdh/ecdh.go:5
// Package ecdh implements Elliptic Curve Diffie-Hellman over
//line /usr/local/go/src/crypto/ecdh/ecdh.go:5
// NIST curves and Curve25519.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
package ecdh

//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
import (
//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
)
//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
import (
//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:7
)

import (
	"crypto"
	"crypto/internal/boring"
	"crypto/subtle"
	"errors"
	"io"
	"sync"
)

type Curve interface {
	// GenerateKey generates a new PrivateKey from rand.
	GenerateKey(rand io.Reader) (*PrivateKey, error)

	// NewPrivateKey checks that key is valid and returns a PrivateKey.
	//
	// For NIST curves, this follows SEC 1, Version 2.0, Section 2.3.6, which
	// amounts to decoding the bytes as a fixed length big endian integer and
	// checking that the result is lower than the order of the curve. The zero
	// private key is also rejected, as the encoding of the corresponding public
	// key would be irregular.
	//
	// For X25519, this only checks the scalar length.
	NewPrivateKey(key []byte) (*PrivateKey, error)

	// NewPublicKey checks that key is valid and returns a PublicKey.
	//
	// For NIST curves, this decodes an uncompressed point according to SEC 1,
	// Version 2.0, Section 2.3.4. Compressed encodings and the point at
	// infinity are rejected.
	//
	// For X25519, this only checks the u-coordinate length. Adversarially
	// selected public keys can cause ECDH to return an error.
	NewPublicKey(key []byte) (*PublicKey, error)

	// ecdh performs a ECDH exchange and returns the shared secret. It's exposed
	// as the PrivateKey.ECDH method.
	//
	// The private method also allow us to expand the ECDH interface with more
	// methods in the future without breaking backwards compatibility.
	ecdh(local *PrivateKey, remote *PublicKey) ([]byte, error)

	// privateKeyToPublicKey converts a PrivateKey to a PublicKey. It's exposed
	// as the PrivateKey.PublicKey method.
	//
	// This method always succeeds: for X25519, the zero key can't be
	// constructed due to clamping; for NIST curves, it is rejected by
	// NewPrivateKey.
	privateKeyToPublicKey(*PrivateKey) *PublicKey
}

// PublicKey is an ECDH public key, usually a peer's ECDH share sent over the wire.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:59
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:59
// These keys can be parsed with [crypto/x509.ParsePKIXPublicKey] and encoded
//line /usr/local/go/src/crypto/ecdh/ecdh.go:59
// with [crypto/x509.MarshalPKIXPublicKey]. For NIST curves, they then need to
//line /usr/local/go/src/crypto/ecdh/ecdh.go:59
// be converted with [crypto/ecdsa.PublicKey.ECDH] after parsing.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:64
type PublicKey struct {
	curve		Curve
	publicKey	[]byte
	boring		*boring.PublicKeyECDH
}

// Bytes returns a copy of the encoding of the public key.
func (k *PublicKey) Bytes() []byte {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:71
	_go_fuzz_dep_.CoverTab[2776]++
	// Copy the public key to a fixed size buffer that can get allocated on the
							// caller's stack after inlining.
							var buf [133]byte
							return append(buf[:0], k.publicKey...)
//line /usr/local/go/src/crypto/ecdh/ecdh.go:75
	// _ = "end of CoverTab[2776]"
}

// Equal returns whether x represents the same public key as k.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:78
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:78
// Note that there can be equivalent public keys with different encodings which
//line /usr/local/go/src/crypto/ecdh/ecdh.go:78
// would return false from this check but behave the same way as inputs to ECDH.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:78
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:78
// This check is performed in constant time as long as the key types and their
//line /usr/local/go/src/crypto/ecdh/ecdh.go:78
// curve match.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:85
func (k *PublicKey) Equal(x crypto.PublicKey) bool {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:85
	_go_fuzz_dep_.CoverTab[2777]++
							xx, ok := x.(*PublicKey)
							if !ok {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:87
		_go_fuzz_dep_.CoverTab[2779]++
								return false
//line /usr/local/go/src/crypto/ecdh/ecdh.go:88
		// _ = "end of CoverTab[2779]"
	} else {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:89
		_go_fuzz_dep_.CoverTab[2780]++
//line /usr/local/go/src/crypto/ecdh/ecdh.go:89
		// _ = "end of CoverTab[2780]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:89
	}
//line /usr/local/go/src/crypto/ecdh/ecdh.go:89
	// _ = "end of CoverTab[2777]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:89
	_go_fuzz_dep_.CoverTab[2778]++
							return k.curve == xx.curve && func() bool {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:90
		_go_fuzz_dep_.CoverTab[2781]++
//line /usr/local/go/src/crypto/ecdh/ecdh.go:90
		return subtle.ConstantTimeCompare(k.publicKey, xx.publicKey) == 1
								// _ = "end of CoverTab[2781]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:91
	}()
//line /usr/local/go/src/crypto/ecdh/ecdh.go:91
	// _ = "end of CoverTab[2778]"
}

func (k *PublicKey) Curve() Curve {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:94
	_go_fuzz_dep_.CoverTab[2782]++
							return k.curve
//line /usr/local/go/src/crypto/ecdh/ecdh.go:95
	// _ = "end of CoverTab[2782]"
}

// PrivateKey is an ECDH private key, usually kept secret.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:98
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:98
// These keys can be parsed with [crypto/x509.ParsePKCS8PrivateKey] and encoded
//line /usr/local/go/src/crypto/ecdh/ecdh.go:98
// with [crypto/x509.MarshalPKCS8PrivateKey]. For NIST curves, they then need to
//line /usr/local/go/src/crypto/ecdh/ecdh.go:98
// be converted with [crypto/ecdsa.PrivateKey.ECDH] after parsing.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:103
type PrivateKey struct {
	curve		Curve
	privateKey	[]byte
	boring		*boring.PrivateKeyECDH
	// publicKey is set under publicKeyOnce, to allow loading private keys with
	// NewPrivateKey without having to perform a scalar multiplication.
	publicKey	*PublicKey
	publicKeyOnce	sync.Once
}

// ECDH performs a ECDH exchange and returns the shared secret. The PrivateKey
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
// and PublicKey must use the same curve.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
// For NIST curves, this performs ECDH as specified in SEC 1, Version 2.0,
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
// Section 3.3.1, and returns the x-coordinate encoded according to SEC 1,
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
// Version 2.0, Section 2.3.5. The result is never the point at infinity.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
// For X25519, this performs ECDH as specified in RFC 7748, Section 6.1. If
//line /usr/local/go/src/crypto/ecdh/ecdh.go:113
// the result is the all-zero value, ECDH returns an error.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:122
func (k *PrivateKey) ECDH(remote *PublicKey) ([]byte, error) {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:122
	_go_fuzz_dep_.CoverTab[2783]++
							if k.curve != remote.curve {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:123
		_go_fuzz_dep_.CoverTab[2785]++
								return nil, errors.New("crypto/ecdh: private key and public key curves do not match")
//line /usr/local/go/src/crypto/ecdh/ecdh.go:124
		// _ = "end of CoverTab[2785]"
	} else {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:125
		_go_fuzz_dep_.CoverTab[2786]++
//line /usr/local/go/src/crypto/ecdh/ecdh.go:125
		// _ = "end of CoverTab[2786]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:125
	}
//line /usr/local/go/src/crypto/ecdh/ecdh.go:125
	// _ = "end of CoverTab[2783]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:125
	_go_fuzz_dep_.CoverTab[2784]++
							return k.curve.ecdh(k, remote)
//line /usr/local/go/src/crypto/ecdh/ecdh.go:126
	// _ = "end of CoverTab[2784]"
}

// Bytes returns a copy of the encoding of the private key.
func (k *PrivateKey) Bytes() []byte {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:130
	_go_fuzz_dep_.CoverTab[2787]++
	// Copy the private key to a fixed size buffer that can get allocated on the
							// caller's stack after inlining.
							var buf [66]byte
							return append(buf[:0], k.privateKey...)
//line /usr/local/go/src/crypto/ecdh/ecdh.go:134
	// _ = "end of CoverTab[2787]"
}

// Equal returns whether x represents the same private key as k.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:137
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:137
// Note that there can be equivalent private keys with different encodings which
//line /usr/local/go/src/crypto/ecdh/ecdh.go:137
// would return false from this check but behave the same way as inputs to ECDH.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:137
//
//line /usr/local/go/src/crypto/ecdh/ecdh.go:137
// This check is performed in constant time as long as the key types and their
//line /usr/local/go/src/crypto/ecdh/ecdh.go:137
// curve match.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:144
func (k *PrivateKey) Equal(x crypto.PrivateKey) bool {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:144
	_go_fuzz_dep_.CoverTab[2788]++
							xx, ok := x.(*PrivateKey)
							if !ok {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:146
		_go_fuzz_dep_.CoverTab[2790]++
								return false
//line /usr/local/go/src/crypto/ecdh/ecdh.go:147
		// _ = "end of CoverTab[2790]"
	} else {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:148
		_go_fuzz_dep_.CoverTab[2791]++
//line /usr/local/go/src/crypto/ecdh/ecdh.go:148
		// _ = "end of CoverTab[2791]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:148
	}
//line /usr/local/go/src/crypto/ecdh/ecdh.go:148
	// _ = "end of CoverTab[2788]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:148
	_go_fuzz_dep_.CoverTab[2789]++
							return k.curve == xx.curve && func() bool {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:149
		_go_fuzz_dep_.CoverTab[2792]++
//line /usr/local/go/src/crypto/ecdh/ecdh.go:149
		return subtle.ConstantTimeCompare(k.privateKey, xx.privateKey) == 1
								// _ = "end of CoverTab[2792]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:150
	}()
//line /usr/local/go/src/crypto/ecdh/ecdh.go:150
	// _ = "end of CoverTab[2789]"
}

func (k *PrivateKey) Curve() Curve {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:153
	_go_fuzz_dep_.CoverTab[2793]++
							return k.curve
//line /usr/local/go/src/crypto/ecdh/ecdh.go:154
	// _ = "end of CoverTab[2793]"
}

func (k *PrivateKey) PublicKey() *PublicKey {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:157
	_go_fuzz_dep_.CoverTab[2794]++
							k.publicKeyOnce.Do(func() {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:158
		_go_fuzz_dep_.CoverTab[2796]++
								if k.boring != nil {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:159
			_go_fuzz_dep_.CoverTab[2797]++

//line /usr/local/go/src/crypto/ecdh/ecdh.go:164
			kpub, err := k.boring.PublicKey()
			if err != nil {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:165
				_go_fuzz_dep_.CoverTab[2799]++
										panic("boringcrypto: " + err.Error())
//line /usr/local/go/src/crypto/ecdh/ecdh.go:166
				// _ = "end of CoverTab[2799]"
			} else {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:167
				_go_fuzz_dep_.CoverTab[2800]++
//line /usr/local/go/src/crypto/ecdh/ecdh.go:167
				// _ = "end of CoverTab[2800]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:167
			}
//line /usr/local/go/src/crypto/ecdh/ecdh.go:167
			// _ = "end of CoverTab[2797]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:167
			_go_fuzz_dep_.CoverTab[2798]++
									k.publicKey = &PublicKey{
				curve:		k.curve,
				publicKey:	kpub.Bytes(),
				boring:		kpub,
			}
//line /usr/local/go/src/crypto/ecdh/ecdh.go:172
			// _ = "end of CoverTab[2798]"
		} else {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:173
			_go_fuzz_dep_.CoverTab[2801]++
									k.publicKey = k.curve.privateKeyToPublicKey(k)
//line /usr/local/go/src/crypto/ecdh/ecdh.go:174
			// _ = "end of CoverTab[2801]"
		}
//line /usr/local/go/src/crypto/ecdh/ecdh.go:175
		// _ = "end of CoverTab[2796]"
	})
//line /usr/local/go/src/crypto/ecdh/ecdh.go:176
	// _ = "end of CoverTab[2794]"
//line /usr/local/go/src/crypto/ecdh/ecdh.go:176
	_go_fuzz_dep_.CoverTab[2795]++
							return k.publicKey
//line /usr/local/go/src/crypto/ecdh/ecdh.go:177
	// _ = "end of CoverTab[2795]"
}

// Public implements the implicit interface of all standard library private
//line /usr/local/go/src/crypto/ecdh/ecdh.go:180
// keys. See the docs of crypto.PrivateKey.
//line /usr/local/go/src/crypto/ecdh/ecdh.go:182
func (k *PrivateKey) Public() crypto.PublicKey {
//line /usr/local/go/src/crypto/ecdh/ecdh.go:182
	_go_fuzz_dep_.CoverTab[2802]++
							return k.PublicKey()
//line /usr/local/go/src/crypto/ecdh/ecdh.go:183
	// _ = "end of CoverTab[2802]"
}

//line /usr/local/go/src/crypto/ecdh/ecdh.go:184
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/ecdh/ecdh.go:184
var _ = _go_fuzz_dep_.CoverTab
