// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/asn1/marshal.go:5
package asn1

//line /usr/local/go/src/encoding/asn1/marshal.go:5
import (
//line /usr/local/go/src/encoding/asn1/marshal.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/asn1/marshal.go:5
)
//line /usr/local/go/src/encoding/asn1/marshal.go:5
import (
//line /usr/local/go/src/encoding/asn1/marshal.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/asn1/marshal.go:5
)

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"time"
	"unicode/utf8"
)

var (
	byte00Encoder	encoder	= byteEncoder(0x00)
	byteFFEncoder	encoder	= byteEncoder(0xff)
)

// encoder represents an ASN.1 element that is waiting to be marshaled.
type encoder interface {
	// Len returns the number of bytes needed to marshal this element.
	Len() int
	// Encode encodes this element by writing Len() bytes to dst.
	Encode(dst []byte)
}

type byteEncoder byte

func (c byteEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:33
	_go_fuzz_dep_.CoverTab[8003]++
							return 1
//line /usr/local/go/src/encoding/asn1/marshal.go:34
	// _ = "end of CoverTab[8003]"
}

func (c byteEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:37
	_go_fuzz_dep_.CoverTab[8004]++
							dst[0] = byte(c)
//line /usr/local/go/src/encoding/asn1/marshal.go:38
	// _ = "end of CoverTab[8004]"
}

type bytesEncoder []byte

func (b bytesEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:43
	_go_fuzz_dep_.CoverTab[8005]++
							return len(b)
//line /usr/local/go/src/encoding/asn1/marshal.go:44
	// _ = "end of CoverTab[8005]"
}

