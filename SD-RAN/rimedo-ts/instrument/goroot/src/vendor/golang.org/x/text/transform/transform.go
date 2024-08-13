// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:5
// Package transform provides reader and writer wrappers that transform the
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:5
// bytes passing through as well as various transformations. Example
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:5
// transformations provided by other packages include normalization and
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:5
// conversion between character sets.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
package transform

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
import (
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:9
)

import (
	"bytes"
	"errors"
	"io"
	"unicode/utf8"
)

var (
	// ErrShortDst means that the destination buffer was too short to
	// receive all of the transformed bytes.
	ErrShortDst	= errors.New("transform: short destination buffer")

	// ErrShortSrc means that the source buffer has insufficient data to
	// complete the transformation.
	ErrShortSrc	= errors.New("transform: short source buffer")

	// ErrEndOfSpan means that the input and output (the transformed input)
	// are not identical.
	ErrEndOfSpan	= errors.New("transform: input and output are not identical")

	// errInconsistentByteCount means that Transform returned success (nil
	// error) but also returned nSrc inconsistent with the src argument.
	errInconsistentByteCount	= errors.New("transform: inconsistent byte count returned")

	// errShortInternal means that an internal buffer is not large enough
	// to make progress and the Transform operation must be aborted.
	errShortInternal	= errors.New("transform: short internal buffer")
)

// Transformer transforms bytes.
type Transformer interface {
	// Transform writes to dst the transformed bytes read from src, and
	// returns the number of dst bytes written and src bytes read. The
	// atEOF argument tells whether src represents the last bytes of the
	// input.
	//
	// Callers should always process the nDst bytes produced and account
	// for the nSrc bytes consumed before considering the error err.
	//
	// A nil error means that all of the transformed bytes (whether freshly
	// transformed from src or left over from previous Transform calls)
	// were written to dst. A nil error can be returned regardless of
	// whether atEOF is true. If err is nil then nSrc must equal len(src);
	// the converse is not necessarily true.
	//
	// ErrShortDst means that dst was too short to receive all of the
	// transformed bytes. ErrShortSrc means that src had insufficient data
	// to complete the transformation. If both conditions apply, then
	// either error may be returned. Other than the error conditions listed
	// here, implementations are free to report other errors that arise.
	Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)

	// Reset resets the state and allows a Transformer to be reused.
	Reset()
}

// SpanningTransformer extends the Transformer interface with a Span method
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:67
// that determines how much of the input already conforms to the Transformer.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:69
type SpanningTransformer interface {
	Transformer

	// Span returns a position in src such that transforming src[:n] results in
	// identical output src[:n] for these bytes. It does not necessarily return
	// the largest such n. The atEOF argument tells whether src represents the
	// last bytes of the input.
	//
	// Callers should always account for the n bytes consumed before
	// considering the error err.
	//
	// A nil error means that all input bytes are known to be identical to the
	// output produced by the Transformer. A nil error can be returned
	// regardless of whether atEOF is true. If err is nil, then n must
	// equal len(src); the converse is not necessarily true.
	//
	// ErrEndOfSpan means that the Transformer output may differ from the
	// input after n bytes. Note that n may be len(src), meaning that the output
	// would contain additional bytes after otherwise identical output.
	// ErrShortSrc means that src had insufficient data to determine whether the
	// remaining bytes would change. Other than the error conditions listed
	// here, implementations are free to report other errors that arise.
	//
	// Calling Span can modify the Transformer state as a side effect. In
	// effect, it does the transformation just as calling Transform would, only
	// without copying to a destination buffer and only up to a point it can
	// determine the input and output bytes are the same. This is obviously more
	// limited than calling Transform, but can be more efficient in terms of
	// copying and allocating buffers. Calls to Span and Transform may be
	// interleaved.
	Span(src []byte, atEOF bool) (n int, err error)
}

// NopResetter can be embedded by implementations of Transformer to add a nop
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:102
// Reset method.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:104
type NopResetter struct{}

