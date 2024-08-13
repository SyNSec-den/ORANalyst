// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/hash/fnv/fnv.go:5
// Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions
//line /usr/local/go/src/hash/fnv/fnv.go:5
// created by Glenn Fowler, Landon Curt Noll, and Phong Vo.
//line /usr/local/go/src/hash/fnv/fnv.go:5
// See
//line /usr/local/go/src/hash/fnv/fnv.go:5
// https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.
//line /usr/local/go/src/hash/fnv/fnv.go:5
//
//line /usr/local/go/src/hash/fnv/fnv.go:5
// All the hash.Hash implementations returned by this package also
//line /usr/local/go/src/hash/fnv/fnv.go:5
// implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
//line /usr/local/go/src/hash/fnv/fnv.go:5
// marshal and unmarshal the internal state of the hash.
//line /usr/local/go/src/hash/fnv/fnv.go:13
package fnv

//line /usr/local/go/src/hash/fnv/fnv.go:13
import (
//line /usr/local/go/src/hash/fnv/fnv.go:13
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/hash/fnv/fnv.go:13
)
//line /usr/local/go/src/hash/fnv/fnv.go:13
import (
//line /usr/local/go/src/hash/fnv/fnv.go:13
	_atomic_ "sync/atomic"
//line /usr/local/go/src/hash/fnv/fnv.go:13
)

import (
	"errors"
	"hash"
	"math/bits"
)

type (
	sum32	uint32
	sum32a	uint32
	sum64	uint64
	sum64a	uint64
	sum128	[2]uint64
	sum128a	[2]uint64
)

const (
	offset32	= 2166136261
	offset64	= 14695981039346656037
	offset128Lower	= 0x62b821756295c58d
	offset128Higher	= 0x6c62272e07bb0142
	prime32		= 16777619
	prime64		= 1099511628211
	prime128Lower	= 0x13b
	prime128Shift	= 24
)

// New32 returns a new 32-bit FNV-1 hash.Hash.
//line /usr/local/go/src/hash/fnv/fnv.go:41
// Its Sum method will lay the value out in big-endian byte order.
//line /usr/local/go/src/hash/fnv/fnv.go:43
func New32() hash.Hash32 {
//line /usr/local/go/src/hash/fnv/fnv.go:43
	_go_fuzz_dep_.CoverTab[48178]++
						var s sum32 = offset32
						return &s
//line /usr/local/go/src/hash/fnv/fnv.go:45
	// _ = "end of CoverTab[48178]"
}

// New32a returns a new 32-bit FNV-1a hash.Hash.
//line /usr/local/go/src/hash/fnv/fnv.go:48
// Its Sum method will lay the value out in big-endian byte order.
//line /usr/local/go/src/hash/fnv/fnv.go:50
func New32a() hash.Hash32 {
//line /usr/local/go/src/hash/fnv/fnv.go:50
	_go_fuzz_dep_.CoverTab[48179]++
						var s sum32a = offset32
						return &s
//line /usr/local/go/src/hash/fnv/fnv.go:52
	// _ = "end of CoverTab[48179]"
}

// New64 returns a new 64-bit FNV-1 hash.Hash.
//line /usr/local/go/src/hash/fnv/fnv.go:55
// Its Sum method will lay the value out in big-endian byte order.
//line /usr/local/go/src/hash/fnv/fnv.go:57
func New64() hash.Hash64 {
//line /usr/local/go/src/hash/fnv/fnv.go:57
	_go_fuzz_dep_.CoverTab[48180]++
						var s sum64 = offset64
						return &s
//line /usr/local/go/src/hash/fnv/fnv.go:59
	// _ = "end of CoverTab[48180]"
}

// New64a returns a new 64-bit FNV-1a hash.Hash.
//line /usr/local/go/src/hash/fnv/fnv.go:62
// Its Sum method will lay the value out in big-endian byte order.
//line /usr/local/go/src/hash/fnv/fnv.go:64
func New64a() hash.Hash64 {
//line /usr/local/go/src/hash/fnv/fnv.go:64
	_go_fuzz_dep_.CoverTab[48181]++
						var s sum64a = offset64
						return &s
//line /usr/local/go/src/hash/fnv/fnv.go:66
	// _ = "end of CoverTab[48181]"
}

// New128 returns a new 128-bit FNV-1 hash.Hash.
//line /usr/local/go/src/hash/fnv/fnv.go:69
// Its Sum method will lay the value out in big-endian byte order.
//line /usr/local/go/src/hash/fnv/fnv.go:71
func New128() hash.Hash {
//line /usr/local/go/src/hash/fnv/fnv.go:71
	_go_fuzz_dep_.CoverTab[48182]++
						var s sum128
						s[0] = offset128Higher
						s[1] = offset128Lower
						return &s
//line /usr/local/go/src/hash/fnv/fnv.go:75
	// _ = "end of CoverTab[48182]"
}

// New128a returns a new 128-bit FNV-1a hash.Hash.
//line /usr/local/go/src/hash/fnv/fnv.go:78
// Its Sum method will lay the value out in big-endian byte order.
//line /usr/local/go/src/hash/fnv/fnv.go:80
func New128a() hash.Hash {
//line /usr/local/go/src/hash/fnv/fnv.go:80
	_go_fuzz_dep_.CoverTab[48183]++
						var s sum128a
						s[0] = offset128Higher
						s[1] = offset128Lower
						return &s
//line /usr/local/go/src/hash/fnv/fnv.go:84
	// _ = "end of CoverTab[48183]"
}

