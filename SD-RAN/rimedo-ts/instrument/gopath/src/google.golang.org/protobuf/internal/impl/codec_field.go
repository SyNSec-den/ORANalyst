// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:5
)

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)

type errInvalidUTF8 struct{}

func (errInvalidUTF8) Error() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:22
	_go_fuzz_dep_.CoverTab[53808]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:22
	return "string field contains invalid UTF-8"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:22
	// _ = "end of CoverTab[53808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:22
}
func (errInvalidUTF8) InvalidUTF8() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:23
	_go_fuzz_dep_.CoverTab[53809]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:23
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:23
	// _ = "end of CoverTab[53809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:23
}
func (errInvalidUTF8) Unwrap() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:24
	_go_fuzz_dep_.CoverTab[53810]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:24
	return errors.Error
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:24
	// _ = "end of CoverTab[53810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:24
}

// initOneofFieldCoders initializes the fast-path functions for the fields in a oneof.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:26
// For size, marshal, and isInit operations, functions are set only on the first field
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:26
// in the oneof. The functions are called when the oneof is non-nil, and will dispatch
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:26
// to the appropriate field-specific function as necessary.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:26
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:26
// The unmarshal function is set on each field individually as usual.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:33
func (mi *MessageInfo) initOneofFieldCoders(od protoreflect.OneofDescriptor, si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:33
	_go_fuzz_dep_.CoverTab[53811]++
													fs := si.oneofsByName[od.Name()]
													ft := fs.Type
													oneofFields := make(map[reflect.Type]*coderFieldInfo)
													needIsInit := false
													fields := od.Fields()
													for i, lim := 0, fields.Len(); i < lim; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:39
		_go_fuzz_dep_.CoverTab[53817]++
														fd := od.Fields().Get(i)
														num := fd.Number()

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:48
		cf := *mi.coderFields[num]
		ot := si.oneofWrappersByNumber[num]
		cf.ft = ot.Field(0).Type
		cf.mi, cf.funcs = fieldCoder(fd, cf.ft)
		oneofFields[ot] = &cf
		if cf.funcs.isInit != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:53
			_go_fuzz_dep_.CoverTab[53819]++
															needIsInit = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:54
			// _ = "end of CoverTab[53819]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:55
			_go_fuzz_dep_.CoverTab[53820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:55
			// _ = "end of CoverTab[53820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:55
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:55
		// _ = "end of CoverTab[53817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:55
		_go_fuzz_dep_.CoverTab[53818]++
														mi.coderFields[num].funcs.unmarshal = func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:56
			_go_fuzz_dep_.CoverTab[53821]++
															var vw reflect.Value	// pointer to wrapper type
															vi := p.AsValueOf(ft).Elem()
															if !vi.IsNil() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				_go_fuzz_dep_.CoverTab[53824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				return !vi.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				// _ = "end of CoverTab[53824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				_go_fuzz_dep_.CoverTab[53825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				return vi.Elem().Elem().Type() == ot
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				// _ = "end of CoverTab[53825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:59
				_go_fuzz_dep_.CoverTab[53826]++
																vw = vi.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:60
				// _ = "end of CoverTab[53826]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:61
				_go_fuzz_dep_.CoverTab[53827]++
																vw = reflect.New(ot)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:62
				// _ = "end of CoverTab[53827]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:63
			// _ = "end of CoverTab[53821]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:63
			_go_fuzz_dep_.CoverTab[53822]++
															out, err := cf.funcs.unmarshal(b, pointerOfValue(vw).Apply(zeroOffset), wtyp, &cf, opts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:65
				_go_fuzz_dep_.CoverTab[53828]++
																return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:66
				// _ = "end of CoverTab[53828]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:67
				_go_fuzz_dep_.CoverTab[53829]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:67
				// _ = "end of CoverTab[53829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:67
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:67
			// _ = "end of CoverTab[53822]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:67
			_go_fuzz_dep_.CoverTab[53823]++
															vi.Set(vw)
															return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:69
			// _ = "end of CoverTab[53823]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:70
		// _ = "end of CoverTab[53818]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:71
	// _ = "end of CoverTab[53811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:71
	_go_fuzz_dep_.CoverTab[53812]++
													getInfo := func(p pointer) (pointer, *coderFieldInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:72
		_go_fuzz_dep_.CoverTab[53830]++
														v := p.AsValueOf(ft).Elem()
														if v.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:74
			_go_fuzz_dep_.CoverTab[53833]++
															return pointer{}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:75
			// _ = "end of CoverTab[53833]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:76
			_go_fuzz_dep_.CoverTab[53834]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:76
			// _ = "end of CoverTab[53834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:76
		// _ = "end of CoverTab[53830]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:76
		_go_fuzz_dep_.CoverTab[53831]++
														v = v.Elem()
														if v.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:78
			_go_fuzz_dep_.CoverTab[53835]++
															return pointer{}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:79
			// _ = "end of CoverTab[53835]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:80
			_go_fuzz_dep_.CoverTab[53836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:80
			// _ = "end of CoverTab[53836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:80
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:80
		// _ = "end of CoverTab[53831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:80
		_go_fuzz_dep_.CoverTab[53832]++
														return pointerOfValue(v).Apply(zeroOffset), oneofFields[v.Elem().Type()]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:81
		// _ = "end of CoverTab[53832]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:82
	// _ = "end of CoverTab[53812]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:82
	_go_fuzz_dep_.CoverTab[53813]++
													first := mi.coderFields[od.Fields().Get(0).Number()]
													first.funcs.size = func(p pointer, _ *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:84
		_go_fuzz_dep_.CoverTab[53837]++
														p, info := getInfo(p)
														if info == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:86
			_go_fuzz_dep_.CoverTab[53839]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:86
			return info.funcs.size == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:86
			// _ = "end of CoverTab[53839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:86
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:86
			_go_fuzz_dep_.CoverTab[53840]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:87
			// _ = "end of CoverTab[53840]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:88
			_go_fuzz_dep_.CoverTab[53841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:88
			// _ = "end of CoverTab[53841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:88
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:88
		// _ = "end of CoverTab[53837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:88
		_go_fuzz_dep_.CoverTab[53838]++
														return info.funcs.size(p, info, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:89
		// _ = "end of CoverTab[53838]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:90
	// _ = "end of CoverTab[53813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:90
	_go_fuzz_dep_.CoverTab[53814]++
													first.funcs.marshal = func(b []byte, p pointer, _ *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:91
		_go_fuzz_dep_.CoverTab[53842]++
														p, info := getInfo(p)
														if info == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:93
			_go_fuzz_dep_.CoverTab[53844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:93
			return info.funcs.marshal == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:93
			// _ = "end of CoverTab[53844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:93
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:93
			_go_fuzz_dep_.CoverTab[53845]++
															return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:94
			// _ = "end of CoverTab[53845]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:95
			_go_fuzz_dep_.CoverTab[53846]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:95
			// _ = "end of CoverTab[53846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:95
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:95
		// _ = "end of CoverTab[53842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:95
		_go_fuzz_dep_.CoverTab[53843]++
														return info.funcs.marshal(b, p, info, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:96
		// _ = "end of CoverTab[53843]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:97
	// _ = "end of CoverTab[53814]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:97
	_go_fuzz_dep_.CoverTab[53815]++
													first.funcs.merge = func(dst, src pointer, _ *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:98
		_go_fuzz_dep_.CoverTab[53847]++
														srcp, srcinfo := getInfo(src)
														if srcinfo == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:100
			_go_fuzz_dep_.CoverTab[53850]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:100
			return srcinfo.funcs.merge == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:100
			// _ = "end of CoverTab[53850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:100
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:100
			_go_fuzz_dep_.CoverTab[53851]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:101
			// _ = "end of CoverTab[53851]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:102
			_go_fuzz_dep_.CoverTab[53852]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:102
			// _ = "end of CoverTab[53852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:102
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:102
		// _ = "end of CoverTab[53847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:102
		_go_fuzz_dep_.CoverTab[53848]++
														dstp, dstinfo := getInfo(dst)
														if dstinfo != srcinfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:104
			_go_fuzz_dep_.CoverTab[53853]++
															dst.AsValueOf(ft).Elem().Set(reflect.New(src.AsValueOf(ft).Elem().Elem().Elem().Type()))
															dstp = pointerOfValue(dst.AsValueOf(ft).Elem().Elem()).Apply(zeroOffset)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:106
			// _ = "end of CoverTab[53853]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:107
			_go_fuzz_dep_.CoverTab[53854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:107
			// _ = "end of CoverTab[53854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:107
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:107
		// _ = "end of CoverTab[53848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:107
		_go_fuzz_dep_.CoverTab[53849]++
														srcinfo.funcs.merge(dstp, srcp, srcinfo, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:108
		// _ = "end of CoverTab[53849]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:109
	// _ = "end of CoverTab[53815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:109
	_go_fuzz_dep_.CoverTab[53816]++
													if needIsInit {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:110
		_go_fuzz_dep_.CoverTab[53855]++
														first.funcs.isInit = func(p pointer, _ *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:111
			_go_fuzz_dep_.CoverTab[53856]++
															p, info := getInfo(p)
															if info == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:113
				_go_fuzz_dep_.CoverTab[53858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:113
				return info.funcs.isInit == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:113
				// _ = "end of CoverTab[53858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:113
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:113
				_go_fuzz_dep_.CoverTab[53859]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:114
				// _ = "end of CoverTab[53859]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:115
				_go_fuzz_dep_.CoverTab[53860]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:115
				// _ = "end of CoverTab[53860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:115
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:115
			// _ = "end of CoverTab[53856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:115
			_go_fuzz_dep_.CoverTab[53857]++
															return info.funcs.isInit(p, info)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:116
			// _ = "end of CoverTab[53857]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:117
		// _ = "end of CoverTab[53855]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:118
		_go_fuzz_dep_.CoverTab[53861]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:118
		// _ = "end of CoverTab[53861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:118
	// _ = "end of CoverTab[53816]"
}

func makeWeakMessageFieldCoder(fd protoreflect.FieldDescriptor) pointerCoderFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:121
	_go_fuzz_dep_.CoverTab[53862]++
													var once sync.Once
													var messageType protoreflect.MessageType
													lazyInit := func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:124
		_go_fuzz_dep_.CoverTab[53864]++
														once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:125
			_go_fuzz_dep_.CoverTab[53865]++
															messageName := fd.Message().FullName()
															messageType, _ = protoregistry.GlobalTypes.FindMessageByName(messageName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:127
			// _ = "end of CoverTab[53865]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:128
		// _ = "end of CoverTab[53864]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:129
	// _ = "end of CoverTab[53862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:129
	_go_fuzz_dep_.CoverTab[53863]++

													return pointerCoderFuncs{
		size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:132
			_go_fuzz_dep_.CoverTab[53866]++
															m, ok := p.WeakFields().get(f.num)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:134
				_go_fuzz_dep_.CoverTab[53869]++
																return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:135
				// _ = "end of CoverTab[53869]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:136
				_go_fuzz_dep_.CoverTab[53870]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:136
				// _ = "end of CoverTab[53870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:136
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:136
			// _ = "end of CoverTab[53866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:136
			_go_fuzz_dep_.CoverTab[53867]++
															lazyInit()
															if messageType == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:138
				_go_fuzz_dep_.CoverTab[53871]++
																panic(fmt.Sprintf("weak message %v is not linked in", fd.Message().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:139
				// _ = "end of CoverTab[53871]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:140
				_go_fuzz_dep_.CoverTab[53872]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:140
				// _ = "end of CoverTab[53872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:140
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:140
			// _ = "end of CoverTab[53867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:140
			_go_fuzz_dep_.CoverTab[53868]++
															return sizeMessage(m, f.tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:141
			// _ = "end of CoverTab[53868]"
		},
		marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:143
			_go_fuzz_dep_.CoverTab[53873]++
															m, ok := p.WeakFields().get(f.num)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:145
				_go_fuzz_dep_.CoverTab[53876]++
																return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:146
				// _ = "end of CoverTab[53876]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:147
				_go_fuzz_dep_.CoverTab[53877]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:147
				// _ = "end of CoverTab[53877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:147
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:147
			// _ = "end of CoverTab[53873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:147
			_go_fuzz_dep_.CoverTab[53874]++
															lazyInit()
															if messageType == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:149
				_go_fuzz_dep_.CoverTab[53878]++
																panic(fmt.Sprintf("weak message %v is not linked in", fd.Message().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:150
				// _ = "end of CoverTab[53878]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:151
				_go_fuzz_dep_.CoverTab[53879]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:151
				// _ = "end of CoverTab[53879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:151
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:151
			// _ = "end of CoverTab[53874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:151
			_go_fuzz_dep_.CoverTab[53875]++
															return appendMessage(b, m, f.wiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:152
			// _ = "end of CoverTab[53875]"
		},
		unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:154
			_go_fuzz_dep_.CoverTab[53880]++
															fs := p.WeakFields()
															m, ok := fs.get(f.num)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:157
				_go_fuzz_dep_.CoverTab[53882]++
																lazyInit()
																if messageType == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:159
					_go_fuzz_dep_.CoverTab[53884]++
																	return unmarshalOutput{}, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:160
					// _ = "end of CoverTab[53884]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:161
					_go_fuzz_dep_.CoverTab[53885]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:161
					// _ = "end of CoverTab[53885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:161
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:161
				// _ = "end of CoverTab[53882]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:161
				_go_fuzz_dep_.CoverTab[53883]++
																m = messageType.New().Interface()
																fs.set(f.num, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:163
				// _ = "end of CoverTab[53883]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:164
				_go_fuzz_dep_.CoverTab[53886]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:164
				// _ = "end of CoverTab[53886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:164
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:164
			// _ = "end of CoverTab[53880]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:164
			_go_fuzz_dep_.CoverTab[53881]++
															return consumeMessage(b, m, wtyp, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:165
			// _ = "end of CoverTab[53881]"
		},
		isInit: func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:167
			_go_fuzz_dep_.CoverTab[53887]++
															m, ok := p.WeakFields().get(f.num)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:169
				_go_fuzz_dep_.CoverTab[53889]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:170
				// _ = "end of CoverTab[53889]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:171
				_go_fuzz_dep_.CoverTab[53890]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:171
				// _ = "end of CoverTab[53890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:171
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:171
			// _ = "end of CoverTab[53887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:171
			_go_fuzz_dep_.CoverTab[53888]++
															return proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:172
			// _ = "end of CoverTab[53888]"
		},
		merge: func(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:174
			_go_fuzz_dep_.CoverTab[53891]++
															sm, ok := src.WeakFields().get(f.num)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:176
				_go_fuzz_dep_.CoverTab[53894]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:177
				// _ = "end of CoverTab[53894]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:178
				_go_fuzz_dep_.CoverTab[53895]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:178
				// _ = "end of CoverTab[53895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:178
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:178
			// _ = "end of CoverTab[53891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:178
			_go_fuzz_dep_.CoverTab[53892]++
															dm, ok := dst.WeakFields().get(f.num)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:180
				_go_fuzz_dep_.CoverTab[53896]++
																lazyInit()
																if messageType == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:182
					_go_fuzz_dep_.CoverTab[53898]++
																	panic(fmt.Sprintf("weak message %v is not linked in", fd.Message().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:183
					// _ = "end of CoverTab[53898]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:184
					_go_fuzz_dep_.CoverTab[53899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:184
					// _ = "end of CoverTab[53899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:184
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:184
				// _ = "end of CoverTab[53896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:184
				_go_fuzz_dep_.CoverTab[53897]++
																dm = messageType.New().Interface()
																dst.WeakFields().set(f.num, dm)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:186
				// _ = "end of CoverTab[53897]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:187
				_go_fuzz_dep_.CoverTab[53900]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:187
				// _ = "end of CoverTab[53900]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:187
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:187
			// _ = "end of CoverTab[53892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:187
			_go_fuzz_dep_.CoverTab[53893]++
															opts.Merge(dm, sm)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:188
			// _ = "end of CoverTab[53893]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:190
	// _ = "end of CoverTab[53863]"
}

func makeMessageFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoderFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:193
	_go_fuzz_dep_.CoverTab[53901]++
													if mi := getMessageInfo(ft); mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:194
		_go_fuzz_dep_.CoverTab[53902]++
														funcs := pointerCoderFuncs{
			size:		sizeMessageInfo,
			marshal:	appendMessageInfo,
			unmarshal:	consumeMessageInfo,
			merge:		mergeMessage,
		}
		if needsInitCheck(mi.Desc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:201
			_go_fuzz_dep_.CoverTab[53904]++
															funcs.isInit = isInitMessageInfo
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:202
			// _ = "end of CoverTab[53904]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:203
			_go_fuzz_dep_.CoverTab[53905]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:203
			// _ = "end of CoverTab[53905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:203
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:203
		// _ = "end of CoverTab[53902]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:203
		_go_fuzz_dep_.CoverTab[53903]++
														return funcs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:204
		// _ = "end of CoverTab[53903]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:205
		_go_fuzz_dep_.CoverTab[53906]++
														return pointerCoderFuncs{
			size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:207
				_go_fuzz_dep_.CoverTab[53907]++
																m := asMessage(p.AsValueOf(ft).Elem())
																return sizeMessage(m, f.tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:209
				// _ = "end of CoverTab[53907]"
			},
			marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:211
				_go_fuzz_dep_.CoverTab[53908]++
																m := asMessage(p.AsValueOf(ft).Elem())
																return appendMessage(b, m, f.wiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:213
				// _ = "end of CoverTab[53908]"
			},
			unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:215
				_go_fuzz_dep_.CoverTab[53909]++
																mp := p.AsValueOf(ft).Elem()
																if mp.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:217
					_go_fuzz_dep_.CoverTab[53911]++
																	mp.Set(reflect.New(ft.Elem()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:218
					// _ = "end of CoverTab[53911]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:219
					_go_fuzz_dep_.CoverTab[53912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:219
					// _ = "end of CoverTab[53912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:219
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:219
				// _ = "end of CoverTab[53909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:219
				_go_fuzz_dep_.CoverTab[53910]++
																return consumeMessage(b, asMessage(mp), wtyp, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:220
				// _ = "end of CoverTab[53910]"
			},
			isInit: func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:222
				_go_fuzz_dep_.CoverTab[53913]++
																m := asMessage(p.AsValueOf(ft).Elem())
																return proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:224
				// _ = "end of CoverTab[53913]"
			},
			merge:	mergeMessage,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:227
		// _ = "end of CoverTab[53906]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:228
	// _ = "end of CoverTab[53901]"
}

func sizeMessageInfo(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:231
	_go_fuzz_dep_.CoverTab[53914]++
													return protowire.SizeBytes(f.mi.sizePointer(p.Elem(), opts)) + f.tagsize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:232
	// _ = "end of CoverTab[53914]"
}

func appendMessageInfo(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:235
	_go_fuzz_dep_.CoverTab[53915]++
													b = protowire.AppendVarint(b, f.wiretag)
													b = protowire.AppendVarint(b, uint64(f.mi.sizePointer(p.Elem(), opts)))
													return f.mi.marshalAppendPointer(b, p.Elem(), opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:238
	// _ = "end of CoverTab[53915]"
}

func consumeMessageInfo(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:241
	_go_fuzz_dep_.CoverTab[53916]++
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:242
		_go_fuzz_dep_.CoverTab[53921]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:243
		// _ = "end of CoverTab[53921]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:244
		_go_fuzz_dep_.CoverTab[53922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:244
		// _ = "end of CoverTab[53922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:244
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:244
	// _ = "end of CoverTab[53916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:244
	_go_fuzz_dep_.CoverTab[53917]++
													v, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:246
		_go_fuzz_dep_.CoverTab[53923]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:247
		// _ = "end of CoverTab[53923]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:248
		_go_fuzz_dep_.CoverTab[53924]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:248
		// _ = "end of CoverTab[53924]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:248
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:248
	// _ = "end of CoverTab[53917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:248
	_go_fuzz_dep_.CoverTab[53918]++
													if p.Elem().IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:249
		_go_fuzz_dep_.CoverTab[53925]++
														p.SetPointer(pointerOfValue(reflect.New(f.mi.GoReflectType.Elem())))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:250
		// _ = "end of CoverTab[53925]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:251
		_go_fuzz_dep_.CoverTab[53926]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:251
		// _ = "end of CoverTab[53926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:251
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:251
	// _ = "end of CoverTab[53918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:251
	_go_fuzz_dep_.CoverTab[53919]++
													o, err := f.mi.unmarshalPointer(v, p.Elem(), 0, opts)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:253
		_go_fuzz_dep_.CoverTab[53927]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:254
		// _ = "end of CoverTab[53927]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:255
		_go_fuzz_dep_.CoverTab[53928]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:255
		// _ = "end of CoverTab[53928]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:255
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:255
	// _ = "end of CoverTab[53919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:255
	_go_fuzz_dep_.CoverTab[53920]++
													out.n = n
													out.initialized = o.initialized
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:258
	// _ = "end of CoverTab[53920]"
}

func isInitMessageInfo(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:261
	_go_fuzz_dep_.CoverTab[53929]++
													return f.mi.checkInitializedPointer(p.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:262
	// _ = "end of CoverTab[53929]"
}

func sizeMessage(m proto.Message, tagsize int, _ marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:265
	_go_fuzz_dep_.CoverTab[53930]++
													return protowire.SizeBytes(proto.Size(m)) + tagsize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:266
	// _ = "end of CoverTab[53930]"
}

func appendMessage(b []byte, m proto.Message, wiretag uint64, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:269
	_go_fuzz_dep_.CoverTab[53931]++
													b = protowire.AppendVarint(b, wiretag)
													b = protowire.AppendVarint(b, uint64(proto.Size(m)))
													return opts.Options().MarshalAppend(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:272
	// _ = "end of CoverTab[53931]"
}

func consumeMessage(b []byte, m proto.Message, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:275
	_go_fuzz_dep_.CoverTab[53932]++
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:276
		_go_fuzz_dep_.CoverTab[53936]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:277
		// _ = "end of CoverTab[53936]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:278
		_go_fuzz_dep_.CoverTab[53937]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:278
		// _ = "end of CoverTab[53937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:278
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:278
	// _ = "end of CoverTab[53932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:278
	_go_fuzz_dep_.CoverTab[53933]++
													v, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:280
		_go_fuzz_dep_.CoverTab[53938]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:281
		// _ = "end of CoverTab[53938]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:282
		_go_fuzz_dep_.CoverTab[53939]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:282
		// _ = "end of CoverTab[53939]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:282
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:282
	// _ = "end of CoverTab[53933]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:282
	_go_fuzz_dep_.CoverTab[53934]++
													o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:		v,
		Message:	m.ProtoReflect(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:287
		_go_fuzz_dep_.CoverTab[53940]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:288
		// _ = "end of CoverTab[53940]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:289
		_go_fuzz_dep_.CoverTab[53941]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:289
		// _ = "end of CoverTab[53941]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:289
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:289
	// _ = "end of CoverTab[53934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:289
	_go_fuzz_dep_.CoverTab[53935]++
													out.n = n
													out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:292
	// _ = "end of CoverTab[53935]"
}

func sizeMessageValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:295
	_go_fuzz_dep_.CoverTab[53942]++
													m := v.Message().Interface()
													return sizeMessage(m, tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:297
	// _ = "end of CoverTab[53942]"
}

func appendMessageValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:300
	_go_fuzz_dep_.CoverTab[53943]++
													m := v.Message().Interface()
													return appendMessage(b, m, wiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:302
	// _ = "end of CoverTab[53943]"
}

func consumeMessageValue(b []byte, v protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (protoreflect.Value, unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:305
	_go_fuzz_dep_.CoverTab[53944]++
													m := v.Message().Interface()
													out, err := consumeMessage(b, m, wtyp, opts)
													return v, out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:308
	// _ = "end of CoverTab[53944]"
}

func isInitMessageValue(v protoreflect.Value) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:311
	_go_fuzz_dep_.CoverTab[53945]++
													m := v.Message().Interface()
													return proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:313
	// _ = "end of CoverTab[53945]"
}

var coderMessageValue = valueCoderFuncs{
	size:		sizeMessageValue,
	marshal:	appendMessageValue,
	unmarshal:	consumeMessageValue,
	isInit:		isInitMessageValue,
	merge:		mergeMessageValue,
}

func sizeGroupValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:324
	_go_fuzz_dep_.CoverTab[53946]++
													m := v.Message().Interface()
													return sizeGroup(m, tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:326
	// _ = "end of CoverTab[53946]"
}

func appendGroupValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:329
	_go_fuzz_dep_.CoverTab[53947]++
													m := v.Message().Interface()
													return appendGroup(b, m, wiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:331
	// _ = "end of CoverTab[53947]"
}

func consumeGroupValue(b []byte, v protoreflect.Value, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (protoreflect.Value, unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:334
	_go_fuzz_dep_.CoverTab[53948]++
													m := v.Message().Interface()
													out, err := consumeGroup(b, m, num, wtyp, opts)
													return v, out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:337
	// _ = "end of CoverTab[53948]"
}

var coderGroupValue = valueCoderFuncs{
	size:		sizeGroupValue,
	marshal:	appendGroupValue,
	unmarshal:	consumeGroupValue,
	isInit:		isInitMessageValue,
	merge:		mergeMessageValue,
}

func makeGroupFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoderFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:348
	_go_fuzz_dep_.CoverTab[53949]++
													num := fd.Number()
													if mi := getMessageInfo(ft); mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:350
		_go_fuzz_dep_.CoverTab[53950]++
														funcs := pointerCoderFuncs{
			size:		sizeGroupType,
			marshal:	appendGroupType,
			unmarshal:	consumeGroupType,
			merge:		mergeMessage,
		}
		if needsInitCheck(mi.Desc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:357
			_go_fuzz_dep_.CoverTab[53952]++
															funcs.isInit = isInitMessageInfo
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:358
			// _ = "end of CoverTab[53952]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:359
			_go_fuzz_dep_.CoverTab[53953]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:359
			// _ = "end of CoverTab[53953]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:359
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:359
		// _ = "end of CoverTab[53950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:359
		_go_fuzz_dep_.CoverTab[53951]++
														return funcs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:360
		// _ = "end of CoverTab[53951]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:361
		_go_fuzz_dep_.CoverTab[53954]++
														return pointerCoderFuncs{
			size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:363
				_go_fuzz_dep_.CoverTab[53955]++
																m := asMessage(p.AsValueOf(ft).Elem())
																return sizeGroup(m, f.tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:365
				// _ = "end of CoverTab[53955]"
			},
			marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:367
				_go_fuzz_dep_.CoverTab[53956]++
																m := asMessage(p.AsValueOf(ft).Elem())
																return appendGroup(b, m, f.wiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:369
				// _ = "end of CoverTab[53956]"
			},
			unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:371
				_go_fuzz_dep_.CoverTab[53957]++
																mp := p.AsValueOf(ft).Elem()
																if mp.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:373
					_go_fuzz_dep_.CoverTab[53959]++
																	mp.Set(reflect.New(ft.Elem()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:374
					// _ = "end of CoverTab[53959]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:375
					_go_fuzz_dep_.CoverTab[53960]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:375
					// _ = "end of CoverTab[53960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:375
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:375
				// _ = "end of CoverTab[53957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:375
				_go_fuzz_dep_.CoverTab[53958]++
																return consumeGroup(b, asMessage(mp), num, wtyp, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:376
				// _ = "end of CoverTab[53958]"
			},
			isInit: func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:378
				_go_fuzz_dep_.CoverTab[53961]++
																m := asMessage(p.AsValueOf(ft).Elem())
																return proto.CheckInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:380
				// _ = "end of CoverTab[53961]"
			},
			merge:	mergeMessage,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:383
		// _ = "end of CoverTab[53954]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:384
	// _ = "end of CoverTab[53949]"
}

func sizeGroupType(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:387
	_go_fuzz_dep_.CoverTab[53962]++
													return 2*f.tagsize + f.mi.sizePointer(p.Elem(), opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:388
	// _ = "end of CoverTab[53962]"
}

func appendGroupType(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:391
	_go_fuzz_dep_.CoverTab[53963]++
													b = protowire.AppendVarint(b, f.wiretag)
													b, err := f.mi.marshalAppendPointer(b, p.Elem(), opts)
													b = protowire.AppendVarint(b, f.wiretag+1)
													return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:395
	// _ = "end of CoverTab[53963]"
}

func consumeGroupType(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:398
	_go_fuzz_dep_.CoverTab[53964]++
													if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:399
		_go_fuzz_dep_.CoverTab[53967]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:400
		// _ = "end of CoverTab[53967]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:401
		_go_fuzz_dep_.CoverTab[53968]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:401
		// _ = "end of CoverTab[53968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:401
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:401
	// _ = "end of CoverTab[53964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:401
	_go_fuzz_dep_.CoverTab[53965]++
													if p.Elem().IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:402
		_go_fuzz_dep_.CoverTab[53969]++
														p.SetPointer(pointerOfValue(reflect.New(f.mi.GoReflectType.Elem())))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:403
		// _ = "end of CoverTab[53969]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:404
		_go_fuzz_dep_.CoverTab[53970]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:404
		// _ = "end of CoverTab[53970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:404
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:404
	// _ = "end of CoverTab[53965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:404
	_go_fuzz_dep_.CoverTab[53966]++
													return f.mi.unmarshalPointer(b, p.Elem(), f.num, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:405
	// _ = "end of CoverTab[53966]"
}

func sizeGroup(m proto.Message, tagsize int, _ marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:408
	_go_fuzz_dep_.CoverTab[53971]++
													return 2*tagsize + proto.Size(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:409
	// _ = "end of CoverTab[53971]"
}

func appendGroup(b []byte, m proto.Message, wiretag uint64, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:412
	_go_fuzz_dep_.CoverTab[53972]++
													b = protowire.AppendVarint(b, wiretag)
													b, err := opts.Options().MarshalAppend(b, m)
													b = protowire.AppendVarint(b, wiretag+1)
													return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:416
	// _ = "end of CoverTab[53972]"
}

func consumeGroup(b []byte, m proto.Message, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:419
	_go_fuzz_dep_.CoverTab[53973]++
													if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:420
		_go_fuzz_dep_.CoverTab[53977]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:421
		// _ = "end of CoverTab[53977]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:422
		_go_fuzz_dep_.CoverTab[53978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:422
		// _ = "end of CoverTab[53978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:422
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:422
	// _ = "end of CoverTab[53973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:422
	_go_fuzz_dep_.CoverTab[53974]++
													b, n := protowire.ConsumeGroup(num, b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:424
		_go_fuzz_dep_.CoverTab[53979]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:425
		// _ = "end of CoverTab[53979]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:426
		_go_fuzz_dep_.CoverTab[53980]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:426
		// _ = "end of CoverTab[53980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:426
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:426
	// _ = "end of CoverTab[53974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:426
	_go_fuzz_dep_.CoverTab[53975]++
													o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:		b,
		Message:	m.ProtoReflect(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:431
		_go_fuzz_dep_.CoverTab[53981]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:432
		// _ = "end of CoverTab[53981]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:433
		_go_fuzz_dep_.CoverTab[53982]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:433
		// _ = "end of CoverTab[53982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:433
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:433
	// _ = "end of CoverTab[53975]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:433
	_go_fuzz_dep_.CoverTab[53976]++
													out.n = n
													out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:436
	// _ = "end of CoverTab[53976]"
}

func makeMessageSliceFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoderFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:439
	_go_fuzz_dep_.CoverTab[53983]++
													if mi := getMessageInfo(ft); mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:440
		_go_fuzz_dep_.CoverTab[53985]++
														funcs := pointerCoderFuncs{
			size:		sizeMessageSliceInfo,
			marshal:	appendMessageSliceInfo,
			unmarshal:	consumeMessageSliceInfo,
			merge:		mergeMessageSlice,
		}
		if needsInitCheck(mi.Desc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:447
			_go_fuzz_dep_.CoverTab[53987]++
															funcs.isInit = isInitMessageSliceInfo
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:448
			// _ = "end of CoverTab[53987]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:449
			_go_fuzz_dep_.CoverTab[53988]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:449
			// _ = "end of CoverTab[53988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:449
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:449
		// _ = "end of CoverTab[53985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:449
		_go_fuzz_dep_.CoverTab[53986]++
														return funcs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:450
		// _ = "end of CoverTab[53986]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:451
		_go_fuzz_dep_.CoverTab[53989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:451
		// _ = "end of CoverTab[53989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:451
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:451
	// _ = "end of CoverTab[53983]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:451
	_go_fuzz_dep_.CoverTab[53984]++
													return pointerCoderFuncs{
		size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:453
			_go_fuzz_dep_.CoverTab[53990]++
															return sizeMessageSlice(p, ft, f.tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:454
			// _ = "end of CoverTab[53990]"
		},
		marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:456
			_go_fuzz_dep_.CoverTab[53991]++
															return appendMessageSlice(b, p, f.wiretag, ft, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:457
			// _ = "end of CoverTab[53991]"
		},
		unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:459
			_go_fuzz_dep_.CoverTab[53992]++
															return consumeMessageSlice(b, p, ft, wtyp, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:460
			// _ = "end of CoverTab[53992]"
		},
		isInit: func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:462
			_go_fuzz_dep_.CoverTab[53993]++
															return isInitMessageSlice(p, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:463
			// _ = "end of CoverTab[53993]"
		},
		merge:	mergeMessageSlice,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:466
	// _ = "end of CoverTab[53984]"
}

func sizeMessageSliceInfo(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:469
	_go_fuzz_dep_.CoverTab[53994]++
													s := p.PointerSlice()
													n := 0
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:472
		_go_fuzz_dep_.CoverTab[53996]++
														n += protowire.SizeBytes(f.mi.sizePointer(v, opts)) + f.tagsize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:473
		// _ = "end of CoverTab[53996]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:474
	// _ = "end of CoverTab[53994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:474
	_go_fuzz_dep_.CoverTab[53995]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:475
	// _ = "end of CoverTab[53995]"
}

func appendMessageSliceInfo(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:478
	_go_fuzz_dep_.CoverTab[53997]++
													s := p.PointerSlice()
													var err error
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:481
		_go_fuzz_dep_.CoverTab[53999]++
														b = protowire.AppendVarint(b, f.wiretag)
														siz := f.mi.sizePointer(v, opts)
														b = protowire.AppendVarint(b, uint64(siz))
														b, err = f.mi.marshalAppendPointer(b, v, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:486
			_go_fuzz_dep_.CoverTab[54000]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:487
			// _ = "end of CoverTab[54000]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:488
			_go_fuzz_dep_.CoverTab[54001]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:488
			// _ = "end of CoverTab[54001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:488
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:488
		// _ = "end of CoverTab[53999]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:489
	// _ = "end of CoverTab[53997]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:489
	_go_fuzz_dep_.CoverTab[53998]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:490
	// _ = "end of CoverTab[53998]"
}

func consumeMessageSliceInfo(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:493
	_go_fuzz_dep_.CoverTab[54002]++
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:494
		_go_fuzz_dep_.CoverTab[54006]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:495
		// _ = "end of CoverTab[54006]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:496
		_go_fuzz_dep_.CoverTab[54007]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:496
		// _ = "end of CoverTab[54007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:496
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:496
	// _ = "end of CoverTab[54002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:496
	_go_fuzz_dep_.CoverTab[54003]++
													v, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:498
		_go_fuzz_dep_.CoverTab[54008]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:499
		// _ = "end of CoverTab[54008]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:500
		_go_fuzz_dep_.CoverTab[54009]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:500
		// _ = "end of CoverTab[54009]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:500
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:500
	// _ = "end of CoverTab[54003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:500
	_go_fuzz_dep_.CoverTab[54004]++
													m := reflect.New(f.mi.GoReflectType.Elem()).Interface()
													mp := pointerOfIface(m)
													o, err := f.mi.unmarshalPointer(v, mp, 0, opts)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:504
		_go_fuzz_dep_.CoverTab[54010]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:505
		// _ = "end of CoverTab[54010]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:506
		_go_fuzz_dep_.CoverTab[54011]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:506
		// _ = "end of CoverTab[54011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:506
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:506
	// _ = "end of CoverTab[54004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:506
	_go_fuzz_dep_.CoverTab[54005]++
													p.AppendPointerSlice(mp)
													out.n = n
													out.initialized = o.initialized
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:510
	// _ = "end of CoverTab[54005]"
}

func isInitMessageSliceInfo(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:513
	_go_fuzz_dep_.CoverTab[54012]++
													s := p.PointerSlice()
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:515
		_go_fuzz_dep_.CoverTab[54014]++
														if err := f.mi.checkInitializedPointer(v); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:516
			_go_fuzz_dep_.CoverTab[54015]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:517
			// _ = "end of CoverTab[54015]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:518
			_go_fuzz_dep_.CoverTab[54016]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:518
			// _ = "end of CoverTab[54016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:518
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:518
		// _ = "end of CoverTab[54014]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:519
	// _ = "end of CoverTab[54012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:519
	_go_fuzz_dep_.CoverTab[54013]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:520
	// _ = "end of CoverTab[54013]"
}

func sizeMessageSlice(p pointer, goType reflect.Type, tagsize int, _ marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:523
	_go_fuzz_dep_.CoverTab[54017]++
													s := p.PointerSlice()
													n := 0
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:526
		_go_fuzz_dep_.CoverTab[54019]++
														m := asMessage(v.AsValueOf(goType.Elem()))
														n += protowire.SizeBytes(proto.Size(m)) + tagsize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:528
		// _ = "end of CoverTab[54019]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:529
	// _ = "end of CoverTab[54017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:529
	_go_fuzz_dep_.CoverTab[54018]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:530
	// _ = "end of CoverTab[54018]"
}

func appendMessageSlice(b []byte, p pointer, wiretag uint64, goType reflect.Type, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:533
	_go_fuzz_dep_.CoverTab[54020]++
													s := p.PointerSlice()
													var err error
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:536
		_go_fuzz_dep_.CoverTab[54022]++
														m := asMessage(v.AsValueOf(goType.Elem()))
														b = protowire.AppendVarint(b, wiretag)
														siz := proto.Size(m)
														b = protowire.AppendVarint(b, uint64(siz))
														b, err = opts.Options().MarshalAppend(b, m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:542
			_go_fuzz_dep_.CoverTab[54023]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:543
			// _ = "end of CoverTab[54023]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:544
			_go_fuzz_dep_.CoverTab[54024]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:544
			// _ = "end of CoverTab[54024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:544
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:544
		// _ = "end of CoverTab[54022]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:545
	// _ = "end of CoverTab[54020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:545
	_go_fuzz_dep_.CoverTab[54021]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:546
	// _ = "end of CoverTab[54021]"
}

func consumeMessageSlice(b []byte, p pointer, goType reflect.Type, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:549
	_go_fuzz_dep_.CoverTab[54025]++
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:550
		_go_fuzz_dep_.CoverTab[54029]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:551
		// _ = "end of CoverTab[54029]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:552
		_go_fuzz_dep_.CoverTab[54030]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:552
		// _ = "end of CoverTab[54030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:552
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:552
	// _ = "end of CoverTab[54025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:552
	_go_fuzz_dep_.CoverTab[54026]++
													v, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:554
		_go_fuzz_dep_.CoverTab[54031]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:555
		// _ = "end of CoverTab[54031]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:556
		_go_fuzz_dep_.CoverTab[54032]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:556
		// _ = "end of CoverTab[54032]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:556
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:556
	// _ = "end of CoverTab[54026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:556
	_go_fuzz_dep_.CoverTab[54027]++
													mp := reflect.New(goType.Elem())
													o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:		v,
		Message:	asMessage(mp).ProtoReflect(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:562
		_go_fuzz_dep_.CoverTab[54033]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:563
		// _ = "end of CoverTab[54033]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:564
		_go_fuzz_dep_.CoverTab[54034]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:564
		// _ = "end of CoverTab[54034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:564
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:564
	// _ = "end of CoverTab[54027]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:564
	_go_fuzz_dep_.CoverTab[54028]++
													p.AppendPointerSlice(pointerOfValue(mp))
													out.n = n
													out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:568
	// _ = "end of CoverTab[54028]"
}

func isInitMessageSlice(p pointer, goType reflect.Type) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:571
	_go_fuzz_dep_.CoverTab[54035]++
													s := p.PointerSlice()
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:573
		_go_fuzz_dep_.CoverTab[54037]++
														m := asMessage(v.AsValueOf(goType.Elem()))
														if err := proto.CheckInitialized(m); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:575
			_go_fuzz_dep_.CoverTab[54038]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:576
			// _ = "end of CoverTab[54038]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:577
			_go_fuzz_dep_.CoverTab[54039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:577
			// _ = "end of CoverTab[54039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:577
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:577
		// _ = "end of CoverTab[54037]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:578
	// _ = "end of CoverTab[54035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:578
	_go_fuzz_dep_.CoverTab[54036]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:579
	// _ = "end of CoverTab[54036]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:584
func sizeMessageSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:584
	_go_fuzz_dep_.CoverTab[54040]++
													list := listv.List()
													n := 0
													for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:587
		_go_fuzz_dep_.CoverTab[54042]++
														m := list.Get(i).Message().Interface()
														n += protowire.SizeBytes(proto.Size(m)) + tagsize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:589
		// _ = "end of CoverTab[54042]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:590
	// _ = "end of CoverTab[54040]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:590
	_go_fuzz_dep_.CoverTab[54041]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:591
	// _ = "end of CoverTab[54041]"
}

func appendMessageSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:594
	_go_fuzz_dep_.CoverTab[54043]++
													list := listv.List()
													mopts := opts.Options()
													for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:597
		_go_fuzz_dep_.CoverTab[54045]++
														m := list.Get(i).Message().Interface()
														b = protowire.AppendVarint(b, wiretag)
														siz := proto.Size(m)
														b = protowire.AppendVarint(b, uint64(siz))
														var err error
														b, err = mopts.MarshalAppend(b, m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:604
			_go_fuzz_dep_.CoverTab[54046]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:605
			// _ = "end of CoverTab[54046]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:606
			_go_fuzz_dep_.CoverTab[54047]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:606
			// _ = "end of CoverTab[54047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:606
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:606
		// _ = "end of CoverTab[54045]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:607
	// _ = "end of CoverTab[54043]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:607
	_go_fuzz_dep_.CoverTab[54044]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:608
	// _ = "end of CoverTab[54044]"
}

func consumeMessageSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:611
	_go_fuzz_dep_.CoverTab[54048]++
													list := listv.List()
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:613
		_go_fuzz_dep_.CoverTab[54052]++
														return protoreflect.Value{}, out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:614
		// _ = "end of CoverTab[54052]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:615
		_go_fuzz_dep_.CoverTab[54053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:615
		// _ = "end of CoverTab[54053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:615
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:615
	// _ = "end of CoverTab[54048]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:615
	_go_fuzz_dep_.CoverTab[54049]++
													v, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:617
		_go_fuzz_dep_.CoverTab[54054]++
														return protoreflect.Value{}, out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:618
		// _ = "end of CoverTab[54054]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:619
		_go_fuzz_dep_.CoverTab[54055]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:619
		// _ = "end of CoverTab[54055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:619
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:619
	// _ = "end of CoverTab[54049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:619
	_go_fuzz_dep_.CoverTab[54050]++
													m := list.NewElement()
													o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:		v,
		Message:	m.Message(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:625
		_go_fuzz_dep_.CoverTab[54056]++
														return protoreflect.Value{}, out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:626
		// _ = "end of CoverTab[54056]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:627
		_go_fuzz_dep_.CoverTab[54057]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:627
		// _ = "end of CoverTab[54057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:627
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:627
	// _ = "end of CoverTab[54050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:627
	_go_fuzz_dep_.CoverTab[54051]++
													list.Append(m)
													out.n = n
													out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
													return listv, out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:631
	// _ = "end of CoverTab[54051]"
}

func isInitMessageSliceValue(listv protoreflect.Value) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:634
	_go_fuzz_dep_.CoverTab[54058]++
													list := listv.List()
													for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:636
		_go_fuzz_dep_.CoverTab[54060]++
														m := list.Get(i).Message().Interface()
														if err := proto.CheckInitialized(m); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:638
			_go_fuzz_dep_.CoverTab[54061]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:639
			// _ = "end of CoverTab[54061]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:640
			_go_fuzz_dep_.CoverTab[54062]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:640
			// _ = "end of CoverTab[54062]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:640
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:640
		// _ = "end of CoverTab[54060]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:641
	// _ = "end of CoverTab[54058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:641
	_go_fuzz_dep_.CoverTab[54059]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:642
	// _ = "end of CoverTab[54059]"
}

var coderMessageSliceValue = valueCoderFuncs{
	size:		sizeMessageSliceValue,
	marshal:	appendMessageSliceValue,
	unmarshal:	consumeMessageSliceValue,
	isInit:		isInitMessageSliceValue,
	merge:		mergeMessageListValue,
}

func sizeGroupSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:653
	_go_fuzz_dep_.CoverTab[54063]++
													list := listv.List()
													n := 0
													for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:656
		_go_fuzz_dep_.CoverTab[54065]++
														m := list.Get(i).Message().Interface()
														n += 2*tagsize + proto.Size(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:658
		// _ = "end of CoverTab[54065]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:659
	// _ = "end of CoverTab[54063]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:659
	_go_fuzz_dep_.CoverTab[54064]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:660
	// _ = "end of CoverTab[54064]"
}

func appendGroupSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:663
	_go_fuzz_dep_.CoverTab[54066]++
													list := listv.List()
													mopts := opts.Options()
													for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:666
		_go_fuzz_dep_.CoverTab[54068]++
														m := list.Get(i).Message().Interface()
														b = protowire.AppendVarint(b, wiretag)
														var err error
														b, err = mopts.MarshalAppend(b, m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:671
			_go_fuzz_dep_.CoverTab[54070]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:672
			// _ = "end of CoverTab[54070]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:673
			_go_fuzz_dep_.CoverTab[54071]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:673
			// _ = "end of CoverTab[54071]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:673
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:673
		// _ = "end of CoverTab[54068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:673
		_go_fuzz_dep_.CoverTab[54069]++
														b = protowire.AppendVarint(b, wiretag+1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:674
		// _ = "end of CoverTab[54069]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:675
	// _ = "end of CoverTab[54066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:675
	_go_fuzz_dep_.CoverTab[54067]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:676
	// _ = "end of CoverTab[54067]"
}

func consumeGroupSliceValue(b []byte, listv protoreflect.Value, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:679
	_go_fuzz_dep_.CoverTab[54072]++
													list := listv.List()
													if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:681
		_go_fuzz_dep_.CoverTab[54076]++
														return protoreflect.Value{}, out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:682
		// _ = "end of CoverTab[54076]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:683
		_go_fuzz_dep_.CoverTab[54077]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:683
		// _ = "end of CoverTab[54077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:683
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:683
	// _ = "end of CoverTab[54072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:683
	_go_fuzz_dep_.CoverTab[54073]++
													b, n := protowire.ConsumeGroup(num, b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:685
		_go_fuzz_dep_.CoverTab[54078]++
														return protoreflect.Value{}, out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:686
		// _ = "end of CoverTab[54078]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:687
		_go_fuzz_dep_.CoverTab[54079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:687
		// _ = "end of CoverTab[54079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:687
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:687
	// _ = "end of CoverTab[54073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:687
	_go_fuzz_dep_.CoverTab[54074]++
													m := list.NewElement()
													o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:		b,
		Message:	m.Message(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:693
		_go_fuzz_dep_.CoverTab[54080]++
														return protoreflect.Value{}, out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:694
		// _ = "end of CoverTab[54080]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:695
		_go_fuzz_dep_.CoverTab[54081]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:695
		// _ = "end of CoverTab[54081]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:695
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:695
	// _ = "end of CoverTab[54074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:695
	_go_fuzz_dep_.CoverTab[54075]++
													list.Append(m)
													out.n = n
													out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
													return listv, out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:699
	// _ = "end of CoverTab[54075]"
}

var coderGroupSliceValue = valueCoderFuncs{
	size:		sizeGroupSliceValue,
	marshal:	appendGroupSliceValue,
	unmarshal:	consumeGroupSliceValue,
	isInit:		isInitMessageSliceValue,
	merge:		mergeMessageListValue,
}

func makeGroupSliceFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoderFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:710
	_go_fuzz_dep_.CoverTab[54082]++
													num := fd.Number()
													if mi := getMessageInfo(ft); mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:712
		_go_fuzz_dep_.CoverTab[54084]++
														funcs := pointerCoderFuncs{
			size:		sizeGroupSliceInfo,
			marshal:	appendGroupSliceInfo,
			unmarshal:	consumeGroupSliceInfo,
			merge:		mergeMessageSlice,
		}
		if needsInitCheck(mi.Desc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:719
			_go_fuzz_dep_.CoverTab[54086]++
															funcs.isInit = isInitMessageSliceInfo
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:720
			// _ = "end of CoverTab[54086]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:721
			_go_fuzz_dep_.CoverTab[54087]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:721
			// _ = "end of CoverTab[54087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:721
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:721
		// _ = "end of CoverTab[54084]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:721
		_go_fuzz_dep_.CoverTab[54085]++
														return funcs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:722
		// _ = "end of CoverTab[54085]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:723
		_go_fuzz_dep_.CoverTab[54088]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:723
		// _ = "end of CoverTab[54088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:723
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:723
	// _ = "end of CoverTab[54082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:723
	_go_fuzz_dep_.CoverTab[54083]++
													return pointerCoderFuncs{
		size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:725
			_go_fuzz_dep_.CoverTab[54089]++
															return sizeGroupSlice(p, ft, f.tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:726
			// _ = "end of CoverTab[54089]"
		},
		marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:728
			_go_fuzz_dep_.CoverTab[54090]++
															return appendGroupSlice(b, p, f.wiretag, ft, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:729
			// _ = "end of CoverTab[54090]"
		},
		unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:731
			_go_fuzz_dep_.CoverTab[54091]++
															return consumeGroupSlice(b, p, num, wtyp, ft, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:732
			// _ = "end of CoverTab[54091]"
		},
		isInit: func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:734
			_go_fuzz_dep_.CoverTab[54092]++
															return isInitMessageSlice(p, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:735
			// _ = "end of CoverTab[54092]"
		},
		merge:	mergeMessageSlice,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:738
	// _ = "end of CoverTab[54083]"
}

func sizeGroupSlice(p pointer, messageType reflect.Type, tagsize int, _ marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:741
	_go_fuzz_dep_.CoverTab[54093]++
													s := p.PointerSlice()
													n := 0
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:744
		_go_fuzz_dep_.CoverTab[54095]++
														m := asMessage(v.AsValueOf(messageType.Elem()))
														n += 2*tagsize + proto.Size(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:746
		// _ = "end of CoverTab[54095]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:747
	// _ = "end of CoverTab[54093]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:747
	_go_fuzz_dep_.CoverTab[54094]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:748
	// _ = "end of CoverTab[54094]"
}

func appendGroupSlice(b []byte, p pointer, wiretag uint64, messageType reflect.Type, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:751
	_go_fuzz_dep_.CoverTab[54096]++
													s := p.PointerSlice()
													var err error
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:754
		_go_fuzz_dep_.CoverTab[54098]++
														m := asMessage(v.AsValueOf(messageType.Elem()))
														b = protowire.AppendVarint(b, wiretag)
														b, err = opts.Options().MarshalAppend(b, m)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:758
			_go_fuzz_dep_.CoverTab[54100]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:759
			// _ = "end of CoverTab[54100]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:760
			_go_fuzz_dep_.CoverTab[54101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:760
			// _ = "end of CoverTab[54101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:760
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:760
		// _ = "end of CoverTab[54098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:760
		_go_fuzz_dep_.CoverTab[54099]++
														b = protowire.AppendVarint(b, wiretag+1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:761
		// _ = "end of CoverTab[54099]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:762
	// _ = "end of CoverTab[54096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:762
	_go_fuzz_dep_.CoverTab[54097]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:763
	// _ = "end of CoverTab[54097]"
}

func consumeGroupSlice(b []byte, p pointer, num protowire.Number, wtyp protowire.Type, goType reflect.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:766
	_go_fuzz_dep_.CoverTab[54102]++
													if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:767
		_go_fuzz_dep_.CoverTab[54106]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:768
		// _ = "end of CoverTab[54106]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:769
		_go_fuzz_dep_.CoverTab[54107]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:769
		// _ = "end of CoverTab[54107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:769
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:769
	// _ = "end of CoverTab[54102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:769
	_go_fuzz_dep_.CoverTab[54103]++
													b, n := protowire.ConsumeGroup(num, b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:771
		_go_fuzz_dep_.CoverTab[54108]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:772
		// _ = "end of CoverTab[54108]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:773
		_go_fuzz_dep_.CoverTab[54109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:773
		// _ = "end of CoverTab[54109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:773
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:773
	// _ = "end of CoverTab[54103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:773
	_go_fuzz_dep_.CoverTab[54104]++
													mp := reflect.New(goType.Elem())
													o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:		b,
		Message:	asMessage(mp).ProtoReflect(),
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:779
		_go_fuzz_dep_.CoverTab[54110]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:780
		// _ = "end of CoverTab[54110]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:781
		_go_fuzz_dep_.CoverTab[54111]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:781
		// _ = "end of CoverTab[54111]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:781
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:781
	// _ = "end of CoverTab[54104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:781
	_go_fuzz_dep_.CoverTab[54105]++
													p.AppendPointerSlice(pointerOfValue(mp))
													out.n = n
													out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:785
	// _ = "end of CoverTab[54105]"
}

func sizeGroupSliceInfo(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:788
	_go_fuzz_dep_.CoverTab[54112]++
													s := p.PointerSlice()
													n := 0
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:791
		_go_fuzz_dep_.CoverTab[54114]++
														n += 2*f.tagsize + f.mi.sizePointer(v, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:792
		// _ = "end of CoverTab[54114]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:793
	// _ = "end of CoverTab[54112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:793
	_go_fuzz_dep_.CoverTab[54113]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:794
	// _ = "end of CoverTab[54113]"
}

func appendGroupSliceInfo(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:797
	_go_fuzz_dep_.CoverTab[54115]++
													s := p.PointerSlice()
													var err error
													for _, v := range s {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:800
		_go_fuzz_dep_.CoverTab[54117]++
														b = protowire.AppendVarint(b, f.wiretag)
														b, err = f.mi.marshalAppendPointer(b, v, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:803
			_go_fuzz_dep_.CoverTab[54119]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:804
			// _ = "end of CoverTab[54119]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:805
			_go_fuzz_dep_.CoverTab[54120]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:805
			// _ = "end of CoverTab[54120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:805
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:805
		// _ = "end of CoverTab[54117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:805
		_go_fuzz_dep_.CoverTab[54118]++
														b = protowire.AppendVarint(b, f.wiretag+1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:806
		// _ = "end of CoverTab[54118]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:807
	// _ = "end of CoverTab[54115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:807
	_go_fuzz_dep_.CoverTab[54116]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:808
	// _ = "end of CoverTab[54116]"
}

func consumeGroupSliceInfo(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:811
	_go_fuzz_dep_.CoverTab[54121]++
													if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:812
		_go_fuzz_dep_.CoverTab[54124]++
														return unmarshalOutput{}, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:813
		// _ = "end of CoverTab[54124]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:814
		_go_fuzz_dep_.CoverTab[54125]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:814
		// _ = "end of CoverTab[54125]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:814
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:814
	// _ = "end of CoverTab[54121]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:814
	_go_fuzz_dep_.CoverTab[54122]++
													m := reflect.New(f.mi.GoReflectType.Elem()).Interface()
													mp := pointerOfIface(m)
													out, err := f.mi.unmarshalPointer(b, mp, f.num, opts)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:818
		_go_fuzz_dep_.CoverTab[54126]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:819
		// _ = "end of CoverTab[54126]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:820
		_go_fuzz_dep_.CoverTab[54127]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:820
		// _ = "end of CoverTab[54127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:820
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:820
	// _ = "end of CoverTab[54122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:820
	_go_fuzz_dep_.CoverTab[54123]++
													p.AppendPointerSlice(mp)
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:822
	// _ = "end of CoverTab[54123]"
}

func asMessage(v reflect.Value) protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:825
	_go_fuzz_dep_.CoverTab[54128]++
													if m, ok := v.Interface().(protoreflect.ProtoMessage); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:826
		_go_fuzz_dep_.CoverTab[54130]++
														return m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:827
		// _ = "end of CoverTab[54130]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:828
		_go_fuzz_dep_.CoverTab[54131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:828
		// _ = "end of CoverTab[54131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:828
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:828
	// _ = "end of CoverTab[54128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:828
	_go_fuzz_dep_.CoverTab[54129]++
													return legacyWrapMessage(v).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:829
	// _ = "end of CoverTab[54129]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:830
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_field.go:830
var _ = _go_fuzz_dep_.CoverTab
