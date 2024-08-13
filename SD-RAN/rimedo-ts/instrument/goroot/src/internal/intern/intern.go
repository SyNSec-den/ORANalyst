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
	_go_fuzz_dep_.CoverTab[11905]++
//line /usr/local/go/src/internal/intern/intern.go:32
	return v.cmpVal
//line /usr/local/go/src/internal/intern/intern.go:32
	// _ = "end of CoverTab[11905]"
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
	_go_fuzz_dep_.CoverTab[11906]++
							if s, ok := cmpVal.(string); ok {
//line /usr/local/go/src/internal/intern/intern.go:47
		_go_fuzz_dep_.CoverTab[11908]++
								return key{s: s, isString: true}
//line /usr/local/go/src/internal/intern/intern.go:48
		// _ = "end of CoverTab[11908]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:49
		_go_fuzz_dep_.CoverTab[11909]++
//line /usr/local/go/src/internal/intern/intern.go:49
		// _ = "end of CoverTab[11909]"
//line /usr/local/go/src/internal/intern/intern.go:49
	}
//line /usr/local/go/src/internal/intern/intern.go:49
	// _ = "end of CoverTab[11906]"
//line /usr/local/go/src/internal/intern/intern.go:49
	_go_fuzz_dep_.CoverTab[11907]++
							return key{cmpVal: cmpVal}
//line /usr/local/go/src/internal/intern/intern.go:50
	// _ = "end of CoverTab[11907]"
}

//line /usr/local/go/src/internal/intern/intern.go:54
func (k key) Value() *Value {
//line /usr/local/go/src/internal/intern/intern.go:54
	_go_fuzz_dep_.CoverTab[11910]++
							if k.isString {
//line /usr/local/go/src/internal/intern/intern.go:55
		_go_fuzz_dep_.CoverTab[11912]++
								return &Value{cmpVal: k.s}
//line /usr/local/go/src/internal/intern/intern.go:56
		// _ = "end of CoverTab[11912]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:57
		_go_fuzz_dep_.CoverTab[11913]++
//line /usr/local/go/src/internal/intern/intern.go:57
		// _ = "end of CoverTab[11913]"
//line /usr/local/go/src/internal/intern/intern.go:57
	}
//line /usr/local/go/src/internal/intern/intern.go:57
	// _ = "end of CoverTab[11910]"
//line /usr/local/go/src/internal/intern/intern.go:57
	_go_fuzz_dep_.CoverTab[11911]++
							return &Value{cmpVal: k.cmpVal}
//line /usr/local/go/src/internal/intern/intern.go:58
	// _ = "end of CoverTab[11911]"
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
	_go_fuzz_dep_.CoverTab[11914]++
							if intern.Value() == "leaky" {
//line /usr/local/go/src/internal/intern/intern.go:74
		_go_fuzz_dep_.CoverTab[11916]++
								return map[key]*Value{}
//line /usr/local/go/src/internal/intern/intern.go:75
		// _ = "end of CoverTab[11916]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:76
		_go_fuzz_dep_.CoverTab[11917]++
//line /usr/local/go/src/internal/intern/intern.go:76
		// _ = "end of CoverTab[11917]"
//line /usr/local/go/src/internal/intern/intern.go:76
	}
//line /usr/local/go/src/internal/intern/intern.go:76
	// _ = "end of CoverTab[11914]"
//line /usr/local/go/src/internal/intern/intern.go:76
	_go_fuzz_dep_.CoverTab[11915]++
							return nil
//line /usr/local/go/src/internal/intern/intern.go:77
	// _ = "end of CoverTab[11915]"
}

//line /usr/local/go/src/internal/intern/intern.go:84
func Get(cmpVal any) *Value {
//line /usr/local/go/src/internal/intern/intern.go:84
	_go_fuzz_dep_.CoverTab[11918]++
							return get(keyFor(cmpVal))
//line /usr/local/go/src/internal/intern/intern.go:85
	// _ = "end of CoverTab[11918]"
}

//line /usr/local/go/src/internal/intern/intern.go:91
func GetByString(s string) *Value {
//line /usr/local/go/src/internal/intern/intern.go:91
	_go_fuzz_dep_.CoverTab[11919]++
							return get(key{s: s, isString: true})
//line /usr/local/go/src/internal/intern/intern.go:92
	// _ = "end of CoverTab[11919]"
}

