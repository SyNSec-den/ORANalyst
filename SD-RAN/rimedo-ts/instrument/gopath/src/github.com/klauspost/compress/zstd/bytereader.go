// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.
// Based on work by Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
package zstd

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:5
)

// byteReader provides a byte reader that reads
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:7
// little endian values from a byte stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:7
// The input stream is manually advanced.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:7
// The reader performs no bounds checks.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:11
type byteReader struct {
	b	[]byte
	off	int
}

// init will initialize the reader and set the input.
func (b *byteReader) init(in []byte) {
	b.b = in
	b.off = 0
}

// advance the stream b n bytes.
func (b *byteReader) advance(n uint) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:23
	_go_fuzz_dep_.CoverTab[91607]++
												b.off += int(n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:24
	// _ = "end of CoverTab[91607]"
}

// overread returns whether we have advanced too far.
func (b *byteReader) overread() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:28
	_go_fuzz_dep_.CoverTab[91608]++
												return b.off > len(b.b)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:29
	// _ = "end of CoverTab[91608]"
}

// Int32 returns a little endian int32 starting at current offset.
func (b byteReader) Int32() int32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:33
	_go_fuzz_dep_.CoverTab[91609]++
												b2 := b.b[b.off:]
												b2 = b2[:4]
												v3 := int32(b2[3])
												v2 := int32(b2[2])
												v1 := int32(b2[1])
												v0 := int32(b2[0])
												return v0 | (v1 << 8) | (v2 << 16) | (v3 << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:40
	// _ = "end of CoverTab[91609]"
}

// Uint8 returns the next byte
func (b *byteReader) Uint8() uint8 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:44
	_go_fuzz_dep_.CoverTab[91610]++
												v := b.b[b.off]
												return v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:46
	// _ = "end of CoverTab[91610]"
}

// Uint32 returns a little endian uint32 starting at current offset.
func (b byteReader) Uint32() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:50
	_go_fuzz_dep_.CoverTab[91611]++
												if r := b.remain(); r < 4 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:51
		_go_fuzz_dep_.CoverTab[91613]++

													v := uint32(0)
													for i := 1; i <= r; i++ {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:54
			_go_fuzz_dep_.CoverTab[91615]++
														v = (v << 8) | uint32(b.b[len(b.b)-i])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:55
			// _ = "end of CoverTab[91615]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:56
		// _ = "end of CoverTab[91613]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:56
		_go_fuzz_dep_.CoverTab[91614]++
													return v
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:57
		// _ = "end of CoverTab[91614]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:58
		_go_fuzz_dep_.CoverTab[91616]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:58
		// _ = "end of CoverTab[91616]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:58
	// _ = "end of CoverTab[91611]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:58
	_go_fuzz_dep_.CoverTab[91612]++
												b2 := b.b[b.off:]
												b2 = b2[:4]
												v3 := uint32(b2[3])
												v2 := uint32(b2[2])
												v1 := uint32(b2[1])
												v0 := uint32(b2[0])
												return v0 | (v1 << 8) | (v2 << 16) | (v3 << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:65
	// _ = "end of CoverTab[91612]"
}

// Uint32NC returns a little endian uint32 starting at current offset.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:68
// The caller must be sure if there are at least 4 bytes left.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:70
func (b byteReader) Uint32NC() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:70
	_go_fuzz_dep_.CoverTab[91617]++
												b2 := b.b[b.off:]
												b2 = b2[:4]
												v3 := uint32(b2[3])
												v2 := uint32(b2[2])
												v1 := uint32(b2[1])
												v0 := uint32(b2[0])
												return v0 | (v1 << 8) | (v2 << 16) | (v3 << 24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:77
	// _ = "end of CoverTab[91617]"
}

// unread returns the unread portion of the input.
func (b byteReader) unread() []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:81
	_go_fuzz_dep_.CoverTab[91618]++
												return b.b[b.off:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:82
	// _ = "end of CoverTab[91618]"
}

// remain will return the number of bytes remaining.
func (b byteReader) remain() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:86
	_go_fuzz_dep_.CoverTab[91619]++
												return len(b.b) - b.off
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:87
	// _ = "end of CoverTab[91619]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:88
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/zstd/bytereader.go:88
var _ = _go_fuzz_dep_.CoverTab
