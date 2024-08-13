//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:1
// Package huff0 provides fast huffman encoding as used in zstd.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:1
// See README.md at https://github.com/klauspost/compress/tree/master/huff0 for details.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
package huff0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:4
)

import (
	"errors"
	"fmt"
	"math"
	"math/bits"

	"github.com/klauspost/compress/fse"
)

const (
	maxSymbolValue	= 255

	// zstandard limits tablelog to 11, see:
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#huffman-tree-description
	tableLogMax	= 11
	tableLogDefault	= 11
	minTablelog	= 5
	huffNodesLen	= 512

	// BlockSizeMax is maximum input size for a single block uncompressed.
	BlockSizeMax	= 1<<18 - 1
)

var (
	// ErrIncompressible is returned when input is judged to be too hard to compress.
	ErrIncompressible	= errors.New("input is not compressible")

	// ErrUseRLE is returned from the compressor when the input is a single byte value repeated.
	ErrUseRLE	= errors.New("input is single value repeated")

	// ErrTooBig is return if input is too large for a single block.
	ErrTooBig	= errors.New("input too big")

	// ErrMaxDecodedSizeExceeded is return if input is too large for a single block.
	ErrMaxDecodedSizeExceeded	= errors.New("maximum output size exceeded")
)

type ReusePolicy uint8

const (
	// ReusePolicyAllow will allow reuse if it produces smaller output.
	ReusePolicyAllow	ReusePolicy	= iota

	// ReusePolicyPrefer will re-use aggressively if possible.
	// This will not check if a new table will produce smaller output,
	// except if the current table is impossible to use or
	// compressed output is bigger than input.
	ReusePolicyPrefer

	// ReusePolicyNone will disable re-use of tables.
	// This is slightly faster than ReusePolicyAllow but may produce larger output.
	ReusePolicyNone

	// ReusePolicyMust must allow reuse and produce smaller output.
	ReusePolicyMust
)

type Scratch struct {
	count	[maxSymbolValue + 1]uint32

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:70
	// Out is output buffer.
	// If the scratch is re-used before the caller is done processing the output,
	// set this field to nil.
	// Otherwise the output buffer will be re-used for next Compression/Decompression step
	// and allocation will be avoided.
	Out	[]byte

	// OutTable will contain the table data only, if a new table has been generated.
	// Slice of the returned data.
	OutTable	[]byte

	// OutData will contain the compressed data.
	// Slice of the returned data.
	OutData	[]byte

	// MaxDecodedSize will set the maximum allowed output size.
	// This value will automatically be set to BlockSizeMax if not set.
	// Decoders will return ErrMaxDecodedSizeExceeded is this limit is exceeded.
	MaxDecodedSize	int

	br	byteReader

	// MaxSymbolValue will override the maximum symbol value of the next block.
	MaxSymbolValue	uint8

	// TableLog will attempt to override the tablelog for the next block.
	// Must be <= 11 and >= 5.
	TableLog	uint8

	// Reuse will specify the reuse policy
	Reuse	ReusePolicy

	// WantLogLess allows to specify a log 2 reduction that should at least be achieved,
	// otherwise the block will be returned as incompressible.
	// The reduction should then at least be (input size >> WantLogLess)
	// If WantLogLess == 0 any improvement will do.
	WantLogLess	uint8

	symbolLen	uint16	// Length of active part of the symbol table.
	maxCount	int	// count of the most probable symbol
	clearCount	bool	// clear count
	actualTableLog	uint8	// Selected tablelog.
	prevTableLog	uint8	// Tablelog for previous table
	prevTable	cTable	// Table used for previous compression.
	cTable		cTable	// compression table
	dt		dTable	// decompression table
	nodes		[]nodeElt
	tmpOut		[4][]byte
	fse		*fse.Scratch
	huffWeight	[maxSymbolValue + 1]byte
}

// TransferCTable will transfer the previously used compression table.
func (s *Scratch) TransferCTable(src *Scratch) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:123
	_go_fuzz_dep_.CoverTab[90270]++
												if cap(s.prevTable) < len(src.prevTable) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:124
		_go_fuzz_dep_.CoverTab[90272]++
													s.prevTable = make(cTable, 0, maxSymbolValue+1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:125
		// _ = "end of CoverTab[90272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:126
		_go_fuzz_dep_.CoverTab[90273]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:126
		// _ = "end of CoverTab[90273]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:126
	// _ = "end of CoverTab[90270]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:126
	_go_fuzz_dep_.CoverTab[90271]++
												s.prevTable = s.prevTable[:len(src.prevTable)]
												copy(s.prevTable, src.prevTable)
												s.prevTableLog = src.prevTableLog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:129
	// _ = "end of CoverTab[90271]"
}

