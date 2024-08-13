// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:5
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:26
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:26
// The decoder introduces its own buffering and may
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:26
// read data from r beyond the JSON values requested.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:30
func NewDecoder(r io.Reader) *Decoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:30
	_go_fuzz_dep_.CoverTab[185679]++
											return &Decoder{r: r}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:31
	// _ = "end of CoverTab[185679]"
}

// UseNumber causes the Decoder to unmarshal a number into an interface{} as a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:34
// Number instead of as a float64.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:36
func (dec *Decoder) UseNumber() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:36
	_go_fuzz_dep_.CoverTab[185680]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:36
	dec.d.useNumber = true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:36
	// _ = "end of CoverTab[185680]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:36
}

// Decode reads the next JSON-encoded value from its
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:38
// input and stores it in the value pointed to by v.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:38
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:38
// See the documentation for Unmarshal for details about
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:38
// the conversion of JSON into a Go value.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:43
func (dec *Decoder) Decode(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:43
	_go_fuzz_dep_.CoverTab[185681]++
											if dec.err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:44
		_go_fuzz_dep_.CoverTab[185686]++
												return dec.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:45
		// _ = "end of CoverTab[185686]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:46
		_go_fuzz_dep_.CoverTab[185687]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:46
		// _ = "end of CoverTab[185687]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:46
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:46
	// _ = "end of CoverTab[185681]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:46
	_go_fuzz_dep_.CoverTab[185682]++

											if err := dec.tokenPrepareForDecode(); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:48
		_go_fuzz_dep_.CoverTab[185688]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:49
		// _ = "end of CoverTab[185688]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:50
		_go_fuzz_dep_.CoverTab[185689]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:50
		// _ = "end of CoverTab[185689]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:50
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:50
	// _ = "end of CoverTab[185682]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:50
	_go_fuzz_dep_.CoverTab[185683]++

											if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:52
		_go_fuzz_dep_.CoverTab[185690]++
												return &SyntaxError{msg: "not at beginning of value"}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:53
		// _ = "end of CoverTab[185690]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:54
		_go_fuzz_dep_.CoverTab[185691]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:54
		// _ = "end of CoverTab[185691]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:54
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:54
	// _ = "end of CoverTab[185683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:54
	_go_fuzz_dep_.CoverTab[185684]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:57
	n, err := dec.readValue()
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:58
		_go_fuzz_dep_.CoverTab[185692]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:59
		// _ = "end of CoverTab[185692]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:60
		_go_fuzz_dep_.CoverTab[185693]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:60
		// _ = "end of CoverTab[185693]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:60
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:60
	// _ = "end of CoverTab[185684]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:60
	_go_fuzz_dep_.CoverTab[185685]++
											dec.d.init(dec.buf[dec.scanp : dec.scanp+n])
											dec.scanp += n

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:67
	err = dec.d.unmarshal(v)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:70
	dec.tokenValueEnd()

											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:72
	// _ = "end of CoverTab[185685]"
}

