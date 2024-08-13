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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:29
)

import (
	"io"
	"reflect"
)

func makeStdDoubleValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:36
	_go_fuzz_dep_.CoverTab[113556]++
											return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:37
			_go_fuzz_dep_.CoverTab[113557]++
													t := ptr.asPointerTo(u.typ).Interface().(*float64)
													v := &float64Value{*t}
													siz := Size(v)
													return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:41
			// _ = "end of CoverTab[113557]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:42
			_go_fuzz_dep_.CoverTab[113558]++
													t := ptr.asPointerTo(u.typ).Interface().(*float64)
													v := &float64Value{*t}
													buf, err := Marshal(v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:46
				_go_fuzz_dep_.CoverTab[113560]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:47
				// _ = "end of CoverTab[113560]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:48
				_go_fuzz_dep_.CoverTab[113561]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:48
				// _ = "end of CoverTab[113561]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:48
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:48
			// _ = "end of CoverTab[113558]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:48
			_go_fuzz_dep_.CoverTab[113559]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(len(buf)))
													b = append(b, buf...)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:52
			// _ = "end of CoverTab[113559]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:53
	// _ = "end of CoverTab[113556]"
}

func makeStdDoubleValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:56
	_go_fuzz_dep_.CoverTab[113562]++
											return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:57
			_go_fuzz_dep_.CoverTab[113563]++
													if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:58
				_go_fuzz_dep_.CoverTab[113565]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:59
				// _ = "end of CoverTab[113565]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:60
				_go_fuzz_dep_.CoverTab[113566]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:60
				// _ = "end of CoverTab[113566]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:60
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:60
			// _ = "end of CoverTab[113563]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:60
			_go_fuzz_dep_.CoverTab[113564]++
													t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*float64)
													v := &float64Value{*t}
													siz := Size(v)
													return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:64
			// _ = "end of CoverTab[113564]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:65
			_go_fuzz_dep_.CoverTab[113567]++
													if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:66
				_go_fuzz_dep_.CoverTab[113570]++
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:67
				// _ = "end of CoverTab[113570]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:68
				_go_fuzz_dep_.CoverTab[113571]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:68
				// _ = "end of CoverTab[113571]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:68
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:68
			// _ = "end of CoverTab[113567]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:68
			_go_fuzz_dep_.CoverTab[113568]++
													t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*float64)
													v := &float64Value{*t}
													buf, err := Marshal(v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:72
				_go_fuzz_dep_.CoverTab[113572]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:73
				// _ = "end of CoverTab[113572]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:74
				_go_fuzz_dep_.CoverTab[113573]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:74
				// _ = "end of CoverTab[113573]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:74
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:74
			// _ = "end of CoverTab[113568]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:74
			_go_fuzz_dep_.CoverTab[113569]++
													b = appendVarint(b, wiretag)
													b = appendVarint(b, uint64(len(buf)))
													b = append(b, buf...)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:78
			// _ = "end of CoverTab[113569]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:79
	// _ = "end of CoverTab[113562]"
}

func makeStdDoubleValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:82
	_go_fuzz_dep_.CoverTab[113574]++
											return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:83
			_go_fuzz_dep_.CoverTab[113575]++
													s := ptr.getSlice(u.typ)
													n := 0
													for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:86
				_go_fuzz_dep_.CoverTab[113577]++
														elem := s.Index(i)
														t := elem.Interface().(float64)
														v := &float64Value{t}
														siz := Size(v)
														n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:91
				// _ = "end of CoverTab[113577]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:92
			// _ = "end of CoverTab[113575]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:92
			_go_fuzz_dep_.CoverTab[113576]++
													return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:93
			// _ = "end of CoverTab[113576]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:95
			_go_fuzz_dep_.CoverTab[113578]++
													s := ptr.getSlice(u.typ)
													for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:97
				_go_fuzz_dep_.CoverTab[113580]++
														elem := s.Index(i)
														t := elem.Interface().(float64)
														v := &float64Value{t}
														siz := Size(v)
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:103
					_go_fuzz_dep_.CoverTab[113582]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:104
					// _ = "end of CoverTab[113582]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:105
					_go_fuzz_dep_.CoverTab[113583]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:105
					// _ = "end of CoverTab[113583]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:105
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:105
				// _ = "end of CoverTab[113580]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:105
				_go_fuzz_dep_.CoverTab[113581]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:108
				// _ = "end of CoverTab[113581]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:109
			// _ = "end of CoverTab[113578]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:109
			_go_fuzz_dep_.CoverTab[113579]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:111
			// _ = "end of CoverTab[113579]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:112
	// _ = "end of CoverTab[113574]"
}

func makeStdDoubleValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:115
	_go_fuzz_dep_.CoverTab[113584]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:116
			_go_fuzz_dep_.CoverTab[113585]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:119
				_go_fuzz_dep_.CoverTab[113587]++
															elem := s.Index(i)
															t := elem.Interface().(*float64)
															v := &float64Value{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:124
				// _ = "end of CoverTab[113587]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:125
			// _ = "end of CoverTab[113585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:125
			_go_fuzz_dep_.CoverTab[113586]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:126
			// _ = "end of CoverTab[113586]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:128
			_go_fuzz_dep_.CoverTab[113588]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:130
				_go_fuzz_dep_.CoverTab[113590]++
															elem := s.Index(i)
															t := elem.Interface().(*float64)
															v := &float64Value{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:136
					_go_fuzz_dep_.CoverTab[113592]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:137
					// _ = "end of CoverTab[113592]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:138
					_go_fuzz_dep_.CoverTab[113593]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:138
					// _ = "end of CoverTab[113593]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:138
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:138
				// _ = "end of CoverTab[113590]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:138
				_go_fuzz_dep_.CoverTab[113591]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:141
				// _ = "end of CoverTab[113591]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:142
			// _ = "end of CoverTab[113588]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:142
			_go_fuzz_dep_.CoverTab[113589]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:144
			// _ = "end of CoverTab[113589]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:145
	// _ = "end of CoverTab[113584]"
}

func makeStdDoubleValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:148
	_go_fuzz_dep_.CoverTab[113594]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:149
		_go_fuzz_dep_.CoverTab[113595]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:150
			_go_fuzz_dep_.CoverTab[113600]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:151
			// _ = "end of CoverTab[113600]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:152
			_go_fuzz_dep_.CoverTab[113601]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:152
			// _ = "end of CoverTab[113601]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:152
		// _ = "end of CoverTab[113595]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:152
		_go_fuzz_dep_.CoverTab[113596]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:154
			_go_fuzz_dep_.CoverTab[113602]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:155
			// _ = "end of CoverTab[113602]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:156
			_go_fuzz_dep_.CoverTab[113603]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:156
			// _ = "end of CoverTab[113603]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:156
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:156
		// _ = "end of CoverTab[113596]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:156
		_go_fuzz_dep_.CoverTab[113597]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:158
			_go_fuzz_dep_.CoverTab[113604]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:159
			// _ = "end of CoverTab[113604]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:160
			_go_fuzz_dep_.CoverTab[113605]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:160
			// _ = "end of CoverTab[113605]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:160
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:160
		// _ = "end of CoverTab[113597]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:160
		_go_fuzz_dep_.CoverTab[113598]++
													m := &float64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:162
			_go_fuzz_dep_.CoverTab[113606]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:163
			// _ = "end of CoverTab[113606]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:164
			_go_fuzz_dep_.CoverTab[113607]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:164
			// _ = "end of CoverTab[113607]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:164
		// _ = "end of CoverTab[113598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:164
		_go_fuzz_dep_.CoverTab[113599]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:167
		// _ = "end of CoverTab[113599]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:168
	// _ = "end of CoverTab[113594]"
}

func makeStdDoubleValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:171
	_go_fuzz_dep_.CoverTab[113608]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:172
		_go_fuzz_dep_.CoverTab[113609]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:173
			_go_fuzz_dep_.CoverTab[113614]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:174
			// _ = "end of CoverTab[113614]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:175
			_go_fuzz_dep_.CoverTab[113615]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:175
			// _ = "end of CoverTab[113615]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:175
		// _ = "end of CoverTab[113609]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:175
		_go_fuzz_dep_.CoverTab[113610]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:177
			_go_fuzz_dep_.CoverTab[113616]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:178
			// _ = "end of CoverTab[113616]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:179
			_go_fuzz_dep_.CoverTab[113617]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:179
			// _ = "end of CoverTab[113617]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:179
		// _ = "end of CoverTab[113610]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:179
		_go_fuzz_dep_.CoverTab[113611]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:181
			_go_fuzz_dep_.CoverTab[113618]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:182
			// _ = "end of CoverTab[113618]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:183
			_go_fuzz_dep_.CoverTab[113619]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:183
			// _ = "end of CoverTab[113619]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:183
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:183
		// _ = "end of CoverTab[113611]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:183
		_go_fuzz_dep_.CoverTab[113612]++
													m := &float64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:185
			_go_fuzz_dep_.CoverTab[113620]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:186
			// _ = "end of CoverTab[113620]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:187
			_go_fuzz_dep_.CoverTab[113621]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:187
			// _ = "end of CoverTab[113621]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:187
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:187
		// _ = "end of CoverTab[113612]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:187
		_go_fuzz_dep_.CoverTab[113613]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:190
		// _ = "end of CoverTab[113613]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:191
	// _ = "end of CoverTab[113608]"
}

func makeStdDoubleValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:194
	_go_fuzz_dep_.CoverTab[113622]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:195
		_go_fuzz_dep_.CoverTab[113623]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:196
			_go_fuzz_dep_.CoverTab[113628]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:197
			// _ = "end of CoverTab[113628]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:198
			_go_fuzz_dep_.CoverTab[113629]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:198
			// _ = "end of CoverTab[113629]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:198
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:198
		// _ = "end of CoverTab[113623]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:198
		_go_fuzz_dep_.CoverTab[113624]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:200
			_go_fuzz_dep_.CoverTab[113630]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:201
			// _ = "end of CoverTab[113630]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:202
			_go_fuzz_dep_.CoverTab[113631]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:202
			// _ = "end of CoverTab[113631]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:202
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:202
		// _ = "end of CoverTab[113624]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:202
		_go_fuzz_dep_.CoverTab[113625]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:204
			_go_fuzz_dep_.CoverTab[113632]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:205
			// _ = "end of CoverTab[113632]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:206
			_go_fuzz_dep_.CoverTab[113633]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:206
			// _ = "end of CoverTab[113633]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:206
		// _ = "end of CoverTab[113625]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:206
		_go_fuzz_dep_.CoverTab[113626]++
													m := &float64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:208
			_go_fuzz_dep_.CoverTab[113634]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:209
			// _ = "end of CoverTab[113634]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:210
			_go_fuzz_dep_.CoverTab[113635]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:210
			// _ = "end of CoverTab[113635]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:210
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:210
		// _ = "end of CoverTab[113626]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:210
		_go_fuzz_dep_.CoverTab[113627]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:214
		// _ = "end of CoverTab[113627]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:215
	// _ = "end of CoverTab[113622]"
}

func makeStdDoubleValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:218
	_go_fuzz_dep_.CoverTab[113636]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:219
		_go_fuzz_dep_.CoverTab[113637]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:220
			_go_fuzz_dep_.CoverTab[113642]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:221
			// _ = "end of CoverTab[113642]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:222
			_go_fuzz_dep_.CoverTab[113643]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:222
			// _ = "end of CoverTab[113643]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:222
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:222
		// _ = "end of CoverTab[113637]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:222
		_go_fuzz_dep_.CoverTab[113638]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:224
			_go_fuzz_dep_.CoverTab[113644]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:225
			// _ = "end of CoverTab[113644]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:226
			_go_fuzz_dep_.CoverTab[113645]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:226
			// _ = "end of CoverTab[113645]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:226
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:226
		// _ = "end of CoverTab[113638]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:226
		_go_fuzz_dep_.CoverTab[113639]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:228
			_go_fuzz_dep_.CoverTab[113646]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:229
			// _ = "end of CoverTab[113646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:230
			_go_fuzz_dep_.CoverTab[113647]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:230
			// _ = "end of CoverTab[113647]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:230
		// _ = "end of CoverTab[113639]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:230
		_go_fuzz_dep_.CoverTab[113640]++
													m := &float64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:232
			_go_fuzz_dep_.CoverTab[113648]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:233
			// _ = "end of CoverTab[113648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:234
			_go_fuzz_dep_.CoverTab[113649]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:234
			// _ = "end of CoverTab[113649]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:234
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:234
		// _ = "end of CoverTab[113640]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:234
		_go_fuzz_dep_.CoverTab[113641]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:238
		// _ = "end of CoverTab[113641]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:239
	// _ = "end of CoverTab[113636]"
}

func makeStdFloatValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:242
	_go_fuzz_dep_.CoverTab[113650]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:243
			_go_fuzz_dep_.CoverTab[113651]++
														t := ptr.asPointerTo(u.typ).Interface().(*float32)
														v := &float32Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:247
			// _ = "end of CoverTab[113651]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:248
			_go_fuzz_dep_.CoverTab[113652]++
														t := ptr.asPointerTo(u.typ).Interface().(*float32)
														v := &float32Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:252
				_go_fuzz_dep_.CoverTab[113654]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:253
				// _ = "end of CoverTab[113654]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:254
				_go_fuzz_dep_.CoverTab[113655]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:254
				// _ = "end of CoverTab[113655]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:254
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:254
			// _ = "end of CoverTab[113652]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:254
			_go_fuzz_dep_.CoverTab[113653]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:258
			// _ = "end of CoverTab[113653]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:259
	// _ = "end of CoverTab[113650]"
}

func makeStdFloatValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:262
	_go_fuzz_dep_.CoverTab[113656]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:263
			_go_fuzz_dep_.CoverTab[113657]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:264
				_go_fuzz_dep_.CoverTab[113659]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:265
				// _ = "end of CoverTab[113659]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:266
				_go_fuzz_dep_.CoverTab[113660]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:266
				// _ = "end of CoverTab[113660]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:266
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:266
			// _ = "end of CoverTab[113657]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:266
			_go_fuzz_dep_.CoverTab[113658]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*float32)
														v := &float32Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:270
			// _ = "end of CoverTab[113658]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:271
			_go_fuzz_dep_.CoverTab[113661]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:272
				_go_fuzz_dep_.CoverTab[113664]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:273
				// _ = "end of CoverTab[113664]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:274
				_go_fuzz_dep_.CoverTab[113665]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:274
				// _ = "end of CoverTab[113665]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:274
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:274
			// _ = "end of CoverTab[113661]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:274
			_go_fuzz_dep_.CoverTab[113662]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*float32)
														v := &float32Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:278
				_go_fuzz_dep_.CoverTab[113666]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:279
				// _ = "end of CoverTab[113666]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:280
				_go_fuzz_dep_.CoverTab[113667]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:280
				// _ = "end of CoverTab[113667]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:280
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:280
			// _ = "end of CoverTab[113662]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:280
			_go_fuzz_dep_.CoverTab[113663]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:284
			// _ = "end of CoverTab[113663]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:285
	// _ = "end of CoverTab[113656]"
}

func makeStdFloatValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:288
	_go_fuzz_dep_.CoverTab[113668]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:289
			_go_fuzz_dep_.CoverTab[113669]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:292
				_go_fuzz_dep_.CoverTab[113671]++
															elem := s.Index(i)
															t := elem.Interface().(float32)
															v := &float32Value{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:297
				// _ = "end of CoverTab[113671]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:298
			// _ = "end of CoverTab[113669]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:298
			_go_fuzz_dep_.CoverTab[113670]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:299
			// _ = "end of CoverTab[113670]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:301
			_go_fuzz_dep_.CoverTab[113672]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:303
				_go_fuzz_dep_.CoverTab[113674]++
															elem := s.Index(i)
															t := elem.Interface().(float32)
															v := &float32Value{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:309
					_go_fuzz_dep_.CoverTab[113676]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:310
					// _ = "end of CoverTab[113676]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:311
					_go_fuzz_dep_.CoverTab[113677]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:311
					// _ = "end of CoverTab[113677]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:311
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:311
				// _ = "end of CoverTab[113674]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:311
				_go_fuzz_dep_.CoverTab[113675]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:314
				// _ = "end of CoverTab[113675]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:315
			// _ = "end of CoverTab[113672]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:315
			_go_fuzz_dep_.CoverTab[113673]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:317
			// _ = "end of CoverTab[113673]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:318
	// _ = "end of CoverTab[113668]"
}

func makeStdFloatValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:321
	_go_fuzz_dep_.CoverTab[113678]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:322
			_go_fuzz_dep_.CoverTab[113679]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:325
				_go_fuzz_dep_.CoverTab[113681]++
															elem := s.Index(i)
															t := elem.Interface().(*float32)
															v := &float32Value{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:330
				// _ = "end of CoverTab[113681]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:331
			// _ = "end of CoverTab[113679]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:331
			_go_fuzz_dep_.CoverTab[113680]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:332
			// _ = "end of CoverTab[113680]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:334
			_go_fuzz_dep_.CoverTab[113682]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:336
				_go_fuzz_dep_.CoverTab[113684]++
															elem := s.Index(i)
															t := elem.Interface().(*float32)
															v := &float32Value{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:342
					_go_fuzz_dep_.CoverTab[113686]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:343
					// _ = "end of CoverTab[113686]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:344
					_go_fuzz_dep_.CoverTab[113687]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:344
					// _ = "end of CoverTab[113687]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:344
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:344
				// _ = "end of CoverTab[113684]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:344
				_go_fuzz_dep_.CoverTab[113685]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:347
				// _ = "end of CoverTab[113685]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:348
			// _ = "end of CoverTab[113682]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:348
			_go_fuzz_dep_.CoverTab[113683]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:350
			// _ = "end of CoverTab[113683]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:351
	// _ = "end of CoverTab[113678]"
}

func makeStdFloatValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:354
	_go_fuzz_dep_.CoverTab[113688]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:355
		_go_fuzz_dep_.CoverTab[113689]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:356
			_go_fuzz_dep_.CoverTab[113694]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:357
			// _ = "end of CoverTab[113694]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:358
			_go_fuzz_dep_.CoverTab[113695]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:358
			// _ = "end of CoverTab[113695]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:358
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:358
		// _ = "end of CoverTab[113689]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:358
		_go_fuzz_dep_.CoverTab[113690]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:360
			_go_fuzz_dep_.CoverTab[113696]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:361
			// _ = "end of CoverTab[113696]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:362
			_go_fuzz_dep_.CoverTab[113697]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:362
			// _ = "end of CoverTab[113697]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:362
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:362
		// _ = "end of CoverTab[113690]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:362
		_go_fuzz_dep_.CoverTab[113691]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:364
			_go_fuzz_dep_.CoverTab[113698]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:365
			// _ = "end of CoverTab[113698]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:366
			_go_fuzz_dep_.CoverTab[113699]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:366
			// _ = "end of CoverTab[113699]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:366
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:366
		// _ = "end of CoverTab[113691]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:366
		_go_fuzz_dep_.CoverTab[113692]++
													m := &float32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:368
			_go_fuzz_dep_.CoverTab[113700]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:369
			// _ = "end of CoverTab[113700]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:370
			_go_fuzz_dep_.CoverTab[113701]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:370
			// _ = "end of CoverTab[113701]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:370
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:370
		// _ = "end of CoverTab[113692]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:370
		_go_fuzz_dep_.CoverTab[113693]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:373
		// _ = "end of CoverTab[113693]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:374
	// _ = "end of CoverTab[113688]"
}

func makeStdFloatValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:377
	_go_fuzz_dep_.CoverTab[113702]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:378
		_go_fuzz_dep_.CoverTab[113703]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:379
			_go_fuzz_dep_.CoverTab[113708]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:380
			// _ = "end of CoverTab[113708]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:381
			_go_fuzz_dep_.CoverTab[113709]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:381
			// _ = "end of CoverTab[113709]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:381
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:381
		// _ = "end of CoverTab[113703]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:381
		_go_fuzz_dep_.CoverTab[113704]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:383
			_go_fuzz_dep_.CoverTab[113710]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:384
			// _ = "end of CoverTab[113710]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:385
			_go_fuzz_dep_.CoverTab[113711]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:385
			// _ = "end of CoverTab[113711]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:385
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:385
		// _ = "end of CoverTab[113704]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:385
		_go_fuzz_dep_.CoverTab[113705]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:387
			_go_fuzz_dep_.CoverTab[113712]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:388
			// _ = "end of CoverTab[113712]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:389
			_go_fuzz_dep_.CoverTab[113713]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:389
			// _ = "end of CoverTab[113713]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:389
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:389
		// _ = "end of CoverTab[113705]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:389
		_go_fuzz_dep_.CoverTab[113706]++
													m := &float32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:391
			_go_fuzz_dep_.CoverTab[113714]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:392
			// _ = "end of CoverTab[113714]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:393
			_go_fuzz_dep_.CoverTab[113715]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:393
			// _ = "end of CoverTab[113715]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:393
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:393
		// _ = "end of CoverTab[113706]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:393
		_go_fuzz_dep_.CoverTab[113707]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:396
		// _ = "end of CoverTab[113707]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:397
	// _ = "end of CoverTab[113702]"
}

func makeStdFloatValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:400
	_go_fuzz_dep_.CoverTab[113716]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:401
		_go_fuzz_dep_.CoverTab[113717]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:402
			_go_fuzz_dep_.CoverTab[113722]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:403
			// _ = "end of CoverTab[113722]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:404
			_go_fuzz_dep_.CoverTab[113723]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:404
			// _ = "end of CoverTab[113723]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:404
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:404
		// _ = "end of CoverTab[113717]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:404
		_go_fuzz_dep_.CoverTab[113718]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:406
			_go_fuzz_dep_.CoverTab[113724]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:407
			// _ = "end of CoverTab[113724]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:408
			_go_fuzz_dep_.CoverTab[113725]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:408
			// _ = "end of CoverTab[113725]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:408
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:408
		// _ = "end of CoverTab[113718]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:408
		_go_fuzz_dep_.CoverTab[113719]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:410
			_go_fuzz_dep_.CoverTab[113726]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:411
			// _ = "end of CoverTab[113726]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:412
			_go_fuzz_dep_.CoverTab[113727]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:412
			// _ = "end of CoverTab[113727]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:412
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:412
		// _ = "end of CoverTab[113719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:412
		_go_fuzz_dep_.CoverTab[113720]++
													m := &float32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:414
			_go_fuzz_dep_.CoverTab[113728]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:415
			// _ = "end of CoverTab[113728]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:416
			_go_fuzz_dep_.CoverTab[113729]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:416
			// _ = "end of CoverTab[113729]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:416
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:416
		// _ = "end of CoverTab[113720]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:416
		_go_fuzz_dep_.CoverTab[113721]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:420
		// _ = "end of CoverTab[113721]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:421
	// _ = "end of CoverTab[113716]"
}

func makeStdFloatValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:424
	_go_fuzz_dep_.CoverTab[113730]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:425
		_go_fuzz_dep_.CoverTab[113731]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:426
			_go_fuzz_dep_.CoverTab[113736]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:427
			// _ = "end of CoverTab[113736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:428
			_go_fuzz_dep_.CoverTab[113737]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:428
			// _ = "end of CoverTab[113737]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:428
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:428
		// _ = "end of CoverTab[113731]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:428
		_go_fuzz_dep_.CoverTab[113732]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:430
			_go_fuzz_dep_.CoverTab[113738]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:431
			// _ = "end of CoverTab[113738]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:432
			_go_fuzz_dep_.CoverTab[113739]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:432
			// _ = "end of CoverTab[113739]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:432
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:432
		// _ = "end of CoverTab[113732]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:432
		_go_fuzz_dep_.CoverTab[113733]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:434
			_go_fuzz_dep_.CoverTab[113740]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:435
			// _ = "end of CoverTab[113740]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:436
			_go_fuzz_dep_.CoverTab[113741]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:436
			// _ = "end of CoverTab[113741]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:436
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:436
		// _ = "end of CoverTab[113733]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:436
		_go_fuzz_dep_.CoverTab[113734]++
													m := &float32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:438
			_go_fuzz_dep_.CoverTab[113742]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:439
			// _ = "end of CoverTab[113742]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:440
			_go_fuzz_dep_.CoverTab[113743]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:440
			// _ = "end of CoverTab[113743]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:440
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:440
		// _ = "end of CoverTab[113734]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:440
		_go_fuzz_dep_.CoverTab[113735]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:444
		// _ = "end of CoverTab[113735]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:445
	// _ = "end of CoverTab[113730]"
}

func makeStdInt64ValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:448
	_go_fuzz_dep_.CoverTab[113744]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:449
			_go_fuzz_dep_.CoverTab[113745]++
														t := ptr.asPointerTo(u.typ).Interface().(*int64)
														v := &int64Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:453
			// _ = "end of CoverTab[113745]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:454
			_go_fuzz_dep_.CoverTab[113746]++
														t := ptr.asPointerTo(u.typ).Interface().(*int64)
														v := &int64Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:458
				_go_fuzz_dep_.CoverTab[113748]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:459
				// _ = "end of CoverTab[113748]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:460
				_go_fuzz_dep_.CoverTab[113749]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:460
				// _ = "end of CoverTab[113749]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:460
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:460
			// _ = "end of CoverTab[113746]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:460
			_go_fuzz_dep_.CoverTab[113747]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:464
			// _ = "end of CoverTab[113747]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:465
	// _ = "end of CoverTab[113744]"
}

func makeStdInt64ValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:468
	_go_fuzz_dep_.CoverTab[113750]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:469
			_go_fuzz_dep_.CoverTab[113751]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:470
				_go_fuzz_dep_.CoverTab[113753]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:471
				// _ = "end of CoverTab[113753]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:472
				_go_fuzz_dep_.CoverTab[113754]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:472
				// _ = "end of CoverTab[113754]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:472
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:472
			// _ = "end of CoverTab[113751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:472
			_go_fuzz_dep_.CoverTab[113752]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*int64)
														v := &int64Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:476
			// _ = "end of CoverTab[113752]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:477
			_go_fuzz_dep_.CoverTab[113755]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:478
				_go_fuzz_dep_.CoverTab[113758]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:479
				// _ = "end of CoverTab[113758]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:480
				_go_fuzz_dep_.CoverTab[113759]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:480
				// _ = "end of CoverTab[113759]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:480
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:480
			// _ = "end of CoverTab[113755]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:480
			_go_fuzz_dep_.CoverTab[113756]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*int64)
														v := &int64Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:484
				_go_fuzz_dep_.CoverTab[113760]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:485
				// _ = "end of CoverTab[113760]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:486
				_go_fuzz_dep_.CoverTab[113761]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:486
				// _ = "end of CoverTab[113761]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:486
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:486
			// _ = "end of CoverTab[113756]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:486
			_go_fuzz_dep_.CoverTab[113757]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:490
			// _ = "end of CoverTab[113757]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:491
	// _ = "end of CoverTab[113750]"
}

