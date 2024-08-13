// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/flate/deflate.go:5
package flate

//line /usr/local/go/src/compress/flate/deflate.go:5
import (
//line /usr/local/go/src/compress/flate/deflate.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/flate/deflate.go:5
)
//line /usr/local/go/src/compress/flate/deflate.go:5
import (
//line /usr/local/go/src/compress/flate/deflate.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/flate/deflate.go:5
)

import (
	"errors"
	"fmt"
	"io"
	"math"
)

const (
	NoCompression		= 0
	BestSpeed		= 1
	BestCompression		= 9
	DefaultCompression	= -1

	// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman
	// entropy encoding. This mode is useful in compressing data that has
	// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)
	// that lacks an entropy encoder. Compression gains are achieved when
	// certain bytes in the input stream occur more frequently than others.
	//
	// Note that HuffmanOnly produces a compressed output that is
	// RFC 1951 compliant. That is, any valid DEFLATE decompressor will
	// continue to be able to decompress this output.
	HuffmanOnly	= -2
)

const (
	logWindowSize	= 15
	windowSize	= 1 << logWindowSize
	windowMask	= windowSize - 1

	// The LZ77 step produces a sequence of literal tokens and <length, offset>
	// pair tokens. The offset is also known as distance. The underlying wire
	// format limits the range of lengths and offsets. For example, there are
	// 256 legitimate lengths: those in the range [3, 258]. This package's
	// compressor uses a higher minimum match length, enabling optimizations
	// such as finding matches via 32-bit loads and compares.
	baseMatchLength	= 3		// The smallest match length per the RFC section 3.2.5
	minMatchLength	= 4		// The smallest match length that the compressor actually emits
	maxMatchLength	= 258		// The largest match length
	baseMatchOffset	= 1		// The smallest match offset
	maxMatchOffset	= 1 << 15	// The largest match offset

	// The maximum number of tokens we put into a single flate block, just to
	// stop things from getting too large.
	maxFlateBlockTokens	= 1 << 14
	maxStoreBlockSize	= 65535
	hashBits		= 17	// After 17 performance degrades
	hashSize		= 1 << hashBits
	hashMask		= (1 << hashBits) - 1
	maxHashOffset		= 1 << 24

	skipNever	= math.MaxInt32
)

type compressionLevel struct {
	level, good, lazy, nice, chain, fastSkipHashing int
}

var levels = []compressionLevel{
							{0, 0, 0, 0, 0, 0},
							{1, 0, 0, 0, 0, 0},

							{2, 4, 0, 16, 8, 5},
							{3, 4, 0, 32, 32, 6},

//line /usr/local/go/src/compress/flate/deflate.go:73
	{4, 4, 4, 16, 16, skipNever},
							{5, 8, 16, 32, 32, skipNever},
							{6, 8, 16, 128, 128, skipNever},
							{7, 8, 32, 128, 256, skipNever},
							{8, 32, 128, 258, 1024, skipNever},
							{9, 32, 258, 258, 4096, skipNever},
}

type compressor struct {
	compressionLevel

	w		*huffmanBitWriter
	bulkHasher	func([]byte, []uint32)

	// compression algorithm
	fill		func(*compressor, []byte) int	// copy data to window
	step		func(*compressor)		// process window
	sync		bool				// requesting flush
	bestSpeed	*deflateFast			// Encoder for BestSpeed

	// Input hash chains
	// hashHead[hashValue] contains the largest inputIndex with the specified hash value
	// If hashHead[hashValue] is within the current window, then
	// hashPrev[hashHead[hashValue] & windowMask] contains the previous index
	// with the same hash value.
	chainHead	int
	hashHead	[hashSize]uint32
	hashPrev	[windowSize]uint32
	hashOffset	int

	// input window: unprocessed data is window[index:windowEnd]
	index		int
	window		[]byte
	windowEnd	int
	blockStart	int	// window index where current tokens start
	byteAvailable	bool	// if true, still need to process window[index-1].

	// queued output tokens
	tokens	[]token

	// deflate state
	length		int
	offset		int
	maxInsertIndex	int
	err		error

	// hashMatch must be able to contain hashes for the maximum match length.
	hashMatch	[maxMatchLength - 1]uint32
}

func (d *compressor) fillDeflate(b []byte) int {
//line /usr/local/go/src/compress/flate/deflate.go:123
	_go_fuzz_dep_.CoverTab[25671]++
							if d.index >= 2*windowSize-(minMatchLength+maxMatchLength) {
//line /usr/local/go/src/compress/flate/deflate.go:124
		_go_fuzz_dep_.CoverTab[25673]++

								copy(d.window, d.window[windowSize:2*windowSize])
								d.index -= windowSize
								d.windowEnd -= windowSize
								if d.blockStart >= windowSize {
//line /usr/local/go/src/compress/flate/deflate.go:129
			_go_fuzz_dep_.CoverTab[25675]++
									d.blockStart -= windowSize
//line /usr/local/go/src/compress/flate/deflate.go:130
			// _ = "end of CoverTab[25675]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:131
			_go_fuzz_dep_.CoverTab[25676]++
									d.blockStart = math.MaxInt32
//line /usr/local/go/src/compress/flate/deflate.go:132
			// _ = "end of CoverTab[25676]"
		}
//line /usr/local/go/src/compress/flate/deflate.go:133
		// _ = "end of CoverTab[25673]"
//line /usr/local/go/src/compress/flate/deflate.go:133
		_go_fuzz_dep_.CoverTab[25674]++
								d.hashOffset += windowSize
								if d.hashOffset > maxHashOffset {
//line /usr/local/go/src/compress/flate/deflate.go:135
			_go_fuzz_dep_.CoverTab[25677]++
									delta := d.hashOffset - 1
									d.hashOffset -= delta
									d.chainHead -= delta

//line /usr/local/go/src/compress/flate/deflate.go:142
			for i, v := range d.hashPrev[:] {
//line /usr/local/go/src/compress/flate/deflate.go:142
				_go_fuzz_dep_.CoverTab[25679]++
										if int(v) > delta {
//line /usr/local/go/src/compress/flate/deflate.go:143
					_go_fuzz_dep_.CoverTab[25680]++
											d.hashPrev[i] = uint32(int(v) - delta)
//line /usr/local/go/src/compress/flate/deflate.go:144
					// _ = "end of CoverTab[25680]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:145
					_go_fuzz_dep_.CoverTab[25681]++
											d.hashPrev[i] = 0
//line /usr/local/go/src/compress/flate/deflate.go:146
					// _ = "end of CoverTab[25681]"
				}
//line /usr/local/go/src/compress/flate/deflate.go:147
				// _ = "end of CoverTab[25679]"
			}
//line /usr/local/go/src/compress/flate/deflate.go:148
			// _ = "end of CoverTab[25677]"
//line /usr/local/go/src/compress/flate/deflate.go:148
			_go_fuzz_dep_.CoverTab[25678]++
									for i, v := range d.hashHead[:] {
//line /usr/local/go/src/compress/flate/deflate.go:149
				_go_fuzz_dep_.CoverTab[25682]++
										if int(v) > delta {
//line /usr/local/go/src/compress/flate/deflate.go:150
					_go_fuzz_dep_.CoverTab[25683]++
											d.hashHead[i] = uint32(int(v) - delta)
//line /usr/local/go/src/compress/flate/deflate.go:151
					// _ = "end of CoverTab[25683]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:152
					_go_fuzz_dep_.CoverTab[25684]++
											d.hashHead[i] = 0
//line /usr/local/go/src/compress/flate/deflate.go:153
					// _ = "end of CoverTab[25684]"
				}
//line /usr/local/go/src/compress/flate/deflate.go:154
				// _ = "end of CoverTab[25682]"
			}
//line /usr/local/go/src/compress/flate/deflate.go:155
			// _ = "end of CoverTab[25678]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:156
			_go_fuzz_dep_.CoverTab[25685]++
//line /usr/local/go/src/compress/flate/deflate.go:156
			// _ = "end of CoverTab[25685]"
//line /usr/local/go/src/compress/flate/deflate.go:156
		}
//line /usr/local/go/src/compress/flate/deflate.go:156
		// _ = "end of CoverTab[25674]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:157
		_go_fuzz_dep_.CoverTab[25686]++
//line /usr/local/go/src/compress/flate/deflate.go:157
		// _ = "end of CoverTab[25686]"
//line /usr/local/go/src/compress/flate/deflate.go:157
	}
//line /usr/local/go/src/compress/flate/deflate.go:157
	// _ = "end of CoverTab[25671]"
//line /usr/local/go/src/compress/flate/deflate.go:157
	_go_fuzz_dep_.CoverTab[25672]++
							n := copy(d.window[d.windowEnd:], b)
							d.windowEnd += n
							return n
//line /usr/local/go/src/compress/flate/deflate.go:160
	// _ = "end of CoverTab[25672]"
}

func (d *compressor) writeBlock(tokens []token, index int) error {
//line /usr/local/go/src/compress/flate/deflate.go:163
	_go_fuzz_dep_.CoverTab[25687]++
							if index > 0 {
//line /usr/local/go/src/compress/flate/deflate.go:164
		_go_fuzz_dep_.CoverTab[25689]++
								var window []byte
								if d.blockStart <= index {
//line /usr/local/go/src/compress/flate/deflate.go:166
			_go_fuzz_dep_.CoverTab[25691]++
									window = d.window[d.blockStart:index]
//line /usr/local/go/src/compress/flate/deflate.go:167
			// _ = "end of CoverTab[25691]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:168
			_go_fuzz_dep_.CoverTab[25692]++
//line /usr/local/go/src/compress/flate/deflate.go:168
			// _ = "end of CoverTab[25692]"
//line /usr/local/go/src/compress/flate/deflate.go:168
		}
//line /usr/local/go/src/compress/flate/deflate.go:168
		// _ = "end of CoverTab[25689]"
//line /usr/local/go/src/compress/flate/deflate.go:168
		_go_fuzz_dep_.CoverTab[25690]++
								d.blockStart = index
								d.w.writeBlock(tokens, false, window)
								return d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:171
		// _ = "end of CoverTab[25690]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:172
		_go_fuzz_dep_.CoverTab[25693]++
//line /usr/local/go/src/compress/flate/deflate.go:172
		// _ = "end of CoverTab[25693]"
//line /usr/local/go/src/compress/flate/deflate.go:172
	}
//line /usr/local/go/src/compress/flate/deflate.go:172
	// _ = "end of CoverTab[25687]"
//line /usr/local/go/src/compress/flate/deflate.go:172
	_go_fuzz_dep_.CoverTab[25688]++
							return nil
//line /usr/local/go/src/compress/flate/deflate.go:173
	// _ = "end of CoverTab[25688]"
}

