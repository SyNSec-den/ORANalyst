// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64

//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
package aes

//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
import (
//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
)
//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
import (
//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:7
)

import (
	"crypto/cipher"
	"crypto/internal/alias"
	"crypto/subtle"
	"errors"
)

//line /usr/local/go/src/crypto/aes/aes_gcm.go:18
//go:noescape
func gcmAesInit(productTable *[256]byte, ks []uint32)

//go:noescape
func gcmAesData(productTable *[256]byte, data []byte, T *[16]byte)

//go:noescape
func gcmAesEnc(productTable *[256]byte, dst, src []byte, ctr, T *[16]byte, ks []uint32)

//go:noescape
func gcmAesDec(productTable *[256]byte, dst, src []byte, ctr, T *[16]byte, ks []uint32)

//go:noescape
func gcmAesFinish(productTable *[256]byte, tagMask, T *[16]byte, pLen, dLen uint64)

const (
	gcmBlockSize		= 16
	gcmTagSize		= 16
	gcmMinimumTagSize	= 12
	gcmStandardNonceSize	= 12
)

var errOpen = errors.New("cipher: message authentication failed")

//line /usr/local/go/src/crypto/aes/aes_gcm.go:43
var _ gcmAble = (*aesCipherGCM)(nil)

//line /usr/local/go/src/crypto/aes/aes_gcm.go:47
func (c *aesCipherGCM) NewGCM(nonceSize, tagSize int) (cipher.AEAD, error) {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:47
	_go_fuzz_dep_.CoverTab[1746]++
							g := &gcmAsm{ks: c.enc, nonceSize: nonceSize, tagSize: tagSize}
							gcmAesInit(&g.productTable, g.ks)
							return g, nil
//line /usr/local/go/src/crypto/aes/aes_gcm.go:50
	// _ = "end of CoverTab[1746]"
}

type gcmAsm struct {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:56
	ks	[]uint32

//line /usr/local/go/src/crypto/aes/aes_gcm.go:59
	productTable	[256]byte

	nonceSize	int

	tagSize	int
}

func (g *gcmAsm) NonceSize() int {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:66
	_go_fuzz_dep_.CoverTab[1747]++
							return g.nonceSize
//line /usr/local/go/src/crypto/aes/aes_gcm.go:67
	// _ = "end of CoverTab[1747]"
}

func (g *gcmAsm) Overhead() int {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:70
	_go_fuzz_dep_.CoverTab[1748]++
							return g.tagSize
//line /usr/local/go/src/crypto/aes/aes_gcm.go:71
	// _ = "end of CoverTab[1748]"
}

