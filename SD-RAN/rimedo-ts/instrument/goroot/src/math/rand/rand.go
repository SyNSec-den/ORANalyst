// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/math/rand/rand.go:18
package rand

//line /usr/local/go/src/math/rand/rand.go:18
import (
//line /usr/local/go/src/math/rand/rand.go:18
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/rand/rand.go:18
)
//line /usr/local/go/src/math/rand/rand.go:18
import (
//line /usr/local/go/src/math/rand/rand.go:18
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/rand/rand.go:18
)

import (
	"internal/godebug"
	"sync"
	_ "unsafe"
)

//line /usr/local/go/src/math/rand/rand.go:30
type Source interface {
	Int63() int64
	Seed(seed int64)
}

//line /usr/local/go/src/math/rand/rand.go:41
type Source64 interface {
	Source
	Uint64() uint64
}

//line /usr/local/go/src/math/rand/rand.go:50
func NewSource(seed int64) Source {
//line /usr/local/go/src/math/rand/rand.go:50
	_go_fuzz_dep_.CoverTab[2972]++
						return newSource(seed)
//line /usr/local/go/src/math/rand/rand.go:51
	// _ = "end of CoverTab[2972]"
}

func newSource(seed int64) *rngSource {
//line /usr/local/go/src/math/rand/rand.go:54
	_go_fuzz_dep_.CoverTab[2973]++
						var rng rngSource
						rng.Seed(seed)
						return &rng
//line /usr/local/go/src/math/rand/rand.go:57
	// _ = "end of CoverTab[2973]"
}

//line /usr/local/go/src/math/rand/rand.go:61
type Rand struct {
						src	Source
						s64	Source64

//line /usr/local/go/src/math/rand/rand.go:69
	readVal	int64

//line /usr/local/go/src/math/rand/rand.go:72
	readPos	int8
}

//line /usr/local/go/src/math/rand/rand.go:77
func New(src Source) *Rand {
//line /usr/local/go/src/math/rand/rand.go:77
	_go_fuzz_dep_.CoverTab[2974]++
						s64, _ := src.(Source64)
						return &Rand{src: src, s64: s64}
//line /usr/local/go/src/math/rand/rand.go:79
	// _ = "end of CoverTab[2974]"
}

//line /usr/local/go/src/math/rand/rand.go:84
func (r *Rand) Seed(seed int64) {
//line /usr/local/go/src/math/rand/rand.go:84
	_go_fuzz_dep_.CoverTab[2975]++
						if lk, ok := r.src.(*lockedSource); ok {
//line /usr/local/go/src/math/rand/rand.go:85
		_go_fuzz_dep_.CoverTab[2977]++
							lk.seedPos(seed, &r.readPos)
							return
//line /usr/local/go/src/math/rand/rand.go:87
		// _ = "end of CoverTab[2977]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:88
		_go_fuzz_dep_.CoverTab[2978]++
//line /usr/local/go/src/math/rand/rand.go:88
		// _ = "end of CoverTab[2978]"
//line /usr/local/go/src/math/rand/rand.go:88
	}
//line /usr/local/go/src/math/rand/rand.go:88
	// _ = "end of CoverTab[2975]"
//line /usr/local/go/src/math/rand/rand.go:88
	_go_fuzz_dep_.CoverTab[2976]++

						r.src.Seed(seed)
						r.readPos = 0
//line /usr/local/go/src/math/rand/rand.go:91
	// _ = "end of CoverTab[2976]"
}

//line /usr/local/go/src/math/rand/rand.go:95
func (r *Rand) Int63() int64 {
//line /usr/local/go/src/math/rand/rand.go:95
	_go_fuzz_dep_.CoverTab[2979]++
//line /usr/local/go/src/math/rand/rand.go:95
	return r.src.Int63()
//line /usr/local/go/src/math/rand/rand.go:95
	// _ = "end of CoverTab[2979]"
//line /usr/local/go/src/math/rand/rand.go:95
}

//line /usr/local/go/src/math/rand/rand.go:98
func (r *Rand) Uint32() uint32 {
//line /usr/local/go/src/math/rand/rand.go:98
	_go_fuzz_dep_.CoverTab[2980]++
//line /usr/local/go/src/math/rand/rand.go:98
	return uint32(r.Int63() >> 31)
//line /usr/local/go/src/math/rand/rand.go:98
	// _ = "end of CoverTab[2980]"
//line /usr/local/go/src/math/rand/rand.go:98
}

