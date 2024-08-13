// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/stream.go:5
package json

//line /usr/local/go/src/encoding/json/stream.go:5
import (
//line /usr/local/go/src/encoding/json/stream.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/stream.go:5
)
//line /usr/local/go/src/encoding/json/stream.go:5
import (
//line /usr/local/go/src/encoding/json/stream.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/stream.go:5
)

import (
	"bytes"
	"errors"
	"io"
)

// A Decoder reads and decodes JSON values from an input stream.
type Decoder struct {
	r	io.Reader
	buf	[]byte
	d	decodeState
	scanp	int	// start of unread data in buf
	scanned	int64	// amount of data already scanned
	scan	scanner
	err	error

	tokenState	int
	tokenStack	[]int
}

// NewDecoder returns a new decoder that reads from r.
//line /usr/local/go/src/encoding/json/stream.go:27
//
//line /usr/local/go/src/encoding/json/stream.go:27
// The decoder introduces its own buffering and may
//line /usr/local/go/src/encoding/json/stream.go:27
// read data from r beyond the JSON values requested.
//line /usr/local/go/src/encoding/json/stream.go:31
func NewDecoder(r io.Reader) *Decoder {
//line /usr/local/go/src/encoding/json/stream.go:31
	_go_fuzz_dep_.CoverTab[28378]++
							return &Decoder{r: r}
//line /usr/local/go/src/encoding/json/stream.go:32
	// _ = "end of CoverTab[28378]"
}

// UseNumber causes the Decoder to unmarshal a number into an interface{} as a
//line /usr/local/go/src/encoding/json/stream.go:35
// Number instead of as a float64.
//line /usr/local/go/src/encoding/json/stream.go:37
func (dec *Decoder) UseNumber() {
//line /usr/local/go/src/encoding/json/stream.go:37
	_go_fuzz_dep_.CoverTab[28379]++
//line /usr/local/go/src/encoding/json/stream.go:37
	dec.d.useNumber = true
//line /usr/local/go/src/encoding/json/stream.go:37
	// _ = "end of CoverTab[28379]"
//line /usr/local/go/src/encoding/json/stream.go:37
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
//line /usr/local/go/src/encoding/json/stream.go:39
// is a struct and the input contains object keys which do not match any
//line /usr/local/go/src/encoding/json/stream.go:39
// non-ignored, exported fields in the destination.
//line /usr/local/go/src/encoding/json/stream.go:42
func (dec *Decoder) DisallowUnknownFields() {
//line /usr/local/go/src/encoding/json/stream.go:42
	_go_fuzz_dep_.CoverTab[28380]++
//line /usr/local/go/src/encoding/json/stream.go:42
	dec.d.disallowUnknownFields = true
//line /usr/local/go/src/encoding/json/stream.go:42
	// _ = "end of CoverTab[28380]"
//line /usr/local/go/src/encoding/json/stream.go:42
}

// Decode reads the next JSON-encoded value from its
//line /usr/local/go/src/encoding/json/stream.go:44
// input and stores it in the value pointed to by v.
//line /usr/local/go/src/encoding/json/stream.go:44
//
//line /usr/local/go/src/encoding/json/stream.go:44
// See the documentation for Unmarshal for details about
//line /usr/local/go/src/encoding/json/stream.go:44
// the conversion of JSON into a Go value.
//line /usr/local/go/src/encoding/json/stream.go:49
func (dec *Decoder) Decode(v any) error {
//line /usr/local/go/src/encoding/json/stream.go:49
	_go_fuzz_dep_.CoverTab[28381]++
							if dec.err != nil {
//line /usr/local/go/src/encoding/json/stream.go:50
		_go_fuzz_dep_.CoverTab[28386]++
								return dec.err
//line /usr/local/go/src/encoding/json/stream.go:51
		// _ = "end of CoverTab[28386]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:52
		_go_fuzz_dep_.CoverTab[28387]++
//line /usr/local/go/src/encoding/json/stream.go:52
		// _ = "end of CoverTab[28387]"
//line /usr/local/go/src/encoding/json/stream.go:52
	}
//line /usr/local/go/src/encoding/json/stream.go:52
	// _ = "end of CoverTab[28381]"
//line /usr/local/go/src/encoding/json/stream.go:52
	_go_fuzz_dep_.CoverTab[28382]++

							if err := dec.tokenPrepareForDecode(); err != nil {
//line /usr/local/go/src/encoding/json/stream.go:54
		_go_fuzz_dep_.CoverTab[28388]++
								return err
//line /usr/local/go/src/encoding/json/stream.go:55
		// _ = "end of CoverTab[28388]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:56
		_go_fuzz_dep_.CoverTab[28389]++
//line /usr/local/go/src/encoding/json/stream.go:56
		// _ = "end of CoverTab[28389]"
//line /usr/local/go/src/encoding/json/stream.go:56
	}
//line /usr/local/go/src/encoding/json/stream.go:56
	// _ = "end of CoverTab[28382]"
//line /usr/local/go/src/encoding/json/stream.go:56
	_go_fuzz_dep_.CoverTab[28383]++

							if !dec.tokenValueAllowed() {
//line /usr/local/go/src/encoding/json/stream.go:58
		_go_fuzz_dep_.CoverTab[28390]++
								return &SyntaxError{msg: "not at beginning of value", Offset: dec.InputOffset()}
//line /usr/local/go/src/encoding/json/stream.go:59
		// _ = "end of CoverTab[28390]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:60
		_go_fuzz_dep_.CoverTab[28391]++
//line /usr/local/go/src/encoding/json/stream.go:60
		// _ = "end of CoverTab[28391]"
//line /usr/local/go/src/encoding/json/stream.go:60
	}
//line /usr/local/go/src/encoding/json/stream.go:60
	// _ = "end of CoverTab[28383]"
//line /usr/local/go/src/encoding/json/stream.go:60
	_go_fuzz_dep_.CoverTab[28384]++

//line /usr/local/go/src/encoding/json/stream.go:63
	n, err := dec.readValue()
	if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:64
		_go_fuzz_dep_.CoverTab[28392]++
								return err
//line /usr/local/go/src/encoding/json/stream.go:65
		// _ = "end of CoverTab[28392]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:66
		_go_fuzz_dep_.CoverTab[28393]++
//line /usr/local/go/src/encoding/json/stream.go:66
		// _ = "end of CoverTab[28393]"
//line /usr/local/go/src/encoding/json/stream.go:66
	}
//line /usr/local/go/src/encoding/json/stream.go:66
	// _ = "end of CoverTab[28384]"
//line /usr/local/go/src/encoding/json/stream.go:66
	_go_fuzz_dep_.CoverTab[28385]++
							dec.d.init(dec.buf[dec.scanp : dec.scanp+n])
							dec.scanp += n

//line /usr/local/go/src/encoding/json/stream.go:73
	err = dec.d.unmarshal(v)

//line /usr/local/go/src/encoding/json/stream.go:76
	dec.tokenValueEnd()

							return err
//line /usr/local/go/src/encoding/json/stream.go:78
	// _ = "end of CoverTab[28385]"
}

// Buffered returns a reader of the data remaining in the Decoder's
//line /usr/local/go/src/encoding/json/stream.go:81
// buffer. The reader is valid until the next call to Decode.
//line /usr/local/go/src/encoding/json/stream.go:83
func (dec *Decoder) Buffered() io.Reader {
//line /usr/local/go/src/encoding/json/stream.go:83
	_go_fuzz_dep_.CoverTab[28394]++
							return bytes.NewReader(dec.buf[dec.scanp:])
//line /usr/local/go/src/encoding/json/stream.go:84
	// _ = "end of CoverTab[28394]"
}

