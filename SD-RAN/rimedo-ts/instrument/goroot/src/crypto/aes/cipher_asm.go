// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64 || ppc64 || ppc64le

//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
package aes

//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
import (
//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
)
//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
import (
//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:7
)

import (
	"crypto/cipher"
	"crypto/internal/alias"
	"crypto/internal/boring"
	"internal/cpu"
	"internal/goarch"
)

//line /usr/local/go/src/crypto/aes/cipher_asm.go:19
//go:noescape
func encryptBlockAsm(nr int, xk *uint32, dst, src *byte)

//go:noescape
func decryptBlockAsm(nr int, xk *uint32, dst, src *byte)

//go:noescape
func expandKeyAsm(nr int, key *byte, enc *uint32, dec *uint32)

type aesCipherAsm struct {
	aesCipher
}

//line /usr/local/go/src/crypto/aes/cipher_asm.go:36
type aesCipherGCM struct {
	aesCipherAsm
}

var supportsAES = cpu.X86.HasAES || func() bool {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	_go_fuzz_dep_.CoverTab[1854]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	return cpu.ARM64.HasAES
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	// _ = "end of CoverTab[1854]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
}() || func() bool {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	_go_fuzz_dep_.CoverTab[1855]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	return goarch.IsPpc64 == 1
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	// _ = "end of CoverTab[1855]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
}() || func() bool {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	_go_fuzz_dep_.CoverTab[1856]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	return goarch.IsPpc64le == 1
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
	// _ = "end of CoverTab[1856]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:40
}()
var supportsGFMUL = cpu.X86.HasPCLMULQDQ || func() bool {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:41
	_go_fuzz_dep_.CoverTab[1857]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:41
	return cpu.ARM64.HasPMULL
//line /usr/local/go/src/crypto/aes/cipher_asm.go:41
	// _ = "end of CoverTab[1857]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:41
}()

