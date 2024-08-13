// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/context/context.go:5
// Package context defines the Context type, which carries deadlines,
//line /usr/local/go/src/context/context.go:5
// cancellation signals, and other request-scoped values across API boundaries
//line /usr/local/go/src/context/context.go:5
// and between processes.
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// Incoming requests to a server should create a Context, and outgoing
//line /usr/local/go/src/context/context.go:5
// calls to servers should accept a Context. The chain of function
//line /usr/local/go/src/context/context.go:5
// calls between them must propagate the Context, optionally replacing
//line /usr/local/go/src/context/context.go:5
// it with a derived Context created using WithCancel, WithDeadline,
//line /usr/local/go/src/context/context.go:5
// WithTimeout, or WithValue. When a Context is canceled, all
//line /usr/local/go/src/context/context.go:5
// Contexts derived from it are also canceled.
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// The WithCancel, WithDeadline, and WithTimeout functions take a
//line /usr/local/go/src/context/context.go:5
// Context (the parent) and return a derived Context (the child) and a
//line /usr/local/go/src/context/context.go:5
// CancelFunc. Calling the CancelFunc cancels the child and its
//line /usr/local/go/src/context/context.go:5
// children, removes the parent's reference to the child, and stops
//line /usr/local/go/src/context/context.go:5
// any associated timers. Failing to call the CancelFunc leaks the
//line /usr/local/go/src/context/context.go:5
// child and its children until the parent is canceled or the timer
//line /usr/local/go/src/context/context.go:5
// fires. The go vet tool checks that CancelFuncs are used on all
//line /usr/local/go/src/context/context.go:5
// control-flow paths.
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// The WithCancelCause function returns a CancelCauseFunc, which
//line /usr/local/go/src/context/context.go:5
// takes an error and records it as the cancellation cause. Calling
//line /usr/local/go/src/context/context.go:5
// Cause on the canceled context or any of its children retrieves
//line /usr/local/go/src/context/context.go:5
// the cause. If no cause is specified, Cause(ctx) returns the same
//line /usr/local/go/src/context/context.go:5
// value as ctx.Err().
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// Programs that use Contexts should follow these rules to keep interfaces
//line /usr/local/go/src/context/context.go:5
// consistent across packages and enable static analysis tools to check context
//line /usr/local/go/src/context/context.go:5
// propagation:
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// Do not store Contexts inside a struct type; instead, pass a Context
//line /usr/local/go/src/context/context.go:5
// explicitly to each function that needs it. The Context should be the first
//line /usr/local/go/src/context/context.go:5
// parameter, typically named ctx:
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
//	func DoSomething(ctx context.Context, arg Arg) error {
//line /usr/local/go/src/context/context.go:5
//		// ... use ctx ...
//line /usr/local/go/src/context/context.go:5
//	}
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// Do not pass a nil Context, even if a function permits it. Pass context.TODO
//line /usr/local/go/src/context/context.go:5
// if you are unsure about which Context to use.
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// Use context Values only for request-scoped data that transits processes and
//line /usr/local/go/src/context/context.go:5
// APIs, not for passing optional parameters to functions.
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// The same Context may be passed to functions running in different goroutines;
//line /usr/local/go/src/context/context.go:5
// Contexts are safe for simultaneous use by multiple goroutines.
//line /usr/local/go/src/context/context.go:5
//
//line /usr/local/go/src/context/context.go:5
// See https://blog.golang.org/context for example code for a server that uses
//line /usr/local/go/src/context/context.go:5
// Contexts.
//line /usr/local/go/src/context/context.go:54
package context

//line /usr/local/go/src/context/context.go:54
import (
//line /usr/local/go/src/context/context.go:54
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/context/context.go:54
)
//line /usr/local/go/src/context/context.go:54
import (
//line /usr/local/go/src/context/context.go:54
	_atomic_ "sync/atomic"
//line /usr/local/go/src/context/context.go:54
)

import (
	"errors"
	"internal/reflectlite"
	"sync"
	"sync/atomic"
	"time"
)

// A Context carries a deadline, a cancellation signal, and other values across
//line /usr/local/go/src/context/context.go:64
// API boundaries.
//line /usr/local/go/src/context/context.go:64
//
//line /usr/local/go/src/context/context.go:64
// Context's methods may be called by multiple goroutines simultaneously.
//line /usr/local/go/src/context/context.go:68
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

// Canceled is the error returned by Context.Err when the context is canceled.
var Canceled = errors.New("context canceled")

// DeadlineExceeded is the error returned by Context.Err when the context's
//line /usr/local/go/src/context/context.go:165
// deadline passes.
//line /usr/local/go/src/context/context.go:167
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string {
//line /usr/local/go/src/context/context.go:171
	_go_fuzz_dep_.CoverTab[2277]++
//line /usr/local/go/src/context/context.go:171
	return "context deadline exceeded"
//line /usr/local/go/src/context/context.go:171
	// _ = "end of CoverTab[2277]"
//line /usr/local/go/src/context/context.go:171
}
func (deadlineExceededError) Timeout() bool {
//line /usr/local/go/src/context/context.go:172
	_go_fuzz_dep_.CoverTab[2278]++
//line /usr/local/go/src/context/context.go:172
	return true
//line /usr/local/go/src/context/context.go:172
	// _ = "end of CoverTab[2278]"
//line /usr/local/go/src/context/context.go:172
}
func (deadlineExceededError) Temporary() bool {
//line /usr/local/go/src/context/context.go:173
	_go_fuzz_dep_.CoverTab[2279]++
//line /usr/local/go/src/context/context.go:173
	return true
//line /usr/local/go/src/context/context.go:173
	// _ = "end of CoverTab[2279]"
//line /usr/local/go/src/context/context.go:173
}

// An emptyCtx is never canceled, has no values, and has no deadline. It is not
//line /usr/local/go/src/context/context.go:175
// struct{}, since vars of this type must have distinct addresses.
//line /usr/local/go/src/context/context.go:177
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
//line /usr/local/go/src/context/context.go:179
	_go_fuzz_dep_.CoverTab[2280]++
							return
//line /usr/local/go/src/context/context.go:180
	// _ = "end of CoverTab[2280]"
}

func (*emptyCtx) Done() <-chan struct{} {
//line /usr/local/go/src/context/context.go:183
	_go_fuzz_dep_.CoverTab[2281]++
							return nil
//line /usr/local/go/src/context/context.go:184
	// _ = "end of CoverTab[2281]"
}

func (*emptyCtx) Err() error {
//line /usr/local/go/src/context/context.go:187
	_go_fuzz_dep_.CoverTab[2282]++
							return nil
//line /usr/local/go/src/context/context.go:188
	// _ = "end of CoverTab[2282]"
}

func (*emptyCtx) Value(key any) any {
//line /usr/local/go/src/context/context.go:191
	_go_fuzz_dep_.CoverTab[2283]++
							return nil
//line /usr/local/go/src/context/context.go:192
	// _ = "end of CoverTab[2283]"
}

