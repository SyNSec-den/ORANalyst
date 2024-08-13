// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/huffman_code.go:5
package flate

//line /usr/local/go/src/compress/flate/huffman_code.go:5
import (
//line /usr/local/go/src/compress/flate/huffman_code.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/huffman_code.go:5
)
//line /usr/local/go/src/compress/flate/huffman_code.go:5
import (
//line /usr/local/go/src/compress/flate/huffman_code.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/huffman_code.go:5
)

import (
	"math"
	"math/bits"
	"sort"
)

// hcode is a huffman code with a bit code and bit length.
type hcode struct {
	code, len uint16
}

type huffmanEncoder struct {
	codes		[]hcode
	freqcache	[]literalNode
	bitCount	[17]int32
	lns		byLiteral	// stored to avoid repeated allocation in generate
	lfs		byFreq		// stored to avoid repeated allocation in generate
}

type literalNode struct {
	literal	uint16
	freq	int32
}

// A levelInfo describes the state of the constructed tree for a given depth.
type levelInfo struct {
	// Our level.  for better printing
	level	int32

	// The frequency of the last node at this level
	lastFreq	int32

	// The frequency of the next character to add to this level
	nextCharFreq	int32

	// The frequency of the next pair (from level below) to add to this level.
	// Only valid if the "needed" value of the next lower level is 0.
	nextPairFreq	int32

	// The number of chains remaining to generate for this level before moving
	// up to the next level
	needed	int32
}

// set sets the code and length of an hcode.
func (h *hcode) set(code uint16, length uint16) {
//line /usr/local/go/src/compress/flate/huffman_code.go:52
	_go_fuzz_dep_.CoverTab[26229]++
								h.len = length
								h.code = code
//line /usr/local/go/src/compress/flate/huffman_code.go:54
	// _ = "end of CoverTab[26229]"
}

func maxNode() literalNode {
//line /usr/local/go/src/compress/flate/huffman_code.go:57
	_go_fuzz_dep_.CoverTab[26230]++
//line /usr/local/go/src/compress/flate/huffman_code.go:57
	return literalNode{math.MaxUint16, math.MaxInt32}
//line /usr/local/go/src/compress/flate/huffman_code.go:57
	// _ = "end of CoverTab[26230]"
//line /usr/local/go/src/compress/flate/huffman_code.go:57
}

func newHuffmanEncoder(size int) *huffmanEncoder {
//line /usr/local/go/src/compress/flate/huffman_code.go:59
	_go_fuzz_dep_.CoverTab[26231]++
								return &huffmanEncoder{codes: make([]hcode, size)}
//line /usr/local/go/src/compress/flate/huffman_code.go:60
	// _ = "end of CoverTab[26231]"
}

