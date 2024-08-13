// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:5
)

import (
	"errors"
	"fmt"
	"sync"
)

// Buffer chunks are allocated from a pool to reduce pressure on GC.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// The maximum wasted space per dataBuffer is 2x the largest size class,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// which happens when the dataBuffer has multiple chunks and there is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// one unread byte in both the first and last chunks. We use a few size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// classes to minimize overheads for servers that typically receive very
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// small request bodies.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// TODO: Benchmark to determine if the pools are necessary. The GC may have
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// improved enough that we can instead allocate chunks like this:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:13
// make([]byte, max(16<<10, expectedBytesRemaining))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:23
var (
	dataChunkSizeClasses	= []int{
		1 << 10,
		2 << 10,
		4 << 10,
		8 << 10,
		16 << 10,
	}
	dataChunkPools	= [...]sync.Pool{
		{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:32
			_go_fuzz_dep_.CoverTab[72340]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:32
			return make([]byte, 1<<10)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:32
			// _ = "end of CoverTab[72340]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:32
		}},
		{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:33
			_go_fuzz_dep_.CoverTab[72341]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:33
			return make([]byte, 2<<10)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:33
			// _ = "end of CoverTab[72341]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:33
		}},
		{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:34
			_go_fuzz_dep_.CoverTab[72342]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:34
			return make([]byte, 4<<10)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:34
			// _ = "end of CoverTab[72342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:34
		}},
		{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:35
			_go_fuzz_dep_.CoverTab[72343]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:35
			return make([]byte, 8<<10)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:35
			// _ = "end of CoverTab[72343]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:35
		}},
		{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:36
			_go_fuzz_dep_.CoverTab[72344]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:36
			return make([]byte, 16<<10)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:36
			// _ = "end of CoverTab[72344]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:36
		}},
	}
)

func getDataBufferChunk(size int64) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:40
	_go_fuzz_dep_.CoverTab[72345]++
											i := 0
											for ; i < len(dataChunkSizeClasses)-1; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:42
		_go_fuzz_dep_.CoverTab[72347]++
												if size <= int64(dataChunkSizeClasses[i]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:43
			_go_fuzz_dep_.CoverTab[72348]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:44
			// _ = "end of CoverTab[72348]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:45
			_go_fuzz_dep_.CoverTab[72349]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:45
			// _ = "end of CoverTab[72349]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:45
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:45
		// _ = "end of CoverTab[72347]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:46
	// _ = "end of CoverTab[72345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:46
	_go_fuzz_dep_.CoverTab[72346]++
											return dataChunkPools[i].Get().([]byte)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:47
	// _ = "end of CoverTab[72346]"
}

