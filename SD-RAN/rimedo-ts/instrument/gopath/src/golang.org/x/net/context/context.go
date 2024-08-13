// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Package context defines the Context type, which carries deadlines,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// cancelation signals, and other request-scoped values across API boundaries
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// and between processes.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// As of Go 1.7 this package is available in the standard library under the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// name context.  https://golang.org/pkg/context.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Incoming requests to a server should create a Context, and outgoing calls to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// servers should accept a Context. The chain of function calls between must
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// propagate the Context, optionally replacing it with a modified copy created
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// using WithDeadline, WithTimeout, WithCancel, or WithValue.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Programs that use Contexts should follow these rules to keep interfaces
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// consistent across packages and enable static analysis tools to check context
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// propagation:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Do not store Contexts inside a struct type; instead, pass a Context
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// explicitly to each function that needs it. The Context should be the first
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// parameter, typically named ctx:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//	func DoSomething(ctx context.Context, arg Arg) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//		// ... use ctx ...
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Do not pass a nil Context, even if a function permits it. Pass context.TODO
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// if you are unsure about which Context to use.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Use context Values only for request-scoped data that transits processes and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// APIs, not for passing optional parameters to functions.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// The same Context may be passed to functions running in different goroutines;
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Contexts are safe for simultaneous use by multiple goroutines.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// See http://blog.golang.org/context for example code for a server that uses
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:5
// Contexts.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
package context

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:39
)

// Background returns a non-nil, empty Context. It is never canceled, has no
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:41
// values, and has no deadline. It is typically used by the main function,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:41
// initialization, and tests, and as the top-level Context for incoming
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:41
// requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:45
func Background() Context {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:45
	_go_fuzz_dep_.CoverTab[131943]++
											return background
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:46
	// _ = "end of CoverTab[131943]"
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:49
// it's unclear which Context to use or it is not yet available (because the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:49
// surrounding function has not yet been extended to accept a Context
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:49
// parameter).  TODO is recognized by static analysis tools that determine
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:49
// whether Contexts are propagated correctly in a program.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:54
func TODO() Context {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:54
	_go_fuzz_dep_.CoverTab[131944]++
											return todo
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:55
	// _ = "end of CoverTab[131944]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:56
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/context.go:56
var _ = _go_fuzz_dep_.CoverTab
