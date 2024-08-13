// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:5
// Package chacha20 implements the ChaCha20 and XChaCha20 encryption algorithms
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:5
// as specified in RFC 8439 and draft-irtf-cfrg-xchacha-01.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
package chacha20

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:7
)

import (
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"math/bits"

	"golang.org/x/crypto/internal/alias"
)

const (
	// KeySize is the size of the key used by this cipher, in bytes.
	KeySize	= 32

	// NonceSize is the size of the nonce used with the standard variant of this
	// cipher, in bytes.
	//
	// Note that this is too short to be safely generated at random if the same
	// key is reused more than 2³² times.
	NonceSize	= 12

	// NonceSizeX is the size of the nonce used with the XChaCha20 variant of
	// this cipher, in bytes.
	NonceSizeX	= 24
)

// Cipher is a stateful instance of ChaCha20 or XChaCha20 using a particular key
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:34
// and nonce. A *Cipher implements the cipher.Stream interface.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:36
type Cipher struct {
	// The ChaCha20 state is 16 words: 4 constant, 8 of key, 1 of counter
	// (incremented after each block), and 3 of nonce.
	key	[8]uint32
	counter	uint32
	nonce	[3]uint32

	// The last len bytes of buf are leftover key stream bytes from the previous
	// XORKeyStream invocation. The size of buf depends on how many blocks are
	// computed at a time by xorKeyStreamBlocks.
	buf	[bufSize]byte
	len	int

	// overflow is set when the counter overflowed, no more blocks can be
	// generated, and the next XORKeyStream call should panic.
	overflow	bool

	// The counter-independent results of the first round are cached after they
	// are computed the first time.
	precompDone		bool
	p1, p5, p9, p13		uint32
	p2, p6, p10, p14	uint32
	p3, p7, p11, p15	uint32
}

var _ cipher.Stream = (*Cipher)(nil)

// NewUnauthenticatedCipher creates a new ChaCha20 stream cipher with the given
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// 32 bytes key and a 12 or 24 bytes nonce. If a nonce of 24 bytes is provided,
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// the XChaCha20 construction will be used. It returns an error if key or nonce
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// have any other length.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// Note that ChaCha20, like all stream ciphers, is not authenticated and allows
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// attackers to silently tamper with the plaintext. For this reason, it is more
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// appropriate as a building block than as a standalone encryption mechanism.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:63
// Instead, consider using package golang.org/x/crypto/chacha20poly1305.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:72
func NewUnauthenticatedCipher(key, nonce []byte) (*Cipher, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:72
	_go_fuzz_dep_.CoverTab[20657]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:76
	c := &Cipher{}
											return newUnauthenticatedCipher(c, key, nonce)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:77
	// _ = "end of CoverTab[20657]"
}