func (e *emptyCtx) String() string {
//line /usr/local/go/src/context/context.go:195
	_go_fuzz_dep_.CoverTab[2284]++
							switch e {
	case background:
//line /usr/local/go/src/context/context.go:197
		_go_fuzz_dep_.CoverTab[2286]++
								return "context.Background"
//line /usr/local/go/src/context/context.go:198
		// _ = "end of CoverTab[2286]"
	case todo:
//line /usr/local/go/src/context/context.go:199
		_go_fuzz_dep_.CoverTab[2287]++
								return "context.TODO"
//line /usr/local/go/src/context/context.go:200
		// _ = "end of CoverTab[2287]"
//line /usr/local/go/src/context/context.go:200
	default:
//line /usr/local/go/src/context/context.go:200
		_go_fuzz_dep_.CoverTab[2288]++
//line /usr/local/go/src/context/context.go:200
		// _ = "end of CoverTab[2288]"
	}
//line /usr/local/go/src/context/context.go:201
	// _ = "end of CoverTab[2284]"
//line /usr/local/go/src/context/context.go:201
	_go_fuzz_dep_.CoverTab[2285]++
							return "unknown empty Context"
//line /usr/local/go/src/context/context.go:202
	// _ = "end of CoverTab[2285]"
}

var (
	background	= new(emptyCtx)
	todo		= new(emptyCtx)
)

// Background returns a non-nil, empty Context. It is never canceled, has no
//line /usr/local/go/src/context/context.go:210
// values, and has no deadline. It is typically used by the main function,
//line /usr/local/go/src/context/context.go:210
// initialization, and tests, and as the top-level Context for incoming
//line /usr/local/go/src/context/context.go:210
// requests.
//line /usr/local/go/src/context/context.go:214
func Background() Context {
//line /usr/local/go/src/context/context.go:214
	_go_fuzz_dep_.CoverTab[2289]++
							return background
//line /usr/local/go/src/context/context.go:215
	// _ = "end of CoverTab[2289]"
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
//line /usr/local/go/src/context/context.go:218
// it's unclear which Context to use or it is not yet available (because the
//line /usr/local/go/src/context/context.go:218
// surrounding function has not yet been extended to accept a Context
//line /usr/local/go/src/context/context.go:218
// parameter).
//line /usr/local/go/src/context/context.go:222
func TODO() Context {
//line /usr/local/go/src/context/context.go:222
	_go_fuzz_dep_.CoverTab[2290]++
							return todo
//line /usr/local/go/src/context/context.go:223
	// _ = "end of CoverTab[2290]"
}

// A CancelFunc tells an operation to abandon its work.
//line /usr/local/go/src/context/context.go:226
// A CancelFunc does not wait for the work to stop.
//line /usr/local/go/src/context/context.go:226
// A CancelFunc may be called by multiple goroutines simultaneously.
//line /usr/local/go/src/context/context.go:226
// After the first call, subsequent calls to a CancelFunc do nothing.
//line /usr/local/go/src/context/context.go:230
type CancelFunc func()

// WithCancel returns a copy of parent with a new Done channel. The returned
//line /usr/local/go/src/context/context.go:232
// context's Done channel is closed when the returned cancel function is called
//line /usr/local/go/src/context/context.go:232
// or when the parent context's Done channel is closed, whichever happens first.
//line /usr/local/go/src/context/context.go:232
//
//line /usr/local/go/src/context/context.go:232
// Canceling this context releases resources associated with it, so code should
//line /usr/local/go/src/context/context.go:232
// call cancel as soon as the operations running in this Context complete.
//line /usr/local/go/src/context/context.go:238
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
//line /usr/local/go/src/context/context.go:238
	_go_fuzz_dep_.CoverTab[2291]++
							c := withCancel(parent)
							return c, func() { _go_fuzz_dep_.CoverTab[2292]++; c.cancel(true, Canceled, nil); // _ = "end of CoverTab[2292]" }
//line /usr/local/go/src/context/context.go:240
	// _ = "end of CoverTab[2291]"
}

// A CancelCauseFunc behaves like a CancelFunc but additionally sets the cancellation cause.
//line /usr/local/go/src/context/context.go:243
// This cause can be retrieved by calling Cause on the canceled Context or on
//line /usr/local/go/src/context/context.go:243
// any of its derived Contexts.
//line /usr/local/go/src/context/context.go:243
//
//line /usr/local/go/src/context/context.go:243
// If the context has already been canceled, CancelCauseFunc does not set the cause.
//line /usr/local/go/src/context/context.go:243
// For example, if childContext is derived from parentContext:
//line /usr/local/go/src/context/context.go:243
//   - if parentContext is canceled with cause1 before childContext is canceled with cause2,
//line /usr/local/go/src/context/context.go:243
//     then Cause(parentContext) == Cause(childContext) == cause1
//line /usr/local/go/src/context/context.go:243
//   - if childContext is canceled with cause2 before parentContext is canceled with cause1,
//line /usr/local/go/src/context/context.go:243
//     then Cause(parentContext) == cause1 and Cause(childContext) == cause2
//line /usr/local/go/src/context/context.go:253
type CancelCauseFunc func(cause error)

// WithCancelCause behaves like WithCancel but returns a CancelCauseFunc instead of a CancelFunc.
//line /usr/local/go/src/context/context.go:255
// Calling cancel with a non-nil error (the "cause") records that error in ctx;
//line /usr/local/go/src/context/context.go:255
// it can then be retrieved using Cause(ctx).
//line /usr/local/go/src/context/context.go:255
// Calling cancel with nil sets the cause to Canceled.
//line /usr/local/go/src/context/context.go:255
//
//line /usr/local/go/src/context/context.go:255
// Example use:
//line /usr/local/go/src/context/context.go:255
//
//line /usr/local/go/src/context/context.go:255
//	ctx, cancel := context.WithCancelCause(parent)
//line /usr/local/go/src/context/context.go:255
//	cancel(myError)
//line /usr/local/go/src/context/context.go:255
//	ctx.Err() // returns context.Canceled
//line /usr/local/go/src/context/context.go:255
//	context.Cause(ctx) // returns myError
//line /usr/local/go/src/context/context.go:266
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc) {
//line /usr/local/go/src/context/context.go:266
	_go_fuzz_dep_.CoverTab[2293]++
							c := withCancel(parent)
							return c, func(cause error) {
//line /usr/local/go/src/context/context.go:268
		_go_fuzz_dep_.CoverTab[2294]++
//line /usr/local/go/src/context/context.go:268
		c.cancel(true, Canceled, cause)
//line /usr/local/go/src/context/context.go:268
		// _ = "end of CoverTab[2294]"
//line /usr/local/go/src/context/context.go:268
	}
//line /usr/local/go/src/context/context.go:268
	// _ = "end of CoverTab[2293]"
}

func withCancel(parent Context) *cancelCtx {
//line /usr/local/go/src/context/context.go:271
	_go_fuzz_dep_.CoverTab[2295]++
							if parent == nil {
//line /usr/local/go/src/context/context.go:272
		_go_fuzz_dep_.CoverTab[2297]++
								panic("cannot create context from nil parent")
//line /usr/local/go/src/context/context.go:273
		// _ = "end of CoverTab[2297]"
	} else {
//line /usr/local/go/src/context/context.go:274
		_go_fuzz_dep_.CoverTab[2298]++
//line /usr/local/go/src/context/context.go:274
		// _ = "end of CoverTab[2298]"
//line /usr/local/go/src/context/context.go:274
	}
//line /usr/local/go/src/context/context.go:274
	// _ = "end of CoverTab[2295]"
//line /usr/local/go/src/context/context.go:274
	_go_fuzz_dep_.CoverTab[2296]++
							c := newCancelCtx(parent)
							propagateCancel(parent, c)
							return c
//line /usr/local/go/src/context/context.go:277
	// _ = "end of CoverTab[2296]"
}

