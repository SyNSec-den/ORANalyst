// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run encgen.go -output enc_helpers.go

//line /usr/local/go/src/encoding/gob/encode.go:5
//go:generate go run encgen.go -output enc_helpers.go

package gob

//line /usr/local/go/src/encoding/gob/encode.go:7
import (
//line /usr/local/go/src/encoding/gob/encode.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/gob/encode.go:7
)
//line /usr/local/go/src/encoding/gob/encode.go:7
import (
//line /usr/local/go/src/encoding/gob/encode.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/gob/encode.go:7
)

import (
	"encoding"
	"encoding/binary"
	"math"
	"math/bits"
	"reflect"
	"sync"
)

const uint64Size = 8

type encHelper func(state *encoderState, v reflect.Value) bool

//line /usr/local/go/src/encoding/gob/encode.go:26
type encoderState struct {
	enc		*Encoder
	b		*encBuffer
	sendZero	bool
	fieldnum	int
	buf		[1 + uint64Size]byte
	next		*encoderState
}

//line /usr/local/go/src/encoding/gob/encode.go:37
type encBuffer struct {
	data	[]byte
	scratch	[64]byte
}

var encBufferPool = sync.Pool{
	New: func() any {
//line /usr/local/go/src/encoding/gob/encode.go:43
		_go_fuzz_dep_.CoverTab[84834]++
								e := new(encBuffer)
								e.data = e.scratch[0:0]
								return e
//line /usr/local/go/src/encoding/gob/encode.go:46
		// _ = "end of CoverTab[84834]"
	},
}

func (e *encBuffer) writeByte(c byte) {
//line /usr/local/go/src/encoding/gob/encode.go:50
	_go_fuzz_dep_.CoverTab[84835]++
							e.data = append(e.data, c)
//line /usr/local/go/src/encoding/gob/encode.go:51
	// _ = "end of CoverTab[84835]"
}

func (e *encBuffer) Write(p []byte) (int, error) {
//line /usr/local/go/src/encoding/gob/encode.go:54
	_go_fuzz_dep_.CoverTab[84836]++
							e.data = append(e.data, p...)
							return len(p), nil
//line /usr/local/go/src/encoding/gob/encode.go:56
	// _ = "end of CoverTab[84836]"
}

func (e *encBuffer) WriteString(s string) {
//line /usr/local/go/src/encoding/gob/encode.go:59
	_go_fuzz_dep_.CoverTab[84837]++
							e.data = append(e.data, s...)
//line /usr/local/go/src/encoding/gob/encode.go:60
	// _ = "end of CoverTab[84837]"
}

func (e *encBuffer) Len() int {
//line /usr/local/go/src/encoding/gob/encode.go:63
	_go_fuzz_dep_.CoverTab[84838]++
							return len(e.data)
//line /usr/local/go/src/encoding/gob/encode.go:64
	// _ = "end of CoverTab[84838]"
}

func (e *encBuffer) Bytes() []byte {
//line /usr/local/go/src/encoding/gob/encode.go:67
	_go_fuzz_dep_.CoverTab[84839]++
							return e.data
//line /usr/local/go/src/encoding/gob/encode.go:68
	// _ = "end of CoverTab[84839]"
}

func (e *encBuffer) Reset() {
//line /usr/local/go/src/encoding/gob/encode.go:71
	_go_fuzz_dep_.CoverTab[84840]++
							if len(e.data) >= tooBig {
//line /usr/local/go/src/encoding/gob/encode.go:72
		_go_fuzz_dep_.CoverTab[84841]++
								e.data = e.scratch[0:0]
//line /usr/local/go/src/encoding/gob/encode.go:73
		// _ = "end of CoverTab[84841]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:74
		_go_fuzz_dep_.CoverTab[84842]++
								e.data = e.data[0:0]
//line /usr/local/go/src/encoding/gob/encode.go:75
		// _ = "end of CoverTab[84842]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:76
	// _ = "end of CoverTab[84840]"
}

func (enc *Encoder) newEncoderState(b *encBuffer) *encoderState {
//line /usr/local/go/src/encoding/gob/encode.go:79
	_go_fuzz_dep_.CoverTab[84843]++
							e := enc.freeList
							if e == nil {
//line /usr/local/go/src/encoding/gob/encode.go:81
		_go_fuzz_dep_.CoverTab[84846]++
								e = new(encoderState)
								e.enc = enc
//line /usr/local/go/src/encoding/gob/encode.go:83
		// _ = "end of CoverTab[84846]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:84
		_go_fuzz_dep_.CoverTab[84847]++
								enc.freeList = e.next
//line /usr/local/go/src/encoding/gob/encode.go:85
		// _ = "end of CoverTab[84847]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:86
	// _ = "end of CoverTab[84843]"
//line /usr/local/go/src/encoding/gob/encode.go:86
	_go_fuzz_dep_.CoverTab[84844]++
							e.sendZero = false
							e.fieldnum = 0
							e.b = b
							if len(b.data) == 0 {
//line /usr/local/go/src/encoding/gob/encode.go:90
		_go_fuzz_dep_.CoverTab[84848]++
								b.data = b.scratch[0:0]
//line /usr/local/go/src/encoding/gob/encode.go:91
		// _ = "end of CoverTab[84848]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:92
		_go_fuzz_dep_.CoverTab[84849]++
//line /usr/local/go/src/encoding/gob/encode.go:92
		// _ = "end of CoverTab[84849]"
//line /usr/local/go/src/encoding/gob/encode.go:92
	}
//line /usr/local/go/src/encoding/gob/encode.go:92
	// _ = "end of CoverTab[84844]"
//line /usr/local/go/src/encoding/gob/encode.go:92
	_go_fuzz_dep_.CoverTab[84845]++
							return e
//line /usr/local/go/src/encoding/gob/encode.go:93
	// _ = "end of CoverTab[84845]"
}

func (enc *Encoder) freeEncoderState(e *encoderState) {
//line /usr/local/go/src/encoding/gob/encode.go:96
	_go_fuzz_dep_.CoverTab[84850]++
							e.next = enc.freeList
							enc.freeList = e
//line /usr/local/go/src/encoding/gob/encode.go:98
	// _ = "end of CoverTab[84850]"
}

//line /usr/local/go/src/encoding/gob/encode.go:107
func (state *encoderState) encodeUint(x uint64) {
//line /usr/local/go/src/encoding/gob/encode.go:107
	_go_fuzz_dep_.CoverTab[84851]++
							if x <= 0x7F {
//line /usr/local/go/src/encoding/gob/encode.go:108
		_go_fuzz_dep_.CoverTab[84853]++
								state.b.writeByte(uint8(x))
								return
//line /usr/local/go/src/encoding/gob/encode.go:110
		// _ = "end of CoverTab[84853]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:111
		_go_fuzz_dep_.CoverTab[84854]++
//line /usr/local/go/src/encoding/gob/encode.go:111
		// _ = "end of CoverTab[84854]"
//line /usr/local/go/src/encoding/gob/encode.go:111
	}
//line /usr/local/go/src/encoding/gob/encode.go:111
	// _ = "end of CoverTab[84851]"
//line /usr/local/go/src/encoding/gob/encode.go:111
	_go_fuzz_dep_.CoverTab[84852]++

							binary.BigEndian.PutUint64(state.buf[1:], x)
							bc := bits.LeadingZeros64(x) >> 3
							state.buf[bc] = uint8(bc - uint64Size)

							state.b.Write(state.buf[bc : uint64Size+1])
//line /usr/local/go/src/encoding/gob/encode.go:117
	// _ = "end of CoverTab[84852]"
}

//line /usr/local/go/src/encoding/gob/encode.go:123
func (state *encoderState) encodeInt(i int64) {
//line /usr/local/go/src/encoding/gob/encode.go:123
	_go_fuzz_dep_.CoverTab[84855]++
							var x uint64
							if i < 0 {
//line /usr/local/go/src/encoding/gob/encode.go:125
		_go_fuzz_dep_.CoverTab[84857]++
								x = uint64(^i<<1) | 1
//line /usr/local/go/src/encoding/gob/encode.go:126
		// _ = "end of CoverTab[84857]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:127
		_go_fuzz_dep_.CoverTab[84858]++
								x = uint64(i << 1)
//line /usr/local/go/src/encoding/gob/encode.go:128
		// _ = "end of CoverTab[84858]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:129
	// _ = "end of CoverTab[84855]"
//line /usr/local/go/src/encoding/gob/encode.go:129
	_go_fuzz_dep_.CoverTab[84856]++
							state.encodeUint(x)
//line /usr/local/go/src/encoding/gob/encode.go:130
	// _ = "end of CoverTab[84856]"
}

//line /usr/local/go/src/encoding/gob/encode.go:134
type encOp func(i *encInstr, state *encoderState, v reflect.Value)

//line /usr/local/go/src/encoding/gob/encode.go:137
type encInstr struct {
	op	encOp
	field	int
	index	[]int
	indir	int
}

//line /usr/local/go/src/encoding/gob/encode.go:146
func (state *encoderState) update(instr *encInstr) {
//line /usr/local/go/src/encoding/gob/encode.go:146
	_go_fuzz_dep_.CoverTab[84859]++
							if instr != nil {
//line /usr/local/go/src/encoding/gob/encode.go:147
		_go_fuzz_dep_.CoverTab[84860]++
								state.encodeUint(uint64(instr.field - state.fieldnum))
								state.fieldnum = instr.field
//line /usr/local/go/src/encoding/gob/encode.go:149
		// _ = "end of CoverTab[84860]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:150
		_go_fuzz_dep_.CoverTab[84861]++
//line /usr/local/go/src/encoding/gob/encode.go:150
		// _ = "end of CoverTab[84861]"
//line /usr/local/go/src/encoding/gob/encode.go:150
	}
//line /usr/local/go/src/encoding/gob/encode.go:150
	// _ = "end of CoverTab[84859]"
}

//line /usr/local/go/src/encoding/gob/encode.go:163
func encIndirect(pv reflect.Value, indir int) reflect.Value {
//line /usr/local/go/src/encoding/gob/encode.go:163
	_go_fuzz_dep_.CoverTab[84862]++
							for ; indir > 0; indir-- {
//line /usr/local/go/src/encoding/gob/encode.go:164
		_go_fuzz_dep_.CoverTab[84864]++
								if pv.IsNil() {
//line /usr/local/go/src/encoding/gob/encode.go:165
			_go_fuzz_dep_.CoverTab[84866]++
									break
//line /usr/local/go/src/encoding/gob/encode.go:166
			// _ = "end of CoverTab[84866]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:167
			_go_fuzz_dep_.CoverTab[84867]++
//line /usr/local/go/src/encoding/gob/encode.go:167
			// _ = "end of CoverTab[84867]"
//line /usr/local/go/src/encoding/gob/encode.go:167
		}
//line /usr/local/go/src/encoding/gob/encode.go:167
		// _ = "end of CoverTab[84864]"
//line /usr/local/go/src/encoding/gob/encode.go:167
		_go_fuzz_dep_.CoverTab[84865]++
								pv = pv.Elem()
//line /usr/local/go/src/encoding/gob/encode.go:168
		// _ = "end of CoverTab[84865]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:169
	// _ = "end of CoverTab[84862]"
//line /usr/local/go/src/encoding/gob/encode.go:169
	_go_fuzz_dep_.CoverTab[84863]++
							return pv
//line /usr/local/go/src/encoding/gob/encode.go:170
	// _ = "end of CoverTab[84863]"
}

//line /usr/local/go/src/encoding/gob/encode.go:174
func encBool(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:174
	_go_fuzz_dep_.CoverTab[84868]++
							b := v.Bool()
							if b || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:176
		_go_fuzz_dep_.CoverTab[84869]++
//line /usr/local/go/src/encoding/gob/encode.go:176
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:176
		// _ = "end of CoverTab[84869]"
//line /usr/local/go/src/encoding/gob/encode.go:176
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:176
		_go_fuzz_dep_.CoverTab[84870]++
								state.update(i)
								if b {
//line /usr/local/go/src/encoding/gob/encode.go:178
			_go_fuzz_dep_.CoverTab[84871]++
									state.encodeUint(1)
//line /usr/local/go/src/encoding/gob/encode.go:179
			// _ = "end of CoverTab[84871]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:180
			_go_fuzz_dep_.CoverTab[84872]++
									state.encodeUint(0)
//line /usr/local/go/src/encoding/gob/encode.go:181
			// _ = "end of CoverTab[84872]"
		}
//line /usr/local/go/src/encoding/gob/encode.go:182
		// _ = "end of CoverTab[84870]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:183
		_go_fuzz_dep_.CoverTab[84873]++
//line /usr/local/go/src/encoding/gob/encode.go:183
		// _ = "end of CoverTab[84873]"
//line /usr/local/go/src/encoding/gob/encode.go:183
	}
//line /usr/local/go/src/encoding/gob/encode.go:183
	// _ = "end of CoverTab[84868]"
}

//line /usr/local/go/src/encoding/gob/encode.go:187
func encInt(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:187
	_go_fuzz_dep_.CoverTab[84874]++
							value := v.Int()
							if value != 0 || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:189
		_go_fuzz_dep_.CoverTab[84875]++
//line /usr/local/go/src/encoding/gob/encode.go:189
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:189
		// _ = "end of CoverTab[84875]"
//line /usr/local/go/src/encoding/gob/encode.go:189
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:189
		_go_fuzz_dep_.CoverTab[84876]++
								state.update(i)
								state.encodeInt(value)
//line /usr/local/go/src/encoding/gob/encode.go:191
		// _ = "end of CoverTab[84876]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:192
		_go_fuzz_dep_.CoverTab[84877]++
//line /usr/local/go/src/encoding/gob/encode.go:192
		// _ = "end of CoverTab[84877]"
//line /usr/local/go/src/encoding/gob/encode.go:192
	}
//line /usr/local/go/src/encoding/gob/encode.go:192
	// _ = "end of CoverTab[84874]"
}

//line /usr/local/go/src/encoding/gob/encode.go:196
func encUint(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:196
	_go_fuzz_dep_.CoverTab[84878]++
							value := v.Uint()
							if value != 0 || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:198
		_go_fuzz_dep_.CoverTab[84879]++
//line /usr/local/go/src/encoding/gob/encode.go:198
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:198
		// _ = "end of CoverTab[84879]"
//line /usr/local/go/src/encoding/gob/encode.go:198
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:198
		_go_fuzz_dep_.CoverTab[84880]++
								state.update(i)
								state.encodeUint(value)
//line /usr/local/go/src/encoding/gob/encode.go:200
		// _ = "end of CoverTab[84880]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:201
		_go_fuzz_dep_.CoverTab[84881]++
//line /usr/local/go/src/encoding/gob/encode.go:201
		// _ = "end of CoverTab[84881]"
//line /usr/local/go/src/encoding/gob/encode.go:201
	}
//line /usr/local/go/src/encoding/gob/encode.go:201
	// _ = "end of CoverTab[84878]"
}

//line /usr/local/go/src/encoding/gob/encode.go:210
func floatBits(f float64) uint64 {
//line /usr/local/go/src/encoding/gob/encode.go:210
	_go_fuzz_dep_.CoverTab[84882]++
							u := math.Float64bits(f)
							return bits.ReverseBytes64(u)
//line /usr/local/go/src/encoding/gob/encode.go:212
	// _ = "end of CoverTab[84882]"
}

//line /usr/local/go/src/encoding/gob/encode.go:216
func encFloat(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:216
	_go_fuzz_dep_.CoverTab[84883]++
							f := v.Float()
							if f != 0 || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:218
		_go_fuzz_dep_.CoverTab[84884]++
//line /usr/local/go/src/encoding/gob/encode.go:218
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:218
		// _ = "end of CoverTab[84884]"
//line /usr/local/go/src/encoding/gob/encode.go:218
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:218
		_go_fuzz_dep_.CoverTab[84885]++
								bits := floatBits(f)
								state.update(i)
								state.encodeUint(bits)
//line /usr/local/go/src/encoding/gob/encode.go:221
		// _ = "end of CoverTab[84885]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:222
		_go_fuzz_dep_.CoverTab[84886]++
//line /usr/local/go/src/encoding/gob/encode.go:222
		// _ = "end of CoverTab[84886]"
//line /usr/local/go/src/encoding/gob/encode.go:222
	}
//line /usr/local/go/src/encoding/gob/encode.go:222
	// _ = "end of CoverTab[84883]"
}

