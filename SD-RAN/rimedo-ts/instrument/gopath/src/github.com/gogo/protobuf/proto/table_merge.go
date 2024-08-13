// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2016 The Go Authors.  All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:32
)

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
)

// Merge merges the src message into dst.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:42
// This assumes that dst and src of the same type and are non-nil.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:44
func (a *InternalMessageInfo) Merge(dst, src Message) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:44
	_go_fuzz_dep_.CoverTab[110697]++
												mi := atomicLoadMergeInfo(&a.merge)
												if mi == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:46
		_go_fuzz_dep_.CoverTab[110699]++
													mi = getMergeInfo(reflect.TypeOf(dst).Elem())
													atomicStoreMergeInfo(&a.merge, mi)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:48
		// _ = "end of CoverTab[110699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:49
		_go_fuzz_dep_.CoverTab[110700]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:49
		// _ = "end of CoverTab[110700]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:49
	// _ = "end of CoverTab[110697]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:49
	_go_fuzz_dep_.CoverTab[110698]++
												mi.merge(toPointer(&dst), toPointer(&src))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:50
	// _ = "end of CoverTab[110698]"
}

type mergeInfo struct {
	typ	reflect.Type

	initialized	int32	// 0: only typ is valid, 1: everything is valid
	lock		sync.Mutex

	fields		[]mergeFieldInfo
	unrecognized	field	// Offset of XXX_unrecognized
}

type mergeFieldInfo struct {
	field	field	// Offset of field, guaranteed to be valid

	// isPointer reports whether the value in the field is a pointer.
	// This is true for the following situations:
	//	* Pointer to struct
	//	* Pointer to basic type (proto2 only)
	//	* Slice (first value in slice header is a pointer)
	//	* String (first value in string header is a pointer)
	isPointer	bool

	// basicWidth reports the width of the field assuming that it is directly
	// embedded in the struct (as is the case for basic types in proto3).
	// The possible values are:
	// 	0: invalid
	//	1: bool
	//	4: int32, uint32, float32
	//	8: int64, uint64, float64
	basicWidth	int

	// Where dst and src are pointers to the types being merged.
	merge	func(dst, src pointer)
}

var (
	mergeInfoMap	= map[reflect.Type]*mergeInfo{}
	mergeInfoLock	sync.Mutex
)

func getMergeInfo(t reflect.Type) *mergeInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:92
	_go_fuzz_dep_.CoverTab[110701]++
												mergeInfoLock.Lock()
												defer mergeInfoLock.Unlock()
												mi := mergeInfoMap[t]
												if mi == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:96
		_go_fuzz_dep_.CoverTab[110703]++
													mi = &mergeInfo{typ: t}
													mergeInfoMap[t] = mi
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:98
		// _ = "end of CoverTab[110703]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:99
		_go_fuzz_dep_.CoverTab[110704]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:99
		// _ = "end of CoverTab[110704]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:99
	// _ = "end of CoverTab[110701]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:99
	_go_fuzz_dep_.CoverTab[110702]++
												return mi
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:100
	// _ = "end of CoverTab[110702]"
}

