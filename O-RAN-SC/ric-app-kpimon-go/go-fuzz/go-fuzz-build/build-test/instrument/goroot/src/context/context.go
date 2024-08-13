// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/context/context.go:5
// Package context defines the Context type, which carries deadlines,
//line /snap/go/10455/src/context/context.go:5
// cancellation signals, and other request-scoped values across API boundaries
//line /snap/go/10455/src/context/context.go:5
// and between processes.
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// Incoming requests to a server should create a [Context], and outgoing
//line /snap/go/10455/src/context/context.go:5
// calls to servers should accept a Context. The chain of function
//line /snap/go/10455/src/context/context.go:5
// calls between them must propagate the Context, optionally replacing
//line /snap/go/10455/src/context/context.go:5
// it with a derived Context created using [WithCancel], [WithDeadline],
//line /snap/go/10455/src/context/context.go:5
// [WithTimeout], or [WithValue]. When a Context is canceled, all
//line /snap/go/10455/src/context/context.go:5
// Contexts derived from it are also canceled.
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// The [WithCancel], [WithDeadline], and [WithTimeout] functions take a
//line /snap/go/10455/src/context/context.go:5
// Context (the parent) and return a derived Context (the child) and a
//line /snap/go/10455/src/context/context.go:5
// [CancelFunc]. Calling the CancelFunc cancels the child and its
//line /snap/go/10455/src/context/context.go:5
// children, removes the parent's reference to the child, and stops
//line /snap/go/10455/src/context/context.go:5
// any associated timers. Failing to call the CancelFunc leaks the
//line /snap/go/10455/src/context/context.go:5
// child and its children until the parent is canceled or the timer
//line /snap/go/10455/src/context/context.go:5
// fires. The go vet tool checks that CancelFuncs are used on all
//line /snap/go/10455/src/context/context.go:5
// control-flow paths.
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// The [WithCancelCause] function returns a [CancelCauseFunc], which
//line /snap/go/10455/src/context/context.go:5
// takes an error and records it as the cancellation cause. Calling
//line /snap/go/10455/src/context/context.go:5
// [Cause] on the canceled context or any of its children retrieves
//line /snap/go/10455/src/context/context.go:5
// the cause. If no cause is specified, Cause(ctx) returns the same
//line /snap/go/10455/src/context/context.go:5
// value as ctx.Err().
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// Programs that use Contexts should follow these rules to keep interfaces
//line /snap/go/10455/src/context/context.go:5
// consistent across packages and enable static analysis tools to check context
//line /snap/go/10455/src/context/context.go:5
// propagation:
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// Do not store Contexts inside a struct type; instead, pass a Context
//line /snap/go/10455/src/context/context.go:5
// explicitly to each function that needs it. The Context should be the first
//line /snap/go/10455/src/context/context.go:5
// parameter, typically named ctx:
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
//	func DoSomething(ctx context.Context, arg Arg) error {
//line /snap/go/10455/src/context/context.go:5
//		// ... use ctx ...
//line /snap/go/10455/src/context/context.go:5
//	}
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// Do not pass a nil [Context], even if a function permits it. Pass [context.TODO]
//line /snap/go/10455/src/context/context.go:5
// if you are unsure about which Context to use.
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// Use context Values only for request-scoped data that transits processes and
//line /snap/go/10455/src/context/context.go:5
// APIs, not for passing optional parameters to functions.
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// The same Context may be passed to functions running in different goroutines;
//line /snap/go/10455/src/context/context.go:5
// Contexts are safe for simultaneous use by multiple goroutines.
//line /snap/go/10455/src/context/context.go:5
//
//line /snap/go/10455/src/context/context.go:5
// See https://blog.golang.org/context for example code for a server that uses
//line /snap/go/10455/src/context/context.go:5
// Contexts.
//line /snap/go/10455/src/context/context.go:54
package context

//line /snap/go/10455/src/context/context.go:54
import (
//line /snap/go/10455/src/context/context.go:54
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/context/context.go:54
)
//line /snap/go/10455/src/context/context.go:54
import (
//line /snap/go/10455/src/context/context.go:54
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/context/context.go:54
)

import (
	"errors"
	"internal/reflectlite"
	"sync"
	"sync/atomic"
	"time"
)

// A Context carries a deadline, a cancellation signal, and other values across
//line /snap/go/10455/src/context/context.go:64
// API boundaries.
//line /snap/go/10455/src/context/context.go:64
//
//line /snap/go/10455/src/context/context.go:64
// Context's methods may be called by multiple goroutines simultaneously.
//line /snap/go/10455/src/context/context.go:68
type Context interface {
	// Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
	Deadline() (deadline time.Time, ok bool)

	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	// The close of the Done channel may happen asynchronously,
	// after the cancel function returns.
	//
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	//
	// Done is provided for use in select statements:
	//
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
	// See https://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancellation.
	Done() <-chan struct{}

	// If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	// After Err returns a non-nil error, successive calls to Err return the same error.
	Err() error

	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	//
	// Use context values only for request-scoped data that transits
	// processes and API boundaries, not for passing optional parameters to
	// functions.
	//
	// A key identifies a specific value in a Context. Functions that wish
	// to store values in Context typically allocate a key in a global
	// variable then use that key as the argument to context.WithValue and
	// Context.Value. A key can be any type that supports equality;
	// packages should define keys as an unexported type to avoid
	// collisions.
	//
	// Packages that define a Context key should provide type-safe accessors
	// for the values stored using that key:
	//
	// 	// Package user defines a User type that's stored in Contexts.
	// 	package user
	//
	// 	import "context"
	//
	// 	// User is the type of value stored in the Contexts.
	// 	type User struct {...}
	//
	// 	// key is an unexported type for keys defined in this package.
	// 	// This prevents collisions with keys defined in other packages.
	// 	type key int
	//
	// 	// userKey is the key for user.User values in Contexts. It is
	// 	// unexported; clients use user.NewContext and user.FromContext
	// 	// instead of using this key directly.
	// 	var userKey key
	//
	// 	// NewContext returns a new Context that carries value u.
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 		return context.WithValue(ctx, userKey, u)
	// 	}
	//
	// 	// FromContext returns the User value stored in ctx, if any.
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	Value(key any) any
}

// Canceled is the error returned by [Context.Err] when the context is canceled.
var Canceled = errors.New("context canceled")

// DeadlineExceeded is the error returned by [Context.Err] when the context's
//line /snap/go/10455/src/context/context.go:165
// deadline passes.
//line /snap/go/10455/src/context/context.go:167
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string {
//line /snap/go/10455/src/context/context.go:171
	_go_fuzz_dep_.CoverTab[2307]++
//line /snap/go/10455/src/context/context.go:171
	return "context deadline exceeded"
//line /snap/go/10455/src/context/context.go:171
	// _ = "end of CoverTab[2307]"
//line /snap/go/10455/src/context/context.go:171
}
func (deadlineExceededError) Timeout() bool {
//line /snap/go/10455/src/context/context.go:172
	_go_fuzz_dep_.CoverTab[2308]++
//line /snap/go/10455/src/context/context.go:172
	return true
//line /snap/go/10455/src/context/context.go:172
	// _ = "end of CoverTab[2308]"
//line /snap/go/10455/src/context/context.go:172
}
func (deadlineExceededError) Temporary() bool {
//line /snap/go/10455/src/context/context.go:173
	_go_fuzz_dep_.CoverTab[2309]++
//line /snap/go/10455/src/context/context.go:173
	return true
//line /snap/go/10455/src/context/context.go:173
	// _ = "end of CoverTab[2309]"
//line /snap/go/10455/src/context/context.go:173
}

// An emptyCtx is never canceled, has no values, and has no deadline.
//line /snap/go/10455/src/context/context.go:175
// It is the common base of backgroundCtx and todoCtx.
//line /snap/go/10455/src/context/context.go:177
type emptyCtx struct{}

func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
//line /snap/go/10455/src/context/context.go:179
	_go_fuzz_dep_.CoverTab[2310]++
							return
//line /snap/go/10455/src/context/context.go:180
	// _ = "end of CoverTab[2310]"
}

func (emptyCtx) Done() <-chan struct{} {
//line /snap/go/10455/src/context/context.go:183
	_go_fuzz_dep_.CoverTab[2311]++
							return nil
//line /snap/go/10455/src/context/context.go:184
	// _ = "end of CoverTab[2311]"
}

func (emptyCtx) Err() error {
//line /snap/go/10455/src/context/context.go:187
	_go_fuzz_dep_.CoverTab[2312]++
							return nil
//line /snap/go/10455/src/context/context.go:188
	// _ = "end of CoverTab[2312]"
}

func (emptyCtx) Value(key any) any {
//line /snap/go/10455/src/context/context.go:191
	_go_fuzz_dep_.CoverTab[2313]++
							return nil
//line /snap/go/10455/src/context/context.go:192
	// _ = "end of CoverTab[2313]"
}

type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
//line /snap/go/10455/src/context/context.go:197
	_go_fuzz_dep_.CoverTab[2314]++
							return "context.Background"
//line /snap/go/10455/src/context/context.go:198
	// _ = "end of CoverTab[2314]"
}

type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
//line /snap/go/10455/src/context/context.go:203
	_go_fuzz_dep_.CoverTab[2315]++
							return "context.TODO"
//line /snap/go/10455/src/context/context.go:204
	// _ = "end of CoverTab[2315]"
}

// Background returns a non-nil, empty [Context]. It is never canceled, has no
//line /snap/go/10455/src/context/context.go:207
// values, and has no deadline. It is typically used by the main function,
//line /snap/go/10455/src/context/context.go:207
// initialization, and tests, and as the top-level Context for incoming
//line /snap/go/10455/src/context/context.go:207
// requests.
//line /snap/go/10455/src/context/context.go:211
func Background() Context {
//line /snap/go/10455/src/context/context.go:211
	_go_fuzz_dep_.CoverTab[2316]++
							return backgroundCtx{}
//line /snap/go/10455/src/context/context.go:212
	// _ = "end of CoverTab[2316]"
}

// TODO returns a non-nil, empty [Context]. Code should use context.TODO when
//line /snap/go/10455/src/context/context.go:215
// it's unclear which Context to use or it is not yet available (because the
//line /snap/go/10455/src/context/context.go:215
// surrounding function has not yet been extended to accept a Context
//line /snap/go/10455/src/context/context.go:215
// parameter).
//line /snap/go/10455/src/context/context.go:219
func TODO() Context {
//line /snap/go/10455/src/context/context.go:219
	_go_fuzz_dep_.CoverTab[2317]++
							return todoCtx{}
//line /snap/go/10455/src/context/context.go:220
	// _ = "end of CoverTab[2317]"
}

