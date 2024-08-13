// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/text/template/funcs.go:5
package template

//line /usr/local/go/src/text/template/funcs.go:5
import (
//line /usr/local/go/src/text/template/funcs.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/template/funcs.go:5
)
//line /usr/local/go/src/text/template/funcs.go:5
import (
//line /usr/local/go/src/text/template/funcs.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/template/funcs.go:5
)

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// FuncMap is the type of the map defining the mapping from names to functions.
//line /usr/local/go/src/text/template/funcs.go:19
// Each function must have either a single return value, or two return values of
//line /usr/local/go/src/text/template/funcs.go:19
// which the second has type error. In that case, if the second (error)
//line /usr/local/go/src/text/template/funcs.go:19
// return value evaluates to non-nil during execution, execution terminates and
//line /usr/local/go/src/text/template/funcs.go:19
// Execute returns that error.
//line /usr/local/go/src/text/template/funcs.go:19
//
//line /usr/local/go/src/text/template/funcs.go:19
// Errors returned by Execute wrap the underlying error; call errors.As to
//line /usr/local/go/src/text/template/funcs.go:19
// uncover them.
//line /usr/local/go/src/text/template/funcs.go:19
//
//line /usr/local/go/src/text/template/funcs.go:19
// When template execution invokes a function with an argument list, that list
//line /usr/local/go/src/text/template/funcs.go:19
// must be assignable to the function's parameter types. Functions meant to
//line /usr/local/go/src/text/template/funcs.go:19
// apply to arguments of arbitrary type can use parameters of type interface{} or
//line /usr/local/go/src/text/template/funcs.go:19
// of type reflect.Value. Similarly, functions meant to return a result of arbitrary
//line /usr/local/go/src/text/template/funcs.go:19
// type can return interface{} or reflect.Value.
//line /usr/local/go/src/text/template/funcs.go:33
type FuncMap map[string]any

// builtins returns the FuncMap.
//line /usr/local/go/src/text/template/funcs.go:35
// It is not a global variable so the linker can dead code eliminate
//line /usr/local/go/src/text/template/funcs.go:35
// more when this isn't called. See golang.org/issue/36021.
//line /usr/local/go/src/text/template/funcs.go:35
// TODO: revert this back to a global map once golang.org/issue/2559 is fixed.
//line /usr/local/go/src/text/template/funcs.go:39
func builtins() FuncMap {
//line /usr/local/go/src/text/template/funcs.go:39
	_go_fuzz_dep_.CoverTab[30191]++
							return FuncMap{
								"and":		and,
								"call":		call,
								"html":		HTMLEscaper,
								"index":	index,
								"slice":	slice,
								"js":		JSEscaper,
								"len":		length,
								"not":		not,
								"or":		or,
								"print":	fmt.Sprint,
								"printf":	fmt.Sprintf,
								"println":	fmt.Sprintln,
								"urlquery":	URLQueryEscaper,

//line /usr/local/go/src/text/template/funcs.go:56
		"eq":	eq,
								"ge":	ge,
								"gt":	gt,
								"le":	le,
								"lt":	lt,
								"ne":	ne,
	}
//line /usr/local/go/src/text/template/funcs.go:62
	// _ = "end of CoverTab[30191]"
}

var builtinFuncsOnce struct {
	sync.Once
	v	map[string]reflect.Value
}

// builtinFuncsOnce lazily computes & caches the builtinFuncs map.
//line /usr/local/go/src/text/template/funcs.go:70
// TODO: revert this back to a global map once golang.org/issue/2559 is fixed.
//line /usr/local/go/src/text/template/funcs.go:72
func builtinFuncs() map[string]reflect.Value {
//line /usr/local/go/src/text/template/funcs.go:72
	_go_fuzz_dep_.CoverTab[30192]++
							builtinFuncsOnce.Do(func() {
//line /usr/local/go/src/text/template/funcs.go:73
		_go_fuzz_dep_.CoverTab[30194]++
								builtinFuncsOnce.v = createValueFuncs(builtins())
//line /usr/local/go/src/text/template/funcs.go:74
		// _ = "end of CoverTab[30194]"
	})
//line /usr/local/go/src/text/template/funcs.go:75
	// _ = "end of CoverTab[30192]"
//line /usr/local/go/src/text/template/funcs.go:75
	_go_fuzz_dep_.CoverTab[30193]++
							return builtinFuncsOnce.v
//line /usr/local/go/src/text/template/funcs.go:76
	// _ = "end of CoverTab[30193]"
}

// createValueFuncs turns a FuncMap into a map[string]reflect.Value
func createValueFuncs(funcMap FuncMap) map[string]reflect.Value {
//line /usr/local/go/src/text/template/funcs.go:80
	_go_fuzz_dep_.CoverTab[30195]++
							m := make(map[string]reflect.Value)
							addValueFuncs(m, funcMap)
							return m
//line /usr/local/go/src/text/template/funcs.go:83
	// _ = "end of CoverTab[30195]"
}

// addValueFuncs adds to values the functions in funcs, converting them to reflect.Values.
func addValueFuncs(out map[string]reflect.Value, in FuncMap) {
//line /usr/local/go/src/text/template/funcs.go:87
	_go_fuzz_dep_.CoverTab[30196]++
							for name, fn := range in {
//line /usr/local/go/src/text/template/funcs.go:88
		_go_fuzz_dep_.CoverTab[30197]++
								if !goodName(name) {
//line /usr/local/go/src/text/template/funcs.go:89
			_go_fuzz_dep_.CoverTab[30201]++
									panic(fmt.Errorf("function name %q is not a valid identifier", name))
//line /usr/local/go/src/text/template/funcs.go:90
			// _ = "end of CoverTab[30201]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:91
			_go_fuzz_dep_.CoverTab[30202]++
//line /usr/local/go/src/text/template/funcs.go:91
			// _ = "end of CoverTab[30202]"
//line /usr/local/go/src/text/template/funcs.go:91
		}
//line /usr/local/go/src/text/template/funcs.go:91
		// _ = "end of CoverTab[30197]"
//line /usr/local/go/src/text/template/funcs.go:91
		_go_fuzz_dep_.CoverTab[30198]++
								v := reflect.ValueOf(fn)
								if v.Kind() != reflect.Func {
//line /usr/local/go/src/text/template/funcs.go:93
			_go_fuzz_dep_.CoverTab[30203]++
									panic("value for " + name + " not a function")
//line /usr/local/go/src/text/template/funcs.go:94
			// _ = "end of CoverTab[30203]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:95
			_go_fuzz_dep_.CoverTab[30204]++
//line /usr/local/go/src/text/template/funcs.go:95
			// _ = "end of CoverTab[30204]"
//line /usr/local/go/src/text/template/funcs.go:95
		}
//line /usr/local/go/src/text/template/funcs.go:95
		// _ = "end of CoverTab[30198]"
//line /usr/local/go/src/text/template/funcs.go:95
		_go_fuzz_dep_.CoverTab[30199]++
								if !goodFunc(v.Type()) {
//line /usr/local/go/src/text/template/funcs.go:96
			_go_fuzz_dep_.CoverTab[30205]++
									panic(fmt.Errorf("can't install method/function %q with %d results", name, v.Type().NumOut()))
//line /usr/local/go/src/text/template/funcs.go:97
			// _ = "end of CoverTab[30205]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:98
			_go_fuzz_dep_.CoverTab[30206]++
//line /usr/local/go/src/text/template/funcs.go:98
			// _ = "end of CoverTab[30206]"
//line /usr/local/go/src/text/template/funcs.go:98
		}
//line /usr/local/go/src/text/template/funcs.go:98
		// _ = "end of CoverTab[30199]"
//line /usr/local/go/src/text/template/funcs.go:98
		_go_fuzz_dep_.CoverTab[30200]++
								out[name] = v
//line /usr/local/go/src/text/template/funcs.go:99
		// _ = "end of CoverTab[30200]"
	}
//line /usr/local/go/src/text/template/funcs.go:100
	// _ = "end of CoverTab[30196]"
}

// addFuncs adds to values the functions in funcs. It does no checking of the input -
//line /usr/local/go/src/text/template/funcs.go:103
// call addValueFuncs first.
//line /usr/local/go/src/text/template/funcs.go:105
func addFuncs(out, in FuncMap) {
//line /usr/local/go/src/text/template/funcs.go:105
	_go_fuzz_dep_.CoverTab[30207]++
							for name, fn := range in {
//line /usr/local/go/src/text/template/funcs.go:106
		_go_fuzz_dep_.CoverTab[30208]++
								out[name] = fn
//line /usr/local/go/src/text/template/funcs.go:107
		// _ = "end of CoverTab[30208]"
	}
//line /usr/local/go/src/text/template/funcs.go:108
	// _ = "end of CoverTab[30207]"
}

// goodFunc reports whether the function or method has the right result signature.
func goodFunc(typ reflect.Type) bool {
//line /usr/local/go/src/text/template/funcs.go:112
	_go_fuzz_dep_.CoverTab[30209]++

							switch {
	case typ.NumOut() == 1:
//line /usr/local/go/src/text/template/funcs.go:115
		_go_fuzz_dep_.CoverTab[30211]++
								return true
//line /usr/local/go/src/text/template/funcs.go:116
		// _ = "end of CoverTab[30211]"
	case typ.NumOut() == 2 && func() bool {
//line /usr/local/go/src/text/template/funcs.go:117
		_go_fuzz_dep_.CoverTab[30214]++
//line /usr/local/go/src/text/template/funcs.go:117
		return typ.Out(1) == errorType
//line /usr/local/go/src/text/template/funcs.go:117
		// _ = "end of CoverTab[30214]"
//line /usr/local/go/src/text/template/funcs.go:117
	}():
//line /usr/local/go/src/text/template/funcs.go:117
		_go_fuzz_dep_.CoverTab[30212]++
								return true
//line /usr/local/go/src/text/template/funcs.go:118
		// _ = "end of CoverTab[30212]"
//line /usr/local/go/src/text/template/funcs.go:118
	default:
//line /usr/local/go/src/text/template/funcs.go:118
		_go_fuzz_dep_.CoverTab[30213]++
//line /usr/local/go/src/text/template/funcs.go:118
		// _ = "end of CoverTab[30213]"
	}
//line /usr/local/go/src/text/template/funcs.go:119
	// _ = "end of CoverTab[30209]"
//line /usr/local/go/src/text/template/funcs.go:119
	_go_fuzz_dep_.CoverTab[30210]++
							return false
//line /usr/local/go/src/text/template/funcs.go:120
	// _ = "end of CoverTab[30210]"
}

// goodName reports whether the function name is a valid identifier.
func goodName(name string) bool {
//line /usr/local/go/src/text/template/funcs.go:124
	_go_fuzz_dep_.CoverTab[30215]++
							if name == "" {
//line /usr/local/go/src/text/template/funcs.go:125
		_go_fuzz_dep_.CoverTab[30218]++
								return false
//line /usr/local/go/src/text/template/funcs.go:126
		// _ = "end of CoverTab[30218]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:127
		_go_fuzz_dep_.CoverTab[30219]++
//line /usr/local/go/src/text/template/funcs.go:127
		// _ = "end of CoverTab[30219]"
//line /usr/local/go/src/text/template/funcs.go:127
	}
//line /usr/local/go/src/text/template/funcs.go:127
	// _ = "end of CoverTab[30215]"
//line /usr/local/go/src/text/template/funcs.go:127
	_go_fuzz_dep_.CoverTab[30216]++
							for i, r := range name {
//line /usr/local/go/src/text/template/funcs.go:128
		_go_fuzz_dep_.CoverTab[30220]++
								switch {
		case r == '_':
//line /usr/local/go/src/text/template/funcs.go:130
			_go_fuzz_dep_.CoverTab[30221]++
//line /usr/local/go/src/text/template/funcs.go:130
			// _ = "end of CoverTab[30221]"
		case i == 0 && func() bool {
//line /usr/local/go/src/text/template/funcs.go:131
			_go_fuzz_dep_.CoverTab[30225]++
//line /usr/local/go/src/text/template/funcs.go:131
			return !unicode.IsLetter(r)
//line /usr/local/go/src/text/template/funcs.go:131
			// _ = "end of CoverTab[30225]"
//line /usr/local/go/src/text/template/funcs.go:131
		}():
//line /usr/local/go/src/text/template/funcs.go:131
			_go_fuzz_dep_.CoverTab[30222]++
									return false
//line /usr/local/go/src/text/template/funcs.go:132
			// _ = "end of CoverTab[30222]"
		case !unicode.IsLetter(r) && func() bool {
//line /usr/local/go/src/text/template/funcs.go:133
			_go_fuzz_dep_.CoverTab[30226]++
//line /usr/local/go/src/text/template/funcs.go:133
			return !unicode.IsDigit(r)
//line /usr/local/go/src/text/template/funcs.go:133
			// _ = "end of CoverTab[30226]"
//line /usr/local/go/src/text/template/funcs.go:133
		}():
//line /usr/local/go/src/text/template/funcs.go:133
			_go_fuzz_dep_.CoverTab[30223]++
									return false
//line /usr/local/go/src/text/template/funcs.go:134
			// _ = "end of CoverTab[30223]"
//line /usr/local/go/src/text/template/funcs.go:134
		default:
//line /usr/local/go/src/text/template/funcs.go:134
			_go_fuzz_dep_.CoverTab[30224]++
//line /usr/local/go/src/text/template/funcs.go:134
			// _ = "end of CoverTab[30224]"
		}
//line /usr/local/go/src/text/template/funcs.go:135
		// _ = "end of CoverTab[30220]"
	}
//line /usr/local/go/src/text/template/funcs.go:136
	// _ = "end of CoverTab[30216]"
//line /usr/local/go/src/text/template/funcs.go:136
	_go_fuzz_dep_.CoverTab[30217]++
							return true
//line /usr/local/go/src/text/template/funcs.go:137
	// _ = "end of CoverTab[30217]"
}

