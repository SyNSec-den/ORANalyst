// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
package hpack

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:5
)

import (
	"io"
)

const (
	uint32Max		= ^uint32(0)
	initialHeaderTableSize	= 4096
)

type Encoder struct {
	dynTab	dynamicTable
	// minSize is the minimum table size set by
	// SetMaxDynamicTableSize after the previous Header Table Size
	// Update.
	minSize	uint32
	// maxSizeLimit is the maximum table size this encoder
	// supports. This will protect the encoder from too large
	// size.
	maxSizeLimit	uint32
	// tableSizeUpdate indicates whether "Header Table Size
	// Update" is required.
	tableSizeUpdate	bool
	w		io.Writer
	buf		[]byte
}

// NewEncoder returns a new Encoder which performs HPACK encoding. An
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:33
// encoded data is written to w.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:35
func NewEncoder(w io.Writer) *Encoder {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:35
	_go_fuzz_dep_.CoverTab[71869]++
											e := &Encoder{
		minSize:		uint32Max,
		maxSizeLimit:		initialHeaderTableSize,
		tableSizeUpdate:	false,
		w:			w,
	}
											e.dynTab.table.init()
											e.dynTab.setMaxSize(initialHeaderTableSize)
											return e
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:44
	// _ = "end of CoverTab[71869]"
}

