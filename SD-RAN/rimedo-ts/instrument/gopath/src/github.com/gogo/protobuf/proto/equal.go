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

// Protocol buffer comparison.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:34
)

import (
	"bytes"
	"log"
	"reflect"
	"strings"
)

/*
Equal returns true iff protocol buffers a and b are equal.
The arguments must both be pointers to protocol buffer structs.

Equality is defined in this way:
  - Two messages are equal iff they are the same type,
    corresponding fields are equal, unknown field sets
    are equal, and extensions sets are equal.
  - Two set scalar fields are equal iff their values are equal.
    If the fields are of a floating-point type, remember that
    NaN != x for all x, including NaN. If the message is defined
    in a proto3 .proto file, fields are not "set"; specifically,
    zero length proto3 "bytes" fields are equal (nil == {}).
  - Two repeated fields are equal iff their lengths are the same,
    and their corresponding elements are equal. Note a "bytes" field,
    although represented by []byte, is not a repeated field and the
    rule for the scalar fields described above applies.
  - Two unset fields are equal.
  - Two unknown field sets are equal if their current
    encoded state is equal.
  - Two extension sets are equal iff they have corresponding
    elements that are pairwise equal.
  - Two map fields are equal iff their lengths are the same,
    and they contain the same set of elements. Zero-length map
    fields are equal.
  - Every other combination of things are not equal.

The return value is undefined if a and b are not protocol buffers.
*/
func Equal(a, b Message) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:72
	_go_fuzz_dep_.CoverTab[108022]++
											if a == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:73
		_go_fuzz_dep_.CoverTab[108027]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:73
		return b == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:73
		// _ = "end of CoverTab[108027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:73
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:73
		_go_fuzz_dep_.CoverTab[108028]++
												return a == b
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:74
		// _ = "end of CoverTab[108028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:75
		_go_fuzz_dep_.CoverTab[108029]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:75
		// _ = "end of CoverTab[108029]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:75
	// _ = "end of CoverTab[108022]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:75
	_go_fuzz_dep_.CoverTab[108023]++
											v1, v2 := reflect.ValueOf(a), reflect.ValueOf(b)
											if v1.Type() != v2.Type() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:77
		_go_fuzz_dep_.CoverTab[108030]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:78
		// _ = "end of CoverTab[108030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:79
		_go_fuzz_dep_.CoverTab[108031]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:79
		// _ = "end of CoverTab[108031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:79
	// _ = "end of CoverTab[108023]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:79
	_go_fuzz_dep_.CoverTab[108024]++
											if v1.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:80
		_go_fuzz_dep_.CoverTab[108032]++
												if v1.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:81
			_go_fuzz_dep_.CoverTab[108035]++
													return v2.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:82
			// _ = "end of CoverTab[108035]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:83
			_go_fuzz_dep_.CoverTab[108036]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:83
			// _ = "end of CoverTab[108036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:83
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:83
		// _ = "end of CoverTab[108032]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:83
		_go_fuzz_dep_.CoverTab[108033]++
												if v2.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:84
			_go_fuzz_dep_.CoverTab[108037]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:85
			// _ = "end of CoverTab[108037]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:86
			_go_fuzz_dep_.CoverTab[108038]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:86
			// _ = "end of CoverTab[108038]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:86
		// _ = "end of CoverTab[108033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:86
		_go_fuzz_dep_.CoverTab[108034]++
												v1, v2 = v1.Elem(), v2.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:87
		// _ = "end of CoverTab[108034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:88
		_go_fuzz_dep_.CoverTab[108039]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:88
		// _ = "end of CoverTab[108039]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:88
	// _ = "end of CoverTab[108024]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:88
	_go_fuzz_dep_.CoverTab[108025]++
											if v1.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:89
		_go_fuzz_dep_.CoverTab[108040]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:90
		// _ = "end of CoverTab[108040]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:91
		_go_fuzz_dep_.CoverTab[108041]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:91
		// _ = "end of CoverTab[108041]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:91
	// _ = "end of CoverTab[108025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:91
	_go_fuzz_dep_.CoverTab[108026]++
											return equalStruct(v1, v2)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:92
	// _ = "end of CoverTab[108026]"
}