// fillWindow will fill the current window with the supplied
//line /usr/local/go/src/compress/flate/deflate.go:176
// dictionary and calculate all hashes.
//line /usr/local/go/src/compress/flate/deflate.go:176
// This is much faster than doing a full encode.
//line /usr/local/go/src/compress/flate/deflate.go:176
// Should only be used after a reset.
//line /usr/local/go/src/compress/flate/deflate.go:180
func (d *compressor) fillWindow(b []byte) {
//line /usr/local/go/src/compress/flate/deflate.go:180
	_go_fuzz_dep_.CoverTab[25694]++

							if d.compressionLevel.level < 2 {
//line /usr/local/go/src/compress/flate/deflate.go:182
		_go_fuzz_dep_.CoverTab[25699]++
								return
//line /usr/local/go/src/compress/flate/deflate.go:183
		// _ = "end of CoverTab[25699]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:184
		_go_fuzz_dep_.CoverTab[25700]++
//line /usr/local/go/src/compress/flate/deflate.go:184
		// _ = "end of CoverTab[25700]"
//line /usr/local/go/src/compress/flate/deflate.go:184
	}
//line /usr/local/go/src/compress/flate/deflate.go:184
	// _ = "end of CoverTab[25694]"
//line /usr/local/go/src/compress/flate/deflate.go:184
	_go_fuzz_dep_.CoverTab[25695]++
							if d.index != 0 || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:185
		_go_fuzz_dep_.CoverTab[25701]++
//line /usr/local/go/src/compress/flate/deflate.go:185
		return d.windowEnd != 0
//line /usr/local/go/src/compress/flate/deflate.go:185
		// _ = "end of CoverTab[25701]"
//line /usr/local/go/src/compress/flate/deflate.go:185
	}() {
//line /usr/local/go/src/compress/flate/deflate.go:185
		_go_fuzz_dep_.CoverTab[25702]++
								panic("internal error: fillWindow called with stale data")
//line /usr/local/go/src/compress/flate/deflate.go:186
		// _ = "end of CoverTab[25702]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:187
		_go_fuzz_dep_.CoverTab[25703]++
//line /usr/local/go/src/compress/flate/deflate.go:187
		// _ = "end of CoverTab[25703]"
//line /usr/local/go/src/compress/flate/deflate.go:187
	}
//line /usr/local/go/src/compress/flate/deflate.go:187
	// _ = "end of CoverTab[25695]"
//line /usr/local/go/src/compress/flate/deflate.go:187
	_go_fuzz_dep_.CoverTab[25696]++

//line /usr/local/go/src/compress/flate/deflate.go:190
	if len(b) > windowSize {
//line /usr/local/go/src/compress/flate/deflate.go:190
		_go_fuzz_dep_.CoverTab[25704]++
								b = b[len(b)-windowSize:]
//line /usr/local/go/src/compress/flate/deflate.go:191
		// _ = "end of CoverTab[25704]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:192
		_go_fuzz_dep_.CoverTab[25705]++
//line /usr/local/go/src/compress/flate/deflate.go:192
		// _ = "end of CoverTab[25705]"
//line /usr/local/go/src/compress/flate/deflate.go:192
	}
//line /usr/local/go/src/compress/flate/deflate.go:192
	// _ = "end of CoverTab[25696]"
//line /usr/local/go/src/compress/flate/deflate.go:192
	_go_fuzz_dep_.CoverTab[25697]++

							n := copy(d.window, b)

//line /usr/local/go/src/compress/flate/deflate.go:197
	loops := (n + 256 - minMatchLength) / 256
	for j := 0; j < loops; j++ {
//line /usr/local/go/src/compress/flate/deflate.go:198
		_go_fuzz_dep_.CoverTab[25706]++
								index := j * 256
								end := index + 256 + minMatchLength - 1
								if end > n {
//line /usr/local/go/src/compress/flate/deflate.go:201
			_go_fuzz_dep_.CoverTab[25709]++
									end = n
//line /usr/local/go/src/compress/flate/deflate.go:202
			// _ = "end of CoverTab[25709]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:203
			_go_fuzz_dep_.CoverTab[25710]++
//line /usr/local/go/src/compress/flate/deflate.go:203
			// _ = "end of CoverTab[25710]"
//line /usr/local/go/src/compress/flate/deflate.go:203
		}
//line /usr/local/go/src/compress/flate/deflate.go:203
		// _ = "end of CoverTab[25706]"
//line /usr/local/go/src/compress/flate/deflate.go:203
		_go_fuzz_dep_.CoverTab[25707]++
								toCheck := d.window[index:end]
								dstSize := len(toCheck) - minMatchLength + 1

								if dstSize <= 0 {
//line /usr/local/go/src/compress/flate/deflate.go:207
			_go_fuzz_dep_.CoverTab[25711]++
									continue
//line /usr/local/go/src/compress/flate/deflate.go:208
			// _ = "end of CoverTab[25711]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:209
			_go_fuzz_dep_.CoverTab[25712]++
//line /usr/local/go/src/compress/flate/deflate.go:209
			// _ = "end of CoverTab[25712]"
//line /usr/local/go/src/compress/flate/deflate.go:209
		}
//line /usr/local/go/src/compress/flate/deflate.go:209
		// _ = "end of CoverTab[25707]"
//line /usr/local/go/src/compress/flate/deflate.go:209
		_go_fuzz_dep_.CoverTab[25708]++

								dst := d.hashMatch[:dstSize]
								d.bulkHasher(toCheck, dst)
								for i, val := range dst {
//line /usr/local/go/src/compress/flate/deflate.go:213
			_go_fuzz_dep_.CoverTab[25713]++
									di := i + index
									hh := &d.hashHead[val&hashMask]

//line /usr/local/go/src/compress/flate/deflate.go:218
			d.hashPrev[di&windowMask] = *hh

									*hh = uint32(di + d.hashOffset)
//line /usr/local/go/src/compress/flate/deflate.go:220
			// _ = "end of CoverTab[25713]"
		}
//line /usr/local/go/src/compress/flate/deflate.go:221
		// _ = "end of CoverTab[25708]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:222
	// _ = "end of CoverTab[25697]"
//line /usr/local/go/src/compress/flate/deflate.go:222
	_go_fuzz_dep_.CoverTab[25698]++

							d.windowEnd = n
							d.index = n
//line /usr/local/go/src/compress/flate/deflate.go:225
	// _ = "end of CoverTab[25698]"
}

// Try to find a match starting at index whose length is greater than prevSize.
//line /usr/local/go/src/compress/flate/deflate.go:228
// We only look at chainCount possibilities before giving up.
//line /usr/local/go/src/compress/flate/deflate.go:230
func (d *compressor) findMatch(pos int, prevHead int, prevLength int, lookahead int) (length, offset int, ok bool) {
//line /usr/local/go/src/compress/flate/deflate.go:230
	_go_fuzz_dep_.CoverTab[25714]++
							minMatchLook := maxMatchLength
							if lookahead < minMatchLook {
//line /usr/local/go/src/compress/flate/deflate.go:232
		_go_fuzz_dep_.CoverTab[25719]++
								minMatchLook = lookahead
//line /usr/local/go/src/compress/flate/deflate.go:233
		// _ = "end of CoverTab[25719]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:234
		_go_fuzz_dep_.CoverTab[25720]++
//line /usr/local/go/src/compress/flate/deflate.go:234
		// _ = "end of CoverTab[25720]"
//line /usr/local/go/src/compress/flate/deflate.go:234
	}
//line /usr/local/go/src/compress/flate/deflate.go:234
	// _ = "end of CoverTab[25714]"
//line /usr/local/go/src/compress/flate/deflate.go:234
	_go_fuzz_dep_.CoverTab[25715]++

							win := d.window[0 : pos+minMatchLook]

//line /usr/local/go/src/compress/flate/deflate.go:239
	nice := len(win) - pos
	if d.nice < nice {
//line /usr/local/go/src/compress/flate/deflate.go:240
		_go_fuzz_dep_.CoverTab[25721]++
								nice = d.nice
//line /usr/local/go/src/compress/flate/deflate.go:241
		// _ = "end of CoverTab[25721]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:242
		_go_fuzz_dep_.CoverTab[25722]++
//line /usr/local/go/src/compress/flate/deflate.go:242
		// _ = "end of CoverTab[25722]"
//line /usr/local/go/src/compress/flate/deflate.go:242
	}
//line /usr/local/go/src/compress/flate/deflate.go:242
	// _ = "end of CoverTab[25715]"
//line /usr/local/go/src/compress/flate/deflate.go:242
	_go_fuzz_dep_.CoverTab[25716]++

//line /usr/local/go/src/compress/flate/deflate.go:245
	tries := d.chain
	length = prevLength
	if length >= d.good {
//line /usr/local/go/src/compress/flate/deflate.go:247
		_go_fuzz_dep_.CoverTab[25723]++
								tries >>= 2
//line /usr/local/go/src/compress/flate/deflate.go:248
		// _ = "end of CoverTab[25723]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:249
		_go_fuzz_dep_.CoverTab[25724]++
//line /usr/local/go/src/compress/flate/deflate.go:249
		// _ = "end of CoverTab[25724]"
//line /usr/local/go/src/compress/flate/deflate.go:249
	}
//line /usr/local/go/src/compress/flate/deflate.go:249
	// _ = "end of CoverTab[25716]"
//line /usr/local/go/src/compress/flate/deflate.go:249
	_go_fuzz_dep_.CoverTab[25717]++

							wEnd := win[pos+length]
							wPos := win[pos:]
							minIndex := pos - windowSize

							for i := prevHead; tries > 0; tries-- {
//line /usr/local/go/src/compress/flate/deflate.go:255
		_go_fuzz_dep_.CoverTab[25725]++
								if wEnd == win[i+length] {
//line /usr/local/go/src/compress/flate/deflate.go:256
			_go_fuzz_dep_.CoverTab[25728]++
									n := matchLen(win[i:], wPos, minMatchLook)

									if n > length && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:259
				_go_fuzz_dep_.CoverTab[25729]++
//line /usr/local/go/src/compress/flate/deflate.go:259
				return (n > minMatchLength || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:259
					_go_fuzz_dep_.CoverTab[25730]++
//line /usr/local/go/src/compress/flate/deflate.go:259
					return pos-i <= 4096
//line /usr/local/go/src/compress/flate/deflate.go:259
					// _ = "end of CoverTab[25730]"
//line /usr/local/go/src/compress/flate/deflate.go:259
				}())
//line /usr/local/go/src/compress/flate/deflate.go:259
				// _ = "end of CoverTab[25729]"
//line /usr/local/go/src/compress/flate/deflate.go:259
			}() {
//line /usr/local/go/src/compress/flate/deflate.go:259
				_go_fuzz_dep_.CoverTab[25731]++
										length = n
										offset = pos - i
										ok = true
										if n >= nice {
//line /usr/local/go/src/compress/flate/deflate.go:263
					_go_fuzz_dep_.CoverTab[25733]++

											break
//line /usr/local/go/src/compress/flate/deflate.go:265
					// _ = "end of CoverTab[25733]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:266
					_go_fuzz_dep_.CoverTab[25734]++
//line /usr/local/go/src/compress/flate/deflate.go:266
					// _ = "end of CoverTab[25734]"
//line /usr/local/go/src/compress/flate/deflate.go:266
				}
//line /usr/local/go/src/compress/flate/deflate.go:266
				// _ = "end of CoverTab[25731]"
//line /usr/local/go/src/compress/flate/deflate.go:266
				_go_fuzz_dep_.CoverTab[25732]++
										wEnd = win[pos+n]
//line /usr/local/go/src/compress/flate/deflate.go:267
				// _ = "end of CoverTab[25732]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:268
				_go_fuzz_dep_.CoverTab[25735]++
//line /usr/local/go/src/compress/flate/deflate.go:268
				// _ = "end of CoverTab[25735]"
//line /usr/local/go/src/compress/flate/deflate.go:268
			}
//line /usr/local/go/src/compress/flate/deflate.go:268
			// _ = "end of CoverTab[25728]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:269
			_go_fuzz_dep_.CoverTab[25736]++
//line /usr/local/go/src/compress/flate/deflate.go:269
			// _ = "end of CoverTab[25736]"
//line /usr/local/go/src/compress/flate/deflate.go:269
		}
