// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/subtle/xor.go:5
package subtle

//line /usr/local/go/src/crypto/subtle/xor.go:5
import (
//line /usr/local/go/src/crypto/subtle/xor.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/subtle/xor.go:5
)
//line /usr/local/go/src/crypto/subtle/xor.go:5
import (
//line /usr/local/go/src/crypto/subtle/xor.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/subtle/xor.go:5
)

// XORBytes sets dst[i] = x[i] ^ y[i] for all i < n = min(len(x), len(y)),
//line /usr/local/go/src/crypto/subtle/xor.go:7
// returning n, the number of bytes written to dst.
//line /usr/local/go/src/crypto/subtle/xor.go:7
// If dst does not have length at least n,
//line /usr/local/go/src/crypto/subtle/xor.go:7
// XORBytes panics without writing anything to dst.
//line /usr/local/go/src/crypto/subtle/xor.go:11
func XORBytes(dst, x, y []byte) int {
//line /usr/local/go/src/crypto/subtle/xor.go:11
	_go_fuzz_dep_.CoverTab[1160]++
							n := len(x)
							if len(y) < n {
//line /usr/local/go/src/crypto/subtle/xor.go:13
		_go_fuzz_dep_.CoverTab[1164]++
								n = len(y)
//line /usr/local/go/src/crypto/subtle/xor.go:14
		// _ = "end of CoverTab[1164]"
	} else {
//line /usr/local/go/src/crypto/subtle/xor.go:15
		_go_fuzz_dep_.CoverTab[1165]++
//line /usr/local/go/src/crypto/subtle/xor.go:15
		// _ = "end of CoverTab[1165]"
//line /usr/local/go/src/crypto/subtle/xor.go:15
	}
//line /usr/local/go/src/crypto/subtle/xor.go:15
	// _ = "end of CoverTab[1160]"
//line /usr/local/go/src/crypto/subtle/xor.go:15
	_go_fuzz_dep_.CoverTab[1161]++
							if n == 0 {
//line /usr/local/go/src/crypto/subtle/xor.go:16
		_go_fuzz_dep_.CoverTab[1166]++
								return 0
//line /usr/local/go/src/crypto/subtle/xor.go:17
		// _ = "end of CoverTab[1166]"
	} else {
//line /usr/local/go/src/crypto/subtle/xor.go:18
		_go_fuzz_dep_.CoverTab[1167]++
//line /usr/local/go/src/crypto/subtle/xor.go:18
		// _ = "end of CoverTab[1167]"
//line /usr/local/go/src/crypto/subtle/xor.go:18
	}
//line /usr/local/go/src/crypto/subtle/xor.go:18
	// _ = "end of CoverTab[1161]"
//line /usr/local/go/src/crypto/subtle/xor.go:18
	_go_fuzz_dep_.CoverTab[1162]++
							if n > len(dst) {
//line /usr/local/go/src/crypto/subtle/xor.go:19
		_go_fuzz_dep_.CoverTab[1168]++
								panic("subtle.XORBytes: dst too short")
//line /usr/local/go/src/crypto/subtle/xor.go:20
		// _ = "end of CoverTab[1168]"
	} else {
//line /usr/local/go/src/crypto/subtle/xor.go:21
		_go_fuzz_dep_.CoverTab[1169]++
//line /usr/local/go/src/crypto/subtle/xor.go:21
		// _ = "end of CoverTab[1169]"
//line /usr/local/go/src/crypto/subtle/xor.go:21
	}
//line /usr/local/go/src/crypto/subtle/xor.go:21
	// _ = "end of CoverTab[1162]"
//line /usr/local/go/src/crypto/subtle/xor.go:21
	_go_fuzz_dep_.CoverTab[1163]++
							xorBytes(&dst[0], &x[0], &y[0], n)
							return n
//line /usr/local/go/src/crypto/subtle/xor.go:23
	// _ = "end of CoverTab[1163]"
}

//line /usr/local/go/src/crypto/subtle/xor.go:24
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/subtle/xor.go:24
var _ = _go_fuzz_dep_.CoverTab