// readValue reads a JSON value into dec.buf.
//line /usr/local/go/src/encoding/json/stream.go:87
// It returns the length of the encoding.
//line /usr/local/go/src/encoding/json/stream.go:89
func (dec *Decoder) readValue() (int, error) {
//line /usr/local/go/src/encoding/json/stream.go:89
	_go_fuzz_dep_.CoverTab[28395]++
							dec.scan.reset()

							scanp := dec.scanp
							var err error
Input:

//line /usr/local/go/src/encoding/json/stream.go:97
	for scanp >= 0 {
//line /usr/local/go/src/encoding/json/stream.go:97
		_go_fuzz_dep_.CoverTab[28397]++

//line /usr/local/go/src/encoding/json/stream.go:100
		for ; scanp < len(dec.buf); scanp++ {
//line /usr/local/go/src/encoding/json/stream.go:100
			_go_fuzz_dep_.CoverTab[28400]++
									c := dec.buf[scanp]
									dec.scan.bytes++
									switch dec.scan.step(&dec.scan, c) {
			case scanEnd:
//line /usr/local/go/src/encoding/json/stream.go:104
				_go_fuzz_dep_.CoverTab[28401]++

//line /usr/local/go/src/encoding/json/stream.go:108
				dec.scan.bytes--
										break Input
//line /usr/local/go/src/encoding/json/stream.go:109
				// _ = "end of CoverTab[28401]"
			case scanEndObject, scanEndArray:
//line /usr/local/go/src/encoding/json/stream.go:110
				_go_fuzz_dep_.CoverTab[28402]++

//line /usr/local/go/src/encoding/json/stream.go:114
				if stateEndValue(&dec.scan, ' ') == scanEnd {
//line /usr/local/go/src/encoding/json/stream.go:114
					_go_fuzz_dep_.CoverTab[28405]++
											scanp++
											break Input
//line /usr/local/go/src/encoding/json/stream.go:116
					// _ = "end of CoverTab[28405]"
				} else {
//line /usr/local/go/src/encoding/json/stream.go:117
					_go_fuzz_dep_.CoverTab[28406]++
//line /usr/local/go/src/encoding/json/stream.go:117
					// _ = "end of CoverTab[28406]"
//line /usr/local/go/src/encoding/json/stream.go:117
				}
//line /usr/local/go/src/encoding/json/stream.go:117
				// _ = "end of CoverTab[28402]"
			case scanError:
//line /usr/local/go/src/encoding/json/stream.go:118
				_go_fuzz_dep_.CoverTab[28403]++
										dec.err = dec.scan.err
										return 0, dec.scan.err
//line /usr/local/go/src/encoding/json/stream.go:120
				// _ = "end of CoverTab[28403]"
//line /usr/local/go/src/encoding/json/stream.go:120
			default:
//line /usr/local/go/src/encoding/json/stream.go:120
				_go_fuzz_dep_.CoverTab[28404]++
//line /usr/local/go/src/encoding/json/stream.go:120
				// _ = "end of CoverTab[28404]"
			}
//line /usr/local/go/src/encoding/json/stream.go:121
			// _ = "end of CoverTab[28400]"
		}
//line /usr/local/go/src/encoding/json/stream.go:122
		// _ = "end of CoverTab[28397]"
//line /usr/local/go/src/encoding/json/stream.go:122
		_go_fuzz_dep_.CoverTab[28398]++

//line /usr/local/go/src/encoding/json/stream.go:126
		if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:126
			_go_fuzz_dep_.CoverTab[28407]++
									if err == io.EOF {
//line /usr/local/go/src/encoding/json/stream.go:127
				_go_fuzz_dep_.CoverTab[28409]++
										if dec.scan.step(&dec.scan, ' ') == scanEnd {
//line /usr/local/go/src/encoding/json/stream.go:128
					_go_fuzz_dep_.CoverTab[28411]++
											break Input
//line /usr/local/go/src/encoding/json/stream.go:129
					// _ = "end of CoverTab[28411]"
				} else {
//line /usr/local/go/src/encoding/json/stream.go:130
					_go_fuzz_dep_.CoverTab[28412]++
//line /usr/local/go/src/encoding/json/stream.go:130
					// _ = "end of CoverTab[28412]"
//line /usr/local/go/src/encoding/json/stream.go:130
				}
//line /usr/local/go/src/encoding/json/stream.go:130
				// _ = "end of CoverTab[28409]"
//line /usr/local/go/src/encoding/json/stream.go:130
				_go_fuzz_dep_.CoverTab[28410]++
										if nonSpace(dec.buf) {
//line /usr/local/go/src/encoding/json/stream.go:131
					_go_fuzz_dep_.CoverTab[28413]++
											err = io.ErrUnexpectedEOF
//line /usr/local/go/src/encoding/json/stream.go:132
					// _ = "end of CoverTab[28413]"
				} else {
//line /usr/local/go/src/encoding/json/stream.go:133
					_go_fuzz_dep_.CoverTab[28414]++
//line /usr/local/go/src/encoding/json/stream.go:133
					// _ = "end of CoverTab[28414]"
//line /usr/local/go/src/encoding/json/stream.go:133
				}
//line /usr/local/go/src/encoding/json/stream.go:133
				// _ = "end of CoverTab[28410]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:134
				_go_fuzz_dep_.CoverTab[28415]++
//line /usr/local/go/src/encoding/json/stream.go:134
				// _ = "end of CoverTab[28415]"
//line /usr/local/go/src/encoding/json/stream.go:134
			}
//line /usr/local/go/src/encoding/json/stream.go:134
			// _ = "end of CoverTab[28407]"
//line /usr/local/go/src/encoding/json/stream.go:134
			_go_fuzz_dep_.CoverTab[28408]++
									dec.err = err
									return 0, err
//line /usr/local/go/src/encoding/json/stream.go:136
			// _ = "end of CoverTab[28408]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:137
			_go_fuzz_dep_.CoverTab[28416]++
//line /usr/local/go/src/encoding/json/stream.go:137
			// _ = "end of CoverTab[28416]"
//line /usr/local/go/src/encoding/json/stream.go:137
		}
//line /usr/local/go/src/encoding/json/stream.go:137
		// _ = "end of CoverTab[28398]"
//line /usr/local/go/src/encoding/json/stream.go:137
		_go_fuzz_dep_.CoverTab[28399]++

								n := scanp - dec.scanp
								err = dec.refill()
								scanp = dec.scanp + n
//line /usr/local/go/src/encoding/json/stream.go:141
		// _ = "end of CoverTab[28399]"
	}
//line /usr/local/go/src/encoding/json/stream.go:142
	// _ = "end of CoverTab[28395]"
//line /usr/local/go/src/encoding/json/stream.go:142
	_go_fuzz_dep_.CoverTab[28396]++
							return scanp - dec.scanp, nil
//line /usr/local/go/src/encoding/json/stream.go:143
	// _ = "end of CoverTab[28396]"
}