func makeStdInt64ValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:494
	_go_fuzz_dep_.CoverTab[113762]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:495
			_go_fuzz_dep_.CoverTab[113763]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:498
				_go_fuzz_dep_.CoverTab[113765]++
															elem := s.Index(i)
															t := elem.Interface().(int64)
															v := &int64Value{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:503
				// _ = "end of CoverTab[113765]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:504
			// _ = "end of CoverTab[113763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:504
			_go_fuzz_dep_.CoverTab[113764]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:505
			// _ = "end of CoverTab[113764]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:507
			_go_fuzz_dep_.CoverTab[113766]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:509
				_go_fuzz_dep_.CoverTab[113768]++
															elem := s.Index(i)
															t := elem.Interface().(int64)
															v := &int64Value{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:515
					_go_fuzz_dep_.CoverTab[113770]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:516
					// _ = "end of CoverTab[113770]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:517
					_go_fuzz_dep_.CoverTab[113771]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:517
					// _ = "end of CoverTab[113771]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:517
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:517
				// _ = "end of CoverTab[113768]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:517
				_go_fuzz_dep_.CoverTab[113769]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:520
				// _ = "end of CoverTab[113769]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:521
			// _ = "end of CoverTab[113766]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:521
			_go_fuzz_dep_.CoverTab[113767]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:523
			// _ = "end of CoverTab[113767]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:524
	// _ = "end of CoverTab[113762]"
}

func makeStdInt64ValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:527
	_go_fuzz_dep_.CoverTab[113772]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:528
			_go_fuzz_dep_.CoverTab[113773]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:531
				_go_fuzz_dep_.CoverTab[113775]++
															elem := s.Index(i)
															t := elem.Interface().(*int64)
															v := &int64Value{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:536
				// _ = "end of CoverTab[113775]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:537
			// _ = "end of CoverTab[113773]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:537
			_go_fuzz_dep_.CoverTab[113774]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:538
			// _ = "end of CoverTab[113774]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:540
			_go_fuzz_dep_.CoverTab[113776]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:542
				_go_fuzz_dep_.CoverTab[113778]++
															elem := s.Index(i)
															t := elem.Interface().(*int64)
															v := &int64Value{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:548
					_go_fuzz_dep_.CoverTab[113780]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:549
					// _ = "end of CoverTab[113780]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:550
					_go_fuzz_dep_.CoverTab[113781]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:550
					// _ = "end of CoverTab[113781]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:550
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:550
				// _ = "end of CoverTab[113778]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:550
				_go_fuzz_dep_.CoverTab[113779]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:553
				// _ = "end of CoverTab[113779]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:554
			// _ = "end of CoverTab[113776]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:554
			_go_fuzz_dep_.CoverTab[113777]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:556
			// _ = "end of CoverTab[113777]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:557
	// _ = "end of CoverTab[113772]"
}

func makeStdInt64ValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:560
	_go_fuzz_dep_.CoverTab[113782]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:561
		_go_fuzz_dep_.CoverTab[113783]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:562
			_go_fuzz_dep_.CoverTab[113788]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:563
			// _ = "end of CoverTab[113788]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:564
			_go_fuzz_dep_.CoverTab[113789]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:564
			// _ = "end of CoverTab[113789]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:564
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:564
		// _ = "end of CoverTab[113783]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:564
		_go_fuzz_dep_.CoverTab[113784]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:566
			_go_fuzz_dep_.CoverTab[113790]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:567
			// _ = "end of CoverTab[113790]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:568
			_go_fuzz_dep_.CoverTab[113791]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:568
			// _ = "end of CoverTab[113791]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:568
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:568
		// _ = "end of CoverTab[113784]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:568
		_go_fuzz_dep_.CoverTab[113785]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:570
			_go_fuzz_dep_.CoverTab[113792]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:571
			// _ = "end of CoverTab[113792]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:572
			_go_fuzz_dep_.CoverTab[113793]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:572
			// _ = "end of CoverTab[113793]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:572
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:572
		// _ = "end of CoverTab[113785]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:572
		_go_fuzz_dep_.CoverTab[113786]++
													m := &int64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:574
			_go_fuzz_dep_.CoverTab[113794]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:575
			// _ = "end of CoverTab[113794]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:576
			_go_fuzz_dep_.CoverTab[113795]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:576
			// _ = "end of CoverTab[113795]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:576
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:576
		// _ = "end of CoverTab[113786]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:576
		_go_fuzz_dep_.CoverTab[113787]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:579
		// _ = "end of CoverTab[113787]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:580
	// _ = "end of CoverTab[113782]"
}

func makeStdInt64ValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:583
	_go_fuzz_dep_.CoverTab[113796]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:584
		_go_fuzz_dep_.CoverTab[113797]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:585
			_go_fuzz_dep_.CoverTab[113802]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:586
			// _ = "end of CoverTab[113802]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:587
			_go_fuzz_dep_.CoverTab[113803]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:587
			// _ = "end of CoverTab[113803]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:587
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:587
		// _ = "end of CoverTab[113797]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:587
		_go_fuzz_dep_.CoverTab[113798]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:589
			_go_fuzz_dep_.CoverTab[113804]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:590
			// _ = "end of CoverTab[113804]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:591
			_go_fuzz_dep_.CoverTab[113805]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:591
			// _ = "end of CoverTab[113805]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:591
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:591
		// _ = "end of CoverTab[113798]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:591
		_go_fuzz_dep_.CoverTab[113799]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:593
			_go_fuzz_dep_.CoverTab[113806]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:594
			// _ = "end of CoverTab[113806]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:595
			_go_fuzz_dep_.CoverTab[113807]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:595
			// _ = "end of CoverTab[113807]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:595
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:595
		// _ = "end of CoverTab[113799]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:595
		_go_fuzz_dep_.CoverTab[113800]++
													m := &int64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:597
			_go_fuzz_dep_.CoverTab[113808]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:598
			// _ = "end of CoverTab[113808]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:599
			_go_fuzz_dep_.CoverTab[113809]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:599
			// _ = "end of CoverTab[113809]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:599
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:599
		// _ = "end of CoverTab[113800]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:599
		_go_fuzz_dep_.CoverTab[113801]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:602
		// _ = "end of CoverTab[113801]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:603
	// _ = "end of CoverTab[113796]"
}

func makeStdInt64ValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:606
	_go_fuzz_dep_.CoverTab[113810]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:607
		_go_fuzz_dep_.CoverTab[113811]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:608
			_go_fuzz_dep_.CoverTab[113816]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:609
			// _ = "end of CoverTab[113816]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:610
			_go_fuzz_dep_.CoverTab[113817]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:610
			// _ = "end of CoverTab[113817]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:610
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:610
		// _ = "end of CoverTab[113811]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:610
		_go_fuzz_dep_.CoverTab[113812]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:612
			_go_fuzz_dep_.CoverTab[113818]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:613
			// _ = "end of CoverTab[113818]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:614
			_go_fuzz_dep_.CoverTab[113819]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:614
			// _ = "end of CoverTab[113819]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:614
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:614
		// _ = "end of CoverTab[113812]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:614
		_go_fuzz_dep_.CoverTab[113813]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:616
			_go_fuzz_dep_.CoverTab[113820]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:617
			// _ = "end of CoverTab[113820]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:618
			_go_fuzz_dep_.CoverTab[113821]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:618
			// _ = "end of CoverTab[113821]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:618
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:618
		// _ = "end of CoverTab[113813]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:618
		_go_fuzz_dep_.CoverTab[113814]++
													m := &int64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:620
			_go_fuzz_dep_.CoverTab[113822]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:621
			// _ = "end of CoverTab[113822]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:622
			_go_fuzz_dep_.CoverTab[113823]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:622
			// _ = "end of CoverTab[113823]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:622
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:622
		// _ = "end of CoverTab[113814]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:622
		_go_fuzz_dep_.CoverTab[113815]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:626
		// _ = "end of CoverTab[113815]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:627
	// _ = "end of CoverTab[113810]"
}

func makeStdInt64ValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:630
	_go_fuzz_dep_.CoverTab[113824]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:631
		_go_fuzz_dep_.CoverTab[113825]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:632
			_go_fuzz_dep_.CoverTab[113830]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:633
			// _ = "end of CoverTab[113830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:634
			_go_fuzz_dep_.CoverTab[113831]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:634
			// _ = "end of CoverTab[113831]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:634
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:634
		// _ = "end of CoverTab[113825]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:634
		_go_fuzz_dep_.CoverTab[113826]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:636
			_go_fuzz_dep_.CoverTab[113832]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:637
			// _ = "end of CoverTab[113832]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:638
			_go_fuzz_dep_.CoverTab[113833]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:638
			// _ = "end of CoverTab[113833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:638
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:638
		// _ = "end of CoverTab[113826]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:638
		_go_fuzz_dep_.CoverTab[113827]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:640
			_go_fuzz_dep_.CoverTab[113834]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:641
			// _ = "end of CoverTab[113834]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:642
			_go_fuzz_dep_.CoverTab[113835]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:642
			// _ = "end of CoverTab[113835]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:642
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:642
		// _ = "end of CoverTab[113827]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:642
		_go_fuzz_dep_.CoverTab[113828]++
													m := &int64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:644
			_go_fuzz_dep_.CoverTab[113836]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:645
			// _ = "end of CoverTab[113836]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:646
			_go_fuzz_dep_.CoverTab[113837]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:646
			// _ = "end of CoverTab[113837]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:646
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:646
		// _ = "end of CoverTab[113828]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:646
		_go_fuzz_dep_.CoverTab[113829]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:650
		// _ = "end of CoverTab[113829]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:651
	// _ = "end of CoverTab[113824]"
}

func makeStdUInt64ValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:654
	_go_fuzz_dep_.CoverTab[113838]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:655
			_go_fuzz_dep_.CoverTab[113839]++
														t := ptr.asPointerTo(u.typ).Interface().(*uint64)
														v := &uint64Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:659
			// _ = "end of CoverTab[113839]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:660
			_go_fuzz_dep_.CoverTab[113840]++
														t := ptr.asPointerTo(u.typ).Interface().(*uint64)
														v := &uint64Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:664
				_go_fuzz_dep_.CoverTab[113842]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:665
				// _ = "end of CoverTab[113842]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:666
				_go_fuzz_dep_.CoverTab[113843]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:666
				// _ = "end of CoverTab[113843]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:666
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:666
			// _ = "end of CoverTab[113840]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:666
			_go_fuzz_dep_.CoverTab[113841]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:670
			// _ = "end of CoverTab[113841]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:671
	// _ = "end of CoverTab[113838]"
}

func makeStdUInt64ValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:674
	_go_fuzz_dep_.CoverTab[113844]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:675
			_go_fuzz_dep_.CoverTab[113845]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:676
				_go_fuzz_dep_.CoverTab[113847]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:677
				// _ = "end of CoverTab[113847]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:678
				_go_fuzz_dep_.CoverTab[113848]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:678
				// _ = "end of CoverTab[113848]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:678
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:678
			// _ = "end of CoverTab[113845]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:678
			_go_fuzz_dep_.CoverTab[113846]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*uint64)
														v := &uint64Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:682
			// _ = "end of CoverTab[113846]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:683
			_go_fuzz_dep_.CoverTab[113849]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:684
				_go_fuzz_dep_.CoverTab[113852]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:685
				// _ = "end of CoverTab[113852]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:686
				_go_fuzz_dep_.CoverTab[113853]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:686
				// _ = "end of CoverTab[113853]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:686
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:686
			// _ = "end of CoverTab[113849]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:686
			_go_fuzz_dep_.CoverTab[113850]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*uint64)
														v := &uint64Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:690
				_go_fuzz_dep_.CoverTab[113854]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:691
				// _ = "end of CoverTab[113854]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:692
				_go_fuzz_dep_.CoverTab[113855]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:692
				// _ = "end of CoverTab[113855]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:692
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:692
			// _ = "end of CoverTab[113850]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:692
			_go_fuzz_dep_.CoverTab[113851]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:696
			// _ = "end of CoverTab[113851]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:697
	// _ = "end of CoverTab[113844]"
}