//line /usr/local/go/src/compress/flate/deflate.go:269
		// _ = "end of CoverTab[25725]"
//line /usr/local/go/src/compress/flate/deflate.go:269
		_go_fuzz_dep_.CoverTab[25726]++
								if i == minIndex {
//line /usr/local/go/src/compress/flate/deflate.go:270
			_go_fuzz_dep_.CoverTab[25737]++

									break
//line /usr/local/go/src/compress/flate/deflate.go:272
			// _ = "end of CoverTab[25737]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:273
			_go_fuzz_dep_.CoverTab[25738]++
//line /usr/local/go/src/compress/flate/deflate.go:273
			// _ = "end of CoverTab[25738]"
//line /usr/local/go/src/compress/flate/deflate.go:273
		}
//line /usr/local/go/src/compress/flate/deflate.go:273
		// _ = "end of CoverTab[25726]"
//line /usr/local/go/src/compress/flate/deflate.go:273
		_go_fuzz_dep_.CoverTab[25727]++
								i = int(d.hashPrev[i&windowMask]) - d.hashOffset
								if i < minIndex || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:275
			_go_fuzz_dep_.CoverTab[25739]++
//line /usr/local/go/src/compress/flate/deflate.go:275
			return i < 0
//line /usr/local/go/src/compress/flate/deflate.go:275
			// _ = "end of CoverTab[25739]"
//line /usr/local/go/src/compress/flate/deflate.go:275
		}() {
//line /usr/local/go/src/compress/flate/deflate.go:275
			_go_fuzz_dep_.CoverTab[25740]++
									break
//line /usr/local/go/src/compress/flate/deflate.go:276
			// _ = "end of CoverTab[25740]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:277
			_go_fuzz_dep_.CoverTab[25741]++
//line /usr/local/go/src/compress/flate/deflate.go:277
			// _ = "end of CoverTab[25741]"
//line /usr/local/go/src/compress/flate/deflate.go:277
		}
//line /usr/local/go/src/compress/flate/deflate.go:277
		// _ = "end of CoverTab[25727]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:278
	// _ = "end of CoverTab[25717]"
//line /usr/local/go/src/compress/flate/deflate.go:278
	_go_fuzz_dep_.CoverTab[25718]++
							return
//line /usr/local/go/src/compress/flate/deflate.go:279
	// _ = "end of CoverTab[25718]"
}

func (d *compressor) writeStoredBlock(buf []byte) error {
//line /usr/local/go/src/compress/flate/deflate.go:282
	_go_fuzz_dep_.CoverTab[25742]++
							if d.w.writeStoredHeader(len(buf), false); d.w.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:283
		_go_fuzz_dep_.CoverTab[25744]++
								return d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:284
		// _ = "end of CoverTab[25744]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:285
		_go_fuzz_dep_.CoverTab[25745]++
//line /usr/local/go/src/compress/flate/deflate.go:285
		// _ = "end of CoverTab[25745]"
//line /usr/local/go/src/compress/flate/deflate.go:285
	}
//line /usr/local/go/src/compress/flate/deflate.go:285
	// _ = "end of CoverTab[25742]"
//line /usr/local/go/src/compress/flate/deflate.go:285
	_go_fuzz_dep_.CoverTab[25743]++
							d.w.writeBytes(buf)
							return d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:287
	// _ = "end of CoverTab[25743]"
}

const hashmul = 0x1e35a7bd

// hash4 returns a hash representation of the first 4 bytes
//line /usr/local/go/src/compress/flate/deflate.go:292
// of the supplied slice.
//line /usr/local/go/src/compress/flate/deflate.go:292
// The caller must ensure that len(b) >= 4.
//line /usr/local/go/src/compress/flate/deflate.go:295
func hash4(b []byte) uint32 {
//line /usr/local/go/src/compress/flate/deflate.go:295
	_go_fuzz_dep_.CoverTab[25746]++
							return ((uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24) * hashmul) >> (32 - hashBits)
//line /usr/local/go/src/compress/flate/deflate.go:296
	// _ = "end of CoverTab[25746]"
}

