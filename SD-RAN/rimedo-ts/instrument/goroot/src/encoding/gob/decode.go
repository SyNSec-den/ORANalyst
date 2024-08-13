// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run decgen.go -output dec_helpers.go

//line /usr/local/go/src/encoding/gob/decode.go:5
//go:generate go run decgen.go -output dec_helpers.go

package gob

//line /usr/local/go/src/encoding/gob/decode.go:7
import (
//line /usr/local/go/src/encoding/gob/decode.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/gob/decode.go:7
)
//line /usr/local/go/src/encoding/gob/decode.go:7
import (
//line /usr/local/go/src/encoding/gob/decode.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/gob/decode.go:7
)

import (
	"encoding"
	"errors"
	"internal/saferio"
	"io"
	"math"
	"math/bits"
	"reflect"
)

var (
	errBadUint	= errors.New("gob: encoded unsigned integer out of range")
	errBadType	= errors.New("gob: unknown type id or corrupted data")
	errRange	= errors.New("gob: bad data: field numbers out of bounds")
)

type decHelper func(state *decoderState, v reflect.Value, length int, ovfl error) bool

//line /usr/local/go/src/encoding/gob/decode.go:29
type decoderState struct {
							dec	*Decoder

//line /usr/local/go/src/encoding/gob/decode.go:33
	b	*decBuffer
	fieldnum	int
	next		*decoderState
}

//line /usr/local/go/src/encoding/gob/decode.go:40
type decBuffer struct {
	data	[]byte
	offset	int
}

func (d *decBuffer) Read(p []byte) (int, error) {
//line /usr/local/go/src/encoding/gob/decode.go:45
	_go_fuzz_dep_.CoverTab[84031]++
							n := copy(p, d.data[d.offset:])
							if n == 0 && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:47
		_go_fuzz_dep_.CoverTab[84033]++
//line /usr/local/go/src/encoding/gob/decode.go:47
		return len(p) != 0
//line /usr/local/go/src/encoding/gob/decode.go:47
		// _ = "end of CoverTab[84033]"
//line /usr/local/go/src/encoding/gob/decode.go:47
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:47
		_go_fuzz_dep_.CoverTab[84034]++
								return 0, io.EOF
//line /usr/local/go/src/encoding/gob/decode.go:48
		// _ = "end of CoverTab[84034]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:49
		_go_fuzz_dep_.CoverTab[84035]++
//line /usr/local/go/src/encoding/gob/decode.go:49
		// _ = "end of CoverTab[84035]"
//line /usr/local/go/src/encoding/gob/decode.go:49
	}
//line /usr/local/go/src/encoding/gob/decode.go:49
	// _ = "end of CoverTab[84031]"
//line /usr/local/go/src/encoding/gob/decode.go:49
	_go_fuzz_dep_.CoverTab[84032]++
							d.offset += n
							return n, nil
//line /usr/local/go/src/encoding/gob/decode.go:51
	// _ = "end of CoverTab[84032]"
}

func (d *decBuffer) Drop(n int) {
//line /usr/local/go/src/encoding/gob/decode.go:54
	_go_fuzz_dep_.CoverTab[84036]++
							if n > d.Len() {
//line /usr/local/go/src/encoding/gob/decode.go:55
		_go_fuzz_dep_.CoverTab[84038]++
								panic("drop")
//line /usr/local/go/src/encoding/gob/decode.go:56
		// _ = "end of CoverTab[84038]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:57
		_go_fuzz_dep_.CoverTab[84039]++
//line /usr/local/go/src/encoding/gob/decode.go:57
		// _ = "end of CoverTab[84039]"
//line /usr/local/go/src/encoding/gob/decode.go:57
	}
//line /usr/local/go/src/encoding/gob/decode.go:57
	// _ = "end of CoverTab[84036]"
//line /usr/local/go/src/encoding/gob/decode.go:57
	_go_fuzz_dep_.CoverTab[84037]++
							d.offset += n
//line /usr/local/go/src/encoding/gob/decode.go:58
	// _ = "end of CoverTab[84037]"
}

func (d *decBuffer) ReadByte() (byte, error) {
//line /usr/local/go/src/encoding/gob/decode.go:61
	_go_fuzz_dep_.CoverTab[84040]++
							if d.offset >= len(d.data) {
//line /usr/local/go/src/encoding/gob/decode.go:62
		_go_fuzz_dep_.CoverTab[84042]++
								return 0, io.EOF
//line /usr/local/go/src/encoding/gob/decode.go:63
		// _ = "end of CoverTab[84042]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:64
		_go_fuzz_dep_.CoverTab[84043]++
//line /usr/local/go/src/encoding/gob/decode.go:64
		// _ = "end of CoverTab[84043]"
//line /usr/local/go/src/encoding/gob/decode.go:64
	}
//line /usr/local/go/src/encoding/gob/decode.go:64
	// _ = "end of CoverTab[84040]"
//line /usr/local/go/src/encoding/gob/decode.go:64
	_go_fuzz_dep_.CoverTab[84041]++
							c := d.data[d.offset]
							d.offset++
							return c, nil
//line /usr/local/go/src/encoding/gob/decode.go:67
	// _ = "end of CoverTab[84041]"
}

func (d *decBuffer) Len() int {
//line /usr/local/go/src/encoding/gob/decode.go:70
	_go_fuzz_dep_.CoverTab[84044]++
							return len(d.data) - d.offset
//line /usr/local/go/src/encoding/gob/decode.go:71
	// _ = "end of CoverTab[84044]"
}

func (d *decBuffer) Bytes() []byte {
//line /usr/local/go/src/encoding/gob/decode.go:74
	_go_fuzz_dep_.CoverTab[84045]++
							return d.data[d.offset:]
//line /usr/local/go/src/encoding/gob/decode.go:75
	// _ = "end of CoverTab[84045]"
}

//line /usr/local/go/src/encoding/gob/decode.go:79
func (d *decBuffer) SetBytes(data []byte) {
//line /usr/local/go/src/encoding/gob/decode.go:79
	_go_fuzz_dep_.CoverTab[84046]++
							d.data = data
							d.offset = 0
//line /usr/local/go/src/encoding/gob/decode.go:81
	// _ = "end of CoverTab[84046]"
}

func (d *decBuffer) Reset() {
//line /usr/local/go/src/encoding/gob/decode.go:84
	_go_fuzz_dep_.CoverTab[84047]++
							d.data = d.data[0:0]
							d.offset = 0
//line /usr/local/go/src/encoding/gob/decode.go:86
	// _ = "end of CoverTab[84047]"
}

//line /usr/local/go/src/encoding/gob/decode.go:91
func (dec *Decoder) newDecoderState(buf *decBuffer) *decoderState {
//line /usr/local/go/src/encoding/gob/decode.go:91
	_go_fuzz_dep_.CoverTab[84048]++
							d := dec.freeList
							if d == nil {
//line /usr/local/go/src/encoding/gob/decode.go:93
		_go_fuzz_dep_.CoverTab[84050]++
								d = new(decoderState)
								d.dec = dec
//line /usr/local/go/src/encoding/gob/decode.go:95
		// _ = "end of CoverTab[84050]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:96
		_go_fuzz_dep_.CoverTab[84051]++
								dec.freeList = d.next
//line /usr/local/go/src/encoding/gob/decode.go:97
		// _ = "end of CoverTab[84051]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:98
	// _ = "end of CoverTab[84048]"
//line /usr/local/go/src/encoding/gob/decode.go:98
	_go_fuzz_dep_.CoverTab[84049]++
							d.b = buf
							return d
//line /usr/local/go/src/encoding/gob/decode.go:100
	// _ = "end of CoverTab[84049]"
}

func (dec *Decoder) freeDecoderState(d *decoderState) {
//line /usr/local/go/src/encoding/gob/decode.go:103
	_go_fuzz_dep_.CoverTab[84052]++
							d.next = dec.freeList
							dec.freeList = d
//line /usr/local/go/src/encoding/gob/decode.go:105
	// _ = "end of CoverTab[84052]"
}

func overflow(name string) error {
//line /usr/local/go/src/encoding/gob/decode.go:108
	_go_fuzz_dep_.CoverTab[84053]++
							return errors.New(`value for "` + name + `" out of range`)
//line /usr/local/go/src/encoding/gob/decode.go:109
	// _ = "end of CoverTab[84053]"
}

//line /usr/local/go/src/encoding/gob/decode.go:114
func decodeUintReader(r io.Reader, buf []byte) (x uint64, width int, err error) {
//line /usr/local/go/src/encoding/gob/decode.go:114
	_go_fuzz_dep_.CoverTab[84054]++
							width = 1
							n, err := io.ReadFull(r, buf[0:width])
							if n == 0 {
//line /usr/local/go/src/encoding/gob/decode.go:117
		_go_fuzz_dep_.CoverTab[84060]++
								return
//line /usr/local/go/src/encoding/gob/decode.go:118
		// _ = "end of CoverTab[84060]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:119
		_go_fuzz_dep_.CoverTab[84061]++
//line /usr/local/go/src/encoding/gob/decode.go:119
		// _ = "end of CoverTab[84061]"
//line /usr/local/go/src/encoding/gob/decode.go:119
	}
//line /usr/local/go/src/encoding/gob/decode.go:119
	// _ = "end of CoverTab[84054]"
//line /usr/local/go/src/encoding/gob/decode.go:119
	_go_fuzz_dep_.CoverTab[84055]++
							b := buf[0]
							if b <= 0x7f {
//line /usr/local/go/src/encoding/gob/decode.go:121
		_go_fuzz_dep_.CoverTab[84062]++
								return uint64(b), width, nil
//line /usr/local/go/src/encoding/gob/decode.go:122
		// _ = "end of CoverTab[84062]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:123
		_go_fuzz_dep_.CoverTab[84063]++
//line /usr/local/go/src/encoding/gob/decode.go:123
		// _ = "end of CoverTab[84063]"
//line /usr/local/go/src/encoding/gob/decode.go:123
	}
//line /usr/local/go/src/encoding/gob/decode.go:123
	// _ = "end of CoverTab[84055]"
//line /usr/local/go/src/encoding/gob/decode.go:123
	_go_fuzz_dep_.CoverTab[84056]++
							n = -int(int8(b))
							if n > uint64Size {
//line /usr/local/go/src/encoding/gob/decode.go:125
		_go_fuzz_dep_.CoverTab[84064]++
								err = errBadUint
								return
//line /usr/local/go/src/encoding/gob/decode.go:127
		// _ = "end of CoverTab[84064]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:128
		_go_fuzz_dep_.CoverTab[84065]++
//line /usr/local/go/src/encoding/gob/decode.go:128
		// _ = "end of CoverTab[84065]"
//line /usr/local/go/src/encoding/gob/decode.go:128
	}
//line /usr/local/go/src/encoding/gob/decode.go:128
	// _ = "end of CoverTab[84056]"
//line /usr/local/go/src/encoding/gob/decode.go:128
	_go_fuzz_dep_.CoverTab[84057]++
							width, err = io.ReadFull(r, buf[0:n])
							if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:130
		_go_fuzz_dep_.CoverTab[84066]++
								if err == io.EOF {
//line /usr/local/go/src/encoding/gob/decode.go:131
			_go_fuzz_dep_.CoverTab[84068]++
									err = io.ErrUnexpectedEOF
//line /usr/local/go/src/encoding/gob/decode.go:132
			// _ = "end of CoverTab[84068]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:133
			_go_fuzz_dep_.CoverTab[84069]++
//line /usr/local/go/src/encoding/gob/decode.go:133
			// _ = "end of CoverTab[84069]"
//line /usr/local/go/src/encoding/gob/decode.go:133
		}
//line /usr/local/go/src/encoding/gob/decode.go:133
		// _ = "end of CoverTab[84066]"
//line /usr/local/go/src/encoding/gob/decode.go:133
		_go_fuzz_dep_.CoverTab[84067]++
								return
//line /usr/local/go/src/encoding/gob/decode.go:134
		// _ = "end of CoverTab[84067]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:135
		_go_fuzz_dep_.CoverTab[84070]++
//line /usr/local/go/src/encoding/gob/decode.go:135
		// _ = "end of CoverTab[84070]"
//line /usr/local/go/src/encoding/gob/decode.go:135
	}
//line /usr/local/go/src/encoding/gob/decode.go:135
	// _ = "end of CoverTab[84057]"
//line /usr/local/go/src/encoding/gob/decode.go:135
	_go_fuzz_dep_.CoverTab[84058]++

							for _, b := range buf[0:width] {
//line /usr/local/go/src/encoding/gob/decode.go:137
		_go_fuzz_dep_.CoverTab[84071]++
								x = x<<8 | uint64(b)
//line /usr/local/go/src/encoding/gob/decode.go:138
		// _ = "end of CoverTab[84071]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:139
	// _ = "end of CoverTab[84058]"
//line /usr/local/go/src/encoding/gob/decode.go:139
	_go_fuzz_dep_.CoverTab[84059]++
							width++
							return
//line /usr/local/go/src/encoding/gob/decode.go:141
	// _ = "end of CoverTab[84059]"
}

//line /usr/local/go/src/encoding/gob/decode.go:146
func (state *decoderState) decodeUint() (x uint64) {
//line /usr/local/go/src/encoding/gob/decode.go:146
	_go_fuzz_dep_.CoverTab[84072]++
							b, err := state.b.ReadByte()
							if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:148
		_go_fuzz_dep_.CoverTab[84078]++
								error_(err)
//line /usr/local/go/src/encoding/gob/decode.go:149
		// _ = "end of CoverTab[84078]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:150
		_go_fuzz_dep_.CoverTab[84079]++
//line /usr/local/go/src/encoding/gob/decode.go:150
		// _ = "end of CoverTab[84079]"
//line /usr/local/go/src/encoding/gob/decode.go:150
	}
//line /usr/local/go/src/encoding/gob/decode.go:150
	// _ = "end of CoverTab[84072]"
//line /usr/local/go/src/encoding/gob/decode.go:150
	_go_fuzz_dep_.CoverTab[84073]++
							if b <= 0x7f {
//line /usr/local/go/src/encoding/gob/decode.go:151
		_go_fuzz_dep_.CoverTab[84080]++
								return uint64(b)
//line /usr/local/go/src/encoding/gob/decode.go:152
		// _ = "end of CoverTab[84080]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:153
		_go_fuzz_dep_.CoverTab[84081]++
//line /usr/local/go/src/encoding/gob/decode.go:153
		// _ = "end of CoverTab[84081]"
//line /usr/local/go/src/encoding/gob/decode.go:153
	}
//line /usr/local/go/src/encoding/gob/decode.go:153
	// _ = "end of CoverTab[84073]"
//line /usr/local/go/src/encoding/gob/decode.go:153
	_go_fuzz_dep_.CoverTab[84074]++
							n := -int(int8(b))
							if n > uint64Size {
//line /usr/local/go/src/encoding/gob/decode.go:155
		_go_fuzz_dep_.CoverTab[84082]++
								error_(errBadUint)
//line /usr/local/go/src/encoding/gob/decode.go:156
		// _ = "end of CoverTab[84082]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:157
		_go_fuzz_dep_.CoverTab[84083]++
//line /usr/local/go/src/encoding/gob/decode.go:157
		// _ = "end of CoverTab[84083]"
//line /usr/local/go/src/encoding/gob/decode.go:157
	}
//line /usr/local/go/src/encoding/gob/decode.go:157
	// _ = "end of CoverTab[84074]"
//line /usr/local/go/src/encoding/gob/decode.go:157
	_go_fuzz_dep_.CoverTab[84075]++
							buf := state.b.Bytes()
							if len(buf) < n {
//line /usr/local/go/src/encoding/gob/decode.go:159
		_go_fuzz_dep_.CoverTab[84084]++
								errorf("invalid uint data length %d: exceeds input size %d", n, len(buf))
//line /usr/local/go/src/encoding/gob/decode.go:160
		// _ = "end of CoverTab[84084]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:161
		_go_fuzz_dep_.CoverTab[84085]++
//line /usr/local/go/src/encoding/gob/decode.go:161
		// _ = "end of CoverTab[84085]"
//line /usr/local/go/src/encoding/gob/decode.go:161
	}
//line /usr/local/go/src/encoding/gob/decode.go:161
	// _ = "end of CoverTab[84075]"
//line /usr/local/go/src/encoding/gob/decode.go:161
	_go_fuzz_dep_.CoverTab[84076]++

//line /usr/local/go/src/encoding/gob/decode.go:164
	for _, b := range buf[0:n] {
//line /usr/local/go/src/encoding/gob/decode.go:164
		_go_fuzz_dep_.CoverTab[84086]++
								x = x<<8 | uint64(b)
//line /usr/local/go/src/encoding/gob/decode.go:165
		// _ = "end of CoverTab[84086]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:166
	// _ = "end of CoverTab[84076]"
//line /usr/local/go/src/encoding/gob/decode.go:166
	_go_fuzz_dep_.CoverTab[84077]++
							state.b.Drop(n)
							return x
//line /usr/local/go/src/encoding/gob/decode.go:168
	// _ = "end of CoverTab[84077]"
}

//line /usr/local/go/src/encoding/gob/decode.go:173
func (state *decoderState) decodeInt() int64 {
//line /usr/local/go/src/encoding/gob/decode.go:173
	_go_fuzz_dep_.CoverTab[84087]++
							x := state.decodeUint()
							if x&1 != 0 {
//line /usr/local/go/src/encoding/gob/decode.go:175
		_go_fuzz_dep_.CoverTab[84089]++
								return ^int64(x >> 1)
//line /usr/local/go/src/encoding/gob/decode.go:176
		// _ = "end of CoverTab[84089]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:177
		_go_fuzz_dep_.CoverTab[84090]++
//line /usr/local/go/src/encoding/gob/decode.go:177
		// _ = "end of CoverTab[84090]"
//line /usr/local/go/src/encoding/gob/decode.go:177
	}
//line /usr/local/go/src/encoding/gob/decode.go:177
	// _ = "end of CoverTab[84087]"
//line /usr/local/go/src/encoding/gob/decode.go:177
	_go_fuzz_dep_.CoverTab[84088]++
							return int64(x >> 1)
//line /usr/local/go/src/encoding/gob/decode.go:178
	// _ = "end of CoverTab[84088]"
}

//line /usr/local/go/src/encoding/gob/decode.go:184
func (state *decoderState) getLength() (int, bool) {
//line /usr/local/go/src/encoding/gob/decode.go:184
	_go_fuzz_dep_.CoverTab[84091]++
							n := int(state.decodeUint())
							if n < 0 || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:186
		_go_fuzz_dep_.CoverTab[84093]++
//line /usr/local/go/src/encoding/gob/decode.go:186
		return state.b.Len() < n
//line /usr/local/go/src/encoding/gob/decode.go:186
		// _ = "end of CoverTab[84093]"
//line /usr/local/go/src/encoding/gob/decode.go:186
	}() || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:186
		_go_fuzz_dep_.CoverTab[84094]++
//line /usr/local/go/src/encoding/gob/decode.go:186
		return tooBig <= n
//line /usr/local/go/src/encoding/gob/decode.go:186
		// _ = "end of CoverTab[84094]"
//line /usr/local/go/src/encoding/gob/decode.go:186
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:186
		_go_fuzz_dep_.CoverTab[84095]++
								return 0, false
//line /usr/local/go/src/encoding/gob/decode.go:187
		// _ = "end of CoverTab[84095]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:188
		_go_fuzz_dep_.CoverTab[84096]++
//line /usr/local/go/src/encoding/gob/decode.go:188
		// _ = "end of CoverTab[84096]"
//line /usr/local/go/src/encoding/gob/decode.go:188
	}
