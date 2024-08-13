// Copyright (c) 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
package edwards25519

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
import (
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:5
)

import (
	"encoding/binary"
	"errors"
)

// A Scalar is an integer modulo
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
//	l = 2^252 + 27742317777372353535851937790883648493
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
// which is the prime order of the edwards25519 group.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
// This type works similarly to math/big.Int, and all arguments and
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
// receivers are allowed to alias.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:12
// The zero value is a valid zero element.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:22
type Scalar struct {
	// s is the scalar in the Montgomery domain, in the format of the
	// fiat-crypto implementation.
	s fiatScalarMontgomeryDomainFieldElement
}

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:56
// NewScalar returns a new zero Scalar.
func NewScalar() *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:57
	_go_fuzz_dep_.CoverTab[9193]++
									return &Scalar{}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:58
	// _ = "end of CoverTab[9193]"
}

// MultiplyAdd sets s = x * y + z mod l, and returns s. It is equivalent to
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:61
// using Multiply and then Add.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:63
func (s *Scalar) MultiplyAdd(x, y, z *Scalar) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:63
	_go_fuzz_dep_.CoverTab[9194]++

									zCopy := new(Scalar).Set(z)
									return s.Multiply(x, y).Add(s, zCopy)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:66
	// _ = "end of CoverTab[9194]"
}

// Add sets s = x + y mod l, and returns s.
func (s *Scalar) Add(x, y *Scalar) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:70
	_go_fuzz_dep_.CoverTab[9195]++

									fiatScalarAdd(&s.s, &x.s, &y.s)
									return s
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:73
	// _ = "end of CoverTab[9195]"
}

// Subtract sets s = x - y mod l, and returns s.
func (s *Scalar) Subtract(x, y *Scalar) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:77
	_go_fuzz_dep_.CoverTab[9196]++

									fiatScalarSub(&s.s, &x.s, &y.s)
									return s
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:80
	// _ = "end of CoverTab[9196]"
}

// Negate sets s = -x mod l, and returns s.
func (s *Scalar) Negate(x *Scalar) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:84
	_go_fuzz_dep_.CoverTab[9197]++

									fiatScalarOpp(&s.s, &x.s)
									return s
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:87
	// _ = "end of CoverTab[9197]"
}

// Multiply sets s = x * y mod l, and returns s.
func (s *Scalar) Multiply(x, y *Scalar) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:91
	_go_fuzz_dep_.CoverTab[9198]++

									fiatScalarMul(&s.s, &x.s, &y.s)
									return s
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:94
	// _ = "end of CoverTab[9198]"
}

// Set sets s = x, and returns s.
func (s *Scalar) Set(x *Scalar) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:98
	_go_fuzz_dep_.CoverTab[9199]++
									*s = *x
									return s
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:100
	// _ = "end of CoverTab[9199]"
}

// SetUniformBytes sets s = x mod l, where x is a 64-byte little-endian integer.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:103
// If x is not of the right length, SetUniformBytes returns nil and an error,
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:103
// and the receiver is unchanged.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:103
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:103
// SetUniformBytes can be used to set s to an uniformly distributed value given
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:103
// 64 uniformly distributed random bytes.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:109
func (s *Scalar) SetUniformBytes(x []byte) (*Scalar, error) {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:109
	_go_fuzz_dep_.CoverTab[9200]++
									if len(x) != 64 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:110
		_go_fuzz_dep_.CoverTab[9202]++
										return nil, errors.New("edwards25519: invalid SetUniformBytes input length")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:111
		// _ = "end of CoverTab[9202]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:112
		_go_fuzz_dep_.CoverTab[9203]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:112
		// _ = "end of CoverTab[9203]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:112
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:112
	// _ = "end of CoverTab[9200]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:112
	_go_fuzz_dep_.CoverTab[9201]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:125
	s.setShortBytes(x[:21])
									t := new(Scalar).setShortBytes(x[21:42])
									s.Add(s, t.Multiply(t, scalarTwo168))
									t.setShortBytes(x[42:])
									s.Add(s, t.Multiply(t, scalarTwo336))

									return s, nil
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:131
	// _ = "end of CoverTab[9201]"
}

// scalarTwo168 and scalarTwo336 are 2^168 and 2^336 modulo l, encoded as a
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:134
// fiatScalarMontgomeryDomainFieldElement, which is a little-endian 4-limb value
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:134
// in the 2^256 Montgomery domain.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:137
var scalarTwo168 = &Scalar{s: [4]uint64{0x5b8ab432eac74798, 0x38afddd6de59d5d7,
	0xa2c131b399411b7c, 0x6329a7ed9ce5a30}}
