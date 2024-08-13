// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go -output md5block.go

//line /usr/local/go/src/crypto/md5/md5.go:5
//go:generate go run gen.go -output md5block.go

//line /usr/local/go/src/crypto/md5/md5.go:11
package md5

//line /usr/local/go/src/crypto/md5/md5.go:11
import (
//line /usr/local/go/src/crypto/md5/md5.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/md5/md5.go:11
)
//line /usr/local/go/src/crypto/md5/md5.go:11
import (
//line /usr/local/go/src/crypto/md5/md5.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/md5/md5.go:11
)

import (
	"crypto"
	"encoding/binary"
	"errors"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.MD5, New)
}

//line /usr/local/go/src/crypto/md5/md5.go:25
const Size = 16

//line /usr/local/go/src/crypto/md5/md5.go:28
const BlockSize = 64

const (
	init0	= 0x67452301
	init1	= 0xEFCDAB89
	init2	= 0x98BADCFE
	init3	= 0x10325476
)

//line /usr/local/go/src/crypto/md5/md5.go:38
type digest struct {
	s	[4]uint32
	x	[BlockSize]byte
	nx	int
	len	uint64
}

func (d *digest) Reset() {
//line /usr/local/go/src/crypto/md5/md5.go:45
	_go_fuzz_dep_.CoverTab[9552]++
						d.s[0] = init0
						d.s[1] = init1
						d.s[2] = init2
						d.s[3] = init3
						d.nx = 0
						d.len = 0
//line /usr/local/go/src/crypto/md5/md5.go:51
	// _ = "end of CoverTab[9552]"
}

const (
	magic		= "md5\x01"
	marshaledSize	= len(magic) + 4*4 + BlockSize + 8
)

func (d *digest) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/crypto/md5/md5.go:59
	_go_fuzz_dep_.CoverTab[9553]++
						b := make([]byte, 0, marshaledSize)
						b = append(b, magic...)
						b = binary.BigEndian.AppendUint32(b, d.s[0])
						b = binary.BigEndian.AppendUint32(b, d.s[1])
						b = binary.BigEndian.AppendUint32(b, d.s[2])
						b = binary.BigEndian.AppendUint32(b, d.s[3])
						b = append(b, d.x[:d.nx]...)
						b = b[:len(b)+len(d.x)-d.nx]
						b = binary.BigEndian.AppendUint64(b, d.len)
						return b, nil
//line /usr/local/go/src/crypto/md5/md5.go:69
	// _ = "end of CoverTab[9553]"
}

func (d *digest) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/crypto/md5/md5.go:72
	_go_fuzz_dep_.CoverTab[9554]++
						if len(b) < len(magic) || func() bool {
//line /usr/local/go/src/crypto/md5/md5.go:73
		_go_fuzz_dep_.CoverTab[9557]++
//line /usr/local/go/src/crypto/md5/md5.go:73
		return string(b[:len(magic)]) != magic
//line /usr/local/go/src/crypto/md5/md5.go:73
		// _ = "end of CoverTab[9557]"
//line /usr/local/go/src/crypto/md5/md5.go:73
	}() {
//line /usr/local/go/src/crypto/md5/md5.go:73
		_go_fuzz_dep_.CoverTab[9558]++
							return errors.New("crypto/md5: invalid hash state identifier")
//line /usr/local/go/src/crypto/md5/md5.go:74
		// _ = "end of CoverTab[9558]"
	} else {
//line /usr/local/go/src/crypto/md5/md5.go:75
		_go_fuzz_dep_.CoverTab[9559]++
//line /usr/local/go/src/crypto/md5/md5.go:75
		// _ = "end of CoverTab[9559]"
//line /usr/local/go/src/crypto/md5/md5.go:75
	}
//line /usr/local/go/src/crypto/md5/md5.go:75
	// _ = "end of CoverTab[9554]"