func (s *sum32) Reset() {
//line /usr/local/go/src/hash/fnv/fnv.go:87
	_go_fuzz_dep_.CoverTab[48184]++
//line /usr/local/go/src/hash/fnv/fnv.go:87
	*s = offset32
//line /usr/local/go/src/hash/fnv/fnv.go:87
	// _ = "end of CoverTab[48184]"
//line /usr/local/go/src/hash/fnv/fnv.go:87
}
func (s *sum32a) Reset() {
//line /usr/local/go/src/hash/fnv/fnv.go:88
	_go_fuzz_dep_.CoverTab[48185]++
//line /usr/local/go/src/hash/fnv/fnv.go:88
	*s = offset32
//line /usr/local/go/src/hash/fnv/fnv.go:88
	// _ = "end of CoverTab[48185]"
//line /usr/local/go/src/hash/fnv/fnv.go:88
}
func (s *sum64) Reset() {
//line /usr/local/go/src/hash/fnv/fnv.go:89
	_go_fuzz_dep_.CoverTab[48186]++
//line /usr/local/go/src/hash/fnv/fnv.go:89
	*s = offset64
//line /usr/local/go/src/hash/fnv/fnv.go:89
	// _ = "end of CoverTab[48186]"
//line /usr/local/go/src/hash/fnv/fnv.go:89
}
func (s *sum64a) Reset() {
//line /usr/local/go/src/hash/fnv/fnv.go:90
	_go_fuzz_dep_.CoverTab[48187]++
//line /usr/local/go/src/hash/fnv/fnv.go:90
	*s = offset64
//line /usr/local/go/src/hash/fnv/fnv.go:90
	// _ = "end of CoverTab[48187]"
//line /usr/local/go/src/hash/fnv/fnv.go:90
}
func (s *sum128) Reset() {
//line /usr/local/go/src/hash/fnv/fnv.go:91
	_go_fuzz_dep_.CoverTab[48188]++
//line /usr/local/go/src/hash/fnv/fnv.go:91
	s[0] = offset128Higher
//line /usr/local/go/src/hash/fnv/fnv.go:91
	s[1] = offset128Lower
//line /usr/local/go/src/hash/fnv/fnv.go:91
	// _ = "end of CoverTab[48188]"
//line /usr/local/go/src/hash/fnv/fnv.go:91
}
func (s *sum128a) Reset() {
//line /usr/local/go/src/hash/fnv/fnv.go:92
	_go_fuzz_dep_.CoverTab[48189]++
//line /usr/local/go/src/hash/fnv/fnv.go:92
	s[0] = offset128Higher
//line /usr/local/go/src/hash/fnv/fnv.go:92
	s[1] = offset128Lower
//line /usr/local/go/src/hash/fnv/fnv.go:92
	// _ = "end of CoverTab[48189]"
//line /usr/local/go/src/hash/fnv/fnv.go:92
}

func (s *sum32) Sum32() uint32 {
//line /usr/local/go/src/hash/fnv/fnv.go:94
	_go_fuzz_dep_.CoverTab[48190]++
//line /usr/local/go/src/hash/fnv/fnv.go:94
	return uint32(*s)
//line /usr/local/go/src/hash/fnv/fnv.go:94
	// _ = "end of CoverTab[48190]"
//line /usr/local/go/src/hash/fnv/fnv.go:94
}
func (s *sum32a) Sum32() uint32 {
//line /usr/local/go/src/hash/fnv/fnv.go:95
	_go_fuzz_dep_.CoverTab[48191]++
//line /usr/local/go/src/hash/fnv/fnv.go:95
	return uint32(*s)
//line /usr/local/go/src/hash/fnv/fnv.go:95
	// _ = "end of CoverTab[48191]"
//line /usr/local/go/src/hash/fnv/fnv.go:95
}
func (s *sum64) Sum64() uint64 {
//line /usr/local/go/src/hash/fnv/fnv.go:96
	_go_fuzz_dep_.CoverTab[48192]++
//line /usr/local/go/src/hash/fnv/fnv.go:96
	return uint64(*s)
//line /usr/local/go/src/hash/fnv/fnv.go:96
	// _ = "end of CoverTab[48192]"
//line /usr/local/go/src/hash/fnv/fnv.go:96
}
func (s *sum64a) Sum64() uint64 {
//line /usr/local/go/src/hash/fnv/fnv.go:97
	_go_fuzz_dep_.CoverTab[48193]++
//line /usr/local/go/src/hash/fnv/fnv.go:97
	return uint64(*s)
//line /usr/local/go/src/hash/fnv/fnv.go:97
	// _ = "end of CoverTab[48193]"
//line /usr/local/go/src/hash/fnv/fnv.go:97
}

func (s *sum32) Write(data []byte) (int, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:99
	_go_fuzz_dep_.CoverTab[48194]++
						hash := *s
						for _, c := range data {
//line /usr/local/go/src/hash/fnv/fnv.go:101
		_go_fuzz_dep_.CoverTab[48196]++
							hash *= prime32
							hash ^= sum32(c)
//line /usr/local/go/src/hash/fnv/fnv.go:103
		// _ = "end of CoverTab[48196]"
	}
//line /usr/local/go/src/hash/fnv/fnv.go:104
	// _ = "end of CoverTab[48194]"
//line /usr/local/go/src/hash/fnv/fnv.go:104
	_go_fuzz_dep_.CoverTab[48195]++
						*s = hash
						return len(data), nil
//line /usr/local/go/src/hash/fnv/fnv.go:106
	// _ = "end of CoverTab[48195]"
}

