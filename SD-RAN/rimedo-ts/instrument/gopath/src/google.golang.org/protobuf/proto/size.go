// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:5
)

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Size returns the size in bytes of the wire-format encoding of m.
func Size(m Message) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:15
	_go_fuzz_dep_.CoverTab[51540]++
											return MarshalOptions{}.Size(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:16
	// _ = "end of CoverTab[51540]"
}

// Size returns the size in bytes of the wire-format encoding of m.
func (o MarshalOptions) Size(m Message) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:20
	_go_fuzz_dep_.CoverTab[51541]++

											if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:22
		_go_fuzz_dep_.CoverTab[51543]++
												return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:23
		// _ = "end of CoverTab[51543]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:24
		_go_fuzz_dep_.CoverTab[51544]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:24
		// _ = "end of CoverTab[51544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:24
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:24
	// _ = "end of CoverTab[51541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:24
	_go_fuzz_dep_.CoverTab[51542]++

											return o.size(m.ProtoReflect())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:26
	// _ = "end of CoverTab[51542]"
}

// size is a centralized function that all size operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:29
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:29
// introducing other code paths for size that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:32
func (o MarshalOptions) size(m protoreflect.Message) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:32
	_go_fuzz_dep_.CoverTab[51545]++
											methods := protoMethods(m)
											if methods != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:34
		_go_fuzz_dep_.CoverTab[51548]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:34
		return methods.Size != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:34
		// _ = "end of CoverTab[51548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:34
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:34
		_go_fuzz_dep_.CoverTab[51549]++
												out := methods.Size(protoiface.SizeInput{
			Message: m,
		})
												return out.Size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:38
		// _ = "end of CoverTab[51549]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:39
		_go_fuzz_dep_.CoverTab[51550]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:39
		// _ = "end of CoverTab[51550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:39
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:39
	// _ = "end of CoverTab[51545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:39
	_go_fuzz_dep_.CoverTab[51546]++
											if methods != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:40
		_go_fuzz_dep_.CoverTab[51551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:40
		return methods.Marshal != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:40
		// _ = "end of CoverTab[51551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:40
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:40
		_go_fuzz_dep_.CoverTab[51552]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:43
		out, _ := methods.Marshal(protoiface.MarshalInput{
			Message: m,
		})
												return len(out.Buf)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:46
		// _ = "end of CoverTab[51552]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:47
		_go_fuzz_dep_.CoverTab[51553]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:47
		// _ = "end of CoverTab[51553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:47
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:47
	// _ = "end of CoverTab[51546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:47
	_go_fuzz_dep_.CoverTab[51547]++
											return o.sizeMessageSlow(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:48
	// _ = "end of CoverTab[51547]"
}

func (o MarshalOptions) sizeMessageSlow(m protoreflect.Message) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:51
	_go_fuzz_dep_.CoverTab[51554]++
											if messageset.IsMessageSet(m.Descriptor()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:52
		_go_fuzz_dep_.CoverTab[51557]++
												return o.sizeMessageSet(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:53
		// _ = "end of CoverTab[51557]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:54
		_go_fuzz_dep_.CoverTab[51558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:54
		// _ = "end of CoverTab[51558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:54
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:54
	// _ = "end of CoverTab[51554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:54
	_go_fuzz_dep_.CoverTab[51555]++
											m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:55
		_go_fuzz_dep_.CoverTab[51559]++
												size += o.sizeField(fd, v)
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:57
		// _ = "end of CoverTab[51559]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:58
	// _ = "end of CoverTab[51555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:58
	_go_fuzz_dep_.CoverTab[51556]++
											size += len(m.GetUnknown())
											return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:60
	// _ = "end of CoverTab[51556]"
}

func (o MarshalOptions) sizeField(fd protoreflect.FieldDescriptor, value protoreflect.Value) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:63
	_go_fuzz_dep_.CoverTab[51560]++
											num := fd.Number()
											switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:66
		_go_fuzz_dep_.CoverTab[51561]++
												return o.sizeList(num, fd, value.List())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:67
		// _ = "end of CoverTab[51561]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:68
		_go_fuzz_dep_.CoverTab[51562]++
												return o.sizeMap(num, fd, value.Map())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:69
		// _ = "end of CoverTab[51562]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:70
		_go_fuzz_dep_.CoverTab[51563]++
												return protowire.SizeTag(num) + o.sizeSingular(num, fd.Kind(), value)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:71
		// _ = "end of CoverTab[51563]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:72
	// _ = "end of CoverTab[51560]"
}

func (o MarshalOptions) sizeList(num protowire.Number, fd protoreflect.FieldDescriptor, list protoreflect.List) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:75
	_go_fuzz_dep_.CoverTab[51564]++
											if fd.IsPacked() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:76
		_go_fuzz_dep_.CoverTab[51567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:76
		return list.Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:76
		// _ = "end of CoverTab[51567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:76
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:76
		_go_fuzz_dep_.CoverTab[51568]++
												content := 0
												for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:78
			_go_fuzz_dep_.CoverTab[51570]++
													content += o.sizeSingular(num, fd.Kind(), list.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:79
			// _ = "end of CoverTab[51570]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:80
		// _ = "end of CoverTab[51568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:80
		_go_fuzz_dep_.CoverTab[51569]++
												return protowire.SizeTag(num) + protowire.SizeBytes(content)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:81
		// _ = "end of CoverTab[51569]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:82
		_go_fuzz_dep_.CoverTab[51571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:82
		// _ = "end of CoverTab[51571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:82
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:82
	// _ = "end of CoverTab[51564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:82
	_go_fuzz_dep_.CoverTab[51565]++

											for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:84
		_go_fuzz_dep_.CoverTab[51572]++
												size += protowire.SizeTag(num) + o.sizeSingular(num, fd.Kind(), list.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:85
		// _ = "end of CoverTab[51572]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:86
	// _ = "end of CoverTab[51565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:86
	_go_fuzz_dep_.CoverTab[51566]++
											return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:87
	// _ = "end of CoverTab[51566]"
}

func (o MarshalOptions) sizeMap(num protowire.Number, fd protoreflect.FieldDescriptor, mapv protoreflect.Map) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:90
	_go_fuzz_dep_.CoverTab[51573]++
											mapv.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:91
		_go_fuzz_dep_.CoverTab[51575]++
												size += protowire.SizeTag(num)
												size += protowire.SizeBytes(o.sizeField(fd.MapKey(), key.Value()) + o.sizeField(fd.MapValue(), value))
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:94
		// _ = "end of CoverTab[51575]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:95
	// _ = "end of CoverTab[51573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:95
	_go_fuzz_dep_.CoverTab[51574]++
											return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:96
	// _ = "end of CoverTab[51574]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:97
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/size.go:97
var _ = _go_fuzz_dep_.CoverTab
