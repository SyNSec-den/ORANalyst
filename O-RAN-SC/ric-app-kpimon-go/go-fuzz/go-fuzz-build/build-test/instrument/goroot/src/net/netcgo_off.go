// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !netcgo

//line /snap/go/10455/src/net/netcgo_off.go:7
package net

//line /snap/go/10455/src/net/netcgo_off.go:7
import (
//line /snap/go/10455/src/net/netcgo_off.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/netcgo_off.go:7
)
//line /snap/go/10455/src/net/netcgo_off.go:7
import (
//line /snap/go/10455/src/net/netcgo_off.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/netcgo_off.go:7
)

const netCgoBuildTag = false

//line /snap/go/10455/src/net/netcgo_off.go:9
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/netcgo_off.go:9
var _ = _go_fuzz_dep_.CoverTab
