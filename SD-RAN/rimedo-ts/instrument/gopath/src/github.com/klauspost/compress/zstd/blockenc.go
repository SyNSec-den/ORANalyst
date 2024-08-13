// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:5
)

import (
	"errors"
	"fmt"
	"math"
	"math/bits"

	"github.com/klauspost/compress/huff0"
)

type blockEnc struct {
	size		int
	literals	[]byte
	sequences	[]seq
	coders		seqCoders
	litEnc		*huff0.Scratch
	dictLitEnc	*huff0.Scratch
	wr		bitWriter

	extraLits		int
	output			[]byte
	recentOffsets		[3]uint32
	prevRecentOffsets	[3]uint32

	last	bool
	lowMem	bool
}

// init should be used once the block has been created.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:34
// If called more than once, the effect is the same as calling reset.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:36
func (b *blockEnc) init() {
	if b.lowMem {

		if cap(b.literals) < 1<<10 {
			b.literals = make([]byte, 0, 1<<10)
		}
		const defSeqs = 20
		if cap(b.sequences) < defSeqs {
			b.sequences = make([]seq, 0, defSeqs)
		}

		if cap(b.output) < 1<<10 {
			b.output = make([]byte, 0, 1<<10)
		}
	} else {
		if cap(b.literals) < maxCompressedBlockSize {
			b.literals = make([]byte, 0, maxCompressedBlockSize)
		}
		const defSeqs = 2000
		if cap(b.sequences) < defSeqs {
			b.sequences = make([]seq, 0, defSeqs)
		}
		if cap(b.output) < maxCompressedBlockSize {
			b.output = make([]byte, 0, maxCompressedBlockSize)
		}
	}

	if b.coders.mlEnc == nil {
		b.coders.mlEnc = &fseEncoder{}
		b.coders.mlPrev = &fseEncoder{}
		b.coders.ofEnc = &fseEncoder{}
		b.coders.ofPrev = &fseEncoder{}
		b.coders.llEnc = &fseEncoder{}
		b.coders.llPrev = &fseEncoder{}
	}
	b.litEnc = &huff0.Scratch{WantLogLess: 4}
	b.reset(nil)
}

// initNewEncode can be used to reset offsets and encoders to the initial state.
func (b *blockEnc) initNewEncode() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:76
	_go_fuzz_dep_.CoverTab[91214]++
												b.recentOffsets = [3]uint32{1, 4, 8}
												b.litEnc.Reuse = huff0.ReusePolicyNone
												b.coders.setPrev(nil, nil, nil)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:79
	// _ = "end of CoverTab[91214]"
}

// reset will reset the block for a new encode, but in the same stream,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:82
// meaning that state will be carried over, but the block content is reset.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:82
// If a previous block is provided, the recent offsets are carried over.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:85
func (b *blockEnc) reset(prev *blockEnc) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:85
	_go_fuzz_dep_.CoverTab[91215]++
												b.extraLits = 0
												b.literals = b.literals[:0]
												b.size = 0
												b.sequences = b.sequences[:0]
												b.output = b.output[:0]
												b.last = false
												if prev != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:92
		_go_fuzz_dep_.CoverTab[91217]++
													b.recentOffsets = prev.prevRecentOffsets
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:93
		// _ = "end of CoverTab[91217]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:94
		_go_fuzz_dep_.CoverTab[91218]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:94
		// _ = "end of CoverTab[91218]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:94
	// _ = "end of CoverTab[91215]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:94
	_go_fuzz_dep_.CoverTab[91216]++
												b.dictLitEnc = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:95
	// _ = "end of CoverTab[91216]"
}

// reset will reset the block for a new encode, but in the same stream,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:98
// meaning that state will be carried over, but the block content is reset.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:98
// If a previous block is provided, the recent offsets are carried over.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:101
func (b *blockEnc) swapEncoders(prev *blockEnc) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:101
	_go_fuzz_dep_.CoverTab[91219]++
												b.coders.swap(&prev.coders)
												b.litEnc, prev.litEnc = prev.litEnc, b.litEnc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:103
	// _ = "end of CoverTab[91219]"
}

// blockHeader contains the information for a block header.
type blockHeader uint32

