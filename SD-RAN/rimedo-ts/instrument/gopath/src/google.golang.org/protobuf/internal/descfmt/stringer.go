// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:5
// Package descfmt provides functionality to format descriptors.
package descfmt

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:6
)

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"google.golang.org/protobuf/internal/detrand"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type list interface {
	Len() int
	pragma.DoNotImplement
}

func FormatList(s fmt.State, r rune, vs list) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:25
	_go_fuzz_dep_.CoverTab[52196]++
													io.WriteString(s, formatListOpt(vs, true, r == 'v' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
		_go_fuzz_dep_.CoverTab[52197]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
		return (s.Flag('+') || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
			_go_fuzz_dep_.CoverTab[52198]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
			return s.Flag('#')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
			// _ = "end of CoverTab[52198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
		// _ = "end of CoverTab[52197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
	}()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:26
	// _ = "end of CoverTab[52196]"
}
func formatListOpt(vs list, isRoot, allowMulti bool) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:28
	_go_fuzz_dep_.CoverTab[52199]++
													start, end := "[", "]"
													if isRoot {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:30
		_go_fuzz_dep_.CoverTab[52201]++
														var name string
														switch vs.(type) {
		case protoreflect.Names:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:33
			_go_fuzz_dep_.CoverTab[52203]++
															name = "Names"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:34
			// _ = "end of CoverTab[52203]"
		case protoreflect.FieldNumbers:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:35
			_go_fuzz_dep_.CoverTab[52204]++
															name = "FieldNumbers"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:36
			// _ = "end of CoverTab[52204]"
		case protoreflect.FieldRanges:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:37
			_go_fuzz_dep_.CoverTab[52205]++
															name = "FieldRanges"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:38
			// _ = "end of CoverTab[52205]"
		case protoreflect.EnumRanges:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:39
			_go_fuzz_dep_.CoverTab[52206]++
															name = "EnumRanges"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:40
			// _ = "end of CoverTab[52206]"
		case protoreflect.FileImports:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:41
			_go_fuzz_dep_.CoverTab[52207]++
															name = "FileImports"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:42
			// _ = "end of CoverTab[52207]"
		case protoreflect.Descriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:43
			_go_fuzz_dep_.CoverTab[52208]++
															name = reflect.ValueOf(vs).MethodByName("Get").Type().Out(0).Name() + "s"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:44
			// _ = "end of CoverTab[52208]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:45
			_go_fuzz_dep_.CoverTab[52209]++
															name = reflect.ValueOf(vs).Elem().Type().Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:46
			// _ = "end of CoverTab[52209]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:47
		// _ = "end of CoverTab[52201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:47
		_go_fuzz_dep_.CoverTab[52202]++
														start, end = name+"{", "}"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:48
		// _ = "end of CoverTab[52202]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:49
		_go_fuzz_dep_.CoverTab[52210]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:49
		// _ = "end of CoverTab[52210]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:49
	// _ = "end of CoverTab[52199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:49
	_go_fuzz_dep_.CoverTab[52200]++

													var ss []string
													switch vs := vs.(type) {
	case protoreflect.Names:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:53
		_go_fuzz_dep_.CoverTab[52211]++
														for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:54
			_go_fuzz_dep_.CoverTab[52223]++
															ss = append(ss, fmt.Sprint(vs.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:55
			// _ = "end of CoverTab[52223]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:56
		// _ = "end of CoverTab[52211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:56
		_go_fuzz_dep_.CoverTab[52212]++
														return start + joinStrings(ss, false) + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:57
		// _ = "end of CoverTab[52212]"
	case protoreflect.FieldNumbers:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:58
		_go_fuzz_dep_.CoverTab[52213]++
														for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:59
			_go_fuzz_dep_.CoverTab[52224]++
															ss = append(ss, fmt.Sprint(vs.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:60
			// _ = "end of CoverTab[52224]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:61
		// _ = "end of CoverTab[52213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:61
		_go_fuzz_dep_.CoverTab[52214]++
														return start + joinStrings(ss, false) + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:62
		// _ = "end of CoverTab[52214]"
	case protoreflect.FieldRanges:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:63
		_go_fuzz_dep_.CoverTab[52215]++
														for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:64
			_go_fuzz_dep_.CoverTab[52225]++
															r := vs.Get(i)
															if r[0]+1 == r[1] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:66
				_go_fuzz_dep_.CoverTab[52226]++
																ss = append(ss, fmt.Sprintf("%d", r[0]))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:67
				// _ = "end of CoverTab[52226]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:68
				_go_fuzz_dep_.CoverTab[52227]++
																ss = append(ss, fmt.Sprintf("%d:%d", r[0], r[1]))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:69
				// _ = "end of CoverTab[52227]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:70
			// _ = "end of CoverTab[52225]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:71
		// _ = "end of CoverTab[52215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:71
		_go_fuzz_dep_.CoverTab[52216]++
														return start + joinStrings(ss, false) + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:72
		// _ = "end of CoverTab[52216]"
	case protoreflect.EnumRanges:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:73
		_go_fuzz_dep_.CoverTab[52217]++
														for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:74
			_go_fuzz_dep_.CoverTab[52228]++
															r := vs.Get(i)
															if r[0] == r[1] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:76
				_go_fuzz_dep_.CoverTab[52229]++
																ss = append(ss, fmt.Sprintf("%d", r[0]))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:77
				// _ = "end of CoverTab[52229]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:78
				_go_fuzz_dep_.CoverTab[52230]++
																ss = append(ss, fmt.Sprintf("%d:%d", r[0], int64(r[1])+1))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:79
				// _ = "end of CoverTab[52230]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:80
			// _ = "end of CoverTab[52228]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:81
		// _ = "end of CoverTab[52217]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:81
		_go_fuzz_dep_.CoverTab[52218]++
														return start + joinStrings(ss, false) + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:82
		// _ = "end of CoverTab[52218]"
	case protoreflect.FileImports:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:83
		_go_fuzz_dep_.CoverTab[52219]++
														for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:84
			_go_fuzz_dep_.CoverTab[52231]++
															var rs records
															rs.Append(reflect.ValueOf(vs.Get(i)), "Path", "Package", "IsPublic", "IsWeak")
															ss = append(ss, "{"+rs.Join()+"}")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:87
			// _ = "end of CoverTab[52231]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:88
		// _ = "end of CoverTab[52219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:88
		_go_fuzz_dep_.CoverTab[52220]++
														return start + joinStrings(ss, allowMulti) + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:89
		// _ = "end of CoverTab[52220]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:90
		_go_fuzz_dep_.CoverTab[52221]++
														_, isEnumValue := vs.(protoreflect.EnumValueDescriptors)
														for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:92
			_go_fuzz_dep_.CoverTab[52232]++
															m := reflect.ValueOf(vs).MethodByName("Get")
															v := m.Call([]reflect.Value{reflect.ValueOf(i)})[0].Interface()
															ss = append(ss, formatDescOpt(v.(protoreflect.Descriptor), false, allowMulti && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:95
				_go_fuzz_dep_.CoverTab[52233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:95
				return !isEnumValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:95
				// _ = "end of CoverTab[52233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:95
			}()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:95
			// _ = "end of CoverTab[52232]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:96
		// _ = "end of CoverTab[52221]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:96
		_go_fuzz_dep_.CoverTab[52222]++
														return start + joinStrings(ss, allowMulti && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:97
			_go_fuzz_dep_.CoverTab[52234]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:97
			return isEnumValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:97
			// _ = "end of CoverTab[52234]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:97
		}()) + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:97
		// _ = "end of CoverTab[52222]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:98
	// _ = "end of CoverTab[52200]"
}

// descriptorAccessors is a list of accessors to print for each descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:101
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:101
// Do not print all accessors since some contain redundant information,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:101
// while others are pointers that we do not want to follow since the descriptor
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:101
// is actually a cyclic graph.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:101
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:101
// Using a list allows us to print the accessors in a sensible order.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:108
var descriptorAccessors = map[reflect.Type][]string{
	reflect.TypeOf((*protoreflect.FileDescriptor)(nil)).Elem():		{"Path", "Package", "Imports", "Messages", "Enums", "Extensions", "Services"},
	reflect.TypeOf((*protoreflect.MessageDescriptor)(nil)).Elem():		{"IsMapEntry", "Fields", "Oneofs", "ReservedNames", "ReservedRanges", "RequiredNumbers", "ExtensionRanges", "Messages", "Enums", "Extensions"},
	reflect.TypeOf((*protoreflect.FieldDescriptor)(nil)).Elem():		{"Number", "Cardinality", "Kind", "HasJSONName", "JSONName", "HasPresence", "IsExtension", "IsPacked", "IsWeak", "IsList", "IsMap", "MapKey", "MapValue", "HasDefault", "Default", "ContainingOneof", "ContainingMessage", "Message", "Enum"},
	reflect.TypeOf((*protoreflect.OneofDescriptor)(nil)).Elem():		{"Fields"},
	reflect.TypeOf((*protoreflect.EnumDescriptor)(nil)).Elem():		{"Values", "ReservedNames", "ReservedRanges"},
	reflect.TypeOf((*protoreflect.EnumValueDescriptor)(nil)).Elem():	{"Number"},
	reflect.TypeOf((*protoreflect.ServiceDescriptor)(nil)).Elem():		{"Methods"},
	reflect.TypeOf((*protoreflect.MethodDescriptor)(nil)).Elem():		{"Input", "Output", "IsStreamingClient", "IsStreamingServer"},
}

func FormatDesc(s fmt.State, r rune, t protoreflect.Descriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:119
	_go_fuzz_dep_.CoverTab[52235]++
													io.WriteString(s, formatDescOpt(t, true, r == 'v' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
		_go_fuzz_dep_.CoverTab[52236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
		return (s.Flag('+') || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
			_go_fuzz_dep_.CoverTab[52237]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
			return s.Flag('#')
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
			// _ = "end of CoverTab[52237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
		// _ = "end of CoverTab[52236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
	}()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:120
	// _ = "end of CoverTab[52235]"
}
func formatDescOpt(t protoreflect.Descriptor, isRoot, allowMulti bool) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:122
	_go_fuzz_dep_.CoverTab[52238]++
													rv := reflect.ValueOf(t)
													rt := rv.MethodByName("ProtoType").Type().In(0)

													start, end := "{", "}"
													if isRoot {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:127
		_go_fuzz_dep_.CoverTab[52241]++
														start = rt.Name() + "{"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:128
		// _ = "end of CoverTab[52241]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:129
		_go_fuzz_dep_.CoverTab[52242]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:129
		// _ = "end of CoverTab[52242]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:129
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:129
	// _ = "end of CoverTab[52238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:129
	_go_fuzz_dep_.CoverTab[52239]++

													_, isFile := t.(protoreflect.FileDescriptor)
													rs := records{allowMulti: allowMulti}
													if t.IsPlaceholder() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:133
		_go_fuzz_dep_.CoverTab[52243]++
														if isFile {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:134
			_go_fuzz_dep_.CoverTab[52244]++
															rs.Append(rv, "Path", "Package", "IsPlaceholder")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:135
			// _ = "end of CoverTab[52244]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:136
			_go_fuzz_dep_.CoverTab[52245]++
															rs.Append(rv, "FullName", "IsPlaceholder")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:137
			// _ = "end of CoverTab[52245]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:138
		// _ = "end of CoverTab[52243]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:139
		_go_fuzz_dep_.CoverTab[52246]++
														switch {
		case isFile:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:141
			_go_fuzz_dep_.CoverTab[52249]++
															rs.Append(rv, "Syntax")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:142
			// _ = "end of CoverTab[52249]"
		case isRoot:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:143
			_go_fuzz_dep_.CoverTab[52250]++
															rs.Append(rv, "Syntax", "FullName")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:144
			// _ = "end of CoverTab[52250]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:145
			_go_fuzz_dep_.CoverTab[52251]++
															rs.Append(rv, "Name")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:146
			// _ = "end of CoverTab[52251]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:147
		// _ = "end of CoverTab[52246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:147
		_go_fuzz_dep_.CoverTab[52247]++
														switch t := t.(type) {
		case protoreflect.FieldDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:149
			_go_fuzz_dep_.CoverTab[52252]++
															for _, s := range descriptorAccessors[rt] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:150
				_go_fuzz_dep_.CoverTab[52256]++
																switch s {
				case "MapKey":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:152
					_go_fuzz_dep_.CoverTab[52257]++
																	if k := t.MapKey(); k != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:153
						_go_fuzz_dep_.CoverTab[52263]++
																		rs.recs = append(rs.recs, [2]string{"MapKey", k.Kind().String()})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:154
						// _ = "end of CoverTab[52263]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:155
						_go_fuzz_dep_.CoverTab[52264]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:155
						// _ = "end of CoverTab[52264]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:155
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:155
					// _ = "end of CoverTab[52257]"
				case "MapValue":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:156
					_go_fuzz_dep_.CoverTab[52258]++
																	if v := t.MapValue(); v != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:157
						_go_fuzz_dep_.CoverTab[52265]++
																		switch v.Kind() {
						case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:159
							_go_fuzz_dep_.CoverTab[52266]++
																			rs.recs = append(rs.recs, [2]string{"MapValue", string(v.Enum().FullName())})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:160
							// _ = "end of CoverTab[52266]"
						case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:161
							_go_fuzz_dep_.CoverTab[52267]++
																			rs.recs = append(rs.recs, [2]string{"MapValue", string(v.Message().FullName())})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:162
							// _ = "end of CoverTab[52267]"
						default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:163
							_go_fuzz_dep_.CoverTab[52268]++
																			rs.recs = append(rs.recs, [2]string{"MapValue", v.Kind().String()})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:164
							// _ = "end of CoverTab[52268]"
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:165
						// _ = "end of CoverTab[52265]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:166
						_go_fuzz_dep_.CoverTab[52269]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:166
						// _ = "end of CoverTab[52269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:166
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:166
					// _ = "end of CoverTab[52258]"
				case "ContainingOneof":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:167
					_go_fuzz_dep_.CoverTab[52259]++
																	if od := t.ContainingOneof(); od != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:168
						_go_fuzz_dep_.CoverTab[52270]++
																		rs.recs = append(rs.recs, [2]string{"Oneof", string(od.Name())})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:169
						// _ = "end of CoverTab[52270]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:170
						_go_fuzz_dep_.CoverTab[52271]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:170
						// _ = "end of CoverTab[52271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:170
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:170
					// _ = "end of CoverTab[52259]"
				case "ContainingMessage":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:171
					_go_fuzz_dep_.CoverTab[52260]++
																	if t.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:172
						_go_fuzz_dep_.CoverTab[52272]++
																		rs.recs = append(rs.recs, [2]string{"Extendee", string(t.ContainingMessage().FullName())})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:173
						// _ = "end of CoverTab[52272]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:174
						_go_fuzz_dep_.CoverTab[52273]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:174
						// _ = "end of CoverTab[52273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:174
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:174
					// _ = "end of CoverTab[52260]"
				case "Message":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:175
					_go_fuzz_dep_.CoverTab[52261]++
																	if !t.IsMap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:176
						_go_fuzz_dep_.CoverTab[52274]++
																		rs.Append(rv, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:177
						// _ = "end of CoverTab[52274]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:178
						_go_fuzz_dep_.CoverTab[52275]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:178
						// _ = "end of CoverTab[52275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:178
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:178
					// _ = "end of CoverTab[52261]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:179
					_go_fuzz_dep_.CoverTab[52262]++
																	rs.Append(rv, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:180
					// _ = "end of CoverTab[52262]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:181
				// _ = "end of CoverTab[52256]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:182
			// _ = "end of CoverTab[52252]"
		case protoreflect.OneofDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:183
			_go_fuzz_dep_.CoverTab[52253]++
															var ss []string
															fs := t.Fields()
															for i := 0; i < fs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:186
				_go_fuzz_dep_.CoverTab[52276]++
																ss = append(ss, string(fs.Get(i).Name()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:187
				// _ = "end of CoverTab[52276]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:188
			// _ = "end of CoverTab[52253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:188
			_go_fuzz_dep_.CoverTab[52254]++
															if len(ss) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:189
				_go_fuzz_dep_.CoverTab[52277]++
																rs.recs = append(rs.recs, [2]string{"Fields", "[" + joinStrings(ss, false) + "]"})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:190
				// _ = "end of CoverTab[52277]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:191
				_go_fuzz_dep_.CoverTab[52278]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:191
				// _ = "end of CoverTab[52278]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:191
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:191
			// _ = "end of CoverTab[52254]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:192
			_go_fuzz_dep_.CoverTab[52255]++
															rs.Append(rv, descriptorAccessors[rt]...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:193
			// _ = "end of CoverTab[52255]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:194
		// _ = "end of CoverTab[52247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:194
		_go_fuzz_dep_.CoverTab[52248]++
														if rv.MethodByName("GoType").IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:195
			_go_fuzz_dep_.CoverTab[52279]++
															rs.Append(rv, "GoType")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:196
			// _ = "end of CoverTab[52279]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:197
			_go_fuzz_dep_.CoverTab[52280]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:197
			// _ = "end of CoverTab[52280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:197
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:197
		// _ = "end of CoverTab[52248]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:198
	// _ = "end of CoverTab[52239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:198
	_go_fuzz_dep_.CoverTab[52240]++
													return start + rs.Join() + end
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:199
	// _ = "end of CoverTab[52240]"
}

type records struct {
	recs		[][2]string
	allowMulti	bool
}

func (rs *records) Append(v reflect.Value, accessors ...string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:207
	_go_fuzz_dep_.CoverTab[52281]++
													for _, a := range accessors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:208
		_go_fuzz_dep_.CoverTab[52282]++
														var rv reflect.Value
														if m := v.MethodByName(a); m.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:210
			_go_fuzz_dep_.CoverTab[52291]++
															rv = m.Call(nil)[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:211
			// _ = "end of CoverTab[52291]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:212
			_go_fuzz_dep_.CoverTab[52292]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:212
			// _ = "end of CoverTab[52292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:212
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:212
		// _ = "end of CoverTab[52282]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:212
		_go_fuzz_dep_.CoverTab[52283]++
														if v.Kind() == reflect.Struct && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:213
			_go_fuzz_dep_.CoverTab[52293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:213
			return !rv.IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:213
			// _ = "end of CoverTab[52293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:213
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:213
			_go_fuzz_dep_.CoverTab[52294]++
															rv = v.FieldByName(a)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:214
			// _ = "end of CoverTab[52294]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:215
			_go_fuzz_dep_.CoverTab[52295]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:215
			// _ = "end of CoverTab[52295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:215
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:215
		// _ = "end of CoverTab[52283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:215
		_go_fuzz_dep_.CoverTab[52284]++
														if !rv.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:216
			_go_fuzz_dep_.CoverTab[52296]++
															panic(fmt.Sprintf("unknown accessor: %v.%s", v.Type(), a))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:217
			// _ = "end of CoverTab[52296]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:218
			_go_fuzz_dep_.CoverTab[52297]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:218
			// _ = "end of CoverTab[52297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:218
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:218
		// _ = "end of CoverTab[52284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:218
		_go_fuzz_dep_.CoverTab[52285]++
														if _, ok := rv.Interface().(protoreflect.Value); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:219
			_go_fuzz_dep_.CoverTab[52298]++
															rv = rv.MethodByName("Interface").Call(nil)[0]
															if !rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:221
				_go_fuzz_dep_.CoverTab[52299]++
																rv = rv.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:222
				// _ = "end of CoverTab[52299]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:223
				_go_fuzz_dep_.CoverTab[52300]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:223
				// _ = "end of CoverTab[52300]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:223
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:223
			// _ = "end of CoverTab[52298]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:224
			_go_fuzz_dep_.CoverTab[52301]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:224
			// _ = "end of CoverTab[52301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:224
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:224
		// _ = "end of CoverTab[52285]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:224
		_go_fuzz_dep_.CoverTab[52286]++

		// Ignore zero values.
		var isZero bool
		switch rv.Kind() {
		case reflect.Interface, reflect.Slice:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:229
			_go_fuzz_dep_.CoverTab[52302]++
															isZero = rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:230
			// _ = "end of CoverTab[52302]"
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:231
			_go_fuzz_dep_.CoverTab[52303]++
															isZero = rv.Bool() == false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:232
			// _ = "end of CoverTab[52303]"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:233
			_go_fuzz_dep_.CoverTab[52304]++
															isZero = rv.Int() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:234
			// _ = "end of CoverTab[52304]"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:235
			_go_fuzz_dep_.CoverTab[52305]++
															isZero = rv.Uint() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:236
			// _ = "end of CoverTab[52305]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:237
			_go_fuzz_dep_.CoverTab[52306]++
															isZero = rv.String() == ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:238
			// _ = "end of CoverTab[52306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:238
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:238
			_go_fuzz_dep_.CoverTab[52307]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:238
			// _ = "end of CoverTab[52307]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:239
		// _ = "end of CoverTab[52286]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:239
		_go_fuzz_dep_.CoverTab[52287]++
														if n, ok := rv.Interface().(list); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:240
			_go_fuzz_dep_.CoverTab[52308]++
															isZero = n.Len() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:241
			// _ = "end of CoverTab[52308]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:242
			_go_fuzz_dep_.CoverTab[52309]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:242
			// _ = "end of CoverTab[52309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:242
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:242
		// _ = "end of CoverTab[52287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:242
		_go_fuzz_dep_.CoverTab[52288]++
														if isZero {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:243
			_go_fuzz_dep_.CoverTab[52310]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:244
			// _ = "end of CoverTab[52310]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:245
			_go_fuzz_dep_.CoverTab[52311]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:245
			// _ = "end of CoverTab[52311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:245
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:245
		// _ = "end of CoverTab[52288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:245
		_go_fuzz_dep_.CoverTab[52289]++

		// Format the value.
		var s string
		v := rv.Interface()
		switch v := v.(type) {
		case list:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:251
			_go_fuzz_dep_.CoverTab[52312]++
															s = formatListOpt(v, false, rs.allowMulti)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:252
			// _ = "end of CoverTab[52312]"
		case protoreflect.FieldDescriptor, protoreflect.OneofDescriptor, protoreflect.EnumValueDescriptor, protoreflect.MethodDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:253
			_go_fuzz_dep_.CoverTab[52313]++
															s = string(v.(protoreflect.Descriptor).Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:254
			// _ = "end of CoverTab[52313]"
		case protoreflect.Descriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:255
			_go_fuzz_dep_.CoverTab[52314]++
															s = string(v.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:256
			// _ = "end of CoverTab[52314]"
		case string:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:257
			_go_fuzz_dep_.CoverTab[52315]++
															s = strconv.Quote(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:258
			// _ = "end of CoverTab[52315]"
		case []byte:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:259
			_go_fuzz_dep_.CoverTab[52316]++
															s = fmt.Sprintf("%q", v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:260
			// _ = "end of CoverTab[52316]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:261
			_go_fuzz_dep_.CoverTab[52317]++
															s = fmt.Sprint(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:262
			// _ = "end of CoverTab[52317]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:263
		// _ = "end of CoverTab[52289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:263
		_go_fuzz_dep_.CoverTab[52290]++
														rs.recs = append(rs.recs, [2]string{a, s})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:264
		// _ = "end of CoverTab[52290]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:265
	// _ = "end of CoverTab[52281]"
}

func (rs *records) Join() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:268
	_go_fuzz_dep_.CoverTab[52318]++
													var ss []string

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:272
	if !rs.allowMulti {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:272
		_go_fuzz_dep_.CoverTab[52322]++
														for _, r := range rs.recs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:273
			_go_fuzz_dep_.CoverTab[52324]++
															ss = append(ss, r[0]+formatColon(0)+r[1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:274
			// _ = "end of CoverTab[52324]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:275
		// _ = "end of CoverTab[52322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:275
		_go_fuzz_dep_.CoverTab[52323]++
														return joinStrings(ss, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:276
		// _ = "end of CoverTab[52323]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:277
		_go_fuzz_dep_.CoverTab[52325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:277
		// _ = "end of CoverTab[52325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:277
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:277
	// _ = "end of CoverTab[52318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:277
	_go_fuzz_dep_.CoverTab[52319]++

	// In allowMulti line mode, align single line records for more readable output.
	var maxLen int
	flush := func(i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:281
		_go_fuzz_dep_.CoverTab[52326]++
														for _, r := range rs.recs[len(ss):i] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:282
			_go_fuzz_dep_.CoverTab[52328]++
															ss = append(ss, r[0]+formatColon(maxLen-len(r[0]))+r[1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:283
			// _ = "end of CoverTab[52328]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:284
		// _ = "end of CoverTab[52326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:284
		_go_fuzz_dep_.CoverTab[52327]++
														maxLen = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:285
		// _ = "end of CoverTab[52327]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:286
	// _ = "end of CoverTab[52319]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:286
	_go_fuzz_dep_.CoverTab[52320]++
													for i, r := range rs.recs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:287
		_go_fuzz_dep_.CoverTab[52329]++
														if isMulti := strings.Contains(r[1], "\n"); isMulti {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:288
			_go_fuzz_dep_.CoverTab[52330]++
															flush(i)
															ss = append(ss, r[0]+formatColon(0)+strings.Join(strings.Split(r[1], "\n"), "\n\t"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:290
			// _ = "end of CoverTab[52330]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:291
			_go_fuzz_dep_.CoverTab[52331]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:291
			if maxLen < len(r[0]) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:291
				_go_fuzz_dep_.CoverTab[52332]++
																maxLen = len(r[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:292
				// _ = "end of CoverTab[52332]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:293
				_go_fuzz_dep_.CoverTab[52333]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:293
				// _ = "end of CoverTab[52333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:293
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:293
			// _ = "end of CoverTab[52331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:293
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:293
		// _ = "end of CoverTab[52329]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:294
	// _ = "end of CoverTab[52320]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:294
	_go_fuzz_dep_.CoverTab[52321]++
													flush(len(rs.recs))
													return joinStrings(ss, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:296
	// _ = "end of CoverTab[52321]"
}

func formatColon(padding int) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:299
	_go_fuzz_dep_.CoverTab[52334]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:303
	if detrand.Bool() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:303
		_go_fuzz_dep_.CoverTab[52335]++
														return ":" + strings.Repeat(" ", 1+padding)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:304
		// _ = "end of CoverTab[52335]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:305
		_go_fuzz_dep_.CoverTab[52336]++
														return ":" + strings.Repeat(" ", 1+padding)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:306
		// _ = "end of CoverTab[52336]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:307
	// _ = "end of CoverTab[52334]"
}

func joinStrings(ss []string, isMulti bool) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:310
	_go_fuzz_dep_.CoverTab[52337]++
													if len(ss) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:311
		_go_fuzz_dep_.CoverTab[52340]++
														return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:312
		// _ = "end of CoverTab[52340]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:313
		_go_fuzz_dep_.CoverTab[52341]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:313
		// _ = "end of CoverTab[52341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:313
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:313
	// _ = "end of CoverTab[52337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:313
	_go_fuzz_dep_.CoverTab[52338]++
													if isMulti {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:314
		_go_fuzz_dep_.CoverTab[52342]++
														return "\n\t" + strings.Join(ss, "\n\t") + "\n"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:315
		// _ = "end of CoverTab[52342]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:316
		_go_fuzz_dep_.CoverTab[52343]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:316
		// _ = "end of CoverTab[52343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:316
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:316
	// _ = "end of CoverTab[52338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:316
	_go_fuzz_dep_.CoverTab[52339]++
													return strings.Join(ss, ", ")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:317
	// _ = "end of CoverTab[52339]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:318
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/descfmt/stringer.go:318
var _ = _go_fuzz_dep_.CoverTab
