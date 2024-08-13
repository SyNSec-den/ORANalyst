// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:5
)

import (
	"crypto/rand"
	"fmt"
	"io"
	rdebug "runtime/debug"
	"sync"

	"github.com/klauspost/compress/zstd/internal/xxhash"
)

// Encoder provides encoding to Zstandard.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:17
// An Encoder can be used for either compressing a stream via the
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:17
// io.WriteCloser interface supported by the Encoder or as multiple independent
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:17
// tasks via the EncodeAll function.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:17
// Smaller encodes are encouraged to use the EncodeAll function.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:17
// Use NewWriter to create a new instance.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:23
type Encoder struct {
	o		encoderOptions
	encoders	chan encoder
	state		encoderState
	init		sync.Once
}

type encoder interface {
	Encode(blk *blockEnc, src []byte)
	EncodeNoHist(blk *blockEnc, src []byte)
	Block() *blockEnc
	CRC() *xxhash.Digest
	AppendCRC([]byte) []byte
	WindowSize(size int64) int32
	UseBlock(*blockEnc)
	Reset(d *dict, singleBlock bool)
}

type encoderState struct {
	w			io.Writer
	filling			[]byte
	current			[]byte
	previous		[]byte
	encoder			encoder
	writing			*blockEnc
	err			error
	writeErr		error
	nWritten		int64
	nInput			int64
	frameContentSize	int64
	headerWritten		bool
	eofWritten		bool
	fullFrameWritten	bool

	// This waitgroup indicates an encode is running.
	wg	sync.WaitGroup
	// This waitgroup indicates we have a block encoding/writing.
	wWg	sync.WaitGroup
}

// NewWriter will create a new Zstandard encoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:63
// If the encoder will be used for encoding blocks a nil writer can be used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:65
func NewWriter(w io.Writer, opts ...EOption) (*Encoder, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:65
	_go_fuzz_dep_.CoverTab[93729]++
												initPredefined()
												var e Encoder
												e.o.setDefault()
												for _, o := range opts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:69
		_go_fuzz_dep_.CoverTab[93732]++
													err := o(&e.o)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:71
			_go_fuzz_dep_.CoverTab[93733]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:72
			// _ = "end of CoverTab[93733]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:73
			_go_fuzz_dep_.CoverTab[93734]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:73
			// _ = "end of CoverTab[93734]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:73
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:73
		// _ = "end of CoverTab[93732]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:74
	// _ = "end of CoverTab[93729]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:74
	_go_fuzz_dep_.CoverTab[93730]++
												if w != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:75
		_go_fuzz_dep_.CoverTab[93735]++
													e.Reset(w)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:76
		// _ = "end of CoverTab[93735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:77
		_go_fuzz_dep_.CoverTab[93736]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:77
		// _ = "end of CoverTab[93736]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:77
	// _ = "end of CoverTab[93730]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:77
	_go_fuzz_dep_.CoverTab[93731]++
												return &e, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:78
	// _ = "end of CoverTab[93731]"
}

func (e *Encoder) initialize() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:81
	_go_fuzz_dep_.CoverTab[93737]++
												if e.o.concurrent == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:82
		_go_fuzz_dep_.CoverTab[93739]++
													e.o.setDefault()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:83
		// _ = "end of CoverTab[93739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:84
		_go_fuzz_dep_.CoverTab[93740]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:84
		// _ = "end of CoverTab[93740]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:84
	// _ = "end of CoverTab[93737]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:84
	_go_fuzz_dep_.CoverTab[93738]++
												e.encoders = make(chan encoder, e.o.concurrent)
												for i := 0; i < e.o.concurrent; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:86
		_go_fuzz_dep_.CoverTab[93741]++
													enc := e.o.encoder()
													e.encoders <- enc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:88
		// _ = "end of CoverTab[93741]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:89
	// _ = "end of CoverTab[93738]"
}