// Cause returns a non-nil error explaining why c was canceled.
//line /usr/local/go/src/context/context.go:280
// The first cancellation of c or one of its parents sets the cause.
//line /usr/local/go/src/context/context.go:280
// If that cancellation happened via a call to CancelCauseFunc(err),
//line /usr/local/go/src/context/context.go:280
// then Cause returns err.
//line /usr/local/go/src/context/context.go:280
// Otherwise Cause(c) returns the same value as c.Err().
//line /usr/local/go/src/context/context.go:280
// Cause returns nil if c has not been canceled yet.
//line /usr/local/go/src/context/context.go:286
func Cause(c Context) error {
//line /usr/local/go/src/context/context.go:286
	_go_fuzz_dep_.CoverTab[2299]++
							if cc, ok := c.Value(&cancelCtxKey).(*cancelCtx); ok {
//line /usr/local/go/src/context/context.go:287
		_go_fuzz_dep_.CoverTab[2301]++
								cc.mu.Lock()
								defer cc.mu.Unlock()
								return cc.cause
//line /usr/local/go/src/context/context.go:290
		// _ = "end of CoverTab[2301]"
	} else {
//line /usr/local/go/src/context/context.go:291
		_go_fuzz_dep_.CoverTab[2302]++
//line /usr/local/go/src/context/context.go:291
		// _ = "end of CoverTab[2302]"
//line /usr/local/go/src/context/context.go:291
	}
//line /usr/local/go/src/context/context.go:291
	// _ = "end of CoverTab[2299]"
//line /usr/local/go/src/context/context.go:291
	_go_fuzz_dep_.CoverTab[2300]++
							return nil
//line /usr/local/go/src/context/context.go:292
	// _ = "end of CoverTab[2300]"
}

// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) *cancelCtx {
//line /usr/local/go/src/context/context.go:296
	_go_fuzz_dep_.CoverTab[2303]++
							return &cancelCtx{Context: parent}
//line /usr/local/go/src/context/context.go:297
	// _ = "end of CoverTab[2303]"
}

// goroutines counts the number of goroutines ever created; for testing.
var goroutines atomic.Int32

// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
//line /usr/local/go/src/context/context.go:304
	_go_fuzz_dep_.CoverTab[2304]++
							done := parent.Done()
							if done == nil {
//line /usr/local/go/src/context/context.go:306
		_go_fuzz_dep_.CoverTab[2307]++
								return
//line /usr/local/go/src/context/context.go:307
		// _ = "end of CoverTab[2307]"
	} else {
//line /usr/local/go/src/context/context.go:308
		_go_fuzz_dep_.CoverTab[2308]++
//line /usr/local/go/src/context/context.go:308
		// _ = "end of CoverTab[2308]"
//line /usr/local/go/src/context/context.go:308
	}
//line /usr/local/go/src/context/context.go:308
	// _ = "end of CoverTab[2304]"
//line /usr/local/go/src/context/context.go:308
	_go_fuzz_dep_.CoverTab[2305]++

							select {
	case <-done:
//line /usr/local/go/src/context/context.go:311
		_go_fuzz_dep_.CoverTab[2309]++

								child.cancel(false, parent.Err(), Cause(parent))
								return
//line /usr/local/go/src/context/context.go:314
		// _ = "end of CoverTab[2309]"
	default:
//line /usr/local/go/src/context/context.go:315
		_go_fuzz_dep_.CoverTab[2310]++
//line /usr/local/go/src/context/context.go:315
		// _ = "end of CoverTab[2310]"
	}
//line /usr/local/go/src/context/context.go:316
	// _ = "end of CoverTab[2305]"
//line /usr/local/go/src/context/context.go:316
	_go_fuzz_dep_.CoverTab[2306]++

							if p, ok := parentCancelCtx(parent); ok {
//line /usr/local/go/src/context/context.go:318
		_go_fuzz_dep_.CoverTab[2311]++
								p.mu.Lock()
								if p.err != nil {
//line /usr/local/go/src/context/context.go:320
			_go_fuzz_dep_.CoverTab[2313]++

									child.cancel(false, p.err, p.cause)
//line /usr/local/go/src/context/context.go:322
			// _ = "end of CoverTab[2313]"
		} else {
//line /usr/local/go/src/context/context.go:323
			_go_fuzz_dep_.CoverTab[2314]++
									if p.children == nil {
//line /usr/local/go/src/context/context.go:324
				_go_fuzz_dep_.CoverTab[2316]++
										p.children = make(map[canceler]struct{})
//line /usr/local/go/src/context/context.go:325
				// _ = "end of CoverTab[2316]"
			} else {
//line /usr/local/go/src/context/context.go:326
				_go_fuzz_dep_.CoverTab[2317]++
//line /usr/local/go/src/context/context.go:326
				// _ = "end of CoverTab[2317]"
//line /usr/local/go/src/context/context.go:326
			}
//line /usr/local/go/src/context/context.go:326
			// _ = "end of CoverTab[2314]"
//line /usr/local/go/src/context/context.go:326
			_go_fuzz_dep_.CoverTab[2315]++
									p.children[child] = struct{}{}
//line /usr/local/go/src/context/context.go:327
			// _ = "end of CoverTab[2315]"
		}
//line /usr/local/go/src/context/context.go:328
		// _ = "end of CoverTab[2311]"
//line /usr/local/go/src/context/context.go:328
		_go_fuzz_dep_.CoverTab[2312]++
								p.mu.Unlock()
//line /usr/local/go/src/context/context.go:329
		// _ = "end of CoverTab[2312]"
	} else {
//line /usr/local/go/src/context/context.go:330
		_go_fuzz_dep_.CoverTab[2318]++
								goroutines.Add(1)
//line /usr/local/go/src/context/context.go:331
		_curRoutineNum0_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/context/context.go:331
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum0_)
								go func() {
//line /usr/local/go/src/context/context.go:332
			_go_fuzz_dep_.CoverTab[2319]++
//line /usr/local/go/src/context/context.go:332
			defer func() {
//line /usr/local/go/src/context/context.go:332
				_go_fuzz_dep_.CoverTab[2320]++
//line /usr/local/go/src/context/context.go:332
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum0_)
//line /usr/local/go/src/context/context.go:332
				// _ = "end of CoverTab[2320]"
//line /usr/local/go/src/context/context.go:332
			}()
									select {
			case <-parent.Done():
//line /usr/local/go/src/context/context.go:334
				_go_fuzz_dep_.CoverTab[2321]++
										child.cancel(false, parent.Err(), Cause(parent))
//line /usr/local/go/src/context/context.go:335
				// _ = "end of CoverTab[2321]"
			case <-child.Done():
//line /usr/local/go/src/context/context.go:336
				_go_fuzz_dep_.CoverTab[2322]++
//line /usr/local/go/src/context/context.go:336
				// _ = "end of CoverTab[2322]"
			}
