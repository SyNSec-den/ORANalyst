// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64 || s390x || ppc64le || ppc64

//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
package sha256

//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
import (
//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
)
//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
import (
//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:7
)

//go:noescape

func block(dig *digest, p []byte)

//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:11
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha256/sha256block_decl.go:11
var _ = _go_fuzz_dep_.CoverTab