// findFunction looks for a function in the template, and global map.
func findFunction(name string, tmpl *Template) (v reflect.Value, isBuiltin, ok bool) {
//line /usr/local/go/src/text/template/funcs.go:141
	_go_fuzz_dep_.CoverTab[30227]++
							if tmpl != nil && func() bool {
//line /usr/local/go/src/text/template/funcs.go:142
		_go_fuzz_dep_.CoverTab[30230]++
//line /usr/local/go/src/text/template/funcs.go:142
		return tmpl.common != nil
//line /usr/local/go/src/text/template/funcs.go:142
		// _ = "end of CoverTab[30230]"
//line /usr/local/go/src/text/template/funcs.go:142
	}() {
//line /usr/local/go/src/text/template/funcs.go:142
		_go_fuzz_dep_.CoverTab[30231]++
								tmpl.muFuncs.RLock()
								defer tmpl.muFuncs.RUnlock()
								if fn := tmpl.execFuncs[name]; fn.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:145
			_go_fuzz_dep_.CoverTab[30232]++
									return fn, false, true
//line /usr/local/go/src/text/template/funcs.go:146
			// _ = "end of CoverTab[30232]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:147
			_go_fuzz_dep_.CoverTab[30233]++
//line /usr/local/go/src/text/template/funcs.go:147
			// _ = "end of CoverTab[30233]"
//line /usr/local/go/src/text/template/funcs.go:147
		}
//line /usr/local/go/src/text/template/funcs.go:147
		// _ = "end of CoverTab[30231]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:148
		_go_fuzz_dep_.CoverTab[30234]++
//line /usr/local/go/src/text/template/funcs.go:148
		// _ = "end of CoverTab[30234]"
//line /usr/local/go/src/text/template/funcs.go:148
	}
//line /usr/local/go/src/text/template/funcs.go:148
	// _ = "end of CoverTab[30227]"
//line /usr/local/go/src/text/template/funcs.go:148
	_go_fuzz_dep_.CoverTab[30228]++
							if fn := builtinFuncs()[name]; fn.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:149
		_go_fuzz_dep_.CoverTab[30235]++
								return fn, true, true
//line /usr/local/go/src/text/template/funcs.go:150
		// _ = "end of CoverTab[30235]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:151
		_go_fuzz_dep_.CoverTab[30236]++
//line /usr/local/go/src/text/template/funcs.go:151
		// _ = "end of CoverTab[30236]"
//line /usr/local/go/src/text/template/funcs.go:151
	}
//line /usr/local/go/src/text/template/funcs.go:151
	// _ = "end of CoverTab[30228]"
//line /usr/local/go/src/text/template/funcs.go:151
	_go_fuzz_dep_.CoverTab[30229]++
							return reflect.Value{}, false, false
//line /usr/local/go/src/text/template/funcs.go:152
	// _ = "end of CoverTab[30229]"
}

// prepareArg checks if value can be used as an argument of type argType, and
//line /usr/local/go/src/text/template/funcs.go:155
// converts an invalid value to appropriate zero if possible.
//line /usr/local/go/src/text/template/funcs.go:157
func prepareArg(value reflect.Value, argType reflect.Type) (reflect.Value, error) {
//line /usr/local/go/src/text/template/funcs.go:157
	_go_fuzz_dep_.CoverTab[30237]++
							if !value.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:158
		_go_fuzz_dep_.CoverTab[30241]++
								if !canBeNil(argType) {
//line /usr/local/go/src/text/template/funcs.go:159
			_go_fuzz_dep_.CoverTab[30243]++
									return reflect.Value{}, fmt.Errorf("value is nil; should be of type %s", argType)
//line /usr/local/go/src/text/template/funcs.go:160
			// _ = "end of CoverTab[30243]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:161
			_go_fuzz_dep_.CoverTab[30244]++
//line /usr/local/go/src/text/template/funcs.go:161
			// _ = "end of CoverTab[30244]"
//line /usr/local/go/src/text/template/funcs.go:161
		}
//line /usr/local/go/src/text/template/funcs.go:161
		// _ = "end of CoverTab[30241]"
//line /usr/local/go/src/text/template/funcs.go:161
		_go_fuzz_dep_.CoverTab[30242]++
								value = reflect.Zero(argType)
//line /usr/local/go/src/text/template/funcs.go:162
		// _ = "end of CoverTab[30242]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:163
		_go_fuzz_dep_.CoverTab[30245]++
//line /usr/local/go/src/text/template/funcs.go:163
		// _ = "end of CoverTab[30245]"
//line /usr/local/go/src/text/template/funcs.go:163
	}
//line /usr/local/go/src/text/template/funcs.go:163
	// _ = "end of CoverTab[30237]"
//line /usr/local/go/src/text/template/funcs.go:163
	_go_fuzz_dep_.CoverTab[30238]++
							if value.Type().AssignableTo(argType) {
//line /usr/local/go/src/text/template/funcs.go:164
		_go_fuzz_dep_.CoverTab[30246]++
								return value, nil
//line /usr/local/go/src/text/template/funcs.go:165
		// _ = "end of CoverTab[30246]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:166
		_go_fuzz_dep_.CoverTab[30247]++
//line /usr/local/go/src/text/template/funcs.go:166
		// _ = "end of CoverTab[30247]"
//line /usr/local/go/src/text/template/funcs.go:166
	}
//line /usr/local/go/src/text/template/funcs.go:166
	// _ = "end of CoverTab[30238]"
//line /usr/local/go/src/text/template/funcs.go:166
	_go_fuzz_dep_.CoverTab[30239]++
							if intLike(value.Kind()) && func() bool {
//line /usr/local/go/src/text/template/funcs.go:167
		_go_fuzz_dep_.CoverTab[30248]++
//line /usr/local/go/src/text/template/funcs.go:167
		return intLike(argType.Kind())
//line /usr/local/go/src/text/template/funcs.go:167
		// _ = "end of CoverTab[30248]"
//line /usr/local/go/src/text/template/funcs.go:167
	}() && func() bool {
//line /usr/local/go/src/text/template/funcs.go:167
		_go_fuzz_dep_.CoverTab[30249]++
//line /usr/local/go/src/text/template/funcs.go:167
		return value.Type().ConvertibleTo(argType)
//line /usr/local/go/src/text/template/funcs.go:167
		// _ = "end of CoverTab[30249]"
//line /usr/local/go/src/text/template/funcs.go:167
	}() {
//line /usr/local/go/src/text/template/funcs.go:167
		_go_fuzz_dep_.CoverTab[30250]++
								value = value.Convert(argType)
								return value, nil
//line /usr/local/go/src/text/template/funcs.go:169
		// _ = "end of CoverTab[30250]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:170
		_go_fuzz_dep_.CoverTab[30251]++
//line /usr/local/go/src/text/template/funcs.go:170
		// _ = "end of CoverTab[30251]"
//line /usr/local/go/src/text/template/funcs.go:170
	}
//line /usr/local/go/src/text/template/funcs.go:170
	// _ = "end of CoverTab[30239]"
//line /usr/local/go/src/text/template/funcs.go:170
	_go_fuzz_dep_.CoverTab[30240]++
							return reflect.Value{}, fmt.Errorf("value has type %s; should be %s", value.Type(), argType)
//line /usr/local/go/src/text/template/funcs.go:171
	// _ = "end of CoverTab[30240]"
}

func intLike(typ reflect.Kind) bool {
//line /usr/local/go/src/text/template/funcs.go:174
	_go_fuzz_dep_.CoverTab[30252]++
							switch typ {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/text/template/funcs.go:176
		_go_fuzz_dep_.CoverTab[30254]++
								return true
//line /usr/local/go/src/text/template/funcs.go:177
		// _ = "end of CoverTab[30254]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/text/template/funcs.go:178
		_go_fuzz_dep_.CoverTab[30255]++
								return true
//line /usr/local/go/src/text/template/funcs.go:179
		// _ = "end of CoverTab[30255]"
//line /usr/local/go/src/text/template/funcs.go:179
	default:
//line /usr/local/go/src/text/template/funcs.go:179
		_go_fuzz_dep_.CoverTab[30256]++
//line /usr/local/go/src/text/template/funcs.go:179
		// _ = "end of CoverTab[30256]"
	}
//line /usr/local/go/src/text/template/funcs.go:180
	// _ = "end of CoverTab[30252]"
//line /usr/local/go/src/text/template/funcs.go:180
	_go_fuzz_dep_.CoverTab[30253]++
							return false
//line /usr/local/go/src/text/template/funcs.go:181
	// _ = "end of CoverTab[30253]"
}

// indexArg checks if a reflect.Value can be used as an index, and converts it to int if possible.
func indexArg(index reflect.Value, cap int) (int, error) {
//line /usr/local/go/src/text/template/funcs.go:185
	_go_fuzz_dep_.CoverTab[30257]++
							var x int64
							switch index.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/text/template/funcs.go:188
		_go_fuzz_dep_.CoverTab[30260]++
								x = index.Int()
//line /usr/local/go/src/text/template/funcs.go:189
		// _ = "end of CoverTab[30260]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/text/template/funcs.go:190
		_go_fuzz_dep_.CoverTab[30261]++
								x = int64(index.Uint())
//line /usr/local/go/src/text/template/funcs.go:191
		// _ = "end of CoverTab[30261]"
	case reflect.Invalid:
//line /usr/local/go/src/text/template/funcs.go:192
		_go_fuzz_dep_.CoverTab[30262]++
								return 0, fmt.Errorf("cannot index slice/array with nil")
//line /usr/local/go/src/text/template/funcs.go:193
		// _ = "end of CoverTab[30262]"
	default:
//line /usr/local/go/src/text/template/funcs.go:194
		_go_fuzz_dep_.CoverTab[30263]++
								return 0, fmt.Errorf("cannot index slice/array with type %s", index.Type())
//line /usr/local/go/src/text/template/funcs.go:195
		// _ = "end of CoverTab[30263]"
	}
//line /usr/local/go/src/text/template/funcs.go:196
	// _ = "end of CoverTab[30257]"
//line /usr/local/go/src/text/template/funcs.go:196
	_go_fuzz_dep_.CoverTab[30258]++
							if x < 0 || func() bool {
//line /usr/local/go/src/text/template/funcs.go:197
		_go_fuzz_dep_.CoverTab[30264]++
//line /usr/local/go/src/text/template/funcs.go:197
		return int(x) < 0
//line /usr/local/go/src/text/template/funcs.go:197
		// _ = "end of CoverTab[30264]"
//line /usr/local/go/src/text/template/funcs.go:197
	}() || func() bool {
//line /usr/local/go/src/text/template/funcs.go:197
		_go_fuzz_dep_.CoverTab[30265]++
//line /usr/local/go/src/text/template/funcs.go:197
		return int(x) > cap
//line /usr/local/go/src/text/template/funcs.go:197
		// _ = "end of CoverTab[30265]"
//line /usr/local/go/src/text/template/funcs.go:197
	}() {
//line /usr/local/go/src/text/template/funcs.go:197
		_go_fuzz_dep_.CoverTab[30266]++
								return 0, fmt.Errorf("index out of range: %d", x)
//line /usr/local/go/src/text/template/funcs.go:198
		// _ = "end of CoverTab[30266]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:199
		_go_fuzz_dep_.CoverTab[30267]++
//line /usr/local/go/src/text/template/funcs.go:199
		// _ = "end of CoverTab[30267]"
//line /usr/local/go/src/text/template/funcs.go:199
	}
//line /usr/local/go/src/text/template/funcs.go:199
	// _ = "end of CoverTab[30258]"
