// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// Package poly1305 implements Poly1305 one-time message authentication code as
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// specified in https://cr.yp.to/mac/poly1305-20050329.pdf.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// Poly1305 is a fast, one-time authentication function. It is infeasible for an
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// attacker to generate an authenticator for a message without the key. However, a
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// key must only be used for a single message. Authenticating two different
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// messages with the same key allows an attacker to forge authenticators for other
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// messages with the same key.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// Poly1305 was originally coupled with AES in order to make Poly1305-AES. AES was
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// used with a fixed key in order to generate one-time keys from an nonce.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// However, in this package AES isn't used and the one-time key is specified
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:5
// directly.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
package poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:18
)

import "crypto/subtle"

// TagSize is the size, in bytes, of a poly1305 authenticator.
const TagSize = 16

// Sum generates an authenticator for msg using a one-time key and puts the
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:25
// 16-byte result into out. Authenticating two different messages with the same
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:25
// key allows an attacker to forge messages at will.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:28
func Sum(out *[16]byte, m []byte, key *[32]byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:28
	_go_fuzz_dep_.CoverTab[20739]++
											h := New(key)
											h.Write(m)
											h.Sum(out[:0])
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:31
	// _ = "end of CoverTab[20739]"
}

// Verify returns true if mac is a valid authenticator for m with the given key.
func Verify(mac *[16]byte, m []byte, key *[32]byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:35
	_go_fuzz_dep_.CoverTab[20740]++
											var tmp [16]byte
											Sum(&tmp, m, key)
											return subtle.ConstantTimeCompare(tmp[:], mac[:]) == 1
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:38
	// _ = "end of CoverTab[20740]"
}

// New returns a new MAC computing an authentication
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// tag of all data written to it with the given key.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// This allows writing the message progressively instead
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// of passing it as a single slice. Common users should use
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// the Sum function instead.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// The key must be unique for each message, as authenticating
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// two different messages with the same key allows an attacker
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:41
// to forge messages at will.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:50
func New(key *[32]byte) *MAC {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:50
	_go_fuzz_dep_.CoverTab[20741]++
											m := &MAC{}
											initialize(key, &m.macState)
											return m
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:53
	// _ = "end of CoverTab[20741]"
}

// MAC is an io.Writer computing an authentication tag
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:56
// of the data written to it.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:56
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:56
// MAC cannot be used like common hash.Hash implementations,
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:56
// because using a poly1305 key twice breaks its security.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:56
// Therefore writing data to a running MAC after calling
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:56
// Sum or Verify causes it to panic.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:63
type MAC struct {
	mac	// platform-dependent implementation

	finalized	bool
}

// Size returns the number of bytes Sum will return.
func (h *MAC) Size() int {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:70
	_go_fuzz_dep_.CoverTab[20742]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:70
	return TagSize
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:70
	// _ = "end of CoverTab[20742]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:70
}

// Write adds more data to the running message authentication code.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:72
// It never returns an error.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:72
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:72
// It must not be called after the first call of Sum or Verify.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:76
func (h *MAC) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:76
	_go_fuzz_dep_.CoverTab[20743]++
											if h.finalized {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:77
		_go_fuzz_dep_.CoverTab[20745]++
												panic("poly1305: write to MAC after Sum or Verify")
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:78
		// _ = "end of CoverTab[20745]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:79
		_go_fuzz_dep_.CoverTab[20746]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:79
		// _ = "end of CoverTab[20746]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:79
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:79
	// _ = "end of CoverTab[20743]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:79
	_go_fuzz_dep_.CoverTab[20744]++
											return h.mac.Write(p)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:80
	// _ = "end of CoverTab[20744]"
}

// Sum computes the authenticator of all data written to the
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:83
// message authentication code.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:85
func (h *MAC) Sum(b []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:85
	_go_fuzz_dep_.CoverTab[20747]++
											var mac [TagSize]byte
											h.mac.Sum(&mac)
											h.finalized = true
											return append(b, mac[:]...)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:89
	// _ = "end of CoverTab[20747]"
}

// Verify returns whether the authenticator of all data written to
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:92
// the message authentication code matches the expected value.
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:94
func (h *MAC) Verify(expected []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:94
	_go_fuzz_dep_.CoverTab[20748]++
											var mac [TagSize]byte
											h.mac.Sum(&mac)
											h.finalized = true
											return subtle.ConstantTimeCompare(expected, mac[:]) == 1
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:98
	// _ = "end of CoverTab[20748]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/poly1305.go:99
var _ = _go_fuzz_dep_.CoverTab
