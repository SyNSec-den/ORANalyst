//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:1
)

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"sync"

	snappy "github.com/eapache/go-xerial-snappy"
	"github.com/pierrec/lz4"
)

var (
	lz4ReaderPool	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:16
			_go_fuzz_dep_.CoverTab[101680]++
													return lz4.NewReader(nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:17
			// _ = "end of CoverTab[101680]"
		},
	}

	gzipReaderPool	sync.Pool
)

func decompress(cc CompressionCodec, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:24
	_go_fuzz_dep_.CoverTab[101681]++
											switch cc {
	case CompressionNone:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:26
		_go_fuzz_dep_.CoverTab[101682]++
												return data, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:27
		// _ = "end of CoverTab[101682]"
	case CompressionGZIP:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:28
		_go_fuzz_dep_.CoverTab[101683]++
												var err error
												reader, ok := gzipReaderPool.Get().(*gzip.Reader)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:31
			_go_fuzz_dep_.CoverTab[101691]++
													reader, err = gzip.NewReader(bytes.NewReader(data))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:32
			// _ = "end of CoverTab[101691]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:33
			_go_fuzz_dep_.CoverTab[101692]++
													err = reader.Reset(bytes.NewReader(data))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:34
			// _ = "end of CoverTab[101692]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:35
		// _ = "end of CoverTab[101683]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:35
		_go_fuzz_dep_.CoverTab[101684]++

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:37
			_go_fuzz_dep_.CoverTab[101693]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:38
			// _ = "end of CoverTab[101693]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:39
			_go_fuzz_dep_.CoverTab[101694]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:39
			// _ = "end of CoverTab[101694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:39
		// _ = "end of CoverTab[101684]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:39
		_go_fuzz_dep_.CoverTab[101685]++

												defer gzipReaderPool.Put(reader)

												return io.ReadAll(reader)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:43
		// _ = "end of CoverTab[101685]"
	case CompressionSnappy:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:44
		_go_fuzz_dep_.CoverTab[101686]++
												return snappy.Decode(data)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:45
		// _ = "end of CoverTab[101686]"
	case CompressionLZ4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:46
		_go_fuzz_dep_.CoverTab[101687]++
												reader, ok := lz4ReaderPool.Get().(*lz4.Reader)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:48
			_go_fuzz_dep_.CoverTab[101695]++
													reader = lz4.NewReader(bytes.NewReader(data))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:49
			// _ = "end of CoverTab[101695]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:50
			_go_fuzz_dep_.CoverTab[101696]++
													reader.Reset(bytes.NewReader(data))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:51
			// _ = "end of CoverTab[101696]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:52
		// _ = "end of CoverTab[101687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:52
		_go_fuzz_dep_.CoverTab[101688]++
												defer lz4ReaderPool.Put(reader)

												return io.ReadAll(reader)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:55
		// _ = "end of CoverTab[101688]"
	case CompressionZSTD:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:56
		_go_fuzz_dep_.CoverTab[101689]++
												return zstdDecompress(ZstdDecoderParams{}, nil, data)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:57
		// _ = "end of CoverTab[101689]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:58
		_go_fuzz_dep_.CoverTab[101690]++
												return nil, PacketDecodingError{fmt.Sprintf("invalid compression specified (%d)", cc)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:59
		// _ = "end of CoverTab[101690]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:60
	// _ = "end of CoverTab[101681]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/decompress.go:61
var _ = _go_fuzz_dep_.CoverTab
