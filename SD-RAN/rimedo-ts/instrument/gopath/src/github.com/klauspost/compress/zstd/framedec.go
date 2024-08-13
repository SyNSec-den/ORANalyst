// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:5
)

import (
	"bytes"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"sync"

	"github.com/klauspost/compress/zstd/internal/xxhash"
)

type frameDec struct {
	o	decoderOptions
	crc	hash.Hash64
	offset	int64

	WindowSize	uint64

	// In order queue of blocks being decoded.
	decoding	chan *blockDec

	// Frame history passed between blocks
	history	history

	rawInput	byteBuffer

	// Byte buffer that can be reused for small input blocks.
	bBuf	byteBuf

	FrameContentSize	uint64
	frameDone		sync.WaitGroup

	DictionaryID	*uint32
	HasCheckSum	bool
	SingleSegment	bool

	// asyncRunning indicates whether the async routine processes input on 'decoding'.
	asyncRunningMu	sync.Mutex
	asyncRunning	bool
}

const (
	// MinWindowSize is the minimum Window Size, which is 1 KB.
	MinWindowSize	= 1 << 10

	// MaxWindowSize is the maximum encoder window size
	// and the default decoder maximum window size.
	MaxWindowSize	= 1 << 29
)

var (
	frameMagic		= []byte{0x28, 0xb5, 0x2f, 0xfd}
	skippableFrameMagic	= []byte{0x2a, 0x4d, 0x18}
)

func newFrameDec(o decoderOptions) *frameDec {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:62
	_go_fuzz_dep_.CoverTab[94098]++
												if o.maxWindowSize > o.maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:63
		_go_fuzz_dep_.CoverTab[94100]++
													o.maxWindowSize = o.maxDecodedSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:64
		// _ = "end of CoverTab[94100]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:65
		_go_fuzz_dep_.CoverTab[94101]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:65
		// _ = "end of CoverTab[94101]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:65
	// _ = "end of CoverTab[94098]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:65
	_go_fuzz_dep_.CoverTab[94099]++
												d := frameDec{
		o: o,
	}
												return &d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:69
	// _ = "end of CoverTab[94099]"
}

