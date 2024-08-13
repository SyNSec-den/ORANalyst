// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:5
// Package transform provides reader and writer wrappers that transform the
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:5
// bytes passing through as well as various transformations. Example
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:5
// transformations provided by other packages include normalization and
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:5
// conversion between character sets.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
package transform

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:9
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:67
// that determines how much of the input already conforms to the Transformer.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:69
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:102
// Reset method.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:104
type NopResetter struct{}

// Reset implements the Reset method of the Transformer interface.
func (NopResetter) Reset()	{ _go_fuzz_dep_.CoverTab[69178]++; // _ = "end of CoverTab[69178]" }

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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:132
// via t. It calls Reset on t.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:134
func NewReader(r io.Reader, t Transformer) *Reader {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:134
	_go_fuzz_dep_.CoverTab[69179]++
											t.Reset()
											return &Reader{
		r:	r,
		t:	t,
		dst:	make([]byte, defaultBufSize),
		src:	make([]byte, defaultBufSize),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:141
	// _ = "end of CoverTab[69179]"
}

// Read implements the io.Reader interface.
func (r *Reader) Read(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:145
	_go_fuzz_dep_.CoverTab[69180]++
											n, err := 0, error(nil)
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:147
		_go_fuzz_dep_.CoverTab[69181]++

												if r.dst0 != r.dst1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:149
			_go_fuzz_dep_.CoverTab[69185]++
													n = copy(p, r.dst[r.dst0:r.dst1])
													r.dst0 += n
													if r.dst0 == r.dst1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:152
				_go_fuzz_dep_.CoverTab[69187]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:152
				return r.transformComplete
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:152
				// _ = "end of CoverTab[69187]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:152
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:152
				_go_fuzz_dep_.CoverTab[69188]++
														return n, r.err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:153
				// _ = "end of CoverTab[69188]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:154
				_go_fuzz_dep_.CoverTab[69189]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:154
				// _ = "end of CoverTab[69189]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:154
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:154
			// _ = "end of CoverTab[69185]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:154
			_go_fuzz_dep_.CoverTab[69186]++
													return n, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:155
			// _ = "end of CoverTab[69186]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:156
			_go_fuzz_dep_.CoverTab[69190]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:156
			if r.transformComplete {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:156
				_go_fuzz_dep_.CoverTab[69191]++
														return 0, r.err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:157
				// _ = "end of CoverTab[69191]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
				_go_fuzz_dep_.CoverTab[69192]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
				// _ = "end of CoverTab[69192]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
			// _ = "end of CoverTab[69190]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
		// _ = "end of CoverTab[69181]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:158
		_go_fuzz_dep_.CoverTab[69182]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:164
		if r.src0 != r.src1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:164
			_go_fuzz_dep_.CoverTab[69193]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:164
			return r.err != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:164
			// _ = "end of CoverTab[69193]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:164
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:164
			_go_fuzz_dep_.CoverTab[69194]++
													r.dst0 = 0
													r.dst1, n, err = r.t.Transform(r.dst, r.src[r.src0:r.src1], r.err == io.EOF)
													r.src0 += n

													switch {
			case err == nil:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:170
				_go_fuzz_dep_.CoverTab[69195]++
														if r.src0 != r.src1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:171
					_go_fuzz_dep_.CoverTab[69201]++
															r.err = errInconsistentByteCount
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:172
					// _ = "end of CoverTab[69201]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:173
					_go_fuzz_dep_.CoverTab[69202]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:173
					// _ = "end of CoverTab[69202]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:173
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:173
				// _ = "end of CoverTab[69195]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:173
				_go_fuzz_dep_.CoverTab[69196]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:176
				r.transformComplete = r.err != nil
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:177
				// _ = "end of CoverTab[69196]"
			case err == ErrShortDst && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
				_go_fuzz_dep_.CoverTab[69203]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
				return (r.dst1 != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
					_go_fuzz_dep_.CoverTab[69204]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
					return n != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
					// _ = "end of CoverTab[69204]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
				}())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
				// _ = "end of CoverTab[69203]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
			}():
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:178
				_go_fuzz_dep_.CoverTab[69197]++

														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:180
				// _ = "end of CoverTab[69197]"
			case err == ErrShortSrc && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				_go_fuzz_dep_.CoverTab[69205]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				return r.src1-r.src0 != len(r.src)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				// _ = "end of CoverTab[69205]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				_go_fuzz_dep_.CoverTab[69206]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				return r.err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				// _ = "end of CoverTab[69206]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
			}():
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				_go_fuzz_dep_.CoverTab[69198]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:181
				// _ = "end of CoverTab[69198]"

			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:183
				_go_fuzz_dep_.CoverTab[69199]++
														r.transformComplete = true

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:187
				if r.err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:187
					_go_fuzz_dep_.CoverTab[69207]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:187
					return r.err == io.EOF
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:187
					// _ = "end of CoverTab[69207]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:187
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:187
					_go_fuzz_dep_.CoverTab[69208]++
															r.err = err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:188
					// _ = "end of CoverTab[69208]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:189
					_go_fuzz_dep_.CoverTab[69209]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:189
					// _ = "end of CoverTab[69209]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:189
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:189
				// _ = "end of CoverTab[69199]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:189
				_go_fuzz_dep_.CoverTab[69200]++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:190
				// _ = "end of CoverTab[69200]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:191
			// _ = "end of CoverTab[69194]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:192
			_go_fuzz_dep_.CoverTab[69210]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:192
			// _ = "end of CoverTab[69210]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:192
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:192
		// _ = "end of CoverTab[69182]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:192
		_go_fuzz_dep_.CoverTab[69183]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:196
		if r.src0 != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:196
			_go_fuzz_dep_.CoverTab[69211]++
													r.src0, r.src1 = 0, copy(r.src, r.src[r.src0:r.src1])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:197
			// _ = "end of CoverTab[69211]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:198
			_go_fuzz_dep_.CoverTab[69212]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:198
			// _ = "end of CoverTab[69212]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:198
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:198
		// _ = "end of CoverTab[69183]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:198
		_go_fuzz_dep_.CoverTab[69184]++
												n, r.err = r.r.Read(r.src[r.src1:])
												r.src1 += n
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:200
		// _ = "end of CoverTab[69184]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:201
	// _ = "end of CoverTab[69180]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:206
// Writer wraps another io.Writer by transforming the bytes read.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:206
// The user needs to call Close to flush unwritten bytes that may
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:206
// be buffered.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:209
type Writer struct {
	w	io.Writer
	t	Transformer
	dst	[]byte

	// src[:n] contains bytes that have not yet passed through t.
	src	[]byte
	n	int
}

// NewWriter returns a new Writer that wraps w by transforming the bytes written
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:219
// via t. It calls Reset on t.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:221
func NewWriter(w io.Writer, t Transformer) *Writer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:221
	_go_fuzz_dep_.CoverTab[69213]++
											t.Reset()
											return &Writer{
		w:	w,
		t:	t,
		dst:	make([]byte, defaultBufSize),
		src:	make([]byte, defaultBufSize),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:228
	// _ = "end of CoverTab[69213]"
}

// Write implements the io.Writer interface. If there are not enough
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:231
// bytes available to complete a Transform, the bytes will be buffered
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:231
// for the next write. Call Close to convert the remaining bytes.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:234
func (w *Writer) Write(data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:234
	_go_fuzz_dep_.CoverTab[69214]++
											src := data
											if w.n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:236
		_go_fuzz_dep_.CoverTab[69216]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:239
		n = copy(w.src[w.n:], data)
												w.n += n
												src = w.src[:w.n]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:241
		// _ = "end of CoverTab[69216]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:242
		_go_fuzz_dep_.CoverTab[69217]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:242
		// _ = "end of CoverTab[69217]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:242
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:242
	// _ = "end of CoverTab[69214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:242
	_go_fuzz_dep_.CoverTab[69215]++
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:243
		_go_fuzz_dep_.CoverTab[69218]++
												nDst, nSrc, err := w.t.Transform(w.dst, src, false)
												if _, werr := w.w.Write(w.dst[:nDst]); werr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:245
			_go_fuzz_dep_.CoverTab[69222]++
													return n, werr
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:246
			// _ = "end of CoverTab[69222]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:247
			_go_fuzz_dep_.CoverTab[69223]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:247
			// _ = "end of CoverTab[69223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:247
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:247
		// _ = "end of CoverTab[69218]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:247
		_go_fuzz_dep_.CoverTab[69219]++
												src = src[nSrc:]
												if w.n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:249
			_go_fuzz_dep_.CoverTab[69224]++
													n += nSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:250
			// _ = "end of CoverTab[69224]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:251
			_go_fuzz_dep_.CoverTab[69225]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:251
			if len(src) <= n {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:251
				_go_fuzz_dep_.CoverTab[69226]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:254
				w.n = 0
				n -= len(src)
				src = data[n:]
				if n < len(data) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
					_go_fuzz_dep_.CoverTab[69227]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
					return (err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
						_go_fuzz_dep_.CoverTab[69228]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
						return err == ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
						// _ = "end of CoverTab[69228]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
					}())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
					// _ = "end of CoverTab[69227]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:257
					_go_fuzz_dep_.CoverTab[69229]++
															continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:258
					// _ = "end of CoverTab[69229]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:259
					_go_fuzz_dep_.CoverTab[69230]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:259
					// _ = "end of CoverTab[69230]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:259
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:259
				// _ = "end of CoverTab[69226]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
				_go_fuzz_dep_.CoverTab[69231]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
				// _ = "end of CoverTab[69231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
			// _ = "end of CoverTab[69225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
		// _ = "end of CoverTab[69219]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:260
		_go_fuzz_dep_.CoverTab[69220]++
												switch err {
		case ErrShortDst:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:262
			_go_fuzz_dep_.CoverTab[69232]++

													if nDst > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:264
				_go_fuzz_dep_.CoverTab[69236]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:264
				return nSrc > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:264
				// _ = "end of CoverTab[69236]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:264
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:264
				_go_fuzz_dep_.CoverTab[69237]++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:265
				// _ = "end of CoverTab[69237]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:266
				_go_fuzz_dep_.CoverTab[69238]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:266
				// _ = "end of CoverTab[69238]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:266
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:266
			// _ = "end of CoverTab[69232]"
		case ErrShortSrc:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:267
			_go_fuzz_dep_.CoverTab[69233]++
													if len(src) < len(w.src) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:268
				_go_fuzz_dep_.CoverTab[69239]++
														m := copy(w.src, src)

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:272
				if w.n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:272
					_go_fuzz_dep_.CoverTab[69241]++
															n += m
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:273
					// _ = "end of CoverTab[69241]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:274
					_go_fuzz_dep_.CoverTab[69242]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:274
					// _ = "end of CoverTab[69242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:274
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:274
				// _ = "end of CoverTab[69239]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:274
				_go_fuzz_dep_.CoverTab[69240]++
														w.n = m
														err = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:276
				// _ = "end of CoverTab[69240]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
				_go_fuzz_dep_.CoverTab[69243]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
				if nDst > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
					_go_fuzz_dep_.CoverTab[69244]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
					return nSrc > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
					// _ = "end of CoverTab[69244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:277
					_go_fuzz_dep_.CoverTab[69245]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:284
					continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:284
					// _ = "end of CoverTab[69245]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:285
					_go_fuzz_dep_.CoverTab[69246]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:285
					// _ = "end of CoverTab[69246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:285
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:285
				// _ = "end of CoverTab[69243]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:285
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:285
			// _ = "end of CoverTab[69233]"
		case nil:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:286
			_go_fuzz_dep_.CoverTab[69234]++
													if w.n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:287
				_go_fuzz_dep_.CoverTab[69247]++
														err = errInconsistentByteCount
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:288
				// _ = "end of CoverTab[69247]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
				_go_fuzz_dep_.CoverTab[69248]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
				// _ = "end of CoverTab[69248]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
			// _ = "end of CoverTab[69234]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
			_go_fuzz_dep_.CoverTab[69235]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:289
			// _ = "end of CoverTab[69235]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:290
		// _ = "end of CoverTab[69220]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:290
		_go_fuzz_dep_.CoverTab[69221]++
												return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:291
		// _ = "end of CoverTab[69221]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:292
	// _ = "end of CoverTab[69215]"
}

// Close implements the io.Closer interface.
func (w *Writer) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:296
	_go_fuzz_dep_.CoverTab[69249]++
											src := w.src[:w.n]
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:298
		_go_fuzz_dep_.CoverTab[69250]++
												nDst, nSrc, err := w.t.Transform(w.dst, src, true)
												if _, werr := w.w.Write(w.dst[:nDst]); werr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:300
			_go_fuzz_dep_.CoverTab[69253]++
													return werr
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:301
			// _ = "end of CoverTab[69253]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:302
			_go_fuzz_dep_.CoverTab[69254]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:302
			// _ = "end of CoverTab[69254]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:302
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:302
		// _ = "end of CoverTab[69250]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:302
		_go_fuzz_dep_.CoverTab[69251]++
												if err != ErrShortDst {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:303
			_go_fuzz_dep_.CoverTab[69255]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:304
			// _ = "end of CoverTab[69255]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:305
			_go_fuzz_dep_.CoverTab[69256]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:305
			// _ = "end of CoverTab[69256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:305
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:305
		// _ = "end of CoverTab[69251]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:305
		_go_fuzz_dep_.CoverTab[69252]++
												src = src[nSrc:]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:306
		// _ = "end of CoverTab[69252]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:307
	// _ = "end of CoverTab[69249]"
}

type nop struct{ NopResetter }

func (nop) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:312
	_go_fuzz_dep_.CoverTab[69257]++
											n := copy(dst, src)
											if n < len(src) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:314
		_go_fuzz_dep_.CoverTab[69259]++
												err = ErrShortDst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:315
		// _ = "end of CoverTab[69259]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:316
		_go_fuzz_dep_.CoverTab[69260]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:316
		// _ = "end of CoverTab[69260]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:316
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:316
	// _ = "end of CoverTab[69257]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:316
	_go_fuzz_dep_.CoverTab[69258]++
											return n, n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:317
	// _ = "end of CoverTab[69258]"
}

func (nop) Span(src []byte, atEOF bool) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:320
	_go_fuzz_dep_.CoverTab[69261]++
											return len(src), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:321
	// _ = "end of CoverTab[69261]"
}

type discard struct{ NopResetter }

func (discard) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:326
	_go_fuzz_dep_.CoverTab[69262]++
											return 0, len(src), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:327
	// _ = "end of CoverTab[69262]"
}

