//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:1
)

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// EOption is an option for creating a encoder.
type EOption func(*encoderOptions) error

// options retains accumulated state of multiple options.
type encoderOptions struct {
	concurrent	int
	level		EncoderLevel
	single		*bool
	pad		int
	blockSize	int
	windowSize	int
	crc		bool
	fullZero	bool
	noEntropy	bool
	allLitEntropy	bool
	customWindow	bool
	customALEntropy	bool
	customBlockSize	bool
	lowMem		bool
	dict		*dict
}

func (o *encoderOptions) setDefault() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:32
	_go_fuzz_dep_.CoverTab[94001]++
													*o = encoderOptions{
		concurrent:	runtime.GOMAXPROCS(0),
		crc:		true,
		single:		nil,
		blockSize:	maxCompressedBlockSize,
		windowSize:	8 << 20,
		level:		SpeedDefault,
		allLitEntropy:	true,
		lowMem:		false,
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:42
	// _ = "end of CoverTab[94001]"
}

// encoder returns an encoder with the selected options.
func (o encoderOptions) encoder() encoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:46
	_go_fuzz_dep_.CoverTab[94002]++
													switch o.level {
	case SpeedFastest:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:48
		_go_fuzz_dep_.CoverTab[94004]++
														if o.dict != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:49
			_go_fuzz_dep_.CoverTab[94012]++
															return &fastEncoderDict{fastEncoder: fastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:50
			// _ = "end of CoverTab[94012]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:51
			_go_fuzz_dep_.CoverTab[94013]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:51
			// _ = "end of CoverTab[94013]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:51
		// _ = "end of CoverTab[94004]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:51
		_go_fuzz_dep_.CoverTab[94005]++
														return &fastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:52
		// _ = "end of CoverTab[94005]"

	case SpeedDefault:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:54
		_go_fuzz_dep_.CoverTab[94006]++
														if o.dict != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:55
			_go_fuzz_dep_.CoverTab[94014]++
															return &doubleFastEncoderDict{fastEncoderDict: fastEncoderDict{fastEncoder: fastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:56
			// _ = "end of CoverTab[94014]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:57
			_go_fuzz_dep_.CoverTab[94015]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:57
			// _ = "end of CoverTab[94015]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:57
		// _ = "end of CoverTab[94006]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:57
		_go_fuzz_dep_.CoverTab[94007]++
														return &doubleFastEncoder{fastEncoder: fastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:58
		// _ = "end of CoverTab[94007]"
	case SpeedBetterCompression:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:59
		_go_fuzz_dep_.CoverTab[94008]++
														if o.dict != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:60
			_go_fuzz_dep_.CoverTab[94016]++
															return &betterFastEncoderDict{betterFastEncoder: betterFastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:61
			// _ = "end of CoverTab[94016]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:62
			_go_fuzz_dep_.CoverTab[94017]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:62
			// _ = "end of CoverTab[94017]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:62
		// _ = "end of CoverTab[94008]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:62
		_go_fuzz_dep_.CoverTab[94009]++
														return &betterFastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:63
		// _ = "end of CoverTab[94009]"
	case SpeedBestCompression:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:64
		_go_fuzz_dep_.CoverTab[94010]++
														return &bestFastEncoder{fastBase: fastBase{maxMatchOff: int32(o.windowSize), lowMem: o.lowMem}}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:65
		// _ = "end of CoverTab[94010]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:65
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:65
		_go_fuzz_dep_.CoverTab[94011]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:65
		// _ = "end of CoverTab[94011]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:66
	// _ = "end of CoverTab[94002]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:66
	_go_fuzz_dep_.CoverTab[94003]++
													panic("unknown compression level")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:67
	// _ = "end of CoverTab[94003]"
}

// WithEncoderCRC will add CRC value to output.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:70
// Output will be 4 bytes larger.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:72
func WithEncoderCRC(b bool) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:72
	_go_fuzz_dep_.CoverTab[94018]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:73
		_go_fuzz_dep_.CoverTab[94019]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:73
		o.crc = b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:73
		return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:73
		// _ = "end of CoverTab[94019]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:73
	// _ = "end of CoverTab[94018]"
}

// WithEncoderConcurrency will set the concurrency,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:76
// meaning the maximum number of encoders to run concurrently.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:76
// The value supplied must be at least 1.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:76
// By default this will be set to GOMAXPROCS.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:80
func WithEncoderConcurrency(n int) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:80
	_go_fuzz_dep_.CoverTab[94020]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:81
		_go_fuzz_dep_.CoverTab[94021]++
														if n <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:82
			_go_fuzz_dep_.CoverTab[94023]++
															return fmt.Errorf("concurrency must be at least 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:83
			// _ = "end of CoverTab[94023]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:84
			_go_fuzz_dep_.CoverTab[94024]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:84
			// _ = "end of CoverTab[94024]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:84
		// _ = "end of CoverTab[94021]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:84
		_go_fuzz_dep_.CoverTab[94022]++
														o.concurrent = n
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:86
		// _ = "end of CoverTab[94022]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:87
	// _ = "end of CoverTab[94020]"
}

// WithWindowSize will set the maximum allowed back-reference distance.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:90
// The value must be a power of two between MinWindowSize and MaxWindowSize.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:90
// A larger value will enable better compression but allocate more memory and,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:90
// for above-default values, take considerably longer.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:90
// The default value is determined by the compression level.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:95
func WithWindowSize(n int) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:95
	_go_fuzz_dep_.CoverTab[94025]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:96
		_go_fuzz_dep_.CoverTab[94026]++
														switch {
		case n < MinWindowSize:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:98
			_go_fuzz_dep_.CoverTab[94029]++
															return fmt.Errorf("window size must be at least %d", MinWindowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:99
			// _ = "end of CoverTab[94029]"
		case n > MaxWindowSize:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:100
			_go_fuzz_dep_.CoverTab[94030]++
															return fmt.Errorf("window size must be at most %d", MaxWindowSize)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:101
			// _ = "end of CoverTab[94030]"
		case (n & (n - 1)) != 0:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:102
			_go_fuzz_dep_.CoverTab[94031]++
															return errors.New("window size must be a power of 2")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:103
			// _ = "end of CoverTab[94031]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:103
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:103
			_go_fuzz_dep_.CoverTab[94032]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:103
			// _ = "end of CoverTab[94032]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:104
		// _ = "end of CoverTab[94026]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:104
		_go_fuzz_dep_.CoverTab[94027]++

														o.windowSize = n
														o.customWindow = true
														if o.blockSize > o.windowSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:108
			_go_fuzz_dep_.CoverTab[94033]++
															o.blockSize = o.windowSize
															o.customBlockSize = true
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:110
			// _ = "end of CoverTab[94033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:111
			_go_fuzz_dep_.CoverTab[94034]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:111
			// _ = "end of CoverTab[94034]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:111
		// _ = "end of CoverTab[94027]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:111
		_go_fuzz_dep_.CoverTab[94028]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:112
		// _ = "end of CoverTab[94028]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:113
	// _ = "end of CoverTab[94025]"
}

// WithEncoderPadding will add padding to all output so the size will be a multiple of n.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:116
// This can be used to obfuscate the exact output size or make blocks of a certain size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:116
// The contents will be a skippable frame, so it will be invisible by the decoder.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:116
// n must be > 0 and <= 1GB, 1<<30 bytes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:116
// The padded area will be filled with data from crypto/rand.Reader.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:116
// If `EncodeAll` is used with data already in the destination, the total size will be multiple of this.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:122
func WithEncoderPadding(n int) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:122
	_go_fuzz_dep_.CoverTab[94035]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:123
		_go_fuzz_dep_.CoverTab[94036]++
														if n <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:124
			_go_fuzz_dep_.CoverTab[94040]++
															return fmt.Errorf("padding must be at least 1")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:125
			// _ = "end of CoverTab[94040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:126
			_go_fuzz_dep_.CoverTab[94041]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:126
			// _ = "end of CoverTab[94041]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:126
		// _ = "end of CoverTab[94036]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:126
		_go_fuzz_dep_.CoverTab[94037]++

														if n == 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:128
			_go_fuzz_dep_.CoverTab[94042]++
															o.pad = 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:129
			// _ = "end of CoverTab[94042]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:130
			_go_fuzz_dep_.CoverTab[94043]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:130
			// _ = "end of CoverTab[94043]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:130
		// _ = "end of CoverTab[94037]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:130
		_go_fuzz_dep_.CoverTab[94038]++
														if n > 1<<30 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:131
			_go_fuzz_dep_.CoverTab[94044]++
															return fmt.Errorf("padding must less than 1GB (1<<30 bytes) ")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:132
			// _ = "end of CoverTab[94044]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:133
			_go_fuzz_dep_.CoverTab[94045]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:133
			// _ = "end of CoverTab[94045]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:133
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:133
		// _ = "end of CoverTab[94038]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:133
		_go_fuzz_dep_.CoverTab[94039]++
														o.pad = n
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:135
		// _ = "end of CoverTab[94039]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:136
	// _ = "end of CoverTab[94035]"
}

// EncoderLevel predefines encoder compression levels.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:139
// Only use the constants made available, since the actual mapping
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:139
// of these values are very likely to change and your compression could change
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:139
// unpredictably when upgrading the library.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:143
type EncoderLevel int

const (
	speedNotSet	EncoderLevel	= iota

	// SpeedFastest will choose the fastest reasonable compression.
	// This is roughly equivalent to the fastest Zstandard mode.
	SpeedFastest

	// SpeedDefault is the default "pretty fast" compression option.
	// This is roughly equivalent to the default Zstandard mode (level 3).
	SpeedDefault

	// SpeedBetterCompression will yield better compression than the default.
	// Currently it is about zstd level 7-8 with ~ 2x-3x the default CPU usage.
	// By using this, notice that CPU usage may go up in the future.
	SpeedBetterCompression

	// SpeedBestCompression will choose the best available compression option.
	// This will offer the best compression no matter the CPU cost.
	SpeedBestCompression

	// speedLast should be kept as the last actual compression option.
	// The is not for external usage, but is used to keep track of the valid options.
	speedLast
)

// EncoderLevelFromString will convert a string representation of an encoding level back
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:170
// to a compression level. The compare is not case sensitive.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:170
// If the string wasn't recognized, (false, SpeedDefault) will be returned.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:173
func EncoderLevelFromString(s string) (bool, EncoderLevel) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:173
	_go_fuzz_dep_.CoverTab[94046]++
													for l := speedNotSet + 1; l < speedLast; l++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:174
		_go_fuzz_dep_.CoverTab[94048]++
														if strings.EqualFold(s, l.String()) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:175
			_go_fuzz_dep_.CoverTab[94049]++
															return true, l
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:176
			// _ = "end of CoverTab[94049]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:177
			_go_fuzz_dep_.CoverTab[94050]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:177
			// _ = "end of CoverTab[94050]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:177
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:177
		// _ = "end of CoverTab[94048]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:178
	// _ = "end of CoverTab[94046]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:178
	_go_fuzz_dep_.CoverTab[94047]++
													return false, SpeedDefault
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:179
	// _ = "end of CoverTab[94047]"
}

// EncoderLevelFromZstd will return an encoder level that closest matches the compression
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:182
// ratio of a specific zstd compression level.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:182
// Many input values will provide the same compression level.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:185
func EncoderLevelFromZstd(level int) EncoderLevel {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:185
	_go_fuzz_dep_.CoverTab[94051]++
													switch {
	case level < 3:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:187
		_go_fuzz_dep_.CoverTab[94052]++
														return SpeedFastest
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:188
		// _ = "end of CoverTab[94052]"
	case level >= 3 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:189
		_go_fuzz_dep_.CoverTab[94056]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:189
		return level < 6
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:189
		// _ = "end of CoverTab[94056]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:189
	}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:189
		_go_fuzz_dep_.CoverTab[94053]++
														return SpeedDefault
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:190
		// _ = "end of CoverTab[94053]"
	case level >= 6 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:191
		_go_fuzz_dep_.CoverTab[94057]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:191
		return level < 10
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:191
		// _ = "end of CoverTab[94057]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:191
	}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:191
		_go_fuzz_dep_.CoverTab[94054]++
														return SpeedBetterCompression
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:192
		// _ = "end of CoverTab[94054]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:193
		_go_fuzz_dep_.CoverTab[94055]++
														return SpeedBestCompression
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:194
		// _ = "end of CoverTab[94055]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:195
	// _ = "end of CoverTab[94051]"
}

// String provides a string representation of the compression level.
func (e EncoderLevel) String() string {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:199
	_go_fuzz_dep_.CoverTab[94058]++
													switch e {
	case SpeedFastest:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:201
		_go_fuzz_dep_.CoverTab[94059]++
														return "fastest"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:202
		// _ = "end of CoverTab[94059]"
	case SpeedDefault:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:203
		_go_fuzz_dep_.CoverTab[94060]++
														return "default"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:204
		// _ = "end of CoverTab[94060]"
	case SpeedBetterCompression:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:205
		_go_fuzz_dep_.CoverTab[94061]++
														return "better"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:206
		// _ = "end of CoverTab[94061]"
	case SpeedBestCompression:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:207
		_go_fuzz_dep_.CoverTab[94062]++
														return "best"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:208
		// _ = "end of CoverTab[94062]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:209
		_go_fuzz_dep_.CoverTab[94063]++
														return "invalid"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:210
		// _ = "end of CoverTab[94063]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:211
	// _ = "end of CoverTab[94058]"
}

// WithEncoderLevel specifies a predefined compression level.
func WithEncoderLevel(l EncoderLevel) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:215
	_go_fuzz_dep_.CoverTab[94064]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:216
		_go_fuzz_dep_.CoverTab[94065]++
														switch {
		case l <= speedNotSet || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:218
			_go_fuzz_dep_.CoverTab[94071]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:218
			return l >= speedLast
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:218
			// _ = "end of CoverTab[94071]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:218
		}():
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:218
			_go_fuzz_dep_.CoverTab[94069]++
															return fmt.Errorf("unknown encoder level")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:219
			// _ = "end of CoverTab[94069]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:219
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:219
			_go_fuzz_dep_.CoverTab[94070]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:219
			// _ = "end of CoverTab[94070]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:220
		// _ = "end of CoverTab[94065]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:220
		_go_fuzz_dep_.CoverTab[94066]++
														o.level = l
														if !o.customWindow {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:222
			_go_fuzz_dep_.CoverTab[94072]++
															switch o.level {
			case SpeedFastest:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:224
				_go_fuzz_dep_.CoverTab[94073]++
																o.windowSize = 4 << 20
																if !o.customBlockSize {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:226
					_go_fuzz_dep_.CoverTab[94078]++
																	o.blockSize = 1 << 16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:227
					// _ = "end of CoverTab[94078]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:228
					_go_fuzz_dep_.CoverTab[94079]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:228
					// _ = "end of CoverTab[94079]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:228
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:228
				// _ = "end of CoverTab[94073]"
			case SpeedDefault:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:229
				_go_fuzz_dep_.CoverTab[94074]++
																o.windowSize = 8 << 20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:230
				// _ = "end of CoverTab[94074]"
			case SpeedBetterCompression:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:231
				_go_fuzz_dep_.CoverTab[94075]++
																o.windowSize = 16 << 20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:232
				// _ = "end of CoverTab[94075]"
			case SpeedBestCompression:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:233
				_go_fuzz_dep_.CoverTab[94076]++
																o.windowSize = 32 << 20
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:234
				// _ = "end of CoverTab[94076]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:234
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:234
				_go_fuzz_dep_.CoverTab[94077]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:234
				// _ = "end of CoverTab[94077]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:235
			// _ = "end of CoverTab[94072]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:236
			_go_fuzz_dep_.CoverTab[94080]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:236
			// _ = "end of CoverTab[94080]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:236
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:236
		// _ = "end of CoverTab[94066]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:236
		_go_fuzz_dep_.CoverTab[94067]++
														if !o.customALEntropy {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:237
			_go_fuzz_dep_.CoverTab[94081]++
															o.allLitEntropy = l > SpeedFastest
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:238
			// _ = "end of CoverTab[94081]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:239
			_go_fuzz_dep_.CoverTab[94082]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:239
			// _ = "end of CoverTab[94082]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:239
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:239
		// _ = "end of CoverTab[94067]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:239
		_go_fuzz_dep_.CoverTab[94068]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:241
		// _ = "end of CoverTab[94068]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:242
	// _ = "end of CoverTab[94064]"
}

// WithZeroFrames will encode 0 length input as full frames.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:245
// This can be needed for compatibility with zstandard usage,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:245
// but is not needed for this package.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:248
func WithZeroFrames(b bool) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:248
	_go_fuzz_dep_.CoverTab[94083]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:249
		_go_fuzz_dep_.CoverTab[94084]++
														o.fullZero = b
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:251
		// _ = "end of CoverTab[94084]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:252
	// _ = "end of CoverTab[94083]"
}

// WithAllLitEntropyCompression will apply entropy compression if no matches are found.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:255
// Disabling this will skip incompressible data faster, but in cases with no matches but
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:255
// skewed character distribution compression is lost.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:255
// Default value depends on the compression level selected.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:259
func WithAllLitEntropyCompression(b bool) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:259
	_go_fuzz_dep_.CoverTab[94085]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:260
		_go_fuzz_dep_.CoverTab[94086]++
														o.customALEntropy = true
														o.allLitEntropy = b
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:263
		// _ = "end of CoverTab[94086]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:264
	// _ = "end of CoverTab[94085]"
}

// WithNoEntropyCompression will always skip entropy compression of literals.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:267
// This can be useful if content has matches, but unlikely to benefit from entropy
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:267
// compression. Usually the slight speed improvement is not worth enabling this.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:270
func WithNoEntropyCompression(b bool) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:270
	_go_fuzz_dep_.CoverTab[94087]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:271
		_go_fuzz_dep_.CoverTab[94088]++
														o.noEntropy = b
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:273
		// _ = "end of CoverTab[94088]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:274
	// _ = "end of CoverTab[94087]"
}

// WithSingleSegment will set the "single segment" flag when EncodeAll is used.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// If this flag is set, data must be regenerated within a single continuous memory segment.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// In this case, Window_Descriptor byte is skipped, but Frame_Content_Size is necessarily present.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// As a consequence, the decoder must allocate a memory segment of size equal or larger than size of your content.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// In order to preserve the decoder from unreasonable memory requirements,
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// a decoder is allowed to reject a compressed frame which requests a memory size beyond decoder's authorized range.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// For broader compatibility, decoders are recommended to support memory sizes of at least 8 MB.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// This is only a recommendation, each decoder is free to support higher or lower limits, depending on local limitations.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// If this is not specified, block encodes will automatically choose this based on the input size.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:277
// This setting has no effect on streamed encodes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:287
func WithSingleSegment(b bool) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:287
	_go_fuzz_dep_.CoverTab[94089]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:288
		_go_fuzz_dep_.CoverTab[94090]++
														o.single = &b
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:290
		// _ = "end of CoverTab[94090]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:291
	// _ = "end of CoverTab[94089]"
}

// WithLowerEncoderMem will trade in some memory cases trade less memory usage for
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:294
// slower encoding speed.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:294
// This will not change the window size which is the primary function for reducing
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:294
// memory usage. See WithWindowSize.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:298
func WithLowerEncoderMem(b bool) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:298
	_go_fuzz_dep_.CoverTab[94091]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:299
		_go_fuzz_dep_.CoverTab[94092]++
														o.lowMem = b
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:301
		// _ = "end of CoverTab[94092]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:302
	// _ = "end of CoverTab[94091]"
}

// WithEncoderDict allows to register a dictionary that will be used for the encode.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:305
// The encoder *may* choose to use no dictionary instead for certain payloads.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:307
func WithEncoderDict(dict []byte) EOption {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:307
	_go_fuzz_dep_.CoverTab[94093]++
													return func(o *encoderOptions) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:308
		_go_fuzz_dep_.CoverTab[94094]++
														d, err := loadDict(dict)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:310
			_go_fuzz_dep_.CoverTab[94096]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:311
			// _ = "end of CoverTab[94096]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:312
			_go_fuzz_dep_.CoverTab[94097]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:312
			// _ = "end of CoverTab[94097]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:312
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:312
		// _ = "end of CoverTab[94094]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:312
		_go_fuzz_dep_.CoverTab[94095]++
														o.dict = d
														return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:314
		// _ = "end of CoverTab[94095]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:315
	// _ = "end of CoverTab[94093]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:316
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/encoder_options.go:316
var _ = _go_fuzz_dep_.CoverTab
