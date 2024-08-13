// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/sha256/sha256.go:5
// Package sha256 implements the SHA224 and SHA256 hash algorithms as defined
//line /usr/local/go/src/crypto/sha256/sha256.go:5
// in FIPS 180-4.
//line /usr/local/go/src/crypto/sha256/sha256.go:7
package sha256

//line /usr/local/go/src/crypto/sha256/sha256.go:7
import (
//line /usr/local/go/src/crypto/sha256/sha256.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha256/sha256.go:7
)
//line /usr/local/go/src/crypto/sha256/sha256.go:7
import (
//line /usr/local/go/src/crypto/sha256/sha256.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha256/sha256.go:7
)

import (
	"crypto"
	"crypto/internal/boring"
	"encoding/binary"
	"errors"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA224, New224)
	crypto.RegisterHash(crypto.SHA256, New)
}

// The size of a SHA256 checksum in bytes.
const Size = 32

// The size of a SHA224 checksum in bytes.
const Size224 = 28

// The blocksize of SHA256 and SHA224 in bytes.
const BlockSize = 64

const (
	chunk		= 64
	init0		= 0x6A09E667
	init1		= 0xBB67AE85
	init2		= 0x3C6EF372
	init3		= 0xA54FF53A
	init4		= 0x510E527F
	init5		= 0x9B05688C
	init6		= 0x1F83D9AB
	init7		= 0x5BE0CD19
	init0_224	= 0xC1059ED8
	init1_224	= 0x367CD507
	init2_224	= 0x3070DD17
	init3_224	= 0xF70E5939
	init4_224	= 0xFFC00B31
	init5_224	= 0x68581511
	init6_224	= 0x64F98FA7
	init7_224	= 0xBEFA4FA4
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	h	[8]uint32
	x	[chunk]byte
	nx	int
	len	uint64
	is224	bool	// mark if this digest is SHA-224
}

const (
	magic224	= "sha\x02"
	magic256	= "sha\x03"
	marshaledSize	= len(magic256) + 8*4 + chunk + 8
)

func (d *digest) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/crypto/sha256/sha256.go:66
	_go_fuzz_dep_.CoverTab[10167]++
							b := make([]byte, 0, marshaledSize)
							if d.is224 {
//line /usr/local/go/src/crypto/sha256/sha256.go:68
		_go_fuzz_dep_.CoverTab[10169]++
								b = append(b, magic224...)
//line /usr/local/go/src/crypto/sha256/sha256.go:69
		// _ = "end of CoverTab[10169]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:70
		_go_fuzz_dep_.CoverTab[10170]++
								b = append(b, magic256...)
//line /usr/local/go/src/crypto/sha256/sha256.go:71
		// _ = "end of CoverTab[10170]"
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:72
	// _ = "end of CoverTab[10167]"
//line /usr/local/go/src/crypto/sha256/sha256.go:72
	_go_fuzz_dep_.CoverTab[10168]++
							b = binary.BigEndian.AppendUint32(b, d.h[0])
							b = binary.BigEndian.AppendUint32(b, d.h[1])
							b = binary.BigEndian.AppendUint32(b, d.h[2])
							b = binary.BigEndian.AppendUint32(b, d.h[3])
							b = binary.BigEndian.AppendUint32(b, d.h[4])
							b = binary.BigEndian.AppendUint32(b, d.h[5])
							b = binary.BigEndian.AppendUint32(b, d.h[6])
							b = binary.BigEndian.AppendUint32(b, d.h[7])
							b = append(b, d.x[:d.nx]...)
							b = b[:len(b)+len(d.x)-d.nx]
							b = binary.BigEndian.AppendUint64(b, d.len)
							return b, nil
//line /usr/local/go/src/crypto/sha256/sha256.go:84
	// _ = "end of CoverTab[10168]"
}