func (s *sum32a) Write(data []byte) (int, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:109
	_go_fuzz_dep_.CoverTab[48197]++
						hash := *s
						for _, c := range data {
//line /usr/local/go/src/hash/fnv/fnv.go:111
		_go_fuzz_dep_.CoverTab[48199]++
							hash ^= sum32a(c)
							hash *= prime32
//line /usr/local/go/src/hash/fnv/fnv.go:113
		// _ = "end of CoverTab[48199]"
	}
//line /usr/local/go/src/hash/fnv/fnv.go:114
	// _ = "end of CoverTab[48197]"
//line /usr/local/go/src/hash/fnv/fnv.go:114
	_go_fuzz_dep_.CoverTab[48198]++
						*s = hash
						return len(data), nil
//line /usr/local/go/src/hash/fnv/fnv.go:116
	// _ = "end of CoverTab[48198]"
}

func (s *sum64) Write(data []byte) (int, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:119
	_go_fuzz_dep_.CoverTab[48200]++
						hash := *s
						for _, c := range data {
//line /usr/local/go/src/hash/fnv/fnv.go:121
		_go_fuzz_dep_.CoverTab[48202]++
							hash *= prime64
							hash ^= sum64(c)
//line /usr/local/go/src/hash/fnv/fnv.go:123
		// _ = "end of CoverTab[48202]"
	}
//line /usr/local/go/src/hash/fnv/fnv.go:124
	// _ = "end of CoverTab[48200]"
//line /usr/local/go/src/hash/fnv/fnv.go:124
	_go_fuzz_dep_.CoverTab[48201]++
						*s = hash
						return len(data), nil
//line /usr/local/go/src/hash/fnv/fnv.go:126
	// _ = "end of CoverTab[48201]"
}

func (s *sum64a) Write(data []byte) (int, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:129
	_go_fuzz_dep_.CoverTab[48203]++
						hash := *s
						for _, c := range data {
//line /usr/local/go/src/hash/fnv/fnv.go:131
		_go_fuzz_dep_.CoverTab[48205]++
							hash ^= sum64a(c)
							hash *= prime64
//line /usr/local/go/src/hash/fnv/fnv.go:133
		// _ = "end of CoverTab[48205]"
	}
//line /usr/local/go/src/hash/fnv/fnv.go:134
	// _ = "end of CoverTab[48203]"
//line /usr/local/go/src/hash/fnv/fnv.go:134
	_go_fuzz_dep_.CoverTab[48204]++
						*s = hash
						return len(data), nil
//line /usr/local/go/src/hash/fnv/fnv.go:136
	// _ = "end of CoverTab[48204]"
}

func (s *sum128) Write(data []byte) (int, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:139
	_go_fuzz_dep_.CoverTab[48206]++
						for _, c := range data {
//line /usr/local/go/src/hash/fnv/fnv.go:140
		_go_fuzz_dep_.CoverTab[48208]++

							s0, s1 := bits.Mul64(prime128Lower, s[1])
							s0 += s[1]<<prime128Shift + prime128Lower*s[0]

							s[1] = s1
							s[0] = s0
							s[1] ^= uint64(c)
//line /usr/local/go/src/hash/fnv/fnv.go:147
		// _ = "end of CoverTab[48208]"
	}
//line /usr/local/go/src/hash/fnv/fnv.go:148
	// _ = "end of CoverTab[48206]"
//line /usr/local/go/src/hash/fnv/fnv.go:148
	_go_fuzz_dep_.CoverTab[48207]++
						return len(data), nil
//line /usr/local/go/src/hash/fnv/fnv.go:149
	// _ = "end of CoverTab[48207]"
}

func (s *sum128a) Write(data []byte) (int, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:152
	_go_fuzz_dep_.CoverTab[48209]++
						for _, c := range data {
//line /usr/local/go/src/hash/fnv/fnv.go:153
		_go_fuzz_dep_.CoverTab[48211]++
							s[1] ^= uint64(c)

							s0, s1 := bits.Mul64(prime128Lower, s[1])
							s0 += s[1]<<prime128Shift + prime128Lower*s[0]

							s[1] = s1
							s[0] = s0
//line /usr/local/go/src/hash/fnv/fnv.go:160
		// _ = "end of CoverTab[48211]"
	}
//line /usr/local/go/src/hash/fnv/fnv.go:161
	// _ = "end of CoverTab[48209]"
//line /usr/local/go/src/hash/fnv/fnv.go:161
	_go_fuzz_dep_.CoverTab[48210]++
						return len(data), nil
//line /usr/local/go/src/hash/fnv/fnv.go:162
	// _ = "end of CoverTab[48210]"
}

