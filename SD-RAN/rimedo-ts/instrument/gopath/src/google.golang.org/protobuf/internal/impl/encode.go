// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:5
)

import (
	"math"
	"sort"
	"sync/atomic"

	"google.golang.org/protobuf/internal/flags"
	proto "google.golang.org/protobuf/proto"
	piface "google.golang.org/protobuf/runtime/protoiface"
)

type marshalOptions struct {
	flags piface.MarshalInputFlags
}

func (o marshalOptions) Options() proto.MarshalOptions {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:21
	_go_fuzz_dep_.CoverTab[57315]++
													return proto.MarshalOptions{
		AllowPartial:	true,
		Deterministic:	o.Deterministic(),
		UseCachedSize:	o.UseCachedSize(),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:26
	// _ = "end of CoverTab[57315]"
}

func (o marshalOptions) Deterministic() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:29
	_go_fuzz_dep_.CoverTab[57316]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:29
	return o.flags&piface.MarshalDeterministic != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:29
	// _ = "end of CoverTab[57316]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:29
}
func (o marshalOptions) UseCachedSize() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:30
	_go_fuzz_dep_.CoverTab[57317]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:30
	return o.flags&piface.MarshalUseCachedSize != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:30
	// _ = "end of CoverTab[57317]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:30
}

// size is protoreflect.Methods.Size.
func (mi *MessageInfo) size(in piface.SizeInput) piface.SizeOutput {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:33
	_go_fuzz_dep_.CoverTab[57318]++
													var p pointer
													if ms, ok := in.Message.(*messageState); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:35
		_go_fuzz_dep_.CoverTab[57320]++
														p = ms.pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:36
		// _ = "end of CoverTab[57320]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:37
		_go_fuzz_dep_.CoverTab[57321]++
														p = in.Message.(*messageReflectWrapper).pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:38
		// _ = "end of CoverTab[57321]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:39
	// _ = "end of CoverTab[57318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:39
	_go_fuzz_dep_.CoverTab[57319]++
													size := mi.sizePointer(p, marshalOptions{
		flags: in.Flags,
	})
													return piface.SizeOutput{Size: size}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:43
	// _ = "end of CoverTab[57319]"
}

func (mi *MessageInfo) sizePointer(p pointer, opts marshalOptions) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:46
	_go_fuzz_dep_.CoverTab[57322]++
													mi.init()
													if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:48
		_go_fuzz_dep_.CoverTab[57325]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:49
		// _ = "end of CoverTab[57325]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:50
		_go_fuzz_dep_.CoverTab[57326]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:50
		// _ = "end of CoverTab[57326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:50
	// _ = "end of CoverTab[57322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:50
	_go_fuzz_dep_.CoverTab[57323]++
													if opts.UseCachedSize() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:51
		_go_fuzz_dep_.CoverTab[57327]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:51
		return mi.sizecacheOffset.IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:51
		// _ = "end of CoverTab[57327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:51
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:51
		_go_fuzz_dep_.CoverTab[57328]++
														if size := atomic.LoadInt32(p.Apply(mi.sizecacheOffset).Int32()); size >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:52
			_go_fuzz_dep_.CoverTab[57329]++
															return int(size)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:53
			// _ = "end of CoverTab[57329]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:54
			_go_fuzz_dep_.CoverTab[57330]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:54
			// _ = "end of CoverTab[57330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:54
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:54
		// _ = "end of CoverTab[57328]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:55
		_go_fuzz_dep_.CoverTab[57331]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:55
		// _ = "end of CoverTab[57331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:55
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:55
	// _ = "end of CoverTab[57323]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:55
	_go_fuzz_dep_.CoverTab[57324]++
													return mi.sizePointerSlow(p, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:56
	// _ = "end of CoverTab[57324]"
}

