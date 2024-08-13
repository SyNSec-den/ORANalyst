// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/gzip/gzip.go:5
package gzip

//line /usr/local/go/src/compress/gzip/gzip.go:5
import (
//line /usr/local/go/src/compress/gzip/gzip.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/gzip/gzip.go:5
)
//line /usr/local/go/src/compress/gzip/gzip.go:5
import (
//line /usr/local/go/src/compress/gzip/gzip.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/gzip/gzip.go:5
)

import (
	"compress/flate"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"time"
)

// These constants are copied from the flate package, so that code that imports
//line /usr/local/go/src/compress/gzip/gzip.go:16
// "compress/gzip" does not also have to import "compress/flate".
//line /usr/local/go/src/compress/gzip/gzip.go:18
const (
	NoCompression		= flate.NoCompression
	BestSpeed		= flate.BestSpeed
	BestCompression		= flate.BestCompression
	DefaultCompression	= flate.DefaultCompression
	HuffmanOnly		= flate.HuffmanOnly
)

// A Writer is an io.WriteCloser.
//line /usr/local/go/src/compress/gzip/gzip.go:26
// Writes to a Writer are compressed and written to w.
//line /usr/local/go/src/compress/gzip/gzip.go:28
type Writer struct {
	Header		// written at first call to Write, Flush, or Close
	w		io.Writer
	level		int
	wroteHeader	bool
	compressor	*flate.Writer
	digest		uint32	// CRC-32, IEEE polynomial (section 8)
	size		uint32	// Uncompressed size (section 2.3.1)
	closed		bool
	buf		[10]byte
	err		error
}

// NewWriter returns a new Writer.
//line /usr/local/go/src/compress/gzip/gzip.go:41
// Writes to the returned writer are compressed and written to w.
//line /usr/local/go/src/compress/gzip/gzip.go:41
//
//line /usr/local/go/src/compress/gzip/gzip.go:41
// It is the caller's responsibility to call Close on the Writer when done.
//line /usr/local/go/src/compress/gzip/gzip.go:41
// Writes may be buffered and not flushed until Close.
//line /usr/local/go/src/compress/gzip/gzip.go:41
//
//line /usr/local/go/src/compress/gzip/gzip.go:41
// Callers that wish to set the fields in Writer.Header must do so before
//line /usr/local/go/src/compress/gzip/gzip.go:41
// the first call to Write, Flush, or Close.
//line /usr/local/go/src/compress/gzip/gzip.go:49
func NewWriter(w io.Writer) *Writer {
//line /usr/local/go/src/compress/gzip/gzip.go:49
	_go_fuzz_dep_.CoverTab[26752]++
							z, _ := NewWriterLevel(w, DefaultCompression)
							return z
//line /usr/local/go/src/compress/gzip/gzip.go:51
	// _ = "end of CoverTab[26752]"
}

// NewWriterLevel is like NewWriter but specifies the compression level instead
//line /usr/local/go/src/compress/gzip/gzip.go:54
// of assuming DefaultCompression.
//line /usr/local/go/src/compress/gzip/gzip.go:54
//
//line /usr/local/go/src/compress/gzip/gzip.go:54
// The compression level can be DefaultCompression, NoCompression, HuffmanOnly
//line /usr/local/go/src/compress/gzip/gzip.go:54
// or any integer value between BestSpeed and BestCompression inclusive.
//line /usr/local/go/src/compress/gzip/gzip.go:54
// The error returned will be nil if the level is valid.
//line /usr/local/go/src/compress/gzip/gzip.go:60
func NewWriterLevel(w io.Writer, level int) (*Writer, error) {
//line /usr/local/go/src/compress/gzip/gzip.go:60
	_go_fuzz_dep_.CoverTab[26753]++
							if level < HuffmanOnly || func() bool {
//line /usr/local/go/src/compress/gzip/gzip.go:61
		_go_fuzz_dep_.CoverTab[26755]++
//line /usr/local/go/src/compress/gzip/gzip.go:61
		return level > BestCompression
//line /usr/local/go/src/compress/gzip/gzip.go:61
		// _ = "end of CoverTab[26755]"
//line /usr/local/go/src/compress/gzip/gzip.go:61
	}() {
//line /usr/local/go/src/compress/gzip/gzip.go:61
		_go_fuzz_dep_.CoverTab[26756]++
								return nil, fmt.Errorf("gzip: invalid compression level: %d", level)
//line /usr/local/go/src/compress/gzip/gzip.go:62
		// _ = "end of CoverTab[26756]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:63
		_go_fuzz_dep_.CoverTab[26757]++
//line /usr/local/go/src/compress/gzip/gzip.go:63
		// _ = "end of CoverTab[26757]"
//line /usr/local/go/src/compress/gzip/gzip.go:63
	}
//line /usr/local/go/src/compress/gzip/gzip.go:63
	// _ = "end of CoverTab[26753]"
//line /usr/local/go/src/compress/gzip/gzip.go:63
	_go_fuzz_dep_.CoverTab[26754]++
							z := new(Writer)
							z.init(w, level)
							return z, nil
//line /usr/local/go/src/compress/gzip/gzip.go:66
	// _ = "end of CoverTab[26754]"
}

