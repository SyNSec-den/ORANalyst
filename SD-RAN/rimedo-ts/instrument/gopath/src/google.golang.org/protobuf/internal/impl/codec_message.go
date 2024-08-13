// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:5
)

import (
	"fmt"
	"reflect"
	"sort"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/order"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// coderMessageInfo contains per-message information used by the fast-path functions.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:19
// This is a different type from MessageInfo to keep MessageInfo as general-purpose as
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:19
// possible.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:22
type coderMessageInfo struct {
	methods	protoiface.Methods

	orderedCoderFields	[]*coderFieldInfo
	denseCoderFields	[]*coderFieldInfo
	coderFields		map[protowire.Number]*coderFieldInfo
	sizecacheOffset		offset
	unknownOffset		offset
	unknownPtrKind		bool
	extensionOffset		offset
	needsInitCheck		bool
	isMessageSet		bool
	numRequiredFields	uint8
}

type coderFieldInfo struct {
	funcs		pointerCoderFuncs	// fast-path per-field functions
	mi		*MessageInfo		// field's message
	ft		reflect.Type
	validation	validationInfo			// information used by message validation
	num		protoreflect.FieldNumber	// field number
	offset		offset				// struct field offset
	wiretag		uint64				// field tag (number + wire type)
	tagsize		int				// size of the varint-encoded tag
	isPointer	bool				// true if IsNil may be called on the struct field
	isRequired	bool				// true if field is required
}

func (mi *MessageInfo) makeCoderMethods(t reflect.Type, si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:50
	_go_fuzz_dep_.CoverTab[56367]++
													mi.sizecacheOffset = invalidOffset
													mi.unknownOffset = invalidOffset
													mi.extensionOffset = invalidOffset

													if si.sizecacheOffset.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:55
		_go_fuzz_dep_.CoverTab[56381]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:55
		return si.sizecacheType == sizecacheType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:55
		// _ = "end of CoverTab[56381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:55
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:55
		_go_fuzz_dep_.CoverTab[56382]++
														mi.sizecacheOffset = si.sizecacheOffset
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:56
		// _ = "end of CoverTab[56382]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:57
		_go_fuzz_dep_.CoverTab[56383]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:57
		// _ = "end of CoverTab[56383]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:57
	// _ = "end of CoverTab[56367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:57
	_go_fuzz_dep_.CoverTab[56368]++
													if si.unknownOffset.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
		_go_fuzz_dep_.CoverTab[56384]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
		return (si.unknownType == unknownFieldsAType || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
			_go_fuzz_dep_.CoverTab[56385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
			return si.unknownType == unknownFieldsBType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
			// _ = "end of CoverTab[56385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
		// _ = "end of CoverTab[56384]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:58
		_go_fuzz_dep_.CoverTab[56386]++
														mi.unknownOffset = si.unknownOffset
														mi.unknownPtrKind = si.unknownType.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:60
		// _ = "end of CoverTab[56386]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:61
		_go_fuzz_dep_.CoverTab[56387]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:61
		// _ = "end of CoverTab[56387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:61
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:61
	// _ = "end of CoverTab[56368]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:61
	_go_fuzz_dep_.CoverTab[56369]++
													if si.extensionOffset.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:62
		_go_fuzz_dep_.CoverTab[56388]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:62
		return si.extensionType == extensionFieldsType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:62
		// _ = "end of CoverTab[56388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:62
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:62
		_go_fuzz_dep_.CoverTab[56389]++
														mi.extensionOffset = si.extensionOffset
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:63
		// _ = "end of CoverTab[56389]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:64
		_go_fuzz_dep_.CoverTab[56390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:64
		// _ = "end of CoverTab[56390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:64
	// _ = "end of CoverTab[56369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:64
	_go_fuzz_dep_.CoverTab[56370]++

													mi.coderFields = make(map[protowire.Number]*coderFieldInfo)
													fields := mi.Desc.Fields()
													preallocFields := make([]coderFieldInfo, fields.Len())
													for i := 0; i < fields.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:69
		_go_fuzz_dep_.CoverTab[56391]++
														fd := fields.Get(i)

														fs := si.fieldsByNumber[fd.Number()]
														isOneof := fd.ContainingOneof() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:73
			_go_fuzz_dep_.CoverTab[56395]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:73
			return !fd.ContainingOneof().IsSynthetic()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:73
			// _ = "end of CoverTab[56395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:73
		}()
														if isOneof {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:74
			_go_fuzz_dep_.CoverTab[56396]++
															fs = si.oneofsByName[fd.ContainingOneof().Name()]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:75
			// _ = "end of CoverTab[56396]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:76
			_go_fuzz_dep_.CoverTab[56397]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:76
			// _ = "end of CoverTab[56397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:76
		// _ = "end of CoverTab[56391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:76
		_go_fuzz_dep_.CoverTab[56392]++
														ft := fs.Type
														var wiretag uint64
														if !fd.IsPacked() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:79
			_go_fuzz_dep_.CoverTab[56398]++
															wiretag = protowire.EncodeTag(fd.Number(), wireTypes[fd.Kind()])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:80
			// _ = "end of CoverTab[56398]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:81
			_go_fuzz_dep_.CoverTab[56399]++
															wiretag = protowire.EncodeTag(fd.Number(), protowire.BytesType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:82
			// _ = "end of CoverTab[56399]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:83
		// _ = "end of CoverTab[56392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:83
		_go_fuzz_dep_.CoverTab[56393]++
														var fieldOffset offset
														var funcs pointerCoderFuncs
														var childMessage *MessageInfo
														switch {
		case ft == nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:88
			_go_fuzz_dep_.CoverTab[56400]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:92
			funcs = pointerCoderFuncs{
				size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:93
					_go_fuzz_dep_.CoverTab[56404]++
																	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:94
					// _ = "end of CoverTab[56404]"
				},
				marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:96
					_go_fuzz_dep_.CoverTab[56405]++
																	return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:97
					// _ = "end of CoverTab[56405]"
				},
				unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:99
						_go_fuzz_dep_.CoverTab[56406]++
																		panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:100
					// _ = "end of CoverTab[56406]"
				},
				isInit: func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:102
					_go_fuzz_dep_.CoverTab[56407]++
																		panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:103
					// _ = "end of CoverTab[56407]"
				},
				merge: func(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:105
					_go_fuzz_dep_.CoverTab[56408]++
																		panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:106
					// _ = "end of CoverTab[56408]"
				},
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:108
			// _ = "end of CoverTab[56400]"
		case isOneof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:109
			_go_fuzz_dep_.CoverTab[56401]++
																fieldOffset = offsetOf(fs, mi.Exporter)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:110
			// _ = "end of CoverTab[56401]"
		case fd.IsWeak():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:111
			_go_fuzz_dep_.CoverTab[56402]++
																fieldOffset = si.weakOffset
																funcs = makeWeakMessageFieldCoder(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:113
			// _ = "end of CoverTab[56402]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:114
			_go_fuzz_dep_.CoverTab[56403]++
																fieldOffset = offsetOf(fs, mi.Exporter)
																childMessage, funcs = fieldCoder(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:116
			// _ = "end of CoverTab[56403]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:117
		// _ = "end of CoverTab[56393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:117
		_go_fuzz_dep_.CoverTab[56394]++
															cf := &preallocFields[i]
															*cf = coderFieldInfo{
			num:		fd.Number(),
			offset:		fieldOffset,
			wiretag:	wiretag,
			ft:		ft,
			tagsize:	protowire.SizeVarint(wiretag),
			funcs:		funcs,
			mi:		childMessage,
			validation:	newFieldValidationInfo(mi, si, fd, ft),
			isPointer: fd.Cardinality() == protoreflect.Repeated || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:128
				_go_fuzz_dep_.CoverTab[56409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:128
				return fd.HasPresence()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:128
				// _ = "end of CoverTab[56409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:128
			}(),
			isRequired:	fd.Cardinality() == protoreflect.Required,
		}
															mi.orderedCoderFields = append(mi.orderedCoderFields, cf)
															mi.coderFields[cf.num] = cf
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:132
		// _ = "end of CoverTab[56394]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:133
	// _ = "end of CoverTab[56370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:133
	_go_fuzz_dep_.CoverTab[56371]++
														for i, oneofs := 0, mi.Desc.Oneofs(); i < oneofs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:134
		_go_fuzz_dep_.CoverTab[56410]++
															if od := oneofs.Get(i); !od.IsSynthetic() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:135
			_go_fuzz_dep_.CoverTab[56411]++
																mi.initOneofFieldCoders(od, si)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:136
			// _ = "end of CoverTab[56411]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:137
			_go_fuzz_dep_.CoverTab[56412]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:137
			// _ = "end of CoverTab[56412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:137
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:137
		// _ = "end of CoverTab[56410]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:138
	// _ = "end of CoverTab[56371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:138
	_go_fuzz_dep_.CoverTab[56372]++
														if messageset.IsMessageSet(mi.Desc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:139
		_go_fuzz_dep_.CoverTab[56413]++
															if !mi.extensionOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:140
			_go_fuzz_dep_.CoverTab[56416]++
																panic(fmt.Sprintf("%v: MessageSet with no extensions field", mi.Desc.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:141
			// _ = "end of CoverTab[56416]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:142
			_go_fuzz_dep_.CoverTab[56417]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:142
			// _ = "end of CoverTab[56417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:142
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:142
		// _ = "end of CoverTab[56413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:142
		_go_fuzz_dep_.CoverTab[56414]++
															if !mi.unknownOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:143
			_go_fuzz_dep_.CoverTab[56418]++
																panic(fmt.Sprintf("%v: MessageSet with no unknown field", mi.Desc.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:144
			// _ = "end of CoverTab[56418]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:145
			_go_fuzz_dep_.CoverTab[56419]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:145
			// _ = "end of CoverTab[56419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:145
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:145
		// _ = "end of CoverTab[56414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:145
		_go_fuzz_dep_.CoverTab[56415]++
															mi.isMessageSet = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:146
		// _ = "end of CoverTab[56415]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:147
		_go_fuzz_dep_.CoverTab[56420]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:147
		// _ = "end of CoverTab[56420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:147
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:147
	// _ = "end of CoverTab[56372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:147
	_go_fuzz_dep_.CoverTab[56373]++
														sort.Slice(mi.orderedCoderFields, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:148
		_go_fuzz_dep_.CoverTab[56421]++
															return mi.orderedCoderFields[i].num < mi.orderedCoderFields[j].num
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:149
		// _ = "end of CoverTab[56421]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:150
	// _ = "end of CoverTab[56373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:150
	_go_fuzz_dep_.CoverTab[56374]++

														var maxDense protoreflect.FieldNumber
														for _, cf := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:153
		_go_fuzz_dep_.CoverTab[56422]++
															if cf.num >= 16 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:154
			_go_fuzz_dep_.CoverTab[56424]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:154
			return cf.num >= 2*maxDense
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:154
			// _ = "end of CoverTab[56424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:154
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:154
			_go_fuzz_dep_.CoverTab[56425]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:155
			// _ = "end of CoverTab[56425]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:156
			_go_fuzz_dep_.CoverTab[56426]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:156
			// _ = "end of CoverTab[56426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:156
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:156
		// _ = "end of CoverTab[56422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:156
		_go_fuzz_dep_.CoverTab[56423]++
															maxDense = cf.num
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:157
		// _ = "end of CoverTab[56423]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:158
	// _ = "end of CoverTab[56374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:158
	_go_fuzz_dep_.CoverTab[56375]++
														mi.denseCoderFields = make([]*coderFieldInfo, maxDense+1)
														for _, cf := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:160
		_go_fuzz_dep_.CoverTab[56427]++
															if int(cf.num) >= len(mi.denseCoderFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:161
			_go_fuzz_dep_.CoverTab[56429]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:162
			// _ = "end of CoverTab[56429]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:163
			_go_fuzz_dep_.CoverTab[56430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:163
			// _ = "end of CoverTab[56430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:163
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:163
		// _ = "end of CoverTab[56427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:163
		_go_fuzz_dep_.CoverTab[56428]++
															mi.denseCoderFields[cf.num] = cf
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:164
		// _ = "end of CoverTab[56428]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:165
	// _ = "end of CoverTab[56375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:165
	_go_fuzz_dep_.CoverTab[56376]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:168
	if mi.Desc.Oneofs().Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:168
		_go_fuzz_dep_.CoverTab[56431]++
															sort.Slice(mi.orderedCoderFields, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:169
			_go_fuzz_dep_.CoverTab[56432]++
																fi := fields.ByNumber(mi.orderedCoderFields[i].num)
																fj := fields.ByNumber(mi.orderedCoderFields[j].num)
																return order.LegacyFieldOrder(fi, fj)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:172
			// _ = "end of CoverTab[56432]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:173
		// _ = "end of CoverTab[56431]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:174
		_go_fuzz_dep_.CoverTab[56433]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:174
		// _ = "end of CoverTab[56433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:174
	// _ = "end of CoverTab[56376]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:174
	_go_fuzz_dep_.CoverTab[56377]++

														mi.needsInitCheck = needsInitCheck(mi.Desc)
														if mi.methods.Marshal == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:177
		_go_fuzz_dep_.CoverTab[56434]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:177
		return mi.methods.Size == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:177
		// _ = "end of CoverTab[56434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:177
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:177
		_go_fuzz_dep_.CoverTab[56435]++
															mi.methods.Flags |= protoiface.SupportMarshalDeterministic
															mi.methods.Marshal = mi.marshal
															mi.methods.Size = mi.size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:180
		// _ = "end of CoverTab[56435]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:181
		_go_fuzz_dep_.CoverTab[56436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:181
		// _ = "end of CoverTab[56436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:181
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:181
	// _ = "end of CoverTab[56377]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:181
	_go_fuzz_dep_.CoverTab[56378]++
														if mi.methods.Unmarshal == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:182
		_go_fuzz_dep_.CoverTab[56437]++
															mi.methods.Flags |= protoiface.SupportUnmarshalDiscardUnknown
															mi.methods.Unmarshal = mi.unmarshal
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:184
		// _ = "end of CoverTab[56437]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:185
		_go_fuzz_dep_.CoverTab[56438]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:185
		// _ = "end of CoverTab[56438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:185
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:185
	// _ = "end of CoverTab[56378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:185
	_go_fuzz_dep_.CoverTab[56379]++
														if mi.methods.CheckInitialized == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:186
		_go_fuzz_dep_.CoverTab[56439]++
															mi.methods.CheckInitialized = mi.checkInitialized
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:187
		// _ = "end of CoverTab[56439]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:188
		_go_fuzz_dep_.CoverTab[56440]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:188
		// _ = "end of CoverTab[56440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:188
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:188
	// _ = "end of CoverTab[56379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:188
	_go_fuzz_dep_.CoverTab[56380]++
														if mi.methods.Merge == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:189
		_go_fuzz_dep_.CoverTab[56441]++
															mi.methods.Merge = mi.merge
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:190
		// _ = "end of CoverTab[56441]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:191
		_go_fuzz_dep_.CoverTab[56442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:191
		// _ = "end of CoverTab[56442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:191
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:191
	// _ = "end of CoverTab[56380]"
}

// getUnknownBytes returns a *[]byte for the unknown fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:194
// It is the caller's responsibility to check whether the pointer is nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:194
// This function is specially designed to be inlineable.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:197
func (mi *MessageInfo) getUnknownBytes(p pointer) *[]byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:197
	_go_fuzz_dep_.CoverTab[56443]++
														if mi.unknownPtrKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:198
		_go_fuzz_dep_.CoverTab[56444]++
															return *p.Apply(mi.unknownOffset).BytesPtr()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:199
		// _ = "end of CoverTab[56444]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:200
		_go_fuzz_dep_.CoverTab[56445]++
															return p.Apply(mi.unknownOffset).Bytes()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:201
		// _ = "end of CoverTab[56445]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:202
	// _ = "end of CoverTab[56443]"
}

// mutableUnknownBytes returns a *[]byte for the unknown fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:205
// The returned pointer is guaranteed to not be nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:207
func (mi *MessageInfo) mutableUnknownBytes(p pointer) *[]byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:207
	_go_fuzz_dep_.CoverTab[56446]++
														if mi.unknownPtrKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:208
		_go_fuzz_dep_.CoverTab[56447]++
															bp := p.Apply(mi.unknownOffset).BytesPtr()
															if *bp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:210
			_go_fuzz_dep_.CoverTab[56449]++
																*bp = new([]byte)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:211
			// _ = "end of CoverTab[56449]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:212
			_go_fuzz_dep_.CoverTab[56450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:212
			// _ = "end of CoverTab[56450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:212
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:212
		// _ = "end of CoverTab[56447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:212
		_go_fuzz_dep_.CoverTab[56448]++
															return *bp
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:213
		// _ = "end of CoverTab[56448]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:214
		_go_fuzz_dep_.CoverTab[56451]++
															return p.Apply(mi.unknownOffset).Bytes()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:215
		// _ = "end of CoverTab[56451]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:216
	// _ = "end of CoverTab[56446]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:217
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_message.go:217
var _ = _go_fuzz_dep_.CoverTab
