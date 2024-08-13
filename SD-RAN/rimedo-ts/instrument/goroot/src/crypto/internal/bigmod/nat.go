// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
package bigmod

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
import (
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
import (
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:5
)

import (
	"errors"
	"math/big"
	"math/bits"
)

const (
	// _W is the number of bits we use for our limbs.
	_W	= bits.UintSize - 1
	// _MASK selects _W bits from a full machine word.
	_MASK	= (1 << _W) - 1
)

// choice represents a constant-time boolean. The value of choice is always
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:20
// either 1 or 0. We use an int instead of bool in order to make decisions in
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:20
// constant time by turning it into a mask.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:23
type choice uint

func not(c choice) choice {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:25
	_go_fuzz_dep_.CoverTab[7188]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:25
	return 1 ^ c
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:25
	// _ = "end of CoverTab[7188]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:25
}

const yes = choice(1)
const no = choice(0)

// ctSelect returns x if on == 1, and y if on == 0. The execution time of this
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:30
// function does not depend on its inputs. If on is any value besides 1 or 0,
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:30
// the result is undefined.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:33
func ctSelect(on choice, x, y uint) uint {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:33
	_go_fuzz_dep_.CoverTab[7189]++

								mask := -uint(on)

								return y ^ (mask & (y ^ x))
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:37
	// _ = "end of CoverTab[7189]"
}

// ctEq returns 1 if x == y, and 0 otherwise. The execution time of this
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:40
// function does not depend on its inputs.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:42
func ctEq(x, y uint) choice {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:42
	_go_fuzz_dep_.CoverTab[7190]++

								_, c1 := bits.Sub(x, y, 0)
								_, c2 := bits.Sub(y, x, 0)
								return not(choice(c1 | c2))
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:46
	// _ = "end of CoverTab[7190]"
}

// ctGeq returns 1 if x >= y, and 0 otherwise. The execution time of this
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:49
// function does not depend on its inputs.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:51
func ctGeq(x, y uint) choice {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:51
	_go_fuzz_dep_.CoverTab[7191]++

								_, carry := bits.Sub(x, y, 0)
								return not(choice(carry))
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:54
	// _ = "end of CoverTab[7191]"
}

// Nat represents an arbitrary natural number
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:57
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:57
// Each Nat has an announced length, which is the number of limbs it has stored.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:57
// Operations on this number are allowed to leak this length, but will not leak
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:57
// any information about the values contained in those limbs.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:62
type Nat struct {
	// limbs is a little-endian representation in base 2^W with
	// W = bits.UintSize - 1. The top bit is always unset between operations.
	//
	// The top bit is left unset to optimize Montgomery multiplication, in the
	// inner loop of exponentiation. Using fully saturated limbs would leave us
	// working with 129-bit numbers on 64-bit platforms, wasting a lot of space,
	// and thus time.
	limbs []uint
}

// preallocTarget is the size in bits of the numbers used to implement the most
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:73
// common and most performant RSA key size. It's also enough to cover some of
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:73
// the operations of key sizes up to 4096.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:76
const preallocTarget = 2048
const preallocLimbs = (preallocTarget + _W - 1) / _W

// NewNat returns a new nat with a size of zero, just like new(Nat), but with
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:79
// the preallocated capacity to hold a number of up to preallocTarget bits.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:79
// NewNat inlines, so the allocation can live on the stack.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:82
func NewNat() *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:82
	_go_fuzz_dep_.CoverTab[7192]++
								limbs := make([]uint, 0, preallocLimbs)
								return &Nat{limbs}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:84
	// _ = "end of CoverTab[7192]"
}