//line /usr/local/go/src/encoding/gob/encode.go:227
func encComplex(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:227
	_go_fuzz_dep_.CoverTab[84887]++
							c := v.Complex()
							if c != 0+0i || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:229
		_go_fuzz_dep_.CoverTab[84888]++
//line /usr/local/go/src/encoding/gob/encode.go:229
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:229
		// _ = "end of CoverTab[84888]"
//line /usr/local/go/src/encoding/gob/encode.go:229
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:229
		_go_fuzz_dep_.CoverTab[84889]++
								rpart := floatBits(real(c))
								ipart := floatBits(imag(c))
								state.update(i)
								state.encodeUint(rpart)
								state.encodeUint(ipart)
//line /usr/local/go/src/encoding/gob/encode.go:234
		// _ = "end of CoverTab[84889]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:235
		_go_fuzz_dep_.CoverTab[84890]++
//line /usr/local/go/src/encoding/gob/encode.go:235
		// _ = "end of CoverTab[84890]"
//line /usr/local/go/src/encoding/gob/encode.go:235
	}
//line /usr/local/go/src/encoding/gob/encode.go:235
	// _ = "end of CoverTab[84887]"
}

//line /usr/local/go/src/encoding/gob/encode.go:240
func encUint8Array(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:240
	_go_fuzz_dep_.CoverTab[84891]++
							b := v.Bytes()
							if len(b) > 0 || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:242
		_go_fuzz_dep_.CoverTab[84892]++
//line /usr/local/go/src/encoding/gob/encode.go:242
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:242
		// _ = "end of CoverTab[84892]"
//line /usr/local/go/src/encoding/gob/encode.go:242
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:242
		_go_fuzz_dep_.CoverTab[84893]++
								state.update(i)
								state.encodeUint(uint64(len(b)))
								state.b.Write(b)
//line /usr/local/go/src/encoding/gob/encode.go:245
		// _ = "end of CoverTab[84893]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:246
		_go_fuzz_dep_.CoverTab[84894]++
//line /usr/local/go/src/encoding/gob/encode.go:246
		// _ = "end of CoverTab[84894]"
//line /usr/local/go/src/encoding/gob/encode.go:246
	}
//line /usr/local/go/src/encoding/gob/encode.go:246
	// _ = "end of CoverTab[84891]"
}

//line /usr/local/go/src/encoding/gob/encode.go:251
func encString(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:251
	_go_fuzz_dep_.CoverTab[84895]++
							s := v.String()
							if len(s) > 0 || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:253
		_go_fuzz_dep_.CoverTab[84896]++
//line /usr/local/go/src/encoding/gob/encode.go:253
		return state.sendZero
//line /usr/local/go/src/encoding/gob/encode.go:253
		// _ = "end of CoverTab[84896]"
//line /usr/local/go/src/encoding/gob/encode.go:253
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:253
		_go_fuzz_dep_.CoverTab[84897]++
								state.update(i)
								state.encodeUint(uint64(len(s)))
								state.b.WriteString(s)
//line /usr/local/go/src/encoding/gob/encode.go:256
		// _ = "end of CoverTab[84897]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:257
		_go_fuzz_dep_.CoverTab[84898]++
//line /usr/local/go/src/encoding/gob/encode.go:257
		// _ = "end of CoverTab[84898]"
//line /usr/local/go/src/encoding/gob/encode.go:257
	}
//line /usr/local/go/src/encoding/gob/encode.go:257
	// _ = "end of CoverTab[84895]"
}

