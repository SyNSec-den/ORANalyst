// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/textproto/writer.go:5
package textproto

//line /usr/local/go/src/net/textproto/writer.go:5
import (
//line /usr/local/go/src/net/textproto/writer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/textproto/writer.go:5
)
//line /usr/local/go/src/net/textproto/writer.go:5
import (
//line /usr/local/go/src/net/textproto/writer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/textproto/writer.go:5
)

import (
	"bufio"
	"fmt"
	"io"
)

// A Writer implements convenience methods for writing
//line /usr/local/go/src/net/textproto/writer.go:13
// requests or responses to a text protocol network connection.
//line /usr/local/go/src/net/textproto/writer.go:15
type Writer struct {
	W	*bufio.Writer
	dot	*dotWriter
}

// NewWriter returns a new Writer writing to w.
func NewWriter(w *bufio.Writer) *Writer {
//line /usr/local/go/src/net/textproto/writer.go:21
	_go_fuzz_dep_.CoverTab[34874]++
							return &Writer{W: w}
//line /usr/local/go/src/net/textproto/writer.go:22
	// _ = "end of CoverTab[34874]"
}

var crnl = []byte{'\r', '\n'}
var dotcrnl = []byte{'.', '\r', '\n'}

// PrintfLine writes the formatted output followed by \r\n.
func (w *Writer) PrintfLine(format string, args ...any) error {
//line /usr/local/go/src/net/textproto/writer.go:29
	_go_fuzz_dep_.CoverTab[34875]++
							w.closeDot()
							fmt.Fprintf(w.W, format, args...)
							w.W.Write(crnl)
							return w.W.Flush()
//line /usr/local/go/src/net/textproto/writer.go:33
	// _ = "end of CoverTab[34875]"
}

// DotWriter returns a writer that can be used to write a dot-encoding to w.
//line /usr/local/go/src/net/textproto/writer.go:36
// It takes care of inserting leading dots when necessary,
//line /usr/local/go/src/net/textproto/writer.go:36
// translating line-ending \n into \r\n, and adding the final .\r\n line
//line /usr/local/go/src/net/textproto/writer.go:36
// when the DotWriter is closed. The caller should close the
//line /usr/local/go/src/net/textproto/writer.go:36
// DotWriter before the next call to a method on w.
//line /usr/local/go/src/net/textproto/writer.go:36
//
//line /usr/local/go/src/net/textproto/writer.go:36
// See the documentation for Reader's DotReader method for details about dot-encoding.
//line /usr/local/go/src/net/textproto/writer.go:43
func (w *Writer) DotWriter() io.WriteCloser {
//line /usr/local/go/src/net/textproto/writer.go:43
	_go_fuzz_dep_.CoverTab[34876]++
							w.closeDot()
							w.dot = &dotWriter{w: w}
							return w.dot
//line /usr/local/go/src/net/textproto/writer.go:46
	// _ = "end of CoverTab[34876]"
}

func (w *Writer) closeDot() {
//line /usr/local/go/src/net/textproto/writer.go:49
	_go_fuzz_dep_.CoverTab[34877]++
							if w.dot != nil {
//line /usr/local/go/src/net/textproto/writer.go:50
		_go_fuzz_dep_.CoverTab[34878]++
								w.dot.Close()
//line /usr/local/go/src/net/textproto/writer.go:51
		// _ = "end of CoverTab[34878]"
	} else {
//line /usr/local/go/src/net/textproto/writer.go:52
		_go_fuzz_dep_.CoverTab[34879]++
//line /usr/local/go/src/net/textproto/writer.go:52
		// _ = "end of CoverTab[34879]"
//line /usr/local/go/src/net/textproto/writer.go:52
	}
//line /usr/local/go/src/net/textproto/writer.go:52
	// _ = "end of CoverTab[34877]"
}

type dotWriter struct {
	w	*Writer
	state	int
}

const (
	wstateBegin	= iota	// initial state; must be zero
	wstateBeginLine		// beginning of line
	wstateCR		// wrote \r (possibly at end of line)
	wstateData		// writing data in middle of line
)

