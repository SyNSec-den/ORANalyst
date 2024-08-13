// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/inflate.go:5
// Package flate implements the DEFLATE compressed data format, described in
//line /usr/local/go/src/compress/flate/inflate.go:5
// RFC 1951.  The gzip and zlib packages implement access to DEFLATE-based file
//line /usr/local/go/src/compress/flate/inflate.go:5
// formats.
//line /usr/local/go/src/compress/flate/inflate.go:8
package flate

//line /usr/local/go/src/compress/flate/inflate.go:8
import (
//line /usr/local/go/src/compress/flate/inflate.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/inflate.go:8
)
//line /usr/local/go/src/compress/flate/inflate.go:8
import (
//line /usr/local/go/src/compress/flate/inflate.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/inflate.go:8
)

import (
	"bufio"
	"io"
	"math/bits"
	"strconv"
	"sync"
)

const (
	maxCodeLen	= 16	// max length of Huffman code
	// The next three numbers come from the RFC section 3.2.7, with the
	// additional proviso in section 3.2.5 which implies that distance codes
	// 30 and 31 should never occur in compressed data.
	maxNumLit	= 286
	maxNumDist	= 30
	numCodes	= 19	// number of codes in Huffman meta-code
)

// Initialize the fixedHuffmanDecoder only once upon first use.
var fixedOnce sync.Once
var fixedHuffmanDecoder huffmanDecoder

// A CorruptInputError reports the presence of corrupt input at a given offset.
type CorruptInputError int64

func (e CorruptInputError) Error() string {
//line /usr/local/go/src/compress/flate/inflate.go:35
	_go_fuzz_dep_.CoverTab[26312]++
							return "flate: corrupt input before offset " + strconv.FormatInt(int64(e), 10)
//line /usr/local/go/src/compress/flate/inflate.go:36
	// _ = "end of CoverTab[26312]"
}

// An InternalError reports an error in the flate code itself.
type InternalError string

func (e InternalError) Error() string {
//line /usr/local/go/src/compress/flate/inflate.go:42
	_go_fuzz_dep_.CoverTab[26313]++
//line /usr/local/go/src/compress/flate/inflate.go:42
	return "flate: internal error: " + string(e)
//line /usr/local/go/src/compress/flate/inflate.go:42
	// _ = "end of CoverTab[26313]"
//line /usr/local/go/src/compress/flate/inflate.go:42
}

// A ReadError reports an error encountered while reading input.
//line /usr/local/go/src/compress/flate/inflate.go:44
//
//line /usr/local/go/src/compress/flate/inflate.go:44
// Deprecated: No longer returned.
//line /usr/local/go/src/compress/flate/inflate.go:47
type ReadError struct {
	Offset	int64	// byte offset where error occurred
	Err	error	// error returned by underlying Read
}

func (e *ReadError) Error() string {
//line /usr/local/go/src/compress/flate/inflate.go:52
	_go_fuzz_dep_.CoverTab[26314]++
							return "flate: read error at offset " + strconv.FormatInt(e.Offset, 10) + ": " + e.Err.Error()
//line /usr/local/go/src/compress/flate/inflate.go:53
	// _ = "end of CoverTab[26314]"
}

// A WriteError reports an error encountered while writing output.
//line /usr/local/go/src/compress/flate/inflate.go:56
//
//line /usr/local/go/src/compress/flate/inflate.go:56
// Deprecated: No longer returned.
//line /usr/local/go/src/compress/flate/inflate.go:59
type WriteError struct {
	Offset	int64	// byte offset where error occurred
	Err	error	// error returned by underlying Write
}

func (e *WriteError) Error() string {
//line /usr/local/go/src/compress/flate/inflate.go:64
	_go_fuzz_dep_.CoverTab[26315]++
							return "flate: write error at offset " + strconv.FormatInt(e.Offset, 10) + ": " + e.Err.Error()
//line /usr/local/go/src/compress/flate/inflate.go:65
	// _ = "end of CoverTab[26315]"
}

// Resetter resets a ReadCloser returned by NewReader or NewReaderDict
//line /usr/local/go/src/compress/flate/inflate.go:68
// to switch to a new underlying Reader. This permits reusing a ReadCloser
//line /usr/local/go/src/compress/flate/inflate.go:68
// instead of allocating a new one.
//line /usr/local/go/src/compress/flate/inflate.go:71
type Resetter interface {
	// Reset discards any buffered data and resets the Resetter as if it was
	// newly initialized with the given reader.
	Reset(r io.Reader, dict []byte) error
}

//line /usr/local/go/src/compress/flate/inflate.go:97
const (
	huffmanChunkBits	= 9
	huffmanNumChunks	= 1 << huffmanChunkBits
	huffmanCountMask	= 15
	huffmanValueShift	= 4
)

type huffmanDecoder struct {
	min		int				// the minimum code length
	chunks		[huffmanNumChunks]uint32	// chunks as described above
	links		[][]uint32			// overflow links
	linkMask	uint32				// mask the width of the link table
}

// Initialize Huffman decoding tables from array of code lengths.
//line /usr/local/go/src/compress/flate/inflate.go:111
// Following this function, h is guaranteed to be initialized into a complete
//line /usr/local/go/src/compress/flate/inflate.go:111
// tree (i.e., neither over-subscribed nor under-subscribed). The exception is a
//line /usr/local/go/src/compress/flate/inflate.go:111
// degenerate case where the tree has only a single symbol with length 1. Empty
//line /usr/local/go/src/compress/flate/inflate.go:111
// trees are permitted.
//line /usr/local/go/src/compress/flate/inflate.go:116
func (h *huffmanDecoder) init(lengths []int) bool {
	// Sanity enables additional runtime tests during Huffman
	// table construction. It's intended to be used during
	// development to supplement the currently ad-hoc unit tests.
	const sanity = false

	if h.min != 0 {
		*h = huffmanDecoder{}
	}

	// Count number of codes of each length,
	// compute min and max length.
	var count [maxCodeLen]int
	var min, max int
	for _, n := range lengths {
		if n == 0 {
			continue
		}
		if min == 0 || n < min {
			min = n
		}
		if n > max {
			max = n
		}
		count[n]++
	}

//line /usr/local/go/src/compress/flate/inflate.go:150
	if max == 0 {
		return true
	}

	code := 0
	var nextcode [maxCodeLen]int
	for i := min; i <= max; i++ {
		code <<= 1
		nextcode[i] = code
		code += count[i]
	}

//line /usr/local/go/src/compress/flate/inflate.go:167
	if code != 1<<uint(max) && !(code == 1 && max == 1) {
		return false
	}

	h.min = min
	if max > huffmanChunkBits {
								numLinks := 1 << (uint(max) - huffmanChunkBits)
								h.linkMask = uint32(numLinks - 1)

//line /usr/local/go/src/compress/flate/inflate.go:177
		link := nextcode[huffmanChunkBits+1] >> 1
		h.links = make([][]uint32, huffmanNumChunks-link)
		for j := uint(link); j < huffmanNumChunks; j++ {
			reverse := int(bits.Reverse16(uint16(j)))
			reverse >>= uint(16 - huffmanChunkBits)
			off := j - uint(link)
			if sanity && h.chunks[reverse] != 0 {
				panic("impossible: overwriting existing chunk")
			}
			h.chunks[reverse] = uint32(off<<huffmanValueShift | (huffmanChunkBits + 1))
			h.links[off] = make([]uint32, numLinks)
		}
	}

	for i, n := range lengths {
		if n == 0 {
			continue
		}
		code := nextcode[n]
		nextcode[n]++
		chunk := uint32(i<<huffmanValueShift | n)
		reverse := int(bits.Reverse16(uint16(code)))
		reverse >>= uint(16 - n)
		if n <= huffmanChunkBits {
			for off := reverse; off < len(h.chunks); off += 1 << uint(n) {

//line /usr/local/go/src/compress/flate/inflate.go:207
				if sanity && h.chunks[off] != 0 {
					panic("impossible: overwriting existing chunk")
				}
				h.chunks[off] = chunk
			}
		} else {
			j := reverse & (huffmanNumChunks - 1)
			if sanity && h.chunks[j]&huffmanCountMask != huffmanChunkBits+1 {

//line /usr/local/go/src/compress/flate/inflate.go:217
				panic("impossible: not an indirect chunk")
			}
			value := h.chunks[j] >> huffmanValueShift
			linktab := h.links[value]
			reverse >>= huffmanChunkBits
			for off := reverse; off < len(linktab); off += 1 << uint(n-huffmanChunkBits) {
				if sanity && linktab[off] != 0 {
					panic("impossible: overwriting existing chunk")
				}
				linktab[off] = chunk
			}
		}
	}

	if sanity {

//line /usr/local/go/src/compress/flate/inflate.go:235
		for i, chunk := range h.chunks {
			if chunk == 0 {

//line /usr/local/go/src/compress/flate/inflate.go:240
				if code == 1 && i%2 == 1 {
					continue
				}
				panic("impossible: missing chunk")
			}
		}
		for _, linktab := range h.links {
			for _, chunk := range linktab {
				if chunk == 0 {
					panic("impossible: missing chunk")
				}
			}
		}
	}

	return true
}

