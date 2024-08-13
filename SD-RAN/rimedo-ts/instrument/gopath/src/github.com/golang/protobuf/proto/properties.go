// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:5
)

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// StructProperties represents protocol buffer type information for a
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:18
// generated protobuf message in the open-struct API.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:18
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:18
// Deprecated: Do not use.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:22
type StructProperties struct {
	// Prop are the properties for each field.
	//
	// Fields belonging to a oneof are stored in OneofTypes instead, with a
	// single Properties representing the parent oneof held here.
	//
	// The order of Prop matches the order of fields in the Go struct.
	// Struct fields that are not related to protobufs have a "XXX_" prefix
	// in the Properties.Name and must be ignored by the user.
	Prop	[]*Properties

	// OneofTypes contains information about the oneof fields in this message.
	// It is keyed by the protobuf field name.
	OneofTypes	map[string]*OneofProperties
}

// Properties represents the type information for a protobuf message field.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:38
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:38
// Deprecated: Do not use.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:41
type Properties struct {
	// Name is a placeholder name with little meaningful semantic value.
	// If the name has an "XXX_" prefix, the entire Properties must be ignored.
	Name	string
	// OrigName is the protobuf field name or oneof name.
	OrigName	string
	// JSONName is the JSON name for the protobuf field.
	JSONName	string
	// Enum is a placeholder name for enums.
	// For historical reasons, this is neither the Go name for the enum,
	// nor the protobuf name for the enum.
	Enum	string	// Deprecated: Do not use.
	// Weak contains the full name of the weakly referenced message.
	Weak	string
	// Wire is a string representation of the wire type.
	Wire	string
	// WireType is the protobuf wire type for the field.
	WireType	int
	// Tag is the protobuf field number.
	Tag	int
	// Required reports whether this is a required field.
	Required	bool
	// Optional reports whether this is a optional field.
	Optional	bool
	// Repeated reports whether this is a repeated field.
	Repeated	bool
	// Packed reports whether this is a packed repeated field of scalars.
	Packed	bool
	// Proto3 reports whether this field operates under the proto3 syntax.
	Proto3	bool
	// Oneof reports whether this field belongs within a oneof.
	Oneof	bool

	// Default is the default value in string form.
	Default	string
	// HasDefault reports whether the field has a default value.
	HasDefault	bool

	// MapKeyProp is the properties for the key field for a map field.
	MapKeyProp	*Properties
	// MapValProp is the properties for the value field for a map field.
	MapValProp	*Properties
}

// OneofProperties represents the type information for a protobuf oneof.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:85
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:85
// Deprecated: Do not use.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:88
type OneofProperties struct {
	// Type is a pointer to the generated wrapper type for the field value.
	// This is nil for messages that are not in the open-struct API.
	Type	reflect.Type
	// Field is the index into StructProperties.Prop for the containing oneof.
	Field	int
	// Prop is the properties for the field.
	Prop	*Properties
}

