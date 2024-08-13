// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/gob/encoder.go:5
package gob

//line /usr/local/go/src/encoding/gob/encoder.go:5
import (
//line /usr/local/go/src/encoding/gob/encoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/gob/encoder.go:5
)
//line /usr/local/go/src/encoding/gob/encoder.go:5
import (
//line /usr/local/go/src/encoding/gob/encoder.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/gob/encoder.go:5
)

import (
	"errors"
	"io"
	"reflect"
	"sync"
)

// An Encoder manages the transmission of type and data information to the
//line /usr/local/go/src/encoding/gob/encoder.go:14
// other side of a connection.  It is safe for concurrent use by multiple
//line /usr/local/go/src/encoding/gob/encoder.go:14
// goroutines.
//line /usr/local/go/src/encoding/gob/encoder.go:17
type Encoder struct {
	mutex		sync.Mutex		// each item must be sent atomically
	w		[]io.Writer		// where to send the data
	sent		map[reflect.Type]typeId	// which types we've already sent
	countState	*encoderState		// stage for writing counts
	freeList	*encoderState		// list of free encoderStates; avoids reallocation
	byteBuf		encBuffer		// buffer for top-level encoderState
	err		error
}

// Before we encode a message, we reserve space at the head of the
//line /usr/local/go/src/encoding/gob/encoder.go:27
// buffer in which to encode its length. This means we can use the
//line /usr/local/go/src/encoding/gob/encoder.go:27
// buffer to assemble the message without another allocation.
//line /usr/local/go/src/encoding/gob/encoder.go:30
const maxLength = 9	// Maximum size of an encoded length.
var spaceForLength = make([]byte, maxLength)

// NewEncoder returns a new encoder that will transmit on the io.Writer.
func NewEncoder(w io.Writer) *Encoder {
//line /usr/local/go/src/encoding/gob/encoder.go:34
	_go_fuzz_dep_.CoverTab[85097]++
							enc := new(Encoder)
							enc.w = []io.Writer{w}
							enc.sent = make(map[reflect.Type]typeId)
							enc.countState = enc.newEncoderState(new(encBuffer))
							return enc
//line /usr/local/go/src/encoding/gob/encoder.go:39
	// _ = "end of CoverTab[85097]"
}

// writer() returns the innermost writer the encoder is using
func (enc *Encoder) writer() io.Writer {
//line /usr/local/go/src/encoding/gob/encoder.go:43
	_go_fuzz_dep_.CoverTab[85098]++
							return enc.w[len(enc.w)-1]
//line /usr/local/go/src/encoding/gob/encoder.go:44
	// _ = "end of CoverTab[85098]"
}

// pushWriter adds a writer to the encoder.
func (enc *Encoder) pushWriter(w io.Writer) {
//line /usr/local/go/src/encoding/gob/encoder.go:48
	_go_fuzz_dep_.CoverTab[85099]++
							enc.w = append(enc.w, w)
//line /usr/local/go/src/encoding/gob/encoder.go:49
	// _ = "end of CoverTab[85099]"
}

// popWriter pops the innermost writer.
func (enc *Encoder) popWriter() {
//line /usr/local/go/src/encoding/gob/encoder.go:53
	_go_fuzz_dep_.CoverTab[85100]++
							enc.w = enc.w[0 : len(enc.w)-1]
//line /usr/local/go/src/encoding/gob/encoder.go:54
	// _ = "end of CoverTab[85100]"
}

func (enc *Encoder) setError(err error) {
//line /usr/local/go/src/encoding/gob/encoder.go:57
	_go_fuzz_dep_.CoverTab[85101]++
							if enc.err == nil {
//line /usr/local/go/src/encoding/gob/encoder.go:58
		_go_fuzz_dep_.CoverTab[85102]++
								enc.err = err
//line /usr/local/go/src/encoding/gob/encoder.go:59
		// _ = "end of CoverTab[85102]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:60
		_go_fuzz_dep_.CoverTab[85103]++
//line /usr/local/go/src/encoding/gob/encoder.go:60
		// _ = "end of CoverTab[85103]"
//line /usr/local/go/src/encoding/gob/encoder.go:60
	}
//line /usr/local/go/src/encoding/gob/encoder.go:60
	// _ = "end of CoverTab[85101]"
}

