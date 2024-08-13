// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
package filedesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:5
)

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"

	"google.golang.org/protobuf/internal/descfmt"
	"google.golang.org/protobuf/internal/descopts"
	"google.golang.org/protobuf/internal/encoding/defval"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:36
type (
	File	struct {
		fileRaw
		L1	FileL1

		once	uint32		// atomically set if L2 is valid
		mu	sync.Mutex	// protects L2
		L2	*FileL2
	}
	FileL1	struct {
		Syntax	protoreflect.Syntax
		Path	string
		Package	protoreflect.FullName

		Enums		Enums
		Messages	Messages
		Extensions	Extensions
		Services	Services
	}
	FileL2	struct {
		Options		func() protoreflect.ProtoMessage
		Imports		FileImports
		Locations	SourceLocations
	}
)

func (fd *File) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:62
	_go_fuzz_dep_.CoverTab[52372]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:62
	return fd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:62
	// _ = "end of CoverTab[52372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:62
}
func (fd *File) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:63
	_go_fuzz_dep_.CoverTab[52373]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:63
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:63
	// _ = "end of CoverTab[52373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:63
}
func (fd *File) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:64
	_go_fuzz_dep_.CoverTab[52374]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:64
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:64
	// _ = "end of CoverTab[52374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:64
}
func (fd *File) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:65
	_go_fuzz_dep_.CoverTab[52375]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:65
	return fd.L1.Syntax
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:65
	// _ = "end of CoverTab[52375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:65
}
func (fd *File) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:66
	_go_fuzz_dep_.CoverTab[52376]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:66
	return fd.L1.Package.Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:66
	// _ = "end of CoverTab[52376]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:66
}
func (fd *File) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:67
	_go_fuzz_dep_.CoverTab[52377]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:67
	return fd.L1.Package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:67
	// _ = "end of CoverTab[52377]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:67
}
func (fd *File) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:68
	_go_fuzz_dep_.CoverTab[52378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:68
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:68
	// _ = "end of CoverTab[52378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:68
}
func (fd *File) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:69
	_go_fuzz_dep_.CoverTab[52379]++
													if f := fd.lazyInit().Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:70
		_go_fuzz_dep_.CoverTab[52381]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:71
		// _ = "end of CoverTab[52381]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:72
		_go_fuzz_dep_.CoverTab[52382]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:72
		// _ = "end of CoverTab[52382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:72
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:72
	// _ = "end of CoverTab[52379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:72
	_go_fuzz_dep_.CoverTab[52380]++
													return descopts.File
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:73
	// _ = "end of CoverTab[52380]"
}
func (fd *File) Path() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:75
	_go_fuzz_dep_.CoverTab[52383]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:75
	return fd.L1.Path
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:75
	// _ = "end of CoverTab[52383]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:75
}
func (fd *File) Package() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:76
	_go_fuzz_dep_.CoverTab[52384]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:76
	return fd.L1.Package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:76
	// _ = "end of CoverTab[52384]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:76
}
func (fd *File) Imports() protoreflect.FileImports {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:77
	_go_fuzz_dep_.CoverTab[52385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:77
	return &fd.lazyInit().Imports
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:77
	// _ = "end of CoverTab[52385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:77
}
func (fd *File) Enums() protoreflect.EnumDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:78
	_go_fuzz_dep_.CoverTab[52386]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:78
	return &fd.L1.Enums
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:78
	// _ = "end of CoverTab[52386]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:78
}
func (fd *File) Messages() protoreflect.MessageDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:79
	_go_fuzz_dep_.CoverTab[52387]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:79
	return &fd.L1.Messages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:79
	// _ = "end of CoverTab[52387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:79
}
func (fd *File) Extensions() protoreflect.ExtensionDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:80
	_go_fuzz_dep_.CoverTab[52388]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:80
	return &fd.L1.Extensions
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:80
	// _ = "end of CoverTab[52388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:80
}
func (fd *File) Services() protoreflect.ServiceDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:81
	_go_fuzz_dep_.CoverTab[52389]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:81
	return &fd.L1.Services
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:81
	// _ = "end of CoverTab[52389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:81
}
func (fd *File) SourceLocations() protoreflect.SourceLocations {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:82
	_go_fuzz_dep_.CoverTab[52390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:82
	return &fd.lazyInit().Locations
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:82
	// _ = "end of CoverTab[52390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:82
}
func (fd *File) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:83
	_go_fuzz_dep_.CoverTab[52391]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:83
	descfmt.FormatDesc(s, r, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:83
	// _ = "end of CoverTab[52391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:83
}
func (fd *File) ProtoType(protoreflect.FileDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:84
	_go_fuzz_dep_.CoverTab[52392]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:84
	// _ = "end of CoverTab[52392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:84
}
func (fd *File) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:85
	_go_fuzz_dep_.CoverTab[52393]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:85
	// _ = "end of CoverTab[52393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:85
}

func (fd *File) lazyInit() *FileL2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:87
	_go_fuzz_dep_.CoverTab[52394]++
													if atomic.LoadUint32(&fd.once) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:88
		_go_fuzz_dep_.CoverTab[52396]++
														fd.lazyInitOnce()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:89
		// _ = "end of CoverTab[52396]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:90
		_go_fuzz_dep_.CoverTab[52397]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:90
		// _ = "end of CoverTab[52397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:90
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:90
	// _ = "end of CoverTab[52394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:90
	_go_fuzz_dep_.CoverTab[52395]++
													return fd.L2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:91
	// _ = "end of CoverTab[52395]"
}

func (fd *File) lazyInitOnce() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:94
	_go_fuzz_dep_.CoverTab[52398]++
													fd.mu.Lock()
													if fd.L2 == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:96
		_go_fuzz_dep_.CoverTab[52400]++
														fd.lazyRawInit()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:97
		// _ = "end of CoverTab[52400]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:98
		_go_fuzz_dep_.CoverTab[52401]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:98
		// _ = "end of CoverTab[52401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:98
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:98
	// _ = "end of CoverTab[52398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:98
	_go_fuzz_dep_.CoverTab[52399]++
													atomic.StoreUint32(&fd.once, 1)
													fd.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:100
	// _ = "end of CoverTab[52399]"
}

// GoPackagePath is a pseudo-internal API for determining the Go package path
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:103
// that this file descriptor is declared in.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:103
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:103
// WARNING: This method is exempt from the compatibility promise and may be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:103
// removed in the future without warning.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:108
func (fd *File) GoPackagePath() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:108
	_go_fuzz_dep_.CoverTab[52402]++
													return fd.builder.GoPackagePath
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:109
	// _ = "end of CoverTab[52402]"
}