//line /usr/local/go/src/encoding/gob/encode.go:262
func encStructTerminator(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:262
	_go_fuzz_dep_.CoverTab[84899]++
							state.encodeUint(0)
//line /usr/local/go/src/encoding/gob/encode.go:263
	// _ = "end of CoverTab[84899]"
}

//line /usr/local/go/src/encoding/gob/encode.go:270
type encEngine struct {
	instr []encInstr
}

const singletonField = 0

//line /usr/local/go/src/encoding/gob/encode.go:278
func valid(v reflect.Value) bool {
//line /usr/local/go/src/encoding/gob/encode.go:278
	_go_fuzz_dep_.CoverTab[84900]++
							switch v.Kind() {
	case reflect.Invalid:
//line /usr/local/go/src/encoding/gob/encode.go:280
		_go_fuzz_dep_.CoverTab[84902]++
								return false
//line /usr/local/go/src/encoding/gob/encode.go:281
		// _ = "end of CoverTab[84902]"
	case reflect.Pointer:
//line /usr/local/go/src/encoding/gob/encode.go:282
		_go_fuzz_dep_.CoverTab[84903]++
								return !v.IsNil()
//line /usr/local/go/src/encoding/gob/encode.go:283
		// _ = "end of CoverTab[84903]"
//line /usr/local/go/src/encoding/gob/encode.go:283
	default:
//line /usr/local/go/src/encoding/gob/encode.go:283
		_go_fuzz_dep_.CoverTab[84904]++
//line /usr/local/go/src/encoding/gob/encode.go:283
		// _ = "end of CoverTab[84904]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:284
	// _ = "end of CoverTab[84900]"
//line /usr/local/go/src/encoding/gob/encode.go:284
	_go_fuzz_dep_.CoverTab[84901]++
							return true
//line /usr/local/go/src/encoding/gob/encode.go:285
	// _ = "end of CoverTab[84901]"
}

//line /usr/local/go/src/encoding/gob/encode.go:289
func (enc *Encoder) encodeSingle(b *encBuffer, engine *encEngine, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:289
	_go_fuzz_dep_.CoverTab[84905]++
							state := enc.newEncoderState(b)
							defer enc.freeEncoderState(state)
							state.fieldnum = singletonField

//line /usr/local/go/src/encoding/gob/encode.go:295
	state.sendZero = true
	instr := &engine.instr[singletonField]
	if instr.indir > 0 {
//line /usr/local/go/src/encoding/gob/encode.go:297
		_go_fuzz_dep_.CoverTab[84907]++
								value = encIndirect(value, instr.indir)
//line /usr/local/go/src/encoding/gob/encode.go:298
		// _ = "end of CoverTab[84907]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:299
		_go_fuzz_dep_.CoverTab[84908]++
//line /usr/local/go/src/encoding/gob/encode.go:299
		// _ = "end of CoverTab[84908]"
//line /usr/local/go/src/encoding/gob/encode.go:299
	}
//line /usr/local/go/src/encoding/gob/encode.go:299
	// _ = "end of CoverTab[84905]"
//line /usr/local/go/src/encoding/gob/encode.go:299
	_go_fuzz_dep_.CoverTab[84906]++
							if valid(value) {
//line /usr/local/go/src/encoding/gob/encode.go:300
		_go_fuzz_dep_.CoverTab[84909]++
								instr.op(instr, state, value)
//line /usr/local/go/src/encoding/gob/encode.go:301
		// _ = "end of CoverTab[84909]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:302
		_go_fuzz_dep_.CoverTab[84910]++
//line /usr/local/go/src/encoding/gob/encode.go:302
		// _ = "end of CoverTab[84910]"
//line /usr/local/go/src/encoding/gob/encode.go:302
	}
//line /usr/local/go/src/encoding/gob/encode.go:302
	// _ = "end of CoverTab[84906]"
}

//line /usr/local/go/src/encoding/gob/encode.go:306
func (enc *Encoder) encodeStruct(b *encBuffer, engine *encEngine, value reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:306
	_go_fuzz_dep_.CoverTab[84911]++
							if !valid(value) {
//line /usr/local/go/src/encoding/gob/encode.go:307
		_go_fuzz_dep_.CoverTab[84913]++
								return
//line /usr/local/go/src/encoding/gob/encode.go:308
		// _ = "end of CoverTab[84913]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:309
		_go_fuzz_dep_.CoverTab[84914]++
//line /usr/local/go/src/encoding/gob/encode.go:309
		// _ = "end of CoverTab[84914]"
//line /usr/local/go/src/encoding/gob/encode.go:309
	}
//line /usr/local/go/src/encoding/gob/encode.go:309
	// _ = "end of CoverTab[84911]"
//line /usr/local/go/src/encoding/gob/encode.go:309
	_go_fuzz_dep_.CoverTab[84912]++
							state := enc.newEncoderState(b)
							defer enc.freeEncoderState(state)
							state.fieldnum = -1
							for i := 0; i < len(engine.instr); i++ {
//line /usr/local/go/src/encoding/gob/encode.go:313
		_go_fuzz_dep_.CoverTab[84915]++
								instr := &engine.instr[i]
								if i >= value.NumField() {
//line /usr/local/go/src/encoding/gob/encode.go:315
			_go_fuzz_dep_.CoverTab[84918]++

									instr.op(instr, state, reflect.Value{})
									break
//line /usr/local/go/src/encoding/gob/encode.go:318
			// _ = "end of CoverTab[84918]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:319
			_go_fuzz_dep_.CoverTab[84919]++
//line /usr/local/go/src/encoding/gob/encode.go:319
			// _ = "end of CoverTab[84919]"
//line /usr/local/go/src/encoding/gob/encode.go:319
		}
//line /usr/local/go/src/encoding/gob/encode.go:319
		// _ = "end of CoverTab[84915]"
//line /usr/local/go/src/encoding/gob/encode.go:319
		_go_fuzz_dep_.CoverTab[84916]++
								field := value.FieldByIndex(instr.index)
								if instr.indir > 0 {
//line /usr/local/go/src/encoding/gob/encode.go:321
			_go_fuzz_dep_.CoverTab[84920]++
									field = encIndirect(field, instr.indir)

									if !valid(field) {
//line /usr/local/go/src/encoding/gob/encode.go:324
				_go_fuzz_dep_.CoverTab[84921]++
										continue
//line /usr/local/go/src/encoding/gob/encode.go:325
				// _ = "end of CoverTab[84921]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:326
				_go_fuzz_dep_.CoverTab[84922]++
//line /usr/local/go/src/encoding/gob/encode.go:326
				// _ = "end of CoverTab[84922]"
//line /usr/local/go/src/encoding/gob/encode.go:326
			}
//line /usr/local/go/src/encoding/gob/encode.go:326
			// _ = "end of CoverTab[84920]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:327
			_go_fuzz_dep_.CoverTab[84923]++
//line /usr/local/go/src/encoding/gob/encode.go:327
			// _ = "end of CoverTab[84923]"
//line /usr/local/go/src/encoding/gob/encode.go:327
		}
//line /usr/local/go/src/encoding/gob/encode.go:327
		// _ = "end of CoverTab[84916]"
//line /usr/local/go/src/encoding/gob/encode.go:327
		_go_fuzz_dep_.CoverTab[84917]++
								instr.op(instr, state, field)
//line /usr/local/go/src/encoding/gob/encode.go:328
		// _ = "end of CoverTab[84917]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:329
	// _ = "end of CoverTab[84912]"
}

//line /usr/local/go/src/encoding/gob/encode.go:333
func (enc *Encoder) encodeArray(b *encBuffer, value reflect.Value, op encOp, elemIndir int, length int, helper encHelper) {
//line /usr/local/go/src/encoding/gob/encode.go:333
	_go_fuzz_dep_.CoverTab[84924]++
							state := enc.newEncoderState(b)
							defer enc.freeEncoderState(state)
							state.fieldnum = -1
							state.sendZero = true
							state.encodeUint(uint64(length))
							if helper != nil && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:339
		_go_fuzz_dep_.CoverTab[84926]++
//line /usr/local/go/src/encoding/gob/encode.go:339
		return helper(state, value)
//line /usr/local/go/src/encoding/gob/encode.go:339
		// _ = "end of CoverTab[84926]"
//line /usr/local/go/src/encoding/gob/encode.go:339
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:339
		_go_fuzz_dep_.CoverTab[84927]++
								return
//line /usr/local/go/src/encoding/gob/encode.go:340
		// _ = "end of CoverTab[84927]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:341
		_go_fuzz_dep_.CoverTab[84928]++
//line /usr/local/go/src/encoding/gob/encode.go:341
		// _ = "end of CoverTab[84928]"
//line /usr/local/go/src/encoding/gob/encode.go:341
	}
//line /usr/local/go/src/encoding/gob/encode.go:341
	// _ = "end of CoverTab[84924]"
//line /usr/local/go/src/encoding/gob/encode.go:341
	_go_fuzz_dep_.CoverTab[84925]++
							for i := 0; i < length; i++ {
//line /usr/local/go/src/encoding/gob/encode.go:342
		_go_fuzz_dep_.CoverTab[84929]++
								elem := value.Index(i)
								if elemIndir > 0 {
//line /usr/local/go/src/encoding/gob/encode.go:344
			_go_fuzz_dep_.CoverTab[84931]++
									elem = encIndirect(elem, elemIndir)

									if !valid(elem) {
//line /usr/local/go/src/encoding/gob/encode.go:347
				_go_fuzz_dep_.CoverTab[84932]++
										errorf("encodeArray: nil element")
//line /usr/local/go/src/encoding/gob/encode.go:348
				// _ = "end of CoverTab[84932]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:349
				_go_fuzz_dep_.CoverTab[84933]++
//line /usr/local/go/src/encoding/gob/encode.go:349
				// _ = "end of CoverTab[84933]"
//line /usr/local/go/src/encoding/gob/encode.go:349
			}
//line /usr/local/go/src/encoding/gob/encode.go:349
			// _ = "end of CoverTab[84931]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:350
			_go_fuzz_dep_.CoverTab[84934]++
//line /usr/local/go/src/encoding/gob/encode.go:350
			// _ = "end of CoverTab[84934]"
//line /usr/local/go/src/encoding/gob/encode.go:350
		}
//line /usr/local/go/src/encoding/gob/encode.go:350
		// _ = "end of CoverTab[84929]"
//line /usr/local/go/src/encoding/gob/encode.go:350
		_go_fuzz_dep_.CoverTab[84930]++
								op(nil, state, elem)
//line /usr/local/go/src/encoding/gob/encode.go:351
		// _ = "end of CoverTab[84930]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:352
	// _ = "end of CoverTab[84925]"
}

//line /usr/local/go/src/encoding/gob/encode.go:356
func encodeReflectValue(state *encoderState, v reflect.Value, op encOp, indir int) {
//line /usr/local/go/src/encoding/gob/encode.go:356
	_go_fuzz_dep_.CoverTab[84935]++
							for i := 0; i < indir && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:357
		_go_fuzz_dep_.CoverTab[84938]++
//line /usr/local/go/src/encoding/gob/encode.go:357
		return v.IsValid()
//line /usr/local/go/src/encoding/gob/encode.go:357
		// _ = "end of CoverTab[84938]"
//line /usr/local/go/src/encoding/gob/encode.go:357
	}(); i++ {
//line /usr/local/go/src/encoding/gob/encode.go:357
		_go_fuzz_dep_.CoverTab[84939]++
								v = reflect.Indirect(v)
//line /usr/local/go/src/encoding/gob/encode.go:358
		// _ = "end of CoverTab[84939]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:359
	// _ = "end of CoverTab[84935]"
//line /usr/local/go/src/encoding/gob/encode.go:359
	_go_fuzz_dep_.CoverTab[84936]++
							if !v.IsValid() {
//line /usr/local/go/src/encoding/gob/encode.go:360
		_go_fuzz_dep_.CoverTab[84940]++
								errorf("encodeReflectValue: nil element")
//line /usr/local/go/src/encoding/gob/encode.go:361
		// _ = "end of CoverTab[84940]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:362
		_go_fuzz_dep_.CoverTab[84941]++
//line /usr/local/go/src/encoding/gob/encode.go:362
		// _ = "end of CoverTab[84941]"
//line /usr/local/go/src/encoding/gob/encode.go:362
	}
//line /usr/local/go/src/encoding/gob/encode.go:362
	// _ = "end of CoverTab[84936]"
//line /usr/local/go/src/encoding/gob/encode.go:362
	_go_fuzz_dep_.CoverTab[84937]++
							op(nil, state, v)
//line /usr/local/go/src/encoding/gob/encode.go:363
	// _ = "end of CoverTab[84937]"
}

//line /usr/local/go/src/encoding/gob/encode.go:367
func (enc *Encoder) encodeMap(b *encBuffer, mv reflect.Value, keyOp, elemOp encOp, keyIndir, elemIndir int) {
//line /usr/local/go/src/encoding/gob/encode.go:367
	_go_fuzz_dep_.CoverTab[84942]++
							state := enc.newEncoderState(b)
							state.fieldnum = -1
							state.sendZero = true
							state.encodeUint(uint64(mv.Len()))
							mi := mv.MapRange()
							for mi.Next() {
//line /usr/local/go/src/encoding/gob/encode.go:373
		_go_fuzz_dep_.CoverTab[84944]++
								encodeReflectValue(state, mi.Key(), keyOp, keyIndir)
								encodeReflectValue(state, mi.Value(), elemOp, elemIndir)
//line /usr/local/go/src/encoding/gob/encode.go:375
		// _ = "end of CoverTab[84944]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:376
	// _ = "end of CoverTab[84942]"
//line /usr/local/go/src/encoding/gob/encode.go:376
	_go_fuzz_dep_.CoverTab[84943]++
							enc.freeEncoderState(state)
//line /usr/local/go/src/encoding/gob/encode.go:377
	// _ = "end of CoverTab[84943]"
}

//line /usr/local/go/src/encoding/gob/encode.go:385
func (enc *Encoder) encodeInterface(b *encBuffer, iv reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:385
	_go_fuzz_dep_.CoverTab[84945]++

//line /usr/local/go/src/encoding/gob/encode.go:388
	elem := iv.Elem()
	if elem.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:389
		_go_fuzz_dep_.CoverTab[84951]++
//line /usr/local/go/src/encoding/gob/encode.go:389
		return elem.IsNil()
//line /usr/local/go/src/encoding/gob/encode.go:389
		// _ = "end of CoverTab[84951]"
//line /usr/local/go/src/encoding/gob/encode.go:389
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:389
		_go_fuzz_dep_.CoverTab[84952]++
								errorf("gob: cannot encode nil pointer of type %s inside interface", iv.Elem().Type())
//line /usr/local/go/src/encoding/gob/encode.go:390
		// _ = "end of CoverTab[84952]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:391
		_go_fuzz_dep_.CoverTab[84953]++
//line /usr/local/go/src/encoding/gob/encode.go:391
		// _ = "end of CoverTab[84953]"
//line /usr/local/go/src/encoding/gob/encode.go:391
	}
//line /usr/local/go/src/encoding/gob/encode.go:391
	// _ = "end of CoverTab[84945]"
//line /usr/local/go/src/encoding/gob/encode.go:391
	_go_fuzz_dep_.CoverTab[84946]++
							state := enc.newEncoderState(b)
							state.fieldnum = -1
							state.sendZero = true
							if iv.IsNil() {
//line /usr/local/go/src/encoding/gob/encode.go:395
		_go_fuzz_dep_.CoverTab[84954]++
								state.encodeUint(0)
								return
//line /usr/local/go/src/encoding/gob/encode.go:397
		// _ = "end of CoverTab[84954]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:398
		_go_fuzz_dep_.CoverTab[84955]++
//line /usr/local/go/src/encoding/gob/encode.go:398
		// _ = "end of CoverTab[84955]"
//line /usr/local/go/src/encoding/gob/encode.go:398
	}
//line /usr/local/go/src/encoding/gob/encode.go:398
	// _ = "end of CoverTab[84946]"
//line /usr/local/go/src/encoding/gob/encode.go:398
	_go_fuzz_dep_.CoverTab[84947]++

							ut := userType(iv.Elem().Type())
							namei, ok := concreteTypeToName.Load(ut.base)
							if !ok {
//line /usr/local/go/src/encoding/gob/encode.go:402
		_go_fuzz_dep_.CoverTab[84956]++
								errorf("type not registered for interface: %s", ut.base)
//line /usr/local/go/src/encoding/gob/encode.go:403
		// _ = "end of CoverTab[84956]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:404
		_go_fuzz_dep_.CoverTab[84957]++
//line /usr/local/go/src/encoding/gob/encode.go:404
		// _ = "end of CoverTab[84957]"
//line /usr/local/go/src/encoding/gob/encode.go:404
	}
//line /usr/local/go/src/encoding/gob/encode.go:404
	// _ = "end of CoverTab[84947]"
//line /usr/local/go/src/encoding/gob/encode.go:404
	_go_fuzz_dep_.CoverTab[84948]++
							name := namei.(string)

//line /usr/local/go/src/encoding/gob/encode.go:408
	state.encodeUint(uint64(len(name)))
							state.b.WriteString(name)

							enc.sendTypeDescriptor(enc.writer(), state, ut)

							enc.sendTypeId(state, ut)

//line /usr/local/go/src/encoding/gob/encode.go:416
	enc.pushWriter(b)
	data := encBufferPool.Get().(*encBuffer)
	data.Write(spaceForLength)
	enc.encode(data, elem, ut)
	if enc.err != nil {
//line /usr/local/go/src/encoding/gob/encode.go:420
		_go_fuzz_dep_.CoverTab[84958]++
								error_(enc.err)
//line /usr/local/go/src/encoding/gob/encode.go:421
		// _ = "end of CoverTab[84958]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:422
		_go_fuzz_dep_.CoverTab[84959]++
//line /usr/local/go/src/encoding/gob/encode.go:422
		// _ = "end of CoverTab[84959]"
//line /usr/local/go/src/encoding/gob/encode.go:422
	}
//line /usr/local/go/src/encoding/gob/encode.go:422
	// _ = "end of CoverTab[84948]"
//line /usr/local/go/src/encoding/gob/encode.go:422
	_go_fuzz_dep_.CoverTab[84949]++
							enc.popWriter()
							enc.writeMessage(b, data)
							data.Reset()
							encBufferPool.Put(data)
							if enc.err != nil {
//line /usr/local/go/src/encoding/gob/encode.go:427
		_go_fuzz_dep_.CoverTab[84960]++
								error_(enc.err)
//line /usr/local/go/src/encoding/gob/encode.go:428
		// _ = "end of CoverTab[84960]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:429
		_go_fuzz_dep_.CoverTab[84961]++
//line /usr/local/go/src/encoding/gob/encode.go:429
		// _ = "end of CoverTab[84961]"
//line /usr/local/go/src/encoding/gob/encode.go:429
	}
//line /usr/local/go/src/encoding/gob/encode.go:429
	// _ = "end of CoverTab[84949]"
//line /usr/local/go/src/encoding/gob/encode.go:429
	_go_fuzz_dep_.CoverTab[84950]++
							enc.freeEncoderState(state)
//line /usr/local/go/src/encoding/gob/encode.go:430
	// _ = "end of CoverTab[84950]"
}

//line /usr/local/go/src/encoding/gob/encode.go:434
func isZero(val reflect.Value) bool {
//line /usr/local/go/src/encoding/gob/encode.go:434
	_go_fuzz_dep_.CoverTab[84962]++
							switch val.Kind() {
	case reflect.Array:
//line /usr/local/go/src/encoding/gob/encode.go:436
		_go_fuzz_dep_.CoverTab[84964]++
								for i := 0; i < val.Len(); i++ {
//line /usr/local/go/src/encoding/gob/encode.go:437
			_go_fuzz_dep_.CoverTab[84976]++
									if !isZero(val.Index(i)) {
//line /usr/local/go/src/encoding/gob/encode.go:438
				_go_fuzz_dep_.CoverTab[84977]++
										return false
//line /usr/local/go/src/encoding/gob/encode.go:439
				// _ = "end of CoverTab[84977]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:440
				_go_fuzz_dep_.CoverTab[84978]++
//line /usr/local/go/src/encoding/gob/encode.go:440
				// _ = "end of CoverTab[84978]"
//line /usr/local/go/src/encoding/gob/encode.go:440
			}
//line /usr/local/go/src/encoding/gob/encode.go:440
			// _ = "end of CoverTab[84976]"
		}
//line /usr/local/go/src/encoding/gob/encode.go:441
		// _ = "end of CoverTab[84964]"
//line /usr/local/go/src/encoding/gob/encode.go:441
		_go_fuzz_dep_.CoverTab[84965]++
								return true
//line /usr/local/go/src/encoding/gob/encode.go:442
		// _ = "end of CoverTab[84965]"
	case reflect.Map, reflect.Slice, reflect.String:
//line /usr/local/go/src/encoding/gob/encode.go:443
		_go_fuzz_dep_.CoverTab[84966]++
								return val.Len() == 0
//line /usr/local/go/src/encoding/gob/encode.go:444
		// _ = "end of CoverTab[84966]"
	case reflect.Bool:
//line /usr/local/go/src/encoding/gob/encode.go:445
		_go_fuzz_dep_.CoverTab[84967]++
								return !val.Bool()
//line /usr/local/go/src/encoding/gob/encode.go:446
		// _ = "end of CoverTab[84967]"
	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/encoding/gob/encode.go:447
		_go_fuzz_dep_.CoverTab[84968]++
								return val.Complex() == 0
//line /usr/local/go/src/encoding/gob/encode.go:448
		// _ = "end of CoverTab[84968]"
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Pointer:
//line /usr/local/go/src/encoding/gob/encode.go:449
		_go_fuzz_dep_.CoverTab[84969]++
								return val.IsNil()
//line /usr/local/go/src/encoding/gob/encode.go:450
		// _ = "end of CoverTab[84969]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/gob/encode.go:451
		_go_fuzz_dep_.CoverTab[84970]++
								return val.Int() == 0
//line /usr/local/go/src/encoding/gob/encode.go:452
		// _ = "end of CoverTab[84970]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/encoding/gob/encode.go:453
		_go_fuzz_dep_.CoverTab[84971]++
								return val.Float() == 0
//line /usr/local/go/src/encoding/gob/encode.go:454
		// _ = "end of CoverTab[84971]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/gob/encode.go:455
		_go_fuzz_dep_.CoverTab[84972]++
								return val.Uint() == 0
//line /usr/local/go/src/encoding/gob/encode.go:456
		// _ = "end of CoverTab[84972]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/gob/encode.go:457
		_go_fuzz_dep_.CoverTab[84973]++
								for i := 0; i < val.NumField(); i++ {
//line /usr/local/go/src/encoding/gob/encode.go:458
			_go_fuzz_dep_.CoverTab[84979]++
									if !isZero(val.Field(i)) {
//line /usr/local/go/src/encoding/gob/encode.go:459
				_go_fuzz_dep_.CoverTab[84980]++
										return false
//line /usr/local/go/src/encoding/gob/encode.go:460
				// _ = "end of CoverTab[84980]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:461
				_go_fuzz_dep_.CoverTab[84981]++
//line /usr/local/go/src/encoding/gob/encode.go:461
				// _ = "end of CoverTab[84981]"
//line /usr/local/go/src/encoding/gob/encode.go:461
			}
//line /usr/local/go/src/encoding/gob/encode.go:461
			// _ = "end of CoverTab[84979]"
		}
//line /usr/local/go/src/encoding/gob/encode.go:462
		// _ = "end of CoverTab[84973]"
//line /usr/local/go/src/encoding/gob/encode.go:462
		_go_fuzz_dep_.CoverTab[84974]++
								return true
//line /usr/local/go/src/encoding/gob/encode.go:463
		// _ = "end of CoverTab[84974]"
//line /usr/local/go/src/encoding/gob/encode.go:463
	default:
//line /usr/local/go/src/encoding/gob/encode.go:463
		_go_fuzz_dep_.CoverTab[84975]++
//line /usr/local/go/src/encoding/gob/encode.go:463
		// _ = "end of CoverTab[84975]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:464
	// _ = "end of CoverTab[84962]"
//line /usr/local/go/src/encoding/gob/encode.go:464
	_go_fuzz_dep_.CoverTab[84963]++
							panic("unknown type in isZero " + val.Type().String())
//line /usr/local/go/src/encoding/gob/encode.go:465
	// _ = "end of CoverTab[84963]"
}

//line /usr/local/go/src/encoding/gob/encode.go:470
func (enc *Encoder) encodeGobEncoder(b *encBuffer, ut *userTypeInfo, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:470
	_go_fuzz_dep_.CoverTab[84982]++

//line /usr/local/go/src/encoding/gob/encode.go:473
	var data []byte
	var err error

	switch ut.externalEnc {
	case xGob:
//line /usr/local/go/src/encoding/gob/encode.go:477
		_go_fuzz_dep_.CoverTab[84985]++
								data, err = v.Interface().(GobEncoder).GobEncode()
//line /usr/local/go/src/encoding/gob/encode.go:478
		// _ = "end of CoverTab[84985]"
	case xBinary:
//line /usr/local/go/src/encoding/gob/encode.go:479
		_go_fuzz_dep_.CoverTab[84986]++
								data, err = v.Interface().(encoding.BinaryMarshaler).MarshalBinary()
//line /usr/local/go/src/encoding/gob/encode.go:480
		// _ = "end of CoverTab[84986]"
	case xText:
//line /usr/local/go/src/encoding/gob/encode.go:481
		_go_fuzz_dep_.CoverTab[84987]++
								data, err = v.Interface().(encoding.TextMarshaler).MarshalText()
//line /usr/local/go/src/encoding/gob/encode.go:482
		// _ = "end of CoverTab[84987]"
//line /usr/local/go/src/encoding/gob/encode.go:482
	default:
//line /usr/local/go/src/encoding/gob/encode.go:482
		_go_fuzz_dep_.CoverTab[84988]++
//line /usr/local/go/src/encoding/gob/encode.go:482
		// _ = "end of CoverTab[84988]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:483
	// _ = "end of CoverTab[84982]"
//line /usr/local/go/src/encoding/gob/encode.go:483
	_go_fuzz_dep_.CoverTab[84983]++
							if err != nil {
//line /usr/local/go/src/encoding/gob/encode.go:484
		_go_fuzz_dep_.CoverTab[84989]++
								error_(err)
//line /usr/local/go/src/encoding/gob/encode.go:485
		// _ = "end of CoverTab[84989]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:486
		_go_fuzz_dep_.CoverTab[84990]++
//line /usr/local/go/src/encoding/gob/encode.go:486
		// _ = "end of CoverTab[84990]"
//line /usr/local/go/src/encoding/gob/encode.go:486
	}
//line /usr/local/go/src/encoding/gob/encode.go:486
	// _ = "end of CoverTab[84983]"
//line /usr/local/go/src/encoding/gob/encode.go:486
	_go_fuzz_dep_.CoverTab[84984]++
							state := enc.newEncoderState(b)
							state.fieldnum = -1
							state.encodeUint(uint64(len(data)))
							state.b.Write(data)
							enc.freeEncoderState(state)
//line /usr/local/go/src/encoding/gob/encode.go:491
	// _ = "end of CoverTab[84984]"
}

var encOpTable = [...]encOp{
	reflect.Bool:		encBool,
	reflect.Int:		encInt,
	reflect.Int8:		encInt,
	reflect.Int16:		encInt,
	reflect.Int32:		encInt,
	reflect.Int64:		encInt,
	reflect.Uint:		encUint,
	reflect.Uint8:		encUint,
	reflect.Uint16:		encUint,
	reflect.Uint32:		encUint,
	reflect.Uint64:		encUint,
	reflect.Uintptr:	encUint,
	reflect.Float32:	encFloat,
	reflect.Float64:	encFloat,
	reflect.Complex64:	encComplex,
	reflect.Complex128:	encComplex,
	reflect.String:		encString,
}

//line /usr/local/go/src/encoding/gob/encode.go:516
func encOpFor(rt reflect.Type, inProgress map[reflect.Type]*encOp, building map[*typeInfo]bool) (*encOp, int) {
//line /usr/local/go/src/encoding/gob/encode.go:516
	_go_fuzz_dep_.CoverTab[84991]++
							ut := userType(rt)

							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/encode.go:519
		_go_fuzz_dep_.CoverTab[84997]++
								return gobEncodeOpFor(ut)
//line /usr/local/go/src/encoding/gob/encode.go:520
		// _ = "end of CoverTab[84997]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:521
		_go_fuzz_dep_.CoverTab[84998]++
//line /usr/local/go/src/encoding/gob/encode.go:521
		// _ = "end of CoverTab[84998]"
//line /usr/local/go/src/encoding/gob/encode.go:521
	}
//line /usr/local/go/src/encoding/gob/encode.go:521
	// _ = "end of CoverTab[84991]"
//line /usr/local/go/src/encoding/gob/encode.go:521
	_go_fuzz_dep_.CoverTab[84992]++

//line /usr/local/go/src/encoding/gob/encode.go:524
	if opPtr := inProgress[rt]; opPtr != nil {
//line /usr/local/go/src/encoding/gob/encode.go:524
		_go_fuzz_dep_.CoverTab[84999]++
								return opPtr, ut.indir
//line /usr/local/go/src/encoding/gob/encode.go:525
		// _ = "end of CoverTab[84999]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:526
		_go_fuzz_dep_.CoverTab[85000]++
//line /usr/local/go/src/encoding/gob/encode.go:526
		// _ = "end of CoverTab[85000]"
//line /usr/local/go/src/encoding/gob/encode.go:526
	}
//line /usr/local/go/src/encoding/gob/encode.go:526
	// _ = "end of CoverTab[84992]"
//line /usr/local/go/src/encoding/gob/encode.go:526
	_go_fuzz_dep_.CoverTab[84993]++
							typ := ut.base
							indir := ut.indir
							k := typ.Kind()
							var op encOp
							if int(k) < len(encOpTable) {
//line /usr/local/go/src/encoding/gob/encode.go:531
		_go_fuzz_dep_.CoverTab[85001]++
								op = encOpTable[k]
//line /usr/local/go/src/encoding/gob/encode.go:532
		// _ = "end of CoverTab[85001]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:533
		_go_fuzz_dep_.CoverTab[85002]++
//line /usr/local/go/src/encoding/gob/encode.go:533
		// _ = "end of CoverTab[85002]"
//line /usr/local/go/src/encoding/gob/encode.go:533
	}
//line /usr/local/go/src/encoding/gob/encode.go:533
	// _ = "end of CoverTab[84993]"
//line /usr/local/go/src/encoding/gob/encode.go:533
	_go_fuzz_dep_.CoverTab[84994]++
							if op == nil {
//line /usr/local/go/src/encoding/gob/encode.go:534
		_go_fuzz_dep_.CoverTab[85003]++
								inProgress[rt] = &op

								switch t := typ; t.Kind() {
		case reflect.Slice:
//line /usr/local/go/src/encoding/gob/encode.go:538
			_go_fuzz_dep_.CoverTab[85004]++
									if t.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/gob/encode.go:539
				_go_fuzz_dep_.CoverTab[85011]++
										op = encUint8Array
										break
//line /usr/local/go/src/encoding/gob/encode.go:541
				// _ = "end of CoverTab[85011]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:542
				_go_fuzz_dep_.CoverTab[85012]++
//line /usr/local/go/src/encoding/gob/encode.go:542
				// _ = "end of CoverTab[85012]"
//line /usr/local/go/src/encoding/gob/encode.go:542
			}
//line /usr/local/go/src/encoding/gob/encode.go:542
			// _ = "end of CoverTab[85004]"
//line /usr/local/go/src/encoding/gob/encode.go:542
			_go_fuzz_dep_.CoverTab[85005]++

									elemOp, elemIndir := encOpFor(t.Elem(), inProgress, building)
									helper := encSliceHelper[t.Elem().Kind()]
									op = func(i *encInstr, state *encoderState, slice reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:546
				_go_fuzz_dep_.CoverTab[85013]++
										if !state.sendZero && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:547
					_go_fuzz_dep_.CoverTab[85015]++
//line /usr/local/go/src/encoding/gob/encode.go:547
					return slice.Len() == 0
//line /usr/local/go/src/encoding/gob/encode.go:547
					// _ = "end of CoverTab[85015]"
//line /usr/local/go/src/encoding/gob/encode.go:547
				}() {
//line /usr/local/go/src/encoding/gob/encode.go:547
					_go_fuzz_dep_.CoverTab[85016]++
											return
//line /usr/local/go/src/encoding/gob/encode.go:548
					// _ = "end of CoverTab[85016]"
				} else {
//line /usr/local/go/src/encoding/gob/encode.go:549
					_go_fuzz_dep_.CoverTab[85017]++
//line /usr/local/go/src/encoding/gob/encode.go:549
					// _ = "end of CoverTab[85017]"
//line /usr/local/go/src/encoding/gob/encode.go:549
				}
//line /usr/local/go/src/encoding/gob/encode.go:549
				// _ = "end of CoverTab[85013]"
//line /usr/local/go/src/encoding/gob/encode.go:549
				_go_fuzz_dep_.CoverTab[85014]++
										state.update(i)
										state.enc.encodeArray(state.b, slice, *elemOp, elemIndir, slice.Len(), helper)
//line /usr/local/go/src/encoding/gob/encode.go:551
				// _ = "end of CoverTab[85014]"
			}
//line /usr/local/go/src/encoding/gob/encode.go:552
			// _ = "end of CoverTab[85005]"
		case reflect.Array:
//line /usr/local/go/src/encoding/gob/encode.go:553
			_go_fuzz_dep_.CoverTab[85006]++

									elemOp, elemIndir := encOpFor(t.Elem(), inProgress, building)
									helper := encArrayHelper[t.Elem().Kind()]
									op = func(i *encInstr, state *encoderState, array reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:557
				_go_fuzz_dep_.CoverTab[85018]++
										state.update(i)
										state.enc.encodeArray(state.b, array, *elemOp, elemIndir, array.Len(), helper)
//line /usr/local/go/src/encoding/gob/encode.go:559
				// _ = "end of CoverTab[85018]"
			}
//line /usr/local/go/src/encoding/gob/encode.go:560
			// _ = "end of CoverTab[85006]"
		case reflect.Map:
//line /usr/local/go/src/encoding/gob/encode.go:561
			_go_fuzz_dep_.CoverTab[85007]++
									keyOp, keyIndir := encOpFor(t.Key(), inProgress, building)
									elemOp, elemIndir := encOpFor(t.Elem(), inProgress, building)
									op = func(i *encInstr, state *encoderState, mv reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:564
				_go_fuzz_dep_.CoverTab[85019]++

//line /usr/local/go/src/encoding/gob/encode.go:567
				if !state.sendZero && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:567
					_go_fuzz_dep_.CoverTab[85021]++
//line /usr/local/go/src/encoding/gob/encode.go:567
					return mv.IsNil()
//line /usr/local/go/src/encoding/gob/encode.go:567
					// _ = "end of CoverTab[85021]"
//line /usr/local/go/src/encoding/gob/encode.go:567
				}() {
//line /usr/local/go/src/encoding/gob/encode.go:567
					_go_fuzz_dep_.CoverTab[85022]++
											return
//line /usr/local/go/src/encoding/gob/encode.go:568
					// _ = "end of CoverTab[85022]"
				} else {
//line /usr/local/go/src/encoding/gob/encode.go:569
					_go_fuzz_dep_.CoverTab[85023]++
//line /usr/local/go/src/encoding/gob/encode.go:569
					// _ = "end of CoverTab[85023]"
//line /usr/local/go/src/encoding/gob/encode.go:569
				}
//line /usr/local/go/src/encoding/gob/encode.go:569
				// _ = "end of CoverTab[85019]"
//line /usr/local/go/src/encoding/gob/encode.go:569
				_go_fuzz_dep_.CoverTab[85020]++
										state.update(i)
										state.enc.encodeMap(state.b, mv, *keyOp, *elemOp, keyIndir, elemIndir)
//line /usr/local/go/src/encoding/gob/encode.go:571
				// _ = "end of CoverTab[85020]"
			}
//line /usr/local/go/src/encoding/gob/encode.go:572
			// _ = "end of CoverTab[85007]"
		case reflect.Struct:
//line /usr/local/go/src/encoding/gob/encode.go:573
			_go_fuzz_dep_.CoverTab[85008]++

									getEncEngine(userType(typ), building)
									info := mustGetTypeInfo(typ)
									op = func(i *encInstr, state *encoderState, sv reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:577
				_go_fuzz_dep_.CoverTab[85024]++
										state.update(i)

										enc := info.encoder.Load()
										state.enc.encodeStruct(state.b, enc, sv)
//line /usr/local/go/src/encoding/gob/encode.go:581
				// _ = "end of CoverTab[85024]"
			}
//line /usr/local/go/src/encoding/gob/encode.go:582
			// _ = "end of CoverTab[85008]"
		case reflect.Interface:
//line /usr/local/go/src/encoding/gob/encode.go:583
			_go_fuzz_dep_.CoverTab[85009]++
									op = func(i *encInstr, state *encoderState, iv reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:584
				_go_fuzz_dep_.CoverTab[85025]++
										if !state.sendZero && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:585
					_go_fuzz_dep_.CoverTab[85027]++
//line /usr/local/go/src/encoding/gob/encode.go:585
					return (!iv.IsValid() || func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:585
						_go_fuzz_dep_.CoverTab[85028]++
//line /usr/local/go/src/encoding/gob/encode.go:585
						return iv.IsNil()
//line /usr/local/go/src/encoding/gob/encode.go:585
						// _ = "end of CoverTab[85028]"
//line /usr/local/go/src/encoding/gob/encode.go:585
					}())
//line /usr/local/go/src/encoding/gob/encode.go:585
					// _ = "end of CoverTab[85027]"
//line /usr/local/go/src/encoding/gob/encode.go:585
				}() {
//line /usr/local/go/src/encoding/gob/encode.go:585
					_go_fuzz_dep_.CoverTab[85029]++
											return
//line /usr/local/go/src/encoding/gob/encode.go:586
					// _ = "end of CoverTab[85029]"
				} else {
//line /usr/local/go/src/encoding/gob/encode.go:587
					_go_fuzz_dep_.CoverTab[85030]++
//line /usr/local/go/src/encoding/gob/encode.go:587
					// _ = "end of CoverTab[85030]"
//line /usr/local/go/src/encoding/gob/encode.go:587
				}
//line /usr/local/go/src/encoding/gob/encode.go:587
				// _ = "end of CoverTab[85025]"
//line /usr/local/go/src/encoding/gob/encode.go:587
				_go_fuzz_dep_.CoverTab[85026]++
										state.update(i)
										state.enc.encodeInterface(state.b, iv)
//line /usr/local/go/src/encoding/gob/encode.go:589
				// _ = "end of CoverTab[85026]"
			}
//line /usr/local/go/src/encoding/gob/encode.go:590
			// _ = "end of CoverTab[85009]"
//line /usr/local/go/src/encoding/gob/encode.go:590
		default:
//line /usr/local/go/src/encoding/gob/encode.go:590
			_go_fuzz_dep_.CoverTab[85010]++
//line /usr/local/go/src/encoding/gob/encode.go:590
			// _ = "end of CoverTab[85010]"
		}
