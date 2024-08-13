// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:5
)

import (
	"fmt"
	"io"
	"io/ioutil"
)

type byteBuffer interface {
	// Read up to 8 bytes.
	// Returns io.ErrUnexpectedEOF if this cannot be satisfied.
	readSmall(n int) ([]byte, error)

	// Read >8 bytes.
	// MAY use the destination slice.
	readBig(n int, dst []byte) ([]byte, error)

	// Read a single byte.
	readByte() (byte, error)

	// Skip n bytes.
	skipN(n int) error
}

// in-memory buffer
type byteBuf []byte

func (b *byteBuf) readSmall(n int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:32
	_go_fuzz_dep_.CoverTab[91553]++
												if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:33
		_go_fuzz_dep_.CoverTab[91556]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:33
		return n > 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:33
		// _ = "end of CoverTab[91556]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:33
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:33
		_go_fuzz_dep_.CoverTab[91557]++
													panic(fmt.Errorf("small read > 8 (%d). use readBig", n))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:34
		// _ = "end of CoverTab[91557]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:35
		_go_fuzz_dep_.CoverTab[91558]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:35
		// _ = "end of CoverTab[91558]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:35
	// _ = "end of CoverTab[91553]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:35
	_go_fuzz_dep_.CoverTab[91554]++
												bb := *b
												if len(bb) < n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:37
		_go_fuzz_dep_.CoverTab[91559]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:38
		// _ = "end of CoverTab[91559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:39
		_go_fuzz_dep_.CoverTab[91560]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:39
		// _ = "end of CoverTab[91560]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:39
	// _ = "end of CoverTab[91554]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:39
	_go_fuzz_dep_.CoverTab[91555]++
												r := bb[:n]
												*b = bb[n:]
												return r, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:42
	// _ = "end of CoverTab[91555]"
}

func (b *byteBuf) readBig(n int, dst []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:45
	_go_fuzz_dep_.CoverTab[91561]++
												bb := *b
												if len(bb) < n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:47
		_go_fuzz_dep_.CoverTab[91563]++
													return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:48
		// _ = "end of CoverTab[91563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:49
		_go_fuzz_dep_.CoverTab[91564]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:49
		// _ = "end of CoverTab[91564]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:49
	// _ = "end of CoverTab[91561]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:49
	_go_fuzz_dep_.CoverTab[91562]++
												r := bb[:n]
												*b = bb[n:]
												return r, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:52
	// _ = "end of CoverTab[91562]"
}

func (b *byteBuf) remain() []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:55
	_go_fuzz_dep_.CoverTab[91565]++
												return *b
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:56
	// _ = "end of CoverTab[91565]"
}

func (b *byteBuf) readByte() (byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:59
	_go_fuzz_dep_.CoverTab[91566]++
												bb := *b
												if len(bb) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:61
		_go_fuzz_dep_.CoverTab[91568]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:62
		// _ = "end of CoverTab[91568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:63
		_go_fuzz_dep_.CoverTab[91569]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:63
		// _ = "end of CoverTab[91569]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:63
	// _ = "end of CoverTab[91566]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:63
	_go_fuzz_dep_.CoverTab[91567]++
												r := bb[0]
												*b = bb[1:]
												return r, nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:66
	// _ = "end of CoverTab[91567]"
}

func (b *byteBuf) skipN(n int) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:69
	_go_fuzz_dep_.CoverTab[91570]++
												bb := *b
												if len(bb) < n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:71
		_go_fuzz_dep_.CoverTab[91572]++
													return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:72
		// _ = "end of CoverTab[91572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:73
		_go_fuzz_dep_.CoverTab[91573]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:73
		// _ = "end of CoverTab[91573]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:73
	// _ = "end of CoverTab[91570]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:73
	_go_fuzz_dep_.CoverTab[91571]++
												*b = bb[n:]
												return nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:75
	// _ = "end of CoverTab[91571]"
}

// wrapper around a reader.
type readerWrapper struct {
	r	io.Reader
	tmp	[8]byte
}

