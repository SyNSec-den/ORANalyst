// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
package filedesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:5
)

import (
	"google.golang.org/protobuf/internal/descopts"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	emptyNames		= new(Names)
	emptyEnumRanges		= new(EnumRanges)
	emptyFieldRanges	= new(FieldRanges)
	emptyFieldNumbers	= new(FieldNumbers)
	emptySourceLocations	= new(SourceLocations)

	emptyFiles	= new(FileImports)
	emptyMessages	= new(Messages)
	emptyFields	= new(Fields)
	emptyOneofs	= new(Oneofs)
	emptyEnums	= new(Enums)
	emptyEnumValues	= new(EnumValues)
	emptyExtensions	= new(Extensions)
	emptyServices	= new(Services)
)

// PlaceholderFile is a placeholder, representing only the file path.
type PlaceholderFile string

func (f PlaceholderFile) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:33
	_go_fuzz_dep_.CoverTab[53428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:33
	return f
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:33
	// _ = "end of CoverTab[53428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:33
}
func (f PlaceholderFile) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:34
	_go_fuzz_dep_.CoverTab[53429]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:34
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:34
	// _ = "end of CoverTab[53429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:34
}
func (f PlaceholderFile) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:35
	_go_fuzz_dep_.CoverTab[53430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:35
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:35
	// _ = "end of CoverTab[53430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:35
}
func (f PlaceholderFile) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:36
	_go_fuzz_dep_.CoverTab[53431]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:36
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:36
	// _ = "end of CoverTab[53431]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:36
}
func (f PlaceholderFile) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:37
	_go_fuzz_dep_.CoverTab[53432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:37
	return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:37
	// _ = "end of CoverTab[53432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:37
}
func (f PlaceholderFile) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:38
	_go_fuzz_dep_.CoverTab[53433]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:38
	return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:38
	// _ = "end of CoverTab[53433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:38
}
func (f PlaceholderFile) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:39
	_go_fuzz_dep_.CoverTab[53434]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:39
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:39
	// _ = "end of CoverTab[53434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:39
}
func (f PlaceholderFile) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:40
	_go_fuzz_dep_.CoverTab[53435]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:40
	return descopts.File
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:40
	// _ = "end of CoverTab[53435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:40
}
func (f PlaceholderFile) Path() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:41
	_go_fuzz_dep_.CoverTab[53436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:41
	return string(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:41
	// _ = "end of CoverTab[53436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:41
}
func (f PlaceholderFile) Package() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:42
	_go_fuzz_dep_.CoverTab[53437]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:42
	return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:42
	// _ = "end of CoverTab[53437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:42
}
func (f PlaceholderFile) Imports() protoreflect.FileImports {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:43
	_go_fuzz_dep_.CoverTab[53438]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:43
	return emptyFiles
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:43
	// _ = "end of CoverTab[53438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:43
}
func (f PlaceholderFile) Messages() protoreflect.MessageDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:44
	_go_fuzz_dep_.CoverTab[53439]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:44
	return emptyMessages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:44
	// _ = "end of CoverTab[53439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:44
}
func (f PlaceholderFile) Enums() protoreflect.EnumDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:45
	_go_fuzz_dep_.CoverTab[53440]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:45
	return emptyEnums
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:45
	// _ = "end of CoverTab[53440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:45
}
func (f PlaceholderFile) Extensions() protoreflect.ExtensionDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:46
	_go_fuzz_dep_.CoverTab[53441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:46
	return emptyExtensions
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:46
	// _ = "end of CoverTab[53441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:46
}
func (f PlaceholderFile) Services() protoreflect.ServiceDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:47
	_go_fuzz_dep_.CoverTab[53442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:47
	return emptyServices
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:47
	// _ = "end of CoverTab[53442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:47
}
func (f PlaceholderFile) SourceLocations() protoreflect.SourceLocations {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:48
	_go_fuzz_dep_.CoverTab[53443]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:48
	return emptySourceLocations
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:48
	// _ = "end of CoverTab[53443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:48
}
func (f PlaceholderFile) ProtoType(protoreflect.FileDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:49
	_go_fuzz_dep_.CoverTab[53444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:49
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:49
	// _ = "end of CoverTab[53444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:49
}
func (f PlaceholderFile) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:50
	_go_fuzz_dep_.CoverTab[53445]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:50
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:50
	// _ = "end of CoverTab[53445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:50
}

// PlaceholderEnum is a placeholder, representing only the full name.
type PlaceholderEnum protoreflect.FullName