func (z *Writer) init(w io.Writer, level int) {
	compressor := z.compressor
	if compressor != nil {
		compressor.Reset(w)
	}
	*z = Writer{
		Header: Header{
			OS: 255,
		},
		w:		w,
		level:		level,
		compressor:	compressor,
	}
}

// Reset discards the Writer z's state and makes it equivalent to the
//line /usr/local/go/src/compress/gzip/gzip.go:84
// result of its original state from NewWriter or NewWriterLevel, but
//line /usr/local/go/src/compress/gzip/gzip.go:84
// writing to w instead. This permits reusing a Writer rather than
//line /usr/local/go/src/compress/gzip/gzip.go:84
// allocating a new one.
//line /usr/local/go/src/compress/gzip/gzip.go:88
func (z *Writer) Reset(w io.Writer) {
//line /usr/local/go/src/compress/gzip/gzip.go:88
	_go_fuzz_dep_.CoverTab[26758]++
							z.init(w, z.level)
//line /usr/local/go/src/compress/gzip/gzip.go:89
	// _ = "end of CoverTab[26758]"
}

// writeBytes writes a length-prefixed byte slice to z.w.
func (z *Writer) writeBytes(b []byte) error {
//line /usr/local/go/src/compress/gzip/gzip.go:93
	_go_fuzz_dep_.CoverTab[26759]++
							if len(b) > 0xffff {
//line /usr/local/go/src/compress/gzip/gzip.go:94
		_go_fuzz_dep_.CoverTab[26762]++
								return errors.New("gzip.Write: Extra data is too large")
//line /usr/local/go/src/compress/gzip/gzip.go:95
		// _ = "end of CoverTab[26762]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:96
		_go_fuzz_dep_.CoverTab[26763]++
//line /usr/local/go/src/compress/gzip/gzip.go:96
		// _ = "end of CoverTab[26763]"
//line /usr/local/go/src/compress/gzip/gzip.go:96
	}
//line /usr/local/go/src/compress/gzip/gzip.go:96
	// _ = "end of CoverTab[26759]"
//line /usr/local/go/src/compress/gzip/gzip.go:96
	_go_fuzz_dep_.CoverTab[26760]++
							le.PutUint16(z.buf[:2], uint16(len(b)))
							_, err := z.w.Write(z.buf[:2])
							if err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:99
		_go_fuzz_dep_.CoverTab[26764]++
								return err
//line /usr/local/go/src/compress/gzip/gzip.go:100
		// _ = "end of CoverTab[26764]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:101
		_go_fuzz_dep_.CoverTab[26765]++
//line /usr/local/go/src/compress/gzip/gzip.go:101
		// _ = "end of CoverTab[26765]"
//line /usr/local/go/src/compress/gzip/gzip.go:101
	}
//line /usr/local/go/src/compress/gzip/gzip.go:101
	// _ = "end of CoverTab[26760]"
//line /usr/local/go/src/compress/gzip/gzip.go:101
	_go_fuzz_dep_.CoverTab[26761]++
							_, err = z.w.Write(b)
							return err
//line /usr/local/go/src/compress/gzip/gzip.go:103
	// _ = "end of CoverTab[26761]"
}