// Reset implements the Reset method of the Transformer interface.
func (NopResetter) Reset()	{ _go_fuzz_dep_.CoverTab[31897]++; // _ = "end of CoverTab[31897]" }

// Reader wraps another io.Reader by transforming the bytes read.
type Reader struct {
	r	io.Reader
	t	Transformer
	err	error

	// dst[dst0:dst1] contains bytes that have been transformed by t but
	// not yet copied out via Read.
	dst		[]byte
	dst0, dst1	int

	// src[src0:src1] contains bytes that have been read from r but not
	// yet transformed through t.
	src		[]byte
	src0, src1	int

	// transformComplete is whether the transformation is complete,
	// regardless of whether or not it was successful.
	transformComplete	bool
}

const defaultBufSize = 4096

// NewReader returns a new Reader that wraps r by transforming the bytes read
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:132
// via t. It calls Reset on t.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:134
func NewReader(r io.Reader, t Transformer) *Reader {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:134
	_go_fuzz_dep_.CoverTab[31898]++
										t.Reset()
										return &Reader{
		r:	r,
		t:	t,
		dst:	make([]byte, defaultBufSize),
		src:	make([]byte, defaultBufSize),
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:141
	// _ = "end of CoverTab[31898]"
}

// Read implements the io.Reader interface.
func (r *Reader) Read(p []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:145
	_go_fuzz_dep_.CoverTab[31899]++
										n, err := 0, error(nil)
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:147
		_go_fuzz_dep_.CoverTab[31900]++

											if r.dst0 != r.dst1 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:149
			_go_fuzz_dep_.CoverTab[31904]++
												n = copy(p, r.dst[r.dst0:r.dst1])
												r.dst0 += n
												if r.dst0 == r.dst1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:152
				_go_fuzz_dep_.CoverTab[31906]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:152
				return r.transformComplete
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:152
				// _ = "end of CoverTab[31906]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:152
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:152
				_go_fuzz_dep_.CoverTab[31907]++
													return n, r.err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:153
				// _ = "end of CoverTab[31907]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:154
				_go_fuzz_dep_.CoverTab[31908]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:154
				// _ = "end of CoverTab[31908]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:154
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:154
			// _ = "end of CoverTab[31904]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:154
			_go_fuzz_dep_.CoverTab[31905]++
												return n, nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:155
			// _ = "end of CoverTab[31905]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:156
			_go_fuzz_dep_.CoverTab[31909]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:156
			if r.transformComplete {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:156
				_go_fuzz_dep_.CoverTab[31910]++
													return 0, r.err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:157
				// _ = "end of CoverTab[31910]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
				_go_fuzz_dep_.CoverTab[31911]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
				// _ = "end of CoverTab[31911]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
			// _ = "end of CoverTab[31909]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
		// _ = "end of CoverTab[31900]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:158
		_go_fuzz_dep_.CoverTab[31901]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:164
		if r.src0 != r.src1 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:164
			_go_fuzz_dep_.CoverTab[31912]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:164
			return r.err != nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:164
			// _ = "end of CoverTab[31912]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:164
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:164
			_go_fuzz_dep_.CoverTab[31913]++
												r.dst0 = 0
												r.dst1, n, err = r.t.Transform(r.dst, r.src[r.src0:r.src1], r.err == io.EOF)
												r.src0 += n

												switch {
			case err == nil:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:170
				_go_fuzz_dep_.CoverTab[31914]++
													if r.src0 != r.src1 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:171
					_go_fuzz_dep_.CoverTab[31920]++
														r.err = errInconsistentByteCount
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:172
					// _ = "end of CoverTab[31920]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:173
					_go_fuzz_dep_.CoverTab[31921]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:173
					// _ = "end of CoverTab[31921]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:173
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:173
				// _ = "end of CoverTab[31914]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:173
				_go_fuzz_dep_.CoverTab[31915]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:176
				r.transformComplete = r.err != nil
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:177
				// _ = "end of CoverTab[31915]"
			case err == ErrShortDst && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
				_go_fuzz_dep_.CoverTab[31922]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
				return (r.dst1 != 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
					_go_fuzz_dep_.CoverTab[31923]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
					return n != 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
					// _ = "end of CoverTab[31923]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
				}())
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
				// _ = "end of CoverTab[31922]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
			}():
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:178
				_go_fuzz_dep_.CoverTab[31916]++

													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:180
				// _ = "end of CoverTab[31916]"
			case err == ErrShortSrc && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				_go_fuzz_dep_.CoverTab[31924]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				return r.src1-r.src0 != len(r.src)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				// _ = "end of CoverTab[31924]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				_go_fuzz_dep_.CoverTab[31925]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				return r.err == nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				// _ = "end of CoverTab[31925]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
			}():
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				_go_fuzz_dep_.CoverTab[31917]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:181
				// _ = "end of CoverTab[31917]"

			default:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:183
				_go_fuzz_dep_.CoverTab[31918]++
													r.transformComplete = true

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:187
				if r.err == nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:187
					_go_fuzz_dep_.CoverTab[31926]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:187
					return r.err == io.EOF
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:187
					// _ = "end of CoverTab[31926]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:187
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:187
					_go_fuzz_dep_.CoverTab[31927]++
														r.err = err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:188
					// _ = "end of CoverTab[31927]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:189
					_go_fuzz_dep_.CoverTab[31928]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:189
					// _ = "end of CoverTab[31928]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:189
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:189
				// _ = "end of CoverTab[31918]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:189
				_go_fuzz_dep_.CoverTab[31919]++
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:190
				// _ = "end of CoverTab[31919]"
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:191
			// _ = "end of CoverTab[31913]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:192
			_go_fuzz_dep_.CoverTab[31929]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:192
			// _ = "end of CoverTab[31929]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:192
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:192
		// _ = "end of CoverTab[31901]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:192
		_go_fuzz_dep_.CoverTab[31902]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:196
		if r.src0 != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:196
			_go_fuzz_dep_.CoverTab[31930]++
												r.src0, r.src1 = 0, copy(r.src, r.src[r.src0:r.src1])
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:197
			// _ = "end of CoverTab[31930]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:198
			_go_fuzz_dep_.CoverTab[31931]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:198
			// _ = "end of CoverTab[31931]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:198
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:198
		// _ = "end of CoverTab[31902]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:198
		_go_fuzz_dep_.CoverTab[31903]++
											n, r.err = r.r.Read(r.src[r.src1:])
											r.src1 += n
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:200
		// _ = "end of CoverTab[31903]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:201
	// _ = "end of CoverTab[31899]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:206
// Writer wraps another io.Writer by transforming the bytes read.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:206
// The user needs to call Close to flush unwritten bytes that may
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:206
// be buffered.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:209
type Writer struct {
	w	io.Writer
	t	Transformer
	dst	[]byte

	// src[:n] contains bytes that have not yet passed through t.
	src	[]byte
	n	int
}

// NewWriter returns a new Writer that wraps w by transforming the bytes written
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:219
// via t. It calls Reset on t.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:221
func NewWriter(w io.Writer, t Transformer) *Writer {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:221
	_go_fuzz_dep_.CoverTab[31932]++
										t.Reset()
										return &Writer{
		w:	w,
		t:	t,
		dst:	make([]byte, defaultBufSize),
		src:	make([]byte, defaultBufSize),
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:228
	// _ = "end of CoverTab[31932]"
}

// Write implements the io.Writer interface. If there are not enough
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:231
// bytes available to complete a Transform, the bytes will be buffered
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:231
// for the next write. Call Close to convert the remaining bytes.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:234
func (w *Writer) Write(data []byte) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:234
	_go_fuzz_dep_.CoverTab[31933]++
										src := data
										if w.n > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:236
		_go_fuzz_dep_.CoverTab[31935]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:239
		n = copy(w.src[w.n:], data)
											w.n += n
											src = w.src[:w.n]
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:241
		// _ = "end of CoverTab[31935]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:242
		_go_fuzz_dep_.CoverTab[31936]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:242
		// _ = "end of CoverTab[31936]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:242
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:242
	// _ = "end of CoverTab[31933]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:242
	_go_fuzz_dep_.CoverTab[31934]++
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:243
		_go_fuzz_dep_.CoverTab[31937]++
											nDst, nSrc, err := w.t.Transform(w.dst, src, false)
											if _, werr := w.w.Write(w.dst[:nDst]); werr != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:245
			_go_fuzz_dep_.CoverTab[31941]++
												return n, werr
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:246
			// _ = "end of CoverTab[31941]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:247
			_go_fuzz_dep_.CoverTab[31942]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:247
			// _ = "end of CoverTab[31942]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:247
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:247
		// _ = "end of CoverTab[31937]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:247
		_go_fuzz_dep_.CoverTab[31938]++
											src = src[nSrc:]
											if w.n == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:249
			_go_fuzz_dep_.CoverTab[31943]++
												n += nSrc
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:250
			// _ = "end of CoverTab[31943]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:251
			_go_fuzz_dep_.CoverTab[31944]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:251
			if len(src) <= n {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:251
				_go_fuzz_dep_.CoverTab[31945]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:254
				w.n = 0
				n -= len(src)
				src = data[n:]
				if n < len(data) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
					_go_fuzz_dep_.CoverTab[31946]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
					return (err == nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
						_go_fuzz_dep_.CoverTab[31947]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
						return err == ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
						// _ = "end of CoverTab[31947]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
					}())
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
					// _ = "end of CoverTab[31946]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:257
					_go_fuzz_dep_.CoverTab[31948]++
														continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:258
					// _ = "end of CoverTab[31948]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:259
					_go_fuzz_dep_.CoverTab[31949]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:259
					// _ = "end of CoverTab[31949]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:259
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:259
				// _ = "end of CoverTab[31945]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
				_go_fuzz_dep_.CoverTab[31950]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
				// _ = "end of CoverTab[31950]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
			// _ = "end of CoverTab[31944]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
		// _ = "end of CoverTab[31938]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:260
		_go_fuzz_dep_.CoverTab[31939]++
											switch err {
		case ErrShortDst:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:262
			_go_fuzz_dep_.CoverTab[31951]++

												if nDst > 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:264
				_go_fuzz_dep_.CoverTab[31955]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:264
				return nSrc > 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:264
				// _ = "end of CoverTab[31955]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:264
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:264
				_go_fuzz_dep_.CoverTab[31956]++
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:265
				// _ = "end of CoverTab[31956]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:266
				_go_fuzz_dep_.CoverTab[31957]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:266
				// _ = "end of CoverTab[31957]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:266
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:266
			// _ = "end of CoverTab[31951]"
		case ErrShortSrc:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:267
			_go_fuzz_dep_.CoverTab[31952]++
												if len(src) < len(w.src) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:268
				_go_fuzz_dep_.CoverTab[31958]++
													m := copy(w.src, src)

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:272
				if w.n == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:272
					_go_fuzz_dep_.CoverTab[31960]++
														n += m
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:273
					// _ = "end of CoverTab[31960]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:274
					_go_fuzz_dep_.CoverTab[31961]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:274
					// _ = "end of CoverTab[31961]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:274
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:274
				// _ = "end of CoverTab[31958]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:274
				_go_fuzz_dep_.CoverTab[31959]++
													w.n = m
													err = nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:276
				// _ = "end of CoverTab[31959]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
				_go_fuzz_dep_.CoverTab[31962]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
				if nDst > 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
					_go_fuzz_dep_.CoverTab[31963]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
					return nSrc > 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
					// _ = "end of CoverTab[31963]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:277
					_go_fuzz_dep_.CoverTab[31964]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:284
					continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:284
					// _ = "end of CoverTab[31964]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:285
					_go_fuzz_dep_.CoverTab[31965]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:285
					// _ = "end of CoverTab[31965]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:285
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:285
				// _ = "end of CoverTab[31962]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:285
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:285
			// _ = "end of CoverTab[31952]"
		case nil:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:286
			_go_fuzz_dep_.CoverTab[31953]++
												if w.n > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:287
				_go_fuzz_dep_.CoverTab[31966]++
													err = errInconsistentByteCount
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:288
				// _ = "end of CoverTab[31966]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
				_go_fuzz_dep_.CoverTab[31967]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
				// _ = "end of CoverTab[31967]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
			// _ = "end of CoverTab[31953]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
			_go_fuzz_dep_.CoverTab[31954]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:289
			// _ = "end of CoverTab[31954]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:290
		// _ = "end of CoverTab[31939]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:290
		_go_fuzz_dep_.CoverTab[31940]++
											return n, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:291
		// _ = "end of CoverTab[31940]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:292
	// _ = "end of CoverTab[31934]"
}

// Close implements the io.Closer interface.
func (w *Writer) Close() error {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:296
	_go_fuzz_dep_.CoverTab[31968]++
										src := w.src[:w.n]
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:298
		_go_fuzz_dep_.CoverTab[31969]++
											nDst, nSrc, err := w.t.Transform(w.dst, src, true)
											if _, werr := w.w.Write(w.dst[:nDst]); werr != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:300
			_go_fuzz_dep_.CoverTab[31972]++
												return werr
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:301
			// _ = "end of CoverTab[31972]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:302
			_go_fuzz_dep_.CoverTab[31973]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:302
			// _ = "end of CoverTab[31973]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:302
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:302
		// _ = "end of CoverTab[31969]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:302
		_go_fuzz_dep_.CoverTab[31970]++
											if err != ErrShortDst {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:303
			_go_fuzz_dep_.CoverTab[31974]++
												return err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:304
			// _ = "end of CoverTab[31974]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:305
			_go_fuzz_dep_.CoverTab[31975]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:305
			// _ = "end of CoverTab[31975]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:305
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:305
		// _ = "end of CoverTab[31970]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:305
		_go_fuzz_dep_.CoverTab[31971]++
											src = src[nSrc:]
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:306
		// _ = "end of CoverTab[31971]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:307
	// _ = "end of CoverTab[31968]"
}

type nop struct{ NopResetter }

func (nop) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:312
	_go_fuzz_dep_.CoverTab[31976]++
										n := copy(dst, src)
										if n < len(src) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:314
		_go_fuzz_dep_.CoverTab[31978]++
											err = ErrShortDst
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:315
		// _ = "end of CoverTab[31978]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:316
		_go_fuzz_dep_.CoverTab[31979]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:316
		// _ = "end of CoverTab[31979]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:316
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:316
	// _ = "end of CoverTab[31976]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:316
	_go_fuzz_dep_.CoverTab[31977]++
										return n, n, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:317
	// _ = "end of CoverTab[31977]"
}

func (nop) Span(src []byte, atEOF bool) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:320
	_go_fuzz_dep_.CoverTab[31980]++
										return len(src), nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:321
	// _ = "end of CoverTab[31980]"
}

type discard struct{ NopResetter }

func (discard) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:326
	_go_fuzz_dep_.CoverTab[31981]++
										return 0, len(src), nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:327
	// _ = "end of CoverTab[31981]"
}

var (
	// Discard is a Transformer for which all Transform calls succeed
	// by consuming all bytes and writing nothing.
	Discard	Transformer	= discard{}

	// Nop is a SpanningTransformer that copies src to dst.
	Nop	SpanningTransformer	= nop{}
)

// chain is a sequence of links. A chain with N Transformers has N+1 links and
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:339
// N+1 buffers. Of those N+1 buffers, the first and last are the src and dst
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:339
// buffers given to chain.Transform and the middle N-1 buffers are intermediate
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:339
// buffers owned by the chain. The i'th link transforms bytes from the i'th
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:339
// buffer chain.link[i].b at read offset chain.link[i].p to the i+1'th buffer
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:339
// chain.link[i+1].b at write offset chain.link[i+1].n, for i in [0, N).
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:345
type chain struct {
	link	[]link
	err	error
	// errStart is the index at which the error occurred plus 1. Processing
	// errStart at this level at the next call to Transform. As long as
	// errStart > 0, chain will not consume any more source bytes.
	errStart	int
}

func (c *chain) fatalError(errIndex int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:354
	_go_fuzz_dep_.CoverTab[31982]++
										if i := errIndex + 1; i > c.errStart {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:355
		_go_fuzz_dep_.CoverTab[31983]++
											c.errStart = i
											c.err = err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:357
		// _ = "end of CoverTab[31983]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:358
		_go_fuzz_dep_.CoverTab[31984]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:358
		// _ = "end of CoverTab[31984]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:358
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:358
	// _ = "end of CoverTab[31982]"
}

type link struct {
	t	Transformer
	// b[p:n] holds the bytes to be transformed by t.
	b	[]byte
	p	int
	n	int
}

func (l *link) src() []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:369
	_go_fuzz_dep_.CoverTab[31985]++
										return l.b[l.p:l.n]
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:370
	// _ = "end of CoverTab[31985]"
}

func (l *link) dst() []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:373
	_go_fuzz_dep_.CoverTab[31986]++
										return l.b[l.n:]
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:374
	// _ = "end of CoverTab[31986]"
}

