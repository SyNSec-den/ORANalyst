// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/root.go:5
package x509

//line /usr/local/go/src/crypto/x509/root.go:5
import (
//line /usr/local/go/src/crypto/x509/root.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/root.go:5
)
//line /usr/local/go/src/crypto/x509/root.go:5
import (
//line /usr/local/go/src/crypto/x509/root.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/root.go:5
)

import (
	"internal/godebug"
	"sync"
)

var (
	once		sync.Once
	systemRootsMu	sync.RWMutex
	systemRoots	*CertPool
	systemRootsErr	error
	fallbacksSet	bool
)

func systemRootsPool() *CertPool {
//line /usr/local/go/src/crypto/x509/root.go:20
	_go_fuzz_dep_.CoverTab[19250]++
							once.Do(initSystemRoots)
							systemRootsMu.RLock()
							defer systemRootsMu.RUnlock()
							return systemRoots
//line /usr/local/go/src/crypto/x509/root.go:24
	// _ = "end of CoverTab[19250]"
}

func initSystemRoots() {
//line /usr/local/go/src/crypto/x509/root.go:27
	_go_fuzz_dep_.CoverTab[19251]++
							systemRootsMu.Lock()
							defer systemRootsMu.Unlock()
							systemRoots, systemRootsErr = loadSystemRoots()
							if systemRootsErr != nil {
//line /usr/local/go/src/crypto/x509/root.go:31
		_go_fuzz_dep_.CoverTab[19252]++
								systemRoots = nil
//line /usr/local/go/src/crypto/x509/root.go:32
		// _ = "end of CoverTab[19252]"
	} else {
//line /usr/local/go/src/crypto/x509/root.go:33
		_go_fuzz_dep_.CoverTab[19253]++
//line /usr/local/go/src/crypto/x509/root.go:33
		// _ = "end of CoverTab[19253]"
//line /usr/local/go/src/crypto/x509/root.go:33
	}
//line /usr/local/go/src/crypto/x509/root.go:33
	// _ = "end of CoverTab[19251]"
}

var forceFallback = godebug.New("x509usefallbackroots")