//line /usr/local/go/src/text/template/funcs.go:199
	_go_fuzz_dep_.CoverTab[30259]++
							return int(x), nil
//line /usr/local/go/src/text/template/funcs.go:200
	// _ = "end of CoverTab[30259]"
}

//line /usr/local/go/src/text/template/funcs.go:205
// index returns the result of indexing its first argument by the following
//line /usr/local/go/src/text/template/funcs.go:205
// arguments. Thus "index x 1 2 3" is, in Go syntax, x[1][2][3]. Each
//line /usr/local/go/src/text/template/funcs.go:205
// indexed item must be a map, slice, or array.
//line /usr/local/go/src/text/template/funcs.go:208
func index(item reflect.Value, indexes ...reflect.Value) (reflect.Value, error) {
//line /usr/local/go/src/text/template/funcs.go:208
	_go_fuzz_dep_.CoverTab[30268]++
							item = indirectInterface(item)
							if !item.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:210
		_go_fuzz_dep_.CoverTab[30271]++
								return reflect.Value{}, fmt.Errorf("index of untyped nil")
//line /usr/local/go/src/text/template/funcs.go:211
		// _ = "end of CoverTab[30271]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:212
		_go_fuzz_dep_.CoverTab[30272]++
//line /usr/local/go/src/text/template/funcs.go:212
		// _ = "end of CoverTab[30272]"
//line /usr/local/go/src/text/template/funcs.go:212
	}
//line /usr/local/go/src/text/template/funcs.go:212
	// _ = "end of CoverTab[30268]"
//line /usr/local/go/src/text/template/funcs.go:212
	_go_fuzz_dep_.CoverTab[30269]++
							for _, index := range indexes {
//line /usr/local/go/src/text/template/funcs.go:213
		_go_fuzz_dep_.CoverTab[30273]++
								index = indirectInterface(index)
								var isNil bool
								if item, isNil = indirect(item); isNil {
//line /usr/local/go/src/text/template/funcs.go:216
			_go_fuzz_dep_.CoverTab[30275]++
									return reflect.Value{}, fmt.Errorf("index of nil pointer")
//line /usr/local/go/src/text/template/funcs.go:217
			// _ = "end of CoverTab[30275]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:218
			_go_fuzz_dep_.CoverTab[30276]++
//line /usr/local/go/src/text/template/funcs.go:218
			// _ = "end of CoverTab[30276]"
//line /usr/local/go/src/text/template/funcs.go:218
		}
//line /usr/local/go/src/text/template/funcs.go:218
		// _ = "end of CoverTab[30273]"
//line /usr/local/go/src/text/template/funcs.go:218
		_go_fuzz_dep_.CoverTab[30274]++
								switch item.Kind() {
		case reflect.Array, reflect.Slice, reflect.String:
//line /usr/local/go/src/text/template/funcs.go:220
			_go_fuzz_dep_.CoverTab[30277]++
									x, err := indexArg(index, item.Len())
									if err != nil {
//line /usr/local/go/src/text/template/funcs.go:222
				_go_fuzz_dep_.CoverTab[30283]++
										return reflect.Value{}, err
//line /usr/local/go/src/text/template/funcs.go:223
				// _ = "end of CoverTab[30283]"
			} else {
//line /usr/local/go/src/text/template/funcs.go:224
				_go_fuzz_dep_.CoverTab[30284]++
//line /usr/local/go/src/text/template/funcs.go:224
				// _ = "end of CoverTab[30284]"
//line /usr/local/go/src/text/template/funcs.go:224
			}
//line /usr/local/go/src/text/template/funcs.go:224
			// _ = "end of CoverTab[30277]"
//line /usr/local/go/src/text/template/funcs.go:224
			_go_fuzz_dep_.CoverTab[30278]++
									item = item.Index(x)
//line /usr/local/go/src/text/template/funcs.go:225
			// _ = "end of CoverTab[30278]"
		case reflect.Map:
//line /usr/local/go/src/text/template/funcs.go:226
			_go_fuzz_dep_.CoverTab[30279]++
									index, err := prepareArg(index, item.Type().Key())
									if err != nil {
//line /usr/local/go/src/text/template/funcs.go:228
				_go_fuzz_dep_.CoverTab[30285]++
										return reflect.Value{}, err
//line /usr/local/go/src/text/template/funcs.go:229
				// _ = "end of CoverTab[30285]"
			} else {
//line /usr/local/go/src/text/template/funcs.go:230
				_go_fuzz_dep_.CoverTab[30286]++
//line /usr/local/go/src/text/template/funcs.go:230
				// _ = "end of CoverTab[30286]"
//line /usr/local/go/src/text/template/funcs.go:230
			}
//line /usr/local/go/src/text/template/funcs.go:230
			// _ = "end of CoverTab[30279]"
//line /usr/local/go/src/text/template/funcs.go:230
			_go_fuzz_dep_.CoverTab[30280]++
									if x := item.MapIndex(index); x.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:231
				_go_fuzz_dep_.CoverTab[30287]++
										item = x
//line /usr/local/go/src/text/template/funcs.go:232
				// _ = "end of CoverTab[30287]"
			} else {
//line /usr/local/go/src/text/template/funcs.go:233
				_go_fuzz_dep_.CoverTab[30288]++
										item = reflect.Zero(item.Type().Elem())
//line /usr/local/go/src/text/template/funcs.go:234
				// _ = "end of CoverTab[30288]"
			}
//line /usr/local/go/src/text/template/funcs.go:235
			// _ = "end of CoverTab[30280]"
		case reflect.Invalid:
//line /usr/local/go/src/text/template/funcs.go:236
			_go_fuzz_dep_.CoverTab[30281]++

									panic("unreachable")
//line /usr/local/go/src/text/template/funcs.go:238
			// _ = "end of CoverTab[30281]"
		default:
//line /usr/local/go/src/text/template/funcs.go:239
			_go_fuzz_dep_.CoverTab[30282]++
									return reflect.Value{}, fmt.Errorf("can't index item of type %s", item.Type())
//line /usr/local/go/src/text/template/funcs.go:240
			// _ = "end of CoverTab[30282]"
		}
//line /usr/local/go/src/text/template/funcs.go:241
		// _ = "end of CoverTab[30274]"
	}
//line /usr/local/go/src/text/template/funcs.go:242
	// _ = "end of CoverTab[30269]"
//line /usr/local/go/src/text/template/funcs.go:242
	_go_fuzz_dep_.CoverTab[30270]++
							return item, nil
//line /usr/local/go/src/text/template/funcs.go:243
	// _ = "end of CoverTab[30270]"
}

//line /usr/local/go/src/text/template/funcs.go:248
// slice returns the result of slicing its first argument by the remaining
//line /usr/local/go/src/text/template/funcs.go:248
// arguments. Thus "slice x 1 2" is, in Go syntax, x[1:2], while "slice x"
//line /usr/local/go/src/text/template/funcs.go:248
// is x[:], "slice x 1" is x[1:], and "slice x 1 2 3" is x[1:2:3]. The first
//line /usr/local/go/src/text/template/funcs.go:248
// argument must be a string, slice, or array.
//line /usr/local/go/src/text/template/funcs.go:252
func slice(item reflect.Value, indexes ...reflect.Value) (reflect.Value, error) {
//line /usr/local/go/src/text/template/funcs.go:252
	_go_fuzz_dep_.CoverTab[30289]++
							item = indirectInterface(item)
							if !item.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:254
		_go_fuzz_dep_.CoverTab[30297]++
								return reflect.Value{}, fmt.Errorf("slice of untyped nil")
//line /usr/local/go/src/text/template/funcs.go:255
		// _ = "end of CoverTab[30297]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:256
		_go_fuzz_dep_.CoverTab[30298]++
//line /usr/local/go/src/text/template/funcs.go:256
		// _ = "end of CoverTab[30298]"
//line /usr/local/go/src/text/template/funcs.go:256
	}
//line /usr/local/go/src/text/template/funcs.go:256
	// _ = "end of CoverTab[30289]"
//line /usr/local/go/src/text/template/funcs.go:256
	_go_fuzz_dep_.CoverTab[30290]++
							if len(indexes) > 3 {
//line /usr/local/go/src/text/template/funcs.go:257
		_go_fuzz_dep_.CoverTab[30299]++
								return reflect.Value{}, fmt.Errorf("too many slice indexes: %d", len(indexes))
//line /usr/local/go/src/text/template/funcs.go:258
		// _ = "end of CoverTab[30299]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:259
		_go_fuzz_dep_.CoverTab[30300]++
//line /usr/local/go/src/text/template/funcs.go:259
		// _ = "end of CoverTab[30300]"
//line /usr/local/go/src/text/template/funcs.go:259
	}
//line /usr/local/go/src/text/template/funcs.go:259
	// _ = "end of CoverTab[30290]"
//line /usr/local/go/src/text/template/funcs.go:259
	_go_fuzz_dep_.CoverTab[30291]++
							var cap int
							switch item.Kind() {
	case reflect.String:
//line /usr/local/go/src/text/template/funcs.go:262
		_go_fuzz_dep_.CoverTab[30301]++
								if len(indexes) == 3 {
//line /usr/local/go/src/text/template/funcs.go:263
			_go_fuzz_dep_.CoverTab[30305]++
									return reflect.Value{}, fmt.Errorf("cannot 3-index slice a string")
//line /usr/local/go/src/text/template/funcs.go:264
			// _ = "end of CoverTab[30305]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:265
			_go_fuzz_dep_.CoverTab[30306]++
//line /usr/local/go/src/text/template/funcs.go:265
			// _ = "end of CoverTab[30306]"
//line /usr/local/go/src/text/template/funcs.go:265
		}
//line /usr/local/go/src/text/template/funcs.go:265
		// _ = "end of CoverTab[30301]"
//line /usr/local/go/src/text/template/funcs.go:265
		_go_fuzz_dep_.CoverTab[30302]++
								cap = item.Len()
//line /usr/local/go/src/text/template/funcs.go:266
		// _ = "end of CoverTab[30302]"
	case reflect.Array, reflect.Slice:
//line /usr/local/go/src/text/template/funcs.go:267
		_go_fuzz_dep_.CoverTab[30303]++
								cap = item.Cap()
//line /usr/local/go/src/text/template/funcs.go:268
		// _ = "end of CoverTab[30303]"
	default:
//line /usr/local/go/src/text/template/funcs.go:269
		_go_fuzz_dep_.CoverTab[30304]++
								return reflect.Value{}, fmt.Errorf("can't slice item of type %s", item.Type())
//line /usr/local/go/src/text/template/funcs.go:270
		// _ = "end of CoverTab[30304]"
	}
//line /usr/local/go/src/text/template/funcs.go:271
	// _ = "end of CoverTab[30291]"
//line /usr/local/go/src/text/template/funcs.go:271
	_go_fuzz_dep_.CoverTab[30292]++

							idx := [3]int{0, item.Len()}
							for i, index := range indexes {
//line /usr/local/go/src/text/template/funcs.go:274
		_go_fuzz_dep_.CoverTab[30307]++
								x, err := indexArg(index, cap)
								if err != nil {
//line /usr/local/go/src/text/template/funcs.go:276
			_go_fuzz_dep_.CoverTab[30309]++
									return reflect.Value{}, err
//line /usr/local/go/src/text/template/funcs.go:277
			// _ = "end of CoverTab[30309]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:278
			_go_fuzz_dep_.CoverTab[30310]++
//line /usr/local/go/src/text/template/funcs.go:278
			// _ = "end of CoverTab[30310]"
//line /usr/local/go/src/text/template/funcs.go:278
		}
//line /usr/local/go/src/text/template/funcs.go:278
		// _ = "end of CoverTab[30307]"
//line /usr/local/go/src/text/template/funcs.go:278
		_go_fuzz_dep_.CoverTab[30308]++
								idx[i] = x
//line /usr/local/go/src/text/template/funcs.go:279
		// _ = "end of CoverTab[30308]"
	}
//line /usr/local/go/src/text/template/funcs.go:280
	// _ = "end of CoverTab[30292]"
//line /usr/local/go/src/text/template/funcs.go:280
	_go_fuzz_dep_.CoverTab[30293]++

							if idx[0] > idx[1] {
//line /usr/local/go/src/text/template/funcs.go:282
		_go_fuzz_dep_.CoverTab[30311]++
								return reflect.Value{}, fmt.Errorf("invalid slice index: %d > %d", idx[0], idx[1])
//line /usr/local/go/src/text/template/funcs.go:283
		// _ = "end of CoverTab[30311]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:284
		_go_fuzz_dep_.CoverTab[30312]++
//line /usr/local/go/src/text/template/funcs.go:284
		// _ = "end of CoverTab[30312]"
//line /usr/local/go/src/text/template/funcs.go:284
	}
//line /usr/local/go/src/text/template/funcs.go:284
	// _ = "end of CoverTab[30293]"