//line /usr/local/go/src/encoding/gob/encode.go:591
		// _ = "end of CoverTab[85003]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:592
		_go_fuzz_dep_.CoverTab[85031]++
//line /usr/local/go/src/encoding/gob/encode.go:592
		// _ = "end of CoverTab[85031]"
//line /usr/local/go/src/encoding/gob/encode.go:592
	}
//line /usr/local/go/src/encoding/gob/encode.go:592
	// _ = "end of CoverTab[84994]"
//line /usr/local/go/src/encoding/gob/encode.go:592
	_go_fuzz_dep_.CoverTab[84995]++
							if op == nil {
//line /usr/local/go/src/encoding/gob/encode.go:593
		_go_fuzz_dep_.CoverTab[85032]++
								errorf("can't happen: encode type %s", rt)
//line /usr/local/go/src/encoding/gob/encode.go:594
		// _ = "end of CoverTab[85032]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:595
		_go_fuzz_dep_.CoverTab[85033]++
//line /usr/local/go/src/encoding/gob/encode.go:595
		// _ = "end of CoverTab[85033]"
//line /usr/local/go/src/encoding/gob/encode.go:595
	}
//line /usr/local/go/src/encoding/gob/encode.go:595
	// _ = "end of CoverTab[84995]"
//line /usr/local/go/src/encoding/gob/encode.go:595
	_go_fuzz_dep_.CoverTab[84996]++
							return &op, indir
