// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
package flate

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
import (
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
import (
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:5
)

import (
	"io"
)

const (
	// The largest offset code.
	offsetCodeCount	= 30

	// The special code used to mark the end of a block.
	endBlockMarker	= 256

	// The first length code.
	lengthCodesStart	= 257

	// The number of codegen codes.
	codegenCodeCount	= 19
	badCode			= 255

	// bufferFlushSize indicates the buffer size
	// after which bytes are flushed to the writer.
	// Should preferably be a multiple of 6, since
	// we accumulate 6 bytes between writes to the buffer.
	bufferFlushSize	= 240

	// bufferSize is the actual output byte buffer size.
	// It must have additional headroom for a flush
	// which can contain up to 8 bytes.
	bufferSize	= bufferFlushSize + 8
)

// The number of extra bits needed by length code X - LENGTH_CODES_START.
var lengthExtraBits = []int8{
	0, 0, 0,
	0, 0, 0, 0, 0, 1, 1, 1, 1, 2,
	2, 2, 2, 3, 3, 3, 3, 4, 4, 4,
	4, 5, 5, 5, 5, 0,
}

// The length indicated by length code X - LENGTH_CODES_START.
var lengthBase = []uint32{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 10,
	12, 14, 16, 20, 24, 28, 32, 40, 48, 56,
	64, 80, 96, 112, 128, 160, 192, 224, 255,
}

// offset code word extra bits.
var offsetExtraBits = []int8{
	0, 0, 0, 0, 1, 1, 2, 2, 3, 3,
	4, 4, 5, 5, 6, 6, 7, 7, 8, 8,
	9, 9, 10, 10, 11, 11, 12, 12, 13, 13,
}

var offsetBase = []uint32{
	0x000000, 0x000001, 0x000002, 0x000003, 0x000004,
	0x000006, 0x000008, 0x00000c, 0x000010, 0x000018,
	0x000020, 0x000030, 0x000040, 0x000060, 0x000080,
	0x0000c0, 0x000100, 0x000180, 0x000200, 0x000300,
	0x000400, 0x000600, 0x000800, 0x000c00, 0x001000,
	0x001800, 0x002000, 0x003000, 0x004000, 0x006000,
}

// The odd order in which the codegen code sizes are written.
var codegenOrder = []uint32{16, 17, 18, 0, 8, 7, 9, 6, 10, 5, 11, 4, 12, 3, 13, 2, 14, 1, 15}

type huffmanBitWriter struct {
	// writer is the underlying writer.
	// Do not use it directly; use the write method, which ensures
	// that Write errors are sticky.
	writer	io.Writer

	// Data waiting to be written is bytes[0:nbytes]
	// and then the low nbits of bits.  Data is always written
	// sequentially into the bytes array.
	bits		uint64
	nbits		uint
	bytes		[bufferSize]byte
	codegenFreq	[codegenCodeCount]int32
	nbytes		int
	literalFreq	[]int32
	offsetFreq	[]int32
	codegen		[]uint8
	literalEncoding	*huffmanEncoder
	offsetEncoding	*huffmanEncoder
	codegenEncoding	*huffmanEncoder
	err		error
}

func newHuffmanBitWriter(w io.Writer) *huffmanBitWriter {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:94
	_go_fuzz_dep_.CoverTab[26031]++
									return &huffmanBitWriter{
		writer:			w,
		literalFreq:		make([]int32, maxNumLit),
		offsetFreq:		make([]int32, offsetCodeCount),
		codegen:		make([]uint8, maxNumLit+offsetCodeCount+1),
		literalEncoding:	newHuffmanEncoder(maxNumLit),
		codegenEncoding:	newHuffmanEncoder(codegenCodeCount),
		offsetEncoding:		newHuffmanEncoder(offsetCodeCount),
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:103
	// _ = "end of CoverTab[26031]"
}

func (w *huffmanBitWriter) reset(writer io.Writer) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:106
	_go_fuzz_dep_.CoverTab[26032]++
									w.writer = writer
									w.bits, w.nbits, w.nbytes, w.err = 0, 0, 0, nil
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:108
	// _ = "end of CoverTab[26032]"
}

func (w *huffmanBitWriter) flush() {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:111
	_go_fuzz_dep_.CoverTab[26033]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:112
		_go_fuzz_dep_.CoverTab[26036]++
										w.nbits = 0
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:114
		// _ = "end of CoverTab[26036]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:115
		_go_fuzz_dep_.CoverTab[26037]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:115
		// _ = "end of CoverTab[26037]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:115
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:115
	// _ = "end of CoverTab[26033]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:115
	_go_fuzz_dep_.CoverTab[26034]++
									n := w.nbytes
									for w.nbits != 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:117
		_go_fuzz_dep_.CoverTab[26038]++
										w.bytes[n] = byte(w.bits)
										w.bits >>= 8
										if w.nbits > 8 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:120
			_go_fuzz_dep_.CoverTab[26040]++
											w.nbits -= 8
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:121
			// _ = "end of CoverTab[26040]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:122
			_go_fuzz_dep_.CoverTab[26041]++
											w.nbits = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:123
			// _ = "end of CoverTab[26041]"
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:124
		// _ = "end of CoverTab[26038]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:124
		_go_fuzz_dep_.CoverTab[26039]++
										n++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:125
		// _ = "end of CoverTab[26039]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:126
	// _ = "end of CoverTab[26034]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:126
	_go_fuzz_dep_.CoverTab[26035]++
									w.bits = 0
									w.write(w.bytes[:n])
									w.nbytes = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:129
	// _ = "end of CoverTab[26035]"
}

func (w *huffmanBitWriter) write(b []byte) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:132
	_go_fuzz_dep_.CoverTab[26042]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:133
		_go_fuzz_dep_.CoverTab[26044]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:134
		// _ = "end of CoverTab[26044]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:135
		_go_fuzz_dep_.CoverTab[26045]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:135
		// _ = "end of CoverTab[26045]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:135
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:135
	// _ = "end of CoverTab[26042]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:135
	_go_fuzz_dep_.CoverTab[26043]++
									_, w.err = w.writer.Write(b)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:136
	// _ = "end of CoverTab[26043]"
}

