//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// Package lz4 implements reading and writing lz4 compressed data (a frame),
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// as specified in http://fastcompression.blogspot.fr/2013/04/lz4-streaming-format-final.html.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// Although the block level compression and decompression functions are exposed and are fully compatible
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// with the lz4 block format definition, they are low level and should not be used directly.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// For a complete description of an lz4 compressed block, see:
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// http://fastcompression.blogspot.fr/2011/05/lz4-explained.html
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:1
// See https://github.com/Cyan4973/lz4 for the reference C implementation.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:11
)

import (
	"math/bits"
	"sync"
)

const (
	// Extension is the LZ4 frame file name extension
	Extension	= ".lz4"
	// Version is the LZ4 frame format version
	Version	= 1

	frameMagic		uint32	= 0x184D2204
	frameSkipMagic		uint32	= 0x184D2A50
	frameMagicLegacy	uint32	= 0x184C2102

	// The following constants are used to setup the compression algorithm.
	minMatch		= 4	// the minimum size of the match sequence size (4 bytes)
	winSizeLog		= 16	// LZ4 64Kb window size limit
	winSize			= 1 << winSizeLog
	winMask			= winSize - 1	// 64Kb window of previous data for dependent blocks
	compressedBlockFlag	= 1 << 31
	compressedBlockMask	= compressedBlockFlag - 1

	// hashLog determines the size of the hash table used to quickly find a previous match position.
	// Its value influences the compression speed and memory usage, the lower the faster,
	// but at the expense of the compression ratio.
	// 16 seems to be the best compromise for fast compression.
	hashLog	= 16
	htSize	= 1 << hashLog

	mfLimit	= 10 + minMatch	// The last match cannot start within the last 14 bytes.
)

// map the block max size id with its value in bytes: 64Kb, 256Kb, 1Mb and 4Mb.
const (
	blockSize64K	= 1 << (16 + 2*iota)
	blockSize256K
	blockSize1M
	blockSize4M
)

var (
	// Keep a pool of buffers for each valid block sizes.
	bsMapValue = [...]*sync.Pool{
		newBufferPool(2 * blockSize64K),
		newBufferPool(2 * blockSize256K),
		newBufferPool(2 * blockSize1M),
		newBufferPool(2 * blockSize4M),
	}
)

// newBufferPool returns a pool for buffers of the given size.
func newBufferPool(size int) *sync.Pool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:65
	_go_fuzz_dep_.CoverTab[95478]++
											return &sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:67
			_go_fuzz_dep_.CoverTab[95479]++
													return make([]byte, size)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:68
			// _ = "end of CoverTab[95479]"
		},
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:70
	// _ = "end of CoverTab[95478]"
}

// getBuffer returns a buffer to its pool.
func getBuffer(size int) []byte {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:74
	_go_fuzz_dep_.CoverTab[95480]++
											idx := blockSizeValueToIndex(size) - 4
											return bsMapValue[idx].Get().([]byte)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:76
	// _ = "end of CoverTab[95480]"
}

// putBuffer returns a buffer to its pool.
func putBuffer(size int, buf []byte) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:80
	_go_fuzz_dep_.CoverTab[95481]++
											if cap(buf) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:81
		_go_fuzz_dep_.CoverTab[95482]++
												idx := blockSizeValueToIndex(size) - 4
												bsMapValue[idx].Put(buf[:cap(buf)])
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:83
		// _ = "end of CoverTab[95482]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:84
		_go_fuzz_dep_.CoverTab[95483]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:84
		// _ = "end of CoverTab[95483]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:84
	// _ = "end of CoverTab[95481]"
}
func blockSizeIndexToValue(i byte) int {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:86
	_go_fuzz_dep_.CoverTab[95484]++
											return 1 << (16 + 2*uint(i))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:87
	// _ = "end of CoverTab[95484]"
}
func isValidBlockSize(size int) bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:89
	_go_fuzz_dep_.CoverTab[95485]++
											const blockSizeMask = blockSize64K | blockSize256K | blockSize1M | blockSize4M

											return size&blockSizeMask > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:92
		_go_fuzz_dep_.CoverTab[95486]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:92
		return bits.OnesCount(uint(size)) == 1
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:92
		// _ = "end of CoverTab[95486]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:92
	}()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:92
	// _ = "end of CoverTab[95485]"
}
func blockSizeValueToIndex(size int) byte {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:94
	_go_fuzz_dep_.CoverTab[95487]++
											return 4 + byte(bits.TrailingZeros(uint(size)>>16)/2)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:95
	// _ = "end of CoverTab[95487]"
}

// Header describes the various flags that can be set on a Writer or obtained from a Reader.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:98
// The default values match those of the LZ4 frame format definition
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:98
// (http://fastcompression.blogspot.com/2013/04/lz4-streaming-format-final.html).
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:98
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:98
// NB. in a Reader, in case of concatenated frames, the Header values may change between Read() calls.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:98
// It is the caller's responsibility to check them if necessary.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:104
type Header struct {
	BlockChecksum		bool	// Compressed blocks checksum flag.
	NoChecksum		bool	// Frame checksum flag.
	BlockMaxSize		int	// Size of the uncompressed data block (one of [64KB, 256KB, 1MB, 4MB]). Default=4MB.
	Size			uint64	// Frame total size. It is _not_ computed by the Writer.
	CompressionLevel	int	// Compression level (higher is better, use 0 for fastest compression).
	done			bool	// Header processed flag (Read or Write and checked).
}

// Reset reset internal status
func (h *Header) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:114
	_go_fuzz_dep_.CoverTab[95488]++
												h.done = false
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:115
	// _ = "end of CoverTab[95488]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:116
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/lz4.go:116
var _ = _go_fuzz_dep_.CoverTab
