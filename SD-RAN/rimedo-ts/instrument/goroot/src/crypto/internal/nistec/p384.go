// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate.go. DO NOT EDIT.

//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
package nistec

//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
import (
//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
import (
//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:7
)

import (
	"crypto/internal/nistec/fiat"
	"crypto/subtle"
	"errors"
	"sync"
)

// p384ElementLength is the length of an element of the base or scalar field,
//line /usr/local/go/src/crypto/internal/nistec/p384.go:16
// which have the same bytes length for all NIST P curves.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:18
const p384ElementLength = 48

// P384Point is a P384 point. The zero value is NOT valid.
type P384Point struct {
	// The point is represented in projective coordinates (X:Y:Z),
	// where x = X/Z and y = Y/Z.
	x, y, z *fiat.P384Element
}

// NewP384Point returns a new P384Point representing the point at infinity point.
func NewP384Point() *P384Point {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:28
	_go_fuzz_dep_.CoverTab[2583]++
								return &P384Point{
		x:	new(fiat.P384Element),
		y:	new(fiat.P384Element).One(),
		z:	new(fiat.P384Element),
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:33
	// _ = "end of CoverTab[2583]"
}

// SetGenerator sets p to the canonical generator and returns p.
func (p *P384Point) SetGenerator() *P384Point {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:37
	_go_fuzz_dep_.CoverTab[2584]++
								p.x.SetBytes([]byte{0xaa, 0x87, 0xca, 0x22, 0xbe, 0x8b, 0x5, 0x37, 0x8e, 0xb1, 0xc7, 0x1e, 0xf3, 0x20, 0xad, 0x74, 0x6e, 0x1d, 0x3b, 0x62, 0x8b, 0xa7, 0x9b, 0x98, 0x59, 0xf7, 0x41, 0xe0, 0x82, 0x54, 0x2a, 0x38, 0x55, 0x2, 0xf2, 0x5d, 0xbf, 0x55, 0x29, 0x6c, 0x3a, 0x54, 0x5e, 0x38, 0x72, 0x76, 0xa, 0xb7})
								p.y.SetBytes([]byte{0x36, 0x17, 0xde, 0x4a, 0x96, 0x26, 0x2c, 0x6f, 0x5d, 0x9e, 0x98, 0xbf, 0x92, 0x92, 0xdc, 0x29, 0xf8, 0xf4, 0x1d, 0xbd, 0x28, 0x9a, 0x14, 0x7c, 0xe9, 0xda, 0x31, 0x13, 0xb5, 0xf0, 0xb8, 0xc0, 0xa, 0x60, 0xb1, 0xce, 0x1d, 0x7e, 0x81, 0x9d, 0x7a, 0x43, 0x1d, 0x7c, 0x90, 0xea, 0xe, 0x5f})
								p.z.One()
								return p
//line /usr/local/go/src/crypto/internal/nistec/p384.go:41
	// _ = "end of CoverTab[2584]"
}

// Set sets p = q and returns p.
func (p *P384Point) Set(q *P384Point) *P384Point {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:45
	_go_fuzz_dep_.CoverTab[2585]++
								p.x.Set(q.x)
								p.y.Set(q.y)
								p.z.Set(q.z)
								return p
//line /usr/local/go/src/crypto/internal/nistec/p384.go:49
	// _ = "end of CoverTab[2585]"
}

// SetBytes sets p to the compressed, uncompressed, or infinity value encoded in
//line /usr/local/go/src/crypto/internal/nistec/p384.go:52
// b, as specified in SEC 1, Version 2.0, Section 2.3.4. If the point is not on
//line /usr/local/go/src/crypto/internal/nistec/p384.go:52
// the curve, it returns nil and an error, and the receiver is unchanged.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:52
// Otherwise, it returns p.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:56
func (p *P384Point) SetBytes(b []byte) (*P384Point, error) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:56
	_go_fuzz_dep_.CoverTab[2586]++
								switch {

	case len(b) == 1 && func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:59
		_go_fuzz_dep_.CoverTab[2596]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:59
		return b[0] == 0
//line /usr/local/go/src/crypto/internal/nistec/p384.go:59
		// _ = "end of CoverTab[2596]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:59
	}():
//line /usr/local/go/src/crypto/internal/nistec/p384.go:59
		_go_fuzz_dep_.CoverTab[2587]++
									return p.Set(NewP384Point()), nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:60
		// _ = "end of CoverTab[2587]"

//line /usr/local/go/src/crypto/internal/nistec/p384.go:63
	case len(b) == 1+2*p384ElementLength && func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:63
		_go_fuzz_dep_.CoverTab[2597]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:63
		return b[0] == 4
//line /usr/local/go/src/crypto/internal/nistec/p384.go:63
		// _ = "end of CoverTab[2597]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:63
	}():
//line /usr/local/go/src/crypto/internal/nistec/p384.go:63
		_go_fuzz_dep_.CoverTab[2588]++
									x, err := new(fiat.P384Element).SetBytes(b[1 : 1+p384ElementLength])
									if err != nil {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:65
			_go_fuzz_dep_.CoverTab[2598]++
										return nil, err
//line /usr/local/go/src/crypto/internal/nistec/p384.go:66
			// _ = "end of CoverTab[2598]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:67
			_go_fuzz_dep_.CoverTab[2599]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:67
			// _ = "end of CoverTab[2599]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:67
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:67
		// _ = "end of CoverTab[2588]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:67
		_go_fuzz_dep_.CoverTab[2589]++
									y, err := new(fiat.P384Element).SetBytes(b[1+p384ElementLength:])
									if err != nil {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:69
			_go_fuzz_dep_.CoverTab[2600]++
										return nil, err
//line /usr/local/go/src/crypto/internal/nistec/p384.go:70
			// _ = "end of CoverTab[2600]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:71
			_go_fuzz_dep_.CoverTab[2601]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:71
			// _ = "end of CoverTab[2601]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:71
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:71
		// _ = "end of CoverTab[2589]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:71
		_go_fuzz_dep_.CoverTab[2590]++
									if err := p384CheckOnCurve(x, y); err != nil {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:72
			_go_fuzz_dep_.CoverTab[2602]++
										return nil, err
//line /usr/local/go/src/crypto/internal/nistec/p384.go:73
			// _ = "end of CoverTab[2602]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:74
			_go_fuzz_dep_.CoverTab[2603]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:74
			// _ = "end of CoverTab[2603]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:74
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:74
		// _ = "end of CoverTab[2590]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:74
		_go_fuzz_dep_.CoverTab[2591]++
									p.x.Set(x)
									p.y.Set(y)
									p.z.One()
									return p, nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:78
		// _ = "end of CoverTab[2591]"

//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
	case len(b) == 1+p384ElementLength && func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
		_go_fuzz_dep_.CoverTab[2604]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
		return (b[0] == 2 || func() bool {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
			_go_fuzz_dep_.CoverTab[2605]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
			return b[0] == 3
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
			// _ = "end of CoverTab[2605]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
		}())
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
		// _ = "end of CoverTab[2604]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
	}():
//line /usr/local/go/src/crypto/internal/nistec/p384.go:81
		_go_fuzz_dep_.CoverTab[2592]++
									x, err := new(fiat.P384Element).SetBytes(b[1:])
									if err != nil {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:83
			_go_fuzz_dep_.CoverTab[2606]++
										return nil, err
//line /usr/local/go/src/crypto/internal/nistec/p384.go:84
			// _ = "end of CoverTab[2606]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:85
			_go_fuzz_dep_.CoverTab[2607]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:85
			// _ = "end of CoverTab[2607]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:85
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:85
		// _ = "end of CoverTab[2592]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:85
		_go_fuzz_dep_.CoverTab[2593]++

//line /usr/local/go/src/crypto/internal/nistec/p384.go:88
		y := p384Polynomial(new(fiat.P384Element), x)
		if !p384Sqrt(y, y) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:89
			_go_fuzz_dep_.CoverTab[2608]++
										return nil, errors.New("invalid P384 compressed point encoding")
//line /usr/local/go/src/crypto/internal/nistec/p384.go:90
			// _ = "end of CoverTab[2608]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:91
			_go_fuzz_dep_.CoverTab[2609]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:91
			// _ = "end of CoverTab[2609]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:91
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:91
		// _ = "end of CoverTab[2593]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:91
		_go_fuzz_dep_.CoverTab[2594]++

//line /usr/local/go/src/crypto/internal/nistec/p384.go:95
		otherRoot := new(fiat.P384Element)
									otherRoot.Sub(otherRoot, y)
									cond := y.Bytes()[p384ElementLength-1]&1 ^ b[0]&1
									y.Select(otherRoot, y, int(cond))

									p.x.Set(x)
									p.y.Set(y)
									p.z.One()
									return p, nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:103
		// _ = "end of CoverTab[2594]"

	default:
//line /usr/local/go/src/crypto/internal/nistec/p384.go:105
		_go_fuzz_dep_.CoverTab[2595]++
									return nil, errors.New("invalid P384 point encoding")
//line /usr/local/go/src/crypto/internal/nistec/p384.go:106
		// _ = "end of CoverTab[2595]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:107
	// _ = "end of CoverTab[2586]"
}

var _p384B *fiat.P384Element
var _p384BOnce sync.Once

func p384B() *fiat.P384Element {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:113
	_go_fuzz_dep_.CoverTab[2610]++
								_p384BOnce.Do(func() {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:114
		_go_fuzz_dep_.CoverTab[2612]++
									_p384B, _ = new(fiat.P384Element).SetBytes([]byte{0xb3, 0x31, 0x2f, 0xa7, 0xe2, 0x3e, 0xe7, 0xe4, 0x98, 0x8e, 0x5, 0x6b, 0xe3, 0xf8, 0x2d, 0x19, 0x18, 0x1d, 0x9c, 0x6e, 0xfe, 0x81, 0x41, 0x12, 0x3, 0x14, 0x8, 0x8f, 0x50, 0x13, 0x87, 0x5a, 0xc6, 0x56, 0x39, 0x8d, 0x8a, 0x2e, 0xd1, 0x9d, 0x2a, 0x85, 0xc8, 0xed, 0xd3, 0xec, 0x2a, 0xef})
//line /usr/local/go/src/crypto/internal/nistec/p384.go:115
		// _ = "end of CoverTab[2612]"
	})
//line /usr/local/go/src/crypto/internal/nistec/p384.go:116
	// _ = "end of CoverTab[2610]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:116
	_go_fuzz_dep_.CoverTab[2611]++
								return _p384B
//line /usr/local/go/src/crypto/internal/nistec/p384.go:117
	// _ = "end of CoverTab[2611]"
}

// p384Polynomial sets y2 to x³ - 3x + b, and returns y2.
func p384Polynomial(y2, x *fiat.P384Element) *fiat.P384Element {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:121
	_go_fuzz_dep_.CoverTab[2613]++
								y2.Square(x)
								y2.Mul(y2, x)

								threeX := new(fiat.P384Element).Add(x, x)
								threeX.Add(threeX, x)
								y2.Sub(y2, threeX)

								return y2.Add(y2, p384B())
//line /usr/local/go/src/crypto/internal/nistec/p384.go:129
	// _ = "end of CoverTab[2613]"
}

func p384CheckOnCurve(x, y *fiat.P384Element) error {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:132
	_go_fuzz_dep_.CoverTab[2614]++

								rhs := p384Polynomial(new(fiat.P384Element), x)
								lhs := new(fiat.P384Element).Square(y)
								if rhs.Equal(lhs) != 1 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:136
		_go_fuzz_dep_.CoverTab[2616]++
									return errors.New("P384 point not on curve")
//line /usr/local/go/src/crypto/internal/nistec/p384.go:137
		// _ = "end of CoverTab[2616]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:138
		_go_fuzz_dep_.CoverTab[2617]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:138
		// _ = "end of CoverTab[2617]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:138
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:138
	// _ = "end of CoverTab[2614]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:138
	_go_fuzz_dep_.CoverTab[2615]++
								return nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:139
	// _ = "end of CoverTab[2615]"
}

// Bytes returns the uncompressed or infinity encoding of p, as specified in
//line /usr/local/go/src/crypto/internal/nistec/p384.go:142
// SEC 1, Version 2.0, Section 2.3.3. Note that the encoding of the point at
//line /usr/local/go/src/crypto/internal/nistec/p384.go:142
// infinity is shorter than all other encodings.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:145
func (p *P384Point) Bytes() []byte {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:145
	_go_fuzz_dep_.CoverTab[2618]++
	// This function is outlined to make the allocations inline in the caller
								// rather than happen on the heap.
								var out [1 + 2*p384ElementLength]byte
								return p.bytes(&out)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:149
	// _ = "end of CoverTab[2618]"
}

func (p *P384Point) bytes(out *[1 + 2*p384ElementLength]byte) []byte {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:152
	_go_fuzz_dep_.CoverTab[2619]++
								if p.z.IsZero() == 1 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:153
		_go_fuzz_dep_.CoverTab[2621]++
									return append(out[:0], 0)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:154
		// _ = "end of CoverTab[2621]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:155
		_go_fuzz_dep_.CoverTab[2622]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:155
		// _ = "end of CoverTab[2622]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:155
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:155
	// _ = "end of CoverTab[2619]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:155
	_go_fuzz_dep_.CoverTab[2620]++

								zinv := new(fiat.P384Element).Invert(p.z)
								x := new(fiat.P384Element).Mul(p.x, zinv)
								y := new(fiat.P384Element).Mul(p.y, zinv)

								buf := append(out[:0], 4)
								buf = append(buf, x.Bytes()...)
								buf = append(buf, y.Bytes()...)
								return buf
//line /usr/local/go/src/crypto/internal/nistec/p384.go:164
	// _ = "end of CoverTab[2620]"
}

// BytesX returns the encoding of the x-coordinate of p, as specified in SEC 1,
//line /usr/local/go/src/crypto/internal/nistec/p384.go:167
// Version 2.0, Section 2.3.5, or an error if p is the point at infinity.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:169
func (p *P384Point) BytesX() ([]byte, error) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:169
	_go_fuzz_dep_.CoverTab[2623]++
	// This function is outlined to make the allocations inline in the caller
								// rather than happen on the heap.
								var out [p384ElementLength]byte
								return p.bytesX(&out)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:173
	// _ = "end of CoverTab[2623]"
}

func (p *P384Point) bytesX(out *[p384ElementLength]byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:176
	_go_fuzz_dep_.CoverTab[2624]++
								if p.z.IsZero() == 1 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:177
		_go_fuzz_dep_.CoverTab[2626]++
									return nil, errors.New("P384 point is the point at infinity")
//line /usr/local/go/src/crypto/internal/nistec/p384.go:178
		// _ = "end of CoverTab[2626]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:179
		_go_fuzz_dep_.CoverTab[2627]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:179
		// _ = "end of CoverTab[2627]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:179
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:179
	// _ = "end of CoverTab[2624]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:179
	_go_fuzz_dep_.CoverTab[2625]++

								zinv := new(fiat.P384Element).Invert(p.z)
								x := new(fiat.P384Element).Mul(p.x, zinv)

								return append(out[:0], x.Bytes()...), nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:184
	// _ = "end of CoverTab[2625]"
}

// BytesCompressed returns the compressed or infinity encoding of p, as
//line /usr/local/go/src/crypto/internal/nistec/p384.go:187
// specified in SEC 1, Version 2.0, Section 2.3.3. Note that the encoding of the
//line /usr/local/go/src/crypto/internal/nistec/p384.go:187
// point at infinity is shorter than all other encodings.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:190
func (p *P384Point) BytesCompressed() []byte {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:190
	_go_fuzz_dep_.CoverTab[2628]++
	// This function is outlined to make the allocations inline in the caller
								// rather than happen on the heap.
								var out [1 + p384ElementLength]byte
								return p.bytesCompressed(&out)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:194
	// _ = "end of CoverTab[2628]"
}

func (p *P384Point) bytesCompressed(out *[1 + p384ElementLength]byte) []byte {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:197
	_go_fuzz_dep_.CoverTab[2629]++
								if p.z.IsZero() == 1 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:198
		_go_fuzz_dep_.CoverTab[2631]++
									return append(out[:0], 0)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:199
		// _ = "end of CoverTab[2631]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:200
		_go_fuzz_dep_.CoverTab[2632]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:200
		// _ = "end of CoverTab[2632]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:200
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:200
	// _ = "end of CoverTab[2629]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:200
	_go_fuzz_dep_.CoverTab[2630]++

								zinv := new(fiat.P384Element).Invert(p.z)
								x := new(fiat.P384Element).Mul(p.x, zinv)
								y := new(fiat.P384Element).Mul(p.y, zinv)

//line /usr/local/go/src/crypto/internal/nistec/p384.go:208
	buf := append(out[:0], 2)
								buf[0] |= y.Bytes()[p384ElementLength-1] & 1
								buf = append(buf, x.Bytes()...)
								return buf
//line /usr/local/go/src/crypto/internal/nistec/p384.go:211
	// _ = "end of CoverTab[2630]"
}

// Add sets q = p1 + p2, and returns q. The points may overlap.
func (q *P384Point) Add(p1, p2 *P384Point) *P384Point {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:215
	_go_fuzz_dep_.CoverTab[2633]++

//line /usr/local/go/src/crypto/internal/nistec/p384.go:219
	t0 := new(fiat.P384Element).Mul(p1.x, p2.x)
								t1 := new(fiat.P384Element).Mul(p1.y, p2.y)
								t2 := new(fiat.P384Element).Mul(p1.z, p2.z)
								t3 := new(fiat.P384Element).Add(p1.x, p1.y)
								t4 := new(fiat.P384Element).Add(p2.x, p2.y)
								t3.Mul(t3, t4)
								t4.Add(t0, t1)
								t3.Sub(t3, t4)
								t4.Add(p1.y, p1.z)
								x3 := new(fiat.P384Element).Add(p2.y, p2.z)
								t4.Mul(t4, x3)
								x3.Add(t1, t2)
								t4.Sub(t4, x3)
								x3.Add(p1.x, p1.z)
								y3 := new(fiat.P384Element).Add(p2.x, p2.z)
								x3.Mul(x3, y3)
								y3.Add(t0, t2)
								y3.Sub(x3, y3)
								z3 := new(fiat.P384Element).Mul(p384B(), t2)
								x3.Sub(y3, z3)
								z3.Add(x3, x3)
								x3.Add(x3, z3)
								z3.Sub(t1, x3)
								x3.Add(t1, x3)
								y3.Mul(p384B(), y3)
								t1.Add(t2, t2)
								t2.Add(t1, t2)
								y3.Sub(y3, t2)
								y3.Sub(y3, t0)
								t1.Add(y3, y3)
								y3.Add(t1, y3)
								t1.Add(t0, t0)
								t0.Add(t1, t0)
								t0.Sub(t0, t2)
								t1.Mul(t4, y3)
								t2.Mul(t0, y3)
								y3.Mul(x3, z3)
								y3.Add(y3, t2)
								x3.Mul(t3, x3)
								x3.Sub(x3, t1)
								z3.Mul(t4, z3)
								t1.Mul(t3, t0)
								z3.Add(z3, t1)

								q.x.Set(x3)
								q.y.Set(y3)
								q.z.Set(z3)
								return q
//line /usr/local/go/src/crypto/internal/nistec/p384.go:266
	// _ = "end of CoverTab[2633]"
}

// Double sets q = p + p, and returns q. The points may overlap.
func (q *P384Point) Double(p *P384Point) *P384Point {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:270
	_go_fuzz_dep_.CoverTab[2634]++

//line /usr/local/go/src/crypto/internal/nistec/p384.go:274
	t0 := new(fiat.P384Element).Square(p.x)
								t1 := new(fiat.P384Element).Square(p.y)
								t2 := new(fiat.P384Element).Square(p.z)
								t3 := new(fiat.P384Element).Mul(p.x, p.y)
								t3.Add(t3, t3)
								z3 := new(fiat.P384Element).Mul(p.x, p.z)
								z3.Add(z3, z3)
								y3 := new(fiat.P384Element).Mul(p384B(), t2)
								y3.Sub(y3, z3)
								x3 := new(fiat.P384Element).Add(y3, y3)
								y3.Add(x3, y3)
								x3.Sub(t1, y3)
								y3.Add(t1, y3)
								y3.Mul(x3, y3)
								x3.Mul(x3, t3)
								t3.Add(t2, t2)
								t2.Add(t2, t3)
								z3.Mul(p384B(), z3)
								z3.Sub(z3, t2)
								z3.Sub(z3, t0)
								t3.Add(z3, z3)
								z3.Add(z3, t3)
								t3.Add(t0, t0)
								t0.Add(t3, t0)
								t0.Sub(t0, t2)
								t0.Mul(t0, z3)
								y3.Add(y3, t0)
								t0.Mul(p.y, p.z)
								t0.Add(t0, t0)
								z3.Mul(t0, z3)
								x3.Sub(x3, z3)
								z3.Mul(t0, t1)
								z3.Add(z3, z3)
								z3.Add(z3, z3)

								q.x.Set(x3)
								q.y.Set(y3)
								q.z.Set(z3)
								return q
//line /usr/local/go/src/crypto/internal/nistec/p384.go:312
	// _ = "end of CoverTab[2634]"
}

// Select sets q to p1 if cond == 1, and to p2 if cond == 0.
func (q *P384Point) Select(p1, p2 *P384Point, cond int) *P384Point {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:316
	_go_fuzz_dep_.CoverTab[2635]++
								q.x.Select(p1.x, p2.x, cond)
								q.y.Select(p1.y, p2.y, cond)
								q.z.Select(p1.z, p2.z, cond)
								return q
//line /usr/local/go/src/crypto/internal/nistec/p384.go:320
	// _ = "end of CoverTab[2635]"
}

// A p384Table holds the first 15 multiples of a point at offset -1, so [1]P
//line /usr/local/go/src/crypto/internal/nistec/p384.go:323
// is at table[0], [15]P is at table[14], and [0]P is implicitly the identity
//line /usr/local/go/src/crypto/internal/nistec/p384.go:323
// point.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:326
type p384Table [15]*P384Point

// Select selects the n-th multiple of the table base point into p. It works in
//line /usr/local/go/src/crypto/internal/nistec/p384.go:328
// constant time by iterating over every entry of the table. n must be in [0, 15].
//line /usr/local/go/src/crypto/internal/nistec/p384.go:330
func (table *p384Table) Select(p *P384Point, n uint8) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:330
	_go_fuzz_dep_.CoverTab[2636]++
								if n >= 16 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:331
		_go_fuzz_dep_.CoverTab[2638]++
									panic("nistec: internal error: p384Table called with out-of-bounds value")
//line /usr/local/go/src/crypto/internal/nistec/p384.go:332
		// _ = "end of CoverTab[2638]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:333
		_go_fuzz_dep_.CoverTab[2639]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:333
		// _ = "end of CoverTab[2639]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:333
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:333
	// _ = "end of CoverTab[2636]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:333
	_go_fuzz_dep_.CoverTab[2637]++
								p.Set(NewP384Point())
								for i := uint8(1); i < 16; i++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:335
		_go_fuzz_dep_.CoverTab[2640]++
									cond := subtle.ConstantTimeByteEq(i, n)
									p.Select(table[i-1], p, cond)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:337
		// _ = "end of CoverTab[2640]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:338
	// _ = "end of CoverTab[2637]"
}

// ScalarMult sets p = scalar * q, and returns p.
func (p *P384Point) ScalarMult(q *P384Point, scalar []byte) (*P384Point, error) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:342
	_go_fuzz_dep_.CoverTab[2641]++
	// Compute a p384Table for the base point q. The explicit NewP384Point
	// calls get inlined, letting the allocations live on the stack.
	var table = p384Table{NewP384Point(), NewP384Point(), NewP384Point(),
		NewP384Point(), NewP384Point(), NewP384Point(), NewP384Point(),
		NewP384Point(), NewP384Point(), NewP384Point(), NewP384Point(),
		NewP384Point(), NewP384Point(), NewP384Point(), NewP384Point()}
	table[0].Set(q)
	for i := 1; i < 15; i += 2 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:350
		_go_fuzz_dep_.CoverTab[2644]++
									table[i].Double(table[i/2])
									table[i+1].Add(table[i], q)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:352
		// _ = "end of CoverTab[2644]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:353
	// _ = "end of CoverTab[2641]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:353
	_go_fuzz_dep_.CoverTab[2642]++

//line /usr/local/go/src/crypto/internal/nistec/p384.go:357
	t := NewP384Point()
	p.Set(NewP384Point())
	for i, byte := range scalar {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:359
		_go_fuzz_dep_.CoverTab[2645]++

//line /usr/local/go/src/crypto/internal/nistec/p384.go:362
		if i != 0 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:362
			_go_fuzz_dep_.CoverTab[2647]++
										p.Double(p)
										p.Double(p)
										p.Double(p)
										p.Double(p)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:366
			// _ = "end of CoverTab[2647]"
		} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:367
			_go_fuzz_dep_.CoverTab[2648]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:367
			// _ = "end of CoverTab[2648]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:367
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:367
		// _ = "end of CoverTab[2645]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:367
		_go_fuzz_dep_.CoverTab[2646]++

									windowValue := byte >> 4
									table.Select(t, windowValue)
									p.Add(p, t)

									p.Double(p)
									p.Double(p)
									p.Double(p)
									p.Double(p)

									windowValue = byte & 0b1111
									table.Select(t, windowValue)
									p.Add(p, t)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:380
		// _ = "end of CoverTab[2646]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:381
	// _ = "end of CoverTab[2642]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:381
	_go_fuzz_dep_.CoverTab[2643]++

								return p, nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:383
	// _ = "end of CoverTab[2643]"
}

