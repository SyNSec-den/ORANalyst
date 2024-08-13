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
	_go_fuzz_dep_.CoverTab[935]++
//line /usr/local/go/src/context/context.go:171
	return "context deadline exceeded"
//line /usr/local/go/src/context/context.go:171
	// _ = "end of CoverTab[935]"
//line /usr/local/go/src/context/context.go:171
}
func (deadlineExceededError) Timeout() bool {
//line /usr/local/go/src/context/context.go:172
	_go_fuzz_dep_.CoverTab[936]++
//line /usr/local/go/src/context/context.go:172
	return true
//line /usr/local/go/src/context/context.go:172
	// _ = "end of CoverTab[936]"
//line /usr/local/go/src/context/context.go:172
}
func (deadlineExceededError) Temporary() bool {
//line /usr/local/go/src/context/context.go:173
	_go_fuzz_dep_.CoverTab[937]++
//line /usr/local/go/src/context/context.go:173
	return true
//line /usr/local/go/src/context/context.go:173
	// _ = "end of CoverTab[937]"
//line /usr/local/go/src/context/context.go:173
}

// An emptyCtx is never canceled, has no values, and has no deadline. It is not
//line /usr/local/go/src/context/context.go:175
// struct{}, since vars of this type must have distinct addresses.
//line /usr/local/go/src/context/context.go:177
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
//line /usr/local/go/src/context/context.go:179
	_go_fuzz_dep_.CoverTab[938]++
							return
//line /usr/local/go/src/context/context.go:180
	// _ = "end of CoverTab[938]"
}

func (*emptyCtx) Done() <-chan struct{} {
//line /usr/local/go/src/context/context.go:183
	_go_fuzz_dep_.CoverTab[939]++
							return nil
//line /usr/local/go/src/context/context.go:184
	// _ = "end of CoverTab[939]"
}

func (*emptyCtx) Err() error {
//line /usr/local/go/src/context/context.go:187
	_go_fuzz_dep_.CoverTab[940]++
							return nil
//line /usr/local/go/src/context/context.go:188
	// _ = "end of CoverTab[940]"
}

func (*emptyCtx) Value(key any) any {
//line /usr/local/go/src/context/context.go:191
	_go_fuzz_dep_.CoverTab[941]++
							return nil
//line /usr/local/go/src/context/context.go:192
	// _ = "end of CoverTab[941]"
}

func (e *emptyCtx) String() string {
//line /usr/local/go/src/context/context.go:195
	_go_fuzz_dep_.CoverTab[942]++
							switch e {
	case background:
//line /usr/local/go/src/context/context.go:197
		_go_fuzz_dep_.CoverTab[944]++
								return "context.Background"
//line /usr/local/go/src/context/context.go:198
		// _ = "end of CoverTab[944]"
	case todo:
//line /usr/local/go/src/context/context.go:199
		_go_fuzz_dep_.CoverTab[945]++
								return "context.TODO"
//line /usr/local/go/src/context/context.go:200
		// _ = "end of CoverTab[945]"
//line /usr/local/go/src/context/context.go:200
	default:
//line /usr/local/go/src/context/context.go:200
		_go_fuzz_dep_.CoverTab[946]++
//line /usr/local/go/src/context/context.go:200
		// _ = "end of CoverTab[946]"
	}
//line /usr/local/go/src/context/context.go:201
	// _ = "end of CoverTab[942]"
//line /usr/local/go/src/context/context.go:201
	_go_fuzz_dep_.CoverTab[943]++
							return "unknown empty Context"
//line /usr/local/go/src/context/context.go:202
	// _ = "end of CoverTab[943]"
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
	_go_fuzz_dep_.CoverTab[947]++
							return background
//line /usr/local/go/src/context/context.go:215
	// _ = "end of CoverTab[947]"
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
	_go_fuzz_dep_.CoverTab[948]++
							return todo
//line /usr/local/go/src/context/context.go:223
	// _ = "end of CoverTab[948]"
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
	_go_fuzz_dep_.CoverTab[949]++
							c := withCancel(parent)
							return c, func() { _go_fuzz_dep_.CoverTab[950]++; c.cancel(true, Canceled, nil); // _ = "end of CoverTab[950]" }
//line /usr/local/go/src/context/context.go:240
	// _ = "end of CoverTab[949]"
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
	_go_fuzz_dep_.CoverTab[951]++
							c := withCancel(parent)
							return c, func(cause error) {
//line /usr/local/go/src/context/context.go:268
		_go_fuzz_dep_.CoverTab[952]++
//line /usr/local/go/src/context/context.go:268
		c.cancel(true, Canceled, cause)
//line /usr/local/go/src/context/context.go:268
		// _ = "end of CoverTab[952]"
//line /usr/local/go/src/context/context.go:268
	}
//line /usr/local/go/src/context/context.go:268
	// _ = "end of CoverTab[951]"
}

func withCancel(parent Context) *cancelCtx {
//line /usr/local/go/src/context/context.go:271
	_go_fuzz_dep_.CoverTab[953]++
							if parent == nil {
//line /usr/local/go/src/context/context.go:272
		_go_fuzz_dep_.CoverTab[955]++
								panic("cannot create context from nil parent")
//line /usr/local/go/src/context/context.go:273
		// _ = "end of CoverTab[955]"
	} else {
//line /usr/local/go/src/context/context.go:274
		_go_fuzz_dep_.CoverTab[956]++
//line /usr/local/go/src/context/context.go:274
		// _ = "end of CoverTab[956]"
//line /usr/local/go/src/context/context.go:274
	}
//line /usr/local/go/src/context/context.go:274
	// _ = "end of CoverTab[953]"
//line /usr/local/go/src/context/context.go:274
	_go_fuzz_dep_.CoverTab[954]++
							c := newCancelCtx(parent)
							propagateCancel(parent, c)
							return c
//line /usr/local/go/src/context/context.go:277
	// _ = "end of CoverTab[954]"
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
	_go_fuzz_dep_.CoverTab[957]++
							if cc, ok := c.Value(&cancelCtxKey).(*cancelCtx); ok {
//line /usr/local/go/src/context/context.go:287
		_go_fuzz_dep_.CoverTab[959]++
								cc.mu.Lock()
								defer cc.mu.Unlock()
								return cc.cause
//line /usr/local/go/src/context/context.go:290
		// _ = "end of CoverTab[959]"
	} else {
//line /usr/local/go/src/context/context.go:291
		_go_fuzz_dep_.CoverTab[960]++
//line /usr/local/go/src/context/context.go:291
		// _ = "end of CoverTab[960]"
//line /usr/local/go/src/context/context.go:291
	}
//line /usr/local/go/src/context/context.go:291
	// _ = "end of CoverTab[957]"
//line /usr/local/go/src/context/context.go:291
	_go_fuzz_dep_.CoverTab[958]++
							return nil
//line /usr/local/go/src/context/context.go:292
	// _ = "end of CoverTab[958]"
}

// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) *cancelCtx {
//line /usr/local/go/src/context/context.go:296
	_go_fuzz_dep_.CoverTab[961]++
							return &cancelCtx{Context: parent}
//line /usr/local/go/src/context/context.go:297
	// _ = "end of CoverTab[961]"
}

