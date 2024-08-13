// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/aes/cipher.go:5
package aes

//line /usr/local/go/src/crypto/aes/cipher.go:5
import (
//line /usr/local/go/src/crypto/aes/cipher.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/aes/cipher.go:5
)
//line /usr/local/go/src/crypto/aes/cipher.go:5
import (
//line /usr/local/go/src/crypto/aes/cipher.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/aes/cipher.go:5
)

import (
	"crypto/cipher"
	"crypto/internal/alias"
	"crypto/internal/boring"
	"strconv"
)

// The AES block size in bytes.
const BlockSize = 16

// A cipher is an instance of AES encryption using a particular key.
type aesCipher struct {
	enc	[]uint32
	dec	[]uint32
}

type KeySizeError int

func (k KeySizeError) Error() string {
//line /usr/local/go/src/crypto/aes/cipher.go:25
	_go_fuzz_dep_.CoverTab[1824]++
							return "crypto/aes: invalid key size " + strconv.Itoa(int(k))
//line /usr/local/go/src/crypto/aes/cipher.go:26
	// _ = "end of CoverTab[1824]"
}

// NewCipher creates and returns a new cipher.Block.
//line /usr/local/go/src/crypto/aes/cipher.go:29
// The key argument should be the AES key,
//line /usr/local/go/src/crypto/aes/cipher.go:29
// either 16, 24, or 32 bytes to select
//line /usr/local/go/src/crypto/aes/cipher.go:29
// AES-128, AES-192, or AES-256.
//line /usr/local/go/src/crypto/aes/cipher.go:33
func NewCipher(key []byte) (cipher.Block, error) {
//line /usr/local/go/src/crypto/aes/cipher.go:33
	_go_fuzz_dep_.CoverTab[1825]++
							k := len(key)
							switch k {
	default:
//line /usr/local/go/src/crypto/aes/cipher.go:36
		_go_fuzz_dep_.CoverTab[1828]++
								return nil, KeySizeError(k)
//line /usr/local/go/src/crypto/aes/cipher.go:37
		// _ = "end of CoverTab[1828]"
	case 16, 24, 32:
//line /usr/local/go/src/crypto/aes/cipher.go:38
		_go_fuzz_dep_.CoverTab[1829]++
								break
//line /usr/local/go/src/crypto/aes/cipher.go:39
		// _ = "end of CoverTab[1829]"
	}
//line /usr/local/go/src/crypto/aes/cipher.go:40
	// _ = "end of CoverTab[1825]"
//line /usr/local/go/src/crypto/aes/cipher.go:40
	_go_fuzz_dep_.CoverTab[1826]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/aes/cipher.go:41
		_go_fuzz_dep_.CoverTab[1830]++
								return boring.NewAESCipher(key)
//line /usr/local/go/src/crypto/aes/cipher.go:42
		// _ = "end of CoverTab[1830]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:43
		_go_fuzz_dep_.CoverTab[1831]++
//line /usr/local/go/src/crypto/aes/cipher.go:43
		// _ = "end of CoverTab[1831]"
//line /usr/local/go/src/crypto/aes/cipher.go:43
	}
//line /usr/local/go/src/crypto/aes/cipher.go:43
	// _ = "end of CoverTab[1826]"
//line /usr/local/go/src/crypto/aes/cipher.go:43
	_go_fuzz_dep_.CoverTab[1827]++
							return newCipher(key)
//line /usr/local/go/src/crypto/aes/cipher.go:44
	// _ = "end of CoverTab[1827]"
}

// newCipherGeneric creates and returns a new cipher.Block
//line /usr/local/go/src/crypto/aes/cipher.go:47
// implemented in pure Go.
//line /usr/local/go/src/crypto/aes/cipher.go:49
func newCipherGeneric(key []byte) (cipher.Block, error) {
//line /usr/local/go/src/crypto/aes/cipher.go:49
	_go_fuzz_dep_.CoverTab[1832]++
							n := len(key) + 28
							c := aesCipher{make([]uint32, n), make([]uint32, n)}
							expandKeyGo(key, c.enc, c.dec)
							return &c, nil
//line /usr/local/go/src/crypto/aes/cipher.go:53
	// _ = "end of CoverTab[1832]"
}

func (c *aesCipher) BlockSize() int {
//line /usr/local/go/src/crypto/aes/cipher.go:56
	_go_fuzz_dep_.CoverTab[1833]++
//line /usr/local/go/src/crypto/aes/cipher.go:56
	return BlockSize
//line /usr/local/go/src/crypto/aes/cipher.go:56
	// _ = "end of CoverTab[1833]"
//line /usr/local/go/src/crypto/aes/cipher.go:56
}

func (c *aesCipher) Encrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/aes/cipher.go:58
	_go_fuzz_dep_.CoverTab[1834]++
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher.go:59
		_go_fuzz_dep_.CoverTab[1838]++
								panic("crypto/aes: input not full block")