// Reset will re-initialize the writer and new writes will encode to the supplied writer
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:92
// as a new, independent stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:94
func (e *Encoder) Reset(w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:94
	_go_fuzz_dep_.CoverTab[93742]++
												s := &e.state
												s.wg.Wait()
												s.wWg.Wait()
												if cap(s.filling) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:98
		_go_fuzz_dep_.CoverTab[93748]++
													s.filling = make([]byte, 0, e.o.blockSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:99
		// _ = "end of CoverTab[93748]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:100
		_go_fuzz_dep_.CoverTab[93749]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:100
		// _ = "end of CoverTab[93749]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:100
	// _ = "end of CoverTab[93742]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:100
	_go_fuzz_dep_.CoverTab[93743]++
												if cap(s.current) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:101
		_go_fuzz_dep_.CoverTab[93750]++
													s.current = make([]byte, 0, e.o.blockSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:102
		// _ = "end of CoverTab[93750]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:103
		_go_fuzz_dep_.CoverTab[93751]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:103
		// _ = "end of CoverTab[93751]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:103
	// _ = "end of CoverTab[93743]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:103
	_go_fuzz_dep_.CoverTab[93744]++
												if cap(s.previous) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:104
		_go_fuzz_dep_.CoverTab[93752]++
													s.previous = make([]byte, 0, e.o.blockSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:105
		// _ = "end of CoverTab[93752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:106
		_go_fuzz_dep_.CoverTab[93753]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:106
		// _ = "end of CoverTab[93753]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:106
	// _ = "end of CoverTab[93744]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:106
	_go_fuzz_dep_.CoverTab[93745]++
												if s.encoder == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:107
		_go_fuzz_dep_.CoverTab[93754]++
													s.encoder = e.o.encoder()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:108
		// _ = "end of CoverTab[93754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:109
		_go_fuzz_dep_.CoverTab[93755]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:109
		// _ = "end of CoverTab[93755]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:109
	// _ = "end of CoverTab[93745]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:109
	_go_fuzz_dep_.CoverTab[93746]++
												if s.writing == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:110
		_go_fuzz_dep_.CoverTab[93756]++
													s.writing = &blockEnc{lowMem: e.o.lowMem}
													s.writing.init()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:112
		// _ = "end of CoverTab[93756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:113
		_go_fuzz_dep_.CoverTab[93757]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:113
		// _ = "end of CoverTab[93757]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:113
	// _ = "end of CoverTab[93746]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:113
	_go_fuzz_dep_.CoverTab[93747]++
												s.writing.initNewEncode()
												s.filling = s.filling[:0]
												s.current = s.current[:0]
												s.previous = s.previous[:0]
												s.encoder.Reset(e.o.dict, false)
												s.headerWritten = false
												s.eofWritten = false
												s.fullFrameWritten = false
												s.w = w
												s.err = nil
												s.nWritten = 0
												s.nInput = 0
												s.writeErr = nil
												s.frameContentSize = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:127
	// _ = "end of CoverTab[93747]"
}

// ResetContentSize will reset and set a content size for the next stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:130
// If the bytes written does not match the size given an error will be returned
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:130
// when calling Close().
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:130
// This is removed when Reset is called.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:130
// Sizes <= 0 results in no content size set.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:135
func (e *Encoder) ResetContentSize(w io.Writer, size int64) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:135
	_go_fuzz_dep_.CoverTab[93758]++
												e.Reset(w)
												if size >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:137
		_go_fuzz_dep_.CoverTab[93759]++
													e.state.frameContentSize = size
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:138
		// _ = "end of CoverTab[93759]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:139
		_go_fuzz_dep_.CoverTab[93760]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:139
		// _ = "end of CoverTab[93760]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:139
	// _ = "end of CoverTab[93758]"
}

// Write data to the encoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:142
// Input data will be buffered and as the buffer fills up
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:142
// content will be compressed and written to the output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:142
// When done writing, use Close to flush the remaining output
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:142
// and write CRC if requested.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:147
func (e *Encoder) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:147
	_go_fuzz_dep_.CoverTab[93761]++
												s := &e.state
												for len(p) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:149
		_go_fuzz_dep_.CoverTab[93763]++
													if len(p)+len(s.filling) < e.o.blockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:150
			_go_fuzz_dep_.CoverTab[93769]++
														if e.o.crc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:151
				_go_fuzz_dep_.CoverTab[93771]++
															_, _ = s.encoder.CRC().Write(p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:152
				// _ = "end of CoverTab[93771]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:153
				_go_fuzz_dep_.CoverTab[93772]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:153
				// _ = "end of CoverTab[93772]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:153
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:153
			// _ = "end of CoverTab[93769]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:153
			_go_fuzz_dep_.CoverTab[93770]++
														s.filling = append(s.filling, p...)
														return n + len(p), nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:155
			// _ = "end of CoverTab[93770]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:156
			_go_fuzz_dep_.CoverTab[93773]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:156
			// _ = "end of CoverTab[93773]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:156
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:156
		// _ = "end of CoverTab[93763]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:156
		_go_fuzz_dep_.CoverTab[93764]++
													add := p
													if len(p)+len(s.filling) > e.o.blockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:158
			_go_fuzz_dep_.CoverTab[93774]++
														add = add[:e.o.blockSize-len(s.filling)]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:159
			// _ = "end of CoverTab[93774]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:160
			_go_fuzz_dep_.CoverTab[93775]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:160
			// _ = "end of CoverTab[93775]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:160
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:160
		// _ = "end of CoverTab[93764]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:160
		_go_fuzz_dep_.CoverTab[93765]++
													if e.o.crc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:161
			_go_fuzz_dep_.CoverTab[93776]++
														_, _ = s.encoder.CRC().Write(add)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:162
			// _ = "end of CoverTab[93776]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:163
			_go_fuzz_dep_.CoverTab[93777]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:163
			// _ = "end of CoverTab[93777]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:163
		// _ = "end of CoverTab[93765]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:163
		_go_fuzz_dep_.CoverTab[93766]++
													s.filling = append(s.filling, add...)
													p = p[len(add):]
													n += len(add)
													if len(s.filling) < e.o.blockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:167
			_go_fuzz_dep_.CoverTab[93778]++
														return n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:168
			// _ = "end of CoverTab[93778]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:169
			_go_fuzz_dep_.CoverTab[93779]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:169
			// _ = "end of CoverTab[93779]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:169
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:169
		// _ = "end of CoverTab[93766]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:169
		_go_fuzz_dep_.CoverTab[93767]++
													err := e.nextBlock(false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:171
			_go_fuzz_dep_.CoverTab[93780]++
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:172
			// _ = "end of CoverTab[93780]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:173
			_go_fuzz_dep_.CoverTab[93781]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:173
			// _ = "end of CoverTab[93781]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:173
		// _ = "end of CoverTab[93767]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:173
		_go_fuzz_dep_.CoverTab[93768]++
													if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:174
			_go_fuzz_dep_.CoverTab[93782]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:174
			return len(s.filling) > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:174
			// _ = "end of CoverTab[93782]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:174
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:174
			_go_fuzz_dep_.CoverTab[93783]++
														panic(len(s.filling))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:175
			// _ = "end of CoverTab[93783]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:176
			_go_fuzz_dep_.CoverTab[93784]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:176
			// _ = "end of CoverTab[93784]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:176
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:176
		// _ = "end of CoverTab[93768]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:177
	// _ = "end of CoverTab[93761]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:177
	_go_fuzz_dep_.CoverTab[93762]++
												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:178
	// _ = "end of CoverTab[93762]"
}

// nextBlock will synchronize and start compressing input in e.state.filling.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:181
// If an error has occurred during encoding it will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:183
func (e *Encoder) nextBlock(final bool) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:183
	_go_fuzz_dep_.CoverTab[93785]++
												s := &e.state

												s.wg.Wait()
												if s.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:187
		_go_fuzz_dep_.CoverTab[93792]++
													return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:188
		// _ = "end of CoverTab[93792]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:189
		_go_fuzz_dep_.CoverTab[93793]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:189
		// _ = "end of CoverTab[93793]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:189
	// _ = "end of CoverTab[93785]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:189
	_go_fuzz_dep_.CoverTab[93786]++
												if len(s.filling) > e.o.blockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:190
		_go_fuzz_dep_.CoverTab[93794]++
													return fmt.Errorf("block > maxStoreBlockSize")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:191
		// _ = "end of CoverTab[93794]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:192
		_go_fuzz_dep_.CoverTab[93795]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:192
		// _ = "end of CoverTab[93795]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:192
	// _ = "end of CoverTab[93786]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:192
	_go_fuzz_dep_.CoverTab[93787]++
												if !s.headerWritten {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:193
		_go_fuzz_dep_.CoverTab[93796]++

													if final && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			_go_fuzz_dep_.CoverTab[93801]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			return len(s.filling) == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			// _ = "end of CoverTab[93801]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			_go_fuzz_dep_.CoverTab[93802]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			return !e.o.fullZero
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			// _ = "end of CoverTab[93802]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:195
			_go_fuzz_dep_.CoverTab[93803]++
														s.headerWritten = true
														s.fullFrameWritten = true
														s.eofWritten = true
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:199
			// _ = "end of CoverTab[93803]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:200
			_go_fuzz_dep_.CoverTab[93804]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:200
			// _ = "end of CoverTab[93804]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:200
		// _ = "end of CoverTab[93796]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:200
		_go_fuzz_dep_.CoverTab[93797]++
													if final && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:201
			_go_fuzz_dep_.CoverTab[93805]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:201
			return len(s.filling) > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:201
			// _ = "end of CoverTab[93805]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:201
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:201
			_go_fuzz_dep_.CoverTab[93806]++
														s.current = e.EncodeAll(s.filling, s.current[:0])
														var n2 int
														n2, s.err = s.w.Write(s.current)
														if s.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:205
				_go_fuzz_dep_.CoverTab[93808]++
															return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:206
				// _ = "end of CoverTab[93808]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:207
				_go_fuzz_dep_.CoverTab[93809]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:207
				// _ = "end of CoverTab[93809]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:207
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:207
			// _ = "end of CoverTab[93806]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:207
			_go_fuzz_dep_.CoverTab[93807]++
														s.nWritten += int64(n2)
														s.nInput += int64(len(s.filling))
														s.current = s.current[:0]
														s.filling = s.filling[:0]
														s.headerWritten = true
														s.fullFrameWritten = true
														s.eofWritten = true
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:215
			// _ = "end of CoverTab[93807]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:216
			_go_fuzz_dep_.CoverTab[93810]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:216
			// _ = "end of CoverTab[93810]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:216
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:216
		// _ = "end of CoverTab[93797]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:216
		_go_fuzz_dep_.CoverTab[93798]++

													var tmp [maxHeaderSize]byte
													fh := frameHeader{
			ContentSize:	uint64(s.frameContentSize),
			WindowSize:	uint32(s.encoder.WindowSize(s.frameContentSize)),
			SingleSegment:	false,
			Checksum:	e.o.crc,
			DictID:		e.o.dict.ID(),
		}

		dst, err := fh.appendTo(tmp[:0])
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:228
			_go_fuzz_dep_.CoverTab[93811]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:229
			// _ = "end of CoverTab[93811]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:230
			_go_fuzz_dep_.CoverTab[93812]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:230
			// _ = "end of CoverTab[93812]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:230
		// _ = "end of CoverTab[93798]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:230
		_go_fuzz_dep_.CoverTab[93799]++
													s.headerWritten = true
													s.wWg.Wait()
													var n2 int
													n2, s.err = s.w.Write(dst)
													if s.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:235
			_go_fuzz_dep_.CoverTab[93813]++
														return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:236
			// _ = "end of CoverTab[93813]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:237
			_go_fuzz_dep_.CoverTab[93814]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:237
			// _ = "end of CoverTab[93814]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:237
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:237
		// _ = "end of CoverTab[93799]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:237
		_go_fuzz_dep_.CoverTab[93800]++
													s.nWritten += int64(n2)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:238
		// _ = "end of CoverTab[93800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:239
		_go_fuzz_dep_.CoverTab[93815]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:239
		// _ = "end of CoverTab[93815]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:239
	// _ = "end of CoverTab[93787]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:239
	_go_fuzz_dep_.CoverTab[93788]++
												if s.eofWritten {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:240
		_go_fuzz_dep_.CoverTab[93816]++

													final = false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:242
		// _ = "end of CoverTab[93816]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:243
		_go_fuzz_dep_.CoverTab[93817]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:243
		// _ = "end of CoverTab[93817]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:243
	// _ = "end of CoverTab[93788]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:243
	_go_fuzz_dep_.CoverTab[93789]++

												if len(s.filling) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:245
		_go_fuzz_dep_.CoverTab[93818]++

													if final {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:247
			_go_fuzz_dep_.CoverTab[93820]++
														enc := s.encoder
														blk := enc.Block()
														blk.reset(nil)
														blk.last = true
														blk.encodeRaw(nil)
														s.wWg.Wait()
														_, s.err = s.w.Write(blk.output)
														s.nWritten += int64(len(blk.output))
														s.eofWritten = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:256
			// _ = "end of CoverTab[93820]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:257
			_go_fuzz_dep_.CoverTab[93821]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:257
			// _ = "end of CoverTab[93821]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:257
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:257
		// _ = "end of CoverTab[93818]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:257
		_go_fuzz_dep_.CoverTab[93819]++
													return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:258
		// _ = "end of CoverTab[93819]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:259
		_go_fuzz_dep_.CoverTab[93822]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:259
		// _ = "end of CoverTab[93822]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:259
	// _ = "end of CoverTab[93789]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:259
	_go_fuzz_dep_.CoverTab[93790]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:262
	s.filling, s.current, s.previous = s.previous[:0], s.filling, s.current
												s.nInput += int64(len(s.current))
												s.wg.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:264
	_curRoutineNum106_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:264
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum106_)
												go func(src []byte) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:265
		_go_fuzz_dep_.CoverTab[93823]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:265
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:265
			_go_fuzz_dep_.CoverTab[93828]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:265
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum106_)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:265
			// _ = "end of CoverTab[93828]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:265
		}()
													if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:266
			_go_fuzz_dep_.CoverTab[93829]++
														println("Adding block,", len(src), "bytes, final:", final)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:267
			// _ = "end of CoverTab[93829]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:268
			_go_fuzz_dep_.CoverTab[93830]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:268
			// _ = "end of CoverTab[93830]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:268
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:268
		// _ = "end of CoverTab[93823]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:268
		_go_fuzz_dep_.CoverTab[93824]++
													defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:269
			_go_fuzz_dep_.CoverTab[93831]++
														if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:270
				_go_fuzz_dep_.CoverTab[93833]++
															s.err = fmt.Errorf("panic while encoding: %v", r)
															rdebug.PrintStack()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:272
				// _ = "end of CoverTab[93833]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:273
				_go_fuzz_dep_.CoverTab[93834]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:273
				// _ = "end of CoverTab[93834]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:273
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:273
			// _ = "end of CoverTab[93831]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:273
			_go_fuzz_dep_.CoverTab[93832]++
														s.wg.Done()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:274
			// _ = "end of CoverTab[93832]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:275
		// _ = "end of CoverTab[93824]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:275
		_go_fuzz_dep_.CoverTab[93825]++
													enc := s.encoder
													blk := enc.Block()
													enc.Encode(blk, src)
													blk.last = final
													if final {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:280
			_go_fuzz_dep_.CoverTab[93835]++
														s.eofWritten = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:281
			// _ = "end of CoverTab[93835]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:282
			_go_fuzz_dep_.CoverTab[93836]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:282
			// _ = "end of CoverTab[93836]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:282
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:282
		// _ = "end of CoverTab[93825]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:282
		_go_fuzz_dep_.CoverTab[93826]++

													s.wWg.Wait()
													if s.writeErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:285
			_go_fuzz_dep_.CoverTab[93837]++
														s.err = s.writeErr
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:287
			// _ = "end of CoverTab[93837]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:288
			_go_fuzz_dep_.CoverTab[93838]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:288
			// _ = "end of CoverTab[93838]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:288
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:288
		// _ = "end of CoverTab[93826]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:288
		_go_fuzz_dep_.CoverTab[93827]++

													blk.swapEncoders(s.writing)

													enc.UseBlock(s.writing)
													s.writing = blk
													s.wWg.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:294
		_curRoutineNum107_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:294
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum107_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:295
			_go_fuzz_dep_.CoverTab[93839]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:295
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:295
				_go_fuzz_dep_.CoverTab[93843]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:295
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum107_)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:295
				// _ = "end of CoverTab[93843]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:295
			}()
														defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:296
				_go_fuzz_dep_.CoverTab[93844]++
															if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:297
					_go_fuzz_dep_.CoverTab[93846]++
																s.writeErr = fmt.Errorf("panic while encoding/writing: %v", r)
																rdebug.PrintStack()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:299
					// _ = "end of CoverTab[93846]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:300
					_go_fuzz_dep_.CoverTab[93847]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:300
					// _ = "end of CoverTab[93847]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:300
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:300
				// _ = "end of CoverTab[93844]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:300
				_go_fuzz_dep_.CoverTab[93845]++
															s.wWg.Done()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:301
				// _ = "end of CoverTab[93845]"
			}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:302
			// _ = "end of CoverTab[93839]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:302
			_go_fuzz_dep_.CoverTab[93840]++
														err := errIncompressible

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:306
			if len(src) != len(blk.literals) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:306
				_go_fuzz_dep_.CoverTab[93848]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:306
				return len(src) != e.o.blockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:306
				// _ = "end of CoverTab[93848]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:306
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:306
				_go_fuzz_dep_.CoverTab[93849]++
															err = blk.encode(src, e.o.noEntropy, !e.o.allLitEntropy)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:307
				// _ = "end of CoverTab[93849]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:308
				_go_fuzz_dep_.CoverTab[93850]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:308
				// _ = "end of CoverTab[93850]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:308
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:308
			// _ = "end of CoverTab[93840]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:308
			_go_fuzz_dep_.CoverTab[93841]++
														switch err {
			case errIncompressible:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:310
				_go_fuzz_dep_.CoverTab[93851]++
															if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:311
					_go_fuzz_dep_.CoverTab[93855]++
																println("Storing incompressible block as raw")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:312
					// _ = "end of CoverTab[93855]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:313
					_go_fuzz_dep_.CoverTab[93856]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:313
					// _ = "end of CoverTab[93856]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:313
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:313
				// _ = "end of CoverTab[93851]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:313
				_go_fuzz_dep_.CoverTab[93852]++
															blk.encodeRaw(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:314
				// _ = "end of CoverTab[93852]"

			case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:316
				_go_fuzz_dep_.CoverTab[93853]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:316
				// _ = "end of CoverTab[93853]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:317
				_go_fuzz_dep_.CoverTab[93854]++
															s.writeErr = err
															return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:319
				// _ = "end of CoverTab[93854]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:320
			// _ = "end of CoverTab[93841]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:320
			_go_fuzz_dep_.CoverTab[93842]++
														_, s.writeErr = s.w.Write(blk.output)
														s.nWritten += int64(len(blk.output))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:322
			// _ = "end of CoverTab[93842]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:323
		// _ = "end of CoverTab[93827]"
	}(s.current)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:324
	// _ = "end of CoverTab[93790]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:324
	_go_fuzz_dep_.CoverTab[93791]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:325
	// _ = "end of CoverTab[93791]"
}

// ReadFrom reads data from r until EOF or error.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:328
// The return value n is the number of bytes read.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:328
// Any error except io.EOF encountered during the read is also returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:328
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:328
// The Copy function uses ReaderFrom if available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:333
func (e *Encoder) ReadFrom(r io.Reader) (n int64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:333
	_go_fuzz_dep_.CoverTab[93857]++
												if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:334
		_go_fuzz_dep_.CoverTab[93860]++
													println("Using ReadFrom")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:335
		// _ = "end of CoverTab[93860]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:336
		_go_fuzz_dep_.CoverTab[93861]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:336
		// _ = "end of CoverTab[93861]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:336
	// _ = "end of CoverTab[93857]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:336
	_go_fuzz_dep_.CoverTab[93858]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:339
	if len(e.state.filling) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:339
		_go_fuzz_dep_.CoverTab[93862]++
													if err := e.nextBlock(false); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:340
			_go_fuzz_dep_.CoverTab[93863]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:341
			// _ = "end of CoverTab[93863]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:342
			_go_fuzz_dep_.CoverTab[93864]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:342
			// _ = "end of CoverTab[93864]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:342
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:342
		// _ = "end of CoverTab[93862]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:343
		_go_fuzz_dep_.CoverTab[93865]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:343
		// _ = "end of CoverTab[93865]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:343
	// _ = "end of CoverTab[93858]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:343
	_go_fuzz_dep_.CoverTab[93859]++
												e.state.filling = e.state.filling[:e.o.blockSize]
												src := e.state.filling
												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:346
		_go_fuzz_dep_.CoverTab[93866]++
													n2, err := r.Read(src)
													if e.o.crc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:348
			_go_fuzz_dep_.CoverTab[93871]++
														_, _ = e.state.encoder.CRC().Write(src[:n2])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:349
			// _ = "end of CoverTab[93871]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:350
			_go_fuzz_dep_.CoverTab[93872]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:350
			// _ = "end of CoverTab[93872]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:350
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:350
		// _ = "end of CoverTab[93866]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:350
		_go_fuzz_dep_.CoverTab[93867]++

													src = src[n2:]
													n += int64(n2)
													switch err {
		case io.EOF:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:355
			_go_fuzz_dep_.CoverTab[93873]++
														e.state.filling = e.state.filling[:len(e.state.filling)-len(src)]
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:357
				_go_fuzz_dep_.CoverTab[93878]++
															println("ReadFrom: got EOF final block:", len(e.state.filling))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:358
				// _ = "end of CoverTab[93878]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:359
				_go_fuzz_dep_.CoverTab[93879]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:359
				// _ = "end of CoverTab[93879]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:359
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:359
			// _ = "end of CoverTab[93873]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:359
			_go_fuzz_dep_.CoverTab[93874]++
														return n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:360
			// _ = "end of CoverTab[93874]"
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:361
			_go_fuzz_dep_.CoverTab[93875]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:361
			// _ = "end of CoverTab[93875]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:362
			_go_fuzz_dep_.CoverTab[93876]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:363
				_go_fuzz_dep_.CoverTab[93880]++
															println("ReadFrom: got error:", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:364
				// _ = "end of CoverTab[93880]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:365
				_go_fuzz_dep_.CoverTab[93881]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:365
				// _ = "end of CoverTab[93881]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:365
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:365
			// _ = "end of CoverTab[93876]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:365
			_go_fuzz_dep_.CoverTab[93877]++
														e.state.err = err
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:367
			// _ = "end of CoverTab[93877]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:368
		// _ = "end of CoverTab[93867]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:368
		_go_fuzz_dep_.CoverTab[93868]++
													if len(src) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:369
			_go_fuzz_dep_.CoverTab[93882]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:370
				_go_fuzz_dep_.CoverTab[93884]++
															println("ReadFrom: got space left in source:", len(src))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:371
				// _ = "end of CoverTab[93884]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:372
				_go_fuzz_dep_.CoverTab[93885]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:372
				// _ = "end of CoverTab[93885]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:372
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:372
			// _ = "end of CoverTab[93882]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:372
			_go_fuzz_dep_.CoverTab[93883]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:373
			// _ = "end of CoverTab[93883]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:374
			_go_fuzz_dep_.CoverTab[93886]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:374
			// _ = "end of CoverTab[93886]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:374
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:374
		// _ = "end of CoverTab[93868]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:374
		_go_fuzz_dep_.CoverTab[93869]++
													err = e.nextBlock(false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:376
			_go_fuzz_dep_.CoverTab[93887]++
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:377
			// _ = "end of CoverTab[93887]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:378
			_go_fuzz_dep_.CoverTab[93888]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:378
			// _ = "end of CoverTab[93888]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:378
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:378
		// _ = "end of CoverTab[93869]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:378
		_go_fuzz_dep_.CoverTab[93870]++
													e.state.filling = e.state.filling[:e.o.blockSize]
													src = e.state.filling
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:380
		// _ = "end of CoverTab[93870]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:381
	// _ = "end of CoverTab[93859]"
}

// Flush will send the currently written data to output
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:384
// and block until everything has been written.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:384
// This should only be used on rare occasions where pushing the currently queued data is critical.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:387
func (e *Encoder) Flush() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:387
	_go_fuzz_dep_.CoverTab[93889]++
												s := &e.state
												if len(s.filling) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:389
		_go_fuzz_dep_.CoverTab[93892]++
													err := e.nextBlock(false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:391
			_go_fuzz_dep_.CoverTab[93893]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:392
			// _ = "end of CoverTab[93893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:393
			_go_fuzz_dep_.CoverTab[93894]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:393
			// _ = "end of CoverTab[93894]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:393
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:393
		// _ = "end of CoverTab[93892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:394
		_go_fuzz_dep_.CoverTab[93895]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:394
		// _ = "end of CoverTab[93895]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:394
	// _ = "end of CoverTab[93889]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:394
	_go_fuzz_dep_.CoverTab[93890]++
												s.wg.Wait()
												s.wWg.Wait()
												if s.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:397
		_go_fuzz_dep_.CoverTab[93896]++
													return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:398
		// _ = "end of CoverTab[93896]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:399
		_go_fuzz_dep_.CoverTab[93897]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:399
		// _ = "end of CoverTab[93897]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:399
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:399
	// _ = "end of CoverTab[93890]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:399
	_go_fuzz_dep_.CoverTab[93891]++
												return s.writeErr
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:400
	// _ = "end of CoverTab[93891]"
}

// Close will flush the final output and close the stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:403
// The function will block until everything has been written.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:403
// The Encoder can still be re-used after calling this.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:406
func (e *Encoder) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:406
	_go_fuzz_dep_.CoverTab[93898]++
												s := &e.state
												if s.encoder == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:408
		_go_fuzz_dep_.CoverTab[93907]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:409
		// _ = "end of CoverTab[93907]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:410
		_go_fuzz_dep_.CoverTab[93908]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:410
		// _ = "end of CoverTab[93908]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:410
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:410
	// _ = "end of CoverTab[93898]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:410
	_go_fuzz_dep_.CoverTab[93899]++
												err := e.nextBlock(true)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:412
		_go_fuzz_dep_.CoverTab[93909]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:413
		// _ = "end of CoverTab[93909]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:414
		_go_fuzz_dep_.CoverTab[93910]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:414
		// _ = "end of CoverTab[93910]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:414
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:414
	// _ = "end of CoverTab[93899]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:414
	_go_fuzz_dep_.CoverTab[93900]++
												if s.frameContentSize > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:415
		_go_fuzz_dep_.CoverTab[93911]++
													if s.nInput != s.frameContentSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:416
			_go_fuzz_dep_.CoverTab[93912]++
														return fmt.Errorf("frame content size %d given, but %d bytes was written", s.frameContentSize, s.nInput)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:417
			// _ = "end of CoverTab[93912]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:418
			_go_fuzz_dep_.CoverTab[93913]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:418
			// _ = "end of CoverTab[93913]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:418
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:418
		// _ = "end of CoverTab[93911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:419
		_go_fuzz_dep_.CoverTab[93914]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:419
		// _ = "end of CoverTab[93914]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:419
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:419
	// _ = "end of CoverTab[93900]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:419
	_go_fuzz_dep_.CoverTab[93901]++
												if e.state.fullFrameWritten {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:420
		_go_fuzz_dep_.CoverTab[93915]++
													return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:421
		// _ = "end of CoverTab[93915]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:422
		_go_fuzz_dep_.CoverTab[93916]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:422
		// _ = "end of CoverTab[93916]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:422
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:422
	// _ = "end of CoverTab[93901]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:422
	_go_fuzz_dep_.CoverTab[93902]++
												s.wg.Wait()
												s.wWg.Wait()

												if s.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:426
		_go_fuzz_dep_.CoverTab[93917]++
													return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:427
		// _ = "end of CoverTab[93917]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:428
		_go_fuzz_dep_.CoverTab[93918]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:428
		// _ = "end of CoverTab[93918]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:428
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:428
	// _ = "end of CoverTab[93902]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:428
	_go_fuzz_dep_.CoverTab[93903]++
												if s.writeErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:429
		_go_fuzz_dep_.CoverTab[93919]++
													return s.writeErr
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:430
		// _ = "end of CoverTab[93919]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:431
		_go_fuzz_dep_.CoverTab[93920]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:431
		// _ = "end of CoverTab[93920]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:431
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:431
	// _ = "end of CoverTab[93903]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:431
	_go_fuzz_dep_.CoverTab[93904]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:434
	if e.o.crc && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:434
		_go_fuzz_dep_.CoverTab[93921]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:434
		return s.err == nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:434
		// _ = "end of CoverTab[93921]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:434
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:434
		_go_fuzz_dep_.CoverTab[93922]++
													// heap alloc.
													var tmp [4]byte
													_, s.err = s.w.Write(s.encoder.AppendCRC(tmp[:0]))
													s.nWritten += 4
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:438
		// _ = "end of CoverTab[93922]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:439
		_go_fuzz_dep_.CoverTab[93923]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:439
		// _ = "end of CoverTab[93923]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:439
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:439
	// _ = "end of CoverTab[93904]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:439
	_go_fuzz_dep_.CoverTab[93905]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:442
	if s.err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:442
		_go_fuzz_dep_.CoverTab[93924]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:442
		return e.o.pad > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:442
		// _ = "end of CoverTab[93924]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:442
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:442
		_go_fuzz_dep_.CoverTab[93925]++
													add := calcSkippableFrame(s.nWritten, int64(e.o.pad))
													frame, err := skippableFrame(s.filling[:0], add, rand.Reader)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:445
			_go_fuzz_dep_.CoverTab[93927]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:446
			// _ = "end of CoverTab[93927]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:447
			_go_fuzz_dep_.CoverTab[93928]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:447
			// _ = "end of CoverTab[93928]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:447
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:447
		// _ = "end of CoverTab[93925]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:447
		_go_fuzz_dep_.CoverTab[93926]++
													_, s.err = s.w.Write(frame)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:448
		// _ = "end of CoverTab[93926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:449
		_go_fuzz_dep_.CoverTab[93929]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:449
		// _ = "end of CoverTab[93929]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:449
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:449
	// _ = "end of CoverTab[93905]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:449
	_go_fuzz_dep_.CoverTab[93906]++
												return s.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:450
	// _ = "end of CoverTab[93906]"
}

// EncodeAll will encode all input in src and append it to dst.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:453
// This function can be called concurrently, but each call will only run on a single goroutine.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:453
// If empty input is given, nothing is returned, unless WithZeroFrames is specified.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:453
// Encoded blocks can be concatenated and the result will be the combined input stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:453
// Data compressed with EncodeAll can be decoded with the Decoder,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:453
// using either a stream or DecodeAll.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:459
func (e *Encoder) EncodeAll(src, dst []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:459
	_go_fuzz_dep_.CoverTab[93930]++
												if len(src) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:460
		_go_fuzz_dep_.CoverTab[93939]++
													if e.o.fullZero {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:461
			_go_fuzz_dep_.CoverTab[93941]++

														fh := frameHeader{
				ContentSize:	0,
				WindowSize:	MinWindowSize,
				SingleSegment:	true,

				Checksum:	false,
				DictID:		0,
			}
			dst, _ = fh.appendTo(dst)

														// Write raw block as last one only.
														var blk blockHeader
														blk.setSize(0)
														blk.setType(blockTypeRaw)
														blk.setLast(true)
														dst = blk.appendTo(dst)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:478
			// _ = "end of CoverTab[93941]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:479
			_go_fuzz_dep_.CoverTab[93942]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:479
			// _ = "end of CoverTab[93942]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:479
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:479
		// _ = "end of CoverTab[93939]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:479
		_go_fuzz_dep_.CoverTab[93940]++
													return dst
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:480
		// _ = "end of CoverTab[93940]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:481
		_go_fuzz_dep_.CoverTab[93943]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:481
		// _ = "end of CoverTab[93943]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:481
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:481
	// _ = "end of CoverTab[93930]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:481
	_go_fuzz_dep_.CoverTab[93931]++
												e.init.Do(e.initialize)
												enc := <-e.encoders
												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:484
		_go_fuzz_dep_.CoverTab[93944]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:487
		e.encoders <- enc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:487
		// _ = "end of CoverTab[93944]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:488
	// _ = "end of CoverTab[93931]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:488
	_go_fuzz_dep_.CoverTab[93932]++

												single := len(src) < 1<<20 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:490
		_go_fuzz_dep_.CoverTab[93945]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:490
		return len(src) > MinWindowSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:490
		// _ = "end of CoverTab[93945]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:490
	}()
												if e.o.single != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:491
		_go_fuzz_dep_.CoverTab[93946]++
													single = *e.o.single
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:492
		// _ = "end of CoverTab[93946]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:493
		_go_fuzz_dep_.CoverTab[93947]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:493
		// _ = "end of CoverTab[93947]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:493
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:493
	// _ = "end of CoverTab[93932]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:493
	_go_fuzz_dep_.CoverTab[93933]++
												fh := frameHeader{
		ContentSize:	uint64(len(src)),
		WindowSize:	uint32(enc.WindowSize(int64(len(src)))),
		SingleSegment:	single,
		Checksum:	e.o.crc,
		DictID:		e.o.dict.ID(),
	}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
	if len(dst) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		_go_fuzz_dep_.CoverTab[93948]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		return cap(dst) == 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		// _ = "end of CoverTab[93948]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		_go_fuzz_dep_.CoverTab[93949]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		return len(src) < 1<<20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		// _ = "end of CoverTab[93949]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		_go_fuzz_dep_.CoverTab[93950]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		return !e.o.lowMem
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		// _ = "end of CoverTab[93950]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:503
		_go_fuzz_dep_.CoverTab[93951]++
													dst = make([]byte, 0, len(src))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:504
		// _ = "end of CoverTab[93951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:505
		_go_fuzz_dep_.CoverTab[93952]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:505
		// _ = "end of CoverTab[93952]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:505
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:505
	// _ = "end of CoverTab[93933]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:505
	_go_fuzz_dep_.CoverTab[93934]++
												dst, err := fh.appendTo(dst)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:507
		_go_fuzz_dep_.CoverTab[93953]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:508
		// _ = "end of CoverTab[93953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:509
		_go_fuzz_dep_.CoverTab[93954]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:509
		// _ = "end of CoverTab[93954]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:509
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:509
	// _ = "end of CoverTab[93934]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:509
	_go_fuzz_dep_.CoverTab[93935]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:512
	if len(src) <= maxCompressedBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:512
		_go_fuzz_dep_.CoverTab[93955]++
													enc.Reset(e.o.dict, true)

													if e.o.crc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:515
			_go_fuzz_dep_.CoverTab[93960]++
														_, _ = enc.CRC().Write(src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:516
			// _ = "end of CoverTab[93960]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:517
			_go_fuzz_dep_.CoverTab[93961]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:517
			// _ = "end of CoverTab[93961]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:517
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:517
		// _ = "end of CoverTab[93955]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:517
		_go_fuzz_dep_.CoverTab[93956]++
													blk := enc.Block()
													blk.last = true
													if e.o.dict == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:520
			_go_fuzz_dep_.CoverTab[93962]++
														enc.EncodeNoHist(blk, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:521
			// _ = "end of CoverTab[93962]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:522
			_go_fuzz_dep_.CoverTab[93963]++
														enc.Encode(blk, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:523
			// _ = "end of CoverTab[93963]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:524
		// _ = "end of CoverTab[93956]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:524
		_go_fuzz_dep_.CoverTab[93957]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:528
		err := errIncompressible
		oldout := blk.output
		if len(blk.literals) != len(src) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:530
			_go_fuzz_dep_.CoverTab[93964]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:530
			return len(src) != e.o.blockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:530
			// _ = "end of CoverTab[93964]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:530
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:530
			_go_fuzz_dep_.CoverTab[93965]++

														blk.output = dst
														err = blk.encode(src, e.o.noEntropy, !e.o.allLitEntropy)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:533
			// _ = "end of CoverTab[93965]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:534
			_go_fuzz_dep_.CoverTab[93966]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:534
			// _ = "end of CoverTab[93966]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:534
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:534
		// _ = "end of CoverTab[93957]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:534
		_go_fuzz_dep_.CoverTab[93958]++

													switch err {
		case errIncompressible:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:537
			_go_fuzz_dep_.CoverTab[93967]++
														if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:538
				_go_fuzz_dep_.CoverTab[93971]++
															println("Storing incompressible block as raw")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:539
				// _ = "end of CoverTab[93971]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:540
				_go_fuzz_dep_.CoverTab[93972]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:540
				// _ = "end of CoverTab[93972]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:540
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:540
			// _ = "end of CoverTab[93967]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:540
			_go_fuzz_dep_.CoverTab[93968]++
														dst = blk.encodeRawTo(dst, src)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:541
			// _ = "end of CoverTab[93968]"
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:542
			_go_fuzz_dep_.CoverTab[93969]++
														dst = blk.output
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:543
			// _ = "end of CoverTab[93969]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:544
			_go_fuzz_dep_.CoverTab[93970]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:545
			// _ = "end of CoverTab[93970]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:546
		// _ = "end of CoverTab[93958]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:546
		_go_fuzz_dep_.CoverTab[93959]++
													blk.output = oldout
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:547
		// _ = "end of CoverTab[93959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:548
		_go_fuzz_dep_.CoverTab[93973]++
													enc.Reset(e.o.dict, false)
													blk := enc.Block()
													for len(src) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:551
			_go_fuzz_dep_.CoverTab[93974]++
														todo := src
														if len(todo) > e.o.blockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:553
				_go_fuzz_dep_.CoverTab[93980]++
															todo = todo[:e.o.blockSize]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:554
				// _ = "end of CoverTab[93980]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:555
				_go_fuzz_dep_.CoverTab[93981]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:555
				// _ = "end of CoverTab[93981]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:555
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:555
			// _ = "end of CoverTab[93974]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:555
			_go_fuzz_dep_.CoverTab[93975]++
														src = src[len(todo):]
														if e.o.crc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:557
				_go_fuzz_dep_.CoverTab[93982]++
															_, _ = enc.CRC().Write(todo)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:558
				// _ = "end of CoverTab[93982]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:559
				_go_fuzz_dep_.CoverTab[93983]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:559
				// _ = "end of CoverTab[93983]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:559
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:559
			// _ = "end of CoverTab[93975]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:559
			_go_fuzz_dep_.CoverTab[93976]++
														blk.pushOffsets()
														enc.Encode(blk, todo)
														if len(src) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:562
				_go_fuzz_dep_.CoverTab[93984]++
															blk.last = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:563
				// _ = "end of CoverTab[93984]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:564
				_go_fuzz_dep_.CoverTab[93985]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:564
				// _ = "end of CoverTab[93985]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:564
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:564
			// _ = "end of CoverTab[93976]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:564
			_go_fuzz_dep_.CoverTab[93977]++
														err := errIncompressible

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:568
			if len(blk.literals) != len(todo) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:568
				_go_fuzz_dep_.CoverTab[93986]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:568
				return len(todo) != e.o.blockSize
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:568
				// _ = "end of CoverTab[93986]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:568
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:568
				_go_fuzz_dep_.CoverTab[93987]++
															err = blk.encode(todo, e.o.noEntropy, !e.o.allLitEntropy)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:569
				// _ = "end of CoverTab[93987]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:570
				_go_fuzz_dep_.CoverTab[93988]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:570
				// _ = "end of CoverTab[93988]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:570
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:570
			// _ = "end of CoverTab[93977]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:570
			_go_fuzz_dep_.CoverTab[93978]++

														switch err {
			case errIncompressible:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:573
				_go_fuzz_dep_.CoverTab[93989]++
															if debugEncoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:574
					_go_fuzz_dep_.CoverTab[93993]++
																println("Storing incompressible block as raw")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:575
					// _ = "end of CoverTab[93993]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:576
					_go_fuzz_dep_.CoverTab[93994]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:576
					// _ = "end of CoverTab[93994]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:576
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:576
				// _ = "end of CoverTab[93989]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:576
				_go_fuzz_dep_.CoverTab[93990]++
															dst = blk.encodeRawTo(dst, todo)
															blk.popOffsets()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:578
				// _ = "end of CoverTab[93990]"
			case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:579
				_go_fuzz_dep_.CoverTab[93991]++
															dst = append(dst, blk.output...)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:580
				// _ = "end of CoverTab[93991]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:581
				_go_fuzz_dep_.CoverTab[93992]++
															panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:582
				// _ = "end of CoverTab[93992]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:583
			// _ = "end of CoverTab[93978]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:583
			_go_fuzz_dep_.CoverTab[93979]++
														blk.reset(nil)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:584
			// _ = "end of CoverTab[93979]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:585
		// _ = "end of CoverTab[93973]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:586
	// _ = "end of CoverTab[93935]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:586
	_go_fuzz_dep_.CoverTab[93936]++
												if e.o.crc {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:587
		_go_fuzz_dep_.CoverTab[93995]++
													dst = enc.AppendCRC(dst)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:588
		// _ = "end of CoverTab[93995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:589
		_go_fuzz_dep_.CoverTab[93996]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:589
		// _ = "end of CoverTab[93996]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:589
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:589
	// _ = "end of CoverTab[93936]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:589
	_go_fuzz_dep_.CoverTab[93937]++

												if e.o.pad > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:591
		_go_fuzz_dep_.CoverTab[93997]++
													add := calcSkippableFrame(int64(len(dst)), int64(e.o.pad))
													dst, err = skippableFrame(dst, add, rand.Reader)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:594
			_go_fuzz_dep_.CoverTab[93998]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:595
			// _ = "end of CoverTab[93998]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:596
			_go_fuzz_dep_.CoverTab[93999]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:596
			// _ = "end of CoverTab[93999]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:596
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:596
		// _ = "end of CoverTab[93997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:597
		_go_fuzz_dep_.CoverTab[94000]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:597
		// _ = "end of CoverTab[94000]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:597
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:597
	// _ = "end of CoverTab[93937]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:597
	_go_fuzz_dep_.CoverTab[93938]++
												return dst
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:598
	// _ = "end of CoverTab[93938]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:599
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder.go:599
var _ = _go_fuzz_dep_.CoverTab