type (
	Enum	struct {
		Base
		L1	EnumL1
		L2	*EnumL2	// protected by fileDesc.once
	}
	EnumL1	struct {
		eagerValues bool	// controls whether EnumL2.Values is already populated
	}
	EnumL2	struct {
		Options		func() protoreflect.ProtoMessage
		Values		EnumValues
		ReservedNames	Names
		ReservedRanges	EnumRanges
	}

	EnumValue	struct {
		Base
		L1	EnumValueL1
	}
	EnumValueL1	struct {
		Options	func() protoreflect.ProtoMessage
		Number	protoreflect.EnumNumber
	}
)

func (ed *Enum) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:138
	_go_fuzz_dep_.CoverTab[52403]++
													if f := ed.lazyInit().Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:139
		_go_fuzz_dep_.CoverTab[52405]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:140
		// _ = "end of CoverTab[52405]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:141
		_go_fuzz_dep_.CoverTab[52406]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:141
		// _ = "end of CoverTab[52406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:141
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:141
	// _ = "end of CoverTab[52403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:141
	_go_fuzz_dep_.CoverTab[52404]++
													return descopts.Enum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:142
	// _ = "end of CoverTab[52404]"
}
func (ed *Enum) Values() protoreflect.EnumValueDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:144
	_go_fuzz_dep_.CoverTab[52407]++
													if ed.L1.eagerValues {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:145
		_go_fuzz_dep_.CoverTab[52409]++
														return &ed.L2.Values
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:146
		// _ = "end of CoverTab[52409]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:147
		_go_fuzz_dep_.CoverTab[52410]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:147
		// _ = "end of CoverTab[52410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:147
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:147
	// _ = "end of CoverTab[52407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:147
	_go_fuzz_dep_.CoverTab[52408]++
													return &ed.lazyInit().Values
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:148
	// _ = "end of CoverTab[52408]"
}
func (ed *Enum) ReservedNames() protoreflect.Names {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:150
	_go_fuzz_dep_.CoverTab[52411]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:150
	return &ed.lazyInit().ReservedNames
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:150
	// _ = "end of CoverTab[52411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:150
}
func (ed *Enum) ReservedRanges() protoreflect.EnumRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:151
	_go_fuzz_dep_.CoverTab[52412]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:151
	return &ed.lazyInit().ReservedRanges
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:151
	// _ = "end of CoverTab[52412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:151
}
func (ed *Enum) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:152
	_go_fuzz_dep_.CoverTab[52413]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:152
	descfmt.FormatDesc(s, r, ed)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:152
	// _ = "end of CoverTab[52413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:152
}
func (ed *Enum) ProtoType(protoreflect.EnumDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:153
	_go_fuzz_dep_.CoverTab[52414]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:153
	// _ = "end of CoverTab[52414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:153
}
func (ed *Enum) lazyInit() *EnumL2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:154
	_go_fuzz_dep_.CoverTab[52415]++
													ed.L0.ParentFile.lazyInit()
													return ed.L2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:156
	// _ = "end of CoverTab[52415]"
}

