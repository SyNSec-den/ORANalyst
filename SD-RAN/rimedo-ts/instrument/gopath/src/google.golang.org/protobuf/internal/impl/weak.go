// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:5
)

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// weakFields adds methods to the exported WeakFields type for internal use.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:14
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:14
// The exported type is an alias to an unnamed type, so methods can't be
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:14
// defined directly on it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:18
type weakFields WeakFields

func (w weakFields) get(num protoreflect.FieldNumber) (protoreflect.ProtoMessage, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:20
	_go_fuzz_dep_.CoverTab[58994]++
												m, ok := w[int32(num)]
												return m, ok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:22
	// _ = "end of CoverTab[58994]"
}

func (w *weakFields) set(num protoreflect.FieldNumber, m protoreflect.ProtoMessage) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:25
	_go_fuzz_dep_.CoverTab[58995]++
												if *w == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:26
		_go_fuzz_dep_.CoverTab[58997]++
													*w = make(weakFields)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:27
		// _ = "end of CoverTab[58997]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:28
		_go_fuzz_dep_.CoverTab[58998]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:28
		// _ = "end of CoverTab[58998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:28
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:28
	// _ = "end of CoverTab[58995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:28
	_go_fuzz_dep_.CoverTab[58996]++
												(*w)[int32(num)] = m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:29
	// _ = "end of CoverTab[58996]"
}

func (w *weakFields) clear(num protoreflect.FieldNumber) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:32
	_go_fuzz_dep_.CoverTab[58999]++
												delete(*w, int32(num))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:33
	// _ = "end of CoverTab[58999]"
}

func (Export) HasWeak(w WeakFields, num protoreflect.FieldNumber) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:36
	_go_fuzz_dep_.CoverTab[59000]++
												_, ok := w[int32(num)]
												return ok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:38
	// _ = "end of CoverTab[59000]"
}

func (Export) ClearWeak(w *WeakFields, num protoreflect.FieldNumber) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:41
	_go_fuzz_dep_.CoverTab[59001]++
												delete(*w, int32(num))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:42
	// _ = "end of CoverTab[59001]"
}

func (Export) GetWeak(w WeakFields, num protoreflect.FieldNumber, name protoreflect.FullName) protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:45
	_go_fuzz_dep_.CoverTab[59002]++
												if m, ok := w[int32(num)]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:46
		_go_fuzz_dep_.CoverTab[59005]++
													return m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:47
		// _ = "end of CoverTab[59005]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:48
		_go_fuzz_dep_.CoverTab[59006]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:48
		// _ = "end of CoverTab[59006]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:48
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:48
	// _ = "end of CoverTab[59002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:48
	_go_fuzz_dep_.CoverTab[59003]++
												mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
												if mt == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:50
		_go_fuzz_dep_.CoverTab[59007]++
													panic(fmt.Sprintf("message %v for weak field is not linked in", name))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:51
		// _ = "end of CoverTab[59007]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:52
		_go_fuzz_dep_.CoverTab[59008]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:52
		// _ = "end of CoverTab[59008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:52
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:52
	// _ = "end of CoverTab[59003]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:52
	_go_fuzz_dep_.CoverTab[59004]++
												return mt.Zero().Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:53
	// _ = "end of CoverTab[59004]"
}

func (Export) SetWeak(w *WeakFields, num protoreflect.FieldNumber, name protoreflect.FullName, m protoreflect.ProtoMessage) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:56
	_go_fuzz_dep_.CoverTab[59009]++
												if m != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:57
		_go_fuzz_dep_.CoverTab[59013]++
													mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
													if mt == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:59
			_go_fuzz_dep_.CoverTab[59015]++
														panic(fmt.Sprintf("message %v for weak field is not linked in", name))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:60
			// _ = "end of CoverTab[59015]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:61
			_go_fuzz_dep_.CoverTab[59016]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:61
			// _ = "end of CoverTab[59016]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:61
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:61
		// _ = "end of CoverTab[59013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:61
		_go_fuzz_dep_.CoverTab[59014]++
													if mt != m.ProtoReflect().Type() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:62
			_go_fuzz_dep_.CoverTab[59017]++
														panic(fmt.Sprintf("invalid message type for weak field: got %T, want %T", m, mt.Zero().Interface()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:63
			// _ = "end of CoverTab[59017]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:64
			_go_fuzz_dep_.CoverTab[59018]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:64
			// _ = "end of CoverTab[59018]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:64
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:64
		// _ = "end of CoverTab[59014]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:65
		_go_fuzz_dep_.CoverTab[59019]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:65
		// _ = "end of CoverTab[59019]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:65
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:65
	// _ = "end of CoverTab[59009]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:65
	_go_fuzz_dep_.CoverTab[59010]++
												if m == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:66
		_go_fuzz_dep_.CoverTab[59020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:66
		return !m.ProtoReflect().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:66
		// _ = "end of CoverTab[59020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:66
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:66
		_go_fuzz_dep_.CoverTab[59021]++
													delete(*w, int32(num))
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:68
		// _ = "end of CoverTab[59021]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:69
		_go_fuzz_dep_.CoverTab[59022]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:69
		// _ = "end of CoverTab[59022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:69
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:69
	// _ = "end of CoverTab[59010]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:69
	_go_fuzz_dep_.CoverTab[59011]++
												if *w == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:70
		_go_fuzz_dep_.CoverTab[59023]++
													*w = make(weakFields)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:71
		// _ = "end of CoverTab[59023]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:72
		_go_fuzz_dep_.CoverTab[59024]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:72
		// _ = "end of CoverTab[59024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:72
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:72
	// _ = "end of CoverTab[59011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:72
	_go_fuzz_dep_.CoverTab[59012]++
												(*w)[int32(num)] = m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:73
	// _ = "end of CoverTab[59012]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/weak.go:74
var _ = _go_fuzz_dep_.CoverTab
