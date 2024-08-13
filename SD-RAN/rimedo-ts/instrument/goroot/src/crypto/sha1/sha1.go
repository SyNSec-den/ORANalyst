// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/sha1/sha1.go:5
// Package sha1 implements the SHA-1 hash algorithm as defined in RFC 3174.
//line /usr/local/go/src/crypto/sha1/sha1.go:5
//
//line /usr/local/go/src/crypto/sha1/sha1.go:5
// SHA-1 is cryptographically broken and should not be used for secure
//line /usr/local/go/src/crypto/sha1/sha1.go:5
// applications.
//line /usr/local/go/src/crypto/sha1/sha1.go:9
package sha1

//line /usr/local/go/src/crypto/sha1/sha1.go:9
import (
//line /usr/local/go/src/crypto/sha1/sha1.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha1/sha1.go:9
)
//line /usr/local/go/src/crypto/sha1/sha1.go:9
import (
//line /usr/local/go/src/crypto/sha1/sha1.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha1/sha1.go:9
)

import (
	"crypto"
	"encoding/binary"
	"errors"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA1, New)
}

// The size of a SHA-1 checksum in bytes.
const Size = 20

// The blocksize of SHA-1 in bytes.
const BlockSize = 64

const (
	chunk	= 64
	init0	= 0x67452301
	init1	= 0xEFCDAB89
	init2	= 0x98BADCFE
	init3	= 0x10325476
	init4	= 0xC3D2E1F0
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	h	[5]uint32
	x	[chunk]byte
	nx	int
	len	uint64
}

const (
	magic		= "sha\x01"
	marshaledSize	= len(magic) + 5*4 + chunk + 8
)

func (d *digest) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/crypto/sha1/sha1.go:50
	_go_fuzz_dep_.CoverTab[10084]++
							b := make([]byte, 0, marshaledSize)
							b = append(b, magic...)
							b = binary.BigEndian.AppendUint32(b, d.h[0])
							b = binary.BigEndian.AppendUint32(b, d.h[1])
							b = binary.BigEndian.AppendUint32(b, d.h[2])
							b = binary.BigEndian.AppendUint32(b, d.h[3])
							b = binary.BigEndian.AppendUint32(b, d.h[4])
							b = append(b, d.x[:d.nx]...)
							b = b[:len(b)+len(d.x)-d.nx]
							b = binary.BigEndian.AppendUint64(b, d.len)
							return b, nil
//line /usr/local/go/src/crypto/sha1/sha1.go:61
	// _ = "end of CoverTab[10084]"
}