//line /usr/local/go/src/context/context.go:337
			// _ = "end of CoverTab[2319]"
		}()
//line /usr/local/go/src/context/context.go:338
		// _ = "end of CoverTab[2318]"
	}
//line /usr/local/go/src/context/context.go:339
	// _ = "end of CoverTab[2306]"
}

// &cancelCtxKey is the key that a cancelCtx returns itself for.
var cancelCtxKey int

// parentCancelCtx returns the underlying *cancelCtx for parent.
//line /usr/local/go/src/context/context.go:345
// It does this by looking up parent.Value(&cancelCtxKey) to find
//line /usr/local/go/src/context/context.go:345
// the innermost enclosing *cancelCtx and then checking whether
//line /usr/local/go/src/context/context.go:345
// parent.Done() matches that *cancelCtx. (If not, the *cancelCtx
//line /usr/local/go/src/context/context.go:345
// has been wrapped in a custom implementation providing a
//line /usr/local/go/src/context/context.go:345
// different done channel, in which case we should not bypass it.)
//line /usr/local/go/src/context/context.go:351
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
//line /usr/local/go/src/context/context.go:351
	_go_fuzz_dep_.CoverTab[2323]++
							done := parent.Done()
							if done == closedchan || func() bool {
//line /usr/local/go/src/context/context.go:353
		_go_fuzz_dep_.CoverTab[2327]++
//line /usr/local/go/src/context/context.go:353
		return done == nil
//line /usr/local/go/src/context/context.go:353
		// _ = "end of CoverTab[2327]"
//line /usr/local/go/src/context/context.go:353
	}() {
//line /usr/local/go/src/context/context.go:353
		_go_fuzz_dep_.CoverTab[2328]++
								return nil, false
//line /usr/local/go/src/context/context.go:354
		// _ = "end of CoverTab[2328]"
	} else {
//line /usr/local/go/src/context/context.go:355
		_go_fuzz_dep_.CoverTab[2329]++
//line /usr/local/go/src/context/context.go:355
		// _ = "end of CoverTab[2329]"
//line /usr/local/go/src/context/context.go:355
	}
//line /usr/local/go/src/context/context.go:355
	// _ = "end of CoverTab[2323]"
//line /usr/local/go/src/context/context.go:355
	_go_fuzz_dep_.CoverTab[2324]++
							p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
							if !ok {
//line /usr/local/go/src/context/context.go:357
		_go_fuzz_dep_.CoverTab[2330]++
								return nil, false
//line /usr/local/go/src/context/context.go:358
		// _ = "end of CoverTab[2330]"
	} else {
//line /usr/local/go/src/context/context.go:359
		_go_fuzz_dep_.CoverTab[2331]++
//line /usr/local/go/src/context/context.go:359
		// _ = "end of CoverTab[2331]"
//line /usr/local/go/src/context/context.go:359
	}
//line /usr/local/go/src/context/context.go:359
	// _ = "end of CoverTab[2324]"
//line /usr/local/go/src/context/context.go:359
	_go_fuzz_dep_.CoverTab[2325]++
							pdone, _ := p.done.Load().(chan struct{})
							if pdone != done {
//line /usr/local/go/src/context/context.go:361
		_go_fuzz_dep_.CoverTab[2332]++
								return nil, false
//line /usr/local/go/src/context/context.go:362
		// _ = "end of CoverTab[2332]"
	} else {
//line /usr/local/go/src/context/context.go:363
		_go_fuzz_dep_.CoverTab[2333]++
//line /usr/local/go/src/context/context.go:363
		// _ = "end of CoverTab[2333]"
//line /usr/local/go/src/context/context.go:363
	}
//line /usr/local/go/src/context/context.go:363
	// _ = "end of CoverTab[2325]"
//line /usr/local/go/src/context/context.go:363
	_go_fuzz_dep_.CoverTab[2326]++
							return p, true
//line /usr/local/go/src/context/context.go:364
	// _ = "end of CoverTab[2326]"
}

// removeChild removes a context from its parent.
func removeChild(parent Context, child canceler) {
//line /usr/local/go/src/context/context.go:368
	_go_fuzz_dep_.CoverTab[2334]++
							p, ok := parentCancelCtx(parent)
							if !ok {
//line /usr/local/go/src/context/context.go:370
		_go_fuzz_dep_.CoverTab[2337]++
								return
//line /usr/local/go/src/context/context.go:371
		// _ = "end of CoverTab[2337]"
	} else {
//line /usr/local/go/src/context/context.go:372
		_go_fuzz_dep_.CoverTab[2338]++
//line /usr/local/go/src/context/context.go:372
		// _ = "end of CoverTab[2338]"
//line /usr/local/go/src/context/context.go:372
	}
//line /usr/local/go/src/context/context.go:372
	// _ = "end of CoverTab[2334]"
//line /usr/local/go/src/context/context.go:372
	_go_fuzz_dep_.CoverTab[2335]++
							p.mu.Lock()
							if p.children != nil {
//line /usr/local/go/src/context/context.go:374
		_go_fuzz_dep_.CoverTab[2339]++
								delete(p.children, child)
//line /usr/local/go/src/context/context.go:375
		// _ = "end of CoverTab[2339]"
	} else {
//line /usr/local/go/src/context/context.go:376
		_go_fuzz_dep_.CoverTab[2340]++
//line /usr/local/go/src/context/context.go:376
		// _ = "end of CoverTab[2340]"
//line /usr/local/go/src/context/context.go:376
	}
//line /usr/local/go/src/context/context.go:376
	// _ = "end of CoverTab[2335]"
//line /usr/local/go/src/context/context.go:376
	_go_fuzz_dep_.CoverTab[2336]++
							p.mu.Unlock()
//line /usr/local/go/src/context/context.go:377
	// _ = "end of CoverTab[2336]"
}

// A canceler is a context type that can be canceled directly. The
//line /usr/local/go/src/context/context.go:380
// implementations are *cancelCtx and *timerCtx.
//line /usr/local/go/src/context/context.go:382
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
//line /usr/local/go/src/context/context.go:394
// that implement canceler.
//line /usr/local/go/src/context/context.go:396
type cancelCtx struct {
	Context

	mu		sync.Mutex		// protects following fields
	done		atomic.Value		// of chan struct{}, created lazily, closed by first cancel call
	children	map[canceler]struct{}	// set to nil by the first cancel call
	err		error			// set to non-nil by the first cancel call
	cause		error			// set to non-nil by the first cancel call
}

func (c *cancelCtx) Value(key any) any {
//line /usr/local/go/src/context/context.go:406
	_go_fuzz_dep_.CoverTab[2341]++
							if key == &cancelCtxKey {
//line /usr/local/go/src/context/context.go:407
		_go_fuzz_dep_.CoverTab[2343]++
								return c
//line /usr/local/go/src/context/context.go:408
		// _ = "end of CoverTab[2343]"
	} else {
//line /usr/local/go/src/context/context.go:409
		_go_fuzz_dep_.CoverTab[2344]++
//line /usr/local/go/src/context/context.go:409
		// _ = "end of CoverTab[2344]"
//line /usr/local/go/src/context/context.go:409
	}
//line /usr/local/go/src/context/context.go:409
	// _ = "end of CoverTab[2341]"
//line /usr/local/go/src/context/context.go:409
	_go_fuzz_dep_.CoverTab[2342]++
							return value(c.Context, key)
//line /usr/local/go/src/context/context.go:410
	// _ = "end of CoverTab[2342]"
}

