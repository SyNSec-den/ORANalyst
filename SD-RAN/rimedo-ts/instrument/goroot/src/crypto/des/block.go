// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/des/block.go:5
package des

//line /usr/local/go/src/crypto/des/block.go:5
import (
//line /usr/local/go/src/crypto/des/block.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/des/block.go:5
)
//line /usr/local/go/src/crypto/des/block.go:5
import (
//line /usr/local/go/src/crypto/des/block.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/des/block.go:5
)

import (
	"encoding/binary"
	"sync"
)

func cryptBlock(subkeys []uint64, dst, src []byte, decrypt bool) {
//line /usr/local/go/src/crypto/des/block.go:12
	_go_fuzz_dep_.CoverTab[1899]++
							b := binary.BigEndian.Uint64(src)
							b = permuteInitialBlock(b)
							left, right := uint32(b>>32), uint32(b)

							left = (left << 1) | (left >> 31)
							right = (right << 1) | (right >> 31)

							if decrypt {
//line /usr/local/go/src/crypto/des/block.go:20
		_go_fuzz_dep_.CoverTab[1901]++
								for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/block.go:21
			_go_fuzz_dep_.CoverTab[1902]++
									left, right = feistel(left, right, subkeys[15-2*i], subkeys[15-(2*i+1)])
//line /usr/local/go/src/crypto/des/block.go:22
			// _ = "end of CoverTab[1902]"
		}
//line /usr/local/go/src/crypto/des/block.go:23
		// _ = "end of CoverTab[1901]"
	} else {
//line /usr/local/go/src/crypto/des/block.go:24
		_go_fuzz_dep_.CoverTab[1903]++
								for i := 0; i < 8; i++ {
//line /usr/local/go/src/crypto/des/block.go:25
			_go_fuzz_dep_.CoverTab[1904]++
									left, right = feistel(left, right, subkeys[2*i], subkeys[2*i+1])
//line /usr/local/go/src/crypto/des/block.go:26
			// _ = "end of CoverTab[1904]"
		}
//line /usr/local/go/src/crypto/des/block.go:27
		// _ = "end of CoverTab[1903]"
	}
//line /usr/local/go/src/crypto/des/block.go:28
	// _ = "end of CoverTab[1899]"
//line /usr/local/go/src/crypto/des/block.go:28
	_go_fuzz_dep_.CoverTab[1900]++

							left = (left << 31) | (left >> 1)
							right = (right << 31) | (right >> 1)

//line /usr/local/go/src/crypto/des/block.go:34
	preOutput := (uint64(right) << 32) | uint64(left)
							binary.BigEndian.PutUint64(dst, permuteFinalBlock(preOutput))
//line /usr/local/go/src/crypto/des/block.go:35
	// _ = "end of CoverTab[1900]"
}

// Encrypt one block from src into dst, using the subkeys.
func encryptBlock(subkeys []uint64, dst, src []byte) {
//line /usr/local/go/src/crypto/des/block.go:39
	_go_fuzz_dep_.CoverTab[1905]++
							cryptBlock(subkeys, dst, src, false)
//line /usr/local/go/src/crypto/des/block.go:40
	// _ = "end of CoverTab[1905]"
}

// Decrypt one block from src into dst, using the subkeys.
func decryptBlock(subkeys []uint64, dst, src []byte) {
//line /usr/local/go/src/crypto/des/block.go:44
	_go_fuzz_dep_.CoverTab[1906]++
							cryptBlock(subkeys, dst, src, true)
//line /usr/local/go/src/crypto/des/block.go:45
	// _ = "end of CoverTab[1906]"
}

// DES Feistel function. feistelBox must be initialized via
//line /usr/local/go/src/crypto/des/block.go:48
// feistelBoxOnce.Do(initFeistelBox) first.
//line /usr/local/go/src/crypto/des/block.go:50
func feistel(l, r uint32, k0, k1 uint64) (lout, rout uint32) {
//line /usr/local/go/src/crypto/des/block.go:50
	_go_fuzz_dep_.CoverTab[1907]++
							var t uint32

							t = r ^ uint32(k0>>32)
							l ^= feistelBox[7][t&0x3f] ^
		feistelBox[5][(t>>8)&0x3f] ^
		feistelBox[3][(t>>16)&0x3f] ^
		feistelBox[1][(t>>24)&0x3f]

	t = ((r << 28) | (r >> 4)) ^ uint32(k0)
	l ^= feistelBox[6][(t)&0x3f] ^
		feistelBox[4][(t>>8)&0x3f] ^
		feistelBox[2][(t>>16)&0x3f] ^
		feistelBox[0][(t>>24)&0x3f]

	t = l ^ uint32(k1>>32)
	r ^= feistelBox[7][t&0x3f] ^
		feistelBox[5][(t>>8)&0x3f] ^
		feistelBox[3][(t>>16)&0x3f] ^
		feistelBox[1][(t>>24)&0x3f]

	t = ((l << 28) | (l >> 4)) ^ uint32(k1)
	r ^= feistelBox[6][(t)&0x3f] ^
		feistelBox[4][(t>>8)&0x3f] ^
		feistelBox[2][(t>>16)&0x3f] ^
		feistelBox[0][(t>>24)&0x3f]

							return l, r
//line /usr/local/go/src/crypto/des/block.go:77
	// _ = "end of CoverTab[1907]"
}

