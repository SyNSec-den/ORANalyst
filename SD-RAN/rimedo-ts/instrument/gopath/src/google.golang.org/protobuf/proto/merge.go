// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:5
)

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Merge merges src into dst, which must be a message with the same descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// Populated scalar fields in src are copied to dst, while populated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// singular messages in src are merged into dst by recursively calling Merge.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// The elements of every list field in src is appended to the corresponded
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// list fields in dst. The entries of every map field in src is copied into
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// the corresponding map field in dst, possibly replacing existing entries.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// The unknown fields of src are appended to the unknown fields of dst.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// It is semantically equivalent to unmarshaling the encoded form of src
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:14
// into dst with the UnmarshalOptions.Merge option specified.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:25
func Merge(dst, src Message) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:25
	_go_fuzz_dep_.CoverTab[51436]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:29
	dstMsg, srcMsg := dst.ProtoReflect(), src.ProtoReflect()
	if dstMsg.Descriptor() != srcMsg.Descriptor() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:30
		_go_fuzz_dep_.CoverTab[51438]++
												if got, want := dstMsg.Descriptor().FullName(), srcMsg.Descriptor().FullName(); got != want {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:31
			_go_fuzz_dep_.CoverTab[51440]++
													panic(fmt.Sprintf("descriptor mismatch: %v != %v", got, want))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:32
			// _ = "end of CoverTab[51440]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:33
			_go_fuzz_dep_.CoverTab[51441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:33
			// _ = "end of CoverTab[51441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:33
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:33
		// _ = "end of CoverTab[51438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:33
		_go_fuzz_dep_.CoverTab[51439]++
												panic("descriptor mismatch")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:34
		// _ = "end of CoverTab[51439]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:35
		_go_fuzz_dep_.CoverTab[51442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:35
		// _ = "end of CoverTab[51442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:35
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:35
	// _ = "end of CoverTab[51436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:35
	_go_fuzz_dep_.CoverTab[51437]++
											mergeOptions{}.mergeMessage(dstMsg, srcMsg)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:36
	// _ = "end of CoverTab[51437]"
}

// Clone returns a deep copy of m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:39
// If the top-level message is invalid, it returns an invalid message as well.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:41
func Clone(m Message) Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:41
	_go_fuzz_dep_.CoverTab[51443]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:50
	if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:50
		_go_fuzz_dep_.CoverTab[51446]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:51
		// _ = "end of CoverTab[51446]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:52
		_go_fuzz_dep_.CoverTab[51447]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:52
		// _ = "end of CoverTab[51447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:52
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:52
	// _ = "end of CoverTab[51443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:52
	_go_fuzz_dep_.CoverTab[51444]++
											src := m.ProtoReflect()
											if !src.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:54
		_go_fuzz_dep_.CoverTab[51448]++
												return src.Type().Zero().Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:55
		// _ = "end of CoverTab[51448]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:56
		_go_fuzz_dep_.CoverTab[51449]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:56
		// _ = "end of CoverTab[51449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:56
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:56
	// _ = "end of CoverTab[51444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:56
	_go_fuzz_dep_.CoverTab[51445]++
											dst := src.New()
											mergeOptions{}.mergeMessage(dst, src)
											return dst.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:59
	// _ = "end of CoverTab[51445]"
}

// mergeOptions provides a namespace for merge functions, and can be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:62
// exported in the future if we add user-visible merge options.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:64
type mergeOptions struct{}

func (o mergeOptions) mergeMessage(dst, src protoreflect.Message) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:66
	_go_fuzz_dep_.CoverTab[51450]++
											methods := protoMethods(dst)
											if methods != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:68
		_go_fuzz_dep_.CoverTab[51454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:68
		return methods.Merge != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:68
		// _ = "end of CoverTab[51454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:68
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:68
		_go_fuzz_dep_.CoverTab[51455]++
												in := protoiface.MergeInput{
			Destination:	dst,
			Source:		src,
		}
		out := methods.Merge(in)
		if out.Flags&protoiface.MergeComplete != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:74
			_go_fuzz_dep_.CoverTab[51456]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:75
			// _ = "end of CoverTab[51456]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:76
			_go_fuzz_dep_.CoverTab[51457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:76
			// _ = "end of CoverTab[51457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:76
		// _ = "end of CoverTab[51455]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:77
		_go_fuzz_dep_.CoverTab[51458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:77
		// _ = "end of CoverTab[51458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:77
	// _ = "end of CoverTab[51450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:77
	_go_fuzz_dep_.CoverTab[51451]++

											if !dst.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:79
		_go_fuzz_dep_.CoverTab[51459]++
												panic(fmt.Sprintf("cannot merge into invalid %v message", dst.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:80
		// _ = "end of CoverTab[51459]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:81
		_go_fuzz_dep_.CoverTab[51460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:81
		// _ = "end of CoverTab[51460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:81
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:81
	// _ = "end of CoverTab[51451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:81
	_go_fuzz_dep_.CoverTab[51452]++

											src.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:83
		_go_fuzz_dep_.CoverTab[51461]++
												switch {
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:85
			_go_fuzz_dep_.CoverTab[51463]++
													o.mergeList(dst.Mutable(fd).List(), v.List(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:86
			// _ = "end of CoverTab[51463]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:87
			_go_fuzz_dep_.CoverTab[51464]++
													o.mergeMap(dst.Mutable(fd).Map(), v.Map(), fd.MapValue())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:88
			// _ = "end of CoverTab[51464]"
		case fd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:89
			_go_fuzz_dep_.CoverTab[51465]++
													o.mergeMessage(dst.Mutable(fd).Message(), v.Message())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:90
			// _ = "end of CoverTab[51465]"
		case fd.Kind() == protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:91
			_go_fuzz_dep_.CoverTab[51466]++
													dst.Set(fd, o.cloneBytes(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:92
			// _ = "end of CoverTab[51466]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:93
			_go_fuzz_dep_.CoverTab[51467]++
													dst.Set(fd, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:94
			// _ = "end of CoverTab[51467]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:95
		// _ = "end of CoverTab[51461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:95
		_go_fuzz_dep_.CoverTab[51462]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:96
		// _ = "end of CoverTab[51462]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:97
	// _ = "end of CoverTab[51452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:97
	_go_fuzz_dep_.CoverTab[51453]++

											if len(src.GetUnknown()) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:99
			_go_fuzz_dep_.CoverTab[51468]++
													dst.SetUnknown(append(dst.GetUnknown(), src.GetUnknown()...))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:100
		// _ = "end of CoverTab[51468]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:101
		_go_fuzz_dep_.CoverTab[51469]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:101
		// _ = "end of CoverTab[51469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:101
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:101
	// _ = "end of CoverTab[51453]"
}

func (o mergeOptions) mergeList(dst, src protoreflect.List, fd protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:104
	_go_fuzz_dep_.CoverTab[51470]++

												for i, n := 0, src.Len(); i < n; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:106
		_go_fuzz_dep_.CoverTab[51471]++
													switch v := src.Get(i); {
		case fd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:108
			_go_fuzz_dep_.CoverTab[51472]++
														dstv := dst.NewElement()
														o.mergeMessage(dstv.Message(), v.Message())
														dst.Append(dstv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:111
			// _ = "end of CoverTab[51472]"
		case fd.Kind() == protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:112
			_go_fuzz_dep_.CoverTab[51473]++
														dst.Append(o.cloneBytes(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:113
			// _ = "end of CoverTab[51473]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:114
			_go_fuzz_dep_.CoverTab[51474]++
														dst.Append(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:115
			// _ = "end of CoverTab[51474]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:116
		// _ = "end of CoverTab[51471]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:117
	// _ = "end of CoverTab[51470]"
}

func (o mergeOptions) mergeMap(dst, src protoreflect.Map, fd protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:120
	_go_fuzz_dep_.CoverTab[51475]++

												src.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:122
		_go_fuzz_dep_.CoverTab[51476]++
													switch {
		case fd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:124
			_go_fuzz_dep_.CoverTab[51478]++
														dstv := dst.NewValue()
														o.mergeMessage(dstv.Message(), v.Message())
														dst.Set(k, dstv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:127
			// _ = "end of CoverTab[51478]"
		case fd.Kind() == protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:128
			_go_fuzz_dep_.CoverTab[51479]++
														dst.Set(k, o.cloneBytes(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:129
			// _ = "end of CoverTab[51479]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:130
			_go_fuzz_dep_.CoverTab[51480]++
														dst.Set(k, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:131
			// _ = "end of CoverTab[51480]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:132
		// _ = "end of CoverTab[51476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:132
		_go_fuzz_dep_.CoverTab[51477]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:133
		// _ = "end of CoverTab[51477]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:134
	// _ = "end of CoverTab[51475]"
}

func (o mergeOptions) cloneBytes(v protoreflect.Value) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:137
	_go_fuzz_dep_.CoverTab[51481]++
												return protoreflect.ValueOfBytes(append([]byte{}, v.Bytes()...))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:138
	// _ = "end of CoverTab[51481]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:139
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/merge.go:139
var _ = _go_fuzz_dep_.CoverTab
