// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:5
// Package hpack implements HPACK, a compression format for
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:5
// efficiently representing HTTP header fields in the context of HTTP/2.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:5
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:5
// See http://tools.ietf.org/html/draft-ietf-httpbis-header-compression-09
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
package hpack

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:9
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
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:22
	_go_fuzz_dep_.CoverTab[35153]++
										return fmt.Sprintf("decoding error: %v", de.Err)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:23
	// _ = "end of CoverTab[35153]"
}

// An InvalidIndexError is returned when an encoder references a table
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:26
// entry before the static table or after the end of the dynamic table.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:28
type InvalidIndexError int

func (e InvalidIndexError) Error() string {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:30
	_go_fuzz_dep_.CoverTab[35154]++
										return fmt.Sprintf("invalid indexed representation index %d", int(e))
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:31
	// _ = "end of CoverTab[35154]"
}

// A HeaderField is a name-value pair. Both the name and value are
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:34
// treated as opaque sequences of octets.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:36
type HeaderField struct {
	Name, Value	string

	// Sensitive means that this header field should never be
	// indexed.
	Sensitive	bool
}

// IsPseudo reports whether the header field is an http2 pseudo header.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:44
// That is, it reports whether it starts with a colon.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:44
// It is not otherwise guaranteed to be a valid pseudo header field,
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:44
// though.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:48
func (hf HeaderField) IsPseudo() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:48
	_go_fuzz_dep_.CoverTab[35155]++
										return len(hf.Name) != 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:49
		_go_fuzz_dep_.CoverTab[35156]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:49
		return hf.Name[0] == ':'
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:49
		// _ = "end of CoverTab[35156]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:49
	}()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:49
	// _ = "end of CoverTab[35155]"
}

func (hf HeaderField) String() string {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:52
	_go_fuzz_dep_.CoverTab[35157]++
										var suffix string
										if hf.Sensitive {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:54
		_go_fuzz_dep_.CoverTab[35159]++
											suffix = " (sensitive)"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:55
		// _ = "end of CoverTab[35159]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:56
		_go_fuzz_dep_.CoverTab[35160]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:56
		// _ = "end of CoverTab[35160]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:56
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:56
	// _ = "end of CoverTab[35157]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:56
	_go_fuzz_dep_.CoverTab[35158]++
										return fmt.Sprintf("header field %q = %q%s", hf.Name, hf.Value, suffix)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:57
	// _ = "end of CoverTab[35158]"
}

// Size returns the size of an entry per RFC 7541 section 4.1.
func (hf HeaderField) Size() uint32 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:61
	_go_fuzz_dep_.CoverTab[35161]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:74
	return uint32(len(hf.Name) + len(hf.Value) + 32)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:74
	// _ = "end of CoverTab[35161]"
}

// A Decoder is the decoding context for incremental processing of
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:77
// header blocks.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:79
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
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:99
// table size. The emitFunc will be called for each valid field
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:99
// parsed, in the same goroutine as calls to Write, before Write returns.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:102
func NewDecoder(maxDynamicTableSize uint32, emitFunc func(f HeaderField)) *Decoder {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:102
	_go_fuzz_dep_.CoverTab[35162]++
										d := &Decoder{
		emit:		emitFunc,
		emitEnabled:	true,
		firstField:	true,
	}
										d.dynTab.table.init()
										d.dynTab.allowedMaxSize = maxDynamicTableSize
										d.dynTab.setMaxSize(maxDynamicTableSize)
										return d
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:111
	// _ = "end of CoverTab[35162]"
}

// ErrStringLength is returned by Decoder.Write when the max string length
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:114
// (as configured by Decoder.SetMaxStringLength) would be violated.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:116
var ErrStringLength = errors.New("hpack: string too long")

// SetMaxStringLength sets the maximum size of a HeaderField name or
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:118
// value string. If a string exceeds this length (even after any
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:118
// decompression), Write will return ErrStringLength.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:118
// A value of 0 means unlimited and is the default from NewDecoder.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:122
func (d *Decoder) SetMaxStringLength(n int) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:122
	_go_fuzz_dep_.CoverTab[35163]++
										d.maxStrLen = n
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:123
	// _ = "end of CoverTab[35163]"
}

