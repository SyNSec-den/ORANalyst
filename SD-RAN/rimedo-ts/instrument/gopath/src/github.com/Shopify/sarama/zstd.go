//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:1
)

import (
	"sync"

	"github.com/klauspost/compress/zstd"
)

type ZstdEncoderParams struct {
	Level int
}
type ZstdDecoderParams struct {
}

var zstdEncMap, zstdDecMap sync.Map

func getEncoder(params ZstdEncoderParams) *zstd.Encoder {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:17
	_go_fuzz_dep_.CoverTab[107068]++
											if ret, ok := zstdEncMap.Load(params); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:18
		_go_fuzz_dep_.CoverTab[107071]++
												return ret.(*zstd.Encoder)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:19
		// _ = "end of CoverTab[107071]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:20
		_go_fuzz_dep_.CoverTab[107072]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:20
		// _ = "end of CoverTab[107072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:20
	// _ = "end of CoverTab[107068]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:20
	_go_fuzz_dep_.CoverTab[107069]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:23
	encoderLevel := zstd.SpeedDefault
	if params.Level != CompressionLevelDefault {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:24
		_go_fuzz_dep_.CoverTab[107073]++
												encoderLevel = zstd.EncoderLevelFromZstd(params.Level)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:25
		// _ = "end of CoverTab[107073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:26
		_go_fuzz_dep_.CoverTab[107074]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:26
		// _ = "end of CoverTab[107074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:26
	// _ = "end of CoverTab[107069]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:26
	_go_fuzz_dep_.CoverTab[107070]++
											zstdEnc, _ := zstd.NewWriter(nil, zstd.WithZeroFrames(true),
		zstd.WithEncoderLevel(encoderLevel))
											zstdEncMap.Store(params, zstdEnc)
											return zstdEnc
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:30
	// _ = "end of CoverTab[107070]"
}

func getDecoder(params ZstdDecoderParams) *zstd.Decoder {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:33
	_go_fuzz_dep_.CoverTab[107075]++
											if ret, ok := zstdDecMap.Load(params); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:34
		_go_fuzz_dep_.CoverTab[107077]++
												return ret.(*zstd.Decoder)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:35
		// _ = "end of CoverTab[107077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:36
		_go_fuzz_dep_.CoverTab[107078]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:36
		// _ = "end of CoverTab[107078]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:36
	// _ = "end of CoverTab[107075]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:36
	_go_fuzz_dep_.CoverTab[107076]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:39
	zstdDec, _ := zstd.NewReader(nil)
											zstdDecMap.Store(params, zstdDec)
											return zstdDec
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:41
	// _ = "end of CoverTab[107076]"
}

func zstdDecompress(params ZstdDecoderParams, dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:44
	_go_fuzz_dep_.CoverTab[107079]++
											return getDecoder(params).DecodeAll(src, dst)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:45
	// _ = "end of CoverTab[107079]"
}

func zstdCompress(params ZstdEncoderParams, dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:48
	_go_fuzz_dep_.CoverTab[107080]++
											return getEncoder(params).EncodeAll(src, dst), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:49
	// _ = "end of CoverTab[107080]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:50
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/zstd.go:50
var _ = _go_fuzz_dep_.CoverTab
