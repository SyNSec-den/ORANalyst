// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
package protodesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:5
)

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/internal/encoding/defval"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/types/descriptorpb"
)

// ToFileDescriptorProto copies a protoreflect.FileDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:19
// google.protobuf.FileDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:21
func ToFileDescriptorProto(file protoreflect.FileDescriptor) *descriptorpb.FileDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:21
	_go_fuzz_dep_.CoverTab[60998]++
													p := &descriptorpb.FileDescriptorProto{
		Name:		proto.String(file.Path()),
		Options:	proto.Clone(file.Options()).(*descriptorpb.FileOptions),
	}
	if file.Package() != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:26
		_go_fuzz_dep_.CoverTab[61007]++
														p.Package = proto.String(string(file.Package()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:27
		// _ = "end of CoverTab[61007]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:28
		_go_fuzz_dep_.CoverTab[61008]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:28
		// _ = "end of CoverTab[61008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:28
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:28
	// _ = "end of CoverTab[60998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:28
	_go_fuzz_dep_.CoverTab[60999]++
													for i, imports := 0, file.Imports(); i < imports.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:29
		_go_fuzz_dep_.CoverTab[61009]++
														imp := imports.Get(i)
														p.Dependency = append(p.Dependency, imp.Path())
														if imp.IsPublic {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:32
			_go_fuzz_dep_.CoverTab[61011]++
															p.PublicDependency = append(p.PublicDependency, int32(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:33
			// _ = "end of CoverTab[61011]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:34
			_go_fuzz_dep_.CoverTab[61012]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:34
			// _ = "end of CoverTab[61012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:34
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:34
		// _ = "end of CoverTab[61009]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:34
		_go_fuzz_dep_.CoverTab[61010]++
														if imp.IsWeak {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:35
			_go_fuzz_dep_.CoverTab[61013]++
															p.WeakDependency = append(p.WeakDependency, int32(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:36
			// _ = "end of CoverTab[61013]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:37
			_go_fuzz_dep_.CoverTab[61014]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:37
			// _ = "end of CoverTab[61014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:37
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:37
		// _ = "end of CoverTab[61010]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:38
	// _ = "end of CoverTab[60999]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:38
	_go_fuzz_dep_.CoverTab[61000]++
													for i, locs := 0, file.SourceLocations(); i < locs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:39
		_go_fuzz_dep_.CoverTab[61015]++
														loc := locs.Get(i)
														l := &descriptorpb.SourceCodeInfo_Location{}
														l.Path = append(l.Path, loc.Path...)
														if loc.StartLine == loc.EndLine {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:43
			_go_fuzz_dep_.CoverTab[61020]++
															l.Span = []int32{int32(loc.StartLine), int32(loc.StartColumn), int32(loc.EndColumn)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:44
			// _ = "end of CoverTab[61020]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:45
			_go_fuzz_dep_.CoverTab[61021]++
															l.Span = []int32{int32(loc.StartLine), int32(loc.StartColumn), int32(loc.EndLine), int32(loc.EndColumn)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:46
			// _ = "end of CoverTab[61021]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:47
		// _ = "end of CoverTab[61015]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:47
		_go_fuzz_dep_.CoverTab[61016]++
														l.LeadingDetachedComments = append([]string(nil), loc.LeadingDetachedComments...)
														if loc.LeadingComments != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:49
			_go_fuzz_dep_.CoverTab[61022]++
															l.LeadingComments = proto.String(loc.LeadingComments)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:50
			// _ = "end of CoverTab[61022]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:51
			_go_fuzz_dep_.CoverTab[61023]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:51
			// _ = "end of CoverTab[61023]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:51
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:51
		// _ = "end of CoverTab[61016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:51
		_go_fuzz_dep_.CoverTab[61017]++
														if loc.TrailingComments != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:52
			_go_fuzz_dep_.CoverTab[61024]++
															l.TrailingComments = proto.String(loc.TrailingComments)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:53
			// _ = "end of CoverTab[61024]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:54
			_go_fuzz_dep_.CoverTab[61025]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:54
			// _ = "end of CoverTab[61025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:54
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:54
		// _ = "end of CoverTab[61017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:54
		_go_fuzz_dep_.CoverTab[61018]++
														if p.SourceCodeInfo == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:55
			_go_fuzz_dep_.CoverTab[61026]++
															p.SourceCodeInfo = &descriptorpb.SourceCodeInfo{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:56
			// _ = "end of CoverTab[61026]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:57
			_go_fuzz_dep_.CoverTab[61027]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:57
			// _ = "end of CoverTab[61027]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:57
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:57
		// _ = "end of CoverTab[61018]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:57
		_go_fuzz_dep_.CoverTab[61019]++
														p.SourceCodeInfo.Location = append(p.SourceCodeInfo.Location, l)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:58
		// _ = "end of CoverTab[61019]"

	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:60
	// _ = "end of CoverTab[61000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:60
	_go_fuzz_dep_.CoverTab[61001]++
													for i, messages := 0, file.Messages(); i < messages.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:61
		_go_fuzz_dep_.CoverTab[61028]++
														p.MessageType = append(p.MessageType, ToDescriptorProto(messages.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:62
		// _ = "end of CoverTab[61028]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:63
	// _ = "end of CoverTab[61001]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:63
	_go_fuzz_dep_.CoverTab[61002]++
													for i, enums := 0, file.Enums(); i < enums.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:64
		_go_fuzz_dep_.CoverTab[61029]++
														p.EnumType = append(p.EnumType, ToEnumDescriptorProto(enums.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:65
		// _ = "end of CoverTab[61029]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:66
	// _ = "end of CoverTab[61002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:66
	_go_fuzz_dep_.CoverTab[61003]++
													for i, services := 0, file.Services(); i < services.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:67
		_go_fuzz_dep_.CoverTab[61030]++
														p.Service = append(p.Service, ToServiceDescriptorProto(services.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:68
		// _ = "end of CoverTab[61030]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:69
	// _ = "end of CoverTab[61003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:69
	_go_fuzz_dep_.CoverTab[61004]++
													for i, exts := 0, file.Extensions(); i < exts.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:70
		_go_fuzz_dep_.CoverTab[61031]++
														p.Extension = append(p.Extension, ToFieldDescriptorProto(exts.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:71
		// _ = "end of CoverTab[61031]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:72
	// _ = "end of CoverTab[61004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:72
	_go_fuzz_dep_.CoverTab[61005]++
													if syntax := file.Syntax(); syntax != protoreflect.Proto2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:73
		_go_fuzz_dep_.CoverTab[61032]++
														p.Syntax = proto.String(file.Syntax().String())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:74
		// _ = "end of CoverTab[61032]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:75
		_go_fuzz_dep_.CoverTab[61033]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:75
		// _ = "end of CoverTab[61033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:75
	// _ = "end of CoverTab[61005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:75
	_go_fuzz_dep_.CoverTab[61006]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:76
	// _ = "end of CoverTab[61006]"
}

// ToDescriptorProto copies a protoreflect.MessageDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:79
// google.protobuf.DescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:81
func ToDescriptorProto(message protoreflect.MessageDescriptor) *descriptorpb.DescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:81
	_go_fuzz_dep_.CoverTab[61034]++
													p := &descriptorpb.DescriptorProto{
		Name:		proto.String(string(message.Name())),
		Options:	proto.Clone(message.Options()).(*descriptorpb.MessageOptions),
	}
	for i, fields := 0, message.Fields(); i < fields.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:86
		_go_fuzz_dep_.CoverTab[61043]++
														p.Field = append(p.Field, ToFieldDescriptorProto(fields.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:87
		// _ = "end of CoverTab[61043]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:88
	// _ = "end of CoverTab[61034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:88
	_go_fuzz_dep_.CoverTab[61035]++
													for i, exts := 0, message.Extensions(); i < exts.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:89
		_go_fuzz_dep_.CoverTab[61044]++
														p.Extension = append(p.Extension, ToFieldDescriptorProto(exts.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:90
		// _ = "end of CoverTab[61044]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:91
	// _ = "end of CoverTab[61035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:91
	_go_fuzz_dep_.CoverTab[61036]++
													for i, messages := 0, message.Messages(); i < messages.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:92
		_go_fuzz_dep_.CoverTab[61045]++
														p.NestedType = append(p.NestedType, ToDescriptorProto(messages.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:93
		// _ = "end of CoverTab[61045]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:94
	// _ = "end of CoverTab[61036]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:94
	_go_fuzz_dep_.CoverTab[61037]++
													for i, enums := 0, message.Enums(); i < enums.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:95
		_go_fuzz_dep_.CoverTab[61046]++
														p.EnumType = append(p.EnumType, ToEnumDescriptorProto(enums.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:96
		// _ = "end of CoverTab[61046]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:97
	// _ = "end of CoverTab[61037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:97
	_go_fuzz_dep_.CoverTab[61038]++
													for i, xranges := 0, message.ExtensionRanges(); i < xranges.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:98
		_go_fuzz_dep_.CoverTab[61047]++
														xrange := xranges.Get(i)
														p.ExtensionRange = append(p.ExtensionRange, &descriptorpb.DescriptorProto_ExtensionRange{
			Start:		proto.Int32(int32(xrange[0])),
			End:		proto.Int32(int32(xrange[1])),
			Options:	proto.Clone(message.ExtensionRangeOptions(i)).(*descriptorpb.ExtensionRangeOptions),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:104
		// _ = "end of CoverTab[61047]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:105
	// _ = "end of CoverTab[61038]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:105
	_go_fuzz_dep_.CoverTab[61039]++
													for i, oneofs := 0, message.Oneofs(); i < oneofs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:106
		_go_fuzz_dep_.CoverTab[61048]++
														p.OneofDecl = append(p.OneofDecl, ToOneofDescriptorProto(oneofs.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:107
		// _ = "end of CoverTab[61048]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:108
	// _ = "end of CoverTab[61039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:108
	_go_fuzz_dep_.CoverTab[61040]++
													for i, ranges := 0, message.ReservedRanges(); i < ranges.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:109
		_go_fuzz_dep_.CoverTab[61049]++
														rrange := ranges.Get(i)
														p.ReservedRange = append(p.ReservedRange, &descriptorpb.DescriptorProto_ReservedRange{
			Start:	proto.Int32(int32(rrange[0])),
			End:	proto.Int32(int32(rrange[1])),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:114
		// _ = "end of CoverTab[61049]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:115
	// _ = "end of CoverTab[61040]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:115
	_go_fuzz_dep_.CoverTab[61041]++
													for i, names := 0, message.ReservedNames(); i < names.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:116
		_go_fuzz_dep_.CoverTab[61050]++
														p.ReservedName = append(p.ReservedName, string(names.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:117
		// _ = "end of CoverTab[61050]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:118
	// _ = "end of CoverTab[61041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:118
	_go_fuzz_dep_.CoverTab[61042]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:119
	// _ = "end of CoverTab[61042]"
}

// ToFieldDescriptorProto copies a protoreflect.FieldDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:122
// google.protobuf.FieldDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:124
func ToFieldDescriptorProto(field protoreflect.FieldDescriptor) *descriptorpb.FieldDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:124
	_go_fuzz_dep_.CoverTab[61051]++
													p := &descriptorpb.FieldDescriptorProto{
		Name:		proto.String(string(field.Name())),
		Number:		proto.Int32(int32(field.Number())),
		Label:		descriptorpb.FieldDescriptorProto_Label(field.Cardinality()).Enum(),
		Options:	proto.Clone(field.Options()).(*descriptorpb.FieldOptions),
	}
	if field.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:131
		_go_fuzz_dep_.CoverTab[61060]++
														p.Extendee = fullNameOf(field.ContainingMessage())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:132
		// _ = "end of CoverTab[61060]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:133
		_go_fuzz_dep_.CoverTab[61061]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:133
		// _ = "end of CoverTab[61061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:133
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:133
	// _ = "end of CoverTab[61051]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:133
	_go_fuzz_dep_.CoverTab[61052]++
													if field.Kind().IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:134
		_go_fuzz_dep_.CoverTab[61062]++
														p.Type = descriptorpb.FieldDescriptorProto_Type(field.Kind()).Enum()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:135
		// _ = "end of CoverTab[61062]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:136
		_go_fuzz_dep_.CoverTab[61063]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:136
		// _ = "end of CoverTab[61063]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:136
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:136
	// _ = "end of CoverTab[61052]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:136
	_go_fuzz_dep_.CoverTab[61053]++
													if field.Enum() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:137
		_go_fuzz_dep_.CoverTab[61064]++
														p.TypeName = fullNameOf(field.Enum())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:138
		// _ = "end of CoverTab[61064]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:139
		_go_fuzz_dep_.CoverTab[61065]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:139
		// _ = "end of CoverTab[61065]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:139
	// _ = "end of CoverTab[61053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:139
	_go_fuzz_dep_.CoverTab[61054]++
													if field.Message() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:140
		_go_fuzz_dep_.CoverTab[61066]++
														p.TypeName = fullNameOf(field.Message())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:141
		// _ = "end of CoverTab[61066]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:142
		_go_fuzz_dep_.CoverTab[61067]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:142
		// _ = "end of CoverTab[61067]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:142
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:142
	// _ = "end of CoverTab[61054]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:142
	_go_fuzz_dep_.CoverTab[61055]++
													if field.HasJSONName() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:143
		_go_fuzz_dep_.CoverTab[61068]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:147
		if field.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:147
			_go_fuzz_dep_.CoverTab[61069]++
															p.JsonName = proto.String(strs.JSONCamelCase(string(field.Name())))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:148
			// _ = "end of CoverTab[61069]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:149
			_go_fuzz_dep_.CoverTab[61070]++
															p.JsonName = proto.String(field.JSONName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:150
			// _ = "end of CoverTab[61070]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:151
		// _ = "end of CoverTab[61068]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:152
		_go_fuzz_dep_.CoverTab[61071]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:152
		// _ = "end of CoverTab[61071]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:152
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:152
	// _ = "end of CoverTab[61055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:152
	_go_fuzz_dep_.CoverTab[61056]++
													if field.Syntax() == protoreflect.Proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:153
		_go_fuzz_dep_.CoverTab[61072]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:153
		return field.HasOptionalKeyword()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:153
		// _ = "end of CoverTab[61072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:153
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:153
		_go_fuzz_dep_.CoverTab[61073]++
														p.Proto3Optional = proto.Bool(true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:154
		// _ = "end of CoverTab[61073]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:155
		_go_fuzz_dep_.CoverTab[61074]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:155
		// _ = "end of CoverTab[61074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:155
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:155
	// _ = "end of CoverTab[61056]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:155
	_go_fuzz_dep_.CoverTab[61057]++
													if field.HasDefault() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:156
		_go_fuzz_dep_.CoverTab[61075]++
														def, err := defval.Marshal(field.Default(), field.DefaultEnumValue(), field.Kind(), defval.Descriptor)
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:158
			_go_fuzz_dep_.CoverTab[61077]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:158
			return field.DefaultEnumValue() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:158
			// _ = "end of CoverTab[61077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:158
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:158
			_go_fuzz_dep_.CoverTab[61078]++
															def = string(field.DefaultEnumValue().Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:159
			// _ = "end of CoverTab[61078]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:160
			_go_fuzz_dep_.CoverTab[61079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:160
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:160
				_go_fuzz_dep_.CoverTab[61080]++
																panic(fmt.Sprintf("%v: %v", field.FullName(), err))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:161
				// _ = "end of CoverTab[61080]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
				_go_fuzz_dep_.CoverTab[61081]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
				// _ = "end of CoverTab[61081]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
			// _ = "end of CoverTab[61079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
		// _ = "end of CoverTab[61075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:162
		_go_fuzz_dep_.CoverTab[61076]++
														p.DefaultValue = proto.String(def)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:163
		// _ = "end of CoverTab[61076]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:164
		_go_fuzz_dep_.CoverTab[61082]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:164
		// _ = "end of CoverTab[61082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:164
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:164
	// _ = "end of CoverTab[61057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:164
	_go_fuzz_dep_.CoverTab[61058]++
													if oneof := field.ContainingOneof(); oneof != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:165
		_go_fuzz_dep_.CoverTab[61083]++
														p.OneofIndex = proto.Int32(int32(oneof.Index()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:166
		// _ = "end of CoverTab[61083]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:167
		_go_fuzz_dep_.CoverTab[61084]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:167
		// _ = "end of CoverTab[61084]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:167
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:167
	// _ = "end of CoverTab[61058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:167
	_go_fuzz_dep_.CoverTab[61059]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:168
	// _ = "end of CoverTab[61059]"
}

// ToOneofDescriptorProto copies a protoreflect.OneofDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:171
// google.protobuf.OneofDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:173
func ToOneofDescriptorProto(oneof protoreflect.OneofDescriptor) *descriptorpb.OneofDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:173
	_go_fuzz_dep_.CoverTab[61085]++
													return &descriptorpb.OneofDescriptorProto{
		Name:		proto.String(string(oneof.Name())),
		Options:	proto.Clone(oneof.Options()).(*descriptorpb.OneofOptions),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:177
	// _ = "end of CoverTab[61085]"
}

// ToEnumDescriptorProto copies a protoreflect.EnumDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:180
// google.protobuf.EnumDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:182
func ToEnumDescriptorProto(enum protoreflect.EnumDescriptor) *descriptorpb.EnumDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:182
	_go_fuzz_dep_.CoverTab[61086]++
													p := &descriptorpb.EnumDescriptorProto{
		Name:		proto.String(string(enum.Name())),
		Options:	proto.Clone(enum.Options()).(*descriptorpb.EnumOptions),
	}
	for i, values := 0, enum.Values(); i < values.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:187
		_go_fuzz_dep_.CoverTab[61090]++
														p.Value = append(p.Value, ToEnumValueDescriptorProto(values.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:188
		// _ = "end of CoverTab[61090]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:189
	// _ = "end of CoverTab[61086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:189
	_go_fuzz_dep_.CoverTab[61087]++
													for i, ranges := 0, enum.ReservedRanges(); i < ranges.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:190
		_go_fuzz_dep_.CoverTab[61091]++
														rrange := ranges.Get(i)
														p.ReservedRange = append(p.ReservedRange, &descriptorpb.EnumDescriptorProto_EnumReservedRange{
			Start:	proto.Int32(int32(rrange[0])),
			End:	proto.Int32(int32(rrange[1])),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:195
		// _ = "end of CoverTab[61091]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:196
	// _ = "end of CoverTab[61087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:196
	_go_fuzz_dep_.CoverTab[61088]++
													for i, names := 0, enum.ReservedNames(); i < names.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:197
		_go_fuzz_dep_.CoverTab[61092]++
														p.ReservedName = append(p.ReservedName, string(names.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:198
		// _ = "end of CoverTab[61092]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:199
	// _ = "end of CoverTab[61088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:199
	_go_fuzz_dep_.CoverTab[61089]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:200
	// _ = "end of CoverTab[61089]"
}

// ToEnumValueDescriptorProto copies a protoreflect.EnumValueDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:203
// google.protobuf.EnumValueDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:205
func ToEnumValueDescriptorProto(value protoreflect.EnumValueDescriptor) *descriptorpb.EnumValueDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:205
	_go_fuzz_dep_.CoverTab[61093]++
													return &descriptorpb.EnumValueDescriptorProto{
		Name:		proto.String(string(value.Name())),
		Number:		proto.Int32(int32(value.Number())),
		Options:	proto.Clone(value.Options()).(*descriptorpb.EnumValueOptions),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:210
	// _ = "end of CoverTab[61093]"
}

// ToServiceDescriptorProto copies a protoreflect.ServiceDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:213
// google.protobuf.ServiceDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:215
func ToServiceDescriptorProto(service protoreflect.ServiceDescriptor) *descriptorpb.ServiceDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:215
	_go_fuzz_dep_.CoverTab[61094]++
													p := &descriptorpb.ServiceDescriptorProto{
		Name:		proto.String(string(service.Name())),
		Options:	proto.Clone(service.Options()).(*descriptorpb.ServiceOptions),
	}
	for i, methods := 0, service.Methods(); i < methods.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:220
		_go_fuzz_dep_.CoverTab[61096]++
														p.Method = append(p.Method, ToMethodDescriptorProto(methods.Get(i)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:221
		// _ = "end of CoverTab[61096]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:222
	// _ = "end of CoverTab[61094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:222
	_go_fuzz_dep_.CoverTab[61095]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:223
	// _ = "end of CoverTab[61095]"
}

// ToMethodDescriptorProto copies a protoreflect.MethodDescriptor into a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:226
// google.protobuf.MethodDescriptorProto message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:228
func ToMethodDescriptorProto(method protoreflect.MethodDescriptor) *descriptorpb.MethodDescriptorProto {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:228
	_go_fuzz_dep_.CoverTab[61097]++
													p := &descriptorpb.MethodDescriptorProto{
		Name:		proto.String(string(method.Name())),
		InputType:	fullNameOf(method.Input()),
		OutputType:	fullNameOf(method.Output()),
		Options:	proto.Clone(method.Options()).(*descriptorpb.MethodOptions),
	}
	if method.IsStreamingClient() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:235
		_go_fuzz_dep_.CoverTab[61100]++
														p.ClientStreaming = proto.Bool(true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:236
		// _ = "end of CoverTab[61100]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:237
		_go_fuzz_dep_.CoverTab[61101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:237
		// _ = "end of CoverTab[61101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:237
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:237
	// _ = "end of CoverTab[61097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:237
	_go_fuzz_dep_.CoverTab[61098]++
													if method.IsStreamingServer() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:238
		_go_fuzz_dep_.CoverTab[61102]++
														p.ServerStreaming = proto.Bool(true)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:239
		// _ = "end of CoverTab[61102]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:240
		_go_fuzz_dep_.CoverTab[61103]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:240
		// _ = "end of CoverTab[61103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:240
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:240
	// _ = "end of CoverTab[61098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:240
	_go_fuzz_dep_.CoverTab[61099]++
													return p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:241
	// _ = "end of CoverTab[61099]"
}

func fullNameOf(d protoreflect.Descriptor) *string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:244
	_go_fuzz_dep_.CoverTab[61104]++
													if d == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:245
		_go_fuzz_dep_.CoverTab[61107]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:246
		// _ = "end of CoverTab[61107]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:247
		_go_fuzz_dep_.CoverTab[61108]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:247
		// _ = "end of CoverTab[61108]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:247
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:247
	// _ = "end of CoverTab[61104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:247
	_go_fuzz_dep_.CoverTab[61105]++
													if strings.HasPrefix(string(d.FullName()), unknownPrefix) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:248
		_go_fuzz_dep_.CoverTab[61109]++
														return proto.String(string(d.FullName()[len(unknownPrefix):]))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:249
		// _ = "end of CoverTab[61109]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:250
		_go_fuzz_dep_.CoverTab[61110]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:250
		// _ = "end of CoverTab[61110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:250
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:250
	// _ = "end of CoverTab[61105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:250
	_go_fuzz_dep_.CoverTab[61106]++
													return proto.String("." + string(d.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:251
	// _ = "end of CoverTab[61106]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:252
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/proto.go:252
var _ = _go_fuzz_dep_.CoverTab
