// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/des/cipher.go:5
package des

//line /usr/local/go/src/crypto/des/cipher.go:5
import (
//line /usr/local/go/src/crypto/des/cipher.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/des/cipher.go:5
)
//line /usr/local/go/src/crypto/des/cipher.go:5
import (
//line /usr/local/go/src/crypto/des/cipher.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/des/cipher.go:5
)

import (
	"crypto/cipher"
	"crypto/internal/alias"
	"encoding/binary"
	"strconv"
)

// The DES block size in bytes.
const BlockSize = 8

type KeySizeError int

func (k KeySizeError) Error() string {
//line /usr/local/go/src/crypto/des/cipher.go:19
	_go_fuzz_dep_.CoverTab[1923]++
							return "crypto/des: invalid key size " + strconv.Itoa(int(k))
//line /usr/local/go/src/crypto/des/cipher.go:20
	// _ = "end of CoverTab[1923]"
}

// desCipher is an instance of DES encryption.
type desCipher struct {
	subkeys [16]uint64
}

// NewCipher creates and returns a new cipher.Block.
func NewCipher(key []byte) (cipher.Block, error) {
//line /usr/local/go/src/crypto/des/cipher.go:29
	_go_fuzz_dep_.CoverTab[1924]++
							if len(key) != 8 {
//line /usr/local/go/src/crypto/des/cipher.go:30
		_go_fuzz_dep_.CoverTab[1926]++
								return nil, KeySizeError(len(key))
//line /usr/local/go/src/crypto/des/cipher.go:31
		// _ = "end of CoverTab[1926]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:32
		_go_fuzz_dep_.CoverTab[1927]++
//line /usr/local/go/src/crypto/des/cipher.go:32
		// _ = "end of CoverTab[1927]"
//line /usr/local/go/src/crypto/des/cipher.go:32
	}
//line /usr/local/go/src/crypto/des/cipher.go:32
	// _ = "end of CoverTab[1924]"
//line /usr/local/go/src/crypto/des/cipher.go:32
	_go_fuzz_dep_.CoverTab[1925]++

							c := new(desCipher)
							c.generateSubkeys(key)
							return c, nil
//line /usr/local/go/src/crypto/des/cipher.go:36
	// _ = "end of CoverTab[1925]"
}

func (c *desCipher) BlockSize() int {
//line /usr/local/go/src/crypto/des/cipher.go:39
	_go_fuzz_dep_.CoverTab[1928]++
//line /usr/local/go/src/crypto/des/cipher.go:39
	return BlockSize
//line /usr/local/go/src/crypto/des/cipher.go:39
	// _ = "end of CoverTab[1928]"
//line /usr/local/go/src/crypto/des/cipher.go:39
}

func (c *desCipher) Encrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/des/cipher.go:41
	_go_fuzz_dep_.CoverTab[1929]++
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:42
		_go_fuzz_dep_.CoverTab[1933]++
								panic("crypto/des: input not full block")
//line /usr/local/go/src/crypto/des/cipher.go:43
		// _ = "end of CoverTab[1933]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:44
		_go_fuzz_dep_.CoverTab[1934]++
//line /usr/local/go/src/crypto/des/cipher.go:44
		// _ = "end of CoverTab[1934]"
//line /usr/local/go/src/crypto/des/cipher.go:44
	}
//line /usr/local/go/src/crypto/des/cipher.go:44
	// _ = "end of CoverTab[1929]"
//line /usr/local/go/src/crypto/des/cipher.go:44
	_go_fuzz_dep_.CoverTab[1930]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:45
		_go_fuzz_dep_.CoverTab[1935]++
								panic("crypto/des: output not full block")
//line /usr/local/go/src/crypto/des/cipher.go:46
		// _ = "end of CoverTab[1935]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:47
		_go_fuzz_dep_.CoverTab[1936]++
//line /usr/local/go/src/crypto/des/cipher.go:47
		// _ = "end of CoverTab[1936]"
