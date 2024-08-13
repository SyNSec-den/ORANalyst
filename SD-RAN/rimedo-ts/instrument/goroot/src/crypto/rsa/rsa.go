// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/rsa/rsa.go:5
// Package rsa implements RSA encryption as specified in PKCS #1 and RFC 8017.
//line /usr/local/go/src/crypto/rsa/rsa.go:5
//
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// RSA is a single, fundamental operation that is used in this package to
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// implement either public-key encryption or public-key signatures.
//line /usr/local/go/src/crypto/rsa/rsa.go:5
//
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// The original specification for encryption and signatures with RSA is PKCS #1
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// and the terms "RSA encryption" and "RSA signatures" by default refer to
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// PKCS #1 version 1.5. However, that specification has flaws and new designs
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// should use version 2, usually called by just OAEP and PSS, where
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// possible.
//line /usr/local/go/src/crypto/rsa/rsa.go:5
//
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// Two sets of interfaces are included in this package. When a more abstract
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// interface isn't necessary, there are functions for encrypting/decrypting
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// with v1.5/OAEP and signing/verifying with v1.5/PSS. If one needs to abstract
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// over the public key primitive, the PrivateKey type implements the
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// Decrypter and Signer interfaces from the crypto package.
//line /usr/local/go/src/crypto/rsa/rsa.go:5
//
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// Operations in this package are implemented using constant-time algorithms,
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// except for [GenerateKey], [PrivateKey.Precompute], and [PrivateKey.Validate].
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// Every other operation only leaks the bit size of the involved values, which
//line /usr/local/go/src/crypto/rsa/rsa.go:5
// all depend on the selected key size.
//line /usr/local/go/src/crypto/rsa/rsa.go:26
package rsa

//line /usr/local/go/src/crypto/rsa/rsa.go:26
import (
//line /usr/local/go/src/crypto/rsa/rsa.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rsa/rsa.go:26
)
//line /usr/local/go/src/crypto/rsa/rsa.go:26
import (
//line /usr/local/go/src/crypto/rsa/rsa.go:26
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rsa/rsa.go:26
)

import (
	"crypto"
	"crypto/internal/bigmod"
	"crypto/internal/boring"
	"crypto/internal/boring/bbig"
	"crypto/internal/randutil"
	"crypto/rand"
	"crypto/subtle"
	"encoding/binary"
	"errors"
	"hash"
	"io"
	"math"
	"math/big"
)

var bigOne = big.NewInt(1)

// A PublicKey represents the public part of an RSA key.
type PublicKey struct {
	N	*big.Int	// modulus
	E	int		// public exponent
}

//line /usr/local/go/src/crypto/rsa/rsa.go:55
// Size returns the modulus size in bytes. Raw signatures and ciphertexts
//line /usr/local/go/src/crypto/rsa/rsa.go:55
// for or by this public key will have the same size.
//line /usr/local/go/src/crypto/rsa/rsa.go:57
func (pub *PublicKey) Size() int {
//line /usr/local/go/src/crypto/rsa/rsa.go:57
	_go_fuzz_dep_.CoverTab[9863]++
						return (pub.N.BitLen() + 7) / 8
//line /usr/local/go/src/crypto/rsa/rsa.go:58
	// _ = "end of CoverTab[9863]"
}

// Equal reports whether pub and x have the same value.
func (pub *PublicKey) Equal(x crypto.PublicKey) bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:62
	_go_fuzz_dep_.CoverTab[9864]++
						xx, ok := x.(*PublicKey)
						if !ok {
//line /usr/local/go/src/crypto/rsa/rsa.go:64
		_go_fuzz_dep_.CoverTab[9866]++
							return false
//line /usr/local/go/src/crypto/rsa/rsa.go:65
		// _ = "end of CoverTab[9866]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:66
		_go_fuzz_dep_.CoverTab[9867]++
//line /usr/local/go/src/crypto/rsa/rsa.go:66
		// _ = "end of CoverTab[9867]"
//line /usr/local/go/src/crypto/rsa/rsa.go:66
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:66
	// _ = "end of CoverTab[9864]"
//line /usr/local/go/src/crypto/rsa/rsa.go:66
	_go_fuzz_dep_.CoverTab[9865]++
						return pub.N.Cmp(xx.N) == 0 && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:67
		_go_fuzz_dep_.CoverTab[9868]++
//line /usr/local/go/src/crypto/rsa/rsa.go:67
		return pub.E == xx.E
//line /usr/local/go/src/crypto/rsa/rsa.go:67
		// _ = "end of CoverTab[9868]"
//line /usr/local/go/src/crypto/rsa/rsa.go:67
	}()
//line /usr/local/go/src/crypto/rsa/rsa.go:67
	// _ = "end of CoverTab[9865]"
}

// OAEPOptions is an interface for passing options to OAEP decryption using the
//line /usr/local/go/src/crypto/rsa/rsa.go:70
// crypto.Decrypter interface.
//line /usr/local/go/src/crypto/rsa/rsa.go:72
type OAEPOptions struct {
	// Hash is the hash function that will be used when generating the mask.
	Hash	crypto.Hash

	// MGFHash is the hash function used for MGF1.
	// If zero, Hash is used instead.
	MGFHash	crypto.Hash

	// Label is an arbitrary byte string that must be equal to the value
	// used when encrypting.
	Label	[]byte
}

var (
	errPublicModulus	= errors.New("crypto/rsa: missing public modulus")
	errPublicExponentSmall	= errors.New("crypto/rsa: public exponent too small")
	errPublicExponentLarge	= errors.New("crypto/rsa: public exponent too large")
)