//line /usr/local/go/src/crypto/md5/md5.go:75
	_go_fuzz_dep_.CoverTab[9555]++
						if len(b) != marshaledSize {
//line /usr/local/go/src/crypto/md5/md5.go:76
		_go_fuzz_dep_.CoverTab[9560]++
							return errors.New("crypto/md5: invalid hash state size")
//line /usr/local/go/src/crypto/md5/md5.go:77
		// _ = "end of CoverTab[9560]"
	} else {
//line /usr/local/go/src/crypto/md5/md5.go:78
		_go_fuzz_dep_.CoverTab[9561]++
//line /usr/local/go/src/crypto/md5/md5.go:78
		// _ = "end of CoverTab[9561]"
//line /usr/local/go/src/crypto/md5/md5.go:78
	}
//line /usr/local/go/src/crypto/md5/md5.go:78
	// _ = "end of CoverTab[9555]"
//line /usr/local/go/src/crypto/md5/md5.go:78
	_go_fuzz_dep_.CoverTab[9556]++
						b = b[len(magic):]
						b, d.s[0] = consumeUint32(b)
						b, d.s[1] = consumeUint32(b)
						b, d.s[2] = consumeUint32(b)
						b, d.s[3] = consumeUint32(b)
						b = b[copy(d.x[:], b):]
						b, d.len = consumeUint64(b)
						d.nx = int(d.len % BlockSize)
						return nil
//line /usr/local/go/src/crypto/md5/md5.go:87
	// _ = "end of CoverTab[9556]"
}

func consumeUint64(b []byte) ([]byte, uint64) {
//line /usr/local/go/src/crypto/md5/md5.go:90
	_go_fuzz_dep_.CoverTab[9562]++
						return b[8:], binary.BigEndian.Uint64(b[0:8])
//line /usr/local/go/src/crypto/md5/md5.go:91
	// _ = "end of CoverTab[9562]"
}

func consumeUint32(b []byte) ([]byte, uint32) {
//line /usr/local/go/src/crypto/md5/md5.go:94
	_go_fuzz_dep_.CoverTab[9563]++
						return b[4:], binary.BigEndian.Uint32(b[0:4])
//line /usr/local/go/src/crypto/md5/md5.go:95
	// _ = "end of CoverTab[9563]"
}

//line /usr/local/go/src/crypto/md5/md5.go:101
func New() hash.Hash {
//line /usr/local/go/src/crypto/md5/md5.go:101
	_go_fuzz_dep_.CoverTab[9564]++
						d := new(digest)
						d.Reset()
						return d
//line /usr/local/go/src/crypto/md5/md5.go:104
	// _ = "end of CoverTab[9564]"
}

func (d *digest) Size() int {
//line /usr/local/go/src/crypto/md5/md5.go:107
	_go_fuzz_dep_.CoverTab[9565]++
//line /usr/local/go/src/crypto/md5/md5.go:107
	return Size
//line /usr/local/go/src/crypto/md5/md5.go:107
	// _ = "end of CoverTab[9565]"
//line /usr/local/go/src/crypto/md5/md5.go:107
}

func (d *digest) BlockSize() int {
//line /usr/local/go/src/crypto/md5/md5.go:109
	_go_fuzz_dep_.CoverTab[9566]++
//line /usr/local/go/src/crypto/md5/md5.go:109
	return BlockSize
//line /usr/local/go/src/crypto/md5/md5.go:109
	// _ = "end of CoverTab[9566]"
//line /usr/local/go/src/crypto/md5/md5.go:109
}

func (d *digest) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/crypto/md5/md5.go:111
	_go_fuzz_dep_.CoverTab[9567]++

