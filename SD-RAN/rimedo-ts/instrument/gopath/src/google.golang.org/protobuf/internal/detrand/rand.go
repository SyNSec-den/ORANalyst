// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:5
// Package detrand provides deterministically random functionality.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:5
// The pseudo-randomness of these functions is seeded by the program binary
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:5
// itself and guarantees that the output does not change within a program,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:5
// while ensuring that the output is unstable across different builds.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
package detrand

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:10
)

import (
	"encoding/binary"
	"hash/fnv"
	"os"
)

// Disable disables detrand such that all functions returns the zero value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:18
// This function is not concurrent-safe and must be called during program init.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:20
func Disable() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:20
	_go_fuzz_dep_.CoverTab[48288]++
													randSeed = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:21
	// _ = "end of CoverTab[48288]"
}

// Bool returns a deterministically random boolean.
func Bool() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:25
	_go_fuzz_dep_.CoverTab[48289]++
													return randSeed%2 == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:26
	// _ = "end of CoverTab[48289]"
}

// Intn returns a deterministically random integer between 0 and n-1, inclusive.
func Intn(n int) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:30
	_go_fuzz_dep_.CoverTab[48290]++
													if n <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:31
		_go_fuzz_dep_.CoverTab[48292]++
														panic("must be positive")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:32
		// _ = "end of CoverTab[48292]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:33
		_go_fuzz_dep_.CoverTab[48293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:33
		// _ = "end of CoverTab[48293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:33
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:33
	// _ = "end of CoverTab[48290]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:33
	_go_fuzz_dep_.CoverTab[48291]++
													return int(randSeed % uint64(n))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:34
	// _ = "end of CoverTab[48291]"
}

// randSeed is a best-effort at an approximate hash of the Go binary.
var randSeed = binaryHash()

func binaryHash() uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:40
	_go_fuzz_dep_.CoverTab[48294]++

													s, err := os.Executable()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:43
		_go_fuzz_dep_.CoverTab[48299]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:44
		// _ = "end of CoverTab[48299]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:45
		_go_fuzz_dep_.CoverTab[48300]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:45
		// _ = "end of CoverTab[48300]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:45
	// _ = "end of CoverTab[48294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:45
	_go_fuzz_dep_.CoverTab[48295]++
													f, err := os.Open(s)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:47
		_go_fuzz_dep_.CoverTab[48301]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:48
		// _ = "end of CoverTab[48301]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:49
		_go_fuzz_dep_.CoverTab[48302]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:49
		// _ = "end of CoverTab[48302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:49
	// _ = "end of CoverTab[48295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:49
	_go_fuzz_dep_.CoverTab[48296]++
													defer f.Close()

	// Hash the size and several samples of the Go binary.
	const numSamples = 8
	var buf [64]byte
	h := fnv.New64()
	fi, err := f.Stat()
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:57
		_go_fuzz_dep_.CoverTab[48303]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:58
		// _ = "end of CoverTab[48303]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:59
		_go_fuzz_dep_.CoverTab[48304]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:59
		// _ = "end of CoverTab[48304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:59
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:59
	// _ = "end of CoverTab[48296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:59
	_go_fuzz_dep_.CoverTab[48297]++
													binary.LittleEndian.PutUint64(buf[:8], uint64(fi.Size()))
													h.Write(buf[:8])
													for i := int64(0); i < numSamples; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:62
		_go_fuzz_dep_.CoverTab[48305]++
														if _, err := f.ReadAt(buf[:], i*fi.Size()/numSamples); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:63
			_go_fuzz_dep_.CoverTab[48307]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:64
			// _ = "end of CoverTab[48307]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:65
			_go_fuzz_dep_.CoverTab[48308]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:65
			// _ = "end of CoverTab[48308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:65
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:65
		// _ = "end of CoverTab[48305]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:65
		_go_fuzz_dep_.CoverTab[48306]++
														h.Write(buf[:])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:66
		// _ = "end of CoverTab[48306]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:67
	// _ = "end of CoverTab[48297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:67
	_go_fuzz_dep_.CoverTab[48298]++
													return h.Sum64()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:68
	// _ = "end of CoverTab[48298]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/detrand/rand.go:69
var _ = _go_fuzz_dep_.CoverTab