// goroutines counts the number of goroutines ever created; for testing.
var goroutines atomic.Int32

// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
//line /usr/local/go/src/context/context.go:304
	_go_fuzz_dep_.CoverTab[962]++
							done := parent.Done()
							if done == nil {
//line /usr/local/go/src/context/context.go:306
		_go_fuzz_dep_.CoverTab[965]++
								return
//line /usr/local/go/src/context/context.go:307
		// _ = "end of CoverTab[965]"
	} else {
//line /usr/local/go/src/context/context.go:308
		_go_fuzz_dep_.CoverTab[966]++
//line /usr/local/go/src/context/context.go:308
		// _ = "end of CoverTab[966]"
//line /usr/local/go/src/context/context.go:308
	}
//line /usr/local/go/src/context/context.go:308
	// _ = "end of CoverTab[962]"
//line /usr/local/go/src/context/context.go:308
	_go_fuzz_dep_.CoverTab[963]++

							select {
	case <-done:
//line /usr/local/go/src/context/context.go:311
		_go_fuzz_dep_.CoverTab[967]++

								child.cancel(false, parent.Err(), Cause(parent))
								return
//line /usr/local/go/src/context/context.go:314
		// _ = "end of CoverTab[967]"
	default:
//line /usr/local/go/src/context/context.go:315
		_go_fuzz_dep_.CoverTab[968]++
//line /usr/local/go/src/context/context.go:315
		// _ = "end of CoverTab[968]"
	}
//line /usr/local/go/src/context/context.go:316
	// _ = "end of CoverTab[963]"
//line /usr/local/go/src/context/context.go:316
	_go_fuzz_dep_.CoverTab[964]++

							if p, ok := parentCancelCtx(parent); ok {
//line /usr/local/go/src/context/context.go:318
		_go_fuzz_dep_.CoverTab[969]++
								p.mu.Lock()
								if p.err != nil {
//line /usr/local/go/src/context/context.go:320
			_go_fuzz_dep_.CoverTab[971]++

									child.cancel(false, p.err, p.cause)
//line /usr/local/go/src/context/context.go:322
			// _ = "end of CoverTab[971]"
		} else {
//line /usr/local/go/src/context/context.go:323
			_go_fuzz_dep_.CoverTab[972]++
									if p.children == nil {
//line /usr/local/go/src/context/context.go:324
				_go_fuzz_dep_.CoverTab[974]++
										p.children = make(map[canceler]struct{})
//line /usr/local/go/src/context/context.go:325
				// _ = "end of CoverTab[974]"
			} else {
//line /usr/local/go/src/context/context.go:326
				_go_fuzz_dep_.CoverTab[975]++
//line /usr/local/go/src/context/context.go:326
				// _ = "end of CoverTab[975]"
//line /usr/local/go/src/context/context.go:326
			}
//line /usr/local/go/src/context/context.go:326
			// _ = "end of CoverTab[972]"
//line /usr/local/go/src/context/context.go:326
			_go_fuzz_dep_.CoverTab[973]++
									p.children[child] = struct{}{}
//line /usr/local/go/src/context/context.go:327
			// _ = "end of CoverTab[973]"
		}
//line /usr/local/go/src/context/context.go:328
		// _ = "end of CoverTab[969]"
//line /usr/local/go/src/context/context.go:328
		_go_fuzz_dep_.CoverTab[970]++
								p.mu.Unlock()
//line /usr/local/go/src/context/context.go:329
		// _ = "end of CoverTab[970]"
	} else {
//line /usr/local/go/src/context/context.go:330
		_go_fuzz_dep_.CoverTab[976]++
								goroutines.Add(1)
//line /usr/local/go/src/context/context.go:331
		_curRoutineNum0_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/context/context.go:331
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum0_)
								go func() {
//line /usr/local/go/src/context/context.go:332
			_go_fuzz_dep_.CoverTab[977]++
//line /usr/local/go/src/context/context.go:332
			defer func() {
//line /usr/local/go/src/context/context.go:332
				_go_fuzz_dep_.CoverTab[978]++
//line /usr/local/go/src/context/context.go:332
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum0_)
//line /usr/local/go/src/context/context.go:332
				// _ = "end of CoverTab[978]"
//line /usr/local/go/src/context/context.go:332
			}()
									select {
			case <-parent.Done():
//line /usr/local/go/src/context/context.go:334
				_go_fuzz_dep_.CoverTab[979]++
										child.cancel(false, parent.Err(), Cause(parent))
//line /usr/local/go/src/context/context.go:335
				// _ = "end of CoverTab[979]"
			case <-child.Done():
//line /usr/local/go/src/context/context.go:336
				_go_fuzz_dep_.CoverTab[980]++
//line /usr/local/go/src/context/context.go:336
				// _ = "end of CoverTab[980]"
			}
