// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
package chacha20poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:5
)

import (
	"encoding/binary"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/internal/alias"
	"golang.org/x/crypto/internal/poly1305"
)

func writeWithPadding(p *poly1305.MAC, b []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:15
	_go_fuzz_dep_.CoverTab[20954]++
													p.Write(b)
													if rem := len(b) % 16; rem != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:17
		_go_fuzz_dep_.CoverTab[20955]++
														var buf [16]byte
														padLen := 16 - rem
														p.Write(buf[:padLen])
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:20
		// _ = "end of CoverTab[20955]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:21
		_go_fuzz_dep_.CoverTab[20956]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:21
		// _ = "end of CoverTab[20956]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:21
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:21
	// _ = "end of CoverTab[20954]"
}

func writeUint64(p *poly1305.MAC, n int) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:24
	_go_fuzz_dep_.CoverTab[20957]++
													var buf [8]byte
													binary.LittleEndian.PutUint64(buf[:], uint64(n))
													p.Write(buf[:])
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:27
	// _ = "end of CoverTab[20957]"
}

func (c *chacha20poly1305) sealGeneric(dst, nonce, plaintext, additionalData []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:30
	_go_fuzz_dep_.CoverTab[20958]++
													ret, out := sliceForAppend(dst, len(plaintext)+poly1305.TagSize)
													ciphertext, tag := out[:len(plaintext)], out[len(plaintext):]
													if alias.InexactOverlap(out, plaintext) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:33
		_go_fuzz_dep_.CoverTab[20960]++
														panic("chacha20poly1305: invalid buffer overlap")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:34
		// _ = "end of CoverTab[20960]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:35
		_go_fuzz_dep_.CoverTab[20961]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:35
		// _ = "end of CoverTab[20961]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:35
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:35
	// _ = "end of CoverTab[20958]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:35
	_go_fuzz_dep_.CoverTab[20959]++

													var polyKey [32]byte
													s, _ := chacha20.NewUnauthenticatedCipher(c.key[:], nonce)
													s.XORKeyStream(polyKey[:], polyKey[:])
													s.SetCounter(1)
													s.XORKeyStream(ciphertext, plaintext)

													p := poly1305.New(&polyKey)
													writeWithPadding(p, additionalData)
													writeWithPadding(p, ciphertext)
													writeUint64(p, len(additionalData))
													writeUint64(p, len(plaintext))
													p.Sum(tag[:0])

													return ret
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:50
	// _ = "end of CoverTab[20959]"
}

func (c *chacha20poly1305) openGeneric(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:53
	_go_fuzz_dep_.CoverTab[20962]++
													tag := ciphertext[len(ciphertext)-16:]
													ciphertext = ciphertext[:len(ciphertext)-16]

													var polyKey [32]byte
													s, _ := chacha20.NewUnauthenticatedCipher(c.key[:], nonce)
													s.XORKeyStream(polyKey[:], polyKey[:])
													s.SetCounter(1)

													p := poly1305.New(&polyKey)
													writeWithPadding(p, additionalData)
													writeWithPadding(p, ciphertext)
													writeUint64(p, len(additionalData))
													writeUint64(p, len(ciphertext))

													ret, out := sliceForAppend(dst, len(ciphertext))
													if alias.InexactOverlap(out, ciphertext) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:69
		_go_fuzz_dep_.CoverTab[20965]++
														panic("chacha20poly1305: invalid buffer overlap")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:70
		// _ = "end of CoverTab[20965]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:71
		_go_fuzz_dep_.CoverTab[20966]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:71
		// _ = "end of CoverTab[20966]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:71
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:71
	// _ = "end of CoverTab[20962]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:71
	_go_fuzz_dep_.CoverTab[20963]++
													if !p.Verify(tag) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:72
		_go_fuzz_dep_.CoverTab[20967]++
														for i := range out {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:73
			_go_fuzz_dep_.CoverTab[20969]++
															out[i] = 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:74
			// _ = "end of CoverTab[20969]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:75
		// _ = "end of CoverTab[20967]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:75
		_go_fuzz_dep_.CoverTab[20968]++
														return nil, errOpen
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:76
		// _ = "end of CoverTab[20968]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:77
		_go_fuzz_dep_.CoverTab[20970]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:77
		// _ = "end of CoverTab[20970]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:77
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:77
	// _ = "end of CoverTab[20963]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:77
	_go_fuzz_dep_.CoverTab[20964]++

													s.XORKeyStream(out, ciphertext)
													return ret, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:80
	// _ = "end of CoverTab[20964]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_generic.go:81
var _ = _go_fuzz_dep_.CoverTab
