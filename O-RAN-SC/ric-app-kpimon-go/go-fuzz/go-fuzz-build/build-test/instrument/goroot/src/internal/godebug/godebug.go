// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/internal/godebug/godebug.go:30
package godebug

//line /snap/go/10455/src/internal/godebug/godebug.go:30
import (
//line /snap/go/10455/src/internal/godebug/godebug.go:30
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/internal/godebug/godebug.go:30
)
//line /snap/go/10455/src/internal/godebug/godebug.go:30
import (
//line /snap/go/10455/src/internal/godebug/godebug.go:30
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/internal/godebug/godebug.go:30
)

//line /snap/go/10455/src/internal/godebug/godebug.go:36
import (
	"internal/bisect"
	"internal/godebugs"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "unsafe"
)

//line /snap/go/10455/src/internal/godebug/godebug.go:46
type Setting struct {
	name	string
	once	sync.Once
	*setting
}

type setting struct {
	value		atomic.Pointer[value]
	nonDefaultOnce	sync.Once
	nonDefault	atomic.Uint64
	info		*godebugs.Info
}

type value struct {
	text	string
	bisect	*bisect.Matcher
}

//line /snap/go/10455/src/internal/godebug/godebug.go:73
func New(name string) *Setting {
//line /snap/go/10455/src/internal/godebug/godebug.go:73
	_go_fuzz_dep_.CoverTab[3766]++
								return &Setting{name: name}
//line /snap/go/10455/src/internal/godebug/godebug.go:74
	// _ = "end of CoverTab[3766]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:78
func (s *Setting) Name() string {
//line /snap/go/10455/src/internal/godebug/godebug.go:78
	_go_fuzz_dep_.CoverTab[3767]++
								if s.name != "" && func() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:79
		_go_fuzz_dep_.CoverTab[3769]++
//line /snap/go/10455/src/internal/godebug/godebug.go:79
		return s.name[0] == '#'
//line /snap/go/10455/src/internal/godebug/godebug.go:79
		// _ = "end of CoverTab[3769]"
//line /snap/go/10455/src/internal/godebug/godebug.go:79
	}() {
//line /snap/go/10455/src/internal/godebug/godebug.go:79
		_go_fuzz_dep_.CoverTab[526944]++
//line /snap/go/10455/src/internal/godebug/godebug.go:79
		_go_fuzz_dep_.CoverTab[3770]++
									return s.name[1:]
//line /snap/go/10455/src/internal/godebug/godebug.go:80
		// _ = "end of CoverTab[3770]"
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:81
		_go_fuzz_dep_.CoverTab[526945]++
//line /snap/go/10455/src/internal/godebug/godebug.go:81
		_go_fuzz_dep_.CoverTab[3771]++
//line /snap/go/10455/src/internal/godebug/godebug.go:81
		// _ = "end of CoverTab[3771]"
//line /snap/go/10455/src/internal/godebug/godebug.go:81
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:81
	// _ = "end of CoverTab[3767]"
//line /snap/go/10455/src/internal/godebug/godebug.go:81
	_go_fuzz_dep_.CoverTab[3768]++
								return s.name
//line /snap/go/10455/src/internal/godebug/godebug.go:82
	// _ = "end of CoverTab[3768]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:86
func (s *Setting) Undocumented() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:86
	_go_fuzz_dep_.CoverTab[3772]++
								return s.name != "" && func() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:87
		_go_fuzz_dep_.CoverTab[3773]++
//line /snap/go/10455/src/internal/godebug/godebug.go:87
		return s.name[0] == '#'
//line /snap/go/10455/src/internal/godebug/godebug.go:87
		// _ = "end of CoverTab[3773]"
//line /snap/go/10455/src/internal/godebug/godebug.go:87
	}()
//line /snap/go/10455/src/internal/godebug/godebug.go:87
	// _ = "end of CoverTab[3772]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:91
func (s *Setting) String() string {
//line /snap/go/10455/src/internal/godebug/godebug.go:91
	_go_fuzz_dep_.CoverTab[3774]++
								return s.Name() + "=" + s.Value()
//line /snap/go/10455/src/internal/godebug/godebug.go:92
	// _ = "end of CoverTab[3774]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:101
func (s *Setting) IncNonDefault() {
//line /snap/go/10455/src/internal/godebug/godebug.go:101
	_go_fuzz_dep_.CoverTab[3775]++
								s.nonDefaultOnce.Do(s.register)
								s.nonDefault.Add(1)
//line /snap/go/10455/src/internal/godebug/godebug.go:103
	// _ = "end of CoverTab[3775]"
}

func (s *Setting) register() {
//line /snap/go/10455/src/internal/godebug/godebug.go:106
	_go_fuzz_dep_.CoverTab[3776]++
								if s.info == nil || func() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:107
		_go_fuzz_dep_.CoverTab[3778]++
//line /snap/go/10455/src/internal/godebug/godebug.go:107
		return s.info.Opaque
//line /snap/go/10455/src/internal/godebug/godebug.go:107
		// _ = "end of CoverTab[3778]"
//line /snap/go/10455/src/internal/godebug/godebug.go:107
	}() {
//line /snap/go/10455/src/internal/godebug/godebug.go:107
		_go_fuzz_dep_.CoverTab[526946]++
//line /snap/go/10455/src/internal/godebug/godebug.go:107
		_go_fuzz_dep_.CoverTab[3779]++
									panic("godebug: unexpected IncNonDefault of " + s.name)
//line /snap/go/10455/src/internal/godebug/godebug.go:108
		// _ = "end of CoverTab[3779]"
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:109
		_go_fuzz_dep_.CoverTab[526947]++
//line /snap/go/10455/src/internal/godebug/godebug.go:109
		_go_fuzz_dep_.CoverTab[3780]++
//line /snap/go/10455/src/internal/godebug/godebug.go:109
		// _ = "end of CoverTab[3780]"
//line /snap/go/10455/src/internal/godebug/godebug.go:109
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:109
	// _ = "end of CoverTab[3776]"
//line /snap/go/10455/src/internal/godebug/godebug.go:109
	_go_fuzz_dep_.CoverTab[3777]++
								registerMetric("/godebug/non-default-behavior/"+s.Name()+":events", s.nonDefault.Load)
//line /snap/go/10455/src/internal/godebug/godebug.go:110
	// _ = "end of CoverTab[3777]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:126
var cache sync.Map

var empty value

//line /snap/go/10455/src/internal/godebug/godebug.go:137
func (s *Setting) Value() string {
//line /snap/go/10455/src/internal/godebug/godebug.go:137
	_go_fuzz_dep_.CoverTab[3781]++
								s.once.Do(func() {
//line /snap/go/10455/src/internal/godebug/godebug.go:138
		_go_fuzz_dep_.CoverTab[3784]++
									s.setting = lookup(s.Name())
									if s.info == nil && func() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:140
			_go_fuzz_dep_.CoverTab[3785]++
//line /snap/go/10455/src/internal/godebug/godebug.go:140
			return !s.Undocumented()
//line /snap/go/10455/src/internal/godebug/godebug.go:140
			// _ = "end of CoverTab[3785]"
//line /snap/go/10455/src/internal/godebug/godebug.go:140
		}() {
//line /snap/go/10455/src/internal/godebug/godebug.go:140
			_go_fuzz_dep_.CoverTab[526948]++
//line /snap/go/10455/src/internal/godebug/godebug.go:140
			_go_fuzz_dep_.CoverTab[3786]++
										panic("godebug: Value of name not listed in godebugs.All: " + s.name)
//line /snap/go/10455/src/internal/godebug/godebug.go:141
			// _ = "end of CoverTab[3786]"
		} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:142
			_go_fuzz_dep_.CoverTab[526949]++
//line /snap/go/10455/src/internal/godebug/godebug.go:142
			_go_fuzz_dep_.CoverTab[3787]++
//line /snap/go/10455/src/internal/godebug/godebug.go:142
			// _ = "end of CoverTab[3787]"
//line /snap/go/10455/src/internal/godebug/godebug.go:142
		}
//line /snap/go/10455/src/internal/godebug/godebug.go:142
		// _ = "end of CoverTab[3784]"
	})
//line /snap/go/10455/src/internal/godebug/godebug.go:143
	// _ = "end of CoverTab[3781]"
//line /snap/go/10455/src/internal/godebug/godebug.go:143
	_go_fuzz_dep_.CoverTab[3782]++
								v := *s.value.Load()
								if v.bisect != nil && func() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:145
		_go_fuzz_dep_.CoverTab[3788]++
//line /snap/go/10455/src/internal/godebug/godebug.go:145
		return !v.bisect.Stack(&stderr)
//line /snap/go/10455/src/internal/godebug/godebug.go:145
		// _ = "end of CoverTab[3788]"
//line /snap/go/10455/src/internal/godebug/godebug.go:145
	}() {
//line /snap/go/10455/src/internal/godebug/godebug.go:145
		_go_fuzz_dep_.CoverTab[526950]++
//line /snap/go/10455/src/internal/godebug/godebug.go:145
		_go_fuzz_dep_.CoverTab[3789]++
									return ""
//line /snap/go/10455/src/internal/godebug/godebug.go:146
		// _ = "end of CoverTab[3789]"
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:147
		_go_fuzz_dep_.CoverTab[526951]++
//line /snap/go/10455/src/internal/godebug/godebug.go:147
		_go_fuzz_dep_.CoverTab[3790]++
//line /snap/go/10455/src/internal/godebug/godebug.go:147
		// _ = "end of CoverTab[3790]"
//line /snap/go/10455/src/internal/godebug/godebug.go:147
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:147
	// _ = "end of CoverTab[3782]"
//line /snap/go/10455/src/internal/godebug/godebug.go:147
	_go_fuzz_dep_.CoverTab[3783]++
								return v.text
//line /snap/go/10455/src/internal/godebug/godebug.go:148
	// _ = "end of CoverTab[3783]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:152
func lookup(name string) *setting {
//line /snap/go/10455/src/internal/godebug/godebug.go:152
	_go_fuzz_dep_.CoverTab[3791]++
								if v, ok := cache.Load(name); ok {
//line /snap/go/10455/src/internal/godebug/godebug.go:153
		_go_fuzz_dep_.CoverTab[526952]++
//line /snap/go/10455/src/internal/godebug/godebug.go:153
		_go_fuzz_dep_.CoverTab[3794]++
									return v.(*setting)
//line /snap/go/10455/src/internal/godebug/godebug.go:154
		// _ = "end of CoverTab[3794]"
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:155
		_go_fuzz_dep_.CoverTab[526953]++
//line /snap/go/10455/src/internal/godebug/godebug.go:155
		_go_fuzz_dep_.CoverTab[3795]++
//line /snap/go/10455/src/internal/godebug/godebug.go:155
		// _ = "end of CoverTab[3795]"
//line /snap/go/10455/src/internal/godebug/godebug.go:155
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:155
	// _ = "end of CoverTab[3791]"
//line /snap/go/10455/src/internal/godebug/godebug.go:155
	_go_fuzz_dep_.CoverTab[3792]++
								s := new(setting)
								s.info = godebugs.Lookup(name)
								s.value.Store(&empty)
								if v, loaded := cache.LoadOrStore(name, s); loaded {
//line /snap/go/10455/src/internal/godebug/godebug.go:159
		_go_fuzz_dep_.CoverTab[526954]++
//line /snap/go/10455/src/internal/godebug/godebug.go:159
		_go_fuzz_dep_.CoverTab[3796]++

									return v.(*setting)
//line /snap/go/10455/src/internal/godebug/godebug.go:161
		// _ = "end of CoverTab[3796]"
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:162
		_go_fuzz_dep_.CoverTab[526955]++
//line /snap/go/10455/src/internal/godebug/godebug.go:162
		_go_fuzz_dep_.CoverTab[3797]++
//line /snap/go/10455/src/internal/godebug/godebug.go:162
		// _ = "end of CoverTab[3797]"
//line /snap/go/10455/src/internal/godebug/godebug.go:162
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:162
	// _ = "end of CoverTab[3792]"
//line /snap/go/10455/src/internal/godebug/godebug.go:162
	_go_fuzz_dep_.CoverTab[3793]++

								return s
//line /snap/go/10455/src/internal/godebug/godebug.go:164
	// _ = "end of CoverTab[3793]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:174
//go:linkname setUpdate
func setUpdate(update func(string, string))

//line /snap/go/10455/src/internal/godebug/godebug.go:180
//go:linkname registerMetric
func registerMetric(name string, read func() uint64)

//line /snap/go/10455/src/internal/godebug/godebug.go:194
//go:linkname setNewIncNonDefault
func setNewIncNonDefault(newIncNonDefault func(string) func())

func init() {
	setUpdate(update)
	setNewIncNonDefault(newIncNonDefault)
}

func newIncNonDefault(name string) func() {
//line /snap/go/10455/src/internal/godebug/godebug.go:202
	_go_fuzz_dep_.CoverTab[3798]++
								s := New(name)
								s.Value()
								return s.IncNonDefault
//line /snap/go/10455/src/internal/godebug/godebug.go:205
	// _ = "end of CoverTab[3798]"
}

var updateMu sync.Mutex

//line /snap/go/10455/src/internal/godebug/godebug.go:213
func update(def, env string) {
//line /snap/go/10455/src/internal/godebug/godebug.go:213
	_go_fuzz_dep_.CoverTab[3799]++
								updateMu.Lock()
								defer updateMu.Unlock()

//line /snap/go/10455/src/internal/godebug/godebug.go:221
	did := make(map[string]bool)
								parse(did, env)
								parse(did, def)

//line /snap/go/10455/src/internal/godebug/godebug.go:226
	cache.Range(func(name, s any) bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:226
		_go_fuzz_dep_.CoverTab[3800]++
									if !did[name.(string)] {
//line /snap/go/10455/src/internal/godebug/godebug.go:227
			_go_fuzz_dep_.CoverTab[526956]++
//line /snap/go/10455/src/internal/godebug/godebug.go:227
			_go_fuzz_dep_.CoverTab[3802]++
										s.(*setting).value.Store(&empty)
//line /snap/go/10455/src/internal/godebug/godebug.go:228
			// _ = "end of CoverTab[3802]"
		} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:229
			_go_fuzz_dep_.CoverTab[526957]++
//line /snap/go/10455/src/internal/godebug/godebug.go:229
			_go_fuzz_dep_.CoverTab[3803]++
//line /snap/go/10455/src/internal/godebug/godebug.go:229
			// _ = "end of CoverTab[3803]"
//line /snap/go/10455/src/internal/godebug/godebug.go:229
		}
//line /snap/go/10455/src/internal/godebug/godebug.go:229
		// _ = "end of CoverTab[3800]"
//line /snap/go/10455/src/internal/godebug/godebug.go:229
		_go_fuzz_dep_.CoverTab[3801]++
									return true
//line /snap/go/10455/src/internal/godebug/godebug.go:230
		// _ = "end of CoverTab[3801]"
	})
//line /snap/go/10455/src/internal/godebug/godebug.go:231
	// _ = "end of CoverTab[3799]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:242
func parse(did map[string]bool, s string) {
//line /snap/go/10455/src/internal/godebug/godebug.go:242
	_go_fuzz_dep_.CoverTab[3804]++

//line /snap/go/10455/src/internal/godebug/godebug.go:248
	end := len(s)
								eq := -1
//line /snap/go/10455/src/internal/godebug/godebug.go:249
	_go_fuzz_dep_.CoverTab[786621] = 0
								for i := end - 1; i >= -1; i-- {
//line /snap/go/10455/src/internal/godebug/godebug.go:250
		if _go_fuzz_dep_.CoverTab[786621] == 0 {
//line /snap/go/10455/src/internal/godebug/godebug.go:250
			_go_fuzz_dep_.CoverTab[526970]++
//line /snap/go/10455/src/internal/godebug/godebug.go:250
		} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:250
			_go_fuzz_dep_.CoverTab[526971]++
//line /snap/go/10455/src/internal/godebug/godebug.go:250
		}
//line /snap/go/10455/src/internal/godebug/godebug.go:250
		_go_fuzz_dep_.CoverTab[786621] = 1
//line /snap/go/10455/src/internal/godebug/godebug.go:250
		_go_fuzz_dep_.CoverTab[3805]++
									if i == -1 || func() bool {
//line /snap/go/10455/src/internal/godebug/godebug.go:251
			_go_fuzz_dep_.CoverTab[3806]++
//line /snap/go/10455/src/internal/godebug/godebug.go:251
			return s[i] == ','
//line /snap/go/10455/src/internal/godebug/godebug.go:251
			// _ = "end of CoverTab[3806]"
//line /snap/go/10455/src/internal/godebug/godebug.go:251
		}() {
//line /snap/go/10455/src/internal/godebug/godebug.go:251
			_go_fuzz_dep_.CoverTab[526958]++
//line /snap/go/10455/src/internal/godebug/godebug.go:251
			_go_fuzz_dep_.CoverTab[3807]++
										if eq >= 0 {
//line /snap/go/10455/src/internal/godebug/godebug.go:252
				_go_fuzz_dep_.CoverTab[526960]++
//line /snap/go/10455/src/internal/godebug/godebug.go:252
				_go_fuzz_dep_.CoverTab[3809]++
											name, arg := s[i+1:eq], s[eq+1:end]
											if !did[name] {
//line /snap/go/10455/src/internal/godebug/godebug.go:254
					_go_fuzz_dep_.CoverTab[526962]++
//line /snap/go/10455/src/internal/godebug/godebug.go:254
					_go_fuzz_dep_.CoverTab[3810]++
												did[name] = true
												v := &value{text: arg}
//line /snap/go/10455/src/internal/godebug/godebug.go:256
					_go_fuzz_dep_.CoverTab[786622] = 0
												for j := 0; j < len(arg); j++ {
//line /snap/go/10455/src/internal/godebug/godebug.go:257
						if _go_fuzz_dep_.CoverTab[786622] == 0 {
//line /snap/go/10455/src/internal/godebug/godebug.go:257
							_go_fuzz_dep_.CoverTab[526974]++
//line /snap/go/10455/src/internal/godebug/godebug.go:257
						} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:257
							_go_fuzz_dep_.CoverTab[526975]++
//line /snap/go/10455/src/internal/godebug/godebug.go:257
						}
//line /snap/go/10455/src/internal/godebug/godebug.go:257
						_go_fuzz_dep_.CoverTab[786622] = 1
//line /snap/go/10455/src/internal/godebug/godebug.go:257
						_go_fuzz_dep_.CoverTab[3812]++
													if arg[j] == '#' {
//line /snap/go/10455/src/internal/godebug/godebug.go:258
							_go_fuzz_dep_.CoverTab[526964]++
//line /snap/go/10455/src/internal/godebug/godebug.go:258
							_go_fuzz_dep_.CoverTab[3813]++
														v.text = arg[:j]
														v.bisect, _ = bisect.New(arg[j+1:])
														break
//line /snap/go/10455/src/internal/godebug/godebug.go:261
							// _ = "end of CoverTab[3813]"
						} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:262
							_go_fuzz_dep_.CoverTab[526965]++
//line /snap/go/10455/src/internal/godebug/godebug.go:262
							_go_fuzz_dep_.CoverTab[3814]++
//line /snap/go/10455/src/internal/godebug/godebug.go:262
							// _ = "end of CoverTab[3814]"
//line /snap/go/10455/src/internal/godebug/godebug.go:262
						}
//line /snap/go/10455/src/internal/godebug/godebug.go:262
						// _ = "end of CoverTab[3812]"
					}
//line /snap/go/10455/src/internal/godebug/godebug.go:263
					if _go_fuzz_dep_.CoverTab[786622] == 0 {
//line /snap/go/10455/src/internal/godebug/godebug.go:263
						_go_fuzz_dep_.CoverTab[526976]++
//line /snap/go/10455/src/internal/godebug/godebug.go:263
					} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:263
						_go_fuzz_dep_.CoverTab[526977]++
//line /snap/go/10455/src/internal/godebug/godebug.go:263
					}
//line /snap/go/10455/src/internal/godebug/godebug.go:263
					// _ = "end of CoverTab[3810]"
//line /snap/go/10455/src/internal/godebug/godebug.go:263
					_go_fuzz_dep_.CoverTab[3811]++
												lookup(name).value.Store(v)
//line /snap/go/10455/src/internal/godebug/godebug.go:264
					// _ = "end of CoverTab[3811]"
				} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:265
					_go_fuzz_dep_.CoverTab[526963]++
//line /snap/go/10455/src/internal/godebug/godebug.go:265
					_go_fuzz_dep_.CoverTab[3815]++
//line /snap/go/10455/src/internal/godebug/godebug.go:265
					// _ = "end of CoverTab[3815]"
//line /snap/go/10455/src/internal/godebug/godebug.go:265
				}
//line /snap/go/10455/src/internal/godebug/godebug.go:265
				// _ = "end of CoverTab[3809]"
			} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:266
				_go_fuzz_dep_.CoverTab[526961]++
//line /snap/go/10455/src/internal/godebug/godebug.go:266
				_go_fuzz_dep_.CoverTab[3816]++
//line /snap/go/10455/src/internal/godebug/godebug.go:266
				// _ = "end of CoverTab[3816]"
//line /snap/go/10455/src/internal/godebug/godebug.go:266
			}
//line /snap/go/10455/src/internal/godebug/godebug.go:266
			// _ = "end of CoverTab[3807]"
//line /snap/go/10455/src/internal/godebug/godebug.go:266
			_go_fuzz_dep_.CoverTab[3808]++
										eq = -1
										end = i
//line /snap/go/10455/src/internal/godebug/godebug.go:268
			// _ = "end of CoverTab[3808]"
		} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:269
			_go_fuzz_dep_.CoverTab[526959]++
