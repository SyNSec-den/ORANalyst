// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
package hpack

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:5
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
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:33
// encoded data is written to w.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:35
func NewEncoder(w io.Writer) *Encoder {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:35
	_go_fuzz_dep_.CoverTab[35087]++
										e := &Encoder{
		minSize:		uint32Max,
		maxSizeLimit:		initialHeaderTableSize,
		tableSizeUpdate:	false,
		w:			w,
	}
										e.dynTab.table.init()
										e.dynTab.setMaxSize(initialHeaderTableSize)
										return e
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:44
	// _ = "end of CoverTab[35087]"
}

// WriteField encodes f into a single Write to e's underlying Writer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:47
// This function may also produce bytes for "Header Table Size Update"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:47
// if necessary. If produced, it is done before encoding f.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:50
func (e *Encoder) WriteField(f HeaderField) error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:50
	_go_fuzz_dep_.CoverTab[35088]++
										e.buf = e.buf[:0]

										if e.tableSizeUpdate {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:53
		_go_fuzz_dep_.CoverTab[35092]++
											e.tableSizeUpdate = false
											if e.minSize < e.dynTab.maxSize {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:55
			_go_fuzz_dep_.CoverTab[35094]++
												e.buf = appendTableSize(e.buf, e.minSize)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:56
			// _ = "end of CoverTab[35094]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:57
			_go_fuzz_dep_.CoverTab[35095]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:57
			// _ = "end of CoverTab[35095]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:57
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:57
		// _ = "end of CoverTab[35092]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:57
		_go_fuzz_dep_.CoverTab[35093]++
											e.minSize = uint32Max
											e.buf = appendTableSize(e.buf, e.dynTab.maxSize)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:59
		// _ = "end of CoverTab[35093]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:60
		_go_fuzz_dep_.CoverTab[35096]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:60
		// _ = "end of CoverTab[35096]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:60
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:60
	// _ = "end of CoverTab[35088]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:60
	_go_fuzz_dep_.CoverTab[35089]++

										idx, nameValueMatch := e.searchTable(f)
										if nameValueMatch {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:63
		_go_fuzz_dep_.CoverTab[35097]++
											e.buf = appendIndexed(e.buf, idx)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:64
		// _ = "end of CoverTab[35097]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:65
		_go_fuzz_dep_.CoverTab[35098]++
											indexing := e.shouldIndex(f)
											if indexing {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:67
			_go_fuzz_dep_.CoverTab[35100]++
												e.dynTab.add(f)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:68
			// _ = "end of CoverTab[35100]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:69
			_go_fuzz_dep_.CoverTab[35101]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:69
			// _ = "end of CoverTab[35101]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:69
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:69
		// _ = "end of CoverTab[35098]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:69
		_go_fuzz_dep_.CoverTab[35099]++

											if idx == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:71
			_go_fuzz_dep_.CoverTab[35102]++
												e.buf = appendNewName(e.buf, f, indexing)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:72
			// _ = "end of CoverTab[35102]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:73
			_go_fuzz_dep_.CoverTab[35103]++
												e.buf = appendIndexedName(e.buf, f, idx, indexing)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:74
			// _ = "end of CoverTab[35103]"
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:75
		// _ = "end of CoverTab[35099]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:76
	// _ = "end of CoverTab[35089]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:76
	_go_fuzz_dep_.CoverTab[35090]++
										n, err := e.w.Write(e.buf)
										if err == nil && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:78
		_go_fuzz_dep_.CoverTab[35104]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:78
		return n != len(e.buf)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:78
		// _ = "end of CoverTab[35104]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:78
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:78
		_go_fuzz_dep_.CoverTab[35105]++
											err = io.ErrShortWrite
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:79
		// _ = "end of CoverTab[35105]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:80
		_go_fuzz_dep_.CoverTab[35106]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:80
		// _ = "end of CoverTab[35106]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:80
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:80
	// _ = "end of CoverTab[35090]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:80
	_go_fuzz_dep_.CoverTab[35091]++
										return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:81
	// _ = "end of CoverTab[35091]"
}

// searchTable searches f in both stable and dynamic header tables.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:84
// The static header table is searched first. Only when there is no
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:84
// exact match for both name and value, the dynamic header table is
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:84
// then searched. If there is no match, i is 0. If both name and value
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:84
// match, i is the matched index and nameValueMatch becomes true. If
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:84
// only name matches, i points to that index and nameValueMatch
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:84
// becomes false.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:91
func (e *Encoder) searchTable(f HeaderField) (i uint64, nameValueMatch bool) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:91
	_go_fuzz_dep_.CoverTab[35107]++
										i, nameValueMatch = staticTable.search(f)
										if nameValueMatch {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:93
		_go_fuzz_dep_.CoverTab[35110]++
											return i, true
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:94
		// _ = "end of CoverTab[35110]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:95
		_go_fuzz_dep_.CoverTab[35111]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:95
		// _ = "end of CoverTab[35111]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:95
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:95
	// _ = "end of CoverTab[35107]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:95
	_go_fuzz_dep_.CoverTab[35108]++

										j, nameValueMatch := e.dynTab.table.search(f)
										if nameValueMatch || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
		_go_fuzz_dep_.CoverTab[35112]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
		return (i == 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
			_go_fuzz_dep_.CoverTab[35113]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
			return j != 0
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
			// _ = "end of CoverTab[35113]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
		}())
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
		// _ = "end of CoverTab[35112]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:98
		_go_fuzz_dep_.CoverTab[35114]++
											return j + uint64(staticTable.len()), nameValueMatch
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:99
		// _ = "end of CoverTab[35114]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:100
		_go_fuzz_dep_.CoverTab[35115]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:100
		// _ = "end of CoverTab[35115]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:100
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:100
	// _ = "end of CoverTab[35108]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:100
	_go_fuzz_dep_.CoverTab[35109]++

										return i, false
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:102
	// _ = "end of CoverTab[35109]"
}

// SetMaxDynamicTableSize changes the dynamic header table size to v.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:105
// The actual size is bounded by the value passed to
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:105
// SetMaxDynamicTableSizeLimit.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:108
func (e *Encoder) SetMaxDynamicTableSize(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:108
	_go_fuzz_dep_.CoverTab[35116]++
										if v > e.maxSizeLimit {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:109
		_go_fuzz_dep_.CoverTab[35119]++
											v = e.maxSizeLimit
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:110
		// _ = "end of CoverTab[35119]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:111
		_go_fuzz_dep_.CoverTab[35120]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:111
		// _ = "end of CoverTab[35120]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:111
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:111
	// _ = "end of CoverTab[35116]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:111
	_go_fuzz_dep_.CoverTab[35117]++
										if v < e.minSize {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:112
		_go_fuzz_dep_.CoverTab[35121]++
											e.minSize = v
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:113
		// _ = "end of CoverTab[35121]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:114
		_go_fuzz_dep_.CoverTab[35122]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:114
		// _ = "end of CoverTab[35122]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:114
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:114
	// _ = "end of CoverTab[35117]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:114
	_go_fuzz_dep_.CoverTab[35118]++
										e.tableSizeUpdate = true
										e.dynTab.setMaxSize(v)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:116
	// _ = "end of CoverTab[35118]"
}

// MaxDynamicTableSize returns the current dynamic header table size.
func (e *Encoder) MaxDynamicTableSize() (v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:120
	_go_fuzz_dep_.CoverTab[35123]++
										return e.dynTab.maxSize
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:121
	// _ = "end of CoverTab[35123]"
}

// SetMaxDynamicTableSizeLimit changes the maximum value that can be
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:124
// specified in SetMaxDynamicTableSize to v. By default, it is set to
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:124
// 4096, which is the same size of the default dynamic header table
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:124
// size described in HPACK specification. If the current maximum
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:124
// dynamic header table size is strictly greater than v, "Header Table
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:124
// Size Update" will be done in the next WriteField call and the
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:124
// maximum dynamic header table size is truncated to v.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:131
func (e *Encoder) SetMaxDynamicTableSizeLimit(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:131
	_go_fuzz_dep_.CoverTab[35124]++
										e.maxSizeLimit = v
										if e.dynTab.maxSize > v {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:133
		_go_fuzz_dep_.CoverTab[35125]++
											e.tableSizeUpdate = true
											e.dynTab.setMaxSize(v)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:135
		// _ = "end of CoverTab[35125]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:136
		_go_fuzz_dep_.CoverTab[35126]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:136
		// _ = "end of CoverTab[35126]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:136
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:136
	// _ = "end of CoverTab[35124]"
}

// shouldIndex reports whether f should be indexed.
func (e *Encoder) shouldIndex(f HeaderField) bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:140
	_go_fuzz_dep_.CoverTab[35127]++
										return !f.Sensitive && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:141
		_go_fuzz_dep_.CoverTab[35128]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:141
		return f.Size() <= e.dynTab.maxSize
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:141
		// _ = "end of CoverTab[35128]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:141
	}()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:141
	// _ = "end of CoverTab[35127]"
}

// appendIndexed appends index i, as encoded in "Indexed Header Field"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:144
// representation, to dst and returns the extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:146
func appendIndexed(dst []byte, i uint64) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:146
	_go_fuzz_dep_.CoverTab[35129]++
										first := len(dst)
										dst = appendVarInt(dst, 7, i)
										dst[first] |= 0x80
										return dst
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:150
	// _ = "end of CoverTab[35129]"
}

// appendNewName appends f, as encoded in one of "Literal Header field
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:153
// - New Name" representation variants, to dst and returns the
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:153
// extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:153
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:153
// If f.Sensitive is true, "Never Indexed" representation is used. If
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:153
// f.Sensitive is false and indexing is true, "Incremental Indexing"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:153
// representation is used.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:160
func appendNewName(dst []byte, f HeaderField, indexing bool) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:160
	_go_fuzz_dep_.CoverTab[35130]++
										dst = append(dst, encodeTypeByte(indexing, f.Sensitive))
										dst = appendHpackString(dst, f.Name)
										return appendHpackString(dst, f.Value)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:163
	// _ = "end of CoverTab[35130]"
}

// appendIndexedName appends f and index i referring indexed name
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:166
// entry, as encoded in one of "Literal Header field - Indexed Name"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:166
// representation variants, to dst and returns the extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:166
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:166
// If f.Sensitive is true, "Never Indexed" representation is used. If
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:166
// f.Sensitive is false and indexing is true, "Incremental Indexing"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:166
// representation is used.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:173
func appendIndexedName(dst []byte, f HeaderField, i uint64, indexing bool) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:173
	_go_fuzz_dep_.CoverTab[35131]++
										first := len(dst)
										var n byte
										if indexing {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:176
		_go_fuzz_dep_.CoverTab[35133]++
											n = 6
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:177
		// _ = "end of CoverTab[35133]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:178
		_go_fuzz_dep_.CoverTab[35134]++
											n = 4
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:179
		// _ = "end of CoverTab[35134]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:180
	// _ = "end of CoverTab[35131]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:180
	_go_fuzz_dep_.CoverTab[35132]++
										dst = appendVarInt(dst, n, i)
										dst[first] |= encodeTypeByte(indexing, f.Sensitive)
										return appendHpackString(dst, f.Value)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:183
	// _ = "end of CoverTab[35132]"
}

// appendTableSize appends v, as encoded in "Header Table Size Update"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:186
// representation, to dst and returns the extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:188
func appendTableSize(dst []byte, v uint32) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:188
	_go_fuzz_dep_.CoverTab[35135]++
										first := len(dst)
										dst = appendVarInt(dst, 5, uint64(v))
										dst[first] |= 0x20
										return dst
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:192
	// _ = "end of CoverTab[35135]"
}