// SetEmitFunc changes the callback used when new header fields
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:126
// are decoded.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:126
// It must be non-nil. It does not affect EmitEnabled.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:129
func (d *Decoder) SetEmitFunc(emitFunc func(f HeaderField)) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:129
	_go_fuzz_dep_.CoverTab[35164]++
										d.emit = emitFunc
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:130
	// _ = "end of CoverTab[35164]"
}

// SetEmitEnabled controls whether the emitFunc provided to NewDecoder
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:133
// should be called. The default is true.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:133
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:133
// This facility exists to let servers enforce MAX_HEADER_LIST_SIZE
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:133
// while still decoding and keeping in-sync with decoder state, but
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:133
// without doing unnecessary decompression or generating unnecessary
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:133
// garbage for header fields past the limit.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:140
func (d *Decoder) SetEmitEnabled(v bool) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:140
	_go_fuzz_dep_.CoverTab[35165]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:140
	d.emitEnabled = v
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:140
	// _ = "end of CoverTab[35165]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:140
}

// EmitEnabled reports whether calls to the emitFunc provided to NewDecoder
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:142
// are currently enabled. The default is true.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:144
func (d *Decoder) EmitEnabled() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:144
	_go_fuzz_dep_.CoverTab[35166]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:144
	return d.emitEnabled
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:144
	// _ = "end of CoverTab[35166]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:144
}

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:149
func (d *Decoder) SetMaxDynamicTableSize(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:149
	_go_fuzz_dep_.CoverTab[35167]++
										d.dynTab.setMaxSize(v)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:150
	// _ = "end of CoverTab[35167]"
}

// SetAllowedMaxDynamicTableSize sets the upper bound that the encoded
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:153
// stream (via dynamic table size updates) may set the maximum size
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:153
// to.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:156
func (d *Decoder) SetAllowedMaxDynamicTableSize(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:156
	_go_fuzz_dep_.CoverTab[35168]++
										d.dynTab.allowedMaxSize = v
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:157
	// _ = "end of CoverTab[35168]"
}

type dynamicTable struct {
	// https://httpwg.org/specs/rfc7541.html#rfc.section.2.3.2
	table		headerFieldTable
	size		uint32	// in bytes
	maxSize		uint32	// current maxSize
	allowedMaxSize	uint32	// maxSize may go up to this, inclusive
}

func (dt *dynamicTable) setMaxSize(v uint32) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:168
	_go_fuzz_dep_.CoverTab[35169]++
										dt.maxSize = v
										dt.evict()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:170
	// _ = "end of CoverTab[35169]"
}

func (dt *dynamicTable) add(f HeaderField) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:173
	_go_fuzz_dep_.CoverTab[35170]++
										dt.table.addEntry(f)
										dt.size += f.Size()
										dt.evict()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:176
	// _ = "end of CoverTab[35170]"
}

// If we're too big, evict old stuff.
func (dt *dynamicTable) evict() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:180
	_go_fuzz_dep_.CoverTab[35171]++
										var n int
										for dt.size > dt.maxSize && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:182
		_go_fuzz_dep_.CoverTab[35173]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:182
		return n < dt.table.len()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:182
		// _ = "end of CoverTab[35173]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:182
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:182
		_go_fuzz_dep_.CoverTab[35174]++
											dt.size -= dt.table.ents[n].Size()
											n++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:184
		// _ = "end of CoverTab[35174]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:185
	// _ = "end of CoverTab[35171]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:185
	_go_fuzz_dep_.CoverTab[35172]++
										dt.table.evictOldest(n)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:186
	// _ = "end of CoverTab[35172]"
}

func (d *Decoder) maxTableIndex() int {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:189
	_go_fuzz_dep_.CoverTab[35175]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:193
	return d.dynTab.table.len() + staticTable.len()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:193
	// _ = "end of CoverTab[35175]"
}

