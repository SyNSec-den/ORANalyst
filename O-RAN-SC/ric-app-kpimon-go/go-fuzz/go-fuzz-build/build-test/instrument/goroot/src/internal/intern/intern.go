// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/internal/intern/intern.go:11
package intern

//line /snap/go/10455/src/internal/intern/intern.go:11
import (
//line /snap/go/10455/src/internal/intern/intern.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/internal/intern/intern.go:11
)
//line /snap/go/10455/src/internal/intern/intern.go:11
import (
//line /snap/go/10455/src/internal/intern/intern.go:11
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/internal/intern/intern.go:11
)

import (
	"internal/godebug"
	"runtime"
	"sync"
	"unsafe"
)

//line /snap/go/10455/src/internal/intern/intern.go:22
type Value struct {
							_	[0]func()
							cmpVal	any

//line /snap/go/10455/src/internal/intern/intern.go:27
	resurrected	bool
}

//line /snap/go/10455/src/internal/intern/intern.go:32
func (v *Value) Get() any {
//line /snap/go/10455/src/internal/intern/intern.go:32
	_go_fuzz_dep_.CoverTab[3851]++
//line /snap/go/10455/src/internal/intern/intern.go:32
	return v.cmpVal
//line /snap/go/10455/src/internal/intern/intern.go:32
	// _ = "end of CoverTab[3851]"
//line /snap/go/10455/src/internal/intern/intern.go:32
}

//line /snap/go/10455/src/internal/intern/intern.go:37
type key struct {
							s	string
							cmpVal	any

//line /snap/go/10455/src/internal/intern/intern.go:42
	isString	bool
}

//line /snap/go/10455/src/internal/intern/intern.go:46
func keyFor(cmpVal any) key {
//line /snap/go/10455/src/internal/intern/intern.go:46
	_go_fuzz_dep_.CoverTab[3852]++
							if s, ok := cmpVal.(string); ok {
//line /snap/go/10455/src/internal/intern/intern.go:47
		_go_fuzz_dep_.CoverTab[526996]++
//line /snap/go/10455/src/internal/intern/intern.go:47
		_go_fuzz_dep_.CoverTab[3854]++
								return key{s: s, isString: true}
//line /snap/go/10455/src/internal/intern/intern.go:48
		// _ = "end of CoverTab[3854]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:49
		_go_fuzz_dep_.CoverTab[526997]++
//line /snap/go/10455/src/internal/intern/intern.go:49
		_go_fuzz_dep_.CoverTab[3855]++
//line /snap/go/10455/src/internal/intern/intern.go:49
		// _ = "end of CoverTab[3855]"
//line /snap/go/10455/src/internal/intern/intern.go:49
	}
//line /snap/go/10455/src/internal/intern/intern.go:49
	// _ = "end of CoverTab[3852]"
//line /snap/go/10455/src/internal/intern/intern.go:49
	_go_fuzz_dep_.CoverTab[3853]++
							return key{cmpVal: cmpVal}
//line /snap/go/10455/src/internal/intern/intern.go:50
	// _ = "end of CoverTab[3853]"
}

//line /snap/go/10455/src/internal/intern/intern.go:54
func (k key) Value() *Value {
//line /snap/go/10455/src/internal/intern/intern.go:54
	_go_fuzz_dep_.CoverTab[3856]++
							if k.isString {
//line /snap/go/10455/src/internal/intern/intern.go:55
		_go_fuzz_dep_.CoverTab[526998]++
//line /snap/go/10455/src/internal/intern/intern.go:55
		_go_fuzz_dep_.CoverTab[3858]++
								return &Value{cmpVal: k.s}
//line /snap/go/10455/src/internal/intern/intern.go:56
		// _ = "end of CoverTab[3858]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:57
		_go_fuzz_dep_.CoverTab[526999]++
//line /snap/go/10455/src/internal/intern/intern.go:57
		_go_fuzz_dep_.CoverTab[3859]++
//line /snap/go/10455/src/internal/intern/intern.go:57
		// _ = "end of CoverTab[3859]"
//line /snap/go/10455/src/internal/intern/intern.go:57
	}
//line /snap/go/10455/src/internal/intern/intern.go:57
	// _ = "end of CoverTab[3856]"
//line /snap/go/10455/src/internal/intern/intern.go:57
	_go_fuzz_dep_.CoverTab[3857]++
							return &Value{cmpVal: k.cmpVal}
//line /snap/go/10455/src/internal/intern/intern.go:58
	// _ = "end of CoverTab[3857]"
}

