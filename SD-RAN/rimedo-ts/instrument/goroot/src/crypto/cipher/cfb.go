// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CFB (Cipher Feedback) Mode.

//line /usr/local/go/src/crypto/cipher/cfb.go:7
package cipher

//line /usr/local/go/src/crypto/cipher/cfb.go:7
import (
//line /usr/local/go/src/crypto/cipher/cfb.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/cfb.go:7
)
//line /usr/local/go/src/crypto/cipher/cfb.go:7
import (
//line /usr/local/go/src/crypto/cipher/cfb.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/cfb.go:7
)

import (
	"crypto/internal/alias"
	"crypto/subtle"
)

type cfb struct {
	b	Block
	next	[]byte
	out	[]byte
	outUsed	int

	decrypt	bool
}

func (x *cfb) XORKeyStream(dst, src []byte) {
//line /usr/local/go/src/crypto/cipher/cfb.go:23
	_go_fuzz_dep_.CoverTab[1538]++
							if len(dst) < len(src) {
//line /usr/local/go/src/crypto/cipher/cfb.go:24
		_go_fuzz_dep_.CoverTab[1541]++
								panic("crypto/cipher: output smaller than input")
//line /usr/local/go/src/crypto/cipher/cfb.go:25
		// _ = "end of CoverTab[1541]"
	} else {
//line /usr/local/go/src/crypto/cipher/cfb.go:26
		_go_fuzz_dep_.CoverTab[1542]++
//line /usr/local/go/src/crypto/cipher/cfb.go:26
		// _ = "end of CoverTab[1542]"
//line /usr/local/go/src/crypto/cipher/cfb.go:26
	}
//line /usr/local/go/src/crypto/cipher/cfb.go:26
	// _ = "end of CoverTab[1538]"
//line /usr/local/go/src/crypto/cipher/cfb.go:26
	_go_fuzz_dep_.CoverTab[1539]++
							if alias.InexactOverlap(dst[:len(src)], src) {
//line /usr/local/go/src/crypto/cipher/cfb.go:27
		_go_fuzz_dep_.CoverTab[1543]++
								panic("crypto/cipher: invalid buffer overlap")
//line /usr/local/go/src/crypto/cipher/cfb.go:28
		// _ = "end of CoverTab[1543]"
	} else {
//line /usr/local/go/src/crypto/cipher/cfb.go:29
		_go_fuzz_dep_.CoverTab[1544]++
//line /usr/local/go/src/crypto/cipher/cfb.go:29
		// _ = "end of CoverTab[1544]"
//line /usr/local/go/src/crypto/cipher/cfb.go:29
	}
//line /usr/local/go/src/crypto/cipher/cfb.go:29
	// _ = "end of CoverTab[1539]"
//line /usr/local/go/src/crypto/cipher/cfb.go:29
	_go_fuzz_dep_.CoverTab[1540]++
							for len(src) > 0 {
//line /usr/local/go/src/crypto/cipher/cfb.go:30
		_go_fuzz_dep_.CoverTab[1545]++
								if x.outUsed == len(x.out) {
//line /usr/local/go/src/crypto/cipher/cfb.go:31
			_go_fuzz_dep_.CoverTab[1549]++
									x.b.Encrypt(x.out, x.next)
									x.outUsed = 0
//line /usr/local/go/src/crypto/cipher/cfb.go:33
			// _ = "end of CoverTab[1549]"
		} else {
//line /usr/local/go/src/crypto/cipher/cfb.go:34
			_go_fuzz_dep_.CoverTab[1550]++
//line /usr/local/go/src/crypto/cipher/cfb.go:34
			// _ = "end of CoverTab[1550]"
//line /usr/local/go/src/crypto/cipher/cfb.go:34
		}
//line /usr/local/go/src/crypto/cipher/cfb.go:34
		// _ = "end of CoverTab[1545]"
//line /usr/local/go/src/crypto/cipher/cfb.go:34
		_go_fuzz_dep_.CoverTab[1546]++

								if x.decrypt {
//line /usr/local/go/src/crypto/cipher/cfb.go:36
			_go_fuzz_dep_.CoverTab[1551]++

//line /usr/local/go/src/crypto/cipher/cfb.go:41
			copy(x.next[x.outUsed:], src)
//line /usr/local/go/src/crypto/cipher/cfb.go:41
			// _ = "end of CoverTab[1551]"
		} else {
//line /usr/local/go/src/crypto/cipher/cfb.go:42
			_go_fuzz_dep_.CoverTab[1552]++
//line /usr/local/go/src/crypto/cipher/cfb.go:42
			// _ = "end of CoverTab[1552]"
//line /usr/local/go/src/crypto/cipher/cfb.go:42
		}
//line /usr/local/go/src/crypto/cipher/cfb.go:42
		// _ = "end of CoverTab[1546]"
//line /usr/local/go/src/crypto/cipher/cfb.go:42
		_go_fuzz_dep_.CoverTab[1547]++
								n := subtle.XORBytes(dst, src, x.out[x.outUsed:])
								if !x.decrypt {
//line /usr/local/go/src/crypto/cipher/cfb.go:44
			_go_fuzz_dep_.CoverTab[1553]++
									copy(x.next[x.outUsed:], dst)
//line /usr/local/go/src/crypto/cipher/cfb.go:45
			// _ = "end of CoverTab[1553]"
		} else {
//line /usr/local/go/src/crypto/cipher/cfb.go:46
			_go_fuzz_dep_.CoverTab[1554]++
//line /usr/local/go/src/crypto/cipher/cfb.go:46
			// _ = "end of CoverTab[1554]"
//line /usr/local/go/src/crypto/cipher/cfb.go:46
		}
//line /usr/local/go/src/crypto/cipher/cfb.go:46
		// _ = "end of CoverTab[1547]"
//line /usr/local/go/src/crypto/cipher/cfb.go:46
		_go_fuzz_dep_.CoverTab[1548]++
								dst = dst[n:]
								src = src[n:]
								x.outUsed += n
//line /usr/local/go/src/crypto/cipher/cfb.go:49
		// _ = "end of CoverTab[1548]"
	}
//line /usr/local/go/src/crypto/cipher/cfb.go:50
	// _ = "end of CoverTab[1540]"
}

// NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode,
//line /usr/local/go/src/crypto/cipher/cfb.go:53
// using the given Block. The iv must be the same length as the Block's block
//line /usr/local/go/src/crypto/cipher/cfb.go:53
// size.
//line /usr/local/go/src/crypto/cipher/cfb.go:56
func NewCFBEncrypter(block Block, iv []byte) Stream {
//line /usr/local/go/src/crypto/cipher/cfb.go:56
	_go_fuzz_dep_.CoverTab[1555]++
							return newCFB(block, iv, false)
//line /usr/local/go/src/crypto/cipher/cfb.go:57
	// _ = "end of CoverTab[1555]"
}

// NewCFBDecrypter returns a Stream which decrypts with cipher feedback mode,
//line /usr/local/go/src/crypto/cipher/cfb.go:60
// using the given Block. The iv must be the same length as the Block's block
//line /usr/local/go/src/crypto/cipher/cfb.go:60
// size.
//line /usr/local/go/src/crypto/cipher/cfb.go:63
func NewCFBDecrypter(block Block, iv []byte) Stream {
//line /usr/local/go/src/crypto/cipher/cfb.go:63
	_go_fuzz_dep_.CoverTab[1556]++
							return newCFB(block, iv, true)
//line /usr/local/go/src/crypto/cipher/cfb.go:64
	// _ = "end of CoverTab[1556]"
}

func newCFB(block Block, iv []byte, decrypt bool) Stream {
//line /usr/local/go/src/crypto/cipher/cfb.go:67
	_go_fuzz_dep_.CoverTab[1557]++
							blockSize := block.BlockSize()
							if len(iv) != blockSize {
//line /usr/local/go/src/crypto/cipher/cfb.go:69
		_go_fuzz_dep_.CoverTab[1559]++

								panic("cipher.newCFB: IV length must equal block size")
//line /usr/local/go/src/crypto/cipher/cfb.go:71
		// _ = "end of CoverTab[1559]"
	} else {
//line /usr/local/go/src/crypto/cipher/cfb.go:72
		_go_fuzz_dep_.CoverTab[1560]++
//line /usr/local/go/src/crypto/cipher/cfb.go:72
		// _ = "end of CoverTab[1560]"
//line /usr/local/go/src/crypto/cipher/cfb.go:72
	}
//line /usr/local/go/src/crypto/cipher/cfb.go:72
	// _ = "end of CoverTab[1557]"
//line /usr/local/go/src/crypto/cipher/cfb.go:72
	_go_fuzz_dep_.CoverTab[1558]++
							x := &cfb{
		b:		block,
		out:		make([]byte, blockSize),
		next:		make([]byte, blockSize),
		outUsed:	blockSize,
		decrypt:	decrypt,
	}
							copy(x.next, iv)

							return x
//line /usr/local/go/src/crypto/cipher/cfb.go:82
	// _ = "end of CoverTab[1558]"
}

//line /usr/local/go/src/crypto/cipher/cfb.go:83
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/cfb.go:83
var _ = _go_fuzz_dep_.CoverTab