func (ed *EnumValue) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:159
	_go_fuzz_dep_.CoverTab[52416]++
													if f := ed.L1.Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:160
		_go_fuzz_dep_.CoverTab[52418]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:161
		// _ = "end of CoverTab[52418]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:162
		_go_fuzz_dep_.CoverTab[52419]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:162
		// _ = "end of CoverTab[52419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:162
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:162
	// _ = "end of CoverTab[52416]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:162
	_go_fuzz_dep_.CoverTab[52417]++
													return descopts.EnumValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:163
	// _ = "end of CoverTab[52417]"
}
func (ed *EnumValue) Number() protoreflect.EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:165
	_go_fuzz_dep_.CoverTab[52420]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:165
	return ed.L1.Number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:165
	// _ = "end of CoverTab[52420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:165
}
func (ed *EnumValue) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:166
	_go_fuzz_dep_.CoverTab[52421]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:166
	descfmt.FormatDesc(s, r, ed)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:166
	// _ = "end of CoverTab[52421]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:166
}
func (ed *EnumValue) ProtoType(protoreflect.EnumValueDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:167
	_go_fuzz_dep_.CoverTab[52422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:167
	// _ = "end of CoverTab[52422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:167
}

type (
	Message	struct {
		Base
		L1	MessageL1
		L2	*MessageL2	// protected by fileDesc.once
	}
	MessageL1	struct {
		Enums		Enums
		Messages	Messages
		Extensions	Extensions
		IsMapEntry	bool	// promoted from google.protobuf.MessageOptions
		IsMessageSet	bool	// promoted from google.protobuf.MessageOptions
	}
	MessageL2	struct {
		Options			func() protoreflect.ProtoMessage
		Fields			Fields
		Oneofs			Oneofs
		ReservedNames		Names
		ReservedRanges		FieldRanges
		RequiredNumbers		FieldNumbers	// must be consistent with Fields.Cardinality
		ExtensionRanges		FieldRanges
		ExtensionRangeOptions	[]func() protoreflect.ProtoMessage	// must be same length as ExtensionRanges
	}

	Field	struct {
		Base
		L1	FieldL1
	}
	FieldL1	struct {
		Options			func() protoreflect.ProtoMessage
		Number			protoreflect.FieldNumber
		Cardinality		protoreflect.Cardinality	// must be consistent with Message.RequiredNumbers
		Kind			protoreflect.Kind
		StringName		stringName
		IsProto3Optional	bool	// promoted from google.protobuf.FieldDescriptorProto
		IsWeak			bool	// promoted from google.protobuf.FieldOptions
		HasPacked		bool	// promoted from google.protobuf.FieldOptions
		IsPacked		bool	// promoted from google.protobuf.FieldOptions
		HasEnforceUTF8		bool	// promoted from google.protobuf.FieldOptions
		EnforceUTF8		bool	// promoted from google.protobuf.FieldOptions
		Default			defaultValue
		ContainingOneof		protoreflect.OneofDescriptor	// must be consistent with Message.Oneofs.Fields
		Enum			protoreflect.EnumDescriptor
		Message			protoreflect.MessageDescriptor
	}

	Oneof	struct {
		Base
		L1	OneofL1
	}
	OneofL1	struct {
		Options	func() protoreflect.ProtoMessage
		Fields	OneofFields	// must be consistent with Message.Fields.ContainingOneof
	}
)

func (md *Message) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:225
	_go_fuzz_dep_.CoverTab[52423]++
													if f := md.lazyInit().Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:226
		_go_fuzz_dep_.CoverTab[52425]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:227
		// _ = "end of CoverTab[52425]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:228
		_go_fuzz_dep_.CoverTab[52426]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:228
		// _ = "end of CoverTab[52426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:228
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:228
	// _ = "end of CoverTab[52423]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:228
	_go_fuzz_dep_.CoverTab[52424]++
													return descopts.Message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:229
	// _ = "end of CoverTab[52424]"
}
func (md *Message) IsMapEntry() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:231
	_go_fuzz_dep_.CoverTab[52427]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:231
	return md.L1.IsMapEntry
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:231
	// _ = "end of CoverTab[52427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:231
}
func (md *Message) Fields() protoreflect.FieldDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:232
	_go_fuzz_dep_.CoverTab[52428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:232
	return &md.lazyInit().Fields
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:232
	// _ = "end of CoverTab[52428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:232
}
func (md *Message) Oneofs() protoreflect.OneofDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:233
	_go_fuzz_dep_.CoverTab[52429]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:233
	return &md.lazyInit().Oneofs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:233
	// _ = "end of CoverTab[52429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:233
}
func (md *Message) ReservedNames() protoreflect.Names {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:234
	_go_fuzz_dep_.CoverTab[52430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:234
	return &md.lazyInit().ReservedNames
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:234
	// _ = "end of CoverTab[52430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:234
}
func (md *Message) ReservedRanges() protoreflect.FieldRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:235
	_go_fuzz_dep_.CoverTab[52431]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:235
	return &md.lazyInit().ReservedRanges
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:235
	// _ = "end of CoverTab[52431]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:235
}
func (md *Message) RequiredNumbers() protoreflect.FieldNumbers {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:236
	_go_fuzz_dep_.CoverTab[52432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:236
	return &md.lazyInit().RequiredNumbers
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:236
	// _ = "end of CoverTab[52432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:236
}
func (md *Message) ExtensionRanges() protoreflect.FieldRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:237
	_go_fuzz_dep_.CoverTab[52433]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:237
	return &md.lazyInit().ExtensionRanges
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:237
	// _ = "end of CoverTab[52433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:237
}
func (md *Message) ExtensionRangeOptions(i int) protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:238
	_go_fuzz_dep_.CoverTab[52434]++
													if f := md.lazyInit().ExtensionRangeOptions[i]; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:239
		_go_fuzz_dep_.CoverTab[52436]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:240
		// _ = "end of CoverTab[52436]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:241
		_go_fuzz_dep_.CoverTab[52437]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:241
		// _ = "end of CoverTab[52437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:241
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:241
	// _ = "end of CoverTab[52434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:241
	_go_fuzz_dep_.CoverTab[52435]++
													return descopts.ExtensionRange
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:242
	// _ = "end of CoverTab[52435]"
}
func (md *Message) Enums() protoreflect.EnumDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:244
	_go_fuzz_dep_.CoverTab[52438]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:244
	return &md.L1.Enums
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:244
	// _ = "end of CoverTab[52438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:244
}
func (md *Message) Messages() protoreflect.MessageDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:245
	_go_fuzz_dep_.CoverTab[52439]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:245
	return &md.L1.Messages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:245
	// _ = "end of CoverTab[52439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:245
}
func (md *Message) Extensions() protoreflect.ExtensionDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:246
	_go_fuzz_dep_.CoverTab[52440]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:246
	return &md.L1.Extensions
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:246
	// _ = "end of CoverTab[52440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:246
}
func (md *Message) ProtoType(protoreflect.MessageDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:247
	_go_fuzz_dep_.CoverTab[52441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:247
	// _ = "end of CoverTab[52441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:247
}
func (md *Message) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:248
	_go_fuzz_dep_.CoverTab[52442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:248
	descfmt.FormatDesc(s, r, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:248
	// _ = "end of CoverTab[52442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:248
}
func (md *Message) lazyInit() *MessageL2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:249
	_go_fuzz_dep_.CoverTab[52443]++
													md.L0.ParentFile.lazyInit()
													return md.L2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:251
	// _ = "end of CoverTab[52443]"
}

// IsMessageSet is a pseudo-internal API for checking whether a message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:254
// should serialize in the proto1 message format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:254
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:254
// WARNING: This method is exempt from the compatibility promise and may be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:254
// removed in the future without warning.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:259
func (md *Message) IsMessageSet() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:259
	_go_fuzz_dep_.CoverTab[52444]++
													return md.L1.IsMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:260
	// _ = "end of CoverTab[52444]"
}