// bulkHash4 will compute hashes using the same
//line /usr/local/go/src/compress/flate/deflate.go:299
// algorithm as hash4.
//line /usr/local/go/src/compress/flate/deflate.go:301
func bulkHash4(b []byte, dst []uint32) {
//line /usr/local/go/src/compress/flate/deflate.go:301
	_go_fuzz_dep_.CoverTab[25747]++
							if len(b) < minMatchLength {
//line /usr/local/go/src/compress/flate/deflate.go:302
		_go_fuzz_dep_.CoverTab[25749]++
								return
//line /usr/local/go/src/compress/flate/deflate.go:303
		// _ = "end of CoverTab[25749]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:304
		_go_fuzz_dep_.CoverTab[25750]++
//line /usr/local/go/src/compress/flate/deflate.go:304
		// _ = "end of CoverTab[25750]"
//line /usr/local/go/src/compress/flate/deflate.go:304
	}
//line /usr/local/go/src/compress/flate/deflate.go:304
	// _ = "end of CoverTab[25747]"
//line /usr/local/go/src/compress/flate/deflate.go:304
	_go_fuzz_dep_.CoverTab[25748]++
							hb := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
							dst[0] = (hb * hashmul) >> (32 - hashBits)
							end := len(b) - minMatchLength + 1
							for i := 1; i < end; i++ {
//line /usr/local/go/src/compress/flate/deflate.go:308
		_go_fuzz_dep_.CoverTab[25751]++
								hb = (hb << 8) | uint32(b[i+3])
								dst[i] = (hb * hashmul) >> (32 - hashBits)
//line /usr/local/go/src/compress/flate/deflate.go:310
		// _ = "end of CoverTab[25751]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:311
	// _ = "end of CoverTab[25748]"
}

// matchLen returns the number of matching bytes in a and b
//line /usr/local/go/src/compress/flate/deflate.go:314
// up to length 'max'. Both slices must be at least 'max'
//line /usr/local/go/src/compress/flate/deflate.go:314
// bytes in size.
//line /usr/local/go/src/compress/flate/deflate.go:317
func matchLen(a, b []byte, max int) int {
//line /usr/local/go/src/compress/flate/deflate.go:317
	_go_fuzz_dep_.CoverTab[25752]++
							a = a[:max]
							b = b[:len(a)]
							for i, av := range a {
//line /usr/local/go/src/compress/flate/deflate.go:320
		_go_fuzz_dep_.CoverTab[25754]++
								if b[i] != av {
//line /usr/local/go/src/compress/flate/deflate.go:321
			_go_fuzz_dep_.CoverTab[25755]++
									return i
//line /usr/local/go/src/compress/flate/deflate.go:322
			// _ = "end of CoverTab[25755]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:323
			_go_fuzz_dep_.CoverTab[25756]++
//line /usr/local/go/src/compress/flate/deflate.go:323
			// _ = "end of CoverTab[25756]"
//line /usr/local/go/src/compress/flate/deflate.go:323
		}
//line /usr/local/go/src/compress/flate/deflate.go:323
		// _ = "end of CoverTab[25754]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:324
	// _ = "end of CoverTab[25752]"
//line /usr/local/go/src/compress/flate/deflate.go:324
	_go_fuzz_dep_.CoverTab[25753]++
							return max
//line /usr/local/go/src/compress/flate/deflate.go:325
	// _ = "end of CoverTab[25753]"
}

// encSpeed will compress and store the currently added data,
//line /usr/local/go/src/compress/flate/deflate.go:328
// if enough has been accumulated or we at the end of the stream.
//line /usr/local/go/src/compress/flate/deflate.go:328
// Any error that occurred will be in d.err
//line /usr/local/go/src/compress/flate/deflate.go:331
func (d *compressor) encSpeed() {
//line /usr/local/go/src/compress/flate/deflate.go:331
	_go_fuzz_dep_.CoverTab[25757]++

							if d.windowEnd < maxStoreBlockSize {
//line /usr/local/go/src/compress/flate/deflate.go:333
		_go_fuzz_dep_.CoverTab[25760]++
								if !d.sync {
//line /usr/local/go/src/compress/flate/deflate.go:334
			_go_fuzz_dep_.CoverTab[25762]++
									return
//line /usr/local/go/src/compress/flate/deflate.go:335
			// _ = "end of CoverTab[25762]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:336
			_go_fuzz_dep_.CoverTab[25763]++
//line /usr/local/go/src/compress/flate/deflate.go:336
			// _ = "end of CoverTab[25763]"
//line /usr/local/go/src/compress/flate/deflate.go:336
		}
//line /usr/local/go/src/compress/flate/deflate.go:336
		// _ = "end of CoverTab[25760]"
//line /usr/local/go/src/compress/flate/deflate.go:336
		_go_fuzz_dep_.CoverTab[25761]++

//line /usr/local/go/src/compress/flate/deflate.go:339
		if d.windowEnd < 128 {
//line /usr/local/go/src/compress/flate/deflate.go:339
			_go_fuzz_dep_.CoverTab[25764]++
									switch {
			case d.windowEnd == 0:
//line /usr/local/go/src/compress/flate/deflate.go:341
				_go_fuzz_dep_.CoverTab[25766]++
										return
//line /usr/local/go/src/compress/flate/deflate.go:342
				// _ = "end of CoverTab[25766]"
			case d.windowEnd <= 16:
//line /usr/local/go/src/compress/flate/deflate.go:343
				_go_fuzz_dep_.CoverTab[25767]++
										d.err = d.writeStoredBlock(d.window[:d.windowEnd])
//line /usr/local/go/src/compress/flate/deflate.go:344
				// _ = "end of CoverTab[25767]"
			default:
//line /usr/local/go/src/compress/flate/deflate.go:345
				_go_fuzz_dep_.CoverTab[25768]++
										d.w.writeBlockHuff(false, d.window[:d.windowEnd])
										d.err = d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:347
				// _ = "end of CoverTab[25768]"
			}
//line /usr/local/go/src/compress/flate/deflate.go:348
			// _ = "end of CoverTab[25764]"
//line /usr/local/go/src/compress/flate/deflate.go:348
			_go_fuzz_dep_.CoverTab[25765]++
									d.windowEnd = 0
									d.bestSpeed.reset()
									return
//line /usr/local/go/src/compress/flate/deflate.go:351
			// _ = "end of CoverTab[25765]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:352
			_go_fuzz_dep_.CoverTab[25769]++
//line /usr/local/go/src/compress/flate/deflate.go:352
			// _ = "end of CoverTab[25769]"
//line /usr/local/go/src/compress/flate/deflate.go:352
		}
//line /usr/local/go/src/compress/flate/deflate.go:352
		// _ = "end of CoverTab[25761]"

	} else {
//line /usr/local/go/src/compress/flate/deflate.go:354
		_go_fuzz_dep_.CoverTab[25770]++
//line /usr/local/go/src/compress/flate/deflate.go:354
		// _ = "end of CoverTab[25770]"
//line /usr/local/go/src/compress/flate/deflate.go:354
	}
//line /usr/local/go/src/compress/flate/deflate.go:354
	// _ = "end of CoverTab[25757]"
//line /usr/local/go/src/compress/flate/deflate.go:354
	_go_fuzz_dep_.CoverTab[25758]++

							d.tokens = d.bestSpeed.encode(d.tokens[:0], d.window[:d.windowEnd])

//line /usr/local/go/src/compress/flate/deflate.go:359
	if len(d.tokens) > d.windowEnd-(d.windowEnd>>4) {
//line /usr/local/go/src/compress/flate/deflate.go:359
		_go_fuzz_dep_.CoverTab[25771]++
								d.w.writeBlockHuff(false, d.window[:d.windowEnd])
//line /usr/local/go/src/compress/flate/deflate.go:360
		// _ = "end of CoverTab[25771]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:361
		_go_fuzz_dep_.CoverTab[25772]++
								d.w.writeBlockDynamic(d.tokens, false, d.window[:d.windowEnd])
//line /usr/local/go/src/compress/flate/deflate.go:362
		// _ = "end of CoverTab[25772]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:363
	// _ = "end of CoverTab[25758]"
//line /usr/local/go/src/compress/flate/deflate.go:363
	_go_fuzz_dep_.CoverTab[25759]++
							d.err = d.w.err
							d.windowEnd = 0
//line /usr/local/go/src/compress/flate/deflate.go:365
	// _ = "end of CoverTab[25759]"
}

func (d *compressor) initDeflate() {
//line /usr/local/go/src/compress/flate/deflate.go:368
	_go_fuzz_dep_.CoverTab[25773]++
							d.window = make([]byte, 2*windowSize)
							d.hashOffset = 1
							d.tokens = make([]token, 0, maxFlateBlockTokens+1)
							d.length = minMatchLength - 1
							d.offset = 0
							d.byteAvailable = false
							d.index = 0
							d.chainHead = -1
							d.bulkHasher = bulkHash4
//line /usr/local/go/src/compress/flate/deflate.go:377
	// _ = "end of CoverTab[25773]"
}

func (d *compressor) deflate() {
//line /usr/local/go/src/compress/flate/deflate.go:380
	_go_fuzz_dep_.CoverTab[25774]++
							if d.windowEnd-d.index < minMatchLength+maxMatchLength && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:381
		_go_fuzz_dep_.CoverTab[25776]++
//line /usr/local/go/src/compress/flate/deflate.go:381
		return !d.sync
//line /usr/local/go/src/compress/flate/deflate.go:381
		// _ = "end of CoverTab[25776]"
//line /usr/local/go/src/compress/flate/deflate.go:381
	}() {
//line /usr/local/go/src/compress/flate/deflate.go:381
		_go_fuzz_dep_.CoverTab[25777]++
								return
//line /usr/local/go/src/compress/flate/deflate.go:382
		// _ = "end of CoverTab[25777]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:383
		_go_fuzz_dep_.CoverTab[25778]++
//line /usr/local/go/src/compress/flate/deflate.go:383
		// _ = "end of CoverTab[25778]"
//line /usr/local/go/src/compress/flate/deflate.go:383
	}
//line /usr/local/go/src/compress/flate/deflate.go:383
	// _ = "end of CoverTab[25774]"
//line /usr/local/go/src/compress/flate/deflate.go:383
	_go_fuzz_dep_.CoverTab[25775]++

							d.maxInsertIndex = d.windowEnd - (minMatchLength - 1)

Loop:
	for {
//line /usr/local/go/src/compress/flate/deflate.go:388
		_go_fuzz_dep_.CoverTab[25779]++
								if d.index > d.windowEnd {
//line /usr/local/go/src/compress/flate/deflate.go:389
			_go_fuzz_dep_.CoverTab[25785]++
									panic("index > windowEnd")
//line /usr/local/go/src/compress/flate/deflate.go:390
			// _ = "end of CoverTab[25785]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:391
			_go_fuzz_dep_.CoverTab[25786]++
//line /usr/local/go/src/compress/flate/deflate.go:391
			// _ = "end of CoverTab[25786]"
//line /usr/local/go/src/compress/flate/deflate.go:391
		}
//line /usr/local/go/src/compress/flate/deflate.go:391
		// _ = "end of CoverTab[25779]"
//line /usr/local/go/src/compress/flate/deflate.go:391
		_go_fuzz_dep_.CoverTab[25780]++
								lookahead := d.windowEnd - d.index
								if lookahead < minMatchLength+maxMatchLength {
//line /usr/local/go/src/compress/flate/deflate.go:393
			_go_fuzz_dep_.CoverTab[25787]++
									if !d.sync {
//line /usr/local/go/src/compress/flate/deflate.go:394
				_go_fuzz_dep_.CoverTab[25790]++
										break Loop
//line /usr/local/go/src/compress/flate/deflate.go:395
				// _ = "end of CoverTab[25790]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:396
				_go_fuzz_dep_.CoverTab[25791]++
//line /usr/local/go/src/compress/flate/deflate.go:396
				// _ = "end of CoverTab[25791]"
//line /usr/local/go/src/compress/flate/deflate.go:396
			}
//line /usr/local/go/src/compress/flate/deflate.go:396
			// _ = "end of CoverTab[25787]"
//line /usr/local/go/src/compress/flate/deflate.go:396
			_go_fuzz_dep_.CoverTab[25788]++
									if d.index > d.windowEnd {
//line /usr/local/go/src/compress/flate/deflate.go:397
				_go_fuzz_dep_.CoverTab[25792]++
										panic("index > windowEnd")
//line /usr/local/go/src/compress/flate/deflate.go:398
				// _ = "end of CoverTab[25792]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:399
				_go_fuzz_dep_.CoverTab[25793]++
//line /usr/local/go/src/compress/flate/deflate.go:399
				// _ = "end of CoverTab[25793]"
//line /usr/local/go/src/compress/flate/deflate.go:399
			}
//line /usr/local/go/src/compress/flate/deflate.go:399
			// _ = "end of CoverTab[25788]"
//line /usr/local/go/src/compress/flate/deflate.go:399
			_go_fuzz_dep_.CoverTab[25789]++
									if lookahead == 0 {
//line /usr/local/go/src/compress/flate/deflate.go:400
				_go_fuzz_dep_.CoverTab[25794]++

										if d.byteAvailable {
//line /usr/local/go/src/compress/flate/deflate.go:402
					_go_fuzz_dep_.CoverTab[25797]++

											d.tokens = append(d.tokens, literalToken(uint32(d.window[d.index-1])))
											d.byteAvailable = false
//line /usr/local/go/src/compress/flate/deflate.go:405
					// _ = "end of CoverTab[25797]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:406
					_go_fuzz_dep_.CoverTab[25798]++
//line /usr/local/go/src/compress/flate/deflate.go:406
					// _ = "end of CoverTab[25798]"
//line /usr/local/go/src/compress/flate/deflate.go:406
				}
//line /usr/local/go/src/compress/flate/deflate.go:406
				// _ = "end of CoverTab[25794]"
//line /usr/local/go/src/compress/flate/deflate.go:406
				_go_fuzz_dep_.CoverTab[25795]++
										if len(d.tokens) > 0 {
//line /usr/local/go/src/compress/flate/deflate.go:407
					_go_fuzz_dep_.CoverTab[25799]++
											if d.err = d.writeBlock(d.tokens, d.index); d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:408
						_go_fuzz_dep_.CoverTab[25801]++
												return
//line /usr/local/go/src/compress/flate/deflate.go:409
						// _ = "end of CoverTab[25801]"
					} else {
//line /usr/local/go/src/compress/flate/deflate.go:410
						_go_fuzz_dep_.CoverTab[25802]++
//line /usr/local/go/src/compress/flate/deflate.go:410
						// _ = "end of CoverTab[25802]"
//line /usr/local/go/src/compress/flate/deflate.go:410
					}
//line /usr/local/go/src/compress/flate/deflate.go:410
					// _ = "end of CoverTab[25799]"
//line /usr/local/go/src/compress/flate/deflate.go:410
					_go_fuzz_dep_.CoverTab[25800]++
											d.tokens = d.tokens[:0]
//line /usr/local/go/src/compress/flate/deflate.go:411
					// _ = "end of CoverTab[25800]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:412
					_go_fuzz_dep_.CoverTab[25803]++
//line /usr/local/go/src/compress/flate/deflate.go:412
					// _ = "end of CoverTab[25803]"
//line /usr/local/go/src/compress/flate/deflate.go:412
				}
//line /usr/local/go/src/compress/flate/deflate.go:412
				// _ = "end of CoverTab[25795]"
//line /usr/local/go/src/compress/flate/deflate.go:412
				_go_fuzz_dep_.CoverTab[25796]++
										break Loop
//line /usr/local/go/src/compress/flate/deflate.go:413
				// _ = "end of CoverTab[25796]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:414
				_go_fuzz_dep_.CoverTab[25804]++
//line /usr/local/go/src/compress/flate/deflate.go:414
				// _ = "end of CoverTab[25804]"
//line /usr/local/go/src/compress/flate/deflate.go:414
			}
//line /usr/local/go/src/compress/flate/deflate.go:414
			// _ = "end of CoverTab[25789]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:415
			_go_fuzz_dep_.CoverTab[25805]++
//line /usr/local/go/src/compress/flate/deflate.go:415
			// _ = "end of CoverTab[25805]"
//line /usr/local/go/src/compress/flate/deflate.go:415
		}
//line /usr/local/go/src/compress/flate/deflate.go:415
		// _ = "end of CoverTab[25780]"
//line /usr/local/go/src/compress/flate/deflate.go:415
		_go_fuzz_dep_.CoverTab[25781]++
								if d.index < d.maxInsertIndex {
//line /usr/local/go/src/compress/flate/deflate.go:416
			_go_fuzz_dep_.CoverTab[25806]++

									hash := hash4(d.window[d.index : d.index+minMatchLength])
									hh := &d.hashHead[hash&hashMask]
									d.chainHead = int(*hh)
									d.hashPrev[d.index&windowMask] = uint32(d.chainHead)
									*hh = uint32(d.index + d.hashOffset)
//line /usr/local/go/src/compress/flate/deflate.go:422
			// _ = "end of CoverTab[25806]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:423
			_go_fuzz_dep_.CoverTab[25807]++
//line /usr/local/go/src/compress/flate/deflate.go:423
			// _ = "end of CoverTab[25807]"
//line /usr/local/go/src/compress/flate/deflate.go:423
		}
//line /usr/local/go/src/compress/flate/deflate.go:423
		// _ = "end of CoverTab[25781]"
//line /usr/local/go/src/compress/flate/deflate.go:423
		_go_fuzz_dep_.CoverTab[25782]++
								prevLength := d.length
								prevOffset := d.offset
								d.length = minMatchLength - 1
								d.offset = 0
								minIndex := d.index - windowSize
								if minIndex < 0 {
//line /usr/local/go/src/compress/flate/deflate.go:429
			_go_fuzz_dep_.CoverTab[25808]++
									minIndex = 0
//line /usr/local/go/src/compress/flate/deflate.go:430
			// _ = "end of CoverTab[25808]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:431
			_go_fuzz_dep_.CoverTab[25809]++
//line /usr/local/go/src/compress/flate/deflate.go:431
			// _ = "end of CoverTab[25809]"
//line /usr/local/go/src/compress/flate/deflate.go:431
		}
//line /usr/local/go/src/compress/flate/deflate.go:431
		// _ = "end of CoverTab[25782]"
//line /usr/local/go/src/compress/flate/deflate.go:431
		_go_fuzz_dep_.CoverTab[25783]++

								if d.chainHead-d.hashOffset >= minIndex && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:433
			_go_fuzz_dep_.CoverTab[25810]++
//line /usr/local/go/src/compress/flate/deflate.go:433
			return (d.fastSkipHashing != skipNever && func() bool {
										_go_fuzz_dep_.CoverTab[25811]++
//line /usr/local/go/src/compress/flate/deflate.go:434
				return lookahead > minMatchLength-1
//line /usr/local/go/src/compress/flate/deflate.go:434
				// _ = "end of CoverTab[25811]"
//line /usr/local/go/src/compress/flate/deflate.go:434
			}() || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:434
				_go_fuzz_dep_.CoverTab[25812]++
//line /usr/local/go/src/compress/flate/deflate.go:434
				return d.fastSkipHashing == skipNever && func() bool {
											_go_fuzz_dep_.CoverTab[25813]++
//line /usr/local/go/src/compress/flate/deflate.go:435
					return lookahead > prevLength
//line /usr/local/go/src/compress/flate/deflate.go:435
					// _ = "end of CoverTab[25813]"
//line /usr/local/go/src/compress/flate/deflate.go:435
				}() && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:435
					_go_fuzz_dep_.CoverTab[25814]++
//line /usr/local/go/src/compress/flate/deflate.go:435
					return prevLength < d.lazy
//line /usr/local/go/src/compress/flate/deflate.go:435
					// _ = "end of CoverTab[25814]"
//line /usr/local/go/src/compress/flate/deflate.go:435
				}()
//line /usr/local/go/src/compress/flate/deflate.go:435
				// _ = "end of CoverTab[25812]"
//line /usr/local/go/src/compress/flate/deflate.go:435
			}())
//line /usr/local/go/src/compress/flate/deflate.go:435
			// _ = "end of CoverTab[25810]"
//line /usr/local/go/src/compress/flate/deflate.go:435
		}() {
//line /usr/local/go/src/compress/flate/deflate.go:435
			_go_fuzz_dep_.CoverTab[25815]++
									if newLength, newOffset, ok := d.findMatch(d.index, d.chainHead-d.hashOffset, minMatchLength-1, lookahead); ok {
//line /usr/local/go/src/compress/flate/deflate.go:436
				_go_fuzz_dep_.CoverTab[25816]++
										d.length = newLength
										d.offset = newOffset
//line /usr/local/go/src/compress/flate/deflate.go:438
				// _ = "end of CoverTab[25816]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:439
				_go_fuzz_dep_.CoverTab[25817]++
//line /usr/local/go/src/compress/flate/deflate.go:439
				// _ = "end of CoverTab[25817]"
//line /usr/local/go/src/compress/flate/deflate.go:439
			}
//line /usr/local/go/src/compress/flate/deflate.go:439
			// _ = "end of CoverTab[25815]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:440
			_go_fuzz_dep_.CoverTab[25818]++
//line /usr/local/go/src/compress/flate/deflate.go:440
			// _ = "end of CoverTab[25818]"
//line /usr/local/go/src/compress/flate/deflate.go:440
		}
