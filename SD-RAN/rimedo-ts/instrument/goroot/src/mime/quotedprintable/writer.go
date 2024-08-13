// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/quotedprintable/writer.go:5
package quotedprintable

//line /usr/local/go/src/mime/quotedprintable/writer.go:5
import (
//line /usr/local/go/src/mime/quotedprintable/writer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/quotedprintable/writer.go:5
)
//line /usr/local/go/src/mime/quotedprintable/writer.go:5
import (
//line /usr/local/go/src/mime/quotedprintable/writer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/quotedprintable/writer.go:5
)

import "io"

const lineMaxLen = 76

// A Writer is a quoted-printable writer that implements io.WriteCloser.
type Writer struct {
	// Binary mode treats the writer's input as pure binary and processes end of
	// line bytes as binary data.
	Binary	bool

	w	io.Writer
	i	int
	line	[78]byte
	cr	bool
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
//line /usr/local/go/src/mime/quotedprintable/writer.go:24
	_go_fuzz_dep_.CoverTab[35975]++
								return &Writer{w: w}
//line /usr/local/go/src/mime/quotedprintable/writer.go:25
	// _ = "end of CoverTab[35975]"
}

// Write encodes p using quoted-printable encoding and writes it to the
//line /usr/local/go/src/mime/quotedprintable/writer.go:28
// underlying io.Writer. It limits line length to 76 characters. The encoded
//line /usr/local/go/src/mime/quotedprintable/writer.go:28
// bytes are not necessarily flushed until the Writer is closed.
//line /usr/local/go/src/mime/quotedprintable/writer.go:31
func (w *Writer) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/mime/quotedprintable/writer.go:31
	_go_fuzz_dep_.CoverTab[35976]++
								for i, b := range p {
//line /usr/local/go/src/mime/quotedprintable/writer.go:32
		_go_fuzz_dep_.CoverTab[35980]++
									switch {

		case b >= '!' && func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			_go_fuzz_dep_.CoverTab[35987]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			return b <= '~'
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			// _ = "end of CoverTab[35987]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
		}() && func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			_go_fuzz_dep_.CoverTab[35988]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			return b != '='
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			// _ = "end of CoverTab[35988]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
		}():
//line /usr/local/go/src/mime/quotedprintable/writer.go:35
			_go_fuzz_dep_.CoverTab[35984]++
										continue
//line /usr/local/go/src/mime/quotedprintable/writer.go:36
			// _ = "end of CoverTab[35984]"
		case isWhitespace(b) || func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
			_go_fuzz_dep_.CoverTab[35989]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
			return !w.Binary && func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
				_go_fuzz_dep_.CoverTab[35990]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
				return (b == '\n' || func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
					_go_fuzz_dep_.CoverTab[35991]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
					return b == '\r'
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
					// _ = "end of CoverTab[35991]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
				}())
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
				// _ = "end of CoverTab[35990]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
			}()
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
			// _ = "end of CoverTab[35989]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
		}():
//line /usr/local/go/src/mime/quotedprintable/writer.go:37
			_go_fuzz_dep_.CoverTab[35985]++
										continue
//line /usr/local/go/src/mime/quotedprintable/writer.go:38
			// _ = "end of CoverTab[35985]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:38
		default:
//line /usr/local/go/src/mime/quotedprintable/writer.go:38
			_go_fuzz_dep_.CoverTab[35986]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:38
			// _ = "end of CoverTab[35986]"
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:39
		// _ = "end of CoverTab[35980]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:39
		_go_fuzz_dep_.CoverTab[35981]++

									if i > n {
//line /usr/local/go/src/mime/quotedprintable/writer.go:41
			_go_fuzz_dep_.CoverTab[35992]++
										if err := w.write(p[n:i]); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:42
				_go_fuzz_dep_.CoverTab[35994]++
											return n, err
//line /usr/local/go/src/mime/quotedprintable/writer.go:43
				// _ = "end of CoverTab[35994]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:44
				_go_fuzz_dep_.CoverTab[35995]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:44
				// _ = "end of CoverTab[35995]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:44
			}
//line /usr/local/go/src/mime/quotedprintable/writer.go:44
			// _ = "end of CoverTab[35992]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:44
			_go_fuzz_dep_.CoverTab[35993]++
										n = i
