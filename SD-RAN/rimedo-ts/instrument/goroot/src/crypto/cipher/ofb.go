// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// OFB (Output Feedback) Mode.

//line /usr/local/go/src/crypto/cipher/ofb.go:7
package cipher

//line /usr/local/go/src/crypto/cipher/ofb.go:7
import (
//line /usr/local/go/src/crypto/cipher/ofb.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/ofb.go:7
)
//line /usr/local/go/src/crypto/cipher/ofb.go:7
import (
//line /usr/local/go/src/crypto/cipher/ofb.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/ofb.go:7
)

import (
	"crypto/internal/alias"
	"crypto/subtle"
)

type ofb struct {
	b	Block
	cipher	[]byte
	out	[]byte
	outUsed	int
}

// NewOFB returns a Stream that encrypts or decrypts using the block cipher b
//line /usr/local/go/src/crypto/cipher/ofb.go:21
// in output feedback mode. The initialization vector iv's length must be equal
//line /usr/local/go/src/crypto/cipher/ofb.go:21
// to b's block size.
//line /usr/local/go/src/crypto/cipher/ofb.go:24
func NewOFB(b Block, iv []byte) Stream {
//line /usr/local/go/src/crypto/cipher/ofb.go:24
	_go_fuzz_dep_.CoverTab[1682]++
							blockSize := b.BlockSize()
							if len(iv) != blockSize {
//line /usr/local/go/src/crypto/cipher/ofb.go:26
		_go_fuzz_dep_.CoverTab[1685]++
								panic("cipher.NewOFB: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/ofb.go:27
		// _ = "end of CoverTab[1685]"
	} else {
//line /usr/local/go/src/crypto/cipher/ofb.go:28
		_go_fuzz_dep_.CoverTab[1686]++
//line /usr/local/go/src/crypto/cipher/ofb.go:28
		// _ = "end of CoverTab[1686]"
//line /usr/local/go/src/crypto/cipher/ofb.go:28
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:28
	// _ = "end of CoverTab[1682]"
//line /usr/local/go/src/crypto/cipher/ofb.go:28
	_go_fuzz_dep_.CoverTab[1683]++
							bufSize := streamBufferSize
							if bufSize < blockSize {
//line /usr/local/go/src/crypto/cipher/ofb.go:30
		_go_fuzz_dep_.CoverTab[1687]++
								bufSize = blockSize
//line /usr/local/go/src/crypto/cipher/ofb.go:31
		// _ = "end of CoverTab[1687]"
	} else {
//line /usr/local/go/src/crypto/cipher/ofb.go:32
		_go_fuzz_dep_.CoverTab[1688]++
//line /usr/local/go/src/crypto/cipher/ofb.go:32
		// _ = "end of CoverTab[1688]"
//line /usr/local/go/src/crypto/cipher/ofb.go:32
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:32
	// _ = "end of CoverTab[1683]"
//line /usr/local/go/src/crypto/cipher/ofb.go:32
	_go_fuzz_dep_.CoverTab[1684]++
							x := &ofb{
		b:		b,
		cipher:		make([]byte, blockSize),
		out:		make([]byte, 0, bufSize),
		outUsed:	0,
	}

							copy(x.cipher, iv)
							return x
//line /usr/local/go/src/crypto/cipher/ofb.go:41
	// _ = "end of CoverTab[1684]"
}

func (x *ofb) refill() {
//line /usr/local/go/src/crypto/cipher/ofb.go:44
	_go_fuzz_dep_.CoverTab[1689]++
							bs := x.b.BlockSize()
							remain := len(x.out) - x.outUsed
							if remain > x.outUsed {
//line /usr/local/go/src/crypto/cipher/ofb.go:47
		_go_fuzz_dep_.CoverTab[1692]++
								return
//line /usr/local/go/src/crypto/cipher/ofb.go:48
		// _ = "end of CoverTab[1692]"
	} else {
//line /usr/local/go/src/crypto/cipher/ofb.go:49
		_go_fuzz_dep_.CoverTab[1693]++
//line /usr/local/go/src/crypto/cipher/ofb.go:49
		// _ = "end of CoverTab[1693]"
//line /usr/local/go/src/crypto/cipher/ofb.go:49
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:49
	// _ = "end of CoverTab[1689]"
//line /usr/local/go/src/crypto/cipher/ofb.go:49
	_go_fuzz_dep_.CoverTab[1690]++
							copy(x.out, x.out[x.outUsed:])
							x.out = x.out[:cap(x.out)]
							for remain < len(x.out)-bs {
//line /usr/local/go/src/crypto/cipher/ofb.go:52
		_go_fuzz_dep_.CoverTab[1694]++
								x.b.Encrypt(x.cipher, x.cipher)
								copy(x.out[remain:], x.cipher)
								remain += bs
//line /usr/local/go/src/crypto/cipher/ofb.go:55
		// _ = "end of CoverTab[1694]"
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:56
	// _ = "end of CoverTab[1690]"
//line /usr/local/go/src/crypto/cipher/ofb.go:56
	_go_fuzz_dep_.CoverTab[1691]++
							x.out = x.out[:remain]
							x.outUsed = 0
//line /usr/local/go/src/crypto/cipher/ofb.go:58
	// _ = "end of CoverTab[1691]"
}

func (x *ofb) XORKeyStream(dst, src []byte) {
//line /usr/local/go/src/crypto/cipher/ofb.go:61
	_go_fuzz_dep_.CoverTab[1695]++
							if len(dst) < len(src) {
//line /usr/local/go/src/crypto/cipher/ofb.go:62
		_go_fuzz_dep_.CoverTab[1698]++
								panic("crypto/cipher: output smaller than input")
//line /usr/local/go/src/crypto/cipher/ofb.go:63
		// _ = "end of CoverTab[1698]"
	} else {
//line /usr/local/go/src/crypto/cipher/ofb.go:64
		_go_fuzz_dep_.CoverTab[1699]++
//line /usr/local/go/src/crypto/cipher/ofb.go:64
		// _ = "end of CoverTab[1699]"
//line /usr/local/go/src/crypto/cipher/ofb.go:64
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:64
	// _ = "end of CoverTab[1695]"
//line /usr/local/go/src/crypto/cipher/ofb.go:64
	_go_fuzz_dep_.CoverTab[1696]++
							if alias.InexactOverlap(dst[:len(src)], src) {
//line /usr/local/go/src/crypto/cipher/ofb.go:65
		_go_fuzz_dep_.CoverTab[1700]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/ofb.go:66
		// _ = "end of CoverTab[1700]"
	} else {
//line /usr/local/go/src/crypto/cipher/ofb.go:67
		_go_fuzz_dep_.CoverTab[1701]++
//line /usr/local/go/src/crypto/cipher/ofb.go:67
		// _ = "end of CoverTab[1701]"
//line /usr/local/go/src/crypto/cipher/ofb.go:67
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:67
	// _ = "end of CoverTab[1696]"
//line /usr/local/go/src/crypto/cipher/ofb.go:67
	_go_fuzz_dep_.CoverTab[1697]++
							for len(src) > 0 {
//line /usr/local/go/src/crypto/cipher/ofb.go:68
		_go_fuzz_dep_.CoverTab[1702]++
								if x.outUsed >= len(x.out)-x.b.BlockSize() {
//line /usr/local/go/src/crypto/cipher/ofb.go:69
			_go_fuzz_dep_.CoverTab[1704]++
									x.refill()
//line /usr/local/go/src/crypto/cipher/ofb.go:70
			// _ = "end of CoverTab[1704]"
		} else {
//line /usr/local/go/src/crypto/cipher/ofb.go:71
			_go_fuzz_dep_.CoverTab[1705]++
//line /usr/local/go/src/crypto/cipher/ofb.go:71
			// _ = "end of CoverTab[1705]"
//line /usr/local/go/src/crypto/cipher/ofb.go:71
		}
//line /usr/local/go/src/crypto/cipher/ofb.go:71
		// _ = "end of CoverTab[1702]"
//line /usr/local/go/src/crypto/cipher/ofb.go:71
		_go_fuzz_dep_.CoverTab[1703]++
								n := subtle.XORBytes(dst, src, x.out[x.outUsed:])
								dst = dst[n:]
								src = src[n:]
								x.outUsed += n
//line /usr/local/go/src/crypto/cipher/ofb.go:75
		// _ = "end of CoverTab[1703]"
	}
//line /usr/local/go/src/crypto/cipher/ofb.go:76
	// _ = "end of CoverTab[1697]"
}

//line /usr/local/go/src/crypto/cipher/ofb.go:77
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/ofb.go:77
var _ = _go_fuzz_dep_.CoverTab
