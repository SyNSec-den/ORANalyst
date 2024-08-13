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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:21
)

import "time"

// MapObjectEncoder is an ObjectEncoder backed by a simple
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:25
// map[string]interface{}. It's not fast enough for production use, but it's
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:25
// helpful in tests.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:28
type MapObjectEncoder struct {
	// Fields contains the entire encoded log context.
	Fields	map[string]interface{}
	// cur is a pointer to the namespace we're currently writing to.
	cur	map[string]interface{}
}

// NewMapObjectEncoder creates a new map-backed ObjectEncoder.
func NewMapObjectEncoder() *MapObjectEncoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:36
	_go_fuzz_dep_.CoverTab[131138]++
											m := make(map[string]interface{})
											return &MapObjectEncoder{
		Fields:	m,
		cur:	m,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:41
	// _ = "end of CoverTab[131138]"
}

// AddArray implements ObjectEncoder.
func (m *MapObjectEncoder) AddArray(key string, v ArrayMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:45
	_go_fuzz_dep_.CoverTab[131139]++
											arr := &sliceArrayEncoder{elems: make([]interface{}, 0)}
											err := v.MarshalLogArray(arr)
											m.cur[key] = arr.elems
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:49
	// _ = "end of CoverTab[131139]"
}

// AddObject implements ObjectEncoder.
func (m *MapObjectEncoder) AddObject(k string, v ObjectMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:53
	_go_fuzz_dep_.CoverTab[131140]++
											newMap := NewMapObjectEncoder()
											m.cur[k] = newMap.Fields
											return v.MarshalLogObject(newMap)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:56
	// _ = "end of CoverTab[131140]"
}

// AddBinary implements ObjectEncoder.
func (m *MapObjectEncoder) AddBinary(k string, v []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:60
	_go_fuzz_dep_.CoverTab[131141]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:60
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:60
	// _ = "end of CoverTab[131141]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:60
}

// AddByteString implements ObjectEncoder.
func (m *MapObjectEncoder) AddByteString(k string, v []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:63
	_go_fuzz_dep_.CoverTab[131142]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:63
	m.cur[k] = string(v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:63
	// _ = "end of CoverTab[131142]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:63
}

// AddBool implements ObjectEncoder.
func (m *MapObjectEncoder) AddBool(k string, v bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:66
	_go_fuzz_dep_.CoverTab[131143]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:66
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:66
	// _ = "end of CoverTab[131143]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:66
}

// AddDuration implements ObjectEncoder.
func (m MapObjectEncoder) AddDuration(k string, v time.Duration) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:69
	_go_fuzz_dep_.CoverTab[131144]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:69
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:69
	// _ = "end of CoverTab[131144]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:69
}

// AddComplex128 implements ObjectEncoder.
func (m *MapObjectEncoder) AddComplex128(k string, v complex128) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:72
	_go_fuzz_dep_.CoverTab[131145]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:72
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:72
	// _ = "end of CoverTab[131145]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:72
}

// AddComplex64 implements ObjectEncoder.
func (m *MapObjectEncoder) AddComplex64(k string, v complex64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:75
	_go_fuzz_dep_.CoverTab[131146]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:75
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:75
	// _ = "end of CoverTab[131146]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:75
}

// AddFloat64 implements ObjectEncoder.
func (m *MapObjectEncoder) AddFloat64(k string, v float64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:78
	_go_fuzz_dep_.CoverTab[131147]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:78
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:78
	// _ = "end of CoverTab[131147]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:78
}

// AddFloat32 implements ObjectEncoder.
func (m *MapObjectEncoder) AddFloat32(k string, v float32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:81
	_go_fuzz_dep_.CoverTab[131148]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:81
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:81
	// _ = "end of CoverTab[131148]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:81
}

// AddInt implements ObjectEncoder.
func (m *MapObjectEncoder) AddInt(k string, v int) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:84
	_go_fuzz_dep_.CoverTab[131149]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:84
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:84
	// _ = "end of CoverTab[131149]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:84
}

// AddInt64 implements ObjectEncoder.
func (m *MapObjectEncoder) AddInt64(k string, v int64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:87
	_go_fuzz_dep_.CoverTab[131150]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:87
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:87
	// _ = "end of CoverTab[131150]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:87
}

// AddInt32 implements ObjectEncoder.
func (m *MapObjectEncoder) AddInt32(k string, v int32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:90
	_go_fuzz_dep_.CoverTab[131151]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:90
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:90
	// _ = "end of CoverTab[131151]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:90
}

