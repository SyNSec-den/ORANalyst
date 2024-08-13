// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:21
)

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"sync"
	"time"
	"unicode/utf8"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/internal/bufferpool"
)

// For JSON-escaping; see jsonEncoder.safeAddString below.
const _hex = "0123456789abcdef"

var _jsonPool = sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:38
	_go_fuzz_dep_.CoverTab[130915]++
											return &jsonEncoder{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:39
	// _ = "end of CoverTab[130915]"
}}

func getJSONEncoder() *jsonEncoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:42
	_go_fuzz_dep_.CoverTab[130916]++
											return _jsonPool.Get().(*jsonEncoder)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:43
	// _ = "end of CoverTab[130916]"
}

func putJSONEncoder(enc *jsonEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:46
	_go_fuzz_dep_.CoverTab[130917]++
											if enc.reflectBuf != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:47
		_go_fuzz_dep_.CoverTab[130919]++
												enc.reflectBuf.Free()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:48
		// _ = "end of CoverTab[130919]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:49
		_go_fuzz_dep_.CoverTab[130920]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:49
		// _ = "end of CoverTab[130920]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:49
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:49
	// _ = "end of CoverTab[130917]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:49
	_go_fuzz_dep_.CoverTab[130918]++
											enc.EncoderConfig = nil
											enc.buf = nil
											enc.spaced = false
											enc.openNamespaces = 0
											enc.reflectBuf = nil
											enc.reflectEnc = nil
											_jsonPool.Put(enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:56
	// _ = "end of CoverTab[130918]"
}

type jsonEncoder struct {
	*EncoderConfig
	buf		*buffer.Buffer
	spaced		bool	// include spaces after colons and commas
	openNamespaces	int

	// for encoding generic values by reflection
	reflectBuf	*buffer.Buffer
	reflectEnc	*json.Encoder
}

// NewJSONEncoder creates a fast, low-allocation JSON encoder. The encoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// appropriately escapes all field keys and values.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// Note that the encoder doesn't deduplicate keys, so it's possible to produce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// a message like
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
//	{"foo":"bar","foo":"baz"}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// This is permitted by the JSON specification, but not encouraged. Many
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// libraries will ignore duplicate key-value pairs (typically keeping the last
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// pair) when unmarshaling, but users should attempt to avoid adding duplicate
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:70
// keys.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:80
func NewJSONEncoder(cfg EncoderConfig) Encoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:80
	_go_fuzz_dep_.CoverTab[130921]++
											return newJSONEncoder(cfg, false)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:81
	// _ = "end of CoverTab[130921]"
}

func newJSONEncoder(cfg EncoderConfig, spaced bool) *jsonEncoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:84
	_go_fuzz_dep_.CoverTab[130922]++
											return &jsonEncoder{
		EncoderConfig:	&cfg,
		buf:		bufferpool.Get(),
		spaced:		spaced,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:89
	// _ = "end of CoverTab[130922]"
}

