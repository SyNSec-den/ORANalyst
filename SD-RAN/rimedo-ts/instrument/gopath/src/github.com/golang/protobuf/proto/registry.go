// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:5
)

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// filePath is the path to the proto source file.
type filePath = string	// e.g., "google/protobuf/descriptor.proto"

// fileDescGZIP is the compressed contents of the encoded FileDescriptorProto.
type fileDescGZIP = []byte

var fileCache sync.Map	// map[filePath]fileDescGZIP

// RegisterFile is called from generated code to register the compressed
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:30
// FileDescriptorProto with the file path for a proto source file.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:30
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:30
// Deprecated: Use protoregistry.GlobalFiles.RegisterFile instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:34
func RegisterFile(s filePath, d fileDescGZIP) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:34
	_go_fuzz_dep_.CoverTab[61556]++

												zr, err := gzip.NewReader(bytes.NewReader(d))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:37
		_go_fuzz_dep_.CoverTab[61559]++
													panic(fmt.Sprintf("proto: invalid compressed file descriptor: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:38
		// _ = "end of CoverTab[61559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:39
		_go_fuzz_dep_.CoverTab[61560]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:39
		// _ = "end of CoverTab[61560]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:39
	// _ = "end of CoverTab[61556]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:39
	_go_fuzz_dep_.CoverTab[61557]++
												b, err := ioutil.ReadAll(zr)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:41
		_go_fuzz_dep_.CoverTab[61561]++
													panic(fmt.Sprintf("proto: invalid compressed file descriptor: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:42
		// _ = "end of CoverTab[61561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:43
		_go_fuzz_dep_.CoverTab[61562]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:43
		// _ = "end of CoverTab[61562]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:43
	// _ = "end of CoverTab[61557]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:43
	_go_fuzz_dep_.CoverTab[61558]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:48
	protoimpl.DescBuilder{RawDescriptor: b}.Build()

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:51
	fileCache.Store(s, d)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:51
	// _ = "end of CoverTab[61558]"
}

// FileDescriptor returns the compressed FileDescriptorProto given the file path
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:54
// for a proto source file. It returns nil if not found.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:54
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:54
// Deprecated: Use protoregistry.GlobalFiles.FindFileByPath instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:58
func FileDescriptor(s filePath) fileDescGZIP {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:58
	_go_fuzz_dep_.CoverTab[61563]++
												if v, ok := fileCache.Load(s); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:59
		_go_fuzz_dep_.CoverTab[61567]++
													return v.(fileDescGZIP)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:60
		// _ = "end of CoverTab[61567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:61
		_go_fuzz_dep_.CoverTab[61568]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:61
		// _ = "end of CoverTab[61568]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:61
	// _ = "end of CoverTab[61563]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:61
	_go_fuzz_dep_.CoverTab[61564]++

	// Find the descriptor in the v2 registry.
	var b []byte
	if fd, _ := protoregistry.GlobalFiles.FindFileByPath(s); fd != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:65
		_go_fuzz_dep_.CoverTab[61569]++
													b, _ = Marshal(protodesc.ToFileDescriptorProto(fd))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:66
		// _ = "end of CoverTab[61569]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:67
		_go_fuzz_dep_.CoverTab[61570]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:67
		// _ = "end of CoverTab[61570]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:67
	// _ = "end of CoverTab[61564]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:67
	_go_fuzz_dep_.CoverTab[61565]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:70
	if len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:70
		_go_fuzz_dep_.CoverTab[61571]++
													v, _ := fileCache.LoadOrStore(s, protoimpl.X.CompressGZIP(b))
													return v.(fileDescGZIP)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:72
		// _ = "end of CoverTab[61571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:73
		_go_fuzz_dep_.CoverTab[61572]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:73
		// _ = "end of CoverTab[61572]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:73
	// _ = "end of CoverTab[61565]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:73
	_go_fuzz_dep_.CoverTab[61566]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:74
	// _ = "end of CoverTab[61566]"
}

// enumName is the name of an enum. For historical reasons, the enum name is
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:77
// neither the full Go name nor the full protobuf name of the enum.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:77
// The name is the dot-separated combination of just the proto package that the
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:77
// enum is declared within followed by the Go type name of the generated enum.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:81
type enumName = string	// e.g., "my.proto.package.GoMessage_GoEnum"

// enumsByName maps enum values by name to their numeric counterpart.
type enumsByName = map[string]int32

// enumsByNumber maps enum values by number to their name counterpart.
type enumsByNumber = map[int32]string

var enumCache sync.Map		// map[enumName]enumsByName
var numFilesCache sync.Map	// map[protoreflect.FullName]int

// RegisterEnum is called from the generated code to register the mapping of
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:92
// enum value names to enum numbers for the enum identified by s.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:92
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:92
// Deprecated: Use protoregistry.GlobalTypes.RegisterEnum instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:96
func RegisterEnum(s enumName, _ enumsByNumber, m enumsByName) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:96
	_go_fuzz_dep_.CoverTab[61573]++
												if _, ok := enumCache.Load(s); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:97
		_go_fuzz_dep_.CoverTab[61575]++
													panic("proto: duplicate enum registered: " + s)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:98
		// _ = "end of CoverTab[61575]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:99
		_go_fuzz_dep_.CoverTab[61576]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:99
		// _ = "end of CoverTab[61576]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:99
	// _ = "end of CoverTab[61573]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:99
	_go_fuzz_dep_.CoverTab[61574]++
												enumCache.Store(s, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:100
	// _ = "end of CoverTab[61574]"

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:104
}

// EnumValueMap returns the mapping from enum value names to enum numbers for
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:106
// the enum of the given name. It returns nil if not found.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:106
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:106
// Deprecated: Use protoregistry.GlobalTypes.FindEnumByName instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:110
func EnumValueMap(s enumName) enumsByName {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:110
	_go_fuzz_dep_.CoverTab[61577]++
												if v, ok := enumCache.Load(s); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:111
		_go_fuzz_dep_.CoverTab[61583]++
													return v.(enumsByName)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:112
		// _ = "end of CoverTab[61583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:113
		_go_fuzz_dep_.CoverTab[61584]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:113
		// _ = "end of CoverTab[61584]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:113
	// _ = "end of CoverTab[61577]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:113
	_go_fuzz_dep_.CoverTab[61578]++

	// Check whether the cache is stale. If the number of files in the current
	// package differs, then it means that some enums may have been recently
	// registered upstream that we do not know about.
	var protoPkg protoreflect.FullName
	if i := strings.LastIndexByte(s, '.'); i >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:119
		_go_fuzz_dep_.CoverTab[61585]++
													protoPkg = protoreflect.FullName(s[:i])
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:120
		// _ = "end of CoverTab[61585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:121
		_go_fuzz_dep_.CoverTab[61586]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:121
		// _ = "end of CoverTab[61586]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:121
	// _ = "end of CoverTab[61578]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:121
	_go_fuzz_dep_.CoverTab[61579]++
												v, _ := numFilesCache.Load(protoPkg)
												numFiles, _ := v.(int)
												if protoregistry.GlobalFiles.NumFilesByPackage(protoPkg) == numFiles {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:124
		_go_fuzz_dep_.CoverTab[61587]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:125
		// _ = "end of CoverTab[61587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:126
		_go_fuzz_dep_.CoverTab[61588]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:126
		// _ = "end of CoverTab[61588]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:126
	// _ = "end of CoverTab[61579]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:126
	_go_fuzz_dep_.CoverTab[61580]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:129
	numFiles = 0
	protoregistry.GlobalFiles.RangeFilesByPackage(protoPkg, func(fd protoreflect.FileDescriptor) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:130
		_go_fuzz_dep_.CoverTab[61589]++
													walkEnums(fd, func(ed protoreflect.EnumDescriptor) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:131
			_go_fuzz_dep_.CoverTab[61591]++
														name := protoimpl.X.LegacyEnumName(ed)
														if _, ok := enumCache.Load(name); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:133
				_go_fuzz_dep_.CoverTab[61592]++
															m := make(enumsByName)
															evs := ed.Values()
															for i := evs.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:136
					_go_fuzz_dep_.CoverTab[61594]++
																ev := evs.Get(i)
																m[string(ev.Name())] = int32(ev.Number())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:138
					// _ = "end of CoverTab[61594]"
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:139
				// _ = "end of CoverTab[61592]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:139
				_go_fuzz_dep_.CoverTab[61593]++
															enumCache.LoadOrStore(name, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:140
				// _ = "end of CoverTab[61593]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:141
				_go_fuzz_dep_.CoverTab[61595]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:141
				// _ = "end of CoverTab[61595]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:141
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:141
			// _ = "end of CoverTab[61591]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:142
		// _ = "end of CoverTab[61589]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:142
		_go_fuzz_dep_.CoverTab[61590]++
													numFiles++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:144
		// _ = "end of CoverTab[61590]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:145
	// _ = "end of CoverTab[61580]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:145
	_go_fuzz_dep_.CoverTab[61581]++
												numFilesCache.Store(protoPkg, numFiles)

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:149
	if v, ok := enumCache.Load(s); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:149
		_go_fuzz_dep_.CoverTab[61596]++
													return v.(enumsByName)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:150
		// _ = "end of CoverTab[61596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:151
		_go_fuzz_dep_.CoverTab[61597]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:151
		// _ = "end of CoverTab[61597]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:151
	// _ = "end of CoverTab[61581]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:151
	_go_fuzz_dep_.CoverTab[61582]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:152
	// _ = "end of CoverTab[61582]"
}

// walkEnums recursively walks all enums declared in d.
func walkEnums(d interface {
	Enums() protoreflect.EnumDescriptors
	Messages() protoreflect.MessageDescriptors
}, f func(protoreflect.EnumDescriptor)) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:159
	_go_fuzz_dep_.CoverTab[61598]++
												eds := d.Enums()
												for i := eds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:161
		_go_fuzz_dep_.CoverTab[61600]++
													f(eds.Get(i))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:162
		// _ = "end of CoverTab[61600]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:163
	// _ = "end of CoverTab[61598]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:163
	_go_fuzz_dep_.CoverTab[61599]++
												mds := d.Messages()
												for i := mds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:165
		_go_fuzz_dep_.CoverTab[61601]++
													walkEnums(mds.Get(i), f)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:166
		// _ = "end of CoverTab[61601]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:167
	// _ = "end of CoverTab[61599]"
}

// messageName is the full name of protobuf message.
type messageName = string

var messageTypeCache sync.Map	// map[messageName]reflect.Type

// RegisterType is called from generated code to register the message Go type
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:175
// for a message of the given name.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:175
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:175
// Deprecated: Use protoregistry.GlobalTypes.RegisterMessage instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:179
func RegisterType(m Message, s messageName) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:179
	_go_fuzz_dep_.CoverTab[61602]++
												mt := protoimpl.X.LegacyMessageTypeOf(m, protoreflect.FullName(s))
												if err := protoregistry.GlobalTypes.RegisterMessage(mt); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:181
		_go_fuzz_dep_.CoverTab[61604]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:182
		// _ = "end of CoverTab[61604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:183
		_go_fuzz_dep_.CoverTab[61605]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:183
		// _ = "end of CoverTab[61605]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:183
	// _ = "end of CoverTab[61602]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:183
	_go_fuzz_dep_.CoverTab[61603]++
												messageTypeCache.Store(s, reflect.TypeOf(m))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:184
	// _ = "end of CoverTab[61603]"
}

// RegisterMapType is called from generated code to register the Go map type
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:187
// for a protobuf message representing a map entry.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:187
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:187
// Deprecated: Do not use.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:191
func RegisterMapType(m interface{}, s messageName) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:191
	_go_fuzz_dep_.CoverTab[61606]++
												t := reflect.TypeOf(m)
												if t.Kind() != reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:193
		_go_fuzz_dep_.CoverTab[61609]++
													panic(fmt.Sprintf("invalid map kind: %v", t))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:194
		// _ = "end of CoverTab[61609]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:195
		_go_fuzz_dep_.CoverTab[61610]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:195
		// _ = "end of CoverTab[61610]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:195
	// _ = "end of CoverTab[61606]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:195
	_go_fuzz_dep_.CoverTab[61607]++
												if _, ok := messageTypeCache.Load(s); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:196
		_go_fuzz_dep_.CoverTab[61611]++
													panic(fmt.Errorf("proto: duplicate proto message registered: %s", s))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:197
		// _ = "end of CoverTab[61611]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:198
		_go_fuzz_dep_.CoverTab[61612]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:198
		// _ = "end of CoverTab[61612]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:198
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:198
	// _ = "end of CoverTab[61607]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:198
	_go_fuzz_dep_.CoverTab[61608]++
												messageTypeCache.Store(s, t)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:199
	// _ = "end of CoverTab[61608]"
}

// MessageType returns the message type for a named message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:202
// It returns nil if not found.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:202
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:202
// Deprecated: Use protoregistry.GlobalTypes.FindMessageByName instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:206
func MessageType(s messageName) reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:206
	_go_fuzz_dep_.CoverTab[61613]++
												if v, ok := messageTypeCache.Load(s); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:207
		_go_fuzz_dep_.CoverTab[61618]++
													return v.(reflect.Type)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:208
		// _ = "end of CoverTab[61618]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:209
		_go_fuzz_dep_.CoverTab[61619]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:209
		// _ = "end of CoverTab[61619]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:209
	// _ = "end of CoverTab[61613]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:209
	_go_fuzz_dep_.CoverTab[61614]++

	// Derive the message type from the v2 registry.
	var t reflect.Type
	if mt, _ := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(s)); mt != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:213
		_go_fuzz_dep_.CoverTab[61620]++
													t = messageGoType(mt)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:214
		// _ = "end of CoverTab[61620]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:215
		_go_fuzz_dep_.CoverTab[61621]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:215
		// _ = "end of CoverTab[61621]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:215
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:215
	// _ = "end of CoverTab[61614]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:215
	_go_fuzz_dep_.CoverTab[61615]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:219
	if t == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:219
		_go_fuzz_dep_.CoverTab[61622]++
													d, _ := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(s))
													if md, _ := d.(protoreflect.MessageDescriptor); md != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:221
			_go_fuzz_dep_.CoverTab[61623]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:221
			return md.IsMapEntry()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:221
			// _ = "end of CoverTab[61623]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:221
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:221
			_go_fuzz_dep_.CoverTab[61624]++
														kt := goTypeForField(md.Fields().ByNumber(1))
														vt := goTypeForField(md.Fields().ByNumber(2))
														t = reflect.MapOf(kt, vt)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:224
			// _ = "end of CoverTab[61624]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:225
			_go_fuzz_dep_.CoverTab[61625]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:225
			// _ = "end of CoverTab[61625]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:225
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:225
		// _ = "end of CoverTab[61622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:226
		_go_fuzz_dep_.CoverTab[61626]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:226
		// _ = "end of CoverTab[61626]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:226
	// _ = "end of CoverTab[61615]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:226
	_go_fuzz_dep_.CoverTab[61616]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:229
	if t != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:229
		_go_fuzz_dep_.CoverTab[61627]++
													v, _ := messageTypeCache.LoadOrStore(s, t)
													return v.(reflect.Type)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:231
		// _ = "end of CoverTab[61627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:232
		_go_fuzz_dep_.CoverTab[61628]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:232
		// _ = "end of CoverTab[61628]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:232
	// _ = "end of CoverTab[61616]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:232
	_go_fuzz_dep_.CoverTab[61617]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:233
	// _ = "end of CoverTab[61617]"
}

func goTypeForField(fd protoreflect.FieldDescriptor) reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:236
	_go_fuzz_dep_.CoverTab[61629]++
												switch k := fd.Kind(); k {
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:238
		_go_fuzz_dep_.CoverTab[61630]++
													if et, _ := protoregistry.GlobalTypes.FindEnumByName(fd.Enum().FullName()); et != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:239
			_go_fuzz_dep_.CoverTab[61635]++
														return enumGoType(et)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:240
			// _ = "end of CoverTab[61635]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:241
			_go_fuzz_dep_.CoverTab[61636]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:241
			// _ = "end of CoverTab[61636]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:241
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:241
		// _ = "end of CoverTab[61630]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:241
		_go_fuzz_dep_.CoverTab[61631]++
													return reflect.TypeOf(protoreflect.EnumNumber(0))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:242
		// _ = "end of CoverTab[61631]"
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:243
		_go_fuzz_dep_.CoverTab[61632]++
													if mt, _ := protoregistry.GlobalTypes.FindMessageByName(fd.Message().FullName()); mt != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:244
			_go_fuzz_dep_.CoverTab[61637]++
														return messageGoType(mt)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:245
			// _ = "end of CoverTab[61637]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:246
			_go_fuzz_dep_.CoverTab[61638]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:246
			// _ = "end of CoverTab[61638]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:246
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:246
		// _ = "end of CoverTab[61632]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:246
		_go_fuzz_dep_.CoverTab[61633]++
													return reflect.TypeOf((*protoreflect.Message)(nil)).Elem()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:247
		// _ = "end of CoverTab[61633]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:248
		_go_fuzz_dep_.CoverTab[61634]++
													return reflect.TypeOf(fd.Default().Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:249
		// _ = "end of CoverTab[61634]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:250
	// _ = "end of CoverTab[61629]"
}

func enumGoType(et protoreflect.EnumType) reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:253
	_go_fuzz_dep_.CoverTab[61639]++
												return reflect.TypeOf(et.New(0))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:254
	// _ = "end of CoverTab[61639]"
}

func messageGoType(mt protoreflect.MessageType) reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:257
	_go_fuzz_dep_.CoverTab[61640]++
												return reflect.TypeOf(MessageV1(mt.Zero().Interface()))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:258
	// _ = "end of CoverTab[61640]"
}

