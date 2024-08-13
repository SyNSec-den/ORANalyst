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

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
package gomock

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:15
)

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Call represents an expected call to a mock.
type Call struct {
	t	TestHelper	// for triggering test failures on invalid call setup

	receiver	interface{}	// the receiver of the method call
	method		string		// the name of the method
	methodType	reflect.Type	// the type of the method
	args		[]Matcher	// the args
	origin		string		// file and line number of call setup

	preReqs	[]*Call	// prerequisite calls

	// Expectations
	minCalls, maxCalls	int

	numCalls	int	// actual number made

	// actions are called when this Call is called. Each action gets the args and
	// can set the return values by returning a non-nil slice. Actions run in the
	// order they are created.
	actions	[]func([]interface{}) []interface{}
}

// newCall creates a *Call. It requires the method type in order to support
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:47
// unexported methods.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:49
func newCall(t TestHelper, receiver interface{}, method string, methodType reflect.Type, args ...interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:49
	_go_fuzz_dep_.CoverTab[142586]++
											t.Helper()

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:53
	mArgs := make([]Matcher, len(args))
	for i, arg := range args {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:54
		_go_fuzz_dep_.CoverTab[142589]++
												if m, ok := arg.(Matcher); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:55
			_go_fuzz_dep_.CoverTab[142590]++
													mArgs[i] = m
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:56
			// _ = "end of CoverTab[142590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:57
			_go_fuzz_dep_.CoverTab[142591]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:57
			if arg == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:57
				_go_fuzz_dep_.CoverTab[142592]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:60
				mArgs[i] = Nil()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:60
				// _ = "end of CoverTab[142592]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:61
				_go_fuzz_dep_.CoverTab[142593]++
														mArgs[i] = Eq(arg)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:62
				// _ = "end of CoverTab[142593]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:63
			// _ = "end of CoverTab[142591]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:63
		// _ = "end of CoverTab[142589]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:64
	// _ = "end of CoverTab[142586]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:64
	_go_fuzz_dep_.CoverTab[142587]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:69
	origin := callerInfo(3)
	actions := []func([]interface{}) []interface{}{func([]interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:70
		_go_fuzz_dep_.CoverTab[142594]++

												rets := make([]interface{}, methodType.NumOut())
												for i := 0; i < methodType.NumOut(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:73
			_go_fuzz_dep_.CoverTab[142596]++
													rets[i] = reflect.Zero(methodType.Out(i)).Interface()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:74
			// _ = "end of CoverTab[142596]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:75
		// _ = "end of CoverTab[142594]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:75
		_go_fuzz_dep_.CoverTab[142595]++
												return rets
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:76
		// _ = "end of CoverTab[142595]"
	}}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:77
	// _ = "end of CoverTab[142587]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:77
	_go_fuzz_dep_.CoverTab[142588]++
											return &Call{t: t, receiver: receiver, method: method, methodType: methodType,
		args:	mArgs, origin: origin, minCalls: 1, maxCalls: 1, actions: actions}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:79
	// _ = "end of CoverTab[142588]"
}

// AnyTimes allows the expectation to be called 0 or more times
func (c *Call) AnyTimes() *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:83
	_go_fuzz_dep_.CoverTab[142597]++
											c.minCalls, c.maxCalls = 0, 1e8
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:85
	// _ = "end of CoverTab[142597]"
}

// MinTimes requires the call to occur at least n times. If AnyTimes or MaxTimes have not been called or if MaxTimes
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:88
// was previously called with 1, MinTimes also sets the maximum number of calls to infinity.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:90
func (c *Call) MinTimes(n int) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:90
	_go_fuzz_dep_.CoverTab[142598]++
											c.minCalls = n
											if c.maxCalls == 1 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:92
		_go_fuzz_dep_.CoverTab[142600]++
												c.maxCalls = 1e8
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:93
		// _ = "end of CoverTab[142600]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:94
		_go_fuzz_dep_.CoverTab[142601]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:94
		// _ = "end of CoverTab[142601]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:94
	// _ = "end of CoverTab[142598]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:94
	_go_fuzz_dep_.CoverTab[142599]++
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:95
	// _ = "end of CoverTab[142599]"
}

