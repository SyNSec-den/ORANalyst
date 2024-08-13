// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
package protoreflect

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:5
)

import (
	"strconv"
)

// SourceLocations is a list of source locations.
type SourceLocations interface {
	// Len reports the number of source locations in the proto file.
	Len() int
	// Get returns the ith SourceLocation. It panics if out of bounds.
	Get(int) SourceLocation

	// ByPath returns the SourceLocation for the given path,
	// returning the first location if multiple exist for the same path.
	// If multiple locations exist for the same path,
	// then SourceLocation.Next index can be used to identify the
	// index of the next SourceLocation.
	// If no location exists for this path, it returns the zero value.
	ByPath(path SourcePath) SourceLocation

	// ByDescriptor returns the SourceLocation for the given descriptor,
	// returning the first location if multiple exist for the same path.
	// If no location exists for this descriptor, it returns the zero value.
	ByDescriptor(desc Descriptor) SourceLocation

	doNotImplement
}

// SourceLocation describes a source location and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:34
// corresponds with the google.protobuf.SourceCodeInfo.Location message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:36
type SourceLocation struct {
	// Path is the path to the declaration from the root file descriptor.
	// The contents of this slice must not be mutated.
	Path	SourcePath

	// StartLine and StartColumn are the zero-indexed starting location
	// in the source file for the declaration.
	StartLine, StartColumn	int
	// EndLine and EndColumn are the zero-indexed ending location
	// in the source file for the declaration.
	// In the descriptor.proto, the end line may be omitted if it is identical
	// to the start line. Here, it is always populated.
	EndLine, EndColumn	int

	// LeadingDetachedComments are the leading detached comments
	// for the declaration. The contents of this slice must not be mutated.
	LeadingDetachedComments	[]string
	// LeadingComments is the leading attached comment for the declaration.
	LeadingComments	string
	// TrailingComments is the trailing attached comment for the declaration.
	TrailingComments	string

	// Next is an index into SourceLocations for the next source location that
	// has the same Path. It is zero if there is no next location.
	Next	int
}

// SourcePath identifies part of a file descriptor for a source location.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:63
// The SourcePath is a sequence of either field numbers or indexes into
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:63
// a repeated field that form a path starting from the root file descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:63
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:63
// See google.protobuf.SourceCodeInfo.Location.path.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:68
type SourcePath []int32

// Equal reports whether p1 equals p2.
func (p1 SourcePath) Equal(p2 SourcePath) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:71
	_go_fuzz_dep_.CoverTab[48608]++
													if len(p1) != len(p2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:72
		_go_fuzz_dep_.CoverTab[48611]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:73
		// _ = "end of CoverTab[48611]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:74
		_go_fuzz_dep_.CoverTab[48612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:74
		// _ = "end of CoverTab[48612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:74
	// _ = "end of CoverTab[48608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:74
	_go_fuzz_dep_.CoverTab[48609]++
													for i := range p1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:75
		_go_fuzz_dep_.CoverTab[48613]++
														if p1[i] != p2[i] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:76
			_go_fuzz_dep_.CoverTab[48614]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:77
			// _ = "end of CoverTab[48614]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:78
			_go_fuzz_dep_.CoverTab[48615]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:78
			// _ = "end of CoverTab[48615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:78
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:78
		// _ = "end of CoverTab[48613]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:79
	// _ = "end of CoverTab[48609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:79
	_go_fuzz_dep_.CoverTab[48610]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:80
	// _ = "end of CoverTab[48610]"
}

// String formats the path in a humanly readable manner.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
// The output is guaranteed to be deterministic,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
// making it suitable for use as a key into a Go map.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
// It is not guaranteed to be stable as the exact output could change
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
// in a future version of this module.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
// Example output:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:83
//	.message_type[6].nested_type[15].field[3]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:92
func (p SourcePath) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:92
	_go_fuzz_dep_.CoverTab[48616]++
													b := p.appendFileDescriptorProto(nil)
													for _, i := range p {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:94
		_go_fuzz_dep_.CoverTab[48618]++
														b = append(b, '.')
														b = strconv.AppendInt(b, int64(i), 10)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:96
		// _ = "end of CoverTab[48618]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:97
	// _ = "end of CoverTab[48616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:97
	_go_fuzz_dep_.CoverTab[48617]++
													return string(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:98
	// _ = "end of CoverTab[48617]"
}

type appendFunc func(*SourcePath, []byte) []byte

func (p *SourcePath) appendSingularField(b []byte, name string, f appendFunc) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:103
	_go_fuzz_dep_.CoverTab[48619]++
														if len(*p) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:104
		_go_fuzz_dep_.CoverTab[48622]++
															return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:105
		// _ = "end of CoverTab[48622]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:106
		_go_fuzz_dep_.CoverTab[48623]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:106
		// _ = "end of CoverTab[48623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:106
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:106
	// _ = "end of CoverTab[48619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:106
	_go_fuzz_dep_.CoverTab[48620]++
														b = append(b, '.')
														b = append(b, name...)
														*p = (*p)[1:]
														if f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:110
		_go_fuzz_dep_.CoverTab[48624]++
															b = f(p, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:111
		// _ = "end of CoverTab[48624]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:112
		_go_fuzz_dep_.CoverTab[48625]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:112
		// _ = "end of CoverTab[48625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:112
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:112
	// _ = "end of CoverTab[48620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:112
	_go_fuzz_dep_.CoverTab[48621]++
														return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:113
	// _ = "end of CoverTab[48621]"
}

func (p *SourcePath) appendRepeatedField(b []byte, name string, f appendFunc) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:116
	_go_fuzz_dep_.CoverTab[48626]++
														b = p.appendSingularField(b, name, nil)
														if len(*p) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:118
		_go_fuzz_dep_.CoverTab[48629]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:118
		return (*p)[0] < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:118
		// _ = "end of CoverTab[48629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:118
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:118
		_go_fuzz_dep_.CoverTab[48630]++
															return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:119
		// _ = "end of CoverTab[48630]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:120
		_go_fuzz_dep_.CoverTab[48631]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:120
		// _ = "end of CoverTab[48631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:120
	// _ = "end of CoverTab[48626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:120
	_go_fuzz_dep_.CoverTab[48627]++
														b = append(b, '[')
														b = strconv.AppendUint(b, uint64((*p)[0]), 10)
														b = append(b, ']')
														*p = (*p)[1:]
														if f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:125
		_go_fuzz_dep_.CoverTab[48632]++
															b = f(p, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:126
		// _ = "end of CoverTab[48632]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:127
		_go_fuzz_dep_.CoverTab[48633]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:127
		// _ = "end of CoverTab[48633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:127
	// _ = "end of CoverTab[48627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:127
	_go_fuzz_dep_.CoverTab[48628]++
														return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:128
	// _ = "end of CoverTab[48628]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:129
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/source.go:129
var _ = _go_fuzz_dep_.CoverTab