var scalarTwo336 = &Scalar{s: [4]uint64{0xbd3d108e2b35ecc5, 0x5c3a3718bdf9c90b,
	0x63aa97a331b4f2ee, 0x3d217f5be65cb5c}}

// setShortBytes sets s = x mod l, where x is a little-endian integer shorter
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:142
// than 32 bytes.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:144
func (s *Scalar) setShortBytes(x []byte) *Scalar {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:144
	_go_fuzz_dep_.CoverTab[9204]++
									if len(x) >= 32 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:145
		_go_fuzz_dep_.CoverTab[9206]++
										panic("edwards25519: internal error: setShortBytes called with a long string")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:146
		// _ = "end of CoverTab[9206]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:147
		_go_fuzz_dep_.CoverTab[9207]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:147
		// _ = "end of CoverTab[9207]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:147
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:147
	// _ = "end of CoverTab[9204]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:147
	_go_fuzz_dep_.CoverTab[9205]++
									var buf [32]byte
									copy(buf[:], x)
									fiatScalarFromBytes((*[4]uint64)(&s.s), &buf)
									fiatScalarToMontgomery(&s.s, (*fiatScalarNonMontgomeryDomainFieldElement)(&s.s))
									return s
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:152
	// _ = "end of CoverTab[9205]"
}