// writeString writes a UTF-8 string s in GZIP's format to z.w.
//line /usr/local/go/src/compress/gzip/gzip.go:106
// GZIP (RFC 1952) specifies that strings are NUL-terminated ISO 8859-1 (Latin-1).
//line /usr/local/go/src/compress/gzip/gzip.go:108
func (z *Writer) writeString(s string) (err error) {
//line /usr/local/go/src/compress/gzip/gzip.go:108
	_go_fuzz_dep_.CoverTab[26766]++

							needconv := false
							for _, v := range s {
//line /usr/local/go/src/compress/gzip/gzip.go:111
		_go_fuzz_dep_.CoverTab[26770]++
								if v == 0 || func() bool {
//line /usr/local/go/src/compress/gzip/gzip.go:112
			_go_fuzz_dep_.CoverTab[26772]++
//line /usr/local/go/src/compress/gzip/gzip.go:112
			return v > 0xff
//line /usr/local/go/src/compress/gzip/gzip.go:112
			// _ = "end of CoverTab[26772]"
//line /usr/local/go/src/compress/gzip/gzip.go:112
		}() {
//line /usr/local/go/src/compress/gzip/gzip.go:112
			_go_fuzz_dep_.CoverTab[26773]++
									return errors.New("gzip.Write: non-Latin-1 header string")
//line /usr/local/go/src/compress/gzip/gzip.go:113
			// _ = "end of CoverTab[26773]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:114
			_go_fuzz_dep_.CoverTab[26774]++
//line /usr/local/go/src/compress/gzip/gzip.go:114
			// _ = "end of CoverTab[26774]"
//line /usr/local/go/src/compress/gzip/gzip.go:114
		}
//line /usr/local/go/src/compress/gzip/gzip.go:114
		// _ = "end of CoverTab[26770]"
//line /usr/local/go/src/compress/gzip/gzip.go:114
		_go_fuzz_dep_.CoverTab[26771]++
								if v > 0x7f {
//line /usr/local/go/src/compress/gzip/gzip.go:115
			_go_fuzz_dep_.CoverTab[26775]++
									needconv = true
//line /usr/local/go/src/compress/gzip/gzip.go:116
			// _ = "end of CoverTab[26775]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:117
			_go_fuzz_dep_.CoverTab[26776]++
//line /usr/local/go/src/compress/gzip/gzip.go:117
			// _ = "end of CoverTab[26776]"
//line /usr/local/go/src/compress/gzip/gzip.go:117
		}
//line /usr/local/go/src/compress/gzip/gzip.go:117
		// _ = "end of CoverTab[26771]"
	}
//line /usr/local/go/src/compress/gzip/gzip.go:118
	// _ = "end of CoverTab[26766]"
//line /usr/local/go/src/compress/gzip/gzip.go:118
	_go_fuzz_dep_.CoverTab[26767]++
							if needconv {
//line /usr/local/go/src/compress/gzip/gzip.go:119
		_go_fuzz_dep_.CoverTab[26777]++
								b := make([]byte, 0, len(s))
								for _, v := range s {
//line /usr/local/go/src/compress/gzip/gzip.go:121
			_go_fuzz_dep_.CoverTab[26779]++
									b = append(b, byte(v))
//line /usr/local/go/src/compress/gzip/gzip.go:122
			// _ = "end of CoverTab[26779]"
		}
//line /usr/local/go/src/compress/gzip/gzip.go:123
		// _ = "end of CoverTab[26777]"
//line /usr/local/go/src/compress/gzip/gzip.go:123
		_go_fuzz_dep_.CoverTab[26778]++
								_, err = z.w.Write(b)
//line /usr/local/go/src/compress/gzip/gzip.go:124
		// _ = "end of CoverTab[26778]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:125
		_go_fuzz_dep_.CoverTab[26780]++
								_, err = io.WriteString(z.w, s)
//line /usr/local/go/src/compress/gzip/gzip.go:126
		// _ = "end of CoverTab[26780]"
	}
//line /usr/local/go/src/compress/gzip/gzip.go:127
	// _ = "end of CoverTab[26767]"
//line /usr/local/go/src/compress/gzip/gzip.go:127
	_go_fuzz_dep_.CoverTab[26768]++
							if err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:128
		_go_fuzz_dep_.CoverTab[26781]++
								return err
//line /usr/local/go/src/compress/gzip/gzip.go:129
		// _ = "end of CoverTab[26781]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:130
		_go_fuzz_dep_.CoverTab[26782]++
//line /usr/local/go/src/compress/gzip/gzip.go:130
		// _ = "end of CoverTab[26782]"
//line /usr/local/go/src/compress/gzip/gzip.go:130
	}
