// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:5
// Package chacha20poly1305 implements the ChaCha20-Poly1305 AEAD and its
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:5
// extended nonce variant XChaCha20-Poly1305, as specified in RFC 8439 and
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:5
// draft-irtf-cfrg-xchacha-01.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
package chacha20poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:8
)

import (
	"crypto/cipher"
	"errors"
)

const (
	// KeySize is the size of the key used by this AEAD, in bytes.
	KeySize	= 32

	// NonceSize is the size of the nonce used with the standard variant of this
	// AEAD, in bytes.
	//
	// Note that this is too short to be safely generated at random if the same
	// key is reused more than 2³² times.
	NonceSize	= 12

	// NonceSizeX is the size of the nonce used with the XChaCha20-Poly1305
	// variant of this AEAD, in bytes.
	NonceSizeX	= 24

	// Overhead is the size of the Poly1305 authentication tag, and the
	// difference between a ciphertext length and its plaintext.
	Overhead	= 16
)

type chacha20poly1305 struct {
	key [KeySize]byte
}

// New returns a ChaCha20-Poly1305 AEAD that uses the given 256-bit key.
func New(key []byte) (cipher.AEAD, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:40
	_go_fuzz_dep_.CoverTab[20906]++
												if len(key) != KeySize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:41
		_go_fuzz_dep_.CoverTab[20908]++
													return nil, errors.New("chacha20poly1305: bad key length")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:42
		// _ = "end of CoverTab[20908]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:43
		_go_fuzz_dep_.CoverTab[20909]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:43
		// _ = "end of CoverTab[20909]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:43
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:43
	// _ = "end of CoverTab[20906]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:43
	_go_fuzz_dep_.CoverTab[20907]++
												ret := new(chacha20poly1305)
												copy(ret.key[:], key)
												return ret, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:46
	// _ = "end of CoverTab[20907]"
}

func (c *chacha20poly1305) NonceSize() int {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:49
	_go_fuzz_dep_.CoverTab[20910]++
												return NonceSize
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:50
	// _ = "end of CoverTab[20910]"
}

func (c *chacha20poly1305) Overhead() int {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:53
	_go_fuzz_dep_.CoverTab[20911]++
												return Overhead
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:54
	// _ = "end of CoverTab[20911]"
}

func (c *chacha20poly1305) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:57
	_go_fuzz_dep_.CoverTab[20912]++
												if len(nonce) != NonceSize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:58
		_go_fuzz_dep_.CoverTab[20915]++
													panic("chacha20poly1305: bad nonce length passed to Seal")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:59
		// _ = "end of CoverTab[20915]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:60
		_go_fuzz_dep_.CoverTab[20916]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:60
		// _ = "end of CoverTab[20916]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:60
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:60
	// _ = "end of CoverTab[20912]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:60
	_go_fuzz_dep_.CoverTab[20913]++

												if uint64(len(plaintext)) > (1<<38)-64 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:62
		_go_fuzz_dep_.CoverTab[20917]++
													panic("chacha20poly1305: plaintext too large")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:63
		// _ = "end of CoverTab[20917]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:64
		_go_fuzz_dep_.CoverTab[20918]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:64
		// _ = "end of CoverTab[20918]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:64
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:64
	// _ = "end of CoverTab[20913]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:64
	_go_fuzz_dep_.CoverTab[20914]++

												return c.seal(dst, nonce, plaintext, additionalData)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:66
	// _ = "end of CoverTab[20914]"
}

var errOpen = errors.New("chacha20poly1305: message authentication failed")

func (c *chacha20poly1305) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:71
	_go_fuzz_dep_.CoverTab[20919]++
												if len(nonce) != NonceSize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:72
		_go_fuzz_dep_.CoverTab[20923]++
													panic("chacha20poly1305: bad nonce length passed to Open")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:73
		// _ = "end of CoverTab[20923]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:74
		_go_fuzz_dep_.CoverTab[20924]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:74
		// _ = "end of CoverTab[20924]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:74
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:74
	// _ = "end of CoverTab[20919]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:74
	_go_fuzz_dep_.CoverTab[20920]++
												if len(ciphertext) < 16 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:75
		_go_fuzz_dep_.CoverTab[20925]++
													return nil, errOpen
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:76
		// _ = "end of CoverTab[20925]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:77
		_go_fuzz_dep_.CoverTab[20926]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:77
		// _ = "end of CoverTab[20926]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:77
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:77
	// _ = "end of CoverTab[20920]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:77
	_go_fuzz_dep_.CoverTab[20921]++
												if uint64(len(ciphertext)) > (1<<38)-48 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:78
		_go_fuzz_dep_.CoverTab[20927]++
													panic("chacha20poly1305: ciphertext too large")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:79
		// _ = "end of CoverTab[20927]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:80
		_go_fuzz_dep_.CoverTab[20928]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:80
		// _ = "end of CoverTab[20928]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:80
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:80
	// _ = "end of CoverTab[20921]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:80
	_go_fuzz_dep_.CoverTab[20922]++

												return c.open(dst, nonce, ciphertext, additionalData)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:82
	// _ = "end of CoverTab[20922]"
}

// sliceForAppend takes a slice and a requested number of bytes. It returns a
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:85
// slice with the contents of the given slice followed by that many bytes and a
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:85
// second slice that aliases into it and contains only the extra bytes. If the
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:85
// original slice has sufficient capacity then no allocation is performed.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:89
func sliceForAppend(in []byte, n int) (head, tail []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:89
	_go_fuzz_dep_.CoverTab[20929]++
												if total := len(in) + n; cap(in) >= total {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:90
		_go_fuzz_dep_.CoverTab[20931]++
													head = in[:total]
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:91
		// _ = "end of CoverTab[20931]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:92
		_go_fuzz_dep_.CoverTab[20932]++
													head = make([]byte, total)
													copy(head, in)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:94
		// _ = "end of CoverTab[20932]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:95
	// _ = "end of CoverTab[20929]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:95
	_go_fuzz_dep_.CoverTab[20930]++
												tail = head[len(in):]
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:97
	// _ = "end of CoverTab[20930]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305.go:98
var _ = _go_fuzz_dep_.CoverTab
