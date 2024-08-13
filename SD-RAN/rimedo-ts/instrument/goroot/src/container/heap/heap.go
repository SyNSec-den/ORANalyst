// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/container/heap/heap.go:5
// Package heap provides heap operations for any type that implements
//line /usr/local/go/src/container/heap/heap.go:5
// heap.Interface. A heap is a tree with the property that each node is the
//line /usr/local/go/src/container/heap/heap.go:5
// minimum-valued node in its subtree.
//line /usr/local/go/src/container/heap/heap.go:5
//
//line /usr/local/go/src/container/heap/heap.go:5
// The minimum element in the tree is the root, at index 0.
//line /usr/local/go/src/container/heap/heap.go:5
//
//line /usr/local/go/src/container/heap/heap.go:5
// A heap is a common way to implement a priority queue. To build a priority
//line /usr/local/go/src/container/heap/heap.go:5
// queue, implement the Heap interface with the (negative) priority as the
//line /usr/local/go/src/container/heap/heap.go:5
// ordering for the Less method, so Push adds items while Pop removes the
//line /usr/local/go/src/container/heap/heap.go:5
// highest-priority item from the queue. The Examples include such an
//line /usr/local/go/src/container/heap/heap.go:5
// implementation; the file example_pq_test.go has the complete source.
//line /usr/local/go/src/container/heap/heap.go:16
package heap

//line /usr/local/go/src/container/heap/heap.go:16
import (
//line /usr/local/go/src/container/heap/heap.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/container/heap/heap.go:16
)
//line /usr/local/go/src/container/heap/heap.go:16
import (
//line /usr/local/go/src/container/heap/heap.go:16
	_atomic_ "sync/atomic"
//line /usr/local/go/src/container/heap/heap.go:16
)

import "sort"