// writeMessage sends the data item preceded by a unsigned count of its length.
func (enc *Encoder) writeMessage(w io.Writer, b *encBuffer) {
//line /usr/local/go/src/encoding/gob/encoder.go:64
	_go_fuzz_dep_.CoverTab[85104]++

//line /usr/local/go/src/encoding/gob/encoder.go:68
	message := b.Bytes()
	messageLen := len(message) - maxLength

	if messageLen >= tooBig {
//line /usr/local/go/src/encoding/gob/encoder.go:71
		_go_fuzz_dep_.CoverTab[85106]++
								enc.setError(errors.New("gob: encoder: message too big"))
								return
//line /usr/local/go/src/encoding/gob/encoder.go:73
		// _ = "end of CoverTab[85106]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:74
		_go_fuzz_dep_.CoverTab[85107]++
//line /usr/local/go/src/encoding/gob/encoder.go:74
		// _ = "end of CoverTab[85107]"
//line /usr/local/go/src/encoding/gob/encoder.go:74
	}
//line /usr/local/go/src/encoding/gob/encoder.go:74
	// _ = "end of CoverTab[85104]"
//line /usr/local/go/src/encoding/gob/encoder.go:74
	_go_fuzz_dep_.CoverTab[85105]++

							enc.countState.b.Reset()
							enc.countState.encodeUint(uint64(messageLen))

							offset := maxLength - enc.countState.b.Len()
							copy(message[offset:], enc.countState.b.Bytes())

							_, err := w.Write(message[offset:])

							b.Reset()
							b.Write(spaceForLength)
							if err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:86
		_go_fuzz_dep_.CoverTab[85108]++
								enc.setError(err)
//line /usr/local/go/src/encoding/gob/encoder.go:87
		// _ = "end of CoverTab[85108]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:88
		_go_fuzz_dep_.CoverTab[85109]++
//line /usr/local/go/src/encoding/gob/encoder.go:88
		// _ = "end of CoverTab[85109]"
//line /usr/local/go/src/encoding/gob/encoder.go:88
	}
//line /usr/local/go/src/encoding/gob/encoder.go:88
	// _ = "end of CoverTab[85105]"
}

// sendActualType sends the requested type, without further investigation, unless
//line /usr/local/go/src/encoding/gob/encoder.go:91
// it's been sent before.
//line /usr/local/go/src/encoding/gob/encoder.go:93
func (enc *Encoder) sendActualType(w io.Writer, state *encoderState, ut *userTypeInfo, actual reflect.Type) (sent bool) {
//line /usr/local/go/src/encoding/gob/encoder.go:93
	_go_fuzz_dep_.CoverTab[85110]++
							if _, alreadySent := enc.sent[actual]; alreadySent {
//line /usr/local/go/src/encoding/gob/encoder.go:94
		_go_fuzz_dep_.CoverTab[85116]++
								return false
//line /usr/local/go/src/encoding/gob/encoder.go:95
		// _ = "end of CoverTab[85116]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:96
		_go_fuzz_dep_.CoverTab[85117]++
//line /usr/local/go/src/encoding/gob/encoder.go:96
		// _ = "end of CoverTab[85117]"
//line /usr/local/go/src/encoding/gob/encoder.go:96
	}
//line /usr/local/go/src/encoding/gob/encoder.go:96
	// _ = "end of CoverTab[85110]"
//line /usr/local/go/src/encoding/gob/encoder.go:96
	_go_fuzz_dep_.CoverTab[85111]++
							info, err := getTypeInfo(ut)
							if err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:98
		_go_fuzz_dep_.CoverTab[85118]++
								enc.setError(err)
								return
//line /usr/local/go/src/encoding/gob/encoder.go:100
		// _ = "end of CoverTab[85118]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:101
		_go_fuzz_dep_.CoverTab[85119]++
//line /usr/local/go/src/encoding/gob/encoder.go:101
		// _ = "end of CoverTab[85119]"
//line /usr/local/go/src/encoding/gob/encoder.go:101
	}