func newCipher(key []byte) (cipher.Block, error) {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:43
	_go_fuzz_dep_.CoverTab[1858]++
							if !supportsAES {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:44
		_go_fuzz_dep_.CoverTab[1862]++
								return newCipherGeneric(key)
//line /usr/local/go/src/crypto/aes/cipher_asm.go:45
		// _ = "end of CoverTab[1862]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:46
		_go_fuzz_dep_.CoverTab[1863]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:46
		// _ = "end of CoverTab[1863]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:46
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:46
	// _ = "end of CoverTab[1858]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:46
	_go_fuzz_dep_.CoverTab[1859]++
							n := len(key) + 28
							c := aesCipherAsm{aesCipher{make([]uint32, n), make([]uint32, n)}}
							var rounds int
							switch len(key) {
	case 128 / 8:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:51
		_go_fuzz_dep_.CoverTab[1864]++
								rounds = 10
//line /usr/local/go/src/crypto/aes/cipher_asm.go:52
		// _ = "end of CoverTab[1864]"
	case 192 / 8:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:53
		_go_fuzz_dep_.CoverTab[1865]++
								rounds = 12
//line /usr/local/go/src/crypto/aes/cipher_asm.go:54
		// _ = "end of CoverTab[1865]"
	case 256 / 8:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:55
		_go_fuzz_dep_.CoverTab[1866]++
								rounds = 14
//line /usr/local/go/src/crypto/aes/cipher_asm.go:56
		// _ = "end of CoverTab[1866]"
	default:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:57
		_go_fuzz_dep_.CoverTab[1867]++
								return nil, KeySizeError(len(key))
//line /usr/local/go/src/crypto/aes/cipher_asm.go:58
		// _ = "end of CoverTab[1867]"
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:59
	// _ = "end of CoverTab[1859]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:59
	_go_fuzz_dep_.CoverTab[1860]++

							expandKeyAsm(rounds, &key[0], &c.enc[0], &c.dec[0])
							if supportsAES && func() bool {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:62
		_go_fuzz_dep_.CoverTab[1868]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:62
		return supportsGFMUL
//line /usr/local/go/src/crypto/aes/cipher_asm.go:62
		// _ = "end of CoverTab[1868]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:62
	}() {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:62
		_go_fuzz_dep_.CoverTab[1869]++
								return &aesCipherGCM{c}, nil
//line /usr/local/go/src/crypto/aes/cipher_asm.go:63
		// _ = "end of CoverTab[1869]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:64
		_go_fuzz_dep_.CoverTab[1870]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:64
		// _ = "end of CoverTab[1870]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:64
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:64
	// _ = "end of CoverTab[1860]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:64
	_go_fuzz_dep_.CoverTab[1861]++
							return &c, nil
//line /usr/local/go/src/crypto/aes/cipher_asm.go:65
	// _ = "end of CoverTab[1861]"
}

func (c *aesCipherAsm) BlockSize() int {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:68
	_go_fuzz_dep_.CoverTab[1871]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:68
	return BlockSize
//line /usr/local/go/src/crypto/aes/cipher_asm.go:68
	// _ = "end of CoverTab[1871]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:68
}

func (c *aesCipherAsm) Encrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:70
	_go_fuzz_dep_.CoverTab[1872]++
							boring.Unreachable()
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:72
		_go_fuzz_dep_.CoverTab[1876]++
								panic("crypto/aes: input not full block")
//line /usr/local/go/src/crypto/aes/cipher_asm.go:73
		// _ = "end of CoverTab[1876]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:74
		_go_fuzz_dep_.CoverTab[1877]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:74
		// _ = "end of CoverTab[1877]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:74
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:74
	// _ = "end of CoverTab[1872]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:74
	_go_fuzz_dep_.CoverTab[1873]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:75
		_go_fuzz_dep_.CoverTab[1878]++
								panic("crypto/aes: output not full block")
//line /usr/local/go/src/crypto/aes/cipher_asm.go:76
		// _ = "end of CoverTab[1878]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:77
		_go_fuzz_dep_.CoverTab[1879]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:77
		// _ = "end of CoverTab[1879]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:77
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:77
	// _ = "end of CoverTab[1873]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:77
	_go_fuzz_dep_.CoverTab[1874]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:78
		_go_fuzz_dep_.CoverTab[1880]++
								panic("crypto/aes: invalid buffer overlap")
//line /usr/local/go/src/crypto/aes/cipher_asm.go:79
		// _ = "end of CoverTab[1880]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:80
		_go_fuzz_dep_.CoverTab[1881]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:80
		// _ = "end of CoverTab[1881]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:80
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:80
	// _ = "end of CoverTab[1874]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:80
	_go_fuzz_dep_.CoverTab[1875]++
							encryptBlockAsm(len(c.enc)/4-1, &c.enc[0], &dst[0], &src[0])
//line /usr/local/go/src/crypto/aes/cipher_asm.go:81
	// _ = "end of CoverTab[1875]"
}

func (c *aesCipherAsm) Decrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:84
	_go_fuzz_dep_.CoverTab[1882]++
							boring.Unreachable()
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:86
		_go_fuzz_dep_.CoverTab[1886]++
								panic("crypto/aes: input not full block")
//line /usr/local/go/src/crypto/aes/cipher_asm.go:87
		// _ = "end of CoverTab[1886]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:88
		_go_fuzz_dep_.CoverTab[1887]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:88
		// _ = "end of CoverTab[1887]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:88
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:88
	// _ = "end of CoverTab[1882]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:88
	_go_fuzz_dep_.CoverTab[1883]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:89
		_go_fuzz_dep_.CoverTab[1888]++
								panic("crypto/aes: output not full block")
//line /usr/local/go/src/crypto/aes/cipher_asm.go:90
		// _ = "end of CoverTab[1888]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:91
		_go_fuzz_dep_.CoverTab[1889]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:91
		// _ = "end of CoverTab[1889]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:91
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:91
	// _ = "end of CoverTab[1883]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:91
	_go_fuzz_dep_.CoverTab[1884]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:92
		_go_fuzz_dep_.CoverTab[1890]++
								panic("crypto/aes: invalid buffer overlap")
//line /usr/local/go/src/crypto/aes/cipher_asm.go:93
		// _ = "end of CoverTab[1890]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:94
		_go_fuzz_dep_.CoverTab[1891]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:94
		// _ = "end of CoverTab[1891]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:94
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:94
	// _ = "end of CoverTab[1884]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:94
	_go_fuzz_dep_.CoverTab[1885]++
							decryptBlockAsm(len(c.dec)/4-1, &c.dec[0], &dst[0], &src[0])
//line /usr/local/go/src/crypto/aes/cipher_asm.go:95
	// _ = "end of CoverTab[1885]"
}

//line /usr/local/go/src/crypto/aes/cipher_asm.go:100
func expandKey(key []byte, enc, dec []uint32) {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:100
	_go_fuzz_dep_.CoverTab[1892]++
							if supportsAES {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:101
		_go_fuzz_dep_.CoverTab[1893]++
								rounds := 10
								switch len(key) {
		case 192 / 8:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:104
			_go_fuzz_dep_.CoverTab[1895]++
									rounds = 12
//line /usr/local/go/src/crypto/aes/cipher_asm.go:105
			// _ = "end of CoverTab[1895]"
		case 256 / 8:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:106
			_go_fuzz_dep_.CoverTab[1896]++
									rounds = 14
//line /usr/local/go/src/crypto/aes/cipher_asm.go:107
			// _ = "end of CoverTab[1896]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:107
		default:
//line /usr/local/go/src/crypto/aes/cipher_asm.go:107
			_go_fuzz_dep_.CoverTab[1897]++
//line /usr/local/go/src/crypto/aes/cipher_asm.go:107
			// _ = "end of CoverTab[1897]"
		}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:108
		// _ = "end of CoverTab[1893]"
//line /usr/local/go/src/crypto/aes/cipher_asm.go:108
		_go_fuzz_dep_.CoverTab[1894]++
								expandKeyAsm(rounds, &key[0], &enc[0], &dec[0])
//line /usr/local/go/src/crypto/aes/cipher_asm.go:109
		// _ = "end of CoverTab[1894]"
	} else {
//line /usr/local/go/src/crypto/aes/cipher_asm.go:110
		_go_fuzz_dep_.CoverTab[1898]++
								expandKeyGo(key, enc, dec)
//line /usr/local/go/src/crypto/aes/cipher_asm.go:111
		// _ = "end of CoverTab[1898]"
	}
//line /usr/local/go/src/crypto/aes/cipher_asm.go:112
	// _ = "end of CoverTab[1892]"
}

//line /usr/local/go/src/crypto/aes/cipher_asm.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/aes/cipher_asm.go:113
var _ = _go_fuzz_dep_.CoverTab
