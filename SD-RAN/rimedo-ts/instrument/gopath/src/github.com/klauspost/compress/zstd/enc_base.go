//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:1
)

import (
	"fmt"
	"math/bits"

	"github.com/klauspost/compress/zstd/internal/xxhash"
)

const (
	dictShardBits = 6
)

type fastBase struct {
	// cur is the offset at the start of hist
	cur	int32
	// maximum offset. Should be at least 2x block size.
	maxMatchOff	int32
	hist		[]byte
	crc		*xxhash.Digest
	tmp		[8]byte
	blk		*blockEnc
	lastDictID	uint32
	lowMem		bool
}

// CRC returns the underlying CRC writer.
func (e *fastBase) CRC() *xxhash.Digest {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:28
	_go_fuzz_dep_.CoverTab[92012]++
												return e.crc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:29
	// _ = "end of CoverTab[92012]"
}

// AppendCRC will append the CRC to the destination slice and return it.
func (e *fastBase) AppendCRC(dst []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:33
	_go_fuzz_dep_.CoverTab[92013]++
												crc := e.crc.Sum(e.tmp[:0])
												dst = append(dst, crc[7], crc[6], crc[5], crc[4])
												return dst
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:36
	// _ = "end of CoverTab[92013]"
}

// WindowSize returns the window size of the encoder,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:39
// or a window size small enough to contain the input size, if > 0.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:41
func (e *fastBase) WindowSize(size int64) int32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:41
	_go_fuzz_dep_.CoverTab[92014]++
												if size > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:42
		_go_fuzz_dep_.CoverTab[92016]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:42
		return size < int64(e.maxMatchOff)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:42
		// _ = "end of CoverTab[92016]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:42
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:42
		_go_fuzz_dep_.CoverTab[92017]++
													b := int32(1) << uint(bits.Len(uint(size)))

													if b < 1024 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:45
			_go_fuzz_dep_.CoverTab[92019]++
														b = 1024
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:46
			// _ = "end of CoverTab[92019]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:47
			_go_fuzz_dep_.CoverTab[92020]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:47
			// _ = "end of CoverTab[92020]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:47
		// _ = "end of CoverTab[92017]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:47
		_go_fuzz_dep_.CoverTab[92018]++
													return b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:48
		// _ = "end of CoverTab[92018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:49
		_go_fuzz_dep_.CoverTab[92021]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:49
		// _ = "end of CoverTab[92021]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:49
	// _ = "end of CoverTab[92014]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:49
	_go_fuzz_dep_.CoverTab[92015]++
												return e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:50
	// _ = "end of CoverTab[92015]"
}

// Block returns the current block.
func (e *fastBase) Block() *blockEnc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:54
	_go_fuzz_dep_.CoverTab[92022]++
												return e.blk
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:55
	// _ = "end of CoverTab[92022]"
}