//line /usr/local/go/src/crypto/des/cipher.go:47
	}
//line /usr/local/go/src/crypto/des/cipher.go:47
	// _ = "end of CoverTab[1930]"
//line /usr/local/go/src/crypto/des/cipher.go:47
	_go_fuzz_dep_.CoverTab[1931]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/des/cipher.go:48
		_go_fuzz_dep_.CoverTab[1937]++
								panic("crypto/des: invalid buffer overlap")
//line /usr/local/go/src/crypto/des/cipher.go:49
		// _ = "end of CoverTab[1937]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:50
		_go_fuzz_dep_.CoverTab[1938]++
//line /usr/local/go/src/crypto/des/cipher.go:50
		// _ = "end of CoverTab[1938]"
//line /usr/local/go/src/crypto/des/cipher.go:50
	}
//line /usr/local/go/src/crypto/des/cipher.go:50
	// _ = "end of CoverTab[1931]"
//line /usr/local/go/src/crypto/des/cipher.go:50
	_go_fuzz_dep_.CoverTab[1932]++
							encryptBlock(c.subkeys[:], dst, src)
//line /usr/local/go/src/crypto/des/cipher.go:51
	// _ = "end of CoverTab[1932]"
}

func (c *desCipher) Decrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/des/cipher.go:54
	_go_fuzz_dep_.CoverTab[1939]++
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:55
		_go_fuzz_dep_.CoverTab[1943]++
								panic("crypto/des: input not full block")
//line /usr/local/go/src/crypto/des/cipher.go:56
		// _ = "end of CoverTab[1943]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:57
		_go_fuzz_dep_.CoverTab[1944]++
//line /usr/local/go/src/crypto/des/cipher.go:57
		// _ = "end of CoverTab[1944]"
//line /usr/local/go/src/crypto/des/cipher.go:57
	}
//line /usr/local/go/src/crypto/des/cipher.go:57
	// _ = "end of CoverTab[1939]"
//line /usr/local/go/src/crypto/des/cipher.go:57
	_go_fuzz_dep_.CoverTab[1940]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:58
		_go_fuzz_dep_.CoverTab[1945]++
								panic("crypto/des: output not full block")
//line /usr/local/go/src/crypto/des/cipher.go:59
		// _ = "end of CoverTab[1945]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:60
		_go_fuzz_dep_.CoverTab[1946]++
//line /usr/local/go/src/crypto/des/cipher.go:60
		// _ = "end of CoverTab[1946]"
//line /usr/local/go/src/crypto/des/cipher.go:60
	}
//line /usr/local/go/src/crypto/des/cipher.go:60
	// _ = "end of CoverTab[1940]"
//line /usr/local/go/src/crypto/des/cipher.go:60
	_go_fuzz_dep_.CoverTab[1941]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/des/cipher.go:61
		_go_fuzz_dep_.CoverTab[1947]++
								panic("crypto/des: invalid buffer overlap")
//line /usr/local/go/src/crypto/des/cipher.go:62
		// _ = "end of CoverTab[1947]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:63
		_go_fuzz_dep_.CoverTab[1948]++
//line /usr/local/go/src/crypto/des/cipher.go:63
		// _ = "end of CoverTab[1948]"
//line /usr/local/go/src/crypto/des/cipher.go:63
	}
//line /usr/local/go/src/crypto/des/cipher.go:63
	// _ = "end of CoverTab[1941]"
//line /usr/local/go/src/crypto/des/cipher.go:63
	_go_fuzz_dep_.CoverTab[1942]++
							decryptBlock(c.subkeys[:], dst, src)
//line /usr/local/go/src/crypto/des/cipher.go:64
	// _ = "end of CoverTab[1942]"
}

// A tripleDESCipher is an instance of TripleDES encryption.
type tripleDESCipher struct {
	cipher1, cipher2, cipher3 desCipher
}