func (c *cancelCtx) Done() <-chan struct{} {
//line /usr/local/go/src/context/context.go:413
	_go_fuzz_dep_.CoverTab[2345]++
							d := c.done.Load()
							if d != nil {
//line /usr/local/go/src/context/context.go:415
		_go_fuzz_dep_.CoverTab[2348]++
								return d.(chan struct{})
//line /usr/local/go/src/context/context.go:416
		// _ = "end of CoverTab[2348]"
	} else {
//line /usr/local/go/src/context/context.go:417
		_go_fuzz_dep_.CoverTab[2349]++
//line /usr/local/go/src/context/context.go:417
		// _ = "end of CoverTab[2349]"
//line /usr/local/go/src/context/context.go:417
	}
//line /usr/local/go/src/context/context.go:417
	// _ = "end of CoverTab[2345]"
//line /usr/local/go/src/context/context.go:417
	_go_fuzz_dep_.CoverTab[2346]++
							c.mu.Lock()
							defer c.mu.Unlock()
							d = c.done.Load()
							if d == nil {
//line /usr/local/go/src/context/context.go:421
		_go_fuzz_dep_.CoverTab[2350]++
								d = make(chan struct{})
								c.done.Store(d)
//line /usr/local/go/src/context/context.go:423
		// _ = "end of CoverTab[2350]"
	} else {
//line /usr/local/go/src/context/context.go:424
		_go_fuzz_dep_.CoverTab[2351]++
//line /usr/local/go/src/context/context.go:424
		// _ = "end of CoverTab[2351]"
//line /usr/local/go/src/context/context.go:424
	}
//line /usr/local/go/src/context/context.go:424
	// _ = "end of CoverTab[2346]"
//line /usr/local/go/src/context/context.go:424
	_go_fuzz_dep_.CoverTab[2347]++
							return d.(chan struct{})
//line /usr/local/go/src/context/context.go:425
	// _ = "end of CoverTab[2347]"
}

func (c *cancelCtx) Err() error {
//line /usr/local/go/src/context/context.go:428
	_go_fuzz_dep_.CoverTab[2352]++
							c.mu.Lock()
							err := c.err
							c.mu.Unlock()
							return err
//line /usr/local/go/src/context/context.go:432
	// _ = "end of CoverTab[2352]"
}

type stringer interface {
	String() string
}

func contextName(c Context) string {
//line /usr/local/go/src/context/context.go:439
	_go_fuzz_dep_.CoverTab[2353]++
							if s, ok := c.(stringer); ok {
//line /usr/local/go/src/context/context.go:440
		_go_fuzz_dep_.CoverTab[2355]++
								return s.String()
//line /usr/local/go/src/context/context.go:441
		// _ = "end of CoverTab[2355]"
	} else {
//line /usr/local/go/src/context/context.go:442
		_go_fuzz_dep_.CoverTab[2356]++
//line /usr/local/go/src/context/context.go:442
		// _ = "end of CoverTab[2356]"
//line /usr/local/go/src/context/context.go:442
	}
//line /usr/local/go/src/context/context.go:442
	// _ = "end of CoverTab[2353]"
//line /usr/local/go/src/context/context.go:442
	_go_fuzz_dep_.CoverTab[2354]++
							return reflectlite.TypeOf(c).String()
//line /usr/local/go/src/context/context.go:443
	// _ = "end of CoverTab[2354]"
}

func (c *cancelCtx) String() string {
//line /usr/local/go/src/context/context.go:446
	_go_fuzz_dep_.CoverTab[2357]++
							return contextName(c.Context) + ".WithCancel"
//line /usr/local/go/src/context/context.go:447
	// _ = "end of CoverTab[2357]"
}

// cancel closes c.done, cancels each of c's children, and, if
//line /usr/local/go/src/context/context.go:450
// removeFromParent is true, removes c from its parent's children.
//line /usr/local/go/src/context/context.go:450
// cancel sets c.cause to cause if this is the first time c is canceled.
//line /usr/local/go/src/context/context.go:453
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
//line /usr/local/go/src/context/context.go:453
	_go_fuzz_dep_.CoverTab[2358]++
							if err == nil {
//line /usr/local/go/src/context/context.go:454
		_go_fuzz_dep_.CoverTab[2364]++
								panic("context: internal error: missing cancel error")
//line /usr/local/go/src/context/context.go:455
		// _ = "end of CoverTab[2364]"
	} else {
//line /usr/local/go/src/context/context.go:456
		_go_fuzz_dep_.CoverTab[2365]++
//line /usr/local/go/src/context/context.go:456
		// _ = "end of CoverTab[2365]"
//line /usr/local/go/src/context/context.go:456
	}
//line /usr/local/go/src/context/context.go:456
	// _ = "end of CoverTab[2358]"
//line /usr/local/go/src/context/context.go:456
	_go_fuzz_dep_.CoverTab[2359]++
							if cause == nil {
//line /usr/local/go/src/context/context.go:457
		_go_fuzz_dep_.CoverTab[2366]++
								cause = err
//line /usr/local/go/src/context/context.go:458
		// _ = "end of CoverTab[2366]"
	} else {
//line /usr/local/go/src/context/context.go:459
		_go_fuzz_dep_.CoverTab[2367]++
//line /usr/local/go/src/context/context.go:459
		// _ = "end of CoverTab[2367]"
//line /usr/local/go/src/context/context.go:459
	}
//line /usr/local/go/src/context/context.go:459
	// _ = "end of CoverTab[2359]"
//line /usr/local/go/src/context/context.go:459
	_go_fuzz_dep_.CoverTab[2360]++
							c.mu.Lock()
							if c.err != nil {
//line /usr/local/go/src/context/context.go:461
		_go_fuzz_dep_.CoverTab[2368]++
								c.mu.Unlock()
								return
//line /usr/local/go/src/context/context.go:463
		// _ = "end of CoverTab[2368]"
	} else {
//line /usr/local/go/src/context/context.go:464
		_go_fuzz_dep_.CoverTab[2369]++
//line /usr/local/go/src/context/context.go:464
		// _ = "end of CoverTab[2369]"
//line /usr/local/go/src/context/context.go:464
	}
//line /usr/local/go/src/context/context.go:464
	// _ = "end of CoverTab[2360]"
//line /usr/local/go/src/context/context.go:464
	_go_fuzz_dep_.CoverTab[2361]++
							c.err = err
							c.cause = cause
							d, _ := c.done.Load().(chan struct{})
							if d == nil {
//line /usr/local/go/src/context/context.go:468
		_go_fuzz_dep_.CoverTab[2370]++
								c.done.Store(closedchan)
//line /usr/local/go/src/context/context.go:469
		// _ = "end of CoverTab[2370]"
	} else {
//line /usr/local/go/src/context/context.go:470
		_go_fuzz_dep_.CoverTab[2371]++
								close(d)
//line /usr/local/go/src/context/context.go:471
		// _ = "end of CoverTab[2371]"
	}