//line /usr/local/go/src/encoding/gob/encoder.go:101
	// _ = "end of CoverTab[85111]"
//line /usr/local/go/src/encoding/gob/encoder.go:101
	_go_fuzz_dep_.CoverTab[85112]++

//line /usr/local/go/src/encoding/gob/encoder.go:104
	state.encodeInt(-int64(info.id))

	enc.encode(state.b, reflect.ValueOf(info.wire), wireTypeUserInfo)
	enc.writeMessage(w, state.b)
	if enc.err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:108
		_go_fuzz_dep_.CoverTab[85120]++
								return
//line /usr/local/go/src/encoding/gob/encoder.go:109
		// _ = "end of CoverTab[85120]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:110
		_go_fuzz_dep_.CoverTab[85121]++
//line /usr/local/go/src/encoding/gob/encoder.go:110
		// _ = "end of CoverTab[85121]"
//line /usr/local/go/src/encoding/gob/encoder.go:110
	}
//line /usr/local/go/src/encoding/gob/encoder.go:110
	// _ = "end of CoverTab[85112]"
//line /usr/local/go/src/encoding/gob/encoder.go:110
	_go_fuzz_dep_.CoverTab[85113]++

//line /usr/local/go/src/encoding/gob/encoder.go:113
	enc.sent[ut.base] = info.id
	if ut.user != ut.base {
//line /usr/local/go/src/encoding/gob/encoder.go:114
		_go_fuzz_dep_.CoverTab[85122]++
								enc.sent[ut.user] = info.id
//line /usr/local/go/src/encoding/gob/encoder.go:115
		// _ = "end of CoverTab[85122]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:116
		_go_fuzz_dep_.CoverTab[85123]++
//line /usr/local/go/src/encoding/gob/encoder.go:116
		// _ = "end of CoverTab[85123]"
//line /usr/local/go/src/encoding/gob/encoder.go:116
	}
//line /usr/local/go/src/encoding/gob/encoder.go:116
	// _ = "end of CoverTab[85113]"
//line /usr/local/go/src/encoding/gob/encoder.go:116
	_go_fuzz_dep_.CoverTab[85114]++

							switch st := actual; st.Kind() {
	case reflect.Struct:
//line /usr/local/go/src/encoding/gob/encoder.go:119
		_go_fuzz_dep_.CoverTab[85124]++
								for i := 0; i < st.NumField(); i++ {
//line /usr/local/go/src/encoding/gob/encoder.go:120
			_go_fuzz_dep_.CoverTab[85128]++
									if isExported(st.Field(i).Name) {
//line /usr/local/go/src/encoding/gob/encoder.go:121
				_go_fuzz_dep_.CoverTab[85129]++
										enc.sendType(w, state, st.Field(i).Type)
//line /usr/local/go/src/encoding/gob/encoder.go:122
				// _ = "end of CoverTab[85129]"
			} else {
//line /usr/local/go/src/encoding/gob/encoder.go:123
				_go_fuzz_dep_.CoverTab[85130]++
//line /usr/local/go/src/encoding/gob/encoder.go:123
				// _ = "end of CoverTab[85130]"
//line /usr/local/go/src/encoding/gob/encoder.go:123
			}
//line /usr/local/go/src/encoding/gob/encoder.go:123
			// _ = "end of CoverTab[85128]"
		}
//line /usr/local/go/src/encoding/gob/encoder.go:124
		// _ = "end of CoverTab[85124]"
	case reflect.Array, reflect.Slice:
//line /usr/local/go/src/encoding/gob/encoder.go:125
		_go_fuzz_dep_.CoverTab[85125]++
								enc.sendType(w, state, st.Elem())
//line /usr/local/go/src/encoding/gob/encoder.go:126
		// _ = "end of CoverTab[85125]"
	case reflect.Map:
//line /usr/local/go/src/encoding/gob/encoder.go:127
		_go_fuzz_dep_.CoverTab[85126]++
								enc.sendType(w, state, st.Key())
								enc.sendType(w, state, st.Elem())
//line /usr/local/go/src/encoding/gob/encoder.go:129
		// _ = "end of CoverTab[85126]"
//line /usr/local/go/src/encoding/gob/encoder.go:129
	default:
//line /usr/local/go/src/encoding/gob/encoder.go:129
		_go_fuzz_dep_.CoverTab[85127]++
//line /usr/local/go/src/encoding/gob/encoder.go:129
		// _ = "end of CoverTab[85127]"
	}