//line /usr/local/go/src/compress/gzip/gzip.go:130
	// _ = "end of CoverTab[26768]"
//line /usr/local/go/src/compress/gzip/gzip.go:130
	_go_fuzz_dep_.CoverTab[26769]++

							z.buf[0] = 0
							_, err = z.w.Write(z.buf[:1])
							return err
//line /usr/local/go/src/compress/gzip/gzip.go:134
	// _ = "end of CoverTab[26769]"
}

// Write writes a compressed form of p to the underlying io.Writer. The
//line /usr/local/go/src/compress/gzip/gzip.go:137
// compressed bytes are not necessarily flushed until the Writer is closed.
//line /usr/local/go/src/compress/gzip/gzip.go:139
func (z *Writer) Write(p []byte) (int, error) {
//line /usr/local/go/src/compress/gzip/gzip.go:139
	_go_fuzz_dep_.CoverTab[26783]++
							if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:140
		_go_fuzz_dep_.CoverTab[26786]++
								return 0, z.err
//line /usr/local/go/src/compress/gzip/gzip.go:141
		// _ = "end of CoverTab[26786]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:142
		_go_fuzz_dep_.CoverTab[26787]++
//line /usr/local/go/src/compress/gzip/gzip.go:142
		// _ = "end of CoverTab[26787]"
//line /usr/local/go/src/compress/gzip/gzip.go:142
	}
//line /usr/local/go/src/compress/gzip/gzip.go:142
	// _ = "end of CoverTab[26783]"
//line /usr/local/go/src/compress/gzip/gzip.go:142
	_go_fuzz_dep_.CoverTab[26784]++
							var n int

							if !z.wroteHeader {
//line /usr/local/go/src/compress/gzip/gzip.go:145
		_go_fuzz_dep_.CoverTab[26788]++
								z.wroteHeader = true
								z.buf = [10]byte{0: gzipID1, 1: gzipID2, 2: gzipDeflate}
								if z.Extra != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:148
			_go_fuzz_dep_.CoverTab[26798]++
									z.buf[3] |= 0x04
//line /usr/local/go/src/compress/gzip/gzip.go:149
			// _ = "end of CoverTab[26798]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:150
			_go_fuzz_dep_.CoverTab[26799]++
//line /usr/local/go/src/compress/gzip/gzip.go:150
			// _ = "end of CoverTab[26799]"
//line /usr/local/go/src/compress/gzip/gzip.go:150
		}
//line /usr/local/go/src/compress/gzip/gzip.go:150
		// _ = "end of CoverTab[26788]"
//line /usr/local/go/src/compress/gzip/gzip.go:150
		_go_fuzz_dep_.CoverTab[26789]++
								if z.Name != "" {
//line /usr/local/go/src/compress/gzip/gzip.go:151
			_go_fuzz_dep_.CoverTab[26800]++
									z.buf[3] |= 0x08
//line /usr/local/go/src/compress/gzip/gzip.go:152
			// _ = "end of CoverTab[26800]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:153
			_go_fuzz_dep_.CoverTab[26801]++
//line /usr/local/go/src/compress/gzip/gzip.go:153
			// _ = "end of CoverTab[26801]"
//line /usr/local/go/src/compress/gzip/gzip.go:153
		}
//line /usr/local/go/src/compress/gzip/gzip.go:153
		// _ = "end of CoverTab[26789]"
//line /usr/local/go/src/compress/gzip/gzip.go:153
		_go_fuzz_dep_.CoverTab[26790]++
								if z.Comment != "" {
//line /usr/local/go/src/compress/gzip/gzip.go:154
			_go_fuzz_dep_.CoverTab[26802]++
									z.buf[3] |= 0x10
//line /usr/local/go/src/compress/gzip/gzip.go:155
			// _ = "end of CoverTab[26802]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:156
			_go_fuzz_dep_.CoverTab[26803]++
//line /usr/local/go/src/compress/gzip/gzip.go:156
			// _ = "end of CoverTab[26803]"
//line /usr/local/go/src/compress/gzip/gzip.go:156
		}
//line /usr/local/go/src/compress/gzip/gzip.go:156
		// _ = "end of CoverTab[26790]"
