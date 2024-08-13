// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.9
// +build go1.9

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
package context

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:8
)

import "context"	// standard library's context, as of Go 1.7

// A Context carries a deadline, a cancelation signal, and other values across
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:12
// API boundaries.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:12
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:12
// Context's methods may be called by multiple goroutines simultaneously.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:16
type Context = context.Context

// A CancelFunc tells an operation to abandon its work.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:18
// A CancelFunc does not wait for the work to stop.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:18
// After the first call, subsequent calls to a CancelFunc do nothing.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:21
type CancelFunc = context.CancelFunc

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:21
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go19.go:21
var _ = _go_fuzz_dep_.CoverTab