// A CancelFunc tells an operation to abandon its work.
//line /snap/go/10455/src/context/context.go:223
// A CancelFunc does not wait for the work to stop.
//line /snap/go/10455/src/context/context.go:223
// A CancelFunc may be called by multiple goroutines simultaneously.
//line /snap/go/10455/src/context/context.go:223
// After the first call, subsequent calls to a CancelFunc do nothing.
//line /snap/go/10455/src/context/context.go:227
type CancelFunc func()

// WithCancel returns a copy of parent with a new Done channel. The returned
//line /snap/go/10455/src/context/context.go:229
// context's Done channel is closed when the returned cancel function is called
//line /snap/go/10455/src/context/context.go:229
// or when the parent context's Done channel is closed, whichever happens first.
//line /snap/go/10455/src/context/context.go:229
//
//line /snap/go/10455/src/context/context.go:229
// Canceling this context releases resources associated with it, so code should
//line /snap/go/10455/src/context/context.go:229
// call cancel as soon as the operations running in this Context complete.
//line /snap/go/10455/src/context/context.go:235
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
//line /snap/go/10455/src/context/context.go:235
	_go_fuzz_dep_.CoverTab[2318]++
							c := withCancel(parent)
							return c, func() { _go_fuzz_dep_.CoverTab[2319]++; c.cancel(true, Canceled, nil); // _ = "end of CoverTab[2319]" }
//line /snap/go/10455/src/context/context.go:237
	// _ = "end of CoverTab[2318]"
}

// A CancelCauseFunc behaves like a [CancelFunc] but additionally sets the cancellation cause.
//line /snap/go/10455/src/context/context.go:240
// This cause can be retrieved by calling [Cause] on the canceled Context or on
//line /snap/go/10455/src/context/context.go:240
// any of its derived Contexts.
//line /snap/go/10455/src/context/context.go:240
//
//line /snap/go/10455/src/context/context.go:240
// If the context has already been canceled, CancelCauseFunc does not set the cause.
//line /snap/go/10455/src/context/context.go:240
// For example, if childContext is derived from parentContext:
//line /snap/go/10455/src/context/context.go:240
//   - if parentContext is canceled with cause1 before childContext is canceled with cause2,
//line /snap/go/10455/src/context/context.go:240
//     then Cause(parentContext) == Cause(childContext) == cause1
//line /snap/go/10455/src/context/context.go:240
//   - if childContext is canceled with cause2 before parentContext is canceled with cause1,
//line /snap/go/10455/src/context/context.go:240
//     then Cause(parentContext) == cause1 and Cause(childContext) == cause2
//line /snap/go/10455/src/context/context.go:250
type CancelCauseFunc func(cause error)

// WithCancelCause behaves like [WithCancel] but returns a [CancelCauseFunc] instead of a [CancelFunc].
//line /snap/go/10455/src/context/context.go:252
// Calling cancel with a non-nil error (the "cause") records that error in ctx;
//line /snap/go/10455/src/context/context.go:252
// it can then be retrieved using Cause(ctx).
//line /snap/go/10455/src/context/context.go:252
// Calling cancel with nil sets the cause to Canceled.
//line /snap/go/10455/src/context/context.go:252
//
//line /snap/go/10455/src/context/context.go:252
// Example use:
//line /snap/go/10455/src/context/context.go:252
//
//line /snap/go/10455/src/context/context.go:252
//	ctx, cancel := context.WithCancelCause(parent)
//line /snap/go/10455/src/context/context.go:252
//	cancel(myError)
//line /snap/go/10455/src/context/context.go:252
//	ctx.Err() // returns context.Canceled
//line /snap/go/10455/src/context/context.go:252
//	context.Cause(ctx) // returns myError
//line /snap/go/10455/src/context/context.go:263
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc) {
//line /snap/go/10455/src/context/context.go:263
	_go_fuzz_dep_.CoverTab[2320]++
							c := withCancel(parent)
							return c, func(cause error) {
//line /snap/go/10455/src/context/context.go:265
		_go_fuzz_dep_.CoverTab[2321]++
//line /snap/go/10455/src/context/context.go:265
		c.cancel(true, Canceled, cause)
//line /snap/go/10455/src/context/context.go:265
		// _ = "end of CoverTab[2321]"
//line /snap/go/10455/src/context/context.go:265
	}
//line /snap/go/10455/src/context/context.go:265
	// _ = "end of CoverTab[2320]"
}

func withCancel(parent Context) *cancelCtx {
//line /snap/go/10455/src/context/context.go:268
	_go_fuzz_dep_.CoverTab[2322]++
							if parent == nil {
//line /snap/go/10455/src/context/context.go:269
		_go_fuzz_dep_.CoverTab[525998]++
//line /snap/go/10455/src/context/context.go:269
		_go_fuzz_dep_.CoverTab[2324]++
								panic("cannot create context from nil parent")
//line /snap/go/10455/src/context/context.go:270
		// _ = "end of CoverTab[2324]"
	} else {
//line /snap/go/10455/src/context/context.go:271
		_go_fuzz_dep_.CoverTab[525999]++
//line /snap/go/10455/src/context/context.go:271
		_go_fuzz_dep_.CoverTab[2325]++
//line /snap/go/10455/src/context/context.go:271
		// _ = "end of CoverTab[2325]"
//line /snap/go/10455/src/context/context.go:271
	}
//line /snap/go/10455/src/context/context.go:271
	// _ = "end of CoverTab[2322]"
//line /snap/go/10455/src/context/context.go:271
	_go_fuzz_dep_.CoverTab[2323]++
							c := &cancelCtx{}
							c.propagateCancel(parent, c)
							return c
//line /snap/go/10455/src/context/context.go:274
	// _ = "end of CoverTab[2323]"
}

// Cause returns a non-nil error explaining why c was canceled.
//line /snap/go/10455/src/context/context.go:277
// The first cancellation of c or one of its parents sets the cause.
//line /snap/go/10455/src/context/context.go:277
// If that cancellation happened via a call to CancelCauseFunc(err),
//line /snap/go/10455/src/context/context.go:277
// then [Cause] returns err.
//line /snap/go/10455/src/context/context.go:277
// Otherwise Cause(c) returns the same value as c.Err().
//line /snap/go/10455/src/context/context.go:277
// Cause returns nil if c has not been canceled yet.
//line /snap/go/10455/src/context/context.go:283
func Cause(c Context) error {
//line /snap/go/10455/src/context/context.go:283
	_go_fuzz_dep_.CoverTab[2326]++
							if cc, ok := c.Value(&cancelCtxKey).(*cancelCtx); ok {
//line /snap/go/10455/src/context/context.go:284
		_go_fuzz_dep_.CoverTab[526000]++
//line /snap/go/10455/src/context/context.go:284
		_go_fuzz_dep_.CoverTab[2328]++
								cc.mu.Lock()
								defer cc.mu.Unlock()
								return cc.cause
//line /snap/go/10455/src/context/context.go:287
		// _ = "end of CoverTab[2328]"
	} else {
//line /snap/go/10455/src/context/context.go:288
		_go_fuzz_dep_.CoverTab[526001]++
//line /snap/go/10455/src/context/context.go:288
		_go_fuzz_dep_.CoverTab[2329]++
//line /snap/go/10455/src/context/context.go:288
		// _ = "end of CoverTab[2329]"
//line /snap/go/10455/src/context/context.go:288
	}
//line /snap/go/10455/src/context/context.go:288
	// _ = "end of CoverTab[2326]"
//line /snap/go/10455/src/context/context.go:288
	_go_fuzz_dep_.CoverTab[2327]++
							return nil
//line /snap/go/10455/src/context/context.go:289
	// _ = "end of CoverTab[2327]"
}

// AfterFunc arranges to call f in its own goroutine after ctx is done
//line /snap/go/10455/src/context/context.go:292
// (cancelled or timed out).
//line /snap/go/10455/src/context/context.go:292
// If ctx is already done, AfterFunc calls f immediately in its own goroutine.
//line /snap/go/10455/src/context/context.go:292
//
//line /snap/go/10455/src/context/context.go:292
// Multiple calls to AfterFunc on a context operate independently;
//line /snap/go/10455/src/context/context.go:292
// one does not replace another.
//line /snap/go/10455/src/context/context.go:292
//
//line /snap/go/10455/src/context/context.go:292
// Calling the returned stop function stops the association of ctx with f.
//line /snap/go/10455/src/context/context.go:292
// It returns true if the call stopped f from being run.
//line /snap/go/10455/src/context/context.go:292
// If stop returns false,
//line /snap/go/10455/src/context/context.go:292
// either the context is done and f has been started in its own goroutine;
//line /snap/go/10455/src/context/context.go:292
// or f was already stopped.
//line /snap/go/10455/src/context/context.go:292
// The stop function does not wait for f to complete before returning.
//line /snap/go/10455/src/context/context.go:292
// If the caller needs to know whether f is completed,
//line /snap/go/10455/src/context/context.go:292
// it must coordinate with f explicitly.
//line /snap/go/10455/src/context/context.go:292
//
//line /snap/go/10455/src/context/context.go:292
// If ctx has a "AfterFunc(func()) func() bool" method,
//line /snap/go/10455/src/context/context.go:292
// AfterFunc will use it to schedule the call.
//line /snap/go/10455/src/context/context.go:310
func AfterFunc(ctx Context, f func()) (stop func() bool) {
//line /snap/go/10455/src/context/context.go:310
	_go_fuzz_dep_.CoverTab[2330]++
							a := &afterFuncCtx{
		f: f,
	}
	a.cancelCtx.propagateCancel(ctx, a)
	return func() bool {
//line /snap/go/10455/src/context/context.go:315
		_go_fuzz_dep_.CoverTab[2331]++
								stopped := false
								a.once.Do(func() {
//line /snap/go/10455/src/context/context.go:317
			_go_fuzz_dep_.CoverTab[2334]++
									stopped = true
//line /snap/go/10455/src/context/context.go:318
			// _ = "end of CoverTab[2334]"
		})
//line /snap/go/10455/src/context/context.go:319
		// _ = "end of CoverTab[2331]"
//line /snap/go/10455/src/context/context.go:319
		_go_fuzz_dep_.CoverTab[2332]++
								if stopped {
//line /snap/go/10455/src/context/context.go:320
			_go_fuzz_dep_.CoverTab[526002]++
//line /snap/go/10455/src/context/context.go:320
			_go_fuzz_dep_.CoverTab[2335]++
									a.cancel(true, Canceled, nil)
//line /snap/go/10455/src/context/context.go:321
			// _ = "end of CoverTab[2335]"
		} else {
//line /snap/go/10455/src/context/context.go:322
			_go_fuzz_dep_.CoverTab[526003]++
//line /snap/go/10455/src/context/context.go:322
			_go_fuzz_dep_.CoverTab[2336]++
//line /snap/go/10455/src/context/context.go:322
			// _ = "end of CoverTab[2336]"
//line /snap/go/10455/src/context/context.go:322
		}
//line /snap/go/10455/src/context/context.go:322
		// _ = "end of CoverTab[2332]"
//line /snap/go/10455/src/context/context.go:322
		_go_fuzz_dep_.CoverTab[2333]++
								return stopped
//line /snap/go/10455/src/context/context.go:323
		// _ = "end of CoverTab[2333]"
	}
//line /snap/go/10455/src/context/context.go:324
	// _ = "end of CoverTab[2330]"
}

