// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/sha512/sha512.go:5
// Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256
//line /usr/local/go/src/crypto/sha512/sha512.go:5
// hash algorithms as defined in FIPS 180-4.
//line /usr/local/go/src/crypto/sha512/sha512.go:5
//
//line /usr/local/go/src/crypto/sha512/sha512.go:5
// All the hash.Hash implementations returned by this package also
//line /usr/local/go/src/crypto/sha512/sha512.go:5
// implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
//line /usr/local/go/src/crypto/sha512/sha512.go:5
// marshal and unmarshal the internal state of the hash.
//line /usr/local/go/src/crypto/sha512/sha512.go:11
package sha512

//line /usr/local/go/src/crypto/sha512/sha512.go:11
import (
//line /usr/local/go/src/crypto/sha512/sha512.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha512/sha512.go:11
)
//line /usr/local/go/src/crypto/sha512/sha512.go:11
import (
//line /usr/local/go/src/crypto/sha512/sha512.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha512/sha512.go:11
)

import (
	"crypto"
	"crypto/internal/boring"
	"encoding/binary"
	"errors"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA384, New384)
	crypto.RegisterHash(crypto.SHA512, New)
	crypto.RegisterHash(crypto.SHA512_224, New512_224)
	crypto.RegisterHash(crypto.SHA512_256, New512_256)
}

const (
	// Size is the size, in bytes, of a SHA-512 checksum.
	Size	= 64

	// Size224 is the size, in bytes, of a SHA-512/224 checksum.
	Size224	= 28

	// Size256 is the size, in bytes, of a SHA-512/256 checksum.
	Size256	= 32

	// Size384 is the size, in bytes, of a SHA-384 checksum.
	Size384	= 48

	// BlockSize is the block size, in bytes, of the SHA-512/224,
	// SHA-512/256, SHA-384 and SHA-512 hash functions.
	BlockSize	= 128
)

const (
	chunk		= 128
	init0		= 0x6a09e667f3bcc908
	init1		= 0xbb67ae8584caa73b
	init2		= 0x3c6ef372fe94f82b
	init3		= 0xa54ff53a5f1d36f1
	init4		= 0x510e527fade682d1
	init5		= 0x9b05688c2b3e6c1f
	init6		= 0x1f83d9abfb41bd6b
	init7		= 0x5be0cd19137e2179
	init0_224	= 0x8c3d37c819544da2
	init1_224	= 0x73e1996689dcd4d6
	init2_224	= 0x1dfab7ae32ff9c82
	init3_224	= 0x679dd514582f9fcf
	init4_224	= 0x0f6d2b697bd44da8
	init5_224	= 0x77e36f7304c48942
	init6_224	= 0x3f9d85a86a1d36c8
	init7_224	= 0x1112e6ad91d692a1
	init0_256	= 0x22312194fc2bf72c
	init1_256	= 0x9f555fa3c84c64c2
	init2_256	= 0x2393b86b6f53b151
	init3_256	= 0x963877195940eabd
	init4_256	= 0x96283ee2a88effe3
	init5_256	= 0xbe5e1e2553863992
	init6_256	= 0x2b0199fc2c85b8aa
	init7_256	= 0x0eb72ddc81c52ca2
	init0_384	= 0xcbbb9d5dc1059ed8
	init1_384	= 0x629a292a367cd507
	init2_384	= 0x9159015a3070dd17
	init3_384	= 0x152fecd8f70e5939
	init4_384	= 0x67332667ffc00b31
	init5_384	= 0x8eb44a8768581511
	init6_384	= 0xdb0c2e0d64f98fa7
	init7_384	= 0x47b5481dbefa4fa4
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	h		[8]uint64
	x		[chunk]byte
	nx		int
	len		uint64
	function	crypto.Hash
}

