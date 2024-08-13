// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:4
)

import (
	"errors"
	"io"
	"sync"
)

// ZipMethodWinZip is the method for Zstandard compressed data inside Zip files for WinZip.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:12
// See https://www.winzip.com/win/en/comp_info.html
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:14
const ZipMethodWinZip = 93

// ZipMethodPKWare is the original method number used by PKWARE to indicate Zstandard compression.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:16
// Deprecated: This has been deprecated by PKWARE, use ZipMethodWinZip instead for compression.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:16
// See https://pkware.cachefly.net/webdocs/APPNOTE/APPNOTE-6.3.9.TXT
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:19
const ZipMethodPKWare = 20

var zipReaderPool sync.Pool

// newZipReader cannot be used since we would leak goroutines...
func newZipReader(r io.Reader) io.ReadCloser {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:24
	_go_fuzz_dep_.CoverTab[95239]++
											dec, ok := zipReaderPool.Get().(*Decoder)
											if ok {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:26
		_go_fuzz_dep_.CoverTab[95241]++
												dec.Reset(r)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:27
		// _ = "end of CoverTab[95241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:28
		_go_fuzz_dep_.CoverTab[95242]++
												d, err := NewReader(r, WithDecoderConcurrency(1), WithDecoderLowmem(true))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:30
			_go_fuzz_dep_.CoverTab[95244]++
													panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:31
			// _ = "end of CoverTab[95244]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:32
			_go_fuzz_dep_.CoverTab[95245]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:32
			// _ = "end of CoverTab[95245]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:32
		// _ = "end of CoverTab[95242]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:32
		_go_fuzz_dep_.CoverTab[95243]++
												dec = d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:33
		// _ = "end of CoverTab[95243]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:34
	// _ = "end of CoverTab[95239]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:34
	_go_fuzz_dep_.CoverTab[95240]++
											return &pooledZipReader{dec: dec}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:35
	// _ = "end of CoverTab[95240]"
}

type pooledZipReader struct {
	mu	sync.Mutex	// guards Close and Read
	dec	*Decoder
}

func (r *pooledZipReader) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:43
	_go_fuzz_dep_.CoverTab[95246]++
											r.mu.Lock()
											defer r.mu.Unlock()
											if r.dec == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:46
		_go_fuzz_dep_.CoverTab[95248]++
												return 0, errors.New("Read after Close")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:47
		// _ = "end of CoverTab[95248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:48
		_go_fuzz_dep_.CoverTab[95249]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:48
		// _ = "end of CoverTab[95249]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:48
	// _ = "end of CoverTab[95246]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:48
	_go_fuzz_dep_.CoverTab[95247]++
											dec, err := r.dec.Read(p)

											return dec, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:51
	// _ = "end of CoverTab[95247]"
}

func (r *pooledZipReader) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:54
	_go_fuzz_dep_.CoverTab[95250]++
											r.mu.Lock()
											defer r.mu.Unlock()
											var err error
											if r.dec != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:58
		_go_fuzz_dep_.CoverTab[95252]++
												err = r.dec.Reset(nil)
												zipReaderPool.Put(r.dec)
												r.dec = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:61
		// _ = "end of CoverTab[95252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:62
		_go_fuzz_dep_.CoverTab[95253]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:62
		// _ = "end of CoverTab[95253]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:62
	// _ = "end of CoverTab[95250]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:62
	_go_fuzz_dep_.CoverTab[95251]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:63
	// _ = "end of CoverTab[95251]"
}

type pooledZipWriter struct {
	mu	sync.Mutex	// guards Close and Read
	enc	*Encoder
	pool	*sync.Pool
}

func (w *pooledZipWriter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:72
	_go_fuzz_dep_.CoverTab[95254]++
											w.mu.Lock()
											defer w.mu.Unlock()
											if w.enc == nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:75
		_go_fuzz_dep_.CoverTab[95256]++
												return 0, errors.New("Write after Close")
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:76
		// _ = "end of CoverTab[95256]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:77
		_go_fuzz_dep_.CoverTab[95257]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:77
		// _ = "end of CoverTab[95257]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:77
	// _ = "end of CoverTab[95254]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:77
	_go_fuzz_dep_.CoverTab[95255]++
											return w.enc.Write(p)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:78
	// _ = "end of CoverTab[95255]"
}

func (w *pooledZipWriter) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:81
	_go_fuzz_dep_.CoverTab[95258]++
											w.mu.Lock()
											defer w.mu.Unlock()
											var err error
											if w.enc != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:85
		_go_fuzz_dep_.CoverTab[95260]++
												err = w.enc.Close()
												w.pool.Put(w.enc)
												w.enc = nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:88
		// _ = "end of CoverTab[95260]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:89
		_go_fuzz_dep_.CoverTab[95261]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:89
		// _ = "end of CoverTab[95261]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:89
	// _ = "end of CoverTab[95258]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:89
	_go_fuzz_dep_.CoverTab[95259]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:90
	// _ = "end of CoverTab[95259]"
}

// ZipCompressor returns a compressor that can be registered with zip libraries.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:93
// The provided encoder options will be used on all encodes.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:95
func ZipCompressor(opts ...EOption) func(w io.Writer) (io.WriteCloser, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:95
	_go_fuzz_dep_.CoverTab[95262]++
											var pool sync.Pool
											return func(w io.Writer) (io.WriteCloser, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:97
		_go_fuzz_dep_.CoverTab[95263]++
												enc, ok := pool.Get().(*Encoder)
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:99
				_go_fuzz_dep_.CoverTab[95265]++
														enc.Reset(w)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:100
			// _ = "end of CoverTab[95265]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:101
			_go_fuzz_dep_.CoverTab[95266]++
														var err error
														enc, err = NewWriter(w, opts...)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:104
				_go_fuzz_dep_.CoverTab[95267]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:105
				// _ = "end of CoverTab[95267]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:106
				_go_fuzz_dep_.CoverTab[95268]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:106
				// _ = "end of CoverTab[95268]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:106
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:106
			// _ = "end of CoverTab[95266]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:107
		// _ = "end of CoverTab[95263]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:107
		_go_fuzz_dep_.CoverTab[95264]++
													return &pooledZipWriter{enc: enc, pool: &pool}, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:108
		// _ = "end of CoverTab[95264]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:109
	// _ = "end of CoverTab[95262]"
}

// ZipDecompressor returns a decompressor that can be registered with zip libraries.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:112
// See ZipCompressor for example.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:114
func ZipDecompressor() func(r io.Reader) io.ReadCloser {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:114
	_go_fuzz_dep_.CoverTab[95269]++
												return func(r io.Reader) io.ReadCloser {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:115
		_go_fuzz_dep_.CoverTab[95270]++
													d, err := NewReader(r, WithDecoderConcurrency(1), WithDecoderLowmem(true))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:117
			_go_fuzz_dep_.CoverTab[95272]++
														panic(err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:118
			// _ = "end of CoverTab[95272]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:119
			_go_fuzz_dep_.CoverTab[95273]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:119
			// _ = "end of CoverTab[95273]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:119
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:119
		// _ = "end of CoverTab[95270]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:119
		_go_fuzz_dep_.CoverTab[95271]++
													return d.IOReadCloser()
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:120
		// _ = "end of CoverTab[95271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:121
	// _ = "end of CoverTab[95269]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/zip.go:122
var _ = _go_fuzz_dep_.CoverTab
