// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/hex/hex.go:5
// Package hex implements hexadecimal encoding and decoding.
package hex

//line /usr/local/go/src/encoding/hex/hex.go:6
import (
//line /usr/local/go/src/encoding/hex/hex.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/hex/hex.go:6
)
//line /usr/local/go/src/encoding/hex/hex.go:6
import (
//line /usr/local/go/src/encoding/hex/hex.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/hex/hex.go:6
)

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	hextable	= "0123456789abcdef"
	reverseHexTable	= "" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\xff\xff\xff\xff\xff\xff" +
		"\xff\x0a\x0b\x0c\x0d\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\x0a\x0b\x0c\x0d\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" +
		"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"
)

// EncodedLen returns the length of an encoding of n source bytes.
//line /usr/local/go/src/encoding/hex/hex.go:36
// Specifically, it returns n * 2.
//line /usr/local/go/src/encoding/hex/hex.go:38
func EncodedLen(n int) int {
//line /usr/local/go/src/encoding/hex/hex.go:38
	_go_fuzz_dep_.CoverTab[10334]++
//line /usr/local/go/src/encoding/hex/hex.go:38
	return n * 2
//line /usr/local/go/src/encoding/hex/hex.go:38
	// _ = "end of CoverTab[10334]"
//line /usr/local/go/src/encoding/hex/hex.go:38
}

// Encode encodes src into EncodedLen(len(src))
//line /usr/local/go/src/encoding/hex/hex.go:40
// bytes of dst. As a convenience, it returns the number
//line /usr/local/go/src/encoding/hex/hex.go:40
// of bytes written to dst, but this value is always EncodedLen(len(src)).
//line /usr/local/go/src/encoding/hex/hex.go:40
// Encode implements hexadecimal encoding.
//line /usr/local/go/src/encoding/hex/hex.go:44
func Encode(dst, src []byte) int {
//line /usr/local/go/src/encoding/hex/hex.go:44
	_go_fuzz_dep_.CoverTab[10335]++
							j := 0
							for _, v := range src {
//line /usr/local/go/src/encoding/hex/hex.go:46
		_go_fuzz_dep_.CoverTab[10337]++
								dst[j] = hextable[v>>4]
								dst[j+1] = hextable[v&0x0f]
								j += 2
//line /usr/local/go/src/encoding/hex/hex.go:49
		// _ = "end of CoverTab[10337]"
	}
//line /usr/local/go/src/encoding/hex/hex.go:50
	// _ = "end of CoverTab[10335]"
//line /usr/local/go/src/encoding/hex/hex.go:50
	_go_fuzz_dep_.CoverTab[10336]++
							return len(src) * 2
//line /usr/local/go/src/encoding/hex/hex.go:51
	// _ = "end of CoverTab[10336]"
}

// ErrLength reports an attempt to decode an odd-length input
//line /usr/local/go/src/encoding/hex/hex.go:54
// using Decode or DecodeString.
//line /usr/local/go/src/encoding/hex/hex.go:54
// The stream-based Decoder returns io.ErrUnexpectedEOF instead of ErrLength.
//line /usr/local/go/src/encoding/hex/hex.go:57
var ErrLength = errors.New("encoding/hex: odd length hex string")

// InvalidByteError values describe errors resulting from an invalid byte in a hex string.
type InvalidByteError byte

func (e InvalidByteError) Error() string {
//line /usr/local/go/src/encoding/hex/hex.go:62
	_go_fuzz_dep_.CoverTab[10338]++
							return fmt.Sprintf("encoding/hex: invalid byte: %#U", rune(e))
//line /usr/local/go/src/encoding/hex/hex.go:63
	// _ = "end of CoverTab[10338]"
}

// DecodedLen returns the length of a decoding of x source bytes.
//line /usr/local/go/src/encoding/hex/hex.go:66
// Specifically, it returns x / 2.
//line /usr/local/go/src/encoding/hex/hex.go:68
func DecodedLen(x int) int {
//line /usr/local/go/src/encoding/hex/hex.go:68
	_go_fuzz_dep_.CoverTab[10339]++
//line /usr/local/go/src/encoding/hex/hex.go:68
	return x / 2
//line /usr/local/go/src/encoding/hex/hex.go:68
	// _ = "end of CoverTab[10339]"
//line /usr/local/go/src/encoding/hex/hex.go:68
}

