// Copyright 2011 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
package snappy

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:5
)

import (
	"encoding/binary"
	"errors"
	"io"
)

var (
	// ErrCorrupt reports that the input is invalid.
	ErrCorrupt	= errors.New("snappy: corrupt input")
	// ErrTooLarge reports that the uncompressed length is too large.
	ErrTooLarge	= errors.New("snappy: decoded block is too large")
	// ErrUnsupported reports that the input isn't supported.
	ErrUnsupported	= errors.New("snappy: unsupported input")

	errUnsupportedLiteralLength	= errors.New("snappy: unsupported literal length")
)

// DecodedLen returns the length of the decoded block.
func DecodedLen(src []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:25
	_go_fuzz_dep_.CoverTab[82033]++
										v, _, err := decodedLen(src)
										return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:27
	// _ = "end of CoverTab[82033]"
}

// decodedLen returns the length of the decoded block and the number of bytes
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:30
// that the length header occupied.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:32
func decodedLen(src []byte) (blockLen, headerLen int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:32
	_go_fuzz_dep_.CoverTab[82034]++
										v, n := binary.Uvarint(src)
										if n <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:34
		_go_fuzz_dep_.CoverTab[82037]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:34
		return v > 0xffffffff
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:34
		// _ = "end of CoverTab[82037]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:34
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:34
		_go_fuzz_dep_.CoverTab[82038]++
											return 0, 0, ErrCorrupt
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:35
		// _ = "end of CoverTab[82038]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:36
		_go_fuzz_dep_.CoverTab[82039]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:36
		// _ = "end of CoverTab[82039]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:36
	// _ = "end of CoverTab[82034]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:36
	_go_fuzz_dep_.CoverTab[82035]++

										const wordSize = 32 << (^uint(0) >> 32 & 1)
										if wordSize == 32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:39
		_go_fuzz_dep_.CoverTab[82040]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:39
		return v > 0x7fffffff
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:39
		// _ = "end of CoverTab[82040]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:39
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:39
		_go_fuzz_dep_.CoverTab[82041]++
											return 0, 0, ErrTooLarge
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:40
		// _ = "end of CoverTab[82041]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:41
		_go_fuzz_dep_.CoverTab[82042]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:41
		// _ = "end of CoverTab[82042]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:41
	// _ = "end of CoverTab[82035]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:41
	_go_fuzz_dep_.CoverTab[82036]++
										return int(v), n, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:42
	// _ = "end of CoverTab[82036]"
}

const (
	decodeErrCodeCorrupt			= 1
	decodeErrCodeUnsupportedLiteralLength	= 2
)