//line /usr/local/go/src/text/template/funcs.go:284
	_go_fuzz_dep_.CoverTab[30294]++
							if len(indexes) < 3 {
//line /usr/local/go/src/text/template/funcs.go:285
		_go_fuzz_dep_.CoverTab[30313]++
								return item.Slice(idx[0], idx[1]), nil
//line /usr/local/go/src/text/template/funcs.go:286
		// _ = "end of CoverTab[30313]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:287
		_go_fuzz_dep_.CoverTab[30314]++
//line /usr/local/go/src/text/template/funcs.go:287
		// _ = "end of CoverTab[30314]"
//line /usr/local/go/src/text/template/funcs.go:287
	}
//line /usr/local/go/src/text/template/funcs.go:287
	// _ = "end of CoverTab[30294]"
//line /usr/local/go/src/text/template/funcs.go:287
	_go_fuzz_dep_.CoverTab[30295]++

							if idx[1] > idx[2] {
//line /usr/local/go/src/text/template/funcs.go:289
		_go_fuzz_dep_.CoverTab[30315]++
								return reflect.Value{}, fmt.Errorf("invalid slice index: %d > %d", idx[1], idx[2])
//line /usr/local/go/src/text/template/funcs.go:290
		// _ = "end of CoverTab[30315]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:291
		_go_fuzz_dep_.CoverTab[30316]++
//line /usr/local/go/src/text/template/funcs.go:291
		// _ = "end of CoverTab[30316]"
//line /usr/local/go/src/text/template/funcs.go:291
	}
//line /usr/local/go/src/text/template/funcs.go:291
	// _ = "end of CoverTab[30295]"
//line /usr/local/go/src/text/template/funcs.go:291
	_go_fuzz_dep_.CoverTab[30296]++
							return item.Slice3(idx[0], idx[1], idx[2]), nil
//line /usr/local/go/src/text/template/funcs.go:292
	// _ = "end of CoverTab[30296]"
}

//line /usr/local/go/src/text/template/funcs.go:297
// length returns the length of the item, with an error if it has no defined length.
func length(item reflect.Value) (int, error) {
//line /usr/local/go/src/text/template/funcs.go:298
	_go_fuzz_dep_.CoverTab[30317]++
							item, isNil := indirect(item)
							if isNil {
//line /usr/local/go/src/text/template/funcs.go:300
		_go_fuzz_dep_.CoverTab[30320]++
								return 0, fmt.Errorf("len of nil pointer")
//line /usr/local/go/src/text/template/funcs.go:301
		// _ = "end of CoverTab[30320]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:302
		_go_fuzz_dep_.CoverTab[30321]++
//line /usr/local/go/src/text/template/funcs.go:302
		// _ = "end of CoverTab[30321]"
//line /usr/local/go/src/text/template/funcs.go:302
	}
//line /usr/local/go/src/text/template/funcs.go:302
	// _ = "end of CoverTab[30317]"
//line /usr/local/go/src/text/template/funcs.go:302
	_go_fuzz_dep_.CoverTab[30318]++
							switch item.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
//line /usr/local/go/src/text/template/funcs.go:304
		_go_fuzz_dep_.CoverTab[30322]++
								return item.Len(), nil
//line /usr/local/go/src/text/template/funcs.go:305
		// _ = "end of CoverTab[30322]"
//line /usr/local/go/src/text/template/funcs.go:305
	default:
//line /usr/local/go/src/text/template/funcs.go:305
		_go_fuzz_dep_.CoverTab[30323]++
//line /usr/local/go/src/text/template/funcs.go:305
		// _ = "end of CoverTab[30323]"
	}
//line /usr/local/go/src/text/template/funcs.go:306
	// _ = "end of CoverTab[30318]"
//line /usr/local/go/src/text/template/funcs.go:306
	_go_fuzz_dep_.CoverTab[30319]++
							return 0, fmt.Errorf("len of type %s", item.Type())
//line /usr/local/go/src/text/template/funcs.go:307
	// _ = "end of CoverTab[30319]"
}

//line /usr/local/go/src/text/template/funcs.go:312
// call returns the result of evaluating the first argument as a function.
//line /usr/local/go/src/text/template/funcs.go:312
// The function must return 1 result, or 2 results, the second of which is an error.
//line /usr/local/go/src/text/template/funcs.go:314
func call(fn reflect.Value, args ...reflect.Value) (reflect.Value, error) {
//line /usr/local/go/src/text/template/funcs.go:314
	_go_fuzz_dep_.CoverTab[30324]++
							fn = indirectInterface(fn)
							if !fn.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:316
		_go_fuzz_dep_.CoverTab[30330]++
								return reflect.Value{}, fmt.Errorf("call of nil")
//line /usr/local/go/src/text/template/funcs.go:317
		// _ = "end of CoverTab[30330]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:318
		_go_fuzz_dep_.CoverTab[30331]++
//line /usr/local/go/src/text/template/funcs.go:318
		// _ = "end of CoverTab[30331]"
//line /usr/local/go/src/text/template/funcs.go:318
	}
//line /usr/local/go/src/text/template/funcs.go:318
	// _ = "end of CoverTab[30324]"
//line /usr/local/go/src/text/template/funcs.go:318
	_go_fuzz_dep_.CoverTab[30325]++
							typ := fn.Type()
							if typ.Kind() != reflect.Func {
//line /usr/local/go/src/text/template/funcs.go:320
		_go_fuzz_dep_.CoverTab[30332]++
								return reflect.Value{}, fmt.Errorf("non-function of type %s", typ)
//line /usr/local/go/src/text/template/funcs.go:321
		// _ = "end of CoverTab[30332]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:322
		_go_fuzz_dep_.CoverTab[30333]++
//line /usr/local/go/src/text/template/funcs.go:322
		// _ = "end of CoverTab[30333]"
//line /usr/local/go/src/text/template/funcs.go:322
	}
//line /usr/local/go/src/text/template/funcs.go:322
	// _ = "end of CoverTab[30325]"
//line /usr/local/go/src/text/template/funcs.go:322
	_go_fuzz_dep_.CoverTab[30326]++
							if !goodFunc(typ) {
//line /usr/local/go/src/text/template/funcs.go:323
		_go_fuzz_dep_.CoverTab[30334]++
								return reflect.Value{}, fmt.Errorf("function called with %d args; should be 1 or 2", typ.NumOut())
//line /usr/local/go/src/text/template/funcs.go:324
		// _ = "end of CoverTab[30334]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:325
		_go_fuzz_dep_.CoverTab[30335]++
//line /usr/local/go/src/text/template/funcs.go:325
		// _ = "end of CoverTab[30335]"
//line /usr/local/go/src/text/template/funcs.go:325
	}
//line /usr/local/go/src/text/template/funcs.go:325
	// _ = "end of CoverTab[30326]"
//line /usr/local/go/src/text/template/funcs.go:325
	_go_fuzz_dep_.CoverTab[30327]++
							numIn := typ.NumIn()
							var dddType reflect.Type
							if typ.IsVariadic() {
//line /usr/local/go/src/text/template/funcs.go:328
		_go_fuzz_dep_.CoverTab[30336]++
								if len(args) < numIn-1 {
//line /usr/local/go/src/text/template/funcs.go:329
			_go_fuzz_dep_.CoverTab[30338]++
									return reflect.Value{}, fmt.Errorf("wrong number of args: got %d want at least %d", len(args), numIn-1)
//line /usr/local/go/src/text/template/funcs.go:330
			// _ = "end of CoverTab[30338]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:331
			_go_fuzz_dep_.CoverTab[30339]++
//line /usr/local/go/src/text/template/funcs.go:331
			// _ = "end of CoverTab[30339]"
//line /usr/local/go/src/text/template/funcs.go:331
		}
//line /usr/local/go/src/text/template/funcs.go:331
		// _ = "end of CoverTab[30336]"
//line /usr/local/go/src/text/template/funcs.go:331
		_go_fuzz_dep_.CoverTab[30337]++
								dddType = typ.In(numIn - 1).Elem()
//line /usr/local/go/src/text/template/funcs.go:332
		// _ = "end of CoverTab[30337]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:333
		_go_fuzz_dep_.CoverTab[30340]++
								if len(args) != numIn {
//line /usr/local/go/src/text/template/funcs.go:334
			_go_fuzz_dep_.CoverTab[30341]++
									return reflect.Value{}, fmt.Errorf("wrong number of args: got %d want %d", len(args), numIn)
//line /usr/local/go/src/text/template/funcs.go:335
			// _ = "end of CoverTab[30341]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:336
			_go_fuzz_dep_.CoverTab[30342]++
//line /usr/local/go/src/text/template/funcs.go:336
			// _ = "end of CoverTab[30342]"
//line /usr/local/go/src/text/template/funcs.go:336
		}
//line /usr/local/go/src/text/template/funcs.go:336
		// _ = "end of CoverTab[30340]"
	}
//line /usr/local/go/src/text/template/funcs.go:337
	// _ = "end of CoverTab[30327]"
//line /usr/local/go/src/text/template/funcs.go:337
	_go_fuzz_dep_.CoverTab[30328]++
							argv := make([]reflect.Value, len(args))
							for i, arg := range args {
//line /usr/local/go/src/text/template/funcs.go:339
		_go_fuzz_dep_.CoverTab[30343]++
								arg = indirectInterface(arg)

								argType := dddType
								if !typ.IsVariadic() || func() bool {
//line /usr/local/go/src/text/template/funcs.go:343
			_go_fuzz_dep_.CoverTab[30345]++
//line /usr/local/go/src/text/template/funcs.go:343
			return i < numIn-1
//line /usr/local/go/src/text/template/funcs.go:343
			// _ = "end of CoverTab[30345]"
//line /usr/local/go/src/text/template/funcs.go:343
		}() {
//line /usr/local/go/src/text/template/funcs.go:343
			_go_fuzz_dep_.CoverTab[30346]++
									argType = typ.In(i)
//line /usr/local/go/src/text/template/funcs.go:344
			// _ = "end of CoverTab[30346]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:345
			_go_fuzz_dep_.CoverTab[30347]++
//line /usr/local/go/src/text/template/funcs.go:345
			// _ = "end of CoverTab[30347]"
//line /usr/local/go/src/text/template/funcs.go:345
		}
//line /usr/local/go/src/text/template/funcs.go:345
		// _ = "end of CoverTab[30343]"
//line /usr/local/go/src/text/template/funcs.go:345
		_go_fuzz_dep_.CoverTab[30344]++

								var err error
								if argv[i], err = prepareArg(arg, argType); err != nil {
//line /usr/local/go/src/text/template/funcs.go:348
			_go_fuzz_dep_.CoverTab[30348]++
									return reflect.Value{}, fmt.Errorf("arg %d: %w", i, err)
//line /usr/local/go/src/text/template/funcs.go:349
			// _ = "end of CoverTab[30348]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:350
			_go_fuzz_dep_.CoverTab[30349]++
//line /usr/local/go/src/text/template/funcs.go:350
			// _ = "end of CoverTab[30349]"
//line /usr/local/go/src/text/template/funcs.go:350
		}
//line /usr/local/go/src/text/template/funcs.go:350
		// _ = "end of CoverTab[30344]"
	}
//line /usr/local/go/src/text/template/funcs.go:351
	// _ = "end of CoverTab[30328]"
//line /usr/local/go/src/text/template/funcs.go:351
	_go_fuzz_dep_.CoverTab[30329]++
							return safeCall(fn, argv)
//line /usr/local/go/src/text/template/funcs.go:352
	// _ = "end of CoverTab[30329]"
}