// merge merges src into dst assuming they are both of type *mi.typ.
func (mi *mergeInfo) merge(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:104
	_go_fuzz_dep_.CoverTab[110705]++
												if dst.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:105
		_go_fuzz_dep_.CoverTab[110711]++
													panic("proto: nil destination")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:106
		// _ = "end of CoverTab[110711]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:107
		_go_fuzz_dep_.CoverTab[110712]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:107
		// _ = "end of CoverTab[110712]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:107
	// _ = "end of CoverTab[110705]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:107
	_go_fuzz_dep_.CoverTab[110706]++
												if src.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:108
		_go_fuzz_dep_.CoverTab[110713]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:109
		// _ = "end of CoverTab[110713]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:110
		_go_fuzz_dep_.CoverTab[110714]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:110
		// _ = "end of CoverTab[110714]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:110
	// _ = "end of CoverTab[110706]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:110
	_go_fuzz_dep_.CoverTab[110707]++

												if atomic.LoadInt32(&mi.initialized) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:112
		_go_fuzz_dep_.CoverTab[110715]++
													mi.computeMergeInfo()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:113
		// _ = "end of CoverTab[110715]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:114
		_go_fuzz_dep_.CoverTab[110716]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:114
		// _ = "end of CoverTab[110716]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:114
	// _ = "end of CoverTab[110707]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:114
	_go_fuzz_dep_.CoverTab[110708]++

												for _, fi := range mi.fields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:116
		_go_fuzz_dep_.CoverTab[110717]++
													sfp := src.offset(fi.field)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:122
		if unsafeAllowed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:122
			_go_fuzz_dep_.CoverTab[110719]++
														if fi.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:123
				_go_fuzz_dep_.CoverTab[110721]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:123
				return sfp.getPointer().isNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:123
				// _ = "end of CoverTab[110721]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:123
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:123
				_go_fuzz_dep_.CoverTab[110722]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:124
				// _ = "end of CoverTab[110722]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:125
				_go_fuzz_dep_.CoverTab[110723]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:125
				// _ = "end of CoverTab[110723]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:125
			// _ = "end of CoverTab[110719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:125
			_go_fuzz_dep_.CoverTab[110720]++
														if fi.basicWidth > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:126
				_go_fuzz_dep_.CoverTab[110724]++
															switch {
				case fi.basicWidth == 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:128
					_go_fuzz_dep_.CoverTab[110729]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:128
					return !*sfp.toBool()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:128
					// _ = "end of CoverTab[110729]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:128
				}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:128
					_go_fuzz_dep_.CoverTab[110725]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:129
					// _ = "end of CoverTab[110725]"
				case fi.basicWidth == 4 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:130
					_go_fuzz_dep_.CoverTab[110730]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:130
					return *sfp.toUint32() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:130
					// _ = "end of CoverTab[110730]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:130
				}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:130
					_go_fuzz_dep_.CoverTab[110726]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:131
					// _ = "end of CoverTab[110726]"
				case fi.basicWidth == 8 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:132
					_go_fuzz_dep_.CoverTab[110731]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:132
					return *sfp.toUint64() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:132
					// _ = "end of CoverTab[110731]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:132
				}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:132
					_go_fuzz_dep_.CoverTab[110727]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:133
					// _ = "end of CoverTab[110727]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:133
				default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:133
					_go_fuzz_dep_.CoverTab[110728]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:133
					// _ = "end of CoverTab[110728]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:134
				// _ = "end of CoverTab[110724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:135
				_go_fuzz_dep_.CoverTab[110732]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:135
				// _ = "end of CoverTab[110732]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:135
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:135
			// _ = "end of CoverTab[110720]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:136
			_go_fuzz_dep_.CoverTab[110733]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:136
			// _ = "end of CoverTab[110733]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:136
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:136
		// _ = "end of CoverTab[110717]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:136
		_go_fuzz_dep_.CoverTab[110718]++

													dfp := dst.offset(fi.field)
													fi.merge(dfp, sfp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:139
		// _ = "end of CoverTab[110718]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:140
	// _ = "end of CoverTab[110708]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:140
	_go_fuzz_dep_.CoverTab[110709]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:143
	out := dst.asPointerTo(mi.typ).Elem()
	in := src.asPointerTo(mi.typ).Elem()
	if emIn, err := extendable(in.Addr().Interface()); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:145
		_go_fuzz_dep_.CoverTab[110734]++
													emOut, _ := extendable(out.Addr().Interface())
													mIn, muIn := emIn.extensionsRead()
													if mIn != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:148
			_go_fuzz_dep_.CoverTab[110735]++
														mOut := emOut.extensionsWrite()
														muIn.Lock()
														mergeExtension(mOut, mIn)
														muIn.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:152
			// _ = "end of CoverTab[110735]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:153
			_go_fuzz_dep_.CoverTab[110736]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:153
			// _ = "end of CoverTab[110736]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:153
		// _ = "end of CoverTab[110734]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:154
		_go_fuzz_dep_.CoverTab[110737]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:154
		// _ = "end of CoverTab[110737]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:154
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:154
	// _ = "end of CoverTab[110709]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:154
	_go_fuzz_dep_.CoverTab[110710]++

												if mi.unrecognized.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:156
		_go_fuzz_dep_.CoverTab[110738]++
													if b := *src.offset(mi.unrecognized).toBytes(); len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:157
			_go_fuzz_dep_.CoverTab[110739]++
														*dst.offset(mi.unrecognized).toBytes() = append([]byte(nil), b...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:158
			// _ = "end of CoverTab[110739]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:159
			_go_fuzz_dep_.CoverTab[110740]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:159
			// _ = "end of CoverTab[110740]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:159
		// _ = "end of CoverTab[110738]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:160
		_go_fuzz_dep_.CoverTab[110741]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:160
		// _ = "end of CoverTab[110741]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:160
	// _ = "end of CoverTab[110710]"
}

func (mi *mergeInfo) computeMergeInfo() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:163
	_go_fuzz_dep_.CoverTab[110742]++
												mi.lock.Lock()
												defer mi.lock.Unlock()
												if mi.initialized != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:166
		_go_fuzz_dep_.CoverTab[110746]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:167
		// _ = "end of CoverTab[110746]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:168
		_go_fuzz_dep_.CoverTab[110747]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:168
		// _ = "end of CoverTab[110747]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:168
	// _ = "end of CoverTab[110742]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:168
	_go_fuzz_dep_.CoverTab[110743]++
												t := mi.typ
												n := t.NumField()

												props := GetProperties(t)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:173
		_go_fuzz_dep_.CoverTab[110748]++
													f := t.Field(i)
													if strings.HasPrefix(f.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:175
			_go_fuzz_dep_.CoverTab[110755]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:176
			// _ = "end of CoverTab[110755]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:177
			_go_fuzz_dep_.CoverTab[110756]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:177
			// _ = "end of CoverTab[110756]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:177
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:177
		// _ = "end of CoverTab[110748]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:177
		_go_fuzz_dep_.CoverTab[110749]++

													mfi := mergeFieldInfo{field: toField(&f)}
													tf := f.Type

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:185
		if unsafeAllowed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:185
			_go_fuzz_dep_.CoverTab[110757]++
														switch tf.Kind() {
			case reflect.Ptr, reflect.Slice, reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:187
				_go_fuzz_dep_.CoverTab[110758]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:191
				mfi.isPointer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:191
				// _ = "end of CoverTab[110758]"
			case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:192
				_go_fuzz_dep_.CoverTab[110759]++
															mfi.basicWidth = 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:193
				// _ = "end of CoverTab[110759]"
			case reflect.Int32, reflect.Uint32, reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:194
				_go_fuzz_dep_.CoverTab[110760]++
															mfi.basicWidth = 4
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:195
				// _ = "end of CoverTab[110760]"
			case reflect.Int64, reflect.Uint64, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:196
				_go_fuzz_dep_.CoverTab[110761]++
															mfi.basicWidth = 8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:197
				// _ = "end of CoverTab[110761]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:197
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:197
				_go_fuzz_dep_.CoverTab[110762]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:197
				// _ = "end of CoverTab[110762]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:198
			// _ = "end of CoverTab[110757]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:199
			_go_fuzz_dep_.CoverTab[110763]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:199
			// _ = "end of CoverTab[110763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:199
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:199
		// _ = "end of CoverTab[110749]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:199
		_go_fuzz_dep_.CoverTab[110750]++

		// Unwrap tf to get at its most basic type.
		var isPointer, isSlice bool
		if tf.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:203
			_go_fuzz_dep_.CoverTab[110764]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:203
			return tf.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:203
			// _ = "end of CoverTab[110764]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:203
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:203
			_go_fuzz_dep_.CoverTab[110765]++
														isSlice = true
														tf = tf.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:205
			// _ = "end of CoverTab[110765]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:206
			_go_fuzz_dep_.CoverTab[110766]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:206
			// _ = "end of CoverTab[110766]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:206
		// _ = "end of CoverTab[110750]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:206
		_go_fuzz_dep_.CoverTab[110751]++
													if tf.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:207
			_go_fuzz_dep_.CoverTab[110767]++
														isPointer = true
														tf = tf.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:209
			// _ = "end of CoverTab[110767]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:210
			_go_fuzz_dep_.CoverTab[110768]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:210
			// _ = "end of CoverTab[110768]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:210
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:210
		// _ = "end of CoverTab[110751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:210
		_go_fuzz_dep_.CoverTab[110752]++
													if isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			_go_fuzz_dep_.CoverTab[110769]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			// _ = "end of CoverTab[110769]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			_go_fuzz_dep_.CoverTab[110770]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			return tf.Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			// _ = "end of CoverTab[110770]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:211
			_go_fuzz_dep_.CoverTab[110771]++
														panic("both pointer and slice for basic type in " + tf.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:212
			// _ = "end of CoverTab[110771]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:213
			_go_fuzz_dep_.CoverTab[110772]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:213
			// _ = "end of CoverTab[110772]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:213
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:213
		// _ = "end of CoverTab[110752]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:213
		_go_fuzz_dep_.CoverTab[110753]++

													switch tf.Kind() {
		case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:216
			_go_fuzz_dep_.CoverTab[110773]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:218
				_go_fuzz_dep_.CoverTab[110786]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:219
					_go_fuzz_dep_.CoverTab[110789]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:231
					sfs := src.getInt32Slice()
					if sfs != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:232
						_go_fuzz_dep_.CoverTab[110790]++
																	dfs := dst.getInt32Slice()
																	dfs = append(dfs, sfs...)
																	if dfs == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:235
							_go_fuzz_dep_.CoverTab[110792]++
																		dfs = []int32{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:236
							// _ = "end of CoverTab[110792]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:237
							_go_fuzz_dep_.CoverTab[110793]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:237
							// _ = "end of CoverTab[110793]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:237
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:237
						// _ = "end of CoverTab[110790]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:237
						_go_fuzz_dep_.CoverTab[110791]++
																	dst.setInt32Slice(dfs)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:238
						// _ = "end of CoverTab[110791]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:239
						_go_fuzz_dep_.CoverTab[110794]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:239
						// _ = "end of CoverTab[110794]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:239
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:239
					// _ = "end of CoverTab[110789]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:240
				// _ = "end of CoverTab[110786]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:241
				_go_fuzz_dep_.CoverTab[110787]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:242
					_go_fuzz_dep_.CoverTab[110795]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:255
					sfp := src.getInt32Ptr()
					if sfp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:256
						_go_fuzz_dep_.CoverTab[110796]++
																	dfp := dst.getInt32Ptr()
																	if dfp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:258
							_go_fuzz_dep_.CoverTab[110797]++
																		dst.setInt32Ptr(*sfp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:259
							// _ = "end of CoverTab[110797]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:260
							_go_fuzz_dep_.CoverTab[110798]++
																		*dfp = *sfp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:261
							// _ = "end of CoverTab[110798]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:262
						// _ = "end of CoverTab[110796]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:263
						_go_fuzz_dep_.CoverTab[110799]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:263
						// _ = "end of CoverTab[110799]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:263
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:263
					// _ = "end of CoverTab[110795]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:264
				// _ = "end of CoverTab[110787]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:265
				_go_fuzz_dep_.CoverTab[110788]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:266
					_go_fuzz_dep_.CoverTab[110800]++
																if v := *src.toInt32(); v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:267
						_go_fuzz_dep_.CoverTab[110801]++
																	*dst.toInt32() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:268
						// _ = "end of CoverTab[110801]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:269
						_go_fuzz_dep_.CoverTab[110802]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:269
						// _ = "end of CoverTab[110802]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:269
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:269
					// _ = "end of CoverTab[110800]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:270
				// _ = "end of CoverTab[110788]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:271
			// _ = "end of CoverTab[110773]"
		case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:272
			_go_fuzz_dep_.CoverTab[110774]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:274
				_go_fuzz_dep_.CoverTab[110803]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:275
					_go_fuzz_dep_.CoverTab[110806]++
																sfsp := src.toInt64Slice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:277
						_go_fuzz_dep_.CoverTab[110807]++
																	dfsp := dst.toInt64Slice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:280
							_go_fuzz_dep_.CoverTab[110808]++
																		*dfsp = []int64{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:281
							// _ = "end of CoverTab[110808]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:282
							_go_fuzz_dep_.CoverTab[110809]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:282
							// _ = "end of CoverTab[110809]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:282
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:282
						// _ = "end of CoverTab[110807]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:283
						_go_fuzz_dep_.CoverTab[110810]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:283
						// _ = "end of CoverTab[110810]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:283
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:283
					// _ = "end of CoverTab[110806]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:284
				// _ = "end of CoverTab[110803]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:285
				_go_fuzz_dep_.CoverTab[110804]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:286
					_go_fuzz_dep_.CoverTab[110811]++
																sfpp := src.toInt64Ptr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:288
						_go_fuzz_dep_.CoverTab[110812]++
																	dfpp := dst.toInt64Ptr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:290
							_go_fuzz_dep_.CoverTab[110813]++
																		*dfpp = Int64(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:291
							// _ = "end of CoverTab[110813]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:292
							_go_fuzz_dep_.CoverTab[110814]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:293
							// _ = "end of CoverTab[110814]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:294
						// _ = "end of CoverTab[110812]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:295
						_go_fuzz_dep_.CoverTab[110815]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:295
						// _ = "end of CoverTab[110815]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:295
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:295
					// _ = "end of CoverTab[110811]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:296
				// _ = "end of CoverTab[110804]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:297
				_go_fuzz_dep_.CoverTab[110805]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:298
					_go_fuzz_dep_.CoverTab[110816]++
																if v := *src.toInt64(); v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:299
						_go_fuzz_dep_.CoverTab[110817]++
																	*dst.toInt64() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:300
						// _ = "end of CoverTab[110817]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:301
						_go_fuzz_dep_.CoverTab[110818]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:301
						// _ = "end of CoverTab[110818]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:301
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:301
					// _ = "end of CoverTab[110816]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:302
				// _ = "end of CoverTab[110805]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:303
			// _ = "end of CoverTab[110774]"
		case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:304
			_go_fuzz_dep_.CoverTab[110775]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:306
				_go_fuzz_dep_.CoverTab[110819]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:307
					_go_fuzz_dep_.CoverTab[110822]++
																sfsp := src.toUint32Slice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:309
						_go_fuzz_dep_.CoverTab[110823]++
																	dfsp := dst.toUint32Slice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:312
							_go_fuzz_dep_.CoverTab[110824]++
																		*dfsp = []uint32{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:313
							// _ = "end of CoverTab[110824]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:314
							_go_fuzz_dep_.CoverTab[110825]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:314
							// _ = "end of CoverTab[110825]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:314
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:314
						// _ = "end of CoverTab[110823]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:315
						_go_fuzz_dep_.CoverTab[110826]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:315
						// _ = "end of CoverTab[110826]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:315
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:315
					// _ = "end of CoverTab[110822]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:316
				// _ = "end of CoverTab[110819]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:317
				_go_fuzz_dep_.CoverTab[110820]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:318
					_go_fuzz_dep_.CoverTab[110827]++
																sfpp := src.toUint32Ptr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:320
						_go_fuzz_dep_.CoverTab[110828]++
																	dfpp := dst.toUint32Ptr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:322
							_go_fuzz_dep_.CoverTab[110829]++
																		*dfpp = Uint32(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:323
							// _ = "end of CoverTab[110829]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:324
							_go_fuzz_dep_.CoverTab[110830]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:325
							// _ = "end of CoverTab[110830]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:326
						// _ = "end of CoverTab[110828]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:327
						_go_fuzz_dep_.CoverTab[110831]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:327
						// _ = "end of CoverTab[110831]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:327
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:327
					// _ = "end of CoverTab[110827]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:328
				// _ = "end of CoverTab[110820]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:329
				_go_fuzz_dep_.CoverTab[110821]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:330
					_go_fuzz_dep_.CoverTab[110832]++
																if v := *src.toUint32(); v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:331
						_go_fuzz_dep_.CoverTab[110833]++
																	*dst.toUint32() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:332
						// _ = "end of CoverTab[110833]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:333
						_go_fuzz_dep_.CoverTab[110834]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:333
						// _ = "end of CoverTab[110834]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:333
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:333
					// _ = "end of CoverTab[110832]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:334
				// _ = "end of CoverTab[110821]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:335
			// _ = "end of CoverTab[110775]"
		case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:336
			_go_fuzz_dep_.CoverTab[110776]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:338
				_go_fuzz_dep_.CoverTab[110835]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:339
					_go_fuzz_dep_.CoverTab[110838]++
																sfsp := src.toUint64Slice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:341
						_go_fuzz_dep_.CoverTab[110839]++
																	dfsp := dst.toUint64Slice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:344
							_go_fuzz_dep_.CoverTab[110840]++
																		*dfsp = []uint64{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:345
							// _ = "end of CoverTab[110840]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:346
							_go_fuzz_dep_.CoverTab[110841]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:346
							// _ = "end of CoverTab[110841]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:346
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:346
						// _ = "end of CoverTab[110839]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:347
						_go_fuzz_dep_.CoverTab[110842]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:347
						// _ = "end of CoverTab[110842]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:347
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:347
					// _ = "end of CoverTab[110838]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:348
				// _ = "end of CoverTab[110835]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:349
				_go_fuzz_dep_.CoverTab[110836]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:350
					_go_fuzz_dep_.CoverTab[110843]++
																sfpp := src.toUint64Ptr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:352
						_go_fuzz_dep_.CoverTab[110844]++
																	dfpp := dst.toUint64Ptr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:354
							_go_fuzz_dep_.CoverTab[110845]++
																		*dfpp = Uint64(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:355
							// _ = "end of CoverTab[110845]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:356
							_go_fuzz_dep_.CoverTab[110846]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:357
							// _ = "end of CoverTab[110846]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:358
						// _ = "end of CoverTab[110844]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:359
						_go_fuzz_dep_.CoverTab[110847]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:359
						// _ = "end of CoverTab[110847]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:359
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:359
					// _ = "end of CoverTab[110843]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:360
				// _ = "end of CoverTab[110836]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:361
				_go_fuzz_dep_.CoverTab[110837]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:362
					_go_fuzz_dep_.CoverTab[110848]++
																if v := *src.toUint64(); v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:363
						_go_fuzz_dep_.CoverTab[110849]++
																	*dst.toUint64() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:364
						// _ = "end of CoverTab[110849]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:365
						_go_fuzz_dep_.CoverTab[110850]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:365
						// _ = "end of CoverTab[110850]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:365
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:365
					// _ = "end of CoverTab[110848]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:366
				// _ = "end of CoverTab[110837]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:367
			// _ = "end of CoverTab[110776]"
		case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:368
			_go_fuzz_dep_.CoverTab[110777]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:370
				_go_fuzz_dep_.CoverTab[110851]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:371
					_go_fuzz_dep_.CoverTab[110854]++
																sfsp := src.toFloat32Slice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:373
						_go_fuzz_dep_.CoverTab[110855]++
																	dfsp := dst.toFloat32Slice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:376
							_go_fuzz_dep_.CoverTab[110856]++
																		*dfsp = []float32{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:377
							// _ = "end of CoverTab[110856]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:378
							_go_fuzz_dep_.CoverTab[110857]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:378
							// _ = "end of CoverTab[110857]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:378
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:378
						// _ = "end of CoverTab[110855]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:379
						_go_fuzz_dep_.CoverTab[110858]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:379
						// _ = "end of CoverTab[110858]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:379
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:379
					// _ = "end of CoverTab[110854]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:380
				// _ = "end of CoverTab[110851]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:381
				_go_fuzz_dep_.CoverTab[110852]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:382
					_go_fuzz_dep_.CoverTab[110859]++
																sfpp := src.toFloat32Ptr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:384
						_go_fuzz_dep_.CoverTab[110860]++
																	dfpp := dst.toFloat32Ptr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:386
							_go_fuzz_dep_.CoverTab[110861]++
																		*dfpp = Float32(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:387
							// _ = "end of CoverTab[110861]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:388
							_go_fuzz_dep_.CoverTab[110862]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:389
							// _ = "end of CoverTab[110862]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:390
						// _ = "end of CoverTab[110860]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:391
						_go_fuzz_dep_.CoverTab[110863]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:391
						// _ = "end of CoverTab[110863]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:391
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:391
					// _ = "end of CoverTab[110859]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:392
				// _ = "end of CoverTab[110852]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:393
				_go_fuzz_dep_.CoverTab[110853]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:394
					_go_fuzz_dep_.CoverTab[110864]++
																if v := *src.toFloat32(); v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:395
						_go_fuzz_dep_.CoverTab[110865]++
																	*dst.toFloat32() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:396
						// _ = "end of CoverTab[110865]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:397
						_go_fuzz_dep_.CoverTab[110866]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:397
						// _ = "end of CoverTab[110866]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:397
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:397
					// _ = "end of CoverTab[110864]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:398
				// _ = "end of CoverTab[110853]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:399
			// _ = "end of CoverTab[110777]"
		case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:400
			_go_fuzz_dep_.CoverTab[110778]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:402
				_go_fuzz_dep_.CoverTab[110867]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:403
					_go_fuzz_dep_.CoverTab[110870]++
																sfsp := src.toFloat64Slice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:405
						_go_fuzz_dep_.CoverTab[110871]++
																	dfsp := dst.toFloat64Slice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:408
							_go_fuzz_dep_.CoverTab[110872]++
																		*dfsp = []float64{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:409
							// _ = "end of CoverTab[110872]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:410
							_go_fuzz_dep_.CoverTab[110873]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:410
							// _ = "end of CoverTab[110873]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:410
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:410
						// _ = "end of CoverTab[110871]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:411
						_go_fuzz_dep_.CoverTab[110874]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:411
						// _ = "end of CoverTab[110874]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:411
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:411
					// _ = "end of CoverTab[110870]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:412
				// _ = "end of CoverTab[110867]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:413
				_go_fuzz_dep_.CoverTab[110868]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:414
					_go_fuzz_dep_.CoverTab[110875]++
																sfpp := src.toFloat64Ptr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:416
						_go_fuzz_dep_.CoverTab[110876]++
																	dfpp := dst.toFloat64Ptr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:418
							_go_fuzz_dep_.CoverTab[110877]++
																		*dfpp = Float64(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:419
							// _ = "end of CoverTab[110877]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:420
							_go_fuzz_dep_.CoverTab[110878]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:421
							// _ = "end of CoverTab[110878]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:422
						// _ = "end of CoverTab[110876]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:423
						_go_fuzz_dep_.CoverTab[110879]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:423
						// _ = "end of CoverTab[110879]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:423
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:423
					// _ = "end of CoverTab[110875]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:424
				// _ = "end of CoverTab[110868]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:425
				_go_fuzz_dep_.CoverTab[110869]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:426
					_go_fuzz_dep_.CoverTab[110880]++
																if v := *src.toFloat64(); v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:427
						_go_fuzz_dep_.CoverTab[110881]++
																	*dst.toFloat64() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:428
						// _ = "end of CoverTab[110881]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:429
						_go_fuzz_dep_.CoverTab[110882]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:429
						// _ = "end of CoverTab[110882]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:429
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:429
					// _ = "end of CoverTab[110880]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:430
				// _ = "end of CoverTab[110869]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:431
			// _ = "end of CoverTab[110778]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:432
			_go_fuzz_dep_.CoverTab[110779]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:434
				_go_fuzz_dep_.CoverTab[110883]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:435
					_go_fuzz_dep_.CoverTab[110886]++
																sfsp := src.toBoolSlice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:437
						_go_fuzz_dep_.CoverTab[110887]++
																	dfsp := dst.toBoolSlice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:440
							_go_fuzz_dep_.CoverTab[110888]++
																		*dfsp = []bool{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:441
							// _ = "end of CoverTab[110888]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:442
							_go_fuzz_dep_.CoverTab[110889]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:442
							// _ = "end of CoverTab[110889]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:442
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:442
						// _ = "end of CoverTab[110887]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:443
						_go_fuzz_dep_.CoverTab[110890]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:443
						// _ = "end of CoverTab[110890]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:443
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:443
					// _ = "end of CoverTab[110886]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:444
				// _ = "end of CoverTab[110883]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:445
				_go_fuzz_dep_.CoverTab[110884]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:446
					_go_fuzz_dep_.CoverTab[110891]++
																sfpp := src.toBoolPtr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:448
						_go_fuzz_dep_.CoverTab[110892]++
																	dfpp := dst.toBoolPtr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:450
							_go_fuzz_dep_.CoverTab[110893]++
																		*dfpp = Bool(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:451
							// _ = "end of CoverTab[110893]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:452
							_go_fuzz_dep_.CoverTab[110894]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:453
							// _ = "end of CoverTab[110894]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:454
						// _ = "end of CoverTab[110892]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:455
						_go_fuzz_dep_.CoverTab[110895]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:455
						// _ = "end of CoverTab[110895]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:455
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:455
					// _ = "end of CoverTab[110891]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:456
				// _ = "end of CoverTab[110884]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:457
				_go_fuzz_dep_.CoverTab[110885]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:458
					_go_fuzz_dep_.CoverTab[110896]++
																if v := *src.toBool(); v {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:459
						_go_fuzz_dep_.CoverTab[110897]++
																	*dst.toBool() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:460
						// _ = "end of CoverTab[110897]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:461
						_go_fuzz_dep_.CoverTab[110898]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:461
						// _ = "end of CoverTab[110898]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:461
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:461
					// _ = "end of CoverTab[110896]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:462
				// _ = "end of CoverTab[110885]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:463
			// _ = "end of CoverTab[110779]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:464
			_go_fuzz_dep_.CoverTab[110780]++
														switch {
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:466
				_go_fuzz_dep_.CoverTab[110899]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:467
					_go_fuzz_dep_.CoverTab[110902]++
																sfsp := src.toStringSlice()
																if *sfsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:469
						_go_fuzz_dep_.CoverTab[110903]++
																	dfsp := dst.toStringSlice()
																	*dfsp = append(*dfsp, *sfsp...)
																	if *dfsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:472
							_go_fuzz_dep_.CoverTab[110904]++
																		*dfsp = []string{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:473
							// _ = "end of CoverTab[110904]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:474
							_go_fuzz_dep_.CoverTab[110905]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:474
							// _ = "end of CoverTab[110905]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:474
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:474
						// _ = "end of CoverTab[110903]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:475
						_go_fuzz_dep_.CoverTab[110906]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:475
						// _ = "end of CoverTab[110906]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:475
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:475
					// _ = "end of CoverTab[110902]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:476
				// _ = "end of CoverTab[110899]"
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:477
				_go_fuzz_dep_.CoverTab[110900]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:478
					_go_fuzz_dep_.CoverTab[110907]++
																sfpp := src.toStringPtr()
																if *sfpp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:480
						_go_fuzz_dep_.CoverTab[110908]++
																	dfpp := dst.toStringPtr()
																	if *dfpp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:482
							_go_fuzz_dep_.CoverTab[110909]++
																		*dfpp = String(**sfpp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:483
							// _ = "end of CoverTab[110909]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:484
							_go_fuzz_dep_.CoverTab[110910]++
																		**dfpp = **sfpp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:485
							// _ = "end of CoverTab[110910]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:486
						// _ = "end of CoverTab[110908]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:487
						_go_fuzz_dep_.CoverTab[110911]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:487
						// _ = "end of CoverTab[110911]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:487
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:487
					// _ = "end of CoverTab[110907]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:488
				// _ = "end of CoverTab[110900]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:489
				_go_fuzz_dep_.CoverTab[110901]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:490
					_go_fuzz_dep_.CoverTab[110912]++
																if v := *src.toString(); v != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:491
						_go_fuzz_dep_.CoverTab[110913]++
																	*dst.toString() = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:492
						// _ = "end of CoverTab[110913]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:493
						_go_fuzz_dep_.CoverTab[110914]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:493
						// _ = "end of CoverTab[110914]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:493
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:493
					// _ = "end of CoverTab[110912]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:494
				// _ = "end of CoverTab[110901]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:495
			// _ = "end of CoverTab[110780]"
		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:496
			_go_fuzz_dep_.CoverTab[110781]++
														isProto3 := props.Prop[i].proto3
														switch {
			case isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:499
				_go_fuzz_dep_.CoverTab[110915]++
															panic("bad pointer in byte slice case in " + tf.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:500
				// _ = "end of CoverTab[110915]"
			case tf.Elem().Kind() != reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:501
				_go_fuzz_dep_.CoverTab[110916]++
															panic("bad element kind in byte slice case in " + tf.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:502
				// _ = "end of CoverTab[110916]"
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:503
				_go_fuzz_dep_.CoverTab[110917]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:504
					_go_fuzz_dep_.CoverTab[110919]++
																sbsp := src.toBytesSlice()
																if *sbsp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:506
						_go_fuzz_dep_.CoverTab[110920]++
																	dbsp := dst.toBytesSlice()
																	for _, sb := range *sbsp {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:508
							_go_fuzz_dep_.CoverTab[110922]++
																		if sb == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:509
								_go_fuzz_dep_.CoverTab[110923]++
																			*dbsp = append(*dbsp, nil)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:510
								// _ = "end of CoverTab[110923]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:511
								_go_fuzz_dep_.CoverTab[110924]++
																			*dbsp = append(*dbsp, append([]byte{}, sb...))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:512
								// _ = "end of CoverTab[110924]"
							}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:513
							// _ = "end of CoverTab[110922]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:514
						// _ = "end of CoverTab[110920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:514
						_go_fuzz_dep_.CoverTab[110921]++
																	if *dbsp == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:515
							_go_fuzz_dep_.CoverTab[110925]++
																		*dbsp = [][]byte{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:516
							// _ = "end of CoverTab[110925]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:517
							_go_fuzz_dep_.CoverTab[110926]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:517
							// _ = "end of CoverTab[110926]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:517
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:517
						// _ = "end of CoverTab[110921]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:518
						_go_fuzz_dep_.CoverTab[110927]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:518
						// _ = "end of CoverTab[110927]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:518
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:518
					// _ = "end of CoverTab[110919]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:519
				// _ = "end of CoverTab[110917]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:520
				_go_fuzz_dep_.CoverTab[110918]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:521
					_go_fuzz_dep_.CoverTab[110928]++
																sbp := src.toBytes()
																if *sbp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:523
						_go_fuzz_dep_.CoverTab[110929]++
																	dbp := dst.toBytes()
																	if !isProto3 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:525
							_go_fuzz_dep_.CoverTab[110930]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:525
							return len(*sbp) > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:525
							// _ = "end of CoverTab[110930]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:525
						}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:525
							_go_fuzz_dep_.CoverTab[110931]++
																		*dbp = append([]byte{}, *sbp...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:526
							// _ = "end of CoverTab[110931]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:527
							_go_fuzz_dep_.CoverTab[110932]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:527
							// _ = "end of CoverTab[110932]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:527
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:527
						// _ = "end of CoverTab[110929]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:528
						_go_fuzz_dep_.CoverTab[110933]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:528
						// _ = "end of CoverTab[110933]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:528
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:528
					// _ = "end of CoverTab[110928]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:529
				// _ = "end of CoverTab[110918]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:530
			// _ = "end of CoverTab[110781]"
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:531
			_go_fuzz_dep_.CoverTab[110782]++
														switch {
			case isSlice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:533
				_go_fuzz_dep_.CoverTab[110938]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:533
				return !isPointer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:533
				// _ = "end of CoverTab[110938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:533
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:533
				_go_fuzz_dep_.CoverTab[110934]++
															mergeInfo := getMergeInfo(tf)
															zero := reflect.Zero(tf)
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:536
					_go_fuzz_dep_.CoverTab[110939]++

																dstsp := dst.asPointerTo(f.Type)
																dsts := dstsp.Elem()
																srcs := src.asPointerTo(f.Type).Elem()
																for i := 0; i < srcs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:541
						_go_fuzz_dep_.CoverTab[110942]++
																	dsts = reflect.Append(dsts, zero)
																	srcElement := srcs.Index(i).Addr()
																	dstElement := dsts.Index(dsts.Len() - 1).Addr()
																	mergeInfo.merge(valToPointer(dstElement), valToPointer(srcElement))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:545
						// _ = "end of CoverTab[110942]"
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:546
					// _ = "end of CoverTab[110939]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:546
					_go_fuzz_dep_.CoverTab[110940]++
																if dsts.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:547
						_go_fuzz_dep_.CoverTab[110943]++
																	dsts = reflect.MakeSlice(f.Type, 0, 0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:548
						// _ = "end of CoverTab[110943]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:549
						_go_fuzz_dep_.CoverTab[110944]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:549
						// _ = "end of CoverTab[110944]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:549
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:549
					// _ = "end of CoverTab[110940]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:549
					_go_fuzz_dep_.CoverTab[110941]++
																dstsp.Elem().Set(dsts)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:550
					// _ = "end of CoverTab[110941]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:551
				// _ = "end of CoverTab[110934]"
			case !isPointer:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:552
				_go_fuzz_dep_.CoverTab[110935]++
															mergeInfo := getMergeInfo(tf)
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:554
					_go_fuzz_dep_.CoverTab[110945]++
																mergeInfo.merge(dst, src)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:555
					// _ = "end of CoverTab[110945]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:556
				// _ = "end of CoverTab[110935]"
			case isSlice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:557
				_go_fuzz_dep_.CoverTab[110936]++
															mergeInfo := getMergeInfo(tf)
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:559
					_go_fuzz_dep_.CoverTab[110946]++
																sps := src.getPointerSlice()
																if sps != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:561
						_go_fuzz_dep_.CoverTab[110947]++
																	dps := dst.getPointerSlice()
																	for _, sp := range sps {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:563
							_go_fuzz_dep_.CoverTab[110950]++
																		var dp pointer
																		if !sp.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:565
								_go_fuzz_dep_.CoverTab[110952]++
																			dp = valToPointer(reflect.New(tf))
																			mergeInfo.merge(dp, sp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:567
								// _ = "end of CoverTab[110952]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:568
								_go_fuzz_dep_.CoverTab[110953]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:568
								// _ = "end of CoverTab[110953]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:568
							}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:568
							// _ = "end of CoverTab[110950]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:568
							_go_fuzz_dep_.CoverTab[110951]++
																		dps = append(dps, dp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:569
							// _ = "end of CoverTab[110951]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:570
						// _ = "end of CoverTab[110947]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:570
						_go_fuzz_dep_.CoverTab[110948]++
																	if dps == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:571
							_go_fuzz_dep_.CoverTab[110954]++
																		dps = []pointer{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:572
							// _ = "end of CoverTab[110954]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:573
							_go_fuzz_dep_.CoverTab[110955]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:573
							// _ = "end of CoverTab[110955]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:573
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:573
						// _ = "end of CoverTab[110948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:573
						_go_fuzz_dep_.CoverTab[110949]++
																	dst.setPointerSlice(dps)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:574
						// _ = "end of CoverTab[110949]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:575
						_go_fuzz_dep_.CoverTab[110956]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:575
						// _ = "end of CoverTab[110956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:575
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:575
					// _ = "end of CoverTab[110946]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:576
				// _ = "end of CoverTab[110936]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:577
				_go_fuzz_dep_.CoverTab[110937]++
															mergeInfo := getMergeInfo(tf)
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:579
					_go_fuzz_dep_.CoverTab[110957]++
																sp := src.getPointer()
																if !sp.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:581
						_go_fuzz_dep_.CoverTab[110958]++
																	dp := dst.getPointer()
																	if dp.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:583
							_go_fuzz_dep_.CoverTab[110960]++
																		dp = valToPointer(reflect.New(tf))
																		dst.setPointer(dp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:585
							// _ = "end of CoverTab[110960]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:586
							_go_fuzz_dep_.CoverTab[110961]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:586
							// _ = "end of CoverTab[110961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:586
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:586
						// _ = "end of CoverTab[110958]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:586
						_go_fuzz_dep_.CoverTab[110959]++
																	mergeInfo.merge(dp, sp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:587
						// _ = "end of CoverTab[110959]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:588
						_go_fuzz_dep_.CoverTab[110962]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:588
						// _ = "end of CoverTab[110962]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:588
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:588
					// _ = "end of CoverTab[110957]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:589
				// _ = "end of CoverTab[110937]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:590
			// _ = "end of CoverTab[110782]"
		case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:591
			_go_fuzz_dep_.CoverTab[110783]++
														switch {
			case isPointer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:593
				_go_fuzz_dep_.CoverTab[110965]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:593
				return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:593
				// _ = "end of CoverTab[110965]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:593
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:593
				_go_fuzz_dep_.CoverTab[110963]++
															panic("bad pointer or slice in map case in " + tf.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:594
				// _ = "end of CoverTab[110963]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:595
				_go_fuzz_dep_.CoverTab[110964]++
															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:596
					_go_fuzz_dep_.CoverTab[110966]++
																sm := src.asPointerTo(tf).Elem()
																if sm.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:598
						_go_fuzz_dep_.CoverTab[110969]++
																	return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:599
						// _ = "end of CoverTab[110969]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:600
						_go_fuzz_dep_.CoverTab[110970]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:600
						// _ = "end of CoverTab[110970]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:600
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:600
					// _ = "end of CoverTab[110966]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:600
					_go_fuzz_dep_.CoverTab[110967]++
																dm := dst.asPointerTo(tf).Elem()
																if dm.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:602
						_go_fuzz_dep_.CoverTab[110971]++
																	dm.Set(reflect.MakeMap(tf))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:603
						// _ = "end of CoverTab[110971]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:604
						_go_fuzz_dep_.CoverTab[110972]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:604
						// _ = "end of CoverTab[110972]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:604
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:604
					// _ = "end of CoverTab[110967]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:604
					_go_fuzz_dep_.CoverTab[110968]++

																switch tf.Elem().Kind() {
					case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:607
						_go_fuzz_dep_.CoverTab[110973]++
																	for _, key := range sm.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:608
							_go_fuzz_dep_.CoverTab[110976]++
																		val := sm.MapIndex(key)
																		val = reflect.ValueOf(Clone(val.Interface().(Message)))
																		dm.SetMapIndex(key, val)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:611
							// _ = "end of CoverTab[110976]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:612
						// _ = "end of CoverTab[110973]"
					case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:613
						_go_fuzz_dep_.CoverTab[110974]++
																	for _, key := range sm.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:614
							_go_fuzz_dep_.CoverTab[110977]++
																		val := sm.MapIndex(key)
																		val = reflect.ValueOf(append([]byte{}, val.Bytes()...))
																		dm.SetMapIndex(key, val)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:617
							// _ = "end of CoverTab[110977]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:618
						// _ = "end of CoverTab[110974]"
					default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:619
						_go_fuzz_dep_.CoverTab[110975]++
																	for _, key := range sm.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:620
							_go_fuzz_dep_.CoverTab[110978]++
																		val := sm.MapIndex(key)
																		dm.SetMapIndex(key, val)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:622
							// _ = "end of CoverTab[110978]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:623
						// _ = "end of CoverTab[110975]"
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:624
					// _ = "end of CoverTab[110968]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:625
				// _ = "end of CoverTab[110964]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:626
			// _ = "end of CoverTab[110783]"
		case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:627
			_go_fuzz_dep_.CoverTab[110784]++

														switch {
			case isPointer || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:630
				_go_fuzz_dep_.CoverTab[110981]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:630
				return isSlice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:630
				// _ = "end of CoverTab[110981]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:630
			}():
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:630
				_go_fuzz_dep_.CoverTab[110979]++
															panic("bad pointer or slice in interface case in " + tf.Name())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:631
				// _ = "end of CoverTab[110979]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:632
				_go_fuzz_dep_.CoverTab[110980]++

															mfi.merge = func(dst, src pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:634
					_go_fuzz_dep_.CoverTab[110982]++
																su := src.asPointerTo(tf).Elem()
																if !su.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:636
						_go_fuzz_dep_.CoverTab[110983]++
																	du := dst.asPointerTo(tf).Elem()
																	typ := su.Elem().Type()
																	if du.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:639
							_go_fuzz_dep_.CoverTab[110987]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:639
							return du.Elem().Type() != typ
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:639
							// _ = "end of CoverTab[110987]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:639
						}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:639
							_go_fuzz_dep_.CoverTab[110988]++
																		du.Set(reflect.New(typ.Elem()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:640
							// _ = "end of CoverTab[110988]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:641
							_go_fuzz_dep_.CoverTab[110989]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:641
							// _ = "end of CoverTab[110989]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:641
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:641
						// _ = "end of CoverTab[110983]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:641
						_go_fuzz_dep_.CoverTab[110984]++
																	sv := su.Elem().Elem().Field(0)
																	if sv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:643
							_go_fuzz_dep_.CoverTab[110990]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:643
							return sv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:643
							// _ = "end of CoverTab[110990]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:643
						}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:643
							_go_fuzz_dep_.CoverTab[110991]++
																		return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:644
							// _ = "end of CoverTab[110991]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:645
							_go_fuzz_dep_.CoverTab[110992]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:645
							// _ = "end of CoverTab[110992]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:645
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:645
						// _ = "end of CoverTab[110984]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:645
						_go_fuzz_dep_.CoverTab[110985]++
																	dv := du.Elem().Elem().Field(0)
																	if dv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:647
							_go_fuzz_dep_.CoverTab[110993]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:647
							return dv.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:647
							// _ = "end of CoverTab[110993]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:647
						}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:647
							_go_fuzz_dep_.CoverTab[110994]++
																		dv.Set(reflect.New(sv.Type().Elem()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:648
							// _ = "end of CoverTab[110994]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:649
							_go_fuzz_dep_.CoverTab[110995]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:649
							// _ = "end of CoverTab[110995]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:649
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:649
						// _ = "end of CoverTab[110985]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:649
						_go_fuzz_dep_.CoverTab[110986]++
																	switch sv.Type().Kind() {
						case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:651
							_go_fuzz_dep_.CoverTab[110996]++
																		Merge(dv.Interface().(Message), sv.Interface().(Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:652
							// _ = "end of CoverTab[110996]"
						case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:653
							_go_fuzz_dep_.CoverTab[110997]++
																		dv.Set(reflect.ValueOf(append([]byte{}, sv.Bytes()...)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:654
							// _ = "end of CoverTab[110997]"
						default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:655
							_go_fuzz_dep_.CoverTab[110998]++
																		dv.Set(sv)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:656
							// _ = "end of CoverTab[110998]"
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:657
						// _ = "end of CoverTab[110986]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:658
						_go_fuzz_dep_.CoverTab[110999]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:658
						// _ = "end of CoverTab[110999]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:658
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:658
					// _ = "end of CoverTab[110982]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:659
				// _ = "end of CoverTab[110980]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:660
			// _ = "end of CoverTab[110784]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:661
			_go_fuzz_dep_.CoverTab[110785]++
														panic(fmt.Sprintf("merger not found for type:%s", tf))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:662
			// _ = "end of CoverTab[110785]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:663
		// _ = "end of CoverTab[110753]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:663
		_go_fuzz_dep_.CoverTab[110754]++
													mi.fields = append(mi.fields, mfi)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:664
		// _ = "end of CoverTab[110754]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:665
	// _ = "end of CoverTab[110743]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:665
	_go_fuzz_dep_.CoverTab[110744]++

												mi.unrecognized = invalidField
												if f, ok := t.FieldByName("XXX_unrecognized"); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:668
		_go_fuzz_dep_.CoverTab[111000]++
													if f.Type != reflect.TypeOf([]byte{}) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:669
			_go_fuzz_dep_.CoverTab[111002]++
														panic("expected XXX_unrecognized to be of type []byte")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:670
			// _ = "end of CoverTab[111002]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:671
			_go_fuzz_dep_.CoverTab[111003]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:671
			// _ = "end of CoverTab[111003]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:671
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:671
		// _ = "end of CoverTab[111000]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:671
		_go_fuzz_dep_.CoverTab[111001]++
													mi.unrecognized = toField(&f)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:672
		// _ = "end of CoverTab[111001]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:673
		_go_fuzz_dep_.CoverTab[111004]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:673
		// _ = "end of CoverTab[111004]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:673
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:673
	// _ = "end of CoverTab[110744]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:673
	_go_fuzz_dep_.CoverTab[110745]++

												atomic.StoreInt32(&mi.initialized, 1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:675
	// _ = "end of CoverTab[110745]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:676
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_merge.go:676
var _ = _go_fuzz_dep_.CoverTab
