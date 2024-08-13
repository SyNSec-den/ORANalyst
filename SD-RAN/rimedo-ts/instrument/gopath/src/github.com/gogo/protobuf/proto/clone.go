// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2011 The Go Authors.  All rights reserved.
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

// Protocol buffer deep copy and merge.
// TODO: RawMessage.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:35
)

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Clone returns a deep copy of a protocol buffer.
func Clone(src Message) Message {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:45
	_go_fuzz_dep_.CoverTab[107574]++
											in := reflect.ValueOf(src)
											if in.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:47
		_go_fuzz_dep_.CoverTab[107576]++
												return src
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:48
		// _ = "end of CoverTab[107576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:49
		_go_fuzz_dep_.CoverTab[107577]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:49
		// _ = "end of CoverTab[107577]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:49
	// _ = "end of CoverTab[107574]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:49
	_go_fuzz_dep_.CoverTab[107575]++
											out := reflect.New(in.Type().Elem())
											dst := out.Interface().(Message)
											Merge(dst, src)
											return dst
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:53
	// _ = "end of CoverTab[107575]"
}

// Merger is the interface representing objects that can merge messages of the same type.
type Merger interface {
	// Merge merges src into this message.
	// Required and optional fields that are set in src will be set to that value in dst.
	// Elements of repeated fields will be appended.
	//
	// Merge may panic if called with a different argument type than the receiver.
	Merge(src Message)
}

// generatedMerger is the custom merge method that generated protos will have.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:66
// We must add this method since a generate Merge method will conflict with
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:66
// many existing protos that have a Merge data field already defined.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:69
type generatedMerger interface {
	XXX_Merge(src Message)
}

// Merge merges src into dst.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:73
// Required and optional fields that are set in src will be set to that value in dst.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:73
// Elements of repeated fields will be appended.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:73
// Merge panics if src and dst are not the same type, or if dst is nil.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:77
func Merge(dst, src Message) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:77
	_go_fuzz_dep_.CoverTab[107578]++
											if m, ok := dst.(Merger); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:78
		_go_fuzz_dep_.CoverTab[107584]++
												m.Merge(src)
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:80
		// _ = "end of CoverTab[107584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:81
		_go_fuzz_dep_.CoverTab[107585]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:81
		// _ = "end of CoverTab[107585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:81
	// _ = "end of CoverTab[107578]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:81
	_go_fuzz_dep_.CoverTab[107579]++

											in := reflect.ValueOf(src)
											out := reflect.ValueOf(dst)
											if out.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:85
		_go_fuzz_dep_.CoverTab[107586]++
												panic("proto: nil destination")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:86
		// _ = "end of CoverTab[107586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:87
		_go_fuzz_dep_.CoverTab[107587]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:87
		// _ = "end of CoverTab[107587]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:87
	// _ = "end of CoverTab[107579]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:87
	_go_fuzz_dep_.CoverTab[107580]++
											if in.Type() != out.Type() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:88
		_go_fuzz_dep_.CoverTab[107588]++
												panic(fmt.Sprintf("proto.Merge(%T, %T) type mismatch", dst, src))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:89
		// _ = "end of CoverTab[107588]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:90
		_go_fuzz_dep_.CoverTab[107589]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:90
		// _ = "end of CoverTab[107589]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:90
	// _ = "end of CoverTab[107580]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:90
	_go_fuzz_dep_.CoverTab[107581]++
											if in.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:91
		_go_fuzz_dep_.CoverTab[107590]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:92
		// _ = "end of CoverTab[107590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:93
		_go_fuzz_dep_.CoverTab[107591]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:93
		// _ = "end of CoverTab[107591]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:93
	// _ = "end of CoverTab[107581]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:93
	_go_fuzz_dep_.CoverTab[107582]++
											if m, ok := dst.(generatedMerger); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:94
		_go_fuzz_dep_.CoverTab[107592]++
												m.XXX_Merge(src)
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:96
		// _ = "end of CoverTab[107592]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:97
		_go_fuzz_dep_.CoverTab[107593]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:97
		// _ = "end of CoverTab[107593]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:97
	// _ = "end of CoverTab[107582]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:97
	_go_fuzz_dep_.CoverTab[107583]++
											mergeStruct(out.Elem(), in.Elem())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:98
	// _ = "end of CoverTab[107583]"
}