type afterFuncer interface {
	AfterFunc(func()) func() bool
}

type afterFuncCtx struct {
	cancelCtx
	once	sync.Once	// either starts running f or stops f from running
	f	func()
}

func (a *afterFuncCtx) cancel(removeFromParent bool, err, cause error) {
//line /snap/go/10455/src/context/context.go:337
	_go_fuzz_dep_.CoverTab[2337]++
							a.cancelCtx.cancel(false, err, cause)
							if removeFromParent {
//line /snap/go/10455/src/context/context.go:339
		_go_fuzz_dep_.CoverTab[526004]++
//line /snap/go/10455/src/context/context.go:339
		_go_fuzz_dep_.CoverTab[2339]++
								removeChild(a.Context, a)
//line /snap/go/10455/src/context/context.go:340
		// _ = "end of CoverTab[2339]"
	} else {
//line /snap/go/10455/src/context/context.go:341
		_go_fuzz_dep_.CoverTab[526005]++
//line /snap/go/10455/src/context/context.go:341
		_go_fuzz_dep_.CoverTab[2340]++
//line /snap/go/10455/src/context/context.go:341
		// _ = "end of CoverTab[2340]"
//line /snap/go/10455/src/context/context.go:341
	}
//line /snap/go/10455/src/context/context.go:341
	// _ = "end of CoverTab[2337]"
//line /snap/go/10455/src/context/context.go:341
	_go_fuzz_dep_.CoverTab[2338]++
							a.once.Do(func() {
//line /snap/go/10455/src/context/context.go:342
		_go_fuzz_dep_.CoverTab[2341]++
//line /snap/go/10455/src/context/context.go:342
		_curRoutineNum0_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/context/context.go:342
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum0_)
								go a.f()
//line /snap/go/10455/src/context/context.go:343
		// _ = "end of CoverTab[2341]"
	})
//line /snap/go/10455/src/context/context.go:344
	// _ = "end of CoverTab[2338]"
}

// A stopCtx is used as the parent context of a cancelCtx when
//line /snap/go/10455/src/context/context.go:347
// an AfterFunc has been registered with the parent.
//line /snap/go/10455/src/context/context.go:347
// It holds the stop function used to unregister the AfterFunc.
//line /snap/go/10455/src/context/context.go:350
type stopCtx struct {
	Context
	stop	func() bool
}

// goroutines counts the number of goroutines ever created; for testing.
var goroutines atomic.Int32

// &cancelCtxKey is the key that a cancelCtx returns itself for.
var cancelCtxKey int

// parentCancelCtx returns the underlying *cancelCtx for parent.
//line /snap/go/10455/src/context/context.go:361
// It does this by looking up parent.Value(&cancelCtxKey) to find
//line /snap/go/10455/src/context/context.go:361
// the innermost enclosing *cancelCtx and then checking whether
//line /snap/go/10455/src/context/context.go:361
// parent.Done() matches that *cancelCtx. (If not, the *cancelCtx
//line /snap/go/10455/src/context/context.go:361
// has been wrapped in a custom implementation providing a
//line /snap/go/10455/src/context/context.go:361
// different done channel, in which case we should not bypass it.)
//line /snap/go/10455/src/context/context.go:367
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
//line /snap/go/10455/src/context/context.go:367
	_go_fuzz_dep_.CoverTab[2342]++
							done := parent.Done()
							if done == closedchan || func() bool {
//line /snap/go/10455/src/context/context.go:369
		_go_fuzz_dep_.CoverTab[2346]++
//line /snap/go/10455/src/context/context.go:369
		return done == nil
//line /snap/go/10455/src/context/context.go:369
		// _ = "end of CoverTab[2346]"
//line /snap/go/10455/src/context/context.go:369
	}() {
//line /snap/go/10455/src/context/context.go:369
		_go_fuzz_dep_.CoverTab[526006]++
//line /snap/go/10455/src/context/context.go:369
		_go_fuzz_dep_.CoverTab[2347]++
								return nil, false
//line /snap/go/10455/src/context/context.go:370
		// _ = "end of CoverTab[2347]"
	} else {
//line /snap/go/10455/src/context/context.go:371
		_go_fuzz_dep_.CoverTab[526007]++
//line /snap/go/10455/src/context/context.go:371
		_go_fuzz_dep_.CoverTab[2348]++
//line /snap/go/10455/src/context/context.go:371
		// _ = "end of CoverTab[2348]"
//line /snap/go/10455/src/context/context.go:371
	}
//line /snap/go/10455/src/context/context.go:371
	// _ = "end of CoverTab[2342]"
//line /snap/go/10455/src/context/context.go:371
	_go_fuzz_dep_.CoverTab[2343]++
							p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
							if !ok {
//line /snap/go/10455/src/context/context.go:373
		_go_fuzz_dep_.CoverTab[526008]++
//line /snap/go/10455/src/context/context.go:373
		_go_fuzz_dep_.CoverTab[2349]++
								return nil, false
//line /snap/go/10455/src/context/context.go:374
		// _ = "end of CoverTab[2349]"
	} else {
//line /snap/go/10455/src/context/context.go:375
		_go_fuzz_dep_.CoverTab[526009]++
//line /snap/go/10455/src/context/context.go:375
		_go_fuzz_dep_.CoverTab[2350]++
//line /snap/go/10455/src/context/context.go:375
		// _ = "end of CoverTab[2350]"
//line /snap/go/10455/src/context/context.go:375
	}
//line /snap/go/10455/src/context/context.go:375
	// _ = "end of CoverTab[2343]"
//line /snap/go/10455/src/context/context.go:375
	_go_fuzz_dep_.CoverTab[2344]++
							pdone, _ := p.done.Load().(chan struct{})
							if pdone != done {
//line /snap/go/10455/src/context/context.go:377
		_go_fuzz_dep_.CoverTab[526010]++
//line /snap/go/10455/src/context/context.go:377
		_go_fuzz_dep_.CoverTab[2351]++
								return nil, false
//line /snap/go/10455/src/context/context.go:378
		// _ = "end of CoverTab[2351]"
	} else {
//line /snap/go/10455/src/context/context.go:379
		_go_fuzz_dep_.CoverTab[526011]++
//line /snap/go/10455/src/context/context.go:379
		_go_fuzz_dep_.CoverTab[2352]++
//line /snap/go/10455/src/context/context.go:379
		// _ = "end of CoverTab[2352]"
//line /snap/go/10455/src/context/context.go:379
	}
//line /snap/go/10455/src/context/context.go:379
	// _ = "end of CoverTab[2344]"
//line /snap/go/10455/src/context/context.go:379
	_go_fuzz_dep_.CoverTab[2345]++
							return p, true
//line /snap/go/10455/src/context/context.go:380
	// _ = "end of CoverTab[2345]"
}

// removeChild removes a context from its parent.
func removeChild(parent Context, child canceler) {
//line /snap/go/10455/src/context/context.go:384
	_go_fuzz_dep_.CoverTab[2353]++
							if s, ok := parent.(stopCtx); ok {
//line /snap/go/10455/src/context/context.go:385
		_go_fuzz_dep_.CoverTab[526012]++
//line /snap/go/10455/src/context/context.go:385
		_go_fuzz_dep_.CoverTab[2357]++
								s.stop()
								return
//line /snap/go/10455/src/context/context.go:387
		// _ = "end of CoverTab[2357]"
	} else {
//line /snap/go/10455/src/context/context.go:388
		_go_fuzz_dep_.CoverTab[526013]++
//line /snap/go/10455/src/context/context.go:388
		_go_fuzz_dep_.CoverTab[2358]++
//line /snap/go/10455/src/context/context.go:388
		// _ = "end of CoverTab[2358]"
//line /snap/go/10455/src/context/context.go:388
	}
//line /snap/go/10455/src/context/context.go:388
	// _ = "end of CoverTab[2353]"
//line /snap/go/10455/src/context/context.go:388
	_go_fuzz_dep_.CoverTab[2354]++
							p, ok := parentCancelCtx(parent)
							if !ok {
//line /snap/go/10455/src/context/context.go:390
		_go_fuzz_dep_.CoverTab[526014]++
//line /snap/go/10455/src/context/context.go:390
		_go_fuzz_dep_.CoverTab[2359]++
								return
//line /snap/go/10455/src/context/context.go:391
		// _ = "end of CoverTab[2359]"
	} else {
//line /snap/go/10455/src/context/context.go:392
		_go_fuzz_dep_.CoverTab[526015]++
//line /snap/go/10455/src/context/context.go:392
		_go_fuzz_dep_.CoverTab[2360]++
//line /snap/go/10455/src/context/context.go:392
		// _ = "end of CoverTab[2360]"
//line /snap/go/10455/src/context/context.go:392
	}
//line /snap/go/10455/src/context/context.go:392
	// _ = "end of CoverTab[2354]"
//line /snap/go/10455/src/context/context.go:392
	_go_fuzz_dep_.CoverTab[2355]++
							p.mu.Lock()
							if p.children != nil {
//line /snap/go/10455/src/context/context.go:394
		_go_fuzz_dep_.CoverTab[526016]++
//line /snap/go/10455/src/context/context.go:394
		_go_fuzz_dep_.CoverTab[2361]++
								delete(p.children, child)
//line /snap/go/10455/src/context/context.go:395
		// _ = "end of CoverTab[2361]"
	} else {
//line /snap/go/10455/src/context/context.go:396
		_go_fuzz_dep_.CoverTab[526017]++
//line /snap/go/10455/src/context/context.go:396
		_go_fuzz_dep_.CoverTab[2362]++
//line /snap/go/10455/src/context/context.go:396
		// _ = "end of CoverTab[2362]"
//line /snap/go/10455/src/context/context.go:396
	}
//line /snap/go/10455/src/context/context.go:396
	// _ = "end of CoverTab[2355]"
//line /snap/go/10455/src/context/context.go:396
	_go_fuzz_dep_.CoverTab[2356]++
							p.mu.Unlock()
//line /snap/go/10455/src/context/context.go:397
	// _ = "end of CoverTab[2356]"
}