func (dec *Decoder) refill() error {
//line /usr/local/go/src/encoding/json/stream.go:146
	_go_fuzz_dep_.CoverTab[28417]++

//line /usr/local/go/src/encoding/json/stream.go:149
	if dec.scanp > 0 {
//line /usr/local/go/src/encoding/json/stream.go:149
		_go_fuzz_dep_.CoverTab[28420]++
								dec.scanned += int64(dec.scanp)
								n := copy(dec.buf, dec.buf[dec.scanp:])
								dec.buf = dec.buf[:n]
								dec.scanp = 0
//line /usr/local/go/src/encoding/json/stream.go:153
		// _ = "end of CoverTab[28420]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:154
		_go_fuzz_dep_.CoverTab[28421]++
//line /usr/local/go/src/encoding/json/stream.go:154
		// _ = "end of CoverTab[28421]"
//line /usr/local/go/src/encoding/json/stream.go:154
	}
//line /usr/local/go/src/encoding/json/stream.go:154
	// _ = "end of CoverTab[28417]"
//line /usr/local/go/src/encoding/json/stream.go:154
	_go_fuzz_dep_.CoverTab[28418]++

	// Grow buffer if not large enough.
	const minRead = 512
	if cap(dec.buf)-len(dec.buf) < minRead {
//line /usr/local/go/src/encoding/json/stream.go:158
		_go_fuzz_dep_.CoverTab[28422]++
								newBuf := make([]byte, len(dec.buf), 2*cap(dec.buf)+minRead)
								copy(newBuf, dec.buf)
								dec.buf = newBuf
//line /usr/local/go/src/encoding/json/stream.go:161
		// _ = "end of CoverTab[28422]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:162
		_go_fuzz_dep_.CoverTab[28423]++
//line /usr/local/go/src/encoding/json/stream.go:162
		// _ = "end of CoverTab[28423]"
//line /usr/local/go/src/encoding/json/stream.go:162
	}
//line /usr/local/go/src/encoding/json/stream.go:162
	// _ = "end of CoverTab[28418]"
//line /usr/local/go/src/encoding/json/stream.go:162
	_go_fuzz_dep_.CoverTab[28419]++

//line /usr/local/go/src/encoding/json/stream.go:165
	n, err := dec.r.Read(dec.buf[len(dec.buf):cap(dec.buf)])
							dec.buf = dec.buf[0 : len(dec.buf)+n]

							return err
//line /usr/local/go/src/encoding/json/stream.go:168
	// _ = "end of CoverTab[28419]"
}

func nonSpace(b []byte) bool {
//line /usr/local/go/src/encoding/json/stream.go:171
	_go_fuzz_dep_.CoverTab[28424]++
							for _, c := range b {
//line /usr/local/go/src/encoding/json/stream.go:172
		_go_fuzz_dep_.CoverTab[28426]++
								if !isSpace(c) {
//line /usr/local/go/src/encoding/json/stream.go:173
			_go_fuzz_dep_.CoverTab[28427]++
									return true
//line /usr/local/go/src/encoding/json/stream.go:174
			// _ = "end of CoverTab[28427]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:175
			_go_fuzz_dep_.CoverTab[28428]++
//line /usr/local/go/src/encoding/json/stream.go:175
			// _ = "end of CoverTab[28428]"
//line /usr/local/go/src/encoding/json/stream.go:175
		}
//line /usr/local/go/src/encoding/json/stream.go:175
		// _ = "end of CoverTab[28426]"
	}
//line /usr/local/go/src/encoding/json/stream.go:176
	// _ = "end of CoverTab[28424]"
//line /usr/local/go/src/encoding/json/stream.go:176
	_go_fuzz_dep_.CoverTab[28425]++
							return false
//line /usr/local/go/src/encoding/json/stream.go:177
	// _ = "end of CoverTab[28425]"
}

// An Encoder writes JSON values to an output stream.
type Encoder struct {
	w		io.Writer
	err		error
	escapeHTML	bool

	indentBuf	*bytes.Buffer
	indentPrefix	string
	indentValue	string
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
//line /usr/local/go/src/encoding/json/stream.go:192
	_go_fuzz_dep_.CoverTab[28429]++
							return &Encoder{w: w, escapeHTML: true}
//line /usr/local/go/src/encoding/json/stream.go:193
	// _ = "end of CoverTab[28429]"
}

// Encode writes the JSON encoding of v to the stream,
//line /usr/local/go/src/encoding/json/stream.go:196
// followed by a newline character.
//line /usr/local/go/src/encoding/json/stream.go:196
//
//line /usr/local/go/src/encoding/json/stream.go:196
// See the documentation for Marshal for details about the
//line /usr/local/go/src/encoding/json/stream.go:196
// conversion of Go values to JSON.
//line /usr/local/go/src/encoding/json/stream.go:201
func (enc *Encoder) Encode(v any) error {
//line /usr/local/go/src/encoding/json/stream.go:201
	_go_fuzz_dep_.CoverTab[28430]++
							if enc.err != nil {
//line /usr/local/go/src/encoding/json/stream.go:202
		_go_fuzz_dep_.CoverTab[28435]++
								return enc.err
//line /usr/local/go/src/encoding/json/stream.go:203
		// _ = "end of CoverTab[28435]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:204
		_go_fuzz_dep_.CoverTab[28436]++
//line /usr/local/go/src/encoding/json/stream.go:204
		// _ = "end of CoverTab[28436]"
//line /usr/local/go/src/encoding/json/stream.go:204
	}
//line /usr/local/go/src/encoding/json/stream.go:204
	// _ = "end of CoverTab[28430]"
//line /usr/local/go/src/encoding/json/stream.go:204
	_go_fuzz_dep_.CoverTab[28431]++

							e := newEncodeState()
							defer encodeStatePool.Put(e)

							err := e.marshal(v, encOpts{escapeHTML: enc.escapeHTML})
							if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:210
		_go_fuzz_dep_.CoverTab[28437]++
								return err
//line /usr/local/go/src/encoding/json/stream.go:211
		// _ = "end of CoverTab[28437]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:212
		_go_fuzz_dep_.CoverTab[28438]++
//line /usr/local/go/src/encoding/json/stream.go:212
		// _ = "end of CoverTab[28438]"
//line /usr/local/go/src/encoding/json/stream.go:212
	}
//line /usr/local/go/src/encoding/json/stream.go:212
	// _ = "end of CoverTab[28431]"
//line /usr/local/go/src/encoding/json/stream.go:212
	_go_fuzz_dep_.CoverTab[28432]++

//line /usr/local/go/src/encoding/json/stream.go:220
	e.WriteByte('\n')

	b := e.Bytes()
	if enc.indentPrefix != "" || func() bool {
//line /usr/local/go/src/encoding/json/stream.go:223
		_go_fuzz_dep_.CoverTab[28439]++
//line /usr/local/go/src/encoding/json/stream.go:223
		return enc.indentValue != ""
//line /usr/local/go/src/encoding/json/stream.go:223
		// _ = "end of CoverTab[28439]"
//line /usr/local/go/src/encoding/json/stream.go:223
	}() {
//line /usr/local/go/src/encoding/json/stream.go:223
		_go_fuzz_dep_.CoverTab[28440]++
								if enc.indentBuf == nil {
//line /usr/local/go/src/encoding/json/stream.go:224
			_go_fuzz_dep_.CoverTab[28443]++
									enc.indentBuf = new(bytes.Buffer)
//line /usr/local/go/src/encoding/json/stream.go:225
			// _ = "end of CoverTab[28443]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:226
			_go_fuzz_dep_.CoverTab[28444]++
//line /usr/local/go/src/encoding/json/stream.go:226
			// _ = "end of CoverTab[28444]"
//line /usr/local/go/src/encoding/json/stream.go:226
		}
//line /usr/local/go/src/encoding/json/stream.go:226
		// _ = "end of CoverTab[28440]"
//line /usr/local/go/src/encoding/json/stream.go:226
		_go_fuzz_dep_.CoverTab[28441]++
								enc.indentBuf.Reset()
								err = Indent(enc.indentBuf, b, enc.indentPrefix, enc.indentValue)
								if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:229
			_go_fuzz_dep_.CoverTab[28445]++
									return err
//line /usr/local/go/src/encoding/json/stream.go:230
			// _ = "end of CoverTab[28445]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:231
			_go_fuzz_dep_.CoverTab[28446]++
//line /usr/local/go/src/encoding/json/stream.go:231
			// _ = "end of CoverTab[28446]"
//line /usr/local/go/src/encoding/json/stream.go:231
		}