//line /usr/local/go/src/context/context.go:337
			// _ = "end of CoverTab[977]"
		}()
//line /usr/local/go/src/context/context.go:338
		// _ = "end of CoverTab[976]"
	}
//line /usr/local/go/src/context/context.go:339
	// _ = "end of CoverTab[964]"
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
	_go_fuzz_dep_.CoverTab[981]++
							done := parent.Done()
							if done == closedchan || func() bool {
//line /usr/local/go/src/context/context.go:353
		_go_fuzz_dep_.CoverTab[985]++
//line /usr/local/go/src/context/context.go:353
		return done == nil
//line /usr/local/go/src/context/context.go:353
		// _ = "end of CoverTab[985]"
//line /usr/local/go/src/context/context.go:353
	}() {
//line /usr/local/go/src/context/context.go:353
		_go_fuzz_dep_.CoverTab[986]++
								return nil, false
//line /usr/local/go/src/context/context.go:354
		// _ = "end of CoverTab[986]"
	} else {
//line /usr/local/go/src/context/context.go:355
		_go_fuzz_dep_.CoverTab[987]++
//line /usr/local/go/src/context/context.go:355
		// _ = "end of CoverTab[987]"
//line /usr/local/go/src/context/context.go:355
	}
//line /usr/local/go/src/context/context.go:355
	// _ = "end of CoverTab[981]"
//line /usr/local/go/src/context/context.go:355
	_go_fuzz_dep_.CoverTab[982]++
							p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
							if !ok {
//line /usr/local/go/src/context/context.go:357
		_go_fuzz_dep_.CoverTab[988]++
								return nil, false
//line /usr/local/go/src/context/context.go:358
		// _ = "end of CoverTab[988]"
	} else {
//line /usr/local/go/src/context/context.go:359
		_go_fuzz_dep_.CoverTab[989]++
//line /usr/local/go/src/context/context.go:359
		// _ = "end of CoverTab[989]"
//line /usr/local/go/src/context/context.go:359
	}
//line /usr/local/go/src/context/context.go:359
	// _ = "end of CoverTab[982]"
//line /usr/local/go/src/context/context.go:359
	_go_fuzz_dep_.CoverTab[983]++
							pdone, _ := p.done.Load().(chan struct{})
							if pdone != done {
//line /usr/local/go/src/context/context.go:361
		_go_fuzz_dep_.CoverTab[990]++
								return nil, false
//line /usr/local/go/src/context/context.go:362
		// _ = "end of CoverTab[990]"
	} else {
//line /usr/local/go/src/context/context.go:363
		_go_fuzz_dep_.CoverTab[991]++
//line /usr/local/go/src/context/context.go:363
		// _ = "end of CoverTab[991]"
//line /usr/local/go/src/context/context.go:363
	}
//line /usr/local/go/src/context/context.go:363
	// _ = "end of CoverTab[983]"
//line /usr/local/go/src/context/context.go:363
	_go_fuzz_dep_.CoverTab[984]++
							return p, true
//line /usr/local/go/src/context/context.go:364
	// _ = "end of CoverTab[984]"
}