// A canceler is a context type that can be canceled directly. The
//line /snap/go/10455/src/context/context.go:400
// implementations are *cancelCtx and *timerCtx.
//line /snap/go/10455/src/context/context.go:402
type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

// A cancelCtx can be canceled. When canceled, it also cancels any children
//line /snap/go/10455/src/context/context.go:414
// that implement canceler.
//line /snap/go/10455/src/context/context.go:416
type cancelCtx struct {
	Context

	mu		sync.Mutex		// protects following fields
	done		atomic.Value		// of chan struct{}, created lazily, closed by first cancel call
	children	map[canceler]struct{}	// set to nil by the first cancel call
	err		error			// set to non-nil by the first cancel call
	cause		error			// set to non-nil by the first cancel call
}

func (c *cancelCtx) Value(key any) any {
//line /snap/go/10455/src/context/context.go:426
	_go_fuzz_dep_.CoverTab[2363]++
							if key == &cancelCtxKey {
//line /snap/go/10455/src/context/context.go:427
		_go_fuzz_dep_.CoverTab[526018]++
//line /snap/go/10455/src/context/context.go:427
		_go_fuzz_dep_.CoverTab[2365]++
								return c
//line /snap/go/10455/src/context/context.go:428
		// _ = "end of CoverTab[2365]"
	} else {
//line /snap/go/10455/src/context/context.go:429
		_go_fuzz_dep_.CoverTab[526019]++
//line /snap/go/10455/src/context/context.go:429
		_go_fuzz_dep_.CoverTab[2366]++
//line /snap/go/10455/src/context/context.go:429
		// _ = "end of CoverTab[2366]"
//line /snap/go/10455/src/context/context.go:429
	}
//line /snap/go/10455/src/context/context.go:429
	// _ = "end of CoverTab[2363]"
//line /snap/go/10455/src/context/context.go:429
	_go_fuzz_dep_.CoverTab[2364]++
							return value(c.Context, key)
//line /snap/go/10455/src/context/context.go:430
	// _ = "end of CoverTab[2364]"
}

func (c *cancelCtx) Done() <-chan struct{} {
//line /snap/go/10455/src/context/context.go:433
	_go_fuzz_dep_.CoverTab[2367]++
							d := c.done.Load()
							if d != nil {
//line /snap/go/10455/src/context/context.go:435
		_go_fuzz_dep_.CoverTab[526020]++
//line /snap/go/10455/src/context/context.go:435
		_go_fuzz_dep_.CoverTab[2370]++
								return d.(chan struct{})
//line /snap/go/10455/src/context/context.go:436
		// _ = "end of CoverTab[2370]"
	} else {
//line /snap/go/10455/src/context/context.go:437
		_go_fuzz_dep_.CoverTab[526021]++
//line /snap/go/10455/src/context/context.go:437
		_go_fuzz_dep_.CoverTab[2371]++
//line /snap/go/10455/src/context/context.go:437
		// _ = "end of CoverTab[2371]"
//line /snap/go/10455/src/context/context.go:437
	}
//line /snap/go/10455/src/context/context.go:437
	// _ = "end of CoverTab[2367]"
//line /snap/go/10455/src/context/context.go:437
	_go_fuzz_dep_.CoverTab[2368]++
							c.mu.Lock()
							defer c.mu.Unlock()
							d = c.done.Load()
							if d == nil {
//line /snap/go/10455/src/context/context.go:441
		_go_fuzz_dep_.CoverTab[526022]++
//line /snap/go/10455/src/context/context.go:441
		_go_fuzz_dep_.CoverTab[2372]++
								d = make(chan struct{})
								c.done.Store(d)
//line /snap/go/10455/src/context/context.go:443
		// _ = "end of CoverTab[2372]"
	} else {
//line /snap/go/10455/src/context/context.go:444
		_go_fuzz_dep_.CoverTab[526023]++
//line /snap/go/10455/src/context/context.go:444
		_go_fuzz_dep_.CoverTab[2373]++
//line /snap/go/10455/src/context/context.go:444
		// _ = "end of CoverTab[2373]"
//line /snap/go/10455/src/context/context.go:444
	}
//line /snap/go/10455/src/context/context.go:444
	// _ = "end of CoverTab[2368]"
//line /snap/go/10455/src/context/context.go:444
	_go_fuzz_dep_.CoverTab[2369]++
							return d.(chan struct{})
//line /snap/go/10455/src/context/context.go:445
	// _ = "end of CoverTab[2369]"
}

func (c *cancelCtx) Err() error {
//line /snap/go/10455/src/context/context.go:448
	_go_fuzz_dep_.CoverTab[2374]++
							c.mu.Lock()
							err := c.err
							c.mu.Unlock()
							return err
//line /snap/go/10455/src/context/context.go:452
	// _ = "end of CoverTab[2374]"
}

// propagateCancel arranges for child to be canceled when parent is.
//line /snap/go/10455/src/context/context.go:455
// It sets the parent context of cancelCtx.
//line /snap/go/10455/src/context/context.go:457
func (c *cancelCtx) propagateCancel(parent Context, child canceler) {
//line /snap/go/10455/src/context/context.go:457
	_go_fuzz_dep_.CoverTab[2375]++
							c.Context = parent

							done := parent.Done()
							if done == nil {
//line /snap/go/10455/src/context/context.go:461
		_go_fuzz_dep_.CoverTab[526024]++
//line /snap/go/10455/src/context/context.go:461
		_go_fuzz_dep_.CoverTab[2380]++
								return
//line /snap/go/10455/src/context/context.go:462
		// _ = "end of CoverTab[2380]"
	} else {
//line /snap/go/10455/src/context/context.go:463
		_go_fuzz_dep_.CoverTab[526025]++
//line /snap/go/10455/src/context/context.go:463
		_go_fuzz_dep_.CoverTab[2381]++
//line /snap/go/10455/src/context/context.go:463
		// _ = "end of CoverTab[2381]"
//line /snap/go/10455/src/context/context.go:463
	}
//line /snap/go/10455/src/context/context.go:463
	// _ = "end of CoverTab[2375]"
//line /snap/go/10455/src/context/context.go:463
	_go_fuzz_dep_.CoverTab[2376]++

							select {
	case <-done:
//line /snap/go/10455/src/context/context.go:466
		_go_fuzz_dep_.CoverTab[2382]++

								child.cancel(false, parent.Err(), Cause(parent))
								return
//line /snap/go/10455/src/context/context.go:469
		// _ = "end of CoverTab[2382]"
	default:
//line /snap/go/10455/src/context/context.go:470
		_go_fuzz_dep_.CoverTab[2383]++
//line /snap/go/10455/src/context/context.go:470
		// _ = "end of CoverTab[2383]"
	}
//line /snap/go/10455/src/context/context.go:471
	// _ = "end of CoverTab[2376]"
//line /snap/go/10455/src/context/context.go:471
	_go_fuzz_dep_.CoverTab[2377]++

							if p, ok := parentCancelCtx(parent); ok {
//line /snap/go/10455/src/context/context.go:473
		_go_fuzz_dep_.CoverTab[526026]++
//line /snap/go/10455/src/context/context.go:473
		_go_fuzz_dep_.CoverTab[2384]++

								p.mu.Lock()
								if p.err != nil {
//line /snap/go/10455/src/context/context.go:476
			_go_fuzz_dep_.CoverTab[526028]++
//line /snap/go/10455/src/context/context.go:476
			_go_fuzz_dep_.CoverTab[2386]++

									child.cancel(false, p.err, p.cause)
//line /snap/go/10455/src/context/context.go:478
			// _ = "end of CoverTab[2386]"
		} else {
//line /snap/go/10455/src/context/context.go:479
			_go_fuzz_dep_.CoverTab[526029]++
//line /snap/go/10455/src/context/context.go:479
			_go_fuzz_dep_.CoverTab[2387]++
									if p.children == nil {
//line /snap/go/10455/src/context/context.go:480
				_go_fuzz_dep_.CoverTab[526030]++
//line /snap/go/10455/src/context/context.go:480
				_go_fuzz_dep_.CoverTab[2389]++
										p.children = make(map[canceler]struct{})
//line /snap/go/10455/src/context/context.go:481
				// _ = "end of CoverTab[2389]"
			} else {
//line /snap/go/10455/src/context/context.go:482
				_go_fuzz_dep_.CoverTab[526031]++
//line /snap/go/10455/src/context/context.go:482
				_go_fuzz_dep_.CoverTab[2390]++
//line /snap/go/10455/src/context/context.go:482
				// _ = "end of CoverTab[2390]"
//line /snap/go/10455/src/context/context.go:482
			}
//line /snap/go/10455/src/context/context.go:482
			// _ = "end of CoverTab[2387]"
//line /snap/go/10455/src/context/context.go:482
			_go_fuzz_dep_.CoverTab[2388]++
									p.children[child] = struct{}{}
//line /snap/go/10455/src/context/context.go:483
			// _ = "end of CoverTab[2388]"
		}
//line /snap/go/10455/src/context/context.go:484
		// _ = "end of CoverTab[2384]"
//line /snap/go/10455/src/context/context.go:484
		_go_fuzz_dep_.CoverTab[2385]++
								p.mu.Unlock()
								return
//line /snap/go/10455/src/context/context.go:486
		// _ = "end of CoverTab[2385]"
	} else {
//line /snap/go/10455/src/context/context.go:487
		_go_fuzz_dep_.CoverTab[526027]++
//line /snap/go/10455/src/context/context.go:487
		_go_fuzz_dep_.CoverTab[2391]++
//line /snap/go/10455/src/context/context.go:487
		// _ = "end of CoverTab[2391]"
//line /snap/go/10455/src/context/context.go:487
	}
//line /snap/go/10455/src/context/context.go:487
	// _ = "end of CoverTab[2377]"
//line /snap/go/10455/src/context/context.go:487
	_go_fuzz_dep_.CoverTab[2378]++

							if a, ok := parent.(afterFuncer); ok {
//line /snap/go/10455/src/context/context.go:489
		_go_fuzz_dep_.CoverTab[526032]++
//line /snap/go/10455/src/context/context.go:489
		_go_fuzz_dep_.CoverTab[2392]++

								c.mu.Lock()
								stop := a.AfterFunc(func() {
//line /snap/go/10455/src/context/context.go:492
			_go_fuzz_dep_.CoverTab[2394]++
									child.cancel(false, parent.Err(), Cause(parent))
//line /snap/go/10455/src/context/context.go:493
			// _ = "end of CoverTab[2394]"
		})
//line /snap/go/10455/src/context/context.go:494
		// _ = "end of CoverTab[2392]"
//line /snap/go/10455/src/context/context.go:494
		_go_fuzz_dep_.CoverTab[2393]++
								c.Context = stopCtx{
			Context:	parent,
			stop:		stop,
		}
								c.mu.Unlock()
								return
//line /snap/go/10455/src/context/context.go:500
		// _ = "end of CoverTab[2393]"
	} else {
//line /snap/go/10455/src/context/context.go:501
		_go_fuzz_dep_.CoverTab[526033]++
//line /snap/go/10455/src/context/context.go:501
		_go_fuzz_dep_.CoverTab[2395]++
//line /snap/go/10455/src/context/context.go:501
		// _ = "end of CoverTab[2395]"
//line /snap/go/10455/src/context/context.go:501
	}
//line /snap/go/10455/src/context/context.go:501
	// _ = "end of CoverTab[2378]"
//line /snap/go/10455/src/context/context.go:501
	_go_fuzz_dep_.CoverTab[2379]++

							goroutines.Add(1)
//line /snap/go/10455/src/context/context.go:503
	_curRoutineNum1_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/context/context.go:503
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum1_)
							go func() {
//line /snap/go/10455/src/context/context.go:504
		_go_fuzz_dep_.CoverTab[2396]++
//line /snap/go/10455/src/context/context.go:504
		defer func() {
//line /snap/go/10455/src/context/context.go:504
			_go_fuzz_dep_.CoverTab[2397]++
//line /snap/go/10455/src/context/context.go:504
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum1_)
//line /snap/go/10455/src/context/context.go:504
			// _ = "end of CoverTab[2397]"
//line /snap/go/10455/src/context/context.go:504
		}()
								select {
		case <-parent.Done():
//line /snap/go/10455/src/context/context.go:506
			_go_fuzz_dep_.CoverTab[2398]++
									child.cancel(false, parent.Err(), Cause(parent))
//line /snap/go/10455/src/context/context.go:507
			// _ = "end of CoverTab[2398]"
		case <-child.Done():
//line /snap/go/10455/src/context/context.go:508
			_go_fuzz_dep_.CoverTab[2399]++
//line /snap/go/10455/src/context/context.go:508
			// _ = "end of CoverTab[2399]"
		}