//line /usr/local/go/src/encoding/json/stream.go:231
		// _ = "end of CoverTab[28441]"
//line /usr/local/go/src/encoding/json/stream.go:231
		_go_fuzz_dep_.CoverTab[28442]++
								b = enc.indentBuf.Bytes()
//line /usr/local/go/src/encoding/json/stream.go:232
		// _ = "end of CoverTab[28442]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:233
		_go_fuzz_dep_.CoverTab[28447]++
//line /usr/local/go/src/encoding/json/stream.go:233
		// _ = "end of CoverTab[28447]"
//line /usr/local/go/src/encoding/json/stream.go:233
	}
//line /usr/local/go/src/encoding/json/stream.go:233
	// _ = "end of CoverTab[28432]"
//line /usr/local/go/src/encoding/json/stream.go:233
	_go_fuzz_dep_.CoverTab[28433]++
							if _, err = enc.w.Write(b); err != nil {
//line /usr/local/go/src/encoding/json/stream.go:234
		_go_fuzz_dep_.CoverTab[28448]++
								enc.err = err
//line /usr/local/go/src/encoding/json/stream.go:235
		// _ = "end of CoverTab[28448]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:236
		_go_fuzz_dep_.CoverTab[28449]++
//line /usr/local/go/src/encoding/json/stream.go:236
		// _ = "end of CoverTab[28449]"
//line /usr/local/go/src/encoding/json/stream.go:236
	}
//line /usr/local/go/src/encoding/json/stream.go:236
	// _ = "end of CoverTab[28433]"
//line /usr/local/go/src/encoding/json/stream.go:236
	_go_fuzz_dep_.CoverTab[28434]++
							return err
//line /usr/local/go/src/encoding/json/stream.go:237
	// _ = "end of CoverTab[28434]"
}

// SetIndent instructs the encoder to format each subsequent encoded
//line /usr/local/go/src/encoding/json/stream.go:240
// value as if indented by the package-level function Indent(dst, src, prefix, indent).
//line /usr/local/go/src/encoding/json/stream.go:240
// Calling SetIndent("", "") disables indentation.
//line /usr/local/go/src/encoding/json/stream.go:243
func (enc *Encoder) SetIndent(prefix, indent string) {
//line /usr/local/go/src/encoding/json/stream.go:243
	_go_fuzz_dep_.CoverTab[28450]++
							enc.indentPrefix = prefix
							enc.indentValue = indent
//line /usr/local/go/src/encoding/json/stream.go:245
	// _ = "end of CoverTab[28450]"
}

// SetEscapeHTML specifies whether problematic HTML characters
//line /usr/local/go/src/encoding/json/stream.go:248
// should be escaped inside JSON quoted strings.
//line /usr/local/go/src/encoding/json/stream.go:248
// The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
//line /usr/local/go/src/encoding/json/stream.go:248
// to avoid certain safety problems that can arise when embedding JSON in HTML.
//line /usr/local/go/src/encoding/json/stream.go:248
//
//line /usr/local/go/src/encoding/json/stream.go:248
// In non-HTML settings where the escaping interferes with the readability
//line /usr/local/go/src/encoding/json/stream.go:248
// of the output, SetEscapeHTML(false) disables this behavior.
//line /usr/local/go/src/encoding/json/stream.go:255
func (enc *Encoder) SetEscapeHTML(on bool) {
//line /usr/local/go/src/encoding/json/stream.go:255
	_go_fuzz_dep_.CoverTab[28451]++
							enc.escapeHTML = on
//line /usr/local/go/src/encoding/json/stream.go:256
	// _ = "end of CoverTab[28451]"
}

// RawMessage is a raw encoded JSON value.
//line /usr/local/go/src/encoding/json/stream.go:259
// It implements Marshaler and Unmarshaler and can
//line /usr/local/go/src/encoding/json/stream.go:259
// be used to delay JSON decoding or precompute a JSON encoding.
//line /usr/local/go/src/encoding/json/stream.go:262
type RawMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
//line /usr/local/go/src/encoding/json/stream.go:265
	_go_fuzz_dep_.CoverTab[28452]++
							if m == nil {
//line /usr/local/go/src/encoding/json/stream.go:266
		_go_fuzz_dep_.CoverTab[28454]++
								return []byte("null"), nil
//line /usr/local/go/src/encoding/json/stream.go:267
		// _ = "end of CoverTab[28454]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:268
		_go_fuzz_dep_.CoverTab[28455]++
//line /usr/local/go/src/encoding/json/stream.go:268
		// _ = "end of CoverTab[28455]"
//line /usr/local/go/src/encoding/json/stream.go:268
	}
//line /usr/local/go/src/encoding/json/stream.go:268
	// _ = "end of CoverTab[28452]"
//line /usr/local/go/src/encoding/json/stream.go:268
	_go_fuzz_dep_.CoverTab[28453]++
							return m, nil
//line /usr/local/go/src/encoding/json/stream.go:269
	// _ = "end of CoverTab[28453]"
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
//line /usr/local/go/src/encoding/json/stream.go:273
	_go_fuzz_dep_.CoverTab[28456]++
							if m == nil {
//line /usr/local/go/src/encoding/json/stream.go:274
		_go_fuzz_dep_.CoverTab[28458]++
								return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
//line /usr/local/go/src/encoding/json/stream.go:275
		// _ = "end of CoverTab[28458]"
	} else {
//line /usr/local/go/src/encoding/json/stream.go:276
		_go_fuzz_dep_.CoverTab[28459]++
//line /usr/local/go/src/encoding/json/stream.go:276
		// _ = "end of CoverTab[28459]"
//line /usr/local/go/src/encoding/json/stream.go:276
	}
//line /usr/local/go/src/encoding/json/stream.go:276
	// _ = "end of CoverTab[28456]"
//line /usr/local/go/src/encoding/json/stream.go:276
	_go_fuzz_dep_.CoverTab[28457]++
							*m = append((*m)[0:0], data...)
							return nil
//line /usr/local/go/src/encoding/json/stream.go:278
	// _ = "end of CoverTab[28457]"
}

var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)

// A Token holds a value of one of these types:
//line /usr/local/go/src/encoding/json/stream.go:284
//
//line /usr/local/go/src/encoding/json/stream.go:284
//	Delim, for the four JSON delimiters [ ] { }
//line /usr/local/go/src/encoding/json/stream.go:284
//	bool, for JSON booleans
//line /usr/local/go/src/encoding/json/stream.go:284
//	float64, for JSON numbers
//line /usr/local/go/src/encoding/json/stream.go:284
//	Number, for JSON numbers
//line /usr/local/go/src/encoding/json/stream.go:284
//	string, for JSON string literals
//line /usr/local/go/src/encoding/json/stream.go:284
//	nil, for JSON null
//line /usr/local/go/src/encoding/json/stream.go:292
type Token any

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
//line /usr/local/go/src/encoding/json/stream.go:307
	_go_fuzz_dep_.CoverTab[28460]++

