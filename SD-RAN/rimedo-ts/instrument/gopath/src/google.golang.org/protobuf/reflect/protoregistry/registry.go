// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// Package protoregistry provides data structures to register and lookup
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// protobuf descriptor types.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// The Files registry contains file descriptors and provides the ability
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// to iterate over the files or lookup a specific descriptor within the files.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// Files only contains protobuf descriptors and has no understanding of Go
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// type information that may be associated with each descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// The Types registry contains descriptor types for which there is a known
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// Go type associated with that descriptor. It provides the ability to iterate
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:5
// over the registered types or lookup a type by name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
package protoregistry

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:16
)

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// conflictPolicy configures the policy for handling registration conflicts.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
// It can be over-written at compile time with a linker-initialized variable:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//	go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
// It can be over-written at program execution with an environment variable:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn ./main
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
// Neither of the above are covered by the compatibility promise and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:30
// may be removed in a future release of this module.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:42
var conflictPolicy = "panic"	// "panic" | "warn" | "ignore"

// ignoreConflict reports whether to ignore a registration conflict
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:44
// given the descriptor being registered and the error.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:44
// It is a variable so that the behavior is easily overridden in another file.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:47
var ignoreConflict = func(d protoreflect.Descriptor, err error) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:47
	_go_fuzz_dep_.CoverTab[50201]++
														const env = "GOLANG_PROTOBUF_REGISTRATION_CONFLICT"
														const faq = "https://developers.google.com/protocol-buffers/docs/reference/go/faq#namespace-conflict"
														policy := conflictPolicy
														if v := os.Getenv(env); v != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:51
		_go_fuzz_dep_.CoverTab[50203]++
															policy = v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:52
		// _ = "end of CoverTab[50203]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:53
		_go_fuzz_dep_.CoverTab[50204]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:53
		// _ = "end of CoverTab[50204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:53
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:53
	// _ = "end of CoverTab[50201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:53
	_go_fuzz_dep_.CoverTab[50202]++
														switch policy {
	case "panic":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:55
		_go_fuzz_dep_.CoverTab[50205]++
															panic(fmt.Sprintf("%v\nSee %v\n", err, faq))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:56
		// _ = "end of CoverTab[50205]"
	case "warn":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:57
		_go_fuzz_dep_.CoverTab[50206]++
															fmt.Fprintf(os.Stderr, "WARNING: %v\nSee %v\n\n", err, faq)
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:59
		// _ = "end of CoverTab[50206]"
	case "ignore":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:60
		_go_fuzz_dep_.CoverTab[50207]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:61
		// _ = "end of CoverTab[50207]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:62
		_go_fuzz_dep_.CoverTab[50208]++
															panic("invalid " + env + " value: " + os.Getenv(env))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:63
		// _ = "end of CoverTab[50208]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:64
	// _ = "end of CoverTab[50202]"
}

var globalMutex sync.RWMutex

// GlobalFiles is a global registry of file descriptors.
var GlobalFiles *Files = new(Files)

// GlobalTypes is the registry used by default for type lookups
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:72
// unless a local registry is provided by the user.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:74
var GlobalTypes *Types = new(Types)

// NotFound is a sentinel error value to indicate that the type was not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:76
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:76
// Since registry lookup can happen in the critical performance path, resolvers
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:76
// must return this exact error value, not an error wrapping it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:80
var NotFound = errors.New("not found")

// Files is a registry for looking up or iterating over files and the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:82
// descriptors contained within them.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:82
// The Find and Range methods are safe for concurrent use.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:85
type Files struct {
	// The map of descsByName contains:
	//	EnumDescriptor
	//	EnumValueDescriptor
	//	MessageDescriptor
	//	ExtensionDescriptor
	//	ServiceDescriptor
	//	*packageDescriptor
	//
	// Note that files are stored as a slice, since a package may contain
	// multiple files. Only top-level declarations are registered.
	// Note that enum values are in the top-level since that are in the same
	// scope as the parent enum.
	descsByName	map[protoreflect.FullName]interface{}
	filesByPath	map[string][]protoreflect.FileDescriptor
	numFiles	int
}

type packageDescriptor struct {
	files []protoreflect.FileDescriptor
}