// NewTripleDESCipher creates and returns a new cipher.Block.
func NewTripleDESCipher(key []byte) (cipher.Block, error) {
//line /usr/local/go/src/crypto/des/cipher.go:73
	_go_fuzz_dep_.CoverTab[1949]++
							if len(key) != 24 {
//line /usr/local/go/src/crypto/des/cipher.go:74
		_go_fuzz_dep_.CoverTab[1951]++
								return nil, KeySizeError(len(key))
//line /usr/local/go/src/crypto/des/cipher.go:75
		// _ = "end of CoverTab[1951]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:76
		_go_fuzz_dep_.CoverTab[1952]++
//line /usr/local/go/src/crypto/des/cipher.go:76
		// _ = "end of CoverTab[1952]"
//line /usr/local/go/src/crypto/des/cipher.go:76
	}
//line /usr/local/go/src/crypto/des/cipher.go:76
	// _ = "end of CoverTab[1949]"
//line /usr/local/go/src/crypto/des/cipher.go:76
	_go_fuzz_dep_.CoverTab[1950]++

							c := new(tripleDESCipher)
							c.cipher1.generateSubkeys(key[:8])
							c.cipher2.generateSubkeys(key[8:16])
							c.cipher3.generateSubkeys(key[16:])
							return c, nil
//line /usr/local/go/src/crypto/des/cipher.go:82
	// _ = "end of CoverTab[1950]"
}

func (c *tripleDESCipher) BlockSize() int {
//line /usr/local/go/src/crypto/des/cipher.go:85
	_go_fuzz_dep_.CoverTab[1953]++
//line /usr/local/go/src/crypto/des/cipher.go:85
	return BlockSize
//line /usr/local/go/src/crypto/des/cipher.go:85
	// _ = "end of CoverTab[1953]"
//line /usr/local/go/src/crypto/des/cipher.go:85
}

func (c *tripleDESCipher) Encrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/des/cipher.go:87
	_go_fuzz_dep_.CoverTab[1954]++
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:88
		_go_fuzz_dep_.CoverTab[1961]++
								panic("crypto/des: input not full block")
//line /usr/local/go/src/crypto/des/cipher.go:89
		// _ = "end of CoverTab[1961]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:90
		_go_fuzz_dep_.CoverTab[1962]++
//line /usr/local/go/src/crypto/des/cipher.go:90
		// _ = "end of CoverTab[1962]"
//line /usr/local/go/src/crypto/des/cipher.go:90
	}
//line /usr/local/go/src/crypto/des/cipher.go:90
	// _ = "end of CoverTab[1954]"
//line /usr/local/go/src/crypto/des/cipher.go:90
	_go_fuzz_dep_.CoverTab[1955]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:91
		_go_fuzz_dep_.CoverTab[1963]++
								panic("crypto/des: output not full block")
//line /usr/local/go/src/crypto/des/cipher.go:92
		// _ = "end of CoverTab[1963]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:93
		_go_fuzz_dep_.CoverTab[1964]++
//line /usr/local/go/src/crypto/des/cipher.go:93
		// _ = "end of CoverTab[1964]"
//line /usr/local/go/src/crypto/des/cipher.go:93
	}
//line /usr/local/go/src/crypto/des/cipher.go:93
	// _ = "end of CoverTab[1955]"
//line /usr/local/go/src/crypto/des/cipher.go:93
	_go_fuzz_dep_.CoverTab[1956]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/des/cipher.go:94
		_go_fuzz_dep_.CoverTab[1965]++
								panic("crypto/des: invalid buffer overlap")
//line /usr/local/go/src/crypto/des/cipher.go:95
		// _ = "end of CoverTab[1965]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:96
		_go_fuzz_dep_.CoverTab[1966]++
//line /usr/local/go/src/crypto/des/cipher.go:96
		// _ = "end of CoverTab[1966]"