func (fd *Field) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:263
	_go_fuzz_dep_.CoverTab[52445]++
													if f := fd.L1.Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:264
		_go_fuzz_dep_.CoverTab[52447]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:265
		// _ = "end of CoverTab[52447]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:266
		_go_fuzz_dep_.CoverTab[52448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:266
		// _ = "end of CoverTab[52448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:266
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:266
	// _ = "end of CoverTab[52445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:266
	_go_fuzz_dep_.CoverTab[52446]++
													return descopts.Field
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:267
	// _ = "end of CoverTab[52446]"
}
func (fd *Field) Number() protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:269
	_go_fuzz_dep_.CoverTab[52449]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:269
	return fd.L1.Number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:269
	// _ = "end of CoverTab[52449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:269
}
func (fd *Field) Cardinality() protoreflect.Cardinality {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:270
	_go_fuzz_dep_.CoverTab[52450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:270
	return fd.L1.Cardinality
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:270
	// _ = "end of CoverTab[52450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:270
}
func (fd *Field) Kind() protoreflect.Kind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:271
	_go_fuzz_dep_.CoverTab[52451]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:271
	return fd.L1.Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:271
	// _ = "end of CoverTab[52451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:271
}
func (fd *Field) HasJSONName() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:272
	_go_fuzz_dep_.CoverTab[52452]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:272
	return fd.L1.StringName.hasJSON
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:272
	// _ = "end of CoverTab[52452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:272
}
func (fd *Field) JSONName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:273
	_go_fuzz_dep_.CoverTab[52453]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:273
	return fd.L1.StringName.getJSON(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:273
	// _ = "end of CoverTab[52453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:273
}
func (fd *Field) TextName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:274
	_go_fuzz_dep_.CoverTab[52454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:274
	return fd.L1.StringName.getText(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:274
	// _ = "end of CoverTab[52454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:274
}
func (fd *Field) HasPresence() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:275
	_go_fuzz_dep_.CoverTab[52455]++
													return fd.L1.Cardinality != protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
		_go_fuzz_dep_.CoverTab[52456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
		return (fd.L0.ParentFile.L1.Syntax == protoreflect.Proto2 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
			_go_fuzz_dep_.CoverTab[52457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
			return fd.L1.Message != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
			// _ = "end of CoverTab[52457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
			_go_fuzz_dep_.CoverTab[52458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
			return fd.L1.ContainingOneof != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
			// _ = "end of CoverTab[52458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
		// _ = "end of CoverTab[52456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:276
	// _ = "end of CoverTab[52455]"
}
func (fd *Field) HasOptionalKeyword() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:278
	_go_fuzz_dep_.CoverTab[52459]++
													return (fd.L0.ParentFile.L1.Syntax == protoreflect.Proto2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		_go_fuzz_dep_.CoverTab[52460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		return fd.L1.Cardinality == protoreflect.Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		// _ = "end of CoverTab[52460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		_go_fuzz_dep_.CoverTab[52461]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		return fd.L1.ContainingOneof == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		// _ = "end of CoverTab[52461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		_go_fuzz_dep_.CoverTab[52462]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		return fd.L1.IsProto3Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
		// _ = "end of CoverTab[52462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:279
	// _ = "end of CoverTab[52459]"
}
func (fd *Field) IsPacked() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:281
	_go_fuzz_dep_.CoverTab[52463]++
													if !fd.L1.HasPacked && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		_go_fuzz_dep_.CoverTab[52465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		return fd.L0.ParentFile.L1.Syntax != protoreflect.Proto2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		// _ = "end of CoverTab[52465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		_go_fuzz_dep_.CoverTab[52466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		return fd.L1.Cardinality == protoreflect.Repeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		// _ = "end of CoverTab[52466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:282
		_go_fuzz_dep_.CoverTab[52467]++
														switch fd.L1.Kind {
		case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:284
			_go_fuzz_dep_.CoverTab[52468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:284
			// _ = "end of CoverTab[52468]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:285
			_go_fuzz_dep_.CoverTab[52469]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:286
			// _ = "end of CoverTab[52469]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:287
		// _ = "end of CoverTab[52467]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:288
		_go_fuzz_dep_.CoverTab[52470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:288
		// _ = "end of CoverTab[52470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:288
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:288
	// _ = "end of CoverTab[52463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:288
	_go_fuzz_dep_.CoverTab[52464]++
													return fd.L1.IsPacked
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:289
	// _ = "end of CoverTab[52464]"
}
func (fd *Field) IsExtension() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:291
	_go_fuzz_dep_.CoverTab[52471]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:291
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:291
	// _ = "end of CoverTab[52471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:291
}
func (fd *Field) IsWeak() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:292
	_go_fuzz_dep_.CoverTab[52472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:292
	return fd.L1.IsWeak
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:292
	// _ = "end of CoverTab[52472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:292
}
func (fd *Field) IsList() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
	_go_fuzz_dep_.CoverTab[52473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
	return fd.Cardinality() == protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
		_go_fuzz_dep_.CoverTab[52474]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
		return !fd.IsMap()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
		// _ = "end of CoverTab[52474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
	// _ = "end of CoverTab[52473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:293
}
func (fd *Field) IsMap() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
	_go_fuzz_dep_.CoverTab[52475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
	return fd.Message() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
		_go_fuzz_dep_.CoverTab[52476]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
		return fd.Message().IsMapEntry()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
		// _ = "end of CoverTab[52476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
	// _ = "end of CoverTab[52475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:294
}
func (fd *Field) MapKey() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:295
	_go_fuzz_dep_.CoverTab[52477]++
													if !fd.IsMap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:296
		_go_fuzz_dep_.CoverTab[52479]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:297
		// _ = "end of CoverTab[52479]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:298
		_go_fuzz_dep_.CoverTab[52480]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:298
		// _ = "end of CoverTab[52480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:298
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:298
	// _ = "end of CoverTab[52477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:298
	_go_fuzz_dep_.CoverTab[52478]++
													return fd.Message().Fields().ByNumber(genid.MapEntry_Key_field_number)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:299
	// _ = "end of CoverTab[52478]"
}
func (fd *Field) MapValue() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:301
	_go_fuzz_dep_.CoverTab[52481]++
													if !fd.IsMap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:302
		_go_fuzz_dep_.CoverTab[52483]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:303
		// _ = "end of CoverTab[52483]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:304
		_go_fuzz_dep_.CoverTab[52484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:304
		// _ = "end of CoverTab[52484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:304
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:304
	// _ = "end of CoverTab[52481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:304
	_go_fuzz_dep_.CoverTab[52482]++
													return fd.Message().Fields().ByNumber(genid.MapEntry_Value_field_number)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:305
	// _ = "end of CoverTab[52482]"
}
func (fd *Field) HasDefault() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:307
	_go_fuzz_dep_.CoverTab[52485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:307
	return fd.L1.Default.has
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:307
	// _ = "end of CoverTab[52485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:307
}
func (fd *Field) Default() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:308
	_go_fuzz_dep_.CoverTab[52486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:308
	return fd.L1.Default.get(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:308
	// _ = "end of CoverTab[52486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:308
}
func (fd *Field) DefaultEnumValue() protoreflect.EnumValueDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:309
	_go_fuzz_dep_.CoverTab[52487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:309
	return fd.L1.Default.enum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:309
	// _ = "end of CoverTab[52487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:309
}
func (fd *Field) ContainingOneof() protoreflect.OneofDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:310
	_go_fuzz_dep_.CoverTab[52488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:310
	return fd.L1.ContainingOneof
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:310
	// _ = "end of CoverTab[52488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:310
}
func (fd *Field) ContainingMessage() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:311
	_go_fuzz_dep_.CoverTab[52489]++
													return fd.L0.Parent.(protoreflect.MessageDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:312
	// _ = "end of CoverTab[52489]"
}
func (fd *Field) Enum() protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:314
	_go_fuzz_dep_.CoverTab[52490]++
													return fd.L1.Enum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:315
	// _ = "end of CoverTab[52490]"
}
func (fd *Field) Message() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:317
	_go_fuzz_dep_.CoverTab[52491]++
													if fd.L1.IsWeak {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:318
		_go_fuzz_dep_.CoverTab[52493]++
														if d, _ := protoregistry.GlobalFiles.FindDescriptorByName(fd.L1.Message.FullName()); d != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:319
			_go_fuzz_dep_.CoverTab[52494]++
															return d.(protoreflect.MessageDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:320
			// _ = "end of CoverTab[52494]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:321
			_go_fuzz_dep_.CoverTab[52495]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:321
			// _ = "end of CoverTab[52495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:321
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:321
		// _ = "end of CoverTab[52493]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:322
		_go_fuzz_dep_.CoverTab[52496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:322
		// _ = "end of CoverTab[52496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:322
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:322
	// _ = "end of CoverTab[52491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:322
	_go_fuzz_dep_.CoverTab[52492]++
													return fd.L1.Message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:323
	// _ = "end of CoverTab[52492]"
}
func (fd *Field) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:325
	_go_fuzz_dep_.CoverTab[52497]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:325
	descfmt.FormatDesc(s, r, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:325
	// _ = "end of CoverTab[52497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:325
}
func (fd *Field) ProtoType(protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:326
	_go_fuzz_dep_.CoverTab[52498]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:326
	// _ = "end of CoverTab[52498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:326
}

// EnforceUTF8 is a pseudo-internal API to determine whether to enforce UTF-8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:328
// validation for the string field. This exists for Google-internal use only
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:328
// since proto3 did not enforce UTF-8 validity prior to the open-source release.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:328
// If this method does not exist, the default is to enforce valid UTF-8.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:328
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:328
// WARNING: This method is exempt from the compatibility promise and may be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:328
// removed in the future without warning.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:335
func (fd *Field) EnforceUTF8() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:335
	_go_fuzz_dep_.CoverTab[52499]++
													if fd.L1.HasEnforceUTF8 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:336
		_go_fuzz_dep_.CoverTab[52501]++
														return fd.L1.EnforceUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:337
		// _ = "end of CoverTab[52501]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:338
		_go_fuzz_dep_.CoverTab[52502]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:338
		// _ = "end of CoverTab[52502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:338
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:338
	// _ = "end of CoverTab[52499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:338
	_go_fuzz_dep_.CoverTab[52500]++
													return fd.L0.ParentFile.L1.Syntax == protoreflect.Proto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:339
	// _ = "end of CoverTab[52500]"
}

func (od *Oneof) IsSynthetic() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:342
	_go_fuzz_dep_.CoverTab[52503]++
													return od.L0.ParentFile.L1.Syntax == protoreflect.Proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
		_go_fuzz_dep_.CoverTab[52504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
		return len(od.L1.Fields.List) == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
		// _ = "end of CoverTab[52504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
		_go_fuzz_dep_.CoverTab[52505]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
		return od.L1.Fields.List[0].HasOptionalKeyword()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
		// _ = "end of CoverTab[52505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:343
	// _ = "end of CoverTab[52503]"
}
func (od *Oneof) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:345
	_go_fuzz_dep_.CoverTab[52506]++
													if f := od.L1.Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:346
		_go_fuzz_dep_.CoverTab[52508]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:347
		// _ = "end of CoverTab[52508]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:348
		_go_fuzz_dep_.CoverTab[52509]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:348
		// _ = "end of CoverTab[52509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:348
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:348
	// _ = "end of CoverTab[52506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:348
	_go_fuzz_dep_.CoverTab[52507]++
													return descopts.Oneof
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:349
	// _ = "end of CoverTab[52507]"
}
func (od *Oneof) Fields() protoreflect.FieldDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:351
	_go_fuzz_dep_.CoverTab[52510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:351
	return &od.L1.Fields
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:351
	// _ = "end of CoverTab[52510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:351
}
func (od *Oneof) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:352
	_go_fuzz_dep_.CoverTab[52511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:352
	descfmt.FormatDesc(s, r, od)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:352
	// _ = "end of CoverTab[52511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:352
}
func (od *Oneof) ProtoType(protoreflect.OneofDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:353
	_go_fuzz_dep_.CoverTab[52512]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:353
	// _ = "end of CoverTab[52512]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:353
}

type (
	Extension	struct {
		Base
		L1	ExtensionL1
		L2	*ExtensionL2	// protected by fileDesc.once
	}
	ExtensionL1	struct {
		Number		protoreflect.FieldNumber
		Extendee	protoreflect.MessageDescriptor
		Cardinality	protoreflect.Cardinality
		Kind		protoreflect.Kind
	}
	ExtensionL2	struct {
		Options			func() protoreflect.ProtoMessage
		StringName		stringName
		IsProto3Optional	bool	// promoted from google.protobuf.FieldDescriptorProto
		IsPacked		bool	// promoted from google.protobuf.FieldOptions
		Default			defaultValue
		Enum			protoreflect.EnumDescriptor
		Message			protoreflect.MessageDescriptor
	}
)

func (xd *Extension) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:378
	_go_fuzz_dep_.CoverTab[52513]++
													if f := xd.lazyInit().Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:379
		_go_fuzz_dep_.CoverTab[52515]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:380
		// _ = "end of CoverTab[52515]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:381
		_go_fuzz_dep_.CoverTab[52516]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:381
		// _ = "end of CoverTab[52516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:381
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:381
	// _ = "end of CoverTab[52513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:381
	_go_fuzz_dep_.CoverTab[52514]++
													return descopts.Field
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:382
	// _ = "end of CoverTab[52514]"
}
func (xd *Extension) Number() protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:384
	_go_fuzz_dep_.CoverTab[52517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:384
	return xd.L1.Number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:384
	// _ = "end of CoverTab[52517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:384
}
func (xd *Extension) Cardinality() protoreflect.Cardinality {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:385
	_go_fuzz_dep_.CoverTab[52518]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:385
	return xd.L1.Cardinality
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:385
	// _ = "end of CoverTab[52518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:385
}
func (xd *Extension) Kind() protoreflect.Kind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:386
	_go_fuzz_dep_.CoverTab[52519]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:386
	return xd.L1.Kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:386
	// _ = "end of CoverTab[52519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:386
}
func (xd *Extension) HasJSONName() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:387
	_go_fuzz_dep_.CoverTab[52520]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:387
	return xd.lazyInit().StringName.hasJSON
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:387
	// _ = "end of CoverTab[52520]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:387
}
func (xd *Extension) JSONName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:388
	_go_fuzz_dep_.CoverTab[52521]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:388
	return xd.lazyInit().StringName.getJSON(xd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:388
	// _ = "end of CoverTab[52521]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:388
}
func (xd *Extension) TextName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:389
	_go_fuzz_dep_.CoverTab[52522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:389
	return xd.lazyInit().StringName.getText(xd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:389
	// _ = "end of CoverTab[52522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:389
}
func (xd *Extension) HasPresence() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:390
	_go_fuzz_dep_.CoverTab[52523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:390
	return xd.L1.Cardinality != protoreflect.Repeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:390
	// _ = "end of CoverTab[52523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:390
}
func (xd *Extension) HasOptionalKeyword() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:391
	_go_fuzz_dep_.CoverTab[52524]++
													return (xd.L0.ParentFile.L1.Syntax == protoreflect.Proto2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
		_go_fuzz_dep_.CoverTab[52525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
		return xd.L1.Cardinality == protoreflect.Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
		// _ = "end of CoverTab[52525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
		_go_fuzz_dep_.CoverTab[52526]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
		return xd.lazyInit().IsProto3Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
		// _ = "end of CoverTab[52526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:392
	// _ = "end of CoverTab[52524]"
}
func (xd *Extension) IsPacked() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:394
	_go_fuzz_dep_.CoverTab[52527]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:394
	return xd.lazyInit().IsPacked
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:394
	// _ = "end of CoverTab[52527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:394
}
func (xd *Extension) IsExtension() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:395
	_go_fuzz_dep_.CoverTab[52528]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:395
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:395
	// _ = "end of CoverTab[52528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:395
}
func (xd *Extension) IsWeak() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:396
	_go_fuzz_dep_.CoverTab[52529]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:396
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:396
	// _ = "end of CoverTab[52529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:396
}
func (xd *Extension) IsList() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:397
	_go_fuzz_dep_.CoverTab[52530]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:397
	return xd.Cardinality() == protoreflect.Repeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:397
	// _ = "end of CoverTab[52530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:397
}
func (xd *Extension) IsMap() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:398
	_go_fuzz_dep_.CoverTab[52531]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:398
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:398
	// _ = "end of CoverTab[52531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:398
}
func (xd *Extension) MapKey() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:399
	_go_fuzz_dep_.CoverTab[52532]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:399
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:399
	// _ = "end of CoverTab[52532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:399
}
func (xd *Extension) MapValue() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:400
	_go_fuzz_dep_.CoverTab[52533]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:400
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:400
	// _ = "end of CoverTab[52533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:400
}
func (xd *Extension) HasDefault() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:401
	_go_fuzz_dep_.CoverTab[52534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:401
	return xd.lazyInit().Default.has
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:401
	// _ = "end of CoverTab[52534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:401
}
func (xd *Extension) Default() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:402
	_go_fuzz_dep_.CoverTab[52535]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:402
	return xd.lazyInit().Default.get(xd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:402
	// _ = "end of CoverTab[52535]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:402
}
func (xd *Extension) DefaultEnumValue() protoreflect.EnumValueDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:403
	_go_fuzz_dep_.CoverTab[52536]++
													return xd.lazyInit().Default.enum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:404
	// _ = "end of CoverTab[52536]"
}
func (xd *Extension) ContainingOneof() protoreflect.OneofDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:406
	_go_fuzz_dep_.CoverTab[52537]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:406
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:406
	// _ = "end of CoverTab[52537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:406
}
func (xd *Extension) ContainingMessage() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:407
	_go_fuzz_dep_.CoverTab[52538]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:407
	return xd.L1.Extendee
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:407
	// _ = "end of CoverTab[52538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:407
}
func (xd *Extension) Enum() protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:408
	_go_fuzz_dep_.CoverTab[52539]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:408
	return xd.lazyInit().Enum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:408
	// _ = "end of CoverTab[52539]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:408
}
func (xd *Extension) Message() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:409
	_go_fuzz_dep_.CoverTab[52540]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:409
	return xd.lazyInit().Message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:409
	// _ = "end of CoverTab[52540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:409
}
func (xd *Extension) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:410
	_go_fuzz_dep_.CoverTab[52541]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:410
	descfmt.FormatDesc(s, r, xd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:410
	// _ = "end of CoverTab[52541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:410
}
func (xd *Extension) ProtoType(protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:411
	_go_fuzz_dep_.CoverTab[52542]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:411
	// _ = "end of CoverTab[52542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:411
}
func (xd *Extension) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:412
	_go_fuzz_dep_.CoverTab[52543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:412
	// _ = "end of CoverTab[52543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:412
}
func (xd *Extension) lazyInit() *ExtensionL2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:413
	_go_fuzz_dep_.CoverTab[52544]++
													xd.L0.ParentFile.lazyInit()
													return xd.L2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:415
	// _ = "end of CoverTab[52544]"
}

