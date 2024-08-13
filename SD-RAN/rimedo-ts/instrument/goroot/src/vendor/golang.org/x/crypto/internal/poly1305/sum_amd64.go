// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc && !purego
// +build gc,!purego

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
package poly1305

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:8
)

//go:noescape
func update(state *macState, msg []byte)

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:18
type mac struct{ macGeneric }

func (h *mac) Write(p []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:20
	_go_fuzz_dep_.CoverTab[20749]++
											nn := len(p)
											if h.offset > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:22
		_go_fuzz_dep_.CoverTab[20753]++
												n := copy(h.buffer[h.offset:], p)
												if h.offset+n < TagSize {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:24
			_go_fuzz_dep_.CoverTab[20755]++
													h.offset += n
													return nn, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:26
			// _ = "end of CoverTab[20755]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:27
			_go_fuzz_dep_.CoverTab[20756]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:27
			// _ = "end of CoverTab[20756]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:27
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:27
		// _ = "end of CoverTab[20753]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:27
		_go_fuzz_dep_.CoverTab[20754]++
												p = p[n:]
												h.offset = 0
												update(&h.macState, h.buffer[:])
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:30
		// _ = "end of CoverTab[20754]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:31
		_go_fuzz_dep_.CoverTab[20757]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:31
		// _ = "end of CoverTab[20757]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:31
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:31
	// _ = "end of CoverTab[20749]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:31
	_go_fuzz_dep_.CoverTab[20750]++
											if n := len(p) - (len(p) % TagSize); n > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:32
		_go_fuzz_dep_.CoverTab[20758]++
												update(&h.macState, p[:n])
												p = p[n:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:34
		// _ = "end of CoverTab[20758]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:35
		_go_fuzz_dep_.CoverTab[20759]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:35
		// _ = "end of CoverTab[20759]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:35
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:35
	// _ = "end of CoverTab[20750]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:35
	_go_fuzz_dep_.CoverTab[20751]++
											if len(p) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:36
		_go_fuzz_dep_.CoverTab[20760]++
												h.offset += copy(h.buffer[h.offset:], p)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:37
		// _ = "end of CoverTab[20760]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:38
		_go_fuzz_dep_.CoverTab[20761]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:38
		// _ = "end of CoverTab[20761]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:38
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:38
	// _ = "end of CoverTab[20751]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:38
	_go_fuzz_dep_.CoverTab[20752]++
											return nn, nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:39
	// _ = "end of CoverTab[20752]"
}

func (h *mac) Sum(out *[16]byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:42
	_go_fuzz_dep_.CoverTab[20762]++
											state := h.macState
											if h.offset > 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:44
		_go_fuzz_dep_.CoverTab[20764]++
												update(&state, h.buffer[:h.offset])
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:45
		// _ = "end of CoverTab[20764]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:46
		_go_fuzz_dep_.CoverTab[20765]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:46
		// _ = "end of CoverTab[20765]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:46
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:46
	// _ = "end of CoverTab[20762]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:46
	_go_fuzz_dep_.CoverTab[20763]++
											finalize(out, &state.h, &state.s)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:47
	// _ = "end of CoverTab[20763]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/internal/poly1305/sum_amd64.go:48
var _ = _go_fuzz_dep_.CoverTab
