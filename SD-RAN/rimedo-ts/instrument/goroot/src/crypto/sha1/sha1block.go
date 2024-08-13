// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/sha1/sha1block.go:5
package sha1

//line /usr/local/go/src/crypto/sha1/sha1block.go:5
import (
//line /usr/local/go/src/crypto/sha1/sha1block.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha1/sha1block.go:5
)
//line /usr/local/go/src/crypto/sha1/sha1block.go:5
import (
//line /usr/local/go/src/crypto/sha1/sha1block.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha1/sha1block.go:5
)

import (
	"math/bits"
)

const (
	_K0	= 0x5A827999
	_K1	= 0x6ED9EBA1
	_K2	= 0x8F1BBCDC
	_K3	= 0xCA62C1D6
)

// blockGeneric is a portable, pure Go version of the SHA-1 block step.
//line /usr/local/go/src/crypto/sha1/sha1block.go:18
// It's used by sha1block_generic.go and tests.
//line /usr/local/go/src/crypto/sha1/sha1block.go:20
func blockGeneric(dig *digest, p []byte) {
//line /usr/local/go/src/crypto/sha1/sha1block.go:20
	_go_fuzz_dep_.CoverTab[10143]++
							var w [16]uint32

							h0, h1, h2, h3, h4 := dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4]
							for len(p) >= chunk {
//line /usr/local/go/src/crypto/sha1/sha1block.go:24
		_go_fuzz_dep_.CoverTab[10145]++

//line /usr/local/go/src/crypto/sha1/sha1block.go:27
		for i := 0; i < 16; i++ {
//line /usr/local/go/src/crypto/sha1/sha1block.go:27
			_go_fuzz_dep_.CoverTab[10152]++
									j := i * 4
									w[i] = uint32(p[j])<<24 | uint32(p[j+1])<<16 | uint32(p[j+2])<<8 | uint32(p[j+3])
//line /usr/local/go/src/crypto/sha1/sha1block.go:29
			// _ = "end of CoverTab[10152]"
		}
//line /usr/local/go/src/crypto/sha1/sha1block.go:30
		// _ = "end of CoverTab[10145]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:30
		_go_fuzz_dep_.CoverTab[10146]++

								a, b, c, d, e := h0, h1, h2, h3, h4

//line /usr/local/go/src/crypto/sha1/sha1block.go:37
		i := 0
		for ; i < 16; i++ {
//line /usr/local/go/src/crypto/sha1/sha1block.go:38
			_go_fuzz_dep_.CoverTab[10153]++
									f := b&c | (^b)&d
									t := bits.RotateLeft32(a, 5) + f + e + w[i&0xf] + _K0
									a, b, c, d, e = t, a, bits.RotateLeft32(b, 30), c, d
//line /usr/local/go/src/crypto/sha1/sha1block.go:41
			// _ = "end of CoverTab[10153]"
		}
//line /usr/local/go/src/crypto/sha1/sha1block.go:42
		// _ = "end of CoverTab[10146]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:42
		_go_fuzz_dep_.CoverTab[10147]++
								for ; i < 20; i++ {
//line /usr/local/go/src/crypto/sha1/sha1block.go:43
			_go_fuzz_dep_.CoverTab[10154]++
									tmp := w[(i-3)&0xf] ^ w[(i-8)&0xf] ^ w[(i-14)&0xf] ^ w[(i)&0xf]
									w[i&0xf] = tmp<<1 | tmp>>(32-1)

									f := b&c | (^b)&d
									t := bits.RotateLeft32(a, 5) + f + e + w[i&0xf] + _K0
									a, b, c, d, e = t, a, bits.RotateLeft32(b, 30), c, d
//line /usr/local/go/src/crypto/sha1/sha1block.go:49
			// _ = "end of CoverTab[10154]"
		}
//line /usr/local/go/src/crypto/sha1/sha1block.go:50
		// _ = "end of CoverTab[10147]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:50
		_go_fuzz_dep_.CoverTab[10148]++
								for ; i < 40; i++ {
//line /usr/local/go/src/crypto/sha1/sha1block.go:51
			_go_fuzz_dep_.CoverTab[10155]++
									tmp := w[(i-3)&0xf] ^ w[(i-8)&0xf] ^ w[(i-14)&0xf] ^ w[(i)&0xf]
									w[i&0xf] = tmp<<1 | tmp>>(32-1)
									f := b ^ c ^ d
									t := bits.RotateLeft32(a, 5) + f + e + w[i&0xf] + _K1
									a, b, c, d, e = t, a, bits.RotateLeft32(b, 30), c, d
//line /usr/local/go/src/crypto/sha1/sha1block.go:56
			// _ = "end of CoverTab[10155]"
		}
//line /usr/local/go/src/crypto/sha1/sha1block.go:57
		// _ = "end of CoverTab[10148]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:57
		_go_fuzz_dep_.CoverTab[10149]++
								for ; i < 60; i++ {
//line /usr/local/go/src/crypto/sha1/sha1block.go:58
			_go_fuzz_dep_.CoverTab[10156]++
									tmp := w[(i-3)&0xf] ^ w[(i-8)&0xf] ^ w[(i-14)&0xf] ^ w[(i)&0xf]
									w[i&0xf] = tmp<<1 | tmp>>(32-1)
									f := ((b | c) & d) | (b & c)
									t := bits.RotateLeft32(a, 5) + f + e + w[i&0xf] + _K2
									a, b, c, d, e = t, a, bits.RotateLeft32(b, 30), c, d
//line /usr/local/go/src/crypto/sha1/sha1block.go:63
			// _ = "end of CoverTab[10156]"
		}
//line /usr/local/go/src/crypto/sha1/sha1block.go:64
		// _ = "end of CoverTab[10149]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:64
		_go_fuzz_dep_.CoverTab[10150]++
								for ; i < 80; i++ {
//line /usr/local/go/src/crypto/sha1/sha1block.go:65
			_go_fuzz_dep_.CoverTab[10157]++
									tmp := w[(i-3)&0xf] ^ w[(i-8)&0xf] ^ w[(i-14)&0xf] ^ w[(i)&0xf]
									w[i&0xf] = tmp<<1 | tmp>>(32-1)
									f := b ^ c ^ d
									t := bits.RotateLeft32(a, 5) + f + e + w[i&0xf] + _K3
									a, b, c, d, e = t, a, bits.RotateLeft32(b, 30), c, d
//line /usr/local/go/src/crypto/sha1/sha1block.go:70
			// _ = "end of CoverTab[10157]"
		}
//line /usr/local/go/src/crypto/sha1/sha1block.go:71
		// _ = "end of CoverTab[10150]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:71
		_go_fuzz_dep_.CoverTab[10151]++

								h0 += a
								h1 += b
								h2 += c
								h3 += d
								h4 += e

								p = p[chunk:]
//line /usr/local/go/src/crypto/sha1/sha1block.go:79
		// _ = "end of CoverTab[10151]"
	}
//line /usr/local/go/src/crypto/sha1/sha1block.go:80
	// _ = "end of CoverTab[10143]"
//line /usr/local/go/src/crypto/sha1/sha1block.go:80
	_go_fuzz_dep_.CoverTab[10144]++

							dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4] = h0, h1, h2, h3, h4
//line /usr/local/go/src/crypto/sha1/sha1block.go:82
	// _ = "end of CoverTab[10144]"
}

//line /usr/local/go/src/crypto/sha1/sha1block.go:83
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha1/sha1block.go:83
var _ = _go_fuzz_dep_.CoverTab
