// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/container/list/list.go:5
// Package list implements a doubly linked list.
//line /usr/local/go/src/container/list/list.go:5
//
//line /usr/local/go/src/container/list/list.go:5
// To iterate over a list (where l is a *List):
//line /usr/local/go/src/container/list/list.go:5
//
//line /usr/local/go/src/container/list/list.go:5
//	for e := l.Front(); e != nil; e = e.Next() {
//line /usr/local/go/src/container/list/list.go:5
//		// do something with e.Value
//line /usr/local/go/src/container/list/list.go:5
//	}
//line /usr/local/go/src/container/list/list.go:12
package list

//line /usr/local/go/src/container/list/list.go:12
import (
//line /usr/local/go/src/container/list/list.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/container/list/list.go:12
)
//line /usr/local/go/src/container/list/list.go:12
import (
//line /usr/local/go/src/container/list/list.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/container/list/list.go:12
)

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev	*Element

	// The list to which this element belongs.
	list	*List

	// The value stored with this element.
	Value	any
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
//line /usr/local/go/src/container/list/list.go:31
	_go_fuzz_dep_.CoverTab[864]++
							if p := e.next; e.list != nil && func() bool {
//line /usr/local/go/src/container/list/list.go:32
		_go_fuzz_dep_.CoverTab[866]++
//line /usr/local/go/src/container/list/list.go:32
		return p != &e.list.root
//line /usr/local/go/src/container/list/list.go:32
		// _ = "end of CoverTab[866]"
//line /usr/local/go/src/container/list/list.go:32
	}() {
//line /usr/local/go/src/container/list/list.go:32
		_go_fuzz_dep_.CoverTab[867]++
								return p
//line /usr/local/go/src/container/list/list.go:33
		// _ = "end of CoverTab[867]"
	} else {
//line /usr/local/go/src/container/list/list.go:34
		_go_fuzz_dep_.CoverTab[868]++
//line /usr/local/go/src/container/list/list.go:34
		// _ = "end of CoverTab[868]"
//line /usr/local/go/src/container/list/list.go:34
	}
//line /usr/local/go/src/container/list/list.go:34
	// _ = "end of CoverTab[864]"
//line /usr/local/go/src/container/list/list.go:34
	_go_fuzz_dep_.CoverTab[865]++
							return nil
//line /usr/local/go/src/container/list/list.go:35
	// _ = "end of CoverTab[865]"
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
//line /usr/local/go/src/container/list/list.go:39
	_go_fuzz_dep_.CoverTab[869]++
							if p := e.prev; e.list != nil && func() bool {
//line /usr/local/go/src/container/list/list.go:40
		_go_fuzz_dep_.CoverTab[871]++
//line /usr/local/go/src/container/list/list.go:40
		return p != &e.list.root
//line /usr/local/go/src/container/list/list.go:40
		// _ = "end of CoverTab[871]"
//line /usr/local/go/src/container/list/list.go:40
	}() {
//line /usr/local/go/src/container/list/list.go:40
		_go_fuzz_dep_.CoverTab[872]++
								return p
//line /usr/local/go/src/container/list/list.go:41
		// _ = "end of CoverTab[872]"
	} else {
//line /usr/local/go/src/container/list/list.go:42
		_go_fuzz_dep_.CoverTab[873]++
//line /usr/local/go/src/container/list/list.go:42
		// _ = "end of CoverTab[873]"
//line /usr/local/go/src/container/list/list.go:42
	}
//line /usr/local/go/src/container/list/list.go:42
	// _ = "end of CoverTab[869]"
//line /usr/local/go/src/container/list/list.go:42
	_go_fuzz_dep_.CoverTab[870]++
							return nil
//line /usr/local/go/src/container/list/list.go:43
	// _ = "end of CoverTab[870]"
}