// feistelBox[s][16*i+j] contains the output of permutationFunction
//line /usr/local/go/src/crypto/des/block.go:80
// for sBoxes[s][i][j] << 4*(7-s)
//line /usr/local/go/src/crypto/des/block.go:82
var feistelBox [8][64]uint32

var feistelBoxOnce sync.Once

// general purpose function to perform DES block permutations.
func permuteBlock(src uint64, permutation []uint8) (block uint64) {
//line /usr/local/go/src/crypto/des/block.go:87
	_go_fuzz_dep_.CoverTab[1908]++
							for position, n := range permutation {
//line /usr/local/go/src/crypto/des/block.go:88
		_go_fuzz_dep_.CoverTab[1910]++
								bit := (src >> n) & 1
								block |= bit << uint((len(permutation)-1)-position)
//line /usr/local/go/src/crypto/des/block.go:90
		// _ = "end of CoverTab[1910]"
	}
//line /usr/local/go/src/crypto/des/block.go:91
	// _ = "end of CoverTab[1908]"
//line /usr/local/go/src/crypto/des/block.go:91
	_go_fuzz_dep_.CoverTab[1909]++
							return
//line /usr/local/go/src/crypto/des/block.go:92
	// _ = "end of CoverTab[1909]"
}

func initFeistelBox() {
//line /usr/local/go/src/crypto/des/block.go:95
	_go_fuzz_dep_.CoverTab[1911]++
							for s := range sBoxes {
//line /usr/local/go/src/crypto/des/block.go:96
		_go_fuzz_dep_.CoverTab[1912]++
								for i := 0; i < 4; i++ {
//line /usr/local/go/src/crypto/des/block.go:97
			_go_fuzz_dep_.CoverTab[1913]++
									for j := 0; j < 16; j++ {
//line /usr/local/go/src/crypto/des/block.go:98
				_go_fuzz_dep_.CoverTab[1914]++
										f := uint64(sBoxes[s][i][j]) << (4 * (7 - uint(s)))
										f = permuteBlock(f, permutationFunction[:])

//line /usr/local/go/src/crypto/des/block.go:104
				row := uint8(((i & 2) << 4) | i&1)
										col := uint8(j << 1)
										t := row | col

//line /usr/local/go/src/crypto/des/block.go:109
				f = (f << 1) | (f >> 31)

										feistelBox[s][t] = uint32(f)
//line /usr/local/go/src/crypto/des/block.go:111
				// _ = "end of CoverTab[1914]"
			}
//line /usr/local/go/src/crypto/des/block.go:112
			// _ = "end of CoverTab[1913]"
		}
//line /usr/local/go/src/crypto/des/block.go:113
		// _ = "end of CoverTab[1912]"
	}
//line /usr/local/go/src/crypto/des/block.go:114
	// _ = "end of CoverTab[1911]"
}

// permuteInitialBlock is equivalent to the permutation defined
//line /usr/local/go/src/crypto/des/block.go:117
// by initialPermutation.
//line /usr/local/go/src/crypto/des/block.go:119
func permuteInitialBlock(block uint64) uint64 {
//line /usr/local/go/src/crypto/des/block.go:119
	_go_fuzz_dep_.CoverTab[1915]++

							b1 := block >> 48
							b2 := block << 48
							block ^= b1 ^ b2 ^ b1<<48 ^ b2>>48

//line /usr/local/go/src/crypto/des/block.go:126
	b1 = block >> 32 & 0xff00ff
							b2 = (block & 0xff00ff00)
							block ^= b1<<32 ^ b2 ^ b1<<8 ^ b2<<24

//line /usr/local/go/src/crypto/des/block.go:141
	b1 = block & 0x0f0f00000f0f0000
							b2 = block & 0x0000f0f00000f0f0
							block ^= b1 ^ b2 ^ b1>>12 ^ b2<<12

//line /usr/local/go/src/crypto/des/block.go:155
	b1 = block & 0x3300330033003300
							b2 = block & 0x00cc00cc00cc00cc
							block ^= b1 ^ b2 ^ b1>>6 ^ b2<<6

//line /usr/local/go/src/crypto/des/block.go:170
	b1 = block & 0xaaaaaaaa55555555
							block ^= b1 ^ b1>>33 ^ b1<<33

//line /usr/local/go/src/crypto/des/block.go:182
	return block
//line /usr/local/go/src/crypto/des/block.go:182
	// _ = "end of CoverTab[1915]"
}