type (
	Service	struct {
		Base
		L1	ServiceL1
		L2	*ServiceL2	// protected by fileDesc.once
	}
	ServiceL1	struct{}
	ServiceL2	struct {
		Options	func() protoreflect.ProtoMessage
		Methods	Methods
	}

	Method	struct {
		Base
		L1	MethodL1
	}
	MethodL1	struct {
		Options			func() protoreflect.ProtoMessage
		Input			protoreflect.MessageDescriptor
		Output			protoreflect.MessageDescriptor
		IsStreamingClient	bool
		IsStreamingServer	bool
	}
)

func (sd *Service) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:443
	_go_fuzz_dep_.CoverTab[52545]++
													if f := sd.lazyInit().Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:444
		_go_fuzz_dep_.CoverTab[52547]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:445
		// _ = "end of CoverTab[52547]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:446
		_go_fuzz_dep_.CoverTab[52548]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:446
		// _ = "end of CoverTab[52548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:446
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:446
	// _ = "end of CoverTab[52545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:446
	_go_fuzz_dep_.CoverTab[52546]++
													return descopts.Service
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:447
	// _ = "end of CoverTab[52546]"
}
func (sd *Service) Methods() protoreflect.MethodDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:449
	_go_fuzz_dep_.CoverTab[52549]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:449
	return &sd.lazyInit().Methods
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:449
	// _ = "end of CoverTab[52549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:449
}
func (sd *Service) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:450
	_go_fuzz_dep_.CoverTab[52550]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:450
	descfmt.FormatDesc(s, r, sd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:450
	// _ = "end of CoverTab[52550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:450
}
func (sd *Service) ProtoType(protoreflect.ServiceDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:451
	_go_fuzz_dep_.CoverTab[52551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:451
	// _ = "end of CoverTab[52551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:451
}
func (sd *Service) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:452
	_go_fuzz_dep_.CoverTab[52552]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:452
	// _ = "end of CoverTab[52552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:452
}
func (sd *Service) lazyInit() *ServiceL2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:453
	_go_fuzz_dep_.CoverTab[52553]++
													sd.L0.ParentFile.lazyInit()
													return sd.L2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:455
	// _ = "end of CoverTab[52553]"
}