func (d *Decoder) at(i uint64) (hf HeaderField, ok bool) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:196
	_go_fuzz_dep_.CoverTab[35176]++

										if i == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:198
		_go_fuzz_dep_.CoverTab[35180]++
											return
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:199
		// _ = "end of CoverTab[35180]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:200
		_go_fuzz_dep_.CoverTab[35181]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:200
		// _ = "end of CoverTab[35181]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:200
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:200
	// _ = "end of CoverTab[35176]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:200
	_go_fuzz_dep_.CoverTab[35177]++
										if i <= uint64(staticTable.len()) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:201
		_go_fuzz_dep_.CoverTab[35182]++
											return staticTable.ents[i-1], true
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:202
		// _ = "end of CoverTab[35182]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:203
		_go_fuzz_dep_.CoverTab[35183]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:203
		// _ = "end of CoverTab[35183]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:203
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:203
	// _ = "end of CoverTab[35177]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:203
	_go_fuzz_dep_.CoverTab[35178]++
										if i > uint64(d.maxTableIndex()) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:204
		_go_fuzz_dep_.CoverTab[35184]++
											return
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:205
		// _ = "end of CoverTab[35184]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:206
		_go_fuzz_dep_.CoverTab[35185]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:206
		// _ = "end of CoverTab[35185]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:206
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:206
	// _ = "end of CoverTab[35178]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:206
	_go_fuzz_dep_.CoverTab[35179]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:210
	dt := d.dynTab.table
										return dt.ents[dt.len()-(int(i)-staticTable.len())], true
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:211
	// _ = "end of CoverTab[35179]"
}

// Decode decodes an entire block.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:214
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:214
// TODO: remove this method and make it incremental later? This is
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:214
// easier for debugging now.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:218
func (d *Decoder) DecodeFull(p []byte) ([]HeaderField, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:218
	_go_fuzz_dep_.CoverTab[35186]++
										var hf []HeaderField
										saveFunc := d.emit
										defer func() { _go_fuzz_dep_.CoverTab[35191]++; d.emit = saveFunc; // _ = "end of CoverTab[35191]" }()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:221
	// _ = "end of CoverTab[35186]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:221
	_go_fuzz_dep_.CoverTab[35187]++
										d.emit = func(f HeaderField) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:222
		_go_fuzz_dep_.CoverTab[35192]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:222
		hf = append(hf, f)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:222
		// _ = "end of CoverTab[35192]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:222
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:222
	// _ = "end of CoverTab[35187]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:222
	_go_fuzz_dep_.CoverTab[35188]++
										if _, err := d.Write(p); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:223
		_go_fuzz_dep_.CoverTab[35193]++
											return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:224
		// _ = "end of CoverTab[35193]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:225
		_go_fuzz_dep_.CoverTab[35194]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:225
		// _ = "end of CoverTab[35194]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:225
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:225
	// _ = "end of CoverTab[35188]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:225
	_go_fuzz_dep_.CoverTab[35189]++
										if err := d.Close(); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:226
		_go_fuzz_dep_.CoverTab[35195]++
											return nil, err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:227
		// _ = "end of CoverTab[35195]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:228
		_go_fuzz_dep_.CoverTab[35196]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:228
		// _ = "end of CoverTab[35196]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:228
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:228
	// _ = "end of CoverTab[35189]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:228
	_go_fuzz_dep_.CoverTab[35190]++
										return hf, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:229
	// _ = "end of CoverTab[35190]"
}

// Close declares that the decoding is complete and resets the Decoder
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:232
// to be reused again for a new header block. If there is any remaining
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:232
// data in the decoder's buffer, Close returns an error.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:235
func (d *Decoder) Close() error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:235
	_go_fuzz_dep_.CoverTab[35197]++
										if d.saveBuf.Len() > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:236
		_go_fuzz_dep_.CoverTab[35199]++
											d.saveBuf.Reset()
											return DecodingError{errors.New("truncated headers")}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:238
		// _ = "end of CoverTab[35199]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:239
		_go_fuzz_dep_.CoverTab[35200]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:239
		// _ = "end of CoverTab[35200]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:239
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:239
	// _ = "end of CoverTab[35197]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:239
	_go_fuzz_dep_.CoverTab[35198]++
										d.firstField = true
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:241
	// _ = "end of CoverTab[35198]"
}

