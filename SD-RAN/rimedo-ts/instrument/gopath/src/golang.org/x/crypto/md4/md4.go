// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:5
// Package md4 implements the MD4 hash algorithm as defined in RFC 1320.
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:5
// Deprecated: MD4 is cryptographically broken and should should only be used
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:5
// where compatibility with legacy systems, not security, is the goal. Instead,
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:5
// use a secure hash like SHA-256 (from crypto/sha256).
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
package md4

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:10
)

import (
	"crypto"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.MD4, New)
}

// The size of an MD4 checksum in bytes.
const Size = 16

// The blocksize of MD4 in bytes.
const BlockSize = 64

const (
	_Chunk	= 64
	_Init0	= 0x67452301
	_Init1	= 0xEFCDAB89
	_Init2	= 0x98BADCFE
	_Init3	= 0x10325476
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	s	[4]uint32
	x	[_Chunk]byte
	nx	int
	len	uint64
}

func (d *digest) Reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:43
	_go_fuzz_dep_.CoverTab[85758]++
													d.s[0] = _Init0
													d.s[1] = _Init1
													d.s[2] = _Init2
													d.s[3] = _Init3
													d.nx = 0
													d.len = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:49
	// _ = "end of CoverTab[85758]"
}

// New returns a new hash.Hash computing the MD4 checksum.
func New() hash.Hash {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:53
	_go_fuzz_dep_.CoverTab[85759]++
													d := new(digest)
													d.Reset()
													return d
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:56
	// _ = "end of CoverTab[85759]"
}

func (d *digest) Size() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:59
	_go_fuzz_dep_.CoverTab[85760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:59
	return Size
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:59
	// _ = "end of CoverTab[85760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:59
}

func (d *digest) BlockSize() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:61
	_go_fuzz_dep_.CoverTab[85761]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:61
	return BlockSize
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:61
	// _ = "end of CoverTab[85761]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:61
}

func (d *digest) Write(p []byte) (nn int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:63
	_go_fuzz_dep_.CoverTab[85762]++
													nn = len(p)
													d.len += uint64(nn)
													if d.nx > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:66
		_go_fuzz_dep_.CoverTab[85765]++
														n := len(p)
														if n > _Chunk-d.nx {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:68
			_go_fuzz_dep_.CoverTab[85769]++
															n = _Chunk - d.nx
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:69
			// _ = "end of CoverTab[85769]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:70
			_go_fuzz_dep_.CoverTab[85770]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:70
			// _ = "end of CoverTab[85770]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:70
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:70
		// _ = "end of CoverTab[85765]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:70
		_go_fuzz_dep_.CoverTab[85766]++
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:71
			_go_fuzz_dep_.CoverTab[85771]++
															d.x[d.nx+i] = p[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:72
			// _ = "end of CoverTab[85771]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:73
		// _ = "end of CoverTab[85766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:73
		_go_fuzz_dep_.CoverTab[85767]++
														d.nx += n
														if d.nx == _Chunk {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:75
			_go_fuzz_dep_.CoverTab[85772]++
															_Block(d, d.x[0:])
															d.nx = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:77
			// _ = "end of CoverTab[85772]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:78
			_go_fuzz_dep_.CoverTab[85773]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:78
			// _ = "end of CoverTab[85773]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:78
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:78
		// _ = "end of CoverTab[85767]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:78
		_go_fuzz_dep_.CoverTab[85768]++
														p = p[n:]
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:79
		// _ = "end of CoverTab[85768]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:80
		_go_fuzz_dep_.CoverTab[85774]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:80
		// _ = "end of CoverTab[85774]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:80
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:80
	// _ = "end of CoverTab[85762]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:80
	_go_fuzz_dep_.CoverTab[85763]++
													n := _Block(d, p)
													p = p[n:]
													if len(p) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:83
		_go_fuzz_dep_.CoverTab[85775]++
														d.nx = copy(d.x[:], p)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:84
		// _ = "end of CoverTab[85775]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:85
		_go_fuzz_dep_.CoverTab[85776]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:85
		// _ = "end of CoverTab[85776]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:85
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:85
	// _ = "end of CoverTab[85763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:85
	_go_fuzz_dep_.CoverTab[85764]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:86
	// _ = "end of CoverTab[85764]"
}

func (d0 *digest) Sum(in []byte) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:89
	_go_fuzz_dep_.CoverTab[85777]++

													d := new(digest)
													*d = *d0

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:95
	len := d.len
	var tmp [64]byte
	tmp[0] = 0x80
	if len%64 < 56 {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:98
		_go_fuzz_dep_.CoverTab[85782]++
														d.Write(tmp[0 : 56-len%64])
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:99
		// _ = "end of CoverTab[85782]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:100
		_go_fuzz_dep_.CoverTab[85783]++
															d.Write(tmp[0 : 64+56-len%64])
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:101
		// _ = "end of CoverTab[85783]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:102
	// _ = "end of CoverTab[85777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:102
	_go_fuzz_dep_.CoverTab[85778]++

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:105
	len <<= 3
	for i := uint(0); i < 8; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:106
		_go_fuzz_dep_.CoverTab[85784]++
															tmp[i] = byte(len >> (8 * i))
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:107
		// _ = "end of CoverTab[85784]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:108
	// _ = "end of CoverTab[85778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:108
	_go_fuzz_dep_.CoverTab[85779]++
														d.Write(tmp[0:8])

														if d.nx != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:111
		_go_fuzz_dep_.CoverTab[85785]++
															panic("d.nx != 0")
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:112
		// _ = "end of CoverTab[85785]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:113
		_go_fuzz_dep_.CoverTab[85786]++
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:113
		// _ = "end of CoverTab[85786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:113
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:113
	// _ = "end of CoverTab[85779]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:113
	_go_fuzz_dep_.CoverTab[85780]++

														for _, s := range d.s {
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:115
		_go_fuzz_dep_.CoverTab[85787]++
															in = append(in, byte(s>>0))
															in = append(in, byte(s>>8))
															in = append(in, byte(s>>16))
															in = append(in, byte(s>>24))
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:119
		// _ = "end of CoverTab[85787]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:120
	// _ = "end of CoverTab[85780]"
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:120
	_go_fuzz_dep_.CoverTab[85781]++
														return in
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:121
	// _ = "end of CoverTab[85781]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/crypto@v0.0.0-20220128200615-198e4374d7ed/md4/md4.go:122
var _ = _go_fuzz_dep_.CoverTab