// reset will read the frame header and prepare for block decoding.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:72
// If nothing can be read from the input, io.EOF will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:72
// Any other error indicated that the stream contained data, but
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:72
// there was a problem.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:76
func (d *frameDec) reset(br byteBuffer) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:76
	_go_fuzz_dep_.CoverTab[94102]++
												d.HasCheckSum = false
												d.WindowSize = 0
												var signature [4]byte
												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:80
		_go_fuzz_dep_.CoverTab[94116]++
													var err error

													b, err := br.readSmall(1)
													switch err {
		case io.EOF, io.ErrUnexpectedEOF:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:85
			_go_fuzz_dep_.CoverTab[94121]++
														return io.EOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:86
			// _ = "end of CoverTab[94121]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:87
			_go_fuzz_dep_.CoverTab[94122]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:88
			// _ = "end of CoverTab[94122]"
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:89
			_go_fuzz_dep_.CoverTab[94123]++
														signature[0] = b[0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:90
			// _ = "end of CoverTab[94123]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:91
		// _ = "end of CoverTab[94116]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:91
		_go_fuzz_dep_.CoverTab[94117]++

													b, err = br.readSmall(3)
													switch err {
		case io.EOF:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:95
			_go_fuzz_dep_.CoverTab[94124]++
														return io.EOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:96
			// _ = "end of CoverTab[94124]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:97
			_go_fuzz_dep_.CoverTab[94125]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:98
			// _ = "end of CoverTab[94125]"
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:99
			_go_fuzz_dep_.CoverTab[94126]++
														copy(signature[1:], b)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:100
			// _ = "end of CoverTab[94126]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:101
		// _ = "end of CoverTab[94117]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:101
		_go_fuzz_dep_.CoverTab[94118]++

													if !bytes.Equal(signature[1:4], skippableFrameMagic) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:103
			_go_fuzz_dep_.CoverTab[94127]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:103
			return signature[0]&0xf0 != 0x50
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:103
			// _ = "end of CoverTab[94127]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:103
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:103
			_go_fuzz_dep_.CoverTab[94128]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:104
				_go_fuzz_dep_.CoverTab[94130]++
															println("Not skippable", hex.EncodeToString(signature[:]), hex.EncodeToString(skippableFrameMagic))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:105
				// _ = "end of CoverTab[94130]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:106
				_go_fuzz_dep_.CoverTab[94131]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:106
				// _ = "end of CoverTab[94131]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:106
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:106
			// _ = "end of CoverTab[94128]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:106
			_go_fuzz_dep_.CoverTab[94129]++

														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:108
			// _ = "end of CoverTab[94129]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:109
			_go_fuzz_dep_.CoverTab[94132]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:109
			// _ = "end of CoverTab[94132]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:109
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:109
		// _ = "end of CoverTab[94118]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:109
		_go_fuzz_dep_.CoverTab[94119]++

													b, err = br.readSmall(4)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:112
			_go_fuzz_dep_.CoverTab[94133]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:113
				_go_fuzz_dep_.CoverTab[94135]++
															println("Reading Frame Size", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:114
				// _ = "end of CoverTab[94135]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:115
				_go_fuzz_dep_.CoverTab[94136]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:115
				// _ = "end of CoverTab[94136]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:115
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:115
			// _ = "end of CoverTab[94133]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:115
			_go_fuzz_dep_.CoverTab[94134]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:116
			// _ = "end of CoverTab[94134]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:117
			_go_fuzz_dep_.CoverTab[94137]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:117
			// _ = "end of CoverTab[94137]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:117
		// _ = "end of CoverTab[94119]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:117
		_go_fuzz_dep_.CoverTab[94120]++
													n := uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
													println("Skipping frame with", n, "bytes.")
													err = br.skipN(int(n))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:121
			_go_fuzz_dep_.CoverTab[94138]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:122
				_go_fuzz_dep_.CoverTab[94140]++
															println("Reading discarded frame", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:123
				// _ = "end of CoverTab[94140]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:124
				_go_fuzz_dep_.CoverTab[94141]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:124
				// _ = "end of CoverTab[94141]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:124
			// _ = "end of CoverTab[94138]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:124
			_go_fuzz_dep_.CoverTab[94139]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:125
			// _ = "end of CoverTab[94139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:126
			_go_fuzz_dep_.CoverTab[94142]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:126
			// _ = "end of CoverTab[94142]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:126
		// _ = "end of CoverTab[94120]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:127
	// _ = "end of CoverTab[94102]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:127
	_go_fuzz_dep_.CoverTab[94103]++
												if !bytes.Equal(signature[:], frameMagic) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:128
		_go_fuzz_dep_.CoverTab[94143]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:129
			_go_fuzz_dep_.CoverTab[94145]++
														println("Got magic numbers: ", signature, "want:", frameMagic)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:130
			// _ = "end of CoverTab[94145]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:131
			_go_fuzz_dep_.CoverTab[94146]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:131
			// _ = "end of CoverTab[94146]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:131
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:131
		// _ = "end of CoverTab[94143]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:131
		_go_fuzz_dep_.CoverTab[94144]++
													return ErrMagicMismatch
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:132
		// _ = "end of CoverTab[94144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:133
		_go_fuzz_dep_.CoverTab[94147]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:133
		// _ = "end of CoverTab[94147]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:133
	// _ = "end of CoverTab[94103]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:133
	_go_fuzz_dep_.CoverTab[94104]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:136
	fhd, err := br.readByte()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:137
		_go_fuzz_dep_.CoverTab[94148]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:138
			_go_fuzz_dep_.CoverTab[94150]++
														println("Reading Frame_Header_Descriptor", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:139
			// _ = "end of CoverTab[94150]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:140
			_go_fuzz_dep_.CoverTab[94151]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:140
			// _ = "end of CoverTab[94151]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:140
		// _ = "end of CoverTab[94148]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:140
		_go_fuzz_dep_.CoverTab[94149]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:141
		// _ = "end of CoverTab[94149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:142
		_go_fuzz_dep_.CoverTab[94152]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:142
		// _ = "end of CoverTab[94152]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:142
	// _ = "end of CoverTab[94104]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:142
	_go_fuzz_dep_.CoverTab[94105]++
												d.SingleSegment = fhd&(1<<5) != 0

												if fhd&(1<<3) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:145
		_go_fuzz_dep_.CoverTab[94153]++
													return errors.New("reserved bit set on frame header")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:146
		// _ = "end of CoverTab[94153]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:147
		_go_fuzz_dep_.CoverTab[94154]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:147
		// _ = "end of CoverTab[94154]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:147
	// _ = "end of CoverTab[94105]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:147
	_go_fuzz_dep_.CoverTab[94106]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:151
	d.WindowSize = 0
	if !d.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:152
		_go_fuzz_dep_.CoverTab[94155]++
													wd, err := br.readByte()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:154
			_go_fuzz_dep_.CoverTab[94157]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:155
				_go_fuzz_dep_.CoverTab[94159]++
															println("Reading Window_Descriptor", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:156
				// _ = "end of CoverTab[94159]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:157
				_go_fuzz_dep_.CoverTab[94160]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:157
				// _ = "end of CoverTab[94160]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:157
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:157
			// _ = "end of CoverTab[94157]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:157
			_go_fuzz_dep_.CoverTab[94158]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:158
			// _ = "end of CoverTab[94158]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:159
			_go_fuzz_dep_.CoverTab[94161]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:159
			// _ = "end of CoverTab[94161]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:159
		// _ = "end of CoverTab[94155]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:159
		_go_fuzz_dep_.CoverTab[94156]++
													printf("raw: %x, mantissa: %d, exponent: %d\n", wd, wd&7, wd>>3)
													windowLog := 10 + (wd >> 3)
													windowBase := uint64(1) << windowLog
													windowAdd := (windowBase / 8) * uint64(wd&0x7)
													d.WindowSize = windowBase + windowAdd
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:164
		// _ = "end of CoverTab[94156]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:165
		_go_fuzz_dep_.CoverTab[94162]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:165
		// _ = "end of CoverTab[94162]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:165
	// _ = "end of CoverTab[94106]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:165
	_go_fuzz_dep_.CoverTab[94107]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:169
	d.DictionaryID = nil
	if size := fhd & 3; size != 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:170
		_go_fuzz_dep_.CoverTab[94163]++
													if size == 3 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:171
			_go_fuzz_dep_.CoverTab[94168]++
														size = 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:172
			// _ = "end of CoverTab[94168]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:173
			_go_fuzz_dep_.CoverTab[94169]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:173
			// _ = "end of CoverTab[94169]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:173
		// _ = "end of CoverTab[94163]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:173
		_go_fuzz_dep_.CoverTab[94164]++

													b, err := br.readSmall(int(size))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:176
			_go_fuzz_dep_.CoverTab[94170]++
														println("Reading Dictionary_ID", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:178
			// _ = "end of CoverTab[94170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:179
			_go_fuzz_dep_.CoverTab[94171]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:179
			// _ = "end of CoverTab[94171]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:179
		// _ = "end of CoverTab[94164]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:179
		_go_fuzz_dep_.CoverTab[94165]++
													var id uint32
													switch size {
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:182
			_go_fuzz_dep_.CoverTab[94172]++
														id = uint32(b[0])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:183
			// _ = "end of CoverTab[94172]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:184
			_go_fuzz_dep_.CoverTab[94173]++
														id = uint32(b[0]) | (uint32(b[1]) << 8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:185
			// _ = "end of CoverTab[94173]"
		case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:186
			_go_fuzz_dep_.CoverTab[94174]++
														id = uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:187
			// _ = "end of CoverTab[94174]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:187
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:187
			_go_fuzz_dep_.CoverTab[94175]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:187
			// _ = "end of CoverTab[94175]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:188
		// _ = "end of CoverTab[94165]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:188
		_go_fuzz_dep_.CoverTab[94166]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:189
			_go_fuzz_dep_.CoverTab[94176]++
														println("Dict size", size, "ID:", id)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:190
			// _ = "end of CoverTab[94176]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:191
			_go_fuzz_dep_.CoverTab[94177]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:191
			// _ = "end of CoverTab[94177]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:191
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:191
		// _ = "end of CoverTab[94166]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:191
		_go_fuzz_dep_.CoverTab[94167]++
													if id > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:192
			_go_fuzz_dep_.CoverTab[94178]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:195
			d.DictionaryID = &id
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:195
			// _ = "end of CoverTab[94178]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:196
			_go_fuzz_dep_.CoverTab[94179]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:196
			// _ = "end of CoverTab[94179]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:196
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:196
		// _ = "end of CoverTab[94167]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:197
		_go_fuzz_dep_.CoverTab[94180]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:197
		// _ = "end of CoverTab[94180]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:197
	// _ = "end of CoverTab[94107]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:197
	_go_fuzz_dep_.CoverTab[94108]++

	// Read Frame_Content_Size
	// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#frame_content_size
	var fcsSize int
	v := fhd >> 6
	switch v {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:204
		_go_fuzz_dep_.CoverTab[94181]++
													if d.SingleSegment {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:205
			_go_fuzz_dep_.CoverTab[94183]++
														fcsSize = 1
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:206
			// _ = "end of CoverTab[94183]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:207
			_go_fuzz_dep_.CoverTab[94184]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:207
			// _ = "end of CoverTab[94184]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:207
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:207
		// _ = "end of CoverTab[94181]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:208
		_go_fuzz_dep_.CoverTab[94182]++
													fcsSize = 1 << v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:209
		// _ = "end of CoverTab[94182]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:210
	// _ = "end of CoverTab[94108]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:210
	_go_fuzz_dep_.CoverTab[94109]++
												d.FrameContentSize = 0
												if fcsSize > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:212
		_go_fuzz_dep_.CoverTab[94185]++
													b, err := br.readSmall(fcsSize)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:214
			_go_fuzz_dep_.CoverTab[94188]++
														println("Reading Frame content", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:216
			// _ = "end of CoverTab[94188]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:217
			_go_fuzz_dep_.CoverTab[94189]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:217
			// _ = "end of CoverTab[94189]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:217
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:217
		// _ = "end of CoverTab[94185]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:217
		_go_fuzz_dep_.CoverTab[94186]++
													switch fcsSize {
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:219
			_go_fuzz_dep_.CoverTab[94190]++
														d.FrameContentSize = uint64(b[0])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:220
			// _ = "end of CoverTab[94190]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:221
			_go_fuzz_dep_.CoverTab[94191]++

														d.FrameContentSize = uint64(b[0]) | (uint64(b[1]) << 8) + 256
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:223
			// _ = "end of CoverTab[94191]"
		case 4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:224
			_go_fuzz_dep_.CoverTab[94192]++
														d.FrameContentSize = uint64(b[0]) | (uint64(b[1]) << 8) | (uint64(b[2]) << 16) | (uint64(b[3]) << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:225
			// _ = "end of CoverTab[94192]"
		case 8:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:226
			_go_fuzz_dep_.CoverTab[94193]++
														d1 := uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
														d2 := uint32(b[4]) | (uint32(b[5]) << 8) | (uint32(b[6]) << 16) | (uint32(b[7]) << 24)
														d.FrameContentSize = uint64(d1) | (uint64(d2) << 32)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:229
			// _ = "end of CoverTab[94193]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:229
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:229
			_go_fuzz_dep_.CoverTab[94194]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:229
			// _ = "end of CoverTab[94194]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:230
		// _ = "end of CoverTab[94186]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:230
		_go_fuzz_dep_.CoverTab[94187]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:231
			_go_fuzz_dep_.CoverTab[94195]++
														println("field size bits:", v, "fcsSize:", fcsSize, "FrameContentSize:", d.FrameContentSize, hex.EncodeToString(b[:fcsSize]), "singleseg:", d.SingleSegment, "window:", d.WindowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:232
			// _ = "end of CoverTab[94195]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:233
			_go_fuzz_dep_.CoverTab[94196]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:233
			// _ = "end of CoverTab[94196]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:233
		// _ = "end of CoverTab[94187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:234
		_go_fuzz_dep_.CoverTab[94197]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:234
		// _ = "end of CoverTab[94197]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:234
	// _ = "end of CoverTab[94109]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:234
	_go_fuzz_dep_.CoverTab[94110]++

												d.HasCheckSum = fhd&(1<<2) != 0
												if d.HasCheckSum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:237
		_go_fuzz_dep_.CoverTab[94198]++
													if d.crc == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:238
			_go_fuzz_dep_.CoverTab[94200]++
														d.crc = xxhash.New()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:239
			// _ = "end of CoverTab[94200]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:240
			_go_fuzz_dep_.CoverTab[94201]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:240
			// _ = "end of CoverTab[94201]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:240
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:240
		// _ = "end of CoverTab[94198]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:240
		_go_fuzz_dep_.CoverTab[94199]++
													d.crc.Reset()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:241
		// _ = "end of CoverTab[94199]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:242
		_go_fuzz_dep_.CoverTab[94202]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:242
		// _ = "end of CoverTab[94202]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:242
	// _ = "end of CoverTab[94110]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:242
	_go_fuzz_dep_.CoverTab[94111]++

												if d.WindowSize == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:244
		_go_fuzz_dep_.CoverTab[94203]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:244
		return d.SingleSegment
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:244
		// _ = "end of CoverTab[94203]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:244
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:244
		_go_fuzz_dep_.CoverTab[94204]++

													d.WindowSize = d.FrameContentSize
													if d.WindowSize < MinWindowSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:247
			_go_fuzz_dep_.CoverTab[94205]++
														d.WindowSize = MinWindowSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:248
			// _ = "end of CoverTab[94205]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:249
			_go_fuzz_dep_.CoverTab[94206]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:249
			// _ = "end of CoverTab[94206]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:249
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:249
		// _ = "end of CoverTab[94204]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:250
		_go_fuzz_dep_.CoverTab[94207]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:250
		// _ = "end of CoverTab[94207]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:250
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:250
	// _ = "end of CoverTab[94111]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:250
	_go_fuzz_dep_.CoverTab[94112]++

												if d.WindowSize > uint64(d.o.maxWindowSize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:252
		_go_fuzz_dep_.CoverTab[94208]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:253
			_go_fuzz_dep_.CoverTab[94210]++
														printf("window size %d > max %d\n", d.WindowSize, d.o.maxWindowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:254
			// _ = "end of CoverTab[94210]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:255
			_go_fuzz_dep_.CoverTab[94211]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:255
			// _ = "end of CoverTab[94211]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:255
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:255
		// _ = "end of CoverTab[94208]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:255
		_go_fuzz_dep_.CoverTab[94209]++
													return ErrWindowSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:256
		// _ = "end of CoverTab[94209]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:257
		_go_fuzz_dep_.CoverTab[94212]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:257
		// _ = "end of CoverTab[94212]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:257
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:257
	// _ = "end of CoverTab[94112]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:257
	_go_fuzz_dep_.CoverTab[94113]++

												if d.WindowSize < MinWindowSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:259
		_go_fuzz_dep_.CoverTab[94213]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:260
			_go_fuzz_dep_.CoverTab[94215]++
														println("got window size: ", d.WindowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:261
			// _ = "end of CoverTab[94215]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:262
			_go_fuzz_dep_.CoverTab[94216]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:262
			// _ = "end of CoverTab[94216]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:262
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:262
		// _ = "end of CoverTab[94213]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:262
		_go_fuzz_dep_.CoverTab[94214]++
													return ErrWindowSizeTooSmall
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:263
		// _ = "end of CoverTab[94214]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:264
		_go_fuzz_dep_.CoverTab[94217]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:264
		// _ = "end of CoverTab[94217]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:264
	// _ = "end of CoverTab[94113]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:264
	_go_fuzz_dep_.CoverTab[94114]++
												d.history.windowSize = int(d.WindowSize)
												if d.o.lowMem && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:266
		_go_fuzz_dep_.CoverTab[94218]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:266
		return d.history.windowSize < maxBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:266
		// _ = "end of CoverTab[94218]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:266
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:266
		_go_fuzz_dep_.CoverTab[94219]++
													d.history.maxSize = d.history.windowSize * 2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:267
		// _ = "end of CoverTab[94219]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:268
		_go_fuzz_dep_.CoverTab[94220]++
													d.history.maxSize = d.history.windowSize + maxBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:269
		// _ = "end of CoverTab[94220]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:270
	// _ = "end of CoverTab[94114]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:270
	_go_fuzz_dep_.CoverTab[94115]++

												d.rawInput = br
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:273
	// _ = "end of CoverTab[94115]"
}

// next will start decoding the next block from stream.
func (d *frameDec) next(block *blockDec) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:277
	_go_fuzz_dep_.CoverTab[94221]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:278
		_go_fuzz_dep_.CoverTab[94227]++
													printf("decoding new block %p:%p", block, block.data)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:279
		// _ = "end of CoverTab[94227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:280
		_go_fuzz_dep_.CoverTab[94228]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:280
		// _ = "end of CoverTab[94228]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:280
	// _ = "end of CoverTab[94221]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:280
	_go_fuzz_dep_.CoverTab[94222]++
												err := block.reset(d.rawInput, d.WindowSize)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:282
		_go_fuzz_dep_.CoverTab[94229]++
													println("block error:", err)

													d.sendErr(block, err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:286
		// _ = "end of CoverTab[94229]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:287
		_go_fuzz_dep_.CoverTab[94230]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:287
		// _ = "end of CoverTab[94230]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:287
	// _ = "end of CoverTab[94222]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:287
	_go_fuzz_dep_.CoverTab[94223]++
												block.input <- struct{}{}
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:289
		_go_fuzz_dep_.CoverTab[94231]++
													println("next block:", block)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:290
		// _ = "end of CoverTab[94231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:291
		_go_fuzz_dep_.CoverTab[94232]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:291
		// _ = "end of CoverTab[94232]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:291
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:291
	// _ = "end of CoverTab[94223]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:291
	_go_fuzz_dep_.CoverTab[94224]++
												d.asyncRunningMu.Lock()
												defer d.asyncRunningMu.Unlock()
												if !d.asyncRunning {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:294
		_go_fuzz_dep_.CoverTab[94233]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:295
		// _ = "end of CoverTab[94233]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:296
		_go_fuzz_dep_.CoverTab[94234]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:296
		// _ = "end of CoverTab[94234]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:296
	// _ = "end of CoverTab[94224]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:296
	_go_fuzz_dep_.CoverTab[94225]++
												if block.Last {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:297
		_go_fuzz_dep_.CoverTab[94235]++

													d.decoding <- block
													return io.EOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:300
		// _ = "end of CoverTab[94235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:301
		_go_fuzz_dep_.CoverTab[94236]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:301
		// _ = "end of CoverTab[94236]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:301
	// _ = "end of CoverTab[94225]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:301
	_go_fuzz_dep_.CoverTab[94226]++
												d.decoding <- block
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:303
	// _ = "end of CoverTab[94226]"
}

// sendEOF will queue an error block on the frame.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:306
// This will cause the frame decoder to return when it encounters the block.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:306
// Returns true if the decoder was added.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:309
func (d *frameDec) sendErr(block *blockDec, err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:309
	_go_fuzz_dep_.CoverTab[94237]++
												d.asyncRunningMu.Lock()
												defer d.asyncRunningMu.Unlock()
												if !d.asyncRunning {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:312
		_go_fuzz_dep_.CoverTab[94239]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:313
		// _ = "end of CoverTab[94239]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:314
		_go_fuzz_dep_.CoverTab[94240]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:314
		// _ = "end of CoverTab[94240]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:314
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:314
	// _ = "end of CoverTab[94237]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:314
	_go_fuzz_dep_.CoverTab[94238]++

												println("sending error", err.Error())
												block.sendErr(err)
												d.decoding <- block
												return true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:319
	// _ = "end of CoverTab[94238]"
}

// checkCRC will check the checksum if the frame has one.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:322
// Will return ErrCRCMismatch if crc check failed, otherwise nil.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:324
func (d *frameDec) checkCRC() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:324
	_go_fuzz_dep_.CoverTab[94241]++
												if !d.HasCheckSum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:325
		_go_fuzz_dep_.CoverTab[94246]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:326
		// _ = "end of CoverTab[94246]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:327
		_go_fuzz_dep_.CoverTab[94247]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:327
		// _ = "end of CoverTab[94247]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:327
	// _ = "end of CoverTab[94241]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:327
	_go_fuzz_dep_.CoverTab[94242]++
												var tmp [4]byte
												got := d.crc.Sum64()

												tmp[0] = byte(got >> 0)
												tmp[1] = byte(got >> 8)
												tmp[2] = byte(got >> 16)
												tmp[3] = byte(got >> 24)

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:337
	want, err := d.rawInput.readSmall(4)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:338
		_go_fuzz_dep_.CoverTab[94248]++
													println("CRC missing?", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:340
		// _ = "end of CoverTab[94248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:341
		_go_fuzz_dep_.CoverTab[94249]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:341
		// _ = "end of CoverTab[94249]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:341
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:341
	// _ = "end of CoverTab[94242]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:341
	_go_fuzz_dep_.CoverTab[94243]++

												if !bytes.Equal(tmp[:], want) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:343
		_go_fuzz_dep_.CoverTab[94250]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:344
			_go_fuzz_dep_.CoverTab[94252]++
														println("CRC Check Failed:", tmp[:], "!=", want)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:345
			// _ = "end of CoverTab[94252]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:346
			_go_fuzz_dep_.CoverTab[94253]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:346
			// _ = "end of CoverTab[94253]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:346
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:346
		// _ = "end of CoverTab[94250]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:346
		_go_fuzz_dep_.CoverTab[94251]++
													return ErrCRCMismatch
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:347
		// _ = "end of CoverTab[94251]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:348
		_go_fuzz_dep_.CoverTab[94254]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:348
		// _ = "end of CoverTab[94254]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:348
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:348
	// _ = "end of CoverTab[94243]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:348
	_go_fuzz_dep_.CoverTab[94244]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:349
		_go_fuzz_dep_.CoverTab[94255]++
													println("CRC ok", tmp[:])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:350
		// _ = "end of CoverTab[94255]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:351
		_go_fuzz_dep_.CoverTab[94256]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:351
		// _ = "end of CoverTab[94256]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:351
	// _ = "end of CoverTab[94244]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:351
	_go_fuzz_dep_.CoverTab[94245]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:352
	// _ = "end of CoverTab[94245]"
}

func (d *frameDec) initAsync() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:355
	_go_fuzz_dep_.CoverTab[94257]++
												if !d.o.lowMem && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:356
		_go_fuzz_dep_.CoverTab[94263]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:356
		return !d.SingleSegment
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:356
		// _ = "end of CoverTab[94263]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:356
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:356
		_go_fuzz_dep_.CoverTab[94264]++

													d.history.maxSize = d.history.windowSize + maxBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:358
		// _ = "end of CoverTab[94264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:359
		_go_fuzz_dep_.CoverTab[94265]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:359
		// _ = "end of CoverTab[94265]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:359
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:359
	// _ = "end of CoverTab[94257]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:359
	_go_fuzz_dep_.CoverTab[94258]++

												if d.o.lowMem && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:361
		_go_fuzz_dep_.CoverTab[94266]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:361
		return cap(d.history.b) > d.history.maxSize+maxBlockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:361
		// _ = "end of CoverTab[94266]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:361
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:361
		_go_fuzz_dep_.CoverTab[94267]++
													d.history.b = make([]byte, 0, d.history.maxSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:362
		// _ = "end of CoverTab[94267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:363
		_go_fuzz_dep_.CoverTab[94268]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:363
		// _ = "end of CoverTab[94268]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:363
	// _ = "end of CoverTab[94258]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:363
	_go_fuzz_dep_.CoverTab[94259]++
												if cap(d.history.b) < d.history.maxSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:364
		_go_fuzz_dep_.CoverTab[94269]++
													d.history.b = make([]byte, 0, d.history.maxSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:365
		// _ = "end of CoverTab[94269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:366
		_go_fuzz_dep_.CoverTab[94270]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:366
		// _ = "end of CoverTab[94270]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:366
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:366
	// _ = "end of CoverTab[94259]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:366
	_go_fuzz_dep_.CoverTab[94260]++
												if cap(d.decoding) < d.o.concurrent {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:367
		_go_fuzz_dep_.CoverTab[94271]++
													d.decoding = make(chan *blockDec, d.o.concurrent)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:368
		// _ = "end of CoverTab[94271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:369
		_go_fuzz_dep_.CoverTab[94272]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:369
		// _ = "end of CoverTab[94272]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:369
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:369
	// _ = "end of CoverTab[94260]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:369
	_go_fuzz_dep_.CoverTab[94261]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:370
		_go_fuzz_dep_.CoverTab[94273]++
													h := d.history
													printf("history init. len: %d, cap: %d", len(h.b), cap(h.b))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:372
		// _ = "end of CoverTab[94273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:373
		_go_fuzz_dep_.CoverTab[94274]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:373
		// _ = "end of CoverTab[94274]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:373
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:373
	// _ = "end of CoverTab[94261]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:373
	_go_fuzz_dep_.CoverTab[94262]++
												d.asyncRunningMu.Lock()
												d.asyncRunning = true
												d.asyncRunningMu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:376
	// _ = "end of CoverTab[94262]"
}

// startDecoder will start decoding blocks and write them to the writer.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:379
// The decoder will stop as soon as an error occurs or at end of frame.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:379
// When the frame has finished decoding the *bufio.Reader
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:379
// containing the remaining input will be sent on frameDec.frameDone.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:383
func (d *frameDec) startDecoder(output chan decodeOutput) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:383
	_go_fuzz_dep_.CoverTab[94275]++
												written := int64(0)

												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:386
		_go_fuzz_dep_.CoverTab[94277]++
													d.asyncRunningMu.Lock()
													d.asyncRunning = false
													d.asyncRunningMu.Unlock()

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:392
		d.history.error = true
	flushdone:
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:394
			_go_fuzz_dep_.CoverTab[94279]++
														select {
			case b := <-d.decoding:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:396
				_go_fuzz_dep_.CoverTab[94280]++
															b.history <- &d.history
															output <- <-b.result
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:398
				// _ = "end of CoverTab[94280]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:399
				_go_fuzz_dep_.CoverTab[94281]++
															break flushdone
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:400
				// _ = "end of CoverTab[94281]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:401
			// _ = "end of CoverTab[94279]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:402
		// _ = "end of CoverTab[94277]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:402
		_go_fuzz_dep_.CoverTab[94278]++
													println("frame decoder done, signalling done")
													d.frameDone.Done()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:404
		// _ = "end of CoverTab[94278]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:405
	// _ = "end of CoverTab[94275]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:405
	_go_fuzz_dep_.CoverTab[94276]++

												block := <-d.decoding
												block.history <- &d.history
												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:409
		_go_fuzz_dep_.CoverTab[94282]++
													var next *blockDec

													r := <-block.result
													if r.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:413
			_go_fuzz_dep_.CoverTab[94290]++
														println("Result contained error", r.err)
														output <- r
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:416
			// _ = "end of CoverTab[94290]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:417
			_go_fuzz_dep_.CoverTab[94291]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:417
			// _ = "end of CoverTab[94291]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:417
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:417
		// _ = "end of CoverTab[94282]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:417
		_go_fuzz_dep_.CoverTab[94283]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:418
			_go_fuzz_dep_.CoverTab[94292]++
														println("got result, from ", d.offset, "to", d.offset+int64(len(r.b)))
														d.offset += int64(len(r.b))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:420
			// _ = "end of CoverTab[94292]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:421
			_go_fuzz_dep_.CoverTab[94293]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:421
			// _ = "end of CoverTab[94293]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:421
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:421
		// _ = "end of CoverTab[94283]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:421
		_go_fuzz_dep_.CoverTab[94284]++
													if !block.Last {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:422
			_go_fuzz_dep_.CoverTab[94294]++

														select {
			case next = <-d.decoding:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:425
				_go_fuzz_dep_.CoverTab[94295]++
															if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:426
					_go_fuzz_dep_.CoverTab[94298]++
																println("Sending ", len(d.history.b), "bytes as history")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:427
					// _ = "end of CoverTab[94298]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:428
					_go_fuzz_dep_.CoverTab[94299]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:428
					// _ = "end of CoverTab[94299]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:428
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:428
				// _ = "end of CoverTab[94295]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:428
				_go_fuzz_dep_.CoverTab[94296]++
															next.history <- &d.history
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:429
				// _ = "end of CoverTab[94296]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:430
				_go_fuzz_dep_.CoverTab[94297]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:433
				next = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:433
				// _ = "end of CoverTab[94297]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:434
			// _ = "end of CoverTab[94294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:435
			_go_fuzz_dep_.CoverTab[94300]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:435
			// _ = "end of CoverTab[94300]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:435
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:435
		// _ = "end of CoverTab[94284]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:435
		_go_fuzz_dep_.CoverTab[94285]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:438
		if d.HasCheckSum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:438
			_go_fuzz_dep_.CoverTab[94301]++
														n, err := d.crc.Write(r.b)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:440
				_go_fuzz_dep_.CoverTab[94302]++
															r.err = err
															if n != len(r.b) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:442
					_go_fuzz_dep_.CoverTab[94304]++
																r.err = io.ErrShortWrite
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:443
					// _ = "end of CoverTab[94304]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:444
					_go_fuzz_dep_.CoverTab[94305]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:444
					// _ = "end of CoverTab[94305]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:444
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:444
				// _ = "end of CoverTab[94302]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:444
				_go_fuzz_dep_.CoverTab[94303]++
															output <- r
															return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:446
				// _ = "end of CoverTab[94303]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:447
				_go_fuzz_dep_.CoverTab[94306]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:447
				// _ = "end of CoverTab[94306]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:447
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:447
			// _ = "end of CoverTab[94301]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:448
			_go_fuzz_dep_.CoverTab[94307]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:448
			// _ = "end of CoverTab[94307]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:448
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:448
		// _ = "end of CoverTab[94285]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:448
		_go_fuzz_dep_.CoverTab[94286]++
													written += int64(len(r.b))
													if d.SingleSegment && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:450
			_go_fuzz_dep_.CoverTab[94308]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:450
			return uint64(written) > d.FrameContentSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:450
			// _ = "end of CoverTab[94308]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:450
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:450
			_go_fuzz_dep_.CoverTab[94309]++
														println("runDecoder: single segment and", uint64(written), ">", d.FrameContentSize)
														r.err = ErrFrameSizeExceeded
														output <- r
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:454
			// _ = "end of CoverTab[94309]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:455
			_go_fuzz_dep_.CoverTab[94310]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:455
			// _ = "end of CoverTab[94310]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:455
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:455
		// _ = "end of CoverTab[94286]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:455
		_go_fuzz_dep_.CoverTab[94287]++
													if block.Last {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:456
			_go_fuzz_dep_.CoverTab[94311]++
														r.err = d.checkCRC()
														output <- r
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:459
			// _ = "end of CoverTab[94311]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:460
			_go_fuzz_dep_.CoverTab[94312]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:460
			// _ = "end of CoverTab[94312]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:460
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:460
		// _ = "end of CoverTab[94287]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:460
		_go_fuzz_dep_.CoverTab[94288]++
													output <- r
													if next == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:462
			_go_fuzz_dep_.CoverTab[94313]++

														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:464
				_go_fuzz_dep_.CoverTab[94315]++
															println("Sending ", len(d.history.b), " bytes as history")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:465
				// _ = "end of CoverTab[94315]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:466
				_go_fuzz_dep_.CoverTab[94316]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:466
				// _ = "end of CoverTab[94316]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:466
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:466
			// _ = "end of CoverTab[94313]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:466
			_go_fuzz_dep_.CoverTab[94314]++
														next = <-d.decoding
														next.history <- &d.history
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:468
			// _ = "end of CoverTab[94314]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:469
			_go_fuzz_dep_.CoverTab[94317]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:469
			// _ = "end of CoverTab[94317]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:469
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:469
		// _ = "end of CoverTab[94288]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:469
		_go_fuzz_dep_.CoverTab[94289]++
													block = next
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:470
		// _ = "end of CoverTab[94289]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:471
	// _ = "end of CoverTab[94276]"
}

// runDecoder will create a sync decoder that will decode a block of data.
func (d *frameDec) runDecoder(dst []byte, dec *blockDec) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:475
	_go_fuzz_dep_.CoverTab[94318]++
												saved := d.history.b

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:479
	d.history.b = dst

	crcStart := len(dst)
	var err error
	for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:483
		_go_fuzz_dep_.CoverTab[94321]++
													err = dec.reset(d.rawInput, d.WindowSize)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:485
			_go_fuzz_dep_.CoverTab[94326]++
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:486
			// _ = "end of CoverTab[94326]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:487
			_go_fuzz_dep_.CoverTab[94327]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:487
			// _ = "end of CoverTab[94327]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:487
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:487
		// _ = "end of CoverTab[94321]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:487
		_go_fuzz_dep_.CoverTab[94322]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:488
			_go_fuzz_dep_.CoverTab[94328]++
														println("next block:", dec)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:489
			// _ = "end of CoverTab[94328]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:490
			_go_fuzz_dep_.CoverTab[94329]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:490
			// _ = "end of CoverTab[94329]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:490
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:490
		// _ = "end of CoverTab[94322]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:490
		_go_fuzz_dep_.CoverTab[94323]++
													err = dec.decodeBuf(&d.history)
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:492
			_go_fuzz_dep_.CoverTab[94330]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:492
			return dec.Last
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:492
			// _ = "end of CoverTab[94330]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:492
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:492
			_go_fuzz_dep_.CoverTab[94331]++
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:493
			// _ = "end of CoverTab[94331]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:494
			_go_fuzz_dep_.CoverTab[94332]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:494
			// _ = "end of CoverTab[94332]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:494
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:494
		// _ = "end of CoverTab[94323]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:494
		_go_fuzz_dep_.CoverTab[94324]++
													if uint64(len(d.history.b)) > d.o.maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:495
			_go_fuzz_dep_.CoverTab[94333]++
														err = ErrDecoderSizeExceeded
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:497
			// _ = "end of CoverTab[94333]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:498
			_go_fuzz_dep_.CoverTab[94334]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:498
			// _ = "end of CoverTab[94334]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:498
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:498
		// _ = "end of CoverTab[94324]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:498
		_go_fuzz_dep_.CoverTab[94325]++
													if d.SingleSegment && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:499
			_go_fuzz_dep_.CoverTab[94335]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:499
			return uint64(len(d.history.b)) > d.o.maxDecodedSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:499
			// _ = "end of CoverTab[94335]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:499
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:499
			_go_fuzz_dep_.CoverTab[94336]++
														println("runDecoder: single segment and", uint64(len(d.history.b)), ">", d.o.maxDecodedSize)
														err = ErrFrameSizeExceeded
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:502
			// _ = "end of CoverTab[94336]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:503
			_go_fuzz_dep_.CoverTab[94337]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:503
			// _ = "end of CoverTab[94337]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:503
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:503
		// _ = "end of CoverTab[94325]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:504
	// _ = "end of CoverTab[94318]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:504
	_go_fuzz_dep_.CoverTab[94319]++
												dst = d.history.b
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:506
		_go_fuzz_dep_.CoverTab[94338]++
													if d.HasCheckSum {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:507
			_go_fuzz_dep_.CoverTab[94339]++
														var n int
														n, err = d.crc.Write(dst[crcStart:])
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:510
				_go_fuzz_dep_.CoverTab[94340]++
															if n != len(dst)-crcStart {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:511
					_go_fuzz_dep_.CoverTab[94341]++
																err = io.ErrShortWrite
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:512
					// _ = "end of CoverTab[94341]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:513
					_go_fuzz_dep_.CoverTab[94342]++
																err = d.checkCRC()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:514
					// _ = "end of CoverTab[94342]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:515
				// _ = "end of CoverTab[94340]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:516
				_go_fuzz_dep_.CoverTab[94343]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:516
				// _ = "end of CoverTab[94343]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:516
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:516
			// _ = "end of CoverTab[94339]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:517
			_go_fuzz_dep_.CoverTab[94344]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:517
			// _ = "end of CoverTab[94344]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:517
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:517
		// _ = "end of CoverTab[94338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:518
		_go_fuzz_dep_.CoverTab[94345]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:518
		// _ = "end of CoverTab[94345]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:518
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:518
	// _ = "end of CoverTab[94319]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:518
	_go_fuzz_dep_.CoverTab[94320]++
												d.history.b = saved
												return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:520
	// _ = "end of CoverTab[94320]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:521
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/framedec.go:521
var _ = _go_fuzz_dep_.CoverTab