//line /usr/local/go/src/math/rand/rand.go:101
func (r *Rand) Uint64() uint64 {
//line /usr/local/go/src/math/rand/rand.go:101
	_go_fuzz_dep_.CoverTab[2981]++
						if r.s64 != nil {
//line /usr/local/go/src/math/rand/rand.go:102
		_go_fuzz_dep_.CoverTab[2983]++
							return r.s64.Uint64()
//line /usr/local/go/src/math/rand/rand.go:103
		// _ = "end of CoverTab[2983]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:104
		_go_fuzz_dep_.CoverTab[2984]++
//line /usr/local/go/src/math/rand/rand.go:104
		// _ = "end of CoverTab[2984]"
//line /usr/local/go/src/math/rand/rand.go:104
	}
//line /usr/local/go/src/math/rand/rand.go:104
	// _ = "end of CoverTab[2981]"
//line /usr/local/go/src/math/rand/rand.go:104
	_go_fuzz_dep_.CoverTab[2982]++
						return uint64(r.Int63())>>31 | uint64(r.Int63())<<32
//line /usr/local/go/src/math/rand/rand.go:105
	// _ = "end of CoverTab[2982]"
}

//line /usr/local/go/src/math/rand/rand.go:109
func (r *Rand) Int31() int32 {
//line /usr/local/go/src/math/rand/rand.go:109
	_go_fuzz_dep_.CoverTab[2985]++
//line /usr/local/go/src/math/rand/rand.go:109
	return int32(r.Int63() >> 32)
//line /usr/local/go/src/math/rand/rand.go:109
	// _ = "end of CoverTab[2985]"
//line /usr/local/go/src/math/rand/rand.go:109
}

//line /usr/local/go/src/math/rand/rand.go:112
func (r *Rand) Int() int {
//line /usr/local/go/src/math/rand/rand.go:112
	_go_fuzz_dep_.CoverTab[2986]++
						u := uint(r.Int63())
						return int(u << 1 >> 1)
//line /usr/local/go/src/math/rand/rand.go:114
	// _ = "end of CoverTab[2986]"
}

//line /usr/local/go/src/math/rand/rand.go:119
func (r *Rand) Int63n(n int64) int64 {
//line /usr/local/go/src/math/rand/rand.go:119
	_go_fuzz_dep_.CoverTab[2987]++
						if n <= 0 {
//line /usr/local/go/src/math/rand/rand.go:120
		_go_fuzz_dep_.CoverTab[2991]++
							panic("invalid argument to Int63n")
//line /usr/local/go/src/math/rand/rand.go:121
		// _ = "end of CoverTab[2991]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:122
		_go_fuzz_dep_.CoverTab[2992]++
//line /usr/local/go/src/math/rand/rand.go:122
		// _ = "end of CoverTab[2992]"
//line /usr/local/go/src/math/rand/rand.go:122
	}
//line /usr/local/go/src/math/rand/rand.go:122
	// _ = "end of CoverTab[2987]"
//line /usr/local/go/src/math/rand/rand.go:122
	_go_fuzz_dep_.CoverTab[2988]++
						if n&(n-1) == 0 {
//line /usr/local/go/src/math/rand/rand.go:123
		_go_fuzz_dep_.CoverTab[2993]++
							return r.Int63() & (n - 1)
//line /usr/local/go/src/math/rand/rand.go:124
		// _ = "end of CoverTab[2993]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:125
		_go_fuzz_dep_.CoverTab[2994]++
//line /usr/local/go/src/math/rand/rand.go:125
		// _ = "end of CoverTab[2994]"
//line /usr/local/go/src/math/rand/rand.go:125
	}
//line /usr/local/go/src/math/rand/rand.go:125
	// _ = "end of CoverTab[2988]"
//line /usr/local/go/src/math/rand/rand.go:125
	_go_fuzz_dep_.CoverTab[2989]++
						max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
						v := r.Int63()
						for v > max {
//line /usr/local/go/src/math/rand/rand.go:128
		_go_fuzz_dep_.CoverTab[2995]++
							v = r.Int63()
//line /usr/local/go/src/math/rand/rand.go:129
		// _ = "end of CoverTab[2995]"
	}
//line /usr/local/go/src/math/rand/rand.go:130
	// _ = "end of CoverTab[2989]"
//line /usr/local/go/src/math/rand/rand.go:130
	_go_fuzz_dep_.CoverTab[2990]++
						return v % n
//line /usr/local/go/src/math/rand/rand.go:131
	// _ = "end of CoverTab[2990]"
}