// removeChild removes a context from its parent.
func removeChild(parent Context, child canceler) {
//line /usr/local/go/src/context/context.go:368
	_go_fuzz_dep_.CoverTab[992]++
							p, ok := parentCancelCtx(parent)
							if !ok {
//line /usr/local/go/src/context/context.go:370
		_go_fuzz_dep_.CoverTab[995]++
								return
//line /usr/local/go/src/context/context.go:371
		// _ = "end of CoverTab[995]"
	} else {
//line /usr/local/go/src/context/context.go:372
		_go_fuzz_dep_.CoverTab[996]++
//line /usr/local/go/src/context/context.go:372
		// _ = "end of CoverTab[996]"
//line /usr/local/go/src/context/context.go:372
	}
//line /usr/local/go/src/context/context.go:372
	// _ = "end of CoverTab[992]"
//line /usr/local/go/src/context/context.go:372
	_go_fuzz_dep_.CoverTab[993]++
							p.mu.Lock()
							if p.children != nil {
//line /usr/local/go/src/context/context.go:374
		_go_fuzz_dep_.CoverTab[997]++
								delete(p.children, child)
//line /usr/local/go/src/context/context.go:375
		// _ = "end of CoverTab[997]"
	} else {
//line /usr/local/go/src/context/context.go:376
		_go_fuzz_dep_.CoverTab[998]++
//line /usr/local/go/src/context/context.go:376
		// _ = "end of CoverTab[998]"
//line /usr/local/go/src/context/context.go:376
	}
//line /usr/local/go/src/context/context.go:376
	// _ = "end of CoverTab[993]"
//line /usr/local/go/src/context/context.go:376
	_go_fuzz_dep_.CoverTab[994]++
							p.mu.Unlock()
//line /usr/local/go/src/context/context.go:377
	// _ = "end of CoverTab[994]"
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
	_go_fuzz_dep_.CoverTab[999]++
							if key == &cancelCtxKey {
//line /usr/local/go/src/context/context.go:407
		_go_fuzz_dep_.CoverTab[1001]++
								return c
//line /usr/local/go/src/context/context.go:408
		// _ = "end of CoverTab[1001]"
	} else {
//line /usr/local/go/src/context/context.go:409
		_go_fuzz_dep_.CoverTab[1002]++
//line /usr/local/go/src/context/context.go:409
		// _ = "end of CoverTab[1002]"
//line /usr/local/go/src/context/context.go:409
	}
//line /usr/local/go/src/context/context.go:409
	// _ = "end of CoverTab[999]"
//line /usr/local/go/src/context/context.go:409
	_go_fuzz_dep_.CoverTab[1000]++
							return value(c.Context, key)
//line /usr/local/go/src/context/context.go:410
	// _ = "end of CoverTab[1000]"
}

func (c *cancelCtx) Done() <-chan struct{} {
//line /usr/local/go/src/context/context.go:413
	_go_fuzz_dep_.CoverTab[1003]++
							d := c.done.Load()
							if d != nil {
//line /usr/local/go/src/context/context.go:415
		_go_fuzz_dep_.CoverTab[1006]++
								return d.(chan struct{})
//line /usr/local/go/src/context/context.go:416
		// _ = "end of CoverTab[1006]"
	} else {
//line /usr/local/go/src/context/context.go:417
		_go_fuzz_dep_.CoverTab[1007]++
//line /usr/local/go/src/context/context.go:417
		// _ = "end of CoverTab[1007]"
//line /usr/local/go/src/context/context.go:417
	}
//line /usr/local/go/src/context/context.go:417
	// _ = "end of CoverTab[1003]"
//line /usr/local/go/src/context/context.go:417
	_go_fuzz_dep_.CoverTab[1004]++
							c.mu.Lock()
							defer c.mu.Unlock()
							d = c.done.Load()
							if d == nil {
//line /usr/local/go/src/context/context.go:421
		_go_fuzz_dep_.CoverTab[1008]++
								d = make(chan struct{})
								c.done.Store(d)
//line /usr/local/go/src/context/context.go:423
		// _ = "end of CoverTab[1008]"
	} else {
//line /usr/local/go/src/context/context.go:424
		_go_fuzz_dep_.CoverTab[1009]++
//line /usr/local/go/src/context/context.go:424
		// _ = "end of CoverTab[1009]"
//line /usr/local/go/src/context/context.go:424
	}
//line /usr/local/go/src/context/context.go:424
	// _ = "end of CoverTab[1004]"
//line /usr/local/go/src/context/context.go:424
	_go_fuzz_dep_.CoverTab[1005]++
							return d.(chan struct{})
//line /usr/local/go/src/context/context.go:425
	// _ = "end of CoverTab[1005]"
}

func (c *cancelCtx) Err() error {
//line /usr/local/go/src/context/context.go:428
	_go_fuzz_dep_.CoverTab[1010]++
							c.mu.Lock()
							err := c.err
							c.mu.Unlock()
							return err
//line /usr/local/go/src/context/context.go:432
	// _ = "end of CoverTab[1010]"
}

type stringer interface {
	String() string
}

func contextName(c Context) string {
//line /usr/local/go/src/context/context.go:439
	_go_fuzz_dep_.CoverTab[1011]++
							if s, ok := c.(stringer); ok {
//line /usr/local/go/src/context/context.go:440
		_go_fuzz_dep_.CoverTab[1013]++
								return s.String()
//line /usr/local/go/src/context/context.go:441
		// _ = "end of CoverTab[1013]"
	} else {
//line /usr/local/go/src/context/context.go:442
		_go_fuzz_dep_.CoverTab[1014]++
//line /usr/local/go/src/context/context.go:442
		// _ = "end of CoverTab[1014]"
//line /usr/local/go/src/context/context.go:442
	}
//line /usr/local/go/src/context/context.go:442
	// _ = "end of CoverTab[1011]"
//line /usr/local/go/src/context/context.go:442
	_go_fuzz_dep_.CoverTab[1012]++
							return reflectlite.TypeOf(c).String()
//line /usr/local/go/src/context/context.go:443
	// _ = "end of CoverTab[1012]"
}

