// Copyright 2016 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine
// +build gc
// +build !noasm
// +build amd64 arm64

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
package snappy

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:10
)

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:14
//go:noescape
func decode(dst, src []byte) int

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:15
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode_asm.go:15
var _ = _go_fuzz_dep_.CoverTab