func (enc *jsonEncoder) AddArray(key string, arr ArrayMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:92
	_go_fuzz_dep_.CoverTab[130923]++
											enc.addKey(key)
											return enc.AppendArray(arr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:94
	// _ = "end of CoverTab[130923]"
}

func (enc *jsonEncoder) AddObject(key string, obj ObjectMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:97
	_go_fuzz_dep_.CoverTab[130924]++
											enc.addKey(key)
											return enc.AppendObject(obj)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:99
	// _ = "end of CoverTab[130924]"
}

func (enc *jsonEncoder) AddBinary(key string, val []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:102
	_go_fuzz_dep_.CoverTab[130925]++
											enc.AddString(key, base64.StdEncoding.EncodeToString(val))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:103
	// _ = "end of CoverTab[130925]"
}

func (enc *jsonEncoder) AddByteString(key string, val []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:106
	_go_fuzz_dep_.CoverTab[130926]++
											enc.addKey(key)
											enc.AppendByteString(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:108
	// _ = "end of CoverTab[130926]"
}

func (enc *jsonEncoder) AddBool(key string, val bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:111
	_go_fuzz_dep_.CoverTab[130927]++
											enc.addKey(key)
											enc.AppendBool(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:113
	// _ = "end of CoverTab[130927]"
}

func (enc *jsonEncoder) AddComplex128(key string, val complex128) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:116
	_go_fuzz_dep_.CoverTab[130928]++
											enc.addKey(key)
											enc.AppendComplex128(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:118
	// _ = "end of CoverTab[130928]"
}

func (enc *jsonEncoder) AddDuration(key string, val time.Duration) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:121
	_go_fuzz_dep_.CoverTab[130929]++
											enc.addKey(key)
											enc.AppendDuration(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:123
	// _ = "end of CoverTab[130929]"
}

func (enc *jsonEncoder) AddFloat64(key string, val float64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:126
	_go_fuzz_dep_.CoverTab[130930]++
											enc.addKey(key)
											enc.AppendFloat64(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:128
	// _ = "end of CoverTab[130930]"
}

func (enc *jsonEncoder) AddInt64(key string, val int64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:131
	_go_fuzz_dep_.CoverTab[130931]++
											enc.addKey(key)
											enc.AppendInt64(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:133
	// _ = "end of CoverTab[130931]"
}

func (enc *jsonEncoder) resetReflectBuf() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:136
	_go_fuzz_dep_.CoverTab[130932]++
											if enc.reflectBuf == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:137
		_go_fuzz_dep_.CoverTab[130933]++
												enc.reflectBuf = bufferpool.Get()
												enc.reflectEnc = json.NewEncoder(enc.reflectBuf)

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:142
		enc.reflectEnc.SetEscapeHTML(false)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:142
		// _ = "end of CoverTab[130933]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:143
		_go_fuzz_dep_.CoverTab[130934]++
												enc.reflectBuf.Reset()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:144
		// _ = "end of CoverTab[130934]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:145
	// _ = "end of CoverTab[130932]"
}

var nullLiteralBytes = []byte("null")

// Only invoke the standard JSON encoder if there is actually something to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:150
// encode; otherwise write JSON null literal directly.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:152
func (enc *jsonEncoder) encodeReflected(obj interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:152
	_go_fuzz_dep_.CoverTab[130935]++
											if obj == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:153
		_go_fuzz_dep_.CoverTab[130938]++
												return nullLiteralBytes, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:154
		// _ = "end of CoverTab[130938]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:155
		_go_fuzz_dep_.CoverTab[130939]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:155
		// _ = "end of CoverTab[130939]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:155
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:155
	// _ = "end of CoverTab[130935]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:155
	_go_fuzz_dep_.CoverTab[130936]++
											enc.resetReflectBuf()
											if err := enc.reflectEnc.Encode(obj); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:157
		_go_fuzz_dep_.CoverTab[130940]++
												return nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:158
		// _ = "end of CoverTab[130940]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:159
		_go_fuzz_dep_.CoverTab[130941]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:159
		// _ = "end of CoverTab[130941]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:159
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:159
	// _ = "end of CoverTab[130936]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:159
	_go_fuzz_dep_.CoverTab[130937]++
											enc.reflectBuf.TrimNewline()
											return enc.reflectBuf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:161
	// _ = "end of CoverTab[130937]"
}

func (enc *jsonEncoder) AddReflected(key string, obj interface{}) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:164
	_go_fuzz_dep_.CoverTab[130942]++
											valueBytes, err := enc.encodeReflected(obj)
											if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:166
		_go_fuzz_dep_.CoverTab[130944]++
												return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:167
		// _ = "end of CoverTab[130944]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:168
		_go_fuzz_dep_.CoverTab[130945]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:168
		// _ = "end of CoverTab[130945]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:168
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:168
	// _ = "end of CoverTab[130942]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:168
	_go_fuzz_dep_.CoverTab[130943]++
											enc.addKey(key)
											_, err = enc.buf.Write(valueBytes)
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:171
	// _ = "end of CoverTab[130943]"
}

func (enc *jsonEncoder) OpenNamespace(key string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:174
	_go_fuzz_dep_.CoverTab[130946]++
											enc.addKey(key)
											enc.buf.AppendByte('{')
											enc.openNamespaces++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:177
	// _ = "end of CoverTab[130946]"
}

func (enc *jsonEncoder) AddString(key, val string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:180
	_go_fuzz_dep_.CoverTab[130947]++
											enc.addKey(key)
											enc.AppendString(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:182
	// _ = "end of CoverTab[130947]"
}

func (enc *jsonEncoder) AddTime(key string, val time.Time) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:185
	_go_fuzz_dep_.CoverTab[130948]++
											enc.addKey(key)
											enc.AppendTime(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:187
	// _ = "end of CoverTab[130948]"
}

func (enc *jsonEncoder) AddUint64(key string, val uint64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:190
	_go_fuzz_dep_.CoverTab[130949]++
											enc.addKey(key)
											enc.AppendUint64(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:192
	// _ = "end of CoverTab[130949]"
}

func (enc *jsonEncoder) AppendArray(arr ArrayMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:195
	_go_fuzz_dep_.CoverTab[130950]++
											enc.addElementSeparator()
											enc.buf.AppendByte('[')
											err := arr.MarshalLogArray(enc)
											enc.buf.AppendByte(']')
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:200
	// _ = "end of CoverTab[130950]"
}

func (enc *jsonEncoder) AppendObject(obj ObjectMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:203
	_go_fuzz_dep_.CoverTab[130951]++
											enc.addElementSeparator()
											enc.buf.AppendByte('{')
											err := obj.MarshalLogObject(enc)
											enc.buf.AppendByte('}')
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:208
	// _ = "end of CoverTab[130951]"
}

func (enc *jsonEncoder) AppendBool(val bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:211
	_go_fuzz_dep_.CoverTab[130952]++
											enc.addElementSeparator()
											enc.buf.AppendBool(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:213
	// _ = "end of CoverTab[130952]"
}

func (enc *jsonEncoder) AppendByteString(val []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:216
	_go_fuzz_dep_.CoverTab[130953]++
											enc.addElementSeparator()
											enc.buf.AppendByte('"')
											enc.safeAddByteString(val)
											enc.buf.AppendByte('"')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:220
	// _ = "end of CoverTab[130953]"
}

func (enc *jsonEncoder) AppendComplex128(val complex128) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:223
	_go_fuzz_dep_.CoverTab[130954]++
											enc.addElementSeparator()

											r, i := float64(real(val)), float64(imag(val))
											enc.buf.AppendByte('"')

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:230
	enc.buf.AppendFloat(r, 64)
											enc.buf.AppendByte('+')
											enc.buf.AppendFloat(i, 64)
											enc.buf.AppendByte('i')
											enc.buf.AppendByte('"')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:234
	// _ = "end of CoverTab[130954]"
}

func (enc *jsonEncoder) AppendDuration(val time.Duration) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:237
	_go_fuzz_dep_.CoverTab[130955]++
											cur := enc.buf.Len()
											if e := enc.EncodeDuration; e != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:239
		_go_fuzz_dep_.CoverTab[130957]++
												e(val, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:240
		// _ = "end of CoverTab[130957]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:241
		_go_fuzz_dep_.CoverTab[130958]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:241
		// _ = "end of CoverTab[130958]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:241
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:241
	// _ = "end of CoverTab[130955]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:241
	_go_fuzz_dep_.CoverTab[130956]++
											if cur == enc.buf.Len() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:242
		_go_fuzz_dep_.CoverTab[130959]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:245
		enc.AppendInt64(int64(val))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:245
		// _ = "end of CoverTab[130959]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:246
		_go_fuzz_dep_.CoverTab[130960]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:246
		// _ = "end of CoverTab[130960]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:246
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:246
	// _ = "end of CoverTab[130956]"
}

func (enc *jsonEncoder) AppendInt64(val int64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:249
	_go_fuzz_dep_.CoverTab[130961]++
											enc.addElementSeparator()
											enc.buf.AppendInt(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:251
	// _ = "end of CoverTab[130961]"
}

func (enc *jsonEncoder) AppendReflected(val interface{}) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:254
	_go_fuzz_dep_.CoverTab[130962]++
											valueBytes, err := enc.encodeReflected(val)
											if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:256
		_go_fuzz_dep_.CoverTab[130964]++
												return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:257
		// _ = "end of CoverTab[130964]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:258
		_go_fuzz_dep_.CoverTab[130965]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:258
		// _ = "end of CoverTab[130965]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:258
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:258
	// _ = "end of CoverTab[130962]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:258
	_go_fuzz_dep_.CoverTab[130963]++
											enc.addElementSeparator()
											_, err = enc.buf.Write(valueBytes)
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:261
	// _ = "end of CoverTab[130963]"
}

func (enc *jsonEncoder) AppendString(val string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:264
	_go_fuzz_dep_.CoverTab[130966]++
											enc.addElementSeparator()
											enc.buf.AppendByte('"')
											enc.safeAddString(val)
											enc.buf.AppendByte('"')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:268
	// _ = "end of CoverTab[130966]"
}

func (enc *jsonEncoder) AppendTimeLayout(time time.Time, layout string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:271
	_go_fuzz_dep_.CoverTab[130967]++
											enc.addElementSeparator()
											enc.buf.AppendByte('"')
											enc.buf.AppendTime(time, layout)
											enc.buf.AppendByte('"')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:275
	// _ = "end of CoverTab[130967]"
}

func (enc *jsonEncoder) AppendTime(val time.Time) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:278
	_go_fuzz_dep_.CoverTab[130968]++
											cur := enc.buf.Len()
											if e := enc.EncodeTime; e != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:280
		_go_fuzz_dep_.CoverTab[130970]++
												e(val, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:281
		// _ = "end of CoverTab[130970]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:282
		_go_fuzz_dep_.CoverTab[130971]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:282
		// _ = "end of CoverTab[130971]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:282
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:282
	// _ = "end of CoverTab[130968]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:282
	_go_fuzz_dep_.CoverTab[130969]++
											if cur == enc.buf.Len() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:283
		_go_fuzz_dep_.CoverTab[130972]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:286
		enc.AppendInt64(val.UnixNano())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:286
		// _ = "end of CoverTab[130972]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:287
		_go_fuzz_dep_.CoverTab[130973]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:287
		// _ = "end of CoverTab[130973]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:287
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:287
	// _ = "end of CoverTab[130969]"
}

func (enc *jsonEncoder) AppendUint64(val uint64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:290
	_go_fuzz_dep_.CoverTab[130974]++
											enc.addElementSeparator()
											enc.buf.AppendUint(val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:292
	// _ = "end of CoverTab[130974]"
}

func (enc *jsonEncoder) AddComplex64(k string, v complex64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:295
	_go_fuzz_dep_.CoverTab[130975]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:295
	enc.AddComplex128(k, complex128(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:295
	// _ = "end of CoverTab[130975]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:295
}
func (enc *jsonEncoder) AddFloat32(k string, v float32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:296
	_go_fuzz_dep_.CoverTab[130976]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:296
	enc.AddFloat64(k, float64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:296
	// _ = "end of CoverTab[130976]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:296
}
func (enc *jsonEncoder) AddInt(k string, v int) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:297
	_go_fuzz_dep_.CoverTab[130977]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:297
	enc.AddInt64(k, int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:297
	// _ = "end of CoverTab[130977]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:297
}
func (enc *jsonEncoder) AddInt32(k string, v int32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:298
	_go_fuzz_dep_.CoverTab[130978]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:298
	enc.AddInt64(k, int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:298
	// _ = "end of CoverTab[130978]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:298
}
func (enc *jsonEncoder) AddInt16(k string, v int16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:299
	_go_fuzz_dep_.CoverTab[130979]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:299
	enc.AddInt64(k, int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:299
	// _ = "end of CoverTab[130979]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:299
}
func (enc *jsonEncoder) AddInt8(k string, v int8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:300
	_go_fuzz_dep_.CoverTab[130980]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:300
	enc.AddInt64(k, int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:300
	// _ = "end of CoverTab[130980]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:300
}
func (enc *jsonEncoder) AddUint(k string, v uint) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:301
	_go_fuzz_dep_.CoverTab[130981]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:301
	enc.AddUint64(k, uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:301
	// _ = "end of CoverTab[130981]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:301
}
func (enc *jsonEncoder) AddUint32(k string, v uint32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:302
	_go_fuzz_dep_.CoverTab[130982]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:302
	enc.AddUint64(k, uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:302
	// _ = "end of CoverTab[130982]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:302
}
func (enc *jsonEncoder) AddUint16(k string, v uint16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:303
	_go_fuzz_dep_.CoverTab[130983]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:303
	enc.AddUint64(k, uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:303
	// _ = "end of CoverTab[130983]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:303
}
func (enc *jsonEncoder) AddUint8(k string, v uint8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:304
	_go_fuzz_dep_.CoverTab[130984]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:304
	enc.AddUint64(k, uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:304
	// _ = "end of CoverTab[130984]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:304
}
func (enc *jsonEncoder) AddUintptr(k string, v uintptr) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:305
	_go_fuzz_dep_.CoverTab[130985]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:305
	enc.AddUint64(k, uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:305
	// _ = "end of CoverTab[130985]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:305
}
func (enc *jsonEncoder) AppendComplex64(v complex64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:306
	_go_fuzz_dep_.CoverTab[130986]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:306
	enc.AppendComplex128(complex128(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:306
	// _ = "end of CoverTab[130986]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:306
}
func (enc *jsonEncoder) AppendFloat64(v float64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:307
	_go_fuzz_dep_.CoverTab[130987]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:307
	enc.appendFloat(v, 64)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:307
	// _ = "end of CoverTab[130987]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:307
}
func (enc *jsonEncoder) AppendFloat32(v float32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:308
	_go_fuzz_dep_.CoverTab[130988]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:308
	enc.appendFloat(float64(v), 32)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:308
	// _ = "end of CoverTab[130988]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:308
}
func (enc *jsonEncoder) AppendInt(v int) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:309
	_go_fuzz_dep_.CoverTab[130989]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:309
	enc.AppendInt64(int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:309
	// _ = "end of CoverTab[130989]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:309
}
func (enc *jsonEncoder) AppendInt32(v int32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:310
	_go_fuzz_dep_.CoverTab[130990]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:310
	enc.AppendInt64(int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:310
	// _ = "end of CoverTab[130990]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:310
}
func (enc *jsonEncoder) AppendInt16(v int16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:311
	_go_fuzz_dep_.CoverTab[130991]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:311
	enc.AppendInt64(int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:311
	// _ = "end of CoverTab[130991]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:311
}
func (enc *jsonEncoder) AppendInt8(v int8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:312
	_go_fuzz_dep_.CoverTab[130992]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:312
	enc.AppendInt64(int64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:312
	// _ = "end of CoverTab[130992]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:312
}
func (enc *jsonEncoder) AppendUint(v uint) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:313
	_go_fuzz_dep_.CoverTab[130993]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:313
	enc.AppendUint64(uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:313
	// _ = "end of CoverTab[130993]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:313
}
func (enc *jsonEncoder) AppendUint32(v uint32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:314
	_go_fuzz_dep_.CoverTab[130994]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:314
	enc.AppendUint64(uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:314
	// _ = "end of CoverTab[130994]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:314
}
func (enc *jsonEncoder) AppendUint16(v uint16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:315
	_go_fuzz_dep_.CoverTab[130995]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:315
	enc.AppendUint64(uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:315
	// _ = "end of CoverTab[130995]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:315
}
func (enc *jsonEncoder) AppendUint8(v uint8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:316
	_go_fuzz_dep_.CoverTab[130996]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:316
	enc.AppendUint64(uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:316
	// _ = "end of CoverTab[130996]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:316
}
func (enc *jsonEncoder) AppendUintptr(v uintptr) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:317
	_go_fuzz_dep_.CoverTab[130997]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:317
	enc.AppendUint64(uint64(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:317
	// _ = "end of CoverTab[130997]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:317
}

func (enc *jsonEncoder) Clone() Encoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:319
	_go_fuzz_dep_.CoverTab[130998]++
											clone := enc.clone()
											clone.buf.Write(enc.buf.Bytes())
											return clone
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:322
	// _ = "end of CoverTab[130998]"
}

func (enc *jsonEncoder) clone() *jsonEncoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:325
	_go_fuzz_dep_.CoverTab[130999]++
											clone := getJSONEncoder()
											clone.EncoderConfig = enc.EncoderConfig
											clone.spaced = enc.spaced
											clone.openNamespaces = enc.openNamespaces
											clone.buf = bufferpool.Get()
											return clone
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:331
	// _ = "end of CoverTab[130999]"
}

func (enc *jsonEncoder) EncodeEntry(ent Entry, fields []Field) (*buffer.Buffer, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:334
	_go_fuzz_dep_.CoverTab[131000]++
											final := enc.clone()
											final.buf.AppendByte('{')

											if final.LevelKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:338
		_go_fuzz_dep_.CoverTab[131009]++
												final.addKey(final.LevelKey)
												cur := final.buf.Len()
												final.EncodeLevel(ent.Level, final)
												if cur == final.buf.Len() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:342
			_go_fuzz_dep_.CoverTab[131010]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:345
			final.AppendString(ent.Level.String())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:345
			// _ = "end of CoverTab[131010]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:346
			_go_fuzz_dep_.CoverTab[131011]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:346
			// _ = "end of CoverTab[131011]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:346
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:346
		// _ = "end of CoverTab[131009]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:347
		_go_fuzz_dep_.CoverTab[131012]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:347
		// _ = "end of CoverTab[131012]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:347
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:347
	// _ = "end of CoverTab[131000]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:347
	_go_fuzz_dep_.CoverTab[131001]++
											if final.TimeKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:348
		_go_fuzz_dep_.CoverTab[131013]++
												final.AddTime(final.TimeKey, ent.Time)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:349
		// _ = "end of CoverTab[131013]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:350
		_go_fuzz_dep_.CoverTab[131014]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:350
		// _ = "end of CoverTab[131014]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:350
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:350
	// _ = "end of CoverTab[131001]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:350
	_go_fuzz_dep_.CoverTab[131002]++
											if ent.LoggerName != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:351
		_go_fuzz_dep_.CoverTab[131015]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:351
		return final.NameKey != ""
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:351
		// _ = "end of CoverTab[131015]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:351
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:351
		_go_fuzz_dep_.CoverTab[131016]++
												final.addKey(final.NameKey)
												cur := final.buf.Len()
												nameEncoder := final.EncodeName

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:358
		if nameEncoder == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:358
			_go_fuzz_dep_.CoverTab[131018]++
													nameEncoder = FullNameEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:359
			// _ = "end of CoverTab[131018]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:360
			_go_fuzz_dep_.CoverTab[131019]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:360
			// _ = "end of CoverTab[131019]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:360
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:360
		// _ = "end of CoverTab[131016]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:360
		_go_fuzz_dep_.CoverTab[131017]++

												nameEncoder(ent.LoggerName, final)
												if cur == final.buf.Len() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:363
			_go_fuzz_dep_.CoverTab[131020]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:366
			final.AppendString(ent.LoggerName)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:366
			// _ = "end of CoverTab[131020]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:367
			_go_fuzz_dep_.CoverTab[131021]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:367
			// _ = "end of CoverTab[131021]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:367
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:367
		// _ = "end of CoverTab[131017]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:368
		_go_fuzz_dep_.CoverTab[131022]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:368
		// _ = "end of CoverTab[131022]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:368
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:368
	// _ = "end of CoverTab[131002]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:368
	_go_fuzz_dep_.CoverTab[131003]++
											if ent.Caller.Defined {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:369
		_go_fuzz_dep_.CoverTab[131023]++
												if final.CallerKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:370
			_go_fuzz_dep_.CoverTab[131025]++
													final.addKey(final.CallerKey)
													cur := final.buf.Len()
													final.EncodeCaller(ent.Caller, final)
													if cur == final.buf.Len() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:374
				_go_fuzz_dep_.CoverTab[131026]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:377
				final.AppendString(ent.Caller.String())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:377
				// _ = "end of CoverTab[131026]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:378
				_go_fuzz_dep_.CoverTab[131027]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:378
				// _ = "end of CoverTab[131027]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:378
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:378
			// _ = "end of CoverTab[131025]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:379
			_go_fuzz_dep_.CoverTab[131028]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:379
			// _ = "end of CoverTab[131028]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:379
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:379
		// _ = "end of CoverTab[131023]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:379
		_go_fuzz_dep_.CoverTab[131024]++
												if final.FunctionKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:380
			_go_fuzz_dep_.CoverTab[131029]++
													final.addKey(final.FunctionKey)
													final.AppendString(ent.Caller.Function)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:382
			// _ = "end of CoverTab[131029]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:383
			_go_fuzz_dep_.CoverTab[131030]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:383
			// _ = "end of CoverTab[131030]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:383
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:383
		// _ = "end of CoverTab[131024]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:384
		_go_fuzz_dep_.CoverTab[131031]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:384
		// _ = "end of CoverTab[131031]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:384
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:384
	// _ = "end of CoverTab[131003]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:384
	_go_fuzz_dep_.CoverTab[131004]++
											if final.MessageKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:385
		_go_fuzz_dep_.CoverTab[131032]++
												final.addKey(enc.MessageKey)
												final.AppendString(ent.Message)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:387
		// _ = "end of CoverTab[131032]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:388
		_go_fuzz_dep_.CoverTab[131033]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:388
		// _ = "end of CoverTab[131033]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:388
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:388
	// _ = "end of CoverTab[131004]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:388
	_go_fuzz_dep_.CoverTab[131005]++
											if enc.buf.Len() > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:389
		_go_fuzz_dep_.CoverTab[131034]++
												final.addElementSeparator()
												final.buf.Write(enc.buf.Bytes())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:391
		// _ = "end of CoverTab[131034]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:392
		_go_fuzz_dep_.CoverTab[131035]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:392
		// _ = "end of CoverTab[131035]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:392
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:392
	// _ = "end of CoverTab[131005]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:392
	_go_fuzz_dep_.CoverTab[131006]++
											addFields(final, fields)
											final.closeOpenNamespaces()
											if ent.Stack != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:395
		_go_fuzz_dep_.CoverTab[131036]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:395
		return final.StacktraceKey != ""
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:395
		// _ = "end of CoverTab[131036]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:395
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:395
		_go_fuzz_dep_.CoverTab[131037]++
												final.AddString(final.StacktraceKey, ent.Stack)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:396
		// _ = "end of CoverTab[131037]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:397
		_go_fuzz_dep_.CoverTab[131038]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:397
		// _ = "end of CoverTab[131038]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:397
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:397
	// _ = "end of CoverTab[131006]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:397
	_go_fuzz_dep_.CoverTab[131007]++
											final.buf.AppendByte('}')
											if final.LineEnding != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:399
		_go_fuzz_dep_.CoverTab[131039]++
												final.buf.AppendString(final.LineEnding)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:400
		// _ = "end of CoverTab[131039]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:401
		_go_fuzz_dep_.CoverTab[131040]++
												final.buf.AppendString(DefaultLineEnding)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:402
		// _ = "end of CoverTab[131040]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:403
	// _ = "end of CoverTab[131007]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:403
	_go_fuzz_dep_.CoverTab[131008]++

											ret := final.buf
											putJSONEncoder(final)
											return ret, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:407
	// _ = "end of CoverTab[131008]"
}

func (enc *jsonEncoder) truncate() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:410
	_go_fuzz_dep_.CoverTab[131041]++
											enc.buf.Reset()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:411
	// _ = "end of CoverTab[131041]"
}

func (enc *jsonEncoder) closeOpenNamespaces() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:414
	_go_fuzz_dep_.CoverTab[131042]++
											for i := 0; i < enc.openNamespaces; i++ {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:415
		_go_fuzz_dep_.CoverTab[131043]++
												enc.buf.AppendByte('}')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:416
		// _ = "end of CoverTab[131043]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:417
	// _ = "end of CoverTab[131042]"
}

func (enc *jsonEncoder) addKey(key string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:420
	_go_fuzz_dep_.CoverTab[131044]++
											enc.addElementSeparator()
											enc.buf.AppendByte('"')
											enc.safeAddString(key)
											enc.buf.AppendByte('"')
											enc.buf.AppendByte(':')
											if enc.spaced {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:426
		_go_fuzz_dep_.CoverTab[131045]++
												enc.buf.AppendByte(' ')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:427
		// _ = "end of CoverTab[131045]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:428
		_go_fuzz_dep_.CoverTab[131046]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:428
		// _ = "end of CoverTab[131046]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:428
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:428
	// _ = "end of CoverTab[131044]"
}

func (enc *jsonEncoder) addElementSeparator() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:431
	_go_fuzz_dep_.CoverTab[131047]++
											last := enc.buf.Len() - 1
											if last < 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:433
		_go_fuzz_dep_.CoverTab[131049]++
												return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:434
		// _ = "end of CoverTab[131049]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:435
		_go_fuzz_dep_.CoverTab[131050]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:435
		// _ = "end of CoverTab[131050]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:435
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:435
	// _ = "end of CoverTab[131047]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:435
	_go_fuzz_dep_.CoverTab[131048]++
											switch enc.buf.Bytes()[last] {
	case '{', '[', ':', ',', ' ':
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:437
		_go_fuzz_dep_.CoverTab[131051]++
												return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:438
		// _ = "end of CoverTab[131051]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:439
		_go_fuzz_dep_.CoverTab[131052]++
												enc.buf.AppendByte(',')
												if enc.spaced {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:441
			_go_fuzz_dep_.CoverTab[131053]++
													enc.buf.AppendByte(' ')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:442
			// _ = "end of CoverTab[131053]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:443
			_go_fuzz_dep_.CoverTab[131054]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:443
			// _ = "end of CoverTab[131054]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:443
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:443
		// _ = "end of CoverTab[131052]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:444
	// _ = "end of CoverTab[131048]"
}

func (enc *jsonEncoder) appendFloat(val float64, bitSize int) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:447
	_go_fuzz_dep_.CoverTab[131055]++
											enc.addElementSeparator()
											switch {
	case math.IsNaN(val):
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:450
		_go_fuzz_dep_.CoverTab[131056]++
												enc.buf.AppendString(`"NaN"`)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:451
		// _ = "end of CoverTab[131056]"
	case math.IsInf(val, 1):
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:452
		_go_fuzz_dep_.CoverTab[131057]++
												enc.buf.AppendString(`"+Inf"`)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:453
		// _ = "end of CoverTab[131057]"
	case math.IsInf(val, -1):
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:454
		_go_fuzz_dep_.CoverTab[131058]++
												enc.buf.AppendString(`"-Inf"`)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:455
		// _ = "end of CoverTab[131058]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:456
		_go_fuzz_dep_.CoverTab[131059]++
												enc.buf.AppendFloat(val, bitSize)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:457
		// _ = "end of CoverTab[131059]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:458
	// _ = "end of CoverTab[131055]"
}

// safeAddString JSON-escapes a string and appends it to the internal buffer.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:461
// Unlike the standard library's encoder, it doesn't attempt to protect the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:461
// user from browser vulnerabilities or JSONP-related problems.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:464
func (enc *jsonEncoder) safeAddString(s string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:464
	_go_fuzz_dep_.CoverTab[131060]++
											for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:465
		_go_fuzz_dep_.CoverTab[131061]++
												if enc.tryAddRuneSelf(s[i]) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:466
			_go_fuzz_dep_.CoverTab[131064]++
													i++
													continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:468
			// _ = "end of CoverTab[131064]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:469
			_go_fuzz_dep_.CoverTab[131065]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:469
			// _ = "end of CoverTab[131065]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:469
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:469
		// _ = "end of CoverTab[131061]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:469
		_go_fuzz_dep_.CoverTab[131062]++
												r, size := utf8.DecodeRuneInString(s[i:])
												if enc.tryAddRuneError(r, size) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:471
			_go_fuzz_dep_.CoverTab[131066]++
													i++
													continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:473
			// _ = "end of CoverTab[131066]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:474
			_go_fuzz_dep_.CoverTab[131067]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:474
			// _ = "end of CoverTab[131067]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:474
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:474
		// _ = "end of CoverTab[131062]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:474
		_go_fuzz_dep_.CoverTab[131063]++
												enc.buf.AppendString(s[i : i+size])
												i += size
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:476
		// _ = "end of CoverTab[131063]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:477
	// _ = "end of CoverTab[131060]"
}

// safeAddByteString is no-alloc equivalent of safeAddString(string(s)) for s []byte.
func (enc *jsonEncoder) safeAddByteString(s []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:481
	_go_fuzz_dep_.CoverTab[131068]++
											for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:482
		_go_fuzz_dep_.CoverTab[131069]++
												if enc.tryAddRuneSelf(s[i]) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:483
			_go_fuzz_dep_.CoverTab[131072]++
													i++
													continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:485
			// _ = "end of CoverTab[131072]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:486
			_go_fuzz_dep_.CoverTab[131073]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:486
			// _ = "end of CoverTab[131073]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:486
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:486
		// _ = "end of CoverTab[131069]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:486
		_go_fuzz_dep_.CoverTab[131070]++
												r, size := utf8.DecodeRune(s[i:])
												if enc.tryAddRuneError(r, size) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:488
			_go_fuzz_dep_.CoverTab[131074]++
													i++
													continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:490
			// _ = "end of CoverTab[131074]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:491
			_go_fuzz_dep_.CoverTab[131075]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:491
			// _ = "end of CoverTab[131075]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:491
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:491
		// _ = "end of CoverTab[131070]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:491
		_go_fuzz_dep_.CoverTab[131071]++
												enc.buf.Write(s[i : i+size])
												i += size
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:493
		// _ = "end of CoverTab[131071]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:494
	// _ = "end of CoverTab[131068]"
}

// tryAddRuneSelf appends b if it is valid UTF-8 character represented in a single byte.
func (enc *jsonEncoder) tryAddRuneSelf(b byte) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:498
	_go_fuzz_dep_.CoverTab[131076]++
											if b >= utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:499
		_go_fuzz_dep_.CoverTab[131080]++
												return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:500
		// _ = "end of CoverTab[131080]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:501
		_go_fuzz_dep_.CoverTab[131081]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:501
		// _ = "end of CoverTab[131081]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:501
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:501
	// _ = "end of CoverTab[131076]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:501
	_go_fuzz_dep_.CoverTab[131077]++
											if 0x20 <= b && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		_go_fuzz_dep_.CoverTab[131082]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		return b != '\\'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		// _ = "end of CoverTab[131082]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		_go_fuzz_dep_.CoverTab[131083]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		return b != '"'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		// _ = "end of CoverTab[131083]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:502
		_go_fuzz_dep_.CoverTab[131084]++
												enc.buf.AppendByte(b)
												return true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:504
		// _ = "end of CoverTab[131084]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:505
		_go_fuzz_dep_.CoverTab[131085]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:505
		// _ = "end of CoverTab[131085]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:505
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:505
	// _ = "end of CoverTab[131077]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:505
	_go_fuzz_dep_.CoverTab[131078]++
											switch b {
	case '\\', '"':
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:507
		_go_fuzz_dep_.CoverTab[131086]++
												enc.buf.AppendByte('\\')
												enc.buf.AppendByte(b)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:509
		// _ = "end of CoverTab[131086]"
	case '\n':
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:510
		_go_fuzz_dep_.CoverTab[131087]++
												enc.buf.AppendByte('\\')
												enc.buf.AppendByte('n')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:512
		// _ = "end of CoverTab[131087]"
	case '\r':
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:513
		_go_fuzz_dep_.CoverTab[131088]++
												enc.buf.AppendByte('\\')
												enc.buf.AppendByte('r')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:515
		// _ = "end of CoverTab[131088]"
	case '\t':
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:516
		_go_fuzz_dep_.CoverTab[131089]++
												enc.buf.AppendByte('\\')
												enc.buf.AppendByte('t')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:518
		// _ = "end of CoverTab[131089]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:519
		_go_fuzz_dep_.CoverTab[131090]++

												enc.buf.AppendString(`\u00`)
												enc.buf.AppendByte(_hex[b>>4])
												enc.buf.AppendByte(_hex[b&0xF])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:523
		// _ = "end of CoverTab[131090]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:524
	// _ = "end of CoverTab[131078]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:524
	_go_fuzz_dep_.CoverTab[131079]++
											return true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:525
	// _ = "end of CoverTab[131079]"
}

func (enc *jsonEncoder) tryAddRuneError(r rune, size int) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:528
	_go_fuzz_dep_.CoverTab[131091]++
											if r == utf8.RuneError && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:529
		_go_fuzz_dep_.CoverTab[131093]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:529
		return size == 1
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:529
		// _ = "end of CoverTab[131093]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:529
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:529
		_go_fuzz_dep_.CoverTab[131094]++
												enc.buf.AppendString(`\ufffd`)
												return true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:531
		// _ = "end of CoverTab[131094]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:532
		_go_fuzz_dep_.CoverTab[131095]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:532
		// _ = "end of CoverTab[131095]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:532
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:532
	// _ = "end of CoverTab[131091]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:532
	_go_fuzz_dep_.CoverTab[131092]++
											return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:533
	// _ = "end of CoverTab[131092]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:534
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/json_encoder.go:534
var _ = _go_fuzz_dep_.CoverTab