// appendVarInt appends i, as encoded in variable integer form using n
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:195
// bit prefix, to dst and returns the extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:195
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:195
// See
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:195
// https://httpwg.org/specs/rfc7541.html#integer.representation
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:200
func appendVarInt(dst []byte, n byte, i uint64) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:200
	_go_fuzz_dep_.CoverTab[35136]++
										k := uint64((1 << n) - 1)
										if i < k {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:202
		_go_fuzz_dep_.CoverTab[35139]++
											return append(dst, byte(i))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:203
		// _ = "end of CoverTab[35139]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:204
		_go_fuzz_dep_.CoverTab[35140]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:204
		// _ = "end of CoverTab[35140]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:204
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:204
	// _ = "end of CoverTab[35136]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:204
	_go_fuzz_dep_.CoverTab[35137]++
										dst = append(dst, byte(k))
										i -= k
										for ; i >= 128; i >>= 7 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:207
		_go_fuzz_dep_.CoverTab[35141]++
											dst = append(dst, byte(0x80|(i&0x7f)))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:208
		// _ = "end of CoverTab[35141]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:209
	// _ = "end of CoverTab[35137]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:209
	_go_fuzz_dep_.CoverTab[35138]++
										return append(dst, byte(i))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:210
	// _ = "end of CoverTab[35138]"
}