//line /usr/local/go/src/context/context.go:472
	// _ = "end of CoverTab[2361]"
//line /usr/local/go/src/context/context.go:472
	_go_fuzz_dep_.CoverTab[2362]++
							for child := range c.children {
//line /usr/local/go/src/context/context.go:473
		_go_fuzz_dep_.CoverTab[2372]++

								child.cancel(false, err, cause)
//line /usr/local/go/src/context/context.go:475
		// _ = "end of CoverTab[2372]"
	}
//line /usr/local/go/src/context/context.go:476
	// _ = "end of CoverTab[2362]"
//line /usr/local/go/src/context/context.go:476
	_go_fuzz_dep_.CoverTab[2363]++
							c.children = nil
							c.mu.Unlock()

							if removeFromParent {
//line /usr/local/go/src/context/context.go:480
		_go_fuzz_dep_.CoverTab[2373]++
								removeChild(c.Context, c)
//line /usr/local/go/src/context/context.go:481
		// _ = "end of CoverTab[2373]"
	} else {
//line /usr/local/go/src/context/context.go:482
		_go_fuzz_dep_.CoverTab[2374]++
//line /usr/local/go/src/context/context.go:482
		// _ = "end of CoverTab[2374]"
//line /usr/local/go/src/context/context.go:482
	}
//line /usr/local/go/src/context/context.go:482
	// _ = "end of CoverTab[2363]"
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
//line /usr/local/go/src/context/context.go:485
// to be no later than d. If the parent's deadline is already earlier than d,
//line /usr/local/go/src/context/context.go:485
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
//line /usr/local/go/src/context/context.go:485
// context's Done channel is closed when the deadline expires, when the returned
//line /usr/local/go/src/context/context.go:485
// cancel function is called, or when the parent context's Done channel is
//line /usr/local/go/src/context/context.go:485
// closed, whichever happens first.
//line /usr/local/go/src/context/context.go:485
//
//line /usr/local/go/src/context/context.go:485
// Canceling this context releases resources associated with it, so code should
//line /usr/local/go/src/context/context.go:485
// call cancel as soon as the operations running in this Context complete.
//line /usr/local/go/src/context/context.go:494
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
//line /usr/local/go/src/context/context.go:494
	_go_fuzz_dep_.CoverTab[2375]++
							if parent == nil {
//line /usr/local/go/src/context/context.go:495
		_go_fuzz_dep_.CoverTab[2380]++
								panic("cannot create context from nil parent")
//line /usr/local/go/src/context/context.go:496
		// _ = "end of CoverTab[2380]"
	} else {
//line /usr/local/go/src/context/context.go:497
		_go_fuzz_dep_.CoverTab[2381]++
//line /usr/local/go/src/context/context.go:497
		// _ = "end of CoverTab[2381]"
//line /usr/local/go/src/context/context.go:497
	}
//line /usr/local/go/src/context/context.go:497
	// _ = "end of CoverTab[2375]"
//line /usr/local/go/src/context/context.go:497
	_go_fuzz_dep_.CoverTab[2376]++
							if cur, ok := parent.Deadline(); ok && func() bool {
//line /usr/local/go/src/context/context.go:498
		_go_fuzz_dep_.CoverTab[2382]++
//line /usr/local/go/src/context/context.go:498
		return cur.Before(d)
//line /usr/local/go/src/context/context.go:498
		// _ = "end of CoverTab[2382]"
//line /usr/local/go/src/context/context.go:498
	}() {
//line /usr/local/go/src/context/context.go:498
		_go_fuzz_dep_.CoverTab[2383]++

								return WithCancel(parent)
//line /usr/local/go/src/context/context.go:500
		// _ = "end of CoverTab[2383]"
	} else {
//line /usr/local/go/src/context/context.go:501
		_go_fuzz_dep_.CoverTab[2384]++
//line /usr/local/go/src/context/context.go:501
		// _ = "end of CoverTab[2384]"
//line /usr/local/go/src/context/context.go:501
	}
//line /usr/local/go/src/context/context.go:501
	// _ = "end of CoverTab[2376]"
//line /usr/local/go/src/context/context.go:501
	_go_fuzz_dep_.CoverTab[2377]++
							c := &timerCtx{
		cancelCtx:	newCancelCtx(parent),
		deadline:	d,
	}
	propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
//line /usr/local/go/src/context/context.go:508
		_go_fuzz_dep_.CoverTab[2385]++
								c.cancel(true, DeadlineExceeded, nil)
								return c, func() { _go_fuzz_dep_.CoverTab[2386]++; c.cancel(false, Canceled, nil); // _ = "end of CoverTab[2386]" }
//line /usr/local/go/src/context/context.go:510
		// _ = "end of CoverTab[2385]"
	} else {
//line /usr/local/go/src/context/context.go:511
		_go_fuzz_dep_.CoverTab[2387]++
//line /usr/local/go/src/context/context.go:511
		// _ = "end of CoverTab[2387]"
//line /usr/local/go/src/context/context.go:511
	}
//line /usr/local/go/src/context/context.go:511
	// _ = "end of CoverTab[2377]"
//line /usr/local/go/src/context/context.go:511
	_go_fuzz_dep_.CoverTab[2378]++
							c.mu.Lock()
							defer c.mu.Unlock()
							if c.err == nil {
//line /usr/local/go/src/context/context.go:514
		_go_fuzz_dep_.CoverTab[2388]++
								c.timer = time.AfterFunc(dur, func() {
//line /usr/local/go/src/context/context.go:515
			_go_fuzz_dep_.CoverTab[2389]++
									c.cancel(true, DeadlineExceeded, nil)
//line /usr/local/go/src/context/context.go:516
			// _ = "end of CoverTab[2389]"
		})
//line /usr/local/go/src/context/context.go:517
		// _ = "end of CoverTab[2388]"
	} else {
//line /usr/local/go/src/context/context.go:518
		_go_fuzz_dep_.CoverTab[2390]++
//line /usr/local/go/src/context/context.go:518
		// _ = "end of CoverTab[2390]"
//line /usr/local/go/src/context/context.go:518
	}
//line /usr/local/go/src/context/context.go:518
	// _ = "end of CoverTab[2378]"
//line /usr/local/go/src/context/context.go:518
	_go_fuzz_dep_.CoverTab[2379]++
							return c, func() { _go_fuzz_dep_.CoverTab[2391]++; c.cancel(true, Canceled, nil); // _ = "end of CoverTab[2391]" }
//line /usr/local/go/src/context/context.go:519
	// _ = "end of CoverTab[2379]"
}

// A timerCtx carries a timer and a deadline. It embeds a cancelCtx to
//line /usr/local/go/src/context/context.go:522
// implement Done and Err. It implements cancel by stopping its timer then
//line /usr/local/go/src/context/context.go:522
// delegating to cancelCtx.cancel.
//line /usr/local/go/src/context/context.go:525
type timerCtx struct {
	*cancelCtx
	timer	*time.Timer	// Under cancelCtx.mu.

	deadline	time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
//line /usr/local/go/src/context/context.go:532
	_go_fuzz_dep_.CoverTab[2392]++
							return c.deadline, true
//line /usr/local/go/src/context/context.go:533
	// _ = "end of CoverTab[2392]"
}