var p384GeneratorTable *[p384ElementLength * 2]p384Table
var p384GeneratorTableOnce sync.Once

// generatorTable returns a sequence of p384Tables. The first table contains
//line /usr/local/go/src/crypto/internal/nistec/p384.go:389
// multiples of G. Each successive table is the previous table doubled four
//line /usr/local/go/src/crypto/internal/nistec/p384.go:389
// times.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:392
func (p *P384Point) generatorTable() *[p384ElementLength * 2]p384Table {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:392
	_go_fuzz_dep_.CoverTab[2649]++
								p384GeneratorTableOnce.Do(func() {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:393
		_go_fuzz_dep_.CoverTab[2651]++
									p384GeneratorTable = new([p384ElementLength * 2]p384Table)
									base := NewP384Point().SetGenerator()
									for i := 0; i < p384ElementLength*2; i++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:396
			_go_fuzz_dep_.CoverTab[2652]++
										p384GeneratorTable[i][0] = NewP384Point().Set(base)
										for j := 1; j < 15; j++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:398
				_go_fuzz_dep_.CoverTab[2654]++
											p384GeneratorTable[i][j] = NewP384Point().Add(p384GeneratorTable[i][j-1], base)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:399
				// _ = "end of CoverTab[2654]"
			}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:400
			// _ = "end of CoverTab[2652]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:400
			_go_fuzz_dep_.CoverTab[2653]++
										base.Double(base)
										base.Double(base)
										base.Double(base)
										base.Double(base)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:404
			// _ = "end of CoverTab[2653]"
		}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:405
		// _ = "end of CoverTab[2651]"
	})
