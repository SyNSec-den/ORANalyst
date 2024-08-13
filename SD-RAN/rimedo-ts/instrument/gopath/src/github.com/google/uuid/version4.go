// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:5
)

import "io"

// New creates a new random UUID or panics.  New is equivalent to
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:9
// the expression
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:9
//	uuid.Must(uuid.NewRandom())
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:13
func New() UUID {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:13
	_go_fuzz_dep_.CoverTab[179536]++
										return Must(NewRandom())
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:14
	// _ = "end of CoverTab[179536]"
}

// NewString creates a new random UUID and returns it as a string or panics.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:17
// NewString is equivalent to the expression
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:17
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:17
//	uuid.New().String()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:21
func NewString() string {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:21
	_go_fuzz_dep_.CoverTab[179537]++
										return Must(NewRandom()).String()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:22
	// _ = "end of CoverTab[179537]"
}

// NewRandom returns a Random (Version 4) UUID.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
// The strength of the UUIDs is based on the strength of the crypto/rand
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
// package.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
// Uses the randomness pool if it was enabled with EnableRandPool.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
// A note about uniqueness derived from the UUID Wikipedia entry:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//	Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//	hit by a meteorite is estimated to be one chance in 17 billion, that
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//	means the probability is about 0.00000000006 (6 × 10−11),
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//	equivalent to the odds of creating a few tens of trillions of UUIDs in a
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:25
//	year and having one duplicate.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:39
func NewRandom() (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:39
	_go_fuzz_dep_.CoverTab[179538]++
										if !poolEnabled {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:40
		_go_fuzz_dep_.CoverTab[179540]++
											return NewRandomFromReader(rander)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:41
		// _ = "end of CoverTab[179540]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:42
		_go_fuzz_dep_.CoverTab[179541]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:42
		// _ = "end of CoverTab[179541]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:42
	// _ = "end of CoverTab[179538]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:42
	_go_fuzz_dep_.CoverTab[179539]++
										return newRandomFromPool()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:43
	// _ = "end of CoverTab[179539]"
}

// NewRandomFromReader returns a UUID based on bytes read from a given io.Reader.
func NewRandomFromReader(r io.Reader) (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:47
	_go_fuzz_dep_.CoverTab[179542]++
										var uuid UUID
										_, err := io.ReadFull(r, uuid[:])
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:50
		_go_fuzz_dep_.CoverTab[179544]++
											return Nil, err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:51
		// _ = "end of CoverTab[179544]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:52
		_go_fuzz_dep_.CoverTab[179545]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:52
		// _ = "end of CoverTab[179545]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:52
	// _ = "end of CoverTab[179542]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:52
	_go_fuzz_dep_.CoverTab[179543]++
										uuid[6] = (uuid[6] & 0x0f) | 0x40
										uuid[8] = (uuid[8] & 0x3f) | 0x80
										return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:55
	// _ = "end of CoverTab[179543]"
}

func newRandomFromPool() (UUID, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:58
	_go_fuzz_dep_.CoverTab[179546]++
										var uuid UUID
										poolMu.Lock()
										if poolPos == randPoolSize {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:61
		_go_fuzz_dep_.CoverTab[179548]++
											_, err := io.ReadFull(rander, pool[:])
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:63
			_go_fuzz_dep_.CoverTab[179550]++
												poolMu.Unlock()
												return Nil, err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:65
			// _ = "end of CoverTab[179550]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:66
			_go_fuzz_dep_.CoverTab[179551]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:66
			// _ = "end of CoverTab[179551]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:66
		// _ = "end of CoverTab[179548]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:66
		_go_fuzz_dep_.CoverTab[179549]++
											poolPos = 0
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:67
		// _ = "end of CoverTab[179549]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:68
		_go_fuzz_dep_.CoverTab[179552]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:68
		// _ = "end of CoverTab[179552]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:68
	// _ = "end of CoverTab[179546]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:68
	_go_fuzz_dep_.CoverTab[179547]++
										copy(uuid[:], pool[poolPos:(poolPos+16)])
										poolPos += 16
										poolMu.Unlock()

										uuid[6] = (uuid[6] & 0x0f) | 0x40
										uuid[8] = (uuid[8] & 0x3f) | 0x80
										return uuid, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:75
	// _ = "end of CoverTab[179547]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:76
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/version4.go:76
var _ = _go_fuzz_dep_.CoverTab