//line /usr/local/go/src/math/rand/rand.go:136
func (r *Rand) Int31n(n int32) int32 {
//line /usr/local/go/src/math/rand/rand.go:136
	_go_fuzz_dep_.CoverTab[2996]++
						if n <= 0 {
//line /usr/local/go/src/math/rand/rand.go:137
		_go_fuzz_dep_.CoverTab[3000]++
							panic("invalid argument to Int31n")
//line /usr/local/go/src/math/rand/rand.go:138
		// _ = "end of CoverTab[3000]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:139
		_go_fuzz_dep_.CoverTab[3001]++
//line /usr/local/go/src/math/rand/rand.go:139
		// _ = "end of CoverTab[3001]"
//line /usr/local/go/src/math/rand/rand.go:139
	}
//line /usr/local/go/src/math/rand/rand.go:139
	// _ = "end of CoverTab[2996]"
//line /usr/local/go/src/math/rand/rand.go:139
	_go_fuzz_dep_.CoverTab[2997]++
						if n&(n-1) == 0 {
//line /usr/local/go/src/math/rand/rand.go:140
		_go_fuzz_dep_.CoverTab[3002]++
							return r.Int31() & (n - 1)
//line /usr/local/go/src/math/rand/rand.go:141
		// _ = "end of CoverTab[3002]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:142
		_go_fuzz_dep_.CoverTab[3003]++
//line /usr/local/go/src/math/rand/rand.go:142
		// _ = "end of CoverTab[3003]"
//line /usr/local/go/src/math/rand/rand.go:142
	}
//line /usr/local/go/src/math/rand/rand.go:142
	// _ = "end of CoverTab[2997]"
//line /usr/local/go/src/math/rand/rand.go:142
	_go_fuzz_dep_.CoverTab[2998]++
						max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
						v := r.Int31()
						for v > max {
//line /usr/local/go/src/math/rand/rand.go:145
		_go_fuzz_dep_.CoverTab[3004]++
							v = r.Int31()
//line /usr/local/go/src/math/rand/rand.go:146
		// _ = "end of CoverTab[3004]"
	}
//line /usr/local/go/src/math/rand/rand.go:147
	// _ = "end of CoverTab[2998]"
//line /usr/local/go/src/math/rand/rand.go:147
	_go_fuzz_dep_.CoverTab[2999]++
						return v % n
//line /usr/local/go/src/math/rand/rand.go:148
	// _ = "end of CoverTab[2999]"
}

//line /usr/local/go/src/math/rand/rand.go:160
func (r *Rand) int31n(n int32) int32 {
//line /usr/local/go/src/math/rand/rand.go:160
	_go_fuzz_dep_.CoverTab[3005]++
						v := r.Uint32()
						prod := uint64(v) * uint64(n)
						low := uint32(prod)
						if low < uint32(n) {
//line /usr/local/go/src/math/rand/rand.go:164
		_go_fuzz_dep_.CoverTab[3007]++
							thresh := uint32(-n) % uint32(n)
							for low < thresh {
//line /usr/local/go/src/math/rand/rand.go:166
			_go_fuzz_dep_.CoverTab[3008]++
								v = r.Uint32()
								prod = uint64(v) * uint64(n)
								low = uint32(prod)
//line /usr/local/go/src/math/rand/rand.go:169
			// _ = "end of CoverTab[3008]"
		}
//line /usr/local/go/src/math/rand/rand.go:170
		// _ = "end of CoverTab[3007]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:171
		_go_fuzz_dep_.CoverTab[3009]++
//line /usr/local/go/src/math/rand/rand.go:171
		// _ = "end of CoverTab[3009]"
//line /usr/local/go/src/math/rand/rand.go:171
	}
//line /usr/local/go/src/math/rand/rand.go:171
	// _ = "end of CoverTab[3005]"
//line /usr/local/go/src/math/rand/rand.go:171
	_go_fuzz_dep_.CoverTab[3006]++
						return int32(prod >> 32)
//line /usr/local/go/src/math/rand/rand.go:172
	// _ = "end of CoverTab[3006]"
}

//line /usr/local/go/src/math/rand/rand.go:177
func (r *Rand) Intn(n int) int {
//line /usr/local/go/src/math/rand/rand.go:177
	_go_fuzz_dep_.CoverTab[3010]++
						if n <= 0 {
//line /usr/local/go/src/math/rand/rand.go:178
		_go_fuzz_dep_.CoverTab[3013]++
							panic("invalid argument to Intn")
//line /usr/local/go/src/math/rand/rand.go:179
		// _ = "end of CoverTab[3013]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:180
		_go_fuzz_dep_.CoverTab[3014]++
//line /usr/local/go/src/math/rand/rand.go:180
		// _ = "end of CoverTab[3014]"
//line /usr/local/go/src/math/rand/rand.go:180
	}
//line /usr/local/go/src/math/rand/rand.go:180
	// _ = "end of CoverTab[3010]"
//line /usr/local/go/src/math/rand/rand.go:180
	_go_fuzz_dep_.CoverTab[3011]++
						if n <= 1<<31-1 {
//line /usr/local/go/src/math/rand/rand.go:181
		_go_fuzz_dep_.CoverTab[3015]++
							return int(r.Int31n(int32(n)))
//line /usr/local/go/src/math/rand/rand.go:182
		// _ = "end of CoverTab[3015]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:183
		_go_fuzz_dep_.CoverTab[3016]++
//line /usr/local/go/src/math/rand/rand.go:183
		// _ = "end of CoverTab[3016]"
//line /usr/local/go/src/math/rand/rand.go:183
	}
//line /usr/local/go/src/math/rand/rand.go:183
	// _ = "end of CoverTab[3011]"
//line /usr/local/go/src/math/rand/rand.go:183
	_go_fuzz_dep_.CoverTab[3012]++
						return int(r.Int63n(int64(n)))
//line /usr/local/go/src/math/rand/rand.go:184
	// _ = "end of CoverTab[3012]"
}