// appendHpackString appends s, as encoded in "String Literal"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:213
// representation, to dst and returns the extended buffer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:213
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:213
// s will be encoded in Huffman codes only when it produces strictly
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:213
// shorter byte string.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:218
func appendHpackString(dst []byte, s string) []byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:218
	_go_fuzz_dep_.CoverTab[35142]++
										huffmanLength := HuffmanEncodeLength(s)
										if huffmanLength < uint64(len(s)) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:220
		_go_fuzz_dep_.CoverTab[35144]++
											first := len(dst)
											dst = appendVarInt(dst, 7, huffmanLength)
											dst = AppendHuffmanString(dst, s)
											dst[first] |= 0x80
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:224
		// _ = "end of CoverTab[35144]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:225
		_go_fuzz_dep_.CoverTab[35145]++
											dst = appendVarInt(dst, 7, uint64(len(s)))
											dst = append(dst, s...)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:227
		// _ = "end of CoverTab[35145]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:228
	// _ = "end of CoverTab[35142]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:228
	_go_fuzz_dep_.CoverTab[35143]++
										return dst
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:229
	// _ = "end of CoverTab[35143]"
}

// encodeTypeByte returns type byte. If sensitive is true, type byte
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:232
// for "Never Indexed" representation is returned. If sensitive is
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:232
// false and indexing is true, type byte for "Incremental Indexing"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:232
// representation is returned. Otherwise, type byte for "Without
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:232
// Indexing" is returned.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:237
func encodeTypeByte(indexing, sensitive bool) byte {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:237
	_go_fuzz_dep_.CoverTab[35146]++
										if sensitive {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:238
		_go_fuzz_dep_.CoverTab[35149]++
											return 0x10
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:239
		// _ = "end of CoverTab[35149]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:240
		_go_fuzz_dep_.CoverTab[35150]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:240
		// _ = "end of CoverTab[35150]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:240
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:240
	// _ = "end of CoverTab[35146]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:240
	_go_fuzz_dep_.CoverTab[35147]++
										if indexing {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:241
		_go_fuzz_dep_.CoverTab[35151]++
											return 0x40
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:242
		// _ = "end of CoverTab[35151]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:243
		_go_fuzz_dep_.CoverTab[35152]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:243
		// _ = "end of CoverTab[35152]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:243
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:243
	// _ = "end of CoverTab[35147]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:243
	_go_fuzz_dep_.CoverTab[35148]++
										return 0
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:244
	// _ = "end of CoverTab[35148]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:245
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/encode.go:245
var _ = _go_fuzz_dep_.CoverTab