//line /snap/go/10455/src/context/context.go:509
		// _ = "end of CoverTab[2396]"
	}()
//line /snap/go/10455/src/context/context.go:510
	// _ = "end of CoverTab[2379]"
}

type stringer interface {
	String() string
}

func contextName(c Context) string {
//line /snap/go/10455/src/context/context.go:517
	_go_fuzz_dep_.CoverTab[2400]++
							if s, ok := c.(stringer); ok {
//line /snap/go/10455/src/context/context.go:518
		_go_fuzz_dep_.CoverTab[526034]++
//line /snap/go/10455/src/context/context.go:518
		_go_fuzz_dep_.CoverTab[2402]++
								return s.String()
//line /snap/go/10455/src/context/context.go:519
		// _ = "end of CoverTab[2402]"
	} else {
//line /snap/go/10455/src/context/context.go:520
		_go_fuzz_dep_.CoverTab[526035]++
//line /snap/go/10455/src/context/context.go:520
		_go_fuzz_dep_.CoverTab[2403]++
//line /snap/go/10455/src/context/context.go:520
		// _ = "end of CoverTab[2403]"
//line /snap/go/10455/src/context/context.go:520
	}
//line /snap/go/10455/src/context/context.go:520
	// _ = "end of CoverTab[2400]"
//line /snap/go/10455/src/context/context.go:520
	_go_fuzz_dep_.CoverTab[2401]++
							return reflectlite.TypeOf(c).String()
//line /snap/go/10455/src/context/context.go:521
	// _ = "end of CoverTab[2401]"
}

func (c *cancelCtx) String() string {
//line /snap/go/10455/src/context/context.go:524
	_go_fuzz_dep_.CoverTab[2404]++
							return contextName(c.Context) + ".WithCancel"
//line /snap/go/10455/src/context/context.go:525
	// _ = "end of CoverTab[2404]"
}

// cancel closes c.done, cancels each of c's children, and, if
//line /snap/go/10455/src/context/context.go:528
// removeFromParent is true, removes c from its parent's children.
//line /snap/go/10455/src/context/context.go:528
// cancel sets c.cause to cause if this is the first time c is canceled.
//line /snap/go/10455/src/context/context.go:531
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
//line /snap/go/10455/src/context/context.go:531
	_go_fuzz_dep_.CoverTab[2405]++
							if err == nil {
//line /snap/go/10455/src/context/context.go:532
		_go_fuzz_dep_.CoverTab[526036]++
//line /snap/go/10455/src/context/context.go:532
		_go_fuzz_dep_.CoverTab[2411]++
								panic("context: internal error: missing cancel error")
//line /snap/go/10455/src/context/context.go:533
		// _ = "end of CoverTab[2411]"
	} else {
//line /snap/go/10455/src/context/context.go:534
		_go_fuzz_dep_.CoverTab[526037]++
//line /snap/go/10455/src/context/context.go:534
		_go_fuzz_dep_.CoverTab[2412]++
//line /snap/go/10455/src/context/context.go:534
		// _ = "end of CoverTab[2412]"
//line /snap/go/10455/src/context/context.go:534
	}
//line /snap/go/10455/src/context/context.go:534
	// _ = "end of CoverTab[2405]"
//line /snap/go/10455/src/context/context.go:534
	_go_fuzz_dep_.CoverTab[2406]++
							if cause == nil {
//line /snap/go/10455/src/context/context.go:535
		_go_fuzz_dep_.CoverTab[526038]++
//line /snap/go/10455/src/context/context.go:535
		_go_fuzz_dep_.CoverTab[2413]++
								cause = err
//line /snap/go/10455/src/context/context.go:536
		// _ = "end of CoverTab[2413]"
	} else {
//line /snap/go/10455/src/context/context.go:537
		_go_fuzz_dep_.CoverTab[526039]++
//line /snap/go/10455/src/context/context.go:537
		_go_fuzz_dep_.CoverTab[2414]++
//line /snap/go/10455/src/context/context.go:537
		// _ = "end of CoverTab[2414]"
//line /snap/go/10455/src/context/context.go:537
	}
//line /snap/go/10455/src/context/context.go:537
	// _ = "end of CoverTab[2406]"
//line /snap/go/10455/src/context/context.go:537
	_go_fuzz_dep_.CoverTab[2407]++
							c.mu.Lock()
							if c.err != nil {
//line /snap/go/10455/src/context/context.go:539
		_go_fuzz_dep_.CoverTab[526040]++
//line /snap/go/10455/src/context/context.go:539
		_go_fuzz_dep_.CoverTab[2415]++
								c.mu.Unlock()
								return
//line /snap/go/10455/src/context/context.go:541
		// _ = "end of CoverTab[2415]"
	} else {
//line /snap/go/10455/src/context/context.go:542
		_go_fuzz_dep_.CoverTab[526041]++
//line /snap/go/10455/src/context/context.go:542
		_go_fuzz_dep_.CoverTab[2416]++
//line /snap/go/10455/src/context/context.go:542
		// _ = "end of CoverTab[2416]"
//line /snap/go/10455/src/context/context.go:542
	}
//line /snap/go/10455/src/context/context.go:542
	// _ = "end of CoverTab[2407]"
//line /snap/go/10455/src/context/context.go:542
	_go_fuzz_dep_.CoverTab[2408]++
							c.err = err
							c.cause = cause
							d, _ := c.done.Load().(chan struct{})
							if d == nil {
//line /snap/go/10455/src/context/context.go:546
		_go_fuzz_dep_.CoverTab[526042]++
//line /snap/go/10455/src/context/context.go:546
		_go_fuzz_dep_.CoverTab[2417]++
								c.done.Store(closedchan)
//line /snap/go/10455/src/context/context.go:547
		// _ = "end of CoverTab[2417]"
	} else {
//line /snap/go/10455/src/context/context.go:548
		_go_fuzz_dep_.CoverTab[526043]++
//line /snap/go/10455/src/context/context.go:548
		_go_fuzz_dep_.CoverTab[2418]++
								close(d)
//line /snap/go/10455/src/context/context.go:549
		// _ = "end of CoverTab[2418]"
	}
//line /snap/go/10455/src/context/context.go:550
	// _ = "end of CoverTab[2408]"
//line /snap/go/10455/src/context/context.go:550
	_go_fuzz_dep_.CoverTab[2409]++
//line /snap/go/10455/src/context/context.go:550
	_go_fuzz_dep_.CoverTab[786572] = 0
							for child := range c.children {
//line /snap/go/10455/src/context/context.go:551
		if _go_fuzz_dep_.CoverTab[786572] == 0 {
//line /snap/go/10455/src/context/context.go:551
			_go_fuzz_dep_.CoverTab[526084]++
//line /snap/go/10455/src/context/context.go:551
		} else {
//line /snap/go/10455/src/context/context.go:551
			_go_fuzz_dep_.CoverTab[526085]++
//line /snap/go/10455/src/context/context.go:551
		}
//line /snap/go/10455/src/context/context.go:551
		_go_fuzz_dep_.CoverTab[786572] = 1
//line /snap/go/10455/src/context/context.go:551
		_go_fuzz_dep_.CoverTab[2419]++

								child.cancel(false, err, cause)
//line /snap/go/10455/src/context/context.go:553
		// _ = "end of CoverTab[2419]"
	}
//line /snap/go/10455/src/context/context.go:554
	if _go_fuzz_dep_.CoverTab[786572] == 0 {
//line /snap/go/10455/src/context/context.go:554
		_go_fuzz_dep_.CoverTab[526086]++
//line /snap/go/10455/src/context/context.go:554
	} else {
//line /snap/go/10455/src/context/context.go:554
		_go_fuzz_dep_.CoverTab[526087]++
//line /snap/go/10455/src/context/context.go:554
	}
