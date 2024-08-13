// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/internal/godebug/godebug.go:23
package godebug

//line /usr/local/go/src/internal/godebug/godebug.go:23
import (
//line /usr/local/go/src/internal/godebug/godebug.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/internal/godebug/godebug.go:23
)
//line /usr/local/go/src/internal/godebug/godebug.go:23
import (
//line /usr/local/go/src/internal/godebug/godebug.go:23
	_atomic_ "sync/atomic"
//line /usr/local/go/src/internal/godebug/godebug.go:23
)

import (
	"sync"
	"sync/atomic"
	_ "unsafe"
)

//line /usr/local/go/src/internal/godebug/godebug.go:32
type Setting struct {
	name	string
	once	sync.Once
	value	*atomic.Pointer[string]
}

//line /usr/local/go/src/internal/godebug/godebug.go:39
func New(name string) *Setting {
//line /usr/local/go/src/internal/godebug/godebug.go:39
	_go_fuzz_dep_.CoverTab[2912]++
								return &Setting{name: name}
//line /usr/local/go/src/internal/godebug/godebug.go:40
	// _ = "end of CoverTab[2912]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:44
func (s *Setting) Name() string {
//line /usr/local/go/src/internal/godebug/godebug.go:44
	_go_fuzz_dep_.CoverTab[2913]++
								return s.name
//line /usr/local/go/src/internal/godebug/godebug.go:45
	// _ = "end of CoverTab[2913]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:49
func (s *Setting) String() string {
//line /usr/local/go/src/internal/godebug/godebug.go:49
	_go_fuzz_dep_.CoverTab[2914]++
								return s.name + "=" + s.Value()
//line /usr/local/go/src/internal/godebug/godebug.go:50
	// _ = "end of CoverTab[2914]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:66
var cache sync.Map

var empty string

//line /usr/local/go/src/internal/godebug/godebug.go:77
func (s *Setting) Value() string {
//line /usr/local/go/src/internal/godebug/godebug.go:77
	_go_fuzz_dep_.CoverTab[2915]++
								s.once.Do(func() {
//line /usr/local/go/src/internal/godebug/godebug.go:78
		_go_fuzz_dep_.CoverTab[2917]++
									v, ok := cache.Load(s.name)
									if !ok {
//line /usr/local/go/src/internal/godebug/godebug.go:80
			_go_fuzz_dep_.CoverTab[2919]++
										p := new(atomic.Pointer[string])
										p.Store(&empty)
										v, _ = cache.LoadOrStore(s.name, p)
//line /usr/local/go/src/internal/godebug/godebug.go:83
			// _ = "end of CoverTab[2919]"
		} else {
//line /usr/local/go/src/internal/godebug/godebug.go:84
			_go_fuzz_dep_.CoverTab[2920]++
//line /usr/local/go/src/internal/godebug/godebug.go:84
			// _ = "end of CoverTab[2920]"
//line /usr/local/go/src/internal/godebug/godebug.go:84
		}
//line /usr/local/go/src/internal/godebug/godebug.go:84
		// _ = "end of CoverTab[2917]"
//line /usr/local/go/src/internal/godebug/godebug.go:84
		_go_fuzz_dep_.CoverTab[2918]++
									s.value = v.(*atomic.Pointer[string])
//line /usr/local/go/src/internal/godebug/godebug.go:85
		// _ = "end of CoverTab[2918]"
	})
//line /usr/local/go/src/internal/godebug/godebug.go:86
	// _ = "end of CoverTab[2915]"
//line /usr/local/go/src/internal/godebug/godebug.go:86
	_go_fuzz_dep_.CoverTab[2916]++
								return *s.value.Load()
//line /usr/local/go/src/internal/godebug/godebug.go:87
	// _ = "end of CoverTab[2916]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:97
//go:linkname setUpdate
func setUpdate(update func(string, string))

func init() {
	setUpdate(update)
}

var updateMu sync.Mutex

//line /usr/local/go/src/internal/godebug/godebug.go:109
func update(def, env string) {
//line /usr/local/go/src/internal/godebug/godebug.go:109
	_go_fuzz_dep_.CoverTab[2921]++
								updateMu.Lock()
								defer updateMu.Unlock()

//line /usr/local/go/src/internal/godebug/godebug.go:117
	did := make(map[string]bool)
								parse(did, env)
								parse(did, def)

//line /usr/local/go/src/internal/godebug/godebug.go:122
	cache.Range(func(name, v any) bool {
//line /usr/local/go/src/internal/godebug/godebug.go:122
		_go_fuzz_dep_.CoverTab[2922]++
									if !did[name.(string)] {
//line /usr/local/go/src/internal/godebug/godebug.go:123
			_go_fuzz_dep_.CoverTab[2924]++
										v.(*atomic.Pointer[string]).Store(&empty)
//line /usr/local/go/src/internal/godebug/godebug.go:124
			// _ = "end of CoverTab[2924]"
		} else {
//line /usr/local/go/src/internal/godebug/godebug.go:125
			_go_fuzz_dep_.CoverTab[2925]++
//line /usr/local/go/src/internal/godebug/godebug.go:125
			// _ = "end of CoverTab[2925]"
//line /usr/local/go/src/internal/godebug/godebug.go:125
		}
//line /usr/local/go/src/internal/godebug/godebug.go:125
		// _ = "end of CoverTab[2922]"
//line /usr/local/go/src/internal/godebug/godebug.go:125
		_go_fuzz_dep_.CoverTab[2923]++
									return true
//line /usr/local/go/src/internal/godebug/godebug.go:126
		// _ = "end of CoverTab[2923]"
	})
//line /usr/local/go/src/internal/godebug/godebug.go:127
	// _ = "end of CoverTab[2921]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:135
func parse(did map[string]bool, s string) {
//line /usr/local/go/src/internal/godebug/godebug.go:135
	_go_fuzz_dep_.CoverTab[2926]++

//line /usr/local/go/src/internal/godebug/godebug.go:141
	end := len(s)
	eq := -1
	for i := end - 1; i >= -1; i-- {
//line /usr/local/go/src/internal/godebug/godebug.go:143
		_go_fuzz_dep_.CoverTab[2927]++
									if i == -1 || func() bool {
//line /usr/local/go/src/internal/godebug/godebug.go:144
			_go_fuzz_dep_.CoverTab[2928]++
//line /usr/local/go/src/internal/godebug/godebug.go:144
			return s[i] == ','
//line /usr/local/go/src/internal/godebug/godebug.go:144
			// _ = "end of CoverTab[2928]"
//line /usr/local/go/src/internal/godebug/godebug.go:144
		}() {
//line /usr/local/go/src/internal/godebug/godebug.go:144
			_go_fuzz_dep_.CoverTab[2929]++
										if eq >= 0 {
//line /usr/local/go/src/internal/godebug/godebug.go:145
				_go_fuzz_dep_.CoverTab[2931]++
											name, value := s[i+1:eq], s[eq+1:end]
											if !did[name] {
//line /usr/local/go/src/internal/godebug/godebug.go:147
					_go_fuzz_dep_.CoverTab[2932]++
												did[name] = true
												v, ok := cache.Load(name)
												if !ok {
//line /usr/local/go/src/internal/godebug/godebug.go:150
						_go_fuzz_dep_.CoverTab[2934]++
													p := new(atomic.Pointer[string])
													p.Store(&empty)
													v, _ = cache.LoadOrStore(name, p)
//line /usr/local/go/src/internal/godebug/godebug.go:153
						// _ = "end of CoverTab[2934]"
					} else {
//line /usr/local/go/src/internal/godebug/godebug.go:154
						_go_fuzz_dep_.CoverTab[2935]++
//line /usr/local/go/src/internal/godebug/godebug.go:154
						// _ = "end of CoverTab[2935]"
//line /usr/local/go/src/internal/godebug/godebug.go:154
					}
//line /usr/local/go/src/internal/godebug/godebug.go:154
					// _ = "end of CoverTab[2932]"
//line /usr/local/go/src/internal/godebug/godebug.go:154
					_go_fuzz_dep_.CoverTab[2933]++
												v.(*atomic.Pointer[string]).Store(&value)
//line /usr/local/go/src/internal/godebug/godebug.go:155
					// _ = "end of CoverTab[2933]"
				} else {
//line /usr/local/go/src/internal/godebug/godebug.go:156
					_go_fuzz_dep_.CoverTab[2936]++
//line /usr/local/go/src/internal/godebug/godebug.go:156
					// _ = "end of CoverTab[2936]"
//line /usr/local/go/src/internal/godebug/godebug.go:156
				}
//line /usr/local/go/src/internal/godebug/godebug.go:156
				// _ = "end of CoverTab[2931]"
			} else {
//line /usr/local/go/src/internal/godebug/godebug.go:157
				_go_fuzz_dep_.CoverTab[2937]++
//line /usr/local/go/src/internal/godebug/godebug.go:157
				// _ = "end of CoverTab[2937]"
//line /usr/local/go/src/internal/godebug/godebug.go:157
			}
//line /usr/local/go/src/internal/godebug/godebug.go:157
			// _ = "end of CoverTab[2929]"
//line /usr/local/go/src/internal/godebug/godebug.go:157
			_go_fuzz_dep_.CoverTab[2930]++
										eq = -1
										end = i
//line /usr/local/go/src/internal/godebug/godebug.go:159
			// _ = "end of CoverTab[2930]"
		} else {
//line /usr/local/go/src/internal/godebug/godebug.go:160
			_go_fuzz_dep_.CoverTab[2938]++
//line /usr/local/go/src/internal/godebug/godebug.go:160
			if s[i] == '=' {
//line /usr/local/go/src/internal/godebug/godebug.go:160
				_go_fuzz_dep_.CoverTab[2939]++
											eq = i
//line /usr/local/go/src/internal/godebug/godebug.go:161
				// _ = "end of CoverTab[2939]"
			} else {
//line /usr/local/go/src/internal/godebug/godebug.go:162
				_go_fuzz_dep_.CoverTab[2940]++
//line /usr/local/go/src/internal/godebug/godebug.go:162
				// _ = "end of CoverTab[2940]"
//line /usr/local/go/src/internal/godebug/godebug.go:162
			}
//line /usr/local/go/src/internal/godebug/godebug.go:162
			// _ = "end of CoverTab[2938]"
//line /usr/local/go/src/internal/godebug/godebug.go:162
		}
//line /usr/local/go/src/internal/godebug/godebug.go:162
		// _ = "end of CoverTab[2927]"
	}
//line /usr/local/go/src/internal/godebug/godebug.go:163
	// _ = "end of CoverTab[2926]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:164
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/godebug/godebug.go:164
var _ = _go_fuzz_dep_.CoverTab