// MessageName returns the full protobuf name for the given message type.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:261
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:261
// Deprecated: Use protoreflect.MessageDescriptor.FullName instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:264
func MessageName(m Message) messageName {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:264
	_go_fuzz_dep_.CoverTab[61641]++
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:265
		_go_fuzz_dep_.CoverTab[61644]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:266
		// _ = "end of CoverTab[61644]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:267
		_go_fuzz_dep_.CoverTab[61645]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:267
		// _ = "end of CoverTab[61645]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:267
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:267
	// _ = "end of CoverTab[61641]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:267
	_go_fuzz_dep_.CoverTab[61642]++
												if m, ok := m.(interface{ XXX_MessageName() messageName }); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:268
		_go_fuzz_dep_.CoverTab[61646]++
													return m.XXX_MessageName()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:269
		// _ = "end of CoverTab[61646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:270
		_go_fuzz_dep_.CoverTab[61647]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:270
		// _ = "end of CoverTab[61647]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:270
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:270
	// _ = "end of CoverTab[61642]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:270
	_go_fuzz_dep_.CoverTab[61643]++
												return messageName(protoimpl.X.MessageDescriptorOf(m).FullName())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:271
	// _ = "end of CoverTab[61643]"
}

// RegisterExtension is called from the generated code to register
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:274
// the extension descriptor.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:274
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:274
// Deprecated: Use protoregistry.GlobalTypes.RegisterExtension instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:278
func RegisterExtension(d *ExtensionDesc) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:278
	_go_fuzz_dep_.CoverTab[61648]++
												if err := protoregistry.GlobalTypes.RegisterExtension(d); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:279
		_go_fuzz_dep_.CoverTab[61649]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:280
		// _ = "end of CoverTab[61649]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:281
		_go_fuzz_dep_.CoverTab[61650]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:281
		// _ = "end of CoverTab[61650]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:281
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:281
	// _ = "end of CoverTab[61648]"
}

