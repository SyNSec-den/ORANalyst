// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:5
)

import (
	"bytes"
	"errors"
	"io"
)

// A Decoder reads and decodes JSON objects from an input stream.
type Decoder struct {
	r	io.Reader
	buf	[]byte
	d	decodeState
	scanp	int	// start of unread data in buf
	scan	scanner
	err	error

	tokenState	int
	tokenStack	[]int
}

// NewDecoder returns a new decoder that reads from r.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:26
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:26
// The decoder introduces its own buffering and may
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:26
// read data from r beyond the JSON values requested.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:30
func NewDecoder(r io.Reader) *Decoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:30
	_go_fuzz_dep_.CoverTab[188737]++
											return &Decoder{r: r}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:31
	// _ = "end of CoverTab[188737]"
}

// UseNumber causes the Decoder to unmarshal a number into an interface{} as a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:34
// Number instead of as a float64.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:36
func (dec *Decoder) UseNumber() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:36
	_go_fuzz_dep_.CoverTab[188738]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:36
	dec.d.useNumber = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:36
	// _ = "end of CoverTab[188738]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:36
}

// Decode reads the next JSON-encoded value from its
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:38
// input and stores it in the value pointed to by v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:38
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:38
// See the documentation for Unmarshal for details about
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:38
// the conversion of JSON into a Go value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:43
func (dec *Decoder) Decode(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:43
	_go_fuzz_dep_.CoverTab[188739]++
											if dec.err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:44
		_go_fuzz_dep_.CoverTab[188744]++
												return dec.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:45
		// _ = "end of CoverTab[188744]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:46
		_go_fuzz_dep_.CoverTab[188745]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:46
		// _ = "end of CoverTab[188745]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:46
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:46
	// _ = "end of CoverTab[188739]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:46
	_go_fuzz_dep_.CoverTab[188740]++

											if err := dec.tokenPrepareForDecode(); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:48
		_go_fuzz_dep_.CoverTab[188746]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:49
		// _ = "end of CoverTab[188746]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:50
		_go_fuzz_dep_.CoverTab[188747]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:50
		// _ = "end of CoverTab[188747]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:50
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:50
	// _ = "end of CoverTab[188740]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:50
	_go_fuzz_dep_.CoverTab[188741]++

											if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:52
		_go_fuzz_dep_.CoverTab[188748]++
												return &SyntaxError{msg: "not at beginning of value"}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:53
		// _ = "end of CoverTab[188748]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:54
		_go_fuzz_dep_.CoverTab[188749]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:54
		// _ = "end of CoverTab[188749]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:54
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:54
	// _ = "end of CoverTab[188741]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:54
	_go_fuzz_dep_.CoverTab[188742]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:57
	n, err := dec.readValue()
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:58
		_go_fuzz_dep_.CoverTab[188750]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:59
		// _ = "end of CoverTab[188750]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:60
		_go_fuzz_dep_.CoverTab[188751]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:60
		// _ = "end of CoverTab[188751]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:60
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:60
	// _ = "end of CoverTab[188742]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:60
	_go_fuzz_dep_.CoverTab[188743]++
											dec.d.init(dec.buf[dec.scanp : dec.scanp+n])
											dec.scanp += n

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:67
	err = dec.d.unmarshal(v)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:70
	dec.tokenValueEnd()

											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:72
	// _ = "end of CoverTab[188743]"
}

