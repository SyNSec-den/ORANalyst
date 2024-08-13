// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/dict_decoder.go:5
package flate

//line /usr/local/go/src/compress/flate/dict_decoder.go:5
import (
//line /usr/local/go/src/compress/flate/dict_decoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/dict_decoder.go:5
)
//line /usr/local/go/src/compress/flate/dict_decoder.go:5
import (
//line /usr/local/go/src/compress/flate/dict_decoder.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/dict_decoder.go:5
)

// dictDecoder implements the LZ77 sliding dictionary as used in decompression.
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
// LZ77 decompresses data through sequences of two forms of commands:
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//   - Literal insertions: Runs of one or more symbols are inserted into the data
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     stream as is. This is accomplished through the writeByte method for a
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     single symbol, or combinations of writeSlice/writeMark for multiple symbols.
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     Any valid stream must start with a literal insertion if no preset dictionary
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     is used.
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//   - Backward copies: Runs of one or more symbols are copied from previously
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     emitted data. Backward copies come as the tuple (dist, length) where dist
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     determines how far back in the stream to copy from and length determines how
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     many bytes to copy. Note that it is valid for the length to be greater than
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     the distance. Since LZ77 uses forward copies, that situation is used to
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     perform a form of run-length encoding on repeated runs of symbols.
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//     The writeCopy and tryWriteCopy are used to implement this command.
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
// For performance reasons, this implementation performs little to no sanity
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
// checks about the arguments. As such, the invariants documented for each
//line /usr/local/go/src/compress/flate/dict_decoder.go:7
// method call must be respected.
//line /usr/local/go/src/compress/flate/dict_decoder.go:27
type dictDecoder struct {
	hist	[]byte	// Sliding window history

	// Invariant: 0 <= rdPos <= wrPos <= len(hist)
	wrPos	int	// Current output position in buffer
	rdPos	int	// Have emitted hist[:rdPos] already
	full	bool	// Has a full window length been written yet?
}

// init initializes dictDecoder to have a sliding window dictionary of the given
//line /usr/local/go/src/compress/flate/dict_decoder.go:36
// size. If a preset dict is provided, it will initialize the dictionary with
//line /usr/local/go/src/compress/flate/dict_decoder.go:36
// the contents of dict.
//line /usr/local/go/src/compress/flate/dict_decoder.go:39
func (dd *dictDecoder) init(size int, dict []byte) {
	*dd = dictDecoder{hist: dd.hist}

	if cap(dd.hist) < size {
		dd.hist = make([]byte, size)
	}
	dd.hist = dd.hist[:size]

	if len(dict) > len(dd.hist) {
		dict = dict[len(dict)-len(dd.hist):]
	}
	dd.wrPos = copy(dd.hist, dict)
	if dd.wrPos == len(dd.hist) {
		dd.wrPos = 0
		dd.full = true
	}
	dd.rdPos = dd.wrPos
}