//line /usr/local/go/src/mime/quotedprintable/writer.go:45
			// _ = "end of CoverTab[35993]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:46
			_go_fuzz_dep_.CoverTab[35996]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:46
			// _ = "end of CoverTab[35996]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:46
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:46
		// _ = "end of CoverTab[35981]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:46
		_go_fuzz_dep_.CoverTab[35982]++

									if err := w.encode(b); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:48
			_go_fuzz_dep_.CoverTab[35997]++
										return n, err
//line /usr/local/go/src/mime/quotedprintable/writer.go:49
			// _ = "end of CoverTab[35997]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:50
			_go_fuzz_dep_.CoverTab[35998]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:50
			// _ = "end of CoverTab[35998]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:50
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:50
		// _ = "end of CoverTab[35982]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:50
		_go_fuzz_dep_.CoverTab[35983]++
									n++
//line /usr/local/go/src/mime/quotedprintable/writer.go:51
		// _ = "end of CoverTab[35983]"
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:52
	// _ = "end of CoverTab[35976]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:52
	_go_fuzz_dep_.CoverTab[35977]++

								if n == len(p) {
//line /usr/local/go/src/mime/quotedprintable/writer.go:54
		_go_fuzz_dep_.CoverTab[35999]++
									return n, nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:55
		// _ = "end of CoverTab[35999]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:56
		_go_fuzz_dep_.CoverTab[36000]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:56
		// _ = "end of CoverTab[36000]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:56
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:56
	// _ = "end of CoverTab[35977]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:56
	_go_fuzz_dep_.CoverTab[35978]++

								if err := w.write(p[n:]); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:58
		_go_fuzz_dep_.CoverTab[36001]++
									return n, err
//line /usr/local/go/src/mime/quotedprintable/writer.go:59
		// _ = "end of CoverTab[36001]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:60
		_go_fuzz_dep_.CoverTab[36002]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:60
		// _ = "end of CoverTab[36002]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:60
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:60
	// _ = "end of CoverTab[35978]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:60
	_go_fuzz_dep_.CoverTab[35979]++

								return len(p), nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:62
	// _ = "end of CoverTab[35979]"
}

// Close closes the Writer, flushing any unwritten data to the underlying
//line /usr/local/go/src/mime/quotedprintable/writer.go:65
// io.Writer, but does not close the underlying io.Writer.
//line /usr/local/go/src/mime/quotedprintable/writer.go:67
func (w *Writer) Close() error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:67
	_go_fuzz_dep_.CoverTab[36003]++
								if err := w.checkLastByte(); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:68
		_go_fuzz_dep_.CoverTab[36005]++
									return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:69
		// _ = "end of CoverTab[36005]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:70
		_go_fuzz_dep_.CoverTab[36006]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:70
		// _ = "end of CoverTab[36006]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:70
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:70
	// _ = "end of CoverTab[36003]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:70
	_go_fuzz_dep_.CoverTab[36004]++

								return w.flush()
//line /usr/local/go/src/mime/quotedprintable/writer.go:72
	// _ = "end of CoverTab[36004]"
}

// write limits text encoded in quoted-printable to 76 characters per line.
func (w *Writer) write(p []byte) error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:76
	_go_fuzz_dep_.CoverTab[36007]++
								for _, b := range p {
//line /usr/local/go/src/mime/quotedprintable/writer.go:77
		_go_fuzz_dep_.CoverTab[36009]++
									if b == '\n' || func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:78
			_go_fuzz_dep_.CoverTab[36012]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:78
			return b == '\r'
//line /usr/local/go/src/mime/quotedprintable/writer.go:78
			// _ = "end of CoverTab[36012]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:78
		}() {
//line /usr/local/go/src/mime/quotedprintable/writer.go:78
			_go_fuzz_dep_.CoverTab[36013]++

										if w.cr && func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:80
				_go_fuzz_dep_.CoverTab[36018]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:80
				return b == '\n'
//line /usr/local/go/src/mime/quotedprintable/writer.go:80
				// _ = "end of CoverTab[36018]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:80
			}() {
//line /usr/local/go/src/mime/quotedprintable/writer.go:80
				_go_fuzz_dep_.CoverTab[36019]++
											w.cr = false
											continue
//line /usr/local/go/src/mime/quotedprintable/writer.go:82
				// _ = "end of CoverTab[36019]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:83
				_go_fuzz_dep_.CoverTab[36020]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:83
				// _ = "end of CoverTab[36020]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:83
			}
