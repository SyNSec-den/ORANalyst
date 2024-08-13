// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:29
)

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"sync"
)

type extensionsBytes interface {
	Message
	ExtensionRangeArray() []ExtensionRange
	GetExtensions() *[]byte
}

type slowExtensionAdapter struct {
	extensionsBytes
}

func (s slowExtensionAdapter) extensionsWrite() map[int32]Extension {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:52
	_go_fuzz_dep_.CoverTab[108378]++
												panic("Please report a bug to github.com/gogo/protobuf if you see this message: Writing extensions is not supported for extensions stored in a byte slice field.")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:53
	// _ = "end of CoverTab[108378]"
}

func (s slowExtensionAdapter) extensionsRead() (map[int32]Extension, sync.Locker) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:56
	_go_fuzz_dep_.CoverTab[108379]++
												b := s.GetExtensions()
												m, err := BytesToExtensionsMap(*b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:59
		_go_fuzz_dep_.CoverTab[108381]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:60
		// _ = "end of CoverTab[108381]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:61
		_go_fuzz_dep_.CoverTab[108382]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:61
		// _ = "end of CoverTab[108382]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:61
	// _ = "end of CoverTab[108379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:61
	_go_fuzz_dep_.CoverTab[108380]++
												return m, notLocker{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:62
	// _ = "end of CoverTab[108380]"
}

func GetBoolExtension(pb Message, extension *ExtensionDesc, ifnotset bool) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:65
	_go_fuzz_dep_.CoverTab[108383]++
												if reflect.ValueOf(pb).IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:66
		_go_fuzz_dep_.CoverTab[108388]++
													return ifnotset
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:67
		// _ = "end of CoverTab[108388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:68
		_go_fuzz_dep_.CoverTab[108389]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:68
		// _ = "end of CoverTab[108389]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:68
	// _ = "end of CoverTab[108383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:68
	_go_fuzz_dep_.CoverTab[108384]++
												value, err := GetExtension(pb, extension)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:70
		_go_fuzz_dep_.CoverTab[108390]++
													return ifnotset
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:71
		// _ = "end of CoverTab[108390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:72
		_go_fuzz_dep_.CoverTab[108391]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:72
		// _ = "end of CoverTab[108391]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:72
	// _ = "end of CoverTab[108384]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:72
	_go_fuzz_dep_.CoverTab[108385]++
												if value == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:73
		_go_fuzz_dep_.CoverTab[108392]++
													return ifnotset
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:74
		// _ = "end of CoverTab[108392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:75
		_go_fuzz_dep_.CoverTab[108393]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:75
		// _ = "end of CoverTab[108393]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:75
	// _ = "end of CoverTab[108385]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:75
	_go_fuzz_dep_.CoverTab[108386]++
												if value.(*bool) == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:76
		_go_fuzz_dep_.CoverTab[108394]++
													return ifnotset
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:77
		// _ = "end of CoverTab[108394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:78
		_go_fuzz_dep_.CoverTab[108395]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:78
		// _ = "end of CoverTab[108395]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:78
	// _ = "end of CoverTab[108386]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:78
	_go_fuzz_dep_.CoverTab[108387]++
												return *(value.(*bool))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:79
	// _ = "end of CoverTab[108387]"
}

func (this *Extension) Equal(that *Extension) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:82
	_go_fuzz_dep_.CoverTab[108396]++
												if err := this.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:83
		_go_fuzz_dep_.CoverTab[108399]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:84
		// _ = "end of CoverTab[108399]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:85
		_go_fuzz_dep_.CoverTab[108400]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:85
		// _ = "end of CoverTab[108400]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:85
	// _ = "end of CoverTab[108396]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:85
	_go_fuzz_dep_.CoverTab[108397]++
												if err := that.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:86
		_go_fuzz_dep_.CoverTab[108401]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:87
		// _ = "end of CoverTab[108401]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:88
		_go_fuzz_dep_.CoverTab[108402]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:88
		// _ = "end of CoverTab[108402]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:88
	// _ = "end of CoverTab[108397]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:88
	_go_fuzz_dep_.CoverTab[108398]++
												return bytes.Equal(this.enc, that.enc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:89
	// _ = "end of CoverTab[108398]"
}

