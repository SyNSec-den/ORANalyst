// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:5
)

import (
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type EnumInfo struct {
	GoReflectType	reflect.Type	// int32 kind
	Desc		protoreflect.EnumDescriptor
}

func (t *EnumInfo) New(n protoreflect.EnumNumber) protoreflect.Enum {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:18
	_go_fuzz_dep_.CoverTab[57420]++
												return reflect.ValueOf(n).Convert(t.GoReflectType).Interface().(protoreflect.Enum)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:19
	// _ = "end of CoverTab[57420]"
}
func (t *EnumInfo) Descriptor() protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:21
	_go_fuzz_dep_.CoverTab[57421]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:21
	return t.Desc
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:21
	// _ = "end of CoverTab[57421]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:21
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:21
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/enum.go:21
var _ = _go_fuzz_dep_.CoverTab
