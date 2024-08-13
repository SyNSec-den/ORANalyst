// Copyright 2010 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// Package gomock is a mock framework for Go.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// Standard usage:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	(1) Define an interface that you wish to mock.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	      type MyInterface interface {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	        SomeMethod(x int64, y string)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	      }
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	(2) Use mockgen to generate a mock from the interface.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	(3) Use the mock in a test:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	      func TestMyThing(t *testing.T) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	        mockCtrl := gomock.NewController(t)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	        defer mockCtrl.Finish()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	        mockObj := something.NewMockMyInterface(mockCtrl)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	        mockObj.EXPECT().SomeMethod(4, "blah")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	        // pass mockObj to a real object and play with it.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	      }
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// By default, expected calls are not enforced to run in any particular order.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// Call order dependency can be enforced by use of InOrder and/or Call.After.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// Call.After can create more varied call order dependencies, but InOrder is
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// often more convenient.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// The following examples create equivalent call order dependencies.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// Example of using Call.After to chain expected call order:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	firstCall := mockObj.EXPECT().SomeMethod(1, "first")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	secondCall := mockObj.EXPECT().SomeMethod(2, "second").After(firstCall)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	mockObj.EXPECT().SomeMethod(3, "third").After(secondCall)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
// Example of using InOrder to declare expected call order:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	gomock.InOrder(
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	    mockObj.EXPECT().SomeMethod(1, "first"),
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	    mockObj.EXPECT().SomeMethod(2, "second"),
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	    mockObj.EXPECT().SomeMethod(3, "third"),
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:15
//	)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
package gomock

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:53
)

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// A TestReporter is something that can be used to report test failures.  It
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:63
// is satisfied by the standard library's *testing.T.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:65
type TestReporter interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// TestHelper is a TestReporter that has the Helper method.  It is satisfied
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:70
// by the standard library's *testing.T.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:72
type TestHelper interface {
	TestReporter
	Helper()
}

// cleanuper is used to check if TestHelper also has the `Cleanup` method. A
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:77
// common pattern is to pass in a `*testing.T` to
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:77
// `NewController(t TestReporter)`. In Go 1.14+, `*testing.T` has a cleanup
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:77
// method. This can be utilized to call `Finish()` so the caller of this library
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:77
// does not have to.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:82
type cleanuper interface {
	Cleanup(func())
}

// A Controller represents the top-level control of a mock ecosystem.  It
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
// defines the scope and lifetime of mock objects, as well as their
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
// expectations.  It is safe to call Controller's methods from multiple
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
// goroutines. Each test should create a new Controller and invoke Finish via
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
// defer.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	func TestFoo(t *testing.T) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  ctrl := gomock.NewController(t)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  defer ctrl.Finish()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  // ..
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	func TestBar(t *testing.T) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  t.Run("Sub-Test-1", st) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	    ctrl := gomock.NewController(st)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	    defer ctrl.Finish()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	    // ..
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  })
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  t.Run("Sub-Test-2", st) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	    ctrl := gomock.NewController(st)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	    defer ctrl.Finish()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	    // ..
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	  })
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:86
//	})
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:110
type Controller struct {
	// T should only be called within a generated mock. It is not intended to
	// be used in user code and may be changed in future versions. T is the
	// TestReporter passed in when creating the Controller via NewController.
	// If the TestReporter does not implement a TestHelper it will be wrapped
	// with a nopTestHelper.
	T		TestHelper
	mu		sync.Mutex
	expectedCalls	*callSet
	finished	bool
}