func (d *digest) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/crypto/sha256/sha256.go:87
	_go_fuzz_dep_.CoverTab[10171]++
							if len(b) < len(magic224) || func() bool {
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		_go_fuzz_dep_.CoverTab[10174]++
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		return (d.is224 && func() bool {
//line /usr/local/go/src/crypto/sha256/sha256.go:88
			_go_fuzz_dep_.CoverTab[10175]++
//line /usr/local/go/src/crypto/sha256/sha256.go:88
			return string(b[:len(magic224)]) != magic224
//line /usr/local/go/src/crypto/sha256/sha256.go:88
			// _ = "end of CoverTab[10175]"
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		}())
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		// _ = "end of CoverTab[10174]"
//line /usr/local/go/src/crypto/sha256/sha256.go:88
	}() || func() bool {
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		_go_fuzz_dep_.CoverTab[10176]++
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		return (!d.is224 && func() bool {
//line /usr/local/go/src/crypto/sha256/sha256.go:88
			_go_fuzz_dep_.CoverTab[10177]++
//line /usr/local/go/src/crypto/sha256/sha256.go:88
			return string(b[:len(magic256)]) != magic256
//line /usr/local/go/src/crypto/sha256/sha256.go:88
			// _ = "end of CoverTab[10177]"
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		}())
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		// _ = "end of CoverTab[10176]"
//line /usr/local/go/src/crypto/sha256/sha256.go:88
	}() {
//line /usr/local/go/src/crypto/sha256/sha256.go:88
		_go_fuzz_dep_.CoverTab[10178]++
								return errors.New("crypto/sha256: invalid hash state identifier")
//line /usr/local/go/src/crypto/sha256/sha256.go:89
		// _ = "end of CoverTab[10178]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:90
		_go_fuzz_dep_.CoverTab[10179]++
//line /usr/local/go/src/crypto/sha256/sha256.go:90
		// _ = "end of CoverTab[10179]"
//line /usr/local/go/src/crypto/sha256/sha256.go:90
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:90
	// _ = "end of CoverTab[10171]"
//line /usr/local/go/src/crypto/sha256/sha256.go:90
	_go_fuzz_dep_.CoverTab[10172]++
							if len(b) != marshaledSize {
//line /usr/local/go/src/crypto/sha256/sha256.go:91
		_go_fuzz_dep_.CoverTab[10180]++
								return errors.New("crypto/sha256: invalid hash state size")
//line /usr/local/go/src/crypto/sha256/sha256.go:92
		// _ = "end of CoverTab[10180]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:93
		_go_fuzz_dep_.CoverTab[10181]++
//line /usr/local/go/src/crypto/sha256/sha256.go:93
		// _ = "end of CoverTab[10181]"
//line /usr/local/go/src/crypto/sha256/sha256.go:93
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:93
	// _ = "end of CoverTab[10172]"
//line /usr/local/go/src/crypto/sha256/sha256.go:93
	_go_fuzz_dep_.CoverTab[10173]++
							b = b[len(magic224):]
							b, d.h[0] = consumeUint32(b)
							b, d.h[1] = consumeUint32(b)
							b, d.h[2] = consumeUint32(b)
							b, d.h[3] = consumeUint32(b)
							b, d.h[4] = consumeUint32(b)
							b, d.h[5] = consumeUint32(b)
							b, d.h[6] = consumeUint32(b)
							b, d.h[7] = consumeUint32(b)
							b = b[copy(d.x[:], b):]
							b, d.len = consumeUint64(b)
							d.nx = int(d.len % chunk)
							return nil
//line /usr/local/go/src/crypto/sha256/sha256.go:106
	// _ = "end of CoverTab[10173]"
}

func consumeUint64(b []byte) ([]byte, uint64) {
//line /usr/local/go/src/crypto/sha256/sha256.go:109
	_go_fuzz_dep_.CoverTab[10182]++
							_ = b[7]
							x := uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
							return b[8:], x
//line /usr/local/go/src/crypto/sha256/sha256.go:113
	// _ = "end of CoverTab[10182]"
}

func consumeUint32(b []byte) ([]byte, uint32) {
//line /usr/local/go/src/crypto/sha256/sha256.go:116
	_go_fuzz_dep_.CoverTab[10183]++
							_ = b[3]
							x := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
							return b[4:], x
//line /usr/local/go/src/crypto/sha256/sha256.go:119
	// _ = "end of CoverTab[10183]"
}

func (d *digest) Reset() {
//line /usr/local/go/src/crypto/sha256/sha256.go:122
	_go_fuzz_dep_.CoverTab[10184]++
							if !d.is224 {
//line /usr/local/go/src/crypto/sha256/sha256.go:123
		_go_fuzz_dep_.CoverTab[10186]++
								d.h[0] = init0
								d.h[1] = init1
								d.h[2] = init2
								d.h[3] = init3
								d.h[4] = init4
								d.h[5] = init5
								d.h[6] = init6
								d.h[7] = init7
//line /usr/local/go/src/crypto/sha256/sha256.go:131
		// _ = "end of CoverTab[10186]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:132
		_go_fuzz_dep_.CoverTab[10187]++
								d.h[0] = init0_224
								d.h[1] = init1_224
								d.h[2] = init2_224
								d.h[3] = init3_224
								d.h[4] = init4_224
								d.h[5] = init5_224
								d.h[6] = init6_224
								d.h[7] = init7_224
//line /usr/local/go/src/crypto/sha256/sha256.go:140
		// _ = "end of CoverTab[10187]"
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:141
	// _ = "end of CoverTab[10184]"
//line /usr/local/go/src/crypto/sha256/sha256.go:141
	_go_fuzz_dep_.CoverTab[10185]++
							d.nx = 0
							d.len = 0
//line /usr/local/go/src/crypto/sha256/sha256.go:143
	// _ = "end of CoverTab[10185]"
}

// New returns a new hash.Hash computing the SHA256 checksum. The Hash
//line /usr/local/go/src/crypto/sha256/sha256.go:146
// also implements encoding.BinaryMarshaler and
//line /usr/local/go/src/crypto/sha256/sha256.go:146
// encoding.BinaryUnmarshaler to marshal and unmarshal the internal
//line /usr/local/go/src/crypto/sha256/sha256.go:146
// state of the hash.
//line /usr/local/go/src/crypto/sha256/sha256.go:150
func New() hash.Hash {
//line /usr/local/go/src/crypto/sha256/sha256.go:150
	_go_fuzz_dep_.CoverTab[10188]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha256/sha256.go:151
		_go_fuzz_dep_.CoverTab[10190]++
								return boring.NewSHA256()
//line /usr/local/go/src/crypto/sha256/sha256.go:152
		// _ = "end of CoverTab[10190]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:153
		_go_fuzz_dep_.CoverTab[10191]++
//line /usr/local/go/src/crypto/sha256/sha256.go:153
		// _ = "end of CoverTab[10191]"
//line /usr/local/go/src/crypto/sha256/sha256.go:153
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:153
	// _ = "end of CoverTab[10188]"
//line /usr/local/go/src/crypto/sha256/sha256.go:153
	_go_fuzz_dep_.CoverTab[10189]++
							d := new(digest)
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha256/sha256.go:156
	// _ = "end of CoverTab[10189]"
}

// New224 returns a new hash.Hash computing the SHA224 checksum.
func New224() hash.Hash {
//line /usr/local/go/src/crypto/sha256/sha256.go:160
	_go_fuzz_dep_.CoverTab[10192]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha256/sha256.go:161
		_go_fuzz_dep_.CoverTab[10194]++
								return boring.NewSHA224()
//line /usr/local/go/src/crypto/sha256/sha256.go:162
		// _ = "end of CoverTab[10194]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:163
		_go_fuzz_dep_.CoverTab[10195]++
//line /usr/local/go/src/crypto/sha256/sha256.go:163
		// _ = "end of CoverTab[10195]"
//line /usr/local/go/src/crypto/sha256/sha256.go:163
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:163
	// _ = "end of CoverTab[10192]"
//line /usr/local/go/src/crypto/sha256/sha256.go:163
	_go_fuzz_dep_.CoverTab[10193]++
							d := new(digest)
							d.is224 = true
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha256/sha256.go:167
	// _ = "end of CoverTab[10193]"
}

func (d *digest) Size() int {
//line /usr/local/go/src/crypto/sha256/sha256.go:170
	_go_fuzz_dep_.CoverTab[10196]++
							if !d.is224 {
//line /usr/local/go/src/crypto/sha256/sha256.go:171
		_go_fuzz_dep_.CoverTab[10198]++
								return Size
//line /usr/local/go/src/crypto/sha256/sha256.go:172
		// _ = "end of CoverTab[10198]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:173
		_go_fuzz_dep_.CoverTab[10199]++
//line /usr/local/go/src/crypto/sha256/sha256.go:173
		// _ = "end of CoverTab[10199]"
//line /usr/local/go/src/crypto/sha256/sha256.go:173
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:173
	// _ = "end of CoverTab[10196]"
//line /usr/local/go/src/crypto/sha256/sha256.go:173
	_go_fuzz_dep_.CoverTab[10197]++
							return Size224
//line /usr/local/go/src/crypto/sha256/sha256.go:174
	// _ = "end of CoverTab[10197]"
}

func (d *digest) BlockSize() int {
//line /usr/local/go/src/crypto/sha256/sha256.go:177
	_go_fuzz_dep_.CoverTab[10200]++
//line /usr/local/go/src/crypto/sha256/sha256.go:177
	return BlockSize
//line /usr/local/go/src/crypto/sha256/sha256.go:177
	// _ = "end of CoverTab[10200]"
//line /usr/local/go/src/crypto/sha256/sha256.go:177
}

func (d *digest) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/crypto/sha256/sha256.go:179
	_go_fuzz_dep_.CoverTab[10201]++
							boring.Unreachable()
							nn = len(p)
							d.len += uint64(nn)
							if d.nx > 0 {
//line /usr/local/go/src/crypto/sha256/sha256.go:183
		_go_fuzz_dep_.CoverTab[10205]++
								n := copy(d.x[d.nx:], p)
								d.nx += n
								if d.nx == chunk {
//line /usr/local/go/src/crypto/sha256/sha256.go:186
			_go_fuzz_dep_.CoverTab[10207]++
									block(d, d.x[:])
									d.nx = 0
//line /usr/local/go/src/crypto/sha256/sha256.go:188
			// _ = "end of CoverTab[10207]"
		} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:189
			_go_fuzz_dep_.CoverTab[10208]++
//line /usr/local/go/src/crypto/sha256/sha256.go:189
			// _ = "end of CoverTab[10208]"
//line /usr/local/go/src/crypto/sha256/sha256.go:189
		}
//line /usr/local/go/src/crypto/sha256/sha256.go:189
		// _ = "end of CoverTab[10205]"
//line /usr/local/go/src/crypto/sha256/sha256.go:189
		_go_fuzz_dep_.CoverTab[10206]++
								p = p[n:]
//line /usr/local/go/src/crypto/sha256/sha256.go:190
		// _ = "end of CoverTab[10206]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:191
		_go_fuzz_dep_.CoverTab[10209]++
//line /usr/local/go/src/crypto/sha256/sha256.go:191
		// _ = "end of CoverTab[10209]"
//line /usr/local/go/src/crypto/sha256/sha256.go:191
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:191
	// _ = "end of CoverTab[10201]"
//line /usr/local/go/src/crypto/sha256/sha256.go:191
	_go_fuzz_dep_.CoverTab[10202]++
							if len(p) >= chunk {
//line /usr/local/go/src/crypto/sha256/sha256.go:192
		_go_fuzz_dep_.CoverTab[10210]++
								n := len(p) &^ (chunk - 1)
								block(d, p[:n])
								p = p[n:]
//line /usr/local/go/src/crypto/sha256/sha256.go:195
		// _ = "end of CoverTab[10210]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:196
		_go_fuzz_dep_.CoverTab[10211]++
//line /usr/local/go/src/crypto/sha256/sha256.go:196
		// _ = "end of CoverTab[10211]"
//line /usr/local/go/src/crypto/sha256/sha256.go:196
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:196
	// _ = "end of CoverTab[10202]"
//line /usr/local/go/src/crypto/sha256/sha256.go:196
	_go_fuzz_dep_.CoverTab[10203]++
							if len(p) > 0 {
//line /usr/local/go/src/crypto/sha256/sha256.go:197
		_go_fuzz_dep_.CoverTab[10212]++
								d.nx = copy(d.x[:], p)
//line /usr/local/go/src/crypto/sha256/sha256.go:198
		// _ = "end of CoverTab[10212]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:199
		_go_fuzz_dep_.CoverTab[10213]++
//line /usr/local/go/src/crypto/sha256/sha256.go:199
		// _ = "end of CoverTab[10213]"
//line /usr/local/go/src/crypto/sha256/sha256.go:199
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:199
	// _ = "end of CoverTab[10203]"
//line /usr/local/go/src/crypto/sha256/sha256.go:199
	_go_fuzz_dep_.CoverTab[10204]++
							return
//line /usr/local/go/src/crypto/sha256/sha256.go:200
	// _ = "end of CoverTab[10204]"
}

func (d *digest) Sum(in []byte) []byte {
//line /usr/local/go/src/crypto/sha256/sha256.go:203
	_go_fuzz_dep_.CoverTab[10214]++
							boring.Unreachable()

							d0 := *d
							hash := d0.checkSum()
							if d0.is224 {
//line /usr/local/go/src/crypto/sha256/sha256.go:208
		_go_fuzz_dep_.CoverTab[10216]++
								return append(in, hash[:Size224]...)
//line /usr/local/go/src/crypto/sha256/sha256.go:209
		// _ = "end of CoverTab[10216]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:210
		_go_fuzz_dep_.CoverTab[10217]++
//line /usr/local/go/src/crypto/sha256/sha256.go:210
		// _ = "end of CoverTab[10217]"
//line /usr/local/go/src/crypto/sha256/sha256.go:210
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:210
	// _ = "end of CoverTab[10214]"
//line /usr/local/go/src/crypto/sha256/sha256.go:210
	_go_fuzz_dep_.CoverTab[10215]++
							return append(in, hash[:]...)
//line /usr/local/go/src/crypto/sha256/sha256.go:211
	// _ = "end of CoverTab[10215]"
}

func (d *digest) checkSum() [Size]byte {
//line /usr/local/go/src/crypto/sha256/sha256.go:214
	_go_fuzz_dep_.CoverTab[10218]++
							len := d.len
	// Padding. Add a 1 bit and 0 bits until 56 bytes mod 64.
	var tmp [64 + 8]byte	// padding + length buffer
	tmp[0] = 0x80
	var t uint64
	if len%64 < 56 {
//line /usr/local/go/src/crypto/sha256/sha256.go:220
		_go_fuzz_dep_.CoverTab[10222]++
								t = 56 - len%64
//line /usr/local/go/src/crypto/sha256/sha256.go:221
		// _ = "end of CoverTab[10222]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:222
		_go_fuzz_dep_.CoverTab[10223]++
								t = 64 + 56 - len%64
//line /usr/local/go/src/crypto/sha256/sha256.go:223
		// _ = "end of CoverTab[10223]"
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:224
	// _ = "end of CoverTab[10218]"
//line /usr/local/go/src/crypto/sha256/sha256.go:224
	_go_fuzz_dep_.CoverTab[10219]++

//line /usr/local/go/src/crypto/sha256/sha256.go:227
	len <<= 3
	padlen := tmp[:t+8]
	binary.BigEndian.PutUint64(padlen[t+0:], len)
	d.Write(padlen)

	if d.nx != 0 {
//line /usr/local/go/src/crypto/sha256/sha256.go:232
		_go_fuzz_dep_.CoverTab[10224]++
								panic("d.nx != 0")
//line /usr/local/go/src/crypto/sha256/sha256.go:233
		// _ = "end of CoverTab[10224]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:234
		_go_fuzz_dep_.CoverTab[10225]++
//line /usr/local/go/src/crypto/sha256/sha256.go:234
		// _ = "end of CoverTab[10225]"
//line /usr/local/go/src/crypto/sha256/sha256.go:234
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:234
	// _ = "end of CoverTab[10219]"
//line /usr/local/go/src/crypto/sha256/sha256.go:234
	_go_fuzz_dep_.CoverTab[10220]++

							var digest [Size]byte

							binary.BigEndian.PutUint32(digest[0:], d.h[0])
							binary.BigEndian.PutUint32(digest[4:], d.h[1])
							binary.BigEndian.PutUint32(digest[8:], d.h[2])
							binary.BigEndian.PutUint32(digest[12:], d.h[3])
							binary.BigEndian.PutUint32(digest[16:], d.h[4])
							binary.BigEndian.PutUint32(digest[20:], d.h[5])
							binary.BigEndian.PutUint32(digest[24:], d.h[6])
							if !d.is224 {
//line /usr/local/go/src/crypto/sha256/sha256.go:245
		_go_fuzz_dep_.CoverTab[10226]++
								binary.BigEndian.PutUint32(digest[28:], d.h[7])
//line /usr/local/go/src/crypto/sha256/sha256.go:246
		// _ = "end of CoverTab[10226]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:247
		_go_fuzz_dep_.CoverTab[10227]++
//line /usr/local/go/src/crypto/sha256/sha256.go:247
		// _ = "end of CoverTab[10227]"
//line /usr/local/go/src/crypto/sha256/sha256.go:247
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:247
	// _ = "end of CoverTab[10220]"
//line /usr/local/go/src/crypto/sha256/sha256.go:247
	_go_fuzz_dep_.CoverTab[10221]++

							return digest
//line /usr/local/go/src/crypto/sha256/sha256.go:249
	// _ = "end of CoverTab[10221]"
}

// Sum256 returns the SHA256 checksum of the data.
func Sum256(data []byte) [Size]byte {
//line /usr/local/go/src/crypto/sha256/sha256.go:253
	_go_fuzz_dep_.CoverTab[10228]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha256/sha256.go:254
		_go_fuzz_dep_.CoverTab[10230]++
								return boring.SHA256(data)
//line /usr/local/go/src/crypto/sha256/sha256.go:255
		// _ = "end of CoverTab[10230]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:256
		_go_fuzz_dep_.CoverTab[10231]++
//line /usr/local/go/src/crypto/sha256/sha256.go:256
		// _ = "end of CoverTab[10231]"
//line /usr/local/go/src/crypto/sha256/sha256.go:256
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:256
	// _ = "end of CoverTab[10228]"
//line /usr/local/go/src/crypto/sha256/sha256.go:256
	_go_fuzz_dep_.CoverTab[10229]++
							var d digest
							d.Reset()
							d.Write(data)
							return d.checkSum()
//line /usr/local/go/src/crypto/sha256/sha256.go:260
	// _ = "end of CoverTab[10229]"
}

// Sum224 returns the SHA224 checksum of the data.
func Sum224(data []byte) [Size224]byte {
//line /usr/local/go/src/crypto/sha256/sha256.go:264
	_go_fuzz_dep_.CoverTab[10232]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha256/sha256.go:265
		_go_fuzz_dep_.CoverTab[10234]++
								return boring.SHA224(data)
//line /usr/local/go/src/crypto/sha256/sha256.go:266
		// _ = "end of CoverTab[10234]"
	} else {
//line /usr/local/go/src/crypto/sha256/sha256.go:267
		_go_fuzz_dep_.CoverTab[10235]++
//line /usr/local/go/src/crypto/sha256/sha256.go:267
		// _ = "end of CoverTab[10235]"
//line /usr/local/go/src/crypto/sha256/sha256.go:267
	}
//line /usr/local/go/src/crypto/sha256/sha256.go:267
	// _ = "end of CoverTab[10232]"
//line /usr/local/go/src/crypto/sha256/sha256.go:267
	_go_fuzz_dep_.CoverTab[10233]++
							var d digest
							d.is224 = true
							d.Reset()
							d.Write(data)
							sum := d.checkSum()
							ap := (*[Size224]byte)(sum[:])
							return *ap
//line /usr/local/go/src/crypto/sha256/sha256.go:274
	// _ = "end of CoverTab[10233]"
}

//line /usr/local/go/src/crypto/sha256/sha256.go:275
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha256/sha256.go:275
var _ = _go_fuzz_dep_.CoverTab