// permuteFinalBlock is equivalent to the permutation defined
//line /usr/local/go/src/crypto/des/block.go:185
// by finalPermutation.
//line /usr/local/go/src/crypto/des/block.go:187
func permuteFinalBlock(block uint64) uint64 {
//line /usr/local/go/src/crypto/des/block.go:187
	_go_fuzz_dep_.CoverTab[1916]++

//line /usr/local/go/src/crypto/des/block.go:190
	b1 := block & 0xaaaaaaaa55555555
							block ^= b1 ^ b1>>33 ^ b1<<33

							b1 = block & 0x3300330033003300
							b2 := block & 0x00cc00cc00cc00cc
							block ^= b1 ^ b2 ^ b1>>6 ^ b2<<6

							b1 = block & 0x0f0f00000f0f0000
							b2 = block & 0x0000f0f00000f0f0
							block ^= b1 ^ b2 ^ b1>>12 ^ b2<<12

							b1 = block >> 32 & 0xff00ff
							b2 = (block & 0xff00ff00)
							block ^= b1<<32 ^ b2 ^ b1<<8 ^ b2<<24

							b1 = block >> 48
							b2 = block << 48
							block ^= b1 ^ b2 ^ b1<<48 ^ b2>>48
							return block
//line /usr/local/go/src/crypto/des/block.go:208
	// _ = "end of CoverTab[1916]"
}

// creates 16 28-bit blocks rotated according
//line /usr/local/go/src/crypto/des/block.go:211
// to the rotation schedule.
//line /usr/local/go/src/crypto/des/block.go:213
func ksRotate(in uint32) (out []uint32) {
//line /usr/local/go/src/crypto/des/block.go:213
	_go_fuzz_dep_.CoverTab[1917]++
							out = make([]uint32, 16)
							last := in
							for i := 0; i < 16; i++ {
//line /usr/local/go/src/crypto/des/block.go:216
		_go_fuzz_dep_.CoverTab[1919]++

								left := (last << (4 + ksRotations[i])) >> 4
								right := (last << 4) >> (32 - ksRotations[i])
								out[i] = left | right
								last = out[i]
//line /usr/local/go/src/crypto/des/block.go:221
		// _ = "end of CoverTab[1919]"
	}
//line /usr/local/go/src/crypto/des/block.go:222
	// _ = "end of CoverTab[1917]"
//line /usr/local/go/src/crypto/des/block.go:222
	_go_fuzz_dep_.CoverTab[1918]++
							return
//line /usr/local/go/src/crypto/des/block.go:223
	// _ = "end of CoverTab[1918]"
}

// creates 16 56-bit subkeys from the original key.
func (c *desCipher) generateSubkeys(keyBytes []byte) {
//line /usr/local/go/src/crypto/des/block.go:227
	_go_fuzz_dep_.CoverTab[1920]++
							feistelBoxOnce.Do(initFeistelBox)

//line /usr/local/go/src/crypto/des/block.go:231
	key := binary.BigEndian.Uint64(keyBytes)
							permutedKey := permuteBlock(key, permutedChoice1[:])

//line /usr/local/go/src/crypto/des/block.go:235
	leftRotations := ksRotate(uint32(permutedKey >> 28))
							rightRotations := ksRotate(uint32(permutedKey<<4) >> 4)

//line /usr/local/go/src/crypto/des/block.go:239
	for i := 0; i < 16; i++ {
//line /usr/local/go/src/crypto/des/block.go:239
		_go_fuzz_dep_.CoverTab[1921]++

								pc2Input := uint64(leftRotations[i])<<28 | uint64(rightRotations[i])

								c.subkeys[i] = unpack(permuteBlock(pc2Input, permutedChoice2[:]))
//line /usr/local/go/src/crypto/des/block.go:243
		// _ = "end of CoverTab[1921]"
	}
//line /usr/local/go/src/crypto/des/block.go:244
	// _ = "end of CoverTab[1920]"
}

// Expand 48-bit input to 64-bit, with each 6-bit block padded by extra two bits at the top.
//line /usr/local/go/src/crypto/des/block.go:247
// By doing so, we can have the input blocks (four bits each), and the key blocks (six bits each) well-aligned without
//line /usr/local/go/src/crypto/des/block.go:247
// extra shifts/rotations for alignments.
//line /usr/local/go/src/crypto/des/block.go:250
func unpack(x uint64) uint64 {
//line /usr/local/go/src/crypto/des/block.go:250
	_go_fuzz_dep_.CoverTab[1922]++
							return ((x>>(6*1))&0xff)<<(8*0) |
		((x>>(6*3))&0xff)<<(8*1) |
		((x>>(6*5))&0xff)<<(8*2) |
		((x>>(6*7))&0xff)<<(8*3) |
		((x>>(6*0))&0xff)<<(8*4) |
		((x>>(6*2))&0xff)<<(8*5) |
		((x>>(6*4))&0xff)<<(8*6) |
		((x>>(6*6))&0xff)<<(8*7)
//line /usr/local/go/src/crypto/des/block.go:258
	// _ = "end of CoverTab[1922]"
}

//line /usr/local/go/src/crypto/des/block.go:259
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/des/block.go:259
var _ = _go_fuzz_dep_.CoverTab