func putDataBufferChunk(p []byte) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:50
	_go_fuzz_dep_.CoverTab[72350]++
											for i, n := range dataChunkSizeClasses {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:51
		_go_fuzz_dep_.CoverTab[72352]++
												if len(p) == n {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:52
			_go_fuzz_dep_.CoverTab[72353]++
													dataChunkPools[i].Put(p)
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:54
			// _ = "end of CoverTab[72353]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:55
			_go_fuzz_dep_.CoverTab[72354]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:55
			// _ = "end of CoverTab[72354]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:55
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:55
		// _ = "end of CoverTab[72352]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:56
	// _ = "end of CoverTab[72350]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:56
	_go_fuzz_dep_.CoverTab[72351]++
											panic(fmt.Sprintf("unexpected buffer len=%v", len(p)))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:57
	// _ = "end of CoverTab[72351]"
}

// dataBuffer is an io.ReadWriter backed by a list of data chunks.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:60
// Each dataBuffer is used to read DATA frames on a single stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:60
// The buffer is divided into chunks so the server can limit the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:60
// total memory used by a single connection without limiting the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:60
// request body size on any single stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:65
type dataBuffer struct {
	chunks		[][]byte
	r		int	// next byte to read is chunks[0][r]
	w		int	// next byte to write is chunks[len(chunks)-1][w]
	size		int	// total buffered bytes
	expected	int64	// we expect at least this many bytes in future Write calls (ignored if <= 0)
}

var errReadEmpty = errors.New("read from empty dataBuffer")

// Read copies bytes from the buffer into p.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:75
// It is an error to read when no data is available.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:77
func (b *dataBuffer) Read(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:77
	_go_fuzz_dep_.CoverTab[72355]++
											if b.size == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:78
		_go_fuzz_dep_.CoverTab[72358]++
												return 0, errReadEmpty
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:79
		// _ = "end of CoverTab[72358]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:80
		_go_fuzz_dep_.CoverTab[72359]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:80
		// _ = "end of CoverTab[72359]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:80
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:80
	// _ = "end of CoverTab[72355]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:80
	_go_fuzz_dep_.CoverTab[72356]++
											var ntotal int
											for len(p) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:82
		_go_fuzz_dep_.CoverTab[72360]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:82
		return b.size > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:82
		// _ = "end of CoverTab[72360]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:82
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:82
		_go_fuzz_dep_.CoverTab[72361]++
												readFrom := b.bytesFromFirstChunk()
												n := copy(p, readFrom)
												p = p[n:]
												ntotal += n
												b.r += n
												b.size -= n

												if b.r == len(b.chunks[0]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:90
			_go_fuzz_dep_.CoverTab[72362]++
													putDataBufferChunk(b.chunks[0])
													end := len(b.chunks) - 1
													copy(b.chunks[:end], b.chunks[1:])
													b.chunks[end] = nil
													b.chunks = b.chunks[:end]
													b.r = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:96
			// _ = "end of CoverTab[72362]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:97
			_go_fuzz_dep_.CoverTab[72363]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:97
			// _ = "end of CoverTab[72363]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:97
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:97
		// _ = "end of CoverTab[72361]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:98
	// _ = "end of CoverTab[72356]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:98
	_go_fuzz_dep_.CoverTab[72357]++
											return ntotal, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:99
	// _ = "end of CoverTab[72357]"
}

func (b *dataBuffer) bytesFromFirstChunk() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:102
	_go_fuzz_dep_.CoverTab[72364]++
											if len(b.chunks) == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:103
		_go_fuzz_dep_.CoverTab[72366]++
												return b.chunks[0][b.r:b.w]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:104
		// _ = "end of CoverTab[72366]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:105
		_go_fuzz_dep_.CoverTab[72367]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:105
		// _ = "end of CoverTab[72367]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:105
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:105
	// _ = "end of CoverTab[72364]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:105
	_go_fuzz_dep_.CoverTab[72365]++
											return b.chunks[0][b.r:]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:106
	// _ = "end of CoverTab[72365]"
}

// Len returns the number of bytes of the unread portion of the buffer.
func (b *dataBuffer) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:110
	_go_fuzz_dep_.CoverTab[72368]++
											return b.size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:111
	// _ = "end of CoverTab[72368]"
}

// Write appends p to the buffer.
func (b *dataBuffer) Write(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:115
	_go_fuzz_dep_.CoverTab[72369]++
											ntotal := len(p)
											for len(p) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:117
		_go_fuzz_dep_.CoverTab[72371]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:121
		want := int64(len(p))
		if b.expected > want {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:122
			_go_fuzz_dep_.CoverTab[72373]++
													want = b.expected
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:123
			// _ = "end of CoverTab[72373]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:124
			_go_fuzz_dep_.CoverTab[72374]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:124
			// _ = "end of CoverTab[72374]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:124
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:124
		// _ = "end of CoverTab[72371]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:124
		_go_fuzz_dep_.CoverTab[72372]++
												chunk := b.lastChunkOrAlloc(want)
												n := copy(chunk[b.w:], p)
												p = p[n:]
												b.w += n
												b.size += n
												b.expected -= int64(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:130
		// _ = "end of CoverTab[72372]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:131
	// _ = "end of CoverTab[72369]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:131
	_go_fuzz_dep_.CoverTab[72370]++
											return ntotal, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:132
	// _ = "end of CoverTab[72370]"
}

func (b *dataBuffer) lastChunkOrAlloc(want int64) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:135
	_go_fuzz_dep_.CoverTab[72375]++
											if len(b.chunks) != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:136
		_go_fuzz_dep_.CoverTab[72377]++
												last := b.chunks[len(b.chunks)-1]
												if b.w < len(last) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:138
			_go_fuzz_dep_.CoverTab[72378]++
													return last
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:139
			// _ = "end of CoverTab[72378]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:140
			_go_fuzz_dep_.CoverTab[72379]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:140
			// _ = "end of CoverTab[72379]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:140
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:140
		// _ = "end of CoverTab[72377]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:141
		_go_fuzz_dep_.CoverTab[72380]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:141
		// _ = "end of CoverTab[72380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:141
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:141
	// _ = "end of CoverTab[72375]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:141
	_go_fuzz_dep_.CoverTab[72376]++
											chunk := getDataBufferChunk(want)
											b.chunks = append(b.chunks, chunk)
											b.w = 0
											return chunk
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:145
	// _ = "end of CoverTab[72376]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:146
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/databuffer.go:146
var _ = _go_fuzz_dep_.CoverTab