func (s *sum32) Size() int {
//line /usr/local/go/src/hash/fnv/fnv.go:165
	_go_fuzz_dep_.CoverTab[48212]++
//line /usr/local/go/src/hash/fnv/fnv.go:165
	return 4
//line /usr/local/go/src/hash/fnv/fnv.go:165
	// _ = "end of CoverTab[48212]"
//line /usr/local/go/src/hash/fnv/fnv.go:165
}
func (s *sum32a) Size() int {
//line /usr/local/go/src/hash/fnv/fnv.go:166
	_go_fuzz_dep_.CoverTab[48213]++
//line /usr/local/go/src/hash/fnv/fnv.go:166
	return 4
//line /usr/local/go/src/hash/fnv/fnv.go:166
	// _ = "end of CoverTab[48213]"
//line /usr/local/go/src/hash/fnv/fnv.go:166
}
func (s *sum64) Size() int {
//line /usr/local/go/src/hash/fnv/fnv.go:167
	_go_fuzz_dep_.CoverTab[48214]++
//line /usr/local/go/src/hash/fnv/fnv.go:167
	return 8
//line /usr/local/go/src/hash/fnv/fnv.go:167
	// _ = "end of CoverTab[48214]"
//line /usr/local/go/src/hash/fnv/fnv.go:167
}
func (s *sum64a) Size() int {
//line /usr/local/go/src/hash/fnv/fnv.go:168
	_go_fuzz_dep_.CoverTab[48215]++
//line /usr/local/go/src/hash/fnv/fnv.go:168
	return 8
//line /usr/local/go/src/hash/fnv/fnv.go:168
	// _ = "end of CoverTab[48215]"
//line /usr/local/go/src/hash/fnv/fnv.go:168
}
func (s *sum128) Size() int {
//line /usr/local/go/src/hash/fnv/fnv.go:169
	_go_fuzz_dep_.CoverTab[48216]++
//line /usr/local/go/src/hash/fnv/fnv.go:169
	return 16
//line /usr/local/go/src/hash/fnv/fnv.go:169
	// _ = "end of CoverTab[48216]"
//line /usr/local/go/src/hash/fnv/fnv.go:169
}
func (s *sum128a) Size() int {
//line /usr/local/go/src/hash/fnv/fnv.go:170
	_go_fuzz_dep_.CoverTab[48217]++
//line /usr/local/go/src/hash/fnv/fnv.go:170
	return 16
//line /usr/local/go/src/hash/fnv/fnv.go:170
	// _ = "end of CoverTab[48217]"
//line /usr/local/go/src/hash/fnv/fnv.go:170
}

func (s *sum32) BlockSize() int {
//line /usr/local/go/src/hash/fnv/fnv.go:172
	_go_fuzz_dep_.CoverTab[48218]++
//line /usr/local/go/src/hash/fnv/fnv.go:172
	return 1
//line /usr/local/go/src/hash/fnv/fnv.go:172
	// _ = "end of CoverTab[48218]"
//line /usr/local/go/src/hash/fnv/fnv.go:172
}
func (s *sum32a) BlockSize() int {
//line /usr/local/go/src/hash/fnv/fnv.go:173
	_go_fuzz_dep_.CoverTab[48219]++
//line /usr/local/go/src/hash/fnv/fnv.go:173
	return 1
//line /usr/local/go/src/hash/fnv/fnv.go:173
	// _ = "end of CoverTab[48219]"
//line /usr/local/go/src/hash/fnv/fnv.go:173
}
func (s *sum64) BlockSize() int {
//line /usr/local/go/src/hash/fnv/fnv.go:174
	_go_fuzz_dep_.CoverTab[48220]++
//line /usr/local/go/src/hash/fnv/fnv.go:174
	return 1
//line /usr/local/go/src/hash/fnv/fnv.go:174
	// _ = "end of CoverTab[48220]"
//line /usr/local/go/src/hash/fnv/fnv.go:174
}
func (s *sum64a) BlockSize() int {
//line /usr/local/go/src/hash/fnv/fnv.go:175
	_go_fuzz_dep_.CoverTab[48221]++
//line /usr/local/go/src/hash/fnv/fnv.go:175
	return 1
//line /usr/local/go/src/hash/fnv/fnv.go:175
	// _ = "end of CoverTab[48221]"
//line /usr/local/go/src/hash/fnv/fnv.go:175
}
func (s *sum128) BlockSize() int {
//line /usr/local/go/src/hash/fnv/fnv.go:176
	_go_fuzz_dep_.CoverTab[48222]++
//line /usr/local/go/src/hash/fnv/fnv.go:176
	return 1
//line /usr/local/go/src/hash/fnv/fnv.go:176
	// _ = "end of CoverTab[48222]"
//line /usr/local/go/src/hash/fnv/fnv.go:176
}
func (s *sum128a) BlockSize() int {
//line /usr/local/go/src/hash/fnv/fnv.go:177
	_go_fuzz_dep_.CoverTab[48223]++
//line /usr/local/go/src/hash/fnv/fnv.go:177
	return 1
//line /usr/local/go/src/hash/fnv/fnv.go:177
	// _ = "end of CoverTab[48223]"
//line /usr/local/go/src/hash/fnv/fnv.go:177
}

func (s *sum32) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:179
	_go_fuzz_dep_.CoverTab[48224]++
						v := uint32(*s)
						return append(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/hash/fnv/fnv.go:181
	// _ = "end of CoverTab[48224]"
}

func (s *sum32a) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:184
	_go_fuzz_dep_.CoverTab[48225]++
						v := uint32(*s)
						return append(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/hash/fnv/fnv.go:186
	// _ = "end of CoverTab[48225]"
}