// Decode decodes src into DecodedLen(len(src)) bytes,
//line /usr/local/go/src/encoding/hex/hex.go:70
// returning the actual number of bytes written to dst.
//line /usr/local/go/src/encoding/hex/hex.go:70
//
//line /usr/local/go/src/encoding/hex/hex.go:70
// Decode expects that src contains only hexadecimal
//line /usr/local/go/src/encoding/hex/hex.go:70
// characters and that src has even length.
//line /usr/local/go/src/encoding/hex/hex.go:70
// If the input is malformed, Decode returns the number
//line /usr/local/go/src/encoding/hex/hex.go:70
// of bytes decoded before the error.
//line /usr/local/go/src/encoding/hex/hex.go:77
func Decode(dst, src []byte) (int, error) {
//line /usr/local/go/src/encoding/hex/hex.go:77
	_go_fuzz_dep_.CoverTab[10340]++
							i, j := 0, 1
							for ; j < len(src); j += 2 {
//line /usr/local/go/src/encoding/hex/hex.go:79
		_go_fuzz_dep_.CoverTab[10343]++
								p := src[j-1]
								q := src[j]

								a := reverseHexTable[p]
								b := reverseHexTable[q]
								if a > 0x0f {
//line /usr/local/go/src/encoding/hex/hex.go:85
			_go_fuzz_dep_.CoverTab[10346]++
									return i, InvalidByteError(p)
//line /usr/local/go/src/encoding/hex/hex.go:86
			// _ = "end of CoverTab[10346]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:87
			_go_fuzz_dep_.CoverTab[10347]++
//line /usr/local/go/src/encoding/hex/hex.go:87
			// _ = "end of CoverTab[10347]"
//line /usr/local/go/src/encoding/hex/hex.go:87
		}
//line /usr/local/go/src/encoding/hex/hex.go:87
		// _ = "end of CoverTab[10343]"
//line /usr/local/go/src/encoding/hex/hex.go:87
		_go_fuzz_dep_.CoverTab[10344]++
								if b > 0x0f {
//line /usr/local/go/src/encoding/hex/hex.go:88
			_go_fuzz_dep_.CoverTab[10348]++
									return i, InvalidByteError(q)
//line /usr/local/go/src/encoding/hex/hex.go:89
			// _ = "end of CoverTab[10348]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:90
			_go_fuzz_dep_.CoverTab[10349]++
//line /usr/local/go/src/encoding/hex/hex.go:90
			// _ = "end of CoverTab[10349]"
//line /usr/local/go/src/encoding/hex/hex.go:90
		}
//line /usr/local/go/src/encoding/hex/hex.go:90
		// _ = "end of CoverTab[10344]"
//line /usr/local/go/src/encoding/hex/hex.go:90
		_go_fuzz_dep_.CoverTab[10345]++
								dst[i] = (a << 4) | b
								i++
//line /usr/local/go/src/encoding/hex/hex.go:92
		// _ = "end of CoverTab[10345]"
	}
//line /usr/local/go/src/encoding/hex/hex.go:93
	// _ = "end of CoverTab[10340]"
//line /usr/local/go/src/encoding/hex/hex.go:93
	_go_fuzz_dep_.CoverTab[10341]++
							if len(src)%2 == 1 {
//line /usr/local/go/src/encoding/hex/hex.go:94
		_go_fuzz_dep_.CoverTab[10350]++

//line /usr/local/go/src/encoding/hex/hex.go:97
		if reverseHexTable[src[j-1]] > 0x0f {
//line /usr/local/go/src/encoding/hex/hex.go:97
			_go_fuzz_dep_.CoverTab[10352]++
									return i, InvalidByteError(src[j-1])
//line /usr/local/go/src/encoding/hex/hex.go:98
			// _ = "end of CoverTab[10352]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:99
			_go_fuzz_dep_.CoverTab[10353]++
//line /usr/local/go/src/encoding/hex/hex.go:99
			// _ = "end of CoverTab[10353]"
//line /usr/local/go/src/encoding/hex/hex.go:99
		}
//line /usr/local/go/src/encoding/hex/hex.go:99
		// _ = "end of CoverTab[10350]"
//line /usr/local/go/src/encoding/hex/hex.go:99
		_go_fuzz_dep_.CoverTab[10351]++
								return i, ErrLength
//line /usr/local/go/src/encoding/hex/hex.go:100
		// _ = "end of CoverTab[10351]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:101
		_go_fuzz_dep_.CoverTab[10354]++
//line /usr/local/go/src/encoding/hex/hex.go:101
		// _ = "end of CoverTab[10354]"
//line /usr/local/go/src/encoding/hex/hex.go:101
	}
//line /usr/local/go/src/encoding/hex/hex.go:101
	// _ = "end of CoverTab[10341]"
//line /usr/local/go/src/encoding/hex/hex.go:101
	_go_fuzz_dep_.CoverTab[10342]++
							return i, nil
//line /usr/local/go/src/encoding/hex/hex.go:102
	// _ = "end of CoverTab[10342]"
}

// EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string {
//line /usr/local/go/src/encoding/hex/hex.go:106
	_go_fuzz_dep_.CoverTab[10355]++
							dst := make([]byte, EncodedLen(len(src)))
							Encode(dst, src)
							return string(dst)
//line /usr/local/go/src/encoding/hex/hex.go:109
	// _ = "end of CoverTab[10355]"
}

// DecodeString returns the bytes represented by the hexadecimal string s.
//line /usr/local/go/src/encoding/hex/hex.go:112
//
//line /usr/local/go/src/encoding/hex/hex.go:112
// DecodeString expects that src contains only hexadecimal
//line /usr/local/go/src/encoding/hex/hex.go:112
// characters and that src has even length.
//line /usr/local/go/src/encoding/hex/hex.go:112
// If the input is malformed, DecodeString returns
//line /usr/local/go/src/encoding/hex/hex.go:112
// the bytes decoded before the error.
//line /usr/local/go/src/encoding/hex/hex.go:118
func DecodeString(s string) ([]byte, error) {
//line /usr/local/go/src/encoding/hex/hex.go:118
	_go_fuzz_dep_.CoverTab[10356]++
							src := []byte(s)

//line /usr/local/go/src/encoding/hex/hex.go:122
	n, err := Decode(src, src)
							return src[:n], err
//line /usr/local/go/src/encoding/hex/hex.go:123
	// _ = "end of CoverTab[10356]"
}

// Dump returns a string that contains a hex dump of the given data. The format
//line /usr/local/go/src/encoding/hex/hex.go:126
// of the hex dump matches the output of `hexdump -C` on the command line.
//line /usr/local/go/src/encoding/hex/hex.go:128
func Dump(data []byte) string {
//line /usr/local/go/src/encoding/hex/hex.go:128
	_go_fuzz_dep_.CoverTab[10357]++
							if len(data) == 0 {
//line /usr/local/go/src/encoding/hex/hex.go:129
		_go_fuzz_dep_.CoverTab[10359]++
								return ""
//line /usr/local/go/src/encoding/hex/hex.go:130
		// _ = "end of CoverTab[10359]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:131
		_go_fuzz_dep_.CoverTab[10360]++
//line /usr/local/go/src/encoding/hex/hex.go:131
		// _ = "end of CoverTab[10360]"
//line /usr/local/go/src/encoding/hex/hex.go:131
	}
//line /usr/local/go/src/encoding/hex/hex.go:131
	// _ = "end of CoverTab[10357]"
//line /usr/local/go/src/encoding/hex/hex.go:131
	_go_fuzz_dep_.CoverTab[10358]++

							var buf strings.Builder

//line /usr/local/go/src/encoding/hex/hex.go:137
	buf.Grow((1 + ((len(data) - 1) / 16)) * 79)

							dumper := Dumper(&buf)
							dumper.Write(data)
							dumper.Close()
							return buf.String()
//line /usr/local/go/src/encoding/hex/hex.go:142
	// _ = "end of CoverTab[10358]"
}

// bufferSize is the number of hexadecimal characters to buffer in encoder and decoder.
const bufferSize = 1024

type encoder struct {
	w	io.Writer
	err	error
	out	[bufferSize]byte	// output buffer
}

// NewEncoder returns an io.Writer that writes lowercase hexadecimal characters to w.
func NewEncoder(w io.Writer) io.Writer {
//line /usr/local/go/src/encoding/hex/hex.go:155
	_go_fuzz_dep_.CoverTab[10361]++
							return &encoder{w: w}
//line /usr/local/go/src/encoding/hex/hex.go:156
	// _ = "end of CoverTab[10361]"
}