func makeStdUInt64ValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:700
	_go_fuzz_dep_.CoverTab[113856]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:701
			_go_fuzz_dep_.CoverTab[113857]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:704
				_go_fuzz_dep_.CoverTab[113859]++
															elem := s.Index(i)
															t := elem.Interface().(uint64)
															v := &uint64Value{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:709
				// _ = "end of CoverTab[113859]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:710
			// _ = "end of CoverTab[113857]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:710
			_go_fuzz_dep_.CoverTab[113858]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:711
			// _ = "end of CoverTab[113858]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:713
			_go_fuzz_dep_.CoverTab[113860]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:715
				_go_fuzz_dep_.CoverTab[113862]++
															elem := s.Index(i)
															t := elem.Interface().(uint64)
															v := &uint64Value{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:721
					_go_fuzz_dep_.CoverTab[113864]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:722
					// _ = "end of CoverTab[113864]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:723
					_go_fuzz_dep_.CoverTab[113865]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:723
					// _ = "end of CoverTab[113865]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:723
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:723
				// _ = "end of CoverTab[113862]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:723
				_go_fuzz_dep_.CoverTab[113863]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:726
				// _ = "end of CoverTab[113863]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:727
			// _ = "end of CoverTab[113860]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:727
			_go_fuzz_dep_.CoverTab[113861]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:729
			// _ = "end of CoverTab[113861]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:730
	// _ = "end of CoverTab[113856]"
}

func makeStdUInt64ValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:733
	_go_fuzz_dep_.CoverTab[113866]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:734
			_go_fuzz_dep_.CoverTab[113867]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:737
				_go_fuzz_dep_.CoverTab[113869]++
															elem := s.Index(i)
															t := elem.Interface().(*uint64)
															v := &uint64Value{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:742
				// _ = "end of CoverTab[113869]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:743
			// _ = "end of CoverTab[113867]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:743
			_go_fuzz_dep_.CoverTab[113868]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:744
			// _ = "end of CoverTab[113868]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:746
			_go_fuzz_dep_.CoverTab[113870]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:748
				_go_fuzz_dep_.CoverTab[113872]++
															elem := s.Index(i)
															t := elem.Interface().(*uint64)
															v := &uint64Value{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:754
					_go_fuzz_dep_.CoverTab[113874]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:755
					// _ = "end of CoverTab[113874]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:756
					_go_fuzz_dep_.CoverTab[113875]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:756
					// _ = "end of CoverTab[113875]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:756
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:756
				// _ = "end of CoverTab[113872]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:756
				_go_fuzz_dep_.CoverTab[113873]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:759
				// _ = "end of CoverTab[113873]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:760
			// _ = "end of CoverTab[113870]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:760
			_go_fuzz_dep_.CoverTab[113871]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:762
			// _ = "end of CoverTab[113871]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:763
	// _ = "end of CoverTab[113866]"
}

func makeStdUInt64ValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:766
	_go_fuzz_dep_.CoverTab[113876]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:767
		_go_fuzz_dep_.CoverTab[113877]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:768
			_go_fuzz_dep_.CoverTab[113882]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:769
			// _ = "end of CoverTab[113882]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:770
			_go_fuzz_dep_.CoverTab[113883]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:770
			// _ = "end of CoverTab[113883]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:770
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:770
		// _ = "end of CoverTab[113877]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:770
		_go_fuzz_dep_.CoverTab[113878]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:772
			_go_fuzz_dep_.CoverTab[113884]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:773
			// _ = "end of CoverTab[113884]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:774
			_go_fuzz_dep_.CoverTab[113885]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:774
			// _ = "end of CoverTab[113885]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:774
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:774
		// _ = "end of CoverTab[113878]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:774
		_go_fuzz_dep_.CoverTab[113879]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:776
			_go_fuzz_dep_.CoverTab[113886]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:777
			// _ = "end of CoverTab[113886]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:778
			_go_fuzz_dep_.CoverTab[113887]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:778
			// _ = "end of CoverTab[113887]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:778
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:778
		// _ = "end of CoverTab[113879]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:778
		_go_fuzz_dep_.CoverTab[113880]++
													m := &uint64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:780
			_go_fuzz_dep_.CoverTab[113888]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:781
			// _ = "end of CoverTab[113888]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:782
			_go_fuzz_dep_.CoverTab[113889]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:782
			// _ = "end of CoverTab[113889]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:782
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:782
		// _ = "end of CoverTab[113880]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:782
		_go_fuzz_dep_.CoverTab[113881]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:785
		// _ = "end of CoverTab[113881]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:786
	// _ = "end of CoverTab[113876]"
}

func makeStdUInt64ValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:789
	_go_fuzz_dep_.CoverTab[113890]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:790
		_go_fuzz_dep_.CoverTab[113891]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:791
			_go_fuzz_dep_.CoverTab[113896]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:792
			// _ = "end of CoverTab[113896]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:793
			_go_fuzz_dep_.CoverTab[113897]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:793
			// _ = "end of CoverTab[113897]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:793
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:793
		// _ = "end of CoverTab[113891]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:793
		_go_fuzz_dep_.CoverTab[113892]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:795
			_go_fuzz_dep_.CoverTab[113898]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:796
			// _ = "end of CoverTab[113898]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:797
			_go_fuzz_dep_.CoverTab[113899]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:797
			// _ = "end of CoverTab[113899]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:797
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:797
		// _ = "end of CoverTab[113892]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:797
		_go_fuzz_dep_.CoverTab[113893]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:799
			_go_fuzz_dep_.CoverTab[113900]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:800
			// _ = "end of CoverTab[113900]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:801
			_go_fuzz_dep_.CoverTab[113901]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:801
			// _ = "end of CoverTab[113901]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:801
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:801
		// _ = "end of CoverTab[113893]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:801
		_go_fuzz_dep_.CoverTab[113894]++
													m := &uint64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:803
			_go_fuzz_dep_.CoverTab[113902]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:804
			// _ = "end of CoverTab[113902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:805
			_go_fuzz_dep_.CoverTab[113903]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:805
			// _ = "end of CoverTab[113903]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:805
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:805
		// _ = "end of CoverTab[113894]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:805
		_go_fuzz_dep_.CoverTab[113895]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:808
		// _ = "end of CoverTab[113895]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:809
	// _ = "end of CoverTab[113890]"
}

func makeStdUInt64ValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:812
	_go_fuzz_dep_.CoverTab[113904]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:813
		_go_fuzz_dep_.CoverTab[113905]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:814
			_go_fuzz_dep_.CoverTab[113910]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:815
			// _ = "end of CoverTab[113910]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:816
			_go_fuzz_dep_.CoverTab[113911]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:816
			// _ = "end of CoverTab[113911]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:816
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:816
		// _ = "end of CoverTab[113905]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:816
		_go_fuzz_dep_.CoverTab[113906]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:818
			_go_fuzz_dep_.CoverTab[113912]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:819
			// _ = "end of CoverTab[113912]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:820
			_go_fuzz_dep_.CoverTab[113913]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:820
			// _ = "end of CoverTab[113913]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:820
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:820
		// _ = "end of CoverTab[113906]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:820
		_go_fuzz_dep_.CoverTab[113907]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:822
			_go_fuzz_dep_.CoverTab[113914]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:823
			// _ = "end of CoverTab[113914]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:824
			_go_fuzz_dep_.CoverTab[113915]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:824
			// _ = "end of CoverTab[113915]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:824
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:824
		// _ = "end of CoverTab[113907]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:824
		_go_fuzz_dep_.CoverTab[113908]++
													m := &uint64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:826
			_go_fuzz_dep_.CoverTab[113916]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:827
			// _ = "end of CoverTab[113916]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:828
			_go_fuzz_dep_.CoverTab[113917]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:828
			// _ = "end of CoverTab[113917]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:828
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:828
		// _ = "end of CoverTab[113908]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:828
		_go_fuzz_dep_.CoverTab[113909]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:832
		// _ = "end of CoverTab[113909]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:833
	// _ = "end of CoverTab[113904]"
}

func makeStdUInt64ValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:836
	_go_fuzz_dep_.CoverTab[113918]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:837
		_go_fuzz_dep_.CoverTab[113919]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:838
			_go_fuzz_dep_.CoverTab[113924]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:839
			// _ = "end of CoverTab[113924]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:840
			_go_fuzz_dep_.CoverTab[113925]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:840
			// _ = "end of CoverTab[113925]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:840
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:840
		// _ = "end of CoverTab[113919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:840
		_go_fuzz_dep_.CoverTab[113920]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:842
			_go_fuzz_dep_.CoverTab[113926]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:843
			// _ = "end of CoverTab[113926]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:844
			_go_fuzz_dep_.CoverTab[113927]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:844
			// _ = "end of CoverTab[113927]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:844
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:844
		// _ = "end of CoverTab[113920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:844
		_go_fuzz_dep_.CoverTab[113921]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:846
			_go_fuzz_dep_.CoverTab[113928]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:847
			// _ = "end of CoverTab[113928]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:848
			_go_fuzz_dep_.CoverTab[113929]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:848
			// _ = "end of CoverTab[113929]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:848
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:848
		// _ = "end of CoverTab[113921]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:848
		_go_fuzz_dep_.CoverTab[113922]++
													m := &uint64Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:850
			_go_fuzz_dep_.CoverTab[113930]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:851
			// _ = "end of CoverTab[113930]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:852
			_go_fuzz_dep_.CoverTab[113931]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:852
			// _ = "end of CoverTab[113931]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:852
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:852
		// _ = "end of CoverTab[113922]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:852
		_go_fuzz_dep_.CoverTab[113923]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:856
		// _ = "end of CoverTab[113923]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:857
	// _ = "end of CoverTab[113918]"
}

func makeStdInt32ValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:860
	_go_fuzz_dep_.CoverTab[113932]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:861
			_go_fuzz_dep_.CoverTab[113933]++
														t := ptr.asPointerTo(u.typ).Interface().(*int32)
														v := &int32Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:865
			// _ = "end of CoverTab[113933]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:866
			_go_fuzz_dep_.CoverTab[113934]++
														t := ptr.asPointerTo(u.typ).Interface().(*int32)
														v := &int32Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:870
				_go_fuzz_dep_.CoverTab[113936]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:871
				// _ = "end of CoverTab[113936]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:872
				_go_fuzz_dep_.CoverTab[113937]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:872
				// _ = "end of CoverTab[113937]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:872
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:872
			// _ = "end of CoverTab[113934]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:872
			_go_fuzz_dep_.CoverTab[113935]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:876
			// _ = "end of CoverTab[113935]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:877
	// _ = "end of CoverTab[113932]"
}

func makeStdInt32ValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:880
	_go_fuzz_dep_.CoverTab[113938]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:881
			_go_fuzz_dep_.CoverTab[113939]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:882
				_go_fuzz_dep_.CoverTab[113941]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:883
				// _ = "end of CoverTab[113941]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:884
				_go_fuzz_dep_.CoverTab[113942]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:884
				// _ = "end of CoverTab[113942]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:884
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:884
			// _ = "end of CoverTab[113939]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:884
			_go_fuzz_dep_.CoverTab[113940]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*int32)
														v := &int32Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:888
			// _ = "end of CoverTab[113940]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:889
			_go_fuzz_dep_.CoverTab[113943]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:890
				_go_fuzz_dep_.CoverTab[113946]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:891
				// _ = "end of CoverTab[113946]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:892
				_go_fuzz_dep_.CoverTab[113947]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:892
				// _ = "end of CoverTab[113947]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:892
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:892
			// _ = "end of CoverTab[113943]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:892
			_go_fuzz_dep_.CoverTab[113944]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*int32)
														v := &int32Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:896
				_go_fuzz_dep_.CoverTab[113948]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:897
				// _ = "end of CoverTab[113948]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:898
				_go_fuzz_dep_.CoverTab[113949]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:898
				// _ = "end of CoverTab[113949]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:898
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:898
			// _ = "end of CoverTab[113944]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:898
			_go_fuzz_dep_.CoverTab[113945]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:902
			// _ = "end of CoverTab[113945]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:903
	// _ = "end of CoverTab[113938]"
}

func makeStdInt32ValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:906
	_go_fuzz_dep_.CoverTab[113950]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:907
			_go_fuzz_dep_.CoverTab[113951]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:910
				_go_fuzz_dep_.CoverTab[113953]++
															elem := s.Index(i)
															t := elem.Interface().(int32)
															v := &int32Value{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:915
				// _ = "end of CoverTab[113953]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:916
			// _ = "end of CoverTab[113951]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:916
			_go_fuzz_dep_.CoverTab[113952]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:917
			// _ = "end of CoverTab[113952]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:919
			_go_fuzz_dep_.CoverTab[113954]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:921
				_go_fuzz_dep_.CoverTab[113956]++
															elem := s.Index(i)
															t := elem.Interface().(int32)
															v := &int32Value{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:927
					_go_fuzz_dep_.CoverTab[113958]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:928
					// _ = "end of CoverTab[113958]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:929
					_go_fuzz_dep_.CoverTab[113959]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:929
					// _ = "end of CoverTab[113959]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:929
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:929
				// _ = "end of CoverTab[113956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:929
				_go_fuzz_dep_.CoverTab[113957]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:932
				// _ = "end of CoverTab[113957]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:933
			// _ = "end of CoverTab[113954]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:933
			_go_fuzz_dep_.CoverTab[113955]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:935
			// _ = "end of CoverTab[113955]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:936
	// _ = "end of CoverTab[113950]"
}

func makeStdInt32ValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:939
	_go_fuzz_dep_.CoverTab[113960]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:940
			_go_fuzz_dep_.CoverTab[113961]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:943
				_go_fuzz_dep_.CoverTab[113963]++
															elem := s.Index(i)
															t := elem.Interface().(*int32)
															v := &int32Value{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:948
				// _ = "end of CoverTab[113963]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:949
			// _ = "end of CoverTab[113961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:949
			_go_fuzz_dep_.CoverTab[113962]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:950
			// _ = "end of CoverTab[113962]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:952
			_go_fuzz_dep_.CoverTab[113964]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:954
				_go_fuzz_dep_.CoverTab[113966]++
															elem := s.Index(i)
															t := elem.Interface().(*int32)
															v := &int32Value{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:960
					_go_fuzz_dep_.CoverTab[113968]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:961
					// _ = "end of CoverTab[113968]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:962
					_go_fuzz_dep_.CoverTab[113969]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:962
					// _ = "end of CoverTab[113969]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:962
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:962
				// _ = "end of CoverTab[113966]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:962
				_go_fuzz_dep_.CoverTab[113967]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:965
				// _ = "end of CoverTab[113967]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:966
			// _ = "end of CoverTab[113964]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:966
			_go_fuzz_dep_.CoverTab[113965]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:968
			// _ = "end of CoverTab[113965]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:969
	// _ = "end of CoverTab[113960]"
}

func makeStdInt32ValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:972
	_go_fuzz_dep_.CoverTab[113970]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:973
		_go_fuzz_dep_.CoverTab[113971]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:974
			_go_fuzz_dep_.CoverTab[113976]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:975
			// _ = "end of CoverTab[113976]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:976
			_go_fuzz_dep_.CoverTab[113977]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:976
			// _ = "end of CoverTab[113977]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:976
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:976
		// _ = "end of CoverTab[113971]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:976
		_go_fuzz_dep_.CoverTab[113972]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:978
			_go_fuzz_dep_.CoverTab[113978]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:979
			// _ = "end of CoverTab[113978]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:980
			_go_fuzz_dep_.CoverTab[113979]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:980
			// _ = "end of CoverTab[113979]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:980
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:980
		// _ = "end of CoverTab[113972]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:980
		_go_fuzz_dep_.CoverTab[113973]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:982
			_go_fuzz_dep_.CoverTab[113980]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:983
			// _ = "end of CoverTab[113980]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:984
			_go_fuzz_dep_.CoverTab[113981]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:984
			// _ = "end of CoverTab[113981]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:984
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:984
		// _ = "end of CoverTab[113973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:984
		_go_fuzz_dep_.CoverTab[113974]++
													m := &int32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:986
			_go_fuzz_dep_.CoverTab[113982]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:987
			// _ = "end of CoverTab[113982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:988
			_go_fuzz_dep_.CoverTab[113983]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:988
			// _ = "end of CoverTab[113983]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:988
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:988
		// _ = "end of CoverTab[113974]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:988
		_go_fuzz_dep_.CoverTab[113975]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:991
		// _ = "end of CoverTab[113975]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:992
	// _ = "end of CoverTab[113970]"
}

func makeStdInt32ValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:995
	_go_fuzz_dep_.CoverTab[113984]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:996
		_go_fuzz_dep_.CoverTab[113985]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:997
			_go_fuzz_dep_.CoverTab[113990]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:998
			// _ = "end of CoverTab[113990]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:999
			_go_fuzz_dep_.CoverTab[113991]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:999
			// _ = "end of CoverTab[113991]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:999
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:999
		// _ = "end of CoverTab[113985]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:999
		_go_fuzz_dep_.CoverTab[113986]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1001
			_go_fuzz_dep_.CoverTab[113992]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1002
			// _ = "end of CoverTab[113992]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1003
			_go_fuzz_dep_.CoverTab[113993]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1003
			// _ = "end of CoverTab[113993]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1003
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1003
		// _ = "end of CoverTab[113986]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1003
		_go_fuzz_dep_.CoverTab[113987]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1005
			_go_fuzz_dep_.CoverTab[113994]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1006
			// _ = "end of CoverTab[113994]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1007
			_go_fuzz_dep_.CoverTab[113995]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1007
			// _ = "end of CoverTab[113995]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1007
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1007
		// _ = "end of CoverTab[113987]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1007
		_go_fuzz_dep_.CoverTab[113988]++
													m := &int32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1009
			_go_fuzz_dep_.CoverTab[113996]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1010
			// _ = "end of CoverTab[113996]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1011
			_go_fuzz_dep_.CoverTab[113997]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1011
			// _ = "end of CoverTab[113997]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1011
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1011
		// _ = "end of CoverTab[113988]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1011
		_go_fuzz_dep_.CoverTab[113989]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1014
		// _ = "end of CoverTab[113989]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1015
	// _ = "end of CoverTab[113984]"
}