// checkPub sanity checks the public key before we use it.
//line /usr/local/go/src/crypto/rsa/rsa.go:91
// We require pub.E to fit into a 32-bit integer so that we
//line /usr/local/go/src/crypto/rsa/rsa.go:91
// do not have different behavior depending on whether
//line /usr/local/go/src/crypto/rsa/rsa.go:91
// int is 32 or 64 bits. See also
//line /usr/local/go/src/crypto/rsa/rsa.go:91
// https://www.imperialviolet.org/2012/03/16/rsae.html.
//line /usr/local/go/src/crypto/rsa/rsa.go:96
func checkPub(pub *PublicKey) error {
//line /usr/local/go/src/crypto/rsa/rsa.go:96
	_go_fuzz_dep_.CoverTab[9869]++
						if pub.N == nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:97
		_go_fuzz_dep_.CoverTab[9873]++
							return errPublicModulus
//line /usr/local/go/src/crypto/rsa/rsa.go:98
		// _ = "end of CoverTab[9873]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:99
		_go_fuzz_dep_.CoverTab[9874]++
//line /usr/local/go/src/crypto/rsa/rsa.go:99
		// _ = "end of CoverTab[9874]"
//line /usr/local/go/src/crypto/rsa/rsa.go:99
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:99
	// _ = "end of CoverTab[9869]"
//line /usr/local/go/src/crypto/rsa/rsa.go:99
	_go_fuzz_dep_.CoverTab[9870]++
						if pub.E < 2 {
//line /usr/local/go/src/crypto/rsa/rsa.go:100
		_go_fuzz_dep_.CoverTab[9875]++
							return errPublicExponentSmall
//line /usr/local/go/src/crypto/rsa/rsa.go:101
		// _ = "end of CoverTab[9875]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:102
		_go_fuzz_dep_.CoverTab[9876]++
//line /usr/local/go/src/crypto/rsa/rsa.go:102
		// _ = "end of CoverTab[9876]"
//line /usr/local/go/src/crypto/rsa/rsa.go:102
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:102
	// _ = "end of CoverTab[9870]"
//line /usr/local/go/src/crypto/rsa/rsa.go:102
	_go_fuzz_dep_.CoverTab[9871]++
						if pub.E > 1<<31-1 {
//line /usr/local/go/src/crypto/rsa/rsa.go:103
		_go_fuzz_dep_.CoverTab[9877]++
							return errPublicExponentLarge
//line /usr/local/go/src/crypto/rsa/rsa.go:104
		// _ = "end of CoverTab[9877]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:105
		_go_fuzz_dep_.CoverTab[9878]++
//line /usr/local/go/src/crypto/rsa/rsa.go:105
		// _ = "end of CoverTab[9878]"
//line /usr/local/go/src/crypto/rsa/rsa.go:105
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:105
	// _ = "end of CoverTab[9871]"
//line /usr/local/go/src/crypto/rsa/rsa.go:105
	_go_fuzz_dep_.CoverTab[9872]++
						return nil
//line /usr/local/go/src/crypto/rsa/rsa.go:106
	// _ = "end of CoverTab[9872]"
}

// A PrivateKey represents an RSA key
type PrivateKey struct {
	PublicKey			// public part.
	D		*big.Int	// private exponent
	Primes		[]*big.Int	// prime factors of N, has >= 2 elements.

	// Precomputed contains precomputed values that speed up RSA operations,
	// if available. It must be generated by calling PrivateKey.Precompute and
	// must not be modified.
	Precomputed	PrecomputedValues
}

// Public returns the public key corresponding to priv.
func (priv *PrivateKey) Public() crypto.PublicKey {
//line /usr/local/go/src/crypto/rsa/rsa.go:122
	_go_fuzz_dep_.CoverTab[9879]++
						return &priv.PublicKey
//line /usr/local/go/src/crypto/rsa/rsa.go:123
	// _ = "end of CoverTab[9879]"
}

// Equal reports whether priv and x have equivalent values. It ignores
//line /usr/local/go/src/crypto/rsa/rsa.go:126
// Precomputed values.
//line /usr/local/go/src/crypto/rsa/rsa.go:128
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:128
	_go_fuzz_dep_.CoverTab[9880]++
						xx, ok := x.(*PrivateKey)
						if !ok {
//line /usr/local/go/src/crypto/rsa/rsa.go:130
		_go_fuzz_dep_.CoverTab[9885]++
							return false
//line /usr/local/go/src/crypto/rsa/rsa.go:131
		// _ = "end of CoverTab[9885]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:132
		_go_fuzz_dep_.CoverTab[9886]++
//line /usr/local/go/src/crypto/rsa/rsa.go:132
		// _ = "end of CoverTab[9886]"
//line /usr/local/go/src/crypto/rsa/rsa.go:132
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:132
	// _ = "end of CoverTab[9880]"
//line /usr/local/go/src/crypto/rsa/rsa.go:132
	_go_fuzz_dep_.CoverTab[9881]++
						if !priv.PublicKey.Equal(&xx.PublicKey) || func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:133
		_go_fuzz_dep_.CoverTab[9887]++
//line /usr/local/go/src/crypto/rsa/rsa.go:133
		return priv.D.Cmp(xx.D) != 0
//line /usr/local/go/src/crypto/rsa/rsa.go:133
		// _ = "end of CoverTab[9887]"
//line /usr/local/go/src/crypto/rsa/rsa.go:133
	}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:133
		_go_fuzz_dep_.CoverTab[9888]++
							return false
//line /usr/local/go/src/crypto/rsa/rsa.go:134
		// _ = "end of CoverTab[9888]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:135
		_go_fuzz_dep_.CoverTab[9889]++
//line /usr/local/go/src/crypto/rsa/rsa.go:135
		// _ = "end of CoverTab[9889]"
//line /usr/local/go/src/crypto/rsa/rsa.go:135
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:135
	// _ = "end of CoverTab[9881]"
//line /usr/local/go/src/crypto/rsa/rsa.go:135
	_go_fuzz_dep_.CoverTab[9882]++
						if len(priv.Primes) != len(xx.Primes) {
//line /usr/local/go/src/crypto/rsa/rsa.go:136
		_go_fuzz_dep_.CoverTab[9890]++
							return false
//line /usr/local/go/src/crypto/rsa/rsa.go:137
		// _ = "end of CoverTab[9890]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:138
		_go_fuzz_dep_.CoverTab[9891]++
//line /usr/local/go/src/crypto/rsa/rsa.go:138
		// _ = "end of CoverTab[9891]"
//line /usr/local/go/src/crypto/rsa/rsa.go:138
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:138
	// _ = "end of CoverTab[9882]"
//line /usr/local/go/src/crypto/rsa/rsa.go:138
	_go_fuzz_dep_.CoverTab[9883]++
						for i := range priv.Primes {
//line /usr/local/go/src/crypto/rsa/rsa.go:139
		_go_fuzz_dep_.CoverTab[9892]++
							if priv.Primes[i].Cmp(xx.Primes[i]) != 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:140
			_go_fuzz_dep_.CoverTab[9893]++
								return false
//line /usr/local/go/src/crypto/rsa/rsa.go:141
			// _ = "end of CoverTab[9893]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:142
			_go_fuzz_dep_.CoverTab[9894]++
//line /usr/local/go/src/crypto/rsa/rsa.go:142
			// _ = "end of CoverTab[9894]"
//line /usr/local/go/src/crypto/rsa/rsa.go:142
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:142
		// _ = "end of CoverTab[9892]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:143
	// _ = "end of CoverTab[9883]"
//line /usr/local/go/src/crypto/rsa/rsa.go:143
	_go_fuzz_dep_.CoverTab[9884]++
						return true
//line /usr/local/go/src/crypto/rsa/rsa.go:144
	// _ = "end of CoverTab[9884]"
}

// Sign signs digest with priv, reading randomness from rand. If opts is a
//line /usr/local/go/src/crypto/rsa/rsa.go:147
// *PSSOptions then the PSS algorithm will be used, otherwise PKCS #1 v1.5 will
//line /usr/local/go/src/crypto/rsa/rsa.go:147
// be used. digest must be the result of hashing the input message using
//line /usr/local/go/src/crypto/rsa/rsa.go:147
// opts.HashFunc().
//line /usr/local/go/src/crypto/rsa/rsa.go:147
//
//line /usr/local/go/src/crypto/rsa/rsa.go:147
// This method implements crypto.Signer, which is an interface to support keys
//line /usr/local/go/src/crypto/rsa/rsa.go:147
// where the private part is kept in, for example, a hardware module. Common
//line /usr/local/go/src/crypto/rsa/rsa.go:147
// uses should use the Sign* functions in this package directly.
//line /usr/local/go/src/crypto/rsa/rsa.go:155
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:155
	_go_fuzz_dep_.CoverTab[9895]++
						if pssOpts, ok := opts.(*PSSOptions); ok {
//line /usr/local/go/src/crypto/rsa/rsa.go:156
		_go_fuzz_dep_.CoverTab[9897]++
							return SignPSS(rand, priv, pssOpts.Hash, digest, pssOpts)
//line /usr/local/go/src/crypto/rsa/rsa.go:157
		// _ = "end of CoverTab[9897]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:158
		_go_fuzz_dep_.CoverTab[9898]++
//line /usr/local/go/src/crypto/rsa/rsa.go:158
		// _ = "end of CoverTab[9898]"
//line /usr/local/go/src/crypto/rsa/rsa.go:158
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:158
	// _ = "end of CoverTab[9895]"
//line /usr/local/go/src/crypto/rsa/rsa.go:158
	_go_fuzz_dep_.CoverTab[9896]++

						return SignPKCS1v15(rand, priv, opts.HashFunc(), digest)
//line /usr/local/go/src/crypto/rsa/rsa.go:160
	// _ = "end of CoverTab[9896]"
}

// Decrypt decrypts ciphertext with priv. If opts is nil or of type
//line /usr/local/go/src/crypto/rsa/rsa.go:163
// *PKCS1v15DecryptOptions then PKCS #1 v1.5 decryption is performed. Otherwise
//line /usr/local/go/src/crypto/rsa/rsa.go:163
// opts must have type *OAEPOptions and OAEP decryption is done.
//line /usr/local/go/src/crypto/rsa/rsa.go:166
func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:166
	_go_fuzz_dep_.CoverTab[9899]++
						if opts == nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:167
		_go_fuzz_dep_.CoverTab[9901]++
							return DecryptPKCS1v15(rand, priv, ciphertext)
//line /usr/local/go/src/crypto/rsa/rsa.go:168
		// _ = "end of CoverTab[9901]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:169
		_go_fuzz_dep_.CoverTab[9902]++
//line /usr/local/go/src/crypto/rsa/rsa.go:169
		// _ = "end of CoverTab[9902]"
//line /usr/local/go/src/crypto/rsa/rsa.go:169
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:169
	// _ = "end of CoverTab[9899]"
//line /usr/local/go/src/crypto/rsa/rsa.go:169
	_go_fuzz_dep_.CoverTab[9900]++

						switch opts := opts.(type) {
	case *OAEPOptions:
//line /usr/local/go/src/crypto/rsa/rsa.go:172
		_go_fuzz_dep_.CoverTab[9903]++
							if opts.MGFHash == 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:173
			_go_fuzz_dep_.CoverTab[9906]++
								return decryptOAEP(opts.Hash.New(), opts.Hash.New(), rand, priv, ciphertext, opts.Label)
//line /usr/local/go/src/crypto/rsa/rsa.go:174
			// _ = "end of CoverTab[9906]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:175
			_go_fuzz_dep_.CoverTab[9907]++
								return decryptOAEP(opts.Hash.New(), opts.MGFHash.New(), rand, priv, ciphertext, opts.Label)
//line /usr/local/go/src/crypto/rsa/rsa.go:176
			// _ = "end of CoverTab[9907]"
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:177
		// _ = "end of CoverTab[9903]"

	case *PKCS1v15DecryptOptions:
//line /usr/local/go/src/crypto/rsa/rsa.go:179
		_go_fuzz_dep_.CoverTab[9904]++
							if l := opts.SessionKeyLen; l > 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:180
			_go_fuzz_dep_.CoverTab[9908]++
								plaintext = make([]byte, l)
								if _, err := io.ReadFull(rand, plaintext); err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:182
				_go_fuzz_dep_.CoverTab[9911]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:183
				// _ = "end of CoverTab[9911]"
			} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:184
				_go_fuzz_dep_.CoverTab[9912]++
//line /usr/local/go/src/crypto/rsa/rsa.go:184
				// _ = "end of CoverTab[9912]"
//line /usr/local/go/src/crypto/rsa/rsa.go:184
			}
//line /usr/local/go/src/crypto/rsa/rsa.go:184
			// _ = "end of CoverTab[9908]"
//line /usr/local/go/src/crypto/rsa/rsa.go:184
			_go_fuzz_dep_.CoverTab[9909]++
								if err := DecryptPKCS1v15SessionKey(rand, priv, ciphertext, plaintext); err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:185
				_go_fuzz_dep_.CoverTab[9913]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:186
				// _ = "end of CoverTab[9913]"
			} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:187
				_go_fuzz_dep_.CoverTab[9914]++
//line /usr/local/go/src/crypto/rsa/rsa.go:187
				// _ = "end of CoverTab[9914]"
//line /usr/local/go/src/crypto/rsa/rsa.go:187
			}