//line /snap/go/10455/src/internal/godebug/godebug.go:269
			_go_fuzz_dep_.CoverTab[3817]++
//line /snap/go/10455/src/internal/godebug/godebug.go:269
			if s[i] == '=' {
//line /snap/go/10455/src/internal/godebug/godebug.go:269
				_go_fuzz_dep_.CoverTab[526966]++
//line /snap/go/10455/src/internal/godebug/godebug.go:269
				_go_fuzz_dep_.CoverTab[3818]++
											eq = i
//line /snap/go/10455/src/internal/godebug/godebug.go:270
				// _ = "end of CoverTab[3818]"
			} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:271
				_go_fuzz_dep_.CoverTab[526967]++
//line /snap/go/10455/src/internal/godebug/godebug.go:271
				_go_fuzz_dep_.CoverTab[3819]++
//line /snap/go/10455/src/internal/godebug/godebug.go:271
				// _ = "end of CoverTab[3819]"
//line /snap/go/10455/src/internal/godebug/godebug.go:271
			}
//line /snap/go/10455/src/internal/godebug/godebug.go:271
			// _ = "end of CoverTab[3817]"
//line /snap/go/10455/src/internal/godebug/godebug.go:271
		}
//line /snap/go/10455/src/internal/godebug/godebug.go:271
		// _ = "end of CoverTab[3805]"
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:272
	if _go_fuzz_dep_.CoverTab[786621] == 0 {
//line /snap/go/10455/src/internal/godebug/godebug.go:272
		_go_fuzz_dep_.CoverTab[526972]++
//line /snap/go/10455/src/internal/godebug/godebug.go:272
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:272
		_go_fuzz_dep_.CoverTab[526973]++
//line /snap/go/10455/src/internal/godebug/godebug.go:272
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:272
	// _ = "end of CoverTab[3804]"
}