func (d *digest) Reset() {
//line /usr/local/go/src/crypto/sha512/sha512.go:91
	_go_fuzz_dep_.CoverTab[7333]++
							switch d.function {
	case crypto.SHA384:
//line /usr/local/go/src/crypto/sha512/sha512.go:93
		_go_fuzz_dep_.CoverTab[7335]++
								d.h[0] = init0_384
								d.h[1] = init1_384
								d.h[2] = init2_384
								d.h[3] = init3_384
								d.h[4] = init4_384
								d.h[5] = init5_384
								d.h[6] = init6_384
								d.h[7] = init7_384
//line /usr/local/go/src/crypto/sha512/sha512.go:101
		// _ = "end of CoverTab[7335]"
	case crypto.SHA512_224:
//line /usr/local/go/src/crypto/sha512/sha512.go:102
		_go_fuzz_dep_.CoverTab[7336]++
								d.h[0] = init0_224
								d.h[1] = init1_224
								d.h[2] = init2_224
								d.h[3] = init3_224
								d.h[4] = init4_224
								d.h[5] = init5_224
								d.h[6] = init6_224
								d.h[7] = init7_224
//line /usr/local/go/src/crypto/sha512/sha512.go:110
		// _ = "end of CoverTab[7336]"
	case crypto.SHA512_256:
//line /usr/local/go/src/crypto/sha512/sha512.go:111
		_go_fuzz_dep_.CoverTab[7337]++
								d.h[0] = init0_256
								d.h[1] = init1_256
								d.h[2] = init2_256
								d.h[3] = init3_256
								d.h[4] = init4_256
								d.h[5] = init5_256
								d.h[6] = init6_256
								d.h[7] = init7_256
//line /usr/local/go/src/crypto/sha512/sha512.go:119
		// _ = "end of CoverTab[7337]"
	default:
//line /usr/local/go/src/crypto/sha512/sha512.go:120
		_go_fuzz_dep_.CoverTab[7338]++
								d.h[0] = init0
								d.h[1] = init1
								d.h[2] = init2
								d.h[3] = init3
								d.h[4] = init4
								d.h[5] = init5
								d.h[6] = init6
								d.h[7] = init7
//line /usr/local/go/src/crypto/sha512/sha512.go:128
		// _ = "end of CoverTab[7338]"
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:129
	// _ = "end of CoverTab[7333]"
//line /usr/local/go/src/crypto/sha512/sha512.go:129
	_go_fuzz_dep_.CoverTab[7334]++
							d.nx = 0
							d.len = 0
//line /usr/local/go/src/crypto/sha512/sha512.go:131
	// _ = "end of CoverTab[7334]"
}

const (
	magic384	= "sha\x04"
	magic512_224	= "sha\x05"
	magic512_256	= "sha\x06"
	magic512	= "sha\x07"
	marshaledSize	= len(magic512) + 8*8 + chunk + 8
)

func (d *digest) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/crypto/sha512/sha512.go:142
	_go_fuzz_dep_.CoverTab[7339]++
							b := make([]byte, 0, marshaledSize)
							switch d.function {
	case crypto.SHA384:
//line /usr/local/go/src/crypto/sha512/sha512.go:145
		_go_fuzz_dep_.CoverTab[7341]++
								b = append(b, magic384...)
//line /usr/local/go/src/crypto/sha512/sha512.go:146
		// _ = "end of CoverTab[7341]"
	case crypto.SHA512_224:
//line /usr/local/go/src/crypto/sha512/sha512.go:147
		_go_fuzz_dep_.CoverTab[7342]++
								b = append(b, magic512_224...)
//line /usr/local/go/src/crypto/sha512/sha512.go:148
		// _ = "end of CoverTab[7342]"
	case crypto.SHA512_256:
//line /usr/local/go/src/crypto/sha512/sha512.go:149
		_go_fuzz_dep_.CoverTab[7343]++
								b = append(b, magic512_256...)
//line /usr/local/go/src/crypto/sha512/sha512.go:150
		// _ = "end of CoverTab[7343]"
	case crypto.SHA512:
//line /usr/local/go/src/crypto/sha512/sha512.go:151
		_go_fuzz_dep_.CoverTab[7344]++
								b = append(b, magic512...)
//line /usr/local/go/src/crypto/sha512/sha512.go:152
		// _ = "end of CoverTab[7344]"
	default:
//line /usr/local/go/src/crypto/sha512/sha512.go:153
		_go_fuzz_dep_.CoverTab[7345]++
								return nil, errors.New("crypto/sha512: invalid hash function")
//line /usr/local/go/src/crypto/sha512/sha512.go:154
		// _ = "end of CoverTab[7345]"
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:155
	// _ = "end of CoverTab[7339]"
//line /usr/local/go/src/crypto/sha512/sha512.go:155
	_go_fuzz_dep_.CoverTab[7340]++
							b = binary.BigEndian.AppendUint64(b, d.h[0])
							b = binary.BigEndian.AppendUint64(b, d.h[1])
							b = binary.BigEndian.AppendUint64(b, d.h[2])
							b = binary.BigEndian.AppendUint64(b, d.h[3])
							b = binary.BigEndian.AppendUint64(b, d.h[4])
							b = binary.BigEndian.AppendUint64(b, d.h[5])
							b = binary.BigEndian.AppendUint64(b, d.h[6])
							b = binary.BigEndian.AppendUint64(b, d.h[7])
							b = append(b, d.x[:d.nx]...)
							b = b[:len(b)+len(d.x)-d.nx]
							b = binary.BigEndian.AppendUint64(b, d.len)
							return b, nil
//line /usr/local/go/src/crypto/sha512/sha512.go:167
	// _ = "end of CoverTab[7340]"
}

func (d *digest) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/crypto/sha512/sha512.go:170
	_go_fuzz_dep_.CoverTab[7346]++
							if len(b) < len(magic512) {
//line /usr/local/go/src/crypto/sha512/sha512.go:171
		_go_fuzz_dep_.CoverTab[7350]++
								return errors.New("crypto/sha512: invalid hash state identifier")
//line /usr/local/go/src/crypto/sha512/sha512.go:172
		// _ = "end of CoverTab[7350]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:173
		_go_fuzz_dep_.CoverTab[7351]++
//line /usr/local/go/src/crypto/sha512/sha512.go:173
		// _ = "end of CoverTab[7351]"
//line /usr/local/go/src/crypto/sha512/sha512.go:173
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:173
	// _ = "end of CoverTab[7346]"
//line /usr/local/go/src/crypto/sha512/sha512.go:173
	_go_fuzz_dep_.CoverTab[7347]++
							switch {
	case d.function == crypto.SHA384 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512.go:175
		_go_fuzz_dep_.CoverTab[7357]++
//line /usr/local/go/src/crypto/sha512/sha512.go:175
		return string(b[:len(magic384)]) == magic384
//line /usr/local/go/src/crypto/sha512/sha512.go:175
		// _ = "end of CoverTab[7357]"
//line /usr/local/go/src/crypto/sha512/sha512.go:175
	}():
//line /usr/local/go/src/crypto/sha512/sha512.go:175
		_go_fuzz_dep_.CoverTab[7352]++
//line /usr/local/go/src/crypto/sha512/sha512.go:175
		// _ = "end of CoverTab[7352]"
	case d.function == crypto.SHA512_224 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512.go:176
		_go_fuzz_dep_.CoverTab[7358]++
//line /usr/local/go/src/crypto/sha512/sha512.go:176
		return string(b[:len(magic512_224)]) == magic512_224
//line /usr/local/go/src/crypto/sha512/sha512.go:176
		// _ = "end of CoverTab[7358]"
//line /usr/local/go/src/crypto/sha512/sha512.go:176
	}():
//line /usr/local/go/src/crypto/sha512/sha512.go:176
		_go_fuzz_dep_.CoverTab[7353]++
//line /usr/local/go/src/crypto/sha512/sha512.go:176
		// _ = "end of CoverTab[7353]"
	case d.function == crypto.SHA512_256 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512.go:177
		_go_fuzz_dep_.CoverTab[7359]++
//line /usr/local/go/src/crypto/sha512/sha512.go:177
		return string(b[:len(magic512_256)]) == magic512_256
//line /usr/local/go/src/crypto/sha512/sha512.go:177
		// _ = "end of CoverTab[7359]"
//line /usr/local/go/src/crypto/sha512/sha512.go:177
	}():
