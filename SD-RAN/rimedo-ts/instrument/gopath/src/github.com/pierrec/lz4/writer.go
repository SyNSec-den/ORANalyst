//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:1
)

import (
	"encoding/binary"
	"fmt"
	"io"
	"runtime"

	"github.com/pierrec/lz4/internal/xxh32"
)

// zResult contains the results of compressing a block.
type zResult struct {
	size		uint32	// Block header
	data		[]byte	// Compressed data
	checksum	uint32	// Data checksum
}

// Writer implements the LZ4 frame encoder.
type Writer struct {
	Header
	// Handler called when a block has been successfully written out.
	// It provides the number of bytes written.
	OnBlockDone	func(size int)

	buf		[19]byte	// magic number(4) + header(flags(2)+[Size(8)+DictID(4)]+checksum(1)) does not exceed 19 bytes
	dst		io.Writer	// Destination.
	checksum	xxh32.XXHZero	// Frame checksum.
	data		[]byte		// Data to be compressed + buffer for compressed data.
	idx		int		// Index into data.
	hashtable	[winSize]int	// Hash table used in CompressBlock().

	// For concurrency.
	c	chan chan zResult	// Channel for block compression goroutines and writer goroutine.
	err	error			// Any error encountered while writing to the underlying destination.
}

// NewWriter returns a new LZ4 frame encoder.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:38
// No access to the underlying io.Writer is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:38
// The supplied Header is checked at the first Write.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:38
// It is ok to change it before the first Write but then not until a Reset() is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:42
func NewWriter(dst io.Writer) *Writer {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:42
	_go_fuzz_dep_.CoverTab[95744]++
												z := new(Writer)
												z.Reset(dst)
												return z
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:45
	// _ = "end of CoverTab[95744]"
}