var (
	// Discard is a Transformer for which all Transform calls succeed
	// by consuming all bytes and writing nothing.
	Discard	Transformer	= discard{}

	// Nop is a SpanningTransformer that copies src to dst.
	Nop	SpanningTransformer	= nop{}
)

// chain is a sequence of links. A chain with N Transformers has N+1 links and
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:339
// N+1 buffers. Of those N+1 buffers, the first and last are the src and dst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:339
// buffers given to chain.Transform and the middle N-1 buffers are intermediate
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:339
// buffers owned by the chain. The i'th link transforms bytes from the i'th
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:339
// buffer chain.link[i].b at read offset chain.link[i].p to the i+1'th buffer
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:339
// chain.link[i+1].b at write offset chain.link[i+1].n, for i in [0, N).
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:345
type chain struct {
	link	[]link
	err	error
	// errStart is the index at which the error occurred plus 1. Processing
	// errStart at this level at the next call to Transform. As long as
	// errStart > 0, chain will not consume any more source bytes.
	errStart	int
}

func (c *chain) fatalError(errIndex int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:354
	_go_fuzz_dep_.CoverTab[69263]++
											if i := errIndex + 1; i > c.errStart {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:355
		_go_fuzz_dep_.CoverTab[69264]++
												c.errStart = i
												c.err = err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:357
		// _ = "end of CoverTab[69264]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:358
		_go_fuzz_dep_.CoverTab[69265]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:358
		// _ = "end of CoverTab[69265]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:358
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:358
	// _ = "end of CoverTab[69263]"
}

type link struct {
	t	Transformer
	// b[p:n] holds the bytes to be transformed by t.
	b	[]byte
	p	int
	n	int
}

func (l *link) src() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:369
	_go_fuzz_dep_.CoverTab[69266]++
											return l.b[l.p:l.n]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:370
	// _ = "end of CoverTab[69266]"
}

