// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:5
// Package defval marshals and unmarshals textual forms of default values.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:5
// This package handles both the form historically used in Go struct field tags
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:5
// and also the form used by google.protobuf.FieldDescriptorProto.default_value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:5
// since they differ in superficial ways.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
package defval

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:10
)

import (
	"fmt"
	"math"
	"strconv"

	ptext "google.golang.org/protobuf/internal/encoding/text"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Format is the serialization format used to represent the default value.
type Format int

const (
	_	Format	= iota

	// Descriptor uses the serialization format that protoc uses with the
	// google.protobuf.FieldDescriptorProto.default_value field.
	Descriptor

	// GoTag uses the historical serialization format in Go struct field tags.
	GoTag
)

// Unmarshal deserializes the default string s according to the given kind k.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:36
// When k is an enum, a list of enum value descriptors must be provided.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:38
func Unmarshal(s string, k protoreflect.Kind, evs protoreflect.EnumValueDescriptors, f Format) (protoreflect.Value, protoreflect.EnumValueDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:38
	_go_fuzz_dep_.CoverTab[52106]++
														switch k {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:40
		_go_fuzz_dep_.CoverTab[52108]++
															if f == GoTag {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:41
			_go_fuzz_dep_.CoverTab[52119]++
																switch s {
			case "1":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:43
				_go_fuzz_dep_.CoverTab[52120]++
																	return protoreflect.ValueOfBool(true), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:44
				// _ = "end of CoverTab[52120]"
			case "0":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:45
				_go_fuzz_dep_.CoverTab[52121]++
																	return protoreflect.ValueOfBool(false), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:46
				// _ = "end of CoverTab[52121]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:46
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:46
				_go_fuzz_dep_.CoverTab[52122]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:46
				// _ = "end of CoverTab[52122]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:47
			// _ = "end of CoverTab[52119]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:48
			_go_fuzz_dep_.CoverTab[52123]++
																switch s {
			case "true":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:50
				_go_fuzz_dep_.CoverTab[52124]++
																	return protoreflect.ValueOfBool(true), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:51
				// _ = "end of CoverTab[52124]"
			case "false":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:52
				_go_fuzz_dep_.CoverTab[52125]++
																	return protoreflect.ValueOfBool(false), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:53
				// _ = "end of CoverTab[52125]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:53
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:53
				_go_fuzz_dep_.CoverTab[52126]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:53
				// _ = "end of CoverTab[52126]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:54
			// _ = "end of CoverTab[52123]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:55
		// _ = "end of CoverTab[52108]"
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:56
		_go_fuzz_dep_.CoverTab[52109]++
															if f == GoTag {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:57
			_go_fuzz_dep_.CoverTab[52127]++

																if n, err := strconv.ParseInt(s, 10, 32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:59
				_go_fuzz_dep_.CoverTab[52128]++
																	if ev := evs.ByNumber(protoreflect.EnumNumber(n)); ev != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:60
					_go_fuzz_dep_.CoverTab[52129]++
																		return protoreflect.ValueOfEnum(ev.Number()), ev, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:61
					// _ = "end of CoverTab[52129]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:62
					_go_fuzz_dep_.CoverTab[52130]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:62
					// _ = "end of CoverTab[52130]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:62
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:62
				// _ = "end of CoverTab[52128]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:63
				_go_fuzz_dep_.CoverTab[52131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:63
				// _ = "end of CoverTab[52131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:63
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:63
			// _ = "end of CoverTab[52127]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:64
			_go_fuzz_dep_.CoverTab[52132]++

																ev := evs.ByName(protoreflect.Name(s))
																if ev != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:67
				_go_fuzz_dep_.CoverTab[52133]++
																	return protoreflect.ValueOfEnum(ev.Number()), ev, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:68
				// _ = "end of CoverTab[52133]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:69
				_go_fuzz_dep_.CoverTab[52134]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:69
				// _ = "end of CoverTab[52134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:69
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:69
			// _ = "end of CoverTab[52132]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:70
		// _ = "end of CoverTab[52109]"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:71
		_go_fuzz_dep_.CoverTab[52110]++
															if v, err := strconv.ParseInt(s, 10, 32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:72
			_go_fuzz_dep_.CoverTab[52135]++
																return protoreflect.ValueOfInt32(int32(v)), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:73
			// _ = "end of CoverTab[52135]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:74
			_go_fuzz_dep_.CoverTab[52136]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:74
			// _ = "end of CoverTab[52136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:74
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:74
		// _ = "end of CoverTab[52110]"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:75
		_go_fuzz_dep_.CoverTab[52111]++
															if v, err := strconv.ParseInt(s, 10, 64); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:76
			_go_fuzz_dep_.CoverTab[52137]++
																return protoreflect.ValueOfInt64(int64(v)), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:77
			// _ = "end of CoverTab[52137]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:78
			_go_fuzz_dep_.CoverTab[52138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:78
			// _ = "end of CoverTab[52138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:78
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:78
		// _ = "end of CoverTab[52111]"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:79
		_go_fuzz_dep_.CoverTab[52112]++
															if v, err := strconv.ParseUint(s, 10, 32); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:80
			_go_fuzz_dep_.CoverTab[52139]++
																return protoreflect.ValueOfUint32(uint32(v)), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:81
			// _ = "end of CoverTab[52139]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:82
			_go_fuzz_dep_.CoverTab[52140]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:82
			// _ = "end of CoverTab[52140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:82
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:82
		// _ = "end of CoverTab[52112]"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:83
		_go_fuzz_dep_.CoverTab[52113]++
															if v, err := strconv.ParseUint(s, 10, 64); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:84
			_go_fuzz_dep_.CoverTab[52141]++
																return protoreflect.ValueOfUint64(uint64(v)), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:85
			// _ = "end of CoverTab[52141]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:86
			_go_fuzz_dep_.CoverTab[52142]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:86
			// _ = "end of CoverTab[52142]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:86
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:86
		// _ = "end of CoverTab[52113]"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:87
		_go_fuzz_dep_.CoverTab[52114]++
															var v float64
															var err error
															switch s {
		case "-inf":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:91
			_go_fuzz_dep_.CoverTab[52143]++
																v = math.Inf(-1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:92
			// _ = "end of CoverTab[52143]"
		case "inf":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:93
			_go_fuzz_dep_.CoverTab[52144]++
																v = math.Inf(+1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:94
			// _ = "end of CoverTab[52144]"
		case "nan":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:95
			_go_fuzz_dep_.CoverTab[52145]++
																v = math.NaN()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:96
			// _ = "end of CoverTab[52145]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:97
			_go_fuzz_dep_.CoverTab[52146]++
																v, err = strconv.ParseFloat(s, 64)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:98
			// _ = "end of CoverTab[52146]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:99
		// _ = "end of CoverTab[52114]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:99
		_go_fuzz_dep_.CoverTab[52115]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:100
			_go_fuzz_dep_.CoverTab[52147]++
																if k == protoreflect.FloatKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:101
				_go_fuzz_dep_.CoverTab[52148]++
																	return protoreflect.ValueOfFloat32(float32(v)), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:102
				// _ = "end of CoverTab[52148]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:103
				_go_fuzz_dep_.CoverTab[52149]++
																	return protoreflect.ValueOfFloat64(float64(v)), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:104
				// _ = "end of CoverTab[52149]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:105
			// _ = "end of CoverTab[52147]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:106
			_go_fuzz_dep_.CoverTab[52150]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:106
			// _ = "end of CoverTab[52150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:106
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:106
		// _ = "end of CoverTab[52115]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:107
		_go_fuzz_dep_.CoverTab[52116]++

															return protoreflect.ValueOfString(s), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:109
		// _ = "end of CoverTab[52116]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:110
		_go_fuzz_dep_.CoverTab[52117]++
															if b, ok := unmarshalBytes(s); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:111
			_go_fuzz_dep_.CoverTab[52151]++
																return protoreflect.ValueOfBytes(b), nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:112
			// _ = "end of CoverTab[52151]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
			_go_fuzz_dep_.CoverTab[52152]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
			// _ = "end of CoverTab[52152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
		// _ = "end of CoverTab[52117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
		_go_fuzz_dep_.CoverTab[52118]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:113
		// _ = "end of CoverTab[52118]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:114
	// _ = "end of CoverTab[52106]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:114
	_go_fuzz_dep_.CoverTab[52107]++
														return protoreflect.Value{}, nil, errors.New("could not parse value for %v: %q", k, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:115
	// _ = "end of CoverTab[52107]"
}

// Marshal serializes v as the default string according to the given kind k.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:118
// When specifying the Descriptor format for an enum kind, the associated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:118
// enum value descriptor must be provided.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:121
func Marshal(v protoreflect.Value, ev protoreflect.EnumValueDescriptor, k protoreflect.Kind, f Format) (string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:121
	_go_fuzz_dep_.CoverTab[52153]++
														switch k {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:123
		_go_fuzz_dep_.CoverTab[52155]++
															if f == GoTag {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:124
			_go_fuzz_dep_.CoverTab[52163]++
																if v.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:125
				_go_fuzz_dep_.CoverTab[52164]++
																	return "1", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:126
				// _ = "end of CoverTab[52164]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:127
				_go_fuzz_dep_.CoverTab[52165]++
																	return "0", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:128
				// _ = "end of CoverTab[52165]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:129
			// _ = "end of CoverTab[52163]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:130
			_go_fuzz_dep_.CoverTab[52166]++
																if v.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:131
				_go_fuzz_dep_.CoverTab[52167]++
																	return "true", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:132
				// _ = "end of CoverTab[52167]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:133
				_go_fuzz_dep_.CoverTab[52168]++
																	return "false", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:134
				// _ = "end of CoverTab[52168]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:135
			// _ = "end of CoverTab[52166]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:136
		// _ = "end of CoverTab[52155]"
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:137
		_go_fuzz_dep_.CoverTab[52156]++
															if f == GoTag {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:138
			_go_fuzz_dep_.CoverTab[52169]++
																return strconv.FormatInt(int64(v.Enum()), 10), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:139
			// _ = "end of CoverTab[52169]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:140
			_go_fuzz_dep_.CoverTab[52170]++
																return string(ev.Name()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:141
			// _ = "end of CoverTab[52170]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:142
		// _ = "end of CoverTab[52156]"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:143
		_go_fuzz_dep_.CoverTab[52157]++
															return strconv.FormatInt(v.Int(), 10), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:144
		// _ = "end of CoverTab[52157]"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:145
		_go_fuzz_dep_.CoverTab[52158]++
															return strconv.FormatUint(v.Uint(), 10), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:146
		// _ = "end of CoverTab[52158]"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:147
		_go_fuzz_dep_.CoverTab[52159]++
															f := v.Float()
															switch {
		case math.IsInf(f, -1):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:150
			_go_fuzz_dep_.CoverTab[52171]++
																return "-inf", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:151
			// _ = "end of CoverTab[52171]"
		case math.IsInf(f, +1):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:152
			_go_fuzz_dep_.CoverTab[52172]++
																return "inf", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:153
			// _ = "end of CoverTab[52172]"
		case math.IsNaN(f):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:154
			_go_fuzz_dep_.CoverTab[52173]++
																return "nan", nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:155
			// _ = "end of CoverTab[52173]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:156
			_go_fuzz_dep_.CoverTab[52174]++
																if k == protoreflect.FloatKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:157
				_go_fuzz_dep_.CoverTab[52175]++
																	return strconv.FormatFloat(f, 'g', -1, 32), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:158
				// _ = "end of CoverTab[52175]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:159
				_go_fuzz_dep_.CoverTab[52176]++
																	return strconv.FormatFloat(f, 'g', -1, 64), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:160
				// _ = "end of CoverTab[52176]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:161
			// _ = "end of CoverTab[52174]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:162
		// _ = "end of CoverTab[52159]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:163
		_go_fuzz_dep_.CoverTab[52160]++

															return v.String(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:165
		// _ = "end of CoverTab[52160]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:166
		_go_fuzz_dep_.CoverTab[52161]++
															if s, ok := marshalBytes(v.Bytes()); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:167
			_go_fuzz_dep_.CoverTab[52177]++
																return s, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:168
			// _ = "end of CoverTab[52177]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
			_go_fuzz_dep_.CoverTab[52178]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
			// _ = "end of CoverTab[52178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
		// _ = "end of CoverTab[52161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
		_go_fuzz_dep_.CoverTab[52162]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:169
		// _ = "end of CoverTab[52162]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:170
	// _ = "end of CoverTab[52153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:170
	_go_fuzz_dep_.CoverTab[52154]++
														return "", errors.New("could not format value for %v: %v", k, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:171
	// _ = "end of CoverTab[52154]"
}

// unmarshalBytes deserializes bytes by applying C unescaping.
func unmarshalBytes(s string) ([]byte, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:175
	_go_fuzz_dep_.CoverTab[52179]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:178
	v, err := ptext.UnmarshalString(`"` + s + `"`)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:179
		_go_fuzz_dep_.CoverTab[52181]++
															return nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:180
		// _ = "end of CoverTab[52181]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:181
		_go_fuzz_dep_.CoverTab[52182]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:181
		// _ = "end of CoverTab[52182]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:181
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:181
	// _ = "end of CoverTab[52179]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:181
	_go_fuzz_dep_.CoverTab[52180]++
														return []byte(v), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:182
	// _ = "end of CoverTab[52180]"
}

// marshalBytes serializes bytes by using C escaping.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:185
// To match the exact output of protoc, this is identical to the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:185
// CEscape function in strutil.cc of the protoc source code.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:188
func marshalBytes(b []byte) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:188
	_go_fuzz_dep_.CoverTab[52183]++
														var s []byte
														for _, c := range b {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:190
		_go_fuzz_dep_.CoverTab[52185]++
															switch c {
		case '\n':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:192
			_go_fuzz_dep_.CoverTab[52186]++
																s = append(s, `\n`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:193
			// _ = "end of CoverTab[52186]"
		case '\r':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:194
			_go_fuzz_dep_.CoverTab[52187]++
																s = append(s, `\r`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:195
			// _ = "end of CoverTab[52187]"
		case '\t':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:196
			_go_fuzz_dep_.CoverTab[52188]++
																s = append(s, `\t`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:197
			// _ = "end of CoverTab[52188]"
		case '"':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:198
			_go_fuzz_dep_.CoverTab[52189]++
																s = append(s, `\"`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:199
			// _ = "end of CoverTab[52189]"
		case '\'':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:200
			_go_fuzz_dep_.CoverTab[52190]++
																s = append(s, `\'`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:201
			// _ = "end of CoverTab[52190]"
		case '\\':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:202
			_go_fuzz_dep_.CoverTab[52191]++
																s = append(s, `\\`...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:203
			// _ = "end of CoverTab[52191]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:204
			_go_fuzz_dep_.CoverTab[52192]++
																if printableASCII := c >= 0x20 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:205
				_go_fuzz_dep_.CoverTab[52193]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:205
				return c <= 0x7e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:205
				// _ = "end of CoverTab[52193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:205
			}(); printableASCII {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:205
				_go_fuzz_dep_.CoverTab[52194]++
																	s = append(s, c)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:206
				// _ = "end of CoverTab[52194]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:207
				_go_fuzz_dep_.CoverTab[52195]++
																	s = append(s, fmt.Sprintf(`\%03o`, c)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:208
				// _ = "end of CoverTab[52195]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:209
			// _ = "end of CoverTab[52192]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:210
		// _ = "end of CoverTab[52185]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:211
	// _ = "end of CoverTab[52183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:211
	_go_fuzz_dep_.CoverTab[52184]++
														return string(s), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:212
	// _ = "end of CoverTab[52184]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:213
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/defval/default.go:213
var _ = _go_fuzz_dep_.CoverTab
