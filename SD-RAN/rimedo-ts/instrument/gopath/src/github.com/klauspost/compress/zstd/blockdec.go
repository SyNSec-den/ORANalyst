// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:5
)

import (
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/klauspost/compress/huff0"
	"github.com/klauspost/compress/zstd/internal/xxhash"
)

type blockType uint8

//go:generate stringer -type=blockType,literalsBlockType,seqCompMode,tableIndex

const (
	blockTypeRaw	blockType	= iota
	blockTypeRLE
	blockTypeCompressed
	blockTypeReserved
)

type literalsBlockType uint8

const (
	literalsBlockRaw	literalsBlockType	= iota
	literalsBlockRLE
	literalsBlockCompressed
	literalsBlockTreeless
)

const (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:39
	maxCompressedBlockSize	= 128 << 10

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:42
	maxBlockSize	= (1 << 21) - 1

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:45
	maxCompressedLiteralSize	= 1 << 18
												maxRLELiteralSize	= 1 << 20
												maxMatchLen		= 131074
												maxSequences		= 0x7f00 + 0xffff

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:52
	maxOffsetBits	= 30
)

var (
	huffDecoderPool	= sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:56
		_go_fuzz_dep_.CoverTab[90898]++
													return &huff0.Scratch{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:57
		// _ = "end of CoverTab[90898]"
	}}

	fseDecoderPool	= sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:60
		_go_fuzz_dep_.CoverTab[90899]++
													return &fseDecoder{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:61
		// _ = "end of CoverTab[90899]"
	}}
)

type blockDec struct {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:67
	data	[]byte
												dataStorage	[]byte

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:71
	dst	[]byte

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:74
	literalBuf	[]byte

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:77
	WindowSize	uint64

												history	chan *history
												input	chan struct{}
												result	chan decodeOutput
												err	error
												decWG	sync.WaitGroup

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:87
	localFrame	*frameDec

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:90
	RLESize	uint32
												tmp	[4]byte

												Type	blockType

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:96
	Last	bool

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:99
	lowMem	bool
}

func (b *blockDec) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:102
	_go_fuzz_dep_.CoverTab[90900]++
												if b == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:103
		_go_fuzz_dep_.CoverTab[90902]++
													return "<nil>"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:104
		// _ = "end of CoverTab[90902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:105
		_go_fuzz_dep_.CoverTab[90903]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:105
		// _ = "end of CoverTab[90903]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:105
	// _ = "end of CoverTab[90900]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:105
	_go_fuzz_dep_.CoverTab[90901]++
												return fmt.Sprintf("Steam Size: %d, Type: %v, Last: %t, Window: %d", len(b.data), b.Type, b.Last, b.WindowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:106
	// _ = "end of CoverTab[90901]"
}

func newBlockDec(lowMem bool) *blockDec {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:109
	_go_fuzz_dep_.CoverTab[90904]++
												b := blockDec{
		lowMem:		lowMem,
		result:		make(chan decodeOutput, 1),
		input:		make(chan struct{}, 1),
		history:	make(chan *history, 1),
	}
												b.decWG.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:116
	_curRoutineNum103_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:116
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum103_)
												go b.startDecoder()
												return &b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:118
	// _ = "end of CoverTab[90904]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:123