// WithConcurrency sets the number of concurrent go routines used for compression.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:48
// A negative value sets the concurrency to GOMAXPROCS.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:50
func (z *Writer) WithConcurrency(n int) *Writer {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:50
	_go_fuzz_dep_.CoverTab[95745]++
												switch {
	case n == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:52
		_go_fuzz_dep_.CoverTab[95751]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:52
		return n == 1
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:52
		// _ = "end of CoverTab[95751]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:52
	}():
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:52
		_go_fuzz_dep_.CoverTab[95748]++
													z.c = nil
													return z
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:54
		// _ = "end of CoverTab[95748]"
	case n < 0:
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:55
		_go_fuzz_dep_.CoverTab[95749]++
													n = runtime.GOMAXPROCS(0)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:56
		// _ = "end of CoverTab[95749]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:56
	default:
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:56
		_go_fuzz_dep_.CoverTab[95750]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:56
		// _ = "end of CoverTab[95750]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:57
	// _ = "end of CoverTab[95745]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:57
	_go_fuzz_dep_.CoverTab[95746]++
												z.c = make(chan chan zResult, n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:58
	_curRoutineNum108_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:58
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum108_)

												go func() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:60
		_go_fuzz_dep_.CoverTab[95752]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:60
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:60
			_go_fuzz_dep_.CoverTab[95753]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:60
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum108_)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:60
			// _ = "end of CoverTab[95753]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:60
		}()

													for c := range z.c {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:62
			_go_fuzz_dep_.CoverTab[95754]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:67
			res := <-c
			n := len(res.data)
			if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:69
				_go_fuzz_dep_.CoverTab[95760]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:72
				close(c)
															return
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:73
				// _ = "end of CoverTab[95760]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:74
				_go_fuzz_dep_.CoverTab[95761]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:74
				// _ = "end of CoverTab[95761]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:74
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:74
			// _ = "end of CoverTab[95754]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:74
			_go_fuzz_dep_.CoverTab[95755]++

														if err := z.writeUint32(res.size); err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:76
				_go_fuzz_dep_.CoverTab[95762]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:76
				return z.err == nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:76
				// _ = "end of CoverTab[95762]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:76
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:76
				_go_fuzz_dep_.CoverTab[95763]++
															z.err = err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:77
				// _ = "end of CoverTab[95763]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:78
				_go_fuzz_dep_.CoverTab[95764]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:78
				// _ = "end of CoverTab[95764]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:78
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:78
			// _ = "end of CoverTab[95755]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:78
			_go_fuzz_dep_.CoverTab[95756]++
														if _, err := z.dst.Write(res.data); err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:79
				_go_fuzz_dep_.CoverTab[95765]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:79
				return z.err == nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:79
				// _ = "end of CoverTab[95765]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:79
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:79
				_go_fuzz_dep_.CoverTab[95766]++
															z.err = err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:80
				// _ = "end of CoverTab[95766]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:81
				_go_fuzz_dep_.CoverTab[95767]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:81
				// _ = "end of CoverTab[95767]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:81
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:81
			// _ = "end of CoverTab[95756]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:81
			_go_fuzz_dep_.CoverTab[95757]++
														if z.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:82
				_go_fuzz_dep_.CoverTab[95768]++
															if err := z.writeUint32(res.checksum); err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:83
					_go_fuzz_dep_.CoverTab[95769]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:83
					return z.err == nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:83
					// _ = "end of CoverTab[95769]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:83
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:83
					_go_fuzz_dep_.CoverTab[95770]++
																z.err = err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:84
					// _ = "end of CoverTab[95770]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:85
					_go_fuzz_dep_.CoverTab[95771]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:85
					// _ = "end of CoverTab[95771]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:85
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:85
				// _ = "end of CoverTab[95768]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:86
				_go_fuzz_dep_.CoverTab[95772]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:86
				// _ = "end of CoverTab[95772]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:86
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:86
			// _ = "end of CoverTab[95757]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:86
			_go_fuzz_dep_.CoverTab[95758]++

														putBuffer(cap(res.data), res.data)
														if h := z.OnBlockDone; h != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:89
				_go_fuzz_dep_.CoverTab[95773]++
															h(n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:90
				// _ = "end of CoverTab[95773]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:91
				_go_fuzz_dep_.CoverTab[95774]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:91
				// _ = "end of CoverTab[95774]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:91
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:91
			// _ = "end of CoverTab[95758]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:91
			_go_fuzz_dep_.CoverTab[95759]++
														close(c)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:92
			// _ = "end of CoverTab[95759]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:93
		// _ = "end of CoverTab[95752]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:94
	// _ = "end of CoverTab[95746]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:94
	_go_fuzz_dep_.CoverTab[95747]++
												return z
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:95
	// _ = "end of CoverTab[95747]"
}

// newBuffers instantiates new buffers which size matches the one in Header.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:98
// The returned buffers are for decompression and compression respectively.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:100
func (z *Writer) newBuffers() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:100
	_go_fuzz_dep_.CoverTab[95775]++
												bSize := z.Header.BlockMaxSize
												buf := getBuffer(bSize)
												z.data = buf[:bSize]
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:103
	// _ = "end of CoverTab[95775]"
}

// freeBuffers puts the writer's buffers back to the pool.
func (z *Writer) freeBuffers() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:107
	_go_fuzz_dep_.CoverTab[95776]++

												putBuffer(z.Header.BlockMaxSize, z.data)
												z.data = nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:110
	// _ = "end of CoverTab[95776]"
}