//line /usr/local/go/src/mime/quotedprintable/writer.go:83
			// _ = "end of CoverTab[36013]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:83
			_go_fuzz_dep_.CoverTab[36014]++

										if b == '\r' {
//line /usr/local/go/src/mime/quotedprintable/writer.go:85
				_go_fuzz_dep_.CoverTab[36021]++
											w.cr = true
//line /usr/local/go/src/mime/quotedprintable/writer.go:86
				// _ = "end of CoverTab[36021]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:87
				_go_fuzz_dep_.CoverTab[36022]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:87
				// _ = "end of CoverTab[36022]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:87
			}
//line /usr/local/go/src/mime/quotedprintable/writer.go:87
			// _ = "end of CoverTab[36014]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:87
			_go_fuzz_dep_.CoverTab[36015]++

										if err := w.checkLastByte(); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:89
				_go_fuzz_dep_.CoverTab[36023]++
											return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:90
				// _ = "end of CoverTab[36023]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:91
				_go_fuzz_dep_.CoverTab[36024]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:91
				// _ = "end of CoverTab[36024]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:91
			}
//line /usr/local/go/src/mime/quotedprintable/writer.go:91
			// _ = "end of CoverTab[36015]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:91
			_go_fuzz_dep_.CoverTab[36016]++
										if err := w.insertCRLF(); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:92
				_go_fuzz_dep_.CoverTab[36025]++
											return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:93
				// _ = "end of CoverTab[36025]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:94
				_go_fuzz_dep_.CoverTab[36026]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:94
				// _ = "end of CoverTab[36026]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:94
			}
//line /usr/local/go/src/mime/quotedprintable/writer.go:94
			// _ = "end of CoverTab[36016]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:94
			_go_fuzz_dep_.CoverTab[36017]++
										continue
//line /usr/local/go/src/mime/quotedprintable/writer.go:95
			// _ = "end of CoverTab[36017]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:96
			_go_fuzz_dep_.CoverTab[36027]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:96
			// _ = "end of CoverTab[36027]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:96
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:96
		// _ = "end of CoverTab[36009]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:96
		_go_fuzz_dep_.CoverTab[36010]++

									if w.i == lineMaxLen-1 {
//line /usr/local/go/src/mime/quotedprintable/writer.go:98
			_go_fuzz_dep_.CoverTab[36028]++
										if err := w.insertSoftLineBreak(); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:99
				_go_fuzz_dep_.CoverTab[36029]++
											return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:100
				// _ = "end of CoverTab[36029]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:101
				_go_fuzz_dep_.CoverTab[36030]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:101
				// _ = "end of CoverTab[36030]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:101
			}
//line /usr/local/go/src/mime/quotedprintable/writer.go:101
			// _ = "end of CoverTab[36028]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:102
			_go_fuzz_dep_.CoverTab[36031]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:102
			// _ = "end of CoverTab[36031]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:102
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:102
		// _ = "end of CoverTab[36010]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:102
		_go_fuzz_dep_.CoverTab[36011]++

									w.line[w.i] = b
									w.i++
									w.cr = false
//line /usr/local/go/src/mime/quotedprintable/writer.go:106
		// _ = "end of CoverTab[36011]"
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:107
	// _ = "end of CoverTab[36007]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:107
	_go_fuzz_dep_.CoverTab[36008]++

								return nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:109
	// _ = "end of CoverTab[36008]"
}

func (w *Writer) encode(b byte) error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:112
	_go_fuzz_dep_.CoverTab[36032]++
								if lineMaxLen-1-w.i < 3 {
//line /usr/local/go/src/mime/quotedprintable/writer.go:113
		_go_fuzz_dep_.CoverTab[36034]++
									if err := w.insertSoftLineBreak(); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:114
			_go_fuzz_dep_.CoverTab[36035]++
										return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:115
			// _ = "end of CoverTab[36035]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:116
			_go_fuzz_dep_.CoverTab[36036]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:116
			// _ = "end of CoverTab[36036]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:116
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:116
		// _ = "end of CoverTab[36034]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:117
		_go_fuzz_dep_.CoverTab[36037]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:117
		// _ = "end of CoverTab[36037]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:117
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:117
	// _ = "end of CoverTab[36032]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:117
	_go_fuzz_dep_.CoverTab[36033]++

								w.line[w.i] = '='
								w.line[w.i+1] = upperhex[b>>4]
								w.line[w.i+2] = upperhex[b&0x0f]
								w.i += 3

								return nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:124
	// _ = "end of CoverTab[36033]"
}

