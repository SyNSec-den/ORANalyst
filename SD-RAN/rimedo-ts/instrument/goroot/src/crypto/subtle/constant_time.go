// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/subtle/constant_time.go:5
// Package subtle implements functions that are often useful in cryptographic
//line /usr/local/go/src/crypto/subtle/constant_time.go:5
// code but require careful thought to use correctly.
//line /usr/local/go/src/crypto/subtle/constant_time.go:7
package subtle

//line /usr/local/go/src/crypto/subtle/constant_time.go:7
import (
//line /usr/local/go/src/crypto/subtle/constant_time.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/subtle/constant_time.go:7
)
//line /usr/local/go/src/crypto/subtle/constant_time.go:7
import (
//line /usr/local/go/src/crypto/subtle/constant_time.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/subtle/constant_time.go:7
)

// ConstantTimeCompare returns 1 if the two slices, x and y, have equal contents
//line /usr/local/go/src/crypto/subtle/constant_time.go:9
// and 0 otherwise. The time taken is a function of the length of the slices and
//line /usr/local/go/src/crypto/subtle/constant_time.go:9
// is independent of the contents. If the lengths of x and y do not match it
//line /usr/local/go/src/crypto/subtle/constant_time.go:9
// returns 0 immediately.
//line /usr/local/go/src/crypto/subtle/constant_time.go:13
func ConstantTimeCompare(x, y []byte) int {
//line /usr/local/go/src/crypto/subtle/constant_time.go:13
	_go_fuzz_dep_.CoverTab[1145]++
								if len(x) != len(y) {
//line /usr/local/go/src/crypto/subtle/constant_time.go:14
		_go_fuzz_dep_.CoverTab[1148]++
									return 0
//line /usr/local/go/src/crypto/subtle/constant_time.go:15
		// _ = "end of CoverTab[1148]"
	} else {
//line /usr/local/go/src/crypto/subtle/constant_time.go:16
		_go_fuzz_dep_.CoverTab[1149]++
//line /usr/local/go/src/crypto/subtle/constant_time.go:16
		// _ = "end of CoverTab[1149]"
//line /usr/local/go/src/crypto/subtle/constant_time.go:16
	}
//line /usr/local/go/src/crypto/subtle/constant_time.go:16
	// _ = "end of CoverTab[1145]"
//line /usr/local/go/src/crypto/subtle/constant_time.go:16
	_go_fuzz_dep_.CoverTab[1146]++

								var v byte

								for i := 0; i < len(x); i++ {
//line /usr/local/go/src/crypto/subtle/constant_time.go:20
		_go_fuzz_dep_.CoverTab[1150]++
									v |= x[i] ^ y[i]
//line /usr/local/go/src/crypto/subtle/constant_time.go:21
		// _ = "end of CoverTab[1150]"
	}
//line /usr/local/go/src/crypto/subtle/constant_time.go:22
	// _ = "end of CoverTab[1146]"
//line /usr/local/go/src/crypto/subtle/constant_time.go:22
	_go_fuzz_dep_.CoverTab[1147]++

								return ConstantTimeByteEq(v, 0)
//line /usr/local/go/src/crypto/subtle/constant_time.go:24
	// _ = "end of CoverTab[1147]"
}

// ConstantTimeSelect returns x if v == 1 and y if v == 0.
//line /usr/local/go/src/crypto/subtle/constant_time.go:27
// Its behavior is undefined if v takes any other value.
//line /usr/local/go/src/crypto/subtle/constant_time.go:29
func ConstantTimeSelect(v, x, y int) int {
//line /usr/local/go/src/crypto/subtle/constant_time.go:29
	_go_fuzz_dep_.CoverTab[1151]++
//line /usr/local/go/src/crypto/subtle/constant_time.go:29
	return ^(v-1)&x | (v-1)&y
//line /usr/local/go/src/crypto/subtle/constant_time.go:29
	// _ = "end of CoverTab[1151]"
//line /usr/local/go/src/crypto/subtle/constant_time.go:29
}