// MaxTimes limits the number of calls to n times. If AnyTimes or MinTimes have not been called or if MinTimes was
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:98
// previously called with 1, MaxTimes also sets the minimum number of calls to 0.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:100
func (c *Call) MaxTimes(n int) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:100
	_go_fuzz_dep_.CoverTab[142602]++
											c.maxCalls = n
											if c.minCalls == 1 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:102
		_go_fuzz_dep_.CoverTab[142604]++
												c.minCalls = 0
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:103
		// _ = "end of CoverTab[142604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:104
		_go_fuzz_dep_.CoverTab[142605]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:104
		// _ = "end of CoverTab[142605]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:104
	// _ = "end of CoverTab[142602]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:104
	_go_fuzz_dep_.CoverTab[142603]++
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:105
	// _ = "end of CoverTab[142603]"
}

// DoAndReturn declares the action to run when the call is matched.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:108
// The return values from this function are returned by the mocked function.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:108
// It takes an interface{} argument to support n-arity functions.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:111
func (c *Call) DoAndReturn(f interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:111
	_go_fuzz_dep_.CoverTab[142606]++

											v := reflect.ValueOf(f)

											c.addAction(func(args []interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:115
		_go_fuzz_dep_.CoverTab[142608]++
												c.t.Helper()
												vArgs := make([]reflect.Value, len(args))
												ft := v.Type()
												if c.methodType.NumIn() != ft.NumIn() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:119
			_go_fuzz_dep_.CoverTab[142612]++
													c.t.Fatalf("wrong number of arguments in DoAndReturn func for %T.%v: got %d, want %d [%s]",
				c.receiver, c.method, ft.NumIn(), c.methodType.NumIn(), c.origin)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:122
			// _ = "end of CoverTab[142612]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:123
			_go_fuzz_dep_.CoverTab[142613]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:123
			// _ = "end of CoverTab[142613]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:123
		// _ = "end of CoverTab[142608]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:123
		_go_fuzz_dep_.CoverTab[142609]++
												for i := 0; i < len(args); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:124
			_go_fuzz_dep_.CoverTab[142614]++
													if args[i] != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:125
				_go_fuzz_dep_.CoverTab[142615]++
														vArgs[i] = reflect.ValueOf(args[i])
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:126
				// _ = "end of CoverTab[142615]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:127
				_go_fuzz_dep_.CoverTab[142616]++

														vArgs[i] = reflect.Zero(ft.In(i))
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:129
				// _ = "end of CoverTab[142616]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:130
			// _ = "end of CoverTab[142614]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:131
		// _ = "end of CoverTab[142609]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:131
		_go_fuzz_dep_.CoverTab[142610]++
												vRets := v.Call(vArgs)
												rets := make([]interface{}, len(vRets))
												for i, ret := range vRets {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:134
			_go_fuzz_dep_.CoverTab[142617]++
													rets[i] = ret.Interface()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:135
			// _ = "end of CoverTab[142617]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:136
		// _ = "end of CoverTab[142610]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:136
		_go_fuzz_dep_.CoverTab[142611]++
												return rets
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:137
		// _ = "end of CoverTab[142611]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:138
	// _ = "end of CoverTab[142606]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:138
	_go_fuzz_dep_.CoverTab[142607]++
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:139
	// _ = "end of CoverTab[142607]"
}

// Do declares the action to run when the call is matched. The function's
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:142
// return values are ignored to retain backward compatibility. To use the
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:142
// return values call DoAndReturn.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:142
// It takes an interface{} argument to support n-arity functions.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:146
func (c *Call) Do(f interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:146
	_go_fuzz_dep_.CoverTab[142618]++

											v := reflect.ValueOf(f)

											c.addAction(func(args []interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:150
		_go_fuzz_dep_.CoverTab[142620]++
												c.t.Helper()
												if c.methodType.NumIn() != v.Type().NumIn() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:152
			_go_fuzz_dep_.CoverTab[142623]++
													c.t.Fatalf("wrong number of arguments in Do func for %T.%v: got %d, want %d [%s]",
				c.receiver, c.method, v.Type().NumIn(), c.methodType.NumIn(), c.origin)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:155
			// _ = "end of CoverTab[142623]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:156
			_go_fuzz_dep_.CoverTab[142624]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:156
			// _ = "end of CoverTab[142624]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:156
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:156
		// _ = "end of CoverTab[142620]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:156
		_go_fuzz_dep_.CoverTab[142621]++
												vArgs := make([]reflect.Value, len(args))
												ft := v.Type()
												for i := 0; i < len(args); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:159
			_go_fuzz_dep_.CoverTab[142625]++
													if args[i] != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:160
				_go_fuzz_dep_.CoverTab[142626]++
														vArgs[i] = reflect.ValueOf(args[i])
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:161
				// _ = "end of CoverTab[142626]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:162
				_go_fuzz_dep_.CoverTab[142627]++

														vArgs[i] = reflect.Zero(ft.In(i))
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:164
				// _ = "end of CoverTab[142627]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:165
			// _ = "end of CoverTab[142625]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:166
		// _ = "end of CoverTab[142621]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:166
		_go_fuzz_dep_.CoverTab[142622]++
												v.Call(vArgs)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:168
		// _ = "end of CoverTab[142622]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:169
	// _ = "end of CoverTab[142618]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:169
	_go_fuzz_dep_.CoverTab[142619]++
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:170
	// _ = "end of CoverTab[142619]"
}

// Return declares the values to be returned by the mocked function call.
func (c *Call) Return(rets ...interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:174
	_go_fuzz_dep_.CoverTab[142628]++
											c.t.Helper()

											mt := c.methodType
											if len(rets) != mt.NumOut() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:178
		_go_fuzz_dep_.CoverTab[142632]++
												c.t.Fatalf("wrong number of arguments to Return for %T.%v: got %d, want %d [%s]",
			c.receiver, c.method, len(rets), mt.NumOut(), c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:180
		// _ = "end of CoverTab[142632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:181
		_go_fuzz_dep_.CoverTab[142633]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:181
		// _ = "end of CoverTab[142633]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:181
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:181
	// _ = "end of CoverTab[142628]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:181
	_go_fuzz_dep_.CoverTab[142629]++
											for i, ret := range rets {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:182
		_go_fuzz_dep_.CoverTab[142634]++
												if got, want := reflect.TypeOf(ret), mt.Out(i); got == want {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:183
			_go_fuzz_dep_.CoverTab[142635]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:183
			// _ = "end of CoverTab[142635]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:185
			_go_fuzz_dep_.CoverTab[142636]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:185
			if got == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:185
				_go_fuzz_dep_.CoverTab[142637]++

														switch want.Kind() {
				case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:188
					_go_fuzz_dep_.CoverTab[142638]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:188
					// _ = "end of CoverTab[142638]"

				default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:190
					_go_fuzz_dep_.CoverTab[142639]++
															c.t.Fatalf("argument %d to Return for %T.%v is nil, but %v is not nillable [%s]",
						i, c.receiver, c.method, want, c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:192
					// _ = "end of CoverTab[142639]"
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:193
				// _ = "end of CoverTab[142637]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:194
				_go_fuzz_dep_.CoverTab[142640]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:194
				if got.AssignableTo(want) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:194
					_go_fuzz_dep_.CoverTab[142641]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:197
					v := reflect.New(want).Elem()
															v.Set(reflect.ValueOf(ret))
															rets[i] = v.Interface()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:199
					// _ = "end of CoverTab[142641]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:200
					_go_fuzz_dep_.CoverTab[142642]++
															c.t.Fatalf("wrong type of argument %d to Return for %T.%v: %v is not assignable to %v [%s]",
						i, c.receiver, c.method, got, want, c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:202
					// _ = "end of CoverTab[142642]"
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:203
				// _ = "end of CoverTab[142640]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:203
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:203
			// _ = "end of CoverTab[142636]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:203
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:203
		// _ = "end of CoverTab[142634]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:204
	// _ = "end of CoverTab[142629]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:204
	_go_fuzz_dep_.CoverTab[142630]++

											c.addAction(func([]interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:206
		_go_fuzz_dep_.CoverTab[142643]++
												return rets
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:207
		// _ = "end of CoverTab[142643]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:208
	// _ = "end of CoverTab[142630]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:208
	_go_fuzz_dep_.CoverTab[142631]++

											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:210
	// _ = "end of CoverTab[142631]"
}

// Times declares the exact number of times a function call is expected to be executed.
func (c *Call) Times(n int) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:214
	_go_fuzz_dep_.CoverTab[142644]++
											c.minCalls, c.maxCalls = n, n
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:216
	// _ = "end of CoverTab[142644]"
}

// SetArg declares an action that will set the nth argument's value,
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:219
// indirected through a pointer. Or, in the case of a slice, SetArg
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:219
// will copy value's elements into the nth argument.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:222
func (c *Call) SetArg(n int, value interface{}) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:222
	_go_fuzz_dep_.CoverTab[142645]++
											c.t.Helper()

											mt := c.methodType

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:228
	if n < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:228
		_go_fuzz_dep_.CoverTab[142649]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:228
		return n >= mt.NumIn()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:228
		// _ = "end of CoverTab[142649]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:228
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:228
		_go_fuzz_dep_.CoverTab[142650]++
												c.t.Fatalf("SetArg(%d, ...) called for a method with %d args [%s]",
			n, mt.NumIn(), c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:230
		// _ = "end of CoverTab[142650]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:231
		_go_fuzz_dep_.CoverTab[142651]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:231
		// _ = "end of CoverTab[142651]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:231
	// _ = "end of CoverTab[142645]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:231
	_go_fuzz_dep_.CoverTab[142646]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:234
	at := mt.In(n)
	switch at.Kind() {
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:236
		_go_fuzz_dep_.CoverTab[142652]++
												dt := at.Elem()
												if vt := reflect.TypeOf(value); !vt.AssignableTo(dt) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:238
			_go_fuzz_dep_.CoverTab[142656]++
													c.t.Fatalf("SetArg(%d, ...) argument is a %v, not assignable to %v [%s]",
				n, vt, dt, c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:240
			// _ = "end of CoverTab[142656]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:241
			_go_fuzz_dep_.CoverTab[142657]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:241
			// _ = "end of CoverTab[142657]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:241
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:241
		// _ = "end of CoverTab[142652]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:242
		_go_fuzz_dep_.CoverTab[142653]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:242
		// _ = "end of CoverTab[142653]"

	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:244
		_go_fuzz_dep_.CoverTab[142654]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:244
		// _ = "end of CoverTab[142654]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:246
		_go_fuzz_dep_.CoverTab[142655]++
												c.t.Fatalf("SetArg(%d, ...) referring to argument of non-pointer non-interface non-slice type %v [%s]",
			n, at, c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:248
		// _ = "end of CoverTab[142655]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:249
	// _ = "end of CoverTab[142646]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:249
	_go_fuzz_dep_.CoverTab[142647]++

											c.addAction(func(args []interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:251
		_go_fuzz_dep_.CoverTab[142658]++
												v := reflect.ValueOf(value)
												switch reflect.TypeOf(args[n]).Kind() {
		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:254
			_go_fuzz_dep_.CoverTab[142660]++
													setSlice(args[n], v)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:255
			// _ = "end of CoverTab[142660]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:256
			_go_fuzz_dep_.CoverTab[142661]++
													reflect.ValueOf(args[n]).Elem().Set(v)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:257
			// _ = "end of CoverTab[142661]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:258
		// _ = "end of CoverTab[142658]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:258
		_go_fuzz_dep_.CoverTab[142659]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:259
		// _ = "end of CoverTab[142659]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:260
	// _ = "end of CoverTab[142647]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:260
	_go_fuzz_dep_.CoverTab[142648]++
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:261
	// _ = "end of CoverTab[142648]"
}

// isPreReq returns true if other is a direct or indirect prerequisite to c.
func (c *Call) isPreReq(other *Call) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:265
	_go_fuzz_dep_.CoverTab[142662]++
											for _, preReq := range c.preReqs {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:266
		_go_fuzz_dep_.CoverTab[142664]++
												if other == preReq || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:267
			_go_fuzz_dep_.CoverTab[142665]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:267
			return preReq.isPreReq(other)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:267
			// _ = "end of CoverTab[142665]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:267
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:267
			_go_fuzz_dep_.CoverTab[142666]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:268
			// _ = "end of CoverTab[142666]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:269
			_go_fuzz_dep_.CoverTab[142667]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:269
			// _ = "end of CoverTab[142667]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:269
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:269
		// _ = "end of CoverTab[142664]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:270
	// _ = "end of CoverTab[142662]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:270
	_go_fuzz_dep_.CoverTab[142663]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:271
	// _ = "end of CoverTab[142663]"
}

// After declares that the call may only match after preReq has been exhausted.
func (c *Call) After(preReq *Call) *Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:275
	_go_fuzz_dep_.CoverTab[142668]++
											c.t.Helper()

											if c == preReq {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:278
		_go_fuzz_dep_.CoverTab[142671]++
												c.t.Fatalf("A call isn't allowed to be its own prerequisite")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:279
		// _ = "end of CoverTab[142671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:280
		_go_fuzz_dep_.CoverTab[142672]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:280
		// _ = "end of CoverTab[142672]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:280
	// _ = "end of CoverTab[142668]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:280
	_go_fuzz_dep_.CoverTab[142669]++
											if preReq.isPreReq(c) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:281
		_go_fuzz_dep_.CoverTab[142673]++
												c.t.Fatalf("Loop in call order: %v is a prerequisite to %v (possibly indirectly).", c, preReq)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:282
		// _ = "end of CoverTab[142673]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:283
		_go_fuzz_dep_.CoverTab[142674]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:283
		// _ = "end of CoverTab[142674]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:283
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:283
	// _ = "end of CoverTab[142669]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:283
	_go_fuzz_dep_.CoverTab[142670]++

											c.preReqs = append(c.preReqs, preReq)
											return c
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:286
	// _ = "end of CoverTab[142670]"
}

// Returns true if the minimum number of calls have been made.
func (c *Call) satisfied() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:290
	_go_fuzz_dep_.CoverTab[142675]++
											return c.numCalls >= c.minCalls
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:291
	// _ = "end of CoverTab[142675]"
}

// Returns true if the maximum number of calls have been made.
func (c *Call) exhausted() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:295
	_go_fuzz_dep_.CoverTab[142676]++
											return c.numCalls >= c.maxCalls
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:296
	// _ = "end of CoverTab[142676]"
}

func (c *Call) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:299
	_go_fuzz_dep_.CoverTab[142677]++
											args := make([]string, len(c.args))
											for i, arg := range c.args {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:301
		_go_fuzz_dep_.CoverTab[142679]++
												args[i] = arg.String()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:302
		// _ = "end of CoverTab[142679]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:303
	// _ = "end of CoverTab[142677]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:303
	_go_fuzz_dep_.CoverTab[142678]++
											arguments := strings.Join(args, ", ")
											return fmt.Sprintf("%T.%v(%s) %s", c.receiver, c.method, arguments, c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:305
	// _ = "end of CoverTab[142678]"
}

// Tests if the given call matches the expected call.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:308
// If yes, returns nil. If no, returns error with message explaining why it does not match.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:310
func (c *Call) matches(args []interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:310
	_go_fuzz_dep_.CoverTab[142680]++
											if !c.methodType.IsVariadic() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:311
		_go_fuzz_dep_.CoverTab[142684]++
												if len(args) != len(c.args) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:312
			_go_fuzz_dep_.CoverTab[142686]++
													return fmt.Errorf("expected call at %s has the wrong number of arguments. Got: %d, want: %d",
				c.origin, len(args), len(c.args))
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:314
			// _ = "end of CoverTab[142686]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:315
			_go_fuzz_dep_.CoverTab[142687]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:315
			// _ = "end of CoverTab[142687]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:315
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:315
		// _ = "end of CoverTab[142684]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:315
		_go_fuzz_dep_.CoverTab[142685]++

												for i, m := range c.args {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:317
			_go_fuzz_dep_.CoverTab[142688]++
													if !m.Matches(args[i]) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:318
				_go_fuzz_dep_.CoverTab[142689]++
														return fmt.Errorf(
					"expected call at %s doesn't match the argument at index %d.\nGot: %v\nWant: %v",
					c.origin, i, formatGottenArg(m, args[i]), m,
				)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:322
				// _ = "end of CoverTab[142689]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:323
				_go_fuzz_dep_.CoverTab[142690]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:323
				// _ = "end of CoverTab[142690]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:323
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:323
			// _ = "end of CoverTab[142688]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:324
		// _ = "end of CoverTab[142685]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:325
		_go_fuzz_dep_.CoverTab[142691]++
												if len(c.args) < c.methodType.NumIn()-1 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:326
			_go_fuzz_dep_.CoverTab[142695]++
													return fmt.Errorf("expected call at %s has the wrong number of matchers. Got: %d, want: %d",
				c.origin, len(c.args), c.methodType.NumIn()-1)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:328
			// _ = "end of CoverTab[142695]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:329
			_go_fuzz_dep_.CoverTab[142696]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:329
			// _ = "end of CoverTab[142696]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:329
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:329
		// _ = "end of CoverTab[142691]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:329
		_go_fuzz_dep_.CoverTab[142692]++
												if len(c.args) != c.methodType.NumIn() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:330
			_go_fuzz_dep_.CoverTab[142697]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:330
			return len(args) != len(c.args)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:330
			// _ = "end of CoverTab[142697]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:330
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:330
			_go_fuzz_dep_.CoverTab[142698]++
													return fmt.Errorf("expected call at %s has the wrong number of arguments. Got: %d, want: %d",
				c.origin, len(args), len(c.args))
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:332
			// _ = "end of CoverTab[142698]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:333
			_go_fuzz_dep_.CoverTab[142699]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:333
			// _ = "end of CoverTab[142699]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:333
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:333
		// _ = "end of CoverTab[142692]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:333
		_go_fuzz_dep_.CoverTab[142693]++
												if len(args) < len(c.args)-1 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:334
			_go_fuzz_dep_.CoverTab[142700]++
													return fmt.Errorf("expected call at %s has the wrong number of arguments. Got: %d, want: greater than or equal to %d",
				c.origin, len(args), len(c.args)-1)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:336
			// _ = "end of CoverTab[142700]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:337
			_go_fuzz_dep_.CoverTab[142701]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:337
			// _ = "end of CoverTab[142701]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:337
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:337
		// _ = "end of CoverTab[142693]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:337
		_go_fuzz_dep_.CoverTab[142694]++

												for i, m := range c.args {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:339
			_go_fuzz_dep_.CoverTab[142702]++
													if i < c.methodType.NumIn()-1 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:340
				_go_fuzz_dep_.CoverTab[142707]++

														if !m.Matches(args[i]) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:342
					_go_fuzz_dep_.CoverTab[142709]++
															return fmt.Errorf("expected call at %s doesn't match the argument at index %s.\nGot: %v\nWant: %v",
						c.origin, strconv.Itoa(i), formatGottenArg(m, args[i]), m)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:344
					// _ = "end of CoverTab[142709]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:345
					_go_fuzz_dep_.CoverTab[142710]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:345
					// _ = "end of CoverTab[142710]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:345
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:345
				// _ = "end of CoverTab[142707]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:345
				_go_fuzz_dep_.CoverTab[142708]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:346
				// _ = "end of CoverTab[142708]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:347
				_go_fuzz_dep_.CoverTab[142711]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:347
				// _ = "end of CoverTab[142711]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:347
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:347
			// _ = "end of CoverTab[142702]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:347
			_go_fuzz_dep_.CoverTab[142703]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:351
			if i < len(c.args) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:351
				_go_fuzz_dep_.CoverTab[142712]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:351
				return i < len(args)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:351
				// _ = "end of CoverTab[142712]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:351
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:351
				_go_fuzz_dep_.CoverTab[142713]++
														if m.Matches(args[i]) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:352
					_go_fuzz_dep_.CoverTab[142714]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:358
					continue
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:358
					// _ = "end of CoverTab[142714]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:359
					_go_fuzz_dep_.CoverTab[142715]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:359
					// _ = "end of CoverTab[142715]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:359
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:359
				// _ = "end of CoverTab[142713]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:360
				_go_fuzz_dep_.CoverTab[142716]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:360
				// _ = "end of CoverTab[142716]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:360
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:360
			// _ = "end of CoverTab[142703]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:360
			_go_fuzz_dep_.CoverTab[142704]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:368
			vArgsType := c.methodType.In(c.methodType.NumIn() - 1)
			vArgs := reflect.MakeSlice(vArgsType, 0, len(args)-i)
			for _, arg := range args[i:] {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:370
				_go_fuzz_dep_.CoverTab[142717]++
														vArgs = reflect.Append(vArgs, reflect.ValueOf(arg))
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:371
				// _ = "end of CoverTab[142717]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:372
			// _ = "end of CoverTab[142704]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:372
			_go_fuzz_dep_.CoverTab[142705]++
													if m.Matches(vArgs.Interface()) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:373
				_go_fuzz_dep_.CoverTab[142718]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:378
				break
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:378
				// _ = "end of CoverTab[142718]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:379
				_go_fuzz_dep_.CoverTab[142719]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:379
				// _ = "end of CoverTab[142719]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:379
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:379
			// _ = "end of CoverTab[142705]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:379
			_go_fuzz_dep_.CoverTab[142706]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:387
			return fmt.Errorf("expected call at %s doesn't match the argument at index %s.\nGot: %v\nWant: %v",
				c.origin, strconv.Itoa(i), formatGottenArg(m, args[i:]), c.args[i])
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:388
			// _ = "end of CoverTab[142706]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:389
		// _ = "end of CoverTab[142694]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:390
	// _ = "end of CoverTab[142680]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:390
	_go_fuzz_dep_.CoverTab[142681]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:393
	for _, preReqCall := range c.preReqs {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:393
		_go_fuzz_dep_.CoverTab[142720]++
												if !preReqCall.satisfied() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:394
			_go_fuzz_dep_.CoverTab[142721]++
													return fmt.Errorf("expected call at %s doesn't have a prerequisite call satisfied:\n%v\nshould be called before:\n%v",
				c.origin, preReqCall, c)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:396
			// _ = "end of CoverTab[142721]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:397
			_go_fuzz_dep_.CoverTab[142722]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:397
			// _ = "end of CoverTab[142722]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:397
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:397
		// _ = "end of CoverTab[142720]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:398
	// _ = "end of CoverTab[142681]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:398
	_go_fuzz_dep_.CoverTab[142682]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:401
	if c.exhausted() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:401
		_go_fuzz_dep_.CoverTab[142723]++
												return fmt.Errorf("expected call at %s has already been called the max number of times", c.origin)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:402
		// _ = "end of CoverTab[142723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:403
		_go_fuzz_dep_.CoverTab[142724]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:403
		// _ = "end of CoverTab[142724]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:403
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:403
	// _ = "end of CoverTab[142682]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:403
	_go_fuzz_dep_.CoverTab[142683]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:405
	// _ = "end of CoverTab[142683]"
}

// dropPrereqs tells the expected Call to not re-check prerequisite calls any
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:408
// longer, and to return its current set.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:410
func (c *Call) dropPrereqs() (preReqs []*Call) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:410
	_go_fuzz_dep_.CoverTab[142725]++
											preReqs = c.preReqs
											c.preReqs = nil
											return
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:413
	// _ = "end of CoverTab[142725]"
}

func (c *Call) call() []func([]interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:416
	_go_fuzz_dep_.CoverTab[142726]++
											c.numCalls++
											return c.actions
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:418
	// _ = "end of CoverTab[142726]"
}

// InOrder declares that the given calls should occur in order.
func InOrder(calls ...*Call) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:422
	_go_fuzz_dep_.CoverTab[142727]++
											for i := 1; i < len(calls); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:423
		_go_fuzz_dep_.CoverTab[142728]++
												calls[i].After(calls[i-1])
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:424
		// _ = "end of CoverTab[142728]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:425
	// _ = "end of CoverTab[142727]"
}

func setSlice(arg interface{}, v reflect.Value) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:428
	_go_fuzz_dep_.CoverTab[142729]++
											va := reflect.ValueOf(arg)
											for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:430
		_go_fuzz_dep_.CoverTab[142730]++
												va.Index(i).Set(v.Index(i))
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:431
		// _ = "end of CoverTab[142730]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:432
	// _ = "end of CoverTab[142729]"
}

func (c *Call) addAction(action func([]interface{}) []interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:435
	_go_fuzz_dep_.CoverTab[142731]++
											c.actions = append(c.actions, action)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:436
	// _ = "end of CoverTab[142731]"
}

func formatGottenArg(m Matcher, arg interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:439
	_go_fuzz_dep_.CoverTab[142732]++
											got := fmt.Sprintf("%v (%T)", arg, arg)
											if gs, ok := m.(GotFormatter); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:441
		_go_fuzz_dep_.CoverTab[142734]++
												got = gs.Got(arg)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:442
		// _ = "end of CoverTab[142734]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:443
		_go_fuzz_dep_.CoverTab[142735]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:443
		// _ = "end of CoverTab[142735]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:443
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:443
	// _ = "end of CoverTab[142732]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:443
	_go_fuzz_dep_.CoverTab[142733]++
											return got
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:444
	// _ = "end of CoverTab[142733]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:445
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/call.go:445
var _ = _go_fuzz_dep_.CoverTab