const upperhex = "0123456789ABCDEF"

// checkLastByte encodes the last buffered byte if it is a space or a tab.
func (w *Writer) checkLastByte() error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:130
	_go_fuzz_dep_.CoverTab[36038]++
								if w.i == 0 {
//line /usr/local/go/src/mime/quotedprintable/writer.go:131
		_go_fuzz_dep_.CoverTab[36041]++
									return nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:132
		// _ = "end of CoverTab[36041]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:133
		_go_fuzz_dep_.CoverTab[36042]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:133
		// _ = "end of CoverTab[36042]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:133
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:133
	// _ = "end of CoverTab[36038]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:133
	_go_fuzz_dep_.CoverTab[36039]++

								b := w.line[w.i-1]
								if isWhitespace(b) {
//line /usr/local/go/src/mime/quotedprintable/writer.go:136
		_go_fuzz_dep_.CoverTab[36043]++
									w.i--
									if err := w.encode(b); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:138
			_go_fuzz_dep_.CoverTab[36044]++
										return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:139
			// _ = "end of CoverTab[36044]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:140
			_go_fuzz_dep_.CoverTab[36045]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:140
			// _ = "end of CoverTab[36045]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:140
		}
//line /usr/local/go/src/mime/quotedprintable/writer.go:140
		// _ = "end of CoverTab[36043]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:141
		_go_fuzz_dep_.CoverTab[36046]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:141
		// _ = "end of CoverTab[36046]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:141
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:141
	// _ = "end of CoverTab[36039]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:141
	_go_fuzz_dep_.CoverTab[36040]++

								return nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:143
	// _ = "end of CoverTab[36040]"
}

func (w *Writer) insertSoftLineBreak() error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:146
	_go_fuzz_dep_.CoverTab[36047]++
								w.line[w.i] = '='
								w.i++

								return w.insertCRLF()
//line /usr/local/go/src/mime/quotedprintable/writer.go:150
	// _ = "end of CoverTab[36047]"
}

func (w *Writer) insertCRLF() error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:153
	_go_fuzz_dep_.CoverTab[36048]++
								w.line[w.i] = '\r'
								w.line[w.i+1] = '\n'
								w.i += 2

								return w.flush()
//line /usr/local/go/src/mime/quotedprintable/writer.go:158
	// _ = "end of CoverTab[36048]"
}

func (w *Writer) flush() error {
//line /usr/local/go/src/mime/quotedprintable/writer.go:161
	_go_fuzz_dep_.CoverTab[36049]++
								if _, err := w.w.Write(w.line[:w.i]); err != nil {
//line /usr/local/go/src/mime/quotedprintable/writer.go:162
		_go_fuzz_dep_.CoverTab[36051]++
									return err
//line /usr/local/go/src/mime/quotedprintable/writer.go:163
		// _ = "end of CoverTab[36051]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/writer.go:164
		_go_fuzz_dep_.CoverTab[36052]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:164
		// _ = "end of CoverTab[36052]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:164
	}
//line /usr/local/go/src/mime/quotedprintable/writer.go:164
	// _ = "end of CoverTab[36049]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:164
	_go_fuzz_dep_.CoverTab[36050]++

								w.i = 0
								return nil
//line /usr/local/go/src/mime/quotedprintable/writer.go:167
	// _ = "end of CoverTab[36050]"
}

func isWhitespace(b byte) bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:170
	_go_fuzz_dep_.CoverTab[36053]++
								return b == ' ' || func() bool {
//line /usr/local/go/src/mime/quotedprintable/writer.go:171
		_go_fuzz_dep_.CoverTab[36054]++
//line /usr/local/go/src/mime/quotedprintable/writer.go:171
		return b == '\t'
//line /usr/local/go/src/mime/quotedprintable/writer.go:171
		// _ = "end of CoverTab[36054]"
//line /usr/local/go/src/mime/quotedprintable/writer.go:171
	}()
//line /usr/local/go/src/mime/quotedprintable/writer.go:171
	// _ = "end of CoverTab[36053]"
}

//line /usr/local/go/src/mime/quotedprintable/writer.go:172
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/quotedprintable/writer.go:172
var _ = _go_fuzz_dep_.CoverTab