// Generates a HuffmanCode corresponding to the fixed literal table.
func generateFixedLiteralEncoding() *huffmanEncoder {
//line /usr/local/go/src/compress/flate/huffman_code.go:64
	_go_fuzz_dep_.CoverTab[26232]++
								h := newHuffmanEncoder(maxNumLit)
								codes := h.codes
								var ch uint16
								for ch = 0; ch < maxNumLit; ch++ {
//line /usr/local/go/src/compress/flate/huffman_code.go:68
		_go_fuzz_dep_.CoverTab[26234]++
									var bits uint16
									var size uint16
									switch {
		case ch < 144:
//line /usr/local/go/src/compress/flate/huffman_code.go:72
			_go_fuzz_dep_.CoverTab[26236]++

										bits = ch + 48
										size = 8
//line /usr/local/go/src/compress/flate/huffman_code.go:75
			// _ = "end of CoverTab[26236]"
		case ch < 256:
//line /usr/local/go/src/compress/flate/huffman_code.go:76
			_go_fuzz_dep_.CoverTab[26237]++

										bits = ch + 400 - 144
										size = 9
//line /usr/local/go/src/compress/flate/huffman_code.go:79
			// _ = "end of CoverTab[26237]"
		case ch < 280:
//line /usr/local/go/src/compress/flate/huffman_code.go:80
			_go_fuzz_dep_.CoverTab[26238]++

										bits = ch - 256
										size = 7
//line /usr/local/go/src/compress/flate/huffman_code.go:83
			// _ = "end of CoverTab[26238]"
		default:
//line /usr/local/go/src/compress/flate/huffman_code.go:84
			_go_fuzz_dep_.CoverTab[26239]++

										bits = ch + 192 - 280
										size = 8
//line /usr/local/go/src/compress/flate/huffman_code.go:87
			// _ = "end of CoverTab[26239]"
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:88
		// _ = "end of CoverTab[26234]"
//line /usr/local/go/src/compress/flate/huffman_code.go:88
		_go_fuzz_dep_.CoverTab[26235]++
									codes[ch] = hcode{code: reverseBits(bits, byte(size)), len: size}
//line /usr/local/go/src/compress/flate/huffman_code.go:89
		// _ = "end of CoverTab[26235]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:90
	// _ = "end of CoverTab[26232]"
//line /usr/local/go/src/compress/flate/huffman_code.go:90
	_go_fuzz_dep_.CoverTab[26233]++
								return h
//line /usr/local/go/src/compress/flate/huffman_code.go:91
	// _ = "end of CoverTab[26233]"
}

func generateFixedOffsetEncoding() *huffmanEncoder {
//line /usr/local/go/src/compress/flate/huffman_code.go:94
	_go_fuzz_dep_.CoverTab[26240]++
								h := newHuffmanEncoder(30)
								codes := h.codes
								for ch := range codes {
//line /usr/local/go/src/compress/flate/huffman_code.go:97
		_go_fuzz_dep_.CoverTab[26242]++
									codes[ch] = hcode{code: reverseBits(uint16(ch), 5), len: 5}
//line /usr/local/go/src/compress/flate/huffman_code.go:98
		// _ = "end of CoverTab[26242]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:99
	// _ = "end of CoverTab[26240]"
//line /usr/local/go/src/compress/flate/huffman_code.go:99
	_go_fuzz_dep_.CoverTab[26241]++
								return h
//line /usr/local/go/src/compress/flate/huffman_code.go:100
	// _ = "end of CoverTab[26241]"
}

var fixedLiteralEncoding *huffmanEncoder = generateFixedLiteralEncoding()
var fixedOffsetEncoding *huffmanEncoder = generateFixedOffsetEncoding()

func (h *huffmanEncoder) bitLength(freq []int32) int {
//line /usr/local/go/src/compress/flate/huffman_code.go:106
	_go_fuzz_dep_.CoverTab[26243]++
								var total int
								for i, f := range freq {
//line /usr/local/go/src/compress/flate/huffman_code.go:108
		_go_fuzz_dep_.CoverTab[26245]++
									if f != 0 {
//line /usr/local/go/src/compress/flate/huffman_code.go:109
			_go_fuzz_dep_.CoverTab[26246]++
										total += int(f) * int(h.codes[i].len)
//line /usr/local/go/src/compress/flate/huffman_code.go:110
			// _ = "end of CoverTab[26246]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:111
			_go_fuzz_dep_.CoverTab[26247]++
//line /usr/local/go/src/compress/flate/huffman_code.go:111
			// _ = "end of CoverTab[26247]"
//line /usr/local/go/src/compress/flate/huffman_code.go:111
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:111
		// _ = "end of CoverTab[26245]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:112
	// _ = "end of CoverTab[26243]"
//line /usr/local/go/src/compress/flate/huffman_code.go:112
	_go_fuzz_dep_.CoverTab[26244]++
								return total
//line /usr/local/go/src/compress/flate/huffman_code.go:113
	// _ = "end of CoverTab[26244]"
}

const maxBitsLimit = 16

// bitCounts computes the number of literals assigned to each bit size in the Huffman encoding.
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// It is only called when list.length >= 3.
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// The cases of 0, 1, and 2 literals are handled by special case code.
//line /usr/local/go/src/compress/flate/huffman_code.go:118
//
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// list is an array of the literals with non-zero frequencies
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// and their associated frequencies. The array is in order of increasing
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// frequency and has as its last element a special element with frequency
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// MaxInt32.
//line /usr/local/go/src/compress/flate/huffman_code.go:118
//
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// maxBits is the maximum number of bits that should be used to encode any literal.
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// It must be less than 16.
//line /usr/local/go/src/compress/flate/huffman_code.go:118
//
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// bitCounts returns an integer slice in which slice[i] indicates the number of literals
//line /usr/local/go/src/compress/flate/huffman_code.go:118
// that should be encoded in i bits.
//line /usr/local/go/src/compress/flate/huffman_code.go:132
func (h *huffmanEncoder) bitCounts(list []literalNode, maxBits int32) []int32 {
//line /usr/local/go/src/compress/flate/huffman_code.go:132
	_go_fuzz_dep_.CoverTab[26248]++
								if maxBits >= maxBitsLimit {
//line /usr/local/go/src/compress/flate/huffman_code.go:133
		_go_fuzz_dep_.CoverTab[26255]++
									panic("flate: maxBits too large")
//line /usr/local/go/src/compress/flate/huffman_code.go:134
		// _ = "end of CoverTab[26255]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:135
		_go_fuzz_dep_.CoverTab[26256]++
//line /usr/local/go/src/compress/flate/huffman_code.go:135
		// _ = "end of CoverTab[26256]"
//line /usr/local/go/src/compress/flate/huffman_code.go:135
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:135
	// _ = "end of CoverTab[26248]"
//line /usr/local/go/src/compress/flate/huffman_code.go:135
	_go_fuzz_dep_.CoverTab[26249]++
								n := int32(len(list))
								list = list[0 : n+1]
								list[n] = maxNode()

//line /usr/local/go/src/compress/flate/huffman_code.go:142
	if maxBits > n-1 {
//line /usr/local/go/src/compress/flate/huffman_code.go:142
		_go_fuzz_dep_.CoverTab[26257]++
									maxBits = n - 1
//line /usr/local/go/src/compress/flate/huffman_code.go:143
		// _ = "end of CoverTab[26257]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:144
		_go_fuzz_dep_.CoverTab[26258]++
//line /usr/local/go/src/compress/flate/huffman_code.go:144
		// _ = "end of CoverTab[26258]"
//line /usr/local/go/src/compress/flate/huffman_code.go:144
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:144
	// _ = "end of CoverTab[26249]"
//line /usr/local/go/src/compress/flate/huffman_code.go:144
	_go_fuzz_dep_.CoverTab[26250]++

	// Create information about each of the levels.
	// A bogus "Level 0" whose sole purpose is so that
	// level1.prev.needed==0.  This makes level1.nextPairFreq
	// be a legitimate value that never gets chosen.
	var levels [maxBitsLimit]levelInfo
	// leafCounts[i] counts the number of literals at the left
	// of ancestors of the rightmost node at level i.
	// leafCounts[i][j] is the number of literals at the left
	// of the level j ancestor.
	var leafCounts [maxBitsLimit][maxBitsLimit]int32

	for level := int32(1); level <= maxBits; level++ {
//line /usr/local/go/src/compress/flate/huffman_code.go:157
		_go_fuzz_dep_.CoverTab[26259]++

//line /usr/local/go/src/compress/flate/huffman_code.go:160
		levels[level] = levelInfo{
			level:		level,
			lastFreq:	list[1].freq,
			nextCharFreq:	list[2].freq,
			nextPairFreq:	list[0].freq + list[1].freq,
		}
		leafCounts[level][level] = 2
		if level == 1 {
//line /usr/local/go/src/compress/flate/huffman_code.go:167
			_go_fuzz_dep_.CoverTab[26260]++
										levels[level].nextPairFreq = math.MaxInt32
//line /usr/local/go/src/compress/flate/huffman_code.go:168
			// _ = "end of CoverTab[26260]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:169
			_go_fuzz_dep_.CoverTab[26261]++
//line /usr/local/go/src/compress/flate/huffman_code.go:169
			// _ = "end of CoverTab[26261]"
//line /usr/local/go/src/compress/flate/huffman_code.go:169
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:169
		// _ = "end of CoverTab[26259]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:170
	// _ = "end of CoverTab[26250]"
//line /usr/local/go/src/compress/flate/huffman_code.go:170
	_go_fuzz_dep_.CoverTab[26251]++

//line /usr/local/go/src/compress/flate/huffman_code.go:173
	levels[maxBits].needed = 2*n - 4

	level := maxBits
	for {
//line /usr/local/go/src/compress/flate/huffman_code.go:176
		_go_fuzz_dep_.CoverTab[26262]++
									l := &levels[level]
									if l.nextPairFreq == math.MaxInt32 && func() bool {
//line /usr/local/go/src/compress/flate/huffman_code.go:178
			_go_fuzz_dep_.CoverTab[26265]++
//line /usr/local/go/src/compress/flate/huffman_code.go:178
			return l.nextCharFreq == math.MaxInt32
//line /usr/local/go/src/compress/flate/huffman_code.go:178
			// _ = "end of CoverTab[26265]"
//line /usr/local/go/src/compress/flate/huffman_code.go:178
		}() {
//line /usr/local/go/src/compress/flate/huffman_code.go:178
			_go_fuzz_dep_.CoverTab[26266]++

//line /usr/local/go/src/compress/flate/huffman_code.go:183
			l.needed = 0
										levels[level+1].nextPairFreq = math.MaxInt32
										level++
										continue
//line /usr/local/go/src/compress/flate/huffman_code.go:186
			// _ = "end of CoverTab[26266]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:187
			_go_fuzz_dep_.CoverTab[26267]++
//line /usr/local/go/src/compress/flate/huffman_code.go:187
			// _ = "end of CoverTab[26267]"
//line /usr/local/go/src/compress/flate/huffman_code.go:187
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:187
		// _ = "end of CoverTab[26262]"
//line /usr/local/go/src/compress/flate/huffman_code.go:187
		_go_fuzz_dep_.CoverTab[26263]++

									prevFreq := l.lastFreq
									if l.nextCharFreq < l.nextPairFreq {
//line /usr/local/go/src/compress/flate/huffman_code.go:190
			_go_fuzz_dep_.CoverTab[26268]++

										n := leafCounts[level][level] + 1
										l.lastFreq = l.nextCharFreq

										leafCounts[level][level] = n
										l.nextCharFreq = list[n].freq
//line /usr/local/go/src/compress/flate/huffman_code.go:196
			// _ = "end of CoverTab[26268]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:197
			_go_fuzz_dep_.CoverTab[26269]++

//line /usr/local/go/src/compress/flate/huffman_code.go:201
			l.lastFreq = l.nextPairFreq

										copy(leafCounts[level][:level], leafCounts[level-1][:level])
										levels[l.level-1].needed = 2
//line /usr/local/go/src/compress/flate/huffman_code.go:204
			// _ = "end of CoverTab[26269]"
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:205
		// _ = "end of CoverTab[26263]"
//line /usr/local/go/src/compress/flate/huffman_code.go:205
		_go_fuzz_dep_.CoverTab[26264]++

									if l.needed--; l.needed == 0 {
//line /usr/local/go/src/compress/flate/huffman_code.go:207
			_go_fuzz_dep_.CoverTab[26270]++

//line /usr/local/go/src/compress/flate/huffman_code.go:212
			if l.level == maxBits {
//line /usr/local/go/src/compress/flate/huffman_code.go:212
				_go_fuzz_dep_.CoverTab[26272]++

											break
//line /usr/local/go/src/compress/flate/huffman_code.go:214
				// _ = "end of CoverTab[26272]"
			} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:215
				_go_fuzz_dep_.CoverTab[26273]++
//line /usr/local/go/src/compress/flate/huffman_code.go:215
				// _ = "end of CoverTab[26273]"
//line /usr/local/go/src/compress/flate/huffman_code.go:215
			}
//line /usr/local/go/src/compress/flate/huffman_code.go:215
			// _ = "end of CoverTab[26270]"
//line /usr/local/go/src/compress/flate/huffman_code.go:215
			_go_fuzz_dep_.CoverTab[26271]++
										levels[l.level+1].nextPairFreq = prevFreq + l.lastFreq
										level++
//line /usr/local/go/src/compress/flate/huffman_code.go:217
			// _ = "end of CoverTab[26271]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:218
			_go_fuzz_dep_.CoverTab[26274]++

										for levels[level-1].needed > 0 {
//line /usr/local/go/src/compress/flate/huffman_code.go:220
				_go_fuzz_dep_.CoverTab[26275]++
											level--
//line /usr/local/go/src/compress/flate/huffman_code.go:221
				// _ = "end of CoverTab[26275]"
			}
//line /usr/local/go/src/compress/flate/huffman_code.go:222
			// _ = "end of CoverTab[26274]"
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:223
		// _ = "end of CoverTab[26264]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:224
	// _ = "end of CoverTab[26251]"
//line /usr/local/go/src/compress/flate/huffman_code.go:224
	_go_fuzz_dep_.CoverTab[26252]++

//line /usr/local/go/src/compress/flate/huffman_code.go:228
	if leafCounts[maxBits][maxBits] != n {
//line /usr/local/go/src/compress/flate/huffman_code.go:228
		_go_fuzz_dep_.CoverTab[26276]++
									panic("leafCounts[maxBits][maxBits] != n")
//line /usr/local/go/src/compress/flate/huffman_code.go:229
		// _ = "end of CoverTab[26276]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:230
		_go_fuzz_dep_.CoverTab[26277]++
//line /usr/local/go/src/compress/flate/huffman_code.go:230
		// _ = "end of CoverTab[26277]"
//line /usr/local/go/src/compress/flate/huffman_code.go:230
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:230
	// _ = "end of CoverTab[26252]"
//line /usr/local/go/src/compress/flate/huffman_code.go:230
	_go_fuzz_dep_.CoverTab[26253]++

								bitCount := h.bitCount[:maxBits+1]
								bits := 1
								counts := &leafCounts[maxBits]
								for level := maxBits; level > 0; level-- {
//line /usr/local/go/src/compress/flate/huffman_code.go:235
		_go_fuzz_dep_.CoverTab[26278]++

//line /usr/local/go/src/compress/flate/huffman_code.go:238
		bitCount[bits] = counts[level] - counts[level-1]
									bits++
//line /usr/local/go/src/compress/flate/huffman_code.go:239
		// _ = "end of CoverTab[26278]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:240
	// _ = "end of CoverTab[26253]"
//line /usr/local/go/src/compress/flate/huffman_code.go:240
	_go_fuzz_dep_.CoverTab[26254]++
								return bitCount
//line /usr/local/go/src/compress/flate/huffman_code.go:241
	// _ = "end of CoverTab[26254]"
}

// Look at the leaves and assign them a bit count and an encoding as specified
//line /usr/local/go/src/compress/flate/huffman_code.go:244
// in RFC 1951 3.2.2
//line /usr/local/go/src/compress/flate/huffman_code.go:246
func (h *huffmanEncoder) assignEncodingAndSize(bitCount []int32, list []literalNode) {
//line /usr/local/go/src/compress/flate/huffman_code.go:246
	_go_fuzz_dep_.CoverTab[26279]++
								code := uint16(0)
								for n, bits := range bitCount {
//line /usr/local/go/src/compress/flate/huffman_code.go:248
		_go_fuzz_dep_.CoverTab[26280]++
									code <<= 1
									if n == 0 || func() bool {
//line /usr/local/go/src/compress/flate/huffman_code.go:250
			_go_fuzz_dep_.CoverTab[26283]++
//line /usr/local/go/src/compress/flate/huffman_code.go:250
			return bits == 0
//line /usr/local/go/src/compress/flate/huffman_code.go:250
			// _ = "end of CoverTab[26283]"
//line /usr/local/go/src/compress/flate/huffman_code.go:250
		}() {
//line /usr/local/go/src/compress/flate/huffman_code.go:250
			_go_fuzz_dep_.CoverTab[26284]++
										continue
//line /usr/local/go/src/compress/flate/huffman_code.go:251
			// _ = "end of CoverTab[26284]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:252
			_go_fuzz_dep_.CoverTab[26285]++
//line /usr/local/go/src/compress/flate/huffman_code.go:252
			// _ = "end of CoverTab[26285]"
//line /usr/local/go/src/compress/flate/huffman_code.go:252
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:252
		// _ = "end of CoverTab[26280]"
//line /usr/local/go/src/compress/flate/huffman_code.go:252
		_go_fuzz_dep_.CoverTab[26281]++

//line /usr/local/go/src/compress/flate/huffman_code.go:257
		chunk := list[len(list)-int(bits):]

		h.lns.sort(chunk)
		for _, node := range chunk {
//line /usr/local/go/src/compress/flate/huffman_code.go:260
			_go_fuzz_dep_.CoverTab[26286]++
										h.codes[node.literal] = hcode{code: reverseBits(code, uint8(n)), len: uint16(n)}
										code++
//line /usr/local/go/src/compress/flate/huffman_code.go:262
			// _ = "end of CoverTab[26286]"
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:263
		// _ = "end of CoverTab[26281]"
//line /usr/local/go/src/compress/flate/huffman_code.go:263
		_go_fuzz_dep_.CoverTab[26282]++
									list = list[0 : len(list)-int(bits)]
//line /usr/local/go/src/compress/flate/huffman_code.go:264
		// _ = "end of CoverTab[26282]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:265
	// _ = "end of CoverTab[26279]"
}

// Update this Huffman Code object to be the minimum code for the specified frequency count.
//line /usr/local/go/src/compress/flate/huffman_code.go:268
//
//line /usr/local/go/src/compress/flate/huffman_code.go:268
// freq is an array of frequencies, in which freq[i] gives the frequency of literal i.
//line /usr/local/go/src/compress/flate/huffman_code.go:268
// maxBits  The maximum number of bits to use for any literal.
//line /usr/local/go/src/compress/flate/huffman_code.go:272
func (h *huffmanEncoder) generate(freq []int32, maxBits int32) {
//line /usr/local/go/src/compress/flate/huffman_code.go:272
	_go_fuzz_dep_.CoverTab[26287]++
								if h.freqcache == nil {
//line /usr/local/go/src/compress/flate/huffman_code.go:273
		_go_fuzz_dep_.CoverTab[26291]++

//line /usr/local/go/src/compress/flate/huffman_code.go:277
		h.freqcache = make([]literalNode, maxNumLit+1)
//line /usr/local/go/src/compress/flate/huffman_code.go:277
		// _ = "end of CoverTab[26291]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:278
		_go_fuzz_dep_.CoverTab[26292]++
//line /usr/local/go/src/compress/flate/huffman_code.go:278
		// _ = "end of CoverTab[26292]"
//line /usr/local/go/src/compress/flate/huffman_code.go:278
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:278
	// _ = "end of CoverTab[26287]"
//line /usr/local/go/src/compress/flate/huffman_code.go:278
	_go_fuzz_dep_.CoverTab[26288]++
								list := h.freqcache[:len(freq)+1]

								count := 0

								for i, f := range freq {
//line /usr/local/go/src/compress/flate/huffman_code.go:283
		_go_fuzz_dep_.CoverTab[26293]++
									if f != 0 {
//line /usr/local/go/src/compress/flate/huffman_code.go:284
			_go_fuzz_dep_.CoverTab[26294]++
										list[count] = literalNode{uint16(i), f}
										count++
//line /usr/local/go/src/compress/flate/huffman_code.go:286
			// _ = "end of CoverTab[26294]"
		} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:287
			_go_fuzz_dep_.CoverTab[26295]++
										h.codes[i].len = 0
//line /usr/local/go/src/compress/flate/huffman_code.go:288
			// _ = "end of CoverTab[26295]"
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:289
		// _ = "end of CoverTab[26293]"
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:290
	// _ = "end of CoverTab[26288]"
//line /usr/local/go/src/compress/flate/huffman_code.go:290
	_go_fuzz_dep_.CoverTab[26289]++

								list = list[:count]
								if count <= 2 {
//line /usr/local/go/src/compress/flate/huffman_code.go:293
		_go_fuzz_dep_.CoverTab[26296]++

//line /usr/local/go/src/compress/flate/huffman_code.go:296
		for i, node := range list {
//line /usr/local/go/src/compress/flate/huffman_code.go:296
			_go_fuzz_dep_.CoverTab[26298]++

										h.codes[node.literal].set(uint16(i), 1)
//line /usr/local/go/src/compress/flate/huffman_code.go:298
			// _ = "end of CoverTab[26298]"
		}
//line /usr/local/go/src/compress/flate/huffman_code.go:299
		// _ = "end of CoverTab[26296]"
//line /usr/local/go/src/compress/flate/huffman_code.go:299
		_go_fuzz_dep_.CoverTab[26297]++
									return
//line /usr/local/go/src/compress/flate/huffman_code.go:300
		// _ = "end of CoverTab[26297]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:301
		_go_fuzz_dep_.CoverTab[26299]++
//line /usr/local/go/src/compress/flate/huffman_code.go:301
		// _ = "end of CoverTab[26299]"
//line /usr/local/go/src/compress/flate/huffman_code.go:301
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:301
	// _ = "end of CoverTab[26289]"
//line /usr/local/go/src/compress/flate/huffman_code.go:301
	_go_fuzz_dep_.CoverTab[26290]++
								h.lfs.sort(list)

//line /usr/local/go/src/compress/flate/huffman_code.go:305
	bitCount := h.bitCounts(list, maxBits)

								h.assignEncodingAndSize(bitCount, list)
//line /usr/local/go/src/compress/flate/huffman_code.go:307
	// _ = "end of CoverTab[26290]"
}

type byLiteral []literalNode

func (s *byLiteral) sort(a []literalNode) {
//line /usr/local/go/src/compress/flate/huffman_code.go:312
	_go_fuzz_dep_.CoverTab[26300]++
								*s = byLiteral(a)
								sort.Sort(s)
//line /usr/local/go/src/compress/flate/huffman_code.go:314
	// _ = "end of CoverTab[26300]"
}

func (s byLiteral) Len() int {
//line /usr/local/go/src/compress/flate/huffman_code.go:317
	_go_fuzz_dep_.CoverTab[26301]++
//line /usr/local/go/src/compress/flate/huffman_code.go:317
	return len(s)
//line /usr/local/go/src/compress/flate/huffman_code.go:317
	// _ = "end of CoverTab[26301]"
//line /usr/local/go/src/compress/flate/huffman_code.go:317
}

func (s byLiteral) Less(i, j int) bool {
//line /usr/local/go/src/compress/flate/huffman_code.go:319
	_go_fuzz_dep_.CoverTab[26302]++
								return s[i].literal < s[j].literal
//line /usr/local/go/src/compress/flate/huffman_code.go:320
	// _ = "end of CoverTab[26302]"
}

func (s byLiteral) Swap(i, j int) {
//line /usr/local/go/src/compress/flate/huffman_code.go:323
	_go_fuzz_dep_.CoverTab[26303]++
//line /usr/local/go/src/compress/flate/huffman_code.go:323
	s[i], s[j] = s[j], s[i]
//line /usr/local/go/src/compress/flate/huffman_code.go:323
	// _ = "end of CoverTab[26303]"
//line /usr/local/go/src/compress/flate/huffman_code.go:323
}

type byFreq []literalNode

func (s *byFreq) sort(a []literalNode) {
//line /usr/local/go/src/compress/flate/huffman_code.go:327
	_go_fuzz_dep_.CoverTab[26304]++
								*s = byFreq(a)
								sort.Sort(s)
//line /usr/local/go/src/compress/flate/huffman_code.go:329
	// _ = "end of CoverTab[26304]"
}

func (s byFreq) Len() int {
//line /usr/local/go/src/compress/flate/huffman_code.go:332
	_go_fuzz_dep_.CoverTab[26305]++
//line /usr/local/go/src/compress/flate/huffman_code.go:332
	return len(s)
//line /usr/local/go/src/compress/flate/huffman_code.go:332
	// _ = "end of CoverTab[26305]"
//line /usr/local/go/src/compress/flate/huffman_code.go:332
}

func (s byFreq) Less(i, j int) bool {
//line /usr/local/go/src/compress/flate/huffman_code.go:334
	_go_fuzz_dep_.CoverTab[26306]++
								if s[i].freq == s[j].freq {
//line /usr/local/go/src/compress/flate/huffman_code.go:335
		_go_fuzz_dep_.CoverTab[26308]++
									return s[i].literal < s[j].literal
//line /usr/local/go/src/compress/flate/huffman_code.go:336
		// _ = "end of CoverTab[26308]"
	} else {
//line /usr/local/go/src/compress/flate/huffman_code.go:337
		_go_fuzz_dep_.CoverTab[26309]++
//line /usr/local/go/src/compress/flate/huffman_code.go:337
		// _ = "end of CoverTab[26309]"
//line /usr/local/go/src/compress/flate/huffman_code.go:337
	}
//line /usr/local/go/src/compress/flate/huffman_code.go:337
	// _ = "end of CoverTab[26306]"
//line /usr/local/go/src/compress/flate/huffman_code.go:337
	_go_fuzz_dep_.CoverTab[26307]++
								return s[i].freq < s[j].freq
//line /usr/local/go/src/compress/flate/huffman_code.go:338
	// _ = "end of CoverTab[26307]"
}

func (s byFreq) Swap(i, j int) {
//line /usr/local/go/src/compress/flate/huffman_code.go:341
	_go_fuzz_dep_.CoverTab[26310]++
//line /usr/local/go/src/compress/flate/huffman_code.go:341
	s[i], s[j] = s[j], s[i]
//line /usr/local/go/src/compress/flate/huffman_code.go:341
	// _ = "end of CoverTab[26310]"
//line /usr/local/go/src/compress/flate/huffman_code.go:341
}

func reverseBits(number uint16, bitLength byte) uint16 {
//line /usr/local/go/src/compress/flate/huffman_code.go:343
	_go_fuzz_dep_.CoverTab[26311]++
								return bits.Reverse16(number << (16 - bitLength))
//line /usr/local/go/src/compress/flate/huffman_code.go:344
	// _ = "end of CoverTab[26311]"
}

//line /usr/local/go/src/compress/flate/huffman_code.go:345
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/huffman_code.go:345
var _ = _go_fuzz_dep_.CoverTab
