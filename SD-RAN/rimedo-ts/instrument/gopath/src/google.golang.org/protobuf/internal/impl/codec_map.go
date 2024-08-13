// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:5
)

import (
	"reflect"
	"sort"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type mapInfo struct {
	goType		reflect.Type
	keyWiretag	uint64
	valWiretag	uint64
	keyFuncs	valueCoderFuncs
	valFuncs	valueCoderFuncs
	keyZero		protoreflect.Value
	keyKind		protoreflect.Kind
	conv		*mapConverter
}

func encoderFuncsForMap(fd protoreflect.FieldDescriptor, ft reflect.Type) (valueMessage *MessageInfo, funcs pointerCoderFuncs) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:27
	_go_fuzz_dep_.CoverTab[56193]++

													keyField := fd.MapKey()
													valField := fd.MapValue()
													keyWiretag := protowire.EncodeTag(1, wireTypes[keyField.Kind()])
													valWiretag := protowire.EncodeTag(2, wireTypes[valField.Kind()])
													keyFuncs := encoderFuncsForValue(keyField)
													valFuncs := encoderFuncsForValue(valField)
													conv := newMapConverter(ft, fd)

													mapi := &mapInfo{
		goType:		ft,
		keyWiretag:	keyWiretag,
		valWiretag:	valWiretag,
		keyFuncs:	keyFuncs,
		valFuncs:	valFuncs,
		keyZero:	keyField.Default(),
		keyKind:	keyField.Kind(),
		conv:		conv,
	}
	if valField.Kind() == protoreflect.MessageKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:47
		_go_fuzz_dep_.CoverTab[56198]++
														valueMessage = getMessageInfo(ft.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:48
		// _ = "end of CoverTab[56198]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:49
		_go_fuzz_dep_.CoverTab[56199]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:49
		// _ = "end of CoverTab[56199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:49
	// _ = "end of CoverTab[56193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:49
	_go_fuzz_dep_.CoverTab[56194]++

													funcs = pointerCoderFuncs{
		size: func(p pointer, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:52
			_go_fuzz_dep_.CoverTab[56200]++
															return sizeMap(p.AsValueOf(ft).Elem(), mapi, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:53
			// _ = "end of CoverTab[56200]"
		},
		marshal: func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:55
			_go_fuzz_dep_.CoverTab[56201]++
															return appendMap(b, p.AsValueOf(ft).Elem(), mapi, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:56
			// _ = "end of CoverTab[56201]"
		},
		unmarshal: func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:58
			_go_fuzz_dep_.CoverTab[56202]++
															mp := p.AsValueOf(ft)
															if mp.Elem().IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:60
				_go_fuzz_dep_.CoverTab[56204]++
																mp.Elem().Set(reflect.MakeMap(mapi.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:61
				// _ = "end of CoverTab[56204]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:62
				_go_fuzz_dep_.CoverTab[56205]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:62
				// _ = "end of CoverTab[56205]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:62
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:62
			// _ = "end of CoverTab[56202]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:62
			_go_fuzz_dep_.CoverTab[56203]++
															if f.mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:63
				_go_fuzz_dep_.CoverTab[56206]++
																return consumeMap(b, mp.Elem(), wtyp, mapi, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:64
				// _ = "end of CoverTab[56206]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:65
				_go_fuzz_dep_.CoverTab[56207]++
																return consumeMapOfMessage(b, mp.Elem(), wtyp, mapi, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:66
				// _ = "end of CoverTab[56207]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:67
			// _ = "end of CoverTab[56203]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:69
	// _ = "end of CoverTab[56194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:69
	_go_fuzz_dep_.CoverTab[56195]++
													switch valField.Kind() {
	case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:71
		_go_fuzz_dep_.CoverTab[56208]++
														funcs.merge = mergeMapOfMessage
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:72
		// _ = "end of CoverTab[56208]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:73
		_go_fuzz_dep_.CoverTab[56209]++
														funcs.merge = mergeMapOfBytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:74
		// _ = "end of CoverTab[56209]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:75
		_go_fuzz_dep_.CoverTab[56210]++
														funcs.merge = mergeMap
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:76
		// _ = "end of CoverTab[56210]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:77
	// _ = "end of CoverTab[56195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:77
	_go_fuzz_dep_.CoverTab[56196]++
													if valFuncs.isInit != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:78
		_go_fuzz_dep_.CoverTab[56211]++
														funcs.isInit = func(p pointer, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:79
			_go_fuzz_dep_.CoverTab[56212]++
															return isInitMap(p.AsValueOf(ft).Elem(), mapi, f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:80
			// _ = "end of CoverTab[56212]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:81
		// _ = "end of CoverTab[56211]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:82
		_go_fuzz_dep_.CoverTab[56213]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:82
		// _ = "end of CoverTab[56213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:82
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:82
	// _ = "end of CoverTab[56196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:82
	_go_fuzz_dep_.CoverTab[56197]++
													return valueMessage, funcs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:83
	// _ = "end of CoverTab[56197]"
}

const (
	mapKeyTagSize	= 1	// field 1, tag size 1.
	mapValTagSize	= 1	// field 2, tag size 2.
)

func sizeMap(mapv reflect.Value, mapi *mapInfo, f *coderFieldInfo, opts marshalOptions) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:91
	_go_fuzz_dep_.CoverTab[56214]++
													if mapv.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:92
		_go_fuzz_dep_.CoverTab[56217]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:93
		// _ = "end of CoverTab[56217]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:94
		_go_fuzz_dep_.CoverTab[56218]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:94
		// _ = "end of CoverTab[56218]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:94
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:94
	// _ = "end of CoverTab[56214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:94
	_go_fuzz_dep_.CoverTab[56215]++
													n := 0
													iter := mapRange(mapv)
													for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:97
		_go_fuzz_dep_.CoverTab[56219]++
														key := mapi.conv.keyConv.PBValueOf(iter.Key()).MapKey()
														keySize := mapi.keyFuncs.size(key.Value(), mapKeyTagSize, opts)
														var valSize int
														value := mapi.conv.valConv.PBValueOf(iter.Value())
														if f.mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:102
			_go_fuzz_dep_.CoverTab[56221]++
															valSize = mapi.valFuncs.size(value, mapValTagSize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:103
			// _ = "end of CoverTab[56221]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:104
			_go_fuzz_dep_.CoverTab[56222]++
															p := pointerOfValue(iter.Value())
															valSize += mapValTagSize
															valSize += protowire.SizeBytes(f.mi.sizePointer(p, opts))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:107
			// _ = "end of CoverTab[56222]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:108
		// _ = "end of CoverTab[56219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:108
		_go_fuzz_dep_.CoverTab[56220]++
														n += f.tagsize + protowire.SizeBytes(keySize+valSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:109
		// _ = "end of CoverTab[56220]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:110
	// _ = "end of CoverTab[56215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:110
	_go_fuzz_dep_.CoverTab[56216]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:111
	// _ = "end of CoverTab[56216]"
}

func consumeMap(b []byte, mapv reflect.Value, wtyp protowire.Type, mapi *mapInfo, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:114
	_go_fuzz_dep_.CoverTab[56223]++
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:115
		_go_fuzz_dep_.CoverTab[56227]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:116
		// _ = "end of CoverTab[56227]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:117
		_go_fuzz_dep_.CoverTab[56228]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:117
		// _ = "end of CoverTab[56228]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:117
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:117
	// _ = "end of CoverTab[56223]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:117
	_go_fuzz_dep_.CoverTab[56224]++
													b, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:119
		_go_fuzz_dep_.CoverTab[56229]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:120
		// _ = "end of CoverTab[56229]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:121
		_go_fuzz_dep_.CoverTab[56230]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:121
		// _ = "end of CoverTab[56230]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:121
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:121
	// _ = "end of CoverTab[56224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:121
	_go_fuzz_dep_.CoverTab[56225]++
													var (
		key	= mapi.keyZero
		val	= mapi.conv.valConv.New()
	)
	for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:126
		_go_fuzz_dep_.CoverTab[56231]++
														num, wtyp, n := protowire.ConsumeTag(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:128
			_go_fuzz_dep_.CoverTab[56236]++
															return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:129
			// _ = "end of CoverTab[56236]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:130
			_go_fuzz_dep_.CoverTab[56237]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:130
			// _ = "end of CoverTab[56237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:130
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:130
		// _ = "end of CoverTab[56231]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:130
		_go_fuzz_dep_.CoverTab[56232]++
														if num > protowire.MaxValidNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:131
			_go_fuzz_dep_.CoverTab[56238]++
															return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:132
			// _ = "end of CoverTab[56238]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:133
			_go_fuzz_dep_.CoverTab[56239]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:133
			// _ = "end of CoverTab[56239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:133
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:133
		// _ = "end of CoverTab[56232]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:133
		_go_fuzz_dep_.CoverTab[56233]++
														b = b[n:]
														err := errUnknown
														switch num {
		case genid.MapEntry_Key_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:137
			_go_fuzz_dep_.CoverTab[56240]++
															var v protoreflect.Value
															var o unmarshalOutput
															v, o, err = mapi.keyFuncs.unmarshal(b, key, num, wtyp, opts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:141
				_go_fuzz_dep_.CoverTab[56245]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:142
				// _ = "end of CoverTab[56245]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:143
				_go_fuzz_dep_.CoverTab[56246]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:143
				// _ = "end of CoverTab[56246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:143
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:143
			// _ = "end of CoverTab[56240]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:143
			_go_fuzz_dep_.CoverTab[56241]++
															key = v
															n = o.n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:145
			// _ = "end of CoverTab[56241]"
		case genid.MapEntry_Value_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:146
			_go_fuzz_dep_.CoverTab[56242]++
															var v protoreflect.Value
															var o unmarshalOutput
															v, o, err = mapi.valFuncs.unmarshal(b, val, num, wtyp, opts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:150
				_go_fuzz_dep_.CoverTab[56247]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:151
				// _ = "end of CoverTab[56247]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:152
				_go_fuzz_dep_.CoverTab[56248]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:152
				// _ = "end of CoverTab[56248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:152
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:152
			// _ = "end of CoverTab[56242]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:152
			_go_fuzz_dep_.CoverTab[56243]++
															val = v
															n = o.n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:154
			// _ = "end of CoverTab[56243]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:154
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:154
			_go_fuzz_dep_.CoverTab[56244]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:154
			// _ = "end of CoverTab[56244]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:155
		// _ = "end of CoverTab[56233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:155
		_go_fuzz_dep_.CoverTab[56234]++
														if err == errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:156
			_go_fuzz_dep_.CoverTab[56249]++
															n = protowire.ConsumeFieldValue(num, wtyp, b)
															if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:158
				_go_fuzz_dep_.CoverTab[56250]++
																return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:159
				// _ = "end of CoverTab[56250]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:160
				_go_fuzz_dep_.CoverTab[56251]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:160
				// _ = "end of CoverTab[56251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:160
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:160
			// _ = "end of CoverTab[56249]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:161
			_go_fuzz_dep_.CoverTab[56252]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:161
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:161
				_go_fuzz_dep_.CoverTab[56253]++
																return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:162
				// _ = "end of CoverTab[56253]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
				_go_fuzz_dep_.CoverTab[56254]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
				// _ = "end of CoverTab[56254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
			// _ = "end of CoverTab[56252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
		// _ = "end of CoverTab[56234]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:163
		_go_fuzz_dep_.CoverTab[56235]++
														b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:164
		// _ = "end of CoverTab[56235]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:165
	// _ = "end of CoverTab[56225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:165
	_go_fuzz_dep_.CoverTab[56226]++
													mapv.SetMapIndex(mapi.conv.keyConv.GoValueOf(key), mapi.conv.valConv.GoValueOf(val))
													out.n = n
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:168
	// _ = "end of CoverTab[56226]"
}

func consumeMapOfMessage(b []byte, mapv reflect.Value, wtyp protowire.Type, mapi *mapInfo, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:171
	_go_fuzz_dep_.CoverTab[56255]++
													if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:172
		_go_fuzz_dep_.CoverTab[56259]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:173
		// _ = "end of CoverTab[56259]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:174
		_go_fuzz_dep_.CoverTab[56260]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:174
		// _ = "end of CoverTab[56260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:174
	// _ = "end of CoverTab[56255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:174
	_go_fuzz_dep_.CoverTab[56256]++
													b, n := protowire.ConsumeBytes(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:176
		_go_fuzz_dep_.CoverTab[56261]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:177
		// _ = "end of CoverTab[56261]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:178
		_go_fuzz_dep_.CoverTab[56262]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:178
		// _ = "end of CoverTab[56262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:178
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:178
	// _ = "end of CoverTab[56256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:178
	_go_fuzz_dep_.CoverTab[56257]++
													var (
		key	= mapi.keyZero
		val	= reflect.New(f.mi.GoReflectType.Elem())
	)
	for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:183
		_go_fuzz_dep_.CoverTab[56263]++
														num, wtyp, n := protowire.ConsumeTag(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:185
			_go_fuzz_dep_.CoverTab[56268]++
															return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:186
			// _ = "end of CoverTab[56268]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:187
			_go_fuzz_dep_.CoverTab[56269]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:187
			// _ = "end of CoverTab[56269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:187
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:187
		// _ = "end of CoverTab[56263]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:187
		_go_fuzz_dep_.CoverTab[56264]++
														if num > protowire.MaxValidNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:188
			_go_fuzz_dep_.CoverTab[56270]++
															return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:189
			// _ = "end of CoverTab[56270]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:190
			_go_fuzz_dep_.CoverTab[56271]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:190
			// _ = "end of CoverTab[56271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:190
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:190
		// _ = "end of CoverTab[56264]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:190
		_go_fuzz_dep_.CoverTab[56265]++
														b = b[n:]
														err := errUnknown
														switch num {
		case 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:194
			_go_fuzz_dep_.CoverTab[56272]++
															var v protoreflect.Value
															var o unmarshalOutput
															v, o, err = mapi.keyFuncs.unmarshal(b, key, num, wtyp, opts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:198
				_go_fuzz_dep_.CoverTab[56278]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:199
				// _ = "end of CoverTab[56278]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:200
				_go_fuzz_dep_.CoverTab[56279]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:200
				// _ = "end of CoverTab[56279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:200
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:200
			// _ = "end of CoverTab[56272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:200
			_go_fuzz_dep_.CoverTab[56273]++
															key = v
															n = o.n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:202
			// _ = "end of CoverTab[56273]"
		case 2:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:203
			_go_fuzz_dep_.CoverTab[56274]++
															if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:204
				_go_fuzz_dep_.CoverTab[56280]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:205
				// _ = "end of CoverTab[56280]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:206
				_go_fuzz_dep_.CoverTab[56281]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:206
				// _ = "end of CoverTab[56281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:206
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:206
			// _ = "end of CoverTab[56274]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:206
			_go_fuzz_dep_.CoverTab[56275]++
															var v []byte
															v, n = protowire.ConsumeBytes(b)
															if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:209
				_go_fuzz_dep_.CoverTab[56282]++
																return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:210
				// _ = "end of CoverTab[56282]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:211
				_go_fuzz_dep_.CoverTab[56283]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:211
				// _ = "end of CoverTab[56283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:211
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:211
			// _ = "end of CoverTab[56275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:211
			_go_fuzz_dep_.CoverTab[56276]++
															var o unmarshalOutput
															o, err = f.mi.unmarshalPointer(v, pointerOfValue(val), 0, opts)
															if o.initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:214
				_go_fuzz_dep_.CoverTab[56284]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:217
				out.initialized = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:217
				// _ = "end of CoverTab[56284]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
				_go_fuzz_dep_.CoverTab[56285]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
				// _ = "end of CoverTab[56285]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
			// _ = "end of CoverTab[56276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
			_go_fuzz_dep_.CoverTab[56277]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:218
			// _ = "end of CoverTab[56277]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:219
		// _ = "end of CoverTab[56265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:219
		_go_fuzz_dep_.CoverTab[56266]++
														if err == errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:220
			_go_fuzz_dep_.CoverTab[56286]++
															n = protowire.ConsumeFieldValue(num, wtyp, b)
															if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:222
				_go_fuzz_dep_.CoverTab[56287]++
																return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:223
				// _ = "end of CoverTab[56287]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:224
				_go_fuzz_dep_.CoverTab[56288]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:224
				// _ = "end of CoverTab[56288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:224
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:224
			// _ = "end of CoverTab[56286]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:225
			_go_fuzz_dep_.CoverTab[56289]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:225
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:225
				_go_fuzz_dep_.CoverTab[56290]++
																return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:226
				// _ = "end of CoverTab[56290]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
				_go_fuzz_dep_.CoverTab[56291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
				// _ = "end of CoverTab[56291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
			// _ = "end of CoverTab[56289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
		// _ = "end of CoverTab[56266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:227
		_go_fuzz_dep_.CoverTab[56267]++
														b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:228
		// _ = "end of CoverTab[56267]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:229
	// _ = "end of CoverTab[56257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:229
	_go_fuzz_dep_.CoverTab[56258]++
													mapv.SetMapIndex(mapi.conv.keyConv.GoValueOf(key), val)
													out.n = n
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:232
	// _ = "end of CoverTab[56258]"
}

func appendMapItem(b []byte, keyrv, valrv reflect.Value, mapi *mapInfo, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:235
	_go_fuzz_dep_.CoverTab[56292]++
													if f.mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:236
		_go_fuzz_dep_.CoverTab[56293]++
														key := mapi.conv.keyConv.PBValueOf(keyrv).MapKey()
														val := mapi.conv.valConv.PBValueOf(valrv)
														size := 0
														size += mapi.keyFuncs.size(key.Value(), mapKeyTagSize, opts)
														size += mapi.valFuncs.size(val, mapValTagSize, opts)
														b = protowire.AppendVarint(b, uint64(size))
														b, err := mapi.keyFuncs.marshal(b, key.Value(), mapi.keyWiretag, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:244
			_go_fuzz_dep_.CoverTab[56295]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:245
			// _ = "end of CoverTab[56295]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:246
			_go_fuzz_dep_.CoverTab[56296]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:246
			// _ = "end of CoverTab[56296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:246
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:246
		// _ = "end of CoverTab[56293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:246
		_go_fuzz_dep_.CoverTab[56294]++
														return mapi.valFuncs.marshal(b, val, mapi.valWiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:247
		// _ = "end of CoverTab[56294]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:248
		_go_fuzz_dep_.CoverTab[56297]++
														key := mapi.conv.keyConv.PBValueOf(keyrv).MapKey()
														val := pointerOfValue(valrv)
														valSize := f.mi.sizePointer(val, opts)
														size := 0
														size += mapi.keyFuncs.size(key.Value(), mapKeyTagSize, opts)
														size += mapValTagSize + protowire.SizeBytes(valSize)
														b = protowire.AppendVarint(b, uint64(size))
														b, err := mapi.keyFuncs.marshal(b, key.Value(), mapi.keyWiretag, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:257
			_go_fuzz_dep_.CoverTab[56299]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:258
			// _ = "end of CoverTab[56299]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:259
			_go_fuzz_dep_.CoverTab[56300]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:259
			// _ = "end of CoverTab[56300]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:259
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:259
		// _ = "end of CoverTab[56297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:259
		_go_fuzz_dep_.CoverTab[56298]++
														b = protowire.AppendVarint(b, mapi.valWiretag)
														b = protowire.AppendVarint(b, uint64(valSize))
														return f.mi.marshalAppendPointer(b, val, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:262
		// _ = "end of CoverTab[56298]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:263
	// _ = "end of CoverTab[56292]"
}

func appendMap(b []byte, mapv reflect.Value, mapi *mapInfo, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:266
	_go_fuzz_dep_.CoverTab[56301]++
													if mapv.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:267
		_go_fuzz_dep_.CoverTab[56305]++
														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:268
		// _ = "end of CoverTab[56305]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:269
		_go_fuzz_dep_.CoverTab[56306]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:269
		// _ = "end of CoverTab[56306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:269
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:269
	// _ = "end of CoverTab[56301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:269
	_go_fuzz_dep_.CoverTab[56302]++
													if opts.Deterministic() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:270
		_go_fuzz_dep_.CoverTab[56307]++
														return appendMapDeterministic(b, mapv, mapi, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:271
		// _ = "end of CoverTab[56307]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:272
		_go_fuzz_dep_.CoverTab[56308]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:272
		// _ = "end of CoverTab[56308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:272
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:272
	// _ = "end of CoverTab[56302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:272
	_go_fuzz_dep_.CoverTab[56303]++
													iter := mapRange(mapv)
													for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:274
		_go_fuzz_dep_.CoverTab[56309]++
														var err error
														b = protowire.AppendVarint(b, f.wiretag)
														b, err = appendMapItem(b, iter.Key(), iter.Value(), mapi, f, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:278
			_go_fuzz_dep_.CoverTab[56310]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:279
			// _ = "end of CoverTab[56310]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:280
			_go_fuzz_dep_.CoverTab[56311]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:280
			// _ = "end of CoverTab[56311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:280
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:280
		// _ = "end of CoverTab[56309]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:281
	// _ = "end of CoverTab[56303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:281
	_go_fuzz_dep_.CoverTab[56304]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:282
	// _ = "end of CoverTab[56304]"
}

func appendMapDeterministic(b []byte, mapv reflect.Value, mapi *mapInfo, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:285
	_go_fuzz_dep_.CoverTab[56312]++
													keys := mapv.MapKeys()
													sort.Slice(keys, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:287
		_go_fuzz_dep_.CoverTab[56315]++
														switch keys[i].Kind() {
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:289
			_go_fuzz_dep_.CoverTab[56316]++
															return !keys[i].Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:290
				_go_fuzz_dep_.CoverTab[56322]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:290
				return keys[j].Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:290
				// _ = "end of CoverTab[56322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:290
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:290
			// _ = "end of CoverTab[56316]"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:291
			_go_fuzz_dep_.CoverTab[56317]++
															return keys[i].Int() < keys[j].Int()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:292
			// _ = "end of CoverTab[56317]"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:293
			_go_fuzz_dep_.CoverTab[56318]++
															return keys[i].Uint() < keys[j].Uint()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:294
			// _ = "end of CoverTab[56318]"
		case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:295
			_go_fuzz_dep_.CoverTab[56319]++
															return keys[i].Float() < keys[j].Float()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:296
			// _ = "end of CoverTab[56319]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:297
			_go_fuzz_dep_.CoverTab[56320]++
															return keys[i].String() < keys[j].String()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:298
			// _ = "end of CoverTab[56320]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:299
			_go_fuzz_dep_.CoverTab[56321]++
															panic("invalid kind: " + keys[i].Kind().String())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:300
			// _ = "end of CoverTab[56321]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:301
		// _ = "end of CoverTab[56315]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:302
	// _ = "end of CoverTab[56312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:302
	_go_fuzz_dep_.CoverTab[56313]++
													for _, key := range keys {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:303
		_go_fuzz_dep_.CoverTab[56323]++
														var err error
														b = protowire.AppendVarint(b, f.wiretag)
														b, err = appendMapItem(b, key, mapv.MapIndex(key), mapi, f, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:307
			_go_fuzz_dep_.CoverTab[56324]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:308
			// _ = "end of CoverTab[56324]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:309
			_go_fuzz_dep_.CoverTab[56325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:309
			// _ = "end of CoverTab[56325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:309
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:309
		// _ = "end of CoverTab[56323]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:310
	// _ = "end of CoverTab[56313]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:310
	_go_fuzz_dep_.CoverTab[56314]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:311
	// _ = "end of CoverTab[56314]"
}

func isInitMap(mapv reflect.Value, mapi *mapInfo, f *coderFieldInfo) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:314
	_go_fuzz_dep_.CoverTab[56326]++
													if mi := f.mi; mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:315
		_go_fuzz_dep_.CoverTab[56328]++
														mi.init()
														if !mi.needsInitCheck {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:317
			_go_fuzz_dep_.CoverTab[56330]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:318
			// _ = "end of CoverTab[56330]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:319
			_go_fuzz_dep_.CoverTab[56331]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:319
			// _ = "end of CoverTab[56331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:319
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:319
		// _ = "end of CoverTab[56328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:319
		_go_fuzz_dep_.CoverTab[56329]++
														iter := mapRange(mapv)
														for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:321
			_go_fuzz_dep_.CoverTab[56332]++
															val := pointerOfValue(iter.Value())
															if err := mi.checkInitializedPointer(val); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:323
				_go_fuzz_dep_.CoverTab[56333]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:324
				// _ = "end of CoverTab[56333]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:325
				_go_fuzz_dep_.CoverTab[56334]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:325
				// _ = "end of CoverTab[56334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:325
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:325
			// _ = "end of CoverTab[56332]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:326
		// _ = "end of CoverTab[56329]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:327
		_go_fuzz_dep_.CoverTab[56335]++
														iter := mapRange(mapv)
														for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:329
			_go_fuzz_dep_.CoverTab[56336]++
															val := mapi.conv.valConv.PBValueOf(iter.Value())
															if err := mapi.valFuncs.isInit(val); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:331
				_go_fuzz_dep_.CoverTab[56337]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:332
				// _ = "end of CoverTab[56337]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:333
				_go_fuzz_dep_.CoverTab[56338]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:333
				// _ = "end of CoverTab[56338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:333
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:333
			// _ = "end of CoverTab[56336]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:334
		// _ = "end of CoverTab[56335]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:335
	// _ = "end of CoverTab[56326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:335
	_go_fuzz_dep_.CoverTab[56327]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:336
	// _ = "end of CoverTab[56327]"
}

func mergeMap(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:339
	_go_fuzz_dep_.CoverTab[56339]++
													dstm := dst.AsValueOf(f.ft).Elem()
													srcm := src.AsValueOf(f.ft).Elem()
													if srcm.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:342
		_go_fuzz_dep_.CoverTab[56342]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:343
		// _ = "end of CoverTab[56342]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:344
		_go_fuzz_dep_.CoverTab[56343]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:344
		// _ = "end of CoverTab[56343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:344
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:344
	// _ = "end of CoverTab[56339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:344
	_go_fuzz_dep_.CoverTab[56340]++
													if dstm.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:345
		_go_fuzz_dep_.CoverTab[56344]++
														dstm.Set(reflect.MakeMap(f.ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:346
		// _ = "end of CoverTab[56344]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:347
		_go_fuzz_dep_.CoverTab[56345]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:347
		// _ = "end of CoverTab[56345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:347
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:347
	// _ = "end of CoverTab[56340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:347
	_go_fuzz_dep_.CoverTab[56341]++
													iter := mapRange(srcm)
													for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:349
		_go_fuzz_dep_.CoverTab[56346]++
														dstm.SetMapIndex(iter.Key(), iter.Value())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:350
		// _ = "end of CoverTab[56346]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:351
	// _ = "end of CoverTab[56341]"
}

func mergeMapOfBytes(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:354
	_go_fuzz_dep_.CoverTab[56347]++
													dstm := dst.AsValueOf(f.ft).Elem()
													srcm := src.AsValueOf(f.ft).Elem()
													if srcm.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:357
		_go_fuzz_dep_.CoverTab[56350]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:358
		// _ = "end of CoverTab[56350]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:359
		_go_fuzz_dep_.CoverTab[56351]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:359
		// _ = "end of CoverTab[56351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:359
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:359
	// _ = "end of CoverTab[56347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:359
	_go_fuzz_dep_.CoverTab[56348]++
													if dstm.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:360
		_go_fuzz_dep_.CoverTab[56352]++
														dstm.Set(reflect.MakeMap(f.ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:361
		// _ = "end of CoverTab[56352]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:362
		_go_fuzz_dep_.CoverTab[56353]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:362
		// _ = "end of CoverTab[56353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:362
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:362
	// _ = "end of CoverTab[56348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:362
	_go_fuzz_dep_.CoverTab[56349]++
													iter := mapRange(srcm)
													for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:364
		_go_fuzz_dep_.CoverTab[56354]++
														dstm.SetMapIndex(iter.Key(), reflect.ValueOf(append(emptyBuf[:], iter.Value().Bytes()...)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:365
		// _ = "end of CoverTab[56354]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:366
	// _ = "end of CoverTab[56349]"
}

func mergeMapOfMessage(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:369
	_go_fuzz_dep_.CoverTab[56355]++
													dstm := dst.AsValueOf(f.ft).Elem()
													srcm := src.AsValueOf(f.ft).Elem()
													if srcm.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:372
		_go_fuzz_dep_.CoverTab[56358]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:373
		// _ = "end of CoverTab[56358]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:374
		_go_fuzz_dep_.CoverTab[56359]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:374
		// _ = "end of CoverTab[56359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:374
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:374
	// _ = "end of CoverTab[56355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:374
	_go_fuzz_dep_.CoverTab[56356]++
													if dstm.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:375
		_go_fuzz_dep_.CoverTab[56360]++
														dstm.Set(reflect.MakeMap(f.ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:376
		// _ = "end of CoverTab[56360]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:377
		_go_fuzz_dep_.CoverTab[56361]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:377
		// _ = "end of CoverTab[56361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:377
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:377
	// _ = "end of CoverTab[56356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:377
	_go_fuzz_dep_.CoverTab[56357]++
													iter := mapRange(srcm)
													for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:379
		_go_fuzz_dep_.CoverTab[56362]++
														val := reflect.New(f.ft.Elem().Elem())
														if f.mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:381
			_go_fuzz_dep_.CoverTab[56364]++
															f.mi.mergePointer(pointerOfValue(val), pointerOfValue(iter.Value()), opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:382
			// _ = "end of CoverTab[56364]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:383
			_go_fuzz_dep_.CoverTab[56365]++
															opts.Merge(asMessage(val), asMessage(iter.Value()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:384
			// _ = "end of CoverTab[56365]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:385
		// _ = "end of CoverTab[56362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:385
		_go_fuzz_dep_.CoverTab[56363]++
														dstm.SetMapIndex(iter.Key(), val)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:386
		// _ = "end of CoverTab[56363]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:387
	// _ = "end of CoverTab[56357]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:388
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_map.go:388
var _ = _go_fuzz_dep_.CoverTab