// writeHeader builds and writes the header (magic+header) to the underlying io.Writer.
func (z *Writer) writeHeader() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:114
	_go_fuzz_dep_.CoverTab[95777]++

												if z.Header.BlockMaxSize == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:116
		_go_fuzz_dep_.CoverTab[95786]++
													z.Header.BlockMaxSize = blockSize4M
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:117
		// _ = "end of CoverTab[95786]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:118
		_go_fuzz_dep_.CoverTab[95787]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:118
		// _ = "end of CoverTab[95787]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:118
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:118
	// _ = "end of CoverTab[95777]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:118
	_go_fuzz_dep_.CoverTab[95778]++

												bSize := z.Header.BlockMaxSize
												if !isValidBlockSize(z.Header.BlockMaxSize) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:121
		_go_fuzz_dep_.CoverTab[95788]++
													return fmt.Errorf("lz4: invalid block max size: %d", bSize)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:122
		// _ = "end of CoverTab[95788]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:123
		_go_fuzz_dep_.CoverTab[95789]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:123
		// _ = "end of CoverTab[95789]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:123
	// _ = "end of CoverTab[95778]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:123
	_go_fuzz_dep_.CoverTab[95779]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:126
	z.newBuffers()
												z.idx = 0

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:130
	buf := z.buf[:]

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:133
	binary.LittleEndian.PutUint32(buf[0:], frameMagic)
	flg := byte(Version << 6)
	flg |= 1 << 5
	if z.Header.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:136
		_go_fuzz_dep_.CoverTab[95790]++
													flg |= 1 << 4
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:137
		// _ = "end of CoverTab[95790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:138
		_go_fuzz_dep_.CoverTab[95791]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:138
		// _ = "end of CoverTab[95791]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:138
	// _ = "end of CoverTab[95779]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:138
	_go_fuzz_dep_.CoverTab[95780]++
												if z.Header.Size > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:139
		_go_fuzz_dep_.CoverTab[95792]++
													flg |= 1 << 3
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:140
		// _ = "end of CoverTab[95792]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:141
		_go_fuzz_dep_.CoverTab[95793]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:141
		// _ = "end of CoverTab[95793]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:141
	// _ = "end of CoverTab[95780]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:141
	_go_fuzz_dep_.CoverTab[95781]++
												if !z.Header.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:142
		_go_fuzz_dep_.CoverTab[95794]++
													flg |= 1 << 2
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:143
		// _ = "end of CoverTab[95794]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:144
		_go_fuzz_dep_.CoverTab[95795]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:144
		// _ = "end of CoverTab[95795]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:144
	// _ = "end of CoverTab[95781]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:144
	_go_fuzz_dep_.CoverTab[95782]++
												buf[4] = flg
												buf[5] = blockSizeValueToIndex(z.Header.BlockMaxSize) << 4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:149
	n := 6

	if z.Header.Size > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:151
		_go_fuzz_dep_.CoverTab[95796]++
													binary.LittleEndian.PutUint64(buf[n:], z.Header.Size)
													n += 8
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:153
		// _ = "end of CoverTab[95796]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:154
		_go_fuzz_dep_.CoverTab[95797]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:154
		// _ = "end of CoverTab[95797]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:154
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:154
	// _ = "end of CoverTab[95782]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:154
	_go_fuzz_dep_.CoverTab[95783]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:157
	buf[n] = byte(xxh32.ChecksumZero(buf[4:n]) >> 8 & 0xFF)
												z.checksum.Reset()

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:161
	if _, err := z.dst.Write(buf[0 : n+1]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:161
		_go_fuzz_dep_.CoverTab[95798]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:162
		// _ = "end of CoverTab[95798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:163
		_go_fuzz_dep_.CoverTab[95799]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:163
		// _ = "end of CoverTab[95799]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:163
	// _ = "end of CoverTab[95783]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:163
	_go_fuzz_dep_.CoverTab[95784]++
												z.Header.done = true
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:165
		_go_fuzz_dep_.CoverTab[95800]++
													debug("wrote header %v", z.Header)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:166
		// _ = "end of CoverTab[95800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:167
		_go_fuzz_dep_.CoverTab[95801]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:167
		// _ = "end of CoverTab[95801]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:167
	// _ = "end of CoverTab[95784]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:167
	_go_fuzz_dep_.CoverTab[95785]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:169
	// _ = "end of CoverTab[95785]"
}

