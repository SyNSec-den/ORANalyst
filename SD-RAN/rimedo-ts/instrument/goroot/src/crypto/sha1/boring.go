// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Extra indirection here so that when building go_bootstrap
// cmd/internal/boring is not even imported, so that we don't
// have to maintain changes to cmd/dist's deps graph.

//go:build !cmd_go_bootstrap && cgo
// +build !cmd_go_bootstrap,cgo

//line /usr/local/go/src/crypto/sha1/boring.go:12
package sha1

//line /usr/local/go/src/crypto/sha1/boring.go:12
import (
//line /usr/local/go/src/crypto/sha1/boring.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha1/boring.go:12
)
//line /usr/local/go/src/crypto/sha1/boring.go:12
import (
//line /usr/local/go/src/crypto/sha1/boring.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha1/boring.go:12
)

import (
	"crypto/internal/boring"
	"hash"
)

const boringEnabled = boring.Enabled

func boringNewSHA1() hash.Hash {
//line /usr/local/go/src/crypto/sha1/boring.go:21
	_go_fuzz_dep_.CoverTab[10081]++
//line /usr/local/go/src/crypto/sha1/boring.go:21
	return boring.NewSHA1()
//line /usr/local/go/src/crypto/sha1/boring.go:21
	// _ = "end of CoverTab[10081]"
//line /usr/local/go/src/crypto/sha1/boring.go:21
}

func boringUnreachable() {
//line /usr/local/go/src/crypto/sha1/boring.go:23
	_go_fuzz_dep_.CoverTab[10082]++
//line /usr/local/go/src/crypto/sha1/boring.go:23
	boring.Unreachable()
//line /usr/local/go/src/crypto/sha1/boring.go:23
	// _ = "end of CoverTab[10082]"
//line /usr/local/go/src/crypto/sha1/boring.go:23
}

func boringSHA1(p []byte) [20]byte {
//line /usr/local/go/src/crypto/sha1/boring.go:25
	_go_fuzz_dep_.CoverTab[10083]++
//line /usr/local/go/src/crypto/sha1/boring.go:25
	return boring.SHA1(p)
//line /usr/local/go/src/crypto/sha1/boring.go:25
	// _ = "end of CoverTab[10083]"
//line /usr/local/go/src/crypto/sha1/boring.go:25
}

//line /usr/local/go/src/crypto/sha1/boring.go:25
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha1/boring.go:25
var _ = _go_fuzz_dep_.CoverTab
