// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
package filedesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:5
)

import (
	"sync"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// fileRaw is a data struct used when initializing a file descriptor from
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:16
// a raw FileDescriptorProto.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:18
type fileRaw struct {
	builder		Builder
	allEnums	[]Enum
	allMessages	[]Message
	allExtensions	[]Extension
	allServices	[]Service
}

func newRawFile(db Builder) *File {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:26
	_go_fuzz_dep_.CoverTab[52636]++
													fd := &File{fileRaw: fileRaw{builder: db}}
													fd.initDecls(db.NumEnums, db.NumMessages, db.NumExtensions, db.NumServices)
													fd.unmarshalSeed(db.RawDescriptor)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:33
	for i := range fd.allExtensions {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:33
		_go_fuzz_dep_.CoverTab[52638]++
														xd := &fd.allExtensions[i]
														xd.L1.Extendee = fd.resolveMessageDependency(xd.L1.Extendee, listExtTargets, int32(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:35
		// _ = "end of CoverTab[52638]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:36
	// _ = "end of CoverTab[52636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:36
	_go_fuzz_dep_.CoverTab[52637]++

													fd.checkDecls()
													return fd
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:39
	// _ = "end of CoverTab[52637]"
}

// initDecls pre-allocates slices for the exact number of enums, messages
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:42
// (including map entries), extensions, and services declared in the proto file.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:42
// This is done to avoid regrowing the slice, which would change the address
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:42
// for any previously seen declaration.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:42
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:42
// The alloc methods "allocates" slices by pulling from the capacity.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:48
func (fd *File) initDecls(numEnums, numMessages, numExtensions, numServices int32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:48
	_go_fuzz_dep_.CoverTab[52639]++
													fd.allEnums = make([]Enum, 0, numEnums)
													fd.allMessages = make([]Message, 0, numMessages)
													fd.allExtensions = make([]Extension, 0, numExtensions)
													fd.allServices = make([]Service, 0, numServices)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:52
	// _ = "end of CoverTab[52639]"
}

func (fd *File) allocEnums(n int) []Enum {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:55
	_go_fuzz_dep_.CoverTab[52640]++
													total := len(fd.allEnums)
													es := fd.allEnums[total : total+n]
													fd.allEnums = fd.allEnums[:total+n]
													return es
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:59
	// _ = "end of CoverTab[52640]"
}
func (fd *File) allocMessages(n int) []Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:61
	_go_fuzz_dep_.CoverTab[52641]++
													total := len(fd.allMessages)
													ms := fd.allMessages[total : total+n]
													fd.allMessages = fd.allMessages[:total+n]
													return ms
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:65
	// _ = "end of CoverTab[52641]"
}
func (fd *File) allocExtensions(n int) []Extension {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:67
	_go_fuzz_dep_.CoverTab[52642]++
													total := len(fd.allExtensions)
													xs := fd.allExtensions[total : total+n]
													fd.allExtensions = fd.allExtensions[:total+n]
													return xs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:71
	// _ = "end of CoverTab[52642]"
}
func (fd *File) allocServices(n int) []Service {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:73
	_go_fuzz_dep_.CoverTab[52643]++
													total := len(fd.allServices)
													xs := fd.allServices[total : total+n]
													fd.allServices = fd.allServices[:total+n]
													return xs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:77
	// _ = "end of CoverTab[52643]"
}

// checkDecls performs a sanity check that the expected number of expected
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:80
// declarations matches the number that were found in the descriptor proto.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:82
func (fd *File) checkDecls() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:82
	_go_fuzz_dep_.CoverTab[52644]++
													switch {
	case len(fd.allEnums) != cap(fd.allEnums):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:84
		_go_fuzz_dep_.CoverTab[52646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:84
		// _ = "end of CoverTab[52646]"
	case len(fd.allMessages) != cap(fd.allMessages):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:85
		_go_fuzz_dep_.CoverTab[52647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:85
		// _ = "end of CoverTab[52647]"
	case len(fd.allExtensions) != cap(fd.allExtensions):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:86
		_go_fuzz_dep_.CoverTab[52648]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:86
		// _ = "end of CoverTab[52648]"
	case len(fd.allServices) != cap(fd.allServices):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:87
		_go_fuzz_dep_.CoverTab[52649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:87
		// _ = "end of CoverTab[52649]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:88
		_go_fuzz_dep_.CoverTab[52650]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:89
		// _ = "end of CoverTab[52650]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:90
	// _ = "end of CoverTab[52644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:90
	_go_fuzz_dep_.CoverTab[52645]++
													panic("mismatching cardinality")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:91
	// _ = "end of CoverTab[52645]"
}

func (fd *File) unmarshalSeed(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:94
	_go_fuzz_dep_.CoverTab[52651]++
													sb := getBuilder()
													defer putBuilder(sb)

													var prevField protoreflect.FieldNumber
													var numEnums, numMessages, numExtensions, numServices int
													var posEnums, posMessages, posExtensions, posServices int
													b0 := b
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:102
		_go_fuzz_dep_.CoverTab[52661]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:106
			_go_fuzz_dep_.CoverTab[52662]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.FileDescriptorProto_Syntax_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:110
				_go_fuzz_dep_.CoverTab[52665]++
																	switch string(v) {
				case "proto2":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:112
					_go_fuzz_dep_.CoverTab[52677]++
																		fd.L1.Syntax = protoreflect.Proto2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:113
					// _ = "end of CoverTab[52677]"
				case "proto3":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:114
					_go_fuzz_dep_.CoverTab[52678]++
																		fd.L1.Syntax = protoreflect.Proto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:115
					// _ = "end of CoverTab[52678]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:116
					_go_fuzz_dep_.CoverTab[52679]++
																		panic("invalid syntax")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:117
					// _ = "end of CoverTab[52679]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:118
				// _ = "end of CoverTab[52665]"
			case genid.FileDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:119
				_go_fuzz_dep_.CoverTab[52666]++
																	fd.L1.Path = sb.MakeString(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:120
				// _ = "end of CoverTab[52666]"
			case genid.FileDescriptorProto_Package_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:121
				_go_fuzz_dep_.CoverTab[52667]++
																	fd.L1.Package = protoreflect.FullName(sb.MakeString(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:122
				// _ = "end of CoverTab[52667]"
			case genid.FileDescriptorProto_EnumType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:123
				_go_fuzz_dep_.CoverTab[52668]++
																	if prevField != genid.FileDescriptorProto_EnumType_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:124
					_go_fuzz_dep_.CoverTab[52680]++
																		if numEnums > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:125
						_go_fuzz_dep_.CoverTab[52682]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:126
						// _ = "end of CoverTab[52682]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:127
						_go_fuzz_dep_.CoverTab[52683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:127
						// _ = "end of CoverTab[52683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:127
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:127
					// _ = "end of CoverTab[52680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:127
					_go_fuzz_dep_.CoverTab[52681]++
																		posEnums = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:128
					// _ = "end of CoverTab[52681]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:129
					_go_fuzz_dep_.CoverTab[52684]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:129
					// _ = "end of CoverTab[52684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:129
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:129
				// _ = "end of CoverTab[52668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:129
				_go_fuzz_dep_.CoverTab[52669]++
																	numEnums++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:130
				// _ = "end of CoverTab[52669]"
			case genid.FileDescriptorProto_MessageType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:131
				_go_fuzz_dep_.CoverTab[52670]++
																	if prevField != genid.FileDescriptorProto_MessageType_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:132
					_go_fuzz_dep_.CoverTab[52685]++
																		if numMessages > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:133
						_go_fuzz_dep_.CoverTab[52687]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:134
						// _ = "end of CoverTab[52687]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:135
						_go_fuzz_dep_.CoverTab[52688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:135
						// _ = "end of CoverTab[52688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:135
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:135
					// _ = "end of CoverTab[52685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:135
					_go_fuzz_dep_.CoverTab[52686]++
																		posMessages = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:136
					// _ = "end of CoverTab[52686]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:137
					_go_fuzz_dep_.CoverTab[52689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:137
					// _ = "end of CoverTab[52689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:137
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:137
				// _ = "end of CoverTab[52670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:137
				_go_fuzz_dep_.CoverTab[52671]++
																	numMessages++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:138
				// _ = "end of CoverTab[52671]"
			case genid.FileDescriptorProto_Extension_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:139
				_go_fuzz_dep_.CoverTab[52672]++
																	if prevField != genid.FileDescriptorProto_Extension_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:140
					_go_fuzz_dep_.CoverTab[52690]++
																		if numExtensions > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:141
						_go_fuzz_dep_.CoverTab[52692]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:142
						// _ = "end of CoverTab[52692]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:143
						_go_fuzz_dep_.CoverTab[52693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:143
						// _ = "end of CoverTab[52693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:143
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:143
					// _ = "end of CoverTab[52690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:143
					_go_fuzz_dep_.CoverTab[52691]++
																		posExtensions = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:144
					// _ = "end of CoverTab[52691]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:145
					_go_fuzz_dep_.CoverTab[52694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:145
					// _ = "end of CoverTab[52694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:145
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:145
				// _ = "end of CoverTab[52672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:145
				_go_fuzz_dep_.CoverTab[52673]++
																	numExtensions++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:146
				// _ = "end of CoverTab[52673]"
			case genid.FileDescriptorProto_Service_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:147
				_go_fuzz_dep_.CoverTab[52674]++
																	if prevField != genid.FileDescriptorProto_Service_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:148
					_go_fuzz_dep_.CoverTab[52695]++
																		if numServices > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:149
						_go_fuzz_dep_.CoverTab[52697]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:150
						// _ = "end of CoverTab[52697]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:151
						_go_fuzz_dep_.CoverTab[52698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:151
						// _ = "end of CoverTab[52698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:151
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:151
					// _ = "end of CoverTab[52695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:151
					_go_fuzz_dep_.CoverTab[52696]++
																		posServices = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:152
					// _ = "end of CoverTab[52696]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:153
					_go_fuzz_dep_.CoverTab[52699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:153
					// _ = "end of CoverTab[52699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:153
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:153
				// _ = "end of CoverTab[52674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:153
				_go_fuzz_dep_.CoverTab[52675]++
																	numServices++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:154
				// _ = "end of CoverTab[52675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:154
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:154
				_go_fuzz_dep_.CoverTab[52676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:154
				// _ = "end of CoverTab[52676]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:155
			// _ = "end of CoverTab[52662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:155
			_go_fuzz_dep_.CoverTab[52663]++
																prevField = num
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:156
			// _ = "end of CoverTab[52663]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:157
			_go_fuzz_dep_.CoverTab[52664]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
																prevField = -1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:160
			// _ = "end of CoverTab[52664]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:161
		// _ = "end of CoverTab[52661]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:162
	// _ = "end of CoverTab[52651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:162
	_go_fuzz_dep_.CoverTab[52652]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:165
	if fd.L1.Syntax == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:165
		_go_fuzz_dep_.CoverTab[52700]++
															fd.L1.Syntax = protoreflect.Proto2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:166
		// _ = "end of CoverTab[52700]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:167
		_go_fuzz_dep_.CoverTab[52701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:167
		// _ = "end of CoverTab[52701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:167
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:167
	// _ = "end of CoverTab[52652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:167
	_go_fuzz_dep_.CoverTab[52653]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:171
	if numEnums > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:171
		_go_fuzz_dep_.CoverTab[52702]++
															fd.L1.Enums.List = fd.allocEnums(numEnums)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:172
		// _ = "end of CoverTab[52702]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:173
		_go_fuzz_dep_.CoverTab[52703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:173
		// _ = "end of CoverTab[52703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:173
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:173
	// _ = "end of CoverTab[52653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:173
	_go_fuzz_dep_.CoverTab[52654]++
														if numMessages > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:174
		_go_fuzz_dep_.CoverTab[52704]++
															fd.L1.Messages.List = fd.allocMessages(numMessages)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:175
		// _ = "end of CoverTab[52704]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:176
		_go_fuzz_dep_.CoverTab[52705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:176
		// _ = "end of CoverTab[52705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:176
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:176
	// _ = "end of CoverTab[52654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:176
	_go_fuzz_dep_.CoverTab[52655]++
														if numExtensions > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:177
		_go_fuzz_dep_.CoverTab[52706]++
															fd.L1.Extensions.List = fd.allocExtensions(numExtensions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:178
		// _ = "end of CoverTab[52706]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:179
		_go_fuzz_dep_.CoverTab[52707]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:179
		// _ = "end of CoverTab[52707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:179
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:179
	// _ = "end of CoverTab[52655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:179
	_go_fuzz_dep_.CoverTab[52656]++
														if numServices > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:180
		_go_fuzz_dep_.CoverTab[52708]++
															fd.L1.Services.List = fd.allocServices(numServices)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:181
		// _ = "end of CoverTab[52708]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:182
		_go_fuzz_dep_.CoverTab[52709]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:182
		// _ = "end of CoverTab[52709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:182
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:182
	// _ = "end of CoverTab[52656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:182
	_go_fuzz_dep_.CoverTab[52657]++

														if numEnums > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:184
		_go_fuzz_dep_.CoverTab[52710]++
															b := b0[posEnums:]
															for i := range fd.L1.Enums.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:186
			_go_fuzz_dep_.CoverTab[52711]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																fd.L1.Enums.List[i].unmarshalSeed(v, sb, fd, fd, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:190
			// _ = "end of CoverTab[52711]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:191
		// _ = "end of CoverTab[52710]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:192
		_go_fuzz_dep_.CoverTab[52712]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:192
		// _ = "end of CoverTab[52712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:192
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:192
	// _ = "end of CoverTab[52657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:192
	_go_fuzz_dep_.CoverTab[52658]++
														if numMessages > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:193
		_go_fuzz_dep_.CoverTab[52713]++
															b := b0[posMessages:]
															for i := range fd.L1.Messages.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:195
			_go_fuzz_dep_.CoverTab[52714]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																fd.L1.Messages.List[i].unmarshalSeed(v, sb, fd, fd, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:199
			// _ = "end of CoverTab[52714]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:200
		// _ = "end of CoverTab[52713]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:201
		_go_fuzz_dep_.CoverTab[52715]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:201
		// _ = "end of CoverTab[52715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:201
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:201
	// _ = "end of CoverTab[52658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:201
	_go_fuzz_dep_.CoverTab[52659]++
														if numExtensions > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:202
		_go_fuzz_dep_.CoverTab[52716]++
															b := b0[posExtensions:]
															for i := range fd.L1.Extensions.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:204
			_go_fuzz_dep_.CoverTab[52717]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																fd.L1.Extensions.List[i].unmarshalSeed(v, sb, fd, fd, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:208
			// _ = "end of CoverTab[52717]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:209
		// _ = "end of CoverTab[52716]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:210
		_go_fuzz_dep_.CoverTab[52718]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:210
		// _ = "end of CoverTab[52718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:210
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:210
	// _ = "end of CoverTab[52659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:210
	_go_fuzz_dep_.CoverTab[52660]++
														if numServices > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:211
		_go_fuzz_dep_.CoverTab[52719]++
															b := b0[posServices:]
															for i := range fd.L1.Services.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:213
			_go_fuzz_dep_.CoverTab[52720]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																fd.L1.Services.List[i].unmarshalSeed(v, sb, fd, fd, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:217
			// _ = "end of CoverTab[52720]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:218
		// _ = "end of CoverTab[52719]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:219
		_go_fuzz_dep_.CoverTab[52721]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:219
		// _ = "end of CoverTab[52721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:219
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:219
	// _ = "end of CoverTab[52660]"
}

func (ed *Enum) unmarshalSeed(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:222
	_go_fuzz_dep_.CoverTab[52722]++
														ed.L0.ParentFile = pf
														ed.L0.Parent = pd
														ed.L0.Index = i

														var numValues int
														for b := b; len(b) > 0; {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:228
		_go_fuzz_dep_.CoverTab[52725]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:232
			_go_fuzz_dep_.CoverTab[52726]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.EnumDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:236
				_go_fuzz_dep_.CoverTab[52728]++
																	ed.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:237
				// _ = "end of CoverTab[52728]"
			case genid.EnumDescriptorProto_Value_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:238
				_go_fuzz_dep_.CoverTab[52729]++
																	numValues++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:239
				// _ = "end of CoverTab[52729]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:239
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:239
				_go_fuzz_dep_.CoverTab[52730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:239
				// _ = "end of CoverTab[52730]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:240
			// _ = "end of CoverTab[52726]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:241
			_go_fuzz_dep_.CoverTab[52727]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:243
			// _ = "end of CoverTab[52727]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:244
		// _ = "end of CoverTab[52725]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:245
	// _ = "end of CoverTab[52722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:245
	_go_fuzz_dep_.CoverTab[52723]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:249
	if pd != pf {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:249
		_go_fuzz_dep_.CoverTab[52731]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:250
		// _ = "end of CoverTab[52731]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:251
		_go_fuzz_dep_.CoverTab[52732]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:251
		// _ = "end of CoverTab[52732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:251
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:251
	// _ = "end of CoverTab[52723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:251
	_go_fuzz_dep_.CoverTab[52724]++
														ed.L1.eagerValues = true
														ed.L2 = new(EnumL2)
														ed.L2.Values.List = make([]EnumValue, numValues)
														for i := 0; len(b) > 0; {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:255
		_go_fuzz_dep_.CoverTab[52733]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:259
			_go_fuzz_dep_.CoverTab[52734]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.EnumDescriptorProto_Value_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:263
				_go_fuzz_dep_.CoverTab[52736]++
																	ed.L2.Values.List[i].unmarshalFull(v, sb, pf, ed, i)
																	i++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:265
				// _ = "end of CoverTab[52736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:265
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:265
				_go_fuzz_dep_.CoverTab[52737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:265
				// _ = "end of CoverTab[52737]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:266
			// _ = "end of CoverTab[52734]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:267
			_go_fuzz_dep_.CoverTab[52735]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:269
			// _ = "end of CoverTab[52735]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:270
		// _ = "end of CoverTab[52733]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:271
	// _ = "end of CoverTab[52724]"
}

func (md *Message) unmarshalSeed(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:274
	_go_fuzz_dep_.CoverTab[52738]++
														md.L0.ParentFile = pf
														md.L0.Parent = pd
														md.L0.Index = i

														var prevField protoreflect.FieldNumber
														var numEnums, numMessages, numExtensions int
														var posEnums, posMessages, posExtensions int
														b0 := b
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:283
		_go_fuzz_dep_.CoverTab[52745]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:287
			_go_fuzz_dep_.CoverTab[52746]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.DescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:291
				_go_fuzz_dep_.CoverTab[52749]++
																	md.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:292
				// _ = "end of CoverTab[52749]"
			case genid.DescriptorProto_EnumType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:293
				_go_fuzz_dep_.CoverTab[52750]++
																	if prevField != genid.DescriptorProto_EnumType_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:294
					_go_fuzz_dep_.CoverTab[52758]++
																		if numEnums > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:295
						_go_fuzz_dep_.CoverTab[52760]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:296
						// _ = "end of CoverTab[52760]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:297
						_go_fuzz_dep_.CoverTab[52761]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:297
						// _ = "end of CoverTab[52761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:297
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:297
					// _ = "end of CoverTab[52758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:297
					_go_fuzz_dep_.CoverTab[52759]++
																		posEnums = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:298
					// _ = "end of CoverTab[52759]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:299
					_go_fuzz_dep_.CoverTab[52762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:299
					// _ = "end of CoverTab[52762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:299
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:299
				// _ = "end of CoverTab[52750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:299
				_go_fuzz_dep_.CoverTab[52751]++
																	numEnums++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:300
				// _ = "end of CoverTab[52751]"
			case genid.DescriptorProto_NestedType_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:301
				_go_fuzz_dep_.CoverTab[52752]++
																	if prevField != genid.DescriptorProto_NestedType_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:302
					_go_fuzz_dep_.CoverTab[52763]++
																		if numMessages > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:303
						_go_fuzz_dep_.CoverTab[52765]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:304
						// _ = "end of CoverTab[52765]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:305
						_go_fuzz_dep_.CoverTab[52766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:305
						// _ = "end of CoverTab[52766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:305
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:305
					// _ = "end of CoverTab[52763]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:305
					_go_fuzz_dep_.CoverTab[52764]++
																		posMessages = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:306
					// _ = "end of CoverTab[52764]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:307
					_go_fuzz_dep_.CoverTab[52767]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:307
					// _ = "end of CoverTab[52767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:307
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:307
				// _ = "end of CoverTab[52752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:307
				_go_fuzz_dep_.CoverTab[52753]++
																	numMessages++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:308
				// _ = "end of CoverTab[52753]"
			case genid.DescriptorProto_Extension_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:309
				_go_fuzz_dep_.CoverTab[52754]++
																	if prevField != genid.DescriptorProto_Extension_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:310
					_go_fuzz_dep_.CoverTab[52768]++
																		if numExtensions > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:311
						_go_fuzz_dep_.CoverTab[52770]++
																			panic("non-contiguous repeated field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:312
						// _ = "end of CoverTab[52770]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:313
						_go_fuzz_dep_.CoverTab[52771]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:313
						// _ = "end of CoverTab[52771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:313
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:313
					// _ = "end of CoverTab[52768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:313
					_go_fuzz_dep_.CoverTab[52769]++
																		posExtensions = len(b0) - len(b) - n - m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:314
					// _ = "end of CoverTab[52769]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:315
					_go_fuzz_dep_.CoverTab[52772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:315
					// _ = "end of CoverTab[52772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:315
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:315
				// _ = "end of CoverTab[52754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:315
				_go_fuzz_dep_.CoverTab[52755]++
																	numExtensions++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:316
				// _ = "end of CoverTab[52755]"
			case genid.DescriptorProto_Options_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:317
				_go_fuzz_dep_.CoverTab[52756]++
																	md.unmarshalSeedOptions(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:318
				// _ = "end of CoverTab[52756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:318
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:318
				_go_fuzz_dep_.CoverTab[52757]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:318
				// _ = "end of CoverTab[52757]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:319
			// _ = "end of CoverTab[52746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:319
			_go_fuzz_dep_.CoverTab[52747]++
																prevField = num
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:320
			// _ = "end of CoverTab[52747]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:321
			_go_fuzz_dep_.CoverTab[52748]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
																prevField = -1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:324
			// _ = "end of CoverTab[52748]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:325
		// _ = "end of CoverTab[52745]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:326
	// _ = "end of CoverTab[52738]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:326
	_go_fuzz_dep_.CoverTab[52739]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:330
	if numEnums > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:330
		_go_fuzz_dep_.CoverTab[52773]++
															md.L1.Enums.List = pf.allocEnums(numEnums)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:331
		// _ = "end of CoverTab[52773]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:332
		_go_fuzz_dep_.CoverTab[52774]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:332
		// _ = "end of CoverTab[52774]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:332
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:332
	// _ = "end of CoverTab[52739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:332
	_go_fuzz_dep_.CoverTab[52740]++
														if numMessages > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:333
		_go_fuzz_dep_.CoverTab[52775]++
															md.L1.Messages.List = pf.allocMessages(numMessages)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:334
		// _ = "end of CoverTab[52775]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:335
		_go_fuzz_dep_.CoverTab[52776]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:335
		// _ = "end of CoverTab[52776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:335
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:335
	// _ = "end of CoverTab[52740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:335
	_go_fuzz_dep_.CoverTab[52741]++
														if numExtensions > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:336
		_go_fuzz_dep_.CoverTab[52777]++
															md.L1.Extensions.List = pf.allocExtensions(numExtensions)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:337
		// _ = "end of CoverTab[52777]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:338
		_go_fuzz_dep_.CoverTab[52778]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:338
		// _ = "end of CoverTab[52778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:338
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:338
	// _ = "end of CoverTab[52741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:338
	_go_fuzz_dep_.CoverTab[52742]++

														if numEnums > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:340
		_go_fuzz_dep_.CoverTab[52779]++
															b := b0[posEnums:]
															for i := range md.L1.Enums.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:342
			_go_fuzz_dep_.CoverTab[52780]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																md.L1.Enums.List[i].unmarshalSeed(v, sb, pf, md, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:346
			// _ = "end of CoverTab[52780]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:347
		// _ = "end of CoverTab[52779]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:348
		_go_fuzz_dep_.CoverTab[52781]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:348
		// _ = "end of CoverTab[52781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:348
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:348
	// _ = "end of CoverTab[52742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:348
	_go_fuzz_dep_.CoverTab[52743]++
														if numMessages > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:349
		_go_fuzz_dep_.CoverTab[52782]++
															b := b0[posMessages:]
															for i := range md.L1.Messages.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:351
			_go_fuzz_dep_.CoverTab[52783]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																md.L1.Messages.List[i].unmarshalSeed(v, sb, pf, md, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:355
			// _ = "end of CoverTab[52783]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:356
		// _ = "end of CoverTab[52782]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:357
		_go_fuzz_dep_.CoverTab[52784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:357
		// _ = "end of CoverTab[52784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:357
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:357
	// _ = "end of CoverTab[52743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:357
	_go_fuzz_dep_.CoverTab[52744]++
														if numExtensions > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:358
		_go_fuzz_dep_.CoverTab[52785]++
															b := b0[posExtensions:]
															for i := range md.L1.Extensions.List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:360
			_go_fuzz_dep_.CoverTab[52786]++
																_, n := protowire.ConsumeVarint(b)
																v, m := protowire.ConsumeBytes(b[n:])
																md.L1.Extensions.List[i].unmarshalSeed(v, sb, pf, md, i)
																b = b[n+m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:364
			// _ = "end of CoverTab[52786]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:365
		// _ = "end of CoverTab[52785]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:366
		_go_fuzz_dep_.CoverTab[52787]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:366
		// _ = "end of CoverTab[52787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:366
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:366
	// _ = "end of CoverTab[52744]"
}

func (md *Message) unmarshalSeedOptions(b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:369
	_go_fuzz_dep_.CoverTab[52788]++
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:370
		_go_fuzz_dep_.CoverTab[52789]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:374
			_go_fuzz_dep_.CoverTab[52790]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.MessageOptions_MapEntry_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:378
				_go_fuzz_dep_.CoverTab[52792]++
																	md.L1.IsMapEntry = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:379
				// _ = "end of CoverTab[52792]"
			case genid.MessageOptions_MessageSetWireFormat_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:380
				_go_fuzz_dep_.CoverTab[52793]++
																	md.L1.IsMessageSet = protowire.DecodeBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:381
				// _ = "end of CoverTab[52793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:381
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:381
				_go_fuzz_dep_.CoverTab[52794]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:381
				// _ = "end of CoverTab[52794]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:382
			// _ = "end of CoverTab[52790]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:383
			_go_fuzz_dep_.CoverTab[52791]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:385
			// _ = "end of CoverTab[52791]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:386
		// _ = "end of CoverTab[52789]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:387
	// _ = "end of CoverTab[52788]"
}

func (xd *Extension) unmarshalSeed(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:390
	_go_fuzz_dep_.CoverTab[52795]++
														xd.L0.ParentFile = pf
														xd.L0.Parent = pd
														xd.L0.Index = i

														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:395
		_go_fuzz_dep_.CoverTab[52796]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:399
			_go_fuzz_dep_.CoverTab[52797]++
																v, m := protowire.ConsumeVarint(b)
																b = b[m:]
																switch num {
			case genid.FieldDescriptorProto_Number_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:403
				_go_fuzz_dep_.CoverTab[52800]++
																	xd.L1.Number = protoreflect.FieldNumber(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:404
				// _ = "end of CoverTab[52800]"
			case genid.FieldDescriptorProto_Label_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:405
				_go_fuzz_dep_.CoverTab[52801]++
																	xd.L1.Cardinality = protoreflect.Cardinality(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:406
				// _ = "end of CoverTab[52801]"
			case genid.FieldDescriptorProto_Type_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:407
				_go_fuzz_dep_.CoverTab[52802]++
																	xd.L1.Kind = protoreflect.Kind(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:408
				// _ = "end of CoverTab[52802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:408
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:408
				_go_fuzz_dep_.CoverTab[52803]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:408
				// _ = "end of CoverTab[52803]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:409
			// _ = "end of CoverTab[52797]"
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:410
			_go_fuzz_dep_.CoverTab[52798]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.FieldDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:414
				_go_fuzz_dep_.CoverTab[52804]++
																	xd.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:415
				// _ = "end of CoverTab[52804]"
			case genid.FieldDescriptorProto_Extendee_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:416
				_go_fuzz_dep_.CoverTab[52805]++
																	xd.L1.Extendee = PlaceholderMessage(makeFullName(sb, v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:417
				// _ = "end of CoverTab[52805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:417
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:417
				_go_fuzz_dep_.CoverTab[52806]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:417
				// _ = "end of CoverTab[52806]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:418
			// _ = "end of CoverTab[52798]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:419
			_go_fuzz_dep_.CoverTab[52799]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:421
			// _ = "end of CoverTab[52799]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:422
		// _ = "end of CoverTab[52796]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:423
	// _ = "end of CoverTab[52795]"
}

func (sd *Service) unmarshalSeed(b []byte, sb *strs.Builder, pf *File, pd protoreflect.Descriptor, i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:426
	_go_fuzz_dep_.CoverTab[52807]++
														sd.L0.ParentFile = pf
														sd.L0.Parent = pd
														sd.L0.Index = i

														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:431
		_go_fuzz_dep_.CoverTab[52808]++
															num, typ, n := protowire.ConsumeTag(b)
															b = b[n:]
															switch typ {
		case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:435
			_go_fuzz_dep_.CoverTab[52809]++
																v, m := protowire.ConsumeBytes(b)
																b = b[m:]
																switch num {
			case genid.ServiceDescriptorProto_Name_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:439
				_go_fuzz_dep_.CoverTab[52811]++
																	sd.L0.FullName = appendFullName(sb, pd.FullName(), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:440
				// _ = "end of CoverTab[52811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:440
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:440
				_go_fuzz_dep_.CoverTab[52812]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:440
				// _ = "end of CoverTab[52812]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:441
			// _ = "end of CoverTab[52809]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:442
			_go_fuzz_dep_.CoverTab[52810]++
																m := protowire.ConsumeFieldValue(num, typ, b)
																b = b[m:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:444
			// _ = "end of CoverTab[52810]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:445
		// _ = "end of CoverTab[52808]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:446
	// _ = "end of CoverTab[52807]"
}

var nameBuilderPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:450
		_go_fuzz_dep_.CoverTab[52813]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:450
		return new(strs.Builder)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:450
		// _ = "end of CoverTab[52813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:450
	},
}

func getBuilder() *strs.Builder {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:453
	_go_fuzz_dep_.CoverTab[52814]++
														return nameBuilderPool.Get().(*strs.Builder)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:454
	// _ = "end of CoverTab[52814]"
}
func putBuilder(b *strs.Builder) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:456
	_go_fuzz_dep_.CoverTab[52815]++
														nameBuilderPool.Put(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:457
	// _ = "end of CoverTab[52815]"
}

// makeFullName converts b to a protoreflect.FullName,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:460
// where b must start with a leading dot.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:462
func makeFullName(sb *strs.Builder, b []byte) protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:462
	_go_fuzz_dep_.CoverTab[52816]++
														if len(b) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:463
		_go_fuzz_dep_.CoverTab[52818]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:463
		return b[0] != '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:463
		// _ = "end of CoverTab[52818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:463
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:463
		_go_fuzz_dep_.CoverTab[52819]++
															panic("name reference must be fully qualified")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:464
		// _ = "end of CoverTab[52819]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:465
		_go_fuzz_dep_.CoverTab[52820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:465
		// _ = "end of CoverTab[52820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:465
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:465
	// _ = "end of CoverTab[52816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:465
	_go_fuzz_dep_.CoverTab[52817]++
														return protoreflect.FullName(sb.MakeString(b[1:]))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:466
	// _ = "end of CoverTab[52817]"
}

func appendFullName(sb *strs.Builder, prefix protoreflect.FullName, suffix []byte) protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:469
	_go_fuzz_dep_.CoverTab[52821]++
														return sb.AppendFullName(prefix, protoreflect.Name(strs.UnsafeString(suffix)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:470
	// _ = "end of CoverTab[52821]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:471
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/filedesc/desc_init.go:471
var _ = _go_fuzz_dep_.CoverTab