func (md *Method) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:458
	_go_fuzz_dep_.CoverTab[52554]++
													if f := md.L1.Options; f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:459
		_go_fuzz_dep_.CoverTab[52556]++
														return f()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:460
		// _ = "end of CoverTab[52556]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:461
		_go_fuzz_dep_.CoverTab[52557]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:461
		// _ = "end of CoverTab[52557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:461
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:461
	// _ = "end of CoverTab[52554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:461
	_go_fuzz_dep_.CoverTab[52555]++
													return descopts.Method
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:462
	// _ = "end of CoverTab[52555]"
}
func (md *Method) Input() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:464
	_go_fuzz_dep_.CoverTab[52558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:464
	return md.L1.Input
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:464
	// _ = "end of CoverTab[52558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:464
}
func (md *Method) Output() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:465
	_go_fuzz_dep_.CoverTab[52559]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:465
	return md.L1.Output
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:465
	// _ = "end of CoverTab[52559]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:465
}
func (md *Method) IsStreamingClient() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:466
	_go_fuzz_dep_.CoverTab[52560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:466
	return md.L1.IsStreamingClient
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:466
	// _ = "end of CoverTab[52560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:466
}
func (md *Method) IsStreamingServer() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:467
	_go_fuzz_dep_.CoverTab[52561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:467
	return md.L1.IsStreamingServer
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:467
	// _ = "end of CoverTab[52561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:467
}
func (md *Method) Format(s fmt.State, r rune) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:468
	_go_fuzz_dep_.CoverTab[52562]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:468
	descfmt.FormatDesc(s, r, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:468
	// _ = "end of CoverTab[52562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:468
}
func (md *Method) ProtoType(protoreflect.MethodDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:469
	_go_fuzz_dep_.CoverTab[52563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:469
	// _ = "end of CoverTab[52563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:469
}
func (md *Method) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:470
	_go_fuzz_dep_.CoverTab[52564]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:470
	// _ = "end of CoverTab[52564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:470
}