//line /usr/local/go/src/encoding/gob/encoder.go:130
	// _ = "end of CoverTab[85114]"
//line /usr/local/go/src/encoding/gob/encoder.go:130
	_go_fuzz_dep_.CoverTab[85115]++
							return true
//line /usr/local/go/src/encoding/gob/encoder.go:131
	// _ = "end of CoverTab[85115]"
}

// sendType sends the type info to the other side, if necessary.
func (enc *Encoder) sendType(w io.Writer, state *encoderState, origt reflect.Type) (sent bool) {
//line /usr/local/go/src/encoding/gob/encoder.go:135
	_go_fuzz_dep_.CoverTab[85131]++
							ut := userType(origt)
							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/encoder.go:137
		_go_fuzz_dep_.CoverTab[85134]++

//line /usr/local/go/src/encoding/gob/encoder.go:140
		return enc.sendActualType(w, state, ut, ut.base)
//line /usr/local/go/src/encoding/gob/encoder.go:140
		// _ = "end of CoverTab[85134]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:141
		_go_fuzz_dep_.CoverTab[85135]++
//line /usr/local/go/src/encoding/gob/encoder.go:141
		// _ = "end of CoverTab[85135]"
//line /usr/local/go/src/encoding/gob/encoder.go:141
	}
//line /usr/local/go/src/encoding/gob/encoder.go:141
	// _ = "end of CoverTab[85131]"
//line /usr/local/go/src/encoding/gob/encoder.go:141
	_go_fuzz_dep_.CoverTab[85132]++

//line /usr/local/go/src/encoding/gob/encoder.go:144
	switch rt := ut.base; rt.Kind() {
	default:
//line /usr/local/go/src/encoding/gob/encoder.go:145
		_go_fuzz_dep_.CoverTab[85136]++

								return
//line /usr/local/go/src/encoding/gob/encoder.go:147
		// _ = "end of CoverTab[85136]"
	case reflect.Slice:
//line /usr/local/go/src/encoding/gob/encoder.go:148
		_go_fuzz_dep_.CoverTab[85137]++

								if rt.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/gob/encoder.go:150
			_go_fuzz_dep_.CoverTab[85143]++
									return
//line /usr/local/go/src/encoding/gob/encoder.go:151
			// _ = "end of CoverTab[85143]"
		} else {
//line /usr/local/go/src/encoding/gob/encoder.go:152
			_go_fuzz_dep_.CoverTab[85144]++
//line /usr/local/go/src/encoding/gob/encoder.go:152
			// _ = "end of CoverTab[85144]"
//line /usr/local/go/src/encoding/gob/encoder.go:152
		}
//line /usr/local/go/src/encoding/gob/encoder.go:152
		// _ = "end of CoverTab[85137]"
//line /usr/local/go/src/encoding/gob/encoder.go:152
		_go_fuzz_dep_.CoverTab[85138]++

								break
//line /usr/local/go/src/encoding/gob/encoder.go:154
		// _ = "end of CoverTab[85138]"
	case reflect.Array:
//line /usr/local/go/src/encoding/gob/encoder.go:155
		_go_fuzz_dep_.CoverTab[85139]++

								break
//line /usr/local/go/src/encoding/gob/encoder.go:157
		// _ = "end of CoverTab[85139]"
	case reflect.Map:
//line /usr/local/go/src/encoding/gob/encoder.go:158
		_go_fuzz_dep_.CoverTab[85140]++

								break
//line /usr/local/go/src/encoding/gob/encoder.go:160
		// _ = "end of CoverTab[85140]"
	case reflect.Struct:
//line /usr/local/go/src/encoding/gob/encoder.go:161
		_go_fuzz_dep_.CoverTab[85141]++

								break
//line /usr/local/go/src/encoding/gob/encoder.go:163
		// _ = "end of CoverTab[85141]"
	case reflect.Chan, reflect.Func:
