// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:5
// Package hpack implements HPACK, a compression format for
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:5
// efficiently representing HTTP header fields in the context of HTTP/2.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:5
// See http://tools.ietf.org/html/draft-ietf-httpbis-header-compression-09
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
package hpack

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:9
)

import (
	"bytes"
	"errors"
	"fmt"
)

// A DecodingError is something the spec defines as a decoding error.
type DecodingError struct {
	Err error
}

func (de DecodingError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:22
	_go_fuzz_dep_.CoverTab[71935]++
											return fmt.Sprintf("decoding error: %v", de.Err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:23
	// _ = "end of CoverTab[71935]"
}

// An InvalidIndexError is returned when an encoder references a table
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:26
// entry before the static table or after the end of the dynamic table.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:28
type InvalidIndexError int

func (e InvalidIndexError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:30
	_go_fuzz_dep_.CoverTab[71936]++
											return fmt.Sprintf("invalid indexed representation index %d", int(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:31
	// _ = "end of CoverTab[71936]"
}

// A HeaderField is a name-value pair. Both the name and value are
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:34
// treated as opaque sequences of octets.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:36
type HeaderField struct {
	Name, Value	string

	// Sensitive means that this header field should never be
	// indexed.
	Sensitive	bool
}

// IsPseudo reports whether the header field is an http2 pseudo header.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:44
// That is, it reports whether it starts with a colon.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:44
// It is not otherwise guaranteed to be a valid pseudo header field,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:44
// though.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:48
func (hf HeaderField) IsPseudo() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:48
	_go_fuzz_dep_.CoverTab[71937]++
											return len(hf.Name) != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:49
		_go_fuzz_dep_.CoverTab[71938]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:49
		return hf.Name[0] == ':'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:49
		// _ = "end of CoverTab[71938]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:49
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:49
	// _ = "end of CoverTab[71937]"
}

func (hf HeaderField) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:52
	_go_fuzz_dep_.CoverTab[71939]++
											var suffix string
											if hf.Sensitive {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:54
		_go_fuzz_dep_.CoverTab[71941]++
												suffix = " (sensitive)"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:55
		// _ = "end of CoverTab[71941]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:56
		_go_fuzz_dep_.CoverTab[71942]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:56
		// _ = "end of CoverTab[71942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:56
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:56
	// _ = "end of CoverTab[71939]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:56
	_go_fuzz_dep_.CoverTab[71940]++
											return fmt.Sprintf("header field %q = %q%s", hf.Name, hf.Value, suffix)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:57
	// _ = "end of CoverTab[71940]"
}

// Size returns the size of an entry per RFC 7541 section 4.1.
func (hf HeaderField) Size() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:61
	_go_fuzz_dep_.CoverTab[71943]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:74
	return uint32(len(hf.Name) + len(hf.Value) + 32)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:74
	// _ = "end of CoverTab[71943]"
}

// A Decoder is the decoding context for incremental processing of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:77
// header blocks.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:79
type Decoder struct {
	dynTab	dynamicTable
	emit	func(f HeaderField)

	emitEnabled	bool	// whether calls to emit are enabled
	maxStrLen	int	// 0 means unlimited

	// buf is the unparsed buffer. It's only written to
	// saveBuf if it was truncated in the middle of a header
	// block. Because it's usually not owned, we can only
	// process it under Write.
	buf	[]byte	// not owned; only valid during Write

	// saveBuf is previous data passed to Write which we weren't able
	// to fully parse before. Unlike buf, we own this data.
	saveBuf	bytes.Buffer

	firstField	bool	// processing the first field of the header block
}

// NewDecoder returns a new decoder with the provided maximum dynamic
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:99
// table size. The emitFunc will be called for each valid field
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:99
// parsed, in the same goroutine as calls to Write, before Write returns.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:102
func NewDecoder(maxDynamicTableSize uint32, emitFunc func(f HeaderField)) *Decoder {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:102
	_go_fuzz_dep_.CoverTab[71944]++
											d := &Decoder{
		emit:		emitFunc,
		emitEnabled:	true,
		firstField:	true,
	}
											d.dynTab.table.init()
											d.dynTab.allowedMaxSize = maxDynamicTableSize
											d.dynTab.setMaxSize(maxDynamicTableSize)
											return d
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:111
	// _ = "end of CoverTab[71944]"
}

// ErrStringLength is returned by Decoder.Write when the max string length
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:114
// (as configured by Decoder.SetMaxStringLength) would be violated.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:116
var ErrStringLength = errors.New("hpack: string too long")

// SetMaxStringLength sets the maximum size of a HeaderField name or
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:118
// value string. If a string exceeds this length (even after any
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:118
// decompression), Write will return ErrStringLength.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:118
// A value of 0 means unlimited and is the default from NewDecoder.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:122
func (d *Decoder) SetMaxStringLength(n int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:122
	_go_fuzz_dep_.CoverTab[71945]++
											d.maxStrLen = n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:123
	// _ = "end of CoverTab[71945]"
}

// SetEmitFunc changes the callback used when new header fields
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:126
// are decoded.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:126
// It must be non-nil. It does not affect EmitEnabled.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:129
func (d *Decoder) SetEmitFunc(emitFunc func(f HeaderField)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:129
	_go_fuzz_dep_.CoverTab[71946]++
											d.emit = emitFunc
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:130
	// _ = "end of CoverTab[71946]"
}

// SetEmitEnabled controls whether the emitFunc provided to NewDecoder
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:133
// should be called. The default is true.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:133
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:133
// This facility exists to let servers enforce MAX_HEADER_LIST_SIZE
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:133
// while still decoding and keeping in-sync with decoder state, but
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:133
// without doing unnecessary decompression or generating unnecessary
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:133
// garbage for header fields past the limit.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:140
func (d *Decoder) SetEmitEnabled(v bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:140
	_go_fuzz_dep_.CoverTab[71947]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:140
	d.emitEnabled = v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:140
	// _ = "end of CoverTab[71947]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:140
}

// EmitEnabled reports whether calls to the emitFunc provided to NewDecoder
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:142
// are currently enabled. The default is true.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:144
func (d *Decoder) EmitEnabled() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:144
	_go_fuzz_dep_.CoverTab[71948]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:144
	return d.emitEnabled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:144
	// _ = "end of CoverTab[71948]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:144
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:149
func (d *Decoder) SetMaxDynamicTableSize(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:149
	_go_fuzz_dep_.CoverTab[71949]++
											d.dynTab.setMaxSize(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:150
	// _ = "end of CoverTab[71949]"
}

// SetAllowedMaxDynamicTableSize sets the upper bound that the encoded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:153
// stream (via dynamic table size updates) may set the maximum size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:153
// to.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:156
func (d *Decoder) SetAllowedMaxDynamicTableSize(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:156
	_go_fuzz_dep_.CoverTab[71950]++
											d.dynTab.allowedMaxSize = v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:157
	// _ = "end of CoverTab[71950]"
}

type dynamicTable struct {
	// https://httpwg.org/specs/rfc7541.html#rfc.section.2.3.2
	table		headerFieldTable
	size		uint32	// in bytes
	maxSize		uint32	// current maxSize
	allowedMaxSize	uint32	// maxSize may go up to this, inclusive
}

func (dt *dynamicTable) setMaxSize(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:168
	_go_fuzz_dep_.CoverTab[71951]++
											dt.maxSize = v
											dt.evict()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:170
	// _ = "end of CoverTab[71951]"
}

func (dt *dynamicTable) add(f HeaderField) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:173
	_go_fuzz_dep_.CoverTab[71952]++
											dt.table.addEntry(f)
											dt.size += f.Size()
											dt.evict()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:176
	// _ = "end of CoverTab[71952]"
}

// If we're too big, evict old stuff.
func (dt *dynamicTable) evict() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:180
	_go_fuzz_dep_.CoverTab[71953]++
											var n int
											for dt.size > dt.maxSize && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:182
		_go_fuzz_dep_.CoverTab[71955]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:182
		return n < dt.table.len()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:182
		// _ = "end of CoverTab[71955]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:182
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:182
		_go_fuzz_dep_.CoverTab[71956]++
												dt.size -= dt.table.ents[n].Size()
												n++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:184
		// _ = "end of CoverTab[71956]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:185
	// _ = "end of CoverTab[71953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:185
	_go_fuzz_dep_.CoverTab[71954]++
											dt.table.evictOldest(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:186
	// _ = "end of CoverTab[71954]"
}

func (d *Decoder) maxTableIndex() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:189
	_go_fuzz_dep_.CoverTab[71957]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:193
	return d.dynTab.table.len() + staticTable.len()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:193
	// _ = "end of CoverTab[71957]"
}

