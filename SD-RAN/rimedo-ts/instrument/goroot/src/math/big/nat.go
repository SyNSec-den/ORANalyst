// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements unsigned multi-precision integers (natural
// numbers). They are the building blocks for the implementation
// of signed integers, rationals, and floating-point numbers.
//
// Caution: This implementation relies on the function "alias"
//          which assumes that (nat) slice capacities are never
//          changed (no 3-operand slice expressions). If that
//          changes, alias needs to be updated for correctness.

//line /usr/local/go/src/math/big/nat.go:14
package big

//line /usr/local/go/src/math/big/nat.go:14
import (
//line /usr/local/go/src/math/big/nat.go:14
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/nat.go:14
)
//line /usr/local/go/src/math/big/nat.go:14
import (
//line /usr/local/go/src/math/big/nat.go:14
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/nat.go:14
)

import (
	"encoding/binary"
	"math/bits"
	"math/rand"
	"sync"
)

// An unsigned integer x of the form
//line /usr/local/go/src/math/big/nat.go:23
//
//line /usr/local/go/src/math/big/nat.go:23
//	x = x[n-1]*_B^(n-1) + x[n-2]*_B^(n-2) + ... + x[1]*_B + x[0]
//line /usr/local/go/src/math/big/nat.go:23
//
//line /usr/local/go/src/math/big/nat.go:23
// with 0 <= x[i] < _B and 0 <= i < n is stored in a slice of length n,
//line /usr/local/go/src/math/big/nat.go:23
// with the digits x[i] as the slice elements.
//line /usr/local/go/src/math/big/nat.go:23
//
//line /usr/local/go/src/math/big/nat.go:23
// A number is normalized if the slice contains no leading 0 digits.
//line /usr/local/go/src/math/big/nat.go:23
// During arithmetic operations, denormalized values may occur but are
//line /usr/local/go/src/math/big/nat.go:23
// always normalized before returning the final result. The normalized
//line /usr/local/go/src/math/big/nat.go:23
// representation of 0 is the empty or nil slice (length = 0).
//line /usr/local/go/src/math/big/nat.go:34
type nat []Word

var (
	natOne	= nat{1}
	natTwo	= nat{2}
	natFive	= nat{5}
	natTen	= nat{10}
)

func (z nat) String() string {
//line /usr/local/go/src/math/big/nat.go:43
	_go_fuzz_dep_.CoverTab[5674]++
						return "0x" + string(z.itoa(false, 16))
//line /usr/local/go/src/math/big/nat.go:44
	// _ = "end of CoverTab[5674]"
}

func (z nat) clear() {
//line /usr/local/go/src/math/big/nat.go:47
	_go_fuzz_dep_.CoverTab[5675]++
						for i := range z {
//line /usr/local/go/src/math/big/nat.go:48
		_go_fuzz_dep_.CoverTab[5676]++
							z[i] = 0
//line /usr/local/go/src/math/big/nat.go:49
		// _ = "end of CoverTab[5676]"
	}
//line /usr/local/go/src/math/big/nat.go:50
	// _ = "end of CoverTab[5675]"
}

func (z nat) norm() nat {
//line /usr/local/go/src/math/big/nat.go:53
	_go_fuzz_dep_.CoverTab[5677]++
						i := len(z)
						for i > 0 && func() bool {
//line /usr/local/go/src/math/big/nat.go:55
		_go_fuzz_dep_.CoverTab[5679]++
//line /usr/local/go/src/math/big/nat.go:55
		return z[i-1] == 0
//line /usr/local/go/src/math/big/nat.go:55
		// _ = "end of CoverTab[5679]"
//line /usr/local/go/src/math/big/nat.go:55
	}() {
//line /usr/local/go/src/math/big/nat.go:55
		_go_fuzz_dep_.CoverTab[5680]++
							i--
//line /usr/local/go/src/math/big/nat.go:56
		// _ = "end of CoverTab[5680]"
	}
//line /usr/local/go/src/math/big/nat.go:57
	// _ = "end of CoverTab[5677]"
//line /usr/local/go/src/math/big/nat.go:57
	_go_fuzz_dep_.CoverTab[5678]++
						return z[0:i]
//line /usr/local/go/src/math/big/nat.go:58
	// _ = "end of CoverTab[5678]"
}

func (z nat) make(n int) nat {
//line /usr/local/go/src/math/big/nat.go:61
	_go_fuzz_dep_.CoverTab[5681]++
						if n <= cap(z) {
//line /usr/local/go/src/math/big/nat.go:62
		_go_fuzz_dep_.CoverTab[5684]++
							return z[:n]
//line /usr/local/go/src/math/big/nat.go:63
		// _ = "end of CoverTab[5684]"
	} else {
//line /usr/local/go/src/math/big/nat.go:64
		_go_fuzz_dep_.CoverTab[5685]++
//line /usr/local/go/src/math/big/nat.go:64
		// _ = "end of CoverTab[5685]"
//line /usr/local/go/src/math/big/nat.go:64
	}
//line /usr/local/go/src/math/big/nat.go:64
	// _ = "end of CoverTab[5681]"
//line /usr/local/go/src/math/big/nat.go:64
	_go_fuzz_dep_.CoverTab[5682]++
						if n == 1 {
//line /usr/local/go/src/math/big/nat.go:65
		_go_fuzz_dep_.CoverTab[5686]++

							return make(nat, 1)
//line /usr/local/go/src/math/big/nat.go:67
		// _ = "end of CoverTab[5686]"
	} else {
//line /usr/local/go/src/math/big/nat.go:68
		_go_fuzz_dep_.CoverTab[5687]++
//line /usr/local/go/src/math/big/nat.go:68
		// _ = "end of CoverTab[5687]"
//line /usr/local/go/src/math/big/nat.go:68
	}
//line /usr/local/go/src/math/big/nat.go:68
	// _ = "end of CoverTab[5682]"
//line /usr/local/go/src/math/big/nat.go:68
	_go_fuzz_dep_.CoverTab[5683]++
	// Choosing a good value for e has significant performance impact
						// because it increases the chance that a value can be reused.
						const e = 4	// extra capacity
						return make(nat, n, n+e)
//line /usr/local/go/src/math/big/nat.go:72
	// _ = "end of CoverTab[5683]"
}

func (z nat) setWord(x Word) nat {
//line /usr/local/go/src/math/big/nat.go:75
	_go_fuzz_dep_.CoverTab[5688]++
						if x == 0 {
//line /usr/local/go/src/math/big/nat.go:76
		_go_fuzz_dep_.CoverTab[5690]++
							return z[:0]
//line /usr/local/go/src/math/big/nat.go:77
		// _ = "end of CoverTab[5690]"
	} else {
//line /usr/local/go/src/math/big/nat.go:78
		_go_fuzz_dep_.CoverTab[5691]++
//line /usr/local/go/src/math/big/nat.go:78
		// _ = "end of CoverTab[5691]"
//line /usr/local/go/src/math/big/nat.go:78
	}
//line /usr/local/go/src/math/big/nat.go:78
	// _ = "end of CoverTab[5688]"
//line /usr/local/go/src/math/big/nat.go:78
	_go_fuzz_dep_.CoverTab[5689]++
						z = z.make(1)
						z[0] = x
						return z
//line /usr/local/go/src/math/big/nat.go:81
	// _ = "end of CoverTab[5689]"
}

func (z nat) setUint64(x uint64) nat {
//line /usr/local/go/src/math/big/nat.go:84
	_go_fuzz_dep_.CoverTab[5692]++

						if w := Word(x); uint64(w) == x {
//line /usr/local/go/src/math/big/nat.go:86
		_go_fuzz_dep_.CoverTab[5694]++
							return z.setWord(w)
//line /usr/local/go/src/math/big/nat.go:87
		// _ = "end of CoverTab[5694]"
	} else {
//line /usr/local/go/src/math/big/nat.go:88
		_go_fuzz_dep_.CoverTab[5695]++
//line /usr/local/go/src/math/big/nat.go:88
		// _ = "end of CoverTab[5695]"
//line /usr/local/go/src/math/big/nat.go:88
	}
//line /usr/local/go/src/math/big/nat.go:88
	// _ = "end of CoverTab[5692]"
//line /usr/local/go/src/math/big/nat.go:88
	_go_fuzz_dep_.CoverTab[5693]++

						z = z.make(2)
						z[1] = Word(x >> 32)
						z[0] = Word(x)
						return z
//line /usr/local/go/src/math/big/nat.go:93
	// _ = "end of CoverTab[5693]"
}

func (z nat) set(x nat) nat {
//line /usr/local/go/src/math/big/nat.go:96
	_go_fuzz_dep_.CoverTab[5696]++
						z = z.make(len(x))
						copy(z, x)
						return z
//line /usr/local/go/src/math/big/nat.go:99
	// _ = "end of CoverTab[5696]"
}

func (z nat) add(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:102
	_go_fuzz_dep_.CoverTab[5697]++
						m := len(x)
						n := len(y)

						switch {
	case m < n:
//line /usr/local/go/src/math/big/nat.go:107
		_go_fuzz_dep_.CoverTab[5700]++
							return z.add(y, x)
//line /usr/local/go/src/math/big/nat.go:108
		// _ = "end of CoverTab[5700]"
	case m == 0:
//line /usr/local/go/src/math/big/nat.go:109
		_go_fuzz_dep_.CoverTab[5701]++

							return z[:0]
//line /usr/local/go/src/math/big/nat.go:111
		// _ = "end of CoverTab[5701]"
	case n == 0:
//line /usr/local/go/src/math/big/nat.go:112
		_go_fuzz_dep_.CoverTab[5702]++

							return z.set(x)
//line /usr/local/go/src/math/big/nat.go:114
		// _ = "end of CoverTab[5702]"
//line /usr/local/go/src/math/big/nat.go:114
	default:
//line /usr/local/go/src/math/big/nat.go:114
		_go_fuzz_dep_.CoverTab[5703]++
//line /usr/local/go/src/math/big/nat.go:114
		// _ = "end of CoverTab[5703]"
	}
//line /usr/local/go/src/math/big/nat.go:115
	// _ = "end of CoverTab[5697]"
//line /usr/local/go/src/math/big/nat.go:115
	_go_fuzz_dep_.CoverTab[5698]++

//line /usr/local/go/src/math/big/nat.go:118
	z = z.make(m + 1)
	c := addVV(z[0:n], x, y)
	if m > n {
//line /usr/local/go/src/math/big/nat.go:120
		_go_fuzz_dep_.CoverTab[5704]++
							c = addVW(z[n:m], x[n:], c)
//line /usr/local/go/src/math/big/nat.go:121
		// _ = "end of CoverTab[5704]"
	} else {
//line /usr/local/go/src/math/big/nat.go:122
		_go_fuzz_dep_.CoverTab[5705]++
//line /usr/local/go/src/math/big/nat.go:122
		// _ = "end of CoverTab[5705]"
//line /usr/local/go/src/math/big/nat.go:122
	}
//line /usr/local/go/src/math/big/nat.go:122
	// _ = "end of CoverTab[5698]"
//line /usr/local/go/src/math/big/nat.go:122
	_go_fuzz_dep_.CoverTab[5699]++
						z[m] = c

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:125
	// _ = "end of CoverTab[5699]"
}

func (z nat) sub(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:128
	_go_fuzz_dep_.CoverTab[5706]++
						m := len(x)
						n := len(y)

						switch {
	case m < n:
//line /usr/local/go/src/math/big/nat.go:133
		_go_fuzz_dep_.CoverTab[5710]++
							panic("underflow")
//line /usr/local/go/src/math/big/nat.go:134
		// _ = "end of CoverTab[5710]"
	case m == 0:
//line /usr/local/go/src/math/big/nat.go:135
		_go_fuzz_dep_.CoverTab[5711]++

							return z[:0]
//line /usr/local/go/src/math/big/nat.go:137
		// _ = "end of CoverTab[5711]"
	case n == 0:
//line /usr/local/go/src/math/big/nat.go:138
		_go_fuzz_dep_.CoverTab[5712]++

							return z.set(x)
//line /usr/local/go/src/math/big/nat.go:140
		// _ = "end of CoverTab[5712]"
//line /usr/local/go/src/math/big/nat.go:140
	default:
//line /usr/local/go/src/math/big/nat.go:140
		_go_fuzz_dep_.CoverTab[5713]++
//line /usr/local/go/src/math/big/nat.go:140
		// _ = "end of CoverTab[5713]"
	}
//line /usr/local/go/src/math/big/nat.go:141
	// _ = "end of CoverTab[5706]"
//line /usr/local/go/src/math/big/nat.go:141
	_go_fuzz_dep_.CoverTab[5707]++

//line /usr/local/go/src/math/big/nat.go:144
	z = z.make(m)
	c := subVV(z[0:n], x, y)
	if m > n {
//line /usr/local/go/src/math/big/nat.go:146
		_go_fuzz_dep_.CoverTab[5714]++
							c = subVW(z[n:], x[n:], c)
//line /usr/local/go/src/math/big/nat.go:147
		// _ = "end of CoverTab[5714]"
	} else {
//line /usr/local/go/src/math/big/nat.go:148
		_go_fuzz_dep_.CoverTab[5715]++
//line /usr/local/go/src/math/big/nat.go:148
		// _ = "end of CoverTab[5715]"
//line /usr/local/go/src/math/big/nat.go:148
	}
//line /usr/local/go/src/math/big/nat.go:148
	// _ = "end of CoverTab[5707]"
//line /usr/local/go/src/math/big/nat.go:148
	_go_fuzz_dep_.CoverTab[5708]++
						if c != 0 {
//line /usr/local/go/src/math/big/nat.go:149
		_go_fuzz_dep_.CoverTab[5716]++
							panic("underflow")
//line /usr/local/go/src/math/big/nat.go:150
		// _ = "end of CoverTab[5716]"
	} else {
//line /usr/local/go/src/math/big/nat.go:151
		_go_fuzz_dep_.CoverTab[5717]++
//line /usr/local/go/src/math/big/nat.go:151
		// _ = "end of CoverTab[5717]"
//line /usr/local/go/src/math/big/nat.go:151
	}
//line /usr/local/go/src/math/big/nat.go:151
	// _ = "end of CoverTab[5708]"
//line /usr/local/go/src/math/big/nat.go:151
	_go_fuzz_dep_.CoverTab[5709]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:153
	// _ = "end of CoverTab[5709]"
}

func (x nat) cmp(y nat) (r int) {
//line /usr/local/go/src/math/big/nat.go:156
	_go_fuzz_dep_.CoverTab[5718]++
						m := len(x)
						n := len(y)
						if m != n || func() bool {
//line /usr/local/go/src/math/big/nat.go:159
		_go_fuzz_dep_.CoverTab[5722]++
//line /usr/local/go/src/math/big/nat.go:159
		return m == 0
//line /usr/local/go/src/math/big/nat.go:159
		// _ = "end of CoverTab[5722]"
//line /usr/local/go/src/math/big/nat.go:159
	}() {
//line /usr/local/go/src/math/big/nat.go:159
		_go_fuzz_dep_.CoverTab[5723]++
							switch {
		case m < n:
//line /usr/local/go/src/math/big/nat.go:161
			_go_fuzz_dep_.CoverTab[5725]++
								r = -1
//line /usr/local/go/src/math/big/nat.go:162
			// _ = "end of CoverTab[5725]"
		case m > n:
//line /usr/local/go/src/math/big/nat.go:163
			_go_fuzz_dep_.CoverTab[5726]++
								r = 1
//line /usr/local/go/src/math/big/nat.go:164
			// _ = "end of CoverTab[5726]"
//line /usr/local/go/src/math/big/nat.go:164
		default:
//line /usr/local/go/src/math/big/nat.go:164
			_go_fuzz_dep_.CoverTab[5727]++
//line /usr/local/go/src/math/big/nat.go:164
			// _ = "end of CoverTab[5727]"
		}
//line /usr/local/go/src/math/big/nat.go:165
		// _ = "end of CoverTab[5723]"
//line /usr/local/go/src/math/big/nat.go:165
		_go_fuzz_dep_.CoverTab[5724]++
							return
//line /usr/local/go/src/math/big/nat.go:166
		// _ = "end of CoverTab[5724]"
	} else {
//line /usr/local/go/src/math/big/nat.go:167
		_go_fuzz_dep_.CoverTab[5728]++
//line /usr/local/go/src/math/big/nat.go:167
		// _ = "end of CoverTab[5728]"
//line /usr/local/go/src/math/big/nat.go:167
	}
//line /usr/local/go/src/math/big/nat.go:167
	// _ = "end of CoverTab[5718]"
//line /usr/local/go/src/math/big/nat.go:167
	_go_fuzz_dep_.CoverTab[5719]++

						i := m - 1
						for i > 0 && func() bool {
//line /usr/local/go/src/math/big/nat.go:170
		_go_fuzz_dep_.CoverTab[5729]++
//line /usr/local/go/src/math/big/nat.go:170
		return x[i] == y[i]
//line /usr/local/go/src/math/big/nat.go:170
		// _ = "end of CoverTab[5729]"
//line /usr/local/go/src/math/big/nat.go:170
	}() {
//line /usr/local/go/src/math/big/nat.go:170
		_go_fuzz_dep_.CoverTab[5730]++
							i--
//line /usr/local/go/src/math/big/nat.go:171
		// _ = "end of CoverTab[5730]"
	}
//line /usr/local/go/src/math/big/nat.go:172
	// _ = "end of CoverTab[5719]"
//line /usr/local/go/src/math/big/nat.go:172
	_go_fuzz_dep_.CoverTab[5720]++

						switch {
	case x[i] < y[i]:
//line /usr/local/go/src/math/big/nat.go:175
		_go_fuzz_dep_.CoverTab[5731]++
							r = -1
//line /usr/local/go/src/math/big/nat.go:176
		// _ = "end of CoverTab[5731]"
	case x[i] > y[i]:
//line /usr/local/go/src/math/big/nat.go:177
		_go_fuzz_dep_.CoverTab[5732]++
							r = 1
//line /usr/local/go/src/math/big/nat.go:178
		// _ = "end of CoverTab[5732]"
//line /usr/local/go/src/math/big/nat.go:178
	default:
//line /usr/local/go/src/math/big/nat.go:178
		_go_fuzz_dep_.CoverTab[5733]++
//line /usr/local/go/src/math/big/nat.go:178
		// _ = "end of CoverTab[5733]"
	}
//line /usr/local/go/src/math/big/nat.go:179
	// _ = "end of CoverTab[5720]"
//line /usr/local/go/src/math/big/nat.go:179
	_go_fuzz_dep_.CoverTab[5721]++
						return
//line /usr/local/go/src/math/big/nat.go:180
	// _ = "end of CoverTab[5721]"
}