//line /usr/local/go/src/encoding/gob/encoder.go:164
		_go_fuzz_dep_.CoverTab[85142]++

								return
//line /usr/local/go/src/encoding/gob/encoder.go:166
		// _ = "end of CoverTab[85142]"
	}
//line /usr/local/go/src/encoding/gob/encoder.go:167
	// _ = "end of CoverTab[85132]"
//line /usr/local/go/src/encoding/gob/encoder.go:167
	_go_fuzz_dep_.CoverTab[85133]++

							return enc.sendActualType(w, state, ut, ut.base)
//line /usr/local/go/src/encoding/gob/encoder.go:169
	// _ = "end of CoverTab[85133]"
}

// Encode transmits the data item represented by the empty interface value,
//line /usr/local/go/src/encoding/gob/encoder.go:172
// guaranteeing that all necessary type information has been transmitted first.
//line /usr/local/go/src/encoding/gob/encoder.go:172
// Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.
//line /usr/local/go/src/encoding/gob/encoder.go:175
func (enc *Encoder) Encode(e any) error {
//line /usr/local/go/src/encoding/gob/encoder.go:175
	_go_fuzz_dep_.CoverTab[85145]++
							return enc.EncodeValue(reflect.ValueOf(e))
//line /usr/local/go/src/encoding/gob/encoder.go:176
	// _ = "end of CoverTab[85145]"
}

// sendTypeDescriptor makes sure the remote side knows about this type.
//line /usr/local/go/src/encoding/gob/encoder.go:179
// It will send a descriptor if this is the first time the type has been
//line /usr/local/go/src/encoding/gob/encoder.go:179
// sent.
//line /usr/local/go/src/encoding/gob/encoder.go:182
func (enc *Encoder) sendTypeDescriptor(w io.Writer, state *encoderState, ut *userTypeInfo) {
//line /usr/local/go/src/encoding/gob/encoder.go:182
	_go_fuzz_dep_.CoverTab[85146]++

//line /usr/local/go/src/encoding/gob/encoder.go:185
	rt := ut.base
	if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/encoder.go:186
		_go_fuzz_dep_.CoverTab[85148]++
								rt = ut.user
//line /usr/local/go/src/encoding/gob/encoder.go:187
		// _ = "end of CoverTab[85148]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:188
		_go_fuzz_dep_.CoverTab[85149]++
//line /usr/local/go/src/encoding/gob/encoder.go:188
		// _ = "end of CoverTab[85149]"
//line /usr/local/go/src/encoding/gob/encoder.go:188
	}
//line /usr/local/go/src/encoding/gob/encoder.go:188
	// _ = "end of CoverTab[85146]"
//line /usr/local/go/src/encoding/gob/encoder.go:188
	_go_fuzz_dep_.CoverTab[85147]++
							if _, alreadySent := enc.sent[rt]; !alreadySent {
//line /usr/local/go/src/encoding/gob/encoder.go:189
		_go_fuzz_dep_.CoverTab[85150]++

								sent := enc.sendType(w, state, rt)
								if enc.err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:192
			_go_fuzz_dep_.CoverTab[85152]++
									return
//line /usr/local/go/src/encoding/gob/encoder.go:193
			// _ = "end of CoverTab[85152]"
		} else {
//line /usr/local/go/src/encoding/gob/encoder.go:194
			_go_fuzz_dep_.CoverTab[85153]++
//line /usr/local/go/src/encoding/gob/encoder.go:194
			// _ = "end of CoverTab[85153]"
//line /usr/local/go/src/encoding/gob/encoder.go:194
		}
//line /usr/local/go/src/encoding/gob/encoder.go:194
		// _ = "end of CoverTab[85150]"
//line /usr/local/go/src/encoding/gob/encoder.go:194
		_go_fuzz_dep_.CoverTab[85151]++

