// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/rc4/rc4.go:5
// Package rc4 implements RC4 encryption, as defined in Bruce Schneier's
//line /usr/local/go/src/crypto/rc4/rc4.go:5
// Applied Cryptography.
//line /usr/local/go/src/crypto/rc4/rc4.go:5
//
//line /usr/local/go/src/crypto/rc4/rc4.go:5
// RC4 is cryptographically broken and should not be used for secure
//line /usr/local/go/src/crypto/rc4/rc4.go:5
// applications.
//line /usr/local/go/src/crypto/rc4/rc4.go:10
package rc4

//line /usr/local/go/src/crypto/rc4/rc4.go:10
import (
//line /usr/local/go/src/crypto/rc4/rc4.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rc4/rc4.go:10
)
//line /usr/local/go/src/crypto/rc4/rc4.go:10
import (
//line /usr/local/go/src/crypto/rc4/rc4.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rc4/rc4.go:10
)

import (
	"crypto/internal/alias"
	"strconv"
)

// A Cipher is an instance of RC4 using a particular key.
type Cipher struct {
	s	[256]uint32
	i, j	uint8
}

type KeySizeError int

func (k KeySizeError) Error() string {
//line /usr/local/go/src/crypto/rc4/rc4.go:25
	_go_fuzz_dep_.CoverTab[9595]++
						return "crypto/rc4: invalid key size " + strconv.Itoa(int(k))
//line /usr/local/go/src/crypto/rc4/rc4.go:26
	// _ = "end of CoverTab[9595]"
}