func (s *sum64) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:189
	_go_fuzz_dep_.CoverTab[48226]++
						v := uint64(*s)
						return append(in, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/hash/fnv/fnv.go:191
	// _ = "end of CoverTab[48226]"
}

func (s *sum64a) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:194
	_go_fuzz_dep_.CoverTab[48227]++
						v := uint64(*s)
						return append(in, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /usr/local/go/src/hash/fnv/fnv.go:196
	// _ = "end of CoverTab[48227]"
}

func (s *sum128) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:199
	_go_fuzz_dep_.CoverTab[48228]++
						return append(in,
		byte(s[0]>>56), byte(s[0]>>48), byte(s[0]>>40), byte(s[0]>>32), byte(s[0]>>24), byte(s[0]>>16), byte(s[0]>>8), byte(s[0]),
		byte(s[1]>>56), byte(s[1]>>48), byte(s[1]>>40), byte(s[1]>>32), byte(s[1]>>24), byte(s[1]>>16), byte(s[1]>>8), byte(s[1]),
	)
//line /usr/local/go/src/hash/fnv/fnv.go:203
	// _ = "end of CoverTab[48228]"
}

func (s *sum128a) Sum(in []byte) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:206
	_go_fuzz_dep_.CoverTab[48229]++
						return append(in,
		byte(s[0]>>56), byte(s[0]>>48), byte(s[0]>>40), byte(s[0]>>32), byte(s[0]>>24), byte(s[0]>>16), byte(s[0]>>8), byte(s[0]),
		byte(s[1]>>56), byte(s[1]>>48), byte(s[1]>>40), byte(s[1]>>32), byte(s[1]>>24), byte(s[1]>>16), byte(s[1]>>8), byte(s[1]),
	)
//line /usr/local/go/src/hash/fnv/fnv.go:210
	// _ = "end of CoverTab[48229]"
}

const (
	magic32			= "fnv\x01"
	magic32a		= "fnv\x02"
	magic64			= "fnv\x03"
	magic64a		= "fnv\x04"
	magic128		= "fnv\x05"
	magic128a		= "fnv\x06"
	marshaledSize32		= len(magic32) + 4
	marshaledSize64		= len(magic64) + 8
	marshaledSize128	= len(magic128) + 8*2
)

func (s *sum32) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:225
	_go_fuzz_dep_.CoverTab[48230]++
						b := make([]byte, 0, marshaledSize32)
						b = append(b, magic32...)
						b = appendUint32(b, uint32(*s))
						return b, nil
//line /usr/local/go/src/hash/fnv/fnv.go:229
	// _ = "end of CoverTab[48230]"
}

func (s *sum32a) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:232
	_go_fuzz_dep_.CoverTab[48231]++
						b := make([]byte, 0, marshaledSize32)
						b = append(b, magic32a...)
						b = appendUint32(b, uint32(*s))
						return b, nil
//line /usr/local/go/src/hash/fnv/fnv.go:236
	// _ = "end of CoverTab[48231]"
}

func (s *sum64) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:239
	_go_fuzz_dep_.CoverTab[48232]++
						b := make([]byte, 0, marshaledSize64)
						b = append(b, magic64...)
						b = appendUint64(b, uint64(*s))
						return b, nil
//line /usr/local/go/src/hash/fnv/fnv.go:243
	// _ = "end of CoverTab[48232]"

}

func (s *sum64a) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:247
	_go_fuzz_dep_.CoverTab[48233]++
						b := make([]byte, 0, marshaledSize64)
						b = append(b, magic64a...)
						b = appendUint64(b, uint64(*s))
						return b, nil
//line /usr/local/go/src/hash/fnv/fnv.go:251
	// _ = "end of CoverTab[48233]"
}

func (s *sum128) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:254
	_go_fuzz_dep_.CoverTab[48234]++
						b := make([]byte, 0, marshaledSize128)
						b = append(b, magic128...)
						b = appendUint64(b, s[0])
						b = appendUint64(b, s[1])
						return b, nil
//line /usr/local/go/src/hash/fnv/fnv.go:259
	// _ = "end of CoverTab[48234]"
}

func (s *sum128a) MarshalBinary() ([]byte, error) {
//line /usr/local/go/src/hash/fnv/fnv.go:262
	_go_fuzz_dep_.CoverTab[48235]++
						b := make([]byte, 0, marshaledSize128)
						b = append(b, magic128a...)
						b = appendUint64(b, s[0])
						b = appendUint64(b, s[1])
						return b, nil
//line /usr/local/go/src/hash/fnv/fnv.go:267
	// _ = "end of CoverTab[48235]"
}

func (s *sum32) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/fnv/fnv.go:270
	_go_fuzz_dep_.CoverTab[48236]++
						if len(b) < len(magic32) || func() bool {
//line /usr/local/go/src/hash/fnv/fnv.go:271
		_go_fuzz_dep_.CoverTab[48239]++
//line /usr/local/go/src/hash/fnv/fnv.go:271
		return string(b[:len(magic32)]) != magic32
//line /usr/local/go/src/hash/fnv/fnv.go:271
		// _ = "end of CoverTab[48239]"
//line /usr/local/go/src/hash/fnv/fnv.go:271
	}() {
//line /usr/local/go/src/hash/fnv/fnv.go:271
		_go_fuzz_dep_.CoverTab[48240]++
							return errors.New("hash/fnv: invalid hash state identifier")
//line /usr/local/go/src/hash/fnv/fnv.go:272
		// _ = "end of CoverTab[48240]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:273
		_go_fuzz_dep_.CoverTab[48241]++
//line /usr/local/go/src/hash/fnv/fnv.go:273
		// _ = "end of CoverTab[48241]"
//line /usr/local/go/src/hash/fnv/fnv.go:273
	}