// AddInt16 implements ObjectEncoder.
func (m *MapObjectEncoder) AddInt16(k string, v int16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:93
	_go_fuzz_dep_.CoverTab[131152]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:93
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:93
	// _ = "end of CoverTab[131152]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:93
}

// AddInt8 implements ObjectEncoder.
func (m *MapObjectEncoder) AddInt8(k string, v int8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:96
	_go_fuzz_dep_.CoverTab[131153]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:96
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:96
	// _ = "end of CoverTab[131153]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:96
}

// AddString implements ObjectEncoder.
func (m *MapObjectEncoder) AddString(k string, v string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:99
	_go_fuzz_dep_.CoverTab[131154]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:99
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:99
	// _ = "end of CoverTab[131154]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:99
}

// AddTime implements ObjectEncoder.
func (m MapObjectEncoder) AddTime(k string, v time.Time) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:102
	_go_fuzz_dep_.CoverTab[131155]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:102
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:102
	// _ = "end of CoverTab[131155]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:102
}

// AddUint implements ObjectEncoder.
func (m *MapObjectEncoder) AddUint(k string, v uint) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:105
	_go_fuzz_dep_.CoverTab[131156]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:105
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:105
	// _ = "end of CoverTab[131156]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:105
}

// AddUint64 implements ObjectEncoder.
func (m *MapObjectEncoder) AddUint64(k string, v uint64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:108
	_go_fuzz_dep_.CoverTab[131157]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:108
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:108
	// _ = "end of CoverTab[131157]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:108
}

// AddUint32 implements ObjectEncoder.
func (m *MapObjectEncoder) AddUint32(k string, v uint32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:111
	_go_fuzz_dep_.CoverTab[131158]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:111
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:111
	// _ = "end of CoverTab[131158]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:111
}

// AddUint16 implements ObjectEncoder.
func (m *MapObjectEncoder) AddUint16(k string, v uint16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:114
	_go_fuzz_dep_.CoverTab[131159]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:114
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:114
	// _ = "end of CoverTab[131159]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:114
}

// AddUint8 implements ObjectEncoder.
func (m *MapObjectEncoder) AddUint8(k string, v uint8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:117
	_go_fuzz_dep_.CoverTab[131160]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:117
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:117
	// _ = "end of CoverTab[131160]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:117
}

// AddUintptr implements ObjectEncoder.
func (m *MapObjectEncoder) AddUintptr(k string, v uintptr) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:120
	_go_fuzz_dep_.CoverTab[131161]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:120
	m.cur[k] = v
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:120
	// _ = "end of CoverTab[131161]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:120
}

// AddReflected implements ObjectEncoder.
func (m *MapObjectEncoder) AddReflected(k string, v interface{}) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:123
	_go_fuzz_dep_.CoverTab[131162]++
												m.cur[k] = v
												return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:125
	// _ = "end of CoverTab[131162]"
}

// OpenNamespace implements ObjectEncoder.
func (m *MapObjectEncoder) OpenNamespace(k string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:129
	_go_fuzz_dep_.CoverTab[131163]++
												ns := make(map[string]interface{})
												m.cur[k] = ns
												m.cur = ns
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:132
	// _ = "end of CoverTab[131163]"
}

// sliceArrayEncoder is an ArrayEncoder backed by a simple []interface{}. Like
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:135
// the MapObjectEncoder, it's not designed for production use.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:137
type sliceArrayEncoder struct {
	elems []interface{}
}

func (s *sliceArrayEncoder) AppendArray(v ArrayMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:141
	_go_fuzz_dep_.CoverTab[131164]++
												enc := &sliceArrayEncoder{}
												err := v.MarshalLogArray(enc)
												s.elems = append(s.elems, enc.elems)
												return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:145
	// _ = "end of CoverTab[131164]"
}

func (s *sliceArrayEncoder) AppendObject(v ObjectMarshaler) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:148
	_go_fuzz_dep_.CoverTab[131165]++
												m := NewMapObjectEncoder()
												err := v.MarshalLogObject(m)
												s.elems = append(s.elems, m.Fields)
												return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:152
	// _ = "end of CoverTab[131165]"
}

func (s *sliceArrayEncoder) AppendReflected(v interface{}) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:155
	_go_fuzz_dep_.CoverTab[131166]++
												s.elems = append(s.elems, v)
												return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:157
	// _ = "end of CoverTab[131166]"
}