//line /snap/go/10455/src/context/context.go:554
	// _ = "end of CoverTab[2409]"
//line /snap/go/10455/src/context/context.go:554
	_go_fuzz_dep_.CoverTab[2410]++
							c.children = nil
							c.mu.Unlock()

							if removeFromParent {
//line /snap/go/10455/src/context/context.go:558
		_go_fuzz_dep_.CoverTab[526044]++
//line /snap/go/10455/src/context/context.go:558
		_go_fuzz_dep_.CoverTab[2420]++
								removeChild(c.Context, c)
//line /snap/go/10455/src/context/context.go:559
		// _ = "end of CoverTab[2420]"
	} else {
//line /snap/go/10455/src/context/context.go:560
		_go_fuzz_dep_.CoverTab[526045]++
//line /snap/go/10455/src/context/context.go:560
		_go_fuzz_dep_.CoverTab[2421]++
//line /snap/go/10455/src/context/context.go:560
		// _ = "end of CoverTab[2421]"
//line /snap/go/10455/src/context/context.go:560
	}
//line /snap/go/10455/src/context/context.go:560
	// _ = "end of CoverTab[2410]"
}

// WithoutCancel returns a copy of parent that is not canceled when parent is canceled.
//line /snap/go/10455/src/context/context.go:563
// The returned context returns no Deadline or Err, and its Done channel is nil.
//line /snap/go/10455/src/context/context.go:563
// Calling [Cause] on the returned context returns nil.
//line /snap/go/10455/src/context/context.go:566
func WithoutCancel(parent Context) Context {
//line /snap/go/10455/src/context/context.go:566
	_go_fuzz_dep_.CoverTab[2422]++
							if parent == nil {
//line /snap/go/10455/src/context/context.go:567
		_go_fuzz_dep_.CoverTab[526046]++
//line /snap/go/10455/src/context/context.go:567
		_go_fuzz_dep_.CoverTab[2424]++
								panic("cannot create context from nil parent")
//line /snap/go/10455/src/context/context.go:568
		// _ = "end of CoverTab[2424]"
	} else {
//line /snap/go/10455/src/context/context.go:569
		_go_fuzz_dep_.CoverTab[526047]++
//line /snap/go/10455/src/context/context.go:569
		_go_fuzz_dep_.CoverTab[2425]++
//line /snap/go/10455/src/context/context.go:569
		// _ = "end of CoverTab[2425]"
//line /snap/go/10455/src/context/context.go:569
	}
//line /snap/go/10455/src/context/context.go:569
	// _ = "end of CoverTab[2422]"
//line /snap/go/10455/src/context/context.go:569
	_go_fuzz_dep_.CoverTab[2423]++
							return withoutCancelCtx{parent}
//line /snap/go/10455/src/context/context.go:570
	// _ = "end of CoverTab[2423]"
}

type withoutCancelCtx struct {
	c Context
}

func (withoutCancelCtx) Deadline() (deadline time.Time, ok bool) {
//line /snap/go/10455/src/context/context.go:577
	_go_fuzz_dep_.CoverTab[2426]++
							return
//line /snap/go/10455/src/context/context.go:578
	// _ = "end of CoverTab[2426]"
}

func (withoutCancelCtx) Done() <-chan struct{} {
//line /snap/go/10455/src/context/context.go:581
	_go_fuzz_dep_.CoverTab[2427]++
							return nil
//line /snap/go/10455/src/context/context.go:582
	// _ = "end of CoverTab[2427]"
}

func (withoutCancelCtx) Err() error {
//line /snap/go/10455/src/context/context.go:585
	_go_fuzz_dep_.CoverTab[2428]++
							return nil
//line /snap/go/10455/src/context/context.go:586
	// _ = "end of CoverTab[2428]"
}

func (c withoutCancelCtx) Value(key any) any {
//line /snap/go/10455/src/context/context.go:589
	_go_fuzz_dep_.CoverTab[2429]++
							return value(c, key)
//line /snap/go/10455/src/context/context.go:590
	// _ = "end of CoverTab[2429]"
}

func (c withoutCancelCtx) String() string {
//line /snap/go/10455/src/context/context.go:593
	_go_fuzz_dep_.CoverTab[2430]++
							return contextName(c.c) + ".WithoutCancel"
//line /snap/go/10455/src/context/context.go:594
	// _ = "end of CoverTab[2430]"
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
//line /snap/go/10455/src/context/context.go:597
// to be no later than d. If the parent's deadline is already earlier than d,
//line /snap/go/10455/src/context/context.go:597
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
//line /snap/go/10455/src/context/context.go:597
// [Context.Done] channel is closed when the deadline expires, when the returned
//line /snap/go/10455/src/context/context.go:597
// cancel function is called, or when the parent context's Done channel is
//line /snap/go/10455/src/context/context.go:597
// closed, whichever happens first.
//line /snap/go/10455/src/context/context.go:597
//
//line /snap/go/10455/src/context/context.go:597
// Canceling this context releases resources associated with it, so code should
//line /snap/go/10455/src/context/context.go:597
// call cancel as soon as the operations running in this [Context] complete.
//line /snap/go/10455/src/context/context.go:606
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
//line /snap/go/10455/src/context/context.go:606
	_go_fuzz_dep_.CoverTab[2431]++
							return WithDeadlineCause(parent, d, nil)
//line /snap/go/10455/src/context/context.go:607
	// _ = "end of CoverTab[2431]"
}