func (z nat) mulAddWW(x nat, y, r Word) nat {
//line /usr/local/go/src/math/big/nat.go:183
	_go_fuzz_dep_.CoverTab[5734]++
						m := len(x)
						if m == 0 || func() bool {
//line /usr/local/go/src/math/big/nat.go:185
		_go_fuzz_dep_.CoverTab[5736]++
//line /usr/local/go/src/math/big/nat.go:185
		return y == 0
//line /usr/local/go/src/math/big/nat.go:185
		// _ = "end of CoverTab[5736]"
//line /usr/local/go/src/math/big/nat.go:185
	}() {
//line /usr/local/go/src/math/big/nat.go:185
		_go_fuzz_dep_.CoverTab[5737]++
							return z.setWord(r)
//line /usr/local/go/src/math/big/nat.go:186
		// _ = "end of CoverTab[5737]"
	} else {
//line /usr/local/go/src/math/big/nat.go:187
		_go_fuzz_dep_.CoverTab[5738]++
//line /usr/local/go/src/math/big/nat.go:187
		// _ = "end of CoverTab[5738]"
//line /usr/local/go/src/math/big/nat.go:187
	}
//line /usr/local/go/src/math/big/nat.go:187
	// _ = "end of CoverTab[5734]"
//line /usr/local/go/src/math/big/nat.go:187
	_go_fuzz_dep_.CoverTab[5735]++

//line /usr/local/go/src/math/big/nat.go:190
	z = z.make(m + 1)
						z[m] = mulAddVWW(z[0:m], x, y, r)

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:193
	// _ = "end of CoverTab[5735]"
}

// basicMul multiplies x and y and leaves the result in z.
//line /usr/local/go/src/math/big/nat.go:196
// The (non-normalized) result is placed in z[0 : len(x) + len(y)].
//line /usr/local/go/src/math/big/nat.go:198
func basicMul(z, x, y nat) {
//line /usr/local/go/src/math/big/nat.go:198
	_go_fuzz_dep_.CoverTab[5739]++
						z[0 : len(x)+len(y)].clear()
						for i, d := range y {
//line /usr/local/go/src/math/big/nat.go:200
		_go_fuzz_dep_.CoverTab[5740]++
							if d != 0 {
//line /usr/local/go/src/math/big/nat.go:201
			_go_fuzz_dep_.CoverTab[5741]++
								z[len(x)+i] = addMulVVW(z[i:i+len(x)], x, d)
//line /usr/local/go/src/math/big/nat.go:202
			// _ = "end of CoverTab[5741]"
		} else {
//line /usr/local/go/src/math/big/nat.go:203
			_go_fuzz_dep_.CoverTab[5742]++
//line /usr/local/go/src/math/big/nat.go:203
			// _ = "end of CoverTab[5742]"
//line /usr/local/go/src/math/big/nat.go:203
		}
//line /usr/local/go/src/math/big/nat.go:203
		// _ = "end of CoverTab[5740]"
	}
//line /usr/local/go/src/math/big/nat.go:204
	// _ = "end of CoverTab[5739]"
}

// montgomery computes z mod m = x*y*2**(-n*_W) mod m,
//line /usr/local/go/src/math/big/nat.go:207
// assuming k = -1/m mod 2**_W.
//line /usr/local/go/src/math/big/nat.go:207
// z is used for storing the result which is returned;
//line /usr/local/go/src/math/big/nat.go:207
// z must not alias x, y or m.
//line /usr/local/go/src/math/big/nat.go:207
// See Gueron, "Efficient Software Implementations of Modular Exponentiation".
//line /usr/local/go/src/math/big/nat.go:207
// https://eprint.iacr.org/2011/239.pdf
//line /usr/local/go/src/math/big/nat.go:207
// In the terminology of that paper, this is an "Almost Montgomery Multiplication":
//line /usr/local/go/src/math/big/nat.go:207
// x and y are required to satisfy 0 <= z < 2**(n*_W) and then the result
//line /usr/local/go/src/math/big/nat.go:207
// z is guaranteed to satisfy 0 <= z < 2**(n*_W), but it may not be < m.
//line /usr/local/go/src/math/big/nat.go:216
func (z nat) montgomery(x, y, m nat, k Word, n int) nat {
//line /usr/local/go/src/math/big/nat.go:216
	_go_fuzz_dep_.CoverTab[5743]++

//line /usr/local/go/src/math/big/nat.go:221
	if len(x) != n || func() bool {
//line /usr/local/go/src/math/big/nat.go:221
		_go_fuzz_dep_.CoverTab[5747]++
//line /usr/local/go/src/math/big/nat.go:221
		return len(y) != n
//line /usr/local/go/src/math/big/nat.go:221
		// _ = "end of CoverTab[5747]"
//line /usr/local/go/src/math/big/nat.go:221
	}() || func() bool {
//line /usr/local/go/src/math/big/nat.go:221
		_go_fuzz_dep_.CoverTab[5748]++
//line /usr/local/go/src/math/big/nat.go:221
		return len(m) != n
//line /usr/local/go/src/math/big/nat.go:221
		// _ = "end of CoverTab[5748]"
//line /usr/local/go/src/math/big/nat.go:221
	}() {
//line /usr/local/go/src/math/big/nat.go:221
		_go_fuzz_dep_.CoverTab[5749]++
							panic("math/big: mismatched montgomery number lengths")
//line /usr/local/go/src/math/big/nat.go:222
		// _ = "end of CoverTab[5749]"
	} else {
//line /usr/local/go/src/math/big/nat.go:223
		_go_fuzz_dep_.CoverTab[5750]++
//line /usr/local/go/src/math/big/nat.go:223
		// _ = "end of CoverTab[5750]"
//line /usr/local/go/src/math/big/nat.go:223
	}
//line /usr/local/go/src/math/big/nat.go:223
	// _ = "end of CoverTab[5743]"
//line /usr/local/go/src/math/big/nat.go:223
	_go_fuzz_dep_.CoverTab[5744]++
						z = z.make(n * 2)
						z.clear()
						var c Word
						for i := 0; i < n; i++ {
//line /usr/local/go/src/math/big/nat.go:227
		_go_fuzz_dep_.CoverTab[5751]++
							d := y[i]
							c2 := addMulVVW(z[i:n+i], x, d)
							t := z[i] * k
							c3 := addMulVVW(z[i:n+i], m, t)
							cx := c + c2
							cy := cx + c3
							z[n+i] = cy
							if cx < c2 || func() bool {
//line /usr/local/go/src/math/big/nat.go:235
			_go_fuzz_dep_.CoverTab[5752]++
//line /usr/local/go/src/math/big/nat.go:235
			return cy < c3
//line /usr/local/go/src/math/big/nat.go:235
			// _ = "end of CoverTab[5752]"
//line /usr/local/go/src/math/big/nat.go:235
		}() {
//line /usr/local/go/src/math/big/nat.go:235
			_go_fuzz_dep_.CoverTab[5753]++
								c = 1
//line /usr/local/go/src/math/big/nat.go:236
			// _ = "end of CoverTab[5753]"
		} else {
//line /usr/local/go/src/math/big/nat.go:237
			_go_fuzz_dep_.CoverTab[5754]++
								c = 0
//line /usr/local/go/src/math/big/nat.go:238
			// _ = "end of CoverTab[5754]"
		}
//line /usr/local/go/src/math/big/nat.go:239
		// _ = "end of CoverTab[5751]"
	}
//line /usr/local/go/src/math/big/nat.go:240
	// _ = "end of CoverTab[5744]"
//line /usr/local/go/src/math/big/nat.go:240
	_go_fuzz_dep_.CoverTab[5745]++
						if c != 0 {
//line /usr/local/go/src/math/big/nat.go:241
		_go_fuzz_dep_.CoverTab[5755]++
							subVV(z[:n], z[n:], m)
//line /usr/local/go/src/math/big/nat.go:242
		// _ = "end of CoverTab[5755]"
	} else {
//line /usr/local/go/src/math/big/nat.go:243
		_go_fuzz_dep_.CoverTab[5756]++
							copy(z[:n], z[n:])
//line /usr/local/go/src/math/big/nat.go:244
		// _ = "end of CoverTab[5756]"
	}
//line /usr/local/go/src/math/big/nat.go:245
	// _ = "end of CoverTab[5745]"
//line /usr/local/go/src/math/big/nat.go:245
	_go_fuzz_dep_.CoverTab[5746]++
						return z[:n]
//line /usr/local/go/src/math/big/nat.go:246
	// _ = "end of CoverTab[5746]"
}

// Fast version of z[0:n+n>>1].add(z[0:n+n>>1], x[0:n]) w/o bounds checks.
//line /usr/local/go/src/math/big/nat.go:249
// Factored out for readability - do not use outside karatsuba.
//line /usr/local/go/src/math/big/nat.go:251
func karatsubaAdd(z, x nat, n int) {
//line /usr/local/go/src/math/big/nat.go:251
	_go_fuzz_dep_.CoverTab[5757]++
						if c := addVV(z[0:n], z, x); c != 0 {
//line /usr/local/go/src/math/big/nat.go:252
		_go_fuzz_dep_.CoverTab[5758]++
							addVW(z[n:n+n>>1], z[n:], c)
//line /usr/local/go/src/math/big/nat.go:253
		// _ = "end of CoverTab[5758]"
	} else {
//line /usr/local/go/src/math/big/nat.go:254
		_go_fuzz_dep_.CoverTab[5759]++
//line /usr/local/go/src/math/big/nat.go:254
		// _ = "end of CoverTab[5759]"
//line /usr/local/go/src/math/big/nat.go:254
	}
//line /usr/local/go/src/math/big/nat.go:254
	// _ = "end of CoverTab[5757]"
}

// Like karatsubaAdd, but does subtract.
func karatsubaSub(z, x nat, n int) {
//line /usr/local/go/src/math/big/nat.go:258
	_go_fuzz_dep_.CoverTab[5760]++
						if c := subVV(z[0:n], z, x); c != 0 {
//line /usr/local/go/src/math/big/nat.go:259
		_go_fuzz_dep_.CoverTab[5761]++
							subVW(z[n:n+n>>1], z[n:], c)
//line /usr/local/go/src/math/big/nat.go:260
		// _ = "end of CoverTab[5761]"
	} else {
//line /usr/local/go/src/math/big/nat.go:261
		_go_fuzz_dep_.CoverTab[5762]++
//line /usr/local/go/src/math/big/nat.go:261
		// _ = "end of CoverTab[5762]"
//line /usr/local/go/src/math/big/nat.go:261
	}
//line /usr/local/go/src/math/big/nat.go:261
	// _ = "end of CoverTab[5760]"
}

// Operands that are shorter than karatsubaThreshold are multiplied using
//line /usr/local/go/src/math/big/nat.go:264
// "grade school" multiplication; for longer operands the Karatsuba algorithm
//line /usr/local/go/src/math/big/nat.go:264
// is used.
//line /usr/local/go/src/math/big/nat.go:267
var karatsubaThreshold = 40	// computed by calibrate_test.go

// karatsuba multiplies x and y and leaves the result in z.
//line /usr/local/go/src/math/big/nat.go:269
// Both x and y must have the same length n and n must be a
//line /usr/local/go/src/math/big/nat.go:269
// power of 2. The result vector z must have len(z) >= 6*n.
//line /usr/local/go/src/math/big/nat.go:269
// The (non-normalized) result is placed in z[0 : 2*n].
//line /usr/local/go/src/math/big/nat.go:273
func karatsuba(z, x, y nat) {
//line /usr/local/go/src/math/big/nat.go:273
	_go_fuzz_dep_.CoverTab[5763]++
						n := len(y)

//line /usr/local/go/src/math/big/nat.go:279
	if n&1 != 0 || func() bool {
//line /usr/local/go/src/math/big/nat.go:279
		_go_fuzz_dep_.CoverTab[5767]++
//line /usr/local/go/src/math/big/nat.go:279
		return n < karatsubaThreshold
//line /usr/local/go/src/math/big/nat.go:279
		// _ = "end of CoverTab[5767]"
//line /usr/local/go/src/math/big/nat.go:279
	}() || func() bool {
//line /usr/local/go/src/math/big/nat.go:279
		_go_fuzz_dep_.CoverTab[5768]++
//line /usr/local/go/src/math/big/nat.go:279
		return n < 2
//line /usr/local/go/src/math/big/nat.go:279
		// _ = "end of CoverTab[5768]"
//line /usr/local/go/src/math/big/nat.go:279
	}() {
//line /usr/local/go/src/math/big/nat.go:279
		_go_fuzz_dep_.CoverTab[5769]++
							basicMul(z, x, y)
							return
//line /usr/local/go/src/math/big/nat.go:281
		// _ = "end of CoverTab[5769]"
	} else {
//line /usr/local/go/src/math/big/nat.go:282
		_go_fuzz_dep_.CoverTab[5770]++
//line /usr/local/go/src/math/big/nat.go:282
		// _ = "end of CoverTab[5770]"
//line /usr/local/go/src/math/big/nat.go:282
	}
//line /usr/local/go/src/math/big/nat.go:282
	// _ = "end of CoverTab[5763]"
//line /usr/local/go/src/math/big/nat.go:282
	_go_fuzz_dep_.CoverTab[5764]++

//line /usr/local/go/src/math/big/nat.go:309
	n2 := n >> 1
						x1, x0 := x[n2:], x[0:n2]
						y1, y0 := y[n2:], y[0:n2]

//line /usr/local/go/src/math/big/nat.go:323
	karatsuba(z, x0, y0)
						karatsuba(z[n:], x1, y1)

//line /usr/local/go/src/math/big/nat.go:327
	s := 1
	xd := z[2*n : 2*n+n2]
	if subVV(xd, x1, x0) != 0 {
//line /usr/local/go/src/math/big/nat.go:329
		_go_fuzz_dep_.CoverTab[5771]++
							s = -s
							subVV(xd, x0, x1)
//line /usr/local/go/src/math/big/nat.go:331
		// _ = "end of CoverTab[5771]"
	} else {
//line /usr/local/go/src/math/big/nat.go:332
		_go_fuzz_dep_.CoverTab[5772]++
//line /usr/local/go/src/math/big/nat.go:332
		// _ = "end of CoverTab[5772]"
//line /usr/local/go/src/math/big/nat.go:332
	}
//line /usr/local/go/src/math/big/nat.go:332
	// _ = "end of CoverTab[5764]"
//line /usr/local/go/src/math/big/nat.go:332
	_go_fuzz_dep_.CoverTab[5765]++

//line /usr/local/go/src/math/big/nat.go:335
	yd := z[2*n+n2 : 3*n]
	if subVV(yd, y0, y1) != 0 {
//line /usr/local/go/src/math/big/nat.go:336
		_go_fuzz_dep_.CoverTab[5773]++
							s = -s
							subVV(yd, y1, y0)
//line /usr/local/go/src/math/big/nat.go:338
		// _ = "end of CoverTab[5773]"
	} else {
//line /usr/local/go/src/math/big/nat.go:339
		_go_fuzz_dep_.CoverTab[5774]++
//line /usr/local/go/src/math/big/nat.go:339
		// _ = "end of CoverTab[5774]"
//line /usr/local/go/src/math/big/nat.go:339
	}
//line /usr/local/go/src/math/big/nat.go:339
	// _ = "end of CoverTab[5765]"
//line /usr/local/go/src/math/big/nat.go:339
	_go_fuzz_dep_.CoverTab[5766]++

//line /usr/local/go/src/math/big/nat.go:343
	p := z[n*3:]
						karatsuba(p, xd, yd)

//line /usr/local/go/src/math/big/nat.go:348
	r := z[n*4:]
						copy(r, z[:n*2])

//line /usr/local/go/src/math/big/nat.go:359
	karatsubaAdd(z[n2:], r, n)
	karatsubaAdd(z[n2:], r[n:], n)
	if s > 0 {
//line /usr/local/go/src/math/big/nat.go:361
		_go_fuzz_dep_.CoverTab[5775]++
							karatsubaAdd(z[n2:], p, n)
//line /usr/local/go/src/math/big/nat.go:362
		// _ = "end of CoverTab[5775]"
	} else {
//line /usr/local/go/src/math/big/nat.go:363
		_go_fuzz_dep_.CoverTab[5776]++
							karatsubaSub(z[n2:], p, n)
//line /usr/local/go/src/math/big/nat.go:364
		// _ = "end of CoverTab[5776]"
	}
//line /usr/local/go/src/math/big/nat.go:365
	// _ = "end of CoverTab[5766]"
}

// alias reports whether x and y share the same base array.
//line /usr/local/go/src/math/big/nat.go:368
//
//line /usr/local/go/src/math/big/nat.go:368
// Note: alias assumes that the capacity of underlying arrays
//line /usr/local/go/src/math/big/nat.go:368
// is never changed for nat values; i.e. that there are
//line /usr/local/go/src/math/big/nat.go:368
// no 3-operand slice expressions in this code (or worse,
//line /usr/local/go/src/math/big/nat.go:368
// reflect-based operations to the same effect).
//line /usr/local/go/src/math/big/nat.go:374
func alias(x, y nat) bool {
//line /usr/local/go/src/math/big/nat.go:374
	_go_fuzz_dep_.CoverTab[5777]++
						return cap(x) > 0 && func() bool {
//line /usr/local/go/src/math/big/nat.go:375
		_go_fuzz_dep_.CoverTab[5778]++
//line /usr/local/go/src/math/big/nat.go:375
		return cap(y) > 0
//line /usr/local/go/src/math/big/nat.go:375
		// _ = "end of CoverTab[5778]"
//line /usr/local/go/src/math/big/nat.go:375
	}() && func() bool {
//line /usr/local/go/src/math/big/nat.go:375
		_go_fuzz_dep_.CoverTab[5779]++
//line /usr/local/go/src/math/big/nat.go:375
		return &x[0:cap(x)][cap(x)-1] == &y[0:cap(y)][cap(y)-1]
//line /usr/local/go/src/math/big/nat.go:375
		// _ = "end of CoverTab[5779]"
//line /usr/local/go/src/math/big/nat.go:375
	}()
//line /usr/local/go/src/math/big/nat.go:375
	// _ = "end of CoverTab[5777]"
}

// addAt implements z += x<<(_W*i); z must be long enough.
//line /usr/local/go/src/math/big/nat.go:378
// (we don't use nat.add because we need z to stay the same
//line /usr/local/go/src/math/big/nat.go:378
// slice, and we don't need to normalize z after each addition)
//line /usr/local/go/src/math/big/nat.go:381
func addAt(z, x nat, i int) {
//line /usr/local/go/src/math/big/nat.go:381
	_go_fuzz_dep_.CoverTab[5780]++
						if n := len(x); n > 0 {
//line /usr/local/go/src/math/big/nat.go:382
		_go_fuzz_dep_.CoverTab[5781]++
							if c := addVV(z[i:i+n], z[i:], x); c != 0 {
//line /usr/local/go/src/math/big/nat.go:383
			_go_fuzz_dep_.CoverTab[5782]++
								j := i + n
								if j < len(z) {
//line /usr/local/go/src/math/big/nat.go:385
				_go_fuzz_dep_.CoverTab[5783]++
									addVW(z[j:], z[j:], c)
//line /usr/local/go/src/math/big/nat.go:386
				// _ = "end of CoverTab[5783]"
			} else {
//line /usr/local/go/src/math/big/nat.go:387
				_go_fuzz_dep_.CoverTab[5784]++
//line /usr/local/go/src/math/big/nat.go:387
				// _ = "end of CoverTab[5784]"
//line /usr/local/go/src/math/big/nat.go:387
			}
//line /usr/local/go/src/math/big/nat.go:387
			// _ = "end of CoverTab[5782]"
		} else {
//line /usr/local/go/src/math/big/nat.go:388
			_go_fuzz_dep_.CoverTab[5785]++
//line /usr/local/go/src/math/big/nat.go:388
			// _ = "end of CoverTab[5785]"
//line /usr/local/go/src/math/big/nat.go:388
		}
//line /usr/local/go/src/math/big/nat.go:388
		// _ = "end of CoverTab[5781]"
	} else {
//line /usr/local/go/src/math/big/nat.go:389
		_go_fuzz_dep_.CoverTab[5786]++
//line /usr/local/go/src/math/big/nat.go:389
		// _ = "end of CoverTab[5786]"
//line /usr/local/go/src/math/big/nat.go:389
	}
//line /usr/local/go/src/math/big/nat.go:389
	// _ = "end of CoverTab[5780]"
}