// List represents a doubly linked list.
//line /usr/local/go/src/container/list/list.go:46
// The zero value for List is an empty list ready to use.
//line /usr/local/go/src/container/list/list.go:48
type List struct {
	root	Element	// sentinel list element, only &root, root.prev, and root.next are used
	len	int	// current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) Init() *List {
//line /usr/local/go/src/container/list/list.go:54
	_go_fuzz_dep_.CoverTab[874]++
							l.root.next = &l.root
							l.root.prev = &l.root
							l.len = 0
							return l
//line /usr/local/go/src/container/list/list.go:58
	// _ = "end of CoverTab[874]"
}

// New returns an initialized list.
func New() *List {
//line /usr/local/go/src/container/list/list.go:62
	_go_fuzz_dep_.CoverTab[875]++
//line /usr/local/go/src/container/list/list.go:62
	return new(List).Init()
//line /usr/local/go/src/container/list/list.go:62
	// _ = "end of CoverTab[875]"
//line /usr/local/go/src/container/list/list.go:62
}

// Len returns the number of elements of list l.
//line /usr/local/go/src/container/list/list.go:64
// The complexity is O(1).
//line /usr/local/go/src/container/list/list.go:66
func (l *List) Len() int	{ _go_fuzz_dep_.CoverTab[876]++; return l.len; // _ = "end of CoverTab[876]" }

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
//line /usr/local/go/src/container/list/list.go:69
	_go_fuzz_dep_.CoverTab[877]++
							if l.len == 0 {
//line /usr/local/go/src/container/list/list.go:70
		_go_fuzz_dep_.CoverTab[879]++
								return nil
//line /usr/local/go/src/container/list/list.go:71
		// _ = "end of CoverTab[879]"
	} else {
//line /usr/local/go/src/container/list/list.go:72
		_go_fuzz_dep_.CoverTab[880]++
//line /usr/local/go/src/container/list/list.go:72
		// _ = "end of CoverTab[880]"
//line /usr/local/go/src/container/list/list.go:72
	}
//line /usr/local/go/src/container/list/list.go:72
	// _ = "end of CoverTab[877]"
//line /usr/local/go/src/container/list/list.go:72
	_go_fuzz_dep_.CoverTab[878]++
							return l.root.next
//line /usr/local/go/src/container/list/list.go:73
	// _ = "end of CoverTab[878]"
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
//line /usr/local/go/src/container/list/list.go:77
	_go_fuzz_dep_.CoverTab[881]++
							if l.len == 0 {
//line /usr/local/go/src/container/list/list.go:78
		_go_fuzz_dep_.CoverTab[883]++
								return nil
//line /usr/local/go/src/container/list/list.go:79
		// _ = "end of CoverTab[883]"
	} else {
//line /usr/local/go/src/container/list/list.go:80
		_go_fuzz_dep_.CoverTab[884]++
//line /usr/local/go/src/container/list/list.go:80
		// _ = "end of CoverTab[884]"
//line /usr/local/go/src/container/list/list.go:80
	}
//line /usr/local/go/src/container/list/list.go:80
	// _ = "end of CoverTab[881]"
//line /usr/local/go/src/container/list/list.go:80
	_go_fuzz_dep_.CoverTab[882]++
							return l.root.prev
//line /usr/local/go/src/container/list/list.go:81
	// _ = "end of CoverTab[882]"
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
//line /usr/local/go/src/container/list/list.go:85
	_go_fuzz_dep_.CoverTab[885]++
							if l.root.next == nil {
//line /usr/local/go/src/container/list/list.go:86
		_go_fuzz_dep_.CoverTab[886]++
								l.Init()
//line /usr/local/go/src/container/list/list.go:87
		// _ = "end of CoverTab[886]"
	} else {
//line /usr/local/go/src/container/list/list.go:88
		_go_fuzz_dep_.CoverTab[887]++
//line /usr/local/go/src/container/list/list.go:88
		// _ = "end of CoverTab[887]"
//line /usr/local/go/src/container/list/list.go:88
	}
//line /usr/local/go/src/container/list/list.go:88
	// _ = "end of CoverTab[885]"
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
//line /usr/local/go/src/container/list/list.go:92
	_go_fuzz_dep_.CoverTab[888]++
							e.prev = at
							e.next = at.next
							e.prev.next = e
							e.next.prev = e
							e.list = l
							l.len++
							return e
//line /usr/local/go/src/container/list/list.go:99
	// _ = "end of CoverTab[888]"
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v any, at *Element) *Element {
//line /usr/local/go/src/container/list/list.go:103
	_go_fuzz_dep_.CoverTab[889]++
							return l.insert(&Element{Value: v}, at)
//line /usr/local/go/src/container/list/list.go:104
	// _ = "end of CoverTab[889]"
}