//line /usr/local/go/src/encoding/json/stream.go:311
	switch dec.tokenState {
	case tokenArrayComma:
//line /usr/local/go/src/encoding/json/stream.go:312
		_go_fuzz_dep_.CoverTab[28462]++
								c, err := dec.peek()
								if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:314
			_go_fuzz_dep_.CoverTab[28469]++
									return err
//line /usr/local/go/src/encoding/json/stream.go:315
			// _ = "end of CoverTab[28469]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:316
			_go_fuzz_dep_.CoverTab[28470]++
//line /usr/local/go/src/encoding/json/stream.go:316
			// _ = "end of CoverTab[28470]"
//line /usr/local/go/src/encoding/json/stream.go:316
		}
//line /usr/local/go/src/encoding/json/stream.go:316
		// _ = "end of CoverTab[28462]"
//line /usr/local/go/src/encoding/json/stream.go:316
		_go_fuzz_dep_.CoverTab[28463]++
								if c != ',' {
//line /usr/local/go/src/encoding/json/stream.go:317
			_go_fuzz_dep_.CoverTab[28471]++
									return &SyntaxError{"expected comma after array element", dec.InputOffset()}
//line /usr/local/go/src/encoding/json/stream.go:318
			// _ = "end of CoverTab[28471]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:319
			_go_fuzz_dep_.CoverTab[28472]++
//line /usr/local/go/src/encoding/json/stream.go:319
			// _ = "end of CoverTab[28472]"
//line /usr/local/go/src/encoding/json/stream.go:319
		}
//line /usr/local/go/src/encoding/json/stream.go:319
		// _ = "end of CoverTab[28463]"
//line /usr/local/go/src/encoding/json/stream.go:319
		_go_fuzz_dep_.CoverTab[28464]++
								dec.scanp++
								dec.tokenState = tokenArrayValue
//line /usr/local/go/src/encoding/json/stream.go:321
		// _ = "end of CoverTab[28464]"
	case tokenObjectColon:
//line /usr/local/go/src/encoding/json/stream.go:322
		_go_fuzz_dep_.CoverTab[28465]++
								c, err := dec.peek()
								if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:324
			_go_fuzz_dep_.CoverTab[28473]++
									return err
//line /usr/local/go/src/encoding/json/stream.go:325
			// _ = "end of CoverTab[28473]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:326
			_go_fuzz_dep_.CoverTab[28474]++
//line /usr/local/go/src/encoding/json/stream.go:326
			// _ = "end of CoverTab[28474]"
//line /usr/local/go/src/encoding/json/stream.go:326
		}
//line /usr/local/go/src/encoding/json/stream.go:326
		// _ = "end of CoverTab[28465]"
//line /usr/local/go/src/encoding/json/stream.go:326
		_go_fuzz_dep_.CoverTab[28466]++
								if c != ':' {
//line /usr/local/go/src/encoding/json/stream.go:327
			_go_fuzz_dep_.CoverTab[28475]++
									return &SyntaxError{"expected colon after object key", dec.InputOffset()}
//line /usr/local/go/src/encoding/json/stream.go:328
			// _ = "end of CoverTab[28475]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:329
			_go_fuzz_dep_.CoverTab[28476]++
//line /usr/local/go/src/encoding/json/stream.go:329
			// _ = "end of CoverTab[28476]"
//line /usr/local/go/src/encoding/json/stream.go:329
		}
//line /usr/local/go/src/encoding/json/stream.go:329
		// _ = "end of CoverTab[28466]"
//line /usr/local/go/src/encoding/json/stream.go:329
		_go_fuzz_dep_.CoverTab[28467]++
								dec.scanp++
								dec.tokenState = tokenObjectValue
//line /usr/local/go/src/encoding/json/stream.go:331
		// _ = "end of CoverTab[28467]"
//line /usr/local/go/src/encoding/json/stream.go:331
	default:
//line /usr/local/go/src/encoding/json/stream.go:331
		_go_fuzz_dep_.CoverTab[28468]++
//line /usr/local/go/src/encoding/json/stream.go:331
		// _ = "end of CoverTab[28468]"
	}
//line /usr/local/go/src/encoding/json/stream.go:332
	// _ = "end of CoverTab[28460]"
//line /usr/local/go/src/encoding/json/stream.go:332
	_go_fuzz_dep_.CoverTab[28461]++
							return nil
//line /usr/local/go/src/encoding/json/stream.go:333
	// _ = "end of CoverTab[28461]"
}

func (dec *Decoder) tokenValueAllowed() bool {
//line /usr/local/go/src/encoding/json/stream.go:336
	_go_fuzz_dep_.CoverTab[28477]++
							switch dec.tokenState {
	case tokenTopValue, tokenArrayStart, tokenArrayValue, tokenObjectValue:
//line /usr/local/go/src/encoding/json/stream.go:338
		_go_fuzz_dep_.CoverTab[28479]++
								return true
//line /usr/local/go/src/encoding/json/stream.go:339
		// _ = "end of CoverTab[28479]"
//line /usr/local/go/src/encoding/json/stream.go:339
	default:
//line /usr/local/go/src/encoding/json/stream.go:339
		_go_fuzz_dep_.CoverTab[28480]++
//line /usr/local/go/src/encoding/json/stream.go:339
		// _ = "end of CoverTab[28480]"
	}
//line /usr/local/go/src/encoding/json/stream.go:340
	// _ = "end of CoverTab[28477]"
//line /usr/local/go/src/encoding/json/stream.go:340
	_go_fuzz_dep_.CoverTab[28478]++
							return false
//line /usr/local/go/src/encoding/json/stream.go:341
	// _ = "end of CoverTab[28478]"
}

func (dec *Decoder) tokenValueEnd() {
//line /usr/local/go/src/encoding/json/stream.go:344
	_go_fuzz_dep_.CoverTab[28481]++
							switch dec.tokenState {
	case tokenArrayStart, tokenArrayValue:
//line /usr/local/go/src/encoding/json/stream.go:346
		_go_fuzz_dep_.CoverTab[28482]++
								dec.tokenState = tokenArrayComma
//line /usr/local/go/src/encoding/json/stream.go:347
		// _ = "end of CoverTab[28482]"
	case tokenObjectValue:
//line /usr/local/go/src/encoding/json/stream.go:348
		_go_fuzz_dep_.CoverTab[28483]++
								dec.tokenState = tokenObjectComma
//line /usr/local/go/src/encoding/json/stream.go:349
		// _ = "end of CoverTab[28483]"
//line /usr/local/go/src/encoding/json/stream.go:349
	default:
//line /usr/local/go/src/encoding/json/stream.go:349
		_go_fuzz_dep_.CoverTab[28484]++
//line /usr/local/go/src/encoding/json/stream.go:349
		// _ = "end of CoverTab[28484]"
	}
//line /usr/local/go/src/encoding/json/stream.go:350
	// _ = "end of CoverTab[28481]"
}

// A Delim is a JSON array or object delimiter, one of [ ] { or }.
type Delim rune

func (d Delim) String() string {
//line /usr/local/go/src/encoding/json/stream.go:356
	_go_fuzz_dep_.CoverTab[28485]++
							return string(d)
//line /usr/local/go/src/encoding/json/stream.go:357
	// _ = "end of CoverTab[28485]"
}