func max(x, y int) int {
//line /usr/local/go/src/math/big/nat.go:392
	_go_fuzz_dep_.CoverTab[5787]++
						if x > y {
//line /usr/local/go/src/math/big/nat.go:393
		_go_fuzz_dep_.CoverTab[5789]++
							return x
//line /usr/local/go/src/math/big/nat.go:394
		// _ = "end of CoverTab[5789]"
	} else {
//line /usr/local/go/src/math/big/nat.go:395
		_go_fuzz_dep_.CoverTab[5790]++
//line /usr/local/go/src/math/big/nat.go:395
		// _ = "end of CoverTab[5790]"
//line /usr/local/go/src/math/big/nat.go:395
	}
//line /usr/local/go/src/math/big/nat.go:395
	// _ = "end of CoverTab[5787]"
//line /usr/local/go/src/math/big/nat.go:395
	_go_fuzz_dep_.CoverTab[5788]++
						return y
//line /usr/local/go/src/math/big/nat.go:396
	// _ = "end of CoverTab[5788]"
}

// karatsubaLen computes an approximation to the maximum k <= n such that
//line /usr/local/go/src/math/big/nat.go:399
// k = p<<i for a number p <= threshold and an i >= 0. Thus, the
//line /usr/local/go/src/math/big/nat.go:399
// result is the largest number that can be divided repeatedly by 2 before
//line /usr/local/go/src/math/big/nat.go:399
// becoming about the value of threshold.
//line /usr/local/go/src/math/big/nat.go:403
func karatsubaLen(n, threshold int) int {
//line /usr/local/go/src/math/big/nat.go:403
	_go_fuzz_dep_.CoverTab[5791]++
						i := uint(0)
						for n > threshold {
//line /usr/local/go/src/math/big/nat.go:405
		_go_fuzz_dep_.CoverTab[5793]++
							n >>= 1
							i++
//line /usr/local/go/src/math/big/nat.go:407
		// _ = "end of CoverTab[5793]"
	}
//line /usr/local/go/src/math/big/nat.go:408
	// _ = "end of CoverTab[5791]"
//line /usr/local/go/src/math/big/nat.go:408
	_go_fuzz_dep_.CoverTab[5792]++
						return n << i
//line /usr/local/go/src/math/big/nat.go:409
	// _ = "end of CoverTab[5792]"
}

func (z nat) mul(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:412
	_go_fuzz_dep_.CoverTab[5794]++
						m := len(x)
						n := len(y)

						switch {
	case m < n:
//line /usr/local/go/src/math/big/nat.go:417
		_go_fuzz_dep_.CoverTab[5799]++
							return z.mul(y, x)
//line /usr/local/go/src/math/big/nat.go:418
		// _ = "end of CoverTab[5799]"
	case m == 0 || func() bool {
//line /usr/local/go/src/math/big/nat.go:419
		_go_fuzz_dep_.CoverTab[5803]++
//line /usr/local/go/src/math/big/nat.go:419
		return n == 0
//line /usr/local/go/src/math/big/nat.go:419
		// _ = "end of CoverTab[5803]"
//line /usr/local/go/src/math/big/nat.go:419
	}():
//line /usr/local/go/src/math/big/nat.go:419
		_go_fuzz_dep_.CoverTab[5800]++
							return z[:0]
//line /usr/local/go/src/math/big/nat.go:420
		// _ = "end of CoverTab[5800]"
	case n == 1:
//line /usr/local/go/src/math/big/nat.go:421
		_go_fuzz_dep_.CoverTab[5801]++
							return z.mulAddWW(x, y[0], 0)
//line /usr/local/go/src/math/big/nat.go:422
		// _ = "end of CoverTab[5801]"
//line /usr/local/go/src/math/big/nat.go:422
	default:
//line /usr/local/go/src/math/big/nat.go:422
		_go_fuzz_dep_.CoverTab[5802]++
//line /usr/local/go/src/math/big/nat.go:422
		// _ = "end of CoverTab[5802]"
	}
//line /usr/local/go/src/math/big/nat.go:423
	// _ = "end of CoverTab[5794]"
//line /usr/local/go/src/math/big/nat.go:423
	_go_fuzz_dep_.CoverTab[5795]++

//line /usr/local/go/src/math/big/nat.go:427
	if alias(z, x) || func() bool {
//line /usr/local/go/src/math/big/nat.go:427
		_go_fuzz_dep_.CoverTab[5804]++
//line /usr/local/go/src/math/big/nat.go:427
		return alias(z, y)
//line /usr/local/go/src/math/big/nat.go:427
		// _ = "end of CoverTab[5804]"
//line /usr/local/go/src/math/big/nat.go:427
	}() {
//line /usr/local/go/src/math/big/nat.go:427
		_go_fuzz_dep_.CoverTab[5805]++
							z = nil
//line /usr/local/go/src/math/big/nat.go:428
		// _ = "end of CoverTab[5805]"
	} else {
//line /usr/local/go/src/math/big/nat.go:429
		_go_fuzz_dep_.CoverTab[5806]++
//line /usr/local/go/src/math/big/nat.go:429
		// _ = "end of CoverTab[5806]"
//line /usr/local/go/src/math/big/nat.go:429
	}
//line /usr/local/go/src/math/big/nat.go:429
	// _ = "end of CoverTab[5795]"
//line /usr/local/go/src/math/big/nat.go:429
	_go_fuzz_dep_.CoverTab[5796]++

//line /usr/local/go/src/math/big/nat.go:432
	if n < karatsubaThreshold {
//line /usr/local/go/src/math/big/nat.go:432
		_go_fuzz_dep_.CoverTab[5807]++
							z = z.make(m + n)
							basicMul(z, x, y)
							return z.norm()
//line /usr/local/go/src/math/big/nat.go:435
		// _ = "end of CoverTab[5807]"
	} else {
//line /usr/local/go/src/math/big/nat.go:436
		_go_fuzz_dep_.CoverTab[5808]++
//line /usr/local/go/src/math/big/nat.go:436
		// _ = "end of CoverTab[5808]"
//line /usr/local/go/src/math/big/nat.go:436
	}
//line /usr/local/go/src/math/big/nat.go:436
	// _ = "end of CoverTab[5796]"
//line /usr/local/go/src/math/big/nat.go:436
	_go_fuzz_dep_.CoverTab[5797]++

//line /usr/local/go/src/math/big/nat.go:445
	k := karatsubaLen(n, karatsubaThreshold)

//line /usr/local/go/src/math/big/nat.go:449
	x0 := x[0:k]
						y0 := y[0:k]
						z = z.make(max(6*k, m+n))
						karatsuba(z, x0, y0)
						z = z[0 : m+n]
						z[2*k:].clear()

//line /usr/local/go/src/math/big/nat.go:469
	if k < n || func() bool {
//line /usr/local/go/src/math/big/nat.go:469
		_go_fuzz_dep_.CoverTab[5809]++
//line /usr/local/go/src/math/big/nat.go:469
		return m != n
//line /usr/local/go/src/math/big/nat.go:469
		// _ = "end of CoverTab[5809]"
//line /usr/local/go/src/math/big/nat.go:469
	}() {
//line /usr/local/go/src/math/big/nat.go:469
		_go_fuzz_dep_.CoverTab[5810]++
							tp := getNat(3 * k)
							t := *tp

//line /usr/local/go/src/math/big/nat.go:474
		x0 := x0.norm()
							y1 := y[k:]
							t = t.mul(x0, y1)
							addAt(z, t, k)

//line /usr/local/go/src/math/big/nat.go:480
		y0 := y0.norm()
		for i := k; i < len(x); i += k {
//line /usr/local/go/src/math/big/nat.go:481
			_go_fuzz_dep_.CoverTab[5812]++
								xi := x[i:]
								if len(xi) > k {
//line /usr/local/go/src/math/big/nat.go:483
				_go_fuzz_dep_.CoverTab[5814]++
									xi = xi[:k]
//line /usr/local/go/src/math/big/nat.go:484
				// _ = "end of CoverTab[5814]"
			} else {
//line /usr/local/go/src/math/big/nat.go:485
				_go_fuzz_dep_.CoverTab[5815]++
//line /usr/local/go/src/math/big/nat.go:485
				// _ = "end of CoverTab[5815]"
//line /usr/local/go/src/math/big/nat.go:485
			}
//line /usr/local/go/src/math/big/nat.go:485
			// _ = "end of CoverTab[5812]"
//line /usr/local/go/src/math/big/nat.go:485
			_go_fuzz_dep_.CoverTab[5813]++
								xi = xi.norm()
								t = t.mul(xi, y0)
								addAt(z, t, i)
								t = t.mul(xi, y1)
								addAt(z, t, i+k)
//line /usr/local/go/src/math/big/nat.go:490
			// _ = "end of CoverTab[5813]"
		}
//line /usr/local/go/src/math/big/nat.go:491
		// _ = "end of CoverTab[5810]"
//line /usr/local/go/src/math/big/nat.go:491
		_go_fuzz_dep_.CoverTab[5811]++

							putNat(tp)
//line /usr/local/go/src/math/big/nat.go:493
		// _ = "end of CoverTab[5811]"
	} else {
//line /usr/local/go/src/math/big/nat.go:494
		_go_fuzz_dep_.CoverTab[5816]++
//line /usr/local/go/src/math/big/nat.go:494
		// _ = "end of CoverTab[5816]"
//line /usr/local/go/src/math/big/nat.go:494
	}
//line /usr/local/go/src/math/big/nat.go:494
	// _ = "end of CoverTab[5797]"
//line /usr/local/go/src/math/big/nat.go:494
	_go_fuzz_dep_.CoverTab[5798]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:496
	// _ = "end of CoverTab[5798]"
}

// basicSqr sets z = x*x and is asymptotically faster than basicMul
//line /usr/local/go/src/math/big/nat.go:499
// by about a factor of 2, but slower for small arguments due to overhead.
//line /usr/local/go/src/math/big/nat.go:499
// Requirements: len(x) > 0, len(z) == 2*len(x)
//line /usr/local/go/src/math/big/nat.go:499
// The (non-normalized) result is placed in z.
//line /usr/local/go/src/math/big/nat.go:503
func basicSqr(z, x nat) {
//line /usr/local/go/src/math/big/nat.go:503
	_go_fuzz_dep_.CoverTab[5817]++
						n := len(x)
						tp := getNat(2 * n)
						t := *tp
						t.clear()
						z[1], z[0] = mulWW(x[0], x[0])
						for i := 1; i < n; i++ {
//line /usr/local/go/src/math/big/nat.go:509
		_go_fuzz_dep_.CoverTab[5819]++
							d := x[i]

							z[2*i+1], z[2*i] = mulWW(d, d)

							t[2*i] = addMulVVW(t[i:2*i], x[0:i], d)
//line /usr/local/go/src/math/big/nat.go:514
		// _ = "end of CoverTab[5819]"
	}
//line /usr/local/go/src/math/big/nat.go:515
	// _ = "end of CoverTab[5817]"
//line /usr/local/go/src/math/big/nat.go:515
	_go_fuzz_dep_.CoverTab[5818]++
						t[2*n-1] = shlVU(t[1:2*n-1], t[1:2*n-1], 1)
						addVV(z, z, t)
						putNat(tp)
//line /usr/local/go/src/math/big/nat.go:518
	// _ = "end of CoverTab[5818]"
}

// karatsubaSqr squares x and leaves the result in z.
//line /usr/local/go/src/math/big/nat.go:521
// len(x) must be a power of 2 and len(z) >= 6*len(x).
//line /usr/local/go/src/math/big/nat.go:521
// The (non-normalized) result is placed in z[0 : 2*len(x)].
//line /usr/local/go/src/math/big/nat.go:521
//
//line /usr/local/go/src/math/big/nat.go:521
// The algorithm and the layout of z are the same as for karatsuba.
//line /usr/local/go/src/math/big/nat.go:526
func karatsubaSqr(z, x nat) {
//line /usr/local/go/src/math/big/nat.go:526
	_go_fuzz_dep_.CoverTab[5820]++
						n := len(x)

						if n&1 != 0 || func() bool {
//line /usr/local/go/src/math/big/nat.go:529
		_go_fuzz_dep_.CoverTab[5823]++
//line /usr/local/go/src/math/big/nat.go:529
		return n < karatsubaSqrThreshold
//line /usr/local/go/src/math/big/nat.go:529
		// _ = "end of CoverTab[5823]"
//line /usr/local/go/src/math/big/nat.go:529
	}() || func() bool {
//line /usr/local/go/src/math/big/nat.go:529
		_go_fuzz_dep_.CoverTab[5824]++
//line /usr/local/go/src/math/big/nat.go:529
		return n < 2
//line /usr/local/go/src/math/big/nat.go:529
		// _ = "end of CoverTab[5824]"
//line /usr/local/go/src/math/big/nat.go:529
	}() {
//line /usr/local/go/src/math/big/nat.go:529
		_go_fuzz_dep_.CoverTab[5825]++
							basicSqr(z[:2*n], x)
							return
//line /usr/local/go/src/math/big/nat.go:531
		// _ = "end of CoverTab[5825]"
	} else {
//line /usr/local/go/src/math/big/nat.go:532
		_go_fuzz_dep_.CoverTab[5826]++
//line /usr/local/go/src/math/big/nat.go:532
		// _ = "end of CoverTab[5826]"
//line /usr/local/go/src/math/big/nat.go:532
	}
//line /usr/local/go/src/math/big/nat.go:532
	// _ = "end of CoverTab[5820]"
//line /usr/local/go/src/math/big/nat.go:532
	_go_fuzz_dep_.CoverTab[5821]++

						n2 := n >> 1
						x1, x0 := x[n2:], x[0:n2]

						karatsubaSqr(z, x0)
						karatsubaSqr(z[n:], x1)

//line /usr/local/go/src/math/big/nat.go:541
	xd := z[2*n : 2*n+n2]
	if subVV(xd, x1, x0) != 0 {
//line /usr/local/go/src/math/big/nat.go:542
		_go_fuzz_dep_.CoverTab[5827]++
							subVV(xd, x0, x1)
//line /usr/local/go/src/math/big/nat.go:543
		// _ = "end of CoverTab[5827]"
	} else {
//line /usr/local/go/src/math/big/nat.go:544
		_go_fuzz_dep_.CoverTab[5828]++
//line /usr/local/go/src/math/big/nat.go:544
		// _ = "end of CoverTab[5828]"
//line /usr/local/go/src/math/big/nat.go:544
	}
//line /usr/local/go/src/math/big/nat.go:544
	// _ = "end of CoverTab[5821]"
//line /usr/local/go/src/math/big/nat.go:544
	_go_fuzz_dep_.CoverTab[5822]++

						p := z[n*3:]
						karatsubaSqr(p, xd)

						r := z[n*4:]
						copy(r, z[:n*2])

						karatsubaAdd(z[n2:], r, n)
						karatsubaAdd(z[n2:], r[n:], n)
						karatsubaSub(z[n2:], p, n)
//line /usr/local/go/src/math/big/nat.go:554
	// _ = "end of CoverTab[5822]"
}

// Operands that are shorter than basicSqrThreshold are squared using
//line /usr/local/go/src/math/big/nat.go:557
// "grade school" multiplication; for operands longer than karatsubaSqrThreshold
//line /usr/local/go/src/math/big/nat.go:557
// we use the Karatsuba algorithm optimized for x == y.
//line /usr/local/go/src/math/big/nat.go:560
var basicSqrThreshold = 20	// computed by calibrate_test.go
var karatsubaSqrThreshold = 260						// computed by calibrate_test.go

// z = x*x
func (z nat) sqr(x nat) nat {
//line /usr/local/go/src/math/big/nat.go:564
	_go_fuzz_dep_.CoverTab[5829]++
						n := len(x)
						switch {
	case n == 0:
//line /usr/local/go/src/math/big/nat.go:567
		_go_fuzz_dep_.CoverTab[5835]++
							return z[:0]
//line /usr/local/go/src/math/big/nat.go:568
		// _ = "end of CoverTab[5835]"
	case n == 1:
//line /usr/local/go/src/math/big/nat.go:569
		_go_fuzz_dep_.CoverTab[5836]++
							d := x[0]
							z = z.make(2)
							z[1], z[0] = mulWW(d, d)
							return z.norm()
//line /usr/local/go/src/math/big/nat.go:573
		// _ = "end of CoverTab[5836]"
//line /usr/local/go/src/math/big/nat.go:573
	default:
//line /usr/local/go/src/math/big/nat.go:573
		_go_fuzz_dep_.CoverTab[5837]++
//line /usr/local/go/src/math/big/nat.go:573
		// _ = "end of CoverTab[5837]"
	}
//line /usr/local/go/src/math/big/nat.go:574
	// _ = "end of CoverTab[5829]"
//line /usr/local/go/src/math/big/nat.go:574
	_go_fuzz_dep_.CoverTab[5830]++

						if alias(z, x) {
//line /usr/local/go/src/math/big/nat.go:576
		_go_fuzz_dep_.CoverTab[5838]++
							z = nil
//line /usr/local/go/src/math/big/nat.go:577
		// _ = "end of CoverTab[5838]"
	} else {
//line /usr/local/go/src/math/big/nat.go:578
		_go_fuzz_dep_.CoverTab[5839]++
//line /usr/local/go/src/math/big/nat.go:578
		// _ = "end of CoverTab[5839]"
//line /usr/local/go/src/math/big/nat.go:578
	}
//line /usr/local/go/src/math/big/nat.go:578
	// _ = "end of CoverTab[5830]"
//line /usr/local/go/src/math/big/nat.go:578
	_go_fuzz_dep_.CoverTab[5831]++

						if n < basicSqrThreshold {
//line /usr/local/go/src/math/big/nat.go:580
		_go_fuzz_dep_.CoverTab[5840]++
							z = z.make(2 * n)
							basicMul(z, x, x)
							return z.norm()
//line /usr/local/go/src/math/big/nat.go:583
		// _ = "end of CoverTab[5840]"
	} else {
//line /usr/local/go/src/math/big/nat.go:584
		_go_fuzz_dep_.CoverTab[5841]++
//line /usr/local/go/src/math/big/nat.go:584
		// _ = "end of CoverTab[5841]"
//line /usr/local/go/src/math/big/nat.go:584
	}
//line /usr/local/go/src/math/big/nat.go:584
	// _ = "end of CoverTab[5831]"
//line /usr/local/go/src/math/big/nat.go:584
	_go_fuzz_dep_.CoverTab[5832]++
						if n < karatsubaSqrThreshold {
//line /usr/local/go/src/math/big/nat.go:585
		_go_fuzz_dep_.CoverTab[5842]++
							z = z.make(2 * n)
							basicSqr(z, x)
							return z.norm()
//line /usr/local/go/src/math/big/nat.go:588
		// _ = "end of CoverTab[5842]"
	} else {
//line /usr/local/go/src/math/big/nat.go:589
		_go_fuzz_dep_.CoverTab[5843]++
//line /usr/local/go/src/math/big/nat.go:589
		// _ = "end of CoverTab[5843]"
//line /usr/local/go/src/math/big/nat.go:589
	}
//line /usr/local/go/src/math/big/nat.go:589
	// _ = "end of CoverTab[5832]"
//line /usr/local/go/src/math/big/nat.go:589
	_go_fuzz_dep_.CoverTab[5833]++

//line /usr/local/go/src/math/big/nat.go:596
	k := karatsubaLen(n, karatsubaSqrThreshold)

	x0 := x[0:k]
	z = z.make(max(6*k, 2*n))
	karatsubaSqr(z, x0)
	z = z[0 : 2*n]
	z[2*k:].clear()

	if k < n {
//line /usr/local/go/src/math/big/nat.go:604
		_go_fuzz_dep_.CoverTab[5844]++
							tp := getNat(2 * k)
							t := *tp
							x0 := x0.norm()
							x1 := x[k:]
							t = t.mul(x0, x1)
							addAt(z, t, k)
							addAt(z, t, k)
							t = t.sqr(x1)
							addAt(z, t, 2*k)
							putNat(tp)
//line /usr/local/go/src/math/big/nat.go:614
		// _ = "end of CoverTab[5844]"
	} else {
//line /usr/local/go/src/math/big/nat.go:615
		_go_fuzz_dep_.CoverTab[5845]++
//line /usr/local/go/src/math/big/nat.go:615
		// _ = "end of CoverTab[5845]"
//line /usr/local/go/src/math/big/nat.go:615
	}
//line /usr/local/go/src/math/big/nat.go:615
	// _ = "end of CoverTab[5833]"
//line /usr/local/go/src/math/big/nat.go:615
	_go_fuzz_dep_.CoverTab[5834]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:617
	// _ = "end of CoverTab[5834]"
}