func (l *link) dst() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:373
	_go_fuzz_dep_.CoverTab[69267]++
											return l.b[l.n:]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:374
	// _ = "end of CoverTab[69267]"
}

// Chain returns a Transformer that applies t in sequence.
func Chain(t ...Transformer) Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:378
	_go_fuzz_dep_.CoverTab[69268]++
											if len(t) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:379
		_go_fuzz_dep_.CoverTab[69272]++
												return nop{}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:380
		// _ = "end of CoverTab[69272]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:381
		_go_fuzz_dep_.CoverTab[69273]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:381
		// _ = "end of CoverTab[69273]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:381
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:381
	// _ = "end of CoverTab[69268]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:381
	_go_fuzz_dep_.CoverTab[69269]++
											c := &chain{link: make([]link, len(t)+1)}
											for i, tt := range t {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:383
		_go_fuzz_dep_.CoverTab[69274]++
												c.link[i].t = tt
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:384
		// _ = "end of CoverTab[69274]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:385
	// _ = "end of CoverTab[69269]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:385
	_go_fuzz_dep_.CoverTab[69270]++

											b := make([][defaultBufSize]byte, len(t)-1)
											for i := range b {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:388
		_go_fuzz_dep_.CoverTab[69275]++
												c.link[i+1].b = b[i][:]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:389
		// _ = "end of CoverTab[69275]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:390
	// _ = "end of CoverTab[69270]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:390
	_go_fuzz_dep_.CoverTab[69271]++
											return c
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:391
	// _ = "end of CoverTab[69271]"
}

// Reset resets the state of Chain. It calls Reset on all the Transformers.
func (c *chain) Reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:395
	_go_fuzz_dep_.CoverTab[69276]++
											for i, l := range c.link {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:396
		_go_fuzz_dep_.CoverTab[69277]++
												if l.t != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:397
			_go_fuzz_dep_.CoverTab[69279]++
													l.t.Reset()
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:398
			// _ = "end of CoverTab[69279]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:399
			_go_fuzz_dep_.CoverTab[69280]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:399
			// _ = "end of CoverTab[69280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:399
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:399
		// _ = "end of CoverTab[69277]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:399
		_go_fuzz_dep_.CoverTab[69278]++
												c.link[i].p, c.link[i].n = 0, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:400
		// _ = "end of CoverTab[69278]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:401
	// _ = "end of CoverTab[69276]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:406
// Transform applies the transformers of c in sequence.
func (c *chain) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:407
	_go_fuzz_dep_.CoverTab[69281]++

											srcL := &c.link[0]
											dstL := &c.link[len(c.link)-1]
											srcL.b, srcL.p, srcL.n = src, 0, len(src)
											dstL.b, dstL.n = dst, 0
											var lastFull, needProgress bool	// for detecting progress

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:420
	for low, i, high := c.errStart, c.errStart, len(c.link)-2; low <= i && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:420
		_go_fuzz_dep_.CoverTab[69284]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:420
		return i <= high
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:420
		// _ = "end of CoverTab[69284]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:420
	}(); {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:420
		_go_fuzz_dep_.CoverTab[69285]++
												in, out := &c.link[i], &c.link[i+1]
												nDst, nSrc, err0 := in.t.Transform(out.dst(), in.src(), atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:422
			_go_fuzz_dep_.CoverTab[69288]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:422
			return low == i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:422
			// _ = "end of CoverTab[69288]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:422
		}())
												out.n += nDst
												in.p += nSrc
												if i > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:425
			_go_fuzz_dep_.CoverTab[69289]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:425
			return in.p == in.n
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:425
			// _ = "end of CoverTab[69289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:425
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:425
			_go_fuzz_dep_.CoverTab[69290]++
													in.p, in.n = 0, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:426
			// _ = "end of CoverTab[69290]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:427
			_go_fuzz_dep_.CoverTab[69291]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:427
			// _ = "end of CoverTab[69291]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:427
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:427
		// _ = "end of CoverTab[69285]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:427
		_go_fuzz_dep_.CoverTab[69286]++
												needProgress, lastFull = lastFull, false
												switch err0 {
		case ErrShortDst:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:430
			_go_fuzz_dep_.CoverTab[69292]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:433
			if i == high {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:433
				_go_fuzz_dep_.CoverTab[69300]++
														return dstL.n, srcL.p, ErrShortDst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:434
				// _ = "end of CoverTab[69300]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:435
				_go_fuzz_dep_.CoverTab[69301]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:435
				// _ = "end of CoverTab[69301]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:435
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:435
			// _ = "end of CoverTab[69292]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:435
			_go_fuzz_dep_.CoverTab[69293]++
													if out.n != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:436
				_go_fuzz_dep_.CoverTab[69302]++
														i++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:442
				lastFull = true
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:443
				// _ = "end of CoverTab[69302]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:444
				_go_fuzz_dep_.CoverTab[69303]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:444
				// _ = "end of CoverTab[69303]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:444
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:444
			// _ = "end of CoverTab[69293]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:444
			_go_fuzz_dep_.CoverTab[69294]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:447
			c.fatalError(i, errShortInternal)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:447
			// _ = "end of CoverTab[69294]"
		case ErrShortSrc:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:448
			_go_fuzz_dep_.CoverTab[69295]++
													if i == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:449
				_go_fuzz_dep_.CoverTab[69304]++

														err = ErrShortSrc
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:452
				// _ = "end of CoverTab[69304]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:453
				_go_fuzz_dep_.CoverTab[69305]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:453
				// _ = "end of CoverTab[69305]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:453
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:453
			// _ = "end of CoverTab[69295]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:453
			_go_fuzz_dep_.CoverTab[69296]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
			if needProgress && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				_go_fuzz_dep_.CoverTab[69306]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				return nSrc == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				// _ = "end of CoverTab[69306]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				_go_fuzz_dep_.CoverTab[69307]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				return in.n-in.p == len(in.b)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				// _ = "end of CoverTab[69307]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:457
				_go_fuzz_dep_.CoverTab[69308]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:461
				c.fatalError(i, errShortInternal)
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:462
				// _ = "end of CoverTab[69308]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:463
				_go_fuzz_dep_.CoverTab[69309]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:463
				// _ = "end of CoverTab[69309]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:463
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:463
			// _ = "end of CoverTab[69296]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:463
			_go_fuzz_dep_.CoverTab[69297]++

													in.p, in.n = 0, copy(in.b, in.src())
													fallthrough
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:466
			// _ = "end of CoverTab[69297]"
		case nil:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:467
			_go_fuzz_dep_.CoverTab[69298]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:471
			if i > low {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:471
				_go_fuzz_dep_.CoverTab[69310]++
														i--
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:473
				// _ = "end of CoverTab[69310]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:474
				_go_fuzz_dep_.CoverTab[69311]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:474
				// _ = "end of CoverTab[69311]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:474
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:474
			// _ = "end of CoverTab[69298]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:475
			_go_fuzz_dep_.CoverTab[69299]++
													c.fatalError(i, err0)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:476
			// _ = "end of CoverTab[69299]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:477
		// _ = "end of CoverTab[69286]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:477
		_go_fuzz_dep_.CoverTab[69287]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:480
		i++
												low = i
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:481
		// _ = "end of CoverTab[69287]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:482
	// _ = "end of CoverTab[69281]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:482
	_go_fuzz_dep_.CoverTab[69282]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:487
	if c.errStart > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:487
		_go_fuzz_dep_.CoverTab[69312]++
												for i := 1; i < c.errStart; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:488
			_go_fuzz_dep_.CoverTab[69314]++
													c.link[i].p, c.link[i].n = 0, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:489
			// _ = "end of CoverTab[69314]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:490
		// _ = "end of CoverTab[69312]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:490
		_go_fuzz_dep_.CoverTab[69313]++
												err, c.errStart, c.err = c.err, 0, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:491
		// _ = "end of CoverTab[69313]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:492
		_go_fuzz_dep_.CoverTab[69315]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:492
		// _ = "end of CoverTab[69315]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:492
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:492
	// _ = "end of CoverTab[69282]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:492
	_go_fuzz_dep_.CoverTab[69283]++
											return dstL.n, srcL.p, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:493
	// _ = "end of CoverTab[69283]"
}

// Deprecated: Use runes.Remove instead.
func RemoveFunc(f func(r rune) bool) Transformer {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:497
	_go_fuzz_dep_.CoverTab[69316]++
											return removeF(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:498
	// _ = "end of CoverTab[69316]"
}

type removeF func(r rune) bool

func (removeF) Reset()	{ _go_fuzz_dep_.CoverTab[69317]++; // _ = "end of CoverTab[69317]" }

// Transform implements the Transformer interface.
func (t removeF) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:506
	_go_fuzz_dep_.CoverTab[69318]++
											for r, sz := rune(0), 0; len(src) > 0; src = src[sz:] {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:507
		_go_fuzz_dep_.CoverTab[69320]++

												if r = rune(src[0]); r < utf8.RuneSelf {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:509
			_go_fuzz_dep_.CoverTab[69323]++
													sz = 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:510
			// _ = "end of CoverTab[69323]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:511
			_go_fuzz_dep_.CoverTab[69324]++
													r, sz = utf8.DecodeRune(src)

													if sz == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:514
				_go_fuzz_dep_.CoverTab[69325]++

														if !atEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:516
					_go_fuzz_dep_.CoverTab[69328]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:516
					return !utf8.FullRune(src)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:516
					// _ = "end of CoverTab[69328]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:516
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:516
					_go_fuzz_dep_.CoverTab[69329]++
															err = ErrShortSrc
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:518
					// _ = "end of CoverTab[69329]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:519
					_go_fuzz_dep_.CoverTab[69330]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:519
					// _ = "end of CoverTab[69330]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:519
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:519
				// _ = "end of CoverTab[69325]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:519
				_go_fuzz_dep_.CoverTab[69326]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:524
				if !t(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:524
					_go_fuzz_dep_.CoverTab[69331]++
															if nDst+3 > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:525
						_go_fuzz_dep_.CoverTab[69333]++
																err = ErrShortDst
																break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:527
						// _ = "end of CoverTab[69333]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:528
						_go_fuzz_dep_.CoverTab[69334]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:528
						// _ = "end of CoverTab[69334]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:528
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:528
					// _ = "end of CoverTab[69331]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:528
					_go_fuzz_dep_.CoverTab[69332]++
															nDst += copy(dst[nDst:], "\uFFFD")
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:529
					// _ = "end of CoverTab[69332]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:530
					_go_fuzz_dep_.CoverTab[69335]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:530
					// _ = "end of CoverTab[69335]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:530
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:530
				// _ = "end of CoverTab[69326]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:530
				_go_fuzz_dep_.CoverTab[69327]++
														nSrc++
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:532
				// _ = "end of CoverTab[69327]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:533
				_go_fuzz_dep_.CoverTab[69336]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:533
				// _ = "end of CoverTab[69336]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:533
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:533
			// _ = "end of CoverTab[69324]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:534
		// _ = "end of CoverTab[69320]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:534
		_go_fuzz_dep_.CoverTab[69321]++

												if !t(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:536
			_go_fuzz_dep_.CoverTab[69337]++
													if nDst+sz > len(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:537
				_go_fuzz_dep_.CoverTab[69339]++
														err = ErrShortDst
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:539
				// _ = "end of CoverTab[69339]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:540
				_go_fuzz_dep_.CoverTab[69340]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:540
				// _ = "end of CoverTab[69340]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:540
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:540
			// _ = "end of CoverTab[69337]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:540
			_go_fuzz_dep_.CoverTab[69338]++
													nDst += copy(dst[nDst:], src[:sz])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:541
			// _ = "end of CoverTab[69338]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:542
			_go_fuzz_dep_.CoverTab[69341]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:542
			// _ = "end of CoverTab[69341]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:542
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:542
		// _ = "end of CoverTab[69321]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:542
		_go_fuzz_dep_.CoverTab[69322]++
												nSrc += sz
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:543
		// _ = "end of CoverTab[69322]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:544
	// _ = "end of CoverTab[69318]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:544
	_go_fuzz_dep_.CoverTab[69319]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:545
	// _ = "end of CoverTab[69319]"
}

// grow returns a new []byte that is longer than b, and copies the first n bytes
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:548
// of b to the start of the new slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:550
func grow(b []byte, n int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:550
	_go_fuzz_dep_.CoverTab[69342]++
											m := len(b)
											if m <= 32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:552
		_go_fuzz_dep_.CoverTab[69344]++
												m = 64
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:553
		// _ = "end of CoverTab[69344]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:554
		_go_fuzz_dep_.CoverTab[69345]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:554
		if m <= 256 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:554
			_go_fuzz_dep_.CoverTab[69346]++
													m *= 2
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:555
			// _ = "end of CoverTab[69346]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:556
			_go_fuzz_dep_.CoverTab[69347]++
													m += m >> 1
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:557
			// _ = "end of CoverTab[69347]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:558
		// _ = "end of CoverTab[69345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:558
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:558
	// _ = "end of CoverTab[69342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:558
	_go_fuzz_dep_.CoverTab[69343]++
											buf := make([]byte, m)
											copy(buf, b[:n])
											return buf
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:561
	// _ = "end of CoverTab[69343]"
}

const initialBufSize = 128

// String returns a string with the result of converting s[:n] using t, where
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:566
// n <= len(s). If err == nil, n will be len(s). It calls Reset on t.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:568
func String(t Transformer, s string) (result string, n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:568
	_go_fuzz_dep_.CoverTab[69348]++
											t.Reset()
											if s == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:570
		_go_fuzz_dep_.CoverTab[69353]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:573
		if _, _, err := t.Transform(nil, nil, true); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:573
			_go_fuzz_dep_.CoverTab[69354]++
													return "", 0, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:574
			// _ = "end of CoverTab[69354]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:575
			_go_fuzz_dep_.CoverTab[69355]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:575
			// _ = "end of CoverTab[69355]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:575
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:575
		// _ = "end of CoverTab[69353]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:576
		_go_fuzz_dep_.CoverTab[69356]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:576
		// _ = "end of CoverTab[69356]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:576
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:576
	// _ = "end of CoverTab[69348]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:576
	_go_fuzz_dep_.CoverTab[69349]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:580
	buf := [2 * initialBufSize]byte{}
											dst := buf[:initialBufSize:initialBufSize]
											src := buf[initialBufSize : 2*initialBufSize]

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:587
	nDst, nSrc := 0, 0
											pDst, pSrc := 0, 0

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:596
	pPrefix := 0
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:597
		_go_fuzz_dep_.CoverTab[69357]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:600
		n := copy(src, s[pSrc:])
												nDst, nSrc, err = t.Transform(dst, src[:n], pSrc+n == len(s))
												pDst += nDst
												pSrc += nSrc

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:607
		if !bytes.Equal(dst[:nDst], src[:nSrc]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:607
			_go_fuzz_dep_.CoverTab[69359]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:608
			// _ = "end of CoverTab[69359]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:609
			_go_fuzz_dep_.CoverTab[69360]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:609
			// _ = "end of CoverTab[69360]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:609
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:609
		// _ = "end of CoverTab[69357]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:609
		_go_fuzz_dep_.CoverTab[69358]++
												pPrefix = pSrc
												if err == ErrShortDst {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:611
			_go_fuzz_dep_.CoverTab[69361]++

													break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:613
			// _ = "end of CoverTab[69361]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:614
			_go_fuzz_dep_.CoverTab[69362]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:614
			if err == ErrShortSrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:614
				_go_fuzz_dep_.CoverTab[69363]++
														if nSrc == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:615
					_go_fuzz_dep_.CoverTab[69364]++

															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:617
					// _ = "end of CoverTab[69364]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:618
					_go_fuzz_dep_.CoverTab[69365]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:618
					// _ = "end of CoverTab[69365]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:618
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:618
				// _ = "end of CoverTab[69363]"

			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
				_go_fuzz_dep_.CoverTab[69366]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
				if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
					_go_fuzz_dep_.CoverTab[69367]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
					return pPrefix == len(s)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
					// _ = "end of CoverTab[69367]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:620
					_go_fuzz_dep_.CoverTab[69368]++
															return string(s[:pPrefix]), pPrefix, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:621
					// _ = "end of CoverTab[69368]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
					_go_fuzz_dep_.CoverTab[69369]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
					// _ = "end of CoverTab[69369]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
				// _ = "end of CoverTab[69366]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
			// _ = "end of CoverTab[69362]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:622
		// _ = "end of CoverTab[69358]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:623
	// _ = "end of CoverTab[69349]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:623
	_go_fuzz_dep_.CoverTab[69350]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:631
	if pPrefix != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:631
		_go_fuzz_dep_.CoverTab[69370]++
												newDst := dst
												if pDst > len(newDst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:633
			_go_fuzz_dep_.CoverTab[69372]++
													newDst = make([]byte, len(s)+nDst-nSrc)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:634
			// _ = "end of CoverTab[69372]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:635
			_go_fuzz_dep_.CoverTab[69373]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:635
			// _ = "end of CoverTab[69373]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:635
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:635
		// _ = "end of CoverTab[69370]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:635
		_go_fuzz_dep_.CoverTab[69371]++
												copy(newDst[pPrefix:pDst], dst[:nDst])
												copy(newDst[:pPrefix], s[:pPrefix])
												dst = newDst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:638
		// _ = "end of CoverTab[69371]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:639
		_go_fuzz_dep_.CoverTab[69374]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:639
		// _ = "end of CoverTab[69374]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:639
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:639
	// _ = "end of CoverTab[69350]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:639
	_go_fuzz_dep_.CoverTab[69351]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
	if (err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
		_go_fuzz_dep_.CoverTab[69375]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
		return pSrc == len(s)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
		// _ = "end of CoverTab[69375]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
		_go_fuzz_dep_.CoverTab[69376]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:643
		return (err != nil && func() bool {
													_go_fuzz_dep_.CoverTab[69377]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
			return err != ErrShortDst
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
			// _ = "end of CoverTab[69377]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
			_go_fuzz_dep_.CoverTab[69378]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
			return err != ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
			// _ = "end of CoverTab[69378]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
		// _ = "end of CoverTab[69376]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:644
		_go_fuzz_dep_.CoverTab[69379]++
												return string(dst[:pDst]), pSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:645
		// _ = "end of CoverTab[69379]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:646
		_go_fuzz_dep_.CoverTab[69380]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:646
		// _ = "end of CoverTab[69380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:646
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:646
	// _ = "end of CoverTab[69351]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:646
	_go_fuzz_dep_.CoverTab[69352]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:649
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:649
		_go_fuzz_dep_.CoverTab[69381]++
												n := copy(src, s[pSrc:])
												atEOF := pSrc+n == len(s)
												nDst, nSrc, err := t.Transform(dst[pDst:], src[:n], atEOF)
												pDst += nDst
												pSrc += nSrc

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:658
		if err == ErrShortDst {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:658
			_go_fuzz_dep_.CoverTab[69382]++
													if nDst == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:659
				_go_fuzz_dep_.CoverTab[69383]++
														dst = grow(dst, pDst)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:660
				// _ = "end of CoverTab[69383]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:661
				_go_fuzz_dep_.CoverTab[69384]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:661
				// _ = "end of CoverTab[69384]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:661
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:661
			// _ = "end of CoverTab[69382]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:662
			_go_fuzz_dep_.CoverTab[69385]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:662
			if err == ErrShortSrc {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:662
				_go_fuzz_dep_.CoverTab[69386]++
														if atEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:663
					_go_fuzz_dep_.CoverTab[69388]++
															return string(dst[:pDst]), pSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:664
					// _ = "end of CoverTab[69388]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:665
					_go_fuzz_dep_.CoverTab[69389]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:665
					// _ = "end of CoverTab[69389]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:665
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:665
				// _ = "end of CoverTab[69386]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:665
				_go_fuzz_dep_.CoverTab[69387]++
														if nSrc == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:666
					_go_fuzz_dep_.CoverTab[69390]++
															src = grow(src, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:667
					// _ = "end of CoverTab[69390]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:668
					_go_fuzz_dep_.CoverTab[69391]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:668
					// _ = "end of CoverTab[69391]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:668
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:668
				// _ = "end of CoverTab[69387]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
				_go_fuzz_dep_.CoverTab[69392]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
				if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
					_go_fuzz_dep_.CoverTab[69393]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
					return pSrc == len(s)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
					// _ = "end of CoverTab[69393]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:669
					_go_fuzz_dep_.CoverTab[69394]++
															return string(dst[:pDst]), pSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:670
					// _ = "end of CoverTab[69394]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
					_go_fuzz_dep_.CoverTab[69395]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
					// _ = "end of CoverTab[69395]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
				// _ = "end of CoverTab[69392]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
			// _ = "end of CoverTab[69385]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:671
		// _ = "end of CoverTab[69381]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:672
	// _ = "end of CoverTab[69352]"
}

// Bytes returns a new byte slice with the result of converting b[:n] using t,
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:675
// where n <= len(b). If err == nil, n will be len(b). It calls Reset on t.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:677
func Bytes(t Transformer, b []byte) (result []byte, n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:677
	_go_fuzz_dep_.CoverTab[69396]++
											return doAppend(t, 0, make([]byte, len(b)), b)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:678
	// _ = "end of CoverTab[69396]"
}

// Append appends the result of converting src[:n] using t to dst, where
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:681
// n <= len(src), If err == nil, n will be len(src). It calls Reset on t.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:683
func Append(t Transformer, dst, src []byte) (result []byte, n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:683
	_go_fuzz_dep_.CoverTab[69397]++
											if len(dst) == cap(dst) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:684
		_go_fuzz_dep_.CoverTab[69399]++
												n := len(src) + len(dst)
												b := make([]byte, n)
												dst = b[:copy(b, dst)]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:687
		// _ = "end of CoverTab[69399]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:688
		_go_fuzz_dep_.CoverTab[69400]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:688
		// _ = "end of CoverTab[69400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:688
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:688
	// _ = "end of CoverTab[69397]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:688
	_go_fuzz_dep_.CoverTab[69398]++
											return doAppend(t, len(dst), dst[:cap(dst)], src)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:689
	// _ = "end of CoverTab[69398]"
}

func doAppend(t Transformer, pDst int, dst, src []byte) (result []byte, n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:692
	_go_fuzz_dep_.CoverTab[69401]++
											t.Reset()
											pSrc := 0
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:695
		_go_fuzz_dep_.CoverTab[69402]++
												nDst, nSrc, err := t.Transform(dst[pDst:], src[pSrc:], true)
												pDst += nDst
												pSrc += nSrc
												if err != ErrShortDst {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:699
			_go_fuzz_dep_.CoverTab[69404]++
													return dst[:pDst], pSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:700
			// _ = "end of CoverTab[69404]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:701
			_go_fuzz_dep_.CoverTab[69405]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:701
			// _ = "end of CoverTab[69405]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:701
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:701
		// _ = "end of CoverTab[69402]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:701
		_go_fuzz_dep_.CoverTab[69403]++

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:705
		if nDst == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:705
			_go_fuzz_dep_.CoverTab[69406]++
													dst = grow(dst, pDst)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:706
			// _ = "end of CoverTab[69406]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:707
			_go_fuzz_dep_.CoverTab[69407]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:707
			// _ = "end of CoverTab[69407]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:707
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:707
		// _ = "end of CoverTab[69403]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:708
	// _ = "end of CoverTab[69401]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:709
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/transform/transform.go:709
var _ = _go_fuzz_dep_.CoverTab
