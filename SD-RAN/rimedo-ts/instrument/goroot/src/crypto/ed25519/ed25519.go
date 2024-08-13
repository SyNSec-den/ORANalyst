// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// Package ed25519 implements the Ed25519 signature algorithm. See
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// https://ed25519.cr.yp.to/.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
//
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// These functions are also compatible with the “Ed25519” function defined in
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// RFC 8032. However, unlike RFC 8032's formulation, this package's private key
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// representation includes a public key suffix to make multiple signing
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// operations with the same key more efficient. This package refers to the RFC
//line /usr/local/go/src/crypto/ed25519/ed25519.go:5
// 8032 private key as the “seed”.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
package ed25519

//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
import (
//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
)
//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
import (
//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:13
)

import (
	"bytes"
	"crypto"
	"crypto/internal/edwards25519"
	cryptorand "crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
	"strconv"
)

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize	= 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize	= 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize	= 64
	// SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.
	SeedSize	= 32
)

// PublicKey is the type of Ed25519 public keys.
type PublicKey []byte

//line /usr/local/go/src/crypto/ed25519/ed25519.go:43
// Equal reports whether pub and x have the same value.
func (pub PublicKey) Equal(x crypto.PublicKey) bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:44
	_go_fuzz_dep_.CoverTab[9395]++
							xx, ok := x.(PublicKey)
							if !ok {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:46
		_go_fuzz_dep_.CoverTab[9397]++
								return false
//line /usr/local/go/src/crypto/ed25519/ed25519.go:47
		// _ = "end of CoverTab[9397]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:48
		_go_fuzz_dep_.CoverTab[9398]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:48
		// _ = "end of CoverTab[9398]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:48
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:48
	// _ = "end of CoverTab[9395]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:48
	_go_fuzz_dep_.CoverTab[9396]++
							return bytes.Equal(pub, xx)
//line /usr/local/go/src/crypto/ed25519/ed25519.go:49
	// _ = "end of CoverTab[9396]"
}

// PrivateKey is the type of Ed25519 private keys. It implements [crypto.Signer].
type PrivateKey []byte

// Public returns the [PublicKey] corresponding to priv.
func (priv PrivateKey) Public() crypto.PublicKey {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:56
	_go_fuzz_dep_.CoverTab[9399]++
							publicKey := make([]byte, PublicKeySize)
							copy(publicKey, priv[32:])
							return PublicKey(publicKey)
//line /usr/local/go/src/crypto/ed25519/ed25519.go:59
	// _ = "end of CoverTab[9399]"
}

// Equal reports whether priv and x have the same value.
func (priv PrivateKey) Equal(x crypto.PrivateKey) bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:63
	_go_fuzz_dep_.CoverTab[9400]++
							xx, ok := x.(PrivateKey)
							if !ok {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:65
		_go_fuzz_dep_.CoverTab[9402]++
								return false
//line /usr/local/go/src/crypto/ed25519/ed25519.go:66
		// _ = "end of CoverTab[9402]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:67
		_go_fuzz_dep_.CoverTab[9403]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:67
		// _ = "end of CoverTab[9403]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:67
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:67
	// _ = "end of CoverTab[9400]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:67
	_go_fuzz_dep_.CoverTab[9401]++
							return bytes.Equal(priv, xx)
//line /usr/local/go/src/crypto/ed25519/ed25519.go:68
	// _ = "end of CoverTab[9401]"
}

// Seed returns the private key seed corresponding to priv. It is provided for
//line /usr/local/go/src/crypto/ed25519/ed25519.go:71
// interoperability with RFC 8032. RFC 8032's private keys correspond to seeds
//line /usr/local/go/src/crypto/ed25519/ed25519.go:71
// in this package.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:74
func (priv PrivateKey) Seed() []byte {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:74
	_go_fuzz_dep_.CoverTab[9404]++
							return bytes.Clone(priv[:SeedSize])
//line /usr/local/go/src/crypto/ed25519/ed25519.go:75
	// _ = "end of CoverTab[9404]"
}

