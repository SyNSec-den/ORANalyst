// Copyright 2011 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
package snappy

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:5
)

import (
	"encoding/binary"
	"errors"
	"io"
)

// Encode returns the encoded form of src. The returned slice may be a sub-
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:13
// slice of dst if dst was large enough to hold the entire encoded block.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:13
// Otherwise, a newly allocated slice will be returned.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:13
// The dst and src must not overlap. It is valid to pass a nil dst.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:13
// Encode handles the Snappy block format, not the Snappy stream format.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:20
func Encode(dst, src []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:20
	_go_fuzz_dep_.CoverTab[82147]++
										if n := MaxEncodedLen(len(src)); n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:21
		_go_fuzz_dep_.CoverTab[82150]++
											panic(ErrTooLarge)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:22
		// _ = "end of CoverTab[82150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:23
		_go_fuzz_dep_.CoverTab[82151]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:23
		if len(dst) < n {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:23
			_go_fuzz_dep_.CoverTab[82152]++
												dst = make([]byte, n)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:24
			// _ = "end of CoverTab[82152]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
			_go_fuzz_dep_.CoverTab[82153]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
			// _ = "end of CoverTab[82153]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
		// _ = "end of CoverTab[82151]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
	// _ = "end of CoverTab[82147]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:25
	_go_fuzz_dep_.CoverTab[82148]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:28
	d := binary.PutUvarint(dst, uint64(len(src)))

	for len(src) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:30
		_go_fuzz_dep_.CoverTab[82154]++
											p := src
											src = nil
											if len(p) > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:33
			_go_fuzz_dep_.CoverTab[82156]++
												p, src = p[:maxBlockSize], p[maxBlockSize:]
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:34
			// _ = "end of CoverTab[82156]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:35
			_go_fuzz_dep_.CoverTab[82157]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:35
			// _ = "end of CoverTab[82157]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:35
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:35
		// _ = "end of CoverTab[82154]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:35
		_go_fuzz_dep_.CoverTab[82155]++
											if len(p) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:36
			_go_fuzz_dep_.CoverTab[82158]++
												d += emitLiteral(dst[d:], p)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:37
			// _ = "end of CoverTab[82158]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:38
			_go_fuzz_dep_.CoverTab[82159]++
												d += encodeBlock(dst[d:], p)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:39
			// _ = "end of CoverTab[82159]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:40
		// _ = "end of CoverTab[82155]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:41
	// _ = "end of CoverTab[82148]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:41
	_go_fuzz_dep_.CoverTab[82149]++
										return dst[:d]
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:42
	// _ = "end of CoverTab[82149]"
}

// inputMargin is the minimum number of extra input bytes to keep, inside
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// encodeBlock's inner loop. On some architectures, this margin lets us
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// implement a fast path for emitLiteral, where the copy of short (<= 16 byte)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// literals can be implemented as a single load to and store from a 16-byte
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// register. That literal's actual length can be as short as 1 byte, so this
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// can copy up to 15 bytes too much, but that's OK as subsequent iterations of
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// the encoding loop will fix up the copy overrun, and this inputMargin ensures
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:45
// that we don't overrun the dst and src buffers.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:53
const inputMargin = 16 - 1

// minNonLiteralBlockSize is the minimum size of the input to encodeBlock that
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// could be encoded with a copy tag. This is the minimum with respect to the
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// algorithm used by encodeBlock, not a minimum enforced by the file format.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// The encoded output must start with at least a 1 byte literal, as there are
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// no previous bytes to copy. A minimal (1 byte) copy after that, generated
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// from an emitCopy call in encodeBlock's main loop, would require at least
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// another inputMargin bytes, for the reason above: we want any emitLiteral
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// calls inside encodeBlock's main loop to use the fast path if possible, which
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// requires being able to overrun by inputMargin bytes. Thus,
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// minNonLiteralBlockSize equals 1 + 1 + inputMargin.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// The C++ code doesn't use this exact threshold, but it could, as discussed at
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// https://groups.google.com/d/topic/snappy-compression/oGbhsdIJSJ8/discussion
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// The difference between Go (2+inputMargin) and C++ (inputMargin) is purely an
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// optimization. It should not affect the encoded form. This is tested by
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:55
// TestSameEncodingAsCppShortCopies.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:72
const minNonLiteralBlockSize = 1 + 1 + inputMargin

// MaxEncodedLen returns the maximum length of a snappy block, given its
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:74
// uncompressed length.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:74
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:74
// It will return a negative value if srcLen is too large to encode.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:78
func MaxEncodedLen(srcLen int) int {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:78
	_go_fuzz_dep_.CoverTab[82160]++
										n := uint64(srcLen)
										if n > 0xffffffff {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:80
		_go_fuzz_dep_.CoverTab[82163]++
											return -1
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:81
		// _ = "end of CoverTab[82163]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:82
		_go_fuzz_dep_.CoverTab[82164]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:82
		// _ = "end of CoverTab[82164]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:82
		// _ = "end of CoverTab[82160]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:82
		_go_fuzz_dep_.CoverTab[82161]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:103
	n = 32 + n + n/6
	if n > 0xffffffff {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:104
		_go_fuzz_dep_.CoverTab[82165]++
												return -1
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:105
		// _ = "end of CoverTab[82165]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:106
		_go_fuzz_dep_.CoverTab[82166]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:106
		// _ = "end of CoverTab[82166]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:106
	// _ = "end of CoverTab[82161]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:106
	_go_fuzz_dep_.CoverTab[82162]++
											return int(n)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:107
	// _ = "end of CoverTab[82162]"
}

var errClosed = errors.New("snappy: Writer is closed")

// NewWriter returns a new Writer that compresses to w.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
// The Writer returned does not buffer writes. There is no need to Flush or
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
// Close such a Writer.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
// Deprecated: the Writer returned is not suitable for many small writes, only
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
// for few large writes. Use NewBufferedWriter instead, which is efficient
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
// regardless of the frequency and shape of the writes, and remember to Close
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:112
// that Writer when done.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:121
func NewWriter(w io.Writer) *Writer {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:121
	_go_fuzz_dep_.CoverTab[82167]++
											return &Writer{
		w:	w,
		obuf:	make([]byte, obufLen),
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:125
	// _ = "end of CoverTab[82167]"
}

// NewBufferedWriter returns a new Writer that compresses to w, using the
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:128
// framing format described at
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:128
// https://github.com/google/snappy/blob/master/framing_format.txt
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:128
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:128
// The Writer returned buffers writes. Users must call Close to guarantee all
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:128
// data has been forwarded to the underlying io.Writer. They may also call
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:128
// Flush zero or more times before calling Close.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:135
func NewBufferedWriter(w io.Writer) *Writer {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:135
	_go_fuzz_dep_.CoverTab[82168]++
											return &Writer{
		w:	w,
		ibuf:	make([]byte, 0, maxBlockSize),
		obuf:	make([]byte, obufLen),
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:140
	// _ = "end of CoverTab[82168]"
}

// Writer is an io.Writer that can write Snappy-compressed bytes.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:143
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:143
// Writer handles the Snappy stream format, not the Snappy block format.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:146
type Writer struct {
	w	io.Writer
	err	error

	// ibuf is a buffer for the incoming (uncompressed) bytes.
	//
	// Its use is optional. For backwards compatibility, Writers created by the
	// NewWriter function have ibuf == nil, do not buffer incoming bytes, and
	// therefore do not need to be Flush'ed or Close'd.
	ibuf	[]byte

	// obuf is a buffer for the outgoing (compressed) bytes.
	obuf	[]byte

	// wroteStreamHeader is whether we have written the stream header.
	wroteStreamHeader	bool
}

// Reset discards the writer's state and switches the Snappy writer to write to
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:164
// w. This permits reusing a Writer rather than allocating a new one.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:166
func (w *Writer) Reset(writer io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:166
	_go_fuzz_dep_.CoverTab[82169]++
											w.w = writer
											w.err = nil
											if w.ibuf != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:169
		_go_fuzz_dep_.CoverTab[82171]++
												w.ibuf = w.ibuf[:0]
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:170
		// _ = "end of CoverTab[82171]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:171
		_go_fuzz_dep_.CoverTab[82172]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:171
		// _ = "end of CoverTab[82172]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:171
	// _ = "end of CoverTab[82169]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:171
	_go_fuzz_dep_.CoverTab[82170]++
											w.wroteStreamHeader = false
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:172
	// _ = "end of CoverTab[82170]"
}

// Write satisfies the io.Writer interface.
func (w *Writer) Write(p []byte) (nRet int, errRet error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:176
	_go_fuzz_dep_.CoverTab[82173]++
											if w.ibuf == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:177
		_go_fuzz_dep_.CoverTab[82177]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:182
		return w.write(p)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:182
		// _ = "end of CoverTab[82177]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:183
		_go_fuzz_dep_.CoverTab[82178]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:183
		// _ = "end of CoverTab[82178]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:183
	// _ = "end of CoverTab[82173]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:183
	_go_fuzz_dep_.CoverTab[82174]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:188
	for len(p) > (cap(w.ibuf)-len(w.ibuf)) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:188
		_go_fuzz_dep_.CoverTab[82179]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:188
		return w.err == nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:188
		// _ = "end of CoverTab[82179]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:188
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:188
		_go_fuzz_dep_.CoverTab[82180]++
												var n int
												if len(w.ibuf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:190
			_go_fuzz_dep_.CoverTab[82182]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:193
			n, _ = w.write(p)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:193
			// _ = "end of CoverTab[82182]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:194
			_go_fuzz_dep_.CoverTab[82183]++
													n = copy(w.ibuf[len(w.ibuf):cap(w.ibuf)], p)
													w.ibuf = w.ibuf[:len(w.ibuf)+n]
													w.Flush()
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:197
			// _ = "end of CoverTab[82183]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:198
		// _ = "end of CoverTab[82180]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:198
		_go_fuzz_dep_.CoverTab[82181]++
												nRet += n
												p = p[n:]
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:200
		// _ = "end of CoverTab[82181]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:201
	// _ = "end of CoverTab[82174]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:201
	_go_fuzz_dep_.CoverTab[82175]++
											if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:202
		_go_fuzz_dep_.CoverTab[82184]++
												return nRet, w.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:203
		// _ = "end of CoverTab[82184]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:204
		_go_fuzz_dep_.CoverTab[82185]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:204
		// _ = "end of CoverTab[82185]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:204
	// _ = "end of CoverTab[82175]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:204
	_go_fuzz_dep_.CoverTab[82176]++
											n := copy(w.ibuf[len(w.ibuf):cap(w.ibuf)], p)
											w.ibuf = w.ibuf[:len(w.ibuf)+n]
											nRet += n
											return nRet, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:208
	// _ = "end of CoverTab[82176]"
}

func (w *Writer) write(p []byte) (nRet int, errRet error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:211
	_go_fuzz_dep_.CoverTab[82186]++
											if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:212
		_go_fuzz_dep_.CoverTab[82189]++
												return 0, w.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:213
		// _ = "end of CoverTab[82189]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:214
		_go_fuzz_dep_.CoverTab[82190]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:214
		// _ = "end of CoverTab[82190]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:214
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:214
	// _ = "end of CoverTab[82186]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:214
	_go_fuzz_dep_.CoverTab[82187]++
											for len(p) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:215
		_go_fuzz_dep_.CoverTab[82191]++
												obufStart := len(magicChunk)
												if !w.wroteStreamHeader {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:217
			_go_fuzz_dep_.CoverTab[82197]++
													w.wroteStreamHeader = true
													copy(w.obuf, magicChunk)
													obufStart = 0
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:220
			// _ = "end of CoverTab[82197]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:221
			_go_fuzz_dep_.CoverTab[82198]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:221
			// _ = "end of CoverTab[82198]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:221
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:221
		// _ = "end of CoverTab[82191]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:221
		_go_fuzz_dep_.CoverTab[82192]++

												var uncompressed []byte
												if len(p) > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:224
			_go_fuzz_dep_.CoverTab[82199]++
													uncompressed, p = p[:maxBlockSize], p[maxBlockSize:]
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:225
			// _ = "end of CoverTab[82199]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:226
			_go_fuzz_dep_.CoverTab[82200]++
													uncompressed, p = p, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:227
			// _ = "end of CoverTab[82200]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:228
		// _ = "end of CoverTab[82192]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:228
		_go_fuzz_dep_.CoverTab[82193]++
												checksum := crc(uncompressed)

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:233
		compressed := Encode(w.obuf[obufHeaderLen:], uncompressed)
		chunkType := uint8(chunkTypeCompressedData)
		chunkLen := 4 + len(compressed)
		obufEnd := obufHeaderLen + len(compressed)
		if len(compressed) >= len(uncompressed)-len(uncompressed)/8 {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:237
			_go_fuzz_dep_.CoverTab[82201]++
													chunkType = chunkTypeUncompressedData
													chunkLen = 4 + len(uncompressed)
													obufEnd = obufHeaderLen
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:240
			// _ = "end of CoverTab[82201]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:241
			_go_fuzz_dep_.CoverTab[82202]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:241
			// _ = "end of CoverTab[82202]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:241
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:241
		// _ = "end of CoverTab[82193]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:241
		_go_fuzz_dep_.CoverTab[82194]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:244
		w.obuf[len(magicChunk)+0] = chunkType
		w.obuf[len(magicChunk)+1] = uint8(chunkLen >> 0)
		w.obuf[len(magicChunk)+2] = uint8(chunkLen >> 8)
		w.obuf[len(magicChunk)+3] = uint8(chunkLen >> 16)
		w.obuf[len(magicChunk)+4] = uint8(checksum >> 0)
		w.obuf[len(magicChunk)+5] = uint8(checksum >> 8)
		w.obuf[len(magicChunk)+6] = uint8(checksum >> 16)
		w.obuf[len(magicChunk)+7] = uint8(checksum >> 24)

		if _, err := w.w.Write(w.obuf[obufStart:obufEnd]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:253
			_go_fuzz_dep_.CoverTab[82203]++
													w.err = err
													return nRet, err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:255
			// _ = "end of CoverTab[82203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:256
			_go_fuzz_dep_.CoverTab[82204]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:256
			// _ = "end of CoverTab[82204]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:256
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:256
		// _ = "end of CoverTab[82194]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:256
		_go_fuzz_dep_.CoverTab[82195]++
												if chunkType == chunkTypeUncompressedData {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:257
			_go_fuzz_dep_.CoverTab[82205]++
													if _, err := w.w.Write(uncompressed); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:258
				_go_fuzz_dep_.CoverTab[82206]++
														w.err = err
														return nRet, err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:260
				// _ = "end of CoverTab[82206]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:261
				_go_fuzz_dep_.CoverTab[82207]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:261
				// _ = "end of CoverTab[82207]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:261
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:261
			// _ = "end of CoverTab[82205]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:262
			_go_fuzz_dep_.CoverTab[82208]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:262
			// _ = "end of CoverTab[82208]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:262
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:262
		// _ = "end of CoverTab[82195]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:262
		_go_fuzz_dep_.CoverTab[82196]++
												nRet += len(uncompressed)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:263
		// _ = "end of CoverTab[82196]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:264
	// _ = "end of CoverTab[82187]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:264
	_go_fuzz_dep_.CoverTab[82188]++
											return nRet, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:265
	// _ = "end of CoverTab[82188]"
}

// Flush flushes the Writer to its underlying io.Writer.
func (w *Writer) Flush() error {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:269
	_go_fuzz_dep_.CoverTab[82209]++
											if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:270
		_go_fuzz_dep_.CoverTab[82212]++
												return w.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:271
		// _ = "end of CoverTab[82212]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:272
		_go_fuzz_dep_.CoverTab[82213]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:272
		// _ = "end of CoverTab[82213]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:272
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:272
	// _ = "end of CoverTab[82209]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:272
	_go_fuzz_dep_.CoverTab[82210]++
											if len(w.ibuf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:273
		_go_fuzz_dep_.CoverTab[82214]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:274
		// _ = "end of CoverTab[82214]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:275
		_go_fuzz_dep_.CoverTab[82215]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:275
		// _ = "end of CoverTab[82215]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:275
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:275
	// _ = "end of CoverTab[82210]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:275
	_go_fuzz_dep_.CoverTab[82211]++
											w.write(w.ibuf)
											w.ibuf = w.ibuf[:0]
											return w.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:278
	// _ = "end of CoverTab[82211]"
}

// Close calls Flush and then closes the Writer.
func (w *Writer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:282
	_go_fuzz_dep_.CoverTab[82216]++
											w.Flush()
											ret := w.err
											if w.err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:285
		_go_fuzz_dep_.CoverTab[82218]++
												w.err = errClosed
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:286
		// _ = "end of CoverTab[82218]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:287
		_go_fuzz_dep_.CoverTab[82219]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:287
		// _ = "end of CoverTab[82219]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:287
	// _ = "end of CoverTab[82216]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:287
	_go_fuzz_dep_.CoverTab[82217]++
											return ret
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:288
	// _ = "end of CoverTab[82217]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:289
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/encode.go:289
var _ = _go_fuzz_dep_.CoverTab