//line /usr/local/go/src/crypto/rsa/rsa.go:187
			// _ = "end of CoverTab[9909]"
//line /usr/local/go/src/crypto/rsa/rsa.go:187
			_go_fuzz_dep_.CoverTab[9910]++
								return plaintext, nil
//line /usr/local/go/src/crypto/rsa/rsa.go:188
			// _ = "end of CoverTab[9910]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:189
			_go_fuzz_dep_.CoverTab[9915]++
								return DecryptPKCS1v15(rand, priv, ciphertext)
//line /usr/local/go/src/crypto/rsa/rsa.go:190
			// _ = "end of CoverTab[9915]"
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:191
		// _ = "end of CoverTab[9904]"

	default:
//line /usr/local/go/src/crypto/rsa/rsa.go:193
		_go_fuzz_dep_.CoverTab[9905]++
							return nil, errors.New("crypto/rsa: invalid options for Decrypt")
//line /usr/local/go/src/crypto/rsa/rsa.go:194
		// _ = "end of CoverTab[9905]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:195
	// _ = "end of CoverTab[9900]"
}

type PrecomputedValues struct {
	Dp, Dq	*big.Int	// D mod (P-1) (or mod Q-1)
	Qinv	*big.Int	// Q^-1 mod P

	// CRTValues is used for the 3rd and subsequent primes. Due to a
	// historical accident, the CRT for the first two primes is handled
	// differently in PKCS #1 and interoperability is sufficiently
	// important that we mirror this.
	//
	// Note: these values are still filled in by Precompute for
	// backwards compatibility but are not used. Multi-prime RSA is very rare,
	// and is implemented by this package without CRT optimizations to limit
	// complexity.
	CRTValues	[]CRTValue

	n, p, q	*bigmod.Modulus	// moduli for CRT with Montgomery precomputed constants
}

// CRTValue contains the precomputed Chinese remainder theorem values.
type CRTValue struct {
	Exp	*big.Int	// D mod (prime-1).
	Coeff	*big.Int	// R·Coeff ≡ 1 mod Prime.
	R	*big.Int	// product of primes prior to this (inc p and q).
}

// Validate performs basic sanity checks on the key.
//line /usr/local/go/src/crypto/rsa/rsa.go:223
// It returns nil if the key is valid, or else an error describing a problem.
//line /usr/local/go/src/crypto/rsa/rsa.go:225
func (priv *PrivateKey) Validate() error {
//line /usr/local/go/src/crypto/rsa/rsa.go:225
	_go_fuzz_dep_.CoverTab[9916]++
						if err := checkPub(&priv.PublicKey); err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:226
		_go_fuzz_dep_.CoverTab[9921]++
							return err
//line /usr/local/go/src/crypto/rsa/rsa.go:227
		// _ = "end of CoverTab[9921]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:228
		_go_fuzz_dep_.CoverTab[9922]++
//line /usr/local/go/src/crypto/rsa/rsa.go:228
		// _ = "end of CoverTab[9922]"
//line /usr/local/go/src/crypto/rsa/rsa.go:228
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:228
	// _ = "end of CoverTab[9916]"
//line /usr/local/go/src/crypto/rsa/rsa.go:228
	_go_fuzz_dep_.CoverTab[9917]++

//line /usr/local/go/src/crypto/rsa/rsa.go:231
	modulus := new(big.Int).Set(bigOne)
	for _, prime := range priv.Primes {
//line /usr/local/go/src/crypto/rsa/rsa.go:232
		_go_fuzz_dep_.CoverTab[9923]++

							if prime.Cmp(bigOne) <= 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:234
			_go_fuzz_dep_.CoverTab[9925]++
								return errors.New("crypto/rsa: invalid prime value")
//line /usr/local/go/src/crypto/rsa/rsa.go:235
			// _ = "end of CoverTab[9925]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:236
			_go_fuzz_dep_.CoverTab[9926]++
//line /usr/local/go/src/crypto/rsa/rsa.go:236
			// _ = "end of CoverTab[9926]"
//line /usr/local/go/src/crypto/rsa/rsa.go:236
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:236
		// _ = "end of CoverTab[9923]"
//line /usr/local/go/src/crypto/rsa/rsa.go:236
		_go_fuzz_dep_.CoverTab[9924]++
							modulus.Mul(modulus, prime)
//line /usr/local/go/src/crypto/rsa/rsa.go:237
		// _ = "end of CoverTab[9924]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:238
	// _ = "end of CoverTab[9917]"
//line /usr/local/go/src/crypto/rsa/rsa.go:238
	_go_fuzz_dep_.CoverTab[9918]++
						if modulus.Cmp(priv.N) != 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:239
		_go_fuzz_dep_.CoverTab[9927]++
							return errors.New("crypto/rsa: invalid modulus")
//line /usr/local/go/src/crypto/rsa/rsa.go:240
		// _ = "end of CoverTab[9927]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:241
		_go_fuzz_dep_.CoverTab[9928]++
//line /usr/local/go/src/crypto/rsa/rsa.go:241
		// _ = "end of CoverTab[9928]"
//line /usr/local/go/src/crypto/rsa/rsa.go:241
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:241
	// _ = "end of CoverTab[9918]"
//line /usr/local/go/src/crypto/rsa/rsa.go:241
	_go_fuzz_dep_.CoverTab[9919]++

//line /usr/local/go/src/crypto/rsa/rsa.go:248
	congruence := new(big.Int)
	de := new(big.Int).SetInt64(int64(priv.E))
	de.Mul(de, priv.D)
	for _, prime := range priv.Primes {
//line /usr/local/go/src/crypto/rsa/rsa.go:251
		_go_fuzz_dep_.CoverTab[9929]++
							pminus1 := new(big.Int).Sub(prime, bigOne)
							congruence.Mod(de, pminus1)
							if congruence.Cmp(bigOne) != 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:254
			_go_fuzz_dep_.CoverTab[9930]++
								return errors.New("crypto/rsa: invalid exponents")
//line /usr/local/go/src/crypto/rsa/rsa.go:255
			// _ = "end of CoverTab[9930]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:256
			_go_fuzz_dep_.CoverTab[9931]++
//line /usr/local/go/src/crypto/rsa/rsa.go:256
			// _ = "end of CoverTab[9931]"
//line /usr/local/go/src/crypto/rsa/rsa.go:256
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:256
		// _ = "end of CoverTab[9929]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:257
	// _ = "end of CoverTab[9919]"
//line /usr/local/go/src/crypto/rsa/rsa.go:257
	_go_fuzz_dep_.CoverTab[9920]++
						return nil
//line /usr/local/go/src/crypto/rsa/rsa.go:258
	// _ = "end of CoverTab[9920]"
}

// GenerateKey generates an RSA keypair of the given bit size using the
//line /usr/local/go/src/crypto/rsa/rsa.go:261
// random source random (for example, crypto/rand.Reader).
//line /usr/local/go/src/crypto/rsa/rsa.go:263
func GenerateKey(random io.Reader, bits int) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:263
	_go_fuzz_dep_.CoverTab[9932]++
						return GenerateMultiPrimeKey(random, 2, bits)
//line /usr/local/go/src/crypto/rsa/rsa.go:264
	// _ = "end of CoverTab[9932]"
}

// GenerateMultiPrimeKey generates a multi-prime RSA keypair of the given bit
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// size and the given random source.
//line /usr/local/go/src/crypto/rsa/rsa.go:267
//
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// Table 1 in "[On the Security of Multi-prime RSA]" suggests maximum numbers of
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// primes for a given bit size.
//line /usr/local/go/src/crypto/rsa/rsa.go:267
//
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// Although the public keys are compatible (actually, indistinguishable) from
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// the 2-prime case, the private keys are not. Thus it may not be possible to
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// export multi-prime private keys in certain formats or to subsequently import
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// them into other code.
//line /usr/local/go/src/crypto/rsa/rsa.go:267
//
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// This package does not implement CRT optimizations for multi-prime RSA, so the
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// keys with more than two primes will have worse performance.
//line /usr/local/go/src/crypto/rsa/rsa.go:267
//
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// Note: The use of this function with a number of primes different from
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// two is not recommended for the above security, compatibility, and performance
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// reasons. Use GenerateKey instead.
//line /usr/local/go/src/crypto/rsa/rsa.go:267
//
//line /usr/local/go/src/crypto/rsa/rsa.go:267
// [On the Security of Multi-prime RSA]: http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf
//line /usr/local/go/src/crypto/rsa/rsa.go:286
func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:286
	_go_fuzz_dep_.CoverTab[9933]++
						randutil.MaybeReadByte(random)

						if boring.Enabled && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		_go_fuzz_dep_.CoverTab[9938]++
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		return random == boring.RandReader
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		// _ = "end of CoverTab[9938]"
//line /usr/local/go/src/crypto/rsa/rsa.go:289
	}() && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		_go_fuzz_dep_.CoverTab[9939]++
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		return nprimes == 2
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		// _ = "end of CoverTab[9939]"
//line /usr/local/go/src/crypto/rsa/rsa.go:289
	}() && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		_go_fuzz_dep_.CoverTab[9940]++
