//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:1
)

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"sync"

	snappy "github.com/eapache/go-xerial-snappy"
	"github.com/pierrec/lz4"
)

var (
	lz4WriterPool	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:15
			_go_fuzz_dep_.CoverTab[100353]++
													return lz4.NewWriter(nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:16
			// _ = "end of CoverTab[100353]"
		},
	}

	gzipWriterPool	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:21
			_go_fuzz_dep_.CoverTab[100354]++
													return gzip.NewWriter(nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:22
			// _ = "end of CoverTab[100354]"
		},
	}
	gzipWriterPoolForCompressionLevel1	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:26
			_go_fuzz_dep_.CoverTab[100355]++
													gz, err := gzip.NewWriterLevel(nil, 1)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:28
				_go_fuzz_dep_.CoverTab[100357]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:29
				// _ = "end of CoverTab[100357]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:30
				_go_fuzz_dep_.CoverTab[100358]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:30
				// _ = "end of CoverTab[100358]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:30
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:30
			// _ = "end of CoverTab[100355]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:30
			_go_fuzz_dep_.CoverTab[100356]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:31
			// _ = "end of CoverTab[100356]"
		},
	}
	gzipWriterPoolForCompressionLevel2	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:35
			_go_fuzz_dep_.CoverTab[100359]++
													gz, err := gzip.NewWriterLevel(nil, 2)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:37
				_go_fuzz_dep_.CoverTab[100361]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:38
				// _ = "end of CoverTab[100361]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:39
				_go_fuzz_dep_.CoverTab[100362]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:39
				// _ = "end of CoverTab[100362]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:39
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:39
			// _ = "end of CoverTab[100359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:39
			_go_fuzz_dep_.CoverTab[100360]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:40
			// _ = "end of CoverTab[100360]"
		},
	}
	gzipWriterPoolForCompressionLevel3	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:44
			_go_fuzz_dep_.CoverTab[100363]++
													gz, err := gzip.NewWriterLevel(nil, 3)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:46
				_go_fuzz_dep_.CoverTab[100365]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:47
				// _ = "end of CoverTab[100365]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:48
				_go_fuzz_dep_.CoverTab[100366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:48
				// _ = "end of CoverTab[100366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:48
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:48
			// _ = "end of CoverTab[100363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:48
			_go_fuzz_dep_.CoverTab[100364]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:49
			// _ = "end of CoverTab[100364]"
		},
	}
	gzipWriterPoolForCompressionLevel4	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:53
			_go_fuzz_dep_.CoverTab[100367]++
													gz, err := gzip.NewWriterLevel(nil, 4)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:55
				_go_fuzz_dep_.CoverTab[100369]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:56
				// _ = "end of CoverTab[100369]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:57
				_go_fuzz_dep_.CoverTab[100370]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:57
				// _ = "end of CoverTab[100370]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:57
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:57
			// _ = "end of CoverTab[100367]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:57
			_go_fuzz_dep_.CoverTab[100368]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:58
			// _ = "end of CoverTab[100368]"
		},
	}
	gzipWriterPoolForCompressionLevel5	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:62
			_go_fuzz_dep_.CoverTab[100371]++
													gz, err := gzip.NewWriterLevel(nil, 5)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:64
				_go_fuzz_dep_.CoverTab[100373]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:65
				// _ = "end of CoverTab[100373]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:66
				_go_fuzz_dep_.CoverTab[100374]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:66
				// _ = "end of CoverTab[100374]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:66
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:66
			// _ = "end of CoverTab[100371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:66
			_go_fuzz_dep_.CoverTab[100372]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:67
			// _ = "end of CoverTab[100372]"
		},
	}
	gzipWriterPoolForCompressionLevel6	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:71
			_go_fuzz_dep_.CoverTab[100375]++
													gz, err := gzip.NewWriterLevel(nil, 6)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:73
				_go_fuzz_dep_.CoverTab[100377]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:74
				// _ = "end of CoverTab[100377]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:75
				_go_fuzz_dep_.CoverTab[100378]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:75
				// _ = "end of CoverTab[100378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:75
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:75
			// _ = "end of CoverTab[100375]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:75
			_go_fuzz_dep_.CoverTab[100376]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:76
			// _ = "end of CoverTab[100376]"
		},
	}
	gzipWriterPoolForCompressionLevel7	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:80
			_go_fuzz_dep_.CoverTab[100379]++
													gz, err := gzip.NewWriterLevel(nil, 7)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:82
				_go_fuzz_dep_.CoverTab[100381]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:83
				// _ = "end of CoverTab[100381]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:84
				_go_fuzz_dep_.CoverTab[100382]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:84
				// _ = "end of CoverTab[100382]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:84
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:84
			// _ = "end of CoverTab[100379]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:84
			_go_fuzz_dep_.CoverTab[100380]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:85
			// _ = "end of CoverTab[100380]"
		},
	}
	gzipWriterPoolForCompressionLevel8	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:89
			_go_fuzz_dep_.CoverTab[100383]++
													gz, err := gzip.NewWriterLevel(nil, 8)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:91
				_go_fuzz_dep_.CoverTab[100385]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:92
				// _ = "end of CoverTab[100385]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:93
				_go_fuzz_dep_.CoverTab[100386]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:93
				// _ = "end of CoverTab[100386]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:93
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:93
			// _ = "end of CoverTab[100383]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:93
			_go_fuzz_dep_.CoverTab[100384]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:94
			// _ = "end of CoverTab[100384]"
		},
	}
	gzipWriterPoolForCompressionLevel9	= sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:98
			_go_fuzz_dep_.CoverTab[100387]++
													gz, err := gzip.NewWriterLevel(nil, 9)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:100
				_go_fuzz_dep_.CoverTab[100389]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:101
				// _ = "end of CoverTab[100389]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:102
				_go_fuzz_dep_.CoverTab[100390]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:102
				// _ = "end of CoverTab[100390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:102
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:102
			// _ = "end of CoverTab[100387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:102
			_go_fuzz_dep_.CoverTab[100388]++
													return gz
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:103
			// _ = "end of CoverTab[100388]"
		},
	}
)