// Surrogate files are can be used to create standalone descriptors
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:472
// where the syntax is only information derived from the parent file.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:474
var (
	SurrogateProto2	= &File{L1: FileL1{Syntax: protoreflect.Proto2}, L2: &FileL2{}}
	SurrogateProto3	= &File{L1: FileL1{Syntax: protoreflect.Proto3}, L2: &FileL2{}}
)

type (
	Base	struct {
		L0 BaseL0
	}
	BaseL0	struct {
		FullName	protoreflect.FullName	// must be populated
		ParentFile	*File			// must be populated
		Parent		protoreflect.Descriptor
		Index		int
	}
)

func (d *Base) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:491
	_go_fuzz_dep_.CoverTab[52565]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:491
	return d.L0.FullName.Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:491
	// _ = "end of CoverTab[52565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:491
}
func (d *Base) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:492
	_go_fuzz_dep_.CoverTab[52566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:492
	return d.L0.FullName
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:492
	// _ = "end of CoverTab[52566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:492
}
func (d *Base) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:493
	_go_fuzz_dep_.CoverTab[52567]++
													if d.L0.ParentFile == SurrogateProto2 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:494
		_go_fuzz_dep_.CoverTab[52569]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:494
		return d.L0.ParentFile == SurrogateProto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:494
		// _ = "end of CoverTab[52569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:494
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:494
		_go_fuzz_dep_.CoverTab[52570]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:495
		// _ = "end of CoverTab[52570]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:496
		_go_fuzz_dep_.CoverTab[52571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:496
		// _ = "end of CoverTab[52571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:496
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:496
	// _ = "end of CoverTab[52567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:496
	_go_fuzz_dep_.CoverTab[52568]++
													return d.L0.ParentFile
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:497
	// _ = "end of CoverTab[52568]"
}
func (d *Base) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:499
	_go_fuzz_dep_.CoverTab[52572]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:499
	return d.L0.Parent
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:499
	// _ = "end of CoverTab[52572]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:499
}
func (d *Base) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:500
	_go_fuzz_dep_.CoverTab[52573]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:500
	return d.L0.Index
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:500
	// _ = "end of CoverTab[52573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:500
}
func (d *Base) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:501
	_go_fuzz_dep_.CoverTab[52574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:501
	return d.L0.ParentFile.Syntax()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:501
	// _ = "end of CoverTab[52574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:501
}
func (d *Base) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:502
	_go_fuzz_dep_.CoverTab[52575]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:502
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:502
	// _ = "end of CoverTab[52575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:502
}
func (d *Base) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:503
	_go_fuzz_dep_.CoverTab[52576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:503
	// _ = "end of CoverTab[52576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:503
}

type stringName struct {
	hasJSON		bool
	once		sync.Once
	nameJSON	string
	nameText	string
}

// InitJSON initializes the name. It is exported for use by other internal packages.
func (s *stringName) InitJSON(name string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:513
	_go_fuzz_dep_.CoverTab[52577]++
													s.hasJSON = true
													s.nameJSON = name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:515
	// _ = "end of CoverTab[52577]"
}