// expand expands x to n limbs, leaving its value unchanged.
func (x *Nat) expand(n int) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:88
	_go_fuzz_dep_.CoverTab[7193]++
								if len(x.limbs) > n {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:89
		_go_fuzz_dep_.CoverTab[7197]++
									panic("bigmod: internal error: shrinking nat")
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:90
		// _ = "end of CoverTab[7197]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:91
		_go_fuzz_dep_.CoverTab[7198]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:91
		// _ = "end of CoverTab[7198]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:91
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:91
	// _ = "end of CoverTab[7193]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:91
	_go_fuzz_dep_.CoverTab[7194]++
								if cap(x.limbs) < n {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:92
		_go_fuzz_dep_.CoverTab[7199]++
									newLimbs := make([]uint, n)
									copy(newLimbs, x.limbs)
									x.limbs = newLimbs
									return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:96
		// _ = "end of CoverTab[7199]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:97
		_go_fuzz_dep_.CoverTab[7200]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:97
		// _ = "end of CoverTab[7200]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:97
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:97
	// _ = "end of CoverTab[7194]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:97
	_go_fuzz_dep_.CoverTab[7195]++
								extraLimbs := x.limbs[len(x.limbs):n]
								for i := range extraLimbs {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:99
		_go_fuzz_dep_.CoverTab[7201]++
									extraLimbs[i] = 0
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:100
		// _ = "end of CoverTab[7201]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:101
	// _ = "end of CoverTab[7195]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:101
	_go_fuzz_dep_.CoverTab[7196]++
								x.limbs = x.limbs[:n]
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:103
	// _ = "end of CoverTab[7196]"
}

// reset returns a zero nat of n limbs, reusing x's storage if n <= cap(x.limbs).
func (x *Nat) reset(n int) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:107
	_go_fuzz_dep_.CoverTab[7202]++
								if cap(x.limbs) < n {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:108
		_go_fuzz_dep_.CoverTab[7205]++
									x.limbs = make([]uint, n)
									return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:110
		// _ = "end of CoverTab[7205]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:111
		_go_fuzz_dep_.CoverTab[7206]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:111
		// _ = "end of CoverTab[7206]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:111
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:111
	// _ = "end of CoverTab[7202]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:111
	_go_fuzz_dep_.CoverTab[7203]++
								for i := range x.limbs {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:112
		_go_fuzz_dep_.CoverTab[7207]++
									x.limbs[i] = 0
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:113
		// _ = "end of CoverTab[7207]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:114
	// _ = "end of CoverTab[7203]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:114
	_go_fuzz_dep_.CoverTab[7204]++
								x.limbs = x.limbs[:n]
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:116
	// _ = "end of CoverTab[7204]"
}

// set assigns x = y, optionally resizing x to the appropriate size.
func (x *Nat) set(y *Nat) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:120
	_go_fuzz_dep_.CoverTab[7208]++
								x.reset(len(y.limbs))
								copy(x.limbs, y.limbs)
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:123
	// _ = "end of CoverTab[7208]"
}

// setBig assigns x = n, optionally resizing n to the appropriate size.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:126
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:126
// The announced length of x is set based on the actual bit size of the input,
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:126
// ignoring leading zeroes.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:130
func (x *Nat) setBig(n *big.Int) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:130
	_go_fuzz_dep_.CoverTab[7209]++
								requiredLimbs := (n.BitLen() + _W - 1) / _W
								x.reset(requiredLimbs)

								outI := 0
								shift := 0
								limbs := n.Bits()
								for i := range limbs {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:137
		_go_fuzz_dep_.CoverTab[7211]++
									xi := uint(limbs[i])
									x.limbs[outI] |= (xi << shift) & _MASK
									outI++
									if outI == requiredLimbs {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:141
			_go_fuzz_dep_.CoverTab[7213]++
										return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:142
			// _ = "end of CoverTab[7213]"
		} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:143
			_go_fuzz_dep_.CoverTab[7214]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:143
			// _ = "end of CoverTab[7214]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:143
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:143
		// _ = "end of CoverTab[7211]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:143
		_go_fuzz_dep_.CoverTab[7212]++
									x.limbs[outI] = xi >> (_W - shift)
									shift++
									if shift == _W {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:146
			_go_fuzz_dep_.CoverTab[7215]++
										shift = 0
										outI++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:148
			// _ = "end of CoverTab[7215]"
		} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:149
			_go_fuzz_dep_.CoverTab[7216]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:149
			// _ = "end of CoverTab[7216]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:149
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:149
		// _ = "end of CoverTab[7212]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:150
	// _ = "end of CoverTab[7209]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:150
	_go_fuzz_dep_.CoverTab[7210]++
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:151
	// _ = "end of CoverTab[7210]"
}

// Bytes returns x as a zero-extended big-endian byte slice. The size of the
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:154
// slice will match the size of m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:154
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:154
// x must have the same size as m and it must be reduced modulo m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:158
func (x *Nat) Bytes(m *Modulus) []byte {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:158
	_go_fuzz_dep_.CoverTab[7217]++
								bytes := make([]byte, m.Size())
								shift := 0
								outI := len(bytes) - 1
								for _, limb := range x.limbs {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:162
		_go_fuzz_dep_.CoverTab[7219]++
									remainingBits := _W
									for remainingBits >= 8 {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:164
			_go_fuzz_dep_.CoverTab[7221]++
										bytes[outI] |= byte(limb) << shift
										consumed := 8 - shift
										limb >>= consumed
										remainingBits -= consumed
										shift = 0
										outI--
										if outI < 0 {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:171
				_go_fuzz_dep_.CoverTab[7222]++
											return bytes
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:172
				// _ = "end of CoverTab[7222]"
			} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:173
				_go_fuzz_dep_.CoverTab[7223]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:173
				// _ = "end of CoverTab[7223]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:173
			}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:173
			// _ = "end of CoverTab[7221]"
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:174
		// _ = "end of CoverTab[7219]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:174
		_go_fuzz_dep_.CoverTab[7220]++
									bytes[outI] = byte(limb)
									shift = remainingBits
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:176
		// _ = "end of CoverTab[7220]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:177
	// _ = "end of CoverTab[7217]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:177
	_go_fuzz_dep_.CoverTab[7218]++
								return bytes
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:178
	// _ = "end of CoverTab[7218]"
}

// SetBytes assigns x = b, where b is a slice of big-endian bytes.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:181
// SetBytes returns an error if b >= m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:181
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:181
// The output will be resized to the size of m and overwritten.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:185
func (x *Nat) SetBytes(b []byte, m *Modulus) (*Nat, error) {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:185
	_go_fuzz_dep_.CoverTab[7224]++
								if err := x.setBytes(b, m); err != nil {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:186
		_go_fuzz_dep_.CoverTab[7227]++
									return nil, err
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:187
		// _ = "end of CoverTab[7227]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:188
		_go_fuzz_dep_.CoverTab[7228]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:188
		// _ = "end of CoverTab[7228]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:188
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:188
	// _ = "end of CoverTab[7224]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:188
	_go_fuzz_dep_.CoverTab[7225]++
								if x.cmpGeq(m.nat) == yes {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:189
		_go_fuzz_dep_.CoverTab[7229]++
									return nil, errors.New("input overflows the modulus")
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:190
		// _ = "end of CoverTab[7229]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:191
		_go_fuzz_dep_.CoverTab[7230]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:191
		// _ = "end of CoverTab[7230]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:191
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:191
	// _ = "end of CoverTab[7225]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:191
	_go_fuzz_dep_.CoverTab[7226]++
								return x, nil
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:192
	// _ = "end of CoverTab[7226]"
}

// SetOverflowingBytes assigns x = b, where b is a slice of big-endian bytes. SetOverflowingBytes
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:195
// returns an error if b has a longer bit length than m, but reduces overflowing
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:195
// values up to 2^⌈log2(m)⌉ - 1.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:195
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:195
// The output will be resized to the size of m and overwritten.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:200
func (x *Nat) SetOverflowingBytes(b []byte, m *Modulus) (*Nat, error) {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:200
	_go_fuzz_dep_.CoverTab[7231]++
								if err := x.setBytes(b, m); err != nil {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:201
		_go_fuzz_dep_.CoverTab[7234]++
									return nil, err
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:202
		// _ = "end of CoverTab[7234]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:203
		_go_fuzz_dep_.CoverTab[7235]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:203
		// _ = "end of CoverTab[7235]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:203
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:203
	// _ = "end of CoverTab[7231]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:203
	_go_fuzz_dep_.CoverTab[7232]++
								leading := _W - bitLen(x.limbs[len(x.limbs)-1])
								if leading < m.leading {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:205
		_go_fuzz_dep_.CoverTab[7236]++
									return nil, errors.New("input overflows the modulus")
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:206
		// _ = "end of CoverTab[7236]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:207
		_go_fuzz_dep_.CoverTab[7237]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:207
		// _ = "end of CoverTab[7237]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:207
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:207
	// _ = "end of CoverTab[7232]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:207
	_go_fuzz_dep_.CoverTab[7233]++
								x.sub(x.cmpGeq(m.nat), m.nat)
								return x, nil
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:209
	// _ = "end of CoverTab[7233]"
}

func (x *Nat) setBytes(b []byte, m *Modulus) error {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:212
	_go_fuzz_dep_.CoverTab[7238]++
								outI := 0
								shift := 0
								x.resetFor(m)
								for i := len(b) - 1; i >= 0; i-- {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:216
		_go_fuzz_dep_.CoverTab[7240]++
									bi := b[i]
									x.limbs[outI] |= uint(bi) << shift
									shift += 8
									if shift >= _W {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:220
			_go_fuzz_dep_.CoverTab[7241]++
										shift -= _W
										x.limbs[outI] &= _MASK
										overflow := bi >> (8 - shift)
										outI++
										if outI >= len(x.limbs) {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:225
				_go_fuzz_dep_.CoverTab[7243]++
											if overflow > 0 || func() bool {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:226
					_go_fuzz_dep_.CoverTab[7245]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:226
					return i > 0
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:226
					// _ = "end of CoverTab[7245]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:226
				}() {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:226
					_go_fuzz_dep_.CoverTab[7246]++
												return errors.New("input overflows the modulus")
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:227
					// _ = "end of CoverTab[7246]"
				} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:228
					_go_fuzz_dep_.CoverTab[7247]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:228
					// _ = "end of CoverTab[7247]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:228
				}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:228
				// _ = "end of CoverTab[7243]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:228
				_go_fuzz_dep_.CoverTab[7244]++
											break
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:229
				// _ = "end of CoverTab[7244]"
			} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:230
				_go_fuzz_dep_.CoverTab[7248]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:230
				// _ = "end of CoverTab[7248]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:230
			}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:230
			// _ = "end of CoverTab[7241]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:230
			_go_fuzz_dep_.CoverTab[7242]++
										x.limbs[outI] = uint(overflow)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:231
			// _ = "end of CoverTab[7242]"
		} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:232
			_go_fuzz_dep_.CoverTab[7249]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:232
			// _ = "end of CoverTab[7249]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:232
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:232
		// _ = "end of CoverTab[7240]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:233
	// _ = "end of CoverTab[7238]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:233
	_go_fuzz_dep_.CoverTab[7239]++
								return nil
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:234
	// _ = "end of CoverTab[7239]"
}

// Equal returns 1 if x == y, and 0 otherwise.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:237
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:237
// Both operands must have the same announced length.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:240
func (x *Nat) Equal(y *Nat) choice {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:240
	_go_fuzz_dep_.CoverTab[7250]++

								size := len(x.limbs)
								xLimbs := x.limbs[:size]
								yLimbs := y.limbs[:size]

								equal := yes
								for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:247
		_go_fuzz_dep_.CoverTab[7252]++
									equal &= ctEq(xLimbs[i], yLimbs[i])
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:248
		// _ = "end of CoverTab[7252]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:249
	// _ = "end of CoverTab[7250]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:249
	_go_fuzz_dep_.CoverTab[7251]++
								return equal
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:250
	// _ = "end of CoverTab[7251]"
}

// IsZero returns 1 if x == 0, and 0 otherwise.
func (x *Nat) IsZero() choice {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:254
	_go_fuzz_dep_.CoverTab[7253]++

								size := len(x.limbs)
								xLimbs := x.limbs[:size]

								zero := yes
								for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:260
		_go_fuzz_dep_.CoverTab[7255]++
									zero &= ctEq(xLimbs[i], 0)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:261
		// _ = "end of CoverTab[7255]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:262
	// _ = "end of CoverTab[7253]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:262
	_go_fuzz_dep_.CoverTab[7254]++
								return zero
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:263
	// _ = "end of CoverTab[7254]"
}

// cmpGeq returns 1 if x >= y, and 0 otherwise.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:266
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:266
// Both operands must have the same announced length.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:269
func (x *Nat) cmpGeq(y *Nat) choice {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:269
	_go_fuzz_dep_.CoverTab[7256]++

								size := len(x.limbs)
								xLimbs := x.limbs[:size]
								yLimbs := y.limbs[:size]

								var c uint
								for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:276
		_go_fuzz_dep_.CoverTab[7258]++
									c = (xLimbs[i] - yLimbs[i] - c) >> _W
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:277
		// _ = "end of CoverTab[7258]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:278
	// _ = "end of CoverTab[7256]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:278
	_go_fuzz_dep_.CoverTab[7257]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:281
	return not(choice(c))
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:281
	// _ = "end of CoverTab[7257]"
}

// assign sets x <- y if on == 1, and does nothing otherwise.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:284
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:284
// Both operands must have the same announced length.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:287
func (x *Nat) assign(on choice, y *Nat) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:287
	_go_fuzz_dep_.CoverTab[7259]++

								size := len(x.limbs)
								xLimbs := x.limbs[:size]
								yLimbs := y.limbs[:size]

								for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:293
		_go_fuzz_dep_.CoverTab[7261]++
									xLimbs[i] = ctSelect(on, yLimbs[i], xLimbs[i])
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:294
		// _ = "end of CoverTab[7261]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:295
	// _ = "end of CoverTab[7259]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:295
	_go_fuzz_dep_.CoverTab[7260]++
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:296
	// _ = "end of CoverTab[7260]"
}

// add computes x += y if on == 1, and does nothing otherwise. It returns the
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:299
// carry of the addition regardless of on.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:299
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:299
// Both operands must have the same announced length.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:303
func (x *Nat) add(on choice, y *Nat) (c uint) {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:303
	_go_fuzz_dep_.CoverTab[7262]++

								size := len(x.limbs)
								xLimbs := x.limbs[:size]
								yLimbs := y.limbs[:size]

								for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:309
		_go_fuzz_dep_.CoverTab[7264]++
									res := xLimbs[i] + yLimbs[i] + c
									xLimbs[i] = ctSelect(on, res&_MASK, xLimbs[i])
									c = res >> _W
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:312
		// _ = "end of CoverTab[7264]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:313
	// _ = "end of CoverTab[7262]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:313
	_go_fuzz_dep_.CoverTab[7263]++
								return
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:314
	// _ = "end of CoverTab[7263]"
}

// sub computes x -= y if on == 1, and does nothing otherwise. It returns the
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:317
// borrow of the subtraction regardless of on.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:317
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:317
// Both operands must have the same announced length.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:321
func (x *Nat) sub(on choice, y *Nat) (c uint) {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:321
	_go_fuzz_dep_.CoverTab[7265]++

								size := len(x.limbs)
								xLimbs := x.limbs[:size]
								yLimbs := y.limbs[:size]

								for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:327
		_go_fuzz_dep_.CoverTab[7267]++
									res := xLimbs[i] - yLimbs[i] - c
									xLimbs[i] = ctSelect(on, res&_MASK, xLimbs[i])
									c = res >> _W
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:330
		// _ = "end of CoverTab[7267]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:331
	// _ = "end of CoverTab[7265]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:331
	_go_fuzz_dep_.CoverTab[7266]++
								return
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:332
	// _ = "end of CoverTab[7266]"
}

// Modulus is used for modular arithmetic, precomputing relevant constants.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:335
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:335
// Moduli are assumed to be odd numbers. Moduli can also leak the exact
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:335
// number of bits needed to store their value, and are stored without padding.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:335
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:335
// Their actual value is still kept secret.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:341
type Modulus struct {
	// The underlying natural number for this modulus.
	//
	// This will be stored without any padding, and shouldn't alias with any
	// other natural number being used.
	nat	*Nat
	leading	int	// number of leading zeros in the modulus
	m0inv	uint	// -nat.limbs[0]⁻¹ mod _W
	rr	*Nat	// R*R for montgomeryRepresentation
}

// rr returns R*R with R = 2^(_W * n) and n = len(m.nat.limbs).
func rr(m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:353
	_go_fuzz_dep_.CoverTab[7268]++
								rr := NewNat().ExpandFor(m)

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:358
	n := len(rr.limbs)
	rr.limbs[n-1] = 1
	for i := n - 1; i < 2*n; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:360
		_go_fuzz_dep_.CoverTab[7270]++
									rr.shiftIn(0, m)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:361
		// _ = "end of CoverTab[7270]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:362
	// _ = "end of CoverTab[7268]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:362
	_go_fuzz_dep_.CoverTab[7269]++
								return rr
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:363
	// _ = "end of CoverTab[7269]"
}

// minusInverseModW computes -x⁻¹ mod _W with x odd.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:366
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:366
// This operation is used to precompute a constant involved in Montgomery
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:366
// multiplication.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:370
func minusInverseModW(x uint) uint {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:370
	_go_fuzz_dep_.CoverTab[7271]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:377
	y := x
	for i := 0; i < 5; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:378
		_go_fuzz_dep_.CoverTab[7273]++
									y = y * (2 - x*y)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:379
		// _ = "end of CoverTab[7273]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:380
	// _ = "end of CoverTab[7271]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:380
	_go_fuzz_dep_.CoverTab[7272]++
								return (1 << _W) - (y & _MASK)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:381
	// _ = "end of CoverTab[7272]"
}

// NewModulusFromBig creates a new Modulus from a [big.Int].
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:384
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:384
// The Int must be odd. The number of significant bits must be leakable.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:387
func NewModulusFromBig(n *big.Int) *Modulus {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:387
	_go_fuzz_dep_.CoverTab[7274]++
								m := &Modulus{}
								m.nat = NewNat().setBig(n)
								m.leading = _W - bitLen(m.nat.limbs[len(m.nat.limbs)-1])
								m.m0inv = minusInverseModW(m.nat.limbs[0])
								m.rr = rr(m)
								return m
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:393
	// _ = "end of CoverTab[7274]"
}

// bitLen is a version of bits.Len that only leaks the bit length of n, but not
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:396
// its value. bits.Len and bits.LeadingZeros use a lookup table for the
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:396
// low-order bits on some architectures.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:399
func bitLen(n uint) int {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:399
	_go_fuzz_dep_.CoverTab[7275]++
								var len int

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:403
	for n != 0 {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:403
		_go_fuzz_dep_.CoverTab[7277]++
									len++
									n >>= 1
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:405
		// _ = "end of CoverTab[7277]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:406
	// _ = "end of CoverTab[7275]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:406
	_go_fuzz_dep_.CoverTab[7276]++
								return len
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:407
	// _ = "end of CoverTab[7276]"
}

// Size returns the size of m in bytes.
func (m *Modulus) Size() int {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:411
	_go_fuzz_dep_.CoverTab[7278]++
								return (m.BitLen() + 7) / 8
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:412
	// _ = "end of CoverTab[7278]"
}

// BitLen returns the size of m in bits.
func (m *Modulus) BitLen() int {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:416
	_go_fuzz_dep_.CoverTab[7279]++
								return len(m.nat.limbs)*_W - int(m.leading)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:417
	// _ = "end of CoverTab[7279]"
}

// Nat returns m as a Nat. The return value must not be written to.
func (m *Modulus) Nat() *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:421
	_go_fuzz_dep_.CoverTab[7280]++
								return m.nat
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:422
	// _ = "end of CoverTab[7280]"
}

// shiftIn calculates x = x << _W + y mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:425
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:425
// This assumes that x is already reduced mod m, and that y < 2^_W.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:428
func (x *Nat) shiftIn(y uint, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:428
	_go_fuzz_dep_.CoverTab[7281]++
								d := NewNat().resetFor(m)

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:432
	size := len(m.nat.limbs)
								xLimbs := x.limbs[:size]
								dLimbs := d.limbs[:size]
								mLimbs := m.nat.limbs[:size]

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:444
	needSubtraction := no
	for i := _W - 1; i >= 0; i-- {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:445
		_go_fuzz_dep_.CoverTab[7283]++
									carry := (y >> i) & 1
									var borrow uint
									for i := 0; i < size; i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:448
			_go_fuzz_dep_.CoverTab[7285]++
										l := ctSelect(needSubtraction, dLimbs[i], xLimbs[i])

										res := l<<1 + carry
										xLimbs[i] = res & _MASK
										carry = res >> _W

										res = xLimbs[i] - mLimbs[i] - borrow
										dLimbs[i] = res & _MASK
										borrow = res >> _W
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:457
			// _ = "end of CoverTab[7285]"
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:458
		// _ = "end of CoverTab[7283]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:458
		_go_fuzz_dep_.CoverTab[7284]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:461
		needSubtraction = ctEq(carry, borrow)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:461
		// _ = "end of CoverTab[7284]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:462
	// _ = "end of CoverTab[7281]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:462
	_go_fuzz_dep_.CoverTab[7282]++
								return x.assign(needSubtraction, d)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:463
	// _ = "end of CoverTab[7282]"
}

// Mod calculates out = x mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:466
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:466
// This works regardless how large the value of x is.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:466
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:466
// The output will be resized to the size of m and overwritten.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:471
func (out *Nat) Mod(x *Nat, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:471
	_go_fuzz_dep_.CoverTab[7286]++
								out.resetFor(m)

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:478
	i := len(x.limbs) - 1

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:481
	start := len(m.nat.limbs) - 2
	if i < start {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:482
		_go_fuzz_dep_.CoverTab[7290]++
									start = i
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:483
		// _ = "end of CoverTab[7290]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:484
		_go_fuzz_dep_.CoverTab[7291]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:484
		// _ = "end of CoverTab[7291]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:484
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:484
	// _ = "end of CoverTab[7286]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:484
	_go_fuzz_dep_.CoverTab[7287]++
								for j := start; j >= 0; j-- {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:485
		_go_fuzz_dep_.CoverTab[7292]++
									out.limbs[j] = x.limbs[i]
									i--
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:487
		// _ = "end of CoverTab[7292]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:488
	// _ = "end of CoverTab[7287]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:488
	_go_fuzz_dep_.CoverTab[7288]++

								for i >= 0 {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:490
		_go_fuzz_dep_.CoverTab[7293]++
									out.shiftIn(x.limbs[i], m)
									i--
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:492
		// _ = "end of CoverTab[7293]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:493
	// _ = "end of CoverTab[7288]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:493
	_go_fuzz_dep_.CoverTab[7289]++
								return out
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:494
	// _ = "end of CoverTab[7289]"
}

// ExpandFor ensures out has the right size to work with operations modulo m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:497
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:497
// The announced size of out must be smaller than or equal to that of m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:500
func (out *Nat) ExpandFor(m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:500
	_go_fuzz_dep_.CoverTab[7294]++
								return out.expand(len(m.nat.limbs))
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:501
	// _ = "end of CoverTab[7294]"
}

// resetFor ensures out has the right size to work with operations modulo m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:504
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:504
// out is zeroed and may start at any size.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:507
func (out *Nat) resetFor(m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:507
	_go_fuzz_dep_.CoverTab[7295]++
								return out.reset(len(m.nat.limbs))
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:508
	// _ = "end of CoverTab[7295]"
}

// Sub computes x = x - y mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:511
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:511
// The length of both operands must be the same as the modulus. Both operands
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:511
// must already be reduced modulo m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:515
func (x *Nat) Sub(y *Nat, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:515
	_go_fuzz_dep_.CoverTab[7296]++
								underflow := x.sub(yes, y)

								x.add(choice(underflow), m.nat)
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:519
	// _ = "end of CoverTab[7296]"
}

// Add computes x = x + y mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:522
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:522
// The length of both operands must be the same as the modulus. Both operands
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:522
// must already be reduced modulo m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:526
func (x *Nat) Add(y *Nat, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:526
	_go_fuzz_dep_.CoverTab[7297]++
								overflow := x.add(yes, y)
								underflow := not(x.cmpGeq(m.nat))

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:552
	needSubtraction := ctEq(overflow, uint(underflow))

								x.sub(needSubtraction, m.nat)
								return x
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:555
	// _ = "end of CoverTab[7297]"
}

// montgomeryRepresentation calculates x = x * R mod m, with R = 2^(_W * n) and
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:558
// n = len(m.nat.limbs).
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:558
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:558
// Faster Montgomery multiplication replaces standard modular multiplication for
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:558
// numbers in this representation.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:558
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:558
// This assumes that x is already reduced mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:565
func (x *Nat) montgomeryRepresentation(m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:565
	_go_fuzz_dep_.CoverTab[7298]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:568
	return x.montgomeryMul(NewNat().set(x), m.rr, m)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:568
	// _ = "end of CoverTab[7298]"
}

// montgomeryReduction calculates x = x / R mod m, with R = 2^(_W * n) and
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:571
// n = len(m.nat.limbs).
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:571
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:571
// This assumes that x is already reduced mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:575
func (x *Nat) montgomeryReduction(m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:575
	_go_fuzz_dep_.CoverTab[7299]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:579
	t0 := NewNat().set(x)
								t1 := NewNat().ExpandFor(m)
								t1.limbs[0] = 1
								return x.montgomeryMul(t0, t1, m)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:582
	// _ = "end of CoverTab[7299]"
}

// montgomeryMul calculates d = a * b / R mod m, with R = 2^(_W * n) and
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:585
// n = len(m.nat.limbs), using the Montgomery Multiplication technique.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:585
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:585
// All inputs should be the same length, not aliasing d, and already
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:585
// reduced modulo m. d will be resized to the size of m and overwritten.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:590
func (d *Nat) montgomeryMul(a *Nat, b *Nat, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:590
	_go_fuzz_dep_.CoverTab[7300]++
								d.resetFor(m)
								if len(a.limbs) != len(m.nat.limbs) || func() bool {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:592
		_go_fuzz_dep_.CoverTab[7302]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:592
		return len(b.limbs) != len(m.nat.limbs)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:592
		// _ = "end of CoverTab[7302]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:592
	}() {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:592
		_go_fuzz_dep_.CoverTab[7303]++
									panic("bigmod: invalid montgomeryMul input")
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:593
		// _ = "end of CoverTab[7303]"
	} else {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:594
		_go_fuzz_dep_.CoverTab[7304]++
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:594
		// _ = "end of CoverTab[7304]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:594
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:594
	// _ = "end of CoverTab[7300]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:594
	_go_fuzz_dep_.CoverTab[7301]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:599
	overflow := montgomeryLoop(d.limbs, a.limbs, b.limbs, m.nat.limbs, m.m0inv)
								underflow := not(d.cmpGeq(m.nat))
								needSubtraction := ctEq(overflow, uint(underflow))
								d.sub(needSubtraction, m.nat)

								return d
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:604
	// _ = "end of CoverTab[7301]"
}

func montgomeryLoopGeneric(d, a, b, m []uint, m0inv uint) (overflow uint) {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:607
	_go_fuzz_dep_.CoverTab[7305]++

								size := len(d)
								a = a[:size]
								b = b[:size]
								m = m[:size]

								for _, ai := range a {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:614
		_go_fuzz_dep_.CoverTab[7307]++

									hi, lo := bits.Mul(ai, b[0])
									z_lo, c := bits.Add(d[0], lo, 0)
									f := (z_lo * m0inv) & _MASK
									z_hi, _ := bits.Add(0, hi, c)
									hi, lo = bits.Mul(f, m[0])
									z_lo, c = bits.Add(z_lo, lo, 0)
									z_hi, _ = bits.Add(z_hi, hi, c)
									carry := z_hi<<1 | z_lo>>_W

									for j := 1; j < size; j++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:625
			_go_fuzz_dep_.CoverTab[7309]++

										hi, lo := bits.Mul(ai, b[j])
										z_lo, c := bits.Add(d[j], lo, 0)
										z_hi, _ := bits.Add(0, hi, c)
										hi, lo = bits.Mul(f, m[j])
										z_lo, c = bits.Add(z_lo, lo, 0)
										z_hi, _ = bits.Add(z_hi, hi, c)
										z_lo, c = bits.Add(z_lo, carry, 0)
										z_hi, _ = bits.Add(z_hi, 0, c)
										d[j-1] = z_lo & _MASK
										carry = z_hi<<1 | z_lo>>_W
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:636
			// _ = "end of CoverTab[7309]"
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:637
		// _ = "end of CoverTab[7307]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:637
		_go_fuzz_dep_.CoverTab[7308]++

									z := overflow + carry
									d[size-1] = z & _MASK
									overflow = z >> _W
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:641
		// _ = "end of CoverTab[7308]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:642
	// _ = "end of CoverTab[7305]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:642
	_go_fuzz_dep_.CoverTab[7306]++
								return
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:643
	// _ = "end of CoverTab[7306]"
}

// Mul calculates x *= y mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:646
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:646
// x and y must already be reduced modulo m, they must share its announced
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:646
// length, and they may not alias.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:650
func (x *Nat) Mul(y *Nat, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:650
	_go_fuzz_dep_.CoverTab[7310]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:653
	xR := NewNat().set(x).montgomeryRepresentation(m)
								return x.montgomeryMul(xR, y, m)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:654
	// _ = "end of CoverTab[7310]"
}

// Exp calculates out = x^e mod m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:657
//
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:657
// The exponent e is represented in big-endian order. The output will be resized
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:657
// to the size of m and overwritten. x must already be reduced modulo m.
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:661
func (out *Nat) Exp(x *Nat, e []byte, m *Modulus) *Nat {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:661
	_go_fuzz_dep_.CoverTab[7311]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:666
	table := [(1 << 4) - 1]*Nat{

		NewNat(), NewNat(), NewNat(), NewNat(), NewNat(),
		NewNat(), NewNat(), NewNat(), NewNat(), NewNat(),
		NewNat(), NewNat(), NewNat(), NewNat(), NewNat(),
	}
	table[0].set(x).montgomeryRepresentation(m)
	for i := 1; i < len(table); i++ {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:673
		_go_fuzz_dep_.CoverTab[7314]++
									table[i].montgomeryMul(table[i-1], table[0], m)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:674
		// _ = "end of CoverTab[7314]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:675
	// _ = "end of CoverTab[7311]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:675
	_go_fuzz_dep_.CoverTab[7312]++

								out.resetFor(m)
								out.limbs[0] = 1
								out.montgomeryRepresentation(m)
								t0 := NewNat().ExpandFor(m)
								t1 := NewNat().ExpandFor(m)
								for _, b := range e {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:682
		_go_fuzz_dep_.CoverTab[7315]++
									for _, j := range []int{4, 0} {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:683
			_go_fuzz_dep_.CoverTab[7316]++

										t1.montgomeryMul(out, out, m)
										out.montgomeryMul(t1, t1, m)
										t1.montgomeryMul(out, out, m)
										out.montgomeryMul(t1, t1, m)

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:691
			k := uint((b >> j) & 0b1111)
			for i := range table {
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:692
				_go_fuzz_dep_.CoverTab[7318]++
											t0.assign(ctEq(k, uint(i+1)), table[i])
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:693
				// _ = "end of CoverTab[7318]"
			}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:694
			// _ = "end of CoverTab[7316]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:694
			_go_fuzz_dep_.CoverTab[7317]++

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:697
			t1.montgomeryMul(out, t0, m)
										out.assign(not(ctEq(k, 0)), t1)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:698
			// _ = "end of CoverTab[7317]"
		}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:699
		// _ = "end of CoverTab[7315]"
	}
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:700
	// _ = "end of CoverTab[7312]"
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:700
	_go_fuzz_dep_.CoverTab[7313]++

								return out.montgomeryReduction(m)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:702
	// _ = "end of CoverTab[7313]"
}

//line /usr/local/go/src/crypto/internal/bigmod/nat.go:703
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/bigmod/nat.go:703
var _ = _go_fuzz_dep_.CoverTab