// String formats the properties in the protobuf struct field tag style.
func (p *Properties) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:99
	_go_fuzz_dep_.CoverTab[61429]++
												s := p.Wire
												s += "," + strconv.Itoa(p.Tag)
												if p.Required {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:102
		_go_fuzz_dep_.CoverTab[61440]++
													s += ",req"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:103
		// _ = "end of CoverTab[61440]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:104
		_go_fuzz_dep_.CoverTab[61441]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:104
		// _ = "end of CoverTab[61441]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:104
	// _ = "end of CoverTab[61429]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:104
	_go_fuzz_dep_.CoverTab[61430]++
												if p.Optional {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:105
		_go_fuzz_dep_.CoverTab[61442]++
													s += ",opt"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:106
		// _ = "end of CoverTab[61442]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:107
		_go_fuzz_dep_.CoverTab[61443]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:107
		// _ = "end of CoverTab[61443]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:107
	// _ = "end of CoverTab[61430]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:107
	_go_fuzz_dep_.CoverTab[61431]++
												if p.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:108
		_go_fuzz_dep_.CoverTab[61444]++
													s += ",rep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:109
		// _ = "end of CoverTab[61444]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:110
		_go_fuzz_dep_.CoverTab[61445]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:110
		// _ = "end of CoverTab[61445]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:110
	// _ = "end of CoverTab[61431]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:110
	_go_fuzz_dep_.CoverTab[61432]++
												if p.Packed {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:111
		_go_fuzz_dep_.CoverTab[61446]++
													s += ",packed"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:112
		// _ = "end of CoverTab[61446]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:113
		_go_fuzz_dep_.CoverTab[61447]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:113
		// _ = "end of CoverTab[61447]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:113
	// _ = "end of CoverTab[61432]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:113
	_go_fuzz_dep_.CoverTab[61433]++
												s += ",name=" + p.OrigName
												if p.JSONName != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:115
		_go_fuzz_dep_.CoverTab[61448]++
													s += ",json=" + p.JSONName
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:116
		// _ = "end of CoverTab[61448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:117
		_go_fuzz_dep_.CoverTab[61449]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:117
		// _ = "end of CoverTab[61449]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:117
	// _ = "end of CoverTab[61433]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:117
	_go_fuzz_dep_.CoverTab[61434]++
												if len(p.Enum) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:118
		_go_fuzz_dep_.CoverTab[61450]++
													s += ",enum=" + p.Enum
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:119
		// _ = "end of CoverTab[61450]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:120
		_go_fuzz_dep_.CoverTab[61451]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:120
		// _ = "end of CoverTab[61451]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:120
	// _ = "end of CoverTab[61434]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:120
	_go_fuzz_dep_.CoverTab[61435]++
												if len(p.Weak) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:121
		_go_fuzz_dep_.CoverTab[61452]++
													s += ",weak=" + p.Weak
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:122
		// _ = "end of CoverTab[61452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:123
		_go_fuzz_dep_.CoverTab[61453]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:123
		// _ = "end of CoverTab[61453]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:123
	// _ = "end of CoverTab[61435]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:123
	_go_fuzz_dep_.CoverTab[61436]++
												if p.Proto3 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:124
		_go_fuzz_dep_.CoverTab[61454]++
													s += ",proto3"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:125
		// _ = "end of CoverTab[61454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:126
		_go_fuzz_dep_.CoverTab[61455]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:126
		// _ = "end of CoverTab[61455]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:126
	// _ = "end of CoverTab[61436]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:126
	_go_fuzz_dep_.CoverTab[61437]++
												if p.Oneof {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:127
		_go_fuzz_dep_.CoverTab[61456]++
													s += ",oneof"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:128
		// _ = "end of CoverTab[61456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:129
		_go_fuzz_dep_.CoverTab[61457]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:129
		// _ = "end of CoverTab[61457]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:129
	// _ = "end of CoverTab[61437]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:129
	_go_fuzz_dep_.CoverTab[61438]++
												if p.HasDefault {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:130
		_go_fuzz_dep_.CoverTab[61458]++
													s += ",def=" + p.Default
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:131
		// _ = "end of CoverTab[61458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:132
		_go_fuzz_dep_.CoverTab[61459]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:132
		// _ = "end of CoverTab[61459]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:132
	// _ = "end of CoverTab[61438]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:132
	_go_fuzz_dep_.CoverTab[61439]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:133
	// _ = "end of CoverTab[61439]"
}

// Parse populates p by parsing a string in the protobuf struct field tag style.
func (p *Properties) Parse(tag string) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:137
	_go_fuzz_dep_.CoverTab[61460]++

												for len(tag) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:139
		_go_fuzz_dep_.CoverTab[61461]++
													i := strings.IndexByte(tag, ',')
													if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:141
			_go_fuzz_dep_.CoverTab[61464]++
														i = len(tag)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:142
			// _ = "end of CoverTab[61464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:143
			_go_fuzz_dep_.CoverTab[61465]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:143
			// _ = "end of CoverTab[61465]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:143
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:143
		// _ = "end of CoverTab[61461]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:143
		_go_fuzz_dep_.CoverTab[61462]++
													switch s := tag[:i]; {
		case strings.HasPrefix(s, "name="):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:145
			_go_fuzz_dep_.CoverTab[61466]++
														p.OrigName = s[len("name="):]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:146
			// _ = "end of CoverTab[61466]"
		case strings.HasPrefix(s, "json="):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:147
			_go_fuzz_dep_.CoverTab[61467]++
														p.JSONName = s[len("json="):]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:148
			// _ = "end of CoverTab[61467]"
		case strings.HasPrefix(s, "enum="):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:149
			_go_fuzz_dep_.CoverTab[61468]++
														p.Enum = s[len("enum="):]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:150
			// _ = "end of CoverTab[61468]"
		case strings.HasPrefix(s, "weak="):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:151
			_go_fuzz_dep_.CoverTab[61469]++
														p.Weak = s[len("weak="):]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:152
			// _ = "end of CoverTab[61469]"
		case strings.Trim(s, "0123456789") == "":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:153
			_go_fuzz_dep_.CoverTab[61470]++
														n, _ := strconv.ParseUint(s, 10, 32)
														p.Tag = int(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:155
			// _ = "end of CoverTab[61470]"
		case s == "opt":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:156
			_go_fuzz_dep_.CoverTab[61471]++
														p.Optional = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:157
			// _ = "end of CoverTab[61471]"
		case s == "req":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:158
			_go_fuzz_dep_.CoverTab[61472]++
														p.Required = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:159
			// _ = "end of CoverTab[61472]"
		case s == "rep":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:160
			_go_fuzz_dep_.CoverTab[61473]++
														p.Repeated = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:161
			// _ = "end of CoverTab[61473]"
		case s == "varint" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			_go_fuzz_dep_.CoverTab[61484]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			return s == "zigzag32"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			// _ = "end of CoverTab[61484]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			_go_fuzz_dep_.CoverTab[61485]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			return s == "zigzag64"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			// _ = "end of CoverTab[61485]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
		}():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:162
			_go_fuzz_dep_.CoverTab[61474]++
														p.Wire = s
														p.WireType = WireVarint
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:164
			// _ = "end of CoverTab[61474]"
		case s == "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:165
			_go_fuzz_dep_.CoverTab[61475]++
														p.Wire = s
														p.WireType = WireFixed32
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:167
			// _ = "end of CoverTab[61475]"
		case s == "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:168
			_go_fuzz_dep_.CoverTab[61476]++
														p.Wire = s
														p.WireType = WireFixed64
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:170
			// _ = "end of CoverTab[61476]"
		case s == "bytes":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:171
			_go_fuzz_dep_.CoverTab[61477]++
														p.Wire = s
														p.WireType = WireBytes
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:173
			// _ = "end of CoverTab[61477]"
		case s == "group":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:174
			_go_fuzz_dep_.CoverTab[61478]++
														p.Wire = s
														p.WireType = WireStartGroup
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:176
			// _ = "end of CoverTab[61478]"
		case s == "packed":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:177
			_go_fuzz_dep_.CoverTab[61479]++
														p.Packed = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:178
			// _ = "end of CoverTab[61479]"
		case s == "proto3":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:179
			_go_fuzz_dep_.CoverTab[61480]++
														p.Proto3 = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:180
			// _ = "end of CoverTab[61480]"
		case s == "oneof":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:181
			_go_fuzz_dep_.CoverTab[61481]++
														p.Oneof = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:182
			// _ = "end of CoverTab[61481]"
		case strings.HasPrefix(s, "def="):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:183
			_go_fuzz_dep_.CoverTab[61482]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:186
			p.HasDefault = true
														p.Default, i = tag[len("def="):], len(tag)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:187
			// _ = "end of CoverTab[61482]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:187
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:187
			_go_fuzz_dep_.CoverTab[61483]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:187
			// _ = "end of CoverTab[61483]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:188
		// _ = "end of CoverTab[61462]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:188
		_go_fuzz_dep_.CoverTab[61463]++
													tag = strings.TrimPrefix(tag[i:], ",")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:189
		// _ = "end of CoverTab[61463]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:190
	// _ = "end of CoverTab[61460]"
}

// Init populates the properties from a protocol buffer struct tag.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:193
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:193
// Deprecated: Do not use.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:196
func (p *Properties) Init(typ reflect.Type, name, tag string, f *reflect.StructField) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:196
	_go_fuzz_dep_.CoverTab[61486]++
												p.Name = name
												p.OrigName = name
												if tag == "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:199
		_go_fuzz_dep_.CoverTab[61488]++
													return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:200
		// _ = "end of CoverTab[61488]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:201
		_go_fuzz_dep_.CoverTab[61489]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:201
		// _ = "end of CoverTab[61489]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:201
	// _ = "end of CoverTab[61486]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:201
	_go_fuzz_dep_.CoverTab[61487]++
												p.Parse(tag)

												if typ != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:204
		_go_fuzz_dep_.CoverTab[61490]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:204
		return typ.Kind() == reflect.Map
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:204
		// _ = "end of CoverTab[61490]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:204
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:204
		_go_fuzz_dep_.CoverTab[61491]++
													p.MapKeyProp = new(Properties)
													p.MapKeyProp.Init(nil, "Key", f.Tag.Get("protobuf_key"), nil)
													p.MapValProp = new(Properties)
													p.MapValProp.Init(nil, "Value", f.Tag.Get("protobuf_val"), nil)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:208
		// _ = "end of CoverTab[61491]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:209
		_go_fuzz_dep_.CoverTab[61492]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:209
		// _ = "end of CoverTab[61492]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:209
	// _ = "end of CoverTab[61487]"
}

var propertiesCache sync.Map	// map[reflect.Type]*StructProperties

// GetProperties returns the list of properties for the type represented by t,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:214
// which must be a generated protocol buffer message in the open-struct API,
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:214
// where protobuf message fields are represented by exported Go struct fields.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:214
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:214
// Deprecated: Use protobuf reflection instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:219
func GetProperties(t reflect.Type) *StructProperties {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:219
	_go_fuzz_dep_.CoverTab[61493]++
												if p, ok := propertiesCache.Load(t); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:220
		_go_fuzz_dep_.CoverTab[61495]++
													return p.(*StructProperties)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:221
		// _ = "end of CoverTab[61495]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:222
		_go_fuzz_dep_.CoverTab[61496]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:222
		// _ = "end of CoverTab[61496]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:222
	// _ = "end of CoverTab[61493]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:222
	_go_fuzz_dep_.CoverTab[61494]++
												p, _ := propertiesCache.LoadOrStore(t, newProperties(t))
												return p.(*StructProperties)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:224
	// _ = "end of CoverTab[61494]"
}

func newProperties(t reflect.Type) *StructProperties {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:227
	_go_fuzz_dep_.CoverTab[61497]++
												if t.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:228
		_go_fuzz_dep_.CoverTab[61501]++
													panic(fmt.Sprintf("%v is not a generated message in the open-struct API", t))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:229
		// _ = "end of CoverTab[61501]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:230
		_go_fuzz_dep_.CoverTab[61502]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:230
		// _ = "end of CoverTab[61502]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:230
	// _ = "end of CoverTab[61497]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:230
	_go_fuzz_dep_.CoverTab[61498]++

												var hasOneof bool
												prop := new(StructProperties)

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:236
	for i := 0; i < t.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:236
		_go_fuzz_dep_.CoverTab[61503]++
													p := new(Properties)
													f := t.Field(i)
													tagField := f.Tag.Get("protobuf")
													p.Init(f.Type, f.Name, tagField, &f)

													tagOneof := f.Tag.Get("protobuf_oneof")
													if tagOneof != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:243
			_go_fuzz_dep_.CoverTab[61506]++
														hasOneof = true
														p.OrigName = tagOneof
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:245
			// _ = "end of CoverTab[61506]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:246
			_go_fuzz_dep_.CoverTab[61507]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:246
			// _ = "end of CoverTab[61507]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:246
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:246
		// _ = "end of CoverTab[61503]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:246
		_go_fuzz_dep_.CoverTab[61504]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
		if tagField == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			_go_fuzz_dep_.CoverTab[61508]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			return tagOneof == ""
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			// _ = "end of CoverTab[61508]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			_go_fuzz_dep_.CoverTab[61509]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			return !strings.HasPrefix(p.Name, "XXX_")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			// _ = "end of CoverTab[61509]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:250
			_go_fuzz_dep_.CoverTab[61510]++
														p.Name = "XXX_" + p.Name
														p.OrigName = "XXX_" + p.OrigName
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:252
			// _ = "end of CoverTab[61510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:253
			_go_fuzz_dep_.CoverTab[61511]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:253
			if p.Weak != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:253
				_go_fuzz_dep_.CoverTab[61512]++
															p.Name = p.OrigName
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:254
				// _ = "end of CoverTab[61512]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
				_go_fuzz_dep_.CoverTab[61513]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
				// _ = "end of CoverTab[61513]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
			// _ = "end of CoverTab[61511]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
		// _ = "end of CoverTab[61504]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:255
		_go_fuzz_dep_.CoverTab[61505]++

													prop.Prop = append(prop.Prop, p)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:257
		// _ = "end of CoverTab[61505]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:258
	// _ = "end of CoverTab[61498]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:258
	_go_fuzz_dep_.CoverTab[61499]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:261
	if hasOneof {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:261
		_go_fuzz_dep_.CoverTab[61514]++
													var oneofWrappers []interface{}
													if fn, ok := reflect.PtrTo(t).MethodByName("XXX_OneofFuncs"); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:263
			_go_fuzz_dep_.CoverTab[61518]++
														oneofWrappers = fn.Func.Call([]reflect.Value{reflect.Zero(fn.Type.In(0))})[3].Interface().([]interface{})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:264
			// _ = "end of CoverTab[61518]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:265
			_go_fuzz_dep_.CoverTab[61519]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:265
			// _ = "end of CoverTab[61519]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:265
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:265
		// _ = "end of CoverTab[61514]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:265
		_go_fuzz_dep_.CoverTab[61515]++
													if fn, ok := reflect.PtrTo(t).MethodByName("XXX_OneofWrappers"); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:266
			_go_fuzz_dep_.CoverTab[61520]++
														oneofWrappers = fn.Func.Call([]reflect.Value{reflect.Zero(fn.Type.In(0))})[0].Interface().([]interface{})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:267
			// _ = "end of CoverTab[61520]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:268
			_go_fuzz_dep_.CoverTab[61521]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:268
			// _ = "end of CoverTab[61521]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:268
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:268
		// _ = "end of CoverTab[61515]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:268
		_go_fuzz_dep_.CoverTab[61516]++
													if m, ok := reflect.Zero(reflect.PtrTo(t)).Interface().(protoreflect.ProtoMessage); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:269
			_go_fuzz_dep_.CoverTab[61522]++
														if m, ok := m.ProtoReflect().(interface{ ProtoMessageInfo() *protoimpl.MessageInfo }); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:270
				_go_fuzz_dep_.CoverTab[61523]++
															oneofWrappers = m.ProtoMessageInfo().OneofWrappers
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:271
				// _ = "end of CoverTab[61523]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:272
				_go_fuzz_dep_.CoverTab[61524]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:272
				// _ = "end of CoverTab[61524]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:272
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:272
			// _ = "end of CoverTab[61522]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:273
			_go_fuzz_dep_.CoverTab[61525]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:273
			// _ = "end of CoverTab[61525]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:273
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:273
		// _ = "end of CoverTab[61516]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:273
		_go_fuzz_dep_.CoverTab[61517]++

													prop.OneofTypes = make(map[string]*OneofProperties)
													for _, wrapper := range oneofWrappers {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:276
			_go_fuzz_dep_.CoverTab[61526]++
														p := &OneofProperties{
				Type:	reflect.ValueOf(wrapper).Type(),
				Prop:	new(Properties),
			}
			f := p.Type.Elem().Field(0)
			p.Prop.Name = f.Name
			p.Prop.Parse(f.Tag.Get("protobuf"))

			// Determine the struct field that contains this oneof.
			// Each wrapper is assignable to exactly one parent field.
			var foundOneof bool
			for i := 0; i < t.NumField() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:288
				_go_fuzz_dep_.CoverTab[61529]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:288
				return !foundOneof
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:288
				// _ = "end of CoverTab[61529]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:288
			}(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:288
				_go_fuzz_dep_.CoverTab[61530]++
															if p.Type.AssignableTo(t.Field(i).Type) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:289
					_go_fuzz_dep_.CoverTab[61531]++
																p.Field = i
																foundOneof = true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:291
					// _ = "end of CoverTab[61531]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:292
					_go_fuzz_dep_.CoverTab[61532]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:292
					// _ = "end of CoverTab[61532]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:292
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:292
				// _ = "end of CoverTab[61530]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:293
			// _ = "end of CoverTab[61526]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:293
			_go_fuzz_dep_.CoverTab[61527]++
														if !foundOneof {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:294
				_go_fuzz_dep_.CoverTab[61533]++
															panic(fmt.Sprintf("%v is not a generated message in the open-struct API", t))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:295
				// _ = "end of CoverTab[61533]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:296
				_go_fuzz_dep_.CoverTab[61534]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:296
				// _ = "end of CoverTab[61534]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:296
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:296
			// _ = "end of CoverTab[61527]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:296
			_go_fuzz_dep_.CoverTab[61528]++
														prop.OneofTypes[p.Prop.OrigName] = p
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:297
			// _ = "end of CoverTab[61528]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:298
		// _ = "end of CoverTab[61517]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:299
		_go_fuzz_dep_.CoverTab[61535]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:299
		// _ = "end of CoverTab[61535]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:299
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:299
	// _ = "end of CoverTab[61499]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:299
	_go_fuzz_dep_.CoverTab[61500]++

												return prop
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:301
	// _ = "end of CoverTab[61500]"
}

func (sp *StructProperties) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:304
	_go_fuzz_dep_.CoverTab[61536]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:304
	return len(sp.Prop)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:304
	// _ = "end of CoverTab[61536]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:304
}
func (sp *StructProperties) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:305
	_go_fuzz_dep_.CoverTab[61537]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:305
	return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:305
	// _ = "end of CoverTab[61537]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:305
}
func (sp *StructProperties) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:306
	_go_fuzz_dep_.CoverTab[61538]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:306
	return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:306
	// _ = "end of CoverTab[61538]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:306
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:306
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/properties.go:306
var _ = _go_fuzz_dep_.CoverTab