// RegisterFile registers the provided file descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:107
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:107
// If any descriptor within the file conflicts with the descriptor of any
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:107
// previously registered file (e.g., two enums with the same full name),
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:107
// then the file is not registered and an error is returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:107
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:107
// It is permitted for multiple files to have the same file path.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:114
func (r *Files) RegisterFile(file protoreflect.FileDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:114
	_go_fuzz_dep_.CoverTab[50209]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:115
		_go_fuzz_dep_.CoverTab[50218]++
															globalMutex.Lock()
															defer globalMutex.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:117
		// _ = "end of CoverTab[50218]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:118
		_go_fuzz_dep_.CoverTab[50219]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:118
		// _ = "end of CoverTab[50219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:118
	// _ = "end of CoverTab[50209]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:118
	_go_fuzz_dep_.CoverTab[50210]++
														if r.descsByName == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:119
		_go_fuzz_dep_.CoverTab[50220]++
															r.descsByName = map[protoreflect.FullName]interface{}{
			"": &packageDescriptor{},
		}
															r.filesByPath = make(map[string][]protoreflect.FileDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:123
		// _ = "end of CoverTab[50220]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:124
		_go_fuzz_dep_.CoverTab[50221]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:124
		// _ = "end of CoverTab[50221]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:124
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:124
	// _ = "end of CoverTab[50210]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:124
	_go_fuzz_dep_.CoverTab[50211]++
														path := file.Path()
														if prev := r.filesByPath[path]; len(prev) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:126
		_go_fuzz_dep_.CoverTab[50222]++
															r.checkGenProtoConflict(path)
															err := errors.New("file %q is already registered", file.Path())
															err = amendErrorWithCaller(err, prev[0], file)
															if !(r == GlobalFiles && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:130
			_go_fuzz_dep_.CoverTab[50223]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:130
			return ignoreConflict(file, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:130
			// _ = "end of CoverTab[50223]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:130
		}()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:130
			_go_fuzz_dep_.CoverTab[50224]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:131
			// _ = "end of CoverTab[50224]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:132
			_go_fuzz_dep_.CoverTab[50225]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:132
			// _ = "end of CoverTab[50225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:132
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:132
		// _ = "end of CoverTab[50222]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:133
		_go_fuzz_dep_.CoverTab[50226]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:133
		// _ = "end of CoverTab[50226]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:133
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:133
	// _ = "end of CoverTab[50211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:133
	_go_fuzz_dep_.CoverTab[50212]++

														for name := file.Package(); name != ""; name = name.Parent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:135
		_go_fuzz_dep_.CoverTab[50227]++
															switch prev := r.descsByName[name]; prev.(type) {
		case nil, *packageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:137
			_go_fuzz_dep_.CoverTab[50228]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:137
			// _ = "end of CoverTab[50228]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:138
			_go_fuzz_dep_.CoverTab[50229]++
																err := errors.New("file %q has a package name conflict over %v", file.Path(), name)
																err = amendErrorWithCaller(err, prev, file)
																if r == GlobalFiles && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:141
				_go_fuzz_dep_.CoverTab[50231]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:141
				return ignoreConflict(file, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:141
				// _ = "end of CoverTab[50231]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:141
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:141
				_go_fuzz_dep_.CoverTab[50232]++
																	err = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:142
				// _ = "end of CoverTab[50232]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:143
				_go_fuzz_dep_.CoverTab[50233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:143
				// _ = "end of CoverTab[50233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:143
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:143
			// _ = "end of CoverTab[50229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:143
			_go_fuzz_dep_.CoverTab[50230]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:144
			// _ = "end of CoverTab[50230]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:145
		// _ = "end of CoverTab[50227]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:146
	// _ = "end of CoverTab[50212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:146
	_go_fuzz_dep_.CoverTab[50213]++
														var err error
														var hasConflict bool
														rangeTopLevelDescriptors(file, func(d protoreflect.Descriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:149
		_go_fuzz_dep_.CoverTab[50234]++
															if prev := r.descsByName[d.FullName()]; prev != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:150
			_go_fuzz_dep_.CoverTab[50235]++
																hasConflict = true
																err = errors.New("file %q has a name conflict over %v", file.Path(), d.FullName())
																err = amendErrorWithCaller(err, prev, file)
																if r == GlobalFiles && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:154
				_go_fuzz_dep_.CoverTab[50236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:154
				return ignoreConflict(d, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:154
				// _ = "end of CoverTab[50236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:154
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:154
				_go_fuzz_dep_.CoverTab[50237]++
																	err = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:155
				// _ = "end of CoverTab[50237]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:156
				_go_fuzz_dep_.CoverTab[50238]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:156
				// _ = "end of CoverTab[50238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:156
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:156
			// _ = "end of CoverTab[50235]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:157
			_go_fuzz_dep_.CoverTab[50239]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:157
			// _ = "end of CoverTab[50239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:157
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:157
		// _ = "end of CoverTab[50234]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:158
	// _ = "end of CoverTab[50213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:158
	_go_fuzz_dep_.CoverTab[50214]++
														if hasConflict {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:159
		_go_fuzz_dep_.CoverTab[50240]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:160
		// _ = "end of CoverTab[50240]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:161
		_go_fuzz_dep_.CoverTab[50241]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:161
		// _ = "end of CoverTab[50241]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:161
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:161
	// _ = "end of CoverTab[50214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:161
	_go_fuzz_dep_.CoverTab[50215]++

														for name := file.Package(); name != ""; name = name.Parent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:163
		_go_fuzz_dep_.CoverTab[50242]++
															if r.descsByName[name] == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:164
			_go_fuzz_dep_.CoverTab[50243]++
																r.descsByName[name] = &packageDescriptor{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:165
			// _ = "end of CoverTab[50243]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:166
			_go_fuzz_dep_.CoverTab[50244]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:166
			// _ = "end of CoverTab[50244]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:166
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:166
		// _ = "end of CoverTab[50242]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:167
	// _ = "end of CoverTab[50215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:167
	_go_fuzz_dep_.CoverTab[50216]++
														p := r.descsByName[file.Package()].(*packageDescriptor)
														p.files = append(p.files, file)
														rangeTopLevelDescriptors(file, func(d protoreflect.Descriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:170
		_go_fuzz_dep_.CoverTab[50245]++
															r.descsByName[d.FullName()] = d
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:171
		// _ = "end of CoverTab[50245]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:172
	// _ = "end of CoverTab[50216]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:172
	_go_fuzz_dep_.CoverTab[50217]++
														r.filesByPath[path] = append(r.filesByPath[path], file)
														r.numFiles++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:175
	// _ = "end of CoverTab[50217]"
}

// Several well-known types were hosted in the google.golang.org/genproto module
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:178
// but were later moved to this module. To avoid a weak dependency on the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:178
// genproto module (and its relatively large set of transitive dependencies),
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:178
// we rely on a registration conflict to determine whether the genproto version
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:178
// is too old (i.e., does not contain aliases to the new type declarations).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:183
func (r *Files) checkGenProtoConflict(path string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:183
	_go_fuzz_dep_.CoverTab[50246]++
														if r != GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:184
		_go_fuzz_dep_.CoverTab[50249]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:185
		// _ = "end of CoverTab[50249]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:186
		_go_fuzz_dep_.CoverTab[50250]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:186
		// _ = "end of CoverTab[50250]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:186
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:186
	// _ = "end of CoverTab[50246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:186
	_go_fuzz_dep_.CoverTab[50247]++
														var prevPath string
														const prevModule = "google.golang.org/genproto"
														const prevVersion = "cb27e3aa (May 26th, 2020)"
														switch path {
	case "google/protobuf/field_mask.proto":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:191
		_go_fuzz_dep_.CoverTab[50251]++
															prevPath = prevModule + "/protobuf/field_mask"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:192
		// _ = "end of CoverTab[50251]"
	case "google/protobuf/api.proto":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:193
		_go_fuzz_dep_.CoverTab[50252]++
															prevPath = prevModule + "/protobuf/api"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:194
		// _ = "end of CoverTab[50252]"
	case "google/protobuf/type.proto":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:195
		_go_fuzz_dep_.CoverTab[50253]++
															prevPath = prevModule + "/protobuf/ptype"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:196
		// _ = "end of CoverTab[50253]"
	case "google/protobuf/source_context.proto":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:197
		_go_fuzz_dep_.CoverTab[50254]++
															prevPath = prevModule + "/protobuf/source_context"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:198
		// _ = "end of CoverTab[50254]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:199
		_go_fuzz_dep_.CoverTab[50255]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:200
		// _ = "end of CoverTab[50255]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:201
	// _ = "end of CoverTab[50247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:201
	_go_fuzz_dep_.CoverTab[50248]++
														pkgName := strings.TrimSuffix(strings.TrimPrefix(path, "google/protobuf/"), ".proto")
														pkgName = strings.Replace(pkgName, "_", "", -1) + "pb"
														currPath := "google.golang.org/protobuf/types/known/" + pkgName
														panic(fmt.Sprintf(""+
		"duplicate registration of %q\n"+
		"\n"+
		"The generated definition for this file has moved:\n"+
		"\tfrom: %q\n"+
		"\tto:   %q\n"+
		"A dependency on the %q module must\n"+
		"be at version %v or higher.\n"+
		"\n"+
		"Upgrade the dependency by running:\n"+
		"\tgo get -u %v\n",
		path, prevPath, currPath, prevModule, prevVersion, prevPath))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:216
	// _ = "end of CoverTab[50248]"
}

// FindDescriptorByName looks up a descriptor by the full name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:219
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:219
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:222
func (r *Files) FindDescriptorByName(name protoreflect.FullName) (protoreflect.Descriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:222
	_go_fuzz_dep_.CoverTab[50256]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:223
		_go_fuzz_dep_.CoverTab[50260]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:224
		// _ = "end of CoverTab[50260]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:225
		_go_fuzz_dep_.CoverTab[50261]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:225
		// _ = "end of CoverTab[50261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:225
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:225
	// _ = "end of CoverTab[50256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:225
	_go_fuzz_dep_.CoverTab[50257]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:226
		_go_fuzz_dep_.CoverTab[50262]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:228
		// _ = "end of CoverTab[50262]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:229
		_go_fuzz_dep_.CoverTab[50263]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:229
		// _ = "end of CoverTab[50263]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:229
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:229
	// _ = "end of CoverTab[50257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:229
	_go_fuzz_dep_.CoverTab[50258]++
														prefix := name
														suffix := nameSuffix("")
														for prefix != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:232
		_go_fuzz_dep_.CoverTab[50264]++
															if d, ok := r.descsByName[prefix]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:233
			_go_fuzz_dep_.CoverTab[50266]++
																switch d := d.(type) {
			case protoreflect.EnumDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:235
				_go_fuzz_dep_.CoverTab[50268]++
																	if d.FullName() == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:236
					_go_fuzz_dep_.CoverTab[50275]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:237
					// _ = "end of CoverTab[50275]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:238
					_go_fuzz_dep_.CoverTab[50276]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:238
					// _ = "end of CoverTab[50276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:238
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:238
				// _ = "end of CoverTab[50268]"
			case protoreflect.EnumValueDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:239
				_go_fuzz_dep_.CoverTab[50269]++
																	if d.FullName() == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:240
					_go_fuzz_dep_.CoverTab[50277]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:241
					// _ = "end of CoverTab[50277]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:242
					_go_fuzz_dep_.CoverTab[50278]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:242
					// _ = "end of CoverTab[50278]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:242
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:242
				// _ = "end of CoverTab[50269]"
			case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:243
				_go_fuzz_dep_.CoverTab[50270]++
																	if d.FullName() == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:244
					_go_fuzz_dep_.CoverTab[50279]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:245
					// _ = "end of CoverTab[50279]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:246
					_go_fuzz_dep_.CoverTab[50280]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:246
					// _ = "end of CoverTab[50280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:246
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:246
				// _ = "end of CoverTab[50270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:246
				_go_fuzz_dep_.CoverTab[50271]++
																	if d := findDescriptorInMessage(d, suffix); d != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:247
					_go_fuzz_dep_.CoverTab[50281]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:247
					return d.FullName() == name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:247
					// _ = "end of CoverTab[50281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:247
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:247
					_go_fuzz_dep_.CoverTab[50282]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:248
					// _ = "end of CoverTab[50282]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:249
					_go_fuzz_dep_.CoverTab[50283]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:249
					// _ = "end of CoverTab[50283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:249
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:249
				// _ = "end of CoverTab[50271]"
			case protoreflect.ExtensionDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:250
				_go_fuzz_dep_.CoverTab[50272]++
																	if d.FullName() == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:251
					_go_fuzz_dep_.CoverTab[50284]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:252
					// _ = "end of CoverTab[50284]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:253
					_go_fuzz_dep_.CoverTab[50285]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:253
					// _ = "end of CoverTab[50285]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:253
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:253
				// _ = "end of CoverTab[50272]"
			case protoreflect.ServiceDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:254
				_go_fuzz_dep_.CoverTab[50273]++
																	if d.FullName() == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:255
					_go_fuzz_dep_.CoverTab[50286]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:256
					// _ = "end of CoverTab[50286]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:257
					_go_fuzz_dep_.CoverTab[50287]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:257
					// _ = "end of CoverTab[50287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:257
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:257
				// _ = "end of CoverTab[50273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:257
				_go_fuzz_dep_.CoverTab[50274]++
																	if d := d.Methods().ByName(suffix.Pop()); d != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:258
					_go_fuzz_dep_.CoverTab[50288]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:258
					return d.FullName() == name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:258
					// _ = "end of CoverTab[50288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:258
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:258
					_go_fuzz_dep_.CoverTab[50289]++
																		return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:259
					// _ = "end of CoverTab[50289]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:260
					_go_fuzz_dep_.CoverTab[50290]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:260
					// _ = "end of CoverTab[50290]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:260
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:260
				// _ = "end of CoverTab[50274]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:261
			// _ = "end of CoverTab[50266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:261
			_go_fuzz_dep_.CoverTab[50267]++
																return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:262
			// _ = "end of CoverTab[50267]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:263
			_go_fuzz_dep_.CoverTab[50291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:263
			// _ = "end of CoverTab[50291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:263
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:263
		// _ = "end of CoverTab[50264]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:263
		_go_fuzz_dep_.CoverTab[50265]++
															prefix = prefix.Parent()
															suffix = nameSuffix(name[len(prefix)+len("."):])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:265
		// _ = "end of CoverTab[50265]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:266
	// _ = "end of CoverTab[50258]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:266
	_go_fuzz_dep_.CoverTab[50259]++
														return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:267
	// _ = "end of CoverTab[50259]"
}

func findDescriptorInMessage(md protoreflect.MessageDescriptor, suffix nameSuffix) protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:270
	_go_fuzz_dep_.CoverTab[50292]++
														name := suffix.Pop()
														if suffix == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:272
		_go_fuzz_dep_.CoverTab[50295]++
															if ed := md.Enums().ByName(name); ed != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:273
			_go_fuzz_dep_.CoverTab[50300]++
																return ed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:274
			// _ = "end of CoverTab[50300]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:275
			_go_fuzz_dep_.CoverTab[50301]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:275
			// _ = "end of CoverTab[50301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:275
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:275
		// _ = "end of CoverTab[50295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:275
		_go_fuzz_dep_.CoverTab[50296]++
															for i := md.Enums().Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:276
			_go_fuzz_dep_.CoverTab[50302]++
																if vd := md.Enums().Get(i).Values().ByName(name); vd != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:277
				_go_fuzz_dep_.CoverTab[50303]++
																	return vd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:278
				// _ = "end of CoverTab[50303]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:279
				_go_fuzz_dep_.CoverTab[50304]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:279
				// _ = "end of CoverTab[50304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:279
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:279
			// _ = "end of CoverTab[50302]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:280
		// _ = "end of CoverTab[50296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:280
		_go_fuzz_dep_.CoverTab[50297]++
															if xd := md.Extensions().ByName(name); xd != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:281
			_go_fuzz_dep_.CoverTab[50305]++
																return xd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:282
			// _ = "end of CoverTab[50305]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:283
			_go_fuzz_dep_.CoverTab[50306]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:283
			// _ = "end of CoverTab[50306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:283
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:283
		// _ = "end of CoverTab[50297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:283
		_go_fuzz_dep_.CoverTab[50298]++
															if fd := md.Fields().ByName(name); fd != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:284
			_go_fuzz_dep_.CoverTab[50307]++
																return fd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:285
			// _ = "end of CoverTab[50307]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:286
			_go_fuzz_dep_.CoverTab[50308]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:286
			// _ = "end of CoverTab[50308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:286
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:286
		// _ = "end of CoverTab[50298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:286
		_go_fuzz_dep_.CoverTab[50299]++
															if od := md.Oneofs().ByName(name); od != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:287
			_go_fuzz_dep_.CoverTab[50309]++
																return od
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:288
			// _ = "end of CoverTab[50309]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:289
			_go_fuzz_dep_.CoverTab[50310]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:289
			// _ = "end of CoverTab[50310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:289
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:289
		// _ = "end of CoverTab[50299]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:290
		_go_fuzz_dep_.CoverTab[50311]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:290
		// _ = "end of CoverTab[50311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:290
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:290
	// _ = "end of CoverTab[50292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:290
	_go_fuzz_dep_.CoverTab[50293]++
														if md := md.Messages().ByName(name); md != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:291
		_go_fuzz_dep_.CoverTab[50312]++
															if suffix == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:292
			_go_fuzz_dep_.CoverTab[50314]++
																return md
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:293
			// _ = "end of CoverTab[50314]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:294
			_go_fuzz_dep_.CoverTab[50315]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:294
			// _ = "end of CoverTab[50315]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:294
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:294
		// _ = "end of CoverTab[50312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:294
		_go_fuzz_dep_.CoverTab[50313]++
															return findDescriptorInMessage(md, suffix)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:295
		// _ = "end of CoverTab[50313]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:296
		_go_fuzz_dep_.CoverTab[50316]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:296
		// _ = "end of CoverTab[50316]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:296
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:296
	// _ = "end of CoverTab[50293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:296
	_go_fuzz_dep_.CoverTab[50294]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:297
	// _ = "end of CoverTab[50294]"
}

type nameSuffix string

func (s *nameSuffix) Pop() (name protoreflect.Name) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:302
	_go_fuzz_dep_.CoverTab[50317]++
														if i := strings.IndexByte(string(*s), '.'); i >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:303
		_go_fuzz_dep_.CoverTab[50319]++
															name, *s = protoreflect.Name((*s)[:i]), (*s)[i+1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:304
		// _ = "end of CoverTab[50319]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:305
		_go_fuzz_dep_.CoverTab[50320]++
															name, *s = protoreflect.Name((*s)), ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:306
		// _ = "end of CoverTab[50320]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:307
	// _ = "end of CoverTab[50317]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:307
	_go_fuzz_dep_.CoverTab[50318]++
														return name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:308
	// _ = "end of CoverTab[50318]"
}

// FindFileByPath looks up a file by the path.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:311
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:311
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:311
// This returns an error if multiple files have the same path.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:315
func (r *Files) FindFileByPath(path string) (protoreflect.FileDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:315
	_go_fuzz_dep_.CoverTab[50321]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:316
		_go_fuzz_dep_.CoverTab[50324]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:317
		// _ = "end of CoverTab[50324]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:318
		_go_fuzz_dep_.CoverTab[50325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:318
		// _ = "end of CoverTab[50325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:318
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:318
	// _ = "end of CoverTab[50321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:318
	_go_fuzz_dep_.CoverTab[50322]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:319
		_go_fuzz_dep_.CoverTab[50326]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:321
		// _ = "end of CoverTab[50326]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:322
		_go_fuzz_dep_.CoverTab[50327]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:322
		// _ = "end of CoverTab[50327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:322
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:322
	// _ = "end of CoverTab[50322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:322
	_go_fuzz_dep_.CoverTab[50323]++
														fds := r.filesByPath[path]
														switch len(fds) {
	case 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:325
		_go_fuzz_dep_.CoverTab[50328]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:326
		// _ = "end of CoverTab[50328]"
	case 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:327
		_go_fuzz_dep_.CoverTab[50329]++
															return fds[0], nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:328
		// _ = "end of CoverTab[50329]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:329
		_go_fuzz_dep_.CoverTab[50330]++
															return nil, errors.New("multiple files named %q", path)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:330
		// _ = "end of CoverTab[50330]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:331
	// _ = "end of CoverTab[50323]"
}

// NumFiles reports the number of registered files,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:334
// including duplicate files with the same name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:336
func (r *Files) NumFiles() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:336
	_go_fuzz_dep_.CoverTab[50331]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:337
		_go_fuzz_dep_.CoverTab[50334]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:338
		// _ = "end of CoverTab[50334]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:339
		_go_fuzz_dep_.CoverTab[50335]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:339
		// _ = "end of CoverTab[50335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:339
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:339
	// _ = "end of CoverTab[50331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:339
	_go_fuzz_dep_.CoverTab[50332]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:340
		_go_fuzz_dep_.CoverTab[50336]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:342
		// _ = "end of CoverTab[50336]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:343
		_go_fuzz_dep_.CoverTab[50337]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:343
		// _ = "end of CoverTab[50337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:343
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:343
	// _ = "end of CoverTab[50332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:343
	_go_fuzz_dep_.CoverTab[50333]++
														return r.numFiles
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:344
	// _ = "end of CoverTab[50333]"
}

// RangeFiles iterates over all registered files while f returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:347
// If multiple files have the same name, RangeFiles iterates over all of them.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:347
// The iteration order is undefined.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:350
func (r *Files) RangeFiles(f func(protoreflect.FileDescriptor) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:350
	_go_fuzz_dep_.CoverTab[50338]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:351
		_go_fuzz_dep_.CoverTab[50341]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:352
		// _ = "end of CoverTab[50341]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:353
		_go_fuzz_dep_.CoverTab[50342]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:353
		// _ = "end of CoverTab[50342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:353
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:353
	// _ = "end of CoverTab[50338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:353
	_go_fuzz_dep_.CoverTab[50339]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:354
		_go_fuzz_dep_.CoverTab[50343]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:356
		// _ = "end of CoverTab[50343]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:357
		_go_fuzz_dep_.CoverTab[50344]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:357
		// _ = "end of CoverTab[50344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:357
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:357
	// _ = "end of CoverTab[50339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:357
	_go_fuzz_dep_.CoverTab[50340]++
														for _, files := range r.filesByPath {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:358
		_go_fuzz_dep_.CoverTab[50345]++
															for _, file := range files {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:359
			_go_fuzz_dep_.CoverTab[50346]++
																if !f(file) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:360
				_go_fuzz_dep_.CoverTab[50347]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:361
				// _ = "end of CoverTab[50347]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:362
				_go_fuzz_dep_.CoverTab[50348]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:362
				// _ = "end of CoverTab[50348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:362
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:362
			// _ = "end of CoverTab[50346]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:363
		// _ = "end of CoverTab[50345]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:364
	// _ = "end of CoverTab[50340]"
}

// NumFilesByPackage reports the number of registered files in a proto package.
func (r *Files) NumFilesByPackage(name protoreflect.FullName) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:368
	_go_fuzz_dep_.CoverTab[50349]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:369
		_go_fuzz_dep_.CoverTab[50353]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:370
		// _ = "end of CoverTab[50353]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:371
		_go_fuzz_dep_.CoverTab[50354]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:371
		// _ = "end of CoverTab[50354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:371
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:371
	// _ = "end of CoverTab[50349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:371
	_go_fuzz_dep_.CoverTab[50350]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:372
		_go_fuzz_dep_.CoverTab[50355]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:374
		// _ = "end of CoverTab[50355]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:375
		_go_fuzz_dep_.CoverTab[50356]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:375
		// _ = "end of CoverTab[50356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:375
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:375
	// _ = "end of CoverTab[50350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:375
	_go_fuzz_dep_.CoverTab[50351]++
														p, ok := r.descsByName[name].(*packageDescriptor)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:377
		_go_fuzz_dep_.CoverTab[50357]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:378
		// _ = "end of CoverTab[50357]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:379
		_go_fuzz_dep_.CoverTab[50358]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:379
		// _ = "end of CoverTab[50358]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:379
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:379
	// _ = "end of CoverTab[50351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:379
	_go_fuzz_dep_.CoverTab[50352]++
														return len(p.files)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:380
	// _ = "end of CoverTab[50352]"
}

// RangeFilesByPackage iterates over all registered files in a given proto package
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:383
// while f returns true. The iteration order is undefined.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:385
func (r *Files) RangeFilesByPackage(name protoreflect.FullName, f func(protoreflect.FileDescriptor) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:385
	_go_fuzz_dep_.CoverTab[50359]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:386
		_go_fuzz_dep_.CoverTab[50363]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:387
		// _ = "end of CoverTab[50363]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:388
		_go_fuzz_dep_.CoverTab[50364]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:388
		// _ = "end of CoverTab[50364]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:388
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:388
	// _ = "end of CoverTab[50359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:388
	_go_fuzz_dep_.CoverTab[50360]++
														if r == GlobalFiles {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:389
		_go_fuzz_dep_.CoverTab[50365]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:391
		// _ = "end of CoverTab[50365]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:392
		_go_fuzz_dep_.CoverTab[50366]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:392
		// _ = "end of CoverTab[50366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:392
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:392
	// _ = "end of CoverTab[50360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:392
	_go_fuzz_dep_.CoverTab[50361]++
														p, ok := r.descsByName[name].(*packageDescriptor)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:394
		_go_fuzz_dep_.CoverTab[50367]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:395
		// _ = "end of CoverTab[50367]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:396
		_go_fuzz_dep_.CoverTab[50368]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:396
		// _ = "end of CoverTab[50368]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:396
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:396
	// _ = "end of CoverTab[50361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:396
	_go_fuzz_dep_.CoverTab[50362]++
														for _, file := range p.files {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:397
		_go_fuzz_dep_.CoverTab[50369]++
															if !f(file) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:398
			_go_fuzz_dep_.CoverTab[50370]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:399
			// _ = "end of CoverTab[50370]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:400
			_go_fuzz_dep_.CoverTab[50371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:400
			// _ = "end of CoverTab[50371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:400
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:400
		// _ = "end of CoverTab[50369]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:401
	// _ = "end of CoverTab[50362]"
}

// rangeTopLevelDescriptors iterates over all top-level descriptors in a file
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:404
// which will be directly entered into the registry.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:406
func rangeTopLevelDescriptors(fd protoreflect.FileDescriptor, f func(protoreflect.Descriptor)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:406
	_go_fuzz_dep_.CoverTab[50372]++
														eds := fd.Enums()
														for i := eds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:408
		_go_fuzz_dep_.CoverTab[50376]++
															f(eds.Get(i))
															vds := eds.Get(i).Values()
															for i := vds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:411
			_go_fuzz_dep_.CoverTab[50377]++
																f(vds.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:412
			// _ = "end of CoverTab[50377]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:413
		// _ = "end of CoverTab[50376]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:414
	// _ = "end of CoverTab[50372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:414
	_go_fuzz_dep_.CoverTab[50373]++
														mds := fd.Messages()
														for i := mds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:416
		_go_fuzz_dep_.CoverTab[50378]++
															f(mds.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:417
		// _ = "end of CoverTab[50378]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:418
	// _ = "end of CoverTab[50373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:418
	_go_fuzz_dep_.CoverTab[50374]++
														xds := fd.Extensions()
														for i := xds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:420
		_go_fuzz_dep_.CoverTab[50379]++
															f(xds.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:421
		// _ = "end of CoverTab[50379]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:422
	// _ = "end of CoverTab[50374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:422
	_go_fuzz_dep_.CoverTab[50375]++
														sds := fd.Services()
														for i := sds.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:424
		_go_fuzz_dep_.CoverTab[50380]++
															f(sds.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:425
		// _ = "end of CoverTab[50380]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:426
	// _ = "end of CoverTab[50375]"
}

// MessageTypeResolver is an interface for looking up messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:429
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:429
// A compliant implementation must deterministically return the same type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:429
// if no error is encountered.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:429
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:429
// The Types type implements this interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:435
type MessageTypeResolver interface {
	// FindMessageByName looks up a message by its full name.
	// E.g., "google.protobuf.Any"
	//
	// This return (nil, NotFound) if not found.
	FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error)

	// FindMessageByURL looks up a message by a URL identifier.
	// See documentation on google.protobuf.Any.type_url for the URL format.
	//
	// This returns (nil, NotFound) if not found.
	FindMessageByURL(url string) (protoreflect.MessageType, error)
}

// ExtensionTypeResolver is an interface for looking up extensions.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:449
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:449
// A compliant implementation must deterministically return the same type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:449
// if no error is encountered.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:449
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:449
// The Types type implements this interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:455
type ExtensionTypeResolver interface {
	// FindExtensionByName looks up a extension field by the field's full name.
	// Note that this is the full name of the field as determined by
	// where the extension is declared and is unrelated to the full name of the
	// message being extended.
	//
	// This returns (nil, NotFound) if not found.
	FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)

	// FindExtensionByNumber looks up a extension field by the field number
	// within some parent message, identified by full name.
	//
	// This returns (nil, NotFound) if not found.
	FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
}

var (
	_	MessageTypeResolver	= (*Types)(nil)
	_	ExtensionTypeResolver	= (*Types)(nil)
)

// Types is a registry for looking up or iterating over descriptor types.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:476
// The Find and Range methods are safe for concurrent use.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:478
type Types struct {
	typesByName		typesByName
	extensionsByMessage	extensionsByMessage

	numEnums	int
	numMessages	int
	numExtensions	int
}

type (
	typesByName		map[protoreflect.FullName]interface{}
	extensionsByMessage	map[protoreflect.FullName]extensionsByNumber
	extensionsByNumber	map[protoreflect.FieldNumber]protoreflect.ExtensionType
)

// RegisterMessage registers the provided message type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:493
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:493
// If a naming conflict occurs, the type is not registered and an error is returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:496
func (r *Types) RegisterMessage(mt protoreflect.MessageType) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:496
	_go_fuzz_dep_.CoverTab[50381]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:499
	md := mt.Descriptor()

	if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:501
		_go_fuzz_dep_.CoverTab[50384]++
															globalMutex.Lock()
															defer globalMutex.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:503
		// _ = "end of CoverTab[50384]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:504
		_go_fuzz_dep_.CoverTab[50385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:504
		// _ = "end of CoverTab[50385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:504
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:504
	// _ = "end of CoverTab[50381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:504
	_go_fuzz_dep_.CoverTab[50382]++

														if err := r.register("message", md, mt); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:506
		_go_fuzz_dep_.CoverTab[50386]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:507
		// _ = "end of CoverTab[50386]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:508
		_go_fuzz_dep_.CoverTab[50387]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:508
		// _ = "end of CoverTab[50387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:508
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:508
	// _ = "end of CoverTab[50382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:508
	_go_fuzz_dep_.CoverTab[50383]++
														r.numMessages++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:510
	// _ = "end of CoverTab[50383]"
}

// RegisterEnum registers the provided enum type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:513
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:513
// If a naming conflict occurs, the type is not registered and an error is returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:516
func (r *Types) RegisterEnum(et protoreflect.EnumType) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:516
	_go_fuzz_dep_.CoverTab[50388]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:519
	ed := et.Descriptor()

	if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:521
		_go_fuzz_dep_.CoverTab[50391]++
															globalMutex.Lock()
															defer globalMutex.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:523
		// _ = "end of CoverTab[50391]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:524
		_go_fuzz_dep_.CoverTab[50392]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:524
		// _ = "end of CoverTab[50392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:524
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:524
	// _ = "end of CoverTab[50388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:524
	_go_fuzz_dep_.CoverTab[50389]++

														if err := r.register("enum", ed, et); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:526
		_go_fuzz_dep_.CoverTab[50393]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:527
		// _ = "end of CoverTab[50393]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:528
		_go_fuzz_dep_.CoverTab[50394]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:528
		// _ = "end of CoverTab[50394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:528
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:528
	// _ = "end of CoverTab[50389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:528
	_go_fuzz_dep_.CoverTab[50390]++
														r.numEnums++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:530
	// _ = "end of CoverTab[50390]"
}

// RegisterExtension registers the provided extension type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:533
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:533
// If a naming conflict occurs, the type is not registered and an error is returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:536
func (r *Types) RegisterExtension(xt protoreflect.ExtensionType) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:536
	_go_fuzz_dep_.CoverTab[50395]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:542
	xd := xt.TypeDescriptor()

	if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:544
		_go_fuzz_dep_.CoverTab[50401]++
															globalMutex.Lock()
															defer globalMutex.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:546
		// _ = "end of CoverTab[50401]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:547
		_go_fuzz_dep_.CoverTab[50402]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:547
		// _ = "end of CoverTab[50402]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:547
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:547
	// _ = "end of CoverTab[50395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:547
	_go_fuzz_dep_.CoverTab[50396]++

														field := xd.Number()
														message := xd.ContainingMessage().FullName()
														if prev := r.extensionsByMessage[message][field]; prev != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:551
		_go_fuzz_dep_.CoverTab[50403]++
															err := errors.New("extension number %d is already registered on message %v", field, message)
															err = amendErrorWithCaller(err, prev, xt)
															if !(r == GlobalTypes && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:554
			_go_fuzz_dep_.CoverTab[50404]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:554
			return ignoreConflict(xd, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:554
			// _ = "end of CoverTab[50404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:554
		}()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:554
			_go_fuzz_dep_.CoverTab[50405]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:555
			// _ = "end of CoverTab[50405]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:556
			_go_fuzz_dep_.CoverTab[50406]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:556
			// _ = "end of CoverTab[50406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:556
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:556
		// _ = "end of CoverTab[50403]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:557
		_go_fuzz_dep_.CoverTab[50407]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:557
		// _ = "end of CoverTab[50407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:557
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:557
	// _ = "end of CoverTab[50396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:557
	_go_fuzz_dep_.CoverTab[50397]++

														if err := r.register("extension", xd, xt); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:559
		_go_fuzz_dep_.CoverTab[50408]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:560
		// _ = "end of CoverTab[50408]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:561
		_go_fuzz_dep_.CoverTab[50409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:561
		// _ = "end of CoverTab[50409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:561
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:561
	// _ = "end of CoverTab[50397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:561
	_go_fuzz_dep_.CoverTab[50398]++
														if r.extensionsByMessage == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:562
		_go_fuzz_dep_.CoverTab[50410]++
															r.extensionsByMessage = make(extensionsByMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:563
		// _ = "end of CoverTab[50410]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:564
		_go_fuzz_dep_.CoverTab[50411]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:564
		// _ = "end of CoverTab[50411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:564
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:564
	// _ = "end of CoverTab[50398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:564
	_go_fuzz_dep_.CoverTab[50399]++
														if r.extensionsByMessage[message] == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:565
		_go_fuzz_dep_.CoverTab[50412]++
															r.extensionsByMessage[message] = make(extensionsByNumber)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:566
		// _ = "end of CoverTab[50412]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:567
		_go_fuzz_dep_.CoverTab[50413]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:567
		// _ = "end of CoverTab[50413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:567
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:567
	// _ = "end of CoverTab[50399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:567
	_go_fuzz_dep_.CoverTab[50400]++
														r.extensionsByMessage[message][field] = xt
														r.numExtensions++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:570
	// _ = "end of CoverTab[50400]"
}

func (r *Types) register(kind string, desc protoreflect.Descriptor, typ interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:573
	_go_fuzz_dep_.CoverTab[50414]++
														name := desc.FullName()
														prev := r.typesByName[name]
														if prev != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:576
		_go_fuzz_dep_.CoverTab[50417]++
															err := errors.New("%v %v is already registered", kind, name)
															err = amendErrorWithCaller(err, prev, typ)
															if !(r == GlobalTypes && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:579
			_go_fuzz_dep_.CoverTab[50418]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:579
			return ignoreConflict(desc, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:579
			// _ = "end of CoverTab[50418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:579
		}()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:579
			_go_fuzz_dep_.CoverTab[50419]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:580
			// _ = "end of CoverTab[50419]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:581
			_go_fuzz_dep_.CoverTab[50420]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:581
			// _ = "end of CoverTab[50420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:581
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:581
		// _ = "end of CoverTab[50417]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:582
		_go_fuzz_dep_.CoverTab[50421]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:582
		// _ = "end of CoverTab[50421]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:582
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:582
	// _ = "end of CoverTab[50414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:582
	_go_fuzz_dep_.CoverTab[50415]++
														if r.typesByName == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:583
		_go_fuzz_dep_.CoverTab[50422]++
															r.typesByName = make(typesByName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:584
		// _ = "end of CoverTab[50422]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:585
		_go_fuzz_dep_.CoverTab[50423]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:585
		// _ = "end of CoverTab[50423]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:585
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:585
	// _ = "end of CoverTab[50415]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:585
	_go_fuzz_dep_.CoverTab[50416]++
														r.typesByName[name] = typ
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:587
	// _ = "end of CoverTab[50416]"
}

// FindEnumByName looks up an enum by its full name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:590
// E.g., "google.protobuf.Field.Kind".
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:590
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:590
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:594
func (r *Types) FindEnumByName(enum protoreflect.FullName) (protoreflect.EnumType, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:594
	_go_fuzz_dep_.CoverTab[50424]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:595
		_go_fuzz_dep_.CoverTab[50428]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:596
		// _ = "end of CoverTab[50428]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:597
		_go_fuzz_dep_.CoverTab[50429]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:597
		// _ = "end of CoverTab[50429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:597
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:597
	// _ = "end of CoverTab[50424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:597
	_go_fuzz_dep_.CoverTab[50425]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:598
		_go_fuzz_dep_.CoverTab[50430]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:600
		// _ = "end of CoverTab[50430]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:601
		_go_fuzz_dep_.CoverTab[50431]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:601
		// _ = "end of CoverTab[50431]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:601
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:601
	// _ = "end of CoverTab[50425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:601
	_go_fuzz_dep_.CoverTab[50426]++
														if v := r.typesByName[enum]; v != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:602
		_go_fuzz_dep_.CoverTab[50432]++
															if et, _ := v.(protoreflect.EnumType); et != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:603
			_go_fuzz_dep_.CoverTab[50434]++
																return et, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:604
			// _ = "end of CoverTab[50434]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:605
			_go_fuzz_dep_.CoverTab[50435]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:605
			// _ = "end of CoverTab[50435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:605
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:605
		// _ = "end of CoverTab[50432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:605
		_go_fuzz_dep_.CoverTab[50433]++
															return nil, errors.New("found wrong type: got %v, want enum", typeName(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:606
		// _ = "end of CoverTab[50433]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:607
		_go_fuzz_dep_.CoverTab[50436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:607
		// _ = "end of CoverTab[50436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:607
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:607
	// _ = "end of CoverTab[50426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:607
	_go_fuzz_dep_.CoverTab[50427]++
														return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:608
	// _ = "end of CoverTab[50427]"
}

// FindMessageByName looks up a message by its full name,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:611
// e.g. "google.protobuf.Any".
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:611
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:611
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:615
func (r *Types) FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:615
	_go_fuzz_dep_.CoverTab[50437]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:616
		_go_fuzz_dep_.CoverTab[50441]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:617
		// _ = "end of CoverTab[50441]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:618
		_go_fuzz_dep_.CoverTab[50442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:618
		// _ = "end of CoverTab[50442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:618
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:618
	// _ = "end of CoverTab[50437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:618
	_go_fuzz_dep_.CoverTab[50438]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:619
		_go_fuzz_dep_.CoverTab[50443]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:621
		// _ = "end of CoverTab[50443]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:622
		_go_fuzz_dep_.CoverTab[50444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:622
		// _ = "end of CoverTab[50444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:622
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:622
	// _ = "end of CoverTab[50438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:622
	_go_fuzz_dep_.CoverTab[50439]++
														if v := r.typesByName[message]; v != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:623
		_go_fuzz_dep_.CoverTab[50445]++
															if mt, _ := v.(protoreflect.MessageType); mt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:624
			_go_fuzz_dep_.CoverTab[50447]++
																return mt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:625
			// _ = "end of CoverTab[50447]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:626
			_go_fuzz_dep_.CoverTab[50448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:626
			// _ = "end of CoverTab[50448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:626
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:626
		// _ = "end of CoverTab[50445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:626
		_go_fuzz_dep_.CoverTab[50446]++
															return nil, errors.New("found wrong type: got %v, want message", typeName(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:627
		// _ = "end of CoverTab[50446]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:628
		_go_fuzz_dep_.CoverTab[50449]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:628
		// _ = "end of CoverTab[50449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:628
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:628
	// _ = "end of CoverTab[50439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:628
	_go_fuzz_dep_.CoverTab[50440]++
														return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:629
	// _ = "end of CoverTab[50440]"
}

// FindMessageByURL looks up a message by a URL identifier.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:632
// See documentation on google.protobuf.Any.type_url for the URL format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:632
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:632
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:636
func (r *Types) FindMessageByURL(url string) (protoreflect.MessageType, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:636
	_go_fuzz_dep_.CoverTab[50450]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:639
	if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:639
		_go_fuzz_dep_.CoverTab[50455]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:640
		// _ = "end of CoverTab[50455]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:641
		_go_fuzz_dep_.CoverTab[50456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:641
		// _ = "end of CoverTab[50456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:641
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:641
	// _ = "end of CoverTab[50450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:641
	_go_fuzz_dep_.CoverTab[50451]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:642
		_go_fuzz_dep_.CoverTab[50457]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:644
		// _ = "end of CoverTab[50457]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:645
		_go_fuzz_dep_.CoverTab[50458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:645
		// _ = "end of CoverTab[50458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:645
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:645
	// _ = "end of CoverTab[50451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:645
	_go_fuzz_dep_.CoverTab[50452]++
														message := protoreflect.FullName(url)
														if i := strings.LastIndexByte(url, '/'); i >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:647
		_go_fuzz_dep_.CoverTab[50459]++
															message = message[i+len("/"):]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:648
		// _ = "end of CoverTab[50459]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:649
		_go_fuzz_dep_.CoverTab[50460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:649
		// _ = "end of CoverTab[50460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:649
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:649
	// _ = "end of CoverTab[50452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:649
	_go_fuzz_dep_.CoverTab[50453]++

														if v := r.typesByName[message]; v != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:651
		_go_fuzz_dep_.CoverTab[50461]++
															if mt, _ := v.(protoreflect.MessageType); mt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:652
			_go_fuzz_dep_.CoverTab[50463]++
																return mt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:653
			// _ = "end of CoverTab[50463]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:654
			_go_fuzz_dep_.CoverTab[50464]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:654
			// _ = "end of CoverTab[50464]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:654
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:654
		// _ = "end of CoverTab[50461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:654
		_go_fuzz_dep_.CoverTab[50462]++
															return nil, errors.New("found wrong type: got %v, want message", typeName(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:655
		// _ = "end of CoverTab[50462]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:656
		_go_fuzz_dep_.CoverTab[50465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:656
		// _ = "end of CoverTab[50465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:656
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:656
	// _ = "end of CoverTab[50453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:656
	_go_fuzz_dep_.CoverTab[50454]++
														return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:657
	// _ = "end of CoverTab[50454]"
}

// FindExtensionByName looks up a extension field by the field's full name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:660
// Note that this is the full name of the field as determined by
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:660
// where the extension is declared and is unrelated to the full name of the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:660
// message being extended.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:660
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:660
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:666
func (r *Types) FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:666
	_go_fuzz_dep_.CoverTab[50466]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:667
		_go_fuzz_dep_.CoverTab[50470]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:668
		// _ = "end of CoverTab[50470]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:669
		_go_fuzz_dep_.CoverTab[50471]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:669
		// _ = "end of CoverTab[50471]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:669
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:669
	// _ = "end of CoverTab[50466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:669
	_go_fuzz_dep_.CoverTab[50467]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:670
		_go_fuzz_dep_.CoverTab[50472]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:672
		// _ = "end of CoverTab[50472]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:673
		_go_fuzz_dep_.CoverTab[50473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:673
		// _ = "end of CoverTab[50473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:673
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:673
	// _ = "end of CoverTab[50467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:673
	_go_fuzz_dep_.CoverTab[50468]++
														if v := r.typesByName[field]; v != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:674
		_go_fuzz_dep_.CoverTab[50474]++
															if xt, _ := v.(protoreflect.ExtensionType); xt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:675
			_go_fuzz_dep_.CoverTab[50477]++
																return xt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:676
			// _ = "end of CoverTab[50477]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:677
			_go_fuzz_dep_.CoverTab[50478]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:677
			// _ = "end of CoverTab[50478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:677
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:677
		// _ = "end of CoverTab[50474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:677
		_go_fuzz_dep_.CoverTab[50475]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:685
		if flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:685
			_go_fuzz_dep_.CoverTab[50479]++
																if _, ok := v.(protoreflect.MessageType); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:686
				_go_fuzz_dep_.CoverTab[50480]++
																	field := field.Append(messageset.ExtensionName)
																	if v := r.typesByName[field]; v != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:688
					_go_fuzz_dep_.CoverTab[50481]++
																		if xt, _ := v.(protoreflect.ExtensionType); xt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:689
						_go_fuzz_dep_.CoverTab[50482]++
																			if messageset.IsMessageSetExtension(xt.TypeDescriptor()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:690
							_go_fuzz_dep_.CoverTab[50483]++
																				return xt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:691
							// _ = "end of CoverTab[50483]"
						} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:692
							_go_fuzz_dep_.CoverTab[50484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:692
							// _ = "end of CoverTab[50484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:692
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:692
						// _ = "end of CoverTab[50482]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:693
						_go_fuzz_dep_.CoverTab[50485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:693
						// _ = "end of CoverTab[50485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:693
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:693
					// _ = "end of CoverTab[50481]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:694
					_go_fuzz_dep_.CoverTab[50486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:694
					// _ = "end of CoverTab[50486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:694
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:694
				// _ = "end of CoverTab[50480]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:695
				_go_fuzz_dep_.CoverTab[50487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:695
				// _ = "end of CoverTab[50487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:695
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:695
			// _ = "end of CoverTab[50479]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:696
			_go_fuzz_dep_.CoverTab[50488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:696
			// _ = "end of CoverTab[50488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:696
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:696
		// _ = "end of CoverTab[50475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:696
		_go_fuzz_dep_.CoverTab[50476]++

															return nil, errors.New("found wrong type: got %v, want extension", typeName(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:698
		// _ = "end of CoverTab[50476]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:699
		_go_fuzz_dep_.CoverTab[50489]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:699
		// _ = "end of CoverTab[50489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:699
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:699
	// _ = "end of CoverTab[50468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:699
	_go_fuzz_dep_.CoverTab[50469]++
														return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:700
	// _ = "end of CoverTab[50469]"
}

// FindExtensionByNumber looks up a extension field by the field number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:703
// within some parent message, identified by full name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:703
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:703
// This returns (nil, NotFound) if not found.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:707
func (r *Types) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:707
	_go_fuzz_dep_.CoverTab[50490]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:708
		_go_fuzz_dep_.CoverTab[50494]++
															return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:709
		// _ = "end of CoverTab[50494]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:710
		_go_fuzz_dep_.CoverTab[50495]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:710
		// _ = "end of CoverTab[50495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:710
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:710
	// _ = "end of CoverTab[50490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:710
	_go_fuzz_dep_.CoverTab[50491]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:711
		_go_fuzz_dep_.CoverTab[50496]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:713
		// _ = "end of CoverTab[50496]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:714
		_go_fuzz_dep_.CoverTab[50497]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:714
		// _ = "end of CoverTab[50497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:714
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:714
	// _ = "end of CoverTab[50491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:714
	_go_fuzz_dep_.CoverTab[50492]++
														if xt, ok := r.extensionsByMessage[message][field]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:715
		_go_fuzz_dep_.CoverTab[50498]++
															return xt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:716
		// _ = "end of CoverTab[50498]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:717
		_go_fuzz_dep_.CoverTab[50499]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:717
		// _ = "end of CoverTab[50499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:717
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:717
	// _ = "end of CoverTab[50492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:717
	_go_fuzz_dep_.CoverTab[50493]++
														return nil, NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:718
	// _ = "end of CoverTab[50493]"
}

// NumEnums reports the number of registered enums.
func (r *Types) NumEnums() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:722
	_go_fuzz_dep_.CoverTab[50500]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:723
		_go_fuzz_dep_.CoverTab[50503]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:724
		// _ = "end of CoverTab[50503]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:725
		_go_fuzz_dep_.CoverTab[50504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:725
		// _ = "end of CoverTab[50504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:725
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:725
	// _ = "end of CoverTab[50500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:725
	_go_fuzz_dep_.CoverTab[50501]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:726
		_go_fuzz_dep_.CoverTab[50505]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:728
		// _ = "end of CoverTab[50505]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:729
		_go_fuzz_dep_.CoverTab[50506]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:729
		// _ = "end of CoverTab[50506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:729
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:729
	// _ = "end of CoverTab[50501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:729
	_go_fuzz_dep_.CoverTab[50502]++
														return r.numEnums
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:730
	// _ = "end of CoverTab[50502]"
}

// RangeEnums iterates over all registered enums while f returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:733
// Iteration order is undefined.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:735
func (r *Types) RangeEnums(f func(protoreflect.EnumType) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:735
	_go_fuzz_dep_.CoverTab[50507]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:736
		_go_fuzz_dep_.CoverTab[50510]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:737
		// _ = "end of CoverTab[50510]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:738
		_go_fuzz_dep_.CoverTab[50511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:738
		// _ = "end of CoverTab[50511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:738
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:738
	// _ = "end of CoverTab[50507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:738
	_go_fuzz_dep_.CoverTab[50508]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:739
		_go_fuzz_dep_.CoverTab[50512]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:741
		// _ = "end of CoverTab[50512]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:742
		_go_fuzz_dep_.CoverTab[50513]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:742
		// _ = "end of CoverTab[50513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:742
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:742
	// _ = "end of CoverTab[50508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:742
	_go_fuzz_dep_.CoverTab[50509]++
														for _, typ := range r.typesByName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:743
		_go_fuzz_dep_.CoverTab[50514]++
															if et, ok := typ.(protoreflect.EnumType); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:744
			_go_fuzz_dep_.CoverTab[50515]++
																if !f(et) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:745
				_go_fuzz_dep_.CoverTab[50516]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:746
				// _ = "end of CoverTab[50516]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:747
				_go_fuzz_dep_.CoverTab[50517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:747
				// _ = "end of CoverTab[50517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:747
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:747
			// _ = "end of CoverTab[50515]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:748
			_go_fuzz_dep_.CoverTab[50518]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:748
			// _ = "end of CoverTab[50518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:748
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:748
		// _ = "end of CoverTab[50514]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:749
	// _ = "end of CoverTab[50509]"
}

// NumMessages reports the number of registered messages.
func (r *Types) NumMessages() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:753
	_go_fuzz_dep_.CoverTab[50519]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:754
		_go_fuzz_dep_.CoverTab[50522]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:755
		// _ = "end of CoverTab[50522]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:756
		_go_fuzz_dep_.CoverTab[50523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:756
		// _ = "end of CoverTab[50523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:756
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:756
	// _ = "end of CoverTab[50519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:756
	_go_fuzz_dep_.CoverTab[50520]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:757
		_go_fuzz_dep_.CoverTab[50524]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:759
		// _ = "end of CoverTab[50524]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:760
		_go_fuzz_dep_.CoverTab[50525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:760
		// _ = "end of CoverTab[50525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:760
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:760
	// _ = "end of CoverTab[50520]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:760
	_go_fuzz_dep_.CoverTab[50521]++
														return r.numMessages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:761
	// _ = "end of CoverTab[50521]"
}

// RangeMessages iterates over all registered messages while f returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:764
// Iteration order is undefined.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:766
func (r *Types) RangeMessages(f func(protoreflect.MessageType) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:766
	_go_fuzz_dep_.CoverTab[50526]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:767
		_go_fuzz_dep_.CoverTab[50529]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:768
		// _ = "end of CoverTab[50529]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:769
		_go_fuzz_dep_.CoverTab[50530]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:769
		// _ = "end of CoverTab[50530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:769
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:769
	// _ = "end of CoverTab[50526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:769
	_go_fuzz_dep_.CoverTab[50527]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:770
		_go_fuzz_dep_.CoverTab[50531]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:772
		// _ = "end of CoverTab[50531]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:773
		_go_fuzz_dep_.CoverTab[50532]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:773
		// _ = "end of CoverTab[50532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:773
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:773
	// _ = "end of CoverTab[50527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:773
	_go_fuzz_dep_.CoverTab[50528]++
														for _, typ := range r.typesByName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:774
		_go_fuzz_dep_.CoverTab[50533]++
															if mt, ok := typ.(protoreflect.MessageType); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:775
			_go_fuzz_dep_.CoverTab[50534]++
																if !f(mt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:776
				_go_fuzz_dep_.CoverTab[50535]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:777
				// _ = "end of CoverTab[50535]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:778
				_go_fuzz_dep_.CoverTab[50536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:778
				// _ = "end of CoverTab[50536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:778
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:778
			// _ = "end of CoverTab[50534]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:779
			_go_fuzz_dep_.CoverTab[50537]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:779
			// _ = "end of CoverTab[50537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:779
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:779
		// _ = "end of CoverTab[50533]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:780
	// _ = "end of CoverTab[50528]"
}

// NumExtensions reports the number of registered extensions.
func (r *Types) NumExtensions() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:784
	_go_fuzz_dep_.CoverTab[50538]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:785
		_go_fuzz_dep_.CoverTab[50541]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:786
		// _ = "end of CoverTab[50541]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:787
		_go_fuzz_dep_.CoverTab[50542]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:787
		// _ = "end of CoverTab[50542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:787
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:787
	// _ = "end of CoverTab[50538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:787
	_go_fuzz_dep_.CoverTab[50539]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:788
		_go_fuzz_dep_.CoverTab[50543]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:790
		// _ = "end of CoverTab[50543]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:791
		_go_fuzz_dep_.CoverTab[50544]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:791
		// _ = "end of CoverTab[50544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:791
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:791
	// _ = "end of CoverTab[50539]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:791
	_go_fuzz_dep_.CoverTab[50540]++
														return r.numExtensions
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:792
	// _ = "end of CoverTab[50540]"
}

// RangeExtensions iterates over all registered extensions while f returns true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:795
// Iteration order is undefined.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:797
func (r *Types) RangeExtensions(f func(protoreflect.ExtensionType) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:797
	_go_fuzz_dep_.CoverTab[50545]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:798
		_go_fuzz_dep_.CoverTab[50548]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:799
		// _ = "end of CoverTab[50548]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:800
		_go_fuzz_dep_.CoverTab[50549]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:800
		// _ = "end of CoverTab[50549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:800
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:800
	// _ = "end of CoverTab[50545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:800
	_go_fuzz_dep_.CoverTab[50546]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:801
		_go_fuzz_dep_.CoverTab[50550]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:803
		// _ = "end of CoverTab[50550]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:804
		_go_fuzz_dep_.CoverTab[50551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:804
		// _ = "end of CoverTab[50551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:804
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:804
	// _ = "end of CoverTab[50546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:804
	_go_fuzz_dep_.CoverTab[50547]++
														for _, typ := range r.typesByName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:805
		_go_fuzz_dep_.CoverTab[50552]++
															if xt, ok := typ.(protoreflect.ExtensionType); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:806
			_go_fuzz_dep_.CoverTab[50553]++
																if !f(xt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:807
				_go_fuzz_dep_.CoverTab[50554]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:808
				// _ = "end of CoverTab[50554]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:809
				_go_fuzz_dep_.CoverTab[50555]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:809
				// _ = "end of CoverTab[50555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:809
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:809
			// _ = "end of CoverTab[50553]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:810
			_go_fuzz_dep_.CoverTab[50556]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:810
			// _ = "end of CoverTab[50556]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:810
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:810
		// _ = "end of CoverTab[50552]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:811
	// _ = "end of CoverTab[50547]"
}

// NumExtensionsByMessage reports the number of registered extensions for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:814
// a given message type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:816
func (r *Types) NumExtensionsByMessage(message protoreflect.FullName) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:816
	_go_fuzz_dep_.CoverTab[50557]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:817
		_go_fuzz_dep_.CoverTab[50560]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:818
		// _ = "end of CoverTab[50560]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:819
		_go_fuzz_dep_.CoverTab[50561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:819
		// _ = "end of CoverTab[50561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:819
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:819
	// _ = "end of CoverTab[50557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:819
	_go_fuzz_dep_.CoverTab[50558]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:820
		_go_fuzz_dep_.CoverTab[50562]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:822
		// _ = "end of CoverTab[50562]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:823
		_go_fuzz_dep_.CoverTab[50563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:823
		// _ = "end of CoverTab[50563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:823
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:823
	// _ = "end of CoverTab[50558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:823
	_go_fuzz_dep_.CoverTab[50559]++
														return len(r.extensionsByMessage[message])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:824
	// _ = "end of CoverTab[50559]"
}

// RangeExtensionsByMessage iterates over all registered extensions filtered
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:827
// by a given message type while f returns true. Iteration order is undefined.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:829
func (r *Types) RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:829
	_go_fuzz_dep_.CoverTab[50564]++
														if r == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:830
		_go_fuzz_dep_.CoverTab[50567]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:831
		// _ = "end of CoverTab[50567]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:832
		_go_fuzz_dep_.CoverTab[50568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:832
		// _ = "end of CoverTab[50568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:832
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:832
	// _ = "end of CoverTab[50564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:832
	_go_fuzz_dep_.CoverTab[50565]++
														if r == GlobalTypes {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:833
		_go_fuzz_dep_.CoverTab[50569]++
															globalMutex.RLock()
															defer globalMutex.RUnlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:835
		// _ = "end of CoverTab[50569]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:836
		_go_fuzz_dep_.CoverTab[50570]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:836
		// _ = "end of CoverTab[50570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:836
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:836
	// _ = "end of CoverTab[50565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:836
	_go_fuzz_dep_.CoverTab[50566]++
														for _, xt := range r.extensionsByMessage[message] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:837
		_go_fuzz_dep_.CoverTab[50571]++
															if !f(xt) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:838
			_go_fuzz_dep_.CoverTab[50572]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:839
			// _ = "end of CoverTab[50572]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:840
			_go_fuzz_dep_.CoverTab[50573]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:840
			// _ = "end of CoverTab[50573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:840
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:840
		// _ = "end of CoverTab[50571]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:841
	// _ = "end of CoverTab[50566]"
}

func typeName(t interface{}) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:844
	_go_fuzz_dep_.CoverTab[50574]++
														switch t.(type) {
	case protoreflect.EnumType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:846
		_go_fuzz_dep_.CoverTab[50575]++
															return "enum"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:847
		// _ = "end of CoverTab[50575]"
	case protoreflect.MessageType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:848
		_go_fuzz_dep_.CoverTab[50576]++
															return "message"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:849
		// _ = "end of CoverTab[50576]"
	case protoreflect.ExtensionType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:850
		_go_fuzz_dep_.CoverTab[50577]++
															return "extension"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:851
		// _ = "end of CoverTab[50577]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:852
		_go_fuzz_dep_.CoverTab[50578]++
															return fmt.Sprintf("%T", t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:853
		// _ = "end of CoverTab[50578]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:854
	// _ = "end of CoverTab[50574]"
}

func amendErrorWithCaller(err error, prev, curr interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:857
	_go_fuzz_dep_.CoverTab[50579]++
														prevPkg := goPackage(prev)
														currPkg := goPackage(curr)
														if prevPkg == "" || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		_go_fuzz_dep_.CoverTab[50581]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		return currPkg == ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		// _ = "end of CoverTab[50581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		_go_fuzz_dep_.CoverTab[50582]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		return prevPkg == currPkg
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		// _ = "end of CoverTab[50582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:860
		_go_fuzz_dep_.CoverTab[50583]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:861
		// _ = "end of CoverTab[50583]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:862
		_go_fuzz_dep_.CoverTab[50584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:862
		// _ = "end of CoverTab[50584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:862
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:862
	// _ = "end of CoverTab[50579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:862
	_go_fuzz_dep_.CoverTab[50580]++
														return errors.New("%s\n\tpreviously from: %q\n\tcurrently from:  %q", err, prevPkg, currPkg)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:863
	// _ = "end of CoverTab[50580]"
}

func goPackage(v interface{}) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:866
	_go_fuzz_dep_.CoverTab[50585]++
														switch d := v.(type) {
	case protoreflect.EnumType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:868
		_go_fuzz_dep_.CoverTab[50589]++
															v = d.Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:869
		// _ = "end of CoverTab[50589]"
	case protoreflect.MessageType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:870
		_go_fuzz_dep_.CoverTab[50590]++
															v = d.Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:871
		// _ = "end of CoverTab[50590]"
	case protoreflect.ExtensionType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:872
		_go_fuzz_dep_.CoverTab[50591]++
															v = d.TypeDescriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:873
		// _ = "end of CoverTab[50591]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:874
	// _ = "end of CoverTab[50585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:874
	_go_fuzz_dep_.CoverTab[50586]++
														if d, ok := v.(protoreflect.Descriptor); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:875
		_go_fuzz_dep_.CoverTab[50592]++
															v = d.ParentFile()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:876
		// _ = "end of CoverTab[50592]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:877
		_go_fuzz_dep_.CoverTab[50593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:877
		// _ = "end of CoverTab[50593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:877
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:877
	// _ = "end of CoverTab[50586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:877
	_go_fuzz_dep_.CoverTab[50587]++
														if d, ok := v.(interface{ GoPackagePath() string }); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:878
		_go_fuzz_dep_.CoverTab[50594]++
															return d.GoPackagePath()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:879
		// _ = "end of CoverTab[50594]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:880
		_go_fuzz_dep_.CoverTab[50595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:880
		// _ = "end of CoverTab[50595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:880
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:880
	// _ = "end of CoverTab[50587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:880
	_go_fuzz_dep_.CoverTab[50588]++
														return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:881
	// _ = "end of CoverTab[50588]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:882
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoregistry/registry.go:882
var _ = _go_fuzz_dep_.CoverTab
