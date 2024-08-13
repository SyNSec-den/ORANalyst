// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:5
/*
Package pbkdf2 implements the key derivation function PBKDF2 as defined in RFC
2898 / PKCS #5 v2.0.

A key derivation function is useful when encrypting data based on a password
or any other not-fully-random data. It uses a pseudorandom function to derive
a secure encryption key based on the password.

While v2.0 of the standard defines only one pseudorandom function to use,
HMAC-SHA1, the drafted v2.1 specification allows use of all five FIPS Approved
Hash Functions SHA-1, SHA-224, SHA-256, SHA-384 and SHA-512 for HMAC. To
choose, you can pass the `New` functions from the different SHA packages to
pbkdf2.Key.
*/
package pbkdf2

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:19
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:19
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:19
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:19
)

import (
	"crypto/hmac"
	"hash"
)

// Key derives a key from the password, salt and iteration count, returning a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// []byte of length keylen that can be used as cryptographic key. The key is
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// derived based on the method described as PBKDF2 with the HMAC variant using
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// the supplied hash function.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// For example, to use a HMAC-SHA-1 based PBKDF2 key derivation function, you
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// can get a derived key for e.g. AES-256 (which needs a 32-byte key) by
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// doing:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
//	dk := pbkdf2.Key([]byte("some password"), salt, 4096, 32, sha1.New)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// Remember to get a good random salt. At least 8 bytes is recommended by the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// RFC.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// Using a higher iteration count will increase the cost of an exhaustive
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:26
// search but will also make derivation proportionally slower.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:42
func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:42
	_go_fuzz_dep_.CoverTab[85703]++
													return Key64(password, salt, int64(iter), int64(keyLen), h)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:43
	// _ = "end of CoverTab[85703]"
}

// Key64 derives a key from the password, salt and iteration count, returning a
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// []byte of length keylen that can be used as cryptographic key. Key64 uses
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// int64 for the iteration count and key length to allow larger values.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// The key is derived based on the method described as PBKDF2 with the HMAC
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// variant using the supplied hash function.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// For example, to use a HMAC-SHA-1 based PBKDF2 key derivation function, you
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// can get a derived key for e.g. AES-256 (which needs a 32-byte key) by
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// doing:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
//	dk := pbkdf2.Key([]byte("some password"), salt, 4096, 32, sha1.New)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// Remember to get a good random salt. At least 8 bytes is recommended by the
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// RFC.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// Using a higher iteration count will increase the cost of an exhaustive
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:46
// search but will also make derivation proportionally slower.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:63
func Key64(password, salt []byte, iter, keyLen int64, h func() hash.Hash) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:63
	_go_fuzz_dep_.CoverTab[85704]++
													prf := hmac.New(h, password)
													hashLen := int64(prf.Size())
													numBlocks := (keyLen + hashLen - 1) / hashLen

													var buf [4]byte
													dk := make([]byte, 0, numBlocks*hashLen)
													U := make([]byte, hashLen)
													for block := int64(1); block <= numBlocks; block++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:71
		_go_fuzz_dep_.CoverTab[85706]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:75
		prf.Reset()
														prf.Write(salt)
														buf[0] = byte(block >> 24)
														buf[1] = byte(block >> 16)
														buf[2] = byte(block >> 8)
														buf[3] = byte(block)
														prf.Write(buf[:4])
														dk = prf.Sum(dk)
														T := dk[int64(len(dk))-hashLen:]
														copy(U, T)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:87
		for n := int64(2); n <= iter; n++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:87
			_go_fuzz_dep_.CoverTab[85707]++
															prf.Reset()
															prf.Write(U)
															U = U[:0]
															U = prf.Sum(U)
															for x := range U {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:92
				_go_fuzz_dep_.CoverTab[85708]++
																T[x] ^= U[x]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:93
				// _ = "end of CoverTab[85708]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:94
			// _ = "end of CoverTab[85707]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:95
		// _ = "end of CoverTab[85706]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:96
	// _ = "end of CoverTab[85704]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:96
	_go_fuzz_dep_.CoverTab[85705]++
													return dk[:keyLen]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:97
	// _ = "end of CoverTab[85705]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gofork@v1.0.0/x/crypto/pbkdf2/pbkdf2.go:98
var _ = _go_fuzz_dep_.CoverTab