//line /usr/local/go/src/crypto/md5/md5.go:115
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 {
//line /usr/local/go/src/crypto/md5/md5.go:117
		_go_fuzz_dep_.CoverTab[9571]++
							n := copy(d.x[d.nx:], p)
							d.nx += n
							if d.nx == BlockSize {
//line /usr/local/go/src/crypto/md5/md5.go:120
			_go_fuzz_dep_.CoverTab[9573]++
								if haveAsm {
//line /usr/local/go/src/crypto/md5/md5.go:121
				_go_fuzz_dep_.CoverTab[9575]++
									block(d, d.x[:])
//line /usr/local/go/src/crypto/md5/md5.go:122
				// _ = "end of CoverTab[9575]"
			} else {
//line /usr/local/go/src/crypto/md5/md5.go:123
				_go_fuzz_dep_.CoverTab[9576]++
									blockGeneric(d, d.x[:])
//line /usr/local/go/src/crypto/md5/md5.go:124
				// _ = "end of CoverTab[9576]"
			}
//line /usr/local/go/src/crypto/md5/md5.go:125
			// _ = "end of CoverTab[9573]"
//line /usr/local/go/src/crypto/md5/md5.go:125
			_go_fuzz_dep_.CoverTab[9574]++
								d.nx = 0
//line /usr/local/go/src/crypto/md5/md5.go:126
			// _ = "end of CoverTab[9574]"
		} else {
//line /usr/local/go/src/crypto/md5/md5.go:127
			_go_fuzz_dep_.CoverTab[9577]++
//line /usr/local/go/src/crypto/md5/md5.go:127
			// _ = "end of CoverTab[9577]"
//line /usr/local/go/src/crypto/md5/md5.go:127
		}
//line /usr/local/go/src/crypto/md5/md5.go:127
		// _ = "end of CoverTab[9571]"
//line /usr/local/go/src/crypto/md5/md5.go:127
		_go_fuzz_dep_.CoverTab[9572]++
							p = p[n:]
//line /usr/local/go/src/crypto/md5/md5.go:128
		// _ = "end of CoverTab[9572]"
	} else {
//line /usr/local/go/src/crypto/md5/md5.go:129
		_go_fuzz_dep_.CoverTab[9578]++
//line /usr/local/go/src/crypto/md5/md5.go:129
		// _ = "end of CoverTab[9578]"
//line /usr/local/go/src/crypto/md5/md5.go:129
	}
//line /usr/local/go/src/crypto/md5/md5.go:129
	// _ = "end of CoverTab[9567]"
//line /usr/local/go/src/crypto/md5/md5.go:129
	_go_fuzz_dep_.CoverTab[9568]++
						if len(p) >= BlockSize {
//line /usr/local/go/src/crypto/md5/md5.go:130
		_go_fuzz_dep_.CoverTab[9579]++
							n := len(p) &^ (BlockSize - 1)
							if haveAsm {
//line /usr/local/go/src/crypto/md5/md5.go:132
			_go_fuzz_dep_.CoverTab[9581]++
								block(d, p[:n])
//line /usr/local/go/src/crypto/md5/md5.go:133
			// _ = "end of CoverTab[9581]"
		} else {
//line /usr/local/go/src/crypto/md5/md5.go:134
			_go_fuzz_dep_.CoverTab[9582]++
								blockGeneric(d, p[:n])
//line /usr/local/go/src/crypto/md5/md5.go:135
			// _ = "end of CoverTab[9582]"
		}
//line /usr/local/go/src/crypto/md5/md5.go:136
		// _ = "end of CoverTab[9579]"
//line /usr/local/go/src/crypto/md5/md5.go:136
		_go_fuzz_dep_.CoverTab[9580]++
							p = p[n:]
//line /usr/local/go/src/crypto/md5/md5.go:137
		// _ = "end of CoverTab[9580]"
	} else {
//line /usr/local/go/src/crypto/md5/md5.go:138
		_go_fuzz_dep_.CoverTab[9583]++
//line /usr/local/go/src/crypto/md5/md5.go:138
		// _ = "end of CoverTab[9583]"
//line /usr/local/go/src/crypto/md5/md5.go:138
	}
//line /usr/local/go/src/crypto/md5/md5.go:138
	// _ = "end of CoverTab[9568]"