//line /usr/local/go/src/compress/flate/deflate.go:440
		// _ = "end of CoverTab[25783]"
//line /usr/local/go/src/compress/flate/deflate.go:440
		_go_fuzz_dep_.CoverTab[25784]++
								if d.fastSkipHashing != skipNever && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:441
			_go_fuzz_dep_.CoverTab[25819]++
//line /usr/local/go/src/compress/flate/deflate.go:441
			return d.length >= minMatchLength
//line /usr/local/go/src/compress/flate/deflate.go:441
			// _ = "end of CoverTab[25819]"
//line /usr/local/go/src/compress/flate/deflate.go:441
		}() || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:441
			_go_fuzz_dep_.CoverTab[25820]++
//line /usr/local/go/src/compress/flate/deflate.go:441
			return d.fastSkipHashing == skipNever && func() bool {
										_go_fuzz_dep_.CoverTab[25821]++
//line /usr/local/go/src/compress/flate/deflate.go:442
				return prevLength >= minMatchLength
//line /usr/local/go/src/compress/flate/deflate.go:442
				// _ = "end of CoverTab[25821]"
//line /usr/local/go/src/compress/flate/deflate.go:442
			}() && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:442
				_go_fuzz_dep_.CoverTab[25822]++
//line /usr/local/go/src/compress/flate/deflate.go:442
				return d.length <= prevLength
//line /usr/local/go/src/compress/flate/deflate.go:442
				// _ = "end of CoverTab[25822]"
//line /usr/local/go/src/compress/flate/deflate.go:442
			}()
//line /usr/local/go/src/compress/flate/deflate.go:442
			// _ = "end of CoverTab[25820]"