func newUnauthenticatedCipher(c *Cipher, key, nonce []byte) (*Cipher, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:80
	_go_fuzz_dep_.CoverTab[20658]++
											if len(key) != KeySize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:81
		_go_fuzz_dep_.CoverTab[20661]++
												return nil, errors.New("chacha20: wrong key size")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:82
		// _ = "end of CoverTab[20661]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:83
		_go_fuzz_dep_.CoverTab[20662]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:83
		// _ = "end of CoverTab[20662]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:83
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:83
	// _ = "end of CoverTab[20658]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:83
	_go_fuzz_dep_.CoverTab[20659]++
											if len(nonce) == NonceSizeX {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:84
		_go_fuzz_dep_.CoverTab[20663]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:88
		key, _ = HChaCha20(key, nonce[0:16])
												cNonce := make([]byte, NonceSize)
												copy(cNonce[4:12], nonce[16:24])
												nonce = cNonce
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:91
		// _ = "end of CoverTab[20663]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:92
		_go_fuzz_dep_.CoverTab[20664]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:92
		if len(nonce) != NonceSize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:92
			_go_fuzz_dep_.CoverTab[20665]++
													return nil, errors.New("chacha20: wrong nonce size")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:93
			// _ = "end of CoverTab[20665]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
			_go_fuzz_dep_.CoverTab[20666]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
			// _ = "end of CoverTab[20666]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
		// _ = "end of CoverTab[20664]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
	// _ = "end of CoverTab[20659]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:94
	_go_fuzz_dep_.CoverTab[20660]++

											key, nonce = key[:KeySize], nonce[:NonceSize]
											c.key = [8]uint32{
		binary.LittleEndian.Uint32(key[0:4]),
		binary.LittleEndian.Uint32(key[4:8]),
		binary.LittleEndian.Uint32(key[8:12]),
		binary.LittleEndian.Uint32(key[12:16]),
		binary.LittleEndian.Uint32(key[16:20]),
		binary.LittleEndian.Uint32(key[20:24]),
		binary.LittleEndian.Uint32(key[24:28]),
		binary.LittleEndian.Uint32(key[28:32]),
	}
	c.nonce = [3]uint32{
		binary.LittleEndian.Uint32(nonce[0:4]),
		binary.LittleEndian.Uint32(nonce[4:8]),
		binary.LittleEndian.Uint32(nonce[8:12]),
	}
											return c, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:112
	// _ = "end of CoverTab[20660]"
}

// The constant first 4 words of the ChaCha20 state.
const (
	j0	uint32	= 0x61707865	// expa
	j1	uint32	= 0x3320646e	// nd 3
	j2	uint32	= 0x79622d32	// 2-by
	j3	uint32	= 0x6b206574	// te k
)

const blockSize = 64

// quarterRound is the core of ChaCha20. It shuffles the bits of 4 state words.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:125
// It's executed 4 times for each of the 20 ChaCha20 rounds, operating on all 16
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:125
// words each round, in columnar or diagonal groups of 4 at a time.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:128
func quarterRound(a, b, c, d uint32) (uint32, uint32, uint32, uint32) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:128
	_go_fuzz_dep_.CoverTab[20667]++
											a += b
											d ^= a
											d = bits.RotateLeft32(d, 16)
											c += d
											b ^= c
											b = bits.RotateLeft32(b, 12)
											a += b
											d ^= a
											d = bits.RotateLeft32(d, 8)
											c += d
											b ^= c
											b = bits.RotateLeft32(b, 7)
											return a, b, c, d
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:141
	// _ = "end of CoverTab[20667]"
}

// SetCounter sets the Cipher counter. The next invocation of XORKeyStream will
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
// behave as if (64 * counter) bytes had been encrypted so far.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
// To prevent accidental counter reuse, SetCounter panics if counter is less
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
// than the current value.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
// Note that the execution time of XORKeyStream is not independent of the
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:144
// counter value.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:152
func (s *Cipher) SetCounter(counter uint32) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:152
	_go_fuzz_dep_.CoverTab[20668]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:157
	outputCounter := s.counter - uint32(s.len)/blockSize
	if s.overflow || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:158
		_go_fuzz_dep_.CoverTab[20670]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:158
		return counter < outputCounter
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:158
		// _ = "end of CoverTab[20670]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:158
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:158
		_go_fuzz_dep_.CoverTab[20671]++
												panic("chacha20: SetCounter attempted to rollback counter")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:159
		// _ = "end of CoverTab[20671]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:160
		_go_fuzz_dep_.CoverTab[20672]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:160
		// _ = "end of CoverTab[20672]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:160
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:160
	// _ = "end of CoverTab[20668]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:160
	_go_fuzz_dep_.CoverTab[20669]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:166
	if counter < s.counter {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:166
		_go_fuzz_dep_.CoverTab[20673]++
												s.len = int(s.counter-counter) * blockSize
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:167
		// _ = "end of CoverTab[20673]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:168
		_go_fuzz_dep_.CoverTab[20674]++
												s.counter = counter
												s.len = 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:170
		// _ = "end of CoverTab[20674]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:171
	// _ = "end of CoverTab[20669]"
}