//line /usr/local/go/src/crypto/aes/aes_gcm.go:78
func sliceForAppend(in []byte, n int) (head, tail []byte) {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:78
	_go_fuzz_dep_.CoverTab[1749]++
							if total := len(in) + n; cap(in) >= total {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:79
		_go_fuzz_dep_.CoverTab[1751]++
								head = in[:total]
//line /usr/local/go/src/crypto/aes/aes_gcm.go:80
		// _ = "end of CoverTab[1751]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:81
		_go_fuzz_dep_.CoverTab[1752]++
								head = make([]byte, total)
								copy(head, in)
//line /usr/local/go/src/crypto/aes/aes_gcm.go:83
		// _ = "end of CoverTab[1752]"
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:84
	// _ = "end of CoverTab[1749]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:84
	_go_fuzz_dep_.CoverTab[1750]++
							tail = head[len(in):]
							return
//line /usr/local/go/src/crypto/aes/aes_gcm.go:86
	// _ = "end of CoverTab[1750]"
}

//line /usr/local/go/src/crypto/aes/aes_gcm.go:91
func (g *gcmAsm) Seal(dst, nonce, plaintext, data []byte) []byte {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:91
	_go_fuzz_dep_.CoverTab[1753]++
							if len(nonce) != g.nonceSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:92
		_go_fuzz_dep_.CoverTab[1759]++
								panic("crypto/cipher: incorrect nonce length given to GCM")
//line /usr/local/go/src/crypto/aes/aes_gcm.go:93
		// _ = "end of CoverTab[1759]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:94
		_go_fuzz_dep_.CoverTab[1760]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:94
		// _ = "end of CoverTab[1760]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:94
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:94
	// _ = "end of CoverTab[1753]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:94
	_go_fuzz_dep_.CoverTab[1754]++
							if uint64(len(plaintext)) > ((1<<32)-2)*BlockSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:95
		_go_fuzz_dep_.CoverTab[1761]++
								panic("crypto/cipher: message too large for GCM")
//line /usr/local/go/src/crypto/aes/aes_gcm.go:96
		// _ = "end of CoverTab[1761]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:97
		_go_fuzz_dep_.CoverTab[1762]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:97
		// _ = "end of CoverTab[1762]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:97
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:97
	// _ = "end of CoverTab[1754]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:97
	_go_fuzz_dep_.CoverTab[1755]++

							var counter, tagMask [gcmBlockSize]byte

							if len(nonce) == gcmStandardNonceSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:101
		_go_fuzz_dep_.CoverTab[1763]++

								copy(counter[:], nonce)
								counter[gcmBlockSize-1] = 1
//line /usr/local/go/src/crypto/aes/aes_gcm.go:104
		// _ = "end of CoverTab[1763]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:105
		_go_fuzz_dep_.CoverTab[1764]++

								gcmAesData(&g.productTable, nonce, &counter)
								gcmAesFinish(&g.productTable, &tagMask, &counter, uint64(len(nonce)), uint64(0))
//line /usr/local/go/src/crypto/aes/aes_gcm.go:108
		// _ = "end of CoverTab[1764]"
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:109
	// _ = "end of CoverTab[1755]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:109
	_go_fuzz_dep_.CoverTab[1756]++

							encryptBlockAsm(len(g.ks)/4-1, &g.ks[0], &tagMask[0], &counter[0])

							var tagOut [gcmTagSize]byte
							gcmAesData(&g.productTable, data, &tagOut)

							ret, out := sliceForAppend(dst, len(plaintext)+g.tagSize)
							if alias.InexactOverlap(out[:len(plaintext)], plaintext) {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:117
		_go_fuzz_dep_.CoverTab[1765]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/aes/aes_gcm.go:118
		// _ = "end of CoverTab[1765]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:119
		_go_fuzz_dep_.CoverTab[1766]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:119
		// _ = "end of CoverTab[1766]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:119
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:119
	// _ = "end of CoverTab[1756]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:119
	_go_fuzz_dep_.CoverTab[1757]++
							if len(plaintext) > 0 {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:120
		_go_fuzz_dep_.CoverTab[1767]++
								gcmAesEnc(&g.productTable, out, plaintext, &counter, &tagOut, g.ks)
//line /usr/local/go/src/crypto/aes/aes_gcm.go:121
		// _ = "end of CoverTab[1767]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:122
		_go_fuzz_dep_.CoverTab[1768]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:122
		// _ = "end of CoverTab[1768]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:122
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:122
	// _ = "end of CoverTab[1757]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:122
	_go_fuzz_dep_.CoverTab[1758]++
							gcmAesFinish(&g.productTable, &tagMask, &tagOut, uint64(len(plaintext)), uint64(len(data)))
							copy(out[len(plaintext):], tagOut[:])

							return ret
//line /usr/local/go/src/crypto/aes/aes_gcm.go:126
	// _ = "end of CoverTab[1758]"
}

//line /usr/local/go/src/crypto/aes/aes_gcm.go:131
func (g *gcmAsm) Open(dst, nonce, ciphertext, data []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:131
	_go_fuzz_dep_.CoverTab[1769]++
							if len(nonce) != g.nonceSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:132
		_go_fuzz_dep_.CoverTab[1778]++
								panic("crypto/cipher: incorrect nonce length given to GCM")
//line /usr/local/go/src/crypto/aes/aes_gcm.go:133
		// _ = "end of CoverTab[1778]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:134
		_go_fuzz_dep_.CoverTab[1779]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:134
		// _ = "end of CoverTab[1779]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:134
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:134
	// _ = "end of CoverTab[1769]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:134
	_go_fuzz_dep_.CoverTab[1770]++

//line /usr/local/go/src/crypto/aes/aes_gcm.go:137
	if g.tagSize < gcmMinimumTagSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:137
		_go_fuzz_dep_.CoverTab[1780]++
								panic("crypto/cipher: incorrect GCM tag size")
//line /usr/local/go/src/crypto/aes/aes_gcm.go:138
		// _ = "end of CoverTab[1780]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:139
		_go_fuzz_dep_.CoverTab[1781]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:139
		// _ = "end of CoverTab[1781]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:139
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:139
	// _ = "end of CoverTab[1770]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:139
	_go_fuzz_dep_.CoverTab[1771]++

							if len(ciphertext) < g.tagSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:141
		_go_fuzz_dep_.CoverTab[1782]++
								return nil, errOpen
//line /usr/local/go/src/crypto/aes/aes_gcm.go:142
		// _ = "end of CoverTab[1782]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:143
		_go_fuzz_dep_.CoverTab[1783]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:143
		// _ = "end of CoverTab[1783]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:143
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:143
	// _ = "end of CoverTab[1771]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:143
	_go_fuzz_dep_.CoverTab[1772]++
							if uint64(len(ciphertext)) > ((1<<32)-2)*uint64(BlockSize)+uint64(g.tagSize) {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:144
		_go_fuzz_dep_.CoverTab[1784]++
								return nil, errOpen
//line /usr/local/go/src/crypto/aes/aes_gcm.go:145
		// _ = "end of CoverTab[1784]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:146
		_go_fuzz_dep_.CoverTab[1785]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:146
		// _ = "end of CoverTab[1785]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:146
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:146
	// _ = "end of CoverTab[1772]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:146
	_go_fuzz_dep_.CoverTab[1773]++

							tag := ciphertext[len(ciphertext)-g.tagSize:]
							ciphertext = ciphertext[:len(ciphertext)-g.tagSize]

//line /usr/local/go/src/crypto/aes/aes_gcm.go:152
	var counter, tagMask [gcmBlockSize]byte

	if len(nonce) == gcmStandardNonceSize {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:154
		_go_fuzz_dep_.CoverTab[1786]++

								copy(counter[:], nonce)
								counter[gcmBlockSize-1] = 1
//line /usr/local/go/src/crypto/aes/aes_gcm.go:157
		// _ = "end of CoverTab[1786]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:158
		_go_fuzz_dep_.CoverTab[1787]++

								gcmAesData(&g.productTable, nonce, &counter)
								gcmAesFinish(&g.productTable, &tagMask, &counter, uint64(len(nonce)), uint64(0))
//line /usr/local/go/src/crypto/aes/aes_gcm.go:161
		// _ = "end of CoverTab[1787]"
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:162
	// _ = "end of CoverTab[1773]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:162
	_go_fuzz_dep_.CoverTab[1774]++

							encryptBlockAsm(len(g.ks)/4-1, &g.ks[0], &tagMask[0], &counter[0])

							var expectedTag [gcmTagSize]byte
							gcmAesData(&g.productTable, data, &expectedTag)

							ret, out := sliceForAppend(dst, len(ciphertext))
							if alias.InexactOverlap(out, ciphertext) {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:170
		_go_fuzz_dep_.CoverTab[1788]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/aes/aes_gcm.go:171
		// _ = "end of CoverTab[1788]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:172
		_go_fuzz_dep_.CoverTab[1789]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:172
		// _ = "end of CoverTab[1789]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:172
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:172
	// _ = "end of CoverTab[1774]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:172
	_go_fuzz_dep_.CoverTab[1775]++
							if len(ciphertext) > 0 {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:173
		_go_fuzz_dep_.CoverTab[1790]++
								gcmAesDec(&g.productTable, out, ciphertext, &counter, &expectedTag, g.ks)
//line /usr/local/go/src/crypto/aes/aes_gcm.go:174
		// _ = "end of CoverTab[1790]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:175
		_go_fuzz_dep_.CoverTab[1791]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:175
		// _ = "end of CoverTab[1791]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:175
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:175
	// _ = "end of CoverTab[1775]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:175
	_go_fuzz_dep_.CoverTab[1776]++
							gcmAesFinish(&g.productTable, &tagMask, &expectedTag, uint64(len(ciphertext)), uint64(len(data)))

							if subtle.ConstantTimeCompare(expectedTag[:g.tagSize], tag) != 1 {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:178
		_go_fuzz_dep_.CoverTab[1792]++
								for i := range out {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:179
			_go_fuzz_dep_.CoverTab[1794]++
									out[i] = 0
//line /usr/local/go/src/crypto/aes/aes_gcm.go:180
			// _ = "end of CoverTab[1794]"
		}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:181
		// _ = "end of CoverTab[1792]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:181
		_go_fuzz_dep_.CoverTab[1793]++
								return nil, errOpen
//line /usr/local/go/src/crypto/aes/aes_gcm.go:182
		// _ = "end of CoverTab[1793]"
	} else {
//line /usr/local/go/src/crypto/aes/aes_gcm.go:183
		_go_fuzz_dep_.CoverTab[1795]++
//line /usr/local/go/src/crypto/aes/aes_gcm.go:183
		// _ = "end of CoverTab[1795]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:183
	}
//line /usr/local/go/src/crypto/aes/aes_gcm.go:183
	// _ = "end of CoverTab[1776]"
//line /usr/local/go/src/crypto/aes/aes_gcm.go:183
	_go_fuzz_dep_.CoverTab[1777]++

							return ret, nil
//line /usr/local/go/src/crypto/aes/aes_gcm.go:185
	// _ = "end of CoverTab[1777]"
}

//line /usr/local/go/src/crypto/aes/aes_gcm.go:186
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/aes/aes_gcm.go:186
var _ = _go_fuzz_dep_.CoverTab