//line /usr/local/go/src/compress/gzip/gzip.go:156
		_go_fuzz_dep_.CoverTab[26791]++
								if z.ModTime.After(time.Unix(0, 0)) {
//line /usr/local/go/src/compress/gzip/gzip.go:157
			_go_fuzz_dep_.CoverTab[26804]++

//line /usr/local/go/src/compress/gzip/gzip.go:160
			le.PutUint32(z.buf[4:8], uint32(z.ModTime.Unix()))
//line /usr/local/go/src/compress/gzip/gzip.go:160
			// _ = "end of CoverTab[26804]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:161
			_go_fuzz_dep_.CoverTab[26805]++
//line /usr/local/go/src/compress/gzip/gzip.go:161
			// _ = "end of CoverTab[26805]"
//line /usr/local/go/src/compress/gzip/gzip.go:161
		}
//line /usr/local/go/src/compress/gzip/gzip.go:161
		// _ = "end of CoverTab[26791]"
//line /usr/local/go/src/compress/gzip/gzip.go:161
		_go_fuzz_dep_.CoverTab[26792]++
								if z.level == BestCompression {
//line /usr/local/go/src/compress/gzip/gzip.go:162
			_go_fuzz_dep_.CoverTab[26806]++
									z.buf[8] = 2
//line /usr/local/go/src/compress/gzip/gzip.go:163
			// _ = "end of CoverTab[26806]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:164
			_go_fuzz_dep_.CoverTab[26807]++
//line /usr/local/go/src/compress/gzip/gzip.go:164
			if z.level == BestSpeed {
//line /usr/local/go/src/compress/gzip/gzip.go:164
				_go_fuzz_dep_.CoverTab[26808]++
										z.buf[8] = 4
//line /usr/local/go/src/compress/gzip/gzip.go:165
				// _ = "end of CoverTab[26808]"
			} else {
//line /usr/local/go/src/compress/gzip/gzip.go:166
				_go_fuzz_dep_.CoverTab[26809]++
//line /usr/local/go/src/compress/gzip/gzip.go:166
				// _ = "end of CoverTab[26809]"
//line /usr/local/go/src/compress/gzip/gzip.go:166
			}
//line /usr/local/go/src/compress/gzip/gzip.go:166
			// _ = "end of CoverTab[26807]"
//line /usr/local/go/src/compress/gzip/gzip.go:166
		}
//line /usr/local/go/src/compress/gzip/gzip.go:166
		// _ = "end of CoverTab[26792]"
//line /usr/local/go/src/compress/gzip/gzip.go:166
		_go_fuzz_dep_.CoverTab[26793]++
								z.buf[9] = z.OS
								_, z.err = z.w.Write(z.buf[:10])
								if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:169
			_go_fuzz_dep_.CoverTab[26810]++
									return 0, z.err
//line /usr/local/go/src/compress/gzip/gzip.go:170
			// _ = "end of CoverTab[26810]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:171
			_go_fuzz_dep_.CoverTab[26811]++
//line /usr/local/go/src/compress/gzip/gzip.go:171
			// _ = "end of CoverTab[26811]"
//line /usr/local/go/src/compress/gzip/gzip.go:171
		}
//line /usr/local/go/src/compress/gzip/gzip.go:171
		// _ = "end of CoverTab[26793]"
//line /usr/local/go/src/compress/gzip/gzip.go:171
		_go_fuzz_dep_.CoverTab[26794]++
								if z.Extra != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:172
			_go_fuzz_dep_.CoverTab[26812]++
									z.err = z.writeBytes(z.Extra)
									if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:174
				_go_fuzz_dep_.CoverTab[26813]++
										return 0, z.err
//line /usr/local/go/src/compress/gzip/gzip.go:175
				// _ = "end of CoverTab[26813]"
			} else {
//line /usr/local/go/src/compress/gzip/gzip.go:176
				_go_fuzz_dep_.CoverTab[26814]++
//line /usr/local/go/src/compress/gzip/gzip.go:176
				// _ = "end of CoverTab[26814]"
//line /usr/local/go/src/compress/gzip/gzip.go:176
			}
//line /usr/local/go/src/compress/gzip/gzip.go:176
			// _ = "end of CoverTab[26812]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:177
			_go_fuzz_dep_.CoverTab[26815]++