//line /usr/local/go/src/encoding/gob/encoder.go:198
		if !sent {
//line /usr/local/go/src/encoding/gob/encoder.go:198
			_go_fuzz_dep_.CoverTab[85154]++
									info, err := getTypeInfo(ut)
									if err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:200
				_go_fuzz_dep_.CoverTab[85156]++
										enc.setError(err)
										return
//line /usr/local/go/src/encoding/gob/encoder.go:202
				// _ = "end of CoverTab[85156]"
			} else {
//line /usr/local/go/src/encoding/gob/encoder.go:203
				_go_fuzz_dep_.CoverTab[85157]++
//line /usr/local/go/src/encoding/gob/encoder.go:203
				// _ = "end of CoverTab[85157]"
//line /usr/local/go/src/encoding/gob/encoder.go:203
			}
//line /usr/local/go/src/encoding/gob/encoder.go:203
			// _ = "end of CoverTab[85154]"
//line /usr/local/go/src/encoding/gob/encoder.go:203
			_go_fuzz_dep_.CoverTab[85155]++
									enc.sent[rt] = info.id
//line /usr/local/go/src/encoding/gob/encoder.go:204
			// _ = "end of CoverTab[85155]"
		} else {
//line /usr/local/go/src/encoding/gob/encoder.go:205
			_go_fuzz_dep_.CoverTab[85158]++
//line /usr/local/go/src/encoding/gob/encoder.go:205
			// _ = "end of CoverTab[85158]"
//line /usr/local/go/src/encoding/gob/encoder.go:205
		}
//line /usr/local/go/src/encoding/gob/encoder.go:205
		// _ = "end of CoverTab[85151]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:206
		_go_fuzz_dep_.CoverTab[85159]++
//line /usr/local/go/src/encoding/gob/encoder.go:206
		// _ = "end of CoverTab[85159]"
//line /usr/local/go/src/encoding/gob/encoder.go:206
	}
//line /usr/local/go/src/encoding/gob/encoder.go:206
	// _ = "end of CoverTab[85147]"
}

// sendTypeId sends the id, which must have already been defined.
func (enc *Encoder) sendTypeId(state *encoderState, ut *userTypeInfo) {
//line /usr/local/go/src/encoding/gob/encoder.go:210
	_go_fuzz_dep_.CoverTab[85160]++

							state.encodeInt(int64(enc.sent[ut.base]))
//line /usr/local/go/src/encoding/gob/encoder.go:212
	// _ = "end of CoverTab[85160]"
}

// EncodeValue transmits the data item represented by the reflection value,
//line /usr/local/go/src/encoding/gob/encoder.go:215
// guaranteeing that all necessary type information has been transmitted first.
//line /usr/local/go/src/encoding/gob/encoder.go:215
// Passing a nil pointer to EncodeValue will panic, as they cannot be transmitted by gob.
//line /usr/local/go/src/encoding/gob/encoder.go:218
func (enc *Encoder) EncodeValue(value reflect.Value) error {
//line /usr/local/go/src/encoding/gob/encoder.go:218
	_go_fuzz_dep_.CoverTab[85161]++
							if value.Kind() == reflect.Invalid {
//line /usr/local/go/src/encoding/gob/encoder.go:219
		_go_fuzz_dep_.CoverTab[85167]++
								return errors.New("gob: cannot encode nil value")
//line /usr/local/go/src/encoding/gob/encoder.go:220
		// _ = "end of CoverTab[85167]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:221
		_go_fuzz_dep_.CoverTab[85168]++
//line /usr/local/go/src/encoding/gob/encoder.go:221
		// _ = "end of CoverTab[85168]"
//line /usr/local/go/src/encoding/gob/encoder.go:221
	}
//line /usr/local/go/src/encoding/gob/encoder.go:221
	// _ = "end of CoverTab[85161]"
//line /usr/local/go/src/encoding/gob/encoder.go:221
	_go_fuzz_dep_.CoverTab[85162]++
							if value.Kind() == reflect.Pointer && func() bool {
//line /usr/local/go/src/encoding/gob/encoder.go:222
		_go_fuzz_dep_.CoverTab[85169]++
//line /usr/local/go/src/encoding/gob/encoder.go:222
		return value.IsNil()
//line /usr/local/go/src/encoding/gob/encoder.go:222
		// _ = "end of CoverTab[85169]"
//line /usr/local/go/src/encoding/gob/encoder.go:222
	}() {
//line /usr/local/go/src/encoding/gob/encoder.go:222
		_go_fuzz_dep_.CoverTab[85170]++
								panic("gob: cannot encode nil pointer of type " + value.Type().String())
//line /usr/local/go/src/encoding/gob/encoder.go:223
		// _ = "end of CoverTab[85170]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:224
		_go_fuzz_dep_.CoverTab[85171]++
//line /usr/local/go/src/encoding/gob/encoder.go:224
		// _ = "end of CoverTab[85171]"
//line /usr/local/go/src/encoding/gob/encoder.go:224
	}
//line /usr/local/go/src/encoding/gob/encoder.go:224
	// _ = "end of CoverTab[85162]"
//line /usr/local/go/src/encoding/gob/encoder.go:224
	_go_fuzz_dep_.CoverTab[85163]++

//line /usr/local/go/src/encoding/gob/encoder.go:228
	enc.mutex.Lock()
							defer enc.mutex.Unlock()

//line /usr/local/go/src/encoding/gob/encoder.go:232
	enc.w = enc.w[0:1]

	ut, err := validUserType(value.Type())
	if err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:235
		_go_fuzz_dep_.CoverTab[85172]++
								return err
//line /usr/local/go/src/encoding/gob/encoder.go:236
		// _ = "end of CoverTab[85172]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:237
		_go_fuzz_dep_.CoverTab[85173]++
//line /usr/local/go/src/encoding/gob/encoder.go:237
		// _ = "end of CoverTab[85173]"
//line /usr/local/go/src/encoding/gob/encoder.go:237
	}
