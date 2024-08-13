// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// W.Hormann, G.Derflinger:
// "Rejection-Inversion to Generate Variates
// from Monotone Discrete Distributions"
// http://eeyore.wu-wien.ac.at/papers/96-04-04.wh-der.ps.gz

//line /usr/local/go/src/math/rand/zipf.go:10
package rand

//line /usr/local/go/src/math/rand/zipf.go:10
import (
//line /usr/local/go/src/math/rand/zipf.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/rand/zipf.go:10
)
//line /usr/local/go/src/math/rand/zipf.go:10
import (
//line /usr/local/go/src/math/rand/zipf.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/rand/zipf.go:10
)

import "math"

// A Zipf generates Zipf distributed variates.
type Zipf struct {
	r		*Rand
	imax		float64
	v		float64
	q		float64
	s		float64
	oneminusQ	float64
	oneminusQinv	float64
	hxm		float64
	hx0minusHxm	float64
}

func (z *Zipf) h(x float64) float64 {
//line /usr/local/go/src/math/rand/zipf.go:27
	_go_fuzz_dep_.CoverTab[3101]++
						return math.Exp(z.oneminusQ*math.Log(z.v+x)) * z.oneminusQinv
//line /usr/local/go/src/math/rand/zipf.go:28
	// _ = "end of CoverTab[3101]"
}

func (z *Zipf) hinv(x float64) float64 {
//line /usr/local/go/src/math/rand/zipf.go:31
	_go_fuzz_dep_.CoverTab[3102]++
						return math.Exp(z.oneminusQinv*math.Log(z.oneminusQ*x)) - z.v
//line /usr/local/go/src/math/rand/zipf.go:32
	// _ = "end of CoverTab[3102]"
}

// NewZipf returns a Zipf variate generator.
//line /usr/local/go/src/math/rand/zipf.go:35
// The generator generates values k ∈ [0, imax]
//line /usr/local/go/src/math/rand/zipf.go:35
// such that P(k) is proportional to (v + k) ** (-s).
//line /usr/local/go/src/math/rand/zipf.go:35
// Requirements: s > 1 and v >= 1.
//line /usr/local/go/src/math/rand/zipf.go:39
func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf {
//line /usr/local/go/src/math/rand/zipf.go:39
	_go_fuzz_dep_.CoverTab[3103]++
						z := new(Zipf)
						if s <= 1.0 || func() bool {
//line /usr/local/go/src/math/rand/zipf.go:41
		_go_fuzz_dep_.CoverTab[3105]++
//line /usr/local/go/src/math/rand/zipf.go:41
		return v < 1
//line /usr/local/go/src/math/rand/zipf.go:41
		// _ = "end of CoverTab[3105]"
//line /usr/local/go/src/math/rand/zipf.go:41
	}() {
//line /usr/local/go/src/math/rand/zipf.go:41
		_go_fuzz_dep_.CoverTab[3106]++
							return nil
//line /usr/local/go/src/math/rand/zipf.go:42
		// _ = "end of CoverTab[3106]"
	} else {
//line /usr/local/go/src/math/rand/zipf.go:43
		_go_fuzz_dep_.CoverTab[3107]++
//line /usr/local/go/src/math/rand/zipf.go:43
		// _ = "end of CoverTab[3107]"
//line /usr/local/go/src/math/rand/zipf.go:43
	}
//line /usr/local/go/src/math/rand/zipf.go:43
	// _ = "end of CoverTab[3103]"
//line /usr/local/go/src/math/rand/zipf.go:43
	_go_fuzz_dep_.CoverTab[3104]++
						z.r = r
						z.imax = float64(imax)
						z.v = v
						z.q = s
						z.oneminusQ = 1.0 - z.q
						z.oneminusQinv = 1.0 / z.oneminusQ
						z.hxm = z.h(z.imax + 0.5)
						z.hx0minusHxm = z.h(0.5) - math.Exp(math.Log(z.v)*(-z.q)) - z.hxm
						z.s = 1 - z.hinv(z.h(1.5)-math.Exp(-z.q*math.Log(z.v+1.0)))
						return z
//line /usr/local/go/src/math/rand/zipf.go:53
	// _ = "end of CoverTab[3104]"
}