// Buffered returns a reader of the data remaining in the Decoder's
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:75
// buffer. The reader is valid until the next call to Decode.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:77
func (dec *Decoder) Buffered() io.Reader {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:77
	_go_fuzz_dep_.CoverTab[185694]++
											return bytes.NewReader(dec.buf[dec.scanp:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:78
	// _ = "end of CoverTab[185694]"
}

// readValue reads a JSON value into dec.buf.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:81
// It returns the length of the encoding.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:83
func (dec *Decoder) readValue() (int, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:83
	_go_fuzz_dep_.CoverTab[185695]++
											dec.scan.reset()

											scanp := dec.scanp
											var err error
Input:
	for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:89
		_go_fuzz_dep_.CoverTab[185697]++

												for i, c := range dec.buf[scanp:] {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:91
			_go_fuzz_dep_.CoverTab[185700]++
													dec.scan.bytes++
													v := dec.scan.step(&dec.scan, c)
													if v == scanEnd {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:94
				_go_fuzz_dep_.CoverTab[185703]++
														scanp += i
														break Input
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:96
				// _ = "end of CoverTab[185703]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:97
				_go_fuzz_dep_.CoverTab[185704]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:97
				// _ = "end of CoverTab[185704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:97
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:97
			// _ = "end of CoverTab[185700]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:97
			_go_fuzz_dep_.CoverTab[185701]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
			if (v == scanEndObject || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				_go_fuzz_dep_.CoverTab[185705]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				return v == scanEndArray
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				// _ = "end of CoverTab[185705]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
			}()) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				_go_fuzz_dep_.CoverTab[185706]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				return dec.scan.step(&dec.scan, ' ') == scanEnd
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				// _ = "end of CoverTab[185706]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:101
				_go_fuzz_dep_.CoverTab[185707]++
														scanp += i + 1
														break Input
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:103
				// _ = "end of CoverTab[185707]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:104
				_go_fuzz_dep_.CoverTab[185708]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:104
				// _ = "end of CoverTab[185708]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:104
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:104
			// _ = "end of CoverTab[185701]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:104
			_go_fuzz_dep_.CoverTab[185702]++
													if v == scanError {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:105
				_go_fuzz_dep_.CoverTab[185709]++
														dec.err = dec.scan.err
														return 0, dec.scan.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:107
				// _ = "end of CoverTab[185709]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:108
				_go_fuzz_dep_.CoverTab[185710]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:108
				// _ = "end of CoverTab[185710]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:108
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:108
			// _ = "end of CoverTab[185702]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:109
		// _ = "end of CoverTab[185697]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:109
		_go_fuzz_dep_.CoverTab[185698]++
												scanp = len(dec.buf)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:114
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:114
			_go_fuzz_dep_.CoverTab[185711]++
													if err == io.EOF {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:115
				_go_fuzz_dep_.CoverTab[185713]++
														if dec.scan.step(&dec.scan, ' ') == scanEnd {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:116
					_go_fuzz_dep_.CoverTab[185715]++
															break Input
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:117
					// _ = "end of CoverTab[185715]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:118
					_go_fuzz_dep_.CoverTab[185716]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:118
					// _ = "end of CoverTab[185716]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:118
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:118
				// _ = "end of CoverTab[185713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:118
				_go_fuzz_dep_.CoverTab[185714]++
														if nonSpace(dec.buf) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:119
					_go_fuzz_dep_.CoverTab[185717]++
															err = io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:120
					// _ = "end of CoverTab[185717]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:121
					_go_fuzz_dep_.CoverTab[185718]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:121
					// _ = "end of CoverTab[185718]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:121
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:121
				// _ = "end of CoverTab[185714]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:122
				_go_fuzz_dep_.CoverTab[185719]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:122
				// _ = "end of CoverTab[185719]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:122
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:122
			// _ = "end of CoverTab[185711]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:122
			_go_fuzz_dep_.CoverTab[185712]++
													dec.err = err
													return 0, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:124
			// _ = "end of CoverTab[185712]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:125
			_go_fuzz_dep_.CoverTab[185720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:125
			// _ = "end of CoverTab[185720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:125
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:125
		// _ = "end of CoverTab[185698]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:125
		_go_fuzz_dep_.CoverTab[185699]++

												n := scanp - dec.scanp
												err = dec.refill()
												scanp = dec.scanp + n
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:129
		// _ = "end of CoverTab[185699]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:130
	// _ = "end of CoverTab[185695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:130
	_go_fuzz_dep_.CoverTab[185696]++
											return scanp - dec.scanp, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:131
	// _ = "end of CoverTab[185696]"
}

func (dec *Decoder) refill() error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:134
	_go_fuzz_dep_.CoverTab[185721]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:137
	if dec.scanp > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:137
		_go_fuzz_dep_.CoverTab[185724]++
												n := copy(dec.buf, dec.buf[dec.scanp:])
												dec.buf = dec.buf[:n]
												dec.scanp = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:140
		// _ = "end of CoverTab[185724]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:141
		_go_fuzz_dep_.CoverTab[185725]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:141
		// _ = "end of CoverTab[185725]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:141
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:141
	// _ = "end of CoverTab[185721]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:141
	_go_fuzz_dep_.CoverTab[185722]++

	// Grow buffer if not large enough.
	const minRead = 512
	if cap(dec.buf)-len(dec.buf) < minRead {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:145
		_go_fuzz_dep_.CoverTab[185726]++
												newBuf := make([]byte, len(dec.buf), 2*cap(dec.buf)+minRead)
												copy(newBuf, dec.buf)
												dec.buf = newBuf
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:148
		// _ = "end of CoverTab[185726]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:149
		_go_fuzz_dep_.CoverTab[185727]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:149
		// _ = "end of CoverTab[185727]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:149
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:149
	// _ = "end of CoverTab[185722]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:149
	_go_fuzz_dep_.CoverTab[185723]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:152
	n, err := dec.r.Read(dec.buf[len(dec.buf):cap(dec.buf)])
											dec.buf = dec.buf[0 : len(dec.buf)+n]

											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:155
	// _ = "end of CoverTab[185723]"
}

func nonSpace(b []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:158
	_go_fuzz_dep_.CoverTab[185728]++
											for _, c := range b {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:159
		_go_fuzz_dep_.CoverTab[185730]++
												if !isSpace(c) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:160
			_go_fuzz_dep_.CoverTab[185731]++
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:161
			// _ = "end of CoverTab[185731]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:162
			_go_fuzz_dep_.CoverTab[185732]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:162
			// _ = "end of CoverTab[185732]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:162
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:162
		// _ = "end of CoverTab[185730]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:163
	// _ = "end of CoverTab[185728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:163
	_go_fuzz_dep_.CoverTab[185729]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:164
	// _ = "end of CoverTab[185729]"
}

// An Encoder writes JSON objects to an output stream.
type Encoder struct {
	w	io.Writer
	err	error
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:174
	_go_fuzz_dep_.CoverTab[185733]++
											return &Encoder{w: w}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:175
	// _ = "end of CoverTab[185733]"
}

// Encode writes the JSON encoding of v to the stream,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:178
// followed by a newline character.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:178
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:178
// See the documentation for Marshal for details about the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:178
// conversion of Go values to JSON.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:183
func (enc *Encoder) Encode(v interface{}) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:183
	_go_fuzz_dep_.CoverTab[185734]++
											if enc.err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:184
		_go_fuzz_dep_.CoverTab[185738]++
												return enc.err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:185
		// _ = "end of CoverTab[185738]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:186
		_go_fuzz_dep_.CoverTab[185739]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:186
		// _ = "end of CoverTab[185739]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:186
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:186
	// _ = "end of CoverTab[185734]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:186
	_go_fuzz_dep_.CoverTab[185735]++
											e := newEncodeState()
											err := e.marshal(v)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:189
		_go_fuzz_dep_.CoverTab[185740]++
												return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:190
		// _ = "end of CoverTab[185740]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:191
		_go_fuzz_dep_.CoverTab[185741]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:191
		// _ = "end of CoverTab[185741]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:191
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:191
	// _ = "end of CoverTab[185735]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:191
	_go_fuzz_dep_.CoverTab[185736]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:199
	e.WriteByte('\n')

	if _, err = enc.w.Write(e.Bytes()); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:201
		_go_fuzz_dep_.CoverTab[185742]++
												enc.err = err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:202
		// _ = "end of CoverTab[185742]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:203
		_go_fuzz_dep_.CoverTab[185743]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:203
		// _ = "end of CoverTab[185743]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:203
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:203
	// _ = "end of CoverTab[185736]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:203
	_go_fuzz_dep_.CoverTab[185737]++
											encodeStatePool.Put(e)
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:205
	// _ = "end of CoverTab[185737]"
}

// RawMessage is a raw encoded JSON object.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:208
// It implements Marshaler and Unmarshaler and can
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:208
// be used to delay JSON decoding or precompute a JSON encoding.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:211
type RawMessage []byte

// MarshalJSON returns *m as the JSON encoding of m.
func (m *RawMessage) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:214
	_go_fuzz_dep_.CoverTab[185744]++
											return *m, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:215
	// _ = "end of CoverTab[185744]"
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:219
	_go_fuzz_dep_.CoverTab[185745]++
											if m == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:220
		_go_fuzz_dep_.CoverTab[185747]++
												return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:221
		// _ = "end of CoverTab[185747]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:222
		_go_fuzz_dep_.CoverTab[185748]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:222
		// _ = "end of CoverTab[185748]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:222
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:222
	// _ = "end of CoverTab[185745]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:222
	_go_fuzz_dep_.CoverTab[185746]++
											*m = append((*m)[0:0], data...)
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:224
	// _ = "end of CoverTab[185746]"
}

var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)

// A Token holds a value of one of these types:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//	Delim, for the four JSON delimiters [ ] { }
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//	bool, for JSON booleans
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//	float64, for JSON numbers
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//	Number, for JSON numbers
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//	string, for JSON string literals
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:230
//	nil, for JSON null
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:239
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:254
	_go_fuzz_dep_.CoverTab[185749]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:258
	switch dec.tokenState {
	case tokenArrayComma:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:259
		_go_fuzz_dep_.CoverTab[185751]++
												c, err := dec.peek()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:261
			_go_fuzz_dep_.CoverTab[185758]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:262
			// _ = "end of CoverTab[185758]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:263
			_go_fuzz_dep_.CoverTab[185759]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:263
			// _ = "end of CoverTab[185759]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:263
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:263
		// _ = "end of CoverTab[185751]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:263
		_go_fuzz_dep_.CoverTab[185752]++
												if c != ',' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:264
			_go_fuzz_dep_.CoverTab[185760]++
													return &SyntaxError{"expected comma after array element", 0}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:265
			// _ = "end of CoverTab[185760]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:266
			_go_fuzz_dep_.CoverTab[185761]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:266
			// _ = "end of CoverTab[185761]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:266
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:266
		// _ = "end of CoverTab[185752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:266
		_go_fuzz_dep_.CoverTab[185753]++
												dec.scanp++
												dec.tokenState = tokenArrayValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:268
		// _ = "end of CoverTab[185753]"
	case tokenObjectColon:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:269
		_go_fuzz_dep_.CoverTab[185754]++
												c, err := dec.peek()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:271
			_go_fuzz_dep_.CoverTab[185762]++
													return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:272
			// _ = "end of CoverTab[185762]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:273
			_go_fuzz_dep_.CoverTab[185763]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:273
			// _ = "end of CoverTab[185763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:273
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:273
		// _ = "end of CoverTab[185754]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:273
		_go_fuzz_dep_.CoverTab[185755]++
												if c != ':' {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:274
			_go_fuzz_dep_.CoverTab[185764]++
													return &SyntaxError{"expected colon after object key", 0}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:275
			// _ = "end of CoverTab[185764]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:276
			_go_fuzz_dep_.CoverTab[185765]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:276
			// _ = "end of CoverTab[185765]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:276
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:276
		// _ = "end of CoverTab[185755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:276
		_go_fuzz_dep_.CoverTab[185756]++
												dec.scanp++
												dec.tokenState = tokenObjectValue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:278
		// _ = "end of CoverTab[185756]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:278
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:278
		_go_fuzz_dep_.CoverTab[185757]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:278
		// _ = "end of CoverTab[185757]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:279
	// _ = "end of CoverTab[185749]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:279
	_go_fuzz_dep_.CoverTab[185750]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:280
	// _ = "end of CoverTab[185750]"
}

func (dec *Decoder) tokenValueAllowed() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:283
	_go_fuzz_dep_.CoverTab[185766]++
											switch dec.tokenState {
	case tokenTopValue, tokenArrayStart, tokenArrayValue, tokenObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:285
		_go_fuzz_dep_.CoverTab[185768]++
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:286
		// _ = "end of CoverTab[185768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:286
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:286
		_go_fuzz_dep_.CoverTab[185769]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:286
		// _ = "end of CoverTab[185769]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:287
	// _ = "end of CoverTab[185766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:287
	_go_fuzz_dep_.CoverTab[185767]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:288
	// _ = "end of CoverTab[185767]"
}

func (dec *Decoder) tokenValueEnd() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:291
	_go_fuzz_dep_.CoverTab[185770]++
											switch dec.tokenState {
	case tokenArrayStart, tokenArrayValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:293
		_go_fuzz_dep_.CoverTab[185771]++
												dec.tokenState = tokenArrayComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:294
		// _ = "end of CoverTab[185771]"
	case tokenObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:295
		_go_fuzz_dep_.CoverTab[185772]++
												dec.tokenState = tokenObjectComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:296
		// _ = "end of CoverTab[185772]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:296
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:296
		_go_fuzz_dep_.CoverTab[185773]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:296
		// _ = "end of CoverTab[185773]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:297
	// _ = "end of CoverTab[185770]"
}

// A Delim is a JSON array or object delimiter, one of [ ] { or }.
type Delim rune

func (d Delim) String() string {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:303
	_go_fuzz_dep_.CoverTab[185774]++
											return string(d)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:304
	// _ = "end of CoverTab[185774]"
}

// Token returns the next JSON token in the input stream.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// At the end of the input stream, Token returns nil, io.EOF.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// Token guarantees that the delimiters [ ] { } it returns are
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// properly nested and matched: if Token encounters an unexpected
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// delimiter in the input, it will return an error.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// The input stream consists of basic JSON values—bool, string,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// number, and null—along with delimiters [ ] { } of type Delim
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// to mark the start and end of arrays and objects.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:307
// Commas and colons are elided.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:318
func (dec *Decoder) Token() (Token, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:318
	_go_fuzz_dep_.CoverTab[185775]++
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:319
		_go_fuzz_dep_.CoverTab[185776]++
												c, err := dec.peek()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:321
			_go_fuzz_dep_.CoverTab[185778]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:322
			// _ = "end of CoverTab[185778]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:323
			_go_fuzz_dep_.CoverTab[185779]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:323
			// _ = "end of CoverTab[185779]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:323
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:323
		// _ = "end of CoverTab[185776]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:323
		_go_fuzz_dep_.CoverTab[185777]++
												switch c {
		case '[':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:325
			_go_fuzz_dep_.CoverTab[185780]++
													if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:326
				_go_fuzz_dep_.CoverTab[185798]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:327
				// _ = "end of CoverTab[185798]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:328
				_go_fuzz_dep_.CoverTab[185799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:328
				// _ = "end of CoverTab[185799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:328
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:328
			// _ = "end of CoverTab[185780]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:328
			_go_fuzz_dep_.CoverTab[185781]++
													dec.scanp++
													dec.tokenStack = append(dec.tokenStack, dec.tokenState)
													dec.tokenState = tokenArrayStart
													return Delim('['), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:332
			// _ = "end of CoverTab[185781]"

		case ']':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:334
			_go_fuzz_dep_.CoverTab[185782]++
													if dec.tokenState != tokenArrayStart && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:335
				_go_fuzz_dep_.CoverTab[185800]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:335
				return dec.tokenState != tokenArrayComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:335
				// _ = "end of CoverTab[185800]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:335
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:335
				_go_fuzz_dep_.CoverTab[185801]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:336
				// _ = "end of CoverTab[185801]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:337
				_go_fuzz_dep_.CoverTab[185802]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:337
				// _ = "end of CoverTab[185802]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:337
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:337
			// _ = "end of CoverTab[185782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:337
			_go_fuzz_dep_.CoverTab[185783]++
													dec.scanp++
													dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
													dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
													dec.tokenValueEnd()
													return Delim(']'), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:342
			// _ = "end of CoverTab[185783]"

		case '{':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:344
			_go_fuzz_dep_.CoverTab[185784]++
													if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:345
				_go_fuzz_dep_.CoverTab[185803]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:346
				// _ = "end of CoverTab[185803]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:347
				_go_fuzz_dep_.CoverTab[185804]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:347
				// _ = "end of CoverTab[185804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:347
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:347
			// _ = "end of CoverTab[185784]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:347
			_go_fuzz_dep_.CoverTab[185785]++
													dec.scanp++
													dec.tokenStack = append(dec.tokenStack, dec.tokenState)
													dec.tokenState = tokenObjectStart
													return Delim('{'), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:351
			// _ = "end of CoverTab[185785]"

		case '}':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:353
			_go_fuzz_dep_.CoverTab[185786]++
													if dec.tokenState != tokenObjectStart && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:354
				_go_fuzz_dep_.CoverTab[185805]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:354
				return dec.tokenState != tokenObjectComma
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:354
				// _ = "end of CoverTab[185805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:354
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:354
				_go_fuzz_dep_.CoverTab[185806]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:355
				// _ = "end of CoverTab[185806]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:356
				_go_fuzz_dep_.CoverTab[185807]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:356
				// _ = "end of CoverTab[185807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:356
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:356
			// _ = "end of CoverTab[185786]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:356
			_go_fuzz_dep_.CoverTab[185787]++
													dec.scanp++
													dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
													dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
													dec.tokenValueEnd()
													return Delim('}'), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:361
			// _ = "end of CoverTab[185787]"

		case ':':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:363
			_go_fuzz_dep_.CoverTab[185788]++
													if dec.tokenState != tokenObjectColon {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:364
				_go_fuzz_dep_.CoverTab[185808]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:365
				// _ = "end of CoverTab[185808]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:366
				_go_fuzz_dep_.CoverTab[185809]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:366
				// _ = "end of CoverTab[185809]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:366
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:366
			// _ = "end of CoverTab[185788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:366
			_go_fuzz_dep_.CoverTab[185789]++
													dec.scanp++
													dec.tokenState = tokenObjectValue
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:369
			// _ = "end of CoverTab[185789]"

		case ',':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:371
			_go_fuzz_dep_.CoverTab[185790]++
													if dec.tokenState == tokenArrayComma {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:372
				_go_fuzz_dep_.CoverTab[185810]++
														dec.scanp++
														dec.tokenState = tokenArrayValue
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:375
				// _ = "end of CoverTab[185810]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:376
				_go_fuzz_dep_.CoverTab[185811]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:376
				// _ = "end of CoverTab[185811]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:376
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:376
			// _ = "end of CoverTab[185790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:376
			_go_fuzz_dep_.CoverTab[185791]++
													if dec.tokenState == tokenObjectComma {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:377
				_go_fuzz_dep_.CoverTab[185812]++
														dec.scanp++
														dec.tokenState = tokenObjectKey
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:380
				// _ = "end of CoverTab[185812]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:381
				_go_fuzz_dep_.CoverTab[185813]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:381
				// _ = "end of CoverTab[185813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:381
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:381
			// _ = "end of CoverTab[185791]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:381
			_go_fuzz_dep_.CoverTab[185792]++
													return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:382
			// _ = "end of CoverTab[185792]"

		case '"':
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:384
			_go_fuzz_dep_.CoverTab[185793]++
													if dec.tokenState == tokenObjectStart || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:385
				_go_fuzz_dep_.CoverTab[185814]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:385
				return dec.tokenState == tokenObjectKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:385
				// _ = "end of CoverTab[185814]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:385
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:385
				_go_fuzz_dep_.CoverTab[185815]++
														var x string
														old := dec.tokenState
														dec.tokenState = tokenTopValue
														err := dec.Decode(&x)
														dec.tokenState = old
														if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:391
					_go_fuzz_dep_.CoverTab[185817]++
															clearOffset(err)
															return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:393
					// _ = "end of CoverTab[185817]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:394
					_go_fuzz_dep_.CoverTab[185818]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:394
					// _ = "end of CoverTab[185818]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:394
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:394
				// _ = "end of CoverTab[185815]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:394
				_go_fuzz_dep_.CoverTab[185816]++
														dec.tokenState = tokenObjectColon
														return x, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:396
				// _ = "end of CoverTab[185816]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:397
				_go_fuzz_dep_.CoverTab[185819]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:397
				// _ = "end of CoverTab[185819]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:397
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:397
			// _ = "end of CoverTab[185793]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:397
			_go_fuzz_dep_.CoverTab[185794]++
													fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:398
			// _ = "end of CoverTab[185794]"

		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:400
			_go_fuzz_dep_.CoverTab[185795]++
													if !dec.tokenValueAllowed() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:401
				_go_fuzz_dep_.CoverTab[185820]++
														return dec.tokenError(c)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:402
				// _ = "end of CoverTab[185820]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:403
				_go_fuzz_dep_.CoverTab[185821]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:403
				// _ = "end of CoverTab[185821]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:403
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:403
			// _ = "end of CoverTab[185795]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:403
			_go_fuzz_dep_.CoverTab[185796]++
													var x interface{}
													if err := dec.Decode(&x); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:405
				_go_fuzz_dep_.CoverTab[185822]++
														clearOffset(err)
														return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:407
				// _ = "end of CoverTab[185822]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:408
				_go_fuzz_dep_.CoverTab[185823]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:408
				// _ = "end of CoverTab[185823]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:408
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:408
			// _ = "end of CoverTab[185796]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:408
			_go_fuzz_dep_.CoverTab[185797]++
													return x, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:409
			// _ = "end of CoverTab[185797]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:410
		// _ = "end of CoverTab[185777]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:411
	// _ = "end of CoverTab[185775]"
}

func clearOffset(err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:414
	_go_fuzz_dep_.CoverTab[185824]++
											if s, ok := err.(*SyntaxError); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:415
		_go_fuzz_dep_.CoverTab[185825]++
												s.Offset = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:416
		// _ = "end of CoverTab[185825]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:417
		_go_fuzz_dep_.CoverTab[185826]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:417
		// _ = "end of CoverTab[185826]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:417
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:417
	// _ = "end of CoverTab[185824]"
}

func (dec *Decoder) tokenError(c byte) (Token, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:420
	_go_fuzz_dep_.CoverTab[185827]++
											var context string
											switch dec.tokenState {
	case tokenTopValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:423
		_go_fuzz_dep_.CoverTab[185829]++
												context = " looking for beginning of value"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:424
		// _ = "end of CoverTab[185829]"
	case tokenArrayStart, tokenArrayValue, tokenObjectValue:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:425
		_go_fuzz_dep_.CoverTab[185830]++
												context = " looking for beginning of value"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:426
		// _ = "end of CoverTab[185830]"
	case tokenArrayComma:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:427
		_go_fuzz_dep_.CoverTab[185831]++
												context = " after array element"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:428
		// _ = "end of CoverTab[185831]"
	case tokenObjectKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:429
		_go_fuzz_dep_.CoverTab[185832]++
												context = " looking for beginning of object key string"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:430
		// _ = "end of CoverTab[185832]"
	case tokenObjectColon:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:431
		_go_fuzz_dep_.CoverTab[185833]++
												context = " after object key"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:432
		// _ = "end of CoverTab[185833]"
	case tokenObjectComma:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:433
		_go_fuzz_dep_.CoverTab[185834]++
												context = " after object key:value pair"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:434
		// _ = "end of CoverTab[185834]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:434
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:434
		_go_fuzz_dep_.CoverTab[185835]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:434
		// _ = "end of CoverTab[185835]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:435
	// _ = "end of CoverTab[185827]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:435
	_go_fuzz_dep_.CoverTab[185828]++
											return nil, &SyntaxError{"invalid character " + quoteChar(c) + " " + context, 0}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:436
	// _ = "end of CoverTab[185828]"
}

// More reports whether there is another element in the
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:439
// current array or object being parsed.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:441
func (dec *Decoder) More() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:441
	_go_fuzz_dep_.CoverTab[185836]++
											c, err := dec.peek()
											return err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
		_go_fuzz_dep_.CoverTab[185837]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
		return c != ']'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
		// _ = "end of CoverTab[185837]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
		_go_fuzz_dep_.CoverTab[185838]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
		return c != '}'
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
		// _ = "end of CoverTab[185838]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:443
	// _ = "end of CoverTab[185836]"
}

func (dec *Decoder) peek() (byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:446
	_go_fuzz_dep_.CoverTab[185839]++
											var err error
											for {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:448
		_go_fuzz_dep_.CoverTab[185840]++
												for i := dec.scanp; i < len(dec.buf); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:449
			_go_fuzz_dep_.CoverTab[185843]++
													c := dec.buf[i]
													if isSpace(c) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:451
				_go_fuzz_dep_.CoverTab[185845]++
														continue
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:452
				// _ = "end of CoverTab[185845]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:453
				_go_fuzz_dep_.CoverTab[185846]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:453
				// _ = "end of CoverTab[185846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:453
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:453
			// _ = "end of CoverTab[185843]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:453
			_go_fuzz_dep_.CoverTab[185844]++
													dec.scanp = i
													return c, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:455
			// _ = "end of CoverTab[185844]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:456
		// _ = "end of CoverTab[185840]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:456
		_go_fuzz_dep_.CoverTab[185841]++

												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:458
			_go_fuzz_dep_.CoverTab[185847]++
													return 0, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:459
			// _ = "end of CoverTab[185847]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:460
			_go_fuzz_dep_.CoverTab[185848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:460
			// _ = "end of CoverTab[185848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:460
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:460
		// _ = "end of CoverTab[185841]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:460
		_go_fuzz_dep_.CoverTab[185842]++
												err = dec.refill()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:461
		// _ = "end of CoverTab[185842]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:462
	// _ = "end of CoverTab[185839]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:463
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/stream.go:463
var _ = _go_fuzz_dep_.CoverTab
