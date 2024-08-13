// Copyright 2018 Klaus Post. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Based on work Copyright (c) 2013, Yann Collet, released under BSD License.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
package huff0

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:6
)

// byteReader provides a byte reader that reads
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:8
// little endian values from a byte stream.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:8
// The input stream is manually advanced.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:8
// The reader performs no bounds checks.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:12
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
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:24
	_go_fuzz_dep_.CoverTab[89530]++
												b.off += int(n)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:25
	// _ = "end of CoverTab[89530]"
}

// Int32 returns a little endian int32 starting at current offset.
func (b byteReader) Int32() int32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:29
	_go_fuzz_dep_.CoverTab[89531]++
												v3 := int32(b.b[b.off+3])
												v2 := int32(b.b[b.off+2])
												v1 := int32(b.b[b.off+1])
												v0 := int32(b.b[b.off])
												return (v3 << 24) | (v2 << 16) | (v1 << 8) | v0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:34
	// _ = "end of CoverTab[89531]"
}

// Uint32 returns a little endian uint32 starting at current offset.
func (b byteReader) Uint32() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:38
	_go_fuzz_dep_.CoverTab[89532]++
												v3 := uint32(b.b[b.off+3])
												v2 := uint32(b.b[b.off+2])
												v1 := uint32(b.b[b.off+1])
												v0 := uint32(b.b[b.off])
												return (v3 << 24) | (v2 << 16) | (v1 << 8) | v0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:43
	// _ = "end of CoverTab[89532]"
}

// unread returns the unread portion of the input.
func (b byteReader) unread() []byte {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:47
	_go_fuzz_dep_.CoverTab[89533]++
												return b.b[b.off:]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:48
	// _ = "end of CoverTab[89533]"
}

// remain will return the number of bytes remaining.
func (b byteReader) remain() int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:52
	_go_fuzz_dep_.CoverTab[89534]++
												return len(b.b) - b.off
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:53
	// _ = "end of CoverTab[89534]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/huff0/bytereader.go:54
var _ = _go_fuzz_dep_.CoverTab