// Uint64 returns a value drawn from the Zipf distribution described
//line /usr/local/go/src/math/rand/zipf.go:56
// by the Zipf object.
//line /usr/local/go/src/math/rand/zipf.go:58
func (z *Zipf) Uint64() uint64 {
//line /usr/local/go/src/math/rand/zipf.go:58
	_go_fuzz_dep_.CoverTab[3108]++
						if z == nil {
//line /usr/local/go/src/math/rand/zipf.go:59
		_go_fuzz_dep_.CoverTab[3111]++
							panic("rand: nil Zipf")
//line /usr/local/go/src/math/rand/zipf.go:60
		// _ = "end of CoverTab[3111]"
	} else {
//line /usr/local/go/src/math/rand/zipf.go:61
		_go_fuzz_dep_.CoverTab[3112]++
//line /usr/local/go/src/math/rand/zipf.go:61
		// _ = "end of CoverTab[3112]"
//line /usr/local/go/src/math/rand/zipf.go:61
	}
//line /usr/local/go/src/math/rand/zipf.go:61
	// _ = "end of CoverTab[3108]"
//line /usr/local/go/src/math/rand/zipf.go:61
	_go_fuzz_dep_.CoverTab[3109]++
						k := 0.0

						for {
//line /usr/local/go/src/math/rand/zipf.go:64
		_go_fuzz_dep_.CoverTab[3113]++
							r := z.r.Float64()
							ur := z.hxm + r*z.hx0minusHxm
							x := z.hinv(ur)
							k = math.Floor(x + 0.5)
							if k-x <= z.s {
//line /usr/local/go/src/math/rand/zipf.go:69
			_go_fuzz_dep_.CoverTab[3115]++
								break
//line /usr/local/go/src/math/rand/zipf.go:70
			// _ = "end of CoverTab[3115]"
		} else {
//line /usr/local/go/src/math/rand/zipf.go:71
			_go_fuzz_dep_.CoverTab[3116]++
//line /usr/local/go/src/math/rand/zipf.go:71
			// _ = "end of CoverTab[3116]"
//line /usr/local/go/src/math/rand/zipf.go:71
		}
//line /usr/local/go/src/math/rand/zipf.go:71
		// _ = "end of CoverTab[3113]"
//line /usr/local/go/src/math/rand/zipf.go:71
		_go_fuzz_dep_.CoverTab[3114]++
							if ur >= z.h(k+0.5)-math.Exp(-math.Log(k+z.v)*z.q) {
//line /usr/local/go/src/math/rand/zipf.go:72
			_go_fuzz_dep_.CoverTab[3117]++
								break
//line /usr/local/go/src/math/rand/zipf.go:73
			// _ = "end of CoverTab[3117]"
		} else {
//line /usr/local/go/src/math/rand/zipf.go:74
			_go_fuzz_dep_.CoverTab[3118]++
//line /usr/local/go/src/math/rand/zipf.go:74
			// _ = "end of CoverTab[3118]"
//line /usr/local/go/src/math/rand/zipf.go:74
		}
//line /usr/local/go/src/math/rand/zipf.go:74
		// _ = "end of CoverTab[3114]"
	}
//line /usr/local/go/src/math/rand/zipf.go:75
	// _ = "end of CoverTab[3109]"
//line /usr/local/go/src/math/rand/zipf.go:75
	_go_fuzz_dep_.CoverTab[3110]++
						return uint64(k)
//line /usr/local/go/src/math/rand/zipf.go:76
	// _ = "end of CoverTab[3110]"
}

//line /usr/local/go/src/math/rand/zipf.go:77
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/rand/zipf.go:77
var _ = _go_fuzz_dep_.CoverTab