func (e *encoder) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/encoding/hex/hex.go:159
	_go_fuzz_dep_.CoverTab[10362]++
							for len(p) > 0 && func() bool {
//line /usr/local/go/src/encoding/hex/hex.go:160
		_go_fuzz_dep_.CoverTab[10364]++
//line /usr/local/go/src/encoding/hex/hex.go:160
		return e.err == nil
//line /usr/local/go/src/encoding/hex/hex.go:160
		// _ = "end of CoverTab[10364]"
//line /usr/local/go/src/encoding/hex/hex.go:160
	}() {
//line /usr/local/go/src/encoding/hex/hex.go:160
		_go_fuzz_dep_.CoverTab[10365]++
								chunkSize := bufferSize / 2
								if len(p) < chunkSize {
//line /usr/local/go/src/encoding/hex/hex.go:162
			_go_fuzz_dep_.CoverTab[10367]++
									chunkSize = len(p)
//line /usr/local/go/src/encoding/hex/hex.go:163
			// _ = "end of CoverTab[10367]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:164
			_go_fuzz_dep_.CoverTab[10368]++
//line /usr/local/go/src/encoding/hex/hex.go:164
			// _ = "end of CoverTab[10368]"
//line /usr/local/go/src/encoding/hex/hex.go:164
		}
//line /usr/local/go/src/encoding/hex/hex.go:164
		// _ = "end of CoverTab[10365]"
//line /usr/local/go/src/encoding/hex/hex.go:164
		_go_fuzz_dep_.CoverTab[10366]++

								var written int
								encoded := Encode(e.out[:], p[:chunkSize])
								written, e.err = e.w.Write(e.out[:encoded])
								n += written / 2
								p = p[chunkSize:]
//line /usr/local/go/src/encoding/hex/hex.go:170
		// _ = "end of CoverTab[10366]"
	}
//line /usr/local/go/src/encoding/hex/hex.go:171
	// _ = "end of CoverTab[10362]"
//line /usr/local/go/src/encoding/hex/hex.go:171
	_go_fuzz_dep_.CoverTab[10363]++
							return n, e.err
//line /usr/local/go/src/encoding/hex/hex.go:172
	// _ = "end of CoverTab[10363]"
}

type decoder struct {
	r	io.Reader
	err	error
	in	[]byte			// input buffer (encoded form)
	arr	[bufferSize]byte	// backing array for in
}

// NewDecoder returns an io.Reader that decodes hexadecimal characters from r.
//line /usr/local/go/src/encoding/hex/hex.go:182
// NewDecoder expects that r contain only an even number of hexadecimal characters.
//line /usr/local/go/src/encoding/hex/hex.go:184
func NewDecoder(r io.Reader) io.Reader {
//line /usr/local/go/src/encoding/hex/hex.go:184
	_go_fuzz_dep_.CoverTab[10369]++
							return &decoder{r: r}
//line /usr/local/go/src/encoding/hex/hex.go:185
	// _ = "end of CoverTab[10369]"
}

func (d *decoder) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/encoding/hex/hex.go:188
	_go_fuzz_dep_.CoverTab[10370]++

							if len(d.in) < 2 && func() bool {
//line /usr/local/go/src/encoding/hex/hex.go:190
		_go_fuzz_dep_.CoverTab[10375]++
//line /usr/local/go/src/encoding/hex/hex.go:190
		return d.err == nil
//line /usr/local/go/src/encoding/hex/hex.go:190
		// _ = "end of CoverTab[10375]"
//line /usr/local/go/src/encoding/hex/hex.go:190
	}() {
//line /usr/local/go/src/encoding/hex/hex.go:190
		_go_fuzz_dep_.CoverTab[10376]++
								var numCopy, numRead int
								numCopy = copy(d.arr[:], d.in)
								numRead, d.err = d.r.Read(d.arr[numCopy:])
								d.in = d.arr[:numCopy+numRead]
								if d.err == io.EOF && func() bool {
//line /usr/local/go/src/encoding/hex/hex.go:195
			_go_fuzz_dep_.CoverTab[10377]++
//line /usr/local/go/src/encoding/hex/hex.go:195
			return len(d.in)%2 != 0
//line /usr/local/go/src/encoding/hex/hex.go:195
			// _ = "end of CoverTab[10377]"
//line /usr/local/go/src/encoding/hex/hex.go:195
		}() {
//line /usr/local/go/src/encoding/hex/hex.go:195
			_go_fuzz_dep_.CoverTab[10378]++

									if a := reverseHexTable[d.in[len(d.in)-1]]; a > 0x0f {
//line /usr/local/go/src/encoding/hex/hex.go:197
				_go_fuzz_dep_.CoverTab[10379]++
										d.err = InvalidByteError(d.in[len(d.in)-1])
//line /usr/local/go/src/encoding/hex/hex.go:198
				// _ = "end of CoverTab[10379]"
			} else {
//line /usr/local/go/src/encoding/hex/hex.go:199
				_go_fuzz_dep_.CoverTab[10380]++
										d.err = io.ErrUnexpectedEOF
//line /usr/local/go/src/encoding/hex/hex.go:200
				// _ = "end of CoverTab[10380]"
			}
//line /usr/local/go/src/encoding/hex/hex.go:201
			// _ = "end of CoverTab[10378]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:202
			_go_fuzz_dep_.CoverTab[10381]++
//line /usr/local/go/src/encoding/hex/hex.go:202
			// _ = "end of CoverTab[10381]"
//line /usr/local/go/src/encoding/hex/hex.go:202
		}
//line /usr/local/go/src/encoding/hex/hex.go:202
		// _ = "end of CoverTab[10376]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:203
		_go_fuzz_dep_.CoverTab[10382]++
//line /usr/local/go/src/encoding/hex/hex.go:203
		// _ = "end of CoverTab[10382]"
//line /usr/local/go/src/encoding/hex/hex.go:203
	}