type runtimeStderr struct{}

var stderr runtimeStderr

func (*runtimeStderr) Write(b []byte) (int, error) {
//line /snap/go/10455/src/internal/godebug/godebug.go:279
	_go_fuzz_dep_.CoverTab[3820]++
								if len(b) > 0 {
//line /snap/go/10455/src/internal/godebug/godebug.go:280
		_go_fuzz_dep_.CoverTab[526968]++
//line /snap/go/10455/src/internal/godebug/godebug.go:280
		_go_fuzz_dep_.CoverTab[3822]++
									write(2, unsafe.Pointer(&b[0]), int32(len(b)))
//line /snap/go/10455/src/internal/godebug/godebug.go:281
		// _ = "end of CoverTab[3822]"
	} else {
//line /snap/go/10455/src/internal/godebug/godebug.go:282
		_go_fuzz_dep_.CoverTab[526969]++
//line /snap/go/10455/src/internal/godebug/godebug.go:282
		_go_fuzz_dep_.CoverTab[3823]++
//line /snap/go/10455/src/internal/godebug/godebug.go:282
		// _ = "end of CoverTab[3823]"
//line /snap/go/10455/src/internal/godebug/godebug.go:282
	}
//line /snap/go/10455/src/internal/godebug/godebug.go:282
	// _ = "end of CoverTab[3820]"
//line /snap/go/10455/src/internal/godebug/godebug.go:282
	_go_fuzz_dep_.CoverTab[3821]++
								return len(b), nil
//line /snap/go/10455/src/internal/godebug/godebug.go:283
	// _ = "end of CoverTab[3821]"
}

//line /snap/go/10455/src/internal/godebug/godebug.go:289
//go:linkname write runtime.write
func write(fd uintptr, p unsafe.Pointer, n int32) int32

//line /snap/go/10455/src/internal/godebug/godebug.go:290
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/internal/godebug/godebug.go:290
var _ = _go_fuzz_dep_.CoverTab