func (c *timerCtx) String() string {
//line /usr/local/go/src/context/context.go:536
	_go_fuzz_dep_.CoverTab[2393]++
							return contextName(c.cancelCtx.Context) + ".WithDeadline(" +
		c.deadline.String() + " [" +
		time.Until(c.deadline).String() + "])"
//line /usr/local/go/src/context/context.go:539
	// _ = "end of CoverTab[2393]"
}

func (c *timerCtx) cancel(removeFromParent bool, err, cause error) {
//line /usr/local/go/src/context/context.go:542
	_go_fuzz_dep_.CoverTab[2394]++
							c.cancelCtx.cancel(false, err, cause)
							if removeFromParent {
//line /usr/local/go/src/context/context.go:544
		_go_fuzz_dep_.CoverTab[2397]++

								removeChild(c.cancelCtx.Context, c)
//line /usr/local/go/src/context/context.go:546
		// _ = "end of CoverTab[2397]"
	} else {
//line /usr/local/go/src/context/context.go:547
		_go_fuzz_dep_.CoverTab[2398]++
//line /usr/local/go/src/context/context.go:547
		// _ = "end of CoverTab[2398]"
//line /usr/local/go/src/context/context.go:547
	}
//line /usr/local/go/src/context/context.go:547
	// _ = "end of CoverTab[2394]"
//line /usr/local/go/src/context/context.go:547
	_go_fuzz_dep_.CoverTab[2395]++
							c.mu.Lock()
							if c.timer != nil {
//line /usr/local/go/src/context/context.go:549
		_go_fuzz_dep_.CoverTab[2399]++
								c.timer.Stop()
								c.timer = nil
//line /usr/local/go/src/context/context.go:551
		// _ = "end of CoverTab[2399]"
	} else {
//line /usr/local/go/src/context/context.go:552
		_go_fuzz_dep_.CoverTab[2400]++
//line /usr/local/go/src/context/context.go:552
		// _ = "end of CoverTab[2400]"
//line /usr/local/go/src/context/context.go:552
	}
//line /usr/local/go/src/context/context.go:552
	// _ = "end of CoverTab[2395]"
//line /usr/local/go/src/context/context.go:552
	_go_fuzz_dep_.CoverTab[2396]++
							c.mu.Unlock()
//line /usr/local/go/src/context/context.go:553
	// _ = "end of CoverTab[2396]"
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//line /usr/local/go/src/context/context.go:556
//
//line /usr/local/go/src/context/context.go:556
// Canceling this context releases resources associated with it, so code should
//line /usr/local/go/src/context/context.go:556
// call cancel as soon as the operations running in this Context complete:
//line /usr/local/go/src/context/context.go:556
//
//line /usr/local/go/src/context/context.go:556
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//line /usr/local/go/src/context/context.go:556
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//line /usr/local/go/src/context/context.go:556
//		defer cancel()  // releases resources if slowOperation completes before timeout elapses
//line /usr/local/go/src/context/context.go:556
//		return slowOperation(ctx)
//line /usr/local/go/src/context/context.go:556
//	}
//line /usr/local/go/src/context/context.go:566
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
//line /usr/local/go/src/context/context.go:566
	_go_fuzz_dep_.CoverTab[2401]++
							return WithDeadline(parent, time.Now().Add(timeout))
//line /usr/local/go/src/context/context.go:567
	// _ = "end of CoverTab[2401]"
}

// WithValue returns a copy of parent in which the value associated with key is
//line /usr/local/go/src/context/context.go:570
// val.
//line /usr/local/go/src/context/context.go:570
//
//line /usr/local/go/src/context/context.go:570
// Use context Values only for request-scoped data that transits processes and
//line /usr/local/go/src/context/context.go:570
// APIs, not for passing optional parameters to functions.
//line /usr/local/go/src/context/context.go:570
//
//line /usr/local/go/src/context/context.go:570
// The provided key must be comparable and should not be of type
//line /usr/local/go/src/context/context.go:570
// string or any other built-in type to avoid collisions between
//line /usr/local/go/src/context/context.go:570
// packages using context. Users of WithValue should define their own
//line /usr/local/go/src/context/context.go:570
// types for keys. To avoid allocating when assigning to an
//line /usr/local/go/src/context/context.go:570
// interface{}, context keys often have concrete type
//line /usr/local/go/src/context/context.go:570
// struct{}. Alternatively, exported context key variables' static
//line /usr/local/go/src/context/context.go:570
// type should be a pointer or interface.
//line /usr/local/go/src/context/context.go:583
func WithValue(parent Context, key, val any) Context {
//line /usr/local/go/src/context/context.go:583
	_go_fuzz_dep_.CoverTab[2402]++
							if parent == nil {
//line /usr/local/go/src/context/context.go:584
		_go_fuzz_dep_.CoverTab[2406]++
								panic("cannot create context from nil parent")
//line /usr/local/go/src/context/context.go:585
		// _ = "end of CoverTab[2406]"
	} else {
//line /usr/local/go/src/context/context.go:586
		_go_fuzz_dep_.CoverTab[2407]++
//line /usr/local/go/src/context/context.go:586
		// _ = "end of CoverTab[2407]"
//line /usr/local/go/src/context/context.go:586
	}
//line /usr/local/go/src/context/context.go:586
	// _ = "end of CoverTab[2402]"
//line /usr/local/go/src/context/context.go:586
	_go_fuzz_dep_.CoverTab[2403]++
							if key == nil {
//line /usr/local/go/src/context/context.go:587
		_go_fuzz_dep_.CoverTab[2408]++
								panic("nil key")
//line /usr/local/go/src/context/context.go:588
		// _ = "end of CoverTab[2408]"
	} else {
//line /usr/local/go/src/context/context.go:589
		_go_fuzz_dep_.CoverTab[2409]++
//line /usr/local/go/src/context/context.go:589
		// _ = "end of CoverTab[2409]"
//line /usr/local/go/src/context/context.go:589
	}
//line /usr/local/go/src/context/context.go:589
	// _ = "end of CoverTab[2403]"
//line /usr/local/go/src/context/context.go:589
	_go_fuzz_dep_.CoverTab[2404]++
							if !reflectlite.TypeOf(key).Comparable() {
//line /usr/local/go/src/context/context.go:590
		_go_fuzz_dep_.CoverTab[2410]++
								panic("key is not comparable")
//line /usr/local/go/src/context/context.go:591
		// _ = "end of CoverTab[2410]"
	} else {
//line /usr/local/go/src/context/context.go:592
		_go_fuzz_dep_.CoverTab[2411]++
//line /usr/local/go/src/context/context.go:592
		// _ = "end of CoverTab[2411]"
//line /usr/local/go/src/context/context.go:592
	}
//line /usr/local/go/src/context/context.go:592
	// _ = "end of CoverTab[2404]"
//line /usr/local/go/src/context/context.go:592
	_go_fuzz_dep_.CoverTab[2405]++
							return &valueCtx{parent, key, val}
//line /usr/local/go/src/context/context.go:593
	// _ = "end of CoverTab[2405]"
}

