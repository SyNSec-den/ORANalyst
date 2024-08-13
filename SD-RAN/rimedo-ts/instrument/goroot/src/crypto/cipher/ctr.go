// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Counter (CTR) mode.

// CTR converts a block cipher into a stream cipher by
// repeatedly encrypting an incrementing counter and
// xoring the resulting stream of data with the input.

// See NIST SP 800-38A, pp 13-15

//line /usr/local/go/src/crypto/cipher/ctr.go:13
package cipher

//line /usr/local/go/src/crypto/cipher/ctr.go:13
import (
//line /usr/local/go/src/crypto/cipher/ctr.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/ctr.go:13
)
//line /usr/local/go/src/crypto/cipher/ctr.go:13
import (
//line /usr/local/go/src/crypto/cipher/ctr.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/ctr.go:13
)

import (
	"bytes"
	"crypto/internal/alias"
	"crypto/subtle"
)

type ctr struct {
	b	Block
	ctr	[]byte
	out	[]byte
	outUsed	int
}

const streamBufferSize = 512

// ctrAble is an interface implemented by ciphers that have a specific optimized
//line /usr/local/go/src/crypto/cipher/ctr.go:30
// implementation of CTR, like crypto/aes. NewCTR will check for this interface
//line /usr/local/go/src/crypto/cipher/ctr.go:30
// and return the specific Stream if found.
//line /usr/local/go/src/crypto/cipher/ctr.go:33
type ctrAble interface {
	NewCTR(iv []byte) Stream
}