func (c *cancelCtx) String() string {
//line /usr/local/go/src/context/context.go:446
	_go_fuzz_dep_.CoverTab[1015]++
							return contextName(c.Context) + ".WithCancel"
//line /usr/local/go/src/context/context.go:447
	// _ = "end of CoverTab[1015]"
}

// cancel closes c.done, cancels each of c's children, and, if
//line /usr/local/go/src/context/context.go:450
// removeFromParent is true, removes c from its parent's children.
//line /usr/local/go/src/context/context.go:450
// cancel sets c.cause to cause if this is the first time c is canceled.
//line /usr/local/go/src/context/context.go:453
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
//line /usr/local/go/src/context/context.go:453
	_go_fuzz_dep_.CoverTab[1016]++
							if err == nil {
//line /usr/local/go/src/context/context.go:454
		_go_fuzz_dep_.CoverTab[1022]++
								panic("context: internal error: missing cancel error")
//line /usr/local/go/src/context/context.go:455
		// _ = "end of CoverTab[1022]"
	} else {
//line /usr/local/go/src/context/context.go:456
		_go_fuzz_dep_.CoverTab[1023]++
//line /usr/local/go/src/context/context.go:456
		// _ = "end of CoverTab[1023]"
//line /usr/local/go/src/context/context.go:456
	}
//line /usr/local/go/src/context/context.go:456
	// _ = "end of CoverTab[1016]"
//line /usr/local/go/src/context/context.go:456
	_go_fuzz_dep_.CoverTab[1017]++
							if cause == nil {
//line /usr/local/go/src/context/context.go:457
		_go_fuzz_dep_.CoverTab[1024]++
								cause = err
//line /usr/local/go/src/context/context.go:458
		// _ = "end of CoverTab[1024]"
	} else {
//line /usr/local/go/src/context/context.go:459
		_go_fuzz_dep_.CoverTab[1025]++
//line /usr/local/go/src/context/context.go:459
		// _ = "end of CoverTab[1025]"
//line /usr/local/go/src/context/context.go:459
	}
//line /usr/local/go/src/context/context.go:459
	// _ = "end of CoverTab[1017]"
//line /usr/local/go/src/context/context.go:459
	_go_fuzz_dep_.CoverTab[1018]++
							c.mu.Lock()
							if c.err != nil {
//line /usr/local/go/src/context/context.go:461
		_go_fuzz_dep_.CoverTab[1026]++
								c.mu.Unlock()
								return
//line /usr/local/go/src/context/context.go:463
		// _ = "end of CoverTab[1026]"
	} else {
//line /usr/local/go/src/context/context.go:464
		_go_fuzz_dep_.CoverTab[1027]++
//line /usr/local/go/src/context/context.go:464
		// _ = "end of CoverTab[1027]"
//line /usr/local/go/src/context/context.go:464
	}
//line /usr/local/go/src/context/context.go:464
	// _ = "end of CoverTab[1018]"
//line /usr/local/go/src/context/context.go:464
	_go_fuzz_dep_.CoverTab[1019]++
							c.err = err
							c.cause = cause
							d, _ := c.done.Load().(chan struct{})
							if d == nil {
//line /usr/local/go/src/context/context.go:468
		_go_fuzz_dep_.CoverTab[1028]++
								c.done.Store(closedchan)
//line /usr/local/go/src/context/context.go:469
		// _ = "end of CoverTab[1028]"
	} else {
//line /usr/local/go/src/context/context.go:470
		_go_fuzz_dep_.CoverTab[1029]++
								close(d)
//line /usr/local/go/src/context/context.go:471
		// _ = "end of CoverTab[1029]"
	}
//line /usr/local/go/src/context/context.go:472
	// _ = "end of CoverTab[1019]"
//line /usr/local/go/src/context/context.go:472
	_go_fuzz_dep_.CoverTab[1020]++
							for child := range c.children {
//line /usr/local/go/src/context/context.go:473
		_go_fuzz_dep_.CoverTab[1030]++

								child.cancel(false, err, cause)
//line /usr/local/go/src/context/context.go:475
		// _ = "end of CoverTab[1030]"
	}
//line /usr/local/go/src/context/context.go:476
	// _ = "end of CoverTab[1020]"
//line /usr/local/go/src/context/context.go:476
	_go_fuzz_dep_.CoverTab[1021]++
							c.children = nil
							c.mu.Unlock()

							if removeFromParent {
//line /usr/local/go/src/context/context.go:480
		_go_fuzz_dep_.CoverTab[1031]++
								removeChild(c.Context, c)
//line /usr/local/go/src/context/context.go:481
		// _ = "end of CoverTab[1031]"
	} else {
//line /usr/local/go/src/context/context.go:482
		_go_fuzz_dep_.CoverTab[1032]++
//line /usr/local/go/src/context/context.go:482
		// _ = "end of CoverTab[1032]"
//line /usr/local/go/src/context/context.go:482
	}