//line /usr/local/go/src/encoding/gob/encoder.go:237
	// _ = "end of CoverTab[85163]"
//line /usr/local/go/src/encoding/gob/encoder.go:237
	_go_fuzz_dep_.CoverTab[85164]++

							enc.err = nil
							enc.byteBuf.Reset()
							enc.byteBuf.Write(spaceForLength)
							state := enc.newEncoderState(&enc.byteBuf)

							enc.sendTypeDescriptor(enc.writer(), state, ut)
							enc.sendTypeId(state, ut)
							if enc.err != nil {
//line /usr/local/go/src/encoding/gob/encoder.go:246
		_go_fuzz_dep_.CoverTab[85174]++
								return enc.err
//line /usr/local/go/src/encoding/gob/encoder.go:247
		// _ = "end of CoverTab[85174]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:248
		_go_fuzz_dep_.CoverTab[85175]++
//line /usr/local/go/src/encoding/gob/encoder.go:248
		// _ = "end of CoverTab[85175]"
//line /usr/local/go/src/encoding/gob/encoder.go:248
	}
//line /usr/local/go/src/encoding/gob/encoder.go:248
	// _ = "end of CoverTab[85164]"
//line /usr/local/go/src/encoding/gob/encoder.go:248
	_go_fuzz_dep_.CoverTab[85165]++

//line /usr/local/go/src/encoding/gob/encoder.go:251
	enc.encode(state.b, value, ut)
	if enc.err == nil {
//line /usr/local/go/src/encoding/gob/encoder.go:252
		_go_fuzz_dep_.CoverTab[85176]++
								enc.writeMessage(enc.writer(), state.b)
//line /usr/local/go/src/encoding/gob/encoder.go:253
		// _ = "end of CoverTab[85176]"
	} else {
//line /usr/local/go/src/encoding/gob/encoder.go:254
		_go_fuzz_dep_.CoverTab[85177]++
//line /usr/local/go/src/encoding/gob/encoder.go:254
		// _ = "end of CoverTab[85177]"
//line /usr/local/go/src/encoding/gob/encoder.go:254
	}
//line /usr/local/go/src/encoding/gob/encoder.go:254
	// _ = "end of CoverTab[85165]"
//line /usr/local/go/src/encoding/gob/encoder.go:254
	_go_fuzz_dep_.CoverTab[85166]++

							enc.freeEncoderState(state)
							return enc.err
//line /usr/local/go/src/encoding/gob/encoder.go:257
	// _ = "end of CoverTab[85166]"
}

//line /usr/local/go/src/encoding/gob/encoder.go:258
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/gob/encoder.go:258
var _ = _go_fuzz_dep_.CoverTab