//line /usr/local/go/src/crypto/des/cipher.go:96
	}
//line /usr/local/go/src/crypto/des/cipher.go:96
	// _ = "end of CoverTab[1956]"
//line /usr/local/go/src/crypto/des/cipher.go:96
	_go_fuzz_dep_.CoverTab[1957]++

							b := binary.BigEndian.Uint64(src)
							b = permuteInitialBlock(b)
							left, right := uint32(b>>32), uint32(b)

							left = (left << 1) | (left >> 31)
							right = (right << 1) | (right >> 31)

							for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/cipher.go:105
		_go_fuzz_dep_.CoverTab[1967]++
								left, right = feistel(left, right, c.cipher1.subkeys[2*i], c.cipher1.subkeys[2*i+1])
//line /usr/local/go/src/crypto/des/cipher.go:106
		// _ = "end of CoverTab[1967]"
	}
//line /usr/local/go/src/crypto/des/cipher.go:107
	// _ = "end of CoverTab[1957]"
//line /usr/local/go/src/crypto/des/cipher.go:107
	_go_fuzz_dep_.CoverTab[1958]++
							for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/cipher.go:108
		_go_fuzz_dep_.CoverTab[1968]++
								right, left = feistel(right, left, c.cipher2.subkeys[15-2*i], c.cipher2.subkeys[15-(2*i+1)])
//line /usr/local/go/src/crypto/des/cipher.go:109
		// _ = "end of CoverTab[1968]"
	}
//line /usr/local/go/src/crypto/des/cipher.go:110
	// _ = "end of CoverTab[1958]"
//line /usr/local/go/src/crypto/des/cipher.go:110
	_go_fuzz_dep_.CoverTab[1959]++
							for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/cipher.go:111
		_go_fuzz_dep_.CoverTab[1969]++
								left, right = feistel(left, right, c.cipher3.subkeys[2*i], c.cipher3.subkeys[2*i+1])
//line /usr/local/go/src/crypto/des/cipher.go:112
		// _ = "end of CoverTab[1969]"
	}
//line /usr/local/go/src/crypto/des/cipher.go:113
	// _ = "end of CoverTab[1959]"
//line /usr/local/go/src/crypto/des/cipher.go:113
	_go_fuzz_dep_.CoverTab[1960]++

							left = (left << 31) | (left >> 1)
							right = (right << 31) | (right >> 1)

							preOutput := (uint64(right) << 32) | uint64(left)
							binary.BigEndian.PutUint64(dst, permuteFinalBlock(preOutput))
//line /usr/local/go/src/crypto/des/cipher.go:119
	// _ = "end of CoverTab[1960]"
}

func (c *tripleDESCipher) Decrypt(dst, src []byte) {
//line /usr/local/go/src/crypto/des/cipher.go:122
	_go_fuzz_dep_.CoverTab[1970]++
							if len(src) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:123
		_go_fuzz_dep_.CoverTab[1977]++
								panic("crypto/des: input not full block")
//line /usr/local/go/src/crypto/des/cipher.go:124
		// _ = "end of CoverTab[1977]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:125
		_go_fuzz_dep_.CoverTab[1978]++
//line /usr/local/go/src/crypto/des/cipher.go:125
		// _ = "end of CoverTab[1978]"
//line /usr/local/go/src/crypto/des/cipher.go:125
	}
//line /usr/local/go/src/crypto/des/cipher.go:125
	// _ = "end of CoverTab[1970]"
//line /usr/local/go/src/crypto/des/cipher.go:125
	_go_fuzz_dep_.CoverTab[1971]++
							if len(dst) < BlockSize {
//line /usr/local/go/src/crypto/des/cipher.go:126
		_go_fuzz_dep_.CoverTab[1979]++
								panic("crypto/des: output not full block")
//line /usr/local/go/src/crypto/des/cipher.go:127
		// _ = "end of CoverTab[1979]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:128
		_go_fuzz_dep_.CoverTab[1980]++
//line /usr/local/go/src/crypto/des/cipher.go:128
		// _ = "end of CoverTab[1980]"
//line /usr/local/go/src/crypto/des/cipher.go:128
	}