func (w *huffmanBitWriter) writeBits(b int32, nb uint) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:139
	_go_fuzz_dep_.CoverTab[26046]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:140
		_go_fuzz_dep_.CoverTab[26048]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:141
		// _ = "end of CoverTab[26048]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:142
		_go_fuzz_dep_.CoverTab[26049]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:142
		// _ = "end of CoverTab[26049]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:142
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:142
	// _ = "end of CoverTab[26046]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:142
	_go_fuzz_dep_.CoverTab[26047]++
									w.bits |= uint64(b) << w.nbits
									w.nbits += nb
									if w.nbits >= 48 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:145
		_go_fuzz_dep_.CoverTab[26050]++
										bits := w.bits
										w.bits >>= 48
										w.nbits -= 48
										n := w.nbytes
										bytes := w.bytes[n : n+6]
										bytes[0] = byte(bits)
										bytes[1] = byte(bits >> 8)
										bytes[2] = byte(bits >> 16)
										bytes[3] = byte(bits >> 24)
										bytes[4] = byte(bits >> 32)
										bytes[5] = byte(bits >> 40)
										n += 6
										if n >= bufferFlushSize {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:158
			_go_fuzz_dep_.CoverTab[26052]++
											w.write(w.bytes[:n])
											n = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:160
			// _ = "end of CoverTab[26052]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:161
			_go_fuzz_dep_.CoverTab[26053]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:161
			// _ = "end of CoverTab[26053]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:161
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:161
		// _ = "end of CoverTab[26050]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:161
		_go_fuzz_dep_.CoverTab[26051]++
										w.nbytes = n
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:162
		// _ = "end of CoverTab[26051]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:163
		_go_fuzz_dep_.CoverTab[26054]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:163
		// _ = "end of CoverTab[26054]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:163
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:163
	// _ = "end of CoverTab[26047]"
}

func (w *huffmanBitWriter) writeBytes(bytes []byte) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:166
	_go_fuzz_dep_.CoverTab[26055]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:167
		_go_fuzz_dep_.CoverTab[26060]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:168
		// _ = "end of CoverTab[26060]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:169
		_go_fuzz_dep_.CoverTab[26061]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:169
		// _ = "end of CoverTab[26061]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:169
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:169
	// _ = "end of CoverTab[26055]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:169
	_go_fuzz_dep_.CoverTab[26056]++
									n := w.nbytes
									if w.nbits&7 != 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:171
		_go_fuzz_dep_.CoverTab[26062]++
										w.err = InternalError("writeBytes with unfinished bits")
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:173
		// _ = "end of CoverTab[26062]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:174
		_go_fuzz_dep_.CoverTab[26063]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:174
		// _ = "end of CoverTab[26063]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:174
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:174
	// _ = "end of CoverTab[26056]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:174
	_go_fuzz_dep_.CoverTab[26057]++
									for w.nbits != 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:175
		_go_fuzz_dep_.CoverTab[26064]++
										w.bytes[n] = byte(w.bits)
										w.bits >>= 8
										w.nbits -= 8
										n++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:179
		// _ = "end of CoverTab[26064]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:180
	// _ = "end of CoverTab[26057]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:180
	_go_fuzz_dep_.CoverTab[26058]++
									if n != 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:181
		_go_fuzz_dep_.CoverTab[26065]++
										w.write(w.bytes[:n])
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:182
		// _ = "end of CoverTab[26065]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:183
		_go_fuzz_dep_.CoverTab[26066]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:183
		// _ = "end of CoverTab[26066]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:183
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:183
	// _ = "end of CoverTab[26058]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:183
	_go_fuzz_dep_.CoverTab[26059]++
									w.nbytes = 0
									w.write(bytes)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:185
	// _ = "end of CoverTab[26059]"
}