//line /usr/local/go/src/math/rand/rand.go:188
func (r *Rand) Float64() float64 {
//line /usr/local/go/src/math/rand/rand.go:188
	_go_fuzz_dep_.CoverTab[3017]++

//line /usr/local/go/src/math/rand/rand.go:205
again:
	f := float64(r.Int63()) / (1 << 63)
	if f == 1 {
//line /usr/local/go/src/math/rand/rand.go:207
		_go_fuzz_dep_.CoverTab[3019]++
							goto again
//line /usr/local/go/src/math/rand/rand.go:208
		// _ = "end of CoverTab[3019]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:209
		_go_fuzz_dep_.CoverTab[3020]++
//line /usr/local/go/src/math/rand/rand.go:209
		// _ = "end of CoverTab[3020]"
//line /usr/local/go/src/math/rand/rand.go:209
	}
//line /usr/local/go/src/math/rand/rand.go:209
	// _ = "end of CoverTab[3017]"
//line /usr/local/go/src/math/rand/rand.go:209
	_go_fuzz_dep_.CoverTab[3018]++
						return f
//line /usr/local/go/src/math/rand/rand.go:210
	// _ = "end of CoverTab[3018]"
}

//line /usr/local/go/src/math/rand/rand.go:214
func (r *Rand) Float32() float32 {
//line /usr/local/go/src/math/rand/rand.go:214
	_go_fuzz_dep_.CoverTab[3021]++

//line /usr/local/go/src/math/rand/rand.go:218
again:
	f := float32(r.Float64())
	if f == 1 {
//line /usr/local/go/src/math/rand/rand.go:220
		_go_fuzz_dep_.CoverTab[3023]++
							goto again
//line /usr/local/go/src/math/rand/rand.go:221
		// _ = "end of CoverTab[3023]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:222
		_go_fuzz_dep_.CoverTab[3024]++
//line /usr/local/go/src/math/rand/rand.go:222
		// _ = "end of CoverTab[3024]"
//line /usr/local/go/src/math/rand/rand.go:222
	}
//line /usr/local/go/src/math/rand/rand.go:222
	// _ = "end of CoverTab[3021]"
//line /usr/local/go/src/math/rand/rand.go:222
	_go_fuzz_dep_.CoverTab[3022]++
						return f
//line /usr/local/go/src/math/rand/rand.go:223
	// _ = "end of CoverTab[3022]"
}

//line /usr/local/go/src/math/rand/rand.go:228
func (r *Rand) Perm(n int) []int {
//line /usr/local/go/src/math/rand/rand.go:228
	_go_fuzz_dep_.CoverTab[3025]++
						m := make([]int, n)

//line /usr/local/go/src/math/rand/rand.go:235
	for i := 0; i < n; i++ {
//line /usr/local/go/src/math/rand/rand.go:235
		_go_fuzz_dep_.CoverTab[3027]++
							j := r.Intn(i + 1)
							m[i] = m[j]
							m[j] = i
//line /usr/local/go/src/math/rand/rand.go:238
		// _ = "end of CoverTab[3027]"
	}
//line /usr/local/go/src/math/rand/rand.go:239
	// _ = "end of CoverTab[3025]"
//line /usr/local/go/src/math/rand/rand.go:239
	_go_fuzz_dep_.CoverTab[3026]++
						return m
//line /usr/local/go/src/math/rand/rand.go:240
	// _ = "end of CoverTab[3026]"
}

