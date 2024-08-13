// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
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
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:37
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:43
import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const debug bool = false

// Constants that identify the encoding of a value on the wire.
const (
	WireVarint	= 0
	WireFixed64	= 1
	WireBytes	= 2
	WireStartGroup	= 3
	WireEndGroup	= 4
	WireFixed32	= 5
)

// tagMap is an optimization over map[int]int for typical protocol buffer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:65
// use-cases. Encoded protocol buffers are often in tag order with small tag
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:65
// numbers.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:68
type tagMap struct {
	fastTags	[]int
	slowTags	map[int]int
}

// tagMapFastLimit is the upper bound on the tag number that will be stored in
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:73
// the tagMap slice rather than its map.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:75
const tagMapFastLimit = 1024

func (p *tagMap) get(t int) (int, bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:77
	_go_fuzz_dep_.CoverTab[108910]++
												if t > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:78
		_go_fuzz_dep_.CoverTab[108912]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:78
		return t < tagMapFastLimit
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:78
		// _ = "end of CoverTab[108912]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:78
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:78
		_go_fuzz_dep_.CoverTab[108913]++
													if t >= len(p.fastTags) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:79
			_go_fuzz_dep_.CoverTab[108915]++
														return 0, false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:80
			// _ = "end of CoverTab[108915]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:81
			_go_fuzz_dep_.CoverTab[108916]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:81
			// _ = "end of CoverTab[108916]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:81
		// _ = "end of CoverTab[108913]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:81
		_go_fuzz_dep_.CoverTab[108914]++
													fi := p.fastTags[t]
													return fi, fi >= 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:83
		// _ = "end of CoverTab[108914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:84
		_go_fuzz_dep_.CoverTab[108917]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:84
		// _ = "end of CoverTab[108917]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:84
	// _ = "end of CoverTab[108910]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:84
	_go_fuzz_dep_.CoverTab[108911]++
												fi, ok := p.slowTags[t]
												return fi, ok
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:86
	// _ = "end of CoverTab[108911]"
}

func (p *tagMap) put(t int, fi int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:89
	_go_fuzz_dep_.CoverTab[108918]++
												if t > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:90
		_go_fuzz_dep_.CoverTab[108921]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:90
		return t < tagMapFastLimit
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:90
		// _ = "end of CoverTab[108921]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:90
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:90
		_go_fuzz_dep_.CoverTab[108922]++
													for len(p.fastTags) < t+1 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:91
			_go_fuzz_dep_.CoverTab[108924]++
														p.fastTags = append(p.fastTags, -1)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:92
			// _ = "end of CoverTab[108924]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:93
		// _ = "end of CoverTab[108922]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:93
		_go_fuzz_dep_.CoverTab[108923]++
													p.fastTags[t] = fi
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:95
		// _ = "end of CoverTab[108923]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:96
		_go_fuzz_dep_.CoverTab[108925]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:96
		// _ = "end of CoverTab[108925]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:96
	// _ = "end of CoverTab[108918]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:96
	_go_fuzz_dep_.CoverTab[108919]++
												if p.slowTags == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:97
		_go_fuzz_dep_.CoverTab[108926]++
													p.slowTags = make(map[int]int)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:98
		// _ = "end of CoverTab[108926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:99
		_go_fuzz_dep_.CoverTab[108927]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:99
		// _ = "end of CoverTab[108927]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:99
	// _ = "end of CoverTab[108919]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:99
	_go_fuzz_dep_.CoverTab[108920]++
												p.slowTags[t] = fi
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:100
	// _ = "end of CoverTab[108920]"
}

// StructProperties represents properties for all the fields of a struct.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:103
// decoderTags and decoderOrigNames should only be used by the decoder.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:105
type StructProperties struct {
	Prop			[]*Properties	// properties for each field
	reqCount		int		// required count
	decoderTags		tagMap		// map from proto tag to struct field number
	decoderOrigNames	map[string]int	// map from original name to struct field number
	order			[]int		// list of struct field numbers in tag order

	// OneofTypes contains information about the oneof fields in this message.
	// It is keyed by the original name of a field.
	OneofTypes	map[string]*OneofProperties
}

// OneofProperties represents information about a specific field in a oneof.
type OneofProperties struct {
	Type	reflect.Type	// pointer to generated struct type for this oneof field
	Field	int		// struct field number of the containing oneof in the message
	Prop	*Properties
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:127
func (sp *StructProperties) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:127
	_go_fuzz_dep_.CoverTab[108928]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:127
	return len(sp.order)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:127
	// _ = "end of CoverTab[108928]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:127
}
func (sp *StructProperties) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:128
	_go_fuzz_dep_.CoverTab[108929]++
												return sp.Prop[sp.order[i]].Tag < sp.Prop[sp.order[j]].Tag
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:129
	// _ = "end of CoverTab[108929]"
}
func (sp *StructProperties) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:131
	_go_fuzz_dep_.CoverTab[108930]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:131
	sp.order[i], sp.order[j] = sp.order[j], sp.order[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:131
	// _ = "end of CoverTab[108930]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:131
}

