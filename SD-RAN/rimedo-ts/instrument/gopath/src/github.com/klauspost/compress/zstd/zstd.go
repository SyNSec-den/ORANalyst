//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:1
// Package zstd provides decompression of zstandard files.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:1
// For advanced usage and examples, go to the README: https://github.com/klauspost/compress/tree/master/zstd#zstd
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:4
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"math"
	"math/bits"
)

// enable debug printing
const debug = false

// enable encoding debug printing
const debugEncoder = debug

// enable decoding debug printing
const debugDecoder = debug

// Enable extra assertions.
const debugAsserts = debug || false

// print sequence details
const debugSequences = false

// print detailed matching information
const debugMatches = false

// force encoder to use predefined tables.
const forcePreDef = false

// zstdMinMatch is the minimum zstd match length.
const zstdMinMatch = 3

// Reset the buffer offset when reaching this.
const bufferReset = math.MaxInt32 - MaxWindowSize

var (
	// ErrReservedBlockType is returned when a reserved block type is found.
	// Typically this indicates wrong or corrupted input.
	ErrReservedBlockType	= errors.New("invalid input: reserved block type encountered")

	// ErrCompressedSizeTooBig is returned when a block is bigger than allowed.
	// Typically this indicates wrong or corrupted input.
	ErrCompressedSizeTooBig	= errors.New("invalid input: compressed size too big")

	// ErrBlockTooSmall is returned when a block is too small to be decoded.
	// Typically returned on invalid input.
	ErrBlockTooSmall	= errors.New("block too small")

	// ErrMagicMismatch is returned when a "magic" number isn't what is expected.
	// Typically this indicates wrong or corrupted input.
	ErrMagicMismatch	= errors.New("invalid input: magic number mismatch")

	// ErrWindowSizeExceeded is returned when a reference exceeds the valid window size.
	// Typically this indicates wrong or corrupted input.
	ErrWindowSizeExceeded	= errors.New("window size exceeded")

	// ErrWindowSizeTooSmall is returned when no window size is specified.
	// Typically this indicates wrong or corrupted input.
	ErrWindowSizeTooSmall	= errors.New("invalid input: window size was too small")

	// ErrDecoderSizeExceeded is returned if decompressed size exceeds the configured limit.
	ErrDecoderSizeExceeded	= errors.New("decompressed size exceeds configured limit")

	// ErrUnknownDictionary is returned if the dictionary ID is unknown.
	// For the time being dictionaries are not supported.
	ErrUnknownDictionary	= errors.New("unknown dictionary")

	// ErrFrameSizeExceeded is returned if the stated frame size is exceeded.
	// This is only returned if SingleSegment is specified on the frame.
	ErrFrameSizeExceeded	= errors.New("frame size exceeded")

	// ErrCRCMismatch is returned if CRC mismatches.
	ErrCRCMismatch	= errors.New("CRC check failed")

	// ErrDecoderClosed will be returned if the Decoder was used after
	// Close has been called.
	ErrDecoderClosed	= errors.New("decoder used after Close")

	// ErrDecoderNilInput is returned when a nil Reader was provided
	// and an operation other than Reset/DecodeAll/Close was attempted.
	ErrDecoderNilInput	= errors.New("nil input provided as reader")
)