var (
//line /snap/go/10455/src/internal/intern/intern.go:64
	mu	sync.Mutex
	valMap	= map[key]uintptr{}
	valSafe	= safeMap()
)

var intern = godebug.New("#intern")

//line /snap/go/10455/src/internal/intern/intern.go:73
func safeMap() map[key]*Value {
//line /snap/go/10455/src/internal/intern/intern.go:73
	_go_fuzz_dep_.CoverTab[3860]++
							if intern.Value() == "leaky" {
//line /snap/go/10455/src/internal/intern/intern.go:74
		_go_fuzz_dep_.CoverTab[527000]++
//line /snap/go/10455/src/internal/intern/intern.go:74
		_go_fuzz_dep_.CoverTab[3862]++
								return map[key]*Value{}
//line /snap/go/10455/src/internal/intern/intern.go:75
		// _ = "end of CoverTab[3862]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:76
		_go_fuzz_dep_.CoverTab[527001]++
//line /snap/go/10455/src/internal/intern/intern.go:76
		_go_fuzz_dep_.CoverTab[3863]++
//line /snap/go/10455/src/internal/intern/intern.go:76
		// _ = "end of CoverTab[3863]"
//line /snap/go/10455/src/internal/intern/intern.go:76
	}
//line /snap/go/10455/src/internal/intern/intern.go:76
	// _ = "end of CoverTab[3860]"
//line /snap/go/10455/src/internal/intern/intern.go:76
	_go_fuzz_dep_.CoverTab[3861]++
							return nil
//line /snap/go/10455/src/internal/intern/intern.go:77
	// _ = "end of CoverTab[3861]"
}

//line /snap/go/10455/src/internal/intern/intern.go:84
func Get(cmpVal any) *Value {
//line /snap/go/10455/src/internal/intern/intern.go:84
	_go_fuzz_dep_.CoverTab[3864]++
							return get(keyFor(cmpVal))
//line /snap/go/10455/src/internal/intern/intern.go:85
	// _ = "end of CoverTab[3864]"
}

//line /snap/go/10455/src/internal/intern/intern.go:91
func GetByString(s string) *Value {
//line /snap/go/10455/src/internal/intern/intern.go:91
	_go_fuzz_dep_.CoverTab[3865]++
							return get(key{s: s, isString: true})
//line /snap/go/10455/src/internal/intern/intern.go:92
	// _ = "end of CoverTab[3865]"
}

