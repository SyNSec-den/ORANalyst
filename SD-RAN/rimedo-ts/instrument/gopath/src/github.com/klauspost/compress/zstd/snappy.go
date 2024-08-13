// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:5
)

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"

	"github.com/klauspost/compress/huff0"
	snappy "github.com/klauspost/compress/internal/snapref"
)

const (
	snappyTagLiteral	= 0x00
	snappyTagCopy1		= 0x01
	snappyTagCopy2		= 0x02
	snappyTagCopy4		= 0x03
)

const (
	snappyChecksumSize	= 4
	snappyMagicBody		= "sNaPpY"

	// snappyMaxBlockSize is the maximum size of the input to encodeBlock. It is not
	// part of the wire format per se, but some parts of the encoder assume
	// that an offset fits into a uint16.
	//
	// Also, for the framing format (Writer type instead of Encode function),
	// https://github.com/google/snappy/blob/master/framing_format.txt says
	// that "the uncompressed data in a chunk must be no longer than 65536
	// bytes".
	snappyMaxBlockSize	= 65536

	// snappyMaxEncodedLenOfMaxBlockSize equals MaxEncodedLen(snappyMaxBlockSize), but is
	// hard coded to be a const instead of a variable, so that obufLen can also
	// be a const. Their equivalence is confirmed by
	// TestMaxEncodedLenOfMaxBlockSize.
	snappyMaxEncodedLenOfMaxBlockSize	= 76490
)

const (
	chunkTypeCompressedData		= 0x00
	chunkTypeUncompressedData	= 0x01
	chunkTypePadding		= 0xfe
	chunkTypeStreamIdentifier	= 0xff
)

var (
	// ErrSnappyCorrupt reports that the input is invalid.
	ErrSnappyCorrupt	= errors.New("snappy: corrupt input")
	// ErrSnappyTooLarge reports that the uncompressed length is too large.
	ErrSnappyTooLarge	= errors.New("snappy: decoded block is too large")
	// ErrSnappyUnsupported reports that the input isn't supported.
	ErrSnappyUnsupported	= errors.New("snappy: unsupported input")

	errUnsupportedLiteralLength	= errors.New("snappy: unsupported literal length")
)

// SnappyConverter can read SnappyConverter-compressed streams and convert them to zstd.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// Conversion is done by converting the stream directly from Snappy without intermediate
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// full decoding.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// Therefore the compression ratio is much less than what can be done by a full decompression
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// and compression, and a faulty Snappy stream may lead to a faulty Zstandard stream without
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// any errors being generated.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// No CRC value is being generated and not all CRC values of the Snappy stream are checked.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// However, it provides really fast recompression of Snappy streams.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:63
// The converter can be reused to avoid allocations, even after errors.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:72
type SnappyConverter struct {
	r	io.Reader
	err	error
	buf	[]byte
	block	*blockEnc
}