//line /usr/local/go/src/math/rand/rand.go:246
func (r *Rand) Shuffle(n int, swap func(i, j int)) {
//line /usr/local/go/src/math/rand/rand.go:246
	_go_fuzz_dep_.CoverTab[3028]++
						if n < 0 {
//line /usr/local/go/src/math/rand/rand.go:247
		_go_fuzz_dep_.CoverTab[3031]++
							panic("invalid argument to Shuffle")
//line /usr/local/go/src/math/rand/rand.go:248
		// _ = "end of CoverTab[3031]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:249
		_go_fuzz_dep_.CoverTab[3032]++
//line /usr/local/go/src/math/rand/rand.go:249
		// _ = "end of CoverTab[3032]"
//line /usr/local/go/src/math/rand/rand.go:249
	}
//line /usr/local/go/src/math/rand/rand.go:249
	// _ = "end of CoverTab[3028]"
//line /usr/local/go/src/math/rand/rand.go:249
	_go_fuzz_dep_.CoverTab[3029]++

//line /usr/local/go/src/math/rand/rand.go:257
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
//line /usr/local/go/src/math/rand/rand.go:258
		_go_fuzz_dep_.CoverTab[3033]++
							j := int(r.Int63n(int64(i + 1)))
							swap(i, j)
//line /usr/local/go/src/math/rand/rand.go:260
		// _ = "end of CoverTab[3033]"
	}
//line /usr/local/go/src/math/rand/rand.go:261
	// _ = "end of CoverTab[3029]"
//line /usr/local/go/src/math/rand/rand.go:261
	_go_fuzz_dep_.CoverTab[3030]++
						for ; i > 0; i-- {
//line /usr/local/go/src/math/rand/rand.go:262
		_go_fuzz_dep_.CoverTab[3034]++
							j := int(r.int31n(int32(i + 1)))
							swap(i, j)
//line /usr/local/go/src/math/rand/rand.go:264
		// _ = "end of CoverTab[3034]"
	}
//line /usr/local/go/src/math/rand/rand.go:265
	// _ = "end of CoverTab[3030]"
}

//line /usr/local/go/src/math/rand/rand.go:271
func (r *Rand) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/math/rand/rand.go:271
	_go_fuzz_dep_.CoverTab[3035]++
						if lk, ok := r.src.(*lockedSource); ok {
//line /usr/local/go/src/math/rand/rand.go:272
		_go_fuzz_dep_.CoverTab[3037]++
							return lk.read(p, &r.readVal, &r.readPos)
//line /usr/local/go/src/math/rand/rand.go:273
		// _ = "end of CoverTab[3037]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:274
		_go_fuzz_dep_.CoverTab[3038]++
//line /usr/local/go/src/math/rand/rand.go:274
		// _ = "end of CoverTab[3038]"
//line /usr/local/go/src/math/rand/rand.go:274
	}
//line /usr/local/go/src/math/rand/rand.go:274
	// _ = "end of CoverTab[3035]"
//line /usr/local/go/src/math/rand/rand.go:274
	_go_fuzz_dep_.CoverTab[3036]++
						return read(p, r.src, &r.readVal, &r.readPos)
//line /usr/local/go/src/math/rand/rand.go:275
	// _ = "end of CoverTab[3036]"
}

func read(p []byte, src Source, readVal *int64, readPos *int8) (n int, err error) {
//line /usr/local/go/src/math/rand/rand.go:278
	_go_fuzz_dep_.CoverTab[3039]++
						pos := *readPos
						val := *readVal
						rng, _ := src.(*rngSource)
						for n = 0; n < len(p); n++ {
//line /usr/local/go/src/math/rand/rand.go:282
		_go_fuzz_dep_.CoverTab[3041]++
							if pos == 0 {
//line /usr/local/go/src/math/rand/rand.go:283
			_go_fuzz_dep_.CoverTab[3043]++
								if rng != nil {
//line /usr/local/go/src/math/rand/rand.go:284
				_go_fuzz_dep_.CoverTab[3045]++
									val = rng.Int63()
//line /usr/local/go/src/math/rand/rand.go:285
				// _ = "end of CoverTab[3045]"
			} else {
//line /usr/local/go/src/math/rand/rand.go:286
				_go_fuzz_dep_.CoverTab[3046]++
									val = src.Int63()
//line /usr/local/go/src/math/rand/rand.go:287
				// _ = "end of CoverTab[3046]"
			}
//line /usr/local/go/src/math/rand/rand.go:288
			// _ = "end of CoverTab[3043]"
//line /usr/local/go/src/math/rand/rand.go:288
			_go_fuzz_dep_.CoverTab[3044]++
								pos = 7
//line /usr/local/go/src/math/rand/rand.go:289
			// _ = "end of CoverTab[3044]"
		} else {
//line /usr/local/go/src/math/rand/rand.go:290
			_go_fuzz_dep_.CoverTab[3047]++
//line /usr/local/go/src/math/rand/rand.go:290
			// _ = "end of CoverTab[3047]"
//line /usr/local/go/src/math/rand/rand.go:290
		}
//line /usr/local/go/src/math/rand/rand.go:290
		// _ = "end of CoverTab[3041]"
//line /usr/local/go/src/math/rand/rand.go:290
		_go_fuzz_dep_.CoverTab[3042]++
							p[n] = byte(val)
							val >>= 8
							pos--
//line /usr/local/go/src/math/rand/rand.go:293
		// _ = "end of CoverTab[3042]"
	}
//line /usr/local/go/src/math/rand/rand.go:294
	// _ = "end of CoverTab[3039]"
//line /usr/local/go/src/math/rand/rand.go:294
	_go_fuzz_dep_.CoverTab[3040]++
						*readPos = pos
						*readVal = val
						return
//line /usr/local/go/src/math/rand/rand.go:297
	// _ = "end of CoverTab[3040]"
}