//line /snap/go/10455/src/internal/intern/intern.go:99
//go:nocheckptr
func get(k key) *Value {
//line /snap/go/10455/src/internal/intern/intern.go:100
	_go_fuzz_dep_.CoverTab[3866]++
								mu.Lock()
								defer mu.Unlock()

								var v *Value
								if valSafe != nil {
//line /snap/go/10455/src/internal/intern/intern.go:105
		_go_fuzz_dep_.CoverTab[527002]++
//line /snap/go/10455/src/internal/intern/intern.go:105
		_go_fuzz_dep_.CoverTab[3870]++
									v = valSafe[k]
//line /snap/go/10455/src/internal/intern/intern.go:106
		// _ = "end of CoverTab[3870]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:107
		_go_fuzz_dep_.CoverTab[527003]++
//line /snap/go/10455/src/internal/intern/intern.go:107
		_go_fuzz_dep_.CoverTab[3871]++
//line /snap/go/10455/src/internal/intern/intern.go:107
		if addr, ok := valMap[k]; ok {
//line /snap/go/10455/src/internal/intern/intern.go:107
			_go_fuzz_dep_.CoverTab[527004]++
//line /snap/go/10455/src/internal/intern/intern.go:107
			_go_fuzz_dep_.CoverTab[3872]++
										v = (*Value)(unsafe.Pointer(addr))
										v.resurrected = true
//line /snap/go/10455/src/internal/intern/intern.go:109
			// _ = "end of CoverTab[3872]"
		} else {
//line /snap/go/10455/src/internal/intern/intern.go:110
			_go_fuzz_dep_.CoverTab[527005]++
//line /snap/go/10455/src/internal/intern/intern.go:110
			_go_fuzz_dep_.CoverTab[3873]++
//line /snap/go/10455/src/internal/intern/intern.go:110
			// _ = "end of CoverTab[3873]"
//line /snap/go/10455/src/internal/intern/intern.go:110
		}
//line /snap/go/10455/src/internal/intern/intern.go:110
		// _ = "end of CoverTab[3871]"
//line /snap/go/10455/src/internal/intern/intern.go:110
	}
//line /snap/go/10455/src/internal/intern/intern.go:110
	// _ = "end of CoverTab[3866]"
//line /snap/go/10455/src/internal/intern/intern.go:110
	_go_fuzz_dep_.CoverTab[3867]++
								if v != nil {
//line /snap/go/10455/src/internal/intern/intern.go:111
		_go_fuzz_dep_.CoverTab[527006]++
//line /snap/go/10455/src/internal/intern/intern.go:111
		_go_fuzz_dep_.CoverTab[3874]++
									return v
//line /snap/go/10455/src/internal/intern/intern.go:112
		// _ = "end of CoverTab[3874]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:113
		_go_fuzz_dep_.CoverTab[527007]++
//line /snap/go/10455/src/internal/intern/intern.go:113
		_go_fuzz_dep_.CoverTab[3875]++
//line /snap/go/10455/src/internal/intern/intern.go:113
		// _ = "end of CoverTab[3875]"
//line /snap/go/10455/src/internal/intern/intern.go:113
	}
//line /snap/go/10455/src/internal/intern/intern.go:113
	// _ = "end of CoverTab[3867]"
//line /snap/go/10455/src/internal/intern/intern.go:113
	_go_fuzz_dep_.CoverTab[3868]++
								v = k.Value()
								if valSafe != nil {
//line /snap/go/10455/src/internal/intern/intern.go:115
		_go_fuzz_dep_.CoverTab[527008]++
//line /snap/go/10455/src/internal/intern/intern.go:115
		_go_fuzz_dep_.CoverTab[3876]++
									valSafe[k] = v
//line /snap/go/10455/src/internal/intern/intern.go:116
		// _ = "end of CoverTab[3876]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:117
		_go_fuzz_dep_.CoverTab[527009]++
//line /snap/go/10455/src/internal/intern/intern.go:117
		_go_fuzz_dep_.CoverTab[3877]++

//line /snap/go/10455/src/internal/intern/intern.go:120
		runtime.SetFinalizer(v, finalize)
									valMap[k] = uintptr(unsafe.Pointer(v))
//line /snap/go/10455/src/internal/intern/intern.go:121
		// _ = "end of CoverTab[3877]"
	}
//line /snap/go/10455/src/internal/intern/intern.go:122
	// _ = "end of CoverTab[3868]"
//line /snap/go/10455/src/internal/intern/intern.go:122
	_go_fuzz_dep_.CoverTab[3869]++
								return v
//line /snap/go/10455/src/internal/intern/intern.go:123
	// _ = "end of CoverTab[3869]"
}

func finalize(v *Value) {
//line /snap/go/10455/src/internal/intern/intern.go:126
	_go_fuzz_dep_.CoverTab[3878]++
								mu.Lock()
								defer mu.Unlock()
								if v.resurrected {
//line /snap/go/10455/src/internal/intern/intern.go:129
		_go_fuzz_dep_.CoverTab[527010]++
//line /snap/go/10455/src/internal/intern/intern.go:129
		_go_fuzz_dep_.CoverTab[3880]++

//line /snap/go/10455/src/internal/intern/intern.go:132
		v.resurrected = false
									runtime.SetFinalizer(v, finalize)
									return
//line /snap/go/10455/src/internal/intern/intern.go:134
		// _ = "end of CoverTab[3880]"
	} else {
//line /snap/go/10455/src/internal/intern/intern.go:135
		_go_fuzz_dep_.CoverTab[527011]++
//line /snap/go/10455/src/internal/intern/intern.go:135
		_go_fuzz_dep_.CoverTab[3881]++
//line /snap/go/10455/src/internal/intern/intern.go:135
		// _ = "end of CoverTab[3881]"
//line /snap/go/10455/src/internal/intern/intern.go:135
	}
//line /snap/go/10455/src/internal/intern/intern.go:135
	// _ = "end of CoverTab[3878]"
//line /snap/go/10455/src/internal/intern/intern.go:135
	_go_fuzz_dep_.CoverTab[3879]++
								delete(valMap, keyFor(v.cmpVal))
//line /snap/go/10455/src/internal/intern/intern.go:136
	// _ = "end of CoverTab[3879]"
}

//line /snap/go/10455/src/internal/intern/intern.go:137
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/internal/intern/intern.go:137
var _ = _go_fuzz_dep_.CoverTab