func (d *digest) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/crypto/sha1/sha1.go:64
	_go_fuzz_dep_.CoverTab[10085]++
							if len(b) < len(magic) || func() bool {
//line /usr/local/go/src/crypto/sha1/sha1.go:65
		_go_fuzz_dep_.CoverTab[10088]++
//line /usr/local/go/src/crypto/sha1/sha1.go:65
		return string(b[:len(magic)]) != magic
//line /usr/local/go/src/crypto/sha1/sha1.go:65
		// _ = "end of CoverTab[10088]"
//line /usr/local/go/src/crypto/sha1/sha1.go:65
	}() {
//line /usr/local/go/src/crypto/sha1/sha1.go:65
		_go_fuzz_dep_.CoverTab[10089]++
								return errors.New("crypto/sha1: invalid hash state identifier")
//line /usr/local/go/src/crypto/sha1/sha1.go:66
		// _ = "end of CoverTab[10089]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:67
		_go_fuzz_dep_.CoverTab[10090]++
//line /usr/local/go/src/crypto/sha1/sha1.go:67
		// _ = "end of CoverTab[10090]"
//line /usr/local/go/src/crypto/sha1/sha1.go:67
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:67
	// _ = "end of CoverTab[10085]"
//line /usr/local/go/src/crypto/sha1/sha1.go:67
	_go_fuzz_dep_.CoverTab[10086]++
							if len(b) != marshaledSize {
//line /usr/local/go/src/crypto/sha1/sha1.go:68
		_go_fuzz_dep_.CoverTab[10091]++
								return errors.New("crypto/sha1: invalid hash state size")
//line /usr/local/go/src/crypto/sha1/sha1.go:69
		// _ = "end of CoverTab[10091]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:70
		_go_fuzz_dep_.CoverTab[10092]++
//line /usr/local/go/src/crypto/sha1/sha1.go:70
		// _ = "end of CoverTab[10092]"
//line /usr/local/go/src/crypto/sha1/sha1.go:70
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:70
	// _ = "end of CoverTab[10086]"
//line /usr/local/go/src/crypto/sha1/sha1.go:70
	_go_fuzz_dep_.CoverTab[10087]++
							b = b[len(magic):]
							b, d.h[0] = consumeUint32(b)
							b, d.h[1] = consumeUint32(b)
							b, d.h[2] = consumeUint32(b)
							b, d.h[3] = consumeUint32(b)
							b, d.h[4] = consumeUint32(b)
							b = b[copy(d.x[:], b):]
							b, d.len = consumeUint64(b)
							d.nx = int(d.len % chunk)
							return nil
//line /usr/local/go/src/crypto/sha1/sha1.go:80
	// _ = "end of CoverTab[10087]"
}

func consumeUint64(b []byte) ([]byte, uint64) {
//line /usr/local/go/src/crypto/sha1/sha1.go:83
	_go_fuzz_dep_.CoverTab[10093]++
							_ = b[7]
							x := uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
							return b[8:], x
//line /usr/local/go/src/crypto/sha1/sha1.go:87
	// _ = "end of CoverTab[10093]"
}

func consumeUint32(b []byte) ([]byte, uint32) {
//line /usr/local/go/src/crypto/sha1/sha1.go:90
	_go_fuzz_dep_.CoverTab[10094]++
							_ = b[3]
							x := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
							return b[4:], x
//line /usr/local/go/src/crypto/sha1/sha1.go:93
	// _ = "end of CoverTab[10094]"
}

func (d *digest) Reset() {
//line /usr/local/go/src/crypto/sha1/sha1.go:96
	_go_fuzz_dep_.CoverTab[10095]++
							d.h[0] = init0
							d.h[1] = init1
							d.h[2] = init2
							d.h[3] = init3
							d.h[4] = init4
							d.nx = 0
							d.len = 0
//line /usr/local/go/src/crypto/sha1/sha1.go:103
	// _ = "end of CoverTab[10095]"
}

// New returns a new hash.Hash computing the SHA1 checksum. The Hash also
//line /usr/local/go/src/crypto/sha1/sha1.go:106
// implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
//line /usr/local/go/src/crypto/sha1/sha1.go:106
// marshal and unmarshal the internal state of the hash.
//line /usr/local/go/src/crypto/sha1/sha1.go:109
func New() hash.Hash {
//line /usr/local/go/src/crypto/sha1/sha1.go:109
	_go_fuzz_dep_.CoverTab[10096]++
							if boringEnabled {
//line /usr/local/go/src/crypto/sha1/sha1.go:110
		_go_fuzz_dep_.CoverTab[10098]++
								return boringNewSHA1()
//line /usr/local/go/src/crypto/sha1/sha1.go:111
		// _ = "end of CoverTab[10098]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:112
		_go_fuzz_dep_.CoverTab[10099]++
//line /usr/local/go/src/crypto/sha1/sha1.go:112
		// _ = "end of CoverTab[10099]"
//line /usr/local/go/src/crypto/sha1/sha1.go:112
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:112
	// _ = "end of CoverTab[10096]"
//line /usr/local/go/src/crypto/sha1/sha1.go:112
	_go_fuzz_dep_.CoverTab[10097]++
							d := new(digest)
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha1/sha1.go:115
	// _ = "end of CoverTab[10097]"
}

func (d *digest) Size() int {
//line /usr/local/go/src/crypto/sha1/sha1.go:118
	_go_fuzz_dep_.CoverTab[10100]++
//line /usr/local/go/src/crypto/sha1/sha1.go:118
	return Size
//line /usr/local/go/src/crypto/sha1/sha1.go:118
	// _ = "end of CoverTab[10100]"
//line /usr/local/go/src/crypto/sha1/sha1.go:118
}

func (d *digest) BlockSize() int {
//line /usr/local/go/src/crypto/sha1/sha1.go:120
	_go_fuzz_dep_.CoverTab[10101]++
//line /usr/local/go/src/crypto/sha1/sha1.go:120
	return BlockSize
//line /usr/local/go/src/crypto/sha1/sha1.go:120
	// _ = "end of CoverTab[10101]"
//line /usr/local/go/src/crypto/sha1/sha1.go:120
}

func (d *digest) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/crypto/sha1/sha1.go:122
	_go_fuzz_dep_.CoverTab[10102]++
							boringUnreachable()
							nn = len(p)
							d.len += uint64(nn)
							if d.nx > 0 {
//line /usr/local/go/src/crypto/sha1/sha1.go:126
		_go_fuzz_dep_.CoverTab[10106]++
								n := copy(d.x[d.nx:], p)
								d.nx += n
								if d.nx == chunk {
//line /usr/local/go/src/crypto/sha1/sha1.go:129
			_go_fuzz_dep_.CoverTab[10108]++
									block(d, d.x[:])
									d.nx = 0
//line /usr/local/go/src/crypto/sha1/sha1.go:131
			// _ = "end of CoverTab[10108]"
		} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:132
			_go_fuzz_dep_.CoverTab[10109]++
//line /usr/local/go/src/crypto/sha1/sha1.go:132
			// _ = "end of CoverTab[10109]"
//line /usr/local/go/src/crypto/sha1/sha1.go:132
		}
//line /usr/local/go/src/crypto/sha1/sha1.go:132
		// _ = "end of CoverTab[10106]"
//line /usr/local/go/src/crypto/sha1/sha1.go:132
		_go_fuzz_dep_.CoverTab[10107]++
								p = p[n:]
//line /usr/local/go/src/crypto/sha1/sha1.go:133
		// _ = "end of CoverTab[10107]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:134
		_go_fuzz_dep_.CoverTab[10110]++
//line /usr/local/go/src/crypto/sha1/sha1.go:134
		// _ = "end of CoverTab[10110]"
//line /usr/local/go/src/crypto/sha1/sha1.go:134
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:134
	// _ = "end of CoverTab[10102]"
//line /usr/local/go/src/crypto/sha1/sha1.go:134
	_go_fuzz_dep_.CoverTab[10103]++
							if len(p) >= chunk {
//line /usr/local/go/src/crypto/sha1/sha1.go:135
		_go_fuzz_dep_.CoverTab[10111]++
								n := len(p) &^ (chunk - 1)
								block(d, p[:n])
								p = p[n:]
//line /usr/local/go/src/crypto/sha1/sha1.go:138
		// _ = "end of CoverTab[10111]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:139
		_go_fuzz_dep_.CoverTab[10112]++
//line /usr/local/go/src/crypto/sha1/sha1.go:139
		// _ = "end of CoverTab[10112]"
//line /usr/local/go/src/crypto/sha1/sha1.go:139
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:139
	// _ = "end of CoverTab[10103]"
//line /usr/local/go/src/crypto/sha1/sha1.go:139
	_go_fuzz_dep_.CoverTab[10104]++
							if len(p) > 0 {
//line /usr/local/go/src/crypto/sha1/sha1.go:140
		_go_fuzz_dep_.CoverTab[10113]++
								d.nx = copy(d.x[:], p)
//line /usr/local/go/src/crypto/sha1/sha1.go:141
		// _ = "end of CoverTab[10113]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:142
		_go_fuzz_dep_.CoverTab[10114]++
//line /usr/local/go/src/crypto/sha1/sha1.go:142
		// _ = "end of CoverTab[10114]"
//line /usr/local/go/src/crypto/sha1/sha1.go:142
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:142
	// _ = "end of CoverTab[10104]"
//line /usr/local/go/src/crypto/sha1/sha1.go:142
	_go_fuzz_dep_.CoverTab[10105]++
							return
//line /usr/local/go/src/crypto/sha1/sha1.go:143
	// _ = "end of CoverTab[10105]"
}

func (d *digest) Sum(in []byte) []byte {
//line /usr/local/go/src/crypto/sha1/sha1.go:146
	_go_fuzz_dep_.CoverTab[10115]++
							boringUnreachable()

							d0 := *d
							hash := d0.checkSum()
							return append(in, hash[:]...)
//line /usr/local/go/src/crypto/sha1/sha1.go:151
	// _ = "end of CoverTab[10115]"
}

func (d *digest) checkSum() [Size]byte {
//line /usr/local/go/src/crypto/sha1/sha1.go:154
	_go_fuzz_dep_.CoverTab[10116]++
							len := d.len
	// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.
	var tmp [64 + 8]byte	// padding + length buffer
	tmp[0] = 0x80
	var t uint64
	if len%64 < 56 {
//line /usr/local/go/src/crypto/sha1/sha1.go:160
		_go_fuzz_dep_.CoverTab[10119]++
								t = 56 - len%64
//line /usr/local/go/src/crypto/sha1/sha1.go:161
		// _ = "end of CoverTab[10119]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:162
		_go_fuzz_dep_.CoverTab[10120]++
								t = 64 + 56 - len%64
//line /usr/local/go/src/crypto/sha1/sha1.go:163
		// _ = "end of CoverTab[10120]"
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:164
	// _ = "end of CoverTab[10116]"
//line /usr/local/go/src/crypto/sha1/sha1.go:164
	_go_fuzz_dep_.CoverTab[10117]++

//line /usr/local/go/src/crypto/sha1/sha1.go:167
	len <<= 3
	padlen := tmp[:t+8]
	binary.BigEndian.PutUint64(padlen[t:], len)
	d.Write(padlen)

	if d.nx != 0 {
//line /usr/local/go/src/crypto/sha1/sha1.go:172
		_go_fuzz_dep_.CoverTab[10121]++
								panic("d.nx != 0")
//line /usr/local/go/src/crypto/sha1/sha1.go:173
		// _ = "end of CoverTab[10121]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:174
		_go_fuzz_dep_.CoverTab[10122]++
//line /usr/local/go/src/crypto/sha1/sha1.go:174
		// _ = "end of CoverTab[10122]"
//line /usr/local/go/src/crypto/sha1/sha1.go:174
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:174
	// _ = "end of CoverTab[10117]"
//line /usr/local/go/src/crypto/sha1/sha1.go:174
	_go_fuzz_dep_.CoverTab[10118]++

							var digest [Size]byte

							binary.BigEndian.PutUint32(digest[0:], d.h[0])
							binary.BigEndian.PutUint32(digest[4:], d.h[1])
							binary.BigEndian.PutUint32(digest[8:], d.h[2])
							binary.BigEndian.PutUint32(digest[12:], d.h[3])
							binary.BigEndian.PutUint32(digest[16:], d.h[4])

							return digest
//line /usr/local/go/src/crypto/sha1/sha1.go:184
	// _ = "end of CoverTab[10118]"
}

// ConstantTimeSum computes the same result of Sum() but in constant time
func (d *digest) ConstantTimeSum(in []byte) []byte {
//line /usr/local/go/src/crypto/sha1/sha1.go:188
	_go_fuzz_dep_.CoverTab[10123]++
							d0 := *d
							hash := d0.constSum()
							return append(in, hash[:]...)
//line /usr/local/go/src/crypto/sha1/sha1.go:191
	// _ = "end of CoverTab[10123]"
}

func (d *digest) constSum() [Size]byte {
//line /usr/local/go/src/crypto/sha1/sha1.go:194
	_go_fuzz_dep_.CoverTab[10124]++
							var length [8]byte
							l := d.len << 3
							for i := uint(0); i < 8; i++ {
//line /usr/local/go/src/crypto/sha1/sha1.go:197
		_go_fuzz_dep_.CoverTab[10130]++
								length[i] = byte(l >> (56 - 8*i))
//line /usr/local/go/src/crypto/sha1/sha1.go:198
		// _ = "end of CoverTab[10130]"
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:199
	// _ = "end of CoverTab[10124]"
//line /usr/local/go/src/crypto/sha1/sha1.go:199
	_go_fuzz_dep_.CoverTab[10125]++

							nx := byte(d.nx)
							t := nx - 56
							mask1b := byte(int8(t) >> 7)

							separator := byte(0x80)
							for i := byte(0); i < chunk; i++ {
//line /usr/local/go/src/crypto/sha1/sha1.go:206
		_go_fuzz_dep_.CoverTab[10131]++
								mask := byte(int8(i-nx) >> 7)

//line /usr/local/go/src/crypto/sha1/sha1.go:210
		d.x[i] = (^mask & separator) | (mask & d.x[i])

//line /usr/local/go/src/crypto/sha1/sha1.go:213
		separator &= mask

		if i >= 56 {
//line /usr/local/go/src/crypto/sha1/sha1.go:215
			_go_fuzz_dep_.CoverTab[10132]++

									d.x[i] |= mask1b & length[i-56]
//line /usr/local/go/src/crypto/sha1/sha1.go:217
			// _ = "end of CoverTab[10132]"
		} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:218
			_go_fuzz_dep_.CoverTab[10133]++
//line /usr/local/go/src/crypto/sha1/sha1.go:218
			// _ = "end of CoverTab[10133]"
//line /usr/local/go/src/crypto/sha1/sha1.go:218
		}
//line /usr/local/go/src/crypto/sha1/sha1.go:218
		// _ = "end of CoverTab[10131]"
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:219
	// _ = "end of CoverTab[10125]"
//line /usr/local/go/src/crypto/sha1/sha1.go:219
	_go_fuzz_dep_.CoverTab[10126]++

//line /usr/local/go/src/crypto/sha1/sha1.go:222
	block(d, d.x[:])

	var digest [Size]byte
	for i, s := range d.h {
//line /usr/local/go/src/crypto/sha1/sha1.go:225
		_go_fuzz_dep_.CoverTab[10134]++
								digest[i*4] = mask1b & byte(s>>24)
								digest[i*4+1] = mask1b & byte(s>>16)
								digest[i*4+2] = mask1b & byte(s>>8)
								digest[i*4+3] = mask1b & byte(s)
//line /usr/local/go/src/crypto/sha1/sha1.go:229
		// _ = "end of CoverTab[10134]"
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:230
	// _ = "end of CoverTab[10126]"
//line /usr/local/go/src/crypto/sha1/sha1.go:230
	_go_fuzz_dep_.CoverTab[10127]++

							for i := byte(0); i < chunk; i++ {
//line /usr/local/go/src/crypto/sha1/sha1.go:232
		_go_fuzz_dep_.CoverTab[10135]++

								if i < 56 {
//line /usr/local/go/src/crypto/sha1/sha1.go:234
			_go_fuzz_dep_.CoverTab[10136]++
									d.x[i] = separator
									separator = 0
//line /usr/local/go/src/crypto/sha1/sha1.go:236
			// _ = "end of CoverTab[10136]"
		} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:237
			_go_fuzz_dep_.CoverTab[10137]++
									d.x[i] = length[i-56]
//line /usr/local/go/src/crypto/sha1/sha1.go:238
			// _ = "end of CoverTab[10137]"
		}
//line /usr/local/go/src/crypto/sha1/sha1.go:239
		// _ = "end of CoverTab[10135]"
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:240
	// _ = "end of CoverTab[10127]"
//line /usr/local/go/src/crypto/sha1/sha1.go:240
	_go_fuzz_dep_.CoverTab[10128]++

//line /usr/local/go/src/crypto/sha1/sha1.go:243
	block(d, d.x[:])

	for i, s := range d.h {
//line /usr/local/go/src/crypto/sha1/sha1.go:245
		_go_fuzz_dep_.CoverTab[10138]++
								digest[i*4] |= ^mask1b & byte(s>>24)
								digest[i*4+1] |= ^mask1b & byte(s>>16)
								digest[i*4+2] |= ^mask1b & byte(s>>8)
								digest[i*4+3] |= ^mask1b & byte(s)
//line /usr/local/go/src/crypto/sha1/sha1.go:249
		// _ = "end of CoverTab[10138]"
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:250
	// _ = "end of CoverTab[10128]"
//line /usr/local/go/src/crypto/sha1/sha1.go:250
	_go_fuzz_dep_.CoverTab[10129]++

							return digest
//line /usr/local/go/src/crypto/sha1/sha1.go:252
	// _ = "end of CoverTab[10129]"
}

// Sum returns the SHA-1 checksum of the data.
func Sum(data []byte) [Size]byte {
//line /usr/local/go/src/crypto/sha1/sha1.go:256
	_go_fuzz_dep_.CoverTab[10139]++
							if boringEnabled {
//line /usr/local/go/src/crypto/sha1/sha1.go:257
		_go_fuzz_dep_.CoverTab[10141]++
								return boringSHA1(data)
//line /usr/local/go/src/crypto/sha1/sha1.go:258
		// _ = "end of CoverTab[10141]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1.go:259
		_go_fuzz_dep_.CoverTab[10142]++
//line /usr/local/go/src/crypto/sha1/sha1.go:259
		// _ = "end of CoverTab[10142]"
//line /usr/local/go/src/crypto/sha1/sha1.go:259
	}
//line /usr/local/go/src/crypto/sha1/sha1.go:259
	// _ = "end of CoverTab[10139]"
//line /usr/local/go/src/crypto/sha1/sha1.go:259
	_go_fuzz_dep_.CoverTab[10140]++
							var d digest
							d.Reset()
							d.Write(data)
							return d.checkSum()
//line /usr/local/go/src/crypto/sha1/sha1.go:263
	// _ = "end of CoverTab[10140]"
}

//line /usr/local/go/src/crypto/sha1/sha1.go:264
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha1/sha1.go:264
var _ = _go_fuzz_dep_.CoverTab