// WriteField encodes f into a single Write to e's underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:47
// This function may also produce bytes for "Header Table Size Update"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:47
// if necessary. If produced, it is done before encoding f.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:50
func (e *Encoder) WriteField(f HeaderField) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:50
	_go_fuzz_dep_.CoverTab[71870]++
											e.buf = e.buf[:0]

											if e.tableSizeUpdate {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:53
		_go_fuzz_dep_.CoverTab[71874]++
												e.tableSizeUpdate = false
												if e.minSize < e.dynTab.maxSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:55
			_go_fuzz_dep_.CoverTab[71876]++
													e.buf = appendTableSize(e.buf, e.minSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:56
			// _ = "end of CoverTab[71876]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:57
			_go_fuzz_dep_.CoverTab[71877]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:57
			// _ = "end of CoverTab[71877]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:57
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:57
		// _ = "end of CoverTab[71874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:57
		_go_fuzz_dep_.CoverTab[71875]++
												e.minSize = uint32Max
												e.buf = appendTableSize(e.buf, e.dynTab.maxSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:59
		// _ = "end of CoverTab[71875]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:60
		_go_fuzz_dep_.CoverTab[71878]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:60
		// _ = "end of CoverTab[71878]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:60
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:60
	// _ = "end of CoverTab[71870]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:60
	_go_fuzz_dep_.CoverTab[71871]++

											idx, nameValueMatch := e.searchTable(f)
											if nameValueMatch {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:63
		_go_fuzz_dep_.CoverTab[71879]++
												e.buf = appendIndexed(e.buf, idx)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:64
		// _ = "end of CoverTab[71879]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:65
		_go_fuzz_dep_.CoverTab[71880]++
												indexing := e.shouldIndex(f)
												if indexing {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:67
			_go_fuzz_dep_.CoverTab[71882]++
													e.dynTab.add(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:68
			// _ = "end of CoverTab[71882]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:69
			_go_fuzz_dep_.CoverTab[71883]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:69
			// _ = "end of CoverTab[71883]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:69
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:69
		// _ = "end of CoverTab[71880]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:69
		_go_fuzz_dep_.CoverTab[71881]++

												if idx == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:71
			_go_fuzz_dep_.CoverTab[71884]++
													e.buf = appendNewName(e.buf, f, indexing)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:72
			// _ = "end of CoverTab[71884]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:73
			_go_fuzz_dep_.CoverTab[71885]++
													e.buf = appendIndexedName(e.buf, f, idx, indexing)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:74
			// _ = "end of CoverTab[71885]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:75
		// _ = "end of CoverTab[71881]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:76
	// _ = "end of CoverTab[71871]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:76
	_go_fuzz_dep_.CoverTab[71872]++
											n, err := e.w.Write(e.buf)
											if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:78
		_go_fuzz_dep_.CoverTab[71886]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:78
		return n != len(e.buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:78
		// _ = "end of CoverTab[71886]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:78
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:78
		_go_fuzz_dep_.CoverTab[71887]++
												err = io.ErrShortWrite
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:79
		// _ = "end of CoverTab[71887]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:80
		_go_fuzz_dep_.CoverTab[71888]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:80
		// _ = "end of CoverTab[71888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:80
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:80
	// _ = "end of CoverTab[71872]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:80
	_go_fuzz_dep_.CoverTab[71873]++
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:81
	// _ = "end of CoverTab[71873]"
}

// searchTable searches f in both stable and dynamic header tables.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:84
// The static header table is searched first. Only when there is no
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:84
// exact match for both name and value, the dynamic header table is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:84
// then searched. If there is no match, i is 0. If both name and value
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:84
// match, i is the matched index and nameValueMatch becomes true. If
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:84
// only name matches, i points to that index and nameValueMatch
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:84
// becomes false.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:91
func (e *Encoder) searchTable(f HeaderField) (i uint64, nameValueMatch bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:91
	_go_fuzz_dep_.CoverTab[71889]++
											i, nameValueMatch = staticTable.search(f)
											if nameValueMatch {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:93
		_go_fuzz_dep_.CoverTab[71892]++
												return i, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:94
		// _ = "end of CoverTab[71892]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:95
		_go_fuzz_dep_.CoverTab[71893]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:95
		// _ = "end of CoverTab[71893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:95
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:95
	// _ = "end of CoverTab[71889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:95
	_go_fuzz_dep_.CoverTab[71890]++

											j, nameValueMatch := e.dynTab.table.search(f)
											if nameValueMatch || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
		_go_fuzz_dep_.CoverTab[71894]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
		return (i == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
			_go_fuzz_dep_.CoverTab[71895]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
			return j != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
			// _ = "end of CoverTab[71895]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
		// _ = "end of CoverTab[71894]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:98
		_go_fuzz_dep_.CoverTab[71896]++
												return j + uint64(staticTable.len()), nameValueMatch
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:99
		// _ = "end of CoverTab[71896]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:100
		_go_fuzz_dep_.CoverTab[71897]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:100
		// _ = "end of CoverTab[71897]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:100
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:100
	// _ = "end of CoverTab[71890]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:100
	_go_fuzz_dep_.CoverTab[71891]++

											return i, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:102
	// _ = "end of CoverTab[71891]"
}

// SetMaxDynamicTableSize changes the dynamic header table size to v.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:105
// The actual size is bounded by the value passed to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:105
// SetMaxDynamicTableSizeLimit.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:108
func (e *Encoder) SetMaxDynamicTableSize(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:108
	_go_fuzz_dep_.CoverTab[71898]++
											if v > e.maxSizeLimit {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:109
		_go_fuzz_dep_.CoverTab[71901]++
												v = e.maxSizeLimit
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:110
		// _ = "end of CoverTab[71901]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:111
		_go_fuzz_dep_.CoverTab[71902]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:111
		// _ = "end of CoverTab[71902]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:111
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:111
	// _ = "end of CoverTab[71898]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:111
	_go_fuzz_dep_.CoverTab[71899]++
											if v < e.minSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:112
		_go_fuzz_dep_.CoverTab[71903]++
												e.minSize = v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:113
		// _ = "end of CoverTab[71903]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:114
		_go_fuzz_dep_.CoverTab[71904]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:114
		// _ = "end of CoverTab[71904]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:114
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:114
	// _ = "end of CoverTab[71899]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:114
	_go_fuzz_dep_.CoverTab[71900]++
											e.tableSizeUpdate = true
											e.dynTab.setMaxSize(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:116
	// _ = "end of CoverTab[71900]"
}

// MaxDynamicTableSize returns the current dynamic header table size.
func (e *Encoder) MaxDynamicTableSize() (v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:120
	_go_fuzz_dep_.CoverTab[71905]++
											return e.dynTab.maxSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:121
	// _ = "end of CoverTab[71905]"
}

// SetMaxDynamicTableSizeLimit changes the maximum value that can be
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:124
// specified in SetMaxDynamicTableSize to v. By default, it is set to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:124
// 4096, which is the same size of the default dynamic header table
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:124
// size described in HPACK specification. If the current maximum
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:124
// dynamic header table size is strictly greater than v, "Header Table
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:124
// Size Update" will be done in the next WriteField call and the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:124
// maximum dynamic header table size is truncated to v.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:131
func (e *Encoder) SetMaxDynamicTableSizeLimit(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:131
	_go_fuzz_dep_.CoverTab[71906]++
											e.maxSizeLimit = v
											if e.dynTab.maxSize > v {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:133
		_go_fuzz_dep_.CoverTab[71907]++
												e.tableSizeUpdate = true
												e.dynTab.setMaxSize(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:135
		// _ = "end of CoverTab[71907]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:136
		_go_fuzz_dep_.CoverTab[71908]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:136
		// _ = "end of CoverTab[71908]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:136
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:136
	// _ = "end of CoverTab[71906]"
}

// shouldIndex reports whether f should be indexed.
func (e *Encoder) shouldIndex(f HeaderField) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:140
	_go_fuzz_dep_.CoverTab[71909]++
											return !f.Sensitive && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:141
		_go_fuzz_dep_.CoverTab[71910]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:141
		return f.Size() <= e.dynTab.maxSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:141
		// _ = "end of CoverTab[71910]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:141
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:141
	// _ = "end of CoverTab[71909]"
}

// appendIndexed appends index i, as encoded in "Indexed Header Field"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:144
// representation, to dst and returns the extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:146
func appendIndexed(dst []byte, i uint64) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:146
	_go_fuzz_dep_.CoverTab[71911]++
											first := len(dst)
											dst = appendVarInt(dst, 7, i)
											dst[first] |= 0x80
											return dst
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:150
	// _ = "end of CoverTab[71911]"
}

// appendNewName appends f, as encoded in one of "Literal Header field
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:153
// - New Name" representation variants, to dst and returns the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:153
// extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:153
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:153
// If f.Sensitive is true, "Never Indexed" representation is used. If
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:153
// f.Sensitive is false and indexing is true, "Incremental Indexing"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:153
// representation is used.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:160
func appendNewName(dst []byte, f HeaderField, indexing bool) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:160
	_go_fuzz_dep_.CoverTab[71912]++
											dst = append(dst, encodeTypeByte(indexing, f.Sensitive))
											dst = appendHpackString(dst, f.Name)
											return appendHpackString(dst, f.Value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:163
	// _ = "end of CoverTab[71912]"
}

// appendIndexedName appends f and index i referring indexed name
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:166
// entry, as encoded in one of "Literal Header field - Indexed Name"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:166
// representation variants, to dst and returns the extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:166
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:166
// If f.Sensitive is true, "Never Indexed" representation is used. If
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:166
// f.Sensitive is false and indexing is true, "Incremental Indexing"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:166
// representation is used.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:173
func appendIndexedName(dst []byte, f HeaderField, i uint64, indexing bool) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:173
	_go_fuzz_dep_.CoverTab[71913]++
											first := len(dst)
											var n byte
											if indexing {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:176
		_go_fuzz_dep_.CoverTab[71915]++
												n = 6
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:177
		// _ = "end of CoverTab[71915]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:178
		_go_fuzz_dep_.CoverTab[71916]++
												n = 4
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:179
		// _ = "end of CoverTab[71916]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:180
	// _ = "end of CoverTab[71913]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:180
	_go_fuzz_dep_.CoverTab[71914]++
											dst = appendVarInt(dst, n, i)
											dst[first] |= encodeTypeByte(indexing, f.Sensitive)
											return appendHpackString(dst, f.Value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:183
	// _ = "end of CoverTab[71914]"
}

// appendTableSize appends v, as encoded in "Header Table Size Update"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:186
// representation, to dst and returns the extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:188
func appendTableSize(dst []byte, v uint32) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:188
	_go_fuzz_dep_.CoverTab[71917]++
											first := len(dst)
											dst = appendVarInt(dst, 5, uint64(v))
											dst[first] |= 0x20
											return dst
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:192
	// _ = "end of CoverTab[71917]"
}

// appendVarInt appends i, as encoded in variable integer form using n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:195
// bit prefix, to dst and returns the extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:195
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:195
// See
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:195
// https://httpwg.org/specs/rfc7541.html#integer.representation
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:200
func appendVarInt(dst []byte, n byte, i uint64) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:200
	_go_fuzz_dep_.CoverTab[71918]++
											k := uint64((1 << n) - 1)
											if i < k {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:202
		_go_fuzz_dep_.CoverTab[71921]++
												return append(dst, byte(i))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:203
		// _ = "end of CoverTab[71921]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:204
		_go_fuzz_dep_.CoverTab[71922]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:204
		// _ = "end of CoverTab[71922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:204
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:204
	// _ = "end of CoverTab[71918]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:204
	_go_fuzz_dep_.CoverTab[71919]++
											dst = append(dst, byte(k))
											i -= k
											for ; i >= 128; i >>= 7 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:207
		_go_fuzz_dep_.CoverTab[71923]++
												dst = append(dst, byte(0x80|(i&0x7f)))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:208
		// _ = "end of CoverTab[71923]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:209
	// _ = "end of CoverTab[71919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:209
	_go_fuzz_dep_.CoverTab[71920]++
											return append(dst, byte(i))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:210
	// _ = "end of CoverTab[71920]"
}

// appendHpackString appends s, as encoded in "String Literal"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:213
// representation, to dst and returns the extended buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:213
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:213
// s will be encoded in Huffman codes only when it produces strictly
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:213
// shorter byte string.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:218
func appendHpackString(dst []byte, s string) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:218
	_go_fuzz_dep_.CoverTab[71924]++
											huffmanLength := HuffmanEncodeLength(s)
											if huffmanLength < uint64(len(s)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:220
		_go_fuzz_dep_.CoverTab[71926]++
												first := len(dst)
												dst = appendVarInt(dst, 7, huffmanLength)
												dst = AppendHuffmanString(dst, s)
												dst[first] |= 0x80
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:224
		// _ = "end of CoverTab[71926]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:225
		_go_fuzz_dep_.CoverTab[71927]++
												dst = appendVarInt(dst, 7, uint64(len(s)))
												dst = append(dst, s...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:227
		// _ = "end of CoverTab[71927]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:228
	// _ = "end of CoverTab[71924]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:228
	_go_fuzz_dep_.CoverTab[71925]++
											return dst
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:229
	// _ = "end of CoverTab[71925]"
}

// encodeTypeByte returns type byte. If sensitive is true, type byte
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:232
// for "Never Indexed" representation is returned. If sensitive is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:232
// false and indexing is true, type byte for "Incremental Indexing"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:232
// representation is returned. Otherwise, type byte for "Without
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:232
// Indexing" is returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:237
func encodeTypeByte(indexing, sensitive bool) byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:237
	_go_fuzz_dep_.CoverTab[71928]++
											if sensitive {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:238
		_go_fuzz_dep_.CoverTab[71931]++
												return 0x10
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:239
		// _ = "end of CoverTab[71931]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:240
		_go_fuzz_dep_.CoverTab[71932]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:240
		// _ = "end of CoverTab[71932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:240
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:240
	// _ = "end of CoverTab[71928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:240
	_go_fuzz_dep_.CoverTab[71929]++
											if indexing {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:241
		_go_fuzz_dep_.CoverTab[71933]++
												return 0x40
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:242
		// _ = "end of CoverTab[71933]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:243
		_go_fuzz_dep_.CoverTab[71934]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:243
		// _ = "end of CoverTab[71934]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:243
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:243
	// _ = "end of CoverTab[71929]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:243
	_go_fuzz_dep_.CoverTab[71930]++
											return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:244
	// _ = "end of CoverTab[71930]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:245
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/encode.go:245
var _ = _go_fuzz_dep_.CoverTab
