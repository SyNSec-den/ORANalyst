// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:5
)

import "io"

type normWriter struct {
	rb	reorderBuffer
	w	io.Writer
	buf	[]byte
}

// Write implements the standard write interface.  If the last characters are
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:15
// not at a normalization boundary, the bytes will be buffered for the next
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:15
// write. The remaining bytes will be written on close.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:18
func (w *normWriter) Write(data []byte) (n int, err error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:18
	_go_fuzz_dep_.CoverTab[33667]++
	// Process data in pieces to keep w.buf size bounded.
	const chunk = 4000

	for len(data) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:22
		_go_fuzz_dep_.CoverTab[33669]++

												m := len(data)
												if m > chunk {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:25
			_go_fuzz_dep_.CoverTab[33672]++
													m = chunk
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:26
			// _ = "end of CoverTab[33672]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:27
			_go_fuzz_dep_.CoverTab[33673]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:27
			// _ = "end of CoverTab[33673]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:27
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:27
		// _ = "end of CoverTab[33669]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:27
		_go_fuzz_dep_.CoverTab[33670]++
												w.rb.src = inputBytes(data[:m])
												w.rb.nsrc = m
												w.buf = doAppend(&w.rb, w.buf, 0)
												data = data[m:]
												n += m

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:36
		i := lastBoundary(&w.rb.f, w.buf)
		if i == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:37
			_go_fuzz_dep_.CoverTab[33674]++
													i = 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:38
			// _ = "end of CoverTab[33674]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:39
			_go_fuzz_dep_.CoverTab[33675]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:39
			// _ = "end of CoverTab[33675]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:39
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:39
		// _ = "end of CoverTab[33670]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:39
		_go_fuzz_dep_.CoverTab[33671]++
												if i > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:40
			_go_fuzz_dep_.CoverTab[33676]++
													if _, err = w.w.Write(w.buf[:i]); err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:41
				_go_fuzz_dep_.CoverTab[33678]++
														break
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:42
				// _ = "end of CoverTab[33678]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:43
				_go_fuzz_dep_.CoverTab[33679]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:43
				// _ = "end of CoverTab[33679]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:43
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:43
			// _ = "end of CoverTab[33676]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:43
			_go_fuzz_dep_.CoverTab[33677]++
													bn := copy(w.buf, w.buf[i:])
													w.buf = w.buf[:bn]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:45
			// _ = "end of CoverTab[33677]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:46
			_go_fuzz_dep_.CoverTab[33680]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:46
			// _ = "end of CoverTab[33680]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:46
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:46
		// _ = "end of CoverTab[33671]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:47
	// _ = "end of CoverTab[33667]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:47
	_go_fuzz_dep_.CoverTab[33668]++
											return n, err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:48
	// _ = "end of CoverTab[33668]"
}

// Close forces data that remains in the buffer to be written.
func (w *normWriter) Close() error {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:52
	_go_fuzz_dep_.CoverTab[33681]++
											if len(w.buf) > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:53
		_go_fuzz_dep_.CoverTab[33683]++
												_, err := w.w.Write(w.buf)
												if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:55
			_go_fuzz_dep_.CoverTab[33684]++
													return err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:56
			// _ = "end of CoverTab[33684]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:57
			_go_fuzz_dep_.CoverTab[33685]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:57
			// _ = "end of CoverTab[33685]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:57
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:57
		// _ = "end of CoverTab[33683]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:58
		_go_fuzz_dep_.CoverTab[33686]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:58
		// _ = "end of CoverTab[33686]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:58
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:58
	// _ = "end of CoverTab[33681]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:58
	_go_fuzz_dep_.CoverTab[33682]++
											return nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:59
	// _ = "end of CoverTab[33682]"
}

// Writer returns a new writer that implements Write(b)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:62
// by writing f(b) to w. The returned writer may use an
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:62
// internal buffer to maintain state across Write calls.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:62
// Calling its Close method writes any buffered data to w.
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:66
func (f Form) Writer(w io.Writer) io.WriteCloser {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:66
	_go_fuzz_dep_.CoverTab[33687]++
											wr := &normWriter{rb: reorderBuffer{}, w: w}
											wr.rb.init(f, nil)
											return wr
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:69
	// _ = "end of CoverTab[33687]"
}

type normReader struct {
	rb		reorderBuffer
	r		io.Reader
	inbuf		[]byte
	outbuf		[]byte
	bufStart	int
	lastBoundary	int
	err		error
}

// Read implements the standard read interface.
func (r *normReader) Read(p []byte) (int, error) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:83
	_go_fuzz_dep_.CoverTab[33688]++
											for {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:84
		_go_fuzz_dep_.CoverTab[33689]++
												if r.lastBoundary-r.bufStart > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:85
			_go_fuzz_dep_.CoverTab[33693]++
													n := copy(p, r.outbuf[r.bufStart:r.lastBoundary])
													r.bufStart += n
													if r.lastBoundary-r.bufStart > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:88
				_go_fuzz_dep_.CoverTab[33695]++
														return n, nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:89
				// _ = "end of CoverTab[33695]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:90
				_go_fuzz_dep_.CoverTab[33696]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:90
				// _ = "end of CoverTab[33696]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:90
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:90
			// _ = "end of CoverTab[33693]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:90
			_go_fuzz_dep_.CoverTab[33694]++
													return n, r.err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:91
			// _ = "end of CoverTab[33694]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:92
			_go_fuzz_dep_.CoverTab[33697]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:92
			// _ = "end of CoverTab[33697]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:92
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:92
		// _ = "end of CoverTab[33689]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:92
		_go_fuzz_dep_.CoverTab[33690]++
												if r.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:93
			_go_fuzz_dep_.CoverTab[33698]++
													return 0, r.err
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:94
			// _ = "end of CoverTab[33698]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:95
			_go_fuzz_dep_.CoverTab[33699]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:95
			// _ = "end of CoverTab[33699]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:95
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:95
		// _ = "end of CoverTab[33690]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:95
		_go_fuzz_dep_.CoverTab[33691]++
												outn := copy(r.outbuf, r.outbuf[r.lastBoundary:])
												r.outbuf = r.outbuf[0:outn]
												r.bufStart = 0

												n, err := r.r.Read(r.inbuf)
												r.rb.src = inputBytes(r.inbuf[0:n])
												r.rb.nsrc, r.err = n, err
												if n > 0 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:103
			_go_fuzz_dep_.CoverTab[33700]++
													r.outbuf = doAppend(&r.rb, r.outbuf, 0)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:104
			// _ = "end of CoverTab[33700]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:105
			_go_fuzz_dep_.CoverTab[33701]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:105
			// _ = "end of CoverTab[33701]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:105
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:105
		// _ = "end of CoverTab[33691]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:105
		_go_fuzz_dep_.CoverTab[33692]++
												if err == io.EOF {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:106
			_go_fuzz_dep_.CoverTab[33702]++
													r.lastBoundary = len(r.outbuf)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:107
			// _ = "end of CoverTab[33702]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:108
			_go_fuzz_dep_.CoverTab[33703]++
													r.lastBoundary = lastBoundary(&r.rb.f, r.outbuf)
													if r.lastBoundary == -1 {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:110
				_go_fuzz_dep_.CoverTab[33704]++
														r.lastBoundary = 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:111
				// _ = "end of CoverTab[33704]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:112
				_go_fuzz_dep_.CoverTab[33705]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:112
				// _ = "end of CoverTab[33705]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:112
			}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:112
			// _ = "end of CoverTab[33703]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:113
		// _ = "end of CoverTab[33692]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:114
	// _ = "end of CoverTab[33688]"
}

// Reader returns a new reader that implements Read
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:117
// by reading data from r and returning f(data).
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:119
func (f Form) Reader(r io.Reader) io.Reader {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:119
	_go_fuzz_dep_.CoverTab[33706]++
											const chunk = 4000
											buf := make([]byte, chunk)
											rr := &normReader{rb: reorderBuffer{}, r: r, inbuf: buf}
											rr.rb.init(f, buf)
											return rr
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:124
	// _ = "end of CoverTab[33706]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/readwriter.go:125
var _ = _go_fuzz_dep_.CoverTab