//line /usr/local/go/src/encoding/gob/encode.go:596
	// _ = "end of CoverTab[84996]"
}

//line /usr/local/go/src/encoding/gob/encode.go:600
func gobEncodeOpFor(ut *userTypeInfo) (*encOp, int) {
//line /usr/local/go/src/encoding/gob/encode.go:600
	_go_fuzz_dep_.CoverTab[85034]++
							rt := ut.user
							if ut.encIndir == -1 {
//line /usr/local/go/src/encoding/gob/encode.go:602
		_go_fuzz_dep_.CoverTab[85037]++
								rt = reflect.PointerTo(rt)
//line /usr/local/go/src/encoding/gob/encode.go:603
		// _ = "end of CoverTab[85037]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:604
		_go_fuzz_dep_.CoverTab[85038]++
//line /usr/local/go/src/encoding/gob/encode.go:604
		if ut.encIndir > 0 {
//line /usr/local/go/src/encoding/gob/encode.go:604
			_go_fuzz_dep_.CoverTab[85039]++
									for i := int8(0); i < ut.encIndir; i++ {
//line /usr/local/go/src/encoding/gob/encode.go:605
				_go_fuzz_dep_.CoverTab[85040]++
										rt = rt.Elem()
//line /usr/local/go/src/encoding/gob/encode.go:606
				// _ = "end of CoverTab[85040]"
			}
//line /usr/local/go/src/encoding/gob/encode.go:607
			// _ = "end of CoverTab[85039]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:608
			_go_fuzz_dep_.CoverTab[85041]++
//line /usr/local/go/src/encoding/gob/encode.go:608
			// _ = "end of CoverTab[85041]"
//line /usr/local/go/src/encoding/gob/encode.go:608
		}
//line /usr/local/go/src/encoding/gob/encode.go:608
		// _ = "end of CoverTab[85038]"
//line /usr/local/go/src/encoding/gob/encode.go:608
	}
//line /usr/local/go/src/encoding/gob/encode.go:608
	// _ = "end of CoverTab[85034]"
//line /usr/local/go/src/encoding/gob/encode.go:608
	_go_fuzz_dep_.CoverTab[85035]++
							var op encOp
							op = func(i *encInstr, state *encoderState, v reflect.Value) {
//line /usr/local/go/src/encoding/gob/encode.go:610
		_go_fuzz_dep_.CoverTab[85042]++
								if ut.encIndir == -1 {
//line /usr/local/go/src/encoding/gob/encode.go:611
			_go_fuzz_dep_.CoverTab[85045]++

									if !v.CanAddr() {
//line /usr/local/go/src/encoding/gob/encode.go:613
				_go_fuzz_dep_.CoverTab[85047]++
										errorf("unaddressable value of type %s", rt)
//line /usr/local/go/src/encoding/gob/encode.go:614
				// _ = "end of CoverTab[85047]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:615
				_go_fuzz_dep_.CoverTab[85048]++
//line /usr/local/go/src/encoding/gob/encode.go:615
				// _ = "end of CoverTab[85048]"
//line /usr/local/go/src/encoding/gob/encode.go:615
			}
//line /usr/local/go/src/encoding/gob/encode.go:615
			// _ = "end of CoverTab[85045]"
//line /usr/local/go/src/encoding/gob/encode.go:615
			_go_fuzz_dep_.CoverTab[85046]++
									v = v.Addr()
//line /usr/local/go/src/encoding/gob/encode.go:616
			// _ = "end of CoverTab[85046]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:617
			_go_fuzz_dep_.CoverTab[85049]++
//line /usr/local/go/src/encoding/gob/encode.go:617
			// _ = "end of CoverTab[85049]"
//line /usr/local/go/src/encoding/gob/encode.go:617
		}
//line /usr/local/go/src/encoding/gob/encode.go:617
		// _ = "end of CoverTab[85042]"
//line /usr/local/go/src/encoding/gob/encode.go:617
		_go_fuzz_dep_.CoverTab[85043]++
								if !state.sendZero && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:618
			_go_fuzz_dep_.CoverTab[85050]++
//line /usr/local/go/src/encoding/gob/encode.go:618
			return isZero(v)
//line /usr/local/go/src/encoding/gob/encode.go:618
			// _ = "end of CoverTab[85050]"
//line /usr/local/go/src/encoding/gob/encode.go:618
		}() {
//line /usr/local/go/src/encoding/gob/encode.go:618
			_go_fuzz_dep_.CoverTab[85051]++
									return
//line /usr/local/go/src/encoding/gob/encode.go:619
			// _ = "end of CoverTab[85051]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:620
			_go_fuzz_dep_.CoverTab[85052]++
//line /usr/local/go/src/encoding/gob/encode.go:620
			// _ = "end of CoverTab[85052]"
//line /usr/local/go/src/encoding/gob/encode.go:620
		}
//line /usr/local/go/src/encoding/gob/encode.go:620
		// _ = "end of CoverTab[85043]"
//line /usr/local/go/src/encoding/gob/encode.go:620
		_go_fuzz_dep_.CoverTab[85044]++
								state.update(i)
								state.enc.encodeGobEncoder(state.b, ut, v)
//line /usr/local/go/src/encoding/gob/encode.go:622
		// _ = "end of CoverTab[85044]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:623
	// _ = "end of CoverTab[85035]"
//line /usr/local/go/src/encoding/gob/encode.go:623
	_go_fuzz_dep_.CoverTab[85036]++
							return &op, int(ut.encIndir)
//line /usr/local/go/src/encoding/gob/encode.go:624
	// _ = "end of CoverTab[85036]"
}