// Write compresses data from the supplied buffer into the underlying io.Writer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:172
// Write does not return until the data has been written.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:174
func (z *Writer) Write(buf []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:174
	_go_fuzz_dep_.CoverTab[95802]++
												if !z.Header.done {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:175
		_go_fuzz_dep_.CoverTab[95806]++
													if err := z.writeHeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:176
			_go_fuzz_dep_.CoverTab[95807]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:177
			// _ = "end of CoverTab[95807]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:178
			_go_fuzz_dep_.CoverTab[95808]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:178
			// _ = "end of CoverTab[95808]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:178
		// _ = "end of CoverTab[95806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:179
		_go_fuzz_dep_.CoverTab[95809]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:179
		// _ = "end of CoverTab[95809]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:179
	// _ = "end of CoverTab[95802]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:179
	_go_fuzz_dep_.CoverTab[95803]++
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:180
		_go_fuzz_dep_.CoverTab[95810]++
													debug("input buffer len=%d index=%d", len(buf), z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:181
		// _ = "end of CoverTab[95810]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:182
		_go_fuzz_dep_.CoverTab[95811]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:182
		// _ = "end of CoverTab[95811]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:182
	// _ = "end of CoverTab[95803]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:182
	_go_fuzz_dep_.CoverTab[95804]++

												zn := len(z.data)
												var n int
												for len(buf) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:186
		_go_fuzz_dep_.CoverTab[95812]++
													if z.idx == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:187
			_go_fuzz_dep_.CoverTab[95817]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:187
			return len(buf) >= zn
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:187
			// _ = "end of CoverTab[95817]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:187
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:187
			_go_fuzz_dep_.CoverTab[95818]++

														if err := z.compressBlock(buf[:zn]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:189
				_go_fuzz_dep_.CoverTab[95820]++
															return n, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:190
				// _ = "end of CoverTab[95820]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:191
				_go_fuzz_dep_.CoverTab[95821]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:191
				// _ = "end of CoverTab[95821]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:191
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:191
			// _ = "end of CoverTab[95818]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:191
			_go_fuzz_dep_.CoverTab[95819]++
														n += zn
														buf = buf[zn:]
														continue
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:194
			// _ = "end of CoverTab[95819]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:195
			_go_fuzz_dep_.CoverTab[95822]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:195
			// _ = "end of CoverTab[95822]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:195
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:195
		// _ = "end of CoverTab[95812]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:195
		_go_fuzz_dep_.CoverTab[95813]++

													m := copy(z.data[z.idx:], buf)
													n += m
													z.idx += m
													buf = buf[m:]
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:201
			_go_fuzz_dep_.CoverTab[95823]++
														debug("%d bytes copied to buf, current index %d", n, z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:202
			// _ = "end of CoverTab[95823]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:203
			_go_fuzz_dep_.CoverTab[95824]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:203
			// _ = "end of CoverTab[95824]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:203
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:203
		// _ = "end of CoverTab[95813]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:203
		_go_fuzz_dep_.CoverTab[95814]++

													if z.idx < len(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:205
			_go_fuzz_dep_.CoverTab[95825]++

														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:207
				_go_fuzz_dep_.CoverTab[95827]++
															debug("need more data for compression")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:208
				// _ = "end of CoverTab[95827]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:209
				_go_fuzz_dep_.CoverTab[95828]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:209
				// _ = "end of CoverTab[95828]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:209
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:209
			// _ = "end of CoverTab[95825]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:209
			_go_fuzz_dep_.CoverTab[95826]++
														return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:210
			// _ = "end of CoverTab[95826]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:211
			_go_fuzz_dep_.CoverTab[95829]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:211
			// _ = "end of CoverTab[95829]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:211
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:211
		// _ = "end of CoverTab[95814]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:211
		_go_fuzz_dep_.CoverTab[95815]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:214
		if err := z.compressBlock(z.data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:214
			_go_fuzz_dep_.CoverTab[95830]++
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:215
			// _ = "end of CoverTab[95830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:216
			_go_fuzz_dep_.CoverTab[95831]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:216
			// _ = "end of CoverTab[95831]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:216
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:216
		// _ = "end of CoverTab[95815]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:216
		_go_fuzz_dep_.CoverTab[95816]++
													z.idx = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:217
		// _ = "end of CoverTab[95816]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:218
	// _ = "end of CoverTab[95804]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:218
	_go_fuzz_dep_.CoverTab[95805]++

												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:220
	// _ = "end of CoverTab[95805]"
}

// compressBlock compresses a block.
func (z *Writer) compressBlock(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:224
	_go_fuzz_dep_.CoverTab[95832]++
												if !z.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:225
		_go_fuzz_dep_.CoverTab[95844]++
													_, _ = z.checksum.Write(data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:226
		// _ = "end of CoverTab[95844]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:227
		_go_fuzz_dep_.CoverTab[95845]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:227
		// _ = "end of CoverTab[95845]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:227
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:227
	// _ = "end of CoverTab[95832]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:227
	_go_fuzz_dep_.CoverTab[95833]++

												if z.c != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:229
		_go_fuzz_dep_.CoverTab[95846]++
													c := make(chan zResult)
													z.c <- c

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:234
		block := getBuffer(z.Header.BlockMaxSize)[:len(data)]
													copy(block, data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:235
		_curRoutineNum109_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:235
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum109_)

													go func() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
			_go_fuzz_dep_.CoverTab[95847]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
				_go_fuzz_dep_.CoverTab[95848]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum109_)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
				// _ = "end of CoverTab[95848]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
			}()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
			writerCompressBlock(c, z.Header, block)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
			// _ = "end of CoverTab[95847]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:237
		}()
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:238
		// _ = "end of CoverTab[95846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:239
		_go_fuzz_dep_.CoverTab[95849]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:239
		// _ = "end of CoverTab[95849]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:239
	// _ = "end of CoverTab[95833]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:239
	_go_fuzz_dep_.CoverTab[95834]++

												zdata := z.data[z.Header.BlockMaxSize:cap(z.data)]
	// The compressed block size cannot exceed the input's.
	var zn int

	if level := z.Header.CompressionLevel; level != 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:245
		_go_fuzz_dep_.CoverTab[95850]++
													zn, _ = CompressBlockHC(data, zdata, level)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:246
		// _ = "end of CoverTab[95850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:247
		_go_fuzz_dep_.CoverTab[95851]++
													zn, _ = CompressBlock(data, zdata, z.hashtable[:])
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:248
		// _ = "end of CoverTab[95851]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:249
	// _ = "end of CoverTab[95834]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:249
	_go_fuzz_dep_.CoverTab[95835]++

												var bLen uint32
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:252
		_go_fuzz_dep_.CoverTab[95852]++
													debug("block compression %d => %d", len(data), zn)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:253
		// _ = "end of CoverTab[95852]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:254
		_go_fuzz_dep_.CoverTab[95853]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:254
		// _ = "end of CoverTab[95853]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:254
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:254
	// _ = "end of CoverTab[95835]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:254
	_go_fuzz_dep_.CoverTab[95836]++
												if zn > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:255
		_go_fuzz_dep_.CoverTab[95854]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:255
		return zn < len(data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:255
		// _ = "end of CoverTab[95854]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:255
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:255
		_go_fuzz_dep_.CoverTab[95855]++

													bLen = uint32(zn)
													zdata = zdata[:zn]
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:258
		// _ = "end of CoverTab[95855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:259
		_go_fuzz_dep_.CoverTab[95856]++

													bLen = uint32(len(data)) | compressedBlockFlag
													zdata = data
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:262
		// _ = "end of CoverTab[95856]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:263
	// _ = "end of CoverTab[95836]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:263
	_go_fuzz_dep_.CoverTab[95837]++
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:264
		_go_fuzz_dep_.CoverTab[95857]++
													debug("block compression to be written len=%d data len=%d", bLen, len(zdata))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:265
		// _ = "end of CoverTab[95857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:266
		_go_fuzz_dep_.CoverTab[95858]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:266
		// _ = "end of CoverTab[95858]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:266
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:266
	// _ = "end of CoverTab[95837]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:266
	_go_fuzz_dep_.CoverTab[95838]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:269
	if err := z.writeUint32(bLen); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:269
		_go_fuzz_dep_.CoverTab[95859]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:270
		// _ = "end of CoverTab[95859]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:271
		_go_fuzz_dep_.CoverTab[95860]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:271
		// _ = "end of CoverTab[95860]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:271
	// _ = "end of CoverTab[95838]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:271
	_go_fuzz_dep_.CoverTab[95839]++
												written, err := z.dst.Write(zdata)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:273
		_go_fuzz_dep_.CoverTab[95861]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:274
		// _ = "end of CoverTab[95861]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:275
		_go_fuzz_dep_.CoverTab[95862]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:275
		// _ = "end of CoverTab[95862]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:275
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:275
	// _ = "end of CoverTab[95839]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:275
	_go_fuzz_dep_.CoverTab[95840]++
												if h := z.OnBlockDone; h != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:276
		_go_fuzz_dep_.CoverTab[95863]++
													h(written)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:277
		// _ = "end of CoverTab[95863]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:278
		_go_fuzz_dep_.CoverTab[95864]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:278
		// _ = "end of CoverTab[95864]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:278
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:278
	// _ = "end of CoverTab[95840]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:278
	_go_fuzz_dep_.CoverTab[95841]++

												if !z.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:280
		_go_fuzz_dep_.CoverTab[95865]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:281
			_go_fuzz_dep_.CoverTab[95867]++
														debug("current frame checksum %x", z.checksum.Sum32())
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:282
			// _ = "end of CoverTab[95867]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:283
			_go_fuzz_dep_.CoverTab[95868]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:283
			// _ = "end of CoverTab[95868]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:283
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:283
		// _ = "end of CoverTab[95865]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:283
		_go_fuzz_dep_.CoverTab[95866]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:284
		// _ = "end of CoverTab[95866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:285
		_go_fuzz_dep_.CoverTab[95869]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:285
		// _ = "end of CoverTab[95869]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:285
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:285
	// _ = "end of CoverTab[95841]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:285
	_go_fuzz_dep_.CoverTab[95842]++
												checksum := xxh32.ChecksumZero(zdata)
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:287
		_go_fuzz_dep_.CoverTab[95870]++
													debug("block checksum %x", checksum)
													defer func() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:289
			_go_fuzz_dep_.CoverTab[95871]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:289
			debug("current frame checksum %x", z.checksum.Sum32())
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:289
			// _ = "end of CoverTab[95871]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:289
		}()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:289
		// _ = "end of CoverTab[95870]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:290
		_go_fuzz_dep_.CoverTab[95872]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:290
		// _ = "end of CoverTab[95872]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:290
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:290
	// _ = "end of CoverTab[95842]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:290
	_go_fuzz_dep_.CoverTab[95843]++
												return z.writeUint32(checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:291
	// _ = "end of CoverTab[95843]"
}

// Flush flushes any pending compressed data to the underlying writer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:294
// Flush does not return until the data has been written.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:294
// If the underlying writer returns an error, Flush returns that error.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:297
func (z *Writer) Flush() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:297
	_go_fuzz_dep_.CoverTab[95873]++
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:298
		_go_fuzz_dep_.CoverTab[95878]++
													debug("flush with index %d", z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:299
		// _ = "end of CoverTab[95878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:300
		_go_fuzz_dep_.CoverTab[95879]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:300
		// _ = "end of CoverTab[95879]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:300
	// _ = "end of CoverTab[95873]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:300
	_go_fuzz_dep_.CoverTab[95874]++
												if z.idx == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:301
		_go_fuzz_dep_.CoverTab[95880]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:302
		// _ = "end of CoverTab[95880]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:303
		_go_fuzz_dep_.CoverTab[95881]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:303
		// _ = "end of CoverTab[95881]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:303
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:303
	// _ = "end of CoverTab[95874]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:303
	_go_fuzz_dep_.CoverTab[95875]++

												data := getBuffer(z.Header.BlockMaxSize)[:len(z.data[:z.idx])]
												copy(data, z.data[:z.idx])

												z.idx = 0
												if z.c == nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:309
		_go_fuzz_dep_.CoverTab[95882]++
													return z.compressBlock(data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:310
		// _ = "end of CoverTab[95882]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:311
		_go_fuzz_dep_.CoverTab[95883]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:311
		// _ = "end of CoverTab[95883]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:311
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:311
	// _ = "end of CoverTab[95875]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:311
	_go_fuzz_dep_.CoverTab[95876]++
												if !z.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:312
		_go_fuzz_dep_.CoverTab[95884]++
													_, _ = z.checksum.Write(data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:313
		// _ = "end of CoverTab[95884]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:314
		_go_fuzz_dep_.CoverTab[95885]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:314
		// _ = "end of CoverTab[95885]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:314
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:314
	// _ = "end of CoverTab[95876]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:314
	_go_fuzz_dep_.CoverTab[95877]++
												c := make(chan zResult)
												z.c <- c
												writerCompressBlock(c, z.Header, data)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:318
	// _ = "end of CoverTab[95877]"
}

func (z *Writer) close() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:321
	_go_fuzz_dep_.CoverTab[95886]++
												if z.c == nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:322
		_go_fuzz_dep_.CoverTab[95888]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:323
		// _ = "end of CoverTab[95888]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:324
		_go_fuzz_dep_.CoverTab[95889]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:324
		// _ = "end of CoverTab[95889]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:324
	// _ = "end of CoverTab[95886]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:324
	_go_fuzz_dep_.CoverTab[95887]++

												c := make(chan zResult)
												z.c <- c
												c <- zResult{}

												<-c

												z.c = nil
												return z.err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:333
	// _ = "end of CoverTab[95887]"
}

// Close closes the Writer, flushing any unwritten data to the underlying io.Writer, but does not close the underlying io.Writer.
func (z *Writer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:337
	_go_fuzz_dep_.CoverTab[95890]++
												if !z.Header.done {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:338
		_go_fuzz_dep_.CoverTab[95898]++
													if err := z.writeHeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:339
			_go_fuzz_dep_.CoverTab[95899]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:340
			// _ = "end of CoverTab[95899]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:341
			_go_fuzz_dep_.CoverTab[95900]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:341
			// _ = "end of CoverTab[95900]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:341
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:341
		// _ = "end of CoverTab[95898]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:342
		_go_fuzz_dep_.CoverTab[95901]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:342
		// _ = "end of CoverTab[95901]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:342
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:342
	// _ = "end of CoverTab[95890]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:342
	_go_fuzz_dep_.CoverTab[95891]++
												if err := z.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:343
		_go_fuzz_dep_.CoverTab[95902]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:344
		// _ = "end of CoverTab[95902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:345
		_go_fuzz_dep_.CoverTab[95903]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:345
		// _ = "end of CoverTab[95903]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:345
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:345
	// _ = "end of CoverTab[95891]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:345
	_go_fuzz_dep_.CoverTab[95892]++
												if err := z.close(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:346
		_go_fuzz_dep_.CoverTab[95904]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:347
		// _ = "end of CoverTab[95904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:348
		_go_fuzz_dep_.CoverTab[95905]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:348
		// _ = "end of CoverTab[95905]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:348
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:348
	// _ = "end of CoverTab[95892]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:348
	_go_fuzz_dep_.CoverTab[95893]++
												z.freeBuffers()

												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:351
		_go_fuzz_dep_.CoverTab[95906]++
													debug("writing last empty block")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:352
		// _ = "end of CoverTab[95906]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:353
		_go_fuzz_dep_.CoverTab[95907]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:353
		// _ = "end of CoverTab[95907]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:353
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:353
	// _ = "end of CoverTab[95893]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:353
	_go_fuzz_dep_.CoverTab[95894]++
												if err := z.writeUint32(0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:354
		_go_fuzz_dep_.CoverTab[95908]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:355
		// _ = "end of CoverTab[95908]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:356
		_go_fuzz_dep_.CoverTab[95909]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:356
		// _ = "end of CoverTab[95909]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:356
	// _ = "end of CoverTab[95894]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:356
	_go_fuzz_dep_.CoverTab[95895]++
												if z.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:357
		_go_fuzz_dep_.CoverTab[95910]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:358
		// _ = "end of CoverTab[95910]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:359
		_go_fuzz_dep_.CoverTab[95911]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:359
		// _ = "end of CoverTab[95911]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:359
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:359
	// _ = "end of CoverTab[95895]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:359
	_go_fuzz_dep_.CoverTab[95896]++
												checksum := z.checksum.Sum32()
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:361
		_go_fuzz_dep_.CoverTab[95912]++
													debug("stream checksum %x", checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:362
		// _ = "end of CoverTab[95912]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:363
		_go_fuzz_dep_.CoverTab[95913]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:363
		// _ = "end of CoverTab[95913]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:363
	// _ = "end of CoverTab[95896]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:363
	_go_fuzz_dep_.CoverTab[95897]++
												return z.writeUint32(checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:364
	// _ = "end of CoverTab[95897]"
}

// Reset clears the state of the Writer z such that it is equivalent to its
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:367
// initial state from NewWriter, but instead writing to w.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:367
// No access to the underlying io.Writer is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:370
func (z *Writer) Reset(w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:370
	_go_fuzz_dep_.CoverTab[95914]++
												n := cap(z.c)
												_ = z.close()
												z.freeBuffers()
												z.Header.Reset()
												z.dst = w
												z.checksum.Reset()
												z.idx = 0
												z.err = nil

												for i := range z.hashtable {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:380
		_go_fuzz_dep_.CoverTab[95916]++
													z.hashtable[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:381
		// _ = "end of CoverTab[95916]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:382
	// _ = "end of CoverTab[95914]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:382
	_go_fuzz_dep_.CoverTab[95915]++
												z.WithConcurrency(n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:383
	// _ = "end of CoverTab[95915]"
}

// writeUint32 writes a uint32 to the underlying writer.
func (z *Writer) writeUint32(x uint32) error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:387
	_go_fuzz_dep_.CoverTab[95917]++
												buf := z.buf[:4]
												binary.LittleEndian.PutUint32(buf, x)
												_, err := z.dst.Write(buf)
												return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:391
	// _ = "end of CoverTab[95917]"
}

// writerCompressBlock compresses data into a pooled buffer and writes its result
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:394
// out to the input channel.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:396
func writerCompressBlock(c chan zResult, header Header, data []byte) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:396
	_go_fuzz_dep_.CoverTab[95918]++
												zdata := getBuffer(header.BlockMaxSize)
	// The compressed block size cannot exceed the input's.
	var zn int
	if level := header.CompressionLevel; level != 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:400
		_go_fuzz_dep_.CoverTab[95922]++
													zn, _ = CompressBlockHC(data, zdata, level)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:401
		// _ = "end of CoverTab[95922]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:402
		_go_fuzz_dep_.CoverTab[95923]++
													var hashTable [winSize]int
													zn, _ = CompressBlock(data, zdata, hashTable[:])
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:404
		// _ = "end of CoverTab[95923]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:405
	// _ = "end of CoverTab[95918]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:405
	_go_fuzz_dep_.CoverTab[95919]++
												var res zResult
												if zn > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:407
		_go_fuzz_dep_.CoverTab[95924]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:407
		return zn < len(data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:407
		// _ = "end of CoverTab[95924]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:407
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:407
		_go_fuzz_dep_.CoverTab[95925]++
													res.size = uint32(zn)
													res.data = zdata[:zn]

													putBuffer(header.BlockMaxSize, data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:411
		// _ = "end of CoverTab[95925]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:412
		_go_fuzz_dep_.CoverTab[95926]++
													res.size = uint32(len(data)) | compressedBlockFlag
													res.data = data

													putBuffer(header.BlockMaxSize, zdata)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:416
		// _ = "end of CoverTab[95926]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:417
	// _ = "end of CoverTab[95919]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:417
	_go_fuzz_dep_.CoverTab[95920]++
												if header.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:418
		_go_fuzz_dep_.CoverTab[95927]++
													res.checksum = xxh32.ChecksumZero(res.data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:419
		// _ = "end of CoverTab[95927]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:420
		_go_fuzz_dep_.CoverTab[95928]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:420
		// _ = "end of CoverTab[95928]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:420
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:420
	// _ = "end of CoverTab[95920]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:420
	_go_fuzz_dep_.CoverTab[95921]++
												c <- res
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:421
	// _ = "end of CoverTab[95921]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:422
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer.go:422
var _ = _go_fuzz_dep_.CoverTab
