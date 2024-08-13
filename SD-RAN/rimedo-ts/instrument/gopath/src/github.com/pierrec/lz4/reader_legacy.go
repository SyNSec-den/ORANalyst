//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
package lz4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:1
)

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ReaderLegacy implements the LZ4Demo frame decoder.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:9
// The Header is set after the first call to Read().
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:11
type ReaderLegacy struct {
	Header
	// Handler called when a block has been successfully read.
	// It provides the number of bytes read.
	OnBlockDone	func(size int)

	lastBlock	bool
	buf		[8]byte		// Scrap buffer.
	pos		int64		// Current position in src.
	src		io.Reader	// Source.
	zdata		[]byte		// Compressed data.
	data		[]byte		// Uncompressed data.
	idx		int		// Index of unread bytes into data.
	skip		int64		// Bytes to skip before next read.
	dpos		int64		// Position in dest
}

// NewReaderLegacy returns a new LZ4Demo frame decoder.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:28
// No access to the underlying io.Reader is performed.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:30
func NewReaderLegacy(src io.Reader) *ReaderLegacy {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:30
	_go_fuzz_dep_.CoverTab[95663]++
													r := &ReaderLegacy{src: src}
													return r
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:32
	// _ = "end of CoverTab[95663]"
}

// readHeader checks the frame magic number and parses the frame descriptoz.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:35
// Skippable frames are supported even as a first frame although the LZ4
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:35
// specifications recommends skippable frames not to be used as first frames.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:38
func (z *ReaderLegacy) readLegacyHeader() error {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:38
	_go_fuzz_dep_.CoverTab[95664]++
													z.lastBlock = false
													magic, err := z.readUint32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:41
		_go_fuzz_dep_.CoverTab[95670]++
														z.pos += 4
														if err == io.ErrUnexpectedEOF {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:43
			_go_fuzz_dep_.CoverTab[95672]++
															return io.EOF
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:44
			// _ = "end of CoverTab[95672]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:45
			_go_fuzz_dep_.CoverTab[95673]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:45
			// _ = "end of CoverTab[95673]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:45
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:45
		// _ = "end of CoverTab[95670]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:45
		_go_fuzz_dep_.CoverTab[95671]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:46
		// _ = "end of CoverTab[95671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:47
		_go_fuzz_dep_.CoverTab[95674]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:47
		// _ = "end of CoverTab[95674]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:47
	// _ = "end of CoverTab[95664]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:47
	_go_fuzz_dep_.CoverTab[95665]++
													if magic != frameMagicLegacy {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:48
		_go_fuzz_dep_.CoverTab[95675]++
														return ErrInvalid
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:49
		// _ = "end of CoverTab[95675]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:50
		_go_fuzz_dep_.CoverTab[95676]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:50
		// _ = "end of CoverTab[95676]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:50
	// _ = "end of CoverTab[95665]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:50
	_go_fuzz_dep_.CoverTab[95666]++
													z.pos += 4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:55
	bSize := blockSize4M * 2

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:59
	if n := 2 * bSize; cap(z.zdata) < n {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:59
		_go_fuzz_dep_.CoverTab[95677]++
														z.zdata = make([]byte, n, n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:60
		// _ = "end of CoverTab[95677]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:61
		_go_fuzz_dep_.CoverTab[95678]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:61
		// _ = "end of CoverTab[95678]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:61
	// _ = "end of CoverTab[95666]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:61
	_go_fuzz_dep_.CoverTab[95667]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:62
		_go_fuzz_dep_.CoverTab[95679]++
														debug("header block max size size=%d", bSize)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:63
		// _ = "end of CoverTab[95679]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:64
		_go_fuzz_dep_.CoverTab[95680]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:64
		// _ = "end of CoverTab[95680]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:64
	// _ = "end of CoverTab[95667]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:64
	_go_fuzz_dep_.CoverTab[95668]++
													z.zdata = z.zdata[:bSize]
													z.data = z.zdata[:cap(z.zdata)][bSize:]
													z.idx = len(z.data)

													z.Header.done = true
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:70
		_go_fuzz_dep_.CoverTab[95681]++
														debug("header read: %v", z.Header)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:71
		// _ = "end of CoverTab[95681]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:72
		_go_fuzz_dep_.CoverTab[95682]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:72
		// _ = "end of CoverTab[95682]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:72
	// _ = "end of CoverTab[95668]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:72
	_go_fuzz_dep_.CoverTab[95669]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:74
	// _ = "end of CoverTab[95669]"
}

// Read decompresses data from the underlying source into the supplied buffer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:77
//
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:77
// Since there can be multiple streams concatenated, Header values may
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:77
// change between calls to Read(). If that is the case, no data is actually read from
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:77
// the underlying io.Reader, to allow for potential input buffer resizing.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:82
func (z *ReaderLegacy) Read(buf []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:82
	_go_fuzz_dep_.CoverTab[95683]++
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:83
		_go_fuzz_dep_.CoverTab[95691]++
														debug("Read buf len=%d", len(buf))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:84
		// _ = "end of CoverTab[95691]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:85
		_go_fuzz_dep_.CoverTab[95692]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:85
		// _ = "end of CoverTab[95692]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:85
	// _ = "end of CoverTab[95683]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:85
	_go_fuzz_dep_.CoverTab[95684]++
													if !z.Header.done {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:86
		_go_fuzz_dep_.CoverTab[95693]++
														if err := z.readLegacyHeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:87
			_go_fuzz_dep_.CoverTab[95695]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:88
			// _ = "end of CoverTab[95695]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:89
			_go_fuzz_dep_.CoverTab[95696]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:89
			// _ = "end of CoverTab[95696]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:89
		// _ = "end of CoverTab[95693]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:89
		_go_fuzz_dep_.CoverTab[95694]++
														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:90
			_go_fuzz_dep_.CoverTab[95697]++
															debug("header read OK compressed buffer %d / %d uncompressed buffer %d : %d index=%d",
				len(z.zdata), cap(z.zdata), len(z.data), cap(z.data), z.idx)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:92
			// _ = "end of CoverTab[95697]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:93
			_go_fuzz_dep_.CoverTab[95698]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:93
			// _ = "end of CoverTab[95698]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:93
		// _ = "end of CoverTab[95694]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:94
		_go_fuzz_dep_.CoverTab[95699]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:94
		// _ = "end of CoverTab[95699]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:94
	// _ = "end of CoverTab[95684]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:94
	_go_fuzz_dep_.CoverTab[95685]++

													if len(buf) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:96
		_go_fuzz_dep_.CoverTab[95700]++
														return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:97
		// _ = "end of CoverTab[95700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:98
		_go_fuzz_dep_.CoverTab[95701]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:98
		// _ = "end of CoverTab[95701]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:98
	// _ = "end of CoverTab[95685]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:98
	_go_fuzz_dep_.CoverTab[95686]++

													if z.idx == len(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:100
		_go_fuzz_dep_.CoverTab[95702]++

														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:102
			_go_fuzz_dep_.CoverTab[95711]++
															debug("  reading block from writer %d %d", z.idx, blockSize4M*2)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:103
			// _ = "end of CoverTab[95711]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:104
			_go_fuzz_dep_.CoverTab[95712]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:104
			// _ = "end of CoverTab[95712]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:104
		// _ = "end of CoverTab[95702]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:104
		_go_fuzz_dep_.CoverTab[95703]++

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:107
		z.data = z.zdata[:cap(z.zdata)][len(z.zdata):]

		bLen, err := z.readUint32()
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:110
			_go_fuzz_dep_.CoverTab[95713]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:111
			// _ = "end of CoverTab[95713]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:112
			_go_fuzz_dep_.CoverTab[95714]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:112
			// _ = "end of CoverTab[95714]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:112
		// _ = "end of CoverTab[95703]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:112
		_go_fuzz_dep_.CoverTab[95704]++
														if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:113
			_go_fuzz_dep_.CoverTab[95715]++
															debug("   bLen %d (0x%x) offset = %d (0x%x)", bLen, bLen, z.pos, z.pos)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:114
			// _ = "end of CoverTab[95715]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:115
			_go_fuzz_dep_.CoverTab[95716]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:115
			// _ = "end of CoverTab[95716]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:115
		// _ = "end of CoverTab[95704]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:115
		_go_fuzz_dep_.CoverTab[95705]++
														z.pos += 4

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:119
		if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:119
			_go_fuzz_dep_.CoverTab[95717]++
															debug("   compressed block size %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:120
			// _ = "end of CoverTab[95717]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:121
			_go_fuzz_dep_.CoverTab[95718]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:121
			// _ = "end of CoverTab[95718]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:121
		// _ = "end of CoverTab[95705]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:121
		_go_fuzz_dep_.CoverTab[95706]++

														if int(bLen) > cap(z.data) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:123
			_go_fuzz_dep_.CoverTab[95719]++
															return 0, fmt.Errorf("lz4: invalid block size: %d", bLen)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:124
			// _ = "end of CoverTab[95719]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:125
			_go_fuzz_dep_.CoverTab[95720]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:125
			// _ = "end of CoverTab[95720]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:125
		// _ = "end of CoverTab[95706]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:125
		_go_fuzz_dep_.CoverTab[95707]++
														zdata := z.zdata[:bLen]
														if _, err := io.ReadFull(z.src, zdata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:127
			_go_fuzz_dep_.CoverTab[95721]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:128
			// _ = "end of CoverTab[95721]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:129
			_go_fuzz_dep_.CoverTab[95722]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:129
			// _ = "end of CoverTab[95722]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:129
		// _ = "end of CoverTab[95707]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:129
		_go_fuzz_dep_.CoverTab[95708]++
														z.pos += int64(bLen)

														n, err := UncompressBlock(zdata, z.data)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:133
			_go_fuzz_dep_.CoverTab[95723]++
															return 0, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:134
			// _ = "end of CoverTab[95723]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:135
			_go_fuzz_dep_.CoverTab[95724]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:135
			// _ = "end of CoverTab[95724]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:135
		// _ = "end of CoverTab[95708]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:135
		_go_fuzz_dep_.CoverTab[95709]++

														z.data = z.data[:n]
														if z.OnBlockDone != nil {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:138
			_go_fuzz_dep_.CoverTab[95725]++
															z.OnBlockDone(n)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:139
			// _ = "end of CoverTab[95725]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:140
			_go_fuzz_dep_.CoverTab[95726]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:140
			// _ = "end of CoverTab[95726]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:140
		// _ = "end of CoverTab[95709]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:140
		_go_fuzz_dep_.CoverTab[95710]++

														z.idx = 0

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:146
		if n < blockSize4M*2 {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:146
			_go_fuzz_dep_.CoverTab[95727]++
															z.lastBlock = true
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:147
			// _ = "end of CoverTab[95727]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:148
			_go_fuzz_dep_.CoverTab[95728]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:148
			// _ = "end of CoverTab[95728]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:148
		}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:148
		// _ = "end of CoverTab[95710]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:149
		_go_fuzz_dep_.CoverTab[95729]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:149
		// _ = "end of CoverTab[95729]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:149
	// _ = "end of CoverTab[95686]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:149
	_go_fuzz_dep_.CoverTab[95687]++

													if z.skip > int64(len(z.data[z.idx:])) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:151
		_go_fuzz_dep_.CoverTab[95730]++
														z.skip -= int64(len(z.data[z.idx:]))
														z.dpos += int64(len(z.data[z.idx:]))
														z.idx = len(z.data)
														return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:155
		// _ = "end of CoverTab[95730]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:156
		_go_fuzz_dep_.CoverTab[95731]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:156
		// _ = "end of CoverTab[95731]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:156
	// _ = "end of CoverTab[95687]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:156
	_go_fuzz_dep_.CoverTab[95688]++

													z.idx += int(z.skip)
													z.dpos += z.skip
													z.skip = 0

													n := copy(buf, z.data[z.idx:])
													z.idx += n
													z.dpos += int64(n)
													if debugFlag {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:165
		_go_fuzz_dep_.CoverTab[95732]++
														debug("%v] copied %d bytes to input (%d:%d)", z.lastBlock, n, z.idx, len(z.data))
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:166
		// _ = "end of CoverTab[95732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:167
		_go_fuzz_dep_.CoverTab[95733]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:167
		// _ = "end of CoverTab[95733]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:167
	// _ = "end of CoverTab[95688]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:167
	_go_fuzz_dep_.CoverTab[95689]++
													if z.lastBlock && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:168
		_go_fuzz_dep_.CoverTab[95734]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:168
		return len(z.data) == z.idx
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:168
		// _ = "end of CoverTab[95734]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:168
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:168
		_go_fuzz_dep_.CoverTab[95735]++
														return n, io.EOF
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:169
		// _ = "end of CoverTab[95735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:170
		_go_fuzz_dep_.CoverTab[95736]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:170
		// _ = "end of CoverTab[95736]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:170
	// _ = "end of CoverTab[95689]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:170
	_go_fuzz_dep_.CoverTab[95690]++
													return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:171
	// _ = "end of CoverTab[95690]"
}

// Seek implements io.Seeker, but supports seeking forward from the current
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:174
// position only. Any other seek will return an error. Allows skipping output
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:174
// bytes which aren't needed, which in some scenarios is faster than reading
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:174
// and discarding them.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:174
// Note this may cause future calls to Read() to read 0 bytes if all of the
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:174
// data they would have returned is skipped.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:180
func (z *ReaderLegacy) Seek(offset int64, whence int) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:180
	_go_fuzz_dep_.CoverTab[95737]++
													if offset < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:181
		_go_fuzz_dep_.CoverTab[95739]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:181
		return whence != io.SeekCurrent
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:181
		// _ = "end of CoverTab[95739]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:181
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:181
		_go_fuzz_dep_.CoverTab[95740]++
														return z.dpos + z.skip, ErrUnsupportedSeek
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:182
		// _ = "end of CoverTab[95740]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:183
		_go_fuzz_dep_.CoverTab[95741]++
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:183
		// _ = "end of CoverTab[95741]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:183
	// _ = "end of CoverTab[95737]"
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:183
	_go_fuzz_dep_.CoverTab[95738]++
													z.skip += offset
													return z.dpos + z.skip, nil
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:185
	// _ = "end of CoverTab[95738]"
}

// Reset discards the Reader's state and makes it equivalent to the
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:188
// result of its original state from NewReader, but reading from r instead.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:188
// This permits reusing a Reader rather than allocating a new one.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:191
func (z *ReaderLegacy) Reset(r io.Reader) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:191
	_go_fuzz_dep_.CoverTab[95742]++
													z.Header = Header{}
													z.pos = 0
													z.src = r
													z.zdata = z.zdata[:0]
													z.data = z.data[:0]
													z.idx = 0
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:197
	// _ = "end of CoverTab[95742]"
}

// readUint32 reads an uint32 into the supplied buffer.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:200
// The idea is to make use of the already allocated buffers avoiding additional allocations.
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:202
func (z *ReaderLegacy) readUint32() (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:202
	_go_fuzz_dep_.CoverTab[95743]++
													buf := z.buf[:4]
													_, err := io.ReadFull(z.src, buf)
													x := binary.LittleEndian.Uint32(buf)
													return x, err
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:206
	// _ = "end of CoverTab[95743]"
}

//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:207
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pierrec/lz4@v2.6.1+incompatible/reader_legacy.go:207
var _ = _go_fuzz_dep_.CoverTab