//line /usr/local/go/src/crypto/des/cipher.go:128
	// _ = "end of CoverTab[1971]"
//line /usr/local/go/src/crypto/des/cipher.go:128
	_go_fuzz_dep_.CoverTab[1972]++
							if alias.InexactOverlap(dst[:BlockSize], src[:BlockSize]) {
//line /usr/local/go/src/crypto/des/cipher.go:129
		_go_fuzz_dep_.CoverTab[1981]++
								panic("crypto/des: invalid buffer overlap")
//line /usr/local/go/src/crypto/des/cipher.go:130
		// _ = "end of CoverTab[1981]"
	} else {
//line /usr/local/go/src/crypto/des/cipher.go:131
		_go_fuzz_dep_.CoverTab[1982]++
//line /usr/local/go/src/crypto/des/cipher.go:131
		// _ = "end of CoverTab[1982]"
//line /usr/local/go/src/crypto/des/cipher.go:131
	}
//line /usr/local/go/src/crypto/des/cipher.go:131
	// _ = "end of CoverTab[1972]"
//line /usr/local/go/src/crypto/des/cipher.go:131
	_go_fuzz_dep_.CoverTab[1973]++

							b := binary.BigEndian.Uint64(src)
							b = permuteInitialBlock(b)
							left, right := uint32(b>>32), uint32(b)

							left = (left << 1) | (left >> 31)
							right = (right << 1) | (right >> 31)

							for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/cipher.go:140
		_go_fuzz_dep_.CoverTab[1983]++
								left, right = feistel(left, right, c.cipher3.subkeys[15-2*i], c.cipher3.subkeys[15-(2*i+1)])
//line /usr/local/go/src/crypto/des/cipher.go:141
		// _ = "end of CoverTab[1983]"
	}
//line /usr/local/go/src/crypto/des/cipher.go:142
	// _ = "end of CoverTab[1973]"
//line /usr/local/go/src/crypto/des/cipher.go:142
	_go_fuzz_dep_.CoverTab[1974]++
							for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/cipher.go:143
		_go_fuzz_dep_.CoverTab[1984]++
								right, left = feistel(right, left, c.cipher2.subkeys[2*i], c.cipher2.subkeys[2*i+1])
//line /usr/local/go/src/crypto/des/cipher.go:144
		// _ = "end of CoverTab[1984]"
	}
//line /usr/local/go/src/crypto/des/cipher.go:145
	// _ = "end of CoverTab[1974]"
//line /usr/local/go/src/crypto/des/cipher.go:145
	_go_fuzz_dep_.CoverTab[1975]++
							for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/cipher.go:146
		_go_fuzz_dep_.CoverTab[1985]++
								left, right = feistel(left, right, c.cipher1.subkeys[15-2*i], c.cipher1.subkeys[15-(2*i+1)])
//line /usr/local/go/src/crypto/des/cipher.go:147
		// _ = "end of CoverTab[1985]"
	}
//line /usr/local/go/src/crypto/des/cipher.go:148
	// _ = "end of CoverTab[1975]"
//line /usr/local/go/src/crypto/des/cipher.go:148
	_go_fuzz_dep_.CoverTab[1976]++

							left = (left << 31) | (left >> 1)
							right = (right << 31) | (right >> 1)

							preOutput := (uint64(right) << 32) | uint64(left)
							binary.BigEndian.PutUint64(dst, permuteFinalBlock(preOutput))
//line /usr/local/go/src/crypto/des/cipher.go:154
	// _ = "end of CoverTab[1976]"
}

//line /usr/local/go/src/crypto/des/cipher.go:155
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/des/cipher.go:155
var _ = _go_fuzz_dep_.CoverTab
