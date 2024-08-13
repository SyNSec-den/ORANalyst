//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:1
)

import (
	"encoding/binary"
	"io"
)

// WriterLegacy implements the LZ4Demo frame decoder.
type WriterLegacy struct {
	Header
	// Handler called when a block has been successfully read.
	// It provides the number of bytes read.
	OnBlockDone	func(size int)

	dst		io.Writer	// Destination.
	data		[]byte		// Data to be compressed + buffer for compressed data.
	idx		int		// Index into data.
	hashtable	[winSize]int	// Hash table used in CompressBlock().
}

// NewWriterLegacy returns a new LZ4 encoder for the legacy frame format.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:21
// No access to the underlying io.Writer is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:21
// The supplied Header is checked at the first Write.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:21
// It is ok to change it before the first Write but then not until a Reset() is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:25
func NewWriterLegacy(dst io.Writer) *WriterLegacy {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:25
	_go_fuzz_dep_.CoverTab[95929]++
													z := new(WriterLegacy)
													z.Reset(dst)
													return z
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:28
	// _ = "end of CoverTab[95929]"
}

// Write compresses data from the supplied buffer into the underlying io.Writer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:31
// Write does not return until the data has been written.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:33
func (z *WriterLegacy) Write(buf []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:33
	_go_fuzz_dep_.CoverTab[95930]++
													if !z.Header.done {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:34
		_go_fuzz_dep_.CoverTab[95934]++
														if err := z.writeHeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:35
			_go_fuzz_dep_.CoverTab[95935]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:36
			// _ = "end of CoverTab[95935]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:37
			_go_fuzz_dep_.CoverTab[95936]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:37
			// _ = "end of CoverTab[95936]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:37
		// _ = "end of CoverTab[95934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:38
		_go_fuzz_dep_.CoverTab[95937]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:38
		// _ = "end of CoverTab[95937]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:38
	// _ = "end of CoverTab[95930]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:38
	_go_fuzz_dep_.CoverTab[95931]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:39
		_go_fuzz_dep_.CoverTab[95938]++
														debug("input buffer len=%d index=%d", len(buf), z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:40
		// _ = "end of CoverTab[95938]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:41
		_go_fuzz_dep_.CoverTab[95939]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:41
		// _ = "end of CoverTab[95939]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:41
	// _ = "end of CoverTab[95931]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:41
	_go_fuzz_dep_.CoverTab[95932]++

													zn := len(z.data)
													var n int
													for len(buf) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:45
		_go_fuzz_dep_.CoverTab[95940]++
														if z.idx == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:46
			_go_fuzz_dep_.CoverTab[95945]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:46
			return len(buf) >= zn
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:46
			// _ = "end of CoverTab[95945]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:46
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:46
			_go_fuzz_dep_.CoverTab[95946]++

															if err := z.compressBlock(buf[:zn]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:48
				_go_fuzz_dep_.CoverTab[95948]++
																return n, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:49
				// _ = "end of CoverTab[95948]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:50
				_go_fuzz_dep_.CoverTab[95949]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:50
				// _ = "end of CoverTab[95949]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:50
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:50
			// _ = "end of CoverTab[95946]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:50
			_go_fuzz_dep_.CoverTab[95947]++
															n += zn
															buf = buf[zn:]
															continue
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:53
			// _ = "end of CoverTab[95947]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:54
			_go_fuzz_dep_.CoverTab[95950]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:54
			// _ = "end of CoverTab[95950]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:54
		// _ = "end of CoverTab[95940]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:54
		_go_fuzz_dep_.CoverTab[95941]++

														m := copy(z.data[z.idx:], buf)
														n += m
														z.idx += m
														buf = buf[m:]
														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:60
			_go_fuzz_dep_.CoverTab[95951]++
															debug("%d bytes copied to buf, current index %d", n, z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:61
			// _ = "end of CoverTab[95951]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:62
			_go_fuzz_dep_.CoverTab[95952]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:62
			// _ = "end of CoverTab[95952]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:62
		// _ = "end of CoverTab[95941]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:62
		_go_fuzz_dep_.CoverTab[95942]++

														if z.idx < len(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:64
			_go_fuzz_dep_.CoverTab[95953]++

															if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:66
				_go_fuzz_dep_.CoverTab[95955]++
																debug("need more data for compression")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:67
				// _ = "end of CoverTab[95955]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:68
				_go_fuzz_dep_.CoverTab[95956]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:68
				// _ = "end of CoverTab[95956]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:68
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:68
			// _ = "end of CoverTab[95953]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:68
			_go_fuzz_dep_.CoverTab[95954]++
															return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:69
			// _ = "end of CoverTab[95954]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:70
			_go_fuzz_dep_.CoverTab[95957]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:70
			// _ = "end of CoverTab[95957]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:70
		// _ = "end of CoverTab[95942]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:70
		_go_fuzz_dep_.CoverTab[95943]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:73
		if err := z.compressBlock(z.data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:73
			_go_fuzz_dep_.CoverTab[95958]++
															return n, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:74
			// _ = "end of CoverTab[95958]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:75
			_go_fuzz_dep_.CoverTab[95959]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:75
			// _ = "end of CoverTab[95959]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:75
		// _ = "end of CoverTab[95943]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:75
		_go_fuzz_dep_.CoverTab[95944]++
														z.idx = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:76
		// _ = "end of CoverTab[95944]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:77
	// _ = "end of CoverTab[95932]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:77
	_go_fuzz_dep_.CoverTab[95933]++

													return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:79
	// _ = "end of CoverTab[95933]"
}

// writeHeader builds and writes the header to the underlying io.Writer.
func (z *WriterLegacy) writeHeader() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:83
	_go_fuzz_dep_.CoverTab[95960]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:86
	bSize := 2 * blockSize4M

													buf := make([]byte, 2*bSize, 2*bSize)
													z.data = buf[:bSize]

													z.idx = 0

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:94
	if err := binary.Write(z.dst, binary.LittleEndian, frameMagicLegacy); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:94
		_go_fuzz_dep_.CoverTab[95963]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:95
		// _ = "end of CoverTab[95963]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:96
		_go_fuzz_dep_.CoverTab[95964]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:96
		// _ = "end of CoverTab[95964]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:96
	// _ = "end of CoverTab[95960]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:96
	_go_fuzz_dep_.CoverTab[95961]++
													z.Header.done = true
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:98
		_go_fuzz_dep_.CoverTab[95965]++
														debug("wrote header %v", z.Header)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:99
		// _ = "end of CoverTab[95965]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:100
		_go_fuzz_dep_.CoverTab[95966]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:100
		// _ = "end of CoverTab[95966]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:100
	// _ = "end of CoverTab[95961]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:100
	_go_fuzz_dep_.CoverTab[95962]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:102
	// _ = "end of CoverTab[95962]"
}

// compressBlock compresses a block.
func (z *WriterLegacy) compressBlock(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:106
	_go_fuzz_dep_.CoverTab[95967]++
													bSize := 2 * blockSize4M
													zdata := z.data[bSize:cap(z.data)]
	// The compressed block size cannot exceed the input's.
	var zn int

	if level := z.Header.CompressionLevel; level != 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:112
		_go_fuzz_dep_.CoverTab[95973]++
														zn, _ = CompressBlockHC(data, zdata, level)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:113
		// _ = "end of CoverTab[95973]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:114
		_go_fuzz_dep_.CoverTab[95974]++
														zn, _ = CompressBlock(data, zdata, z.hashtable[:])
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:115
		// _ = "end of CoverTab[95974]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:116
	// _ = "end of CoverTab[95967]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:116
	_go_fuzz_dep_.CoverTab[95968]++

													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:118
		_go_fuzz_dep_.CoverTab[95975]++
														debug("block compression %d => %d", len(data), zn)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:119
		// _ = "end of CoverTab[95975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:120
		_go_fuzz_dep_.CoverTab[95976]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:120
		// _ = "end of CoverTab[95976]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:120
	// _ = "end of CoverTab[95968]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:120
	_go_fuzz_dep_.CoverTab[95969]++
													zdata = zdata[:zn]

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:124
	if err := binary.Write(z.dst, binary.LittleEndian, uint32(zn)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:124
		_go_fuzz_dep_.CoverTab[95977]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:125
		// _ = "end of CoverTab[95977]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:126
		_go_fuzz_dep_.CoverTab[95978]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:126
		// _ = "end of CoverTab[95978]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:126
	// _ = "end of CoverTab[95969]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:126
	_go_fuzz_dep_.CoverTab[95970]++
													written, err := z.dst.Write(zdata)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:128
		_go_fuzz_dep_.CoverTab[95979]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:129
		// _ = "end of CoverTab[95979]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:130
		_go_fuzz_dep_.CoverTab[95980]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:130
		// _ = "end of CoverTab[95980]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:130
	// _ = "end of CoverTab[95970]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:130
	_go_fuzz_dep_.CoverTab[95971]++
													if h := z.OnBlockDone; h != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:131
		_go_fuzz_dep_.CoverTab[95981]++
														h(written)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:132
		// _ = "end of CoverTab[95981]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:133
		_go_fuzz_dep_.CoverTab[95982]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:133
		// _ = "end of CoverTab[95982]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:133
	// _ = "end of CoverTab[95971]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:133
	_go_fuzz_dep_.CoverTab[95972]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:134
	// _ = "end of CoverTab[95972]"
}

// Flush flushes any pending compressed data to the underlying writer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:137
// Flush does not return until the data has been written.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:137
// If the underlying writer returns an error, Flush returns that error.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:140
func (z *WriterLegacy) Flush() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:140
	_go_fuzz_dep_.CoverTab[95983]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:141
		_go_fuzz_dep_.CoverTab[95986]++
														debug("flush with index %d", z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:142
		// _ = "end of CoverTab[95986]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:143
		_go_fuzz_dep_.CoverTab[95987]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:143
		// _ = "end of CoverTab[95987]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:143
	// _ = "end of CoverTab[95983]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:143
	_go_fuzz_dep_.CoverTab[95984]++
													if z.idx == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:144
		_go_fuzz_dep_.CoverTab[95988]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:145
		// _ = "end of CoverTab[95988]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:146
		_go_fuzz_dep_.CoverTab[95989]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:146
		// _ = "end of CoverTab[95989]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:146
	// _ = "end of CoverTab[95984]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:146
	_go_fuzz_dep_.CoverTab[95985]++

													data := z.data[:z.idx]
													z.idx = 0
													return z.compressBlock(data)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:150
	// _ = "end of CoverTab[95985]"
}

// Close closes the WriterLegacy, flushing any unwritten data to the underlying io.Writer, but does not close the underlying io.Writer.
func (z *WriterLegacy) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:154
	_go_fuzz_dep_.CoverTab[95990]++
													if !z.Header.done {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:155
		_go_fuzz_dep_.CoverTab[95994]++
														if err := z.writeHeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:156
			_go_fuzz_dep_.CoverTab[95995]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:157
			// _ = "end of CoverTab[95995]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:158
			_go_fuzz_dep_.CoverTab[95996]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:158
			// _ = "end of CoverTab[95996]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:158
		// _ = "end of CoverTab[95994]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:159
		_go_fuzz_dep_.CoverTab[95997]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:159
		// _ = "end of CoverTab[95997]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:159
	// _ = "end of CoverTab[95990]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:159
	_go_fuzz_dep_.CoverTab[95991]++
													if err := z.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:160
		_go_fuzz_dep_.CoverTab[95998]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:161
		// _ = "end of CoverTab[95998]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:162
		_go_fuzz_dep_.CoverTab[95999]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:162
		// _ = "end of CoverTab[95999]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:162
	// _ = "end of CoverTab[95991]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:162
	_go_fuzz_dep_.CoverTab[95992]++

													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:164
		_go_fuzz_dep_.CoverTab[96000]++
														debug("writing last empty block")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:165
		// _ = "end of CoverTab[96000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:166
		_go_fuzz_dep_.CoverTab[96001]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:166
		// _ = "end of CoverTab[96001]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:166
	// _ = "end of CoverTab[95992]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:166
	_go_fuzz_dep_.CoverTab[95993]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:168
	// _ = "end of CoverTab[95993]"
}

// Reset clears the state of the WriterLegacy z such that it is equivalent to its
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:171
// initial state from NewWriterLegacy, but instead writing to w.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:171
// No access to the underlying io.Writer is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:174
func (z *WriterLegacy) Reset(w io.Writer) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:174
	_go_fuzz_dep_.CoverTab[96002]++
													z.Header.Reset()
													z.dst = w
													z.idx = 0

													for i := range z.hashtable {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:179
		_go_fuzz_dep_.CoverTab[96003]++
														z.hashtable[i] = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:180
		// _ = "end of CoverTab[96003]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:181
	// _ = "end of CoverTab[96002]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:182
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/writer_legacy.go:182
var _ = _go_fuzz_dep_.CoverTab