//line /usr/local/go/src/encoding/gob/encode.go:628
func compileEnc(ut *userTypeInfo, building map[*typeInfo]bool) *encEngine {
//line /usr/local/go/src/encoding/gob/encode.go:628
	_go_fuzz_dep_.CoverTab[85053]++
							srt := ut.base
							engine := new(encEngine)
							seen := make(map[reflect.Type]*encOp)
							rt := ut.base
							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/encode.go:633
		_go_fuzz_dep_.CoverTab[85056]++
								rt = ut.user
//line /usr/local/go/src/encoding/gob/encode.go:634
		// _ = "end of CoverTab[85056]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:635
		_go_fuzz_dep_.CoverTab[85057]++
//line /usr/local/go/src/encoding/gob/encode.go:635
		// _ = "end of CoverTab[85057]"
//line /usr/local/go/src/encoding/gob/encode.go:635
	}
//line /usr/local/go/src/encoding/gob/encode.go:635
	// _ = "end of CoverTab[85053]"
//line /usr/local/go/src/encoding/gob/encode.go:635
	_go_fuzz_dep_.CoverTab[85054]++
							if ut.externalEnc == 0 && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:636
		_go_fuzz_dep_.CoverTab[85058]++
//line /usr/local/go/src/encoding/gob/encode.go:636
		return srt.Kind() == reflect.Struct