// Token returns the next JSON token in the input stream.
//line /usr/local/go/src/encoding/json/stream.go:360
// At the end of the input stream, Token returns nil, io.EOF.
//line /usr/local/go/src/encoding/json/stream.go:360
//
//line /usr/local/go/src/encoding/json/stream.go:360
// Token guarantees that the delimiters [ ] { } it returns are
//line /usr/local/go/src/encoding/json/stream.go:360
// properly nested and matched: if Token encounters an unexpected
//line /usr/local/go/src/encoding/json/stream.go:360
// delimiter in the input, it will return an error.
//line /usr/local/go/src/encoding/json/stream.go:360
//
//line /usr/local/go/src/encoding/json/stream.go:360
// The input stream consists of basic JSON values—bool, string,
//line /usr/local/go/src/encoding/json/stream.go:360
// number, and null—along with delimiters [ ] { } of type Delim
//line /usr/local/go/src/encoding/json/stream.go:360
// to mark the start and end of arrays and objects.
//line /usr/local/go/src/encoding/json/stream.go:360
// Commas and colons are elided.
//line /usr/local/go/src/encoding/json/stream.go:371
func (dec *Decoder) Token() (Token, error) {
//line /usr/local/go/src/encoding/json/stream.go:371
	_go_fuzz_dep_.CoverTab[28486]++
							for {
//line /usr/local/go/src/encoding/json/stream.go:372
		_go_fuzz_dep_.CoverTab[28487]++
								c, err := dec.peek()
								if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:374
			_go_fuzz_dep_.CoverTab[28489]++
									return nil, err
//line /usr/local/go/src/encoding/json/stream.go:375
			// _ = "end of CoverTab[28489]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:376
			_go_fuzz_dep_.CoverTab[28490]++
//line /usr/local/go/src/encoding/json/stream.go:376
			// _ = "end of CoverTab[28490]"
//line /usr/local/go/src/encoding/json/stream.go:376
		}
//line /usr/local/go/src/encoding/json/stream.go:376
		// _ = "end of CoverTab[28487]"
//line /usr/local/go/src/encoding/json/stream.go:376
		_go_fuzz_dep_.CoverTab[28488]++
								switch c {
		case '[':
//line /usr/local/go/src/encoding/json/stream.go:378
			_go_fuzz_dep_.CoverTab[28491]++
									if !dec.tokenValueAllowed() {
//line /usr/local/go/src/encoding/json/stream.go:379
				_go_fuzz_dep_.CoverTab[28509]++
										return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:380
				// _ = "end of CoverTab[28509]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:381
				_go_fuzz_dep_.CoverTab[28510]++
//line /usr/local/go/src/encoding/json/stream.go:381
				// _ = "end of CoverTab[28510]"
//line /usr/local/go/src/encoding/json/stream.go:381
			}
//line /usr/local/go/src/encoding/json/stream.go:381
			// _ = "end of CoverTab[28491]"
//line /usr/local/go/src/encoding/json/stream.go:381
			_go_fuzz_dep_.CoverTab[28492]++
									dec.scanp++
									dec.tokenStack = append(dec.tokenStack, dec.tokenState)
									dec.tokenState = tokenArrayStart
									return Delim('['), nil
//line /usr/local/go/src/encoding/json/stream.go:385
			// _ = "end of CoverTab[28492]"

		case ']':
//line /usr/local/go/src/encoding/json/stream.go:387
			_go_fuzz_dep_.CoverTab[28493]++
									if dec.tokenState != tokenArrayStart && func() bool {
//line /usr/local/go/src/encoding/json/stream.go:388
				_go_fuzz_dep_.CoverTab[28511]++
//line /usr/local/go/src/encoding/json/stream.go:388
				return dec.tokenState != tokenArrayComma
//line /usr/local/go/src/encoding/json/stream.go:388
				// _ = "end of CoverTab[28511]"
//line /usr/local/go/src/encoding/json/stream.go:388
			}() {
//line /usr/local/go/src/encoding/json/stream.go:388
				_go_fuzz_dep_.CoverTab[28512]++
										return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:389
				// _ = "end of CoverTab[28512]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:390
				_go_fuzz_dep_.CoverTab[28513]++
//line /usr/local/go/src/encoding/json/stream.go:390
				// _ = "end of CoverTab[28513]"
//line /usr/local/go/src/encoding/json/stream.go:390
			}
//line /usr/local/go/src/encoding/json/stream.go:390
			// _ = "end of CoverTab[28493]"
//line /usr/local/go/src/encoding/json/stream.go:390
			_go_fuzz_dep_.CoverTab[28494]++
									dec.scanp++
									dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
									dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
									dec.tokenValueEnd()
									return Delim(']'), nil
//line /usr/local/go/src/encoding/json/stream.go:395
			// _ = "end of CoverTab[28494]"

		case '{':
//line /usr/local/go/src/encoding/json/stream.go:397
			_go_fuzz_dep_.CoverTab[28495]++
									if !dec.tokenValueAllowed() {
//line /usr/local/go/src/encoding/json/stream.go:398
				_go_fuzz_dep_.CoverTab[28514]++
										return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:399
				// _ = "end of CoverTab[28514]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:400
				_go_fuzz_dep_.CoverTab[28515]++
//line /usr/local/go/src/encoding/json/stream.go:400
				// _ = "end of CoverTab[28515]"
//line /usr/local/go/src/encoding/json/stream.go:400
			}
//line /usr/local/go/src/encoding/json/stream.go:400
			// _ = "end of CoverTab[28495]"
//line /usr/local/go/src/encoding/json/stream.go:400
			_go_fuzz_dep_.CoverTab[28496]++
									dec.scanp++
									dec.tokenStack = append(dec.tokenStack, dec.tokenState)
									dec.tokenState = tokenObjectStart
									return Delim('{'), nil
//line /usr/local/go/src/encoding/json/stream.go:404
			// _ = "end of CoverTab[28496]"

		case '}':
//line /usr/local/go/src/encoding/json/stream.go:406
			_go_fuzz_dep_.CoverTab[28497]++
									if dec.tokenState != tokenObjectStart && func() bool {
//line /usr/local/go/src/encoding/json/stream.go:407
				_go_fuzz_dep_.CoverTab[28516]++
//line /usr/local/go/src/encoding/json/stream.go:407
				return dec.tokenState != tokenObjectComma
//line /usr/local/go/src/encoding/json/stream.go:407
				// _ = "end of CoverTab[28516]"
//line /usr/local/go/src/encoding/json/stream.go:407
			}() {
//line /usr/local/go/src/encoding/json/stream.go:407
				_go_fuzz_dep_.CoverTab[28517]++
										return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:408
				// _ = "end of CoverTab[28517]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:409
				_go_fuzz_dep_.CoverTab[28518]++
//line /usr/local/go/src/encoding/json/stream.go:409
				// _ = "end of CoverTab[28518]"
//line /usr/local/go/src/encoding/json/stream.go:409
			}
//line /usr/local/go/src/encoding/json/stream.go:409
			// _ = "end of CoverTab[28497]"
//line /usr/local/go/src/encoding/json/stream.go:409
			_go_fuzz_dep_.CoverTab[28498]++
									dec.scanp++
									dec.tokenState = dec.tokenStack[len(dec.tokenStack)-1]
									dec.tokenStack = dec.tokenStack[:len(dec.tokenStack)-1]
									dec.tokenValueEnd()
									return Delim('}'), nil
//line /usr/local/go/src/encoding/json/stream.go:414
			// _ = "end of CoverTab[28498]"

		case ':':
//line /usr/local/go/src/encoding/json/stream.go:416
			_go_fuzz_dep_.CoverTab[28499]++
									if dec.tokenState != tokenObjectColon {
//line /usr/local/go/src/encoding/json/stream.go:417
				_go_fuzz_dep_.CoverTab[28519]++
										return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:418
				// _ = "end of CoverTab[28519]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:419
				_go_fuzz_dep_.CoverTab[28520]++
//line /usr/local/go/src/encoding/json/stream.go:419
				// _ = "end of CoverTab[28520]"
//line /usr/local/go/src/encoding/json/stream.go:419
			}
//line /usr/local/go/src/encoding/json/stream.go:419
			// _ = "end of CoverTab[28499]"
//line /usr/local/go/src/encoding/json/stream.go:419
			_go_fuzz_dep_.CoverTab[28500]++
									dec.scanp++
									dec.tokenState = tokenObjectValue
									continue
//line /usr/local/go/src/encoding/json/stream.go:422
			// _ = "end of CoverTab[28500]"

		case ',':
//line /usr/local/go/src/encoding/json/stream.go:424
			_go_fuzz_dep_.CoverTab[28501]++
									if dec.tokenState == tokenArrayComma {
//line /usr/local/go/src/encoding/json/stream.go:425
				_go_fuzz_dep_.CoverTab[28521]++
										dec.scanp++
										dec.tokenState = tokenArrayValue
										continue
//line /usr/local/go/src/encoding/json/stream.go:428
				// _ = "end of CoverTab[28521]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:429
				_go_fuzz_dep_.CoverTab[28522]++
//line /usr/local/go/src/encoding/json/stream.go:429
				// _ = "end of CoverTab[28522]"
//line /usr/local/go/src/encoding/json/stream.go:429
			}
//line /usr/local/go/src/encoding/json/stream.go:429
			// _ = "end of CoverTab[28501]"
//line /usr/local/go/src/encoding/json/stream.go:429
			_go_fuzz_dep_.CoverTab[28502]++
									if dec.tokenState == tokenObjectComma {
//line /usr/local/go/src/encoding/json/stream.go:430
				_go_fuzz_dep_.CoverTab[28523]++
										dec.scanp++
										dec.tokenState = tokenObjectKey
										continue
//line /usr/local/go/src/encoding/json/stream.go:433
				// _ = "end of CoverTab[28523]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:434
				_go_fuzz_dep_.CoverTab[28524]++
//line /usr/local/go/src/encoding/json/stream.go:434
				// _ = "end of CoverTab[28524]"
//line /usr/local/go/src/encoding/json/stream.go:434
			}
//line /usr/local/go/src/encoding/json/stream.go:434
			// _ = "end of CoverTab[28502]"
//line /usr/local/go/src/encoding/json/stream.go:434
			_go_fuzz_dep_.CoverTab[28503]++
									return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:435
			// _ = "end of CoverTab[28503]"

		case '"':
//line /usr/local/go/src/encoding/json/stream.go:437
			_go_fuzz_dep_.CoverTab[28504]++
									if dec.tokenState == tokenObjectStart || func() bool {
//line /usr/local/go/src/encoding/json/stream.go:438
				_go_fuzz_dep_.CoverTab[28525]++
//line /usr/local/go/src/encoding/json/stream.go:438
				return dec.tokenState == tokenObjectKey
//line /usr/local/go/src/encoding/json/stream.go:438
				// _ = "end of CoverTab[28525]"
//line /usr/local/go/src/encoding/json/stream.go:438
			}() {
//line /usr/local/go/src/encoding/json/stream.go:438
				_go_fuzz_dep_.CoverTab[28526]++
										var x string
										old := dec.tokenState
										dec.tokenState = tokenTopValue
										err := dec.Decode(&x)
										dec.tokenState = old
										if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:444
					_go_fuzz_dep_.CoverTab[28528]++
											return nil, err
//line /usr/local/go/src/encoding/json/stream.go:445
					// _ = "end of CoverTab[28528]"
				} else {
//line /usr/local/go/src/encoding/json/stream.go:446
					_go_fuzz_dep_.CoverTab[28529]++
//line /usr/local/go/src/encoding/json/stream.go:446
					// _ = "end of CoverTab[28529]"
//line /usr/local/go/src/encoding/json/stream.go:446
				}
//line /usr/local/go/src/encoding/json/stream.go:446
				// _ = "end of CoverTab[28526]"
//line /usr/local/go/src/encoding/json/stream.go:446
				_go_fuzz_dep_.CoverTab[28527]++
										dec.tokenState = tokenObjectColon
										return x, nil
//line /usr/local/go/src/encoding/json/stream.go:448
				// _ = "end of CoverTab[28527]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:449
				_go_fuzz_dep_.CoverTab[28530]++
//line /usr/local/go/src/encoding/json/stream.go:449
				// _ = "end of CoverTab[28530]"
//line /usr/local/go/src/encoding/json/stream.go:449
			}
//line /usr/local/go/src/encoding/json/stream.go:449
			// _ = "end of CoverTab[28504]"
//line /usr/local/go/src/encoding/json/stream.go:449
			_go_fuzz_dep_.CoverTab[28505]++
									fallthrough
//line /usr/local/go/src/encoding/json/stream.go:450
			// _ = "end of CoverTab[28505]"

		default:
//line /usr/local/go/src/encoding/json/stream.go:452
			_go_fuzz_dep_.CoverTab[28506]++
									if !dec.tokenValueAllowed() {
//line /usr/local/go/src/encoding/json/stream.go:453
				_go_fuzz_dep_.CoverTab[28531]++
										return dec.tokenError(c)
//line /usr/local/go/src/encoding/json/stream.go:454
				// _ = "end of CoverTab[28531]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:455
				_go_fuzz_dep_.CoverTab[28532]++
//line /usr/local/go/src/encoding/json/stream.go:455
				// _ = "end of CoverTab[28532]"
//line /usr/local/go/src/encoding/json/stream.go:455
			}
//line /usr/local/go/src/encoding/json/stream.go:455
			// _ = "end of CoverTab[28506]"
//line /usr/local/go/src/encoding/json/stream.go:455
			_go_fuzz_dep_.CoverTab[28507]++
									var x any
									if err := dec.Decode(&x); err != nil {
//line /usr/local/go/src/encoding/json/stream.go:457
				_go_fuzz_dep_.CoverTab[28533]++
										return nil, err
//line /usr/local/go/src/encoding/json/stream.go:458
				// _ = "end of CoverTab[28533]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:459
				_go_fuzz_dep_.CoverTab[28534]++
//line /usr/local/go/src/encoding/json/stream.go:459
				// _ = "end of CoverTab[28534]"
//line /usr/local/go/src/encoding/json/stream.go:459
			}
//line /usr/local/go/src/encoding/json/stream.go:459
			// _ = "end of CoverTab[28507]"
//line /usr/local/go/src/encoding/json/stream.go:459
			_go_fuzz_dep_.CoverTab[28508]++
									return x, nil
//line /usr/local/go/src/encoding/json/stream.go:460
			// _ = "end of CoverTab[28508]"
		}