func (d *Decoder) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:244
	_go_fuzz_dep_.CoverTab[35201]++
										if len(p) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:245
		_go_fuzz_dep_.CoverTab[35205]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:249
		return
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:249
		// _ = "end of CoverTab[35205]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:250
		_go_fuzz_dep_.CoverTab[35206]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:250
		// _ = "end of CoverTab[35206]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:250
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:250
	// _ = "end of CoverTab[35201]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:250
	_go_fuzz_dep_.CoverTab[35202]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:253
	if d.saveBuf.Len() == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:253
		_go_fuzz_dep_.CoverTab[35207]++
											d.buf = p
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:254
		// _ = "end of CoverTab[35207]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:255
		_go_fuzz_dep_.CoverTab[35208]++
											d.saveBuf.Write(p)
											d.buf = d.saveBuf.Bytes()
											d.saveBuf.Reset()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:258
		// _ = "end of CoverTab[35208]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:259
	// _ = "end of CoverTab[35202]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:259
	_go_fuzz_dep_.CoverTab[35203]++

										for len(d.buf) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:261
		_go_fuzz_dep_.CoverTab[35209]++
											err = d.parseHeaderFieldRepr()
											if err == errNeedMore {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:263
			_go_fuzz_dep_.CoverTab[35211]++
			// Extra paranoia, making sure saveBuf won't
			// get too large. All the varint and string
			// reading code earlier should already catch
			// overlong things and return ErrStringLength,
			// but keep this as a last resort.
			const varIntOverhead = 8	// conservative
			if d.maxStrLen != 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:270
				_go_fuzz_dep_.CoverTab[35213]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:270
				return int64(len(d.buf)) > 2*(int64(d.maxStrLen)+varIntOverhead)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:270
				// _ = "end of CoverTab[35213]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:270
			}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:270
				_go_fuzz_dep_.CoverTab[35214]++
													return 0, ErrStringLength
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:271
				// _ = "end of CoverTab[35214]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:272
				_go_fuzz_dep_.CoverTab[35215]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:272
				// _ = "end of CoverTab[35215]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:272
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:272
			// _ = "end of CoverTab[35211]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:272
			_go_fuzz_dep_.CoverTab[35212]++
												d.saveBuf.Write(d.buf)
												return len(p), nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:274
			// _ = "end of CoverTab[35212]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:275
			_go_fuzz_dep_.CoverTab[35216]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:275
			// _ = "end of CoverTab[35216]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:275
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:275
		// _ = "end of CoverTab[35209]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:275
		_go_fuzz_dep_.CoverTab[35210]++
											d.firstField = false
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:277
			_go_fuzz_dep_.CoverTab[35217]++
												break
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:278
			// _ = "end of CoverTab[35217]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:279
			_go_fuzz_dep_.CoverTab[35218]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:279
			// _ = "end of CoverTab[35218]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:279
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:279
		// _ = "end of CoverTab[35210]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:280
	// _ = "end of CoverTab[35203]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:280
	_go_fuzz_dep_.CoverTab[35204]++
										return len(p), err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:281
	// _ = "end of CoverTab[35204]"
}

// errNeedMore is an internal sentinel error value that means the
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:284
// buffer is truncated and we need to read more data before we can
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:284
// continue parsing.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:287
var errNeedMore = errors.New("need more data")

type indexType int

const (
	indexedTrue	indexType	= iota
	indexedFalse
	indexedNever
)

func (v indexType) indexed() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:297
	_go_fuzz_dep_.CoverTab[35219]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:297
	return v == indexedTrue
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:297
	// _ = "end of CoverTab[35219]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:297
}
func (v indexType) sensitive() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:298
	_go_fuzz_dep_.CoverTab[35220]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:298
	return v == indexedNever
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:298
	// _ = "end of CoverTab[35220]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:298
}

// returns errNeedMore if there isn't enough data available.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:300
// any other error is fatal.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:300
// consumes d.buf iff it returns nil.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:300
// precondition: must be called with len(d.buf) > 0
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:304
func (d *Decoder) parseHeaderFieldRepr() error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:304
	_go_fuzz_dep_.CoverTab[35221]++
										b := d.buf[0]
										switch {
	case b&128 != 0:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:307
		_go_fuzz_dep_.CoverTab[35223]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:311
		return d.parseFieldIndexed()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:311
		// _ = "end of CoverTab[35223]"
	case b&192 == 64:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:312
		_go_fuzz_dep_.CoverTab[35224]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:316
		return d.parseFieldLiteral(6, indexedTrue)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:316
		// _ = "end of CoverTab[35224]"
	case b&240 == 0:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:317
		_go_fuzz_dep_.CoverTab[35225]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:321
		return d.parseFieldLiteral(4, indexedFalse)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:321
		// _ = "end of CoverTab[35225]"
	case b&240 == 16:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:322
		_go_fuzz_dep_.CoverTab[35226]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:326
		return d.parseFieldLiteral(4, indexedNever)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:326
		// _ = "end of CoverTab[35226]"
	case b&224 == 32:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:327
		_go_fuzz_dep_.CoverTab[35227]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:331
		return d.parseDynamicTableSizeUpdate()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:331
		// _ = "end of CoverTab[35227]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:331
	default:
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:331
		_go_fuzz_dep_.CoverTab[35228]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:331
		// _ = "end of CoverTab[35228]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:332
	// _ = "end of CoverTab[35221]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:332
	_go_fuzz_dep_.CoverTab[35222]++

										return DecodingError{errors.New("invalid encoding")}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:334
	// _ = "end of CoverTab[35222]"
}

