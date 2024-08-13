//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:1
)

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pierrec/lz4/internal/xxh32"
)

// Reader implements the LZ4 frame decoder.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:12
// The Header is set after the first call to Read().
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:12
// The Header may change between Read() calls in case of concatenated frames.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:15
type Reader struct {
	Header
	// Handler called when a block has been successfully read.
	// It provides the number of bytes read.
	OnBlockDone	func(size int)

	buf		[8]byte		// Scrap buffer.
	pos		int64		// Current position in src.
	src		io.Reader	// Source.
	zdata		[]byte		// Compressed data.
	data		[]byte		// Uncompressed data.
	idx		int		// Index of unread bytes into data.
	checksum	xxh32.XXHZero	// Frame hash.
	skip		int64		// Bytes to skip before next read.
	dpos		int64		// Position in dest
}

// NewReader returns a new LZ4 frame decoder.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:32
// No access to the underlying io.Reader is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:34
func NewReader(src io.Reader) *Reader {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:34
	_go_fuzz_dep_.CoverTab[95503]++
												r := &Reader{src: src}
												return r
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:36
	// _ = "end of CoverTab[95503]"
}

// readHeader checks the frame magic number and parses the frame descriptoz.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:39
// Skippable frames are supported even as a first frame although the LZ4
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:39
// specifications recommends skippable frames not to be used as first frames.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:42
func (z *Reader) readHeader(first bool) error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:42
	_go_fuzz_dep_.CoverTab[95504]++
												defer z.checksum.Reset()

												buf := z.buf[:]
												for {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:46
		_go_fuzz_dep_.CoverTab[95516]++
													magic, err := z.readUint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:48
			_go_fuzz_dep_.CoverTab[95522]++
														z.pos += 4
														if !first && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:50
				_go_fuzz_dep_.CoverTab[95524]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:50
				return err == io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:50
				// _ = "end of CoverTab[95524]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:50
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:50
				_go_fuzz_dep_.CoverTab[95525]++
															return io.EOF
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:51
				// _ = "end of CoverTab[95525]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:52
				_go_fuzz_dep_.CoverTab[95526]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:52
				// _ = "end of CoverTab[95526]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:52
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:52
			// _ = "end of CoverTab[95522]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:52
			_go_fuzz_dep_.CoverTab[95523]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:53
			// _ = "end of CoverTab[95523]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:54
			_go_fuzz_dep_.CoverTab[95527]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:54
			// _ = "end of CoverTab[95527]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:54
		// _ = "end of CoverTab[95516]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:54
		_go_fuzz_dep_.CoverTab[95517]++
													if magic == frameMagic {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:55
			_go_fuzz_dep_.CoverTab[95528]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:56
			// _ = "end of CoverTab[95528]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:57
			_go_fuzz_dep_.CoverTab[95529]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:57
			// _ = "end of CoverTab[95529]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:57
		// _ = "end of CoverTab[95517]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:57
		_go_fuzz_dep_.CoverTab[95518]++
													if magic>>8 != frameSkipMagic>>8 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:58
			_go_fuzz_dep_.CoverTab[95530]++
														return ErrInvalid
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:59
			// _ = "end of CoverTab[95530]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:60
			_go_fuzz_dep_.CoverTab[95531]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:60
			// _ = "end of CoverTab[95531]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:60
		// _ = "end of CoverTab[95518]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:60
		_go_fuzz_dep_.CoverTab[95519]++
													skipSize, err := z.readUint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:62
			_go_fuzz_dep_.CoverTab[95532]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:63
			// _ = "end of CoverTab[95532]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:64
			_go_fuzz_dep_.CoverTab[95533]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:64
			// _ = "end of CoverTab[95533]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:64
		// _ = "end of CoverTab[95519]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:64
		_go_fuzz_dep_.CoverTab[95520]++
													z.pos += 4
													m, err := io.CopyN(ioutil.Discard, z.src, int64(skipSize))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:67
			_go_fuzz_dep_.CoverTab[95534]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:68
			// _ = "end of CoverTab[95534]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:69
			_go_fuzz_dep_.CoverTab[95535]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:69
			// _ = "end of CoverTab[95535]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:69
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:69
		// _ = "end of CoverTab[95520]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:69
		_go_fuzz_dep_.CoverTab[95521]++
													z.pos += m
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:70
		// _ = "end of CoverTab[95521]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:71
	// _ = "end of CoverTab[95504]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:71
	_go_fuzz_dep_.CoverTab[95505]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:74
	if _, err := io.ReadFull(z.src, buf[:2]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:74
		_go_fuzz_dep_.CoverTab[95536]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:75
		// _ = "end of CoverTab[95536]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:76
		_go_fuzz_dep_.CoverTab[95537]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:76
		// _ = "end of CoverTab[95537]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:76
	// _ = "end of CoverTab[95505]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:76
	_go_fuzz_dep_.CoverTab[95506]++
												z.pos += 8

												b := buf[0]
												if v := b >> 6; v != Version {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:80
		_go_fuzz_dep_.CoverTab[95538]++
													return fmt.Errorf("lz4: invalid version: got %d; expected %d", v, Version)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:81
		// _ = "end of CoverTab[95538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:82
		_go_fuzz_dep_.CoverTab[95539]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:82
		// _ = "end of CoverTab[95539]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:82
	// _ = "end of CoverTab[95506]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:82
	_go_fuzz_dep_.CoverTab[95507]++
												if b>>5&1 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:83
		_go_fuzz_dep_.CoverTab[95540]++
													return ErrBlockDependency
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:84
		// _ = "end of CoverTab[95540]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:85
		_go_fuzz_dep_.CoverTab[95541]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:85
		// _ = "end of CoverTab[95541]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:85
	// _ = "end of CoverTab[95507]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:85
	_go_fuzz_dep_.CoverTab[95508]++
												z.BlockChecksum = b>>4&1 > 0
												frameSize := b>>3&1 > 0
												z.NoChecksum = b>>2&1 == 0

												bmsID := buf[1] >> 4 & 0x7
												if bmsID < 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:91
		_go_fuzz_dep_.CoverTab[95542]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:91
		return bmsID > 7
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:91
		// _ = "end of CoverTab[95542]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:91
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:91
		_go_fuzz_dep_.CoverTab[95543]++
													return fmt.Errorf("lz4: invalid block max size ID: %d", bmsID)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:92
		// _ = "end of CoverTab[95543]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:93
		_go_fuzz_dep_.CoverTab[95544]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:93
		// _ = "end of CoverTab[95544]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:93
	// _ = "end of CoverTab[95508]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:93
	_go_fuzz_dep_.CoverTab[95509]++
												bSize := blockSizeIndexToValue(bmsID - 4)
												z.BlockMaxSize = bSize

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:99
	if n := 2 * bSize; cap(z.zdata) < n {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:99
		_go_fuzz_dep_.CoverTab[95545]++
													z.zdata = make([]byte, n, n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:100
		// _ = "end of CoverTab[95545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:101
		_go_fuzz_dep_.CoverTab[95546]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:101
		// _ = "end of CoverTab[95546]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:101
	// _ = "end of CoverTab[95509]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:101
	_go_fuzz_dep_.CoverTab[95510]++
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:102
		_go_fuzz_dep_.CoverTab[95547]++
													debug("header block max size id=%d size=%d", bmsID, bSize)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:103
		// _ = "end of CoverTab[95547]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:104
		_go_fuzz_dep_.CoverTab[95548]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:104
		// _ = "end of CoverTab[95548]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:104
	// _ = "end of CoverTab[95510]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:104
	_go_fuzz_dep_.CoverTab[95511]++
												z.zdata = z.zdata[:bSize]
												z.data = z.zdata[:cap(z.zdata)][bSize:]
												z.idx = len(z.data)

												_, _ = z.checksum.Write(buf[0:2])

												if frameSize {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:111
		_go_fuzz_dep_.CoverTab[95549]++
													buf := buf[:8]
													if _, err := io.ReadFull(z.src, buf); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:113
			_go_fuzz_dep_.CoverTab[95551]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:114
			// _ = "end of CoverTab[95551]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:115
			_go_fuzz_dep_.CoverTab[95552]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:115
			// _ = "end of CoverTab[95552]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:115
		// _ = "end of CoverTab[95549]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:115
		_go_fuzz_dep_.CoverTab[95550]++
													z.Size = binary.LittleEndian.Uint64(buf)
													z.pos += 8
													_, _ = z.checksum.Write(buf)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:118
		// _ = "end of CoverTab[95550]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:119
		_go_fuzz_dep_.CoverTab[95553]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:119
		// _ = "end of CoverTab[95553]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:119
	// _ = "end of CoverTab[95511]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:119
	_go_fuzz_dep_.CoverTab[95512]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:122
	if _, err := io.ReadFull(z.src, buf[:1]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:122
		_go_fuzz_dep_.CoverTab[95554]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:123
		// _ = "end of CoverTab[95554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:124
		_go_fuzz_dep_.CoverTab[95555]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:124
		// _ = "end of CoverTab[95555]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:124
	// _ = "end of CoverTab[95512]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:124
	_go_fuzz_dep_.CoverTab[95513]++
												z.pos++
												if h := byte(z.checksum.Sum32() >> 8 & 0xFF); h != buf[0] {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:126
		_go_fuzz_dep_.CoverTab[95556]++
													return fmt.Errorf("lz4: invalid header checksum: got %x; expected %x", buf[0], h)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:127
		// _ = "end of CoverTab[95556]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:128
		_go_fuzz_dep_.CoverTab[95557]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:128
		// _ = "end of CoverTab[95557]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:128
	// _ = "end of CoverTab[95513]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:128
	_go_fuzz_dep_.CoverTab[95514]++

												z.Header.done = true
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:131
		_go_fuzz_dep_.CoverTab[95558]++
													debug("header read: %v", z.Header)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:132
		// _ = "end of CoverTab[95558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:133
		_go_fuzz_dep_.CoverTab[95559]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:133
		// _ = "end of CoverTab[95559]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:133
	// _ = "end of CoverTab[95514]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:133
	_go_fuzz_dep_.CoverTab[95515]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:135
	// _ = "end of CoverTab[95515]"
}

// Read decompresses data from the underlying source into the supplied buffer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:138
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:138
// Since there can be multiple streams concatenated, Header values may
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:138
// change between calls to Read(). If that is the case, no data is actually read from
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:138
// the underlying io.Reader, to allow for potential input buffer resizing.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:143
func (z *Reader) Read(buf []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:143
	_go_fuzz_dep_.CoverTab[95560]++
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:144
		_go_fuzz_dep_.CoverTab[95567]++
													debug("Read buf len=%d", len(buf))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:145
		// _ = "end of CoverTab[95567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:146
		_go_fuzz_dep_.CoverTab[95568]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:146
		// _ = "end of CoverTab[95568]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:146
	// _ = "end of CoverTab[95560]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:146
	_go_fuzz_dep_.CoverTab[95561]++
												if !z.Header.done {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:147
		_go_fuzz_dep_.CoverTab[95569]++
													if err := z.readHeader(true); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:148
			_go_fuzz_dep_.CoverTab[95571]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:149
			// _ = "end of CoverTab[95571]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:150
			_go_fuzz_dep_.CoverTab[95572]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:150
			// _ = "end of CoverTab[95572]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:150
		// _ = "end of CoverTab[95569]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:150
		_go_fuzz_dep_.CoverTab[95570]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:151
			_go_fuzz_dep_.CoverTab[95573]++
														debug("header read OK compressed buffer %d / %d uncompressed buffer %d : %d index=%d",
				len(z.zdata), cap(z.zdata), len(z.data), cap(z.data), z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:153
			// _ = "end of CoverTab[95573]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:154
			_go_fuzz_dep_.CoverTab[95574]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:154
			// _ = "end of CoverTab[95574]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:154
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:154
		// _ = "end of CoverTab[95570]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:155
		_go_fuzz_dep_.CoverTab[95575]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:155
		// _ = "end of CoverTab[95575]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:155
	// _ = "end of CoverTab[95561]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:155
	_go_fuzz_dep_.CoverTab[95562]++

												if len(buf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:157
		_go_fuzz_dep_.CoverTab[95576]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:158
		// _ = "end of CoverTab[95576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:159
		_go_fuzz_dep_.CoverTab[95577]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:159
		// _ = "end of CoverTab[95577]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:159
	// _ = "end of CoverTab[95562]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:159
	_go_fuzz_dep_.CoverTab[95563]++

												if z.idx == len(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:161
		_go_fuzz_dep_.CoverTab[95578]++

													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:163
			_go_fuzz_dep_.CoverTab[95585]++
														debug("reading block from writer")
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:164
			// _ = "end of CoverTab[95585]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:165
			_go_fuzz_dep_.CoverTab[95586]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:165
			// _ = "end of CoverTab[95586]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:165
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:165
		// _ = "end of CoverTab[95578]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:165
		_go_fuzz_dep_.CoverTab[95579]++

													z.data = z.zdata[:cap(z.zdata)][len(z.zdata):]

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:170
		bLen, err := z.readUint32()
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:171
			_go_fuzz_dep_.CoverTab[95587]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:172
			// _ = "end of CoverTab[95587]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:173
			_go_fuzz_dep_.CoverTab[95588]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:173
			// _ = "end of CoverTab[95588]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:173
		// _ = "end of CoverTab[95579]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:173
		_go_fuzz_dep_.CoverTab[95580]++
													z.pos += 4

													if bLen == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:176
			_go_fuzz_dep_.CoverTab[95589]++

														if !z.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:178
				_go_fuzz_dep_.CoverTab[95591]++

															checksum, err := z.readUint32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:181
					_go_fuzz_dep_.CoverTab[95594]++
																return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:182
					// _ = "end of CoverTab[95594]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:183
					_go_fuzz_dep_.CoverTab[95595]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:183
					// _ = "end of CoverTab[95595]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:183
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:183
				// _ = "end of CoverTab[95591]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:183
				_go_fuzz_dep_.CoverTab[95592]++
															if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:184
					_go_fuzz_dep_.CoverTab[95596]++
																debug("frame checksum got=%x / want=%x", z.checksum.Sum32(), checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:185
					// _ = "end of CoverTab[95596]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:186
					_go_fuzz_dep_.CoverTab[95597]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:186
					// _ = "end of CoverTab[95597]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:186
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:186
				// _ = "end of CoverTab[95592]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:186
				_go_fuzz_dep_.CoverTab[95593]++
															z.pos += 4
															if h := z.checksum.Sum32(); checksum != h {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:188
					_go_fuzz_dep_.CoverTab[95598]++
																return 0, fmt.Errorf("lz4: invalid frame checksum: got %x; expected %x", h, checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:189
					// _ = "end of CoverTab[95598]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:190
					_go_fuzz_dep_.CoverTab[95599]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:190
					// _ = "end of CoverTab[95599]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:190
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:190
				// _ = "end of CoverTab[95593]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:191
				_go_fuzz_dep_.CoverTab[95600]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:191
				// _ = "end of CoverTab[95600]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:191
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:191
			// _ = "end of CoverTab[95589]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:191
			_go_fuzz_dep_.CoverTab[95590]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:194
			pos := z.pos
														z.Reset(z.src)
														z.pos = pos

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:199
			return 0, z.readHeader(false)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:199
			// _ = "end of CoverTab[95590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:200
			_go_fuzz_dep_.CoverTab[95601]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:200
			// _ = "end of CoverTab[95601]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:200
		// _ = "end of CoverTab[95580]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:200
		_go_fuzz_dep_.CoverTab[95581]++

													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:202
			_go_fuzz_dep_.CoverTab[95602]++
														debug("raw block size %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:203
			// _ = "end of CoverTab[95602]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:204
			_go_fuzz_dep_.CoverTab[95603]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:204
			// _ = "end of CoverTab[95603]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:204
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:204
		// _ = "end of CoverTab[95581]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:204
		_go_fuzz_dep_.CoverTab[95582]++
													if bLen&compressedBlockFlag > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:205
			_go_fuzz_dep_.CoverTab[95604]++

														bLen &= compressedBlockMask
														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:208
				_go_fuzz_dep_.CoverTab[95609]++
															debug("uncompressed block size %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:209
				// _ = "end of CoverTab[95609]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:210
				_go_fuzz_dep_.CoverTab[95610]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:210
				// _ = "end of CoverTab[95610]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:210
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:210
			// _ = "end of CoverTab[95604]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:210
			_go_fuzz_dep_.CoverTab[95605]++
														if int(bLen) > cap(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:211
				_go_fuzz_dep_.CoverTab[95611]++
															return 0, fmt.Errorf("lz4: invalid block size: %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:212
				// _ = "end of CoverTab[95611]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:213
				_go_fuzz_dep_.CoverTab[95612]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:213
				// _ = "end of CoverTab[95612]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:213
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:213
			// _ = "end of CoverTab[95605]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:213
			_go_fuzz_dep_.CoverTab[95606]++
														z.data = z.data[:bLen]
														if _, err := io.ReadFull(z.src, z.data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:215
				_go_fuzz_dep_.CoverTab[95613]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:216
				// _ = "end of CoverTab[95613]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:217
				_go_fuzz_dep_.CoverTab[95614]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:217
				// _ = "end of CoverTab[95614]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:217
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:217
			// _ = "end of CoverTab[95606]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:217
			_go_fuzz_dep_.CoverTab[95607]++
														z.pos += int64(bLen)
														if z.OnBlockDone != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:219
				_go_fuzz_dep_.CoverTab[95615]++
															z.OnBlockDone(int(bLen))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:220
				// _ = "end of CoverTab[95615]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:221
				_go_fuzz_dep_.CoverTab[95616]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:221
				// _ = "end of CoverTab[95616]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:221
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:221
			// _ = "end of CoverTab[95607]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:221
			_go_fuzz_dep_.CoverTab[95608]++

														if z.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:223
				_go_fuzz_dep_.CoverTab[95617]++
															checksum, err := z.readUint32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:225
					_go_fuzz_dep_.CoverTab[95619]++
																return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:226
					// _ = "end of CoverTab[95619]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:227
					_go_fuzz_dep_.CoverTab[95620]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:227
					// _ = "end of CoverTab[95620]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:227
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:227
				// _ = "end of CoverTab[95617]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:227
				_go_fuzz_dep_.CoverTab[95618]++
															z.pos += 4

															if h := xxh32.ChecksumZero(z.data); h != checksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:230
					_go_fuzz_dep_.CoverTab[95621]++
																return 0, fmt.Errorf("lz4: invalid block checksum: got %x; expected %x", h, checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:231
					// _ = "end of CoverTab[95621]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:232
					_go_fuzz_dep_.CoverTab[95622]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:232
					// _ = "end of CoverTab[95622]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:232
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:232
				// _ = "end of CoverTab[95618]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:233
				_go_fuzz_dep_.CoverTab[95623]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:233
				// _ = "end of CoverTab[95623]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:233
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:233
			// _ = "end of CoverTab[95608]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:235
			_go_fuzz_dep_.CoverTab[95624]++

														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:237
				_go_fuzz_dep_.CoverTab[95630]++
															debug("compressed block size %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:238
				// _ = "end of CoverTab[95630]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:239
				_go_fuzz_dep_.CoverTab[95631]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:239
				// _ = "end of CoverTab[95631]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:239
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:239
			// _ = "end of CoverTab[95624]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:239
			_go_fuzz_dep_.CoverTab[95625]++
														if int(bLen) > cap(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:240
				_go_fuzz_dep_.CoverTab[95632]++
															return 0, fmt.Errorf("lz4: invalid block size: %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:241
				// _ = "end of CoverTab[95632]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:242
				_go_fuzz_dep_.CoverTab[95633]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:242
				// _ = "end of CoverTab[95633]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:242
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:242
			// _ = "end of CoverTab[95625]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:242
			_go_fuzz_dep_.CoverTab[95626]++
														zdata := z.zdata[:bLen]
														if _, err := io.ReadFull(z.src, zdata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:244
				_go_fuzz_dep_.CoverTab[95634]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:245
				// _ = "end of CoverTab[95634]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:246
				_go_fuzz_dep_.CoverTab[95635]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:246
				// _ = "end of CoverTab[95635]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:246
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:246
			// _ = "end of CoverTab[95626]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:246
			_go_fuzz_dep_.CoverTab[95627]++
														z.pos += int64(bLen)

														if z.BlockChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:249
				_go_fuzz_dep_.CoverTab[95636]++
															checksum, err := z.readUint32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:251
					_go_fuzz_dep_.CoverTab[95638]++
																return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:252
					// _ = "end of CoverTab[95638]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:253
					_go_fuzz_dep_.CoverTab[95639]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:253
					// _ = "end of CoverTab[95639]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:253
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:253
				// _ = "end of CoverTab[95636]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:253
				_go_fuzz_dep_.CoverTab[95637]++
															z.pos += 4

															if h := xxh32.ChecksumZero(zdata); h != checksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:256
					_go_fuzz_dep_.CoverTab[95640]++
																return 0, fmt.Errorf("lz4: invalid block checksum: got %x; expected %x", h, checksum)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:257
					// _ = "end of CoverTab[95640]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:258
					_go_fuzz_dep_.CoverTab[95641]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:258
					// _ = "end of CoverTab[95641]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:258
				}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:258
				// _ = "end of CoverTab[95637]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:259
				_go_fuzz_dep_.CoverTab[95642]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:259
				// _ = "end of CoverTab[95642]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:259
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:259
			// _ = "end of CoverTab[95627]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:259
			_go_fuzz_dep_.CoverTab[95628]++

														n, err := UncompressBlock(zdata, z.data)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:262
				_go_fuzz_dep_.CoverTab[95643]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:263
				// _ = "end of CoverTab[95643]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:264
				_go_fuzz_dep_.CoverTab[95644]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:264
				// _ = "end of CoverTab[95644]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:264
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:264
			// _ = "end of CoverTab[95628]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:264
			_go_fuzz_dep_.CoverTab[95629]++
														z.data = z.data[:n]
														if z.OnBlockDone != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:266
				_go_fuzz_dep_.CoverTab[95645]++
															z.OnBlockDone(n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:267
				// _ = "end of CoverTab[95645]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:268
				_go_fuzz_dep_.CoverTab[95646]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:268
				// _ = "end of CoverTab[95646]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:268
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:268
			// _ = "end of CoverTab[95629]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:269
		// _ = "end of CoverTab[95582]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:269
		_go_fuzz_dep_.CoverTab[95583]++

													if !z.NoChecksum {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:271
			_go_fuzz_dep_.CoverTab[95647]++
														_, _ = z.checksum.Write(z.data)
														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:273
				_go_fuzz_dep_.CoverTab[95648]++
															debug("current frame checksum %x", z.checksum.Sum32())
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:274
				// _ = "end of CoverTab[95648]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:275
				_go_fuzz_dep_.CoverTab[95649]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:275
				// _ = "end of CoverTab[95649]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:275
			}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:275
			// _ = "end of CoverTab[95647]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:276
			_go_fuzz_dep_.CoverTab[95650]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:276
			// _ = "end of CoverTab[95650]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:276
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:276
		// _ = "end of CoverTab[95583]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:276
		_go_fuzz_dep_.CoverTab[95584]++
													z.idx = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:277
		// _ = "end of CoverTab[95584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:278
		_go_fuzz_dep_.CoverTab[95651]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:278
		// _ = "end of CoverTab[95651]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:278
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:278
	// _ = "end of CoverTab[95563]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:278
	_go_fuzz_dep_.CoverTab[95564]++

												if z.skip > int64(len(z.data[z.idx:])) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:280
		_go_fuzz_dep_.CoverTab[95652]++
													z.skip -= int64(len(z.data[z.idx:]))
													z.dpos += int64(len(z.data[z.idx:]))
													z.idx = len(z.data)
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:284
		// _ = "end of CoverTab[95652]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:285
		_go_fuzz_dep_.CoverTab[95653]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:285
		// _ = "end of CoverTab[95653]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:285
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:285
	// _ = "end of CoverTab[95564]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:285
	_go_fuzz_dep_.CoverTab[95565]++

												z.idx += int(z.skip)
												z.dpos += z.skip
												z.skip = 0

												n := copy(buf, z.data[z.idx:])
												z.idx += n
												z.dpos += int64(n)
												if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:294
		_go_fuzz_dep_.CoverTab[95654]++
													debug("copied %d bytes to input", n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:295
		// _ = "end of CoverTab[95654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:296
		_go_fuzz_dep_.CoverTab[95655]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:296
		// _ = "end of CoverTab[95655]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:296
	// _ = "end of CoverTab[95565]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:296
	_go_fuzz_dep_.CoverTab[95566]++

												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:298
	// _ = "end of CoverTab[95566]"
}

// Seek implements io.Seeker, but supports seeking forward from the current
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:301
// position only. Any other seek will return an error. Allows skipping output
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:301
// bytes which aren't needed, which in some scenarios is faster than reading
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:301
// and discarding them.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:301
// Note this may cause future calls to Read() to read 0 bytes if all of the
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:301
// data they would have returned is skipped.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:307
func (z *Reader) Seek(offset int64, whence int) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:307
	_go_fuzz_dep_.CoverTab[95656]++
												if offset < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:308
		_go_fuzz_dep_.CoverTab[95658]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:308
		return whence != io.SeekCurrent
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:308
		// _ = "end of CoverTab[95658]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:308
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:308
		_go_fuzz_dep_.CoverTab[95659]++
													return z.dpos + z.skip, ErrUnsupportedSeek
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:309
		// _ = "end of CoverTab[95659]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:310
		_go_fuzz_dep_.CoverTab[95660]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:310
		// _ = "end of CoverTab[95660]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:310
	// _ = "end of CoverTab[95656]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:310
	_go_fuzz_dep_.CoverTab[95657]++
												z.skip += offset
												return z.dpos + z.skip, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:312
	// _ = "end of CoverTab[95657]"
}

// Reset discards the Reader's state and makes it equivalent to the
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:315
// result of its original state from NewReader, but reading from r instead.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:315
// This permits reusing a Reader rather than allocating a new one.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:318
func (z *Reader) Reset(r io.Reader) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:318
	_go_fuzz_dep_.CoverTab[95661]++
												z.Header = Header{}
												z.pos = 0
												z.src = r
												z.zdata = z.zdata[:0]
												z.data = z.data[:0]
												z.idx = 0
												z.checksum.Reset()
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:325
	// _ = "end of CoverTab[95661]"
}

// readUint32 reads an uint32 into the supplied buffer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:328
// The idea is to make use of the already allocated buffers avoiding additional allocations.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:330
func (z *Reader) readUint32() (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:330
	_go_fuzz_dep_.CoverTab[95662]++
												buf := z.buf[:4]
												_, err := io.ReadFull(z.src, buf)
												x := binary.LittleEndian.Uint32(buf)
												return x, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:334
	// _ = "end of CoverTab[95662]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:335
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader.go:335
var _ = _go_fuzz_dep_.CoverTab