// safeCall runs fun.Call(args), and returns the resulting value and error, if
//line /usr/local/go/src/text/template/funcs.go:355
// any. If the call panics, the panic value is returned as an error.
//line /usr/local/go/src/text/template/funcs.go:357
func safeCall(fun reflect.Value, args []reflect.Value) (val reflect.Value, err error) {
//line /usr/local/go/src/text/template/funcs.go:357
	_go_fuzz_dep_.CoverTab[30350]++
							defer func() {
//line /usr/local/go/src/text/template/funcs.go:358
		_go_fuzz_dep_.CoverTab[30353]++
								if r := recover(); r != nil {
//line /usr/local/go/src/text/template/funcs.go:359
			_go_fuzz_dep_.CoverTab[30354]++
									if e, ok := r.(error); ok {
//line /usr/local/go/src/text/template/funcs.go:360
				_go_fuzz_dep_.CoverTab[30355]++
										err = e
//line /usr/local/go/src/text/template/funcs.go:361
				// _ = "end of CoverTab[30355]"
			} else {
//line /usr/local/go/src/text/template/funcs.go:362
				_go_fuzz_dep_.CoverTab[30356]++
										err = fmt.Errorf("%v", r)
//line /usr/local/go/src/text/template/funcs.go:363
				// _ = "end of CoverTab[30356]"
			}
//line /usr/local/go/src/text/template/funcs.go:364
			// _ = "end of CoverTab[30354]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:365
			_go_fuzz_dep_.CoverTab[30357]++
//line /usr/local/go/src/text/template/funcs.go:365
			// _ = "end of CoverTab[30357]"
//line /usr/local/go/src/text/template/funcs.go:365
		}
//line /usr/local/go/src/text/template/funcs.go:365
		// _ = "end of CoverTab[30353]"
	}()
//line /usr/local/go/src/text/template/funcs.go:366
	// _ = "end of CoverTab[30350]"
//line /usr/local/go/src/text/template/funcs.go:366
	_go_fuzz_dep_.CoverTab[30351]++
							ret := fun.Call(args)
							if len(ret) == 2 && func() bool {
//line /usr/local/go/src/text/template/funcs.go:368
		_go_fuzz_dep_.CoverTab[30358]++
//line /usr/local/go/src/text/template/funcs.go:368
		return !ret[1].IsNil()
//line /usr/local/go/src/text/template/funcs.go:368
		// _ = "end of CoverTab[30358]"
//line /usr/local/go/src/text/template/funcs.go:368
	}() {
//line /usr/local/go/src/text/template/funcs.go:368
		_go_fuzz_dep_.CoverTab[30359]++
								return ret[0], ret[1].Interface().(error)
//line /usr/local/go/src/text/template/funcs.go:369
		// _ = "end of CoverTab[30359]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:370
		_go_fuzz_dep_.CoverTab[30360]++
//line /usr/local/go/src/text/template/funcs.go:370
		// _ = "end of CoverTab[30360]"
//line /usr/local/go/src/text/template/funcs.go:370
	}
//line /usr/local/go/src/text/template/funcs.go:370
	// _ = "end of CoverTab[30351]"
//line /usr/local/go/src/text/template/funcs.go:370
	_go_fuzz_dep_.CoverTab[30352]++
							return ret[0], nil
//line /usr/local/go/src/text/template/funcs.go:371
	// _ = "end of CoverTab[30352]"
}

//line /usr/local/go/src/text/template/funcs.go:376
func truth(arg reflect.Value) bool {
//line /usr/local/go/src/text/template/funcs.go:376
	_go_fuzz_dep_.CoverTab[30361]++
							t, _ := isTrue(indirectInterface(arg))
							return t
//line /usr/local/go/src/text/template/funcs.go:378
	// _ = "end of CoverTab[30361]"
}

// and computes the Boolean AND of its arguments, returning
//line /usr/local/go/src/text/template/funcs.go:381
// the first false argument it encounters, or the last argument.
//line /usr/local/go/src/text/template/funcs.go:383
func and(arg0 reflect.Value, args ...reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/funcs.go:383
	_go_fuzz_dep_.CoverTab[30362]++
							panic("unreachable")
//line /usr/local/go/src/text/template/funcs.go:384
	// _ = "end of CoverTab[30362]"
}

// or computes the Boolean OR of its arguments, returning
//line /usr/local/go/src/text/template/funcs.go:387
// the first true argument it encounters, or the last argument.
//line /usr/local/go/src/text/template/funcs.go:389
func or(arg0 reflect.Value, args ...reflect.Value) reflect.Value {
//line /usr/local/go/src/text/template/funcs.go:389
	_go_fuzz_dep_.CoverTab[30363]++
							panic("unreachable")
//line /usr/local/go/src/text/template/funcs.go:390
	// _ = "end of CoverTab[30363]"
}

// not returns the Boolean negation of its argument.
func not(arg reflect.Value) bool {
//line /usr/local/go/src/text/template/funcs.go:394
	_go_fuzz_dep_.CoverTab[30364]++
							return !truth(arg)
//line /usr/local/go/src/text/template/funcs.go:395
	// _ = "end of CoverTab[30364]"
}

//line /usr/local/go/src/text/template/funcs.go:402
var (
	errBadComparisonType	= errors.New("invalid type for comparison")
	errBadComparison	= errors.New("incompatible types for comparison")
	errNoComparison		= errors.New("missing argument for comparison")
)

type kind int

const (
	invalidKind	kind	= iota
	boolKind
	complexKind
	intKind
	floatKind
	stringKind
	uintKind
)

func basicKind(v reflect.Value) (kind, error) {
//line /usr/local/go/src/text/template/funcs.go:420
	_go_fuzz_dep_.CoverTab[30365]++
							switch v.Kind() {
	case reflect.Bool:
//line /usr/local/go/src/text/template/funcs.go:422
		_go_fuzz_dep_.CoverTab[30367]++
								return boolKind, nil
//line /usr/local/go/src/text/template/funcs.go:423
		// _ = "end of CoverTab[30367]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/text/template/funcs.go:424
		_go_fuzz_dep_.CoverTab[30368]++
								return intKind, nil
//line /usr/local/go/src/text/template/funcs.go:425
		// _ = "end of CoverTab[30368]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/text/template/funcs.go:426
		_go_fuzz_dep_.CoverTab[30369]++
								return uintKind, nil
//line /usr/local/go/src/text/template/funcs.go:427
		// _ = "end of CoverTab[30369]"
	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/text/template/funcs.go:428
		_go_fuzz_dep_.CoverTab[30370]++
								return floatKind, nil
//line /usr/local/go/src/text/template/funcs.go:429
		// _ = "end of CoverTab[30370]"
	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/text/template/funcs.go:430
		_go_fuzz_dep_.CoverTab[30371]++
								return complexKind, nil
//line /usr/local/go/src/text/template/funcs.go:431
		// _ = "end of CoverTab[30371]"
	case reflect.String:
//line /usr/local/go/src/text/template/funcs.go:432
		_go_fuzz_dep_.CoverTab[30372]++
								return stringKind, nil
//line /usr/local/go/src/text/template/funcs.go:433
		// _ = "end of CoverTab[30372]"
//line /usr/local/go/src/text/template/funcs.go:433
	default:
//line /usr/local/go/src/text/template/funcs.go:433
		_go_fuzz_dep_.CoverTab[30373]++
//line /usr/local/go/src/text/template/funcs.go:433
		// _ = "end of CoverTab[30373]"
	}
//line /usr/local/go/src/text/template/funcs.go:434
	// _ = "end of CoverTab[30365]"
//line /usr/local/go/src/text/template/funcs.go:434
	_go_fuzz_dep_.CoverTab[30366]++
							return invalidKind, errBadComparisonType
//line /usr/local/go/src/text/template/funcs.go:435
	// _ = "end of CoverTab[30366]"
}

// isNil returns true if v is the zero reflect.Value, or nil of its type.
func isNil(v reflect.Value) bool {
//line /usr/local/go/src/text/template/funcs.go:439
	_go_fuzz_dep_.CoverTab[30374]++
							if !v.IsValid() {
//line /usr/local/go/src/text/template/funcs.go:440
		_go_fuzz_dep_.CoverTab[30377]++
								return true
//line /usr/local/go/src/text/template/funcs.go:441
		// _ = "end of CoverTab[30377]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:442
		_go_fuzz_dep_.CoverTab[30378]++
//line /usr/local/go/src/text/template/funcs.go:442
		// _ = "end of CoverTab[30378]"
//line /usr/local/go/src/text/template/funcs.go:442
	}
//line /usr/local/go/src/text/template/funcs.go:442
	// _ = "end of CoverTab[30374]"
//line /usr/local/go/src/text/template/funcs.go:442
	_go_fuzz_dep_.CoverTab[30375]++
							switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
//line /usr/local/go/src/text/template/funcs.go:444
		_go_fuzz_dep_.CoverTab[30379]++
								return v.IsNil()
//line /usr/local/go/src/text/template/funcs.go:445
		// _ = "end of CoverTab[30379]"
//line /usr/local/go/src/text/template/funcs.go:445
	default:
//line /usr/local/go/src/text/template/funcs.go:445
		_go_fuzz_dep_.CoverTab[30380]++
//line /usr/local/go/src/text/template/funcs.go:445
		// _ = "end of CoverTab[30380]"
	}
//line /usr/local/go/src/text/template/funcs.go:446
	// _ = "end of CoverTab[30375]"
//line /usr/local/go/src/text/template/funcs.go:446
	_go_fuzz_dep_.CoverTab[30376]++
							return false
//line /usr/local/go/src/text/template/funcs.go:447
	// _ = "end of CoverTab[30376]"
}

// canCompare reports whether v1 and v2 are both the same kind, or one is nil.
//line /usr/local/go/src/text/template/funcs.go:450
// Called only when dealing with nillable types, or there's about to be an error.
//line /usr/local/go/src/text/template/funcs.go:452
func canCompare(v1, v2 reflect.Value) bool {
//line /usr/local/go/src/text/template/funcs.go:452
	_go_fuzz_dep_.CoverTab[30381]++
							k1 := v1.Kind()
							k2 := v2.Kind()
							if k1 == k2 {
//line /usr/local/go/src/text/template/funcs.go:455
		_go_fuzz_dep_.CoverTab[30383]++
								return true
//line /usr/local/go/src/text/template/funcs.go:456
		// _ = "end of CoverTab[30383]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:457
		_go_fuzz_dep_.CoverTab[30384]++
//line /usr/local/go/src/text/template/funcs.go:457
		// _ = "end of CoverTab[30384]"
//line /usr/local/go/src/text/template/funcs.go:457
	}
//line /usr/local/go/src/text/template/funcs.go:457
	// _ = "end of CoverTab[30381]"
//line /usr/local/go/src/text/template/funcs.go:457
	_go_fuzz_dep_.CoverTab[30382]++

							return k1 == reflect.Invalid || func() bool {
//line /usr/local/go/src/text/template/funcs.go:459
		_go_fuzz_dep_.CoverTab[30385]++
//line /usr/local/go/src/text/template/funcs.go:459
		return k2 == reflect.Invalid
//line /usr/local/go/src/text/template/funcs.go:459
		// _ = "end of CoverTab[30385]"
//line /usr/local/go/src/text/template/funcs.go:459
	}()
//line /usr/local/go/src/text/template/funcs.go:459
	// _ = "end of CoverTab[30382]"
}

