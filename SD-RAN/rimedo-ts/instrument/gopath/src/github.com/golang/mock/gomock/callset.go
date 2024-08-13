// Copyright 2011 Google Inc.
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

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
package gomock

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:15
)

import (
	"bytes"
	"errors"
	"fmt"
)

// callSet represents a set of expected calls, indexed by receiver and method
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:23
// name.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:25
type callSet struct {
	// Calls that are still expected.
	expected	map[callSetKey][]*Call
	// Calls that have been exhausted.
	exhausted	map[callSetKey][]*Call
}

// callSetKey is the key in the maps in callSet
type callSetKey struct {
	receiver	interface{}
	fname		string
}

func newCallSet() *callSet {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:38
	_go_fuzz_dep_.CoverTab[142736]++
											return &callSet{make(map[callSetKey][]*Call), make(map[callSetKey][]*Call)}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:39
	// _ = "end of CoverTab[142736]"
}

// Add adds a new expected call.
func (cs callSet) Add(call *Call) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:43
	_go_fuzz_dep_.CoverTab[142737]++
											key := callSetKey{call.receiver, call.method}
											m := cs.expected
											if call.exhausted() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:46
		_go_fuzz_dep_.CoverTab[142739]++
												m = cs.exhausted
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:47
		// _ = "end of CoverTab[142739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:48
		_go_fuzz_dep_.CoverTab[142740]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:48
		// _ = "end of CoverTab[142740]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:48
	// _ = "end of CoverTab[142737]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:48
	_go_fuzz_dep_.CoverTab[142738]++
											m[key] = append(m[key], call)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:49
	// _ = "end of CoverTab[142738]"
}

// Remove removes an expected call.
func (cs callSet) Remove(call *Call) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:53
	_go_fuzz_dep_.CoverTab[142741]++
											key := callSetKey{call.receiver, call.method}
											calls := cs.expected[key]
											for i, c := range calls {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:56
		_go_fuzz_dep_.CoverTab[142742]++
												if c == call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:57
			_go_fuzz_dep_.CoverTab[142743]++

													cs.expected[key] = append(calls[:i], calls[i+1:]...)
													cs.exhausted[key] = append(cs.exhausted[key], call)
													break
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:61
			// _ = "end of CoverTab[142743]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:62
			_go_fuzz_dep_.CoverTab[142744]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:62
			// _ = "end of CoverTab[142744]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:62
		// _ = "end of CoverTab[142742]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:63
	// _ = "end of CoverTab[142741]"
}

// FindMatch searches for a matching call. Returns error with explanation message if no call matched.
func (cs callSet) FindMatch(receiver interface{}, method string, args []interface{}) (*Call, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:67
	_go_fuzz_dep_.CoverTab[142745]++
											key := callSetKey{receiver, method}

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:71
	expected := cs.expected[key]
	var callsErrors bytes.Buffer
	for _, call := range expected {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:73
		_go_fuzz_dep_.CoverTab[142749]++
												err := call.matches(args)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:75
			_go_fuzz_dep_.CoverTab[142750]++
													_, _ = fmt.Fprintf(&callsErrors, "\n%v", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:76
			// _ = "end of CoverTab[142750]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:77
			_go_fuzz_dep_.CoverTab[142751]++
													return call, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:78
			// _ = "end of CoverTab[142751]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:79
		// _ = "end of CoverTab[142749]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:80
	// _ = "end of CoverTab[142745]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:80
	_go_fuzz_dep_.CoverTab[142746]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:84
	exhausted := cs.exhausted[key]
	for _, call := range exhausted {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:85
		_go_fuzz_dep_.CoverTab[142752]++
												if err := call.matches(args); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:86
			_go_fuzz_dep_.CoverTab[142754]++
													_, _ = fmt.Fprintf(&callsErrors, "\n%v", err)
													continue
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:88
			// _ = "end of CoverTab[142754]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:89
			_go_fuzz_dep_.CoverTab[142755]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:89
			// _ = "end of CoverTab[142755]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:89
		// _ = "end of CoverTab[142752]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:89
		_go_fuzz_dep_.CoverTab[142753]++
												_, _ = fmt.Fprintf(
			&callsErrors, "all expected calls for method %q have been exhausted", method,
		)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:92
		// _ = "end of CoverTab[142753]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:93
	// _ = "end of CoverTab[142746]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:93
	_go_fuzz_dep_.CoverTab[142747]++

											if len(expected)+len(exhausted) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:95
		_go_fuzz_dep_.CoverTab[142756]++
												_, _ = fmt.Fprintf(&callsErrors, "there are no expected calls of the method %q for that receiver", method)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:96
		// _ = "end of CoverTab[142756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:97
		_go_fuzz_dep_.CoverTab[142757]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:97
		// _ = "end of CoverTab[142757]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:97
	// _ = "end of CoverTab[142747]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:97
	_go_fuzz_dep_.CoverTab[142748]++

											return nil, errors.New(callsErrors.String())
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:99
	// _ = "end of CoverTab[142748]"
}

// Failures returns the calls that are not satisfied.
func (cs callSet) Failures() []*Call {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:103
	_go_fuzz_dep_.CoverTab[142758]++
											failures := make([]*Call, 0, len(cs.expected))
											for _, calls := range cs.expected {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:105
		_go_fuzz_dep_.CoverTab[142760]++
												for _, call := range calls {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:106
			_go_fuzz_dep_.CoverTab[142761]++
													if !call.satisfied() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:107
				_go_fuzz_dep_.CoverTab[142762]++
														failures = append(failures, call)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:108
				// _ = "end of CoverTab[142762]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:109
				_go_fuzz_dep_.CoverTab[142763]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:109
				// _ = "end of CoverTab[142763]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:109
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:109
			// _ = "end of CoverTab[142761]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:110
		// _ = "end of CoverTab[142760]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:111
	// _ = "end of CoverTab[142758]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:111
	_go_fuzz_dep_.CoverTab[142759]++
											return failures
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:112
	// _ = "end of CoverTab[142759]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/callset.go:113
var _ = _go_fuzz_dep_.CoverTab