func (d *dotWriter) Write(b []byte) (n int, err error) {
//line /usr/local/go/src/net/textproto/writer.go:67
	_go_fuzz_dep_.CoverTab[34880]++
							bw := d.w.W
							for n < len(b) {
//line /usr/local/go/src/net/textproto/writer.go:69
		_go_fuzz_dep_.CoverTab[34882]++
								c := b[n]
								switch d.state {
		case wstateBegin, wstateBeginLine:
//line /usr/local/go/src/net/textproto/writer.go:72
			_go_fuzz_dep_.CoverTab[34885]++
									d.state = wstateData
									if c == '.' {
//line /usr/local/go/src/net/textproto/writer.go:74
				_go_fuzz_dep_.CoverTab[34891]++

										bw.WriteByte('.')
//line /usr/local/go/src/net/textproto/writer.go:76
				// _ = "end of CoverTab[34891]"
			} else {
//line /usr/local/go/src/net/textproto/writer.go:77
				_go_fuzz_dep_.CoverTab[34892]++
//line /usr/local/go/src/net/textproto/writer.go:77
				// _ = "end of CoverTab[34892]"
//line /usr/local/go/src/net/textproto/writer.go:77
			}
//line /usr/local/go/src/net/textproto/writer.go:77
			// _ = "end of CoverTab[34885]"
//line /usr/local/go/src/net/textproto/writer.go:77
			_go_fuzz_dep_.CoverTab[34886]++
									fallthrough
//line /usr/local/go/src/net/textproto/writer.go:78
			// _ = "end of CoverTab[34886]"

		case wstateData:
//line /usr/local/go/src/net/textproto/writer.go:80
			_go_fuzz_dep_.CoverTab[34887]++
									if c == '\r' {
//line /usr/local/go/src/net/textproto/writer.go:81
				_go_fuzz_dep_.CoverTab[34893]++
										d.state = wstateCR
//line /usr/local/go/src/net/textproto/writer.go:82
				// _ = "end of CoverTab[34893]"
			} else {
//line /usr/local/go/src/net/textproto/writer.go:83
				_go_fuzz_dep_.CoverTab[34894]++
//line /usr/local/go/src/net/textproto/writer.go:83
				// _ = "end of CoverTab[34894]"
//line /usr/local/go/src/net/textproto/writer.go:83
			}
//line /usr/local/go/src/net/textproto/writer.go:83
			// _ = "end of CoverTab[34887]"
//line /usr/local/go/src/net/textproto/writer.go:83
			_go_fuzz_dep_.CoverTab[34888]++
									if c == '\n' {
//line /usr/local/go/src/net/textproto/writer.go:84
				_go_fuzz_dep_.CoverTab[34895]++
										bw.WriteByte('\r')
										d.state = wstateBeginLine
//line /usr/local/go/src/net/textproto/writer.go:86
				// _ = "end of CoverTab[34895]"
			} else {
//line /usr/local/go/src/net/textproto/writer.go:87
				_go_fuzz_dep_.CoverTab[34896]++
//line /usr/local/go/src/net/textproto/writer.go:87
				// _ = "end of CoverTab[34896]"
//line /usr/local/go/src/net/textproto/writer.go:87
			}
//line /usr/local/go/src/net/textproto/writer.go:87
			// _ = "end of CoverTab[34888]"

		case wstateCR:
//line /usr/local/go/src/net/textproto/writer.go:89
			_go_fuzz_dep_.CoverTab[34889]++
									d.state = wstateData
									if c == '\n' {
//line /usr/local/go/src/net/textproto/writer.go:91
				_go_fuzz_dep_.CoverTab[34897]++
										d.state = wstateBeginLine
//line /usr/local/go/src/net/textproto/writer.go:92
				// _ = "end of CoverTab[34897]"
			} else {
//line /usr/local/go/src/net/textproto/writer.go:93
				_go_fuzz_dep_.CoverTab[34898]++
//line /usr/local/go/src/net/textproto/writer.go:93
				// _ = "end of CoverTab[34898]"
//line /usr/local/go/src/net/textproto/writer.go:93
			}
//line /usr/local/go/src/net/textproto/writer.go:93
			// _ = "end of CoverTab[34889]"
//line /usr/local/go/src/net/textproto/writer.go:93
		default:
//line /usr/local/go/src/net/textproto/writer.go:93
			_go_fuzz_dep_.CoverTab[34890]++
//line /usr/local/go/src/net/textproto/writer.go:93
			// _ = "end of CoverTab[34890]"
		}
//line /usr/local/go/src/net/textproto/writer.go:94
		// _ = "end of CoverTab[34882]"
//line /usr/local/go/src/net/textproto/writer.go:94
		_go_fuzz_dep_.CoverTab[34883]++
								if err = bw.WriteByte(c); err != nil {
//line /usr/local/go/src/net/textproto/writer.go:95
			_go_fuzz_dep_.CoverTab[34899]++
									break
//line /usr/local/go/src/net/textproto/writer.go:96
			// _ = "end of CoverTab[34899]"
		} else {
//line /usr/local/go/src/net/textproto/writer.go:97
			_go_fuzz_dep_.CoverTab[34900]++
//line /usr/local/go/src/net/textproto/writer.go:97
			// _ = "end of CoverTab[34900]"
//line /usr/local/go/src/net/textproto/writer.go:97
		}
//line /usr/local/go/src/net/textproto/writer.go:97
		// _ = "end of CoverTab[34883]"
//line /usr/local/go/src/net/textproto/writer.go:97
		_go_fuzz_dep_.CoverTab[34884]++
								n++
//line /usr/local/go/src/net/textproto/writer.go:98
		// _ = "end of CoverTab[34884]"
	}
//line /usr/local/go/src/net/textproto/writer.go:99
	// _ = "end of CoverTab[34880]"
//line /usr/local/go/src/net/textproto/writer.go:99
	_go_fuzz_dep_.CoverTab[34881]++
							return
//line /usr/local/go/src/net/textproto/writer.go:100
	// _ = "end of CoverTab[34881]"
}

