// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:5
)

import (
	"errors"
	"runtime"
)

// DOption is an option for creating a decoder.
type DOption func(*decoderOptions) error

// options retains accumulated state of multiple options.
type decoderOptions struct {
	lowMem		bool
	concurrent	int
	maxDecodedSize	uint64
	maxWindowSize	uint64
	dicts		[]dict
}

func (o *decoderOptions) setDefault() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:24
	_go_fuzz_dep_.CoverTab[91923]++
													*o = decoderOptions{

		lowMem:		true,
		concurrent:	runtime.GOMAXPROCS(0),
		maxWindowSize:	MaxWindowSize,
	}
													o.maxDecodedSize = 1 << 63
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:31
	// _ = "end of CoverTab[91923]"
}

// WithDecoderLowmem will set whether to use a lower amount of memory,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:34
// but possibly have to allocate more while running.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:36
func WithDecoderLowmem(b bool) DOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:36
	_go_fuzz_dep_.CoverTab[91924]++
													return func(o *decoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:37
		_go_fuzz_dep_.CoverTab[91925]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:37
		o.lowMem = b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:37
		return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:37
		// _ = "end of CoverTab[91925]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:37
	// _ = "end of CoverTab[91924]"
}

// WithDecoderConcurrency will set the concurrency,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:40
// meaning the maximum number of decoders to run concurrently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:40
// The value supplied must be at least 1.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:40
// By default this will be set to GOMAXPROCS.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:44
func WithDecoderConcurrency(n int) DOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:44
	_go_fuzz_dep_.CoverTab[91926]++
													return func(o *decoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:45
		_go_fuzz_dep_.CoverTab[91927]++
														if n <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:46
			_go_fuzz_dep_.CoverTab[91929]++
															return errors.New("concurrency must be at least 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:47
			// _ = "end of CoverTab[91929]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:48
			_go_fuzz_dep_.CoverTab[91930]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:48
			// _ = "end of CoverTab[91930]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:48
		// _ = "end of CoverTab[91927]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:48
		_go_fuzz_dep_.CoverTab[91928]++
														o.concurrent = n
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:50
		// _ = "end of CoverTab[91928]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:51
	// _ = "end of CoverTab[91926]"
}

// WithDecoderMaxMemory allows to set a maximum decoded size for in-memory
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:54
// non-streaming operations or maximum window size for streaming operations.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:54
// This can be used to control memory usage of potentially hostile content.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:54
// Maximum and default is 1 << 63 bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:58
func WithDecoderMaxMemory(n uint64) DOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:58
	_go_fuzz_dep_.CoverTab[91931]++
													return func(o *decoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:59
		_go_fuzz_dep_.CoverTab[91932]++
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:60
			_go_fuzz_dep_.CoverTab[91935]++
															return errors.New("WithDecoderMaxMemory must be at least 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:61
			// _ = "end of CoverTab[91935]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:62
			_go_fuzz_dep_.CoverTab[91936]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:62
			// _ = "end of CoverTab[91936]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:62
		// _ = "end of CoverTab[91932]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:62
		_go_fuzz_dep_.CoverTab[91933]++
														if n > 1<<63 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:63
			_go_fuzz_dep_.CoverTab[91937]++
															return errors.New("WithDecoderMaxmemory must be less than 1 << 63")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:64
			// _ = "end of CoverTab[91937]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:65
			_go_fuzz_dep_.CoverTab[91938]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:65
			// _ = "end of CoverTab[91938]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:65
		// _ = "end of CoverTab[91933]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:65
		_go_fuzz_dep_.CoverTab[91934]++
														o.maxDecodedSize = n
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:67
		// _ = "end of CoverTab[91934]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:68
	// _ = "end of CoverTab[91931]"
}

// WithDecoderDicts allows to register one or more dictionaries for the decoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:71
// If several dictionaries with the same ID is provided the last one will be used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:73
func WithDecoderDicts(dicts ...[]byte) DOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:73
	_go_fuzz_dep_.CoverTab[91939]++
													return func(o *decoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:74
		_go_fuzz_dep_.CoverTab[91940]++
														for _, b := range dicts {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:75
			_go_fuzz_dep_.CoverTab[91942]++
															d, err := loadDict(b)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:77
				_go_fuzz_dep_.CoverTab[91944]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:78
				// _ = "end of CoverTab[91944]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:79
				_go_fuzz_dep_.CoverTab[91945]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:79
				// _ = "end of CoverTab[91945]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:79
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:79
			// _ = "end of CoverTab[91942]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:79
			_go_fuzz_dep_.CoverTab[91943]++
															o.dicts = append(o.dicts, *d)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:80
			// _ = "end of CoverTab[91943]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:81
		// _ = "end of CoverTab[91940]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:81
		_go_fuzz_dep_.CoverTab[91941]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:82
		// _ = "end of CoverTab[91941]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:83
	// _ = "end of CoverTab[91939]"
}

// WithDecoderMaxWindow allows to set a maximum window size for decodes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:86
// This allows rejecting packets that will cause big memory usage.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:86
// The Decoder will likely allocate more memory based on the WithDecoderLowmem setting.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:86
// If WithDecoderMaxMemory is set to a lower value, that will be used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:86
// Default is 512MB, Maximum is ~3.75 TB as per zstandard spec.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:91
func WithDecoderMaxWindow(size uint64) DOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:91
	_go_fuzz_dep_.CoverTab[91946]++
													return func(o *decoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:92
		_go_fuzz_dep_.CoverTab[91947]++
														if size < MinWindowSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:93
			_go_fuzz_dep_.CoverTab[91950]++
															return errors.New("WithMaxWindowSize must be at least 1KB, 1024 bytes")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:94
			// _ = "end of CoverTab[91950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:95
			_go_fuzz_dep_.CoverTab[91951]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:95
			// _ = "end of CoverTab[91951]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:95
		// _ = "end of CoverTab[91947]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:95
		_go_fuzz_dep_.CoverTab[91948]++
														if size > (1<<41)+7*(1<<38) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:96
			_go_fuzz_dep_.CoverTab[91952]++
															return errors.New("WithMaxWindowSize must be less than (1<<41) + 7*(1<<38) ~ 3.75TB")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:97
			// _ = "end of CoverTab[91952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:98
			_go_fuzz_dep_.CoverTab[91953]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:98
			// _ = "end of CoverTab[91953]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:98
		// _ = "end of CoverTab[91948]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:98
		_go_fuzz_dep_.CoverTab[91949]++
														o.maxWindowSize = size
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:100
		// _ = "end of CoverTab[91949]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:101
	// _ = "end of CoverTab[91946]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:102
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/decoder_options.go:102
var _ = _go_fuzz_dep_.CoverTab
