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

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
package gomock

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:15
)

import (
	"fmt"
	"reflect"
	"strings"
)

// A Matcher is a representation of a class of values.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:23
// It is used to represent the valid or expected arguments to a mocked method.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:25
type Matcher interface {
	// Matches returns whether x is a match.
	Matches(x interface{}) bool

	// String describes what the matcher matches.
	String() string
}

// WantFormatter modifies the given Matcher's String() method to the given
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:33
// Stringer. This allows for control on how the "Want" is formatted when
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:33
// printing .
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:36
func WantFormatter(s fmt.Stringer, m Matcher) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:36
	_go_fuzz_dep_.CoverTab[142833]++
											type matcher interface {
		Matches(x interface{}) bool
	}

	return struct {
		matcher
		fmt.Stringer
	}{
		matcher:	m,
		Stringer:	s,
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:47
	// _ = "end of CoverTab[142833]"
}

// StringerFunc type is an adapter to allow the use of ordinary functions as
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:50
// a Stringer. If f is a function with the appropriate signature,
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:50
// StringerFunc(f) is a Stringer that calls f.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:53
type StringerFunc func() string

// String implements fmt.Stringer.
func (f StringerFunc) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:56
	_go_fuzz_dep_.CoverTab[142834]++
											return f()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:57
	// _ = "end of CoverTab[142834]"
}

// GotFormatter is used to better print failure messages. If a matcher
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:60
// implements GotFormatter, it will use the result from Got when printing
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:60
// the failure message.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:63
type GotFormatter interface {
	// Got is invoked with the received value. The result is used when
	// printing the failure message.
	Got(got interface{}) string
}

// GotFormatterFunc type is an adapter to allow the use of ordinary
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:69
// functions as a GotFormatter. If f is a function with the appropriate
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:69
// signature, GotFormatterFunc(f) is a GotFormatter that calls f.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:72
type GotFormatterFunc func(got interface{}) string

// Got implements GotFormatter.
func (f GotFormatterFunc) Got(got interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:75
	_go_fuzz_dep_.CoverTab[142835]++
											return f(got)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:76
	// _ = "end of CoverTab[142835]"
}

// GotFormatterAdapter attaches a GotFormatter to a Matcher.
func GotFormatterAdapter(s GotFormatter, m Matcher) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:80
	_go_fuzz_dep_.CoverTab[142836]++
											return struct {
		GotFormatter
		Matcher
	}{
		GotFormatter:	s,
		Matcher:	m,
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:87
	// _ = "end of CoverTab[142836]"
}

type anyMatcher struct{}

func (anyMatcher) Matches(interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:92
	_go_fuzz_dep_.CoverTab[142837]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:93
	// _ = "end of CoverTab[142837]"
}

func (anyMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:96
	_go_fuzz_dep_.CoverTab[142838]++
											return "is anything"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:97
	// _ = "end of CoverTab[142838]"
}

type eqMatcher struct {
	x interface{}
}