//line /usr/local/go/src/encoding/gob/decode.go:188
	// _ = "end of CoverTab[84091]"
//line /usr/local/go/src/encoding/gob/decode.go:188
	_go_fuzz_dep_.CoverTab[84092]++
							return n, true
//line /usr/local/go/src/encoding/gob/decode.go:189
	// _ = "end of CoverTab[84092]"
}

//line /usr/local/go/src/encoding/gob/decode.go:193
type decOp func(i *decInstr, state *decoderState, v reflect.Value)

//line /usr/local/go/src/encoding/gob/decode.go:196
type decInstr struct {
	op	decOp
	field	int
	index	[]int
	ovfl	error
}

//line /usr/local/go/src/encoding/gob/decode.go:204
func ignoreUint(i *decInstr, state *decoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:204
	_go_fuzz_dep_.CoverTab[84097]++
							state.decodeUint()
//line /usr/local/go/src/encoding/gob/decode.go:205
	// _ = "end of CoverTab[84097]"
}

//line /usr/local/go/src/encoding/gob/decode.go:210
func ignoreTwoUints(i *decInstr, state *decoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:210
	_go_fuzz_dep_.CoverTab[84098]++
							state.decodeUint()
							state.decodeUint()
//line /usr/local/go/src/encoding/gob/decode.go:212
	// _ = "end of CoverTab[84098]"
}

//line /usr/local/go/src/encoding/gob/decode.go:226
func decAlloc(v reflect.Value) reflect.Value {
//line /usr/local/go/src/encoding/gob/decode.go:226
	_go_fuzz_dep_.CoverTab[84099]++
							for v.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/gob/decode.go:227
		_go_fuzz_dep_.CoverTab[84101]++
								if v.IsNil() {
//line /usr/local/go/src/encoding/gob/decode.go:228
			_go_fuzz_dep_.CoverTab[84103]++
									v.Set(reflect.New(v.Type().Elem()))
//line /usr/local/go/src/encoding/gob/decode.go:229
			// _ = "end of CoverTab[84103]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:230
			_go_fuzz_dep_.CoverTab[84104]++
//line /usr/local/go/src/encoding/gob/decode.go:230
			// _ = "end of CoverTab[84104]"
//line /usr/local/go/src/encoding/gob/decode.go:230
		}
//line /usr/local/go/src/encoding/gob/decode.go:230
		// _ = "end of CoverTab[84101]"
//line /usr/local/go/src/encoding/gob/decode.go:230
		_go_fuzz_dep_.CoverTab[84102]++
								v = v.Elem()
//line /usr/local/go/src/encoding/gob/decode.go:231
		// _ = "end of CoverTab[84102]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:232
	// _ = "end of CoverTab[84099]"
//line /usr/local/go/src/encoding/gob/decode.go:232
	_go_fuzz_dep_.CoverTab[84100]++
							return v
//line /usr/local/go/src/encoding/gob/decode.go:233
	// _ = "end of CoverTab[84100]"
}

//line /usr/local/go/src/encoding/gob/decode.go:237
func decBool(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:237
	_go_fuzz_dep_.CoverTab[84105]++
							value.SetBool(state.decodeUint() != 0)
//line /usr/local/go/src/encoding/gob/decode.go:238
	// _ = "end of CoverTab[84105]"
}

//line /usr/local/go/src/encoding/gob/decode.go:242
func decInt8(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:242
	_go_fuzz_dep_.CoverTab[84106]++
							v := state.decodeInt()
							if v < math.MinInt8 || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:244
		_go_fuzz_dep_.CoverTab[84108]++
//line /usr/local/go/src/encoding/gob/decode.go:244
		return math.MaxInt8 < v
//line /usr/local/go/src/encoding/gob/decode.go:244
		// _ = "end of CoverTab[84108]"
//line /usr/local/go/src/encoding/gob/decode.go:244
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:244
		_go_fuzz_dep_.CoverTab[84109]++
								error_(i.ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:245
		// _ = "end of CoverTab[84109]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:246
		_go_fuzz_dep_.CoverTab[84110]++
//line /usr/local/go/src/encoding/gob/decode.go:246
		// _ = "end of CoverTab[84110]"
//line /usr/local/go/src/encoding/gob/decode.go:246
	}
//line /usr/local/go/src/encoding/gob/decode.go:246
	// _ = "end of CoverTab[84106]"
//line /usr/local/go/src/encoding/gob/decode.go:246
	_go_fuzz_dep_.CoverTab[84107]++
							value.SetInt(v)
//line /usr/local/go/src/encoding/gob/decode.go:247
	// _ = "end of CoverTab[84107]"
}

//line /usr/local/go/src/encoding/gob/decode.go:251
func decUint8(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:251
	_go_fuzz_dep_.CoverTab[84111]++
							v := state.decodeUint()
							if math.MaxUint8 < v {
//line /usr/local/go/src/encoding/gob/decode.go:253
		_go_fuzz_dep_.CoverTab[84113]++
								error_(i.ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:254
		// _ = "end of CoverTab[84113]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:255
		_go_fuzz_dep_.CoverTab[84114]++
//line /usr/local/go/src/encoding/gob/decode.go:255
		// _ = "end of CoverTab[84114]"
//line /usr/local/go/src/encoding/gob/decode.go:255
	}
//line /usr/local/go/src/encoding/gob/decode.go:255
	// _ = "end of CoverTab[84111]"
//line /usr/local/go/src/encoding/gob/decode.go:255
	_go_fuzz_dep_.CoverTab[84112]++
							value.SetUint(v)
//line /usr/local/go/src/encoding/gob/decode.go:256
	// _ = "end of CoverTab[84112]"
}

//line /usr/local/go/src/encoding/gob/decode.go:260
func decInt16(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:260
	_go_fuzz_dep_.CoverTab[84115]++
							v := state.decodeInt()
							if v < math.MinInt16 || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:262
		_go_fuzz_dep_.CoverTab[84117]++
//line /usr/local/go/src/encoding/gob/decode.go:262
		return math.MaxInt16 < v
//line /usr/local/go/src/encoding/gob/decode.go:262
		// _ = "end of CoverTab[84117]"
//line /usr/local/go/src/encoding/gob/decode.go:262
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:262
		_go_fuzz_dep_.CoverTab[84118]++
								error_(i.ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:263
		// _ = "end of CoverTab[84118]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:264
		_go_fuzz_dep_.CoverTab[84119]++
//line /usr/local/go/src/encoding/gob/decode.go:264
		// _ = "end of CoverTab[84119]"
//line /usr/local/go/src/encoding/gob/decode.go:264
	}
//line /usr/local/go/src/encoding/gob/decode.go:264
	// _ = "end of CoverTab[84115]"
//line /usr/local/go/src/encoding/gob/decode.go:264
	_go_fuzz_dep_.CoverTab[84116]++
							value.SetInt(v)
//line /usr/local/go/src/encoding/gob/decode.go:265
	// _ = "end of CoverTab[84116]"
}

//line /usr/local/go/src/encoding/gob/decode.go:269
func decUint16(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:269
	_go_fuzz_dep_.CoverTab[84120]++
							v := state.decodeUint()
							if math.MaxUint16 < v {
//line /usr/local/go/src/encoding/gob/decode.go:271
		_go_fuzz_dep_.CoverTab[84122]++
								error_(i.ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:272
		// _ = "end of CoverTab[84122]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:273
		_go_fuzz_dep_.CoverTab[84123]++
//line /usr/local/go/src/encoding/gob/decode.go:273
		// _ = "end of CoverTab[84123]"
//line /usr/local/go/src/encoding/gob/decode.go:273
	}
//line /usr/local/go/src/encoding/gob/decode.go:273
	// _ = "end of CoverTab[84120]"
//line /usr/local/go/src/encoding/gob/decode.go:273
	_go_fuzz_dep_.CoverTab[84121]++
							value.SetUint(v)
//line /usr/local/go/src/encoding/gob/decode.go:274
	// _ = "end of CoverTab[84121]"
}

//line /usr/local/go/src/encoding/gob/decode.go:278
func decInt32(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:278
	_go_fuzz_dep_.CoverTab[84124]++
							v := state.decodeInt()
							if v < math.MinInt32 || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:280
		_go_fuzz_dep_.CoverTab[84126]++
//line /usr/local/go/src/encoding/gob/decode.go:280
		return math.MaxInt32 < v
//line /usr/local/go/src/encoding/gob/decode.go:280
		// _ = "end of CoverTab[84126]"
//line /usr/local/go/src/encoding/gob/decode.go:280
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:280
		_go_fuzz_dep_.CoverTab[84127]++
								error_(i.ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:281
		// _ = "end of CoverTab[84127]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:282
		_go_fuzz_dep_.CoverTab[84128]++
//line /usr/local/go/src/encoding/gob/decode.go:282
		// _ = "end of CoverTab[84128]"
//line /usr/local/go/src/encoding/gob/decode.go:282
	}
//line /usr/local/go/src/encoding/gob/decode.go:282
	// _ = "end of CoverTab[84124]"
//line /usr/local/go/src/encoding/gob/decode.go:282
	_go_fuzz_dep_.CoverTab[84125]++
							value.SetInt(v)
//line /usr/local/go/src/encoding/gob/decode.go:283
	// _ = "end of CoverTab[84125]"
}

//line /usr/local/go/src/encoding/gob/decode.go:287
func decUint32(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:287
	_go_fuzz_dep_.CoverTab[84129]++
							v := state.decodeUint()
							if math.MaxUint32 < v {
//line /usr/local/go/src/encoding/gob/decode.go:289
		_go_fuzz_dep_.CoverTab[84131]++
								error_(i.ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:290
		// _ = "end of CoverTab[84131]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:291
		_go_fuzz_dep_.CoverTab[84132]++
//line /usr/local/go/src/encoding/gob/decode.go:291
		// _ = "end of CoverTab[84132]"
//line /usr/local/go/src/encoding/gob/decode.go:291
	}
//line /usr/local/go/src/encoding/gob/decode.go:291
	// _ = "end of CoverTab[84129]"
//line /usr/local/go/src/encoding/gob/decode.go:291
	_go_fuzz_dep_.CoverTab[84130]++
							value.SetUint(v)
//line /usr/local/go/src/encoding/gob/decode.go:292
	// _ = "end of CoverTab[84130]"
}

//line /usr/local/go/src/encoding/gob/decode.go:296
func decInt64(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:296
	_go_fuzz_dep_.CoverTab[84133]++
							v := state.decodeInt()
							value.SetInt(v)
//line /usr/local/go/src/encoding/gob/decode.go:298
	// _ = "end of CoverTab[84133]"
}

//line /usr/local/go/src/encoding/gob/decode.go:302
func decUint64(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:302
	_go_fuzz_dep_.CoverTab[84134]++
							v := state.decodeUint()
							value.SetUint(v)
//line /usr/local/go/src/encoding/gob/decode.go:304
	// _ = "end of CoverTab[84134]"
}

//line /usr/local/go/src/encoding/gob/decode.go:312
func float64FromBits(u uint64) float64 {
//line /usr/local/go/src/encoding/gob/decode.go:312
	_go_fuzz_dep_.CoverTab[84135]++
							v := bits.ReverseBytes64(u)
							return math.Float64frombits(v)
//line /usr/local/go/src/encoding/gob/decode.go:314
	// _ = "end of CoverTab[84135]"
}

//line /usr/local/go/src/encoding/gob/decode.go:321
func float32FromBits(u uint64, ovfl error) float64 {
//line /usr/local/go/src/encoding/gob/decode.go:321
	_go_fuzz_dep_.CoverTab[84136]++
							v := float64FromBits(u)
							av := v
							if av < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:324
		_go_fuzz_dep_.CoverTab[84139]++
								av = -av
//line /usr/local/go/src/encoding/gob/decode.go:325
		// _ = "end of CoverTab[84139]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:326
		_go_fuzz_dep_.CoverTab[84140]++
//line /usr/local/go/src/encoding/gob/decode.go:326
		// _ = "end of CoverTab[84140]"
//line /usr/local/go/src/encoding/gob/decode.go:326
	}
//line /usr/local/go/src/encoding/gob/decode.go:326
	// _ = "end of CoverTab[84136]"
//line /usr/local/go/src/encoding/gob/decode.go:326
	_go_fuzz_dep_.CoverTab[84137]++

							if math.MaxFloat32 < av && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:328
		_go_fuzz_dep_.CoverTab[84141]++
//line /usr/local/go/src/encoding/gob/decode.go:328
		return av <= math.MaxFloat64
//line /usr/local/go/src/encoding/gob/decode.go:328
		// _ = "end of CoverTab[84141]"
//line /usr/local/go/src/encoding/gob/decode.go:328
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:328
		_go_fuzz_dep_.CoverTab[84142]++
								error_(ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:329
		// _ = "end of CoverTab[84142]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:330
		_go_fuzz_dep_.CoverTab[84143]++
//line /usr/local/go/src/encoding/gob/decode.go:330
		// _ = "end of CoverTab[84143]"
//line /usr/local/go/src/encoding/gob/decode.go:330
	}
//line /usr/local/go/src/encoding/gob/decode.go:330
	// _ = "end of CoverTab[84137]"
//line /usr/local/go/src/encoding/gob/decode.go:330
	_go_fuzz_dep_.CoverTab[84138]++
							return v
//line /usr/local/go/src/encoding/gob/decode.go:331
	// _ = "end of CoverTab[84138]"
}

//line /usr/local/go/src/encoding/gob/decode.go:336
func decFloat32(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:336
	_go_fuzz_dep_.CoverTab[84144]++
							value.SetFloat(float32FromBits(state.decodeUint(), i.ovfl))
//line /usr/local/go/src/encoding/gob/decode.go:337
	// _ = "end of CoverTab[84144]"
}

//line /usr/local/go/src/encoding/gob/decode.go:342
func decFloat64(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:342
	_go_fuzz_dep_.CoverTab[84145]++
							value.SetFloat(float64FromBits(state.decodeUint()))
//line /usr/local/go/src/encoding/gob/decode.go:343
	// _ = "end of CoverTab[84145]"
}

//line /usr/local/go/src/encoding/gob/decode.go:349
func decComplex64(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:349
	_go_fuzz_dep_.CoverTab[84146]++
							real := float32FromBits(state.decodeUint(), i.ovfl)
							imag := float32FromBits(state.decodeUint(), i.ovfl)
							value.SetComplex(complex(real, imag))
//line /usr/local/go/src/encoding/gob/decode.go:352
	// _ = "end of CoverTab[84146]"
}

//line /usr/local/go/src/encoding/gob/decode.go:358
func decComplex128(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:358
	_go_fuzz_dep_.CoverTab[84147]++
							real := float64FromBits(state.decodeUint())
							imag := float64FromBits(state.decodeUint())
							value.SetComplex(complex(real, imag))
//line /usr/local/go/src/encoding/gob/decode.go:361
	// _ = "end of CoverTab[84147]"
}

//line /usr/local/go/src/encoding/gob/decode.go:367
func decUint8Slice(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:367
	_go_fuzz_dep_.CoverTab[84148]++
							n, ok := state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:369
		_go_fuzz_dep_.CoverTab[84150]++
								errorf("bad %s slice length: %d", value.Type(), n)
//line /usr/local/go/src/encoding/gob/decode.go:370
		// _ = "end of CoverTab[84150]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:371
		_go_fuzz_dep_.CoverTab[84151]++
//line /usr/local/go/src/encoding/gob/decode.go:371
		// _ = "end of CoverTab[84151]"
//line /usr/local/go/src/encoding/gob/decode.go:371
	}
//line /usr/local/go/src/encoding/gob/decode.go:371
	// _ = "end of CoverTab[84148]"
//line /usr/local/go/src/encoding/gob/decode.go:371
	_go_fuzz_dep_.CoverTab[84149]++
							if value.Cap() < n {
//line /usr/local/go/src/encoding/gob/decode.go:372
		_go_fuzz_dep_.CoverTab[84152]++
								safe := saferio.SliceCap((*byte)(nil), uint64(n))
								if safe < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:374
			_go_fuzz_dep_.CoverTab[84154]++
									errorf("%s slice too big: %d elements", value.Type(), n)
//line /usr/local/go/src/encoding/gob/decode.go:375
			// _ = "end of CoverTab[84154]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:376
			_go_fuzz_dep_.CoverTab[84155]++
//line /usr/local/go/src/encoding/gob/decode.go:376
			// _ = "end of CoverTab[84155]"
//line /usr/local/go/src/encoding/gob/decode.go:376
		}
//line /usr/local/go/src/encoding/gob/decode.go:376
		// _ = "end of CoverTab[84152]"
//line /usr/local/go/src/encoding/gob/decode.go:376
		_go_fuzz_dep_.CoverTab[84153]++
								value.Set(reflect.MakeSlice(value.Type(), safe, safe))
								ln := safe
								i := 0
								for i < n {
//line /usr/local/go/src/encoding/gob/decode.go:380
			_go_fuzz_dep_.CoverTab[84156]++
									if i >= ln {
//line /usr/local/go/src/encoding/gob/decode.go:381
				_go_fuzz_dep_.CoverTab[84160]++

//line /usr/local/go/src/encoding/gob/decode.go:387
				value.Set(reflect.Append(value, reflect.Zero(value.Type().Elem())))
//line /usr/local/go/src/encoding/gob/decode.go:387
				// _ = "end of CoverTab[84160]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:388
				_go_fuzz_dep_.CoverTab[84161]++
//line /usr/local/go/src/encoding/gob/decode.go:388
				// _ = "end of CoverTab[84161]"
//line /usr/local/go/src/encoding/gob/decode.go:388
			}
//line /usr/local/go/src/encoding/gob/decode.go:388
			// _ = "end of CoverTab[84156]"
//line /usr/local/go/src/encoding/gob/decode.go:388
			_go_fuzz_dep_.CoverTab[84157]++

//line /usr/local/go/src/encoding/gob/decode.go:391
			ln = value.Cap()
			if ln > n {
//line /usr/local/go/src/encoding/gob/decode.go:392
				_go_fuzz_dep_.CoverTab[84162]++
										ln = n
//line /usr/local/go/src/encoding/gob/decode.go:393
				// _ = "end of CoverTab[84162]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:394
				_go_fuzz_dep_.CoverTab[84163]++
//line /usr/local/go/src/encoding/gob/decode.go:394
				// _ = "end of CoverTab[84163]"
//line /usr/local/go/src/encoding/gob/decode.go:394
			}
//line /usr/local/go/src/encoding/gob/decode.go:394
			// _ = "end of CoverTab[84157]"
//line /usr/local/go/src/encoding/gob/decode.go:394
			_go_fuzz_dep_.CoverTab[84158]++
									value.SetLen(ln)
									sub := value.Slice(i, ln)
									if _, err := state.b.Read(sub.Bytes()); err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:397
				_go_fuzz_dep_.CoverTab[84164]++
										errorf("error decoding []byte at %d: %s", err, i)
//line /usr/local/go/src/encoding/gob/decode.go:398
				// _ = "end of CoverTab[84164]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:399
				_go_fuzz_dep_.CoverTab[84165]++
//line /usr/local/go/src/encoding/gob/decode.go:399
				// _ = "end of CoverTab[84165]"
//line /usr/local/go/src/encoding/gob/decode.go:399
			}
//line /usr/local/go/src/encoding/gob/decode.go:399
			// _ = "end of CoverTab[84158]"
//line /usr/local/go/src/encoding/gob/decode.go:399
			_go_fuzz_dep_.CoverTab[84159]++
									i = ln
//line /usr/local/go/src/encoding/gob/decode.go:400
			// _ = "end of CoverTab[84159]"
		}
//line /usr/local/go/src/encoding/gob/decode.go:401
		// _ = "end of CoverTab[84153]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:402
		_go_fuzz_dep_.CoverTab[84166]++
								value.SetLen(n)
								if _, err := state.b.Read(value.Bytes()); err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:404
			_go_fuzz_dep_.CoverTab[84167]++
									errorf("error decoding []byte: %s", err)
//line /usr/local/go/src/encoding/gob/decode.go:405
			// _ = "end of CoverTab[84167]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:406
			_go_fuzz_dep_.CoverTab[84168]++
//line /usr/local/go/src/encoding/gob/decode.go:406
			// _ = "end of CoverTab[84168]"
//line /usr/local/go/src/encoding/gob/decode.go:406
		}