// RFC 1951 3.2.7 specifies a special run-length encoding for specifying
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
// the literal and offset lengths arrays (which are concatenated into a single
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
// array).  This method generates that run-length encoding.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
//
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
// The result is written into the codegen array, and the frequencies
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
// of each code is written into the codegenFreq array.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
// Codes 0-15 are single byte codes. Codes 16-18 are followed by additional
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
// information. Code badCode is an end marker
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
//
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
//	numLiterals      The number of literals in literalEncoding
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
//	numOffsets       The number of offsets in offsetEncoding
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:188
//	litenc, offenc   The literal and offset encoder to use
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:200
func (w *huffmanBitWriter) generateCodegen(numLiterals int, numOffsets int, litEnc, offEnc *huffmanEncoder) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:200
	_go_fuzz_dep_.CoverTab[26067]++
									for i := range w.codegenFreq {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:201
		_go_fuzz_dep_.CoverTab[26072]++
										w.codegenFreq[i] = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:202
		// _ = "end of CoverTab[26072]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:203
	// _ = "end of CoverTab[26067]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:203
	_go_fuzz_dep_.CoverTab[26068]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:208
	codegen := w.codegen

	cgnl := codegen[:numLiterals]
	for i := range cgnl {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:211
		_go_fuzz_dep_.CoverTab[26073]++
										cgnl[i] = uint8(litEnc.codes[i].len)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:212
		// _ = "end of CoverTab[26073]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:213
	// _ = "end of CoverTab[26068]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:213
	_go_fuzz_dep_.CoverTab[26069]++

									cgnl = codegen[numLiterals : numLiterals+numOffsets]
									for i := range cgnl {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:216
		_go_fuzz_dep_.CoverTab[26074]++
										cgnl[i] = uint8(offEnc.codes[i].len)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:217
		// _ = "end of CoverTab[26074]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:218
	// _ = "end of CoverTab[26069]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:218
	_go_fuzz_dep_.CoverTab[26070]++
									codegen[numLiterals+numOffsets] = badCode

									size := codegen[0]
									count := 1
									outIndex := 0
									for inIndex := 1; size != badCode; inIndex++ {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:224
		_go_fuzz_dep_.CoverTab[26075]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:227
		nextSize := codegen[inIndex]
		if nextSize == size {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:228
			_go_fuzz_dep_.CoverTab[26079]++
											count++
											continue
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:230
			// _ = "end of CoverTab[26079]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:231
			_go_fuzz_dep_.CoverTab[26080]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:231
			// _ = "end of CoverTab[26080]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:231
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:231
		// _ = "end of CoverTab[26075]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:231
		_go_fuzz_dep_.CoverTab[26076]++

										if size != 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:233
			_go_fuzz_dep_.CoverTab[26081]++
											codegen[outIndex] = size
											outIndex++
											w.codegenFreq[size]++
											count--
											for count >= 3 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:238
				_go_fuzz_dep_.CoverTab[26082]++
												n := 6
												if n > count {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:240
					_go_fuzz_dep_.CoverTab[26084]++
													n = count
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:241
					// _ = "end of CoverTab[26084]"
				} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:242
					_go_fuzz_dep_.CoverTab[26085]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:242
					// _ = "end of CoverTab[26085]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:242
				}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:242
				// _ = "end of CoverTab[26082]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:242
				_go_fuzz_dep_.CoverTab[26083]++
												codegen[outIndex] = 16
												outIndex++
												codegen[outIndex] = uint8(n - 3)
												outIndex++
												w.codegenFreq[16]++
												count -= n
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:248
				// _ = "end of CoverTab[26083]"
			}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:249
			// _ = "end of CoverTab[26081]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:250
			_go_fuzz_dep_.CoverTab[26086]++
											for count >= 11 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:251
				_go_fuzz_dep_.CoverTab[26088]++
												n := 138
												if n > count {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:253
					_go_fuzz_dep_.CoverTab[26090]++
													n = count
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:254
					// _ = "end of CoverTab[26090]"
				} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:255
					_go_fuzz_dep_.CoverTab[26091]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:255
					// _ = "end of CoverTab[26091]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:255
				}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:255
				// _ = "end of CoverTab[26088]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:255
				_go_fuzz_dep_.CoverTab[26089]++
												codegen[outIndex] = 18
												outIndex++
												codegen[outIndex] = uint8(n - 11)
												outIndex++
												w.codegenFreq[18]++
												count -= n
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:261
				// _ = "end of CoverTab[26089]"
			}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:262
			// _ = "end of CoverTab[26086]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:262
			_go_fuzz_dep_.CoverTab[26087]++
											if count >= 3 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:263
				_go_fuzz_dep_.CoverTab[26092]++

												codegen[outIndex] = 17
												outIndex++
												codegen[outIndex] = uint8(count - 3)
												outIndex++
												w.codegenFreq[17]++
												count = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:270
				// _ = "end of CoverTab[26092]"
			} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:271
				_go_fuzz_dep_.CoverTab[26093]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:271
				// _ = "end of CoverTab[26093]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:271
			}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:271
			// _ = "end of CoverTab[26087]"
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:272
		// _ = "end of CoverTab[26076]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:272
		_go_fuzz_dep_.CoverTab[26077]++
										count--
										for ; count >= 0; count-- {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:274
			_go_fuzz_dep_.CoverTab[26094]++
											codegen[outIndex] = size
											outIndex++
											w.codegenFreq[size]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:277
			// _ = "end of CoverTab[26094]"
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:278
		// _ = "end of CoverTab[26077]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:278
		_go_fuzz_dep_.CoverTab[26078]++

										size = nextSize
										count = 1
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:281
		// _ = "end of CoverTab[26078]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:282
	// _ = "end of CoverTab[26070]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:282
	_go_fuzz_dep_.CoverTab[26071]++

									codegen[outIndex] = badCode
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:284
	// _ = "end of CoverTab[26071]"
}

// dynamicSize returns the size of dynamically encoded data in bits.
func (w *huffmanBitWriter) dynamicSize(litEnc, offEnc *huffmanEncoder, extraBits int) (size, numCodegens int) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:288
	_go_fuzz_dep_.CoverTab[26095]++
									numCodegens = len(w.codegenFreq)
									for numCodegens > 4 && func() bool {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:290
		_go_fuzz_dep_.CoverTab[26097]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:290
		return w.codegenFreq[codegenOrder[numCodegens-1]] == 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:290
		// _ = "end of CoverTab[26097]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:290
	}() {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:290
		_go_fuzz_dep_.CoverTab[26098]++
										numCodegens--
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:291
		// _ = "end of CoverTab[26098]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:292
	// _ = "end of CoverTab[26095]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:292
	_go_fuzz_dep_.CoverTab[26096]++
									header := 3 + 5 + 5 + 4 + (3 * numCodegens) +
		w.codegenEncoding.bitLength(w.codegenFreq[:]) +
		int(w.codegenFreq[16])*2 +
		int(w.codegenFreq[17])*3 +
		int(w.codegenFreq[18])*7
	size = header +
		litEnc.bitLength(w.literalFreq) +
		offEnc.bitLength(w.offsetFreq) +
		extraBits

									return size, numCodegens
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:303
	// _ = "end of CoverTab[26096]"
}