//line /usr/local/go/src/compress/flate/deflate.go:442
		}() {
//line /usr/local/go/src/compress/flate/deflate.go:442
			_go_fuzz_dep_.CoverTab[25823]++

//line /usr/local/go/src/compress/flate/deflate.go:445
			if d.fastSkipHashing != skipNever {
//line /usr/local/go/src/compress/flate/deflate.go:445
				_go_fuzz_dep_.CoverTab[25826]++
										d.tokens = append(d.tokens, matchToken(uint32(d.length-baseMatchLength), uint32(d.offset-baseMatchOffset)))
//line /usr/local/go/src/compress/flate/deflate.go:446
				// _ = "end of CoverTab[25826]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:447
				_go_fuzz_dep_.CoverTab[25827]++
										d.tokens = append(d.tokens, matchToken(uint32(prevLength-baseMatchLength), uint32(prevOffset-baseMatchOffset)))
//line /usr/local/go/src/compress/flate/deflate.go:448
				// _ = "end of CoverTab[25827]"
			}
//line /usr/local/go/src/compress/flate/deflate.go:449
			// _ = "end of CoverTab[25823]"
//line /usr/local/go/src/compress/flate/deflate.go:449
			_go_fuzz_dep_.CoverTab[25824]++

//line /usr/local/go/src/compress/flate/deflate.go:454
			if d.length <= d.fastSkipHashing {
//line /usr/local/go/src/compress/flate/deflate.go:454
				_go_fuzz_dep_.CoverTab[25828]++
										var newIndex int
										if d.fastSkipHashing != skipNever {
//line /usr/local/go/src/compress/flate/deflate.go:456
					_go_fuzz_dep_.CoverTab[25831]++
											newIndex = d.index + d.length
//line /usr/local/go/src/compress/flate/deflate.go:457
					// _ = "end of CoverTab[25831]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:458
					_go_fuzz_dep_.CoverTab[25832]++
											newIndex = d.index + prevLength - 1
//line /usr/local/go/src/compress/flate/deflate.go:459
					// _ = "end of CoverTab[25832]"
				}
//line /usr/local/go/src/compress/flate/deflate.go:460
				// _ = "end of CoverTab[25828]"
//line /usr/local/go/src/compress/flate/deflate.go:460
				_go_fuzz_dep_.CoverTab[25829]++
										index := d.index
										for index++; index < newIndex; index++ {
//line /usr/local/go/src/compress/flate/deflate.go:462
					_go_fuzz_dep_.CoverTab[25833]++
											if index < d.maxInsertIndex {
//line /usr/local/go/src/compress/flate/deflate.go:463
						_go_fuzz_dep_.CoverTab[25834]++
												hash := hash4(d.window[index : index+minMatchLength])

//line /usr/local/go/src/compress/flate/deflate.go:467
						hh := &d.hashHead[hash&hashMask]
												d.hashPrev[index&windowMask] = *hh

												*hh = uint32(index + d.hashOffset)
//line /usr/local/go/src/compress/flate/deflate.go:470
						// _ = "end of CoverTab[25834]"
					} else {
//line /usr/local/go/src/compress/flate/deflate.go:471
						_go_fuzz_dep_.CoverTab[25835]++
//line /usr/local/go/src/compress/flate/deflate.go:471
						// _ = "end of CoverTab[25835]"
//line /usr/local/go/src/compress/flate/deflate.go:471
					}
//line /usr/local/go/src/compress/flate/deflate.go:471
					// _ = "end of CoverTab[25833]"
				}
//line /usr/local/go/src/compress/flate/deflate.go:472
				// _ = "end of CoverTab[25829]"
//line /usr/local/go/src/compress/flate/deflate.go:472
				_go_fuzz_dep_.CoverTab[25830]++
										d.index = index

										if d.fastSkipHashing == skipNever {
//line /usr/local/go/src/compress/flate/deflate.go:475
					_go_fuzz_dep_.CoverTab[25836]++
											d.byteAvailable = false
											d.length = minMatchLength - 1
//line /usr/local/go/src/compress/flate/deflate.go:477
					// _ = "end of CoverTab[25836]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:478
					_go_fuzz_dep_.CoverTab[25837]++
//line /usr/local/go/src/compress/flate/deflate.go:478
					// _ = "end of CoverTab[25837]"
//line /usr/local/go/src/compress/flate/deflate.go:478
				}
//line /usr/local/go/src/compress/flate/deflate.go:478
				// _ = "end of CoverTab[25830]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:479
				_go_fuzz_dep_.CoverTab[25838]++

//line /usr/local/go/src/compress/flate/deflate.go:482
				d.index += d.length
//line /usr/local/go/src/compress/flate/deflate.go:482
				// _ = "end of CoverTab[25838]"
			}
//line /usr/local/go/src/compress/flate/deflate.go:483
			// _ = "end of CoverTab[25824]"
//line /usr/local/go/src/compress/flate/deflate.go:483
			_go_fuzz_dep_.CoverTab[25825]++
									if len(d.tokens) == maxFlateBlockTokens {
//line /usr/local/go/src/compress/flate/deflate.go:484
				_go_fuzz_dep_.CoverTab[25839]++

										if d.err = d.writeBlock(d.tokens, d.index); d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:486
					_go_fuzz_dep_.CoverTab[25841]++
											return
//line /usr/local/go/src/compress/flate/deflate.go:487
					// _ = "end of CoverTab[25841]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:488
					_go_fuzz_dep_.CoverTab[25842]++
//line /usr/local/go/src/compress/flate/deflate.go:488
					// _ = "end of CoverTab[25842]"
//line /usr/local/go/src/compress/flate/deflate.go:488
				}
//line /usr/local/go/src/compress/flate/deflate.go:488
				// _ = "end of CoverTab[25839]"
//line /usr/local/go/src/compress/flate/deflate.go:488
				_go_fuzz_dep_.CoverTab[25840]++
										d.tokens = d.tokens[:0]
//line /usr/local/go/src/compress/flate/deflate.go:489
				// _ = "end of CoverTab[25840]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:490
				_go_fuzz_dep_.CoverTab[25843]++
//line /usr/local/go/src/compress/flate/deflate.go:490
				// _ = "end of CoverTab[25843]"
//line /usr/local/go/src/compress/flate/deflate.go:490
			}
//line /usr/local/go/src/compress/flate/deflate.go:490
			// _ = "end of CoverTab[25825]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:491
			_go_fuzz_dep_.CoverTab[25844]++
									if d.fastSkipHashing != skipNever || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:492
				_go_fuzz_dep_.CoverTab[25846]++
//line /usr/local/go/src/compress/flate/deflate.go:492
				return d.byteAvailable
//line /usr/local/go/src/compress/flate/deflate.go:492
				// _ = "end of CoverTab[25846]"
//line /usr/local/go/src/compress/flate/deflate.go:492
			}() {
//line /usr/local/go/src/compress/flate/deflate.go:492
				_go_fuzz_dep_.CoverTab[25847]++
										i := d.index - 1
										if d.fastSkipHashing != skipNever {
//line /usr/local/go/src/compress/flate/deflate.go:494
					_go_fuzz_dep_.CoverTab[25849]++
											i = d.index
//line /usr/local/go/src/compress/flate/deflate.go:495
					// _ = "end of CoverTab[25849]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:496
					_go_fuzz_dep_.CoverTab[25850]++
//line /usr/local/go/src/compress/flate/deflate.go:496
					// _ = "end of CoverTab[25850]"
//line /usr/local/go/src/compress/flate/deflate.go:496
				}
//line /usr/local/go/src/compress/flate/deflate.go:496
				// _ = "end of CoverTab[25847]"
//line /usr/local/go/src/compress/flate/deflate.go:496
				_go_fuzz_dep_.CoverTab[25848]++
										d.tokens = append(d.tokens, literalToken(uint32(d.window[i])))
										if len(d.tokens) == maxFlateBlockTokens {
//line /usr/local/go/src/compress/flate/deflate.go:498
					_go_fuzz_dep_.CoverTab[25851]++
											if d.err = d.writeBlock(d.tokens, i+1); d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:499
						_go_fuzz_dep_.CoverTab[25853]++
												return
//line /usr/local/go/src/compress/flate/deflate.go:500
						// _ = "end of CoverTab[25853]"
					} else {
//line /usr/local/go/src/compress/flate/deflate.go:501
						_go_fuzz_dep_.CoverTab[25854]++
//line /usr/local/go/src/compress/flate/deflate.go:501
						// _ = "end of CoverTab[25854]"
//line /usr/local/go/src/compress/flate/deflate.go:501
					}
//line /usr/local/go/src/compress/flate/deflate.go:501
					// _ = "end of CoverTab[25851]"
//line /usr/local/go/src/compress/flate/deflate.go:501
					_go_fuzz_dep_.CoverTab[25852]++
											d.tokens = d.tokens[:0]
//line /usr/local/go/src/compress/flate/deflate.go:502
					// _ = "end of CoverTab[25852]"
				} else {
//line /usr/local/go/src/compress/flate/deflate.go:503
					_go_fuzz_dep_.CoverTab[25855]++
//line /usr/local/go/src/compress/flate/deflate.go:503
					// _ = "end of CoverTab[25855]"
//line /usr/local/go/src/compress/flate/deflate.go:503
				}
//line /usr/local/go/src/compress/flate/deflate.go:503
				// _ = "end of CoverTab[25848]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:504
				_go_fuzz_dep_.CoverTab[25856]++
//line /usr/local/go/src/compress/flate/deflate.go:504
				// _ = "end of CoverTab[25856]"
//line /usr/local/go/src/compress/flate/deflate.go:504
			}
//line /usr/local/go/src/compress/flate/deflate.go:504
			// _ = "end of CoverTab[25844]"
//line /usr/local/go/src/compress/flate/deflate.go:504
			_go_fuzz_dep_.CoverTab[25845]++
									d.index++
									if d.fastSkipHashing == skipNever {
//line /usr/local/go/src/compress/flate/deflate.go:506
				_go_fuzz_dep_.CoverTab[25857]++
										d.byteAvailable = true
//line /usr/local/go/src/compress/flate/deflate.go:507
				// _ = "end of CoverTab[25857]"
			} else {
//line /usr/local/go/src/compress/flate/deflate.go:508
				_go_fuzz_dep_.CoverTab[25858]++
//line /usr/local/go/src/compress/flate/deflate.go:508
				// _ = "end of CoverTab[25858]"
//line /usr/local/go/src/compress/flate/deflate.go:508
			}
//line /usr/local/go/src/compress/flate/deflate.go:508
			// _ = "end of CoverTab[25845]"
		}
//line /usr/local/go/src/compress/flate/deflate.go:509
		// _ = "end of CoverTab[25784]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:510
	// _ = "end of CoverTab[25775]"
}

func (d *compressor) fillStore(b []byte) int {
//line /usr/local/go/src/compress/flate/deflate.go:513
	_go_fuzz_dep_.CoverTab[25859]++
							n := copy(d.window[d.windowEnd:], b)
							d.windowEnd += n
							return n
//line /usr/local/go/src/compress/flate/deflate.go:516
	// _ = "end of CoverTab[25859]"
}

func (d *compressor) store() {
//line /usr/local/go/src/compress/flate/deflate.go:519
	_go_fuzz_dep_.CoverTab[25860]++
							if d.windowEnd > 0 && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:520
		_go_fuzz_dep_.CoverTab[25861]++
//line /usr/local/go/src/compress/flate/deflate.go:520
		return (d.windowEnd == maxStoreBlockSize || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:520
			_go_fuzz_dep_.CoverTab[25862]++
//line /usr/local/go/src/compress/flate/deflate.go:520
			return d.sync
//line /usr/local/go/src/compress/flate/deflate.go:520
			// _ = "end of CoverTab[25862]"
//line /usr/local/go/src/compress/flate/deflate.go:520
		}())
//line /usr/local/go/src/compress/flate/deflate.go:520
		// _ = "end of CoverTab[25861]"
//line /usr/local/go/src/compress/flate/deflate.go:520
	}() {
//line /usr/local/go/src/compress/flate/deflate.go:520
		_go_fuzz_dep_.CoverTab[25863]++
								d.err = d.writeStoredBlock(d.window[:d.windowEnd])
								d.windowEnd = 0
//line /usr/local/go/src/compress/flate/deflate.go:522
		// _ = "end of CoverTab[25863]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:523
		_go_fuzz_dep_.CoverTab[25864]++
//line /usr/local/go/src/compress/flate/deflate.go:523
		// _ = "end of CoverTab[25864]"
//line /usr/local/go/src/compress/flate/deflate.go:523
	}
//line /usr/local/go/src/compress/flate/deflate.go:523
	// _ = "end of CoverTab[25860]"
}

// storeHuff compresses and stores the currently added data
//line /usr/local/go/src/compress/flate/deflate.go:526
// when the d.window is full or we are at the end of the stream.
//line /usr/local/go/src/compress/flate/deflate.go:526
// Any error that occurred will be in d.err
//line /usr/local/go/src/compress/flate/deflate.go:529
func (d *compressor) storeHuff() {
//line /usr/local/go/src/compress/flate/deflate.go:529
	_go_fuzz_dep_.CoverTab[25865]++
							if d.windowEnd < len(d.window) && func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:530
		_go_fuzz_dep_.CoverTab[25867]++
//line /usr/local/go/src/compress/flate/deflate.go:530
		return !d.sync
//line /usr/local/go/src/compress/flate/deflate.go:530
		// _ = "end of CoverTab[25867]"
//line /usr/local/go/src/compress/flate/deflate.go:530
	}() || func() bool {
//line /usr/local/go/src/compress/flate/deflate.go:530
		_go_fuzz_dep_.CoverTab[25868]++
//line /usr/local/go/src/compress/flate/deflate.go:530
		return d.windowEnd == 0
//line /usr/local/go/src/compress/flate/deflate.go:530
		// _ = "end of CoverTab[25868]"
//line /usr/local/go/src/compress/flate/deflate.go:530
	}() {
//line /usr/local/go/src/compress/flate/deflate.go:530
		_go_fuzz_dep_.CoverTab[25869]++
								return
//line /usr/local/go/src/compress/flate/deflate.go:531
		// _ = "end of CoverTab[25869]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:532
		_go_fuzz_dep_.CoverTab[25870]++
//line /usr/local/go/src/compress/flate/deflate.go:532
		// _ = "end of CoverTab[25870]"
//line /usr/local/go/src/compress/flate/deflate.go:532
	}