//line /usr/local/go/src/crypto/md5/md5.go:138
	_go_fuzz_dep_.CoverTab[9569]++
						if len(p) > 0 {
//line /usr/local/go/src/crypto/md5/md5.go:139
		_go_fuzz_dep_.CoverTab[9584]++
							d.nx = copy(d.x[:], p)
//line /usr/local/go/src/crypto/md5/md5.go:140
		// _ = "end of CoverTab[9584]"
	} else {
//line /usr/local/go/src/crypto/md5/md5.go:141
		_go_fuzz_dep_.CoverTab[9585]++
//line /usr/local/go/src/crypto/md5/md5.go:141
		// _ = "end of CoverTab[9585]"
//line /usr/local/go/src/crypto/md5/md5.go:141
	}
//line /usr/local/go/src/crypto/md5/md5.go:141
	// _ = "end of CoverTab[9569]"
//line /usr/local/go/src/crypto/md5/md5.go:141
	_go_fuzz_dep_.CoverTab[9570]++
						return
//line /usr/local/go/src/crypto/md5/md5.go:142
	// _ = "end of CoverTab[9570]"
}

func (d *digest) Sum(in []byte) []byte {
//line /usr/local/go/src/crypto/md5/md5.go:145
	_go_fuzz_dep_.CoverTab[9586]++

						d0 := *d
						hash := d0.checkSum()
						return append(in, hash[:]...)
//line /usr/local/go/src/crypto/md5/md5.go:149
	// _ = "end of CoverTab[9586]"
}

func (d *digest) checkSum() [Size]byte {
//line /usr/local/go/src/crypto/md5/md5.go:152
	_go_fuzz_dep_.CoverTab[9587]++

//line /usr/local/go/src/crypto/md5/md5.go:158
	tmp := [1 + 63 + 8]byte{0x80}
						pad := (55 - d.len) % 64
						binary.LittleEndian.PutUint64(tmp[1+pad:], d.len<<3)
						d.Write(tmp[:1+pad+8])

//line /usr/local/go/src/crypto/md5/md5.go:165
	if d.nx != 0 {
//line /usr/local/go/src/crypto/md5/md5.go:165
		_go_fuzz_dep_.CoverTab[9589]++
							panic("d.nx != 0")
//line /usr/local/go/src/crypto/md5/md5.go:166
		// _ = "end of CoverTab[9589]"
	} else {
//line /usr/local/go/src/crypto/md5/md5.go:167
		_go_fuzz_dep_.CoverTab[9590]++
//line /usr/local/go/src/crypto/md5/md5.go:167
		// _ = "end of CoverTab[9590]"
//line /usr/local/go/src/crypto/md5/md5.go:167
	}
//line /usr/local/go/src/crypto/md5/md5.go:167
	// _ = "end of CoverTab[9587]"
//line /usr/local/go/src/crypto/md5/md5.go:167
	_go_fuzz_dep_.CoverTab[9588]++

						var digest [Size]byte
						binary.LittleEndian.PutUint32(digest[0:], d.s[0])
						binary.LittleEndian.PutUint32(digest[4:], d.s[1])
						binary.LittleEndian.PutUint32(digest[8:], d.s[2])
						binary.LittleEndian.PutUint32(digest[12:], d.s[3])
						return digest
//line /usr/local/go/src/crypto/md5/md5.go:174
	// _ = "end of CoverTab[9588]"
}

//line /usr/local/go/src/crypto/md5/md5.go:178
func Sum(data []byte) [Size]byte {
//line /usr/local/go/src/crypto/md5/md5.go:178
	_go_fuzz_dep_.CoverTab[9591]++
						var d digest
						d.Reset()
						d.Write(data)
						return d.checkSum()
//line /usr/local/go/src/crypto/md5/md5.go:182
	// _ = "end of CoverTab[9591]"
}

//line /usr/local/go/src/crypto/md5/md5.go:183
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/md5/md5.go:183
var _ = _go_fuzz_dep_.CoverTab
