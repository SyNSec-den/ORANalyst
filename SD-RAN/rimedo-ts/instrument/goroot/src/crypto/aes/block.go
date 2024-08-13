// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This Go implementation is derived in part from the reference
// ANSI C implementation, which carries the following notice:
//
//	rijndael-alg-fst.c
//
//	@version 3.0 (December 2000)
//
//	Optimised ANSI C code for the Rijndael cipher (now AES)
//
//	@author Vincent Rijmen <vincent.rijmen@esat.kuleuven.ac.be>
//	@author Antoon Bosselaers <antoon.bosselaers@esat.kuleuven.ac.be>
//	@author Paulo Barreto <paulo.barreto@terra.com.br>
//
//	This code is hereby placed in the public domain.
//
//	THIS SOFTWARE IS PROVIDED BY THE AUTHORS ''AS IS'' AND ANY EXPRESS
//	OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//	WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
//	ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHORS OR CONTRIBUTORS BE
//	LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
//	CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
//	SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
//	BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
//	WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
//	OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,
//	EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// See FIPS 197 for specification, and see Daemen and Rijmen's Rijndael submission
// for implementation details.
//	https://csrc.nist.gov/csrc/media/publications/fips/197/final/documents/fips-197.pdf
//	https://csrc.nist.gov/archive/aes/rijndael/Rijndael-ammended.pdf

//line /usr/local/go/src/crypto/aes/block.go:37
package aes

//line /usr/local/go/src/crypto/aes/block.go:37
import (
//line /usr/local/go/src/crypto/aes/block.go:37
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/aes/block.go:37
)
//line /usr/local/go/src/crypto/aes/block.go:37
import (
//line /usr/local/go/src/crypto/aes/block.go:37
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/aes/block.go:37
)

import (
	"encoding/binary"
)

// Encrypt one block from src into dst, using the expanded key xk.
func encryptBlockGo(xk []uint32, dst, src []byte) {
//line /usr/local/go/src/crypto/aes/block.go:44
	_go_fuzz_dep_.CoverTab[1796]++
							_ = src[15]
							s0 := binary.BigEndian.Uint32(src[0:4])
							s1 := binary.BigEndian.Uint32(src[4:8])
							s2 := binary.BigEndian.Uint32(src[8:12])
							s3 := binary.BigEndian.Uint32(src[12:16])

//line /usr/local/go/src/crypto/aes/block.go:52
	s0 ^= xk[0]
							s1 ^= xk[1]
							s2 ^= xk[2]
							s3 ^= xk[3]

//line /usr/local/go/src/crypto/aes/block.go:59
	nr := len(xk)/4 - 2
	k := 4
	var t0, t1, t2, t3 uint32
	for r := 0; r < nr; r++ {
//line /usr/local/go/src/crypto/aes/block.go:62
		_go_fuzz_dep_.CoverTab[1798]++
								t0 = xk[k+0] ^ te0[uint8(s0>>24)] ^ te1[uint8(s1>>16)] ^ te2[uint8(s2>>8)] ^ te3[uint8(s3)]
								t1 = xk[k+1] ^ te0[uint8(s1>>24)] ^ te1[uint8(s2>>16)] ^ te2[uint8(s3>>8)] ^ te3[uint8(s0)]
								t2 = xk[k+2] ^ te0[uint8(s2>>24)] ^ te1[uint8(s3>>16)] ^ te2[uint8(s0>>8)] ^ te3[uint8(s1)]
								t3 = xk[k+3] ^ te0[uint8(s3>>24)] ^ te1[uint8(s0>>16)] ^ te2[uint8(s1>>8)] ^ te3[uint8(s2)]
								k += 4
								s0, s1, s2, s3 = t0, t1, t2, t3
//line /usr/local/go/src/crypto/aes/block.go:68
		// _ = "end of CoverTab[1798]"
	}
//line /usr/local/go/src/crypto/aes/block.go:69
	// _ = "end of CoverTab[1796]"
//line /usr/local/go/src/crypto/aes/block.go:69
	_go_fuzz_dep_.CoverTab[1797]++

//line /usr/local/go/src/crypto/aes/block.go:72
	s0 = uint32(sbox0[t0>>24])<<24 | uint32(sbox0[t1>>16&0xff])<<16 | uint32(sbox0[t2>>8&0xff])<<8 | uint32(sbox0[t3&0xff])
							s1 = uint32(sbox0[t1>>24])<<24 | uint32(sbox0[t2>>16&0xff])<<16 | uint32(sbox0[t3>>8&0xff])<<8 | uint32(sbox0[t0&0xff])
							s2 = uint32(sbox0[t2>>24])<<24 | uint32(sbox0[t3>>16&0xff])<<16 | uint32(sbox0[t0>>8&0xff])<<8 | uint32(sbox0[t1&0xff])
							s3 = uint32(sbox0[t3>>24])<<24 | uint32(sbox0[t0>>16&0xff])<<16 | uint32(sbox0[t1>>8&0xff])<<8 | uint32(sbox0[t2&0xff])

							s0 ^= xk[k+0]
							s1 ^= xk[k+1]
							s2 ^= xk[k+2]
							s3 ^= xk[k+3]

							_ = dst[15]
							binary.BigEndian.PutUint32(dst[0:4], s0)
							binary.BigEndian.PutUint32(dst[4:8], s1)
							binary.BigEndian.PutUint32(dst[8:12], s2)
							binary.BigEndian.PutUint32(dst[12:16], s3)
//line /usr/local/go/src/crypto/aes/block.go:86
	// _ = "end of CoverTab[1797]"
}