//line /usr/local/go/src/context/context.go:482
	// _ = "end of CoverTab[1021]"
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
	_go_fuzz_dep_.CoverTab[1033]++
							if parent == nil {
//line /usr/local/go/src/context/context.go:495
		_go_fuzz_dep_.CoverTab[1038]++
								panic("cannot create context from nil parent")
//line /usr/local/go/src/context/context.go:496
		// _ = "end of CoverTab[1038]"
	} else {
//line /usr/local/go/src/context/context.go:497
		_go_fuzz_dep_.CoverTab[1039]++
//line /usr/local/go/src/context/context.go:497
		// _ = "end of CoverTab[1039]"
//line /usr/local/go/src/context/context.go:497
	}
//line /usr/local/go/src/context/context.go:497
	// _ = "end of CoverTab[1033]"
//line /usr/local/go/src/context/context.go:497
	_go_fuzz_dep_.CoverTab[1034]++
							if cur, ok := parent.Deadline(); ok && func() bool {
//line /usr/local/go/src/context/context.go:498
		_go_fuzz_dep_.CoverTab[1040]++
//line /usr/local/go/src/context/context.go:498
		return cur.Before(d)
//line /usr/local/go/src/context/context.go:498
		// _ = "end of CoverTab[1040]"
//line /usr/local/go/src/context/context.go:498
	}() {
//line /usr/local/go/src/context/context.go:498
		_go_fuzz_dep_.CoverTab[1041]++

								return WithCancel(parent)
//line /usr/local/go/src/context/context.go:500
		// _ = "end of CoverTab[1041]"
	} else {
//line /usr/local/go/src/context/context.go:501
		_go_fuzz_dep_.CoverTab[1042]++
//line /usr/local/go/src/context/context.go:501
		// _ = "end of CoverTab[1042]"
//line /usr/local/go/src/context/context.go:501
	}
//line /usr/local/go/src/context/context.go:501
	// _ = "end of CoverTab[1034]"
//line /usr/local/go/src/context/context.go:501
	_go_fuzz_dep_.CoverTab[1035]++
							c := &timerCtx{
		cancelCtx:	newCancelCtx(parent),
		deadline:	d,
	}
	propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
//line /usr/local/go/src/context/context.go:508
		_go_fuzz_dep_.CoverTab[1043]++
								c.cancel(true, DeadlineExceeded, nil)
								return c, func() { _go_fuzz_dep_.CoverTab[1044]++; c.cancel(false, Canceled, nil); // _ = "end of CoverTab[1044]" }
//line /usr/local/go/src/context/context.go:510
		// _ = "end of CoverTab[1043]"
	} else {
//line /usr/local/go/src/context/context.go:511
		_go_fuzz_dep_.CoverTab[1045]++
//line /usr/local/go/src/context/context.go:511
		// _ = "end of CoverTab[1045]"
//line /usr/local/go/src/context/context.go:511
	}
//line /usr/local/go/src/context/context.go:511
	// _ = "end of CoverTab[1035]"
//line /usr/local/go/src/context/context.go:511
	_go_fuzz_dep_.CoverTab[1036]++
							c.mu.Lock()
							defer c.mu.Unlock()
							if c.err == nil {
//line /usr/local/go/src/context/context.go:514
		_go_fuzz_dep_.CoverTab[1046]++
								c.timer = time.AfterFunc(dur, func() {
//line /usr/local/go/src/context/context.go:515
			_go_fuzz_dep_.CoverTab[1047]++
									c.cancel(true, DeadlineExceeded, nil)
//line /usr/local/go/src/context/context.go:516
			// _ = "end of CoverTab[1047]"
		})
//line /usr/local/go/src/context/context.go:517
		// _ = "end of CoverTab[1046]"
	} else {
//line /usr/local/go/src/context/context.go:518
		_go_fuzz_dep_.CoverTab[1048]++
//line /usr/local/go/src/context/context.go:518
		// _ = "end of CoverTab[1048]"
//line /usr/local/go/src/context/context.go:518
	}
//line /usr/local/go/src/context/context.go:518
	// _ = "end of CoverTab[1036]"
//line /usr/local/go/src/context/context.go:518
	_go_fuzz_dep_.CoverTab[1037]++
							return c, func() { _go_fuzz_dep_.CoverTab[1049]++; c.cancel(true, Canceled, nil); // _ = "end of CoverTab[1049]" }
//line /usr/local/go/src/context/context.go:519
	// _ = "end of CoverTab[1037]"
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
	_go_fuzz_dep_.CoverTab[1050]++
							return c.deadline, true
//line /usr/local/go/src/context/context.go:533
	// _ = "end of CoverTab[1050]"
}

func (c *timerCtx) String() string {
//line /usr/local/go/src/context/context.go:536
	_go_fuzz_dep_.CoverTab[1051]++
							return contextName(c.cancelCtx.Context) + ".WithDeadline(" +
		c.deadline.String() + " [" +
		time.Until(c.deadline).String() + "])"
//line /usr/local/go/src/context/context.go:539
	// _ = "end of CoverTab[1051]"
}