// fixedSize returns the size of dynamically encoded data in bits.
func (w *huffmanBitWriter) fixedSize(extraBits int) int {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:307
	_go_fuzz_dep_.CoverTab[26099]++
									return 3 +
		fixedLiteralEncoding.bitLength(w.literalFreq) +
		fixedOffsetEncoding.bitLength(w.offsetFreq) +
		extraBits
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:311
	// _ = "end of CoverTab[26099]"
}

// storedSize calculates the stored size, including header.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:314
// The function returns the size in bits and whether the block
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:314
// fits inside a single block.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:317
func (w *huffmanBitWriter) storedSize(in []byte) (int, bool) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:317
	_go_fuzz_dep_.CoverTab[26100]++
									if in == nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:318
		_go_fuzz_dep_.CoverTab[26103]++
										return 0, false
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:319
		// _ = "end of CoverTab[26103]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:320
		_go_fuzz_dep_.CoverTab[26104]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:320
		// _ = "end of CoverTab[26104]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:320
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:320
	// _ = "end of CoverTab[26100]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:320
	_go_fuzz_dep_.CoverTab[26101]++
									if len(in) <= maxStoreBlockSize {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:321
		_go_fuzz_dep_.CoverTab[26105]++
										return (len(in) + 5) * 8, true
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:322
		// _ = "end of CoverTab[26105]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:323
		_go_fuzz_dep_.CoverTab[26106]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:323
		// _ = "end of CoverTab[26106]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:323
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:323
	// _ = "end of CoverTab[26101]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:323
	_go_fuzz_dep_.CoverTab[26102]++
									return 0, false
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:324
	// _ = "end of CoverTab[26102]"
}

func (w *huffmanBitWriter) writeCode(c hcode) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:327
	_go_fuzz_dep_.CoverTab[26107]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:328
		_go_fuzz_dep_.CoverTab[26109]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:329
		// _ = "end of CoverTab[26109]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:330
		_go_fuzz_dep_.CoverTab[26110]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:330
		// _ = "end of CoverTab[26110]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:330
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:330
	// _ = "end of CoverTab[26107]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:330
	_go_fuzz_dep_.CoverTab[26108]++
									w.bits |= uint64(c.code) << w.nbits
									w.nbits += uint(c.len)
									if w.nbits >= 48 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:333
		_go_fuzz_dep_.CoverTab[26111]++
										bits := w.bits
										w.bits >>= 48
										w.nbits -= 48
										n := w.nbytes
										bytes := w.bytes[n : n+6]
										bytes[0] = byte(bits)
										bytes[1] = byte(bits >> 8)
										bytes[2] = byte(bits >> 16)
										bytes[3] = byte(bits >> 24)
										bytes[4] = byte(bits >> 32)
										bytes[5] = byte(bits >> 40)
										n += 6
										if n >= bufferFlushSize {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:346
			_go_fuzz_dep_.CoverTab[26113]++
											w.write(w.bytes[:n])
											n = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:348
			// _ = "end of CoverTab[26113]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:349
			_go_fuzz_dep_.CoverTab[26114]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:349
			// _ = "end of CoverTab[26114]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:349
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:349
		// _ = "end of CoverTab[26111]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:349
		_go_fuzz_dep_.CoverTab[26112]++
										w.nbytes = n
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:350
		// _ = "end of CoverTab[26112]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:351
		_go_fuzz_dep_.CoverTab[26115]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:351
		// _ = "end of CoverTab[26115]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:351
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:351
	// _ = "end of CoverTab[26108]"
}