// (same invariants and behavior as parseHeaderFieldRepr)
func (d *Decoder) parseFieldIndexed() error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:338
	_go_fuzz_dep_.CoverTab[35229]++
										buf := d.buf
										idx, buf, err := readVarInt(7, buf)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:341
		_go_fuzz_dep_.CoverTab[35232]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:342
		// _ = "end of CoverTab[35232]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:343
		_go_fuzz_dep_.CoverTab[35233]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:343
		// _ = "end of CoverTab[35233]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:343
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:343
	// _ = "end of CoverTab[35229]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:343
	_go_fuzz_dep_.CoverTab[35230]++
										hf, ok := d.at(idx)
										if !ok {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:345
		_go_fuzz_dep_.CoverTab[35234]++
											return DecodingError{InvalidIndexError(idx)}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:346
		// _ = "end of CoverTab[35234]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:347
		_go_fuzz_dep_.CoverTab[35235]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:347
		// _ = "end of CoverTab[35235]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:347
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:347
	// _ = "end of CoverTab[35230]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:347
	_go_fuzz_dep_.CoverTab[35231]++
										d.buf = buf
										return d.callEmit(HeaderField{Name: hf.Name, Value: hf.Value})
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:349
	// _ = "end of CoverTab[35231]"
}

// (same invariants and behavior as parseHeaderFieldRepr)
func (d *Decoder) parseFieldLiteral(n uint8, it indexType) error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:353
	_go_fuzz_dep_.CoverTab[35236]++
										buf := d.buf
										nameIdx, buf, err := readVarInt(n, buf)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:356
		_go_fuzz_dep_.CoverTab[35242]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:357
		// _ = "end of CoverTab[35242]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:358
		_go_fuzz_dep_.CoverTab[35243]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:358
		// _ = "end of CoverTab[35243]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:358
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:358
	// _ = "end of CoverTab[35236]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:358
	_go_fuzz_dep_.CoverTab[35237]++

										var hf HeaderField
										wantStr := d.emitEnabled || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:361
		_go_fuzz_dep_.CoverTab[35244]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:361
		return it.indexed()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:361
		// _ = "end of CoverTab[35244]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:361
	}()
										var undecodedName undecodedString
										if nameIdx > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:363
		_go_fuzz_dep_.CoverTab[35245]++
											ihf, ok := d.at(nameIdx)
											if !ok {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:365
			_go_fuzz_dep_.CoverTab[35247]++
												return DecodingError{InvalidIndexError(nameIdx)}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:366
			// _ = "end of CoverTab[35247]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:367
			_go_fuzz_dep_.CoverTab[35248]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:367
			// _ = "end of CoverTab[35248]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:367
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:367
		// _ = "end of CoverTab[35245]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:367
		_go_fuzz_dep_.CoverTab[35246]++
											hf.Name = ihf.Name
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:368
		// _ = "end of CoverTab[35246]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:369
		_go_fuzz_dep_.CoverTab[35249]++
											undecodedName, buf, err = d.readString(buf)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:371
			_go_fuzz_dep_.CoverTab[35250]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:372
			// _ = "end of CoverTab[35250]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:373
			_go_fuzz_dep_.CoverTab[35251]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:373
			// _ = "end of CoverTab[35251]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:373
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:373
		// _ = "end of CoverTab[35249]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:374
	// _ = "end of CoverTab[35237]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:374
	_go_fuzz_dep_.CoverTab[35238]++
										undecodedValue, buf, err := d.readString(buf)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:376
		_go_fuzz_dep_.CoverTab[35252]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:377
		// _ = "end of CoverTab[35252]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:378
		_go_fuzz_dep_.CoverTab[35253]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:378
		// _ = "end of CoverTab[35253]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:378
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:378
	// _ = "end of CoverTab[35238]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:378
	_go_fuzz_dep_.CoverTab[35239]++
										if wantStr {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:379
		_go_fuzz_dep_.CoverTab[35254]++
											if nameIdx <= 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:380
			_go_fuzz_dep_.CoverTab[35256]++
												hf.Name, err = d.decodeString(undecodedName)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:382
				_go_fuzz_dep_.CoverTab[35257]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:383
				// _ = "end of CoverTab[35257]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:384
				_go_fuzz_dep_.CoverTab[35258]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:384
				// _ = "end of CoverTab[35258]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:384
			}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:384
			// _ = "end of CoverTab[35256]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:385
			_go_fuzz_dep_.CoverTab[35259]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:385
			// _ = "end of CoverTab[35259]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:385
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:385
		// _ = "end of CoverTab[35254]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:385
		_go_fuzz_dep_.CoverTab[35255]++
											hf.Value, err = d.decodeString(undecodedValue)
											if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:387
			_go_fuzz_dep_.CoverTab[35260]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:388
			// _ = "end of CoverTab[35260]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:389
			_go_fuzz_dep_.CoverTab[35261]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:389
			// _ = "end of CoverTab[35261]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:389
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:389
		// _ = "end of CoverTab[35255]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:390
		_go_fuzz_dep_.CoverTab[35262]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:390
		// _ = "end of CoverTab[35262]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:390
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:390
	// _ = "end of CoverTab[35239]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:390
	_go_fuzz_dep_.CoverTab[35240]++
										d.buf = buf
										if it.indexed() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:392
		_go_fuzz_dep_.CoverTab[35263]++
											d.dynTab.add(hf)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:393
		// _ = "end of CoverTab[35263]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:394
		_go_fuzz_dep_.CoverTab[35264]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:394
		// _ = "end of CoverTab[35264]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:394
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:394
	// _ = "end of CoverTab[35240]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:394
	_go_fuzz_dep_.CoverTab[35241]++
										hf.Sensitive = it.sensitive()
										return d.callEmit(hf)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:396
	// _ = "end of CoverTab[35241]"
}

