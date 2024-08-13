// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
package filedesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:5
)

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/descopts"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (fd *File) lazyRawInit() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:19
	_go_fuzz_dep_.CoverTab[52822]++
													fd.unmarshalFull(fd.builder.RawDescriptor)
													fd.resolveMessages()
													fd.resolveExtensions()
													fd.resolveServices()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:23
	// _ = "end of CoverTab[52822]"
}

func (file *File) resolveMessages() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:26
	_go_fuzz_dep_.CoverTab[52823]++
													var depIdx int32
													for i := range file.allMessages {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:28
		_go_fuzz_dep_.CoverTab[52824]++
														md := &file.allMessages[i]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:32
		for j := range md.L2.Fields.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:32
			_go_fuzz_dep_.CoverTab[52825]++
															fd := &md.L2.Fields.List[j]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:36
			if fd.L1.IsWeak {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:36
				_go_fuzz_dep_.CoverTab[52828]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:37
				// _ = "end of CoverTab[52828]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:38
				_go_fuzz_dep_.CoverTab[52829]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:38
				// _ = "end of CoverTab[52829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:38
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:38
			// _ = "end of CoverTab[52825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:38
			_go_fuzz_dep_.CoverTab[52826]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:41
			switch fd.L1.Kind {
			case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:42
				_go_fuzz_dep_.CoverTab[52830]++
																fd.L1.Enum = file.resolveEnumDependency(fd.L1.Enum, listFieldDeps, depIdx)
																depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:44
				// _ = "end of CoverTab[52830]"
			case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:45
				_go_fuzz_dep_.CoverTab[52831]++
																fd.L1.Message = file.resolveMessageDependency(fd.L1.Message, listFieldDeps, depIdx)
																depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:47
				// _ = "end of CoverTab[52831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:47
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:47
				_go_fuzz_dep_.CoverTab[52832]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:47
				// _ = "end of CoverTab[52832]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:48
			// _ = "end of CoverTab[52826]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:48
			_go_fuzz_dep_.CoverTab[52827]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:51
			if v := fd.L1.Default.val; v.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:51
				_go_fuzz_dep_.CoverTab[52833]++
																fd.L1.Default = unmarshalDefault(v.Bytes(), fd.L1.Kind, file, fd.L1.Enum)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:52
				// _ = "end of CoverTab[52833]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:53
				_go_fuzz_dep_.CoverTab[52834]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:53
				// _ = "end of CoverTab[52834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:53
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:53
			// _ = "end of CoverTab[52827]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:54
		// _ = "end of CoverTab[52824]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:55
	// _ = "end of CoverTab[52823]"
}

func (file *File) resolveExtensions() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:58
	_go_fuzz_dep_.CoverTab[52835]++
													var depIdx int32
													for i := range file.allExtensions {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:60
		_go_fuzz_dep_.CoverTab[52836]++
														xd := &file.allExtensions[i]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:64
		switch xd.L1.Kind {
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:65
			_go_fuzz_dep_.CoverTab[52838]++
															xd.L2.Enum = file.resolveEnumDependency(xd.L2.Enum, listExtDeps, depIdx)
															depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:67
			// _ = "end of CoverTab[52838]"
		case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:68
			_go_fuzz_dep_.CoverTab[52839]++
															xd.L2.Message = file.resolveMessageDependency(xd.L2.Message, listExtDeps, depIdx)
															depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:70
			// _ = "end of CoverTab[52839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:70
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:70
			_go_fuzz_dep_.CoverTab[52840]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:70
			// _ = "end of CoverTab[52840]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:71
		// _ = "end of CoverTab[52836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:71
		_go_fuzz_dep_.CoverTab[52837]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:74
		if v := xd.L2.Default.val; v.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:74
			_go_fuzz_dep_.CoverTab[52841]++
															xd.L2.Default = unmarshalDefault(v.Bytes(), xd.L1.Kind, file, xd.L2.Enum)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:75
			// _ = "end of CoverTab[52841]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:76
			_go_fuzz_dep_.CoverTab[52842]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:76
			// _ = "end of CoverTab[52842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:76
		// _ = "end of CoverTab[52837]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:77
	// _ = "end of CoverTab[52835]"
}

func (file *File) resolveServices() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:80
	_go_fuzz_dep_.CoverTab[52843]++
													var depIdx int32
													for i := range file.allServices {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:82
		_go_fuzz_dep_.CoverTab[52844]++
														sd := &file.allServices[i]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:86
		for j := range sd.L2.Methods.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:86
			_go_fuzz_dep_.CoverTab[52845]++
															md := &sd.L2.Methods.List[j]
															md.L1.Input = file.resolveMessageDependency(md.L1.Input, listMethInDeps, depIdx)
															md.L1.Output = file.resolveMessageDependency(md.L1.Output, listMethOutDeps, depIdx)
															depIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:90
			// _ = "end of CoverTab[52845]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:91
		// _ = "end of CoverTab[52844]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:92
	// _ = "end of CoverTab[52843]"
}

func (file *File) resolveEnumDependency(ed protoreflect.EnumDescriptor, i, j int32) protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:95
	_go_fuzz_dep_.CoverTab[52846]++
													r := file.builder.FileRegistry
													if r, ok := r.(resolverByIndex); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:97
		_go_fuzz_dep_.CoverTab[52850]++
														if ed2 := r.FindEnumByIndex(i, j, file.allEnums, file.allMessages); ed2 != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:98
			_go_fuzz_dep_.CoverTab[52851]++
															return ed2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:99
			// _ = "end of CoverTab[52851]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:100
			_go_fuzz_dep_.CoverTab[52852]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:100
			// _ = "end of CoverTab[52852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:100
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:100
		// _ = "end of CoverTab[52850]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:101
		_go_fuzz_dep_.CoverTab[52853]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:101
		// _ = "end of CoverTab[52853]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:101
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:101
	// _ = "end of CoverTab[52846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:101
	_go_fuzz_dep_.CoverTab[52847]++
														for i := range file.allEnums {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:102
		_go_fuzz_dep_.CoverTab[52854]++
															if ed2 := &file.allEnums[i]; ed2.L0.FullName == ed.FullName() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:103
			_go_fuzz_dep_.CoverTab[52855]++
																return ed2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:104
			// _ = "end of CoverTab[52855]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:105
			_go_fuzz_dep_.CoverTab[52856]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:105
			// _ = "end of CoverTab[52856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:105
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:105
		// _ = "end of CoverTab[52854]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:106
	// _ = "end of CoverTab[52847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:106
	_go_fuzz_dep_.CoverTab[52848]++
														if d, _ := r.FindDescriptorByName(ed.FullName()); d != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:107
		_go_fuzz_dep_.CoverTab[52857]++
															return d.(protoreflect.EnumDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:108
		// _ = "end of CoverTab[52857]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:109
		_go_fuzz_dep_.CoverTab[52858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:109
		// _ = "end of CoverTab[52858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:109
	// _ = "end of CoverTab[52848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:109
	_go_fuzz_dep_.CoverTab[52849]++
														return ed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:110
	// _ = "end of CoverTab[52849]"
}

func (file *File) resolveMessageDependency(md protoreflect.MessageDescriptor, i, j int32) protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:113
	_go_fuzz_dep_.CoverTab[52859]++
														r := file.builder.FileRegistry
														if r, ok := r.(resolverByIndex); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:115
		_go_fuzz_dep_.CoverTab[52863]++
															if md2 := r.FindMessageByIndex(i, j, file.allEnums, file.allMessages); md2 != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:116
			_go_fuzz_dep_.CoverTab[52864]++
																return md2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:117
			// _ = "end of CoverTab[52864]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:118
			_go_fuzz_dep_.CoverTab[52865]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:118
			// _ = "end of CoverTab[52865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:118
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:118
		// _ = "end of CoverTab[52863]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:119
		_go_fuzz_dep_.CoverTab[52866]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:119
		// _ = "end of CoverTab[52866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:119
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:119
	// _ = "end of CoverTab[52859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:119
	_go_fuzz_dep_.CoverTab[52860]++
														for i := range file.allMessages {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:120
		_go_fuzz_dep_.CoverTab[52867]++
															if md2 := &file.allMessages[i]; md2.L0.FullName == md.FullName() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:121
			_go_fuzz_dep_.CoverTab[52868]++
																return md2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:122
			// _ = "end of CoverTab[52868]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:123
			_go_fuzz_dep_.CoverTab[52869]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:123
			// _ = "end of CoverTab[52869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:123
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:123
		// _ = "end of CoverTab[52867]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:124
	// _ = "end of CoverTab[52860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:124
	_go_fuzz_dep_.CoverTab[52861]++
														if d, _ := r.FindDescriptorByName(md.FullName()); d != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:125
		_go_fuzz_dep_.CoverTab[52870]++
															return d.(protoreflect.MessageDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:126
		// _ = "end of CoverTab[52870]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:127
		_go_fuzz_dep_.CoverTab[52871]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:127
		// _ = "end of CoverTab[52871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:127
	// _ = "end of CoverTab[52861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:127
	_go_fuzz_dep_.CoverTab[52862]++
														return md
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:128
	// _ = "end of CoverTab[52862]"
}

func (fd *File) unmarshalFull(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:131
	_go_fuzz_dep_.CoverTab[52872]++
														sb := getBuilder()
														defer putBuilder(sb)

														var enumIdx, messageIdx, extensionIdx, serviceIdx int
														var rawOptions []byte
														fd.L2 = new(FileL2)
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:138
		_go_fuzz_dep_.CoverTab[52874]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:142
			_go_fuzz_dep_.CoverTab[52875]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.FileDescriptorProto_PublicDependency_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:146
				_go_fuzz_dep_.CoverTab[52878]++
																	fd.L2.Imports[v].IsPublic = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:147
				// _ = "end of CoverTab[52878]"
			case genid.FileDescriptorProto_WeakDependency_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:148
				_go_fuzz_dep_.CoverTab[52879]++
																	fd.L2.Imports[v].IsWeak = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:149
				// _ = "end of CoverTab[52879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:149
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:149
				_go_fuzz_dep_.CoverTab[52880]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:149
				// _ = "end of CoverTab[52880]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:150
			// _ = "end of CoverTab[52875]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:151
			_go_fuzz_dep_.CoverTab[52876]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.FileDescriptorProto_Dependency_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:155
				_go_fuzz_dep_.CoverTab[52881]++
																	path := sb.MakeString(v)
																	imp, _ := fd.builder.FileRegistry.FindFileByPath(path)
																	if imp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:158
					_go_fuzz_dep_.CoverTab[52889]++
																		imp = PlaceholderFile(path)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:159
					// _ = "end of CoverTab[52889]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:160
					_go_fuzz_dep_.CoverTab[52890]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:160
					// _ = "end of CoverTab[52890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:160
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:160
				// _ = "end of CoverTab[52881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:160
				_go_fuzz_dep_.CoverTab[52882]++
																	fd.L2.Imports = append(fd.L2.Imports, protoreflect.FileImport{FileDescriptor: imp})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:161
				// _ = "end of CoverTab[52882]"
			case genid.FileDescriptorProto_EnumType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:162
				_go_fuzz_dep_.CoverTab[52883]++
																	fd.L1.Enums.List[enumIdx].unmarshalFull(v, sb)
																	enumIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:164
				// _ = "end of CoverTab[52883]"
			case genid.FileDescriptorProto_MessageType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:165
				_go_fuzz_dep_.CoverTab[52884]++
																	fd.L1.Messages.List[messageIdx].unmarshalFull(v, sb)
																	messageIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:167
				// _ = "end of CoverTab[52884]"
			case genid.FileDescriptorProto_Extension_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:168
				_go_fuzz_dep_.CoverTab[52885]++
																	fd.L1.Extensions.List[extensionIdx].unmarshalFull(v, sb)
																	extensionIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:170
				// _ = "end of CoverTab[52885]"
			case genid.FileDescriptorProto_Service_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:171
				_go_fuzz_dep_.CoverTab[52886]++
																	fd.L1.Services.List[serviceIdx].unmarshalFull(v, sb)
																	serviceIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:173
				// _ = "end of CoverTab[52886]"
			case genid.FileDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:174
				_go_fuzz_dep_.CoverTab[52887]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:175
				// _ = "end of CoverTab[52887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:175
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:175
				_go_fuzz_dep_.CoverTab[52888]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:175
				// _ = "end of CoverTab[52888]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:176
			// _ = "end of CoverTab[52876]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:177
			_go_fuzz_dep_.CoverTab[52877]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:179
			// _ = "end of CoverTab[52877]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:180
		// _ = "end of CoverTab[52874]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:181
	// _ = "end of CoverTab[52872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:181
	_go_fuzz_dep_.CoverTab[52873]++
														fd.L2.Options = fd.builder.optionsUnmarshaler(&descopts.File, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:182
	// _ = "end of CoverTab[52873]"
}

func (ed *Enum) unmarshalFull(b []byte, sb *strs.Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:185
	_go_fuzz_dep_.CoverTab[52891]++
														var rawValues [][]byte
														var rawOptions []byte
														if !ed.L1.eagerValues {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:188
		_go_fuzz_dep_.CoverTab[52895]++
															ed.L2 = new(EnumL2)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:189
		// _ = "end of CoverTab[52895]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:190
		_go_fuzz_dep_.CoverTab[52896]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:190
		// _ = "end of CoverTab[52896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:190
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:190
	// _ = "end of CoverTab[52891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:190
	_go_fuzz_dep_.CoverTab[52892]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:191
		_go_fuzz_dep_.CoverTab[52897]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:195
			_go_fuzz_dep_.CoverTab[52898]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.EnumDescriptorProto_Value_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:199
				_go_fuzz_dep_.CoverTab[52900]++
																	rawValues = append(rawValues, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:200
				// _ = "end of CoverTab[52900]"
			case genid.EnumDescriptorProto_ReservedName_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:201
				_go_fuzz_dep_.CoverTab[52901]++
																	ed.L2.ReservedNames.List = append(ed.L2.ReservedNames.List, protoreflect.Name(sb.MakeString(v)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:202
				// _ = "end of CoverTab[52901]"
			case genid.EnumDescriptorProto_ReservedRange_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:203
				_go_fuzz_dep_.CoverTab[52902]++
																	ed.L2.ReservedRanges.List = append(ed.L2.ReservedRanges.List, unmarshalEnumReservedRange(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:204
				// _ = "end of CoverTab[52902]"
			case genid.EnumDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:205
				_go_fuzz_dep_.CoverTab[52903]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:206
				// _ = "end of CoverTab[52903]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:206
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:206
				_go_fuzz_dep_.CoverTab[52904]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:206
				// _ = "end of CoverTab[52904]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:207
			// _ = "end of CoverTab[52898]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:208
			_go_fuzz_dep_.CoverTab[52899]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:210
			// _ = "end of CoverTab[52899]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:211
		// _ = "end of CoverTab[52897]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:212
	// _ = "end of CoverTab[52892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:212
	_go_fuzz_dep_.CoverTab[52893]++
														if !ed.L1.eagerValues && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:213
		_go_fuzz_dep_.CoverTab[52905]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:213
		return len(rawValues) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:213
		// _ = "end of CoverTab[52905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:213
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:213
		_go_fuzz_dep_.CoverTab[52906]++
															ed.L2.Values.List = make([]EnumValue, len(rawValues))
															for i, b := range rawValues {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:215
			_go_fuzz_dep_.CoverTab[52907]++
																ed.L2.Values.List[i].unmarshalFull(b, sb, ed.L0.ParentFile, ed, i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:216
			// _ = "end of CoverTab[52907]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:217
		// _ = "end of CoverTab[52906]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:218
		_go_fuzz_dep_.CoverTab[52908]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:218
		// _ = "end of CoverTab[52908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:218
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:218
	// _ = "end of CoverTab[52893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:218
	_go_fuzz_dep_.CoverTab[52894]++
														ed.L2.Options = ed.L0.ParentFile.builder.optionsUnmarshaler(&descopts.Enum, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:219
	// _ = "end of CoverTab[52894]"
}

func unmarshalEnumReservedRange(b []byte) (r [2]protoreflect.EnumNumber) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:222
	_go_fuzz_dep_.CoverTab[52909]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:223
		_go_fuzz_dep_.CoverTab[52911]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:227
			_go_fuzz_dep_.CoverTab[52912]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.EnumDescriptorProto_EnumReservedRange_Start_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:231
				_go_fuzz_dep_.CoverTab[52914]++
																	r[0] = protoreflect.EnumNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:232
				// _ = "end of CoverTab[52914]"
			case genid.EnumDescriptorProto_EnumReservedRange_End_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:233
				_go_fuzz_dep_.CoverTab[52915]++
																	r[1] = protoreflect.EnumNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:234
				// _ = "end of CoverTab[52915]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:234
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:234
				_go_fuzz_dep_.CoverTab[52916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:234
				// _ = "end of CoverTab[52916]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:235
			// _ = "end of CoverTab[52912]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:236
			_go_fuzz_dep_.CoverTab[52913]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:238
			// _ = "end of CoverTab[52913]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:239
		// _ = "end of CoverTab[52911]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:240
	// _ = "end of CoverTab[52909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:240
	_go_fuzz_dep_.CoverTab[52910]++
														return r
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:241
	// _ = "end of CoverTab[52910]"
}

func (vd *EnumValue) unmarshalFull(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:244
	_go_fuzz_dep_.CoverTab[52917]++
														vd.L0.ParentFile = pf
														vd.L0.Parent = pd
														vd.L0.Index = i

														var rawOptions []byte
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:250
		_go_fuzz_dep_.CoverTab[52919]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:254
			_go_fuzz_dep_.CoverTab[52920]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.EnumValueDescriptorProto_Number_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:258
				_go_fuzz_dep_.CoverTab[52923]++
																	vd.L1.Number = protoreflect.EnumNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:259
				// _ = "end of CoverTab[52923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:259
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:259
				_go_fuzz_dep_.CoverTab[52924]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:259
				// _ = "end of CoverTab[52924]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:260
			// _ = "end of CoverTab[52920]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:261
			_go_fuzz_dep_.CoverTab[52921]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.EnumValueDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:265
				_go_fuzz_dep_.CoverTab[52925]++

																	vd.L0.FullName = appendFullName(sb, pd.Parent().FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:267
				// _ = "end of CoverTab[52925]"
			case genid.EnumValueDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:268
				_go_fuzz_dep_.CoverTab[52926]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:269
				// _ = "end of CoverTab[52926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:269
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:269
				_go_fuzz_dep_.CoverTab[52927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:269
				// _ = "end of CoverTab[52927]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:270
			// _ = "end of CoverTab[52921]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:271
			_go_fuzz_dep_.CoverTab[52922]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:273
			// _ = "end of CoverTab[52922]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:274
		// _ = "end of CoverTab[52919]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:275
	// _ = "end of CoverTab[52917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:275
	_go_fuzz_dep_.CoverTab[52918]++
														vd.L1.Options = pf.builder.optionsUnmarshaler(&descopts.EnumValue, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:276
	// _ = "end of CoverTab[52918]"
}

func (md *Message) unmarshalFull(b []byte, sb *strs.Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:279
	_go_fuzz_dep_.CoverTab[52928]++
														var rawFields, rawOneofs [][]byte
														var enumIdx, messageIdx, extensionIdx int
														var rawOptions []byte
														md.L2 = new(MessageL2)
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:284
		_go_fuzz_dep_.CoverTab[52931]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:288
			_go_fuzz_dep_.CoverTab[52932]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.DescriptorProto_Field_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:292
				_go_fuzz_dep_.CoverTab[52934]++
																	rawFields = append(rawFields, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:293
				// _ = "end of CoverTab[52934]"
			case genid.DescriptorProto_OneofDecl_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:294
				_go_fuzz_dep_.CoverTab[52935]++
																	rawOneofs = append(rawOneofs, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:295
				// _ = "end of CoverTab[52935]"
			case genid.DescriptorProto_ReservedName_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:296
				_go_fuzz_dep_.CoverTab[52936]++
																	md.L2.ReservedNames.List = append(md.L2.ReservedNames.List, protoreflect.Name(sb.MakeString(v)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:297
				// _ = "end of CoverTab[52936]"
			case genid.DescriptorProto_ReservedRange_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:298
				_go_fuzz_dep_.CoverTab[52937]++
																	md.L2.ReservedRanges.List = append(md.L2.ReservedRanges.List, unmarshalMessageReservedRange(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:299
				// _ = "end of CoverTab[52937]"
			case genid.DescriptorProto_ExtensionRange_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:300
				_go_fuzz_dep_.CoverTab[52938]++
																	r, rawOptions := unmarshalMessageExtensionRange(v)
																	opts := md.L0.ParentFile.builder.optionsUnmarshaler(&descopts.ExtensionRange, rawOptions)
																	md.L2.ExtensionRanges.List = append(md.L2.ExtensionRanges.List, r)
																	md.L2.ExtensionRangeOptions = append(md.L2.ExtensionRangeOptions, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:304
				// _ = "end of CoverTab[52938]"
			case genid.DescriptorProto_EnumType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:305
				_go_fuzz_dep_.CoverTab[52939]++
																	md.L1.Enums.List[enumIdx].unmarshalFull(v, sb)
																	enumIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:307
				// _ = "end of CoverTab[52939]"
			case genid.DescriptorProto_NestedType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:308
				_go_fuzz_dep_.CoverTab[52940]++
																	md.L1.Messages.List[messageIdx].unmarshalFull(v, sb)
																	messageIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:310
				// _ = "end of CoverTab[52940]"
			case genid.DescriptorProto_Extension_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:311
				_go_fuzz_dep_.CoverTab[52941]++
																	md.L1.Extensions.List[extensionIdx].unmarshalFull(v, sb)
																	extensionIdx++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:313
				// _ = "end of CoverTab[52941]"
			case genid.DescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:314
				_go_fuzz_dep_.CoverTab[52942]++
																	md.unmarshalOptions(v)
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:316
				// _ = "end of CoverTab[52942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:316
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:316
				_go_fuzz_dep_.CoverTab[52943]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:316
				// _ = "end of CoverTab[52943]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:317
			// _ = "end of CoverTab[52932]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:318
			_go_fuzz_dep_.CoverTab[52933]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:320
			// _ = "end of CoverTab[52933]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:321
		// _ = "end of CoverTab[52931]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:322
	// _ = "end of CoverTab[52928]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:322
	_go_fuzz_dep_.CoverTab[52929]++
														if len(rawFields) > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:323
		_go_fuzz_dep_.CoverTab[52944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:323
		return len(rawOneofs) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:323
		// _ = "end of CoverTab[52944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:323
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:323
		_go_fuzz_dep_.CoverTab[52945]++
															md.L2.Fields.List = make([]Field, len(rawFields))
															md.L2.Oneofs.List = make([]Oneof, len(rawOneofs))
															for i, b := range rawFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:326
			_go_fuzz_dep_.CoverTab[52947]++
																fd := &md.L2.Fields.List[i]
																fd.unmarshalFull(b, sb, md.L0.ParentFile, md, i)
																if fd.L1.Cardinality == protoreflect.Required {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:329
				_go_fuzz_dep_.CoverTab[52948]++
																	md.L2.RequiredNumbers.List = append(md.L2.RequiredNumbers.List, fd.L1.Number)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:330
				// _ = "end of CoverTab[52948]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:331
				_go_fuzz_dep_.CoverTab[52949]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:331
				// _ = "end of CoverTab[52949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:331
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:331
			// _ = "end of CoverTab[52947]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:332
		// _ = "end of CoverTab[52945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:332
		_go_fuzz_dep_.CoverTab[52946]++
															for i, b := range rawOneofs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:333
			_go_fuzz_dep_.CoverTab[52950]++
																od := &md.L2.Oneofs.List[i]
																od.unmarshalFull(b, sb, md.L0.ParentFile, md, i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:335
			// _ = "end of CoverTab[52950]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:336
		// _ = "end of CoverTab[52946]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:337
		_go_fuzz_dep_.CoverTab[52951]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:337
		// _ = "end of CoverTab[52951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:337
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:337
	// _ = "end of CoverTab[52929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:337
	_go_fuzz_dep_.CoverTab[52930]++
														md.L2.Options = md.L0.ParentFile.builder.optionsUnmarshaler(&descopts.Message, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:338
	// _ = "end of CoverTab[52930]"
}

func (md *Message) unmarshalOptions(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:341
	_go_fuzz_dep_.CoverTab[52952]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:342
		_go_fuzz_dep_.CoverTab[52953]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:346
			_go_fuzz_dep_.CoverTab[52954]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.MessageOptions_MapEntry_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:350
				_go_fuzz_dep_.CoverTab[52956]++
																	md.L1.IsMapEntry = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:351
				// _ = "end of CoverTab[52956]"
			case genid.MessageOptions_MessageSetWireFormat_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:352
				_go_fuzz_dep_.CoverTab[52957]++
																	md.L1.IsMessageSet = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:353
				// _ = "end of CoverTab[52957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:353
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:353
				_go_fuzz_dep_.CoverTab[52958]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:353
				// _ = "end of CoverTab[52958]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:354
			// _ = "end of CoverTab[52954]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:355
			_go_fuzz_dep_.CoverTab[52955]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:357
			// _ = "end of CoverTab[52955]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:358
		// _ = "end of CoverTab[52953]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:359
	// _ = "end of CoverTab[52952]"
}

func unmarshalMessageReservedRange(b []byte) (r [2]protoreflect.FieldNumber) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:362
	_go_fuzz_dep_.CoverTab[52959]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:363
		_go_fuzz_dep_.CoverTab[52961]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:367
			_go_fuzz_dep_.CoverTab[52962]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.DescriptorProto_ReservedRange_Start_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:371
				_go_fuzz_dep_.CoverTab[52964]++
																	r[0] = protoreflect.FieldNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:372
				// _ = "end of CoverTab[52964]"
			case genid.DescriptorProto_ReservedRange_End_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:373
				_go_fuzz_dep_.CoverTab[52965]++
																	r[1] = protoreflect.FieldNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:374
				// _ = "end of CoverTab[52965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:374
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:374
				_go_fuzz_dep_.CoverTab[52966]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:374
				// _ = "end of CoverTab[52966]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:375
			// _ = "end of CoverTab[52962]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:376
			_go_fuzz_dep_.CoverTab[52963]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:378
			// _ = "end of CoverTab[52963]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:379
		// _ = "end of CoverTab[52961]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:380
	// _ = "end of CoverTab[52959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:380
	_go_fuzz_dep_.CoverTab[52960]++
														return r
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:381
	// _ = "end of CoverTab[52960]"
}

func unmarshalMessageExtensionRange(b []byte) (r [2]protoreflect.FieldNumber, rawOptions []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:384
	_go_fuzz_dep_.CoverTab[52967]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:385
		_go_fuzz_dep_.CoverTab[52969]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:389
			_go_fuzz_dep_.CoverTab[52970]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.DescriptorProto_ExtensionRange_Start_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:393
				_go_fuzz_dep_.CoverTab[52973]++
																	r[0] = protoreflect.FieldNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:394
				// _ = "end of CoverTab[52973]"
			case genid.DescriptorProto_ExtensionRange_End_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:395
				_go_fuzz_dep_.CoverTab[52974]++
																	r[1] = protoreflect.FieldNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:396
				// _ = "end of CoverTab[52974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:396
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:396
				_go_fuzz_dep_.CoverTab[52975]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:396
				// _ = "end of CoverTab[52975]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:397
			// _ = "end of CoverTab[52970]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:398
			_go_fuzz_dep_.CoverTab[52971]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.DescriptorProto_ExtensionRange_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:402
				_go_fuzz_dep_.CoverTab[52976]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:403
				// _ = "end of CoverTab[52976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:403
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:403
				_go_fuzz_dep_.CoverTab[52977]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:403
				// _ = "end of CoverTab[52977]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:404
			// _ = "end of CoverTab[52971]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:405
			_go_fuzz_dep_.CoverTab[52972]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:407
			// _ = "end of CoverTab[52972]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:408
		// _ = "end of CoverTab[52969]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:409
	// _ = "end of CoverTab[52967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:409
	_go_fuzz_dep_.CoverTab[52968]++
														return r, rawOptions
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:410
	// _ = "end of CoverTab[52968]"
}

func (fd *Field) unmarshalFull(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:413
	_go_fuzz_dep_.CoverTab[52978]++
														fd.L0.ParentFile = pf
														fd.L0.Parent = pd
														fd.L0.Index = i

														var rawTypeName []byte
														var rawOptions []byte
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:420
		_go_fuzz_dep_.CoverTab[52981]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:424
			_go_fuzz_dep_.CoverTab[52982]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.FieldDescriptorProto_Number_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:428
				_go_fuzz_dep_.CoverTab[52985]++
																	fd.L1.Number = protoreflect.FieldNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:429
				// _ = "end of CoverTab[52985]"
			case genid.FieldDescriptorProto_Label_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:430
				_go_fuzz_dep_.CoverTab[52986]++
																	fd.L1.Cardinality = protoreflect.Cardinality(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:431
				// _ = "end of CoverTab[52986]"
			case genid.FieldDescriptorProto_Type_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:432
				_go_fuzz_dep_.CoverTab[52987]++
																	fd.L1.Kind = protoreflect.Kind(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:433
				// _ = "end of CoverTab[52987]"
			case genid.FieldDescriptorProto_OneofIndex_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:434
				_go_fuzz_dep_.CoverTab[52988]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:438
				od := &pd.(*Message).L2.Oneofs.List[v]
				od.L1.Fields.List = append(od.L1.Fields.List, fd)
				if fd.L1.ContainingOneof != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:440
					_go_fuzz_dep_.CoverTab[52992]++
																		panic("oneof type already set")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:441
					// _ = "end of CoverTab[52992]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:442
					_go_fuzz_dep_.CoverTab[52993]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:442
					// _ = "end of CoverTab[52993]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:442
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:442
				// _ = "end of CoverTab[52988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:442
				_go_fuzz_dep_.CoverTab[52989]++
																	fd.L1.ContainingOneof = od
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:443
				// _ = "end of CoverTab[52989]"
			case genid.FieldDescriptorProto_Proto3Optional_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:444
				_go_fuzz_dep_.CoverTab[52990]++
																	fd.L1.IsProto3Optional = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:445
				// _ = "end of CoverTab[52990]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:445
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:445
				_go_fuzz_dep_.CoverTab[52991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:445
				// _ = "end of CoverTab[52991]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:446
			// _ = "end of CoverTab[52982]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:447
			_go_fuzz_dep_.CoverTab[52983]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.FieldDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:451
				_go_fuzz_dep_.CoverTab[52994]++
																	fd.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:452
				// _ = "end of CoverTab[52994]"
			case genid.FieldDescriptorProto_JsonName_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:453
				_go_fuzz_dep_.CoverTab[52995]++
																	fd.L1.StringName.InitJSON(sb.MakeString(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:454
				// _ = "end of CoverTab[52995]"
			case genid.FieldDescriptorProto_DefaultValue_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:455
				_go_fuzz_dep_.CoverTab[52996]++
																	fd.L1.Default.val = protoreflect.ValueOfBytes(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:456
				// _ = "end of CoverTab[52996]"
			case genid.FieldDescriptorProto_TypeName_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:457
				_go_fuzz_dep_.CoverTab[52997]++
																	rawTypeName = v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:458
				// _ = "end of CoverTab[52997]"
			case genid.FieldDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:459
				_go_fuzz_dep_.CoverTab[52998]++
																	fd.unmarshalOptions(v)
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:461
				// _ = "end of CoverTab[52998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:461
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:461
				_go_fuzz_dep_.CoverTab[52999]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:461
				// _ = "end of CoverTab[52999]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:462
			// _ = "end of CoverTab[52983]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:463
			_go_fuzz_dep_.CoverTab[52984]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:465
			// _ = "end of CoverTab[52984]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:466
		// _ = "end of CoverTab[52981]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:467
	// _ = "end of CoverTab[52978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:467
	_go_fuzz_dep_.CoverTab[52979]++
														if rawTypeName != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:468
		_go_fuzz_dep_.CoverTab[53000]++
															name := makeFullName(sb, rawTypeName)
															switch fd.L1.Kind {
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:471
			_go_fuzz_dep_.CoverTab[53001]++
																fd.L1.Enum = PlaceholderEnum(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:472
			// _ = "end of CoverTab[53001]"
		case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:473
			_go_fuzz_dep_.CoverTab[53002]++
																fd.L1.Message = PlaceholderMessage(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:474
			// _ = "end of CoverTab[53002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:474
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:474
			_go_fuzz_dep_.CoverTab[53003]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:474
			// _ = "end of CoverTab[53003]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:475
		// _ = "end of CoverTab[53000]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:476
		_go_fuzz_dep_.CoverTab[53004]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:476
		// _ = "end of CoverTab[53004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:476
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:476
	// _ = "end of CoverTab[52979]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:476
	_go_fuzz_dep_.CoverTab[52980]++
														fd.L1.Options = pf.builder.optionsUnmarshaler(&descopts.Field, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:477
	// _ = "end of CoverTab[52980]"
}

func (fd *Field) unmarshalOptions(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:480
	_go_fuzz_dep_.CoverTab[53005]++
														const FieldOptions_EnforceUTF8 = 13

														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:483
		_go_fuzz_dep_.CoverTab[53006]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:487
			_go_fuzz_dep_.CoverTab[53007]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.FieldOptions_Packed_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:491
				_go_fuzz_dep_.CoverTab[53009]++
																	fd.L1.HasPacked = true
																	fd.L1.IsPacked = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:493
				// _ = "end of CoverTab[53009]"
			case genid.FieldOptions_Weak_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:494
				_go_fuzz_dep_.CoverTab[53010]++
																	fd.L1.IsWeak = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:495
				// _ = "end of CoverTab[53010]"
			case FieldOptions_EnforceUTF8:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:496
				_go_fuzz_dep_.CoverTab[53011]++
																	fd.L1.HasEnforceUTF8 = true
																	fd.L1.EnforceUTF8 = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:498
				// _ = "end of CoverTab[53011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:498
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:498
				_go_fuzz_dep_.CoverTab[53012]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:498
				// _ = "end of CoverTab[53012]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:499
			// _ = "end of CoverTab[53007]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:500
			_go_fuzz_dep_.CoverTab[53008]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:502
			// _ = "end of CoverTab[53008]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:503
		// _ = "end of CoverTab[53006]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:504
	// _ = "end of CoverTab[53005]"
}

func (od *Oneof) unmarshalFull(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:507
	_go_fuzz_dep_.CoverTab[53013]++
														od.L0.ParentFile = pf
														od.L0.Parent = pd
														od.L0.Index = i

														var rawOptions []byte
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:513
		_go_fuzz_dep_.CoverTab[53015]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:517
			_go_fuzz_dep_.CoverTab[53016]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.OneofDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:521
				_go_fuzz_dep_.CoverTab[53018]++
																	od.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:522
				// _ = "end of CoverTab[53018]"
			case genid.OneofDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:523
				_go_fuzz_dep_.CoverTab[53019]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:524
				// _ = "end of CoverTab[53019]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:524
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:524
				_go_fuzz_dep_.CoverTab[53020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:524
				// _ = "end of CoverTab[53020]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:525
			// _ = "end of CoverTab[53016]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:526
			_go_fuzz_dep_.CoverTab[53017]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:528
			// _ = "end of CoverTab[53017]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:529
		// _ = "end of CoverTab[53015]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:530
	// _ = "end of CoverTab[53013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:530
	_go_fuzz_dep_.CoverTab[53014]++
														od.L1.Options = pf.builder.optionsUnmarshaler(&descopts.Oneof, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:531
	// _ = "end of CoverTab[53014]"
}

func (xd *Extension) unmarshalFull(b []byte, sb *strs.Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:534
	_go_fuzz_dep_.CoverTab[53021]++
														var rawTypeName []byte
														var rawOptions []byte
														xd.L2 = new(ExtensionL2)
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:538
		_go_fuzz_dep_.CoverTab[53024]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:542
			_go_fuzz_dep_.CoverTab[53025]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.FieldDescriptorProto_Proto3Optional_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:546
				_go_fuzz_dep_.CoverTab[53028]++
																	xd.L2.IsProto3Optional = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:547
				// _ = "end of CoverTab[53028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:547
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:547
				_go_fuzz_dep_.CoverTab[53029]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:547
				// _ = "end of CoverTab[53029]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:548
			// _ = "end of CoverTab[53025]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:549
			_go_fuzz_dep_.CoverTab[53026]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.FieldDescriptorProto_JsonName_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:553
				_go_fuzz_dep_.CoverTab[53030]++
																	xd.L2.StringName.InitJSON(sb.MakeString(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:554
				// _ = "end of CoverTab[53030]"
			case genid.FieldDescriptorProto_DefaultValue_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:555
				_go_fuzz_dep_.CoverTab[53031]++
																	xd.L2.Default.val = protoreflect.ValueOfBytes(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:556
				// _ = "end of CoverTab[53031]"
			case genid.FieldDescriptorProto_TypeName_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:557
				_go_fuzz_dep_.CoverTab[53032]++
																	rawTypeName = v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:558
				// _ = "end of CoverTab[53032]"
			case genid.FieldDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:559
				_go_fuzz_dep_.CoverTab[53033]++
																	xd.unmarshalOptions(v)
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:561
				// _ = "end of CoverTab[53033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:561
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:561
				_go_fuzz_dep_.CoverTab[53034]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:561
				// _ = "end of CoverTab[53034]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:562
			// _ = "end of CoverTab[53026]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:563
			_go_fuzz_dep_.CoverTab[53027]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:565
			// _ = "end of CoverTab[53027]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:566
		// _ = "end of CoverTab[53024]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:567
	// _ = "end of CoverTab[53021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:567
	_go_fuzz_dep_.CoverTab[53022]++
														if rawTypeName != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:568
		_go_fuzz_dep_.CoverTab[53035]++
															name := makeFullName(sb, rawTypeName)
															switch xd.L1.Kind {
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:571
			_go_fuzz_dep_.CoverTab[53036]++
																xd.L2.Enum = PlaceholderEnum(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:572
			// _ = "end of CoverTab[53036]"
		case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:573
			_go_fuzz_dep_.CoverTab[53037]++
																xd.L2.Message = PlaceholderMessage(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:574
			// _ = "end of CoverTab[53037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:574
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:574
			_go_fuzz_dep_.CoverTab[53038]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:574
			// _ = "end of CoverTab[53038]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:575
		// _ = "end of CoverTab[53035]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:576
		_go_fuzz_dep_.CoverTab[53039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:576
		// _ = "end of CoverTab[53039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:576
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:576
	// _ = "end of CoverTab[53022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:576
	_go_fuzz_dep_.CoverTab[53023]++
														xd.L2.Options = xd.L0.ParentFile.builder.optionsUnmarshaler(&descopts.Field, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:577
	// _ = "end of CoverTab[53023]"
}

func (xd *Extension) unmarshalOptions(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:580
	_go_fuzz_dep_.CoverTab[53040]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:581
		_go_fuzz_dep_.CoverTab[53041]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:585
			_go_fuzz_dep_.CoverTab[53042]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.FieldOptions_Packed_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:589
				_go_fuzz_dep_.CoverTab[53044]++
																	xd.L2.IsPacked = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:590
				// _ = "end of CoverTab[53044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:590
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:590
				_go_fuzz_dep_.CoverTab[53045]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:590
				// _ = "end of CoverTab[53045]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:591
			// _ = "end of CoverTab[53042]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:592
			_go_fuzz_dep_.CoverTab[53043]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:594
			// _ = "end of CoverTab[53043]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:595
		// _ = "end of CoverTab[53041]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:596
	// _ = "end of CoverTab[53040]"
}

func (sd *Service) unmarshalFull(b []byte, sb *strs.Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:599
	_go_fuzz_dep_.CoverTab[53046]++
														var rawMethods [][]byte
														var rawOptions []byte
														sd.L2 = new(ServiceL2)
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:603
		_go_fuzz_dep_.CoverTab[53049]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:607
			_go_fuzz_dep_.CoverTab[53050]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.ServiceDescriptorProto_Method_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:611
				_go_fuzz_dep_.CoverTab[53052]++
																	rawMethods = append(rawMethods, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:612
				// _ = "end of CoverTab[53052]"
			case genid.ServiceDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:613
				_go_fuzz_dep_.CoverTab[53053]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:614
				// _ = "end of CoverTab[53053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:614
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:614
				_go_fuzz_dep_.CoverTab[53054]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:614
				// _ = "end of CoverTab[53054]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:615
			// _ = "end of CoverTab[53050]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:616
			_go_fuzz_dep_.CoverTab[53051]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:618
			// _ = "end of CoverTab[53051]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:619
		// _ = "end of CoverTab[53049]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:620
	// _ = "end of CoverTab[53046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:620
	_go_fuzz_dep_.CoverTab[53047]++
														if len(rawMethods) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:621
		_go_fuzz_dep_.CoverTab[53055]++
															sd.L2.Methods.List = make([]Method, len(rawMethods))
															for i, b := range rawMethods {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:623
			_go_fuzz_dep_.CoverTab[53056]++
																sd.L2.Methods.List[i].unmarshalFull(b, sb, sd.L0.ParentFile, sd, i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:624
			// _ = "end of CoverTab[53056]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:625
		// _ = "end of CoverTab[53055]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:626
		_go_fuzz_dep_.CoverTab[53057]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:626
		// _ = "end of CoverTab[53057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:626
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:626
	// _ = "end of CoverTab[53047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:626
	_go_fuzz_dep_.CoverTab[53048]++
														sd.L2.Options = sd.L0.ParentFile.builder.optionsUnmarshaler(&descopts.Service, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:627
	// _ = "end of CoverTab[53048]"
}

func (md *Method) unmarshalFull(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:630
	_go_fuzz_dep_.CoverTab[53058]++
														md.L0.ParentFile = pf
														md.L0.Parent = pd
														md.L0.Index = i

														var rawOptions []byte
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:636
		_go_fuzz_dep_.CoverTab[53060]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:640
			_go_fuzz_dep_.CoverTab[53061]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.MethodDescriptorProto_ClientStreaming_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:644
				_go_fuzz_dep_.CoverTab[53064]++
																	md.L1.IsStreamingClient = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:645
				// _ = "end of CoverTab[53064]"
			case genid.MethodDescriptorProto_ServerStreaming_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:646
				_go_fuzz_dep_.CoverTab[53065]++
																	md.L1.IsStreamingServer = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:647
				// _ = "end of CoverTab[53065]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:647
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:647
				_go_fuzz_dep_.CoverTab[53066]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:647
				// _ = "end of CoverTab[53066]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:648
			// _ = "end of CoverTab[53061]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:649
			_go_fuzz_dep_.CoverTab[53062]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.MethodDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:653
				_go_fuzz_dep_.CoverTab[53067]++
																	md.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:654
				// _ = "end of CoverTab[53067]"
			case genid.MethodDescriptorProto_InputType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:655
				_go_fuzz_dep_.CoverTab[53068]++
																	md.L1.Input = PlaceholderMessage(makeFullName(sb, v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:656
				// _ = "end of CoverTab[53068]"
			case genid.MethodDescriptorProto_OutputType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:657
				_go_fuzz_dep_.CoverTab[53069]++
																	md.L1.Output = PlaceholderMessage(makeFullName(sb, v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:658
				// _ = "end of CoverTab[53069]"
			case genid.MethodDescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:659
				_go_fuzz_dep_.CoverTab[53070]++
																	rawOptions = appendOptions(rawOptions, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:660
				// _ = "end of CoverTab[53070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:660
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:660
				_go_fuzz_dep_.CoverTab[53071]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:660
				// _ = "end of CoverTab[53071]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:661
			// _ = "end of CoverTab[53062]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:662
			_go_fuzz_dep_.CoverTab[53063]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:664
			// _ = "end of CoverTab[53063]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:665
		// _ = "end of CoverTab[53060]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:666
	// _ = "end of CoverTab[53058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:666
	_go_fuzz_dep_.CoverTab[53059]++
														md.L1.Options = pf.builder.optionsUnmarshaler(&descopts.Method, rawOptions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:667
	// _ = "end of CoverTab[53059]"
}

// appendOptions appends src to dst, where the returned slice is never nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:670
// This is necessary to distinguish between empty and unpopulated options.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:672
func appendOptions(dst, src []byte) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:672
	_go_fuzz_dep_.CoverTab[53072]++
														if dst == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:673
		_go_fuzz_dep_.CoverTab[53074]++
															dst = []byte{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:674
		// _ = "end of CoverTab[53074]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:675
		_go_fuzz_dep_.CoverTab[53075]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:675
		// _ = "end of CoverTab[53075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:675
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:675
	// _ = "end of CoverTab[53072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:675
	_go_fuzz_dep_.CoverTab[53073]++
														return append(dst, src...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:676
	// _ = "end of CoverTab[53073]"
}

// optionsUnmarshaler constructs a lazy unmarshal function for an options message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:679
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:679
// The type of message to unmarshal to is passed as a pointer since the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:679
// vars in descopts may not yet be populated at the time this function is called.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:683
func (db *Builder) optionsUnmarshaler(p *protoreflect.ProtoMessage, b []byte) func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:683
	_go_fuzz_dep_.CoverTab[53076]++
														if b == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:684
		_go_fuzz_dep_.CoverTab[53078]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:685
		// _ = "end of CoverTab[53078]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:686
		_go_fuzz_dep_.CoverTab[53079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:686
		// _ = "end of CoverTab[53079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:686
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:686
	// _ = "end of CoverTab[53076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:686
	_go_fuzz_dep_.CoverTab[53077]++
														var opts protoreflect.ProtoMessage
														var once sync.Once
														return func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:689
		_go_fuzz_dep_.CoverTab[53080]++
															once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:690
			_go_fuzz_dep_.CoverTab[53082]++
																if *p == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:691
				_go_fuzz_dep_.CoverTab[53084]++
																	panic("Descriptor.Options called without importing the descriptor package")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:692
				// _ = "end of CoverTab[53084]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:693
				_go_fuzz_dep_.CoverTab[53085]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:693
				// _ = "end of CoverTab[53085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:693
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:693
			// _ = "end of CoverTab[53082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:693
			_go_fuzz_dep_.CoverTab[53083]++
																opts = reflect.New(reflect.TypeOf(*p).Elem()).Interface().(protoreflect.ProtoMessage)
																if err := (proto.UnmarshalOptions{
				AllowPartial:	true,
				Resolver:	db.TypeResolver,
			}).Unmarshal(b, opts); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:698
				_go_fuzz_dep_.CoverTab[53086]++
																	panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:699
				// _ = "end of CoverTab[53086]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:700
				_go_fuzz_dep_.CoverTab[53087]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:700
				// _ = "end of CoverTab[53087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:700
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:700
			// _ = "end of CoverTab[53083]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:701
		// _ = "end of CoverTab[53080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:701
		_go_fuzz_dep_.CoverTab[53081]++
															return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:702
		// _ = "end of CoverTab[53081]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:703
	// _ = "end of CoverTab[53077]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:704
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_lazy.go:704
var _ = _go_fuzz_dep_.CoverTab
