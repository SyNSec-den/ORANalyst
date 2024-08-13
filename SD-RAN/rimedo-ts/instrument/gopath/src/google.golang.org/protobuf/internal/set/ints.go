// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:5
// Package set provides simple set data structures for uint64s.
package set

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:6
)

import "math/bits"

// int64s represents a set of integers within the range of 0..63.
type int64s uint64

func (bs *int64s) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:13
	_go_fuzz_dep_.CoverTab[50181]++
												return bits.OnesCount64(uint64(*bs))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:14
	// _ = "end of CoverTab[50181]"
}
func (bs *int64s) Has(n uint64) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:16
	_go_fuzz_dep_.CoverTab[50182]++
												return uint64(*bs)&(uint64(1)<<n) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:17
	// _ = "end of CoverTab[50182]"
}
func (bs *int64s) Set(n uint64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:19
	_go_fuzz_dep_.CoverTab[50183]++
												*(*uint64)(bs) |= uint64(1) << n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:20
	// _ = "end of CoverTab[50183]"
}
func (bs *int64s) Clear(n uint64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:22
	_go_fuzz_dep_.CoverTab[50184]++
												*(*uint64)(bs) &^= uint64(1) << n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:23
	// _ = "end of CoverTab[50184]"
}

// Ints represents a set of integers within the range of 0..math.MaxUint64.
type Ints struct {
	lo	int64s
	hi	map[uint64]struct{}
}

func (bs *Ints) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:32
	_go_fuzz_dep_.CoverTab[50185]++
												return bs.lo.Len() + len(bs.hi)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:33
	// _ = "end of CoverTab[50185]"
}
func (bs *Ints) Has(n uint64) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:35
	_go_fuzz_dep_.CoverTab[50186]++
												if n < 64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:36
		_go_fuzz_dep_.CoverTab[50188]++
													return bs.lo.Has(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:37
		// _ = "end of CoverTab[50188]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:38
		_go_fuzz_dep_.CoverTab[50189]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:38
		// _ = "end of CoverTab[50189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:38
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:38
	// _ = "end of CoverTab[50186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:38
	_go_fuzz_dep_.CoverTab[50187]++
												_, ok := bs.hi[n]
												return ok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:40
	// _ = "end of CoverTab[50187]"
}
func (bs *Ints) Set(n uint64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:42
	_go_fuzz_dep_.CoverTab[50190]++
												if n < 64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:43
		_go_fuzz_dep_.CoverTab[50193]++
													bs.lo.Set(n)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:45
		// _ = "end of CoverTab[50193]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:46
		_go_fuzz_dep_.CoverTab[50194]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:46
		// _ = "end of CoverTab[50194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:46
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:46
	// _ = "end of CoverTab[50190]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:46
	_go_fuzz_dep_.CoverTab[50191]++
												if bs.hi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:47
		_go_fuzz_dep_.CoverTab[50195]++
													bs.hi = make(map[uint64]struct{})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:48
		// _ = "end of CoverTab[50195]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:49
		_go_fuzz_dep_.CoverTab[50196]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:49
		// _ = "end of CoverTab[50196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:49
	// _ = "end of CoverTab[50191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:49
	_go_fuzz_dep_.CoverTab[50192]++
												bs.hi[n] = struct{}{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:50
	// _ = "end of CoverTab[50192]"
}
func (bs *Ints) Clear(n uint64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:52
	_go_fuzz_dep_.CoverTab[50197]++
												if n < 64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:53
		_go_fuzz_dep_.CoverTab[50199]++
													bs.lo.Clear(n)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:55
		// _ = "end of CoverTab[50199]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:56
		_go_fuzz_dep_.CoverTab[50200]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:56
		// _ = "end of CoverTab[50200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:56
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:56
	// _ = "end of CoverTab[50197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:56
	_go_fuzz_dep_.CoverTab[50198]++
												delete(bs.hi, n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:57
	// _ = "end of CoverTab[50198]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:58
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/set/ints.go:58
var _ = _go_fuzz_dep_.CoverTab
