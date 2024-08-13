// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.7
// +build go1.7

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
package context

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:8
)

import (
	"context"	// standard library's context, as of Go 1.7
	"time"
)

var (
	todo		= context.TODO()
	background	= context.Background()
)

// Canceled is the error returned by Context.Err when the context is canceled.
var Canceled = context.Canceled

// DeadlineExceeded is the error returned by Context.Err when the context's
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:23
// deadline passes.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:25
var DeadlineExceeded = context.DeadlineExceeded

// WithCancel returns a copy of parent with a new Done channel. The returned
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:27
// context's Done channel is closed when the returned cancel function is called
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:27
// or when the parent context's Done channel is closed, whichever happens first.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:27
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:27
// Canceling this context releases resources associated with it, so code should
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:27
// call cancel as soon as the operations running in this Context complete.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:33
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:33
	_go_fuzz_dep_.CoverTab[131945]++
										ctx, f := context.WithCancel(parent)
										return ctx, f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:35
	// _ = "end of CoverTab[131945]"
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// to be no later than d. If the parent's deadline is already earlier than d,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// context's Done channel is closed when the deadline expires, when the returned
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// cancel function is called, or when the parent context's Done channel is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// closed, whichever happens first.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// Canceling this context releases resources associated with it, so code should
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:38
// call cancel as soon as the operations running in this Context complete.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:47
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:47
	_go_fuzz_dep_.CoverTab[131946]++
										ctx, f := context.WithDeadline(parent, deadline)
										return ctx, f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:49
	// _ = "end of CoverTab[131946]"
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
// Canceling this context releases resources associated with it, so code should
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
// call cancel as soon as the operations running in this Context complete:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//		defer cancel()  // releases resources if slowOperation completes before timeout elapses
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//		return slowOperation(ctx)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:52
//	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:62
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:62
	_go_fuzz_dep_.CoverTab[131947]++
										return WithDeadline(parent, time.Now().Add(timeout))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:63
	// _ = "end of CoverTab[131947]"
}

// WithValue returns a copy of parent in which the value associated with key is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:66
// val.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:66
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:66
// Use context Values only for request-scoped data that transits processes and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:66
// APIs, not for passing optional parameters to functions.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:71
func WithValue(parent Context, key interface{}, val interface{}) Context {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:71
	_go_fuzz_dep_.CoverTab[131948]++
										return context.WithValue(parent, key, val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:72
	// _ = "end of CoverTab[131948]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/go17.go:73
var _ = _go_fuzz_dep_.CoverTab