//line /usr/local/go/src/encoding/gob/decode.go:406
		// _ = "end of CoverTab[84166]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:407
	// _ = "end of CoverTab[84149]"
}

//line /usr/local/go/src/encoding/gob/decode.go:413
func decString(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:413
	_go_fuzz_dep_.CoverTab[84169]++
							n, ok := state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:415
		_go_fuzz_dep_.CoverTab[84172]++
								errorf("bad %s slice length: %d", value.Type(), n)
//line /usr/local/go/src/encoding/gob/decode.go:416
		// _ = "end of CoverTab[84172]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:417
		_go_fuzz_dep_.CoverTab[84173]++
//line /usr/local/go/src/encoding/gob/decode.go:417
		// _ = "end of CoverTab[84173]"
//line /usr/local/go/src/encoding/gob/decode.go:417
	}
//line /usr/local/go/src/encoding/gob/decode.go:417
	// _ = "end of CoverTab[84169]"
//line /usr/local/go/src/encoding/gob/decode.go:417
	_go_fuzz_dep_.CoverTab[84170]++

							data := state.b.Bytes()
							if len(data) < n {
//line /usr/local/go/src/encoding/gob/decode.go:420
		_go_fuzz_dep_.CoverTab[84174]++
								errorf("invalid string length %d: exceeds input size %d", n, len(data))
//line /usr/local/go/src/encoding/gob/decode.go:421
		// _ = "end of CoverTab[84174]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:422
		_go_fuzz_dep_.CoverTab[84175]++
//line /usr/local/go/src/encoding/gob/decode.go:422
		// _ = "end of CoverTab[84175]"
//line /usr/local/go/src/encoding/gob/decode.go:422
	}
//line /usr/local/go/src/encoding/gob/decode.go:422
	// _ = "end of CoverTab[84170]"
//line /usr/local/go/src/encoding/gob/decode.go:422
	_go_fuzz_dep_.CoverTab[84171]++
							s := string(data[:n])
							state.b.Drop(n)
							value.SetString(s)
//line /usr/local/go/src/encoding/gob/decode.go:425
	// _ = "end of CoverTab[84171]"
}

//line /usr/local/go/src/encoding/gob/decode.go:429
func ignoreUint8Array(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:429
	_go_fuzz_dep_.CoverTab[84176]++
							n, ok := state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:431
		_go_fuzz_dep_.CoverTab[84179]++
								errorf("slice length too large")
//line /usr/local/go/src/encoding/gob/decode.go:432
		// _ = "end of CoverTab[84179]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:433
		_go_fuzz_dep_.CoverTab[84180]++
//line /usr/local/go/src/encoding/gob/decode.go:433
		// _ = "end of CoverTab[84180]"
//line /usr/local/go/src/encoding/gob/decode.go:433
	}
//line /usr/local/go/src/encoding/gob/decode.go:433
	// _ = "end of CoverTab[84176]"
//line /usr/local/go/src/encoding/gob/decode.go:433
	_go_fuzz_dep_.CoverTab[84177]++
							bn := state.b.Len()
							if bn < n {
//line /usr/local/go/src/encoding/gob/decode.go:435
		_go_fuzz_dep_.CoverTab[84181]++
								errorf("invalid slice length %d: exceeds input size %d", n, bn)
//line /usr/local/go/src/encoding/gob/decode.go:436
		// _ = "end of CoverTab[84181]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:437
		_go_fuzz_dep_.CoverTab[84182]++
//line /usr/local/go/src/encoding/gob/decode.go:437
		// _ = "end of CoverTab[84182]"
//line /usr/local/go/src/encoding/gob/decode.go:437
	}
//line /usr/local/go/src/encoding/gob/decode.go:437
	// _ = "end of CoverTab[84177]"
//line /usr/local/go/src/encoding/gob/decode.go:437
	_go_fuzz_dep_.CoverTab[84178]++
							state.b.Drop(n)
//line /usr/local/go/src/encoding/gob/decode.go:438
	// _ = "end of CoverTab[84178]"
}

//line /usr/local/go/src/encoding/gob/decode.go:445
type decEngine struct {
	instr		[]decInstr
	numInstr	int
}

//line /usr/local/go/src/encoding/gob/decode.go:453
func (dec *Decoder) decodeSingle(engine *decEngine, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:453
	_go_fuzz_dep_.CoverTab[84183]++
							state := dec.newDecoderState(&dec.buf)
							defer dec.freeDecoderState(state)
							state.fieldnum = singletonField
							if state.decodeUint() != 0 {
//line /usr/local/go/src/encoding/gob/decode.go:457
		_go_fuzz_dep_.CoverTab[84185]++
								errorf("decode: corrupted data: non-zero delta for singleton")
//line /usr/local/go/src/encoding/gob/decode.go:458
		// _ = "end of CoverTab[84185]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:459
		_go_fuzz_dep_.CoverTab[84186]++
//line /usr/local/go/src/encoding/gob/decode.go:459
		// _ = "end of CoverTab[84186]"
//line /usr/local/go/src/encoding/gob/decode.go:459
	}
//line /usr/local/go/src/encoding/gob/decode.go:459
	// _ = "end of CoverTab[84183]"
//line /usr/local/go/src/encoding/gob/decode.go:459
	_go_fuzz_dep_.CoverTab[84184]++
							instr := &engine.instr[singletonField]
							instr.op(instr, state, value)
//line /usr/local/go/src/encoding/gob/decode.go:461
	// _ = "end of CoverTab[84184]"
}

//line /usr/local/go/src/encoding/gob/decode.go:469
func (dec *Decoder) decodeStruct(engine *decEngine, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:469
	_go_fuzz_dep_.CoverTab[84187]++
							state := dec.newDecoderState(&dec.buf)
							defer dec.freeDecoderState(state)
							state.fieldnum = -1
							for state.b.Len() > 0 {
//line /usr/local/go/src/encoding/gob/decode.go:473
		_go_fuzz_dep_.CoverTab[84188]++
								delta := int(state.decodeUint())
								if delta < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:475
			_go_fuzz_dep_.CoverTab[84193]++
									errorf("decode: corrupted data: negative delta")
//line /usr/local/go/src/encoding/gob/decode.go:476
			// _ = "end of CoverTab[84193]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:477
			_go_fuzz_dep_.CoverTab[84194]++
//line /usr/local/go/src/encoding/gob/decode.go:477
			// _ = "end of CoverTab[84194]"
//line /usr/local/go/src/encoding/gob/decode.go:477
		}
//line /usr/local/go/src/encoding/gob/decode.go:477
		// _ = "end of CoverTab[84188]"
//line /usr/local/go/src/encoding/gob/decode.go:477
		_go_fuzz_dep_.CoverTab[84189]++
								if delta == 0 {
//line /usr/local/go/src/encoding/gob/decode.go:478
			_go_fuzz_dep_.CoverTab[84195]++
									break
//line /usr/local/go/src/encoding/gob/decode.go:479
			// _ = "end of CoverTab[84195]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:480
			_go_fuzz_dep_.CoverTab[84196]++
//line /usr/local/go/src/encoding/gob/decode.go:480
			// _ = "end of CoverTab[84196]"
//line /usr/local/go/src/encoding/gob/decode.go:480
		}
//line /usr/local/go/src/encoding/gob/decode.go:480
		// _ = "end of CoverTab[84189]"
//line /usr/local/go/src/encoding/gob/decode.go:480
		_go_fuzz_dep_.CoverTab[84190]++
								if state.fieldnum >= len(engine.instr)-delta {
//line /usr/local/go/src/encoding/gob/decode.go:481
			_go_fuzz_dep_.CoverTab[84197]++
									error_(errRange)
//line /usr/local/go/src/encoding/gob/decode.go:482
			// _ = "end of CoverTab[84197]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:483
			_go_fuzz_dep_.CoverTab[84198]++
//line /usr/local/go/src/encoding/gob/decode.go:483
			// _ = "end of CoverTab[84198]"
//line /usr/local/go/src/encoding/gob/decode.go:483
		}
//line /usr/local/go/src/encoding/gob/decode.go:483
		// _ = "end of CoverTab[84190]"
//line /usr/local/go/src/encoding/gob/decode.go:483
		_go_fuzz_dep_.CoverTab[84191]++
								fieldnum := state.fieldnum + delta
								instr := &engine.instr[fieldnum]
								var field reflect.Value
								if instr.index != nil {
//line /usr/local/go/src/encoding/gob/decode.go:487
			_go_fuzz_dep_.CoverTab[84199]++

									field = value.FieldByIndex(instr.index)
									if field.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/gob/decode.go:490
				_go_fuzz_dep_.CoverTab[84200]++
										field = decAlloc(field)
//line /usr/local/go/src/encoding/gob/decode.go:491
				// _ = "end of CoverTab[84200]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:492
				_go_fuzz_dep_.CoverTab[84201]++
//line /usr/local/go/src/encoding/gob/decode.go:492
				// _ = "end of CoverTab[84201]"
//line /usr/local/go/src/encoding/gob/decode.go:492
			}
//line /usr/local/go/src/encoding/gob/decode.go:492
			// _ = "end of CoverTab[84199]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:493
			_go_fuzz_dep_.CoverTab[84202]++
//line /usr/local/go/src/encoding/gob/decode.go:493
			// _ = "end of CoverTab[84202]"
//line /usr/local/go/src/encoding/gob/decode.go:493
		}
//line /usr/local/go/src/encoding/gob/decode.go:493
		// _ = "end of CoverTab[84191]"
//line /usr/local/go/src/encoding/gob/decode.go:493
		_go_fuzz_dep_.CoverTab[84192]++
								instr.op(instr, state, field)
								state.fieldnum = fieldnum
//line /usr/local/go/src/encoding/gob/decode.go:495
		// _ = "end of CoverTab[84192]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:496
	// _ = "end of CoverTab[84187]"
}

var noValue reflect.Value

//line /usr/local/go/src/encoding/gob/decode.go:502
func (dec *Decoder) ignoreStruct(engine *decEngine) {
//line /usr/local/go/src/encoding/gob/decode.go:502
	_go_fuzz_dep_.CoverTab[84203]++
							state := dec.newDecoderState(&dec.buf)
							defer dec.freeDecoderState(state)
							state.fieldnum = -1
							for state.b.Len() > 0 {
//line /usr/local/go/src/encoding/gob/decode.go:506
		_go_fuzz_dep_.CoverTab[84204]++
								delta := int(state.decodeUint())
								if delta < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:508
			_go_fuzz_dep_.CoverTab[84208]++
									errorf("ignore decode: corrupted data: negative delta")
//line /usr/local/go/src/encoding/gob/decode.go:509
			// _ = "end of CoverTab[84208]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:510
			_go_fuzz_dep_.CoverTab[84209]++
//line /usr/local/go/src/encoding/gob/decode.go:510
			// _ = "end of CoverTab[84209]"
//line /usr/local/go/src/encoding/gob/decode.go:510
		}
//line /usr/local/go/src/encoding/gob/decode.go:510
		// _ = "end of CoverTab[84204]"
//line /usr/local/go/src/encoding/gob/decode.go:510
		_go_fuzz_dep_.CoverTab[84205]++
								if delta == 0 {
//line /usr/local/go/src/encoding/gob/decode.go:511
			_go_fuzz_dep_.CoverTab[84210]++
									break
//line /usr/local/go/src/encoding/gob/decode.go:512
			// _ = "end of CoverTab[84210]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:513
			_go_fuzz_dep_.CoverTab[84211]++
//line /usr/local/go/src/encoding/gob/decode.go:513
			// _ = "end of CoverTab[84211]"
//line /usr/local/go/src/encoding/gob/decode.go:513
		}
//line /usr/local/go/src/encoding/gob/decode.go:513
		// _ = "end of CoverTab[84205]"
//line /usr/local/go/src/encoding/gob/decode.go:513
		_go_fuzz_dep_.CoverTab[84206]++
								fieldnum := state.fieldnum + delta
								if fieldnum >= len(engine.instr) {
//line /usr/local/go/src/encoding/gob/decode.go:515
			_go_fuzz_dep_.CoverTab[84212]++
									error_(errRange)
//line /usr/local/go/src/encoding/gob/decode.go:516
			// _ = "end of CoverTab[84212]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:517
			_go_fuzz_dep_.CoverTab[84213]++
//line /usr/local/go/src/encoding/gob/decode.go:517
			// _ = "end of CoverTab[84213]"
//line /usr/local/go/src/encoding/gob/decode.go:517
		}
//line /usr/local/go/src/encoding/gob/decode.go:517
		// _ = "end of CoverTab[84206]"
//line /usr/local/go/src/encoding/gob/decode.go:517
		_go_fuzz_dep_.CoverTab[84207]++
								instr := &engine.instr[fieldnum]
								instr.op(instr, state, noValue)
								state.fieldnum = fieldnum
//line /usr/local/go/src/encoding/gob/decode.go:520
		// _ = "end of CoverTab[84207]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:521
	// _ = "end of CoverTab[84203]"
}

//line /usr/local/go/src/encoding/gob/decode.go:526
func (dec *Decoder) ignoreSingle(engine *decEngine) {
//line /usr/local/go/src/encoding/gob/decode.go:526
	_go_fuzz_dep_.CoverTab[84214]++
							state := dec.newDecoderState(&dec.buf)
							defer dec.freeDecoderState(state)
							state.fieldnum = singletonField
							delta := int(state.decodeUint())
							if delta != 0 {
//line /usr/local/go/src/encoding/gob/decode.go:531
		_go_fuzz_dep_.CoverTab[84216]++
								errorf("decode: corrupted data: non-zero delta for singleton")
//line /usr/local/go/src/encoding/gob/decode.go:532
		// _ = "end of CoverTab[84216]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:533
		_go_fuzz_dep_.CoverTab[84217]++
//line /usr/local/go/src/encoding/gob/decode.go:533
		// _ = "end of CoverTab[84217]"
//line /usr/local/go/src/encoding/gob/decode.go:533
	}
//line /usr/local/go/src/encoding/gob/decode.go:533
	// _ = "end of CoverTab[84214]"
//line /usr/local/go/src/encoding/gob/decode.go:533
	_go_fuzz_dep_.CoverTab[84215]++
							instr := &engine.instr[singletonField]
							instr.op(instr, state, noValue)
//line /usr/local/go/src/encoding/gob/decode.go:535
	// _ = "end of CoverTab[84215]"
}