//line /usr/local/go/src/crypto/sha512/sha512.go:177
		_go_fuzz_dep_.CoverTab[7354]++
//line /usr/local/go/src/crypto/sha512/sha512.go:177
		// _ = "end of CoverTab[7354]"
	case d.function == crypto.SHA512 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512.go:178
		_go_fuzz_dep_.CoverTab[7360]++
//line /usr/local/go/src/crypto/sha512/sha512.go:178
		return string(b[:len(magic512)]) == magic512
//line /usr/local/go/src/crypto/sha512/sha512.go:178
		// _ = "end of CoverTab[7360]"
//line /usr/local/go/src/crypto/sha512/sha512.go:178
	}():
//line /usr/local/go/src/crypto/sha512/sha512.go:178
		_go_fuzz_dep_.CoverTab[7355]++
//line /usr/local/go/src/crypto/sha512/sha512.go:178
		// _ = "end of CoverTab[7355]"
	default:
//line /usr/local/go/src/crypto/sha512/sha512.go:179
		_go_fuzz_dep_.CoverTab[7356]++
								return errors.New("crypto/sha512: invalid hash state identifier")
//line /usr/local/go/src/crypto/sha512/sha512.go:180
		// _ = "end of CoverTab[7356]"
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:181
	// _ = "end of CoverTab[7347]"
