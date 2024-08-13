// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
package chacha20poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:5
)

import (
	"crypto/cipher"
	"errors"

	"golang.org/x/crypto/chacha20"
)

type xchacha20poly1305 struct {
	key [KeySize]byte
}

// NewX returns a XChaCha20-Poly1305 AEAD that uses the given 256-bit key.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:18
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:18
// XChaCha20-Poly1305 is a ChaCha20-Poly1305 variant that takes a longer nonce,
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:18
// suitable to be generated randomly without risk of collisions. It should be
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:18
// preferred when nonce uniqueness cannot be trivially ensured, or whenever
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:18
// nonces are randomly generated.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:24
func NewX(key []byte) (cipher.AEAD, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:24
	_go_fuzz_dep_.CoverTab[20971]++
												if len(key) != KeySize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:25
		_go_fuzz_dep_.CoverTab[20973]++
													return nil, errors.New("chacha20poly1305: bad key length")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:26
		// _ = "end of CoverTab[20973]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:27
		_go_fuzz_dep_.CoverTab[20974]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:27
		// _ = "end of CoverTab[20974]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:27
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:27
	// _ = "end of CoverTab[20971]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:27
	_go_fuzz_dep_.CoverTab[20972]++
												ret := new(xchacha20poly1305)
												copy(ret.key[:], key)
												return ret, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:30
	// _ = "end of CoverTab[20972]"
}

func (*xchacha20poly1305) NonceSize() int {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:33
	_go_fuzz_dep_.CoverTab[20975]++
												return NonceSizeX
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:34
	// _ = "end of CoverTab[20975]"
}

func (*xchacha20poly1305) Overhead() int {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:37
	_go_fuzz_dep_.CoverTab[20976]++
												return Overhead
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:38
	// _ = "end of CoverTab[20976]"
}

func (x *xchacha20poly1305) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:41
	_go_fuzz_dep_.CoverTab[20977]++
												if len(nonce) != NonceSizeX {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:42
		_go_fuzz_dep_.CoverTab[20980]++
													panic("chacha20poly1305: bad nonce length passed to Seal")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:43
		// _ = "end of CoverTab[20980]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:44
		_go_fuzz_dep_.CoverTab[20981]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:44
		// _ = "end of CoverTab[20981]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:44
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:44
	// _ = "end of CoverTab[20977]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:44
	_go_fuzz_dep_.CoverTab[20978]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:51
	if uint64(len(plaintext)) > (1<<38)-64 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:51
		_go_fuzz_dep_.CoverTab[20982]++
													panic("chacha20poly1305: plaintext too large")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:52
		// _ = "end of CoverTab[20982]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:53
		_go_fuzz_dep_.CoverTab[20983]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:53
		// _ = "end of CoverTab[20983]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:53
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:53
	// _ = "end of CoverTab[20978]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:53
	_go_fuzz_dep_.CoverTab[20979]++

												c := new(chacha20poly1305)
												hKey, _ := chacha20.HChaCha20(x.key[:], nonce[0:16])
												copy(c.key[:], hKey)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:60
	cNonce := make([]byte, NonceSize)
												copy(cNonce[4:12], nonce[16:24])

												return c.seal(dst, cNonce[:], plaintext, additionalData)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:63
	// _ = "end of CoverTab[20979]"
}

func (x *xchacha20poly1305) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:66
	_go_fuzz_dep_.CoverTab[20984]++
												if len(nonce) != NonceSizeX {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:67
		_go_fuzz_dep_.CoverTab[20988]++
													panic("chacha20poly1305: bad nonce length passed to Open")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:68
		// _ = "end of CoverTab[20988]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:69
		_go_fuzz_dep_.CoverTab[20989]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:69
		// _ = "end of CoverTab[20989]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:69
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:69
	// _ = "end of CoverTab[20984]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:69
	_go_fuzz_dep_.CoverTab[20985]++
												if len(ciphertext) < 16 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:70
		_go_fuzz_dep_.CoverTab[20990]++
													return nil, errOpen
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:71
		// _ = "end of CoverTab[20990]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:72
		_go_fuzz_dep_.CoverTab[20991]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:72
		// _ = "end of CoverTab[20991]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:72
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:72
	// _ = "end of CoverTab[20985]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:72
	_go_fuzz_dep_.CoverTab[20986]++
												if uint64(len(ciphertext)) > (1<<38)-48 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:73
		_go_fuzz_dep_.CoverTab[20992]++
													panic("chacha20poly1305: ciphertext too large")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:74
		// _ = "end of CoverTab[20992]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:75
		_go_fuzz_dep_.CoverTab[20993]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:75
		// _ = "end of CoverTab[20993]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:75
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:75
	// _ = "end of CoverTab[20986]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:75
	_go_fuzz_dep_.CoverTab[20987]++

												c := new(chacha20poly1305)
												hKey, _ := chacha20.HChaCha20(x.key[:], nonce[0:16])
												copy(c.key[:], hKey)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:82
	cNonce := make([]byte, NonceSize)
												copy(cNonce[4:12], nonce[16:24])

												return c.open(dst, cNonce[:], ciphertext, additionalData)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:85
	// _ = "end of CoverTab[20987]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/xchacha20poly1305.go:86
var _ = _go_fuzz_dep_.CoverTab