//line /usr/local/go/src/encoding/gob/decode.go:539
func (dec *Decoder) decodeArrayHelper(state *decoderState, value reflect.Value, elemOp decOp, length int, ovfl error, helper decHelper) {
//line /usr/local/go/src/encoding/gob/decode.go:539
	_go_fuzz_dep_.CoverTab[84218]++
							if helper != nil && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:540
		_go_fuzz_dep_.CoverTab[84220]++
//line /usr/local/go/src/encoding/gob/decode.go:540
		return helper(state, value, length, ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:540
		// _ = "end of CoverTab[84220]"
//line /usr/local/go/src/encoding/gob/decode.go:540
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:540
		_go_fuzz_dep_.CoverTab[84221]++
								return
//line /usr/local/go/src/encoding/gob/decode.go:541
		// _ = "end of CoverTab[84221]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:542
		_go_fuzz_dep_.CoverTab[84222]++
//line /usr/local/go/src/encoding/gob/decode.go:542
		// _ = "end of CoverTab[84222]"
//line /usr/local/go/src/encoding/gob/decode.go:542
	}
//line /usr/local/go/src/encoding/gob/decode.go:542
	// _ = "end of CoverTab[84218]"
//line /usr/local/go/src/encoding/gob/decode.go:542
	_go_fuzz_dep_.CoverTab[84219]++
							instr := &decInstr{elemOp, 0, nil, ovfl}
							isPtr := value.Type().Elem().Kind() == reflect.Pointer
							ln := value.Len()
							for i := 0; i < length; i++ {
//line /usr/local/go/src/encoding/gob/decode.go:546
		_go_fuzz_dep_.CoverTab[84223]++
								if state.b.Len() == 0 {
//line /usr/local/go/src/encoding/gob/decode.go:547
			_go_fuzz_dep_.CoverTab[84227]++
									errorf("decoding array or slice: length exceeds input size (%d elements)", length)
//line /usr/local/go/src/encoding/gob/decode.go:548
			// _ = "end of CoverTab[84227]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:549
			_go_fuzz_dep_.CoverTab[84228]++
//line /usr/local/go/src/encoding/gob/decode.go:549
			// _ = "end of CoverTab[84228]"
//line /usr/local/go/src/encoding/gob/decode.go:549
		}
//line /usr/local/go/src/encoding/gob/decode.go:549
		// _ = "end of CoverTab[84223]"
//line /usr/local/go/src/encoding/gob/decode.go:549
		_go_fuzz_dep_.CoverTab[84224]++
								if i >= ln {
//line /usr/local/go/src/encoding/gob/decode.go:550
			_go_fuzz_dep_.CoverTab[84229]++

//line /usr/local/go/src/encoding/gob/decode.go:553
			value.Set(reflect.Append(value, reflect.Zero(value.Type().Elem())))
			cp := value.Cap()
			if cp > length {
//line /usr/local/go/src/encoding/gob/decode.go:555
				_go_fuzz_dep_.CoverTab[84231]++
										cp = length
//line /usr/local/go/src/encoding/gob/decode.go:556
				// _ = "end of CoverTab[84231]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:557
				_go_fuzz_dep_.CoverTab[84232]++
//line /usr/local/go/src/encoding/gob/decode.go:557
				// _ = "end of CoverTab[84232]"
//line /usr/local/go/src/encoding/gob/decode.go:557
			}
//line /usr/local/go/src/encoding/gob/decode.go:557
			// _ = "end of CoverTab[84229]"
//line /usr/local/go/src/encoding/gob/decode.go:557
			_go_fuzz_dep_.CoverTab[84230]++
									value.SetLen(cp)
									ln = cp
//line /usr/local/go/src/encoding/gob/decode.go:559
			// _ = "end of CoverTab[84230]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:560
			_go_fuzz_dep_.CoverTab[84233]++
//line /usr/local/go/src/encoding/gob/decode.go:560
			// _ = "end of CoverTab[84233]"
//line /usr/local/go/src/encoding/gob/decode.go:560
		}
//line /usr/local/go/src/encoding/gob/decode.go:560
		// _ = "end of CoverTab[84224]"
//line /usr/local/go/src/encoding/gob/decode.go:560
		_go_fuzz_dep_.CoverTab[84225]++
								v := value.Index(i)
								if isPtr {
//line /usr/local/go/src/encoding/gob/decode.go:562
			_go_fuzz_dep_.CoverTab[84234]++
									v = decAlloc(v)
//line /usr/local/go/src/encoding/gob/decode.go:563
			// _ = "end of CoverTab[84234]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:564
			_go_fuzz_dep_.CoverTab[84235]++
//line /usr/local/go/src/encoding/gob/decode.go:564
			// _ = "end of CoverTab[84235]"
//line /usr/local/go/src/encoding/gob/decode.go:564
		}
//line /usr/local/go/src/encoding/gob/decode.go:564
		// _ = "end of CoverTab[84225]"
//line /usr/local/go/src/encoding/gob/decode.go:564
		_go_fuzz_dep_.CoverTab[84226]++
								elemOp(instr, state, v)
//line /usr/local/go/src/encoding/gob/decode.go:565
		// _ = "end of CoverTab[84226]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:566
	// _ = "end of CoverTab[84219]"
}

//line /usr/local/go/src/encoding/gob/decode.go:572
func (dec *Decoder) decodeArray(state *decoderState, value reflect.Value, elemOp decOp, length int, ovfl error, helper decHelper) {
//line /usr/local/go/src/encoding/gob/decode.go:572
	_go_fuzz_dep_.CoverTab[84236]++
							if n := state.decodeUint(); n != uint64(length) {
//line /usr/local/go/src/encoding/gob/decode.go:573
		_go_fuzz_dep_.CoverTab[84238]++
								errorf("length mismatch in decodeArray")
//line /usr/local/go/src/encoding/gob/decode.go:574
		// _ = "end of CoverTab[84238]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:575
		_go_fuzz_dep_.CoverTab[84239]++
//line /usr/local/go/src/encoding/gob/decode.go:575
		// _ = "end of CoverTab[84239]"
//line /usr/local/go/src/encoding/gob/decode.go:575
	}
//line /usr/local/go/src/encoding/gob/decode.go:575
	// _ = "end of CoverTab[84236]"
//line /usr/local/go/src/encoding/gob/decode.go:575
	_go_fuzz_dep_.CoverTab[84237]++
							dec.decodeArrayHelper(state, value, elemOp, length, ovfl, helper)
//line /usr/local/go/src/encoding/gob/decode.go:576
	// _ = "end of CoverTab[84237]"
}

//line /usr/local/go/src/encoding/gob/decode.go:580
func decodeIntoValue(state *decoderState, op decOp, isPtr bool, value reflect.Value, instr *decInstr) reflect.Value {
//line /usr/local/go/src/encoding/gob/decode.go:580
	_go_fuzz_dep_.CoverTab[84240]++
							v := value
							if isPtr {
//line /usr/local/go/src/encoding/gob/decode.go:582
		_go_fuzz_dep_.CoverTab[84242]++
								v = decAlloc(value)
//line /usr/local/go/src/encoding/gob/decode.go:583
		// _ = "end of CoverTab[84242]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:584
		_go_fuzz_dep_.CoverTab[84243]++
//line /usr/local/go/src/encoding/gob/decode.go:584
		// _ = "end of CoverTab[84243]"
//line /usr/local/go/src/encoding/gob/decode.go:584
	}
//line /usr/local/go/src/encoding/gob/decode.go:584
	// _ = "end of CoverTab[84240]"
//line /usr/local/go/src/encoding/gob/decode.go:584
	_go_fuzz_dep_.CoverTab[84241]++

							op(instr, state, v)
							return value
//line /usr/local/go/src/encoding/gob/decode.go:587
	// _ = "end of CoverTab[84241]"
}

//line /usr/local/go/src/encoding/gob/decode.go:594
func (dec *Decoder) decodeMap(mtyp reflect.Type, state *decoderState, value reflect.Value, keyOp, elemOp decOp, ovfl error) {
//line /usr/local/go/src/encoding/gob/decode.go:594
	_go_fuzz_dep_.CoverTab[84244]++
							n := int(state.decodeUint())
							if value.IsNil() {
//line /usr/local/go/src/encoding/gob/decode.go:596
		_go_fuzz_dep_.CoverTab[84246]++
								value.Set(reflect.MakeMapWithSize(mtyp, n))
//line /usr/local/go/src/encoding/gob/decode.go:597
		// _ = "end of CoverTab[84246]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:598
		_go_fuzz_dep_.CoverTab[84247]++
//line /usr/local/go/src/encoding/gob/decode.go:598
		// _ = "end of CoverTab[84247]"
//line /usr/local/go/src/encoding/gob/decode.go:598
	}
//line /usr/local/go/src/encoding/gob/decode.go:598
	// _ = "end of CoverTab[84244]"
//line /usr/local/go/src/encoding/gob/decode.go:598
	_go_fuzz_dep_.CoverTab[84245]++
							keyIsPtr := mtyp.Key().Kind() == reflect.Pointer
							elemIsPtr := mtyp.Elem().Kind() == reflect.Pointer
							keyInstr := &decInstr{keyOp, 0, nil, ovfl}
							elemInstr := &decInstr{elemOp, 0, nil, ovfl}
							keyP := reflect.New(mtyp.Key())
							keyZ := reflect.Zero(mtyp.Key())
							elemP := reflect.New(mtyp.Elem())
							elemZ := reflect.Zero(mtyp.Elem())
							for i := 0; i < n; i++ {
//line /usr/local/go/src/encoding/gob/decode.go:607
		_go_fuzz_dep_.CoverTab[84248]++
								key := decodeIntoValue(state, keyOp, keyIsPtr, keyP.Elem(), keyInstr)
								elem := decodeIntoValue(state, elemOp, elemIsPtr, elemP.Elem(), elemInstr)
								value.SetMapIndex(key, elem)
								keyP.Elem().Set(keyZ)
								elemP.Elem().Set(elemZ)
//line /usr/local/go/src/encoding/gob/decode.go:612
		// _ = "end of CoverTab[84248]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:613
	// _ = "end of CoverTab[84245]"
}

//line /usr/local/go/src/encoding/gob/decode.go:617
func (dec *Decoder) ignoreArrayHelper(state *decoderState, elemOp decOp, length int) {
//line /usr/local/go/src/encoding/gob/decode.go:617
	_go_fuzz_dep_.CoverTab[84249]++
							instr := &decInstr{elemOp, 0, nil, errors.New("no error")}
							for i := 0; i < length; i++ {
//line /usr/local/go/src/encoding/gob/decode.go:619
		_go_fuzz_dep_.CoverTab[84250]++
								if state.b.Len() == 0 {
//line /usr/local/go/src/encoding/gob/decode.go:620
			_go_fuzz_dep_.CoverTab[84252]++
									errorf("decoding array or slice: length exceeds input size (%d elements)", length)
//line /usr/local/go/src/encoding/gob/decode.go:621
			// _ = "end of CoverTab[84252]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:622
			_go_fuzz_dep_.CoverTab[84253]++
//line /usr/local/go/src/encoding/gob/decode.go:622
			// _ = "end of CoverTab[84253]"
//line /usr/local/go/src/encoding/gob/decode.go:622
		}
//line /usr/local/go/src/encoding/gob/decode.go:622
		// _ = "end of CoverTab[84250]"
//line /usr/local/go/src/encoding/gob/decode.go:622
		_go_fuzz_dep_.CoverTab[84251]++
								elemOp(instr, state, noValue)
//line /usr/local/go/src/encoding/gob/decode.go:623
		// _ = "end of CoverTab[84251]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:624
	// _ = "end of CoverTab[84249]"
}

//line /usr/local/go/src/encoding/gob/decode.go:628
func (dec *Decoder) ignoreArray(state *decoderState, elemOp decOp, length int) {
//line /usr/local/go/src/encoding/gob/decode.go:628
	_go_fuzz_dep_.CoverTab[84254]++
							if n := state.decodeUint(); n != uint64(length) {
//line /usr/local/go/src/encoding/gob/decode.go:629
		_go_fuzz_dep_.CoverTab[84256]++
								errorf("length mismatch in ignoreArray")
//line /usr/local/go/src/encoding/gob/decode.go:630
		// _ = "end of CoverTab[84256]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:631
		_go_fuzz_dep_.CoverTab[84257]++
//line /usr/local/go/src/encoding/gob/decode.go:631
		// _ = "end of CoverTab[84257]"
//line /usr/local/go/src/encoding/gob/decode.go:631
	}
//line /usr/local/go/src/encoding/gob/decode.go:631
	// _ = "end of CoverTab[84254]"
//line /usr/local/go/src/encoding/gob/decode.go:631
	_go_fuzz_dep_.CoverTab[84255]++
							dec.ignoreArrayHelper(state, elemOp, length)
//line /usr/local/go/src/encoding/gob/decode.go:632
	// _ = "end of CoverTab[84255]"
}

//line /usr/local/go/src/encoding/gob/decode.go:636
func (dec *Decoder) ignoreMap(state *decoderState, keyOp, elemOp decOp) {
//line /usr/local/go/src/encoding/gob/decode.go:636
	_go_fuzz_dep_.CoverTab[84258]++
							n := int(state.decodeUint())
							keyInstr := &decInstr{keyOp, 0, nil, errors.New("no error")}
							elemInstr := &decInstr{elemOp, 0, nil, errors.New("no error")}
							for i := 0; i < n; i++ {
//line /usr/local/go/src/encoding/gob/decode.go:640
		_go_fuzz_dep_.CoverTab[84259]++
								keyOp(keyInstr, state, noValue)
								elemOp(elemInstr, state, noValue)
//line /usr/local/go/src/encoding/gob/decode.go:642
		// _ = "end of CoverTab[84259]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:643
	// _ = "end of CoverTab[84258]"
}

//line /usr/local/go/src/encoding/gob/decode.go:648
func (dec *Decoder) decodeSlice(state *decoderState, value reflect.Value, elemOp decOp, ovfl error, helper decHelper) {
//line /usr/local/go/src/encoding/gob/decode.go:648
	_go_fuzz_dep_.CoverTab[84260]++
							u := state.decodeUint()
							typ := value.Type()
							size := uint64(typ.Elem().Size())
							nBytes := u * size
							n := int(u)

							if n < 0 || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:655
		_go_fuzz_dep_.CoverTab[84263]++
//line /usr/local/go/src/encoding/gob/decode.go:655
		return uint64(n) != u
//line /usr/local/go/src/encoding/gob/decode.go:655
		// _ = "end of CoverTab[84263]"
//line /usr/local/go/src/encoding/gob/decode.go:655
	}() || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:655
		_go_fuzz_dep_.CoverTab[84264]++
//line /usr/local/go/src/encoding/gob/decode.go:655
		return nBytes > tooBig
//line /usr/local/go/src/encoding/gob/decode.go:655
		// _ = "end of CoverTab[84264]"
//line /usr/local/go/src/encoding/gob/decode.go:655
	}() || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:655
		_go_fuzz_dep_.CoverTab[84265]++
//line /usr/local/go/src/encoding/gob/decode.go:655
		return (size > 0 && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:655
			_go_fuzz_dep_.CoverTab[84266]++
//line /usr/local/go/src/encoding/gob/decode.go:655
			return nBytes/size != u
//line /usr/local/go/src/encoding/gob/decode.go:655
			// _ = "end of CoverTab[84266]"
//line /usr/local/go/src/encoding/gob/decode.go:655
		}())
//line /usr/local/go/src/encoding/gob/decode.go:655
		// _ = "end of CoverTab[84265]"
//line /usr/local/go/src/encoding/gob/decode.go:655
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:655
		_go_fuzz_dep_.CoverTab[84267]++

//line /usr/local/go/src/encoding/gob/decode.go:658
		errorf("%s slice too big: %d elements of %d bytes", typ.Elem(), u, size)
//line /usr/local/go/src/encoding/gob/decode.go:658
		// _ = "end of CoverTab[84267]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:659
		_go_fuzz_dep_.CoverTab[84268]++
//line /usr/local/go/src/encoding/gob/decode.go:659
		// _ = "end of CoverTab[84268]"
//line /usr/local/go/src/encoding/gob/decode.go:659
	}
//line /usr/local/go/src/encoding/gob/decode.go:659
	// _ = "end of CoverTab[84260]"
//line /usr/local/go/src/encoding/gob/decode.go:659
	_go_fuzz_dep_.CoverTab[84261]++
							if value.Cap() < n {
//line /usr/local/go/src/encoding/gob/decode.go:660
		_go_fuzz_dep_.CoverTab[84269]++
								safe := saferio.SliceCap(reflect.Zero(reflect.PtrTo(typ.Elem())).Interface(), uint64(n))
								if safe < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:662
			_go_fuzz_dep_.CoverTab[84271]++
									errorf("%s slice too big: %d elements of %d bytes", typ.Elem(), u, size)
//line /usr/local/go/src/encoding/gob/decode.go:663
			// _ = "end of CoverTab[84271]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:664
			_go_fuzz_dep_.CoverTab[84272]++
//line /usr/local/go/src/encoding/gob/decode.go:664
			// _ = "end of CoverTab[84272]"
//line /usr/local/go/src/encoding/gob/decode.go:664
		}
//line /usr/local/go/src/encoding/gob/decode.go:664
		// _ = "end of CoverTab[84269]"
//line /usr/local/go/src/encoding/gob/decode.go:664
		_go_fuzz_dep_.CoverTab[84270]++
								value.Set(reflect.MakeSlice(typ, safe, safe))
//line /usr/local/go/src/encoding/gob/decode.go:665
		// _ = "end of CoverTab[84270]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:666
		_go_fuzz_dep_.CoverTab[84273]++
								value.SetLen(n)
//line /usr/local/go/src/encoding/gob/decode.go:667
		// _ = "end of CoverTab[84273]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:668
	// _ = "end of CoverTab[84261]"
//line /usr/local/go/src/encoding/gob/decode.go:668
	_go_fuzz_dep_.CoverTab[84262]++
							dec.decodeArrayHelper(state, value, elemOp, n, ovfl, helper)
//line /usr/local/go/src/encoding/gob/decode.go:669
	// _ = "end of CoverTab[84262]"
}

//line /usr/local/go/src/encoding/gob/decode.go:673
func (dec *Decoder) ignoreSlice(state *decoderState, elemOp decOp) {
//line /usr/local/go/src/encoding/gob/decode.go:673
	_go_fuzz_dep_.CoverTab[84274]++
							dec.ignoreArrayHelper(state, elemOp, int(state.decodeUint()))
//line /usr/local/go/src/encoding/gob/decode.go:674
	// _ = "end of CoverTab[84274]"
}

//line /usr/local/go/src/encoding/gob/decode.go:680
func (dec *Decoder) decodeInterface(ityp reflect.Type, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:680
	_go_fuzz_dep_.CoverTab[84275]++

							nr := state.decodeUint()
							if nr > 1<<31 {
//line /usr/local/go/src/encoding/gob/decode.go:683
		_go_fuzz_dep_.CoverTab[84284]++
								errorf("invalid type name length %d", nr)
//line /usr/local/go/src/encoding/gob/decode.go:684
		// _ = "end of CoverTab[84284]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:685
		_go_fuzz_dep_.CoverTab[84285]++
//line /usr/local/go/src/encoding/gob/decode.go:685
		// _ = "end of CoverTab[84285]"
//line /usr/local/go/src/encoding/gob/decode.go:685
	}
//line /usr/local/go/src/encoding/gob/decode.go:685
	// _ = "end of CoverTab[84275]"
//line /usr/local/go/src/encoding/gob/decode.go:685
	_go_fuzz_dep_.CoverTab[84276]++
							if nr > uint64(state.b.Len()) {
//line /usr/local/go/src/encoding/gob/decode.go:686
		_go_fuzz_dep_.CoverTab[84286]++
								errorf("invalid type name length %d: exceeds input size", nr)
//line /usr/local/go/src/encoding/gob/decode.go:687
		// _ = "end of CoverTab[84286]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:688
		_go_fuzz_dep_.CoverTab[84287]++
//line /usr/local/go/src/encoding/gob/decode.go:688
		// _ = "end of CoverTab[84287]"
//line /usr/local/go/src/encoding/gob/decode.go:688
	}
//line /usr/local/go/src/encoding/gob/decode.go:688
	// _ = "end of CoverTab[84276]"
//line /usr/local/go/src/encoding/gob/decode.go:688
	_go_fuzz_dep_.CoverTab[84277]++
							n := int(nr)
							name := state.b.Bytes()[:n]
							state.b.Drop(n)

							if len(name) == 0 {
//line /usr/local/go/src/encoding/gob/decode.go:693
		_go_fuzz_dep_.CoverTab[84288]++

								value.Set(reflect.Zero(value.Type()))
								return
//line /usr/local/go/src/encoding/gob/decode.go:696
		// _ = "end of CoverTab[84288]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:697
		_go_fuzz_dep_.CoverTab[84289]++
//line /usr/local/go/src/encoding/gob/decode.go:697
		// _ = "end of CoverTab[84289]"
//line /usr/local/go/src/encoding/gob/decode.go:697
	}
//line /usr/local/go/src/encoding/gob/decode.go:697
	// _ = "end of CoverTab[84277]"
//line /usr/local/go/src/encoding/gob/decode.go:697
	_go_fuzz_dep_.CoverTab[84278]++
							if len(name) > 1024 {
//line /usr/local/go/src/encoding/gob/decode.go:698
		_go_fuzz_dep_.CoverTab[84290]++
								errorf("name too long (%d bytes): %.20q...", len(name), name)
//line /usr/local/go/src/encoding/gob/decode.go:699
		// _ = "end of CoverTab[84290]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:700
		_go_fuzz_dep_.CoverTab[84291]++
//line /usr/local/go/src/encoding/gob/decode.go:700
		// _ = "end of CoverTab[84291]"
//line /usr/local/go/src/encoding/gob/decode.go:700
	}
//line /usr/local/go/src/encoding/gob/decode.go:700
	// _ = "end of CoverTab[84278]"
//line /usr/local/go/src/encoding/gob/decode.go:700
	_go_fuzz_dep_.CoverTab[84279]++

							typi, ok := nameToConcreteType.Load(string(name))
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:703
		_go_fuzz_dep_.CoverTab[84292]++
								errorf("name not registered for interface: %q", name)
//line /usr/local/go/src/encoding/gob/decode.go:704
		// _ = "end of CoverTab[84292]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:705
		_go_fuzz_dep_.CoverTab[84293]++
//line /usr/local/go/src/encoding/gob/decode.go:705
		// _ = "end of CoverTab[84293]"
//line /usr/local/go/src/encoding/gob/decode.go:705
	}