// remove removes e from its list, decrements l.len
func (l *List) remove(e *Element) {
//line /usr/local/go/src/container/list/list.go:108
	_go_fuzz_dep_.CoverTab[890]++
							e.prev.next = e.next
							e.next.prev = e.prev
							e.next = nil
							e.prev = nil
							e.list = nil
							l.len--
//line /usr/local/go/src/container/list/list.go:114
	// _ = "end of CoverTab[890]"
}

// move moves e to next to at.
func (l *List) move(e, at *Element) {
//line /usr/local/go/src/container/list/list.go:118
	_go_fuzz_dep_.CoverTab[891]++
							if e == at {
//line /usr/local/go/src/container/list/list.go:119
		_go_fuzz_dep_.CoverTab[893]++
								return
//line /usr/local/go/src/container/list/list.go:120
		// _ = "end of CoverTab[893]"
	} else {
//line /usr/local/go/src/container/list/list.go:121
		_go_fuzz_dep_.CoverTab[894]++
//line /usr/local/go/src/container/list/list.go:121
		// _ = "end of CoverTab[894]"
//line /usr/local/go/src/container/list/list.go:121
	}
//line /usr/local/go/src/container/list/list.go:121
	// _ = "end of CoverTab[891]"
//line /usr/local/go/src/container/list/list.go:121
	_go_fuzz_dep_.CoverTab[892]++
							e.prev.next = e.next
							e.next.prev = e.prev

							e.prev = at
							e.next = at.next
							e.prev.next = e
							e.next.prev = e
//line /usr/local/go/src/container/list/list.go:128
	// _ = "end of CoverTab[892]"
}