// Properties represents the protocol-specific behavior of a single struct field.
type Properties struct {
	Name		string	// name of the field, for error messages
	OrigName	string	// original name before protocol compiler (always set)
	JSONName	string	// name to use for JSON; determined by protoc
	Wire		string
	WireType	int
	Tag		int
	Required	bool
	Optional	bool
	Repeated	bool
	Packed		bool	// relevant for repeated primitives only
	Enum		string	// set for enum types only
	proto3		bool	// whether this is known to be a proto3 field
	oneof		bool	// whether this is a oneof field

	Default		string	// default value
	HasDefault	bool	// whether an explicit default was provided
	CustomType	string
	CastType	string
	StdTime		bool
	StdDuration	bool
	WktPointer	bool

	stype	reflect.Type		// set for struct types only
	ctype	reflect.Type		// set for custom types only
	sprop	*StructProperties	// set for struct types only

	mtype		reflect.Type	// set for map types only
	MapKeyProp	*Properties	// set for map types only
	MapValProp	*Properties	// set for map types only
}

// String formats the properties in the protobuf struct field tag style.
func (p *Properties) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:167
	_go_fuzz_dep_.CoverTab[108931]++
												s := p.Wire
												s += ","
												s += strconv.Itoa(p.Tag)
												if p.Required {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:171
		_go_fuzz_dep_.CoverTab[108941]++
													s += ",req"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:172
		// _ = "end of CoverTab[108941]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:173
		_go_fuzz_dep_.CoverTab[108942]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:173
		// _ = "end of CoverTab[108942]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:173
	// _ = "end of CoverTab[108931]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:173
	_go_fuzz_dep_.CoverTab[108932]++
												if p.Optional {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:174
		_go_fuzz_dep_.CoverTab[108943]++
													s += ",opt"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:175
		// _ = "end of CoverTab[108943]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:176
		_go_fuzz_dep_.CoverTab[108944]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:176
		// _ = "end of CoverTab[108944]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:176
	// _ = "end of CoverTab[108932]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:176
	_go_fuzz_dep_.CoverTab[108933]++
												if p.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:177
		_go_fuzz_dep_.CoverTab[108945]++
													s += ",rep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:178
		// _ = "end of CoverTab[108945]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:179
		_go_fuzz_dep_.CoverTab[108946]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:179
		// _ = "end of CoverTab[108946]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:179
	// _ = "end of CoverTab[108933]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:179
	_go_fuzz_dep_.CoverTab[108934]++
												if p.Packed {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:180
		_go_fuzz_dep_.CoverTab[108947]++
													s += ",packed"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:181
		// _ = "end of CoverTab[108947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:182
		_go_fuzz_dep_.CoverTab[108948]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:182
		// _ = "end of CoverTab[108948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:182
	// _ = "end of CoverTab[108934]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:182
	_go_fuzz_dep_.CoverTab[108935]++
												s += ",name=" + p.OrigName
												if p.JSONName != p.OrigName {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:184
		_go_fuzz_dep_.CoverTab[108949]++
													s += ",json=" + p.JSONName
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:185
		// _ = "end of CoverTab[108949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:186
		_go_fuzz_dep_.CoverTab[108950]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:186
		// _ = "end of CoverTab[108950]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:186
	// _ = "end of CoverTab[108935]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:186
	_go_fuzz_dep_.CoverTab[108936]++
												if p.proto3 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:187
		_go_fuzz_dep_.CoverTab[108951]++
													s += ",proto3"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:188
		// _ = "end of CoverTab[108951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:189
		_go_fuzz_dep_.CoverTab[108952]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:189
		// _ = "end of CoverTab[108952]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:189
	// _ = "end of CoverTab[108936]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:189
	_go_fuzz_dep_.CoverTab[108937]++
												if p.oneof {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:190
		_go_fuzz_dep_.CoverTab[108953]++
													s += ",oneof"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:191
		// _ = "end of CoverTab[108953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:192
		_go_fuzz_dep_.CoverTab[108954]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:192
		// _ = "end of CoverTab[108954]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:192
	// _ = "end of CoverTab[108937]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:192
	_go_fuzz_dep_.CoverTab[108938]++
												if len(p.Enum) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:193
		_go_fuzz_dep_.CoverTab[108955]++
													s += ",enum=" + p.Enum
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:194
		// _ = "end of CoverTab[108955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:195
		_go_fuzz_dep_.CoverTab[108956]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:195
		// _ = "end of CoverTab[108956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:195
	// _ = "end of CoverTab[108938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:195
	_go_fuzz_dep_.CoverTab[108939]++
												if p.HasDefault {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:196
		_go_fuzz_dep_.CoverTab[108957]++
													s += ",def=" + p.Default
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:197
		// _ = "end of CoverTab[108957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:198
		_go_fuzz_dep_.CoverTab[108958]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:198
		// _ = "end of CoverTab[108958]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:198
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:198
	// _ = "end of CoverTab[108939]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:198
	_go_fuzz_dep_.CoverTab[108940]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:199
	// _ = "end of CoverTab[108940]"
}

// Parse populates p by parsing a string in the protobuf struct field tag style.
func (p *Properties) Parse(s string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:203
	_go_fuzz_dep_.CoverTab[108959]++

												fields := strings.Split(s, ",")
												if len(fields) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:206
		_go_fuzz_dep_.CoverTab[108963]++
													log.Printf("proto: tag has too few fields: %q", s)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:208
		// _ = "end of CoverTab[108963]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:209
		_go_fuzz_dep_.CoverTab[108964]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:209
		// _ = "end of CoverTab[108964]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:209
	// _ = "end of CoverTab[108959]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:209
	_go_fuzz_dep_.CoverTab[108960]++

												p.Wire = fields[0]
												switch p.Wire {
	case "varint":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:213
		_go_fuzz_dep_.CoverTab[108965]++
													p.WireType = WireVarint
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:214
		// _ = "end of CoverTab[108965]"
	case "fixed32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:215
		_go_fuzz_dep_.CoverTab[108966]++
													p.WireType = WireFixed32
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:216
		// _ = "end of CoverTab[108966]"
	case "fixed64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:217
		_go_fuzz_dep_.CoverTab[108967]++
													p.WireType = WireFixed64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:218
		// _ = "end of CoverTab[108967]"
	case "zigzag32":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:219
		_go_fuzz_dep_.CoverTab[108968]++
													p.WireType = WireVarint
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:220
		// _ = "end of CoverTab[108968]"
	case "zigzag64":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:221
		_go_fuzz_dep_.CoverTab[108969]++
													p.WireType = WireVarint
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:222
		// _ = "end of CoverTab[108969]"
	case "bytes", "group":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:223
		_go_fuzz_dep_.CoverTab[108970]++
													p.WireType = WireBytes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:224
		// _ = "end of CoverTab[108970]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:226
		_go_fuzz_dep_.CoverTab[108971]++
													log.Printf("proto: tag has unknown wire type: %q", s)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:228
		// _ = "end of CoverTab[108971]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:229
	// _ = "end of CoverTab[108960]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:229
	_go_fuzz_dep_.CoverTab[108961]++

												var err error
												p.Tag, err = strconv.Atoi(fields[1])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:233
		_go_fuzz_dep_.CoverTab[108972]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:234
		// _ = "end of CoverTab[108972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:235
		_go_fuzz_dep_.CoverTab[108973]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:235
		// _ = "end of CoverTab[108973]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:235
	// _ = "end of CoverTab[108961]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:235
	_go_fuzz_dep_.CoverTab[108962]++

outer:
	for i := 2; i < len(fields); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:238
		_go_fuzz_dep_.CoverTab[108974]++
													f := fields[i]
													switch {
		case f == "req":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:241
			_go_fuzz_dep_.CoverTab[108975]++
														p.Required = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:242
			// _ = "end of CoverTab[108975]"
		case f == "opt":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:243
			_go_fuzz_dep_.CoverTab[108976]++
														p.Optional = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:244
			// _ = "end of CoverTab[108976]"
		case f == "rep":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:245
			_go_fuzz_dep_.CoverTab[108977]++
														p.Repeated = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:246
			// _ = "end of CoverTab[108977]"
		case f == "packed":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:247
			_go_fuzz_dep_.CoverTab[108978]++
														p.Packed = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:248
			// _ = "end of CoverTab[108978]"
		case strings.HasPrefix(f, "name="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:249
			_go_fuzz_dep_.CoverTab[108979]++
														p.OrigName = f[5:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:250
			// _ = "end of CoverTab[108979]"
		case strings.HasPrefix(f, "json="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:251
			_go_fuzz_dep_.CoverTab[108980]++
														p.JSONName = f[5:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:252
			// _ = "end of CoverTab[108980]"
		case strings.HasPrefix(f, "enum="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:253
			_go_fuzz_dep_.CoverTab[108981]++
														p.Enum = f[5:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:254
			// _ = "end of CoverTab[108981]"
		case f == "proto3":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:255
			_go_fuzz_dep_.CoverTab[108982]++
														p.proto3 = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:256
			// _ = "end of CoverTab[108982]"
		case f == "oneof":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:257
			_go_fuzz_dep_.CoverTab[108983]++
														p.oneof = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:258
			// _ = "end of CoverTab[108983]"
		case strings.HasPrefix(f, "def="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:259
			_go_fuzz_dep_.CoverTab[108984]++
														p.HasDefault = true
														p.Default = f[4:]
														if i+1 < len(fields) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:262
				_go_fuzz_dep_.CoverTab[108992]++

															p.Default += "," + strings.Join(fields[i+1:], ",")
															break outer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:265
				// _ = "end of CoverTab[108992]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:266
				_go_fuzz_dep_.CoverTab[108993]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:266
				// _ = "end of CoverTab[108993]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:266
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:266
			// _ = "end of CoverTab[108984]"
		case strings.HasPrefix(f, "embedded="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:267
			_go_fuzz_dep_.CoverTab[108985]++
														p.OrigName = strings.Split(f, "=")[1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:268
			// _ = "end of CoverTab[108985]"
		case strings.HasPrefix(f, "customtype="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:269
			_go_fuzz_dep_.CoverTab[108986]++
														p.CustomType = strings.Split(f, "=")[1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:270
			// _ = "end of CoverTab[108986]"
		case strings.HasPrefix(f, "casttype="):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:271
			_go_fuzz_dep_.CoverTab[108987]++
														p.CastType = strings.Split(f, "=")[1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:272
			// _ = "end of CoverTab[108987]"
		case f == "stdtime":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:273
			_go_fuzz_dep_.CoverTab[108988]++
														p.StdTime = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:274
			// _ = "end of CoverTab[108988]"
		case f == "stdduration":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:275
			_go_fuzz_dep_.CoverTab[108989]++
														p.StdDuration = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:276
			// _ = "end of CoverTab[108989]"
		case f == "wktptr":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:277
			_go_fuzz_dep_.CoverTab[108990]++
														p.WktPointer = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:278
			// _ = "end of CoverTab[108990]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:278
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:278
			_go_fuzz_dep_.CoverTab[108991]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:278
			// _ = "end of CoverTab[108991]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:279
		// _ = "end of CoverTab[108974]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:280
	// _ = "end of CoverTab[108962]"
}

var protoMessageType = reflect.TypeOf((*Message)(nil)).Elem()

// setFieldProps initializes the field properties for submessages and maps.
func (p *Properties) setFieldProps(typ reflect.Type, f *reflect.StructField, lockGetProp bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:286
	_go_fuzz_dep_.CoverTab[108994]++
												isMap := typ.Kind() == reflect.Map
												if len(p.CustomType) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:288
		_go_fuzz_dep_.CoverTab[109000]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:288
		return !isMap
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:288
		// _ = "end of CoverTab[109000]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:288
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:288
		_go_fuzz_dep_.CoverTab[109001]++
													p.ctype = typ
													p.setTag(lockGetProp)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:291
		// _ = "end of CoverTab[109001]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:292
		_go_fuzz_dep_.CoverTab[109002]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:292
		// _ = "end of CoverTab[109002]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:292
	// _ = "end of CoverTab[108994]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:292
	_go_fuzz_dep_.CoverTab[108995]++
												if p.StdTime && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:293
		_go_fuzz_dep_.CoverTab[109003]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:293
		return !isMap
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:293
		// _ = "end of CoverTab[109003]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:293
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:293
		_go_fuzz_dep_.CoverTab[109004]++
													p.setTag(lockGetProp)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:295
		// _ = "end of CoverTab[109004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:296
		_go_fuzz_dep_.CoverTab[109005]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:296
		// _ = "end of CoverTab[109005]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:296
	// _ = "end of CoverTab[108995]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:296
	_go_fuzz_dep_.CoverTab[108996]++
												if p.StdDuration && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:297
		_go_fuzz_dep_.CoverTab[109006]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:297
		return !isMap
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:297
		// _ = "end of CoverTab[109006]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:297
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:297
		_go_fuzz_dep_.CoverTab[109007]++
													p.setTag(lockGetProp)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:299
		// _ = "end of CoverTab[109007]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:300
		_go_fuzz_dep_.CoverTab[109008]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:300
		// _ = "end of CoverTab[109008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:300
	// _ = "end of CoverTab[108996]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:300
	_go_fuzz_dep_.CoverTab[108997]++
												if p.WktPointer && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:301
		_go_fuzz_dep_.CoverTab[109009]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:301
		return !isMap
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:301
		// _ = "end of CoverTab[109009]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:301
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:301
		_go_fuzz_dep_.CoverTab[109010]++
													p.setTag(lockGetProp)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:303
		// _ = "end of CoverTab[109010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:304
		_go_fuzz_dep_.CoverTab[109011]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:304
		// _ = "end of CoverTab[109011]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:304
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:304
	// _ = "end of CoverTab[108997]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:304
	_go_fuzz_dep_.CoverTab[108998]++
												switch t1 := typ; t1.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:306
		_go_fuzz_dep_.CoverTab[109012]++
													p.stype = typ
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:307
		// _ = "end of CoverTab[109012]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:308
		_go_fuzz_dep_.CoverTab[109013]++
													if t1.Elem().Kind() == reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:309
			_go_fuzz_dep_.CoverTab[109018]++
														p.stype = t1.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:310
			// _ = "end of CoverTab[109018]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:311
			_go_fuzz_dep_.CoverTab[109019]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:311
			// _ = "end of CoverTab[109019]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:311
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:311
		// _ = "end of CoverTab[109013]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:312
		_go_fuzz_dep_.CoverTab[109014]++
													switch t2 := t1.Elem(); t2.Kind() {
		case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:314
			_go_fuzz_dep_.CoverTab[109020]++
														switch t3 := t2.Elem(); t3.Kind() {
			case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:316
				_go_fuzz_dep_.CoverTab[109023]++
															p.stype = t3
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:317
				// _ = "end of CoverTab[109023]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:317
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:317
				_go_fuzz_dep_.CoverTab[109024]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:317
				// _ = "end of CoverTab[109024]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:318
			// _ = "end of CoverTab[109020]"
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:319
			_go_fuzz_dep_.CoverTab[109021]++
														p.stype = t2
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:320
			// _ = "end of CoverTab[109021]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:320
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:320
			_go_fuzz_dep_.CoverTab[109022]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:320
			// _ = "end of CoverTab[109022]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:321
		// _ = "end of CoverTab[109014]"

	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:323
		_go_fuzz_dep_.CoverTab[109015]++

													p.mtype = t1
													p.MapKeyProp = &Properties{}
													p.MapKeyProp.init(reflect.PtrTo(p.mtype.Key()), "Key", f.Tag.Get("protobuf_key"), nil, lockGetProp)
													p.MapValProp = &Properties{}
													vtype := p.mtype.Elem()
													if vtype.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:330
			_go_fuzz_dep_.CoverTab[109025]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:330
			return vtype.Kind() != reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:330
			// _ = "end of CoverTab[109025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:330
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:330
			_go_fuzz_dep_.CoverTab[109026]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:333
			vtype = reflect.PtrTo(vtype)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:333
			// _ = "end of CoverTab[109026]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:334
			_go_fuzz_dep_.CoverTab[109027]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:334
			// _ = "end of CoverTab[109027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:334
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:334
		// _ = "end of CoverTab[109015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:334
		_go_fuzz_dep_.CoverTab[109016]++

													p.MapValProp.CustomType = p.CustomType
													p.MapValProp.StdDuration = p.StdDuration
													p.MapValProp.StdTime = p.StdTime
													p.MapValProp.WktPointer = p.WktPointer
													p.MapValProp.init(vtype, "Value", f.Tag.Get("protobuf_val"), nil, lockGetProp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:340
		// _ = "end of CoverTab[109016]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:340
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:340
		_go_fuzz_dep_.CoverTab[109017]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:340
		// _ = "end of CoverTab[109017]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:341
	// _ = "end of CoverTab[108998]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:341
	_go_fuzz_dep_.CoverTab[108999]++
												p.setTag(lockGetProp)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:342
	// _ = "end of CoverTab[108999]"
}

func (p *Properties) setTag(lockGetProp bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:345
	_go_fuzz_dep_.CoverTab[109028]++
												if p.stype != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:346
		_go_fuzz_dep_.CoverTab[109029]++
													if lockGetProp {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:347
			_go_fuzz_dep_.CoverTab[109030]++
														p.sprop = GetProperties(p.stype)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:348
			// _ = "end of CoverTab[109030]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:349
			_go_fuzz_dep_.CoverTab[109031]++
														p.sprop = getPropertiesLocked(p.stype)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:350
			// _ = "end of CoverTab[109031]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:351
		// _ = "end of CoverTab[109029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:352
		_go_fuzz_dep_.CoverTab[109032]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:352
		// _ = "end of CoverTab[109032]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:352
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:352
	// _ = "end of CoverTab[109028]"
}

var (
	marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()
)

// Init populates the properties from a protocol buffer struct tag.
func (p *Properties) Init(typ reflect.Type, name, tag string, f *reflect.StructField) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:360
	_go_fuzz_dep_.CoverTab[109033]++
												p.init(typ, name, tag, f, true)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:361
	// _ = "end of CoverTab[109033]"
}

func (p *Properties) init(typ reflect.Type, name, tag string, f *reflect.StructField, lockGetProp bool) {

	p.Name = name
	p.OrigName = name
	if tag == "" {
		return
	}
	p.Parse(tag)
	p.setFieldProps(typ, f, lockGetProp)
}

var (
	propertiesMu	sync.RWMutex
	propertiesMap	= make(map[reflect.Type]*StructProperties)
)

// GetProperties returns the list of properties for the type represented by t.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:380
// t must represent a generated struct type of a protocol message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:382
func GetProperties(t reflect.Type) *StructProperties {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:382
	_go_fuzz_dep_.CoverTab[109034]++
												if t.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:383
		_go_fuzz_dep_.CoverTab[109037]++
													panic("proto: type must have kind struct")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:384
		// _ = "end of CoverTab[109037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:385
		_go_fuzz_dep_.CoverTab[109038]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:385
		// _ = "end of CoverTab[109038]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:385
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:385
	// _ = "end of CoverTab[109034]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:385
	_go_fuzz_dep_.CoverTab[109035]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:389
	propertiesMu.RLock()
	sprop, ok := propertiesMap[t]
	propertiesMu.RUnlock()
	if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:392
		_go_fuzz_dep_.CoverTab[109039]++
													return sprop
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:393
		// _ = "end of CoverTab[109039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:394
		_go_fuzz_dep_.CoverTab[109040]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:394
		// _ = "end of CoverTab[109040]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:394
	// _ = "end of CoverTab[109035]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:394
	_go_fuzz_dep_.CoverTab[109036]++

												propertiesMu.Lock()
												sprop = getPropertiesLocked(t)
												propertiesMu.Unlock()
												return sprop
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:399
	// _ = "end of CoverTab[109036]"
}

type (
	oneofFuncsIface	interface {
		XXX_OneofFuncs() (func(Message, *Buffer) error, func(Message, int, int, *Buffer) (bool, error), func(Message) int, []interface{})
	}
	oneofWrappersIface	interface {
		XXX_OneofWrappers() []interface{}
	}
)

// getPropertiesLocked requires that propertiesMu is held.
func getPropertiesLocked(t reflect.Type) *StructProperties {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:412
	_go_fuzz_dep_.CoverTab[109041]++
												if prop, ok := propertiesMap[t]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:413
		_go_fuzz_dep_.CoverTab[109046]++
													return prop
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:414
		// _ = "end of CoverTab[109046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:415
		_go_fuzz_dep_.CoverTab[109047]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:415
		// _ = "end of CoverTab[109047]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:415
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:415
	// _ = "end of CoverTab[109041]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:415
	_go_fuzz_dep_.CoverTab[109042]++

												prop := new(StructProperties)

												propertiesMap[t] = prop

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:422
	prop.Prop = make([]*Properties, t.NumField())
	prop.order = make([]int, t.NumField())

	isOneofMessage := false
	for i := 0; i < t.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:426
		_go_fuzz_dep_.CoverTab[109048]++
													f := t.Field(i)
													p := new(Properties)
													name := f.Name
													p.init(f.Type, name, f.Tag.Get("protobuf"), &f, false)

													oneof := f.Tag.Get("protobuf_oneof")
													if oneof != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:433
			_go_fuzz_dep_.CoverTab[109050]++
														isOneofMessage = true

														p.OrigName = oneof
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:436
			// _ = "end of CoverTab[109050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:437
			_go_fuzz_dep_.CoverTab[109051]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:437
			// _ = "end of CoverTab[109051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:437
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:437
		// _ = "end of CoverTab[109048]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:437
		_go_fuzz_dep_.CoverTab[109049]++
													prop.Prop[i] = p
													prop.order[i] = i
													if debug {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:440
			_go_fuzz_dep_.CoverTab[109052]++
														print(i, " ", f.Name, " ", t.String(), " ")
														if p.Tag > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:442
				_go_fuzz_dep_.CoverTab[109054]++
															print(p.String())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:443
				// _ = "end of CoverTab[109054]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:444
				_go_fuzz_dep_.CoverTab[109055]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:444
				// _ = "end of CoverTab[109055]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:444
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:444
			// _ = "end of CoverTab[109052]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:444
			_go_fuzz_dep_.CoverTab[109053]++
														print("\n")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:445
			// _ = "end of CoverTab[109053]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:446
			_go_fuzz_dep_.CoverTab[109056]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:446
			// _ = "end of CoverTab[109056]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:446
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:446
		// _ = "end of CoverTab[109049]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:447
	// _ = "end of CoverTab[109042]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:447
	_go_fuzz_dep_.CoverTab[109043]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:450
	sort.Sort(prop)

	if isOneofMessage {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:452
		_go_fuzz_dep_.CoverTab[109057]++
													var oots []interface{}
													switch m := reflect.Zero(reflect.PtrTo(t)).Interface().(type) {
		case oneofFuncsIface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:455
			_go_fuzz_dep_.CoverTab[109059]++
														_, _, _, oots = m.XXX_OneofFuncs()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:456
			// _ = "end of CoverTab[109059]"
		case oneofWrappersIface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:457
			_go_fuzz_dep_.CoverTab[109060]++
														oots = m.XXX_OneofWrappers()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:458
			// _ = "end of CoverTab[109060]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:459
		// _ = "end of CoverTab[109057]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:459
		_go_fuzz_dep_.CoverTab[109058]++
													if len(oots) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:460
			_go_fuzz_dep_.CoverTab[109061]++

														prop.OneofTypes = make(map[string]*OneofProperties)
														for _, oot := range oots {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:463
				_go_fuzz_dep_.CoverTab[109062]++
															oop := &OneofProperties{
					Type:	reflect.ValueOf(oot).Type(),
					Prop:	new(Properties),
				}
															sft := oop.Type.Elem().Field(0)
															oop.Prop.Name = sft.Name
															oop.Prop.Parse(sft.Tag.Get("protobuf"))

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:473
				for i := 0; i < t.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:473
					_go_fuzz_dep_.CoverTab[109064]++
																f := t.Field(i)
																if f.Type.Kind() != reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:475
						_go_fuzz_dep_.CoverTab[109067]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:476
						// _ = "end of CoverTab[109067]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:477
						_go_fuzz_dep_.CoverTab[109068]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:477
						// _ = "end of CoverTab[109068]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:477
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:477
					// _ = "end of CoverTab[109064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:477
					_go_fuzz_dep_.CoverTab[109065]++
																if !oop.Type.AssignableTo(f.Type) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:478
						_go_fuzz_dep_.CoverTab[109069]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:479
						// _ = "end of CoverTab[109069]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:480
						_go_fuzz_dep_.CoverTab[109070]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:480
						// _ = "end of CoverTab[109070]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:480
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:480
					// _ = "end of CoverTab[109065]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:480
					_go_fuzz_dep_.CoverTab[109066]++
																oop.Field = i
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:482
					// _ = "end of CoverTab[109066]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:483
				// _ = "end of CoverTab[109062]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:483
				_go_fuzz_dep_.CoverTab[109063]++
															prop.OneofTypes[oop.Prop.OrigName] = oop
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:484
				// _ = "end of CoverTab[109063]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:485
			// _ = "end of CoverTab[109061]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:486
			_go_fuzz_dep_.CoverTab[109071]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:486
			// _ = "end of CoverTab[109071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:486
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:486
		// _ = "end of CoverTab[109058]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:487
		_go_fuzz_dep_.CoverTab[109072]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:487
		// _ = "end of CoverTab[109072]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:487
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:487
	// _ = "end of CoverTab[109043]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:487
	_go_fuzz_dep_.CoverTab[109044]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:491
	reqCount := 0
	prop.decoderOrigNames = make(map[string]int)
	for i, p := range prop.Prop {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:493
		_go_fuzz_dep_.CoverTab[109073]++
													if strings.HasPrefix(p.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:494
			_go_fuzz_dep_.CoverTab[109076]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:497
			continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:497
			// _ = "end of CoverTab[109076]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:498
			_go_fuzz_dep_.CoverTab[109077]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:498
			// _ = "end of CoverTab[109077]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:498
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:498
		// _ = "end of CoverTab[109073]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:498
		_go_fuzz_dep_.CoverTab[109074]++
													if p.Required {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:499
			_go_fuzz_dep_.CoverTab[109078]++
														reqCount++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:500
			// _ = "end of CoverTab[109078]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:501
			_go_fuzz_dep_.CoverTab[109079]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:501
			// _ = "end of CoverTab[109079]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:501
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:501
		// _ = "end of CoverTab[109074]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:501
		_go_fuzz_dep_.CoverTab[109075]++
													prop.decoderTags.put(p.Tag, i)
													prop.decoderOrigNames[p.OrigName] = i
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:503
		// _ = "end of CoverTab[109075]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:504
	// _ = "end of CoverTab[109044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:504
	_go_fuzz_dep_.CoverTab[109045]++
												prop.reqCount = reqCount

												return prop
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:507
	// _ = "end of CoverTab[109045]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:513
var enumValueMaps = make(map[string]map[string]int32)
var enumStringMaps = make(map[string]map[int32]string)

// RegisterEnum is called from the generated code to install the enum descriptor
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:516
// maps into the global table to aid parsing text format protocol buffers.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:518
func RegisterEnum(typeName string, unusedNameMap map[int32]string, valueMap map[string]int32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:518
	_go_fuzz_dep_.CoverTab[109080]++
												if _, ok := enumValueMaps[typeName]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:519
		_go_fuzz_dep_.CoverTab[109083]++
													panic("proto: duplicate enum registered: " + typeName)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:520
		// _ = "end of CoverTab[109083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:521
		_go_fuzz_dep_.CoverTab[109084]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:521
		// _ = "end of CoverTab[109084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:521
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:521
	// _ = "end of CoverTab[109080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:521
	_go_fuzz_dep_.CoverTab[109081]++
												enumValueMaps[typeName] = valueMap
												if _, ok := enumStringMaps[typeName]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:523
		_go_fuzz_dep_.CoverTab[109085]++
													panic("proto: duplicate enum registered: " + typeName)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:524
		// _ = "end of CoverTab[109085]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:525
		_go_fuzz_dep_.CoverTab[109086]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:525
		// _ = "end of CoverTab[109086]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:525
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:525
	// _ = "end of CoverTab[109081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:525
	_go_fuzz_dep_.CoverTab[109082]++
												enumStringMaps[typeName] = unusedNameMap
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:526
	// _ = "end of CoverTab[109082]"
}

// EnumValueMap returns the mapping from names to integers of the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:529
// enum type enumType, or a nil if not found.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:531
func EnumValueMap(enumType string) map[string]int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:531
	_go_fuzz_dep_.CoverTab[109087]++
												return enumValueMaps[enumType]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:532
	// _ = "end of CoverTab[109087]"
}

// A registry of all linked message types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:535
// The string is a fully-qualified proto name ("pkg.Message").
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:537
var (
	protoTypedNils	= make(map[string]Message)	// a map from proto names to typed nil pointers
	protoMapTypes	= make(map[string]reflect.Type)	// a map from proto names to map types
	revProtoTypes	= make(map[reflect.Type]string)
)

// RegisterType is called from generated code and maps from the fully qualified
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:543
// proto name to the type (pointer to struct) of the protocol buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:545
func RegisterType(x Message, name string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:545
	_go_fuzz_dep_.CoverTab[109088]++
												if _, ok := protoTypedNils[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:546
		_go_fuzz_dep_.CoverTab[109091]++

													log.Printf("proto: duplicate proto type registered: %s", name)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:549
		// _ = "end of CoverTab[109091]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:550
		_go_fuzz_dep_.CoverTab[109092]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:550
		// _ = "end of CoverTab[109092]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:550
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:550
	// _ = "end of CoverTab[109088]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:550
	_go_fuzz_dep_.CoverTab[109089]++
												t := reflect.TypeOf(x)
												if v := reflect.ValueOf(x); v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:552
		_go_fuzz_dep_.CoverTab[109093]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:552
		return v.Pointer() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:552
		// _ = "end of CoverTab[109093]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:552
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:552
		_go_fuzz_dep_.CoverTab[109094]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:555
		protoTypedNils[name] = x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:555
		// _ = "end of CoverTab[109094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:556
		_go_fuzz_dep_.CoverTab[109095]++
													protoTypedNils[name] = reflect.Zero(t).Interface().(Message)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:557
		// _ = "end of CoverTab[109095]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:558
	// _ = "end of CoverTab[109089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:558
	_go_fuzz_dep_.CoverTab[109090]++
												revProtoTypes[t] = name
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:559
	// _ = "end of CoverTab[109090]"
}

// RegisterMapType is called from generated code and maps from the fully qualified
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:562
// proto name to the native map type of the proto map definition.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:564
func RegisterMapType(x interface{}, name string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:564
	_go_fuzz_dep_.CoverTab[109096]++
												if reflect.TypeOf(x).Kind() != reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:565
		_go_fuzz_dep_.CoverTab[109099]++
													panic(fmt.Sprintf("RegisterMapType(%T, %q); want map", x, name))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:566
		// _ = "end of CoverTab[109099]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:567
		_go_fuzz_dep_.CoverTab[109100]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:567
		// _ = "end of CoverTab[109100]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:567
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:567
	// _ = "end of CoverTab[109096]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:567
	_go_fuzz_dep_.CoverTab[109097]++
												if _, ok := protoMapTypes[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:568
		_go_fuzz_dep_.CoverTab[109101]++
													log.Printf("proto: duplicate proto type registered: %s", name)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:570
		// _ = "end of CoverTab[109101]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:571
		_go_fuzz_dep_.CoverTab[109102]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:571
		// _ = "end of CoverTab[109102]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:571
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:571
	// _ = "end of CoverTab[109097]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:571
	_go_fuzz_dep_.CoverTab[109098]++
												t := reflect.TypeOf(x)
												protoMapTypes[name] = t
												revProtoTypes[t] = name
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:574
	// _ = "end of CoverTab[109098]"
}

// MessageName returns the fully-qualified proto name for the given message type.
func MessageName(x Message) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:578
	_go_fuzz_dep_.CoverTab[109103]++
												type xname interface {
		XXX_MessageName() string
	}
	if m, ok := x.(xname); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:582
		_go_fuzz_dep_.CoverTab[109105]++
													return m.XXX_MessageName()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:583
		// _ = "end of CoverTab[109105]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:584
		_go_fuzz_dep_.CoverTab[109106]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:584
		// _ = "end of CoverTab[109106]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:584
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:584
	// _ = "end of CoverTab[109103]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:584
	_go_fuzz_dep_.CoverTab[109104]++
												return revProtoTypes[reflect.TypeOf(x)]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:585
	// _ = "end of CoverTab[109104]"
}

// MessageType returns the message type (pointer to struct) for a named message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:588
// The type is not guaranteed to implement proto.Message if the name refers to a
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:588
// map entry.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:591
func MessageType(name string) reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:591
	_go_fuzz_dep_.CoverTab[109107]++
												if t, ok := protoTypedNils[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:592
		_go_fuzz_dep_.CoverTab[109109]++
													return reflect.TypeOf(t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:593
		// _ = "end of CoverTab[109109]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:594
		_go_fuzz_dep_.CoverTab[109110]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:594
		// _ = "end of CoverTab[109110]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:594
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:594
	// _ = "end of CoverTab[109107]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:594
	_go_fuzz_dep_.CoverTab[109108]++
												return protoMapTypes[name]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:595
	// _ = "end of CoverTab[109108]"
}

// A registry of all linked proto files.
var (
	protoFiles = make(map[string][]byte)	// file name => fileDescriptor
)

// RegisterFile is called from generated code and maps from the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:603
// full file name of a .proto file to its compressed FileDescriptorProto.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:605
func RegisterFile(filename string, fileDescriptor []byte) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:605
	_go_fuzz_dep_.CoverTab[109111]++
												protoFiles[filename] = fileDescriptor
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:606
	// _ = "end of CoverTab[109111]"
}

// FileDescriptor returns the compressed FileDescriptorProto for a .proto file.
func FileDescriptor(filename string) []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:610
	_go_fuzz_dep_.CoverTab[109112]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:610
	return protoFiles[filename]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:610
	// _ = "end of CoverTab[109112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:610
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:610
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/properties.go:610
var _ = _go_fuzz_dep_.CoverTab