// Write the header of a dynamic Huffman block to the output stream.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:354
//
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:354
//	numLiterals  The number of literals specified in codegen
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:354
//	numOffsets   The number of offsets specified in codegen
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:354
//	numCodegens  The number of codegens used in codegen
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:359
func (w *huffmanBitWriter) writeDynamicHeader(numLiterals int, numOffsets int, numCodegens int, isEof bool) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:359
	_go_fuzz_dep_.CoverTab[26116]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:360
		_go_fuzz_dep_.CoverTab[26120]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:361
		// _ = "end of CoverTab[26120]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:362
		_go_fuzz_dep_.CoverTab[26121]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:362
		// _ = "end of CoverTab[26121]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:362
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:362
	// _ = "end of CoverTab[26116]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:362
	_go_fuzz_dep_.CoverTab[26117]++
									var firstBits int32 = 4
									if isEof {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:364
		_go_fuzz_dep_.CoverTab[26122]++
										firstBits = 5
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:365
		// _ = "end of CoverTab[26122]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:366
		_go_fuzz_dep_.CoverTab[26123]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:366
		// _ = "end of CoverTab[26123]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:366
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:366
	// _ = "end of CoverTab[26117]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:366
	_go_fuzz_dep_.CoverTab[26118]++
									w.writeBits(firstBits, 3)
									w.writeBits(int32(numLiterals-257), 5)
									w.writeBits(int32(numOffsets-1), 5)
									w.writeBits(int32(numCodegens-4), 4)

									for i := 0; i < numCodegens; i++ {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:372
		_go_fuzz_dep_.CoverTab[26124]++
										value := uint(w.codegenEncoding.codes[codegenOrder[i]].len)
										w.writeBits(int32(value), 3)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:374
		// _ = "end of CoverTab[26124]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:375
	// _ = "end of CoverTab[26118]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:375
	_go_fuzz_dep_.CoverTab[26119]++

									i := 0
									for {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:378
		_go_fuzz_dep_.CoverTab[26125]++
										var codeWord int = int(w.codegen[i])
										i++
										if codeWord == badCode {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:381
			_go_fuzz_dep_.CoverTab[26127]++
											break
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:382
			// _ = "end of CoverTab[26127]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:383
			_go_fuzz_dep_.CoverTab[26128]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:383
			// _ = "end of CoverTab[26128]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:383
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:383
		// _ = "end of CoverTab[26125]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:383
		_go_fuzz_dep_.CoverTab[26126]++
										w.writeCode(w.codegenEncoding.codes[uint32(codeWord)])

										switch codeWord {
		case 16:
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:387
			_go_fuzz_dep_.CoverTab[26129]++
											w.writeBits(int32(w.codegen[i]), 2)
											i++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:389
			// _ = "end of CoverTab[26129]"
		case 17:
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:390
			_go_fuzz_dep_.CoverTab[26130]++
											w.writeBits(int32(w.codegen[i]), 3)
											i++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:392
			// _ = "end of CoverTab[26130]"
		case 18:
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:393
			_go_fuzz_dep_.CoverTab[26131]++
											w.writeBits(int32(w.codegen[i]), 7)
											i++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:395
			// _ = "end of CoverTab[26131]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:395
		default:
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:395
			_go_fuzz_dep_.CoverTab[26132]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:395
			// _ = "end of CoverTab[26132]"
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:396
		// _ = "end of CoverTab[26126]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:397
	// _ = "end of CoverTab[26119]"
}

func (w *huffmanBitWriter) writeStoredHeader(length int, isEof bool) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:400
	_go_fuzz_dep_.CoverTab[26133]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:401
		_go_fuzz_dep_.CoverTab[26136]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:402
		// _ = "end of CoverTab[26136]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:403
		_go_fuzz_dep_.CoverTab[26137]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:403
		// _ = "end of CoverTab[26137]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:403
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:403
	// _ = "end of CoverTab[26133]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:403
	_go_fuzz_dep_.CoverTab[26134]++
									var flag int32
									if isEof {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:405
		_go_fuzz_dep_.CoverTab[26138]++
										flag = 1
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:406
		// _ = "end of CoverTab[26138]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:407
		_go_fuzz_dep_.CoverTab[26139]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:407
		// _ = "end of CoverTab[26139]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:407
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:407
	// _ = "end of CoverTab[26134]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:407
	_go_fuzz_dep_.CoverTab[26135]++
									w.writeBits(flag, 3)
									w.flush()
									w.writeBits(int32(length), 16)
									w.writeBits(int32(^uint16(length)), 16)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:411
	// _ = "end of CoverTab[26135]"
}

func (w *huffmanBitWriter) writeFixedHeader(isEof bool) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:414
	_go_fuzz_dep_.CoverTab[26140]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:415
		_go_fuzz_dep_.CoverTab[26143]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:416
		// _ = "end of CoverTab[26143]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:417
		_go_fuzz_dep_.CoverTab[26144]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:417
		// _ = "end of CoverTab[26144]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:417
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:417
	// _ = "end of CoverTab[26140]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:417
	_go_fuzz_dep_.CoverTab[26141]++
	// Indicate that we are a fixed Huffman block
	var value int32 = 2
	if isEof {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:420
		_go_fuzz_dep_.CoverTab[26145]++
										value = 3
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:421
		// _ = "end of CoverTab[26145]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:422
		_go_fuzz_dep_.CoverTab[26146]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:422
		// _ = "end of CoverTab[26146]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:422
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:422
	// _ = "end of CoverTab[26141]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:422
	_go_fuzz_dep_.CoverTab[26142]++
									w.writeBits(value, 3)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:423
	// _ = "end of CoverTab[26142]"
}

// writeBlock will write a block of tokens with the smallest encoding.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:426
// The original input can be supplied, and if the huffman encoded data
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:426
// is larger than the original bytes, the data will be written as a
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:426
// stored block.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:426
// If the input is nil, the tokens will always be Huffman encoded.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:431
func (w *huffmanBitWriter) writeBlock(tokens []token, eof bool, input []byte) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:431
	_go_fuzz_dep_.CoverTab[26147]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:432
		_go_fuzz_dep_.CoverTab[26153]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:433
		// _ = "end of CoverTab[26153]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:434
		_go_fuzz_dep_.CoverTab[26154]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:434
		// _ = "end of CoverTab[26154]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:434
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:434
	// _ = "end of CoverTab[26147]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:434
	_go_fuzz_dep_.CoverTab[26148]++

									tokens = append(tokens, endBlockMarker)
									numLiterals, numOffsets := w.indexTokens(tokens)

									var extraBits int
									storedSize, storable := w.storedSize(input)
									if storable {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:441
		_go_fuzz_dep_.CoverTab[26155]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:446
		for lengthCode := lengthCodesStart + 8; lengthCode < numLiterals; lengthCode++ {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:446
			_go_fuzz_dep_.CoverTab[26157]++

											extraBits += int(w.literalFreq[lengthCode]) * int(lengthExtraBits[lengthCode-lengthCodesStart])
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:448
			// _ = "end of CoverTab[26157]"
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:449
		// _ = "end of CoverTab[26155]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:449
		_go_fuzz_dep_.CoverTab[26156]++
										for offsetCode := 4; offsetCode < numOffsets; offsetCode++ {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:450
			_go_fuzz_dep_.CoverTab[26158]++

											extraBits += int(w.offsetFreq[offsetCode]) * int(offsetExtraBits[offsetCode])
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:452
			// _ = "end of CoverTab[26158]"
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:453
		// _ = "end of CoverTab[26156]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:454
		_go_fuzz_dep_.CoverTab[26159]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:454
		// _ = "end of CoverTab[26159]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:454
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:454
	// _ = "end of CoverTab[26148]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:454
	_go_fuzz_dep_.CoverTab[26149]++

	// Figure out smallest code.
	// Fixed Huffman baseline.
	var literalEncoding = fixedLiteralEncoding
	var offsetEncoding = fixedOffsetEncoding
	var size = w.fixedSize(extraBits)

									// Dynamic Huffman?
									var numCodegens int

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:467
	w.generateCodegen(numLiterals, numOffsets, w.literalEncoding, w.offsetEncoding)
	w.codegenEncoding.generate(w.codegenFreq[:], 7)
	dynamicSize, numCodegens := w.dynamicSize(w.literalEncoding, w.offsetEncoding, extraBits)

	if dynamicSize < size {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:471
		_go_fuzz_dep_.CoverTab[26160]++
										size = dynamicSize
										literalEncoding = w.literalEncoding
										offsetEncoding = w.offsetEncoding
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:474
		// _ = "end of CoverTab[26160]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:475
		_go_fuzz_dep_.CoverTab[26161]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:475
		// _ = "end of CoverTab[26161]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:475
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:475
	// _ = "end of CoverTab[26149]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:475
	_go_fuzz_dep_.CoverTab[26150]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:478
	if storable && func() bool {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:478
		_go_fuzz_dep_.CoverTab[26162]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:478
		return storedSize < size
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:478
		// _ = "end of CoverTab[26162]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:478
	}() {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:478
		_go_fuzz_dep_.CoverTab[26163]++
										w.writeStoredHeader(len(input), eof)
										w.writeBytes(input)
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:481
		// _ = "end of CoverTab[26163]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:482
		_go_fuzz_dep_.CoverTab[26164]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:482
		// _ = "end of CoverTab[26164]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:482
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:482
	// _ = "end of CoverTab[26150]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:482
	_go_fuzz_dep_.CoverTab[26151]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:485
	if literalEncoding == fixedLiteralEncoding {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:485
		_go_fuzz_dep_.CoverTab[26165]++
										w.writeFixedHeader(eof)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:486
		// _ = "end of CoverTab[26165]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:487
		_go_fuzz_dep_.CoverTab[26166]++
										w.writeDynamicHeader(numLiterals, numOffsets, numCodegens, eof)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:488
		// _ = "end of CoverTab[26166]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:489
	// _ = "end of CoverTab[26151]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:489
	_go_fuzz_dep_.CoverTab[26152]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:492
	w.writeTokens(tokens, literalEncoding.codes, offsetEncoding.codes)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:492
	// _ = "end of CoverTab[26152]"
}

// writeBlockDynamic encodes a block using a dynamic Huffman table.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:495
// This should be used if the symbols used have a disproportionate
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:495
// histogram distribution.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:495
// If input is supplied and the compression savings are below 1/16th of the
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:495
// input size the block is stored.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:500
func (w *huffmanBitWriter) writeBlockDynamic(tokens []token, eof bool, input []byte) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:500
	_go_fuzz_dep_.CoverTab[26167]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:501
		_go_fuzz_dep_.CoverTab[26170]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:502
		// _ = "end of CoverTab[26170]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:503
		_go_fuzz_dep_.CoverTab[26171]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:503
		// _ = "end of CoverTab[26171]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:503
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:503
	// _ = "end of CoverTab[26167]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:503
	_go_fuzz_dep_.CoverTab[26168]++

									tokens = append(tokens, endBlockMarker)
									numLiterals, numOffsets := w.indexTokens(tokens)

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:510
	w.generateCodegen(numLiterals, numOffsets, w.literalEncoding, w.offsetEncoding)
									w.codegenEncoding.generate(w.codegenFreq[:], 7)
									size, numCodegens := w.dynamicSize(w.literalEncoding, w.offsetEncoding, 0)

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:515
	if ssize, storable := w.storedSize(input); storable && func() bool {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:515
		_go_fuzz_dep_.CoverTab[26172]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:515
		return ssize < (size + size>>4)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:515
		// _ = "end of CoverTab[26172]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:515
	}() {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:515
		_go_fuzz_dep_.CoverTab[26173]++
										w.writeStoredHeader(len(input), eof)
										w.writeBytes(input)
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:518
		// _ = "end of CoverTab[26173]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:519
		_go_fuzz_dep_.CoverTab[26174]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:519
		// _ = "end of CoverTab[26174]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:519
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:519
	// _ = "end of CoverTab[26168]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:519
	_go_fuzz_dep_.CoverTab[26169]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:522
	w.writeDynamicHeader(numLiterals, numOffsets, numCodegens, eof)

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:525
	w.writeTokens(tokens, w.literalEncoding.codes, w.offsetEncoding.codes)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:525
	// _ = "end of CoverTab[26169]"
}

// indexTokens indexes a slice of tokens, and updates
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:528
// literalFreq and offsetFreq, and generates literalEncoding
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:528
// and offsetEncoding.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:528
// The number of literal and offset tokens is returned.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:532
func (w *huffmanBitWriter) indexTokens(tokens []token) (numLiterals, numOffsets int) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:532
	_go_fuzz_dep_.CoverTab[26175]++
									for i := range w.literalFreq {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:533
		_go_fuzz_dep_.CoverTab[26182]++
										w.literalFreq[i] = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:534
		// _ = "end of CoverTab[26182]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:535
	// _ = "end of CoverTab[26175]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:535
	_go_fuzz_dep_.CoverTab[26176]++
									for i := range w.offsetFreq {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:536
		_go_fuzz_dep_.CoverTab[26183]++
										w.offsetFreq[i] = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:537
		// _ = "end of CoverTab[26183]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:538
	// _ = "end of CoverTab[26176]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:538
	_go_fuzz_dep_.CoverTab[26177]++

									for _, t := range tokens {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:540
		_go_fuzz_dep_.CoverTab[26184]++
										if t < matchType {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:541
			_go_fuzz_dep_.CoverTab[26186]++
											w.literalFreq[t.literal()]++
											continue
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:543
			// _ = "end of CoverTab[26186]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:544
			_go_fuzz_dep_.CoverTab[26187]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:544
			// _ = "end of CoverTab[26187]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:544
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:544
		// _ = "end of CoverTab[26184]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:544
		_go_fuzz_dep_.CoverTab[26185]++
										length := t.length()
										offset := t.offset()
										w.literalFreq[lengthCodesStart+lengthCode(length)]++
										w.offsetFreq[offsetCode(offset)]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:548
		// _ = "end of CoverTab[26185]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:549
	// _ = "end of CoverTab[26177]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:549
	_go_fuzz_dep_.CoverTab[26178]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:552
	numLiterals = len(w.literalFreq)
	for w.literalFreq[numLiterals-1] == 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:553
		_go_fuzz_dep_.CoverTab[26188]++
										numLiterals--
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:554
		// _ = "end of CoverTab[26188]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:555
	// _ = "end of CoverTab[26178]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:555
	_go_fuzz_dep_.CoverTab[26179]++

									numOffsets = len(w.offsetFreq)
									for numOffsets > 0 && func() bool {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:558
		_go_fuzz_dep_.CoverTab[26189]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:558
		return w.offsetFreq[numOffsets-1] == 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:558
		// _ = "end of CoverTab[26189]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:558
	}() {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:558
		_go_fuzz_dep_.CoverTab[26190]++
										numOffsets--
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:559
		// _ = "end of CoverTab[26190]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:560
	// _ = "end of CoverTab[26179]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:560
	_go_fuzz_dep_.CoverTab[26180]++
									if numOffsets == 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:561
		_go_fuzz_dep_.CoverTab[26191]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:564
		w.offsetFreq[0] = 1
										numOffsets = 1
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:565
		// _ = "end of CoverTab[26191]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:566
		_go_fuzz_dep_.CoverTab[26192]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:566
		// _ = "end of CoverTab[26192]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:566
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:566
	// _ = "end of CoverTab[26180]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:566
	_go_fuzz_dep_.CoverTab[26181]++
									w.literalEncoding.generate(w.literalFreq, 15)
									w.offsetEncoding.generate(w.offsetFreq, 15)
									return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:569
	// _ = "end of CoverTab[26181]"
}

// writeTokens writes a slice of tokens to the output.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:572
// codes for literal and offset encoding must be supplied.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:574
func (w *huffmanBitWriter) writeTokens(tokens []token, leCodes, oeCodes []hcode) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:574
	_go_fuzz_dep_.CoverTab[26193]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:575
		_go_fuzz_dep_.CoverTab[26195]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:576
		// _ = "end of CoverTab[26195]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:577
		_go_fuzz_dep_.CoverTab[26196]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:577
		// _ = "end of CoverTab[26196]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:577
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:577
	// _ = "end of CoverTab[26193]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:577
	_go_fuzz_dep_.CoverTab[26194]++
									for _, t := range tokens {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:578
		_go_fuzz_dep_.CoverTab[26197]++
										if t < matchType {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:579
			_go_fuzz_dep_.CoverTab[26200]++
											w.writeCode(leCodes[t.literal()])
											continue
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:581
			// _ = "end of CoverTab[26200]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:582
			_go_fuzz_dep_.CoverTab[26201]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:582
			// _ = "end of CoverTab[26201]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:582
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:582
		// _ = "end of CoverTab[26197]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:582
		_go_fuzz_dep_.CoverTab[26198]++

										length := t.length()
										lengthCode := lengthCode(length)
										w.writeCode(leCodes[lengthCode+lengthCodesStart])
										extraLengthBits := uint(lengthExtraBits[lengthCode])
										if extraLengthBits > 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:588
			_go_fuzz_dep_.CoverTab[26202]++
											extraLength := int32(length - lengthBase[lengthCode])
											w.writeBits(extraLength, extraLengthBits)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:590
			// _ = "end of CoverTab[26202]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:591
			_go_fuzz_dep_.CoverTab[26203]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:591
			// _ = "end of CoverTab[26203]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:591
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:591
		// _ = "end of CoverTab[26198]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:591
		_go_fuzz_dep_.CoverTab[26199]++

										offset := t.offset()
										offsetCode := offsetCode(offset)
										w.writeCode(oeCodes[offsetCode])
										extraOffsetBits := uint(offsetExtraBits[offsetCode])
										if extraOffsetBits > 0 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:597
			_go_fuzz_dep_.CoverTab[26204]++
											extraOffset := int32(offset - offsetBase[offsetCode])
											w.writeBits(extraOffset, extraOffsetBits)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:599
			// _ = "end of CoverTab[26204]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:600
			_go_fuzz_dep_.CoverTab[26205]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:600
			// _ = "end of CoverTab[26205]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:600
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:600
		// _ = "end of CoverTab[26199]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:601
	// _ = "end of CoverTab[26194]"
}

// huffOffset is a static offset encoder used for huffman only encoding.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:604
// It can be reused since we will not be encoding offset values.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:606
var huffOffset *huffmanEncoder

func init() {
	offsetFreq := make([]int32, offsetCodeCount)
	offsetFreq[0] = 1
	huffOffset = newHuffmanEncoder(offsetCodeCount)
	huffOffset.generate(offsetFreq, 15)
}

// writeBlockHuff encodes a block of bytes as either
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:615
// Huffman encoded literals or uncompressed bytes if the
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:615
// results only gains very little from compression.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:618
func (w *huffmanBitWriter) writeBlockHuff(eof bool, input []byte) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:618
	_go_fuzz_dep_.CoverTab[26206]++
									if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:619
		_go_fuzz_dep_.CoverTab[26211]++
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:620
		// _ = "end of CoverTab[26211]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:621
		_go_fuzz_dep_.CoverTab[26212]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:621
		// _ = "end of CoverTab[26212]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:621
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:621
	// _ = "end of CoverTab[26206]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:621
	_go_fuzz_dep_.CoverTab[26207]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:624
	for i := range w.literalFreq {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:624
		_go_fuzz_dep_.CoverTab[26213]++
										w.literalFreq[i] = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:625
		// _ = "end of CoverTab[26213]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:626
	// _ = "end of CoverTab[26207]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:626
	_go_fuzz_dep_.CoverTab[26208]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:629
	histogram(input, w.literalFreq)

	w.literalFreq[endBlockMarker] = 1

	const numLiterals = endBlockMarker + 1
	w.offsetFreq[0] = 1
	const numOffsets = 1

	w.literalEncoding.generate(w.literalFreq, 15)

	// Figure out smallest code.
									// Always use dynamic Huffman or Store
									var numCodegens int

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:645
	w.generateCodegen(numLiterals, numOffsets, w.literalEncoding, huffOffset)
									w.codegenEncoding.generate(w.codegenFreq[:], 7)
									size, numCodegens := w.dynamicSize(w.literalEncoding, huffOffset, 0)

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:650
	if ssize, storable := w.storedSize(input); storable && func() bool {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:650
		_go_fuzz_dep_.CoverTab[26214]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:650
		return ssize < (size + size>>4)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:650
		// _ = "end of CoverTab[26214]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:650
	}() {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:650
		_go_fuzz_dep_.CoverTab[26215]++
										w.writeStoredHeader(len(input), eof)
										w.writeBytes(input)
										return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:653
		// _ = "end of CoverTab[26215]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:654
		_go_fuzz_dep_.CoverTab[26216]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:654
		// _ = "end of CoverTab[26216]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:654
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:654
	// _ = "end of CoverTab[26208]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:654
	_go_fuzz_dep_.CoverTab[26209]++

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:657
	w.writeDynamicHeader(numLiterals, numOffsets, numCodegens, eof)
	encoding := w.literalEncoding.codes[:257]
	n := w.nbytes
	for _, t := range input {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:660
		_go_fuzz_dep_.CoverTab[26217]++

										c := encoding[t]
										w.bits |= uint64(c.code) << w.nbits
										w.nbits += uint(c.len)
										if w.nbits < 48 {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:665
			_go_fuzz_dep_.CoverTab[26221]++
											continue
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:666
			// _ = "end of CoverTab[26221]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:667
			_go_fuzz_dep_.CoverTab[26222]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:667
			// _ = "end of CoverTab[26222]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:667
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:667
		// _ = "end of CoverTab[26217]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:667
		_go_fuzz_dep_.CoverTab[26218]++

										bits := w.bits
										w.bits >>= 48
										w.nbits -= 48
										bytes := w.bytes[n : n+6]
										bytes[0] = byte(bits)
										bytes[1] = byte(bits >> 8)
										bytes[2] = byte(bits >> 16)
										bytes[3] = byte(bits >> 24)
										bytes[4] = byte(bits >> 32)
										bytes[5] = byte(bits >> 40)
										n += 6
										if n < bufferFlushSize {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:680
			_go_fuzz_dep_.CoverTab[26223]++
											continue
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:681
			// _ = "end of CoverTab[26223]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:682
			_go_fuzz_dep_.CoverTab[26224]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:682
			// _ = "end of CoverTab[26224]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:682
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:682
		// _ = "end of CoverTab[26218]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:682
		_go_fuzz_dep_.CoverTab[26219]++
										w.write(w.bytes[:n])
										if w.err != nil {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:684
			_go_fuzz_dep_.CoverTab[26225]++
											return
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:685
			// _ = "end of CoverTab[26225]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:686
			_go_fuzz_dep_.CoverTab[26226]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:686
			// _ = "end of CoverTab[26226]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:686
		}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:686
		// _ = "end of CoverTab[26219]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:686
		_go_fuzz_dep_.CoverTab[26220]++
										n = 0
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:687
		// _ = "end of CoverTab[26220]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:688
	// _ = "end of CoverTab[26209]"
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:688
	_go_fuzz_dep_.CoverTab[26210]++
									w.nbytes = n
									w.writeCode(encoding[endBlockMarker])
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:690
	// _ = "end of CoverTab[26210]"
}

// histogram accumulates a histogram of b in h.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:693
//
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:693
// len(h) must be >= 256, and h's elements must be all zeroes.
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:696
func histogram(b []byte, h []int32) {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:696
	_go_fuzz_dep_.CoverTab[26227]++
									h = h[:256]
									for _, t := range b {
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:698
		_go_fuzz_dep_.CoverTab[26228]++
										h[t]++
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:699
		// _ = "end of CoverTab[26228]"
	}
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:700
	// _ = "end of CoverTab[26227]"
}

//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:701
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/huffman_bit_writer.go:701
var _ = _go_fuzz_dep_.CoverTab
