// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
package sortkeys

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:29
)

import (
	"sort"
)

func Strings(l []string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:35
	_go_fuzz_dep_.CoverTab[134193]++
												sort.Strings(l)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:36
	// _ = "end of CoverTab[134193]"
}

func Float64s(l []float64) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:39
	_go_fuzz_dep_.CoverTab[134194]++
												sort.Float64s(l)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:40
	// _ = "end of CoverTab[134194]"
}

func Float32s(l []float32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:43
	_go_fuzz_dep_.CoverTab[134195]++
												sort.Sort(Float32Slice(l))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:44
	// _ = "end of CoverTab[134195]"
}

func Int64s(l []int64) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:47
	_go_fuzz_dep_.CoverTab[134196]++
												sort.Sort(Int64Slice(l))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:48
	// _ = "end of CoverTab[134196]"
}

func Int32s(l []int32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:51
	_go_fuzz_dep_.CoverTab[134197]++
												sort.Sort(Int32Slice(l))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:52
	// _ = "end of CoverTab[134197]"
}

func Uint64s(l []uint64) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:55
	_go_fuzz_dep_.CoverTab[134198]++
												sort.Sort(Uint64Slice(l))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:56
	// _ = "end of CoverTab[134198]"
}

func Uint32s(l []uint32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:59
	_go_fuzz_dep_.CoverTab[134199]++
												sort.Sort(Uint32Slice(l))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:60
	// _ = "end of CoverTab[134199]"
}

func Bools(l []bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:63
	_go_fuzz_dep_.CoverTab[134200]++
												sort.Sort(BoolSlice(l))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:64
	// _ = "end of CoverTab[134200]"
}

type BoolSlice []bool

func (p BoolSlice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:69
	_go_fuzz_dep_.CoverTab[134201]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:69
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:69
	// _ = "end of CoverTab[134201]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:69
}
func (p BoolSlice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:70
	_go_fuzz_dep_.CoverTab[134202]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:70
	return p[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:70
	// _ = "end of CoverTab[134202]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:70
}
func (p BoolSlice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:71
	_go_fuzz_dep_.CoverTab[134203]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:71
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:71
	// _ = "end of CoverTab[134203]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:71
}

type Int64Slice []int64

func (p Int64Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:75
	_go_fuzz_dep_.CoverTab[134204]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:75
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:75
	// _ = "end of CoverTab[134204]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:75
}
func (p Int64Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:76
	_go_fuzz_dep_.CoverTab[134205]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:76
	return p[i] < p[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:76
	// _ = "end of CoverTab[134205]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:76
}
func (p Int64Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:77
	_go_fuzz_dep_.CoverTab[134206]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:77
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:77
	// _ = "end of CoverTab[134206]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:77
}

type Int32Slice []int32

func (p Int32Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:81
	_go_fuzz_dep_.CoverTab[134207]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:81
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:81
	// _ = "end of CoverTab[134207]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:81
}
func (p Int32Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:82
	_go_fuzz_dep_.CoverTab[134208]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:82
	return p[i] < p[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:82
	// _ = "end of CoverTab[134208]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:82
}
func (p Int32Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:83
	_go_fuzz_dep_.CoverTab[134209]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:83
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:83
	// _ = "end of CoverTab[134209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:83
}

type Uint64Slice []uint64

func (p Uint64Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:87
	_go_fuzz_dep_.CoverTab[134210]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:87
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:87
	// _ = "end of CoverTab[134210]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:87
}
func (p Uint64Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:88
	_go_fuzz_dep_.CoverTab[134211]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:88
	return p[i] < p[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:88
	// _ = "end of CoverTab[134211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:88
}
func (p Uint64Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:89
	_go_fuzz_dep_.CoverTab[134212]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:89
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:89
	// _ = "end of CoverTab[134212]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:89
}

type Uint32Slice []uint32

func (p Uint32Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:93
	_go_fuzz_dep_.CoverTab[134213]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:93
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:93
	// _ = "end of CoverTab[134213]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:93
}
func (p Uint32Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:94
	_go_fuzz_dep_.CoverTab[134214]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:94
	return p[i] < p[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:94
	// _ = "end of CoverTab[134214]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:94
}
func (p Uint32Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:95
	_go_fuzz_dep_.CoverTab[134215]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:95
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:95
	// _ = "end of CoverTab[134215]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:95
}

type Float32Slice []float32

func (p Float32Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:99
	_go_fuzz_dep_.CoverTab[134216]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:99
	return len(p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:99
	// _ = "end of CoverTab[134216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:99
}
func (p Float32Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:100
	_go_fuzz_dep_.CoverTab[134217]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:100
	return p[i] < p[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:100
	// _ = "end of CoverTab[134217]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:100
}
func (p Float32Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:101
	_go_fuzz_dep_.CoverTab[134218]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:101
	p[i], p[j] = p[j], p[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:101
	// _ = "end of CoverTab[134218]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:101
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:101
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/sortkeys/sortkeys.go:101
var _ = _go_fuzz_dep_.CoverTab