//line /usr/local/go/src/crypto/sha512/sha512.go:181
	_go_fuzz_dep_.CoverTab[7348]++
							if len(b) != marshaledSize {
//line /usr/local/go/src/crypto/sha512/sha512.go:182
		_go_fuzz_dep_.CoverTab[7361]++
								return errors.New("crypto/sha512: invalid hash state size")
//line /usr/local/go/src/crypto/sha512/sha512.go:183
		// _ = "end of CoverTab[7361]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:184
		_go_fuzz_dep_.CoverTab[7362]++
//line /usr/local/go/src/crypto/sha512/sha512.go:184
		// _ = "end of CoverTab[7362]"
//line /usr/local/go/src/crypto/sha512/sha512.go:184
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:184
	// _ = "end of CoverTab[7348]"
//line /usr/local/go/src/crypto/sha512/sha512.go:184
	_go_fuzz_dep_.CoverTab[7349]++
							b = b[len(magic512):]
							b, d.h[0] = consumeUint64(b)
							b, d.h[1] = consumeUint64(b)
							b, d.h[2] = consumeUint64(b)
							b, d.h[3] = consumeUint64(b)
							b, d.h[4] = consumeUint64(b)
							b, d.h[5] = consumeUint64(b)
							b, d.h[6] = consumeUint64(b)
							b, d.h[7] = consumeUint64(b)
							b = b[copy(d.x[:], b):]
							b, d.len = consumeUint64(b)
							d.nx = int(d.len % chunk)
							return nil
//line /usr/local/go/src/crypto/sha512/sha512.go:197
	// _ = "end of CoverTab[7349]"
}

func consumeUint64(b []byte) ([]byte, uint64) {
//line /usr/local/go/src/crypto/sha512/sha512.go:200
	_go_fuzz_dep_.CoverTab[7363]++
							_ = b[7]
							x := uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
							return b[8:], x
//line /usr/local/go/src/crypto/sha512/sha512.go:204
	// _ = "end of CoverTab[7363]"
}

// New returns a new hash.Hash computing the SHA-512 checksum.
func New() hash.Hash {
//line /usr/local/go/src/crypto/sha512/sha512.go:208
	_go_fuzz_dep_.CoverTab[7364]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha512/sha512.go:209
		_go_fuzz_dep_.CoverTab[7366]++
								return boring.NewSHA512()
//line /usr/local/go/src/crypto/sha512/sha512.go:210
		// _ = "end of CoverTab[7366]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:211
		_go_fuzz_dep_.CoverTab[7367]++
//line /usr/local/go/src/crypto/sha512/sha512.go:211
		// _ = "end of CoverTab[7367]"
//line /usr/local/go/src/crypto/sha512/sha512.go:211
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:211
	// _ = "end of CoverTab[7364]"
//line /usr/local/go/src/crypto/sha512/sha512.go:211
	_go_fuzz_dep_.CoverTab[7365]++
							d := &digest{function: crypto.SHA512}
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha512/sha512.go:214
	// _ = "end of CoverTab[7365]"
}

// New512_224 returns a new hash.Hash computing the SHA-512/224 checksum.
func New512_224() hash.Hash {
//line /usr/local/go/src/crypto/sha512/sha512.go:218
	_go_fuzz_dep_.CoverTab[7368]++
							d := &digest{function: crypto.SHA512_224}
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha512/sha512.go:221
	// _ = "end of CoverTab[7368]"
}

// New512_256 returns a new hash.Hash computing the SHA-512/256 checksum.
func New512_256() hash.Hash {
//line /usr/local/go/src/crypto/sha512/sha512.go:225
	_go_fuzz_dep_.CoverTab[7369]++
							d := &digest{function: crypto.SHA512_256}
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha512/sha512.go:228
	// _ = "end of CoverTab[7369]"
}