// Decode returns the decoded form of src. The returned slice may be a sub-
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:50
// slice of dst if dst was large enough to hold the entire decoded block.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:50
// Otherwise, a newly allocated slice will be returned.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:50
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:50
// The dst and src must not overlap. It is valid to pass a nil dst.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:50
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:50
// Decode handles the Snappy block format, not the Snappy stream format.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:57
func Decode(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:57
	_go_fuzz_dep_.CoverTab[82043]++
										dLen, s, err := decodedLen(src)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:59
		_go_fuzz_dep_.CoverTab[82047]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:60
		// _ = "end of CoverTab[82047]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:61
		_go_fuzz_dep_.CoverTab[82048]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:61
		// _ = "end of CoverTab[82048]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:61
	// _ = "end of CoverTab[82043]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:61
	_go_fuzz_dep_.CoverTab[82044]++
										if dLen <= len(dst) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:62
		_go_fuzz_dep_.CoverTab[82049]++
											dst = dst[:dLen]
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:63
		// _ = "end of CoverTab[82049]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:64
		_go_fuzz_dep_.CoverTab[82050]++
											dst = make([]byte, dLen)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:65
		// _ = "end of CoverTab[82050]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:66
	// _ = "end of CoverTab[82044]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:66
	_go_fuzz_dep_.CoverTab[82045]++
										switch decode(dst, src[s:]) {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:68
		_go_fuzz_dep_.CoverTab[82051]++
											return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:69
		// _ = "end of CoverTab[82051]"
	case decodeErrCodeUnsupportedLiteralLength:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:70
		_go_fuzz_dep_.CoverTab[82052]++
											return nil, errUnsupportedLiteralLength
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:71
		// _ = "end of CoverTab[82052]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:71
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:71
		_go_fuzz_dep_.CoverTab[82053]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:71
		// _ = "end of CoverTab[82053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:72
	// _ = "end of CoverTab[82045]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:72
	_go_fuzz_dep_.CoverTab[82046]++
										return nil, ErrCorrupt
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:73
	// _ = "end of CoverTab[82046]"
}

// NewReader returns a new Reader that decompresses from r, using the framing
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:76
// format described at
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:76
// https://github.com/google/snappy/blob/master/framing_format.txt
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:79
func NewReader(r io.Reader) *Reader {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:79
	_go_fuzz_dep_.CoverTab[82054]++
										return &Reader{
		r:		r,
		decoded:	make([]byte, maxBlockSize),
		buf:		make([]byte, maxEncodedLenOfMaxBlockSize+checksumSize),
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:84
	// _ = "end of CoverTab[82054]"
}

// Reader is an io.Reader that can read Snappy-compressed bytes.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:87
//
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:87
// Reader handles the Snappy stream format, not the Snappy block format.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:90
type Reader struct {
	r	io.Reader
	err	error
	decoded	[]byte
	buf	[]byte
	// decoded[i:j] contains decoded bytes that have not yet been passed on.
	i, j		int
	readHeader	bool
}

// Reset discards any buffered data, resets all state, and switches the Snappy
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:100
// reader to read from r. This permits reusing a Reader rather than allocating
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:100
// a new one.
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:103
func (r *Reader) Reset(reader io.Reader) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:103
	_go_fuzz_dep_.CoverTab[82055]++
											r.r = reader
											r.err = nil
											r.i = 0
											r.j = 0
											r.readHeader = false
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:108
	// _ = "end of CoverTab[82055]"
}

func (r *Reader) readFull(p []byte, allowEOF bool) (ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:111
	_go_fuzz_dep_.CoverTab[82056]++
											if _, r.err = io.ReadFull(r.r, p); r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:112
		_go_fuzz_dep_.CoverTab[82058]++
												if r.err == io.ErrUnexpectedEOF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
			_go_fuzz_dep_.CoverTab[82060]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
			return (r.err == io.EOF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
				_go_fuzz_dep_.CoverTab[82061]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
				return !allowEOF
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
				// _ = "end of CoverTab[82061]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
			}())
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
			// _ = "end of CoverTab[82060]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:113
			_go_fuzz_dep_.CoverTab[82062]++
													r.err = ErrCorrupt
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:114
			// _ = "end of CoverTab[82062]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:115
			_go_fuzz_dep_.CoverTab[82063]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:115
			// _ = "end of CoverTab[82063]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:115
		// _ = "end of CoverTab[82058]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:115
		_go_fuzz_dep_.CoverTab[82059]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:116
		// _ = "end of CoverTab[82059]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:117
		_go_fuzz_dep_.CoverTab[82064]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:117
		// _ = "end of CoverTab[82064]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:117
	// _ = "end of CoverTab[82056]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:117
	_go_fuzz_dep_.CoverTab[82057]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:118
	// _ = "end of CoverTab[82057]"
}

