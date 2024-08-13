// Copyright 2020+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:4
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// HeaderMaxSize is the maximum size of a Frame and Block Header.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:13
// If less is sent to Header.Decode it *may* still contain enough information.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:15
const HeaderMaxSize = 14 + 3

// Header contains information about the first frame and block within that.
type Header struct {
	// SingleSegment specifies whether the data is to be decompressed into a
	// single contiguous memory segment.
	// It implies that WindowSize is invalid and that FrameContentSize is valid.
	SingleSegment	bool

	// WindowSize is the window of data to keep while decoding.
	// Will only be set if SingleSegment is false.
	WindowSize	uint64

	// Dictionary ID.
	// If 0, no dictionary.
	DictionaryID	uint32

	// HasFCS specifies whether FrameContentSize has a valid value.
	HasFCS	bool

	// FrameContentSize is the expected uncompressed size of the entire frame.
	FrameContentSize	uint64

	// Skippable will be true if the frame is meant to be skipped.
	// This implies that FirstBlock.OK is false.
	Skippable	bool

	// SkippableID is the user-specific ID for the skippable frame.
	// Valid values are between 0 to 15, inclusive.
	SkippableID	int

	// SkippableSize is the length of the user data to skip following
	// the header.
	SkippableSize	uint32

	// HeaderSize is the raw size of the frame header.
	//
	// For normal frames, it includes the size of the magic number and
	// the size of the header (per section 3.1.1.1).
	// It does not include the size for any data blocks (section 3.1.1.2) nor
	// the size for the trailing content checksum.
	//
	// For skippable frames, this counts the size of the magic number
	// along with the size of the size field of the payload.
	// It does not include the size of the skippable payload itself.
	// The total frame size is the HeaderSize plus the SkippableSize.
	HeaderSize	int

	// First block information.
	FirstBlock	struct {
		// OK will be set if first block could be decoded.
		OK	bool

		// Is this the last block of a frame?
		Last	bool

		// Is the data compressed?
		// If true CompressedSize will be populated.
		// Unfortunately DecompressedSize cannot be determined
		// without decoding the blocks.
		Compressed	bool

		// DecompressedSize is the expected decompressed size of the block.
		// Will be 0 if it cannot be determined.
		DecompressedSize	int

		// CompressedSize of the data in the block.
		// Does not include the block header.
		// Will be equal to DecompressedSize if not Compressed.
		CompressedSize	int
	}

	// If set there is a checksum present for the block content.
	// The checksum field at the end is always 4 bytes long.
	HasCheckSum	bool
}

