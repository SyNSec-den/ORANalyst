// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/cipher/gcm.go:5
package cipher

//line /usr/local/go/src/crypto/cipher/gcm.go:5
import (
//line /usr/local/go/src/crypto/cipher/gcm.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/gcm.go:5
)
//line /usr/local/go/src/crypto/cipher/gcm.go:5
import (
//line /usr/local/go/src/crypto/cipher/gcm.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/gcm.go:5
)

import (
	"crypto/internal/alias"
	"crypto/subtle"
	"encoding/binary"
	"errors"
)

// AEAD is a cipher mode providing authenticated encryption with associated
//line /usr/local/go/src/crypto/cipher/gcm.go:14
// data. For a description of the methodology, see
//line /usr/local/go/src/crypto/cipher/gcm.go:14
// https://en.wikipedia.org/wiki/Authenticated_encryption.
//line /usr/local/go/src/crypto/cipher/gcm.go:17
type AEAD interface {
	// NonceSize returns the size of the nonce that must be passed to Seal
	// and Open.
	NonceSize() int

	// Overhead returns the maximum difference between the lengths of a
	// plaintext and its ciphertext.
	Overhead() int

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	//
	// To reuse plaintext's storage for the encrypted output, use plaintext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	Seal(dst, nonce, plaintext, additionalData []byte) []byte

	// Open decrypts and authenticates ciphertext, authenticates the
	// additional data and, if successful, appends the resulting plaintext
	// to dst, returning the updated slice. The nonce must be NonceSize()
	// bytes long and both it and the additional data must match the
	// value passed to Seal.
	//
	// To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	//
	// Even if the function fails, the contents of dst, up to its capacity,
	// may be overwritten.
	Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}

// gcmAble is an interface implemented by ciphers that have a specific optimized
//line /usr/local/go/src/crypto/cipher/gcm.go:49
// implementation of GCM, like crypto/aes. NewGCM will check for this interface
//line /usr/local/go/src/crypto/cipher/gcm.go:49
// and return the specific AEAD if found.
//line /usr/local/go/src/crypto/cipher/gcm.go:52
type gcmAble interface {
	NewGCM(nonceSize, tagSize int) (AEAD, error)
}

// gcmFieldElement represents a value in GF(2¹²⁸). In order to reflect the GCM
//line /usr/local/go/src/crypto/cipher/gcm.go:56
// standard and make binary.BigEndian suitable for marshaling these values, the
//line /usr/local/go/src/crypto/cipher/gcm.go:56
// bits are stored in big endian order. For example:
//line /usr/local/go/src/crypto/cipher/gcm.go:56
//
//line /usr/local/go/src/crypto/cipher/gcm.go:56
//	the coefficient of x⁰ can be obtained by v.low >> 63.
//line /usr/local/go/src/crypto/cipher/gcm.go:56
//	the coefficient of x⁶³ can be obtained by v.low & 1.
//line /usr/local/go/src/crypto/cipher/gcm.go:56
//	the coefficient of x⁶⁴ can be obtained by v.high >> 63.
//line /usr/local/go/src/crypto/cipher/gcm.go:56
//	the coefficient of x¹²⁷ can be obtained by v.high & 1.
//line /usr/local/go/src/crypto/cipher/gcm.go:64
type gcmFieldElement struct {
	low, high uint64
}

// gcm represents a Galois Counter Mode with a specific key. See
//line /usr/local/go/src/crypto/cipher/gcm.go:68
// https://csrc.nist.gov/groups/ST/toolkit/BCM/documents/proposedmodes/gcm/gcm-revised-spec.pdf
//line /usr/local/go/src/crypto/cipher/gcm.go:70
type gcm struct {
	cipher		Block
	nonceSize	int
	tagSize		int
	// productTable contains the first sixteen powers of the key, H.
	// However, they are in bit reversed order. See NewGCMWithNonceSize.
	productTable	[16]gcmFieldElement
}

// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode
//line /usr/local/go/src/crypto/cipher/gcm.go:79
// with the standard nonce length.
//line /usr/local/go/src/crypto/cipher/gcm.go:79
//
//line /usr/local/go/src/crypto/cipher/gcm.go:79
// In general, the GHASH operation performed by this implementation of GCM is not constant-time.
//line /usr/local/go/src/crypto/cipher/gcm.go:79
// An exception is when the underlying Block was created by aes.NewCipher
//line /usr/local/go/src/crypto/cipher/gcm.go:79
// on systems with hardware support for AES. See the crypto/aes package documentation for details.
//line /usr/local/go/src/crypto/cipher/gcm.go:85
func NewGCM(cipher Block) (AEAD, error) {
//line /usr/local/go/src/crypto/cipher/gcm.go:85
	_go_fuzz_dep_.CoverTab[1588]++
							return newGCMWithNonceAndTagSize(cipher, gcmStandardNonceSize, gcmTagSize)
//line /usr/local/go/src/crypto/cipher/gcm.go:86
	// _ = "end of CoverTab[1588]"
}

// NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois
//line /usr/local/go/src/crypto/cipher/gcm.go:89
// Counter Mode, which accepts nonces of the given length. The length must not
//line /usr/local/go/src/crypto/cipher/gcm.go:89
// be zero.
//line /usr/local/go/src/crypto/cipher/gcm.go:89
//
//line /usr/local/go/src/crypto/cipher/gcm.go:89
// Only use this function if you require compatibility with an existing
//line /usr/local/go/src/crypto/cipher/gcm.go:89
// cryptosystem that uses non-standard nonce lengths. All other users should use
//line /usr/local/go/src/crypto/cipher/gcm.go:89
// NewGCM, which is faster and more resistant to misuse.
//line /usr/local/go/src/crypto/cipher/gcm.go:96
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error) {
//line /usr/local/go/src/crypto/cipher/gcm.go:96
	_go_fuzz_dep_.CoverTab[1589]++
							return newGCMWithNonceAndTagSize(cipher, size, gcmTagSize)
//line /usr/local/go/src/crypto/cipher/gcm.go:97
	// _ = "end of CoverTab[1589]"
}

// NewGCMWithTagSize returns the given 128-bit, block cipher wrapped in Galois
//line /usr/local/go/src/crypto/cipher/gcm.go:100
// Counter Mode, which generates tags with the given length.
//line /usr/local/go/src/crypto/cipher/gcm.go:100
//
//line /usr/local/go/src/crypto/cipher/gcm.go:100
// Tag sizes between 12 and 16 bytes are allowed.
//line /usr/local/go/src/crypto/cipher/gcm.go:100
//
//line /usr/local/go/src/crypto/cipher/gcm.go:100
// Only use this function if you require compatibility with an existing
//line /usr/local/go/src/crypto/cipher/gcm.go:100
// cryptosystem that uses non-standard tag lengths. All other users should use
//line /usr/local/go/src/crypto/cipher/gcm.go:100
// NewGCM, which is more resistant to misuse.
//line /usr/local/go/src/crypto/cipher/gcm.go:108
func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error) {
//line /usr/local/go/src/crypto/cipher/gcm.go:108
	_go_fuzz_dep_.CoverTab[1590]++
							return newGCMWithNonceAndTagSize(cipher, gcmStandardNonceSize, tagSize)
//line /usr/local/go/src/crypto/cipher/gcm.go:109
	// _ = "end of CoverTab[1590]"
}

