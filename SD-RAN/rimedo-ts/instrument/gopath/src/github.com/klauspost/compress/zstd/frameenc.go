// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:5
)

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/bits"
)

type frameHeader struct {
	ContentSize	uint64
	WindowSize	uint32
	SingleSegment	bool
	Checksum	bool
	DictID		uint32
}

const maxHeaderSize = 14

func (f frameHeader) appendTo(dst []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:25
	_go_fuzz_dep_.CoverTab[94346]++
												dst = append(dst, frameMagic...)
												var fhd uint8
												if f.Checksum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:28
		_go_fuzz_dep_.CoverTab[94356]++
													fhd |= 1 << 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:29
		// _ = "end of CoverTab[94356]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:30
		_go_fuzz_dep_.CoverTab[94357]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:30
		// _ = "end of CoverTab[94357]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:30
	// _ = "end of CoverTab[94346]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:30
	_go_fuzz_dep_.CoverTab[94347]++
												if f.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:31
		_go_fuzz_dep_.CoverTab[94358]++
													fhd |= 1 << 5
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:32
		// _ = "end of CoverTab[94358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:33
		_go_fuzz_dep_.CoverTab[94359]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:33
		// _ = "end of CoverTab[94359]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:33
	// _ = "end of CoverTab[94347]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:33
	_go_fuzz_dep_.CoverTab[94348]++

												var dictIDContent []byte
												if f.DictID > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:36
		_go_fuzz_dep_.CoverTab[94360]++
													var tmp [4]byte
													if f.DictID < 256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:38
			_go_fuzz_dep_.CoverTab[94361]++
														fhd |= 1
														tmp[0] = uint8(f.DictID)
														dictIDContent = tmp[:1]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:41
			// _ = "end of CoverTab[94361]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:42
			_go_fuzz_dep_.CoverTab[94362]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:42
			if f.DictID < 1<<16 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:42
				_go_fuzz_dep_.CoverTab[94363]++
															fhd |= 2
															binary.LittleEndian.PutUint16(tmp[:2], uint16(f.DictID))
															dictIDContent = tmp[:2]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:45
				// _ = "end of CoverTab[94363]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:46
				_go_fuzz_dep_.CoverTab[94364]++
															fhd |= 3
															binary.LittleEndian.PutUint32(tmp[:4], f.DictID)
															dictIDContent = tmp[:4]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:49
				// _ = "end of CoverTab[94364]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:50
			// _ = "end of CoverTab[94362]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:50
		// _ = "end of CoverTab[94360]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:51
		_go_fuzz_dep_.CoverTab[94365]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:51
		// _ = "end of CoverTab[94365]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:51
	// _ = "end of CoverTab[94348]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:51
	_go_fuzz_dep_.CoverTab[94349]++
												var fcs uint8
												if f.ContentSize >= 256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:53
		_go_fuzz_dep_.CoverTab[94366]++
													fcs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:54
		// _ = "end of CoverTab[94366]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:55
		_go_fuzz_dep_.CoverTab[94367]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:55
		// _ = "end of CoverTab[94367]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:55
	// _ = "end of CoverTab[94349]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:55
	_go_fuzz_dep_.CoverTab[94350]++
												if f.ContentSize >= 65536+256 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:56
		_go_fuzz_dep_.CoverTab[94368]++
													fcs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:57
		// _ = "end of CoverTab[94368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:58
		_go_fuzz_dep_.CoverTab[94369]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:58
		// _ = "end of CoverTab[94369]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:58
	// _ = "end of CoverTab[94350]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:58
	_go_fuzz_dep_.CoverTab[94351]++
												if f.ContentSize >= 0xffffffff {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:59
		_go_fuzz_dep_.CoverTab[94370]++
													fcs++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:60
		// _ = "end of CoverTab[94370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:61
		_go_fuzz_dep_.CoverTab[94371]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:61
		// _ = "end of CoverTab[94371]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:61
	// _ = "end of CoverTab[94351]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:61
	_go_fuzz_dep_.CoverTab[94352]++

												fhd |= fcs << 6

												dst = append(dst, fhd)
												if !f.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:66
		_go_fuzz_dep_.CoverTab[94372]++
													const winLogMin = 10
													windowLog := (bits.Len32(f.WindowSize-1) - winLogMin) << 3
													dst = append(dst, uint8(windowLog))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:69
		// _ = "end of CoverTab[94372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:70
		_go_fuzz_dep_.CoverTab[94373]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:70
		// _ = "end of CoverTab[94373]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:70
	// _ = "end of CoverTab[94352]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:70
	_go_fuzz_dep_.CoverTab[94353]++
												if f.DictID > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:71
		_go_fuzz_dep_.CoverTab[94374]++
													dst = append(dst, dictIDContent...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:72
		// _ = "end of CoverTab[94374]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:73
		_go_fuzz_dep_.CoverTab[94375]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:73
		// _ = "end of CoverTab[94375]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:73
	// _ = "end of CoverTab[94353]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:73
	_go_fuzz_dep_.CoverTab[94354]++
												switch fcs {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:75
		_go_fuzz_dep_.CoverTab[94376]++
													if f.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:76
			_go_fuzz_dep_.CoverTab[94381]++
														dst = append(dst, uint8(f.ContentSize))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:77
			// _ = "end of CoverTab[94381]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:78
			_go_fuzz_dep_.CoverTab[94382]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:78
			// _ = "end of CoverTab[94382]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:78
		// _ = "end of CoverTab[94376]"

	case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:80
		_go_fuzz_dep_.CoverTab[94377]++
													f.ContentSize -= 256
													dst = append(dst, uint8(f.ContentSize), uint8(f.ContentSize>>8))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:82
		// _ = "end of CoverTab[94377]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:83
		_go_fuzz_dep_.CoverTab[94378]++
													dst = append(dst, uint8(f.ContentSize), uint8(f.ContentSize>>8), uint8(f.ContentSize>>16), uint8(f.ContentSize>>24))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:84
		// _ = "end of CoverTab[94378]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:85
		_go_fuzz_dep_.CoverTab[94379]++
													dst = append(dst, uint8(f.ContentSize), uint8(f.ContentSize>>8), uint8(f.ContentSize>>16), uint8(f.ContentSize>>24),
			uint8(f.ContentSize>>32), uint8(f.ContentSize>>40), uint8(f.ContentSize>>48), uint8(f.ContentSize>>56))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:87
		// _ = "end of CoverTab[94379]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:88
		_go_fuzz_dep_.CoverTab[94380]++
													panic("invalid fcs")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:89
		// _ = "end of CoverTab[94380]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:90
	// _ = "end of CoverTab[94354]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:90
	_go_fuzz_dep_.CoverTab[94355]++
												return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:91
	// _ = "end of CoverTab[94355]"
}

const skippableFrameHeader = 4 + 4

// calcSkippableFrame will return a total size to be added for written
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:96
// to be divisible by multiple.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:96
// The value will always be > skippableFrameHeader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:96
// The function will panic if written < 0 or wantMultiple <= 0.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:100
func calcSkippableFrame(written, wantMultiple int64) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:100
	_go_fuzz_dep_.CoverTab[94383]++
												if wantMultiple <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:101
		_go_fuzz_dep_.CoverTab[94388]++
													panic("wantMultiple <= 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:102
		// _ = "end of CoverTab[94388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:103
		_go_fuzz_dep_.CoverTab[94389]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:103
		// _ = "end of CoverTab[94389]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:103
	// _ = "end of CoverTab[94383]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:103
	_go_fuzz_dep_.CoverTab[94384]++
												if written < 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:104
		_go_fuzz_dep_.CoverTab[94390]++
													panic("written < 0")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:105
		// _ = "end of CoverTab[94390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:106
		_go_fuzz_dep_.CoverTab[94391]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:106
		// _ = "end of CoverTab[94391]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:106
	// _ = "end of CoverTab[94384]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:106
	_go_fuzz_dep_.CoverTab[94385]++
												leftOver := written % wantMultiple
												if leftOver == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:108
		_go_fuzz_dep_.CoverTab[94392]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:109
		// _ = "end of CoverTab[94392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:110
		_go_fuzz_dep_.CoverTab[94393]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:110
		// _ = "end of CoverTab[94393]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:110
	// _ = "end of CoverTab[94385]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:110
	_go_fuzz_dep_.CoverTab[94386]++
												toAdd := wantMultiple - leftOver
												for toAdd < skippableFrameHeader {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:112
		_go_fuzz_dep_.CoverTab[94394]++
													toAdd += wantMultiple
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:113
		// _ = "end of CoverTab[94394]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:114
	// _ = "end of CoverTab[94386]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:114
	_go_fuzz_dep_.CoverTab[94387]++
												return int(toAdd)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:115
	// _ = "end of CoverTab[94387]"
}

// skippableFrame will add a skippable frame with a total size of bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:118
// total should be >= skippableFrameHeader and < math.MaxUint32.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:120
func skippableFrame(dst []byte, total int, r io.Reader) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:120
	_go_fuzz_dep_.CoverTab[94395]++
												if total == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:121
		_go_fuzz_dep_.CoverTab[94399]++
													return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:122
		// _ = "end of CoverTab[94399]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:123
		_go_fuzz_dep_.CoverTab[94400]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:123
		// _ = "end of CoverTab[94400]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:123
	// _ = "end of CoverTab[94395]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:123
	_go_fuzz_dep_.CoverTab[94396]++
												if total < skippableFrameHeader {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:124
		_go_fuzz_dep_.CoverTab[94401]++
													return dst, fmt.Errorf("requested skippable frame (%d) < 8", total)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:125
		// _ = "end of CoverTab[94401]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:126
		_go_fuzz_dep_.CoverTab[94402]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:126
		// _ = "end of CoverTab[94402]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:126
	// _ = "end of CoverTab[94396]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:126
	_go_fuzz_dep_.CoverTab[94397]++
												if int64(total) > math.MaxUint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:127
		_go_fuzz_dep_.CoverTab[94403]++
													return dst, fmt.Errorf("requested skippable frame (%d) > max uint32", total)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:128
		// _ = "end of CoverTab[94403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:129
		_go_fuzz_dep_.CoverTab[94404]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:129
		// _ = "end of CoverTab[94404]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:129
	// _ = "end of CoverTab[94397]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:129
	_go_fuzz_dep_.CoverTab[94398]++
												dst = append(dst, 0x50, 0x2a, 0x4d, 0x18)
												f := uint32(total - skippableFrameHeader)
												dst = append(dst, uint8(f), uint8(f>>8), uint8(f>>16), uint8(f>>24))
												start := len(dst)
												dst = append(dst, make([]byte, f)...)
												_, err := io.ReadFull(r, dst[start:])
												return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:136
	// _ = "end of CoverTab[94398]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:137
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/frameenc.go:137
var _ = _go_fuzz_dep_.CoverTab