func (e *fastBase) addBlock(src []byte) int32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:58
	_go_fuzz_dep_.CoverTab[92023]++
												if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:59
		_go_fuzz_dep_.CoverTab[92026]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:59
		return e.cur > bufferReset
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:59
		// _ = "end of CoverTab[92026]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:59
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:59
		_go_fuzz_dep_.CoverTab[92027]++
													panic(fmt.Sprintf("ecur (%d) > buffer reset (%d)", e.cur, bufferReset))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:60
		// _ = "end of CoverTab[92027]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:61
		_go_fuzz_dep_.CoverTab[92028]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:61
		// _ = "end of CoverTab[92028]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:61
	// _ = "end of CoverTab[92023]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:61
	_go_fuzz_dep_.CoverTab[92024]++

												if len(e.hist)+len(src) > cap(e.hist) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:63
		_go_fuzz_dep_.CoverTab[92029]++
													if cap(e.hist) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:64
			_go_fuzz_dep_.CoverTab[92030]++
														e.ensureHist(len(src))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:65
			// _ = "end of CoverTab[92030]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:66
			_go_fuzz_dep_.CoverTab[92031]++
														if cap(e.hist) < int(e.maxMatchOff+maxCompressedBlockSize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:67
				_go_fuzz_dep_.CoverTab[92033]++
															panic(fmt.Errorf("unexpected buffer cap %d, want at least %d with window %d", cap(e.hist), e.maxMatchOff+maxCompressedBlockSize, e.maxMatchOff))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:68
				// _ = "end of CoverTab[92033]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:69
				_go_fuzz_dep_.CoverTab[92034]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:69
				// _ = "end of CoverTab[92034]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:69
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:69
			// _ = "end of CoverTab[92031]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:69
			_go_fuzz_dep_.CoverTab[92032]++

														offset := int32(len(e.hist)) - e.maxMatchOff
														copy(e.hist[0:e.maxMatchOff], e.hist[offset:])
														e.cur += offset
														e.hist = e.hist[:e.maxMatchOff]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:74
			// _ = "end of CoverTab[92032]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:75
		// _ = "end of CoverTab[92029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:76
		_go_fuzz_dep_.CoverTab[92035]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:76
		// _ = "end of CoverTab[92035]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:76
	// _ = "end of CoverTab[92024]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:76
	_go_fuzz_dep_.CoverTab[92025]++
												s := int32(len(e.hist))
												e.hist = append(e.hist, src...)
												return s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:79
	// _ = "end of CoverTab[92025]"
}

// ensureHist will ensure that history can keep at least this many bytes.
func (e *fastBase) ensureHist(n int) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:83
	_go_fuzz_dep_.CoverTab[92036]++
												if cap(e.hist) >= n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:84
		_go_fuzz_dep_.CoverTab[92041]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:85
		// _ = "end of CoverTab[92041]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:86
		_go_fuzz_dep_.CoverTab[92042]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:86
		// _ = "end of CoverTab[92042]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:86
	// _ = "end of CoverTab[92036]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:86
	_go_fuzz_dep_.CoverTab[92037]++
												l := e.maxMatchOff
												if (e.lowMem && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		_go_fuzz_dep_.CoverTab[92043]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		return e.maxMatchOff > maxCompressedBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		// _ = "end of CoverTab[92043]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		_go_fuzz_dep_.CoverTab[92044]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		return e.maxMatchOff <= maxCompressedBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		// _ = "end of CoverTab[92044]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:88
		_go_fuzz_dep_.CoverTab[92045]++
													l += maxCompressedBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:89
		// _ = "end of CoverTab[92045]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:90
		_go_fuzz_dep_.CoverTab[92046]++
													l += e.maxMatchOff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:91
		// _ = "end of CoverTab[92046]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:92
	// _ = "end of CoverTab[92037]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:92
	_go_fuzz_dep_.CoverTab[92038]++

												if l < 1<<20 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:94
		_go_fuzz_dep_.CoverTab[92047]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:94
		return !e.lowMem
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:94
		// _ = "end of CoverTab[92047]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:94
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:94
		_go_fuzz_dep_.CoverTab[92048]++
													l = 1 << 20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:95
		// _ = "end of CoverTab[92048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:96
		_go_fuzz_dep_.CoverTab[92049]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:96
		// _ = "end of CoverTab[92049]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:96
	// _ = "end of CoverTab[92038]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:96
	_go_fuzz_dep_.CoverTab[92039]++

												if l < int32(n) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:98
		_go_fuzz_dep_.CoverTab[92050]++
													l = int32(n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:99
		// _ = "end of CoverTab[92050]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:100
		_go_fuzz_dep_.CoverTab[92051]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:100
		// _ = "end of CoverTab[92051]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:100
	// _ = "end of CoverTab[92039]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:100
	_go_fuzz_dep_.CoverTab[92040]++
												e.hist = make([]byte, 0, l)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:101
	// _ = "end of CoverTab[92040]"
}

// useBlock will replace the block with the provided one,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:104
// but transfer recent offsets from the previous.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:106
func (e *fastBase) UseBlock(enc *blockEnc) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:106
	_go_fuzz_dep_.CoverTab[92052]++
												enc.reset(e.blk)
												e.blk = enc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:108
	// _ = "end of CoverTab[92052]"
}