// Convert the Snappy stream supplied in 'in' and write the zStandard stream to 'w'.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:79
// If any error is detected on the Snappy stream it is returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:79
// The number of bytes written is returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:82
func (r *SnappyConverter) Convert(in io.Reader, w io.Writer) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:82
	_go_fuzz_dep_.CoverTab[95062]++
												initPredefined()
												r.err = nil
												r.r = in
												if r.block == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:86
		_go_fuzz_dep_.CoverTab[95066]++
													r.block = &blockEnc{}
													r.block.init()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:88
		// _ = "end of CoverTab[95066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:89
		_go_fuzz_dep_.CoverTab[95067]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:89
		// _ = "end of CoverTab[95067]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:89
	// _ = "end of CoverTab[95062]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:89
	_go_fuzz_dep_.CoverTab[95063]++
												r.block.initNewEncode()
												if len(r.buf) != snappyMaxEncodedLenOfMaxBlockSize+snappyChecksumSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:91
		_go_fuzz_dep_.CoverTab[95068]++
													r.buf = make([]byte, snappyMaxEncodedLenOfMaxBlockSize+snappyChecksumSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:92
		// _ = "end of CoverTab[95068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:93
		_go_fuzz_dep_.CoverTab[95069]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:93
		// _ = "end of CoverTab[95069]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:93
	// _ = "end of CoverTab[95063]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:93
	_go_fuzz_dep_.CoverTab[95064]++
												r.block.litEnc.Reuse = huff0.ReusePolicyNone
												var written int64
												var readHeader bool
												{
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:97
		_go_fuzz_dep_.CoverTab[95070]++
													var header []byte
													var n int
													header, r.err = frameHeader{WindowSize: snappyMaxBlockSize}.appendTo(r.buf[:0])

													n, r.err = w.Write(header)
													if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:103
			_go_fuzz_dep_.CoverTab[95072]++
														return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:104
			// _ = "end of CoverTab[95072]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:105
			_go_fuzz_dep_.CoverTab[95073]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:105
			// _ = "end of CoverTab[95073]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:105
		// _ = "end of CoverTab[95070]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:105
		_go_fuzz_dep_.CoverTab[95071]++
													written += int64(n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:106
		// _ = "end of CoverTab[95071]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:107
	// _ = "end of CoverTab[95064]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:107
	_go_fuzz_dep_.CoverTab[95065]++

												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:109
		_go_fuzz_dep_.CoverTab[95074]++
													if !r.readFull(r.buf[:4], true) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:110
			_go_fuzz_dep_.CoverTab[95080]++

														r.block.reset(nil)
														r.block.last = true
														err := r.block.encodeLits(r.block.literals, false)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:115
				_go_fuzz_dep_.CoverTab[95083]++
															return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:116
				// _ = "end of CoverTab[95083]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:117
				_go_fuzz_dep_.CoverTab[95084]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:117
				// _ = "end of CoverTab[95084]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:117
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:117
			// _ = "end of CoverTab[95080]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:117
			_go_fuzz_dep_.CoverTab[95081]++
														n, err := w.Write(r.block.output)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:119
				_go_fuzz_dep_.CoverTab[95085]++
															return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:120
				// _ = "end of CoverTab[95085]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:121
				_go_fuzz_dep_.CoverTab[95086]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:121
				// _ = "end of CoverTab[95086]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:121
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:121
			// _ = "end of CoverTab[95081]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:121
			_go_fuzz_dep_.CoverTab[95082]++
														written += int64(n)

														return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:124
			// _ = "end of CoverTab[95082]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:125
			_go_fuzz_dep_.CoverTab[95087]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:125
			// _ = "end of CoverTab[95087]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:125
		// _ = "end of CoverTab[95074]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:125
		_go_fuzz_dep_.CoverTab[95075]++
													chunkType := r.buf[0]
													if !readHeader {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:127
			_go_fuzz_dep_.CoverTab[95088]++
														if chunkType != chunkTypeStreamIdentifier {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:128
				_go_fuzz_dep_.CoverTab[95090]++
															println("chunkType != chunkTypeStreamIdentifier", chunkType)
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:131
				// _ = "end of CoverTab[95090]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:132
				_go_fuzz_dep_.CoverTab[95091]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:132
				// _ = "end of CoverTab[95091]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:132
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:132
			// _ = "end of CoverTab[95088]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:132
			_go_fuzz_dep_.CoverTab[95089]++
														readHeader = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:133
			// _ = "end of CoverTab[95089]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:134
			_go_fuzz_dep_.CoverTab[95092]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:134
			// _ = "end of CoverTab[95092]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:134
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:134
		// _ = "end of CoverTab[95075]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:134
		_go_fuzz_dep_.CoverTab[95076]++
													chunkLen := int(r.buf[1]) | int(r.buf[2])<<8 | int(r.buf[3])<<16
													if chunkLen > len(r.buf) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:136
			_go_fuzz_dep_.CoverTab[95093]++
														println("chunkLen > len(r.buf)", chunkType)
														r.err = ErrSnappyUnsupported
														return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:139
			// _ = "end of CoverTab[95093]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:140
			_go_fuzz_dep_.CoverTab[95094]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:140
			// _ = "end of CoverTab[95094]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:140
		// _ = "end of CoverTab[95076]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:140
		_go_fuzz_dep_.CoverTab[95077]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:144
		switch chunkType {
		case chunkTypeCompressedData:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:145
			_go_fuzz_dep_.CoverTab[95095]++

														if chunkLen < snappyChecksumSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:147
				_go_fuzz_dep_.CoverTab[95119]++
															println("chunkLen < snappyChecksumSize", chunkLen, snappyChecksumSize)
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:150
				// _ = "end of CoverTab[95119]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:151
				_go_fuzz_dep_.CoverTab[95120]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:151
				// _ = "end of CoverTab[95120]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:151
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:151
			// _ = "end of CoverTab[95095]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:151
			_go_fuzz_dep_.CoverTab[95096]++
														buf := r.buf[:chunkLen]
														if !r.readFull(buf, false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:153
				_go_fuzz_dep_.CoverTab[95121]++
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:154
				// _ = "end of CoverTab[95121]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:155
				_go_fuzz_dep_.CoverTab[95122]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:155
				// _ = "end of CoverTab[95122]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:155
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:155
			// _ = "end of CoverTab[95096]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:155
			_go_fuzz_dep_.CoverTab[95097]++

														buf = buf[snappyChecksumSize:]

														n, hdr, err := snappyDecodedLen(buf)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:160
				_go_fuzz_dep_.CoverTab[95123]++
															r.err = err
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:162
				// _ = "end of CoverTab[95123]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:163
				_go_fuzz_dep_.CoverTab[95124]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:163
				// _ = "end of CoverTab[95124]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:163
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:163
			// _ = "end of CoverTab[95097]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:163
			_go_fuzz_dep_.CoverTab[95098]++
														buf = buf[hdr:]
														if n > snappyMaxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:165
				_go_fuzz_dep_.CoverTab[95125]++
															println("n > snappyMaxBlockSize", n, snappyMaxBlockSize)
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:168
				// _ = "end of CoverTab[95125]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:169
				_go_fuzz_dep_.CoverTab[95126]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:169
				// _ = "end of CoverTab[95126]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:169
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:169
			// _ = "end of CoverTab[95098]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:169
			_go_fuzz_dep_.CoverTab[95099]++
														r.block.reset(nil)
														r.block.pushOffsets()
														if err := decodeSnappy(r.block, buf); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:172
				_go_fuzz_dep_.CoverTab[95127]++
															r.err = err
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:174
				// _ = "end of CoverTab[95127]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:175
				_go_fuzz_dep_.CoverTab[95128]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:175
				// _ = "end of CoverTab[95128]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:175
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:175
			// _ = "end of CoverTab[95099]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:175
			_go_fuzz_dep_.CoverTab[95100]++
														if r.block.size+r.block.extraLits != n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:176
				_go_fuzz_dep_.CoverTab[95129]++
															printf("invalid size, want %d, got %d\n", n, r.block.size+r.block.extraLits)
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:179
				// _ = "end of CoverTab[95129]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:180
				_go_fuzz_dep_.CoverTab[95130]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:180
				// _ = "end of CoverTab[95130]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:180
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:180
			// _ = "end of CoverTab[95100]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:180
			_go_fuzz_dep_.CoverTab[95101]++
														err = r.block.encode(nil, false, false)
														switch err {
			case errIncompressible:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:183
				_go_fuzz_dep_.CoverTab[95131]++
															r.block.popOffsets()
															r.block.reset(nil)
															r.block.literals, err = snappy.Decode(r.block.literals[:n], r.buf[snappyChecksumSize:chunkLen])
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:187
					_go_fuzz_dep_.CoverTab[95135]++
																return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:188
					// _ = "end of CoverTab[95135]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:189
					_go_fuzz_dep_.CoverTab[95136]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:189
					// _ = "end of CoverTab[95136]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:189
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:189
				// _ = "end of CoverTab[95131]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:189
				_go_fuzz_dep_.CoverTab[95132]++
															err = r.block.encodeLits(r.block.literals, false)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:191
					_go_fuzz_dep_.CoverTab[95137]++
																return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:192
					// _ = "end of CoverTab[95137]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:193
					_go_fuzz_dep_.CoverTab[95138]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:193
					// _ = "end of CoverTab[95138]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:193
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:193
				// _ = "end of CoverTab[95132]"
			case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:194
				_go_fuzz_dep_.CoverTab[95133]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:194
				// _ = "end of CoverTab[95133]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:195
				_go_fuzz_dep_.CoverTab[95134]++
															return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:196
				// _ = "end of CoverTab[95134]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:197
			// _ = "end of CoverTab[95101]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:197
			_go_fuzz_dep_.CoverTab[95102]++

														n, r.err = w.Write(r.block.output)
														if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:200
				_go_fuzz_dep_.CoverTab[95139]++
															return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:201
				// _ = "end of CoverTab[95139]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:202
				_go_fuzz_dep_.CoverTab[95140]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:202
				// _ = "end of CoverTab[95140]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:202
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:202
			// _ = "end of CoverTab[95102]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:202
			_go_fuzz_dep_.CoverTab[95103]++
														written += int64(n)
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:204
			// _ = "end of CoverTab[95103]"
		case chunkTypeUncompressedData:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:205
			_go_fuzz_dep_.CoverTab[95104]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:206
				_go_fuzz_dep_.CoverTab[95141]++
															println("Uncompressed, chunklen", chunkLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:207
				// _ = "end of CoverTab[95141]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:208
				_go_fuzz_dep_.CoverTab[95142]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:208
				// _ = "end of CoverTab[95142]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:208
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:208
			// _ = "end of CoverTab[95104]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:208
			_go_fuzz_dep_.CoverTab[95105]++

														if chunkLen < snappyChecksumSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:210
				_go_fuzz_dep_.CoverTab[95143]++
															println("chunkLen < snappyChecksumSize", chunkLen, snappyChecksumSize)
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:213
				// _ = "end of CoverTab[95143]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:214
				_go_fuzz_dep_.CoverTab[95144]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:214
				// _ = "end of CoverTab[95144]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:214
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:214
			// _ = "end of CoverTab[95105]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:214
			_go_fuzz_dep_.CoverTab[95106]++
														r.block.reset(nil)
														buf := r.buf[:snappyChecksumSize]
														if !r.readFull(buf, false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:217
				_go_fuzz_dep_.CoverTab[95145]++
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:218
				// _ = "end of CoverTab[95145]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:219
				_go_fuzz_dep_.CoverTab[95146]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:219
				// _ = "end of CoverTab[95146]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:219
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:219
			// _ = "end of CoverTab[95106]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:219
			_go_fuzz_dep_.CoverTab[95107]++
														checksum := uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24

														n := chunkLen - snappyChecksumSize
														if n > snappyMaxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:223
				_go_fuzz_dep_.CoverTab[95147]++
															println("n > snappyMaxBlockSize", n, snappyMaxBlockSize)
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:226
				// _ = "end of CoverTab[95147]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:227
				_go_fuzz_dep_.CoverTab[95148]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:227
				// _ = "end of CoverTab[95148]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:227
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:227
			// _ = "end of CoverTab[95107]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:227
			_go_fuzz_dep_.CoverTab[95108]++
														r.block.literals = r.block.literals[:n]
														if !r.readFull(r.block.literals, false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:229
				_go_fuzz_dep_.CoverTab[95149]++
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:230
				// _ = "end of CoverTab[95149]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:231
				_go_fuzz_dep_.CoverTab[95150]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:231
				// _ = "end of CoverTab[95150]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:231
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:231
			// _ = "end of CoverTab[95108]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:231
			_go_fuzz_dep_.CoverTab[95109]++
														if snappyCRC(r.block.literals) != checksum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:232
				_go_fuzz_dep_.CoverTab[95151]++
															println("literals crc mismatch")
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:235
				// _ = "end of CoverTab[95151]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:236
				_go_fuzz_dep_.CoverTab[95152]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:236
				// _ = "end of CoverTab[95152]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:236
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:236
			// _ = "end of CoverTab[95109]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:236
			_go_fuzz_dep_.CoverTab[95110]++
														err := r.block.encodeLits(r.block.literals, false)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:238
				_go_fuzz_dep_.CoverTab[95153]++
															return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:239
				// _ = "end of CoverTab[95153]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:240
				_go_fuzz_dep_.CoverTab[95154]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:240
				// _ = "end of CoverTab[95154]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:240
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:240
			// _ = "end of CoverTab[95110]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:240
			_go_fuzz_dep_.CoverTab[95111]++
														n, r.err = w.Write(r.block.output)
														if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:242
				_go_fuzz_dep_.CoverTab[95155]++
															return written, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:243
				// _ = "end of CoverTab[95155]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:244
				_go_fuzz_dep_.CoverTab[95156]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:244
				// _ = "end of CoverTab[95156]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:244
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:244
			// _ = "end of CoverTab[95111]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:244
			_go_fuzz_dep_.CoverTab[95112]++
														written += int64(n)
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:246
			// _ = "end of CoverTab[95112]"

		case chunkTypeStreamIdentifier:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:248
			_go_fuzz_dep_.CoverTab[95113]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:249
				_go_fuzz_dep_.CoverTab[95157]++
															println("stream id", chunkLen, len(snappyMagicBody))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:250
				// _ = "end of CoverTab[95157]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:251
				_go_fuzz_dep_.CoverTab[95158]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:251
				// _ = "end of CoverTab[95158]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:251
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:251
			// _ = "end of CoverTab[95113]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:251
			_go_fuzz_dep_.CoverTab[95114]++

														if chunkLen != len(snappyMagicBody) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:253
				_go_fuzz_dep_.CoverTab[95159]++
															println("chunkLen != len(snappyMagicBody)", chunkLen, len(snappyMagicBody))
															r.err = ErrSnappyCorrupt
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:256
				// _ = "end of CoverTab[95159]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:257
				_go_fuzz_dep_.CoverTab[95160]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:257
				// _ = "end of CoverTab[95160]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:257
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:257
			// _ = "end of CoverTab[95114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:257
			_go_fuzz_dep_.CoverTab[95115]++
														if !r.readFull(r.buf[:len(snappyMagicBody)], false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:258
				_go_fuzz_dep_.CoverTab[95161]++
															return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:259
				// _ = "end of CoverTab[95161]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:260
				_go_fuzz_dep_.CoverTab[95162]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:260
				// _ = "end of CoverTab[95162]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:260
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:260
			// _ = "end of CoverTab[95115]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:260
			_go_fuzz_dep_.CoverTab[95116]++
														for i := 0; i < len(snappyMagicBody); i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:261
				_go_fuzz_dep_.CoverTab[95163]++
															if r.buf[i] != snappyMagicBody[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:262
					_go_fuzz_dep_.CoverTab[95164]++
																println("r.buf[i] != snappyMagicBody[i]", r.buf[i], snappyMagicBody[i], i)
																r.err = ErrSnappyCorrupt
																return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:265
					// _ = "end of CoverTab[95164]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:266
					_go_fuzz_dep_.CoverTab[95165]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:266
					// _ = "end of CoverTab[95165]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:266
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:266
				// _ = "end of CoverTab[95163]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:267
			// _ = "end of CoverTab[95116]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:267
			_go_fuzz_dep_.CoverTab[95117]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:268
			// _ = "end of CoverTab[95117]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:268
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:268
			_go_fuzz_dep_.CoverTab[95118]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:268
			// _ = "end of CoverTab[95118]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:269
		// _ = "end of CoverTab[95077]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:269
		_go_fuzz_dep_.CoverTab[95078]++

													if chunkType <= 0x7f {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:271
			_go_fuzz_dep_.CoverTab[95166]++

														println("chunkType <= 0x7f")
														r.err = ErrSnappyUnsupported
														return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:275
			// _ = "end of CoverTab[95166]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:276
			_go_fuzz_dep_.CoverTab[95167]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:276
			// _ = "end of CoverTab[95167]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:276
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:276
		// _ = "end of CoverTab[95078]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:276
		_go_fuzz_dep_.CoverTab[95079]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:279
		if !r.readFull(r.buf[:chunkLen], false) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:279
			_go_fuzz_dep_.CoverTab[95168]++
														return written, r.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:280
			// _ = "end of CoverTab[95168]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:281
			_go_fuzz_dep_.CoverTab[95169]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:281
			// _ = "end of CoverTab[95169]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:281
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:281
		// _ = "end of CoverTab[95079]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:282
	// _ = "end of CoverTab[95065]"
}

// decodeSnappy writes the decoding of src to dst. It assumes that the varint-encoded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:285
// length of the decompressed bytes has already been read.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:287
func decodeSnappy(blk *blockEnc, src []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:287
	_go_fuzz_dep_.CoverTab[95170]++
	//decodeRef(make([]byte, snappyMaxBlockSize), src)
	var s, length int
	lits := blk.extraLits
	var offset uint32
	for s < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:292
		_go_fuzz_dep_.CoverTab[95172]++
													switch src[s] & 0x03 {
		case snappyTagLiteral:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:294
			_go_fuzz_dep_.CoverTab[95176]++
														x := uint32(src[s] >> 2)
														switch {
			case x < 60:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:297
				_go_fuzz_dep_.CoverTab[95187]++
															s++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:298
				// _ = "end of CoverTab[95187]"
			case x == 60:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:299
				_go_fuzz_dep_.CoverTab[95188]++
															s += 2
															if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:301
					_go_fuzz_dep_.CoverTab[95197]++
																println("uint(s) > uint(len(src)", s, src)
																return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:303
					// _ = "end of CoverTab[95197]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:304
					_go_fuzz_dep_.CoverTab[95198]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:304
					// _ = "end of CoverTab[95198]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:304
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:304
				// _ = "end of CoverTab[95188]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:304
				_go_fuzz_dep_.CoverTab[95189]++
															x = uint32(src[s-1])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:305
				// _ = "end of CoverTab[95189]"
			case x == 61:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:306
				_go_fuzz_dep_.CoverTab[95190]++
															s += 3
															if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:308
					_go_fuzz_dep_.CoverTab[95199]++
																println("uint(s) > uint(len(src)", s, src)
																return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:310
					// _ = "end of CoverTab[95199]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:311
					_go_fuzz_dep_.CoverTab[95200]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:311
					// _ = "end of CoverTab[95200]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:311
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:311
				// _ = "end of CoverTab[95190]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:311
				_go_fuzz_dep_.CoverTab[95191]++
															x = uint32(src[s-2]) | uint32(src[s-1])<<8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:312
				// _ = "end of CoverTab[95191]"
			case x == 62:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:313
				_go_fuzz_dep_.CoverTab[95192]++
															s += 4
															if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:315
					_go_fuzz_dep_.CoverTab[95201]++
																println("uint(s) > uint(len(src)", s, src)
																return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:317
					// _ = "end of CoverTab[95201]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:318
					_go_fuzz_dep_.CoverTab[95202]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:318
					// _ = "end of CoverTab[95202]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:318
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:318
				// _ = "end of CoverTab[95192]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:318
				_go_fuzz_dep_.CoverTab[95193]++
															x = uint32(src[s-3]) | uint32(src[s-2])<<8 | uint32(src[s-1])<<16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:319
				// _ = "end of CoverTab[95193]"
			case x == 63:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:320
				_go_fuzz_dep_.CoverTab[95194]++
															s += 5
															if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:322
					_go_fuzz_dep_.CoverTab[95203]++
																println("uint(s) > uint(len(src)", s, src)
																return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:324
					// _ = "end of CoverTab[95203]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:325
					_go_fuzz_dep_.CoverTab[95204]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:325
					// _ = "end of CoverTab[95204]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:325
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:325
				// _ = "end of CoverTab[95194]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:325
				_go_fuzz_dep_.CoverTab[95195]++
															x = uint32(src[s-4]) | uint32(src[s-3])<<8 | uint32(src[s-2])<<16 | uint32(src[s-1])<<24
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:326
				// _ = "end of CoverTab[95195]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:326
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:326
				_go_fuzz_dep_.CoverTab[95196]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:326
				// _ = "end of CoverTab[95196]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:327
			// _ = "end of CoverTab[95176]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:327
			_go_fuzz_dep_.CoverTab[95177]++
														if x > snappyMaxBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:328
				_go_fuzz_dep_.CoverTab[95205]++
															println("x > snappyMaxBlockSize", x, snappyMaxBlockSize)
															return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:330
				// _ = "end of CoverTab[95205]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:331
				_go_fuzz_dep_.CoverTab[95206]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:331
				// _ = "end of CoverTab[95206]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:331
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:331
			// _ = "end of CoverTab[95177]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:331
			_go_fuzz_dep_.CoverTab[95178]++
														length = int(x) + 1
														if length <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:333
				_go_fuzz_dep_.CoverTab[95207]++
															println("length <= 0 ", length)

															return errUnsupportedLiteralLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:336
				// _ = "end of CoverTab[95207]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:337
				_go_fuzz_dep_.CoverTab[95208]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:337
				// _ = "end of CoverTab[95208]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:337
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:337
			// _ = "end of CoverTab[95178]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:337
			_go_fuzz_dep_.CoverTab[95179]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:342
			blk.literals = append(blk.literals, src[s:s+length]...)

														lits += length
														s += length
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:346
			// _ = "end of CoverTab[95179]"

		case snappyTagCopy1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:348
			_go_fuzz_dep_.CoverTab[95180]++
														s += 2
														if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:350
				_go_fuzz_dep_.CoverTab[95209]++
															println("uint(s) > uint(len(src)", s, len(src))
															return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:352
				// _ = "end of CoverTab[95209]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:353
				_go_fuzz_dep_.CoverTab[95210]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:353
				// _ = "end of CoverTab[95210]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:353
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:353
			// _ = "end of CoverTab[95180]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:353
			_go_fuzz_dep_.CoverTab[95181]++
														length = 4 + int(src[s-2])>>2&0x7
														offset = uint32(src[s-2])&0xe0<<3 | uint32(src[s-1])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:355
			// _ = "end of CoverTab[95181]"

		case snappyTagCopy2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:357
			_go_fuzz_dep_.CoverTab[95182]++
														s += 3
														if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:359
				_go_fuzz_dep_.CoverTab[95211]++
															println("uint(s) > uint(len(src)", s, len(src))
															return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:361
				// _ = "end of CoverTab[95211]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:362
				_go_fuzz_dep_.CoverTab[95212]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:362
				// _ = "end of CoverTab[95212]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:362
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:362
			// _ = "end of CoverTab[95182]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:362
			_go_fuzz_dep_.CoverTab[95183]++
														length = 1 + int(src[s-3])>>2
														offset = uint32(src[s-2]) | uint32(src[s-1])<<8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:364
			// _ = "end of CoverTab[95183]"

		case snappyTagCopy4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:366
			_go_fuzz_dep_.CoverTab[95184]++
														s += 5
														if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:368
				_go_fuzz_dep_.CoverTab[95213]++
															println("uint(s) > uint(len(src)", s, len(src))
															return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:370
				// _ = "end of CoverTab[95213]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:371
				_go_fuzz_dep_.CoverTab[95214]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:371
				// _ = "end of CoverTab[95214]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:371
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:371
			// _ = "end of CoverTab[95184]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:371
			_go_fuzz_dep_.CoverTab[95185]++
														length = 1 + int(src[s-5])>>2
														offset = uint32(src[s-4]) | uint32(src[s-3])<<8 | uint32(src[s-2])<<16 | uint32(src[s-1])<<24
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:373
			// _ = "end of CoverTab[95185]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:373
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:373
			_go_fuzz_dep_.CoverTab[95186]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:373
			// _ = "end of CoverTab[95186]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:374
		// _ = "end of CoverTab[95172]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:374
		_go_fuzz_dep_.CoverTab[95173]++

													if offset <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:376
			_go_fuzz_dep_.CoverTab[95215]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:376
			return blk.size+lits < int(offset)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:376
			// _ = "end of CoverTab[95215]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:376
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:376
			_go_fuzz_dep_.CoverTab[95216]++
														println("offset <= 0 || blk.size+lits < int(offset)", offset, blk.size+lits, int(offset), blk.size, lits)

														return ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:379
			// _ = "end of CoverTab[95216]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:380
			_go_fuzz_dep_.CoverTab[95217]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:380
			// _ = "end of CoverTab[95217]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:380
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:380
		// _ = "end of CoverTab[95173]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:380
		_go_fuzz_dep_.CoverTab[95174]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:385
		if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:385
			_go_fuzz_dep_.CoverTab[95218]++
														offset = blk.matchOffset(offset, uint32(lits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:386
			// _ = "end of CoverTab[95218]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:387
			_go_fuzz_dep_.CoverTab[95219]++
														offset += 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:388
			// _ = "end of CoverTab[95219]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:389
		// _ = "end of CoverTab[95174]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:389
		_go_fuzz_dep_.CoverTab[95175]++

													blk.sequences = append(blk.sequences, seq{
			litLen:		uint32(lits),
			offset:		offset,
			matchLen:	uint32(length) - zstdMinMatch,
		})
													blk.size += length + lits
													lits = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:397
		// _ = "end of CoverTab[95175]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:398
	// _ = "end of CoverTab[95170]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:398
	_go_fuzz_dep_.CoverTab[95171]++
												blk.extraLits = lits
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:400
	// _ = "end of CoverTab[95171]"
}

func (r *SnappyConverter) readFull(p []byte, allowEOF bool) (ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:403
	_go_fuzz_dep_.CoverTab[95220]++
												if _, r.err = io.ReadFull(r.r, p); r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:404
		_go_fuzz_dep_.CoverTab[95222]++
													if r.err == io.ErrUnexpectedEOF || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
			_go_fuzz_dep_.CoverTab[95224]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
			return (r.err == io.EOF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
				_go_fuzz_dep_.CoverTab[95225]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
				return !allowEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
				// _ = "end of CoverTab[95225]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
			}())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
			// _ = "end of CoverTab[95224]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:405
			_go_fuzz_dep_.CoverTab[95226]++
														r.err = ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:406
			// _ = "end of CoverTab[95226]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:407
			_go_fuzz_dep_.CoverTab[95227]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:407
			// _ = "end of CoverTab[95227]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:407
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:407
		// _ = "end of CoverTab[95222]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:407
		_go_fuzz_dep_.CoverTab[95223]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:408
		// _ = "end of CoverTab[95223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:409
		_go_fuzz_dep_.CoverTab[95228]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:409
		// _ = "end of CoverTab[95228]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:409
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:409
	// _ = "end of CoverTab[95220]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:409
	_go_fuzz_dep_.CoverTab[95221]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:410
	// _ = "end of CoverTab[95221]"
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

// crc implements the checksum specified in section 3 of
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:415
// https://github.com/google/snappy/blob/master/framing_format.txt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:417
func snappyCRC(b []byte) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:417
	_go_fuzz_dep_.CoverTab[95229]++
												c := crc32.Update(0, crcTable, b)
												return c>>15 | c<<17 + 0xa282ead8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:419
	// _ = "end of CoverTab[95229]"
}

// snappyDecodedLen returns the length of the decoded block and the number of bytes
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:422
// that the length header occupied.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:424
func snappyDecodedLen(src []byte) (blockLen, headerLen int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:424
	_go_fuzz_dep_.CoverTab[95230]++
												v, n := binary.Uvarint(src)
												if n <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:426
		_go_fuzz_dep_.CoverTab[95233]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:426
		return v > 0xffffffff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:426
		// _ = "end of CoverTab[95233]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:426
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:426
		_go_fuzz_dep_.CoverTab[95234]++
													return 0, 0, ErrSnappyCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:427
		// _ = "end of CoverTab[95234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:428
		_go_fuzz_dep_.CoverTab[95235]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:428
		// _ = "end of CoverTab[95235]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:428
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:428
	// _ = "end of CoverTab[95230]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:428
	_go_fuzz_dep_.CoverTab[95231]++

												const wordSize = 32 << (^uint(0) >> 32 & 1)
												if wordSize == 32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:431
		_go_fuzz_dep_.CoverTab[95236]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:431
		return v > 0x7fffffff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:431
		// _ = "end of CoverTab[95236]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:431
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:431
		_go_fuzz_dep_.CoverTab[95237]++
													return 0, 0, ErrSnappyTooLarge
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:432
		// _ = "end of CoverTab[95237]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:433
		_go_fuzz_dep_.CoverTab[95238]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:433
		// _ = "end of CoverTab[95238]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:433
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:433
	// _ = "end of CoverTab[95231]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:433
	_go_fuzz_dep_.CoverTab[95232]++
												return int(v), n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:434
	// _ = "end of CoverTab[95232]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:435
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/snappy.go:435
var _ = _go_fuzz_dep_.CoverTab