//line /usr/local/go/src/encoding/hex/hex.go:203
	// _ = "end of CoverTab[10370]"
//line /usr/local/go/src/encoding/hex/hex.go:203
	_go_fuzz_dep_.CoverTab[10371]++

//line /usr/local/go/src/encoding/hex/hex.go:206
	if numAvail := len(d.in) / 2; len(p) > numAvail {
//line /usr/local/go/src/encoding/hex/hex.go:206
		_go_fuzz_dep_.CoverTab[10383]++
								p = p[:numAvail]
//line /usr/local/go/src/encoding/hex/hex.go:207
		// _ = "end of CoverTab[10383]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:208
		_go_fuzz_dep_.CoverTab[10384]++
//line /usr/local/go/src/encoding/hex/hex.go:208
		// _ = "end of CoverTab[10384]"
//line /usr/local/go/src/encoding/hex/hex.go:208
	}
//line /usr/local/go/src/encoding/hex/hex.go:208
	// _ = "end of CoverTab[10371]"
//line /usr/local/go/src/encoding/hex/hex.go:208
	_go_fuzz_dep_.CoverTab[10372]++
							numDec, err := Decode(p, d.in[:len(p)*2])
							d.in = d.in[2*numDec:]
							if err != nil {
//line /usr/local/go/src/encoding/hex/hex.go:211
		_go_fuzz_dep_.CoverTab[10385]++
								d.in, d.err = nil, err
//line /usr/local/go/src/encoding/hex/hex.go:212
		// _ = "end of CoverTab[10385]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:213
		_go_fuzz_dep_.CoverTab[10386]++
//line /usr/local/go/src/encoding/hex/hex.go:213
		// _ = "end of CoverTab[10386]"
//line /usr/local/go/src/encoding/hex/hex.go:213
	}
//line /usr/local/go/src/encoding/hex/hex.go:213
	// _ = "end of CoverTab[10372]"
//line /usr/local/go/src/encoding/hex/hex.go:213
	_go_fuzz_dep_.CoverTab[10373]++

							if len(d.in) < 2 {
//line /usr/local/go/src/encoding/hex/hex.go:215
		_go_fuzz_dep_.CoverTab[10387]++
								return numDec, d.err
//line /usr/local/go/src/encoding/hex/hex.go:216
		// _ = "end of CoverTab[10387]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:217
		_go_fuzz_dep_.CoverTab[10388]++
//line /usr/local/go/src/encoding/hex/hex.go:217
		// _ = "end of CoverTab[10388]"
//line /usr/local/go/src/encoding/hex/hex.go:217
	}
//line /usr/local/go/src/encoding/hex/hex.go:217
	// _ = "end of CoverTab[10373]"
//line /usr/local/go/src/encoding/hex/hex.go:217
	_go_fuzz_dep_.CoverTab[10374]++
							return numDec, nil
//line /usr/local/go/src/encoding/hex/hex.go:218
	// _ = "end of CoverTab[10374]"
}

// Dumper returns a WriteCloser that writes a hex dump of all written data to
//line /usr/local/go/src/encoding/hex/hex.go:221
// w. The format of the dump matches the output of `hexdump -C` on the command
//line /usr/local/go/src/encoding/hex/hex.go:221
// line.
//line /usr/local/go/src/encoding/hex/hex.go:224
func Dumper(w io.Writer) io.WriteCloser {
//line /usr/local/go/src/encoding/hex/hex.go:224
	_go_fuzz_dep_.CoverTab[10389]++
							return &dumper{w: w}
//line /usr/local/go/src/encoding/hex/hex.go:225
	// _ = "end of CoverTab[10389]"
}

type dumper struct {
	w		io.Writer
	rightChars	[18]byte
	buf		[14]byte
	used		int	// number of bytes in the current line
	n		uint	// number of bytes, total
	closed		bool
}

