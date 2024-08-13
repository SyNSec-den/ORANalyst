// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
package subtle

//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
import (
//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
)
//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
import (
//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/subtle/xor_amd64.go:7
)

//go:noescape
func xorBytes(dst, a, b *byte, n int)

//line /usr/local/go/src/crypto/subtle/xor_amd64.go:10
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/subtle/xor_amd64.go:10
var _ = _go_fuzz_dep_.CoverTab
