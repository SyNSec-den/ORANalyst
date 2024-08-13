// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:5
)

import (
	"errors"
	"io"
	"sync"
)

// Decoder provides decoding of zstandard streams.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:13
// The decoder has been designed to operate without allocations after a warmup.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:13
// This means that you should store the decoder for best performance.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:13
// To re-use a stream decoder, use the Reset(r io.Reader) error to switch to another stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:13
// A decoder can safely be re-used even if the previous stream failed.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:13
// To release the resources, you must call the Close() function on a decoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:19
type Decoder struct {
	o	decoderOptions

	// Unreferenced decoders, ready for use.
	decoders	chan *blockDec

	// Streams ready to be decoded.
	stream	chan decodeStream

	// Current read position used for Reader functionality.
	current	decoderState

	// Custom dictionaries.
	// Always uses copies.
	dicts	map[uint32]dict

	// streamWg is the waitgroup for all streams
	streamWg	sync.WaitGroup
}

// decoderState is used for maintaining state when the decoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:39
// is used for streaming.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:41
type decoderState struct {
	// current block being written to stream.
	decodeOutput

	// output in order to be written to stream.
	output	chan decodeOutput

	// cancel remaining output.
	cancel	chan struct{}

	flushed	bool
}

var (
	// Check the interfaces we want to support.
	_	= io.WriterTo(&Decoder{})
	_	= io.Reader(&Decoder{})
)

// NewReader creates a new decoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// A nil Reader can be provided in which case Reset can be used to start a decode.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// A Decoder can be used in two modes:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// 1) As a stream, or
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// 2) For stateless decoding using DecodeAll.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// Only a single stream can be decoded concurrently, but the same decoder
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// can run multiple concurrent stateless decodes. It is even possible to
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// use stateless decodes while a stream is being decoded.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// The Reset function can be used to initiate a new stream, which is will considerably
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:60
// reduce the allocations normally caused by NewReader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:74
func NewReader(r io.Reader, opts ...DOption) (*Decoder, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:74
	_go_fuzz_dep_.CoverTab[91684]++
												initPredefined()
												var d Decoder
												d.o.setDefault()
												for _, o := range opts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:78
		_go_fuzz_dep_.CoverTab[91690]++
													err := o(&d.o)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:80
			_go_fuzz_dep_.CoverTab[91691]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:81
			// _ = "end of CoverTab[91691]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:82
			_go_fuzz_dep_.CoverTab[91692]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:82
			// _ = "end of CoverTab[91692]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:82
		// _ = "end of CoverTab[91690]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:83
	// _ = "end of CoverTab[91684]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:83
	_go_fuzz_dep_.CoverTab[91685]++
												d.current.output = make(chan decodeOutput, d.o.concurrent)
												d.current.flushed = true

												if r == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:87
		_go_fuzz_dep_.CoverTab[91693]++
													d.current.err = ErrDecoderNilInput
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:88
		// _ = "end of CoverTab[91693]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:89
		_go_fuzz_dep_.CoverTab[91694]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:89
		// _ = "end of CoverTab[91694]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:89
	// _ = "end of CoverTab[91685]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:89
	_go_fuzz_dep_.CoverTab[91686]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:92
	d.dicts = make(map[uint32]dict, len(d.o.dicts))
	for _, dc := range d.o.dicts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:93
		_go_fuzz_dep_.CoverTab[91695]++
													d.dicts[dc.id] = dc
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:94
		// _ = "end of CoverTab[91695]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:95
	// _ = "end of CoverTab[91686]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:95
	_go_fuzz_dep_.CoverTab[91687]++
												d.o.dicts = nil

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:99
	d.decoders = make(chan *blockDec, d.o.concurrent)
	for i := 0; i < d.o.concurrent; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:100
		_go_fuzz_dep_.CoverTab[91696]++
													dec := newBlockDec(d.o.lowMem)
													dec.localFrame = newFrameDec(d.o)
													d.decoders <- dec
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:103
		// _ = "end of CoverTab[91696]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:104
	// _ = "end of CoverTab[91687]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:104
	_go_fuzz_dep_.CoverTab[91688]++

												if r == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:106
		_go_fuzz_dep_.CoverTab[91697]++
													return &d, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:107
		// _ = "end of CoverTab[91697]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:108
		_go_fuzz_dep_.CoverTab[91698]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:108
		// _ = "end of CoverTab[91698]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:108
	// _ = "end of CoverTab[91688]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:108
	_go_fuzz_dep_.CoverTab[91689]++
												return &d, d.Reset(r)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:109
	// _ = "end of CoverTab[91689]"
}

