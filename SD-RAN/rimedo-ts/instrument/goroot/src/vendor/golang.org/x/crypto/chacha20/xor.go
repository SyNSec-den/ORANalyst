// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found src the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
package chacha20

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:5
)

import "runtime"

// Platforms that have fast unaligned 32-bit little endian accesses.
const unaligned = runtime.GOARCH == "386" ||
	runtime.GOARCH == "amd64" ||
	runtime.GOARCH == "arm64" ||
	runtime.GOARCH == "ppc64le" ||
	runtime.GOARCH == "s390x"

// addXor reads a little endian uint32 from src, XORs it with (a + b) and
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:16
// places the result in little endian byte order in dst.
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:18
func addXor(dst, src []byte, a, b uint32) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:18
	_go_fuzz_dep_.CoverTab[20733]++
									_, _ = src[3], dst[3]
									if unaligned {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:20
		_go_fuzz_dep_.CoverTab[20734]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:26
		v := uint32(src[0])
										v |= uint32(src[1]) << 8
										v |= uint32(src[2]) << 16
										v |= uint32(src[3]) << 24
										v ^= a + b
										dst[0] = byte(v)
										dst[1] = byte(v >> 8)
										dst[2] = byte(v >> 16)
										dst[3] = byte(v >> 24)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:34
		// _ = "end of CoverTab[20734]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:35
		_go_fuzz_dep_.CoverTab[20735]++
										a += b
										dst[0] = src[0] ^ byte(a)
										dst[1] = src[1] ^ byte(a>>8)
										dst[2] = src[2] ^ byte(a>>16)
										dst[3] = src[3] ^ byte(a>>24)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:40
		// _ = "end of CoverTab[20735]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:41
	// _ = "end of CoverTab[20733]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/xor.go:42
var _ = _go_fuzz_dep_.CoverTab