func (r *Reader) fill() error {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:121
	_go_fuzz_dep_.CoverTab[82065]++
											for r.i >= r.j {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:122
		_go_fuzz_dep_.CoverTab[82067]++
												if !r.readFull(r.buf[:4], true) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:123
			_go_fuzz_dep_.CoverTab[82073]++
													return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:124
			// _ = "end of CoverTab[82073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:125
			_go_fuzz_dep_.CoverTab[82074]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:125
			// _ = "end of CoverTab[82074]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:125
		// _ = "end of CoverTab[82067]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:125
		_go_fuzz_dep_.CoverTab[82068]++
												chunkType := r.buf[0]
												if !r.readHeader {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:127
			_go_fuzz_dep_.CoverTab[82075]++
													if chunkType != chunkTypeStreamIdentifier {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:128
				_go_fuzz_dep_.CoverTab[82077]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:130
				// _ = "end of CoverTab[82077]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:131
				_go_fuzz_dep_.CoverTab[82078]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:131
				// _ = "end of CoverTab[82078]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:131
			// _ = "end of CoverTab[82075]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:131
			_go_fuzz_dep_.CoverTab[82076]++
													r.readHeader = true
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:132
			// _ = "end of CoverTab[82076]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:133
			_go_fuzz_dep_.CoverTab[82079]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:133
			// _ = "end of CoverTab[82079]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:133
		// _ = "end of CoverTab[82068]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:133
		_go_fuzz_dep_.CoverTab[82069]++
												chunkLen := int(r.buf[1]) | int(r.buf[2])<<8 | int(r.buf[3])<<16
												if chunkLen > len(r.buf) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:135
			_go_fuzz_dep_.CoverTab[82080]++
													r.err = ErrUnsupported
													return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:137
			// _ = "end of CoverTab[82080]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:138
			_go_fuzz_dep_.CoverTab[82081]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:138
			// _ = "end of CoverTab[82081]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:138
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:138
		// _ = "end of CoverTab[82069]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:138
		_go_fuzz_dep_.CoverTab[82070]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:142
		switch chunkType {
		case chunkTypeCompressedData:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:143
			_go_fuzz_dep_.CoverTab[82082]++

													if chunkLen < checksumSize {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:145
				_go_fuzz_dep_.CoverTab[82100]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:147
				// _ = "end of CoverTab[82100]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:148
				_go_fuzz_dep_.CoverTab[82101]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:148
				// _ = "end of CoverTab[82101]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:148
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:148
			// _ = "end of CoverTab[82082]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:148
			_go_fuzz_dep_.CoverTab[82083]++
													buf := r.buf[:chunkLen]
													if !r.readFull(buf, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:150
				_go_fuzz_dep_.CoverTab[82102]++
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:151
				// _ = "end of CoverTab[82102]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:152
				_go_fuzz_dep_.CoverTab[82103]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:152
				// _ = "end of CoverTab[82103]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:152
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:152
			// _ = "end of CoverTab[82083]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:152
			_go_fuzz_dep_.CoverTab[82084]++
													checksum := uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24
													buf = buf[checksumSize:]

													n, err := DecodedLen(buf)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:157
				_go_fuzz_dep_.CoverTab[82104]++
														r.err = err
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:159
				// _ = "end of CoverTab[82104]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:160
				_go_fuzz_dep_.CoverTab[82105]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:160
				// _ = "end of CoverTab[82105]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:160
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:160
			// _ = "end of CoverTab[82084]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:160
			_go_fuzz_dep_.CoverTab[82085]++
													if n > len(r.decoded) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:161
				_go_fuzz_dep_.CoverTab[82106]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:163
				// _ = "end of CoverTab[82106]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:164
				_go_fuzz_dep_.CoverTab[82107]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:164
				// _ = "end of CoverTab[82107]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:164
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:164
			// _ = "end of CoverTab[82085]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:164
			_go_fuzz_dep_.CoverTab[82086]++
													if _, err := Decode(r.decoded, buf); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:165
				_go_fuzz_dep_.CoverTab[82108]++
														r.err = err
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:167
				// _ = "end of CoverTab[82108]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:168
				_go_fuzz_dep_.CoverTab[82109]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:168
				// _ = "end of CoverTab[82109]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:168
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:168
			// _ = "end of CoverTab[82086]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:168
			_go_fuzz_dep_.CoverTab[82087]++
													if crc(r.decoded[:n]) != checksum {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:169
				_go_fuzz_dep_.CoverTab[82110]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:171
				// _ = "end of CoverTab[82110]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:172
				_go_fuzz_dep_.CoverTab[82111]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:172
				// _ = "end of CoverTab[82111]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:172
			// _ = "end of CoverTab[82087]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:172
			_go_fuzz_dep_.CoverTab[82088]++
													r.i, r.j = 0, n
													continue
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:174
			// _ = "end of CoverTab[82088]"

		case chunkTypeUncompressedData:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:176
			_go_fuzz_dep_.CoverTab[82089]++

													if chunkLen < checksumSize {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:178
				_go_fuzz_dep_.CoverTab[82112]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:180
				// _ = "end of CoverTab[82112]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:181
				_go_fuzz_dep_.CoverTab[82113]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:181
				// _ = "end of CoverTab[82113]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:181
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:181
			// _ = "end of CoverTab[82089]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:181
			_go_fuzz_dep_.CoverTab[82090]++
													buf := r.buf[:checksumSize]
													if !r.readFull(buf, false) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:183
				_go_fuzz_dep_.CoverTab[82114]++
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:184
				// _ = "end of CoverTab[82114]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:185
				_go_fuzz_dep_.CoverTab[82115]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:185
				// _ = "end of CoverTab[82115]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:185
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:185
			// _ = "end of CoverTab[82090]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:185
			_go_fuzz_dep_.CoverTab[82091]++
													checksum := uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24

													n := chunkLen - checksumSize
													if n > len(r.decoded) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:189
				_go_fuzz_dep_.CoverTab[82116]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:191
				// _ = "end of CoverTab[82116]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:192
				_go_fuzz_dep_.CoverTab[82117]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:192
				// _ = "end of CoverTab[82117]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:192
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:192
			// _ = "end of CoverTab[82091]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:192
			_go_fuzz_dep_.CoverTab[82092]++
													if !r.readFull(r.decoded[:n], false) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:193
				_go_fuzz_dep_.CoverTab[82118]++
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:194
				// _ = "end of CoverTab[82118]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:195
				_go_fuzz_dep_.CoverTab[82119]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:195
				// _ = "end of CoverTab[82119]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:195
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:195
			// _ = "end of CoverTab[82092]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:195
			_go_fuzz_dep_.CoverTab[82093]++
													if crc(r.decoded[:n]) != checksum {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:196
				_go_fuzz_dep_.CoverTab[82120]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:198
				// _ = "end of CoverTab[82120]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:199
				_go_fuzz_dep_.CoverTab[82121]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:199
				// _ = "end of CoverTab[82121]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:199
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:199
			// _ = "end of CoverTab[82093]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:199
			_go_fuzz_dep_.CoverTab[82094]++
													r.i, r.j = 0, n
													continue
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:201
			// _ = "end of CoverTab[82094]"

		case chunkTypeStreamIdentifier:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:203
			_go_fuzz_dep_.CoverTab[82095]++

													if chunkLen != len(magicBody) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:205
				_go_fuzz_dep_.CoverTab[82122]++
														r.err = ErrCorrupt
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:207
				// _ = "end of CoverTab[82122]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:208
				_go_fuzz_dep_.CoverTab[82123]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:208
				// _ = "end of CoverTab[82123]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:208
			// _ = "end of CoverTab[82095]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:208
			_go_fuzz_dep_.CoverTab[82096]++
													if !r.readFull(r.buf[:len(magicBody)], false) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:209
				_go_fuzz_dep_.CoverTab[82124]++
														return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:210
				// _ = "end of CoverTab[82124]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:211
				_go_fuzz_dep_.CoverTab[82125]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:211
				// _ = "end of CoverTab[82125]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:211
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:211
			// _ = "end of CoverTab[82096]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:211
			_go_fuzz_dep_.CoverTab[82097]++
													for i := 0; i < len(magicBody); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:212
				_go_fuzz_dep_.CoverTab[82126]++
														if r.buf[i] != magicBody[i] {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:213
					_go_fuzz_dep_.CoverTab[82127]++
															r.err = ErrCorrupt
															return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:215
					// _ = "end of CoverTab[82127]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:216
					_go_fuzz_dep_.CoverTab[82128]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:216
					// _ = "end of CoverTab[82128]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:216
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:216
				// _ = "end of CoverTab[82126]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:217
			// _ = "end of CoverTab[82097]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:217
			_go_fuzz_dep_.CoverTab[82098]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:218
			// _ = "end of CoverTab[82098]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:218
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:218
			_go_fuzz_dep_.CoverTab[82099]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:218
			// _ = "end of CoverTab[82099]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:219
		// _ = "end of CoverTab[82070]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:219
		_go_fuzz_dep_.CoverTab[82071]++

												if chunkType <= 0x7f {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:221
			_go_fuzz_dep_.CoverTab[82129]++

													r.err = ErrUnsupported
													return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:224
			// _ = "end of CoverTab[82129]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:225
			_go_fuzz_dep_.CoverTab[82130]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:225
			// _ = "end of CoverTab[82130]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:225
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:225
		// _ = "end of CoverTab[82071]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:225
		_go_fuzz_dep_.CoverTab[82072]++

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:228
		if !r.readFull(r.buf[:chunkLen], false) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:228
			_go_fuzz_dep_.CoverTab[82131]++
													return r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:229
			// _ = "end of CoverTab[82131]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:230
			_go_fuzz_dep_.CoverTab[82132]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:230
			// _ = "end of CoverTab[82132]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:230
		// _ = "end of CoverTab[82072]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:231
	// _ = "end of CoverTab[82065]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:231
	_go_fuzz_dep_.CoverTab[82066]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:233
	// _ = "end of CoverTab[82066]"
}

// Read satisfies the io.Reader interface.
func (r *Reader) Read(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:237
	_go_fuzz_dep_.CoverTab[82133]++
											if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:238
		_go_fuzz_dep_.CoverTab[82136]++
												return 0, r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:239
		// _ = "end of CoverTab[82136]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:240
		_go_fuzz_dep_.CoverTab[82137]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:240
		// _ = "end of CoverTab[82137]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:240
	// _ = "end of CoverTab[82133]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:240
	_go_fuzz_dep_.CoverTab[82134]++

											if err := r.fill(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:242
		_go_fuzz_dep_.CoverTab[82138]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:243
		// _ = "end of CoverTab[82138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:244
		_go_fuzz_dep_.CoverTab[82139]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:244
		// _ = "end of CoverTab[82139]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:244
	// _ = "end of CoverTab[82134]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:244
	_go_fuzz_dep_.CoverTab[82135]++

											n := copy(p, r.decoded[r.i:r.j])
											r.i += n
											return n, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:248
	// _ = "end of CoverTab[82135]"
}

// ReadByte satisfies the io.ByteReader interface.
func (r *Reader) ReadByte() (byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:252
	_go_fuzz_dep_.CoverTab[82140]++
											if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:253
		_go_fuzz_dep_.CoverTab[82143]++
												return 0, r.err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:254
		// _ = "end of CoverTab[82143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:255
		_go_fuzz_dep_.CoverTab[82144]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:255
		// _ = "end of CoverTab[82144]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:255
	// _ = "end of CoverTab[82140]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:255
	_go_fuzz_dep_.CoverTab[82141]++

											if err := r.fill(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:257
		_go_fuzz_dep_.CoverTab[82145]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:258
		// _ = "end of CoverTab[82145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:259
		_go_fuzz_dep_.CoverTab[82146]++
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:259
		// _ = "end of CoverTab[82146]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:259
	// _ = "end of CoverTab[82141]"
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:259
	_go_fuzz_dep_.CoverTab[82142]++

											c := r.decoded[r.i]
											r.i++
											return c, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:263
	// _ = "end of CoverTab[82142]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:264
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/snappy@v0.0.4/decode.go:264
var _ = _go_fuzz_dep_.CoverTab