// SetFallbackRoots sets the roots to use during certificate verification, if no
//line /usr/local/go/src/crypto/x509/root.go:38
// custom roots are specified and a platform verifier or a system certificate
//line /usr/local/go/src/crypto/x509/root.go:38
// pool is not available (for instance in a container which does not have a root
//line /usr/local/go/src/crypto/x509/root.go:38
// certificate bundle). SetFallbackRoots will panic if roots is nil.
//line /usr/local/go/src/crypto/x509/root.go:38
//
//line /usr/local/go/src/crypto/x509/root.go:38
// SetFallbackRoots may only be called once, if called multiple times it will
//line /usr/local/go/src/crypto/x509/root.go:38
// panic.
//line /usr/local/go/src/crypto/x509/root.go:38
//
//line /usr/local/go/src/crypto/x509/root.go:38
// The fallback behavior can be forced on all platforms, even when there is a
//line /usr/local/go/src/crypto/x509/root.go:38
// system certificate pool, by setting GODEBUG=x509usefallbackroots=1 (note that
//line /usr/local/go/src/crypto/x509/root.go:38
// on Windows and macOS this will disable usage of the platform verification
//line /usr/local/go/src/crypto/x509/root.go:38
// APIs and cause the pure Go verifier to be used). Setting
//line /usr/local/go/src/crypto/x509/root.go:38
// x509usefallbackroots=1 without calling SetFallbackRoots has no effect.
//line /usr/local/go/src/crypto/x509/root.go:51
func SetFallbackRoots(roots *CertPool) {
//line /usr/local/go/src/crypto/x509/root.go:51
	_go_fuzz_dep_.CoverTab[19254]++
							if roots == nil {
//line /usr/local/go/src/crypto/x509/root.go:52
		_go_fuzz_dep_.CoverTab[19258]++
								panic("roots must be non-nil")
//line /usr/local/go/src/crypto/x509/root.go:53
		// _ = "end of CoverTab[19258]"
	} else {
//line /usr/local/go/src/crypto/x509/root.go:54
		_go_fuzz_dep_.CoverTab[19259]++
//line /usr/local/go/src/crypto/x509/root.go:54
		// _ = "end of CoverTab[19259]"
//line /usr/local/go/src/crypto/x509/root.go:54
	}
//line /usr/local/go/src/crypto/x509/root.go:54
	// _ = "end of CoverTab[19254]"
//line /usr/local/go/src/crypto/x509/root.go:54
	_go_fuzz_dep_.CoverTab[19255]++

//line /usr/local/go/src/crypto/x509/root.go:58
	_ = systemRootsPool()

	systemRootsMu.Lock()
	defer systemRootsMu.Unlock()

	if fallbacksSet {
//line /usr/local/go/src/crypto/x509/root.go:63
		_go_fuzz_dep_.CoverTab[19260]++
								panic("SetFallbackRoots has already been called")
//line /usr/local/go/src/crypto/x509/root.go:64
		// _ = "end of CoverTab[19260]"
	} else {
//line /usr/local/go/src/crypto/x509/root.go:65
		_go_fuzz_dep_.CoverTab[19261]++
//line /usr/local/go/src/crypto/x509/root.go:65
		// _ = "end of CoverTab[19261]"
//line /usr/local/go/src/crypto/x509/root.go:65
	}
//line /usr/local/go/src/crypto/x509/root.go:65
	// _ = "end of CoverTab[19255]"
//line /usr/local/go/src/crypto/x509/root.go:65
	_go_fuzz_dep_.CoverTab[19256]++
							fallbacksSet = true

							if systemRoots != nil && func() bool {
//line /usr/local/go/src/crypto/x509/root.go:68
		_go_fuzz_dep_.CoverTab[19262]++
//line /usr/local/go/src/crypto/x509/root.go:68
		return (systemRoots.len() > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/root.go:68
			_go_fuzz_dep_.CoverTab[19263]++
//line /usr/local/go/src/crypto/x509/root.go:68
			return systemRoots.systemPool
//line /usr/local/go/src/crypto/x509/root.go:68
			// _ = "end of CoverTab[19263]"
//line /usr/local/go/src/crypto/x509/root.go:68
		}())
//line /usr/local/go/src/crypto/x509/root.go:68
		// _ = "end of CoverTab[19262]"
//line /usr/local/go/src/crypto/x509/root.go:68
	}() && func() bool {
//line /usr/local/go/src/crypto/x509/root.go:68
		_go_fuzz_dep_.CoverTab[19264]++
//line /usr/local/go/src/crypto/x509/root.go:68
		return forceFallback.Value() != "1"
//line /usr/local/go/src/crypto/x509/root.go:68
		// _ = "end of CoverTab[19264]"
//line /usr/local/go/src/crypto/x509/root.go:68
	}() {
//line /usr/local/go/src/crypto/x509/root.go:68
		_go_fuzz_dep_.CoverTab[19265]++
								return
//line /usr/local/go/src/crypto/x509/root.go:69
		// _ = "end of CoverTab[19265]"
	} else {
//line /usr/local/go/src/crypto/x509/root.go:70
		_go_fuzz_dep_.CoverTab[19266]++
//line /usr/local/go/src/crypto/x509/root.go:70
		// _ = "end of CoverTab[19266]"
//line /usr/local/go/src/crypto/x509/root.go:70
	}
//line /usr/local/go/src/crypto/x509/root.go:70
	// _ = "end of CoverTab[19256]"
//line /usr/local/go/src/crypto/x509/root.go:70
	_go_fuzz_dep_.CoverTab[19257]++
							systemRoots, systemRootsErr = roots, nil
//line /usr/local/go/src/crypto/x509/root.go:71
	// _ = "end of CoverTab[19257]"
}

//line /usr/local/go/src/crypto/x509/root.go:72
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/root.go:72
var _ = _go_fuzz_dep_.CoverTab
