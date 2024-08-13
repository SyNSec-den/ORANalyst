// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.13
// +build go1.13

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// Package ed25519 implements the Ed25519 signature algorithm. See
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// https://ed25519.cr.yp.to/.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
//
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// These functions are also compatible with the “Ed25519” function defined in
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// RFC 8032. However, unlike RFC 8032's formulation, this package's private key
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// representation includes a public key suffix to make multiple signing
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// operations with the same key more efficient. This package refers to the RFC
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// 8032 private key as the “seed”.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
//
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// Beginning with Go 1.13, the functionality of this package was moved to the
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// standard library as crypto/ed25519. This package only acts as a compatibility
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:8
// wrapper.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
package ed25519

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:20
)

import (
	"crypto/ed25519"
	"io"
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
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:38
//
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:38
// This type is an alias for crypto/ed25519's PublicKey type.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:38
// See the crypto/ed25519 package for the methods on this type.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:42
type PublicKey = ed25519.PublicKey

// PrivateKey is the type of Ed25519 private keys. It implements crypto.Signer.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:44
//
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:44
// This type is an alias for crypto/ed25519's PrivateKey type.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:44
// See the crypto/ed25519 package for the methods on this type.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:48
type PrivateKey = ed25519.PrivateKey

// GenerateKey generates a public/private key pair using entropy from rand.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:50
// If rand is nil, crypto/rand.Reader will be used.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:52
func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:52
	_go_fuzz_dep_.CoverTab[187340]++
															return ed25519.GenerateKey(rand)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:53
	// _ = "end of CoverTab[187340]"
}

// NewKeyFromSeed calculates a private key from a seed. It will panic if
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:56
// len(seed) is not SeedSize. This function is provided for interoperability
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:56
// with RFC 8032. RFC 8032's private keys correspond to seeds in this
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:56
// package.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:60
func NewKeyFromSeed(seed []byte) PrivateKey {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:60
	_go_fuzz_dep_.CoverTab[187341]++
															return ed25519.NewKeyFromSeed(seed)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:61
	// _ = "end of CoverTab[187341]"
}

// Sign signs the message with privateKey and returns a signature. It will
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:64
// panic if len(privateKey) is not PrivateKeySize.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:66
func Sign(privateKey PrivateKey, message []byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:66
	_go_fuzz_dep_.CoverTab[187342]++
															return ed25519.Sign(privateKey, message)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:67
	// _ = "end of CoverTab[187342]"
}

// Verify reports whether sig is a valid signature of message by publicKey. It
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:70
// will panic if len(publicKey) is not PublicKeySize.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:72
func Verify(publicKey PublicKey, message, sig []byte) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:72
	_go_fuzz_dep_.CoverTab[187343]++
															return ed25519.Verify(publicKey, message, sig)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:73
	// _ = "end of CoverTab[187343]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/ed25519/ed25519_go113.go:74
var _ = _go_fuzz_dep_.CoverTab