// Read bytes from the decompressed stream into p.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:112
// Returns the number of bytes written and any error that occurred.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:112
// When the stream is done, io.EOF will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:115
func (d *Decoder) Read(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:115
	_go_fuzz_dep_.CoverTab[91699]++
												var n int
												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:117
		_go_fuzz_dep_.CoverTab[91704]++
													if len(d.current.b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:118
			_go_fuzz_dep_.CoverTab[91707]++
														filled := copy(p, d.current.b)
														p = p[filled:]
														d.current.b = d.current.b[filled:]
														n += filled
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:122
			// _ = "end of CoverTab[91707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:123
			_go_fuzz_dep_.CoverTab[91708]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:123
			// _ = "end of CoverTab[91708]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:123
		// _ = "end of CoverTab[91704]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:123
		_go_fuzz_dep_.CoverTab[91705]++
													if len(p) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:124
			_go_fuzz_dep_.CoverTab[91709]++
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:125
			// _ = "end of CoverTab[91709]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:126
			_go_fuzz_dep_.CoverTab[91710]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:126
			// _ = "end of CoverTab[91710]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:126
		// _ = "end of CoverTab[91705]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:126
		_go_fuzz_dep_.CoverTab[91706]++
													if len(d.current.b) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:127
			_go_fuzz_dep_.CoverTab[91711]++

														if d.current.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:129
				_go_fuzz_dep_.CoverTab[91713]++
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:130
				// _ = "end of CoverTab[91713]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:131
				_go_fuzz_dep_.CoverTab[91714]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:131
				// _ = "end of CoverTab[91714]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:131
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:131
			// _ = "end of CoverTab[91711]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:131
			_go_fuzz_dep_.CoverTab[91712]++
														if !d.nextBlock(n == 0) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:132
				_go_fuzz_dep_.CoverTab[91715]++
															return n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:133
				// _ = "end of CoverTab[91715]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:134
				_go_fuzz_dep_.CoverTab[91716]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:134
				// _ = "end of CoverTab[91716]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:134
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:134
			// _ = "end of CoverTab[91712]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:135
			_go_fuzz_dep_.CoverTab[91717]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:135
			// _ = "end of CoverTab[91717]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:135
		// _ = "end of CoverTab[91706]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:136
	// _ = "end of CoverTab[91699]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:136
	_go_fuzz_dep_.CoverTab[91700]++
												if len(d.current.b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:137
		_go_fuzz_dep_.CoverTab[91718]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:138
			_go_fuzz_dep_.CoverTab[91720]++
														println("returning", n, "still bytes left:", len(d.current.b))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:139
			// _ = "end of CoverTab[91720]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:140
			_go_fuzz_dep_.CoverTab[91721]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:140
			// _ = "end of CoverTab[91721]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:140
		// _ = "end of CoverTab[91718]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:140
		_go_fuzz_dep_.CoverTab[91719]++

													return n, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:142
		// _ = "end of CoverTab[91719]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:143
		_go_fuzz_dep_.CoverTab[91722]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:143
		// _ = "end of CoverTab[91722]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:143
	// _ = "end of CoverTab[91700]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:143
	_go_fuzz_dep_.CoverTab[91701]++
												if d.current.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:144
		_go_fuzz_dep_.CoverTab[91723]++
													d.drainOutput()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:145
		// _ = "end of CoverTab[91723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:146
		_go_fuzz_dep_.CoverTab[91724]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:146
		// _ = "end of CoverTab[91724]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:146
	// _ = "end of CoverTab[91701]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:146
	_go_fuzz_dep_.CoverTab[91702]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:147
		_go_fuzz_dep_.CoverTab[91725]++
													println("returning", n, d.current.err, len(d.decoders))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:148
		// _ = "end of CoverTab[91725]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:149
		_go_fuzz_dep_.CoverTab[91726]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:149
		// _ = "end of CoverTab[91726]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:149
	// _ = "end of CoverTab[91702]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:149
	_go_fuzz_dep_.CoverTab[91703]++
												return n, d.current.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:150
	// _ = "end of CoverTab[91703]"
}

// Reset will reset the decoder the supplied stream after the current has finished processing.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:153
// Note that this functionality cannot be used after Close has been called.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:153
// Reset can be called with a nil reader to release references to the previous reader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:153
// After being called with a nil reader, no other operations than Reset or DecodeAll or Close
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:153
// should be used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:158
func (d *Decoder) Reset(r io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:158
	_go_fuzz_dep_.CoverTab[91727]++
												if d.current.err == ErrDecoderClosed {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:159
		_go_fuzz_dep_.CoverTab[91732]++
													return d.current.err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:160
		// _ = "end of CoverTab[91732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:161
		_go_fuzz_dep_.CoverTab[91733]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:161
		// _ = "end of CoverTab[91733]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:161
	// _ = "end of CoverTab[91727]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:161
	_go_fuzz_dep_.CoverTab[91728]++

												d.drainOutput()

												if r == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:165
		_go_fuzz_dep_.CoverTab[91734]++
													d.current.err = ErrDecoderNilInput
													if len(d.current.b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:167
			_go_fuzz_dep_.CoverTab[91736]++
														d.current.b = d.current.b[:0]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:168
			// _ = "end of CoverTab[91736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:169
			_go_fuzz_dep_.CoverTab[91737]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:169
			// _ = "end of CoverTab[91737]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:169
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:169
		// _ = "end of CoverTab[91734]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:169
		_go_fuzz_dep_.CoverTab[91735]++
													d.current.flushed = true
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:171
		// _ = "end of CoverTab[91735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:172
		_go_fuzz_dep_.CoverTab[91738]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:172
		// _ = "end of CoverTab[91738]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:172
	// _ = "end of CoverTab[91728]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:172
	_go_fuzz_dep_.CoverTab[91729]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:175
	if bb, ok := r.(byter); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:175
		_go_fuzz_dep_.CoverTab[91739]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:175
		return bb.Len() < 5<<20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:175
		// _ = "end of CoverTab[91739]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:175
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:175
		_go_fuzz_dep_.CoverTab[91740]++
													bb2 := bb
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:177
			_go_fuzz_dep_.CoverTab[91745]++
														println("*bytes.Buffer detected, doing sync decode, len:", bb.Len())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:178
			// _ = "end of CoverTab[91745]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:179
			_go_fuzz_dep_.CoverTab[91746]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:179
			// _ = "end of CoverTab[91746]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:179
		// _ = "end of CoverTab[91740]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:179
		_go_fuzz_dep_.CoverTab[91741]++
													b := bb2.Bytes()
													var dst []byte
													if cap(d.current.b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:182
			_go_fuzz_dep_.CoverTab[91747]++
														dst = d.current.b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:183
			// _ = "end of CoverTab[91747]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:184
			_go_fuzz_dep_.CoverTab[91748]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:184
			// _ = "end of CoverTab[91748]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:184
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:184
		// _ = "end of CoverTab[91741]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:184
		_go_fuzz_dep_.CoverTab[91742]++

													dst, err := d.DecodeAll(b, dst[:0])
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:187
			_go_fuzz_dep_.CoverTab[91749]++
														err = io.EOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:188
			// _ = "end of CoverTab[91749]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:189
			_go_fuzz_dep_.CoverTab[91750]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:189
			// _ = "end of CoverTab[91750]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:189
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:189
		// _ = "end of CoverTab[91742]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:189
		_go_fuzz_dep_.CoverTab[91743]++
													d.current.b = dst
													d.current.err = err
													d.current.flushed = true
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:193
			_go_fuzz_dep_.CoverTab[91751]++
														println("sync decode to", len(dst), "bytes, err:", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:194
			// _ = "end of CoverTab[91751]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:195
			_go_fuzz_dep_.CoverTab[91752]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:195
			// _ = "end of CoverTab[91752]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:195
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:195
		// _ = "end of CoverTab[91743]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:195
		_go_fuzz_dep_.CoverTab[91744]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:196
		// _ = "end of CoverTab[91744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:197
		_go_fuzz_dep_.CoverTab[91753]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:197
		// _ = "end of CoverTab[91753]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:197
	// _ = "end of CoverTab[91729]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:197
	_go_fuzz_dep_.CoverTab[91730]++

												if d.stream == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:199
		_go_fuzz_dep_.CoverTab[91754]++
													d.stream = make(chan decodeStream, 1)
													d.streamWg.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:201
		_curRoutineNum104_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:201
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum104_)
													go d.startStreamDecoder(d.stream)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:202
		// _ = "end of CoverTab[91754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:203
		_go_fuzz_dep_.CoverTab[91755]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:203
		// _ = "end of CoverTab[91755]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:203
	// _ = "end of CoverTab[91730]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:203
	_go_fuzz_dep_.CoverTab[91731]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:206
	d.current.decodeOutput = decodeOutput{}
	d.current.err = nil
	d.current.cancel = make(chan struct{})
	d.current.flushed = false
	d.current.d = nil

	d.stream <- decodeStream{
		r:	r,
		output:	d.current.output,
		cancel:	d.current.cancel,
	}
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:217
	// _ = "end of CoverTab[91731]"
}

// drainOutput will drain the output until errEndOfStream is sent.
func (d *Decoder) drainOutput() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:221
	_go_fuzz_dep_.CoverTab[91756]++
												if d.current.cancel != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:222
		_go_fuzz_dep_.CoverTab[91760]++
													println("cancelling current")
													close(d.current.cancel)
													d.current.cancel = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:225
		// _ = "end of CoverTab[91760]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:226
		_go_fuzz_dep_.CoverTab[91761]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:226
		// _ = "end of CoverTab[91761]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:226
	// _ = "end of CoverTab[91756]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:226
	_go_fuzz_dep_.CoverTab[91757]++
												if d.current.d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:227
		_go_fuzz_dep_.CoverTab[91762]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:228
			_go_fuzz_dep_.CoverTab[91764]++
														printf("re-adding current decoder %p, decoders: %d", d.current.d, len(d.decoders))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:229
			// _ = "end of CoverTab[91764]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:230
			_go_fuzz_dep_.CoverTab[91765]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:230
			// _ = "end of CoverTab[91765]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:230
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:230
		// _ = "end of CoverTab[91762]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:230
		_go_fuzz_dep_.CoverTab[91763]++
													d.decoders <- d.current.d
													d.current.d = nil
													d.current.b = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:233
		// _ = "end of CoverTab[91763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:234
		_go_fuzz_dep_.CoverTab[91766]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:234
		// _ = "end of CoverTab[91766]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:234
	// _ = "end of CoverTab[91757]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:234
	_go_fuzz_dep_.CoverTab[91758]++
												if d.current.output == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:235
		_go_fuzz_dep_.CoverTab[91767]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:235
		return d.current.flushed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:235
		// _ = "end of CoverTab[91767]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:235
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:235
		_go_fuzz_dep_.CoverTab[91768]++
													println("current already flushed")
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:237
		// _ = "end of CoverTab[91768]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:238
		_go_fuzz_dep_.CoverTab[91769]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:238
		// _ = "end of CoverTab[91769]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:238
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:238
	// _ = "end of CoverTab[91758]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:238
	_go_fuzz_dep_.CoverTab[91759]++
												for v := range d.current.output {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:239
		_go_fuzz_dep_.CoverTab[91770]++
													if v.d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:240
			_go_fuzz_dep_.CoverTab[91772]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:241
				_go_fuzz_dep_.CoverTab[91774]++
															printf("re-adding decoder %p", v.d)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:242
				// _ = "end of CoverTab[91774]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:243
				_go_fuzz_dep_.CoverTab[91775]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:243
				// _ = "end of CoverTab[91775]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:243
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:243
			// _ = "end of CoverTab[91772]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:243
			_go_fuzz_dep_.CoverTab[91773]++
														d.decoders <- v.d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:244
			// _ = "end of CoverTab[91773]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:245
			_go_fuzz_dep_.CoverTab[91776]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:245
			// _ = "end of CoverTab[91776]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:245
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:245
		// _ = "end of CoverTab[91770]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:245
		_go_fuzz_dep_.CoverTab[91771]++
													if v.err == errEndOfStream {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:246
			_go_fuzz_dep_.CoverTab[91777]++
														println("current flushed")
														d.current.flushed = true
														return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:249
			// _ = "end of CoverTab[91777]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:250
			_go_fuzz_dep_.CoverTab[91778]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:250
			// _ = "end of CoverTab[91778]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:250
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:250
		// _ = "end of CoverTab[91771]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:251
	// _ = "end of CoverTab[91759]"
}

// WriteTo writes data to w until there's no more data to write or when an error occurs.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:254
// The return value n is the number of bytes written.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:254
// Any error encountered during the write is also returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:257
func (d *Decoder) WriteTo(w io.Writer) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:257
	_go_fuzz_dep_.CoverTab[91779]++
												var n int64
												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:259
		_go_fuzz_dep_.CoverTab[91783]++
													if len(d.current.b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:260
			_go_fuzz_dep_.CoverTab[91786]++
														n2, err2 := w.Write(d.current.b)
														n += int64(n2)
														if err2 != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
				_go_fuzz_dep_.CoverTab[91787]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
				return (d.current.err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
					_go_fuzz_dep_.CoverTab[91788]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
					return d.current.err == io.EOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
					// _ = "end of CoverTab[91788]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
				}())
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
				// _ = "end of CoverTab[91787]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:263
				_go_fuzz_dep_.CoverTab[91789]++
															d.current.err = err2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:264
				// _ = "end of CoverTab[91789]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:265
				_go_fuzz_dep_.CoverTab[91790]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:265
				if n2 != len(d.current.b) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:265
					_go_fuzz_dep_.CoverTab[91791]++
																d.current.err = io.ErrShortWrite
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:266
					// _ = "end of CoverTab[91791]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:267
					_go_fuzz_dep_.CoverTab[91792]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:267
					// _ = "end of CoverTab[91792]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:267
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:267
				// _ = "end of CoverTab[91790]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:267
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:267
			// _ = "end of CoverTab[91786]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:268
			_go_fuzz_dep_.CoverTab[91793]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:268
			// _ = "end of CoverTab[91793]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:268
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:268
		// _ = "end of CoverTab[91783]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:268
		_go_fuzz_dep_.CoverTab[91784]++
													if d.current.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:269
			_go_fuzz_dep_.CoverTab[91794]++
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:270
			// _ = "end of CoverTab[91794]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:271
			_go_fuzz_dep_.CoverTab[91795]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:271
			// _ = "end of CoverTab[91795]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:271
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:271
		// _ = "end of CoverTab[91784]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:271
		_go_fuzz_dep_.CoverTab[91785]++
													d.nextBlock(true)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:272
		// _ = "end of CoverTab[91785]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:273
	// _ = "end of CoverTab[91779]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:273
	_go_fuzz_dep_.CoverTab[91780]++
												err := d.current.err
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:275
		_go_fuzz_dep_.CoverTab[91796]++
													d.drainOutput()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:276
		// _ = "end of CoverTab[91796]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:277
		_go_fuzz_dep_.CoverTab[91797]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:277
		// _ = "end of CoverTab[91797]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:277
	// _ = "end of CoverTab[91780]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:277
	_go_fuzz_dep_.CoverTab[91781]++
												if err == io.EOF {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:278
		_go_fuzz_dep_.CoverTab[91798]++
													err = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:279
		// _ = "end of CoverTab[91798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:280
		_go_fuzz_dep_.CoverTab[91799]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:280
		// _ = "end of CoverTab[91799]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:280
	// _ = "end of CoverTab[91781]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:280
	_go_fuzz_dep_.CoverTab[91782]++
												return n, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:281
	// _ = "end of CoverTab[91782]"
}

// DecodeAll allows stateless decoding of a blob of bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:284
// Output will be appended to dst, so if the destination size is known
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:284
// you can pre-allocate the destination slice to avoid allocations.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:284
// DecodeAll can be used concurrently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:284
// The Decoder concurrency limits will be respected.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:289
func (d *Decoder) DecodeAll(input, dst []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:289
	_go_fuzz_dep_.CoverTab[91800]++
												if d.current.err == ErrDecoderClosed {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:290
		_go_fuzz_dep_.CoverTab[91804]++
													return dst, ErrDecoderClosed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:291
		// _ = "end of CoverTab[91804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:292
		_go_fuzz_dep_.CoverTab[91805]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:292
		// _ = "end of CoverTab[91805]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:292
	// _ = "end of CoverTab[91800]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:292
	_go_fuzz_dep_.CoverTab[91801]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:295
	block := <-d.decoders
	frame := block.localFrame
	defer func() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:297
		_go_fuzz_dep_.CoverTab[91806]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:298
			_go_fuzz_dep_.CoverTab[91808]++
														printf("re-adding decoder: %p", block)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:299
			// _ = "end of CoverTab[91808]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:300
			_go_fuzz_dep_.CoverTab[91809]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:300
			// _ = "end of CoverTab[91809]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:300
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:300
		// _ = "end of CoverTab[91806]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:300
		_go_fuzz_dep_.CoverTab[91807]++
													frame.rawInput = nil
													frame.bBuf = nil
													d.decoders <- block
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:303
		// _ = "end of CoverTab[91807]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:304
	// _ = "end of CoverTab[91801]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:304
	_go_fuzz_dep_.CoverTab[91802]++
												frame.bBuf = input

												for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:307
		_go_fuzz_dep_.CoverTab[91810]++
													frame.history.reset()
													err := frame.reset(&frame.bBuf)
													if err == io.EOF {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:310
			_go_fuzz_dep_.CoverTab[91818]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:311
				_go_fuzz_dep_.CoverTab[91820]++
															println("frame reset return EOF")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:312
				// _ = "end of CoverTab[91820]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:313
				_go_fuzz_dep_.CoverTab[91821]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:313
				// _ = "end of CoverTab[91821]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:313
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:313
			// _ = "end of CoverTab[91818]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:313
			_go_fuzz_dep_.CoverTab[91819]++
														return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:314
			// _ = "end of CoverTab[91819]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:315
			_go_fuzz_dep_.CoverTab[91822]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:315
			// _ = "end of CoverTab[91822]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:315
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:315
		// _ = "end of CoverTab[91810]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:315
		_go_fuzz_dep_.CoverTab[91811]++
													if frame.DictionaryID != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:316
			_go_fuzz_dep_.CoverTab[91823]++
														dict, ok := d.dicts[*frame.DictionaryID]
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:318
				_go_fuzz_dep_.CoverTab[91825]++
															return nil, ErrUnknownDictionary
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:319
				// _ = "end of CoverTab[91825]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:320
				_go_fuzz_dep_.CoverTab[91826]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:320
				// _ = "end of CoverTab[91826]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:320
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:320
			// _ = "end of CoverTab[91823]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:320
			_go_fuzz_dep_.CoverTab[91824]++
														frame.history.setDict(&dict)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:321
			// _ = "end of CoverTab[91824]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:322
			_go_fuzz_dep_.CoverTab[91827]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:322
			// _ = "end of CoverTab[91827]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:322
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:322
		// _ = "end of CoverTab[91811]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:322
		_go_fuzz_dep_.CoverTab[91812]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:323
			_go_fuzz_dep_.CoverTab[91828]++
														return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:324
			// _ = "end of CoverTab[91828]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:325
			_go_fuzz_dep_.CoverTab[91829]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:325
			// _ = "end of CoverTab[91829]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:325
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:325
		// _ = "end of CoverTab[91812]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:325
		_go_fuzz_dep_.CoverTab[91813]++
													if frame.FrameContentSize > d.o.maxDecodedSize-uint64(len(dst)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:326
			_go_fuzz_dep_.CoverTab[91830]++
														return dst, ErrDecoderSizeExceeded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:327
			// _ = "end of CoverTab[91830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:328
			_go_fuzz_dep_.CoverTab[91831]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:328
			// _ = "end of CoverTab[91831]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:328
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:328
		// _ = "end of CoverTab[91813]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:328
		_go_fuzz_dep_.CoverTab[91814]++
													if frame.FrameContentSize > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:329
			_go_fuzz_dep_.CoverTab[91832]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:329
			return frame.FrameContentSize < 1<<30
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:329
			// _ = "end of CoverTab[91832]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:329
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:329
			_go_fuzz_dep_.CoverTab[91833]++

														if cap(dst)-len(dst) < int(frame.FrameContentSize) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:331
				_go_fuzz_dep_.CoverTab[91834]++
															dst2 := make([]byte, len(dst), len(dst)+int(frame.FrameContentSize))
															copy(dst2, dst)
															dst = dst2
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:334
				// _ = "end of CoverTab[91834]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:335
				_go_fuzz_dep_.CoverTab[91835]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:335
				// _ = "end of CoverTab[91835]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:335
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:335
			// _ = "end of CoverTab[91833]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:336
			_go_fuzz_dep_.CoverTab[91836]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:336
			// _ = "end of CoverTab[91836]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:336
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:336
		// _ = "end of CoverTab[91814]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:336
		_go_fuzz_dep_.CoverTab[91815]++
													if cap(dst) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:337
			_go_fuzz_dep_.CoverTab[91837]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:340
			size := len(input) * 2

			if size > 1<<20 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:342
				_go_fuzz_dep_.CoverTab[91840]++
															size = 1 << 20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:343
				// _ = "end of CoverTab[91840]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:344
				_go_fuzz_dep_.CoverTab[91841]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:344
				// _ = "end of CoverTab[91841]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:344
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:344
			// _ = "end of CoverTab[91837]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:344
			_go_fuzz_dep_.CoverTab[91838]++
														if uint64(size) > d.o.maxDecodedSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:345
				_go_fuzz_dep_.CoverTab[91842]++
															size = int(d.o.maxDecodedSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:346
				// _ = "end of CoverTab[91842]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:347
				_go_fuzz_dep_.CoverTab[91843]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:347
				// _ = "end of CoverTab[91843]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:347
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:347
			// _ = "end of CoverTab[91838]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:347
			_go_fuzz_dep_.CoverTab[91839]++
														dst = make([]byte, 0, size)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:348
			// _ = "end of CoverTab[91839]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:349
			_go_fuzz_dep_.CoverTab[91844]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:349
			// _ = "end of CoverTab[91844]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:349
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:349
		// _ = "end of CoverTab[91815]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:349
		_go_fuzz_dep_.CoverTab[91816]++

													dst, err = frame.runDecoder(dst, block)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:352
			_go_fuzz_dep_.CoverTab[91845]++
														return dst, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:353
			// _ = "end of CoverTab[91845]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:354
			_go_fuzz_dep_.CoverTab[91846]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:354
			// _ = "end of CoverTab[91846]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:354
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:354
		// _ = "end of CoverTab[91816]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:354
		_go_fuzz_dep_.CoverTab[91817]++
													if len(frame.bBuf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:355
			_go_fuzz_dep_.CoverTab[91847]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:356
				_go_fuzz_dep_.CoverTab[91849]++
															println("frame dbuf empty")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:357
				// _ = "end of CoverTab[91849]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:358
				_go_fuzz_dep_.CoverTab[91850]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:358
				// _ = "end of CoverTab[91850]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:358
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:358
			// _ = "end of CoverTab[91847]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:358
			_go_fuzz_dep_.CoverTab[91848]++
														break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:359
			// _ = "end of CoverTab[91848]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:360
			_go_fuzz_dep_.CoverTab[91851]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:360
			// _ = "end of CoverTab[91851]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:360
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:360
		// _ = "end of CoverTab[91817]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:361
	// _ = "end of CoverTab[91802]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:361
	_go_fuzz_dep_.CoverTab[91803]++
												return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:362
	// _ = "end of CoverTab[91803]"
}

// nextBlock returns the next block.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:365
// If an error occurs d.err will be set.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:365
// Optionally the function can block for new output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:365
// If non-blocking mode is used the returned boolean will be false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:365
// if no data was available without blocking.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:370
func (d *Decoder) nextBlock(blocking bool) (ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:370
	_go_fuzz_dep_.CoverTab[91852]++
												if d.current.d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:371
		_go_fuzz_dep_.CoverTab[91857]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:372
			_go_fuzz_dep_.CoverTab[91859]++
														printf("re-adding current decoder %p", d.current.d)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:373
			// _ = "end of CoverTab[91859]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:374
			_go_fuzz_dep_.CoverTab[91860]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:374
			// _ = "end of CoverTab[91860]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:374
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:374
		// _ = "end of CoverTab[91857]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:374
		_go_fuzz_dep_.CoverTab[91858]++
													d.decoders <- d.current.d
													d.current.d = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:376
		// _ = "end of CoverTab[91858]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:377
		_go_fuzz_dep_.CoverTab[91861]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:377
		// _ = "end of CoverTab[91861]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:377
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:377
	// _ = "end of CoverTab[91852]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:377
	_go_fuzz_dep_.CoverTab[91853]++
												if d.current.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:378
		_go_fuzz_dep_.CoverTab[91862]++

													return blocking
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:380
		// _ = "end of CoverTab[91862]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:381
		_go_fuzz_dep_.CoverTab[91863]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:381
		// _ = "end of CoverTab[91863]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:381
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:381
	// _ = "end of CoverTab[91853]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:381
	_go_fuzz_dep_.CoverTab[91854]++

												if blocking {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:383
		_go_fuzz_dep_.CoverTab[91864]++
													d.current.decodeOutput = <-d.current.output
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:384
		// _ = "end of CoverTab[91864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:385
		_go_fuzz_dep_.CoverTab[91865]++
													select {
		case d.current.decodeOutput = <-d.current.output:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:387
			_go_fuzz_dep_.CoverTab[91866]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:387
			// _ = "end of CoverTab[91866]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:388
			_go_fuzz_dep_.CoverTab[91867]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:389
			// _ = "end of CoverTab[91867]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:390
		// _ = "end of CoverTab[91865]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:391
	// _ = "end of CoverTab[91854]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:391
	_go_fuzz_dep_.CoverTab[91855]++
												if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:392
		_go_fuzz_dep_.CoverTab[91868]++
													println("got", len(d.current.b), "bytes, error:", d.current.err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:393
		// _ = "end of CoverTab[91868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:394
		_go_fuzz_dep_.CoverTab[91869]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:394
		// _ = "end of CoverTab[91869]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:394
	// _ = "end of CoverTab[91855]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:394
	_go_fuzz_dep_.CoverTab[91856]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:395
	// _ = "end of CoverTab[91856]"
}

// Close will release all resources.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:398
// It is NOT possible to reuse the decoder after this.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:400
func (d *Decoder) Close() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:400
	_go_fuzz_dep_.CoverTab[91870]++
												if d.current.err == ErrDecoderClosed {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:401
		_go_fuzz_dep_.CoverTab[91875]++
													return
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:402
		// _ = "end of CoverTab[91875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:403
		_go_fuzz_dep_.CoverTab[91876]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:403
		// _ = "end of CoverTab[91876]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:403
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:403
	// _ = "end of CoverTab[91870]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:403
	_go_fuzz_dep_.CoverTab[91871]++
												d.drainOutput()
												if d.stream != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:405
		_go_fuzz_dep_.CoverTab[91877]++
													close(d.stream)
													d.streamWg.Wait()
													d.stream = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:408
		// _ = "end of CoverTab[91877]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:409
		_go_fuzz_dep_.CoverTab[91878]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:409
		// _ = "end of CoverTab[91878]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:409
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:409
	// _ = "end of CoverTab[91871]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:409
	_go_fuzz_dep_.CoverTab[91872]++
												if d.decoders != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:410
		_go_fuzz_dep_.CoverTab[91879]++
													close(d.decoders)
													for dec := range d.decoders {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:412
			_go_fuzz_dep_.CoverTab[91881]++
														dec.Close()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:413
			// _ = "end of CoverTab[91881]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:414
		// _ = "end of CoverTab[91879]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:414
		_go_fuzz_dep_.CoverTab[91880]++
													d.decoders = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:415
		// _ = "end of CoverTab[91880]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:416
		_go_fuzz_dep_.CoverTab[91882]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:416
		// _ = "end of CoverTab[91882]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:416
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:416
	// _ = "end of CoverTab[91872]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:416
	_go_fuzz_dep_.CoverTab[91873]++
												if d.current.d != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:417
		_go_fuzz_dep_.CoverTab[91883]++
													d.current.d.Close()
													d.current.d = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:419
		// _ = "end of CoverTab[91883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:420
		_go_fuzz_dep_.CoverTab[91884]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:420
		// _ = "end of CoverTab[91884]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:420
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:420
	// _ = "end of CoverTab[91873]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:420
	_go_fuzz_dep_.CoverTab[91874]++
												d.current.err = ErrDecoderClosed
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:421
	// _ = "end of CoverTab[91874]"
}

// IOReadCloser returns the decoder as an io.ReadCloser for convenience.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:424
// Any changes to the decoder will be reflected, so the returned ReadCloser
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:424
// can be reused along with the decoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:424
// io.WriterTo is also supported by the returned ReadCloser.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:428
func (d *Decoder) IOReadCloser() io.ReadCloser {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:428
	_go_fuzz_dep_.CoverTab[91885]++
												return closeWrapper{d: d}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:429
	// _ = "end of CoverTab[91885]"
}

// closeWrapper wraps a function call as a closer.
type closeWrapper struct {
	d *Decoder
}

// WriteTo forwards WriteTo calls to the decoder.
func (c closeWrapper) WriteTo(w io.Writer) (n int64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:438
	_go_fuzz_dep_.CoverTab[91886]++
												return c.d.WriteTo(w)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:439
	// _ = "end of CoverTab[91886]"
}

// Read forwards read calls to the decoder.
func (c closeWrapper) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:443
	_go_fuzz_dep_.CoverTab[91887]++
												return c.d.Read(p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:444
	// _ = "end of CoverTab[91887]"
}

// Close closes the decoder.
func (c closeWrapper) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:448
	_go_fuzz_dep_.CoverTab[91888]++
												c.d.Close()
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:450
	// _ = "end of CoverTab[91888]"
}

type decodeOutput struct {
	d	*blockDec
	b	[]byte
	err	error
}

type decodeStream struct {
	r	io.Reader

	// Blocks ready to be written to output.
	output	chan decodeOutput

	// cancel reading from the input
	cancel	chan struct{}
}

// errEndOfStream indicates that everything from the stream was read.
var errEndOfStream = errors.New("end-of-stream")

// Create Decoder:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
// Spawn n block decoders. These accept tasks to decode a block.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
// Create goroutine that handles stream processing, this will send history to decoders as they are available.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
// Decoders update the history as they decode.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
// When a block is returned:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
//	a) history is sent to the next decoder,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
//	b) content written to CRC.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
//	c) return data to WRITER.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
//	d) wait for next block to return data.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:472
// Once WRITTEN, the decoders reused by the writer frame decoder for re-use.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:482
func (d *Decoder) startStreamDecoder(inStream chan decodeStream) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:482
	_go_fuzz_dep_.CoverTab[91889]++
												defer d.streamWg.Done()
												frame := newFrameDec(d.o)
												for stream := range inStream {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:485
		_go_fuzz_dep_.CoverTab[91890]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:486
			_go_fuzz_dep_.CoverTab[91893]++
														println("got new stream")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:487
			// _ = "end of CoverTab[91893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:488
			_go_fuzz_dep_.CoverTab[91894]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:488
			// _ = "end of CoverTab[91894]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:488
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:488
		// _ = "end of CoverTab[91890]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:488
		_go_fuzz_dep_.CoverTab[91891]++
													br := readerWrapper{r: stream.r}
	decodeStream:
		for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:491
			_go_fuzz_dep_.CoverTab[91895]++
														frame.history.reset()
														err := frame.reset(&br)
														if debugDecoder && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:494
				_go_fuzz_dep_.CoverTab[91901]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:494
				return err != nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:494
				// _ = "end of CoverTab[91901]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:494
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:494
				_go_fuzz_dep_.CoverTab[91902]++
															println("Frame decoder returned", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:495
				// _ = "end of CoverTab[91902]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:496
				_go_fuzz_dep_.CoverTab[91903]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:496
				// _ = "end of CoverTab[91903]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:496
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:496
			// _ = "end of CoverTab[91895]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:496
			_go_fuzz_dep_.CoverTab[91896]++
														if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:497
				_go_fuzz_dep_.CoverTab[91904]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:497
				return frame.DictionaryID != nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:497
				// _ = "end of CoverTab[91904]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:497
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:497
				_go_fuzz_dep_.CoverTab[91905]++
															dict, ok := d.dicts[*frame.DictionaryID]
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:499
					_go_fuzz_dep_.CoverTab[91906]++
																err = ErrUnknownDictionary
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:500
					// _ = "end of CoverTab[91906]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:501
					_go_fuzz_dep_.CoverTab[91907]++
																frame.history.setDict(&dict)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:502
					// _ = "end of CoverTab[91907]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:503
				// _ = "end of CoverTab[91905]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:504
				_go_fuzz_dep_.CoverTab[91908]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:504
				// _ = "end of CoverTab[91908]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:504
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:504
			// _ = "end of CoverTab[91896]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:504
			_go_fuzz_dep_.CoverTab[91897]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:505
				_go_fuzz_dep_.CoverTab[91909]++
															stream.output <- decodeOutput{
					err: err,
				}
															break
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:509
				// _ = "end of CoverTab[91909]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:510
				_go_fuzz_dep_.CoverTab[91910]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:510
				// _ = "end of CoverTab[91910]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:510
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:510
			// _ = "end of CoverTab[91897]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:510
			_go_fuzz_dep_.CoverTab[91898]++
														if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:511
				_go_fuzz_dep_.CoverTab[91911]++
															println("starting frame decoder")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:512
				// _ = "end of CoverTab[91911]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:513
				_go_fuzz_dep_.CoverTab[91912]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:513
				// _ = "end of CoverTab[91912]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:513
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:513
			// _ = "end of CoverTab[91898]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:513
			_go_fuzz_dep_.CoverTab[91899]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:516
			frame.frameDone.Add(1)
														frame.initAsync()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:517
			_curRoutineNum105_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:517
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum105_)

														go frame.startDecoder(stream.output)
		decodeFrame:

			for {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:522
				_go_fuzz_dep_.CoverTab[91913]++
															dec := <-d.decoders
															select {
				case <-stream.cancel:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:525
					_go_fuzz_dep_.CoverTab[91915]++
																if !frame.sendErr(dec, io.EOF) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:526
						_go_fuzz_dep_.CoverTab[91918]++

																	stream.output <- decodeOutput{d: dec}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:528
						// _ = "end of CoverTab[91918]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:529
						_go_fuzz_dep_.CoverTab[91919]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:529
						// _ = "end of CoverTab[91919]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:529
					}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:529
					// _ = "end of CoverTab[91915]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:529
					_go_fuzz_dep_.CoverTab[91916]++
																break decodeStream
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:530
					// _ = "end of CoverTab[91916]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:531
					_go_fuzz_dep_.CoverTab[91917]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:531
					// _ = "end of CoverTab[91917]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:532
				// _ = "end of CoverTab[91913]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:532
				_go_fuzz_dep_.CoverTab[91914]++
															err := frame.next(dec)
															switch err {
				case io.EOF:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:535
					_go_fuzz_dep_.CoverTab[91920]++

																println("EOF on next block")
																break decodeFrame
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:538
					// _ = "end of CoverTab[91920]"
				case nil:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:539
					_go_fuzz_dep_.CoverTab[91921]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:540
					// _ = "end of CoverTab[91921]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:541
					_go_fuzz_dep_.CoverTab[91922]++
																println("block decoder returned", err)
																break decodeStream
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:543
					// _ = "end of CoverTab[91922]"
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:544
				// _ = "end of CoverTab[91914]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:545
			// _ = "end of CoverTab[91899]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:545
			_go_fuzz_dep_.CoverTab[91900]++

														println("waiting for done")
														frame.frameDone.Wait()
														println("done waiting...")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:549
			// _ = "end of CoverTab[91900]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:550
		// _ = "end of CoverTab[91891]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:550
		_go_fuzz_dep_.CoverTab[91892]++
													frame.frameDone.Wait()
													println("Sending EOS")
													stream.output <- decodeOutput{err: errEndOfStream}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:553
		// _ = "end of CoverTab[91892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:554
	// _ = "end of CoverTab[91889]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:555
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder.go:555
var _ = _go_fuzz_dep_.CoverTab