//line /usr/local/go/src/encoding/gob/decode.go:705
	// _ = "end of CoverTab[84279]"
//line /usr/local/go/src/encoding/gob/decode.go:705
	_go_fuzz_dep_.CoverTab[84280]++
							typ := typi.(reflect.Type)

//line /usr/local/go/src/encoding/gob/decode.go:709
	concreteId := dec.decodeTypeSequence(true)
	if concreteId < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:710
		_go_fuzz_dep_.CoverTab[84294]++
								error_(dec.err)
//line /usr/local/go/src/encoding/gob/decode.go:711
		// _ = "end of CoverTab[84294]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:712
		_go_fuzz_dep_.CoverTab[84295]++
//line /usr/local/go/src/encoding/gob/decode.go:712
		// _ = "end of CoverTab[84295]"
//line /usr/local/go/src/encoding/gob/decode.go:712
	}
//line /usr/local/go/src/encoding/gob/decode.go:712
	// _ = "end of CoverTab[84280]"
//line /usr/local/go/src/encoding/gob/decode.go:712
	_go_fuzz_dep_.CoverTab[84281]++

//line /usr/local/go/src/encoding/gob/decode.go:715
	state.decodeUint()

	v := allocValue(typ)
	dec.decodeValue(concreteId, v)
	if dec.err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:719
		_go_fuzz_dep_.CoverTab[84296]++
								error_(dec.err)
//line /usr/local/go/src/encoding/gob/decode.go:720
		// _ = "end of CoverTab[84296]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:721
		_go_fuzz_dep_.CoverTab[84297]++
//line /usr/local/go/src/encoding/gob/decode.go:721
		// _ = "end of CoverTab[84297]"
//line /usr/local/go/src/encoding/gob/decode.go:721
	}
//line /usr/local/go/src/encoding/gob/decode.go:721
	// _ = "end of CoverTab[84281]"
//line /usr/local/go/src/encoding/gob/decode.go:721
	_go_fuzz_dep_.CoverTab[84282]++

//line /usr/local/go/src/encoding/gob/decode.go:724
	if !typ.AssignableTo(ityp) {
//line /usr/local/go/src/encoding/gob/decode.go:724
		_go_fuzz_dep_.CoverTab[84298]++
								errorf("%s is not assignable to type %s", typ, ityp)
//line /usr/local/go/src/encoding/gob/decode.go:725
		// _ = "end of CoverTab[84298]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:726
		_go_fuzz_dep_.CoverTab[84299]++
//line /usr/local/go/src/encoding/gob/decode.go:726
		// _ = "end of CoverTab[84299]"
//line /usr/local/go/src/encoding/gob/decode.go:726
	}
//line /usr/local/go/src/encoding/gob/decode.go:726
	// _ = "end of CoverTab[84282]"
//line /usr/local/go/src/encoding/gob/decode.go:726
	_go_fuzz_dep_.CoverTab[84283]++

							value.Set(v)
//line /usr/local/go/src/encoding/gob/decode.go:728
	// _ = "end of CoverTab[84283]"
}

//line /usr/local/go/src/encoding/gob/decode.go:732
func (dec *Decoder) ignoreInterface(state *decoderState) {
//line /usr/local/go/src/encoding/gob/decode.go:732
	_go_fuzz_dep_.CoverTab[84300]++

							n, ok := state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:735
		_go_fuzz_dep_.CoverTab[84305]++
								errorf("bad interface encoding: name too large for buffer")
//line /usr/local/go/src/encoding/gob/decode.go:736
		// _ = "end of CoverTab[84305]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:737
		_go_fuzz_dep_.CoverTab[84306]++
//line /usr/local/go/src/encoding/gob/decode.go:737
		// _ = "end of CoverTab[84306]"
//line /usr/local/go/src/encoding/gob/decode.go:737
	}
//line /usr/local/go/src/encoding/gob/decode.go:737
	// _ = "end of CoverTab[84300]"
//line /usr/local/go/src/encoding/gob/decode.go:737
	_go_fuzz_dep_.CoverTab[84301]++
							bn := state.b.Len()
							if bn < n {
//line /usr/local/go/src/encoding/gob/decode.go:739
		_go_fuzz_dep_.CoverTab[84307]++
								errorf("invalid interface value length %d: exceeds input size %d", n, bn)
//line /usr/local/go/src/encoding/gob/decode.go:740
		// _ = "end of CoverTab[84307]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:741
		_go_fuzz_dep_.CoverTab[84308]++
//line /usr/local/go/src/encoding/gob/decode.go:741
		// _ = "end of CoverTab[84308]"
//line /usr/local/go/src/encoding/gob/decode.go:741
	}
//line /usr/local/go/src/encoding/gob/decode.go:741
	// _ = "end of CoverTab[84301]"
//line /usr/local/go/src/encoding/gob/decode.go:741
	_go_fuzz_dep_.CoverTab[84302]++
							state.b.Drop(n)
							id := dec.decodeTypeSequence(true)
							if id < 0 {
//line /usr/local/go/src/encoding/gob/decode.go:744
		_go_fuzz_dep_.CoverTab[84309]++
								error_(dec.err)
//line /usr/local/go/src/encoding/gob/decode.go:745
		// _ = "end of CoverTab[84309]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:746
		_go_fuzz_dep_.CoverTab[84310]++
//line /usr/local/go/src/encoding/gob/decode.go:746
		// _ = "end of CoverTab[84310]"
//line /usr/local/go/src/encoding/gob/decode.go:746
	}
//line /usr/local/go/src/encoding/gob/decode.go:746
	// _ = "end of CoverTab[84302]"
//line /usr/local/go/src/encoding/gob/decode.go:746
	_go_fuzz_dep_.CoverTab[84303]++

							n, ok = state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:749
		_go_fuzz_dep_.CoverTab[84311]++
								errorf("bad interface encoding: data length too large for buffer")
//line /usr/local/go/src/encoding/gob/decode.go:750
		// _ = "end of CoverTab[84311]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:751
		_go_fuzz_dep_.CoverTab[84312]++
//line /usr/local/go/src/encoding/gob/decode.go:751
		// _ = "end of CoverTab[84312]"
//line /usr/local/go/src/encoding/gob/decode.go:751
	}
//line /usr/local/go/src/encoding/gob/decode.go:751
	// _ = "end of CoverTab[84303]"
//line /usr/local/go/src/encoding/gob/decode.go:751
	_go_fuzz_dep_.CoverTab[84304]++
							state.b.Drop(n)
//line /usr/local/go/src/encoding/gob/decode.go:752
	// _ = "end of CoverTab[84304]"
}

//line /usr/local/go/src/encoding/gob/decode.go:757
func (dec *Decoder) decodeGobDecoder(ut *userTypeInfo, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:757
	_go_fuzz_dep_.CoverTab[84313]++

							n, ok := state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:760
		_go_fuzz_dep_.CoverTab[84317]++
								errorf("GobDecoder: length too large for buffer")
//line /usr/local/go/src/encoding/gob/decode.go:761
		// _ = "end of CoverTab[84317]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:762
		_go_fuzz_dep_.CoverTab[84318]++
//line /usr/local/go/src/encoding/gob/decode.go:762
		// _ = "end of CoverTab[84318]"
//line /usr/local/go/src/encoding/gob/decode.go:762
	}
//line /usr/local/go/src/encoding/gob/decode.go:762
	// _ = "end of CoverTab[84313]"
//line /usr/local/go/src/encoding/gob/decode.go:762
	_go_fuzz_dep_.CoverTab[84314]++
							b := state.b.Bytes()
							if len(b) < n {
//line /usr/local/go/src/encoding/gob/decode.go:764
		_go_fuzz_dep_.CoverTab[84319]++
								errorf("GobDecoder: invalid data length %d: exceeds input size %d", n, len(b))
//line /usr/local/go/src/encoding/gob/decode.go:765
		// _ = "end of CoverTab[84319]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:766
		_go_fuzz_dep_.CoverTab[84320]++
//line /usr/local/go/src/encoding/gob/decode.go:766
		// _ = "end of CoverTab[84320]"
//line /usr/local/go/src/encoding/gob/decode.go:766
	}
//line /usr/local/go/src/encoding/gob/decode.go:766
	// _ = "end of CoverTab[84314]"
//line /usr/local/go/src/encoding/gob/decode.go:766
	_go_fuzz_dep_.CoverTab[84315]++
							b = b[:n]
							state.b.Drop(n)
							var err error

							switch ut.externalDec {
	case xGob:
//line /usr/local/go/src/encoding/gob/decode.go:772
		_go_fuzz_dep_.CoverTab[84321]++
								err = value.Interface().(GobDecoder).GobDecode(b)
//line /usr/local/go/src/encoding/gob/decode.go:773
		// _ = "end of CoverTab[84321]"
	case xBinary:
//line /usr/local/go/src/encoding/gob/decode.go:774
		_go_fuzz_dep_.CoverTab[84322]++
								err = value.Interface().(encoding.BinaryUnmarshaler).UnmarshalBinary(b)
//line /usr/local/go/src/encoding/gob/decode.go:775
		// _ = "end of CoverTab[84322]"
	case xText:
//line /usr/local/go/src/encoding/gob/decode.go:776
		_go_fuzz_dep_.CoverTab[84323]++
								err = value.Interface().(encoding.TextUnmarshaler).UnmarshalText(b)
//line /usr/local/go/src/encoding/gob/decode.go:777
		// _ = "end of CoverTab[84323]"
//line /usr/local/go/src/encoding/gob/decode.go:777
	default:
//line /usr/local/go/src/encoding/gob/decode.go:777
		_go_fuzz_dep_.CoverTab[84324]++
//line /usr/local/go/src/encoding/gob/decode.go:777
		// _ = "end of CoverTab[84324]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:778
	// _ = "end of CoverTab[84315]"
//line /usr/local/go/src/encoding/gob/decode.go:778
	_go_fuzz_dep_.CoverTab[84316]++
							if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:779
		_go_fuzz_dep_.CoverTab[84325]++
								error_(err)
//line /usr/local/go/src/encoding/gob/decode.go:780
		// _ = "end of CoverTab[84325]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:781
		_go_fuzz_dep_.CoverTab[84326]++
//line /usr/local/go/src/encoding/gob/decode.go:781
		// _ = "end of CoverTab[84326]"
//line /usr/local/go/src/encoding/gob/decode.go:781
	}
//line /usr/local/go/src/encoding/gob/decode.go:781
	// _ = "end of CoverTab[84316]"
}

//line /usr/local/go/src/encoding/gob/decode.go:785
func (dec *Decoder) ignoreGobDecoder(state *decoderState) {
//line /usr/local/go/src/encoding/gob/decode.go:785
	_go_fuzz_dep_.CoverTab[84327]++

							n, ok := state.getLength()
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:788
		_go_fuzz_dep_.CoverTab[84330]++
								errorf("GobDecoder: length too large for buffer")
//line /usr/local/go/src/encoding/gob/decode.go:789
		// _ = "end of CoverTab[84330]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:790
		_go_fuzz_dep_.CoverTab[84331]++
//line /usr/local/go/src/encoding/gob/decode.go:790
		// _ = "end of CoverTab[84331]"
//line /usr/local/go/src/encoding/gob/decode.go:790
	}
//line /usr/local/go/src/encoding/gob/decode.go:790
	// _ = "end of CoverTab[84327]"
//line /usr/local/go/src/encoding/gob/decode.go:790
	_go_fuzz_dep_.CoverTab[84328]++
							bn := state.b.Len()
							if bn < n {
//line /usr/local/go/src/encoding/gob/decode.go:792
		_go_fuzz_dep_.CoverTab[84332]++
								errorf("GobDecoder: invalid data length %d: exceeds input size %d", n, bn)
//line /usr/local/go/src/encoding/gob/decode.go:793
		// _ = "end of CoverTab[84332]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:794
		_go_fuzz_dep_.CoverTab[84333]++
//line /usr/local/go/src/encoding/gob/decode.go:794
		// _ = "end of CoverTab[84333]"
//line /usr/local/go/src/encoding/gob/decode.go:794
	}
//line /usr/local/go/src/encoding/gob/decode.go:794
	// _ = "end of CoverTab[84328]"
//line /usr/local/go/src/encoding/gob/decode.go:794
	_go_fuzz_dep_.CoverTab[84329]++
							state.b.Drop(n)
//line /usr/local/go/src/encoding/gob/decode.go:795
	// _ = "end of CoverTab[84329]"
}

//line /usr/local/go/src/encoding/gob/decode.go:799
var decOpTable = [...]decOp{
	reflect.Bool:		decBool,
	reflect.Int8:		decInt8,
	reflect.Int16:		decInt16,
	reflect.Int32:		decInt32,
	reflect.Int64:		decInt64,
	reflect.Uint8:		decUint8,
	reflect.Uint16:		decUint16,
	reflect.Uint32:		decUint32,
	reflect.Uint64:		decUint64,
	reflect.Float32:	decFloat32,
	reflect.Float64:	decFloat64,
	reflect.Complex64:	decComplex64,
	reflect.Complex128:	decComplex128,
	reflect.String:		decString,
}

//line /usr/local/go/src/encoding/gob/decode.go:817
var decIgnoreOpMap = map[typeId]decOp{
	tBool:		ignoreUint,
	tInt:		ignoreUint,
	tUint:		ignoreUint,
	tFloat:		ignoreUint,
	tBytes:		ignoreUint8Array,
	tString:	ignoreUint8Array,
	tComplex:	ignoreTwoUints,
}

//line /usr/local/go/src/encoding/gob/decode.go:829
func (dec *Decoder) decOpFor(wireId typeId, rt reflect.Type, name string, inProgress map[reflect.Type]*decOp) *decOp {
//line /usr/local/go/src/encoding/gob/decode.go:829
	_go_fuzz_dep_.CoverTab[84334]++
							ut := userType(rt)

							if ut.externalDec != 0 {
//line /usr/local/go/src/encoding/gob/decode.go:832
		_go_fuzz_dep_.CoverTab[84340]++
								return dec.gobDecodeOpFor(ut)
//line /usr/local/go/src/encoding/gob/decode.go:833
		// _ = "end of CoverTab[84340]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:834
		_go_fuzz_dep_.CoverTab[84341]++
//line /usr/local/go/src/encoding/gob/decode.go:834
		// _ = "end of CoverTab[84341]"
//line /usr/local/go/src/encoding/gob/decode.go:834
	}
//line /usr/local/go/src/encoding/gob/decode.go:834
	// _ = "end of CoverTab[84334]"
//line /usr/local/go/src/encoding/gob/decode.go:834
	_go_fuzz_dep_.CoverTab[84335]++

//line /usr/local/go/src/encoding/gob/decode.go:838
	if opPtr := inProgress[rt]; opPtr != nil {
//line /usr/local/go/src/encoding/gob/decode.go:838
		_go_fuzz_dep_.CoverTab[84342]++
								return opPtr
//line /usr/local/go/src/encoding/gob/decode.go:839
		// _ = "end of CoverTab[84342]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:840
		_go_fuzz_dep_.CoverTab[84343]++
//line /usr/local/go/src/encoding/gob/decode.go:840
		// _ = "end of CoverTab[84343]"
//line /usr/local/go/src/encoding/gob/decode.go:840
	}
//line /usr/local/go/src/encoding/gob/decode.go:840
	// _ = "end of CoverTab[84335]"
//line /usr/local/go/src/encoding/gob/decode.go:840
	_go_fuzz_dep_.CoverTab[84336]++
							typ := ut.base
							var op decOp
							k := typ.Kind()
							if int(k) < len(decOpTable) {
//line /usr/local/go/src/encoding/gob/decode.go:844
		_go_fuzz_dep_.CoverTab[84344]++
								op = decOpTable[k]
//line /usr/local/go/src/encoding/gob/decode.go:845
		// _ = "end of CoverTab[84344]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:846
		_go_fuzz_dep_.CoverTab[84345]++
//line /usr/local/go/src/encoding/gob/decode.go:846
		// _ = "end of CoverTab[84345]"
//line /usr/local/go/src/encoding/gob/decode.go:846
	}
//line /usr/local/go/src/encoding/gob/decode.go:846
	// _ = "end of CoverTab[84336]"
//line /usr/local/go/src/encoding/gob/decode.go:846
	_go_fuzz_dep_.CoverTab[84337]++
							if op == nil {
//line /usr/local/go/src/encoding/gob/decode.go:847
		_go_fuzz_dep_.CoverTab[84346]++
								inProgress[rt] = &op

								switch t := typ; t.Kind() {
		case reflect.Array:
//line /usr/local/go/src/encoding/gob/decode.go:851
			_go_fuzz_dep_.CoverTab[84347]++
									name = "element of " + name
									elemId := dec.wireType[wireId].ArrayT.Elem
									elemOp := dec.decOpFor(elemId, t.Elem(), name, inProgress)
									ovfl := overflow(name)
									helper := decArrayHelper[t.Elem().Kind()]
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:857
				_go_fuzz_dep_.CoverTab[84356]++
										state.dec.decodeArray(state, value, *elemOp, t.Len(), ovfl, helper)
//line /usr/local/go/src/encoding/gob/decode.go:858
				// _ = "end of CoverTab[84356]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:859
			// _ = "end of CoverTab[84347]"

		case reflect.Map:
//line /usr/local/go/src/encoding/gob/decode.go:861
			_go_fuzz_dep_.CoverTab[84348]++
									keyId := dec.wireType[wireId].MapT.Key
									elemId := dec.wireType[wireId].MapT.Elem
									keyOp := dec.decOpFor(keyId, t.Key(), "key of "+name, inProgress)
									elemOp := dec.decOpFor(elemId, t.Elem(), "element of "+name, inProgress)
									ovfl := overflow(name)
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:867
				_go_fuzz_dep_.CoverTab[84357]++
										state.dec.decodeMap(t, state, value, *keyOp, *elemOp, ovfl)
//line /usr/local/go/src/encoding/gob/decode.go:868
				// _ = "end of CoverTab[84357]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:869
			// _ = "end of CoverTab[84348]"

		case reflect.Slice:
//line /usr/local/go/src/encoding/gob/decode.go:871
			_go_fuzz_dep_.CoverTab[84349]++
									name = "element of " + name
									if t.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/gob/decode.go:873
				_go_fuzz_dep_.CoverTab[84358]++
										op = decUint8Slice
										break
//line /usr/local/go/src/encoding/gob/decode.go:875
				// _ = "end of CoverTab[84358]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:876
				_go_fuzz_dep_.CoverTab[84359]++
//line /usr/local/go/src/encoding/gob/decode.go:876
				// _ = "end of CoverTab[84359]"
//line /usr/local/go/src/encoding/gob/decode.go:876
			}
//line /usr/local/go/src/encoding/gob/decode.go:876
			// _ = "end of CoverTab[84349]"
//line /usr/local/go/src/encoding/gob/decode.go:876
			_go_fuzz_dep_.CoverTab[84350]++
									var elemId typeId
									if tt, ok := builtinIdToType[wireId]; ok {
//line /usr/local/go/src/encoding/gob/decode.go:878
				_go_fuzz_dep_.CoverTab[84360]++
										elemId = tt.(*sliceType).Elem
//line /usr/local/go/src/encoding/gob/decode.go:879
				// _ = "end of CoverTab[84360]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:880
				_go_fuzz_dep_.CoverTab[84361]++
										elemId = dec.wireType[wireId].SliceT.Elem
//line /usr/local/go/src/encoding/gob/decode.go:881
				// _ = "end of CoverTab[84361]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:882
			// _ = "end of CoverTab[84350]"
//line /usr/local/go/src/encoding/gob/decode.go:882
			_go_fuzz_dep_.CoverTab[84351]++
									elemOp := dec.decOpFor(elemId, t.Elem(), name, inProgress)
									ovfl := overflow(name)
									helper := decSliceHelper[t.Elem().Kind()]
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:886
				_go_fuzz_dep_.CoverTab[84362]++
										state.dec.decodeSlice(state, value, *elemOp, ovfl, helper)
//line /usr/local/go/src/encoding/gob/decode.go:887
				// _ = "end of CoverTab[84362]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:888
			// _ = "end of CoverTab[84351]"

		case reflect.Struct:
//line /usr/local/go/src/encoding/gob/decode.go:890
			_go_fuzz_dep_.CoverTab[84352]++

									ut := userType(typ)
									enginePtr, err := dec.getDecEnginePtr(wireId, ut)
									if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:894
				_go_fuzz_dep_.CoverTab[84363]++
										error_(err)
//line /usr/local/go/src/encoding/gob/decode.go:895
				// _ = "end of CoverTab[84363]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:896
				_go_fuzz_dep_.CoverTab[84364]++
//line /usr/local/go/src/encoding/gob/decode.go:896
				// _ = "end of CoverTab[84364]"
//line /usr/local/go/src/encoding/gob/decode.go:896
			}
//line /usr/local/go/src/encoding/gob/decode.go:896
			// _ = "end of CoverTab[84352]"
//line /usr/local/go/src/encoding/gob/decode.go:896
			_go_fuzz_dep_.CoverTab[84353]++
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:897
				_go_fuzz_dep_.CoverTab[84365]++

										dec.decodeStruct(*enginePtr, value)
//line /usr/local/go/src/encoding/gob/decode.go:899
				// _ = "end of CoverTab[84365]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:900
			// _ = "end of CoverTab[84353]"
		case reflect.Interface:
//line /usr/local/go/src/encoding/gob/decode.go:901
			_go_fuzz_dep_.CoverTab[84354]++
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:902
				_go_fuzz_dep_.CoverTab[84366]++
										state.dec.decodeInterface(t, state, value)
//line /usr/local/go/src/encoding/gob/decode.go:903
				// _ = "end of CoverTab[84366]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:904
			// _ = "end of CoverTab[84354]"
//line /usr/local/go/src/encoding/gob/decode.go:904
		default:
//line /usr/local/go/src/encoding/gob/decode.go:904
			_go_fuzz_dep_.CoverTab[84355]++
//line /usr/local/go/src/encoding/gob/decode.go:904
			// _ = "end of CoverTab[84355]"
		}
//line /usr/local/go/src/encoding/gob/decode.go:905
		// _ = "end of CoverTab[84346]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:906
		_go_fuzz_dep_.CoverTab[84367]++
//line /usr/local/go/src/encoding/gob/decode.go:906
		// _ = "end of CoverTab[84367]"
//line /usr/local/go/src/encoding/gob/decode.go:906
	}