func (this *Extension) Compare(that *Extension) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:92
	_go_fuzz_dep_.CoverTab[108403]++
												if err := this.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:93
		_go_fuzz_dep_.CoverTab[108406]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:94
		// _ = "end of CoverTab[108406]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:95
		_go_fuzz_dep_.CoverTab[108407]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:95
		// _ = "end of CoverTab[108407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:95
	// _ = "end of CoverTab[108403]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:95
	_go_fuzz_dep_.CoverTab[108404]++
												if err := that.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:96
		_go_fuzz_dep_.CoverTab[108408]++
													return -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:97
		// _ = "end of CoverTab[108408]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:98
		_go_fuzz_dep_.CoverTab[108409]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:98
		// _ = "end of CoverTab[108409]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:98
	// _ = "end of CoverTab[108404]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:98
	_go_fuzz_dep_.CoverTab[108405]++
												return bytes.Compare(this.enc, that.enc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:99
	// _ = "end of CoverTab[108405]"
}

func SizeOfInternalExtension(m extendableProto) (n int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:102
	_go_fuzz_dep_.CoverTab[108410]++
												info := getMarshalInfo(reflect.TypeOf(m))
												return info.sizeV1Extensions(m.extensionsWrite())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:104
	// _ = "end of CoverTab[108410]"
}

type sortableMapElem struct {
	field	int32
	ext	Extension
}

func newSortableExtensionsFromMap(m map[int32]Extension) sortableExtensions {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:112
	_go_fuzz_dep_.CoverTab[108411]++
												s := make(sortableExtensions, 0, len(m))
												for k, v := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:114
		_go_fuzz_dep_.CoverTab[108413]++
													s = append(s, &sortableMapElem{field: k, ext: v})
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:115
		// _ = "end of CoverTab[108413]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:116
	// _ = "end of CoverTab[108411]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:116
	_go_fuzz_dep_.CoverTab[108412]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:117
	// _ = "end of CoverTab[108412]"
}

type sortableExtensions []*sortableMapElem

func (this sortableExtensions) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:122
	_go_fuzz_dep_.CoverTab[108414]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:122
	return len(this)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:122
	// _ = "end of CoverTab[108414]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:122
}

func (this sortableExtensions) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:124
	_go_fuzz_dep_.CoverTab[108415]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:124
	this[i], this[j] = this[j], this[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:124
	// _ = "end of CoverTab[108415]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:124
}

func (this sortableExtensions) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:126
	_go_fuzz_dep_.CoverTab[108416]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:126
	return this[i].field < this[j].field
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:126
	// _ = "end of CoverTab[108416]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:126
}

func (this sortableExtensions) String() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:128
	_go_fuzz_dep_.CoverTab[108417]++
												sort.Sort(this)
												ss := make([]string, len(this))
												for i := range this {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:131
		_go_fuzz_dep_.CoverTab[108419]++
													ss[i] = fmt.Sprintf("%d: %v", this[i].field, this[i].ext)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:132
		// _ = "end of CoverTab[108419]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:133
	// _ = "end of CoverTab[108417]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:133
	_go_fuzz_dep_.CoverTab[108418]++
												return "map[" + strings.Join(ss, ",") + "]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:134
	// _ = "end of CoverTab[108418]"
}

func StringFromInternalExtension(m extendableProto) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:137
	_go_fuzz_dep_.CoverTab[108420]++
												return StringFromExtensionsMap(m.extensionsWrite())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:138
	// _ = "end of CoverTab[108420]"
}

func StringFromExtensionsMap(m map[int32]Extension) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:141
	_go_fuzz_dep_.CoverTab[108421]++
												return newSortableExtensionsFromMap(m).String()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:142
	// _ = "end of CoverTab[108421]"
}

