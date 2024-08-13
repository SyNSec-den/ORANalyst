// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/internal/singleflight/singleflight.go:5
// Package singleflight provides a duplicate function call suppression
//line /snap/go/10455/src/internal/singleflight/singleflight.go:5
// mechanism.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
package singleflight

//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
import (
//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
)
//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
import (
//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:7
)

import "sync"

// call is an in-flight or completed singleflight.Do call
type call struct {
	wg	sync.WaitGroup

	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val	any
	err	error

	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
	dups	int
	chans	[]chan<- Result
}

// Group represents a class of work and forms a namespace in
//line /snap/go/10455/src/internal/singleflight/singleflight.go:27
// which units of work can be executed with duplicate suppression.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:29
type Group struct {
	mu	sync.Mutex		// protects m
	m	map[string]*call	// lazily initialized
}

// Result holds the results of Do, so they can be passed
//line /snap/go/10455/src/internal/singleflight/singleflight.go:34
// on a channel.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:36
type Result struct {
	Val	any
	Err	error
	Shared	bool
}

// Do executes and returns the results of the given function, making
//line /snap/go/10455/src/internal/singleflight/singleflight.go:42
// sure that only one execution is in-flight for a given key at a
//line /snap/go/10455/src/internal/singleflight/singleflight.go:42
// time. If a duplicate comes in, the duplicate caller waits for the
//line /snap/go/10455/src/internal/singleflight/singleflight.go:42
// original to complete and receives the same results.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:42
// The return value shared indicates whether v was given to multiple callers.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:47
func (g *Group) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:47
	_go_fuzz_dep_.CoverTab[3824]++
									g.mu.Lock()
									if g.m == nil {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:49
		_go_fuzz_dep_.CoverTab[526978]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:49
		_go_fuzz_dep_.CoverTab[3827]++
										g.m = make(map[string]*call)
//line /snap/go/10455/src/internal/singleflight/singleflight.go:50
		// _ = "end of CoverTab[3827]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:51
		_go_fuzz_dep_.CoverTab[526979]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:51
		_go_fuzz_dep_.CoverTab[3828]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:51
		// _ = "end of CoverTab[3828]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:51
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:51
	// _ = "end of CoverTab[3824]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:51
	_go_fuzz_dep_.CoverTab[3825]++
									if c, ok := g.m[key]; ok {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:52
		_go_fuzz_dep_.CoverTab[526980]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:52
		_go_fuzz_dep_.CoverTab[3829]++
										c.dups++
										g.mu.Unlock()
										c.wg.Wait()
										return c.val, c.err, true
//line /snap/go/10455/src/internal/singleflight/singleflight.go:56
		// _ = "end of CoverTab[3829]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:57
		_go_fuzz_dep_.CoverTab[526981]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:57
		_go_fuzz_dep_.CoverTab[3830]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:57
		// _ = "end of CoverTab[3830]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:57
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:57
	// _ = "end of CoverTab[3825]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:57
	_go_fuzz_dep_.CoverTab[3826]++
									c := new(call)
									c.wg.Add(1)
									g.m[key] = c
									g.mu.Unlock()

									g.doCall(c, key, fn)
									return c.val, c.err, c.dups > 0
//line /snap/go/10455/src/internal/singleflight/singleflight.go:64
	// _ = "end of CoverTab[3826]"
}

// DoChan is like Do but returns a channel that will receive the
//line /snap/go/10455/src/internal/singleflight/singleflight.go:67
// results when they are ready.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:69
func (g *Group) DoChan(key string, fn func() (any, error)) <-chan Result {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:69
	_go_fuzz_dep_.CoverTab[3831]++
									ch := make(chan Result, 1)
									g.mu.Lock()
									if g.m == nil {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:72
		_go_fuzz_dep_.CoverTab[526982]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:72
		_go_fuzz_dep_.CoverTab[3834]++
										g.m = make(map[string]*call)
//line /snap/go/10455/src/internal/singleflight/singleflight.go:73
		// _ = "end of CoverTab[3834]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:74
		_go_fuzz_dep_.CoverTab[526983]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:74
		_go_fuzz_dep_.CoverTab[3835]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:74
		// _ = "end of CoverTab[3835]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:74
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:74
	// _ = "end of CoverTab[3831]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:74
	_go_fuzz_dep_.CoverTab[3832]++
									if c, ok := g.m[key]; ok {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:75
		_go_fuzz_dep_.CoverTab[526984]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:75
		_go_fuzz_dep_.CoverTab[3836]++
										c.dups++
										c.chans = append(c.chans, ch)
										g.mu.Unlock()
										return ch
//line /snap/go/10455/src/internal/singleflight/singleflight.go:79
		// _ = "end of CoverTab[3836]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:80
		_go_fuzz_dep_.CoverTab[526985]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:80
		_go_fuzz_dep_.CoverTab[3837]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:80
		// _ = "end of CoverTab[3837]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:80
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:80
	// _ = "end of CoverTab[3832]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:80
	_go_fuzz_dep_.CoverTab[3833]++
									c := &call{chans: []chan<- Result{ch}}
									c.wg.Add(1)
									g.m[key] = c
									g.mu.Unlock()
//line /snap/go/10455/src/internal/singleflight/singleflight.go:84
	_curRoutineNum2_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/internal/singleflight/singleflight.go:84
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum2_)

									go g.doCall(c, key, fn)

									return ch
//line /snap/go/10455/src/internal/singleflight/singleflight.go:88
	// _ = "end of CoverTab[3833]"
}

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func() (any, error)) {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:92
	_go_fuzz_dep_.CoverTab[3838]++
									c.val, c.err = fn()

									g.mu.Lock()
									c.wg.Done()
									if g.m[key] == c {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:97
		_go_fuzz_dep_.CoverTab[526986]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:97
		_go_fuzz_dep_.CoverTab[3841]++
										delete(g.m, key)
//line /snap/go/10455/src/internal/singleflight/singleflight.go:98
		// _ = "end of CoverTab[3841]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
		_go_fuzz_dep_.CoverTab[526987]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
		_go_fuzz_dep_.CoverTab[3842]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
		// _ = "end of CoverTab[3842]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
	// _ = "end of CoverTab[3838]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
	_go_fuzz_dep_.CoverTab[3839]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:99
	_go_fuzz_dep_.CoverTab[786623] = 0
									for _, ch := range c.chans {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
		if _go_fuzz_dep_.CoverTab[786623] == 0 {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
			_go_fuzz_dep_.CoverTab[526992]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
		} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
			_go_fuzz_dep_.CoverTab[526993]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
		}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
		_go_fuzz_dep_.CoverTab[786623] = 1
//line /snap/go/10455/src/internal/singleflight/singleflight.go:100
		_go_fuzz_dep_.CoverTab[3843]++
										ch <- Result{c.val, c.err, c.dups > 0}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:101
		// _ = "end of CoverTab[3843]"
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
	if _go_fuzz_dep_.CoverTab[786623] == 0 {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
		_go_fuzz_dep_.CoverTab[526994]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
		_go_fuzz_dep_.CoverTab[526995]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
	// _ = "end of CoverTab[3839]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:102
	_go_fuzz_dep_.CoverTab[3840]++
									g.mu.Unlock()
//line /snap/go/10455/src/internal/singleflight/singleflight.go:103
	// _ = "end of CoverTab[3840]"
}

// ForgetUnshared tells the singleflight to forget about a key if it is not
//line /snap/go/10455/src/internal/singleflight/singleflight.go:106
// shared with any other goroutines. Future calls to Do for a forgotten key
//line /snap/go/10455/src/internal/singleflight/singleflight.go:106
// will call the function rather than waiting for an earlier call to complete.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:106
// Returns whether the key was forgotten or unknown--that is, whether no
//line /snap/go/10455/src/internal/singleflight/singleflight.go:106
// other goroutines are waiting for the result.
//line /snap/go/10455/src/internal/singleflight/singleflight.go:111
func (g *Group) ForgetUnshared(key string) bool {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:111
	_go_fuzz_dep_.CoverTab[3844]++
									g.mu.Lock()
									defer g.mu.Unlock()
									c, ok := g.m[key]
									if !ok {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:115
		_go_fuzz_dep_.CoverTab[526988]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:115
		_go_fuzz_dep_.CoverTab[3847]++
										return true
//line /snap/go/10455/src/internal/singleflight/singleflight.go:116
		// _ = "end of CoverTab[3847]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:117
		_go_fuzz_dep_.CoverTab[526989]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:117
		_go_fuzz_dep_.CoverTab[3848]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:117
		// _ = "end of CoverTab[3848]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:117
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:117
	// _ = "end of CoverTab[3844]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:117
	_go_fuzz_dep_.CoverTab[3845]++
									if c.dups == 0 {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:118
		_go_fuzz_dep_.CoverTab[526990]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:118
		_go_fuzz_dep_.CoverTab[3849]++
										delete(g.m, key)
										return true
//line /snap/go/10455/src/internal/singleflight/singleflight.go:120
		// _ = "end of CoverTab[3849]"
	} else {
//line /snap/go/10455/src/internal/singleflight/singleflight.go:121
		_go_fuzz_dep_.CoverTab[526991]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:121
		_go_fuzz_dep_.CoverTab[3850]++
//line /snap/go/10455/src/internal/singleflight/singleflight.go:121
		// _ = "end of CoverTab[3850]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:121
	}
//line /snap/go/10455/src/internal/singleflight/singleflight.go:121
	// _ = "end of CoverTab[3845]"
//line /snap/go/10455/src/internal/singleflight/singleflight.go:121
	_go_fuzz_dep_.CoverTab[3846]++
									return false
//line /snap/go/10455/src/internal/singleflight/singleflight.go:122
	// _ = "end of CoverTab[3846]"
}

//line /snap/go/10455/src/internal/singleflight/singleflight.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/internal/singleflight/singleflight.go:123
var _ = _go_fuzz_dep_.CoverTab
