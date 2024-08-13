// Copyright 2011 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
package snapref

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:5
)

import (
	"encoding/binary"
	"errors"
	"io"
)

// Encode returns the encoded form of src. The returned slice may be a sub-
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:13
// slice of dst if dst was large enough to hold the entire encoded block.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:13
// Otherwise, a newly allocated slice will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:13
// The dst and src must not overlap. It is valid to pass a nil dst.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:13
// Encode handles the Snappy block format, not the Snappy stream format.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:20
func Encode(dst, src []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:20
	_go_fuzz_dep_.CoverTab[90538]++
													if n := MaxEncodedLen(len(src)); n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:21
		_go_fuzz_dep_.CoverTab[90541]++
														panic(ErrTooLarge)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:22
		// _ = "end of CoverTab[90541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:23
		_go_fuzz_dep_.CoverTab[90542]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:23
		if len(dst) < n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:23
			_go_fuzz_dep_.CoverTab[90543]++
															dst = make([]byte, n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:24
			// _ = "end of CoverTab[90543]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
			_go_fuzz_dep_.CoverTab[90544]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
			// _ = "end of CoverTab[90544]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
		// _ = "end of CoverTab[90542]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
	// _ = "end of CoverTab[90538]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:25
	_go_fuzz_dep_.CoverTab[90539]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:28
	d := binary.PutUvarint(dst, uint64(len(src)))

	for len(src) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:30
		_go_fuzz_dep_.CoverTab[90545]++
														p := src
														src = nil
														if len(p) > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:33
			_go_fuzz_dep_.CoverTab[90547]++
															p, src = p[:maxBlockSize], p[maxBlockSize:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:34
			// _ = "end of CoverTab[90547]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:35
			_go_fuzz_dep_.CoverTab[90548]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:35
			// _ = "end of CoverTab[90548]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:35
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:35
		// _ = "end of CoverTab[90545]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:35
		_go_fuzz_dep_.CoverTab[90546]++
														if len(p) < minNonLiteralBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:36
			_go_fuzz_dep_.CoverTab[90549]++
															d += emitLiteral(dst[d:], p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:37
			// _ = "end of CoverTab[90549]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:38
			_go_fuzz_dep_.CoverTab[90550]++
															d += encodeBlock(dst[d:], p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:39
			// _ = "end of CoverTab[90550]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:40
		// _ = "end of CoverTab[90546]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:41
	// _ = "end of CoverTab[90539]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:41
	_go_fuzz_dep_.CoverTab[90540]++
													return dst[:d]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:42
	// _ = "end of CoverTab[90540]"
}

// inputMargin is the minimum number of extra input bytes to keep, inside
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// encodeBlock's inner loop. On some architectures, this margin lets us
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// implement a fast path for emitLiteral, where the copy of short (<= 16 byte)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// literals can be implemented as a single load to and store from a 16-byte
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// register. That literal's actual length can be as short as 1 byte, so this
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// can copy up to 15 bytes too much, but that's OK as subsequent iterations of
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// the encoding loop will fix up the copy overrun, and this inputMargin ensures
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:45
// that we don't overrun the dst and src buffers.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:53
const inputMargin = 16 - 1

// minNonLiteralBlockSize is the minimum size of the input to encodeBlock that
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// could be encoded with a copy tag. This is the minimum with respect to the
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// algorithm used by encodeBlock, not a minimum enforced by the file format.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// The encoded output must start with at least a 1 byte literal, as there are
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// no previous bytes to copy. A minimal (1 byte) copy after that, generated
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// from an emitCopy call in encodeBlock's main loop, would require at least
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// another inputMargin bytes, for the reason above: we want any emitLiteral
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// calls inside encodeBlock's main loop to use the fast path if possible, which
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// requires being able to overrun by inputMargin bytes. Thus,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// minNonLiteralBlockSize equals 1 + 1 + inputMargin.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// The C++ code doesn't use this exact threshold, but it could, as discussed at
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// https://groups.google.com/d/topic/snappy-compression/oGbhsdIJSJ8/discussion
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// The difference between Go (2+inputMargin) and C++ (inputMargin) is purely an
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// optimization. It should not affect the encoded form. This is tested by
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:55
// TestSameEncodingAsCppShortCopies.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:72
const minNonLiteralBlockSize = 1 + 1 + inputMargin

// MaxEncodedLen returns the maximum length of a snappy block, given its
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:74
// uncompressed length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:74
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:74
// It will return a negative value if srcLen is too large to encode.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:78
func MaxEncodedLen(srcLen int) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:78
	_go_fuzz_dep_.CoverTab[90551]++
													n := uint64(srcLen)
													if n > 0xffffffff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:80
		_go_fuzz_dep_.CoverTab[90554]++
														return -1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:81
		// _ = "end of CoverTab[90554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:82
		_go_fuzz_dep_.CoverTab[90555]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:82
		// _ = "end of CoverTab[90555]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:82
	// _ = "end of CoverTab[90551]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:82
	_go_fuzz_dep_.CoverTab[90552]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:103
	n = 32 + n + n/6
	if n > 0xffffffff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:104
		_go_fuzz_dep_.CoverTab[90556]++
														return -1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:105
		// _ = "end of CoverTab[90556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:106
		_go_fuzz_dep_.CoverTab[90557]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:106
		// _ = "end of CoverTab[90557]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:106
	// _ = "end of CoverTab[90552]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:106
	_go_fuzz_dep_.CoverTab[90553]++
													return int(n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:107
	// _ = "end of CoverTab[90553]"
}

var errClosed = errors.New("snappy: Writer is closed")

// NewWriter returns a new Writer that compresses to w.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
// The Writer returned does not buffer writes. There is no need to Flush or
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
// Close such a Writer.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
// Deprecated: the Writer returned is not suitable for many small writes, only
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
// for few large writes. Use NewBufferedWriter instead, which is efficient
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
// regardless of the frequency and shape of the writes, and remember to Close
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:112
// that Writer when done.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:121
func NewWriter(w io.Writer) *Writer {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:121
	_go_fuzz_dep_.CoverTab[90558]++
													return &Writer{
		w:	w,
		obuf:	make([]byte, obufLen),
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:125
	// _ = "end of CoverTab[90558]"
}

// NewBufferedWriter returns a new Writer that compresses to w, using the
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:128
// framing format described at
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:128
// https://github.com/google/snappy/blob/master/framing_format.txt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:128
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:128
// The Writer returned buffers writes. Users must call Close to guarantee all
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:128
// data has been forwarded to the underlying io.Writer. They may also call
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:128
// Flush zero or more times before calling Close.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:135
func NewBufferedWriter(w io.Writer) *Writer {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:135
	_go_fuzz_dep_.CoverTab[90559]++
													return &Writer{
		w:	w,
		ibuf:	make([]byte, 0, maxBlockSize),
		obuf:	make([]byte, obufLen),
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:140
	// _ = "end of CoverTab[90559]"
}

// Writer is an io.Writer that can write Snappy-compressed bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:143
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:143
// Writer handles the Snappy stream format, not the Snappy block format.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:146
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:164
// w. This permits reusing a Writer rather than allocating a new one.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:166
func (w *Writer) Reset(writer io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:166
	_go_fuzz_dep_.CoverTab[90560]++
													w.w = writer
													w.err = nil
													if w.ibuf != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:169
		_go_fuzz_dep_.CoverTab[90562]++
														w.ibuf = w.ibuf[:0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:170
		// _ = "end of CoverTab[90562]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:171
		_go_fuzz_dep_.CoverTab[90563]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:171
		// _ = "end of CoverTab[90563]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:171
	// _ = "end of CoverTab[90560]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:171
	_go_fuzz_dep_.CoverTab[90561]++
													w.wroteStreamHeader = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:172
	// _ = "end of CoverTab[90561]"
}

// Write satisfies the io.Writer interface.
func (w *Writer) Write(p []byte) (nRet int, errRet error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:176
	_go_fuzz_dep_.CoverTab[90564]++
													if w.ibuf == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:177
		_go_fuzz_dep_.CoverTab[90568]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:182
		return w.write(p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:182
		// _ = "end of CoverTab[90568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:183
		_go_fuzz_dep_.CoverTab[90569]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:183
		// _ = "end of CoverTab[90569]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:183
	// _ = "end of CoverTab[90564]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:183
	_go_fuzz_dep_.CoverTab[90565]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:188
	for len(p) > (cap(w.ibuf)-len(w.ibuf)) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:188
		_go_fuzz_dep_.CoverTab[90570]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:188
		return w.err == nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:188
		// _ = "end of CoverTab[90570]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:188
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:188
		_go_fuzz_dep_.CoverTab[90571]++
														var n int
														if len(w.ibuf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:190
			_go_fuzz_dep_.CoverTab[90573]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:193
			n, _ = w.write(p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:193
			// _ = "end of CoverTab[90573]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:194
			_go_fuzz_dep_.CoverTab[90574]++
															n = copy(w.ibuf[len(w.ibuf):cap(w.ibuf)], p)
															w.ibuf = w.ibuf[:len(w.ibuf)+n]
															w.Flush()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:197
			// _ = "end of CoverTab[90574]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:198
		// _ = "end of CoverTab[90571]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:198
		_go_fuzz_dep_.CoverTab[90572]++
														nRet += n
														p = p[n:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:200
		// _ = "end of CoverTab[90572]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:201
	// _ = "end of CoverTab[90565]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:201
	_go_fuzz_dep_.CoverTab[90566]++
													if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:202
		_go_fuzz_dep_.CoverTab[90575]++
														return nRet, w.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:203
		// _ = "end of CoverTab[90575]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:204
		_go_fuzz_dep_.CoverTab[90576]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:204
		// _ = "end of CoverTab[90576]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:204
	// _ = "end of CoverTab[90566]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:204
	_go_fuzz_dep_.CoverTab[90567]++
													n := copy(w.ibuf[len(w.ibuf):cap(w.ibuf)], p)
													w.ibuf = w.ibuf[:len(w.ibuf)+n]
													nRet += n
													return nRet, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:208
	// _ = "end of CoverTab[90567]"
}

func (w *Writer) write(p []byte) (nRet int, errRet error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:211
	_go_fuzz_dep_.CoverTab[90577]++
													if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:212
		_go_fuzz_dep_.CoverTab[90580]++
														return 0, w.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:213
		// _ = "end of CoverTab[90580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:214
		_go_fuzz_dep_.CoverTab[90581]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:214
		// _ = "end of CoverTab[90581]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:214
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:214
	// _ = "end of CoverTab[90577]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:214
	_go_fuzz_dep_.CoverTab[90578]++
													for len(p) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:215
		_go_fuzz_dep_.CoverTab[90582]++
														obufStart := len(magicChunk)
														if !w.wroteStreamHeader {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:217
			_go_fuzz_dep_.CoverTab[90588]++
															w.wroteStreamHeader = true
															copy(w.obuf, magicChunk)
															obufStart = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:220
			// _ = "end of CoverTab[90588]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:221
			_go_fuzz_dep_.CoverTab[90589]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:221
			// _ = "end of CoverTab[90589]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:221
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:221
		// _ = "end of CoverTab[90582]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:221
		_go_fuzz_dep_.CoverTab[90583]++

														var uncompressed []byte
														if len(p) > maxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:224
			_go_fuzz_dep_.CoverTab[90590]++
															uncompressed, p = p[:maxBlockSize], p[maxBlockSize:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:225
			// _ = "end of CoverTab[90590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:226
			_go_fuzz_dep_.CoverTab[90591]++
															uncompressed, p = p, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:227
			// _ = "end of CoverTab[90591]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:228
		// _ = "end of CoverTab[90583]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:228
		_go_fuzz_dep_.CoverTab[90584]++
														checksum := crc(uncompressed)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:233
		compressed := Encode(w.obuf[obufHeaderLen:], uncompressed)
		chunkType := uint8(chunkTypeCompressedData)
		chunkLen := 4 + len(compressed)
		obufEnd := obufHeaderLen + len(compressed)
		if len(compressed) >= len(uncompressed)-len(uncompressed)/8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:237
			_go_fuzz_dep_.CoverTab[90592]++
															chunkType = chunkTypeUncompressedData
															chunkLen = 4 + len(uncompressed)
															obufEnd = obufHeaderLen
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:240
			// _ = "end of CoverTab[90592]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:241
			_go_fuzz_dep_.CoverTab[90593]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:241
			// _ = "end of CoverTab[90593]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:241
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:241
		// _ = "end of CoverTab[90584]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:241
		_go_fuzz_dep_.CoverTab[90585]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:244
		w.obuf[len(magicChunk)+0] = chunkType
		w.obuf[len(magicChunk)+1] = uint8(chunkLen >> 0)
		w.obuf[len(magicChunk)+2] = uint8(chunkLen >> 8)
		w.obuf[len(magicChunk)+3] = uint8(chunkLen >> 16)
		w.obuf[len(magicChunk)+4] = uint8(checksum >> 0)
		w.obuf[len(magicChunk)+5] = uint8(checksum >> 8)
		w.obuf[len(magicChunk)+6] = uint8(checksum >> 16)
		w.obuf[len(magicChunk)+7] = uint8(checksum >> 24)

		if _, err := w.w.Write(w.obuf[obufStart:obufEnd]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:253
			_go_fuzz_dep_.CoverTab[90594]++
															w.err = err
															return nRet, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:255
			// _ = "end of CoverTab[90594]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:256
			_go_fuzz_dep_.CoverTab[90595]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:256
			// _ = "end of CoverTab[90595]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:256
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:256
		// _ = "end of CoverTab[90585]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:256
		_go_fuzz_dep_.CoverTab[90586]++
														if chunkType == chunkTypeUncompressedData {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:257
			_go_fuzz_dep_.CoverTab[90596]++
															if _, err := w.w.Write(uncompressed); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:258
				_go_fuzz_dep_.CoverTab[90597]++
																w.err = err
																return nRet, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:260
				// _ = "end of CoverTab[90597]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:261
				_go_fuzz_dep_.CoverTab[90598]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:261
				// _ = "end of CoverTab[90598]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:261
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:261
			// _ = "end of CoverTab[90596]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:262
			_go_fuzz_dep_.CoverTab[90599]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:262
			// _ = "end of CoverTab[90599]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:262
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:262
		// _ = "end of CoverTab[90586]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:262
		_go_fuzz_dep_.CoverTab[90587]++
														nRet += len(uncompressed)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:263
		// _ = "end of CoverTab[90587]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:264
	// _ = "end of CoverTab[90578]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:264
	_go_fuzz_dep_.CoverTab[90579]++
													return nRet, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:265
	// _ = "end of CoverTab[90579]"
}

// Flush flushes the Writer to its underlying io.Writer.
func (w *Writer) Flush() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:269
	_go_fuzz_dep_.CoverTab[90600]++
													if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:270
		_go_fuzz_dep_.CoverTab[90603]++
														return w.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:271
		// _ = "end of CoverTab[90603]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:272
		_go_fuzz_dep_.CoverTab[90604]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:272
		// _ = "end of CoverTab[90604]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:272
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:272
	// _ = "end of CoverTab[90600]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:272
	_go_fuzz_dep_.CoverTab[90601]++
													if len(w.ibuf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:273
		_go_fuzz_dep_.CoverTab[90605]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:274
		// _ = "end of CoverTab[90605]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:275
		_go_fuzz_dep_.CoverTab[90606]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:275
		// _ = "end of CoverTab[90606]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:275
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:275
	// _ = "end of CoverTab[90601]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:275
	_go_fuzz_dep_.CoverTab[90602]++
													w.write(w.ibuf)
													w.ibuf = w.ibuf[:0]
													return w.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:278
	// _ = "end of CoverTab[90602]"
}

// Close calls Flush and then closes the Writer.
func (w *Writer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:282
	_go_fuzz_dep_.CoverTab[90607]++
													w.Flush()
													ret := w.err
													if w.err == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:285
		_go_fuzz_dep_.CoverTab[90609]++
														w.err = errClosed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:286
		// _ = "end of CoverTab[90609]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:287
		_go_fuzz_dep_.CoverTab[90610]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:287
		// _ = "end of CoverTab[90610]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:287
	// _ = "end of CoverTab[90607]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:287
	_go_fuzz_dep_.CoverTab[90608]++
													return ret
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:288
	// _ = "end of CoverTab[90608]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:289
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/encode.go:289
var _ = _go_fuzz_dep_.CoverTab