// mulRange computes the product of all the unsigned integers in the
//line /usr/local/go/src/math/big/nat.go:620
// range [a, b] inclusively. If a > b (empty range), the result is 1.
//line /usr/local/go/src/math/big/nat.go:622
func (z nat) mulRange(a, b uint64) nat {
//line /usr/local/go/src/math/big/nat.go:622
	_go_fuzz_dep_.CoverTab[5846]++
						switch {
	case a == 0:
//line /usr/local/go/src/math/big/nat.go:624
		_go_fuzz_dep_.CoverTab[5848]++

							return z.setUint64(0)
//line /usr/local/go/src/math/big/nat.go:626
		// _ = "end of CoverTab[5848]"
	case a > b:
//line /usr/local/go/src/math/big/nat.go:627
		_go_fuzz_dep_.CoverTab[5849]++
							return z.setUint64(1)
//line /usr/local/go/src/math/big/nat.go:628
		// _ = "end of CoverTab[5849]"
	case a == b:
//line /usr/local/go/src/math/big/nat.go:629
		_go_fuzz_dep_.CoverTab[5850]++
							return z.setUint64(a)
//line /usr/local/go/src/math/big/nat.go:630
		// _ = "end of CoverTab[5850]"
	case a+1 == b:
//line /usr/local/go/src/math/big/nat.go:631
		_go_fuzz_dep_.CoverTab[5851]++
							return z.mul(nat(nil).setUint64(a), nat(nil).setUint64(b))
//line /usr/local/go/src/math/big/nat.go:632
		// _ = "end of CoverTab[5851]"
//line /usr/local/go/src/math/big/nat.go:632
	default:
//line /usr/local/go/src/math/big/nat.go:632
		_go_fuzz_dep_.CoverTab[5852]++
//line /usr/local/go/src/math/big/nat.go:632
		// _ = "end of CoverTab[5852]"
	}
//line /usr/local/go/src/math/big/nat.go:633
	// _ = "end of CoverTab[5846]"
//line /usr/local/go/src/math/big/nat.go:633
	_go_fuzz_dep_.CoverTab[5847]++
						m := (a + b) / 2
						return z.mul(nat(nil).mulRange(a, m), nat(nil).mulRange(m+1, b))
//line /usr/local/go/src/math/big/nat.go:635
	// _ = "end of CoverTab[5847]"
}

// getNat returns a *nat of len n. The contents may not be zero.
//line /usr/local/go/src/math/big/nat.go:638
// The pool holds *nat to avoid allocation when converting to interface{}.
//line /usr/local/go/src/math/big/nat.go:640
func getNat(n int) *nat {
//line /usr/local/go/src/math/big/nat.go:640
	_go_fuzz_dep_.CoverTab[5853]++
						var z *nat
						if v := natPool.Get(); v != nil {
//line /usr/local/go/src/math/big/nat.go:642
		_go_fuzz_dep_.CoverTab[5857]++
							z = v.(*nat)
//line /usr/local/go/src/math/big/nat.go:643
		// _ = "end of CoverTab[5857]"
	} else {
//line /usr/local/go/src/math/big/nat.go:644
		_go_fuzz_dep_.CoverTab[5858]++
//line /usr/local/go/src/math/big/nat.go:644
		// _ = "end of CoverTab[5858]"
//line /usr/local/go/src/math/big/nat.go:644
	}
//line /usr/local/go/src/math/big/nat.go:644
	// _ = "end of CoverTab[5853]"
//line /usr/local/go/src/math/big/nat.go:644
	_go_fuzz_dep_.CoverTab[5854]++
						if z == nil {
//line /usr/local/go/src/math/big/nat.go:645
		_go_fuzz_dep_.CoverTab[5859]++
							z = new(nat)
//line /usr/local/go/src/math/big/nat.go:646
		// _ = "end of CoverTab[5859]"
	} else {
//line /usr/local/go/src/math/big/nat.go:647
		_go_fuzz_dep_.CoverTab[5860]++
//line /usr/local/go/src/math/big/nat.go:647
		// _ = "end of CoverTab[5860]"
//line /usr/local/go/src/math/big/nat.go:647
	}
//line /usr/local/go/src/math/big/nat.go:647
	// _ = "end of CoverTab[5854]"
//line /usr/local/go/src/math/big/nat.go:647
	_go_fuzz_dep_.CoverTab[5855]++
						*z = z.make(n)
						if n > 0 {
//line /usr/local/go/src/math/big/nat.go:649
		_go_fuzz_dep_.CoverTab[5861]++
							(*z)[0] = 0xfedcb
//line /usr/local/go/src/math/big/nat.go:650
		// _ = "end of CoverTab[5861]"
	} else {
//line /usr/local/go/src/math/big/nat.go:651
		_go_fuzz_dep_.CoverTab[5862]++
//line /usr/local/go/src/math/big/nat.go:651
		// _ = "end of CoverTab[5862]"
//line /usr/local/go/src/math/big/nat.go:651
	}
//line /usr/local/go/src/math/big/nat.go:651
	// _ = "end of CoverTab[5855]"
//line /usr/local/go/src/math/big/nat.go:651
	_go_fuzz_dep_.CoverTab[5856]++
						return z
//line /usr/local/go/src/math/big/nat.go:652
	// _ = "end of CoverTab[5856]"
}

func putNat(x *nat) {
//line /usr/local/go/src/math/big/nat.go:655
	_go_fuzz_dep_.CoverTab[5863]++
						natPool.Put(x)
//line /usr/local/go/src/math/big/nat.go:656
	// _ = "end of CoverTab[5863]"
}

var natPool sync.Pool

// bitLen returns the length of x in bits.
//line /usr/local/go/src/math/big/nat.go:661
// Unlike most methods, it works even if x is not normalized.
//line /usr/local/go/src/math/big/nat.go:663
func (x nat) bitLen() int {
//line /usr/local/go/src/math/big/nat.go:663
	_go_fuzz_dep_.CoverTab[5864]++

//line /usr/local/go/src/math/big/nat.go:667
	if i := len(x) - 1; i >= 0 {
//line /usr/local/go/src/math/big/nat.go:667
		_go_fuzz_dep_.CoverTab[5866]++

//line /usr/local/go/src/math/big/nat.go:671
		top := uint(x[i])
							top |= top >> 1
							top |= top >> 2
							top |= top >> 4
							top |= top >> 8
							top |= top >> 16
							top |= top >> 16 >> 16
							return i*_W + bits.Len(top)
//line /usr/local/go/src/math/big/nat.go:678
		// _ = "end of CoverTab[5866]"
	} else {
//line /usr/local/go/src/math/big/nat.go:679
		_go_fuzz_dep_.CoverTab[5867]++
//line /usr/local/go/src/math/big/nat.go:679
		// _ = "end of CoverTab[5867]"
//line /usr/local/go/src/math/big/nat.go:679
	}
//line /usr/local/go/src/math/big/nat.go:679
	// _ = "end of CoverTab[5864]"
//line /usr/local/go/src/math/big/nat.go:679
	_go_fuzz_dep_.CoverTab[5865]++
						return 0
//line /usr/local/go/src/math/big/nat.go:680
	// _ = "end of CoverTab[5865]"
}

// trailingZeroBits returns the number of consecutive least significant zero
//line /usr/local/go/src/math/big/nat.go:683
// bits of x.
//line /usr/local/go/src/math/big/nat.go:685
func (x nat) trailingZeroBits() uint {
//line /usr/local/go/src/math/big/nat.go:685
	_go_fuzz_dep_.CoverTab[5868]++
						if len(x) == 0 {
//line /usr/local/go/src/math/big/nat.go:686
		_go_fuzz_dep_.CoverTab[5871]++
							return 0
//line /usr/local/go/src/math/big/nat.go:687
		// _ = "end of CoverTab[5871]"
	} else {
//line /usr/local/go/src/math/big/nat.go:688
		_go_fuzz_dep_.CoverTab[5872]++
//line /usr/local/go/src/math/big/nat.go:688
		// _ = "end of CoverTab[5872]"
//line /usr/local/go/src/math/big/nat.go:688
	}
//line /usr/local/go/src/math/big/nat.go:688
	// _ = "end of CoverTab[5868]"
//line /usr/local/go/src/math/big/nat.go:688
	_go_fuzz_dep_.CoverTab[5869]++
						var i uint
						for x[i] == 0 {
//line /usr/local/go/src/math/big/nat.go:690
		_go_fuzz_dep_.CoverTab[5873]++
							i++
//line /usr/local/go/src/math/big/nat.go:691
		// _ = "end of CoverTab[5873]"
	}
//line /usr/local/go/src/math/big/nat.go:692
	// _ = "end of CoverTab[5869]"
//line /usr/local/go/src/math/big/nat.go:692
	_go_fuzz_dep_.CoverTab[5870]++

						return i*_W + uint(bits.TrailingZeros(uint(x[i])))
//line /usr/local/go/src/math/big/nat.go:694
	// _ = "end of CoverTab[5870]"
}

// isPow2 returns i, true when x == 2**i and 0, false otherwise.
func (x nat) isPow2() (uint, bool) {
//line /usr/local/go/src/math/big/nat.go:698
	_go_fuzz_dep_.CoverTab[5874]++
						var i uint
						for x[i] == 0 {
//line /usr/local/go/src/math/big/nat.go:700
		_go_fuzz_dep_.CoverTab[5877]++
							i++
//line /usr/local/go/src/math/big/nat.go:701
		// _ = "end of CoverTab[5877]"
	}
//line /usr/local/go/src/math/big/nat.go:702
	// _ = "end of CoverTab[5874]"
//line /usr/local/go/src/math/big/nat.go:702
	_go_fuzz_dep_.CoverTab[5875]++
						if i == uint(len(x))-1 && func() bool {
//line /usr/local/go/src/math/big/nat.go:703
		_go_fuzz_dep_.CoverTab[5878]++
//line /usr/local/go/src/math/big/nat.go:703
		return x[i]&(x[i]-1) == 0
//line /usr/local/go/src/math/big/nat.go:703
		// _ = "end of CoverTab[5878]"
//line /usr/local/go/src/math/big/nat.go:703
	}() {
//line /usr/local/go/src/math/big/nat.go:703
		_go_fuzz_dep_.CoverTab[5879]++
							return i*_W + uint(bits.TrailingZeros(uint(x[i]))), true
//line /usr/local/go/src/math/big/nat.go:704
		// _ = "end of CoverTab[5879]"
	} else {
//line /usr/local/go/src/math/big/nat.go:705
		_go_fuzz_dep_.CoverTab[5880]++
//line /usr/local/go/src/math/big/nat.go:705
		// _ = "end of CoverTab[5880]"
//line /usr/local/go/src/math/big/nat.go:705
	}
//line /usr/local/go/src/math/big/nat.go:705
	// _ = "end of CoverTab[5875]"
//line /usr/local/go/src/math/big/nat.go:705
	_go_fuzz_dep_.CoverTab[5876]++
						return 0, false
//line /usr/local/go/src/math/big/nat.go:706
	// _ = "end of CoverTab[5876]"
}

func same(x, y nat) bool {
//line /usr/local/go/src/math/big/nat.go:709
	_go_fuzz_dep_.CoverTab[5881]++
						return len(x) == len(y) && func() bool {
//line /usr/local/go/src/math/big/nat.go:710
		_go_fuzz_dep_.CoverTab[5882]++
//line /usr/local/go/src/math/big/nat.go:710
		return len(x) > 0
//line /usr/local/go/src/math/big/nat.go:710
		// _ = "end of CoverTab[5882]"
//line /usr/local/go/src/math/big/nat.go:710
	}() && func() bool {
//line /usr/local/go/src/math/big/nat.go:710
		_go_fuzz_dep_.CoverTab[5883]++
//line /usr/local/go/src/math/big/nat.go:710
		return &x[0] == &y[0]
//line /usr/local/go/src/math/big/nat.go:710
		// _ = "end of CoverTab[5883]"
//line /usr/local/go/src/math/big/nat.go:710
	}()
//line /usr/local/go/src/math/big/nat.go:710
	// _ = "end of CoverTab[5881]"
}

// z = x << s
func (z nat) shl(x nat, s uint) nat {
//line /usr/local/go/src/math/big/nat.go:714
	_go_fuzz_dep_.CoverTab[5884]++
						if s == 0 {
//line /usr/local/go/src/math/big/nat.go:715
		_go_fuzz_dep_.CoverTab[5887]++
							if same(z, x) {
//line /usr/local/go/src/math/big/nat.go:716
			_go_fuzz_dep_.CoverTab[5889]++
								return z
//line /usr/local/go/src/math/big/nat.go:717
			// _ = "end of CoverTab[5889]"
		} else {
//line /usr/local/go/src/math/big/nat.go:718
			_go_fuzz_dep_.CoverTab[5890]++
//line /usr/local/go/src/math/big/nat.go:718
			// _ = "end of CoverTab[5890]"
//line /usr/local/go/src/math/big/nat.go:718
		}
//line /usr/local/go/src/math/big/nat.go:718
		// _ = "end of CoverTab[5887]"
//line /usr/local/go/src/math/big/nat.go:718
		_go_fuzz_dep_.CoverTab[5888]++
							if !alias(z, x) {
//line /usr/local/go/src/math/big/nat.go:719
			_go_fuzz_dep_.CoverTab[5891]++
								return z.set(x)
//line /usr/local/go/src/math/big/nat.go:720
			// _ = "end of CoverTab[5891]"
		} else {
//line /usr/local/go/src/math/big/nat.go:721
			_go_fuzz_dep_.CoverTab[5892]++
//line /usr/local/go/src/math/big/nat.go:721
			// _ = "end of CoverTab[5892]"
//line /usr/local/go/src/math/big/nat.go:721
		}
//line /usr/local/go/src/math/big/nat.go:721
		// _ = "end of CoverTab[5888]"
	} else {
//line /usr/local/go/src/math/big/nat.go:722
		_go_fuzz_dep_.CoverTab[5893]++
//line /usr/local/go/src/math/big/nat.go:722
		// _ = "end of CoverTab[5893]"
//line /usr/local/go/src/math/big/nat.go:722
	}
//line /usr/local/go/src/math/big/nat.go:722
	// _ = "end of CoverTab[5884]"
//line /usr/local/go/src/math/big/nat.go:722
	_go_fuzz_dep_.CoverTab[5885]++

						m := len(x)
						if m == 0 {
//line /usr/local/go/src/math/big/nat.go:725
		_go_fuzz_dep_.CoverTab[5894]++
							return z[:0]
//line /usr/local/go/src/math/big/nat.go:726
		// _ = "end of CoverTab[5894]"
	} else {
//line /usr/local/go/src/math/big/nat.go:727
		_go_fuzz_dep_.CoverTab[5895]++
//line /usr/local/go/src/math/big/nat.go:727
		// _ = "end of CoverTab[5895]"
//line /usr/local/go/src/math/big/nat.go:727
	}
//line /usr/local/go/src/math/big/nat.go:727
	// _ = "end of CoverTab[5885]"
//line /usr/local/go/src/math/big/nat.go:727
	_go_fuzz_dep_.CoverTab[5886]++

//line /usr/local/go/src/math/big/nat.go:730
	n := m + int(s/_W)
						z = z.make(n + 1)
						z[n] = shlVU(z[n-m:n], x, s%_W)
						z[0 : n-m].clear()

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:735
	// _ = "end of CoverTab[5886]"
}

// z = x >> s
func (z nat) shr(x nat, s uint) nat {
//line /usr/local/go/src/math/big/nat.go:739
	_go_fuzz_dep_.CoverTab[5896]++
						if s == 0 {
//line /usr/local/go/src/math/big/nat.go:740
		_go_fuzz_dep_.CoverTab[5899]++
							if same(z, x) {
//line /usr/local/go/src/math/big/nat.go:741
			_go_fuzz_dep_.CoverTab[5901]++
								return z
//line /usr/local/go/src/math/big/nat.go:742
			// _ = "end of CoverTab[5901]"
		} else {
//line /usr/local/go/src/math/big/nat.go:743
			_go_fuzz_dep_.CoverTab[5902]++
//line /usr/local/go/src/math/big/nat.go:743
			// _ = "end of CoverTab[5902]"
//line /usr/local/go/src/math/big/nat.go:743
		}
//line /usr/local/go/src/math/big/nat.go:743
		// _ = "end of CoverTab[5899]"
//line /usr/local/go/src/math/big/nat.go:743
		_go_fuzz_dep_.CoverTab[5900]++
							if !alias(z, x) {
//line /usr/local/go/src/math/big/nat.go:744
			_go_fuzz_dep_.CoverTab[5903]++
								return z.set(x)
//line /usr/local/go/src/math/big/nat.go:745
			// _ = "end of CoverTab[5903]"
		} else {
//line /usr/local/go/src/math/big/nat.go:746
			_go_fuzz_dep_.CoverTab[5904]++
//line /usr/local/go/src/math/big/nat.go:746
			// _ = "end of CoverTab[5904]"
//line /usr/local/go/src/math/big/nat.go:746
		}
//line /usr/local/go/src/math/big/nat.go:746
		// _ = "end of CoverTab[5900]"
	} else {
//line /usr/local/go/src/math/big/nat.go:747
		_go_fuzz_dep_.CoverTab[5905]++
//line /usr/local/go/src/math/big/nat.go:747
		// _ = "end of CoverTab[5905]"
//line /usr/local/go/src/math/big/nat.go:747
	}
//line /usr/local/go/src/math/big/nat.go:747
	// _ = "end of CoverTab[5896]"
//line /usr/local/go/src/math/big/nat.go:747
	_go_fuzz_dep_.CoverTab[5897]++

						m := len(x)
						n := m - int(s/_W)
						if n <= 0 {
//line /usr/local/go/src/math/big/nat.go:751
		_go_fuzz_dep_.CoverTab[5906]++
							return z[:0]
//line /usr/local/go/src/math/big/nat.go:752
		// _ = "end of CoverTab[5906]"
	} else {
//line /usr/local/go/src/math/big/nat.go:753
		_go_fuzz_dep_.CoverTab[5907]++
//line /usr/local/go/src/math/big/nat.go:753
		// _ = "end of CoverTab[5907]"
//line /usr/local/go/src/math/big/nat.go:753
	}
//line /usr/local/go/src/math/big/nat.go:753
	// _ = "end of CoverTab[5897]"
//line /usr/local/go/src/math/big/nat.go:753
	_go_fuzz_dep_.CoverTab[5898]++

//line /usr/local/go/src/math/big/nat.go:756
	z = z.make(n)
						shrVU(z, x[m-n:], s%_W)

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:759
	// _ = "end of CoverTab[5898]"
}