// SetCanonicalBytes sets s = x, where x is a 32-byte little-endian encoding of
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:155
// s, and returns s. If x is not a canonical encoding of s, SetCanonicalBytes
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:155
// returns nil and an error, and the receiver is unchanged.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:158
func (s *Scalar) SetCanonicalBytes(x []byte) (*Scalar, error) {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:158
	_go_fuzz_dep_.CoverTab[9208]++
									if len(x) != 32 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:159
		_go_fuzz_dep_.CoverTab[9211]++
										return nil, errors.New("invalid scalar length")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:160
		// _ = "end of CoverTab[9211]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:161
		_go_fuzz_dep_.CoverTab[9212]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:161
		// _ = "end of CoverTab[9212]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:161
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:161
	// _ = "end of CoverTab[9208]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:161
	_go_fuzz_dep_.CoverTab[9209]++
									if !isReduced(x) {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:162
		_go_fuzz_dep_.CoverTab[9213]++
										return nil, errors.New("invalid scalar encoding")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:163
		// _ = "end of CoverTab[9213]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:164
		_go_fuzz_dep_.CoverTab[9214]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:164
		// _ = "end of CoverTab[9214]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:164
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:164
	// _ = "end of CoverTab[9209]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:164
	_go_fuzz_dep_.CoverTab[9210]++

									fiatScalarFromBytes((*[4]uint64)(&s.s), (*[32]byte)(x))
									fiatScalarToMontgomery(&s.s, (*fiatScalarNonMontgomeryDomainFieldElement)(&s.s))

									return s, nil
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:169
	// _ = "end of CoverTab[9210]"
}

// scalarMinusOneBytes is l - 1 in little endian.
var scalarMinusOneBytes = [32]byte{236, 211, 245, 92, 26, 99, 18, 88, 214, 156, 247, 162, 222, 249, 222, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16}

// isReduced returns whether the given scalar in 32-byte little endian encoded
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:175
// form is reduced modulo l.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:177
func isReduced(s []byte) bool {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:177
	_go_fuzz_dep_.CoverTab[9215]++
									if len(s) != 32 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:178
		_go_fuzz_dep_.CoverTab[9218]++
										return false
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:179
		// _ = "end of CoverTab[9218]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:180
		_go_fuzz_dep_.CoverTab[9219]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:180
		// _ = "end of CoverTab[9219]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:180
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:180
	// _ = "end of CoverTab[9215]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:180
	_go_fuzz_dep_.CoverTab[9216]++

									for i := len(s) - 1; i >= 0; i-- {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:182
		_go_fuzz_dep_.CoverTab[9220]++
										switch {
		case s[i] > scalarMinusOneBytes[i]:
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:184
			_go_fuzz_dep_.CoverTab[9221]++
											return false
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:185
			// _ = "end of CoverTab[9221]"
		case s[i] < scalarMinusOneBytes[i]:
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:186
			_go_fuzz_dep_.CoverTab[9222]++
											return true
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:187
			// _ = "end of CoverTab[9222]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:187
		default:
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:187
			_go_fuzz_dep_.CoverTab[9223]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:187
			// _ = "end of CoverTab[9223]"
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:188
		// _ = "end of CoverTab[9220]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:189
	// _ = "end of CoverTab[9216]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:189
	_go_fuzz_dep_.CoverTab[9217]++
									return true
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:190
	// _ = "end of CoverTab[9217]"
}

// SetBytesWithClamping applies the buffer pruning described in RFC 8032,
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// Section 5.1.5 (also known as clamping) and sets s to the result. The input
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// must be 32 bytes, and it is not modified. If x is not of the right length,
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// SetBytesWithClamping returns nil and an error, and the receiver is unchanged.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// Note that since Scalar values are always reduced modulo the prime order of
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// the curve, the resulting value will not preserve any of the cofactor-clearing
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// properties that clamping is meant to provide. It will however work as
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// expected as long as it is applied to points on the prime order subgroup, like
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// in Ed25519. In fact, it is lost to history why RFC 8032 adopted the
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:193
// irrelevant RFC 7748 clamping, but it is now required for compatibility.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:204
func (s *Scalar) SetBytesWithClamping(x []byte) (*Scalar, error) {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:204
	_go_fuzz_dep_.CoverTab[9224]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:209
	if len(x) != 32 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:209
		_go_fuzz_dep_.CoverTab[9226]++
										return nil, errors.New("edwards25519: invalid SetBytesWithClamping input length")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:210
		// _ = "end of CoverTab[9226]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:211
		_go_fuzz_dep_.CoverTab[9227]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:211
		// _ = "end of CoverTab[9227]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:211
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:211
	// _ = "end of CoverTab[9224]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:211
	_go_fuzz_dep_.CoverTab[9225]++

	// We need to use the wide reduction from SetUniformBytes, since clamping
									// sets the 2^254 bit, making the value higher than the order.
									var wideBytes [64]byte
									copy(wideBytes[:], x[:])
									wideBytes[0] &= 248
									wideBytes[31] &= 63
									wideBytes[31] |= 64
									return s.SetUniformBytes(wideBytes[:])
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:220
	// _ = "end of CoverTab[9225]"
}

// Bytes returns the canonical 32-byte little-endian encoding of s.
func (s *Scalar) Bytes() []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:224
	_go_fuzz_dep_.CoverTab[9228]++
	// This function is outlined to make the allocations inline in the caller
									// rather than happen on the heap.
									var encoded [32]byte
									return s.bytes(&encoded)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:228
	// _ = "end of CoverTab[9228]"
}

func (s *Scalar) bytes(out *[32]byte) []byte {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:231
	_go_fuzz_dep_.CoverTab[9229]++
									var ss fiatScalarNonMontgomeryDomainFieldElement
									fiatScalarFromMontgomery(&ss, &s.s)
									fiatScalarToBytes(out, (*[4]uint64)(&ss))
									return out[:]
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:235
	// _ = "end of CoverTab[9229]"
}

// Equal returns 1 if s and t are equal, and 0 otherwise.
func (s *Scalar) Equal(t *Scalar) int {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:239
	_go_fuzz_dep_.CoverTab[9230]++
									var diff fiatScalarMontgomeryDomainFieldElement
									fiatScalarSub(&diff, &s.s, &t.s)
									var nonzero uint64
									fiatScalarNonzero(&nonzero, (*[4]uint64)(&diff))
									nonzero |= nonzero >> 32
									nonzero |= nonzero >> 16
									nonzero |= nonzero >> 8
									nonzero |= nonzero >> 4
									nonzero |= nonzero >> 2
									nonzero |= nonzero >> 1
									return int(^nonzero) & 1
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:250
	// _ = "end of CoverTab[9230]"
}

// nonAdjacentForm computes a width-w non-adjacent form for this scalar.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:253
//
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:253
// w must be between 2 and 8, or nonAdjacentForm will panic.
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:256
func (s *Scalar) nonAdjacentForm(w uint) [256]int8 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:256
	_go_fuzz_dep_.CoverTab[9231]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:260
	b := s.Bytes()
	if b[31] > 127 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:261
		_go_fuzz_dep_.CoverTab[9236]++
										panic("scalar has high bit set illegally")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:262
		// _ = "end of CoverTab[9236]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:263
		_go_fuzz_dep_.CoverTab[9237]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:263
		// _ = "end of CoverTab[9237]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:263
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:263
	// _ = "end of CoverTab[9231]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:263
	_go_fuzz_dep_.CoverTab[9232]++
									if w < 2 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:264
		_go_fuzz_dep_.CoverTab[9238]++
										panic("w must be at least 2 by the definition of NAF")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:265
		// _ = "end of CoverTab[9238]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:266
		_go_fuzz_dep_.CoverTab[9239]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:266
		if w > 8 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:266
			_go_fuzz_dep_.CoverTab[9240]++
											panic("NAF digits must fit in int8")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:267
			// _ = "end of CoverTab[9240]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
			_go_fuzz_dep_.CoverTab[9241]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
			// _ = "end of CoverTab[9241]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
		// _ = "end of CoverTab[9239]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
	// _ = "end of CoverTab[9232]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:268
	_go_fuzz_dep_.CoverTab[9233]++

									var naf [256]int8
									var digits [5]uint64

									for i := 0; i < 4; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:273
		_go_fuzz_dep_.CoverTab[9242]++
										digits[i] = binary.LittleEndian.Uint64(b[i*8:])
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:274
		// _ = "end of CoverTab[9242]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:275
	// _ = "end of CoverTab[9233]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:275
	_go_fuzz_dep_.CoverTab[9234]++

									width := uint64(1 << w)
									windowMask := uint64(width - 1)

									pos := uint(0)
									carry := uint64(0)
									for pos < 256 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:282
		_go_fuzz_dep_.CoverTab[9243]++
										indexU64 := pos / 64
										indexBit := pos % 64
										var bitBuf uint64
										if indexBit < 64-w {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:286
			_go_fuzz_dep_.CoverTab[9247]++

											bitBuf = digits[indexU64] >> indexBit
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:288
			// _ = "end of CoverTab[9247]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:289
			_go_fuzz_dep_.CoverTab[9248]++

											bitBuf = (digits[indexU64] >> indexBit) | (digits[1+indexU64] << (64 - indexBit))
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:291
			// _ = "end of CoverTab[9248]"
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:292
		// _ = "end of CoverTab[9243]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:292
		_go_fuzz_dep_.CoverTab[9244]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:295
		window := carry + (bitBuf & windowMask)

		if window&1 == 0 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:297
			_go_fuzz_dep_.CoverTab[9249]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:304
			pos += 1
											continue
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:305
			// _ = "end of CoverTab[9249]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:306
			_go_fuzz_dep_.CoverTab[9250]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:306
			// _ = "end of CoverTab[9250]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:306
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:306
		// _ = "end of CoverTab[9244]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:306
		_go_fuzz_dep_.CoverTab[9245]++

										if window < width/2 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:308
			_go_fuzz_dep_.CoverTab[9251]++
											carry = 0
											naf[pos] = int8(window)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:310
			// _ = "end of CoverTab[9251]"
		} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:311
			_go_fuzz_dep_.CoverTab[9252]++
											carry = 1
											naf[pos] = int8(window) - int8(width)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:313
			// _ = "end of CoverTab[9252]"
		}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:314
		// _ = "end of CoverTab[9245]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:314
		_go_fuzz_dep_.CoverTab[9246]++

										pos += w
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:316
		// _ = "end of CoverTab[9246]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:317
	// _ = "end of CoverTab[9234]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:317
	_go_fuzz_dep_.CoverTab[9235]++
									return naf
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:318
	// _ = "end of CoverTab[9235]"
}