//line /usr/local/go/src/encoding/gob/decode.go:906
	// _ = "end of CoverTab[84337]"
//line /usr/local/go/src/encoding/gob/decode.go:906
	_go_fuzz_dep_.CoverTab[84338]++
							if op == nil {
//line /usr/local/go/src/encoding/gob/decode.go:907
		_go_fuzz_dep_.CoverTab[84368]++
								errorf("decode can't handle type %s", rt)
//line /usr/local/go/src/encoding/gob/decode.go:908
		// _ = "end of CoverTab[84368]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:909
		_go_fuzz_dep_.CoverTab[84369]++
//line /usr/local/go/src/encoding/gob/decode.go:909
		// _ = "end of CoverTab[84369]"
//line /usr/local/go/src/encoding/gob/decode.go:909
	}
//line /usr/local/go/src/encoding/gob/decode.go:909
	// _ = "end of CoverTab[84338]"
//line /usr/local/go/src/encoding/gob/decode.go:909
	_go_fuzz_dep_.CoverTab[84339]++
							return &op
//line /usr/local/go/src/encoding/gob/decode.go:910
	// _ = "end of CoverTab[84339]"
}

var maxIgnoreNestingDepth = 10000

//line /usr/local/go/src/encoding/gob/decode.go:916
func (dec *Decoder) decIgnoreOpFor(wireId typeId, inProgress map[typeId]*decOp, depth int) *decOp {
//line /usr/local/go/src/encoding/gob/decode.go:916
	_go_fuzz_dep_.CoverTab[84370]++
							if depth > maxIgnoreNestingDepth {
//line /usr/local/go/src/encoding/gob/decode.go:917
		_go_fuzz_dep_.CoverTab[84375]++
								error_(errors.New("invalid nesting depth"))
//line /usr/local/go/src/encoding/gob/decode.go:918
		// _ = "end of CoverTab[84375]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:919
		_go_fuzz_dep_.CoverTab[84376]++
//line /usr/local/go/src/encoding/gob/decode.go:919
		// _ = "end of CoverTab[84376]"
//line /usr/local/go/src/encoding/gob/decode.go:919
	}
//line /usr/local/go/src/encoding/gob/decode.go:919
	// _ = "end of CoverTab[84370]"
//line /usr/local/go/src/encoding/gob/decode.go:919
	_go_fuzz_dep_.CoverTab[84371]++

//line /usr/local/go/src/encoding/gob/decode.go:922
	if opPtr := inProgress[wireId]; opPtr != nil {
//line /usr/local/go/src/encoding/gob/decode.go:922
		_go_fuzz_dep_.CoverTab[84377]++
								return opPtr
//line /usr/local/go/src/encoding/gob/decode.go:923
		// _ = "end of CoverTab[84377]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:924
		_go_fuzz_dep_.CoverTab[84378]++
//line /usr/local/go/src/encoding/gob/decode.go:924
		// _ = "end of CoverTab[84378]"
//line /usr/local/go/src/encoding/gob/decode.go:924
	}
//line /usr/local/go/src/encoding/gob/decode.go:924
	// _ = "end of CoverTab[84371]"
//line /usr/local/go/src/encoding/gob/decode.go:924
	_go_fuzz_dep_.CoverTab[84372]++
							op, ok := decIgnoreOpMap[wireId]
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:926
		_go_fuzz_dep_.CoverTab[84379]++
								inProgress[wireId] = &op
								if wireId == tInterface {
//line /usr/local/go/src/encoding/gob/decode.go:928
			_go_fuzz_dep_.CoverTab[84381]++

//line /usr/local/go/src/encoding/gob/decode.go:931
			op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:931
				_go_fuzz_dep_.CoverTab[84383]++
										state.dec.ignoreInterface(state)
//line /usr/local/go/src/encoding/gob/decode.go:932
				// _ = "end of CoverTab[84383]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:933
			// _ = "end of CoverTab[84381]"
//line /usr/local/go/src/encoding/gob/decode.go:933
			_go_fuzz_dep_.CoverTab[84382]++
									return &op
//line /usr/local/go/src/encoding/gob/decode.go:934
			// _ = "end of CoverTab[84382]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:935
			_go_fuzz_dep_.CoverTab[84384]++
//line /usr/local/go/src/encoding/gob/decode.go:935
			// _ = "end of CoverTab[84384]"
//line /usr/local/go/src/encoding/gob/decode.go:935
		}
//line /usr/local/go/src/encoding/gob/decode.go:935
		// _ = "end of CoverTab[84379]"
//line /usr/local/go/src/encoding/gob/decode.go:935
		_go_fuzz_dep_.CoverTab[84380]++

								wire := dec.wireType[wireId]
								switch {
		case wire == nil:
//line /usr/local/go/src/encoding/gob/decode.go:939
			_go_fuzz_dep_.CoverTab[84385]++
									errorf("bad data: undefined type %s", wireId.string())
//line /usr/local/go/src/encoding/gob/decode.go:940
			// _ = "end of CoverTab[84385]"
		case wire.ArrayT != nil:
//line /usr/local/go/src/encoding/gob/decode.go:941
			_go_fuzz_dep_.CoverTab[84386]++
									elemId := wire.ArrayT.Elem
									elemOp := dec.decIgnoreOpFor(elemId, inProgress, depth+1)
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:944
				_go_fuzz_dep_.CoverTab[84393]++
										state.dec.ignoreArray(state, *elemOp, wire.ArrayT.Len)
//line /usr/local/go/src/encoding/gob/decode.go:945
				// _ = "end of CoverTab[84393]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:946
			// _ = "end of CoverTab[84386]"

		case wire.MapT != nil:
//line /usr/local/go/src/encoding/gob/decode.go:948
			_go_fuzz_dep_.CoverTab[84387]++
									keyId := dec.wireType[wireId].MapT.Key
									elemId := dec.wireType[wireId].MapT.Elem
									keyOp := dec.decIgnoreOpFor(keyId, inProgress, depth+1)
									elemOp := dec.decIgnoreOpFor(elemId, inProgress, depth+1)
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:953
				_go_fuzz_dep_.CoverTab[84394]++
										state.dec.ignoreMap(state, *keyOp, *elemOp)
//line /usr/local/go/src/encoding/gob/decode.go:954
				// _ = "end of CoverTab[84394]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:955
			// _ = "end of CoverTab[84387]"

		case wire.SliceT != nil:
//line /usr/local/go/src/encoding/gob/decode.go:957
			_go_fuzz_dep_.CoverTab[84388]++
									elemId := wire.SliceT.Elem
									elemOp := dec.decIgnoreOpFor(elemId, inProgress, depth+1)
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:960
				_go_fuzz_dep_.CoverTab[84395]++
										state.dec.ignoreSlice(state, *elemOp)
//line /usr/local/go/src/encoding/gob/decode.go:961
				// _ = "end of CoverTab[84395]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:962
			// _ = "end of CoverTab[84388]"

		case wire.StructT != nil:
//line /usr/local/go/src/encoding/gob/decode.go:964
			_go_fuzz_dep_.CoverTab[84389]++

									enginePtr, err := dec.getIgnoreEnginePtr(wireId)
									if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:967
				_go_fuzz_dep_.CoverTab[84396]++
										error_(err)
//line /usr/local/go/src/encoding/gob/decode.go:968
				// _ = "end of CoverTab[84396]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:969
				_go_fuzz_dep_.CoverTab[84397]++
//line /usr/local/go/src/encoding/gob/decode.go:969
				// _ = "end of CoverTab[84397]"
//line /usr/local/go/src/encoding/gob/decode.go:969
			}
//line /usr/local/go/src/encoding/gob/decode.go:969
			// _ = "end of CoverTab[84389]"
//line /usr/local/go/src/encoding/gob/decode.go:969
			_go_fuzz_dep_.CoverTab[84390]++
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:970
				_go_fuzz_dep_.CoverTab[84398]++

										state.dec.ignoreStruct(*enginePtr)
//line /usr/local/go/src/encoding/gob/decode.go:972
				// _ = "end of CoverTab[84398]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:973
			// _ = "end of CoverTab[84390]"

		case wire.GobEncoderT != nil, wire.BinaryMarshalerT != nil, wire.TextMarshalerT != nil:
//line /usr/local/go/src/encoding/gob/decode.go:975
			_go_fuzz_dep_.CoverTab[84391]++
									op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:976
				_go_fuzz_dep_.CoverTab[84399]++
										state.dec.ignoreGobDecoder(state)
//line /usr/local/go/src/encoding/gob/decode.go:977
				// _ = "end of CoverTab[84399]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:978
			// _ = "end of CoverTab[84391]"
//line /usr/local/go/src/encoding/gob/decode.go:978
		default:
//line /usr/local/go/src/encoding/gob/decode.go:978
			_go_fuzz_dep_.CoverTab[84392]++
//line /usr/local/go/src/encoding/gob/decode.go:978
			// _ = "end of CoverTab[84392]"
		}
//line /usr/local/go/src/encoding/gob/decode.go:979
		// _ = "end of CoverTab[84380]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:980
		_go_fuzz_dep_.CoverTab[84400]++
//line /usr/local/go/src/encoding/gob/decode.go:980
		// _ = "end of CoverTab[84400]"
//line /usr/local/go/src/encoding/gob/decode.go:980
	}
//line /usr/local/go/src/encoding/gob/decode.go:980
	// _ = "end of CoverTab[84372]"
//line /usr/local/go/src/encoding/gob/decode.go:980
	_go_fuzz_dep_.CoverTab[84373]++
							if op == nil {
//line /usr/local/go/src/encoding/gob/decode.go:981
		_go_fuzz_dep_.CoverTab[84401]++
								errorf("bad data: ignore can't handle type %s", wireId.string())
//line /usr/local/go/src/encoding/gob/decode.go:982
		// _ = "end of CoverTab[84401]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:983
		_go_fuzz_dep_.CoverTab[84402]++
//line /usr/local/go/src/encoding/gob/decode.go:983
		// _ = "end of CoverTab[84402]"
//line /usr/local/go/src/encoding/gob/decode.go:983
	}
//line /usr/local/go/src/encoding/gob/decode.go:983
	// _ = "end of CoverTab[84373]"
//line /usr/local/go/src/encoding/gob/decode.go:983
	_go_fuzz_dep_.CoverTab[84374]++
							return &op
//line /usr/local/go/src/encoding/gob/decode.go:984
	// _ = "end of CoverTab[84374]"
}

//line /usr/local/go/src/encoding/gob/decode.go:989
func (dec *Decoder) gobDecodeOpFor(ut *userTypeInfo) *decOp {
//line /usr/local/go/src/encoding/gob/decode.go:989
	_go_fuzz_dep_.CoverTab[84403]++
							rcvrType := ut.user
							if ut.decIndir == -1 {
//line /usr/local/go/src/encoding/gob/decode.go:991
		_go_fuzz_dep_.CoverTab[84406]++
								rcvrType = reflect.PointerTo(rcvrType)
//line /usr/local/go/src/encoding/gob/decode.go:992
		// _ = "end of CoverTab[84406]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:993
		_go_fuzz_dep_.CoverTab[84407]++
//line /usr/local/go/src/encoding/gob/decode.go:993
		if ut.decIndir > 0 {
//line /usr/local/go/src/encoding/gob/decode.go:993
			_go_fuzz_dep_.CoverTab[84408]++
									for i := int8(0); i < ut.decIndir; i++ {
//line /usr/local/go/src/encoding/gob/decode.go:994
				_go_fuzz_dep_.CoverTab[84409]++
										rcvrType = rcvrType.Elem()
//line /usr/local/go/src/encoding/gob/decode.go:995
				// _ = "end of CoverTab[84409]"
			}
//line /usr/local/go/src/encoding/gob/decode.go:996
			// _ = "end of CoverTab[84408]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:997
			_go_fuzz_dep_.CoverTab[84410]++
//line /usr/local/go/src/encoding/gob/decode.go:997
			// _ = "end of CoverTab[84410]"
//line /usr/local/go/src/encoding/gob/decode.go:997
		}
//line /usr/local/go/src/encoding/gob/decode.go:997
		// _ = "end of CoverTab[84407]"
//line /usr/local/go/src/encoding/gob/decode.go:997
	}
//line /usr/local/go/src/encoding/gob/decode.go:997
	// _ = "end of CoverTab[84403]"
//line /usr/local/go/src/encoding/gob/decode.go:997
	_go_fuzz_dep_.CoverTab[84404]++
							var op decOp
							op = func(i *decInstr, state *decoderState, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:999
		_go_fuzz_dep_.CoverTab[84411]++

								if value.Kind() != reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1001
			_go_fuzz_dep_.CoverTab[84413]++
//line /usr/local/go/src/encoding/gob/decode.go:1001
			return rcvrType.Kind() == reflect.Pointer
//line /usr/local/go/src/encoding/gob/decode.go:1001
			// _ = "end of CoverTab[84413]"
//line /usr/local/go/src/encoding/gob/decode.go:1001
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1001
			_go_fuzz_dep_.CoverTab[84414]++
									value = value.Addr()
//line /usr/local/go/src/encoding/gob/decode.go:1002
			// _ = "end of CoverTab[84414]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1003
			_go_fuzz_dep_.CoverTab[84415]++
//line /usr/local/go/src/encoding/gob/decode.go:1003
			// _ = "end of CoverTab[84415]"
//line /usr/local/go/src/encoding/gob/decode.go:1003
		}
//line /usr/local/go/src/encoding/gob/decode.go:1003
		// _ = "end of CoverTab[84411]"
//line /usr/local/go/src/encoding/gob/decode.go:1003
		_go_fuzz_dep_.CoverTab[84412]++
								state.dec.decodeGobDecoder(ut, state, value)
//line /usr/local/go/src/encoding/gob/decode.go:1004
		// _ = "end of CoverTab[84412]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:1005
	// _ = "end of CoverTab[84404]"
//line /usr/local/go/src/encoding/gob/decode.go:1005
	_go_fuzz_dep_.CoverTab[84405]++
							return &op
//line /usr/local/go/src/encoding/gob/decode.go:1006
	// _ = "end of CoverTab[84405]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1013
func (dec *Decoder) compatibleType(fr reflect.Type, fw typeId, inProgress map[reflect.Type]typeId) bool {
//line /usr/local/go/src/encoding/gob/decode.go:1013
	_go_fuzz_dep_.CoverTab[84416]++
							if rhs, ok := inProgress[fr]; ok {
//line /usr/local/go/src/encoding/gob/decode.go:1014
		_go_fuzz_dep_.CoverTab[84420]++
								return rhs == fw
//line /usr/local/go/src/encoding/gob/decode.go:1015
		// _ = "end of CoverTab[84420]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1016
		_go_fuzz_dep_.CoverTab[84421]++
//line /usr/local/go/src/encoding/gob/decode.go:1016
		// _ = "end of CoverTab[84421]"
//line /usr/local/go/src/encoding/gob/decode.go:1016
	}
//line /usr/local/go/src/encoding/gob/decode.go:1016
	// _ = "end of CoverTab[84416]"
//line /usr/local/go/src/encoding/gob/decode.go:1016
	_go_fuzz_dep_.CoverTab[84417]++
							inProgress[fr] = fw
							ut := userType(fr)
							wire, ok := dec.wireType[fw]

//line /usr/local/go/src/encoding/gob/decode.go:1026
	if (ut.externalDec == xGob) != (ok && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1026
		_go_fuzz_dep_.CoverTab[84422]++
//line /usr/local/go/src/encoding/gob/decode.go:1026
		return wire.GobEncoderT != nil
//line /usr/local/go/src/encoding/gob/decode.go:1026
		// _ = "end of CoverTab[84422]"
//line /usr/local/go/src/encoding/gob/decode.go:1026
	}()) || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1026
		_go_fuzz_dep_.CoverTab[84423]++
//line /usr/local/go/src/encoding/gob/decode.go:1026
		return (ut.externalDec == xBinary) != (ok && func() bool {
									_go_fuzz_dep_.CoverTab[84424]++
//line /usr/local/go/src/encoding/gob/decode.go:1027
			return wire.BinaryMarshalerT != nil
//line /usr/local/go/src/encoding/gob/decode.go:1027
			// _ = "end of CoverTab[84424]"
//line /usr/local/go/src/encoding/gob/decode.go:1027
		}())
//line /usr/local/go/src/encoding/gob/decode.go:1027
		// _ = "end of CoverTab[84423]"
//line /usr/local/go/src/encoding/gob/decode.go:1027
	}() || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1027
		_go_fuzz_dep_.CoverTab[84425]++