// histSize reports the total amount of historical data in the dictionary.
func (dd *dictDecoder) histSize() int {
//line /usr/local/go/src/compress/flate/dict_decoder.go:59
	_go_fuzz_dep_.CoverTab[26002]++
								if dd.full {
//line /usr/local/go/src/compress/flate/dict_decoder.go:60
		_go_fuzz_dep_.CoverTab[26004]++
									return len(dd.hist)
//line /usr/local/go/src/compress/flate/dict_decoder.go:61
		// _ = "end of CoverTab[26004]"
	} else {
//line /usr/local/go/src/compress/flate/dict_decoder.go:62
		_go_fuzz_dep_.CoverTab[26005]++
//line /usr/local/go/src/compress/flate/dict_decoder.go:62
		// _ = "end of CoverTab[26005]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:62
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:62
	// _ = "end of CoverTab[26002]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:62
	_go_fuzz_dep_.CoverTab[26003]++
								return dd.wrPos
//line /usr/local/go/src/compress/flate/dict_decoder.go:63
	// _ = "end of CoverTab[26003]"
}

// availRead reports the number of bytes that can be flushed by readFlush.
func (dd *dictDecoder) availRead() int {
//line /usr/local/go/src/compress/flate/dict_decoder.go:67
	_go_fuzz_dep_.CoverTab[26006]++
								return dd.wrPos - dd.rdPos
//line /usr/local/go/src/compress/flate/dict_decoder.go:68
	// _ = "end of CoverTab[26006]"
}

// availWrite reports the available amount of output buffer space.
func (dd *dictDecoder) availWrite() int {
//line /usr/local/go/src/compress/flate/dict_decoder.go:72
	_go_fuzz_dep_.CoverTab[26007]++
								return len(dd.hist) - dd.wrPos
//line /usr/local/go/src/compress/flate/dict_decoder.go:73
	// _ = "end of CoverTab[26007]"
}

// writeSlice returns a slice of the available buffer to write data to.
//line /usr/local/go/src/compress/flate/dict_decoder.go:76
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:76
// This invariant will be kept: len(s) <= availWrite()
//line /usr/local/go/src/compress/flate/dict_decoder.go:79
func (dd *dictDecoder) writeSlice() []byte {
//line /usr/local/go/src/compress/flate/dict_decoder.go:79
	_go_fuzz_dep_.CoverTab[26008]++
								return dd.hist[dd.wrPos:]
//line /usr/local/go/src/compress/flate/dict_decoder.go:80
	// _ = "end of CoverTab[26008]"
}

// writeMark advances the writer pointer by cnt.
//line /usr/local/go/src/compress/flate/dict_decoder.go:83
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:83
// This invariant must be kept: 0 <= cnt <= availWrite()
//line /usr/local/go/src/compress/flate/dict_decoder.go:86
func (dd *dictDecoder) writeMark(cnt int) {
//line /usr/local/go/src/compress/flate/dict_decoder.go:86
	_go_fuzz_dep_.CoverTab[26009]++
								dd.wrPos += cnt
//line /usr/local/go/src/compress/flate/dict_decoder.go:87
	// _ = "end of CoverTab[26009]"
}

// writeByte writes a single byte to the dictionary.
//line /usr/local/go/src/compress/flate/dict_decoder.go:90
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:90
// This invariant must be kept: 0 < availWrite()
//line /usr/local/go/src/compress/flate/dict_decoder.go:93
func (dd *dictDecoder) writeByte(c byte) {
//line /usr/local/go/src/compress/flate/dict_decoder.go:93
	_go_fuzz_dep_.CoverTab[26010]++
								dd.hist[dd.wrPos] = c
								dd.wrPos++
//line /usr/local/go/src/compress/flate/dict_decoder.go:95
	// _ = "end of CoverTab[26010]"
}

// writeCopy copies a string at a given (dist, length) to the output.
//line /usr/local/go/src/compress/flate/dict_decoder.go:98
// This returns the number of bytes copied and may be less than the requested
//line /usr/local/go/src/compress/flate/dict_decoder.go:98
// length if the available space in the output buffer is too small.
//line /usr/local/go/src/compress/flate/dict_decoder.go:98
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:98
// This invariant must be kept: 0 < dist <= histSize()
//line /usr/local/go/src/compress/flate/dict_decoder.go:103
func (dd *dictDecoder) writeCopy(dist, length int) int {
//line /usr/local/go/src/compress/flate/dict_decoder.go:103
	_go_fuzz_dep_.CoverTab[26011]++
								dstBase := dd.wrPos
								dstPos := dstBase
								srcPos := dstPos - dist
								endPos := dstPos + length
								if endPos > len(dd.hist) {
//line /usr/local/go/src/compress/flate/dict_decoder.go:108
		_go_fuzz_dep_.CoverTab[26015]++
									endPos = len(dd.hist)
//line /usr/local/go/src/compress/flate/dict_decoder.go:109
		// _ = "end of CoverTab[26015]"
	} else {
//line /usr/local/go/src/compress/flate/dict_decoder.go:110
		_go_fuzz_dep_.CoverTab[26016]++
//line /usr/local/go/src/compress/flate/dict_decoder.go:110
		// _ = "end of CoverTab[26016]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:110
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:110
	// _ = "end of CoverTab[26011]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:110
	_go_fuzz_dep_.CoverTab[26012]++

//line /usr/local/go/src/compress/flate/dict_decoder.go:119
	if srcPos < 0 {
//line /usr/local/go/src/compress/flate/dict_decoder.go:119
		_go_fuzz_dep_.CoverTab[26017]++
									srcPos += len(dd.hist)
									dstPos += copy(dd.hist[dstPos:endPos], dd.hist[srcPos:])
									srcPos = 0
//line /usr/local/go/src/compress/flate/dict_decoder.go:122
		// _ = "end of CoverTab[26017]"
	} else {
//line /usr/local/go/src/compress/flate/dict_decoder.go:123
		_go_fuzz_dep_.CoverTab[26018]++
//line /usr/local/go/src/compress/flate/dict_decoder.go:123
		// _ = "end of CoverTab[26018]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:123
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:123
	// _ = "end of CoverTab[26012]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:123
	_go_fuzz_dep_.CoverTab[26013]++

//line /usr/local/go/src/compress/flate/dict_decoder.go:139
	for dstPos < endPos {
//line /usr/local/go/src/compress/flate/dict_decoder.go:139
		_go_fuzz_dep_.CoverTab[26019]++
									dstPos += copy(dd.hist[dstPos:endPos], dd.hist[srcPos:dstPos])
//line /usr/local/go/src/compress/flate/dict_decoder.go:140
		// _ = "end of CoverTab[26019]"
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:141
	// _ = "end of CoverTab[26013]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:141
	_go_fuzz_dep_.CoverTab[26014]++

								dd.wrPos = dstPos
								return dstPos - dstBase
//line /usr/local/go/src/compress/flate/dict_decoder.go:144
	// _ = "end of CoverTab[26014]"
}

// tryWriteCopy tries to copy a string at a given (distance, length) to the
//line /usr/local/go/src/compress/flate/dict_decoder.go:147
// output. This specialized version is optimized for short distances.
//line /usr/local/go/src/compress/flate/dict_decoder.go:147
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:147
// This method is designed to be inlined for performance reasons.
//line /usr/local/go/src/compress/flate/dict_decoder.go:147
//
//line /usr/local/go/src/compress/flate/dict_decoder.go:147
// This invariant must be kept: 0 < dist <= histSize()
//line /usr/local/go/src/compress/flate/dict_decoder.go:153
func (dd *dictDecoder) tryWriteCopy(dist, length int) int {
//line /usr/local/go/src/compress/flate/dict_decoder.go:153
	_go_fuzz_dep_.CoverTab[26020]++
								dstPos := dd.wrPos
								endPos := dstPos + length
								if dstPos < dist || func() bool {
//line /usr/local/go/src/compress/flate/dict_decoder.go:156
		_go_fuzz_dep_.CoverTab[26023]++
//line /usr/local/go/src/compress/flate/dict_decoder.go:156
		return endPos > len(dd.hist)
//line /usr/local/go/src/compress/flate/dict_decoder.go:156
		// _ = "end of CoverTab[26023]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:156
	}() {
//line /usr/local/go/src/compress/flate/dict_decoder.go:156
		_go_fuzz_dep_.CoverTab[26024]++
									return 0
//line /usr/local/go/src/compress/flate/dict_decoder.go:157
		// _ = "end of CoverTab[26024]"
	} else {
//line /usr/local/go/src/compress/flate/dict_decoder.go:158
		_go_fuzz_dep_.CoverTab[26025]++
//line /usr/local/go/src/compress/flate/dict_decoder.go:158
		// _ = "end of CoverTab[26025]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:158
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:158
	// _ = "end of CoverTab[26020]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:158
	_go_fuzz_dep_.CoverTab[26021]++
								dstBase := dstPos
								srcPos := dstPos - dist

//line /usr/local/go/src/compress/flate/dict_decoder.go:163
	for dstPos < endPos {
//line /usr/local/go/src/compress/flate/dict_decoder.go:163
		_go_fuzz_dep_.CoverTab[26026]++
									dstPos += copy(dd.hist[dstPos:endPos], dd.hist[srcPos:dstPos])
//line /usr/local/go/src/compress/flate/dict_decoder.go:164
		// _ = "end of CoverTab[26026]"
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:165
	// _ = "end of CoverTab[26021]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:165
	_go_fuzz_dep_.CoverTab[26022]++

								dd.wrPos = dstPos
								return dstPos - dstBase
//line /usr/local/go/src/compress/flate/dict_decoder.go:168
	// _ = "end of CoverTab[26022]"
}

// readFlush returns a slice of the historical buffer that is ready to be
//line /usr/local/go/src/compress/flate/dict_decoder.go:171
// emitted to the user. The data returned by readFlush must be fully consumed
//line /usr/local/go/src/compress/flate/dict_decoder.go:171
// before calling any other dictDecoder methods.
//line /usr/local/go/src/compress/flate/dict_decoder.go:174
func (dd *dictDecoder) readFlush() []byte {
//line /usr/local/go/src/compress/flate/dict_decoder.go:174
	_go_fuzz_dep_.CoverTab[26027]++
								toRead := dd.hist[dd.rdPos:dd.wrPos]
								dd.rdPos = dd.wrPos
								if dd.wrPos == len(dd.hist) {
//line /usr/local/go/src/compress/flate/dict_decoder.go:177
		_go_fuzz_dep_.CoverTab[26029]++
									dd.wrPos, dd.rdPos = 0, 0
									dd.full = true
//line /usr/local/go/src/compress/flate/dict_decoder.go:179
		// _ = "end of CoverTab[26029]"
	} else {
//line /usr/local/go/src/compress/flate/dict_decoder.go:180
		_go_fuzz_dep_.CoverTab[26030]++
//line /usr/local/go/src/compress/flate/dict_decoder.go:180
		// _ = "end of CoverTab[26030]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:180
	}
//line /usr/local/go/src/compress/flate/dict_decoder.go:180
	// _ = "end of CoverTab[26027]"
//line /usr/local/go/src/compress/flate/dict_decoder.go:180
	_go_fuzz_dep_.CoverTab[26028]++
								return toRead
//line /usr/local/go/src/compress/flate/dict_decoder.go:181
	// _ = "end of CoverTab[26028]"
}

//line /usr/local/go/src/compress/flate/dict_decoder.go:182
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/dict_decoder.go:182
var _ = _go_fuzz_dep_.CoverTab