// ConstantTimeByteEq returns 1 if x == y and 0 otherwise.
func ConstantTimeByteEq(x, y uint8) int {
//line /usr/local/go/src/crypto/subtle/constant_time.go:32
	_go_fuzz_dep_.CoverTab[1152]++
								return int((uint32(x^y) - 1) >> 31)
//line /usr/local/go/src/crypto/subtle/constant_time.go:33
	// _ = "end of CoverTab[1152]"
}

// ConstantTimeEq returns 1 if x == y and 0 otherwise.
func ConstantTimeEq(x, y int32) int {
//line /usr/local/go/src/crypto/subtle/constant_time.go:37
	_go_fuzz_dep_.CoverTab[1153]++
								return int((uint64(uint32(x^y)) - 1) >> 63)
//line /usr/local/go/src/crypto/subtle/constant_time.go:38
	// _ = "end of CoverTab[1153]"
}

// ConstantTimeCopy copies the contents of y into x (a slice of equal length)
//line /usr/local/go/src/crypto/subtle/constant_time.go:41
// if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v
//line /usr/local/go/src/crypto/subtle/constant_time.go:41
// takes any other value.
//line /usr/local/go/src/crypto/subtle/constant_time.go:44
func ConstantTimeCopy(v int, x, y []byte) {
//line /usr/local/go/src/crypto/subtle/constant_time.go:44
	_go_fuzz_dep_.CoverTab[1154]++
								if len(x) != len(y) {
//line /usr/local/go/src/crypto/subtle/constant_time.go:45
		_go_fuzz_dep_.CoverTab[1156]++
									panic("subtle: slices have different lengths")
//line /usr/local/go/src/crypto/subtle/constant_time.go:46
		// _ = "end of CoverTab[1156]"
	} else {
//line /usr/local/go/src/crypto/subtle/constant_time.go:47
		_go_fuzz_dep_.CoverTab[1157]++
//line /usr/local/go/src/crypto/subtle/constant_time.go:47
		// _ = "end of CoverTab[1157]"
//line /usr/local/go/src/crypto/subtle/constant_time.go:47
	}
//line /usr/local/go/src/crypto/subtle/constant_time.go:47
	// _ = "end of CoverTab[1154]"
//line /usr/local/go/src/crypto/subtle/constant_time.go:47
	_go_fuzz_dep_.CoverTab[1155]++

								xmask := byte(v - 1)
								ymask := byte(^(v - 1))
								for i := 0; i < len(x); i++ {
//line /usr/local/go/src/crypto/subtle/constant_time.go:51
		_go_fuzz_dep_.CoverTab[1158]++
									x[i] = x[i]&xmask | y[i]&ymask
//line /usr/local/go/src/crypto/subtle/constant_time.go:52
		// _ = "end of CoverTab[1158]"
	}
//line /usr/local/go/src/crypto/subtle/constant_time.go:53
	// _ = "end of CoverTab[1155]"
}

// ConstantTimeLessOrEq returns 1 if x <= y and 0 otherwise.
//line /usr/local/go/src/crypto/subtle/constant_time.go:56
// Its behavior is undefined if x or y are negative or > 2**31 - 1.
//line /usr/local/go/src/crypto/subtle/constant_time.go:58
func ConstantTimeLessOrEq(x, y int) int {
//line /usr/local/go/src/crypto/subtle/constant_time.go:58
	_go_fuzz_dep_.CoverTab[1159]++
								x32 := int32(x)
								y32 := int32(y)
								return int(((x32 - y32 - 1) >> 31) & 1)
//line /usr/local/go/src/crypto/subtle/constant_time.go:61
	// _ = "end of CoverTab[1159]"
}

//line /usr/local/go/src/crypto/subtle/constant_time.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/subtle/constant_time.go:62
var _ = _go_fuzz_dep_.CoverTab