// Sign signs the given message with priv. rand is ignored.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
//
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
// If opts.HashFunc() is [crypto.SHA512], the pre-hashed variant Ed25519ph is used
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
// and message is expected to be a SHA-512 hash, otherwise opts.HashFunc() must
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
// be [crypto.Hash](0) and the message must not be hashed, as Ed25519 performs two
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
// passes over messages to be signed.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
//
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
// A value of type [Options] can be used as opts, or crypto.Hash(0) or
//line /usr/local/go/src/crypto/ed25519/ed25519.go:78
// crypto.SHA512 directly to select plain Ed25519 or Ed25519ph, respectively.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:87
func (priv PrivateKey) Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error) {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:87
	_go_fuzz_dep_.CoverTab[9405]++
							hash := opts.HashFunc()
							context := ""
							if opts, ok := opts.(*Options); ok {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:90
		_go_fuzz_dep_.CoverTab[9408]++
								context = opts.Context
//line /usr/local/go/src/crypto/ed25519/ed25519.go:91
		// _ = "end of CoverTab[9408]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:92
		_go_fuzz_dep_.CoverTab[9409]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:92
		// _ = "end of CoverTab[9409]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:92
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:92
	// _ = "end of CoverTab[9405]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:92
	_go_fuzz_dep_.CoverTab[9406]++
							if l := len(context); l > 255 {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:93
		_go_fuzz_dep_.CoverTab[9410]++
								return nil, errors.New("ed25519: bad Ed25519ph context length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:94
		// _ = "end of CoverTab[9410]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:95
		_go_fuzz_dep_.CoverTab[9411]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:95
		// _ = "end of CoverTab[9411]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:95
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:95
	// _ = "end of CoverTab[9406]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:95
	_go_fuzz_dep_.CoverTab[9407]++
							switch {
	case hash == crypto.SHA512:
//line /usr/local/go/src/crypto/ed25519/ed25519.go:97
		_go_fuzz_dep_.CoverTab[9412]++
								if l := len(message); l != sha512.Size {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:98
			_go_fuzz_dep_.CoverTab[9417]++
									return nil, errors.New("ed25519: bad Ed25519ph message hash length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:99
			// _ = "end of CoverTab[9417]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:100
			_go_fuzz_dep_.CoverTab[9418]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:100
			// _ = "end of CoverTab[9418]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:100
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:100
		// _ = "end of CoverTab[9412]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:100
		_go_fuzz_dep_.CoverTab[9413]++
								signature := make([]byte, SignatureSize)
								sign(signature, priv, message, domPrefixPh, context)
								return signature, nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:103
		// _ = "end of CoverTab[9413]"
	case hash == crypto.Hash(0) && func() bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:104
		_go_fuzz_dep_.CoverTab[9419]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:104
		return context != ""
//line /usr/local/go/src/crypto/ed25519/ed25519.go:104
		// _ = "end of CoverTab[9419]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:104
	}():
//line /usr/local/go/src/crypto/ed25519/ed25519.go:104
		_go_fuzz_dep_.CoverTab[9414]++
								signature := make([]byte, SignatureSize)
								sign(signature, priv, message, domPrefixCtx, context)
								return signature, nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:107
		// _ = "end of CoverTab[9414]"
	case hash == crypto.Hash(0):
//line /usr/local/go/src/crypto/ed25519/ed25519.go:108
		_go_fuzz_dep_.CoverTab[9415]++
								return Sign(priv, message), nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:109
		// _ = "end of CoverTab[9415]"
	default:
//line /usr/local/go/src/crypto/ed25519/ed25519.go:110
		_go_fuzz_dep_.CoverTab[9416]++
								return nil, errors.New("ed25519: expected opts.HashFunc() zero (unhashed message, for standard Ed25519) or SHA-512 (for Ed25519ph)")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:111
		// _ = "end of CoverTab[9416]"
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:112
	// _ = "end of CoverTab[9407]"
}

// Options can be used with [PrivateKey.Sign] or [VerifyWithOptions]
//line /usr/local/go/src/crypto/ed25519/ed25519.go:115
// to select Ed25519 variants.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:117
type Options struct {
	// Hash can be zero for regular Ed25519, or crypto.SHA512 for Ed25519ph.
	Hash	crypto.Hash

	// Context, if not empty, selects Ed25519ctx or provides the context string
	// for Ed25519ph. It can be at most 255 bytes in length.
	Context	string
}

// HashFunc returns o.Hash.
func (o *Options) HashFunc() crypto.Hash {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:127
	_go_fuzz_dep_.CoverTab[9420]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:127
	return o.Hash
//line /usr/local/go/src/crypto/ed25519/ed25519.go:127
	// _ = "end of CoverTab[9420]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:127
}

// GenerateKey generates a public/private key pair using entropy from rand.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:129
// If rand is nil, [crypto/rand.Reader] will be used.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:131
func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error) {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:131
	_go_fuzz_dep_.CoverTab[9421]++
							if rand == nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:132
		_go_fuzz_dep_.CoverTab[9424]++
								rand = cryptorand.Reader
//line /usr/local/go/src/crypto/ed25519/ed25519.go:133
		// _ = "end of CoverTab[9424]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:134
		_go_fuzz_dep_.CoverTab[9425]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:134
		// _ = "end of CoverTab[9425]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:134
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:134
	// _ = "end of CoverTab[9421]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:134
	_go_fuzz_dep_.CoverTab[9422]++

							seed := make([]byte, SeedSize)
							if _, err := io.ReadFull(rand, seed); err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:137
		_go_fuzz_dep_.CoverTab[9426]++
								return nil, nil, err
//line /usr/local/go/src/crypto/ed25519/ed25519.go:138
		// _ = "end of CoverTab[9426]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:139
		_go_fuzz_dep_.CoverTab[9427]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:139
		// _ = "end of CoverTab[9427]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:139
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:139
	// _ = "end of CoverTab[9422]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:139
	_go_fuzz_dep_.CoverTab[9423]++

							privateKey := NewKeyFromSeed(seed)
							publicKey := make([]byte, PublicKeySize)
							copy(publicKey, privateKey[32:])

							return publicKey, privateKey, nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:145
	// _ = "end of CoverTab[9423]"
}

// NewKeyFromSeed calculates a private key from a seed. It will panic if
//line /usr/local/go/src/crypto/ed25519/ed25519.go:148
// len(seed) is not [SeedSize]. This function is provided for interoperability
//line /usr/local/go/src/crypto/ed25519/ed25519.go:148
// with RFC 8032. RFC 8032's private keys correspond to seeds in this
//line /usr/local/go/src/crypto/ed25519/ed25519.go:148
// package.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:152
func NewKeyFromSeed(seed []byte) PrivateKey {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:152
	_go_fuzz_dep_.CoverTab[9428]++

							privateKey := make([]byte, PrivateKeySize)
							newKeyFromSeed(privateKey, seed)
							return privateKey
//line /usr/local/go/src/crypto/ed25519/ed25519.go:156
	// _ = "end of CoverTab[9428]"
}

func newKeyFromSeed(privateKey, seed []byte) {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:159
	_go_fuzz_dep_.CoverTab[9429]++
							if l := len(seed); l != SeedSize {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:160
		_go_fuzz_dep_.CoverTab[9432]++
								panic("ed25519: bad seed length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:161
		// _ = "end of CoverTab[9432]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:162
		_go_fuzz_dep_.CoverTab[9433]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:162
		// _ = "end of CoverTab[9433]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:162
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:162
	// _ = "end of CoverTab[9429]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:162
	_go_fuzz_dep_.CoverTab[9430]++

							h := sha512.Sum512(seed)
							s, err := edwards25519.NewScalar().SetBytesWithClamping(h[:32])
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:166
		_go_fuzz_dep_.CoverTab[9434]++
								panic("ed25519: internal error: setting scalar failed")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:167
		// _ = "end of CoverTab[9434]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:168
		_go_fuzz_dep_.CoverTab[9435]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:168
		// _ = "end of CoverTab[9435]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:168
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:168
	// _ = "end of CoverTab[9430]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:168
	_go_fuzz_dep_.CoverTab[9431]++
							A := (&edwards25519.Point{}).ScalarBaseMult(s)

							publicKey := A.Bytes()

							copy(privateKey, seed)
							copy(privateKey[32:], publicKey)
//line /usr/local/go/src/crypto/ed25519/ed25519.go:174
	// _ = "end of CoverTab[9431]"
}

// Sign signs the message with privateKey and returns a signature. It will
//line /usr/local/go/src/crypto/ed25519/ed25519.go:177
// panic if len(privateKey) is not [PrivateKeySize].
//line /usr/local/go/src/crypto/ed25519/ed25519.go:179
func Sign(privateKey PrivateKey, message []byte) []byte {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:179
	_go_fuzz_dep_.CoverTab[9436]++

//line /usr/local/go/src/crypto/ed25519/ed25519.go:182
	signature := make([]byte, SignatureSize)
							sign(signature, privateKey, message, domPrefixPure, "")
							return signature
//line /usr/local/go/src/crypto/ed25519/ed25519.go:184
	// _ = "end of CoverTab[9436]"
}

// Domain separation prefixes used to disambiguate Ed25519/Ed25519ph/Ed25519ctx.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:187
// See RFC 8032, Section 2 and Section 5.1.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:189
const (
	// domPrefixPure is empty for pure Ed25519.
	domPrefixPure	= ""
	// domPrefixPh is dom2(phflag=1) for Ed25519ph. It must be followed by the
	// uint8-length prefixed context.
	domPrefixPh	= "SigEd25519 no Ed25519 collisions\x01"
	// domPrefixCtx is dom2(phflag=0) for Ed25519ctx. It must be followed by the
	// uint8-length prefixed context.
	domPrefixCtx	= "SigEd25519 no Ed25519 collisions\x00"
)

func sign(signature, privateKey, message []byte, domPrefix, context string) {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:200
	_go_fuzz_dep_.CoverTab[9437]++
							if l := len(privateKey); l != PrivateKeySize {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:201
		_go_fuzz_dep_.CoverTab[9444]++
								panic("ed25519: bad private key length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:202
		// _ = "end of CoverTab[9444]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:203
		_go_fuzz_dep_.CoverTab[9445]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:203
		// _ = "end of CoverTab[9445]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:203
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:203
	// _ = "end of CoverTab[9437]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:203
	_go_fuzz_dep_.CoverTab[9438]++
							seed, publicKey := privateKey[:SeedSize], privateKey[SeedSize:]

							h := sha512.Sum512(seed)
							s, err := edwards25519.NewScalar().SetBytesWithClamping(h[:32])
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:208
		_go_fuzz_dep_.CoverTab[9446]++
								panic("ed25519: internal error: setting scalar failed")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:209
		// _ = "end of CoverTab[9446]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:210
		_go_fuzz_dep_.CoverTab[9447]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:210
		// _ = "end of CoverTab[9447]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:210
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:210
	// _ = "end of CoverTab[9438]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:210
	_go_fuzz_dep_.CoverTab[9439]++
							prefix := h[32:]

							mh := sha512.New()
							if domPrefix != domPrefixPure {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:214
		_go_fuzz_dep_.CoverTab[9448]++
								mh.Write([]byte(domPrefix))
								mh.Write([]byte{byte(len(context))})
								mh.Write([]byte(context))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:217
		// _ = "end of CoverTab[9448]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:218
		_go_fuzz_dep_.CoverTab[9449]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:218
		// _ = "end of CoverTab[9449]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:218
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:218
	// _ = "end of CoverTab[9439]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:218
	_go_fuzz_dep_.CoverTab[9440]++
							mh.Write(prefix)
							mh.Write(message)
							messageDigest := make([]byte, 0, sha512.Size)
							messageDigest = mh.Sum(messageDigest)
							r, err := edwards25519.NewScalar().SetUniformBytes(messageDigest)
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:224
		_go_fuzz_dep_.CoverTab[9450]++
								panic("ed25519: internal error: setting scalar failed")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:225
		// _ = "end of CoverTab[9450]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:226
		_go_fuzz_dep_.CoverTab[9451]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:226
		// _ = "end of CoverTab[9451]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:226
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:226
	// _ = "end of CoverTab[9440]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:226
	_go_fuzz_dep_.CoverTab[9441]++

							R := (&edwards25519.Point{}).ScalarBaseMult(r)

							kh := sha512.New()
							if domPrefix != domPrefixPure {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:231
		_go_fuzz_dep_.CoverTab[9452]++
								kh.Write([]byte(domPrefix))
								kh.Write([]byte{byte(len(context))})
								kh.Write([]byte(context))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:234
		// _ = "end of CoverTab[9452]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:235
		_go_fuzz_dep_.CoverTab[9453]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:235
		// _ = "end of CoverTab[9453]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:235
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:235
	// _ = "end of CoverTab[9441]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:235
	_go_fuzz_dep_.CoverTab[9442]++
							kh.Write(R.Bytes())
							kh.Write(publicKey)
							kh.Write(message)
							hramDigest := make([]byte, 0, sha512.Size)
							hramDigest = kh.Sum(hramDigest)
							k, err := edwards25519.NewScalar().SetUniformBytes(hramDigest)
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:242
		_go_fuzz_dep_.CoverTab[9454]++
								panic("ed25519: internal error: setting scalar failed")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:243
		// _ = "end of CoverTab[9454]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:244
		_go_fuzz_dep_.CoverTab[9455]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:244
		// _ = "end of CoverTab[9455]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:244
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:244
	// _ = "end of CoverTab[9442]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:244
	_go_fuzz_dep_.CoverTab[9443]++

							S := edwards25519.NewScalar().MultiplyAdd(k, s, r)

							copy(signature[:32], R.Bytes())
							copy(signature[32:], S.Bytes())
//line /usr/local/go/src/crypto/ed25519/ed25519.go:249
	// _ = "end of CoverTab[9443]"
}

// Verify reports whether sig is a valid signature of message by publicKey. It
//line /usr/local/go/src/crypto/ed25519/ed25519.go:252
// will panic if len(publicKey) is not [PublicKeySize].
//line /usr/local/go/src/crypto/ed25519/ed25519.go:254
func Verify(publicKey PublicKey, message, sig []byte) bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:254
	_go_fuzz_dep_.CoverTab[9456]++
							return verify(publicKey, message, sig, domPrefixPure, "")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:255
	// _ = "end of CoverTab[9456]"
}

// VerifyWithOptions reports whether sig is a valid signature of message by
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
// publicKey. A valid signature is indicated by returning a nil error. It will
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
// panic if len(publicKey) is not [PublicKeySize].
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
//
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
// If opts.Hash is [crypto.SHA512], the pre-hashed variant Ed25519ph is used and
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
// message is expected to be a SHA-512 hash, otherwise opts.Hash must be
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
// [crypto.Hash](0) and the message must not be hashed, as Ed25519 performs two
//line /usr/local/go/src/crypto/ed25519/ed25519.go:258
// passes over messages to be signed.
//line /usr/local/go/src/crypto/ed25519/ed25519.go:266
func VerifyWithOptions(publicKey PublicKey, message, sig []byte, opts *Options) error {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:266
	_go_fuzz_dep_.CoverTab[9457]++
							switch {
	case opts.Hash == crypto.SHA512:
//line /usr/local/go/src/crypto/ed25519/ed25519.go:268
		_go_fuzz_dep_.CoverTab[9458]++
								if l := len(message); l != sha512.Size {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:269
			_go_fuzz_dep_.CoverTab[9468]++
									return errors.New("ed25519: bad Ed25519ph message hash length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:270
			// _ = "end of CoverTab[9468]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:271
			_go_fuzz_dep_.CoverTab[9469]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:271
			// _ = "end of CoverTab[9469]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:271
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:271
		// _ = "end of CoverTab[9458]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:271
		_go_fuzz_dep_.CoverTab[9459]++
								if l := len(opts.Context); l > 255 {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:272
			_go_fuzz_dep_.CoverTab[9470]++
									return errors.New("ed25519: bad Ed25519ph context length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:273
			// _ = "end of CoverTab[9470]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:274
			_go_fuzz_dep_.CoverTab[9471]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:274
			// _ = "end of CoverTab[9471]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:274
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:274
		// _ = "end of CoverTab[9459]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:274
		_go_fuzz_dep_.CoverTab[9460]++
								if !verify(publicKey, message, sig, domPrefixPh, opts.Context) {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:275
			_go_fuzz_dep_.CoverTab[9472]++
									return errors.New("ed25519: invalid signature")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:276
			// _ = "end of CoverTab[9472]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:277
			_go_fuzz_dep_.CoverTab[9473]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:277
			// _ = "end of CoverTab[9473]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:277
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:277
		// _ = "end of CoverTab[9460]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:277
		_go_fuzz_dep_.CoverTab[9461]++
								return nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:278
		// _ = "end of CoverTab[9461]"
	case opts.Hash == crypto.Hash(0) && func() bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:279
		_go_fuzz_dep_.CoverTab[9474]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:279
		return opts.Context != ""
//line /usr/local/go/src/crypto/ed25519/ed25519.go:279
		// _ = "end of CoverTab[9474]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:279
	}():
//line /usr/local/go/src/crypto/ed25519/ed25519.go:279
		_go_fuzz_dep_.CoverTab[9462]++
								if l := len(opts.Context); l > 255 {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:280
			_go_fuzz_dep_.CoverTab[9475]++
									return errors.New("ed25519: bad Ed25519ctx context length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:281
			// _ = "end of CoverTab[9475]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:282
			_go_fuzz_dep_.CoverTab[9476]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:282
			// _ = "end of CoverTab[9476]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:282
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:282
		// _ = "end of CoverTab[9462]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:282
		_go_fuzz_dep_.CoverTab[9463]++
								if !verify(publicKey, message, sig, domPrefixCtx, opts.Context) {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:283
			_go_fuzz_dep_.CoverTab[9477]++
									return errors.New("ed25519: invalid signature")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:284
			// _ = "end of CoverTab[9477]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:285
			_go_fuzz_dep_.CoverTab[9478]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:285
			// _ = "end of CoverTab[9478]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:285
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:285
		// _ = "end of CoverTab[9463]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:285
		_go_fuzz_dep_.CoverTab[9464]++
								return nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:286
		// _ = "end of CoverTab[9464]"
	case opts.Hash == crypto.Hash(0):
//line /usr/local/go/src/crypto/ed25519/ed25519.go:287
		_go_fuzz_dep_.CoverTab[9465]++
								if !verify(publicKey, message, sig, domPrefixPure, "") {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:288
			_go_fuzz_dep_.CoverTab[9479]++
									return errors.New("ed25519: invalid signature")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:289
			// _ = "end of CoverTab[9479]"
		} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:290
			_go_fuzz_dep_.CoverTab[9480]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:290
			// _ = "end of CoverTab[9480]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:290
		}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:290
		// _ = "end of CoverTab[9465]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:290
		_go_fuzz_dep_.CoverTab[9466]++
								return nil
//line /usr/local/go/src/crypto/ed25519/ed25519.go:291
		// _ = "end of CoverTab[9466]"
	default:
//line /usr/local/go/src/crypto/ed25519/ed25519.go:292
		_go_fuzz_dep_.CoverTab[9467]++
								return errors.New("ed25519: expected opts.Hash zero (unhashed message, for standard Ed25519) or SHA-512 (for Ed25519ph)")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:293
		// _ = "end of CoverTab[9467]"
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:294
	// _ = "end of CoverTab[9457]"
}

func verify(publicKey PublicKey, message, sig []byte, domPrefix, context string) bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:297
	_go_fuzz_dep_.CoverTab[9481]++
							if l := len(publicKey); l != PublicKeySize {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:298
		_go_fuzz_dep_.CoverTab[9488]++
								panic("ed25519: bad public key length: " + strconv.Itoa(l))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:299
		// _ = "end of CoverTab[9488]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:300
		_go_fuzz_dep_.CoverTab[9489]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:300
		// _ = "end of CoverTab[9489]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:300
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:300
	// _ = "end of CoverTab[9481]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:300
	_go_fuzz_dep_.CoverTab[9482]++

							if len(sig) != SignatureSize || func() bool {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:302
		_go_fuzz_dep_.CoverTab[9490]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:302
		return sig[63]&224 != 0
//line /usr/local/go/src/crypto/ed25519/ed25519.go:302
		// _ = "end of CoverTab[9490]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:302
	}() {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:302
		_go_fuzz_dep_.CoverTab[9491]++
								return false
//line /usr/local/go/src/crypto/ed25519/ed25519.go:303
		// _ = "end of CoverTab[9491]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:304
		_go_fuzz_dep_.CoverTab[9492]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:304
		// _ = "end of CoverTab[9492]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:304
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:304
	// _ = "end of CoverTab[9482]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:304
	_go_fuzz_dep_.CoverTab[9483]++

							A, err := (&edwards25519.Point{}).SetBytes(publicKey)
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:307
		_go_fuzz_dep_.CoverTab[9493]++
								return false
//line /usr/local/go/src/crypto/ed25519/ed25519.go:308
		// _ = "end of CoverTab[9493]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:309
		_go_fuzz_dep_.CoverTab[9494]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:309
		// _ = "end of CoverTab[9494]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:309
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:309
	// _ = "end of CoverTab[9483]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:309
	_go_fuzz_dep_.CoverTab[9484]++

							kh := sha512.New()
							if domPrefix != domPrefixPure {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:312
		_go_fuzz_dep_.CoverTab[9495]++
								kh.Write([]byte(domPrefix))
								kh.Write([]byte{byte(len(context))})
								kh.Write([]byte(context))
//line /usr/local/go/src/crypto/ed25519/ed25519.go:315
		// _ = "end of CoverTab[9495]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:316
		_go_fuzz_dep_.CoverTab[9496]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:316
		// _ = "end of CoverTab[9496]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:316
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:316
	// _ = "end of CoverTab[9484]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:316
	_go_fuzz_dep_.CoverTab[9485]++
							kh.Write(sig[:32])
							kh.Write(publicKey)
							kh.Write(message)
							hramDigest := make([]byte, 0, sha512.Size)
							hramDigest = kh.Sum(hramDigest)
							k, err := edwards25519.NewScalar().SetUniformBytes(hramDigest)
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:323
		_go_fuzz_dep_.CoverTab[9497]++
								panic("ed25519: internal error: setting scalar failed")
//line /usr/local/go/src/crypto/ed25519/ed25519.go:324
		// _ = "end of CoverTab[9497]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:325
		_go_fuzz_dep_.CoverTab[9498]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:325
		// _ = "end of CoverTab[9498]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:325
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:325
	// _ = "end of CoverTab[9485]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:325
	_go_fuzz_dep_.CoverTab[9486]++

							S, err := edwards25519.NewScalar().SetCanonicalBytes(sig[32:])
							if err != nil {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:328
		_go_fuzz_dep_.CoverTab[9499]++
								return false
//line /usr/local/go/src/crypto/ed25519/ed25519.go:329
		// _ = "end of CoverTab[9499]"
	} else {
//line /usr/local/go/src/crypto/ed25519/ed25519.go:330
		_go_fuzz_dep_.CoverTab[9500]++
//line /usr/local/go/src/crypto/ed25519/ed25519.go:330
		// _ = "end of CoverTab[9500]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:330
	}
//line /usr/local/go/src/crypto/ed25519/ed25519.go:330
	// _ = "end of CoverTab[9486]"
//line /usr/local/go/src/crypto/ed25519/ed25519.go:330
	_go_fuzz_dep_.CoverTab[9487]++

//line /usr/local/go/src/crypto/ed25519/ed25519.go:333
	minusA := (&edwards25519.Point{}).Negate(A)
							R := (&edwards25519.Point{}).VarTimeDoubleScalarBaseMult(k, minusA, S)

							return bytes.Equal(sig[:32], R.Bytes())
//line /usr/local/go/src/crypto/ed25519/ed25519.go:336
	// _ = "end of CoverTab[9487]"
}

//line /usr/local/go/src/crypto/ed25519/ed25519.go:337
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/ed25519/ed25519.go:337
var _ = _go_fuzz_dep_.CoverTab