//line /usr/local/go/src/encoding/gob/decode.go:1027
		return (ut.externalDec == xText) != (ok && func() bool {
									_go_fuzz_dep_.CoverTab[84426]++
//line /usr/local/go/src/encoding/gob/decode.go:1028
			return wire.TextMarshalerT != nil
//line /usr/local/go/src/encoding/gob/decode.go:1028
			// _ = "end of CoverTab[84426]"
//line /usr/local/go/src/encoding/gob/decode.go:1028
		}())
//line /usr/local/go/src/encoding/gob/decode.go:1028
		// _ = "end of CoverTab[84425]"
//line /usr/local/go/src/encoding/gob/decode.go:1028
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:1028
		_go_fuzz_dep_.CoverTab[84427]++
								return false
//line /usr/local/go/src/encoding/gob/decode.go:1029
		// _ = "end of CoverTab[84427]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1030
		_go_fuzz_dep_.CoverTab[84428]++
//line /usr/local/go/src/encoding/gob/decode.go:1030
		// _ = "end of CoverTab[84428]"
//line /usr/local/go/src/encoding/gob/decode.go:1030
	}
//line /usr/local/go/src/encoding/gob/decode.go:1030
	// _ = "end of CoverTab[84417]"
//line /usr/local/go/src/encoding/gob/decode.go:1030
	_go_fuzz_dep_.CoverTab[84418]++
							if ut.externalDec != 0 {
//line /usr/local/go/src/encoding/gob/decode.go:1031
		_go_fuzz_dep_.CoverTab[84429]++
								return true
//line /usr/local/go/src/encoding/gob/decode.go:1032
		// _ = "end of CoverTab[84429]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1033
		_go_fuzz_dep_.CoverTab[84430]++
//line /usr/local/go/src/encoding/gob/decode.go:1033
		// _ = "end of CoverTab[84430]"
//line /usr/local/go/src/encoding/gob/decode.go:1033
	}
//line /usr/local/go/src/encoding/gob/decode.go:1033
	// _ = "end of CoverTab[84418]"
//line /usr/local/go/src/encoding/gob/decode.go:1033
	_go_fuzz_dep_.CoverTab[84419]++
							switch t := ut.base; t.Kind() {
	default:
//line /usr/local/go/src/encoding/gob/decode.go:1035
		_go_fuzz_dep_.CoverTab[84431]++

								return false
//line /usr/local/go/src/encoding/gob/decode.go:1037
		// _ = "end of CoverTab[84431]"
	case reflect.Bool:
//line /usr/local/go/src/encoding/gob/decode.go:1038
		_go_fuzz_dep_.CoverTab[84432]++
								return fw == tBool
//line /usr/local/go/src/encoding/gob/decode.go:1039
		// _ = "end of CoverTab[84432]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/gob/decode.go:1040
		_go_fuzz_dep_.CoverTab[84433]++
								return fw == tInt
//line /usr/local/go/src/encoding/gob/decode.go:1041
		// _ = "end of CoverTab[84433]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/gob/decode.go:1042
		_go_fuzz_dep_.CoverTab[84434]++
								return fw == tUint
//line /usr/local/go/src/encoding/gob/decode.go:1043
		// _ = "end of CoverTab[84434]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/encoding/gob/decode.go:1044
		_go_fuzz_dep_.CoverTab[84435]++
								return fw == tFloat
//line /usr/local/go/src/encoding/gob/decode.go:1045
		// _ = "end of CoverTab[84435]"
	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/encoding/gob/decode.go:1046
		_go_fuzz_dep_.CoverTab[84436]++
								return fw == tComplex
//line /usr/local/go/src/encoding/gob/decode.go:1047
		// _ = "end of CoverTab[84436]"
	case reflect.String:
//line /usr/local/go/src/encoding/gob/decode.go:1048
		_go_fuzz_dep_.CoverTab[84437]++
								return fw == tString
//line /usr/local/go/src/encoding/gob/decode.go:1049
		// _ = "end of CoverTab[84437]"
	case reflect.Interface:
//line /usr/local/go/src/encoding/gob/decode.go:1050
		_go_fuzz_dep_.CoverTab[84438]++
								return fw == tInterface
//line /usr/local/go/src/encoding/gob/decode.go:1051
		// _ = "end of CoverTab[84438]"
	case reflect.Array:
//line /usr/local/go/src/encoding/gob/decode.go:1052
		_go_fuzz_dep_.CoverTab[84439]++
								if !ok || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1053
			_go_fuzz_dep_.CoverTab[84447]++
//line /usr/local/go/src/encoding/gob/decode.go:1053
			return wire.ArrayT == nil
//line /usr/local/go/src/encoding/gob/decode.go:1053
			// _ = "end of CoverTab[84447]"
//line /usr/local/go/src/encoding/gob/decode.go:1053
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1053
			_go_fuzz_dep_.CoverTab[84448]++
									return false
//line /usr/local/go/src/encoding/gob/decode.go:1054
			// _ = "end of CoverTab[84448]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1055
			_go_fuzz_dep_.CoverTab[84449]++
//line /usr/local/go/src/encoding/gob/decode.go:1055
			// _ = "end of CoverTab[84449]"
//line /usr/local/go/src/encoding/gob/decode.go:1055
		}
//line /usr/local/go/src/encoding/gob/decode.go:1055
		// _ = "end of CoverTab[84439]"
//line /usr/local/go/src/encoding/gob/decode.go:1055
		_go_fuzz_dep_.CoverTab[84440]++
								array := wire.ArrayT
								return t.Len() == array.Len && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1057
			_go_fuzz_dep_.CoverTab[84450]++
//line /usr/local/go/src/encoding/gob/decode.go:1057
			return dec.compatibleType(t.Elem(), array.Elem, inProgress)
//line /usr/local/go/src/encoding/gob/decode.go:1057
			// _ = "end of CoverTab[84450]"
//line /usr/local/go/src/encoding/gob/decode.go:1057
		}()
//line /usr/local/go/src/encoding/gob/decode.go:1057
		// _ = "end of CoverTab[84440]"
	case reflect.Map:
//line /usr/local/go/src/encoding/gob/decode.go:1058
		_go_fuzz_dep_.CoverTab[84441]++
								if !ok || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1059
			_go_fuzz_dep_.CoverTab[84451]++
//line /usr/local/go/src/encoding/gob/decode.go:1059
			return wire.MapT == nil
//line /usr/local/go/src/encoding/gob/decode.go:1059
			// _ = "end of CoverTab[84451]"
//line /usr/local/go/src/encoding/gob/decode.go:1059
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1059
			_go_fuzz_dep_.CoverTab[84452]++
									return false
//line /usr/local/go/src/encoding/gob/decode.go:1060
			// _ = "end of CoverTab[84452]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1061
			_go_fuzz_dep_.CoverTab[84453]++
//line /usr/local/go/src/encoding/gob/decode.go:1061
			// _ = "end of CoverTab[84453]"
//line /usr/local/go/src/encoding/gob/decode.go:1061
		}
//line /usr/local/go/src/encoding/gob/decode.go:1061
		// _ = "end of CoverTab[84441]"
//line /usr/local/go/src/encoding/gob/decode.go:1061
		_go_fuzz_dep_.CoverTab[84442]++
								MapType := wire.MapT
								return dec.compatibleType(t.Key(), MapType.Key, inProgress) && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1063
			_go_fuzz_dep_.CoverTab[84454]++
//line /usr/local/go/src/encoding/gob/decode.go:1063
			return dec.compatibleType(t.Elem(), MapType.Elem, inProgress)
//line /usr/local/go/src/encoding/gob/decode.go:1063
			// _ = "end of CoverTab[84454]"
//line /usr/local/go/src/encoding/gob/decode.go:1063
		}()
//line /usr/local/go/src/encoding/gob/decode.go:1063
		// _ = "end of CoverTab[84442]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/gob/decode.go:1064
		_go_fuzz_dep_.CoverTab[84443]++

								if t.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/gob/decode.go:1066
			_go_fuzz_dep_.CoverTab[84455]++
									return fw == tBytes
//line /usr/local/go/src/encoding/gob/decode.go:1067
			// _ = "end of CoverTab[84455]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1068
			_go_fuzz_dep_.CoverTab[84456]++
//line /usr/local/go/src/encoding/gob/decode.go:1068
			// _ = "end of CoverTab[84456]"
//line /usr/local/go/src/encoding/gob/decode.go:1068
		}
//line /usr/local/go/src/encoding/gob/decode.go:1068
		// _ = "end of CoverTab[84443]"
//line /usr/local/go/src/encoding/gob/decode.go:1068
		_go_fuzz_dep_.CoverTab[84444]++

								var sw *sliceType
								if tt, ok := builtinIdToType[fw]; ok {
//line /usr/local/go/src/encoding/gob/decode.go:1071
			_go_fuzz_dep_.CoverTab[84457]++
									sw, _ = tt.(*sliceType)
//line /usr/local/go/src/encoding/gob/decode.go:1072
			// _ = "end of CoverTab[84457]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1073
			_go_fuzz_dep_.CoverTab[84458]++
//line /usr/local/go/src/encoding/gob/decode.go:1073
			if wire != nil {
//line /usr/local/go/src/encoding/gob/decode.go:1073
				_go_fuzz_dep_.CoverTab[84459]++
										sw = wire.SliceT
//line /usr/local/go/src/encoding/gob/decode.go:1074
				// _ = "end of CoverTab[84459]"
			} else {
//line /usr/local/go/src/encoding/gob/decode.go:1075
				_go_fuzz_dep_.CoverTab[84460]++
//line /usr/local/go/src/encoding/gob/decode.go:1075
				// _ = "end of CoverTab[84460]"
//line /usr/local/go/src/encoding/gob/decode.go:1075
			}
//line /usr/local/go/src/encoding/gob/decode.go:1075
			// _ = "end of CoverTab[84458]"
//line /usr/local/go/src/encoding/gob/decode.go:1075
		}
//line /usr/local/go/src/encoding/gob/decode.go:1075
		// _ = "end of CoverTab[84444]"
//line /usr/local/go/src/encoding/gob/decode.go:1075
		_go_fuzz_dep_.CoverTab[84445]++
								elem := userType(t.Elem()).base
								return sw != nil && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1077
			_go_fuzz_dep_.CoverTab[84461]++
//line /usr/local/go/src/encoding/gob/decode.go:1077
			return dec.compatibleType(elem, sw.Elem, inProgress)
//line /usr/local/go/src/encoding/gob/decode.go:1077
			// _ = "end of CoverTab[84461]"
//line /usr/local/go/src/encoding/gob/decode.go:1077
		}()
//line /usr/local/go/src/encoding/gob/decode.go:1077
		// _ = "end of CoverTab[84445]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/gob/decode.go:1078
		_go_fuzz_dep_.CoverTab[84446]++
								return true
//line /usr/local/go/src/encoding/gob/decode.go:1079
		// _ = "end of CoverTab[84446]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:1080
	// _ = "end of CoverTab[84419]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1084