//line /usr/local/go/src/math/rand/rand.go:304
var globalRand = New(new(lockedSource))

//line /usr/local/go/src/math/rand/rand.go:324
func Seed(seed int64) {
//line /usr/local/go/src/math/rand/rand.go:324
	_go_fuzz_dep_.CoverTab[3048]++
//line /usr/local/go/src/math/rand/rand.go:324
	globalRand.Seed(seed)
//line /usr/local/go/src/math/rand/rand.go:324
	// _ = "end of CoverTab[3048]"
//line /usr/local/go/src/math/rand/rand.go:324
}

//line /usr/local/go/src/math/rand/rand.go:328
func Int63() int64 {
//line /usr/local/go/src/math/rand/rand.go:328
	_go_fuzz_dep_.CoverTab[3049]++
//line /usr/local/go/src/math/rand/rand.go:328
	return globalRand.Int63()
//line /usr/local/go/src/math/rand/rand.go:328
	// _ = "end of CoverTab[3049]"
//line /usr/local/go/src/math/rand/rand.go:328
}

//line /usr/local/go/src/math/rand/rand.go:332
func Uint32() uint32 {
//line /usr/local/go/src/math/rand/rand.go:332
	_go_fuzz_dep_.CoverTab[3050]++
//line /usr/local/go/src/math/rand/rand.go:332
	return globalRand.Uint32()
//line /usr/local/go/src/math/rand/rand.go:332
	// _ = "end of CoverTab[3050]"
//line /usr/local/go/src/math/rand/rand.go:332
}

//line /usr/local/go/src/math/rand/rand.go:336
func Uint64() uint64 {
//line /usr/local/go/src/math/rand/rand.go:336
	_go_fuzz_dep_.CoverTab[3051]++
//line /usr/local/go/src/math/rand/rand.go:336
	return globalRand.Uint64()
//line /usr/local/go/src/math/rand/rand.go:336
	// _ = "end of CoverTab[3051]"
//line /usr/local/go/src/math/rand/rand.go:336
}

//line /usr/local/go/src/math/rand/rand.go:340
func Int31() int32 {
//line /usr/local/go/src/math/rand/rand.go:340
	_go_fuzz_dep_.CoverTab[3052]++
//line /usr/local/go/src/math/rand/rand.go:340
	return globalRand.Int31()
//line /usr/local/go/src/math/rand/rand.go:340
	// _ = "end of CoverTab[3052]"
//line /usr/local/go/src/math/rand/rand.go:340
}

//line /usr/local/go/src/math/rand/rand.go:343
func Int() int {
//line /usr/local/go/src/math/rand/rand.go:343
	_go_fuzz_dep_.CoverTab[3053]++
//line /usr/local/go/src/math/rand/rand.go:343
	return globalRand.Int()
//line /usr/local/go/src/math/rand/rand.go:343
	// _ = "end of CoverTab[3053]"
//line /usr/local/go/src/math/rand/rand.go:343
}

//line /usr/local/go/src/math/rand/rand.go:348
func Int63n(n int64) int64 {
//line /usr/local/go/src/math/rand/rand.go:348
	_go_fuzz_dep_.CoverTab[3054]++
//line /usr/local/go/src/math/rand/rand.go:348
	return globalRand.Int63n(n)
//line /usr/local/go/src/math/rand/rand.go:348
	// _ = "end of CoverTab[3054]"
//line /usr/local/go/src/math/rand/rand.go:348
}

//line /usr/local/go/src/math/rand/rand.go:353
func Int31n(n int32) int32 {
//line /usr/local/go/src/math/rand/rand.go:353
	_go_fuzz_dep_.CoverTab[3055]++
//line /usr/local/go/src/math/rand/rand.go:353
	return globalRand.Int31n(n)
//line /usr/local/go/src/math/rand/rand.go:353
	// _ = "end of CoverTab[3055]"
//line /usr/local/go/src/math/rand/rand.go:353
}