func (c *timerCtx) cancel(removeFromParent bool, err, cause error) {
//line /usr/local/go/src/context/context.go:542
	_go_fuzz_dep_.CoverTab[1052]++
							c.cancelCtx.cancel(false, err, cause)
							if removeFromParent {
//line /usr/local/go/src/context/context.go:544
		_go_fuzz_dep_.CoverTab[1055]++

								removeChild(c.cancelCtx.Context, c)
//line /usr/local/go/src/context/context.go:546
		// _ = "end of CoverTab[1055]"
	} else {
//line /usr/local/go/src/context/context.go:547
		_go_fuzz_dep_.CoverTab[1056]++
//line /usr/local/go/src/context/context.go:547
		// _ = "end of CoverTab[1056]"
//line /usr/local/go/src/context/context.go:547
	}
//line /usr/local/go/src/context/context.go:547
	// _ = "end of CoverTab[1052]"
//line /usr/local/go/src/context/context.go:547
	_go_fuzz_dep_.CoverTab[1053]++
							c.mu.Lock()
							if c.timer != nil {
//line /usr/local/go/src/context/context.go:549
		_go_fuzz_dep_.CoverTab[1057]++
								c.timer.Stop()
								c.timer = nil
//line /usr/local/go/src/context/context.go:551
		// _ = "end of CoverTab[1057]"
	} else {
//line /usr/local/go/src/context/context.go:552
		_go_fuzz_dep_.CoverTab[1058]++
//line /usr/local/go/src/context/context.go:552
		// _ = "end of CoverTab[1058]"
//line /usr/local/go/src/context/context.go:552
	}
//line /usr/local/go/src/context/context.go:552
	// _ = "end of CoverTab[1053]"
//line /usr/local/go/src/context/context.go:552
	_go_fuzz_dep_.CoverTab[1054]++
							c.mu.Unlock()
//line /usr/local/go/src/context/context.go:553
	// _ = "end of CoverTab[1054]"
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
	_go_fuzz_dep_.CoverTab[1059]++
							return WithDeadline(parent, time.Now().Add(timeout))
//line /usr/local/go/src/context/context.go:567
	// _ = "end of CoverTab[1059]"
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
	_go_fuzz_dep_.CoverTab[1060]++
							if parent == nil {
//line /usr/local/go/src/context/context.go:584
		_go_fuzz_dep_.CoverTab[1064]++
								panic("cannot create context from nil parent")
//line /usr/local/go/src/context/context.go:585
		// _ = "end of CoverTab[1064]"
	} else {
//line /usr/local/go/src/context/context.go:586
		_go_fuzz_dep_.CoverTab[1065]++
//line /usr/local/go/src/context/context.go:586
		// _ = "end of CoverTab[1065]"
//line /usr/local/go/src/context/context.go:586
	}
//line /usr/local/go/src/context/context.go:586
	// _ = "end of CoverTab[1060]"
//line /usr/local/go/src/context/context.go:586
	_go_fuzz_dep_.CoverTab[1061]++
							if key == nil {
//line /usr/local/go/src/context/context.go:587
		_go_fuzz_dep_.CoverTab[1066]++
								panic("nil key")
//line /usr/local/go/src/context/context.go:588
		// _ = "end of CoverTab[1066]"
	} else {
//line /usr/local/go/src/context/context.go:589
		_go_fuzz_dep_.CoverTab[1067]++
//line /usr/local/go/src/context/context.go:589
		// _ = "end of CoverTab[1067]"
//line /usr/local/go/src/context/context.go:589
	}
//line /usr/local/go/src/context/context.go:589
	// _ = "end of CoverTab[1061]"
//line /usr/local/go/src/context/context.go:589
	_go_fuzz_dep_.CoverTab[1062]++
							if !reflectlite.TypeOf(key).Comparable() {
//line /usr/local/go/src/context/context.go:590
		_go_fuzz_dep_.CoverTab[1068]++
								panic("key is not comparable")
//line /usr/local/go/src/context/context.go:591
		// _ = "end of CoverTab[1068]"
	} else {
//line /usr/local/go/src/context/context.go:592
		_go_fuzz_dep_.CoverTab[1069]++
//line /usr/local/go/src/context/context.go:592
		// _ = "end of CoverTab[1069]"
//line /usr/local/go/src/context/context.go:592
	}
//line /usr/local/go/src/context/context.go:592
	// _ = "end of CoverTab[1062]"
//line /usr/local/go/src/context/context.go:592
	_go_fuzz_dep_.CoverTab[1063]++
							return &valueCtx{parent, key, val}
//line /usr/local/go/src/context/context.go:593
	// _ = "end of CoverTab[1063]"
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
	_go_fuzz_dep_.CoverTab[1070]++
							switch s := v.(type) {
	case stringer:
//line /usr/local/go/src/context/context.go:608
		_go_fuzz_dep_.CoverTab[1072]++
								return s.String()
//line /usr/local/go/src/context/context.go:609
		// _ = "end of CoverTab[1072]"
	case string:
//line /usr/local/go/src/context/context.go:610
		_go_fuzz_dep_.CoverTab[1073]++
								return s
//line /usr/local/go/src/context/context.go:611
		// _ = "end of CoverTab[1073]"
	}
//line /usr/local/go/src/context/context.go:612
	// _ = "end of CoverTab[1070]"
//line /usr/local/go/src/context/context.go:612
	_go_fuzz_dep_.CoverTab[1071]++
							return "<not Stringer>"
//line /usr/local/go/src/context/context.go:613
	// _ = "end of CoverTab[1071]"
}

func (c *valueCtx) String() string {
//line /usr/local/go/src/context/context.go:616
	_go_fuzz_dep_.CoverTab[1074]++
							return contextName(c.Context) + ".WithValue(type " +
		reflectlite.TypeOf(c.key).String() +
		", val " + stringify(c.val) + ")"
//line /usr/local/go/src/context/context.go:619
	// _ = "end of CoverTab[1074]"
}

