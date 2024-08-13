// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
package genid

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:5
)

import protoreflect "google.golang.org/protobuf/reflect/protoreflect"

// Generic field names and numbers for synthetic map entry messages.
const (
	MapEntry_Key_field_name		protoreflect.Name	= "key"
	MapEntry_Value_field_name	protoreflect.Name	= "value"

	MapEntry_Key_field_number	protoreflect.FieldNumber	= 1
	MapEntry_Value_field_number	protoreflect.FieldNumber	= 2
)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/genid/map_entry.go:16
var _ = _go_fuzz_dep_.CoverTab