// A valueCtx carries a key-value pair. It implements Value for that key and
//line /usr/local/go/src/context/context.go:596
// delegates all other calls to the embedded Context.
//line /usr/local/go/src/context/context.go:598
type valueCtx struct {
	Context
	key, val	any
}

// stringify tries a bit to stringify v, without using fmt, since we don't
//line /usr/local/go/src/context/context.go:603
// want context depending on the unicode tables. This is only used by
//line /usr/local/go/src/context/context.go:603
// *valueCtx.String().
//line /usr/local/go/src/context/context.go:606
func stringify(v any) string {
//line /usr/local/go/src/context/context.go:606
	_go_fuzz_dep_.CoverTab[2412]++
							switch s := v.(type) {
	case stringer:
//line /usr/local/go/src/context/context.go:608
		_go_fuzz_dep_.CoverTab[2414]++
								return s.String()
//line /usr/local/go/src/context/context.go:609
		// _ = "end of CoverTab[2414]"
	case string:
//line /usr/local/go/src/context/context.go:610
		_go_fuzz_dep_.CoverTab[2415]++
								return s
//line /usr/local/go/src/context/context.go:611
		// _ = "end of CoverTab[2415]"
	}
//line /usr/local/go/src/context/context.go:612
	// _ = "end of CoverTab[2412]"
//line /usr/local/go/src/context/context.go:612
	_go_fuzz_dep_.CoverTab[2413]++
							return "<not Stringer>"
//line /usr/local/go/src/context/context.go:613
	// _ = "end of CoverTab[2413]"
}

func (c *valueCtx) String() string {
//line /usr/local/go/src/context/context.go:616
	_go_fuzz_dep_.CoverTab[2416]++
							return contextName(c.Context) + ".WithValue(type " +
		reflectlite.TypeOf(c.key).String() +
		", val " + stringify(c.val) + ")"
//line /usr/local/go/src/context/context.go:619
	// _ = "end of CoverTab[2416]"
}

func (c *valueCtx) Value(key any) any {
//line /usr/local/go/src/context/context.go:622
	_go_fuzz_dep_.CoverTab[2417]++
							if c.key == key {
//line /usr/local/go/src/context/context.go:623
		_go_fuzz_dep_.CoverTab[2419]++
								return c.val
//line /usr/local/go/src/context/context.go:624
		// _ = "end of CoverTab[2419]"
	} else {
//line /usr/local/go/src/context/context.go:625
		_go_fuzz_dep_.CoverTab[2420]++
//line /usr/local/go/src/context/context.go:625
		// _ = "end of CoverTab[2420]"
//line /usr/local/go/src/context/context.go:625
	}
//line /usr/local/go/src/context/context.go:625
	// _ = "end of CoverTab[2417]"
//line /usr/local/go/src/context/context.go:625
	_go_fuzz_dep_.CoverTab[2418]++
							return value(c.Context, key)
//line /usr/local/go/src/context/context.go:626
	// _ = "end of CoverTab[2418]"
}

func value(c Context, key any) any {
//line /usr/local/go/src/context/context.go:629
	_go_fuzz_dep_.CoverTab[2421]++
							for {
//line /usr/local/go/src/context/context.go:630
		_go_fuzz_dep_.CoverTab[2422]++
								switch ctx := c.(type) {
		case *valueCtx:
//line /usr/local/go/src/context/context.go:632
			_go_fuzz_dep_.CoverTab[2423]++
									if key == ctx.key {
//line /usr/local/go/src/context/context.go:633
				_go_fuzz_dep_.CoverTab[2431]++
										return ctx.val
//line /usr/local/go/src/context/context.go:634
				// _ = "end of CoverTab[2431]"
			} else {
//line /usr/local/go/src/context/context.go:635
				_go_fuzz_dep_.CoverTab[2432]++
//line /usr/local/go/src/context/context.go:635
				// _ = "end of CoverTab[2432]"
//line /usr/local/go/src/context/context.go:635
			}
//line /usr/local/go/src/context/context.go:635
			// _ = "end of CoverTab[2423]"
//line /usr/local/go/src/context/context.go:635
			_go_fuzz_dep_.CoverTab[2424]++
									c = ctx.Context
//line /usr/local/go/src/context/context.go:636
			// _ = "end of CoverTab[2424]"
		case *cancelCtx:
//line /usr/local/go/src/context/context.go:637
			_go_fuzz_dep_.CoverTab[2425]++
									if key == &cancelCtxKey {
//line /usr/local/go/src/context/context.go:638
				_go_fuzz_dep_.CoverTab[2433]++
										return c
//line /usr/local/go/src/context/context.go:639
				// _ = "end of CoverTab[2433]"
			} else {
//line /usr/local/go/src/context/context.go:640
				_go_fuzz_dep_.CoverTab[2434]++
//line /usr/local/go/src/context/context.go:640
				// _ = "end of CoverTab[2434]"
//line /usr/local/go/src/context/context.go:640
			}
//line /usr/local/go/src/context/context.go:640
			// _ = "end of CoverTab[2425]"
//line /usr/local/go/src/context/context.go:640
			_go_fuzz_dep_.CoverTab[2426]++
									c = ctx.Context
//line /usr/local/go/src/context/context.go:641
			// _ = "end of CoverTab[2426]"
		case *timerCtx:
//line /usr/local/go/src/context/context.go:642
			_go_fuzz_dep_.CoverTab[2427]++
									if key == &cancelCtxKey {
//line /usr/local/go/src/context/context.go:643
				_go_fuzz_dep_.CoverTab[2435]++
										return ctx.cancelCtx
//line /usr/local/go/src/context/context.go:644
				// _ = "end of CoverTab[2435]"
			} else {
//line /usr/local/go/src/context/context.go:645
				_go_fuzz_dep_.CoverTab[2436]++
//line /usr/local/go/src/context/context.go:645
				// _ = "end of CoverTab[2436]"
//line /usr/local/go/src/context/context.go:645
			}
//line /usr/local/go/src/context/context.go:645
			// _ = "end of CoverTab[2427]"
//line /usr/local/go/src/context/context.go:645
			_go_fuzz_dep_.CoverTab[2428]++
									c = ctx.Context
//line /usr/local/go/src/context/context.go:646
			// _ = "end of CoverTab[2428]"
		case *emptyCtx:
//line /usr/local/go/src/context/context.go:647
			_go_fuzz_dep_.CoverTab[2429]++
									return nil
//line /usr/local/go/src/context/context.go:648
			// _ = "end of CoverTab[2429]"
		default:
//line /usr/local/go/src/context/context.go:649
			_go_fuzz_dep_.CoverTab[2430]++
									return c.Value(key)
//line /usr/local/go/src/context/context.go:650
			// _ = "end of CoverTab[2430]"
		}
//line /usr/local/go/src/context/context.go:651
		// _ = "end of CoverTab[2422]"
	}
//line /usr/local/go/src/context/context.go:652
	// _ = "end of CoverTab[2421]"
}

//line /usr/local/go/src/context/context.go:653
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/context/context.go:653
var _ = _go_fuzz_dep_.CoverTab