func mergeStruct(out, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:101
	_go_fuzz_dep_.CoverTab[107594]++
											sprop := GetProperties(in.Type())
											for i := 0; i < in.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:103
		_go_fuzz_dep_.CoverTab[107598]++
												f := in.Type().Field(i)
												if strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:105
			_go_fuzz_dep_.CoverTab[107600]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:106
			// _ = "end of CoverTab[107600]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:107
			_go_fuzz_dep_.CoverTab[107601]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:107
			// _ = "end of CoverTab[107601]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:107
		// _ = "end of CoverTab[107598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:107
		_go_fuzz_dep_.CoverTab[107599]++
												mergeAny(out.Field(i), in.Field(i), false, sprop.Prop[i])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:108
		// _ = "end of CoverTab[107599]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:109
	// _ = "end of CoverTab[107594]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:109
	_go_fuzz_dep_.CoverTab[107595]++

											if emIn, ok := in.Addr().Interface().(extensionsBytes); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:111
		_go_fuzz_dep_.CoverTab[107602]++
												emOut := out.Addr().Interface().(extensionsBytes)
												bIn := emIn.GetExtensions()
												bOut := emOut.GetExtensions()
												*bOut = append(*bOut, *bIn...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:115
		// _ = "end of CoverTab[107602]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:116
		_go_fuzz_dep_.CoverTab[107603]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:116
		if emIn, err := extendable(in.Addr().Interface()); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:116
			_go_fuzz_dep_.CoverTab[107604]++
													emOut, _ := extendable(out.Addr().Interface())
													mIn, muIn := emIn.extensionsRead()
													if mIn != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:119
				_go_fuzz_dep_.CoverTab[107605]++
														mOut := emOut.extensionsWrite()
														muIn.Lock()
														mergeExtension(mOut, mIn)
														muIn.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:123
				// _ = "end of CoverTab[107605]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:124
				_go_fuzz_dep_.CoverTab[107606]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:124
				// _ = "end of CoverTab[107606]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:124
			// _ = "end of CoverTab[107604]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
			_go_fuzz_dep_.CoverTab[107607]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
			// _ = "end of CoverTab[107607]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
		// _ = "end of CoverTab[107603]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
	// _ = "end of CoverTab[107595]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:125
	_go_fuzz_dep_.CoverTab[107596]++

											uf := in.FieldByName("XXX_unrecognized")
											if !uf.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:128
		_go_fuzz_dep_.CoverTab[107608]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:129
		// _ = "end of CoverTab[107608]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:130
		_go_fuzz_dep_.CoverTab[107609]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:130
		// _ = "end of CoverTab[107609]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:130
	// _ = "end of CoverTab[107596]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:130
	_go_fuzz_dep_.CoverTab[107597]++
											uin := uf.Bytes()
											if len(uin) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:132
		_go_fuzz_dep_.CoverTab[107610]++
												out.FieldByName("XXX_unrecognized").SetBytes(append([]byte(nil), uin...))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:133
		// _ = "end of CoverTab[107610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:134
		_go_fuzz_dep_.CoverTab[107611]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:134
		// _ = "end of CoverTab[107611]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:134
	// _ = "end of CoverTab[107597]"
}

// mergeAny performs a merge between two values of the same type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:137
// viaPtr indicates whether the values were indirected through a pointer (implying proto2).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:137
// prop is set if this is a struct field (it may be nil).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:140
func mergeAny(out, in reflect.Value, viaPtr bool, prop *Properties) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:140
	_go_fuzz_dep_.CoverTab[107612]++
											if in.Type() == protoMessageType {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:141
		_go_fuzz_dep_.CoverTab[107614]++
												if !in.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:142
			_go_fuzz_dep_.CoverTab[107616]++
													if out.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:143
				_go_fuzz_dep_.CoverTab[107617]++
														out.Set(reflect.ValueOf(Clone(in.Interface().(Message))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:144
				// _ = "end of CoverTab[107617]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:145
				_go_fuzz_dep_.CoverTab[107618]++
														Merge(out.Interface().(Message), in.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:146
				// _ = "end of CoverTab[107618]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:147
			// _ = "end of CoverTab[107616]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:148
			_go_fuzz_dep_.CoverTab[107619]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:148
			// _ = "end of CoverTab[107619]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:148
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:148
		// _ = "end of CoverTab[107614]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:148
		_go_fuzz_dep_.CoverTab[107615]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:149
		// _ = "end of CoverTab[107615]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:150
		_go_fuzz_dep_.CoverTab[107620]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:150
		// _ = "end of CoverTab[107620]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:150
	// _ = "end of CoverTab[107612]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:150
	_go_fuzz_dep_.CoverTab[107613]++
											switch in.Kind() {
	case reflect.Bool, reflect.Float32, reflect.Float64, reflect.Int32, reflect.Int64,
		reflect.String, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:153
		_go_fuzz_dep_.CoverTab[107621]++
												if !viaPtr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:154
			_go_fuzz_dep_.CoverTab[107638]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:154
			return isProto3Zero(in)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:154
			// _ = "end of CoverTab[107638]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:154
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:154
			_go_fuzz_dep_.CoverTab[107639]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:155
			// _ = "end of CoverTab[107639]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:156
			_go_fuzz_dep_.CoverTab[107640]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:156
			// _ = "end of CoverTab[107640]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:156
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:156
		// _ = "end of CoverTab[107621]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:156
		_go_fuzz_dep_.CoverTab[107622]++
												out.Set(in)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:157
		// _ = "end of CoverTab[107622]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:158
		_go_fuzz_dep_.CoverTab[107623]++

												if in.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:160
			_go_fuzz_dep_.CoverTab[107641]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:161
			// _ = "end of CoverTab[107641]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:162
			_go_fuzz_dep_.CoverTab[107642]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:162
			// _ = "end of CoverTab[107642]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:162
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:162
		// _ = "end of CoverTab[107623]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:162
		_go_fuzz_dep_.CoverTab[107624]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:165
		if out.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:165
			_go_fuzz_dep_.CoverTab[107643]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:165
			return out.Elem().Type() != in.Elem().Type()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:165
			// _ = "end of CoverTab[107643]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:165
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:165
			_go_fuzz_dep_.CoverTab[107644]++
													out.Set(reflect.New(in.Elem().Elem().Type()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:166
			// _ = "end of CoverTab[107644]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:167
			_go_fuzz_dep_.CoverTab[107645]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:167
			// _ = "end of CoverTab[107645]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:167
		// _ = "end of CoverTab[107624]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:167
		_go_fuzz_dep_.CoverTab[107625]++
												mergeAny(out.Elem(), in.Elem(), false, nil)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:168
		// _ = "end of CoverTab[107625]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:169
		_go_fuzz_dep_.CoverTab[107626]++
												if in.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:170
			_go_fuzz_dep_.CoverTab[107646]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:171
			// _ = "end of CoverTab[107646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:172
			_go_fuzz_dep_.CoverTab[107647]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:172
			// _ = "end of CoverTab[107647]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:172
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:172
		// _ = "end of CoverTab[107626]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:172
		_go_fuzz_dep_.CoverTab[107627]++
												if out.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:173
			_go_fuzz_dep_.CoverTab[107648]++
													out.Set(reflect.MakeMap(in.Type()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:174
			// _ = "end of CoverTab[107648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:175
			_go_fuzz_dep_.CoverTab[107649]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:175
			// _ = "end of CoverTab[107649]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:175
		// _ = "end of CoverTab[107627]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:175
		_go_fuzz_dep_.CoverTab[107628]++

												elemKind := in.Type().Elem().Kind()
												for _, key := range in.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:178
			_go_fuzz_dep_.CoverTab[107650]++
													var val reflect.Value
													switch elemKind {
			case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:181
				_go_fuzz_dep_.CoverTab[107652]++
														val = reflect.New(in.Type().Elem().Elem())
														mergeAny(val, in.MapIndex(key), false, nil)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:183
				// _ = "end of CoverTab[107652]"
			case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:184
				_go_fuzz_dep_.CoverTab[107653]++
														val = in.MapIndex(key)
														val = reflect.ValueOf(append([]byte{}, val.Bytes()...))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:186
				// _ = "end of CoverTab[107653]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:187
				_go_fuzz_dep_.CoverTab[107654]++
														val = in.MapIndex(key)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:188
				// _ = "end of CoverTab[107654]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:189
			// _ = "end of CoverTab[107650]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:189
			_go_fuzz_dep_.CoverTab[107651]++
													out.SetMapIndex(key, val)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:190
			// _ = "end of CoverTab[107651]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:191
		// _ = "end of CoverTab[107628]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:192
		_go_fuzz_dep_.CoverTab[107629]++
												if in.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:193
			_go_fuzz_dep_.CoverTab[107655]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:194
			// _ = "end of CoverTab[107655]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:195
			_go_fuzz_dep_.CoverTab[107656]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:195
			// _ = "end of CoverTab[107656]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:195
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:195
		// _ = "end of CoverTab[107629]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:195
		_go_fuzz_dep_.CoverTab[107630]++
												if out.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:196
			_go_fuzz_dep_.CoverTab[107657]++
													out.Set(reflect.New(in.Elem().Type()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:197
			// _ = "end of CoverTab[107657]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:198
			_go_fuzz_dep_.CoverTab[107658]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:198
			// _ = "end of CoverTab[107658]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:198
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:198
		// _ = "end of CoverTab[107630]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:198
		_go_fuzz_dep_.CoverTab[107631]++
												mergeAny(out.Elem(), in.Elem(), true, nil)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:199
		// _ = "end of CoverTab[107631]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:200
		_go_fuzz_dep_.CoverTab[107632]++
												if in.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:201
			_go_fuzz_dep_.CoverTab[107659]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:202
			// _ = "end of CoverTab[107659]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:203
			_go_fuzz_dep_.CoverTab[107660]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:203
			// _ = "end of CoverTab[107660]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:203
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:203
		// _ = "end of CoverTab[107632]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:203
		_go_fuzz_dep_.CoverTab[107633]++
												if in.Type().Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:204
			_go_fuzz_dep_.CoverTab[107661]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
			if prop != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				_go_fuzz_dep_.CoverTab[107663]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				return prop.proto3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				// _ = "end of CoverTab[107663]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				_go_fuzz_dep_.CoverTab[107664]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				return in.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				// _ = "end of CoverTab[107664]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:210
				_go_fuzz_dep_.CoverTab[107665]++
														return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:211
				// _ = "end of CoverTab[107665]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:212
				_go_fuzz_dep_.CoverTab[107666]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:212
				// _ = "end of CoverTab[107666]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:212
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:212
			// _ = "end of CoverTab[107661]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:212
			_go_fuzz_dep_.CoverTab[107662]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:217
			out.SetBytes(append([]byte{}, in.Bytes()...))
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:218
			// _ = "end of CoverTab[107662]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:219
			_go_fuzz_dep_.CoverTab[107667]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:219
			// _ = "end of CoverTab[107667]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:219
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:219
		// _ = "end of CoverTab[107633]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:219
		_go_fuzz_dep_.CoverTab[107634]++
												n := in.Len()
												if out.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:221
			_go_fuzz_dep_.CoverTab[107668]++
													out.Set(reflect.MakeSlice(in.Type(), 0, n))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:222
			// _ = "end of CoverTab[107668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:223
			_go_fuzz_dep_.CoverTab[107669]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:223
			// _ = "end of CoverTab[107669]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:223
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:223
		// _ = "end of CoverTab[107634]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:223
		_go_fuzz_dep_.CoverTab[107635]++
												switch in.Type().Elem().Kind() {
		case reflect.Bool, reflect.Float32, reflect.Float64, reflect.Int32, reflect.Int64,
			reflect.String, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:226
			_go_fuzz_dep_.CoverTab[107670]++
													out.Set(reflect.AppendSlice(out, in))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:227
			// _ = "end of CoverTab[107670]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:228
			_go_fuzz_dep_.CoverTab[107671]++
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:229
				_go_fuzz_dep_.CoverTab[107672]++
														x := reflect.Indirect(reflect.New(in.Type().Elem()))
														mergeAny(x, in.Index(i), false, nil)
														out.Set(reflect.Append(out, x))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:232
				// _ = "end of CoverTab[107672]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:233
			// _ = "end of CoverTab[107671]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:234
		// _ = "end of CoverTab[107635]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:235
		_go_fuzz_dep_.CoverTab[107636]++
												mergeStruct(out, in)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:236
		// _ = "end of CoverTab[107636]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:237
		_go_fuzz_dep_.CoverTab[107637]++

												log.Printf("proto: don't know how to copy %v", in)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:239
		// _ = "end of CoverTab[107637]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:240
	// _ = "end of CoverTab[107613]"
}

func mergeExtension(out, in map[int32]Extension) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:243
	_go_fuzz_dep_.CoverTab[107673]++
											for extNum, eIn := range in {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:244
		_go_fuzz_dep_.CoverTab[107674]++
												eOut := Extension{desc: eIn.desc}
												if eIn.value != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:246
			_go_fuzz_dep_.CoverTab[107677]++
													v := reflect.New(reflect.TypeOf(eIn.value)).Elem()
													mergeAny(v, reflect.ValueOf(eIn.value), false, nil)
													eOut.value = v.Interface()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:249
			// _ = "end of CoverTab[107677]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:250
			_go_fuzz_dep_.CoverTab[107678]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:250
			// _ = "end of CoverTab[107678]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:250
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:250
		// _ = "end of CoverTab[107674]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:250
		_go_fuzz_dep_.CoverTab[107675]++
												if eIn.enc != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:251
			_go_fuzz_dep_.CoverTab[107679]++
													eOut.enc = make([]byte, len(eIn.enc))
													copy(eOut.enc, eIn.enc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:253
			// _ = "end of CoverTab[107679]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:254
			_go_fuzz_dep_.CoverTab[107680]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:254
			// _ = "end of CoverTab[107680]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:254
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:254
		// _ = "end of CoverTab[107675]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:254
		_go_fuzz_dep_.CoverTab[107676]++

												out[extNum] = eOut
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:256
		// _ = "end of CoverTab[107676]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:257
	// _ = "end of CoverTab[107673]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:258
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/clone.go:258
var _ = _go_fuzz_dep_.CoverTab