func (s *stringName) lazyInit(fd protoreflect.FieldDescriptor) *stringName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:518
	_go_fuzz_dep_.CoverTab[52578]++
													s.once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:519
		_go_fuzz_dep_.CoverTab[52580]++
														if fd.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:520
			_go_fuzz_dep_.CoverTab[52581]++
			// For extensions, JSON and text are formatted the same way.
			var name string
			if messageset.IsMessageSetExtension(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:523
				_go_fuzz_dep_.CoverTab[52583]++
																name = string("[" + fd.FullName().Parent() + "]")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:524
				// _ = "end of CoverTab[52583]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:525
				_go_fuzz_dep_.CoverTab[52584]++
																name = string("[" + fd.FullName() + "]")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:526
				// _ = "end of CoverTab[52584]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:527
			// _ = "end of CoverTab[52581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:527
			_go_fuzz_dep_.CoverTab[52582]++
															s.nameJSON = name
															s.nameText = name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:529
			// _ = "end of CoverTab[52582]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:530
			_go_fuzz_dep_.CoverTab[52585]++

															if !s.hasJSON {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:532
				_go_fuzz_dep_.CoverTab[52587]++
																s.nameJSON = strs.JSONCamelCase(string(fd.Name()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:533
				// _ = "end of CoverTab[52587]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:534
				_go_fuzz_dep_.CoverTab[52588]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:534
				// _ = "end of CoverTab[52588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:534
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:534
			// _ = "end of CoverTab[52585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:534
			_go_fuzz_dep_.CoverTab[52586]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:537
			s.nameText = string(fd.Name())
			if fd.Kind() == protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:538
				_go_fuzz_dep_.CoverTab[52589]++
																s.nameText = string(fd.Message().Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:539
				// _ = "end of CoverTab[52589]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:540
				_go_fuzz_dep_.CoverTab[52590]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:540
				// _ = "end of CoverTab[52590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:540
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:540
			// _ = "end of CoverTab[52586]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:541
		// _ = "end of CoverTab[52580]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:542
	// _ = "end of CoverTab[52578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:542
	_go_fuzz_dep_.CoverTab[52579]++
													return s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:543
	// _ = "end of CoverTab[52579]"
}

func (s *stringName) getJSON(fd protoreflect.FieldDescriptor) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:546
	_go_fuzz_dep_.CoverTab[52591]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:546
	return s.lazyInit(fd).nameJSON
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:546
	// _ = "end of CoverTab[52591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:546
}
func (s *stringName) getText(fd protoreflect.FieldDescriptor) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:547
	_go_fuzz_dep_.CoverTab[52592]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:547
	return s.lazyInit(fd).nameText
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:547
	// _ = "end of CoverTab[52592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:547
}

func DefaultValue(v protoreflect.Value, ev protoreflect.EnumValueDescriptor) defaultValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:549
	_go_fuzz_dep_.CoverTab[52593]++
													dv := defaultValue{has: v.IsValid(), val: v, enum: ev}
													if b, ok := v.Interface().([]byte); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:551
		_go_fuzz_dep_.CoverTab[52595]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:554
		dv.bytes = append([]byte(nil), b...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:554
		// _ = "end of CoverTab[52595]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:555
		_go_fuzz_dep_.CoverTab[52596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:555
		// _ = "end of CoverTab[52596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:555
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:555
	// _ = "end of CoverTab[52593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:555
	_go_fuzz_dep_.CoverTab[52594]++
													return dv
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:556
	// _ = "end of CoverTab[52594]"
}

func unmarshalDefault(b []byte, k protoreflect.Kind, pf *File, ed protoreflect.EnumDescriptor) defaultValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:559
	_go_fuzz_dep_.CoverTab[52597]++
													var evs protoreflect.EnumValueDescriptors
													if k == protoreflect.EnumKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:561
		_go_fuzz_dep_.CoverTab[52600]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:564
		if e, ok := ed.(*Enum); ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:564
			_go_fuzz_dep_.CoverTab[52602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:564
			return e.L0.ParentFile == pf
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:564
			// _ = "end of CoverTab[52602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:564
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:564
			_go_fuzz_dep_.CoverTab[52603]++
															evs = &e.L2.Values
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:565
			// _ = "end of CoverTab[52603]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:566
			_go_fuzz_dep_.CoverTab[52604]++
															evs = ed.Values()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:567
			// _ = "end of CoverTab[52604]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:568
		// _ = "end of CoverTab[52600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:568
		_go_fuzz_dep_.CoverTab[52601]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:572
		if ed.IsPlaceholder() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:572
			_go_fuzz_dep_.CoverTab[52605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:572
			return protoreflect.Name(b).IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:572
			// _ = "end of CoverTab[52605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:572
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:572
			_go_fuzz_dep_.CoverTab[52606]++
															v := protoreflect.ValueOfEnum(0)
															ev := PlaceholderEnumValue(ed.FullName().Parent().Append(protoreflect.Name(b)))
															return DefaultValue(v, ev)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:575
			// _ = "end of CoverTab[52606]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:576
			_go_fuzz_dep_.CoverTab[52607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:576
			// _ = "end of CoverTab[52607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:576
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:576
		// _ = "end of CoverTab[52601]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:577
		_go_fuzz_dep_.CoverTab[52608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:577
		// _ = "end of CoverTab[52608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:577
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:577
	// _ = "end of CoverTab[52597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:577
	_go_fuzz_dep_.CoverTab[52598]++

													v, ev, err := defval.Unmarshal(string(b), k, evs, defval.Descriptor)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:580
		_go_fuzz_dep_.CoverTab[52609]++
														panic(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:581
		// _ = "end of CoverTab[52609]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:582
		_go_fuzz_dep_.CoverTab[52610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:582
		// _ = "end of CoverTab[52610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:582
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:582
	// _ = "end of CoverTab[52598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:582
	_go_fuzz_dep_.CoverTab[52599]++
													return DefaultValue(v, ev)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:583
	// _ = "end of CoverTab[52599]"
}

type defaultValue struct {
	has	bool
	val	protoreflect.Value
	enum	protoreflect.EnumValueDescriptor
	bytes	[]byte
}

func (dv *defaultValue) get(fd protoreflect.FieldDescriptor) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:593
	_go_fuzz_dep_.CoverTab[52611]++

													if !dv.has {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:595
		_go_fuzz_dep_.CoverTab[52614]++
														if fd.Cardinality() == protoreflect.Repeated {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:596
			_go_fuzz_dep_.CoverTab[52616]++
															return protoreflect.Value{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:597
			// _ = "end of CoverTab[52616]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:598
			_go_fuzz_dep_.CoverTab[52617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:598
			// _ = "end of CoverTab[52617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:598
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:598
		// _ = "end of CoverTab[52614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:598
		_go_fuzz_dep_.CoverTab[52615]++
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:600
			_go_fuzz_dep_.CoverTab[52618]++
															return protoreflect.ValueOfBool(false)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:601
			// _ = "end of CoverTab[52618]"
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:602
			_go_fuzz_dep_.CoverTab[52619]++
															return protoreflect.ValueOfInt32(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:603
			// _ = "end of CoverTab[52619]"
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:604
			_go_fuzz_dep_.CoverTab[52620]++
															return protoreflect.ValueOfInt64(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:605
			// _ = "end of CoverTab[52620]"
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:606
			_go_fuzz_dep_.CoverTab[52621]++
															return protoreflect.ValueOfUint32(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:607
			// _ = "end of CoverTab[52621]"
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:608
			_go_fuzz_dep_.CoverTab[52622]++
															return protoreflect.ValueOfUint64(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:609
			// _ = "end of CoverTab[52622]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:610
			_go_fuzz_dep_.CoverTab[52623]++
															return protoreflect.ValueOfFloat32(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:611
			// _ = "end of CoverTab[52623]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:612
			_go_fuzz_dep_.CoverTab[52624]++
															return protoreflect.ValueOfFloat64(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:613
			// _ = "end of CoverTab[52624]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:614
			_go_fuzz_dep_.CoverTab[52625]++
															return protoreflect.ValueOfString("")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:615
			// _ = "end of CoverTab[52625]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:616
			_go_fuzz_dep_.CoverTab[52626]++
															return protoreflect.ValueOfBytes(nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:617
			// _ = "end of CoverTab[52626]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:618
			_go_fuzz_dep_.CoverTab[52627]++
															if evs := fd.Enum().Values(); evs.Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:619
				_go_fuzz_dep_.CoverTab[52630]++
																return protoreflect.ValueOfEnum(evs.Get(0).Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:620
				// _ = "end of CoverTab[52630]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:621
				_go_fuzz_dep_.CoverTab[52631]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:621
				// _ = "end of CoverTab[52631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:621
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:621
			// _ = "end of CoverTab[52627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:621
			_go_fuzz_dep_.CoverTab[52628]++
															return protoreflect.ValueOfEnum(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:622
			// _ = "end of CoverTab[52628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:622
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:622
			_go_fuzz_dep_.CoverTab[52629]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:622
			// _ = "end of CoverTab[52629]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:623
		// _ = "end of CoverTab[52615]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:624
		_go_fuzz_dep_.CoverTab[52632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:624
		// _ = "end of CoverTab[52632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:624
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:624
	// _ = "end of CoverTab[52611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:624
	_go_fuzz_dep_.CoverTab[52612]++

													if len(dv.bytes) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:626
		_go_fuzz_dep_.CoverTab[52633]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:626
		return !bytes.Equal(dv.bytes, dv.val.Bytes())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:626
		// _ = "end of CoverTab[52633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:626
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:626
		_go_fuzz_dep_.CoverTab[52634]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:630
		panic(fmt.Sprintf("detected mutation on the default bytes for %v", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:630
		// _ = "end of CoverTab[52634]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:631
		_go_fuzz_dep_.CoverTab[52635]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:631
		// _ = "end of CoverTab[52635]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:631
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:631
	// _ = "end of CoverTab[52612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:631
	_go_fuzz_dep_.CoverTab[52613]++
													return dv.val
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:632
	// _ = "end of CoverTab[52613]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:633
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc.go:633
var _ = _go_fuzz_dep_.CoverTab