func (e PlaceholderEnum) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:55
	_go_fuzz_dep_.CoverTab[53446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:55
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:55
	// _ = "end of CoverTab[53446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:55
}
func (e PlaceholderEnum) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:56
	_go_fuzz_dep_.CoverTab[53447]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:56
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:56
	// _ = "end of CoverTab[53447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:56
}
func (e PlaceholderEnum) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:57
	_go_fuzz_dep_.CoverTab[53448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:57
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:57
	// _ = "end of CoverTab[53448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:57
}
func (e PlaceholderEnum) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:58
	_go_fuzz_dep_.CoverTab[53449]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:58
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:58
	// _ = "end of CoverTab[53449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:58
}
func (e PlaceholderEnum) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:59
	_go_fuzz_dep_.CoverTab[53450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:59
	return protoreflect.FullName(e).Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:59
	// _ = "end of CoverTab[53450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:59
}
func (e PlaceholderEnum) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:60
	_go_fuzz_dep_.CoverTab[53451]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:60
	return protoreflect.FullName(e)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:60
	// _ = "end of CoverTab[53451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:60
}
func (e PlaceholderEnum) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:61
	_go_fuzz_dep_.CoverTab[53452]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:61
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:61
	// _ = "end of CoverTab[53452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:61
}
func (e PlaceholderEnum) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:62
	_go_fuzz_dep_.CoverTab[53453]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:62
	return descopts.Enum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:62
	// _ = "end of CoverTab[53453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:62
}
func (e PlaceholderEnum) Values() protoreflect.EnumValueDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:63
	_go_fuzz_dep_.CoverTab[53454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:63
	return emptyEnumValues
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:63
	// _ = "end of CoverTab[53454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:63
}
func (e PlaceholderEnum) ReservedNames() protoreflect.Names {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:64
	_go_fuzz_dep_.CoverTab[53455]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:64
	return emptyNames
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:64
	// _ = "end of CoverTab[53455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:64
}
func (e PlaceholderEnum) ReservedRanges() protoreflect.EnumRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:65
	_go_fuzz_dep_.CoverTab[53456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:65
	return emptyEnumRanges
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:65
	// _ = "end of CoverTab[53456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:65
}
func (e PlaceholderEnum) ProtoType(protoreflect.EnumDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:66
	_go_fuzz_dep_.CoverTab[53457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:66
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:66
	// _ = "end of CoverTab[53457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:66
}
func (e PlaceholderEnum) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:67
	_go_fuzz_dep_.CoverTab[53458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:67
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:67
	// _ = "end of CoverTab[53458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:67
}

// PlaceholderEnumValue is a placeholder, representing only the full name.
type PlaceholderEnumValue protoreflect.FullName

func (e PlaceholderEnumValue) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:72
	_go_fuzz_dep_.CoverTab[53459]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:72
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:72
	// _ = "end of CoverTab[53459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:72
}
func (e PlaceholderEnumValue) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:73
	_go_fuzz_dep_.CoverTab[53460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:73
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:73
	// _ = "end of CoverTab[53460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:73
}
func (e PlaceholderEnumValue) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:74
	_go_fuzz_dep_.CoverTab[53461]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:74
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:74
	// _ = "end of CoverTab[53461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:74
}
func (e PlaceholderEnumValue) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:75
	_go_fuzz_dep_.CoverTab[53462]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:75
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:75
	// _ = "end of CoverTab[53462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:75
}
func (e PlaceholderEnumValue) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:76
	_go_fuzz_dep_.CoverTab[53463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:76
	return protoreflect.FullName(e).Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:76
	// _ = "end of CoverTab[53463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:76
}
func (e PlaceholderEnumValue) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:77
	_go_fuzz_dep_.CoverTab[53464]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:77
	return protoreflect.FullName(e)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:77
	// _ = "end of CoverTab[53464]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:77
}
func (e PlaceholderEnumValue) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:78
	_go_fuzz_dep_.CoverTab[53465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:78
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:78
	// _ = "end of CoverTab[53465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:78
}
func (e PlaceholderEnumValue) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:79
	_go_fuzz_dep_.CoverTab[53466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:79
	return descopts.EnumValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:79
	// _ = "end of CoverTab[53466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:79
}
func (e PlaceholderEnumValue) Number() protoreflect.EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:80
	_go_fuzz_dep_.CoverTab[53467]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:80
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:80
	// _ = "end of CoverTab[53467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:80
}
func (e PlaceholderEnumValue) ProtoType(protoreflect.EnumValueDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:81
	_go_fuzz_dep_.CoverTab[53468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:81
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:81
	// _ = "end of CoverTab[53468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:81
}
func (e PlaceholderEnumValue) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:82
	_go_fuzz_dep_.CoverTab[53469]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:82
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:82
	// _ = "end of CoverTab[53469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:82
}

