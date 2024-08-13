// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:5
// Package tag marshals and unmarshals the legacy struct tags as generated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:5
// by historical versions of protoc-gen-go.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
package tag

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:7
)

import (
	"reflect"
	"strconv"
	"strings"

	"google.golang.org/protobuf/internal/encoding/defval"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var byteType = reflect.TypeOf(byte(0))

// Unmarshal decodes the tag into a prototype.Field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// The goType is needed to determine the original protoreflect.Kind since the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// tag does not record sufficient information to determine that.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// The type is the underlying field type (e.g., a repeated field may be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// represented by []T, but the Go type passed in is just T).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// A list of enum value descriptors must be provided for enum fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// This does not populate the Enum or Message (except for weak message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:22
// This function is a best effort attempt; parsing errors are ignored.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:32
func Unmarshal(tag string, goType reflect.Type, evs protoreflect.EnumValueDescriptors) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:32
	_go_fuzz_dep_.CoverTab[53491]++
													f := new(filedesc.Field)
													f.L0.ParentFile = filedesc.SurrogateProto2
													for len(tag) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:35
		_go_fuzz_dep_.CoverTab[53494]++
														i := strings.IndexByte(tag, ',')
														if i < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:37
			_go_fuzz_dep_.CoverTab[53497]++
															i = len(tag)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:38
			// _ = "end of CoverTab[53497]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:39
			_go_fuzz_dep_.CoverTab[53498]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:39
			// _ = "end of CoverTab[53498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:39
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:39
		// _ = "end of CoverTab[53494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:39
		_go_fuzz_dep_.CoverTab[53495]++
														switch s := tag[:i]; {
		case strings.HasPrefix(s, "name="):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:41
			_go_fuzz_dep_.CoverTab[53499]++
															f.L0.FullName = protoreflect.FullName(s[len("name="):])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:42
			// _ = "end of CoverTab[53499]"
		case strings.Trim(s, "0123456789") == "":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:43
			_go_fuzz_dep_.CoverTab[53500]++
															n, _ := strconv.ParseUint(s, 10, 32)
															f.L1.Number = protoreflect.FieldNumber(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:45
			// _ = "end of CoverTab[53500]"
		case s == "opt":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:46
			_go_fuzz_dep_.CoverTab[53501]++
															f.L1.Cardinality = protoreflect.Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:47
			// _ = "end of CoverTab[53501]"
		case s == "req":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:48
			_go_fuzz_dep_.CoverTab[53502]++
															f.L1.Cardinality = protoreflect.Required
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:49
			// _ = "end of CoverTab[53502]"
		case s == "rep":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:50
			_go_fuzz_dep_.CoverTab[53503]++
															f.L1.Cardinality = protoreflect.Repeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:51
			// _ = "end of CoverTab[53503]"
		case s == "varint":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:52
			_go_fuzz_dep_.CoverTab[53504]++
															switch goType.Kind() {
			case reflect.Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:54
				_go_fuzz_dep_.CoverTab[53518]++
																f.L1.Kind = protoreflect.BoolKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:55
				// _ = "end of CoverTab[53518]"
			case reflect.Int32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:56
				_go_fuzz_dep_.CoverTab[53519]++
																f.L1.Kind = protoreflect.Int32Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:57
				// _ = "end of CoverTab[53519]"
			case reflect.Int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:58
				_go_fuzz_dep_.CoverTab[53520]++
																f.L1.Kind = protoreflect.Int64Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:59
				// _ = "end of CoverTab[53520]"
			case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:60
				_go_fuzz_dep_.CoverTab[53521]++
																f.L1.Kind = protoreflect.Uint32Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:61
				// _ = "end of CoverTab[53521]"
			case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:62
				_go_fuzz_dep_.CoverTab[53522]++
																f.L1.Kind = protoreflect.Uint64Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:63
				// _ = "end of CoverTab[53522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:63
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:63
				_go_fuzz_dep_.CoverTab[53523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:63
				// _ = "end of CoverTab[53523]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:64
			// _ = "end of CoverTab[53504]"
		case s == "zigzag32":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:65
			_go_fuzz_dep_.CoverTab[53505]++
															if goType.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:66
				_go_fuzz_dep_.CoverTab[53524]++
																f.L1.Kind = protoreflect.Sint32Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:67
				// _ = "end of CoverTab[53524]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:68
				_go_fuzz_dep_.CoverTab[53525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:68
				// _ = "end of CoverTab[53525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:68
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:68
			// _ = "end of CoverTab[53505]"
		case s == "zigzag64":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:69
			_go_fuzz_dep_.CoverTab[53506]++
															if goType.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:70
				_go_fuzz_dep_.CoverTab[53526]++
																f.L1.Kind = protoreflect.Sint64Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:71
				// _ = "end of CoverTab[53526]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:72
				_go_fuzz_dep_.CoverTab[53527]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:72
				// _ = "end of CoverTab[53527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:72
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:72
			// _ = "end of CoverTab[53506]"
		case s == "fixed32":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:73
			_go_fuzz_dep_.CoverTab[53507]++
															switch goType.Kind() {
			case reflect.Int32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:75
				_go_fuzz_dep_.CoverTab[53528]++
																f.L1.Kind = protoreflect.Sfixed32Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:76
				// _ = "end of CoverTab[53528]"
			case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:77
				_go_fuzz_dep_.CoverTab[53529]++
																f.L1.Kind = protoreflect.Fixed32Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:78
				// _ = "end of CoverTab[53529]"
			case reflect.Float32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:79
				_go_fuzz_dep_.CoverTab[53530]++
																f.L1.Kind = protoreflect.FloatKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:80
				// _ = "end of CoverTab[53530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:80
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:80
				_go_fuzz_dep_.CoverTab[53531]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:80
				// _ = "end of CoverTab[53531]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:81
			// _ = "end of CoverTab[53507]"
		case s == "fixed64":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:82
			_go_fuzz_dep_.CoverTab[53508]++
															switch goType.Kind() {
			case reflect.Int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:84
				_go_fuzz_dep_.CoverTab[53532]++
																f.L1.Kind = protoreflect.Sfixed64Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:85
				// _ = "end of CoverTab[53532]"
			case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:86
				_go_fuzz_dep_.CoverTab[53533]++
																f.L1.Kind = protoreflect.Fixed64Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:87
				// _ = "end of CoverTab[53533]"
			case reflect.Float64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:88
				_go_fuzz_dep_.CoverTab[53534]++
																f.L1.Kind = protoreflect.DoubleKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:89
				// _ = "end of CoverTab[53534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:89
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:89
				_go_fuzz_dep_.CoverTab[53535]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:89
				// _ = "end of CoverTab[53535]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:90
			// _ = "end of CoverTab[53508]"
		case s == "bytes":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:91
			_go_fuzz_dep_.CoverTab[53509]++
															switch {
			case goType.Kind() == reflect.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:93
				_go_fuzz_dep_.CoverTab[53536]++
																f.L1.Kind = protoreflect.StringKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:94
				// _ = "end of CoverTab[53536]"
			case goType.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:95
				_go_fuzz_dep_.CoverTab[53539]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:95
				return goType.Elem() == byteType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:95
				// _ = "end of CoverTab[53539]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:95
			}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:95
				_go_fuzz_dep_.CoverTab[53537]++
																f.L1.Kind = protoreflect.BytesKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:96
				// _ = "end of CoverTab[53537]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:97
				_go_fuzz_dep_.CoverTab[53538]++
																f.L1.Kind = protoreflect.MessageKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:98
				// _ = "end of CoverTab[53538]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:99
			// _ = "end of CoverTab[53509]"
		case s == "group":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:100
			_go_fuzz_dep_.CoverTab[53510]++
															f.L1.Kind = protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:101
			// _ = "end of CoverTab[53510]"
		case strings.HasPrefix(s, "enum="):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:102
			_go_fuzz_dep_.CoverTab[53511]++
															f.L1.Kind = protoreflect.EnumKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:103
			// _ = "end of CoverTab[53511]"
		case strings.HasPrefix(s, "json="):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:104
			_go_fuzz_dep_.CoverTab[53512]++
															jsonName := s[len("json="):]
															if jsonName != strs.JSONCamelCase(string(f.L0.FullName.Name())) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:106
				_go_fuzz_dep_.CoverTab[53540]++
																f.L1.StringName.InitJSON(jsonName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:107
				// _ = "end of CoverTab[53540]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:108
				_go_fuzz_dep_.CoverTab[53541]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:108
				// _ = "end of CoverTab[53541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:108
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:108
			// _ = "end of CoverTab[53512]"
		case s == "packed":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:109
			_go_fuzz_dep_.CoverTab[53513]++
															f.L1.HasPacked = true
															f.L1.IsPacked = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:111
			// _ = "end of CoverTab[53513]"
		case strings.HasPrefix(s, "weak="):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:112
			_go_fuzz_dep_.CoverTab[53514]++
															f.L1.IsWeak = true
															f.L1.Message = filedesc.PlaceholderMessage(protoreflect.FullName(s[len("weak="):]))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:114
			// _ = "end of CoverTab[53514]"
		case strings.HasPrefix(s, "def="):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:115
			_go_fuzz_dep_.CoverTab[53515]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:118
			s, i = tag[len("def="):], len(tag)
															v, ev, _ := defval.Unmarshal(s, f.L1.Kind, evs, defval.GoTag)
															f.L1.Default = filedesc.DefaultValue(v, ev)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:120
			// _ = "end of CoverTab[53515]"
		case s == "proto3":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:121
			_go_fuzz_dep_.CoverTab[53516]++
															f.L0.ParentFile = filedesc.SurrogateProto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:122
			// _ = "end of CoverTab[53516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:122
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:122
			_go_fuzz_dep_.CoverTab[53517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:122
			// _ = "end of CoverTab[53517]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:123
		// _ = "end of CoverTab[53495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:123
		_go_fuzz_dep_.CoverTab[53496]++
														tag = strings.TrimPrefix(tag[i:], ",")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:124
		// _ = "end of CoverTab[53496]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:125
	// _ = "end of CoverTab[53491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:125
	_go_fuzz_dep_.CoverTab[53492]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:129
	if f.L1.Kind == protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:129
		_go_fuzz_dep_.CoverTab[53542]++
														f.L0.FullName = protoreflect.FullName(strings.ToLower(string(f.L0.FullName)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:130
		// _ = "end of CoverTab[53542]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:131
		_go_fuzz_dep_.CoverTab[53543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:131
		// _ = "end of CoverTab[53543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:131
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:131
	// _ = "end of CoverTab[53492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:131
	_go_fuzz_dep_.CoverTab[53493]++
													return f
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:132
	// _ = "end of CoverTab[53493]"
}

// Marshal encodes the protoreflect.FieldDescriptor as a tag.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
// The enumName must be provided if the kind is an enum.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
// Historically, the formulation of the enum "name" was the proto package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
// dot-concatenated with the generated Go identifier for the enum type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
// Depending on the context on how Marshal is called, there are different ways
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
// through which that information is determined. As such it is the caller's
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:135
// responsibility to provide a function to obtain that information.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:143
func Marshal(fd protoreflect.FieldDescriptor, enumName string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:143
	_go_fuzz_dep_.CoverTab[53544]++
													var tag []string
													switch fd.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind, protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Int64Kind, protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:146
		_go_fuzz_dep_.CoverTab[53555]++
														tag = append(tag, "varint")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:147
		// _ = "end of CoverTab[53555]"
	case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:148
		_go_fuzz_dep_.CoverTab[53556]++
														tag = append(tag, "zigzag32")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:149
		// _ = "end of CoverTab[53556]"
	case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:150
		_go_fuzz_dep_.CoverTab[53557]++
														tag = append(tag, "zigzag64")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:151
		// _ = "end of CoverTab[53557]"
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:152
		_go_fuzz_dep_.CoverTab[53558]++
														tag = append(tag, "fixed32")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:153
		// _ = "end of CoverTab[53558]"
	case protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:154
		_go_fuzz_dep_.CoverTab[53559]++
														tag = append(tag, "fixed64")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:155
		// _ = "end of CoverTab[53559]"
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:156
		_go_fuzz_dep_.CoverTab[53560]++
														tag = append(tag, "bytes")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:157
		// _ = "end of CoverTab[53560]"
	case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:158
		_go_fuzz_dep_.CoverTab[53561]++
														tag = append(tag, "group")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:159
		// _ = "end of CoverTab[53561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:159
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:159
		_go_fuzz_dep_.CoverTab[53562]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:159
		// _ = "end of CoverTab[53562]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:160
	// _ = "end of CoverTab[53544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:160
	_go_fuzz_dep_.CoverTab[53545]++
													tag = append(tag, strconv.Itoa(int(fd.Number())))
													switch fd.Cardinality() {
	case protoreflect.Optional:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:163
		_go_fuzz_dep_.CoverTab[53563]++
														tag = append(tag, "opt")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:164
		// _ = "end of CoverTab[53563]"
	case protoreflect.Required:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:165
		_go_fuzz_dep_.CoverTab[53564]++
														tag = append(tag, "req")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:166
		// _ = "end of CoverTab[53564]"
	case protoreflect.Repeated:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:167
		_go_fuzz_dep_.CoverTab[53565]++
														tag = append(tag, "rep")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:168
		// _ = "end of CoverTab[53565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:168
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:168
		_go_fuzz_dep_.CoverTab[53566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:168
		// _ = "end of CoverTab[53566]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:169
	// _ = "end of CoverTab[53545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:169
	_go_fuzz_dep_.CoverTab[53546]++
													if fd.IsPacked() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:170
		_go_fuzz_dep_.CoverTab[53567]++
														tag = append(tag, "packed")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:171
		// _ = "end of CoverTab[53567]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:172
		_go_fuzz_dep_.CoverTab[53568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:172
		// _ = "end of CoverTab[53568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:172
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:172
	// _ = "end of CoverTab[53546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:172
	_go_fuzz_dep_.CoverTab[53547]++
													name := string(fd.Name())
													if fd.Kind() == protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:174
		_go_fuzz_dep_.CoverTab[53569]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:178
		name = string(fd.Message().Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:178
		// _ = "end of CoverTab[53569]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:179
		_go_fuzz_dep_.CoverTab[53570]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:179
		// _ = "end of CoverTab[53570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:179
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:179
	// _ = "end of CoverTab[53547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:179
	_go_fuzz_dep_.CoverTab[53548]++
													tag = append(tag, "name="+name)
													if jsonName := fd.JSONName(); jsonName != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		_go_fuzz_dep_.CoverTab[53571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		return jsonName != name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		// _ = "end of CoverTab[53571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		_go_fuzz_dep_.CoverTab[53572]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		return !fd.IsExtension()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		// _ = "end of CoverTab[53572]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:181
		_go_fuzz_dep_.CoverTab[53573]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:184
		tag = append(tag, "json="+jsonName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:184
		// _ = "end of CoverTab[53573]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:185
		_go_fuzz_dep_.CoverTab[53574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:185
		// _ = "end of CoverTab[53574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:185
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:185
	// _ = "end of CoverTab[53548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:185
	_go_fuzz_dep_.CoverTab[53549]++
													if fd.IsWeak() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:186
		_go_fuzz_dep_.CoverTab[53575]++
														tag = append(tag, "weak="+string(fd.Message().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:187
		// _ = "end of CoverTab[53575]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:188
		_go_fuzz_dep_.CoverTab[53576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:188
		// _ = "end of CoverTab[53576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:188
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:188
	// _ = "end of CoverTab[53549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:188
	_go_fuzz_dep_.CoverTab[53550]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:192
	if fd.Syntax() == protoreflect.Proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:192
		_go_fuzz_dep_.CoverTab[53577]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:192
		return !fd.IsExtension()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:192
		// _ = "end of CoverTab[53577]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:192
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:192
		_go_fuzz_dep_.CoverTab[53578]++
														tag = append(tag, "proto3")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:193
		// _ = "end of CoverTab[53578]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:194
		_go_fuzz_dep_.CoverTab[53579]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:194
		// _ = "end of CoverTab[53579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:194
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:194
	// _ = "end of CoverTab[53550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:194
	_go_fuzz_dep_.CoverTab[53551]++
													if fd.Kind() == protoreflect.EnumKind && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:195
		_go_fuzz_dep_.CoverTab[53580]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:195
		return enumName != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:195
		// _ = "end of CoverTab[53580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:195
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:195
		_go_fuzz_dep_.CoverTab[53581]++
														tag = append(tag, "enum="+enumName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:196
		// _ = "end of CoverTab[53581]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:197
		_go_fuzz_dep_.CoverTab[53582]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:197
		// _ = "end of CoverTab[53582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:197
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:197
	// _ = "end of CoverTab[53551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:197
	_go_fuzz_dep_.CoverTab[53552]++
													if fd.ContainingOneof() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:198
		_go_fuzz_dep_.CoverTab[53583]++
														tag = append(tag, "oneof")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:199
		// _ = "end of CoverTab[53583]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:200
		_go_fuzz_dep_.CoverTab[53584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:200
		// _ = "end of CoverTab[53584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:200
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:200
	// _ = "end of CoverTab[53552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:200
	_go_fuzz_dep_.CoverTab[53553]++

													if fd.HasDefault() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:202
		_go_fuzz_dep_.CoverTab[53585]++
														def, _ := defval.Marshal(fd.Default(), fd.DefaultEnumValue(), fd.Kind(), defval.GoTag)
														tag = append(tag, "def="+def)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:204
		// _ = "end of CoverTab[53585]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:205
		_go_fuzz_dep_.CoverTab[53586]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:205
		// _ = "end of CoverTab[53586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:205
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:205
	// _ = "end of CoverTab[53553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:205
	_go_fuzz_dep_.CoverTab[53554]++
													return strings.Join(tag, ",")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:206
	// _ = "end of CoverTab[53554]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:207
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/tag/tag.go:207
var _ = _go_fuzz_dep_.CoverTab