//line /usr/local/go/src/encoding/gob/encode.go:636
		// _ = "end of CoverTab[85058]"
//line /usr/local/go/src/encoding/gob/encode.go:636
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:636
		_go_fuzz_dep_.CoverTab[85059]++
								for fieldNum, wireFieldNum := 0, 0; fieldNum < srt.NumField(); fieldNum++ {
//line /usr/local/go/src/encoding/gob/encode.go:637
			_go_fuzz_dep_.CoverTab[85062]++
									f := srt.Field(fieldNum)
									if !isSent(&f) {
//line /usr/local/go/src/encoding/gob/encode.go:639
				_go_fuzz_dep_.CoverTab[85064]++
										continue
//line /usr/local/go/src/encoding/gob/encode.go:640
				// _ = "end of CoverTab[85064]"
			} else {
//line /usr/local/go/src/encoding/gob/encode.go:641
				_go_fuzz_dep_.CoverTab[85065]++
//line /usr/local/go/src/encoding/gob/encode.go:641
				// _ = "end of CoverTab[85065]"
//line /usr/local/go/src/encoding/gob/encode.go:641
			}
//line /usr/local/go/src/encoding/gob/encode.go:641
			// _ = "end of CoverTab[85062]"
//line /usr/local/go/src/encoding/gob/encode.go:641
			_go_fuzz_dep_.CoverTab[85063]++
									op, indir := encOpFor(f.Type, seen, building)
									engine.instr = append(engine.instr, encInstr{*op, wireFieldNum, f.Index, indir})
									wireFieldNum++
//line /usr/local/go/src/encoding/gob/encode.go:644
			// _ = "end of CoverTab[85063]"
		}
//line /usr/local/go/src/encoding/gob/encode.go:645
		// _ = "end of CoverTab[85059]"
//line /usr/local/go/src/encoding/gob/encode.go:645
		_go_fuzz_dep_.CoverTab[85060]++
								if srt.NumField() > 0 && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:646
			_go_fuzz_dep_.CoverTab[85066]++