//line /usr/local/go/src/crypto/rsa/rsa.go:289
		return (bits == 2048 || func() bool {
								_go_fuzz_dep_.CoverTab[9941]++
//line /usr/local/go/src/crypto/rsa/rsa.go:290
			return bits == 3072
//line /usr/local/go/src/crypto/rsa/rsa.go:290
			// _ = "end of CoverTab[9941]"
//line /usr/local/go/src/crypto/rsa/rsa.go:290
		}() || func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:290
			_go_fuzz_dep_.CoverTab[9942]++
//line /usr/local/go/src/crypto/rsa/rsa.go:290
			return bits == 4096
//line /usr/local/go/src/crypto/rsa/rsa.go:290
			// _ = "end of CoverTab[9942]"
//line /usr/local/go/src/crypto/rsa/rsa.go:290
		}())
//line /usr/local/go/src/crypto/rsa/rsa.go:290
		// _ = "end of CoverTab[9940]"
//line /usr/local/go/src/crypto/rsa/rsa.go:290
	}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:290
		_go_fuzz_dep_.CoverTab[9943]++
							bN, bE, bD, bP, bQ, bDp, bDq, bQinv, err := boring.GenerateKeyRSA(bits)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:292
			_go_fuzz_dep_.CoverTab[9946]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:293
			// _ = "end of CoverTab[9946]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:294
			_go_fuzz_dep_.CoverTab[9947]++
//line /usr/local/go/src/crypto/rsa/rsa.go:294
			// _ = "end of CoverTab[9947]"
//line /usr/local/go/src/crypto/rsa/rsa.go:294
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:294
		// _ = "end of CoverTab[9943]"
//line /usr/local/go/src/crypto/rsa/rsa.go:294
		_go_fuzz_dep_.CoverTab[9944]++
							N := bbig.Dec(bN)
							E := bbig.Dec(bE)
							D := bbig.Dec(bD)
							P := bbig.Dec(bP)
							Q := bbig.Dec(bQ)
							Dp := bbig.Dec(bDp)
							Dq := bbig.Dec(bDq)
							Qinv := bbig.Dec(bQinv)
							e64 := E.Int64()
							if !E.IsInt64() || func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:304
			_go_fuzz_dep_.CoverTab[9948]++
//line /usr/local/go/src/crypto/rsa/rsa.go:304
			return int64(int(e64)) != e64
//line /usr/local/go/src/crypto/rsa/rsa.go:304
			// _ = "end of CoverTab[9948]"
//line /usr/local/go/src/crypto/rsa/rsa.go:304
		}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:304
			_go_fuzz_dep_.CoverTab[9949]++
								return nil, errors.New("crypto/rsa: generated key exponent too large")
//line /usr/local/go/src/crypto/rsa/rsa.go:305
			// _ = "end of CoverTab[9949]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:306
			_go_fuzz_dep_.CoverTab[9950]++
//line /usr/local/go/src/crypto/rsa/rsa.go:306
			// _ = "end of CoverTab[9950]"
//line /usr/local/go/src/crypto/rsa/rsa.go:306
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:306
		// _ = "end of CoverTab[9944]"
//line /usr/local/go/src/crypto/rsa/rsa.go:306
		_go_fuzz_dep_.CoverTab[9945]++
							key := &PrivateKey{
			PublicKey: PublicKey{
				N:	N,
				E:	int(e64),
			},
			D:	D,
			Primes:	[]*big.Int{P, Q},
			Precomputed: PrecomputedValues{
				Dp:		Dp,
				Dq:		Dq,
				Qinv:		Qinv,
				CRTValues:	make([]CRTValue, 0),
				n:		bigmod.NewModulusFromBig(N),
				p:		bigmod.NewModulusFromBig(P),
				q:		bigmod.NewModulusFromBig(Q),
			},
		}
							return key, nil
//line /usr/local/go/src/crypto/rsa/rsa.go:324
		// _ = "end of CoverTab[9945]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:325
		_go_fuzz_dep_.CoverTab[9951]++
//line /usr/local/go/src/crypto/rsa/rsa.go:325
		// _ = "end of CoverTab[9951]"
//line /usr/local/go/src/crypto/rsa/rsa.go:325
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:325
	// _ = "end of CoverTab[9933]"