//line /usr/local/go/src/hash/fnv/fnv.go:273
	// _ = "end of CoverTab[48236]"
//line /usr/local/go/src/hash/fnv/fnv.go:273
	_go_fuzz_dep_.CoverTab[48237]++
						if len(b) != marshaledSize32 {
//line /usr/local/go/src/hash/fnv/fnv.go:274
		_go_fuzz_dep_.CoverTab[48242]++
							return errors.New("hash/fnv: invalid hash state size")
//line /usr/local/go/src/hash/fnv/fnv.go:275
		// _ = "end of CoverTab[48242]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:276
		_go_fuzz_dep_.CoverTab[48243]++
//line /usr/local/go/src/hash/fnv/fnv.go:276
		// _ = "end of CoverTab[48243]"
//line /usr/local/go/src/hash/fnv/fnv.go:276
	}
//line /usr/local/go/src/hash/fnv/fnv.go:276
	// _ = "end of CoverTab[48237]"
//line /usr/local/go/src/hash/fnv/fnv.go:276
	_go_fuzz_dep_.CoverTab[48238]++
						*s = sum32(readUint32(b[4:]))
						return nil
//line /usr/local/go/src/hash/fnv/fnv.go:278
	// _ = "end of CoverTab[48238]"
}

func (s *sum32a) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/fnv/fnv.go:281
	_go_fuzz_dep_.CoverTab[48244]++
						if len(b) < len(magic32a) || func() bool {
//line /usr/local/go/src/hash/fnv/fnv.go:282
		_go_fuzz_dep_.CoverTab[48247]++
//line /usr/local/go/src/hash/fnv/fnv.go:282
		return string(b[:len(magic32a)]) != magic32a
//line /usr/local/go/src/hash/fnv/fnv.go:282
		// _ = "end of CoverTab[48247]"
//line /usr/local/go/src/hash/fnv/fnv.go:282
	}() {
//line /usr/local/go/src/hash/fnv/fnv.go:282
		_go_fuzz_dep_.CoverTab[48248]++
							return errors.New("hash/fnv: invalid hash state identifier")
//line /usr/local/go/src/hash/fnv/fnv.go:283
		// _ = "end of CoverTab[48248]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:284
		_go_fuzz_dep_.CoverTab[48249]++
//line /usr/local/go/src/hash/fnv/fnv.go:284
		// _ = "end of CoverTab[48249]"
//line /usr/local/go/src/hash/fnv/fnv.go:284
	}
//line /usr/local/go/src/hash/fnv/fnv.go:284
	// _ = "end of CoverTab[48244]"
//line /usr/local/go/src/hash/fnv/fnv.go:284
	_go_fuzz_dep_.CoverTab[48245]++
						if len(b) != marshaledSize32 {
//line /usr/local/go/src/hash/fnv/fnv.go:285
		_go_fuzz_dep_.CoverTab[48250]++
							return errors.New("hash/fnv: invalid hash state size")
//line /usr/local/go/src/hash/fnv/fnv.go:286
		// _ = "end of CoverTab[48250]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:287
		_go_fuzz_dep_.CoverTab[48251]++
//line /usr/local/go/src/hash/fnv/fnv.go:287
		// _ = "end of CoverTab[48251]"
//line /usr/local/go/src/hash/fnv/fnv.go:287
	}
//line /usr/local/go/src/hash/fnv/fnv.go:287
	// _ = "end of CoverTab[48245]"
//line /usr/local/go/src/hash/fnv/fnv.go:287
	_go_fuzz_dep_.CoverTab[48246]++
						*s = sum32a(readUint32(b[4:]))
						return nil
//line /usr/local/go/src/hash/fnv/fnv.go:289
	// _ = "end of CoverTab[48246]"
}

func (s *sum64) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/fnv/fnv.go:292
	_go_fuzz_dep_.CoverTab[48252]++
						if len(b) < len(magic64) || func() bool {
//line /usr/local/go/src/hash/fnv/fnv.go:293
		_go_fuzz_dep_.CoverTab[48255]++
//line /usr/local/go/src/hash/fnv/fnv.go:293
		return string(b[:len(magic64)]) != magic64
//line /usr/local/go/src/hash/fnv/fnv.go:293
		// _ = "end of CoverTab[48255]"
//line /usr/local/go/src/hash/fnv/fnv.go:293
	}() {
//line /usr/local/go/src/hash/fnv/fnv.go:293
		_go_fuzz_dep_.CoverTab[48256]++
							return errors.New("hash/fnv: invalid hash state identifier")
//line /usr/local/go/src/hash/fnv/fnv.go:294
		// _ = "end of CoverTab[48256]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:295
		_go_fuzz_dep_.CoverTab[48257]++
//line /usr/local/go/src/hash/fnv/fnv.go:295
		// _ = "end of CoverTab[48257]"
//line /usr/local/go/src/hash/fnv/fnv.go:295
	}