// The Interface type describes the requirements
//line /usr/local/go/src/container/heap/heap.go:20
// for a type using the routines in this package.
//line /usr/local/go/src/container/heap/heap.go:20
// Any type that implements it may be used as a
//line /usr/local/go/src/container/heap/heap.go:20
// min-heap with the following invariants (established after
//line /usr/local/go/src/container/heap/heap.go:20
// Init has been called or if the data is empty or sorted):
//line /usr/local/go/src/container/heap/heap.go:20
//
//line /usr/local/go/src/container/heap/heap.go:20
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//line /usr/local/go/src/container/heap/heap.go:20
//
//line /usr/local/go/src/container/heap/heap.go:20
// Note that Push and Pop in this interface are for package heap's
//line /usr/local/go/src/container/heap/heap.go:20
// implementation to call. To add and remove things from the heap,
//line /usr/local/go/src/container/heap/heap.go:20
// use heap.Push and heap.Pop.
//line /usr/local/go/src/container/heap/heap.go:31
type Interface interface {
	sort.Interface
	Push(x any)	// add x as element Len()
	Pop() any	// remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
//line /usr/local/go/src/container/heap/heap.go:37
// Init is idempotent with respect to the heap invariants
//line /usr/local/go/src/container/heap/heap.go:37
// and may be called whenever the heap invariants may have been invalidated.
//line /usr/local/go/src/container/heap/heap.go:37
// The complexity is O(n) where n = h.Len().
//line /usr/local/go/src/container/heap/heap.go:41
func Init(h Interface) {
//line /usr/local/go/src/container/heap/heap.go:41
	_go_fuzz_dep_.CoverTab[81497]++

							n := h.Len()
							for i := n/2 - 1; i >= 0; i-- {
//line /usr/local/go/src/container/heap/heap.go:44
		_go_fuzz_dep_.CoverTab[81498]++
								down(h, i, n)
//line /usr/local/go/src/container/heap/heap.go:45
		// _ = "end of CoverTab[81498]"
	}
//line /usr/local/go/src/container/heap/heap.go:46
	// _ = "end of CoverTab[81497]"
}

// Push pushes the element x onto the heap.
//line /usr/local/go/src/container/heap/heap.go:49
// The complexity is O(log n) where n = h.Len().
//line /usr/local/go/src/container/heap/heap.go:51
func Push(h Interface, x any) {
//line /usr/local/go/src/container/heap/heap.go:51
	_go_fuzz_dep_.CoverTab[81499]++
							h.Push(x)
							up(h, h.Len()-1)
//line /usr/local/go/src/container/heap/heap.go:53
	// _ = "end of CoverTab[81499]"
}

// Pop removes and returns the minimum element (according to Less) from the heap.
//line /usr/local/go/src/container/heap/heap.go:56
// The complexity is O(log n) where n = h.Len().
//line /usr/local/go/src/container/heap/heap.go:56
// Pop is equivalent to Remove(h, 0).
//line /usr/local/go/src/container/heap/heap.go:59
func Pop(h Interface) any {
//line /usr/local/go/src/container/heap/heap.go:59
	_go_fuzz_dep_.CoverTab[81500]++
							n := h.Len() - 1
							h.Swap(0, n)
							down(h, 0, n)
							return h.Pop()
//line /usr/local/go/src/container/heap/heap.go:63
	// _ = "end of CoverTab[81500]"
}

// Remove removes and returns the element at index i from the heap.
//line /usr/local/go/src/container/heap/heap.go:66
// The complexity is O(log n) where n = h.Len().
//line /usr/local/go/src/container/heap/heap.go:68
func Remove(h Interface, i int) any {
//line /usr/local/go/src/container/heap/heap.go:68
	_go_fuzz_dep_.CoverTab[81501]++
							n := h.Len() - 1
							if n != i {
//line /usr/local/go/src/container/heap/heap.go:70
		_go_fuzz_dep_.CoverTab[81503]++
								h.Swap(i, n)
								if !down(h, i, n) {
//line /usr/local/go/src/container/heap/heap.go:72
			_go_fuzz_dep_.CoverTab[81504]++
									up(h, i)
//line /usr/local/go/src/container/heap/heap.go:73
			// _ = "end of CoverTab[81504]"
		} else {
//line /usr/local/go/src/container/heap/heap.go:74
			_go_fuzz_dep_.CoverTab[81505]++
//line /usr/local/go/src/container/heap/heap.go:74
			// _ = "end of CoverTab[81505]"
//line /usr/local/go/src/container/heap/heap.go:74
		}
//line /usr/local/go/src/container/heap/heap.go:74
		// _ = "end of CoverTab[81503]"
	} else {
//line /usr/local/go/src/container/heap/heap.go:75
		_go_fuzz_dep_.CoverTab[81506]++
//line /usr/local/go/src/container/heap/heap.go:75
		// _ = "end of CoverTab[81506]"
//line /usr/local/go/src/container/heap/heap.go:75
	}
//line /usr/local/go/src/container/heap/heap.go:75
	// _ = "end of CoverTab[81501]"
//line /usr/local/go/src/container/heap/heap.go:75
	_go_fuzz_dep_.CoverTab[81502]++
							return h.Pop()
//line /usr/local/go/src/container/heap/heap.go:76
	// _ = "end of CoverTab[81502]"
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
//line /usr/local/go/src/container/heap/heap.go:79
// Changing the value of the element at index i and then calling Fix is equivalent to,
//line /usr/local/go/src/container/heap/heap.go:79
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
//line /usr/local/go/src/container/heap/heap.go:79
// The complexity is O(log n) where n = h.Len().
//line /usr/local/go/src/container/heap/heap.go:83
func Fix(h Interface, i int) {
//line /usr/local/go/src/container/heap/heap.go:83
	_go_fuzz_dep_.CoverTab[81507]++
							if !down(h, i, h.Len()) {
//line /usr/local/go/src/container/heap/heap.go:84
		_go_fuzz_dep_.CoverTab[81508]++
								up(h, i)
//line /usr/local/go/src/container/heap/heap.go:85
		// _ = "end of CoverTab[81508]"
	} else {
//line /usr/local/go/src/container/heap/heap.go:86
		_go_fuzz_dep_.CoverTab[81509]++
//line /usr/local/go/src/container/heap/heap.go:86
		// _ = "end of CoverTab[81509]"
//line /usr/local/go/src/container/heap/heap.go:86
	}
//line /usr/local/go/src/container/heap/heap.go:86
	// _ = "end of CoverTab[81507]"
}

func up(h Interface, j int) {
//line /usr/local/go/src/container/heap/heap.go:89
	_go_fuzz_dep_.CoverTab[81510]++
							for {
//line /usr/local/go/src/container/heap/heap.go:90
		_go_fuzz_dep_.CoverTab[81511]++
								i := (j - 1) / 2
								if i == j || func() bool {
//line /usr/local/go/src/container/heap/heap.go:92
			_go_fuzz_dep_.CoverTab[81513]++
//line /usr/local/go/src/container/heap/heap.go:92
			return !h.Less(j, i)
//line /usr/local/go/src/container/heap/heap.go:92
			// _ = "end of CoverTab[81513]"
//line /usr/local/go/src/container/heap/heap.go:92
		}() {
//line /usr/local/go/src/container/heap/heap.go:92
			_go_fuzz_dep_.CoverTab[81514]++
									break
//line /usr/local/go/src/container/heap/heap.go:93
			// _ = "end of CoverTab[81514]"
		} else {
//line /usr/local/go/src/container/heap/heap.go:94
			_go_fuzz_dep_.CoverTab[81515]++
//line /usr/local/go/src/container/heap/heap.go:94
			// _ = "end of CoverTab[81515]"
//line /usr/local/go/src/container/heap/heap.go:94
		}
//line /usr/local/go/src/container/heap/heap.go:94
		// _ = "end of CoverTab[81511]"
//line /usr/local/go/src/container/heap/heap.go:94
		_go_fuzz_dep_.CoverTab[81512]++
								h.Swap(i, j)
								j = i
//line /usr/local/go/src/container/heap/heap.go:96
		// _ = "end of CoverTab[81512]"
	}
//line /usr/local/go/src/container/heap/heap.go:97
	// _ = "end of CoverTab[81510]"
}

func down(h Interface, i0, n int) bool {
//line /usr/local/go/src/container/heap/heap.go:100
	_go_fuzz_dep_.CoverTab[81516]++
							i := i0
							for {
//line /usr/local/go/src/container/heap/heap.go:102
		_go_fuzz_dep_.CoverTab[81518]++
								j1 := 2*i + 1
								if j1 >= n || func() bool {
//line /usr/local/go/src/container/heap/heap.go:104
			_go_fuzz_dep_.CoverTab[81522]++
//line /usr/local/go/src/container/heap/heap.go:104
			return j1 < 0
//line /usr/local/go/src/container/heap/heap.go:104
			// _ = "end of CoverTab[81522]"
//line /usr/local/go/src/container/heap/heap.go:104
		}() {
//line /usr/local/go/src/container/heap/heap.go:104
			_go_fuzz_dep_.CoverTab[81523]++
									break
//line /usr/local/go/src/container/heap/heap.go:105
			// _ = "end of CoverTab[81523]"
		} else {
//line /usr/local/go/src/container/heap/heap.go:106
			_go_fuzz_dep_.CoverTab[81524]++
//line /usr/local/go/src/container/heap/heap.go:106
			// _ = "end of CoverTab[81524]"
//line /usr/local/go/src/container/heap/heap.go:106
		}
//line /usr/local/go/src/container/heap/heap.go:106
		// _ = "end of CoverTab[81518]"
//line /usr/local/go/src/container/heap/heap.go:106
		_go_fuzz_dep_.CoverTab[81519]++
								j := j1
								if j2 := j1 + 1; j2 < n && func() bool {
//line /usr/local/go/src/container/heap/heap.go:108
			_go_fuzz_dep_.CoverTab[81525]++
//line /usr/local/go/src/container/heap/heap.go:108
			return h.Less(j2, j1)
//line /usr/local/go/src/container/heap/heap.go:108
			// _ = "end of CoverTab[81525]"
//line /usr/local/go/src/container/heap/heap.go:108
		}() {
//line /usr/local/go/src/container/heap/heap.go:108
			_go_fuzz_dep_.CoverTab[81526]++
									j = j2
//line /usr/local/go/src/container/heap/heap.go:109
			// _ = "end of CoverTab[81526]"
		} else {
//line /usr/local/go/src/container/heap/heap.go:110
			_go_fuzz_dep_.CoverTab[81527]++
//line /usr/local/go/src/container/heap/heap.go:110
			// _ = "end of CoverTab[81527]"
//line /usr/local/go/src/container/heap/heap.go:110
		}
//line /usr/local/go/src/container/heap/heap.go:110
		// _ = "end of CoverTab[81519]"
//line /usr/local/go/src/container/heap/heap.go:110
		_go_fuzz_dep_.CoverTab[81520]++
								if !h.Less(j, i) {
//line /usr/local/go/src/container/heap/heap.go:111
			_go_fuzz_dep_.CoverTab[81528]++
									break
//line /usr/local/go/src/container/heap/heap.go:112
			// _ = "end of CoverTab[81528]"
		} else {
//line /usr/local/go/src/container/heap/heap.go:113
			_go_fuzz_dep_.CoverTab[81529]++
//line /usr/local/go/src/container/heap/heap.go:113
			// _ = "end of CoverTab[81529]"
//line /usr/local/go/src/container/heap/heap.go:113
		}
//line /usr/local/go/src/container/heap/heap.go:113
		// _ = "end of CoverTab[81520]"
//line /usr/local/go/src/container/heap/heap.go:113
		_go_fuzz_dep_.CoverTab[81521]++
								h.Swap(i, j)
								i = j
//line /usr/local/go/src/container/heap/heap.go:115
		// _ = "end of CoverTab[81521]"
	}
//line /usr/local/go/src/container/heap/heap.go:116
	// _ = "end of CoverTab[81516]"
//line /usr/local/go/src/container/heap/heap.go:116
	_go_fuzz_dep_.CoverTab[81517]++
							return i > i0
//line /usr/local/go/src/container/heap/heap.go:117
	// _ = "end of CoverTab[81517]"
}

//line /usr/local/go/src/container/heap/heap.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/container/heap/heap.go:118
var _ = _go_fuzz_dep_.CoverTab