func newGCMWithNonceAndTagSize(cipher Block, nonceSize, tagSize int) (AEAD, error) {
//line /usr/local/go/src/crypto/cipher/gcm.go:112
	_go_fuzz_dep_.CoverTab[1591]++
							if tagSize < gcmMinimumTagSize || func() bool {
//line /usr/local/go/src/crypto/cipher/gcm.go:113
		_go_fuzz_dep_.CoverTab[1597]++
//line /usr/local/go/src/crypto/cipher/gcm.go:113
		return tagSize > gcmBlockSize
//line /usr/local/go/src/crypto/cipher/gcm.go:113
		// _ = "end of CoverTab[1597]"
//line /usr/local/go/src/crypto/cipher/gcm.go:113
	}() {
//line /usr/local/go/src/crypto/cipher/gcm.go:113
		_go_fuzz_dep_.CoverTab[1598]++
								return nil, errors.New("cipher: incorrect tag size given to GCM")
//line /usr/local/go/src/crypto/cipher/gcm.go:114
		// _ = "end of CoverTab[1598]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:115
		_go_fuzz_dep_.CoverTab[1599]++
//line /usr/local/go/src/crypto/cipher/gcm.go:115
		// _ = "end of CoverTab[1599]"
//line /usr/local/go/src/crypto/cipher/gcm.go:115
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:115
	// _ = "end of CoverTab[1591]"
//line /usr/local/go/src/crypto/cipher/gcm.go:115
	_go_fuzz_dep_.CoverTab[1592]++

							if nonceSize <= 0 {
//line /usr/local/go/src/crypto/cipher/gcm.go:117
		_go_fuzz_dep_.CoverTab[1600]++
								return nil, errors.New("cipher: the nonce can't have zero length, or the security of the key will be immediately compromised")
//line /usr/local/go/src/crypto/cipher/gcm.go:118
		// _ = "end of CoverTab[1600]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:119
		_go_fuzz_dep_.CoverTab[1601]++
//line /usr/local/go/src/crypto/cipher/gcm.go:119
		// _ = "end of CoverTab[1601]"
//line /usr/local/go/src/crypto/cipher/gcm.go:119
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:119
	// _ = "end of CoverTab[1592]"
//line /usr/local/go/src/crypto/cipher/gcm.go:119
	_go_fuzz_dep_.CoverTab[1593]++

							if cipher, ok := cipher.(gcmAble); ok {
//line /usr/local/go/src/crypto/cipher/gcm.go:121
		_go_fuzz_dep_.CoverTab[1602]++
								return cipher.NewGCM(nonceSize, tagSize)
//line /usr/local/go/src/crypto/cipher/gcm.go:122
		// _ = "end of CoverTab[1602]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:123
		_go_fuzz_dep_.CoverTab[1603]++
//line /usr/local/go/src/crypto/cipher/gcm.go:123
		// _ = "end of CoverTab[1603]"
//line /usr/local/go/src/crypto/cipher/gcm.go:123
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:123
	// _ = "end of CoverTab[1593]"
//line /usr/local/go/src/crypto/cipher/gcm.go:123
	_go_fuzz_dep_.CoverTab[1594]++

							if cipher.BlockSize() != gcmBlockSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:125
		_go_fuzz_dep_.CoverTab[1604]++
								return nil, errors.New("cipher: NewGCM requires 128-bit block cipher")
//line /usr/local/go/src/crypto/cipher/gcm.go:126
		// _ = "end of CoverTab[1604]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:127
		_go_fuzz_dep_.CoverTab[1605]++
//line /usr/local/go/src/crypto/cipher/gcm.go:127
		// _ = "end of CoverTab[1605]"
//line /usr/local/go/src/crypto/cipher/gcm.go:127
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:127
	// _ = "end of CoverTab[1594]"
//line /usr/local/go/src/crypto/cipher/gcm.go:127
	_go_fuzz_dep_.CoverTab[1595]++

							var key [gcmBlockSize]byte
							cipher.Encrypt(key[:], key[:])

							g := &gcm{cipher: cipher, nonceSize: nonceSize, tagSize: tagSize}

//line /usr/local/go/src/crypto/cipher/gcm.go:139
	x := gcmFieldElement{
		binary.BigEndian.Uint64(key[:8]),
		binary.BigEndian.Uint64(key[8:]),
	}
	g.productTable[reverseBits(1)] = x

	for i := 2; i < 16; i += 2 {
//line /usr/local/go/src/crypto/cipher/gcm.go:145
		_go_fuzz_dep_.CoverTab[1606]++
								g.productTable[reverseBits(i)] = gcmDouble(&g.productTable[reverseBits(i/2)])
								g.productTable[reverseBits(i+1)] = gcmAdd(&g.productTable[reverseBits(i)], &x)
//line /usr/local/go/src/crypto/cipher/gcm.go:147
		// _ = "end of CoverTab[1606]"
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:148
	// _ = "end of CoverTab[1595]"
//line /usr/local/go/src/crypto/cipher/gcm.go:148
	_go_fuzz_dep_.CoverTab[1596]++

							return g, nil
//line /usr/local/go/src/crypto/cipher/gcm.go:150
	// _ = "end of CoverTab[1596]"
}

const (
	gcmBlockSize		= 16
	gcmTagSize		= 16
	gcmMinimumTagSize	= 12	// NIST SP 800-38D recommends tags with 12 or more bytes.
	gcmStandardNonceSize	= 12
)

func (g *gcm) NonceSize() int {
//line /usr/local/go/src/crypto/cipher/gcm.go:160
	_go_fuzz_dep_.CoverTab[1607]++
							return g.nonceSize
//line /usr/local/go/src/crypto/cipher/gcm.go:161
	// _ = "end of CoverTab[1607]"
}

func (g *gcm) Overhead() int {
//line /usr/local/go/src/crypto/cipher/gcm.go:164
	_go_fuzz_dep_.CoverTab[1608]++
							return g.tagSize
//line /usr/local/go/src/crypto/cipher/gcm.go:165
	// _ = "end of CoverTab[1608]"
}

func (g *gcm) Seal(dst, nonce, plaintext, data []byte) []byte {
//line /usr/local/go/src/crypto/cipher/gcm.go:168
	_go_fuzz_dep_.CoverTab[1609]++
							if len(nonce) != g.nonceSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:169
		_go_fuzz_dep_.CoverTab[1613]++
								panic("crypto/cipher: incorrect nonce length given to GCM")
//line /usr/local/go/src/crypto/cipher/gcm.go:170
		// _ = "end of CoverTab[1613]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:171
		_go_fuzz_dep_.CoverTab[1614]++
//line /usr/local/go/src/crypto/cipher/gcm.go:171
		// _ = "end of CoverTab[1614]"
//line /usr/local/go/src/crypto/cipher/gcm.go:171
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:171
	// _ = "end of CoverTab[1609]"
//line /usr/local/go/src/crypto/cipher/gcm.go:171
	_go_fuzz_dep_.CoverTab[1610]++
							if uint64(len(plaintext)) > ((1<<32)-2)*uint64(g.cipher.BlockSize()) {
//line /usr/local/go/src/crypto/cipher/gcm.go:172
		_go_fuzz_dep_.CoverTab[1615]++
								panic("crypto/cipher: message too large for GCM")
//line /usr/local/go/src/crypto/cipher/gcm.go:173
		// _ = "end of CoverTab[1615]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:174
		_go_fuzz_dep_.CoverTab[1616]++
//line /usr/local/go/src/crypto/cipher/gcm.go:174
		// _ = "end of CoverTab[1616]"
//line /usr/local/go/src/crypto/cipher/gcm.go:174
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:174
	// _ = "end of CoverTab[1610]"
//line /usr/local/go/src/crypto/cipher/gcm.go:174
	_go_fuzz_dep_.CoverTab[1611]++

							ret, out := sliceForAppend(dst, len(plaintext)+g.tagSize)
							if alias.InexactOverlap(out, plaintext) {
//line /usr/local/go/src/crypto/cipher/gcm.go:177
		_go_fuzz_dep_.CoverTab[1617]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/gcm.go:178
		// _ = "end of CoverTab[1617]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:179
		_go_fuzz_dep_.CoverTab[1618]++
//line /usr/local/go/src/crypto/cipher/gcm.go:179
		// _ = "end of CoverTab[1618]"
//line /usr/local/go/src/crypto/cipher/gcm.go:179
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:179
	// _ = "end of CoverTab[1611]"
//line /usr/local/go/src/crypto/cipher/gcm.go:179
	_go_fuzz_dep_.CoverTab[1612]++

							var counter, tagMask [gcmBlockSize]byte
							g.deriveCounter(&counter, nonce)

							g.cipher.Encrypt(tagMask[:], counter[:])
							gcmInc32(&counter)

							g.counterCrypt(out, plaintext, &counter)

							var tag [gcmTagSize]byte
							g.auth(tag[:], out[:len(plaintext)], data, &tagMask)
							copy(out[len(plaintext):], tag[:])

							return ret
//line /usr/local/go/src/crypto/cipher/gcm.go:193
	// _ = "end of CoverTab[1612]"
}

var errOpen = errors.New("cipher: message authentication failed")

func (g *gcm) Open(dst, nonce, ciphertext, data []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/cipher/gcm.go:198
	_go_fuzz_dep_.CoverTab[1619]++
							if len(nonce) != g.nonceSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:199
		_go_fuzz_dep_.CoverTab[1626]++
								panic("crypto/cipher: incorrect nonce length given to GCM")
//line /usr/local/go/src/crypto/cipher/gcm.go:200
		// _ = "end of CoverTab[1626]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:201
		_go_fuzz_dep_.CoverTab[1627]++
//line /usr/local/go/src/crypto/cipher/gcm.go:201
		// _ = "end of CoverTab[1627]"
//line /usr/local/go/src/crypto/cipher/gcm.go:201
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:201
	// _ = "end of CoverTab[1619]"
//line /usr/local/go/src/crypto/cipher/gcm.go:201
	_go_fuzz_dep_.CoverTab[1620]++

//line /usr/local/go/src/crypto/cipher/gcm.go:204
	if g.tagSize < gcmMinimumTagSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:204
		_go_fuzz_dep_.CoverTab[1628]++
								panic("crypto/cipher: incorrect GCM tag size")
//line /usr/local/go/src/crypto/cipher/gcm.go:205
		// _ = "end of CoverTab[1628]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:206
		_go_fuzz_dep_.CoverTab[1629]++
//line /usr/local/go/src/crypto/cipher/gcm.go:206
		// _ = "end of CoverTab[1629]"
//line /usr/local/go/src/crypto/cipher/gcm.go:206
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:206
	// _ = "end of CoverTab[1620]"
//line /usr/local/go/src/crypto/cipher/gcm.go:206
	_go_fuzz_dep_.CoverTab[1621]++

							if len(ciphertext) < g.tagSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:208
		_go_fuzz_dep_.CoverTab[1630]++
								return nil, errOpen
//line /usr/local/go/src/crypto/cipher/gcm.go:209
		// _ = "end of CoverTab[1630]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:210
		_go_fuzz_dep_.CoverTab[1631]++
//line /usr/local/go/src/crypto/cipher/gcm.go:210
		// _ = "end of CoverTab[1631]"
//line /usr/local/go/src/crypto/cipher/gcm.go:210
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:210
	// _ = "end of CoverTab[1621]"
//line /usr/local/go/src/crypto/cipher/gcm.go:210
	_go_fuzz_dep_.CoverTab[1622]++
							if uint64(len(ciphertext)) > ((1<<32)-2)*uint64(g.cipher.BlockSize())+uint64(g.tagSize) {
//line /usr/local/go/src/crypto/cipher/gcm.go:211
		_go_fuzz_dep_.CoverTab[1632]++
								return nil, errOpen
//line /usr/local/go/src/crypto/cipher/gcm.go:212
		// _ = "end of CoverTab[1632]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:213
		_go_fuzz_dep_.CoverTab[1633]++
//line /usr/local/go/src/crypto/cipher/gcm.go:213
		// _ = "end of CoverTab[1633]"
//line /usr/local/go/src/crypto/cipher/gcm.go:213
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:213
	// _ = "end of CoverTab[1622]"
//line /usr/local/go/src/crypto/cipher/gcm.go:213
	_go_fuzz_dep_.CoverTab[1623]++

							tag := ciphertext[len(ciphertext)-g.tagSize:]
							ciphertext = ciphertext[:len(ciphertext)-g.tagSize]

							var counter, tagMask [gcmBlockSize]byte
							g.deriveCounter(&counter, nonce)

							g.cipher.Encrypt(tagMask[:], counter[:])
							gcmInc32(&counter)

							var expectedTag [gcmTagSize]byte
							g.auth(expectedTag[:], ciphertext, data, &tagMask)

							ret, out := sliceForAppend(dst, len(ciphertext))
							if alias.InexactOverlap(out, ciphertext) {
//line /usr/local/go/src/crypto/cipher/gcm.go:228
		_go_fuzz_dep_.CoverTab[1634]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/gcm.go:229
		// _ = "end of CoverTab[1634]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:230
		_go_fuzz_dep_.CoverTab[1635]++
//line /usr/local/go/src/crypto/cipher/gcm.go:230
		// _ = "end of CoverTab[1635]"
//line /usr/local/go/src/crypto/cipher/gcm.go:230
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:230
	// _ = "end of CoverTab[1623]"
//line /usr/local/go/src/crypto/cipher/gcm.go:230
	_go_fuzz_dep_.CoverTab[1624]++

							if subtle.ConstantTimeCompare(expectedTag[:g.tagSize], tag) != 1 {
//line /usr/local/go/src/crypto/cipher/gcm.go:232
		_go_fuzz_dep_.CoverTab[1636]++

//line /usr/local/go/src/crypto/cipher/gcm.go:237
		for i := range out {
//line /usr/local/go/src/crypto/cipher/gcm.go:237
			_go_fuzz_dep_.CoverTab[1638]++
									out[i] = 0
//line /usr/local/go/src/crypto/cipher/gcm.go:238
			// _ = "end of CoverTab[1638]"
		}
//line /usr/local/go/src/crypto/cipher/gcm.go:239
		// _ = "end of CoverTab[1636]"
//line /usr/local/go/src/crypto/cipher/gcm.go:239
		_go_fuzz_dep_.CoverTab[1637]++
								return nil, errOpen
//line /usr/local/go/src/crypto/cipher/gcm.go:240
		// _ = "end of CoverTab[1637]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:241
		_go_fuzz_dep_.CoverTab[1639]++
//line /usr/local/go/src/crypto/cipher/gcm.go:241
		// _ = "end of CoverTab[1639]"
//line /usr/local/go/src/crypto/cipher/gcm.go:241
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:241
	// _ = "end of CoverTab[1624]"
//line /usr/local/go/src/crypto/cipher/gcm.go:241
	_go_fuzz_dep_.CoverTab[1625]++

							g.counterCrypt(out, ciphertext, &counter)

							return ret, nil
//line /usr/local/go/src/crypto/cipher/gcm.go:245
	// _ = "end of CoverTab[1625]"
}

// reverseBits reverses the order of the bits of 4-bit number in i.
func reverseBits(i int) int {
//line /usr/local/go/src/crypto/cipher/gcm.go:249
	_go_fuzz_dep_.CoverTab[1640]++
							i = ((i << 2) & 0xc) | ((i >> 2) & 0x3)
							i = ((i << 1) & 0xa) | ((i >> 1) & 0x5)
							return i
//line /usr/local/go/src/crypto/cipher/gcm.go:252
	// _ = "end of CoverTab[1640]"
}

// gcmAdd adds two elements of GF(2¹²⁸) and returns the sum.
func gcmAdd(x, y *gcmFieldElement) gcmFieldElement {
//line /usr/local/go/src/crypto/cipher/gcm.go:256
	_go_fuzz_dep_.CoverTab[1641]++

							return gcmFieldElement{x.low ^ y.low, x.high ^ y.high}
//line /usr/local/go/src/crypto/cipher/gcm.go:258
	// _ = "end of CoverTab[1641]"
}

// gcmDouble returns the result of doubling an element of GF(2¹²⁸).
func gcmDouble(x *gcmFieldElement) (double gcmFieldElement) {
//line /usr/local/go/src/crypto/cipher/gcm.go:262
	_go_fuzz_dep_.CoverTab[1642]++
							msbSet := x.high&1 == 1

//line /usr/local/go/src/crypto/cipher/gcm.go:266
	double.high = x.high >> 1
							double.high |= x.low << 63
							double.low = x.low >> 1

//line /usr/local/go/src/crypto/cipher/gcm.go:277
	if msbSet {
//line /usr/local/go/src/crypto/cipher/gcm.go:277
		_go_fuzz_dep_.CoverTab[1644]++
								double.low ^= 0xe100000000000000
//line /usr/local/go/src/crypto/cipher/gcm.go:278
		// _ = "end of CoverTab[1644]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:279
		_go_fuzz_dep_.CoverTab[1645]++
//line /usr/local/go/src/crypto/cipher/gcm.go:279
		// _ = "end of CoverTab[1645]"
//line /usr/local/go/src/crypto/cipher/gcm.go:279
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:279
	// _ = "end of CoverTab[1642]"
//line /usr/local/go/src/crypto/cipher/gcm.go:279
	_go_fuzz_dep_.CoverTab[1643]++

							return
//line /usr/local/go/src/crypto/cipher/gcm.go:281
	// _ = "end of CoverTab[1643]"
}

var gcmReductionTable = []uint16{
	0x0000, 0x1c20, 0x3840, 0x2460, 0x7080, 0x6ca0, 0x48c0, 0x54e0,
	0xe100, 0xfd20, 0xd940, 0xc560, 0x9180, 0x8da0, 0xa9c0, 0xb5e0,
}

// mul sets y to y*H, where H is the GCM key, fixed during NewGCMWithNonceSize.
func (g *gcm) mul(y *gcmFieldElement) {
//line /usr/local/go/src/crypto/cipher/gcm.go:290
	_go_fuzz_dep_.CoverTab[1646]++
							var z gcmFieldElement

							for i := 0; i < 2; i++ {
//line /usr/local/go/src/crypto/cipher/gcm.go:293
		_go_fuzz_dep_.CoverTab[1648]++
								word := y.high
								if i == 1 {
//line /usr/local/go/src/crypto/cipher/gcm.go:295
			_go_fuzz_dep_.CoverTab[1650]++
									word = y.low
//line /usr/local/go/src/crypto/cipher/gcm.go:296
			// _ = "end of CoverTab[1650]"
		} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:297
			_go_fuzz_dep_.CoverTab[1651]++
//line /usr/local/go/src/crypto/cipher/gcm.go:297
			// _ = "end of CoverTab[1651]"
//line /usr/local/go/src/crypto/cipher/gcm.go:297
		}
//line /usr/local/go/src/crypto/cipher/gcm.go:297
		// _ = "end of CoverTab[1648]"
//line /usr/local/go/src/crypto/cipher/gcm.go:297
		_go_fuzz_dep_.CoverTab[1649]++

//line /usr/local/go/src/crypto/cipher/gcm.go:301
		for j := 0; j < 64; j += 4 {
//line /usr/local/go/src/crypto/cipher/gcm.go:301
			_go_fuzz_dep_.CoverTab[1652]++
									msw := z.high & 0xf
									z.high >>= 4
									z.high |= z.low << 60
									z.low >>= 4
									z.low ^= uint64(gcmReductionTable[msw]) << 48

//line /usr/local/go/src/crypto/cipher/gcm.go:311
			t := &g.productTable[word&0xf]

									z.low ^= t.low
									z.high ^= t.high
									word >>= 4
//line /usr/local/go/src/crypto/cipher/gcm.go:315
			// _ = "end of CoverTab[1652]"
		}
//line /usr/local/go/src/crypto/cipher/gcm.go:316
		// _ = "end of CoverTab[1649]"
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:317
	// _ = "end of CoverTab[1646]"
//line /usr/local/go/src/crypto/cipher/gcm.go:317
	_go_fuzz_dep_.CoverTab[1647]++

							*y = z
//line /usr/local/go/src/crypto/cipher/gcm.go:319
	// _ = "end of CoverTab[1647]"
}

// updateBlocks extends y with more polynomial terms from blocks, based on
//line /usr/local/go/src/crypto/cipher/gcm.go:322
// Horner's rule. There must be a multiple of gcmBlockSize bytes in blocks.
//line /usr/local/go/src/crypto/cipher/gcm.go:324
func (g *gcm) updateBlocks(y *gcmFieldElement, blocks []byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:324
	_go_fuzz_dep_.CoverTab[1653]++
							for len(blocks) > 0 {
//line /usr/local/go/src/crypto/cipher/gcm.go:325
		_go_fuzz_dep_.CoverTab[1654]++
								y.low ^= binary.BigEndian.Uint64(blocks)
								y.high ^= binary.BigEndian.Uint64(blocks[8:])
								g.mul(y)
								blocks = blocks[gcmBlockSize:]
//line /usr/local/go/src/crypto/cipher/gcm.go:329
		// _ = "end of CoverTab[1654]"
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:330
	// _ = "end of CoverTab[1653]"
}

// update extends y with more polynomial terms from data. If data is not a
//line /usr/local/go/src/crypto/cipher/gcm.go:333
// multiple of gcmBlockSize bytes long then the remainder is zero padded.
//line /usr/local/go/src/crypto/cipher/gcm.go:335
func (g *gcm) update(y *gcmFieldElement, data []byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:335
	_go_fuzz_dep_.CoverTab[1655]++
							fullBlocks := (len(data) >> 4) << 4
							g.updateBlocks(y, data[:fullBlocks])

							if len(data) != fullBlocks {
//line /usr/local/go/src/crypto/cipher/gcm.go:339
		_go_fuzz_dep_.CoverTab[1656]++
								var partialBlock [gcmBlockSize]byte
								copy(partialBlock[:], data[fullBlocks:])
								g.updateBlocks(y, partialBlock[:])
//line /usr/local/go/src/crypto/cipher/gcm.go:342
		// _ = "end of CoverTab[1656]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:343
		_go_fuzz_dep_.CoverTab[1657]++
//line /usr/local/go/src/crypto/cipher/gcm.go:343
		// _ = "end of CoverTab[1657]"
//line /usr/local/go/src/crypto/cipher/gcm.go:343
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:343
	// _ = "end of CoverTab[1655]"
}

// gcmInc32 treats the final four bytes of counterBlock as a big-endian value
//line /usr/local/go/src/crypto/cipher/gcm.go:346
// and increments it.
//line /usr/local/go/src/crypto/cipher/gcm.go:348
func gcmInc32(counterBlock *[16]byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:348
	_go_fuzz_dep_.CoverTab[1658]++
							ctr := counterBlock[len(counterBlock)-4:]
							binary.BigEndian.PutUint32(ctr, binary.BigEndian.Uint32(ctr)+1)
//line /usr/local/go/src/crypto/cipher/gcm.go:350
	// _ = "end of CoverTab[1658]"
}

// sliceForAppend takes a slice and a requested number of bytes. It returns a
//line /usr/local/go/src/crypto/cipher/gcm.go:353
// slice with the contents of the given slice followed by that many bytes and a
//line /usr/local/go/src/crypto/cipher/gcm.go:353
// second slice that aliases into it and contains only the extra bytes. If the
//line /usr/local/go/src/crypto/cipher/gcm.go:353
// original slice has sufficient capacity then no allocation is performed.
//line /usr/local/go/src/crypto/cipher/gcm.go:357
func sliceForAppend(in []byte, n int) (head, tail []byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:357
	_go_fuzz_dep_.CoverTab[1659]++
							if total := len(in) + n; cap(in) >= total {
//line /usr/local/go/src/crypto/cipher/gcm.go:358
		_go_fuzz_dep_.CoverTab[1661]++
								head = in[:total]
//line /usr/local/go/src/crypto/cipher/gcm.go:359
		// _ = "end of CoverTab[1661]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:360
		_go_fuzz_dep_.CoverTab[1662]++
								head = make([]byte, total)
								copy(head, in)
//line /usr/local/go/src/crypto/cipher/gcm.go:362
		// _ = "end of CoverTab[1662]"
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:363
	// _ = "end of CoverTab[1659]"
//line /usr/local/go/src/crypto/cipher/gcm.go:363
	_go_fuzz_dep_.CoverTab[1660]++
							tail = head[len(in):]
							return
//line /usr/local/go/src/crypto/cipher/gcm.go:365
	// _ = "end of CoverTab[1660]"
}

// counterCrypt crypts in to out using g.cipher in counter mode.
func (g *gcm) counterCrypt(out, in []byte, counter *[gcmBlockSize]byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:369
	_go_fuzz_dep_.CoverTab[1663]++
							var mask [gcmBlockSize]byte

							for len(in) >= gcmBlockSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:372
		_go_fuzz_dep_.CoverTab[1665]++
								g.cipher.Encrypt(mask[:], counter[:])
								gcmInc32(counter)

								subtle.XORBytes(out, in, mask[:])
								out = out[gcmBlockSize:]
								in = in[gcmBlockSize:]
//line /usr/local/go/src/crypto/cipher/gcm.go:378
		// _ = "end of CoverTab[1665]"
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:379
	// _ = "end of CoverTab[1663]"
//line /usr/local/go/src/crypto/cipher/gcm.go:379
	_go_fuzz_dep_.CoverTab[1664]++

							if len(in) > 0 {
//line /usr/local/go/src/crypto/cipher/gcm.go:381
		_go_fuzz_dep_.CoverTab[1666]++
								g.cipher.Encrypt(mask[:], counter[:])
								gcmInc32(counter)
								subtle.XORBytes(out, in, mask[:])
//line /usr/local/go/src/crypto/cipher/gcm.go:384
		// _ = "end of CoverTab[1666]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:385
		_go_fuzz_dep_.CoverTab[1667]++
//line /usr/local/go/src/crypto/cipher/gcm.go:385
		// _ = "end of CoverTab[1667]"
//line /usr/local/go/src/crypto/cipher/gcm.go:385
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:385
	// _ = "end of CoverTab[1664]"
}

// deriveCounter computes the initial GCM counter state from the given nonce.
//line /usr/local/go/src/crypto/cipher/gcm.go:388
// See NIST SP 800-38D, section 7.1. This assumes that counter is filled with
//line /usr/local/go/src/crypto/cipher/gcm.go:388
// zeros on entry.
//line /usr/local/go/src/crypto/cipher/gcm.go:391
func (g *gcm) deriveCounter(counter *[gcmBlockSize]byte, nonce []byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:391
	_go_fuzz_dep_.CoverTab[1668]++

//line /usr/local/go/src/crypto/cipher/gcm.go:398
	if len(nonce) == gcmStandardNonceSize {
//line /usr/local/go/src/crypto/cipher/gcm.go:398
		_go_fuzz_dep_.CoverTab[1669]++
								copy(counter[:], nonce)
								counter[gcmBlockSize-1] = 1
//line /usr/local/go/src/crypto/cipher/gcm.go:400
		// _ = "end of CoverTab[1669]"
	} else {
//line /usr/local/go/src/crypto/cipher/gcm.go:401
		_go_fuzz_dep_.CoverTab[1670]++
								var y gcmFieldElement
								g.update(&y, nonce)
								y.high ^= uint64(len(nonce)) * 8
								g.mul(&y)
								binary.BigEndian.PutUint64(counter[:8], y.low)
								binary.BigEndian.PutUint64(counter[8:], y.high)
//line /usr/local/go/src/crypto/cipher/gcm.go:407
		// _ = "end of CoverTab[1670]"
	}
//line /usr/local/go/src/crypto/cipher/gcm.go:408
	// _ = "end of CoverTab[1668]"
}

// auth calculates GHASH(ciphertext, additionalData), masks the result with
//line /usr/local/go/src/crypto/cipher/gcm.go:411
// tagMask and writes the result to out.
//line /usr/local/go/src/crypto/cipher/gcm.go:413
func (g *gcm) auth(out, ciphertext, additionalData []byte, tagMask *[gcmTagSize]byte) {
//line /usr/local/go/src/crypto/cipher/gcm.go:413
	_go_fuzz_dep_.CoverTab[1671]++
							var y gcmFieldElement
							g.update(&y, additionalData)
							g.update(&y, ciphertext)

							y.low ^= uint64(len(additionalData)) * 8
							y.high ^= uint64(len(ciphertext)) * 8

							g.mul(&y)

							binary.BigEndian.PutUint64(out, y.low)
							binary.BigEndian.PutUint64(out[8:], y.high)

							subtle.XORBytes(out, out, tagMask[:])
//line /usr/local/go/src/crypto/cipher/gcm.go:426
	// _ = "end of CoverTab[1671]"
}

//line /usr/local/go/src/crypto/cipher/gcm.go:427
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/gcm.go:427
var _ = _go_fuzz_dep_.CoverTab