// setLast sets the 'last' indicator on a block.
func (h *blockHeader) setLast(b bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:110
	_go_fuzz_dep_.CoverTab[91220]++
												if b {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:111
		_go_fuzz_dep_.CoverTab[91221]++
													*h = *h | 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:112
		// _ = "end of CoverTab[91221]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:113
		_go_fuzz_dep_.CoverTab[91222]++
													const mask = (1 << 24) - 2
													*h = *h & mask
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:115
		// _ = "end of CoverTab[91222]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:116
	// _ = "end of CoverTab[91220]"
}

// setSize will store the compressed size of a block.
func (h *blockHeader) setSize(v uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:120
	_go_fuzz_dep_.CoverTab[91223]++
												const mask = 7
												*h = (*h)&mask | blockHeader(v<<3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:122
	// _ = "end of CoverTab[91223]"
}

// setType sets the block type.
func (h *blockHeader) setType(t blockType) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:126
	_go_fuzz_dep_.CoverTab[91224]++
												const mask = 1 | (((1 << 24) - 1) ^ 7)
												*h = (*h & mask) | blockHeader(t<<1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:128
	// _ = "end of CoverTab[91224]"
}

// appendTo will append the block header to a slice.
func (h blockHeader) appendTo(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:132
	_go_fuzz_dep_.CoverTab[91225]++
												return append(b, uint8(h), uint8(h>>8), uint8(h>>16))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:133
	// _ = "end of CoverTab[91225]"
}

// String returns a string representation of the block.
func (h blockHeader) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:137
	_go_fuzz_dep_.CoverTab[91226]++
												return fmt.Sprintf("Type: %d, Size: %d, Last:%t", (h>>1)&3, h>>3, h&1 == 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:138
	// _ = "end of CoverTab[91226]"
}

// literalsHeader contains literals header information.
type literalsHeader uint64

// setType can be used to set the type of literal block.
func (h *literalsHeader) setType(t literalsBlockType) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:145
	_go_fuzz_dep_.CoverTab[91227]++
												const mask = math.MaxUint64 - 3
												*h = (*h & mask) | literalsHeader(t)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:147
	// _ = "end of CoverTab[91227]"
}

// setSize can be used to set a single size, for uncompressed and RLE content.
func (h *literalsHeader) setSize(regenLen int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:151
	_go_fuzz_dep_.CoverTab[91228]++
												inBits := bits.Len32(uint32(regenLen))
	// Only retain 2 bits
	const mask = 3
	lh := uint64(*h & mask)
	switch {
	case inBits < 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:157
		_go_fuzz_dep_.CoverTab[91230]++
													lh |= (uint64(regenLen) << 3) | (1 << 60)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:159
			_go_fuzz_dep_.CoverTab[91234]++
														got := int(lh>>3) & 0xff
														if got != regenLen {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:161
				_go_fuzz_dep_.CoverTab[91235]++
															panic(fmt.Sprint("litRegenSize = ", regenLen, "(want) != ", got, "(got)"))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:162
				// _ = "end of CoverTab[91235]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:163
				_go_fuzz_dep_.CoverTab[91236]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:163
				// _ = "end of CoverTab[91236]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:163
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:163
			// _ = "end of CoverTab[91234]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:164
			_go_fuzz_dep_.CoverTab[91237]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:164
			// _ = "end of CoverTab[91237]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:164
		// _ = "end of CoverTab[91230]"
	case inBits < 12:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:165
		_go_fuzz_dep_.CoverTab[91231]++
													lh |= (1 << 2) | (uint64(regenLen) << 4) | (2 << 60)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:166
		// _ = "end of CoverTab[91231]"
	case inBits < 20:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:167
		_go_fuzz_dep_.CoverTab[91232]++
													lh |= (3 << 2) | (uint64(regenLen) << 4) | (3 << 60)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:168
		// _ = "end of CoverTab[91232]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:169
		_go_fuzz_dep_.CoverTab[91233]++
													panic(fmt.Errorf("internal error: block too big (%d)", regenLen))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:170
		// _ = "end of CoverTab[91233]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:171
	// _ = "end of CoverTab[91228]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:171
	_go_fuzz_dep_.CoverTab[91229]++
												*h = literalsHeader(lh)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:172
	// _ = "end of CoverTab[91229]"
}

// setSizes will set the size of a compressed literals section and the input length.
func (h *literalsHeader) setSizes(compLen, inLen int, single bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:176
	_go_fuzz_dep_.CoverTab[91238]++
												compBits, inBits := bits.Len32(uint32(compLen)), bits.Len32(uint32(inLen))
	// Only retain 2 bits
	const mask = 3
	lh := uint64(*h & mask)
	switch {
	case compBits <= 10 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:182
		_go_fuzz_dep_.CoverTab[91245]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:182
		return inBits <= 10
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:182
		// _ = "end of CoverTab[91245]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:182
	}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:182
		_go_fuzz_dep_.CoverTab[91240]++
													if !single {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:183
			_go_fuzz_dep_.CoverTab[91246]++
														lh |= 1 << 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:184
			// _ = "end of CoverTab[91246]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:185
			_go_fuzz_dep_.CoverTab[91247]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:185
			// _ = "end of CoverTab[91247]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:185
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:185
		// _ = "end of CoverTab[91240]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:185
		_go_fuzz_dep_.CoverTab[91241]++
													lh |= (uint64(inLen) << 4) | (uint64(compLen) << (10 + 4)) | (3 << 60)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:187
			_go_fuzz_dep_.CoverTab[91248]++
														const mmask = (1 << 24) - 1
														n := (lh >> 4) & mmask
														if int(n&1023) != inLen {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:190
				_go_fuzz_dep_.CoverTab[91250]++
															panic(fmt.Sprint("regensize:", int(n&1023), "!=", inLen, inBits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:191
				// _ = "end of CoverTab[91250]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:192
				_go_fuzz_dep_.CoverTab[91251]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:192
				// _ = "end of CoverTab[91251]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:192
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:192
			// _ = "end of CoverTab[91248]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:192
			_go_fuzz_dep_.CoverTab[91249]++
														if int(n>>10) != compLen {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:193
				_go_fuzz_dep_.CoverTab[91252]++
															panic(fmt.Sprint("compsize:", int(n>>10), "!=", compLen, compBits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:194
				// _ = "end of CoverTab[91252]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:195
				_go_fuzz_dep_.CoverTab[91253]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:195
				// _ = "end of CoverTab[91253]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:195
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:195
			// _ = "end of CoverTab[91249]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:196
			_go_fuzz_dep_.CoverTab[91254]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:196
			// _ = "end of CoverTab[91254]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:196
		// _ = "end of CoverTab[91241]"
	case compBits <= 14 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:197
		_go_fuzz_dep_.CoverTab[91255]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:197
		return inBits <= 14
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:197
		// _ = "end of CoverTab[91255]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:197
	}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:197
		_go_fuzz_dep_.CoverTab[91242]++
													lh |= (2 << 2) | (uint64(inLen) << 4) | (uint64(compLen) << (14 + 4)) | (4 << 60)
													if single {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:199
			_go_fuzz_dep_.CoverTab[91256]++
														panic("single stream used with more than 10 bits length.")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:200
			// _ = "end of CoverTab[91256]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:201
			_go_fuzz_dep_.CoverTab[91257]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:201
			// _ = "end of CoverTab[91257]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:201
		// _ = "end of CoverTab[91242]"
	case compBits <= 18 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:202
		_go_fuzz_dep_.CoverTab[91258]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:202
		return inBits <= 18
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:202
		// _ = "end of CoverTab[91258]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:202
	}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:202
		_go_fuzz_dep_.CoverTab[91243]++
													lh |= (3 << 2) | (uint64(inLen) << 4) | (uint64(compLen) << (18 + 4)) | (5 << 60)
													if single {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:204
			_go_fuzz_dep_.CoverTab[91259]++
														panic("single stream used with more than 10 bits length.")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:205
			// _ = "end of CoverTab[91259]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:206
			_go_fuzz_dep_.CoverTab[91260]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:206
			// _ = "end of CoverTab[91260]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:206
		// _ = "end of CoverTab[91243]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:207
		_go_fuzz_dep_.CoverTab[91244]++
													panic("internal error: block too big")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:208
		// _ = "end of CoverTab[91244]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:209
	// _ = "end of CoverTab[91238]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:209
	_go_fuzz_dep_.CoverTab[91239]++
												*h = literalsHeader(lh)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:210
	// _ = "end of CoverTab[91239]"
}

// appendTo will append the literals header to a byte slice.
func (h literalsHeader) appendTo(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:214
	_go_fuzz_dep_.CoverTab[91261]++
												size := uint8(h >> 60)
												switch size {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:217
		_go_fuzz_dep_.CoverTab[91263]++
													b = append(b, uint8(h))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:218
		// _ = "end of CoverTab[91263]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:219
		_go_fuzz_dep_.CoverTab[91264]++
													b = append(b, uint8(h), uint8(h>>8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:220
		// _ = "end of CoverTab[91264]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:221
		_go_fuzz_dep_.CoverTab[91265]++
													b = append(b, uint8(h), uint8(h>>8), uint8(h>>16))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:222
		// _ = "end of CoverTab[91265]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:223
		_go_fuzz_dep_.CoverTab[91266]++
													b = append(b, uint8(h), uint8(h>>8), uint8(h>>16), uint8(h>>24))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:224
		// _ = "end of CoverTab[91266]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:225
		_go_fuzz_dep_.CoverTab[91267]++
													b = append(b, uint8(h), uint8(h>>8), uint8(h>>16), uint8(h>>24), uint8(h>>32))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:226
		// _ = "end of CoverTab[91267]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:227
		_go_fuzz_dep_.CoverTab[91268]++
													panic(fmt.Errorf("internal error: literalsHeader has invalid size (%d)", size))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:228
		// _ = "end of CoverTab[91268]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:229
	// _ = "end of CoverTab[91261]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:229
	_go_fuzz_dep_.CoverTab[91262]++
												return b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:230
	// _ = "end of CoverTab[91262]"
}

// size returns the output size with currently set values.
func (h literalsHeader) size() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:234
	_go_fuzz_dep_.CoverTab[91269]++
												return int(h >> 60)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:235
	// _ = "end of CoverTab[91269]"
}

func (h literalsHeader) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:238
	_go_fuzz_dep_.CoverTab[91270]++
												return fmt.Sprintf("Type: %d, SizeFormat: %d, Size: 0x%d, Bytes:%d", literalsBlockType(h&3), (h>>2)&3, h&((1<<60)-1)>>4, h>>60)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:239
	// _ = "end of CoverTab[91270]"
}

// pushOffsets will push the recent offsets to the backup store.
func (b *blockEnc) pushOffsets() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:243
	_go_fuzz_dep_.CoverTab[91271]++
												b.prevRecentOffsets = b.recentOffsets
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:244
	// _ = "end of CoverTab[91271]"
}

// pushOffsets will push the recent offsets to the backup store.
func (b *blockEnc) popOffsets() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:248
	_go_fuzz_dep_.CoverTab[91272]++
												b.recentOffsets = b.prevRecentOffsets
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:249
	// _ = "end of CoverTab[91272]"
}

// matchOffset will adjust recent offsets and return the adjusted one,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:252
// if it matches a previous offset.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:254
func (b *blockEnc) matchOffset(offset, lits uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:254
	_go_fuzz_dep_.CoverTab[91273]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:258
	if true {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:258
		_go_fuzz_dep_.CoverTab[91275]++
													if lits > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:259
			_go_fuzz_dep_.CoverTab[91276]++
														switch offset {
			case b.recentOffsets[0]:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:261
				_go_fuzz_dep_.CoverTab[91277]++
															offset = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:262
				// _ = "end of CoverTab[91277]"
			case b.recentOffsets[1]:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:263
				_go_fuzz_dep_.CoverTab[91278]++
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset = 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:266
				// _ = "end of CoverTab[91278]"
			case b.recentOffsets[2]:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:267
				_go_fuzz_dep_.CoverTab[91279]++
															b.recentOffsets[2] = b.recentOffsets[1]
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset = 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:271
				// _ = "end of CoverTab[91279]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:272
				_go_fuzz_dep_.CoverTab[91280]++
															b.recentOffsets[2] = b.recentOffsets[1]
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset += 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:276
				// _ = "end of CoverTab[91280]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:277
			// _ = "end of CoverTab[91276]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:278
			_go_fuzz_dep_.CoverTab[91281]++
														switch offset {
			case b.recentOffsets[1]:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:280
				_go_fuzz_dep_.CoverTab[91282]++
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:283
				// _ = "end of CoverTab[91282]"
			case b.recentOffsets[2]:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:284
				_go_fuzz_dep_.CoverTab[91283]++
															b.recentOffsets[2] = b.recentOffsets[1]
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset = 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:288
				// _ = "end of CoverTab[91283]"
			case b.recentOffsets[0] - 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:289
				_go_fuzz_dep_.CoverTab[91284]++
															b.recentOffsets[2] = b.recentOffsets[1]
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset = 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:293
				// _ = "end of CoverTab[91284]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:294
				_go_fuzz_dep_.CoverTab[91285]++
															b.recentOffsets[2] = b.recentOffsets[1]
															b.recentOffsets[1] = b.recentOffsets[0]
															b.recentOffsets[0] = offset
															offset += 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:298
				// _ = "end of CoverTab[91285]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:299
			// _ = "end of CoverTab[91281]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:300
		// _ = "end of CoverTab[91275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:301
		_go_fuzz_dep_.CoverTab[91286]++
													offset += 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:302
		// _ = "end of CoverTab[91286]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:303
	// _ = "end of CoverTab[91273]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:303
	_go_fuzz_dep_.CoverTab[91274]++
												return offset
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:304
	// _ = "end of CoverTab[91274]"
}

// encodeRaw can be used to set the output to a raw representation of supplied bytes.
func (b *blockEnc) encodeRaw(a []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:308
	_go_fuzz_dep_.CoverTab[91287]++
												var bh blockHeader
												bh.setLast(b.last)
												bh.setSize(uint32(len(a)))
												bh.setType(blockTypeRaw)
												b.output = bh.appendTo(b.output[:0])
												b.output = append(b.output, a...)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:315
		_go_fuzz_dep_.CoverTab[91288]++
													println("Adding RAW block, length", len(a), "last:", b.last)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:316
		// _ = "end of CoverTab[91288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:317
		_go_fuzz_dep_.CoverTab[91289]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:317
		// _ = "end of CoverTab[91289]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:317
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:317
	// _ = "end of CoverTab[91287]"
}

// encodeRaw can be used to set the output to a raw representation of supplied bytes.
func (b *blockEnc) encodeRawTo(dst, src []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:321
	_go_fuzz_dep_.CoverTab[91290]++
												var bh blockHeader
												bh.setLast(b.last)
												bh.setSize(uint32(len(src)))
												bh.setType(blockTypeRaw)
												dst = bh.appendTo(dst)
												dst = append(dst, src...)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:328
		_go_fuzz_dep_.CoverTab[91292]++
													println("Adding RAW block, length", len(src), "last:", b.last)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:329
		// _ = "end of CoverTab[91292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:330
		_go_fuzz_dep_.CoverTab[91293]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:330
		// _ = "end of CoverTab[91293]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:330
	// _ = "end of CoverTab[91290]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:330
	_go_fuzz_dep_.CoverTab[91291]++
												return dst
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:331
	// _ = "end of CoverTab[91291]"
}

// encodeLits can be used if the block is only litLen.
func (b *blockEnc) encodeLits(lits []byte, raw bool) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:335
	_go_fuzz_dep_.CoverTab[91294]++
												var bh blockHeader
												bh.setLast(b.last)
												bh.setSize(uint32(len(lits)))

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
	if len(lits) < 8 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		_go_fuzz_dep_.CoverTab[91300]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		return (len(lits) < 32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
			_go_fuzz_dep_.CoverTab[91301]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
			return b.dictLitEnc == nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
			// _ = "end of CoverTab[91301]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		}())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		// _ = "end of CoverTab[91300]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		_go_fuzz_dep_.CoverTab[91302]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		return raw
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		// _ = "end of CoverTab[91302]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:341
		_go_fuzz_dep_.CoverTab[91303]++
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:342
			_go_fuzz_dep_.CoverTab[91305]++
														println("Adding RAW block, length", len(lits), "last:", b.last)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:343
			// _ = "end of CoverTab[91305]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:344
			_go_fuzz_dep_.CoverTab[91306]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:344
			// _ = "end of CoverTab[91306]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:344
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:344
		// _ = "end of CoverTab[91303]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:344
		_go_fuzz_dep_.CoverTab[91304]++
													bh.setType(blockTypeRaw)
													b.output = bh.appendTo(b.output)
													b.output = append(b.output, lits...)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:348
		// _ = "end of CoverTab[91304]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:349
		_go_fuzz_dep_.CoverTab[91307]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:349
		// _ = "end of CoverTab[91307]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:349
	// _ = "end of CoverTab[91294]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:349
	_go_fuzz_dep_.CoverTab[91295]++

												var (
		out		[]byte
		reUsed, single	bool
		err		error
	)
	if b.dictLitEnc != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:356
		_go_fuzz_dep_.CoverTab[91308]++
													b.litEnc.TransferCTable(b.dictLitEnc)
													b.litEnc.Reuse = huff0.ReusePolicyAllow
													b.dictLitEnc = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:359
		// _ = "end of CoverTab[91308]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:360
		_go_fuzz_dep_.CoverTab[91309]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:360
		// _ = "end of CoverTab[91309]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:360
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:360
	// _ = "end of CoverTab[91295]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:360
	_go_fuzz_dep_.CoverTab[91296]++
												if len(lits) >= 1024 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:361
		_go_fuzz_dep_.CoverTab[91310]++

													out, reUsed, err = huff0.Compress4X(lits, b.litEnc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:363
		// _ = "end of CoverTab[91310]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:364
		_go_fuzz_dep_.CoverTab[91311]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:364
		if len(lits) > 32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:364
			_go_fuzz_dep_.CoverTab[91312]++

														single = true
														out, reUsed, err = huff0.Compress1X(lits, b.litEnc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:367
			// _ = "end of CoverTab[91312]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:368
			_go_fuzz_dep_.CoverTab[91313]++
														err = huff0.ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:369
			// _ = "end of CoverTab[91313]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:370
		// _ = "end of CoverTab[91311]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:370
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:370
	// _ = "end of CoverTab[91296]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:370
	_go_fuzz_dep_.CoverTab[91297]++

												switch err {
	case huff0.ErrIncompressible:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:373
		_go_fuzz_dep_.CoverTab[91314]++
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:374
			_go_fuzz_dep_.CoverTab[91320]++
														println("Adding RAW block, length", len(lits), "last:", b.last)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:375
			// _ = "end of CoverTab[91320]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:376
			_go_fuzz_dep_.CoverTab[91321]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:376
			// _ = "end of CoverTab[91321]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:376
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:376
		// _ = "end of CoverTab[91314]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:376
		_go_fuzz_dep_.CoverTab[91315]++
													bh.setType(blockTypeRaw)
													b.output = bh.appendTo(b.output)
													b.output = append(b.output, lits...)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:380
		// _ = "end of CoverTab[91315]"
	case huff0.ErrUseRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:381
		_go_fuzz_dep_.CoverTab[91316]++
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:382
			_go_fuzz_dep_.CoverTab[91322]++
														println("Adding RLE block, length", len(lits))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:383
			// _ = "end of CoverTab[91322]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:384
			_go_fuzz_dep_.CoverTab[91323]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:384
			// _ = "end of CoverTab[91323]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:384
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:384
		// _ = "end of CoverTab[91316]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:384
		_go_fuzz_dep_.CoverTab[91317]++
													bh.setType(blockTypeRLE)
													b.output = bh.appendTo(b.output)
													b.output = append(b.output, lits[0])
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:388
		// _ = "end of CoverTab[91317]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:389
		_go_fuzz_dep_.CoverTab[91318]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:389
		// _ = "end of CoverTab[91318]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:390
		_go_fuzz_dep_.CoverTab[91319]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:391
		// _ = "end of CoverTab[91319]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:392
	// _ = "end of CoverTab[91297]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:392
	_go_fuzz_dep_.CoverTab[91298]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:395
	b.litEnc.Reuse = huff0.ReusePolicyAllow
	bh.setType(blockTypeCompressed)
	var lh literalsHeader
	if reUsed {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:398
		_go_fuzz_dep_.CoverTab[91324]++
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:399
			_go_fuzz_dep_.CoverTab[91326]++
														println("Reused tree, compressed to", len(out))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:400
			// _ = "end of CoverTab[91326]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:401
			_go_fuzz_dep_.CoverTab[91327]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:401
			// _ = "end of CoverTab[91327]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:401
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:401
		// _ = "end of CoverTab[91324]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:401
		_go_fuzz_dep_.CoverTab[91325]++
													lh.setType(literalsBlockTreeless)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:402
		// _ = "end of CoverTab[91325]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:403
		_go_fuzz_dep_.CoverTab[91328]++
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:404
			_go_fuzz_dep_.CoverTab[91330]++
														println("New tree, compressed to", len(out), "tree size:", len(b.litEnc.OutTable))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:405
			// _ = "end of CoverTab[91330]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:406
			_go_fuzz_dep_.CoverTab[91331]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:406
			// _ = "end of CoverTab[91331]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:406
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:406
		// _ = "end of CoverTab[91328]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:406
		_go_fuzz_dep_.CoverTab[91329]++
													lh.setType(literalsBlockCompressed)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:407
		// _ = "end of CoverTab[91329]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:408
	// _ = "end of CoverTab[91298]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:408
	_go_fuzz_dep_.CoverTab[91299]++

												lh.setSizes(len(out), len(lits), single)
												bh.setSize(uint32(len(out) + lh.size() + 1))

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:414
	b.output = bh.appendTo(b.output)
												b.output = lh.appendTo(b.output)

												b.output = append(b.output, out...)

												b.output = append(b.output, 0)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:420
	// _ = "end of CoverTab[91299]"
}

// fuzzFseEncoder can be used to fuzz the FSE encoder.
func fuzzFseEncoder(data []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:424
	_go_fuzz_dep_.CoverTab[91332]++
												if len(data) > maxSequences || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:425
		_go_fuzz_dep_.CoverTab[91340]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:425
		return len(data) < 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:425
		// _ = "end of CoverTab[91340]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:425
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:425
		_go_fuzz_dep_.CoverTab[91341]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:426
		// _ = "end of CoverTab[91341]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:427
		_go_fuzz_dep_.CoverTab[91342]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:427
		// _ = "end of CoverTab[91342]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:427
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:427
	// _ = "end of CoverTab[91332]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:427
	_go_fuzz_dep_.CoverTab[91333]++
												enc := fseEncoder{}
												hist := enc.Histogram()
												maxSym := uint8(0)
												for i, v := range data {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:431
		_go_fuzz_dep_.CoverTab[91343]++
													v = v & 63
													data[i] = v
													hist[v]++
													if v > maxSym {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:435
			_go_fuzz_dep_.CoverTab[91344]++
														maxSym = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:436
			// _ = "end of CoverTab[91344]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:437
			_go_fuzz_dep_.CoverTab[91345]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:437
			// _ = "end of CoverTab[91345]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:437
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:437
		// _ = "end of CoverTab[91343]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:438
	// _ = "end of CoverTab[91333]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:438
	_go_fuzz_dep_.CoverTab[91334]++
												if maxSym == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:439
		_go_fuzz_dep_.CoverTab[91346]++

													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:441
		// _ = "end of CoverTab[91346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:442
		_go_fuzz_dep_.CoverTab[91347]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:442
		// _ = "end of CoverTab[91347]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:442
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:442
	// _ = "end of CoverTab[91334]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:442
	_go_fuzz_dep_.CoverTab[91335]++
												maxCount := func(a []uint32) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:443
		_go_fuzz_dep_.CoverTab[91348]++
													var max uint32
													for _, v := range a {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:445
			_go_fuzz_dep_.CoverTab[91350]++
														if v > max {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:446
				_go_fuzz_dep_.CoverTab[91351]++
															max = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:447
				// _ = "end of CoverTab[91351]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:448
				_go_fuzz_dep_.CoverTab[91352]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:448
				// _ = "end of CoverTab[91352]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:448
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:448
			// _ = "end of CoverTab[91350]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:449
		// _ = "end of CoverTab[91348]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:449
		_go_fuzz_dep_.CoverTab[91349]++
													return int(max)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:450
		// _ = "end of CoverTab[91349]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:451
	// _ = "end of CoverTab[91335]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:451
	_go_fuzz_dep_.CoverTab[91336]++
												cnt := maxCount(hist[:maxSym])
												if cnt == len(data) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:453
		_go_fuzz_dep_.CoverTab[91353]++

													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:455
		// _ = "end of CoverTab[91353]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:456
		_go_fuzz_dep_.CoverTab[91354]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:456
		// _ = "end of CoverTab[91354]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:456
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:456
	// _ = "end of CoverTab[91336]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:456
	_go_fuzz_dep_.CoverTab[91337]++
												enc.HistogramFinished(maxSym, cnt)
												err := enc.normalizeCount(len(data))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:459
		_go_fuzz_dep_.CoverTab[91355]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:460
		// _ = "end of CoverTab[91355]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:461
		_go_fuzz_dep_.CoverTab[91356]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:461
		// _ = "end of CoverTab[91356]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:461
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:461
	// _ = "end of CoverTab[91337]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:461
	_go_fuzz_dep_.CoverTab[91338]++
												_, err = enc.writeCount(nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:463
		_go_fuzz_dep_.CoverTab[91357]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:464
		// _ = "end of CoverTab[91357]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:465
		_go_fuzz_dep_.CoverTab[91358]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:465
		// _ = "end of CoverTab[91358]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:465
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:465
	// _ = "end of CoverTab[91338]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:465
	_go_fuzz_dep_.CoverTab[91339]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:466
	// _ = "end of CoverTab[91339]"
}

// encode will encode the block and append the output in b.output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:469
// Previous offset codes must be pushed if more blocks are expected.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:471
func (b *blockEnc) encode(org []byte, raw, rawAllLits bool) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:471
	_go_fuzz_dep_.CoverTab[91359]++
												if len(b.sequences) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:472
		_go_fuzz_dep_.CoverTab[91384]++
													return b.encodeLits(b.literals, rawAllLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:473
		// _ = "end of CoverTab[91384]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:474
		_go_fuzz_dep_.CoverTab[91385]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:474
		// _ = "end of CoverTab[91385]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:474
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:474
	// _ = "end of CoverTab[91359]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:474
	_go_fuzz_dep_.CoverTab[91360]++

												saved := b.size - len(b.literals) - (b.size >> 5)
												if saved < 16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:477
		_go_fuzz_dep_.CoverTab[91386]++
													if org == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:478
			_go_fuzz_dep_.CoverTab[91388]++
														return errIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:479
			// _ = "end of CoverTab[91388]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:480
			_go_fuzz_dep_.CoverTab[91389]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:480
			// _ = "end of CoverTab[91389]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:480
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:480
		// _ = "end of CoverTab[91386]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:480
		_go_fuzz_dep_.CoverTab[91387]++
													b.popOffsets()
													return b.encodeLits(org, rawAllLits)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:482
		// _ = "end of CoverTab[91387]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:483
		_go_fuzz_dep_.CoverTab[91390]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:483
		// _ = "end of CoverTab[91390]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:483
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:483
	// _ = "end of CoverTab[91360]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:483
	_go_fuzz_dep_.CoverTab[91361]++

												var bh blockHeader
												var lh literalsHeader
												bh.setLast(b.last)
												bh.setType(blockTypeCompressed)

												bhOffset := len(b.output)
												b.output = bh.appendTo(b.output)

												var (
		out		[]byte
		reUsed, single	bool
		err		error
	)
	if b.dictLitEnc != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:498
		_go_fuzz_dep_.CoverTab[91391]++
													b.litEnc.TransferCTable(b.dictLitEnc)
													b.litEnc.Reuse = huff0.ReusePolicyAllow
													b.dictLitEnc = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:501
		// _ = "end of CoverTab[91391]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:502
		_go_fuzz_dep_.CoverTab[91392]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:502
		// _ = "end of CoverTab[91392]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:502
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:502
	// _ = "end of CoverTab[91361]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:502
	_go_fuzz_dep_.CoverTab[91362]++
												if len(b.literals) >= 1024 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:503
		_go_fuzz_dep_.CoverTab[91393]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:503
		return !raw
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:503
		// _ = "end of CoverTab[91393]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:503
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:503
		_go_fuzz_dep_.CoverTab[91394]++

													out, reUsed, err = huff0.Compress4X(b.literals, b.litEnc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:505
		// _ = "end of CoverTab[91394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
		_go_fuzz_dep_.CoverTab[91395]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
		if len(b.literals) > 32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
			_go_fuzz_dep_.CoverTab[91396]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
			return !raw
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
			// _ = "end of CoverTab[91396]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:506
			_go_fuzz_dep_.CoverTab[91397]++

														single = true
														out, reUsed, err = huff0.Compress1X(b.literals, b.litEnc)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:509
			// _ = "end of CoverTab[91397]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:510
			_go_fuzz_dep_.CoverTab[91398]++
														err = huff0.ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:511
			// _ = "end of CoverTab[91398]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:512
		// _ = "end of CoverTab[91395]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:512
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:512
	// _ = "end of CoverTab[91362]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:512
	_go_fuzz_dep_.CoverTab[91363]++

												switch err {
	case huff0.ErrIncompressible:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:515
		_go_fuzz_dep_.CoverTab[91399]++
													lh.setType(literalsBlockRaw)
													lh.setSize(len(b.literals))
													b.output = lh.appendTo(b.output)
													b.output = append(b.output, b.literals...)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:520
			_go_fuzz_dep_.CoverTab[91406]++
														println("Adding literals RAW, length", len(b.literals))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:521
			// _ = "end of CoverTab[91406]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:522
			_go_fuzz_dep_.CoverTab[91407]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:522
			// _ = "end of CoverTab[91407]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:522
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:522
		// _ = "end of CoverTab[91399]"
	case huff0.ErrUseRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:523
		_go_fuzz_dep_.CoverTab[91400]++
													lh.setType(literalsBlockRLE)
													lh.setSize(len(b.literals))
													b.output = lh.appendTo(b.output)
													b.output = append(b.output, b.literals[0])
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:528
			_go_fuzz_dep_.CoverTab[91408]++
														println("Adding literals RLE")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:529
			// _ = "end of CoverTab[91408]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:530
			_go_fuzz_dep_.CoverTab[91409]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:530
			// _ = "end of CoverTab[91409]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:530
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:530
		// _ = "end of CoverTab[91400]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:531
		_go_fuzz_dep_.CoverTab[91401]++

													if reUsed {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:533
			_go_fuzz_dep_.CoverTab[91410]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:534
				_go_fuzz_dep_.CoverTab[91412]++
															println("reused tree")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:535
				// _ = "end of CoverTab[91412]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:536
				_go_fuzz_dep_.CoverTab[91413]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:536
				// _ = "end of CoverTab[91413]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:536
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:536
			// _ = "end of CoverTab[91410]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:536
			_go_fuzz_dep_.CoverTab[91411]++
														lh.setType(literalsBlockTreeless)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:537
			// _ = "end of CoverTab[91411]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:538
			_go_fuzz_dep_.CoverTab[91414]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:539
				_go_fuzz_dep_.CoverTab[91416]++
															println("new tree, size:", len(b.litEnc.OutTable))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:540
				// _ = "end of CoverTab[91416]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:541
				_go_fuzz_dep_.CoverTab[91417]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:541
				// _ = "end of CoverTab[91417]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:541
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:541
			// _ = "end of CoverTab[91414]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:541
			_go_fuzz_dep_.CoverTab[91415]++
														lh.setType(literalsBlockCompressed)
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:543
				_go_fuzz_dep_.CoverTab[91418]++
															_, _, err := huff0.ReadTable(out, nil)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:545
					_go_fuzz_dep_.CoverTab[91419]++
																panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:546
					// _ = "end of CoverTab[91419]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:547
					_go_fuzz_dep_.CoverTab[91420]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:547
					// _ = "end of CoverTab[91420]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:547
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:547
				// _ = "end of CoverTab[91418]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:548
				_go_fuzz_dep_.CoverTab[91421]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:548
				// _ = "end of CoverTab[91421]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:548
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:548
			// _ = "end of CoverTab[91415]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:549
		// _ = "end of CoverTab[91401]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:549
		_go_fuzz_dep_.CoverTab[91402]++
													lh.setSizes(len(out), len(b.literals), single)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:551
			_go_fuzz_dep_.CoverTab[91422]++
														printf("Compressed %d literals to %d bytes", len(b.literals), len(out))
														println("Adding literal header:", lh)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:553
			// _ = "end of CoverTab[91422]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:554
			_go_fuzz_dep_.CoverTab[91423]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:554
			// _ = "end of CoverTab[91423]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:554
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:554
		// _ = "end of CoverTab[91402]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:554
		_go_fuzz_dep_.CoverTab[91403]++
													b.output = lh.appendTo(b.output)
													b.output = append(b.output, out...)
													b.litEnc.Reuse = huff0.ReusePolicyAllow
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:558
			_go_fuzz_dep_.CoverTab[91424]++
														println("Adding literals compressed")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:559
			// _ = "end of CoverTab[91424]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:560
			_go_fuzz_dep_.CoverTab[91425]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:560
			// _ = "end of CoverTab[91425]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:560
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:560
		// _ = "end of CoverTab[91403]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:561
		_go_fuzz_dep_.CoverTab[91404]++
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:562
			_go_fuzz_dep_.CoverTab[91426]++
														println("Adding literals ERROR:", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:563
			// _ = "end of CoverTab[91426]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:564
			_go_fuzz_dep_.CoverTab[91427]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:564
			// _ = "end of CoverTab[91427]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:564
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:564
		// _ = "end of CoverTab[91404]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:564
		_go_fuzz_dep_.CoverTab[91405]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:565
		// _ = "end of CoverTab[91405]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:566
	// _ = "end of CoverTab[91363]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:566
	_go_fuzz_dep_.CoverTab[91364]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:570
	switch {
	case len(b.sequences) < 128:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:571
		_go_fuzz_dep_.CoverTab[91428]++
													b.output = append(b.output, uint8(len(b.sequences)))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:572
		// _ = "end of CoverTab[91428]"
	case len(b.sequences) < 0x7f00:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:573
		_go_fuzz_dep_.CoverTab[91429]++
													n := len(b.sequences)
													b.output = append(b.output, 128+uint8(n>>8), uint8(n))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:575
		// _ = "end of CoverTab[91429]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:576
		_go_fuzz_dep_.CoverTab[91430]++
													n := len(b.sequences) - 0x7f00
													b.output = append(b.output, 255, uint8(n), uint8(n>>8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:578
		// _ = "end of CoverTab[91430]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:579
	// _ = "end of CoverTab[91364]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:579
	_go_fuzz_dep_.CoverTab[91365]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:580
		_go_fuzz_dep_.CoverTab[91431]++
													println("Encoding", len(b.sequences), "sequences")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:581
		// _ = "end of CoverTab[91431]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:582
		_go_fuzz_dep_.CoverTab[91432]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:582
		// _ = "end of CoverTab[91432]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:582
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:582
	// _ = "end of CoverTab[91365]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:582
	_go_fuzz_dep_.CoverTab[91366]++
												b.genCodes()
												llEnc := b.coders.llEnc
												ofEnc := b.coders.ofEnc
												mlEnc := b.coders.mlEnc
												err = llEnc.normalizeCount(len(b.sequences))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:588
		_go_fuzz_dep_.CoverTab[91433]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:589
		// _ = "end of CoverTab[91433]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:590
		_go_fuzz_dep_.CoverTab[91434]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:590
		// _ = "end of CoverTab[91434]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:590
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:590
	// _ = "end of CoverTab[91366]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:590
	_go_fuzz_dep_.CoverTab[91367]++
												err = ofEnc.normalizeCount(len(b.sequences))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:592
		_go_fuzz_dep_.CoverTab[91435]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:593
		// _ = "end of CoverTab[91435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:594
		_go_fuzz_dep_.CoverTab[91436]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:594
		// _ = "end of CoverTab[91436]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:594
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:594
	// _ = "end of CoverTab[91367]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:594
	_go_fuzz_dep_.CoverTab[91368]++
												err = mlEnc.normalizeCount(len(b.sequences))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:596
		_go_fuzz_dep_.CoverTab[91437]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:597
		// _ = "end of CoverTab[91437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:598
		_go_fuzz_dep_.CoverTab[91438]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:598
		// _ = "end of CoverTab[91438]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:598
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:598
	// _ = "end of CoverTab[91368]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:598
	_go_fuzz_dep_.CoverTab[91369]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:602
	chooseComp := func(cur, prev, preDef *fseEncoder) (*fseEncoder, seqCompMode) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:602
		_go_fuzz_dep_.CoverTab[91439]++

													hist := cur.count[:cur.symbolLen]
													nSize := cur.approxSize(hist) + cur.maxHeaderSize()
													predefSize := preDef.approxSize(hist)
													prevSize := prev.approxSize(hist)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:611
		nSize = nSize + (nSize+2*8*16)>>4
		switch {
		case predefSize <= prevSize && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			_go_fuzz_dep_.CoverTab[91446]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			return predefSize <= nSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			// _ = "end of CoverTab[91446]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			_go_fuzz_dep_.CoverTab[91447]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			return forcePreDef
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			// _ = "end of CoverTab[91447]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
		}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:613
			_go_fuzz_dep_.CoverTab[91440]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:614
				_go_fuzz_dep_.CoverTab[91448]++
															println("Using predefined", predefSize>>3, "<=", nSize>>3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:615
				// _ = "end of CoverTab[91448]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:616
				_go_fuzz_dep_.CoverTab[91449]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:616
				// _ = "end of CoverTab[91449]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:616
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:616
			// _ = "end of CoverTab[91440]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:616
			_go_fuzz_dep_.CoverTab[91441]++
														return preDef, compModePredefined
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:617
			// _ = "end of CoverTab[91441]"
		case prevSize <= nSize:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:618
			_go_fuzz_dep_.CoverTab[91442]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:619
				_go_fuzz_dep_.CoverTab[91450]++
															println("Using previous", prevSize>>3, "<=", nSize>>3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:620
				// _ = "end of CoverTab[91450]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:621
				_go_fuzz_dep_.CoverTab[91451]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:621
				// _ = "end of CoverTab[91451]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:621
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:621
			// _ = "end of CoverTab[91442]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:621
			_go_fuzz_dep_.CoverTab[91443]++
														return prev, compModeRepeat
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:622
			// _ = "end of CoverTab[91443]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:623
			_go_fuzz_dep_.CoverTab[91444]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:624
				_go_fuzz_dep_.CoverTab[91452]++
															println("Using new, predef", predefSize>>3, ". previous:", prevSize>>3, ">", nSize>>3, "header max:", cur.maxHeaderSize()>>3, "bytes")
															println("tl:", cur.actualTableLog, "symbolLen:", cur.symbolLen, "norm:", cur.norm[:cur.symbolLen], "hist", cur.count[:cur.symbolLen])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:626
				// _ = "end of CoverTab[91452]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:627
				_go_fuzz_dep_.CoverTab[91453]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:627
				// _ = "end of CoverTab[91453]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:627
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:627
			// _ = "end of CoverTab[91444]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:627
			_go_fuzz_dep_.CoverTab[91445]++
														return cur, compModeFSE
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:628
			// _ = "end of CoverTab[91445]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:629
		// _ = "end of CoverTab[91439]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:630
	// _ = "end of CoverTab[91369]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:630
	_go_fuzz_dep_.CoverTab[91370]++

	// Write compression mode
	var mode uint8
	if llEnc.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:634
		_go_fuzz_dep_.CoverTab[91454]++
													mode |= uint8(compModeRLE) << 6
													llEnc.setRLE(b.sequences[0].llCode)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:637
			_go_fuzz_dep_.CoverTab[91455]++
														println("llEnc.useRLE")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:638
			// _ = "end of CoverTab[91455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:639
			_go_fuzz_dep_.CoverTab[91456]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:639
			// _ = "end of CoverTab[91456]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:639
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:639
		// _ = "end of CoverTab[91454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:640
		_go_fuzz_dep_.CoverTab[91457]++
													var m seqCompMode
													llEnc, m = chooseComp(llEnc, b.coders.llPrev, &fsePredefEnc[tableLiteralLengths])
													mode |= uint8(m) << 6
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:643
		// _ = "end of CoverTab[91457]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:644
	// _ = "end of CoverTab[91370]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:644
	_go_fuzz_dep_.CoverTab[91371]++
												if ofEnc.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:645
		_go_fuzz_dep_.CoverTab[91458]++
													mode |= uint8(compModeRLE) << 4
													ofEnc.setRLE(b.sequences[0].ofCode)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:648
			_go_fuzz_dep_.CoverTab[91459]++
														println("ofEnc.useRLE")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:649
			// _ = "end of CoverTab[91459]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:650
			_go_fuzz_dep_.CoverTab[91460]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:650
			// _ = "end of CoverTab[91460]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:650
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:650
		// _ = "end of CoverTab[91458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:651
		_go_fuzz_dep_.CoverTab[91461]++
													var m seqCompMode
													ofEnc, m = chooseComp(ofEnc, b.coders.ofPrev, &fsePredefEnc[tableOffsets])
													mode |= uint8(m) << 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:654
		// _ = "end of CoverTab[91461]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:655
	// _ = "end of CoverTab[91371]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:655
	_go_fuzz_dep_.CoverTab[91372]++

												if mlEnc.useRLE {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:657
		_go_fuzz_dep_.CoverTab[91462]++
													mode |= uint8(compModeRLE) << 2
													mlEnc.setRLE(b.sequences[0].mlCode)
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:660
			_go_fuzz_dep_.CoverTab[91463]++
														println("mlEnc.useRLE, code: ", b.sequences[0].mlCode, "value", b.sequences[0].matchLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:661
			// _ = "end of CoverTab[91463]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:662
			_go_fuzz_dep_.CoverTab[91464]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:662
			// _ = "end of CoverTab[91464]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:662
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:662
		// _ = "end of CoverTab[91462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:663
		_go_fuzz_dep_.CoverTab[91465]++
													var m seqCompMode
													mlEnc, m = chooseComp(mlEnc, b.coders.mlPrev, &fsePredefEnc[tableMatchLengths])
													mode |= uint8(m) << 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:666
		// _ = "end of CoverTab[91465]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:667
	// _ = "end of CoverTab[91372]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:667
	_go_fuzz_dep_.CoverTab[91373]++
												b.output = append(b.output, mode)
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:669
		_go_fuzz_dep_.CoverTab[91466]++
													printf("Compression modes: 0b%b", mode)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:670
		// _ = "end of CoverTab[91466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:671
		_go_fuzz_dep_.CoverTab[91467]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:671
		// _ = "end of CoverTab[91467]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:671
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:671
	// _ = "end of CoverTab[91373]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:671
	_go_fuzz_dep_.CoverTab[91374]++
												b.output, err = llEnc.writeCount(b.output)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:673
		_go_fuzz_dep_.CoverTab[91468]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:674
		// _ = "end of CoverTab[91468]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:675
		_go_fuzz_dep_.CoverTab[91469]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:675
		// _ = "end of CoverTab[91469]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:675
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:675
	// _ = "end of CoverTab[91374]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:675
	_go_fuzz_dep_.CoverTab[91375]++
												start := len(b.output)
												b.output, err = ofEnc.writeCount(b.output)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:678
		_go_fuzz_dep_.CoverTab[91470]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:679
		// _ = "end of CoverTab[91470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:680
		_go_fuzz_dep_.CoverTab[91471]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:680
		// _ = "end of CoverTab[91471]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:680
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:680
	// _ = "end of CoverTab[91375]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:680
	_go_fuzz_dep_.CoverTab[91376]++
												if false {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:681
		_go_fuzz_dep_.CoverTab[91472]++
													println("block:", b.output[start:], "tablelog", ofEnc.actualTableLog, "maxcount:", ofEnc.maxCount)
													fmt.Printf("selected TableLog: %d, Symbol length: %d\n", ofEnc.actualTableLog, ofEnc.symbolLen)
													for i, v := range ofEnc.norm[:ofEnc.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:684
			_go_fuzz_dep_.CoverTab[91473]++
														fmt.Printf("%3d: %5d -> %4d \n", i, ofEnc.count[i], v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:685
			// _ = "end of CoverTab[91473]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:686
		// _ = "end of CoverTab[91472]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:687
		_go_fuzz_dep_.CoverTab[91474]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:687
		// _ = "end of CoverTab[91474]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:687
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:687
	// _ = "end of CoverTab[91376]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:687
	_go_fuzz_dep_.CoverTab[91377]++
												b.output, err = mlEnc.writeCount(b.output)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:689
		_go_fuzz_dep_.CoverTab[91475]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:690
		// _ = "end of CoverTab[91475]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:691
		_go_fuzz_dep_.CoverTab[91476]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:691
		// _ = "end of CoverTab[91476]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:691
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:691
	// _ = "end of CoverTab[91377]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:691
	_go_fuzz_dep_.CoverTab[91378]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:694
	wr := &b.wr
												wr.reset(b.output)

												var ll, of, ml cState

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:700
	seq := len(b.sequences) - 1
												s := b.sequences[seq]
												llEnc.setBits(llBitsTable[:])
												mlEnc.setBits(mlBitsTable[:])
												ofEnc.setBits(nil)

												llTT, ofTT, mlTT := llEnc.ct.symbolTT[:256], ofEnc.ct.symbolTT[:256], mlEnc.ct.symbolTT[:256]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:710
	llB, ofB, mlB := llTT[s.llCode], ofTT[s.ofCode], mlTT[s.mlCode]
												ll.init(wr, &llEnc.ct, llB)
												of.init(wr, &ofEnc.ct, ofB)
												wr.flush32()
												ml.init(wr, &mlEnc.ct, mlB)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:717
	wr.addBits32NC(s.litLen, llB.outBits)
	wr.addBits32NC(s.matchLen, mlB.outBits)
	wr.flush32()
	wr.addBits32NC(s.offset, ofB.outBits)
	if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:721
		_go_fuzz_dep_.CoverTab[91477]++
													println("Encoded seq", seq, s, "codes:", s.llCode, s.mlCode, s.ofCode, "states:", ll.state, ml.state, of.state, "bits:", llB, mlB, ofB)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:722
		// _ = "end of CoverTab[91477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:723
		_go_fuzz_dep_.CoverTab[91478]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:723
		// _ = "end of CoverTab[91478]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:723
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:723
	// _ = "end of CoverTab[91378]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:723
	_go_fuzz_dep_.CoverTab[91379]++
												seq--

												for seq >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:726
		_go_fuzz_dep_.CoverTab[91479]++
													s = b.sequences[seq]

													ofB := ofTT[s.ofCode]
													wr.flush32()

													nbBitsOut := (uint32(of.state) + ofB.deltaNbBits) >> 16
													dstState := int32(of.state>>(nbBitsOut&15)) + int32(ofB.deltaFindState)
													wr.addBits16NC(of.state, uint8(nbBitsOut))
													of.state = of.stateTable[dstState]

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:738
		outBits := ofB.outBits & 31
		extraBits := uint64(s.offset & bitMask32[outBits])
		extraBitsN := outBits

		mlB := mlTT[s.mlCode]

		nbBitsOut = (uint32(ml.state) + mlB.deltaNbBits) >> 16
		dstState = int32(ml.state>>(nbBitsOut&15)) + int32(mlB.deltaFindState)
		wr.addBits16NC(ml.state, uint8(nbBitsOut))
		ml.state = ml.stateTable[dstState]

		outBits = mlB.outBits & 31
		extraBits = extraBits<<outBits | uint64(s.matchLen&bitMask32[outBits])
		extraBitsN += outBits

		llB := llTT[s.llCode]

		nbBitsOut = (uint32(ll.state) + llB.deltaNbBits) >> 16
		dstState = int32(ll.state>>(nbBitsOut&15)) + int32(llB.deltaFindState)
		wr.addBits16NC(ll.state, uint8(nbBitsOut))
		ll.state = ll.stateTable[dstState]

		outBits = llB.outBits & 31
		extraBits = extraBits<<outBits | uint64(s.litLen&bitMask32[outBits])
		extraBitsN += outBits

		wr.flush32()
		wr.addBits64NC(extraBits, extraBitsN)

		if debugSequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:767
			_go_fuzz_dep_.CoverTab[91481]++
														println("Encoded seq", seq, s)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:768
			// _ = "end of CoverTab[91481]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:769
			_go_fuzz_dep_.CoverTab[91482]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:769
			// _ = "end of CoverTab[91482]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:769
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:769
		// _ = "end of CoverTab[91479]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:769
		_go_fuzz_dep_.CoverTab[91480]++

													seq--
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:771
		// _ = "end of CoverTab[91480]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:772
	// _ = "end of CoverTab[91379]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:772
	_go_fuzz_dep_.CoverTab[91380]++
												ml.flush(mlEnc.actualTableLog)
												of.flush(ofEnc.actualTableLog)
												ll.flush(llEnc.actualTableLog)
												err = wr.close()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:777
		_go_fuzz_dep_.CoverTab[91483]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:778
		// _ = "end of CoverTab[91483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:779
		_go_fuzz_dep_.CoverTab[91484]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:779
		// _ = "end of CoverTab[91484]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:779
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:779
	// _ = "end of CoverTab[91380]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:779
	_go_fuzz_dep_.CoverTab[91381]++
												b.output = wr.out

												if len(b.output)-3-bhOffset >= b.size {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:782
		_go_fuzz_dep_.CoverTab[91485]++

													b.litEnc.Reuse = huff0.ReusePolicyNone
													return errIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:785
		// _ = "end of CoverTab[91485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:786
		_go_fuzz_dep_.CoverTab[91486]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:786
		// _ = "end of CoverTab[91486]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:786
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:786
	// _ = "end of CoverTab[91381]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:786
	_go_fuzz_dep_.CoverTab[91382]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:789
	bh.setSize(uint32(len(b.output)-bhOffset) - 3)
	if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:790
		_go_fuzz_dep_.CoverTab[91487]++
													println("Rewriting block header", bh)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:791
		// _ = "end of CoverTab[91487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:792
		_go_fuzz_dep_.CoverTab[91488]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:792
		// _ = "end of CoverTab[91488]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:792
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:792
	// _ = "end of CoverTab[91382]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:792
	_go_fuzz_dep_.CoverTab[91383]++
												_ = bh.appendTo(b.output[bhOffset:bhOffset])
												b.coders.setPrev(llEnc, mlEnc, ofEnc)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:795
	// _ = "end of CoverTab[91383]"
}

var errIncompressible = errors.New("incompressible")

func (b *blockEnc) genCodes() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:800
	_go_fuzz_dep_.CoverTab[91489]++
												if len(b.sequences) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:801
		_go_fuzz_dep_.CoverTab[91500]++

													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:803
		// _ = "end of CoverTab[91500]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:804
		_go_fuzz_dep_.CoverTab[91501]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:804
		// _ = "end of CoverTab[91501]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:804
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:804
	// _ = "end of CoverTab[91489]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:804
	_go_fuzz_dep_.CoverTab[91490]++
												if len(b.sequences) > math.MaxUint16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:805
		_go_fuzz_dep_.CoverTab[91502]++
													panic("can only encode up to 64K sequences")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:806
		// _ = "end of CoverTab[91502]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:807
		_go_fuzz_dep_.CoverTab[91503]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:807
		// _ = "end of CoverTab[91503]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:807
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:807
	// _ = "end of CoverTab[91490]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:807
	_go_fuzz_dep_.CoverTab[91491]++

												llH := b.coders.llEnc.Histogram()
												ofH := b.coders.ofEnc.Histogram()
												mlH := b.coders.mlEnc.Histogram()
												for i := range llH {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:812
		_go_fuzz_dep_.CoverTab[91504]++
													llH[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:813
		// _ = "end of CoverTab[91504]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:814
	// _ = "end of CoverTab[91491]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:814
	_go_fuzz_dep_.CoverTab[91492]++
												for i := range ofH {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:815
		_go_fuzz_dep_.CoverTab[91505]++
													ofH[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:816
		// _ = "end of CoverTab[91505]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:817
	// _ = "end of CoverTab[91492]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:817
	_go_fuzz_dep_.CoverTab[91493]++
												for i := range mlH {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:818
		_go_fuzz_dep_.CoverTab[91506]++
													mlH[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:819
		// _ = "end of CoverTab[91506]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:820
	// _ = "end of CoverTab[91493]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:820
	_go_fuzz_dep_.CoverTab[91494]++

												var llMax, ofMax, mlMax uint8
												for i := range b.sequences {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:823
		_go_fuzz_dep_.CoverTab[91507]++
													seq := &b.sequences[i]
													v := llCode(seq.litLen)
													seq.llCode = v
													llH[v]++
													if v > llMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:828
			_go_fuzz_dep_.CoverTab[91510]++
														llMax = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:829
			// _ = "end of CoverTab[91510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:830
			_go_fuzz_dep_.CoverTab[91511]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:830
			// _ = "end of CoverTab[91511]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:830
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:830
		// _ = "end of CoverTab[91507]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:830
		_go_fuzz_dep_.CoverTab[91508]++

													v = ofCode(seq.offset)
													seq.ofCode = v
													ofH[v]++
													if v > ofMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:835
			_go_fuzz_dep_.CoverTab[91512]++
														ofMax = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:836
			// _ = "end of CoverTab[91512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:837
			_go_fuzz_dep_.CoverTab[91513]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:837
			// _ = "end of CoverTab[91513]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:837
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:837
		// _ = "end of CoverTab[91508]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:837
		_go_fuzz_dep_.CoverTab[91509]++

													v = mlCode(seq.matchLen)
													seq.mlCode = v
													mlH[v]++
													if v > mlMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:842
			_go_fuzz_dep_.CoverTab[91514]++
														mlMax = v
														if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:844
				_go_fuzz_dep_.CoverTab[91515]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:844
				return mlMax > maxMatchLengthSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:844
				// _ = "end of CoverTab[91515]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:844
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:844
				_go_fuzz_dep_.CoverTab[91516]++
															panic(fmt.Errorf("mlMax > maxMatchLengthSymbol (%d), matchlen: %d", mlMax, seq.matchLen))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:845
				// _ = "end of CoverTab[91516]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:846
				_go_fuzz_dep_.CoverTab[91517]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:846
				// _ = "end of CoverTab[91517]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:846
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:846
			// _ = "end of CoverTab[91514]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:847
			_go_fuzz_dep_.CoverTab[91518]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:847
			// _ = "end of CoverTab[91518]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:847
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:847
		// _ = "end of CoverTab[91509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:848
	// _ = "end of CoverTab[91494]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:848
	_go_fuzz_dep_.CoverTab[91495]++
												maxCount := func(a []uint32) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:849
		_go_fuzz_dep_.CoverTab[91519]++
													var max uint32
													for _, v := range a {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:851
			_go_fuzz_dep_.CoverTab[91521]++
														if v > max {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:852
				_go_fuzz_dep_.CoverTab[91522]++
															max = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:853
				// _ = "end of CoverTab[91522]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:854
				_go_fuzz_dep_.CoverTab[91523]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:854
				// _ = "end of CoverTab[91523]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:854
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:854
			// _ = "end of CoverTab[91521]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:855
		// _ = "end of CoverTab[91519]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:855
		_go_fuzz_dep_.CoverTab[91520]++
													return int(max)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:856
		// _ = "end of CoverTab[91520]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:857
	// _ = "end of CoverTab[91495]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:857
	_go_fuzz_dep_.CoverTab[91496]++
												if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:858
		_go_fuzz_dep_.CoverTab[91524]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:858
		return mlMax > maxMatchLengthSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:858
		// _ = "end of CoverTab[91524]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:858
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:858
		_go_fuzz_dep_.CoverTab[91525]++
													panic(fmt.Errorf("mlMax > maxMatchLengthSymbol (%d)", mlMax))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:859
		// _ = "end of CoverTab[91525]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:860
		_go_fuzz_dep_.CoverTab[91526]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:860
		// _ = "end of CoverTab[91526]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:860
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:860
	// _ = "end of CoverTab[91496]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:860
	_go_fuzz_dep_.CoverTab[91497]++
												if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:861
		_go_fuzz_dep_.CoverTab[91527]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:861
		return ofMax > maxOffsetBits
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:861
		// _ = "end of CoverTab[91527]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:861
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:861
		_go_fuzz_dep_.CoverTab[91528]++
													panic(fmt.Errorf("ofMax > maxOffsetBits (%d)", ofMax))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:862
		// _ = "end of CoverTab[91528]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:863
		_go_fuzz_dep_.CoverTab[91529]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:863
		// _ = "end of CoverTab[91529]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:863
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:863
	// _ = "end of CoverTab[91497]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:863
	_go_fuzz_dep_.CoverTab[91498]++
												if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:864
		_go_fuzz_dep_.CoverTab[91530]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:864
		return llMax > maxLiteralLengthSymbol
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:864
		// _ = "end of CoverTab[91530]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:864
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:864
		_go_fuzz_dep_.CoverTab[91531]++
													panic(fmt.Errorf("llMax > maxLiteralLengthSymbol (%d)", llMax))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:865
		// _ = "end of CoverTab[91531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:866
		_go_fuzz_dep_.CoverTab[91532]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:866
		// _ = "end of CoverTab[91532]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:866
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:866
	// _ = "end of CoverTab[91498]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:866
	_go_fuzz_dep_.CoverTab[91499]++

												b.coders.mlEnc.HistogramFinished(mlMax, maxCount(mlH[:mlMax+1]))
												b.coders.ofEnc.HistogramFinished(ofMax, maxCount(ofH[:ofMax+1]))
												b.coders.llEnc.HistogramFinished(llMax, maxCount(llH[:llMax+1]))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:870
	// _ = "end of CoverTab[91499]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:871
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockenc.go:871
var _ = _go_fuzz_dep_.CoverTab