//line /usr/local/go/src/compress/gzip/gzip.go:177
			// _ = "end of CoverTab[26815]"
//line /usr/local/go/src/compress/gzip/gzip.go:177
		}
//line /usr/local/go/src/compress/gzip/gzip.go:177
		// _ = "end of CoverTab[26794]"
//line /usr/local/go/src/compress/gzip/gzip.go:177
		_go_fuzz_dep_.CoverTab[26795]++
								if z.Name != "" {
//line /usr/local/go/src/compress/gzip/gzip.go:178
			_go_fuzz_dep_.CoverTab[26816]++
									z.err = z.writeString(z.Name)
									if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:180
				_go_fuzz_dep_.CoverTab[26817]++
										return 0, z.err
//line /usr/local/go/src/compress/gzip/gzip.go:181
				// _ = "end of CoverTab[26817]"
			} else {
//line /usr/local/go/src/compress/gzip/gzip.go:182
				_go_fuzz_dep_.CoverTab[26818]++
//line /usr/local/go/src/compress/gzip/gzip.go:182
				// _ = "end of CoverTab[26818]"
//line /usr/local/go/src/compress/gzip/gzip.go:182
			}
//line /usr/local/go/src/compress/gzip/gzip.go:182
			// _ = "end of CoverTab[26816]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:183
			_go_fuzz_dep_.CoverTab[26819]++
//line /usr/local/go/src/compress/gzip/gzip.go:183
			// _ = "end of CoverTab[26819]"
//line /usr/local/go/src/compress/gzip/gzip.go:183
		}
//line /usr/local/go/src/compress/gzip/gzip.go:183
		// _ = "end of CoverTab[26795]"
//line /usr/local/go/src/compress/gzip/gzip.go:183
		_go_fuzz_dep_.CoverTab[26796]++
								if z.Comment != "" {
//line /usr/local/go/src/compress/gzip/gzip.go:184
			_go_fuzz_dep_.CoverTab[26820]++
									z.err = z.writeString(z.Comment)
									if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:186
				_go_fuzz_dep_.CoverTab[26821]++
										return 0, z.err
//line /usr/local/go/src/compress/gzip/gzip.go:187
				// _ = "end of CoverTab[26821]"
			} else {
//line /usr/local/go/src/compress/gzip/gzip.go:188
				_go_fuzz_dep_.CoverTab[26822]++
//line /usr/local/go/src/compress/gzip/gzip.go:188
				// _ = "end of CoverTab[26822]"
//line /usr/local/go/src/compress/gzip/gzip.go:188
			}
//line /usr/local/go/src/compress/gzip/gzip.go:188
			// _ = "end of CoverTab[26820]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:189
			_go_fuzz_dep_.CoverTab[26823]++
//line /usr/local/go/src/compress/gzip/gzip.go:189
			// _ = "end of CoverTab[26823]"
//line /usr/local/go/src/compress/gzip/gzip.go:189
		}
//line /usr/local/go/src/compress/gzip/gzip.go:189
		// _ = "end of CoverTab[26796]"
//line /usr/local/go/src/compress/gzip/gzip.go:189
		_go_fuzz_dep_.CoverTab[26797]++
								if z.compressor == nil {
//line /usr/local/go/src/compress/gzip/gzip.go:190
			_go_fuzz_dep_.CoverTab[26824]++
									z.compressor, _ = flate.NewWriter(z.w, z.level)
//line /usr/local/go/src/compress/gzip/gzip.go:191
			// _ = "end of CoverTab[26824]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:192
			_go_fuzz_dep_.CoverTab[26825]++
//line /usr/local/go/src/compress/gzip/gzip.go:192
			// _ = "end of CoverTab[26825]"
//line /usr/local/go/src/compress/gzip/gzip.go:192
		}
//line /usr/local/go/src/compress/gzip/gzip.go:192
		// _ = "end of CoverTab[26797]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:193
		_go_fuzz_dep_.CoverTab[26826]++
//line /usr/local/go/src/compress/gzip/gzip.go:193
		// _ = "end of CoverTab[26826]"
//line /usr/local/go/src/compress/gzip/gzip.go:193
	}
//line /usr/local/go/src/compress/gzip/gzip.go:193
	// _ = "end of CoverTab[26784]"