func (s *sliceArrayEncoder) AppendBool(v bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:160
	_go_fuzz_dep_.CoverTab[131167]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:160
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:160
	// _ = "end of CoverTab[131167]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:160
}
func (s *sliceArrayEncoder) AppendByteString(v []byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:161
	_go_fuzz_dep_.CoverTab[131168]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:161
	s.elems = append(s.elems, string(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:161
	// _ = "end of CoverTab[131168]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:161
}
func (s *sliceArrayEncoder) AppendComplex128(v complex128) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:162
	_go_fuzz_dep_.CoverTab[131169]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:162
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:162
	// _ = "end of CoverTab[131169]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:162
}
func (s *sliceArrayEncoder) AppendComplex64(v complex64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:163
	_go_fuzz_dep_.CoverTab[131170]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:163
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:163
	// _ = "end of CoverTab[131170]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:163
}
func (s *sliceArrayEncoder) AppendDuration(v time.Duration) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:164
	_go_fuzz_dep_.CoverTab[131171]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:164
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:164
	// _ = "end of CoverTab[131171]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:164
}
func (s *sliceArrayEncoder) AppendFloat64(v float64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:165
	_go_fuzz_dep_.CoverTab[131172]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:165
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:165
	// _ = "end of CoverTab[131172]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:165
}
func (s *sliceArrayEncoder) AppendFloat32(v float32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:166
	_go_fuzz_dep_.CoverTab[131173]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:166
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:166
	// _ = "end of CoverTab[131173]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:166
}
func (s *sliceArrayEncoder) AppendInt(v int) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:167
	_go_fuzz_dep_.CoverTab[131174]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:167
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:167
	// _ = "end of CoverTab[131174]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:167
}
func (s *sliceArrayEncoder) AppendInt64(v int64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:168
	_go_fuzz_dep_.CoverTab[131175]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:168
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:168
	// _ = "end of CoverTab[131175]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:168
}
func (s *sliceArrayEncoder) AppendInt32(v int32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:169
	_go_fuzz_dep_.CoverTab[131176]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:169
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:169
	// _ = "end of CoverTab[131176]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:169
}
func (s *sliceArrayEncoder) AppendInt16(v int16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:170
	_go_fuzz_dep_.CoverTab[131177]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:170
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:170
	// _ = "end of CoverTab[131177]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:170
}
func (s *sliceArrayEncoder) AppendInt8(v int8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:171
	_go_fuzz_dep_.CoverTab[131178]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:171
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:171
	// _ = "end of CoverTab[131178]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:171
}
func (s *sliceArrayEncoder) AppendString(v string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:172
	_go_fuzz_dep_.CoverTab[131179]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:172
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:172
	// _ = "end of CoverTab[131179]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:172
}
func (s *sliceArrayEncoder) AppendTime(v time.Time) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:173
	_go_fuzz_dep_.CoverTab[131180]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:173
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:173
	// _ = "end of CoverTab[131180]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:173
}
func (s *sliceArrayEncoder) AppendUint(v uint) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:174
	_go_fuzz_dep_.CoverTab[131181]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:174
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:174
	// _ = "end of CoverTab[131181]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:174
}
func (s *sliceArrayEncoder) AppendUint64(v uint64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:175
	_go_fuzz_dep_.CoverTab[131182]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:175
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:175
	// _ = "end of CoverTab[131182]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:175
}
func (s *sliceArrayEncoder) AppendUint32(v uint32) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:176
	_go_fuzz_dep_.CoverTab[131183]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:176
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:176
	// _ = "end of CoverTab[131183]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:176
}
func (s *sliceArrayEncoder) AppendUint16(v uint16) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:177
	_go_fuzz_dep_.CoverTab[131184]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:177
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:177
	// _ = "end of CoverTab[131184]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:177
}
func (s *sliceArrayEncoder) AppendUint8(v uint8) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:178
	_go_fuzz_dep_.CoverTab[131185]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:178
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:178
	// _ = "end of CoverTab[131185]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:178
}
func (s *sliceArrayEncoder) AppendUintptr(v uintptr) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:179
	_go_fuzz_dep_.CoverTab[131186]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:179
	s.elems = append(s.elems, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:179
	// _ = "end of CoverTab[131186]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:179
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:179
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/memory_encoder.go:179
var _ = _go_fuzz_dep_.CoverTab