func (b *blockDec) reset(br byteBuffer, windowSize uint64) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:123
	_go_fuzz_dep_.CoverTab[90905]++
												b.WindowSize = windowSize
												tmp, err := br.readSmall(3)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:126
		_go_fuzz_dep_.CoverTab[90911]++
													println("Reading block header:", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:128
		// _ = "end of CoverTab[90911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:129
		_go_fuzz_dep_.CoverTab[90912]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:129
		// _ = "end of CoverTab[90912]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:129
	// _ = "end of CoverTab[90905]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:129
	_go_fuzz_dep_.CoverTab[90906]++
												bh := uint32(tmp[0]) | (uint32(tmp[1]) << 8) | (uint32(tmp[2]) << 16)
												b.Last = bh&1 != 0
												b.Type = blockType((bh >> 1) & 3)

												cSize := int(bh >> 3)
												maxSize := maxBlockSize
												switch b.Type {
	case blockTypeReserved:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:137
		_go_fuzz_dep_.CoverTab[90913]++
													return ErrReservedBlockType
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:138
		// _ = "end of CoverTab[90913]"
	case blockTypeRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:139
		_go_fuzz_dep_.CoverTab[90914]++
													b.RLESize = uint32(cSize)
													if b.lowMem {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:141
			_go_fuzz_dep_.CoverTab[90921]++
														maxSize = cSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:142
			// _ = "end of CoverTab[90921]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:143
			_go_fuzz_dep_.CoverTab[90922]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:143
			// _ = "end of CoverTab[90922]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:143
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:143
		// _ = "end of CoverTab[90914]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:143
		_go_fuzz_dep_.CoverTab[90915]++
													cSize = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:144
		// _ = "end of CoverTab[90915]"
	case blockTypeCompressed:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:145
		_go_fuzz_dep_.CoverTab[90916]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:146
			_go_fuzz_dep_.CoverTab[90923]++
														println("Data size on stream:", cSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:147
			// _ = "end of CoverTab[90923]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:148
			_go_fuzz_dep_.CoverTab[90924]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:148
			// _ = "end of CoverTab[90924]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:148
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:148
		// _ = "end of CoverTab[90916]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:148
		_go_fuzz_dep_.CoverTab[90917]++
													b.RLESize = 0
													maxSize = maxCompressedBlockSize
													if windowSize < maxCompressedBlockSize && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:151
			_go_fuzz_dep_.CoverTab[90925]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:151
			return b.lowMem
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:151
			// _ = "end of CoverTab[90925]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:151
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:151
			_go_fuzz_dep_.CoverTab[90926]++
														maxSize = int(windowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:152
			// _ = "end of CoverTab[90926]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:153
			_go_fuzz_dep_.CoverTab[90927]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:153
			// _ = "end of CoverTab[90927]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:153
		// _ = "end of CoverTab[90917]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:153
		_go_fuzz_dep_.CoverTab[90918]++
													if cSize > maxCompressedBlockSize || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:154
			_go_fuzz_dep_.CoverTab[90928]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:154
			return uint64(cSize) > b.WindowSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:154
			// _ = "end of CoverTab[90928]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:154
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:154
			_go_fuzz_dep_.CoverTab[90929]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:155
				_go_fuzz_dep_.CoverTab[90931]++
															printf("compressed block too big: csize:%d block: %+v\n", uint64(cSize), b)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:156
				// _ = "end of CoverTab[90931]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:157
				_go_fuzz_dep_.CoverTab[90932]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:157
				// _ = "end of CoverTab[90932]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:157
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:157
			// _ = "end of CoverTab[90929]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:157
			_go_fuzz_dep_.CoverTab[90930]++
														return ErrCompressedSizeTooBig
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:158
			// _ = "end of CoverTab[90930]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:159
			_go_fuzz_dep_.CoverTab[90933]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:159
			// _ = "end of CoverTab[90933]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:159
		// _ = "end of CoverTab[90918]"
	case blockTypeRaw:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:160
		_go_fuzz_dep_.CoverTab[90919]++
													b.RLESize = 0

													maxSize = -1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:163
		// _ = "end of CoverTab[90919]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:164
		_go_fuzz_dep_.CoverTab[90920]++
													panic("Invalid block type")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:165
		// _ = "end of CoverTab[90920]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:166
	// _ = "end of CoverTab[90906]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:166
	_go_fuzz_dep_.CoverTab[90907]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:169
	if cap(b.dataStorage) < cSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:169
		_go_fuzz_dep_.CoverTab[90934]++
													if b.lowMem || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:170
			_go_fuzz_dep_.CoverTab[90935]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:170
			return cSize > maxCompressedBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:170
			// _ = "end of CoverTab[90935]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:170
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:170
			_go_fuzz_dep_.CoverTab[90936]++
														b.dataStorage = make([]byte, 0, cSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:171
			// _ = "end of CoverTab[90936]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:172
			_go_fuzz_dep_.CoverTab[90937]++
														b.dataStorage = make([]byte, 0, maxCompressedBlockSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:173
			// _ = "end of CoverTab[90937]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:174
		// _ = "end of CoverTab[90934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:175
		_go_fuzz_dep_.CoverTab[90938]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:175
		// _ = "end of CoverTab[90938]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:175
	// _ = "end of CoverTab[90907]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:175
	_go_fuzz_dep_.CoverTab[90908]++
												if cap(b.dst) <= maxSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:176
		_go_fuzz_dep_.CoverTab[90939]++
													b.dst = make([]byte, 0, maxSize+1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:177
		// _ = "end of CoverTab[90939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:178
		_go_fuzz_dep_.CoverTab[90940]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:178
		// _ = "end of CoverTab[90940]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:178
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:178
	// _ = "end of CoverTab[90908]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:178
	_go_fuzz_dep_.CoverTab[90909]++
												b.data, err = br.readBig(cSize, b.dataStorage)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:180
		_go_fuzz_dep_.CoverTab[90941]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:181
			_go_fuzz_dep_.CoverTab[90943]++
														println("Reading block:", err, "(", cSize, ")", len(b.data))
														printf("%T", br)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:183
			// _ = "end of CoverTab[90943]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:184
			_go_fuzz_dep_.CoverTab[90944]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:184
			// _ = "end of CoverTab[90944]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:184
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:184
		// _ = "end of CoverTab[90941]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:184
		_go_fuzz_dep_.CoverTab[90942]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:185
		// _ = "end of CoverTab[90942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:186
		_go_fuzz_dep_.CoverTab[90945]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:186
		// _ = "end of CoverTab[90945]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:186
	// _ = "end of CoverTab[90909]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:186
	_go_fuzz_dep_.CoverTab[90910]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:187
	// _ = "end of CoverTab[90910]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:191
func (b *blockDec) sendErr(err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:191
	_go_fuzz_dep_.CoverTab[90946]++
												b.Last = true
												b.Type = blockTypeReserved
												b.err = err
												b.input <- struct{}{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:195
	// _ = "end of CoverTab[90946]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:200
func (b *blockDec) Close() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:200
	_go_fuzz_dep_.CoverTab[90947]++
												close(b.input)
												close(b.history)
												close(b.result)
												b.decWG.Wait()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:204
	// _ = "end of CoverTab[90947]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:209
func (b *blockDec) startDecoder() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:209
	_go_fuzz_dep_.CoverTab[90948]++
												defer b.decWG.Done()
												for range b.input {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:211
		_go_fuzz_dep_.CoverTab[90949]++

													switch b.Type {
		case blockTypeRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:214
			_go_fuzz_dep_.CoverTab[90951]++
														if cap(b.dst) < int(b.RLESize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:215
				_go_fuzz_dep_.CoverTab[90959]++
															if b.lowMem {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:216
					_go_fuzz_dep_.CoverTab[90960]++
																b.dst = make([]byte, b.RLESize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:217
					// _ = "end of CoverTab[90960]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:218
					_go_fuzz_dep_.CoverTab[90961]++
																b.dst = make([]byte, maxBlockSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:219
					// _ = "end of CoverTab[90961]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:220
				// _ = "end of CoverTab[90959]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:221
				_go_fuzz_dep_.CoverTab[90962]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:221
				// _ = "end of CoverTab[90962]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:221
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:221
			// _ = "end of CoverTab[90951]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:221
			_go_fuzz_dep_.CoverTab[90952]++
														o := decodeOutput{
				d:	b,
				b:	b.dst[:b.RLESize],
				err:	nil,
			}
			v := b.data[0]
			for i := range o.b {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:228
				_go_fuzz_dep_.CoverTab[90963]++
															o.b[i] = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:229
				// _ = "end of CoverTab[90963]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:230
			// _ = "end of CoverTab[90952]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:230
			_go_fuzz_dep_.CoverTab[90953]++
														hist := <-b.history
														hist.append(o.b)
														b.result <- o
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:233
			// _ = "end of CoverTab[90953]"
		case blockTypeRaw:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:234
			_go_fuzz_dep_.CoverTab[90954]++
														o := decodeOutput{
				d:	b,
				b:	b.data,
				err:	nil,
			}
														hist := <-b.history
														hist.append(o.b)
														b.result <- o
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:242
			// _ = "end of CoverTab[90954]"
		case blockTypeCompressed:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:243
			_go_fuzz_dep_.CoverTab[90955]++
														b.dst = b.dst[:0]
														err := b.decodeCompressed(nil)
														o := decodeOutput{
				d:	b,
				b:	b.dst,
				err:	err,
			}
			if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:251
				_go_fuzz_dep_.CoverTab[90964]++
															println("Decompressed to", len(b.dst), "bytes, error:", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:252
				// _ = "end of CoverTab[90964]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:253
				_go_fuzz_dep_.CoverTab[90965]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:253
				// _ = "end of CoverTab[90965]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:253
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:253
			// _ = "end of CoverTab[90955]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:253
			_go_fuzz_dep_.CoverTab[90956]++
														b.result <- o
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:254
			// _ = "end of CoverTab[90956]"
		case blockTypeReserved:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:255
			_go_fuzz_dep_.CoverTab[90957]++

														<-b.history
														b.result <- decodeOutput{
				d:	b,
				b:	nil,
				err:	b.err,
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:262
			// _ = "end of CoverTab[90957]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:263
			_go_fuzz_dep_.CoverTab[90958]++
														panic("Invalid block type")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:264
			// _ = "end of CoverTab[90958]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:265
		// _ = "end of CoverTab[90949]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:265
		_go_fuzz_dep_.CoverTab[90950]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:266
			_go_fuzz_dep_.CoverTab[90966]++
														println("blockDec: Finished block")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:267
			// _ = "end of CoverTab[90966]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:268
			_go_fuzz_dep_.CoverTab[90967]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:268
			// _ = "end of CoverTab[90967]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:268
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:268
		// _ = "end of CoverTab[90950]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:269
	// _ = "end of CoverTab[90948]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:274
func (b *blockDec) decodeBuf(hist *history) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:274
	_go_fuzz_dep_.CoverTab[90968]++
												switch b.Type {
	case blockTypeRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:276
		_go_fuzz_dep_.CoverTab[90969]++
													if cap(b.dst) < int(b.RLESize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:277
			_go_fuzz_dep_.CoverTab[90977]++
														if b.lowMem {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:278
				_go_fuzz_dep_.CoverTab[90978]++
															b.dst = make([]byte, b.RLESize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:279
				// _ = "end of CoverTab[90978]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:280
				_go_fuzz_dep_.CoverTab[90979]++
															b.dst = make([]byte, maxBlockSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:281
				// _ = "end of CoverTab[90979]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:282
			// _ = "end of CoverTab[90977]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:283
			_go_fuzz_dep_.CoverTab[90980]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:283
			// _ = "end of CoverTab[90980]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:283
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:283
		// _ = "end of CoverTab[90969]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:283
		_go_fuzz_dep_.CoverTab[90970]++
													b.dst = b.dst[:b.RLESize]
													v := b.data[0]
													for i := range b.dst {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:286
			_go_fuzz_dep_.CoverTab[90981]++
														b.dst[i] = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:287
			// _ = "end of CoverTab[90981]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:288
		// _ = "end of CoverTab[90970]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:288
		_go_fuzz_dep_.CoverTab[90971]++
													hist.appendKeep(b.dst)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:290
		// _ = "end of CoverTab[90971]"
	case blockTypeRaw:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:291
		_go_fuzz_dep_.CoverTab[90972]++
													hist.appendKeep(b.data)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:293
		// _ = "end of CoverTab[90972]"
	case blockTypeCompressed:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:294
		_go_fuzz_dep_.CoverTab[90973]++
													saved := b.dst
													b.dst = hist.b
													hist.b = nil
													err := b.decodeCompressed(hist)
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:299
			_go_fuzz_dep_.CoverTab[90982]++
														println("Decompressed to total", len(b.dst), "bytes, hash:", xxhash.Sum64(b.dst), "error:", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:300
			// _ = "end of CoverTab[90982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:301
			_go_fuzz_dep_.CoverTab[90983]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:301
			// _ = "end of CoverTab[90983]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:301
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:301
		// _ = "end of CoverTab[90973]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:301
		_go_fuzz_dep_.CoverTab[90974]++
													hist.b = b.dst
													b.dst = saved
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:304
		// _ = "end of CoverTab[90974]"
	case blockTypeReserved:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:305
		_go_fuzz_dep_.CoverTab[90975]++

													return b.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:307
		// _ = "end of CoverTab[90975]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:308
		_go_fuzz_dep_.CoverTab[90976]++
													panic("Invalid block type")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:309
		// _ = "end of CoverTab[90976]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:310
	// _ = "end of CoverTab[90968]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:316
func (b *blockDec) decodeCompressed(hist *history) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:316
	_go_fuzz_dep_.CoverTab[90984]++
												in := b.data
												delayedHistory := hist == nil

												if delayedHistory {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:320
		_go_fuzz_dep_.CoverTab[91010]++

													defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:322
			_go_fuzz_dep_.CoverTab[91011]++
														if hist == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:323
				_go_fuzz_dep_.CoverTab[91012]++
															<-b.history
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:324
				// _ = "end of CoverTab[91012]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:325
				_go_fuzz_dep_.CoverTab[91013]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:325
				// _ = "end of CoverTab[91013]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:325
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:325
			// _ = "end of CoverTab[91011]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:326
		// _ = "end of CoverTab[91010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:327
		_go_fuzz_dep_.CoverTab[91014]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:327
		// _ = "end of CoverTab[91014]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:327
	// _ = "end of CoverTab[90984]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:327
	_go_fuzz_dep_.CoverTab[90985]++

												if len(in) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:329
		_go_fuzz_dep_.CoverTab[91015]++
													return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:330
		// _ = "end of CoverTab[91015]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:331
		_go_fuzz_dep_.CoverTab[91016]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:331
		// _ = "end of CoverTab[91016]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:331
	// _ = "end of CoverTab[90985]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:331
	_go_fuzz_dep_.CoverTab[90986]++
												litType := literalsBlockType(in[0] & 3)
												var litRegenSize int
												var litCompSize int
												sizeFormat := (in[0] >> 2) & 3
												var fourStreams bool
												switch litType {
	case literalsBlockRaw, literalsBlockRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:338
		_go_fuzz_dep_.CoverTab[91017]++
													switch sizeFormat {
		case 0, 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:340
			_go_fuzz_dep_.CoverTab[91020]++

														litRegenSize = int(in[0] >> 3)
														in = in[1:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:343
			// _ = "end of CoverTab[91020]"
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:344
			_go_fuzz_dep_.CoverTab[91021]++

														litRegenSize = int(in[0]>>4) + (int(in[1]) << 4)
														in = in[2:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:347
			// _ = "end of CoverTab[91021]"
		case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:348
			_go_fuzz_dep_.CoverTab[91022]++

														if len(in) < 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:350
				_go_fuzz_dep_.CoverTab[91025]++
															println("too small: litType:", litType, " sizeFormat", sizeFormat, len(in))
															return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:352
				// _ = "end of CoverTab[91025]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:353
				_go_fuzz_dep_.CoverTab[91026]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:353
				// _ = "end of CoverTab[91026]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:353
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:353
			// _ = "end of CoverTab[91022]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:353
			_go_fuzz_dep_.CoverTab[91023]++
														litRegenSize = int(in[0]>>4) + (int(in[1]) << 4) + (int(in[2]) << 12)
														in = in[3:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:355
			// _ = "end of CoverTab[91023]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:355
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:355
			_go_fuzz_dep_.CoverTab[91024]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:355
			// _ = "end of CoverTab[91024]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:356
		// _ = "end of CoverTab[91017]"
	case literalsBlockCompressed, literalsBlockTreeless:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:357
		_go_fuzz_dep_.CoverTab[91018]++
													switch sizeFormat {
		case 0, 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:359
			_go_fuzz_dep_.CoverTab[91027]++

														if len(in) < 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:361
				_go_fuzz_dep_.CoverTab[91034]++
															println("too small: litType:", litType, " sizeFormat", sizeFormat, len(in))
															return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:363
				// _ = "end of CoverTab[91034]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:364
				_go_fuzz_dep_.CoverTab[91035]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:364
				// _ = "end of CoverTab[91035]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:364
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:364
			// _ = "end of CoverTab[91027]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:364
			_go_fuzz_dep_.CoverTab[91028]++
														n := uint64(in[0]>>4) + (uint64(in[1]) << 4) + (uint64(in[2]) << 12)
														litRegenSize = int(n & 1023)
														litCompSize = int(n >> 10)
														fourStreams = sizeFormat == 1
														in = in[3:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:369
			// _ = "end of CoverTab[91028]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:370
			_go_fuzz_dep_.CoverTab[91029]++
														fourStreams = true
														if len(in) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:372
				_go_fuzz_dep_.CoverTab[91036]++
															println("too small: litType:", litType, " sizeFormat", sizeFormat, len(in))
															return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:374
				// _ = "end of CoverTab[91036]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:375
				_go_fuzz_dep_.CoverTab[91037]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:375
				// _ = "end of CoverTab[91037]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:375
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:375
			// _ = "end of CoverTab[91029]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:375
			_go_fuzz_dep_.CoverTab[91030]++
														n := uint64(in[0]>>4) + (uint64(in[1]) << 4) + (uint64(in[2]) << 12) + (uint64(in[3]) << 20)
														litRegenSize = int(n & 16383)
														litCompSize = int(n >> 14)
														in = in[4:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:379
			// _ = "end of CoverTab[91030]"
		case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:380
			_go_fuzz_dep_.CoverTab[91031]++
														fourStreams = true
														if len(in) < 5 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:382
				_go_fuzz_dep_.CoverTab[91038]++
															println("too small: litType:", litType, " sizeFormat", sizeFormat, len(in))
															return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:384
				// _ = "end of CoverTab[91038]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:385
				_go_fuzz_dep_.CoverTab[91039]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:385
				// _ = "end of CoverTab[91039]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:385
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:385
			// _ = "end of CoverTab[91031]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:385
			_go_fuzz_dep_.CoverTab[91032]++
														n := uint64(in[0]>>4) + (uint64(in[1]) << 4) + (uint64(in[2]) << 12) + (uint64(in[3]) << 20) + (uint64(in[4]) << 28)
														litRegenSize = int(n & 262143)
														litCompSize = int(n >> 18)
														in = in[5:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:389
			// _ = "end of CoverTab[91032]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:389
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:389
			_go_fuzz_dep_.CoverTab[91033]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:389
			// _ = "end of CoverTab[91033]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:390
		// _ = "end of CoverTab[91018]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:390
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:390
		_go_fuzz_dep_.CoverTab[91019]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:390
		// _ = "end of CoverTab[91019]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:391
	// _ = "end of CoverTab[90986]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:391
	_go_fuzz_dep_.CoverTab[90987]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:392
		_go_fuzz_dep_.CoverTab[91040]++
													println("literals type:", litType, "litRegenSize:", litRegenSize, "litCompSize:", litCompSize, "sizeFormat:", sizeFormat, "4X:", fourStreams)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:393
		// _ = "end of CoverTab[91040]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:394
		_go_fuzz_dep_.CoverTab[91041]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:394
		// _ = "end of CoverTab[91041]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:394
	// _ = "end of CoverTab[90987]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:394
	_go_fuzz_dep_.CoverTab[90988]++
												var literals []byte
												var huff *huff0.Scratch
												switch litType {
	case literalsBlockRaw:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:398
		_go_fuzz_dep_.CoverTab[91042]++
													if len(in) < litRegenSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:399
			_go_fuzz_dep_.CoverTab[91059]++
														println("too small: litType:", litType, " sizeFormat", sizeFormat, "remain:", len(in), "want:", litRegenSize)
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:401
			// _ = "end of CoverTab[91059]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:402
			_go_fuzz_dep_.CoverTab[91060]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:402
			// _ = "end of CoverTab[91060]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:402
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:402
		// _ = "end of CoverTab[91042]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:402
		_go_fuzz_dep_.CoverTab[91043]++
													literals = in[:litRegenSize]
													in = in[litRegenSize:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:404
		// _ = "end of CoverTab[91043]"

	case literalsBlockRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:406
		_go_fuzz_dep_.CoverTab[91044]++
													if len(in) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:407
			_go_fuzz_dep_.CoverTab[91061]++
														println("too small: litType:", litType, " sizeFormat", sizeFormat, "remain:", len(in), "want:", 1)
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:409
			// _ = "end of CoverTab[91061]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:410
			_go_fuzz_dep_.CoverTab[91062]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:410
			// _ = "end of CoverTab[91062]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:410
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:410
		// _ = "end of CoverTab[91044]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:410
		_go_fuzz_dep_.CoverTab[91045]++
													if cap(b.literalBuf) < litRegenSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:411
			_go_fuzz_dep_.CoverTab[91063]++
														if b.lowMem {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:412
				_go_fuzz_dep_.CoverTab[91064]++
															b.literalBuf = make([]byte, litRegenSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:413
				// _ = "end of CoverTab[91064]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:414
				_go_fuzz_dep_.CoverTab[91065]++
															if litRegenSize > maxCompressedLiteralSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:415
					_go_fuzz_dep_.CoverTab[91066]++

																b.literalBuf = make([]byte, litRegenSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:417
					// _ = "end of CoverTab[91066]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:418
					_go_fuzz_dep_.CoverTab[91067]++
																b.literalBuf = make([]byte, litRegenSize, maxCompressedLiteralSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:419
					// _ = "end of CoverTab[91067]"

				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:421
				// _ = "end of CoverTab[91065]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:422
			// _ = "end of CoverTab[91063]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:423
			_go_fuzz_dep_.CoverTab[91068]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:423
			// _ = "end of CoverTab[91068]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:423
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:423
		// _ = "end of CoverTab[91045]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:423
		_go_fuzz_dep_.CoverTab[91046]++
													literals = b.literalBuf[:litRegenSize]
													v := in[0]
													for i := range literals {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:426
			_go_fuzz_dep_.CoverTab[91069]++
														literals[i] = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:427
			// _ = "end of CoverTab[91069]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:428
		// _ = "end of CoverTab[91046]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:428
		_go_fuzz_dep_.CoverTab[91047]++
													in = in[1:]
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:430
			_go_fuzz_dep_.CoverTab[91070]++
														printf("Found %d RLE compressed literals\n", litRegenSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:431
			// _ = "end of CoverTab[91070]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:432
			_go_fuzz_dep_.CoverTab[91071]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:432
			// _ = "end of CoverTab[91071]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:432
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:432
		// _ = "end of CoverTab[91047]"
	case literalsBlockTreeless:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:433
		_go_fuzz_dep_.CoverTab[91048]++
													if len(in) < litCompSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:434
			_go_fuzz_dep_.CoverTab[91072]++
														println("too small: litType:", litType, " sizeFormat", sizeFormat, "remain:", len(in), "want:", litCompSize)
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:436
			// _ = "end of CoverTab[91072]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:437
			_go_fuzz_dep_.CoverTab[91073]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:437
			// _ = "end of CoverTab[91073]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:437
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:437
		// _ = "end of CoverTab[91048]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:437
		_go_fuzz_dep_.CoverTab[91049]++

													literals = in[:litCompSize]
													in = in[litCompSize:]
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:441
			_go_fuzz_dep_.CoverTab[91074]++
														printf("Found %d compressed literals\n", litCompSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:442
			// _ = "end of CoverTab[91074]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:443
			_go_fuzz_dep_.CoverTab[91075]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:443
			// _ = "end of CoverTab[91075]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:443
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:443
		// _ = "end of CoverTab[91049]"
	case literalsBlockCompressed:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:444
		_go_fuzz_dep_.CoverTab[91050]++
													if len(in) < litCompSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:445
			_go_fuzz_dep_.CoverTab[91076]++
														println("too small: litType:", litType, " sizeFormat", sizeFormat, "remain:", len(in), "want:", litCompSize)
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:447
			// _ = "end of CoverTab[91076]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:448
			_go_fuzz_dep_.CoverTab[91077]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:448
			// _ = "end of CoverTab[91077]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:448
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:448
		// _ = "end of CoverTab[91050]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:448
		_go_fuzz_dep_.CoverTab[91051]++
													literals = in[:litCompSize]
													in = in[litCompSize:]
													huff = huffDecoderPool.Get().(*huff0.Scratch)
													var err error

													if cap(b.literalBuf) < litRegenSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:454
			_go_fuzz_dep_.CoverTab[91078]++
														if b.lowMem {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:455
				_go_fuzz_dep_.CoverTab[91079]++
															b.literalBuf = make([]byte, 0, litRegenSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:456
				// _ = "end of CoverTab[91079]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:457
				_go_fuzz_dep_.CoverTab[91080]++
															b.literalBuf = make([]byte, 0, maxCompressedLiteralSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:458
				// _ = "end of CoverTab[91080]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:459
			// _ = "end of CoverTab[91078]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:460
			_go_fuzz_dep_.CoverTab[91081]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:460
			// _ = "end of CoverTab[91081]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:460
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:460
		// _ = "end of CoverTab[91051]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:460
		_go_fuzz_dep_.CoverTab[91052]++
													if huff == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:461
			_go_fuzz_dep_.CoverTab[91082]++
														huff = &huff0.Scratch{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:462
			// _ = "end of CoverTab[91082]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:463
			_go_fuzz_dep_.CoverTab[91083]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:463
			// _ = "end of CoverTab[91083]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:463
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:463
		// _ = "end of CoverTab[91052]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:463
		_go_fuzz_dep_.CoverTab[91053]++
													huff, literals, err = huff0.ReadTable(literals, huff)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:465
			_go_fuzz_dep_.CoverTab[91084]++
														println("reading huffman table:", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:467
			// _ = "end of CoverTab[91084]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:468
			_go_fuzz_dep_.CoverTab[91085]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:468
			// _ = "end of CoverTab[91085]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:468
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:468
		// _ = "end of CoverTab[91053]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:468
		_go_fuzz_dep_.CoverTab[91054]++

													if fourStreams {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:470
			_go_fuzz_dep_.CoverTab[91086]++
														literals, err = huff.Decoder().Decompress4X(b.literalBuf[:0:litRegenSize], literals)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:471
			// _ = "end of CoverTab[91086]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:472
			_go_fuzz_dep_.CoverTab[91087]++
														literals, err = huff.Decoder().Decompress1X(b.literalBuf[:0:litRegenSize], literals)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:473
			// _ = "end of CoverTab[91087]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:474
		// _ = "end of CoverTab[91054]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:474
		_go_fuzz_dep_.CoverTab[91055]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:475
			_go_fuzz_dep_.CoverTab[91088]++
														println("decoding compressed literals:", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:477
			// _ = "end of CoverTab[91088]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:478
			_go_fuzz_dep_.CoverTab[91089]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:478
			// _ = "end of CoverTab[91089]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:478
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:478
		// _ = "end of CoverTab[91055]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:478
		_go_fuzz_dep_.CoverTab[91056]++

													if len(literals) != litRegenSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:480
			_go_fuzz_dep_.CoverTab[91090]++
														return fmt.Errorf("literal output size mismatch want %d, got %d", litRegenSize, len(literals))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:481
			// _ = "end of CoverTab[91090]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:482
			_go_fuzz_dep_.CoverTab[91091]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:482
			// _ = "end of CoverTab[91091]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:482
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:482
		// _ = "end of CoverTab[91056]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:482
		_go_fuzz_dep_.CoverTab[91057]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:483
			_go_fuzz_dep_.CoverTab[91092]++
														printf("Decompressed %d literals into %d bytes\n", litCompSize, litRegenSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:484
			// _ = "end of CoverTab[91092]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
			_go_fuzz_dep_.CoverTab[91093]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
			// _ = "end of CoverTab[91093]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
		// _ = "end of CoverTab[91057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
		_go_fuzz_dep_.CoverTab[91058]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:485
		// _ = "end of CoverTab[91058]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:486
	// _ = "end of CoverTab[90988]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:486
	_go_fuzz_dep_.CoverTab[90989]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:490
	if len(in) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:490
		_go_fuzz_dep_.CoverTab[91094]++
													return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:491
		// _ = "end of CoverTab[91094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:492
		_go_fuzz_dep_.CoverTab[91095]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:492
		// _ = "end of CoverTab[91095]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:492
	// _ = "end of CoverTab[90989]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:492
	_go_fuzz_dep_.CoverTab[90990]++
												seqHeader := in[0]
												nSeqs := 0
												switch {
	case seqHeader == 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:496
		_go_fuzz_dep_.CoverTab[91096]++
													in = in[1:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:497
		// _ = "end of CoverTab[91096]"
	case seqHeader < 128:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:498
		_go_fuzz_dep_.CoverTab[91097]++
													nSeqs = int(seqHeader)
													in = in[1:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:500
		// _ = "end of CoverTab[91097]"
	case seqHeader < 255:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:501
		_go_fuzz_dep_.CoverTab[91098]++
													if len(in) < 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:502
			_go_fuzz_dep_.CoverTab[91103]++
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:503
			// _ = "end of CoverTab[91103]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:504
			_go_fuzz_dep_.CoverTab[91104]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:504
			// _ = "end of CoverTab[91104]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:504
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:504
		// _ = "end of CoverTab[91098]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:504
		_go_fuzz_dep_.CoverTab[91099]++
													nSeqs = int(seqHeader-128)<<8 | int(in[1])
													in = in[2:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:506
		// _ = "end of CoverTab[91099]"
	case seqHeader == 255:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:507
		_go_fuzz_dep_.CoverTab[91100]++
													if len(in) < 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:508
			_go_fuzz_dep_.CoverTab[91105]++
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:509
			// _ = "end of CoverTab[91105]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:510
			_go_fuzz_dep_.CoverTab[91106]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:510
			// _ = "end of CoverTab[91106]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:510
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:510
		// _ = "end of CoverTab[91100]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:510
		_go_fuzz_dep_.CoverTab[91101]++
													nSeqs = 0x7f00 + int(in[1]) + (int(in[2]) << 8)
													in = in[3:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:512
		// _ = "end of CoverTab[91101]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:512
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:512
		_go_fuzz_dep_.CoverTab[91102]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:512
		// _ = "end of CoverTab[91102]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:513
	// _ = "end of CoverTab[90990]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:513
	_go_fuzz_dep_.CoverTab[90991]++

												var seqs = &sequenceDecs{}
												if nSeqs > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:516
		_go_fuzz_dep_.CoverTab[91107]++
													if len(in) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:517
			_go_fuzz_dep_.CoverTab[91111]++
														return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:518
			// _ = "end of CoverTab[91111]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:519
			_go_fuzz_dep_.CoverTab[91112]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:519
			// _ = "end of CoverTab[91112]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:519
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:519
		// _ = "end of CoverTab[91107]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:519
		_go_fuzz_dep_.CoverTab[91108]++
													br := byteReader{b: in, off: 0}
													compMode := br.Uint8()
													br.advance(1)
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:523
			_go_fuzz_dep_.CoverTab[91113]++
														printf("Compression modes: 0b%b", compMode)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:524
			// _ = "end of CoverTab[91113]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:525
			_go_fuzz_dep_.CoverTab[91114]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:525
			// _ = "end of CoverTab[91114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:525
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:525
		// _ = "end of CoverTab[91108]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:525
		_go_fuzz_dep_.CoverTab[91109]++
													for i := uint(0); i < 3; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:526
			_go_fuzz_dep_.CoverTab[91115]++
														mode := seqCompMode((compMode >> (6 - i*2)) & 3)
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:528
				_go_fuzz_dep_.CoverTab[91119]++
															println("Table", tableIndex(i), "is", mode)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:529
				// _ = "end of CoverTab[91119]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:530
				_go_fuzz_dep_.CoverTab[91120]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:530
				// _ = "end of CoverTab[91120]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:530
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:530
			// _ = "end of CoverTab[91115]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:530
			_go_fuzz_dep_.CoverTab[91116]++
														var seq *sequenceDec
														switch tableIndex(i) {
			case tableLiteralLengths:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:533
				_go_fuzz_dep_.CoverTab[91121]++
															seq = &seqs.litLengths
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:534
				// _ = "end of CoverTab[91121]"
			case tableOffsets:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:535
				_go_fuzz_dep_.CoverTab[91122]++
															seq = &seqs.offsets
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:536
				// _ = "end of CoverTab[91122]"
			case tableMatchLengths:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:537
				_go_fuzz_dep_.CoverTab[91123]++
															seq = &seqs.matchLengths
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:538
				// _ = "end of CoverTab[91123]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:539
				_go_fuzz_dep_.CoverTab[91124]++
															panic("unknown table")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:540
				// _ = "end of CoverTab[91124]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:541
			// _ = "end of CoverTab[91116]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:541
			_go_fuzz_dep_.CoverTab[91117]++
														switch mode {
			case compModePredefined:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:543
				_go_fuzz_dep_.CoverTab[91125]++
															seq.fse = &fsePredef[i]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:544
				// _ = "end of CoverTab[91125]"
			case compModeRLE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:545
				_go_fuzz_dep_.CoverTab[91126]++
															if br.remain() < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:546
					_go_fuzz_dep_.CoverTab[91135]++
																return ErrBlockTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:547
					// _ = "end of CoverTab[91135]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:548
					_go_fuzz_dep_.CoverTab[91136]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:548
					// _ = "end of CoverTab[91136]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:548
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:548
				// _ = "end of CoverTab[91126]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:548
				_go_fuzz_dep_.CoverTab[91127]++
															v := br.Uint8()
															br.advance(1)
															dec := fseDecoderPool.Get().(*fseDecoder)
															symb, err := decSymbolValue(v, symbolTableX[i])
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:553
					_go_fuzz_dep_.CoverTab[91137]++
																printf("RLE Transform table (%v) error: %v", tableIndex(i), err)
																return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:555
					// _ = "end of CoverTab[91137]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:556
					_go_fuzz_dep_.CoverTab[91138]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:556
					// _ = "end of CoverTab[91138]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:556
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:556
				// _ = "end of CoverTab[91127]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:556
				_go_fuzz_dep_.CoverTab[91128]++
															dec.setRLE(symb)
															seq.fse = dec
															if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:559
					_go_fuzz_dep_.CoverTab[91139]++
																printf("RLE set to %+v, code: %v", symb, v)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:560
					// _ = "end of CoverTab[91139]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:561
					_go_fuzz_dep_.CoverTab[91140]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:561
					// _ = "end of CoverTab[91140]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:561
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:561
				// _ = "end of CoverTab[91128]"
			case compModeFSE:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:562
				_go_fuzz_dep_.CoverTab[91129]++
															println("Reading table for", tableIndex(i))
															dec := fseDecoderPool.Get().(*fseDecoder)
															err := dec.readNCount(&br, uint16(maxTableSymbol[i]))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:566
					_go_fuzz_dep_.CoverTab[91141]++
																println("Read table error:", err)
																return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:568
					// _ = "end of CoverTab[91141]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:569
					_go_fuzz_dep_.CoverTab[91142]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:569
					// _ = "end of CoverTab[91142]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:569
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:569
				// _ = "end of CoverTab[91129]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:569
				_go_fuzz_dep_.CoverTab[91130]++
															err = dec.transform(symbolTableX[i])
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:571
					_go_fuzz_dep_.CoverTab[91143]++
																println("Transform table error:", err)
																return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:573
					// _ = "end of CoverTab[91143]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:574
					_go_fuzz_dep_.CoverTab[91144]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:574
					// _ = "end of CoverTab[91144]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:574
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:574
				// _ = "end of CoverTab[91130]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:574
				_go_fuzz_dep_.CoverTab[91131]++
															if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:575
					_go_fuzz_dep_.CoverTab[91145]++
																println("Read table ok", "symbolLen:", dec.symbolLen)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:576
					// _ = "end of CoverTab[91145]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:577
					_go_fuzz_dep_.CoverTab[91146]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:577
					// _ = "end of CoverTab[91146]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:577
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:577
				// _ = "end of CoverTab[91131]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:577
				_go_fuzz_dep_.CoverTab[91132]++
															seq.fse = dec
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:578
				// _ = "end of CoverTab[91132]"
			case compModeRepeat:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:579
				_go_fuzz_dep_.CoverTab[91133]++
															seq.repeat = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:580
				// _ = "end of CoverTab[91133]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:580
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:580
				_go_fuzz_dep_.CoverTab[91134]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:580
				// _ = "end of CoverTab[91134]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:581
			// _ = "end of CoverTab[91117]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:581
			_go_fuzz_dep_.CoverTab[91118]++
														if br.overread() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:582
				_go_fuzz_dep_.CoverTab[91147]++
															return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:583
				// _ = "end of CoverTab[91147]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:584
				_go_fuzz_dep_.CoverTab[91148]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:584
				// _ = "end of CoverTab[91148]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:584
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:584
			// _ = "end of CoverTab[91118]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:585
		// _ = "end of CoverTab[91109]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:585
		_go_fuzz_dep_.CoverTab[91110]++
													in = br.unread()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:586
		// _ = "end of CoverTab[91110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:587
		_go_fuzz_dep_.CoverTab[91149]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:587
		// _ = "end of CoverTab[91149]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:587
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:587
	// _ = "end of CoverTab[90991]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:587
	_go_fuzz_dep_.CoverTab[90992]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:591
	if hist == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:591
		_go_fuzz_dep_.CoverTab[91150]++
													hist = <-b.history
													if hist.error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:593
			_go_fuzz_dep_.CoverTab[91151]++
														return ErrDecoderClosed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:594
			// _ = "end of CoverTab[91151]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:595
			_go_fuzz_dep_.CoverTab[91152]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:595
			// _ = "end of CoverTab[91152]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:595
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:595
		// _ = "end of CoverTab[91150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:596
		_go_fuzz_dep_.CoverTab[91153]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:596
		// _ = "end of CoverTab[91153]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:596
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:596
	// _ = "end of CoverTab[90992]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:596
	_go_fuzz_dep_.CoverTab[90993]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:599
	if litType == literalsBlockTreeless {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:599
		_go_fuzz_dep_.CoverTab[91154]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:605
		if hist.huffTree == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:605
			_go_fuzz_dep_.CoverTab[91159]++
														return errors.New("literal block was treeless, but no history was defined")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:606
			// _ = "end of CoverTab[91159]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:607
			_go_fuzz_dep_.CoverTab[91160]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:607
			// _ = "end of CoverTab[91160]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:607
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:607
		// _ = "end of CoverTab[91154]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:607
		_go_fuzz_dep_.CoverTab[91155]++

													if cap(b.literalBuf) < litRegenSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:609
			_go_fuzz_dep_.CoverTab[91161]++
														if b.lowMem {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:610
				_go_fuzz_dep_.CoverTab[91162]++
															b.literalBuf = make([]byte, 0, litRegenSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:611
				// _ = "end of CoverTab[91162]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:612
				_go_fuzz_dep_.CoverTab[91163]++
															b.literalBuf = make([]byte, 0, maxCompressedLiteralSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:613
				// _ = "end of CoverTab[91163]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:614
			// _ = "end of CoverTab[91161]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:615
			_go_fuzz_dep_.CoverTab[91164]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:615
			// _ = "end of CoverTab[91164]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:615
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:615
		// _ = "end of CoverTab[91155]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:615
		_go_fuzz_dep_.CoverTab[91156]++
													var err error

													huff = hist.huffTree
													if fourStreams {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:619
			_go_fuzz_dep_.CoverTab[91165]++
														literals, err = huff.Decoder().Decompress4X(b.literalBuf[:0:litRegenSize], literals)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:620
			// _ = "end of CoverTab[91165]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:621
			_go_fuzz_dep_.CoverTab[91166]++
														literals, err = huff.Decoder().Decompress1X(b.literalBuf[:0:litRegenSize], literals)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:622
			// _ = "end of CoverTab[91166]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:623
		// _ = "end of CoverTab[91156]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:623
		_go_fuzz_dep_.CoverTab[91157]++

													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:625
			_go_fuzz_dep_.CoverTab[91167]++
														println("decompressing literals:", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:627
			// _ = "end of CoverTab[91167]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:628
			_go_fuzz_dep_.CoverTab[91168]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:628
			// _ = "end of CoverTab[91168]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:628
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:628
		// _ = "end of CoverTab[91157]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:628
		_go_fuzz_dep_.CoverTab[91158]++
													if len(literals) != litRegenSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:629
			_go_fuzz_dep_.CoverTab[91169]++
														return fmt.Errorf("literal output size mismatch want %d, got %d", litRegenSize, len(literals))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:630
			// _ = "end of CoverTab[91169]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:631
			_go_fuzz_dep_.CoverTab[91170]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:631
			// _ = "end of CoverTab[91170]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:631
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:631
		// _ = "end of CoverTab[91158]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:632
		_go_fuzz_dep_.CoverTab[91171]++
													if hist.huffTree != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:633
			_go_fuzz_dep_.CoverTab[91172]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:633
			return huff != nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:633
			// _ = "end of CoverTab[91172]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:633
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:633
			_go_fuzz_dep_.CoverTab[91173]++
														if hist.dict == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:634
				_go_fuzz_dep_.CoverTab[91175]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:634
				return hist.dict.litEnc != hist.huffTree
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:634
				// _ = "end of CoverTab[91175]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:634
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:634
				_go_fuzz_dep_.CoverTab[91176]++
															huffDecoderPool.Put(hist.huffTree)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:635
				// _ = "end of CoverTab[91176]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:636
				_go_fuzz_dep_.CoverTab[91177]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:636
				// _ = "end of CoverTab[91177]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:636
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:636
			// _ = "end of CoverTab[91173]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:636
			_go_fuzz_dep_.CoverTab[91174]++
														hist.huffTree = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:637
			// _ = "end of CoverTab[91174]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:638
			_go_fuzz_dep_.CoverTab[91178]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:638
			// _ = "end of CoverTab[91178]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:638
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:638
		// _ = "end of CoverTab[91171]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:639
	// _ = "end of CoverTab[90993]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:639
	_go_fuzz_dep_.CoverTab[90994]++
												if huff != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:640
		_go_fuzz_dep_.CoverTab[91179]++
													hist.huffTree = huff
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:641
		// _ = "end of CoverTab[91179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:642
		_go_fuzz_dep_.CoverTab[91180]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:642
		// _ = "end of CoverTab[91180]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:642
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:642
	// _ = "end of CoverTab[90994]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:642
	_go_fuzz_dep_.CoverTab[90995]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:643
		_go_fuzz_dep_.CoverTab[91181]++
													println("Final literals:", len(literals), "hash:", xxhash.Sum64(literals), "and", nSeqs, "sequences.")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:644
		// _ = "end of CoverTab[91181]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:645
		_go_fuzz_dep_.CoverTab[91182]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:645
		// _ = "end of CoverTab[91182]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:645
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:645
	// _ = "end of CoverTab[90995]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:645
	_go_fuzz_dep_.CoverTab[90996]++

												if nSeqs == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:647
		_go_fuzz_dep_.CoverTab[91183]++

													b.dst = append(b.dst, literals...)
													if delayedHistory {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:650
			_go_fuzz_dep_.CoverTab[91185]++
														hist.append(literals)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:651
			// _ = "end of CoverTab[91185]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:652
			_go_fuzz_dep_.CoverTab[91186]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:652
			// _ = "end of CoverTab[91186]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:652
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:652
		// _ = "end of CoverTab[91183]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:652
		_go_fuzz_dep_.CoverTab[91184]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:653
		// _ = "end of CoverTab[91184]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:654
		_go_fuzz_dep_.CoverTab[91187]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:654
		// _ = "end of CoverTab[91187]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:654
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:654
	// _ = "end of CoverTab[90996]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:654
	_go_fuzz_dep_.CoverTab[90997]++

												seqs, err := seqs.mergeHistory(&hist.decoders)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:657
		_go_fuzz_dep_.CoverTab[91188]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:658
		// _ = "end of CoverTab[91188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:659
		_go_fuzz_dep_.CoverTab[91189]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:659
		// _ = "end of CoverTab[91189]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:659
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:659
	// _ = "end of CoverTab[90997]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:659
	_go_fuzz_dep_.CoverTab[90998]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:660
		_go_fuzz_dep_.CoverTab[91190]++
													println("History merged ok")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:661
		// _ = "end of CoverTab[91190]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:662
		_go_fuzz_dep_.CoverTab[91191]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:662
		// _ = "end of CoverTab[91191]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:662
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:662
	// _ = "end of CoverTab[90998]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:662
	_go_fuzz_dep_.CoverTab[90999]++
												br := &bitReader{}
												if err := br.init(in); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:664
		_go_fuzz_dep_.CoverTab[91192]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:665
		// _ = "end of CoverTab[91192]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:666
		_go_fuzz_dep_.CoverTab[91193]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:666
		// _ = "end of CoverTab[91193]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:666
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:666
	// _ = "end of CoverTab[90999]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:666
	_go_fuzz_dep_.CoverTab[91000]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:673
	hbytes := hist.b
	if len(hbytes) > hist.windowSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:674
		_go_fuzz_dep_.CoverTab[91194]++
													hbytes = hbytes[len(hbytes)-hist.windowSize:]

													if hist.dict != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:677
			_go_fuzz_dep_.CoverTab[91195]++
														hist.dict.content = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:678
			// _ = "end of CoverTab[91195]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:679
			_go_fuzz_dep_.CoverTab[91196]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:679
			// _ = "end of CoverTab[91196]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:679
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:679
		// _ = "end of CoverTab[91194]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:680
		_go_fuzz_dep_.CoverTab[91197]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:680
		// _ = "end of CoverTab[91197]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:680
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:680
	// _ = "end of CoverTab[91000]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:680
	_go_fuzz_dep_.CoverTab[91001]++

												if err := seqs.initialize(br, hist, literals, b.dst); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:682
		_go_fuzz_dep_.CoverTab[91198]++
													println("initializing sequences:", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:684
		// _ = "end of CoverTab[91198]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:685
		_go_fuzz_dep_.CoverTab[91199]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:685
		// _ = "end of CoverTab[91199]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:685
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:685
	// _ = "end of CoverTab[91001]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:685
	_go_fuzz_dep_.CoverTab[91002]++

												err = seqs.decode(nSeqs, br, hbytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:688
		_go_fuzz_dep_.CoverTab[91200]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:689
		// _ = "end of CoverTab[91200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:690
		_go_fuzz_dep_.CoverTab[91201]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:690
		// _ = "end of CoverTab[91201]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:690
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:690
	// _ = "end of CoverTab[91002]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:690
	_go_fuzz_dep_.CoverTab[91003]++
												if !br.finished() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:691
		_go_fuzz_dep_.CoverTab[91202]++
													return fmt.Errorf("%d extra bits on block, should be 0", br.remain())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:692
		// _ = "end of CoverTab[91202]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:693
		_go_fuzz_dep_.CoverTab[91203]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:693
		// _ = "end of CoverTab[91203]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:693
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:693
	// _ = "end of CoverTab[91003]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:693
	_go_fuzz_dep_.CoverTab[91004]++

												err = br.close()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:696
		_go_fuzz_dep_.CoverTab[91204]++
													printf("Closing sequences: %v, %+v\n", err, *br)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:697
		// _ = "end of CoverTab[91204]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:698
		_go_fuzz_dep_.CoverTab[91205]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:698
		// _ = "end of CoverTab[91205]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:698
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:698
	// _ = "end of CoverTab[91004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:698
	_go_fuzz_dep_.CoverTab[91005]++
												if len(b.data) > maxCompressedBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:699
		_go_fuzz_dep_.CoverTab[91206]++
													return fmt.Errorf("compressed block size too large (%d)", len(b.data))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:700
		// _ = "end of CoverTab[91206]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:701
		_go_fuzz_dep_.CoverTab[91207]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:701
		// _ = "end of CoverTab[91207]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:701
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:701
	// _ = "end of CoverTab[91005]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:701
	_go_fuzz_dep_.CoverTab[91006]++

												b.dst = seqs.out
												seqs.out, seqs.literals, seqs.hist = nil, nil, nil

												if !delayedHistory {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:706
		_go_fuzz_dep_.CoverTab[91208]++

													hist.recentOffsets = seqs.prevOffset
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:709
		// _ = "end of CoverTab[91208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:710
		_go_fuzz_dep_.CoverTab[91209]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:710
		// _ = "end of CoverTab[91209]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:710
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:710
	// _ = "end of CoverTab[91006]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:710
	_go_fuzz_dep_.CoverTab[91007]++
												if b.Last {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:711
		_go_fuzz_dep_.CoverTab[91210]++

													println("Last block, no history returned")
													hist.b = hist.b[:0]
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:715
		// _ = "end of CoverTab[91210]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:716
		_go_fuzz_dep_.CoverTab[91211]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:716
		// _ = "end of CoverTab[91211]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:716
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:716
	// _ = "end of CoverTab[91007]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:716
	_go_fuzz_dep_.CoverTab[91008]++
												hist.append(b.dst)
												hist.recentOffsets = seqs.prevOffset
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:719
		_go_fuzz_dep_.CoverTab[91212]++
													println("Finished block with literals:", len(literals), "and", nSeqs, "sequences.")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:720
		// _ = "end of CoverTab[91212]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:721
		_go_fuzz_dep_.CoverTab[91213]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:721
		// _ = "end of CoverTab[91213]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:721
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:721
	// _ = "end of CoverTab[91008]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:721
	_go_fuzz_dep_.CoverTab[91009]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:723
	// _ = "end of CoverTab[91009]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:724
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/blockdec.go:724
var _ = _go_fuzz_dep_.CoverTab
