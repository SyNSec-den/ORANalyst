// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2018, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:29
)

import (
	"reflect"
	"time"
)

// makeMessageRefMarshaler differs a bit from makeMessageMarshaler
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:36
// It marshal a message T instead of a *T
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:38
func makeMessageRefMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:38
	_go_fuzz_dep_.CoverTab[110558]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:39
			_go_fuzz_dep_.CoverTab[110559]++
															siz := u.size(ptr)
															return siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:41
			// _ = "end of CoverTab[110559]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:43
			_go_fuzz_dep_.CoverTab[110560]++
															b = appendVarint(b, wiretag)
															siz := u.cachedsize(ptr)
															b = appendVarint(b, uint64(siz))
															return u.marshal(b, ptr, deterministic)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:47
			// _ = "end of CoverTab[110560]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:48
	// _ = "end of CoverTab[110558]"
}

// makeMessageRefSliceMarshaler differs quite a lot from makeMessageSliceMarshaler
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:51
// It marshals a slice of messages []T instead of []*T
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:53
func makeMessageRefSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:53
	_go_fuzz_dep_.CoverTab[110561]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:54
			_go_fuzz_dep_.CoverTab[110562]++
															s := ptr.getSlice(u.typ)
															n := 0
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:57
				_go_fuzz_dep_.CoverTab[110564]++
																elem := s.Index(i)
																e := elem.Interface()
																v := toAddrPointer(&e, false)
																siz := u.size(v)
																n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:62
				// _ = "end of CoverTab[110564]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:63
			// _ = "end of CoverTab[110562]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:63
			_go_fuzz_dep_.CoverTab[110563]++
															return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:64
			// _ = "end of CoverTab[110563]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:66
			_go_fuzz_dep_.CoverTab[110565]++
															s := ptr.getSlice(u.typ)
															var err, errreq error
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:69
				_go_fuzz_dep_.CoverTab[110567]++
																elem := s.Index(i)
																e := elem.Interface()
																v := toAddrPointer(&e, false)
																b = appendVarint(b, wiretag)
																siz := u.size(v)
																b = appendVarint(b, uint64(siz))
																b, err = u.marshal(b, v, deterministic)

																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:78
					_go_fuzz_dep_.CoverTab[110568]++
																	if _, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:79
						_go_fuzz_dep_.CoverTab[110571]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:82
						if errreq == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:82
							_go_fuzz_dep_.CoverTab[110573]++
																			errreq = err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:83
							// _ = "end of CoverTab[110573]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:84
							_go_fuzz_dep_.CoverTab[110574]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:84
							// _ = "end of CoverTab[110574]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:84
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:84
						// _ = "end of CoverTab[110571]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:84
						_go_fuzz_dep_.CoverTab[110572]++
																		continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:85
						// _ = "end of CoverTab[110572]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:86
						_go_fuzz_dep_.CoverTab[110575]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:86
						// _ = "end of CoverTab[110575]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:86
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:86
					// _ = "end of CoverTab[110568]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:86
					_go_fuzz_dep_.CoverTab[110569]++
																	if err == ErrNil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:87
						_go_fuzz_dep_.CoverTab[110576]++
																		err = errRepeatedHasNil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:88
						// _ = "end of CoverTab[110576]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:89
						_go_fuzz_dep_.CoverTab[110577]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:89
						// _ = "end of CoverTab[110577]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:89
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:89
					// _ = "end of CoverTab[110569]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:89
					_go_fuzz_dep_.CoverTab[110570]++
																	return b, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:90
					// _ = "end of CoverTab[110570]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:91
					_go_fuzz_dep_.CoverTab[110578]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:91
					// _ = "end of CoverTab[110578]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:91
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:91
				// _ = "end of CoverTab[110567]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:92
			// _ = "end of CoverTab[110565]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:92
			_go_fuzz_dep_.CoverTab[110566]++

															return b, errreq
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:94
			// _ = "end of CoverTab[110566]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:95
	// _ = "end of CoverTab[110561]"
}

func makeCustomPtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:98
	_go_fuzz_dep_.CoverTab[110579]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:99
			_go_fuzz_dep_.CoverTab[110580]++
															if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:100
				_go_fuzz_dep_.CoverTab[110582]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:101
				// _ = "end of CoverTab[110582]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:102
				_go_fuzz_dep_.CoverTab[110583]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:102
				// _ = "end of CoverTab[110583]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:102
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:102
			// _ = "end of CoverTab[110580]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:102
			_go_fuzz_dep_.CoverTab[110581]++
															m := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(custom)
															siz := m.Size()
															return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:105
			// _ = "end of CoverTab[110581]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:106
			_go_fuzz_dep_.CoverTab[110584]++
															if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:107
				_go_fuzz_dep_.CoverTab[110587]++
																return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:108
				// _ = "end of CoverTab[110587]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:109
				_go_fuzz_dep_.CoverTab[110588]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:109
				// _ = "end of CoverTab[110588]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:109
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:109
			// _ = "end of CoverTab[110584]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:109
			_go_fuzz_dep_.CoverTab[110585]++
															m := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(custom)
															siz := m.Size()
															buf, err := m.Marshal()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:113
				_go_fuzz_dep_.CoverTab[110589]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:114
				// _ = "end of CoverTab[110589]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:115
				_go_fuzz_dep_.CoverTab[110590]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:115
				// _ = "end of CoverTab[110590]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:115
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:115
			// _ = "end of CoverTab[110585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:115
			_go_fuzz_dep_.CoverTab[110586]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:119
			// _ = "end of CoverTab[110586]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:120
	// _ = "end of CoverTab[110579]"
}

func makeCustomMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:123
	_go_fuzz_dep_.CoverTab[110591]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:124
			_go_fuzz_dep_.CoverTab[110592]++
															m := ptr.asPointerTo(u.typ).Interface().(custom)
															siz := m.Size()
															return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:127
			// _ = "end of CoverTab[110592]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:128
			_go_fuzz_dep_.CoverTab[110593]++
															m := ptr.asPointerTo(u.typ).Interface().(custom)
															siz := m.Size()
															buf, err := m.Marshal()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:132
				_go_fuzz_dep_.CoverTab[110595]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:133
				// _ = "end of CoverTab[110595]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:134
				_go_fuzz_dep_.CoverTab[110596]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:134
				// _ = "end of CoverTab[110596]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:134
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:134
			// _ = "end of CoverTab[110593]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:134
			_go_fuzz_dep_.CoverTab[110594]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:138
			// _ = "end of CoverTab[110594]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:139
	// _ = "end of CoverTab[110591]"
}

func makeTimeMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:142
	_go_fuzz_dep_.CoverTab[110597]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:143
			_go_fuzz_dep_.CoverTab[110598]++
															t := ptr.asPointerTo(u.typ).Interface().(*time.Time)
															ts, err := timestampProto(*t)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:146
				_go_fuzz_dep_.CoverTab[110600]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:147
				// _ = "end of CoverTab[110600]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:148
				_go_fuzz_dep_.CoverTab[110601]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:148
				// _ = "end of CoverTab[110601]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:148
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:148
			// _ = "end of CoverTab[110598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:148
			_go_fuzz_dep_.CoverTab[110599]++
															siz := Size(ts)
															return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:150
			// _ = "end of CoverTab[110599]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:151
			_go_fuzz_dep_.CoverTab[110602]++
															t := ptr.asPointerTo(u.typ).Interface().(*time.Time)
															ts, err := timestampProto(*t)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:154
				_go_fuzz_dep_.CoverTab[110605]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:155
				// _ = "end of CoverTab[110605]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:156
				_go_fuzz_dep_.CoverTab[110606]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:156
				// _ = "end of CoverTab[110606]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:156
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:156
			// _ = "end of CoverTab[110602]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:156
			_go_fuzz_dep_.CoverTab[110603]++
															buf, err := Marshal(ts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:158
				_go_fuzz_dep_.CoverTab[110607]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:159
				// _ = "end of CoverTab[110607]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:160
				_go_fuzz_dep_.CoverTab[110608]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:160
				// _ = "end of CoverTab[110608]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:160
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:160
			// _ = "end of CoverTab[110603]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:160
			_go_fuzz_dep_.CoverTab[110604]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(len(buf)))
															b = append(b, buf...)
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:164
			// _ = "end of CoverTab[110604]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:165
	// _ = "end of CoverTab[110597]"
}

func makeTimePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:168
	_go_fuzz_dep_.CoverTab[110609]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:169
			_go_fuzz_dep_.CoverTab[110610]++
															if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:170
				_go_fuzz_dep_.CoverTab[110613]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:171
				// _ = "end of CoverTab[110613]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:172
				_go_fuzz_dep_.CoverTab[110614]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:172
				// _ = "end of CoverTab[110614]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:172
			// _ = "end of CoverTab[110610]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:172
			_go_fuzz_dep_.CoverTab[110611]++
															t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*time.Time)
															ts, err := timestampProto(*t)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:175
				_go_fuzz_dep_.CoverTab[110615]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:176
				// _ = "end of CoverTab[110615]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:177
				_go_fuzz_dep_.CoverTab[110616]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:177
				// _ = "end of CoverTab[110616]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:177
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:177
			// _ = "end of CoverTab[110611]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:177
			_go_fuzz_dep_.CoverTab[110612]++
															siz := Size(ts)
															return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:179
			// _ = "end of CoverTab[110612]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:180
			_go_fuzz_dep_.CoverTab[110617]++
															if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:181
				_go_fuzz_dep_.CoverTab[110621]++
																return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:182
				// _ = "end of CoverTab[110621]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:183
				_go_fuzz_dep_.CoverTab[110622]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:183
				// _ = "end of CoverTab[110622]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:183
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:183
			// _ = "end of CoverTab[110617]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:183
			_go_fuzz_dep_.CoverTab[110618]++
															t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*time.Time)
															ts, err := timestampProto(*t)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:186
				_go_fuzz_dep_.CoverTab[110623]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:187
				// _ = "end of CoverTab[110623]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:188
				_go_fuzz_dep_.CoverTab[110624]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:188
				// _ = "end of CoverTab[110624]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:188
			// _ = "end of CoverTab[110618]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:188
			_go_fuzz_dep_.CoverTab[110619]++
															buf, err := Marshal(ts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:190
				_go_fuzz_dep_.CoverTab[110625]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:191
				// _ = "end of CoverTab[110625]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:192
				_go_fuzz_dep_.CoverTab[110626]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:192
				// _ = "end of CoverTab[110626]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:192
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:192
			// _ = "end of CoverTab[110619]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:192
			_go_fuzz_dep_.CoverTab[110620]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(len(buf)))
															b = append(b, buf...)
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:196
			// _ = "end of CoverTab[110620]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:197
	// _ = "end of CoverTab[110609]"
}

func makeTimeSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:200
	_go_fuzz_dep_.CoverTab[110627]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:201
			_go_fuzz_dep_.CoverTab[110628]++
															s := ptr.getSlice(u.typ)
															n := 0
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:204
				_go_fuzz_dep_.CoverTab[110630]++
																elem := s.Index(i)
																t := elem.Interface().(time.Time)
																ts, err := timestampProto(t)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:208
					_go_fuzz_dep_.CoverTab[110632]++
																	return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:209
					// _ = "end of CoverTab[110632]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:210
					_go_fuzz_dep_.CoverTab[110633]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:210
					// _ = "end of CoverTab[110633]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:210
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:210
				// _ = "end of CoverTab[110630]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:210
				_go_fuzz_dep_.CoverTab[110631]++
																siz := Size(ts)
																n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:212
				// _ = "end of CoverTab[110631]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:213
			// _ = "end of CoverTab[110628]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:213
			_go_fuzz_dep_.CoverTab[110629]++
															return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:214
			// _ = "end of CoverTab[110629]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:216
			_go_fuzz_dep_.CoverTab[110634]++
															s := ptr.getSlice(u.typ)
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:218
				_go_fuzz_dep_.CoverTab[110636]++
																elem := s.Index(i)
																t := elem.Interface().(time.Time)
																ts, err := timestampProto(t)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:222
					_go_fuzz_dep_.CoverTab[110639]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:223
					// _ = "end of CoverTab[110639]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:224
					_go_fuzz_dep_.CoverTab[110640]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:224
					// _ = "end of CoverTab[110640]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:224
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:224
				// _ = "end of CoverTab[110636]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:224
				_go_fuzz_dep_.CoverTab[110637]++
																siz := Size(ts)
																buf, err := Marshal(ts)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:227
					_go_fuzz_dep_.CoverTab[110641]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:228
					// _ = "end of CoverTab[110641]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:229
					_go_fuzz_dep_.CoverTab[110642]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:229
					// _ = "end of CoverTab[110642]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:229
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:229
				// _ = "end of CoverTab[110637]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:229
				_go_fuzz_dep_.CoverTab[110638]++
																b = appendVarint(b, wiretag)
																b = appendVarint(b, uint64(siz))
																b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:232
				// _ = "end of CoverTab[110638]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:233
			// _ = "end of CoverTab[110634]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:233
			_go_fuzz_dep_.CoverTab[110635]++

															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:235
			// _ = "end of CoverTab[110635]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:236
	// _ = "end of CoverTab[110627]"
}

func makeTimePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:239
	_go_fuzz_dep_.CoverTab[110643]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:240
			_go_fuzz_dep_.CoverTab[110644]++
															s := ptr.getSlice(reflect.PtrTo(u.typ))
															n := 0
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:243
				_go_fuzz_dep_.CoverTab[110646]++
																elem := s.Index(i)
																t := elem.Interface().(*time.Time)
																ts, err := timestampProto(*t)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:247
					_go_fuzz_dep_.CoverTab[110648]++
																	return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:248
					// _ = "end of CoverTab[110648]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:249
					_go_fuzz_dep_.CoverTab[110649]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:249
					// _ = "end of CoverTab[110649]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:249
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:249
				// _ = "end of CoverTab[110646]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:249
				_go_fuzz_dep_.CoverTab[110647]++
																siz := Size(ts)
																n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:251
				// _ = "end of CoverTab[110647]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:252
			// _ = "end of CoverTab[110644]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:252
			_go_fuzz_dep_.CoverTab[110645]++
															return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:253
			// _ = "end of CoverTab[110645]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:255
			_go_fuzz_dep_.CoverTab[110650]++
															s := ptr.getSlice(reflect.PtrTo(u.typ))
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:257
				_go_fuzz_dep_.CoverTab[110652]++
																elem := s.Index(i)
																t := elem.Interface().(*time.Time)
																ts, err := timestampProto(*t)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:261
					_go_fuzz_dep_.CoverTab[110655]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:262
					// _ = "end of CoverTab[110655]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:263
					_go_fuzz_dep_.CoverTab[110656]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:263
					// _ = "end of CoverTab[110656]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:263
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:263
				// _ = "end of CoverTab[110652]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:263
				_go_fuzz_dep_.CoverTab[110653]++
																siz := Size(ts)
																buf, err := Marshal(ts)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:266
					_go_fuzz_dep_.CoverTab[110657]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:267
					// _ = "end of CoverTab[110657]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:268
					_go_fuzz_dep_.CoverTab[110658]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:268
					// _ = "end of CoverTab[110658]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:268
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:268
				// _ = "end of CoverTab[110653]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:268
				_go_fuzz_dep_.CoverTab[110654]++
																b = appendVarint(b, wiretag)
																b = appendVarint(b, uint64(siz))
																b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:271
				// _ = "end of CoverTab[110654]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:272
			// _ = "end of CoverTab[110650]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:272
			_go_fuzz_dep_.CoverTab[110651]++

															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:274
			// _ = "end of CoverTab[110651]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:275
	// _ = "end of CoverTab[110643]"
}

func makeDurationMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:278
	_go_fuzz_dep_.CoverTab[110659]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:279
			_go_fuzz_dep_.CoverTab[110660]++
															d := ptr.asPointerTo(u.typ).Interface().(*time.Duration)
															dur := durationProto(*d)
															siz := Size(dur)
															return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:283
			// _ = "end of CoverTab[110660]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:284
			_go_fuzz_dep_.CoverTab[110661]++
															d := ptr.asPointerTo(u.typ).Interface().(*time.Duration)
															dur := durationProto(*d)
															buf, err := Marshal(dur)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:288
				_go_fuzz_dep_.CoverTab[110663]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:289
				// _ = "end of CoverTab[110663]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:290
				_go_fuzz_dep_.CoverTab[110664]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:290
				// _ = "end of CoverTab[110664]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:290
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:290
			// _ = "end of CoverTab[110661]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:290
			_go_fuzz_dep_.CoverTab[110662]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(len(buf)))
															b = append(b, buf...)
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:294
			// _ = "end of CoverTab[110662]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:295
	// _ = "end of CoverTab[110659]"
}

func makeDurationPtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:298
	_go_fuzz_dep_.CoverTab[110665]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:299
			_go_fuzz_dep_.CoverTab[110666]++
															if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:300
				_go_fuzz_dep_.CoverTab[110668]++
																return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:301
				// _ = "end of CoverTab[110668]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:302
				_go_fuzz_dep_.CoverTab[110669]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:302
				// _ = "end of CoverTab[110669]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:302
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:302
			// _ = "end of CoverTab[110666]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:302
			_go_fuzz_dep_.CoverTab[110667]++
															d := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*time.Duration)
															dur := durationProto(*d)
															siz := Size(dur)
															return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:306
			// _ = "end of CoverTab[110667]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:307
			_go_fuzz_dep_.CoverTab[110670]++
															if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:308
				_go_fuzz_dep_.CoverTab[110673]++
																return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:309
				// _ = "end of CoverTab[110673]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:310
				_go_fuzz_dep_.CoverTab[110674]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:310
				// _ = "end of CoverTab[110674]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:310
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:310
			// _ = "end of CoverTab[110670]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:310
			_go_fuzz_dep_.CoverTab[110671]++
															d := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*time.Duration)
															dur := durationProto(*d)
															buf, err := Marshal(dur)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:314
				_go_fuzz_dep_.CoverTab[110675]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:315
				// _ = "end of CoverTab[110675]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:316
				_go_fuzz_dep_.CoverTab[110676]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:316
				// _ = "end of CoverTab[110676]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:316
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:316
			// _ = "end of CoverTab[110671]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:316
			_go_fuzz_dep_.CoverTab[110672]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(len(buf)))
															b = append(b, buf...)
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:320
			// _ = "end of CoverTab[110672]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:321
	// _ = "end of CoverTab[110665]"
}

func makeDurationSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:324
	_go_fuzz_dep_.CoverTab[110677]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:325
			_go_fuzz_dep_.CoverTab[110678]++
															s := ptr.getSlice(u.typ)
															n := 0
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:328
				_go_fuzz_dep_.CoverTab[110680]++
																elem := s.Index(i)
																d := elem.Interface().(time.Duration)
																dur := durationProto(d)
																siz := Size(dur)
																n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:333
				// _ = "end of CoverTab[110680]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:334
			// _ = "end of CoverTab[110678]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:334
			_go_fuzz_dep_.CoverTab[110679]++
															return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:335
			// _ = "end of CoverTab[110679]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:337
			_go_fuzz_dep_.CoverTab[110681]++
															s := ptr.getSlice(u.typ)
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:339
				_go_fuzz_dep_.CoverTab[110683]++
																elem := s.Index(i)
																d := elem.Interface().(time.Duration)
																dur := durationProto(d)
																siz := Size(dur)
																buf, err := Marshal(dur)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:345
					_go_fuzz_dep_.CoverTab[110685]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:346
					// _ = "end of CoverTab[110685]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:347
					_go_fuzz_dep_.CoverTab[110686]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:347
					// _ = "end of CoverTab[110686]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:347
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:347
				// _ = "end of CoverTab[110683]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:347
				_go_fuzz_dep_.CoverTab[110684]++
																b = appendVarint(b, wiretag)
																b = appendVarint(b, uint64(siz))
																b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:350
				// _ = "end of CoverTab[110684]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:351
			// _ = "end of CoverTab[110681]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:351
			_go_fuzz_dep_.CoverTab[110682]++

															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:353
			// _ = "end of CoverTab[110682]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:354
	// _ = "end of CoverTab[110677]"
}

func makeDurationPtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:357
	_go_fuzz_dep_.CoverTab[110687]++
													return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:358
			_go_fuzz_dep_.CoverTab[110688]++
															s := ptr.getSlice(reflect.PtrTo(u.typ))
															n := 0
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:361
				_go_fuzz_dep_.CoverTab[110690]++
																elem := s.Index(i)
																d := elem.Interface().(*time.Duration)
																dur := durationProto(*d)
																siz := Size(dur)
																n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:366
				// _ = "end of CoverTab[110690]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:367
			// _ = "end of CoverTab[110688]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:367
			_go_fuzz_dep_.CoverTab[110689]++
															return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:368
			// _ = "end of CoverTab[110689]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:370
			_go_fuzz_dep_.CoverTab[110691]++
															s := ptr.getSlice(reflect.PtrTo(u.typ))
															for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:372
				_go_fuzz_dep_.CoverTab[110693]++
																elem := s.Index(i)
																d := elem.Interface().(*time.Duration)
																dur := durationProto(*d)
																siz := Size(dur)
																buf, err := Marshal(dur)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:378
					_go_fuzz_dep_.CoverTab[110695]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:379
					// _ = "end of CoverTab[110695]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:380
					_go_fuzz_dep_.CoverTab[110696]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:380
					// _ = "end of CoverTab[110696]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:380
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:380
				// _ = "end of CoverTab[110693]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:380
				_go_fuzz_dep_.CoverTab[110694]++
																b = appendVarint(b, wiretag)
																b = appendVarint(b, uint64(siz))
																b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:383
				// _ = "end of CoverTab[110694]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:384
			// _ = "end of CoverTab[110691]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:384
			_go_fuzz_dep_.CoverTab[110692]++

															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:386
			// _ = "end of CoverTab[110692]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:387
	// _ = "end of CoverTab[110687]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:388
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_marshal_gogo.go:388
var _ = _go_fuzz_dep_.CoverTab