func (r *readerWrapper) readSmall(n int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:84
	_go_fuzz_dep_.CoverTab[91574]++
												if debugAsserts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:85
		_go_fuzz_dep_.CoverTab[91577]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:85
		return n > 8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:85
		// _ = "end of CoverTab[91577]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:85
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:85
		_go_fuzz_dep_.CoverTab[91578]++
													panic(fmt.Errorf("small read > 8 (%d). use readBig", n))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:86
		// _ = "end of CoverTab[91578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:87
		_go_fuzz_dep_.CoverTab[91579]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:87
		// _ = "end of CoverTab[91579]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:87
	// _ = "end of CoverTab[91574]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:87
	_go_fuzz_dep_.CoverTab[91575]++
												n2, err := io.ReadFull(r.r, r.tmp[:n])

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:90
		_go_fuzz_dep_.CoverTab[91580]++
													if err == io.EOF {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:91
			_go_fuzz_dep_.CoverTab[91583]++
														return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:92
			// _ = "end of CoverTab[91583]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:93
			_go_fuzz_dep_.CoverTab[91584]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:93
			// _ = "end of CoverTab[91584]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:93
		// _ = "end of CoverTab[91580]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:93
		_go_fuzz_dep_.CoverTab[91581]++
													if debugDecoder {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:94
			_go_fuzz_dep_.CoverTab[91585]++
														println("readSmall: got", n2, "want", n, "err", err)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:95
			// _ = "end of CoverTab[91585]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:96
			_go_fuzz_dep_.CoverTab[91586]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:96
			// _ = "end of CoverTab[91586]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:96
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:96
		// _ = "end of CoverTab[91581]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:96
		_go_fuzz_dep_.CoverTab[91582]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:97
		// _ = "end of CoverTab[91582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:98
		_go_fuzz_dep_.CoverTab[91587]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:98
		// _ = "end of CoverTab[91587]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:98
	// _ = "end of CoverTab[91575]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:98
	_go_fuzz_dep_.CoverTab[91576]++
												return r.tmp[:n], nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:99
	// _ = "end of CoverTab[91576]"
}

func (r *readerWrapper) readBig(n int, dst []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:102
	_go_fuzz_dep_.CoverTab[91588]++
												if cap(dst) < n {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:103
		_go_fuzz_dep_.CoverTab[91591]++
													dst = make([]byte, n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:104
		// _ = "end of CoverTab[91591]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:105
		_go_fuzz_dep_.CoverTab[91592]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:105
		// _ = "end of CoverTab[91592]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:105
	// _ = "end of CoverTab[91588]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:105
	_go_fuzz_dep_.CoverTab[91589]++
												n2, err := io.ReadFull(r.r, dst[:n])
												if err == io.EOF && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:107
		_go_fuzz_dep_.CoverTab[91593]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:107
		return n > 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:107
		// _ = "end of CoverTab[91593]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:107
	}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:107
		_go_fuzz_dep_.CoverTab[91594]++
													err = io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:108
		// _ = "end of CoverTab[91594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:109
		_go_fuzz_dep_.CoverTab[91595]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:109
		// _ = "end of CoverTab[91595]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:109
	// _ = "end of CoverTab[91589]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:109
	_go_fuzz_dep_.CoverTab[91590]++
												return dst[:n2], err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:110
	// _ = "end of CoverTab[91590]"
}

func (r *readerWrapper) readByte() (byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:113
	_go_fuzz_dep_.CoverTab[91596]++
												n2, err := r.r.Read(r.tmp[:1])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:115
		_go_fuzz_dep_.CoverTab[91599]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:116
		// _ = "end of CoverTab[91599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:117
		_go_fuzz_dep_.CoverTab[91600]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:117
		// _ = "end of CoverTab[91600]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:117
	// _ = "end of CoverTab[91596]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:117
	_go_fuzz_dep_.CoverTab[91597]++
												if n2 != 1 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:118
		_go_fuzz_dep_.CoverTab[91601]++
													return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:119
		// _ = "end of CoverTab[91601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:120
		_go_fuzz_dep_.CoverTab[91602]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:120
		// _ = "end of CoverTab[91602]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:120
	// _ = "end of CoverTab[91597]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:120
	_go_fuzz_dep_.CoverTab[91598]++
												return r.tmp[0], nil
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:121
	// _ = "end of CoverTab[91598]"
}

func (r *readerWrapper) skipN(n int) error {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:124
	_go_fuzz_dep_.CoverTab[91603]++
												n2, err := io.CopyN(ioutil.Discard, r.r, int64(n))
												if n2 != int64(n) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:126
		_go_fuzz_dep_.CoverTab[91605]++
													err = io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:127
		// _ = "end of CoverTab[91605]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:128
		_go_fuzz_dep_.CoverTab[91606]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:128
		// _ = "end of CoverTab[91606]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:128
	// _ = "end of CoverTab[91603]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:128
	_go_fuzz_dep_.CoverTab[91604]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:129
	// _ = "end of CoverTab[91604]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytebuf.go:130
var _ = _go_fuzz_dep_.CoverTab
