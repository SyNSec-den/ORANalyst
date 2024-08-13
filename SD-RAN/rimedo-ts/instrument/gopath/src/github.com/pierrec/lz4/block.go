//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:1
)

import (
	"encoding/binary"
	"math/bits"
	"sync"
)

// blockHash hashes the lower 6 bytes into a value < htSize.
func blockHash(x uint64) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:10
	_go_fuzz_dep_.CoverTab[95343]++
												const prime6bytes = 227718039650203
												return uint32(((x << (64 - 48)) * prime6bytes) >> (64 - hashLog))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:12
	// _ = "end of CoverTab[95343]"
}

// CompressBlockBound returns the maximum size of a given buffer of size n, when not compressible.
func CompressBlockBound(n int) int {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:16
	_go_fuzz_dep_.CoverTab[95344]++
												return n + n/255 + 16
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:17
	// _ = "end of CoverTab[95344]"
}

// UncompressBlock uncompresses the source buffer into the destination one,
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:20
// and returns the uncompressed size.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:20
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:20
// The destination buffer must be sized appropriately.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:20
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:20
// An error is returned if the source data is invalid or the destination buffer is too small.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:26
func UncompressBlock(src, dst []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:26
	_go_fuzz_dep_.CoverTab[95345]++
												if len(src) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:27
		_go_fuzz_dep_.CoverTab[95348]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:28
		// _ = "end of CoverTab[95348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:29
		_go_fuzz_dep_.CoverTab[95349]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:29
		// _ = "end of CoverTab[95349]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:29
	// _ = "end of CoverTab[95345]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:29
	_go_fuzz_dep_.CoverTab[95346]++
												if di := decodeBlock(dst, src); di >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:30
		_go_fuzz_dep_.CoverTab[95350]++
													return di, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:31
		// _ = "end of CoverTab[95350]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:32
		_go_fuzz_dep_.CoverTab[95351]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:32
		// _ = "end of CoverTab[95351]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:32
	// _ = "end of CoverTab[95346]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:32
	_go_fuzz_dep_.CoverTab[95347]++
												return 0, ErrInvalidSourceShortBuffer
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:33
	// _ = "end of CoverTab[95347]"
}