// NewCTR returns a Stream which encrypts/decrypts using the given Block in
//line /usr/local/go/src/crypto/cipher/ctr.go:37
// counter mode. The length of iv must be the same as the Block's block size.
//line /usr/local/go/src/crypto/cipher/ctr.go:39
func NewCTR(block Block, iv []byte) Stream {
//line /usr/local/go/src/crypto/cipher/ctr.go:39
	_go_fuzz_dep_.CoverTab[1561]++
							if ctr, ok := block.(ctrAble); ok {
//line /usr/local/go/src/crypto/cipher/ctr.go:40
		_go_fuzz_dep_.CoverTab[1565]++
								return ctr.NewCTR(iv)
//line /usr/local/go/src/crypto/cipher/ctr.go:41
		// _ = "end of CoverTab[1565]"
	} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:42
		_go_fuzz_dep_.CoverTab[1566]++
//line /usr/local/go/src/crypto/cipher/ctr.go:42
		// _ = "end of CoverTab[1566]"
//line /usr/local/go/src/crypto/cipher/ctr.go:42
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:42
	// _ = "end of CoverTab[1561]"
//line /usr/local/go/src/crypto/cipher/ctr.go:42
	_go_fuzz_dep_.CoverTab[1562]++
							if len(iv) != block.BlockSize() {
//line /usr/local/go/src/crypto/cipher/ctr.go:43
		_go_fuzz_dep_.CoverTab[1567]++
								panic("cipher.NewCTR: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/ctr.go:44
		// _ = "end of CoverTab[1567]"
	} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:45
		_go_fuzz_dep_.CoverTab[1568]++
//line /usr/local/go/src/crypto/cipher/ctr.go:45
		// _ = "end of CoverTab[1568]"
//line /usr/local/go/src/crypto/cipher/ctr.go:45
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:45
	// _ = "end of CoverTab[1562]"
//line /usr/local/go/src/crypto/cipher/ctr.go:45
	_go_fuzz_dep_.CoverTab[1563]++
							bufSize := streamBufferSize
							if bufSize < block.BlockSize() {
//line /usr/local/go/src/crypto/cipher/ctr.go:47
		_go_fuzz_dep_.CoverTab[1569]++
								bufSize = block.BlockSize()
//line /usr/local/go/src/crypto/cipher/ctr.go:48
		// _ = "end of CoverTab[1569]"
	} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:49
		_go_fuzz_dep_.CoverTab[1570]++
//line /usr/local/go/src/crypto/cipher/ctr.go:49
		// _ = "end of CoverTab[1570]"
//line /usr/local/go/src/crypto/cipher/ctr.go:49
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:49
	// _ = "end of CoverTab[1563]"
//line /usr/local/go/src/crypto/cipher/ctr.go:49
	_go_fuzz_dep_.CoverTab[1564]++
							return &ctr{
		b:		block,
		ctr:		bytes.Clone(iv),
		out:		make([]byte, 0, bufSize),
		outUsed:	0,
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:55
	// _ = "end of CoverTab[1564]"
}

func (x *ctr) refill() {
//line /usr/local/go/src/crypto/cipher/ctr.go:58
	_go_fuzz_dep_.CoverTab[1571]++
							remain := len(x.out) - x.outUsed
							copy(x.out, x.out[x.outUsed:])
							x.out = x.out[:cap(x.out)]
							bs := x.b.BlockSize()
							for remain <= len(x.out)-bs {
//line /usr/local/go/src/crypto/cipher/ctr.go:63
		_go_fuzz_dep_.CoverTab[1573]++
								x.b.Encrypt(x.out[remain:], x.ctr)
								remain += bs

//line /usr/local/go/src/crypto/cipher/ctr.go:68
		for i := len(x.ctr) - 1; i >= 0; i-- {
//line /usr/local/go/src/crypto/cipher/ctr.go:68
			_go_fuzz_dep_.CoverTab[1574]++
									x.ctr[i]++
									if x.ctr[i] != 0 {
//line /usr/local/go/src/crypto/cipher/ctr.go:70
				_go_fuzz_dep_.CoverTab[1575]++
										break
//line /usr/local/go/src/crypto/cipher/ctr.go:71
				// _ = "end of CoverTab[1575]"
			} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:72
				_go_fuzz_dep_.CoverTab[1576]++
//line /usr/local/go/src/crypto/cipher/ctr.go:72
				// _ = "end of CoverTab[1576]"
//line /usr/local/go/src/crypto/cipher/ctr.go:72
			}
//line /usr/local/go/src/crypto/cipher/ctr.go:72
			// _ = "end of CoverTab[1574]"
		}
//line /usr/local/go/src/crypto/cipher/ctr.go:73
		// _ = "end of CoverTab[1573]"
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:74
	// _ = "end of CoverTab[1571]"
//line /usr/local/go/src/crypto/cipher/ctr.go:74
	_go_fuzz_dep_.CoverTab[1572]++
							x.out = x.out[:remain]
							x.outUsed = 0
//line /usr/local/go/src/crypto/cipher/ctr.go:76
	// _ = "end of CoverTab[1572]"
}

func (x *ctr) XORKeyStream(dst, src []byte) {
//line /usr/local/go/src/crypto/cipher/ctr.go:79
	_go_fuzz_dep_.CoverTab[1577]++
							if len(dst) < len(src) {
//line /usr/local/go/src/crypto/cipher/ctr.go:80
		_go_fuzz_dep_.CoverTab[1580]++
								panic("crypto/cipher: output smaller than input")
//line /usr/local/go/src/crypto/cipher/ctr.go:81
		// _ = "end of CoverTab[1580]"
	} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:82
		_go_fuzz_dep_.CoverTab[1581]++
//line /usr/local/go/src/crypto/cipher/ctr.go:82
		// _ = "end of CoverTab[1581]"
//line /usr/local/go/src/crypto/cipher/ctr.go:82
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:82
	// _ = "end of CoverTab[1577]"
//line /usr/local/go/src/crypto/cipher/ctr.go:82
	_go_fuzz_dep_.CoverTab[1578]++
							if alias.InexactOverlap(dst[:len(src)], src) {
//line /usr/local/go/src/crypto/cipher/ctr.go:83
		_go_fuzz_dep_.CoverTab[1582]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/ctr.go:84
		// _ = "end of CoverTab[1582]"
	} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:85
		_go_fuzz_dep_.CoverTab[1583]++
//line /usr/local/go/src/crypto/cipher/ctr.go:85
		// _ = "end of CoverTab[1583]"
//line /usr/local/go/src/crypto/cipher/ctr.go:85
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:85
	// _ = "end of CoverTab[1578]"
//line /usr/local/go/src/crypto/cipher/ctr.go:85
	_go_fuzz_dep_.CoverTab[1579]++
							for len(src) > 0 {
//line /usr/local/go/src/crypto/cipher/ctr.go:86
		_go_fuzz_dep_.CoverTab[1584]++
								if x.outUsed >= len(x.out)-x.b.BlockSize() {
//line /usr/local/go/src/crypto/cipher/ctr.go:87
			_go_fuzz_dep_.CoverTab[1586]++
									x.refill()
//line /usr/local/go/src/crypto/cipher/ctr.go:88
			// _ = "end of CoverTab[1586]"
		} else {
//line /usr/local/go/src/crypto/cipher/ctr.go:89
			_go_fuzz_dep_.CoverTab[1587]++
//line /usr/local/go/src/crypto/cipher/ctr.go:89
			// _ = "end of CoverTab[1587]"
//line /usr/local/go/src/crypto/cipher/ctr.go:89
		}
//line /usr/local/go/src/crypto/cipher/ctr.go:89
		// _ = "end of CoverTab[1584]"
//line /usr/local/go/src/crypto/cipher/ctr.go:89
		_go_fuzz_dep_.CoverTab[1585]++
								n := subtle.XORBytes(dst, src, x.out[x.outUsed:])
								dst = dst[n:]
								src = src[n:]
								x.outUsed += n
//line /usr/local/go/src/crypto/cipher/ctr.go:93
		// _ = "end of CoverTab[1585]"
	}
//line /usr/local/go/src/crypto/cipher/ctr.go:94
	// _ = "end of CoverTab[1579]"
}

//line /usr/local/go/src/crypto/cipher/ctr.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/ctr.go:95
var _ = _go_fuzz_dep_.CoverTab
