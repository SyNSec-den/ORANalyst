// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:5
)

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

// Reset implements the Reset method of the transform.Transformer interface.
func (Form) Reset()	{ _go_fuzz_dep_.CoverTab[71212]++; // _ = "end of CoverTab[71212]" }

// Transform implements the Transform method of the transform.Transformer
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:16
// interface. It may need to write segments of up to MaxSegmentSize at once.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:16
// Users should either catch ErrShortDst and allow dst to grow or have dst be at
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:16
// least of size MaxTransformChunkSize to be guaranteed of progress.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:20
func (f Form) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:20
	_go_fuzz_dep_.CoverTab[71213]++

												b := src
												eof := atEOF
												if ns := len(dst); ns < len(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:24
		_go_fuzz_dep_.CoverTab[71217]++
													err = transform.ErrShortDst
													eof = false
													b = b[:ns]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:27
		// _ = "end of CoverTab[71217]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:28
		_go_fuzz_dep_.CoverTab[71218]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:28
		// _ = "end of CoverTab[71218]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:28
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:28
	// _ = "end of CoverTab[71213]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:28
	_go_fuzz_dep_.CoverTab[71214]++
												i, ok := formTable[f].quickSpan(inputBytes(b), 0, len(b), eof)
												n := copy(dst, b[:i])
												if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:31
		_go_fuzz_dep_.CoverTab[71219]++
													nDst, nSrc, err = f.transform(dst[n:], src[n:], atEOF)
													return nDst + n, nSrc + n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:33
		// _ = "end of CoverTab[71219]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:34
		_go_fuzz_dep_.CoverTab[71220]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:34
		// _ = "end of CoverTab[71220]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:34
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:34
	// _ = "end of CoverTab[71214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:34
	_go_fuzz_dep_.CoverTab[71215]++

												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		_go_fuzz_dep_.CoverTab[71221]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		return n < len(src)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		// _ = "end of CoverTab[71221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		_go_fuzz_dep_.CoverTab[71222]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		return !atEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		// _ = "end of CoverTab[71222]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:36
		_go_fuzz_dep_.CoverTab[71223]++
													err = transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:37
		// _ = "end of CoverTab[71223]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:38
		_go_fuzz_dep_.CoverTab[71224]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:38
		// _ = "end of CoverTab[71224]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:38
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:38
	// _ = "end of CoverTab[71215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:38
	_go_fuzz_dep_.CoverTab[71216]++
												return n, n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:39
	// _ = "end of CoverTab[71216]"
}

func flushTransform(rb *reorderBuffer) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:42
	_go_fuzz_dep_.CoverTab[71225]++

												if len(rb.out) < rb.nrune*utf8.UTFMax {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:44
		_go_fuzz_dep_.CoverTab[71227]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:45
		// _ = "end of CoverTab[71227]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:46
		_go_fuzz_dep_.CoverTab[71228]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:46
		// _ = "end of CoverTab[71228]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:46
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:46
	// _ = "end of CoverTab[71225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:46
	_go_fuzz_dep_.CoverTab[71226]++
												rb.out = rb.out[rb.flushCopy(rb.out):]
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:48
	// _ = "end of CoverTab[71226]"
}

var errs = []error{nil, transform.ErrShortDst, transform.ErrShortSrc}

// transform implements the transform.Transformer interface. It is only called
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:53
// when quickSpan does not pass for a given string.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:55
func (f Form) transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:55
	_go_fuzz_dep_.CoverTab[71229]++

												rb := reorderBuffer{}
												rb.init(f, src)
												for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:59
		_go_fuzz_dep_.CoverTab[71230]++

													rb.setFlusher(dst[nDst:], flushTransform)
													end := decomposeSegment(&rb, nSrc, atEOF)
													if end < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:63
			_go_fuzz_dep_.CoverTab[71233]++
														return nDst, nSrc, errs[-end]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:64
			// _ = "end of CoverTab[71233]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:65
			_go_fuzz_dep_.CoverTab[71234]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:65
			// _ = "end of CoverTab[71234]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:65
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:65
		// _ = "end of CoverTab[71230]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:65
		_go_fuzz_dep_.CoverTab[71231]++
													nDst = len(dst) - len(rb.out)
													nSrc = end

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:70
		end = rb.nsrc
		eof := atEOF
		if n := nSrc + len(dst) - nDst; n < end {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:72
			_go_fuzz_dep_.CoverTab[71235]++
														err = transform.ErrShortDst
														end = n
														eof = false
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:75
			// _ = "end of CoverTab[71235]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:76
			_go_fuzz_dep_.CoverTab[71236]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:76
			// _ = "end of CoverTab[71236]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:76
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:76
		// _ = "end of CoverTab[71231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:76
		_go_fuzz_dep_.CoverTab[71232]++
													end, ok := rb.f.quickSpan(rb.src, nSrc, end, eof)
													n := copy(dst[nDst:], rb.src.bytes[nSrc:end])
													nSrc += n
													nDst += n
													if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:81
			_go_fuzz_dep_.CoverTab[71237]++
														if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				_go_fuzz_dep_.CoverTab[71239]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				return n < rb.nsrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				// _ = "end of CoverTab[71239]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				_go_fuzz_dep_.CoverTab[71240]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				return !atEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				// _ = "end of CoverTab[71240]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:82
				_go_fuzz_dep_.CoverTab[71241]++
															err = transform.ErrShortSrc
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:83
				// _ = "end of CoverTab[71241]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:84
				_go_fuzz_dep_.CoverTab[71242]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:84
				// _ = "end of CoverTab[71242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:84
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:84
			// _ = "end of CoverTab[71237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:84
			_go_fuzz_dep_.CoverTab[71238]++
														return nDst, nSrc, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:85
			// _ = "end of CoverTab[71238]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:86
			_go_fuzz_dep_.CoverTab[71243]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:86
			// _ = "end of CoverTab[71243]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:86
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:86
		// _ = "end of CoverTab[71232]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:87
	// _ = "end of CoverTab[71229]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/transform.go:88
var _ = _go_fuzz_dep_.CoverTab