//line /usr/local/go/src/encoding/json/stream.go:461
		// _ = "end of CoverTab[28488]"
	}
//line /usr/local/go/src/encoding/json/stream.go:462
	// _ = "end of CoverTab[28486]"
}

func (dec *Decoder) tokenError(c byte) (Token, error) {
//line /usr/local/go/src/encoding/json/stream.go:465
	_go_fuzz_dep_.CoverTab[28535]++
							var context string
							switch dec.tokenState {
	case tokenTopValue:
//line /usr/local/go/src/encoding/json/stream.go:468
		_go_fuzz_dep_.CoverTab[28537]++
								context = " looking for beginning of value"
//line /usr/local/go/src/encoding/json/stream.go:469
		// _ = "end of CoverTab[28537]"
	case tokenArrayStart, tokenArrayValue, tokenObjectValue:
//line /usr/local/go/src/encoding/json/stream.go:470
		_go_fuzz_dep_.CoverTab[28538]++
								context = " looking for beginning of value"
//line /usr/local/go/src/encoding/json/stream.go:471
		// _ = "end of CoverTab[28538]"
	case tokenArrayComma:
//line /usr/local/go/src/encoding/json/stream.go:472
		_go_fuzz_dep_.CoverTab[28539]++
								context = " after array element"
//line /usr/local/go/src/encoding/json/stream.go:473
		// _ = "end of CoverTab[28539]"
	case tokenObjectKey:
//line /usr/local/go/src/encoding/json/stream.go:474
		_go_fuzz_dep_.CoverTab[28540]++
								context = " looking for beginning of object key string"
//line /usr/local/go/src/encoding/json/stream.go:475
		// _ = "end of CoverTab[28540]"
	case tokenObjectColon:
//line /usr/local/go/src/encoding/json/stream.go:476
		_go_fuzz_dep_.CoverTab[28541]++
								context = " after object key"
//line /usr/local/go/src/encoding/json/stream.go:477
		// _ = "end of CoverTab[28541]"
	case tokenObjectComma:
//line /usr/local/go/src/encoding/json/stream.go:478
		_go_fuzz_dep_.CoverTab[28542]++
								context = " after object key:value pair"
//line /usr/local/go/src/encoding/json/stream.go:479
		// _ = "end of CoverTab[28542]"
//line /usr/local/go/src/encoding/json/stream.go:479
	default:
//line /usr/local/go/src/encoding/json/stream.go:479
		_go_fuzz_dep_.CoverTab[28543]++
//line /usr/local/go/src/encoding/json/stream.go:479
		// _ = "end of CoverTab[28543]"
	}
//line /usr/local/go/src/encoding/json/stream.go:480
	// _ = "end of CoverTab[28535]"
//line /usr/local/go/src/encoding/json/stream.go:480
	_go_fuzz_dep_.CoverTab[28536]++
							return nil, &SyntaxError{"invalid character " + quoteChar(c) + context, dec.InputOffset()}
//line /usr/local/go/src/encoding/json/stream.go:481
	// _ = "end of CoverTab[28536]"
}