// New384 returns a new hash.Hash computing the SHA-384 checksum.
func New384() hash.Hash {
//line /usr/local/go/src/crypto/sha512/sha512.go:232
	_go_fuzz_dep_.CoverTab[7370]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha512/sha512.go:233
		_go_fuzz_dep_.CoverTab[7372]++
								return boring.NewSHA384()
//line /usr/local/go/src/crypto/sha512/sha512.go:234
		// _ = "end of CoverTab[7372]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:235
		_go_fuzz_dep_.CoverTab[7373]++
//line /usr/local/go/src/crypto/sha512/sha512.go:235
		// _ = "end of CoverTab[7373]"
//line /usr/local/go/src/crypto/sha512/sha512.go:235
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:235
	// _ = "end of CoverTab[7370]"
//line /usr/local/go/src/crypto/sha512/sha512.go:235
	_go_fuzz_dep_.CoverTab[7371]++
							d := &digest{function: crypto.SHA384}
							d.Reset()
							return d
//line /usr/local/go/src/crypto/sha512/sha512.go:238
	// _ = "end of CoverTab[7371]"
}

func (d *digest) Size() int {
//line /usr/local/go/src/crypto/sha512/sha512.go:241
	_go_fuzz_dep_.CoverTab[7374]++
							switch d.function {
	case crypto.SHA512_224:
//line /usr/local/go/src/crypto/sha512/sha512.go:243
		_go_fuzz_dep_.CoverTab[7375]++
								return Size224
//line /usr/local/go/src/crypto/sha512/sha512.go:244
		// _ = "end of CoverTab[7375]"
	case crypto.SHA512_256:
//line /usr/local/go/src/crypto/sha512/sha512.go:245
		_go_fuzz_dep_.CoverTab[7376]++
								return Size256
//line /usr/local/go/src/crypto/sha512/sha512.go:246
		// _ = "end of CoverTab[7376]"
	case crypto.SHA384:
//line /usr/local/go/src/crypto/sha512/sha512.go:247
		_go_fuzz_dep_.CoverTab[7377]++
								return Size384
//line /usr/local/go/src/crypto/sha512/sha512.go:248
		// _ = "end of CoverTab[7377]"
	default:
//line /usr/local/go/src/crypto/sha512/sha512.go:249
		_go_fuzz_dep_.CoverTab[7378]++
								return Size
//line /usr/local/go/src/crypto/sha512/sha512.go:250
		// _ = "end of CoverTab[7378]"
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:251
	// _ = "end of CoverTab[7374]"
}

func (d *digest) BlockSize() int {
//line /usr/local/go/src/crypto/sha512/sha512.go:254
	_go_fuzz_dep_.CoverTab[7379]++
//line /usr/local/go/src/crypto/sha512/sha512.go:254
	return BlockSize
//line /usr/local/go/src/crypto/sha512/sha512.go:254
	// _ = "end of CoverTab[7379]"
//line /usr/local/go/src/crypto/sha512/sha512.go:254
}

func (d *digest) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/crypto/sha512/sha512.go:256
	_go_fuzz_dep_.CoverTab[7380]++
							if d.function != crypto.SHA512_224 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512.go:257
		_go_fuzz_dep_.CoverTab[7385]++
//line /usr/local/go/src/crypto/sha512/sha512.go:257
		return d.function != crypto.SHA512_256
//line /usr/local/go/src/crypto/sha512/sha512.go:257
		// _ = "end of CoverTab[7385]"
//line /usr/local/go/src/crypto/sha512/sha512.go:257
	}() {
//line /usr/local/go/src/crypto/sha512/sha512.go:257
		_go_fuzz_dep_.CoverTab[7386]++
								boring.Unreachable()
//line /usr/local/go/src/crypto/sha512/sha512.go:258
		// _ = "end of CoverTab[7386]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:259
		_go_fuzz_dep_.CoverTab[7387]++
//line /usr/local/go/src/crypto/sha512/sha512.go:259
		// _ = "end of CoverTab[7387]"
//line /usr/local/go/src/crypto/sha512/sha512.go:259
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:259
	// _ = "end of CoverTab[7380]"
//line /usr/local/go/src/crypto/sha512/sha512.go:259
	_go_fuzz_dep_.CoverTab[7381]++
							nn = len(p)
							d.len += uint64(nn)
							if d.nx > 0 {
//line /usr/local/go/src/crypto/sha512/sha512.go:262
		_go_fuzz_dep_.CoverTab[7388]++
								n := copy(d.x[d.nx:], p)
								d.nx += n
								if d.nx == chunk {
//line /usr/local/go/src/crypto/sha512/sha512.go:265
			_go_fuzz_dep_.CoverTab[7390]++
									block(d, d.x[:])
									d.nx = 0
//line /usr/local/go/src/crypto/sha512/sha512.go:267
			// _ = "end of CoverTab[7390]"
		} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:268
			_go_fuzz_dep_.CoverTab[7391]++
//line /usr/local/go/src/crypto/sha512/sha512.go:268
			// _ = "end of CoverTab[7391]"
//line /usr/local/go/src/crypto/sha512/sha512.go:268
		}
//line /usr/local/go/src/crypto/sha512/sha512.go:268
		// _ = "end of CoverTab[7388]"
//line /usr/local/go/src/crypto/sha512/sha512.go:268
		_go_fuzz_dep_.CoverTab[7389]++
								p = p[n:]
//line /usr/local/go/src/crypto/sha512/sha512.go:269
		// _ = "end of CoverTab[7389]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:270
		_go_fuzz_dep_.CoverTab[7392]++
//line /usr/local/go/src/crypto/sha512/sha512.go:270
		// _ = "end of CoverTab[7392]"
//line /usr/local/go/src/crypto/sha512/sha512.go:270
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:270
	// _ = "end of CoverTab[7381]"
//line /usr/local/go/src/crypto/sha512/sha512.go:270
	_go_fuzz_dep_.CoverTab[7382]++
							if len(p) >= chunk {
//line /usr/local/go/src/crypto/sha512/sha512.go:271
		_go_fuzz_dep_.CoverTab[7393]++
								n := len(p) &^ (chunk - 1)
								block(d, p[:n])
								p = p[n:]
//line /usr/local/go/src/crypto/sha512/sha512.go:274
		// _ = "end of CoverTab[7393]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:275
		_go_fuzz_dep_.CoverTab[7394]++
//line /usr/local/go/src/crypto/sha512/sha512.go:275
		// _ = "end of CoverTab[7394]"
//line /usr/local/go/src/crypto/sha512/sha512.go:275
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:275
	// _ = "end of CoverTab[7382]"
//line /usr/local/go/src/crypto/sha512/sha512.go:275
	_go_fuzz_dep_.CoverTab[7383]++
							if len(p) > 0 {
//line /usr/local/go/src/crypto/sha512/sha512.go:276
		_go_fuzz_dep_.CoverTab[7395]++
								d.nx = copy(d.x[:], p)
//line /usr/local/go/src/crypto/sha512/sha512.go:277
		// _ = "end of CoverTab[7395]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:278
		_go_fuzz_dep_.CoverTab[7396]++
//line /usr/local/go/src/crypto/sha512/sha512.go:278
		// _ = "end of CoverTab[7396]"
//line /usr/local/go/src/crypto/sha512/sha512.go:278
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:278
	// _ = "end of CoverTab[7383]"
//line /usr/local/go/src/crypto/sha512/sha512.go:278
	_go_fuzz_dep_.CoverTab[7384]++
							return
//line /usr/local/go/src/crypto/sha512/sha512.go:279
	// _ = "end of CoverTab[7384]"
}

func (d *digest) Sum(in []byte) []byte {
//line /usr/local/go/src/crypto/sha512/sha512.go:282
	_go_fuzz_dep_.CoverTab[7397]++
							if d.function != crypto.SHA512_224 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512.go:283
		_go_fuzz_dep_.CoverTab[7399]++
//line /usr/local/go/src/crypto/sha512/sha512.go:283
		return d.function != crypto.SHA512_256
//line /usr/local/go/src/crypto/sha512/sha512.go:283
		// _ = "end of CoverTab[7399]"
//line /usr/local/go/src/crypto/sha512/sha512.go:283
	}() {
//line /usr/local/go/src/crypto/sha512/sha512.go:283
		_go_fuzz_dep_.CoverTab[7400]++
								boring.Unreachable()
//line /usr/local/go/src/crypto/sha512/sha512.go:284
		// _ = "end of CoverTab[7400]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:285
		_go_fuzz_dep_.CoverTab[7401]++
//line /usr/local/go/src/crypto/sha512/sha512.go:285
		// _ = "end of CoverTab[7401]"
//line /usr/local/go/src/crypto/sha512/sha512.go:285
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:285
	// _ = "end of CoverTab[7397]"
//line /usr/local/go/src/crypto/sha512/sha512.go:285
	_go_fuzz_dep_.CoverTab[7398]++

							d0 := new(digest)
							*d0 = *d
							hash := d0.checkSum()
							switch d0.function {
	case crypto.SHA384:
//line /usr/local/go/src/crypto/sha512/sha512.go:291
		_go_fuzz_dep_.CoverTab[7402]++
								return append(in, hash[:Size384]...)
//line /usr/local/go/src/crypto/sha512/sha512.go:292
		// _ = "end of CoverTab[7402]"
	case crypto.SHA512_224:
//line /usr/local/go/src/crypto/sha512/sha512.go:293
		_go_fuzz_dep_.CoverTab[7403]++
								return append(in, hash[:Size224]...)
//line /usr/local/go/src/crypto/sha512/sha512.go:294
		// _ = "end of CoverTab[7403]"
	case crypto.SHA512_256:
//line /usr/local/go/src/crypto/sha512/sha512.go:295
		_go_fuzz_dep_.CoverTab[7404]++
								return append(in, hash[:Size256]...)
//line /usr/local/go/src/crypto/sha512/sha512.go:296
		// _ = "end of CoverTab[7404]"
	default:
//line /usr/local/go/src/crypto/sha512/sha512.go:297
		_go_fuzz_dep_.CoverTab[7405]++
								return append(in, hash[:]...)
//line /usr/local/go/src/crypto/sha512/sha512.go:298
		// _ = "end of CoverTab[7405]"
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:299
	// _ = "end of CoverTab[7398]"
}

func (d *digest) checkSum() [Size]byte {
//line /usr/local/go/src/crypto/sha512/sha512.go:302
	_go_fuzz_dep_.CoverTab[7406]++

							len := d.len
							var tmp [128 + 16]byte	// padding + length buffer
							tmp[0] = 0x80
							var t uint64
							if len%128 < 112 {
//line /usr/local/go/src/crypto/sha512/sha512.go:308
		_go_fuzz_dep_.CoverTab[7410]++
								t = 112 - len%128
//line /usr/local/go/src/crypto/sha512/sha512.go:309
		// _ = "end of CoverTab[7410]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:310
		_go_fuzz_dep_.CoverTab[7411]++
								t = 128 + 112 - len%128
//line /usr/local/go/src/crypto/sha512/sha512.go:311
		// _ = "end of CoverTab[7411]"
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:312
	// _ = "end of CoverTab[7406]"
//line /usr/local/go/src/crypto/sha512/sha512.go:312
	_go_fuzz_dep_.CoverTab[7407]++

//line /usr/local/go/src/crypto/sha512/sha512.go:315
	len <<= 3
							padlen := tmp[:t+16]

//line /usr/local/go/src/crypto/sha512/sha512.go:320
	binary.BigEndian.PutUint64(padlen[t+8:], len)
	d.Write(padlen)

	if d.nx != 0 {
//line /usr/local/go/src/crypto/sha512/sha512.go:323
		_go_fuzz_dep_.CoverTab[7412]++
								panic("d.nx != 0")
//line /usr/local/go/src/crypto/sha512/sha512.go:324
		// _ = "end of CoverTab[7412]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:325
		_go_fuzz_dep_.CoverTab[7413]++
//line /usr/local/go/src/crypto/sha512/sha512.go:325
		// _ = "end of CoverTab[7413]"
//line /usr/local/go/src/crypto/sha512/sha512.go:325
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:325
	// _ = "end of CoverTab[7407]"
//line /usr/local/go/src/crypto/sha512/sha512.go:325
	_go_fuzz_dep_.CoverTab[7408]++

							var digest [Size]byte
							binary.BigEndian.PutUint64(digest[0:], d.h[0])
							binary.BigEndian.PutUint64(digest[8:], d.h[1])
							binary.BigEndian.PutUint64(digest[16:], d.h[2])
							binary.BigEndian.PutUint64(digest[24:], d.h[3])
							binary.BigEndian.PutUint64(digest[32:], d.h[4])
							binary.BigEndian.PutUint64(digest[40:], d.h[5])
							if d.function != crypto.SHA384 {
//line /usr/local/go/src/crypto/sha512/sha512.go:334
		_go_fuzz_dep_.CoverTab[7414]++
								binary.BigEndian.PutUint64(digest[48:], d.h[6])
								binary.BigEndian.PutUint64(digest[56:], d.h[7])
//line /usr/local/go/src/crypto/sha512/sha512.go:336
		// _ = "end of CoverTab[7414]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:337
		_go_fuzz_dep_.CoverTab[7415]++
//line /usr/local/go/src/crypto/sha512/sha512.go:337
		// _ = "end of CoverTab[7415]"
//line /usr/local/go/src/crypto/sha512/sha512.go:337
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:337
	// _ = "end of CoverTab[7408]"
//line /usr/local/go/src/crypto/sha512/sha512.go:337
	_go_fuzz_dep_.CoverTab[7409]++

							return digest
//line /usr/local/go/src/crypto/sha512/sha512.go:339
	// _ = "end of CoverTab[7409]"
}

// Sum512 returns the SHA512 checksum of the data.
func Sum512(data []byte) [Size]byte {
//line /usr/local/go/src/crypto/sha512/sha512.go:343
	_go_fuzz_dep_.CoverTab[7416]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha512/sha512.go:344
		_go_fuzz_dep_.CoverTab[7418]++
								return boring.SHA512(data)
//line /usr/local/go/src/crypto/sha512/sha512.go:345
		// _ = "end of CoverTab[7418]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:346
		_go_fuzz_dep_.CoverTab[7419]++
//line /usr/local/go/src/crypto/sha512/sha512.go:346
		// _ = "end of CoverTab[7419]"
//line /usr/local/go/src/crypto/sha512/sha512.go:346
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:346
	// _ = "end of CoverTab[7416]"
//line /usr/local/go/src/crypto/sha512/sha512.go:346
	_go_fuzz_dep_.CoverTab[7417]++
							d := digest{function: crypto.SHA512}
							d.Reset()
							d.Write(data)
							return d.checkSum()
//line /usr/local/go/src/crypto/sha512/sha512.go:350
	// _ = "end of CoverTab[7417]"
}

// Sum384 returns the SHA384 checksum of the data.
func Sum384(data []byte) [Size384]byte {
//line /usr/local/go/src/crypto/sha512/sha512.go:354
	_go_fuzz_dep_.CoverTab[7420]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/sha512/sha512.go:355
		_go_fuzz_dep_.CoverTab[7422]++
								return boring.SHA384(data)
//line /usr/local/go/src/crypto/sha512/sha512.go:356
		// _ = "end of CoverTab[7422]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512.go:357
		_go_fuzz_dep_.CoverTab[7423]++
//line /usr/local/go/src/crypto/sha512/sha512.go:357
		// _ = "end of CoverTab[7423]"
//line /usr/local/go/src/crypto/sha512/sha512.go:357
	}
//line /usr/local/go/src/crypto/sha512/sha512.go:357
	// _ = "end of CoverTab[7420]"
//line /usr/local/go/src/crypto/sha512/sha512.go:357
	_go_fuzz_dep_.CoverTab[7421]++
							d := digest{function: crypto.SHA384}
							d.Reset()
							d.Write(data)
							sum := d.checkSum()
							ap := (*[Size384]byte)(sum[:])
							return *ap
//line /usr/local/go/src/crypto/sha512/sha512.go:363
	// _ = "end of CoverTab[7421]"
}

// Sum512_224 returns the Sum512/224 checksum of the data.
func Sum512_224(data []byte) [Size224]byte {
//line /usr/local/go/src/crypto/sha512/sha512.go:367
	_go_fuzz_dep_.CoverTab[7424]++
							d := digest{function: crypto.SHA512_224}
							d.Reset()
							d.Write(data)
							sum := d.checkSum()
							ap := (*[Size224]byte)(sum[:])
							return *ap
//line /usr/local/go/src/crypto/sha512/sha512.go:373
	// _ = "end of CoverTab[7424]"
}

// Sum512_256 returns the Sum512/256 checksum of the data.
func Sum512_256(data []byte) [Size256]byte {
//line /usr/local/go/src/crypto/sha512/sha512.go:377
	_go_fuzz_dep_.CoverTab[7425]++
							d := digest{function: crypto.SHA512_256}
							d.Reset()
							d.Write(data)
							sum := d.checkSum()
							ap := (*[Size256]byte)(sum[:])
							return *ap
//line /usr/local/go/src/crypto/sha512/sha512.go:383
	// _ = "end of CoverTab[7425]"
}

//line /usr/local/go/src/crypto/sha512/sha512.go:384
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha512/sha512.go:384
var _ = _go_fuzz_dep_.CoverTab