// XORKeyStream XORs each byte in the given slice with a byte from the
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// cipher's key stream. Dst and src must overlap entirely or not at all.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// If len(dst) < len(src), XORKeyStream will panic. It is acceptable
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// to pass a dst bigger than src, and in that case, XORKeyStream will
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// only update dst[:len(src)] and will not touch the rest of dst.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// Multiple calls to XORKeyStream behave as if the concatenation of
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// the src buffers was passed in a single run. That is, Cipher
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:174
// maintains state and does not reset at each XORKeyStream call.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:184
func (s *Cipher) XORKeyStream(dst, src []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:184
	_go_fuzz_dep_.CoverTab[20675]++
											if len(src) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:185
		_go_fuzz_dep_.CoverTab[20684]++
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:186
		// _ = "end of CoverTab[20684]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:187
		_go_fuzz_dep_.CoverTab[20685]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:187
		// _ = "end of CoverTab[20685]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:187
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:187
	// _ = "end of CoverTab[20675]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:187
	_go_fuzz_dep_.CoverTab[20676]++
											if len(dst) < len(src) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:188
		_go_fuzz_dep_.CoverTab[20686]++
												panic("chacha20: output smaller than input")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:189
		// _ = "end of CoverTab[20686]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:190
		_go_fuzz_dep_.CoverTab[20687]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:190
		// _ = "end of CoverTab[20687]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:190
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:190
	// _ = "end of CoverTab[20676]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:190
	_go_fuzz_dep_.CoverTab[20677]++
											dst = dst[:len(src)]
											if alias.InexactOverlap(dst, src) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:192
		_go_fuzz_dep_.CoverTab[20688]++
												panic("chacha20: invalid buffer overlap")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:193
		// _ = "end of CoverTab[20688]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:194
		_go_fuzz_dep_.CoverTab[20689]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:194
		// _ = "end of CoverTab[20689]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:194
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:194
	// _ = "end of CoverTab[20677]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:194
	_go_fuzz_dep_.CoverTab[20678]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:197
	if s.len != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:197
		_go_fuzz_dep_.CoverTab[20690]++
												keyStream := s.buf[bufSize-s.len:]
												if len(src) < len(keyStream) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:199
			_go_fuzz_dep_.CoverTab[20693]++
													keyStream = keyStream[:len(src)]
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:200
			// _ = "end of CoverTab[20693]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:201
			_go_fuzz_dep_.CoverTab[20694]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:201
			// _ = "end of CoverTab[20694]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:201
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:201
		// _ = "end of CoverTab[20690]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:201
		_go_fuzz_dep_.CoverTab[20691]++
												_ = src[len(keyStream)-1]
												for i, b := range keyStream {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:203
			_go_fuzz_dep_.CoverTab[20695]++
													dst[i] = src[i] ^ b
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:204
			// _ = "end of CoverTab[20695]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:205
		// _ = "end of CoverTab[20691]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:205
		_go_fuzz_dep_.CoverTab[20692]++
												s.len -= len(keyStream)
												dst, src = dst[len(keyStream):], src[len(keyStream):]
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:207
		// _ = "end of CoverTab[20692]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:208
		_go_fuzz_dep_.CoverTab[20696]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:208
		// _ = "end of CoverTab[20696]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:208
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:208
	// _ = "end of CoverTab[20678]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:208
	_go_fuzz_dep_.CoverTab[20679]++
											if len(src) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:209
		_go_fuzz_dep_.CoverTab[20697]++
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:210
		// _ = "end of CoverTab[20697]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:211
		_go_fuzz_dep_.CoverTab[20698]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:211
		// _ = "end of CoverTab[20698]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:211
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:211
	// _ = "end of CoverTab[20679]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:211
	_go_fuzz_dep_.CoverTab[20680]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:216
	numBlocks := (uint64(len(src)) + blockSize - 1) / blockSize
	if s.overflow || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:217
		_go_fuzz_dep_.CoverTab[20699]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:217
		return uint64(s.counter)+numBlocks > 1<<32
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:217
		// _ = "end of CoverTab[20699]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:217
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:217
		_go_fuzz_dep_.CoverTab[20700]++
												panic("chacha20: counter overflow")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:218
		// _ = "end of CoverTab[20700]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:219
		_go_fuzz_dep_.CoverTab[20701]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:219
		if uint64(s.counter)+numBlocks == 1<<32 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:219
			_go_fuzz_dep_.CoverTab[20702]++
													s.overflow = true
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:220
			// _ = "end of CoverTab[20702]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
			_go_fuzz_dep_.CoverTab[20703]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
			// _ = "end of CoverTab[20703]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
		// _ = "end of CoverTab[20701]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
	// _ = "end of CoverTab[20680]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:221
	_go_fuzz_dep_.CoverTab[20681]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:227
	full := len(src) - len(src)%bufSize
	if full > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:228
		_go_fuzz_dep_.CoverTab[20704]++
												s.xorKeyStreamBlocks(dst[:full], src[:full])
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:229
		// _ = "end of CoverTab[20704]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:230
		_go_fuzz_dep_.CoverTab[20705]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:230
		// _ = "end of CoverTab[20705]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:230
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:230
	// _ = "end of CoverTab[20681]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:230
	_go_fuzz_dep_.CoverTab[20682]++
											dst, src = dst[full:], src[full:]

	// If using a multi-block xorKeyStreamBlocks would overflow, use the generic
	// one that does one block at a time.
	const blocksPerBuf = bufSize / blockSize
	if uint64(s.counter)+blocksPerBuf > 1<<32 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:236
		_go_fuzz_dep_.CoverTab[20706]++
												s.buf = [bufSize]byte{}
												numBlocks := (len(src) + blockSize - 1) / blockSize
												buf := s.buf[bufSize-numBlocks*blockSize:]
												copy(buf, src)
												s.xorKeyStreamBlocksGeneric(buf, buf)
												s.len = len(buf) - copy(dst, buf)
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:243
		// _ = "end of CoverTab[20706]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:244
		_go_fuzz_dep_.CoverTab[20707]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:244
		// _ = "end of CoverTab[20707]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:244
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:244
	// _ = "end of CoverTab[20682]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:244
	_go_fuzz_dep_.CoverTab[20683]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:248
	if len(src) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:248
		_go_fuzz_dep_.CoverTab[20708]++
												s.buf = [bufSize]byte{}
												copy(s.buf[:], src)
												s.xorKeyStreamBlocks(s.buf[:], s.buf[:])
												s.len = bufSize - copy(dst, s.buf[:])
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:252
		// _ = "end of CoverTab[20708]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:253
		_go_fuzz_dep_.CoverTab[20709]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:253
		// _ = "end of CoverTab[20709]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:253
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:253
	// _ = "end of CoverTab[20683]"
}

func (s *Cipher) xorKeyStreamBlocksGeneric(dst, src []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:256
	_go_fuzz_dep_.CoverTab[20710]++
											if len(dst) != len(src) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:257
		_go_fuzz_dep_.CoverTab[20713]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:257
		return len(dst)%blockSize != 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:257
		// _ = "end of CoverTab[20713]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:257
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:257
		_go_fuzz_dep_.CoverTab[20714]++
												panic("chacha20: internal error: wrong dst and/or src length")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:258
		// _ = "end of CoverTab[20714]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:259
		_go_fuzz_dep_.CoverTab[20715]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:259
		// _ = "end of CoverTab[20715]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:259
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:259
	// _ = "end of CoverTab[20710]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:259
	_go_fuzz_dep_.CoverTab[20711]++

	// To generate each block of key stream, the initial cipher state
	// (represented below) is passed through 20 rounds of shuffling,
	// alternatively applying quarterRounds by columns (like 1, 5, 9, 13)
	// or by diagonals (like 1, 6, 11, 12).
	//
	//      0:cccccccc   1:cccccccc   2:cccccccc   3:cccccccc
	//      4:kkkkkkkk   5:kkkkkkkk   6:kkkkkkkk   7:kkkkkkkk
	//      8:kkkkkkkk   9:kkkkkkkk  10:kkkkkkkk  11:kkkkkkkk
	//     12:bbbbbbbb  13:nnnnnnnn  14:nnnnnnnn  15:nnnnnnnn
	//
	//            c=constant k=key b=blockcount n=nonce
	var (
		c0, c1, c2, c3		= j0, j1, j2, j3
		c4, c5, c6, c7		= s.key[0], s.key[1], s.key[2], s.key[3]
		c8, c9, c10, c11	= s.key[4], s.key[5], s.key[6], s.key[7]
		_, c13, c14, c15	= s.counter, s.nonce[0], s.nonce[1], s.nonce[2]
	)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:282
	if !s.precompDone {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:282
		_go_fuzz_dep_.CoverTab[20716]++
												s.p1, s.p5, s.p9, s.p13 = quarterRound(c1, c5, c9, c13)
												s.p2, s.p6, s.p10, s.p14 = quarterRound(c2, c6, c10, c14)
												s.p3, s.p7, s.p11, s.p15 = quarterRound(c3, c7, c11, c15)
												s.precompDone = true
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:286
		// _ = "end of CoverTab[20716]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:287
		_go_fuzz_dep_.CoverTab[20717]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:287
		// _ = "end of CoverTab[20717]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:287
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:287
	// _ = "end of CoverTab[20711]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:287
	_go_fuzz_dep_.CoverTab[20712]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:291
	for len(src) >= 64 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:291
		_go_fuzz_dep_.CoverTab[20718]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:291
		return len(dst) >= 64
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:291
		// _ = "end of CoverTab[20718]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:291
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:291
		_go_fuzz_dep_.CoverTab[20719]++

												fcr0, fcr4, fcr8, fcr12 := quarterRound(c0, c4, c8, s.counter)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:296
		x0, x5, x10, x15 := quarterRound(fcr0, s.p5, s.p10, s.p15)
												x1, x6, x11, x12 := quarterRound(s.p1, s.p6, s.p11, fcr12)
												x2, x7, x8, x13 := quarterRound(s.p2, s.p7, fcr8, s.p13)
												x3, x4, x9, x14 := quarterRound(s.p3, fcr4, s.p9, s.p14)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:302
		for i := 0; i < 9; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:302
			_go_fuzz_dep_.CoverTab[20721]++

													x0, x4, x8, x12 = quarterRound(x0, x4, x8, x12)
													x1, x5, x9, x13 = quarterRound(x1, x5, x9, x13)
													x2, x6, x10, x14 = quarterRound(x2, x6, x10, x14)
													x3, x7, x11, x15 = quarterRound(x3, x7, x11, x15)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:310
			x0, x5, x10, x15 = quarterRound(x0, x5, x10, x15)
													x1, x6, x11, x12 = quarterRound(x1, x6, x11, x12)
													x2, x7, x8, x13 = quarterRound(x2, x7, x8, x13)
													x3, x4, x9, x14 = quarterRound(x3, x4, x9, x14)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:313
			// _ = "end of CoverTab[20721]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:314
		// _ = "end of CoverTab[20719]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:314
		_go_fuzz_dep_.CoverTab[20720]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:318
		addXor(dst[0:4], src[0:4], x0, c0)
												addXor(dst[4:8], src[4:8], x1, c1)
												addXor(dst[8:12], src[8:12], x2, c2)
												addXor(dst[12:16], src[12:16], x3, c3)
												addXor(dst[16:20], src[16:20], x4, c4)
												addXor(dst[20:24], src[20:24], x5, c5)
												addXor(dst[24:28], src[24:28], x6, c6)
												addXor(dst[28:32], src[28:32], x7, c7)
												addXor(dst[32:36], src[32:36], x8, c8)
												addXor(dst[36:40], src[36:40], x9, c9)
												addXor(dst[40:44], src[40:44], x10, c10)
												addXor(dst[44:48], src[44:48], x11, c11)
												addXor(dst[48:52], src[48:52], x12, s.counter)
												addXor(dst[52:56], src[52:56], x13, c13)
												addXor(dst[56:60], src[56:60], x14, c14)
												addXor(dst[60:64], src[60:64], x15, c15)

												s.counter += 1

												src, dst = src[blockSize:], dst[blockSize:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:337
		// _ = "end of CoverTab[20720]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:338
	// _ = "end of CoverTab[20712]"
}

// HChaCha20 uses the ChaCha20 core to generate a derived key from a 32 bytes
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:341
// key and a 16 bytes nonce. It returns an error if key or nonce have any other
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:341
// length. It is used as part of the XChaCha20 construction.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:344
func HChaCha20(key, nonce []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:344
	_go_fuzz_dep_.CoverTab[20722]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:348
	out := make([]byte, 32)
											return hChaCha20(out, key, nonce)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:349
	// _ = "end of CoverTab[20722]"
}

func hChaCha20(out, key, nonce []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:352
	_go_fuzz_dep_.CoverTab[20723]++
											if len(key) != KeySize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:353
		_go_fuzz_dep_.CoverTab[20727]++
												return nil, errors.New("chacha20: wrong HChaCha20 key size")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:354
		// _ = "end of CoverTab[20727]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:355
		_go_fuzz_dep_.CoverTab[20728]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:355
		// _ = "end of CoverTab[20728]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:355
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:355
	// _ = "end of CoverTab[20723]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:355
	_go_fuzz_dep_.CoverTab[20724]++
											if len(nonce) != 16 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:356
		_go_fuzz_dep_.CoverTab[20729]++
												return nil, errors.New("chacha20: wrong HChaCha20 nonce size")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:357
		// _ = "end of CoverTab[20729]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:358
		_go_fuzz_dep_.CoverTab[20730]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:358
		// _ = "end of CoverTab[20730]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:358
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:358
	// _ = "end of CoverTab[20724]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:358
	_go_fuzz_dep_.CoverTab[20725]++

											x0, x1, x2, x3 := j0, j1, j2, j3
											x4 := binary.LittleEndian.Uint32(key[0:4])
											x5 := binary.LittleEndian.Uint32(key[4:8])
											x6 := binary.LittleEndian.Uint32(key[8:12])
											x7 := binary.LittleEndian.Uint32(key[12:16])
											x8 := binary.LittleEndian.Uint32(key[16:20])
											x9 := binary.LittleEndian.Uint32(key[20:24])
											x10 := binary.LittleEndian.Uint32(key[24:28])
											x11 := binary.LittleEndian.Uint32(key[28:32])
											x12 := binary.LittleEndian.Uint32(nonce[0:4])
											x13 := binary.LittleEndian.Uint32(nonce[4:8])
											x14 := binary.LittleEndian.Uint32(nonce[8:12])
											x15 := binary.LittleEndian.Uint32(nonce[12:16])

											for i := 0; i < 10; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:374
		_go_fuzz_dep_.CoverTab[20731]++

												x0, x4, x8, x12 = quarterRound(x0, x4, x8, x12)
												x1, x5, x9, x13 = quarterRound(x1, x5, x9, x13)
												x2, x6, x10, x14 = quarterRound(x2, x6, x10, x14)
												x3, x7, x11, x15 = quarterRound(x3, x7, x11, x15)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:382
		x0, x5, x10, x15 = quarterRound(x0, x5, x10, x15)
												x1, x6, x11, x12 = quarterRound(x1, x6, x11, x12)
												x2, x7, x8, x13 = quarterRound(x2, x7, x8, x13)
												x3, x4, x9, x14 = quarterRound(x3, x4, x9, x14)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:385
		// _ = "end of CoverTab[20731]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:386
	// _ = "end of CoverTab[20725]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:386
	_go_fuzz_dep_.CoverTab[20726]++

											_ = out[31]
											binary.LittleEndian.PutUint32(out[0:4], x0)
											binary.LittleEndian.PutUint32(out[4:8], x1)
											binary.LittleEndian.PutUint32(out[8:12], x2)
											binary.LittleEndian.PutUint32(out[12:16], x3)
											binary.LittleEndian.PutUint32(out[16:20], x12)
											binary.LittleEndian.PutUint32(out[20:24], x13)
											binary.LittleEndian.PutUint32(out[24:28], x14)
											binary.LittleEndian.PutUint32(out[28:32], x15)
											return out, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:397
	// _ = "end of CoverTab[20726]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:398
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_generic.go:398
var _ = _go_fuzz_dep_.CoverTab
