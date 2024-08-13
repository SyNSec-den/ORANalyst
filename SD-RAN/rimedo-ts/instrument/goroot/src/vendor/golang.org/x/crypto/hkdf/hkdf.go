// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:5
// Package hkdf implements the HMAC-based Extract-and-Expand Key Derivation
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:5
// Function (HKDF) as defined in RFC 5869.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:5
// HKDF is a cryptographic key derivation function (KDF) with the goal of
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:5
// expanding limited input keying material into one or more cryptographically
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:5
// strong secret keys.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
package hkdf

//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:11
)

import (
	"crypto/hmac"
	"errors"
	"hash"
	"io"
)

// Extract generates a pseudorandom key for use with Expand from an input secret
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:20
// and an optional independent salt.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:20
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:20
// Only use this function if you need to reuse the extracted key with multiple
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:20
// Expand invocations and different context values. Most common scenarios,
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:20
// including the generation of multiple keys, should use New instead.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:26
func Extract(hash func() hash.Hash, secret, salt []byte) []byte {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:26
	_go_fuzz_dep_.CoverTab[20994]++
									if salt == nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:27
		_go_fuzz_dep_.CoverTab[20996]++
										salt = make([]byte, hash().Size())
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:28
		// _ = "end of CoverTab[20996]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:29
		_go_fuzz_dep_.CoverTab[20997]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:29
		// _ = "end of CoverTab[20997]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:29
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:29
	// _ = "end of CoverTab[20994]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:29
	_go_fuzz_dep_.CoverTab[20995]++
									extractor := hmac.New(hash, salt)
									extractor.Write(secret)
									return extractor.Sum(nil)
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:32
	// _ = "end of CoverTab[20995]"
}

type hkdf struct {
	expander	hash.Hash
	size		int

	info	[]byte
	counter	byte

	prev	[]byte
	buf	[]byte
}

func (f *hkdf) Read(p []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:46
	_go_fuzz_dep_.CoverTab[20998]++

									need := len(p)
									remains := len(f.buf) + int(255-f.counter+1)*f.size
									if remains < need {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:50
		_go_fuzz_dep_.CoverTab[21001]++
										return 0, errors.New("hkdf: entropy limit reached")
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:51
		// _ = "end of CoverTab[21001]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:52
		_go_fuzz_dep_.CoverTab[21002]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:52
		// _ = "end of CoverTab[21002]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:52
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:52
	// _ = "end of CoverTab[20998]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:52
	_go_fuzz_dep_.CoverTab[20999]++

									n := copy(p, f.buf)
									p = p[n:]

//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:58
	for len(p) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:58
		_go_fuzz_dep_.CoverTab[21003]++
										f.expander.Reset()
										f.expander.Write(f.prev)
										f.expander.Write(f.info)
										f.expander.Write([]byte{f.counter})
										f.prev = f.expander.Sum(f.prev[:0])
										f.counter++

//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:67
		f.buf = f.prev
										n = copy(p, f.buf)
										p = p[n:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:69
		// _ = "end of CoverTab[21003]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:70
	// _ = "end of CoverTab[20999]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:70
	_go_fuzz_dep_.CoverTab[21000]++

									f.buf = f.buf[n:]

									return need, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:74
	// _ = "end of CoverTab[21000]"
}

// Expand returns a Reader, from which keys can be read, using the given
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:77
// pseudorandom key and optional context info, skipping the extraction step.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:77
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:77
// The pseudorandomKey should have been generated by Extract, or be a uniformly
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:77
// random or pseudorandom cryptographically strong key. See RFC 5869, Section
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:77
// 3.3. Most common scenarios will want to use New instead.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:83
func Expand(hash func() hash.Hash, pseudorandomKey, info []byte) io.Reader {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:83
	_go_fuzz_dep_.CoverTab[21004]++
									expander := hmac.New(hash, pseudorandomKey)
									return &hkdf{expander, expander.Size(), info, 1, nil, nil}
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:85
	// _ = "end of CoverTab[21004]"
}

// New returns a Reader, from which keys can be read, using the given hash,
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:88
// secret, salt and context info. Salt and info can be nil.
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:90
func New(hash func() hash.Hash, secret, salt, info []byte) io.Reader {
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:90
	_go_fuzz_dep_.CoverTab[21005]++
									prk := Extract(hash, secret, salt)
									return Expand(hash, prk, info)
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:92
	// _ = "end of CoverTab[21005]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:93
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/hkdf/hkdf.go:93
var _ = _go_fuzz_dep_.CoverTab