func makeStdInt32ValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1018
	_go_fuzz_dep_.CoverTab[113998]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1019
		_go_fuzz_dep_.CoverTab[113999]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1020
			_go_fuzz_dep_.CoverTab[114004]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1021
			// _ = "end of CoverTab[114004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1022
			_go_fuzz_dep_.CoverTab[114005]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1022
			// _ = "end of CoverTab[114005]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1022
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1022
		// _ = "end of CoverTab[113999]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1022
		_go_fuzz_dep_.CoverTab[114000]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1024
			_go_fuzz_dep_.CoverTab[114006]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1025
			// _ = "end of CoverTab[114006]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1026
			_go_fuzz_dep_.CoverTab[114007]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1026
			// _ = "end of CoverTab[114007]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1026
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1026
		// _ = "end of CoverTab[114000]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1026
		_go_fuzz_dep_.CoverTab[114001]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1028
			_go_fuzz_dep_.CoverTab[114008]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1029
			// _ = "end of CoverTab[114008]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1030
			_go_fuzz_dep_.CoverTab[114009]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1030
			// _ = "end of CoverTab[114009]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1030
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1030
		// _ = "end of CoverTab[114001]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1030
		_go_fuzz_dep_.CoverTab[114002]++
													m := &int32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1032
			_go_fuzz_dep_.CoverTab[114010]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1033
			// _ = "end of CoverTab[114010]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1034
			_go_fuzz_dep_.CoverTab[114011]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1034
			// _ = "end of CoverTab[114011]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1034
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1034
		// _ = "end of CoverTab[114002]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1034
		_go_fuzz_dep_.CoverTab[114003]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1038
		// _ = "end of CoverTab[114003]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1039
	// _ = "end of CoverTab[113998]"
}

func makeStdInt32ValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1042
	_go_fuzz_dep_.CoverTab[114012]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1043
		_go_fuzz_dep_.CoverTab[114013]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1044
			_go_fuzz_dep_.CoverTab[114018]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1045
			// _ = "end of CoverTab[114018]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1046
			_go_fuzz_dep_.CoverTab[114019]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1046
			// _ = "end of CoverTab[114019]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1046
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1046
		// _ = "end of CoverTab[114013]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1046
		_go_fuzz_dep_.CoverTab[114014]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1048
			_go_fuzz_dep_.CoverTab[114020]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1049
			// _ = "end of CoverTab[114020]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1050
			_go_fuzz_dep_.CoverTab[114021]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1050
			// _ = "end of CoverTab[114021]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1050
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1050
		// _ = "end of CoverTab[114014]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1050
		_go_fuzz_dep_.CoverTab[114015]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1052
			_go_fuzz_dep_.CoverTab[114022]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1053
			// _ = "end of CoverTab[114022]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1054
			_go_fuzz_dep_.CoverTab[114023]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1054
			// _ = "end of CoverTab[114023]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1054
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1054
		// _ = "end of CoverTab[114015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1054
		_go_fuzz_dep_.CoverTab[114016]++
													m := &int32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1056
			_go_fuzz_dep_.CoverTab[114024]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1057
			// _ = "end of CoverTab[114024]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1058
			_go_fuzz_dep_.CoverTab[114025]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1058
			// _ = "end of CoverTab[114025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1058
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1058
		// _ = "end of CoverTab[114016]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1058
		_go_fuzz_dep_.CoverTab[114017]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1062
		// _ = "end of CoverTab[114017]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1063
	// _ = "end of CoverTab[114012]"
}

func makeStdUInt32ValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1066
	_go_fuzz_dep_.CoverTab[114026]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1067
			_go_fuzz_dep_.CoverTab[114027]++
														t := ptr.asPointerTo(u.typ).Interface().(*uint32)
														v := &uint32Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1071
			// _ = "end of CoverTab[114027]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1072
			_go_fuzz_dep_.CoverTab[114028]++
														t := ptr.asPointerTo(u.typ).Interface().(*uint32)
														v := &uint32Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1076
				_go_fuzz_dep_.CoverTab[114030]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1077
				// _ = "end of CoverTab[114030]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1078
				_go_fuzz_dep_.CoverTab[114031]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1078
				// _ = "end of CoverTab[114031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1078
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1078
			// _ = "end of CoverTab[114028]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1078
			_go_fuzz_dep_.CoverTab[114029]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1082
			// _ = "end of CoverTab[114029]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1083
	// _ = "end of CoverTab[114026]"
}

func makeStdUInt32ValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1086
	_go_fuzz_dep_.CoverTab[114032]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1087
			_go_fuzz_dep_.CoverTab[114033]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1088
				_go_fuzz_dep_.CoverTab[114035]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1089
				// _ = "end of CoverTab[114035]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1090
				_go_fuzz_dep_.CoverTab[114036]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1090
				// _ = "end of CoverTab[114036]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1090
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1090
			// _ = "end of CoverTab[114033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1090
			_go_fuzz_dep_.CoverTab[114034]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*uint32)
														v := &uint32Value{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1094
			// _ = "end of CoverTab[114034]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1095
			_go_fuzz_dep_.CoverTab[114037]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1096
				_go_fuzz_dep_.CoverTab[114040]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1097
				// _ = "end of CoverTab[114040]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1098
				_go_fuzz_dep_.CoverTab[114041]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1098
				// _ = "end of CoverTab[114041]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1098
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1098
			// _ = "end of CoverTab[114037]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1098
			_go_fuzz_dep_.CoverTab[114038]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*uint32)
														v := &uint32Value{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1102
				_go_fuzz_dep_.CoverTab[114042]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1103
				// _ = "end of CoverTab[114042]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1104
				_go_fuzz_dep_.CoverTab[114043]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1104
				// _ = "end of CoverTab[114043]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1104
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1104
			// _ = "end of CoverTab[114038]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1104
			_go_fuzz_dep_.CoverTab[114039]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1108
			// _ = "end of CoverTab[114039]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1109
	// _ = "end of CoverTab[114032]"
}

func makeStdUInt32ValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1112
	_go_fuzz_dep_.CoverTab[114044]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1113
			_go_fuzz_dep_.CoverTab[114045]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1116
				_go_fuzz_dep_.CoverTab[114047]++
															elem := s.Index(i)
															t := elem.Interface().(uint32)
															v := &uint32Value{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1121
				// _ = "end of CoverTab[114047]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1122
			// _ = "end of CoverTab[114045]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1122
			_go_fuzz_dep_.CoverTab[114046]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1123
			// _ = "end of CoverTab[114046]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1125
			_go_fuzz_dep_.CoverTab[114048]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1127
				_go_fuzz_dep_.CoverTab[114050]++
															elem := s.Index(i)
															t := elem.Interface().(uint32)
															v := &uint32Value{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1133
					_go_fuzz_dep_.CoverTab[114052]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1134
					// _ = "end of CoverTab[114052]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1135
					_go_fuzz_dep_.CoverTab[114053]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1135
					// _ = "end of CoverTab[114053]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1135
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1135
				// _ = "end of CoverTab[114050]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1135
				_go_fuzz_dep_.CoverTab[114051]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1138
				// _ = "end of CoverTab[114051]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1139
			// _ = "end of CoverTab[114048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1139
			_go_fuzz_dep_.CoverTab[114049]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1141
			// _ = "end of CoverTab[114049]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1142
	// _ = "end of CoverTab[114044]"
}

func makeStdUInt32ValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1145
	_go_fuzz_dep_.CoverTab[114054]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1146
			_go_fuzz_dep_.CoverTab[114055]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1149
				_go_fuzz_dep_.CoverTab[114057]++
															elem := s.Index(i)
															t := elem.Interface().(*uint32)
															v := &uint32Value{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1154
				// _ = "end of CoverTab[114057]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1155
			// _ = "end of CoverTab[114055]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1155
			_go_fuzz_dep_.CoverTab[114056]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1156
			// _ = "end of CoverTab[114056]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1158
			_go_fuzz_dep_.CoverTab[114058]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1160
				_go_fuzz_dep_.CoverTab[114060]++
															elem := s.Index(i)
															t := elem.Interface().(*uint32)
															v := &uint32Value{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1166
					_go_fuzz_dep_.CoverTab[114062]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1167
					// _ = "end of CoverTab[114062]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1168
					_go_fuzz_dep_.CoverTab[114063]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1168
					// _ = "end of CoverTab[114063]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1168
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1168
				// _ = "end of CoverTab[114060]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1168
				_go_fuzz_dep_.CoverTab[114061]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1171
				// _ = "end of CoverTab[114061]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1172
			// _ = "end of CoverTab[114058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1172
			_go_fuzz_dep_.CoverTab[114059]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1174
			// _ = "end of CoverTab[114059]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1175
	// _ = "end of CoverTab[114054]"
}

func makeStdUInt32ValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1178
	_go_fuzz_dep_.CoverTab[114064]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1179
		_go_fuzz_dep_.CoverTab[114065]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1180
			_go_fuzz_dep_.CoverTab[114070]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1181
			// _ = "end of CoverTab[114070]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1182
			_go_fuzz_dep_.CoverTab[114071]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1182
			// _ = "end of CoverTab[114071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1182
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1182
		// _ = "end of CoverTab[114065]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1182
		_go_fuzz_dep_.CoverTab[114066]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1184
			_go_fuzz_dep_.CoverTab[114072]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1185
			// _ = "end of CoverTab[114072]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1186
			_go_fuzz_dep_.CoverTab[114073]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1186
			// _ = "end of CoverTab[114073]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1186
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1186
		// _ = "end of CoverTab[114066]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1186
		_go_fuzz_dep_.CoverTab[114067]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1188
			_go_fuzz_dep_.CoverTab[114074]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1189
			// _ = "end of CoverTab[114074]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1190
			_go_fuzz_dep_.CoverTab[114075]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1190
			// _ = "end of CoverTab[114075]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1190
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1190
		// _ = "end of CoverTab[114067]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1190
		_go_fuzz_dep_.CoverTab[114068]++
													m := &uint32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1192
			_go_fuzz_dep_.CoverTab[114076]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1193
			// _ = "end of CoverTab[114076]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1194
			_go_fuzz_dep_.CoverTab[114077]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1194
			// _ = "end of CoverTab[114077]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1194
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1194
		// _ = "end of CoverTab[114068]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1194
		_go_fuzz_dep_.CoverTab[114069]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1197
		// _ = "end of CoverTab[114069]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1198
	// _ = "end of CoverTab[114064]"
}

func makeStdUInt32ValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1201
	_go_fuzz_dep_.CoverTab[114078]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1202
		_go_fuzz_dep_.CoverTab[114079]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1203
			_go_fuzz_dep_.CoverTab[114084]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1204
			// _ = "end of CoverTab[114084]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1205
			_go_fuzz_dep_.CoverTab[114085]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1205
			// _ = "end of CoverTab[114085]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1205
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1205
		// _ = "end of CoverTab[114079]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1205
		_go_fuzz_dep_.CoverTab[114080]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1207
			_go_fuzz_dep_.CoverTab[114086]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1208
			// _ = "end of CoverTab[114086]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1209
			_go_fuzz_dep_.CoverTab[114087]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1209
			// _ = "end of CoverTab[114087]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1209
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1209
		// _ = "end of CoverTab[114080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1209
		_go_fuzz_dep_.CoverTab[114081]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1211
			_go_fuzz_dep_.CoverTab[114088]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1212
			// _ = "end of CoverTab[114088]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1213
			_go_fuzz_dep_.CoverTab[114089]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1213
			// _ = "end of CoverTab[114089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1213
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1213
		// _ = "end of CoverTab[114081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1213
		_go_fuzz_dep_.CoverTab[114082]++
													m := &uint32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1215
			_go_fuzz_dep_.CoverTab[114090]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1216
			// _ = "end of CoverTab[114090]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1217
			_go_fuzz_dep_.CoverTab[114091]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1217
			// _ = "end of CoverTab[114091]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1217
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1217
		// _ = "end of CoverTab[114082]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1217
		_go_fuzz_dep_.CoverTab[114083]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1220
		// _ = "end of CoverTab[114083]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1221
	// _ = "end of CoverTab[114078]"
}