// Chain returns a Transformer that applies t in sequence.
func Chain(t ...Transformer) Transformer {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:378
	_go_fuzz_dep_.CoverTab[31987]++
										if len(t) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:379
		_go_fuzz_dep_.CoverTab[31991]++
											return nop{}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:380
		// _ = "end of CoverTab[31991]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:381
		_go_fuzz_dep_.CoverTab[31992]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:381
		// _ = "end of CoverTab[31992]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:381
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:381
	// _ = "end of CoverTab[31987]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:381
	_go_fuzz_dep_.CoverTab[31988]++
										c := &chain{link: make([]link, len(t)+1)}
										for i, tt := range t {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:383
		_go_fuzz_dep_.CoverTab[31993]++
											c.link[i].t = tt
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:384
		// _ = "end of CoverTab[31993]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:385
	// _ = "end of CoverTab[31988]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:385
	_go_fuzz_dep_.CoverTab[31989]++

										b := make([][defaultBufSize]byte, len(t)-1)
										for i := range b {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:388
		_go_fuzz_dep_.CoverTab[31994]++
											c.link[i+1].b = b[i][:]
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:389
		// _ = "end of CoverTab[31994]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:390
	// _ = "end of CoverTab[31989]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:390
	_go_fuzz_dep_.CoverTab[31990]++
										return c
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:391
	// _ = "end of CoverTab[31990]"
}

// Reset resets the state of Chain. It calls Reset on all the Transformers.
func (c *chain) Reset() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:395
	_go_fuzz_dep_.CoverTab[31995]++
										for i, l := range c.link {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:396
		_go_fuzz_dep_.CoverTab[31996]++
											if l.t != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:397
			_go_fuzz_dep_.CoverTab[31998]++
												l.t.Reset()
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:398
			// _ = "end of CoverTab[31998]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:399
			_go_fuzz_dep_.CoverTab[31999]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:399
			// _ = "end of CoverTab[31999]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:399
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:399
		// _ = "end of CoverTab[31996]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:399
		_go_fuzz_dep_.CoverTab[31997]++
											c.link[i].p, c.link[i].n = 0, 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:400
		// _ = "end of CoverTab[31997]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:401
	// _ = "end of CoverTab[31995]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:406
// Transform applies the transformers of c in sequence.
func (c *chain) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:407
	_go_fuzz_dep_.CoverTab[32000]++

										srcL := &c.link[0]
										dstL := &c.link[len(c.link)-1]
										srcL.b, srcL.p, srcL.n = src, 0, len(src)
										dstL.b, dstL.n = dst, 0
										var lastFull, needProgress bool	// for detecting progress

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:420
	for low, i, high := c.errStart, c.errStart, len(c.link)-2; low <= i && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:420
		_go_fuzz_dep_.CoverTab[32003]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:420
		return i <= high
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:420
		// _ = "end of CoverTab[32003]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:420
	}(); {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:420
		_go_fuzz_dep_.CoverTab[32004]++
											in, out := &c.link[i], &c.link[i+1]
											nDst, nSrc, err0 := in.t.Transform(out.dst(), in.src(), atEOF && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:422
			_go_fuzz_dep_.CoverTab[32007]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:422
			return low == i
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:422
			// _ = "end of CoverTab[32007]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:422
		}())
											out.n += nDst
											in.p += nSrc
											if i > 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:425
			_go_fuzz_dep_.CoverTab[32008]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:425
			return in.p == in.n
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:425
			// _ = "end of CoverTab[32008]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:425
		}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:425
			_go_fuzz_dep_.CoverTab[32009]++
												in.p, in.n = 0, 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:426
			// _ = "end of CoverTab[32009]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:427
			_go_fuzz_dep_.CoverTab[32010]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:427
			// _ = "end of CoverTab[32010]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:427
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:427
		// _ = "end of CoverTab[32004]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:427
		_go_fuzz_dep_.CoverTab[32005]++
											needProgress, lastFull = lastFull, false
											switch err0 {
		case ErrShortDst:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:430
			_go_fuzz_dep_.CoverTab[32011]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:433
			if i == high {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:433
				_go_fuzz_dep_.CoverTab[32019]++
													return dstL.n, srcL.p, ErrShortDst
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:434
				// _ = "end of CoverTab[32019]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:435
				_go_fuzz_dep_.CoverTab[32020]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:435
				// _ = "end of CoverTab[32020]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:435
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:435
			// _ = "end of CoverTab[32011]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:435
			_go_fuzz_dep_.CoverTab[32012]++
												if out.n != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:436
				_go_fuzz_dep_.CoverTab[32021]++
													i++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:442
				lastFull = true
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:443
				// _ = "end of CoverTab[32021]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:444
				_go_fuzz_dep_.CoverTab[32022]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:444
				// _ = "end of CoverTab[32022]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:444
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:444
			// _ = "end of CoverTab[32012]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:444
			_go_fuzz_dep_.CoverTab[32013]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:447
			c.fatalError(i, errShortInternal)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:447
			// _ = "end of CoverTab[32013]"
		case ErrShortSrc:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:448
			_go_fuzz_dep_.CoverTab[32014]++
												if i == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:449
				_go_fuzz_dep_.CoverTab[32023]++

													err = ErrShortSrc
													break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:452
				// _ = "end of CoverTab[32023]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:453
				_go_fuzz_dep_.CoverTab[32024]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:453
				// _ = "end of CoverTab[32024]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:453
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:453
			// _ = "end of CoverTab[32014]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:453
			_go_fuzz_dep_.CoverTab[32015]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
			if needProgress && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				_go_fuzz_dep_.CoverTab[32025]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				return nSrc == 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				// _ = "end of CoverTab[32025]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
			}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				_go_fuzz_dep_.CoverTab[32026]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				return in.n-in.p == len(in.b)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				// _ = "end of CoverTab[32026]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:457
				_go_fuzz_dep_.CoverTab[32027]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:461
				c.fatalError(i, errShortInternal)
													break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:462
				// _ = "end of CoverTab[32027]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:463
				_go_fuzz_dep_.CoverTab[32028]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:463
				// _ = "end of CoverTab[32028]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:463
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:463
			// _ = "end of CoverTab[32015]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:463
			_go_fuzz_dep_.CoverTab[32016]++

												in.p, in.n = 0, copy(in.b, in.src())
												fallthrough
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:466
			// _ = "end of CoverTab[32016]"
		case nil:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:467
			_go_fuzz_dep_.CoverTab[32017]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:471
			if i > low {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:471
				_go_fuzz_dep_.CoverTab[32029]++
													i--
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:473
				// _ = "end of CoverTab[32029]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:474
				_go_fuzz_dep_.CoverTab[32030]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:474
				// _ = "end of CoverTab[32030]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:474
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:474
			// _ = "end of CoverTab[32017]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:475
			_go_fuzz_dep_.CoverTab[32018]++
												c.fatalError(i, err0)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:476
			// _ = "end of CoverTab[32018]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:477
		// _ = "end of CoverTab[32005]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:477
		_go_fuzz_dep_.CoverTab[32006]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:480
		i++
											low = i
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:481
		// _ = "end of CoverTab[32006]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:482
	// _ = "end of CoverTab[32000]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:482
	_go_fuzz_dep_.CoverTab[32001]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:487
	if c.errStart > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:487
		_go_fuzz_dep_.CoverTab[32031]++
											for i := 1; i < c.errStart; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:488
			_go_fuzz_dep_.CoverTab[32033]++
												c.link[i].p, c.link[i].n = 0, 0
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:489
			// _ = "end of CoverTab[32033]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:490
		// _ = "end of CoverTab[32031]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:490
		_go_fuzz_dep_.CoverTab[32032]++
											err, c.errStart, c.err = c.err, 0, nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:491
		// _ = "end of CoverTab[32032]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:492
		_go_fuzz_dep_.CoverTab[32034]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:492
		// _ = "end of CoverTab[32034]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:492
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:492
	// _ = "end of CoverTab[32001]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:492
	_go_fuzz_dep_.CoverTab[32002]++
										return dstL.n, srcL.p, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:493
	// _ = "end of CoverTab[32002]"
}

// Deprecated: Use runes.Remove instead.
func RemoveFunc(f func(r rune) bool) Transformer {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:497
	_go_fuzz_dep_.CoverTab[32035]++
										return removeF(f)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:498
	// _ = "end of CoverTab[32035]"
}

type removeF func(r rune) bool

func (removeF) Reset()	{ _go_fuzz_dep_.CoverTab[32036]++; // _ = "end of CoverTab[32036]" }

// Transform implements the Transformer interface.
func (t removeF) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:506
	_go_fuzz_dep_.CoverTab[32037]++
										for r, sz := rune(0), 0; len(src) > 0; src = src[sz:] {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:507
		_go_fuzz_dep_.CoverTab[32039]++

											if r = rune(src[0]); r < utf8.RuneSelf {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:509
			_go_fuzz_dep_.CoverTab[32042]++
												sz = 1
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:510
			// _ = "end of CoverTab[32042]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:511
			_go_fuzz_dep_.CoverTab[32043]++
												r, sz = utf8.DecodeRune(src)

												if sz == 1 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:514
				_go_fuzz_dep_.CoverTab[32044]++

													if !atEOF && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:516
					_go_fuzz_dep_.CoverTab[32047]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:516
					return !utf8.FullRune(src)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:516
					// _ = "end of CoverTab[32047]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:516
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:516
					_go_fuzz_dep_.CoverTab[32048]++
														err = ErrShortSrc
														break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:518
					// _ = "end of CoverTab[32048]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:519
					_go_fuzz_dep_.CoverTab[32049]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:519
					// _ = "end of CoverTab[32049]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:519
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:519
				// _ = "end of CoverTab[32044]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:519
				_go_fuzz_dep_.CoverTab[32045]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:524
				if !t(r) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:524
					_go_fuzz_dep_.CoverTab[32050]++
														if nDst+3 > len(dst) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:525
						_go_fuzz_dep_.CoverTab[32052]++
															err = ErrShortDst
															break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:527
						// _ = "end of CoverTab[32052]"
					} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:528
						_go_fuzz_dep_.CoverTab[32053]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:528
						// _ = "end of CoverTab[32053]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:528
					}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:528
					// _ = "end of CoverTab[32050]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:528
					_go_fuzz_dep_.CoverTab[32051]++
														nDst += copy(dst[nDst:], "\uFFFD")
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:529
					// _ = "end of CoverTab[32051]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:530
					_go_fuzz_dep_.CoverTab[32054]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:530
					// _ = "end of CoverTab[32054]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:530
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:530
				// _ = "end of CoverTab[32045]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:530
				_go_fuzz_dep_.CoverTab[32046]++
													nSrc++
													continue
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:532
				// _ = "end of CoverTab[32046]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:533
				_go_fuzz_dep_.CoverTab[32055]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:533
				// _ = "end of CoverTab[32055]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:533
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:533
			// _ = "end of CoverTab[32043]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:534
		// _ = "end of CoverTab[32039]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:534
		_go_fuzz_dep_.CoverTab[32040]++

											if !t(r) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:536
			_go_fuzz_dep_.CoverTab[32056]++
												if nDst+sz > len(dst) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:537
				_go_fuzz_dep_.CoverTab[32058]++
													err = ErrShortDst
													break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:539
				// _ = "end of CoverTab[32058]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:540
				_go_fuzz_dep_.CoverTab[32059]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:540
				// _ = "end of CoverTab[32059]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:540
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:540
			// _ = "end of CoverTab[32056]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:540
			_go_fuzz_dep_.CoverTab[32057]++
												nDst += copy(dst[nDst:], src[:sz])
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:541
			// _ = "end of CoverTab[32057]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:542
			_go_fuzz_dep_.CoverTab[32060]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:542
			// _ = "end of CoverTab[32060]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:542
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:542
		// _ = "end of CoverTab[32040]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:542
		_go_fuzz_dep_.CoverTab[32041]++
											nSrc += sz
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:543
		// _ = "end of CoverTab[32041]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:544
	// _ = "end of CoverTab[32037]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:544
	_go_fuzz_dep_.CoverTab[32038]++
										return
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:545
	// _ = "end of CoverTab[32038]"
}

// grow returns a new []byte that is longer than b, and copies the first n bytes
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:548
// of b to the start of the new slice.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:550
func grow(b []byte, n int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:550
	_go_fuzz_dep_.CoverTab[32061]++
										m := len(b)
										if m <= 32 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:552
		_go_fuzz_dep_.CoverTab[32063]++
											m = 64
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:553
		// _ = "end of CoverTab[32063]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:554
		_go_fuzz_dep_.CoverTab[32064]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:554
		if m <= 256 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:554
			_go_fuzz_dep_.CoverTab[32065]++
												m *= 2
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:555
			// _ = "end of CoverTab[32065]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:556
			_go_fuzz_dep_.CoverTab[32066]++
												m += m >> 1
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:557
			// _ = "end of CoverTab[32066]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:558
		// _ = "end of CoverTab[32064]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:558
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:558
	// _ = "end of CoverTab[32061]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:558
	_go_fuzz_dep_.CoverTab[32062]++
										buf := make([]byte, m)
										copy(buf, b[:n])
										return buf
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:561
	// _ = "end of CoverTab[32062]"
}

const initialBufSize = 128

// String returns a string with the result of converting s[:n] using t, where
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:566
// n <= len(s). If err == nil, n will be len(s). It calls Reset on t.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:568
func String(t Transformer, s string) (result string, n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:568
	_go_fuzz_dep_.CoverTab[32067]++
										t.Reset()
										if s == "" {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:570
		_go_fuzz_dep_.CoverTab[32072]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:573
		if _, _, err := t.Transform(nil, nil, true); err == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:573
			_go_fuzz_dep_.CoverTab[32073]++
												return "", 0, nil
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:574
			// _ = "end of CoverTab[32073]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:575
			_go_fuzz_dep_.CoverTab[32074]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:575
			// _ = "end of CoverTab[32074]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:575
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:575
		// _ = "end of CoverTab[32072]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:576
		_go_fuzz_dep_.CoverTab[32075]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:576
		// _ = "end of CoverTab[32075]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:576
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:576
	// _ = "end of CoverTab[32067]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:576
	_go_fuzz_dep_.CoverTab[32068]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:580
	buf := [2 * initialBufSize]byte{}
										dst := buf[:initialBufSize:initialBufSize]
										src := buf[initialBufSize : 2*initialBufSize]

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:587
	nDst, nSrc := 0, 0
										pDst, pSrc := 0, 0

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:596
	pPrefix := 0
	for {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:597
		_go_fuzz_dep_.CoverTab[32076]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:600
		n := copy(src, s[pSrc:])
											nDst, nSrc, err = t.Transform(dst, src[:n], pSrc+n == len(s))
											pDst += nDst
											pSrc += nSrc

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:607
		if !bytes.Equal(dst[:nDst], src[:nSrc]) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:607
			_go_fuzz_dep_.CoverTab[32078]++
												break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:608
			// _ = "end of CoverTab[32078]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:609
			_go_fuzz_dep_.CoverTab[32079]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:609
			// _ = "end of CoverTab[32079]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:609
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:609
		// _ = "end of CoverTab[32076]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:609
		_go_fuzz_dep_.CoverTab[32077]++
											pPrefix = pSrc
											if err == ErrShortDst {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:611
			_go_fuzz_dep_.CoverTab[32080]++

												break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:613
			// _ = "end of CoverTab[32080]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:614
			_go_fuzz_dep_.CoverTab[32081]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:614
			if err == ErrShortSrc {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:614
				_go_fuzz_dep_.CoverTab[32082]++
													if nSrc == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:615
					_go_fuzz_dep_.CoverTab[32083]++

														break
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:617
					// _ = "end of CoverTab[32083]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:618
					_go_fuzz_dep_.CoverTab[32084]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:618
					// _ = "end of CoverTab[32084]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:618
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:618
				// _ = "end of CoverTab[32082]"

			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
				_go_fuzz_dep_.CoverTab[32085]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
				if err != nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
					_go_fuzz_dep_.CoverTab[32086]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
					return pPrefix == len(s)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
					// _ = "end of CoverTab[32086]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:620
					_go_fuzz_dep_.CoverTab[32087]++
														return string(s[:pPrefix]), pPrefix, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:621
					// _ = "end of CoverTab[32087]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
					_go_fuzz_dep_.CoverTab[32088]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
					// _ = "end of CoverTab[32088]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
				// _ = "end of CoverTab[32085]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
			// _ = "end of CoverTab[32081]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:622
		// _ = "end of CoverTab[32077]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:623
	// _ = "end of CoverTab[32068]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:623
	_go_fuzz_dep_.CoverTab[32069]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:631
	if pPrefix != 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:631
		_go_fuzz_dep_.CoverTab[32089]++
											newDst := dst
											if pDst > len(newDst) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:633
			_go_fuzz_dep_.CoverTab[32091]++
												newDst = make([]byte, len(s)+nDst-nSrc)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:634
			// _ = "end of CoverTab[32091]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:635
			_go_fuzz_dep_.CoverTab[32092]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:635
			// _ = "end of CoverTab[32092]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:635
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:635
		// _ = "end of CoverTab[32089]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:635
		_go_fuzz_dep_.CoverTab[32090]++
											copy(newDst[pPrefix:pDst], dst[:nDst])
											copy(newDst[:pPrefix], s[:pPrefix])
											dst = newDst
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:638
		// _ = "end of CoverTab[32090]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:639
		_go_fuzz_dep_.CoverTab[32093]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:639
		// _ = "end of CoverTab[32093]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:639
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:639
	// _ = "end of CoverTab[32069]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:639
	_go_fuzz_dep_.CoverTab[32070]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
	if (err == nil && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
		_go_fuzz_dep_.CoverTab[32094]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
		return pSrc == len(s)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
		// _ = "end of CoverTab[32094]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
	}()) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
		_go_fuzz_dep_.CoverTab[32095]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:643
		return (err != nil && func() bool {
												_go_fuzz_dep_.CoverTab[32096]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
			return err != ErrShortDst
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
			// _ = "end of CoverTab[32096]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
		}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
			_go_fuzz_dep_.CoverTab[32097]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
			return err != ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
			// _ = "end of CoverTab[32097]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
		}())
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
		// _ = "end of CoverTab[32095]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:644
		_go_fuzz_dep_.CoverTab[32098]++
											return string(dst[:pDst]), pSrc, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:645
		// _ = "end of CoverTab[32098]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:646
		_go_fuzz_dep_.CoverTab[32099]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:646
		// _ = "end of CoverTab[32099]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:646
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:646
	// _ = "end of CoverTab[32070]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:646
	_go_fuzz_dep_.CoverTab[32071]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:649
	for {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:649
		_go_fuzz_dep_.CoverTab[32100]++
											n := copy(src, s[pSrc:])
											atEOF := pSrc+n == len(s)
											nDst, nSrc, err := t.Transform(dst[pDst:], src[:n], atEOF)
											pDst += nDst
											pSrc += nSrc

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:658
		if err == ErrShortDst {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:658
			_go_fuzz_dep_.CoverTab[32101]++
												if nDst == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:659
				_go_fuzz_dep_.CoverTab[32102]++
													dst = grow(dst, pDst)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:660
				// _ = "end of CoverTab[32102]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:661
				_go_fuzz_dep_.CoverTab[32103]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:661
				// _ = "end of CoverTab[32103]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:661
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:661
			// _ = "end of CoverTab[32101]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:662
			_go_fuzz_dep_.CoverTab[32104]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:662
			if err == ErrShortSrc {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:662
				_go_fuzz_dep_.CoverTab[32105]++
													if atEOF {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:663
					_go_fuzz_dep_.CoverTab[32107]++
														return string(dst[:pDst]), pSrc, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:664
					// _ = "end of CoverTab[32107]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:665
					_go_fuzz_dep_.CoverTab[32108]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:665
					// _ = "end of CoverTab[32108]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:665
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:665
				// _ = "end of CoverTab[32105]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:665
				_go_fuzz_dep_.CoverTab[32106]++
													if nSrc == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:666
					_go_fuzz_dep_.CoverTab[32109]++
														src = grow(src, 0)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:667
					// _ = "end of CoverTab[32109]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:668
					_go_fuzz_dep_.CoverTab[32110]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:668
					// _ = "end of CoverTab[32110]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:668
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:668
				// _ = "end of CoverTab[32106]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
				_go_fuzz_dep_.CoverTab[32111]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
				if err != nil || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
					_go_fuzz_dep_.CoverTab[32112]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
					return pSrc == len(s)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
					// _ = "end of CoverTab[32112]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
				}() {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:669
					_go_fuzz_dep_.CoverTab[32113]++
														return string(dst[:pDst]), pSrc, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:670
					// _ = "end of CoverTab[32113]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
					_go_fuzz_dep_.CoverTab[32114]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
					// _ = "end of CoverTab[32114]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
				}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
				// _ = "end of CoverTab[32111]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
			}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
			// _ = "end of CoverTab[32104]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:671
		// _ = "end of CoverTab[32100]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:672
	// _ = "end of CoverTab[32071]"
}

// Bytes returns a new byte slice with the result of converting b[:n] using t,
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:675
// where n <= len(b). If err == nil, n will be len(b). It calls Reset on t.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:677
func Bytes(t Transformer, b []byte) (result []byte, n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:677
	_go_fuzz_dep_.CoverTab[32115]++
										return doAppend(t, 0, make([]byte, len(b)), b)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:678
	// _ = "end of CoverTab[32115]"
}

// Append appends the result of converting src[:n] using t to dst, where
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:681
// n <= len(src), If err == nil, n will be len(src). It calls Reset on t.
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:683
func Append(t Transformer, dst, src []byte) (result []byte, n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:683
	_go_fuzz_dep_.CoverTab[32116]++
										if len(dst) == cap(dst) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:684
		_go_fuzz_dep_.CoverTab[32118]++
											n := len(src) + len(dst)
											b := make([]byte, n)
											dst = b[:copy(b, dst)]
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:687
		// _ = "end of CoverTab[32118]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:688
		_go_fuzz_dep_.CoverTab[32119]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:688
		// _ = "end of CoverTab[32119]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:688
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:688
	// _ = "end of CoverTab[32116]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:688
	_go_fuzz_dep_.CoverTab[32117]++
										return doAppend(t, len(dst), dst[:cap(dst)], src)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:689
	// _ = "end of CoverTab[32117]"
}

func doAppend(t Transformer, pDst int, dst, src []byte) (result []byte, n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:692
	_go_fuzz_dep_.CoverTab[32120]++
										t.Reset()
										pSrc := 0
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:695
		_go_fuzz_dep_.CoverTab[32121]++
											nDst, nSrc, err := t.Transform(dst[pDst:], src[pSrc:], true)
											pDst += nDst
											pSrc += nSrc
											if err != ErrShortDst {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:699
			_go_fuzz_dep_.CoverTab[32123]++
												return dst[:pDst], pSrc, err
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:700
			// _ = "end of CoverTab[32123]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:701
			_go_fuzz_dep_.CoverTab[32124]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:701
			// _ = "end of CoverTab[32124]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:701
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:701
		// _ = "end of CoverTab[32121]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:701
		_go_fuzz_dep_.CoverTab[32122]++

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:705
		if nDst == 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:705
			_go_fuzz_dep_.CoverTab[32125]++
												dst = grow(dst, pDst)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:706
			// _ = "end of CoverTab[32125]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:707
			_go_fuzz_dep_.CoverTab[32126]++
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:707
			// _ = "end of CoverTab[32126]"
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:707
		}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:707
		// _ = "end of CoverTab[32122]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:708
	// _ = "end of CoverTab[32120]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:709
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/transform/transform.go:709
var _ = _go_fuzz_dep_.CoverTab