func println(a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:90
	_go_fuzz_dep_.CoverTab[95274]++
												if debug || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		_go_fuzz_dep_.CoverTab[95275]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		return debugDecoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		// _ = "end of CoverTab[95275]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		_go_fuzz_dep_.CoverTab[95276]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		return debugEncoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		// _ = "end of CoverTab[95276]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:91
		_go_fuzz_dep_.CoverTab[95277]++
													log.Println(a...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:92
		// _ = "end of CoverTab[95277]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:93
		_go_fuzz_dep_.CoverTab[95278]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:93
		// _ = "end of CoverTab[95278]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:93
	// _ = "end of CoverTab[95274]"
}

func printf(format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:96
	_go_fuzz_dep_.CoverTab[95279]++
												if debug || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		_go_fuzz_dep_.CoverTab[95280]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		return debugDecoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		// _ = "end of CoverTab[95280]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		_go_fuzz_dep_.CoverTab[95281]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		return debugEncoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		// _ = "end of CoverTab[95281]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:97
		_go_fuzz_dep_.CoverTab[95282]++
													log.Printf(format, a...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:98
		// _ = "end of CoverTab[95282]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:99
		_go_fuzz_dep_.CoverTab[95283]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:99
		// _ = "end of CoverTab[95283]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:99
	// _ = "end of CoverTab[95279]"
}

// matchLenFast does matching, but will not match the last up to 7 bytes.
func matchLenFast(a, b []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:103
	_go_fuzz_dep_.CoverTab[95284]++
												endI := len(a) & (math.MaxInt32 - 7)
												for i := 0; i < endI; i += 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:105
		_go_fuzz_dep_.CoverTab[95286]++
													if diff := load64(a, i) ^ load64(b, i); diff != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:106
			_go_fuzz_dep_.CoverTab[95287]++
														return i + bits.TrailingZeros64(diff)>>3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:107
			// _ = "end of CoverTab[95287]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:108
			_go_fuzz_dep_.CoverTab[95288]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:108
			// _ = "end of CoverTab[95288]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:108
		// _ = "end of CoverTab[95286]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:109
	// _ = "end of CoverTab[95284]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:109
	_go_fuzz_dep_.CoverTab[95285]++
												return endI
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:110
	// _ = "end of CoverTab[95285]"
}

// matchLen returns the maximum length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:113
// a must be the shortest of the two.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:113
// The function also returns whether all bytes matched.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:116
func matchLen(a, b []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:116
	_go_fuzz_dep_.CoverTab[95289]++
												b = b[:len(a)]
												for i := 0; i < len(a)-7; i += 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:118
		_go_fuzz_dep_.CoverTab[95292]++
													if diff := load64(a, i) ^ load64(b, i); diff != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:119
			_go_fuzz_dep_.CoverTab[95293]++
														return i + (bits.TrailingZeros64(diff) >> 3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:120
			// _ = "end of CoverTab[95293]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:121
			_go_fuzz_dep_.CoverTab[95294]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:121
			// _ = "end of CoverTab[95294]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:121
		// _ = "end of CoverTab[95292]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:122
	// _ = "end of CoverTab[95289]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:122
	_go_fuzz_dep_.CoverTab[95290]++

												checked := (len(a) >> 3) << 3
												a = a[checked:]
												b = b[checked:]
												for i := range a {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:127
		_go_fuzz_dep_.CoverTab[95295]++
													if a[i] != b[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:128
			_go_fuzz_dep_.CoverTab[95296]++
														return i + checked
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:129
			// _ = "end of CoverTab[95296]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:130
			_go_fuzz_dep_.CoverTab[95297]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:130
			// _ = "end of CoverTab[95297]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:130
		// _ = "end of CoverTab[95295]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:131
	// _ = "end of CoverTab[95290]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:131
	_go_fuzz_dep_.CoverTab[95291]++
												return len(a) + checked
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:132
	// _ = "end of CoverTab[95291]"
}

func load3232(b []byte, i int32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:135
	_go_fuzz_dep_.CoverTab[95298]++
												return binary.LittleEndian.Uint32(b[i:])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:136
	// _ = "end of CoverTab[95298]"
}

func load6432(b []byte, i int32) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:139
	_go_fuzz_dep_.CoverTab[95299]++
												return binary.LittleEndian.Uint64(b[i:])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:140
	// _ = "end of CoverTab[95299]"
}

func load64(b []byte, i int) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:143
	_go_fuzz_dep_.CoverTab[95300]++
												return binary.LittleEndian.Uint64(b[i:])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:144
	// _ = "end of CoverTab[95300]"
}

type byter interface {
	Bytes() []byte
	Len() int
}

var _ byter = &bytes.Buffer{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:152
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zstd.go:152
var _ = _go_fuzz_dep_.CoverTab