// NewController returns a new Controller. It is the preferred way to create a
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:122
// Controller.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:122
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:122
// New in go1.14+, if you are passing a *testing.T into this function you no
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:122
// longer need to call ctrl.Finish() in your test methods.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:127
func NewController(t TestReporter) *Controller {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:127
	_go_fuzz_dep_.CoverTab[142764]++
												h, ok := t.(TestHelper)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:129
		_go_fuzz_dep_.CoverTab[142767]++
													h = &nopTestHelper{t}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:130
		// _ = "end of CoverTab[142767]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:131
		_go_fuzz_dep_.CoverTab[142768]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:131
		// _ = "end of CoverTab[142768]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:131
	// _ = "end of CoverTab[142764]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:131
	_go_fuzz_dep_.CoverTab[142765]++
												ctrl := &Controller{
		T:		h,
		expectedCalls:	newCallSet(),
	}
	if c, ok := isCleanuper(ctrl.T); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:136
		_go_fuzz_dep_.CoverTab[142769]++
													c.Cleanup(func() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:137
			_go_fuzz_dep_.CoverTab[142770]++
														ctrl.T.Helper()
														ctrl.finish(true, nil)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:139
			// _ = "end of CoverTab[142770]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:140
		// _ = "end of CoverTab[142769]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:141
		_go_fuzz_dep_.CoverTab[142771]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:141
		// _ = "end of CoverTab[142771]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:141
	// _ = "end of CoverTab[142765]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:141
	_go_fuzz_dep_.CoverTab[142766]++

												return ctrl
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:143
	// _ = "end of CoverTab[142766]"
}

type cancelReporter struct {
	t	TestHelper
	cancel	func()
}

func (r *cancelReporter) Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:151
	_go_fuzz_dep_.CoverTab[142772]++
												r.t.Errorf(format, args...)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:152
	// _ = "end of CoverTab[142772]"
}
func (r *cancelReporter) Fatalf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:154
	_go_fuzz_dep_.CoverTab[142773]++
												defer r.cancel()
												r.t.Fatalf(format, args...)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:156
	// _ = "end of CoverTab[142773]"
}

func (r *cancelReporter) Helper() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:159
	_go_fuzz_dep_.CoverTab[142774]++
												r.t.Helper()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:160
	// _ = "end of CoverTab[142774]"
}

// WithContext returns a new Controller and a Context, which is cancelled on any
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:163
// fatal failure.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:165
func WithContext(ctx context.Context, t TestReporter) (*Controller, context.Context) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:165
	_go_fuzz_dep_.CoverTab[142775]++
												h, ok := t.(TestHelper)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:167
		_go_fuzz_dep_.CoverTab[142777]++
													h = &nopTestHelper{t: t}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:168
		// _ = "end of CoverTab[142777]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:169
		_go_fuzz_dep_.CoverTab[142778]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:169
		// _ = "end of CoverTab[142778]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:169
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:169
	// _ = "end of CoverTab[142775]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:169
	_go_fuzz_dep_.CoverTab[142776]++

												ctx, cancel := context.WithCancel(ctx)
												return NewController(&cancelReporter{t: h, cancel: cancel}), ctx
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:172
	// _ = "end of CoverTab[142776]"
}

type nopTestHelper struct {
	t TestReporter
}

func (h *nopTestHelper) Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:179
	_go_fuzz_dep_.CoverTab[142779]++
												h.t.Errorf(format, args...)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:180
	// _ = "end of CoverTab[142779]"
}
func (h *nopTestHelper) Fatalf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:182
	_go_fuzz_dep_.CoverTab[142780]++
												h.t.Fatalf(format, args...)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:183
	// _ = "end of CoverTab[142780]"
}

