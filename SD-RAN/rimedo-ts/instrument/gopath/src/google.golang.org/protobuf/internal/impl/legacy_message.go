// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:5
)

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"google.golang.org/protobuf/internal/descopts"
	ptag "google.golang.org/protobuf/internal/encoding/tag"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// legacyWrapMessage wraps v as a protoreflect.Message,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:22
// where v must be a *struct kind and not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:24
func legacyWrapMessage(v reflect.Value) protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:24
	_go_fuzz_dep_.CoverTab[57637]++
														t := v.Type()
														if t.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:26
		_go_fuzz_dep_.CoverTab[57639]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:26
		return t.Elem().Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:26
		// _ = "end of CoverTab[57639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:26
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:26
		_go_fuzz_dep_.CoverTab[57640]++
															return aberrantMessage{v: v}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:27
		// _ = "end of CoverTab[57640]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:28
		_go_fuzz_dep_.CoverTab[57641]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:28
		// _ = "end of CoverTab[57641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:28
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:28
	// _ = "end of CoverTab[57637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:28
	_go_fuzz_dep_.CoverTab[57638]++
														mt := legacyLoadMessageInfo(t, "")
														return mt.MessageOf(v.Interface())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:30
	// _ = "end of CoverTab[57638]"
}

// legacyLoadMessageType dynamically loads a protoreflect.Type for t,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:33
// where t must be not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:33
// The provided name is used if it cannot be determined from the message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:36
func legacyLoadMessageType(t reflect.Type, name protoreflect.FullName) protoreflect.MessageType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:36
	_go_fuzz_dep_.CoverTab[57642]++
														if t.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:37
		_go_fuzz_dep_.CoverTab[57644]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:37
		return t.Elem().Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:37
		// _ = "end of CoverTab[57644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:37
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:37
		_go_fuzz_dep_.CoverTab[57645]++
															return aberrantMessageType{t}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:38
		// _ = "end of CoverTab[57645]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:39
		_go_fuzz_dep_.CoverTab[57646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:39
		// _ = "end of CoverTab[57646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:39
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:39
	// _ = "end of CoverTab[57642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:39
	_go_fuzz_dep_.CoverTab[57643]++
														return legacyLoadMessageInfo(t, name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:40
	// _ = "end of CoverTab[57643]"
}

var legacyMessageTypeCache sync.Map	// map[reflect.Type]*MessageInfo

// legacyLoadMessageInfo dynamically loads a *MessageInfo for t,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:45
// where t must be a *struct kind and not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:45
// The provided name is used if it cannot be determined from the message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:48
func legacyLoadMessageInfo(t reflect.Type, name protoreflect.FullName) *MessageInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:48
	_go_fuzz_dep_.CoverTab[57647]++

														if mt, ok := legacyMessageTypeCache.Load(t); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:50
		_go_fuzz_dep_.CoverTab[57653]++
															return mt.(*MessageInfo)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:51
		// _ = "end of CoverTab[57653]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:52
		_go_fuzz_dep_.CoverTab[57654]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:52
		// _ = "end of CoverTab[57654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:52
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:52
	// _ = "end of CoverTab[57647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:52
	_go_fuzz_dep_.CoverTab[57648]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:55
	mi := &MessageInfo{
		Desc:		legacyLoadMessageDesc(t, name),
		GoReflectType:	t,
	}

	var hasMarshal, hasUnmarshal bool
	v := reflect.Zero(t).Interface()
	if _, hasMarshal = v.(legacyMarshaler); hasMarshal {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:62
		_go_fuzz_dep_.CoverTab[57655]++
															mi.methods.Marshal = legacyMarshal

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:69
		mi.methods.Flags |= protoiface.SupportMarshalDeterministic
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:69
		// _ = "end of CoverTab[57655]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:70
		_go_fuzz_dep_.CoverTab[57656]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:70
		// _ = "end of CoverTab[57656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:70
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:70
	// _ = "end of CoverTab[57648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:70
	_go_fuzz_dep_.CoverTab[57649]++
														if _, hasUnmarshal = v.(legacyUnmarshaler); hasUnmarshal {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:71
		_go_fuzz_dep_.CoverTab[57657]++
															mi.methods.Unmarshal = legacyUnmarshal
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:72
		// _ = "end of CoverTab[57657]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:73
		_go_fuzz_dep_.CoverTab[57658]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:73
		// _ = "end of CoverTab[57658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:73
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:73
	// _ = "end of CoverTab[57649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:73
	_go_fuzz_dep_.CoverTab[57650]++
														if _, hasMerge := v.(legacyMerger); hasMerge || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
		_go_fuzz_dep_.CoverTab[57659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
		return (hasMarshal && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
			_go_fuzz_dep_.CoverTab[57660]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
			return hasUnmarshal
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
			// _ = "end of CoverTab[57660]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
		// _ = "end of CoverTab[57659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:74
		_go_fuzz_dep_.CoverTab[57661]++
															mi.methods.Merge = legacyMerge
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:75
		// _ = "end of CoverTab[57661]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:76
		_go_fuzz_dep_.CoverTab[57662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:76
		// _ = "end of CoverTab[57662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:76
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:76
	// _ = "end of CoverTab[57650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:76
	_go_fuzz_dep_.CoverTab[57651]++

														if mi, ok := legacyMessageTypeCache.LoadOrStore(t, mi); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:78
		_go_fuzz_dep_.CoverTab[57663]++
															return mi.(*MessageInfo)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:79
		// _ = "end of CoverTab[57663]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:80
		_go_fuzz_dep_.CoverTab[57664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:80
		// _ = "end of CoverTab[57664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:80
	// _ = "end of CoverTab[57651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:80
	_go_fuzz_dep_.CoverTab[57652]++
														return mi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:81
	// _ = "end of CoverTab[57652]"
}

var legacyMessageDescCache sync.Map	// map[reflect.Type]protoreflect.MessageDescriptor

// LegacyLoadMessageDesc returns an MessageDescriptor derived from the Go type,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:86
// which should be a *struct kind and must not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:86
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:86
// This is exported for testing purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:90
func LegacyLoadMessageDesc(t reflect.Type) protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:90
	_go_fuzz_dep_.CoverTab[57665]++
														return legacyLoadMessageDesc(t, "")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:91
	// _ = "end of CoverTab[57665]"
}
func legacyLoadMessageDesc(t reflect.Type, name protoreflect.FullName) protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:93
	_go_fuzz_dep_.CoverTab[57666]++

														if mi, ok := legacyMessageDescCache.Load(t); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:95
		_go_fuzz_dep_.CoverTab[57676]++
															return mi.(protoreflect.MessageDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:96
		// _ = "end of CoverTab[57676]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:97
		_go_fuzz_dep_.CoverTab[57677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:97
		// _ = "end of CoverTab[57677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:97
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:97
	// _ = "end of CoverTab[57666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:97
	_go_fuzz_dep_.CoverTab[57667]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:100
	mv := reflect.Zero(t).Interface()
	if _, ok := mv.(protoreflect.ProtoMessage); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:101
		_go_fuzz_dep_.CoverTab[57678]++
															panic(fmt.Sprintf("%v already implements proto.Message", t))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:102
		// _ = "end of CoverTab[57678]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:103
		_go_fuzz_dep_.CoverTab[57679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:103
		// _ = "end of CoverTab[57679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:103
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:103
	// _ = "end of CoverTab[57667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:103
	_go_fuzz_dep_.CoverTab[57668]++
														mdV1, ok := mv.(messageV1)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:105
		_go_fuzz_dep_.CoverTab[57680]++
															return aberrantLoadMessageDesc(t, name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:106
		// _ = "end of CoverTab[57680]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:107
		_go_fuzz_dep_.CoverTab[57681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:107
		// _ = "end of CoverTab[57681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:107
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:107
	// _ = "end of CoverTab[57668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:107
	_go_fuzz_dep_.CoverTab[57669]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:113
	b, idxs := func() ([]byte, []int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:113
		_go_fuzz_dep_.CoverTab[57682]++
															defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:114
			_go_fuzz_dep_.CoverTab[57684]++
																recover()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:115
			// _ = "end of CoverTab[57684]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:116
		// _ = "end of CoverTab[57682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:116
		_go_fuzz_dep_.CoverTab[57683]++
															return mdV1.Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:117
		// _ = "end of CoverTab[57683]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:118
	// _ = "end of CoverTab[57669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:118
	_go_fuzz_dep_.CoverTab[57670]++
														if b == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:119
		_go_fuzz_dep_.CoverTab[57685]++
															return aberrantLoadMessageDesc(t, name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:120
		// _ = "end of CoverTab[57685]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:121
		_go_fuzz_dep_.CoverTab[57686]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:121
		// _ = "end of CoverTab[57686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:121
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:121
	// _ = "end of CoverTab[57670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:121
	_go_fuzz_dep_.CoverTab[57671]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:126
	if t.Elem().Kind() == reflect.Struct {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:126
		_go_fuzz_dep_.CoverTab[57687]++
															if nfield := t.Elem().NumField(); nfield > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:127
			_go_fuzz_dep_.CoverTab[57688]++
																hasProtoField := false
																for i := 0; i < nfield; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:129
				_go_fuzz_dep_.CoverTab[57690]++
																	f := t.Elem().Field(i)
																	if f.Tag.Get("protobuf") != "" || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					_go_fuzz_dep_.CoverTab[57691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					return f.Tag.Get("protobuf_oneof") != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					// _ = "end of CoverTab[57691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					_go_fuzz_dep_.CoverTab[57692]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					return strings.HasPrefix(f.Name, "XXX_")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					// _ = "end of CoverTab[57692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:131
					_go_fuzz_dep_.CoverTab[57693]++
																		hasProtoField = true
																		break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:133
					// _ = "end of CoverTab[57693]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:134
					_go_fuzz_dep_.CoverTab[57694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:134
					// _ = "end of CoverTab[57694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:134
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:134
				// _ = "end of CoverTab[57690]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:135
			// _ = "end of CoverTab[57688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:135
			_go_fuzz_dep_.CoverTab[57689]++
																if !hasProtoField {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:136
				_go_fuzz_dep_.CoverTab[57695]++
																	return aberrantLoadMessageDesc(t, name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:137
				// _ = "end of CoverTab[57695]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:138
				_go_fuzz_dep_.CoverTab[57696]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:138
				// _ = "end of CoverTab[57696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:138
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:138
			// _ = "end of CoverTab[57689]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:139
			_go_fuzz_dep_.CoverTab[57697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:139
			// _ = "end of CoverTab[57697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:139
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:139
		// _ = "end of CoverTab[57687]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:140
		_go_fuzz_dep_.CoverTab[57698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:140
		// _ = "end of CoverTab[57698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:140
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:140
	// _ = "end of CoverTab[57671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:140
	_go_fuzz_dep_.CoverTab[57672]++

														md := legacyLoadFileDesc(b).Messages().Get(idxs[0])
														for _, i := range idxs[1:] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:143
		_go_fuzz_dep_.CoverTab[57699]++
															md = md.Messages().Get(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:144
		// _ = "end of CoverTab[57699]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:145
	// _ = "end of CoverTab[57672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:145
	_go_fuzz_dep_.CoverTab[57673]++
														if name != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:146
		_go_fuzz_dep_.CoverTab[57700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:146
		return md.FullName() != name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:146
		// _ = "end of CoverTab[57700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:146
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:146
		_go_fuzz_dep_.CoverTab[57701]++
															panic(fmt.Sprintf("mismatching message name: got %v, want %v", md.FullName(), name))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:147
		// _ = "end of CoverTab[57701]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:148
		_go_fuzz_dep_.CoverTab[57702]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:148
		// _ = "end of CoverTab[57702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:148
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:148
	// _ = "end of CoverTab[57673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:148
	_go_fuzz_dep_.CoverTab[57674]++
														if md, ok := legacyMessageDescCache.LoadOrStore(t, md); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:149
		_go_fuzz_dep_.CoverTab[57703]++
															return md.(protoreflect.MessageDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:150
		// _ = "end of CoverTab[57703]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:151
		_go_fuzz_dep_.CoverTab[57704]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:151
		// _ = "end of CoverTab[57704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:151
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:151
	// _ = "end of CoverTab[57674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:151
	_go_fuzz_dep_.CoverTab[57675]++
														return md
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:152
	// _ = "end of CoverTab[57675]"
}

var (
	aberrantMessageDescLock		sync.Mutex
	aberrantMessageDescCache	map[reflect.Type]protoreflect.MessageDescriptor
)

// aberrantLoadMessageDesc returns an MessageDescriptor derived from the Go type,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:160
// which must not implement protoreflect.ProtoMessage or messageV1.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:160
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:160
// This is a best-effort derivation of the message descriptor using the protobuf
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:160
// tags on the struct fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:165
func aberrantLoadMessageDesc(t reflect.Type, name protoreflect.FullName) protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:165
	_go_fuzz_dep_.CoverTab[57705]++
														aberrantMessageDescLock.Lock()
														defer aberrantMessageDescLock.Unlock()
														if aberrantMessageDescCache == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:168
		_go_fuzz_dep_.CoverTab[57707]++
															aberrantMessageDescCache = make(map[reflect.Type]protoreflect.MessageDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:169
		// _ = "end of CoverTab[57707]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:170
		_go_fuzz_dep_.CoverTab[57708]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:170
		// _ = "end of CoverTab[57708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:170
	// _ = "end of CoverTab[57705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:170
	_go_fuzz_dep_.CoverTab[57706]++
														return aberrantLoadMessageDescReentrant(t, name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:171
	// _ = "end of CoverTab[57706]"
}
func aberrantLoadMessageDescReentrant(t reflect.Type, name protoreflect.FullName) protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:173
	_go_fuzz_dep_.CoverTab[57709]++

														if md, ok := aberrantMessageDescCache[t]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:175
		_go_fuzz_dep_.CoverTab[57716]++
															return md
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:176
		// _ = "end of CoverTab[57716]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:177
		_go_fuzz_dep_.CoverTab[57717]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:177
		// _ = "end of CoverTab[57717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:177
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:177
	// _ = "end of CoverTab[57709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:177
	_go_fuzz_dep_.CoverTab[57710]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:182
	md := &filedesc.Message{L2: new(filedesc.MessageL2)}
	md.L0.FullName = aberrantDeriveMessageName(t, name)
	md.L0.ParentFile = filedesc.SurrogateProto2
	aberrantMessageDescCache[t] = md

	if t.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:187
		_go_fuzz_dep_.CoverTab[57718]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:187
		return t.Elem().Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:187
		// _ = "end of CoverTab[57718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:187
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:187
		_go_fuzz_dep_.CoverTab[57719]++
															return md
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:188
		// _ = "end of CoverTab[57719]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:189
		_go_fuzz_dep_.CoverTab[57720]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:189
		// _ = "end of CoverTab[57720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:189
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:189
	// _ = "end of CoverTab[57710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:189
	_go_fuzz_dep_.CoverTab[57711]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:192
	for i := 0; i < t.Elem().NumField(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:192
		_go_fuzz_dep_.CoverTab[57721]++
															f := t.Elem().Field(i)
															if tag := f.Tag.Get("protobuf"); tag != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:194
			_go_fuzz_dep_.CoverTab[57722]++
																switch f.Type.Kind() {
			case reflect.Bool, reflect.Int32, reflect.Int64, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:196
				_go_fuzz_dep_.CoverTab[57724]++
																	md.L0.ParentFile = filedesc.SurrogateProto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:197
				// _ = "end of CoverTab[57724]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:197
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:197
				_go_fuzz_dep_.CoverTab[57725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:197
				// _ = "end of CoverTab[57725]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:198
			// _ = "end of CoverTab[57722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:198
			_go_fuzz_dep_.CoverTab[57723]++
																for _, s := range strings.Split(tag, ",") {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:199
				_go_fuzz_dep_.CoverTab[57726]++
																	if s == "proto3" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:200
					_go_fuzz_dep_.CoverTab[57727]++
																		md.L0.ParentFile = filedesc.SurrogateProto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:201
					// _ = "end of CoverTab[57727]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:202
					_go_fuzz_dep_.CoverTab[57728]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:202
					// _ = "end of CoverTab[57728]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:202
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:202
				// _ = "end of CoverTab[57726]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:203
			// _ = "end of CoverTab[57723]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:204
			_go_fuzz_dep_.CoverTab[57729]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:204
			// _ = "end of CoverTab[57729]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:204
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:204
		// _ = "end of CoverTab[57721]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:205
	// _ = "end of CoverTab[57711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:205
	_go_fuzz_dep_.CoverTab[57712]++

	// Obtain a list of oneof wrapper types.
	var oneofWrappers []reflect.Type
	for _, method := range []string{"XXX_OneofFuncs", "XXX_OneofWrappers"} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:209
		_go_fuzz_dep_.CoverTab[57730]++
															if fn, ok := t.MethodByName(method); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:210
			_go_fuzz_dep_.CoverTab[57731]++
																for _, v := range fn.Func.Call([]reflect.Value{reflect.Zero(fn.Type.In(0))}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:211
				_go_fuzz_dep_.CoverTab[57732]++
																	if vs, ok := v.Interface().([]interface{}); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:212
					_go_fuzz_dep_.CoverTab[57733]++
																		for _, v := range vs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:213
						_go_fuzz_dep_.CoverTab[57734]++
																			oneofWrappers = append(oneofWrappers, reflect.TypeOf(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:214
						// _ = "end of CoverTab[57734]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:215
					// _ = "end of CoverTab[57733]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:216
					_go_fuzz_dep_.CoverTab[57735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:216
					// _ = "end of CoverTab[57735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:216
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:216
				// _ = "end of CoverTab[57732]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:217
			// _ = "end of CoverTab[57731]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:218
			_go_fuzz_dep_.CoverTab[57736]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:218
			// _ = "end of CoverTab[57736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:218
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:218
		// _ = "end of CoverTab[57730]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:219
	// _ = "end of CoverTab[57712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:219
	_go_fuzz_dep_.CoverTab[57713]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:222
	if fn, ok := t.MethodByName("ExtensionRangeArray"); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:222
		_go_fuzz_dep_.CoverTab[57737]++
															vs := fn.Func.Call([]reflect.Value{reflect.Zero(fn.Type.In(0))})[0]
															for i := 0; i < vs.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:224
			_go_fuzz_dep_.CoverTab[57738]++
																v := vs.Index(i)
																md.L2.ExtensionRanges.List = append(md.L2.ExtensionRanges.List, [2]protoreflect.FieldNumber{
				protoreflect.FieldNumber(v.FieldByName("Start").Int()),
				protoreflect.FieldNumber(v.FieldByName("End").Int() + 1),
			})
																md.L2.ExtensionRangeOptions = append(md.L2.ExtensionRangeOptions, nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:230
			// _ = "end of CoverTab[57738]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:231
		// _ = "end of CoverTab[57737]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:232
		_go_fuzz_dep_.CoverTab[57739]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:232
		// _ = "end of CoverTab[57739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:232
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:232
	// _ = "end of CoverTab[57713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:232
	_go_fuzz_dep_.CoverTab[57714]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:235
	for i := 0; i < t.Elem().NumField(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:235
		_go_fuzz_dep_.CoverTab[57740]++
															f := t.Elem().Field(i)
															if tag := f.Tag.Get("protobuf"); tag != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:237
			_go_fuzz_dep_.CoverTab[57742]++
																tagKey := f.Tag.Get("protobuf_key")
																tagVal := f.Tag.Get("protobuf_val")
																aberrantAppendField(md, f.Type, tag, tagKey, tagVal)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:240
			// _ = "end of CoverTab[57742]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:241
			_go_fuzz_dep_.CoverTab[57743]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:241
			// _ = "end of CoverTab[57743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:241
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:241
		// _ = "end of CoverTab[57740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:241
		_go_fuzz_dep_.CoverTab[57741]++
															if tag := f.Tag.Get("protobuf_oneof"); tag != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:242
			_go_fuzz_dep_.CoverTab[57744]++
																n := len(md.L2.Oneofs.List)
																md.L2.Oneofs.List = append(md.L2.Oneofs.List, filedesc.Oneof{})
																od := &md.L2.Oneofs.List[n]
																od.L0.FullName = md.FullName().Append(protoreflect.Name(tag))
																od.L0.ParentFile = md.L0.ParentFile
																od.L0.Parent = md
																od.L0.Index = n

																for _, t := range oneofWrappers {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:251
				_go_fuzz_dep_.CoverTab[57745]++
																	if t.Implements(f.Type) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:252
					_go_fuzz_dep_.CoverTab[57746]++
																		f := t.Elem().Field(0)
																		if tag := f.Tag.Get("protobuf"); tag != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:254
						_go_fuzz_dep_.CoverTab[57747]++
																			aberrantAppendField(md, f.Type, tag, "", "")
																			fd := &md.L2.Fields.List[len(md.L2.Fields.List)-1]
																			fd.L1.ContainingOneof = od
																			od.L1.Fields.List = append(od.L1.Fields.List, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:258
						// _ = "end of CoverTab[57747]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:259
						_go_fuzz_dep_.CoverTab[57748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:259
						// _ = "end of CoverTab[57748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:259
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:259
					// _ = "end of CoverTab[57746]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:260
					_go_fuzz_dep_.CoverTab[57749]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:260
					// _ = "end of CoverTab[57749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:260
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:260
				// _ = "end of CoverTab[57745]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:261
			// _ = "end of CoverTab[57744]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:262
			_go_fuzz_dep_.CoverTab[57750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:262
			// _ = "end of CoverTab[57750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:262
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:262
		// _ = "end of CoverTab[57741]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:263
	// _ = "end of CoverTab[57714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:263
	_go_fuzz_dep_.CoverTab[57715]++

														return md
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:265
	// _ = "end of CoverTab[57715]"
}

func aberrantDeriveMessageName(t reflect.Type, name protoreflect.FullName) protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:268
	_go_fuzz_dep_.CoverTab[57751]++
														if name.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:269
		_go_fuzz_dep_.CoverTab[57756]++
															return name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:270
		// _ = "end of CoverTab[57756]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:271
		_go_fuzz_dep_.CoverTab[57757]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:271
		// _ = "end of CoverTab[57757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:271
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:271
	// _ = "end of CoverTab[57751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:271
	_go_fuzz_dep_.CoverTab[57752]++
														func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:272
		_go_fuzz_dep_.CoverTab[57758]++
															defer func() { _go_fuzz_dep_.CoverTab[57760]++; recover(); // _ = "end of CoverTab[57760]" }()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:273
		// _ = "end of CoverTab[57758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:273
		_go_fuzz_dep_.CoverTab[57759]++
															if m, ok := reflect.Zero(t).Interface().(interface{ XXX_MessageName() string }); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:274
			_go_fuzz_dep_.CoverTab[57761]++
																name = protoreflect.FullName(m.XXX_MessageName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:275
			// _ = "end of CoverTab[57761]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:276
			_go_fuzz_dep_.CoverTab[57762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:276
			// _ = "end of CoverTab[57762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:276
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:276
		// _ = "end of CoverTab[57759]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:277
	// _ = "end of CoverTab[57752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:277
	_go_fuzz_dep_.CoverTab[57753]++
														if name.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:278
		_go_fuzz_dep_.CoverTab[57763]++
															return name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:279
		// _ = "end of CoverTab[57763]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:280
		_go_fuzz_dep_.CoverTab[57764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:280
		// _ = "end of CoverTab[57764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:280
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:280
	// _ = "end of CoverTab[57753]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:280
	_go_fuzz_dep_.CoverTab[57754]++
														if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:281
		_go_fuzz_dep_.CoverTab[57765]++
															t = t.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:282
		// _ = "end of CoverTab[57765]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:283
		_go_fuzz_dep_.CoverTab[57766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:283
		// _ = "end of CoverTab[57766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:283
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:283
	// _ = "end of CoverTab[57754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:283
	_go_fuzz_dep_.CoverTab[57755]++
														return AberrantDeriveFullName(t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:284
	// _ = "end of CoverTab[57755]"
}

func aberrantAppendField(md *filedesc.Message, goType reflect.Type, tag, tagKey, tagVal string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:287
	_go_fuzz_dep_.CoverTab[57767]++
														t := goType
														isOptional := t.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:289
		_go_fuzz_dep_.CoverTab[57771]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:289
		return t.Elem().Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:289
		// _ = "end of CoverTab[57771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:289
	}()
														isRepeated := t.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:290
		_go_fuzz_dep_.CoverTab[57772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:290
		return t.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:290
		// _ = "end of CoverTab[57772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:290
	}()
														if isOptional || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:291
		_go_fuzz_dep_.CoverTab[57773]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:291
		return isRepeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:291
		// _ = "end of CoverTab[57773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:291
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:291
		_go_fuzz_dep_.CoverTab[57774]++
															t = t.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:292
		// _ = "end of CoverTab[57774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:293
		_go_fuzz_dep_.CoverTab[57775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:293
		// _ = "end of CoverTab[57775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:293
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:293
	// _ = "end of CoverTab[57767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:293
	_go_fuzz_dep_.CoverTab[57768]++
														fd := ptag.Unmarshal(tag, t, placeholderEnumValues{}).(*filedesc.Field)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:297
	n := len(md.L2.Fields.List)
	md.L2.Fields.List = append(md.L2.Fields.List, *fd)
	fd = &md.L2.Fields.List[n]
	fd.L0.FullName = md.FullName().Append(fd.Name())
	fd.L0.ParentFile = md.L0.ParentFile
	fd.L0.Parent = md
	fd.L0.Index = n

	if fd.L1.IsWeak || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:305
		_go_fuzz_dep_.CoverTab[57776]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:305
		return fd.L1.HasPacked
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:305
		// _ = "end of CoverTab[57776]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:305
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:305
		_go_fuzz_dep_.CoverTab[57777]++
															fd.L1.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:306
			_go_fuzz_dep_.CoverTab[57778]++
																opts := descopts.Field.ProtoReflect().New()
																if fd.L1.IsWeak {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:308
				_go_fuzz_dep_.CoverTab[57781]++
																	opts.Set(opts.Descriptor().Fields().ByName("weak"), protoreflect.ValueOfBool(true))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:309
				// _ = "end of CoverTab[57781]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:310
				_go_fuzz_dep_.CoverTab[57782]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:310
				// _ = "end of CoverTab[57782]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:310
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:310
			// _ = "end of CoverTab[57778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:310
			_go_fuzz_dep_.CoverTab[57779]++
																if fd.L1.HasPacked {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:311
				_go_fuzz_dep_.CoverTab[57783]++
																	opts.Set(opts.Descriptor().Fields().ByName("packed"), protoreflect.ValueOfBool(fd.L1.IsPacked))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:312
				// _ = "end of CoverTab[57783]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:313
				_go_fuzz_dep_.CoverTab[57784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:313
				// _ = "end of CoverTab[57784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:313
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:313
			// _ = "end of CoverTab[57779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:313
			_go_fuzz_dep_.CoverTab[57780]++
																return opts.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:314
			// _ = "end of CoverTab[57780]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:315
		// _ = "end of CoverTab[57777]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:316
		_go_fuzz_dep_.CoverTab[57785]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:316
		// _ = "end of CoverTab[57785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:316
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:316
	// _ = "end of CoverTab[57768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:316
	_go_fuzz_dep_.CoverTab[57769]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:319
	if fd.Enum() == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:319
		_go_fuzz_dep_.CoverTab[57786]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:319
		return fd.Kind() == protoreflect.EnumKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:319
		// _ = "end of CoverTab[57786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:319
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:319
		_go_fuzz_dep_.CoverTab[57787]++
															switch v := reflect.Zero(t).Interface().(type) {
		case protoreflect.Enum:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:321
			_go_fuzz_dep_.CoverTab[57788]++
																fd.L1.Enum = v.Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:322
			// _ = "end of CoverTab[57788]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:323
			_go_fuzz_dep_.CoverTab[57789]++
																fd.L1.Enum = LegacyLoadEnumDesc(t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:324
			// _ = "end of CoverTab[57789]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:325
		// _ = "end of CoverTab[57787]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:326
		_go_fuzz_dep_.CoverTab[57790]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:326
		// _ = "end of CoverTab[57790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:326
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:326
	// _ = "end of CoverTab[57769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:326
	_go_fuzz_dep_.CoverTab[57770]++
														if fd.Message() == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
		_go_fuzz_dep_.CoverTab[57791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
		return (fd.Kind() == protoreflect.MessageKind || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
			_go_fuzz_dep_.CoverTab[57792]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
			return fd.Kind() == protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
			// _ = "end of CoverTab[57792]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
		// _ = "end of CoverTab[57791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:327
		_go_fuzz_dep_.CoverTab[57793]++
															switch v := reflect.Zero(t).Interface().(type) {
		case protoreflect.ProtoMessage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:329
			_go_fuzz_dep_.CoverTab[57794]++
																fd.L1.Message = v.ProtoReflect().Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:330
			// _ = "end of CoverTab[57794]"
		case messageV1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:331
			_go_fuzz_dep_.CoverTab[57795]++
																fd.L1.Message = LegacyLoadMessageDesc(t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:332
			// _ = "end of CoverTab[57795]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:333
			_go_fuzz_dep_.CoverTab[57796]++
																if t.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:334
				_go_fuzz_dep_.CoverTab[57798]++
																	n := len(md.L1.Messages.List)
																	md.L1.Messages.List = append(md.L1.Messages.List, filedesc.Message{L2: new(filedesc.MessageL2)})
																	md2 := &md.L1.Messages.List[n]
																	md2.L0.FullName = md.FullName().Append(protoreflect.Name(strs.MapEntryName(string(fd.Name()))))
																	md2.L0.ParentFile = md.L0.ParentFile
																	md2.L0.Parent = md
																	md2.L0.Index = n

																	md2.L1.IsMapEntry = true
																	md2.L2.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:344
					_go_fuzz_dep_.CoverTab[57800]++
																		opts := descopts.Message.ProtoReflect().New()
																		opts.Set(opts.Descriptor().Fields().ByName("map_entry"), protoreflect.ValueOfBool(true))
																		return opts.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:347
					// _ = "end of CoverTab[57800]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:348
				// _ = "end of CoverTab[57798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:348
				_go_fuzz_dep_.CoverTab[57799]++

																	aberrantAppendField(md2, t.Key(), tagKey, "", "")
																	aberrantAppendField(md2, t.Elem(), tagVal, "", "")

																	fd.L1.Message = md2
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:354
				// _ = "end of CoverTab[57799]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:355
				_go_fuzz_dep_.CoverTab[57801]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:355
				// _ = "end of CoverTab[57801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:355
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:355
			// _ = "end of CoverTab[57796]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:355
			_go_fuzz_dep_.CoverTab[57797]++
																fd.L1.Message = aberrantLoadMessageDescReentrant(t, "")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:356
			// _ = "end of CoverTab[57797]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:357
		// _ = "end of CoverTab[57793]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:358
		_go_fuzz_dep_.CoverTab[57802]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:358
		// _ = "end of CoverTab[57802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:358
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:358
	// _ = "end of CoverTab[57770]"
}

type placeholderEnumValues struct {
	protoreflect.EnumValueDescriptors
}

func (placeholderEnumValues) ByNumber(n protoreflect.EnumNumber) protoreflect.EnumValueDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:365
	_go_fuzz_dep_.CoverTab[57803]++
														return filedesc.PlaceholderEnumValue(protoreflect.FullName(fmt.Sprintf("UNKNOWN_%d", n)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:366
	// _ = "end of CoverTab[57803]"
}

// legacyMarshaler is the proto.Marshaler interface superseded by protoiface.Methoder.
type legacyMarshaler interface {
	Marshal() ([]byte, error)
}

// legacyUnmarshaler is the proto.Unmarshaler interface superseded by protoiface.Methoder.
type legacyUnmarshaler interface {
	Unmarshal([]byte) error
}

// legacyMerger is the proto.Merger interface superseded by protoiface.Methoder.
type legacyMerger interface {
	Merge(protoiface.MessageV1)
}

var aberrantProtoMethods = &protoiface.Methods{
														Marshal:	legacyMarshal,
														Unmarshal:	legacyUnmarshal,
														Merge:		legacyMerge,

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:393
	Flags:	protoiface.SupportMarshalDeterministic,
}

func legacyMarshal(in protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:396
	_go_fuzz_dep_.CoverTab[57804]++
														v := in.Message.(unwrapper).protoUnwrap()
														marshaler, ok := v.(legacyMarshaler)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:399
		_go_fuzz_dep_.CoverTab[57807]++
															return protoiface.MarshalOutput{}, errors.New("%T does not implement Marshal", v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:400
		// _ = "end of CoverTab[57807]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:401
		_go_fuzz_dep_.CoverTab[57808]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:401
		// _ = "end of CoverTab[57808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:401
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:401
	// _ = "end of CoverTab[57804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:401
	_go_fuzz_dep_.CoverTab[57805]++
														out, err := marshaler.Marshal()
														if in.Buf != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:403
		_go_fuzz_dep_.CoverTab[57809]++
															out = append(in.Buf, out...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:404
		// _ = "end of CoverTab[57809]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:405
		_go_fuzz_dep_.CoverTab[57810]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:405
		// _ = "end of CoverTab[57810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:405
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:405
	// _ = "end of CoverTab[57805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:405
	_go_fuzz_dep_.CoverTab[57806]++
														return protoiface.MarshalOutput{
		Buf: out,
	}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:408
	// _ = "end of CoverTab[57806]"
}

func legacyUnmarshal(in protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:411
	_go_fuzz_dep_.CoverTab[57811]++
														v := in.Message.(unwrapper).protoUnwrap()
														unmarshaler, ok := v.(legacyUnmarshaler)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:414
		_go_fuzz_dep_.CoverTab[57813]++
															return protoiface.UnmarshalOutput{}, errors.New("%T does not implement Unmarshal", v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:415
		// _ = "end of CoverTab[57813]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:416
		_go_fuzz_dep_.CoverTab[57814]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:416
		// _ = "end of CoverTab[57814]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:416
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:416
	// _ = "end of CoverTab[57811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:416
	_go_fuzz_dep_.CoverTab[57812]++
														return protoiface.UnmarshalOutput{}, unmarshaler.Unmarshal(in.Buf)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:417
	// _ = "end of CoverTab[57812]"
}

func legacyMerge(in protoiface.MergeInput) protoiface.MergeOutput {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:420
	_go_fuzz_dep_.CoverTab[57815]++

														dstv := in.Destination.(unwrapper).protoUnwrap()
														merger, ok := dstv.(legacyMerger)
														if ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:424
		_go_fuzz_dep_.CoverTab[57822]++
															merger.Merge(Export{}.ProtoMessageV1Of(in.Source))
															return protoiface.MergeOutput{Flags: protoiface.MergeComplete}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:426
		// _ = "end of CoverTab[57822]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:427
		_go_fuzz_dep_.CoverTab[57823]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:427
		// _ = "end of CoverTab[57823]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:427
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:427
	// _ = "end of CoverTab[57815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:427
	_go_fuzz_dep_.CoverTab[57816]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:431
	srcv := in.Source.(unwrapper).protoUnwrap()
	marshaler, ok := srcv.(legacyMarshaler)
	if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:433
		_go_fuzz_dep_.CoverTab[57824]++
															return protoiface.MergeOutput{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:434
		// _ = "end of CoverTab[57824]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:435
		_go_fuzz_dep_.CoverTab[57825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:435
		// _ = "end of CoverTab[57825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:435
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:435
	// _ = "end of CoverTab[57816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:435
	_go_fuzz_dep_.CoverTab[57817]++
														dstv = in.Destination.(unwrapper).protoUnwrap()
														unmarshaler, ok := dstv.(legacyUnmarshaler)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:438
		_go_fuzz_dep_.CoverTab[57826]++
															return protoiface.MergeOutput{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:439
		// _ = "end of CoverTab[57826]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:440
		_go_fuzz_dep_.CoverTab[57827]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:440
		// _ = "end of CoverTab[57827]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:440
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:440
	// _ = "end of CoverTab[57817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:440
	_go_fuzz_dep_.CoverTab[57818]++
														if !in.Source.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:441
		_go_fuzz_dep_.CoverTab[57828]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:446
		return protoiface.MergeOutput{Flags: protoiface.MergeComplete}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:446
		// _ = "end of CoverTab[57828]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:447
		_go_fuzz_dep_.CoverTab[57829]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:447
		// _ = "end of CoverTab[57829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:447
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:447
	// _ = "end of CoverTab[57818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:447
	_go_fuzz_dep_.CoverTab[57819]++
														b, err := marshaler.Marshal()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:449
		_go_fuzz_dep_.CoverTab[57830]++
															return protoiface.MergeOutput{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:450
		// _ = "end of CoverTab[57830]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:451
		_go_fuzz_dep_.CoverTab[57831]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:451
		// _ = "end of CoverTab[57831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:451
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:451
	// _ = "end of CoverTab[57819]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:451
	_go_fuzz_dep_.CoverTab[57820]++
														err = unmarshaler.Unmarshal(b)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:453
		_go_fuzz_dep_.CoverTab[57832]++
															return protoiface.MergeOutput{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:454
		// _ = "end of CoverTab[57832]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:455
		_go_fuzz_dep_.CoverTab[57833]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:455
		// _ = "end of CoverTab[57833]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:455
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:455
	// _ = "end of CoverTab[57820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:455
	_go_fuzz_dep_.CoverTab[57821]++
														return protoiface.MergeOutput{Flags: protoiface.MergeComplete}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:456
	// _ = "end of CoverTab[57821]"
}

// aberrantMessageType implements MessageType for all types other than pointer-to-struct.
type aberrantMessageType struct {
	t reflect.Type
}

func (mt aberrantMessageType) New() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:464
	_go_fuzz_dep_.CoverTab[57834]++
														if mt.t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:465
		_go_fuzz_dep_.CoverTab[57836]++
															return aberrantMessage{reflect.New(mt.t.Elem())}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:466
		// _ = "end of CoverTab[57836]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:467
		_go_fuzz_dep_.CoverTab[57837]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:467
		// _ = "end of CoverTab[57837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:467
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:467
	// _ = "end of CoverTab[57834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:467
	_go_fuzz_dep_.CoverTab[57835]++
														return aberrantMessage{reflect.Zero(mt.t)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:468
	// _ = "end of CoverTab[57835]"
}
func (mt aberrantMessageType) Zero() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:470
	_go_fuzz_dep_.CoverTab[57838]++
														return aberrantMessage{reflect.Zero(mt.t)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:471
	// _ = "end of CoverTab[57838]"
}
func (mt aberrantMessageType) GoType() reflect.Type {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:473
	_go_fuzz_dep_.CoverTab[57839]++
														return mt.t
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:474
	// _ = "end of CoverTab[57839]"
}
func (mt aberrantMessageType) Descriptor() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:476
	_go_fuzz_dep_.CoverTab[57840]++
														return LegacyLoadMessageDesc(mt.t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:477
	// _ = "end of CoverTab[57840]"
}

// aberrantMessage implements Message for all types other than pointer-to-struct.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:480
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:480
// When the underlying type implements legacyMarshaler or legacyUnmarshaler,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:480
// the aberrant Message can be marshaled or unmarshaled. Otherwise, there is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:480
// not much that can be done with values of this type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:485
type aberrantMessage struct {
	v reflect.Value
}

// Reset implements the v1 proto.Message.Reset method.
func (m aberrantMessage) Reset() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:490
	_go_fuzz_dep_.CoverTab[57841]++
														if mr, ok := m.v.Interface().(interface{ Reset() }); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:491
		_go_fuzz_dep_.CoverTab[57843]++
															mr.Reset()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:493
		// _ = "end of CoverTab[57843]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:494
		_go_fuzz_dep_.CoverTab[57844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:494
		// _ = "end of CoverTab[57844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:494
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:494
	// _ = "end of CoverTab[57841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:494
	_go_fuzz_dep_.CoverTab[57842]++
														if m.v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:495
		_go_fuzz_dep_.CoverTab[57845]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:495
		return !m.v.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:495
		// _ = "end of CoverTab[57845]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:495
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:495
		_go_fuzz_dep_.CoverTab[57846]++
															m.v.Elem().Set(reflect.Zero(m.v.Type().Elem()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:496
		// _ = "end of CoverTab[57846]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:497
		_go_fuzz_dep_.CoverTab[57847]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:497
		// _ = "end of CoverTab[57847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:497
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:497
	// _ = "end of CoverTab[57842]"
}

func (m aberrantMessage) ProtoReflect() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:500
	_go_fuzz_dep_.CoverTab[57848]++
														return m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:501
	// _ = "end of CoverTab[57848]"
}

func (m aberrantMessage) Descriptor() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:504
	_go_fuzz_dep_.CoverTab[57849]++
														return LegacyLoadMessageDesc(m.v.Type())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:505
	// _ = "end of CoverTab[57849]"
}
func (m aberrantMessage) Type() protoreflect.MessageType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:507
	_go_fuzz_dep_.CoverTab[57850]++
														return aberrantMessageType{m.v.Type()}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:508
	// _ = "end of CoverTab[57850]"
}
func (m aberrantMessage) New() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:510
	_go_fuzz_dep_.CoverTab[57851]++
														if m.v.Type().Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:511
		_go_fuzz_dep_.CoverTab[57853]++
															return aberrantMessage{reflect.New(m.v.Type().Elem())}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:512
		// _ = "end of CoverTab[57853]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:513
		_go_fuzz_dep_.CoverTab[57854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:513
		// _ = "end of CoverTab[57854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:513
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:513
	// _ = "end of CoverTab[57851]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:513
	_go_fuzz_dep_.CoverTab[57852]++
														return aberrantMessage{reflect.Zero(m.v.Type())}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:514
	// _ = "end of CoverTab[57852]"
}
func (m aberrantMessage) Interface() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:516
	_go_fuzz_dep_.CoverTab[57855]++
														return m
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:517
	// _ = "end of CoverTab[57855]"
}
func (m aberrantMessage) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:519
	_go_fuzz_dep_.CoverTab[57856]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:520
	// _ = "end of CoverTab[57856]"
}
func (m aberrantMessage) Has(protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:522
	_go_fuzz_dep_.CoverTab[57857]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:523
	// _ = "end of CoverTab[57857]"
}
func (m aberrantMessage) Clear(protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:525
	_go_fuzz_dep_.CoverTab[57858]++
														panic("invalid Message.Clear on " + string(m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:526
	// _ = "end of CoverTab[57858]"
}
func (m aberrantMessage) Get(fd protoreflect.FieldDescriptor) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:528
	_go_fuzz_dep_.CoverTab[57859]++
														if fd.Default().IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:529
		_go_fuzz_dep_.CoverTab[57861]++
															return fd.Default()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:530
		// _ = "end of CoverTab[57861]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:531
		_go_fuzz_dep_.CoverTab[57862]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:531
		// _ = "end of CoverTab[57862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:531
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:531
	// _ = "end of CoverTab[57859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:531
	_go_fuzz_dep_.CoverTab[57860]++
														panic("invalid Message.Get on " + string(m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:532
	// _ = "end of CoverTab[57860]"
}
func (m aberrantMessage) Set(protoreflect.FieldDescriptor, protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:534
	_go_fuzz_dep_.CoverTab[57863]++
														panic("invalid Message.Set on " + string(m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:535
	// _ = "end of CoverTab[57863]"
}
func (m aberrantMessage) Mutable(protoreflect.FieldDescriptor) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:537
	_go_fuzz_dep_.CoverTab[57864]++
														panic("invalid Message.Mutable on " + string(m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:538
	// _ = "end of CoverTab[57864]"
}
func (m aberrantMessage) NewField(protoreflect.FieldDescriptor) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:540
	_go_fuzz_dep_.CoverTab[57865]++
														panic("invalid Message.NewField on " + string(m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:541
	// _ = "end of CoverTab[57865]"
}
func (m aberrantMessage) WhichOneof(protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:543
	_go_fuzz_dep_.CoverTab[57866]++
														panic("invalid Message.WhichOneof descriptor on " + string(m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:544
	// _ = "end of CoverTab[57866]"
}
func (m aberrantMessage) GetUnknown() protoreflect.RawFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:546
	_go_fuzz_dep_.CoverTab[57867]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:547
	// _ = "end of CoverTab[57867]"
}
func (m aberrantMessage) SetUnknown(protoreflect.RawFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:549
	_go_fuzz_dep_.CoverTab[57868]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:549
	// _ = "end of CoverTab[57868]"

}
func (m aberrantMessage) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:552
	_go_fuzz_dep_.CoverTab[57869]++
														if m.v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:553
		_go_fuzz_dep_.CoverTab[57871]++
															return !m.v.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:554
		// _ = "end of CoverTab[57871]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:555
		_go_fuzz_dep_.CoverTab[57872]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:555
		// _ = "end of CoverTab[57872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:555
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:555
	// _ = "end of CoverTab[57869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:555
	_go_fuzz_dep_.CoverTab[57870]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:556
	// _ = "end of CoverTab[57870]"
}
func (m aberrantMessage) ProtoMethods() *protoiface.Methods {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:558
	_go_fuzz_dep_.CoverTab[57873]++
														return aberrantProtoMethods
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:559
	// _ = "end of CoverTab[57873]"
}
func (m aberrantMessage) protoUnwrap() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:561
	_go_fuzz_dep_.CoverTab[57874]++
														return m.v.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:562
	// _ = "end of CoverTab[57874]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:563
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_message.go:563
var _ = _go_fuzz_dep_.CoverTab