func (c *valueCtx) Value(key any) any {
//line /usr/local/go/src/context/context.go:622
	_go_fuzz_dep_.CoverTab[1075]++
							if c.key == key {
//line /usr/local/go/src/context/context.go:623
		_go_fuzz_dep_.CoverTab[1077]++
								return c.val
//line /usr/local/go/src/context/context.go:624
		// _ = "end of CoverTab[1077]"
	} else {
//line /usr/local/go/src/context/context.go:625
		_go_fuzz_dep_.CoverTab[1078]++
//line /usr/local/go/src/context/context.go:625
		// _ = "end of CoverTab[1078]"
//line /usr/local/go/src/context/context.go:625
	}
//line /usr/local/go/src/context/context.go:625
	// _ = "end of CoverTab[1075]"
//line /usr/local/go/src/context/context.go:625
	_go_fuzz_dep_.CoverTab[1076]++
							return value(c.Context, key)
//line /usr/local/go/src/context/context.go:626
	// _ = "end of CoverTab[1076]"
}

func value(c Context, key any) any {
//line /usr/local/go/src/context/context.go:629
	_go_fuzz_dep_.CoverTab[1079]++
							for {
//line /usr/local/go/src/context/context.go:630
		_go_fuzz_dep_.CoverTab[1080]++
								switch ctx := c.(type) {
		case *valueCtx:
//line /usr/local/go/src/context/context.go:632
			_go_fuzz_dep_.CoverTab[1081]++
									if key == ctx.key {
//line /usr/local/go/src/context/context.go:633
				_go_fuzz_dep_.CoverTab[1089]++
										return ctx.val
//line /usr/local/go/src/context/context.go:634
				// _ = "end of CoverTab[1089]"
			} else {
//line /usr/local/go/src/context/context.go:635
				_go_fuzz_dep_.CoverTab[1090]++
//line /usr/local/go/src/context/context.go:635
				// _ = "end of CoverTab[1090]"
//line /usr/local/go/src/context/context.go:635
			}
//line /usr/local/go/src/context/context.go:635
			// _ = "end of CoverTab[1081]"
//line /usr/local/go/src/context/context.go:635
			_go_fuzz_dep_.CoverTab[1082]++
									c = ctx.Context
//line /usr/local/go/src/context/context.go:636
			// _ = "end of CoverTab[1082]"
		case *cancelCtx:
//line /usr/local/go/src/context/context.go:637
			_go_fuzz_dep_.CoverTab[1083]++
									if key == &cancelCtxKey {
//line /usr/local/go/src/context/context.go:638
				_go_fuzz_dep_.CoverTab[1091]++
										return c
//line /usr/local/go/src/context/context.go:639
				// _ = "end of CoverTab[1091]"
			} else {
//line /usr/local/go/src/context/context.go:640
				_go_fuzz_dep_.CoverTab[1092]++
//line /usr/local/go/src/context/context.go:640
				// _ = "end of CoverTab[1092]"
//line /usr/local/go/src/context/context.go:640
			}
//line /usr/local/go/src/context/context.go:640
			// _ = "end of CoverTab[1083]"
//line /usr/local/go/src/context/context.go:640
			_go_fuzz_dep_.CoverTab[1084]++
									c = ctx.Context
//line /usr/local/go/src/context/context.go:641
			// _ = "end of CoverTab[1084]"
		case *timerCtx:
//line /usr/local/go/src/context/context.go:642
			_go_fuzz_dep_.CoverTab[1085]++
									if key == &cancelCtxKey {
//line /usr/local/go/src/context/context.go:643
				_go_fuzz_dep_.CoverTab[1093]++
										return ctx.cancelCtx
//line /usr/local/go/src/context/context.go:644
				// _ = "end of CoverTab[1093]"
			} else {
//line /usr/local/go/src/context/context.go:645
				_go_fuzz_dep_.CoverTab[1094]++
//line /usr/local/go/src/context/context.go:645
				// _ = "end of CoverTab[1094]"
//line /usr/local/go/src/context/context.go:645
			}
//line /usr/local/go/src/context/context.go:645
			// _ = "end of CoverTab[1085]"
//line /usr/local/go/src/context/context.go:645
			_go_fuzz_dep_.CoverTab[1086]++
									c = ctx.Context
//line /usr/local/go/src/context/context.go:646
			// _ = "end of CoverTab[1086]"
		case *emptyCtx:
//line /usr/local/go/src/context/context.go:647
			_go_fuzz_dep_.CoverTab[1087]++
									return nil
//line /usr/local/go/src/context/context.go:648
			// _ = "end of CoverTab[1087]"
		default:
//line /usr/local/go/src/context/context.go:649
			_go_fuzz_dep_.CoverTab[1088]++
									return c.Value(key)
//line /usr/local/go/src/context/context.go:650
			// _ = "end of CoverTab[1088]"
		}
//line /usr/local/go/src/context/context.go:651
		// _ = "end of CoverTab[1080]"
	}
//line /usr/local/go/src/context/context.go:652
	// _ = "end of CoverTab[1079]"
}

//line /usr/local/go/src/context/context.go:653
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/context/context.go:653
var _ = _go_fuzz_dep_.CoverTab
