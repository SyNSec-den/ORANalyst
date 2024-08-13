// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2017 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:32
)

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
)

type generatedDiscarder interface {
	XXX_DiscardUnknown()
}

// DiscardUnknown recursively discards all unknown fields from this message
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// and all embedded messages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// When unmarshaling a message with unrecognized fields, the tags and values
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// of such fields are preserved in the Message. This allows a later call to
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// marshal to be able to produce a message that continues to have those
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// unrecognized fields. To avoid this, DiscardUnknown is used to
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// explicitly clear the unknown fields after unmarshaling.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// For proto2 messages, the unknown fields of message extensions are only
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:46
// discarded from messages that have been accessed via GetExtension.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:57
func DiscardUnknown(m Message) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:57
	_go_fuzz_dep_.CoverTab[107803]++
											if m, ok := m.(generatedDiscarder); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:58
		_go_fuzz_dep_.CoverTab[107805]++
												m.XXX_DiscardUnknown()
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:60
		// _ = "end of CoverTab[107805]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:61
		_go_fuzz_dep_.CoverTab[107806]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:61
		// _ = "end of CoverTab[107806]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:61
	// _ = "end of CoverTab[107803]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:61
	_go_fuzz_dep_.CoverTab[107804]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:65
	discardLegacy(m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:65
	// _ = "end of CoverTab[107804]"
}

// DiscardUnknown recursively discards all unknown fields.
func (a *InternalMessageInfo) DiscardUnknown(m Message) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:69
	_go_fuzz_dep_.CoverTab[107807]++
											di := atomicLoadDiscardInfo(&a.discard)
											if di == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:71
		_go_fuzz_dep_.CoverTab[107809]++
												di = getDiscardInfo(reflect.TypeOf(m).Elem())
												atomicStoreDiscardInfo(&a.discard, di)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:73
		// _ = "end of CoverTab[107809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:74
		_go_fuzz_dep_.CoverTab[107810]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:74
		// _ = "end of CoverTab[107810]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:74
	// _ = "end of CoverTab[107807]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:74
	_go_fuzz_dep_.CoverTab[107808]++
											di.discard(toPointer(&m))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:75
	// _ = "end of CoverTab[107808]"
}

type discardInfo struct {
	typ	reflect.Type

	initialized	int32	// 0: only typ is valid, 1: everything is valid
	lock		sync.Mutex

	fields		[]discardFieldInfo
	unrecognized	field
}

type discardFieldInfo struct {
	field	field	// Offset of field, guaranteed to be valid
	discard	func(src pointer)
}

var (
	discardInfoMap	= map[reflect.Type]*discardInfo{}
	discardInfoLock	sync.Mutex
)

func getDiscardInfo(t reflect.Type) *discardInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:98
	_go_fuzz_dep_.CoverTab[107811]++
											discardInfoLock.Lock()
											defer discardInfoLock.Unlock()
											di := discardInfoMap[t]
											if di == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:102
		_go_fuzz_dep_.CoverTab[107813]++
												di = &discardInfo{typ: t}
												discardInfoMap[t] = di
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:104
		// _ = "end of CoverTab[107813]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:105
		_go_fuzz_dep_.CoverTab[107814]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:105
		// _ = "end of CoverTab[107814]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:105
	// _ = "end of CoverTab[107811]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:105
	_go_fuzz_dep_.CoverTab[107812]++
											return di
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:106
	// _ = "end of CoverTab[107812]"
}