func (z nat) setBit(x nat, i uint, b uint) nat {
//line /usr/local/go/src/math/big/nat.go:762
	_go_fuzz_dep_.CoverTab[5908]++
						j := int(i / _W)
						m := Word(1) << (i % _W)
						n := len(x)
						switch b {
	case 0:
//line /usr/local/go/src/math/big/nat.go:767
		_go_fuzz_dep_.CoverTab[5910]++
							z = z.make(n)
							copy(z, x)
							if j >= n {
//line /usr/local/go/src/math/big/nat.go:770
			_go_fuzz_dep_.CoverTab[5915]++

								return z
//line /usr/local/go/src/math/big/nat.go:772
			// _ = "end of CoverTab[5915]"
		} else {
//line /usr/local/go/src/math/big/nat.go:773
			_go_fuzz_dep_.CoverTab[5916]++
//line /usr/local/go/src/math/big/nat.go:773
			// _ = "end of CoverTab[5916]"
//line /usr/local/go/src/math/big/nat.go:773
		}
//line /usr/local/go/src/math/big/nat.go:773
		// _ = "end of CoverTab[5910]"
//line /usr/local/go/src/math/big/nat.go:773
		_go_fuzz_dep_.CoverTab[5911]++
							z[j] &^= m
							return z.norm()
//line /usr/local/go/src/math/big/nat.go:775
		// _ = "end of CoverTab[5911]"
	case 1:
//line /usr/local/go/src/math/big/nat.go:776
		_go_fuzz_dep_.CoverTab[5912]++
							if j >= n {
//line /usr/local/go/src/math/big/nat.go:777
			_go_fuzz_dep_.CoverTab[5917]++
								z = z.make(j + 1)
								z[n:].clear()
//line /usr/local/go/src/math/big/nat.go:779
			// _ = "end of CoverTab[5917]"
		} else {
//line /usr/local/go/src/math/big/nat.go:780
			_go_fuzz_dep_.CoverTab[5918]++
								z = z.make(n)
//line /usr/local/go/src/math/big/nat.go:781
			// _ = "end of CoverTab[5918]"
		}
//line /usr/local/go/src/math/big/nat.go:782
		// _ = "end of CoverTab[5912]"
//line /usr/local/go/src/math/big/nat.go:782
		_go_fuzz_dep_.CoverTab[5913]++
							copy(z, x)
							z[j] |= m

							return z
//line /usr/local/go/src/math/big/nat.go:786
		// _ = "end of CoverTab[5913]"
//line /usr/local/go/src/math/big/nat.go:786
	default:
//line /usr/local/go/src/math/big/nat.go:786
		_go_fuzz_dep_.CoverTab[5914]++
//line /usr/local/go/src/math/big/nat.go:786
		// _ = "end of CoverTab[5914]"
	}
//line /usr/local/go/src/math/big/nat.go:787
	// _ = "end of CoverTab[5908]"
//line /usr/local/go/src/math/big/nat.go:787
	_go_fuzz_dep_.CoverTab[5909]++
						panic("set bit is not 0 or 1")
//line /usr/local/go/src/math/big/nat.go:788
	// _ = "end of CoverTab[5909]"
}

// bit returns the value of the i'th bit, with lsb == bit 0.
func (x nat) bit(i uint) uint {
//line /usr/local/go/src/math/big/nat.go:792
	_go_fuzz_dep_.CoverTab[5919]++
						j := i / _W
						if j >= uint(len(x)) {
//line /usr/local/go/src/math/big/nat.go:794
		_go_fuzz_dep_.CoverTab[5921]++
							return 0
//line /usr/local/go/src/math/big/nat.go:795
		// _ = "end of CoverTab[5921]"
	} else {
//line /usr/local/go/src/math/big/nat.go:796
		_go_fuzz_dep_.CoverTab[5922]++
//line /usr/local/go/src/math/big/nat.go:796
		// _ = "end of CoverTab[5922]"
//line /usr/local/go/src/math/big/nat.go:796
	}
//line /usr/local/go/src/math/big/nat.go:796
	// _ = "end of CoverTab[5919]"
//line /usr/local/go/src/math/big/nat.go:796
	_go_fuzz_dep_.CoverTab[5920]++

						return uint(x[j] >> (i % _W) & 1)
//line /usr/local/go/src/math/big/nat.go:798
	// _ = "end of CoverTab[5920]"
}

// sticky returns 1 if there's a 1 bit within the
//line /usr/local/go/src/math/big/nat.go:801
// i least significant bits, otherwise it returns 0.
//line /usr/local/go/src/math/big/nat.go:803
func (x nat) sticky(i uint) uint {
//line /usr/local/go/src/math/big/nat.go:803
	_go_fuzz_dep_.CoverTab[5923]++
						j := i / _W
						if j >= uint(len(x)) {
//line /usr/local/go/src/math/big/nat.go:805
		_go_fuzz_dep_.CoverTab[5927]++
							if len(x) == 0 {
//line /usr/local/go/src/math/big/nat.go:806
			_go_fuzz_dep_.CoverTab[5929]++
								return 0
//line /usr/local/go/src/math/big/nat.go:807
			// _ = "end of CoverTab[5929]"
		} else {
//line /usr/local/go/src/math/big/nat.go:808
			_go_fuzz_dep_.CoverTab[5930]++
//line /usr/local/go/src/math/big/nat.go:808
			// _ = "end of CoverTab[5930]"
//line /usr/local/go/src/math/big/nat.go:808
		}
//line /usr/local/go/src/math/big/nat.go:808
		// _ = "end of CoverTab[5927]"
//line /usr/local/go/src/math/big/nat.go:808
		_go_fuzz_dep_.CoverTab[5928]++
							return 1
//line /usr/local/go/src/math/big/nat.go:809
		// _ = "end of CoverTab[5928]"
	} else {
//line /usr/local/go/src/math/big/nat.go:810
		_go_fuzz_dep_.CoverTab[5931]++
//line /usr/local/go/src/math/big/nat.go:810
		// _ = "end of CoverTab[5931]"
//line /usr/local/go/src/math/big/nat.go:810
	}
//line /usr/local/go/src/math/big/nat.go:810
	// _ = "end of CoverTab[5923]"
//line /usr/local/go/src/math/big/nat.go:810
	_go_fuzz_dep_.CoverTab[5924]++

						for _, x := range x[:j] {
//line /usr/local/go/src/math/big/nat.go:812
		_go_fuzz_dep_.CoverTab[5932]++
							if x != 0 {
//line /usr/local/go/src/math/big/nat.go:813
			_go_fuzz_dep_.CoverTab[5933]++
								return 1
//line /usr/local/go/src/math/big/nat.go:814
			// _ = "end of CoverTab[5933]"
		} else {
//line /usr/local/go/src/math/big/nat.go:815
			_go_fuzz_dep_.CoverTab[5934]++
//line /usr/local/go/src/math/big/nat.go:815
			// _ = "end of CoverTab[5934]"
//line /usr/local/go/src/math/big/nat.go:815
		}
//line /usr/local/go/src/math/big/nat.go:815
		// _ = "end of CoverTab[5932]"
	}
//line /usr/local/go/src/math/big/nat.go:816
	// _ = "end of CoverTab[5924]"
//line /usr/local/go/src/math/big/nat.go:816
	_go_fuzz_dep_.CoverTab[5925]++
						if x[j]<<(_W-i%_W) != 0 {
//line /usr/local/go/src/math/big/nat.go:817
		_go_fuzz_dep_.CoverTab[5935]++
							return 1
//line /usr/local/go/src/math/big/nat.go:818
		// _ = "end of CoverTab[5935]"
	} else {
//line /usr/local/go/src/math/big/nat.go:819
		_go_fuzz_dep_.CoverTab[5936]++
//line /usr/local/go/src/math/big/nat.go:819
		// _ = "end of CoverTab[5936]"
//line /usr/local/go/src/math/big/nat.go:819
	}
//line /usr/local/go/src/math/big/nat.go:819
	// _ = "end of CoverTab[5925]"
//line /usr/local/go/src/math/big/nat.go:819
	_go_fuzz_dep_.CoverTab[5926]++
						return 0
//line /usr/local/go/src/math/big/nat.go:820
	// _ = "end of CoverTab[5926]"
}

func (z nat) and(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:823
	_go_fuzz_dep_.CoverTab[5937]++
						m := len(x)
						n := len(y)
						if m > n {
//line /usr/local/go/src/math/big/nat.go:826
		_go_fuzz_dep_.CoverTab[5940]++
							m = n
//line /usr/local/go/src/math/big/nat.go:827
		// _ = "end of CoverTab[5940]"
	} else {
//line /usr/local/go/src/math/big/nat.go:828
		_go_fuzz_dep_.CoverTab[5941]++
//line /usr/local/go/src/math/big/nat.go:828
		// _ = "end of CoverTab[5941]"
//line /usr/local/go/src/math/big/nat.go:828
	}
//line /usr/local/go/src/math/big/nat.go:828
	// _ = "end of CoverTab[5937]"
//line /usr/local/go/src/math/big/nat.go:828
	_go_fuzz_dep_.CoverTab[5938]++

//line /usr/local/go/src/math/big/nat.go:831
	z = z.make(m)
	for i := 0; i < m; i++ {
//line /usr/local/go/src/math/big/nat.go:832
		_go_fuzz_dep_.CoverTab[5942]++
							z[i] = x[i] & y[i]
//line /usr/local/go/src/math/big/nat.go:833
		// _ = "end of CoverTab[5942]"
	}
//line /usr/local/go/src/math/big/nat.go:834
	// _ = "end of CoverTab[5938]"
//line /usr/local/go/src/math/big/nat.go:834
	_go_fuzz_dep_.CoverTab[5939]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:836
	// _ = "end of CoverTab[5939]"
}

// trunc returns z = x mod 2ⁿ.
func (z nat) trunc(x nat, n uint) nat {
//line /usr/local/go/src/math/big/nat.go:840
	_go_fuzz_dep_.CoverTab[5943]++
						w := (n + _W - 1) / _W
						if uint(len(x)) < w {
//line /usr/local/go/src/math/big/nat.go:842
		_go_fuzz_dep_.CoverTab[5946]++
							return z.set(x)
//line /usr/local/go/src/math/big/nat.go:843
		// _ = "end of CoverTab[5946]"
	} else {
//line /usr/local/go/src/math/big/nat.go:844
		_go_fuzz_dep_.CoverTab[5947]++
//line /usr/local/go/src/math/big/nat.go:844
		// _ = "end of CoverTab[5947]"
//line /usr/local/go/src/math/big/nat.go:844
	}
//line /usr/local/go/src/math/big/nat.go:844
	// _ = "end of CoverTab[5943]"
//line /usr/local/go/src/math/big/nat.go:844
	_go_fuzz_dep_.CoverTab[5944]++
						z = z.make(int(w))
						copy(z, x)
						if n%_W != 0 {
//line /usr/local/go/src/math/big/nat.go:847
		_go_fuzz_dep_.CoverTab[5948]++
							z[len(z)-1] &= 1<<(n%_W) - 1
//line /usr/local/go/src/math/big/nat.go:848
		// _ = "end of CoverTab[5948]"
	} else {
//line /usr/local/go/src/math/big/nat.go:849
		_go_fuzz_dep_.CoverTab[5949]++
//line /usr/local/go/src/math/big/nat.go:849
		// _ = "end of CoverTab[5949]"
//line /usr/local/go/src/math/big/nat.go:849
	}
//line /usr/local/go/src/math/big/nat.go:849
	// _ = "end of CoverTab[5944]"
//line /usr/local/go/src/math/big/nat.go:849
	_go_fuzz_dep_.CoverTab[5945]++
						return z.norm()
//line /usr/local/go/src/math/big/nat.go:850
	// _ = "end of CoverTab[5945]"
}

func (z nat) andNot(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:853
	_go_fuzz_dep_.CoverTab[5950]++
						m := len(x)
						n := len(y)
						if n > m {
//line /usr/local/go/src/math/big/nat.go:856
		_go_fuzz_dep_.CoverTab[5953]++
							n = m
//line /usr/local/go/src/math/big/nat.go:857
		// _ = "end of CoverTab[5953]"
	} else {
//line /usr/local/go/src/math/big/nat.go:858
		_go_fuzz_dep_.CoverTab[5954]++
//line /usr/local/go/src/math/big/nat.go:858
		// _ = "end of CoverTab[5954]"
//line /usr/local/go/src/math/big/nat.go:858
	}
//line /usr/local/go/src/math/big/nat.go:858
	// _ = "end of CoverTab[5950]"
//line /usr/local/go/src/math/big/nat.go:858
	_go_fuzz_dep_.CoverTab[5951]++

//line /usr/local/go/src/math/big/nat.go:861
	z = z.make(m)
	for i := 0; i < n; i++ {
//line /usr/local/go/src/math/big/nat.go:862
		_go_fuzz_dep_.CoverTab[5955]++
							z[i] = x[i] &^ y[i]
//line /usr/local/go/src/math/big/nat.go:863
		// _ = "end of CoverTab[5955]"
	}
//line /usr/local/go/src/math/big/nat.go:864
	// _ = "end of CoverTab[5951]"
//line /usr/local/go/src/math/big/nat.go:864
	_go_fuzz_dep_.CoverTab[5952]++
						copy(z[n:m], x[n:m])

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:867
	// _ = "end of CoverTab[5952]"
}

func (z nat) or(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:870
	_go_fuzz_dep_.CoverTab[5956]++
						m := len(x)
						n := len(y)
						s := x
						if m < n {
//line /usr/local/go/src/math/big/nat.go:874
		_go_fuzz_dep_.CoverTab[5959]++
							n, m = m, n
							s = y
//line /usr/local/go/src/math/big/nat.go:876
		// _ = "end of CoverTab[5959]"
	} else {
//line /usr/local/go/src/math/big/nat.go:877
		_go_fuzz_dep_.CoverTab[5960]++
//line /usr/local/go/src/math/big/nat.go:877
		// _ = "end of CoverTab[5960]"
//line /usr/local/go/src/math/big/nat.go:877
	}
//line /usr/local/go/src/math/big/nat.go:877
	// _ = "end of CoverTab[5956]"
//line /usr/local/go/src/math/big/nat.go:877
	_go_fuzz_dep_.CoverTab[5957]++

//line /usr/local/go/src/math/big/nat.go:880
	z = z.make(m)
	for i := 0; i < n; i++ {
//line /usr/local/go/src/math/big/nat.go:881
		_go_fuzz_dep_.CoverTab[5961]++
							z[i] = x[i] | y[i]
//line /usr/local/go/src/math/big/nat.go:882
		// _ = "end of CoverTab[5961]"
	}
//line /usr/local/go/src/math/big/nat.go:883
	// _ = "end of CoverTab[5957]"
//line /usr/local/go/src/math/big/nat.go:883
	_go_fuzz_dep_.CoverTab[5958]++
						copy(z[n:m], s[n:m])

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:886
	// _ = "end of CoverTab[5958]"
}

func (z nat) xor(x, y nat) nat {
//line /usr/local/go/src/math/big/nat.go:889
	_go_fuzz_dep_.CoverTab[5962]++
						m := len(x)
						n := len(y)
						s := x
						if m < n {
//line /usr/local/go/src/math/big/nat.go:893
		_go_fuzz_dep_.CoverTab[5965]++
							n, m = m, n
							s = y
//line /usr/local/go/src/math/big/nat.go:895
		// _ = "end of CoverTab[5965]"
	} else {
//line /usr/local/go/src/math/big/nat.go:896
		_go_fuzz_dep_.CoverTab[5966]++
//line /usr/local/go/src/math/big/nat.go:896
		// _ = "end of CoverTab[5966]"
//line /usr/local/go/src/math/big/nat.go:896
	}
//line /usr/local/go/src/math/big/nat.go:896
	// _ = "end of CoverTab[5962]"
//line /usr/local/go/src/math/big/nat.go:896
	_go_fuzz_dep_.CoverTab[5963]++

//line /usr/local/go/src/math/big/nat.go:899
	z = z.make(m)
	for i := 0; i < n; i++ {
//line /usr/local/go/src/math/big/nat.go:900
		_go_fuzz_dep_.CoverTab[5967]++
							z[i] = x[i] ^ y[i]
//line /usr/local/go/src/math/big/nat.go:901
		// _ = "end of CoverTab[5967]"
	}
//line /usr/local/go/src/math/big/nat.go:902
	// _ = "end of CoverTab[5963]"
//line /usr/local/go/src/math/big/nat.go:902
	_go_fuzz_dep_.CoverTab[5964]++
						copy(z[n:m], s[n:m])

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:905
	// _ = "end of CoverTab[5964]"
}

