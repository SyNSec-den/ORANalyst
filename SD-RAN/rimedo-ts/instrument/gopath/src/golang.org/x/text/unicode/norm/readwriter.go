// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:5
)

import "io"

type normWriter struct {
	rb	reorderBuffer
	w	io.Writer
	buf	[]byte
}

// Write implements the standard write interface.  If the last characters are
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:15
// not at a normalization boundary, the bytes will be buffered for the next
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:15
// write. The remaining bytes will be written on close.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:18
func (w *normWriter) Write(data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:18
	_go_fuzz_dep_.CoverTab[70948]++
	// Process data in pieces to keep w.buf size bounded.
	const chunk = 4000

	for len(data) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:22
		_go_fuzz_dep_.CoverTab[70950]++

													m := len(data)
													if m > chunk {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:25
			_go_fuzz_dep_.CoverTab[70953]++
														m = chunk
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:26
			// _ = "end of CoverTab[70953]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:27
			_go_fuzz_dep_.CoverTab[70954]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:27
			// _ = "end of CoverTab[70954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:27
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:27
		// _ = "end of CoverTab[70950]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:27
		_go_fuzz_dep_.CoverTab[70951]++
													w.rb.src = inputBytes(data[:m])
													w.rb.nsrc = m
													w.buf = doAppend(&w.rb, w.buf, 0)
													data = data[m:]
													n += m

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:36
		i := lastBoundary(&w.rb.f, w.buf)
		if i == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:37
			_go_fuzz_dep_.CoverTab[70955]++
														i = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:38
			// _ = "end of CoverTab[70955]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:39
			_go_fuzz_dep_.CoverTab[70956]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:39
			// _ = "end of CoverTab[70956]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:39
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:39
		// _ = "end of CoverTab[70951]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:39
		_go_fuzz_dep_.CoverTab[70952]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:40
			_go_fuzz_dep_.CoverTab[70957]++
														if _, err = w.w.Write(w.buf[:i]); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:41
				_go_fuzz_dep_.CoverTab[70959]++
															break
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:42
				// _ = "end of CoverTab[70959]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:43
				_go_fuzz_dep_.CoverTab[70960]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:43
				// _ = "end of CoverTab[70960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:43
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:43
			// _ = "end of CoverTab[70957]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:43
			_go_fuzz_dep_.CoverTab[70958]++
														bn := copy(w.buf, w.buf[i:])
														w.buf = w.buf[:bn]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:45
			// _ = "end of CoverTab[70958]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:46
			_go_fuzz_dep_.CoverTab[70961]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:46
			// _ = "end of CoverTab[70961]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:46
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:46
		// _ = "end of CoverTab[70952]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:47
	// _ = "end of CoverTab[70948]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:47
	_go_fuzz_dep_.CoverTab[70949]++
												return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:48
	// _ = "end of CoverTab[70949]"
}

// Close forces data that remains in the buffer to be written.
func (w *normWriter) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:52
	_go_fuzz_dep_.CoverTab[70962]++
												if len(w.buf) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:53
		_go_fuzz_dep_.CoverTab[70964]++
													_, err := w.w.Write(w.buf)
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:55
			_go_fuzz_dep_.CoverTab[70965]++
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:56
			// _ = "end of CoverTab[70965]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:57
			_go_fuzz_dep_.CoverTab[70966]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:57
			// _ = "end of CoverTab[70966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:57
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:57
		// _ = "end of CoverTab[70964]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:58
		_go_fuzz_dep_.CoverTab[70967]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:58
		// _ = "end of CoverTab[70967]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:58
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:58
	// _ = "end of CoverTab[70962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:58
	_go_fuzz_dep_.CoverTab[70963]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:59
	// _ = "end of CoverTab[70963]"
}

