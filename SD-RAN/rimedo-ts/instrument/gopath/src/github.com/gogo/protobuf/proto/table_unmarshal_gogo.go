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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:29
)

import (
	"io"
	"reflect"
)

func makeUnmarshalMessage(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:36
	_go_fuzz_dep_.CoverTab[112260]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:37
		_go_fuzz_dep_.CoverTab[112261]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:38
			_go_fuzz_dep_.CoverTab[112267]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:39
			// _ = "end of CoverTab[112267]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:40
			_go_fuzz_dep_.CoverTab[112268]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:40
			// _ = "end of CoverTab[112268]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:40
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:40
		// _ = "end of CoverTab[112261]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:40
		_go_fuzz_dep_.CoverTab[112262]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:42
			_go_fuzz_dep_.CoverTab[112269]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:43
			// _ = "end of CoverTab[112269]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:44
			_go_fuzz_dep_.CoverTab[112270]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:44
			// _ = "end of CoverTab[112270]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:44
		// _ = "end of CoverTab[112262]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:44
		_go_fuzz_dep_.CoverTab[112263]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:46
			_go_fuzz_dep_.CoverTab[112271]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:47
			// _ = "end of CoverTab[112271]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:48
			_go_fuzz_dep_.CoverTab[112272]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:48
			// _ = "end of CoverTab[112272]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:48
		// _ = "end of CoverTab[112263]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:48
		_go_fuzz_dep_.CoverTab[112264]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:53
		v := f
		if v.isNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:54
			_go_fuzz_dep_.CoverTab[112273]++
															v = valToPointer(reflect.New(sub.typ))
															f.setPointer(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:56
			// _ = "end of CoverTab[112273]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:57
			_go_fuzz_dep_.CoverTab[112274]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:57
			// _ = "end of CoverTab[112274]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:57
		// _ = "end of CoverTab[112264]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:57
		_go_fuzz_dep_.CoverTab[112265]++
														err := sub.unmarshal(v, b[:x])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:59
			_go_fuzz_dep_.CoverTab[112275]++
															if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:60
				_go_fuzz_dep_.CoverTab[112276]++
																r.field = name + "." + r.field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:61
				// _ = "end of CoverTab[112276]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:62
				_go_fuzz_dep_.CoverTab[112277]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:63
				// _ = "end of CoverTab[112277]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:64
			// _ = "end of CoverTab[112275]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:65
			_go_fuzz_dep_.CoverTab[112278]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:65
			// _ = "end of CoverTab[112278]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:65
		// _ = "end of CoverTab[112265]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:65
		_go_fuzz_dep_.CoverTab[112266]++
														return b[x:], err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:66
		// _ = "end of CoverTab[112266]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:67
	// _ = "end of CoverTab[112260]"
}

func makeUnmarshalMessageSlice(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:70
	_go_fuzz_dep_.CoverTab[112279]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:71
		_go_fuzz_dep_.CoverTab[112280]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:72
			_go_fuzz_dep_.CoverTab[112285]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:73
			// _ = "end of CoverTab[112285]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:74
			_go_fuzz_dep_.CoverTab[112286]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:74
			// _ = "end of CoverTab[112286]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:74
		// _ = "end of CoverTab[112280]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:74
		_go_fuzz_dep_.CoverTab[112281]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:76
			_go_fuzz_dep_.CoverTab[112287]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:77
			// _ = "end of CoverTab[112287]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:78
			_go_fuzz_dep_.CoverTab[112288]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:78
			// _ = "end of CoverTab[112288]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:78
		// _ = "end of CoverTab[112281]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:78
		_go_fuzz_dep_.CoverTab[112282]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:80
			_go_fuzz_dep_.CoverTab[112289]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:81
			// _ = "end of CoverTab[112289]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:82
			_go_fuzz_dep_.CoverTab[112290]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:82
			// _ = "end of CoverTab[112290]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:82
		// _ = "end of CoverTab[112282]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:82
		_go_fuzz_dep_.CoverTab[112283]++
														v := valToPointer(reflect.New(sub.typ))
														err := sub.unmarshal(v, b[:x])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:85
			_go_fuzz_dep_.CoverTab[112291]++
															if r, ok := err.(*RequiredNotSetError); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:86
				_go_fuzz_dep_.CoverTab[112292]++
																r.field = name + "." + r.field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:87
				// _ = "end of CoverTab[112292]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:88
				_go_fuzz_dep_.CoverTab[112293]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:89
				// _ = "end of CoverTab[112293]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:90
			// _ = "end of CoverTab[112291]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:91
			_go_fuzz_dep_.CoverTab[112294]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:91
			// _ = "end of CoverTab[112294]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:91
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:91
		// _ = "end of CoverTab[112283]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:91
		_go_fuzz_dep_.CoverTab[112284]++
														f.appendRef(v, sub.typ)
														return b[x:], err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:93
		// _ = "end of CoverTab[112284]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:94
	// _ = "end of CoverTab[112279]"
}

func makeUnmarshalCustomPtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:97
	_go_fuzz_dep_.CoverTab[112295]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:98
		_go_fuzz_dep_.CoverTab[112296]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:99
			_go_fuzz_dep_.CoverTab[112301]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:100
			// _ = "end of CoverTab[112301]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:101
			_go_fuzz_dep_.CoverTab[112302]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:101
			// _ = "end of CoverTab[112302]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:101
		// _ = "end of CoverTab[112296]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:101
		_go_fuzz_dep_.CoverTab[112297]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:103
			_go_fuzz_dep_.CoverTab[112303]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:104
			// _ = "end of CoverTab[112303]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:105
			_go_fuzz_dep_.CoverTab[112304]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:105
			// _ = "end of CoverTab[112304]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:105
		// _ = "end of CoverTab[112297]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:105
		_go_fuzz_dep_.CoverTab[112298]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:107
			_go_fuzz_dep_.CoverTab[112305]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:108
			// _ = "end of CoverTab[112305]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:109
			_go_fuzz_dep_.CoverTab[112306]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:109
			// _ = "end of CoverTab[112306]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:109
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:109
		// _ = "end of CoverTab[112298]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:109
		_go_fuzz_dep_.CoverTab[112299]++

														s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
														s.Set(reflect.New(sub.typ))
														m := s.Interface().(custom)
														if err := m.Unmarshal(b[:x]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:114
			_go_fuzz_dep_.CoverTab[112307]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:115
			// _ = "end of CoverTab[112307]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:116
			_go_fuzz_dep_.CoverTab[112308]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:116
			// _ = "end of CoverTab[112308]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:116
		// _ = "end of CoverTab[112299]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:116
		_go_fuzz_dep_.CoverTab[112300]++
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:117
		// _ = "end of CoverTab[112300]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:118
	// _ = "end of CoverTab[112295]"
}

func makeUnmarshalCustomSlice(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:121
	_go_fuzz_dep_.CoverTab[112309]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:122
		_go_fuzz_dep_.CoverTab[112310]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:123
			_go_fuzz_dep_.CoverTab[112315]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:124
			// _ = "end of CoverTab[112315]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:125
			_go_fuzz_dep_.CoverTab[112316]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:125
			// _ = "end of CoverTab[112316]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:125
		// _ = "end of CoverTab[112310]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:125
		_go_fuzz_dep_.CoverTab[112311]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:127
			_go_fuzz_dep_.CoverTab[112317]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:128
			// _ = "end of CoverTab[112317]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:129
			_go_fuzz_dep_.CoverTab[112318]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:129
			// _ = "end of CoverTab[112318]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:129
		// _ = "end of CoverTab[112311]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:129
		_go_fuzz_dep_.CoverTab[112312]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:131
			_go_fuzz_dep_.CoverTab[112319]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:132
			// _ = "end of CoverTab[112319]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:133
			_go_fuzz_dep_.CoverTab[112320]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:133
			// _ = "end of CoverTab[112320]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:133
		// _ = "end of CoverTab[112312]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:133
		_go_fuzz_dep_.CoverTab[112313]++
														m := reflect.New(sub.typ)
														c := m.Interface().(custom)
														if err := c.Unmarshal(b[:x]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:136
			_go_fuzz_dep_.CoverTab[112321]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:137
			// _ = "end of CoverTab[112321]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:138
			_go_fuzz_dep_.CoverTab[112322]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:138
			// _ = "end of CoverTab[112322]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:138
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:138
		// _ = "end of CoverTab[112313]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:138
		_go_fuzz_dep_.CoverTab[112314]++
														v := valToPointer(m)
														f.appendRef(v, sub.typ)
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:141
		// _ = "end of CoverTab[112314]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:142
	// _ = "end of CoverTab[112309]"
}

func makeUnmarshalCustom(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:145
	_go_fuzz_dep_.CoverTab[112323]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:146
		_go_fuzz_dep_.CoverTab[112324]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:147
			_go_fuzz_dep_.CoverTab[112329]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:148
			// _ = "end of CoverTab[112329]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:149
			_go_fuzz_dep_.CoverTab[112330]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:149
			// _ = "end of CoverTab[112330]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:149
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:149
		// _ = "end of CoverTab[112324]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:149
		_go_fuzz_dep_.CoverTab[112325]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:151
			_go_fuzz_dep_.CoverTab[112331]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:152
			// _ = "end of CoverTab[112331]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:153
			_go_fuzz_dep_.CoverTab[112332]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:153
			// _ = "end of CoverTab[112332]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:153
		// _ = "end of CoverTab[112325]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:153
		_go_fuzz_dep_.CoverTab[112326]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:155
			_go_fuzz_dep_.CoverTab[112333]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:156
			// _ = "end of CoverTab[112333]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:157
			_go_fuzz_dep_.CoverTab[112334]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:157
			// _ = "end of CoverTab[112334]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:157
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:157
		// _ = "end of CoverTab[112326]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:157
		_go_fuzz_dep_.CoverTab[112327]++

														m := f.asPointerTo(sub.typ).Interface().(custom)
														if err := m.Unmarshal(b[:x]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:160
			_go_fuzz_dep_.CoverTab[112335]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:161
			// _ = "end of CoverTab[112335]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:162
			_go_fuzz_dep_.CoverTab[112336]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:162
			// _ = "end of CoverTab[112336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:162
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:162
		// _ = "end of CoverTab[112327]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:162
		_go_fuzz_dep_.CoverTab[112328]++
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:163
		// _ = "end of CoverTab[112328]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:164
	// _ = "end of CoverTab[112323]"
}

func makeUnmarshalTime(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:167
	_go_fuzz_dep_.CoverTab[112337]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:168
		_go_fuzz_dep_.CoverTab[112338]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:169
			_go_fuzz_dep_.CoverTab[112344]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:170
			// _ = "end of CoverTab[112344]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:171
			_go_fuzz_dep_.CoverTab[112345]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:171
			// _ = "end of CoverTab[112345]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:171
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:171
		// _ = "end of CoverTab[112338]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:171
		_go_fuzz_dep_.CoverTab[112339]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:173
			_go_fuzz_dep_.CoverTab[112346]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:174
			// _ = "end of CoverTab[112346]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:175
			_go_fuzz_dep_.CoverTab[112347]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:175
			// _ = "end of CoverTab[112347]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:175
		// _ = "end of CoverTab[112339]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:175
		_go_fuzz_dep_.CoverTab[112340]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:177
			_go_fuzz_dep_.CoverTab[112348]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:178
			// _ = "end of CoverTab[112348]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:179
			_go_fuzz_dep_.CoverTab[112349]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:179
			// _ = "end of CoverTab[112349]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:179
		// _ = "end of CoverTab[112340]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:179
		_go_fuzz_dep_.CoverTab[112341]++
														m := &timestamp{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:181
			_go_fuzz_dep_.CoverTab[112350]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:182
			// _ = "end of CoverTab[112350]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:183
			_go_fuzz_dep_.CoverTab[112351]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:183
			// _ = "end of CoverTab[112351]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:183
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:183
		// _ = "end of CoverTab[112341]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:183
		_go_fuzz_dep_.CoverTab[112342]++
														t, err := timestampFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:185
			_go_fuzz_dep_.CoverTab[112352]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:186
			// _ = "end of CoverTab[112352]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:187
			_go_fuzz_dep_.CoverTab[112353]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:187
			// _ = "end of CoverTab[112353]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:187
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:187
		// _ = "end of CoverTab[112342]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:187
		_go_fuzz_dep_.CoverTab[112343]++
														s := f.asPointerTo(sub.typ).Elem()
														s.Set(reflect.ValueOf(t))
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:190
		// _ = "end of CoverTab[112343]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:191
	// _ = "end of CoverTab[112337]"
}

func makeUnmarshalTimePtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:194
	_go_fuzz_dep_.CoverTab[112354]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:195
		_go_fuzz_dep_.CoverTab[112355]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:196
			_go_fuzz_dep_.CoverTab[112361]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:197
			// _ = "end of CoverTab[112361]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:198
			_go_fuzz_dep_.CoverTab[112362]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:198
			// _ = "end of CoverTab[112362]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:198
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:198
		// _ = "end of CoverTab[112355]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:198
		_go_fuzz_dep_.CoverTab[112356]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:200
			_go_fuzz_dep_.CoverTab[112363]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:201
			// _ = "end of CoverTab[112363]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:202
			_go_fuzz_dep_.CoverTab[112364]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:202
			// _ = "end of CoverTab[112364]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:202
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:202
		// _ = "end of CoverTab[112356]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:202
		_go_fuzz_dep_.CoverTab[112357]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:204
			_go_fuzz_dep_.CoverTab[112365]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:205
			// _ = "end of CoverTab[112365]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:206
			_go_fuzz_dep_.CoverTab[112366]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:206
			// _ = "end of CoverTab[112366]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:206
		// _ = "end of CoverTab[112357]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:206
		_go_fuzz_dep_.CoverTab[112358]++
														m := &timestamp{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:208
			_go_fuzz_dep_.CoverTab[112367]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:209
			// _ = "end of CoverTab[112367]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:210
			_go_fuzz_dep_.CoverTab[112368]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:210
			// _ = "end of CoverTab[112368]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:210
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:210
		// _ = "end of CoverTab[112358]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:210
		_go_fuzz_dep_.CoverTab[112359]++
														t, err := timestampFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:212
			_go_fuzz_dep_.CoverTab[112369]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:213
			// _ = "end of CoverTab[112369]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:214
			_go_fuzz_dep_.CoverTab[112370]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:214
			// _ = "end of CoverTab[112370]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:214
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:214
		// _ = "end of CoverTab[112359]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:214
		_go_fuzz_dep_.CoverTab[112360]++
														s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
														s.Set(reflect.ValueOf(&t))
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:217
		// _ = "end of CoverTab[112360]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:218
	// _ = "end of CoverTab[112354]"
}

func makeUnmarshalTimePtrSlice(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:221
	_go_fuzz_dep_.CoverTab[112371]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:222
		_go_fuzz_dep_.CoverTab[112372]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:223
			_go_fuzz_dep_.CoverTab[112378]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:224
			// _ = "end of CoverTab[112378]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:225
			_go_fuzz_dep_.CoverTab[112379]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:225
			// _ = "end of CoverTab[112379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:225
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:225
		// _ = "end of CoverTab[112372]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:225
		_go_fuzz_dep_.CoverTab[112373]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:227
			_go_fuzz_dep_.CoverTab[112380]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:228
			// _ = "end of CoverTab[112380]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:229
			_go_fuzz_dep_.CoverTab[112381]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:229
			// _ = "end of CoverTab[112381]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:229
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:229
		// _ = "end of CoverTab[112373]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:229
		_go_fuzz_dep_.CoverTab[112374]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:231
			_go_fuzz_dep_.CoverTab[112382]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:232
			// _ = "end of CoverTab[112382]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:233
			_go_fuzz_dep_.CoverTab[112383]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:233
			// _ = "end of CoverTab[112383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:233
		// _ = "end of CoverTab[112374]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:233
		_go_fuzz_dep_.CoverTab[112375]++
														m := &timestamp{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:235
			_go_fuzz_dep_.CoverTab[112384]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:236
			// _ = "end of CoverTab[112384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:237
			_go_fuzz_dep_.CoverTab[112385]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:237
			// _ = "end of CoverTab[112385]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:237
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:237
		// _ = "end of CoverTab[112375]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:237
		_go_fuzz_dep_.CoverTab[112376]++
														t, err := timestampFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:239
			_go_fuzz_dep_.CoverTab[112386]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:240
			// _ = "end of CoverTab[112386]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:241
			_go_fuzz_dep_.CoverTab[112387]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:241
			// _ = "end of CoverTab[112387]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:241
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:241
		// _ = "end of CoverTab[112376]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:241
		_go_fuzz_dep_.CoverTab[112377]++
														slice := f.getSlice(reflect.PtrTo(sub.typ))
														newSlice := reflect.Append(slice, reflect.ValueOf(&t))
														slice.Set(newSlice)
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:245
		// _ = "end of CoverTab[112377]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:246
	// _ = "end of CoverTab[112371]"
}

func makeUnmarshalTimeSlice(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:249
	_go_fuzz_dep_.CoverTab[112388]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:250
		_go_fuzz_dep_.CoverTab[112389]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:251
			_go_fuzz_dep_.CoverTab[112395]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:252
			// _ = "end of CoverTab[112395]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:253
			_go_fuzz_dep_.CoverTab[112396]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:253
			// _ = "end of CoverTab[112396]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:253
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:253
		// _ = "end of CoverTab[112389]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:253
		_go_fuzz_dep_.CoverTab[112390]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:255
			_go_fuzz_dep_.CoverTab[112397]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:256
			// _ = "end of CoverTab[112397]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:257
			_go_fuzz_dep_.CoverTab[112398]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:257
			// _ = "end of CoverTab[112398]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:257
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:257
		// _ = "end of CoverTab[112390]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:257
		_go_fuzz_dep_.CoverTab[112391]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:259
			_go_fuzz_dep_.CoverTab[112399]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:260
			// _ = "end of CoverTab[112399]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:261
			_go_fuzz_dep_.CoverTab[112400]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:261
			// _ = "end of CoverTab[112400]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:261
		// _ = "end of CoverTab[112391]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:261
		_go_fuzz_dep_.CoverTab[112392]++
														m := &timestamp{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:263
			_go_fuzz_dep_.CoverTab[112401]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:264
			// _ = "end of CoverTab[112401]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:265
			_go_fuzz_dep_.CoverTab[112402]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:265
			// _ = "end of CoverTab[112402]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:265
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:265
		// _ = "end of CoverTab[112392]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:265
		_go_fuzz_dep_.CoverTab[112393]++
														t, err := timestampFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:267
			_go_fuzz_dep_.CoverTab[112403]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:268
			// _ = "end of CoverTab[112403]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:269
			_go_fuzz_dep_.CoverTab[112404]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:269
			// _ = "end of CoverTab[112404]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:269
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:269
		// _ = "end of CoverTab[112393]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:269
		_go_fuzz_dep_.CoverTab[112394]++
														slice := f.getSlice(sub.typ)
														newSlice := reflect.Append(slice, reflect.ValueOf(t))
														slice.Set(newSlice)
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:273
		// _ = "end of CoverTab[112394]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:274
	// _ = "end of CoverTab[112388]"
}

func makeUnmarshalDurationPtr(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:277
	_go_fuzz_dep_.CoverTab[112405]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:278
		_go_fuzz_dep_.CoverTab[112406]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:279
			_go_fuzz_dep_.CoverTab[112412]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:280
			// _ = "end of CoverTab[112412]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:281
			_go_fuzz_dep_.CoverTab[112413]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:281
			// _ = "end of CoverTab[112413]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:281
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:281
		// _ = "end of CoverTab[112406]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:281
		_go_fuzz_dep_.CoverTab[112407]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:283
			_go_fuzz_dep_.CoverTab[112414]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:284
			// _ = "end of CoverTab[112414]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:285
			_go_fuzz_dep_.CoverTab[112415]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:285
			// _ = "end of CoverTab[112415]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:285
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:285
		// _ = "end of CoverTab[112407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:285
		_go_fuzz_dep_.CoverTab[112408]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:287
			_go_fuzz_dep_.CoverTab[112416]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:288
			// _ = "end of CoverTab[112416]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:289
			_go_fuzz_dep_.CoverTab[112417]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:289
			// _ = "end of CoverTab[112417]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:289
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:289
		// _ = "end of CoverTab[112408]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:289
		_go_fuzz_dep_.CoverTab[112409]++
														m := &duration{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:291
			_go_fuzz_dep_.CoverTab[112418]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:292
			// _ = "end of CoverTab[112418]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:293
			_go_fuzz_dep_.CoverTab[112419]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:293
			// _ = "end of CoverTab[112419]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:293
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:293
		// _ = "end of CoverTab[112409]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:293
		_go_fuzz_dep_.CoverTab[112410]++
														d, err := durationFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:295
			_go_fuzz_dep_.CoverTab[112420]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:296
			// _ = "end of CoverTab[112420]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:297
			_go_fuzz_dep_.CoverTab[112421]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:297
			// _ = "end of CoverTab[112421]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:297
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:297
		// _ = "end of CoverTab[112410]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:297
		_go_fuzz_dep_.CoverTab[112411]++
														s := f.asPointerTo(reflect.PtrTo(sub.typ)).Elem()
														s.Set(reflect.ValueOf(&d))
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:300
		// _ = "end of CoverTab[112411]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:301
	// _ = "end of CoverTab[112405]"
}

func makeUnmarshalDuration(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:304
	_go_fuzz_dep_.CoverTab[112422]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:305
		_go_fuzz_dep_.CoverTab[112423]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:306
			_go_fuzz_dep_.CoverTab[112429]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:307
			// _ = "end of CoverTab[112429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:308
			_go_fuzz_dep_.CoverTab[112430]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:308
			// _ = "end of CoverTab[112430]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:308
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:308
		// _ = "end of CoverTab[112423]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:308
		_go_fuzz_dep_.CoverTab[112424]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:310
			_go_fuzz_dep_.CoverTab[112431]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:311
			// _ = "end of CoverTab[112431]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:312
			_go_fuzz_dep_.CoverTab[112432]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:312
			// _ = "end of CoverTab[112432]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:312
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:312
		// _ = "end of CoverTab[112424]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:312
		_go_fuzz_dep_.CoverTab[112425]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:314
			_go_fuzz_dep_.CoverTab[112433]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:315
			// _ = "end of CoverTab[112433]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:316
			_go_fuzz_dep_.CoverTab[112434]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:316
			// _ = "end of CoverTab[112434]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:316
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:316
		// _ = "end of CoverTab[112425]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:316
		_go_fuzz_dep_.CoverTab[112426]++
														m := &duration{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:318
			_go_fuzz_dep_.CoverTab[112435]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:319
			// _ = "end of CoverTab[112435]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:320
			_go_fuzz_dep_.CoverTab[112436]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:320
			// _ = "end of CoverTab[112436]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:320
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:320
		// _ = "end of CoverTab[112426]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:320
		_go_fuzz_dep_.CoverTab[112427]++
														d, err := durationFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:322
			_go_fuzz_dep_.CoverTab[112437]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:323
			// _ = "end of CoverTab[112437]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:324
			_go_fuzz_dep_.CoverTab[112438]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:324
			// _ = "end of CoverTab[112438]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:324
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:324
		// _ = "end of CoverTab[112427]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:324
		_go_fuzz_dep_.CoverTab[112428]++
														s := f.asPointerTo(sub.typ).Elem()
														s.Set(reflect.ValueOf(d))
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:327
		// _ = "end of CoverTab[112428]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:328
	// _ = "end of CoverTab[112422]"
}

func makeUnmarshalDurationPtrSlice(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:331
	_go_fuzz_dep_.CoverTab[112439]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:332
		_go_fuzz_dep_.CoverTab[112440]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:333
			_go_fuzz_dep_.CoverTab[112446]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:334
			// _ = "end of CoverTab[112446]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:335
			_go_fuzz_dep_.CoverTab[112447]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:335
			// _ = "end of CoverTab[112447]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:335
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:335
		// _ = "end of CoverTab[112440]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:335
		_go_fuzz_dep_.CoverTab[112441]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:337
			_go_fuzz_dep_.CoverTab[112448]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:338
			// _ = "end of CoverTab[112448]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:339
			_go_fuzz_dep_.CoverTab[112449]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:339
			// _ = "end of CoverTab[112449]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:339
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:339
		// _ = "end of CoverTab[112441]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:339
		_go_fuzz_dep_.CoverTab[112442]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:341
			_go_fuzz_dep_.CoverTab[112450]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:342
			// _ = "end of CoverTab[112450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:343
			_go_fuzz_dep_.CoverTab[112451]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:343
			// _ = "end of CoverTab[112451]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:343
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:343
		// _ = "end of CoverTab[112442]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:343
		_go_fuzz_dep_.CoverTab[112443]++
														m := &duration{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:345
			_go_fuzz_dep_.CoverTab[112452]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:346
			// _ = "end of CoverTab[112452]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:347
			_go_fuzz_dep_.CoverTab[112453]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:347
			// _ = "end of CoverTab[112453]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:347
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:347
		// _ = "end of CoverTab[112443]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:347
		_go_fuzz_dep_.CoverTab[112444]++
														d, err := durationFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:349
			_go_fuzz_dep_.CoverTab[112454]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:350
			// _ = "end of CoverTab[112454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:351
			_go_fuzz_dep_.CoverTab[112455]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:351
			// _ = "end of CoverTab[112455]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:351
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:351
		// _ = "end of CoverTab[112444]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:351
		_go_fuzz_dep_.CoverTab[112445]++
														slice := f.getSlice(reflect.PtrTo(sub.typ))
														newSlice := reflect.Append(slice, reflect.ValueOf(&d))
														slice.Set(newSlice)
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:355
		// _ = "end of CoverTab[112445]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:356
	// _ = "end of CoverTab[112439]"
}

func makeUnmarshalDurationSlice(sub *unmarshalInfo, name string) unmarshaler {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:359
	_go_fuzz_dep_.CoverTab[112456]++
													return func(b []byte, f pointer, w int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:360
		_go_fuzz_dep_.CoverTab[112457]++
														if w != WireBytes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:361
			_go_fuzz_dep_.CoverTab[112463]++
															return nil, errInternalBadWireType
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:362
			// _ = "end of CoverTab[112463]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:363
			_go_fuzz_dep_.CoverTab[112464]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:363
			// _ = "end of CoverTab[112464]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:363
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:363
		// _ = "end of CoverTab[112457]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:363
		_go_fuzz_dep_.CoverTab[112458]++
														x, n := decodeVarint(b)
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:365
			_go_fuzz_dep_.CoverTab[112465]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:366
			// _ = "end of CoverTab[112465]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:367
			_go_fuzz_dep_.CoverTab[112466]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:367
			// _ = "end of CoverTab[112466]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:367
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:367
		// _ = "end of CoverTab[112458]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:367
		_go_fuzz_dep_.CoverTab[112459]++
														b = b[n:]
														if x > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:369
			_go_fuzz_dep_.CoverTab[112467]++
															return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:370
			// _ = "end of CoverTab[112467]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:371
			_go_fuzz_dep_.CoverTab[112468]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:371
			// _ = "end of CoverTab[112468]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:371
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:371
		// _ = "end of CoverTab[112459]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:371
		_go_fuzz_dep_.CoverTab[112460]++
														m := &duration{}
														if err := Unmarshal(b[:x], m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:373
			_go_fuzz_dep_.CoverTab[112469]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:374
			// _ = "end of CoverTab[112469]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:375
			_go_fuzz_dep_.CoverTab[112470]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:375
			// _ = "end of CoverTab[112470]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:375
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:375
		// _ = "end of CoverTab[112460]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:375
		_go_fuzz_dep_.CoverTab[112461]++
														d, err := durationFromProto(m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:377
			_go_fuzz_dep_.CoverTab[112471]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:378
			// _ = "end of CoverTab[112471]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:379
			_go_fuzz_dep_.CoverTab[112472]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:379
			// _ = "end of CoverTab[112472]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:379
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:379
		// _ = "end of CoverTab[112461]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:379
		_go_fuzz_dep_.CoverTab[112462]++
														slice := f.getSlice(sub.typ)
														newSlice := reflect.Append(slice, reflect.ValueOf(d))
														slice.Set(newSlice)
														return b[x:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:383
		// _ = "end of CoverTab[112462]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:384
	// _ = "end of CoverTab[112456]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:385
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/table_unmarshal_gogo.go:385
var _ = _go_fuzz_dep_.CoverTab