func (dec *Decoder) typeString(remoteId typeId) string {
//line /usr/local/go/src/encoding/gob/decode.go:1084
	_go_fuzz_dep_.CoverTab[84462]++
							typeLock.Lock()
							defer typeLock.Unlock()
							if t := idToType[remoteId]; t != nil {
//line /usr/local/go/src/encoding/gob/decode.go:1087
		_go_fuzz_dep_.CoverTab[84464]++

								return t.string()
//line /usr/local/go/src/encoding/gob/decode.go:1089
		// _ = "end of CoverTab[84464]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1090
		_go_fuzz_dep_.CoverTab[84465]++
//line /usr/local/go/src/encoding/gob/decode.go:1090
		// _ = "end of CoverTab[84465]"
//line /usr/local/go/src/encoding/gob/decode.go:1090
	}
//line /usr/local/go/src/encoding/gob/decode.go:1090
	// _ = "end of CoverTab[84462]"
//line /usr/local/go/src/encoding/gob/decode.go:1090
	_go_fuzz_dep_.CoverTab[84463]++
							return dec.wireType[remoteId].string()
//line /usr/local/go/src/encoding/gob/decode.go:1091
	// _ = "end of CoverTab[84463]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1096
func (dec *Decoder) compileSingle(remoteId typeId, ut *userTypeInfo) (engine *decEngine, err error) {
//line /usr/local/go/src/encoding/gob/decode.go:1096
	_go_fuzz_dep_.CoverTab[84466]++
							rt := ut.user
							engine = new(decEngine)
							engine.instr = make([]decInstr, 1)
							name := rt.String()
							if !dec.compatibleType(rt, remoteId, make(map[reflect.Type]typeId)) {
//line /usr/local/go/src/encoding/gob/decode.go:1101
		_go_fuzz_dep_.CoverTab[84468]++
								remoteType := dec.typeString(remoteId)

								if ut.base.Kind() == reflect.Interface && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1104
			_go_fuzz_dep_.CoverTab[84470]++
//line /usr/local/go/src/encoding/gob/decode.go:1104
			return remoteId != tInterface
//line /usr/local/go/src/encoding/gob/decode.go:1104
			// _ = "end of CoverTab[84470]"
//line /usr/local/go/src/encoding/gob/decode.go:1104
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1104
			_go_fuzz_dep_.CoverTab[84471]++
									return nil, errors.New("gob: local interface type " + name + " can only be decoded from remote interface type; received concrete type " + remoteType)
//line /usr/local/go/src/encoding/gob/decode.go:1105
			// _ = "end of CoverTab[84471]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1106
			_go_fuzz_dep_.CoverTab[84472]++
//line /usr/local/go/src/encoding/gob/decode.go:1106
			// _ = "end of CoverTab[84472]"
//line /usr/local/go/src/encoding/gob/decode.go:1106
		}
//line /usr/local/go/src/encoding/gob/decode.go:1106
		// _ = "end of CoverTab[84468]"
//line /usr/local/go/src/encoding/gob/decode.go:1106
		_go_fuzz_dep_.CoverTab[84469]++
								return nil, errors.New("gob: decoding into local type " + name + ", received remote type " + remoteType)
//line /usr/local/go/src/encoding/gob/decode.go:1107
		// _ = "end of CoverTab[84469]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1108
		_go_fuzz_dep_.CoverTab[84473]++
//line /usr/local/go/src/encoding/gob/decode.go:1108
		// _ = "end of CoverTab[84473]"
//line /usr/local/go/src/encoding/gob/decode.go:1108
	}
//line /usr/local/go/src/encoding/gob/decode.go:1108
	// _ = "end of CoverTab[84466]"
//line /usr/local/go/src/encoding/gob/decode.go:1108
	_go_fuzz_dep_.CoverTab[84467]++
							op := dec.decOpFor(remoteId, rt, name, make(map[reflect.Type]*decOp))
							ovfl := errors.New(`value for "` + name + `" out of range`)
							engine.instr[singletonField] = decInstr{*op, singletonField, nil, ovfl}
							engine.numInstr = 1
							return
//line /usr/local/go/src/encoding/gob/decode.go:1113
	// _ = "end of CoverTab[84467]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1117
func (dec *Decoder) compileIgnoreSingle(remoteId typeId) *decEngine {
//line /usr/local/go/src/encoding/gob/decode.go:1117
	_go_fuzz_dep_.CoverTab[84474]++
							engine := new(decEngine)
							engine.instr = make([]decInstr, 1)
							op := dec.decIgnoreOpFor(remoteId, make(map[typeId]*decOp), 0)
							ovfl := overflow(dec.typeString(remoteId))
							engine.instr[0] = decInstr{*op, 0, nil, ovfl}
							engine.numInstr = 1
							return engine
//line /usr/local/go/src/encoding/gob/decode.go:1124
	// _ = "end of CoverTab[84474]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1129
func (dec *Decoder) compileDec(remoteId typeId, ut *userTypeInfo) (engine *decEngine, err error) {
//line /usr/local/go/src/encoding/gob/decode.go:1129
	_go_fuzz_dep_.CoverTab[84475]++
							defer catchError(&err)
							rt := ut.base
							srt := rt
							if srt.Kind() != reflect.Struct || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1133
		_go_fuzz_dep_.CoverTab[84480]++
//line /usr/local/go/src/encoding/gob/decode.go:1133
		return ut.externalDec != 0
//line /usr/local/go/src/encoding/gob/decode.go:1133
		// _ = "end of CoverTab[84480]"
//line /usr/local/go/src/encoding/gob/decode.go:1133
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:1133
		_go_fuzz_dep_.CoverTab[84481]++
								return dec.compileSingle(remoteId, ut)
//line /usr/local/go/src/encoding/gob/decode.go:1134
		// _ = "end of CoverTab[84481]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1135
		_go_fuzz_dep_.CoverTab[84482]++
//line /usr/local/go/src/encoding/gob/decode.go:1135
		// _ = "end of CoverTab[84482]"
//line /usr/local/go/src/encoding/gob/decode.go:1135
	}
//line /usr/local/go/src/encoding/gob/decode.go:1135
	// _ = "end of CoverTab[84475]"
//line /usr/local/go/src/encoding/gob/decode.go:1135
	_go_fuzz_dep_.CoverTab[84476]++
							var wireStruct *structType

//line /usr/local/go/src/encoding/gob/decode.go:1139
	if t, ok := builtinIdToType[remoteId]; ok {
//line /usr/local/go/src/encoding/gob/decode.go:1139
		_go_fuzz_dep_.CoverTab[84483]++
								wireStruct, _ = t.(*structType)
//line /usr/local/go/src/encoding/gob/decode.go:1140
		// _ = "end of CoverTab[84483]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1141
		_go_fuzz_dep_.CoverTab[84484]++
								wire := dec.wireType[remoteId]
								if wire == nil {
//line /usr/local/go/src/encoding/gob/decode.go:1143
			_go_fuzz_dep_.CoverTab[84486]++
									error_(errBadType)
//line /usr/local/go/src/encoding/gob/decode.go:1144
			// _ = "end of CoverTab[84486]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1145
			_go_fuzz_dep_.CoverTab[84487]++
//line /usr/local/go/src/encoding/gob/decode.go:1145
			// _ = "end of CoverTab[84487]"
//line /usr/local/go/src/encoding/gob/decode.go:1145
		}
//line /usr/local/go/src/encoding/gob/decode.go:1145
		// _ = "end of CoverTab[84484]"
//line /usr/local/go/src/encoding/gob/decode.go:1145
		_go_fuzz_dep_.CoverTab[84485]++
								wireStruct = wire.StructT
//line /usr/local/go/src/encoding/gob/decode.go:1146
		// _ = "end of CoverTab[84485]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:1147
	// _ = "end of CoverTab[84476]"
//line /usr/local/go/src/encoding/gob/decode.go:1147
	_go_fuzz_dep_.CoverTab[84477]++
							if wireStruct == nil {
//line /usr/local/go/src/encoding/gob/decode.go:1148
		_go_fuzz_dep_.CoverTab[84488]++
								errorf("type mismatch in decoder: want struct type %s; got non-struct", rt)
//line /usr/local/go/src/encoding/gob/decode.go:1149
		// _ = "end of CoverTab[84488]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1150
		_go_fuzz_dep_.CoverTab[84489]++
//line /usr/local/go/src/encoding/gob/decode.go:1150
		// _ = "end of CoverTab[84489]"
//line /usr/local/go/src/encoding/gob/decode.go:1150
	}
//line /usr/local/go/src/encoding/gob/decode.go:1150
	// _ = "end of CoverTab[84477]"
//line /usr/local/go/src/encoding/gob/decode.go:1150
	_go_fuzz_dep_.CoverTab[84478]++
							engine = new(decEngine)
							engine.instr = make([]decInstr, len(wireStruct.Field))
							seen := make(map[reflect.Type]*decOp)

							for fieldnum := 0; fieldnum < len(wireStruct.Field); fieldnum++ {
//line /usr/local/go/src/encoding/gob/decode.go:1155
		_go_fuzz_dep_.CoverTab[84490]++
								wireField := wireStruct.Field[fieldnum]
								if wireField.Name == "" {
//line /usr/local/go/src/encoding/gob/decode.go:1157
			_go_fuzz_dep_.CoverTab[84494]++
									errorf("empty name for remote field of type %s", wireStruct.Name)
//line /usr/local/go/src/encoding/gob/decode.go:1158
			// _ = "end of CoverTab[84494]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1159
			_go_fuzz_dep_.CoverTab[84495]++
//line /usr/local/go/src/encoding/gob/decode.go:1159
			// _ = "end of CoverTab[84495]"
//line /usr/local/go/src/encoding/gob/decode.go:1159
		}
//line /usr/local/go/src/encoding/gob/decode.go:1159
		// _ = "end of CoverTab[84490]"
//line /usr/local/go/src/encoding/gob/decode.go:1159
		_go_fuzz_dep_.CoverTab[84491]++
								ovfl := overflow(wireField.Name)

								localField, present := srt.FieldByName(wireField.Name)

								if !present || func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1164
			_go_fuzz_dep_.CoverTab[84496]++
//line /usr/local/go/src/encoding/gob/decode.go:1164
			return !isExported(wireField.Name)
//line /usr/local/go/src/encoding/gob/decode.go:1164
			// _ = "end of CoverTab[84496]"
//line /usr/local/go/src/encoding/gob/decode.go:1164
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1164
			_go_fuzz_dep_.CoverTab[84497]++
									op := dec.decIgnoreOpFor(wireField.Id, make(map[typeId]*decOp), 0)
									engine.instr[fieldnum] = decInstr{*op, fieldnum, nil, ovfl}
									continue
//line /usr/local/go/src/encoding/gob/decode.go:1167
			// _ = "end of CoverTab[84497]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1168
			_go_fuzz_dep_.CoverTab[84498]++
//line /usr/local/go/src/encoding/gob/decode.go:1168
			// _ = "end of CoverTab[84498]"
//line /usr/local/go/src/encoding/gob/decode.go:1168
		}
//line /usr/local/go/src/encoding/gob/decode.go:1168
		// _ = "end of CoverTab[84491]"
//line /usr/local/go/src/encoding/gob/decode.go:1168
		_go_fuzz_dep_.CoverTab[84492]++
								if !dec.compatibleType(localField.Type, wireField.Id, make(map[reflect.Type]typeId)) {
//line /usr/local/go/src/encoding/gob/decode.go:1169
			_go_fuzz_dep_.CoverTab[84499]++
									errorf("wrong type (%s) for received field %s.%s", localField.Type, wireStruct.Name, wireField.Name)
//line /usr/local/go/src/encoding/gob/decode.go:1170
			// _ = "end of CoverTab[84499]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1171
			_go_fuzz_dep_.CoverTab[84500]++
//line /usr/local/go/src/encoding/gob/decode.go:1171
			// _ = "end of CoverTab[84500]"
//line /usr/local/go/src/encoding/gob/decode.go:1171
		}
//line /usr/local/go/src/encoding/gob/decode.go:1171
		// _ = "end of CoverTab[84492]"
//line /usr/local/go/src/encoding/gob/decode.go:1171
		_go_fuzz_dep_.CoverTab[84493]++
								op := dec.decOpFor(wireField.Id, localField.Type, localField.Name, seen)
								engine.instr[fieldnum] = decInstr{*op, fieldnum, localField.Index, ovfl}
								engine.numInstr++
//line /usr/local/go/src/encoding/gob/decode.go:1174
		// _ = "end of CoverTab[84493]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:1175
	// _ = "end of CoverTab[84478]"
//line /usr/local/go/src/encoding/gob/decode.go:1175
	_go_fuzz_dep_.CoverTab[84479]++
							return
//line /usr/local/go/src/encoding/gob/decode.go:1176
	// _ = "end of CoverTab[84479]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1180
func (dec *Decoder) getDecEnginePtr(remoteId typeId, ut *userTypeInfo) (enginePtr **decEngine, err error) {
//line /usr/local/go/src/encoding/gob/decode.go:1180
	_go_fuzz_dep_.CoverTab[84501]++
							rt := ut.user
							decoderMap, ok := dec.decoderCache[rt]
							if !ok {
//line /usr/local/go/src/encoding/gob/decode.go:1183
		_go_fuzz_dep_.CoverTab[84504]++
								decoderMap = make(map[typeId]**decEngine)
								dec.decoderCache[rt] = decoderMap
//line /usr/local/go/src/encoding/gob/decode.go:1185
		// _ = "end of CoverTab[84504]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1186
		_go_fuzz_dep_.CoverTab[84505]++
//line /usr/local/go/src/encoding/gob/decode.go:1186
		// _ = "end of CoverTab[84505]"
//line /usr/local/go/src/encoding/gob/decode.go:1186
	}
//line /usr/local/go/src/encoding/gob/decode.go:1186
	// _ = "end of CoverTab[84501]"
//line /usr/local/go/src/encoding/gob/decode.go:1186
	_go_fuzz_dep_.CoverTab[84502]++
							if enginePtr, ok = decoderMap[remoteId]; !ok {
//line /usr/local/go/src/encoding/gob/decode.go:1187
		_go_fuzz_dep_.CoverTab[84506]++

								enginePtr = new(*decEngine)
								decoderMap[remoteId] = enginePtr
								*enginePtr, err = dec.compileDec(remoteId, ut)
								if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:1192
			_go_fuzz_dep_.CoverTab[84507]++
									delete(decoderMap, remoteId)
//line /usr/local/go/src/encoding/gob/decode.go:1193
			// _ = "end of CoverTab[84507]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1194
			_go_fuzz_dep_.CoverTab[84508]++
//line /usr/local/go/src/encoding/gob/decode.go:1194
			// _ = "end of CoverTab[84508]"
//line /usr/local/go/src/encoding/gob/decode.go:1194
		}
//line /usr/local/go/src/encoding/gob/decode.go:1194
		// _ = "end of CoverTab[84506]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1195
		_go_fuzz_dep_.CoverTab[84509]++
//line /usr/local/go/src/encoding/gob/decode.go:1195
		// _ = "end of CoverTab[84509]"
//line /usr/local/go/src/encoding/gob/decode.go:1195
	}
//line /usr/local/go/src/encoding/gob/decode.go:1195
	// _ = "end of CoverTab[84502]"
//line /usr/local/go/src/encoding/gob/decode.go:1195
	_go_fuzz_dep_.CoverTab[84503]++
							return
//line /usr/local/go/src/encoding/gob/decode.go:1196
	// _ = "end of CoverTab[84503]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1200
type emptyStruct struct{}

var emptyStructType = reflect.TypeOf(emptyStruct{})

//line /usr/local/go/src/encoding/gob/decode.go:1205
func (dec *Decoder) getIgnoreEnginePtr(wireId typeId) (enginePtr **decEngine, err error) {
//line /usr/local/go/src/encoding/gob/decode.go:1205
	_go_fuzz_dep_.CoverTab[84510]++
							var ok bool
							if enginePtr, ok = dec.ignorerCache[wireId]; !ok {
//line /usr/local/go/src/encoding/gob/decode.go:1207
		_go_fuzz_dep_.CoverTab[84512]++

								enginePtr = new(*decEngine)
								dec.ignorerCache[wireId] = enginePtr
								wire := dec.wireType[wireId]
								if wire != nil && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1212
			_go_fuzz_dep_.CoverTab[84514]++
//line /usr/local/go/src/encoding/gob/decode.go:1212
			return wire.StructT != nil
//line /usr/local/go/src/encoding/gob/decode.go:1212
			// _ = "end of CoverTab[84514]"
//line /usr/local/go/src/encoding/gob/decode.go:1212
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1212
			_go_fuzz_dep_.CoverTab[84515]++
									*enginePtr, err = dec.compileDec(wireId, userType(emptyStructType))
//line /usr/local/go/src/encoding/gob/decode.go:1213
			// _ = "end of CoverTab[84515]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1214
			_go_fuzz_dep_.CoverTab[84516]++
									*enginePtr = dec.compileIgnoreSingle(wireId)
//line /usr/local/go/src/encoding/gob/decode.go:1215
			// _ = "end of CoverTab[84516]"
		}
//line /usr/local/go/src/encoding/gob/decode.go:1216
		// _ = "end of CoverTab[84512]"
//line /usr/local/go/src/encoding/gob/decode.go:1216
		_go_fuzz_dep_.CoverTab[84513]++
								if err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:1217
			_go_fuzz_dep_.CoverTab[84517]++
									delete(dec.ignorerCache, wireId)
//line /usr/local/go/src/encoding/gob/decode.go:1218
			// _ = "end of CoverTab[84517]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1219
			_go_fuzz_dep_.CoverTab[84518]++
//line /usr/local/go/src/encoding/gob/decode.go:1219
			// _ = "end of CoverTab[84518]"
//line /usr/local/go/src/encoding/gob/decode.go:1219
		}
//line /usr/local/go/src/encoding/gob/decode.go:1219
		// _ = "end of CoverTab[84513]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1220
		_go_fuzz_dep_.CoverTab[84519]++
//line /usr/local/go/src/encoding/gob/decode.go:1220
		// _ = "end of CoverTab[84519]"
//line /usr/local/go/src/encoding/gob/decode.go:1220
	}
//line /usr/local/go/src/encoding/gob/decode.go:1220
	// _ = "end of CoverTab[84510]"
//line /usr/local/go/src/encoding/gob/decode.go:1220
	_go_fuzz_dep_.CoverTab[84511]++
							return
//line /usr/local/go/src/encoding/gob/decode.go:1221
	// _ = "end of CoverTab[84511]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1225
func (dec *Decoder) decodeValue(wireId typeId, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/decode.go:1225
	_go_fuzz_dep_.CoverTab[84520]++
							defer catchError(&dec.err)

							if !value.IsValid() {
//line /usr/local/go/src/encoding/gob/decode.go:1228
		_go_fuzz_dep_.CoverTab[84523]++
								dec.decodeIgnoredValue(wireId)
								return
//line /usr/local/go/src/encoding/gob/decode.go:1230
		// _ = "end of CoverTab[84523]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1231
		_go_fuzz_dep_.CoverTab[84524]++
//line /usr/local/go/src/encoding/gob/decode.go:1231
		// _ = "end of CoverTab[84524]"
//line /usr/local/go/src/encoding/gob/decode.go:1231
	}
//line /usr/local/go/src/encoding/gob/decode.go:1231
	// _ = "end of CoverTab[84520]"
//line /usr/local/go/src/encoding/gob/decode.go:1231
	_go_fuzz_dep_.CoverTab[84521]++

							ut := userType(value.Type())
							base := ut.base
							var enginePtr **decEngine
							enginePtr, dec.err = dec.getDecEnginePtr(wireId, ut)
							if dec.err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:1237
		_go_fuzz_dep_.CoverTab[84525]++
								return
//line /usr/local/go/src/encoding/gob/decode.go:1238
		// _ = "end of CoverTab[84525]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1239
		_go_fuzz_dep_.CoverTab[84526]++
//line /usr/local/go/src/encoding/gob/decode.go:1239
		// _ = "end of CoverTab[84526]"
//line /usr/local/go/src/encoding/gob/decode.go:1239
	}
//line /usr/local/go/src/encoding/gob/decode.go:1239
	// _ = "end of CoverTab[84521]"
//line /usr/local/go/src/encoding/gob/decode.go:1239
	_go_fuzz_dep_.CoverTab[84522]++
							value = decAlloc(value)
							engine := *enginePtr
							if st := base; st.Kind() == reflect.Struct && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1242
		_go_fuzz_dep_.CoverTab[84527]++
//line /usr/local/go/src/encoding/gob/decode.go:1242
		return ut.externalDec == 0
//line /usr/local/go/src/encoding/gob/decode.go:1242
		// _ = "end of CoverTab[84527]"
//line /usr/local/go/src/encoding/gob/decode.go:1242
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:1242
		_go_fuzz_dep_.CoverTab[84528]++
								wt := dec.wireType[wireId]
								if engine.numInstr == 0 && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1244
			_go_fuzz_dep_.CoverTab[84530]++
//line /usr/local/go/src/encoding/gob/decode.go:1244
			return st.NumField() > 0
//line /usr/local/go/src/encoding/gob/decode.go:1244
			// _ = "end of CoverTab[84530]"
//line /usr/local/go/src/encoding/gob/decode.go:1244
		}() && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1244
			_go_fuzz_dep_.CoverTab[84531]++
//line /usr/local/go/src/encoding/gob/decode.go:1244
			return wt != nil
									// _ = "end of CoverTab[84531]"
//line /usr/local/go/src/encoding/gob/decode.go:1245
		}() && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1245
			_go_fuzz_dep_.CoverTab[84532]++
//line /usr/local/go/src/encoding/gob/decode.go:1245
			return len(wt.StructT.Field) > 0
//line /usr/local/go/src/encoding/gob/decode.go:1245
			// _ = "end of CoverTab[84532]"
//line /usr/local/go/src/encoding/gob/decode.go:1245
		}() {
//line /usr/local/go/src/encoding/gob/decode.go:1245
			_go_fuzz_dep_.CoverTab[84533]++
									name := base.Name()
									errorf("type mismatch: no fields matched compiling decoder for %s", name)
//line /usr/local/go/src/encoding/gob/decode.go:1247
			// _ = "end of CoverTab[84533]"
		} else {
//line /usr/local/go/src/encoding/gob/decode.go:1248
			_go_fuzz_dep_.CoverTab[84534]++
//line /usr/local/go/src/encoding/gob/decode.go:1248
			// _ = "end of CoverTab[84534]"
//line /usr/local/go/src/encoding/gob/decode.go:1248
		}
//line /usr/local/go/src/encoding/gob/decode.go:1248
		// _ = "end of CoverTab[84528]"
//line /usr/local/go/src/encoding/gob/decode.go:1248
		_go_fuzz_dep_.CoverTab[84529]++
								dec.decodeStruct(engine, value)
//line /usr/local/go/src/encoding/gob/decode.go:1249
		// _ = "end of CoverTab[84529]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1250
		_go_fuzz_dep_.CoverTab[84535]++
								dec.decodeSingle(engine, value)
//line /usr/local/go/src/encoding/gob/decode.go:1251
		// _ = "end of CoverTab[84535]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:1252
	// _ = "end of CoverTab[84522]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1256
func (dec *Decoder) decodeIgnoredValue(wireId typeId) {
//line /usr/local/go/src/encoding/gob/decode.go:1256
	_go_fuzz_dep_.CoverTab[84536]++
							var enginePtr **decEngine
							enginePtr, dec.err = dec.getIgnoreEnginePtr(wireId)
							if dec.err != nil {
//line /usr/local/go/src/encoding/gob/decode.go:1259
		_go_fuzz_dep_.CoverTab[84538]++
								return
//line /usr/local/go/src/encoding/gob/decode.go:1260
		// _ = "end of CoverTab[84538]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1261
		_go_fuzz_dep_.CoverTab[84539]++
//line /usr/local/go/src/encoding/gob/decode.go:1261
		// _ = "end of CoverTab[84539]"
//line /usr/local/go/src/encoding/gob/decode.go:1261
	}
//line /usr/local/go/src/encoding/gob/decode.go:1261
	// _ = "end of CoverTab[84536]"
//line /usr/local/go/src/encoding/gob/decode.go:1261
	_go_fuzz_dep_.CoverTab[84537]++
							wire := dec.wireType[wireId]
							if wire != nil && func() bool {
//line /usr/local/go/src/encoding/gob/decode.go:1263
		_go_fuzz_dep_.CoverTab[84540]++
//line /usr/local/go/src/encoding/gob/decode.go:1263
		return wire.StructT != nil
//line /usr/local/go/src/encoding/gob/decode.go:1263
		// _ = "end of CoverTab[84540]"
//line /usr/local/go/src/encoding/gob/decode.go:1263
	}() {
//line /usr/local/go/src/encoding/gob/decode.go:1263
		_go_fuzz_dep_.CoverTab[84541]++
								dec.ignoreStruct(*enginePtr)
//line /usr/local/go/src/encoding/gob/decode.go:1264
		// _ = "end of CoverTab[84541]"
	} else {
//line /usr/local/go/src/encoding/gob/decode.go:1265
		_go_fuzz_dep_.CoverTab[84542]++
								dec.ignoreSingle(*enginePtr)
//line /usr/local/go/src/encoding/gob/decode.go:1266
		// _ = "end of CoverTab[84542]"
	}
//line /usr/local/go/src/encoding/gob/decode.go:1267
	// _ = "end of CoverTab[84537]"
}

const (
	intBits		= 32 << (^uint(0) >> 63)
	uintptrBits	= 32 << (^uintptr(0) >> 63)
)

func init() {
	var iop, uop decOp
	switch intBits {
	case 32:
		iop = decInt32
		uop = decUint32
	case 64:
		iop = decInt64
		uop = decUint64
	default:
		panic("gob: unknown size of int/uint")
	}
							decOpTable[reflect.Int] = iop
							decOpTable[reflect.Uint] = uop

//line /usr/local/go/src/encoding/gob/decode.go:1291
	switch uintptrBits {
	case 32:
		uop = decUint32
	case 64:
		uop = decUint64
	default:
		panic("gob: unknown size of uintptr")
	}
	decOpTable[reflect.Uintptr] = uop
}

//line /usr/local/go/src/encoding/gob/decode.go:1306
func allocValue(t reflect.Type) reflect.Value {
//line /usr/local/go/src/encoding/gob/decode.go:1306
	_go_fuzz_dep_.CoverTab[84543]++
							return reflect.New(t).Elem()
//line /usr/local/go/src/encoding/gob/decode.go:1307
	// _ = "end of CoverTab[84543]"
}

//line /usr/local/go/src/encoding/gob/decode.go:1308
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/gob/decode.go:1308
var _ = _go_fuzz_dep_.CoverTab