// Writer returns a new writer that implements Write(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:62
// by writing f(b) to w. The returned writer may use an
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:62
// internal buffer to maintain state across Write calls.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:62
// Calling its Close method writes any buffered data to w.
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:66
func (f Form) Writer(w io.Writer) io.WriteCloser {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:66
	_go_fuzz_dep_.CoverTab[70968]++
												wr := &normWriter{rb: reorderBuffer{}, w: w}
												wr.rb.init(f, nil)
												return wr
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:69
	// _ = "end of CoverTab[70968]"
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
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:83
	_go_fuzz_dep_.CoverTab[70969]++
												for {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:84
		_go_fuzz_dep_.CoverTab[70970]++
													if r.lastBoundary-r.bufStart > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:85
			_go_fuzz_dep_.CoverTab[70974]++
														n := copy(p, r.outbuf[r.bufStart:r.lastBoundary])
														r.bufStart += n
														if r.lastBoundary-r.bufStart > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:88
				_go_fuzz_dep_.CoverTab[70976]++
															return n, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:89
				// _ = "end of CoverTab[70976]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:90
				_go_fuzz_dep_.CoverTab[70977]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:90
				// _ = "end of CoverTab[70977]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:90
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:90
			// _ = "end of CoverTab[70974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:90
			_go_fuzz_dep_.CoverTab[70975]++
														return n, r.err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:91
			// _ = "end of CoverTab[70975]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:92
			_go_fuzz_dep_.CoverTab[70978]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:92
			// _ = "end of CoverTab[70978]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:92
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:92
		// _ = "end of CoverTab[70970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:92
		_go_fuzz_dep_.CoverTab[70971]++
													if r.err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:93
			_go_fuzz_dep_.CoverTab[70979]++
														return 0, r.err
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:94
			// _ = "end of CoverTab[70979]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:95
			_go_fuzz_dep_.CoverTab[70980]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:95
			// _ = "end of CoverTab[70980]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:95
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:95
		// _ = "end of CoverTab[70971]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:95
		_go_fuzz_dep_.CoverTab[70972]++
													outn := copy(r.outbuf, r.outbuf[r.lastBoundary:])
													r.outbuf = r.outbuf[0:outn]
													r.bufStart = 0

													n, err := r.r.Read(r.inbuf)
													r.rb.src = inputBytes(r.inbuf[0:n])
													r.rb.nsrc, r.err = n, err
													if n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:103
			_go_fuzz_dep_.CoverTab[70981]++
														r.outbuf = doAppend(&r.rb, r.outbuf, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:104
			// _ = "end of CoverTab[70981]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:105
			_go_fuzz_dep_.CoverTab[70982]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:105
			// _ = "end of CoverTab[70982]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:105
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:105
		// _ = "end of CoverTab[70972]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:105
		_go_fuzz_dep_.CoverTab[70973]++
													if err == io.EOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:106
			_go_fuzz_dep_.CoverTab[70983]++
														r.lastBoundary = len(r.outbuf)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:107
			// _ = "end of CoverTab[70983]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:108
			_go_fuzz_dep_.CoverTab[70984]++
														r.lastBoundary = lastBoundary(&r.rb.f, r.outbuf)
														if r.lastBoundary == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:110
				_go_fuzz_dep_.CoverTab[70985]++
															r.lastBoundary = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:111
				// _ = "end of CoverTab[70985]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:112
				_go_fuzz_dep_.CoverTab[70986]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:112
				// _ = "end of CoverTab[70986]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:112
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:112
			// _ = "end of CoverTab[70984]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:113
		// _ = "end of CoverTab[70973]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:114
	// _ = "end of CoverTab[70969]"
}

// Reader returns a new reader that implements Read
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:117
// by reading data from r and returning f(data).
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:119
func (f Form) Reader(r io.Reader) io.Reader {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:119
	_go_fuzz_dep_.CoverTab[70987]++
												const chunk = 4000
												buf := make([]byte, chunk)
												rr := &normReader{rb: reorderBuffer{}, r: r, inbuf: buf}
												rr.rb.init(f, buf)
												return rr
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:124
	// _ = "end of CoverTab[70987]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/readwriter.go:125
var _ = _go_fuzz_dep_.CoverTab