// Decode the header from the beginning of the stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:92
// This will decode the frame header and the first block header if enough bytes are provided.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:92
// It is recommended to provide at least HeaderMaxSize bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:92
// If the frame header cannot be read an error will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:92
// If there isn't enough input, io.ErrUnexpectedEOF is returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:92
// The FirstBlock.OK will indicate if enough information was available to decode the first block header.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:98
func (h *Header) Decode(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:98
	_go_fuzz_dep_.CoverTab[91620]++
													*h = Header{}
													if len(in) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:100
		_go_fuzz_dep_.CoverTab[91631]++
														return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:101
		// _ = "end of CoverTab[91631]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:102
		_go_fuzz_dep_.CoverTab[91632]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:102
		// _ = "end of CoverTab[91632]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:102
	// _ = "end of CoverTab[91620]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:102
	_go_fuzz_dep_.CoverTab[91621]++
													h.HeaderSize += 4
													b, in := in[:4], in[4:]
													if !bytes.Equal(b, frameMagic) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:105
		_go_fuzz_dep_.CoverTab[91633]++
														if !bytes.Equal(b[1:4], skippableFrameMagic) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:106
			_go_fuzz_dep_.CoverTab[91636]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:106
			return b[0]&0xf0 != 0x50
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:106
			// _ = "end of CoverTab[91636]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:106
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:106
			_go_fuzz_dep_.CoverTab[91637]++
															return ErrMagicMismatch
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:107
			// _ = "end of CoverTab[91637]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:108
			_go_fuzz_dep_.CoverTab[91638]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:108
			// _ = "end of CoverTab[91638]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:108
		// _ = "end of CoverTab[91633]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:108
		_go_fuzz_dep_.CoverTab[91634]++
														if len(in) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:109
			_go_fuzz_dep_.CoverTab[91639]++
															return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:110
			// _ = "end of CoverTab[91639]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:111
			_go_fuzz_dep_.CoverTab[91640]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:111
			// _ = "end of CoverTab[91640]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:111
		// _ = "end of CoverTab[91634]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:111
		_go_fuzz_dep_.CoverTab[91635]++
														h.HeaderSize += 4
														h.Skippable = true
														h.SkippableID = int(b[0] & 0xf)
														h.SkippableSize = binary.LittleEndian.Uint32(in)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:116
		// _ = "end of CoverTab[91635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:117
		_go_fuzz_dep_.CoverTab[91641]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:117
		// _ = "end of CoverTab[91641]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:117
	// _ = "end of CoverTab[91621]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:117
	_go_fuzz_dep_.CoverTab[91622]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:121
	if len(in) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:121
		_go_fuzz_dep_.CoverTab[91642]++
														return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:122
		// _ = "end of CoverTab[91642]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:123
		_go_fuzz_dep_.CoverTab[91643]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:123
		// _ = "end of CoverTab[91643]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:123
	// _ = "end of CoverTab[91622]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:123
	_go_fuzz_dep_.CoverTab[91623]++
													fhd, in := in[0], in[1:]
													h.HeaderSize++
													h.SingleSegment = fhd&(1<<5) != 0
													h.HasCheckSum = fhd&(1<<2) != 0
													if fhd&(1<<3) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:128
		_go_fuzz_dep_.CoverTab[91644]++
														return errors.New("reserved bit set on frame header")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:129
		// _ = "end of CoverTab[91644]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:130
		_go_fuzz_dep_.CoverTab[91645]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:130
		// _ = "end of CoverTab[91645]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:130
	// _ = "end of CoverTab[91623]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:130
	_go_fuzz_dep_.CoverTab[91624]++

													if !h.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:132
		_go_fuzz_dep_.CoverTab[91646]++
														if len(in) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:133
			_go_fuzz_dep_.CoverTab[91648]++
															return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:134
			// _ = "end of CoverTab[91648]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:135
			_go_fuzz_dep_.CoverTab[91649]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:135
			// _ = "end of CoverTab[91649]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:135
		// _ = "end of CoverTab[91646]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:135
		_go_fuzz_dep_.CoverTab[91647]++
														var wd byte
														wd, in = in[0], in[1:]
														h.HeaderSize++
														windowLog := 10 + (wd >> 3)
														windowBase := uint64(1) << windowLog
														windowAdd := (windowBase / 8) * uint64(wd&0x7)
														h.WindowSize = windowBase + windowAdd
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:142
		// _ = "end of CoverTab[91647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:143
		_go_fuzz_dep_.CoverTab[91650]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:143
		// _ = "end of CoverTab[91650]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:143
	// _ = "end of CoverTab[91624]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:143
	_go_fuzz_dep_.CoverTab[91625]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:147
	if size := fhd & 3; size != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:147
		_go_fuzz_dep_.CoverTab[91651]++
														if size == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:148
			_go_fuzz_dep_.CoverTab[91654]++
															size = 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:149
			// _ = "end of CoverTab[91654]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:150
			_go_fuzz_dep_.CoverTab[91655]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:150
			// _ = "end of CoverTab[91655]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:150
		// _ = "end of CoverTab[91651]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:150
		_go_fuzz_dep_.CoverTab[91652]++
														if len(in) < int(size) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:151
			_go_fuzz_dep_.CoverTab[91656]++
															return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:152
			// _ = "end of CoverTab[91656]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:153
			_go_fuzz_dep_.CoverTab[91657]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:153
			// _ = "end of CoverTab[91657]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:153
		// _ = "end of CoverTab[91652]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:153
		_go_fuzz_dep_.CoverTab[91653]++
														b, in = in[:size], in[size:]
														h.HeaderSize += int(size)
														switch size {
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:157
			_go_fuzz_dep_.CoverTab[91658]++
															h.DictionaryID = uint32(b[0])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:158
			// _ = "end of CoverTab[91658]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:159
			_go_fuzz_dep_.CoverTab[91659]++
															h.DictionaryID = uint32(b[0]) | (uint32(b[1]) << 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:160
			// _ = "end of CoverTab[91659]"
		case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:161
			_go_fuzz_dep_.CoverTab[91660]++
															h.DictionaryID = uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:162
			// _ = "end of CoverTab[91660]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:162
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:162
			_go_fuzz_dep_.CoverTab[91661]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:162
			// _ = "end of CoverTab[91661]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:163
		// _ = "end of CoverTab[91653]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:164
		_go_fuzz_dep_.CoverTab[91662]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:164
		// _ = "end of CoverTab[91662]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:164
	// _ = "end of CoverTab[91625]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:164
	_go_fuzz_dep_.CoverTab[91626]++

	// Read Frame_Content_Size
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#frame_content_size
	var fcsSize int
	v := fhd >> 6
	switch v {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:171
		_go_fuzz_dep_.CoverTab[91663]++
														if h.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:172
			_go_fuzz_dep_.CoverTab[91665]++
															fcsSize = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:173
			// _ = "end of CoverTab[91665]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:174
			_go_fuzz_dep_.CoverTab[91666]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:174
			// _ = "end of CoverTab[91666]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:174
		// _ = "end of CoverTab[91663]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:175
		_go_fuzz_dep_.CoverTab[91664]++
														fcsSize = 1 << v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:176
		// _ = "end of CoverTab[91664]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:177
	// _ = "end of CoverTab[91626]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:177
	_go_fuzz_dep_.CoverTab[91627]++

													if fcsSize > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:179
		_go_fuzz_dep_.CoverTab[91667]++
														h.HasFCS = true
														if len(in) < fcsSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:181
			_go_fuzz_dep_.CoverTab[91669]++
															return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:182
			// _ = "end of CoverTab[91669]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:183
			_go_fuzz_dep_.CoverTab[91670]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:183
			// _ = "end of CoverTab[91670]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:183
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:183
		// _ = "end of CoverTab[91667]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:183
		_go_fuzz_dep_.CoverTab[91668]++
														b, in = in[:fcsSize], in[fcsSize:]
														h.HeaderSize += int(fcsSize)
														switch fcsSize {
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:187
			_go_fuzz_dep_.CoverTab[91671]++
															h.FrameContentSize = uint64(b[0])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:188
			// _ = "end of CoverTab[91671]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:189
			_go_fuzz_dep_.CoverTab[91672]++

															h.FrameContentSize = uint64(b[0]) | (uint64(b[1]) << 8) + 256
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:191
			// _ = "end of CoverTab[91672]"
		case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:192
			_go_fuzz_dep_.CoverTab[91673]++
															h.FrameContentSize = uint64(b[0]) | (uint64(b[1]) << 8) | (uint64(b[2]) << 16) | (uint64(b[3]) << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:193
			// _ = "end of CoverTab[91673]"
		case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:194
			_go_fuzz_dep_.CoverTab[91674]++
															d1 := uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
															d2 := uint32(b[4]) | (uint32(b[5]) << 8) | (uint32(b[6]) << 16) | (uint32(b[7]) << 24)
															h.FrameContentSize = uint64(d1) | (uint64(d2) << 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:197
			// _ = "end of CoverTab[91674]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:197
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:197
			_go_fuzz_dep_.CoverTab[91675]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:197
			// _ = "end of CoverTab[91675]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:198
		// _ = "end of CoverTab[91668]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:199
		_go_fuzz_dep_.CoverTab[91676]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:199
		// _ = "end of CoverTab[91676]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:199
	// _ = "end of CoverTab[91627]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:199
	_go_fuzz_dep_.CoverTab[91628]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:202
	if len(in) < 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:202
		_go_fuzz_dep_.CoverTab[91677]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:203
		// _ = "end of CoverTab[91677]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:204
		_go_fuzz_dep_.CoverTab[91678]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:204
		// _ = "end of CoverTab[91678]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:204
	// _ = "end of CoverTab[91628]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:204
	_go_fuzz_dep_.CoverTab[91629]++
													tmp := in[:3]
													bh := uint32(tmp[0]) | (uint32(tmp[1]) << 8) | (uint32(tmp[2]) << 16)
													h.FirstBlock.Last = bh&1 != 0
													blockType := blockType((bh >> 1) & 3)

													cSize := int(bh >> 3)
													switch blockType {
	case blockTypeReserved:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:212
		_go_fuzz_dep_.CoverTab[91679]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:213
		// _ = "end of CoverTab[91679]"
	case blockTypeRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:214
		_go_fuzz_dep_.CoverTab[91680]++
														h.FirstBlock.Compressed = true
														h.FirstBlock.DecompressedSize = cSize
														h.FirstBlock.CompressedSize = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:217
		// _ = "end of CoverTab[91680]"
	case blockTypeCompressed:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:218
		_go_fuzz_dep_.CoverTab[91681]++
														h.FirstBlock.Compressed = true
														h.FirstBlock.CompressedSize = cSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:220
		// _ = "end of CoverTab[91681]"
	case blockTypeRaw:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:221
		_go_fuzz_dep_.CoverTab[91682]++
														h.FirstBlock.DecompressedSize = cSize
														h.FirstBlock.CompressedSize = cSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:223
		// _ = "end of CoverTab[91682]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:224
		_go_fuzz_dep_.CoverTab[91683]++
														panic("Invalid block type")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:225
		// _ = "end of CoverTab[91683]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:226
	// _ = "end of CoverTab[91629]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:226
	_go_fuzz_dep_.CoverTab[91630]++

													h.FirstBlock.OK = true
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:229
	// _ = "end of CoverTab[91630]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:230
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decodeheader.go:230
var _ = _go_fuzz_dep_.CoverTab