func makeStdUInt32ValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1224
	_go_fuzz_dep_.CoverTab[114092]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1225
		_go_fuzz_dep_.CoverTab[114093]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1226
			_go_fuzz_dep_.CoverTab[114098]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1227
			// _ = "end of CoverTab[114098]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1228
			_go_fuzz_dep_.CoverTab[114099]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1228
			// _ = "end of CoverTab[114099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1228
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1228
		// _ = "end of CoverTab[114093]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1228
		_go_fuzz_dep_.CoverTab[114094]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1230
			_go_fuzz_dep_.CoverTab[114100]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1231
			// _ = "end of CoverTab[114100]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1232
			_go_fuzz_dep_.CoverTab[114101]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1232
			// _ = "end of CoverTab[114101]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1232
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1232
		// _ = "end of CoverTab[114094]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1232
		_go_fuzz_dep_.CoverTab[114095]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1234
			_go_fuzz_dep_.CoverTab[114102]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1235
			// _ = "end of CoverTab[114102]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1236
			_go_fuzz_dep_.CoverTab[114103]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1236
			// _ = "end of CoverTab[114103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1236
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1236
		// _ = "end of CoverTab[114095]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1236
		_go_fuzz_dep_.CoverTab[114096]++
													m := &uint32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1238
			_go_fuzz_dep_.CoverTab[114104]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1239
			// _ = "end of CoverTab[114104]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1240
			_go_fuzz_dep_.CoverTab[114105]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1240
			// _ = "end of CoverTab[114105]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1240
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1240
		// _ = "end of CoverTab[114096]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1240
		_go_fuzz_dep_.CoverTab[114097]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1244
		// _ = "end of CoverTab[114097]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1245
	// _ = "end of CoverTab[114092]"
}

func makeStdUInt32ValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1248
	_go_fuzz_dep_.CoverTab[114106]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1249
		_go_fuzz_dep_.CoverTab[114107]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1250
			_go_fuzz_dep_.CoverTab[114112]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1251
			// _ = "end of CoverTab[114112]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1252
			_go_fuzz_dep_.CoverTab[114113]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1252
			// _ = "end of CoverTab[114113]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1252
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1252
		// _ = "end of CoverTab[114107]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1252
		_go_fuzz_dep_.CoverTab[114108]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1254
			_go_fuzz_dep_.CoverTab[114114]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1255
			// _ = "end of CoverTab[114114]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1256
			_go_fuzz_dep_.CoverTab[114115]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1256
			// _ = "end of CoverTab[114115]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1256
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1256
		// _ = "end of CoverTab[114108]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1256
		_go_fuzz_dep_.CoverTab[114109]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1258
			_go_fuzz_dep_.CoverTab[114116]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1259
			// _ = "end of CoverTab[114116]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1260
			_go_fuzz_dep_.CoverTab[114117]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1260
			// _ = "end of CoverTab[114117]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1260
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1260
		// _ = "end of CoverTab[114109]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1260
		_go_fuzz_dep_.CoverTab[114110]++
													m := &uint32Value{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1262
			_go_fuzz_dep_.CoverTab[114118]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1263
			// _ = "end of CoverTab[114118]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1264
			_go_fuzz_dep_.CoverTab[114119]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1264
			// _ = "end of CoverTab[114119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1264
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1264
		// _ = "end of CoverTab[114110]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1264
		_go_fuzz_dep_.CoverTab[114111]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1268
		// _ = "end of CoverTab[114111]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1269
	// _ = "end of CoverTab[114106]"
}

func makeStdBoolValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1272
	_go_fuzz_dep_.CoverTab[114120]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1273
			_go_fuzz_dep_.CoverTab[114121]++
														t := ptr.asPointerTo(u.typ).Interface().(*bool)
														v := &boolValue{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1277
			// _ = "end of CoverTab[114121]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1278
			_go_fuzz_dep_.CoverTab[114122]++
														t := ptr.asPointerTo(u.typ).Interface().(*bool)
														v := &boolValue{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1282
				_go_fuzz_dep_.CoverTab[114124]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1283
				// _ = "end of CoverTab[114124]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1284
				_go_fuzz_dep_.CoverTab[114125]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1284
				// _ = "end of CoverTab[114125]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1284
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1284
			// _ = "end of CoverTab[114122]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1284
			_go_fuzz_dep_.CoverTab[114123]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1288
			// _ = "end of CoverTab[114123]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1289
	// _ = "end of CoverTab[114120]"
}

func makeStdBoolValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1292
	_go_fuzz_dep_.CoverTab[114126]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1293
			_go_fuzz_dep_.CoverTab[114127]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1294
				_go_fuzz_dep_.CoverTab[114129]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1295
				// _ = "end of CoverTab[114129]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1296
				_go_fuzz_dep_.CoverTab[114130]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1296
				// _ = "end of CoverTab[114130]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1296
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1296
			// _ = "end of CoverTab[114127]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1296
			_go_fuzz_dep_.CoverTab[114128]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*bool)
														v := &boolValue{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1300
			// _ = "end of CoverTab[114128]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1301
			_go_fuzz_dep_.CoverTab[114131]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1302
				_go_fuzz_dep_.CoverTab[114134]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1303
				// _ = "end of CoverTab[114134]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1304
				_go_fuzz_dep_.CoverTab[114135]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1304
				// _ = "end of CoverTab[114135]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1304
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1304
			// _ = "end of CoverTab[114131]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1304
			_go_fuzz_dep_.CoverTab[114132]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*bool)
														v := &boolValue{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1308
				_go_fuzz_dep_.CoverTab[114136]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1309
				// _ = "end of CoverTab[114136]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1310
				_go_fuzz_dep_.CoverTab[114137]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1310
				// _ = "end of CoverTab[114137]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1310
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1310
			// _ = "end of CoverTab[114132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1310
			_go_fuzz_dep_.CoverTab[114133]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1314
			// _ = "end of CoverTab[114133]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1315
	// _ = "end of CoverTab[114126]"
}

func makeStdBoolValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1318
	_go_fuzz_dep_.CoverTab[114138]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1319
			_go_fuzz_dep_.CoverTab[114139]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1322
				_go_fuzz_dep_.CoverTab[114141]++
															elem := s.Index(i)
															t := elem.Interface().(bool)
															v := &boolValue{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1327
				// _ = "end of CoverTab[114141]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1328
			// _ = "end of CoverTab[114139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1328
			_go_fuzz_dep_.CoverTab[114140]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1329
			// _ = "end of CoverTab[114140]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1331
			_go_fuzz_dep_.CoverTab[114142]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1333
				_go_fuzz_dep_.CoverTab[114144]++
															elem := s.Index(i)
															t := elem.Interface().(bool)
															v := &boolValue{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1339
					_go_fuzz_dep_.CoverTab[114146]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1340
					// _ = "end of CoverTab[114146]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1341
					_go_fuzz_dep_.CoverTab[114147]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1341
					// _ = "end of CoverTab[114147]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1341
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1341
				// _ = "end of CoverTab[114144]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1341
				_go_fuzz_dep_.CoverTab[114145]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1344
				// _ = "end of CoverTab[114145]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1345
			// _ = "end of CoverTab[114142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1345
			_go_fuzz_dep_.CoverTab[114143]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1347
			// _ = "end of CoverTab[114143]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1348
	// _ = "end of CoverTab[114138]"
}

func makeStdBoolValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1351
	_go_fuzz_dep_.CoverTab[114148]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1352
			_go_fuzz_dep_.CoverTab[114149]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1355
				_go_fuzz_dep_.CoverTab[114151]++
															elem := s.Index(i)
															t := elem.Interface().(*bool)
															v := &boolValue{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1360
				// _ = "end of CoverTab[114151]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1361
			// _ = "end of CoverTab[114149]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1361
			_go_fuzz_dep_.CoverTab[114150]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1362
			// _ = "end of CoverTab[114150]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1364
			_go_fuzz_dep_.CoverTab[114152]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1366
				_go_fuzz_dep_.CoverTab[114154]++
															elem := s.Index(i)
															t := elem.Interface().(*bool)
															v := &boolValue{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1372
					_go_fuzz_dep_.CoverTab[114156]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1373
					// _ = "end of CoverTab[114156]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1374
					_go_fuzz_dep_.CoverTab[114157]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1374
					// _ = "end of CoverTab[114157]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1374
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1374
				// _ = "end of CoverTab[114154]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1374
				_go_fuzz_dep_.CoverTab[114155]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1377
				// _ = "end of CoverTab[114155]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1378
			// _ = "end of CoverTab[114152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1378
			_go_fuzz_dep_.CoverTab[114153]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1380
			// _ = "end of CoverTab[114153]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1381
	// _ = "end of CoverTab[114148]"
}

func makeStdBoolValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1384
	_go_fuzz_dep_.CoverTab[114158]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1385
		_go_fuzz_dep_.CoverTab[114159]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1386
			_go_fuzz_dep_.CoverTab[114164]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1387
			// _ = "end of CoverTab[114164]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1388
			_go_fuzz_dep_.CoverTab[114165]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1388
			// _ = "end of CoverTab[114165]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1388
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1388
		// _ = "end of CoverTab[114159]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1388
		_go_fuzz_dep_.CoverTab[114160]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1390
			_go_fuzz_dep_.CoverTab[114166]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1391
			// _ = "end of CoverTab[114166]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1392
			_go_fuzz_dep_.CoverTab[114167]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1392
			// _ = "end of CoverTab[114167]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1392
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1392
		// _ = "end of CoverTab[114160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1392
		_go_fuzz_dep_.CoverTab[114161]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1394
			_go_fuzz_dep_.CoverTab[114168]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1395
			// _ = "end of CoverTab[114168]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1396
			_go_fuzz_dep_.CoverTab[114169]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1396
			// _ = "end of CoverTab[114169]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1396
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1396
		// _ = "end of CoverTab[114161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1396
		_go_fuzz_dep_.CoverTab[114162]++
													m := &boolValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1398
			_go_fuzz_dep_.CoverTab[114170]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1399
			// _ = "end of CoverTab[114170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1400
			_go_fuzz_dep_.CoverTab[114171]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1400
			// _ = "end of CoverTab[114171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1400
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1400
		// _ = "end of CoverTab[114162]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1400
		_go_fuzz_dep_.CoverTab[114163]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1403
		// _ = "end of CoverTab[114163]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1404
	// _ = "end of CoverTab[114158]"
}

func makeStdBoolValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1407
	_go_fuzz_dep_.CoverTab[114172]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1408
		_go_fuzz_dep_.CoverTab[114173]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1409
			_go_fuzz_dep_.CoverTab[114178]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1410
			// _ = "end of CoverTab[114178]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1411
			_go_fuzz_dep_.CoverTab[114179]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1411
			// _ = "end of CoverTab[114179]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1411
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1411
		// _ = "end of CoverTab[114173]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1411
		_go_fuzz_dep_.CoverTab[114174]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1413
			_go_fuzz_dep_.CoverTab[114180]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1414
			// _ = "end of CoverTab[114180]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1415
			_go_fuzz_dep_.CoverTab[114181]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1415
			// _ = "end of CoverTab[114181]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1415
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1415
		// _ = "end of CoverTab[114174]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1415
		_go_fuzz_dep_.CoverTab[114175]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1417
			_go_fuzz_dep_.CoverTab[114182]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1418
			// _ = "end of CoverTab[114182]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1419
			_go_fuzz_dep_.CoverTab[114183]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1419
			// _ = "end of CoverTab[114183]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1419
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1419
		// _ = "end of CoverTab[114175]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1419
		_go_fuzz_dep_.CoverTab[114176]++
													m := &boolValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1421
			_go_fuzz_dep_.CoverTab[114184]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1422
			// _ = "end of CoverTab[114184]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1423
			_go_fuzz_dep_.CoverTab[114185]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1423
			// _ = "end of CoverTab[114185]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1423
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1423
		// _ = "end of CoverTab[114176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1423
		_go_fuzz_dep_.CoverTab[114177]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1426
		// _ = "end of CoverTab[114177]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1427
	// _ = "end of CoverTab[114172]"
}

func makeStdBoolValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1430
	_go_fuzz_dep_.CoverTab[114186]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1431
		_go_fuzz_dep_.CoverTab[114187]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1432
			_go_fuzz_dep_.CoverTab[114192]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1433
			// _ = "end of CoverTab[114192]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1434
			_go_fuzz_dep_.CoverTab[114193]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1434
			// _ = "end of CoverTab[114193]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1434
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1434
		// _ = "end of CoverTab[114187]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1434
		_go_fuzz_dep_.CoverTab[114188]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1436
			_go_fuzz_dep_.CoverTab[114194]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1437
			// _ = "end of CoverTab[114194]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1438
			_go_fuzz_dep_.CoverTab[114195]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1438
			// _ = "end of CoverTab[114195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1438
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1438
		// _ = "end of CoverTab[114188]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1438
		_go_fuzz_dep_.CoverTab[114189]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1440
			_go_fuzz_dep_.CoverTab[114196]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1441
			// _ = "end of CoverTab[114196]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1442
			_go_fuzz_dep_.CoverTab[114197]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1442
			// _ = "end of CoverTab[114197]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1442
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1442
		// _ = "end of CoverTab[114189]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1442
		_go_fuzz_dep_.CoverTab[114190]++
													m := &boolValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1444
			_go_fuzz_dep_.CoverTab[114198]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1445
			// _ = "end of CoverTab[114198]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1446
			_go_fuzz_dep_.CoverTab[114199]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1446
			// _ = "end of CoverTab[114199]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1446
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1446
		// _ = "end of CoverTab[114190]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1446
		_go_fuzz_dep_.CoverTab[114191]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1450
		// _ = "end of CoverTab[114191]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1451
	// _ = "end of CoverTab[114186]"
}

func makeStdBoolValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1454
	_go_fuzz_dep_.CoverTab[114200]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1455
		_go_fuzz_dep_.CoverTab[114201]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1456
			_go_fuzz_dep_.CoverTab[114206]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1457
			// _ = "end of CoverTab[114206]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1458
			_go_fuzz_dep_.CoverTab[114207]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1458
			// _ = "end of CoverTab[114207]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1458
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1458
		// _ = "end of CoverTab[114201]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1458
		_go_fuzz_dep_.CoverTab[114202]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1460
			_go_fuzz_dep_.CoverTab[114208]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1461
			// _ = "end of CoverTab[114208]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1462
			_go_fuzz_dep_.CoverTab[114209]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1462
			// _ = "end of CoverTab[114209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1462
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1462
		// _ = "end of CoverTab[114202]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1462
		_go_fuzz_dep_.CoverTab[114203]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1464
			_go_fuzz_dep_.CoverTab[114210]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1465
			// _ = "end of CoverTab[114210]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1466
			_go_fuzz_dep_.CoverTab[114211]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1466
			// _ = "end of CoverTab[114211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1466
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1466
		// _ = "end of CoverTab[114203]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1466
		_go_fuzz_dep_.CoverTab[114204]++
													m := &boolValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1468
			_go_fuzz_dep_.CoverTab[114212]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1469
			// _ = "end of CoverTab[114212]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1470
			_go_fuzz_dep_.CoverTab[114213]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1470
			// _ = "end of CoverTab[114213]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1470
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1470
		// _ = "end of CoverTab[114204]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1470
		_go_fuzz_dep_.CoverTab[114205]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1474
		// _ = "end of CoverTab[114205]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1475
	// _ = "end of CoverTab[114200]"
}

func makeStdStringValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1478
	_go_fuzz_dep_.CoverTab[114214]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1479
			_go_fuzz_dep_.CoverTab[114215]++
														t := ptr.asPointerTo(u.typ).Interface().(*string)
														v := &stringValue{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1483
			// _ = "end of CoverTab[114215]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1484
			_go_fuzz_dep_.CoverTab[114216]++
														t := ptr.asPointerTo(u.typ).Interface().(*string)
														v := &stringValue{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1488
				_go_fuzz_dep_.CoverTab[114218]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1489
				// _ = "end of CoverTab[114218]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1490
				_go_fuzz_dep_.CoverTab[114219]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1490
				// _ = "end of CoverTab[114219]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1490
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1490
			// _ = "end of CoverTab[114216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1490
			_go_fuzz_dep_.CoverTab[114217]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1494
			// _ = "end of CoverTab[114217]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1495
	// _ = "end of CoverTab[114214]"
}

func makeStdStringValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1498
	_go_fuzz_dep_.CoverTab[114220]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1499
			_go_fuzz_dep_.CoverTab[114221]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1500
				_go_fuzz_dep_.CoverTab[114223]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1501
				// _ = "end of CoverTab[114223]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1502
				_go_fuzz_dep_.CoverTab[114224]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1502
				// _ = "end of CoverTab[114224]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1502
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1502
			// _ = "end of CoverTab[114221]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1502
			_go_fuzz_dep_.CoverTab[114222]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*string)
														v := &stringValue{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1506
			// _ = "end of CoverTab[114222]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1507
			_go_fuzz_dep_.CoverTab[114225]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1508
				_go_fuzz_dep_.CoverTab[114228]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1509
				// _ = "end of CoverTab[114228]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1510
				_go_fuzz_dep_.CoverTab[114229]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1510
				// _ = "end of CoverTab[114229]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1510
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1510
			// _ = "end of CoverTab[114225]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1510
			_go_fuzz_dep_.CoverTab[114226]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*string)
														v := &stringValue{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1514
				_go_fuzz_dep_.CoverTab[114230]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1515
				// _ = "end of CoverTab[114230]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1516
				_go_fuzz_dep_.CoverTab[114231]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1516
				// _ = "end of CoverTab[114231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1516
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1516
			// _ = "end of CoverTab[114226]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1516
			_go_fuzz_dep_.CoverTab[114227]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1520
			// _ = "end of CoverTab[114227]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1521
	// _ = "end of CoverTab[114220]"
}

func makeStdStringValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1524
	_go_fuzz_dep_.CoverTab[114232]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1525
			_go_fuzz_dep_.CoverTab[114233]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1528
				_go_fuzz_dep_.CoverTab[114235]++
															elem := s.Index(i)
															t := elem.Interface().(string)
															v := &stringValue{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1533
				// _ = "end of CoverTab[114235]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1534
			// _ = "end of CoverTab[114233]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1534
			_go_fuzz_dep_.CoverTab[114234]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1535
			// _ = "end of CoverTab[114234]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1537
			_go_fuzz_dep_.CoverTab[114236]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1539
				_go_fuzz_dep_.CoverTab[114238]++
															elem := s.Index(i)
															t := elem.Interface().(string)
															v := &stringValue{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1545
					_go_fuzz_dep_.CoverTab[114240]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1546
					// _ = "end of CoverTab[114240]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1547
					_go_fuzz_dep_.CoverTab[114241]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1547
					// _ = "end of CoverTab[114241]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1547
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1547
				// _ = "end of CoverTab[114238]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1547
				_go_fuzz_dep_.CoverTab[114239]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1550
				// _ = "end of CoverTab[114239]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1551
			// _ = "end of CoverTab[114236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1551
			_go_fuzz_dep_.CoverTab[114237]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1553
			// _ = "end of CoverTab[114237]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1554
	// _ = "end of CoverTab[114232]"
}

func makeStdStringValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1557
	_go_fuzz_dep_.CoverTab[114242]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1558
			_go_fuzz_dep_.CoverTab[114243]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1561
				_go_fuzz_dep_.CoverTab[114245]++
															elem := s.Index(i)
															t := elem.Interface().(*string)
															v := &stringValue{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1566
				// _ = "end of CoverTab[114245]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1567
			// _ = "end of CoverTab[114243]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1567
			_go_fuzz_dep_.CoverTab[114244]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1568
			// _ = "end of CoverTab[114244]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1570
			_go_fuzz_dep_.CoverTab[114246]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1572
				_go_fuzz_dep_.CoverTab[114248]++
															elem := s.Index(i)
															t := elem.Interface().(*string)
															v := &stringValue{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1578
					_go_fuzz_dep_.CoverTab[114250]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1579
					// _ = "end of CoverTab[114250]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1580
					_go_fuzz_dep_.CoverTab[114251]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1580
					// _ = "end of CoverTab[114251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1580
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1580
				// _ = "end of CoverTab[114248]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1580
				_go_fuzz_dep_.CoverTab[114249]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1583
				// _ = "end of CoverTab[114249]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1584
			// _ = "end of CoverTab[114246]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1584
			_go_fuzz_dep_.CoverTab[114247]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1586
			// _ = "end of CoverTab[114247]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1587
	// _ = "end of CoverTab[114242]"
}

func makeStdStringValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1590
	_go_fuzz_dep_.CoverTab[114252]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1591
		_go_fuzz_dep_.CoverTab[114253]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1592
			_go_fuzz_dep_.CoverTab[114258]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1593
			// _ = "end of CoverTab[114258]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1594
			_go_fuzz_dep_.CoverTab[114259]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1594
			// _ = "end of CoverTab[114259]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1594
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1594
		// _ = "end of CoverTab[114253]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1594
		_go_fuzz_dep_.CoverTab[114254]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1596
			_go_fuzz_dep_.CoverTab[114260]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1597
			// _ = "end of CoverTab[114260]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1598
			_go_fuzz_dep_.CoverTab[114261]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1598
			// _ = "end of CoverTab[114261]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1598
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1598
		// _ = "end of CoverTab[114254]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1598
		_go_fuzz_dep_.CoverTab[114255]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1600
			_go_fuzz_dep_.CoverTab[114262]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1601
			// _ = "end of CoverTab[114262]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1602
			_go_fuzz_dep_.CoverTab[114263]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1602
			// _ = "end of CoverTab[114263]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1602
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1602
		// _ = "end of CoverTab[114255]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1602
		_go_fuzz_dep_.CoverTab[114256]++
													m := &stringValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1604
			_go_fuzz_dep_.CoverTab[114264]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1605
			// _ = "end of CoverTab[114264]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1606
			_go_fuzz_dep_.CoverTab[114265]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1606
			// _ = "end of CoverTab[114265]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1606
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1606
		// _ = "end of CoverTab[114256]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1606
		_go_fuzz_dep_.CoverTab[114257]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1609
		// _ = "end of CoverTab[114257]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1610
	// _ = "end of CoverTab[114252]"
}

func makeStdStringValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1613
	_go_fuzz_dep_.CoverTab[114266]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1614
		_go_fuzz_dep_.CoverTab[114267]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1615
			_go_fuzz_dep_.CoverTab[114272]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1616
			// _ = "end of CoverTab[114272]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1617
			_go_fuzz_dep_.CoverTab[114273]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1617
			// _ = "end of CoverTab[114273]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1617
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1617
		// _ = "end of CoverTab[114267]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1617
		_go_fuzz_dep_.CoverTab[114268]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1619
			_go_fuzz_dep_.CoverTab[114274]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1620
			// _ = "end of CoverTab[114274]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1621
			_go_fuzz_dep_.CoverTab[114275]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1621
			// _ = "end of CoverTab[114275]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1621
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1621
		// _ = "end of CoverTab[114268]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1621
		_go_fuzz_dep_.CoverTab[114269]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1623
			_go_fuzz_dep_.CoverTab[114276]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1624
			// _ = "end of CoverTab[114276]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1625
			_go_fuzz_dep_.CoverTab[114277]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1625
			// _ = "end of CoverTab[114277]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1625
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1625
		// _ = "end of CoverTab[114269]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1625
		_go_fuzz_dep_.CoverTab[114270]++
													m := &stringValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1627
			_go_fuzz_dep_.CoverTab[114278]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1628
			// _ = "end of CoverTab[114278]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1629
			_go_fuzz_dep_.CoverTab[114279]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1629
			// _ = "end of CoverTab[114279]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1629
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1629
		// _ = "end of CoverTab[114270]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1629
		_go_fuzz_dep_.CoverTab[114271]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1632
		// _ = "end of CoverTab[114271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1633
	// _ = "end of CoverTab[114266]"
}

func makeStdStringValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1636
	_go_fuzz_dep_.CoverTab[114280]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1637
		_go_fuzz_dep_.CoverTab[114281]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1638
			_go_fuzz_dep_.CoverTab[114286]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1639
			// _ = "end of CoverTab[114286]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1640
			_go_fuzz_dep_.CoverTab[114287]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1640
			// _ = "end of CoverTab[114287]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1640
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1640
		// _ = "end of CoverTab[114281]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1640
		_go_fuzz_dep_.CoverTab[114282]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1642
			_go_fuzz_dep_.CoverTab[114288]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1643
			// _ = "end of CoverTab[114288]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1644
			_go_fuzz_dep_.CoverTab[114289]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1644
			// _ = "end of CoverTab[114289]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1644
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1644
		// _ = "end of CoverTab[114282]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1644
		_go_fuzz_dep_.CoverTab[114283]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1646
			_go_fuzz_dep_.CoverTab[114290]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1647
			// _ = "end of CoverTab[114290]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1648
			_go_fuzz_dep_.CoverTab[114291]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1648
			// _ = "end of CoverTab[114291]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1648
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1648
		// _ = "end of CoverTab[114283]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1648
		_go_fuzz_dep_.CoverTab[114284]++
													m := &stringValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1650
			_go_fuzz_dep_.CoverTab[114292]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1651
			// _ = "end of CoverTab[114292]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1652
			_go_fuzz_dep_.CoverTab[114293]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1652
			// _ = "end of CoverTab[114293]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1652
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1652
		// _ = "end of CoverTab[114284]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1652
		_go_fuzz_dep_.CoverTab[114285]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1656
		// _ = "end of CoverTab[114285]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1657
	// _ = "end of CoverTab[114280]"
}

func makeStdStringValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1660
	_go_fuzz_dep_.CoverTab[114294]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1661
		_go_fuzz_dep_.CoverTab[114295]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1662
			_go_fuzz_dep_.CoverTab[114300]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1663
			// _ = "end of CoverTab[114300]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1664
			_go_fuzz_dep_.CoverTab[114301]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1664
			// _ = "end of CoverTab[114301]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1664
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1664
		// _ = "end of CoverTab[114295]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1664
		_go_fuzz_dep_.CoverTab[114296]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1666
			_go_fuzz_dep_.CoverTab[114302]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1667
			// _ = "end of CoverTab[114302]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1668
			_go_fuzz_dep_.CoverTab[114303]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1668
			// _ = "end of CoverTab[114303]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1668
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1668
		// _ = "end of CoverTab[114296]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1668
		_go_fuzz_dep_.CoverTab[114297]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1670
			_go_fuzz_dep_.CoverTab[114304]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1671
			// _ = "end of CoverTab[114304]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1672
			_go_fuzz_dep_.CoverTab[114305]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1672
			// _ = "end of CoverTab[114305]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1672
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1672
		// _ = "end of CoverTab[114297]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1672
		_go_fuzz_dep_.CoverTab[114298]++
													m := &stringValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1674
			_go_fuzz_dep_.CoverTab[114306]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1675
			// _ = "end of CoverTab[114306]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1676
			_go_fuzz_dep_.CoverTab[114307]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1676
			// _ = "end of CoverTab[114307]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1676
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1676
		// _ = "end of CoverTab[114298]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1676
		_go_fuzz_dep_.CoverTab[114299]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1680
		// _ = "end of CoverTab[114299]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1681
	// _ = "end of CoverTab[114294]"
}

func makeStdBytesValueMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1684
	_go_fuzz_dep_.CoverTab[114308]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1685
			_go_fuzz_dep_.CoverTab[114309]++
														t := ptr.asPointerTo(u.typ).Interface().(*[]byte)
														v := &bytesValue{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1689
			// _ = "end of CoverTab[114309]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1690
			_go_fuzz_dep_.CoverTab[114310]++
														t := ptr.asPointerTo(u.typ).Interface().(*[]byte)
														v := &bytesValue{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1694
				_go_fuzz_dep_.CoverTab[114312]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1695
				// _ = "end of CoverTab[114312]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1696
				_go_fuzz_dep_.CoverTab[114313]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1696
				// _ = "end of CoverTab[114313]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1696
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1696
			// _ = "end of CoverTab[114310]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1696
			_go_fuzz_dep_.CoverTab[114311]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1700
			// _ = "end of CoverTab[114311]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1701
	// _ = "end of CoverTab[114308]"
}

func makeStdBytesValuePtrMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1704
	_go_fuzz_dep_.CoverTab[114314]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1705
			_go_fuzz_dep_.CoverTab[114315]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1706
				_go_fuzz_dep_.CoverTab[114317]++
															return 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1707
				// _ = "end of CoverTab[114317]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1708
				_go_fuzz_dep_.CoverTab[114318]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1708
				// _ = "end of CoverTab[114318]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1708
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1708
			// _ = "end of CoverTab[114315]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1708
			_go_fuzz_dep_.CoverTab[114316]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*[]byte)
														v := &bytesValue{*t}
														siz := Size(v)
														return tagsize + SizeVarint(uint64(siz)) + siz
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1712
			// _ = "end of CoverTab[114316]"
		}, func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1713
			_go_fuzz_dep_.CoverTab[114319]++
														if ptr.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1714
				_go_fuzz_dep_.CoverTab[114322]++
															return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1715
				// _ = "end of CoverTab[114322]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1716
				_go_fuzz_dep_.CoverTab[114323]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1716
				// _ = "end of CoverTab[114323]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1716
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1716
			// _ = "end of CoverTab[114319]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1716
			_go_fuzz_dep_.CoverTab[114320]++
														t := ptr.asPointerTo(reflect.PtrTo(u.typ)).Elem().Interface().(*[]byte)
														v := &bytesValue{*t}
														buf, err := Marshal(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1720
				_go_fuzz_dep_.CoverTab[114324]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1721
				// _ = "end of CoverTab[114324]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1722
				_go_fuzz_dep_.CoverTab[114325]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1722
				// _ = "end of CoverTab[114325]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1722
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1722
			// _ = "end of CoverTab[114320]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1722
			_go_fuzz_dep_.CoverTab[114321]++
														b = appendVarint(b, wiretag)
														b = appendVarint(b, uint64(len(buf)))
														b = append(b, buf...)
														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1726
			// _ = "end of CoverTab[114321]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1727
	// _ = "end of CoverTab[114314]"
}

func makeStdBytesValueSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1730
	_go_fuzz_dep_.CoverTab[114326]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1731
			_go_fuzz_dep_.CoverTab[114327]++
														s := ptr.getSlice(u.typ)
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1734
				_go_fuzz_dep_.CoverTab[114329]++
															elem := s.Index(i)
															t := elem.Interface().([]byte)
															v := &bytesValue{t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1739
				// _ = "end of CoverTab[114329]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1740
			// _ = "end of CoverTab[114327]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1740
			_go_fuzz_dep_.CoverTab[114328]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1741
			// _ = "end of CoverTab[114328]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1743
			_go_fuzz_dep_.CoverTab[114330]++
														s := ptr.getSlice(u.typ)
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1745
				_go_fuzz_dep_.CoverTab[114332]++
															elem := s.Index(i)
															t := elem.Interface().([]byte)
															v := &bytesValue{t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1751
					_go_fuzz_dep_.CoverTab[114334]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1752
					// _ = "end of CoverTab[114334]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1753
					_go_fuzz_dep_.CoverTab[114335]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1753
					// _ = "end of CoverTab[114335]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1753
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1753
				// _ = "end of CoverTab[114332]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1753
				_go_fuzz_dep_.CoverTab[114333]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1756
				// _ = "end of CoverTab[114333]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1757
			// _ = "end of CoverTab[114330]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1757
			_go_fuzz_dep_.CoverTab[114331]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1759
			// _ = "end of CoverTab[114331]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1760
	// _ = "end of CoverTab[114326]"
}

func makeStdBytesValuePtrSliceMarshaler(u *marshalInfo) (sizer, marshaler) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1763
	_go_fuzz_dep_.CoverTab[114336]++
												return func(ptr pointer, tagsize int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1764
			_go_fuzz_dep_.CoverTab[114337]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														n := 0
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1767
				_go_fuzz_dep_.CoverTab[114339]++
															elem := s.Index(i)
															t := elem.Interface().(*[]byte)
															v := &bytesValue{*t}
															siz := Size(v)
															n += siz + SizeVarint(uint64(siz)) + tagsize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1772
				// _ = "end of CoverTab[114339]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1773
			// _ = "end of CoverTab[114337]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1773
			_go_fuzz_dep_.CoverTab[114338]++
														return n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1774
			// _ = "end of CoverTab[114338]"
		},
		func(b []byte, ptr pointer, wiretag uint64, deterministic bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1776
			_go_fuzz_dep_.CoverTab[114340]++
														s := ptr.getSlice(reflect.PtrTo(u.typ))
														for i := 0; i < s.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1778
				_go_fuzz_dep_.CoverTab[114342]++
															elem := s.Index(i)
															t := elem.Interface().(*[]byte)
															v := &bytesValue{*t}
															siz := Size(v)
															buf, err := Marshal(v)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1784
					_go_fuzz_dep_.CoverTab[114344]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1785
					// _ = "end of CoverTab[114344]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1786
					_go_fuzz_dep_.CoverTab[114345]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1786
					// _ = "end of CoverTab[114345]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1786
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1786
				// _ = "end of CoverTab[114342]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1786
				_go_fuzz_dep_.CoverTab[114343]++
															b = appendVarint(b, wiretag)
															b = appendVarint(b, uint64(siz))
															b = append(b, buf...)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1789
				// _ = "end of CoverTab[114343]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1790
			// _ = "end of CoverTab[114340]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1790
			_go_fuzz_dep_.CoverTab[114341]++

														return b, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1792
			// _ = "end of CoverTab[114341]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1793
	// _ = "end of CoverTab[114336]"
}

func makeStdBytesValueUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1796
	_go_fuzz_dep_.CoverTab[114346]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1797
		_go_fuzz_dep_.CoverTab[114347]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1798
			_go_fuzz_dep_.CoverTab[114352]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1799
			// _ = "end of CoverTab[114352]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1800
			_go_fuzz_dep_.CoverTab[114353]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1800
			// _ = "end of CoverTab[114353]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1800
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1800
		// _ = "end of CoverTab[114347]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1800
		_go_fuzz_dep_.CoverTab[114348]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1802
			_go_fuzz_dep_.CoverTab[114354]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1803
			// _ = "end of CoverTab[114354]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1804
			_go_fuzz_dep_.CoverTab[114355]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1804
			// _ = "end of CoverTab[114355]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1804
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1804
		// _ = "end of CoverTab[114348]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1804
		_go_fuzz_dep_.CoverTab[114349]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1806
			_go_fuzz_dep_.CoverTab[114356]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1807
			// _ = "end of CoverTab[114356]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1808
			_go_fuzz_dep_.CoverTab[114357]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1808
			// _ = "end of CoverTab[114357]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1808
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1808
		// _ = "end of CoverTab[114349]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1808
		_go_fuzz_dep_.CoverTab[114350]++
													m := &bytesValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1810
			_go_fuzz_dep_.CoverTab[114358]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1811
			// _ = "end of CoverTab[114358]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1812
			_go_fuzz_dep_.CoverTab[114359]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1812
			// _ = "end of CoverTab[114359]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1812
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1812
		// _ = "end of CoverTab[114350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1812
		_go_fuzz_dep_.CoverTab[114351]++
													s := f.asPointerTo(sub.typ).Elem()
													s.Set(reflect.ValueOf(m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1815
		// _ = "end of CoverTab[114351]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1816
	// _ = "end of CoverTab[114346]"
}

func makeStdBytesValuePtrUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1819
	_go_fuzz_dep_.CoverTab[114360]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1820
		_go_fuzz_dep_.CoverTab[114361]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1821
			_go_fuzz_dep_.CoverTab[114366]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1822
			// _ = "end of CoverTab[114366]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1823
			_go_fuzz_dep_.CoverTab[114367]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1823
			// _ = "end of CoverTab[114367]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1823
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1823
		// _ = "end of CoverTab[114361]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1823
		_go_fuzz_dep_.CoverTab[114362]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1825
			_go_fuzz_dep_.CoverTab[114368]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1826
			// _ = "end of CoverTab[114368]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1827
			_go_fuzz_dep_.CoverTab[114369]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1827
			// _ = "end of CoverTab[114369]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1827
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1827
		// _ = "end of CoverTab[114362]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1827
		_go_fuzz_dep_.CoverTab[114363]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1829
			_go_fuzz_dep_.CoverTab[114370]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1830
			// _ = "end of CoverTab[114370]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1831
			_go_fuzz_dep_.CoverTab[114371]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1831
			// _ = "end of CoverTab[114371]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1831
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1831
		// _ = "end of CoverTab[114363]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1831
		_go_fuzz_dep_.CoverTab[114364]++
													m := &bytesValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1833
			_go_fuzz_dep_.CoverTab[114372]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1834
			// _ = "end of CoverTab[114372]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1835
			_go_fuzz_dep_.CoverTab[114373]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1835
			// _ = "end of CoverTab[114373]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1835
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1835
		// _ = "end of CoverTab[114364]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1835
		_go_fuzz_dep_.CoverTab[114365]++
													s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
													s.Set(reflect.ValueOf(&m.Value))
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1838
		// _ = "end of CoverTab[114365]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1839
	// _ = "end of CoverTab[114360]"
}

func makeStdBytesValuePtrSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1842
	_go_fuzz_dep_.CoverTab[114374]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1843
		_go_fuzz_dep_.CoverTab[114375]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1844
			_go_fuzz_dep_.CoverTab[114380]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1845
			// _ = "end of CoverTab[114380]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1846
			_go_fuzz_dep_.CoverTab[114381]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1846
			// _ = "end of CoverTab[114381]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1846
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1846
		// _ = "end of CoverTab[114375]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1846
		_go_fuzz_dep_.CoverTab[114376]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1848
			_go_fuzz_dep_.CoverTab[114382]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1849
			// _ = "end of CoverTab[114382]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1850
			_go_fuzz_dep_.CoverTab[114383]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1850
			// _ = "end of CoverTab[114383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1850
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1850
		// _ = "end of CoverTab[114376]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1850
		_go_fuzz_dep_.CoverTab[114377]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1852
			_go_fuzz_dep_.CoverTab[114384]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1853
			// _ = "end of CoverTab[114384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1854
			_go_fuzz_dep_.CoverTab[114385]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1854
			// _ = "end of CoverTab[114385]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1854
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1854
		// _ = "end of CoverTab[114377]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1854
		_go_fuzz_dep_.CoverTab[114378]++
													m := &bytesValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1856
			_go_fuzz_dep_.CoverTab[114386]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1857
			// _ = "end of CoverTab[114386]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1858
			_go_fuzz_dep_.CoverTab[114387]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1858
			// _ = "end of CoverTab[114387]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1858
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1858
		// _ = "end of CoverTab[114378]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1858
		_go_fuzz_dep_.CoverTab[114379]++
													slice := f.getSlice(reflect.PtrTo(sub.typ))
													newSlice := reflect.Append(slice, reflect.ValueOf(&m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1862
		// _ = "end of CoverTab[114379]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1863
	// _ = "end of CoverTab[114374]"
}

func makeStdBytesValueSliceUnmarshaler(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1866
	_go_fuzz_dep_.CoverTab[114388]++
												return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1867
		_go_fuzz_dep_.CoverTab[114389]++
													if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1868
			_go_fuzz_dep_.CoverTab[114394]++
														return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1869
			// _ = "end of CoverTab[114394]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1870
			_go_fuzz_dep_.CoverTab[114395]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1870
			// _ = "end of CoverTab[114395]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1870
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1870
		// _ = "end of CoverTab[114389]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1870
		_go_fuzz_dep_.CoverTab[114390]++
													x, n := decodeVarint(b)
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1872
			_go_fuzz_dep_.CoverTab[114396]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1873
			// _ = "end of CoverTab[114396]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1874
			_go_fuzz_dep_.CoverTab[114397]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1874
			// _ = "end of CoverTab[114397]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1874
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1874
		// _ = "end of CoverTab[114390]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1874
		_go_fuzz_dep_.CoverTab[114391]++
													b = b[n:]
													if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1876
			_go_fuzz_dep_.CoverTab[114398]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1877
			// _ = "end of CoverTab[114398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1878
			_go_fuzz_dep_.CoverTab[114399]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1878
			// _ = "end of CoverTab[114399]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1878
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1878
		// _ = "end of CoverTab[114391]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1878
		_go_fuzz_dep_.CoverTab[114392]++
													m := &bytesValue{}
													if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1880
			_go_fuzz_dep_.CoverTab[114400]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1881
			// _ = "end of CoverTab[114400]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1882
			_go_fuzz_dep_.CoverTab[114401]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1882
			// _ = "end of CoverTab[114401]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1882
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1882
		// _ = "end of CoverTab[114392]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1882
		_go_fuzz_dep_.CoverTab[114393]++
													slice := f.getSlice(sub.typ)
													newSlice := reflect.Append(slice, reflect.ValueOf(m.Value))
													slice.Set(newSlice)
													return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1886
		// _ = "end of CoverTab[114393]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1887
	// _ = "end of CoverTab[114388]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1888
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/wrappers.go:1888
var _ = _go_fuzz_dep_.CoverTab