func compress(cc CompressionCodec, level int, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:108
	_go_fuzz_dep_.CoverTab[100391]++
											switch cc {
	case CompressionNone:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:110
		_go_fuzz_dep_.CoverTab[100392]++
												return data, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:111
		// _ = "end of CoverTab[100392]"
	case CompressionGZIP:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:112
		_go_fuzz_dep_.CoverTab[100393]++
												var (
			err	error
			buf	bytes.Buffer
			writer	*gzip.Writer
		)

		switch level {
		case CompressionLevelDefault:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:120
			_go_fuzz_dep_.CoverTab[100403]++
													writer = gzipWriterPool.Get().(*gzip.Writer)
													defer gzipWriterPool.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:123
			// _ = "end of CoverTab[100403]"
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:124
			_go_fuzz_dep_.CoverTab[100404]++
													writer = gzipWriterPoolForCompressionLevel1.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel1.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:127
			// _ = "end of CoverTab[100404]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:128
			_go_fuzz_dep_.CoverTab[100405]++
													writer = gzipWriterPoolForCompressionLevel2.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel2.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:131
			// _ = "end of CoverTab[100405]"
		case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:132
			_go_fuzz_dep_.CoverTab[100406]++
													writer = gzipWriterPoolForCompressionLevel3.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel3.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:135
			// _ = "end of CoverTab[100406]"
		case 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:136
			_go_fuzz_dep_.CoverTab[100407]++
													writer = gzipWriterPoolForCompressionLevel4.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel4.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:139
			// _ = "end of CoverTab[100407]"
		case 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:140
			_go_fuzz_dep_.CoverTab[100408]++
													writer = gzipWriterPoolForCompressionLevel5.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel5.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:143
			// _ = "end of CoverTab[100408]"
		case 6:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:144
			_go_fuzz_dep_.CoverTab[100409]++
													writer = gzipWriterPoolForCompressionLevel6.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel6.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:147
			// _ = "end of CoverTab[100409]"
		case 7:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:148
			_go_fuzz_dep_.CoverTab[100410]++
													writer = gzipWriterPoolForCompressionLevel7.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel7.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:151
			// _ = "end of CoverTab[100410]"
		case 8:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:152
			_go_fuzz_dep_.CoverTab[100411]++
													writer = gzipWriterPoolForCompressionLevel8.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel8.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:155
			// _ = "end of CoverTab[100411]"
		case 9:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:156
			_go_fuzz_dep_.CoverTab[100412]++
													writer = gzipWriterPoolForCompressionLevel9.Get().(*gzip.Writer)
													defer gzipWriterPoolForCompressionLevel9.Put(writer)
													writer.Reset(&buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:159
			// _ = "end of CoverTab[100412]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:160
			_go_fuzz_dep_.CoverTab[100413]++
													writer, err = gzip.NewWriterLevel(&buf, level)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:162
				_go_fuzz_dep_.CoverTab[100414]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:163
				// _ = "end of CoverTab[100414]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:164
				_go_fuzz_dep_.CoverTab[100415]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:164
				// _ = "end of CoverTab[100415]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:164
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:164
			// _ = "end of CoverTab[100413]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:165
		// _ = "end of CoverTab[100393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:165
		_go_fuzz_dep_.CoverTab[100394]++
												if _, err := writer.Write(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:166
			_go_fuzz_dep_.CoverTab[100416]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:167
			// _ = "end of CoverTab[100416]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:168
			_go_fuzz_dep_.CoverTab[100417]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:168
			// _ = "end of CoverTab[100417]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:168
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:168
		// _ = "end of CoverTab[100394]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:168
		_go_fuzz_dep_.CoverTab[100395]++
												if err := writer.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:169
			_go_fuzz_dep_.CoverTab[100418]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:170
			// _ = "end of CoverTab[100418]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:171
			_go_fuzz_dep_.CoverTab[100419]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:171
			// _ = "end of CoverTab[100419]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:171
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:171
		// _ = "end of CoverTab[100395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:171
		_go_fuzz_dep_.CoverTab[100396]++
												return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:172
		// _ = "end of CoverTab[100396]"
	case CompressionSnappy:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:173
		_go_fuzz_dep_.CoverTab[100397]++
												return snappy.Encode(data), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:174
		// _ = "end of CoverTab[100397]"
	case CompressionLZ4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:175
		_go_fuzz_dep_.CoverTab[100398]++
												writer := lz4WriterPool.Get().(*lz4.Writer)
												defer lz4WriterPool.Put(writer)

												var buf bytes.Buffer
												writer.Reset(&buf)

												if _, err := writer.Write(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:182
			_go_fuzz_dep_.CoverTab[100420]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:183
			// _ = "end of CoverTab[100420]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:184
			_go_fuzz_dep_.CoverTab[100421]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:184
			// _ = "end of CoverTab[100421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:184
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:184
		// _ = "end of CoverTab[100398]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:184
		_go_fuzz_dep_.CoverTab[100399]++
												if err := writer.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:185
			_go_fuzz_dep_.CoverTab[100422]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:186
			// _ = "end of CoverTab[100422]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:187
			_go_fuzz_dep_.CoverTab[100423]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:187
			// _ = "end of CoverTab[100423]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:187
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:187
		// _ = "end of CoverTab[100399]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:187
		_go_fuzz_dep_.CoverTab[100400]++
												return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:188
		// _ = "end of CoverTab[100400]"
	case CompressionZSTD:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:189
		_go_fuzz_dep_.CoverTab[100401]++
												return zstdCompress(ZstdEncoderParams{level}, nil, data)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:190
		// _ = "end of CoverTab[100401]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:191
		_go_fuzz_dep_.CoverTab[100402]++
												return nil, PacketEncodingError{fmt.Sprintf("unsupported compression codec (%d)", cc)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:192
		// _ = "end of CoverTab[100402]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:193
	// _ = "end of CoverTab[100391]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:194
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/compress.go:194
var _ = _go_fuzz_dep_.CoverTab