func (e *fastBase) matchlen(s, t int32, src []byte) int32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:111
	_go_fuzz_dep_.CoverTab[92053]++
												if debugAsserts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:112
		_go_fuzz_dep_.CoverTab[92057]++
													if s < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:113
			_go_fuzz_dep_.CoverTab[92061]++
														err := fmt.Sprintf("s (%d) < 0", s)
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:115
			// _ = "end of CoverTab[92061]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:116
			_go_fuzz_dep_.CoverTab[92062]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:116
			// _ = "end of CoverTab[92062]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:116
		// _ = "end of CoverTab[92057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:116
		_go_fuzz_dep_.CoverTab[92058]++
													if t < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:117
			_go_fuzz_dep_.CoverTab[92063]++
														err := fmt.Sprintf("s (%d) < 0", s)
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:119
			// _ = "end of CoverTab[92063]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:120
			_go_fuzz_dep_.CoverTab[92064]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:120
			// _ = "end of CoverTab[92064]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:120
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:120
		// _ = "end of CoverTab[92058]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:120
		_go_fuzz_dep_.CoverTab[92059]++
													if s-t > e.maxMatchOff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:121
			_go_fuzz_dep_.CoverTab[92065]++
														err := fmt.Sprintf("s (%d) - t (%d) > maxMatchOff (%d)", s, t, e.maxMatchOff)
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:123
			// _ = "end of CoverTab[92065]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:124
			_go_fuzz_dep_.CoverTab[92066]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:124
			// _ = "end of CoverTab[92066]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:124
		// _ = "end of CoverTab[92059]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:124
		_go_fuzz_dep_.CoverTab[92060]++
													if len(src)-int(s) > maxCompressedBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:125
			_go_fuzz_dep_.CoverTab[92067]++
														panic(fmt.Sprintf("len(src)-s (%d) > maxCompressedBlockSize (%d)", len(src)-int(s), maxCompressedBlockSize))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:126
			// _ = "end of CoverTab[92067]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:127
			_go_fuzz_dep_.CoverTab[92068]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:127
			// _ = "end of CoverTab[92068]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:127
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:127
		// _ = "end of CoverTab[92060]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:128
		_go_fuzz_dep_.CoverTab[92069]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:128
		// _ = "end of CoverTab[92069]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:128
	// _ = "end of CoverTab[92053]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:128
	_go_fuzz_dep_.CoverTab[92054]++
												a := src[s:]
												b := src[t:]
												b = b[:len(a)]
												end := int32((len(a) >> 3) << 3)
												for i := int32(0); i < end; i += 8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:133
		_go_fuzz_dep_.CoverTab[92070]++
													if diff := load6432(a, i) ^ load6432(b, i); diff != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:134
			_go_fuzz_dep_.CoverTab[92071]++
														return i + int32(bits.TrailingZeros64(diff)>>3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:135
			// _ = "end of CoverTab[92071]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:136
			_go_fuzz_dep_.CoverTab[92072]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:136
			// _ = "end of CoverTab[92072]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:136
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:136
		// _ = "end of CoverTab[92070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:137
	// _ = "end of CoverTab[92054]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:137
	_go_fuzz_dep_.CoverTab[92055]++

												a = a[end:]
												b = b[end:]
												for i := range a {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:141
		_go_fuzz_dep_.CoverTab[92073]++
													if a[i] != b[i] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:142
			_go_fuzz_dep_.CoverTab[92074]++
														return int32(i) + end
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:143
			// _ = "end of CoverTab[92074]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:144
			_go_fuzz_dep_.CoverTab[92075]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:144
			// _ = "end of CoverTab[92075]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:144
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:144
		// _ = "end of CoverTab[92073]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:145
	// _ = "end of CoverTab[92055]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:145
	_go_fuzz_dep_.CoverTab[92056]++
												return int32(len(a)) + end
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:146
	// _ = "end of CoverTab[92056]"
}

// Reset the encoding table.
func (e *fastBase) resetBase(d *dict, singleBlock bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:150
	_go_fuzz_dep_.CoverTab[92076]++
												if e.blk == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:151
		_go_fuzz_dep_.CoverTab[92081]++
													e.blk = &blockEnc{lowMem: e.lowMem}
													e.blk.init()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:153
		// _ = "end of CoverTab[92081]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:154
		_go_fuzz_dep_.CoverTab[92082]++
													e.blk.reset(nil)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:155
		// _ = "end of CoverTab[92082]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:156
	// _ = "end of CoverTab[92076]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:156
	_go_fuzz_dep_.CoverTab[92077]++
												e.blk.initNewEncode()
												if e.crc == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:158
		_go_fuzz_dep_.CoverTab[92083]++
													e.crc = xxhash.New()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:159
		// _ = "end of CoverTab[92083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:160
		_go_fuzz_dep_.CoverTab[92084]++
													e.crc.Reset()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:161
		// _ = "end of CoverTab[92084]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:162
	// _ = "end of CoverTab[92077]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:162
	_go_fuzz_dep_.CoverTab[92078]++
												if d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:163
		_go_fuzz_dep_.CoverTab[92085]++
													low := e.lowMem
													if singleBlock {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:165
			_go_fuzz_dep_.CoverTab[92087]++
														e.lowMem = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:166
			// _ = "end of CoverTab[92087]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:167
			_go_fuzz_dep_.CoverTab[92088]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:167
			// _ = "end of CoverTab[92088]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:167
		// _ = "end of CoverTab[92085]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:167
		_go_fuzz_dep_.CoverTab[92086]++
													e.ensureHist(d.DictContentSize() + maxCompressedBlockSize)
													e.lowMem = low
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:169
		// _ = "end of CoverTab[92086]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:170
		_go_fuzz_dep_.CoverTab[92089]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:170
		// _ = "end of CoverTab[92089]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:170
	// _ = "end of CoverTab[92078]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:170
	_go_fuzz_dep_.CoverTab[92079]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:174
	if e.cur < bufferReset {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:174
		_go_fuzz_dep_.CoverTab[92090]++
													e.cur += e.maxMatchOff + int32(len(e.hist))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:175
		// _ = "end of CoverTab[92090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:176
		_go_fuzz_dep_.CoverTab[92091]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:176
		// _ = "end of CoverTab[92091]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:176
	// _ = "end of CoverTab[92079]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:176
	_go_fuzz_dep_.CoverTab[92080]++
												e.hist = e.hist[:0]
												if d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:178
		_go_fuzz_dep_.CoverTab[92092]++

													for i, off := range d.offsets {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:180
			_go_fuzz_dep_.CoverTab[92094]++
														e.blk.recentOffsets[i] = uint32(off)
														e.blk.prevRecentOffsets[i] = e.blk.recentOffsets[i]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:182
			// _ = "end of CoverTab[92094]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:183
		// _ = "end of CoverTab[92092]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:183
		_go_fuzz_dep_.CoverTab[92093]++

													e.blk.dictLitEnc = d.litEnc
													e.hist = append(e.hist, d.content...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:186
		// _ = "end of CoverTab[92093]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:187
		_go_fuzz_dep_.CoverTab[92095]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:187
		// _ = "end of CoverTab[92095]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:187
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:187
	// _ = "end of CoverTab[92080]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:188
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/enc_base.go:188
var _ = _go_fuzz_dep_.CoverTab