func (b bytesEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:47
	_go_fuzz_dep_.CoverTab[8006]++
							if copy(dst, b) != len(b) {
//line /usr/local/go/src/encoding/asn1/marshal.go:48
		_go_fuzz_dep_.CoverTab[8007]++
								panic("internal error")
//line /usr/local/go/src/encoding/asn1/marshal.go:49
		// _ = "end of CoverTab[8007]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:50
		_go_fuzz_dep_.CoverTab[8008]++
//line /usr/local/go/src/encoding/asn1/marshal.go:50
		// _ = "end of CoverTab[8008]"
//line /usr/local/go/src/encoding/asn1/marshal.go:50
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:50
	// _ = "end of CoverTab[8006]"
}

type stringEncoder string

func (s stringEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:55
	_go_fuzz_dep_.CoverTab[8009]++
							return len(s)
//line /usr/local/go/src/encoding/asn1/marshal.go:56
	// _ = "end of CoverTab[8009]"
}

func (s stringEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:59
	_go_fuzz_dep_.CoverTab[8010]++
							if copy(dst, s) != len(s) {
//line /usr/local/go/src/encoding/asn1/marshal.go:60
		_go_fuzz_dep_.CoverTab[8011]++
								panic("internal error")
//line /usr/local/go/src/encoding/asn1/marshal.go:61
		// _ = "end of CoverTab[8011]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:62
		_go_fuzz_dep_.CoverTab[8012]++
//line /usr/local/go/src/encoding/asn1/marshal.go:62
		// _ = "end of CoverTab[8012]"
//line /usr/local/go/src/encoding/asn1/marshal.go:62
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:62
	// _ = "end of CoverTab[8010]"
}

type multiEncoder []encoder

func (m multiEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:67
	_go_fuzz_dep_.CoverTab[8013]++
							var size int
							for _, e := range m {
//line /usr/local/go/src/encoding/asn1/marshal.go:69
		_go_fuzz_dep_.CoverTab[8015]++
								size += e.Len()
//line /usr/local/go/src/encoding/asn1/marshal.go:70
		// _ = "end of CoverTab[8015]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:71
	// _ = "end of CoverTab[8013]"
//line /usr/local/go/src/encoding/asn1/marshal.go:71
	_go_fuzz_dep_.CoverTab[8014]++
							return size
//line /usr/local/go/src/encoding/asn1/marshal.go:72
	// _ = "end of CoverTab[8014]"
}

func (m multiEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:75
	_go_fuzz_dep_.CoverTab[8016]++
							var off int
							for _, e := range m {
//line /usr/local/go/src/encoding/asn1/marshal.go:77
		_go_fuzz_dep_.CoverTab[8017]++
								e.Encode(dst[off:])
								off += e.Len()
//line /usr/local/go/src/encoding/asn1/marshal.go:79
		// _ = "end of CoverTab[8017]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:80
	// _ = "end of CoverTab[8016]"
}

type setEncoder []encoder

func (s setEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:85
	_go_fuzz_dep_.CoverTab[8018]++
							var size int
							for _, e := range s {
//line /usr/local/go/src/encoding/asn1/marshal.go:87
		_go_fuzz_dep_.CoverTab[8020]++
								size += e.Len()
//line /usr/local/go/src/encoding/asn1/marshal.go:88
		// _ = "end of CoverTab[8020]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:89
	// _ = "end of CoverTab[8018]"
//line /usr/local/go/src/encoding/asn1/marshal.go:89
	_go_fuzz_dep_.CoverTab[8019]++
							return size
//line /usr/local/go/src/encoding/asn1/marshal.go:90
	// _ = "end of CoverTab[8019]"
}

func (s setEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:93
	_go_fuzz_dep_.CoverTab[8021]++

//line /usr/local/go/src/encoding/asn1/marshal.go:102
	l := make([][]byte, len(s))
	for i, e := range s {
//line /usr/local/go/src/encoding/asn1/marshal.go:103
		_go_fuzz_dep_.CoverTab[8024]++
								l[i] = make([]byte, e.Len())
								e.Encode(l[i])
//line /usr/local/go/src/encoding/asn1/marshal.go:105
		// _ = "end of CoverTab[8024]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:106
	// _ = "end of CoverTab[8021]"
//line /usr/local/go/src/encoding/asn1/marshal.go:106
	_go_fuzz_dep_.CoverTab[8022]++

							sort.Slice(l, func(i, j int) bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:108
		_go_fuzz_dep_.CoverTab[8025]++

//line /usr/local/go/src/encoding/asn1/marshal.go:115
		return bytes.Compare(l[i], l[j]) < 0
//line /usr/local/go/src/encoding/asn1/marshal.go:115
		// _ = "end of CoverTab[8025]"
	})
//line /usr/local/go/src/encoding/asn1/marshal.go:116
	// _ = "end of CoverTab[8022]"
//line /usr/local/go/src/encoding/asn1/marshal.go:116
	_go_fuzz_dep_.CoverTab[8023]++

							var off int
							for _, b := range l {
//line /usr/local/go/src/encoding/asn1/marshal.go:119
		_go_fuzz_dep_.CoverTab[8026]++
								copy(dst[off:], b)
								off += len(b)
//line /usr/local/go/src/encoding/asn1/marshal.go:121
		// _ = "end of CoverTab[8026]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:122
	// _ = "end of CoverTab[8023]"
}

type taggedEncoder struct {
	// scratch contains temporary space for encoding the tag and length of
	// an element in order to avoid extra allocations.
	scratch	[8]byte
	tag	encoder
	body	encoder
}

func (t *taggedEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:133
	_go_fuzz_dep_.CoverTab[8027]++
							return t.tag.Len() + t.body.Len()
//line /usr/local/go/src/encoding/asn1/marshal.go:134
	// _ = "end of CoverTab[8027]"
}

func (t *taggedEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:137
	_go_fuzz_dep_.CoverTab[8028]++
							t.tag.Encode(dst)
							t.body.Encode(dst[t.tag.Len():])
//line /usr/local/go/src/encoding/asn1/marshal.go:139
	// _ = "end of CoverTab[8028]"
}

type int64Encoder int64

func (i int64Encoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:144
	_go_fuzz_dep_.CoverTab[8029]++
							n := 1

							for i > 127 {
//line /usr/local/go/src/encoding/asn1/marshal.go:147
		_go_fuzz_dep_.CoverTab[8032]++
								n++
								i >>= 8
//line /usr/local/go/src/encoding/asn1/marshal.go:149
		// _ = "end of CoverTab[8032]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:150
	// _ = "end of CoverTab[8029]"
//line /usr/local/go/src/encoding/asn1/marshal.go:150
	_go_fuzz_dep_.CoverTab[8030]++

							for i < -128 {
//line /usr/local/go/src/encoding/asn1/marshal.go:152
		_go_fuzz_dep_.CoverTab[8033]++
								n++
								i >>= 8
//line /usr/local/go/src/encoding/asn1/marshal.go:154
		// _ = "end of CoverTab[8033]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:155
	// _ = "end of CoverTab[8030]"
//line /usr/local/go/src/encoding/asn1/marshal.go:155
	_go_fuzz_dep_.CoverTab[8031]++

							return n
//line /usr/local/go/src/encoding/asn1/marshal.go:157
	// _ = "end of CoverTab[8031]"
}

func (i int64Encoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:160
	_go_fuzz_dep_.CoverTab[8034]++
							n := i.Len()

							for j := 0; j < n; j++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:163
		_go_fuzz_dep_.CoverTab[8035]++
								dst[j] = byte(i >> uint((n-1-j)*8))
//line /usr/local/go/src/encoding/asn1/marshal.go:164
		// _ = "end of CoverTab[8035]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:165
	// _ = "end of CoverTab[8034]"
}

func base128IntLength(n int64) int {
//line /usr/local/go/src/encoding/asn1/marshal.go:168
	_go_fuzz_dep_.CoverTab[8036]++
							if n == 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:169
		_go_fuzz_dep_.CoverTab[8039]++
								return 1
//line /usr/local/go/src/encoding/asn1/marshal.go:170
		// _ = "end of CoverTab[8039]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:171
		_go_fuzz_dep_.CoverTab[8040]++
//line /usr/local/go/src/encoding/asn1/marshal.go:171
		// _ = "end of CoverTab[8040]"
//line /usr/local/go/src/encoding/asn1/marshal.go:171
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:171
	// _ = "end of CoverTab[8036]"
//line /usr/local/go/src/encoding/asn1/marshal.go:171
	_go_fuzz_dep_.CoverTab[8037]++

							l := 0
							for i := n; i > 0; i >>= 7 {
//line /usr/local/go/src/encoding/asn1/marshal.go:174
		_go_fuzz_dep_.CoverTab[8041]++
								l++
//line /usr/local/go/src/encoding/asn1/marshal.go:175
		// _ = "end of CoverTab[8041]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:176
	// _ = "end of CoverTab[8037]"
//line /usr/local/go/src/encoding/asn1/marshal.go:176
	_go_fuzz_dep_.CoverTab[8038]++

							return l
//line /usr/local/go/src/encoding/asn1/marshal.go:178
	// _ = "end of CoverTab[8038]"
}

func appendBase128Int(dst []byte, n int64) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:181
	_go_fuzz_dep_.CoverTab[8042]++
							l := base128IntLength(n)

							for i := l - 1; i >= 0; i-- {
//line /usr/local/go/src/encoding/asn1/marshal.go:184
		_go_fuzz_dep_.CoverTab[8044]++
								o := byte(n >> uint(i*7))
								o &= 0x7f
								if i != 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:187
			_go_fuzz_dep_.CoverTab[8046]++
									o |= 0x80
//line /usr/local/go/src/encoding/asn1/marshal.go:188
			// _ = "end of CoverTab[8046]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:189
			_go_fuzz_dep_.CoverTab[8047]++
//line /usr/local/go/src/encoding/asn1/marshal.go:189
			// _ = "end of CoverTab[8047]"
//line /usr/local/go/src/encoding/asn1/marshal.go:189
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:189
		// _ = "end of CoverTab[8044]"
//line /usr/local/go/src/encoding/asn1/marshal.go:189
		_go_fuzz_dep_.CoverTab[8045]++

								dst = append(dst, o)
//line /usr/local/go/src/encoding/asn1/marshal.go:191
		// _ = "end of CoverTab[8045]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:192
	// _ = "end of CoverTab[8042]"
//line /usr/local/go/src/encoding/asn1/marshal.go:192
	_go_fuzz_dep_.CoverTab[8043]++

							return dst
//line /usr/local/go/src/encoding/asn1/marshal.go:194
	// _ = "end of CoverTab[8043]"
}

func makeBigInt(n *big.Int) (encoder, error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:197
	_go_fuzz_dep_.CoverTab[8048]++
							if n == nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:198
		_go_fuzz_dep_.CoverTab[8050]++
								return nil, StructuralError{"empty integer"}
//line /usr/local/go/src/encoding/asn1/marshal.go:199
		// _ = "end of CoverTab[8050]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:200
		_go_fuzz_dep_.CoverTab[8051]++
//line /usr/local/go/src/encoding/asn1/marshal.go:200
		// _ = "end of CoverTab[8051]"
//line /usr/local/go/src/encoding/asn1/marshal.go:200
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:200
	// _ = "end of CoverTab[8048]"
//line /usr/local/go/src/encoding/asn1/marshal.go:200
	_go_fuzz_dep_.CoverTab[8049]++

							if n.Sign() < 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:202
		_go_fuzz_dep_.CoverTab[8052]++

//line /usr/local/go/src/encoding/asn1/marshal.go:207
		nMinus1 := new(big.Int).Neg(n)
		nMinus1.Sub(nMinus1, bigOne)
		bytes := nMinus1.Bytes()
		for i := range bytes {
//line /usr/local/go/src/encoding/asn1/marshal.go:210
			_go_fuzz_dep_.CoverTab[8055]++
									bytes[i] ^= 0xff
//line /usr/local/go/src/encoding/asn1/marshal.go:211
			// _ = "end of CoverTab[8055]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:212
		// _ = "end of CoverTab[8052]"
//line /usr/local/go/src/encoding/asn1/marshal.go:212
		_go_fuzz_dep_.CoverTab[8053]++
								if len(bytes) == 0 || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:213
			_go_fuzz_dep_.CoverTab[8056]++
//line /usr/local/go/src/encoding/asn1/marshal.go:213
			return bytes[0]&0x80 == 0
//line /usr/local/go/src/encoding/asn1/marshal.go:213
			// _ = "end of CoverTab[8056]"
//line /usr/local/go/src/encoding/asn1/marshal.go:213
		}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:213
			_go_fuzz_dep_.CoverTab[8057]++
									return multiEncoder([]encoder{byteFFEncoder, bytesEncoder(bytes)}), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:214
			// _ = "end of CoverTab[8057]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:215
			_go_fuzz_dep_.CoverTab[8058]++
//line /usr/local/go/src/encoding/asn1/marshal.go:215
			// _ = "end of CoverTab[8058]"
//line /usr/local/go/src/encoding/asn1/marshal.go:215
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:215
		// _ = "end of CoverTab[8053]"
//line /usr/local/go/src/encoding/asn1/marshal.go:215
		_go_fuzz_dep_.CoverTab[8054]++
								return bytesEncoder(bytes), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:216
		// _ = "end of CoverTab[8054]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:217
		_go_fuzz_dep_.CoverTab[8059]++
//line /usr/local/go/src/encoding/asn1/marshal.go:217
		if n.Sign() == 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:217
			_go_fuzz_dep_.CoverTab[8060]++

									return byte00Encoder, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:219
			// _ = "end of CoverTab[8060]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:220
			_go_fuzz_dep_.CoverTab[8061]++
									bytes := n.Bytes()
									if len(bytes) > 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:222
				_go_fuzz_dep_.CoverTab[8063]++
//line /usr/local/go/src/encoding/asn1/marshal.go:222
				return bytes[0]&0x80 != 0
//line /usr/local/go/src/encoding/asn1/marshal.go:222
				// _ = "end of CoverTab[8063]"
//line /usr/local/go/src/encoding/asn1/marshal.go:222
			}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:222
				_go_fuzz_dep_.CoverTab[8064]++

//line /usr/local/go/src/encoding/asn1/marshal.go:225
				return multiEncoder([]encoder{byte00Encoder, bytesEncoder(bytes)}), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:225
				// _ = "end of CoverTab[8064]"
			} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:226
				_go_fuzz_dep_.CoverTab[8065]++
//line /usr/local/go/src/encoding/asn1/marshal.go:226
				// _ = "end of CoverTab[8065]"
//line /usr/local/go/src/encoding/asn1/marshal.go:226
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:226
			// _ = "end of CoverTab[8061]"
//line /usr/local/go/src/encoding/asn1/marshal.go:226
			_go_fuzz_dep_.CoverTab[8062]++
									return bytesEncoder(bytes), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:227
			// _ = "end of CoverTab[8062]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:228
		// _ = "end of CoverTab[8059]"
//line /usr/local/go/src/encoding/asn1/marshal.go:228
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:228
	// _ = "end of CoverTab[8049]"
}

func appendLength(dst []byte, i int) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:231
	_go_fuzz_dep_.CoverTab[8066]++
							n := lengthLength(i)

							for ; n > 0; n-- {
//line /usr/local/go/src/encoding/asn1/marshal.go:234
		_go_fuzz_dep_.CoverTab[8068]++
								dst = append(dst, byte(i>>uint((n-1)*8)))
//line /usr/local/go/src/encoding/asn1/marshal.go:235
		// _ = "end of CoverTab[8068]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:236
	// _ = "end of CoverTab[8066]"
//line /usr/local/go/src/encoding/asn1/marshal.go:236
	_go_fuzz_dep_.CoverTab[8067]++

							return dst
//line /usr/local/go/src/encoding/asn1/marshal.go:238
	// _ = "end of CoverTab[8067]"
}

func lengthLength(i int) (numBytes int) {
//line /usr/local/go/src/encoding/asn1/marshal.go:241
	_go_fuzz_dep_.CoverTab[8069]++
							numBytes = 1
							for i > 255 {
//line /usr/local/go/src/encoding/asn1/marshal.go:243
		_go_fuzz_dep_.CoverTab[8071]++
								numBytes++
								i >>= 8
//line /usr/local/go/src/encoding/asn1/marshal.go:245
		// _ = "end of CoverTab[8071]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:246
	// _ = "end of CoverTab[8069]"
//line /usr/local/go/src/encoding/asn1/marshal.go:246
	_go_fuzz_dep_.CoverTab[8070]++
							return
//line /usr/local/go/src/encoding/asn1/marshal.go:247
	// _ = "end of CoverTab[8070]"
}

func appendTagAndLength(dst []byte, t tagAndLength) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:250
	_go_fuzz_dep_.CoverTab[8072]++
							b := uint8(t.class) << 6
							if t.isCompound {
//line /usr/local/go/src/encoding/asn1/marshal.go:252
		_go_fuzz_dep_.CoverTab[8076]++
								b |= 0x20
//line /usr/local/go/src/encoding/asn1/marshal.go:253
		// _ = "end of CoverTab[8076]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:254
		_go_fuzz_dep_.CoverTab[8077]++
//line /usr/local/go/src/encoding/asn1/marshal.go:254
		// _ = "end of CoverTab[8077]"
//line /usr/local/go/src/encoding/asn1/marshal.go:254
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:254
	// _ = "end of CoverTab[8072]"
//line /usr/local/go/src/encoding/asn1/marshal.go:254
	_go_fuzz_dep_.CoverTab[8073]++
							if t.tag >= 31 {
//line /usr/local/go/src/encoding/asn1/marshal.go:255
		_go_fuzz_dep_.CoverTab[8078]++
								b |= 0x1f
								dst = append(dst, b)
								dst = appendBase128Int(dst, int64(t.tag))
//line /usr/local/go/src/encoding/asn1/marshal.go:258
		// _ = "end of CoverTab[8078]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:259
		_go_fuzz_dep_.CoverTab[8079]++
								b |= uint8(t.tag)
								dst = append(dst, b)
//line /usr/local/go/src/encoding/asn1/marshal.go:261
		// _ = "end of CoverTab[8079]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:262
	// _ = "end of CoverTab[8073]"
//line /usr/local/go/src/encoding/asn1/marshal.go:262
	_go_fuzz_dep_.CoverTab[8074]++

							if t.length >= 128 {
//line /usr/local/go/src/encoding/asn1/marshal.go:264
		_go_fuzz_dep_.CoverTab[8080]++
								l := lengthLength(t.length)
								dst = append(dst, 0x80|byte(l))
								dst = appendLength(dst, t.length)
//line /usr/local/go/src/encoding/asn1/marshal.go:267
		// _ = "end of CoverTab[8080]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:268
		_go_fuzz_dep_.CoverTab[8081]++
								dst = append(dst, byte(t.length))
//line /usr/local/go/src/encoding/asn1/marshal.go:269
		// _ = "end of CoverTab[8081]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:270
	// _ = "end of CoverTab[8074]"
//line /usr/local/go/src/encoding/asn1/marshal.go:270
	_go_fuzz_dep_.CoverTab[8075]++

							return dst
//line /usr/local/go/src/encoding/asn1/marshal.go:272
	// _ = "end of CoverTab[8075]"
}

type bitStringEncoder BitString

func (b bitStringEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:277
	_go_fuzz_dep_.CoverTab[8082]++
							return len(b.Bytes) + 1
//line /usr/local/go/src/encoding/asn1/marshal.go:278
	// _ = "end of CoverTab[8082]"
}

func (b bitStringEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:281
	_go_fuzz_dep_.CoverTab[8083]++
							dst[0] = byte((8 - b.BitLength%8) % 8)
							if copy(dst[1:], b.Bytes) != len(b.Bytes) {
//line /usr/local/go/src/encoding/asn1/marshal.go:283
		_go_fuzz_dep_.CoverTab[8084]++
								panic("internal error")
//line /usr/local/go/src/encoding/asn1/marshal.go:284
		// _ = "end of CoverTab[8084]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:285
		_go_fuzz_dep_.CoverTab[8085]++
//line /usr/local/go/src/encoding/asn1/marshal.go:285
		// _ = "end of CoverTab[8085]"
//line /usr/local/go/src/encoding/asn1/marshal.go:285
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:285
	// _ = "end of CoverTab[8083]"
}

type oidEncoder []int

func (oid oidEncoder) Len() int {
//line /usr/local/go/src/encoding/asn1/marshal.go:290
	_go_fuzz_dep_.CoverTab[8086]++
							l := base128IntLength(int64(oid[0]*40 + oid[1]))
							for i := 2; i < len(oid); i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:292
		_go_fuzz_dep_.CoverTab[8088]++
								l += base128IntLength(int64(oid[i]))
//line /usr/local/go/src/encoding/asn1/marshal.go:293
		// _ = "end of CoverTab[8088]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:294
	// _ = "end of CoverTab[8086]"
//line /usr/local/go/src/encoding/asn1/marshal.go:294
	_go_fuzz_dep_.CoverTab[8087]++
							return l
//line /usr/local/go/src/encoding/asn1/marshal.go:295
	// _ = "end of CoverTab[8087]"
}

func (oid oidEncoder) Encode(dst []byte) {
//line /usr/local/go/src/encoding/asn1/marshal.go:298
	_go_fuzz_dep_.CoverTab[8089]++
							dst = appendBase128Int(dst[:0], int64(oid[0]*40+oid[1]))
							for i := 2; i < len(oid); i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:300
		_go_fuzz_dep_.CoverTab[8090]++
								dst = appendBase128Int(dst, int64(oid[i]))
//line /usr/local/go/src/encoding/asn1/marshal.go:301
		// _ = "end of CoverTab[8090]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:302
	// _ = "end of CoverTab[8089]"
}

func makeObjectIdentifier(oid []int) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:305
	_go_fuzz_dep_.CoverTab[8091]++
							if len(oid) < 2 || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		_go_fuzz_dep_.CoverTab[8093]++
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		return oid[0] > 2
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		// _ = "end of CoverTab[8093]"
//line /usr/local/go/src/encoding/asn1/marshal.go:306
	}() || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		_go_fuzz_dep_.CoverTab[8094]++
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		return (oid[0] < 2 && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:306
			_go_fuzz_dep_.CoverTab[8095]++
//line /usr/local/go/src/encoding/asn1/marshal.go:306
			return oid[1] >= 40
//line /usr/local/go/src/encoding/asn1/marshal.go:306
			// _ = "end of CoverTab[8095]"
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		}())
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		// _ = "end of CoverTab[8094]"
//line /usr/local/go/src/encoding/asn1/marshal.go:306
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:306
		_go_fuzz_dep_.CoverTab[8096]++
								return nil, StructuralError{"invalid object identifier"}
//line /usr/local/go/src/encoding/asn1/marshal.go:307
		// _ = "end of CoverTab[8096]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:308
		_go_fuzz_dep_.CoverTab[8097]++
//line /usr/local/go/src/encoding/asn1/marshal.go:308
		// _ = "end of CoverTab[8097]"
//line /usr/local/go/src/encoding/asn1/marshal.go:308
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:308
	// _ = "end of CoverTab[8091]"
//line /usr/local/go/src/encoding/asn1/marshal.go:308
	_go_fuzz_dep_.CoverTab[8092]++

							return oidEncoder(oid), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:310
	// _ = "end of CoverTab[8092]"
}

func makePrintableString(s string) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:313
	_go_fuzz_dep_.CoverTab[8098]++
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:314
		_go_fuzz_dep_.CoverTab[8100]++

//line /usr/local/go/src/encoding/asn1/marshal.go:321
		if !isPrintable(s[i], allowAsterisk, rejectAmpersand) {
//line /usr/local/go/src/encoding/asn1/marshal.go:321
			_go_fuzz_dep_.CoverTab[8101]++
									return nil, StructuralError{"PrintableString contains invalid character"}
//line /usr/local/go/src/encoding/asn1/marshal.go:322
			// _ = "end of CoverTab[8101]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:323
			_go_fuzz_dep_.CoverTab[8102]++
//line /usr/local/go/src/encoding/asn1/marshal.go:323
			// _ = "end of CoverTab[8102]"
//line /usr/local/go/src/encoding/asn1/marshal.go:323
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:323
		// _ = "end of CoverTab[8100]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:324
	// _ = "end of CoverTab[8098]"
//line /usr/local/go/src/encoding/asn1/marshal.go:324
	_go_fuzz_dep_.CoverTab[8099]++

							return stringEncoder(s), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:326
	// _ = "end of CoverTab[8099]"
}

func makeIA5String(s string) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:329
	_go_fuzz_dep_.CoverTab[8103]++
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:330
		_go_fuzz_dep_.CoverTab[8105]++
								if s[i] > 127 {
//line /usr/local/go/src/encoding/asn1/marshal.go:331
			_go_fuzz_dep_.CoverTab[8106]++
									return nil, StructuralError{"IA5String contains invalid character"}
//line /usr/local/go/src/encoding/asn1/marshal.go:332
			// _ = "end of CoverTab[8106]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:333
			_go_fuzz_dep_.CoverTab[8107]++
//line /usr/local/go/src/encoding/asn1/marshal.go:333
			// _ = "end of CoverTab[8107]"
//line /usr/local/go/src/encoding/asn1/marshal.go:333
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:333
		// _ = "end of CoverTab[8105]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:334
	// _ = "end of CoverTab[8103]"
//line /usr/local/go/src/encoding/asn1/marshal.go:334
	_go_fuzz_dep_.CoverTab[8104]++

							return stringEncoder(s), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:336
	// _ = "end of CoverTab[8104]"
}

func makeNumericString(s string) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:339
	_go_fuzz_dep_.CoverTab[8108]++
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:340
		_go_fuzz_dep_.CoverTab[8110]++
								if !isNumeric(s[i]) {
//line /usr/local/go/src/encoding/asn1/marshal.go:341
			_go_fuzz_dep_.CoverTab[8111]++
									return nil, StructuralError{"NumericString contains invalid character"}
//line /usr/local/go/src/encoding/asn1/marshal.go:342
			// _ = "end of CoverTab[8111]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:343
			_go_fuzz_dep_.CoverTab[8112]++
//line /usr/local/go/src/encoding/asn1/marshal.go:343
			// _ = "end of CoverTab[8112]"
//line /usr/local/go/src/encoding/asn1/marshal.go:343
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:343
		// _ = "end of CoverTab[8110]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:344
	// _ = "end of CoverTab[8108]"
//line /usr/local/go/src/encoding/asn1/marshal.go:344
	_go_fuzz_dep_.CoverTab[8109]++

							return stringEncoder(s), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:346
	// _ = "end of CoverTab[8109]"
}

func makeUTF8String(s string) encoder {
//line /usr/local/go/src/encoding/asn1/marshal.go:349
	_go_fuzz_dep_.CoverTab[8113]++
							return stringEncoder(s)
//line /usr/local/go/src/encoding/asn1/marshal.go:350
	// _ = "end of CoverTab[8113]"
}

func appendTwoDigits(dst []byte, v int) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:353
	_go_fuzz_dep_.CoverTab[8114]++
							return append(dst, byte('0'+(v/10)%10), byte('0'+v%10))
//line /usr/local/go/src/encoding/asn1/marshal.go:354
	// _ = "end of CoverTab[8114]"
}

func appendFourDigits(dst []byte, v int) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:357
	_go_fuzz_dep_.CoverTab[8115]++
							var bytes [4]byte
							for i := range bytes {
//line /usr/local/go/src/encoding/asn1/marshal.go:359
		_go_fuzz_dep_.CoverTab[8117]++
								bytes[3-i] = '0' + byte(v%10)
								v /= 10
//line /usr/local/go/src/encoding/asn1/marshal.go:361
		// _ = "end of CoverTab[8117]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:362
	// _ = "end of CoverTab[8115]"
//line /usr/local/go/src/encoding/asn1/marshal.go:362
	_go_fuzz_dep_.CoverTab[8116]++
							return append(dst, bytes[:]...)
//line /usr/local/go/src/encoding/asn1/marshal.go:363
	// _ = "end of CoverTab[8116]"
}

func outsideUTCRange(t time.Time) bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:366
	_go_fuzz_dep_.CoverTab[8118]++
							year := t.Year()
							return year < 1950 || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:368
		_go_fuzz_dep_.CoverTab[8119]++
//line /usr/local/go/src/encoding/asn1/marshal.go:368
		return year >= 2050
//line /usr/local/go/src/encoding/asn1/marshal.go:368
		// _ = "end of CoverTab[8119]"
//line /usr/local/go/src/encoding/asn1/marshal.go:368
	}()
//line /usr/local/go/src/encoding/asn1/marshal.go:368
	// _ = "end of CoverTab[8118]"
}

func makeUTCTime(t time.Time) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:371
	_go_fuzz_dep_.CoverTab[8120]++
							dst := make([]byte, 0, 18)

							dst, err = appendUTCTime(dst, t)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:375
		_go_fuzz_dep_.CoverTab[8122]++
								return nil, err
//line /usr/local/go/src/encoding/asn1/marshal.go:376
		// _ = "end of CoverTab[8122]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:377
		_go_fuzz_dep_.CoverTab[8123]++
//line /usr/local/go/src/encoding/asn1/marshal.go:377
		// _ = "end of CoverTab[8123]"
//line /usr/local/go/src/encoding/asn1/marshal.go:377
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:377
	// _ = "end of CoverTab[8120]"
//line /usr/local/go/src/encoding/asn1/marshal.go:377
	_go_fuzz_dep_.CoverTab[8121]++

							return bytesEncoder(dst), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:379
	// _ = "end of CoverTab[8121]"
}

func makeGeneralizedTime(t time.Time) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:382
	_go_fuzz_dep_.CoverTab[8124]++
							dst := make([]byte, 0, 20)

							dst, err = appendGeneralizedTime(dst, t)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:386
		_go_fuzz_dep_.CoverTab[8126]++
								return nil, err
//line /usr/local/go/src/encoding/asn1/marshal.go:387
		// _ = "end of CoverTab[8126]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:388
		_go_fuzz_dep_.CoverTab[8127]++
//line /usr/local/go/src/encoding/asn1/marshal.go:388
		// _ = "end of CoverTab[8127]"
//line /usr/local/go/src/encoding/asn1/marshal.go:388
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:388
	// _ = "end of CoverTab[8124]"
//line /usr/local/go/src/encoding/asn1/marshal.go:388
	_go_fuzz_dep_.CoverTab[8125]++

							return bytesEncoder(dst), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:390
	// _ = "end of CoverTab[8125]"
}

func appendUTCTime(dst []byte, t time.Time) (ret []byte, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:393
	_go_fuzz_dep_.CoverTab[8128]++
							year := t.Year()

							switch {
	case 1950 <= year && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:397
		_go_fuzz_dep_.CoverTab[8133]++
//line /usr/local/go/src/encoding/asn1/marshal.go:397
		return year < 2000
//line /usr/local/go/src/encoding/asn1/marshal.go:397
		// _ = "end of CoverTab[8133]"
//line /usr/local/go/src/encoding/asn1/marshal.go:397
	}():
//line /usr/local/go/src/encoding/asn1/marshal.go:397
		_go_fuzz_dep_.CoverTab[8130]++
								dst = appendTwoDigits(dst, year-1900)
//line /usr/local/go/src/encoding/asn1/marshal.go:398
		// _ = "end of CoverTab[8130]"
	case 2000 <= year && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:399
		_go_fuzz_dep_.CoverTab[8134]++
//line /usr/local/go/src/encoding/asn1/marshal.go:399
		return year < 2050
//line /usr/local/go/src/encoding/asn1/marshal.go:399
		// _ = "end of CoverTab[8134]"
//line /usr/local/go/src/encoding/asn1/marshal.go:399
	}():
//line /usr/local/go/src/encoding/asn1/marshal.go:399
		_go_fuzz_dep_.CoverTab[8131]++
								dst = appendTwoDigits(dst, year-2000)
//line /usr/local/go/src/encoding/asn1/marshal.go:400
		// _ = "end of CoverTab[8131]"
	default:
//line /usr/local/go/src/encoding/asn1/marshal.go:401
		_go_fuzz_dep_.CoverTab[8132]++
								return nil, StructuralError{"cannot represent time as UTCTime"}
//line /usr/local/go/src/encoding/asn1/marshal.go:402
		// _ = "end of CoverTab[8132]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:403
	// _ = "end of CoverTab[8128]"
//line /usr/local/go/src/encoding/asn1/marshal.go:403
	_go_fuzz_dep_.CoverTab[8129]++

							return appendTimeCommon(dst, t), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:405
	// _ = "end of CoverTab[8129]"
}

func appendGeneralizedTime(dst []byte, t time.Time) (ret []byte, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:408
	_go_fuzz_dep_.CoverTab[8135]++
							year := t.Year()
							if year < 0 || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:410
		_go_fuzz_dep_.CoverTab[8137]++
//line /usr/local/go/src/encoding/asn1/marshal.go:410
		return year > 9999
//line /usr/local/go/src/encoding/asn1/marshal.go:410
		// _ = "end of CoverTab[8137]"
//line /usr/local/go/src/encoding/asn1/marshal.go:410
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:410
		_go_fuzz_dep_.CoverTab[8138]++
								return nil, StructuralError{"cannot represent time as GeneralizedTime"}
//line /usr/local/go/src/encoding/asn1/marshal.go:411
		// _ = "end of CoverTab[8138]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:412
		_go_fuzz_dep_.CoverTab[8139]++
//line /usr/local/go/src/encoding/asn1/marshal.go:412
		// _ = "end of CoverTab[8139]"
//line /usr/local/go/src/encoding/asn1/marshal.go:412
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:412
	// _ = "end of CoverTab[8135]"
//line /usr/local/go/src/encoding/asn1/marshal.go:412
	_go_fuzz_dep_.CoverTab[8136]++

							dst = appendFourDigits(dst, year)

							return appendTimeCommon(dst, t), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:416
	// _ = "end of CoverTab[8136]"
}

func appendTimeCommon(dst []byte, t time.Time) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:419
	_go_fuzz_dep_.CoverTab[8140]++
							_, month, day := t.Date()

							dst = appendTwoDigits(dst, int(month))
							dst = appendTwoDigits(dst, day)

							hour, min, sec := t.Clock()

							dst = appendTwoDigits(dst, hour)
							dst = appendTwoDigits(dst, min)
							dst = appendTwoDigits(dst, sec)

							_, offset := t.Zone()

							switch {
	case offset/60 == 0:
//line /usr/local/go/src/encoding/asn1/marshal.go:434
		_go_fuzz_dep_.CoverTab[8143]++
								return append(dst, 'Z')
//line /usr/local/go/src/encoding/asn1/marshal.go:435
		// _ = "end of CoverTab[8143]"
	case offset > 0:
//line /usr/local/go/src/encoding/asn1/marshal.go:436
		_go_fuzz_dep_.CoverTab[8144]++
								dst = append(dst, '+')
//line /usr/local/go/src/encoding/asn1/marshal.go:437
		// _ = "end of CoverTab[8144]"
	case offset < 0:
//line /usr/local/go/src/encoding/asn1/marshal.go:438
		_go_fuzz_dep_.CoverTab[8145]++
								dst = append(dst, '-')
//line /usr/local/go/src/encoding/asn1/marshal.go:439
		// _ = "end of CoverTab[8145]"
//line /usr/local/go/src/encoding/asn1/marshal.go:439
	default:
//line /usr/local/go/src/encoding/asn1/marshal.go:439
		_go_fuzz_dep_.CoverTab[8146]++
//line /usr/local/go/src/encoding/asn1/marshal.go:439
		// _ = "end of CoverTab[8146]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:440
	// _ = "end of CoverTab[8140]"
//line /usr/local/go/src/encoding/asn1/marshal.go:440
	_go_fuzz_dep_.CoverTab[8141]++

							offsetMinutes := offset / 60
							if offsetMinutes < 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:443
		_go_fuzz_dep_.CoverTab[8147]++
								offsetMinutes = -offsetMinutes
//line /usr/local/go/src/encoding/asn1/marshal.go:444
		// _ = "end of CoverTab[8147]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:445
		_go_fuzz_dep_.CoverTab[8148]++
//line /usr/local/go/src/encoding/asn1/marshal.go:445
		// _ = "end of CoverTab[8148]"
//line /usr/local/go/src/encoding/asn1/marshal.go:445
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:445
	// _ = "end of CoverTab[8141]"
//line /usr/local/go/src/encoding/asn1/marshal.go:445
	_go_fuzz_dep_.CoverTab[8142]++

							dst = appendTwoDigits(dst, offsetMinutes/60)
							dst = appendTwoDigits(dst, offsetMinutes%60)

							return dst
//line /usr/local/go/src/encoding/asn1/marshal.go:450
	// _ = "end of CoverTab[8142]"
}

func stripTagAndLength(in []byte) []byte {
//line /usr/local/go/src/encoding/asn1/marshal.go:453
	_go_fuzz_dep_.CoverTab[8149]++
							_, offset, err := parseTagAndLength(in, 0)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:455
		_go_fuzz_dep_.CoverTab[8151]++
								return in
//line /usr/local/go/src/encoding/asn1/marshal.go:456
		// _ = "end of CoverTab[8151]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:457
		_go_fuzz_dep_.CoverTab[8152]++
//line /usr/local/go/src/encoding/asn1/marshal.go:457
		// _ = "end of CoverTab[8152]"
//line /usr/local/go/src/encoding/asn1/marshal.go:457
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:457
	// _ = "end of CoverTab[8149]"
//line /usr/local/go/src/encoding/asn1/marshal.go:457
	_go_fuzz_dep_.CoverTab[8150]++
							return in[offset:]
//line /usr/local/go/src/encoding/asn1/marshal.go:458
	// _ = "end of CoverTab[8150]"
}

func makeBody(value reflect.Value, params fieldParameters) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:461
	_go_fuzz_dep_.CoverTab[8153]++
							switch value.Type() {
	case flagType:
//line /usr/local/go/src/encoding/asn1/marshal.go:463
		_go_fuzz_dep_.CoverTab[8156]++
								return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:464
		// _ = "end of CoverTab[8156]"
	case timeType:
//line /usr/local/go/src/encoding/asn1/marshal.go:465
		_go_fuzz_dep_.CoverTab[8157]++
								t := value.Interface().(time.Time)
								if params.timeType == TagGeneralizedTime || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:467
			_go_fuzz_dep_.CoverTab[8163]++
//line /usr/local/go/src/encoding/asn1/marshal.go:467
			return outsideUTCRange(t)
//line /usr/local/go/src/encoding/asn1/marshal.go:467
			// _ = "end of CoverTab[8163]"
//line /usr/local/go/src/encoding/asn1/marshal.go:467
		}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:467
			_go_fuzz_dep_.CoverTab[8164]++
									return makeGeneralizedTime(t)
//line /usr/local/go/src/encoding/asn1/marshal.go:468
			// _ = "end of CoverTab[8164]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:469
			_go_fuzz_dep_.CoverTab[8165]++
//line /usr/local/go/src/encoding/asn1/marshal.go:469
			// _ = "end of CoverTab[8165]"
//line /usr/local/go/src/encoding/asn1/marshal.go:469
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:469
		// _ = "end of CoverTab[8157]"
//line /usr/local/go/src/encoding/asn1/marshal.go:469
		_go_fuzz_dep_.CoverTab[8158]++
								return makeUTCTime(t)
//line /usr/local/go/src/encoding/asn1/marshal.go:470
		// _ = "end of CoverTab[8158]"
	case bitStringType:
//line /usr/local/go/src/encoding/asn1/marshal.go:471
		_go_fuzz_dep_.CoverTab[8159]++
								return bitStringEncoder(value.Interface().(BitString)), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:472
		// _ = "end of CoverTab[8159]"
	case objectIdentifierType:
//line /usr/local/go/src/encoding/asn1/marshal.go:473
		_go_fuzz_dep_.CoverTab[8160]++
								return makeObjectIdentifier(value.Interface().(ObjectIdentifier))
//line /usr/local/go/src/encoding/asn1/marshal.go:474
		// _ = "end of CoverTab[8160]"
	case bigIntType:
//line /usr/local/go/src/encoding/asn1/marshal.go:475
		_go_fuzz_dep_.CoverTab[8161]++
								return makeBigInt(value.Interface().(*big.Int))
//line /usr/local/go/src/encoding/asn1/marshal.go:476
		// _ = "end of CoverTab[8161]"
//line /usr/local/go/src/encoding/asn1/marshal.go:476
	default:
//line /usr/local/go/src/encoding/asn1/marshal.go:476
		_go_fuzz_dep_.CoverTab[8162]++
//line /usr/local/go/src/encoding/asn1/marshal.go:476
		// _ = "end of CoverTab[8162]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:477
	// _ = "end of CoverTab[8153]"
//line /usr/local/go/src/encoding/asn1/marshal.go:477
	_go_fuzz_dep_.CoverTab[8154]++

							switch v := value; v.Kind() {
	case reflect.Bool:
//line /usr/local/go/src/encoding/asn1/marshal.go:480
		_go_fuzz_dep_.CoverTab[8166]++
								if v.Bool() {
//line /usr/local/go/src/encoding/asn1/marshal.go:481
			_go_fuzz_dep_.CoverTab[8177]++
									return byteFFEncoder, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:482
			// _ = "end of CoverTab[8177]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:483
			_go_fuzz_dep_.CoverTab[8178]++
//line /usr/local/go/src/encoding/asn1/marshal.go:483
			// _ = "end of CoverTab[8178]"
//line /usr/local/go/src/encoding/asn1/marshal.go:483
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:483
		// _ = "end of CoverTab[8166]"
//line /usr/local/go/src/encoding/asn1/marshal.go:483
		_go_fuzz_dep_.CoverTab[8167]++
								return byte00Encoder, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:484
		// _ = "end of CoverTab[8167]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/asn1/marshal.go:485
		_go_fuzz_dep_.CoverTab[8168]++
								return int64Encoder(v.Int()), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:486
		// _ = "end of CoverTab[8168]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/asn1/marshal.go:487
		_go_fuzz_dep_.CoverTab[8169]++
								t := v.Type()

								for i := 0; i < t.NumField(); i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:490
			_go_fuzz_dep_.CoverTab[8179]++
									if !t.Field(i).IsExported() {
//line /usr/local/go/src/encoding/asn1/marshal.go:491
				_go_fuzz_dep_.CoverTab[8180]++
										return nil, StructuralError{"struct contains unexported fields"}
//line /usr/local/go/src/encoding/asn1/marshal.go:492
				// _ = "end of CoverTab[8180]"
			} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:493
				_go_fuzz_dep_.CoverTab[8181]++
//line /usr/local/go/src/encoding/asn1/marshal.go:493
				// _ = "end of CoverTab[8181]"
//line /usr/local/go/src/encoding/asn1/marshal.go:493
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:493
			// _ = "end of CoverTab[8179]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:494
		// _ = "end of CoverTab[8169]"
//line /usr/local/go/src/encoding/asn1/marshal.go:494
		_go_fuzz_dep_.CoverTab[8170]++

								startingField := 0

								n := t.NumField()
								if n == 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:499
			_go_fuzz_dep_.CoverTab[8182]++
									return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:500
			// _ = "end of CoverTab[8182]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:501
			_go_fuzz_dep_.CoverTab[8183]++
//line /usr/local/go/src/encoding/asn1/marshal.go:501
			// _ = "end of CoverTab[8183]"
//line /usr/local/go/src/encoding/asn1/marshal.go:501
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:501
		// _ = "end of CoverTab[8170]"
//line /usr/local/go/src/encoding/asn1/marshal.go:501
		_go_fuzz_dep_.CoverTab[8171]++

//line /usr/local/go/src/encoding/asn1/marshal.go:505
		if t.Field(0).Type == rawContentsType {
//line /usr/local/go/src/encoding/asn1/marshal.go:505
			_go_fuzz_dep_.CoverTab[8184]++
									s := v.Field(0)
									if s.Len() > 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:507
				_go_fuzz_dep_.CoverTab[8186]++
										bytes := s.Bytes()

//line /usr/local/go/src/encoding/asn1/marshal.go:513
				return bytesEncoder(stripTagAndLength(bytes)), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:513
				// _ = "end of CoverTab[8186]"
			} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:514
				_go_fuzz_dep_.CoverTab[8187]++
//line /usr/local/go/src/encoding/asn1/marshal.go:514
				// _ = "end of CoverTab[8187]"
//line /usr/local/go/src/encoding/asn1/marshal.go:514
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:514
			// _ = "end of CoverTab[8184]"
//line /usr/local/go/src/encoding/asn1/marshal.go:514
			_go_fuzz_dep_.CoverTab[8185]++

									startingField = 1
//line /usr/local/go/src/encoding/asn1/marshal.go:516
			// _ = "end of CoverTab[8185]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:517
			_go_fuzz_dep_.CoverTab[8188]++
//line /usr/local/go/src/encoding/asn1/marshal.go:517
			// _ = "end of CoverTab[8188]"
//line /usr/local/go/src/encoding/asn1/marshal.go:517
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:517
		// _ = "end of CoverTab[8171]"
//line /usr/local/go/src/encoding/asn1/marshal.go:517
		_go_fuzz_dep_.CoverTab[8172]++

								switch n1 := n - startingField; n1 {
		case 0:
//line /usr/local/go/src/encoding/asn1/marshal.go:520
			_go_fuzz_dep_.CoverTab[8189]++
									return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:521
			// _ = "end of CoverTab[8189]"
		case 1:
//line /usr/local/go/src/encoding/asn1/marshal.go:522
			_go_fuzz_dep_.CoverTab[8190]++
									return makeField(v.Field(startingField), parseFieldParameters(t.Field(startingField).Tag.Get("asn1")))
//line /usr/local/go/src/encoding/asn1/marshal.go:523
			// _ = "end of CoverTab[8190]"
		default:
//line /usr/local/go/src/encoding/asn1/marshal.go:524
			_go_fuzz_dep_.CoverTab[8191]++
									m := make([]encoder, n1)
									for i := 0; i < n1; i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:526
				_go_fuzz_dep_.CoverTab[8193]++
										m[i], err = makeField(v.Field(i+startingField), parseFieldParameters(t.Field(i+startingField).Tag.Get("asn1")))
										if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:528
					_go_fuzz_dep_.CoverTab[8194]++
											return nil, err
//line /usr/local/go/src/encoding/asn1/marshal.go:529
					// _ = "end of CoverTab[8194]"
				} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:530
					_go_fuzz_dep_.CoverTab[8195]++
//line /usr/local/go/src/encoding/asn1/marshal.go:530
					// _ = "end of CoverTab[8195]"
//line /usr/local/go/src/encoding/asn1/marshal.go:530
				}
//line /usr/local/go/src/encoding/asn1/marshal.go:530
				// _ = "end of CoverTab[8193]"
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:531
			// _ = "end of CoverTab[8191]"
//line /usr/local/go/src/encoding/asn1/marshal.go:531
			_go_fuzz_dep_.CoverTab[8192]++

									return multiEncoder(m), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:533
			// _ = "end of CoverTab[8192]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:534
		// _ = "end of CoverTab[8172]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/asn1/marshal.go:535
		_go_fuzz_dep_.CoverTab[8173]++
								sliceType := v.Type()
								if sliceType.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/asn1/marshal.go:537
			_go_fuzz_dep_.CoverTab[8196]++
									return bytesEncoder(v.Bytes()), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:538
			// _ = "end of CoverTab[8196]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:539
			_go_fuzz_dep_.CoverTab[8197]++
//line /usr/local/go/src/encoding/asn1/marshal.go:539
			// _ = "end of CoverTab[8197]"
//line /usr/local/go/src/encoding/asn1/marshal.go:539
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:539
		// _ = "end of CoverTab[8173]"
//line /usr/local/go/src/encoding/asn1/marshal.go:539
		_go_fuzz_dep_.CoverTab[8174]++

								var fp fieldParameters

								switch l := v.Len(); l {
		case 0:
//line /usr/local/go/src/encoding/asn1/marshal.go:544
			_go_fuzz_dep_.CoverTab[8198]++
									return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:545
			// _ = "end of CoverTab[8198]"
		case 1:
//line /usr/local/go/src/encoding/asn1/marshal.go:546
			_go_fuzz_dep_.CoverTab[8199]++
									return makeField(v.Index(0), fp)
//line /usr/local/go/src/encoding/asn1/marshal.go:547
			// _ = "end of CoverTab[8199]"
		default:
//line /usr/local/go/src/encoding/asn1/marshal.go:548
			_go_fuzz_dep_.CoverTab[8200]++
									m := make([]encoder, l)

									for i := 0; i < l; i++ {
//line /usr/local/go/src/encoding/asn1/marshal.go:551
				_go_fuzz_dep_.CoverTab[8203]++
										m[i], err = makeField(v.Index(i), fp)
										if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:553
					_go_fuzz_dep_.CoverTab[8204]++
											return nil, err
//line /usr/local/go/src/encoding/asn1/marshal.go:554
					// _ = "end of CoverTab[8204]"
				} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:555
					_go_fuzz_dep_.CoverTab[8205]++
//line /usr/local/go/src/encoding/asn1/marshal.go:555
					// _ = "end of CoverTab[8205]"
//line /usr/local/go/src/encoding/asn1/marshal.go:555
				}
//line /usr/local/go/src/encoding/asn1/marshal.go:555
				// _ = "end of CoverTab[8203]"
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:556
			// _ = "end of CoverTab[8200]"
//line /usr/local/go/src/encoding/asn1/marshal.go:556
			_go_fuzz_dep_.CoverTab[8201]++

									if params.set {
//line /usr/local/go/src/encoding/asn1/marshal.go:558
				_go_fuzz_dep_.CoverTab[8206]++
										return setEncoder(m), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:559
				// _ = "end of CoverTab[8206]"
			} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:560
				_go_fuzz_dep_.CoverTab[8207]++
//line /usr/local/go/src/encoding/asn1/marshal.go:560
				// _ = "end of CoverTab[8207]"
//line /usr/local/go/src/encoding/asn1/marshal.go:560
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:560
			// _ = "end of CoverTab[8201]"
//line /usr/local/go/src/encoding/asn1/marshal.go:560
			_go_fuzz_dep_.CoverTab[8202]++
									return multiEncoder(m), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:561
			// _ = "end of CoverTab[8202]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:562
		// _ = "end of CoverTab[8174]"
	case reflect.String:
//line /usr/local/go/src/encoding/asn1/marshal.go:563
		_go_fuzz_dep_.CoverTab[8175]++
								switch params.stringType {
		case TagIA5String:
//line /usr/local/go/src/encoding/asn1/marshal.go:565
			_go_fuzz_dep_.CoverTab[8208]++
									return makeIA5String(v.String())
//line /usr/local/go/src/encoding/asn1/marshal.go:566
			// _ = "end of CoverTab[8208]"
		case TagPrintableString:
//line /usr/local/go/src/encoding/asn1/marshal.go:567
			_go_fuzz_dep_.CoverTab[8209]++
									return makePrintableString(v.String())
//line /usr/local/go/src/encoding/asn1/marshal.go:568
			// _ = "end of CoverTab[8209]"
		case TagNumericString:
//line /usr/local/go/src/encoding/asn1/marshal.go:569
			_go_fuzz_dep_.CoverTab[8210]++
									return makeNumericString(v.String())
//line /usr/local/go/src/encoding/asn1/marshal.go:570
			// _ = "end of CoverTab[8210]"
		default:
//line /usr/local/go/src/encoding/asn1/marshal.go:571
			_go_fuzz_dep_.CoverTab[8211]++
									return makeUTF8String(v.String()), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:572
			// _ = "end of CoverTab[8211]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:573
		// _ = "end of CoverTab[8175]"
//line /usr/local/go/src/encoding/asn1/marshal.go:573
	default:
//line /usr/local/go/src/encoding/asn1/marshal.go:573
		_go_fuzz_dep_.CoverTab[8176]++
//line /usr/local/go/src/encoding/asn1/marshal.go:573
		// _ = "end of CoverTab[8176]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:574
	// _ = "end of CoverTab[8154]"
//line /usr/local/go/src/encoding/asn1/marshal.go:574
	_go_fuzz_dep_.CoverTab[8155]++

							return nil, StructuralError{"unknown Go type"}
//line /usr/local/go/src/encoding/asn1/marshal.go:576
	// _ = "end of CoverTab[8155]"
}

func makeField(v reflect.Value, params fieldParameters) (e encoder, err error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:579
	_go_fuzz_dep_.CoverTab[8212]++
							if !v.IsValid() {
//line /usr/local/go/src/encoding/asn1/marshal.go:580
		_go_fuzz_dep_.CoverTab[8227]++
								return nil, fmt.Errorf("asn1: cannot marshal nil value")
//line /usr/local/go/src/encoding/asn1/marshal.go:581
		// _ = "end of CoverTab[8227]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:582
		_go_fuzz_dep_.CoverTab[8228]++
//line /usr/local/go/src/encoding/asn1/marshal.go:582
		// _ = "end of CoverTab[8228]"
//line /usr/local/go/src/encoding/asn1/marshal.go:582
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:582
	// _ = "end of CoverTab[8212]"
//line /usr/local/go/src/encoding/asn1/marshal.go:582
	_go_fuzz_dep_.CoverTab[8213]++

							if v.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:584
		_go_fuzz_dep_.CoverTab[8229]++
//line /usr/local/go/src/encoding/asn1/marshal.go:584
		return v.Type().NumMethod() == 0
//line /usr/local/go/src/encoding/asn1/marshal.go:584
		// _ = "end of CoverTab[8229]"
//line /usr/local/go/src/encoding/asn1/marshal.go:584
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:584
		_go_fuzz_dep_.CoverTab[8230]++
								return makeField(v.Elem(), params)
//line /usr/local/go/src/encoding/asn1/marshal.go:585
		// _ = "end of CoverTab[8230]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:586
		_go_fuzz_dep_.CoverTab[8231]++
//line /usr/local/go/src/encoding/asn1/marshal.go:586
		// _ = "end of CoverTab[8231]"
//line /usr/local/go/src/encoding/asn1/marshal.go:586
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:586
	// _ = "end of CoverTab[8213]"
//line /usr/local/go/src/encoding/asn1/marshal.go:586
	_go_fuzz_dep_.CoverTab[8214]++

							if v.Kind() == reflect.Slice && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		_go_fuzz_dep_.CoverTab[8232]++
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		return v.Len() == 0
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		// _ = "end of CoverTab[8232]"
//line /usr/local/go/src/encoding/asn1/marshal.go:588
	}() && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		_go_fuzz_dep_.CoverTab[8233]++
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		return params.omitEmpty
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		// _ = "end of CoverTab[8233]"
//line /usr/local/go/src/encoding/asn1/marshal.go:588
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:588
		_go_fuzz_dep_.CoverTab[8234]++
								return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:589
		// _ = "end of CoverTab[8234]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:590
		_go_fuzz_dep_.CoverTab[8235]++
//line /usr/local/go/src/encoding/asn1/marshal.go:590
		// _ = "end of CoverTab[8235]"
//line /usr/local/go/src/encoding/asn1/marshal.go:590
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:590
	// _ = "end of CoverTab[8214]"
//line /usr/local/go/src/encoding/asn1/marshal.go:590
	_go_fuzz_dep_.CoverTab[8215]++

							if params.optional && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		_go_fuzz_dep_.CoverTab[8236]++
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		return params.defaultValue != nil
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		// _ = "end of CoverTab[8236]"
//line /usr/local/go/src/encoding/asn1/marshal.go:592
	}() && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		_go_fuzz_dep_.CoverTab[8237]++
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		return canHaveDefaultValue(v.Kind())
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		// _ = "end of CoverTab[8237]"
//line /usr/local/go/src/encoding/asn1/marshal.go:592
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:592
		_go_fuzz_dep_.CoverTab[8238]++
								defaultValue := reflect.New(v.Type()).Elem()
								defaultValue.SetInt(*params.defaultValue)

								if reflect.DeepEqual(v.Interface(), defaultValue.Interface()) {
//line /usr/local/go/src/encoding/asn1/marshal.go:596
			_go_fuzz_dep_.CoverTab[8239]++
									return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:597
			// _ = "end of CoverTab[8239]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:598
			_go_fuzz_dep_.CoverTab[8240]++
//line /usr/local/go/src/encoding/asn1/marshal.go:598
			// _ = "end of CoverTab[8240]"
//line /usr/local/go/src/encoding/asn1/marshal.go:598
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:598
		// _ = "end of CoverTab[8238]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:599
		_go_fuzz_dep_.CoverTab[8241]++
//line /usr/local/go/src/encoding/asn1/marshal.go:599
		// _ = "end of CoverTab[8241]"
//line /usr/local/go/src/encoding/asn1/marshal.go:599
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:599
	// _ = "end of CoverTab[8215]"
//line /usr/local/go/src/encoding/asn1/marshal.go:599
	_go_fuzz_dep_.CoverTab[8216]++

//line /usr/local/go/src/encoding/asn1/marshal.go:604
	if params.optional && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:604
		_go_fuzz_dep_.CoverTab[8242]++
//line /usr/local/go/src/encoding/asn1/marshal.go:604
		return params.defaultValue == nil
//line /usr/local/go/src/encoding/asn1/marshal.go:604
		// _ = "end of CoverTab[8242]"
//line /usr/local/go/src/encoding/asn1/marshal.go:604
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:604
		_go_fuzz_dep_.CoverTab[8243]++
								if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
//line /usr/local/go/src/encoding/asn1/marshal.go:605
			_go_fuzz_dep_.CoverTab[8244]++
									return bytesEncoder(nil), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:606
			// _ = "end of CoverTab[8244]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:607
			_go_fuzz_dep_.CoverTab[8245]++
//line /usr/local/go/src/encoding/asn1/marshal.go:607
			// _ = "end of CoverTab[8245]"
//line /usr/local/go/src/encoding/asn1/marshal.go:607
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:607
		// _ = "end of CoverTab[8243]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:608
		_go_fuzz_dep_.CoverTab[8246]++
//line /usr/local/go/src/encoding/asn1/marshal.go:608
		// _ = "end of CoverTab[8246]"
//line /usr/local/go/src/encoding/asn1/marshal.go:608
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:608
	// _ = "end of CoverTab[8216]"
//line /usr/local/go/src/encoding/asn1/marshal.go:608
	_go_fuzz_dep_.CoverTab[8217]++

							if v.Type() == rawValueType {
//line /usr/local/go/src/encoding/asn1/marshal.go:610
		_go_fuzz_dep_.CoverTab[8247]++
								rv := v.Interface().(RawValue)
								if len(rv.FullBytes) != 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:612
			_go_fuzz_dep_.CoverTab[8249]++
									return bytesEncoder(rv.FullBytes), nil
//line /usr/local/go/src/encoding/asn1/marshal.go:613
			// _ = "end of CoverTab[8249]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:614
			_go_fuzz_dep_.CoverTab[8250]++
//line /usr/local/go/src/encoding/asn1/marshal.go:614
			// _ = "end of CoverTab[8250]"
//line /usr/local/go/src/encoding/asn1/marshal.go:614
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:614
		// _ = "end of CoverTab[8247]"
//line /usr/local/go/src/encoding/asn1/marshal.go:614
		_go_fuzz_dep_.CoverTab[8248]++

								t := new(taggedEncoder)

								t.tag = bytesEncoder(appendTagAndLength(t.scratch[:0], tagAndLength{rv.Class, rv.Tag, len(rv.Bytes), rv.IsCompound}))
								t.body = bytesEncoder(rv.Bytes)

								return t, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:621
		// _ = "end of CoverTab[8248]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:622
		_go_fuzz_dep_.CoverTab[8251]++
//line /usr/local/go/src/encoding/asn1/marshal.go:622
		// _ = "end of CoverTab[8251]"
//line /usr/local/go/src/encoding/asn1/marshal.go:622
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:622
	// _ = "end of CoverTab[8217]"
//line /usr/local/go/src/encoding/asn1/marshal.go:622
	_go_fuzz_dep_.CoverTab[8218]++

							matchAny, tag, isCompound, ok := getUniversalType(v.Type())
							if !ok || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:625
		_go_fuzz_dep_.CoverTab[8252]++
//line /usr/local/go/src/encoding/asn1/marshal.go:625
		return matchAny
//line /usr/local/go/src/encoding/asn1/marshal.go:625
		// _ = "end of CoverTab[8252]"
//line /usr/local/go/src/encoding/asn1/marshal.go:625
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:625
		_go_fuzz_dep_.CoverTab[8253]++
								return nil, StructuralError{fmt.Sprintf("unknown Go type: %v", v.Type())}
//line /usr/local/go/src/encoding/asn1/marshal.go:626
		// _ = "end of CoverTab[8253]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:627
		_go_fuzz_dep_.CoverTab[8254]++
//line /usr/local/go/src/encoding/asn1/marshal.go:627
		// _ = "end of CoverTab[8254]"
//line /usr/local/go/src/encoding/asn1/marshal.go:627
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:627
	// _ = "end of CoverTab[8218]"
//line /usr/local/go/src/encoding/asn1/marshal.go:627
	_go_fuzz_dep_.CoverTab[8219]++

							if params.timeType != 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:629
		_go_fuzz_dep_.CoverTab[8255]++
//line /usr/local/go/src/encoding/asn1/marshal.go:629
		return tag != TagUTCTime
//line /usr/local/go/src/encoding/asn1/marshal.go:629
		// _ = "end of CoverTab[8255]"
//line /usr/local/go/src/encoding/asn1/marshal.go:629
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:629
		_go_fuzz_dep_.CoverTab[8256]++
								return nil, StructuralError{"explicit time type given to non-time member"}
//line /usr/local/go/src/encoding/asn1/marshal.go:630
		// _ = "end of CoverTab[8256]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:631
		_go_fuzz_dep_.CoverTab[8257]++
//line /usr/local/go/src/encoding/asn1/marshal.go:631
		// _ = "end of CoverTab[8257]"
//line /usr/local/go/src/encoding/asn1/marshal.go:631
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:631
	// _ = "end of CoverTab[8219]"
//line /usr/local/go/src/encoding/asn1/marshal.go:631
	_go_fuzz_dep_.CoverTab[8220]++

							if params.stringType != 0 && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:633
		_go_fuzz_dep_.CoverTab[8258]++
//line /usr/local/go/src/encoding/asn1/marshal.go:633
		return tag != TagPrintableString
//line /usr/local/go/src/encoding/asn1/marshal.go:633
		// _ = "end of CoverTab[8258]"
//line /usr/local/go/src/encoding/asn1/marshal.go:633
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:633
		_go_fuzz_dep_.CoverTab[8259]++
								return nil, StructuralError{"explicit string type given to non-string member"}
//line /usr/local/go/src/encoding/asn1/marshal.go:634
		// _ = "end of CoverTab[8259]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:635
		_go_fuzz_dep_.CoverTab[8260]++
//line /usr/local/go/src/encoding/asn1/marshal.go:635
		// _ = "end of CoverTab[8260]"
//line /usr/local/go/src/encoding/asn1/marshal.go:635
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:635
	// _ = "end of CoverTab[8220]"
//line /usr/local/go/src/encoding/asn1/marshal.go:635
	_go_fuzz_dep_.CoverTab[8221]++

							switch tag {
	case TagPrintableString:
//line /usr/local/go/src/encoding/asn1/marshal.go:638
		_go_fuzz_dep_.CoverTab[8261]++
								if params.stringType == 0 {
//line /usr/local/go/src/encoding/asn1/marshal.go:639
			_go_fuzz_dep_.CoverTab[8264]++

//line /usr/local/go/src/encoding/asn1/marshal.go:643
			for _, r := range v.String() {
//line /usr/local/go/src/encoding/asn1/marshal.go:643
				_go_fuzz_dep_.CoverTab[8265]++
										if r >= utf8.RuneSelf || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:644
					_go_fuzz_dep_.CoverTab[8266]++
//line /usr/local/go/src/encoding/asn1/marshal.go:644
					return !isPrintable(byte(r), rejectAsterisk, rejectAmpersand)
//line /usr/local/go/src/encoding/asn1/marshal.go:644
					// _ = "end of CoverTab[8266]"
//line /usr/local/go/src/encoding/asn1/marshal.go:644
				}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:644
					_go_fuzz_dep_.CoverTab[8267]++
											if !utf8.ValidString(v.String()) {
//line /usr/local/go/src/encoding/asn1/marshal.go:645
						_go_fuzz_dep_.CoverTab[8269]++
												return nil, errors.New("asn1: string not valid UTF-8")
//line /usr/local/go/src/encoding/asn1/marshal.go:646
						// _ = "end of CoverTab[8269]"
					} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:647
						_go_fuzz_dep_.CoverTab[8270]++
//line /usr/local/go/src/encoding/asn1/marshal.go:647
						// _ = "end of CoverTab[8270]"
//line /usr/local/go/src/encoding/asn1/marshal.go:647
					}
//line /usr/local/go/src/encoding/asn1/marshal.go:647
					// _ = "end of CoverTab[8267]"
//line /usr/local/go/src/encoding/asn1/marshal.go:647
					_go_fuzz_dep_.CoverTab[8268]++
											tag = TagUTF8String
											break
//line /usr/local/go/src/encoding/asn1/marshal.go:649
					// _ = "end of CoverTab[8268]"
				} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:650
					_go_fuzz_dep_.CoverTab[8271]++
//line /usr/local/go/src/encoding/asn1/marshal.go:650
					// _ = "end of CoverTab[8271]"
//line /usr/local/go/src/encoding/asn1/marshal.go:650
				}
//line /usr/local/go/src/encoding/asn1/marshal.go:650
				// _ = "end of CoverTab[8265]"
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:651
			// _ = "end of CoverTab[8264]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:652
			_go_fuzz_dep_.CoverTab[8272]++
									tag = params.stringType
//line /usr/local/go/src/encoding/asn1/marshal.go:653
			// _ = "end of CoverTab[8272]"
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:654
		// _ = "end of CoverTab[8261]"
	case TagUTCTime:
//line /usr/local/go/src/encoding/asn1/marshal.go:655
		_go_fuzz_dep_.CoverTab[8262]++
								if params.timeType == TagGeneralizedTime || func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:656
			_go_fuzz_dep_.CoverTab[8273]++
//line /usr/local/go/src/encoding/asn1/marshal.go:656
			return outsideUTCRange(v.Interface().(time.Time))
//line /usr/local/go/src/encoding/asn1/marshal.go:656
			// _ = "end of CoverTab[8273]"
//line /usr/local/go/src/encoding/asn1/marshal.go:656
		}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:656
			_go_fuzz_dep_.CoverTab[8274]++
									tag = TagGeneralizedTime
//line /usr/local/go/src/encoding/asn1/marshal.go:657
			// _ = "end of CoverTab[8274]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:658
			_go_fuzz_dep_.CoverTab[8275]++
//line /usr/local/go/src/encoding/asn1/marshal.go:658
			// _ = "end of CoverTab[8275]"
//line /usr/local/go/src/encoding/asn1/marshal.go:658
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:658
		// _ = "end of CoverTab[8262]"
//line /usr/local/go/src/encoding/asn1/marshal.go:658
	default:
//line /usr/local/go/src/encoding/asn1/marshal.go:658
		_go_fuzz_dep_.CoverTab[8263]++
//line /usr/local/go/src/encoding/asn1/marshal.go:658
		// _ = "end of CoverTab[8263]"
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:659
	// _ = "end of CoverTab[8221]"
//line /usr/local/go/src/encoding/asn1/marshal.go:659
	_go_fuzz_dep_.CoverTab[8222]++

							if params.set {
//line /usr/local/go/src/encoding/asn1/marshal.go:661
		_go_fuzz_dep_.CoverTab[8276]++
								if tag != TagSequence {
//line /usr/local/go/src/encoding/asn1/marshal.go:662
			_go_fuzz_dep_.CoverTab[8278]++
									return nil, StructuralError{"non sequence tagged as set"}
//line /usr/local/go/src/encoding/asn1/marshal.go:663
			// _ = "end of CoverTab[8278]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:664
			_go_fuzz_dep_.CoverTab[8279]++
//line /usr/local/go/src/encoding/asn1/marshal.go:664
			// _ = "end of CoverTab[8279]"
//line /usr/local/go/src/encoding/asn1/marshal.go:664
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:664
		// _ = "end of CoverTab[8276]"
//line /usr/local/go/src/encoding/asn1/marshal.go:664
		_go_fuzz_dep_.CoverTab[8277]++
								tag = TagSet
//line /usr/local/go/src/encoding/asn1/marshal.go:665
		// _ = "end of CoverTab[8277]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:666
		_go_fuzz_dep_.CoverTab[8280]++
//line /usr/local/go/src/encoding/asn1/marshal.go:666
		// _ = "end of CoverTab[8280]"
//line /usr/local/go/src/encoding/asn1/marshal.go:666
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:666
	// _ = "end of CoverTab[8222]"
//line /usr/local/go/src/encoding/asn1/marshal.go:666
	_go_fuzz_dep_.CoverTab[8223]++

//line /usr/local/go/src/encoding/asn1/marshal.go:673
	if tag == TagSet && func() bool {
//line /usr/local/go/src/encoding/asn1/marshal.go:673
		_go_fuzz_dep_.CoverTab[8281]++
//line /usr/local/go/src/encoding/asn1/marshal.go:673
		return !params.set
//line /usr/local/go/src/encoding/asn1/marshal.go:673
		// _ = "end of CoverTab[8281]"
//line /usr/local/go/src/encoding/asn1/marshal.go:673
	}() {
//line /usr/local/go/src/encoding/asn1/marshal.go:673
		_go_fuzz_dep_.CoverTab[8282]++
								params.set = true
//line /usr/local/go/src/encoding/asn1/marshal.go:674
		// _ = "end of CoverTab[8282]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:675
		_go_fuzz_dep_.CoverTab[8283]++
//line /usr/local/go/src/encoding/asn1/marshal.go:675
		// _ = "end of CoverTab[8283]"
//line /usr/local/go/src/encoding/asn1/marshal.go:675
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:675
	// _ = "end of CoverTab[8223]"
//line /usr/local/go/src/encoding/asn1/marshal.go:675
	_go_fuzz_dep_.CoverTab[8224]++

							t := new(taggedEncoder)

							t.body, err = makeBody(v, params)
							if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:680
		_go_fuzz_dep_.CoverTab[8284]++
								return nil, err
//line /usr/local/go/src/encoding/asn1/marshal.go:681
		// _ = "end of CoverTab[8284]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:682
		_go_fuzz_dep_.CoverTab[8285]++
//line /usr/local/go/src/encoding/asn1/marshal.go:682
		// _ = "end of CoverTab[8285]"
//line /usr/local/go/src/encoding/asn1/marshal.go:682
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:682
	// _ = "end of CoverTab[8224]"
//line /usr/local/go/src/encoding/asn1/marshal.go:682
	_go_fuzz_dep_.CoverTab[8225]++

							bodyLen := t.body.Len()

							class := ClassUniversal
							if params.tag != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:687
		_go_fuzz_dep_.CoverTab[8286]++
								if params.application {
//line /usr/local/go/src/encoding/asn1/marshal.go:688
			_go_fuzz_dep_.CoverTab[8289]++
									class = ClassApplication
//line /usr/local/go/src/encoding/asn1/marshal.go:689
			// _ = "end of CoverTab[8289]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:690
			_go_fuzz_dep_.CoverTab[8290]++
//line /usr/local/go/src/encoding/asn1/marshal.go:690
			if params.private {
//line /usr/local/go/src/encoding/asn1/marshal.go:690
				_go_fuzz_dep_.CoverTab[8291]++
										class = ClassPrivate
//line /usr/local/go/src/encoding/asn1/marshal.go:691
				// _ = "end of CoverTab[8291]"
			} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:692
				_go_fuzz_dep_.CoverTab[8292]++
										class = ClassContextSpecific
//line /usr/local/go/src/encoding/asn1/marshal.go:693
				// _ = "end of CoverTab[8292]"
			}
//line /usr/local/go/src/encoding/asn1/marshal.go:694
			// _ = "end of CoverTab[8290]"
//line /usr/local/go/src/encoding/asn1/marshal.go:694
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:694
		// _ = "end of CoverTab[8286]"
//line /usr/local/go/src/encoding/asn1/marshal.go:694
		_go_fuzz_dep_.CoverTab[8287]++

								if params.explicit {
//line /usr/local/go/src/encoding/asn1/marshal.go:696
			_go_fuzz_dep_.CoverTab[8293]++
									t.tag = bytesEncoder(appendTagAndLength(t.scratch[:0], tagAndLength{ClassUniversal, tag, bodyLen, isCompound}))

									tt := new(taggedEncoder)

									tt.body = t

									tt.tag = bytesEncoder(appendTagAndLength(tt.scratch[:0], tagAndLength{
				class:		class,
				tag:		*params.tag,
				length:		bodyLen + t.tag.Len(),
				isCompound:	true,
			}))

									return tt, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:710
			// _ = "end of CoverTab[8293]"
		} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:711
			_go_fuzz_dep_.CoverTab[8294]++
//line /usr/local/go/src/encoding/asn1/marshal.go:711
			// _ = "end of CoverTab[8294]"
//line /usr/local/go/src/encoding/asn1/marshal.go:711
		}
//line /usr/local/go/src/encoding/asn1/marshal.go:711
		// _ = "end of CoverTab[8287]"
//line /usr/local/go/src/encoding/asn1/marshal.go:711
		_go_fuzz_dep_.CoverTab[8288]++

//line /usr/local/go/src/encoding/asn1/marshal.go:714
		tag = *params.tag
//line /usr/local/go/src/encoding/asn1/marshal.go:714
		// _ = "end of CoverTab[8288]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:715
		_go_fuzz_dep_.CoverTab[8295]++
//line /usr/local/go/src/encoding/asn1/marshal.go:715
		// _ = "end of CoverTab[8295]"
//line /usr/local/go/src/encoding/asn1/marshal.go:715
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:715
	// _ = "end of CoverTab[8225]"
//line /usr/local/go/src/encoding/asn1/marshal.go:715
	_go_fuzz_dep_.CoverTab[8226]++

							t.tag = bytesEncoder(appendTagAndLength(t.scratch[:0], tagAndLength{class, tag, bodyLen, isCompound}))

							return t, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:719
	// _ = "end of CoverTab[8226]"
}

// Marshal returns the ASN.1 encoding of val.
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//
//line /usr/local/go/src/encoding/asn1/marshal.go:722
// In addition to the struct tags recognised by Unmarshal, the following can be
//line /usr/local/go/src/encoding/asn1/marshal.go:722
// used:
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//	ia5:         causes strings to be marshaled as ASN.1, IA5String values
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//	omitempty:   causes empty slices to be skipped
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//	printable:   causes strings to be marshaled as ASN.1, PrintableString values
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//	utf8:        causes strings to be marshaled as ASN.1, UTF8String values
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//	utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
//line /usr/local/go/src/encoding/asn1/marshal.go:722
//	generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
//line /usr/local/go/src/encoding/asn1/marshal.go:733
func Marshal(val any) ([]byte, error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:733
	_go_fuzz_dep_.CoverTab[8296]++
							return MarshalWithParams(val, "")
//line /usr/local/go/src/encoding/asn1/marshal.go:734
	// _ = "end of CoverTab[8296]"
}

// MarshalWithParams allows field parameters to be specified for the
//line /usr/local/go/src/encoding/asn1/marshal.go:737
// top-level element. The form of the params is the same as the field tags.
//line /usr/local/go/src/encoding/asn1/marshal.go:739
func MarshalWithParams(val any, params string) ([]byte, error) {
//line /usr/local/go/src/encoding/asn1/marshal.go:739
	_go_fuzz_dep_.CoverTab[8297]++
							e, err := makeField(reflect.ValueOf(val), parseFieldParameters(params))
							if err != nil {
//line /usr/local/go/src/encoding/asn1/marshal.go:741
		_go_fuzz_dep_.CoverTab[8299]++
								return nil, err
//line /usr/local/go/src/encoding/asn1/marshal.go:742
		// _ = "end of CoverTab[8299]"
	} else {
//line /usr/local/go/src/encoding/asn1/marshal.go:743
		_go_fuzz_dep_.CoverTab[8300]++
//line /usr/local/go/src/encoding/asn1/marshal.go:743
		// _ = "end of CoverTab[8300]"
//line /usr/local/go/src/encoding/asn1/marshal.go:743
	}
//line /usr/local/go/src/encoding/asn1/marshal.go:743
	// _ = "end of CoverTab[8297]"
//line /usr/local/go/src/encoding/asn1/marshal.go:743
	_go_fuzz_dep_.CoverTab[8298]++
							b := make([]byte, e.Len())
							e.Encode(b)
							return b, nil
//line /usr/local/go/src/encoding/asn1/marshal.go:746
	// _ = "end of CoverTab[8298]"
}

//line /usr/local/go/src/encoding/asn1/marshal.go:747
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/asn1/marshal.go:747
var _ = _go_fuzz_dep_.CoverTab
