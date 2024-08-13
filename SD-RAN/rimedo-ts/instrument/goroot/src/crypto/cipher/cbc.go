// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Cipher block chaining (CBC) mode.

// CBC provides confidentiality by xoring (chaining) each plaintext block
// with the previous ciphertext block before applying the block cipher.

// See NIST SP 800-38A, pp 10-11

//line /usr/local/go/src/crypto/cipher/cbc.go:12
package cipher

//line /usr/local/go/src/crypto/cipher/cbc.go:12
import (
//line /usr/local/go/src/crypto/cipher/cbc.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/cbc.go:12
)
//line /usr/local/go/src/crypto/cipher/cbc.go:12
import (
//line /usr/local/go/src/crypto/cipher/cbc.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/cbc.go:12
)

import (
	"bytes"
	"crypto/internal/alias"
	"crypto/subtle"
)

type cbc struct {
	b		Block
	blockSize	int
	iv		[]byte
	tmp		[]byte
}

func newCBC(b Block, iv []byte) *cbc {
//line /usr/local/go/src/crypto/cipher/cbc.go:27
	_go_fuzz_dep_.CoverTab[1478]++
							return &cbc{
		b:		b,
		blockSize:	b.BlockSize(),
		iv:		bytes.Clone(iv),
		tmp:		make([]byte, b.BlockSize()),
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:33
	// _ = "end of CoverTab[1478]"
}

type cbcEncrypter cbc

// cbcEncAble is an interface implemented by ciphers that have a specific
//line /usr/local/go/src/crypto/cipher/cbc.go:38
// optimized implementation of CBC encryption, like crypto/aes.
//line /usr/local/go/src/crypto/cipher/cbc.go:38
// NewCBCEncrypter will check for this interface and return the specific
//line /usr/local/go/src/crypto/cipher/cbc.go:38
// BlockMode if found.
//line /usr/local/go/src/crypto/cipher/cbc.go:42
type cbcEncAble interface {
	NewCBCEncrypter(iv []byte) BlockMode
}

// NewCBCEncrypter returns a BlockMode which encrypts in cipher block chaining
//line /usr/local/go/src/crypto/cipher/cbc.go:46
// mode, using the given Block. The length of iv must be the same as the
//line /usr/local/go/src/crypto/cipher/cbc.go:46
// Block's block size.
//line /usr/local/go/src/crypto/cipher/cbc.go:49
func NewCBCEncrypter(b Block, iv []byte) BlockMode {
//line /usr/local/go/src/crypto/cipher/cbc.go:49
	_go_fuzz_dep_.CoverTab[1479]++
							if len(iv) != b.BlockSize() {
//line /usr/local/go/src/crypto/cipher/cbc.go:50
		_go_fuzz_dep_.CoverTab[1482]++
								panic("cipher.NewCBCEncrypter: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/cbc.go:51
		// _ = "end of CoverTab[1482]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:52
		_go_fuzz_dep_.CoverTab[1483]++
//line /usr/local/go/src/crypto/cipher/cbc.go:52
		// _ = "end of CoverTab[1483]"
//line /usr/local/go/src/crypto/cipher/cbc.go:52
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:52
	// _ = "end of CoverTab[1479]"
//line /usr/local/go/src/crypto/cipher/cbc.go:52
	_go_fuzz_dep_.CoverTab[1480]++
							if cbc, ok := b.(cbcEncAble); ok {
//line /usr/local/go/src/crypto/cipher/cbc.go:53
		_go_fuzz_dep_.CoverTab[1484]++
								return cbc.NewCBCEncrypter(iv)
//line /usr/local/go/src/crypto/cipher/cbc.go:54
		// _ = "end of CoverTab[1484]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:55
		_go_fuzz_dep_.CoverTab[1485]++
//line /usr/local/go/src/crypto/cipher/cbc.go:55
		// _ = "end of CoverTab[1485]"
//line /usr/local/go/src/crypto/cipher/cbc.go:55
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:55
	// _ = "end of CoverTab[1480]"
//line /usr/local/go/src/crypto/cipher/cbc.go:55
	_go_fuzz_dep_.CoverTab[1481]++
							return (*cbcEncrypter)(newCBC(b, iv))
//line /usr/local/go/src/crypto/cipher/cbc.go:56
	// _ = "end of CoverTab[1481]"
}

// newCBCGenericEncrypter returns a BlockMode which encrypts in cipher block chaining
//line /usr/local/go/src/crypto/cipher/cbc.go:59
// mode, using the given Block. The length of iv must be the same as the
//line /usr/local/go/src/crypto/cipher/cbc.go:59
// Block's block size. This always returns the generic non-asm encrypter for use
//line /usr/local/go/src/crypto/cipher/cbc.go:59
// in fuzz testing.
//line /usr/local/go/src/crypto/cipher/cbc.go:63
func newCBCGenericEncrypter(b Block, iv []byte) BlockMode {
//line /usr/local/go/src/crypto/cipher/cbc.go:63
	_go_fuzz_dep_.CoverTab[1486]++
							if len(iv) != b.BlockSize() {
//line /usr/local/go/src/crypto/cipher/cbc.go:64
		_go_fuzz_dep_.CoverTab[1488]++
								panic("cipher.NewCBCEncrypter: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/cbc.go:65
		// _ = "end of CoverTab[1488]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:66
		_go_fuzz_dep_.CoverTab[1489]++
//line /usr/local/go/src/crypto/cipher/cbc.go:66
		// _ = "end of CoverTab[1489]"
//line /usr/local/go/src/crypto/cipher/cbc.go:66
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:66
	// _ = "end of CoverTab[1486]"
//line /usr/local/go/src/crypto/cipher/cbc.go:66
	_go_fuzz_dep_.CoverTab[1487]++
							return (*cbcEncrypter)(newCBC(b, iv))
//line /usr/local/go/src/crypto/cipher/cbc.go:67
	// _ = "end of CoverTab[1487]"
}

func (x *cbcEncrypter) BlockSize() int {
//line /usr/local/go/src/crypto/cipher/cbc.go:70
	_go_fuzz_dep_.CoverTab[1490]++
//line /usr/local/go/src/crypto/cipher/cbc.go:70
	return x.blockSize
//line /usr/local/go/src/crypto/cipher/cbc.go:70
	// _ = "end of CoverTab[1490]"
//line /usr/local/go/src/crypto/cipher/cbc.go:70
}

func (x *cbcEncrypter) CryptBlocks(dst, src []byte) {
//line /usr/local/go/src/crypto/cipher/cbc.go:72
	_go_fuzz_dep_.CoverTab[1491]++
							if len(src)%x.blockSize != 0 {
//line /usr/local/go/src/crypto/cipher/cbc.go:73
		_go_fuzz_dep_.CoverTab[1496]++
								panic("crypto/cipher: input not full blocks")
//line /usr/local/go/src/crypto/cipher/cbc.go:74
		// _ = "end of CoverTab[1496]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:75
		_go_fuzz_dep_.CoverTab[1497]++
//line /usr/local/go/src/crypto/cipher/cbc.go:75
		// _ = "end of CoverTab[1497]"
//line /usr/local/go/src/crypto/cipher/cbc.go:75
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:75
	// _ = "end of CoverTab[1491]"
//line /usr/local/go/src/crypto/cipher/cbc.go:75
	_go_fuzz_dep_.CoverTab[1492]++
							if len(dst) < len(src) {
//line /usr/local/go/src/crypto/cipher/cbc.go:76
		_go_fuzz_dep_.CoverTab[1498]++
								panic("crypto/cipher: output smaller than input")
//line /usr/local/go/src/crypto/cipher/cbc.go:77
		// _ = "end of CoverTab[1498]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:78
		_go_fuzz_dep_.CoverTab[1499]++
//line /usr/local/go/src/crypto/cipher/cbc.go:78
		// _ = "end of CoverTab[1499]"
//line /usr/local/go/src/crypto/cipher/cbc.go:78
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:78
	// _ = "end of CoverTab[1492]"
//line /usr/local/go/src/crypto/cipher/cbc.go:78
	_go_fuzz_dep_.CoverTab[1493]++
							if alias.InexactOverlap(dst[:len(src)], src) {
//line /usr/local/go/src/crypto/cipher/cbc.go:79
		_go_fuzz_dep_.CoverTab[1500]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/cbc.go:80
		// _ = "end of CoverTab[1500]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:81
		_go_fuzz_dep_.CoverTab[1501]++
//line /usr/local/go/src/crypto/cipher/cbc.go:81
		// _ = "end of CoverTab[1501]"
//line /usr/local/go/src/crypto/cipher/cbc.go:81
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:81
	// _ = "end of CoverTab[1493]"
//line /usr/local/go/src/crypto/cipher/cbc.go:81
	_go_fuzz_dep_.CoverTab[1494]++

							iv := x.iv

							for len(src) > 0 {
//line /usr/local/go/src/crypto/cipher/cbc.go:85
		_go_fuzz_dep_.CoverTab[1502]++

								subtle.XORBytes(dst[:x.blockSize], src[:x.blockSize], iv)
								x.b.Encrypt(dst[:x.blockSize], dst[:x.blockSize])

//line /usr/local/go/src/crypto/cipher/cbc.go:91
		iv = dst[:x.blockSize]
								src = src[x.blockSize:]
								dst = dst[x.blockSize:]
//line /usr/local/go/src/crypto/cipher/cbc.go:93
		// _ = "end of CoverTab[1502]"
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:94
	// _ = "end of CoverTab[1494]"
//line /usr/local/go/src/crypto/cipher/cbc.go:94
	_go_fuzz_dep_.CoverTab[1495]++

//line /usr/local/go/src/crypto/cipher/cbc.go:97
	copy(x.iv, iv)
//line /usr/local/go/src/crypto/cipher/cbc.go:97
	// _ = "end of CoverTab[1495]"
}

func (x *cbcEncrypter) SetIV(iv []byte) {
//line /usr/local/go/src/crypto/cipher/cbc.go:100
	_go_fuzz_dep_.CoverTab[1503]++
							if len(iv) != len(x.iv) {
//line /usr/local/go/src/crypto/cipher/cbc.go:101
		_go_fuzz_dep_.CoverTab[1505]++
								panic("cipher: incorrect length IV")
//line /usr/local/go/src/crypto/cipher/cbc.go:102
		// _ = "end of CoverTab[1505]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:103
		_go_fuzz_dep_.CoverTab[1506]++
//line /usr/local/go/src/crypto/cipher/cbc.go:103
		// _ = "end of CoverTab[1506]"
//line /usr/local/go/src/crypto/cipher/cbc.go:103
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:103
	// _ = "end of CoverTab[1503]"
//line /usr/local/go/src/crypto/cipher/cbc.go:103
	_go_fuzz_dep_.CoverTab[1504]++
							copy(x.iv, iv)
//line /usr/local/go/src/crypto/cipher/cbc.go:104
	// _ = "end of CoverTab[1504]"
}

type cbcDecrypter cbc

// cbcDecAble is an interface implemented by ciphers that have a specific
//line /usr/local/go/src/crypto/cipher/cbc.go:109
// optimized implementation of CBC decryption, like crypto/aes.
//line /usr/local/go/src/crypto/cipher/cbc.go:109
// NewCBCDecrypter will check for this interface and return the specific
//line /usr/local/go/src/crypto/cipher/cbc.go:109
// BlockMode if found.
//line /usr/local/go/src/crypto/cipher/cbc.go:113
type cbcDecAble interface {
	NewCBCDecrypter(iv []byte) BlockMode
}

// NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining
//line /usr/local/go/src/crypto/cipher/cbc.go:117
// mode, using the given Block. The length of iv must be the same as the
//line /usr/local/go/src/crypto/cipher/cbc.go:117
// Block's block size and must match the iv used to encrypt the data.
//line /usr/local/go/src/crypto/cipher/cbc.go:120
func NewCBCDecrypter(b Block, iv []byte) BlockMode {
//line /usr/local/go/src/crypto/cipher/cbc.go:120
	_go_fuzz_dep_.CoverTab[1507]++
							if len(iv) != b.BlockSize() {
//line /usr/local/go/src/crypto/cipher/cbc.go:121
		_go_fuzz_dep_.CoverTab[1510]++
								panic("cipher.NewCBCDecrypter: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/cbc.go:122
		// _ = "end of CoverTab[1510]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:123
		_go_fuzz_dep_.CoverTab[1511]++
//line /usr/local/go/src/crypto/cipher/cbc.go:123
		// _ = "end of CoverTab[1511]"
//line /usr/local/go/src/crypto/cipher/cbc.go:123
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:123
	// _ = "end of CoverTab[1507]"
//line /usr/local/go/src/crypto/cipher/cbc.go:123
	_go_fuzz_dep_.CoverTab[1508]++
							if cbc, ok := b.(cbcDecAble); ok {
//line /usr/local/go/src/crypto/cipher/cbc.go:124
		_go_fuzz_dep_.CoverTab[1512]++
								return cbc.NewCBCDecrypter(iv)
//line /usr/local/go/src/crypto/cipher/cbc.go:125
		// _ = "end of CoverTab[1512]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:126
		_go_fuzz_dep_.CoverTab[1513]++
//line /usr/local/go/src/crypto/cipher/cbc.go:126
		// _ = "end of CoverTab[1513]"
//line /usr/local/go/src/crypto/cipher/cbc.go:126
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:126
	// _ = "end of CoverTab[1508]"
//line /usr/local/go/src/crypto/cipher/cbc.go:126
	_go_fuzz_dep_.CoverTab[1509]++
							return (*cbcDecrypter)(newCBC(b, iv))
//line /usr/local/go/src/crypto/cipher/cbc.go:127
	// _ = "end of CoverTab[1509]"
}

// newCBCGenericDecrypter returns a BlockMode which encrypts in cipher block chaining
//line /usr/local/go/src/crypto/cipher/cbc.go:130
// mode, using the given Block. The length of iv must be the same as the
//line /usr/local/go/src/crypto/cipher/cbc.go:130
// Block's block size. This always returns the generic non-asm decrypter for use in
//line /usr/local/go/src/crypto/cipher/cbc.go:130
// fuzz testing.
//line /usr/local/go/src/crypto/cipher/cbc.go:134
func newCBCGenericDecrypter(b Block, iv []byte) BlockMode {
//line /usr/local/go/src/crypto/cipher/cbc.go:134
	_go_fuzz_dep_.CoverTab[1514]++
							if len(iv) != b.BlockSize() {
//line /usr/local/go/src/crypto/cipher/cbc.go:135
		_go_fuzz_dep_.CoverTab[1516]++
								panic("cipher.NewCBCDecrypter: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/cbc.go:136
		// _ = "end of CoverTab[1516]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:137
		_go_fuzz_dep_.CoverTab[1517]++
//line /usr/local/go/src/crypto/cipher/cbc.go:137
		// _ = "end of CoverTab[1517]"
//line /usr/local/go/src/crypto/cipher/cbc.go:137
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:137
	// _ = "end of CoverTab[1514]"
//line /usr/local/go/src/crypto/cipher/cbc.go:137
	_go_fuzz_dep_.CoverTab[1515]++
							return (*cbcDecrypter)(newCBC(b, iv))
//line /usr/local/go/src/crypto/cipher/cbc.go:138
	// _ = "end of CoverTab[1515]"
}

func (x *cbcDecrypter) BlockSize() int {
//line /usr/local/go/src/crypto/cipher/cbc.go:141
	_go_fuzz_dep_.CoverTab[1518]++
//line /usr/local/go/src/crypto/cipher/cbc.go:141
	return x.blockSize
//line /usr/local/go/src/crypto/cipher/cbc.go:141
	// _ = "end of CoverTab[1518]"
//line /usr/local/go/src/crypto/cipher/cbc.go:141
}

func (x *cbcDecrypter) CryptBlocks(dst, src []byte) {
//line /usr/local/go/src/crypto/cipher/cbc.go:143
	_go_fuzz_dep_.CoverTab[1519]++
							if len(src)%x.blockSize != 0 {
//line /usr/local/go/src/crypto/cipher/cbc.go:144
		_go_fuzz_dep_.CoverTab[1525]++
								panic("crypto/cipher: input not full blocks")
//line /usr/local/go/src/crypto/cipher/cbc.go:145
		// _ = "end of CoverTab[1525]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:146
		_go_fuzz_dep_.CoverTab[1526]++
//line /usr/local/go/src/crypto/cipher/cbc.go:146
		// _ = "end of CoverTab[1526]"
//line /usr/local/go/src/crypto/cipher/cbc.go:146
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:146
	// _ = "end of CoverTab[1519]"
//line /usr/local/go/src/crypto/cipher/cbc.go:146
	_go_fuzz_dep_.CoverTab[1520]++
							if len(dst) < len(src) {
//line /usr/local/go/src/crypto/cipher/cbc.go:147
		_go_fuzz_dep_.CoverTab[1527]++
								panic("crypto/cipher: output smaller than input")
//line /usr/local/go/src/crypto/cipher/cbc.go:148
		// _ = "end of CoverTab[1527]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:149
		_go_fuzz_dep_.CoverTab[1528]++
//line /usr/local/go/src/crypto/cipher/cbc.go:149
		// _ = "end of CoverTab[1528]"
//line /usr/local/go/src/crypto/cipher/cbc.go:149
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:149
	// _ = "end of CoverTab[1520]"
//line /usr/local/go/src/crypto/cipher/cbc.go:149
	_go_fuzz_dep_.CoverTab[1521]++
							if alias.InexactOverlap(dst[:len(src)], src) {
//line /usr/local/go/src/crypto/cipher/cbc.go:150
		_go_fuzz_dep_.CoverTab[1529]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/cbc.go:151
		// _ = "end of CoverTab[1529]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:152
		_go_fuzz_dep_.CoverTab[1530]++
//line /usr/local/go/src/crypto/cipher/cbc.go:152
		// _ = "end of CoverTab[1530]"
//line /usr/local/go/src/crypto/cipher/cbc.go:152
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:152
	// _ = "end of CoverTab[1521]"
//line /usr/local/go/src/crypto/cipher/cbc.go:152
	_go_fuzz_dep_.CoverTab[1522]++
							if len(src) == 0 {
//line /usr/local/go/src/crypto/cipher/cbc.go:153
		_go_fuzz_dep_.CoverTab[1531]++
								return
//line /usr/local/go/src/crypto/cipher/cbc.go:154
		// _ = "end of CoverTab[1531]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:155
		_go_fuzz_dep_.CoverTab[1532]++
//line /usr/local/go/src/crypto/cipher/cbc.go:155
		// _ = "end of CoverTab[1532]"
//line /usr/local/go/src/crypto/cipher/cbc.go:155
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:155
	// _ = "end of CoverTab[1522]"
//line /usr/local/go/src/crypto/cipher/cbc.go:155
	_go_fuzz_dep_.CoverTab[1523]++

//line /usr/local/go/src/crypto/cipher/cbc.go:159
	end := len(src)
							start := end - x.blockSize
							prev := start - x.blockSize

//line /usr/local/go/src/crypto/cipher/cbc.go:164
	copy(x.tmp, src[start:end])

//line /usr/local/go/src/crypto/cipher/cbc.go:167
	for start > 0 {
//line /usr/local/go/src/crypto/cipher/cbc.go:167
		_go_fuzz_dep_.CoverTab[1533]++
								x.b.Decrypt(dst[start:end], src[start:end])
								subtle.XORBytes(dst[start:end], dst[start:end], src[prev:start])

								end = start
								start = prev
								prev -= x.blockSize
//line /usr/local/go/src/crypto/cipher/cbc.go:173
		// _ = "end of CoverTab[1533]"
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:174
	// _ = "end of CoverTab[1523]"
//line /usr/local/go/src/crypto/cipher/cbc.go:174
	_go_fuzz_dep_.CoverTab[1524]++

//line /usr/local/go/src/crypto/cipher/cbc.go:177
	x.b.Decrypt(dst[start:end], src[start:end])
							subtle.XORBytes(dst[start:end], dst[start:end], x.iv)

//line /usr/local/go/src/crypto/cipher/cbc.go:181
	x.iv, x.tmp = x.tmp, x.iv
//line /usr/local/go/src/crypto/cipher/cbc.go:181
	// _ = "end of CoverTab[1524]"
}

func (x *cbcDecrypter) SetIV(iv []byte) {
//line /usr/local/go/src/crypto/cipher/cbc.go:184
	_go_fuzz_dep_.CoverTab[1534]++
							if len(iv) != len(x.iv) {
//line /usr/local/go/src/crypto/cipher/cbc.go:185
		_go_fuzz_dep_.CoverTab[1536]++
								panic("cipher: incorrect length IV")
//line /usr/local/go/src/crypto/cipher/cbc.go:186
		// _ = "end of CoverTab[1536]"
	} else {
//line /usr/local/go/src/crypto/cipher/cbc.go:187
		_go_fuzz_dep_.CoverTab[1537]++
//line /usr/local/go/src/crypto/cipher/cbc.go:187
		// _ = "end of CoverTab[1537]"
//line /usr/local/go/src/crypto/cipher/cbc.go:187
	}
//line /usr/local/go/src/crypto/cipher/cbc.go:187
	// _ = "end of CoverTab[1534]"
//line /usr/local/go/src/crypto/cipher/cbc.go:187
	_go_fuzz_dep_.CoverTab[1535]++
							copy(x.iv, iv)
//line /usr/local/go/src/crypto/cipher/cbc.go:188
	// _ = "end of CoverTab[1535]"
}

//line /usr/local/go/src/crypto/cipher/cbc.go:189
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/cbc.go:189
var _ = _go_fuzz_dep_.CoverTab