func (d *Decoder) callEmit(hf HeaderField) error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:399
	_go_fuzz_dep_.CoverTab[35265]++
										if d.maxStrLen != 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:400
		_go_fuzz_dep_.CoverTab[35268]++
											if len(hf.Name) > d.maxStrLen || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:401
			_go_fuzz_dep_.CoverTab[35269]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:401
			return len(hf.Value) > d.maxStrLen
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:401
			// _ = "end of CoverTab[35269]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:401
		}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:401
			_go_fuzz_dep_.CoverTab[35270]++
												return ErrStringLength
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:402
			// _ = "end of CoverTab[35270]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:403
			_go_fuzz_dep_.CoverTab[35271]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:403
			// _ = "end of CoverTab[35271]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:403
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:403
		// _ = "end of CoverTab[35268]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:404
		_go_fuzz_dep_.CoverTab[35272]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:404
		// _ = "end of CoverTab[35272]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:404
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:404
	// _ = "end of CoverTab[35265]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:404
	_go_fuzz_dep_.CoverTab[35266]++
										if d.emitEnabled {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:405
		_go_fuzz_dep_.CoverTab[35273]++
											d.emit(hf)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:406
		// _ = "end of CoverTab[35273]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:407
		_go_fuzz_dep_.CoverTab[35274]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:407
		// _ = "end of CoverTab[35274]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:407
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:407
	// _ = "end of CoverTab[35266]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:407
	_go_fuzz_dep_.CoverTab[35267]++
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:408
	// _ = "end of CoverTab[35267]"
}

