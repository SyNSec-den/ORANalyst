// Copyright 2011 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
package snapref

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:5
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:25
	_go_fuzz_dep_.CoverTab[90367]++
													v, _, err := decodedLen(src)
													return v, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:27
	// _ = "end of CoverTab[90367]"
}

// decodedLen returns the length of the decoded block and the number of bytes
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:30
// that the length header occupied.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:32
func decodedLen(src []byte) (blockLen, headerLen int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:32
	_go_fuzz_dep_.CoverTab[90368]++
													v, n := binary.Uvarint(src)
													if n <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:34
		_go_fuzz_dep_.CoverTab[90371]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:34
		return v > 0xffffffff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:34
		// _ = "end of CoverTab[90371]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:34
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:34
		_go_fuzz_dep_.CoverTab[90372]++
														return 0, 0, ErrCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:35
		// _ = "end of CoverTab[90372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:36
		_go_fuzz_dep_.CoverTab[90373]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:36
		// _ = "end of CoverTab[90373]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:36
	// _ = "end of CoverTab[90368]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:36
	_go_fuzz_dep_.CoverTab[90369]++

													const wordSize = 32 << (^uint(0) >> 32 & 1)
													if wordSize == 32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:39
		_go_fuzz_dep_.CoverTab[90374]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:39
		return v > 0x7fffffff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:39
		// _ = "end of CoverTab[90374]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:39
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:39
		_go_fuzz_dep_.CoverTab[90375]++
														return 0, 0, ErrTooLarge
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:40
		// _ = "end of CoverTab[90375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:41
		_go_fuzz_dep_.CoverTab[90376]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:41
		// _ = "end of CoverTab[90376]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:41
	// _ = "end of CoverTab[90369]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:41
	_go_fuzz_dep_.CoverTab[90370]++
													return int(v), n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:42
	// _ = "end of CoverTab[90370]"
}

const (
	decodeErrCodeCorrupt			= 1
	decodeErrCodeUnsupportedLiteralLength	= 2
)