func toChar(b byte) byte {
//line /usr/local/go/src/encoding/hex/hex.go:237
	_go_fuzz_dep_.CoverTab[10390]++
							if b < 32 || func() bool {
//line /usr/local/go/src/encoding/hex/hex.go:238
		_go_fuzz_dep_.CoverTab[10392]++
//line /usr/local/go/src/encoding/hex/hex.go:238
		return b > 126
//line /usr/local/go/src/encoding/hex/hex.go:238
		// _ = "end of CoverTab[10392]"
//line /usr/local/go/src/encoding/hex/hex.go:238
	}() {
//line /usr/local/go/src/encoding/hex/hex.go:238
		_go_fuzz_dep_.CoverTab[10393]++
								return '.'
//line /usr/local/go/src/encoding/hex/hex.go:239
		// _ = "end of CoverTab[10393]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:240
		_go_fuzz_dep_.CoverTab[10394]++
//line /usr/local/go/src/encoding/hex/hex.go:240
		// _ = "end of CoverTab[10394]"
//line /usr/local/go/src/encoding/hex/hex.go:240
	}
//line /usr/local/go/src/encoding/hex/hex.go:240
	// _ = "end of CoverTab[10390]"
//line /usr/local/go/src/encoding/hex/hex.go:240
	_go_fuzz_dep_.CoverTab[10391]++
							return b
//line /usr/local/go/src/encoding/hex/hex.go:241
	// _ = "end of CoverTab[10391]"
}

