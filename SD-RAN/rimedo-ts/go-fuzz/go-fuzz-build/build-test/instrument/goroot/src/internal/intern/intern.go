// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/internal/intern/intern.go:11
package intern

//line /usr/local/go/src/internal/intern/intern.go:11
import (
//line /usr/local/go/src/internal/intern/intern.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/internal/intern/intern.go:11
)
//line /usr/local/go/src/internal/intern/intern.go:11
import (
//line /usr/local/go/src/internal/intern/intern.go:11
	_atomic_ "sync/atomic"
//line /usr/local/go/src/internal/intern/intern.go:11
)

import (
	"internal/godebug"
	"runtime"
	"sync"
	"unsafe"
)

//line /usr/local/go/src/internal/intern/intern.go:22
type Value struct {
							_	[0]func()
							cmpVal	any

//line /usr/local/go/src/internal/intern/intern.go:27
	resurrected	bool
}

//line /usr/local/go/src/internal/intern/intern.go:32
func (v *Value) Get() any {
//line /usr/local/go/src/internal/intern/intern.go:32
	_go_fuzz_dep_.CoverTab[3515]++
//line /usr/local/go/src/internal/intern/intern.go:32
	return v.cmpVal
//line /usr/local/go/src/internal/intern/intern.go:32
	// _ = "end of CoverTab[3515]"
//line /usr/local/go/src/internal/intern/intern.go:32
}

//line /usr/local/go/src/internal/intern/intern.go:37
type key struct {
							s	string
							cmpVal	any

//line /usr/local/go/src/internal/intern/intern.go:42
	isString	bool
}

//line /usr/local/go/src/internal/intern/intern.go:46
func keyFor(cmpVal any) key {
//line /usr/local/go/src/internal/intern/intern.go:46
	_go_fuzz_dep_.CoverTab[3516]++
							if s, ok := cmpVal.(string); ok {
//line /usr/local/go/src/internal/intern/intern.go:47
		_go_fuzz_dep_.CoverTab[3518]++
								return key{s: s, isString: true}
//line /usr/local/go/src/internal/intern/intern.go:48
		// _ = "end of CoverTab[3518]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:49
		_go_fuzz_dep_.CoverTab[3519]++
//line /usr/local/go/src/internal/intern/intern.go:49
		// _ = "end of CoverTab[3519]"
//line /usr/local/go/src/internal/intern/intern.go:49
	}
//line /usr/local/go/src/internal/intern/intern.go:49
	// _ = "end of CoverTab[3516]"
//line /usr/local/go/src/internal/intern/intern.go:49
	_go_fuzz_dep_.CoverTab[3517]++
							return key{cmpVal: cmpVal}
//line /usr/local/go/src/internal/intern/intern.go:50
	// _ = "end of CoverTab[3517]"
}

//line /usr/local/go/src/internal/intern/intern.go:54
func (k key) Value() *Value {
//line /usr/local/go/src/internal/intern/intern.go:54
	_go_fuzz_dep_.CoverTab[3520]++
							if k.isString {
//line /usr/local/go/src/internal/intern/intern.go:55
		_go_fuzz_dep_.CoverTab[3522]++
								return &Value{cmpVal: k.s}
//line /usr/local/go/src/internal/intern/intern.go:56
		// _ = "end of CoverTab[3522]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:57
		_go_fuzz_dep_.CoverTab[3523]++
//line /usr/local/go/src/internal/intern/intern.go:57
		// _ = "end of CoverTab[3523]"
//line /usr/local/go/src/internal/intern/intern.go:57
	}
//line /usr/local/go/src/internal/intern/intern.go:57
	// _ = "end of CoverTab[3520]"
//line /usr/local/go/src/internal/intern/intern.go:57
	_go_fuzz_dep_.CoverTab[3521]++
							return &Value{cmpVal: k.cmpVal}
//line /usr/local/go/src/internal/intern/intern.go:58
	// _ = "end of CoverTab[3521]"
}

var (
//line /usr/local/go/src/internal/intern/intern.go:64
	mu	sync.Mutex
	valMap	= map[key]uintptr{}
	valSafe	= safeMap()
)

var intern = godebug.New("intern")

//line /usr/local/go/src/internal/intern/intern.go:73
func safeMap() map[key]*Value {
//line /usr/local/go/src/internal/intern/intern.go:73
	_go_fuzz_dep_.CoverTab[3524]++
							if intern.Value() == "leaky" {
//line /usr/local/go/src/internal/intern/intern.go:74
		_go_fuzz_dep_.CoverTab[3526]++
								return map[key]*Value{}
//line /usr/local/go/src/internal/intern/intern.go:75
		// _ = "end of CoverTab[3526]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:76
		_go_fuzz_dep_.CoverTab[3527]++
//line /usr/local/go/src/internal/intern/intern.go:76
		// _ = "end of CoverTab[3527]"
//line /usr/local/go/src/internal/intern/intern.go:76
	}
//line /usr/local/go/src/internal/intern/intern.go:76
	// _ = "end of CoverTab[3524]"
//line /usr/local/go/src/internal/intern/intern.go:76
	_go_fuzz_dep_.CoverTab[3525]++
							return nil
//line /usr/local/go/src/internal/intern/intern.go:77
	// _ = "end of CoverTab[3525]"
}

//line /usr/local/go/src/internal/intern/intern.go:84
func Get(cmpVal any) *Value {
//line /usr/local/go/src/internal/intern/intern.go:84
	_go_fuzz_dep_.CoverTab[3528]++
							return get(keyFor(cmpVal))
//line /usr/local/go/src/internal/intern/intern.go:85
	// _ = "end of CoverTab[3528]"
}

