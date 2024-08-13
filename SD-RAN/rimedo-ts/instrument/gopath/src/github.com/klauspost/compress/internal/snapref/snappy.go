// Copyright 2011 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// Package snapref implements the Snappy compression format. It aims for very
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// high speeds and reasonable compression.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// There are actually two Snappy formats: block and stream. They are related,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// but different: trying to decompress block-compressed data as a Snappy stream
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// will fail, and vice versa. The block format is the Decode and Encode
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// functions and the stream format is the Reader and Writer types.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// The block format, the more common case, is used when the complete size (the
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// number of bytes) of the original data is known upfront, at the time
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// compression starts. The stream format, also known as the framing format, is
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// for when that isn't always true.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// The canonical, C++ implementation is at https://github.com/google/snappy and
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:5
// it only implements the block format.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
package snapref

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:20
)

import (
	"hash/crc32"
)

/*
Each encoded block begins with the varint-encoded length of the decoded data,
followed by a sequence of chunks. Chunks begin and end on byte boundaries. The
first byte of each chunk is broken into its 2 least and 6 most significant bits
called l and m: l ranges in [0, 4) and m ranges in [0, 64). l is the chunk tag.
Zero means a literal tag. All other values mean a copy tag.

For literal tags:
  - If m < 60, the next 1 + m bytes are literal bytes.
  - Otherwise, let n be the little-endian unsigned integer denoted by the next
    m - 59 bytes. The next 1 + n bytes after that are literal bytes.

For copy tags, length bytes are copied from offset bytes ago, in the style of
Lempel-Ziv compression algorithms. In particular:
  - For l == 1, the offset ranges in [0, 1<<11) and the length in [4, 12).
    The length is 4 + the low 3 bits of m. The high 3 bits of m form bits 8-10
    of the offset. The next byte is bits 0-7 of the offset.
  - For l == 2, the offset ranges in [0, 1<<16) and the length in [1, 65).
    The length is 1 + m. The offset is the little-endian unsigned integer
    denoted by the next 2 bytes.
  - For l == 3, this tag is a legacy format that is no longer issued by most
    encoders. Nonetheless, the offset ranges in [0, 1<<32) and the length in
    [1, 65). The length is 1 + m. The offset is the little-endian unsigned
    integer denoted by the next 4 bytes.
*/
const (
	tagLiteral	= 0x00
	tagCopy1	= 0x01
	tagCopy2	= 0x02
	tagCopy4	= 0x03
)

const (
	checksumSize	= 4
	chunkHeaderSize	= 4
	magicChunk	= "\xff\x06\x00\x00" + magicBody
	magicBody	= "sNaPpY"

	// maxBlockSize is the maximum size of the input to encodeBlock. It is not
	// part of the wire format per se, but some parts of the encoder assume
	// that an offset fits into a uint16.
	//
	// Also, for the framing format (Writer type instead of Encode function),
	// https://github.com/google/snappy/blob/master/framing_format.txt says
	// that "the uncompressed data in a chunk must be no longer than 65536
	// bytes".
	maxBlockSize	= 65536

	// maxEncodedLenOfMaxBlockSize equals MaxEncodedLen(maxBlockSize), but is
	// hard coded to be a const instead of a variable, so that obufLen can also
	// be a const. Their equivalence is confirmed by
	// TestMaxEncodedLenOfMaxBlockSize.
	maxEncodedLenOfMaxBlockSize	= 76490

	obufHeaderLen	= len(magicChunk) + checksumSize + chunkHeaderSize
	obufLen		= obufHeaderLen + maxEncodedLenOfMaxBlockSize
)

const (
	chunkTypeCompressedData		= 0x00
	chunkTypeUncompressedData	= 0x01
	chunkTypePadding		= 0xfe
	chunkTypeStreamIdentifier	= 0xff
)

var crcTable = crc32.MakeTable(crc32.Castagnoli)

// crc implements the checksum specified in section 3 of
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:93
// https://github.com/google/snappy/blob/master/framing_format.txt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:95
func crc(b []byte) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:95
	_go_fuzz_dep_.CoverTab[90658]++
													c := crc32.Update(0, crcTable, b)
													return uint32(c>>15|c<<17) + 0xa282ead8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:97
	// _ = "end of CoverTab[90658]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/snappy.go:98
var _ = _go_fuzz_dep_.CoverTab
