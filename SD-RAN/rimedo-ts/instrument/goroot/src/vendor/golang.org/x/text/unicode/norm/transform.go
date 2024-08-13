// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:5
)

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

// Reset implements the Reset method of the transform.Transformer interface.
func (Form) Reset()	{ _go_fuzz_dep_.CoverTab[33931]++; // _ = "end of CoverTab[33931]" }

// Transform implements the Transform method of the transform.Transformer
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:16
// interface. It may need to write segments of up to MaxSegmentSize at once.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:16
// Users should either catch ErrShortDst and allow dst to grow or have dst be at
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:16
// least of size MaxTransformChunkSize to be guaranteed of progress.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:20
func (f Form) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:20
	_go_fuzz_dep_.CoverTab[33932]++

										b := src
										eof := atEOF
										if ns := len(dst); ns < len(b) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:24
		_go_fuzz_dep_.CoverTab[33936]++
											err = transform.ErrShortDst
											eof = false
											b = b[:ns]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:27
		// _ = "end of CoverTab[33936]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:28
		_go_fuzz_dep_.CoverTab[33937]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:28
		// _ = "end of CoverTab[33937]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:28
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:28
	// _ = "end of CoverTab[33932]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:28
	_go_fuzz_dep_.CoverTab[33933]++
										i, ok := formTable[f].quickSpan(inputBytes(b), 0, len(b), eof)
										n := copy(dst, b[:i])
										if !ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:31
		_go_fuzz_dep_.CoverTab[33938]++
											nDst, nSrc, err = f.transform(dst[n:], src[n:], atEOF)
											return nDst + n, nSrc + n, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:33
		// _ = "end of CoverTab[33938]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:34
		_go_fuzz_dep_.CoverTab[33939]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:34
		// _ = "end of CoverTab[33939]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:34
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:34
	// _ = "end of CoverTab[33933]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:34
	_go_fuzz_dep_.CoverTab[33934]++

										if err == nil && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		_go_fuzz_dep_.CoverTab[33940]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		return n < len(src)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		// _ = "end of CoverTab[33940]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
	}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		_go_fuzz_dep_.CoverTab[33941]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		return !atEOF
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		// _ = "end of CoverTab[33941]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
	}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:36
		_go_fuzz_dep_.CoverTab[33942]++
											err = transform.ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:37
		// _ = "end of CoverTab[33942]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:38
		_go_fuzz_dep_.CoverTab[33943]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:38
		// _ = "end of CoverTab[33943]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:38
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:38
	// _ = "end of CoverTab[33934]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:38
	_go_fuzz_dep_.CoverTab[33935]++
										return n, n, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:39
	// _ = "end of CoverTab[33935]"
}

func flushTransform(rb *reorderBuffer) bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:42
	_go_fuzz_dep_.CoverTab[33944]++

										if len(rb.out) < rb.nrune*utf8.UTFMax {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:44
		_go_fuzz_dep_.CoverTab[33946]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:45
		// _ = "end of CoverTab[33946]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:46
		_go_fuzz_dep_.CoverTab[33947]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:46
		// _ = "end of CoverTab[33947]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:46
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:46
	// _ = "end of CoverTab[33944]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:46
	_go_fuzz_dep_.CoverTab[33945]++
										rb.out = rb.out[rb.flushCopy(rb.out):]
										return true
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:48
	// _ = "end of CoverTab[33945]"
}

var errs = []error{nil, transform.ErrShortDst, transform.ErrShortSrc}

// transform implements the transform.Transformer interface. It is only called
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:53
// when quickSpan does not pass for a given string.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:55
func (f Form) transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:55
	_go_fuzz_dep_.CoverTab[33948]++

										rb := reorderBuffer{}
										rb.init(f, src)
										for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:59
		_go_fuzz_dep_.CoverTab[33949]++

											rb.setFlusher(dst[nDst:], flushTransform)
											end := decomposeSegment(&rb, nSrc, atEOF)
											if end < 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:63
			_go_fuzz_dep_.CoverTab[33952]++
												return nDst, nSrc, errs[-end]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:64
			// _ = "end of CoverTab[33952]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:65
			_go_fuzz_dep_.CoverTab[33953]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:65
			// _ = "end of CoverTab[33953]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:65
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:65
		// _ = "end of CoverTab[33949]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:65
		_go_fuzz_dep_.CoverTab[33950]++
											nDst = len(dst) - len(rb.out)
											nSrc = end

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:70
		end = rb.nsrc
		eof := atEOF
		if n := nSrc + len(dst) - nDst; n < end {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:72
			_go_fuzz_dep_.CoverTab[33954]++
												err = transform.ErrShortDst
												end = n
												eof = false
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:75
			// _ = "end of CoverTab[33954]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:76
			_go_fuzz_dep_.CoverTab[33955]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:76
			// _ = "end of CoverTab[33955]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:76
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:76
		// _ = "end of CoverTab[33950]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:76
		_go_fuzz_dep_.CoverTab[33951]++
											end, ok := rb.f.quickSpan(rb.src, nSrc, end, eof)
											n := copy(dst[nDst:], rb.src.bytes[nSrc:end])
											nSrc += n
											nDst += n
											if ok {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:81
			_go_fuzz_dep_.CoverTab[33956]++
												if err == nil && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				_go_fuzz_dep_.CoverTab[33958]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				return n < rb.nsrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				// _ = "end of CoverTab[33958]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
			}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				_go_fuzz_dep_.CoverTab[33959]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				return !atEOF
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				// _ = "end of CoverTab[33959]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
			}() {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:82
				_go_fuzz_dep_.CoverTab[33960]++
													err = transform.ErrShortSrc
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:83
				// _ = "end of CoverTab[33960]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:84
				_go_fuzz_dep_.CoverTab[33961]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:84
				// _ = "end of CoverTab[33961]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:84
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:84
			// _ = "end of CoverTab[33956]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:84
			_go_fuzz_dep_.CoverTab[33957]++
												return nDst, nSrc, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:85
			// _ = "end of CoverTab[33957]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:86
			_go_fuzz_dep_.CoverTab[33962]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:86
			// _ = "end of CoverTab[33962]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:86
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:86
		// _ = "end of CoverTab[33951]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:87
	// _ = "end of CoverTab[33948]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/transform.go:88
var _ = _go_fuzz_dep_.CoverTab