// random creates a random integer in [0..limit), using the space in z if
//line /usr/local/go/src/math/big/nat.go:908
// possible. n is the bit length of limit.
//line /usr/local/go/src/math/big/nat.go:910
func (z nat) random(rand *rand.Rand, limit nat, n int) nat {
//line /usr/local/go/src/math/big/nat.go:910
	_go_fuzz_dep_.CoverTab[5968]++
						if alias(z, limit) {
//line /usr/local/go/src/math/big/nat.go:911
		_go_fuzz_dep_.CoverTab[5972]++
							z = nil
//line /usr/local/go/src/math/big/nat.go:912
		// _ = "end of CoverTab[5972]"
	} else {
//line /usr/local/go/src/math/big/nat.go:913
		_go_fuzz_dep_.CoverTab[5973]++
//line /usr/local/go/src/math/big/nat.go:913
		// _ = "end of CoverTab[5973]"
//line /usr/local/go/src/math/big/nat.go:913
	}
//line /usr/local/go/src/math/big/nat.go:913
	// _ = "end of CoverTab[5968]"
//line /usr/local/go/src/math/big/nat.go:913
	_go_fuzz_dep_.CoverTab[5969]++
						z = z.make(len(limit))

						bitLengthOfMSW := uint(n % _W)
						if bitLengthOfMSW == 0 {
//line /usr/local/go/src/math/big/nat.go:917
		_go_fuzz_dep_.CoverTab[5974]++
							bitLengthOfMSW = _W
//line /usr/local/go/src/math/big/nat.go:918
		// _ = "end of CoverTab[5974]"
	} else {
//line /usr/local/go/src/math/big/nat.go:919
		_go_fuzz_dep_.CoverTab[5975]++
//line /usr/local/go/src/math/big/nat.go:919
		// _ = "end of CoverTab[5975]"
//line /usr/local/go/src/math/big/nat.go:919
	}
//line /usr/local/go/src/math/big/nat.go:919
	// _ = "end of CoverTab[5969]"
//line /usr/local/go/src/math/big/nat.go:919
	_go_fuzz_dep_.CoverTab[5970]++
						mask := Word((1 << bitLengthOfMSW) - 1)

						for {
//line /usr/local/go/src/math/big/nat.go:922
		_go_fuzz_dep_.CoverTab[5976]++
							switch _W {
		case 32:
//line /usr/local/go/src/math/big/nat.go:924
			_go_fuzz_dep_.CoverTab[5978]++
								for i := range z {
//line /usr/local/go/src/math/big/nat.go:925
				_go_fuzz_dep_.CoverTab[5981]++
									z[i] = Word(rand.Uint32())
//line /usr/local/go/src/math/big/nat.go:926
				// _ = "end of CoverTab[5981]"
			}
//line /usr/local/go/src/math/big/nat.go:927
			// _ = "end of CoverTab[5978]"
		case 64:
//line /usr/local/go/src/math/big/nat.go:928
			_go_fuzz_dep_.CoverTab[5979]++
								for i := range z {
//line /usr/local/go/src/math/big/nat.go:929
				_go_fuzz_dep_.CoverTab[5982]++
									z[i] = Word(rand.Uint32()) | Word(rand.Uint32())<<32
//line /usr/local/go/src/math/big/nat.go:930
				// _ = "end of CoverTab[5982]"
			}
//line /usr/local/go/src/math/big/nat.go:931
			// _ = "end of CoverTab[5979]"
		default:
//line /usr/local/go/src/math/big/nat.go:932
			_go_fuzz_dep_.CoverTab[5980]++
								panic("unknown word size")
//line /usr/local/go/src/math/big/nat.go:933
			// _ = "end of CoverTab[5980]"
		}
//line /usr/local/go/src/math/big/nat.go:934
		// _ = "end of CoverTab[5976]"
//line /usr/local/go/src/math/big/nat.go:934
		_go_fuzz_dep_.CoverTab[5977]++
							z[len(limit)-1] &= mask
							if z.cmp(limit) < 0 {
//line /usr/local/go/src/math/big/nat.go:936
			_go_fuzz_dep_.CoverTab[5983]++
								break
//line /usr/local/go/src/math/big/nat.go:937
			// _ = "end of CoverTab[5983]"
		} else {
//line /usr/local/go/src/math/big/nat.go:938
			_go_fuzz_dep_.CoverTab[5984]++
//line /usr/local/go/src/math/big/nat.go:938
			// _ = "end of CoverTab[5984]"
//line /usr/local/go/src/math/big/nat.go:938
		}
//line /usr/local/go/src/math/big/nat.go:938
		// _ = "end of CoverTab[5977]"
	}
//line /usr/local/go/src/math/big/nat.go:939
	// _ = "end of CoverTab[5970]"
//line /usr/local/go/src/math/big/nat.go:939
	_go_fuzz_dep_.CoverTab[5971]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:941
	// _ = "end of CoverTab[5971]"
}

// If m != 0 (i.e., len(m) != 0), expNN sets z to x**y mod m;
//line /usr/local/go/src/math/big/nat.go:944
// otherwise it sets z to x**y. The result is the value of z.
//line /usr/local/go/src/math/big/nat.go:946
func (z nat) expNN(x, y, m nat, slow bool) nat {
//line /usr/local/go/src/math/big/nat.go:946
	_go_fuzz_dep_.CoverTab[5985]++
						if alias(z, x) || func() bool {
//line /usr/local/go/src/math/big/nat.go:947
		_go_fuzz_dep_.CoverTab[5995]++
//line /usr/local/go/src/math/big/nat.go:947
		return alias(z, y)
//line /usr/local/go/src/math/big/nat.go:947
		// _ = "end of CoverTab[5995]"
//line /usr/local/go/src/math/big/nat.go:947
	}() {
//line /usr/local/go/src/math/big/nat.go:947
		_go_fuzz_dep_.CoverTab[5996]++

							z = nil
//line /usr/local/go/src/math/big/nat.go:949
		// _ = "end of CoverTab[5996]"
	} else {
//line /usr/local/go/src/math/big/nat.go:950
		_go_fuzz_dep_.CoverTab[5997]++
//line /usr/local/go/src/math/big/nat.go:950
		// _ = "end of CoverTab[5997]"
//line /usr/local/go/src/math/big/nat.go:950
	}
//line /usr/local/go/src/math/big/nat.go:950
	// _ = "end of CoverTab[5985]"
//line /usr/local/go/src/math/big/nat.go:950
	_go_fuzz_dep_.CoverTab[5986]++

//line /usr/local/go/src/math/big/nat.go:953
	if len(m) == 1 && func() bool {
//line /usr/local/go/src/math/big/nat.go:953
		_go_fuzz_dep_.CoverTab[5998]++
//line /usr/local/go/src/math/big/nat.go:953
		return m[0] == 1
//line /usr/local/go/src/math/big/nat.go:953
		// _ = "end of CoverTab[5998]"
//line /usr/local/go/src/math/big/nat.go:953
	}() {
//line /usr/local/go/src/math/big/nat.go:953
		_go_fuzz_dep_.CoverTab[5999]++
							return z.setWord(0)
//line /usr/local/go/src/math/big/nat.go:954
		// _ = "end of CoverTab[5999]"
	} else {
//line /usr/local/go/src/math/big/nat.go:955
		_go_fuzz_dep_.CoverTab[6000]++
//line /usr/local/go/src/math/big/nat.go:955
		// _ = "end of CoverTab[6000]"
//line /usr/local/go/src/math/big/nat.go:955
	}
//line /usr/local/go/src/math/big/nat.go:955
	// _ = "end of CoverTab[5986]"
//line /usr/local/go/src/math/big/nat.go:955
	_go_fuzz_dep_.CoverTab[5987]++

//line /usr/local/go/src/math/big/nat.go:959
	if len(y) == 0 {
//line /usr/local/go/src/math/big/nat.go:959
		_go_fuzz_dep_.CoverTab[6001]++
							return z.setWord(1)
//line /usr/local/go/src/math/big/nat.go:960
		// _ = "end of CoverTab[6001]"
	} else {
//line /usr/local/go/src/math/big/nat.go:961
		_go_fuzz_dep_.CoverTab[6002]++
//line /usr/local/go/src/math/big/nat.go:961
		// _ = "end of CoverTab[6002]"
//line /usr/local/go/src/math/big/nat.go:961
	}
//line /usr/local/go/src/math/big/nat.go:961
	// _ = "end of CoverTab[5987]"
//line /usr/local/go/src/math/big/nat.go:961
	_go_fuzz_dep_.CoverTab[5988]++

//line /usr/local/go/src/math/big/nat.go:965
	if len(x) == 0 {
//line /usr/local/go/src/math/big/nat.go:965
		_go_fuzz_dep_.CoverTab[6003]++
							return z.setWord(0)
//line /usr/local/go/src/math/big/nat.go:966
		// _ = "end of CoverTab[6003]"
	} else {
//line /usr/local/go/src/math/big/nat.go:967
		_go_fuzz_dep_.CoverTab[6004]++
//line /usr/local/go/src/math/big/nat.go:967
		// _ = "end of CoverTab[6004]"
//line /usr/local/go/src/math/big/nat.go:967
	}
//line /usr/local/go/src/math/big/nat.go:967
	// _ = "end of CoverTab[5988]"
//line /usr/local/go/src/math/big/nat.go:967
	_go_fuzz_dep_.CoverTab[5989]++

//line /usr/local/go/src/math/big/nat.go:971
	if len(x) == 1 && func() bool {
//line /usr/local/go/src/math/big/nat.go:971
		_go_fuzz_dep_.CoverTab[6005]++
//line /usr/local/go/src/math/big/nat.go:971
		return x[0] == 1
//line /usr/local/go/src/math/big/nat.go:971
		// _ = "end of CoverTab[6005]"
//line /usr/local/go/src/math/big/nat.go:971
	}() {
//line /usr/local/go/src/math/big/nat.go:971
		_go_fuzz_dep_.CoverTab[6006]++
							return z.setWord(1)
//line /usr/local/go/src/math/big/nat.go:972
		// _ = "end of CoverTab[6006]"
	} else {
//line /usr/local/go/src/math/big/nat.go:973
		_go_fuzz_dep_.CoverTab[6007]++
//line /usr/local/go/src/math/big/nat.go:973
		// _ = "end of CoverTab[6007]"
//line /usr/local/go/src/math/big/nat.go:973
	}
//line /usr/local/go/src/math/big/nat.go:973
	// _ = "end of CoverTab[5989]"
//line /usr/local/go/src/math/big/nat.go:973
	_go_fuzz_dep_.CoverTab[5990]++

//line /usr/local/go/src/math/big/nat.go:977
	if len(y) == 1 && func() bool {
//line /usr/local/go/src/math/big/nat.go:977
		_go_fuzz_dep_.CoverTab[6008]++
//line /usr/local/go/src/math/big/nat.go:977
		return y[0] == 1
//line /usr/local/go/src/math/big/nat.go:977
		// _ = "end of CoverTab[6008]"
//line /usr/local/go/src/math/big/nat.go:977
	}() {
//line /usr/local/go/src/math/big/nat.go:977
		_go_fuzz_dep_.CoverTab[6009]++
							if len(m) != 0 {
//line /usr/local/go/src/math/big/nat.go:978
			_go_fuzz_dep_.CoverTab[6011]++
								return z.rem(x, m)
//line /usr/local/go/src/math/big/nat.go:979
			// _ = "end of CoverTab[6011]"
		} else {
//line /usr/local/go/src/math/big/nat.go:980
			_go_fuzz_dep_.CoverTab[6012]++
//line /usr/local/go/src/math/big/nat.go:980
			// _ = "end of CoverTab[6012]"
//line /usr/local/go/src/math/big/nat.go:980
		}
//line /usr/local/go/src/math/big/nat.go:980
		// _ = "end of CoverTab[6009]"
//line /usr/local/go/src/math/big/nat.go:980
		_go_fuzz_dep_.CoverTab[6010]++
							return z.set(x)
//line /usr/local/go/src/math/big/nat.go:981
		// _ = "end of CoverTab[6010]"
	} else {
//line /usr/local/go/src/math/big/nat.go:982
		_go_fuzz_dep_.CoverTab[6013]++
//line /usr/local/go/src/math/big/nat.go:982
		// _ = "end of CoverTab[6013]"
//line /usr/local/go/src/math/big/nat.go:982
	}
//line /usr/local/go/src/math/big/nat.go:982
	// _ = "end of CoverTab[5990]"
//line /usr/local/go/src/math/big/nat.go:982
	_go_fuzz_dep_.CoverTab[5991]++

//line /usr/local/go/src/math/big/nat.go:985
	if len(m) != 0 {
//line /usr/local/go/src/math/big/nat.go:985
		_go_fuzz_dep_.CoverTab[6014]++

							z = z.make(len(m))

//line /usr/local/go/src/math/big/nat.go:994
		if len(y) > 1 && func() bool {
//line /usr/local/go/src/math/big/nat.go:994
			_go_fuzz_dep_.CoverTab[6015]++
//line /usr/local/go/src/math/big/nat.go:994
			return !slow
//line /usr/local/go/src/math/big/nat.go:994
			// _ = "end of CoverTab[6015]"
//line /usr/local/go/src/math/big/nat.go:994
		}() {
//line /usr/local/go/src/math/big/nat.go:994
			_go_fuzz_dep_.CoverTab[6016]++
								if m[0]&1 == 1 {
//line /usr/local/go/src/math/big/nat.go:995
				_go_fuzz_dep_.CoverTab[6019]++
									return z.expNNMontgomery(x, y, m)
//line /usr/local/go/src/math/big/nat.go:996
				// _ = "end of CoverTab[6019]"
			} else {
//line /usr/local/go/src/math/big/nat.go:997
				_go_fuzz_dep_.CoverTab[6020]++
//line /usr/local/go/src/math/big/nat.go:997
				// _ = "end of CoverTab[6020]"
//line /usr/local/go/src/math/big/nat.go:997
			}
//line /usr/local/go/src/math/big/nat.go:997
			// _ = "end of CoverTab[6016]"
//line /usr/local/go/src/math/big/nat.go:997
			_go_fuzz_dep_.CoverTab[6017]++
								if logM, ok := m.isPow2(); ok {
//line /usr/local/go/src/math/big/nat.go:998
				_go_fuzz_dep_.CoverTab[6021]++
									return z.expNNWindowed(x, y, logM)
//line /usr/local/go/src/math/big/nat.go:999
				// _ = "end of CoverTab[6021]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1000
				_go_fuzz_dep_.CoverTab[6022]++
//line /usr/local/go/src/math/big/nat.go:1000
				// _ = "end of CoverTab[6022]"
//line /usr/local/go/src/math/big/nat.go:1000
			}
//line /usr/local/go/src/math/big/nat.go:1000
			// _ = "end of CoverTab[6017]"
//line /usr/local/go/src/math/big/nat.go:1000
			_go_fuzz_dep_.CoverTab[6018]++
								return z.expNNMontgomeryEven(x, y, m)
//line /usr/local/go/src/math/big/nat.go:1001
			// _ = "end of CoverTab[6018]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1002
			_go_fuzz_dep_.CoverTab[6023]++
//line /usr/local/go/src/math/big/nat.go:1002
			// _ = "end of CoverTab[6023]"
//line /usr/local/go/src/math/big/nat.go:1002
		}
//line /usr/local/go/src/math/big/nat.go:1002
		// _ = "end of CoverTab[6014]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1003
		_go_fuzz_dep_.CoverTab[6024]++
//line /usr/local/go/src/math/big/nat.go:1003
		// _ = "end of CoverTab[6024]"
//line /usr/local/go/src/math/big/nat.go:1003
	}
//line /usr/local/go/src/math/big/nat.go:1003
	// _ = "end of CoverTab[5991]"
//line /usr/local/go/src/math/big/nat.go:1003
	_go_fuzz_dep_.CoverTab[5992]++

						z = z.set(x)
						v := y[len(y)-1]
						shift := nlz(v) + 1
						v <<= shift
						var q nat

						const mask = 1 << (_W - 1)

//line /usr/local/go/src/math/big/nat.go:1017
	w := _W - int(shift)
	// zz and r are used to avoid allocating in mul and div as
	// otherwise the arguments would alias.
	var zz, r nat
	for j := 0; j < w; j++ {
//line /usr/local/go/src/math/big/nat.go:1021
		_go_fuzz_dep_.CoverTab[6025]++
							zz = zz.sqr(z)
							zz, z = z, zz

							if v&mask != 0 {
//line /usr/local/go/src/math/big/nat.go:1025
			_go_fuzz_dep_.CoverTab[6028]++
								zz = zz.mul(z, x)
								zz, z = z, zz
//line /usr/local/go/src/math/big/nat.go:1027
			// _ = "end of CoverTab[6028]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1028
			_go_fuzz_dep_.CoverTab[6029]++
//line /usr/local/go/src/math/big/nat.go:1028
			// _ = "end of CoverTab[6029]"
//line /usr/local/go/src/math/big/nat.go:1028
		}
//line /usr/local/go/src/math/big/nat.go:1028
		// _ = "end of CoverTab[6025]"
//line /usr/local/go/src/math/big/nat.go:1028
		_go_fuzz_dep_.CoverTab[6026]++

							if len(m) != 0 {
//line /usr/local/go/src/math/big/nat.go:1030
			_go_fuzz_dep_.CoverTab[6030]++
								zz, r = zz.div(r, z, m)
								zz, r, q, z = q, z, zz, r
//line /usr/local/go/src/math/big/nat.go:1032
			// _ = "end of CoverTab[6030]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1033
			_go_fuzz_dep_.CoverTab[6031]++
//line /usr/local/go/src/math/big/nat.go:1033
			// _ = "end of CoverTab[6031]"
//line /usr/local/go/src/math/big/nat.go:1033
		}
//line /usr/local/go/src/math/big/nat.go:1033
		// _ = "end of CoverTab[6026]"
//line /usr/local/go/src/math/big/nat.go:1033
		_go_fuzz_dep_.CoverTab[6027]++

							v <<= 1
//line /usr/local/go/src/math/big/nat.go:1035
		// _ = "end of CoverTab[6027]"
	}
//line /usr/local/go/src/math/big/nat.go:1036
	// _ = "end of CoverTab[5992]"
//line /usr/local/go/src/math/big/nat.go:1036
	_go_fuzz_dep_.CoverTab[5993]++

						for i := len(y) - 2; i >= 0; i-- {
//line /usr/local/go/src/math/big/nat.go:1038
		_go_fuzz_dep_.CoverTab[6032]++
							v = y[i]

							for j := 0; j < _W; j++ {
//line /usr/local/go/src/math/big/nat.go:1041
			_go_fuzz_dep_.CoverTab[6033]++
								zz = zz.sqr(z)
								zz, z = z, zz

								if v&mask != 0 {
//line /usr/local/go/src/math/big/nat.go:1045
				_go_fuzz_dep_.CoverTab[6036]++
									zz = zz.mul(z, x)
									zz, z = z, zz
//line /usr/local/go/src/math/big/nat.go:1047
				// _ = "end of CoverTab[6036]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1048
				_go_fuzz_dep_.CoverTab[6037]++
//line /usr/local/go/src/math/big/nat.go:1048
				// _ = "end of CoverTab[6037]"
//line /usr/local/go/src/math/big/nat.go:1048
			}
//line /usr/local/go/src/math/big/nat.go:1048
			// _ = "end of CoverTab[6033]"
//line /usr/local/go/src/math/big/nat.go:1048
			_go_fuzz_dep_.CoverTab[6034]++

								if len(m) != 0 {
//line /usr/local/go/src/math/big/nat.go:1050
				_go_fuzz_dep_.CoverTab[6038]++
									zz, r = zz.div(r, z, m)
									zz, r, q, z = q, z, zz, r
//line /usr/local/go/src/math/big/nat.go:1052
				// _ = "end of CoverTab[6038]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1053
				_go_fuzz_dep_.CoverTab[6039]++
//line /usr/local/go/src/math/big/nat.go:1053
				// _ = "end of CoverTab[6039]"
//line /usr/local/go/src/math/big/nat.go:1053
			}
//line /usr/local/go/src/math/big/nat.go:1053
			// _ = "end of CoverTab[6034]"
//line /usr/local/go/src/math/big/nat.go:1053
			_go_fuzz_dep_.CoverTab[6035]++

								v <<= 1
//line /usr/local/go/src/math/big/nat.go:1055
			// _ = "end of CoverTab[6035]"
		}
//line /usr/local/go/src/math/big/nat.go:1056
		// _ = "end of CoverTab[6032]"
	}
//line /usr/local/go/src/math/big/nat.go:1057
	// _ = "end of CoverTab[5993]"
//line /usr/local/go/src/math/big/nat.go:1057
	_go_fuzz_dep_.CoverTab[5994]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:1059
	// _ = "end of CoverTab[5994]"
}