// eq evaluates the comparison a == b || a == c || ...
func eq(arg1 reflect.Value, arg2 ...reflect.Value) (bool, error) {
//line /usr/local/go/src/text/template/funcs.go:463
	_go_fuzz_dep_.CoverTab[30386]++
							arg1 = indirectInterface(arg1)
							if len(arg2) == 0 {
//line /usr/local/go/src/text/template/funcs.go:465
		_go_fuzz_dep_.CoverTab[30389]++
								return false, errNoComparison
//line /usr/local/go/src/text/template/funcs.go:466
		// _ = "end of CoverTab[30389]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:467
		_go_fuzz_dep_.CoverTab[30390]++
//line /usr/local/go/src/text/template/funcs.go:467
		// _ = "end of CoverTab[30390]"
//line /usr/local/go/src/text/template/funcs.go:467
	}
//line /usr/local/go/src/text/template/funcs.go:467
	// _ = "end of CoverTab[30386]"
//line /usr/local/go/src/text/template/funcs.go:467
	_go_fuzz_dep_.CoverTab[30387]++
							k1, _ := basicKind(arg1)
							for _, arg := range arg2 {
//line /usr/local/go/src/text/template/funcs.go:469
		_go_fuzz_dep_.CoverTab[30391]++
								arg = indirectInterface(arg)
								k2, _ := basicKind(arg)
								truth := false
								if k1 != k2 {
//line /usr/local/go/src/text/template/funcs.go:473
			_go_fuzz_dep_.CoverTab[30393]++

									switch {
			case k1 == intKind && func() bool {
//line /usr/local/go/src/text/template/funcs.go:476
				_go_fuzz_dep_.CoverTab[30397]++
//line /usr/local/go/src/text/template/funcs.go:476
				return k2 == uintKind
//line /usr/local/go/src/text/template/funcs.go:476
				// _ = "end of CoverTab[30397]"
//line /usr/local/go/src/text/template/funcs.go:476
			}():
//line /usr/local/go/src/text/template/funcs.go:476
				_go_fuzz_dep_.CoverTab[30394]++
										truth = arg1.Int() >= 0 && func() bool {
//line /usr/local/go/src/text/template/funcs.go:477
					_go_fuzz_dep_.CoverTab[30398]++
//line /usr/local/go/src/text/template/funcs.go:477
					return uint64(arg1.Int()) == arg.Uint()
//line /usr/local/go/src/text/template/funcs.go:477
					// _ = "end of CoverTab[30398]"
//line /usr/local/go/src/text/template/funcs.go:477
				}()
//line /usr/local/go/src/text/template/funcs.go:477
				// _ = "end of CoverTab[30394]"
			case k1 == uintKind && func() bool {
//line /usr/local/go/src/text/template/funcs.go:478
				_go_fuzz_dep_.CoverTab[30399]++
//line /usr/local/go/src/text/template/funcs.go:478
				return k2 == intKind
//line /usr/local/go/src/text/template/funcs.go:478
				// _ = "end of CoverTab[30399]"
//line /usr/local/go/src/text/template/funcs.go:478
			}():
//line /usr/local/go/src/text/template/funcs.go:478
				_go_fuzz_dep_.CoverTab[30395]++
										truth = arg.Int() >= 0 && func() bool {
//line /usr/local/go/src/text/template/funcs.go:479
					_go_fuzz_dep_.CoverTab[30400]++
//line /usr/local/go/src/text/template/funcs.go:479
					return arg1.Uint() == uint64(arg.Int())
//line /usr/local/go/src/text/template/funcs.go:479
					// _ = "end of CoverTab[30400]"
//line /usr/local/go/src/text/template/funcs.go:479
				}()
//line /usr/local/go/src/text/template/funcs.go:479
				// _ = "end of CoverTab[30395]"
			default:
//line /usr/local/go/src/text/template/funcs.go:480
				_go_fuzz_dep_.CoverTab[30396]++
										if arg1 != zero && func() bool {
//line /usr/local/go/src/text/template/funcs.go:481
					_go_fuzz_dep_.CoverTab[30401]++
//line /usr/local/go/src/text/template/funcs.go:481
					return arg != zero
//line /usr/local/go/src/text/template/funcs.go:481
					// _ = "end of CoverTab[30401]"
//line /usr/local/go/src/text/template/funcs.go:481
				}() {
//line /usr/local/go/src/text/template/funcs.go:481
					_go_fuzz_dep_.CoverTab[30402]++
											return false, errBadComparison
//line /usr/local/go/src/text/template/funcs.go:482
					// _ = "end of CoverTab[30402]"
				} else {
//line /usr/local/go/src/text/template/funcs.go:483
					_go_fuzz_dep_.CoverTab[30403]++
//line /usr/local/go/src/text/template/funcs.go:483
					// _ = "end of CoverTab[30403]"
//line /usr/local/go/src/text/template/funcs.go:483
				}
//line /usr/local/go/src/text/template/funcs.go:483
				// _ = "end of CoverTab[30396]"
			}
//line /usr/local/go/src/text/template/funcs.go:484
			// _ = "end of CoverTab[30393]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:485
			_go_fuzz_dep_.CoverTab[30404]++
									switch k1 {
			case boolKind:
//line /usr/local/go/src/text/template/funcs.go:487
				_go_fuzz_dep_.CoverTab[30405]++
										truth = arg1.Bool() == arg.Bool()
//line /usr/local/go/src/text/template/funcs.go:488
				// _ = "end of CoverTab[30405]"
			case complexKind:
//line /usr/local/go/src/text/template/funcs.go:489
				_go_fuzz_dep_.CoverTab[30406]++
										truth = arg1.Complex() == arg.Complex()
//line /usr/local/go/src/text/template/funcs.go:490
				// _ = "end of CoverTab[30406]"
			case floatKind:
//line /usr/local/go/src/text/template/funcs.go:491
				_go_fuzz_dep_.CoverTab[30407]++
										truth = arg1.Float() == arg.Float()
//line /usr/local/go/src/text/template/funcs.go:492
				// _ = "end of CoverTab[30407]"
			case intKind:
//line /usr/local/go/src/text/template/funcs.go:493
				_go_fuzz_dep_.CoverTab[30408]++
										truth = arg1.Int() == arg.Int()
//line /usr/local/go/src/text/template/funcs.go:494
				// _ = "end of CoverTab[30408]"
			case stringKind:
//line /usr/local/go/src/text/template/funcs.go:495
				_go_fuzz_dep_.CoverTab[30409]++
										truth = arg1.String() == arg.String()
//line /usr/local/go/src/text/template/funcs.go:496
				// _ = "end of CoverTab[30409]"
			case uintKind:
//line /usr/local/go/src/text/template/funcs.go:497
				_go_fuzz_dep_.CoverTab[30410]++
										truth = arg1.Uint() == arg.Uint()
//line /usr/local/go/src/text/template/funcs.go:498
				// _ = "end of CoverTab[30410]"
			default:
//line /usr/local/go/src/text/template/funcs.go:499
				_go_fuzz_dep_.CoverTab[30411]++
										if !canCompare(arg1, arg) {
//line /usr/local/go/src/text/template/funcs.go:500
					_go_fuzz_dep_.CoverTab[30413]++
											return false, fmt.Errorf("non-comparable types %s: %v, %s: %v", arg1, arg1.Type(), arg.Type(), arg)
//line /usr/local/go/src/text/template/funcs.go:501
					// _ = "end of CoverTab[30413]"
				} else {
//line /usr/local/go/src/text/template/funcs.go:502
					_go_fuzz_dep_.CoverTab[30414]++
//line /usr/local/go/src/text/template/funcs.go:502
					// _ = "end of CoverTab[30414]"
//line /usr/local/go/src/text/template/funcs.go:502
				}
//line /usr/local/go/src/text/template/funcs.go:502
				// _ = "end of CoverTab[30411]"
//line /usr/local/go/src/text/template/funcs.go:502
				_go_fuzz_dep_.CoverTab[30412]++
										if isNil(arg1) || func() bool {
//line /usr/local/go/src/text/template/funcs.go:503
					_go_fuzz_dep_.CoverTab[30415]++
//line /usr/local/go/src/text/template/funcs.go:503
					return isNil(arg)
//line /usr/local/go/src/text/template/funcs.go:503
					// _ = "end of CoverTab[30415]"
//line /usr/local/go/src/text/template/funcs.go:503
				}() {
//line /usr/local/go/src/text/template/funcs.go:503
					_go_fuzz_dep_.CoverTab[30416]++
											truth = isNil(arg) == isNil(arg1)
//line /usr/local/go/src/text/template/funcs.go:504
					// _ = "end of CoverTab[30416]"
				} else {
//line /usr/local/go/src/text/template/funcs.go:505
					_go_fuzz_dep_.CoverTab[30417]++
											if !arg.Type().Comparable() {
//line /usr/local/go/src/text/template/funcs.go:506
						_go_fuzz_dep_.CoverTab[30419]++
												return false, fmt.Errorf("non-comparable type %s: %v", arg, arg.Type())
//line /usr/local/go/src/text/template/funcs.go:507
						// _ = "end of CoverTab[30419]"
					} else {
//line /usr/local/go/src/text/template/funcs.go:508
						_go_fuzz_dep_.CoverTab[30420]++
//line /usr/local/go/src/text/template/funcs.go:508
						// _ = "end of CoverTab[30420]"
//line /usr/local/go/src/text/template/funcs.go:508
					}
//line /usr/local/go/src/text/template/funcs.go:508
					// _ = "end of CoverTab[30417]"
//line /usr/local/go/src/text/template/funcs.go:508
					_go_fuzz_dep_.CoverTab[30418]++
											truth = arg1.Interface() == arg.Interface()
//line /usr/local/go/src/text/template/funcs.go:509
					// _ = "end of CoverTab[30418]"
				}
//line /usr/local/go/src/text/template/funcs.go:510
				// _ = "end of CoverTab[30412]"
			}
//line /usr/local/go/src/text/template/funcs.go:511
			// _ = "end of CoverTab[30404]"
		}
//line /usr/local/go/src/text/template/funcs.go:512
		// _ = "end of CoverTab[30391]"
//line /usr/local/go/src/text/template/funcs.go:512
		_go_fuzz_dep_.CoverTab[30392]++
								if truth {
//line /usr/local/go/src/text/template/funcs.go:513
			_go_fuzz_dep_.CoverTab[30421]++
									return true, nil
//line /usr/local/go/src/text/template/funcs.go:514
			// _ = "end of CoverTab[30421]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:515
			_go_fuzz_dep_.CoverTab[30422]++
//line /usr/local/go/src/text/template/funcs.go:515
			// _ = "end of CoverTab[30422]"
//line /usr/local/go/src/text/template/funcs.go:515
		}
//line /usr/local/go/src/text/template/funcs.go:515
		// _ = "end of CoverTab[30392]"
	}
//line /usr/local/go/src/text/template/funcs.go:516
	// _ = "end of CoverTab[30387]"
//line /usr/local/go/src/text/template/funcs.go:516
	_go_fuzz_dep_.CoverTab[30388]++
							return false, nil
//line /usr/local/go/src/text/template/funcs.go:517
	// _ = "end of CoverTab[30388]"
}

// ne evaluates the comparison a != b.
func ne(arg1, arg2 reflect.Value) (bool, error) {
//line /usr/local/go/src/text/template/funcs.go:521
	_go_fuzz_dep_.CoverTab[30423]++

							equal, err := eq(arg1, arg2)
							return !equal, err
//line /usr/local/go/src/text/template/funcs.go:524
	// _ = "end of CoverTab[30423]"
}

// lt evaluates the comparison a < b.
func lt(arg1, arg2 reflect.Value) (bool, error) {
//line /usr/local/go/src/text/template/funcs.go:528
	_go_fuzz_dep_.CoverTab[30424]++
							arg1 = indirectInterface(arg1)
							k1, err := basicKind(arg1)
							if err != nil {
//line /usr/local/go/src/text/template/funcs.go:531
		_go_fuzz_dep_.CoverTab[30428]++
								return false, err
//line /usr/local/go/src/text/template/funcs.go:532
		// _ = "end of CoverTab[30428]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:533
		_go_fuzz_dep_.CoverTab[30429]++
//line /usr/local/go/src/text/template/funcs.go:533
		// _ = "end of CoverTab[30429]"
//line /usr/local/go/src/text/template/funcs.go:533
	}
//line /usr/local/go/src/text/template/funcs.go:533
	// _ = "end of CoverTab[30424]"
//line /usr/local/go/src/text/template/funcs.go:533
	_go_fuzz_dep_.CoverTab[30425]++
							arg2 = indirectInterface(arg2)
							k2, err := basicKind(arg2)
							if err != nil {
//line /usr/local/go/src/text/template/funcs.go:536
		_go_fuzz_dep_.CoverTab[30430]++
								return false, err
//line /usr/local/go/src/text/template/funcs.go:537
		// _ = "end of CoverTab[30430]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:538
		_go_fuzz_dep_.CoverTab[30431]++
//line /usr/local/go/src/text/template/funcs.go:538
		// _ = "end of CoverTab[30431]"
//line /usr/local/go/src/text/template/funcs.go:538
	}
//line /usr/local/go/src/text/template/funcs.go:538
	// _ = "end of CoverTab[30425]"
//line /usr/local/go/src/text/template/funcs.go:538
	_go_fuzz_dep_.CoverTab[30426]++
							truth := false
							if k1 != k2 {
//line /usr/local/go/src/text/template/funcs.go:540
		_go_fuzz_dep_.CoverTab[30432]++

								switch {
		case k1 == intKind && func() bool {
//line /usr/local/go/src/text/template/funcs.go:543
			_go_fuzz_dep_.CoverTab[30436]++
//line /usr/local/go/src/text/template/funcs.go:543
			return k2 == uintKind
//line /usr/local/go/src/text/template/funcs.go:543
			// _ = "end of CoverTab[30436]"
//line /usr/local/go/src/text/template/funcs.go:543
		}():
//line /usr/local/go/src/text/template/funcs.go:543
			_go_fuzz_dep_.CoverTab[30433]++
									truth = arg1.Int() < 0 || func() bool {
//line /usr/local/go/src/text/template/funcs.go:544
				_go_fuzz_dep_.CoverTab[30437]++
//line /usr/local/go/src/text/template/funcs.go:544
				return uint64(arg1.Int()) < arg2.Uint()
//line /usr/local/go/src/text/template/funcs.go:544
				// _ = "end of CoverTab[30437]"
//line /usr/local/go/src/text/template/funcs.go:544
			}()
//line /usr/local/go/src/text/template/funcs.go:544
			// _ = "end of CoverTab[30433]"
		case k1 == uintKind && func() bool {
//line /usr/local/go/src/text/template/funcs.go:545
			_go_fuzz_dep_.CoverTab[30438]++
//line /usr/local/go/src/text/template/funcs.go:545
			return k2 == intKind
//line /usr/local/go/src/text/template/funcs.go:545
			// _ = "end of CoverTab[30438]"
//line /usr/local/go/src/text/template/funcs.go:545
		}():
//line /usr/local/go/src/text/template/funcs.go:545
			_go_fuzz_dep_.CoverTab[30434]++
									truth = arg2.Int() >= 0 && func() bool {
//line /usr/local/go/src/text/template/funcs.go:546
				_go_fuzz_dep_.CoverTab[30439]++
//line /usr/local/go/src/text/template/funcs.go:546
				return arg1.Uint() < uint64(arg2.Int())
//line /usr/local/go/src/text/template/funcs.go:546
				// _ = "end of CoverTab[30439]"
//line /usr/local/go/src/text/template/funcs.go:546
			}()
//line /usr/local/go/src/text/template/funcs.go:546
			// _ = "end of CoverTab[30434]"
		default:
//line /usr/local/go/src/text/template/funcs.go:547
			_go_fuzz_dep_.CoverTab[30435]++
									return false, errBadComparison
//line /usr/local/go/src/text/template/funcs.go:548
			// _ = "end of CoverTab[30435]"
		}
//line /usr/local/go/src/text/template/funcs.go:549
		// _ = "end of CoverTab[30432]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:550
		_go_fuzz_dep_.CoverTab[30440]++
								switch k1 {
		case boolKind, complexKind:
//line /usr/local/go/src/text/template/funcs.go:552
			_go_fuzz_dep_.CoverTab[30441]++
									return false, errBadComparisonType
//line /usr/local/go/src/text/template/funcs.go:553
			// _ = "end of CoverTab[30441]"
		case floatKind:
//line /usr/local/go/src/text/template/funcs.go:554
			_go_fuzz_dep_.CoverTab[30442]++
									truth = arg1.Float() < arg2.Float()
//line /usr/local/go/src/text/template/funcs.go:555
			// _ = "end of CoverTab[30442]"
		case intKind:
//line /usr/local/go/src/text/template/funcs.go:556
			_go_fuzz_dep_.CoverTab[30443]++
									truth = arg1.Int() < arg2.Int()
//line /usr/local/go/src/text/template/funcs.go:557
			// _ = "end of CoverTab[30443]"
		case stringKind:
//line /usr/local/go/src/text/template/funcs.go:558
			_go_fuzz_dep_.CoverTab[30444]++
									truth = arg1.String() < arg2.String()
//line /usr/local/go/src/text/template/funcs.go:559
			// _ = "end of CoverTab[30444]"
		case uintKind:
//line /usr/local/go/src/text/template/funcs.go:560
			_go_fuzz_dep_.CoverTab[30445]++
									truth = arg1.Uint() < arg2.Uint()
//line /usr/local/go/src/text/template/funcs.go:561
			// _ = "end of CoverTab[30445]"
		default:
//line /usr/local/go/src/text/template/funcs.go:562
			_go_fuzz_dep_.CoverTab[30446]++
									panic("invalid kind")
//line /usr/local/go/src/text/template/funcs.go:563
			// _ = "end of CoverTab[30446]"
		}
//line /usr/local/go/src/text/template/funcs.go:564
		// _ = "end of CoverTab[30440]"
	}
//line /usr/local/go/src/text/template/funcs.go:565
	// _ = "end of CoverTab[30426]"
//line /usr/local/go/src/text/template/funcs.go:565
	_go_fuzz_dep_.CoverTab[30427]++
							return truth, nil
//line /usr/local/go/src/text/template/funcs.go:566
	// _ = "end of CoverTab[30427]"
}