//line /usr/local/go/src/compress/gzip/gzip.go:193
	_go_fuzz_dep_.CoverTab[26785]++
							z.size += uint32(len(p))
							z.digest = crc32.Update(z.digest, crc32.IEEETable, p)
							n, z.err = z.compressor.Write(p)
							return n, z.err
//line /usr/local/go/src/compress/gzip/gzip.go:197
	// _ = "end of CoverTab[26785]"
}

// Flush flushes any pending compressed data to the underlying writer.
//line /usr/local/go/src/compress/gzip/gzip.go:200
//
//line /usr/local/go/src/compress/gzip/gzip.go:200
// It is useful mainly in compressed network protocols, to ensure that
//line /usr/local/go/src/compress/gzip/gzip.go:200
// a remote reader has enough data to reconstruct a packet. Flush does
//line /usr/local/go/src/compress/gzip/gzip.go:200
// not return until the data has been written. If the underlying
//line /usr/local/go/src/compress/gzip/gzip.go:200
// writer returns an error, Flush returns that error.
//line /usr/local/go/src/compress/gzip/gzip.go:200
//
//line /usr/local/go/src/compress/gzip/gzip.go:200
// In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
//line /usr/local/go/src/compress/gzip/gzip.go:208
func (z *Writer) Flush() error {
//line /usr/local/go/src/compress/gzip/gzip.go:208
	_go_fuzz_dep_.CoverTab[26827]++
							if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:209
		_go_fuzz_dep_.CoverTab[26831]++
								return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:210
		// _ = "end of CoverTab[26831]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:211
		_go_fuzz_dep_.CoverTab[26832]++
//line /usr/local/go/src/compress/gzip/gzip.go:211
		// _ = "end of CoverTab[26832]"
//line /usr/local/go/src/compress/gzip/gzip.go:211
	}
//line /usr/local/go/src/compress/gzip/gzip.go:211
	// _ = "end of CoverTab[26827]"
//line /usr/local/go/src/compress/gzip/gzip.go:211
	_go_fuzz_dep_.CoverTab[26828]++
							if z.closed {
//line /usr/local/go/src/compress/gzip/gzip.go:212
		_go_fuzz_dep_.CoverTab[26833]++
								return nil
//line /usr/local/go/src/compress/gzip/gzip.go:213
		// _ = "end of CoverTab[26833]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:214
		_go_fuzz_dep_.CoverTab[26834]++
//line /usr/local/go/src/compress/gzip/gzip.go:214
		// _ = "end of CoverTab[26834]"
//line /usr/local/go/src/compress/gzip/gzip.go:214
	}
//line /usr/local/go/src/compress/gzip/gzip.go:214
	// _ = "end of CoverTab[26828]"
//line /usr/local/go/src/compress/gzip/gzip.go:214
	_go_fuzz_dep_.CoverTab[26829]++
							if !z.wroteHeader {
//line /usr/local/go/src/compress/gzip/gzip.go:215
		_go_fuzz_dep_.CoverTab[26835]++
								z.Write(nil)
								if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:217
			_go_fuzz_dep_.CoverTab[26836]++
									return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:218
			// _ = "end of CoverTab[26836]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:219
			_go_fuzz_dep_.CoverTab[26837]++
//line /usr/local/go/src/compress/gzip/gzip.go:219
			// _ = "end of CoverTab[26837]"
//line /usr/local/go/src/compress/gzip/gzip.go:219
		}
//line /usr/local/go/src/compress/gzip/gzip.go:219
		// _ = "end of CoverTab[26835]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:220
		_go_fuzz_dep_.CoverTab[26838]++
//line /usr/local/go/src/compress/gzip/gzip.go:220
		// _ = "end of CoverTab[26838]"
//line /usr/local/go/src/compress/gzip/gzip.go:220
	}
//line /usr/local/go/src/compress/gzip/gzip.go:220
	// _ = "end of CoverTab[26829]"
//line /usr/local/go/src/compress/gzip/gzip.go:220
	_go_fuzz_dep_.CoverTab[26830]++
							z.err = z.compressor.Flush()
							return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:222
	// _ = "end of CoverTab[26830]"
}