//line /usr/local/go/src/internal/intern/intern.go:91
func GetByString(s string) *Value {
//line /usr/local/go/src/internal/intern/intern.go:91
	_go_fuzz_dep_.CoverTab[3529]++
							return get(key{s: s, isString: true})
//line /usr/local/go/src/internal/intern/intern.go:92
	// _ = "end of CoverTab[3529]"
}

//line /usr/local/go/src/internal/intern/intern.go:99
//go:nocheckptr
func get(k key) *Value {
//line /usr/local/go/src/internal/intern/intern.go:100
	_go_fuzz_dep_.CoverTab[3530]++
							mu.Lock()
							defer mu.Unlock()

							var v *Value
							if valSafe != nil {
//line /usr/local/go/src/internal/intern/intern.go:105
		_go_fuzz_dep_.CoverTab[3534]++
								v = valSafe[k]
//line /usr/local/go/src/internal/intern/intern.go:106
		// _ = "end of CoverTab[3534]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:107
		_go_fuzz_dep_.CoverTab[3535]++
//line /usr/local/go/src/internal/intern/intern.go:107
		if addr, ok := valMap[k]; ok {
//line /usr/local/go/src/internal/intern/intern.go:107
			_go_fuzz_dep_.CoverTab[3536]++
									v = (*Value)(unsafe.Pointer(addr))
									v.resurrected = true
//line /usr/local/go/src/internal/intern/intern.go:109
			// _ = "end of CoverTab[3536]"
		} else {
//line /usr/local/go/src/internal/intern/intern.go:110
			_go_fuzz_dep_.CoverTab[3537]++
//line /usr/local/go/src/internal/intern/intern.go:110
			// _ = "end of CoverTab[3537]"
//line /usr/local/go/src/internal/intern/intern.go:110
		}
//line /usr/local/go/src/internal/intern/intern.go:110
		// _ = "end of CoverTab[3535]"
//line /usr/local/go/src/internal/intern/intern.go:110
	}
//line /usr/local/go/src/internal/intern/intern.go:110
	// _ = "end of CoverTab[3530]"
//line /usr/local/go/src/internal/intern/intern.go:110
	_go_fuzz_dep_.CoverTab[3531]++
							if v != nil {
//line /usr/local/go/src/internal/intern/intern.go:111
		_go_fuzz_dep_.CoverTab[3538]++
								return v
//line /usr/local/go/src/internal/intern/intern.go:112
		// _ = "end of CoverTab[3538]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:113
		_go_fuzz_dep_.CoverTab[3539]++
//line /usr/local/go/src/internal/intern/intern.go:113
		// _ = "end of CoverTab[3539]"
//line /usr/local/go/src/internal/intern/intern.go:113
	}
//line /usr/local/go/src/internal/intern/intern.go:113
	// _ = "end of CoverTab[3531]"
//line /usr/local/go/src/internal/intern/intern.go:113
	_go_fuzz_dep_.CoverTab[3532]++
							v = k.Value()
							if valSafe != nil {
//line /usr/local/go/src/internal/intern/intern.go:115
		_go_fuzz_dep_.CoverTab[3540]++
								valSafe[k] = v
//line /usr/local/go/src/internal/intern/intern.go:116
		// _ = "end of CoverTab[3540]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:117
		_go_fuzz_dep_.CoverTab[3541]++

//line /usr/local/go/src/internal/intern/intern.go:120
		runtime.SetFinalizer(v, finalize)
								valMap[k] = uintptr(unsafe.Pointer(v))
//line /usr/local/go/src/internal/intern/intern.go:121
		// _ = "end of CoverTab[3541]"
	}
//line /usr/local/go/src/internal/intern/intern.go:122
	// _ = "end of CoverTab[3532]"
//line /usr/local/go/src/internal/intern/intern.go:122
	_go_fuzz_dep_.CoverTab[3533]++
							return v
//line /usr/local/go/src/internal/intern/intern.go:123
	// _ = "end of CoverTab[3533]"
}

func finalize(v *Value) {
//line /usr/local/go/src/internal/intern/intern.go:126
	_go_fuzz_dep_.CoverTab[3542]++
							mu.Lock()
							defer mu.Unlock()
							if v.resurrected {
//line /usr/local/go/src/internal/intern/intern.go:129
		_go_fuzz_dep_.CoverTab[3544]++

//line /usr/local/go/src/internal/intern/intern.go:132
		v.resurrected = false
								runtime.SetFinalizer(v, finalize)
								return
//line /usr/local/go/src/internal/intern/intern.go:134
		// _ = "end of CoverTab[3544]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:135
		_go_fuzz_dep_.CoverTab[3545]++
//line /usr/local/go/src/internal/intern/intern.go:135
		// _ = "end of CoverTab[3545]"
//line /usr/local/go/src/internal/intern/intern.go:135
	}
//line /usr/local/go/src/internal/intern/intern.go:135
	// _ = "end of CoverTab[3542]"
//line /usr/local/go/src/internal/intern/intern.go:135
	_go_fuzz_dep_.CoverTab[3543]++
							delete(valMap, keyFor(v.cmpVal))
//line /usr/local/go/src/internal/intern/intern.go:136
	// _ = "end of CoverTab[3543]"
}

//line /usr/local/go/src/internal/intern/intern.go:137
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/intern/intern.go:137
var _ = _go_fuzz_dep_.CoverTab