type extensionsByNumber = map[int32]*ExtensionDesc

var extensionCache sync.Map	// map[messageName]extensionsByNumber

// RegisteredExtensions returns a map of the registered extensions for the
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:288
// provided protobuf message, indexed by the extension field number.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:288
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:288
// Deprecated: Use protoregistry.GlobalTypes.RangeExtensionsByMessage instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:292
func RegisteredExtensions(m Message) extensionsByNumber {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:292
	_go_fuzz_dep_.CoverTab[61651]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:296
	s := MessageName(m)
	v, _ := extensionCache.Load(s)
	xs, _ := v.(extensionsByNumber)
	if protoregistry.GlobalTypes.NumExtensionsByMessage(protoreflect.FullName(s)) == len(xs) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:299
		_go_fuzz_dep_.CoverTab[61654]++
													return xs
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:300
		// _ = "end of CoverTab[61654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:301
		_go_fuzz_dep_.CoverTab[61655]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:301
		// _ = "end of CoverTab[61655]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:301
	// _ = "end of CoverTab[61651]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:301
	_go_fuzz_dep_.CoverTab[61652]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:304
	xs = make(extensionsByNumber)
	protoregistry.GlobalTypes.RangeExtensionsByMessage(protoreflect.FullName(s), func(xt protoreflect.ExtensionType) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:305
		_go_fuzz_dep_.CoverTab[61656]++
													if xd, ok := xt.(*ExtensionDesc); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:306
			_go_fuzz_dep_.CoverTab[61658]++
														xs[int32(xt.TypeDescriptor().Number())] = xd
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:307
			// _ = "end of CoverTab[61658]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:308
			_go_fuzz_dep_.CoverTab[61659]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:308
			// _ = "end of CoverTab[61659]"

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:312
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:312
		// _ = "end of CoverTab[61656]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:312
		_go_fuzz_dep_.CoverTab[61657]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:313
		// _ = "end of CoverTab[61657]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:314
	// _ = "end of CoverTab[61652]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:314
	_go_fuzz_dep_.CoverTab[61653]++
												extensionCache.Store(s, xs)
												return xs
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:316
	// _ = "end of CoverTab[61653]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:317
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/registry.go:317
var _ = _go_fuzz_dep_.CoverTab