// The actual read interface needed by NewReader.
//line /usr/local/go/src/compress/flate/inflate.go:258
// If the passed in io.Reader does not also have ReadByte,
//line /usr/local/go/src/compress/flate/inflate.go:258
// the NewReader will introduce its own buffering.
//line /usr/local/go/src/compress/flate/inflate.go:261
type Reader interface {
	io.Reader
	io.ByteReader
}

// Decompress state.
type decompressor struct {
	// Input source.
	r	Reader
	roffset	int64

	// Input bits, in top of b.
	b	uint32
	nb	uint

	// Huffman decoders for literal/length, distance.
	h1, h2	huffmanDecoder

	// Length arrays used to define Huffman codes.
	bits		*[maxNumLit + maxNumDist]int
	codebits	*[numCodes]int

	// Output history, buffer.
	dict	dictDecoder

	// Temporary buffer (avoids repeated allocation).
	buf	[4]byte

	// Next step in the decompression,
	// and decompression state.
	step		func(*decompressor)
	stepState	int
	final		bool
	err		error
	toRead		[]byte
	hl, hd		*huffmanDecoder
	copyLen		int
	copyDist	int
}

func (f *decompressor) nextBlock() {
//line /usr/local/go/src/compress/flate/inflate.go:301
	_go_fuzz_dep_.CoverTab[26316]++
							for f.nb < 1+2 {
//line /usr/local/go/src/compress/flate/inflate.go:302
		_go_fuzz_dep_.CoverTab[26318]++
								if f.err = f.moreBits(); f.err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:303
			_go_fuzz_dep_.CoverTab[26319]++
									return
//line /usr/local/go/src/compress/flate/inflate.go:304
			// _ = "end of CoverTab[26319]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:305
			_go_fuzz_dep_.CoverTab[26320]++
//line /usr/local/go/src/compress/flate/inflate.go:305
			// _ = "end of CoverTab[26320]"
//line /usr/local/go/src/compress/flate/inflate.go:305
		}
//line /usr/local/go/src/compress/flate/inflate.go:305
		// _ = "end of CoverTab[26318]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:306
	// _ = "end of CoverTab[26316]"
//line /usr/local/go/src/compress/flate/inflate.go:306
	_go_fuzz_dep_.CoverTab[26317]++
							f.final = f.b&1 == 1
							f.b >>= 1
							typ := f.b & 3
							f.b >>= 2
							f.nb -= 1 + 2
							switch typ {
	case 0:
//line /usr/local/go/src/compress/flate/inflate.go:313
		_go_fuzz_dep_.CoverTab[26321]++
								f.dataBlock()
//line /usr/local/go/src/compress/flate/inflate.go:314
		// _ = "end of CoverTab[26321]"
	case 1:
//line /usr/local/go/src/compress/flate/inflate.go:315
		_go_fuzz_dep_.CoverTab[26322]++

								f.hl = &fixedHuffmanDecoder
								f.hd = nil
								f.huffmanBlock()
//line /usr/local/go/src/compress/flate/inflate.go:319
		// _ = "end of CoverTab[26322]"
	case 2:
//line /usr/local/go/src/compress/flate/inflate.go:320
		_go_fuzz_dep_.CoverTab[26323]++

								if f.err = f.readHuffman(); f.err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:322
			_go_fuzz_dep_.CoverTab[26326]++
									break
//line /usr/local/go/src/compress/flate/inflate.go:323
			// _ = "end of CoverTab[26326]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:324
			_go_fuzz_dep_.CoverTab[26327]++
//line /usr/local/go/src/compress/flate/inflate.go:324
			// _ = "end of CoverTab[26327]"
//line /usr/local/go/src/compress/flate/inflate.go:324
		}
//line /usr/local/go/src/compress/flate/inflate.go:324
		// _ = "end of CoverTab[26323]"
//line /usr/local/go/src/compress/flate/inflate.go:324
		_go_fuzz_dep_.CoverTab[26324]++
								f.hl = &f.h1
								f.hd = &f.h2
								f.huffmanBlock()
//line /usr/local/go/src/compress/flate/inflate.go:327
		// _ = "end of CoverTab[26324]"
	default:
//line /usr/local/go/src/compress/flate/inflate.go:328
		_go_fuzz_dep_.CoverTab[26325]++

								f.err = CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:330
		// _ = "end of CoverTab[26325]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:331
	// _ = "end of CoverTab[26317]"
}

func (f *decompressor) Read(b []byte) (int, error) {
//line /usr/local/go/src/compress/flate/inflate.go:334
	_go_fuzz_dep_.CoverTab[26328]++
							for {
//line /usr/local/go/src/compress/flate/inflate.go:335
		_go_fuzz_dep_.CoverTab[26329]++
								if len(f.toRead) > 0 {
//line /usr/local/go/src/compress/flate/inflate.go:336
			_go_fuzz_dep_.CoverTab[26332]++
									n := copy(b, f.toRead)
									f.toRead = f.toRead[n:]
									if len(f.toRead) == 0 {
//line /usr/local/go/src/compress/flate/inflate.go:339
				_go_fuzz_dep_.CoverTab[26334]++
										return n, f.err
//line /usr/local/go/src/compress/flate/inflate.go:340
				// _ = "end of CoverTab[26334]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:341
				_go_fuzz_dep_.CoverTab[26335]++
//line /usr/local/go/src/compress/flate/inflate.go:341
				// _ = "end of CoverTab[26335]"
//line /usr/local/go/src/compress/flate/inflate.go:341
			}
//line /usr/local/go/src/compress/flate/inflate.go:341
			// _ = "end of CoverTab[26332]"
//line /usr/local/go/src/compress/flate/inflate.go:341
			_go_fuzz_dep_.CoverTab[26333]++
									return n, nil
//line /usr/local/go/src/compress/flate/inflate.go:342
			// _ = "end of CoverTab[26333]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:343
			_go_fuzz_dep_.CoverTab[26336]++
//line /usr/local/go/src/compress/flate/inflate.go:343
			// _ = "end of CoverTab[26336]"
//line /usr/local/go/src/compress/flate/inflate.go:343
		}
//line /usr/local/go/src/compress/flate/inflate.go:343
		// _ = "end of CoverTab[26329]"
//line /usr/local/go/src/compress/flate/inflate.go:343
		_go_fuzz_dep_.CoverTab[26330]++
								if f.err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:344
			_go_fuzz_dep_.CoverTab[26337]++
									return 0, f.err
//line /usr/local/go/src/compress/flate/inflate.go:345
			// _ = "end of CoverTab[26337]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:346
			_go_fuzz_dep_.CoverTab[26338]++
//line /usr/local/go/src/compress/flate/inflate.go:346
			// _ = "end of CoverTab[26338]"
//line /usr/local/go/src/compress/flate/inflate.go:346
		}
//line /usr/local/go/src/compress/flate/inflate.go:346
		// _ = "end of CoverTab[26330]"
//line /usr/local/go/src/compress/flate/inflate.go:346
		_go_fuzz_dep_.CoverTab[26331]++
								f.step(f)
								if f.err != nil && func() bool {
//line /usr/local/go/src/compress/flate/inflate.go:348
			_go_fuzz_dep_.CoverTab[26339]++
//line /usr/local/go/src/compress/flate/inflate.go:348
			return len(f.toRead) == 0
//line /usr/local/go/src/compress/flate/inflate.go:348
			// _ = "end of CoverTab[26339]"
//line /usr/local/go/src/compress/flate/inflate.go:348
		}() {
//line /usr/local/go/src/compress/flate/inflate.go:348
			_go_fuzz_dep_.CoverTab[26340]++
									f.toRead = f.dict.readFlush()
//line /usr/local/go/src/compress/flate/inflate.go:349
			// _ = "end of CoverTab[26340]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:350
			_go_fuzz_dep_.CoverTab[26341]++
//line /usr/local/go/src/compress/flate/inflate.go:350
			// _ = "end of CoverTab[26341]"
//line /usr/local/go/src/compress/flate/inflate.go:350
		}
//line /usr/local/go/src/compress/flate/inflate.go:350
		// _ = "end of CoverTab[26331]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:351
	// _ = "end of CoverTab[26328]"
}

func (f *decompressor) Close() error {
//line /usr/local/go/src/compress/flate/inflate.go:354
	_go_fuzz_dep_.CoverTab[26342]++
							if f.err == io.EOF {
//line /usr/local/go/src/compress/flate/inflate.go:355
		_go_fuzz_dep_.CoverTab[26344]++
								return nil
//line /usr/local/go/src/compress/flate/inflate.go:356
		// _ = "end of CoverTab[26344]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:357
		_go_fuzz_dep_.CoverTab[26345]++
//line /usr/local/go/src/compress/flate/inflate.go:357
		// _ = "end of CoverTab[26345]"
//line /usr/local/go/src/compress/flate/inflate.go:357
	}
//line /usr/local/go/src/compress/flate/inflate.go:357
	// _ = "end of CoverTab[26342]"
//line /usr/local/go/src/compress/flate/inflate.go:357
	_go_fuzz_dep_.CoverTab[26343]++
							return f.err
//line /usr/local/go/src/compress/flate/inflate.go:358
	// _ = "end of CoverTab[26343]"
}