//line /usr/local/go/src/math/rand/rand.go:358
func Intn(n int) int {
//line /usr/local/go/src/math/rand/rand.go:358
	_go_fuzz_dep_.CoverTab[3056]++
//line /usr/local/go/src/math/rand/rand.go:358
	return globalRand.Intn(n)
//line /usr/local/go/src/math/rand/rand.go:358
	// _ = "end of CoverTab[3056]"
//line /usr/local/go/src/math/rand/rand.go:358
}

//line /usr/local/go/src/math/rand/rand.go:362
func Float64() float64 {
//line /usr/local/go/src/math/rand/rand.go:362
	_go_fuzz_dep_.CoverTab[3057]++
//line /usr/local/go/src/math/rand/rand.go:362
	return globalRand.Float64()
//line /usr/local/go/src/math/rand/rand.go:362
	// _ = "end of CoverTab[3057]"
//line /usr/local/go/src/math/rand/rand.go:362
}

//line /usr/local/go/src/math/rand/rand.go:366
func Float32() float32 {
//line /usr/local/go/src/math/rand/rand.go:366
	_go_fuzz_dep_.CoverTab[3058]++
//line /usr/local/go/src/math/rand/rand.go:366
	return globalRand.Float32()
//line /usr/local/go/src/math/rand/rand.go:366
	// _ = "end of CoverTab[3058]"
//line /usr/local/go/src/math/rand/rand.go:366
}

//line /usr/local/go/src/math/rand/rand.go:370
func Perm(n int) []int {
//line /usr/local/go/src/math/rand/rand.go:370
	_go_fuzz_dep_.CoverTab[3059]++
//line /usr/local/go/src/math/rand/rand.go:370
	return globalRand.Perm(n)
//line /usr/local/go/src/math/rand/rand.go:370
	// _ = "end of CoverTab[3059]"
//line /usr/local/go/src/math/rand/rand.go:370
}

//line /usr/local/go/src/math/rand/rand.go:375
func Shuffle(n int, swap func(i, j int)) {
//line /usr/local/go/src/math/rand/rand.go:375
	_go_fuzz_dep_.CoverTab[3060]++
//line /usr/local/go/src/math/rand/rand.go:375
	globalRand.Shuffle(n, swap)
//line /usr/local/go/src/math/rand/rand.go:375
	// _ = "end of CoverTab[3060]"
//line /usr/local/go/src/math/rand/rand.go:375
}

//line /usr/local/go/src/math/rand/rand.go:382
func Read(p []byte) (n int, err error) {
//line /usr/local/go/src/math/rand/rand.go:382
	_go_fuzz_dep_.CoverTab[3061]++
//line /usr/local/go/src/math/rand/rand.go:382
	return globalRand.Read(p)
//line /usr/local/go/src/math/rand/rand.go:382
	// _ = "end of CoverTab[3061]"
//line /usr/local/go/src/math/rand/rand.go:382
}

//line /usr/local/go/src/math/rand/rand.go:392
func NormFloat64() float64 {
//line /usr/local/go/src/math/rand/rand.go:392
	_go_fuzz_dep_.CoverTab[3062]++
//line /usr/local/go/src/math/rand/rand.go:392
	return globalRand.NormFloat64()
//line /usr/local/go/src/math/rand/rand.go:392
	// _ = "end of CoverTab[3062]"
//line /usr/local/go/src/math/rand/rand.go:392
}

//line /usr/local/go/src/math/rand/rand.go:401
func ExpFloat64() float64 {
//line /usr/local/go/src/math/rand/rand.go:401
	_go_fuzz_dep_.CoverTab[3063]++
//line /usr/local/go/src/math/rand/rand.go:401
	return globalRand.ExpFloat64()
//line /usr/local/go/src/math/rand/rand.go:401
	// _ = "end of CoverTab[3063]"
//line /usr/local/go/src/math/rand/rand.go:401
}

type lockedSource struct {
	lk	sync.Mutex
	s	*rngSource
}

//go:linkname fastrand64
func fastrand64() uint64

var randautoseed = godebug.New("randautoseed")