func (h *dumper) Write(data []byte) (n int, err error) {
//line /usr/local/go/src/encoding/hex/hex.go:244
	_go_fuzz_dep_.CoverTab[10395]++
							if h.closed {
//line /usr/local/go/src/encoding/hex/hex.go:245
		_go_fuzz_dep_.CoverTab[10398]++
								return 0, errors.New("encoding/hex: dumper closed")
//line /usr/local/go/src/encoding/hex/hex.go:246
		// _ = "end of CoverTab[10398]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:247
		_go_fuzz_dep_.CoverTab[10399]++
//line /usr/local/go/src/encoding/hex/hex.go:247
		// _ = "end of CoverTab[10399]"
//line /usr/local/go/src/encoding/hex/hex.go:247
	}
//line /usr/local/go/src/encoding/hex/hex.go:247
	// _ = "end of CoverTab[10395]"
//line /usr/local/go/src/encoding/hex/hex.go:247
	_go_fuzz_dep_.CoverTab[10396]++

//line /usr/local/go/src/encoding/hex/hex.go:252
	for i := range data {
//line /usr/local/go/src/encoding/hex/hex.go:252
		_go_fuzz_dep_.CoverTab[10400]++
								if h.used == 0 {
//line /usr/local/go/src/encoding/hex/hex.go:253
			_go_fuzz_dep_.CoverTab[10404]++

//line /usr/local/go/src/encoding/hex/hex.go:256
			h.buf[0] = byte(h.n >> 24)
			h.buf[1] = byte(h.n >> 16)
			h.buf[2] = byte(h.n >> 8)
			h.buf[3] = byte(h.n)
			Encode(h.buf[4:], h.buf[:4])
			h.buf[12] = ' '
			h.buf[13] = ' '
			_, err = h.w.Write(h.buf[4:])
			if err != nil {
//line /usr/local/go/src/encoding/hex/hex.go:264
				_go_fuzz_dep_.CoverTab[10405]++
										return
//line /usr/local/go/src/encoding/hex/hex.go:265
				// _ = "end of CoverTab[10405]"
			} else {
//line /usr/local/go/src/encoding/hex/hex.go:266
				_go_fuzz_dep_.CoverTab[10406]++
//line /usr/local/go/src/encoding/hex/hex.go:266
				// _ = "end of CoverTab[10406]"
//line /usr/local/go/src/encoding/hex/hex.go:266
			}
//line /usr/local/go/src/encoding/hex/hex.go:266
			// _ = "end of CoverTab[10404]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:267
			_go_fuzz_dep_.CoverTab[10407]++
//line /usr/local/go/src/encoding/hex/hex.go:267
			// _ = "end of CoverTab[10407]"
//line /usr/local/go/src/encoding/hex/hex.go:267
		}
//line /usr/local/go/src/encoding/hex/hex.go:267
		// _ = "end of CoverTab[10400]"
//line /usr/local/go/src/encoding/hex/hex.go:267
		_go_fuzz_dep_.CoverTab[10401]++
								Encode(h.buf[:], data[i:i+1])
								h.buf[2] = ' '
								l := 3
								if h.used == 7 {
//line /usr/local/go/src/encoding/hex/hex.go:271
			_go_fuzz_dep_.CoverTab[10408]++

									h.buf[3] = ' '
									l = 4
//line /usr/local/go/src/encoding/hex/hex.go:274
			// _ = "end of CoverTab[10408]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:275
			_go_fuzz_dep_.CoverTab[10409]++
//line /usr/local/go/src/encoding/hex/hex.go:275
			if h.used == 15 {
//line /usr/local/go/src/encoding/hex/hex.go:275
				_go_fuzz_dep_.CoverTab[10410]++

//line /usr/local/go/src/encoding/hex/hex.go:278
				h.buf[3] = ' '
										h.buf[4] = '|'
										l = 5
//line /usr/local/go/src/encoding/hex/hex.go:280
				// _ = "end of CoverTab[10410]"
			} else {
//line /usr/local/go/src/encoding/hex/hex.go:281
				_go_fuzz_dep_.CoverTab[10411]++
//line /usr/local/go/src/encoding/hex/hex.go:281
				// _ = "end of CoverTab[10411]"
//line /usr/local/go/src/encoding/hex/hex.go:281
			}
//line /usr/local/go/src/encoding/hex/hex.go:281
			// _ = "end of CoverTab[10409]"
//line /usr/local/go/src/encoding/hex/hex.go:281
		}
//line /usr/local/go/src/encoding/hex/hex.go:281
		// _ = "end of CoverTab[10401]"
//line /usr/local/go/src/encoding/hex/hex.go:281
		_go_fuzz_dep_.CoverTab[10402]++
								_, err = h.w.Write(h.buf[:l])
								if err != nil {
//line /usr/local/go/src/encoding/hex/hex.go:283
			_go_fuzz_dep_.CoverTab[10412]++
									return
//line /usr/local/go/src/encoding/hex/hex.go:284
			// _ = "end of CoverTab[10412]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:285
			_go_fuzz_dep_.CoverTab[10413]++
//line /usr/local/go/src/encoding/hex/hex.go:285
			// _ = "end of CoverTab[10413]"
//line /usr/local/go/src/encoding/hex/hex.go:285
		}
//line /usr/local/go/src/encoding/hex/hex.go:285
		// _ = "end of CoverTab[10402]"
//line /usr/local/go/src/encoding/hex/hex.go:285
		_go_fuzz_dep_.CoverTab[10403]++
								n++
								h.rightChars[h.used] = toChar(data[i])
								h.used++
								h.n++
								if h.used == 16 {
//line /usr/local/go/src/encoding/hex/hex.go:290
			_go_fuzz_dep_.CoverTab[10414]++
									h.rightChars[16] = '|'
									h.rightChars[17] = '\n'
									_, err = h.w.Write(h.rightChars[:])
									if err != nil {
//line /usr/local/go/src/encoding/hex/hex.go:294
				_go_fuzz_dep_.CoverTab[10416]++
										return
//line /usr/local/go/src/encoding/hex/hex.go:295
				// _ = "end of CoverTab[10416]"
			} else {
//line /usr/local/go/src/encoding/hex/hex.go:296
				_go_fuzz_dep_.CoverTab[10417]++
//line /usr/local/go/src/encoding/hex/hex.go:296
				// _ = "end of CoverTab[10417]"
//line /usr/local/go/src/encoding/hex/hex.go:296
			}
//line /usr/local/go/src/encoding/hex/hex.go:296
			// _ = "end of CoverTab[10414]"
//line /usr/local/go/src/encoding/hex/hex.go:296
			_go_fuzz_dep_.CoverTab[10415]++
									h.used = 0
//line /usr/local/go/src/encoding/hex/hex.go:297
			// _ = "end of CoverTab[10415]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:298
			_go_fuzz_dep_.CoverTab[10418]++
//line /usr/local/go/src/encoding/hex/hex.go:298
			// _ = "end of CoverTab[10418]"
//line /usr/local/go/src/encoding/hex/hex.go:298
		}
//line /usr/local/go/src/encoding/hex/hex.go:298
		// _ = "end of CoverTab[10403]"
	}
//line /usr/local/go/src/encoding/hex/hex.go:299
	// _ = "end of CoverTab[10396]"
//line /usr/local/go/src/encoding/hex/hex.go:299
	_go_fuzz_dep_.CoverTab[10397]++
							return
//line /usr/local/go/src/encoding/hex/hex.go:300
	// _ = "end of CoverTab[10397]"
}