//line /usr/local/go/src/compress/flate/inflate.go:364
var codeOrder = [...]int{16, 17, 18, 0, 8, 7, 9, 6, 10, 5, 11, 4, 12, 3, 13, 2, 14, 1, 15}

func (f *decompressor) readHuffman() error {
//line /usr/local/go/src/compress/flate/inflate.go:366
	_go_fuzz_dep_.CoverTab[26346]++

							for f.nb < 5+5+4 {
//line /usr/local/go/src/compress/flate/inflate.go:368
		_go_fuzz_dep_.CoverTab[26356]++
								if err := f.moreBits(); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:369
			_go_fuzz_dep_.CoverTab[26357]++
									return err
//line /usr/local/go/src/compress/flate/inflate.go:370
			// _ = "end of CoverTab[26357]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:371
			_go_fuzz_dep_.CoverTab[26358]++
//line /usr/local/go/src/compress/flate/inflate.go:371
			// _ = "end of CoverTab[26358]"
//line /usr/local/go/src/compress/flate/inflate.go:371
		}
//line /usr/local/go/src/compress/flate/inflate.go:371
		// _ = "end of CoverTab[26356]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:372
	// _ = "end of CoverTab[26346]"
//line /usr/local/go/src/compress/flate/inflate.go:372
	_go_fuzz_dep_.CoverTab[26347]++
							nlit := int(f.b&0x1F) + 257
							if nlit > maxNumLit {
//line /usr/local/go/src/compress/flate/inflate.go:374
		_go_fuzz_dep_.CoverTab[26359]++
								return CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:375
		// _ = "end of CoverTab[26359]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:376
		_go_fuzz_dep_.CoverTab[26360]++
//line /usr/local/go/src/compress/flate/inflate.go:376
		// _ = "end of CoverTab[26360]"
//line /usr/local/go/src/compress/flate/inflate.go:376
	}
//line /usr/local/go/src/compress/flate/inflate.go:376
	// _ = "end of CoverTab[26347]"
//line /usr/local/go/src/compress/flate/inflate.go:376
	_go_fuzz_dep_.CoverTab[26348]++
							f.b >>= 5
							ndist := int(f.b&0x1F) + 1
							if ndist > maxNumDist {
//line /usr/local/go/src/compress/flate/inflate.go:379
		_go_fuzz_dep_.CoverTab[26361]++
								return CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:380
		// _ = "end of CoverTab[26361]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:381
		_go_fuzz_dep_.CoverTab[26362]++
//line /usr/local/go/src/compress/flate/inflate.go:381
		// _ = "end of CoverTab[26362]"
//line /usr/local/go/src/compress/flate/inflate.go:381
	}
//line /usr/local/go/src/compress/flate/inflate.go:381
	// _ = "end of CoverTab[26348]"
//line /usr/local/go/src/compress/flate/inflate.go:381
	_go_fuzz_dep_.CoverTab[26349]++
							f.b >>= 5
							nclen := int(f.b&0xF) + 4

							f.b >>= 4
							f.nb -= 5 + 5 + 4

//line /usr/local/go/src/compress/flate/inflate.go:389
	for i := 0; i < nclen; i++ {
//line /usr/local/go/src/compress/flate/inflate.go:389
		_go_fuzz_dep_.CoverTab[26363]++
								for f.nb < 3 {
//line /usr/local/go/src/compress/flate/inflate.go:390
			_go_fuzz_dep_.CoverTab[26365]++
									if err := f.moreBits(); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:391
				_go_fuzz_dep_.CoverTab[26366]++
										return err
//line /usr/local/go/src/compress/flate/inflate.go:392
				// _ = "end of CoverTab[26366]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:393
				_go_fuzz_dep_.CoverTab[26367]++
//line /usr/local/go/src/compress/flate/inflate.go:393
				// _ = "end of CoverTab[26367]"
//line /usr/local/go/src/compress/flate/inflate.go:393
			}
//line /usr/local/go/src/compress/flate/inflate.go:393
			// _ = "end of CoverTab[26365]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:394
		// _ = "end of CoverTab[26363]"
//line /usr/local/go/src/compress/flate/inflate.go:394
		_go_fuzz_dep_.CoverTab[26364]++
								f.codebits[codeOrder[i]] = int(f.b & 0x7)
								f.b >>= 3
								f.nb -= 3
//line /usr/local/go/src/compress/flate/inflate.go:397
		// _ = "end of CoverTab[26364]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:398
	// _ = "end of CoverTab[26349]"
//line /usr/local/go/src/compress/flate/inflate.go:398
	_go_fuzz_dep_.CoverTab[26350]++
							for i := nclen; i < len(codeOrder); i++ {
//line /usr/local/go/src/compress/flate/inflate.go:399
		_go_fuzz_dep_.CoverTab[26368]++
								f.codebits[codeOrder[i]] = 0
//line /usr/local/go/src/compress/flate/inflate.go:400
		// _ = "end of CoverTab[26368]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:401
	// _ = "end of CoverTab[26350]"
//line /usr/local/go/src/compress/flate/inflate.go:401
	_go_fuzz_dep_.CoverTab[26351]++
							if !f.h1.init(f.codebits[0:]) {
//line /usr/local/go/src/compress/flate/inflate.go:402
		_go_fuzz_dep_.CoverTab[26369]++
								return CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:403
		// _ = "end of CoverTab[26369]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:404
		_go_fuzz_dep_.CoverTab[26370]++
//line /usr/local/go/src/compress/flate/inflate.go:404
		// _ = "end of CoverTab[26370]"
//line /usr/local/go/src/compress/flate/inflate.go:404
	}
//line /usr/local/go/src/compress/flate/inflate.go:404
	// _ = "end of CoverTab[26351]"
//line /usr/local/go/src/compress/flate/inflate.go:404
	_go_fuzz_dep_.CoverTab[26352]++

//line /usr/local/go/src/compress/flate/inflate.go:408
	for i, n := 0, nlit+ndist; i < n; {
//line /usr/local/go/src/compress/flate/inflate.go:408
		_go_fuzz_dep_.CoverTab[26371]++
								x, err := f.huffSym(&f.h1)
								if err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:410
			_go_fuzz_dep_.CoverTab[26377]++
									return err
//line /usr/local/go/src/compress/flate/inflate.go:411
			// _ = "end of CoverTab[26377]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:412
			_go_fuzz_dep_.CoverTab[26378]++
//line /usr/local/go/src/compress/flate/inflate.go:412
			// _ = "end of CoverTab[26378]"
//line /usr/local/go/src/compress/flate/inflate.go:412
		}
//line /usr/local/go/src/compress/flate/inflate.go:412
		// _ = "end of CoverTab[26371]"
//line /usr/local/go/src/compress/flate/inflate.go:412
		_go_fuzz_dep_.CoverTab[26372]++
								if x < 16 {
//line /usr/local/go/src/compress/flate/inflate.go:413
			_go_fuzz_dep_.CoverTab[26379]++

									f.bits[i] = x
									i++
									continue
//line /usr/local/go/src/compress/flate/inflate.go:417
			// _ = "end of CoverTab[26379]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:418
			_go_fuzz_dep_.CoverTab[26380]++
//line /usr/local/go/src/compress/flate/inflate.go:418
			// _ = "end of CoverTab[26380]"
//line /usr/local/go/src/compress/flate/inflate.go:418
		}
//line /usr/local/go/src/compress/flate/inflate.go:418
		// _ = "end of CoverTab[26372]"
//line /usr/local/go/src/compress/flate/inflate.go:418
		_go_fuzz_dep_.CoverTab[26373]++
		// Repeat previous length or zero.
		var rep int
		var nb uint
		var b int
		switch x {
		default:
//line /usr/local/go/src/compress/flate/inflate.go:424
			_go_fuzz_dep_.CoverTab[26381]++
									return InternalError("unexpected length code")
//line /usr/local/go/src/compress/flate/inflate.go:425
			// _ = "end of CoverTab[26381]"
		case 16:
//line /usr/local/go/src/compress/flate/inflate.go:426
			_go_fuzz_dep_.CoverTab[26382]++
									rep = 3
									nb = 2
									if i == 0 {
//line /usr/local/go/src/compress/flate/inflate.go:429
				_go_fuzz_dep_.CoverTab[26386]++
										return CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:430
				// _ = "end of CoverTab[26386]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:431
				_go_fuzz_dep_.CoverTab[26387]++
//line /usr/local/go/src/compress/flate/inflate.go:431
				// _ = "end of CoverTab[26387]"
//line /usr/local/go/src/compress/flate/inflate.go:431
			}
//line /usr/local/go/src/compress/flate/inflate.go:431
			// _ = "end of CoverTab[26382]"
//line /usr/local/go/src/compress/flate/inflate.go:431
			_go_fuzz_dep_.CoverTab[26383]++
									b = f.bits[i-1]
//line /usr/local/go/src/compress/flate/inflate.go:432
			// _ = "end of CoverTab[26383]"
		case 17:
//line /usr/local/go/src/compress/flate/inflate.go:433
			_go_fuzz_dep_.CoverTab[26384]++
									rep = 3
									nb = 3
									b = 0
//line /usr/local/go/src/compress/flate/inflate.go:436
			// _ = "end of CoverTab[26384]"
		case 18:
//line /usr/local/go/src/compress/flate/inflate.go:437
			_go_fuzz_dep_.CoverTab[26385]++
									rep = 11
									nb = 7
									b = 0
//line /usr/local/go/src/compress/flate/inflate.go:440
			// _ = "end of CoverTab[26385]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:441
		// _ = "end of CoverTab[26373]"
//line /usr/local/go/src/compress/flate/inflate.go:441
		_go_fuzz_dep_.CoverTab[26374]++
								for f.nb < nb {
//line /usr/local/go/src/compress/flate/inflate.go:442
			_go_fuzz_dep_.CoverTab[26388]++
									if err := f.moreBits(); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:443
				_go_fuzz_dep_.CoverTab[26389]++
										return err
//line /usr/local/go/src/compress/flate/inflate.go:444
				// _ = "end of CoverTab[26389]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:445
				_go_fuzz_dep_.CoverTab[26390]++
//line /usr/local/go/src/compress/flate/inflate.go:445
				// _ = "end of CoverTab[26390]"
//line /usr/local/go/src/compress/flate/inflate.go:445
			}
//line /usr/local/go/src/compress/flate/inflate.go:445
			// _ = "end of CoverTab[26388]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:446
		// _ = "end of CoverTab[26374]"
//line /usr/local/go/src/compress/flate/inflate.go:446
		_go_fuzz_dep_.CoverTab[26375]++
								rep += int(f.b & uint32(1<<nb-1))
								f.b >>= nb
								f.nb -= nb
								if i+rep > n {
//line /usr/local/go/src/compress/flate/inflate.go:450
			_go_fuzz_dep_.CoverTab[26391]++
									return CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:451
			// _ = "end of CoverTab[26391]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:452
			_go_fuzz_dep_.CoverTab[26392]++
//line /usr/local/go/src/compress/flate/inflate.go:452
			// _ = "end of CoverTab[26392]"
//line /usr/local/go/src/compress/flate/inflate.go:452
		}
//line /usr/local/go/src/compress/flate/inflate.go:452
		// _ = "end of CoverTab[26375]"
//line /usr/local/go/src/compress/flate/inflate.go:452
		_go_fuzz_dep_.CoverTab[26376]++
								for j := 0; j < rep; j++ {
//line /usr/local/go/src/compress/flate/inflate.go:453
			_go_fuzz_dep_.CoverTab[26393]++
									f.bits[i] = b
									i++
//line /usr/local/go/src/compress/flate/inflate.go:455
			// _ = "end of CoverTab[26393]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:456
		// _ = "end of CoverTab[26376]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:457
	// _ = "end of CoverTab[26352]"
//line /usr/local/go/src/compress/flate/inflate.go:457
	_go_fuzz_dep_.CoverTab[26353]++

							if !f.h1.init(f.bits[0:nlit]) || func() bool {
//line /usr/local/go/src/compress/flate/inflate.go:459
		_go_fuzz_dep_.CoverTab[26394]++
//line /usr/local/go/src/compress/flate/inflate.go:459
		return !f.h2.init(f.bits[nlit : nlit+ndist])
//line /usr/local/go/src/compress/flate/inflate.go:459
		// _ = "end of CoverTab[26394]"
//line /usr/local/go/src/compress/flate/inflate.go:459
	}() {
//line /usr/local/go/src/compress/flate/inflate.go:459
		_go_fuzz_dep_.CoverTab[26395]++
								return CorruptInputError(f.roffset)
//line /usr/local/go/src/compress/flate/inflate.go:460
		// _ = "end of CoverTab[26395]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:461
		_go_fuzz_dep_.CoverTab[26396]++
//line /usr/local/go/src/compress/flate/inflate.go:461
		// _ = "end of CoverTab[26396]"
//line /usr/local/go/src/compress/flate/inflate.go:461
	}
//line /usr/local/go/src/compress/flate/inflate.go:461
	// _ = "end of CoverTab[26353]"
//line /usr/local/go/src/compress/flate/inflate.go:461
	_go_fuzz_dep_.CoverTab[26354]++

//line /usr/local/go/src/compress/flate/inflate.go:467
	if f.h1.min < f.bits[endBlockMarker] {
//line /usr/local/go/src/compress/flate/inflate.go:467
		_go_fuzz_dep_.CoverTab[26397]++
								f.h1.min = f.bits[endBlockMarker]
//line /usr/local/go/src/compress/flate/inflate.go:468
		// _ = "end of CoverTab[26397]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:469
		_go_fuzz_dep_.CoverTab[26398]++
//line /usr/local/go/src/compress/flate/inflate.go:469
		// _ = "end of CoverTab[26398]"
//line /usr/local/go/src/compress/flate/inflate.go:469
	}
//line /usr/local/go/src/compress/flate/inflate.go:469
	// _ = "end of CoverTab[26354]"
//line /usr/local/go/src/compress/flate/inflate.go:469
	_go_fuzz_dep_.CoverTab[26355]++

							return nil
//line /usr/local/go/src/compress/flate/inflate.go:471
	// _ = "end of CoverTab[26355]"
}

// Decode a single Huffman block from f.
//line /usr/local/go/src/compress/flate/inflate.go:474
// hl and hd are the Huffman states for the lit/length values
//line /usr/local/go/src/compress/flate/inflate.go:474
// and the distance values, respectively. If hd == nil, using the
//line /usr/local/go/src/compress/flate/inflate.go:474
// fixed distance encoding associated with fixed Huffman blocks.
//line /usr/local/go/src/compress/flate/inflate.go:478
func (f *decompressor) huffmanBlock() {
//line /usr/local/go/src/compress/flate/inflate.go:478
	_go_fuzz_dep_.CoverTab[26399]++
							const (
		stateInit	= iota	// Zero value must be stateInit
		stateDict
	)

	switch f.stepState {
	case stateInit:
//line /usr/local/go/src/compress/flate/inflate.go:485
		_go_fuzz_dep_.CoverTab[26402]++
								goto readLiteral
//line /usr/local/go/src/compress/flate/inflate.go:486
		// _ = "end of CoverTab[26402]"
	case stateDict:
//line /usr/local/go/src/compress/flate/inflate.go:487
		_go_fuzz_dep_.CoverTab[26403]++
								goto copyHistory
//line /usr/local/go/src/compress/flate/inflate.go:488
		// _ = "end of CoverTab[26403]"
//line /usr/local/go/src/compress/flate/inflate.go:488
	default:
//line /usr/local/go/src/compress/flate/inflate.go:488
		_go_fuzz_dep_.CoverTab[26404]++
//line /usr/local/go/src/compress/flate/inflate.go:488
		// _ = "end of CoverTab[26404]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:489
	// _ = "end of CoverTab[26399]"
//line /usr/local/go/src/compress/flate/inflate.go:489
	_go_fuzz_dep_.CoverTab[26400]++

readLiteral:

	{
//line /usr/local/go/src/compress/flate/inflate.go:493
		_go_fuzz_dep_.CoverTab[26405]++
								v, err := f.huffSym(f.hl)
								if err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:495
			_go_fuzz_dep_.CoverTab[26412]++
									f.err = err
									return
//line /usr/local/go/src/compress/flate/inflate.go:497
			// _ = "end of CoverTab[26412]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:498
			_go_fuzz_dep_.CoverTab[26413]++
//line /usr/local/go/src/compress/flate/inflate.go:498
			// _ = "end of CoverTab[26413]"
//line /usr/local/go/src/compress/flate/inflate.go:498
		}
//line /usr/local/go/src/compress/flate/inflate.go:498
		// _ = "end of CoverTab[26405]"
//line /usr/local/go/src/compress/flate/inflate.go:498
		_go_fuzz_dep_.CoverTab[26406]++
								var n uint	// number of bits extra
								var length int
								switch {
		case v < 256:
//line /usr/local/go/src/compress/flate/inflate.go:502
			_go_fuzz_dep_.CoverTab[26414]++
									f.dict.writeByte(byte(v))
									if f.dict.availWrite() == 0 {
//line /usr/local/go/src/compress/flate/inflate.go:504
				_go_fuzz_dep_.CoverTab[26425]++
										f.toRead = f.dict.readFlush()
										f.step = (*decompressor).huffmanBlock
										f.stepState = stateInit
										return
//line /usr/local/go/src/compress/flate/inflate.go:508
				// _ = "end of CoverTab[26425]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:509
				_go_fuzz_dep_.CoverTab[26426]++
//line /usr/local/go/src/compress/flate/inflate.go:509
				// _ = "end of CoverTab[26426]"
//line /usr/local/go/src/compress/flate/inflate.go:509
			}
//line /usr/local/go/src/compress/flate/inflate.go:509
			// _ = "end of CoverTab[26414]"
//line /usr/local/go/src/compress/flate/inflate.go:509
			_go_fuzz_dep_.CoverTab[26415]++
									goto readLiteral
//line /usr/local/go/src/compress/flate/inflate.go:510
			// _ = "end of CoverTab[26415]"
		case v == 256:
//line /usr/local/go/src/compress/flate/inflate.go:511
			_go_fuzz_dep_.CoverTab[26416]++
									f.finishBlock()
									return
//line /usr/local/go/src/compress/flate/inflate.go:513
			// _ = "end of CoverTab[26416]"

		case v < 265:
//line /usr/local/go/src/compress/flate/inflate.go:515
			_go_fuzz_dep_.CoverTab[26417]++
									length = v - (257 - 3)
									n = 0
//line /usr/local/go/src/compress/flate/inflate.go:517
			// _ = "end of CoverTab[26417]"
		case v < 269:
//line /usr/local/go/src/compress/flate/inflate.go:518
			_go_fuzz_dep_.CoverTab[26418]++
									length = v*2 - (265*2 - 11)
									n = 1
//line /usr/local/go/src/compress/flate/inflate.go:520
			// _ = "end of CoverTab[26418]"
		case v < 273:
//line /usr/local/go/src/compress/flate/inflate.go:521
			_go_fuzz_dep_.CoverTab[26419]++
									length = v*4 - (269*4 - 19)
									n = 2
//line /usr/local/go/src/compress/flate/inflate.go:523
			// _ = "end of CoverTab[26419]"
		case v < 277:
//line /usr/local/go/src/compress/flate/inflate.go:524
			_go_fuzz_dep_.CoverTab[26420]++
									length = v*8 - (273*8 - 35)
									n = 3
//line /usr/local/go/src/compress/flate/inflate.go:526
			// _ = "end of CoverTab[26420]"
		case v < 281:
//line /usr/local/go/src/compress/flate/inflate.go:527
			_go_fuzz_dep_.CoverTab[26421]++
									length = v*16 - (277*16 - 67)
									n = 4
//line /usr/local/go/src/compress/flate/inflate.go:529
			// _ = "end of CoverTab[26421]"
		case v < 285:
//line /usr/local/go/src/compress/flate/inflate.go:530
			_go_fuzz_dep_.CoverTab[26422]++
									length = v*32 - (281*32 - 131)
									n = 5
//line /usr/local/go/src/compress/flate/inflate.go:532
			// _ = "end of CoverTab[26422]"
		case v < maxNumLit:
//line /usr/local/go/src/compress/flate/inflate.go:533
			_go_fuzz_dep_.CoverTab[26423]++
									length = 258
									n = 0
//line /usr/local/go/src/compress/flate/inflate.go:535
			// _ = "end of CoverTab[26423]"
		default:
//line /usr/local/go/src/compress/flate/inflate.go:536
			_go_fuzz_dep_.CoverTab[26424]++
									f.err = CorruptInputError(f.roffset)
									return
//line /usr/local/go/src/compress/flate/inflate.go:538
			// _ = "end of CoverTab[26424]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:539
		// _ = "end of CoverTab[26406]"
//line /usr/local/go/src/compress/flate/inflate.go:539
		_go_fuzz_dep_.CoverTab[26407]++
								if n > 0 {
//line /usr/local/go/src/compress/flate/inflate.go:540
			_go_fuzz_dep_.CoverTab[26427]++
									for f.nb < n {
//line /usr/local/go/src/compress/flate/inflate.go:541
				_go_fuzz_dep_.CoverTab[26429]++
										if err = f.moreBits(); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:542
					_go_fuzz_dep_.CoverTab[26430]++
											f.err = err
											return
//line /usr/local/go/src/compress/flate/inflate.go:544
					// _ = "end of CoverTab[26430]"
				} else {
//line /usr/local/go/src/compress/flate/inflate.go:545
					_go_fuzz_dep_.CoverTab[26431]++
//line /usr/local/go/src/compress/flate/inflate.go:545
					// _ = "end of CoverTab[26431]"
//line /usr/local/go/src/compress/flate/inflate.go:545
				}
//line /usr/local/go/src/compress/flate/inflate.go:545
				// _ = "end of CoverTab[26429]"
			}
//line /usr/local/go/src/compress/flate/inflate.go:546
			// _ = "end of CoverTab[26427]"
//line /usr/local/go/src/compress/flate/inflate.go:546
			_go_fuzz_dep_.CoverTab[26428]++
									length += int(f.b & uint32(1<<n-1))
									f.b >>= n
									f.nb -= n
//line /usr/local/go/src/compress/flate/inflate.go:549
			// _ = "end of CoverTab[26428]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:550
			_go_fuzz_dep_.CoverTab[26432]++
//line /usr/local/go/src/compress/flate/inflate.go:550
			// _ = "end of CoverTab[26432]"
//line /usr/local/go/src/compress/flate/inflate.go:550
		}
//line /usr/local/go/src/compress/flate/inflate.go:550
		// _ = "end of CoverTab[26407]"
//line /usr/local/go/src/compress/flate/inflate.go:550
		_go_fuzz_dep_.CoverTab[26408]++

								var dist int
								if f.hd == nil {
//line /usr/local/go/src/compress/flate/inflate.go:553
			_go_fuzz_dep_.CoverTab[26433]++
									for f.nb < 5 {
//line /usr/local/go/src/compress/flate/inflate.go:554
				_go_fuzz_dep_.CoverTab[26435]++
										if err = f.moreBits(); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:555
					_go_fuzz_dep_.CoverTab[26436]++
											f.err = err
											return
//line /usr/local/go/src/compress/flate/inflate.go:557
					// _ = "end of CoverTab[26436]"
				} else {
//line /usr/local/go/src/compress/flate/inflate.go:558
					_go_fuzz_dep_.CoverTab[26437]++
//line /usr/local/go/src/compress/flate/inflate.go:558
					// _ = "end of CoverTab[26437]"
//line /usr/local/go/src/compress/flate/inflate.go:558
				}
//line /usr/local/go/src/compress/flate/inflate.go:558
				// _ = "end of CoverTab[26435]"
			}
//line /usr/local/go/src/compress/flate/inflate.go:559
			// _ = "end of CoverTab[26433]"
//line /usr/local/go/src/compress/flate/inflate.go:559
			_go_fuzz_dep_.CoverTab[26434]++
									dist = int(bits.Reverse8(uint8(f.b & 0x1F << 3)))
									f.b >>= 5
									f.nb -= 5
//line /usr/local/go/src/compress/flate/inflate.go:562
			// _ = "end of CoverTab[26434]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:563
			_go_fuzz_dep_.CoverTab[26438]++
									if dist, err = f.huffSym(f.hd); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:564
				_go_fuzz_dep_.CoverTab[26439]++
										f.err = err
										return
//line /usr/local/go/src/compress/flate/inflate.go:566
				// _ = "end of CoverTab[26439]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:567
				_go_fuzz_dep_.CoverTab[26440]++
//line /usr/local/go/src/compress/flate/inflate.go:567
				// _ = "end of CoverTab[26440]"
//line /usr/local/go/src/compress/flate/inflate.go:567
			}
//line /usr/local/go/src/compress/flate/inflate.go:567
			// _ = "end of CoverTab[26438]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:568
		// _ = "end of CoverTab[26408]"
//line /usr/local/go/src/compress/flate/inflate.go:568
		_go_fuzz_dep_.CoverTab[26409]++

								switch {
		case dist < 4:
//line /usr/local/go/src/compress/flate/inflate.go:571
			_go_fuzz_dep_.CoverTab[26441]++
									dist++
//line /usr/local/go/src/compress/flate/inflate.go:572
			// _ = "end of CoverTab[26441]"
		case dist < maxNumDist:
//line /usr/local/go/src/compress/flate/inflate.go:573
			_go_fuzz_dep_.CoverTab[26442]++
									nb := uint(dist-2) >> 1

									extra := (dist & 1) << nb
									for f.nb < nb {
//line /usr/local/go/src/compress/flate/inflate.go:577
				_go_fuzz_dep_.CoverTab[26445]++
										if err = f.moreBits(); err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:578
					_go_fuzz_dep_.CoverTab[26446]++
											f.err = err
											return
//line /usr/local/go/src/compress/flate/inflate.go:580
					// _ = "end of CoverTab[26446]"
				} else {
//line /usr/local/go/src/compress/flate/inflate.go:581
					_go_fuzz_dep_.CoverTab[26447]++
//line /usr/local/go/src/compress/flate/inflate.go:581
					// _ = "end of CoverTab[26447]"
//line /usr/local/go/src/compress/flate/inflate.go:581
				}
//line /usr/local/go/src/compress/flate/inflate.go:581
				// _ = "end of CoverTab[26445]"
			}
//line /usr/local/go/src/compress/flate/inflate.go:582
			// _ = "end of CoverTab[26442]"
//line /usr/local/go/src/compress/flate/inflate.go:582
			_go_fuzz_dep_.CoverTab[26443]++
									extra |= int(f.b & uint32(1<<nb-1))
									f.b >>= nb
									f.nb -= nb
									dist = 1<<(nb+1) + 1 + extra
//line /usr/local/go/src/compress/flate/inflate.go:586
			// _ = "end of CoverTab[26443]"
		default:
//line /usr/local/go/src/compress/flate/inflate.go:587
			_go_fuzz_dep_.CoverTab[26444]++
									f.err = CorruptInputError(f.roffset)
									return
//line /usr/local/go/src/compress/flate/inflate.go:589
			// _ = "end of CoverTab[26444]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:590
		// _ = "end of CoverTab[26409]"
//line /usr/local/go/src/compress/flate/inflate.go:590
		_go_fuzz_dep_.CoverTab[26410]++

//line /usr/local/go/src/compress/flate/inflate.go:593
		if dist > f.dict.histSize() {
//line /usr/local/go/src/compress/flate/inflate.go:593
			_go_fuzz_dep_.CoverTab[26448]++
									f.err = CorruptInputError(f.roffset)
									return
//line /usr/local/go/src/compress/flate/inflate.go:595
			// _ = "end of CoverTab[26448]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:596
			_go_fuzz_dep_.CoverTab[26449]++
//line /usr/local/go/src/compress/flate/inflate.go:596
			// _ = "end of CoverTab[26449]"
//line /usr/local/go/src/compress/flate/inflate.go:596
		}
//line /usr/local/go/src/compress/flate/inflate.go:596
		// _ = "end of CoverTab[26410]"
//line /usr/local/go/src/compress/flate/inflate.go:596
		_go_fuzz_dep_.CoverTab[26411]++

								f.copyLen, f.copyDist = length, dist
								goto copyHistory
//line /usr/local/go/src/compress/flate/inflate.go:599
		// _ = "end of CoverTab[26411]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:600
	// _ = "end of CoverTab[26400]"
//line /usr/local/go/src/compress/flate/inflate.go:600
	_go_fuzz_dep_.CoverTab[26401]++

copyHistory:

	{
//line /usr/local/go/src/compress/flate/inflate.go:604
		_go_fuzz_dep_.CoverTab[26450]++
								cnt := f.dict.tryWriteCopy(f.copyDist, f.copyLen)
								if cnt == 0 {
//line /usr/local/go/src/compress/flate/inflate.go:606
			_go_fuzz_dep_.CoverTab[26453]++
									cnt = f.dict.writeCopy(f.copyDist, f.copyLen)
//line /usr/local/go/src/compress/flate/inflate.go:607
			// _ = "end of CoverTab[26453]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:608
			_go_fuzz_dep_.CoverTab[26454]++
//line /usr/local/go/src/compress/flate/inflate.go:608
			// _ = "end of CoverTab[26454]"
//line /usr/local/go/src/compress/flate/inflate.go:608
		}
//line /usr/local/go/src/compress/flate/inflate.go:608
		// _ = "end of CoverTab[26450]"
//line /usr/local/go/src/compress/flate/inflate.go:608
		_go_fuzz_dep_.CoverTab[26451]++
								f.copyLen -= cnt

								if f.dict.availWrite() == 0 || func() bool {
//line /usr/local/go/src/compress/flate/inflate.go:611
			_go_fuzz_dep_.CoverTab[26455]++
//line /usr/local/go/src/compress/flate/inflate.go:611
			return f.copyLen > 0
//line /usr/local/go/src/compress/flate/inflate.go:611
			// _ = "end of CoverTab[26455]"
//line /usr/local/go/src/compress/flate/inflate.go:611
		}() {
//line /usr/local/go/src/compress/flate/inflate.go:611
			_go_fuzz_dep_.CoverTab[26456]++
									f.toRead = f.dict.readFlush()
									f.step = (*decompressor).huffmanBlock
									f.stepState = stateDict
									return
//line /usr/local/go/src/compress/flate/inflate.go:615
			// _ = "end of CoverTab[26456]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:616
			_go_fuzz_dep_.CoverTab[26457]++
//line /usr/local/go/src/compress/flate/inflate.go:616
			// _ = "end of CoverTab[26457]"
//line /usr/local/go/src/compress/flate/inflate.go:616
		}
//line /usr/local/go/src/compress/flate/inflate.go:616
		// _ = "end of CoverTab[26451]"
//line /usr/local/go/src/compress/flate/inflate.go:616
		_go_fuzz_dep_.CoverTab[26452]++
								goto readLiteral
//line /usr/local/go/src/compress/flate/inflate.go:617
		// _ = "end of CoverTab[26452]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:618
	// _ = "end of CoverTab[26401]"
}

// Copy a single uncompressed data block from input to output.
func (f *decompressor) dataBlock() {
//line /usr/local/go/src/compress/flate/inflate.go:622
	_go_fuzz_dep_.CoverTab[26458]++

//line /usr/local/go/src/compress/flate/inflate.go:625
	f.nb = 0
							f.b = 0

//line /usr/local/go/src/compress/flate/inflate.go:629
	nr, err := io.ReadFull(f.r, f.buf[0:4])
	f.roffset += int64(nr)
	if err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:631
		_go_fuzz_dep_.CoverTab[26462]++
								f.err = noEOF(err)
								return
//line /usr/local/go/src/compress/flate/inflate.go:633
		// _ = "end of CoverTab[26462]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:634
		_go_fuzz_dep_.CoverTab[26463]++
//line /usr/local/go/src/compress/flate/inflate.go:634
		// _ = "end of CoverTab[26463]"
//line /usr/local/go/src/compress/flate/inflate.go:634
	}
//line /usr/local/go/src/compress/flate/inflate.go:634
	// _ = "end of CoverTab[26458]"
//line /usr/local/go/src/compress/flate/inflate.go:634
	_go_fuzz_dep_.CoverTab[26459]++
							n := int(f.buf[0]) | int(f.buf[1])<<8
							nn := int(f.buf[2]) | int(f.buf[3])<<8
							if uint16(nn) != uint16(^n) {
//line /usr/local/go/src/compress/flate/inflate.go:637
		_go_fuzz_dep_.CoverTab[26464]++
								f.err = CorruptInputError(f.roffset)
								return
//line /usr/local/go/src/compress/flate/inflate.go:639
		// _ = "end of CoverTab[26464]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:640
		_go_fuzz_dep_.CoverTab[26465]++
//line /usr/local/go/src/compress/flate/inflate.go:640
		// _ = "end of CoverTab[26465]"
//line /usr/local/go/src/compress/flate/inflate.go:640
	}
//line /usr/local/go/src/compress/flate/inflate.go:640
	// _ = "end of CoverTab[26459]"
//line /usr/local/go/src/compress/flate/inflate.go:640
	_go_fuzz_dep_.CoverTab[26460]++

							if n == 0 {
//line /usr/local/go/src/compress/flate/inflate.go:642
		_go_fuzz_dep_.CoverTab[26466]++
								f.toRead = f.dict.readFlush()
								f.finishBlock()
								return
//line /usr/local/go/src/compress/flate/inflate.go:645
		// _ = "end of CoverTab[26466]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:646
		_go_fuzz_dep_.CoverTab[26467]++
//line /usr/local/go/src/compress/flate/inflate.go:646
		// _ = "end of CoverTab[26467]"
//line /usr/local/go/src/compress/flate/inflate.go:646
	}
//line /usr/local/go/src/compress/flate/inflate.go:646
	// _ = "end of CoverTab[26460]"
//line /usr/local/go/src/compress/flate/inflate.go:646
	_go_fuzz_dep_.CoverTab[26461]++

							f.copyLen = n
							f.copyData()
//line /usr/local/go/src/compress/flate/inflate.go:649
	// _ = "end of CoverTab[26461]"
}

// copyData copies f.copyLen bytes from the underlying reader into f.hist.
//line /usr/local/go/src/compress/flate/inflate.go:652
// It pauses for reads when f.hist is full.
//line /usr/local/go/src/compress/flate/inflate.go:654
func (f *decompressor) copyData() {
//line /usr/local/go/src/compress/flate/inflate.go:654
	_go_fuzz_dep_.CoverTab[26468]++
							buf := f.dict.writeSlice()
							if len(buf) > f.copyLen {
//line /usr/local/go/src/compress/flate/inflate.go:656
		_go_fuzz_dep_.CoverTab[26472]++
								buf = buf[:f.copyLen]
//line /usr/local/go/src/compress/flate/inflate.go:657
		// _ = "end of CoverTab[26472]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:658
		_go_fuzz_dep_.CoverTab[26473]++
//line /usr/local/go/src/compress/flate/inflate.go:658
		// _ = "end of CoverTab[26473]"
//line /usr/local/go/src/compress/flate/inflate.go:658
	}
//line /usr/local/go/src/compress/flate/inflate.go:658
	// _ = "end of CoverTab[26468]"
//line /usr/local/go/src/compress/flate/inflate.go:658
	_go_fuzz_dep_.CoverTab[26469]++

							cnt, err := io.ReadFull(f.r, buf)
							f.roffset += int64(cnt)
							f.copyLen -= cnt
							f.dict.writeMark(cnt)
							if err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:664
		_go_fuzz_dep_.CoverTab[26474]++
								f.err = noEOF(err)
								return
//line /usr/local/go/src/compress/flate/inflate.go:666
		// _ = "end of CoverTab[26474]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:667
		_go_fuzz_dep_.CoverTab[26475]++
//line /usr/local/go/src/compress/flate/inflate.go:667
		// _ = "end of CoverTab[26475]"
//line /usr/local/go/src/compress/flate/inflate.go:667
	}
//line /usr/local/go/src/compress/flate/inflate.go:667
	// _ = "end of CoverTab[26469]"
//line /usr/local/go/src/compress/flate/inflate.go:667
	_go_fuzz_dep_.CoverTab[26470]++

							if f.dict.availWrite() == 0 || func() bool {
//line /usr/local/go/src/compress/flate/inflate.go:669
		_go_fuzz_dep_.CoverTab[26476]++
//line /usr/local/go/src/compress/flate/inflate.go:669
		return f.copyLen > 0
//line /usr/local/go/src/compress/flate/inflate.go:669
		// _ = "end of CoverTab[26476]"
//line /usr/local/go/src/compress/flate/inflate.go:669
	}() {
//line /usr/local/go/src/compress/flate/inflate.go:669
		_go_fuzz_dep_.CoverTab[26477]++
								f.toRead = f.dict.readFlush()
								f.step = (*decompressor).copyData
								return
//line /usr/local/go/src/compress/flate/inflate.go:672
		// _ = "end of CoverTab[26477]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:673
		_go_fuzz_dep_.CoverTab[26478]++
//line /usr/local/go/src/compress/flate/inflate.go:673
		// _ = "end of CoverTab[26478]"
//line /usr/local/go/src/compress/flate/inflate.go:673
	}
//line /usr/local/go/src/compress/flate/inflate.go:673
	// _ = "end of CoverTab[26470]"
//line /usr/local/go/src/compress/flate/inflate.go:673
	_go_fuzz_dep_.CoverTab[26471]++
							f.finishBlock()
//line /usr/local/go/src/compress/flate/inflate.go:674
	// _ = "end of CoverTab[26471]"
}

func (f *decompressor) finishBlock() {
//line /usr/local/go/src/compress/flate/inflate.go:677
	_go_fuzz_dep_.CoverTab[26479]++
							if f.final {
//line /usr/local/go/src/compress/flate/inflate.go:678
		_go_fuzz_dep_.CoverTab[26481]++
								if f.dict.availRead() > 0 {
//line /usr/local/go/src/compress/flate/inflate.go:679
			_go_fuzz_dep_.CoverTab[26483]++
									f.toRead = f.dict.readFlush()
//line /usr/local/go/src/compress/flate/inflate.go:680
			// _ = "end of CoverTab[26483]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:681
			_go_fuzz_dep_.CoverTab[26484]++
//line /usr/local/go/src/compress/flate/inflate.go:681
			// _ = "end of CoverTab[26484]"
//line /usr/local/go/src/compress/flate/inflate.go:681
		}
//line /usr/local/go/src/compress/flate/inflate.go:681
		// _ = "end of CoverTab[26481]"
//line /usr/local/go/src/compress/flate/inflate.go:681
		_go_fuzz_dep_.CoverTab[26482]++
								f.err = io.EOF
//line /usr/local/go/src/compress/flate/inflate.go:682
		// _ = "end of CoverTab[26482]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:683
		_go_fuzz_dep_.CoverTab[26485]++
//line /usr/local/go/src/compress/flate/inflate.go:683
		// _ = "end of CoverTab[26485]"
//line /usr/local/go/src/compress/flate/inflate.go:683
	}
//line /usr/local/go/src/compress/flate/inflate.go:683
	// _ = "end of CoverTab[26479]"
//line /usr/local/go/src/compress/flate/inflate.go:683
	_go_fuzz_dep_.CoverTab[26480]++
							f.step = (*decompressor).nextBlock
//line /usr/local/go/src/compress/flate/inflate.go:684
	// _ = "end of CoverTab[26480]"
}

// noEOF returns err, unless err == io.EOF, in which case it returns io.ErrUnexpectedEOF.
func noEOF(e error) error {
//line /usr/local/go/src/compress/flate/inflate.go:688
	_go_fuzz_dep_.CoverTab[26486]++
							if e == io.EOF {
//line /usr/local/go/src/compress/flate/inflate.go:689
		_go_fuzz_dep_.CoverTab[26488]++
								return io.ErrUnexpectedEOF
//line /usr/local/go/src/compress/flate/inflate.go:690
		// _ = "end of CoverTab[26488]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:691
		_go_fuzz_dep_.CoverTab[26489]++
//line /usr/local/go/src/compress/flate/inflate.go:691
		// _ = "end of CoverTab[26489]"
//line /usr/local/go/src/compress/flate/inflate.go:691
	}
//line /usr/local/go/src/compress/flate/inflate.go:691
	// _ = "end of CoverTab[26486]"
//line /usr/local/go/src/compress/flate/inflate.go:691
	_go_fuzz_dep_.CoverTab[26487]++
							return e
//line /usr/local/go/src/compress/flate/inflate.go:692
	// _ = "end of CoverTab[26487]"
}

func (f *decompressor) moreBits() error {
//line /usr/local/go/src/compress/flate/inflate.go:695
	_go_fuzz_dep_.CoverTab[26490]++
							c, err := f.r.ReadByte()
							if err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:697
		_go_fuzz_dep_.CoverTab[26492]++
								return noEOF(err)
//line /usr/local/go/src/compress/flate/inflate.go:698
		// _ = "end of CoverTab[26492]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:699
		_go_fuzz_dep_.CoverTab[26493]++
//line /usr/local/go/src/compress/flate/inflate.go:699
		// _ = "end of CoverTab[26493]"
//line /usr/local/go/src/compress/flate/inflate.go:699
	}
//line /usr/local/go/src/compress/flate/inflate.go:699
	// _ = "end of CoverTab[26490]"
//line /usr/local/go/src/compress/flate/inflate.go:699
	_go_fuzz_dep_.CoverTab[26491]++
							f.roffset++
							f.b |= uint32(c) << f.nb
							f.nb += 8
							return nil
//line /usr/local/go/src/compress/flate/inflate.go:703
	// _ = "end of CoverTab[26491]"
}

// Read the next Huffman-encoded symbol from f according to h.
func (f *decompressor) huffSym(h *huffmanDecoder) (int, error) {
//line /usr/local/go/src/compress/flate/inflate.go:707
	_go_fuzz_dep_.CoverTab[26494]++

//line /usr/local/go/src/compress/flate/inflate.go:712
	n := uint(h.min)

//line /usr/local/go/src/compress/flate/inflate.go:716
	nb, b := f.nb, f.b
	for {
//line /usr/local/go/src/compress/flate/inflate.go:717
		_go_fuzz_dep_.CoverTab[26495]++
								for nb < n {
//line /usr/local/go/src/compress/flate/inflate.go:718
			_go_fuzz_dep_.CoverTab[26498]++
									c, err := f.r.ReadByte()
									if err != nil {
//line /usr/local/go/src/compress/flate/inflate.go:720
				_go_fuzz_dep_.CoverTab[26500]++
										f.b = b
										f.nb = nb
										return 0, noEOF(err)
//line /usr/local/go/src/compress/flate/inflate.go:723
				// _ = "end of CoverTab[26500]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:724
				_go_fuzz_dep_.CoverTab[26501]++
//line /usr/local/go/src/compress/flate/inflate.go:724
				// _ = "end of CoverTab[26501]"
//line /usr/local/go/src/compress/flate/inflate.go:724
			}
//line /usr/local/go/src/compress/flate/inflate.go:724
			// _ = "end of CoverTab[26498]"
//line /usr/local/go/src/compress/flate/inflate.go:724
			_go_fuzz_dep_.CoverTab[26499]++
									f.roffset++
									b |= uint32(c) << (nb & 31)
									nb += 8
//line /usr/local/go/src/compress/flate/inflate.go:727
			// _ = "end of CoverTab[26499]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:728
		// _ = "end of CoverTab[26495]"
//line /usr/local/go/src/compress/flate/inflate.go:728
		_go_fuzz_dep_.CoverTab[26496]++
								chunk := h.chunks[b&(huffmanNumChunks-1)]
								n = uint(chunk & huffmanCountMask)
								if n > huffmanChunkBits {
//line /usr/local/go/src/compress/flate/inflate.go:731
			_go_fuzz_dep_.CoverTab[26502]++
									chunk = h.links[chunk>>huffmanValueShift][(b>>huffmanChunkBits)&h.linkMask]
									n = uint(chunk & huffmanCountMask)
//line /usr/local/go/src/compress/flate/inflate.go:733
			// _ = "end of CoverTab[26502]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:734
			_go_fuzz_dep_.CoverTab[26503]++
//line /usr/local/go/src/compress/flate/inflate.go:734
			// _ = "end of CoverTab[26503]"
//line /usr/local/go/src/compress/flate/inflate.go:734
		}
//line /usr/local/go/src/compress/flate/inflate.go:734
		// _ = "end of CoverTab[26496]"
//line /usr/local/go/src/compress/flate/inflate.go:734
		_go_fuzz_dep_.CoverTab[26497]++
								if n <= nb {
//line /usr/local/go/src/compress/flate/inflate.go:735
			_go_fuzz_dep_.CoverTab[26504]++
									if n == 0 {
//line /usr/local/go/src/compress/flate/inflate.go:736
				_go_fuzz_dep_.CoverTab[26506]++
										f.b = b
										f.nb = nb
										f.err = CorruptInputError(f.roffset)
										return 0, f.err
//line /usr/local/go/src/compress/flate/inflate.go:740
				// _ = "end of CoverTab[26506]"
			} else {
//line /usr/local/go/src/compress/flate/inflate.go:741
				_go_fuzz_dep_.CoverTab[26507]++
//line /usr/local/go/src/compress/flate/inflate.go:741
				// _ = "end of CoverTab[26507]"
//line /usr/local/go/src/compress/flate/inflate.go:741
			}
//line /usr/local/go/src/compress/flate/inflate.go:741
			// _ = "end of CoverTab[26504]"
//line /usr/local/go/src/compress/flate/inflate.go:741
			_go_fuzz_dep_.CoverTab[26505]++
									f.b = b >> (n & 31)
									f.nb = nb - n
									return int(chunk >> huffmanValueShift), nil
//line /usr/local/go/src/compress/flate/inflate.go:744
			// _ = "end of CoverTab[26505]"
		} else {
//line /usr/local/go/src/compress/flate/inflate.go:745
			_go_fuzz_dep_.CoverTab[26508]++
//line /usr/local/go/src/compress/flate/inflate.go:745
			// _ = "end of CoverTab[26508]"
//line /usr/local/go/src/compress/flate/inflate.go:745
		}
//line /usr/local/go/src/compress/flate/inflate.go:745
		// _ = "end of CoverTab[26497]"
	}
//line /usr/local/go/src/compress/flate/inflate.go:746
	// _ = "end of CoverTab[26494]"
}

func makeReader(r io.Reader) Reader {
//line /usr/local/go/src/compress/flate/inflate.go:749
	_go_fuzz_dep_.CoverTab[26509]++
							if rr, ok := r.(Reader); ok {
//line /usr/local/go/src/compress/flate/inflate.go:750
		_go_fuzz_dep_.CoverTab[26511]++
								return rr
//line /usr/local/go/src/compress/flate/inflate.go:751
		// _ = "end of CoverTab[26511]"
	} else {
//line /usr/local/go/src/compress/flate/inflate.go:752
		_go_fuzz_dep_.CoverTab[26512]++
//line /usr/local/go/src/compress/flate/inflate.go:752
		// _ = "end of CoverTab[26512]"
//line /usr/local/go/src/compress/flate/inflate.go:752
	}
//line /usr/local/go/src/compress/flate/inflate.go:752
	// _ = "end of CoverTab[26509]"
//line /usr/local/go/src/compress/flate/inflate.go:752
	_go_fuzz_dep_.CoverTab[26510]++
							return bufio.NewReader(r)
//line /usr/local/go/src/compress/flate/inflate.go:753
	// _ = "end of CoverTab[26510]"
}

func fixedHuffmanDecoderInit() {
//line /usr/local/go/src/compress/flate/inflate.go:756
	_go_fuzz_dep_.CoverTab[26513]++
							fixedOnce.Do(func() {
//line /usr/local/go/src/compress/flate/inflate.go:757
		_go_fuzz_dep_.CoverTab[26514]++
		// These come from the RFC section 3.2.6.
		var bits [288]int
		for i := 0; i < 144; i++ {
//line /usr/local/go/src/compress/flate/inflate.go:760
			_go_fuzz_dep_.CoverTab[26519]++
									bits[i] = 8
//line /usr/local/go/src/compress/flate/inflate.go:761
			// _ = "end of CoverTab[26519]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:762
		// _ = "end of CoverTab[26514]"
//line /usr/local/go/src/compress/flate/inflate.go:762
		_go_fuzz_dep_.CoverTab[26515]++
								for i := 144; i < 256; i++ {
//line /usr/local/go/src/compress/flate/inflate.go:763
			_go_fuzz_dep_.CoverTab[26520]++
									bits[i] = 9
//line /usr/local/go/src/compress/flate/inflate.go:764
			// _ = "end of CoverTab[26520]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:765
		// _ = "end of CoverTab[26515]"
//line /usr/local/go/src/compress/flate/inflate.go:765
		_go_fuzz_dep_.CoverTab[26516]++
								for i := 256; i < 280; i++ {
//line /usr/local/go/src/compress/flate/inflate.go:766
			_go_fuzz_dep_.CoverTab[26521]++
									bits[i] = 7
//line /usr/local/go/src/compress/flate/inflate.go:767
			// _ = "end of CoverTab[26521]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:768
		// _ = "end of CoverTab[26516]"
//line /usr/local/go/src/compress/flate/inflate.go:768
		_go_fuzz_dep_.CoverTab[26517]++
								for i := 280; i < 288; i++ {
//line /usr/local/go/src/compress/flate/inflate.go:769
			_go_fuzz_dep_.CoverTab[26522]++
									bits[i] = 8
//line /usr/local/go/src/compress/flate/inflate.go:770
			// _ = "end of CoverTab[26522]"
		}
//line /usr/local/go/src/compress/flate/inflate.go:771
		// _ = "end of CoverTab[26517]"
//line /usr/local/go/src/compress/flate/inflate.go:771
		_go_fuzz_dep_.CoverTab[26518]++
								fixedHuffmanDecoder.init(bits[:])
//line /usr/local/go/src/compress/flate/inflate.go:772
		// _ = "end of CoverTab[26518]"
	})
//line /usr/local/go/src/compress/flate/inflate.go:773
	// _ = "end of CoverTab[26513]"
}

func (f *decompressor) Reset(r io.Reader, dict []byte) error {
//line /usr/local/go/src/compress/flate/inflate.go:776
	_go_fuzz_dep_.CoverTab[26523]++
							*f = decompressor{
		r:		makeReader(r),
		bits:		f.bits,
		codebits:	f.codebits,
		dict:		f.dict,
		step:		(*decompressor).nextBlock,
	}
							f.dict.init(maxMatchOffset, dict)
							return nil
//line /usr/local/go/src/compress/flate/inflate.go:785
	// _ = "end of CoverTab[26523]"
}

// NewReader returns a new ReadCloser that can be used
//line /usr/local/go/src/compress/flate/inflate.go:788
// to read the uncompressed version of r.
//line /usr/local/go/src/compress/flate/inflate.go:788
// If r does not also implement io.ByteReader,
//line /usr/local/go/src/compress/flate/inflate.go:788
// the decompressor may read more data than necessary from r.
//line /usr/local/go/src/compress/flate/inflate.go:788
// The reader returns io.EOF after the final block in the DEFLATE stream has
//line /usr/local/go/src/compress/flate/inflate.go:788
// been encountered. Any trailing data after the final block is ignored.
//line /usr/local/go/src/compress/flate/inflate.go:788
//
//line /usr/local/go/src/compress/flate/inflate.go:788
// The ReadCloser returned by NewReader also implements Resetter.
//line /usr/local/go/src/compress/flate/inflate.go:796
func NewReader(r io.Reader) io.ReadCloser {
//line /usr/local/go/src/compress/flate/inflate.go:796
	_go_fuzz_dep_.CoverTab[26524]++
							fixedHuffmanDecoderInit()

							var f decompressor
							f.r = makeReader(r)
							f.bits = new([maxNumLit + maxNumDist]int)
							f.codebits = new([numCodes]int)
							f.step = (*decompressor).nextBlock
							f.dict.init(maxMatchOffset, nil)
							return &f
//line /usr/local/go/src/compress/flate/inflate.go:805
	// _ = "end of CoverTab[26524]"
}

// NewReaderDict is like NewReader but initializes the reader
//line /usr/local/go/src/compress/flate/inflate.go:808
// with a preset dictionary. The returned Reader behaves as if
//line /usr/local/go/src/compress/flate/inflate.go:808
// the uncompressed data stream started with the given dictionary,
//line /usr/local/go/src/compress/flate/inflate.go:808
// which has already been read. NewReaderDict is typically used
//line /usr/local/go/src/compress/flate/inflate.go:808
// to read data compressed by NewWriterDict.
//line /usr/local/go/src/compress/flate/inflate.go:808
//
//line /usr/local/go/src/compress/flate/inflate.go:808
// The ReadCloser returned by NewReader also implements Resetter.
//line /usr/local/go/src/compress/flate/inflate.go:815
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser {
//line /usr/local/go/src/compress/flate/inflate.go:815
	_go_fuzz_dep_.CoverTab[26525]++
							fixedHuffmanDecoderInit()

							var f decompressor
							f.r = makeReader(r)
							f.bits = new([maxNumLit + maxNumDist]int)
							f.codebits = new([numCodes]int)
							f.step = (*decompressor).nextBlock
							f.dict.init(maxMatchOffset, dict)
							return &f
//line /usr/local/go/src/compress/flate/inflate.go:824
	// _ = "end of CoverTab[26525]"
}

//line /usr/local/go/src/compress/flate/inflate.go:825
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/inflate.go:825
var _ = _go_fuzz_dep_.CoverTab