// Decrypt one block from src into dst, using the expanded key xk.
func decryptBlockGo(xk []uint32, dst, src []byte) {
//line /usr/local/go/src/crypto/aes/block.go:90
	_go_fuzz_dep_.CoverTab[1799]++
							_ = src[15]
							s0 := binary.BigEndian.Uint32(src[0:4])
							s1 := binary.BigEndian.Uint32(src[4:8])
							s2 := binary.BigEndian.Uint32(src[8:12])
							s3 := binary.BigEndian.Uint32(src[12:16])

//line /usr/local/go/src/crypto/aes/block.go:98
	s0 ^= xk[0]
							s1 ^= xk[1]
							s2 ^= xk[2]
							s3 ^= xk[3]

//line /usr/local/go/src/crypto/aes/block.go:105
	nr := len(xk)/4 - 2
	k := 4
	var t0, t1, t2, t3 uint32
	for r := 0; r < nr; r++ {
//line /usr/local/go/src/crypto/aes/block.go:108
		_go_fuzz_dep_.CoverTab[1801]++
								t0 = xk[k+0] ^ td0[uint8(s0>>24)] ^ td1[uint8(s3>>16)] ^ td2[uint8(s2>>8)] ^ td3[uint8(s1)]
								t1 = xk[k+1] ^ td0[uint8(s1>>24)] ^ td1[uint8(s0>>16)] ^ td2[uint8(s3>>8)] ^ td3[uint8(s2)]
								t2 = xk[k+2] ^ td0[uint8(s2>>24)] ^ td1[uint8(s1>>16)] ^ td2[uint8(s0>>8)] ^ td3[uint8(s3)]
								t3 = xk[k+3] ^ td0[uint8(s3>>24)] ^ td1[uint8(s2>>16)] ^ td2[uint8(s1>>8)] ^ td3[uint8(s0)]
								k += 4
								s0, s1, s2, s3 = t0, t1, t2, t3
//line /usr/local/go/src/crypto/aes/block.go:114
		// _ = "end of CoverTab[1801]"
	}
//line /usr/local/go/src/crypto/aes/block.go:115
	// _ = "end of CoverTab[1799]"
//line /usr/local/go/src/crypto/aes/block.go:115
	_go_fuzz_dep_.CoverTab[1800]++

//line /usr/local/go/src/crypto/aes/block.go:118
	s0 = uint32(sbox1[t0>>24])<<24 | uint32(sbox1[t3>>16&0xff])<<16 | uint32(sbox1[t2>>8&0xff])<<8 | uint32(sbox1[t1&0xff])
							s1 = uint32(sbox1[t1>>24])<<24 | uint32(sbox1[t0>>16&0xff])<<16 | uint32(sbox1[t3>>8&0xff])<<8 | uint32(sbox1[t2&0xff])
							s2 = uint32(sbox1[t2>>24])<<24 | uint32(sbox1[t1>>16&0xff])<<16 | uint32(sbox1[t0>>8&0xff])<<8 | uint32(sbox1[t3&0xff])
							s3 = uint32(sbox1[t3>>24])<<24 | uint32(sbox1[t2>>16&0xff])<<16 | uint32(sbox1[t1>>8&0xff])<<8 | uint32(sbox1[t0&0xff])

							s0 ^= xk[k+0]
							s1 ^= xk[k+1]
							s2 ^= xk[k+2]
							s3 ^= xk[k+3]

							_ = dst[15]
							binary.BigEndian.PutUint32(dst[0:4], s0)
							binary.BigEndian.PutUint32(dst[4:8], s1)
							binary.BigEndian.PutUint32(dst[8:12], s2)
							binary.BigEndian.PutUint32(dst[12:16], s3)
//line /usr/local/go/src/crypto/aes/block.go:132
	// _ = "end of CoverTab[1800]"
}