// Close closes the Writer by flushing any unwritten data to the underlying
//line /usr/local/go/src/compress/gzip/gzip.go:225
// io.Writer and writing the GZIP footer.
//line /usr/local/go/src/compress/gzip/gzip.go:225
// It does not close the underlying io.Writer.
//line /usr/local/go/src/compress/gzip/gzip.go:228
func (z *Writer) Close() error {
//line /usr/local/go/src/compress/gzip/gzip.go:228
	_go_fuzz_dep_.CoverTab[26839]++
							if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:229
		_go_fuzz_dep_.CoverTab[26844]++
								return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:230
		// _ = "end of CoverTab[26844]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:231
		_go_fuzz_dep_.CoverTab[26845]++
//line /usr/local/go/src/compress/gzip/gzip.go:231
		// _ = "end of CoverTab[26845]"
//line /usr/local/go/src/compress/gzip/gzip.go:231
	}
//line /usr/local/go/src/compress/gzip/gzip.go:231
	// _ = "end of CoverTab[26839]"
//line /usr/local/go/src/compress/gzip/gzip.go:231
	_go_fuzz_dep_.CoverTab[26840]++
							if z.closed {
//line /usr/local/go/src/compress/gzip/gzip.go:232
		_go_fuzz_dep_.CoverTab[26846]++
								return nil
//line /usr/local/go/src/compress/gzip/gzip.go:233
		// _ = "end of CoverTab[26846]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:234
		_go_fuzz_dep_.CoverTab[26847]++
//line /usr/local/go/src/compress/gzip/gzip.go:234
		// _ = "end of CoverTab[26847]"
//line /usr/local/go/src/compress/gzip/gzip.go:234
	}
//line /usr/local/go/src/compress/gzip/gzip.go:234
	// _ = "end of CoverTab[26840]"
//line /usr/local/go/src/compress/gzip/gzip.go:234
	_go_fuzz_dep_.CoverTab[26841]++
							z.closed = true
							if !z.wroteHeader {
//line /usr/local/go/src/compress/gzip/gzip.go:236
		_go_fuzz_dep_.CoverTab[26848]++
								z.Write(nil)
								if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:238
			_go_fuzz_dep_.CoverTab[26849]++
									return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:239
			// _ = "end of CoverTab[26849]"
		} else {
//line /usr/local/go/src/compress/gzip/gzip.go:240
			_go_fuzz_dep_.CoverTab[26850]++
//line /usr/local/go/src/compress/gzip/gzip.go:240
			// _ = "end of CoverTab[26850]"
//line /usr/local/go/src/compress/gzip/gzip.go:240
		}
//line /usr/local/go/src/compress/gzip/gzip.go:240
		// _ = "end of CoverTab[26848]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:241
		_go_fuzz_dep_.CoverTab[26851]++
//line /usr/local/go/src/compress/gzip/gzip.go:241
		// _ = "end of CoverTab[26851]"
//line /usr/local/go/src/compress/gzip/gzip.go:241
	}
//line /usr/local/go/src/compress/gzip/gzip.go:241
	// _ = "end of CoverTab[26841]"
//line /usr/local/go/src/compress/gzip/gzip.go:241
	_go_fuzz_dep_.CoverTab[26842]++
							z.err = z.compressor.Close()
							if z.err != nil {
//line /usr/local/go/src/compress/gzip/gzip.go:243
		_go_fuzz_dep_.CoverTab[26852]++
								return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:244
		// _ = "end of CoverTab[26852]"
	} else {
//line /usr/local/go/src/compress/gzip/gzip.go:245
		_go_fuzz_dep_.CoverTab[26853]++
//line /usr/local/go/src/compress/gzip/gzip.go:245
		// _ = "end of CoverTab[26853]"
//line /usr/local/go/src/compress/gzip/gzip.go:245
	}
//line /usr/local/go/src/compress/gzip/gzip.go:245
	// _ = "end of CoverTab[26842]"
//line /usr/local/go/src/compress/gzip/gzip.go:245
	_go_fuzz_dep_.CoverTab[26843]++
							le.PutUint32(z.buf[:4], z.digest)
							le.PutUint32(z.buf[4:8], z.size)
							_, z.err = z.w.Write(z.buf[:8])
							return z.err
//line /usr/local/go/src/compress/gzip/gzip.go:249
	// _ = "end of CoverTab[26843]"
}

//line /usr/local/go/src/compress/gzip/gzip.go:250
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/gzip/gzip.go:250
var _ = _go_fuzz_dep_.CoverTab
