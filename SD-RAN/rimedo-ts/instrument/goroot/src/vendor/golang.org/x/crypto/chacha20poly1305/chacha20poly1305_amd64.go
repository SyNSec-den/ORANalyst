// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc && !purego
// +build gc,!purego

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
package chacha20poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:8
)

import (
	"encoding/binary"

	"golang.org/x/crypto/internal/alias"
	"golang.org/x/sys/cpu"
)

//go:noescape
func chacha20Poly1305Open(dst []byte, key []uint32, src, ad []byte) bool

//go:noescape
func chacha20Poly1305Seal(dst []byte, key []uint32, src, ad []byte)

var (
	useAVX2 = cpu.X86.HasAVX2 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:24
		_go_fuzz_dep_.CoverTab[20933]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:24
		return cpu.X86.HasBMI2
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:24
		// _ = "end of CoverTab[20933]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:24
	}()
)

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:29
func setupState(state *[16]uint32, key *[32]byte, nonce []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:29
	_go_fuzz_dep_.CoverTab[20934]++
													state[0] = 0x61707865
													state[1] = 0x3320646e
													state[2] = 0x79622d32
													state[3] = 0x6b206574

													state[4] = binary.LittleEndian.Uint32(key[0:4])
													state[5] = binary.LittleEndian.Uint32(key[4:8])
													state[6] = binary.LittleEndian.Uint32(key[8:12])
													state[7] = binary.LittleEndian.Uint32(key[12:16])
													state[8] = binary.LittleEndian.Uint32(key[16:20])
													state[9] = binary.LittleEndian.Uint32(key[20:24])
													state[10] = binary.LittleEndian.Uint32(key[24:28])
													state[11] = binary.LittleEndian.Uint32(key[28:32])

													state[12] = 0
													state[13] = binary.LittleEndian.Uint32(nonce[0:4])
													state[14] = binary.LittleEndian.Uint32(nonce[4:8])
													state[15] = binary.LittleEndian.Uint32(nonce[8:12])
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:47
	// _ = "end of CoverTab[20934]"
}

func (c *chacha20poly1305) seal(dst, nonce, plaintext, additionalData []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:50
	_go_fuzz_dep_.CoverTab[20935]++
													if !cpu.X86.HasSSSE3 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:51
		_go_fuzz_dep_.CoverTab[20938]++
														return c.sealGeneric(dst, nonce, plaintext, additionalData)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:52
		// _ = "end of CoverTab[20938]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:53
		_go_fuzz_dep_.CoverTab[20939]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:53
		// _ = "end of CoverTab[20939]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:53
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:53
	// _ = "end of CoverTab[20935]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:53
	_go_fuzz_dep_.CoverTab[20936]++

													var state [16]uint32
													setupState(&state, &c.key, nonce)

													ret, out := sliceForAppend(dst, len(plaintext)+16)
													if alias.InexactOverlap(out, plaintext) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:59
		_go_fuzz_dep_.CoverTab[20940]++
														panic("chacha20poly1305: invalid buffer overlap")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:60
		// _ = "end of CoverTab[20940]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:61
		_go_fuzz_dep_.CoverTab[20941]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:61
		// _ = "end of CoverTab[20941]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:61
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:61
	// _ = "end of CoverTab[20936]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:61
	_go_fuzz_dep_.CoverTab[20937]++
													chacha20Poly1305Seal(out[:], state[:], plaintext, additionalData)
													return ret
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:63
	// _ = "end of CoverTab[20937]"
}

func (c *chacha20poly1305) open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:66
	_go_fuzz_dep_.CoverTab[20942]++
													if !cpu.X86.HasSSSE3 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:67
		_go_fuzz_dep_.CoverTab[20946]++
														return c.openGeneric(dst, nonce, ciphertext, additionalData)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:68
		// _ = "end of CoverTab[20946]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:69
		_go_fuzz_dep_.CoverTab[20947]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:69
		// _ = "end of CoverTab[20947]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:69
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:69
	// _ = "end of CoverTab[20942]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:69
	_go_fuzz_dep_.CoverTab[20943]++

													var state [16]uint32
													setupState(&state, &c.key, nonce)

													ciphertext = ciphertext[:len(ciphertext)-16]
													ret, out := sliceForAppend(dst, len(ciphertext))
													if alias.InexactOverlap(out, ciphertext) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:76
		_go_fuzz_dep_.CoverTab[20948]++
														panic("chacha20poly1305: invalid buffer overlap")
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:77
		// _ = "end of CoverTab[20948]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:78
		_go_fuzz_dep_.CoverTab[20949]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:78
		// _ = "end of CoverTab[20949]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:78
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:78
	// _ = "end of CoverTab[20943]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:78
	_go_fuzz_dep_.CoverTab[20944]++
													if !chacha20Poly1305Open(out, state[:], ciphertext, additionalData) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:79
		_go_fuzz_dep_.CoverTab[20950]++
														for i := range out {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:80
			_go_fuzz_dep_.CoverTab[20952]++
															out[i] = 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:81
			// _ = "end of CoverTab[20952]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:82
		// _ = "end of CoverTab[20950]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:82
		_go_fuzz_dep_.CoverTab[20951]++
														return nil, errOpen
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:83
		// _ = "end of CoverTab[20951]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:84
		_go_fuzz_dep_.CoverTab[20953]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:84
		// _ = "end of CoverTab[20953]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:84
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:84
	// _ = "end of CoverTab[20944]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:84
	_go_fuzz_dep_.CoverTab[20945]++

													return ret, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:86
	// _ = "end of CoverTab[20945]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20poly1305/chacha20poly1305_amd64.go:87
var _ = _go_fuzz_dep_.CoverTab
