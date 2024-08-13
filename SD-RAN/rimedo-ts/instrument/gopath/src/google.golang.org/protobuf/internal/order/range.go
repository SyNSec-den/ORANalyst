// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:5
// Package order provides ordered access to messages and maps.
package order

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:6
)

import (
	"sort"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type messageField struct {
	fd	protoreflect.FieldDescriptor
	v	protoreflect.Value
}

var messageFieldPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:21
		_go_fuzz_dep_.CoverTab[50149]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:21
		return new([]messageField)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:21
		// _ = "end of CoverTab[50149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:21
	},
}

type (
	// FieldRnger is an interface for visiting all fields in a message.
	// The protoreflect.Message type implements this interface.
	FieldRanger	interface{ Range(VisitField) }
	// VisitField is called every time a message field is visited.
	VisitField	= func(protoreflect.FieldDescriptor, protoreflect.Value) bool
)

// RangeFields iterates over the fields of fs according to the specified order.
func RangeFields(fs FieldRanger, less FieldOrder, fn VisitField) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:33
	_go_fuzz_dep_.CoverTab[50150]++
													if less == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:34
		_go_fuzz_dep_.CoverTab[50155]++
														fs.Range(fn)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:36
		// _ = "end of CoverTab[50155]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:37
		_go_fuzz_dep_.CoverTab[50156]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:37
		// _ = "end of CoverTab[50156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:37
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:37
	// _ = "end of CoverTab[50150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:37
	_go_fuzz_dep_.CoverTab[50151]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:40
	p := messageFieldPool.Get().(*[]messageField)
	fields := (*p)[:0]
	defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:42
		_go_fuzz_dep_.CoverTab[50157]++
														if cap(fields) < 1024 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:43
			_go_fuzz_dep_.CoverTab[50158]++
															*p = fields
															messageFieldPool.Put(p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:45
			// _ = "end of CoverTab[50158]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:46
			_go_fuzz_dep_.CoverTab[50159]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:46
			// _ = "end of CoverTab[50159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:46
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:46
		// _ = "end of CoverTab[50157]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:47
	// _ = "end of CoverTab[50151]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:47
	_go_fuzz_dep_.CoverTab[50152]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:50
	fs.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:50
		_go_fuzz_dep_.CoverTab[50160]++
														fields = append(fields, messageField{fd, v})
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:52
		// _ = "end of CoverTab[50160]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:53
	// _ = "end of CoverTab[50152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:53
	_go_fuzz_dep_.CoverTab[50153]++
													sort.Slice(fields, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:54
		_go_fuzz_dep_.CoverTab[50161]++
														return less(fields[i].fd, fields[j].fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:55
		// _ = "end of CoverTab[50161]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:56
	// _ = "end of CoverTab[50153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:56
	_go_fuzz_dep_.CoverTab[50154]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:59
	for _, f := range fields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:59
		_go_fuzz_dep_.CoverTab[50162]++
														if !fn(f.fd, f.v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:60
			_go_fuzz_dep_.CoverTab[50163]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:61
			// _ = "end of CoverTab[50163]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:62
			_go_fuzz_dep_.CoverTab[50164]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:62
			// _ = "end of CoverTab[50164]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:62
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:62
		// _ = "end of CoverTab[50162]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:63
	// _ = "end of CoverTab[50154]"
}

type mapEntry struct {
	k	protoreflect.MapKey
	v	protoreflect.Value
}

var mapEntryPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:72
		_go_fuzz_dep_.CoverTab[50165]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:72
		return new([]mapEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:72
		// _ = "end of CoverTab[50165]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:72
	},
}

type (
	// EntryRanger is an interface for visiting all fields in a message.
	// The protoreflect.Map type implements this interface.
	EntryRanger	interface{ Range(VisitEntry) }
	// VisitEntry is called every time a map entry is visited.
	VisitEntry	= func(protoreflect.MapKey, protoreflect.Value) bool
)

// RangeEntries iterates over the entries of es according to the specified order.
func RangeEntries(es EntryRanger, less KeyOrder, fn VisitEntry) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:84
	_go_fuzz_dep_.CoverTab[50166]++
													if less == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:85
		_go_fuzz_dep_.CoverTab[50171]++
														es.Range(fn)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:87
		// _ = "end of CoverTab[50171]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:88
		_go_fuzz_dep_.CoverTab[50172]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:88
		// _ = "end of CoverTab[50172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:88
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:88
	// _ = "end of CoverTab[50166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:88
	_go_fuzz_dep_.CoverTab[50167]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:91
	p := mapEntryPool.Get().(*[]mapEntry)
	entries := (*p)[:0]
	defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:93
		_go_fuzz_dep_.CoverTab[50173]++
														if cap(entries) < 1024 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:94
			_go_fuzz_dep_.CoverTab[50174]++
															*p = entries
															mapEntryPool.Put(p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:96
			// _ = "end of CoverTab[50174]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:97
			_go_fuzz_dep_.CoverTab[50175]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:97
			// _ = "end of CoverTab[50175]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:97
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:97
		// _ = "end of CoverTab[50173]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:98
	// _ = "end of CoverTab[50167]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:98
	_go_fuzz_dep_.CoverTab[50168]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:101
	es.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:101
		_go_fuzz_dep_.CoverTab[50176]++
														entries = append(entries, mapEntry{k, v})
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:103
		// _ = "end of CoverTab[50176]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:104
	// _ = "end of CoverTab[50168]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:104
	_go_fuzz_dep_.CoverTab[50169]++
													sort.Slice(entries, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:105
		_go_fuzz_dep_.CoverTab[50177]++
														return less(entries[i].k, entries[j].k)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:106
		// _ = "end of CoverTab[50177]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:107
	// _ = "end of CoverTab[50169]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:107
	_go_fuzz_dep_.CoverTab[50170]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:110
	for _, e := range entries {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:110
		_go_fuzz_dep_.CoverTab[50178]++
														if !fn(e.k, e.v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:111
			_go_fuzz_dep_.CoverTab[50179]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:112
			// _ = "end of CoverTab[50179]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:113
			_go_fuzz_dep_.CoverTab[50180]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:113
			// _ = "end of CoverTab[50180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:113
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:113
		// _ = "end of CoverTab[50178]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:114
	// _ = "end of CoverTab[50170]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:115
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/range.go:115
var _ = _go_fuzz_dep_.CoverTab