// Apply sbox0 to each byte in w.
func subw(w uint32) uint32 {
//line /usr/local/go/src/crypto/aes/block.go:136
	_go_fuzz_dep_.CoverTab[1802]++
							return uint32(sbox0[w>>24])<<24 |
		uint32(sbox0[w>>16&0xff])<<16 |
		uint32(sbox0[w>>8&0xff])<<8 |
		uint32(sbox0[w&0xff])
//line /usr/local/go/src/crypto/aes/block.go:140
	// _ = "end of CoverTab[1802]"
}

// Rotate
func rotw(w uint32) uint32 {
//line /usr/local/go/src/crypto/aes/block.go:144
	_go_fuzz_dep_.CoverTab[1803]++
//line /usr/local/go/src/crypto/aes/block.go:144
	return w<<8 | w>>24
//line /usr/local/go/src/crypto/aes/block.go:144
	// _ = "end of CoverTab[1803]"
//line /usr/local/go/src/crypto/aes/block.go:144
}

// Key expansion algorithm. See FIPS-197, Figure 11.
//line /usr/local/go/src/crypto/aes/block.go:146
// Their rcon[i] is our powx[i-1] << 24.
//line /usr/local/go/src/crypto/aes/block.go:148
func expandKeyGo(key []byte, enc, dec []uint32) {
//line /usr/local/go/src/crypto/aes/block.go:148
	_go_fuzz_dep_.CoverTab[1804]++
	// Encryption key setup.
	var i int
	nk := len(key) / 4
	for i = 0; i < nk; i++ {
//line /usr/local/go/src/crypto/aes/block.go:152
		_go_fuzz_dep_.CoverTab[1808]++
								enc[i] = binary.BigEndian.Uint32(key[4*i:])
//line /usr/local/go/src/crypto/aes/block.go:153
		// _ = "end of CoverTab[1808]"
	}
//line /usr/local/go/src/crypto/aes/block.go:154
	// _ = "end of CoverTab[1804]"
//line /usr/local/go/src/crypto/aes/block.go:154
	_go_fuzz_dep_.CoverTab[1805]++
							for ; i < len(enc); i++ {
//line /usr/local/go/src/crypto/aes/block.go:155
		_go_fuzz_dep_.CoverTab[1809]++
								t := enc[i-1]
								if i%nk == 0 {
//line /usr/local/go/src/crypto/aes/block.go:157
			_go_fuzz_dep_.CoverTab[1811]++
									t = subw(rotw(t)) ^ (uint32(powx[i/nk-1]) << 24)
//line /usr/local/go/src/crypto/aes/block.go:158
			// _ = "end of CoverTab[1811]"
		} else {
//line /usr/local/go/src/crypto/aes/block.go:159
			_go_fuzz_dep_.CoverTab[1812]++
//line /usr/local/go/src/crypto/aes/block.go:159
			if nk > 6 && func() bool {
//line /usr/local/go/src/crypto/aes/block.go:159
				_go_fuzz_dep_.CoverTab[1813]++
//line /usr/local/go/src/crypto/aes/block.go:159
				return i%nk == 4
//line /usr/local/go/src/crypto/aes/block.go:159
				// _ = "end of CoverTab[1813]"
//line /usr/local/go/src/crypto/aes/block.go:159
			}() {
//line /usr/local/go/src/crypto/aes/block.go:159
				_go_fuzz_dep_.CoverTab[1814]++
										t = subw(t)
//line /usr/local/go/src/crypto/aes/block.go:160
				// _ = "end of CoverTab[1814]"
			} else {
//line /usr/local/go/src/crypto/aes/block.go:161
				_go_fuzz_dep_.CoverTab[1815]++
//line /usr/local/go/src/crypto/aes/block.go:161
				// _ = "end of CoverTab[1815]"
//line /usr/local/go/src/crypto/aes/block.go:161
			}
//line /usr/local/go/src/crypto/aes/block.go:161
			// _ = "end of CoverTab[1812]"
//line /usr/local/go/src/crypto/aes/block.go:161
		}
//line /usr/local/go/src/crypto/aes/block.go:161
		// _ = "end of CoverTab[1809]"
//line /usr/local/go/src/crypto/aes/block.go:161
		_go_fuzz_dep_.CoverTab[1810]++
								enc[i] = enc[i-nk] ^ t
//line /usr/local/go/src/crypto/aes/block.go:162
		// _ = "end of CoverTab[1810]"
	}
//line /usr/local/go/src/crypto/aes/block.go:163
	// _ = "end of CoverTab[1805]"
//line /usr/local/go/src/crypto/aes/block.go:163
	_go_fuzz_dep_.CoverTab[1806]++

//line /usr/local/go/src/crypto/aes/block.go:168
	if dec == nil {
//line /usr/local/go/src/crypto/aes/block.go:168
		_go_fuzz_dep_.CoverTab[1816]++
								return
//line /usr/local/go/src/crypto/aes/block.go:169
		// _ = "end of CoverTab[1816]"
	} else {
//line /usr/local/go/src/crypto/aes/block.go:170
		_go_fuzz_dep_.CoverTab[1817]++
//line /usr/local/go/src/crypto/aes/block.go:170
		// _ = "end of CoverTab[1817]"
//line /usr/local/go/src/crypto/aes/block.go:170
	}
//line /usr/local/go/src/crypto/aes/block.go:170
	// _ = "end of CoverTab[1806]"
//line /usr/local/go/src/crypto/aes/block.go:170
	_go_fuzz_dep_.CoverTab[1807]++
							n := len(enc)
							for i := 0; i < n; i += 4 {
//line /usr/local/go/src/crypto/aes/block.go:172
		_go_fuzz_dep_.CoverTab[1818]++
								ei := n - i - 4
								for j := 0; j < 4; j++ {
//line /usr/local/go/src/crypto/aes/block.go:174
			_go_fuzz_dep_.CoverTab[1819]++
									x := enc[ei+j]
									if i > 0 && func() bool {
//line /usr/local/go/src/crypto/aes/block.go:176
				_go_fuzz_dep_.CoverTab[1821]++
//line /usr/local/go/src/crypto/aes/block.go:176
				return i+4 < n
//line /usr/local/go/src/crypto/aes/block.go:176
				// _ = "end of CoverTab[1821]"
//line /usr/local/go/src/crypto/aes/block.go:176
			}() {
//line /usr/local/go/src/crypto/aes/block.go:176
				_go_fuzz_dep_.CoverTab[1822]++
										x = td0[sbox0[x>>24]] ^ td1[sbox0[x>>16&0xff]] ^ td2[sbox0[x>>8&0xff]] ^ td3[sbox0[x&0xff]]
//line /usr/local/go/src/crypto/aes/block.go:177
				// _ = "end of CoverTab[1822]"
			} else {
//line /usr/local/go/src/crypto/aes/block.go:178
				_go_fuzz_dep_.CoverTab[1823]++
//line /usr/local/go/src/crypto/aes/block.go:178
				// _ = "end of CoverTab[1823]"
//line /usr/local/go/src/crypto/aes/block.go:178
			}
//line /usr/local/go/src/crypto/aes/block.go:178
			// _ = "end of CoverTab[1819]"
//line /usr/local/go/src/crypto/aes/block.go:178
			_go_fuzz_dep_.CoverTab[1820]++
									dec[i+j] = x
//line /usr/local/go/src/crypto/aes/block.go:179
			// _ = "end of CoverTab[1820]"
		}
//line /usr/local/go/src/crypto/aes/block.go:180
		// _ = "end of CoverTab[1818]"
	}
//line /usr/local/go/src/crypto/aes/block.go:181
	// _ = "end of CoverTab[1807]"
}

//line /usr/local/go/src/crypto/aes/block.go:182
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/aes/block.go:182
var _ = _go_fuzz_dep_.CoverTab
