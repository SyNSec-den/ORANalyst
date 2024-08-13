// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:5
)

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// Reset clears every field in the message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:13
// The resulting message shares no observable memory with its previous state
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:13
// other than the memory for the message itself.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:16
func Reset(m Message) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:16
	_go_fuzz_dep_.CoverTab[51527]++
											if mr, ok := m.(interface{ Reset() }); ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:17
		_go_fuzz_dep_.CoverTab[51529]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:17
		return hasProtoMethods
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:17
		// _ = "end of CoverTab[51529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:17
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:17
		_go_fuzz_dep_.CoverTab[51530]++
												mr.Reset()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:19
		// _ = "end of CoverTab[51530]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:20
		_go_fuzz_dep_.CoverTab[51531]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:20
		// _ = "end of CoverTab[51531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:20
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:20
	// _ = "end of CoverTab[51527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:20
	_go_fuzz_dep_.CoverTab[51528]++
											resetMessage(m.ProtoReflect())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:21
	// _ = "end of CoverTab[51528]"
}

func resetMessage(m protoreflect.Message) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:24
	_go_fuzz_dep_.CoverTab[51532]++
											if !m.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:25
		_go_fuzz_dep_.CoverTab[51536]++
												panic(fmt.Sprintf("cannot reset invalid %v message", m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:26
		// _ = "end of CoverTab[51536]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:27
		_go_fuzz_dep_.CoverTab[51537]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:27
		// _ = "end of CoverTab[51537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:27
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:27
	// _ = "end of CoverTab[51532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:27
	_go_fuzz_dep_.CoverTab[51533]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:30
	fds := m.Descriptor().Fields()
	for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:31
		_go_fuzz_dep_.CoverTab[51538]++
												m.Clear(fds.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:32
		// _ = "end of CoverTab[51538]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:33
	// _ = "end of CoverTab[51533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:33
	_go_fuzz_dep_.CoverTab[51534]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:36
	m.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:36
		_go_fuzz_dep_.CoverTab[51539]++
												m.Clear(fd)
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:38
		// _ = "end of CoverTab[51539]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:39
	// _ = "end of CoverTab[51534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:39
	_go_fuzz_dep_.CoverTab[51535]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:42
	m.SetUnknown(nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:42
	// _ = "end of CoverTab[51535]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/reset.go:43
var _ = _go_fuzz_dep_.CoverTab