//line /usr/local/go/src/crypto/rsa/rsa.go:325
	_go_fuzz_dep_.CoverTab[9934]++

						priv := new(PrivateKey)
						priv.E = 65537

						if nprimes < 2 {
//line /usr/local/go/src/crypto/rsa/rsa.go:330
		_go_fuzz_dep_.CoverTab[9952]++
							return nil, errors.New("crypto/rsa: GenerateMultiPrimeKey: nprimes must be >= 2")
//line /usr/local/go/src/crypto/rsa/rsa.go:331
		// _ = "end of CoverTab[9952]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:332
		_go_fuzz_dep_.CoverTab[9953]++
//line /usr/local/go/src/crypto/rsa/rsa.go:332
		// _ = "end of CoverTab[9953]"
//line /usr/local/go/src/crypto/rsa/rsa.go:332
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:332
	// _ = "end of CoverTab[9934]"
//line /usr/local/go/src/crypto/rsa/rsa.go:332
	_go_fuzz_dep_.CoverTab[9935]++

						if bits < 64 {
//line /usr/local/go/src/crypto/rsa/rsa.go:334
		_go_fuzz_dep_.CoverTab[9954]++
							primeLimit := float64(uint64(1) << uint(bits/nprimes))

							pi := primeLimit / (math.Log(primeLimit) - 1)

//line /usr/local/go/src/crypto/rsa/rsa.go:340
		pi /= 4

//line /usr/local/go/src/crypto/rsa/rsa.go:343
		pi /= 2
		if pi <= float64(nprimes) {
//line /usr/local/go/src/crypto/rsa/rsa.go:344
			_go_fuzz_dep_.CoverTab[9955]++
								return nil, errors.New("crypto/rsa: too few primes of given length to generate an RSA key")
//line /usr/local/go/src/crypto/rsa/rsa.go:345
			// _ = "end of CoverTab[9955]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:346
			_go_fuzz_dep_.CoverTab[9956]++
//line /usr/local/go/src/crypto/rsa/rsa.go:346
			// _ = "end of CoverTab[9956]"
//line /usr/local/go/src/crypto/rsa/rsa.go:346
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:346
		// _ = "end of CoverTab[9954]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:347
		_go_fuzz_dep_.CoverTab[9957]++
//line /usr/local/go/src/crypto/rsa/rsa.go:347
		// _ = "end of CoverTab[9957]"
//line /usr/local/go/src/crypto/rsa/rsa.go:347
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:347
	// _ = "end of CoverTab[9935]"
//line /usr/local/go/src/crypto/rsa/rsa.go:347
	_go_fuzz_dep_.CoverTab[9936]++

						primes := make([]*big.Int, nprimes)

NextSetOfPrimes:
	for {
//line /usr/local/go/src/crypto/rsa/rsa.go:352
		_go_fuzz_dep_.CoverTab[9958]++
							todo := bits

//line /usr/local/go/src/crypto/rsa/rsa.go:365
		if nprimes >= 7 {
//line /usr/local/go/src/crypto/rsa/rsa.go:365
			_go_fuzz_dep_.CoverTab[9964]++
								todo += (nprimes - 2) / 5
//line /usr/local/go/src/crypto/rsa/rsa.go:366
			// _ = "end of CoverTab[9964]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:367
			_go_fuzz_dep_.CoverTab[9965]++
//line /usr/local/go/src/crypto/rsa/rsa.go:367
			// _ = "end of CoverTab[9965]"
//line /usr/local/go/src/crypto/rsa/rsa.go:367
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:367
		// _ = "end of CoverTab[9958]"
//line /usr/local/go/src/crypto/rsa/rsa.go:367
		_go_fuzz_dep_.CoverTab[9959]++
							for i := 0; i < nprimes; i++ {
//line /usr/local/go/src/crypto/rsa/rsa.go:368
			_go_fuzz_dep_.CoverTab[9966]++
								var err error
								primes[i], err = rand.Prime(random, todo/(nprimes-i))
								if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:371
				_go_fuzz_dep_.CoverTab[9968]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:372
				// _ = "end of CoverTab[9968]"
			} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:373
				_go_fuzz_dep_.CoverTab[9969]++
//line /usr/local/go/src/crypto/rsa/rsa.go:373
				// _ = "end of CoverTab[9969]"
//line /usr/local/go/src/crypto/rsa/rsa.go:373
			}
//line /usr/local/go/src/crypto/rsa/rsa.go:373
			// _ = "end of CoverTab[9966]"
//line /usr/local/go/src/crypto/rsa/rsa.go:373
			_go_fuzz_dep_.CoverTab[9967]++
								todo -= primes[i].BitLen()
//line /usr/local/go/src/crypto/rsa/rsa.go:374
			// _ = "end of CoverTab[9967]"
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:375
		// _ = "end of CoverTab[9959]"
//line /usr/local/go/src/crypto/rsa/rsa.go:375
		_go_fuzz_dep_.CoverTab[9960]++

//line /usr/local/go/src/crypto/rsa/rsa.go:378
		for i, prime := range primes {
//line /usr/local/go/src/crypto/rsa/rsa.go:378
			_go_fuzz_dep_.CoverTab[9970]++
								for j := 0; j < i; j++ {
//line /usr/local/go/src/crypto/rsa/rsa.go:379
				_go_fuzz_dep_.CoverTab[9971]++
									if prime.Cmp(primes[j]) == 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:380
					_go_fuzz_dep_.CoverTab[9972]++
										continue NextSetOfPrimes
//line /usr/local/go/src/crypto/rsa/rsa.go:381
					// _ = "end of CoverTab[9972]"
				} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:382
					_go_fuzz_dep_.CoverTab[9973]++
//line /usr/local/go/src/crypto/rsa/rsa.go:382
					// _ = "end of CoverTab[9973]"
//line /usr/local/go/src/crypto/rsa/rsa.go:382
				}
//line /usr/local/go/src/crypto/rsa/rsa.go:382
				// _ = "end of CoverTab[9971]"
			}
//line /usr/local/go/src/crypto/rsa/rsa.go:383
			// _ = "end of CoverTab[9970]"
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:384
		// _ = "end of CoverTab[9960]"
//line /usr/local/go/src/crypto/rsa/rsa.go:384
		_go_fuzz_dep_.CoverTab[9961]++

							n := new(big.Int).Set(bigOne)
							totient := new(big.Int).Set(bigOne)
							pminus1 := new(big.Int)
							for _, prime := range primes {
//line /usr/local/go/src/crypto/rsa/rsa.go:389
			_go_fuzz_dep_.CoverTab[9974]++
								n.Mul(n, prime)
								pminus1.Sub(prime, bigOne)
								totient.Mul(totient, pminus1)
//line /usr/local/go/src/crypto/rsa/rsa.go:392
			// _ = "end of CoverTab[9974]"
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:393
		// _ = "end of CoverTab[9961]"
//line /usr/local/go/src/crypto/rsa/rsa.go:393
		_go_fuzz_dep_.CoverTab[9962]++
							if n.BitLen() != bits {
//line /usr/local/go/src/crypto/rsa/rsa.go:394
			_go_fuzz_dep_.CoverTab[9975]++

//line /usr/local/go/src/crypto/rsa/rsa.go:398
			continue NextSetOfPrimes
//line /usr/local/go/src/crypto/rsa/rsa.go:398
			// _ = "end of CoverTab[9975]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:399
			_go_fuzz_dep_.CoverTab[9976]++
//line /usr/local/go/src/crypto/rsa/rsa.go:399
			// _ = "end of CoverTab[9976]"
//line /usr/local/go/src/crypto/rsa/rsa.go:399
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:399
		// _ = "end of CoverTab[9962]"
//line /usr/local/go/src/crypto/rsa/rsa.go:399
		_go_fuzz_dep_.CoverTab[9963]++

							priv.D = new(big.Int)
							e := big.NewInt(int64(priv.E))
							ok := priv.D.ModInverse(e, totient)

							if ok != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:405
			_go_fuzz_dep_.CoverTab[9977]++
								priv.Primes = primes
								priv.N = n
								break
//line /usr/local/go/src/crypto/rsa/rsa.go:408
			// _ = "end of CoverTab[9977]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:409
			_go_fuzz_dep_.CoverTab[9978]++
//line /usr/local/go/src/crypto/rsa/rsa.go:409
			// _ = "end of CoverTab[9978]"
//line /usr/local/go/src/crypto/rsa/rsa.go:409
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:409
		// _ = "end of CoverTab[9963]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:410
	// _ = "end of CoverTab[9936]"
//line /usr/local/go/src/crypto/rsa/rsa.go:410
	_go_fuzz_dep_.CoverTab[9937]++

						priv.Precompute()
						return priv, nil
//line /usr/local/go/src/crypto/rsa/rsa.go:413
	// _ = "end of CoverTab[9937]"
}

// incCounter increments a four byte, big-endian counter.
func incCounter(c *[4]byte) {
//line /usr/local/go/src/crypto/rsa/rsa.go:417
	_go_fuzz_dep_.CoverTab[9979]++
						if c[3]++; c[3] != 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:418
		_go_fuzz_dep_.CoverTab[9983]++
							return
//line /usr/local/go/src/crypto/rsa/rsa.go:419
		// _ = "end of CoverTab[9983]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:420
		_go_fuzz_dep_.CoverTab[9984]++
//line /usr/local/go/src/crypto/rsa/rsa.go:420
		// _ = "end of CoverTab[9984]"
//line /usr/local/go/src/crypto/rsa/rsa.go:420
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:420
	// _ = "end of CoverTab[9979]"
//line /usr/local/go/src/crypto/rsa/rsa.go:420
	_go_fuzz_dep_.CoverTab[9980]++
						if c[2]++; c[2] != 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:421
		_go_fuzz_dep_.CoverTab[9985]++
							return
//line /usr/local/go/src/crypto/rsa/rsa.go:422
		// _ = "end of CoverTab[9985]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:423
		_go_fuzz_dep_.CoverTab[9986]++
//line /usr/local/go/src/crypto/rsa/rsa.go:423
		// _ = "end of CoverTab[9986]"
//line /usr/local/go/src/crypto/rsa/rsa.go:423
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:423
	// _ = "end of CoverTab[9980]"
//line /usr/local/go/src/crypto/rsa/rsa.go:423
	_go_fuzz_dep_.CoverTab[9981]++
						if c[1]++; c[1] != 0 {
//line /usr/local/go/src/crypto/rsa/rsa.go:424
		_go_fuzz_dep_.CoverTab[9987]++
							return
//line /usr/local/go/src/crypto/rsa/rsa.go:425
		// _ = "end of CoverTab[9987]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:426
		_go_fuzz_dep_.CoverTab[9988]++
//line /usr/local/go/src/crypto/rsa/rsa.go:426
		// _ = "end of CoverTab[9988]"
//line /usr/local/go/src/crypto/rsa/rsa.go:426
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:426
	// _ = "end of CoverTab[9981]"
//line /usr/local/go/src/crypto/rsa/rsa.go:426
	_go_fuzz_dep_.CoverTab[9982]++
						c[0]++
//line /usr/local/go/src/crypto/rsa/rsa.go:427
	// _ = "end of CoverTab[9982]"
}

// mgf1XOR XORs the bytes in out with a mask generated using the MGF1 function
//line /usr/local/go/src/crypto/rsa/rsa.go:430
// specified in PKCS #1 v2.1.
//line /usr/local/go/src/crypto/rsa/rsa.go:432
func mgf1XOR(out []byte, hash hash.Hash, seed []byte) {
//line /usr/local/go/src/crypto/rsa/rsa.go:432
	_go_fuzz_dep_.CoverTab[9989]++
						var counter [4]byte
						var digest []byte

						done := 0
						for done < len(out) {
//line /usr/local/go/src/crypto/rsa/rsa.go:437
		_go_fuzz_dep_.CoverTab[9990]++
							hash.Write(seed)
							hash.Write(counter[0:4])
							digest = hash.Sum(digest[:0])
							hash.Reset()

							for i := 0; i < len(digest) && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:443
			_go_fuzz_dep_.CoverTab[9992]++
//line /usr/local/go/src/crypto/rsa/rsa.go:443
			return done < len(out)
//line /usr/local/go/src/crypto/rsa/rsa.go:443
			// _ = "end of CoverTab[9992]"
//line /usr/local/go/src/crypto/rsa/rsa.go:443
		}(); i++ {
//line /usr/local/go/src/crypto/rsa/rsa.go:443
			_go_fuzz_dep_.CoverTab[9993]++
								out[done] ^= digest[i]
								done++
//line /usr/local/go/src/crypto/rsa/rsa.go:445
			// _ = "end of CoverTab[9993]"
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:446
		// _ = "end of CoverTab[9990]"
//line /usr/local/go/src/crypto/rsa/rsa.go:446
		_go_fuzz_dep_.CoverTab[9991]++
							incCounter(&counter)
//line /usr/local/go/src/crypto/rsa/rsa.go:447
		// _ = "end of CoverTab[9991]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:448
	// _ = "end of CoverTab[9989]"
}

// ErrMessageTooLong is returned when attempting to encrypt or sign a message
//line /usr/local/go/src/crypto/rsa/rsa.go:451
// which is too large for the size of the key. When using SignPSS, this can also
//line /usr/local/go/src/crypto/rsa/rsa.go:451
// be returned if the size of the salt is too large.
//line /usr/local/go/src/crypto/rsa/rsa.go:454
var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA key size")

func encrypt(pub *PublicKey, plaintext []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:456
	_go_fuzz_dep_.CoverTab[9994]++
						boring.Unreachable()

						N := bigmod.NewModulusFromBig(pub.N)
						m, err := bigmod.NewNat().SetBytes(plaintext, N)
						if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:461
		_go_fuzz_dep_.CoverTab[9996]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:462
		// _ = "end of CoverTab[9996]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:463
		_go_fuzz_dep_.CoverTab[9997]++
//line /usr/local/go/src/crypto/rsa/rsa.go:463
		// _ = "end of CoverTab[9997]"
//line /usr/local/go/src/crypto/rsa/rsa.go:463
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:463
	// _ = "end of CoverTab[9994]"
//line /usr/local/go/src/crypto/rsa/rsa.go:463
	_go_fuzz_dep_.CoverTab[9995]++
						e := intToBytes(pub.E)

						return bigmod.NewNat().Exp(m, e, N).Bytes(N), nil
//line /usr/local/go/src/crypto/rsa/rsa.go:466
	// _ = "end of CoverTab[9995]"
}

// intToBytes returns i as a big-endian slice of bytes with no leading zeroes,
//line /usr/local/go/src/crypto/rsa/rsa.go:469
// leaking only the bit size of i through timing side-channels.
//line /usr/local/go/src/crypto/rsa/rsa.go:471
func intToBytes(i int) []byte {
//line /usr/local/go/src/crypto/rsa/rsa.go:471
	_go_fuzz_dep_.CoverTab[9998]++
						b := make([]byte, 8)
						binary.BigEndian.PutUint64(b, uint64(i))
						for len(b) > 1 && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:474
		_go_fuzz_dep_.CoverTab[10000]++
//line /usr/local/go/src/crypto/rsa/rsa.go:474
		return b[0] == 0
//line /usr/local/go/src/crypto/rsa/rsa.go:474
		// _ = "end of CoverTab[10000]"
//line /usr/local/go/src/crypto/rsa/rsa.go:474
	}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:474
		_go_fuzz_dep_.CoverTab[10001]++
							b = b[1:]
//line /usr/local/go/src/crypto/rsa/rsa.go:475
		// _ = "end of CoverTab[10001]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:476
	// _ = "end of CoverTab[9998]"
//line /usr/local/go/src/crypto/rsa/rsa.go:476
	_go_fuzz_dep_.CoverTab[9999]++
						return b
//line /usr/local/go/src/crypto/rsa/rsa.go:477
	// _ = "end of CoverTab[9999]"
}

// EncryptOAEP encrypts the given message with RSA-OAEP.
//line /usr/local/go/src/crypto/rsa/rsa.go:480
//
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// OAEP is parameterised by a hash function that is used as a random oracle.
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// Encryption and decryption of a given message must use the same hash function
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// and sha256.New() is a reasonable choice.
//line /usr/local/go/src/crypto/rsa/rsa.go:480
//
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// The random parameter is used as a source of entropy to ensure that
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// encrypting the same message twice doesn't result in the same ciphertext.
//line /usr/local/go/src/crypto/rsa/rsa.go:480
//
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// The label parameter may contain arbitrary data that will not be encrypted,
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// but which gives important context to the message. For example, if a given
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// public key is used to encrypt two types of messages then distinct label
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// values could be used to ensure that a ciphertext for one purpose cannot be
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// used for another by an attacker. If not required it can be empty.
//line /usr/local/go/src/crypto/rsa/rsa.go:480
//
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// The message must be no longer than the length of the public modulus minus
//line /usr/local/go/src/crypto/rsa/rsa.go:480
// twice the hash length, minus a further 2.
//line /usr/local/go/src/crypto/rsa/rsa.go:497
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:497
	_go_fuzz_dep_.CoverTab[10002]++
						if err := checkPub(pub); err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:498
		_go_fuzz_dep_.CoverTab[10008]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:499
		// _ = "end of CoverTab[10008]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:500
		_go_fuzz_dep_.CoverTab[10009]++
//line /usr/local/go/src/crypto/rsa/rsa.go:500
		// _ = "end of CoverTab[10009]"
//line /usr/local/go/src/crypto/rsa/rsa.go:500
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:500
	// _ = "end of CoverTab[10002]"
//line /usr/local/go/src/crypto/rsa/rsa.go:500
	_go_fuzz_dep_.CoverTab[10003]++
						hash.Reset()
						k := pub.Size()
						if len(msg) > k-2*hash.Size()-2 {
//line /usr/local/go/src/crypto/rsa/rsa.go:503
		_go_fuzz_dep_.CoverTab[10010]++
							return nil, ErrMessageTooLong
//line /usr/local/go/src/crypto/rsa/rsa.go:504
		// _ = "end of CoverTab[10010]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:505
		_go_fuzz_dep_.CoverTab[10011]++
//line /usr/local/go/src/crypto/rsa/rsa.go:505
		// _ = "end of CoverTab[10011]"
//line /usr/local/go/src/crypto/rsa/rsa.go:505
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:505
	// _ = "end of CoverTab[10003]"
//line /usr/local/go/src/crypto/rsa/rsa.go:505
	_go_fuzz_dep_.CoverTab[10004]++

						if boring.Enabled && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:507
		_go_fuzz_dep_.CoverTab[10012]++
//line /usr/local/go/src/crypto/rsa/rsa.go:507
		return random == boring.RandReader
//line /usr/local/go/src/crypto/rsa/rsa.go:507
		// _ = "end of CoverTab[10012]"
//line /usr/local/go/src/crypto/rsa/rsa.go:507
	}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:507
		_go_fuzz_dep_.CoverTab[10013]++
							bkey, err := boringPublicKey(pub)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:509
			_go_fuzz_dep_.CoverTab[10015]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:510
			// _ = "end of CoverTab[10015]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:511
			_go_fuzz_dep_.CoverTab[10016]++
//line /usr/local/go/src/crypto/rsa/rsa.go:511
			// _ = "end of CoverTab[10016]"
//line /usr/local/go/src/crypto/rsa/rsa.go:511
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:511
		// _ = "end of CoverTab[10013]"
//line /usr/local/go/src/crypto/rsa/rsa.go:511
		_go_fuzz_dep_.CoverTab[10014]++
							return boring.EncryptRSAOAEP(hash, hash, bkey, msg, label)
//line /usr/local/go/src/crypto/rsa/rsa.go:512
		// _ = "end of CoverTab[10014]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:513
		_go_fuzz_dep_.CoverTab[10017]++
//line /usr/local/go/src/crypto/rsa/rsa.go:513
		// _ = "end of CoverTab[10017]"
//line /usr/local/go/src/crypto/rsa/rsa.go:513
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:513
	// _ = "end of CoverTab[10004]"
//line /usr/local/go/src/crypto/rsa/rsa.go:513
	_go_fuzz_dep_.CoverTab[10005]++
						boring.UnreachableExceptTests()

						hash.Write(label)
						lHash := hash.Sum(nil)
						hash.Reset()

						em := make([]byte, k)
						seed := em[1 : 1+hash.Size()]
						db := em[1+hash.Size():]

						copy(db[0:hash.Size()], lHash)
						db[len(db)-len(msg)-1] = 1
						copy(db[len(db)-len(msg):], msg)

						_, err := io.ReadFull(random, seed)
						if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:529
		_go_fuzz_dep_.CoverTab[10018]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:530
		// _ = "end of CoverTab[10018]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:531
		_go_fuzz_dep_.CoverTab[10019]++
//line /usr/local/go/src/crypto/rsa/rsa.go:531
		// _ = "end of CoverTab[10019]"
//line /usr/local/go/src/crypto/rsa/rsa.go:531
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:531
	// _ = "end of CoverTab[10005]"
//line /usr/local/go/src/crypto/rsa/rsa.go:531
	_go_fuzz_dep_.CoverTab[10006]++

						mgf1XOR(db, hash, seed)
						mgf1XOR(seed, hash, db)

						if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/rsa.go:536
		_go_fuzz_dep_.CoverTab[10020]++
							var bkey *boring.PublicKeyRSA
							bkey, err = boringPublicKey(pub)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:539
			_go_fuzz_dep_.CoverTab[10022]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:540
			// _ = "end of CoverTab[10022]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:541
			_go_fuzz_dep_.CoverTab[10023]++
//line /usr/local/go/src/crypto/rsa/rsa.go:541
			// _ = "end of CoverTab[10023]"
//line /usr/local/go/src/crypto/rsa/rsa.go:541
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:541
		// _ = "end of CoverTab[10020]"
//line /usr/local/go/src/crypto/rsa/rsa.go:541
		_go_fuzz_dep_.CoverTab[10021]++
							return boring.EncryptRSANoPadding(bkey, em)
//line /usr/local/go/src/crypto/rsa/rsa.go:542
		// _ = "end of CoverTab[10021]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:543
		_go_fuzz_dep_.CoverTab[10024]++
//line /usr/local/go/src/crypto/rsa/rsa.go:543
		// _ = "end of CoverTab[10024]"
//line /usr/local/go/src/crypto/rsa/rsa.go:543
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:543
	// _ = "end of CoverTab[10006]"
//line /usr/local/go/src/crypto/rsa/rsa.go:543
	_go_fuzz_dep_.CoverTab[10007]++

						return encrypt(pub, em)
//line /usr/local/go/src/crypto/rsa/rsa.go:545
	// _ = "end of CoverTab[10007]"
}

// ErrDecryption represents a failure to decrypt a message.
//line /usr/local/go/src/crypto/rsa/rsa.go:548
// It is deliberately vague to avoid adaptive attacks.
//line /usr/local/go/src/crypto/rsa/rsa.go:550
var ErrDecryption = errors.New("crypto/rsa: decryption error")

// ErrVerification represents a failure to verify a signature.
//line /usr/local/go/src/crypto/rsa/rsa.go:552
// It is deliberately vague to avoid adaptive attacks.
//line /usr/local/go/src/crypto/rsa/rsa.go:554
var ErrVerification = errors.New("crypto/rsa: verification error")

// Precompute performs some calculations that speed up private key operations
//line /usr/local/go/src/crypto/rsa/rsa.go:556
// in the future.
//line /usr/local/go/src/crypto/rsa/rsa.go:558
func (priv *PrivateKey) Precompute() {
//line /usr/local/go/src/crypto/rsa/rsa.go:558
	_go_fuzz_dep_.CoverTab[10025]++
						if priv.Precomputed.n == nil && func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:559
		_go_fuzz_dep_.CoverTab[10028]++
//line /usr/local/go/src/crypto/rsa/rsa.go:559
		return len(priv.Primes) == 2
//line /usr/local/go/src/crypto/rsa/rsa.go:559
		// _ = "end of CoverTab[10028]"
//line /usr/local/go/src/crypto/rsa/rsa.go:559
	}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:559
		_go_fuzz_dep_.CoverTab[10029]++
							priv.Precomputed.n = bigmod.NewModulusFromBig(priv.N)
							priv.Precomputed.p = bigmod.NewModulusFromBig(priv.Primes[0])
							priv.Precomputed.q = bigmod.NewModulusFromBig(priv.Primes[1])
//line /usr/local/go/src/crypto/rsa/rsa.go:562
		// _ = "end of CoverTab[10029]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:563
		_go_fuzz_dep_.CoverTab[10030]++
//line /usr/local/go/src/crypto/rsa/rsa.go:563
		// _ = "end of CoverTab[10030]"
//line /usr/local/go/src/crypto/rsa/rsa.go:563
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:563
	// _ = "end of CoverTab[10025]"
//line /usr/local/go/src/crypto/rsa/rsa.go:563
	_go_fuzz_dep_.CoverTab[10026]++

//line /usr/local/go/src/crypto/rsa/rsa.go:566
	if priv.Precomputed.Dp != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:566
		_go_fuzz_dep_.CoverTab[10031]++
							return
//line /usr/local/go/src/crypto/rsa/rsa.go:567
		// _ = "end of CoverTab[10031]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:568
		_go_fuzz_dep_.CoverTab[10032]++
//line /usr/local/go/src/crypto/rsa/rsa.go:568
		// _ = "end of CoverTab[10032]"
//line /usr/local/go/src/crypto/rsa/rsa.go:568
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:568
	// _ = "end of CoverTab[10026]"
//line /usr/local/go/src/crypto/rsa/rsa.go:568
	_go_fuzz_dep_.CoverTab[10027]++

						priv.Precomputed.Dp = new(big.Int).Sub(priv.Primes[0], bigOne)
						priv.Precomputed.Dp.Mod(priv.D, priv.Precomputed.Dp)

						priv.Precomputed.Dq = new(big.Int).Sub(priv.Primes[1], bigOne)
						priv.Precomputed.Dq.Mod(priv.D, priv.Precomputed.Dq)

						priv.Precomputed.Qinv = new(big.Int).ModInverse(priv.Primes[1], priv.Primes[0])

						r := new(big.Int).Mul(priv.Primes[0], priv.Primes[1])
						priv.Precomputed.CRTValues = make([]CRTValue, len(priv.Primes)-2)
						for i := 2; i < len(priv.Primes); i++ {
//line /usr/local/go/src/crypto/rsa/rsa.go:580
		_go_fuzz_dep_.CoverTab[10033]++
							prime := priv.Primes[i]
							values := &priv.Precomputed.CRTValues[i-2]

							values.Exp = new(big.Int).Sub(prime, bigOne)
							values.Exp.Mod(priv.D, values.Exp)

							values.R = new(big.Int).Set(r)
							values.Coeff = new(big.Int).ModInverse(r, prime)

							r.Mul(r, prime)
//line /usr/local/go/src/crypto/rsa/rsa.go:590
		// _ = "end of CoverTab[10033]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:591
	// _ = "end of CoverTab[10027]"
}

const withCheck = true
const noCheck = false

// decrypt performs an RSA decryption of ciphertext into out. If check is true,
//line /usr/local/go/src/crypto/rsa/rsa.go:597
// m^e is calculated and compared with ciphertext, in order to defend against
//line /usr/local/go/src/crypto/rsa/rsa.go:597
// errors in the CRT computation.
//line /usr/local/go/src/crypto/rsa/rsa.go:600
func decrypt(priv *PrivateKey, ciphertext []byte, check bool) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:600
	_go_fuzz_dep_.CoverTab[10034]++
						if len(priv.Primes) <= 2 {
//line /usr/local/go/src/crypto/rsa/rsa.go:601
		_go_fuzz_dep_.CoverTab[10038]++
							boring.Unreachable()
//line /usr/local/go/src/crypto/rsa/rsa.go:602
		// _ = "end of CoverTab[10038]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:603
		_go_fuzz_dep_.CoverTab[10039]++
//line /usr/local/go/src/crypto/rsa/rsa.go:603
		// _ = "end of CoverTab[10039]"
//line /usr/local/go/src/crypto/rsa/rsa.go:603
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:603
	// _ = "end of CoverTab[10034]"
//line /usr/local/go/src/crypto/rsa/rsa.go:603
	_go_fuzz_dep_.CoverTab[10035]++

						var (
		err	error
		m, c	*bigmod.Nat
		N	*bigmod.Modulus
		t0	= bigmod.NewNat()
	)
	if priv.Precomputed.n == nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:611
		_go_fuzz_dep_.CoverTab[10040]++
							N = bigmod.NewModulusFromBig(priv.N)
							c, err = bigmod.NewNat().SetBytes(ciphertext, N)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:614
			_go_fuzz_dep_.CoverTab[10042]++
								return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:615
			// _ = "end of CoverTab[10042]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:616
			_go_fuzz_dep_.CoverTab[10043]++
//line /usr/local/go/src/crypto/rsa/rsa.go:616
			// _ = "end of CoverTab[10043]"
//line /usr/local/go/src/crypto/rsa/rsa.go:616
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:616
		// _ = "end of CoverTab[10040]"
//line /usr/local/go/src/crypto/rsa/rsa.go:616
		_go_fuzz_dep_.CoverTab[10041]++
							m = bigmod.NewNat().Exp(c, priv.D.Bytes(), N)
//line /usr/local/go/src/crypto/rsa/rsa.go:617
		// _ = "end of CoverTab[10041]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:618
		_go_fuzz_dep_.CoverTab[10044]++
							N = priv.Precomputed.n
							P, Q := priv.Precomputed.p, priv.Precomputed.q
							Qinv, err := bigmod.NewNat().SetBytes(priv.Precomputed.Qinv.Bytes(), P)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:622
			_go_fuzz_dep_.CoverTab[10047]++
								return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:623
			// _ = "end of CoverTab[10047]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:624
			_go_fuzz_dep_.CoverTab[10048]++
//line /usr/local/go/src/crypto/rsa/rsa.go:624
			// _ = "end of CoverTab[10048]"
//line /usr/local/go/src/crypto/rsa/rsa.go:624
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:624
		// _ = "end of CoverTab[10044]"
//line /usr/local/go/src/crypto/rsa/rsa.go:624
		_go_fuzz_dep_.CoverTab[10045]++
							c, err = bigmod.NewNat().SetBytes(ciphertext, N)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:626
			_go_fuzz_dep_.CoverTab[10049]++
								return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:627
			// _ = "end of CoverTab[10049]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:628
			_go_fuzz_dep_.CoverTab[10050]++
//line /usr/local/go/src/crypto/rsa/rsa.go:628
			// _ = "end of CoverTab[10050]"
//line /usr/local/go/src/crypto/rsa/rsa.go:628
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:628
		// _ = "end of CoverTab[10045]"
//line /usr/local/go/src/crypto/rsa/rsa.go:628
		_go_fuzz_dep_.CoverTab[10046]++

//line /usr/local/go/src/crypto/rsa/rsa.go:631
		m = bigmod.NewNat().Exp(t0.Mod(c, P), priv.Precomputed.Dp.Bytes(), P)

							m2 := bigmod.NewNat().Exp(t0.Mod(c, Q), priv.Precomputed.Dq.Bytes(), Q)

							m.Sub(t0.Mod(m2, P), P)

							m.Mul(Qinv, P)

							m.ExpandFor(N).Mul(t0.Mod(Q.Nat(), N), N)

							m.Add(m2.ExpandFor(N), N)
//line /usr/local/go/src/crypto/rsa/rsa.go:641
		// _ = "end of CoverTab[10046]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:642
	// _ = "end of CoverTab[10035]"
//line /usr/local/go/src/crypto/rsa/rsa.go:642
	_go_fuzz_dep_.CoverTab[10036]++

						if check {
//line /usr/local/go/src/crypto/rsa/rsa.go:644
		_go_fuzz_dep_.CoverTab[10051]++
							c1 := bigmod.NewNat().Exp(m, intToBytes(priv.E), N)
							if c1.Equal(c) != 1 {
//line /usr/local/go/src/crypto/rsa/rsa.go:646
			_go_fuzz_dep_.CoverTab[10052]++
								return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:647
			// _ = "end of CoverTab[10052]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:648
			_go_fuzz_dep_.CoverTab[10053]++
//line /usr/local/go/src/crypto/rsa/rsa.go:648
			// _ = "end of CoverTab[10053]"
//line /usr/local/go/src/crypto/rsa/rsa.go:648
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:648
		// _ = "end of CoverTab[10051]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:649
		_go_fuzz_dep_.CoverTab[10054]++
//line /usr/local/go/src/crypto/rsa/rsa.go:649
		// _ = "end of CoverTab[10054]"
//line /usr/local/go/src/crypto/rsa/rsa.go:649
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:649
	// _ = "end of CoverTab[10036]"
//line /usr/local/go/src/crypto/rsa/rsa.go:649
	_go_fuzz_dep_.CoverTab[10037]++

						return m.Bytes(N), nil
//line /usr/local/go/src/crypto/rsa/rsa.go:651
	// _ = "end of CoverTab[10037]"
}

// DecryptOAEP decrypts ciphertext using RSA-OAEP.
//line /usr/local/go/src/crypto/rsa/rsa.go:654
//
//line /usr/local/go/src/crypto/rsa/rsa.go:654
// OAEP is parameterised by a hash function that is used as a random oracle.
//line /usr/local/go/src/crypto/rsa/rsa.go:654
// Encryption and decryption of a given message must use the same hash function
//line /usr/local/go/src/crypto/rsa/rsa.go:654
// and sha256.New() is a reasonable choice.
//line /usr/local/go/src/crypto/rsa/rsa.go:654
//
//line /usr/local/go/src/crypto/rsa/rsa.go:654
// The random parameter is legacy and ignored, and it can be as nil.
//line /usr/local/go/src/crypto/rsa/rsa.go:654
//
//line /usr/local/go/src/crypto/rsa/rsa.go:654
// The label parameter must match the value given when encrypting. See
//line /usr/local/go/src/crypto/rsa/rsa.go:654
// EncryptOAEP for details.
//line /usr/local/go/src/crypto/rsa/rsa.go:664
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:664
	_go_fuzz_dep_.CoverTab[10055]++
						return decryptOAEP(hash, hash, random, priv, ciphertext, label)
//line /usr/local/go/src/crypto/rsa/rsa.go:665
	// _ = "end of CoverTab[10055]"
}

func decryptOAEP(hash, mgfHash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/rsa.go:668
	_go_fuzz_dep_.CoverTab[10056]++
						if err := checkPub(&priv.PublicKey); err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:669
		_go_fuzz_dep_.CoverTab[10063]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:670
		// _ = "end of CoverTab[10063]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:671
		_go_fuzz_dep_.CoverTab[10064]++
//line /usr/local/go/src/crypto/rsa/rsa.go:671
		// _ = "end of CoverTab[10064]"
//line /usr/local/go/src/crypto/rsa/rsa.go:671
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:671
	// _ = "end of CoverTab[10056]"
//line /usr/local/go/src/crypto/rsa/rsa.go:671
	_go_fuzz_dep_.CoverTab[10057]++
						k := priv.Size()
						if len(ciphertext) > k || func() bool {
//line /usr/local/go/src/crypto/rsa/rsa.go:673
		_go_fuzz_dep_.CoverTab[10065]++
//line /usr/local/go/src/crypto/rsa/rsa.go:673
		return k < hash.Size()*2+2
							// _ = "end of CoverTab[10065]"
//line /usr/local/go/src/crypto/rsa/rsa.go:674
	}() {
//line /usr/local/go/src/crypto/rsa/rsa.go:674
		_go_fuzz_dep_.CoverTab[10066]++
							return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:675
		// _ = "end of CoverTab[10066]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:676
		_go_fuzz_dep_.CoverTab[10067]++
//line /usr/local/go/src/crypto/rsa/rsa.go:676
		// _ = "end of CoverTab[10067]"
//line /usr/local/go/src/crypto/rsa/rsa.go:676
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:676
	// _ = "end of CoverTab[10057]"
//line /usr/local/go/src/crypto/rsa/rsa.go:676
	_go_fuzz_dep_.CoverTab[10058]++

						if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/rsa.go:678
		_go_fuzz_dep_.CoverTab[10068]++
							bkey, err := boringPrivateKey(priv)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:680
			_go_fuzz_dep_.CoverTab[10071]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:681
			// _ = "end of CoverTab[10071]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:682
			_go_fuzz_dep_.CoverTab[10072]++
//line /usr/local/go/src/crypto/rsa/rsa.go:682
			// _ = "end of CoverTab[10072]"
//line /usr/local/go/src/crypto/rsa/rsa.go:682
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:682
		// _ = "end of CoverTab[10068]"
//line /usr/local/go/src/crypto/rsa/rsa.go:682
		_go_fuzz_dep_.CoverTab[10069]++
							out, err := boring.DecryptRSAOAEP(hash, mgfHash, bkey, ciphertext, label)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:684
			_go_fuzz_dep_.CoverTab[10073]++
								return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:685
			// _ = "end of CoverTab[10073]"
		} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:686
			_go_fuzz_dep_.CoverTab[10074]++
//line /usr/local/go/src/crypto/rsa/rsa.go:686
			// _ = "end of CoverTab[10074]"
//line /usr/local/go/src/crypto/rsa/rsa.go:686
		}
//line /usr/local/go/src/crypto/rsa/rsa.go:686
		// _ = "end of CoverTab[10069]"
//line /usr/local/go/src/crypto/rsa/rsa.go:686
		_go_fuzz_dep_.CoverTab[10070]++
							return out, nil
//line /usr/local/go/src/crypto/rsa/rsa.go:687
		// _ = "end of CoverTab[10070]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:688
		_go_fuzz_dep_.CoverTab[10075]++
//line /usr/local/go/src/crypto/rsa/rsa.go:688
		// _ = "end of CoverTab[10075]"
//line /usr/local/go/src/crypto/rsa/rsa.go:688
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:688
	// _ = "end of CoverTab[10058]"
//line /usr/local/go/src/crypto/rsa/rsa.go:688
	_go_fuzz_dep_.CoverTab[10059]++

						em, err := decrypt(priv, ciphertext, noCheck)
						if err != nil {
//line /usr/local/go/src/crypto/rsa/rsa.go:691
		_go_fuzz_dep_.CoverTab[10076]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/rsa.go:692
		// _ = "end of CoverTab[10076]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:693
		_go_fuzz_dep_.CoverTab[10077]++
//line /usr/local/go/src/crypto/rsa/rsa.go:693
		// _ = "end of CoverTab[10077]"
//line /usr/local/go/src/crypto/rsa/rsa.go:693
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:693
	// _ = "end of CoverTab[10059]"
//line /usr/local/go/src/crypto/rsa/rsa.go:693
	_go_fuzz_dep_.CoverTab[10060]++

						hash.Write(label)
						lHash := hash.Sum(nil)
						hash.Reset()

						firstByteIsZero := subtle.ConstantTimeByteEq(em[0], 0)

						seed := em[1 : hash.Size()+1]
						db := em[hash.Size()+1:]

						mgf1XOR(seed, mgfHash, db)
						mgf1XOR(db, mgfHash, seed)

						lHash2 := db[0:hash.Size()]

//line /usr/local/go/src/crypto/rsa/rsa.go:713
	lHash2Good := subtle.ConstantTimeCompare(lHash, lHash2)

	// The remainder of the plaintext must be zero or more 0x00, followed
	// by 0x01, followed by the message.
	//   lookingForIndex: 1 iff we are still looking for the 0x01
	//   index: the offset of the first 0x01 byte
	//   invalid: 1 iff we saw a non-zero byte before the 0x01.
	var lookingForIndex, index, invalid int
	lookingForIndex = 1
	rest := db[hash.Size():]

	for i := 0; i < len(rest); i++ {
//line /usr/local/go/src/crypto/rsa/rsa.go:724
		_go_fuzz_dep_.CoverTab[10078]++
							equals0 := subtle.ConstantTimeByteEq(rest[i], 0)
							equals1 := subtle.ConstantTimeByteEq(rest[i], 1)
							index = subtle.ConstantTimeSelect(lookingForIndex&equals1, i, index)
							lookingForIndex = subtle.ConstantTimeSelect(equals1, 0, lookingForIndex)
							invalid = subtle.ConstantTimeSelect(lookingForIndex&^equals0, 1, invalid)
//line /usr/local/go/src/crypto/rsa/rsa.go:729
		// _ = "end of CoverTab[10078]"
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:730
	// _ = "end of CoverTab[10060]"
//line /usr/local/go/src/crypto/rsa/rsa.go:730
	_go_fuzz_dep_.CoverTab[10061]++

						if firstByteIsZero&lHash2Good&^invalid&^lookingForIndex != 1 {
//line /usr/local/go/src/crypto/rsa/rsa.go:732
		_go_fuzz_dep_.CoverTab[10079]++
							return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/rsa.go:733
		// _ = "end of CoverTab[10079]"
	} else {
//line /usr/local/go/src/crypto/rsa/rsa.go:734
		_go_fuzz_dep_.CoverTab[10080]++
//line /usr/local/go/src/crypto/rsa/rsa.go:734
		// _ = "end of CoverTab[10080]"
//line /usr/local/go/src/crypto/rsa/rsa.go:734
	}
//line /usr/local/go/src/crypto/rsa/rsa.go:734
	// _ = "end of CoverTab[10061]"
//line /usr/local/go/src/crypto/rsa/rsa.go:734
	_go_fuzz_dep_.CoverTab[10062]++

						return rest[index+1:], nil
//line /usr/local/go/src/crypto/rsa/rsa.go:736
	// _ = "end of CoverTab[10062]"
}

//line /usr/local/go/src/crypto/rsa/rsa.go:737
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rsa/rsa.go:737
var _ = _go_fuzz_dep_.CoverTab