//line /usr/local/go/src/math/rand/rand.go:415
func (r *lockedSource) source() *rngSource {
//line /usr/local/go/src/math/rand/rand.go:415
	_go_fuzz_dep_.CoverTab[3064]++
						if r.s == nil {
//line /usr/local/go/src/math/rand/rand.go:416
		_go_fuzz_dep_.CoverTab[3066]++
							var seed int64
							if randautoseed.Value() == "0" {
//line /usr/local/go/src/math/rand/rand.go:418
			_go_fuzz_dep_.CoverTab[3068]++
								seed = 1
//line /usr/local/go/src/math/rand/rand.go:419
			// _ = "end of CoverTab[3068]"
		} else {
//line /usr/local/go/src/math/rand/rand.go:420
			_go_fuzz_dep_.CoverTab[3069]++
								seed = int64(fastrand64())
//line /usr/local/go/src/math/rand/rand.go:421
			// _ = "end of CoverTab[3069]"
		}
//line /usr/local/go/src/math/rand/rand.go:422
		// _ = "end of CoverTab[3066]"
//line /usr/local/go/src/math/rand/rand.go:422
		_go_fuzz_dep_.CoverTab[3067]++
							r.s = newSource(seed)
//line /usr/local/go/src/math/rand/rand.go:423
		// _ = "end of CoverTab[3067]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:424
		_go_fuzz_dep_.CoverTab[3070]++
//line /usr/local/go/src/math/rand/rand.go:424
		// _ = "end of CoverTab[3070]"
//line /usr/local/go/src/math/rand/rand.go:424
	}
//line /usr/local/go/src/math/rand/rand.go:424
	// _ = "end of CoverTab[3064]"
//line /usr/local/go/src/math/rand/rand.go:424
	_go_fuzz_dep_.CoverTab[3065]++
						return r.s
//line /usr/local/go/src/math/rand/rand.go:425
	// _ = "end of CoverTab[3065]"
}

func (r *lockedSource) Int63() (n int64) {
//line /usr/local/go/src/math/rand/rand.go:428
	_go_fuzz_dep_.CoverTab[3071]++
						r.lk.Lock()
						n = r.source().Int63()
						r.lk.Unlock()
						return
//line /usr/local/go/src/math/rand/rand.go:432
	// _ = "end of CoverTab[3071]"
}

func (r *lockedSource) Uint64() (n uint64) {
//line /usr/local/go/src/math/rand/rand.go:435
	_go_fuzz_dep_.CoverTab[3072]++
						r.lk.Lock()
						n = r.source().Uint64()
						r.lk.Unlock()
						return
//line /usr/local/go/src/math/rand/rand.go:439
	// _ = "end of CoverTab[3072]"
}

func (r *lockedSource) Seed(seed int64) {
//line /usr/local/go/src/math/rand/rand.go:442
	_go_fuzz_dep_.CoverTab[3073]++
						r.lk.Lock()
						r.seed(seed)
						r.lk.Unlock()
//line /usr/local/go/src/math/rand/rand.go:445
	// _ = "end of CoverTab[3073]"
}

//line /usr/local/go/src/math/rand/rand.go:449
func (r *lockedSource) seedPos(seed int64, readPos *int8) {
//line /usr/local/go/src/math/rand/rand.go:449
	_go_fuzz_dep_.CoverTab[3074]++
						r.lk.Lock()
						r.seed(seed)
						*readPos = 0
						r.lk.Unlock()
//line /usr/local/go/src/math/rand/rand.go:453
	// _ = "end of CoverTab[3074]"
}

//line /usr/local/go/src/math/rand/rand.go:458
func (r *lockedSource) seed(seed int64) {
//line /usr/local/go/src/math/rand/rand.go:458
	_go_fuzz_dep_.CoverTab[3075]++
						if r.s == nil {
//line /usr/local/go/src/math/rand/rand.go:459
		_go_fuzz_dep_.CoverTab[3076]++
							r.s = newSource(seed)
//line /usr/local/go/src/math/rand/rand.go:460
		// _ = "end of CoverTab[3076]"
	} else {
//line /usr/local/go/src/math/rand/rand.go:461
		_go_fuzz_dep_.CoverTab[3077]++
							r.s.Seed(seed)
//line /usr/local/go/src/math/rand/rand.go:462
		// _ = "end of CoverTab[3077]"
	}
//line /usr/local/go/src/math/rand/rand.go:463
	// _ = "end of CoverTab[3075]"
}

//line /usr/local/go/src/math/rand/rand.go:467
func (r *lockedSource) read(p []byte, readVal *int64, readPos *int8) (n int, err error) {
//line /usr/local/go/src/math/rand/rand.go:467
	_go_fuzz_dep_.CoverTab[3078]++
						r.lk.Lock()
						n, err = read(p, r.source(), readVal, readPos)
						r.lk.Unlock()
						return
//line /usr/local/go/src/math/rand/rand.go:471
	// _ = "end of CoverTab[3078]"
}

//line /usr/local/go/src/math/rand/rand.go:472
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/rand/rand.go:472
var _ = _go_fuzz_dep_.CoverTab
