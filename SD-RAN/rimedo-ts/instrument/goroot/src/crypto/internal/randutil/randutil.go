// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/randutil/randutil.go:5
// Package randutil contains internal randomness utilities for various
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:5
// crypto packages.
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
package randutil

//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
import (
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
)
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
import (
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:7
)

import (
	"io"
	"sync"
)

var (
	closedChanOnce	sync.Once
	closedChan	chan struct{}
)

// MaybeReadByte reads a single byte from r with ~50% probability. This is used
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:19
// to ensure that callers do not depend on non-guaranteed behaviour, e.g.
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:19
// assuming that rsa.GenerateKey is deterministic w.r.t. a given random stream.
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:19
//
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:19
// This does not affect tests that pass a stream of fixed bytes as the random
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:19
// source (e.g. a zeroReader).
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:25
func MaybeReadByte(r io.Reader) {
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:25
	_go_fuzz_dep_.CoverTab[2771]++
									closedChanOnce.Do(func() {
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:26
		_go_fuzz_dep_.CoverTab[2773]++
										closedChan = make(chan struct{})
										close(closedChan)
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:28
		// _ = "end of CoverTab[2773]"
	})
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:29
	// _ = "end of CoverTab[2771]"
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:29
	_go_fuzz_dep_.CoverTab[2772]++

									select {
	case <-closedChan:
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:32
		_go_fuzz_dep_.CoverTab[2774]++
										return
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:33
		// _ = "end of CoverTab[2774]"
	case <-closedChan:
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:34
		_go_fuzz_dep_.CoverTab[2775]++
										var buf [1]byte
										r.Read(buf[:])
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:36
		// _ = "end of CoverTab[2775]"
	}
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:37
	// _ = "end of CoverTab[2772]"
}

//line /usr/local/go/src/crypto/internal/randutil/randutil.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/randutil/randutil.go:38
var _ = _go_fuzz_dep_.CoverTab