// v1 and v2 are known to have the same type.
func equalStruct(v1, v2 reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:96
	_go_fuzz_dep_.CoverTab[108042]++
											sprop := GetProperties(v1.Type())
											for i := 0; i < v1.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:98
		_go_fuzz_dep_.CoverTab[108047]++
												f := v1.Type().Field(i)
												if strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:100
			_go_fuzz_dep_.CoverTab[108050]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:101
			// _ = "end of CoverTab[108050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:102
			_go_fuzz_dep_.CoverTab[108051]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:102
			// _ = "end of CoverTab[108051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:102
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:102
		// _ = "end of CoverTab[108047]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:102
		_go_fuzz_dep_.CoverTab[108048]++
												f1, f2 := v1.Field(i), v2.Field(i)
												if f.Type.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:104
			_go_fuzz_dep_.CoverTab[108052]++
													if n1, n2 := f1.IsNil(), f2.IsNil(); n1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:105
				_go_fuzz_dep_.CoverTab[108054]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:105
				return n2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:105
				// _ = "end of CoverTab[108054]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:105
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:105
				_go_fuzz_dep_.CoverTab[108055]++

														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:107
				// _ = "end of CoverTab[108055]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:108
				_go_fuzz_dep_.CoverTab[108056]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:108
				if n1 != n2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:108
					_go_fuzz_dep_.CoverTab[108057]++

															return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:110
					// _ = "end of CoverTab[108057]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
					_go_fuzz_dep_.CoverTab[108058]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
					// _ = "end of CoverTab[108058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
				// _ = "end of CoverTab[108056]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
			// _ = "end of CoverTab[108052]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:111
			_go_fuzz_dep_.CoverTab[108053]++
													f1, f2 = f1.Elem(), f2.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:112
			// _ = "end of CoverTab[108053]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:113
			_go_fuzz_dep_.CoverTab[108059]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:113
			// _ = "end of CoverTab[108059]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:113
		// _ = "end of CoverTab[108048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:113
		_go_fuzz_dep_.CoverTab[108049]++
												if !equalAny(f1, f2, sprop.Prop[i]) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:114
			_go_fuzz_dep_.CoverTab[108060]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:115
			// _ = "end of CoverTab[108060]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:116
			_go_fuzz_dep_.CoverTab[108061]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:116
			// _ = "end of CoverTab[108061]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:116
		// _ = "end of CoverTab[108049]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:117
	// _ = "end of CoverTab[108042]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:117
	_go_fuzz_dep_.CoverTab[108043]++

											if em1 := v1.FieldByName("XXX_InternalExtensions"); em1.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:119
		_go_fuzz_dep_.CoverTab[108062]++
												em2 := v2.FieldByName("XXX_InternalExtensions")
												if !equalExtensions(v1.Type(), em1.Interface().(XXX_InternalExtensions), em2.Interface().(XXX_InternalExtensions)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:121
			_go_fuzz_dep_.CoverTab[108063]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:122
			// _ = "end of CoverTab[108063]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:123
			_go_fuzz_dep_.CoverTab[108064]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:123
			// _ = "end of CoverTab[108064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:123
		// _ = "end of CoverTab[108062]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:124
		_go_fuzz_dep_.CoverTab[108065]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:124
		// _ = "end of CoverTab[108065]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:124
	// _ = "end of CoverTab[108043]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:124
	_go_fuzz_dep_.CoverTab[108044]++

											if em1 := v1.FieldByName("XXX_extensions"); em1.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:126
		_go_fuzz_dep_.CoverTab[108066]++
												em2 := v2.FieldByName("XXX_extensions")
												if !equalExtMap(v1.Type(), em1.Interface().(map[int32]Extension), em2.Interface().(map[int32]Extension)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:128
			_go_fuzz_dep_.CoverTab[108067]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:129
			// _ = "end of CoverTab[108067]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:130
			_go_fuzz_dep_.CoverTab[108068]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:130
			// _ = "end of CoverTab[108068]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:130
		// _ = "end of CoverTab[108066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:131
		_go_fuzz_dep_.CoverTab[108069]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:131
		// _ = "end of CoverTab[108069]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:131
	// _ = "end of CoverTab[108044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:131
	_go_fuzz_dep_.CoverTab[108045]++

											uf := v1.FieldByName("XXX_unrecognized")
											if !uf.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:134
		_go_fuzz_dep_.CoverTab[108070]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:135
		// _ = "end of CoverTab[108070]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:136
		_go_fuzz_dep_.CoverTab[108071]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:136
		// _ = "end of CoverTab[108071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:136
	// _ = "end of CoverTab[108045]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:136
	_go_fuzz_dep_.CoverTab[108046]++

											u1 := uf.Bytes()
											u2 := v2.FieldByName("XXX_unrecognized").Bytes()
											return bytes.Equal(u1, u2)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:140
	// _ = "end of CoverTab[108046]"
}

// v1 and v2 are known to have the same type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:143
// prop may be nil.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:145
func equalAny(v1, v2 reflect.Value, prop *Properties) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:145
	_go_fuzz_dep_.CoverTab[108072]++
											if v1.Type() == protoMessageType {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:146
		_go_fuzz_dep_.CoverTab[108075]++
												m1, _ := v1.Interface().(Message)
												m2, _ := v2.Interface().(Message)
												return Equal(m1, m2)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:149
		// _ = "end of CoverTab[108075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:150
		_go_fuzz_dep_.CoverTab[108076]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:150
		// _ = "end of CoverTab[108076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:150
	// _ = "end of CoverTab[108072]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:150
	_go_fuzz_dep_.CoverTab[108073]++
											switch v1.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:152
		_go_fuzz_dep_.CoverTab[108077]++
												return v1.Bool() == v2.Bool()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:153
		// _ = "end of CoverTab[108077]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:154
		_go_fuzz_dep_.CoverTab[108078]++
												return v1.Float() == v2.Float()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:155
		// _ = "end of CoverTab[108078]"
	case reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:156
		_go_fuzz_dep_.CoverTab[108079]++
												return v1.Int() == v2.Int()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:157
		// _ = "end of CoverTab[108079]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:158
		_go_fuzz_dep_.CoverTab[108080]++

												n1, n2 := v1.IsNil(), v2.IsNil()
												if n1 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:161
			_go_fuzz_dep_.CoverTab[108097]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:161
			return n2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:161
			// _ = "end of CoverTab[108097]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:161
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:161
			_go_fuzz_dep_.CoverTab[108098]++
													return n1 == n2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:162
			// _ = "end of CoverTab[108098]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:163
			_go_fuzz_dep_.CoverTab[108099]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:163
			// _ = "end of CoverTab[108099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:163
		// _ = "end of CoverTab[108080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:163
		_go_fuzz_dep_.CoverTab[108081]++
												e1, e2 := v1.Elem(), v2.Elem()
												if e1.Type() != e2.Type() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:165
			_go_fuzz_dep_.CoverTab[108100]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:166
			// _ = "end of CoverTab[108100]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:167
			_go_fuzz_dep_.CoverTab[108101]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:167
			// _ = "end of CoverTab[108101]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:167
		// _ = "end of CoverTab[108081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:167
		_go_fuzz_dep_.CoverTab[108082]++
												return equalAny(e1, e2, nil)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:168
		// _ = "end of CoverTab[108082]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:169
		_go_fuzz_dep_.CoverTab[108083]++
												if v1.Len() != v2.Len() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:170
			_go_fuzz_dep_.CoverTab[108102]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:171
			// _ = "end of CoverTab[108102]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:172
			_go_fuzz_dep_.CoverTab[108103]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:172
			// _ = "end of CoverTab[108103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:172
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:172
		// _ = "end of CoverTab[108083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:172
		_go_fuzz_dep_.CoverTab[108084]++
												for _, key := range v1.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:173
			_go_fuzz_dep_.CoverTab[108104]++
													val2 := v2.MapIndex(key)
													if !val2.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:175
				_go_fuzz_dep_.CoverTab[108106]++

														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:177
				// _ = "end of CoverTab[108106]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:178
				_go_fuzz_dep_.CoverTab[108107]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:178
				// _ = "end of CoverTab[108107]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:178
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:178
			// _ = "end of CoverTab[108104]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:178
			_go_fuzz_dep_.CoverTab[108105]++
													if !equalAny(v1.MapIndex(key), val2, nil) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:179
				_go_fuzz_dep_.CoverTab[108108]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:180
				// _ = "end of CoverTab[108108]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:181
				_go_fuzz_dep_.CoverTab[108109]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:181
				// _ = "end of CoverTab[108109]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:181
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:181
			// _ = "end of CoverTab[108105]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:182
		// _ = "end of CoverTab[108084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:182
		_go_fuzz_dep_.CoverTab[108085]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:183
		// _ = "end of CoverTab[108085]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:184
		_go_fuzz_dep_.CoverTab[108086]++

												if v1.IsNil() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:186
			_go_fuzz_dep_.CoverTab[108110]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:186
			return v2.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:186
			// _ = "end of CoverTab[108110]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:186
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:186
			_go_fuzz_dep_.CoverTab[108111]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:187
			// _ = "end of CoverTab[108111]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:188
			_go_fuzz_dep_.CoverTab[108112]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:188
			// _ = "end of CoverTab[108112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:188
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:188
		// _ = "end of CoverTab[108086]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:188
		_go_fuzz_dep_.CoverTab[108087]++
												if v1.IsNil() != v2.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:189
			_go_fuzz_dep_.CoverTab[108113]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:190
			// _ = "end of CoverTab[108113]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:191
			_go_fuzz_dep_.CoverTab[108114]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:191
			// _ = "end of CoverTab[108114]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:191
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:191
		// _ = "end of CoverTab[108087]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:191
		_go_fuzz_dep_.CoverTab[108088]++
												return equalAny(v1.Elem(), v2.Elem(), prop)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:192
		// _ = "end of CoverTab[108088]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:193
		_go_fuzz_dep_.CoverTab[108089]++
												if v1.Type().Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:194
			_go_fuzz_dep_.CoverTab[108115]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
			if prop != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				_go_fuzz_dep_.CoverTab[108118]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				return prop.proto3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				// _ = "end of CoverTab[108118]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				_go_fuzz_dep_.CoverTab[108119]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				return v1.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				// _ = "end of CoverTab[108119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				_go_fuzz_dep_.CoverTab[108120]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				return v2.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				// _ = "end of CoverTab[108120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:199
				_go_fuzz_dep_.CoverTab[108121]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:200
				// _ = "end of CoverTab[108121]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:201
				_go_fuzz_dep_.CoverTab[108122]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:201
				// _ = "end of CoverTab[108122]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:201
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:201
			// _ = "end of CoverTab[108115]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:201
			_go_fuzz_dep_.CoverTab[108116]++
													if v1.IsNil() != v2.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:202
				_go_fuzz_dep_.CoverTab[108123]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:203
				// _ = "end of CoverTab[108123]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:204
				_go_fuzz_dep_.CoverTab[108124]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:204
				// _ = "end of CoverTab[108124]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:204
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:204
			// _ = "end of CoverTab[108116]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:204
			_go_fuzz_dep_.CoverTab[108117]++
													return bytes.Equal(v1.Interface().([]byte), v2.Interface().([]byte))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:205
			// _ = "end of CoverTab[108117]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:206
			_go_fuzz_dep_.CoverTab[108125]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:206
			// _ = "end of CoverTab[108125]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:206
		// _ = "end of CoverTab[108089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:206
		_go_fuzz_dep_.CoverTab[108090]++

												if v1.Len() != v2.Len() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:208
			_go_fuzz_dep_.CoverTab[108126]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:209
			// _ = "end of CoverTab[108126]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:210
			_go_fuzz_dep_.CoverTab[108127]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:210
			// _ = "end of CoverTab[108127]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:210
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:210
		// _ = "end of CoverTab[108090]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:210
		_go_fuzz_dep_.CoverTab[108091]++
												for i := 0; i < v1.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:211
			_go_fuzz_dep_.CoverTab[108128]++
													if !equalAny(v1.Index(i), v2.Index(i), prop) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:212
				_go_fuzz_dep_.CoverTab[108129]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:213
				// _ = "end of CoverTab[108129]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:214
				_go_fuzz_dep_.CoverTab[108130]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:214
				// _ = "end of CoverTab[108130]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:214
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:214
			// _ = "end of CoverTab[108128]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:215
		// _ = "end of CoverTab[108091]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:215
		_go_fuzz_dep_.CoverTab[108092]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:216
		// _ = "end of CoverTab[108092]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:217
		_go_fuzz_dep_.CoverTab[108093]++
												return v1.Interface().(string) == v2.Interface().(string)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:218
		// _ = "end of CoverTab[108093]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:219
		_go_fuzz_dep_.CoverTab[108094]++
												return equalStruct(v1, v2)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:220
		// _ = "end of CoverTab[108094]"
	case reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:221
		_go_fuzz_dep_.CoverTab[108095]++
												return v1.Uint() == v2.Uint()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:222
		// _ = "end of CoverTab[108095]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:222
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:222
		_go_fuzz_dep_.CoverTab[108096]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:222
		// _ = "end of CoverTab[108096]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:223
	// _ = "end of CoverTab[108073]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:223
	_go_fuzz_dep_.CoverTab[108074]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:226
	log.Printf("proto: don't know how to compare %v", v1)
											return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:227
	// _ = "end of CoverTab[108074]"
}

// base is the struct type that the extensions are based on.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:230
// x1 and x2 are InternalExtensions.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:232
func equalExtensions(base reflect.Type, x1, x2 XXX_InternalExtensions) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:232
	_go_fuzz_dep_.CoverTab[108131]++
											em1, _ := x1.extensionsRead()
											em2, _ := x2.extensionsRead()
											return equalExtMap(base, em1, em2)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:235
	// _ = "end of CoverTab[108131]"
}

func equalExtMap(base reflect.Type, em1, em2 map[int32]Extension) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:238
	_go_fuzz_dep_.CoverTab[108132]++
											if len(em1) != len(em2) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:239
		_go_fuzz_dep_.CoverTab[108135]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:240
		// _ = "end of CoverTab[108135]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:241
		_go_fuzz_dep_.CoverTab[108136]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:241
		// _ = "end of CoverTab[108136]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:241
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:241
	// _ = "end of CoverTab[108132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:241
	_go_fuzz_dep_.CoverTab[108133]++

											for extNum, e1 := range em1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:243
		_go_fuzz_dep_.CoverTab[108137]++
												e2, ok := em2[extNum]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:245
			_go_fuzz_dep_.CoverTab[108146]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:246
			// _ = "end of CoverTab[108146]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:247
			_go_fuzz_dep_.CoverTab[108147]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:247
			// _ = "end of CoverTab[108147]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:247
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:247
		// _ = "end of CoverTab[108137]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:247
		_go_fuzz_dep_.CoverTab[108138]++

												m1, m2 := e1.value, e2.value

												if m1 == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:251
			_go_fuzz_dep_.CoverTab[108148]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:251
			return m2 == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:251
			// _ = "end of CoverTab[108148]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:251
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:251
			_go_fuzz_dep_.CoverTab[108149]++

													if bytes.Equal(e1.enc, e2.enc) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:253
				_go_fuzz_dep_.CoverTab[108150]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:254
				// _ = "end of CoverTab[108150]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:255
				_go_fuzz_dep_.CoverTab[108151]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:255
				// _ = "end of CoverTab[108151]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:255
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:255
			// _ = "end of CoverTab[108149]"

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:258
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:258
			_go_fuzz_dep_.CoverTab[108152]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:258
			// _ = "end of CoverTab[108152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:258
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:258
		// _ = "end of CoverTab[108138]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:258
		_go_fuzz_dep_.CoverTab[108139]++

												if m1 != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:260
			_go_fuzz_dep_.CoverTab[108153]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:260
			return m2 != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:260
			// _ = "end of CoverTab[108153]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:260
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:260
			_go_fuzz_dep_.CoverTab[108154]++

													if !equalAny(reflect.ValueOf(m1), reflect.ValueOf(m2), nil) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:262
				_go_fuzz_dep_.CoverTab[108156]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:263
				// _ = "end of CoverTab[108156]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:264
				_go_fuzz_dep_.CoverTab[108157]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:264
				// _ = "end of CoverTab[108157]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:264
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:264
			// _ = "end of CoverTab[108154]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:264
			_go_fuzz_dep_.CoverTab[108155]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:265
			// _ = "end of CoverTab[108155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:266
			_go_fuzz_dep_.CoverTab[108158]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:266
			// _ = "end of CoverTab[108158]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:266
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:266
		// _ = "end of CoverTab[108139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:266
		_go_fuzz_dep_.CoverTab[108140]++

		// At least one is encoded. To do a semantically correct comparison
		// we need to unmarshal them first.
		var desc *ExtensionDesc
		if m := extensionMaps[base]; m != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:271
			_go_fuzz_dep_.CoverTab[108159]++
													desc = m[extNum]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:272
			// _ = "end of CoverTab[108159]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:273
			_go_fuzz_dep_.CoverTab[108160]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:273
			// _ = "end of CoverTab[108160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:273
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:273
		// _ = "end of CoverTab[108140]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:273
		_go_fuzz_dep_.CoverTab[108141]++
												if desc == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:274
			_go_fuzz_dep_.CoverTab[108161]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:279
			log.Printf("proto: don't know how to compare extension %d of %v", extNum, base)
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:280
			// _ = "end of CoverTab[108161]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:281
			_go_fuzz_dep_.CoverTab[108162]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:281
			// _ = "end of CoverTab[108162]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:281
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:281
		// _ = "end of CoverTab[108141]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:281
		_go_fuzz_dep_.CoverTab[108142]++
												var err error
												if m1 == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:283
			_go_fuzz_dep_.CoverTab[108163]++
													m1, err = decodeExtension(e1.enc, desc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:284
			// _ = "end of CoverTab[108163]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:285
			_go_fuzz_dep_.CoverTab[108164]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:285
			// _ = "end of CoverTab[108164]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:285
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:285
		// _ = "end of CoverTab[108142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:285
		_go_fuzz_dep_.CoverTab[108143]++
												if m2 == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:286
			_go_fuzz_dep_.CoverTab[108165]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:286
			return err == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:286
			// _ = "end of CoverTab[108165]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:286
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:286
			_go_fuzz_dep_.CoverTab[108166]++
													m2, err = decodeExtension(e2.enc, desc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:287
			// _ = "end of CoverTab[108166]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:288
			_go_fuzz_dep_.CoverTab[108167]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:288
			// _ = "end of CoverTab[108167]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:288
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:288
		// _ = "end of CoverTab[108143]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:288
		_go_fuzz_dep_.CoverTab[108144]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:289
			_go_fuzz_dep_.CoverTab[108168]++

													log.Printf("proto: badly encoded extension %d of %v: %v", extNum, base, err)
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:292
			// _ = "end of CoverTab[108168]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:293
			_go_fuzz_dep_.CoverTab[108169]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:293
			// _ = "end of CoverTab[108169]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:293
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:293
		// _ = "end of CoverTab[108144]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:293
		_go_fuzz_dep_.CoverTab[108145]++
												if !equalAny(reflect.ValueOf(m1), reflect.ValueOf(m2), nil) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:294
			_go_fuzz_dep_.CoverTab[108170]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:295
			// _ = "end of CoverTab[108170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:296
			_go_fuzz_dep_.CoverTab[108171]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:296
			// _ = "end of CoverTab[108171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:296
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:296
		// _ = "end of CoverTab[108145]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:297
	// _ = "end of CoverTab[108133]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:297
	_go_fuzz_dep_.CoverTab[108134]++

											return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:299
	// _ = "end of CoverTab[108134]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:300
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/equal.go:300
var _ = _go_fuzz_dep_.CoverTab
