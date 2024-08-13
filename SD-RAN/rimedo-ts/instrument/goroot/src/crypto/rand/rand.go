// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/rand/rand.go:5
// Package rand implements a cryptographically secure
//line /usr/local/go/src/crypto/rand/rand.go:5
// random number generator.
//line /usr/local/go/src/crypto/rand/rand.go:7
package rand

//line /usr/local/go/src/crypto/rand/rand.go:7
import (
//line /usr/local/go/src/crypto/rand/rand.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rand/rand.go:7
)
//line /usr/local/go/src/crypto/rand/rand.go:7
import (
//line /usr/local/go/src/crypto/rand/rand.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rand/rand.go:7
)

import "io"

// Reader is a global, shared instance of a cryptographically
//line /usr/local/go/src/crypto/rand/rand.go:11
// secure random number generator.
//line /usr/local/go/src/crypto/rand/rand.go:11
//
//line /usr/local/go/src/crypto/rand/rand.go:11
// On Linux, FreeBSD, Dragonfly and Solaris, Reader uses getrandom(2) if
//line /usr/local/go/src/crypto/rand/rand.go:11
// available, /dev/urandom otherwise.
//line /usr/local/go/src/crypto/rand/rand.go:11
// On OpenBSD and macOS, Reader uses getentropy(2).
//line /usr/local/go/src/crypto/rand/rand.go:11
// On other Unix-like systems, Reader reads from /dev/urandom.
//line /usr/local/go/src/crypto/rand/rand.go:11
// On Windows systems, Reader uses the RtlGenRandom API.
//line /usr/local/go/src/crypto/rand/rand.go:11
// On Wasm, Reader uses the Web Crypto API.
//line /usr/local/go/src/crypto/rand/rand.go:20
var Reader io.Reader

// Read is a helper function that calls Reader.Read using io.ReadFull.
//line /usr/local/go/src/crypto/rand/rand.go:22
// On return, n == len(b) if and only if err == nil.
//line /usr/local/go/src/crypto/rand/rand.go:24
func Read(b []byte) (n int, err error) {
//line /usr/local/go/src/crypto/rand/rand.go:24
	_go_fuzz_dep_.CoverTab[9321]++
							return io.ReadFull(Reader, b)
//line /usr/local/go/src/crypto/rand/rand.go:25
	// _ = "end of CoverTab[9321]"
}

// batched returns a function that calls f to populate a []byte by chunking it
//line /usr/local/go/src/crypto/rand/rand.go:28
// into subslices of, at most, readMax bytes.
//line /usr/local/go/src/crypto/rand/rand.go:30
func batched(f func([]byte) error, readMax int) func([]byte) error {
//line /usr/local/go/src/crypto/rand/rand.go:30
	_go_fuzz_dep_.CoverTab[9322]++
							return func(out []byte) error {
//line /usr/local/go/src/crypto/rand/rand.go:31
		_go_fuzz_dep_.CoverTab[9323]++
								for len(out) > 0 {
//line /usr/local/go/src/crypto/rand/rand.go:32
			_go_fuzz_dep_.CoverTab[9325]++
									read := len(out)
									if read > readMax {
//line /usr/local/go/src/crypto/rand/rand.go:34
				_go_fuzz_dep_.CoverTab[9328]++
										read = readMax
//line /usr/local/go/src/crypto/rand/rand.go:35
				// _ = "end of CoverTab[9328]"
			} else {
//line /usr/local/go/src/crypto/rand/rand.go:36
				_go_fuzz_dep_.CoverTab[9329]++
//line /usr/local/go/src/crypto/rand/rand.go:36
				// _ = "end of CoverTab[9329]"
//line /usr/local/go/src/crypto/rand/rand.go:36
			}
//line /usr/local/go/src/crypto/rand/rand.go:36
			// _ = "end of CoverTab[9325]"
//line /usr/local/go/src/crypto/rand/rand.go:36
			_go_fuzz_dep_.CoverTab[9326]++
									if err := f(out[:read]); err != nil {
//line /usr/local/go/src/crypto/rand/rand.go:37
				_go_fuzz_dep_.CoverTab[9330]++
										return err
//line /usr/local/go/src/crypto/rand/rand.go:38
				// _ = "end of CoverTab[9330]"
			} else {
//line /usr/local/go/src/crypto/rand/rand.go:39
				_go_fuzz_dep_.CoverTab[9331]++
//line /usr/local/go/src/crypto/rand/rand.go:39
				// _ = "end of CoverTab[9331]"
//line /usr/local/go/src/crypto/rand/rand.go:39
			}
//line /usr/local/go/src/crypto/rand/rand.go:39
			// _ = "end of CoverTab[9326]"
//line /usr/local/go/src/crypto/rand/rand.go:39
			_go_fuzz_dep_.CoverTab[9327]++
									out = out[read:]
//line /usr/local/go/src/crypto/rand/rand.go:40
			// _ = "end of CoverTab[9327]"
		}
//line /usr/local/go/src/crypto/rand/rand.go:41
		// _ = "end of CoverTab[9323]"
//line /usr/local/go/src/crypto/rand/rand.go:41
		_go_fuzz_dep_.CoverTab[9324]++
								return nil
//line /usr/local/go/src/crypto/rand/rand.go:42
		// _ = "end of CoverTab[9324]"
	}
//line /usr/local/go/src/crypto/rand/rand.go:43
	// _ = "end of CoverTab[9322]"
}

//line /usr/local/go/src/crypto/rand/rand.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rand/rand.go:44
var _ = _go_fuzz_dep_.CoverTab