func StringFromExtensionsBytes(ext []byte) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:145
	_go_fuzz_dep_.CoverTab[108422]++
												m, err := BytesToExtensionsMap(ext)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:147
		_go_fuzz_dep_.CoverTab[108424]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:148
		// _ = "end of CoverTab[108424]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:149
		_go_fuzz_dep_.CoverTab[108425]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:149
		// _ = "end of CoverTab[108425]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:149
	// _ = "end of CoverTab[108422]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:149
	_go_fuzz_dep_.CoverTab[108423]++
												return StringFromExtensionsMap(m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:150
	// _ = "end of CoverTab[108423]"
}

func EncodeInternalExtension(m extendableProto, data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:153
	_go_fuzz_dep_.CoverTab[108426]++
												return EncodeExtensionMap(m.extensionsWrite(), data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:154
	// _ = "end of CoverTab[108426]"
}

func EncodeInternalExtensionBackwards(m extendableProto, data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:157
	_go_fuzz_dep_.CoverTab[108427]++
												return EncodeExtensionMapBackwards(m.extensionsWrite(), data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:158
	// _ = "end of CoverTab[108427]"
}

func EncodeExtensionMap(m map[int32]Extension, data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:161
	_go_fuzz_dep_.CoverTab[108428]++
												o := 0
												for _, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:163
		_go_fuzz_dep_.CoverTab[108430]++
													if err := e.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:164
			_go_fuzz_dep_.CoverTab[108433]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:165
			// _ = "end of CoverTab[108433]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:166
			_go_fuzz_dep_.CoverTab[108434]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:166
			// _ = "end of CoverTab[108434]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:166
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:166
		// _ = "end of CoverTab[108430]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:166
		_go_fuzz_dep_.CoverTab[108431]++
													n := copy(data[o:], e.enc)
													if n != len(e.enc) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:168
			_go_fuzz_dep_.CoverTab[108435]++
														return 0, io.ErrShortBuffer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:169
			// _ = "end of CoverTab[108435]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:170
			_go_fuzz_dep_.CoverTab[108436]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:170
			// _ = "end of CoverTab[108436]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:170
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:170
		// _ = "end of CoverTab[108431]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:170
		_go_fuzz_dep_.CoverTab[108432]++
													o += n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:171
		// _ = "end of CoverTab[108432]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:172
	// _ = "end of CoverTab[108428]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:172
	_go_fuzz_dep_.CoverTab[108429]++
												return o, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:173
	// _ = "end of CoverTab[108429]"
}

func EncodeExtensionMapBackwards(m map[int32]Extension, data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:176
	_go_fuzz_dep_.CoverTab[108437]++
												o := 0
												end := len(data)
												for _, e := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:179
		_go_fuzz_dep_.CoverTab[108439]++
													if err := e.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:180
			_go_fuzz_dep_.CoverTab[108442]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:181
			// _ = "end of CoverTab[108442]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:182
			_go_fuzz_dep_.CoverTab[108443]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:182
			// _ = "end of CoverTab[108443]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:182
		// _ = "end of CoverTab[108439]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:182
		_go_fuzz_dep_.CoverTab[108440]++
													n := copy(data[end-len(e.enc):], e.enc)
													if n != len(e.enc) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:184
			_go_fuzz_dep_.CoverTab[108444]++
														return 0, io.ErrShortBuffer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:185
			// _ = "end of CoverTab[108444]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:186
			_go_fuzz_dep_.CoverTab[108445]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:186
			// _ = "end of CoverTab[108445]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:186
		// _ = "end of CoverTab[108440]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:186
		_go_fuzz_dep_.CoverTab[108441]++
													end -= n
													o += n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:188
		// _ = "end of CoverTab[108441]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:189
	// _ = "end of CoverTab[108437]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:189
	_go_fuzz_dep_.CoverTab[108438]++
												return o, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:190
	// _ = "end of CoverTab[108438]"
}

func GetRawExtension(m map[int32]Extension, id int32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:193
	_go_fuzz_dep_.CoverTab[108446]++
												e := m[id]
												if err := e.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:195
		_go_fuzz_dep_.CoverTab[108448]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:196
		// _ = "end of CoverTab[108448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:197
		_go_fuzz_dep_.CoverTab[108449]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:197
		// _ = "end of CoverTab[108449]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:197
	// _ = "end of CoverTab[108446]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:197
	_go_fuzz_dep_.CoverTab[108447]++
												return e.enc, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:198
	// _ = "end of CoverTab[108447]"
}

func size(buf []byte, wire int) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:201
	_go_fuzz_dep_.CoverTab[108450]++
												switch wire {
	case WireVarint:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:203
		_go_fuzz_dep_.CoverTab[108452]++
													_, n := DecodeVarint(buf)
													return n, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:205
		// _ = "end of CoverTab[108452]"
	case WireFixed64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:206
		_go_fuzz_dep_.CoverTab[108453]++
													return 8, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:207
		// _ = "end of CoverTab[108453]"
	case WireBytes:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:208
		_go_fuzz_dep_.CoverTab[108454]++
													v, n := DecodeVarint(buf)
													return int(v) + n, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:210
		// _ = "end of CoverTab[108454]"
	case WireFixed32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:211
		_go_fuzz_dep_.CoverTab[108455]++
													return 4, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:212
		// _ = "end of CoverTab[108455]"
	case WireStartGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:213
		_go_fuzz_dep_.CoverTab[108456]++
													offset := 0
													for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:215
			_go_fuzz_dep_.CoverTab[108458]++
														u, n := DecodeVarint(buf[offset:])
														fwire := int(u & 0x7)
														offset += n
														if fwire == WireEndGroup {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:219
				_go_fuzz_dep_.CoverTab[108461]++
															return offset, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:220
				// _ = "end of CoverTab[108461]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:221
				_go_fuzz_dep_.CoverTab[108462]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:221
				// _ = "end of CoverTab[108462]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:221
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:221
			// _ = "end of CoverTab[108458]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:221
			_go_fuzz_dep_.CoverTab[108459]++
														s, err := size(buf[offset:], wire)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:223
				_go_fuzz_dep_.CoverTab[108463]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:224
				// _ = "end of CoverTab[108463]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:225
				_go_fuzz_dep_.CoverTab[108464]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:225
				// _ = "end of CoverTab[108464]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:225
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:225
			// _ = "end of CoverTab[108459]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:225
			_go_fuzz_dep_.CoverTab[108460]++
														offset += s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:226
			// _ = "end of CoverTab[108460]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:227
		// _ = "end of CoverTab[108456]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:227
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:227
		_go_fuzz_dep_.CoverTab[108457]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:227
		// _ = "end of CoverTab[108457]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:228
	// _ = "end of CoverTab[108450]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:228
	_go_fuzz_dep_.CoverTab[108451]++
												return 0, fmt.Errorf("proto: can't get size for unknown wire type %d", wire)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:229
	// _ = "end of CoverTab[108451]"
}

func BytesToExtensionsMap(buf []byte) (map[int32]Extension, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:232
	_go_fuzz_dep_.CoverTab[108465]++
												m := make(map[int32]Extension)
												i := 0
												for i < len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:235
		_go_fuzz_dep_.CoverTab[108467]++
													tag, n := DecodeVarint(buf[i:])
													if n <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:237
			_go_fuzz_dep_.CoverTab[108470]++
														return nil, fmt.Errorf("unable to decode varint")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:238
			// _ = "end of CoverTab[108470]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:239
			_go_fuzz_dep_.CoverTab[108471]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:239
			// _ = "end of CoverTab[108471]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:239
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:239
		// _ = "end of CoverTab[108467]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:239
		_go_fuzz_dep_.CoverTab[108468]++
													fieldNum := int32(tag >> 3)
													wireType := int(tag & 0x7)
													l, err := size(buf[i+n:], wireType)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:243
			_go_fuzz_dep_.CoverTab[108472]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:244
			// _ = "end of CoverTab[108472]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:245
			_go_fuzz_dep_.CoverTab[108473]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:245
			// _ = "end of CoverTab[108473]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:245
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:245
		// _ = "end of CoverTab[108468]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:245
		_go_fuzz_dep_.CoverTab[108469]++
													end := i + int(l) + n
													m[int32(fieldNum)] = Extension{enc: buf[i:end]}
													i = end
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:248
		// _ = "end of CoverTab[108469]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:249
	// _ = "end of CoverTab[108465]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:249
	_go_fuzz_dep_.CoverTab[108466]++
												return m, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:250
	// _ = "end of CoverTab[108466]"
}

func NewExtension(e []byte) Extension {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:253
	_go_fuzz_dep_.CoverTab[108474]++
												ee := Extension{enc: make([]byte, len(e))}
												copy(ee.enc, e)
												return ee
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:256
	// _ = "end of CoverTab[108474]"
}

func AppendExtension(e Message, tag int32, buf []byte) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:259
	_go_fuzz_dep_.CoverTab[108475]++
												if ee, eok := e.(extensionsBytes); eok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:260
		_go_fuzz_dep_.CoverTab[108477]++
													ext := ee.GetExtensions()
													*ext = append(*ext, buf...)
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:263
		// _ = "end of CoverTab[108477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:264
		_go_fuzz_dep_.CoverTab[108478]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:264
		// _ = "end of CoverTab[108478]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:264
	// _ = "end of CoverTab[108475]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:264
	_go_fuzz_dep_.CoverTab[108476]++
												if ee, eok := e.(extendableProto); eok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:265
		_go_fuzz_dep_.CoverTab[108479]++
													m := ee.extensionsWrite()
													ext := m[int32(tag)]
													ext.enc = append(ext.enc, buf...)
													m[int32(tag)] = ext
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:269
		// _ = "end of CoverTab[108479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:270
		_go_fuzz_dep_.CoverTab[108480]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:270
		// _ = "end of CoverTab[108480]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:270
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:270
	// _ = "end of CoverTab[108476]"
}

func encodeExtension(extension *ExtensionDesc, value interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:273
	_go_fuzz_dep_.CoverTab[108481]++
												u := getMarshalInfo(reflect.TypeOf(extension.ExtendedType))
												ei := u.getExtElemInfo(extension)
												v := value
												p := toAddrPointer(&v, ei.isptr)
												siz := ei.sizer(p, SizeVarint(ei.wiretag))
												buf := make([]byte, 0, siz)
												return ei.marshaler(buf, p, ei.wiretag, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:280
	// _ = "end of CoverTab[108481]"
}

func decodeExtensionFromBytes(extension *ExtensionDesc, buf []byte) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:283
	_go_fuzz_dep_.CoverTab[108482]++
												o := 0
												for o < len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:285
		_go_fuzz_dep_.CoverTab[108484]++
													tag, n := DecodeVarint((buf)[o:])
													fieldNum := int32(tag >> 3)
													wireType := int(tag & 0x7)
													if o+n > len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:289
			_go_fuzz_dep_.CoverTab[108488]++
														return nil, fmt.Errorf("unable to decode extension")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:290
			// _ = "end of CoverTab[108488]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:291
			_go_fuzz_dep_.CoverTab[108489]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:291
			// _ = "end of CoverTab[108489]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:291
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:291
		// _ = "end of CoverTab[108484]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:291
		_go_fuzz_dep_.CoverTab[108485]++
													l, err := size((buf)[o+n:], wireType)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:293
			_go_fuzz_dep_.CoverTab[108490]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:294
			// _ = "end of CoverTab[108490]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:295
			_go_fuzz_dep_.CoverTab[108491]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:295
			// _ = "end of CoverTab[108491]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:295
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:295
		// _ = "end of CoverTab[108485]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:295
		_go_fuzz_dep_.CoverTab[108486]++
													if int32(fieldNum) == extension.Field {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:296
			_go_fuzz_dep_.CoverTab[108492]++
														if o+n+l > len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:297
				_go_fuzz_dep_.CoverTab[108495]++
															return nil, fmt.Errorf("unable to decode extension")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:298
				// _ = "end of CoverTab[108495]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:299
				_go_fuzz_dep_.CoverTab[108496]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:299
				// _ = "end of CoverTab[108496]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:299
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:299
			// _ = "end of CoverTab[108492]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:299
			_go_fuzz_dep_.CoverTab[108493]++
														v, err := decodeExtension((buf)[o:o+n+l], extension)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:301
				_go_fuzz_dep_.CoverTab[108497]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:302
				// _ = "end of CoverTab[108497]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:303
				_go_fuzz_dep_.CoverTab[108498]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:303
				// _ = "end of CoverTab[108498]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:303
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:303
			// _ = "end of CoverTab[108493]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:303
			_go_fuzz_dep_.CoverTab[108494]++
														return v, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:304
			// _ = "end of CoverTab[108494]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:305
			_go_fuzz_dep_.CoverTab[108499]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:305
			// _ = "end of CoverTab[108499]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:305
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:305
		// _ = "end of CoverTab[108486]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:305
		_go_fuzz_dep_.CoverTab[108487]++
													o += n + l
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:306
		// _ = "end of CoverTab[108487]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:307
	// _ = "end of CoverTab[108482]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:307
	_go_fuzz_dep_.CoverTab[108483]++
												return defaultExtensionValue(extension)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:308
	// _ = "end of CoverTab[108483]"
}

func (this *Extension) Encode() error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:311
	_go_fuzz_dep_.CoverTab[108500]++
												if this.enc == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:312
		_go_fuzz_dep_.CoverTab[108502]++
													var err error
													this.enc, err = encodeExtension(this.desc, this.value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:315
			_go_fuzz_dep_.CoverTab[108503]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:316
			// _ = "end of CoverTab[108503]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:317
			_go_fuzz_dep_.CoverTab[108504]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:317
			// _ = "end of CoverTab[108504]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:317
		// _ = "end of CoverTab[108502]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:318
		_go_fuzz_dep_.CoverTab[108505]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:318
		// _ = "end of CoverTab[108505]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:318
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:318
	// _ = "end of CoverTab[108500]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:318
	_go_fuzz_dep_.CoverTab[108501]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:319
	// _ = "end of CoverTab[108501]"
}

func (this Extension) GoString() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:322
	_go_fuzz_dep_.CoverTab[108506]++
												if err := this.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:323
		_go_fuzz_dep_.CoverTab[108508]++
													return fmt.Sprintf("error encoding extension: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:324
		// _ = "end of CoverTab[108508]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:325
		_go_fuzz_dep_.CoverTab[108509]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:325
		// _ = "end of CoverTab[108509]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:325
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:325
	// _ = "end of CoverTab[108506]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:325
	_go_fuzz_dep_.CoverTab[108507]++
												return fmt.Sprintf("proto.NewExtension(%#v)", this.enc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:326
	// _ = "end of CoverTab[108507]"
}

func SetUnsafeExtension(pb Message, fieldNum int32, value interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:329
	_go_fuzz_dep_.CoverTab[108510]++
												typ := reflect.TypeOf(pb).Elem()
												ext, ok := extensionMaps[typ]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:332
		_go_fuzz_dep_.CoverTab[108513]++
													return fmt.Errorf("proto: bad extended type; %s is not extendable", typ.String())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:333
		// _ = "end of CoverTab[108513]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:334
		_go_fuzz_dep_.CoverTab[108514]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:334
		// _ = "end of CoverTab[108514]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:334
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:334
	// _ = "end of CoverTab[108510]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:334
	_go_fuzz_dep_.CoverTab[108511]++
												desc, ok := ext[fieldNum]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:336
		_go_fuzz_dep_.CoverTab[108515]++
													return errors.New("proto: bad extension number; not in declared ranges")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:337
		// _ = "end of CoverTab[108515]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:338
		_go_fuzz_dep_.CoverTab[108516]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:338
		// _ = "end of CoverTab[108516]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:338
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:338
	// _ = "end of CoverTab[108511]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:338
	_go_fuzz_dep_.CoverTab[108512]++
												return SetExtension(pb, desc, value)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:339
	// _ = "end of CoverTab[108512]"
}

func GetUnsafeExtension(pb Message, fieldNum int32) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:342
	_go_fuzz_dep_.CoverTab[108517]++
												typ := reflect.TypeOf(pb).Elem()
												ext, ok := extensionMaps[typ]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:345
		_go_fuzz_dep_.CoverTab[108520]++
													return nil, fmt.Errorf("proto: bad extended type; %s is not extendable", typ.String())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:346
		// _ = "end of CoverTab[108520]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:347
		_go_fuzz_dep_.CoverTab[108521]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:347
		// _ = "end of CoverTab[108521]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:347
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:347
	// _ = "end of CoverTab[108517]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:347
	_go_fuzz_dep_.CoverTab[108518]++
												desc, ok := ext[fieldNum]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:349
		_go_fuzz_dep_.CoverTab[108522]++
													return nil, fmt.Errorf("unregistered field number %d", fieldNum)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:350
		// _ = "end of CoverTab[108522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:351
		_go_fuzz_dep_.CoverTab[108523]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:351
		// _ = "end of CoverTab[108523]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:351
	// _ = "end of CoverTab[108518]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:351
	_go_fuzz_dep_.CoverTab[108519]++
												return GetExtension(pb, desc)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:352
	// _ = "end of CoverTab[108519]"
}

func NewUnsafeXXX_InternalExtensions(m map[int32]Extension) XXX_InternalExtensions {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:355
	_go_fuzz_dep_.CoverTab[108524]++
												x := &XXX_InternalExtensions{
		p: new(struct {
			mu		sync.Mutex
			extensionMap	map[int32]Extension
		}),
	}
												x.p.extensionMap = m
												return *x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:363
	// _ = "end of CoverTab[108524]"
}

func GetUnsafeExtensionsMap(extendable Message) map[int32]Extension {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:366
	_go_fuzz_dep_.CoverTab[108525]++
												pb := extendable.(extendableProto)
												return pb.extensionsWrite()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:368
	// _ = "end of CoverTab[108525]"
}

func deleteExtension(pb extensionsBytes, theFieldNum int32, offset int) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:371
	_go_fuzz_dep_.CoverTab[108526]++
												ext := pb.GetExtensions()
												for offset < len(*ext) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:373
		_go_fuzz_dep_.CoverTab[108528]++
													tag, n1 := DecodeVarint((*ext)[offset:])
													fieldNum := int32(tag >> 3)
													wireType := int(tag & 0x7)
													n2, err := size((*ext)[offset+n1:], wireType)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:378
			_go_fuzz_dep_.CoverTab[108531]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:379
			// _ = "end of CoverTab[108531]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:380
			_go_fuzz_dep_.CoverTab[108532]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:380
			// _ = "end of CoverTab[108532]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:380
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:380
		// _ = "end of CoverTab[108528]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:380
		_go_fuzz_dep_.CoverTab[108529]++
													newOffset := offset + n1 + n2
													if fieldNum == theFieldNum {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:382
			_go_fuzz_dep_.CoverTab[108533]++
														*ext = append((*ext)[:offset], (*ext)[newOffset:]...)
														return offset
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:384
			// _ = "end of CoverTab[108533]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:385
			_go_fuzz_dep_.CoverTab[108534]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:385
			// _ = "end of CoverTab[108534]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:385
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:385
		// _ = "end of CoverTab[108529]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:385
		_go_fuzz_dep_.CoverTab[108530]++
													offset = newOffset
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:386
		// _ = "end of CoverTab[108530]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:387
	// _ = "end of CoverTab[108526]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:387
	_go_fuzz_dep_.CoverTab[108527]++
												return -1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:388
	// _ = "end of CoverTab[108527]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:389
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/extensions_gogo.go:389
var _ = _go_fuzz_dep_.CoverTab