// WithDeadlineCause behaves like [WithDeadline] but also sets the cause of the
//line /snap/go/10455/src/context/context.go:610
// returned Context when the deadline is exceeded. The returned [CancelFunc] does
//line /snap/go/10455/src/context/context.go:610
// not set the cause.
//line /snap/go/10455/src/context/context.go:613
func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc) {
//line /snap/go/10455/src/context/context.go:613
	_go_fuzz_dep_.CoverTab[2432]++
							if parent == nil {
//line /snap/go/10455/src/context/context.go:614
		_go_fuzz_dep_.CoverTab[526048]++
//line /snap/go/10455/src/context/context.go:614
		_go_fuzz_dep_.CoverTab[2437]++
								panic("cannot create context from nil parent")
//line /snap/go/10455/src/context/context.go:615
		// _ = "end of CoverTab[2437]"
	} else {
//line /snap/go/10455/src/context/context.go:616
		_go_fuzz_dep_.CoverTab[526049]++
//line /snap/go/10455/src/context/context.go:616
		_go_fuzz_dep_.CoverTab[2438]++
//line /snap/go/10455/src/context/context.go:616
		// _ = "end of CoverTab[2438]"
//line /snap/go/10455/src/context/context.go:616
	}
//line /snap/go/10455/src/context/context.go:616
	// _ = "end of CoverTab[2432]"
//line /snap/go/10455/src/context/context.go:616
	_go_fuzz_dep_.CoverTab[2433]++
							if cur, ok := parent.Deadline(); ok && func() bool {
//line /snap/go/10455/src/context/context.go:617
		_go_fuzz_dep_.CoverTab[2439]++
//line /snap/go/10455/src/context/context.go:617
		return cur.Before(d)
//line /snap/go/10455/src/context/context.go:617
		// _ = "end of CoverTab[2439]"
//line /snap/go/10455/src/context/context.go:617
	}() {
//line /snap/go/10455/src/context/context.go:617
		_go_fuzz_dep_.CoverTab[526050]++
//line /snap/go/10455/src/context/context.go:617
		_go_fuzz_dep_.CoverTab[2440]++

								return WithCancel(parent)
//line /snap/go/10455/src/context/context.go:619
		// _ = "end of CoverTab[2440]"
	} else {
//line /snap/go/10455/src/context/context.go:620
		_go_fuzz_dep_.CoverTab[526051]++
//line /snap/go/10455/src/context/context.go:620
		_go_fuzz_dep_.CoverTab[2441]++
//line /snap/go/10455/src/context/context.go:620
		// _ = "end of CoverTab[2441]"
//line /snap/go/10455/src/context/context.go:620
	}
//line /snap/go/10455/src/context/context.go:620
	// _ = "end of CoverTab[2433]"
//line /snap/go/10455/src/context/context.go:620
	_go_fuzz_dep_.CoverTab[2434]++
							c := &timerCtx{
		deadline: d,
	}
	c.cancelCtx.propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
//line /snap/go/10455/src/context/context.go:626
		_go_fuzz_dep_.CoverTab[526052]++
//line /snap/go/10455/src/context/context.go:626
		_go_fuzz_dep_.CoverTab[2442]++
								c.cancel(true, DeadlineExceeded, cause)
								return c, func() { _go_fuzz_dep_.CoverTab[2443]++; c.cancel(false, Canceled, nil); // _ = "end of CoverTab[2443]" }
//line /snap/go/10455/src/context/context.go:628
		// _ = "end of CoverTab[2442]"
	} else {
//line /snap/go/10455/src/context/context.go:629
		_go_fuzz_dep_.CoverTab[526053]++
//line /snap/go/10455/src/context/context.go:629
		_go_fuzz_dep_.CoverTab[2444]++
//line /snap/go/10455/src/context/context.go:629
		// _ = "end of CoverTab[2444]"
//line /snap/go/10455/src/context/context.go:629
	}
//line /snap/go/10455/src/context/context.go:629
	// _ = "end of CoverTab[2434]"
//line /snap/go/10455/src/context/context.go:629
	_go_fuzz_dep_.CoverTab[2435]++
							c.mu.Lock()
							defer c.mu.Unlock()
							if c.err == nil {
//line /snap/go/10455/src/context/context.go:632
		_go_fuzz_dep_.CoverTab[526054]++
//line /snap/go/10455/src/context/context.go:632
		_go_fuzz_dep_.CoverTab[2445]++
								c.timer = time.AfterFunc(dur, func() {
//line /snap/go/10455/src/context/context.go:633
			_go_fuzz_dep_.CoverTab[2446]++
									c.cancel(true, DeadlineExceeded, cause)
//line /snap/go/10455/src/context/context.go:634
			// _ = "end of CoverTab[2446]"
		})
//line /snap/go/10455/src/context/context.go:635
		// _ = "end of CoverTab[2445]"
	} else {
//line /snap/go/10455/src/context/context.go:636
		_go_fuzz_dep_.CoverTab[526055]++
//line /snap/go/10455/src/context/context.go:636
		_go_fuzz_dep_.CoverTab[2447]++
//line /snap/go/10455/src/context/context.go:636
		// _ = "end of CoverTab[2447]"
//line /snap/go/10455/src/context/context.go:636
	}
//line /snap/go/10455/src/context/context.go:636
	// _ = "end of CoverTab[2435]"
//line /snap/go/10455/src/context/context.go:636
	_go_fuzz_dep_.CoverTab[2436]++
							return c, func() { _go_fuzz_dep_.CoverTab[2448]++; c.cancel(true, Canceled, nil); // _ = "end of CoverTab[2448]" }
//line /snap/go/10455/src/context/context.go:637
	// _ = "end of CoverTab[2436]"
}

// A timerCtx carries a timer and a deadline. It embeds a cancelCtx to
//line /snap/go/10455/src/context/context.go:640
// implement Done and Err. It implements cancel by stopping its timer then
//line /snap/go/10455/src/context/context.go:640
// delegating to cancelCtx.cancel.
//line /snap/go/10455/src/context/context.go:643
type timerCtx struct {
	cancelCtx
	timer	*time.Timer	// Under cancelCtx.mu.

	deadline	time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
//line /snap/go/10455/src/context/context.go:650
	_go_fuzz_dep_.CoverTab[2449]++
							return c.deadline, true
//line /snap/go/10455/src/context/context.go:651
	// _ = "end of CoverTab[2449]"
}

func (c *timerCtx) String() string {
//line /snap/go/10455/src/context/context.go:654
	_go_fuzz_dep_.CoverTab[2450]++
							return contextName(c.cancelCtx.Context) + ".WithDeadline(" +
		c.deadline.String() + " [" +
		time.Until(c.deadline).String() + "])"
//line /snap/go/10455/src/context/context.go:657
	// _ = "end of CoverTab[2450]"
}

func (c *timerCtx) cancel(removeFromParent bool, err, cause error) {
//line /snap/go/10455/src/context/context.go:660
	_go_fuzz_dep_.CoverTab[2451]++
							c.cancelCtx.cancel(false, err, cause)
							if removeFromParent {
//line /snap/go/10455/src/context/context.go:662
		_go_fuzz_dep_.CoverTab[526056]++
//line /snap/go/10455/src/context/context.go:662
		_go_fuzz_dep_.CoverTab[2454]++

								removeChild(c.cancelCtx.Context, c)
//line /snap/go/10455/src/context/context.go:664
		// _ = "end of CoverTab[2454]"
	} else {
//line /snap/go/10455/src/context/context.go:665
		_go_fuzz_dep_.CoverTab[526057]++
//line /snap/go/10455/src/context/context.go:665
		_go_fuzz_dep_.CoverTab[2455]++
//line /snap/go/10455/src/context/context.go:665
		// _ = "end of CoverTab[2455]"
//line /snap/go/10455/src/context/context.go:665
	}
//line /snap/go/10455/src/context/context.go:665
	// _ = "end of CoverTab[2451]"
//line /snap/go/10455/src/context/context.go:665
	_go_fuzz_dep_.CoverTab[2452]++
							c.mu.Lock()
							if c.timer != nil {
//line /snap/go/10455/src/context/context.go:667
		_go_fuzz_dep_.CoverTab[526058]++
//line /snap/go/10455/src/context/context.go:667
		_go_fuzz_dep_.CoverTab[2456]++
								c.timer.Stop()
								c.timer = nil
//line /snap/go/10455/src/context/context.go:669
		// _ = "end of CoverTab[2456]"
	} else {
//line /snap/go/10455/src/context/context.go:670
		_go_fuzz_dep_.CoverTab[526059]++
//line /snap/go/10455/src/context/context.go:670
		_go_fuzz_dep_.CoverTab[2457]++
//line /snap/go/10455/src/context/context.go:670
		// _ = "end of CoverTab[2457]"
//line /snap/go/10455/src/context/context.go:670
	}
//line /snap/go/10455/src/context/context.go:670
	// _ = "end of CoverTab[2452]"
//line /snap/go/10455/src/context/context.go:670
	_go_fuzz_dep_.CoverTab[2453]++
							c.mu.Unlock()
//line /snap/go/10455/src/context/context.go:671
	// _ = "end of CoverTab[2453]"
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//line /snap/go/10455/src/context/context.go:674
//
//line /snap/go/10455/src/context/context.go:674
// Canceling this context releases resources associated with it, so code should
//line /snap/go/10455/src/context/context.go:674
// call cancel as soon as the operations running in this [Context] complete:
//line /snap/go/10455/src/context/context.go:674
//
//line /snap/go/10455/src/context/context.go:674
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//line /snap/go/10455/src/context/context.go:674
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//line /snap/go/10455/src/context/context.go:674
//		defer cancel()  // releases resources if slowOperation completes before timeout elapses
//line /snap/go/10455/src/context/context.go:674
//		return slowOperation(ctx)
//line /snap/go/10455/src/context/context.go:674
//	}
//line /snap/go/10455/src/context/context.go:684
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
//line /snap/go/10455/src/context/context.go:684
	_go_fuzz_dep_.CoverTab[2458]++
							return WithDeadline(parent, time.Now().Add(timeout))
//line /snap/go/10455/src/context/context.go:685
	// _ = "end of CoverTab[2458]"
}

// WithTimeoutCause behaves like [WithTimeout] but also sets the cause of the
//line /snap/go/10455/src/context/context.go:688
// returned Context when the timeout expires. The returned [CancelFunc] does
//line /snap/go/10455/src/context/context.go:688
// not set the cause.
//line /snap/go/10455/src/context/context.go:691
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (Context, CancelFunc) {
//line /snap/go/10455/src/context/context.go:691
	_go_fuzz_dep_.CoverTab[2459]++
							return WithDeadlineCause(parent, time.Now().Add(timeout), cause)
//line /snap/go/10455/src/context/context.go:692
	// _ = "end of CoverTab[2459]"
}

// WithValue returns a copy of parent in which the value associated with key is
//line /snap/go/10455/src/context/context.go:695
// val.
//line /snap/go/10455/src/context/context.go:695
//
//line /snap/go/10455/src/context/context.go:695
// Use context Values only for request-scoped data that transits processes and
//line /snap/go/10455/src/context/context.go:695
// APIs, not for passing optional parameters to functions.
//line /snap/go/10455/src/context/context.go:695
//
//line /snap/go/10455/src/context/context.go:695
// The provided key must be comparable and should not be of type
//line /snap/go/10455/src/context/context.go:695
// string or any other built-in type to avoid collisions between
//line /snap/go/10455/src/context/context.go:695
// packages using context. Users of WithValue should define their own
//line /snap/go/10455/src/context/context.go:695
// types for keys. To avoid allocating when assigning to an
//line /snap/go/10455/src/context/context.go:695
// interface{}, context keys often have concrete type
//line /snap/go/10455/src/context/context.go:695
// struct{}. Alternatively, exported context key variables' static
//line /snap/go/10455/src/context/context.go:695
// type should be a pointer or interface.
//line /snap/go/10455/src/context/context.go:708
func WithValue(parent Context, key, val any) Context {
//line /snap/go/10455/src/context/context.go:708
	_go_fuzz_dep_.CoverTab[2460]++
							if parent == nil {
//line /snap/go/10455/src/context/context.go:709
		_go_fuzz_dep_.CoverTab[526060]++
//line /snap/go/10455/src/context/context.go:709
		_go_fuzz_dep_.CoverTab[2464]++
								panic("cannot create context from nil parent")
//line /snap/go/10455/src/context/context.go:710
		// _ = "end of CoverTab[2464]"
	} else {
//line /snap/go/10455/src/context/context.go:711
		_go_fuzz_dep_.CoverTab[526061]++
//line /snap/go/10455/src/context/context.go:711
		_go_fuzz_dep_.CoverTab[2465]++
//line /snap/go/10455/src/context/context.go:711
		// _ = "end of CoverTab[2465]"
//line /snap/go/10455/src/context/context.go:711
	}
//line /snap/go/10455/src/context/context.go:711
	// _ = "end of CoverTab[2460]"
//line /snap/go/10455/src/context/context.go:711
	_go_fuzz_dep_.CoverTab[2461]++
							if key == nil {
//line /snap/go/10455/src/context/context.go:712
		_go_fuzz_dep_.CoverTab[526062]++
//line /snap/go/10455/src/context/context.go:712
		_go_fuzz_dep_.CoverTab[2466]++
								panic("nil key")
//line /snap/go/10455/src/context/context.go:713
		// _ = "end of CoverTab[2466]"
	} else {
//line /snap/go/10455/src/context/context.go:714
		_go_fuzz_dep_.CoverTab[526063]++
//line /snap/go/10455/src/context/context.go:714
		_go_fuzz_dep_.CoverTab[2467]++
//line /snap/go/10455/src/context/context.go:714
		// _ = "end of CoverTab[2467]"
//line /snap/go/10455/src/context/context.go:714
	}
//line /snap/go/10455/src/context/context.go:714
	// _ = "end of CoverTab[2461]"
//line /snap/go/10455/src/context/context.go:714
	_go_fuzz_dep_.CoverTab[2462]++
							if !reflectlite.TypeOf(key).Comparable() {
//line /snap/go/10455/src/context/context.go:715
		_go_fuzz_dep_.CoverTab[526064]++
//line /snap/go/10455/src/context/context.go:715
		_go_fuzz_dep_.CoverTab[2468]++
								panic("key is not comparable")
//line /snap/go/10455/src/context/context.go:716
		// _ = "end of CoverTab[2468]"
	} else {
//line /snap/go/10455/src/context/context.go:717
		_go_fuzz_dep_.CoverTab[526065]++
//line /snap/go/10455/src/context/context.go:717
		_go_fuzz_dep_.CoverTab[2469]++
//line /snap/go/10455/src/context/context.go:717
		// _ = "end of CoverTab[2469]"
//line /snap/go/10455/src/context/context.go:717
	}
//line /snap/go/10455/src/context/context.go:717
	// _ = "end of CoverTab[2462]"
//line /snap/go/10455/src/context/context.go:717
	_go_fuzz_dep_.CoverTab[2463]++
							return &valueCtx{parent, key, val}
//line /snap/go/10455/src/context/context.go:718
	// _ = "end of CoverTab[2463]"
}

// A valueCtx carries a key-value pair. It implements Value for that key and
//line /snap/go/10455/src/context/context.go:721
// delegates all other calls to the embedded Context.
//line /snap/go/10455/src/context/context.go:723
type valueCtx struct {
	Context
	key, val	any
}

// stringify tries a bit to stringify v, without using fmt, since we don't
//line /snap/go/10455/src/context/context.go:728
// want context depending on the unicode tables. This is only used by
//line /snap/go/10455/src/context/context.go:728
// *valueCtx.String().
//line /snap/go/10455/src/context/context.go:731
func stringify(v any) string {
//line /snap/go/10455/src/context/context.go:731
	_go_fuzz_dep_.CoverTab[2470]++
							switch s := v.(type) {
	case stringer:
//line /snap/go/10455/src/context/context.go:733
		_go_fuzz_dep_.CoverTab[526066]++
//line /snap/go/10455/src/context/context.go:733
		_go_fuzz_dep_.CoverTab[2472]++
								return s.String()
//line /snap/go/10455/src/context/context.go:734
		// _ = "end of CoverTab[2472]"
	case string:
//line /snap/go/10455/src/context/context.go:735
		_go_fuzz_dep_.CoverTab[526067]++
//line /snap/go/10455/src/context/context.go:735
		_go_fuzz_dep_.CoverTab[2473]++
								return s
//line /snap/go/10455/src/context/context.go:736
		// _ = "end of CoverTab[2473]"
	}
//line /snap/go/10455/src/context/context.go:737
	// _ = "end of CoverTab[2470]"
//line /snap/go/10455/src/context/context.go:737
	_go_fuzz_dep_.CoverTab[2471]++
							return "<not Stringer>"
//line /snap/go/10455/src/context/context.go:738
	// _ = "end of CoverTab[2471]"
}

func (c *valueCtx) String() string {
//line /snap/go/10455/src/context/context.go:741
	_go_fuzz_dep_.CoverTab[2474]++
							return contextName(c.Context) + ".WithValue(type " +
		reflectlite.TypeOf(c.key).String() +
		", val " + stringify(c.val) + ")"
//line /snap/go/10455/src/context/context.go:744
	// _ = "end of CoverTab[2474]"
}

func (c *valueCtx) Value(key any) any {
//line /snap/go/10455/src/context/context.go:747
	_go_fuzz_dep_.CoverTab[2475]++
							if c.key == key {
//line /snap/go/10455/src/context/context.go:748
		_go_fuzz_dep_.CoverTab[526068]++
//line /snap/go/10455/src/context/context.go:748
		_go_fuzz_dep_.CoverTab[2477]++
								return c.val
//line /snap/go/10455/src/context/context.go:749
		// _ = "end of CoverTab[2477]"
	} else {
//line /snap/go/10455/src/context/context.go:750
		_go_fuzz_dep_.CoverTab[526069]++
//line /snap/go/10455/src/context/context.go:750
		_go_fuzz_dep_.CoverTab[2478]++
//line /snap/go/10455/src/context/context.go:750
		// _ = "end of CoverTab[2478]"
//line /snap/go/10455/src/context/context.go:750
	}
//line /snap/go/10455/src/context/context.go:750
	// _ = "end of CoverTab[2475]"
//line /snap/go/10455/src/context/context.go:750
	_go_fuzz_dep_.CoverTab[2476]++
							return value(c.Context, key)
//line /snap/go/10455/src/context/context.go:751
	// _ = "end of CoverTab[2476]"
}

func value(c Context, key any) any {
//line /snap/go/10455/src/context/context.go:754
	_go_fuzz_dep_.CoverTab[2479]++
//line /snap/go/10455/src/context/context.go:754
	_go_fuzz_dep_.CoverTab[786573] = 0
							for {
//line /snap/go/10455/src/context/context.go:755
		if _go_fuzz_dep_.CoverTab[786573] == 0 {
//line /snap/go/10455/src/context/context.go:755
			_go_fuzz_dep_.CoverTab[526088]++
//line /snap/go/10455/src/context/context.go:755
		} else {
//line /snap/go/10455/src/context/context.go:755
			_go_fuzz_dep_.CoverTab[526089]++
//line /snap/go/10455/src/context/context.go:755
		}
//line /snap/go/10455/src/context/context.go:755
		_go_fuzz_dep_.CoverTab[786573] = 1
//line /snap/go/10455/src/context/context.go:755
		_go_fuzz_dep_.CoverTab[2480]++
								switch ctx := c.(type) {
		case *valueCtx:
//line /snap/go/10455/src/context/context.go:757
			_go_fuzz_dep_.CoverTab[526070]++
//line /snap/go/10455/src/context/context.go:757
			_go_fuzz_dep_.CoverTab[2481]++
									if key == ctx.key {
//line /snap/go/10455/src/context/context.go:758
				_go_fuzz_dep_.CoverTab[526076]++
//line /snap/go/10455/src/context/context.go:758
				_go_fuzz_dep_.CoverTab[2491]++
										return ctx.val
//line /snap/go/10455/src/context/context.go:759
				// _ = "end of CoverTab[2491]"
			} else {
//line /snap/go/10455/src/context/context.go:760
				_go_fuzz_dep_.CoverTab[526077]++
//line /snap/go/10455/src/context/context.go:760
				_go_fuzz_dep_.CoverTab[2492]++
//line /snap/go/10455/src/context/context.go:760
				// _ = "end of CoverTab[2492]"
//line /snap/go/10455/src/context/context.go:760
			}
//line /snap/go/10455/src/context/context.go:760
			// _ = "end of CoverTab[2481]"
//line /snap/go/10455/src/context/context.go:760
			_go_fuzz_dep_.CoverTab[2482]++
									c = ctx.Context
//line /snap/go/10455/src/context/context.go:761
			// _ = "end of CoverTab[2482]"
		case *cancelCtx:
//line /snap/go/10455/src/context/context.go:762
			_go_fuzz_dep_.CoverTab[526071]++
//line /snap/go/10455/src/context/context.go:762
			_go_fuzz_dep_.CoverTab[2483]++
									if key == &cancelCtxKey {
//line /snap/go/10455/src/context/context.go:763
				_go_fuzz_dep_.CoverTab[526078]++
//line /snap/go/10455/src/context/context.go:763
				_go_fuzz_dep_.CoverTab[2493]++
										return c
//line /snap/go/10455/src/context/context.go:764
				// _ = "end of CoverTab[2493]"
			} else {
//line /snap/go/10455/src/context/context.go:765
				_go_fuzz_dep_.CoverTab[526079]++
//line /snap/go/10455/src/context/context.go:765
				_go_fuzz_dep_.CoverTab[2494]++
//line /snap/go/10455/src/context/context.go:765
				// _ = "end of CoverTab[2494]"
//line /snap/go/10455/src/context/context.go:765
			}
//line /snap/go/10455/src/context/context.go:765
			// _ = "end of CoverTab[2483]"
//line /snap/go/10455/src/context/context.go:765
			_go_fuzz_dep_.CoverTab[2484]++
									c = ctx.Context
//line /snap/go/10455/src/context/context.go:766
			// _ = "end of CoverTab[2484]"
		case withoutCancelCtx:
//line /snap/go/10455/src/context/context.go:767
			_go_fuzz_dep_.CoverTab[526072]++
//line /snap/go/10455/src/context/context.go:767
			_go_fuzz_dep_.CoverTab[2485]++
									if key == &cancelCtxKey {
//line /snap/go/10455/src/context/context.go:768
				_go_fuzz_dep_.CoverTab[526080]++
//line /snap/go/10455/src/context/context.go:768
				_go_fuzz_dep_.CoverTab[2495]++

//line /snap/go/10455/src/context/context.go:771
				return nil
//line /snap/go/10455/src/context/context.go:771
				// _ = "end of CoverTab[2495]"
			} else {
//line /snap/go/10455/src/context/context.go:772
				_go_fuzz_dep_.CoverTab[526081]++
//line /snap/go/10455/src/context/context.go:772
				_go_fuzz_dep_.CoverTab[2496]++
//line /snap/go/10455/src/context/context.go:772
				// _ = "end of CoverTab[2496]"
//line /snap/go/10455/src/context/context.go:772
			}
//line /snap/go/10455/src/context/context.go:772
			// _ = "end of CoverTab[2485]"
//line /snap/go/10455/src/context/context.go:772
			_go_fuzz_dep_.CoverTab[2486]++
									c = ctx.c
//line /snap/go/10455/src/context/context.go:773
			// _ = "end of CoverTab[2486]"
		case *timerCtx:
//line /snap/go/10455/src/context/context.go:774
			_go_fuzz_dep_.CoverTab[526073]++
//line /snap/go/10455/src/context/context.go:774
			_go_fuzz_dep_.CoverTab[2487]++
									if key == &cancelCtxKey {
//line /snap/go/10455/src/context/context.go:775
				_go_fuzz_dep_.CoverTab[526082]++
//line /snap/go/10455/src/context/context.go:775
				_go_fuzz_dep_.CoverTab[2497]++
										return &ctx.cancelCtx
//line /snap/go/10455/src/context/context.go:776
				// _ = "end of CoverTab[2497]"
			} else {
//line /snap/go/10455/src/context/context.go:777
				_go_fuzz_dep_.CoverTab[526083]++
//line /snap/go/10455/src/context/context.go:777
				_go_fuzz_dep_.CoverTab[2498]++
//line /snap/go/10455/src/context/context.go:777
				// _ = "end of CoverTab[2498]"
//line /snap/go/10455/src/context/context.go:777
			}
//line /snap/go/10455/src/context/context.go:777
			// _ = "end of CoverTab[2487]"
//line /snap/go/10455/src/context/context.go:777
			_go_fuzz_dep_.CoverTab[2488]++
									c = ctx.Context
//line /snap/go/10455/src/context/context.go:778
			// _ = "end of CoverTab[2488]"
		case backgroundCtx, todoCtx:
//line /snap/go/10455/src/context/context.go:779
			_go_fuzz_dep_.CoverTab[526074]++
//line /snap/go/10455/src/context/context.go:779
			_go_fuzz_dep_.CoverTab[2489]++
									return nil
//line /snap/go/10455/src/context/context.go:780
			// _ = "end of CoverTab[2489]"
		default:
//line /snap/go/10455/src/context/context.go:781
			_go_fuzz_dep_.CoverTab[526075]++
//line /snap/go/10455/src/context/context.go:781
			_go_fuzz_dep_.CoverTab[2490]++
									return c.Value(key)
//line /snap/go/10455/src/context/context.go:782
			// _ = "end of CoverTab[2490]"
		}
//line /snap/go/10455/src/context/context.go:783
		// _ = "end of CoverTab[2480]"
	}
//line /snap/go/10455/src/context/context.go:784
	// _ = "end of CoverTab[2479]"
}

//line /snap/go/10455/src/context/context.go:785
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/context/context.go:785
var _ = _go_fuzz_dep_.CoverTab