// Remove removes e from l if e is an element of list l.
//line /usr/local/go/src/container/list/list.go:131
// It returns the element value e.Value.
//line /usr/local/go/src/container/list/list.go:131
// The element must not be nil.
//line /usr/local/go/src/container/list/list.go:134
func (l *List) Remove(e *Element) any {
//line /usr/local/go/src/container/list/list.go:134
	_go_fuzz_dep_.CoverTab[895]++
							if e.list == l {
//line /usr/local/go/src/container/list/list.go:135
		_go_fuzz_dep_.CoverTab[897]++

//line /usr/local/go/src/container/list/list.go:138
		l.remove(e)
//line /usr/local/go/src/container/list/list.go:138
		// _ = "end of CoverTab[897]"
	} else {
//line /usr/local/go/src/container/list/list.go:139
		_go_fuzz_dep_.CoverTab[898]++
//line /usr/local/go/src/container/list/list.go:139
		// _ = "end of CoverTab[898]"
//line /usr/local/go/src/container/list/list.go:139
	}
//line /usr/local/go/src/container/list/list.go:139
	// _ = "end of CoverTab[895]"
//line /usr/local/go/src/container/list/list.go:139
	_go_fuzz_dep_.CoverTab[896]++
							return e.Value
//line /usr/local/go/src/container/list/list.go:140
	// _ = "end of CoverTab[896]"
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v any) *Element {
//line /usr/local/go/src/container/list/list.go:144
	_go_fuzz_dep_.CoverTab[899]++
							l.lazyInit()
							return l.insertValue(v, &l.root)
//line /usr/local/go/src/container/list/list.go:146
	// _ = "end of CoverTab[899]"
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v any) *Element {
//line /usr/local/go/src/container/list/list.go:150
	_go_fuzz_dep_.CoverTab[900]++
							l.lazyInit()
							return l.insertValue(v, l.root.prev)
//line /usr/local/go/src/container/list/list.go:152
	// _ = "end of CoverTab[900]"
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
//line /usr/local/go/src/container/list/list.go:155
// If mark is not an element of l, the list is not modified.
//line /usr/local/go/src/container/list/list.go:155
// The mark must not be nil.
//line /usr/local/go/src/container/list/list.go:158
func (l *List) InsertBefore(v any, mark *Element) *Element {
//line /usr/local/go/src/container/list/list.go:158
	_go_fuzz_dep_.CoverTab[901]++
							if mark.list != l {
//line /usr/local/go/src/container/list/list.go:159
		_go_fuzz_dep_.CoverTab[903]++
								return nil
//line /usr/local/go/src/container/list/list.go:160
		// _ = "end of CoverTab[903]"
	} else {
//line /usr/local/go/src/container/list/list.go:161
		_go_fuzz_dep_.CoverTab[904]++
//line /usr/local/go/src/container/list/list.go:161
		// _ = "end of CoverTab[904]"
//line /usr/local/go/src/container/list/list.go:161
	}
//line /usr/local/go/src/container/list/list.go:161
	// _ = "end of CoverTab[901]"
//line /usr/local/go/src/container/list/list.go:161
	_go_fuzz_dep_.CoverTab[902]++

							return l.insertValue(v, mark.prev)
//line /usr/local/go/src/container/list/list.go:163
	// _ = "end of CoverTab[902]"
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
//line /usr/local/go/src/container/list/list.go:166
// If mark is not an element of l, the list is not modified.
//line /usr/local/go/src/container/list/list.go:166
// The mark must not be nil.
//line /usr/local/go/src/container/list/list.go:169
func (l *List) InsertAfter(v any, mark *Element) *Element {
//line /usr/local/go/src/container/list/list.go:169
	_go_fuzz_dep_.CoverTab[905]++
							if mark.list != l {
//line /usr/local/go/src/container/list/list.go:170
		_go_fuzz_dep_.CoverTab[907]++
								return nil
//line /usr/local/go/src/container/list/list.go:171
		// _ = "end of CoverTab[907]"
	} else {
//line /usr/local/go/src/container/list/list.go:172
		_go_fuzz_dep_.CoverTab[908]++
//line /usr/local/go/src/container/list/list.go:172
		// _ = "end of CoverTab[908]"
//line /usr/local/go/src/container/list/list.go:172
	}
//line /usr/local/go/src/container/list/list.go:172
	// _ = "end of CoverTab[905]"
//line /usr/local/go/src/container/list/list.go:172
	_go_fuzz_dep_.CoverTab[906]++

							return l.insertValue(v, mark)
//line /usr/local/go/src/container/list/list.go:174
	// _ = "end of CoverTab[906]"
}

// MoveToFront moves element e to the front of list l.
//line /usr/local/go/src/container/list/list.go:177
// If e is not an element of l, the list is not modified.
//line /usr/local/go/src/container/list/list.go:177
// The element must not be nil.
//line /usr/local/go/src/container/list/list.go:180
func (l *List) MoveToFront(e *Element) {
//line /usr/local/go/src/container/list/list.go:180
	_go_fuzz_dep_.CoverTab[909]++
							if e.list != l || func() bool {
//line /usr/local/go/src/container/list/list.go:181
		_go_fuzz_dep_.CoverTab[911]++
//line /usr/local/go/src/container/list/list.go:181
		return l.root.next == e
//line /usr/local/go/src/container/list/list.go:181
		// _ = "end of CoverTab[911]"
//line /usr/local/go/src/container/list/list.go:181
	}() {
//line /usr/local/go/src/container/list/list.go:181
		_go_fuzz_dep_.CoverTab[912]++
								return
//line /usr/local/go/src/container/list/list.go:182
		// _ = "end of CoverTab[912]"
	} else {
//line /usr/local/go/src/container/list/list.go:183
		_go_fuzz_dep_.CoverTab[913]++
//line /usr/local/go/src/container/list/list.go:183
		// _ = "end of CoverTab[913]"
//line /usr/local/go/src/container/list/list.go:183
	}
//line /usr/local/go/src/container/list/list.go:183
	// _ = "end of CoverTab[909]"
//line /usr/local/go/src/container/list/list.go:183
	_go_fuzz_dep_.CoverTab[910]++

							l.move(e, &l.root)
//line /usr/local/go/src/container/list/list.go:185
	// _ = "end of CoverTab[910]"
}

// MoveToBack moves element e to the back of list l.
//line /usr/local/go/src/container/list/list.go:188
// If e is not an element of l, the list is not modified.
//line /usr/local/go/src/container/list/list.go:188
// The element must not be nil.
//line /usr/local/go/src/container/list/list.go:191
func (l *List) MoveToBack(e *Element) {
//line /usr/local/go/src/container/list/list.go:191
	_go_fuzz_dep_.CoverTab[914]++
							if e.list != l || func() bool {
//line /usr/local/go/src/container/list/list.go:192
		_go_fuzz_dep_.CoverTab[916]++
//line /usr/local/go/src/container/list/list.go:192
		return l.root.prev == e
//line /usr/local/go/src/container/list/list.go:192
		// _ = "end of CoverTab[916]"
//line /usr/local/go/src/container/list/list.go:192
	}() {
//line /usr/local/go/src/container/list/list.go:192
		_go_fuzz_dep_.CoverTab[917]++
								return
//line /usr/local/go/src/container/list/list.go:193
		// _ = "end of CoverTab[917]"
	} else {
//line /usr/local/go/src/container/list/list.go:194
		_go_fuzz_dep_.CoverTab[918]++
//line /usr/local/go/src/container/list/list.go:194
		// _ = "end of CoverTab[918]"
//line /usr/local/go/src/container/list/list.go:194
	}
//line /usr/local/go/src/container/list/list.go:194
	// _ = "end of CoverTab[914]"
//line /usr/local/go/src/container/list/list.go:194
	_go_fuzz_dep_.CoverTab[915]++

							l.move(e, l.root.prev)
//line /usr/local/go/src/container/list/list.go:196
	// _ = "end of CoverTab[915]"
}

// MoveBefore moves element e to its new position before mark.
//line /usr/local/go/src/container/list/list.go:199
// If e or mark is not an element of l, or e == mark, the list is not modified.
//line /usr/local/go/src/container/list/list.go:199
// The element and mark must not be nil.
//line /usr/local/go/src/container/list/list.go:202
func (l *List) MoveBefore(e, mark *Element) {
//line /usr/local/go/src/container/list/list.go:202
	_go_fuzz_dep_.CoverTab[919]++
							if e.list != l || func() bool {
//line /usr/local/go/src/container/list/list.go:203
		_go_fuzz_dep_.CoverTab[921]++
//line /usr/local/go/src/container/list/list.go:203
		return e == mark
//line /usr/local/go/src/container/list/list.go:203
		// _ = "end of CoverTab[921]"
//line /usr/local/go/src/container/list/list.go:203
	}() || func() bool {
//line /usr/local/go/src/container/list/list.go:203
		_go_fuzz_dep_.CoverTab[922]++
//line /usr/local/go/src/container/list/list.go:203
		return mark.list != l
//line /usr/local/go/src/container/list/list.go:203
		// _ = "end of CoverTab[922]"
//line /usr/local/go/src/container/list/list.go:203
	}() {
//line /usr/local/go/src/container/list/list.go:203
		_go_fuzz_dep_.CoverTab[923]++
								return
//line /usr/local/go/src/container/list/list.go:204
		// _ = "end of CoverTab[923]"
	} else {
//line /usr/local/go/src/container/list/list.go:205
		_go_fuzz_dep_.CoverTab[924]++
//line /usr/local/go/src/container/list/list.go:205
		// _ = "end of CoverTab[924]"
//line /usr/local/go/src/container/list/list.go:205
	}
//line /usr/local/go/src/container/list/list.go:205
	// _ = "end of CoverTab[919]"
//line /usr/local/go/src/container/list/list.go:205
	_go_fuzz_dep_.CoverTab[920]++
							l.move(e, mark.prev)
//line /usr/local/go/src/container/list/list.go:206
	// _ = "end of CoverTab[920]"
}

// MoveAfter moves element e to its new position after mark.
//line /usr/local/go/src/container/list/list.go:209
// If e or mark is not an element of l, or e == mark, the list is not modified.
//line /usr/local/go/src/container/list/list.go:209
// The element and mark must not be nil.
//line /usr/local/go/src/container/list/list.go:212
func (l *List) MoveAfter(e, mark *Element) {
//line /usr/local/go/src/container/list/list.go:212
	_go_fuzz_dep_.CoverTab[925]++
							if e.list != l || func() bool {
//line /usr/local/go/src/container/list/list.go:213
		_go_fuzz_dep_.CoverTab[927]++
//line /usr/local/go/src/container/list/list.go:213
		return e == mark
//line /usr/local/go/src/container/list/list.go:213
		// _ = "end of CoverTab[927]"
//line /usr/local/go/src/container/list/list.go:213
	}() || func() bool {
//line /usr/local/go/src/container/list/list.go:213
		_go_fuzz_dep_.CoverTab[928]++
//line /usr/local/go/src/container/list/list.go:213
		return mark.list != l
//line /usr/local/go/src/container/list/list.go:213
		// _ = "end of CoverTab[928]"
//line /usr/local/go/src/container/list/list.go:213
	}() {
//line /usr/local/go/src/container/list/list.go:213
		_go_fuzz_dep_.CoverTab[929]++
								return
//line /usr/local/go/src/container/list/list.go:214
		// _ = "end of CoverTab[929]"
	} else {
//line /usr/local/go/src/container/list/list.go:215
		_go_fuzz_dep_.CoverTab[930]++
//line /usr/local/go/src/container/list/list.go:215
		// _ = "end of CoverTab[930]"
//line /usr/local/go/src/container/list/list.go:215
	}
//line /usr/local/go/src/container/list/list.go:215
	// _ = "end of CoverTab[925]"
//line /usr/local/go/src/container/list/list.go:215
	_go_fuzz_dep_.CoverTab[926]++
							l.move(e, mark)
//line /usr/local/go/src/container/list/list.go:216
	// _ = "end of CoverTab[926]"
}

// PushBackList inserts a copy of another list at the back of list l.
//line /usr/local/go/src/container/list/list.go:219
// The lists l and other may be the same. They must not be nil.
//line /usr/local/go/src/container/list/list.go:221
func (l *List) PushBackList(other *List) {
//line /usr/local/go/src/container/list/list.go:221
	_go_fuzz_dep_.CoverTab[931]++
							l.lazyInit()
							for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
//line /usr/local/go/src/container/list/list.go:223
		_go_fuzz_dep_.CoverTab[932]++
								l.insertValue(e.Value, l.root.prev)
//line /usr/local/go/src/container/list/list.go:224
		// _ = "end of CoverTab[932]"
	}
//line /usr/local/go/src/container/list/list.go:225
	// _ = "end of CoverTab[931]"
}

// PushFrontList inserts a copy of another list at the front of list l.
//line /usr/local/go/src/container/list/list.go:228
// The lists l and other may be the same. They must not be nil.
//line /usr/local/go/src/container/list/list.go:230
func (l *List) PushFrontList(other *List) {
//line /usr/local/go/src/container/list/list.go:230
	_go_fuzz_dep_.CoverTab[933]++
							l.lazyInit()
							for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
//line /usr/local/go/src/container/list/list.go:232
		_go_fuzz_dep_.CoverTab[934]++
								l.insertValue(e.Value, &l.root)
//line /usr/local/go/src/container/list/list.go:233
		// _ = "end of CoverTab[934]"
	}
//line /usr/local/go/src/container/list/list.go:234
	// _ = "end of CoverTab[933]"
}

//line /usr/local/go/src/container/list/list.go:235
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/container/list/list.go:235
var _ = _go_fuzz_dep_.CoverTab