// More reports whether there is another element in the
//line /usr/local/go/src/encoding/json/stream.go:484
// current array or object being parsed.
//line /usr/local/go/src/encoding/json/stream.go:486
func (dec *Decoder) More() bool {
//line /usr/local/go/src/encoding/json/stream.go:486
	_go_fuzz_dep_.CoverTab[28544]++
							c, err := dec.peek()
							return err == nil && func() bool {
//line /usr/local/go/src/encoding/json/stream.go:488
		_go_fuzz_dep_.CoverTab[28545]++
//line /usr/local/go/src/encoding/json/stream.go:488
		return c != ']'
//line /usr/local/go/src/encoding/json/stream.go:488
		// _ = "end of CoverTab[28545]"
//line /usr/local/go/src/encoding/json/stream.go:488
	}() && func() bool {
//line /usr/local/go/src/encoding/json/stream.go:488
		_go_fuzz_dep_.CoverTab[28546]++
//line /usr/local/go/src/encoding/json/stream.go:488
		return c != '}'
//line /usr/local/go/src/encoding/json/stream.go:488
		// _ = "end of CoverTab[28546]"
//line /usr/local/go/src/encoding/json/stream.go:488
	}()
//line /usr/local/go/src/encoding/json/stream.go:488
	// _ = "end of CoverTab[28544]"
}

func (dec *Decoder) peek() (byte, error) {
//line /usr/local/go/src/encoding/json/stream.go:491
	_go_fuzz_dep_.CoverTab[28547]++
							var err error
							for {
//line /usr/local/go/src/encoding/json/stream.go:493
		_go_fuzz_dep_.CoverTab[28548]++
								for i := dec.scanp; i < len(dec.buf); i++ {
//line /usr/local/go/src/encoding/json/stream.go:494
			_go_fuzz_dep_.CoverTab[28551]++
									c := dec.buf[i]
									if isSpace(c) {
//line /usr/local/go/src/encoding/json/stream.go:496
				_go_fuzz_dep_.CoverTab[28553]++
										continue
//line /usr/local/go/src/encoding/json/stream.go:497
				// _ = "end of CoverTab[28553]"
			} else {
//line /usr/local/go/src/encoding/json/stream.go:498
				_go_fuzz_dep_.CoverTab[28554]++
//line /usr/local/go/src/encoding/json/stream.go:498
				// _ = "end of CoverTab[28554]"
//line /usr/local/go/src/encoding/json/stream.go:498
			}
//line /usr/local/go/src/encoding/json/stream.go:498
			// _ = "end of CoverTab[28551]"
//line /usr/local/go/src/encoding/json/stream.go:498
			_go_fuzz_dep_.CoverTab[28552]++
									dec.scanp = i
									return c, nil
//line /usr/local/go/src/encoding/json/stream.go:500
			// _ = "end of CoverTab[28552]"
		}
//line /usr/local/go/src/encoding/json/stream.go:501
		// _ = "end of CoverTab[28548]"
//line /usr/local/go/src/encoding/json/stream.go:501
		_go_fuzz_dep_.CoverTab[28549]++

								if err != nil {
//line /usr/local/go/src/encoding/json/stream.go:503
			_go_fuzz_dep_.CoverTab[28555]++
									return 0, err
//line /usr/local/go/src/encoding/json/stream.go:504
			// _ = "end of CoverTab[28555]"
		} else {
//line /usr/local/go/src/encoding/json/stream.go:505
			_go_fuzz_dep_.CoverTab[28556]++
//line /usr/local/go/src/encoding/json/stream.go:505
			// _ = "end of CoverTab[28556]"
//line /usr/local/go/src/encoding/json/stream.go:505
		}
//line /usr/local/go/src/encoding/json/stream.go:505
		// _ = "end of CoverTab[28549]"
//line /usr/local/go/src/encoding/json/stream.go:505
		_go_fuzz_dep_.CoverTab[28550]++
								err = dec.refill()
//line /usr/local/go/src/encoding/json/stream.go:506
		// _ = "end of CoverTab[28550]"
	}
//line /usr/local/go/src/encoding/json/stream.go:507
	// _ = "end of CoverTab[28547]"
}

// InputOffset returns the input stream byte offset of the current decoder position.
//line /usr/local/go/src/encoding/json/stream.go:510
// The offset gives the location of the end of the most recently returned token
//line /usr/local/go/src/encoding/json/stream.go:510
// and the beginning of the next token.
//line /usr/local/go/src/encoding/json/stream.go:513
func (dec *Decoder) InputOffset() int64 {
//line /usr/local/go/src/encoding/json/stream.go:513
	_go_fuzz_dep_.CoverTab[28557]++
							return dec.scanned + int64(dec.scanp)
//line /usr/local/go/src/encoding/json/stream.go:514
	// _ = "end of CoverTab[28557]"
}

//line /usr/local/go/src/encoding/json/stream.go:515
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/stream.go:515
var _ = _go_fuzz_dep_.CoverTab