//line /usr/local/go/src/encoding/gob/encode.go:646
			return len(engine.instr) == 0
//line /usr/local/go/src/encoding/gob/encode.go:646
			// _ = "end of CoverTab[85066]"
//line /usr/local/go/src/encoding/gob/encode.go:646
		}() {
//line /usr/local/go/src/encoding/gob/encode.go:646
			_go_fuzz_dep_.CoverTab[85067]++
									errorf("type %s has no exported fields", rt)
//line /usr/local/go/src/encoding/gob/encode.go:647
			// _ = "end of CoverTab[85067]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:648
			_go_fuzz_dep_.CoverTab[85068]++
//line /usr/local/go/src/encoding/gob/encode.go:648
			// _ = "end of CoverTab[85068]"
//line /usr/local/go/src/encoding/gob/encode.go:648
		}
//line /usr/local/go/src/encoding/gob/encode.go:648
		// _ = "end of CoverTab[85060]"
//line /usr/local/go/src/encoding/gob/encode.go:648
		_go_fuzz_dep_.CoverTab[85061]++
								engine.instr = append(engine.instr, encInstr{encStructTerminator, 0, nil, 0})
//line /usr/local/go/src/encoding/gob/encode.go:649
		// _ = "end of CoverTab[85061]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:650
		_go_fuzz_dep_.CoverTab[85069]++
								engine.instr = make([]encInstr, 1)
								op, indir := encOpFor(rt, seen, building)
								engine.instr[0] = encInstr{*op, singletonField, nil, indir}
//line /usr/local/go/src/encoding/gob/encode.go:653
		// _ = "end of CoverTab[85069]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:654
	// _ = "end of CoverTab[85054]"
//line /usr/local/go/src/encoding/gob/encode.go:654
	_go_fuzz_dep_.CoverTab[85055]++
							return engine
//line /usr/local/go/src/encoding/gob/encode.go:655
	// _ = "end of CoverTab[85055]"
}

//line /usr/local/go/src/encoding/gob/encode.go:659
func getEncEngine(ut *userTypeInfo, building map[*typeInfo]bool) *encEngine {
//line /usr/local/go/src/encoding/gob/encode.go:659
	_go_fuzz_dep_.CoverTab[85070]++
							info, err := getTypeInfo(ut)
							if err != nil {
//line /usr/local/go/src/encoding/gob/encode.go:661
		_go_fuzz_dep_.CoverTab[85073]++
								error_(err)
//line /usr/local/go/src/encoding/gob/encode.go:662
		// _ = "end of CoverTab[85073]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:663
		_go_fuzz_dep_.CoverTab[85074]++
//line /usr/local/go/src/encoding/gob/encode.go:663
		// _ = "end of CoverTab[85074]"
//line /usr/local/go/src/encoding/gob/encode.go:663
	}
//line /usr/local/go/src/encoding/gob/encode.go:663
	// _ = "end of CoverTab[85070]"
//line /usr/local/go/src/encoding/gob/encode.go:663
	_go_fuzz_dep_.CoverTab[85071]++
							enc := info.encoder.Load()
							if enc == nil {
//line /usr/local/go/src/encoding/gob/encode.go:665
		_go_fuzz_dep_.CoverTab[85075]++
								enc = buildEncEngine(info, ut, building)
//line /usr/local/go/src/encoding/gob/encode.go:666
		// _ = "end of CoverTab[85075]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:667
		_go_fuzz_dep_.CoverTab[85076]++
//line /usr/local/go/src/encoding/gob/encode.go:667
		// _ = "end of CoverTab[85076]"
//line /usr/local/go/src/encoding/gob/encode.go:667
	}
//line /usr/local/go/src/encoding/gob/encode.go:667
	// _ = "end of CoverTab[85071]"
//line /usr/local/go/src/encoding/gob/encode.go:667
	_go_fuzz_dep_.CoverTab[85072]++
							return enc
//line /usr/local/go/src/encoding/gob/encode.go:668
	// _ = "end of CoverTab[85072]"
}

func buildEncEngine(info *typeInfo, ut *userTypeInfo, building map[*typeInfo]bool) *encEngine {
//line /usr/local/go/src/encoding/gob/encode.go:671
	_go_fuzz_dep_.CoverTab[85077]++

							if building != nil && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:673
		_go_fuzz_dep_.CoverTab[85080]++
//line /usr/local/go/src/encoding/gob/encode.go:673
		return building[info]
//line /usr/local/go/src/encoding/gob/encode.go:673
		// _ = "end of CoverTab[85080]"
//line /usr/local/go/src/encoding/gob/encode.go:673
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:673
		_go_fuzz_dep_.CoverTab[85081]++
								return nil
//line /usr/local/go/src/encoding/gob/encode.go:674
		// _ = "end of CoverTab[85081]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:675
		_go_fuzz_dep_.CoverTab[85082]++
//line /usr/local/go/src/encoding/gob/encode.go:675
		// _ = "end of CoverTab[85082]"
//line /usr/local/go/src/encoding/gob/encode.go:675
	}
//line /usr/local/go/src/encoding/gob/encode.go:675
	// _ = "end of CoverTab[85077]"
//line /usr/local/go/src/encoding/gob/encode.go:675
	_go_fuzz_dep_.CoverTab[85078]++
							info.encInit.Lock()
							defer info.encInit.Unlock()
							enc := info.encoder.Load()
							if enc == nil {
//line /usr/local/go/src/encoding/gob/encode.go:679
		_go_fuzz_dep_.CoverTab[85083]++
								if building == nil {
//line /usr/local/go/src/encoding/gob/encode.go:680
			_go_fuzz_dep_.CoverTab[85085]++
									building = make(map[*typeInfo]bool)
//line /usr/local/go/src/encoding/gob/encode.go:681
			// _ = "end of CoverTab[85085]"
		} else {
//line /usr/local/go/src/encoding/gob/encode.go:682
			_go_fuzz_dep_.CoverTab[85086]++
//line /usr/local/go/src/encoding/gob/encode.go:682
			// _ = "end of CoverTab[85086]"
//line /usr/local/go/src/encoding/gob/encode.go:682
		}
//line /usr/local/go/src/encoding/gob/encode.go:682
		// _ = "end of CoverTab[85083]"
//line /usr/local/go/src/encoding/gob/encode.go:682
		_go_fuzz_dep_.CoverTab[85084]++
								building[info] = true
								enc = compileEnc(ut, building)
								info.encoder.Store(enc)
//line /usr/local/go/src/encoding/gob/encode.go:685
		// _ = "end of CoverTab[85084]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:686
		_go_fuzz_dep_.CoverTab[85087]++
//line /usr/local/go/src/encoding/gob/encode.go:686
		// _ = "end of CoverTab[85087]"
//line /usr/local/go/src/encoding/gob/encode.go:686
	}
//line /usr/local/go/src/encoding/gob/encode.go:686
	// _ = "end of CoverTab[85078]"
//line /usr/local/go/src/encoding/gob/encode.go:686
	_go_fuzz_dep_.CoverTab[85079]++
							return enc
//line /usr/local/go/src/encoding/gob/encode.go:687
	// _ = "end of CoverTab[85079]"
}

func (enc *Encoder) encode(b *encBuffer, value reflect.Value, ut *userTypeInfo) {
//line /usr/local/go/src/encoding/gob/encode.go:690
	_go_fuzz_dep_.CoverTab[85088]++
							defer catchError(&enc.err)
							engine := getEncEngine(ut, nil)
							indir := ut.indir
							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/encode.go:694
		_go_fuzz_dep_.CoverTab[85091]++
								indir = int(ut.encIndir)
//line /usr/local/go/src/encoding/gob/encode.go:695
		// _ = "end of CoverTab[85091]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:696
		_go_fuzz_dep_.CoverTab[85092]++
//line /usr/local/go/src/encoding/gob/encode.go:696
		// _ = "end of CoverTab[85092]"
//line /usr/local/go/src/encoding/gob/encode.go:696
	}
//line /usr/local/go/src/encoding/gob/encode.go:696
	// _ = "end of CoverTab[85088]"
//line /usr/local/go/src/encoding/gob/encode.go:696
	_go_fuzz_dep_.CoverTab[85089]++
							for i := 0; i < indir; i++ {
//line /usr/local/go/src/encoding/gob/encode.go:697
		_go_fuzz_dep_.CoverTab[85093]++
								value = reflect.Indirect(value)
//line /usr/local/go/src/encoding/gob/encode.go:698
		// _ = "end of CoverTab[85093]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:699
	// _ = "end of CoverTab[85089]"
//line /usr/local/go/src/encoding/gob/encode.go:699
	_go_fuzz_dep_.CoverTab[85090]++
							if ut.externalEnc == 0 && func() bool {
//line /usr/local/go/src/encoding/gob/encode.go:700
		_go_fuzz_dep_.CoverTab[85094]++
//line /usr/local/go/src/encoding/gob/encode.go:700
		return value.Type().Kind() == reflect.Struct
//line /usr/local/go/src/encoding/gob/encode.go:700
		// _ = "end of CoverTab[85094]"
//line /usr/local/go/src/encoding/gob/encode.go:700
	}() {
//line /usr/local/go/src/encoding/gob/encode.go:700
		_go_fuzz_dep_.CoverTab[85095]++
								enc.encodeStruct(b, engine, value)
//line /usr/local/go/src/encoding/gob/encode.go:701
		// _ = "end of CoverTab[85095]"
	} else {
//line /usr/local/go/src/encoding/gob/encode.go:702
		_go_fuzz_dep_.CoverTab[85096]++
								enc.encodeSingle(b, engine, value)
//line /usr/local/go/src/encoding/gob/encode.go:703
		// _ = "end of CoverTab[85096]"
	}
//line /usr/local/go/src/encoding/gob/encode.go:704
	// _ = "end of CoverTab[85090]"
}

//line /usr/local/go/src/encoding/gob/encode.go:705
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/gob/encode.go:705
var _ = _go_fuzz_dep_.CoverTab