// PlaceholderMessage is a placeholder, representing only the full name.
type PlaceholderMessage protoreflect.FullName

func (m PlaceholderMessage) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:87
	_go_fuzz_dep_.CoverTab[53470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:87
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:87
	// _ = "end of CoverTab[53470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:87
}
func (m PlaceholderMessage) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:88
	_go_fuzz_dep_.CoverTab[53471]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:88
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:88
	// _ = "end of CoverTab[53471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:88
}
func (m PlaceholderMessage) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:89
	_go_fuzz_dep_.CoverTab[53472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:89
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:89
	// _ = "end of CoverTab[53472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:89
}
func (m PlaceholderMessage) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:90
	_go_fuzz_dep_.CoverTab[53473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:90
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:90
	// _ = "end of CoverTab[53473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:90
}
func (m PlaceholderMessage) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:91
	_go_fuzz_dep_.CoverTab[53474]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:91
	return protoreflect.FullName(m).Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:91
	// _ = "end of CoverTab[53474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:91
}
func (m PlaceholderMessage) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:92
	_go_fuzz_dep_.CoverTab[53475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:92
	return protoreflect.FullName(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:92
	// _ = "end of CoverTab[53475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:92
}
func (m PlaceholderMessage) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:93
	_go_fuzz_dep_.CoverTab[53476]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:93
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:93
	// _ = "end of CoverTab[53476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:93
}
func (m PlaceholderMessage) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:94
	_go_fuzz_dep_.CoverTab[53477]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:94
	return descopts.Message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:94
	// _ = "end of CoverTab[53477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:94
}
func (m PlaceholderMessage) IsMapEntry() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:95
	_go_fuzz_dep_.CoverTab[53478]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:95
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:95
	// _ = "end of CoverTab[53478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:95
}
func (m PlaceholderMessage) Fields() protoreflect.FieldDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:96
	_go_fuzz_dep_.CoverTab[53479]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:96
	return emptyFields
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:96
	// _ = "end of CoverTab[53479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:96
}
func (m PlaceholderMessage) Oneofs() protoreflect.OneofDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:97
	_go_fuzz_dep_.CoverTab[53480]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:97
	return emptyOneofs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:97
	// _ = "end of CoverTab[53480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:97
}
func (m PlaceholderMessage) ReservedNames() protoreflect.Names {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:98
	_go_fuzz_dep_.CoverTab[53481]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:98
	return emptyNames
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:98
	// _ = "end of CoverTab[53481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:98
}
func (m PlaceholderMessage) ReservedRanges() protoreflect.FieldRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:99
	_go_fuzz_dep_.CoverTab[53482]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:99
	return emptyFieldRanges
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:99
	// _ = "end of CoverTab[53482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:99
}
func (m PlaceholderMessage) RequiredNumbers() protoreflect.FieldNumbers {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:100
	_go_fuzz_dep_.CoverTab[53483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:100
	return emptyFieldNumbers
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:100
	// _ = "end of CoverTab[53483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:100
}
func (m PlaceholderMessage) ExtensionRanges() protoreflect.FieldRanges {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:101
	_go_fuzz_dep_.CoverTab[53484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:101
	return emptyFieldRanges
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:101
	// _ = "end of CoverTab[53484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:101
}
func (m PlaceholderMessage) ExtensionRangeOptions(int) protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:102
	_go_fuzz_dep_.CoverTab[53485]++
														panic("index out of range")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:103
	// _ = "end of CoverTab[53485]"
}
func (m PlaceholderMessage) Messages() protoreflect.MessageDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:105
	_go_fuzz_dep_.CoverTab[53486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:105
	return emptyMessages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:105
	// _ = "end of CoverTab[53486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:105
}
func (m PlaceholderMessage) Enums() protoreflect.EnumDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:106
	_go_fuzz_dep_.CoverTab[53487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:106
	return emptyEnums
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:106
	// _ = "end of CoverTab[53487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:106
}
func (m PlaceholderMessage) Extensions() protoreflect.ExtensionDescriptors {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:107
	_go_fuzz_dep_.CoverTab[53488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:107
	return emptyExtensions
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:107
	// _ = "end of CoverTab[53488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:107
}
func (m PlaceholderMessage) ProtoType(protoreflect.MessageDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:108
	_go_fuzz_dep_.CoverTab[53489]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:108
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:108
	// _ = "end of CoverTab[53489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:108
}
func (m PlaceholderMessage) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:109
	_go_fuzz_dep_.CoverTab[53490]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:109
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:109
	// _ = "end of CoverTab[53490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:109
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/placeholder.go:109
var _ = _go_fuzz_dep_.CoverTab