// NewCipher creates and returns a new Cipher. The key argument should be the
//line /usr/local/go/src/crypto/rc4/rc4.go:29
// RC4 key, at least 1 byte and at most 256 bytes.
//line /usr/local/go/src/crypto/rc4/rc4.go:31
func NewCipher(key []byte) (*Cipher, error) {
//line /usr/local/go/src/crypto/rc4/rc4.go:31
	_go_fuzz_dep_.CoverTab[9596]++
						k := len(key)
						if k < 1 || func() bool {
//line /usr/local/go/src/crypto/rc4/rc4.go:33
		_go_fuzz_dep_.CoverTab[9600]++
//line /usr/local/go/src/crypto/rc4/rc4.go:33
		return k > 256
//line /usr/local/go/src/crypto/rc4/rc4.go:33
		// _ = "end of CoverTab[9600]"
//line /usr/local/go/src/crypto/rc4/rc4.go:33
	}() {
//line /usr/local/go/src/crypto/rc4/rc4.go:33
		_go_fuzz_dep_.CoverTab[9601]++
							return nil, KeySizeError(k)
//line /usr/local/go/src/crypto/rc4/rc4.go:34
		// _ = "end of CoverTab[9601]"
	} else {
//line /usr/local/go/src/crypto/rc4/rc4.go:35
		_go_fuzz_dep_.CoverTab[9602]++
//line /usr/local/go/src/crypto/rc4/rc4.go:35
		// _ = "end of CoverTab[9602]"
//line /usr/local/go/src/crypto/rc4/rc4.go:35
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:35
	// _ = "end of CoverTab[9596]"
//line /usr/local/go/src/crypto/rc4/rc4.go:35
	_go_fuzz_dep_.CoverTab[9597]++
						var c Cipher
						for i := 0; i < 256; i++ {
//line /usr/local/go/src/crypto/rc4/rc4.go:37
		_go_fuzz_dep_.CoverTab[9603]++
							c.s[i] = uint32(i)
//line /usr/local/go/src/crypto/rc4/rc4.go:38
		// _ = "end of CoverTab[9603]"
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:39
	// _ = "end of CoverTab[9597]"
//line /usr/local/go/src/crypto/rc4/rc4.go:39
	_go_fuzz_dep_.CoverTab[9598]++
						var j uint8 = 0
						for i := 0; i < 256; i++ {
//line /usr/local/go/src/crypto/rc4/rc4.go:41
		_go_fuzz_dep_.CoverTab[9604]++
							j += uint8(c.s[i]) + key[i%k]
							c.s[i], c.s[j] = c.s[j], c.s[i]
//line /usr/local/go/src/crypto/rc4/rc4.go:43
		// _ = "end of CoverTab[9604]"
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:44
	// _ = "end of CoverTab[9598]"
//line /usr/local/go/src/crypto/rc4/rc4.go:44
	_go_fuzz_dep_.CoverTab[9599]++
						return &c, nil
//line /usr/local/go/src/crypto/rc4/rc4.go:45
	// _ = "end of CoverTab[9599]"
}

// Reset zeros the key data and makes the Cipher unusable.
//line /usr/local/go/src/crypto/rc4/rc4.go:48
//
//line /usr/local/go/src/crypto/rc4/rc4.go:48
// Deprecated: Reset can't guarantee that the key will be entirely removed from
//line /usr/local/go/src/crypto/rc4/rc4.go:48
// the process's memory.
//line /usr/local/go/src/crypto/rc4/rc4.go:52
func (c *Cipher) Reset() {
//line /usr/local/go/src/crypto/rc4/rc4.go:52
	_go_fuzz_dep_.CoverTab[9605]++
						for i := range c.s {
//line /usr/local/go/src/crypto/rc4/rc4.go:53
		_go_fuzz_dep_.CoverTab[9607]++
							c.s[i] = 0
//line /usr/local/go/src/crypto/rc4/rc4.go:54
		// _ = "end of CoverTab[9607]"
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:55
	// _ = "end of CoverTab[9605]"
//line /usr/local/go/src/crypto/rc4/rc4.go:55
	_go_fuzz_dep_.CoverTab[9606]++
						c.i, c.j = 0, 0
//line /usr/local/go/src/crypto/rc4/rc4.go:56
	// _ = "end of CoverTab[9606]"
}

// XORKeyStream sets dst to the result of XORing src with the key stream.
//line /usr/local/go/src/crypto/rc4/rc4.go:59
// Dst and src must overlap entirely or not at all.
//line /usr/local/go/src/crypto/rc4/rc4.go:61
func (c *Cipher) XORKeyStream(dst, src []byte) {
//line /usr/local/go/src/crypto/rc4/rc4.go:61
	_go_fuzz_dep_.CoverTab[9608]++
						if len(src) == 0 {
//line /usr/local/go/src/crypto/rc4/rc4.go:62
		_go_fuzz_dep_.CoverTab[9612]++
							return
//line /usr/local/go/src/crypto/rc4/rc4.go:63
		// _ = "end of CoverTab[9612]"
	} else {
//line /usr/local/go/src/crypto/rc4/rc4.go:64
		_go_fuzz_dep_.CoverTab[9613]++
//line /usr/local/go/src/crypto/rc4/rc4.go:64
		// _ = "end of CoverTab[9613]"
//line /usr/local/go/src/crypto/rc4/rc4.go:64
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:64
	// _ = "end of CoverTab[9608]"
//line /usr/local/go/src/crypto/rc4/rc4.go:64
	_go_fuzz_dep_.CoverTab[9609]++
						if alias.InexactOverlap(dst[:len(src)], src) {
//line /usr/local/go/src/crypto/rc4/rc4.go:65
		_go_fuzz_dep_.CoverTab[9614]++
							panic("crypto/rc4: invalid buffer overlap")
//line /usr/local/go/src/crypto/rc4/rc4.go:66
		// _ = "end of CoverTab[9614]"
	} else {
//line /usr/local/go/src/crypto/rc4/rc4.go:67
		_go_fuzz_dep_.CoverTab[9615]++
//line /usr/local/go/src/crypto/rc4/rc4.go:67
		// _ = "end of CoverTab[9615]"
//line /usr/local/go/src/crypto/rc4/rc4.go:67
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:67
	// _ = "end of CoverTab[9609]"
//line /usr/local/go/src/crypto/rc4/rc4.go:67
	_go_fuzz_dep_.CoverTab[9610]++
						i, j := c.i, c.j
						_ = dst[len(src)-1]
						dst = dst[:len(src)]
						for k, v := range src {
//line /usr/local/go/src/crypto/rc4/rc4.go:71
		_go_fuzz_dep_.CoverTab[9616]++
							i += 1
							x := c.s[i]
							j += uint8(x)
							y := c.s[j]
							c.s[i], c.s[j] = y, x
							dst[k] = v ^ uint8(c.s[uint8(x+y)])
//line /usr/local/go/src/crypto/rc4/rc4.go:77
		// _ = "end of CoverTab[9616]"
	}
//line /usr/local/go/src/crypto/rc4/rc4.go:78
	// _ = "end of CoverTab[9610]"
//line /usr/local/go/src/crypto/rc4/rc4.go:78
	_go_fuzz_dep_.CoverTab[9611]++
						c.i, c.j = i, j
//line /usr/local/go/src/crypto/rc4/rc4.go:79
	// _ = "end of CoverTab[9611]"
}

//line /usr/local/go/src/crypto/rc4/rc4.go:80
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rc4/rc4.go:80
var _ = _go_fuzz_dep_.CoverTab