//line /usr/local/go/src/crypto/internal/nistec/p384.go:406
	// _ = "end of CoverTab[2649]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:406
	_go_fuzz_dep_.CoverTab[2650]++
								return p384GeneratorTable
//line /usr/local/go/src/crypto/internal/nistec/p384.go:407
	// _ = "end of CoverTab[2650]"
}

// ScalarBaseMult sets p = scalar * B, where B is the canonical generator, and
//line /usr/local/go/src/crypto/internal/nistec/p384.go:410
// returns p.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:412
func (p *P384Point) ScalarBaseMult(scalar []byte) (*P384Point, error) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:412
	_go_fuzz_dep_.CoverTab[2655]++
								if len(scalar) != p384ElementLength {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:413
		_go_fuzz_dep_.CoverTab[2658]++
									return nil, errors.New("invalid scalar length")
//line /usr/local/go/src/crypto/internal/nistec/p384.go:414
		// _ = "end of CoverTab[2658]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:415
		_go_fuzz_dep_.CoverTab[2659]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:415
		// _ = "end of CoverTab[2659]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:415
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:415
	// _ = "end of CoverTab[2655]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:415
	_go_fuzz_dep_.CoverTab[2656]++
								tables := p.generatorTable()

//line /usr/local/go/src/crypto/internal/nistec/p384.go:424
	t := NewP384Point()
	p.Set(NewP384Point())
	tableIndex := len(tables) - 1
	for _, byte := range scalar {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:427
		_go_fuzz_dep_.CoverTab[2660]++
									windowValue := byte >> 4
									tables[tableIndex].Select(t, windowValue)
									p.Add(p, t)
									tableIndex--

									windowValue = byte & 0b1111
									tables[tableIndex].Select(t, windowValue)
									p.Add(p, t)
									tableIndex--
//line /usr/local/go/src/crypto/internal/nistec/p384.go:436
		// _ = "end of CoverTab[2660]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:437
	// _ = "end of CoverTab[2656]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:437
	_go_fuzz_dep_.CoverTab[2657]++

								return p, nil
//line /usr/local/go/src/crypto/internal/nistec/p384.go:439
	// _ = "end of CoverTab[2657]"
}

// p384Sqrt sets e to a square root of x. If x is not a square, p384Sqrt returns
//line /usr/local/go/src/crypto/internal/nistec/p384.go:442
// false and e is unchanged. e and x can overlap.
//line /usr/local/go/src/crypto/internal/nistec/p384.go:444
func p384Sqrt(e, x *fiat.P384Element) (isSquare bool) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:444
	_go_fuzz_dep_.CoverTab[2661]++
								candidate := new(fiat.P384Element)
								p384SqrtCandidate(candidate, x)
								square := new(fiat.P384Element).Square(candidate)
								if square.Equal(x) != 1 {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:448
		_go_fuzz_dep_.CoverTab[2663]++
									return false
//line /usr/local/go/src/crypto/internal/nistec/p384.go:449
		// _ = "end of CoverTab[2663]"
	} else {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:450
		_go_fuzz_dep_.CoverTab[2664]++
//line /usr/local/go/src/crypto/internal/nistec/p384.go:450
		// _ = "end of CoverTab[2664]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:450
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:450
	// _ = "end of CoverTab[2661]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:450
	_go_fuzz_dep_.CoverTab[2662]++
								e.Set(candidate)
								return true
//line /usr/local/go/src/crypto/internal/nistec/p384.go:452
	// _ = "end of CoverTab[2662]"
}

// p384SqrtCandidate sets z to a square root candidate for x. z and x must not overlap.
func p384SqrtCandidate(z, x *fiat.P384Element) {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:456
	_go_fuzz_dep_.CoverTab[2665]++
	// Since p = 3 mod 4, exponentiation by (p + 1) / 4 yields a square root candidate.
	//
	// The sequence of 14 multiplications and 381 squarings is derived from the
	// following addition chain generated with github.com/mmcloughlin/addchain v0.4.0.
	//
	//	_10      = 2*1
	//	_11      = 1 + _10
	//	_110     = 2*_11
	//	_111     = 1 + _110
	//	_111000  = _111 << 3
	//	_111111  = _111 + _111000
	//	_1111110 = 2*_111111
	//	_1111111 = 1 + _1111110
	//	x12      = _1111110 << 5 + _111111
	//	x24      = x12 << 12 + x12
	//	x31      = x24 << 7 + _1111111
	//	x32      = 2*x31 + 1
	//	x63      = x32 << 31 + x31
	//	x126     = x63 << 63 + x63
	//	x252     = x126 << 126 + x126
	//	x255     = x252 << 3 + _111
	//	return     ((x255 << 33 + x32) << 64 + 1) << 30
	//
	var t0 = new(fiat.P384Element)
	var t1 = new(fiat.P384Element)
	var t2 = new(fiat.P384Element)

	z.Square(x)
	z.Mul(x, z)
	z.Square(z)
	t0.Mul(x, z)
	z.Square(t0)
	for s := 1; s < 3; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:489
		_go_fuzz_dep_.CoverTab[2676]++
									z.Square(z)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:490
		// _ = "end of CoverTab[2676]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:491
	// _ = "end of CoverTab[2665]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:491
	_go_fuzz_dep_.CoverTab[2666]++
								t1.Mul(t0, z)
								t2.Square(t1)
								z.Mul(x, t2)
								for s := 0; s < 5; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:495
		_go_fuzz_dep_.CoverTab[2677]++
									t2.Square(t2)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:496
		// _ = "end of CoverTab[2677]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:497
	// _ = "end of CoverTab[2666]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:497
	_go_fuzz_dep_.CoverTab[2667]++
								t1.Mul(t1, t2)
								t2.Square(t1)
								for s := 1; s < 12; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:500
		_go_fuzz_dep_.CoverTab[2678]++
									t2.Square(t2)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:501
		// _ = "end of CoverTab[2678]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:502
	// _ = "end of CoverTab[2667]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:502
	_go_fuzz_dep_.CoverTab[2668]++
								t1.Mul(t1, t2)
								for s := 0; s < 7; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:504
		_go_fuzz_dep_.CoverTab[2679]++
									t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:505
		// _ = "end of CoverTab[2679]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:506
	// _ = "end of CoverTab[2668]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:506
	_go_fuzz_dep_.CoverTab[2669]++
								t1.Mul(z, t1)
								z.Square(t1)
								z.Mul(x, z)
								t2.Square(z)
								for s := 1; s < 31; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:511
		_go_fuzz_dep_.CoverTab[2680]++
									t2.Square(t2)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:512
		// _ = "end of CoverTab[2680]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:513
	// _ = "end of CoverTab[2669]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:513
	_go_fuzz_dep_.CoverTab[2670]++
								t1.Mul(t1, t2)
								t2.Square(t1)
								for s := 1; s < 63; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:516
		_go_fuzz_dep_.CoverTab[2681]++
									t2.Square(t2)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:517
		// _ = "end of CoverTab[2681]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:518
	// _ = "end of CoverTab[2670]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:518
	_go_fuzz_dep_.CoverTab[2671]++
								t1.Mul(t1, t2)
								t2.Square(t1)
								for s := 1; s < 126; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:521
		_go_fuzz_dep_.CoverTab[2682]++
									t2.Square(t2)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:522
		// _ = "end of CoverTab[2682]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:523
	// _ = "end of CoverTab[2671]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:523
	_go_fuzz_dep_.CoverTab[2672]++
								t1.Mul(t1, t2)
								for s := 0; s < 3; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:525
		_go_fuzz_dep_.CoverTab[2683]++
									t1.Square(t1)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:526
		// _ = "end of CoverTab[2683]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:527
	// _ = "end of CoverTab[2672]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:527
	_go_fuzz_dep_.CoverTab[2673]++
								t0.Mul(t0, t1)
								for s := 0; s < 33; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:529
		_go_fuzz_dep_.CoverTab[2684]++
									t0.Square(t0)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:530
		// _ = "end of CoverTab[2684]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:531
	// _ = "end of CoverTab[2673]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:531
	_go_fuzz_dep_.CoverTab[2674]++
								z.Mul(z, t0)
								for s := 0; s < 64; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:533
		_go_fuzz_dep_.CoverTab[2685]++
									z.Square(z)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:534
		// _ = "end of CoverTab[2685]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:535
	// _ = "end of CoverTab[2674]"
//line /usr/local/go/src/crypto/internal/nistec/p384.go:535
	_go_fuzz_dep_.CoverTab[2675]++
								z.Mul(x, z)
								for s := 0; s < 30; s++ {
//line /usr/local/go/src/crypto/internal/nistec/p384.go:537
		_go_fuzz_dep_.CoverTab[2686]++
									z.Square(z)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:538
		// _ = "end of CoverTab[2686]"
	}
//line /usr/local/go/src/crypto/internal/nistec/p384.go:539
	// _ = "end of CoverTab[2675]"
}

//line /usr/local/go/src/crypto/internal/nistec/p384.go:540
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/nistec/p384.go:540
var _ = _go_fuzz_dep_.CoverTab