// le evaluates the comparison <= b.
func le(arg1, arg2 reflect.Value) (bool, error) {
//line /usr/local/go/src/text/template/funcs.go:570
	_go_fuzz_dep_.CoverTab[30447]++

							lessThan, err := lt(arg1, arg2)
							if lessThan || func() bool {
//line /usr/local/go/src/text/template/funcs.go:573
		_go_fuzz_dep_.CoverTab[30449]++
//line /usr/local/go/src/text/template/funcs.go:573
		return err != nil
//line /usr/local/go/src/text/template/funcs.go:573
		// _ = "end of CoverTab[30449]"
//line /usr/local/go/src/text/template/funcs.go:573
	}() {
//line /usr/local/go/src/text/template/funcs.go:573
		_go_fuzz_dep_.CoverTab[30450]++
								return lessThan, err
//line /usr/local/go/src/text/template/funcs.go:574
		// _ = "end of CoverTab[30450]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:575
		_go_fuzz_dep_.CoverTab[30451]++
//line /usr/local/go/src/text/template/funcs.go:575
		// _ = "end of CoverTab[30451]"
//line /usr/local/go/src/text/template/funcs.go:575
	}
//line /usr/local/go/src/text/template/funcs.go:575
	// _ = "end of CoverTab[30447]"
//line /usr/local/go/src/text/template/funcs.go:575
	_go_fuzz_dep_.CoverTab[30448]++
							return eq(arg1, arg2)
//line /usr/local/go/src/text/template/funcs.go:576
	// _ = "end of CoverTab[30448]"
}

// gt evaluates the comparison a > b.
func gt(arg1, arg2 reflect.Value) (bool, error) {
//line /usr/local/go/src/text/template/funcs.go:580
	_go_fuzz_dep_.CoverTab[30452]++

							lessOrEqual, err := le(arg1, arg2)
							if err != nil {
//line /usr/local/go/src/text/template/funcs.go:583
		_go_fuzz_dep_.CoverTab[30454]++
								return false, err
//line /usr/local/go/src/text/template/funcs.go:584
		// _ = "end of CoverTab[30454]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:585
		_go_fuzz_dep_.CoverTab[30455]++
//line /usr/local/go/src/text/template/funcs.go:585
		// _ = "end of CoverTab[30455]"
//line /usr/local/go/src/text/template/funcs.go:585
	}
//line /usr/local/go/src/text/template/funcs.go:585
	// _ = "end of CoverTab[30452]"
//line /usr/local/go/src/text/template/funcs.go:585
	_go_fuzz_dep_.CoverTab[30453]++
							return !lessOrEqual, nil
//line /usr/local/go/src/text/template/funcs.go:586
	// _ = "end of CoverTab[30453]"
}

// ge evaluates the comparison a >= b.
func ge(arg1, arg2 reflect.Value) (bool, error) {
//line /usr/local/go/src/text/template/funcs.go:590
	_go_fuzz_dep_.CoverTab[30456]++

							lessThan, err := lt(arg1, arg2)
							if err != nil {
//line /usr/local/go/src/text/template/funcs.go:593
		_go_fuzz_dep_.CoverTab[30458]++
								return false, err
//line /usr/local/go/src/text/template/funcs.go:594
		// _ = "end of CoverTab[30458]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:595
		_go_fuzz_dep_.CoverTab[30459]++
//line /usr/local/go/src/text/template/funcs.go:595
		// _ = "end of CoverTab[30459]"
//line /usr/local/go/src/text/template/funcs.go:595
	}
//line /usr/local/go/src/text/template/funcs.go:595
	// _ = "end of CoverTab[30456]"
//line /usr/local/go/src/text/template/funcs.go:595
	_go_fuzz_dep_.CoverTab[30457]++
							return !lessThan, nil
//line /usr/local/go/src/text/template/funcs.go:596
	// _ = "end of CoverTab[30457]"
}

//line /usr/local/go/src/text/template/funcs.go:601
var (
	htmlQuot	= []byte("&#34;")	// shorter than "&quot;"
	htmlApos	= []byte("&#39;")	// shorter than "&apos;" and apos was not in HTML until HTML5
	htmlAmp		= []byte("&amp;")
	htmlLt		= []byte("&lt;")
	htmlGt		= []byte("&gt;")
	htmlNull	= []byte("\uFFFD")
)

// HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
func HTMLEscape(w io.Writer, b []byte) {
//line /usr/local/go/src/text/template/funcs.go:611
	_go_fuzz_dep_.CoverTab[30460]++
							last := 0
							for i, c := range b {
//line /usr/local/go/src/text/template/funcs.go:613
		_go_fuzz_dep_.CoverTab[30462]++
								var html []byte
								switch c {
		case '\000':
//line /usr/local/go/src/text/template/funcs.go:616
			_go_fuzz_dep_.CoverTab[30464]++
									html = htmlNull
//line /usr/local/go/src/text/template/funcs.go:617
			// _ = "end of CoverTab[30464]"
		case '"':
//line /usr/local/go/src/text/template/funcs.go:618
			_go_fuzz_dep_.CoverTab[30465]++
									html = htmlQuot
//line /usr/local/go/src/text/template/funcs.go:619
			// _ = "end of CoverTab[30465]"
		case '\'':
//line /usr/local/go/src/text/template/funcs.go:620
			_go_fuzz_dep_.CoverTab[30466]++
									html = htmlApos
//line /usr/local/go/src/text/template/funcs.go:621
			// _ = "end of CoverTab[30466]"
		case '&':
//line /usr/local/go/src/text/template/funcs.go:622
			_go_fuzz_dep_.CoverTab[30467]++
									html = htmlAmp
//line /usr/local/go/src/text/template/funcs.go:623
			// _ = "end of CoverTab[30467]"
		case '<':
//line /usr/local/go/src/text/template/funcs.go:624
			_go_fuzz_dep_.CoverTab[30468]++
									html = htmlLt
//line /usr/local/go/src/text/template/funcs.go:625
			// _ = "end of CoverTab[30468]"
		case '>':
//line /usr/local/go/src/text/template/funcs.go:626
			_go_fuzz_dep_.CoverTab[30469]++
									html = htmlGt
//line /usr/local/go/src/text/template/funcs.go:627
			// _ = "end of CoverTab[30469]"
		default:
//line /usr/local/go/src/text/template/funcs.go:628
			_go_fuzz_dep_.CoverTab[30470]++
									continue
//line /usr/local/go/src/text/template/funcs.go:629
			// _ = "end of CoverTab[30470]"
		}
//line /usr/local/go/src/text/template/funcs.go:630
		// _ = "end of CoverTab[30462]"
//line /usr/local/go/src/text/template/funcs.go:630
		_go_fuzz_dep_.CoverTab[30463]++
								w.Write(b[last:i])
								w.Write(html)
								last = i + 1
//line /usr/local/go/src/text/template/funcs.go:633
		// _ = "end of CoverTab[30463]"
	}
//line /usr/local/go/src/text/template/funcs.go:634
	// _ = "end of CoverTab[30460]"
//line /usr/local/go/src/text/template/funcs.go:634
	_go_fuzz_dep_.CoverTab[30461]++
							w.Write(b[last:])
//line /usr/local/go/src/text/template/funcs.go:635
	// _ = "end of CoverTab[30461]"
}

// HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
func HTMLEscapeString(s string) string {
//line /usr/local/go/src/text/template/funcs.go:639
	_go_fuzz_dep_.CoverTab[30471]++

							if !strings.ContainsAny(s, "'\"&<>\000") {
//line /usr/local/go/src/text/template/funcs.go:641
		_go_fuzz_dep_.CoverTab[30473]++
								return s
//line /usr/local/go/src/text/template/funcs.go:642
		// _ = "end of CoverTab[30473]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:643
		_go_fuzz_dep_.CoverTab[30474]++
//line /usr/local/go/src/text/template/funcs.go:643
		// _ = "end of CoverTab[30474]"
//line /usr/local/go/src/text/template/funcs.go:643
	}
//line /usr/local/go/src/text/template/funcs.go:643
	// _ = "end of CoverTab[30471]"
//line /usr/local/go/src/text/template/funcs.go:643
	_go_fuzz_dep_.CoverTab[30472]++
							var b strings.Builder
							HTMLEscape(&b, []byte(s))
							return b.String()
//line /usr/local/go/src/text/template/funcs.go:646
	// _ = "end of CoverTab[30472]"
}

// HTMLEscaper returns the escaped HTML equivalent of the textual
//line /usr/local/go/src/text/template/funcs.go:649
// representation of its arguments.
//line /usr/local/go/src/text/template/funcs.go:651
func HTMLEscaper(args ...any) string {
//line /usr/local/go/src/text/template/funcs.go:651
	_go_fuzz_dep_.CoverTab[30475]++
							return HTMLEscapeString(evalArgs(args))
//line /usr/local/go/src/text/template/funcs.go:652
	// _ = "end of CoverTab[30475]"
}

//line /usr/local/go/src/text/template/funcs.go:657
var (
	jsLowUni	= []byte(`\u00`)
	hex		= []byte("0123456789ABCDEF")

	jsBackslash	= []byte(`\\`)
	jsApos		= []byte(`\'`)
	jsQuot		= []byte(`\"`)
	jsLt		= []byte(`\u003C`)
	jsGt		= []byte(`\u003E`)
	jsAmp		= []byte(`\u0026`)
	jsEq		= []byte(`\u003D`)
)

// JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.
func JSEscape(w io.Writer, b []byte) {
//line /usr/local/go/src/text/template/funcs.go:671
	_go_fuzz_dep_.CoverTab[30476]++
							last := 0
							for i := 0; i < len(b); i++ {
//line /usr/local/go/src/text/template/funcs.go:673
		_go_fuzz_dep_.CoverTab[30478]++
								c := b[i]

								if !jsIsSpecial(rune(c)) {
//line /usr/local/go/src/text/template/funcs.go:676
			_go_fuzz_dep_.CoverTab[30481]++

									continue
//line /usr/local/go/src/text/template/funcs.go:678
			// _ = "end of CoverTab[30481]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:679
			_go_fuzz_dep_.CoverTab[30482]++
//line /usr/local/go/src/text/template/funcs.go:679
			// _ = "end of CoverTab[30482]"
//line /usr/local/go/src/text/template/funcs.go:679
		}
//line /usr/local/go/src/text/template/funcs.go:679
		// _ = "end of CoverTab[30478]"
//line /usr/local/go/src/text/template/funcs.go:679
		_go_fuzz_dep_.CoverTab[30479]++
								w.Write(b[last:i])

								if c < utf8.RuneSelf {
//line /usr/local/go/src/text/template/funcs.go:682
			_go_fuzz_dep_.CoverTab[30483]++

//line /usr/local/go/src/text/template/funcs.go:685
			switch c {
			case '\\':
//line /usr/local/go/src/text/template/funcs.go:686
				_go_fuzz_dep_.CoverTab[30484]++
										w.Write(jsBackslash)
//line /usr/local/go/src/text/template/funcs.go:687
				// _ = "end of CoverTab[30484]"
			case '\'':
//line /usr/local/go/src/text/template/funcs.go:688
				_go_fuzz_dep_.CoverTab[30485]++
										w.Write(jsApos)
//line /usr/local/go/src/text/template/funcs.go:689
				// _ = "end of CoverTab[30485]"
			case '"':
//line /usr/local/go/src/text/template/funcs.go:690
				_go_fuzz_dep_.CoverTab[30486]++
										w.Write(jsQuot)
//line /usr/local/go/src/text/template/funcs.go:691
				// _ = "end of CoverTab[30486]"
			case '<':
//line /usr/local/go/src/text/template/funcs.go:692
				_go_fuzz_dep_.CoverTab[30487]++
										w.Write(jsLt)
//line /usr/local/go/src/text/template/funcs.go:693
				// _ = "end of CoverTab[30487]"
			case '>':
//line /usr/local/go/src/text/template/funcs.go:694
				_go_fuzz_dep_.CoverTab[30488]++
										w.Write(jsGt)
//line /usr/local/go/src/text/template/funcs.go:695
				// _ = "end of CoverTab[30488]"
			case '&':
//line /usr/local/go/src/text/template/funcs.go:696
				_go_fuzz_dep_.CoverTab[30489]++
										w.Write(jsAmp)
//line /usr/local/go/src/text/template/funcs.go:697
				// _ = "end of CoverTab[30489]"
			case '=':
//line /usr/local/go/src/text/template/funcs.go:698
				_go_fuzz_dep_.CoverTab[30490]++
										w.Write(jsEq)
//line /usr/local/go/src/text/template/funcs.go:699
				// _ = "end of CoverTab[30490]"
			default:
//line /usr/local/go/src/text/template/funcs.go:700
				_go_fuzz_dep_.CoverTab[30491]++
										w.Write(jsLowUni)
										t, b := c>>4, c&0x0f
										w.Write(hex[t : t+1])
										w.Write(hex[b : b+1])
//line /usr/local/go/src/text/template/funcs.go:704
				// _ = "end of CoverTab[30491]"
			}
//line /usr/local/go/src/text/template/funcs.go:705
			// _ = "end of CoverTab[30483]"
		} else {
//line /usr/local/go/src/text/template/funcs.go:706
			_go_fuzz_dep_.CoverTab[30492]++

									r, size := utf8.DecodeRune(b[i:])
									if unicode.IsPrint(r) {
//line /usr/local/go/src/text/template/funcs.go:709
				_go_fuzz_dep_.CoverTab[30494]++
										w.Write(b[i : i+size])
//line /usr/local/go/src/text/template/funcs.go:710
				// _ = "end of CoverTab[30494]"
			} else {
//line /usr/local/go/src/text/template/funcs.go:711
				_go_fuzz_dep_.CoverTab[30495]++
										fmt.Fprintf(w, "\\u%04X", r)
//line /usr/local/go/src/text/template/funcs.go:712
				// _ = "end of CoverTab[30495]"
			}
//line /usr/local/go/src/text/template/funcs.go:713
			// _ = "end of CoverTab[30492]"
//line /usr/local/go/src/text/template/funcs.go:713
			_go_fuzz_dep_.CoverTab[30493]++
									i += size - 1
//line /usr/local/go/src/text/template/funcs.go:714
			// _ = "end of CoverTab[30493]"
		}
//line /usr/local/go/src/text/template/funcs.go:715
		// _ = "end of CoverTab[30479]"
//line /usr/local/go/src/text/template/funcs.go:715
		_go_fuzz_dep_.CoverTab[30480]++
								last = i + 1
//line /usr/local/go/src/text/template/funcs.go:716
		// _ = "end of CoverTab[30480]"
	}
//line /usr/local/go/src/text/template/funcs.go:717
	// _ = "end of CoverTab[30476]"
//line /usr/local/go/src/text/template/funcs.go:717
	_go_fuzz_dep_.CoverTab[30477]++
							w.Write(b[last:])
//line /usr/local/go/src/text/template/funcs.go:718
	// _ = "end of CoverTab[30477]"
}

// JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.
func JSEscapeString(s string) string {
//line /usr/local/go/src/text/template/funcs.go:722
	_go_fuzz_dep_.CoverTab[30496]++

							if strings.IndexFunc(s, jsIsSpecial) < 0 {
//line /usr/local/go/src/text/template/funcs.go:724
		_go_fuzz_dep_.CoverTab[30498]++
								return s
//line /usr/local/go/src/text/template/funcs.go:725
		// _ = "end of CoverTab[30498]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:726
		_go_fuzz_dep_.CoverTab[30499]++
//line /usr/local/go/src/text/template/funcs.go:726
		// _ = "end of CoverTab[30499]"
//line /usr/local/go/src/text/template/funcs.go:726
	}
//line /usr/local/go/src/text/template/funcs.go:726
	// _ = "end of CoverTab[30496]"
//line /usr/local/go/src/text/template/funcs.go:726
	_go_fuzz_dep_.CoverTab[30497]++
							var b strings.Builder
							JSEscape(&b, []byte(s))
							return b.String()
//line /usr/local/go/src/text/template/funcs.go:729
	// _ = "end of CoverTab[30497]"
}

func jsIsSpecial(r rune) bool {
//line /usr/local/go/src/text/template/funcs.go:732
	_go_fuzz_dep_.CoverTab[30500]++
							switch r {
	case '\\', '\'', '"', '<', '>', '&', '=':
//line /usr/local/go/src/text/template/funcs.go:734
		_go_fuzz_dep_.CoverTab[30502]++
								return true
//line /usr/local/go/src/text/template/funcs.go:735
		// _ = "end of CoverTab[30502]"
//line /usr/local/go/src/text/template/funcs.go:735
	default:
//line /usr/local/go/src/text/template/funcs.go:735
		_go_fuzz_dep_.CoverTab[30503]++
//line /usr/local/go/src/text/template/funcs.go:735
		// _ = "end of CoverTab[30503]"
	}
//line /usr/local/go/src/text/template/funcs.go:736
	// _ = "end of CoverTab[30500]"
//line /usr/local/go/src/text/template/funcs.go:736
	_go_fuzz_dep_.CoverTab[30501]++
							return r < ' ' || func() bool {
//line /usr/local/go/src/text/template/funcs.go:737
		_go_fuzz_dep_.CoverTab[30504]++
//line /usr/local/go/src/text/template/funcs.go:737
		return utf8.RuneSelf <= r
//line /usr/local/go/src/text/template/funcs.go:737
		// _ = "end of CoverTab[30504]"
//line /usr/local/go/src/text/template/funcs.go:737
	}()
//line /usr/local/go/src/text/template/funcs.go:737
	// _ = "end of CoverTab[30501]"
}

// JSEscaper returns the escaped JavaScript equivalent of the textual
//line /usr/local/go/src/text/template/funcs.go:740
// representation of its arguments.
//line /usr/local/go/src/text/template/funcs.go:742
func JSEscaper(args ...any) string {
//line /usr/local/go/src/text/template/funcs.go:742
	_go_fuzz_dep_.CoverTab[30505]++
							return JSEscapeString(evalArgs(args))
//line /usr/local/go/src/text/template/funcs.go:743
	// _ = "end of CoverTab[30505]"
}

// URLQueryEscaper returns the escaped value of the textual representation of
//line /usr/local/go/src/text/template/funcs.go:746
// its arguments in a form suitable for embedding in a URL query.
//line /usr/local/go/src/text/template/funcs.go:748
func URLQueryEscaper(args ...any) string {
//line /usr/local/go/src/text/template/funcs.go:748
	_go_fuzz_dep_.CoverTab[30506]++
							return url.QueryEscape(evalArgs(args))
//line /usr/local/go/src/text/template/funcs.go:749
	// _ = "end of CoverTab[30506]"
}

// evalArgs formats the list of arguments into a string. It is therefore equivalent to
//line /usr/local/go/src/text/template/funcs.go:752
//
//line /usr/local/go/src/text/template/funcs.go:752
//	fmt.Sprint(args...)
//line /usr/local/go/src/text/template/funcs.go:752
//
//line /usr/local/go/src/text/template/funcs.go:752
// except that each argument is indirected (if a pointer), as required,
//line /usr/local/go/src/text/template/funcs.go:752
// using the same rules as the default string evaluation during template
//line /usr/local/go/src/text/template/funcs.go:752
// execution.
//line /usr/local/go/src/text/template/funcs.go:759
func evalArgs(args []any) string {
//line /usr/local/go/src/text/template/funcs.go:759
	_go_fuzz_dep_.CoverTab[30507]++
							ok := false
							var s string

							if len(args) == 1 {
//line /usr/local/go/src/text/template/funcs.go:763
		_go_fuzz_dep_.CoverTab[30510]++
								s, ok = args[0].(string)
//line /usr/local/go/src/text/template/funcs.go:764
		// _ = "end of CoverTab[30510]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:765
		_go_fuzz_dep_.CoverTab[30511]++
//line /usr/local/go/src/text/template/funcs.go:765
		// _ = "end of CoverTab[30511]"
//line /usr/local/go/src/text/template/funcs.go:765
	}
//line /usr/local/go/src/text/template/funcs.go:765
	// _ = "end of CoverTab[30507]"
//line /usr/local/go/src/text/template/funcs.go:765
	_go_fuzz_dep_.CoverTab[30508]++
							if !ok {
//line /usr/local/go/src/text/template/funcs.go:766
		_go_fuzz_dep_.CoverTab[30512]++
								for i, arg := range args {
//line /usr/local/go/src/text/template/funcs.go:767
			_go_fuzz_dep_.CoverTab[30514]++
									a, ok := printableValue(reflect.ValueOf(arg))
									if ok {
//line /usr/local/go/src/text/template/funcs.go:769
				_go_fuzz_dep_.CoverTab[30515]++
										args[i] = a
//line /usr/local/go/src/text/template/funcs.go:770
				// _ = "end of CoverTab[30515]"
			} else {
//line /usr/local/go/src/text/template/funcs.go:771
				_go_fuzz_dep_.CoverTab[30516]++
//line /usr/local/go/src/text/template/funcs.go:771
				// _ = "end of CoverTab[30516]"
//line /usr/local/go/src/text/template/funcs.go:771
			}
//line /usr/local/go/src/text/template/funcs.go:771
			// _ = "end of CoverTab[30514]"
		}
//line /usr/local/go/src/text/template/funcs.go:772
		// _ = "end of CoverTab[30512]"
//line /usr/local/go/src/text/template/funcs.go:772
		_go_fuzz_dep_.CoverTab[30513]++
								s = fmt.Sprint(args...)
//line /usr/local/go/src/text/template/funcs.go:773
		// _ = "end of CoverTab[30513]"
	} else {
//line /usr/local/go/src/text/template/funcs.go:774
		_go_fuzz_dep_.CoverTab[30517]++
//line /usr/local/go/src/text/template/funcs.go:774
		// _ = "end of CoverTab[30517]"
//line /usr/local/go/src/text/template/funcs.go:774
	}
//line /usr/local/go/src/text/template/funcs.go:774
	// _ = "end of CoverTab[30508]"
//line /usr/local/go/src/text/template/funcs.go:774
	_go_fuzz_dep_.CoverTab[30509]++
							return s
//line /usr/local/go/src/text/template/funcs.go:775
	// _ = "end of CoverTab[30509]"
}

//line /usr/local/go/src/text/template/funcs.go:776
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/template/funcs.go:776
var _ = _go_fuzz_dep_.CoverTab