// Buffered returns a reader of the data remaining in the Decoder's
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:75
// buffer. The reader is valid until the next call to Decode.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:77
func (dec *Decoder) Buffered() io.Reader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:77
	_go_fuzz_dep_.CoverTab[188752]++
											return bytes.NewReader(dec.buf[dec.scanp:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:78
	// _ = "end of CoverTab[188752]"
}

// readValue reads a JSON value into dec.buf.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:81
// It returns the length of the encoding.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:83
func (dec *Decoder) readValue() (int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:83
	_go_fuzz_dep_.CoverTab[188753]++
											dec.scan.reset()

											scanp := dec.scanp
											var err error
Input:
	for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:89
		_go_fuzz_dep_.CoverTab[188755]++

												for i, c := range dec.buf[scanp:] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:91
			_go_fuzz_dep_.CoverTab[188758]++
													dec.scan.bytes++
													v := dec.scan.step(&dec.scan, c)
													if v == scanEnd {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:94
				_go_fuzz_dep_.CoverTab[188761]++
														scanp += i
														break Input
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:96
				// _ = "end of CoverTab[188761]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:97
				_go_fuzz_dep_.CoverTab[188762]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:97
				// _ = "end of CoverTab[188762]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:97
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:97
			// _ = "end of CoverTab[188758]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:97
			_go_fuzz_dep_.CoverTab[188759]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
			if (v == scanEndObject || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				_go_fuzz_dep_.CoverTab[188763]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				return v == scanEndArray
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				// _ = "end of CoverTab[188763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				_go_fuzz_dep_.CoverTab[188764]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				return dec.scan.step(&dec.scan, ' ') == scanEnd
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				// _ = "end of CoverTab[188764]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:101
				_go_fuzz_dep_.CoverTab[188765]++
														scanp += i + 1
														break Input
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:103
				// _ = "end of CoverTab[188765]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:104
				_go_fuzz_dep_.CoverTab[188766]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:104
				// _ = "end of CoverTab[188766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:104
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:104
			// _ = "end of CoverTab[188759]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:104
			_go_fuzz_dep_.CoverTab[188760]++
													if v == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:105
				_go_fuzz_dep_.CoverTab[188767]++
														dec.err = dec.scan.err
														return 0, dec.scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:107
				// _ = "end of CoverTab[188767]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:108
				_go_fuzz_dep_.CoverTab[188768]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:108
				// _ = "end of CoverTab[188768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:108
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:108
			// _ = "end of CoverTab[188760]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:109
		// _ = "end of CoverTab[188755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:109
		_go_fuzz_dep_.CoverTab[188756]++
												scanp = len(dec.buf)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:114
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:114
			_go_fuzz_dep_.CoverTab[188769]++
													if err == io.EOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:115
				_go_fuzz_dep_.CoverTab[188771]++
														if dec.scan.step(&dec.scan, ' ') == scanEnd {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:116
					_go_fuzz_dep_.CoverTab[188773]++
															break Input
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:117
					// _ = "end of CoverTab[188773]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:118
					_go_fuzz_dep_.CoverTab[188774]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:118
					// _ = "end of CoverTab[188774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:118
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:118
				// _ = "end of CoverTab[188771]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:118
				_go_fuzz_dep_.CoverTab[188772]++
														if nonSpace(dec.buf) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:119
					_go_fuzz_dep_.CoverTab[188775]++
															err = io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:120
					// _ = "end of CoverTab[188775]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:121
					_go_fuzz_dep_.CoverTab[188776]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:121
					// _ = "end of CoverTab[188776]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:121
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:121
				// _ = "end of CoverTab[188772]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:122
				_go_fuzz_dep_.CoverTab[188777]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:122
				// _ = "end of CoverTab[188777]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:122
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:122
			// _ = "end of CoverTab[188769]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:122
			_go_fuzz_dep_.CoverTab[188770]++
													dec.err = err
													return 0, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:124
			// _ = "end of CoverTab[188770]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:125
			_go_fuzz_dep_.CoverTab[188778]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:125
			// _ = "end of CoverTab[188778]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:125
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:125
		// _ = "end of CoverTab[188756]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:125
		_go_fuzz_dep_.CoverTab[188757]++

												n := scanp - dec.scanp
												err = dec.refill()
												scanp = dec.scanp + n
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:129
		// _ = "end of CoverTab[188757]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:130
	// _ = "end of CoverTab[188753]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:130
	_go_fuzz_dep_.CoverTab[188754]++
											return scanp - dec.scanp, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:131
	// _ = "end of CoverTab[188754]"
}

func (dec *Decoder) refill() error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:134
	_go_fuzz_dep_.CoverTab[188779]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:137
	if dec.scanp > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:137
		_go_fuzz_dep_.CoverTab[188782]++
												n := copy(dec.buf, dec.buf[dec.scanp:])
												dec.buf = dec.buf[:n]
												dec.scanp = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:140
		// _ = "end of CoverTab[188782]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:141
		_go_fuzz_dep_.CoverTab[188783]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:141
		// _ = "end of CoverTab[188783]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:141
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:141
	// _ = "end of CoverTab[188779]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:141
	_go_fuzz_dep_.CoverTab[188780]++

	// Grow buffer if not large enough.
	const minRead = 512
	if cap(dec.buf)-len(dec.buf) < minRead {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:145
		_go_fuzz_dep_.CoverTab[188784]++
												newBuf := make([]byte, len(dec.buf), 2*cap(dec.buf)+minRead)
												copy(newBuf, dec.buf)
												dec.buf = newBuf
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:148
		// _ = "end of CoverTab[188784]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:149
		_go_fuzz_dep_.CoverTab[188785]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:149
		// _ = "end of CoverTab[188785]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:149
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:149
	// _ = "end of CoverTab[188780]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:149
	_go_fuzz_dep_.CoverTab[188781]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:152
	n, err := dec.r.Read(dec.buf[len(dec.buf):cap(dec.buf)])
											dec.buf = dec.buf[0 : len(dec.buf)+n]

											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:155
	// _ = "end of CoverTab[188781]"
}

func nonSpace(b []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:158
	_go_fuzz_dep_.CoverTab[188786]++
											for _, c := range b {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:159
		_go_fuzz_dep_.CoverTab[188788]++
												if !isSpace(c) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:160
			_go_fuzz_dep_.CoverTab[188789]++
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:161
			// _ = "end of CoverTab[188789]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:162
			_go_fuzz_dep_.CoverTab[188790]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:162
			// _ = "end of CoverTab[188790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:162
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:162
		// _ = "end of CoverTab[188788]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:163
	// _ = "end of CoverTab[188786]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:163
	_go_fuzz_dep_.CoverTab[188787]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:164
	// _ = "end of CoverTab[188787]"
}

// An Encoder writes JSON objects to an output stream.
type Encoder struct {
	w	io.Writer
	err	error
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:174
	_go_fuzz_dep_.CoverTab[188791]++
											return &Encoder{w: w}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:175
	// _ = "end of CoverTab[188791]"
}

// Encode writes the JSON encoding of v to the stream,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:178
// followed by a newline character.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:178
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:178
// See the documentation for Marshal for details about the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:178
// conversion of Go values to JSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:183
func (enc *Encoder) Encode(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:183
	_go_fuzz_dep_.CoverTab[188792]++
											if enc.err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:184
		_go_fuzz_dep_.CoverTab[188796]++
												return enc.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:185
		// _ = "end of CoverTab[188796]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:186
		_go_fuzz_dep_.CoverTab[188797]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:186
		// _ = "end of CoverTab[188797]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:186
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:186
	// _ = "end of CoverTab[188792]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:186
	_go_fuzz_dep_.CoverTab[188793]++
											e := newEncodeState()
											err := e.marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:189
		_go_fuzz_dep_.CoverTab[188798]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:190
		// _ = "end of CoverTab[188798]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:191
		_go_fuzz_dep_.CoverTab[188799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:191
		// _ = "end of CoverTab[188799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:191
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:191
	// _ = "end of CoverTab[188793]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:191
	_go_fuzz_dep_.CoverTab[188794]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:199
	e.WriteByte('\n')

	if _, err = enc.w.Write(e.Bytes()); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:201
		_go_fuzz_dep_.CoverTab[188800]++
												enc.err = err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:202
		// _ = "end of CoverTab[188800]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:203
		_go_fuzz_dep_.CoverTab[188801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:203
		// _ = "end of CoverTab[188801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:203
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:203
	// _ = "end of CoverTab[188794]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:203
	_go_fuzz_dep_.CoverTab[188795]++
											encodeStatePool.Put(e)
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:205
	// _ = "end of CoverTab[188795]"
}

// RawMessage is a raw encoded JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:208
// It implements Marshaler and Unmarshaler and can
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:208
// be used to delay JSON decoding or precompute a JSON encoding.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:211
type RawMessage []byte

// MarshalJSON returns *m as the JSON encoding of m.
func (m *RawMessage) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:214
	_go_fuzz_dep_.CoverTab[188802]++
											return *m, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:215
	// _ = "end of CoverTab[188802]"
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:219
	_go_fuzz_dep_.CoverTab[188803]++
											if m == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:220
		_go_fuzz_dep_.CoverTab[188805]++
												return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:221
		// _ = "end of CoverTab[188805]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:222
		_go_fuzz_dep_.CoverTab[188806]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:222
		// _ = "end of CoverTab[188806]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:222
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:222
	// _ = "end of CoverTab[188803]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:222
	_go_fuzz_dep_.CoverTab[188804]++
											*m = append((*m)[0:0], data...)
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:224
	// _ = "end of CoverTab[188804]"
}

var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)

// A Token holds a value of one of these types:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//	Delim, for the four JSON delimiters [ ] { }
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//	bool, for JSON booleans
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//	float64, for JSON numbers
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//	Number, for JSON numbers
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//	string, for JSON string literals
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:230
//	nil, for JSON null
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:239
type Token interface{}

const (
	tokenTopValue	= iota
	tokenArrayStart
	tokenArrayValue
	tokenArrayComma
	tokenObjectStart
	tokenObjectKey
	tokenObjectColon
	tokenObjectValue
	tokenObjectComma
)

// advance tokenstate from a separator state to a value state
func (dec *Decoder) tokenPrepareForDecode() error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:254
	_go_fuzz_dep_.CoverTab[188807]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:258
	switch dec.tokenState {
	case tokenArrayComma:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:259
		_go_fuzz_dep_.CoverTab[188809]++
												c, err := dec.peek()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:261
			_go_fuzz_dep_.CoverTab[188816]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:262
			// _ = "end of CoverTab[188816]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:263
			_go_fuzz_dep_.CoverTab[188817]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:263
			// _ = "end of CoverTab[188817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:263
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:263
		// _ = "end of CoverTab[188809]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:263
		_go_fuzz_dep_.CoverTab[188810]++
												if c != ',' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:264
			_go_fuzz_dep_.CoverTab[188818]++
													return &SyntaxError{"expected comma after array element", 0}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:265
			// _ = "end of CoverTab[188818]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:266
			_go_fuzz_dep_.CoverTab[188819]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:266
			// _ = "end of CoverTab[188819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:266
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:266
		// _ = "end of CoverTab[188810]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:266
		_go_fuzz_dep_.CoverTab[188811]++
												dec.scanp++
												dec.tokenState = tokenArrayValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:268
		// _ = "end of CoverTab[188811]"
	case tokenObjectColon:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:269
		_go_fuzz_dep_.CoverTab[188812]++
												c, err := dec.peek()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:271
			_go_fuzz_dep_.CoverTab[188820]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:272
			// _ = "end of CoverTab[188820]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:273
			_go_fuzz_dep_.CoverTab[188821]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:273
			// _ = "end of CoverTab[188821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:273
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:273
		// _ = "end of CoverTab[188812]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:273
		_go_fuzz_dep_.CoverTab[188813]++
												if c != ':' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:274
			_go_fuzz_dep_.CoverTab[188822]++
													return &SyntaxError{"expected colon after object key", 0}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:275
			// _ = "end of CoverTab[188822]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:276
			_go_fuzz_dep_.CoverTab[188823]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:276
			// _ = "end of CoverTab[188823]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:276
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:276
		// _ = "end of CoverTab[188813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:276
		_go_fuzz_dep_.CoverTab[188814]++
												dec.scanp++
												dec.tokenState = tokenObjectValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:278
		// _ = "end of CoverTab[188814]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:278
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:278
		_go_fuzz_dep_.CoverTab[188815]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:278
		// _ = "end of CoverTab[188815]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:279
	// _ = "end of CoverTab[188807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:279
	_go_fuzz_dep_.CoverTab[188808]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:280
	// _ = "end of CoverTab[188808]"
}

func (dec *Decoder) tokenValueAllowed() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:283
	_go_fuzz_dep_.CoverTab[188824]++
											switch dec.tokenState {
	case tokenTopValue, tokenArrayStart, tokenArrayValue, tokenObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:285
		_go_fuzz_dep_.CoverTab[188826]++
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:286
		// _ = "end of CoverTab[188826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:286
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:286
		_go_fuzz_dep_.CoverTab[188827]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:286
		// _ = "end of CoverTab[188827]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:287
	// _ = "end of CoverTab[188824]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:287
	_go_fuzz_dep_.CoverTab[188825]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:288
	// _ = "end of CoverTab[188825]"
}

func (dec *Decoder) tokenValueEnd() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:291
	_go_fuzz_dep_.CoverTab[188828]++
											switch dec.tokenState {
	case tokenArrayStart, tokenArrayValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:293
		_go_fuzz_dep_.CoverTab[188829]++
												dec.tokenState = tokenArrayComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:294
		// _ = "end of CoverTab[188829]"
	case tokenObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:295
		_go_fuzz_dep_.CoverTab[188830]++
												dec.tokenState = tokenObjectComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:296
		// _ = "end of CoverTab[188830]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:296
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:296
		_go_fuzz_dep_.CoverTab[188831]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:296
		// _ = "end of CoverTab[188831]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:297
	// _ = "end of CoverTab[188828]"
}

// A Delim is a JSON array or object delimiter, one of [ ] { or }.
type Delim rune

func (d Delim) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:303
	_go_fuzz_dep_.CoverTab[188832]++
											return string(d)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:304
	// _ = "end of CoverTab[188832]"
}

// Token returns the next JSON token in the input stream.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// At the end of the input stream, Token returns nil, io.EOF.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// Token guarantees that the delimiters [ ] { } it returns are
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// properly nested and matched: if Token encounters an unexpected
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// delimiter in the input, it will return an error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// The input stream consists of basic JSON values—bool, string,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// number, and null—along with delimiters [ ] { } of type Delim
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// to mark the start and end of arrays and objects.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:307
// Commas and colons are elided.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:318
func (dec *Decoder) Token() (Token, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:318
	_go_fuzz_dep_.CoverTab[188833]++
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:319
		_go_fuzz_dep_.CoverTab[188834]++
												c, err := dec.peek()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:321
			_go_fuzz_dep_.CoverTab[188836]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:322
			// _ = "end of CoverTab[188836]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:323
			_go_fuzz_dep_.CoverTab[188837]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:323
			// _ = "end of CoverTab[188837]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:323
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:323
		// _ = "end of CoverTab[188834]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:323
		_go_fuzz_dep_.CoverTab[188835]++
												switch c {
		case '[':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:325
			_go_fuzz_dep_.CoverTab[188838]++
													if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:326
				_go_fuzz_dep_.CoverTab[188856]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:327
				// _ = "end of CoverTab[188856]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:328
				_go_fuzz_dep_.CoverTab[188857]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:328
				// _ = "end of CoverTab[188857]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:328
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:328
			// _ = "end of CoverTab[188838]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:328
			_go_fuzz_dep_.CoverTab[188839]++
													dec.scanp++
													dec.tokenStack = append(dec.tokenStack, dec.tokenState)
													dec.tokenState = tokenArrayStart
													return Delim('['), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:332
			// _ = "end of CoverTab[188839]"

		case ']':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:334
			_go_fuzz_dep_.CoverTab[188840]++
													if dec.tokenState != tokenArrayStart && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:335
				_go_fuzz_dep_.CoverTab[188858]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:335
				return dec.tokenState != tokenArrayComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:335
				// _ = "end of CoverTab[188858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:335
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:335
				_go_fuzz_dep_.CoverTab[188859]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:336
				// _ = "end of CoverTab[188859]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:337
				_go_fuzz_dep_.CoverTab[188860]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:337
				// _ = "end of CoverTab[188860]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:337
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:337
			// _ = "end of CoverTab[188840]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:337
			_go_fuzz_dep_.CoverTab[188841]++
													dec.scanp++
													dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
													dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
													dec.tokenValueEnd()
													return Delim(']'), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:342
			// _ = "end of CoverTab[188841]"

		case '{':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:344
			_go_fuzz_dep_.CoverTab[188842]++
													if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:345
				_go_fuzz_dep_.CoverTab[188861]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:346
				// _ = "end of CoverTab[188861]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:347
				_go_fuzz_dep_.CoverTab[188862]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:347
				// _ = "end of CoverTab[188862]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:347
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:347
			// _ = "end of CoverTab[188842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:347
			_go_fuzz_dep_.CoverTab[188843]++
													dec.scanp++
													dec.tokenStack = append(dec.tokenStack, dec.tokenState)
													dec.tokenState = tokenObjectStart
													return Delim('{'), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:351
			// _ = "end of CoverTab[188843]"

		case '}':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:353
			_go_fuzz_dep_.CoverTab[188844]++
													if dec.tokenState != tokenObjectStart && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:354
				_go_fuzz_dep_.CoverTab[188863]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:354
				return dec.tokenState != tokenObjectComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:354
				// _ = "end of CoverTab[188863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:354
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:354
				_go_fuzz_dep_.CoverTab[188864]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:355
				// _ = "end of CoverTab[188864]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:356
				_go_fuzz_dep_.CoverTab[188865]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:356
				// _ = "end of CoverTab[188865]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:356
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:356
			// _ = "end of CoverTab[188844]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:356
			_go_fuzz_dep_.CoverTab[188845]++
													dec.scanp++
													dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
													dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
													dec.tokenValueEnd()
													return Delim('}'), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:361
			// _ = "end of CoverTab[188845]"

		case ':':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:363
			_go_fuzz_dep_.CoverTab[188846]++
													if dec.tokenState != tokenObjectColon {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:364
				_go_fuzz_dep_.CoverTab[188866]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:365
				// _ = "end of CoverTab[188866]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:366
				_go_fuzz_dep_.CoverTab[188867]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:366
				// _ = "end of CoverTab[188867]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:366
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:366
			// _ = "end of CoverTab[188846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:366
			_go_fuzz_dep_.CoverTab[188847]++
													dec.scanp++
													dec.tokenState = tokenObjectValue
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:369
			// _ = "end of CoverTab[188847]"

		case ',':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:371
			_go_fuzz_dep_.CoverTab[188848]++
													if dec.tokenState == tokenArrayComma {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:372
				_go_fuzz_dep_.CoverTab[188868]++
														dec.scanp++
														dec.tokenState = tokenArrayValue
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:375
				// _ = "end of CoverTab[188868]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:376
				_go_fuzz_dep_.CoverTab[188869]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:376
				// _ = "end of CoverTab[188869]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:376
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:376
			// _ = "end of CoverTab[188848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:376
			_go_fuzz_dep_.CoverTab[188849]++
													if dec.tokenState == tokenObjectComma {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:377
				_go_fuzz_dep_.CoverTab[188870]++
														dec.scanp++
														dec.tokenState = tokenObjectKey
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:380
				// _ = "end of CoverTab[188870]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:381
				_go_fuzz_dep_.CoverTab[188871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:381
				// _ = "end of CoverTab[188871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:381
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:381
			// _ = "end of CoverTab[188849]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:381
			_go_fuzz_dep_.CoverTab[188850]++
													return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:382
			// _ = "end of CoverTab[188850]"

		case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:384
			_go_fuzz_dep_.CoverTab[188851]++
													if dec.tokenState == tokenObjectStart || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:385
				_go_fuzz_dep_.CoverTab[188872]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:385
				return dec.tokenState == tokenObjectKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:385
				// _ = "end of CoverTab[188872]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:385
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:385
				_go_fuzz_dep_.CoverTab[188873]++
														var x string
														old := dec.tokenState
														dec.tokenState = tokenTopValue
														err := dec.Decode(&x)
														dec.tokenState = old
														if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:391
					_go_fuzz_dep_.CoverTab[188875]++
															clearOffset(err)
															return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:393
					// _ = "end of CoverTab[188875]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:394
					_go_fuzz_dep_.CoverTab[188876]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:394
					// _ = "end of CoverTab[188876]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:394
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:394
				// _ = "end of CoverTab[188873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:394
				_go_fuzz_dep_.CoverTab[188874]++
														dec.tokenState = tokenObjectColon
														return x, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:396
				// _ = "end of CoverTab[188874]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:397
				_go_fuzz_dep_.CoverTab[188877]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:397
				// _ = "end of CoverTab[188877]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:397
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:397
			// _ = "end of CoverTab[188851]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:397
			_go_fuzz_dep_.CoverTab[188852]++
													fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:398
			// _ = "end of CoverTab[188852]"

		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:400
			_go_fuzz_dep_.CoverTab[188853]++
													if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:401
				_go_fuzz_dep_.CoverTab[188878]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:402
				// _ = "end of CoverTab[188878]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:403
				_go_fuzz_dep_.CoverTab[188879]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:403
				// _ = "end of CoverTab[188879]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:403
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:403
			// _ = "end of CoverTab[188853]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:403
			_go_fuzz_dep_.CoverTab[188854]++
													var x interface{}
													if err := dec.Decode(&x); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:405
				_go_fuzz_dep_.CoverTab[188880]++
														clearOffset(err)
														return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:407
				// _ = "end of CoverTab[188880]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:408
				_go_fuzz_dep_.CoverTab[188881]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:408
				// _ = "end of CoverTab[188881]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:408
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:408
			// _ = "end of CoverTab[188854]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:408
			_go_fuzz_dep_.CoverTab[188855]++
													return x, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:409
			// _ = "end of CoverTab[188855]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:410
		// _ = "end of CoverTab[188835]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:411
	// _ = "end of CoverTab[188833]"
}

func clearOffset(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:414
	_go_fuzz_dep_.CoverTab[188882]++
											if s, ok := err.(*SyntaxError); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:415
		_go_fuzz_dep_.CoverTab[188883]++
												s.Offset = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:416
		// _ = "end of CoverTab[188883]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:417
		_go_fuzz_dep_.CoverTab[188884]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:417
		// _ = "end of CoverTab[188884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:417
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:417
	// _ = "end of CoverTab[188882]"
}

func (dec *Decoder) tokenError(c byte) (Token, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:420
	_go_fuzz_dep_.CoverTab[188885]++
											var context string
											switch dec.tokenState {
	case tokenTopValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:423
		_go_fuzz_dep_.CoverTab[188887]++
												context = " looking for beginning of value"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:424
		// _ = "end of CoverTab[188887]"
	case tokenArrayStart, tokenArrayValue, tokenObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:425
		_go_fuzz_dep_.CoverTab[188888]++
												context = " looking for beginning of value"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:426
		// _ = "end of CoverTab[188888]"
	case tokenArrayComma:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:427
		_go_fuzz_dep_.CoverTab[188889]++
												context = " after array element"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:428
		// _ = "end of CoverTab[188889]"
	case tokenObjectKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:429
		_go_fuzz_dep_.CoverTab[188890]++
												context = " looking for beginning of object key string"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:430
		// _ = "end of CoverTab[188890]"
	case tokenObjectColon:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:431
		_go_fuzz_dep_.CoverTab[188891]++
												context = " after object key"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:432
		// _ = "end of CoverTab[188891]"
	case tokenObjectComma:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:433
		_go_fuzz_dep_.CoverTab[188892]++
												context = " after object key:value pair"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:434
		// _ = "end of CoverTab[188892]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:434
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:434
		_go_fuzz_dep_.CoverTab[188893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:434
		// _ = "end of CoverTab[188893]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:435
	// _ = "end of CoverTab[188885]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:435
	_go_fuzz_dep_.CoverTab[188886]++
											return nil, &SyntaxError{"invalid character " + quoteChar(c) + " " + context, 0}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:436
	// _ = "end of CoverTab[188886]"
}

// More reports whether there is another element in the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:439
// current array or object being parsed.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:441
func (dec *Decoder) More() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:441
	_go_fuzz_dep_.CoverTab[188894]++
											c, err := dec.peek()
											return err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
		_go_fuzz_dep_.CoverTab[188895]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
		return c != ']'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
		// _ = "end of CoverTab[188895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
		_go_fuzz_dep_.CoverTab[188896]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
		return c != '}'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
		// _ = "end of CoverTab[188896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:443
	// _ = "end of CoverTab[188894]"
}

func (dec *Decoder) peek() (byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:446
	_go_fuzz_dep_.CoverTab[188897]++
											var err error
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:448
		_go_fuzz_dep_.CoverTab[188898]++
												for i := dec.scanp; i < len(dec.buf); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:449
			_go_fuzz_dep_.CoverTab[188901]++
													c := dec.buf[i]
													if isSpace(c) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:451
				_go_fuzz_dep_.CoverTab[188903]++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:452
				// _ = "end of CoverTab[188903]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:453
				_go_fuzz_dep_.CoverTab[188904]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:453
				// _ = "end of CoverTab[188904]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:453
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:453
			// _ = "end of CoverTab[188901]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:453
			_go_fuzz_dep_.CoverTab[188902]++
													dec.scanp = i
													return c, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:455
			// _ = "end of CoverTab[188902]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:456
		// _ = "end of CoverTab[188898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:456
		_go_fuzz_dep_.CoverTab[188899]++

												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:458
			_go_fuzz_dep_.CoverTab[188905]++
													return 0, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:459
			// _ = "end of CoverTab[188905]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:460
			_go_fuzz_dep_.CoverTab[188906]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:460
			// _ = "end of CoverTab[188906]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:460
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:460
		// _ = "end of CoverTab[188899]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:460
		_go_fuzz_dep_.CoverTab[188900]++
												err = dec.refill()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:461
		// _ = "end of CoverTab[188900]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:462
	// _ = "end of CoverTab[188897]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:463
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/json/stream.go:463
var _ = _go_fuzz_dep_.CoverTab