func (mi *MessageInfo) sizePointerSlow(p pointer, opts marshalOptions) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:59
	_go_fuzz_dep_.CoverTab[57332]++
													if flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:60
		_go_fuzz_dep_.CoverTab[57338]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:60
		return mi.isMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:60
		// _ = "end of CoverTab[57338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:60
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:60
		_go_fuzz_dep_.CoverTab[57339]++
														size = sizeMessageSet(mi, p, opts)
														if mi.sizecacheOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:62
			_go_fuzz_dep_.CoverTab[57341]++
															atomic.StoreInt32(p.Apply(mi.sizecacheOffset).Int32(), int32(size))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:63
			// _ = "end of CoverTab[57341]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:64
			_go_fuzz_dep_.CoverTab[57342]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:64
			// _ = "end of CoverTab[57342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:64
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:64
		// _ = "end of CoverTab[57339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:64
		_go_fuzz_dep_.CoverTab[57340]++
														return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:65
		// _ = "end of CoverTab[57340]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:66
		_go_fuzz_dep_.CoverTab[57343]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:66
		// _ = "end of CoverTab[57343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:66
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:66
	// _ = "end of CoverTab[57332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:66
	_go_fuzz_dep_.CoverTab[57333]++
													if mi.extensionOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:67
		_go_fuzz_dep_.CoverTab[57344]++
														e := p.Apply(mi.extensionOffset).Extensions()
														size += mi.sizeExtensions(e, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:69
		// _ = "end of CoverTab[57344]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:70
		_go_fuzz_dep_.CoverTab[57345]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:70
		// _ = "end of CoverTab[57345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:70
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:70
	// _ = "end of CoverTab[57333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:70
	_go_fuzz_dep_.CoverTab[57334]++
													for _, f := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:71
		_go_fuzz_dep_.CoverTab[57346]++
														if f.funcs.size == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:72
			_go_fuzz_dep_.CoverTab[57349]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:73
			// _ = "end of CoverTab[57349]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:74
			_go_fuzz_dep_.CoverTab[57350]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:74
			// _ = "end of CoverTab[57350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:74
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:74
		// _ = "end of CoverTab[57346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:74
		_go_fuzz_dep_.CoverTab[57347]++
														fptr := p.Apply(f.offset)
														if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:76
			_go_fuzz_dep_.CoverTab[57351]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:76
			return fptr.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:76
			// _ = "end of CoverTab[57351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:76
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:76
			_go_fuzz_dep_.CoverTab[57352]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:77
			// _ = "end of CoverTab[57352]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:78
			_go_fuzz_dep_.CoverTab[57353]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:78
			// _ = "end of CoverTab[57353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:78
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:78
		// _ = "end of CoverTab[57347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:78
		_go_fuzz_dep_.CoverTab[57348]++
														size += f.funcs.size(fptr, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:79
		// _ = "end of CoverTab[57348]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:80
	// _ = "end of CoverTab[57334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:80
	_go_fuzz_dep_.CoverTab[57335]++
													if mi.unknownOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:81
		_go_fuzz_dep_.CoverTab[57354]++
														if u := mi.getUnknownBytes(p); u != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:82
			_go_fuzz_dep_.CoverTab[57355]++
															size += len(*u)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:83
			// _ = "end of CoverTab[57355]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:84
			_go_fuzz_dep_.CoverTab[57356]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:84
			// _ = "end of CoverTab[57356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:84
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:84
		// _ = "end of CoverTab[57354]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:85
		_go_fuzz_dep_.CoverTab[57357]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:85
		// _ = "end of CoverTab[57357]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:85
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:85
	// _ = "end of CoverTab[57335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:85
	_go_fuzz_dep_.CoverTab[57336]++
													if mi.sizecacheOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:86
		_go_fuzz_dep_.CoverTab[57358]++
														if size > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:87
			_go_fuzz_dep_.CoverTab[57359]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:91
			atomic.StoreInt32(p.Apply(mi.sizecacheOffset).Int32(), -1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:91
			// _ = "end of CoverTab[57359]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:92
			_go_fuzz_dep_.CoverTab[57360]++
															atomic.StoreInt32(p.Apply(mi.sizecacheOffset).Int32(), int32(size))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:93
			// _ = "end of CoverTab[57360]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:94
		// _ = "end of CoverTab[57358]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:95
		_go_fuzz_dep_.CoverTab[57361]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:95
		// _ = "end of CoverTab[57361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:95
	// _ = "end of CoverTab[57336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:95
	_go_fuzz_dep_.CoverTab[57337]++
													return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:96
	// _ = "end of CoverTab[57337]"
}

// marshal is protoreflect.Methods.Marshal.
func (mi *MessageInfo) marshal(in piface.MarshalInput) (out piface.MarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:100
	_go_fuzz_dep_.CoverTab[57362]++
													var p pointer
													if ms, ok := in.Message.(*messageState); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:102
		_go_fuzz_dep_.CoverTab[57364]++
														p = ms.pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:103
		// _ = "end of CoverTab[57364]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:104
		_go_fuzz_dep_.CoverTab[57365]++
														p = in.Message.(*messageReflectWrapper).pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:105
		// _ = "end of CoverTab[57365]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:106
	// _ = "end of CoverTab[57362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:106
	_go_fuzz_dep_.CoverTab[57363]++
													b, err := mi.marshalAppendPointer(in.Buf, p, marshalOptions{
		flags: in.Flags,
	})
													return piface.MarshalOutput{Buf: b}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:110
	// _ = "end of CoverTab[57363]"
}

func (mi *MessageInfo) marshalAppendPointer(b []byte, p pointer, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:113
	_go_fuzz_dep_.CoverTab[57366]++
													mi.init()
													if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:115
		_go_fuzz_dep_.CoverTab[57372]++
														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:116
		// _ = "end of CoverTab[57372]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:117
		_go_fuzz_dep_.CoverTab[57373]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:117
		// _ = "end of CoverTab[57373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:117
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:117
	// _ = "end of CoverTab[57366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:117
	_go_fuzz_dep_.CoverTab[57367]++
													if flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:118
		_go_fuzz_dep_.CoverTab[57374]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:118
		return mi.isMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:118
		// _ = "end of CoverTab[57374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:118
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:118
		_go_fuzz_dep_.CoverTab[57375]++
														return marshalMessageSet(mi, b, p, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:119
		// _ = "end of CoverTab[57375]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:120
		_go_fuzz_dep_.CoverTab[57376]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:120
		// _ = "end of CoverTab[57376]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:120
	// _ = "end of CoverTab[57367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:120
	_go_fuzz_dep_.CoverTab[57368]++
													var err error

													if mi.extensionOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:123
		_go_fuzz_dep_.CoverTab[57377]++
														e := p.Apply(mi.extensionOffset).Extensions()

														b, err = mi.appendExtensions(b, e, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:127
			_go_fuzz_dep_.CoverTab[57378]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:128
			// _ = "end of CoverTab[57378]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:129
			_go_fuzz_dep_.CoverTab[57379]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:129
			// _ = "end of CoverTab[57379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:129
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:129
		// _ = "end of CoverTab[57377]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:130
		_go_fuzz_dep_.CoverTab[57380]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:130
		// _ = "end of CoverTab[57380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:130
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:130
	// _ = "end of CoverTab[57368]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:130
	_go_fuzz_dep_.CoverTab[57369]++
													for _, f := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:131
		_go_fuzz_dep_.CoverTab[57381]++
														if f.funcs.marshal == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:132
			_go_fuzz_dep_.CoverTab[57384]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:133
			// _ = "end of CoverTab[57384]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:134
			_go_fuzz_dep_.CoverTab[57385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:134
			// _ = "end of CoverTab[57385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:134
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:134
		// _ = "end of CoverTab[57381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:134
		_go_fuzz_dep_.CoverTab[57382]++
														fptr := p.Apply(f.offset)
														if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:136
			_go_fuzz_dep_.CoverTab[57386]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:136
			return fptr.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:136
			// _ = "end of CoverTab[57386]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:136
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:136
			_go_fuzz_dep_.CoverTab[57387]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:137
			// _ = "end of CoverTab[57387]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:138
			_go_fuzz_dep_.CoverTab[57388]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:138
			// _ = "end of CoverTab[57388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:138
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:138
		// _ = "end of CoverTab[57382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:138
		_go_fuzz_dep_.CoverTab[57383]++
														b, err = f.funcs.marshal(b, fptr, f, opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:140
			_go_fuzz_dep_.CoverTab[57389]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:141
			// _ = "end of CoverTab[57389]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:142
			_go_fuzz_dep_.CoverTab[57390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:142
			// _ = "end of CoverTab[57390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:142
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:142
		// _ = "end of CoverTab[57383]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:143
	// _ = "end of CoverTab[57369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:143
	_go_fuzz_dep_.CoverTab[57370]++
													if mi.unknownOffset.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:144
		_go_fuzz_dep_.CoverTab[57391]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:144
		return !mi.isMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:144
		// _ = "end of CoverTab[57391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:144
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:144
		_go_fuzz_dep_.CoverTab[57392]++
														if u := mi.getUnknownBytes(p); u != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:145
			_go_fuzz_dep_.CoverTab[57393]++
															b = append(b, (*u)...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:146
			// _ = "end of CoverTab[57393]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:147
			_go_fuzz_dep_.CoverTab[57394]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:147
			// _ = "end of CoverTab[57394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:147
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:147
		// _ = "end of CoverTab[57392]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:148
		_go_fuzz_dep_.CoverTab[57395]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:148
		// _ = "end of CoverTab[57395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:148
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:148
	// _ = "end of CoverTab[57370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:148
	_go_fuzz_dep_.CoverTab[57371]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:149
	// _ = "end of CoverTab[57371]"
}

func (mi *MessageInfo) sizeExtensions(ext *map[int32]ExtensionField, opts marshalOptions) (n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:152
	_go_fuzz_dep_.CoverTab[57396]++
													if ext == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:153
		_go_fuzz_dep_.CoverTab[57399]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:154
		// _ = "end of CoverTab[57399]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:155
		_go_fuzz_dep_.CoverTab[57400]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:155
		// _ = "end of CoverTab[57400]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:155
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:155
	// _ = "end of CoverTab[57396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:155
	_go_fuzz_dep_.CoverTab[57397]++
													for _, x := range *ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:156
		_go_fuzz_dep_.CoverTab[57401]++
														xi := getExtensionFieldInfo(x.Type())
														if xi.funcs.size == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:158
			_go_fuzz_dep_.CoverTab[57403]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:159
			// _ = "end of CoverTab[57403]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:160
			_go_fuzz_dep_.CoverTab[57404]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:160
			// _ = "end of CoverTab[57404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:160
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:160
		// _ = "end of CoverTab[57401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:160
		_go_fuzz_dep_.CoverTab[57402]++
														n += xi.funcs.size(x.Value(), xi.tagsize, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:161
		// _ = "end of CoverTab[57402]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:162
	// _ = "end of CoverTab[57397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:162
	_go_fuzz_dep_.CoverTab[57398]++
													return n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:163
	// _ = "end of CoverTab[57398]"
}

func (mi *MessageInfo) appendExtensions(b []byte, ext *map[int32]ExtensionField, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:166
	_go_fuzz_dep_.CoverTab[57405]++
													if ext == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:167
		_go_fuzz_dep_.CoverTab[57407]++
														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:168
		// _ = "end of CoverTab[57407]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:169
		_go_fuzz_dep_.CoverTab[57408]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:169
		// _ = "end of CoverTab[57408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:169
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:169
	// _ = "end of CoverTab[57405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:169
	_go_fuzz_dep_.CoverTab[57406]++

													switch len(*ext) {
	case 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:172
		_go_fuzz_dep_.CoverTab[57409]++
														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:173
		// _ = "end of CoverTab[57409]"
	case 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:174
		_go_fuzz_dep_.CoverTab[57410]++
		// Fast-path for one extension: Don't bother sorting the keys.
		var err error
		for _, x := range *ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:177
			_go_fuzz_dep_.CoverTab[57415]++
															xi := getExtensionFieldInfo(x.Type())
															b, err = xi.funcs.marshal(b, x.Value(), xi.wiretag, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:179
			// _ = "end of CoverTab[57415]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:180
		// _ = "end of CoverTab[57410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:180
		_go_fuzz_dep_.CoverTab[57411]++
														return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:181
		// _ = "end of CoverTab[57411]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:182
		_go_fuzz_dep_.CoverTab[57412]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:185
		keys := make([]int, 0, len(*ext))
		for k := range *ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:186
			_go_fuzz_dep_.CoverTab[57416]++
															keys = append(keys, int(k))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:187
			// _ = "end of CoverTab[57416]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:188
		// _ = "end of CoverTab[57412]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:188
		_go_fuzz_dep_.CoverTab[57413]++
														sort.Ints(keys)
														var err error
														for _, k := range keys {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:191
			_go_fuzz_dep_.CoverTab[57417]++
															x := (*ext)[int32(k)]
															xi := getExtensionFieldInfo(x.Type())
															b, err = xi.funcs.marshal(b, x.Value(), xi.wiretag, opts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:195
				_go_fuzz_dep_.CoverTab[57418]++
																return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:196
				// _ = "end of CoverTab[57418]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:197
				_go_fuzz_dep_.CoverTab[57419]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:197
				// _ = "end of CoverTab[57419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:197
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:197
			// _ = "end of CoverTab[57417]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:198
		// _ = "end of CoverTab[57413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:198
		_go_fuzz_dep_.CoverTab[57414]++
														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:199
		// _ = "end of CoverTab[57414]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:200
	// _ = "end of CoverTab[57406]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:201
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/encode.go:201
var _ = _go_fuzz_dep_.CoverTab