//line /usr/local/go/src/hash/fnv/fnv.go:295
	// _ = "end of CoverTab[48252]"
//line /usr/local/go/src/hash/fnv/fnv.go:295
	_go_fuzz_dep_.CoverTab[48253]++
						if len(b) != marshaledSize64 {
//line /usr/local/go/src/hash/fnv/fnv.go:296
		_go_fuzz_dep_.CoverTab[48258]++
							return errors.New("hash/fnv: invalid hash state size")
//line /usr/local/go/src/hash/fnv/fnv.go:297
		// _ = "end of CoverTab[48258]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:298
		_go_fuzz_dep_.CoverTab[48259]++
//line /usr/local/go/src/hash/fnv/fnv.go:298
		// _ = "end of CoverTab[48259]"
//line /usr/local/go/src/hash/fnv/fnv.go:298
	}
//line /usr/local/go/src/hash/fnv/fnv.go:298
	// _ = "end of CoverTab[48253]"
//line /usr/local/go/src/hash/fnv/fnv.go:298
	_go_fuzz_dep_.CoverTab[48254]++
						*s = sum64(readUint64(b[4:]))
						return nil
//line /usr/local/go/src/hash/fnv/fnv.go:300
	// _ = "end of CoverTab[48254]"
}

func (s *sum64a) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/fnv/fnv.go:303
	_go_fuzz_dep_.CoverTab[48260]++
						if len(b) < len(magic64a) || func() bool {
//line /usr/local/go/src/hash/fnv/fnv.go:304
		_go_fuzz_dep_.CoverTab[48263]++
//line /usr/local/go/src/hash/fnv/fnv.go:304
		return string(b[:len(magic64a)]) != magic64a
//line /usr/local/go/src/hash/fnv/fnv.go:304
		// _ = "end of CoverTab[48263]"
//line /usr/local/go/src/hash/fnv/fnv.go:304
	}() {
//line /usr/local/go/src/hash/fnv/fnv.go:304
		_go_fuzz_dep_.CoverTab[48264]++
							return errors.New("hash/fnv: invalid hash state identifier")
//line /usr/local/go/src/hash/fnv/fnv.go:305
		// _ = "end of CoverTab[48264]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:306
		_go_fuzz_dep_.CoverTab[48265]++
//line /usr/local/go/src/hash/fnv/fnv.go:306
		// _ = "end of CoverTab[48265]"
//line /usr/local/go/src/hash/fnv/fnv.go:306
	}
//line /usr/local/go/src/hash/fnv/fnv.go:306
	// _ = "end of CoverTab[48260]"
//line /usr/local/go/src/hash/fnv/fnv.go:306
	_go_fuzz_dep_.CoverTab[48261]++
						if len(b) != marshaledSize64 {
//line /usr/local/go/src/hash/fnv/fnv.go:307
		_go_fuzz_dep_.CoverTab[48266]++
							return errors.New("hash/fnv: invalid hash state size")
//line /usr/local/go/src/hash/fnv/fnv.go:308
		// _ = "end of CoverTab[48266]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:309
		_go_fuzz_dep_.CoverTab[48267]++
//line /usr/local/go/src/hash/fnv/fnv.go:309
		// _ = "end of CoverTab[48267]"
//line /usr/local/go/src/hash/fnv/fnv.go:309
	}
//line /usr/local/go/src/hash/fnv/fnv.go:309
	// _ = "end of CoverTab[48261]"
//line /usr/local/go/src/hash/fnv/fnv.go:309
	_go_fuzz_dep_.CoverTab[48262]++
						*s = sum64a(readUint64(b[4:]))
						return nil
//line /usr/local/go/src/hash/fnv/fnv.go:311
	// _ = "end of CoverTab[48262]"
}

func (s *sum128) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/fnv/fnv.go:314
	_go_fuzz_dep_.CoverTab[48268]++
						if len(b) < len(magic128) || func() bool {
//line /usr/local/go/src/hash/fnv/fnv.go:315
		_go_fuzz_dep_.CoverTab[48271]++
//line /usr/local/go/src/hash/fnv/fnv.go:315
		return string(b[:len(magic128)]) != magic128
//line /usr/local/go/src/hash/fnv/fnv.go:315
		// _ = "end of CoverTab[48271]"
//line /usr/local/go/src/hash/fnv/fnv.go:315
	}() {
//line /usr/local/go/src/hash/fnv/fnv.go:315
		_go_fuzz_dep_.CoverTab[48272]++
							return errors.New("hash/fnv: invalid hash state identifier")
//line /usr/local/go/src/hash/fnv/fnv.go:316
		// _ = "end of CoverTab[48272]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:317
		_go_fuzz_dep_.CoverTab[48273]++
//line /usr/local/go/src/hash/fnv/fnv.go:317
		// _ = "end of CoverTab[48273]"
//line /usr/local/go/src/hash/fnv/fnv.go:317
	}
//line /usr/local/go/src/hash/fnv/fnv.go:317
	// _ = "end of CoverTab[48268]"