func (d *dotWriter) Close() error {
//line /usr/local/go/src/net/textproto/writer.go:103
	_go_fuzz_dep_.CoverTab[34901]++
							if d.w.dot == d {
//line /usr/local/go/src/net/textproto/writer.go:104
		_go_fuzz_dep_.CoverTab[34904]++
								d.w.dot = nil
//line /usr/local/go/src/net/textproto/writer.go:105
		// _ = "end of CoverTab[34904]"
	} else {
//line /usr/local/go/src/net/textproto/writer.go:106
		_go_fuzz_dep_.CoverTab[34905]++
//line /usr/local/go/src/net/textproto/writer.go:106
		// _ = "end of CoverTab[34905]"
//line /usr/local/go/src/net/textproto/writer.go:106
	}
//line /usr/local/go/src/net/textproto/writer.go:106
	// _ = "end of CoverTab[34901]"
//line /usr/local/go/src/net/textproto/writer.go:106
	_go_fuzz_dep_.CoverTab[34902]++
							bw := d.w.W
							switch d.state {
	default:
//line /usr/local/go/src/net/textproto/writer.go:109
		_go_fuzz_dep_.CoverTab[34906]++
								bw.WriteByte('\r')
								fallthrough
//line /usr/local/go/src/net/textproto/writer.go:111
		// _ = "end of CoverTab[34906]"
	case wstateCR:
//line /usr/local/go/src/net/textproto/writer.go:112
		_go_fuzz_dep_.CoverTab[34907]++
								bw.WriteByte('\n')
								fallthrough
//line /usr/local/go/src/net/textproto/writer.go:114
		// _ = "end of CoverTab[34907]"
	case wstateBeginLine:
//line /usr/local/go/src/net/textproto/writer.go:115
		_go_fuzz_dep_.CoverTab[34908]++
								bw.Write(dotcrnl)
//line /usr/local/go/src/net/textproto/writer.go:116
		// _ = "end of CoverTab[34908]"
	}
//line /usr/local/go/src/net/textproto/writer.go:117
	// _ = "end of CoverTab[34902]"
//line /usr/local/go/src/net/textproto/writer.go:117
	_go_fuzz_dep_.CoverTab[34903]++
							return bw.Flush()
//line /usr/local/go/src/net/textproto/writer.go:118
	// _ = "end of CoverTab[34903]"
}

//line /usr/local/go/src/net/textproto/writer.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/textproto/writer.go:119
var _ = _go_fuzz_dep_.CoverTab