func (h *dumper) Close() (err error) {
//line /usr/local/go/src/encoding/hex/hex.go:303
	_go_fuzz_dep_.CoverTab[10419]++

							if h.closed {
//line /usr/local/go/src/encoding/hex/hex.go:305
		_go_fuzz_dep_.CoverTab[10423]++
								return
//line /usr/local/go/src/encoding/hex/hex.go:306
		// _ = "end of CoverTab[10423]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:307
		_go_fuzz_dep_.CoverTab[10424]++
//line /usr/local/go/src/encoding/hex/hex.go:307
		// _ = "end of CoverTab[10424]"
//line /usr/local/go/src/encoding/hex/hex.go:307
	}
//line /usr/local/go/src/encoding/hex/hex.go:307
	// _ = "end of CoverTab[10419]"
//line /usr/local/go/src/encoding/hex/hex.go:307
	_go_fuzz_dep_.CoverTab[10420]++
							h.closed = true
							if h.used == 0 {
//line /usr/local/go/src/encoding/hex/hex.go:309
		_go_fuzz_dep_.CoverTab[10425]++
								return
//line /usr/local/go/src/encoding/hex/hex.go:310
		// _ = "end of CoverTab[10425]"
	} else {
//line /usr/local/go/src/encoding/hex/hex.go:311
		_go_fuzz_dep_.CoverTab[10426]++
//line /usr/local/go/src/encoding/hex/hex.go:311
		// _ = "end of CoverTab[10426]"
//line /usr/local/go/src/encoding/hex/hex.go:311
	}
//line /usr/local/go/src/encoding/hex/hex.go:311
	// _ = "end of CoverTab[10420]"
//line /usr/local/go/src/encoding/hex/hex.go:311
	_go_fuzz_dep_.CoverTab[10421]++
							h.buf[0] = ' '
							h.buf[1] = ' '
							h.buf[2] = ' '
							h.buf[3] = ' '
							h.buf[4] = '|'
							nBytes := h.used
							for h.used < 16 {
//line /usr/local/go/src/encoding/hex/hex.go:318
		_go_fuzz_dep_.CoverTab[10427]++
								l := 3
								if h.used == 7 {
//line /usr/local/go/src/encoding/hex/hex.go:320
			_go_fuzz_dep_.CoverTab[10430]++
									l = 4
//line /usr/local/go/src/encoding/hex/hex.go:321
			// _ = "end of CoverTab[10430]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:322
			_go_fuzz_dep_.CoverTab[10431]++
//line /usr/local/go/src/encoding/hex/hex.go:322
			if h.used == 15 {
//line /usr/local/go/src/encoding/hex/hex.go:322
				_go_fuzz_dep_.CoverTab[10432]++
										l = 5
//line /usr/local/go/src/encoding/hex/hex.go:323
				// _ = "end of CoverTab[10432]"
			} else {
//line /usr/local/go/src/encoding/hex/hex.go:324
				_go_fuzz_dep_.CoverTab[10433]++
//line /usr/local/go/src/encoding/hex/hex.go:324
				// _ = "end of CoverTab[10433]"
//line /usr/local/go/src/encoding/hex/hex.go:324
			}
//line /usr/local/go/src/encoding/hex/hex.go:324
			// _ = "end of CoverTab[10431]"
//line /usr/local/go/src/encoding/hex/hex.go:324
		}
//line /usr/local/go/src/encoding/hex/hex.go:324
		// _ = "end of CoverTab[10427]"
//line /usr/local/go/src/encoding/hex/hex.go:324
		_go_fuzz_dep_.CoverTab[10428]++
								_, err = h.w.Write(h.buf[:l])
								if err != nil {
//line /usr/local/go/src/encoding/hex/hex.go:326
			_go_fuzz_dep_.CoverTab[10434]++
									return
//line /usr/local/go/src/encoding/hex/hex.go:327
			// _ = "end of CoverTab[10434]"
		} else {
//line /usr/local/go/src/encoding/hex/hex.go:328
			_go_fuzz_dep_.CoverTab[10435]++
//line /usr/local/go/src/encoding/hex/hex.go:328
			// _ = "end of CoverTab[10435]"
//line /usr/local/go/src/encoding/hex/hex.go:328
		}
//line /usr/local/go/src/encoding/hex/hex.go:328
		// _ = "end of CoverTab[10428]"
//line /usr/local/go/src/encoding/hex/hex.go:328
		_go_fuzz_dep_.CoverTab[10429]++
								h.used++
//line /usr/local/go/src/encoding/hex/hex.go:329
		// _ = "end of CoverTab[10429]"
	}
//line /usr/local/go/src/encoding/hex/hex.go:330
	// _ = "end of CoverTab[10421]"
//line /usr/local/go/src/encoding/hex/hex.go:330
	_go_fuzz_dep_.CoverTab[10422]++
							h.rightChars[nBytes] = '|'
							h.rightChars[nBytes+1] = '\n'
							_, err = h.w.Write(h.rightChars[:nBytes+2])
							return
//line /usr/local/go/src/encoding/hex/hex.go:334
	// _ = "end of CoverTab[10422]"
}

//line /usr/local/go/src/encoding/hex/hex.go:335
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/hex/hex.go:335
var _ = _go_fuzz_dep_.CoverTab