// expNNMontgomeryEven calculates x**y mod m where m = m1 × m2 for m1 = 2ⁿ and m2 odd.
//line /usr/local/go/src/math/big/nat.go:1062
// It uses two recursive calls to expNN for x**y mod m1 and x**y mod m2
//line /usr/local/go/src/math/big/nat.go:1062
// and then uses the Chinese Remainder Theorem to combine the results.
//line /usr/local/go/src/math/big/nat.go:1062
// The recursive call using m1 will use expNNWindowed,
//line /usr/local/go/src/math/big/nat.go:1062
// while the recursive call using m2 will use expNNMontgomery.
//line /usr/local/go/src/math/big/nat.go:1062
// For more details, see Ç. K. Koç, “Montgomery Reduction with Even Modulus”,
//line /usr/local/go/src/math/big/nat.go:1062
// IEE Proceedings: Computers and Digital Techniques, 141(5) 314-316, September 1994.
//line /usr/local/go/src/math/big/nat.go:1062
// http://www.people.vcu.edu/~jwang3/CMSC691/j34monex.pdf
//line /usr/local/go/src/math/big/nat.go:1070
func (z nat) expNNMontgomeryEven(x, y, m nat) nat {
//line /usr/local/go/src/math/big/nat.go:1070
	_go_fuzz_dep_.CoverTab[6040]++

						n := m.trailingZeroBits()
						m1 := nat(nil).shl(natOne, n)
						m2 := nat(nil).shr(m, n)

//line /usr/local/go/src/math/big/nat.go:1082
	z1 := nat(nil).expNN(x, y, m1, false)
						z2 := nat(nil).expNN(x, y, m2, false)

//line /usr/local/go/src/math/big/nat.go:1095
	z = z.set(z2)

//line /usr/local/go/src/math/big/nat.go:1098
	z1 = z1.subMod2N(z1, z2, n)

//line /usr/local/go/src/math/big/nat.go:1101
	m2inv := nat(nil).modInverse(m2, m1)
						z2 = z2.mul(z1, m2inv)
						z2 = z2.trunc(z2, n)

//line /usr/local/go/src/math/big/nat.go:1106
	z = z.add(z, z1.mul(z2, m2))

						return z
//line /usr/local/go/src/math/big/nat.go:1108
	// _ = "end of CoverTab[6040]"
}

// expNNWindowed calculates x**y mod m using a fixed, 4-bit window,
//line /usr/local/go/src/math/big/nat.go:1111
// where m = 2**logM.
//line /usr/local/go/src/math/big/nat.go:1113
func (z nat) expNNWindowed(x, y nat, logM uint) nat {
//line /usr/local/go/src/math/big/nat.go:1113
	_go_fuzz_dep_.CoverTab[6041]++
						if len(y) <= 1 {
//line /usr/local/go/src/math/big/nat.go:1114
		_go_fuzz_dep_.CoverTab[6051]++
							panic("big: misuse of expNNWindowed")
//line /usr/local/go/src/math/big/nat.go:1115
		// _ = "end of CoverTab[6051]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1116
		_go_fuzz_dep_.CoverTab[6052]++
//line /usr/local/go/src/math/big/nat.go:1116
		// _ = "end of CoverTab[6052]"
//line /usr/local/go/src/math/big/nat.go:1116
	}
//line /usr/local/go/src/math/big/nat.go:1116
	// _ = "end of CoverTab[6041]"
//line /usr/local/go/src/math/big/nat.go:1116
	_go_fuzz_dep_.CoverTab[6042]++
						if x[0]&1 == 0 {
//line /usr/local/go/src/math/big/nat.go:1117
		_go_fuzz_dep_.CoverTab[6053]++

//line /usr/local/go/src/math/big/nat.go:1120
		return z.setWord(0)
//line /usr/local/go/src/math/big/nat.go:1120
		// _ = "end of CoverTab[6053]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1121
		_go_fuzz_dep_.CoverTab[6054]++
//line /usr/local/go/src/math/big/nat.go:1121
		// _ = "end of CoverTab[6054]"
//line /usr/local/go/src/math/big/nat.go:1121
	}
//line /usr/local/go/src/math/big/nat.go:1121
	// _ = "end of CoverTab[6042]"
//line /usr/local/go/src/math/big/nat.go:1121
	_go_fuzz_dep_.CoverTab[6043]++
						if logM == 1 {
//line /usr/local/go/src/math/big/nat.go:1122
		_go_fuzz_dep_.CoverTab[6055]++
							return z.setWord(1)
//line /usr/local/go/src/math/big/nat.go:1123
		// _ = "end of CoverTab[6055]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1124
		_go_fuzz_dep_.CoverTab[6056]++
//line /usr/local/go/src/math/big/nat.go:1124
		// _ = "end of CoverTab[6056]"
//line /usr/local/go/src/math/big/nat.go:1124
	}
//line /usr/local/go/src/math/big/nat.go:1124
	// _ = "end of CoverTab[6043]"
//line /usr/local/go/src/math/big/nat.go:1124
	_go_fuzz_dep_.CoverTab[6044]++

//line /usr/local/go/src/math/big/nat.go:1128
	w := int((logM + _W - 1) / _W)
	zzp := getNat(w)
	zz := *zzp

	const n = 4
	// powers[i] contains x^i.
	var powers [1 << n]*nat
	for i := range powers {
//line /usr/local/go/src/math/big/nat.go:1135
		_go_fuzz_dep_.CoverTab[6057]++
							powers[i] = getNat(w)
//line /usr/local/go/src/math/big/nat.go:1136
		// _ = "end of CoverTab[6057]"
	}
//line /usr/local/go/src/math/big/nat.go:1137
	// _ = "end of CoverTab[6044]"
//line /usr/local/go/src/math/big/nat.go:1137
	_go_fuzz_dep_.CoverTab[6045]++
						*powers[0] = powers[0].set(natOne)
						*powers[1] = powers[1].trunc(x, logM)
						for i := 2; i < 1<<n; i += 2 {
//line /usr/local/go/src/math/big/nat.go:1140
		_go_fuzz_dep_.CoverTab[6058]++
							p2, p, p1 := powers[i/2], powers[i], powers[i+1]
							*p = p.sqr(*p2)
							*p = p.trunc(*p, logM)
							*p1 = p1.mul(*p, x)
							*p1 = p1.trunc(*p1, logM)
//line /usr/local/go/src/math/big/nat.go:1145
		// _ = "end of CoverTab[6058]"
	}
//line /usr/local/go/src/math/big/nat.go:1146
	// _ = "end of CoverTab[6045]"
//line /usr/local/go/src/math/big/nat.go:1146
	_go_fuzz_dep_.CoverTab[6046]++

//line /usr/local/go/src/math/big/nat.go:1153
	i := len(y) - 1
	mtop := int((logM - 2) / _W)
	mmask := ^Word(0)
	if mbits := (logM - 1) & (_W - 1); mbits != 0 {
//line /usr/local/go/src/math/big/nat.go:1156
		_go_fuzz_dep_.CoverTab[6059]++
							mmask = (1 << mbits) - 1
//line /usr/local/go/src/math/big/nat.go:1157
		// _ = "end of CoverTab[6059]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1158
		_go_fuzz_dep_.CoverTab[6060]++
//line /usr/local/go/src/math/big/nat.go:1158
		// _ = "end of CoverTab[6060]"
//line /usr/local/go/src/math/big/nat.go:1158
	}
//line /usr/local/go/src/math/big/nat.go:1158
	// _ = "end of CoverTab[6046]"
//line /usr/local/go/src/math/big/nat.go:1158
	_go_fuzz_dep_.CoverTab[6047]++
						if i > mtop {
//line /usr/local/go/src/math/big/nat.go:1159
		_go_fuzz_dep_.CoverTab[6061]++
							i = mtop
//line /usr/local/go/src/math/big/nat.go:1160
		// _ = "end of CoverTab[6061]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1161
		_go_fuzz_dep_.CoverTab[6062]++
//line /usr/local/go/src/math/big/nat.go:1161
		// _ = "end of CoverTab[6062]"
//line /usr/local/go/src/math/big/nat.go:1161
	}
//line /usr/local/go/src/math/big/nat.go:1161
	// _ = "end of CoverTab[6047]"
//line /usr/local/go/src/math/big/nat.go:1161
	_go_fuzz_dep_.CoverTab[6048]++
						advance := false
						z = z.setWord(1)
						for ; i >= 0; i-- {
//line /usr/local/go/src/math/big/nat.go:1164
		_go_fuzz_dep_.CoverTab[6063]++
							yi := y[i]
							if i == mtop {
//line /usr/local/go/src/math/big/nat.go:1166
			_go_fuzz_dep_.CoverTab[6065]++
								yi &= mmask
//line /usr/local/go/src/math/big/nat.go:1167
			// _ = "end of CoverTab[6065]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1168
			_go_fuzz_dep_.CoverTab[6066]++
//line /usr/local/go/src/math/big/nat.go:1168
			// _ = "end of CoverTab[6066]"
//line /usr/local/go/src/math/big/nat.go:1168
		}
//line /usr/local/go/src/math/big/nat.go:1168
		// _ = "end of CoverTab[6063]"
//line /usr/local/go/src/math/big/nat.go:1168
		_go_fuzz_dep_.CoverTab[6064]++
							for j := 0; j < _W; j += n {
//line /usr/local/go/src/math/big/nat.go:1169
			_go_fuzz_dep_.CoverTab[6067]++
								if advance {
//line /usr/local/go/src/math/big/nat.go:1170
				_go_fuzz_dep_.CoverTab[6069]++

//line /usr/local/go/src/math/big/nat.go:1175
				zz = zz.sqr(z)
									zz, z = z, zz
									z = z.trunc(z, logM)

									zz = zz.sqr(z)
									zz, z = z, zz
									z = z.trunc(z, logM)

									zz = zz.sqr(z)
									zz, z = z, zz
									z = z.trunc(z, logM)

									zz = zz.sqr(z)
									zz, z = z, zz
									z = z.trunc(z, logM)
//line /usr/local/go/src/math/big/nat.go:1189
				// _ = "end of CoverTab[6069]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1190
				_go_fuzz_dep_.CoverTab[6070]++
//line /usr/local/go/src/math/big/nat.go:1190
				// _ = "end of CoverTab[6070]"
//line /usr/local/go/src/math/big/nat.go:1190
			}
//line /usr/local/go/src/math/big/nat.go:1190
			// _ = "end of CoverTab[6067]"
//line /usr/local/go/src/math/big/nat.go:1190
			_go_fuzz_dep_.CoverTab[6068]++

								zz = zz.mul(z, *powers[yi>>(_W-n)])
								zz, z = z, zz
								z = z.trunc(z, logM)

								yi <<= n
								advance = true
//line /usr/local/go/src/math/big/nat.go:1197
			// _ = "end of CoverTab[6068]"
		}
//line /usr/local/go/src/math/big/nat.go:1198
		// _ = "end of CoverTab[6064]"
	}
//line /usr/local/go/src/math/big/nat.go:1199
	// _ = "end of CoverTab[6048]"
//line /usr/local/go/src/math/big/nat.go:1199
	_go_fuzz_dep_.CoverTab[6049]++

						*zzp = zz
						putNat(zzp)
						for i := range powers {
//line /usr/local/go/src/math/big/nat.go:1203
		_go_fuzz_dep_.CoverTab[6071]++
							putNat(powers[i])
//line /usr/local/go/src/math/big/nat.go:1204
		// _ = "end of CoverTab[6071]"
	}
//line /usr/local/go/src/math/big/nat.go:1205
	// _ = "end of CoverTab[6049]"
//line /usr/local/go/src/math/big/nat.go:1205
	_go_fuzz_dep_.CoverTab[6050]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:1207
	// _ = "end of CoverTab[6050]"
}

// expNNMontgomery calculates x**y mod m using a fixed, 4-bit window.
//line /usr/local/go/src/math/big/nat.go:1210
// Uses Montgomery representation.
//line /usr/local/go/src/math/big/nat.go:1212
func (z nat) expNNMontgomery(x, y, m nat) nat {
//line /usr/local/go/src/math/big/nat.go:1212
	_go_fuzz_dep_.CoverTab[6072]++
						numWords := len(m)

//line /usr/local/go/src/math/big/nat.go:1217
	if len(x) > numWords {
//line /usr/local/go/src/math/big/nat.go:1217
		_go_fuzz_dep_.CoverTab[6080]++
							_, x = nat(nil).div(nil, x, m)
//line /usr/local/go/src/math/big/nat.go:1218
		// _ = "end of CoverTab[6080]"

	} else {
//line /usr/local/go/src/math/big/nat.go:1220
		_go_fuzz_dep_.CoverTab[6081]++
//line /usr/local/go/src/math/big/nat.go:1220
		// _ = "end of CoverTab[6081]"
//line /usr/local/go/src/math/big/nat.go:1220
	}
//line /usr/local/go/src/math/big/nat.go:1220
	// _ = "end of CoverTab[6072]"
//line /usr/local/go/src/math/big/nat.go:1220
	_go_fuzz_dep_.CoverTab[6073]++
						if len(x) < numWords {
//line /usr/local/go/src/math/big/nat.go:1221
		_go_fuzz_dep_.CoverTab[6082]++
							rr := make(nat, numWords)
							copy(rr, x)
							x = rr
//line /usr/local/go/src/math/big/nat.go:1224
		// _ = "end of CoverTab[6082]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1225
		_go_fuzz_dep_.CoverTab[6083]++
//line /usr/local/go/src/math/big/nat.go:1225
		// _ = "end of CoverTab[6083]"
//line /usr/local/go/src/math/big/nat.go:1225
	}
//line /usr/local/go/src/math/big/nat.go:1225
	// _ = "end of CoverTab[6073]"
//line /usr/local/go/src/math/big/nat.go:1225
	_go_fuzz_dep_.CoverTab[6074]++

//line /usr/local/go/src/math/big/nat.go:1230
	k0 := 2 - m[0]
	t := m[0] - 1
	for i := 1; i < _W; i <<= 1 {
//line /usr/local/go/src/math/big/nat.go:1232
		_go_fuzz_dep_.CoverTab[6084]++
							t *= t
							k0 *= (t + 1)
//line /usr/local/go/src/math/big/nat.go:1234
		// _ = "end of CoverTab[6084]"
	}
//line /usr/local/go/src/math/big/nat.go:1235
	// _ = "end of CoverTab[6074]"
//line /usr/local/go/src/math/big/nat.go:1235
	_go_fuzz_dep_.CoverTab[6075]++
						k0 = -k0

//line /usr/local/go/src/math/big/nat.go:1239
	RR := nat(nil).setWord(1)
	zz := nat(nil).shl(RR, uint(2*numWords*_W))
	_, RR = nat(nil).div(RR, zz, m)
	if len(RR) < numWords {
//line /usr/local/go/src/math/big/nat.go:1242
		_go_fuzz_dep_.CoverTab[6085]++
							zz = zz.make(numWords)
							copy(zz, RR)
							RR = zz
//line /usr/local/go/src/math/big/nat.go:1245
		// _ = "end of CoverTab[6085]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1246
		_go_fuzz_dep_.CoverTab[6086]++
//line /usr/local/go/src/math/big/nat.go:1246
		// _ = "end of CoverTab[6086]"
//line /usr/local/go/src/math/big/nat.go:1246
	}
//line /usr/local/go/src/math/big/nat.go:1246
	// _ = "end of CoverTab[6075]"
//line /usr/local/go/src/math/big/nat.go:1246
	_go_fuzz_dep_.CoverTab[6076]++

						one := make(nat, numWords)
						one[0] = 1

						const n = 4
	// powers[i] contains x^i
	var powers [1 << n]nat
	powers[0] = powers[0].montgomery(one, RR, m, k0, numWords)
	powers[1] = powers[1].montgomery(x, RR, m, k0, numWords)
	for i := 2; i < 1<<n; i++ {
//line /usr/local/go/src/math/big/nat.go:1256
		_go_fuzz_dep_.CoverTab[6087]++
							powers[i] = powers[i].montgomery(powers[i-1], powers[1], m, k0, numWords)
//line /usr/local/go/src/math/big/nat.go:1257
		// _ = "end of CoverTab[6087]"
	}
//line /usr/local/go/src/math/big/nat.go:1258
	// _ = "end of CoverTab[6076]"
//line /usr/local/go/src/math/big/nat.go:1258
	_go_fuzz_dep_.CoverTab[6077]++

//line /usr/local/go/src/math/big/nat.go:1261
	z = z.make(numWords)
						copy(z, powers[0])

						zz = zz.make(numWords)

//line /usr/local/go/src/math/big/nat.go:1267
	for i := len(y) - 1; i >= 0; i-- {
//line /usr/local/go/src/math/big/nat.go:1267
		_go_fuzz_dep_.CoverTab[6088]++
							yi := y[i]
							for j := 0; j < _W; j += n {
//line /usr/local/go/src/math/big/nat.go:1269
			_go_fuzz_dep_.CoverTab[6089]++
								if i != len(y)-1 || func() bool {
//line /usr/local/go/src/math/big/nat.go:1270
				_go_fuzz_dep_.CoverTab[6091]++
//line /usr/local/go/src/math/big/nat.go:1270
				return j != 0
//line /usr/local/go/src/math/big/nat.go:1270
				// _ = "end of CoverTab[6091]"
//line /usr/local/go/src/math/big/nat.go:1270
			}() {
//line /usr/local/go/src/math/big/nat.go:1270
				_go_fuzz_dep_.CoverTab[6092]++
									zz = zz.montgomery(z, z, m, k0, numWords)
									z = z.montgomery(zz, zz, m, k0, numWords)
									zz = zz.montgomery(z, z, m, k0, numWords)
									z = z.montgomery(zz, zz, m, k0, numWords)
//line /usr/local/go/src/math/big/nat.go:1274
				// _ = "end of CoverTab[6092]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1275
				_go_fuzz_dep_.CoverTab[6093]++
//line /usr/local/go/src/math/big/nat.go:1275
				// _ = "end of CoverTab[6093]"
//line /usr/local/go/src/math/big/nat.go:1275
			}
//line /usr/local/go/src/math/big/nat.go:1275
			// _ = "end of CoverTab[6089]"
//line /usr/local/go/src/math/big/nat.go:1275
			_go_fuzz_dep_.CoverTab[6090]++
								zz = zz.montgomery(z, powers[yi>>(_W-n)], m, k0, numWords)
								z, zz = zz, z
								yi <<= n
//line /usr/local/go/src/math/big/nat.go:1278
			// _ = "end of CoverTab[6090]"
		}
//line /usr/local/go/src/math/big/nat.go:1279
		// _ = "end of CoverTab[6088]"
	}
//line /usr/local/go/src/math/big/nat.go:1280
	// _ = "end of CoverTab[6077]"
//line /usr/local/go/src/math/big/nat.go:1280
	_go_fuzz_dep_.CoverTab[6078]++

						zz = zz.montgomery(z, one, m, k0, numWords)

//line /usr/local/go/src/math/big/nat.go:1286
	if zz.cmp(m) >= 0 {
//line /usr/local/go/src/math/big/nat.go:1286
		_go_fuzz_dep_.CoverTab[6094]++

//line /usr/local/go/src/math/big/nat.go:1294
		zz = zz.sub(zz, m)
		if zz.cmp(m) >= 0 {
//line /usr/local/go/src/math/big/nat.go:1295
			_go_fuzz_dep_.CoverTab[6095]++
								_, zz = nat(nil).div(nil, zz, m)
//line /usr/local/go/src/math/big/nat.go:1296
			// _ = "end of CoverTab[6095]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1297
			_go_fuzz_dep_.CoverTab[6096]++
//line /usr/local/go/src/math/big/nat.go:1297
			// _ = "end of CoverTab[6096]"
//line /usr/local/go/src/math/big/nat.go:1297
		}
//line /usr/local/go/src/math/big/nat.go:1297
		// _ = "end of CoverTab[6094]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1298
		_go_fuzz_dep_.CoverTab[6097]++
//line /usr/local/go/src/math/big/nat.go:1298
		// _ = "end of CoverTab[6097]"
//line /usr/local/go/src/math/big/nat.go:1298
	}
//line /usr/local/go/src/math/big/nat.go:1298
	// _ = "end of CoverTab[6078]"
//line /usr/local/go/src/math/big/nat.go:1298
	_go_fuzz_dep_.CoverTab[6079]++

						return zz.norm()
//line /usr/local/go/src/math/big/nat.go:1300
	// _ = "end of CoverTab[6079]"
}