//line /usr/local/go/src/compress/flate/deflate.go:532
	// _ = "end of CoverTab[25865]"
//line /usr/local/go/src/compress/flate/deflate.go:532
	_go_fuzz_dep_.CoverTab[25866]++
							d.w.writeBlockHuff(false, d.window[:d.windowEnd])
							d.err = d.w.err
							d.windowEnd = 0
//line /usr/local/go/src/compress/flate/deflate.go:535
	// _ = "end of CoverTab[25866]"
}

func (d *compressor) write(b []byte) (n int, err error) {
//line /usr/local/go/src/compress/flate/deflate.go:538
	_go_fuzz_dep_.CoverTab[25871]++
							if d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:539
		_go_fuzz_dep_.CoverTab[25874]++
								return 0, d.err
//line /usr/local/go/src/compress/flate/deflate.go:540
		// _ = "end of CoverTab[25874]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:541
		_go_fuzz_dep_.CoverTab[25875]++
//line /usr/local/go/src/compress/flate/deflate.go:541
		// _ = "end of CoverTab[25875]"
//line /usr/local/go/src/compress/flate/deflate.go:541
	}
//line /usr/local/go/src/compress/flate/deflate.go:541
	// _ = "end of CoverTab[25871]"
//line /usr/local/go/src/compress/flate/deflate.go:541
	_go_fuzz_dep_.CoverTab[25872]++
							n = len(b)
							for len(b) > 0 {
//line /usr/local/go/src/compress/flate/deflate.go:543
		_go_fuzz_dep_.CoverTab[25876]++
								d.step(d)
								b = b[d.fill(d, b):]
								if d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:546
			_go_fuzz_dep_.CoverTab[25877]++
									return 0, d.err
//line /usr/local/go/src/compress/flate/deflate.go:547
			// _ = "end of CoverTab[25877]"
		} else {
//line /usr/local/go/src/compress/flate/deflate.go:548
			_go_fuzz_dep_.CoverTab[25878]++
//line /usr/local/go/src/compress/flate/deflate.go:548
			// _ = "end of CoverTab[25878]"
//line /usr/local/go/src/compress/flate/deflate.go:548
		}
//line /usr/local/go/src/compress/flate/deflate.go:548
		// _ = "end of CoverTab[25876]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:549
	// _ = "end of CoverTab[25872]"
//line /usr/local/go/src/compress/flate/deflate.go:549
	_go_fuzz_dep_.CoverTab[25873]++
							return n, nil
//line /usr/local/go/src/compress/flate/deflate.go:550
	// _ = "end of CoverTab[25873]"
}

func (d *compressor) syncFlush() error {
//line /usr/local/go/src/compress/flate/deflate.go:553
	_go_fuzz_dep_.CoverTab[25879]++
							if d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:554
		_go_fuzz_dep_.CoverTab[25882]++
								return d.err
//line /usr/local/go/src/compress/flate/deflate.go:555
		// _ = "end of CoverTab[25882]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:556
		_go_fuzz_dep_.CoverTab[25883]++
//line /usr/local/go/src/compress/flate/deflate.go:556
		// _ = "end of CoverTab[25883]"
//line /usr/local/go/src/compress/flate/deflate.go:556
	}
//line /usr/local/go/src/compress/flate/deflate.go:556
	// _ = "end of CoverTab[25879]"
//line /usr/local/go/src/compress/flate/deflate.go:556
	_go_fuzz_dep_.CoverTab[25880]++
							d.sync = true
							d.step(d)
							if d.err == nil {
//line /usr/local/go/src/compress/flate/deflate.go:559
		_go_fuzz_dep_.CoverTab[25884]++
								d.w.writeStoredHeader(0, false)
								d.w.flush()
								d.err = d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:562
		// _ = "end of CoverTab[25884]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:563
		_go_fuzz_dep_.CoverTab[25885]++
//line /usr/local/go/src/compress/flate/deflate.go:563
		// _ = "end of CoverTab[25885]"
//line /usr/local/go/src/compress/flate/deflate.go:563
	}
//line /usr/local/go/src/compress/flate/deflate.go:563
	// _ = "end of CoverTab[25880]"
//line /usr/local/go/src/compress/flate/deflate.go:563
	_go_fuzz_dep_.CoverTab[25881]++
							d.sync = false
							return d.err
//line /usr/local/go/src/compress/flate/deflate.go:565
	// _ = "end of CoverTab[25881]"
}

func (d *compressor) init(w io.Writer, level int) (err error) {
	d.w = newHuffmanBitWriter(w)

	switch {
	case level == NoCompression:
		d.window = make([]byte, maxStoreBlockSize)
		d.fill = (*compressor).fillStore
		d.step = (*compressor).store
	case level == HuffmanOnly:
		d.window = make([]byte, maxStoreBlockSize)
		d.fill = (*compressor).fillStore
		d.step = (*compressor).storeHuff
	case level == BestSpeed:
		d.compressionLevel = levels[level]
		d.window = make([]byte, maxStoreBlockSize)
		d.fill = (*compressor).fillStore
		d.step = (*compressor).encSpeed
		d.bestSpeed = newDeflateFast()
		d.tokens = make([]token, maxStoreBlockSize)
	case level == DefaultCompression:
		level = 6
		fallthrough
	case 2 <= level && level <= 9:
		d.compressionLevel = levels[level]
		d.initDeflate()
		d.fill = (*compressor).fillDeflate
		d.step = (*compressor).deflate
	default:
		return fmt.Errorf("flate: invalid compression level %d: want value in range [-2, 9]", level)
	}
	return nil
}

func (d *compressor) reset(w io.Writer) {
//line /usr/local/go/src/compress/flate/deflate.go:601
	_go_fuzz_dep_.CoverTab[25886]++
							d.w.reset(w)
							d.sync = false
							d.err = nil
							switch d.compressionLevel.level {
	case NoCompression:
//line /usr/local/go/src/compress/flate/deflate.go:606
		_go_fuzz_dep_.CoverTab[25887]++
								d.windowEnd = 0
//line /usr/local/go/src/compress/flate/deflate.go:607
		// _ = "end of CoverTab[25887]"
	case BestSpeed:
//line /usr/local/go/src/compress/flate/deflate.go:608
		_go_fuzz_dep_.CoverTab[25888]++
								d.windowEnd = 0
								d.tokens = d.tokens[:0]
								d.bestSpeed.reset()
//line /usr/local/go/src/compress/flate/deflate.go:611
		// _ = "end of CoverTab[25888]"
	default:
//line /usr/local/go/src/compress/flate/deflate.go:612
		_go_fuzz_dep_.CoverTab[25889]++
								d.chainHead = -1
								for i := range d.hashHead {
//line /usr/local/go/src/compress/flate/deflate.go:614
			_go_fuzz_dep_.CoverTab[25892]++
									d.hashHead[i] = 0
//line /usr/local/go/src/compress/flate/deflate.go:615
			// _ = "end of CoverTab[25892]"
		}
//line /usr/local/go/src/compress/flate/deflate.go:616
		// _ = "end of CoverTab[25889]"
//line /usr/local/go/src/compress/flate/deflate.go:616
		_go_fuzz_dep_.CoverTab[25890]++
								for i := range d.hashPrev {
//line /usr/local/go/src/compress/flate/deflate.go:617
			_go_fuzz_dep_.CoverTab[25893]++
									d.hashPrev[i] = 0
//line /usr/local/go/src/compress/flate/deflate.go:618
			// _ = "end of CoverTab[25893]"
		}
//line /usr/local/go/src/compress/flate/deflate.go:619
		// _ = "end of CoverTab[25890]"
//line /usr/local/go/src/compress/flate/deflate.go:619
		_go_fuzz_dep_.CoverTab[25891]++
								d.hashOffset = 1
								d.index, d.windowEnd = 0, 0
								d.blockStart, d.byteAvailable = 0, false
								d.tokens = d.tokens[:0]
								d.length = minMatchLength - 1
								d.offset = 0
								d.maxInsertIndex = 0
//line /usr/local/go/src/compress/flate/deflate.go:626
		// _ = "end of CoverTab[25891]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:627
	// _ = "end of CoverTab[25886]"
}

func (d *compressor) close() error {
//line /usr/local/go/src/compress/flate/deflate.go:630
	_go_fuzz_dep_.CoverTab[25894]++
							if d.err == errWriterClosed {
//line /usr/local/go/src/compress/flate/deflate.go:631
		_go_fuzz_dep_.CoverTab[25900]++
								return nil
//line /usr/local/go/src/compress/flate/deflate.go:632
		// _ = "end of CoverTab[25900]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:633
		_go_fuzz_dep_.CoverTab[25901]++
//line /usr/local/go/src/compress/flate/deflate.go:633
		// _ = "end of CoverTab[25901]"
//line /usr/local/go/src/compress/flate/deflate.go:633
	}
//line /usr/local/go/src/compress/flate/deflate.go:633
	// _ = "end of CoverTab[25894]"
//line /usr/local/go/src/compress/flate/deflate.go:633
	_go_fuzz_dep_.CoverTab[25895]++
							if d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:634
		_go_fuzz_dep_.CoverTab[25902]++
								return d.err
//line /usr/local/go/src/compress/flate/deflate.go:635
		// _ = "end of CoverTab[25902]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:636
		_go_fuzz_dep_.CoverTab[25903]++
//line /usr/local/go/src/compress/flate/deflate.go:636
		// _ = "end of CoverTab[25903]"
//line /usr/local/go/src/compress/flate/deflate.go:636
	}
//line /usr/local/go/src/compress/flate/deflate.go:636
	// _ = "end of CoverTab[25895]"
//line /usr/local/go/src/compress/flate/deflate.go:636
	_go_fuzz_dep_.CoverTab[25896]++
							d.sync = true
							d.step(d)
							if d.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:639
		_go_fuzz_dep_.CoverTab[25904]++
								return d.err
//line /usr/local/go/src/compress/flate/deflate.go:640
		// _ = "end of CoverTab[25904]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:641
		_go_fuzz_dep_.CoverTab[25905]++
//line /usr/local/go/src/compress/flate/deflate.go:641
		// _ = "end of CoverTab[25905]"
//line /usr/local/go/src/compress/flate/deflate.go:641
	}