// (same invariants and behavior as parseHeaderFieldRepr)
func (d *Decoder) parseDynamicTableSizeUpdate() error {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:412
	_go_fuzz_dep_.CoverTab[35275]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:415
	if !d.firstField && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:415
		_go_fuzz_dep_.CoverTab[35279]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:415
		return d.dynTab.size > 0
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:415
		// _ = "end of CoverTab[35279]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:415
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:415
		_go_fuzz_dep_.CoverTab[35280]++
											return DecodingError{errors.New("dynamic table size update MUST occur at the beginning of a header block")}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:416
		// _ = "end of CoverTab[35280]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:417
		_go_fuzz_dep_.CoverTab[35281]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:417
		// _ = "end of CoverTab[35281]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:417
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:417
	// _ = "end of CoverTab[35275]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:417
	_go_fuzz_dep_.CoverTab[35276]++

										buf := d.buf
										size, buf, err := readVarInt(5, buf)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:421
		_go_fuzz_dep_.CoverTab[35282]++
											return err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:422
		// _ = "end of CoverTab[35282]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:423
		_go_fuzz_dep_.CoverTab[35283]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:423
		// _ = "end of CoverTab[35283]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:423
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:423
	// _ = "end of CoverTab[35276]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:423
	_go_fuzz_dep_.CoverTab[35277]++
										if size > uint64(d.dynTab.allowedMaxSize) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:424
		_go_fuzz_dep_.CoverTab[35284]++
											return DecodingError{errors.New("dynamic table size update too large")}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:425
		// _ = "end of CoverTab[35284]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:426
		_go_fuzz_dep_.CoverTab[35285]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:426
		// _ = "end of CoverTab[35285]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:426
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:426
	// _ = "end of CoverTab[35277]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:426
	_go_fuzz_dep_.CoverTab[35278]++
										d.dynTab.setMaxSize(uint32(size))
										d.buf = buf
										return nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:429
	// _ = "end of CoverTab[35278]"
}

var errVarintOverflow = DecodingError{errors.New("varint integer overflow")}

// readVarInt reads an unsigned variable length integer off the
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
// beginning of p. n is the parameter as described in
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
// https://httpwg.org/specs/rfc7541.html#rfc.section.5.1.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
// n must always be between 1 and 8.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
// The returned remain buffer is either a smaller suffix of p, or err != nil.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:434
// The error is errNeedMore if p doesn't contain a complete integer.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:442
func readVarInt(n byte, p []byte) (i uint64, remain []byte, err error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:442
	_go_fuzz_dep_.CoverTab[35286]++
										if n < 1 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:443
		_go_fuzz_dep_.CoverTab[35292]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:443
		return n > 8
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:443
		// _ = "end of CoverTab[35292]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:443
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:443
		_go_fuzz_dep_.CoverTab[35293]++
											panic("bad n")
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:444
		// _ = "end of CoverTab[35293]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:445
		_go_fuzz_dep_.CoverTab[35294]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:445
		// _ = "end of CoverTab[35294]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:445
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:445
	// _ = "end of CoverTab[35286]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:445
	_go_fuzz_dep_.CoverTab[35287]++
										if len(p) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:446
		_go_fuzz_dep_.CoverTab[35295]++
											return 0, p, errNeedMore
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:447
		// _ = "end of CoverTab[35295]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:448
		_go_fuzz_dep_.CoverTab[35296]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:448
		// _ = "end of CoverTab[35296]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:448
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:448
	// _ = "end of CoverTab[35287]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:448
	_go_fuzz_dep_.CoverTab[35288]++
										i = uint64(p[0])
										if n < 8 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:450
		_go_fuzz_dep_.CoverTab[35297]++
											i &= (1 << uint64(n)) - 1
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:451
		// _ = "end of CoverTab[35297]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:452
		_go_fuzz_dep_.CoverTab[35298]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:452
		// _ = "end of CoverTab[35298]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:452
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:452
	// _ = "end of CoverTab[35288]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:452
	_go_fuzz_dep_.CoverTab[35289]++
										if i < (1<<uint64(n))-1 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:453
		_go_fuzz_dep_.CoverTab[35299]++
											return i, p[1:], nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:454
		// _ = "end of CoverTab[35299]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:455
		_go_fuzz_dep_.CoverTab[35300]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:455
		// _ = "end of CoverTab[35300]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:455
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:455
	// _ = "end of CoverTab[35289]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:455
	_go_fuzz_dep_.CoverTab[35290]++

										origP := p
										p = p[1:]
										var m uint64
										for len(p) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:460
		_go_fuzz_dep_.CoverTab[35301]++
											b := p[0]
											p = p[1:]
											i += uint64(b&127) << m
											if b&128 == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:464
			_go_fuzz_dep_.CoverTab[35303]++
												return i, p, nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:465
			// _ = "end of CoverTab[35303]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:466
			_go_fuzz_dep_.CoverTab[35304]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:466
			// _ = "end of CoverTab[35304]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:466
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:466
		// _ = "end of CoverTab[35301]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:466
		_go_fuzz_dep_.CoverTab[35302]++
											m += 7
											if m >= 63 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:468
			_go_fuzz_dep_.CoverTab[35305]++
												return 0, origP, errVarintOverflow
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:469
			// _ = "end of CoverTab[35305]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:470
			_go_fuzz_dep_.CoverTab[35306]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:470
			// _ = "end of CoverTab[35306]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:470
		}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:470
		// _ = "end of CoverTab[35302]"
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:471
	// _ = "end of CoverTab[35290]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:471
	_go_fuzz_dep_.CoverTab[35291]++
										return 0, origP, errNeedMore
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:472
	// _ = "end of CoverTab[35291]"
}