func (s *Scalar) signedRadix16() [64]int8 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:321
	_go_fuzz_dep_.CoverTab[9253]++
									b := s.Bytes()
									if b[31] > 127 {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:323
		_go_fuzz_dep_.CoverTab[9257]++
										panic("scalar has high bit set illegally")
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:324
		// _ = "end of CoverTab[9257]"
	} else {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:325
		_go_fuzz_dep_.CoverTab[9258]++
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:325
		// _ = "end of CoverTab[9258]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:325
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:325
	// _ = "end of CoverTab[9253]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:325
	_go_fuzz_dep_.CoverTab[9254]++

									var digits [64]int8

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:330
	for i := 0; i < 32; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:330
		_go_fuzz_dep_.CoverTab[9259]++
										digits[2*i] = int8(b[i] & 15)
										digits[2*i+1] = int8((b[i] >> 4) & 15)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:332
		// _ = "end of CoverTab[9259]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:333
	// _ = "end of CoverTab[9254]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:333
	_go_fuzz_dep_.CoverTab[9255]++

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:336
	for i := 0; i < 63; i++ {
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:336
		_go_fuzz_dep_.CoverTab[9260]++
										carry := (digits[i] + 8) >> 4
										digits[i] -= carry << 4
										digits[i+1] += carry
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:339
		// _ = "end of CoverTab[9260]"
	}
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:340
	// _ = "end of CoverTab[9255]"
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:340
	_go_fuzz_dep_.CoverTab[9256]++

									return digits
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:342
	// _ = "end of CoverTab[9256]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:343
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/scalar.go:343
var _ = _go_fuzz_dep_.CoverTab