//line /usr/local/go/src/compress/flate/deflate.go:641
	// _ = "end of CoverTab[25896]"
//line /usr/local/go/src/compress/flate/deflate.go:641
	_go_fuzz_dep_.CoverTab[25897]++
							if d.w.writeStoredHeader(0, true); d.w.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:642
		_go_fuzz_dep_.CoverTab[25906]++
								return d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:643
		// _ = "end of CoverTab[25906]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:644
		_go_fuzz_dep_.CoverTab[25907]++
//line /usr/local/go/src/compress/flate/deflate.go:644
		// _ = "end of CoverTab[25907]"
//line /usr/local/go/src/compress/flate/deflate.go:644
	}
//line /usr/local/go/src/compress/flate/deflate.go:644
	// _ = "end of CoverTab[25897]"
//line /usr/local/go/src/compress/flate/deflate.go:644
	_go_fuzz_dep_.CoverTab[25898]++
							d.w.flush()
							if d.w.err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:646
		_go_fuzz_dep_.CoverTab[25908]++
								return d.w.err
//line /usr/local/go/src/compress/flate/deflate.go:647
		// _ = "end of CoverTab[25908]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:648
		_go_fuzz_dep_.CoverTab[25909]++
//line /usr/local/go/src/compress/flate/deflate.go:648
		// _ = "end of CoverTab[25909]"
//line /usr/local/go/src/compress/flate/deflate.go:648
	}
//line /usr/local/go/src/compress/flate/deflate.go:648
	// _ = "end of CoverTab[25898]"
//line /usr/local/go/src/compress/flate/deflate.go:648
	_go_fuzz_dep_.CoverTab[25899]++
							d.err = errWriterClosed
							return nil
//line /usr/local/go/src/compress/flate/deflate.go:650
	// _ = "end of CoverTab[25899]"
}

// NewWriter returns a new Writer compressing data at the given level.
//line /usr/local/go/src/compress/flate/deflate.go:653
// Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression);
//line /usr/local/go/src/compress/flate/deflate.go:653
// higher levels typically run slower but compress more. Level 0
//line /usr/local/go/src/compress/flate/deflate.go:653
// (NoCompression) does not attempt any compression; it only adds the
//line /usr/local/go/src/compress/flate/deflate.go:653
// necessary DEFLATE framing.
//line /usr/local/go/src/compress/flate/deflate.go:653
// Level -1 (DefaultCompression) uses the default compression level.
//line /usr/local/go/src/compress/flate/deflate.go:653
// Level -2 (HuffmanOnly) will use Huffman compression only, giving
//line /usr/local/go/src/compress/flate/deflate.go:653
// a very fast compression for all types of input, but sacrificing considerable
//line /usr/local/go/src/compress/flate/deflate.go:653
// compression efficiency.
//line /usr/local/go/src/compress/flate/deflate.go:653
//
//line /usr/local/go/src/compress/flate/deflate.go:653
// If level is in the range [-2, 9] then the error returned will be nil.
//line /usr/local/go/src/compress/flate/deflate.go:653
// Otherwise the error returned will be non-nil.
//line /usr/local/go/src/compress/flate/deflate.go:665
func NewWriter(w io.Writer, level int) (*Writer, error) {
//line /usr/local/go/src/compress/flate/deflate.go:665
	_go_fuzz_dep_.CoverTab[25910]++
							var dw Writer
							if err := dw.d.init(w, level); err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:667
		_go_fuzz_dep_.CoverTab[25912]++
								return nil, err
//line /usr/local/go/src/compress/flate/deflate.go:668
		// _ = "end of CoverTab[25912]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:669
		_go_fuzz_dep_.CoverTab[25913]++
//line /usr/local/go/src/compress/flate/deflate.go:669
		// _ = "end of CoverTab[25913]"
//line /usr/local/go/src/compress/flate/deflate.go:669
	}
//line /usr/local/go/src/compress/flate/deflate.go:669
	// _ = "end of CoverTab[25910]"
//line /usr/local/go/src/compress/flate/deflate.go:669
	_go_fuzz_dep_.CoverTab[25911]++
							return &dw, nil
//line /usr/local/go/src/compress/flate/deflate.go:670
	// _ = "end of CoverTab[25911]"
}

// NewWriterDict is like NewWriter but initializes the new
//line /usr/local/go/src/compress/flate/deflate.go:673
// Writer with a preset dictionary. The returned Writer behaves
//line /usr/local/go/src/compress/flate/deflate.go:673
// as if the dictionary had been written to it without producing
//line /usr/local/go/src/compress/flate/deflate.go:673
// any compressed output. The compressed data written to w
//line /usr/local/go/src/compress/flate/deflate.go:673
// can only be decompressed by a Reader initialized with the
//line /usr/local/go/src/compress/flate/deflate.go:673
// same dictionary.
//line /usr/local/go/src/compress/flate/deflate.go:679
func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error) {
//line /usr/local/go/src/compress/flate/deflate.go:679
	_go_fuzz_dep_.CoverTab[25914]++
							dw := &dictWriter{w}
							zw, err := NewWriter(dw, level)
							if err != nil {
//line /usr/local/go/src/compress/flate/deflate.go:682
		_go_fuzz_dep_.CoverTab[25916]++
								return nil, err
//line /usr/local/go/src/compress/flate/deflate.go:683
		// _ = "end of CoverTab[25916]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:684
		_go_fuzz_dep_.CoverTab[25917]++
//line /usr/local/go/src/compress/flate/deflate.go:684
		// _ = "end of CoverTab[25917]"
//line /usr/local/go/src/compress/flate/deflate.go:684
	}
//line /usr/local/go/src/compress/flate/deflate.go:684
	// _ = "end of CoverTab[25914]"
//line /usr/local/go/src/compress/flate/deflate.go:684
	_go_fuzz_dep_.CoverTab[25915]++
							zw.d.fillWindow(dict)
							zw.dict = append(zw.dict, dict...)
							return zw, err
//line /usr/local/go/src/compress/flate/deflate.go:687
	// _ = "end of CoverTab[25915]"
}

type dictWriter struct {
	w io.Writer
}

func (w *dictWriter) Write(b []byte) (n int, err error) {
//line /usr/local/go/src/compress/flate/deflate.go:694
	_go_fuzz_dep_.CoverTab[25918]++
							return w.w.Write(b)
//line /usr/local/go/src/compress/flate/deflate.go:695
	// _ = "end of CoverTab[25918]"
}

var errWriterClosed = errors.New("flate: closed writer")

// A Writer takes data written to it and writes the compressed
//line /usr/local/go/src/compress/flate/deflate.go:700
// form of that data to an underlying writer (see NewWriter).
//line /usr/local/go/src/compress/flate/deflate.go:702
type Writer struct {
	d	compressor
	dict	[]byte
}

// Write writes data to w, which will eventually write the
//line /usr/local/go/src/compress/flate/deflate.go:707
// compressed form of data to its underlying writer.
//line /usr/local/go/src/compress/flate/deflate.go:709
func (w *Writer) Write(data []byte) (n int, err error) {
//line /usr/local/go/src/compress/flate/deflate.go:709
	_go_fuzz_dep_.CoverTab[25919]++
							return w.d.write(data)
//line /usr/local/go/src/compress/flate/deflate.go:710
	// _ = "end of CoverTab[25919]"
}

// Flush flushes any pending data to the underlying writer.
//line /usr/local/go/src/compress/flate/deflate.go:713
// It is useful mainly in compressed network protocols, to ensure that
//line /usr/local/go/src/compress/flate/deflate.go:713
// a remote reader has enough data to reconstruct a packet.
//line /usr/local/go/src/compress/flate/deflate.go:713
// Flush does not return until the data has been written.
//line /usr/local/go/src/compress/flate/deflate.go:713
// Calling Flush when there is no pending data still causes the Writer
//line /usr/local/go/src/compress/flate/deflate.go:713
// to emit a sync marker of at least 4 bytes.
//line /usr/local/go/src/compress/flate/deflate.go:713
// If the underlying writer returns an error, Flush returns that error.
//line /usr/local/go/src/compress/flate/deflate.go:713
//
//line /usr/local/go/src/compress/flate/deflate.go:713
// In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
//line /usr/local/go/src/compress/flate/deflate.go:722
func (w *Writer) Flush() error {
//line /usr/local/go/src/compress/flate/deflate.go:722
	_go_fuzz_dep_.CoverTab[25920]++

//line /usr/local/go/src/compress/flate/deflate.go:725
	return w.d.syncFlush()
//line /usr/local/go/src/compress/flate/deflate.go:725
	// _ = "end of CoverTab[25920]"
}

// Close flushes and closes the writer.
func (w *Writer) Close() error {
//line /usr/local/go/src/compress/flate/deflate.go:729
	_go_fuzz_dep_.CoverTab[25921]++
							return w.d.close()
//line /usr/local/go/src/compress/flate/deflate.go:730
	// _ = "end of CoverTab[25921]"
}

// Reset discards the writer's state and makes it equivalent to
//line /usr/local/go/src/compress/flate/deflate.go:733
// the result of NewWriter or NewWriterDict called with dst
//line /usr/local/go/src/compress/flate/deflate.go:733
// and w's level and dictionary.
//line /usr/local/go/src/compress/flate/deflate.go:736
func (w *Writer) Reset(dst io.Writer) {
//line /usr/local/go/src/compress/flate/deflate.go:736
	_go_fuzz_dep_.CoverTab[25922]++
							if dw, ok := w.d.w.writer.(*dictWriter); ok {
//line /usr/local/go/src/compress/flate/deflate.go:737
		_go_fuzz_dep_.CoverTab[25923]++

								dw.w = dst
								w.d.reset(dw)
								w.d.fillWindow(w.dict)
//line /usr/local/go/src/compress/flate/deflate.go:741
		// _ = "end of CoverTab[25923]"
	} else {
//line /usr/local/go/src/compress/flate/deflate.go:742
		_go_fuzz_dep_.CoverTab[25924]++

								w.d.reset(dst)
//line /usr/local/go/src/compress/flate/deflate.go:744
		// _ = "end of CoverTab[25924]"
	}
//line /usr/local/go/src/compress/flate/deflate.go:745
	// _ = "end of CoverTab[25922]"
}

//line /usr/local/go/src/compress/flate/deflate.go:746
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/flate/deflate.go:746
var _ = _go_fuzz_dep_.CoverTab