func (s *Scratch) prepare(in []byte) (*Scratch, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:132
	_go_fuzz_dep_.CoverTab[90274]++
												if len(in) > BlockSizeMax {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:133
		_go_fuzz_dep_.CoverTab[90285]++
													return nil, ErrTooBig
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:134
		// _ = "end of CoverTab[90285]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:135
		_go_fuzz_dep_.CoverTab[90286]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:135
		// _ = "end of CoverTab[90286]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:135
	// _ = "end of CoverTab[90274]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:135
	_go_fuzz_dep_.CoverTab[90275]++
												if s == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:136
		_go_fuzz_dep_.CoverTab[90287]++
													s = &Scratch{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:137
		// _ = "end of CoverTab[90287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:138
		_go_fuzz_dep_.CoverTab[90288]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:138
		// _ = "end of CoverTab[90288]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:138
	// _ = "end of CoverTab[90275]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:138
	_go_fuzz_dep_.CoverTab[90276]++
												if s.MaxSymbolValue == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:139
		_go_fuzz_dep_.CoverTab[90289]++
													s.MaxSymbolValue = maxSymbolValue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:140
		// _ = "end of CoverTab[90289]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:141
		_go_fuzz_dep_.CoverTab[90290]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:141
		// _ = "end of CoverTab[90290]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:141
	// _ = "end of CoverTab[90276]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:141
	_go_fuzz_dep_.CoverTab[90277]++
												if s.TableLog == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:142
		_go_fuzz_dep_.CoverTab[90291]++
													s.TableLog = tableLogDefault
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:143
		// _ = "end of CoverTab[90291]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:144
		_go_fuzz_dep_.CoverTab[90292]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:144
		// _ = "end of CoverTab[90292]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:144
	// _ = "end of CoverTab[90277]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:144
	_go_fuzz_dep_.CoverTab[90278]++
												if s.TableLog > tableLogMax || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:145
		_go_fuzz_dep_.CoverTab[90293]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:145
		return s.TableLog < minTablelog
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:145
		// _ = "end of CoverTab[90293]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:145
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:145
		_go_fuzz_dep_.CoverTab[90294]++
													return nil, fmt.Errorf(" invalid tableLog %d (%d -> %d)", s.TableLog, minTablelog, tableLogMax)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:146
		// _ = "end of CoverTab[90294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:147
		_go_fuzz_dep_.CoverTab[90295]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:147
		// _ = "end of CoverTab[90295]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:147
	// _ = "end of CoverTab[90278]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:147
	_go_fuzz_dep_.CoverTab[90279]++
												if s.MaxDecodedSize <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:148
		_go_fuzz_dep_.CoverTab[90296]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:148
		return s.MaxDecodedSize > BlockSizeMax
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:148
		// _ = "end of CoverTab[90296]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:148
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:148
		_go_fuzz_dep_.CoverTab[90297]++
													s.MaxDecodedSize = BlockSizeMax
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:149
		// _ = "end of CoverTab[90297]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:150
		_go_fuzz_dep_.CoverTab[90298]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:150
		// _ = "end of CoverTab[90298]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:150
	// _ = "end of CoverTab[90279]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:150
	_go_fuzz_dep_.CoverTab[90280]++
												if s.clearCount && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:151
		_go_fuzz_dep_.CoverTab[90299]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:151
		return s.maxCount == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:151
		// _ = "end of CoverTab[90299]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:151
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:151
		_go_fuzz_dep_.CoverTab[90300]++
													for i := range s.count {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:152
			_go_fuzz_dep_.CoverTab[90302]++
														s.count[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:153
			// _ = "end of CoverTab[90302]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:154
		// _ = "end of CoverTab[90300]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:154
		_go_fuzz_dep_.CoverTab[90301]++
													s.clearCount = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:155
		// _ = "end of CoverTab[90301]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:156
		_go_fuzz_dep_.CoverTab[90303]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:156
		// _ = "end of CoverTab[90303]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:156
	// _ = "end of CoverTab[90280]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:156
	_go_fuzz_dep_.CoverTab[90281]++
												if cap(s.Out) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:157
		_go_fuzz_dep_.CoverTab[90304]++
													s.Out = make([]byte, 0, len(in))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:158
		// _ = "end of CoverTab[90304]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:159
		_go_fuzz_dep_.CoverTab[90305]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:159
		// _ = "end of CoverTab[90305]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:159
	// _ = "end of CoverTab[90281]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:159
	_go_fuzz_dep_.CoverTab[90282]++
												s.Out = s.Out[:0]

												s.OutTable = nil
												s.OutData = nil
												if cap(s.nodes) < huffNodesLen+1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:164
		_go_fuzz_dep_.CoverTab[90306]++
													s.nodes = make([]nodeElt, 0, huffNodesLen+1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:165
		// _ = "end of CoverTab[90306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:166
		_go_fuzz_dep_.CoverTab[90307]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:166
		// _ = "end of CoverTab[90307]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:166
	// _ = "end of CoverTab[90282]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:166
	_go_fuzz_dep_.CoverTab[90283]++
												s.nodes = s.nodes[:0]
												if s.fse == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:168
		_go_fuzz_dep_.CoverTab[90308]++
													s.fse = &fse.Scratch{}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:169
		// _ = "end of CoverTab[90308]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:170
		_go_fuzz_dep_.CoverTab[90309]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:170
		// _ = "end of CoverTab[90309]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:170
	// _ = "end of CoverTab[90283]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:170
	_go_fuzz_dep_.CoverTab[90284]++
												s.br.init(in)

												return s, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:173
	// _ = "end of CoverTab[90284]"
}

type cTable []cTableEntry

func (c cTable) write(s *Scratch) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:178
	_go_fuzz_dep_.CoverTab[90310]++
												var (
		// precomputed conversion table
		bitsToWeight	[tableLogMax + 1]byte
		huffLog		= s.actualTableLog
		// last weight is not saved.
		maxSymbolValue	= uint8(s.symbolLen - 1)
		huffWeight	= s.huffWeight[:256]
	)
	const (
		maxFSETableLog = 6
	)

	bitsToWeight[0] = 0
	for n := uint8(1); n < huffLog+1; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:192
		_go_fuzz_dep_.CoverTab[90317]++
													bitsToWeight[n] = huffLog + 1 - n
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:193
		// _ = "end of CoverTab[90317]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:194
	// _ = "end of CoverTab[90310]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:194
	_go_fuzz_dep_.CoverTab[90311]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:197
	hist := s.fse.Histogram()
	hist = hist[:256]
	for i := range hist[:16] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:199
		_go_fuzz_dep_.CoverTab[90318]++
													hist[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:200
		// _ = "end of CoverTab[90318]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:201
	// _ = "end of CoverTab[90311]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:201
	_go_fuzz_dep_.CoverTab[90312]++
												for n := uint8(0); n < maxSymbolValue; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:202
		_go_fuzz_dep_.CoverTab[90319]++
													v := bitsToWeight[c[n].nBits] & 15
													huffWeight[n] = v
													hist[v]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:205
		// _ = "end of CoverTab[90319]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:206
	// _ = "end of CoverTab[90312]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:206
	_go_fuzz_dep_.CoverTab[90313]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:209
	if maxSymbolValue >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:209
		_go_fuzz_dep_.CoverTab[90320]++
													huffMaxCnt := uint32(0)
													huffMax := uint8(0)
													for i, v := range hist[:16] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:212
			_go_fuzz_dep_.CoverTab[90322]++
														if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:213
				_go_fuzz_dep_.CoverTab[90324]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:214
				// _ = "end of CoverTab[90324]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:215
				_go_fuzz_dep_.CoverTab[90325]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:215
				// _ = "end of CoverTab[90325]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:215
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:215
			// _ = "end of CoverTab[90322]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:215
			_go_fuzz_dep_.CoverTab[90323]++
														huffMax = byte(i)
														if v > huffMaxCnt {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:217
				_go_fuzz_dep_.CoverTab[90326]++
															huffMaxCnt = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:218
				// _ = "end of CoverTab[90326]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:219
				_go_fuzz_dep_.CoverTab[90327]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:219
				// _ = "end of CoverTab[90327]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:219
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:219
			// _ = "end of CoverTab[90323]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:220
		// _ = "end of CoverTab[90320]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:220
		_go_fuzz_dep_.CoverTab[90321]++
													s.fse.HistogramFinished(huffMax, int(huffMaxCnt))
													s.fse.TableLog = maxFSETableLog
													b, err := fse.Compress(huffWeight[:maxSymbolValue], s.fse)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:224
			_go_fuzz_dep_.CoverTab[90328]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:224
			return len(b) < int(s.symbolLen>>1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:224
			// _ = "end of CoverTab[90328]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:224
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:224
			_go_fuzz_dep_.CoverTab[90329]++
														s.Out = append(s.Out, uint8(len(b)))
														s.Out = append(s.Out, b...)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:227
			// _ = "end of CoverTab[90329]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:228
			_go_fuzz_dep_.CoverTab[90330]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:228
			// _ = "end of CoverTab[90330]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:228
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:228
		// _ = "end of CoverTab[90321]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:230
		_go_fuzz_dep_.CoverTab[90331]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:230
		// _ = "end of CoverTab[90331]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:230
	// _ = "end of CoverTab[90313]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:230
	_go_fuzz_dep_.CoverTab[90314]++

												if maxSymbolValue > (256 - 128) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:232
		_go_fuzz_dep_.CoverTab[90332]++

													return ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:234
		// _ = "end of CoverTab[90332]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:235
		_go_fuzz_dep_.CoverTab[90333]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:235
		// _ = "end of CoverTab[90333]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:235
	// _ = "end of CoverTab[90314]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:235
	_go_fuzz_dep_.CoverTab[90315]++
												op := s.Out

												op = append(op, 128|(maxSymbolValue-1))

												huffWeight[maxSymbolValue] = 0
												for n := uint16(0); n < uint16(maxSymbolValue); n += 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:241
		_go_fuzz_dep_.CoverTab[90334]++
													op = append(op, (huffWeight[n]<<4)|huffWeight[n+1])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:242
		// _ = "end of CoverTab[90334]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:243
	// _ = "end of CoverTab[90315]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:243
	_go_fuzz_dep_.CoverTab[90316]++
												s.Out = op
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:245
	// _ = "end of CoverTab[90316]"
}

func (c cTable) estTableSize(s *Scratch) (sz int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:248
	_go_fuzz_dep_.CoverTab[90335]++
												var (
		// precomputed conversion table
		bitsToWeight	[tableLogMax + 1]byte
		huffLog		= s.actualTableLog
		// last weight is not saved.
		maxSymbolValue	= uint8(s.symbolLen - 1)
		huffWeight	= s.huffWeight[:256]
	)
	const (
		maxFSETableLog = 6
	)

	bitsToWeight[0] = 0
	for n := uint8(1); n < huffLog+1; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:262
		_go_fuzz_dep_.CoverTab[90341]++
													bitsToWeight[n] = huffLog + 1 - n
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:263
		// _ = "end of CoverTab[90341]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:264
	// _ = "end of CoverTab[90335]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:264
	_go_fuzz_dep_.CoverTab[90336]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:267
	hist := s.fse.Histogram()
	hist = hist[:256]
	for i := range hist[:16] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:269
		_go_fuzz_dep_.CoverTab[90342]++
													hist[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:270
		// _ = "end of CoverTab[90342]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:271
	// _ = "end of CoverTab[90336]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:271
	_go_fuzz_dep_.CoverTab[90337]++
												for n := uint8(0); n < maxSymbolValue; n++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:272
		_go_fuzz_dep_.CoverTab[90343]++
													v := bitsToWeight[c[n].nBits] & 15
													huffWeight[n] = v
													hist[v]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:275
		// _ = "end of CoverTab[90343]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:276
	// _ = "end of CoverTab[90337]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:276
	_go_fuzz_dep_.CoverTab[90338]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:279
	if maxSymbolValue >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:279
		_go_fuzz_dep_.CoverTab[90344]++
													huffMaxCnt := uint32(0)
													huffMax := uint8(0)
													for i, v := range hist[:16] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:282
			_go_fuzz_dep_.CoverTab[90346]++
														if v == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:283
				_go_fuzz_dep_.CoverTab[90348]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:284
				// _ = "end of CoverTab[90348]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:285
				_go_fuzz_dep_.CoverTab[90349]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:285
				// _ = "end of CoverTab[90349]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:285
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:285
			// _ = "end of CoverTab[90346]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:285
			_go_fuzz_dep_.CoverTab[90347]++
														huffMax = byte(i)
														if v > huffMaxCnt {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:287
				_go_fuzz_dep_.CoverTab[90350]++
															huffMaxCnt = v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:288
				// _ = "end of CoverTab[90350]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:289
				_go_fuzz_dep_.CoverTab[90351]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:289
				// _ = "end of CoverTab[90351]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:289
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:289
			// _ = "end of CoverTab[90347]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:290
		// _ = "end of CoverTab[90344]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:290
		_go_fuzz_dep_.CoverTab[90345]++
													s.fse.HistogramFinished(huffMax, int(huffMaxCnt))
													s.fse.TableLog = maxFSETableLog
													b, err := fse.Compress(huffWeight[:maxSymbolValue], s.fse)
													if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:294
			_go_fuzz_dep_.CoverTab[90352]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:294
			return len(b) < int(s.symbolLen>>1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:294
			// _ = "end of CoverTab[90352]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:294
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:294
			_go_fuzz_dep_.CoverTab[90353]++
														sz += 1 + len(b)
														return sz, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:296
			// _ = "end of CoverTab[90353]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:297
			_go_fuzz_dep_.CoverTab[90354]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:297
			// _ = "end of CoverTab[90354]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:297
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:297
		// _ = "end of CoverTab[90345]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:299
		_go_fuzz_dep_.CoverTab[90355]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:299
		// _ = "end of CoverTab[90355]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:299
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:299
	// _ = "end of CoverTab[90338]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:299
	_go_fuzz_dep_.CoverTab[90339]++

												if maxSymbolValue > (256 - 128) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:301
		_go_fuzz_dep_.CoverTab[90356]++

													return 0, ErrIncompressible
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:303
		// _ = "end of CoverTab[90356]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:304
		_go_fuzz_dep_.CoverTab[90357]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:304
		// _ = "end of CoverTab[90357]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:304
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:304
	// _ = "end of CoverTab[90339]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:304
	_go_fuzz_dep_.CoverTab[90340]++

												sz += 1 + int(maxSymbolValue/2)
												return sz, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:307
	// _ = "end of CoverTab[90340]"
}

// estimateSize returns the estimated size in bytes of the input represented in the
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:310
// histogram supplied.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:312
func (c cTable) estimateSize(hist []uint32) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:312
	_go_fuzz_dep_.CoverTab[90358]++
												nbBits := uint32(7)
												for i, v := range c[:len(hist)] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:314
		_go_fuzz_dep_.CoverTab[90360]++
													nbBits += uint32(v.nBits) * hist[i]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:315
		// _ = "end of CoverTab[90360]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:316
	// _ = "end of CoverTab[90358]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:316
	_go_fuzz_dep_.CoverTab[90359]++
												return int(nbBits >> 3)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:317
	// _ = "end of CoverTab[90359]"
}

// minSize returns the minimum possible size considering the shannon limit.
func (s *Scratch) minSize(total int) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:321
	_go_fuzz_dep_.CoverTab[90361]++
												nbBits := float64(7)
												fTotal := float64(total)
												for _, v := range s.count[:s.symbolLen] {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:324
		_go_fuzz_dep_.CoverTab[90363]++
													n := float64(v)
													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:326
			_go_fuzz_dep_.CoverTab[90364]++
														nbBits += math.Log2(fTotal/n) * n
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:327
			// _ = "end of CoverTab[90364]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:328
			_go_fuzz_dep_.CoverTab[90365]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:328
			// _ = "end of CoverTab[90365]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:328
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:328
		// _ = "end of CoverTab[90363]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:329
	// _ = "end of CoverTab[90361]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:329
	_go_fuzz_dep_.CoverTab[90362]++
												return int(nbBits) >> 3
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:330
	// _ = "end of CoverTab[90362]"
}

func highBit32(val uint32) (n uint32) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:333
	_go_fuzz_dep_.CoverTab[90366]++
												return uint32(bits.Len32(val) - 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:334
	// _ = "end of CoverTab[90366]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:335
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/huff0.go:335
var _ = _go_fuzz_dep_.CoverTab