//line /usr/local/go/src/internal/intern/intern.go:99
//go:nocheckptr
func get(k key) *Value {
//line /usr/local/go/src/internal/intern/intern.go:100
	_go_fuzz_dep_.CoverTab[11920]++
							mu.Lock()
							defer mu.Unlock()

							var v *Value
							if valSafe != nil {
//line /usr/local/go/src/internal/intern/intern.go:105
		_go_fuzz_dep_.CoverTab[11924]++
								v = valSafe[k]
//line /usr/local/go/src/internal/intern/intern.go:106
		// _ = "end of CoverTab[11924]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:107
		_go_fuzz_dep_.CoverTab[11925]++
//line /usr/local/go/src/internal/intern/intern.go:107
		if addr, ok := valMap[k]; ok {
//line /usr/local/go/src/internal/intern/intern.go:107
			_go_fuzz_dep_.CoverTab[11926]++
									v = (*Value)(unsafe.Pointer(addr))
									v.resurrected = true
//line /usr/local/go/src/internal/intern/intern.go:109
			// _ = "end of CoverTab[11926]"
		} else {
//line /usr/local/go/src/internal/intern/intern.go:110
			_go_fuzz_dep_.CoverTab[11927]++
//line /usr/local/go/src/internal/intern/intern.go:110
			// _ = "end of CoverTab[11927]"
//line /usr/local/go/src/internal/intern/intern.go:110
		}
//line /usr/local/go/src/internal/intern/intern.go:110
		// _ = "end of CoverTab[11925]"
//line /usr/local/go/src/internal/intern/intern.go:110
	}
//line /usr/local/go/src/internal/intern/intern.go:110
	// _ = "end of CoverTab[11920]"
//line /usr/local/go/src/internal/intern/intern.go:110
	_go_fuzz_dep_.CoverTab[11921]++
							if v != nil {
//line /usr/local/go/src/internal/intern/intern.go:111
		_go_fuzz_dep_.CoverTab[11928]++
								return v
//line /usr/local/go/src/internal/intern/intern.go:112
		// _ = "end of CoverTab[11928]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:113
		_go_fuzz_dep_.CoverTab[11929]++
//line /usr/local/go/src/internal/intern/intern.go:113
		// _ = "end of CoverTab[11929]"
//line /usr/local/go/src/internal/intern/intern.go:113
	}
//line /usr/local/go/src/internal/intern/intern.go:113
	// _ = "end of CoverTab[11921]"
//line /usr/local/go/src/internal/intern/intern.go:113
	_go_fuzz_dep_.CoverTab[11922]++
							v = k.Value()
							if valSafe != nil {
//line /usr/local/go/src/internal/intern/intern.go:115
		_go_fuzz_dep_.CoverTab[11930]++
								valSafe[k] = v
//line /usr/local/go/src/internal/intern/intern.go:116
		// _ = "end of CoverTab[11930]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:117
		_go_fuzz_dep_.CoverTab[11931]++

//line /usr/local/go/src/internal/intern/intern.go:120
		runtime.SetFinalizer(v, finalize)
								valMap[k] = uintptr(unsafe.Pointer(v))
//line /usr/local/go/src/internal/intern/intern.go:121
		// _ = "end of CoverTab[11931]"
	}
//line /usr/local/go/src/internal/intern/intern.go:122
	// _ = "end of CoverTab[11922]"
//line /usr/local/go/src/internal/intern/intern.go:122
	_go_fuzz_dep_.CoverTab[11923]++
							return v
//line /usr/local/go/src/internal/intern/intern.go:123
	// _ = "end of CoverTab[11923]"
}

func finalize(v *Value) {
//line /usr/local/go/src/internal/intern/intern.go:126
	_go_fuzz_dep_.CoverTab[11932]++
							mu.Lock()
							defer mu.Unlock()
							if v.resurrected {
//line /usr/local/go/src/internal/intern/intern.go:129
		_go_fuzz_dep_.CoverTab[11934]++

//line /usr/local/go/src/internal/intern/intern.go:132
		v.resurrected = false
								runtime.SetFinalizer(v, finalize)
								return
//line /usr/local/go/src/internal/intern/intern.go:134
		// _ = "end of CoverTab[11934]"
	} else {
//line /usr/local/go/src/internal/intern/intern.go:135
		_go_fuzz_dep_.CoverTab[11935]++
//line /usr/local/go/src/internal/intern/intern.go:135
		// _ = "end of CoverTab[11935]"
//line /usr/local/go/src/internal/intern/intern.go:135
	}
//line /usr/local/go/src/internal/intern/intern.go:135
	// _ = "end of CoverTab[11932]"
//line /usr/local/go/src/internal/intern/intern.go:135
	_go_fuzz_dep_.CoverTab[11933]++
							delete(valMap, keyFor(v.cmpVal))
//line /usr/local/go/src/internal/intern/intern.go:136
	// _ = "end of CoverTab[11933]"
}

//line /usr/local/go/src/internal/intern/intern.go:137
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/intern/intern.go:137
var _ = _go_fuzz_dep_.CoverTab