func (d *Decoder) at(i uint64) (hf HeaderField, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:196
	_go_fuzz_dep_.CoverTab[71958]++

											if i == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:198
		_go_fuzz_dep_.CoverTab[71962]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:199
		// _ = "end of CoverTab[71962]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:200
		_go_fuzz_dep_.CoverTab[71963]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:200
		// _ = "end of CoverTab[71963]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:200
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:200
	// _ = "end of CoverTab[71958]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:200
	_go_fuzz_dep_.CoverTab[71959]++
											if i <= uint64(staticTable.len()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:201
		_go_fuzz_dep_.CoverTab[71964]++
												return staticTable.ents[i-1], true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:202
		// _ = "end of CoverTab[71964]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:203
		_go_fuzz_dep_.CoverTab[71965]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:203
		// _ = "end of CoverTab[71965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:203
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:203
	// _ = "end of CoverTab[71959]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:203
	_go_fuzz_dep_.CoverTab[71960]++
											if i > uint64(d.maxTableIndex()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:204
		_go_fuzz_dep_.CoverTab[71966]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:205
		// _ = "end of CoverTab[71966]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:206
		_go_fuzz_dep_.CoverTab[71967]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:206
		// _ = "end of CoverTab[71967]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:206
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:206
	// _ = "end of CoverTab[71960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:206
	_go_fuzz_dep_.CoverTab[71961]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:210
	dt := d.dynTab.table
											return dt.ents[dt.len()-(int(i)-staticTable.len())], true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:211
	// _ = "end of CoverTab[71961]"
}

// DecodeFull decodes an entire block.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:214
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:214
// TODO: remove this method and make it incremental later? This is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:214
// easier for debugging now.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:218
func (d *Decoder) DecodeFull(p []byte) ([]HeaderField, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:218
	_go_fuzz_dep_.CoverTab[71968]++
											var hf []HeaderField
											saveFunc := d.emit
											defer func() { _go_fuzz_dep_.CoverTab[71973]++; d.emit = saveFunc; // _ = "end of CoverTab[71973]" }()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:221
	// _ = "end of CoverTab[71968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:221
	_go_fuzz_dep_.CoverTab[71969]++
											d.emit = func(f HeaderField) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:222
		_go_fuzz_dep_.CoverTab[71974]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:222
		hf = append(hf, f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:222
		// _ = "end of CoverTab[71974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:222
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:222
	// _ = "end of CoverTab[71969]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:222
	_go_fuzz_dep_.CoverTab[71970]++
											if _, err := d.Write(p); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:223
		_go_fuzz_dep_.CoverTab[71975]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:224
		// _ = "end of CoverTab[71975]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:225
		_go_fuzz_dep_.CoverTab[71976]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:225
		// _ = "end of CoverTab[71976]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:225
	// _ = "end of CoverTab[71970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:225
	_go_fuzz_dep_.CoverTab[71971]++
											if err := d.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:226
		_go_fuzz_dep_.CoverTab[71977]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:227
		// _ = "end of CoverTab[71977]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:228
		_go_fuzz_dep_.CoverTab[71978]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:228
		// _ = "end of CoverTab[71978]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:228
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:228
	// _ = "end of CoverTab[71971]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:228
	_go_fuzz_dep_.CoverTab[71972]++
											return hf, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:229
	// _ = "end of CoverTab[71972]"
}

// Close declares that the decoding is complete and resets the Decoder
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:232
// to be reused again for a new header block. If there is any remaining
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:232
// data in the decoder's buffer, Close returns an error.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:235
func (d *Decoder) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:235
	_go_fuzz_dep_.CoverTab[71979]++
											if d.saveBuf.Len() > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:236
		_go_fuzz_dep_.CoverTab[71981]++
												d.saveBuf.Reset()
												return DecodingError{errors.New("truncated headers")}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:238
		// _ = "end of CoverTab[71981]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:239
		_go_fuzz_dep_.CoverTab[71982]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:239
		// _ = "end of CoverTab[71982]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:239
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:239
	// _ = "end of CoverTab[71979]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:239
	_go_fuzz_dep_.CoverTab[71980]++
											d.firstField = true
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:241
	// _ = "end of CoverTab[71980]"
}

func (d *Decoder) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:244
	_go_fuzz_dep_.CoverTab[71983]++
											if len(p) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:245
		_go_fuzz_dep_.CoverTab[71987]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:249
		return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:249
		// _ = "end of CoverTab[71987]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:250
		_go_fuzz_dep_.CoverTab[71988]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:250
		// _ = "end of CoverTab[71988]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:250
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:250
	// _ = "end of CoverTab[71983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:250
	_go_fuzz_dep_.CoverTab[71984]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:253
	if d.saveBuf.Len() == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:253
		_go_fuzz_dep_.CoverTab[71989]++
												d.buf = p
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:254
		// _ = "end of CoverTab[71989]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:255
		_go_fuzz_dep_.CoverTab[71990]++
												d.saveBuf.Write(p)
												d.buf = d.saveBuf.Bytes()
												d.saveBuf.Reset()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:258
		// _ = "end of CoverTab[71990]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:259
	// _ = "end of CoverTab[71984]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:259
	_go_fuzz_dep_.CoverTab[71985]++

											for len(d.buf) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:261
		_go_fuzz_dep_.CoverTab[71991]++
												err = d.parseHeaderFieldRepr()
												if err == errNeedMore {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:263
			_go_fuzz_dep_.CoverTab[71993]++
			// Extra paranoia, making sure saveBuf won't
			// get too large. All the varint and string
			// reading code earlier should already catch
			// overlong things and return ErrStringLength,
			// but keep this as a last resort.
			const varIntOverhead = 8	// conservative
			if d.maxStrLen != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:270
				_go_fuzz_dep_.CoverTab[71995]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:270
				return int64(len(d.buf)) > 2*(int64(d.maxStrLen)+varIntOverhead)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:270
				// _ = "end of CoverTab[71995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:270
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:270
				_go_fuzz_dep_.CoverTab[71996]++
														return 0, ErrStringLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:271
				// _ = "end of CoverTab[71996]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:272
				_go_fuzz_dep_.CoverTab[71997]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:272
				// _ = "end of CoverTab[71997]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:272
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:272
			// _ = "end of CoverTab[71993]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:272
			_go_fuzz_dep_.CoverTab[71994]++
													d.saveBuf.Write(d.buf)
													return len(p), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:274
			// _ = "end of CoverTab[71994]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:275
			_go_fuzz_dep_.CoverTab[71998]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:275
			// _ = "end of CoverTab[71998]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:275
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:275
		// _ = "end of CoverTab[71991]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:275
		_go_fuzz_dep_.CoverTab[71992]++
												d.firstField = false
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:277
			_go_fuzz_dep_.CoverTab[71999]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:278
			// _ = "end of CoverTab[71999]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:279
			_go_fuzz_dep_.CoverTab[72000]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:279
			// _ = "end of CoverTab[72000]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:279
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:279
		// _ = "end of CoverTab[71992]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:280
	// _ = "end of CoverTab[71985]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:280
	_go_fuzz_dep_.CoverTab[71986]++
											return len(p), err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:281
	// _ = "end of CoverTab[71986]"
}

// errNeedMore is an internal sentinel error value that means the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:284
// buffer is truncated and we need to read more data before we can
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:284
// continue parsing.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:287
var errNeedMore = errors.New("need more data")

type indexType int

const (
	indexedTrue	indexType	= iota
	indexedFalse
	indexedNever
)

func (v indexType) indexed() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:297
	_go_fuzz_dep_.CoverTab[72001]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:297
	return v == indexedTrue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:297
	// _ = "end of CoverTab[72001]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:297
}
func (v indexType) sensitive() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:298
	_go_fuzz_dep_.CoverTab[72002]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:298
	return v == indexedNever
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:298
	// _ = "end of CoverTab[72002]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:298
}

// returns errNeedMore if there isn't enough data available.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:300
// any other error is fatal.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:300
// consumes d.buf iff it returns nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:300
// precondition: must be called with len(d.buf) > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:304
func (d *Decoder) parseHeaderFieldRepr() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:304
	_go_fuzz_dep_.CoverTab[72003]++
											b := d.buf[0]
											switch {
	case b&128 != 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:307
		_go_fuzz_dep_.CoverTab[72005]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:311
		return d.parseFieldIndexed()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:311
		// _ = "end of CoverTab[72005]"
	case b&192 == 64:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:312
		_go_fuzz_dep_.CoverTab[72006]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:316
		return d.parseFieldLiteral(6, indexedTrue)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:316
		// _ = "end of CoverTab[72006]"
	case b&240 == 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:317
		_go_fuzz_dep_.CoverTab[72007]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:321
		return d.parseFieldLiteral(4, indexedFalse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:321
		// _ = "end of CoverTab[72007]"
	case b&240 == 16:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:322
		_go_fuzz_dep_.CoverTab[72008]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:326
		return d.parseFieldLiteral(4, indexedNever)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:326
		// _ = "end of CoverTab[72008]"
	case b&224 == 32:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:327
		_go_fuzz_dep_.CoverTab[72009]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:331
		return d.parseDynamicTableSizeUpdate()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:331
		// _ = "end of CoverTab[72009]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:331
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:331
		_go_fuzz_dep_.CoverTab[72010]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:331
		// _ = "end of CoverTab[72010]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:332
	// _ = "end of CoverTab[72003]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:332
	_go_fuzz_dep_.CoverTab[72004]++

											return DecodingError{errors.New("invalid encoding")}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:334
	// _ = "end of CoverTab[72004]"
}

// (same invariants and behavior as parseHeaderFieldRepr)
func (d *Decoder) parseFieldIndexed() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:338
	_go_fuzz_dep_.CoverTab[72011]++
											buf := d.buf
											idx, buf, err := readVarInt(7, buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:341
		_go_fuzz_dep_.CoverTab[72014]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:342
		// _ = "end of CoverTab[72014]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:343
		_go_fuzz_dep_.CoverTab[72015]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:343
		// _ = "end of CoverTab[72015]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:343
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:343
	// _ = "end of CoverTab[72011]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:343
	_go_fuzz_dep_.CoverTab[72012]++
											hf, ok := d.at(idx)
											if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:345
		_go_fuzz_dep_.CoverTab[72016]++
												return DecodingError{InvalidIndexError(idx)}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:346
		// _ = "end of CoverTab[72016]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:347
		_go_fuzz_dep_.CoverTab[72017]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:347
		// _ = "end of CoverTab[72017]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:347
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:347
	// _ = "end of CoverTab[72012]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:347
	_go_fuzz_dep_.CoverTab[72013]++
											d.buf = buf
											return d.callEmit(HeaderField{Name: hf.Name, Value: hf.Value})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:349
	// _ = "end of CoverTab[72013]"
}

// (same invariants and behavior as parseHeaderFieldRepr)
func (d *Decoder) parseFieldLiteral(n uint8, it indexType) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:353
	_go_fuzz_dep_.CoverTab[72018]++
											buf := d.buf
											nameIdx, buf, err := readVarInt(n, buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:356
		_go_fuzz_dep_.CoverTab[72024]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:357
		// _ = "end of CoverTab[72024]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:358
		_go_fuzz_dep_.CoverTab[72025]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:358
		// _ = "end of CoverTab[72025]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:358
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:358
	// _ = "end of CoverTab[72018]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:358
	_go_fuzz_dep_.CoverTab[72019]++

											var hf HeaderField
											wantStr := d.emitEnabled || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:361
		_go_fuzz_dep_.CoverTab[72026]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:361
		return it.indexed()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:361
		// _ = "end of CoverTab[72026]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:361
	}()
											var undecodedName undecodedString
											if nameIdx > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:363
		_go_fuzz_dep_.CoverTab[72027]++
												ihf, ok := d.at(nameIdx)
												if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:365
			_go_fuzz_dep_.CoverTab[72029]++
													return DecodingError{InvalidIndexError(nameIdx)}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:366
			// _ = "end of CoverTab[72029]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:367
			_go_fuzz_dep_.CoverTab[72030]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:367
			// _ = "end of CoverTab[72030]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:367
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:367
		// _ = "end of CoverTab[72027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:367
		_go_fuzz_dep_.CoverTab[72028]++
												hf.Name = ihf.Name
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:368
		// _ = "end of CoverTab[72028]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:369
		_go_fuzz_dep_.CoverTab[72031]++
												undecodedName, buf, err = d.readString(buf)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:371
			_go_fuzz_dep_.CoverTab[72032]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:372
			// _ = "end of CoverTab[72032]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:373
			_go_fuzz_dep_.CoverTab[72033]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:373
			// _ = "end of CoverTab[72033]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:373
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:373
		// _ = "end of CoverTab[72031]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:374
	// _ = "end of CoverTab[72019]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:374
	_go_fuzz_dep_.CoverTab[72020]++
											undecodedValue, buf, err := d.readString(buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:376
		_go_fuzz_dep_.CoverTab[72034]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:377
		// _ = "end of CoverTab[72034]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:378
		_go_fuzz_dep_.CoverTab[72035]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:378
		// _ = "end of CoverTab[72035]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:378
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:378
	// _ = "end of CoverTab[72020]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:378
	_go_fuzz_dep_.CoverTab[72021]++
											if wantStr {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:379
		_go_fuzz_dep_.CoverTab[72036]++
												if nameIdx <= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:380
			_go_fuzz_dep_.CoverTab[72038]++
													hf.Name, err = d.decodeString(undecodedName)
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:382
				_go_fuzz_dep_.CoverTab[72039]++
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:383
				// _ = "end of CoverTab[72039]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:384
				_go_fuzz_dep_.CoverTab[72040]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:384
				// _ = "end of CoverTab[72040]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:384
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:384
			// _ = "end of CoverTab[72038]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:385
			_go_fuzz_dep_.CoverTab[72041]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:385
			// _ = "end of CoverTab[72041]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:385
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:385
		// _ = "end of CoverTab[72036]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:385
		_go_fuzz_dep_.CoverTab[72037]++
												hf.Value, err = d.decodeString(undecodedValue)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:387
			_go_fuzz_dep_.CoverTab[72042]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:388
			// _ = "end of CoverTab[72042]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:389
			_go_fuzz_dep_.CoverTab[72043]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:389
			// _ = "end of CoverTab[72043]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:389
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:389
		// _ = "end of CoverTab[72037]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:390
		_go_fuzz_dep_.CoverTab[72044]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:390
		// _ = "end of CoverTab[72044]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:390
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:390
	// _ = "end of CoverTab[72021]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:390
	_go_fuzz_dep_.CoverTab[72022]++
											d.buf = buf
											if it.indexed() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:392
		_go_fuzz_dep_.CoverTab[72045]++
												d.dynTab.add(hf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:393
		// _ = "end of CoverTab[72045]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:394
		_go_fuzz_dep_.CoverTab[72046]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:394
		// _ = "end of CoverTab[72046]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:394
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:394
	// _ = "end of CoverTab[72022]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:394
	_go_fuzz_dep_.CoverTab[72023]++
											hf.Sensitive = it.sensitive()
											return d.callEmit(hf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:396
	// _ = "end of CoverTab[72023]"
}

func (d *Decoder) callEmit(hf HeaderField) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:399
	_go_fuzz_dep_.CoverTab[72047]++
											if d.maxStrLen != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:400
		_go_fuzz_dep_.CoverTab[72050]++
												if len(hf.Name) > d.maxStrLen || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:401
			_go_fuzz_dep_.CoverTab[72051]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:401
			return len(hf.Value) > d.maxStrLen
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:401
			// _ = "end of CoverTab[72051]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:401
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:401
			_go_fuzz_dep_.CoverTab[72052]++
													return ErrStringLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:402
			// _ = "end of CoverTab[72052]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:403
			_go_fuzz_dep_.CoverTab[72053]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:403
			// _ = "end of CoverTab[72053]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:403
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:403
		// _ = "end of CoverTab[72050]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:404
		_go_fuzz_dep_.CoverTab[72054]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:404
		// _ = "end of CoverTab[72054]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:404
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:404
	// _ = "end of CoverTab[72047]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:404
	_go_fuzz_dep_.CoverTab[72048]++
											if d.emitEnabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:405
		_go_fuzz_dep_.CoverTab[72055]++
												d.emit(hf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:406
		// _ = "end of CoverTab[72055]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:407
		_go_fuzz_dep_.CoverTab[72056]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:407
		// _ = "end of CoverTab[72056]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:407
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:407
	// _ = "end of CoverTab[72048]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:407
	_go_fuzz_dep_.CoverTab[72049]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:408
	// _ = "end of CoverTab[72049]"
}

// (same invariants and behavior as parseHeaderFieldRepr)
func (d *Decoder) parseDynamicTableSizeUpdate() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:412
	_go_fuzz_dep_.CoverTab[72057]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:415
	if !d.firstField && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:415
		_go_fuzz_dep_.CoverTab[72061]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:415
		return d.dynTab.size > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:415
		// _ = "end of CoverTab[72061]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:415
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:415
		_go_fuzz_dep_.CoverTab[72062]++
												return DecodingError{errors.New("dynamic table size update MUST occur at the beginning of a header block")}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:416
		// _ = "end of CoverTab[72062]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:417
		_go_fuzz_dep_.CoverTab[72063]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:417
		// _ = "end of CoverTab[72063]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:417
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:417
	// _ = "end of CoverTab[72057]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:417
	_go_fuzz_dep_.CoverTab[72058]++

											buf := d.buf
											size, buf, err := readVarInt(5, buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:421
		_go_fuzz_dep_.CoverTab[72064]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:422
		// _ = "end of CoverTab[72064]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:423
		_go_fuzz_dep_.CoverTab[72065]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:423
		// _ = "end of CoverTab[72065]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:423
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:423
	// _ = "end of CoverTab[72058]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:423
	_go_fuzz_dep_.CoverTab[72059]++
											if size > uint64(d.dynTab.allowedMaxSize) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:424
		_go_fuzz_dep_.CoverTab[72066]++
												return DecodingError{errors.New("dynamic table size update too large")}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:425
		// _ = "end of CoverTab[72066]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:426
		_go_fuzz_dep_.CoverTab[72067]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:426
		// _ = "end of CoverTab[72067]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:426
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:426
	// _ = "end of CoverTab[72059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:426
	_go_fuzz_dep_.CoverTab[72060]++
											d.dynTab.setMaxSize(uint32(size))
											d.buf = buf
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:429
	// _ = "end of CoverTab[72060]"
}

var errVarintOverflow = DecodingError{errors.New("varint integer overflow")}

// readVarInt reads an unsigned variable length integer off the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
// beginning of p. n is the parameter as described in
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
// https://httpwg.org/specs/rfc7541.html#rfc.section.5.1.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
// n must always be between 1 and 8.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
// The returned remain buffer is either a smaller suffix of p, or err != nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:434
// The error is errNeedMore if p doesn't contain a complete integer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:442
func readVarInt(n byte, p []byte) (i uint64, remain []byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:442
	_go_fuzz_dep_.CoverTab[72068]++
											if n < 1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:443
		_go_fuzz_dep_.CoverTab[72074]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:443
		return n > 8
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:443
		// _ = "end of CoverTab[72074]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:443
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:443
		_go_fuzz_dep_.CoverTab[72075]++
												panic("bad n")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:444
		// _ = "end of CoverTab[72075]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:445
		_go_fuzz_dep_.CoverTab[72076]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:445
		// _ = "end of CoverTab[72076]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:445
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:445
	// _ = "end of CoverTab[72068]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:445
	_go_fuzz_dep_.CoverTab[72069]++
											if len(p) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:446
		_go_fuzz_dep_.CoverTab[72077]++
												return 0, p, errNeedMore
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:447
		// _ = "end of CoverTab[72077]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:448
		_go_fuzz_dep_.CoverTab[72078]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:448
		// _ = "end of CoverTab[72078]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:448
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:448
	// _ = "end of CoverTab[72069]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:448
	_go_fuzz_dep_.CoverTab[72070]++
											i = uint64(p[0])
											if n < 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:450
		_go_fuzz_dep_.CoverTab[72079]++
												i &= (1 << uint64(n)) - 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:451
		// _ = "end of CoverTab[72079]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:452
		_go_fuzz_dep_.CoverTab[72080]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:452
		// _ = "end of CoverTab[72080]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:452
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:452
	// _ = "end of CoverTab[72070]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:452
	_go_fuzz_dep_.CoverTab[72071]++
											if i < (1<<uint64(n))-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:453
		_go_fuzz_dep_.CoverTab[72081]++
												return i, p[1:], nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:454
		// _ = "end of CoverTab[72081]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:455
		_go_fuzz_dep_.CoverTab[72082]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:455
		// _ = "end of CoverTab[72082]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:455
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:455
	// _ = "end of CoverTab[72071]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:455
	_go_fuzz_dep_.CoverTab[72072]++

											origP := p
											p = p[1:]
											var m uint64
											for len(p) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:460
		_go_fuzz_dep_.CoverTab[72083]++
												b := p[0]
												p = p[1:]
												i += uint64(b&127) << m
												if b&128 == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:464
			_go_fuzz_dep_.CoverTab[72085]++
													return i, p, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:465
			// _ = "end of CoverTab[72085]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:466
			_go_fuzz_dep_.CoverTab[72086]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:466
			// _ = "end of CoverTab[72086]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:466
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:466
		// _ = "end of CoverTab[72083]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:466
		_go_fuzz_dep_.CoverTab[72084]++
												m += 7
												if m >= 63 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:468
			_go_fuzz_dep_.CoverTab[72087]++
													return 0, origP, errVarintOverflow
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:469
			// _ = "end of CoverTab[72087]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:470
			_go_fuzz_dep_.CoverTab[72088]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:470
			// _ = "end of CoverTab[72088]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:470
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:470
		// _ = "end of CoverTab[72084]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:471
	// _ = "end of CoverTab[72072]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:471
	_go_fuzz_dep_.CoverTab[72073]++
											return 0, origP, errNeedMore
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:472
	// _ = "end of CoverTab[72073]"
}

// readString reads an hpack string from p.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:475
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:475
// It returns a reference to the encoded string data to permit deferring decode costs
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:475
// until after the caller verifies all data is present.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:479
func (d *Decoder) readString(p []byte) (u undecodedString, remain []byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:479
	_go_fuzz_dep_.CoverTab[72089]++
											if len(p) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:480
		_go_fuzz_dep_.CoverTab[72094]++
												return u, p, errNeedMore
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:481
		// _ = "end of CoverTab[72094]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:482
		_go_fuzz_dep_.CoverTab[72095]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:482
		// _ = "end of CoverTab[72095]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:482
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:482
	// _ = "end of CoverTab[72089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:482
	_go_fuzz_dep_.CoverTab[72090]++
											isHuff := p[0]&128 != 0
											strLen, p, err := readVarInt(7, p)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:485
		_go_fuzz_dep_.CoverTab[72096]++
												return u, p, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:486
		// _ = "end of CoverTab[72096]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:487
		_go_fuzz_dep_.CoverTab[72097]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:487
		// _ = "end of CoverTab[72097]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:487
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:487
	// _ = "end of CoverTab[72090]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:487
	_go_fuzz_dep_.CoverTab[72091]++
											if d.maxStrLen != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:488
		_go_fuzz_dep_.CoverTab[72098]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:488
		return strLen > uint64(d.maxStrLen)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:488
		// _ = "end of CoverTab[72098]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:488
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:488
		_go_fuzz_dep_.CoverTab[72099]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:494
		return u, nil, ErrStringLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:494
		// _ = "end of CoverTab[72099]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:495
		_go_fuzz_dep_.CoverTab[72100]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:495
		// _ = "end of CoverTab[72100]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:495
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:495
	// _ = "end of CoverTab[72091]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:495
	_go_fuzz_dep_.CoverTab[72092]++
											if uint64(len(p)) < strLen {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:496
		_go_fuzz_dep_.CoverTab[72101]++
												return u, p, errNeedMore
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:497
		// _ = "end of CoverTab[72101]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:498
		_go_fuzz_dep_.CoverTab[72102]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:498
		// _ = "end of CoverTab[72102]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:498
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:498
	// _ = "end of CoverTab[72092]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:498
	_go_fuzz_dep_.CoverTab[72093]++
											u.isHuff = isHuff
											u.b = p[:strLen]
											return u, p[strLen:], nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:501
	// _ = "end of CoverTab[72093]"
}

type undecodedString struct {
	isHuff	bool
	b	[]byte
}

func (d *Decoder) decodeString(u undecodedString) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:509
	_go_fuzz_dep_.CoverTab[72103]++
											if !u.isHuff {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:510
		_go_fuzz_dep_.CoverTab[72106]++
												return string(u.b), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:511
		// _ = "end of CoverTab[72106]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:512
		_go_fuzz_dep_.CoverTab[72107]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:512
		// _ = "end of CoverTab[72107]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:512
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:512
	// _ = "end of CoverTab[72103]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:512
	_go_fuzz_dep_.CoverTab[72104]++
											buf := bufPool.Get().(*bytes.Buffer)
											buf.Reset()
											var s string
											err := huffmanDecode(buf, d.maxStrLen, u.b)
											if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:517
		_go_fuzz_dep_.CoverTab[72108]++
												s = buf.String()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:518
		// _ = "end of CoverTab[72108]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:519
		_go_fuzz_dep_.CoverTab[72109]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:519
		// _ = "end of CoverTab[72109]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:519
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:519
	// _ = "end of CoverTab[72104]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:519
	_go_fuzz_dep_.CoverTab[72105]++
											buf.Reset()
											bufPool.Put(buf)
											return s, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:522
	// _ = "end of CoverTab[72105]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:523
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/hpack/hpack.go:523
var _ = _go_fuzz_dep_.CoverTab
