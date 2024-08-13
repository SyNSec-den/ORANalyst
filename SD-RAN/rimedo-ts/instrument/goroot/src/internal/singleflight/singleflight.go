// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/internal/singleflight/singleflight.go:5
// Package singleflight provides a duplicate function call suppression
//line /usr/local/go/src/internal/singleflight/singleflight.go:5
// mechanism.
//line /usr/local/go/src/internal/singleflight/singleflight.go:7
package singleflight

//line /usr/local/go/src/internal/singleflight/singleflight.go:7
import (
//line /usr/local/go/src/internal/singleflight/singleflight.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/internal/singleflight/singleflight.go:7
)
//line /usr/local/go/src/internal/singleflight/singleflight.go:7
import (
//line /usr/local/go/src/internal/singleflight/singleflight.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/internal/singleflight/singleflight.go:7
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
//line /usr/local/go/src/internal/singleflight/singleflight.go:27
// which units of work can be executed with duplicate suppression.
//line /usr/local/go/src/internal/singleflight/singleflight.go:29
type Group struct {
	mu	sync.Mutex		// protects m
	m	map[string]*call	// lazily initialized
}

// Result holds the results of Do, so they can be passed
//line /usr/local/go/src/internal/singleflight/singleflight.go:34
// on a channel.
//line /usr/local/go/src/internal/singleflight/singleflight.go:36
type Result struct {
	Val	any
	Err	error
	Shared	bool
}

// Do executes and returns the results of the given function, making
//line /usr/local/go/src/internal/singleflight/singleflight.go:42
// sure that only one execution is in-flight for a given key at a
//line /usr/local/go/src/internal/singleflight/singleflight.go:42
// time. If a duplicate comes in, the duplicate caller waits for the
//line /usr/local/go/src/internal/singleflight/singleflight.go:42
// original to complete and receives the same results.
//line /usr/local/go/src/internal/singleflight/singleflight.go:42
// The return value shared indicates whether v was given to multiple callers.
//line /usr/local/go/src/internal/singleflight/singleflight.go:47
func (g *Group) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
//line /usr/local/go/src/internal/singleflight/singleflight.go:47
	_go_fuzz_dep_.CoverTab[11878]++
									g.mu.Lock()
									if g.m == nil {
//line /usr/local/go/src/internal/singleflight/singleflight.go:49
		_go_fuzz_dep_.CoverTab[11881]++
										g.m = make(map[string]*call)
//line /usr/local/go/src/internal/singleflight/singleflight.go:50
		// _ = "end of CoverTab[11881]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:51
		_go_fuzz_dep_.CoverTab[11882]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:51
		// _ = "end of CoverTab[11882]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:51
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:51
	// _ = "end of CoverTab[11878]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:51
	_go_fuzz_dep_.CoverTab[11879]++
									if c, ok := g.m[key]; ok {
//line /usr/local/go/src/internal/singleflight/singleflight.go:52
		_go_fuzz_dep_.CoverTab[11883]++
										c.dups++
										g.mu.Unlock()
										c.wg.Wait()
										return c.val, c.err, true
//line /usr/local/go/src/internal/singleflight/singleflight.go:56
		// _ = "end of CoverTab[11883]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:57
		_go_fuzz_dep_.CoverTab[11884]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:57
		// _ = "end of CoverTab[11884]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:57
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:57
	// _ = "end of CoverTab[11879]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:57
	_go_fuzz_dep_.CoverTab[11880]++
									c := new(call)
									c.wg.Add(1)
									g.m[key] = c
									g.mu.Unlock()

									g.doCall(c, key, fn)
									return c.val, c.err, c.dups > 0
//line /usr/local/go/src/internal/singleflight/singleflight.go:64
	// _ = "end of CoverTab[11880]"
}

// DoChan is like Do but returns a channel that will receive the
//line /usr/local/go/src/internal/singleflight/singleflight.go:67
// results when they are ready.
//line /usr/local/go/src/internal/singleflight/singleflight.go:69
func (g *Group) DoChan(key string, fn func() (any, error)) <-chan Result {
//line /usr/local/go/src/internal/singleflight/singleflight.go:69
	_go_fuzz_dep_.CoverTab[11885]++
									ch := make(chan Result, 1)
									g.mu.Lock()
									if g.m == nil {
//line /usr/local/go/src/internal/singleflight/singleflight.go:72
		_go_fuzz_dep_.CoverTab[11888]++
										g.m = make(map[string]*call)
//line /usr/local/go/src/internal/singleflight/singleflight.go:73
		// _ = "end of CoverTab[11888]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:74
		_go_fuzz_dep_.CoverTab[11889]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:74
		// _ = "end of CoverTab[11889]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:74
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:74
	// _ = "end of CoverTab[11885]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:74
	_go_fuzz_dep_.CoverTab[11886]++
									if c, ok := g.m[key]; ok {
//line /usr/local/go/src/internal/singleflight/singleflight.go:75
		_go_fuzz_dep_.CoverTab[11890]++
										c.dups++
										c.chans = append(c.chans, ch)
										g.mu.Unlock()
										return ch
//line /usr/local/go/src/internal/singleflight/singleflight.go:79
		// _ = "end of CoverTab[11890]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:80
		_go_fuzz_dep_.CoverTab[11891]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:80
		// _ = "end of CoverTab[11891]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:80
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:80
	// _ = "end of CoverTab[11886]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:80
	_go_fuzz_dep_.CoverTab[11887]++
									c := &call{chans: []chan<- Result{ch}}
									c.wg.Add(1)
									g.m[key] = c
									g.mu.Unlock()
//line /usr/local/go/src/internal/singleflight/singleflight.go:84
	_curRoutineNum1_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/internal/singleflight/singleflight.go:84
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum1_)

									go g.doCall(c, key, fn)

									return ch
//line /usr/local/go/src/internal/singleflight/singleflight.go:88
	// _ = "end of CoverTab[11887]"
}

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func() (any, error)) {
//line /usr/local/go/src/internal/singleflight/singleflight.go:92
	_go_fuzz_dep_.CoverTab[11892]++
									c.val, c.err = fn()

									g.mu.Lock()
									c.wg.Done()
									if g.m[key] == c {
//line /usr/local/go/src/internal/singleflight/singleflight.go:97
		_go_fuzz_dep_.CoverTab[11895]++
										delete(g.m, key)
//line /usr/local/go/src/internal/singleflight/singleflight.go:98
		// _ = "end of CoverTab[11895]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:99
		_go_fuzz_dep_.CoverTab[11896]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:99
		// _ = "end of CoverTab[11896]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:99
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:99
	// _ = "end of CoverTab[11892]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:99
	_go_fuzz_dep_.CoverTab[11893]++
									for _, ch := range c.chans {
//line /usr/local/go/src/internal/singleflight/singleflight.go:100
		_go_fuzz_dep_.CoverTab[11897]++
										ch <- Result{c.val, c.err, c.dups > 0}
//line /usr/local/go/src/internal/singleflight/singleflight.go:101
		// _ = "end of CoverTab[11897]"
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:102
	// _ = "end of CoverTab[11893]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:102
	_go_fuzz_dep_.CoverTab[11894]++
									g.mu.Unlock()
//line /usr/local/go/src/internal/singleflight/singleflight.go:103
	// _ = "end of CoverTab[11894]"
}

// ForgetUnshared tells the singleflight to forget about a key if it is not
//line /usr/local/go/src/internal/singleflight/singleflight.go:106
// shared with any other goroutines. Future calls to Do for a forgotten key
//line /usr/local/go/src/internal/singleflight/singleflight.go:106
// will call the function rather than waiting for an earlier call to complete.
//line /usr/local/go/src/internal/singleflight/singleflight.go:106
// Returns whether the key was forgotten or unknown--that is, whether no
//line /usr/local/go/src/internal/singleflight/singleflight.go:106
// other goroutines are waiting for the result.
//line /usr/local/go/src/internal/singleflight/singleflight.go:111
func (g *Group) ForgetUnshared(key string) bool {
//line /usr/local/go/src/internal/singleflight/singleflight.go:111
	_go_fuzz_dep_.CoverTab[11898]++
									g.mu.Lock()
									defer g.mu.Unlock()
									c, ok := g.m[key]
									if !ok {
//line /usr/local/go/src/internal/singleflight/singleflight.go:115
		_go_fuzz_dep_.CoverTab[11901]++
										return true
//line /usr/local/go/src/internal/singleflight/singleflight.go:116
		// _ = "end of CoverTab[11901]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:117
		_go_fuzz_dep_.CoverTab[11902]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:117
		// _ = "end of CoverTab[11902]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:117
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:117
	// _ = "end of CoverTab[11898]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:117
	_go_fuzz_dep_.CoverTab[11899]++
									if c.dups == 0 {
//line /usr/local/go/src/internal/singleflight/singleflight.go:118
		_go_fuzz_dep_.CoverTab[11903]++
										delete(g.m, key)
										return true
//line /usr/local/go/src/internal/singleflight/singleflight.go:120
		// _ = "end of CoverTab[11903]"
	} else {
//line /usr/local/go/src/internal/singleflight/singleflight.go:121
		_go_fuzz_dep_.CoverTab[11904]++
//line /usr/local/go/src/internal/singleflight/singleflight.go:121
		// _ = "end of CoverTab[11904]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:121
	}
//line /usr/local/go/src/internal/singleflight/singleflight.go:121
	// _ = "end of CoverTab[11899]"
//line /usr/local/go/src/internal/singleflight/singleflight.go:121
	_go_fuzz_dep_.CoverTab[11900]++
									return false
//line /usr/local/go/src/internal/singleflight/singleflight.go:122
	// _ = "end of CoverTab[11900]"
}

//line /usr/local/go/src/internal/singleflight/singleflight.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/singleflight/singleflight.go:123
var _ = _go_fuzz_dep_.CoverTab