//line /usr/local/go/src/hash/fnv/fnv.go:317
	_go_fuzz_dep_.CoverTab[48269]++
						if len(b) != marshaledSize128 {
//line /usr/local/go/src/hash/fnv/fnv.go:318
		_go_fuzz_dep_.CoverTab[48274]++
							return errors.New("hash/fnv: invalid hash state size")
//line /usr/local/go/src/hash/fnv/fnv.go:319
		// _ = "end of CoverTab[48274]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:320
		_go_fuzz_dep_.CoverTab[48275]++
//line /usr/local/go/src/hash/fnv/fnv.go:320
		// _ = "end of CoverTab[48275]"
//line /usr/local/go/src/hash/fnv/fnv.go:320
	}
//line /usr/local/go/src/hash/fnv/fnv.go:320
	// _ = "end of CoverTab[48269]"
//line /usr/local/go/src/hash/fnv/fnv.go:320
	_go_fuzz_dep_.CoverTab[48270]++
						s[0] = readUint64(b[4:])
						s[1] = readUint64(b[12:])
						return nil
//line /usr/local/go/src/hash/fnv/fnv.go:323
	// _ = "end of CoverTab[48270]"
}

func (s *sum128a) UnmarshalBinary(b []byte) error {
//line /usr/local/go/src/hash/fnv/fnv.go:326
	_go_fuzz_dep_.CoverTab[48276]++
						if len(b) < len(magic128a) || func() bool {
//line /usr/local/go/src/hash/fnv/fnv.go:327
		_go_fuzz_dep_.CoverTab[48279]++
//line /usr/local/go/src/hash/fnv/fnv.go:327
		return string(b[:len(magic128a)]) != magic128a
//line /usr/local/go/src/hash/fnv/fnv.go:327
		// _ = "end of CoverTab[48279]"
//line /usr/local/go/src/hash/fnv/fnv.go:327
	}() {
//line /usr/local/go/src/hash/fnv/fnv.go:327
		_go_fuzz_dep_.CoverTab[48280]++
							return errors.New("hash/fnv: invalid hash state identifier")
//line /usr/local/go/src/hash/fnv/fnv.go:328
		// _ = "end of CoverTab[48280]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:329
		_go_fuzz_dep_.CoverTab[48281]++
//line /usr/local/go/src/hash/fnv/fnv.go:329
		// _ = "end of CoverTab[48281]"
//line /usr/local/go/src/hash/fnv/fnv.go:329
	}
//line /usr/local/go/src/hash/fnv/fnv.go:329
	// _ = "end of CoverTab[48276]"
//line /usr/local/go/src/hash/fnv/fnv.go:329
	_go_fuzz_dep_.CoverTab[48277]++
						if len(b) != marshaledSize128 {
//line /usr/local/go/src/hash/fnv/fnv.go:330
		_go_fuzz_dep_.CoverTab[48282]++
							return errors.New("hash/fnv: invalid hash state size")
//line /usr/local/go/src/hash/fnv/fnv.go:331
		// _ = "end of CoverTab[48282]"
	} else {
//line /usr/local/go/src/hash/fnv/fnv.go:332
		_go_fuzz_dep_.CoverTab[48283]++
//line /usr/local/go/src/hash/fnv/fnv.go:332
		// _ = "end of CoverTab[48283]"
//line /usr/local/go/src/hash/fnv/fnv.go:332
	}
//line /usr/local/go/src/hash/fnv/fnv.go:332
	// _ = "end of CoverTab[48277]"
//line /usr/local/go/src/hash/fnv/fnv.go:332
	_go_fuzz_dep_.CoverTab[48278]++
						s[0] = readUint64(b[4:])
						s[1] = readUint64(b[12:])
						return nil
//line /usr/local/go/src/hash/fnv/fnv.go:335
	// _ = "end of CoverTab[48278]"
}

func readUint32(b []byte) uint32 {
//line /usr/local/go/src/hash/fnv/fnv.go:338
	_go_fuzz_dep_.CoverTab[48284]++
						_ = b[3]
						return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
//line /usr/local/go/src/hash/fnv/fnv.go:340
	// _ = "end of CoverTab[48284]"
}

func appendUint32(b []byte, x uint32) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:343
	_go_fuzz_dep_.CoverTab[48285]++
						a := [4]byte{
		byte(x >> 24),
		byte(x >> 16),
		byte(x >> 8),
		byte(x),
	}
						return append(b, a[:]...)
//line /usr/local/go/src/hash/fnv/fnv.go:350
	// _ = "end of CoverTab[48285]"
}

func appendUint64(b []byte, x uint64) []byte {
//line /usr/local/go/src/hash/fnv/fnv.go:353
	_go_fuzz_dep_.CoverTab[48286]++
						a := [8]byte{
		byte(x >> 56),
		byte(x >> 48),
		byte(x >> 40),
		byte(x >> 32),
		byte(x >> 24),
		byte(x >> 16),
		byte(x >> 8),
		byte(x),
	}
						return append(b, a[:]...)
//line /usr/local/go/src/hash/fnv/fnv.go:364
	// _ = "end of CoverTab[48286]"
}

func readUint64(b []byte) uint64 {
//line /usr/local/go/src/hash/fnv/fnv.go:367
	_go_fuzz_dep_.CoverTab[48287]++
						_ = b[7]
						return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
//line /usr/local/go/src/hash/fnv/fnv.go:370
	// _ = "end of CoverTab[48287]"
}

//line /usr/local/go/src/hash/fnv/fnv.go:371
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/hash/fnv/fnv.go:371
var _ = _go_fuzz_dep_.CoverTab