func (h nopTestHelper) Helper()	{ _go_fuzz_dep_.CoverTab[142781]++; // _ = "end of CoverTab[142781]" }

// RecordCall is called by a mock. It should not be called by user code.
func (ctrl *Controller) RecordCall(receiver interface{}, method string, args ...interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:189
	_go_fuzz_dep_.CoverTab[142782]++
												ctrl.T.Helper()

												recv := reflect.ValueOf(receiver)
												for i := 0; i < recv.Type().NumMethod(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:193
		_go_fuzz_dep_.CoverTab[142784]++
													if recv.Type().Method(i).Name == method {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:194
			_go_fuzz_dep_.CoverTab[142785]++
														return ctrl.RecordCallWithMethodType(receiver, method, recv.Method(i).Type(), args...)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:195
			// _ = "end of CoverTab[142785]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:196
			_go_fuzz_dep_.CoverTab[142786]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:196
			// _ = "end of CoverTab[142786]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:196
		// _ = "end of CoverTab[142784]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:197
	// _ = "end of CoverTab[142782]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:197
	_go_fuzz_dep_.CoverTab[142783]++
												ctrl.T.Fatalf("gomock: failed finding method %s on %T", method, receiver)
												panic("unreachable")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:199
	// _ = "end of CoverTab[142783]"
}

// RecordCallWithMethodType is called by a mock. It should not be called by user code.
func (ctrl *Controller) RecordCallWithMethodType(receiver interface{}, method string, methodType reflect.Type, args ...interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:203
	_go_fuzz_dep_.CoverTab[142787]++
												ctrl.T.Helper()

												call := newCall(ctrl.T, receiver, method, methodType, args...)

												ctrl.mu.Lock()
												defer ctrl.mu.Unlock()
												ctrl.expectedCalls.Add(call)

												return call
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:212
	// _ = "end of CoverTab[142787]"
}

// Call is called by a mock. It should not be called by user code.
func (ctrl *Controller) Call(receiver interface{}, method string, args ...interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:216
	_go_fuzz_dep_.CoverTab[142788]++
												ctrl.T.Helper()

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:220
	actions := func() []func([]interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:220
		_go_fuzz_dep_.CoverTab[142791]++
													ctrl.T.Helper()
													ctrl.mu.Lock()
													defer ctrl.mu.Unlock()

													expected, err := ctrl.expectedCalls.FindMatch(receiver, method, args)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:226
			_go_fuzz_dep_.CoverTab[142795]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:230
			origin := callerInfo(3)
														ctrl.T.Fatalf("Unexpected call to %T.%v(%v) at %s because: %s", receiver, method, args, origin, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:231
			// _ = "end of CoverTab[142795]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:232
			_go_fuzz_dep_.CoverTab[142796]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:232
			// _ = "end of CoverTab[142796]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:232
		// _ = "end of CoverTab[142791]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:232
		_go_fuzz_dep_.CoverTab[142792]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:237
		preReqCalls := expected.dropPrereqs()
		for _, preReqCall := range preReqCalls {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:238
			_go_fuzz_dep_.CoverTab[142797]++
														ctrl.expectedCalls.Remove(preReqCall)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:239
			// _ = "end of CoverTab[142797]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:240
		// _ = "end of CoverTab[142792]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:240
		_go_fuzz_dep_.CoverTab[142793]++

													actions := expected.call()
													if expected.exhausted() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:243
			_go_fuzz_dep_.CoverTab[142798]++
														ctrl.expectedCalls.Remove(expected)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:244
			// _ = "end of CoverTab[142798]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:245
			_go_fuzz_dep_.CoverTab[142799]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:245
			// _ = "end of CoverTab[142799]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:245
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:245
		// _ = "end of CoverTab[142793]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:245
		_go_fuzz_dep_.CoverTab[142794]++
													return actions
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:246
		// _ = "end of CoverTab[142794]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:247
	// _ = "end of CoverTab[142788]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:247
	_go_fuzz_dep_.CoverTab[142789]++

												var rets []interface{}
												for _, action := range actions {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:250
		_go_fuzz_dep_.CoverTab[142800]++
													if r := action(args); r != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:251
			_go_fuzz_dep_.CoverTab[142801]++
														rets = r
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:252
			// _ = "end of CoverTab[142801]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:253
			_go_fuzz_dep_.CoverTab[142802]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:253
			// _ = "end of CoverTab[142802]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:253
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:253
		// _ = "end of CoverTab[142800]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:254
	// _ = "end of CoverTab[142789]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:254
	_go_fuzz_dep_.CoverTab[142790]++

												return rets
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:256
	// _ = "end of CoverTab[142790]"
}

// Finish checks to see if all the methods that were expected to be called
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:259
// were called. It should be invoked for each Controller. It is not idempotent
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:259
// and therefore can only be invoked once.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:259
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:259
// New in go1.14+, if you are passing a *testing.T into NewController function you no
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:259
// longer need to call ctrl.Finish() in your test methods.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:265
func (ctrl *Controller) Finish() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:265
	_go_fuzz_dep_.CoverTab[142803]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:268
	err := recover()
												ctrl.finish(false, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:269
	// _ = "end of CoverTab[142803]"
}

func (ctrl *Controller) finish(cleanup bool, panicErr interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:272
	_go_fuzz_dep_.CoverTab[142804]++
												ctrl.T.Helper()

												ctrl.mu.Lock()
												defer ctrl.mu.Unlock()

												if ctrl.finished {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:278
		_go_fuzz_dep_.CoverTab[142808]++
													if _, ok := isCleanuper(ctrl.T); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:279
			_go_fuzz_dep_.CoverTab[142810]++
														ctrl.T.Fatalf("Controller.Finish was called more than once. It has to be called exactly once.")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:280
			// _ = "end of CoverTab[142810]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:281
			_go_fuzz_dep_.CoverTab[142811]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:281
			// _ = "end of CoverTab[142811]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:281
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:281
		// _ = "end of CoverTab[142808]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:281
		_go_fuzz_dep_.CoverTab[142809]++
													return
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:282
		// _ = "end of CoverTab[142809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:283
		_go_fuzz_dep_.CoverTab[142812]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:283
		// _ = "end of CoverTab[142812]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:283
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:283
	// _ = "end of CoverTab[142804]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:283
	_go_fuzz_dep_.CoverTab[142805]++
												ctrl.finished = true

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:287
	if panicErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:287
		_go_fuzz_dep_.CoverTab[142813]++
													panic(panicErr)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:288
		// _ = "end of CoverTab[142813]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:289
		_go_fuzz_dep_.CoverTab[142814]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:289
		// _ = "end of CoverTab[142814]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:289
	// _ = "end of CoverTab[142805]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:289
	_go_fuzz_dep_.CoverTab[142806]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:292
	failures := ctrl.expectedCalls.Failures()
	for _, call := range failures {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:293
		_go_fuzz_dep_.CoverTab[142815]++
													ctrl.T.Errorf("missing call(s) to %v", call)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:294
		// _ = "end of CoverTab[142815]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:295
	// _ = "end of CoverTab[142806]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:295
	_go_fuzz_dep_.CoverTab[142807]++
												if len(failures) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:296
		_go_fuzz_dep_.CoverTab[142816]++
													if !cleanup {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:297
			_go_fuzz_dep_.CoverTab[142818]++
														ctrl.T.Fatalf("aborting test due to missing call(s)")
														return
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:299
			// _ = "end of CoverTab[142818]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:300
			_go_fuzz_dep_.CoverTab[142819]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:300
			// _ = "end of CoverTab[142819]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:300
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:300
		// _ = "end of CoverTab[142816]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:300
		_go_fuzz_dep_.CoverTab[142817]++
													ctrl.T.Errorf("aborting test due to missing call(s)")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:301
		// _ = "end of CoverTab[142817]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:302
		_go_fuzz_dep_.CoverTab[142820]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:302
		// _ = "end of CoverTab[142820]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:302
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:302
	// _ = "end of CoverTab[142807]"
}

// callerInfo returns the file:line of the call site. skip is the number
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:305
// of stack frames to skip when reporting. 0 is callerInfo's call site.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:307
func callerInfo(skip int) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:307
	_go_fuzz_dep_.CoverTab[142821]++
												if _, file, line, ok := runtime.Caller(skip + 1); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:308
		_go_fuzz_dep_.CoverTab[142823]++
													return fmt.Sprintf("%s:%d", file, line)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:309
		// _ = "end of CoverTab[142823]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:310
		_go_fuzz_dep_.CoverTab[142824]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:310
		// _ = "end of CoverTab[142824]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:310
	// _ = "end of CoverTab[142821]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:310
	_go_fuzz_dep_.CoverTab[142822]++
												return "unknown file"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:311
	// _ = "end of CoverTab[142822]"
}

// isCleanuper checks it if t's base TestReporter has a Cleanup method.
func isCleanuper(t TestReporter) (cleanuper, bool) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:315
	_go_fuzz_dep_.CoverTab[142825]++
												tr := unwrapTestReporter(t)
												c, ok := tr.(cleanuper)
												return c, ok
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:318
	// _ = "end of CoverTab[142825]"
}

// unwrapTestReporter unwraps TestReporter to the base implementation.
func unwrapTestReporter(t TestReporter) TestReporter {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:322
	_go_fuzz_dep_.CoverTab[142826]++
												tr := t
												switch nt := t.(type) {
	case *cancelReporter:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:325
		_go_fuzz_dep_.CoverTab[142828]++
													tr = nt.t
													if h, check := tr.(*nopTestHelper); check {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:327
			_go_fuzz_dep_.CoverTab[142831]++
														tr = h.t
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:328
			// _ = "end of CoverTab[142831]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:329
			_go_fuzz_dep_.CoverTab[142832]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:329
			// _ = "end of CoverTab[142832]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:329
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:329
		// _ = "end of CoverTab[142828]"
	case *nopTestHelper:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:330
		_go_fuzz_dep_.CoverTab[142829]++
													tr = nt.t
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:331
		// _ = "end of CoverTab[142829]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:332
		_go_fuzz_dep_.CoverTab[142830]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:332
		// _ = "end of CoverTab[142830]"

	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:334
	// _ = "end of CoverTab[142826]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:334
	_go_fuzz_dep_.CoverTab[142827]++
												return tr
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:335
	// _ = "end of CoverTab[142827]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:336
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/controller.go:336
var _ = _go_fuzz_dep_.CoverTab