// CompressBlock compresses the source buffer into the destination one.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// This is the fast version of LZ4 compression and also the default one.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// The argument hashTable is scratch space for a hash table used by the
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// compressor. If provided, it should have length at least 1<<16. If it is
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// shorter (or nil), CompressBlock allocates its own hash table.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// The size of the compressed data is returned.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// If the destination buffer size is lower than CompressBlockBound and
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// the compressed size is 0 and no error, then the data is incompressible.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:36
// An error is returned if the destination buffer is too small.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:49
func CompressBlock(src, dst []byte, hashTable []int) (_ int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:49
	_go_fuzz_dep_.CoverTab[95352]++
												defer recoverBlock(&err)

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:53
	isNotCompressible := len(dst) < CompressBlockBound(len(src))

	// adaptSkipLog sets how quickly the compressor begins skipping blocks when data is incompressible.
	// This significantly speeds up incompressible data and usually has very small impact on compression.
	// bytes to skip =  1 + (bytes since last match >> adaptSkipLog)
	const adaptSkipLog = 7
	if len(hashTable) < htSize {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:59
		_go_fuzz_dep_.CoverTab[95359]++
													htIface := htPool.Get()
													defer htPool.Put(htIface)
													hashTable = (*(htIface).(*[htSize]int))[:]
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:62
		// _ = "end of CoverTab[95359]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:63
		_go_fuzz_dep_.CoverTab[95360]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:63
		// _ = "end of CoverTab[95360]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:63
	// _ = "end of CoverTab[95352]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:63
	_go_fuzz_dep_.CoverTab[95353]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:66
	hashTable = hashTable[:htSize]

	// si: Current position of the search.
	// anchor: Position of the current literals.
	var si, di, anchor int
	sn := len(src) - mfLimit
	if sn <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:72
		_go_fuzz_dep_.CoverTab[95361]++
													goto lastLiterals
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:73
		// _ = "end of CoverTab[95361]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:74
		_go_fuzz_dep_.CoverTab[95362]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:74
		// _ = "end of CoverTab[95362]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:74
	// _ = "end of CoverTab[95353]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:74
	_go_fuzz_dep_.CoverTab[95354]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:77
	for si < sn {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:77
		_go_fuzz_dep_.CoverTab[95363]++

													match := binary.LittleEndian.Uint64(src[si:])
													h := blockHash(match)
													h2 := blockHash(match >> 8)

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:85
		ref := hashTable[h]
													ref2 := hashTable[h2]
													hashTable[h] = si
													hashTable[h2] = si + 1
													offset := si - ref

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
		if offset <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
			_go_fuzz_dep_.CoverTab[95371]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
			return offset >= winSize
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
			// _ = "end of CoverTab[95371]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
			_go_fuzz_dep_.CoverTab[95372]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:92
			return uint32(match) != binary.LittleEndian.Uint32(src[ref:])
														// _ = "end of CoverTab[95372]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:93
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:93
			_go_fuzz_dep_.CoverTab[95373]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:96
			h = blockHash(match >> 16)
														ref = hashTable[h]

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:100
			si += 1
			offset = si - ref2

			if offset <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:103
				_go_fuzz_dep_.CoverTab[95374]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:103
				return offset >= winSize
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:103
				// _ = "end of CoverTab[95374]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:103
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:103
				_go_fuzz_dep_.CoverTab[95375]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:103
				return uint32(match>>8) != binary.LittleEndian.Uint32(src[ref2:])
															// _ = "end of CoverTab[95375]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:104
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:104
				_go_fuzz_dep_.CoverTab[95376]++

															si += 1
															offset = si - ref
															hashTable[h] = si

															if offset <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:110
					_go_fuzz_dep_.CoverTab[95377]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:110
					return offset >= winSize
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:110
					// _ = "end of CoverTab[95377]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:110
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:110
					_go_fuzz_dep_.CoverTab[95378]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:110
					return uint32(match>>16) != binary.LittleEndian.Uint32(src[ref:])
																// _ = "end of CoverTab[95378]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:111
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:111
					_go_fuzz_dep_.CoverTab[95379]++

																si += 2 + (si-anchor)>>adaptSkipLog
																continue
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:114
					// _ = "end of CoverTab[95379]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:115
					_go_fuzz_dep_.CoverTab[95380]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:115
					// _ = "end of CoverTab[95380]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:115
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:115
				// _ = "end of CoverTab[95376]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:116
				_go_fuzz_dep_.CoverTab[95381]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:116
				// _ = "end of CoverTab[95381]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:116
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:116
			// _ = "end of CoverTab[95373]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:117
			_go_fuzz_dep_.CoverTab[95382]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:117
			// _ = "end of CoverTab[95382]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:117
		// _ = "end of CoverTab[95363]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:117
		_go_fuzz_dep_.CoverTab[95364]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:120
		lLen := si - anchor

													mLen := 4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:125
		tOff := si - offset - 1
		for lLen > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			_go_fuzz_dep_.CoverTab[95383]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			return tOff >= 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			// _ = "end of CoverTab[95383]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			_go_fuzz_dep_.CoverTab[95384]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			return src[si-1] == src[tOff]
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			// _ = "end of CoverTab[95384]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:126
			_go_fuzz_dep_.CoverTab[95385]++
														si--
														tOff--
														lLen--
														mLen++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:130
			// _ = "end of CoverTab[95385]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:131
		// _ = "end of CoverTab[95364]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:131
		_go_fuzz_dep_.CoverTab[95365]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:135
		si, mLen = si+mLen, si+minMatch

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:138
		for si+8 < sn {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:138
			_go_fuzz_dep_.CoverTab[95386]++
														x := binary.LittleEndian.Uint64(src[si:]) ^ binary.LittleEndian.Uint64(src[si-offset:])
														if x == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:140
				_go_fuzz_dep_.CoverTab[95387]++
															si += 8
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:141
				// _ = "end of CoverTab[95387]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:142
				_go_fuzz_dep_.CoverTab[95388]++

															si += bits.TrailingZeros64(x) >> 3
															break
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:145
				// _ = "end of CoverTab[95388]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:146
			// _ = "end of CoverTab[95386]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:147
		// _ = "end of CoverTab[95365]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:147
		_go_fuzz_dep_.CoverTab[95366]++

													mLen = si - mLen
													if mLen < 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:150
			_go_fuzz_dep_.CoverTab[95389]++
														dst[di] = byte(mLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:151
			// _ = "end of CoverTab[95389]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:152
			_go_fuzz_dep_.CoverTab[95390]++
														dst[di] = 0xF
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:153
			// _ = "end of CoverTab[95390]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:154
		// _ = "end of CoverTab[95366]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:154
		_go_fuzz_dep_.CoverTab[95367]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:157
		if lLen < 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:157
			_go_fuzz_dep_.CoverTab[95391]++
														dst[di] |= byte(lLen << 4)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:158
			// _ = "end of CoverTab[95391]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:159
			_go_fuzz_dep_.CoverTab[95392]++
														dst[di] |= 0xF0
														di++
														l := lLen - 0xF
														for ; l >= 0xFF; l -= 0xFF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:163
				_go_fuzz_dep_.CoverTab[95394]++
															dst[di] = 0xFF
															di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:165
				// _ = "end of CoverTab[95394]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:166
			// _ = "end of CoverTab[95392]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:166
			_go_fuzz_dep_.CoverTab[95393]++
														dst[di] = byte(l)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:167
			// _ = "end of CoverTab[95393]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:168
		// _ = "end of CoverTab[95367]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:168
		_go_fuzz_dep_.CoverTab[95368]++
													di++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:172
		copy(dst[di:di+lLen], src[anchor:anchor+lLen])
													di += lLen + 2
													anchor = si

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:177
		_ = dst[di]
													dst[di-2], dst[di-1] = byte(offset), byte(offset>>8)

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:181
		if mLen >= 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:181
			_go_fuzz_dep_.CoverTab[95395]++
														for mLen -= 0xF; mLen >= 0xFF; mLen -= 0xFF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:182
				_go_fuzz_dep_.CoverTab[95397]++
															dst[di] = 0xFF
															di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:184
				// _ = "end of CoverTab[95397]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:185
			// _ = "end of CoverTab[95395]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:185
			_go_fuzz_dep_.CoverTab[95396]++
														dst[di] = byte(mLen)
														di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:187
			// _ = "end of CoverTab[95396]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:188
			_go_fuzz_dep_.CoverTab[95398]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:188
			// _ = "end of CoverTab[95398]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:188
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:188
		// _ = "end of CoverTab[95368]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:188
		_go_fuzz_dep_.CoverTab[95369]++

													if si >= sn {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:190
			_go_fuzz_dep_.CoverTab[95399]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:191
			// _ = "end of CoverTab[95399]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:192
			_go_fuzz_dep_.CoverTab[95400]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:192
			// _ = "end of CoverTab[95400]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:192
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:192
		// _ = "end of CoverTab[95369]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:192
		_go_fuzz_dep_.CoverTab[95370]++

													h = blockHash(binary.LittleEndian.Uint64(src[si-2:]))
													hashTable[h] = si - 2
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:195
		// _ = "end of CoverTab[95370]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:196
	// _ = "end of CoverTab[95354]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:196
	_go_fuzz_dep_.CoverTab[95355]++

lastLiterals:
	if isNotCompressible && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:199
		_go_fuzz_dep_.CoverTab[95401]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:199
		return anchor == 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:199
		// _ = "end of CoverTab[95401]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:199
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:199
		_go_fuzz_dep_.CoverTab[95402]++

													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:201
		// _ = "end of CoverTab[95402]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:202
		_go_fuzz_dep_.CoverTab[95403]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:202
		// _ = "end of CoverTab[95403]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:202
	// _ = "end of CoverTab[95355]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:202
	_go_fuzz_dep_.CoverTab[95356]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:205
	lLen := len(src) - anchor
	if lLen < 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:206
		_go_fuzz_dep_.CoverTab[95404]++
													dst[di] = byte(lLen << 4)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:207
		// _ = "end of CoverTab[95404]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:208
		_go_fuzz_dep_.CoverTab[95405]++
													dst[di] = 0xF0
													di++
													for lLen -= 0xF; lLen >= 0xFF; lLen -= 0xFF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:211
			_go_fuzz_dep_.CoverTab[95407]++
														dst[di] = 0xFF
														di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:213
			// _ = "end of CoverTab[95407]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:214
		// _ = "end of CoverTab[95405]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:214
		_go_fuzz_dep_.CoverTab[95406]++
													dst[di] = byte(lLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:215
		// _ = "end of CoverTab[95406]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:216
	// _ = "end of CoverTab[95356]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:216
	_go_fuzz_dep_.CoverTab[95357]++
												di++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:220
	if isNotCompressible && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:220
		_go_fuzz_dep_.CoverTab[95408]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:220
		return di >= anchor
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:220
		// _ = "end of CoverTab[95408]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:220
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:220
		_go_fuzz_dep_.CoverTab[95409]++

													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:222
		// _ = "end of CoverTab[95409]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:223
		_go_fuzz_dep_.CoverTab[95410]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:223
		// _ = "end of CoverTab[95410]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:223
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:223
	// _ = "end of CoverTab[95357]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:223
	_go_fuzz_dep_.CoverTab[95358]++
												di += copy(dst[di:di+len(src)-anchor], src[anchor:])
												return di, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:225
	// _ = "end of CoverTab[95358]"
}

// Pool of hash tables for CompressBlock.
var htPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:230
		_go_fuzz_dep_.CoverTab[95411]++
													return new([htSize]int)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:231
		// _ = "end of CoverTab[95411]"
	},
}

// blockHash hashes 4 bytes into a value < winSize.
func blockHashHC(x uint32) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:236
	_go_fuzz_dep_.CoverTab[95412]++
												const hasher uint32 = 2654435761	// Knuth multiplicative hash.
												return x * hasher >> (32 - winSizeLog)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:238
	// _ = "end of CoverTab[95412]"
}

// CompressBlockHC compresses the source buffer src into the destination dst
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
// with max search depth (use 0 or negative value for no max).
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
// CompressBlockHC compression ratio is better than CompressBlock but it is also slower.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
// The size of the compressed data is returned.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
// If the destination buffer size is lower than CompressBlockBound and
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
// the compressed size is 0 and no error, then the data is incompressible.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:241
// An error is returned if the destination buffer is too small.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:252
func CompressBlockHC(src, dst []byte, depth int) (_ int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:252
	_go_fuzz_dep_.CoverTab[95413]++
												defer recoverBlock(&err)

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:256
	isNotCompressible := len(dst) < CompressBlockBound(len(src))

	// adaptSkipLog sets how quickly the compressor begins skipping blocks when data is incompressible.
	// This significantly speeds up incompressible data and usually has very small impact on compression.
	// bytes to skip =  1 + (bytes since last match >> adaptSkipLog)
	const adaptSkipLog = 7

	var si, di, anchor int

	// hashTable: stores the last position found for a given hash
	// chainTable: stores previous positions for a given hash
	var hashTable, chainTable [winSize]int

	if depth <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:269
		_go_fuzz_dep_.CoverTab[95420]++
													depth = winSize
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:270
		// _ = "end of CoverTab[95420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:271
		_go_fuzz_dep_.CoverTab[95421]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:271
		// _ = "end of CoverTab[95421]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:271
	// _ = "end of CoverTab[95413]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:271
	_go_fuzz_dep_.CoverTab[95414]++

												sn := len(src) - mfLimit
												if sn <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:274
		_go_fuzz_dep_.CoverTab[95422]++
													goto lastLiterals
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:275
		// _ = "end of CoverTab[95422]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:276
		_go_fuzz_dep_.CoverTab[95423]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:276
		// _ = "end of CoverTab[95423]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:276
	// _ = "end of CoverTab[95414]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:276
	_go_fuzz_dep_.CoverTab[95415]++

												for si < sn {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:278
		_go_fuzz_dep_.CoverTab[95424]++

													match := binary.LittleEndian.Uint32(src[si:])
													h := blockHashHC(match)

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:284
		mLen := 0
		offset := 0
		for next, try := hashTable[h], depth; try > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			_go_fuzz_dep_.CoverTab[95431]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			return next > 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			// _ = "end of CoverTab[95431]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			_go_fuzz_dep_.CoverTab[95432]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			return si-next < winSize
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			// _ = "end of CoverTab[95432]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
		}(); next = chainTable[next&winMask] {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:286
			_go_fuzz_dep_.CoverTab[95433]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:289
			if src[next+mLen] != src[si+mLen] {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:289
				_go_fuzz_dep_.CoverTab[95437]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:290
				// _ = "end of CoverTab[95437]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:291
				_go_fuzz_dep_.CoverTab[95438]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:291
				// _ = "end of CoverTab[95438]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:291
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:291
			// _ = "end of CoverTab[95433]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:291
			_go_fuzz_dep_.CoverTab[95434]++
														ml := 0

														for ml < sn-si {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:294
				_go_fuzz_dep_.CoverTab[95439]++
															x := binary.LittleEndian.Uint64(src[next+ml:]) ^ binary.LittleEndian.Uint64(src[si+ml:])
															if x == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:296
					_go_fuzz_dep_.CoverTab[95440]++
																ml += 8
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:297
					// _ = "end of CoverTab[95440]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:298
					_go_fuzz_dep_.CoverTab[95441]++

																ml += bits.TrailingZeros64(x) >> 3
																break
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:301
					// _ = "end of CoverTab[95441]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:302
				// _ = "end of CoverTab[95439]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:303
			// _ = "end of CoverTab[95434]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:303
			_go_fuzz_dep_.CoverTab[95435]++
														if ml < minMatch || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:304
				_go_fuzz_dep_.CoverTab[95442]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:304
				return ml <= mLen
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:304
				// _ = "end of CoverTab[95442]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:304
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:304
				_go_fuzz_dep_.CoverTab[95443]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:306
				// _ = "end of CoverTab[95443]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:307
				_go_fuzz_dep_.CoverTab[95444]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:307
				// _ = "end of CoverTab[95444]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:307
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:307
			// _ = "end of CoverTab[95435]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:307
			_go_fuzz_dep_.CoverTab[95436]++

														mLen = ml
														offset = si - next

														try--
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:312
			// _ = "end of CoverTab[95436]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:313
		// _ = "end of CoverTab[95424]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:313
		_go_fuzz_dep_.CoverTab[95425]++
													chainTable[si&winMask] = hashTable[h]
													hashTable[h] = si

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:318
		if mLen == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:318
			_go_fuzz_dep_.CoverTab[95445]++
														si += 1 + (si-anchor)>>adaptSkipLog
														continue
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:320
			// _ = "end of CoverTab[95445]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:321
			_go_fuzz_dep_.CoverTab[95446]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:321
			// _ = "end of CoverTab[95446]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:321
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:321
		// _ = "end of CoverTab[95425]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:321
		_go_fuzz_dep_.CoverTab[95426]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:326
		winStart := si + 1
		if ws := si + mLen - winSize; ws > winStart {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:327
			_go_fuzz_dep_.CoverTab[95447]++
														winStart = ws
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:328
			// _ = "end of CoverTab[95447]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:329
			_go_fuzz_dep_.CoverTab[95448]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:329
			// _ = "end of CoverTab[95448]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:329
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:329
		// _ = "end of CoverTab[95426]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:329
		_go_fuzz_dep_.CoverTab[95427]++
													for si, ml := winStart, si+mLen; si < ml; {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:330
			_go_fuzz_dep_.CoverTab[95449]++
														match >>= 8
														match |= uint32(src[si+3]) << 24
														h := blockHashHC(match)
														chainTable[si&winMask] = hashTable[h]
														hashTable[h] = si
														si++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:336
			// _ = "end of CoverTab[95449]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:337
		// _ = "end of CoverTab[95427]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:337
		_go_fuzz_dep_.CoverTab[95428]++

													lLen := si - anchor
													si += mLen
													mLen -= minMatch

													if mLen < 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:343
			_go_fuzz_dep_.CoverTab[95450]++
														dst[di] = byte(mLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:344
			// _ = "end of CoverTab[95450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:345
			_go_fuzz_dep_.CoverTab[95451]++
														dst[di] = 0xF
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:346
			// _ = "end of CoverTab[95451]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:347
		// _ = "end of CoverTab[95428]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:347
		_go_fuzz_dep_.CoverTab[95429]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:350
		if lLen < 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:350
			_go_fuzz_dep_.CoverTab[95452]++
														dst[di] |= byte(lLen << 4)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:351
			// _ = "end of CoverTab[95452]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:352
			_go_fuzz_dep_.CoverTab[95453]++
														dst[di] |= 0xF0
														di++
														l := lLen - 0xF
														for ; l >= 0xFF; l -= 0xFF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:356
				_go_fuzz_dep_.CoverTab[95455]++
															dst[di] = 0xFF
															di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:358
				// _ = "end of CoverTab[95455]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:359
			// _ = "end of CoverTab[95453]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:359
			_go_fuzz_dep_.CoverTab[95454]++
														dst[di] = byte(l)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:360
			// _ = "end of CoverTab[95454]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:361
		// _ = "end of CoverTab[95429]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:361
		_go_fuzz_dep_.CoverTab[95430]++
													di++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:365
		copy(dst[di:di+lLen], src[anchor:anchor+lLen])
													di += lLen
													anchor = si

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:370
		di += 2
													dst[di-2], dst[di-1] = byte(offset), byte(offset>>8)

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:374
		if mLen >= 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:374
			_go_fuzz_dep_.CoverTab[95456]++
														for mLen -= 0xF; mLen >= 0xFF; mLen -= 0xFF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:375
				_go_fuzz_dep_.CoverTab[95458]++
															dst[di] = 0xFF
															di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:377
				// _ = "end of CoverTab[95458]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:378
			// _ = "end of CoverTab[95456]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:378
			_go_fuzz_dep_.CoverTab[95457]++
														dst[di] = byte(mLen)
														di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:380
			// _ = "end of CoverTab[95457]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:381
			_go_fuzz_dep_.CoverTab[95459]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:381
			// _ = "end of CoverTab[95459]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:381
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:381
		// _ = "end of CoverTab[95430]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:382
	// _ = "end of CoverTab[95415]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:382
	_go_fuzz_dep_.CoverTab[95416]++

												if isNotCompressible && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:384
		_go_fuzz_dep_.CoverTab[95460]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:384
		return anchor == 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:384
		// _ = "end of CoverTab[95460]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:384
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:384
		_go_fuzz_dep_.CoverTab[95461]++

													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:386
		// _ = "end of CoverTab[95461]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:387
		_go_fuzz_dep_.CoverTab[95462]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:387
		// _ = "end of CoverTab[95462]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:387
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:387
	// _ = "end of CoverTab[95416]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:387
	_go_fuzz_dep_.CoverTab[95417]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:390
lastLiterals:
	lLen := len(src) - anchor
	if lLen < 0xF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:392
		_go_fuzz_dep_.CoverTab[95463]++
													dst[di] = byte(lLen << 4)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:393
		// _ = "end of CoverTab[95463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:394
		_go_fuzz_dep_.CoverTab[95464]++
													dst[di] = 0xF0
													di++
													lLen -= 0xF
													for ; lLen >= 0xFF; lLen -= 0xFF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:398
			_go_fuzz_dep_.CoverTab[95466]++
														dst[di] = 0xFF
														di++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:400
			// _ = "end of CoverTab[95466]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:401
		// _ = "end of CoverTab[95464]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:401
		_go_fuzz_dep_.CoverTab[95465]++
													dst[di] = byte(lLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:402
		// _ = "end of CoverTab[95465]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:403
	// _ = "end of CoverTab[95417]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:403
	_go_fuzz_dep_.CoverTab[95418]++
												di++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:407
	if isNotCompressible && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:407
		_go_fuzz_dep_.CoverTab[95467]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:407
		return di >= anchor
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:407
		// _ = "end of CoverTab[95467]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:407
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:407
		_go_fuzz_dep_.CoverTab[95468]++

													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:409
		// _ = "end of CoverTab[95468]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:410
		_go_fuzz_dep_.CoverTab[95469]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:410
		// _ = "end of CoverTab[95469]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:410
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:410
	// _ = "end of CoverTab[95418]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:410
	_go_fuzz_dep_.CoverTab[95419]++
												di += copy(dst[di:di+len(src)-anchor], src[anchor:])
												return di, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:412
	// _ = "end of CoverTab[95419]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:413
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/block.go:413
var _ = _go_fuzz_dep_.CoverTab