func (di *discardInfo) discard(src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:109
	_go_fuzz_dep_.CoverTab[107815]++
											if src.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:110
		_go_fuzz_dep_.CoverTab[107820]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:111
		// _ = "end of CoverTab[107820]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:112
		_go_fuzz_dep_.CoverTab[107821]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:112
		// _ = "end of CoverTab[107821]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:112
	// _ = "end of CoverTab[107815]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:112
	_go_fuzz_dep_.CoverTab[107816]++

											if atomic.LoadInt32(&di.initialized) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:114
		_go_fuzz_dep_.CoverTab[107822]++
												di.computeDiscardInfo()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:115
		// _ = "end of CoverTab[107822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:116
		_go_fuzz_dep_.CoverTab[107823]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:116
		// _ = "end of CoverTab[107823]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:116
	// _ = "end of CoverTab[107816]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:116
	_go_fuzz_dep_.CoverTab[107817]++

											for _, fi := range di.fields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:118
		_go_fuzz_dep_.CoverTab[107824]++
												sfp := src.offset(fi.field)
												fi.discard(sfp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:120
		// _ = "end of CoverTab[107824]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:121
	// _ = "end of CoverTab[107817]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:121
	_go_fuzz_dep_.CoverTab[107818]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:125
	if em, err := extendable(src.asPointerTo(di.typ).Interface()); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:125
		_go_fuzz_dep_.CoverTab[107825]++

												emm, _ := em.extensionsRead()
												for _, mx := range emm {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:128
			_go_fuzz_dep_.CoverTab[107826]++
													if m, ok := mx.value.(Message); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:129
				_go_fuzz_dep_.CoverTab[107827]++
														DiscardUnknown(m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:130
				// _ = "end of CoverTab[107827]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:131
				_go_fuzz_dep_.CoverTab[107828]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:131
				// _ = "end of CoverTab[107828]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:131
			// _ = "end of CoverTab[107826]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:132
		// _ = "end of CoverTab[107825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:133
		_go_fuzz_dep_.CoverTab[107829]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:133
		// _ = "end of CoverTab[107829]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:133
	// _ = "end of CoverTab[107818]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:133
	_go_fuzz_dep_.CoverTab[107819]++

											if di.unrecognized.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:135
		_go_fuzz_dep_.CoverTab[107830]++
												*src.offset(di.unrecognized).toBytes() = nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:136
		// _ = "end of CoverTab[107830]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:137
		_go_fuzz_dep_.CoverTab[107831]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:137
		// _ = "end of CoverTab[107831]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:137
	// _ = "end of CoverTab[107819]"
}

func (di *discardInfo) computeDiscardInfo() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:140
	_go_fuzz_dep_.CoverTab[107832]++
											di.lock.Lock()
											defer di.lock.Unlock()
											if di.initialized != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:143
		_go_fuzz_dep_.CoverTab[107836]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:144
		// _ = "end of CoverTab[107836]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:145
		_go_fuzz_dep_.CoverTab[107837]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:145
		// _ = "end of CoverTab[107837]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:145
	// _ = "end of CoverTab[107832]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:145
	_go_fuzz_dep_.CoverTab[107833]++
											t := di.typ
											n := t.NumField()

											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:149
		_go_fuzz_dep_.CoverTab[107838]++
												f := t.Field(i)
												if strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:151
			_go_fuzz_dep_.CoverTab[107844]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:152
			// _ = "end of CoverTab[107844]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:153
			_go_fuzz_dep_.CoverTab[107845]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:153
			// _ = "end of CoverTab[107845]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:153
		// _ = "end of CoverTab[107838]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:153
		_go_fuzz_dep_.CoverTab[107839]++

												dfi := discardFieldInfo{field: toField(&f)}
												tf := f.Type

		// Unwrap tf to get its most basic type.
		var isPointer, isSlice bool
		if tf.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:160
			_go_fuzz_dep_.CoverTab[107846]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:160
			return tf.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:160
			// _ = "end of CoverTab[107846]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:160
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:160
			_go_fuzz_dep_.CoverTab[107847]++
													isSlice = true
													tf = tf.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:162
			// _ = "end of CoverTab[107847]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:163
			_go_fuzz_dep_.CoverTab[107848]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:163
			// _ = "end of CoverTab[107848]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:163
		// _ = "end of CoverTab[107839]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:163
		_go_fuzz_dep_.CoverTab[107840]++
												if tf.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:164
			_go_fuzz_dep_.CoverTab[107849]++
													isPointer = true
													tf = tf.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:166
			// _ = "end of CoverTab[107849]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:167
			_go_fuzz_dep_.CoverTab[107850]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:167
			// _ = "end of CoverTab[107850]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:167
		// _ = "end of CoverTab[107840]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:167
		_go_fuzz_dep_.CoverTab[107841]++
												if isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			_go_fuzz_dep_.CoverTab[107851]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			// _ = "end of CoverTab[107851]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			_go_fuzz_dep_.CoverTab[107852]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			return tf.Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			// _ = "end of CoverTab[107852]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:168
			_go_fuzz_dep_.CoverTab[107853]++
													panic(fmt.Sprintf("%v.%s cannot be a slice of pointers to primitive types", t, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:169
			// _ = "end of CoverTab[107853]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:170
			_go_fuzz_dep_.CoverTab[107854]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:170
			// _ = "end of CoverTab[107854]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:170
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:170
		// _ = "end of CoverTab[107841]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:170
		_go_fuzz_dep_.CoverTab[107842]++

												switch tf.Kind() {
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:173
			_go_fuzz_dep_.CoverTab[107855]++
													switch {
			case !isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:175
				_go_fuzz_dep_.CoverTab[107859]++
														panic(fmt.Sprintf("%v.%s cannot be a direct struct value", t, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:176
				// _ = "end of CoverTab[107859]"
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:177
				_go_fuzz_dep_.CoverTab[107860]++
														discardInfo := getDiscardInfo(tf)
														dfi.discard = func(src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:179
					_go_fuzz_dep_.CoverTab[107862]++
															sps := src.getPointerSlice()
															for _, sp := range sps {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:181
						_go_fuzz_dep_.CoverTab[107863]++
																if !sp.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:182
							_go_fuzz_dep_.CoverTab[107864]++
																	discardInfo.discard(sp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:183
							// _ = "end of CoverTab[107864]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:184
							_go_fuzz_dep_.CoverTab[107865]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:184
							// _ = "end of CoverTab[107865]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:184
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:184
						// _ = "end of CoverTab[107863]"
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:185
					// _ = "end of CoverTab[107862]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:186
				// _ = "end of CoverTab[107860]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:187
				_go_fuzz_dep_.CoverTab[107861]++
														discardInfo := getDiscardInfo(tf)
														dfi.discard = func(src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:189
					_go_fuzz_dep_.CoverTab[107866]++
															sp := src.getPointer()
															if !sp.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:191
						_go_fuzz_dep_.CoverTab[107867]++
																discardInfo.discard(sp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:192
						// _ = "end of CoverTab[107867]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:193
						_go_fuzz_dep_.CoverTab[107868]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:193
						// _ = "end of CoverTab[107868]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:193
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:193
					// _ = "end of CoverTab[107866]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:194
				// _ = "end of CoverTab[107861]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:195
			// _ = "end of CoverTab[107855]"
		case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:196
			_go_fuzz_dep_.CoverTab[107856]++
													switch {
			case isPointer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:198
				_go_fuzz_dep_.CoverTab[107871]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:198
				return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:198
				// _ = "end of CoverTab[107871]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:198
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:198
				_go_fuzz_dep_.CoverTab[107869]++
														panic(fmt.Sprintf("%v.%s cannot be a pointer to a map or a slice of map values", t, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:199
				// _ = "end of CoverTab[107869]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:200
				_go_fuzz_dep_.CoverTab[107870]++
														if tf.Elem().Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:201
					_go_fuzz_dep_.CoverTab[107872]++
															dfi.discard = func(src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:202
						_go_fuzz_dep_.CoverTab[107873]++
																sm := src.asPointerTo(tf).Elem()
																if sm.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:204
							_go_fuzz_dep_.CoverTab[107875]++
																	return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:205
							// _ = "end of CoverTab[107875]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:206
							_go_fuzz_dep_.CoverTab[107876]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:206
							// _ = "end of CoverTab[107876]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:206
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:206
						// _ = "end of CoverTab[107873]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:206
						_go_fuzz_dep_.CoverTab[107874]++
																for _, key := range sm.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:207
							_go_fuzz_dep_.CoverTab[107877]++
																	val := sm.MapIndex(key)
																	DiscardUnknown(val.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:209
							// _ = "end of CoverTab[107877]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:210
						// _ = "end of CoverTab[107874]"
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:211
					// _ = "end of CoverTab[107872]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:212
					_go_fuzz_dep_.CoverTab[107878]++
															dfi.discard = func(pointer) { _go_fuzz_dep_.CoverTab[107879]++; // _ = "end of CoverTab[107879]" }
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:213
					// _ = "end of CoverTab[107878]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:214
				// _ = "end of CoverTab[107870]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:215
			// _ = "end of CoverTab[107856]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:216
			_go_fuzz_dep_.CoverTab[107857]++

													switch {
			case isPointer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:219
				_go_fuzz_dep_.CoverTab[107882]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:219
				return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:219
				// _ = "end of CoverTab[107882]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:219
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:219
				_go_fuzz_dep_.CoverTab[107880]++
														panic(fmt.Sprintf("%v.%s cannot be a pointer to a interface or a slice of interface values", t, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:220
				// _ = "end of CoverTab[107880]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:221
				_go_fuzz_dep_.CoverTab[107881]++

														dfi.discard = func(src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:223
					_go_fuzz_dep_.CoverTab[107883]++
															su := src.asPointerTo(tf).Elem()
															if !su.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:225
						_go_fuzz_dep_.CoverTab[107884]++
																sv := su.Elem().Elem().Field(0)
																if sv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:227
							_go_fuzz_dep_.CoverTab[107886]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:227
							return sv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:227
							// _ = "end of CoverTab[107886]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:227
						}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:227
							_go_fuzz_dep_.CoverTab[107887]++
																	return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:228
							// _ = "end of CoverTab[107887]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:229
							_go_fuzz_dep_.CoverTab[107888]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:229
							// _ = "end of CoverTab[107888]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:229
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:229
						// _ = "end of CoverTab[107884]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:229
						_go_fuzz_dep_.CoverTab[107885]++
																switch sv.Type().Kind() {
						case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:231
							_go_fuzz_dep_.CoverTab[107889]++
																	DiscardUnknown(sv.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:232
							// _ = "end of CoverTab[107889]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:232
						default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:232
							_go_fuzz_dep_.CoverTab[107890]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:232
							// _ = "end of CoverTab[107890]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:233
						// _ = "end of CoverTab[107885]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:234
						_go_fuzz_dep_.CoverTab[107891]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:234
						// _ = "end of CoverTab[107891]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:234
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:234
					// _ = "end of CoverTab[107883]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:235
				// _ = "end of CoverTab[107881]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:236
			// _ = "end of CoverTab[107857]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:237
			_go_fuzz_dep_.CoverTab[107858]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:238
			// _ = "end of CoverTab[107858]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:239
		// _ = "end of CoverTab[107842]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:239
		_go_fuzz_dep_.CoverTab[107843]++
												di.fields = append(di.fields, dfi)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:240
		// _ = "end of CoverTab[107843]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:241
	// _ = "end of CoverTab[107833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:241
	_go_fuzz_dep_.CoverTab[107834]++

											di.unrecognized = invalidField
											if f, ok := t.FieldByName("XXX_unrecognized"); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:244
		_go_fuzz_dep_.CoverTab[107892]++
												if f.Type != reflect.TypeOf([]byte{}) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:245
			_go_fuzz_dep_.CoverTab[107894]++
													panic("expected XXX_unrecognized to be of type []byte")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:246
			// _ = "end of CoverTab[107894]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:247
			_go_fuzz_dep_.CoverTab[107895]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:247
			// _ = "end of CoverTab[107895]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:247
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:247
		// _ = "end of CoverTab[107892]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:247
		_go_fuzz_dep_.CoverTab[107893]++
												di.unrecognized = toField(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:248
		// _ = "end of CoverTab[107893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:249
		_go_fuzz_dep_.CoverTab[107896]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:249
		// _ = "end of CoverTab[107896]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:249
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:249
	// _ = "end of CoverTab[107834]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:249
	_go_fuzz_dep_.CoverTab[107835]++

											atomic.StoreInt32(&di.initialized, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:251
	// _ = "end of CoverTab[107835]"
}

func discardLegacy(m Message) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:254
	_go_fuzz_dep_.CoverTab[107897]++
											v := reflect.ValueOf(m)
											if v.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:256
		_go_fuzz_dep_.CoverTab[107902]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:256
		return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:256
		// _ = "end of CoverTab[107902]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:256
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:256
		_go_fuzz_dep_.CoverTab[107903]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:257
		// _ = "end of CoverTab[107903]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:258
		_go_fuzz_dep_.CoverTab[107904]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:258
		// _ = "end of CoverTab[107904]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:258
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:258
	// _ = "end of CoverTab[107897]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:258
	_go_fuzz_dep_.CoverTab[107898]++
											v = v.Elem()
											if v.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:260
		_go_fuzz_dep_.CoverTab[107905]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:261
		// _ = "end of CoverTab[107905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:262
		_go_fuzz_dep_.CoverTab[107906]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:262
		// _ = "end of CoverTab[107906]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:262
	// _ = "end of CoverTab[107898]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:262
	_go_fuzz_dep_.CoverTab[107899]++
											t := v.Type()

											for i := 0; i < v.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:265
		_go_fuzz_dep_.CoverTab[107907]++
												f := t.Field(i)
												if strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:267
			_go_fuzz_dep_.CoverTab[107912]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:268
			// _ = "end of CoverTab[107912]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:269
			_go_fuzz_dep_.CoverTab[107913]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:269
			// _ = "end of CoverTab[107913]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:269
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:269
		// _ = "end of CoverTab[107907]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:269
		_go_fuzz_dep_.CoverTab[107908]++
												vf := v.Field(i)
												tf := f.Type

		// Unwrap tf to get its most basic type.
		var isPointer, isSlice bool
		if tf.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:275
			_go_fuzz_dep_.CoverTab[107914]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:275
			return tf.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:275
			// _ = "end of CoverTab[107914]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:275
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:275
			_go_fuzz_dep_.CoverTab[107915]++
													isSlice = true
													tf = tf.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:277
			// _ = "end of CoverTab[107915]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:278
			_go_fuzz_dep_.CoverTab[107916]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:278
			// _ = "end of CoverTab[107916]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:278
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:278
		// _ = "end of CoverTab[107908]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:278
		_go_fuzz_dep_.CoverTab[107909]++
												if tf.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:279
			_go_fuzz_dep_.CoverTab[107917]++
													isPointer = true
													tf = tf.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:281
			// _ = "end of CoverTab[107917]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:282
			_go_fuzz_dep_.CoverTab[107918]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:282
			// _ = "end of CoverTab[107918]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:282
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:282
		// _ = "end of CoverTab[107909]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:282
		_go_fuzz_dep_.CoverTab[107910]++
												if isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			_go_fuzz_dep_.CoverTab[107919]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			// _ = "end of CoverTab[107919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			_go_fuzz_dep_.CoverTab[107920]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			return tf.Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			// _ = "end of CoverTab[107920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:283
			_go_fuzz_dep_.CoverTab[107921]++
													panic(fmt.Sprintf("%T.%s cannot be a slice of pointers to primitive types", m, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:284
			// _ = "end of CoverTab[107921]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:285
			_go_fuzz_dep_.CoverTab[107922]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:285
			// _ = "end of CoverTab[107922]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:285
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:285
		// _ = "end of CoverTab[107910]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:285
		_go_fuzz_dep_.CoverTab[107911]++

												switch tf.Kind() {
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:288
			_go_fuzz_dep_.CoverTab[107923]++
													switch {
			case !isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:290
				_go_fuzz_dep_.CoverTab[107927]++
														panic(fmt.Sprintf("%T.%s cannot be a direct struct value", m, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:291
				// _ = "end of CoverTab[107927]"
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:292
				_go_fuzz_dep_.CoverTab[107928]++
														for j := 0; j < vf.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:293
					_go_fuzz_dep_.CoverTab[107930]++
															discardLegacy(vf.Index(j).Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:294
					// _ = "end of CoverTab[107930]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:295
				// _ = "end of CoverTab[107928]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:296
				_go_fuzz_dep_.CoverTab[107929]++
														discardLegacy(vf.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:297
				// _ = "end of CoverTab[107929]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:298
			// _ = "end of CoverTab[107923]"
		case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:299
			_go_fuzz_dep_.CoverTab[107924]++
													switch {
			case isPointer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:301
				_go_fuzz_dep_.CoverTab[107933]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:301
				return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:301
				// _ = "end of CoverTab[107933]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:301
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:301
				_go_fuzz_dep_.CoverTab[107931]++
														panic(fmt.Sprintf("%T.%s cannot be a pointer to a map or a slice of map values", m, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:302
				// _ = "end of CoverTab[107931]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:303
				_go_fuzz_dep_.CoverTab[107932]++
														tv := vf.Type().Elem()
														if tv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:305
					_go_fuzz_dep_.CoverTab[107934]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:305
					return tv.Implements(protoMessageType)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:305
					// _ = "end of CoverTab[107934]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:305
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:305
					_go_fuzz_dep_.CoverTab[107935]++
															for _, key := range vf.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:306
						_go_fuzz_dep_.CoverTab[107936]++
																val := vf.MapIndex(key)
																discardLegacy(val.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:308
						// _ = "end of CoverTab[107936]"
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:309
					// _ = "end of CoverTab[107935]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:310
					_go_fuzz_dep_.CoverTab[107937]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:310
					// _ = "end of CoverTab[107937]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:310
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:310
				// _ = "end of CoverTab[107932]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:311
			// _ = "end of CoverTab[107924]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:312
			_go_fuzz_dep_.CoverTab[107925]++

													switch {
			case isPointer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:315
				_go_fuzz_dep_.CoverTab[107940]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:315
				return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:315
				// _ = "end of CoverTab[107940]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:315
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:315
				_go_fuzz_dep_.CoverTab[107938]++
														panic(fmt.Sprintf("%T.%s cannot be a pointer to a interface or a slice of interface values", m, f.Name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:316
				// _ = "end of CoverTab[107938]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:317
				_go_fuzz_dep_.CoverTab[107939]++
														if !vf.IsNil() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:318
					_go_fuzz_dep_.CoverTab[107941]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:318
					return f.Tag.Get("protobuf_oneof") != ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:318
					// _ = "end of CoverTab[107941]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:318
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:318
					_go_fuzz_dep_.CoverTab[107942]++
															vf = vf.Elem()
															if !vf.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:320
						_go_fuzz_dep_.CoverTab[107943]++
																vf = vf.Elem()
																vf = vf.Field(0)
																if vf.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:323
							_go_fuzz_dep_.CoverTab[107944]++
																	discardLegacy(vf.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:324
							// _ = "end of CoverTab[107944]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:325
							_go_fuzz_dep_.CoverTab[107945]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:325
							// _ = "end of CoverTab[107945]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:325
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:325
						// _ = "end of CoverTab[107943]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:326
						_go_fuzz_dep_.CoverTab[107946]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:326
						// _ = "end of CoverTab[107946]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:326
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:326
					// _ = "end of CoverTab[107942]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:327
					_go_fuzz_dep_.CoverTab[107947]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:327
					// _ = "end of CoverTab[107947]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:327
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:327
				// _ = "end of CoverTab[107939]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:328
			// _ = "end of CoverTab[107925]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:328
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:328
			_go_fuzz_dep_.CoverTab[107926]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:328
			// _ = "end of CoverTab[107926]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:329
		// _ = "end of CoverTab[107911]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:330
	// _ = "end of CoverTab[107899]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:330
	_go_fuzz_dep_.CoverTab[107900]++

											if vf := v.FieldByName("XXX_unrecognized"); vf.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:332
		_go_fuzz_dep_.CoverTab[107948]++
												if vf.Type() != reflect.TypeOf([]byte{}) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:333
			_go_fuzz_dep_.CoverTab[107950]++
													panic("expected XXX_unrecognized to be of type []byte")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:334
			// _ = "end of CoverTab[107950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:335
			_go_fuzz_dep_.CoverTab[107951]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:335
			// _ = "end of CoverTab[107951]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:335
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:335
		// _ = "end of CoverTab[107948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:335
		_go_fuzz_dep_.CoverTab[107949]++
												vf.Set(reflect.ValueOf([]byte(nil)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:336
		// _ = "end of CoverTab[107949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:337
		_go_fuzz_dep_.CoverTab[107952]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:337
		// _ = "end of CoverTab[107952]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:337
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:337
	// _ = "end of CoverTab[107900]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:337
	_go_fuzz_dep_.CoverTab[107901]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:341
	if em, err := extendable(m); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:341
		_go_fuzz_dep_.CoverTab[107953]++

												emm, _ := em.extensionsRead()
												for _, mx := range emm {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:344
			_go_fuzz_dep_.CoverTab[107954]++
													if m, ok := mx.value.(Message); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:345
				_go_fuzz_dep_.CoverTab[107955]++
														discardLegacy(m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:346
				// _ = "end of CoverTab[107955]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:347
				_go_fuzz_dep_.CoverTab[107956]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:347
				// _ = "end of CoverTab[107956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:347
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:347
			// _ = "end of CoverTab[107954]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:348
		// _ = "end of CoverTab[107953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:349
		_go_fuzz_dep_.CoverTab[107957]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:349
		// _ = "end of CoverTab[107957]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:349
	// _ = "end of CoverTab[107901]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:350
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/discard.go:350
var _ = _go_fuzz_dep_.CoverTab