// bytes writes the value of z into buf using big-endian encoding.
//line /usr/local/go/src/math/big/nat.go:1303
// The value of z is encoded in the slice buf[i:]. If the value of z
//line /usr/local/go/src/math/big/nat.go:1303
// cannot be represented in buf, bytes panics. The number i of unused
//line /usr/local/go/src/math/big/nat.go:1303
// bytes at the beginning of buf is returned as result.
//line /usr/local/go/src/math/big/nat.go:1307
func (z nat) bytes(buf []byte) (i int) {
//line /usr/local/go/src/math/big/nat.go:1307
	_go_fuzz_dep_.CoverTab[6098]++

//line /usr/local/go/src/math/big/nat.go:1311
	i = len(buf)
	for _, d := range z {
//line /usr/local/go/src/math/big/nat.go:1312
		_go_fuzz_dep_.CoverTab[6102]++
							for j := 0; j < _S; j++ {
//line /usr/local/go/src/math/big/nat.go:1313
			_go_fuzz_dep_.CoverTab[6103]++
								i--
								if i >= 0 {
//line /usr/local/go/src/math/big/nat.go:1315
				_go_fuzz_dep_.CoverTab[6105]++
									buf[i] = byte(d)
//line /usr/local/go/src/math/big/nat.go:1316
				// _ = "end of CoverTab[6105]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1317
				_go_fuzz_dep_.CoverTab[6106]++
//line /usr/local/go/src/math/big/nat.go:1317
				if byte(d) != 0 {
//line /usr/local/go/src/math/big/nat.go:1317
					_go_fuzz_dep_.CoverTab[6107]++
										panic("math/big: buffer too small to fit value")
//line /usr/local/go/src/math/big/nat.go:1318
					// _ = "end of CoverTab[6107]"
				} else {
//line /usr/local/go/src/math/big/nat.go:1319
					_go_fuzz_dep_.CoverTab[6108]++
//line /usr/local/go/src/math/big/nat.go:1319
					// _ = "end of CoverTab[6108]"
//line /usr/local/go/src/math/big/nat.go:1319
				}
//line /usr/local/go/src/math/big/nat.go:1319
				// _ = "end of CoverTab[6106]"
//line /usr/local/go/src/math/big/nat.go:1319
			}
//line /usr/local/go/src/math/big/nat.go:1319
			// _ = "end of CoverTab[6103]"
//line /usr/local/go/src/math/big/nat.go:1319
			_go_fuzz_dep_.CoverTab[6104]++
								d >>= 8
//line /usr/local/go/src/math/big/nat.go:1320
			// _ = "end of CoverTab[6104]"
		}
//line /usr/local/go/src/math/big/nat.go:1321
		// _ = "end of CoverTab[6102]"
	}
//line /usr/local/go/src/math/big/nat.go:1322
	// _ = "end of CoverTab[6098]"
//line /usr/local/go/src/math/big/nat.go:1322
	_go_fuzz_dep_.CoverTab[6099]++

						if i < 0 {
//line /usr/local/go/src/math/big/nat.go:1324
		_go_fuzz_dep_.CoverTab[6109]++
							i = 0
//line /usr/local/go/src/math/big/nat.go:1325
		// _ = "end of CoverTab[6109]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1326
		_go_fuzz_dep_.CoverTab[6110]++
//line /usr/local/go/src/math/big/nat.go:1326
		// _ = "end of CoverTab[6110]"
//line /usr/local/go/src/math/big/nat.go:1326
	}
//line /usr/local/go/src/math/big/nat.go:1326
	// _ = "end of CoverTab[6099]"
//line /usr/local/go/src/math/big/nat.go:1326
	_go_fuzz_dep_.CoverTab[6100]++
						for i < len(buf) && func() bool {
//line /usr/local/go/src/math/big/nat.go:1327
		_go_fuzz_dep_.CoverTab[6111]++
//line /usr/local/go/src/math/big/nat.go:1327
		return buf[i] == 0
//line /usr/local/go/src/math/big/nat.go:1327
		// _ = "end of CoverTab[6111]"
//line /usr/local/go/src/math/big/nat.go:1327
	}() {
//line /usr/local/go/src/math/big/nat.go:1327
		_go_fuzz_dep_.CoverTab[6112]++
							i++
//line /usr/local/go/src/math/big/nat.go:1328
		// _ = "end of CoverTab[6112]"
	}
//line /usr/local/go/src/math/big/nat.go:1329
	// _ = "end of CoverTab[6100]"
//line /usr/local/go/src/math/big/nat.go:1329
	_go_fuzz_dep_.CoverTab[6101]++

						return
//line /usr/local/go/src/math/big/nat.go:1331
	// _ = "end of CoverTab[6101]"
}

// bigEndianWord returns the contents of buf interpreted as a big-endian encoded Word value.
func bigEndianWord(buf []byte) Word {
//line /usr/local/go/src/math/big/nat.go:1335
	_go_fuzz_dep_.CoverTab[6113]++
						if _W == 64 {
//line /usr/local/go/src/math/big/nat.go:1336
		_go_fuzz_dep_.CoverTab[6115]++
							return Word(binary.BigEndian.Uint64(buf))
//line /usr/local/go/src/math/big/nat.go:1337
		// _ = "end of CoverTab[6115]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1338
		_go_fuzz_dep_.CoverTab[6116]++
//line /usr/local/go/src/math/big/nat.go:1338
		// _ = "end of CoverTab[6116]"
//line /usr/local/go/src/math/big/nat.go:1338
	}
//line /usr/local/go/src/math/big/nat.go:1338
	// _ = "end of CoverTab[6113]"
//line /usr/local/go/src/math/big/nat.go:1338
	_go_fuzz_dep_.CoverTab[6114]++
						return Word(binary.BigEndian.Uint32(buf))
//line /usr/local/go/src/math/big/nat.go:1339
	// _ = "end of CoverTab[6114]"
}

// setBytes interprets buf as the bytes of a big-endian unsigned
//line /usr/local/go/src/math/big/nat.go:1342
// integer, sets z to that value, and returns z.
//line /usr/local/go/src/math/big/nat.go:1344
func (z nat) setBytes(buf []byte) nat {
//line /usr/local/go/src/math/big/nat.go:1344
	_go_fuzz_dep_.CoverTab[6117]++
						z = z.make((len(buf) + _S - 1) / _S)

						i := len(buf)
						for k := 0; i >= _S; k++ {
//line /usr/local/go/src/math/big/nat.go:1348
		_go_fuzz_dep_.CoverTab[6120]++
							z[k] = bigEndianWord(buf[i-_S : i])
							i -= _S
//line /usr/local/go/src/math/big/nat.go:1350
		// _ = "end of CoverTab[6120]"
	}
//line /usr/local/go/src/math/big/nat.go:1351
	// _ = "end of CoverTab[6117]"
//line /usr/local/go/src/math/big/nat.go:1351
	_go_fuzz_dep_.CoverTab[6118]++
						if i > 0 {
//line /usr/local/go/src/math/big/nat.go:1352
		_go_fuzz_dep_.CoverTab[6121]++
							var d Word
							for s := uint(0); i > 0; s += 8 {
//line /usr/local/go/src/math/big/nat.go:1354
			_go_fuzz_dep_.CoverTab[6123]++
								d |= Word(buf[i-1]) << s
								i--
//line /usr/local/go/src/math/big/nat.go:1356
			// _ = "end of CoverTab[6123]"
		}
//line /usr/local/go/src/math/big/nat.go:1357
		// _ = "end of CoverTab[6121]"
//line /usr/local/go/src/math/big/nat.go:1357
		_go_fuzz_dep_.CoverTab[6122]++
							z[len(z)-1] = d
//line /usr/local/go/src/math/big/nat.go:1358
		// _ = "end of CoverTab[6122]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1359
		_go_fuzz_dep_.CoverTab[6124]++
//line /usr/local/go/src/math/big/nat.go:1359
		// _ = "end of CoverTab[6124]"
//line /usr/local/go/src/math/big/nat.go:1359
	}
//line /usr/local/go/src/math/big/nat.go:1359
	// _ = "end of CoverTab[6118]"
//line /usr/local/go/src/math/big/nat.go:1359
	_go_fuzz_dep_.CoverTab[6119]++

						return z.norm()
//line /usr/local/go/src/math/big/nat.go:1361
	// _ = "end of CoverTab[6119]"
}

// sqrt sets z = ⌊√x⌋
func (z nat) sqrt(x nat) nat {
//line /usr/local/go/src/math/big/nat.go:1365
	_go_fuzz_dep_.CoverTab[6125]++
						if x.cmp(natOne) <= 0 {
//line /usr/local/go/src/math/big/nat.go:1366
		_go_fuzz_dep_.CoverTab[6128]++
							return z.set(x)
//line /usr/local/go/src/math/big/nat.go:1367
		// _ = "end of CoverTab[6128]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1368
		_go_fuzz_dep_.CoverTab[6129]++
//line /usr/local/go/src/math/big/nat.go:1368
		// _ = "end of CoverTab[6129]"
//line /usr/local/go/src/math/big/nat.go:1368
	}
//line /usr/local/go/src/math/big/nat.go:1368
	// _ = "end of CoverTab[6125]"
//line /usr/local/go/src/math/big/nat.go:1368
	_go_fuzz_dep_.CoverTab[6126]++
						if alias(z, x) {
//line /usr/local/go/src/math/big/nat.go:1369
		_go_fuzz_dep_.CoverTab[6130]++
							z = nil
//line /usr/local/go/src/math/big/nat.go:1370
		// _ = "end of CoverTab[6130]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1371
		_go_fuzz_dep_.CoverTab[6131]++
//line /usr/local/go/src/math/big/nat.go:1371
		// _ = "end of CoverTab[6131]"
//line /usr/local/go/src/math/big/nat.go:1371
	}
//line /usr/local/go/src/math/big/nat.go:1371
	// _ = "end of CoverTab[6126]"
//line /usr/local/go/src/math/big/nat.go:1371
	_go_fuzz_dep_.CoverTab[6127]++

	// Start with value known to be too large and repeat "z = ⌊(z + ⌊x/z⌋)/2⌋" until it stops getting smaller.
	// See Brent and Zimmermann, Modern Computer Arithmetic, Algorithm 1.13 (SqrtInt).
	// https://members.loria.fr/PZimmermann/mca/pub226.html
	// If x is one less than a perfect square, the sequence oscillates between the correct z and z+1;
	// otherwise it converges to the correct z and stays there.
	var z1, z2 nat
	z1 = z
	z1 = z1.setUint64(1)
	z1 = z1.shl(z1, uint(x.bitLen()+1)/2)
	for n := 0; ; n++ {
//line /usr/local/go/src/math/big/nat.go:1382
		_go_fuzz_dep_.CoverTab[6132]++
							z2, _ = z2.div(nil, x, z1)
							z2 = z2.add(z2, z1)
							z2 = z2.shr(z2, 1)
							if z2.cmp(z1) >= 0 {
//line /usr/local/go/src/math/big/nat.go:1386
			_go_fuzz_dep_.CoverTab[6134]++

//line /usr/local/go/src/math/big/nat.go:1389
			if n&1 == 0 {
//line /usr/local/go/src/math/big/nat.go:1389
				_go_fuzz_dep_.CoverTab[6136]++
									return z1
//line /usr/local/go/src/math/big/nat.go:1390
				// _ = "end of CoverTab[6136]"
			} else {
//line /usr/local/go/src/math/big/nat.go:1391
				_go_fuzz_dep_.CoverTab[6137]++
//line /usr/local/go/src/math/big/nat.go:1391
				// _ = "end of CoverTab[6137]"
//line /usr/local/go/src/math/big/nat.go:1391
			}
//line /usr/local/go/src/math/big/nat.go:1391
			// _ = "end of CoverTab[6134]"
//line /usr/local/go/src/math/big/nat.go:1391
			_go_fuzz_dep_.CoverTab[6135]++
								return z.set(z1)
//line /usr/local/go/src/math/big/nat.go:1392
			// _ = "end of CoverTab[6135]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1393
			_go_fuzz_dep_.CoverTab[6138]++
//line /usr/local/go/src/math/big/nat.go:1393
			// _ = "end of CoverTab[6138]"
//line /usr/local/go/src/math/big/nat.go:1393
		}
//line /usr/local/go/src/math/big/nat.go:1393
		// _ = "end of CoverTab[6132]"
//line /usr/local/go/src/math/big/nat.go:1393
		_go_fuzz_dep_.CoverTab[6133]++
							z1, z2 = z2, z1
//line /usr/local/go/src/math/big/nat.go:1394
		// _ = "end of CoverTab[6133]"
	}
//line /usr/local/go/src/math/big/nat.go:1395
	// _ = "end of CoverTab[6127]"
}

// subMod2N returns z = (x - y) mod 2ⁿ.
func (z nat) subMod2N(x, y nat, n uint) nat {
//line /usr/local/go/src/math/big/nat.go:1399
	_go_fuzz_dep_.CoverTab[6139]++
						if uint(x.bitLen()) > n {
//line /usr/local/go/src/math/big/nat.go:1400
		_go_fuzz_dep_.CoverTab[6145]++
							if alias(z, x) {
//line /usr/local/go/src/math/big/nat.go:1401
			_go_fuzz_dep_.CoverTab[6146]++

								x = x.trunc(x, n)
//line /usr/local/go/src/math/big/nat.go:1403
			// _ = "end of CoverTab[6146]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1404
			_go_fuzz_dep_.CoverTab[6147]++
								x = nat(nil).trunc(x, n)
//line /usr/local/go/src/math/big/nat.go:1405
			// _ = "end of CoverTab[6147]"
		}
//line /usr/local/go/src/math/big/nat.go:1406
		// _ = "end of CoverTab[6145]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1407
		_go_fuzz_dep_.CoverTab[6148]++
//line /usr/local/go/src/math/big/nat.go:1407
		// _ = "end of CoverTab[6148]"
//line /usr/local/go/src/math/big/nat.go:1407
	}
//line /usr/local/go/src/math/big/nat.go:1407
	// _ = "end of CoverTab[6139]"
//line /usr/local/go/src/math/big/nat.go:1407
	_go_fuzz_dep_.CoverTab[6140]++
						if uint(y.bitLen()) > n {
//line /usr/local/go/src/math/big/nat.go:1408
		_go_fuzz_dep_.CoverTab[6149]++
							if alias(z, y) {
//line /usr/local/go/src/math/big/nat.go:1409
			_go_fuzz_dep_.CoverTab[6150]++

								y = y.trunc(y, n)
//line /usr/local/go/src/math/big/nat.go:1411
			// _ = "end of CoverTab[6150]"
		} else {
//line /usr/local/go/src/math/big/nat.go:1412
			_go_fuzz_dep_.CoverTab[6151]++
								y = nat(nil).trunc(y, n)
//line /usr/local/go/src/math/big/nat.go:1413
			// _ = "end of CoverTab[6151]"
		}
//line /usr/local/go/src/math/big/nat.go:1414
		// _ = "end of CoverTab[6149]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1415
		_go_fuzz_dep_.CoverTab[6152]++
//line /usr/local/go/src/math/big/nat.go:1415
		// _ = "end of CoverTab[6152]"
//line /usr/local/go/src/math/big/nat.go:1415
	}
//line /usr/local/go/src/math/big/nat.go:1415
	// _ = "end of CoverTab[6140]"
//line /usr/local/go/src/math/big/nat.go:1415
	_go_fuzz_dep_.CoverTab[6141]++
						if x.cmp(y) >= 0 {
//line /usr/local/go/src/math/big/nat.go:1416
		_go_fuzz_dep_.CoverTab[6153]++
							return z.sub(x, y)
//line /usr/local/go/src/math/big/nat.go:1417
		// _ = "end of CoverTab[6153]"
	} else {
//line /usr/local/go/src/math/big/nat.go:1418
		_go_fuzz_dep_.CoverTab[6154]++
//line /usr/local/go/src/math/big/nat.go:1418
		// _ = "end of CoverTab[6154]"
//line /usr/local/go/src/math/big/nat.go:1418
	}
//line /usr/local/go/src/math/big/nat.go:1418
	// _ = "end of CoverTab[6141]"
//line /usr/local/go/src/math/big/nat.go:1418
	_go_fuzz_dep_.CoverTab[6142]++

						z = z.sub(y, x)
						for uint(len(z))*_W < n {
//line /usr/local/go/src/math/big/nat.go:1421
		_go_fuzz_dep_.CoverTab[6155]++
							z = append(z, 0)
//line /usr/local/go/src/math/big/nat.go:1422
		// _ = "end of CoverTab[6155]"
	}
//line /usr/local/go/src/math/big/nat.go:1423
	// _ = "end of CoverTab[6142]"
//line /usr/local/go/src/math/big/nat.go:1423
	_go_fuzz_dep_.CoverTab[6143]++
						for i := range z {
//line /usr/local/go/src/math/big/nat.go:1424
		_go_fuzz_dep_.CoverTab[6156]++
							z[i] = ^z[i]
//line /usr/local/go/src/math/big/nat.go:1425
		// _ = "end of CoverTab[6156]"
	}
//line /usr/local/go/src/math/big/nat.go:1426
	// _ = "end of CoverTab[6143]"
//line /usr/local/go/src/math/big/nat.go:1426
	_go_fuzz_dep_.CoverTab[6144]++
						z = z.trunc(z, n)
						return z.add(z, natOne)
//line /usr/local/go/src/math/big/nat.go:1428
	// _ = "end of CoverTab[6144]"
}

//line /usr/local/go/src/math/big/nat.go:1429
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/nat.go:1429
var _ = _go_fuzz_dep_.CoverTab