//line /usr/local/go/src/crypto/aes/cipher.go:60
		// _ = "end of CoverTab[1838]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:61
		_go_fuzz_dep_.CoverTab[1839]++
//line /usr/local/go/src/crypto/aes/cipher.go:61
		// _ = "end of CoverTab[1839]"
//line /usr/local/go/src/crypto/aes/cipher.go:61
	}
//line /usr/local/go/src/crypto/aes/cipher.go:61
	// _ = "end of CoverTab[1834]"
//line /usr/local/go/src/crypto/aes/cipher.go:61
	_go_fuzz_dep_.CoverTab[1835]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher.go:62
		_go_fuzz_dep_.CoverTab[1840]++
								panic("crypto/aes: output not full block")
//line /usr/local/go/src/crypto/aes/cipher.go:63
		// _ = "end of CoverTab[1840]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:64
		_go_fuzz_dep_.CoverTab[1841]++
//line /usr/local/go/src/crypto/aes/cipher.go:64
		// _ = "end of CoverTab[1841]"
//line /usr/local/go/src/crypto/aes/cipher.go:64
	}
//line /usr/local/go/src/crypto/aes/cipher.go:64
	// _ = "end of CoverTab[1835]"
//line /usr/local/go/src/crypto/aes/cipher.go:64
	_go_fuzz_dep_.CoverTab[1836]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/aes/cipher.go:65
		_go_fuzz_dep_.CoverTab[1842]++
								panic("crypto/aes: invalid buffer overlap")
//line /usr/local/go/src/crypto/aes/cipher.go:66
		// _ = "end of CoverTab[1842]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:67
		_go_fuzz_dep_.CoverTab[1843]++
//line /usr/local/go/src/crypto/aes/cipher.go:67
		// _ = "end of CoverTab[1843]"
//line /usr/local/go/src/crypto/aes/cipher.go:67
	}
//line /usr/local/go/src/crypto/aes/cipher.go:67
	// _ = "end of CoverTab[1836]"
//line /usr/local/go/src/crypto/aes/cipher.go:67
	_go_fuzz_dep_.CoverTab[1837]++
							encryptBlockGo(c.enc, dst, src)
//line /usr/local/go/src/crypto/aes/cipher.go:68
	// _ = "end of CoverTab[1837]"
}

func (c *aesCipher) Decrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/aes/cipher.go:71
	_go_fuzz_dep_.CoverTab[1844]++
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher.go:72
		_go_fuzz_dep_.CoverTab[1848]++
								panic("crypto/aes: input not full block")
//line /usr/local/go/src/crypto/aes/cipher.go:73
		// _ = "end of CoverTab[1848]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:74
		_go_fuzz_dep_.CoverTab[1849]++
//line /usr/local/go/src/crypto/aes/cipher.go:74
		// _ = "end of CoverTab[1849]"
//line /usr/local/go/src/crypto/aes/cipher.go:74
	}
//line /usr/local/go/src/crypto/aes/cipher.go:74
	// _ = "end of CoverTab[1844]"
//line /usr/local/go/src/crypto/aes/cipher.go:74
	_go_fuzz_dep_.CoverTab[1845]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher.go:75
		_go_fuzz_dep_.CoverTab[1850]++
								panic("crypto/aes: output not full block")
//line /usr/local/go/src/crypto/aes/cipher.go:76
		// _ = "end of CoverTab[1850]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:77
		_go_fuzz_dep_.CoverTab[1851]++
//line /usr/local/go/src/crypto/aes/cipher.go:77
		// _ = "end of CoverTab[1851]"
//line /usr/local/go/src/crypto/aes/cipher.go:77
	}
//line /usr/local/go/src/crypto/aes/cipher.go:77
	// _ = "end of CoverTab[1845]"
//line /usr/local/go/src/crypto/aes/cipher.go:77
	_go_fuzz_dep_.CoverTab[1846]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/aes/cipher.go:78
		_go_fuzz_dep_.CoverTab[1852]++
								panic("crypto/aes: invalid buffer overlap")
//line /usr/local/go/src/crypto/aes/cipher.go:79
		// _ = "end of CoverTab[1852]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher.go:80
		_go_fuzz_dep_.CoverTab[1853]++
//line /usr/local/go/src/crypto/aes/cipher.go:80
		// _ = "end of CoverTab[1853]"
//line /usr/local/go/src/crypto/aes/cipher.go:80
	}
//line /usr/local/go/src/crypto/aes/cipher.go:80
	// _ = "end of CoverTab[1846]"
//line /usr/local/go/src/crypto/aes/cipher.go:80
	_go_fuzz_dep_.CoverTab[1847]++
							decryptBlockGo(c.dec, dst, src)
//line /usr/local/go/src/crypto/aes/cipher.go:81
	// _ = "end of CoverTab[1847]"
}

//line /usr/local/go/src/crypto/aes/cipher.go:82
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/aes/cipher.go:82
var _ = _go_fuzz_dep_.CoverTab