// Decode returns the decoded form of src. The returned slice may be a sub-
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:50
// slice of dst if dst was large enough to hold the entire decoded block.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:50
// Otherwise, a newly allocated slice will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:50
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:50
// The dst and src must not overlap. It is valid to pass a nil dst.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:50
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:50
// Decode handles the Snappy block format, not the Snappy stream format.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:57
func Decode(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:57
	_go_fuzz_dep_.CoverTab[90377]++
													dLen, s, err := decodedLen(src)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:59
		_go_fuzz_dep_.CoverTab[90381]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:60
		// _ = "end of CoverTab[90381]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:61
		_go_fuzz_dep_.CoverTab[90382]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:61
		// _ = "end of CoverTab[90382]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:61
	// _ = "end of CoverTab[90377]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:61
	_go_fuzz_dep_.CoverTab[90378]++
													if dLen <= len(dst) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:62
		_go_fuzz_dep_.CoverTab[90383]++
														dst = dst[:dLen]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:63
		// _ = "end of CoverTab[90383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:64
		_go_fuzz_dep_.CoverTab[90384]++
														dst = make([]byte, dLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:65
		// _ = "end of CoverTab[90384]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:66
	// _ = "end of CoverTab[90378]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:66
	_go_fuzz_dep_.CoverTab[90379]++
													switch decode(dst, src[s:]) {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:68
		_go_fuzz_dep_.CoverTab[90385]++
														return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:69
		// _ = "end of CoverTab[90385]"
	case decodeErrCodeUnsupportedLiteralLength:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:70
		_go_fuzz_dep_.CoverTab[90386]++
														return nil, errUnsupportedLiteralLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:71
		// _ = "end of CoverTab[90386]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:71
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:71
		_go_fuzz_dep_.CoverTab[90387]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:71
		// _ = "end of CoverTab[90387]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:72
	// _ = "end of CoverTab[90379]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:72
	_go_fuzz_dep_.CoverTab[90380]++
													return nil, ErrCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:73
	// _ = "end of CoverTab[90380]"
}

// NewReader returns a new Reader that decompresses from r, using the framing
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:76
// format described at
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:76
// https://github.com/google/snappy/blob/master/framing_format.txt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:79
func NewReader(r io.Reader) *Reader {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:79
	_go_fuzz_dep_.CoverTab[90388]++
													return &Reader{
		r:		r,
		decoded:	make([]byte, maxBlockSize),
		buf:		make([]byte, maxEncodedLenOfMaxBlockSize+checksumSize),
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:84
	// _ = "end of CoverTab[90388]"
}

// Reader is an io.Reader that can read Snappy-compressed bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:87
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:87
// Reader handles the Snappy stream format, not the Snappy block format.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:90
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:100
// reader to read from r. This permits reusing a Reader rather than allocating
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:100
// a new one.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:103
func (r *Reader) Reset(reader io.Reader) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:103
	_go_fuzz_dep_.CoverTab[90389]++
													r.r = reader
													r.err = nil
													r.i = 0
													r.j = 0
													r.readHeader = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:108
	// _ = "end of CoverTab[90389]"
}

func (r *Reader) readFull(p []byte, allowEOF bool) (ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:111
	_go_fuzz_dep_.CoverTab[90390]++
													if _, r.err = io.ReadFull(r.r, p); r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:112
		_go_fuzz_dep_.CoverTab[90392]++
														if r.err == io.ErrUnexpectedEOF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
			_go_fuzz_dep_.CoverTab[90394]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
			return (r.err == io.EOF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
				_go_fuzz_dep_.CoverTab[90395]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
				return !allowEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
				// _ = "end of CoverTab[90395]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
			}())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
			// _ = "end of CoverTab[90394]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:113
			_go_fuzz_dep_.CoverTab[90396]++
															r.err = ErrCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:114
			// _ = "end of CoverTab[90396]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:115
			_go_fuzz_dep_.CoverTab[90397]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:115
			// _ = "end of CoverTab[90397]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:115
		// _ = "end of CoverTab[90392]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:115
		_go_fuzz_dep_.CoverTab[90393]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:116
		// _ = "end of CoverTab[90393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:117
		_go_fuzz_dep_.CoverTab[90398]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:117
		// _ = "end of CoverTab[90398]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:117
	// _ = "end of CoverTab[90390]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:117
	_go_fuzz_dep_.CoverTab[90391]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:118
	// _ = "end of CoverTab[90391]"
}

func (r *Reader) fill() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:121
	_go_fuzz_dep_.CoverTab[90399]++
													for r.i >= r.j {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:122
		_go_fuzz_dep_.CoverTab[90401]++
														if !r.readFull(r.buf[:4], true) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:123
			_go_fuzz_dep_.CoverTab[90407]++
															return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:124
			// _ = "end of CoverTab[90407]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:125
			_go_fuzz_dep_.CoverTab[90408]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:125
			// _ = "end of CoverTab[90408]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:125
		// _ = "end of CoverTab[90401]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:125
		_go_fuzz_dep_.CoverTab[90402]++
														chunkType := r.buf[0]
														if !r.readHeader {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:127
			_go_fuzz_dep_.CoverTab[90409]++
															if chunkType != chunkTypeStreamIdentifier {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:128
				_go_fuzz_dep_.CoverTab[90411]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:130
				// _ = "end of CoverTab[90411]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:131
				_go_fuzz_dep_.CoverTab[90412]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:131
				// _ = "end of CoverTab[90412]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:131
			// _ = "end of CoverTab[90409]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:131
			_go_fuzz_dep_.CoverTab[90410]++
															r.readHeader = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:132
			// _ = "end of CoverTab[90410]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:133
			_go_fuzz_dep_.CoverTab[90413]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:133
			// _ = "end of CoverTab[90413]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:133
		// _ = "end of CoverTab[90402]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:133
		_go_fuzz_dep_.CoverTab[90403]++
														chunkLen := int(r.buf[1]) | int(r.buf[2])<<8 | int(r.buf[3])<<16
														if chunkLen > len(r.buf) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:135
			_go_fuzz_dep_.CoverTab[90414]++
															r.err = ErrUnsupported
															return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:137
			// _ = "end of CoverTab[90414]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:138
			_go_fuzz_dep_.CoverTab[90415]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:138
			// _ = "end of CoverTab[90415]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:138
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:138
		// _ = "end of CoverTab[90403]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:138
		_go_fuzz_dep_.CoverTab[90404]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:142
		switch chunkType {
		case chunkTypeCompressedData:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:143
			_go_fuzz_dep_.CoverTab[90416]++

															if chunkLen < checksumSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:145
				_go_fuzz_dep_.CoverTab[90434]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:147
				// _ = "end of CoverTab[90434]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:148
				_go_fuzz_dep_.CoverTab[90435]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:148
				// _ = "end of CoverTab[90435]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:148
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:148
			// _ = "end of CoverTab[90416]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:148
			_go_fuzz_dep_.CoverTab[90417]++
															buf := r.buf[:chunkLen]
															if !r.readFull(buf, false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:150
				_go_fuzz_dep_.CoverTab[90436]++
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:151
				// _ = "end of CoverTab[90436]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:152
				_go_fuzz_dep_.CoverTab[90437]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:152
				// _ = "end of CoverTab[90437]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:152
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:152
			// _ = "end of CoverTab[90417]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:152
			_go_fuzz_dep_.CoverTab[90418]++
															checksum := uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24
															buf = buf[checksumSize:]

															n, err := DecodedLen(buf)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:157
				_go_fuzz_dep_.CoverTab[90438]++
																r.err = err
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:159
				// _ = "end of CoverTab[90438]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:160
				_go_fuzz_dep_.CoverTab[90439]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:160
				// _ = "end of CoverTab[90439]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:160
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:160
			// _ = "end of CoverTab[90418]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:160
			_go_fuzz_dep_.CoverTab[90419]++
															if n > len(r.decoded) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:161
				_go_fuzz_dep_.CoverTab[90440]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:163
				// _ = "end of CoverTab[90440]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:164
				_go_fuzz_dep_.CoverTab[90441]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:164
				// _ = "end of CoverTab[90441]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:164
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:164
			// _ = "end of CoverTab[90419]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:164
			_go_fuzz_dep_.CoverTab[90420]++
															if _, err := Decode(r.decoded, buf); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:165
				_go_fuzz_dep_.CoverTab[90442]++
																r.err = err
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:167
				// _ = "end of CoverTab[90442]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:168
				_go_fuzz_dep_.CoverTab[90443]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:168
				// _ = "end of CoverTab[90443]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:168
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:168
			// _ = "end of CoverTab[90420]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:168
			_go_fuzz_dep_.CoverTab[90421]++
															if crc(r.decoded[:n]) != checksum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:169
				_go_fuzz_dep_.CoverTab[90444]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:171
				// _ = "end of CoverTab[90444]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:172
				_go_fuzz_dep_.CoverTab[90445]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:172
				// _ = "end of CoverTab[90445]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:172
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:172
			// _ = "end of CoverTab[90421]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:172
			_go_fuzz_dep_.CoverTab[90422]++
															r.i, r.j = 0, n
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:174
			// _ = "end of CoverTab[90422]"

		case chunkTypeUncompressedData:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:176
			_go_fuzz_dep_.CoverTab[90423]++

															if chunkLen < checksumSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:178
				_go_fuzz_dep_.CoverTab[90446]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:180
				// _ = "end of CoverTab[90446]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:181
				_go_fuzz_dep_.CoverTab[90447]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:181
				// _ = "end of CoverTab[90447]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:181
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:181
			// _ = "end of CoverTab[90423]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:181
			_go_fuzz_dep_.CoverTab[90424]++
															buf := r.buf[:checksumSize]
															if !r.readFull(buf, false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:183
				_go_fuzz_dep_.CoverTab[90448]++
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:184
				// _ = "end of CoverTab[90448]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:185
				_go_fuzz_dep_.CoverTab[90449]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:185
				// _ = "end of CoverTab[90449]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:185
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:185
			// _ = "end of CoverTab[90424]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:185
			_go_fuzz_dep_.CoverTab[90425]++
															checksum := uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24

															n := chunkLen - checksumSize
															if n > len(r.decoded) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:189
				_go_fuzz_dep_.CoverTab[90450]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:191
				// _ = "end of CoverTab[90450]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:192
				_go_fuzz_dep_.CoverTab[90451]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:192
				// _ = "end of CoverTab[90451]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:192
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:192
			// _ = "end of CoverTab[90425]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:192
			_go_fuzz_dep_.CoverTab[90426]++
															if !r.readFull(r.decoded[:n], false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:193
				_go_fuzz_dep_.CoverTab[90452]++
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:194
				// _ = "end of CoverTab[90452]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:195
				_go_fuzz_dep_.CoverTab[90453]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:195
				// _ = "end of CoverTab[90453]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:195
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:195
			// _ = "end of CoverTab[90426]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:195
			_go_fuzz_dep_.CoverTab[90427]++
															if crc(r.decoded[:n]) != checksum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:196
				_go_fuzz_dep_.CoverTab[90454]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:198
				// _ = "end of CoverTab[90454]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:199
				_go_fuzz_dep_.CoverTab[90455]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:199
				// _ = "end of CoverTab[90455]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:199
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:199
			// _ = "end of CoverTab[90427]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:199
			_go_fuzz_dep_.CoverTab[90428]++
															r.i, r.j = 0, n
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:201
			// _ = "end of CoverTab[90428]"

		case chunkTypeStreamIdentifier:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:203
			_go_fuzz_dep_.CoverTab[90429]++

															if chunkLen != len(magicBody) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:205
				_go_fuzz_dep_.CoverTab[90456]++
																r.err = ErrCorrupt
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:207
				// _ = "end of CoverTab[90456]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:208
				_go_fuzz_dep_.CoverTab[90457]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:208
				// _ = "end of CoverTab[90457]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:208
			// _ = "end of CoverTab[90429]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:208
			_go_fuzz_dep_.CoverTab[90430]++
															if !r.readFull(r.buf[:len(magicBody)], false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:209
				_go_fuzz_dep_.CoverTab[90458]++
																return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:210
				// _ = "end of CoverTab[90458]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:211
				_go_fuzz_dep_.CoverTab[90459]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:211
				// _ = "end of CoverTab[90459]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:211
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:211
			// _ = "end of CoverTab[90430]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:211
			_go_fuzz_dep_.CoverTab[90431]++
															for i := 0; i < len(magicBody); i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:212
				_go_fuzz_dep_.CoverTab[90460]++
																if r.buf[i] != magicBody[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:213
					_go_fuzz_dep_.CoverTab[90461]++
																	r.err = ErrCorrupt
																	return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:215
					// _ = "end of CoverTab[90461]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:216
					_go_fuzz_dep_.CoverTab[90462]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:216
					// _ = "end of CoverTab[90462]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:216
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:216
				// _ = "end of CoverTab[90460]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:217
			// _ = "end of CoverTab[90431]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:217
			_go_fuzz_dep_.CoverTab[90432]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:218
			// _ = "end of CoverTab[90432]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:218
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:218
			_go_fuzz_dep_.CoverTab[90433]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:218
			// _ = "end of CoverTab[90433]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:219
		// _ = "end of CoverTab[90404]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:219
		_go_fuzz_dep_.CoverTab[90405]++

														if chunkType <= 0x7f {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:221
			_go_fuzz_dep_.CoverTab[90463]++

															r.err = ErrUnsupported
															return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:224
			// _ = "end of CoverTab[90463]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:225
			_go_fuzz_dep_.CoverTab[90464]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:225
			// _ = "end of CoverTab[90464]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:225
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:225
		// _ = "end of CoverTab[90405]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:225
		_go_fuzz_dep_.CoverTab[90406]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:228
		if !r.readFull(r.buf[:chunkLen], false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:228
			_go_fuzz_dep_.CoverTab[90465]++
															return r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:229
			// _ = "end of CoverTab[90465]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:230
			_go_fuzz_dep_.CoverTab[90466]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:230
			// _ = "end of CoverTab[90466]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:230
		// _ = "end of CoverTab[90406]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:231
	// _ = "end of CoverTab[90399]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:231
	_go_fuzz_dep_.CoverTab[90400]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:233
	// _ = "end of CoverTab[90400]"
}

// Read satisfies the io.Reader interface.
func (r *Reader) Read(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:237
	_go_fuzz_dep_.CoverTab[90467]++
													if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:238
		_go_fuzz_dep_.CoverTab[90470]++
														return 0, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:239
		// _ = "end of CoverTab[90470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:240
		_go_fuzz_dep_.CoverTab[90471]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:240
		// _ = "end of CoverTab[90471]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:240
	// _ = "end of CoverTab[90467]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:240
	_go_fuzz_dep_.CoverTab[90468]++

													if err := r.fill(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:242
		_go_fuzz_dep_.CoverTab[90472]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:243
		// _ = "end of CoverTab[90472]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:244
		_go_fuzz_dep_.CoverTab[90473]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:244
		// _ = "end of CoverTab[90473]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:244
	// _ = "end of CoverTab[90468]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:244
	_go_fuzz_dep_.CoverTab[90469]++

													n := copy(p, r.decoded[r.i:r.j])
													r.i += n
													return n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:248
	// _ = "end of CoverTab[90469]"
}

// ReadByte satisfies the io.ByteReader interface.
func (r *Reader) ReadByte() (byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:252
	_go_fuzz_dep_.CoverTab[90474]++
													if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:253
		_go_fuzz_dep_.CoverTab[90477]++
														return 0, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:254
		// _ = "end of CoverTab[90477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:255
		_go_fuzz_dep_.CoverTab[90478]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:255
		// _ = "end of CoverTab[90478]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:255
	// _ = "end of CoverTab[90474]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:255
	_go_fuzz_dep_.CoverTab[90475]++

													if err := r.fill(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:257
		_go_fuzz_dep_.CoverTab[90479]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:258
		// _ = "end of CoverTab[90479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:259
		_go_fuzz_dep_.CoverTab[90480]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:259
		// _ = "end of CoverTab[90480]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:259
	// _ = "end of CoverTab[90475]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:259
	_go_fuzz_dep_.CoverTab[90476]++

													c := r.decoded[r.i]
													r.i++
													return c, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:263
	// _ = "end of CoverTab[90476]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:264
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode.go:264
var _ = _go_fuzz_dep_.CoverTab