func (e eqMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:104
	_go_fuzz_dep_.CoverTab[142839]++

											if e.x == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:106
		_go_fuzz_dep_.CoverTab[142842]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:106
		return x == nil
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:106
		// _ = "end of CoverTab[142842]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:106
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:106
		_go_fuzz_dep_.CoverTab[142843]++
												return reflect.DeepEqual(e.x, x)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:107
		// _ = "end of CoverTab[142843]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:108
		_go_fuzz_dep_.CoverTab[142844]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:108
		// _ = "end of CoverTab[142844]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:108
	// _ = "end of CoverTab[142839]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:108
	_go_fuzz_dep_.CoverTab[142840]++

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:111
	x1Val := reflect.ValueOf(e.x)
	x2Val := reflect.ValueOf(x)

	if x1Val.Type().AssignableTo(x2Val.Type()) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:114
		_go_fuzz_dep_.CoverTab[142845]++
												x1ValConverted := x1Val.Convert(x2Val.Type())
												return reflect.DeepEqual(x1ValConverted.Interface(), x2Val.Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:116
		// _ = "end of CoverTab[142845]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:117
		_go_fuzz_dep_.CoverTab[142846]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:117
		// _ = "end of CoverTab[142846]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:117
	// _ = "end of CoverTab[142840]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:117
	_go_fuzz_dep_.CoverTab[142841]++

											return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:119
	// _ = "end of CoverTab[142841]"
}

func (e eqMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:122
	_go_fuzz_dep_.CoverTab[142847]++
											return fmt.Sprintf("is equal to %v (%T)", e.x, e.x)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:123
	// _ = "end of CoverTab[142847]"
}

type nilMatcher struct{}

func (nilMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:128
	_go_fuzz_dep_.CoverTab[142848]++
											if x == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:129
		_go_fuzz_dep_.CoverTab[142851]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:130
		// _ = "end of CoverTab[142851]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:131
		_go_fuzz_dep_.CoverTab[142852]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:131
		// _ = "end of CoverTab[142852]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:131
	// _ = "end of CoverTab[142848]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:131
	_go_fuzz_dep_.CoverTab[142849]++

											v := reflect.ValueOf(x)
											switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:136
		_go_fuzz_dep_.CoverTab[142853]++
												return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:137
		// _ = "end of CoverTab[142853]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:137
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:137
		_go_fuzz_dep_.CoverTab[142854]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:137
		// _ = "end of CoverTab[142854]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:138
	// _ = "end of CoverTab[142849]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:138
	_go_fuzz_dep_.CoverTab[142850]++

											return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:140
	// _ = "end of CoverTab[142850]"
}

func (nilMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:143
	_go_fuzz_dep_.CoverTab[142855]++
											return "is nil"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:144
	// _ = "end of CoverTab[142855]"
}

type notMatcher struct {
	m Matcher
}

func (n notMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:151
	_go_fuzz_dep_.CoverTab[142856]++
											return !n.m.Matches(x)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:152
	// _ = "end of CoverTab[142856]"
}

func (n notMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:155
	_go_fuzz_dep_.CoverTab[142857]++
											return "not(" + n.m.String() + ")"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:156
	// _ = "end of CoverTab[142857]"
}

type assignableToTypeOfMatcher struct {
	targetType reflect.Type
}

func (m assignableToTypeOfMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:163
	_go_fuzz_dep_.CoverTab[142858]++
											return reflect.TypeOf(x).AssignableTo(m.targetType)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:164
	// _ = "end of CoverTab[142858]"
}

func (m assignableToTypeOfMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:167
	_go_fuzz_dep_.CoverTab[142859]++
											return "is assignable to " + m.targetType.Name()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:168
	// _ = "end of CoverTab[142859]"
}

type allMatcher struct {
	matchers []Matcher
}

func (am allMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:175
	_go_fuzz_dep_.CoverTab[142860]++
											for _, m := range am.matchers {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:176
		_go_fuzz_dep_.CoverTab[142862]++
												if !m.Matches(x) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:177
			_go_fuzz_dep_.CoverTab[142863]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:178
			// _ = "end of CoverTab[142863]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:179
			_go_fuzz_dep_.CoverTab[142864]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:179
			// _ = "end of CoverTab[142864]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:179
		// _ = "end of CoverTab[142862]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:180
	// _ = "end of CoverTab[142860]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:180
	_go_fuzz_dep_.CoverTab[142861]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:181
	// _ = "end of CoverTab[142861]"
}

func (am allMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:184
	_go_fuzz_dep_.CoverTab[142865]++
											ss := make([]string, 0, len(am.matchers))
											for _, matcher := range am.matchers {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:186
		_go_fuzz_dep_.CoverTab[142867]++
												ss = append(ss, matcher.String())
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:187
		// _ = "end of CoverTab[142867]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:188
	// _ = "end of CoverTab[142865]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:188
	_go_fuzz_dep_.CoverTab[142866]++
											return strings.Join(ss, "; ")
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:189
	// _ = "end of CoverTab[142866]"
}

type lenMatcher struct {
	i int
}

func (m lenMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:196
	_go_fuzz_dep_.CoverTab[142868]++
											v := reflect.ValueOf(x)
											switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:199
		_go_fuzz_dep_.CoverTab[142869]++
												return v.Len() == m.i
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:200
		// _ = "end of CoverTab[142869]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:201
		_go_fuzz_dep_.CoverTab[142870]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:202
		// _ = "end of CoverTab[142870]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:203
	// _ = "end of CoverTab[142868]"
}

func (m lenMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:206
	_go_fuzz_dep_.CoverTab[142871]++
											return fmt.Sprintf("has length %d", m.i)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:207
	// _ = "end of CoverTab[142871]"
}

type inAnyOrderMatcher struct {
	x interface{}
}

func (m inAnyOrderMatcher) Matches(x interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:214
	_go_fuzz_dep_.CoverTab[142872]++
											given, ok := m.prepareValue(x)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:216
		_go_fuzz_dep_.CoverTab[142879]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:217
		// _ = "end of CoverTab[142879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:218
		_go_fuzz_dep_.CoverTab[142880]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:218
		// _ = "end of CoverTab[142880]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:218
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:218
	// _ = "end of CoverTab[142872]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:218
	_go_fuzz_dep_.CoverTab[142873]++
											wanted, ok := m.prepareValue(m.x)
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:220
		_go_fuzz_dep_.CoverTab[142881]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:221
		// _ = "end of CoverTab[142881]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:222
		_go_fuzz_dep_.CoverTab[142882]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:222
		// _ = "end of CoverTab[142882]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:222
	// _ = "end of CoverTab[142873]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:222
	_go_fuzz_dep_.CoverTab[142874]++

											if given.Len() != wanted.Len() {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:224
		_go_fuzz_dep_.CoverTab[142883]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:225
		// _ = "end of CoverTab[142883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:226
		_go_fuzz_dep_.CoverTab[142884]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:226
		// _ = "end of CoverTab[142884]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:226
	// _ = "end of CoverTab[142874]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:226
	_go_fuzz_dep_.CoverTab[142875]++

											usedFromGiven := make([]bool, given.Len())
											foundFromWanted := make([]bool, wanted.Len())
											for i := 0; i < wanted.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:230
		_go_fuzz_dep_.CoverTab[142885]++
												wantedMatcher := Eq(wanted.Index(i).Interface())
												for j := 0; j < given.Len(); j++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:232
			_go_fuzz_dep_.CoverTab[142886]++
													if usedFromGiven[j] {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:233
				_go_fuzz_dep_.CoverTab[142888]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:234
				// _ = "end of CoverTab[142888]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:235
				_go_fuzz_dep_.CoverTab[142889]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:235
				// _ = "end of CoverTab[142889]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:235
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:235
			// _ = "end of CoverTab[142886]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:235
			_go_fuzz_dep_.CoverTab[142887]++
													if wantedMatcher.Matches(given.Index(j).Interface()) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:236
				_go_fuzz_dep_.CoverTab[142890]++
														foundFromWanted[i] = true
														usedFromGiven[j] = true
														break
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:239
				// _ = "end of CoverTab[142890]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:240
				_go_fuzz_dep_.CoverTab[142891]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:240
				// _ = "end of CoverTab[142891]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:240
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:240
			// _ = "end of CoverTab[142887]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:241
		// _ = "end of CoverTab[142885]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:242
	// _ = "end of CoverTab[142875]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:242
	_go_fuzz_dep_.CoverTab[142876]++

											missingFromWanted := 0
											for _, found := range foundFromWanted {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:245
		_go_fuzz_dep_.CoverTab[142892]++
												if !found {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:246
			_go_fuzz_dep_.CoverTab[142893]++
													missingFromWanted++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:247
			// _ = "end of CoverTab[142893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:248
			_go_fuzz_dep_.CoverTab[142894]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:248
			// _ = "end of CoverTab[142894]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:248
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:248
		// _ = "end of CoverTab[142892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:249
	// _ = "end of CoverTab[142876]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:249
	_go_fuzz_dep_.CoverTab[142877]++
											extraInGiven := 0
											for _, used := range usedFromGiven {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:251
		_go_fuzz_dep_.CoverTab[142895]++
												if !used {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:252
			_go_fuzz_dep_.CoverTab[142896]++
													extraInGiven++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:253
			// _ = "end of CoverTab[142896]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:254
			_go_fuzz_dep_.CoverTab[142897]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:254
			// _ = "end of CoverTab[142897]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:254
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:254
		// _ = "end of CoverTab[142895]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:255
	// _ = "end of CoverTab[142877]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:255
	_go_fuzz_dep_.CoverTab[142878]++

											return extraInGiven == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:257
		_go_fuzz_dep_.CoverTab[142898]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:257
		return missingFromWanted == 0
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:257
		// _ = "end of CoverTab[142898]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:257
	}()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:257
	// _ = "end of CoverTab[142878]"
}

func (m inAnyOrderMatcher) prepareValue(x interface{}) (reflect.Value, bool) {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:260
	_go_fuzz_dep_.CoverTab[142899]++
											xValue := reflect.ValueOf(x)
											switch xValue.Kind() {
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:263
		_go_fuzz_dep_.CoverTab[142900]++
												return xValue, true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:264
		// _ = "end of CoverTab[142900]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:265
		_go_fuzz_dep_.CoverTab[142901]++
												return reflect.Value{}, false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:266
		// _ = "end of CoverTab[142901]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:267
	// _ = "end of CoverTab[142899]"
}

func (m inAnyOrderMatcher) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:270
	_go_fuzz_dep_.CoverTab[142902]++
											return fmt.Sprintf("has the same elements as %v", m.x)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:271
	// _ = "end of CoverTab[142902]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:276
// All returns a composite Matcher that returns true if and only all of the
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:276
// matchers return true.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:278
func All(ms ...Matcher) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:278
	_go_fuzz_dep_.CoverTab[142903]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:278
	return allMatcher{ms}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:278
	// _ = "end of CoverTab[142903]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:278
}

// Any returns a matcher that always matches.
func Any() Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:281
	_go_fuzz_dep_.CoverTab[142904]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:281
	return anyMatcher{}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:281
	// _ = "end of CoverTab[142904]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:281
}

// Eq returns a matcher that matches on equality.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:283
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:283
// Example usage:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:283
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:283
//	Eq(5).Matches(5) // returns true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:283
//	Eq(5).Matches(4) // returns false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:288
func Eq(x interface{}) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:288
	_go_fuzz_dep_.CoverTab[142905]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:288
	return eqMatcher{x}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:288
	// _ = "end of CoverTab[142905]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:288
}

// Len returns a matcher that matches on length. This matcher returns false if
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:290
// is compared to a type that is not an array, chan, map, slice, or string.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:292
func Len(i int) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:292
	_go_fuzz_dep_.CoverTab[142906]++
											return lenMatcher{i}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:293
	// _ = "end of CoverTab[142906]"
}

// Nil returns a matcher that matches if the received value is nil.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
// Example usage:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
//	var x *bytes.Buffer
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
//	Nil().Matches(x) // returns true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
//	x = &bytes.Buffer{}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:296
//	Nil().Matches(x) // returns false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:303
func Nil() Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:303
	_go_fuzz_dep_.CoverTab[142907]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:303
	return nilMatcher{}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:303
	// _ = "end of CoverTab[142907]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:303
}

// Not reverses the results of its given child matcher.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:305
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:305
// Example usage:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:305
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:305
//	Not(Eq(5)).Matches(4) // returns true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:305
//	Not(Eq(5)).Matches(5) // returns false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:310
func Not(x interface{}) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:310
	_go_fuzz_dep_.CoverTab[142908]++
											if m, ok := x.(Matcher); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:311
		_go_fuzz_dep_.CoverTab[142910]++
												return notMatcher{m}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:312
		// _ = "end of CoverTab[142910]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:313
		_go_fuzz_dep_.CoverTab[142911]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:313
		// _ = "end of CoverTab[142911]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:313
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:313
	// _ = "end of CoverTab[142908]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:313
	_go_fuzz_dep_.CoverTab[142909]++
											return notMatcher{Eq(x)}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:314
	// _ = "end of CoverTab[142909]"
}

// AssignableToTypeOf is a Matcher that matches if the parameter to the mock
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
// function is assignable to the type of the parameter to this function.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
// Example usage:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//	var s fmt.Stringer = &bytes.Buffer{}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//	AssignableToTypeOf(s).Matches(time.Second) // returns true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//	AssignableToTypeOf(s).Matches(99) // returns false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//	var ctx = reflect.TypeOf((*context.Context)(nil)).Elem()
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:317
//	AssignableToTypeOf(ctx).Matches(context.Background()) // returns true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:327
func AssignableToTypeOf(x interface{}) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:327
	_go_fuzz_dep_.CoverTab[142912]++
											if xt, ok := x.(reflect.Type); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:328
		_go_fuzz_dep_.CoverTab[142914]++
												return assignableToTypeOfMatcher{xt}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:329
		// _ = "end of CoverTab[142914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:330
		_go_fuzz_dep_.CoverTab[142915]++
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:330
		// _ = "end of CoverTab[142915]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:330
	// _ = "end of CoverTab[142912]"
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:330
	_go_fuzz_dep_.CoverTab[142913]++
											return assignableToTypeOfMatcher{reflect.TypeOf(x)}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:331
	// _ = "end of CoverTab[142913]"
}

// InAnyOrder is a Matcher that returns true for collections of the same elements ignoring the order.
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:334
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:334
// Example usage:
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:334
//
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:334
//	InAnyOrder([]int{1, 2, 3}).Matches([]int{1, 3, 2}) // returns true
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:334
//	InAnyOrder([]int{1, 2, 3}).Matches([]int{1, 2}) // returns false
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:339
func InAnyOrder(x interface{}) Matcher {
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:339
	_go_fuzz_dep_.CoverTab[142916]++
											return inAnyOrderMatcher{x}
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:340
	// _ = "end of CoverTab[142916]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:341
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/mock@v1.6.0/gomock/matchers.go:341
var _ = _go_fuzz_dep_.CoverTab