// readString reads an hpack string from p.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:475
//
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:475
// It returns a reference to the encoded string data to permit deferring decode costs
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:475
// until after the caller verifies all data is present.
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:479
func (d *Decoder) readString(p []byte) (u undecodedString, remain []byte, err error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:479
	_go_fuzz_dep_.CoverTab[35307]++
										if len(p) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:480
		_go_fuzz_dep_.CoverTab[35312]++
											return u, p, errNeedMore
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:481
		// _ = "end of CoverTab[35312]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:482
		_go_fuzz_dep_.CoverTab[35313]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:482
		// _ = "end of CoverTab[35313]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:482
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:482
	// _ = "end of CoverTab[35307]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:482
	_go_fuzz_dep_.CoverTab[35308]++
										isHuff := p[0]&128 != 0
										strLen, p, err := readVarInt(7, p)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:485
		_go_fuzz_dep_.CoverTab[35314]++
											return u, p, err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:486
		// _ = "end of CoverTab[35314]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:487
		_go_fuzz_dep_.CoverTab[35315]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:487
		// _ = "end of CoverTab[35315]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:487
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:487
	// _ = "end of CoverTab[35308]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:487
	_go_fuzz_dep_.CoverTab[35309]++
										if d.maxStrLen != 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:488
		_go_fuzz_dep_.CoverTab[35316]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:488
		return strLen > uint64(d.maxStrLen)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:488
		// _ = "end of CoverTab[35316]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:488
	}() {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:488
		_go_fuzz_dep_.CoverTab[35317]++

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:494
		return u, nil, ErrStringLength
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:494
		// _ = "end of CoverTab[35317]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:495
		_go_fuzz_dep_.CoverTab[35318]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:495
		// _ = "end of CoverTab[35318]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:495
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:495
	// _ = "end of CoverTab[35309]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:495
	_go_fuzz_dep_.CoverTab[35310]++
										if uint64(len(p)) < strLen {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:496
		_go_fuzz_dep_.CoverTab[35319]++
											return u, p, errNeedMore
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:497
		// _ = "end of CoverTab[35319]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:498
		_go_fuzz_dep_.CoverTab[35320]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:498
		// _ = "end of CoverTab[35320]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:498
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:498
	// _ = "end of CoverTab[35310]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:498
	_go_fuzz_dep_.CoverTab[35311]++
										u.isHuff = isHuff
										u.b = p[:strLen]
										return u, p[strLen:], nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:501
	// _ = "end of CoverTab[35311]"
}

type undecodedString struct {
	isHuff	bool
	b	[]byte
}

func (d *Decoder) decodeString(u undecodedString) (string, error) {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:509
	_go_fuzz_dep_.CoverTab[35321]++
										if !u.isHuff {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:510
		_go_fuzz_dep_.CoverTab[35324]++
											return string(u.b), nil
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:511
		// _ = "end of CoverTab[35324]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:512
		_go_fuzz_dep_.CoverTab[35325]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:512
		// _ = "end of CoverTab[35325]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:512
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:512
	// _ = "end of CoverTab[35321]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:512
	_go_fuzz_dep_.CoverTab[35322]++
										buf := bufPool.Get().(*bytes.Buffer)
										buf.Reset()
										var s string
										err := huffmanDecode(buf, d.maxStrLen, u.b)
										if err == nil {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:517
		_go_fuzz_dep_.CoverTab[35326]++
											s = buf.String()
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:518
		// _ = "end of CoverTab[35326]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:519
		_go_fuzz_dep_.CoverTab[35327]++
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:519
		// _ = "end of CoverTab[35327]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:519
	}
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:519
	// _ = "end of CoverTab[35322]"
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:519
	_go_fuzz_dep_.CoverTab[35323]++
										buf.Reset()
										bufPool.Put(buf)
										return s, err
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:522
	// _ = "end of CoverTab[35323]"
}

//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:523
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/net/http2/hpack/hpack.go:523
var _ = _go_fuzz_dep_.CoverTab
