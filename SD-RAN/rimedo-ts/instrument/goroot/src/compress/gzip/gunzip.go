// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/compress/gzip/gunzip.go:5
// Package gzip implements reading and writing of gzip format compressed files,
//line /usr/local/go/src/compress/gzip/gunzip.go:5
// as specified in RFC 1952.
//line /usr/local/go/src/compress/gzip/gunzip.go:7
package gzip

//line /usr/local/go/src/compress/gzip/gunzip.go:7
import (
//line /usr/local/go/src/compress/gzip/gunzip.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/compress/gzip/gunzip.go:7
)
//line /usr/local/go/src/compress/gzip/gunzip.go:7
import (
//line /usr/local/go/src/compress/gzip/gunzip.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/compress/gzip/gunzip.go:7
)

import (
	"bufio"
	"compress/flate"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
	"time"
)

const (
	gzipID1		= 0x1f
	gzipID2		= 0x8b
	gzipDeflate	= 8
	flagText	= 1 << 0
	flagHdrCrc	= 1 << 1
	flagExtra	= 1 << 2
	flagName	= 1 << 3
	flagComment	= 1 << 4
)

var (
	// ErrChecksum is returned when reading GZIP data that has an invalid checksum.
	ErrChecksum	= errors.New("gzip: invalid checksum")
	// ErrHeader is returned when reading GZIP data that has an invalid header.
	ErrHeader	= errors.New("gzip: invalid header")
)

var le = binary.LittleEndian

// noEOF converts io.EOF to io.ErrUnexpectedEOF.
func noEOF(err error) error {
//line /usr/local/go/src/compress/gzip/gunzip.go:40
	_go_fuzz_dep_.CoverTab[26655]++
							if err == io.EOF {
//line /usr/local/go/src/compress/gzip/gunzip.go:41
		_go_fuzz_dep_.CoverTab[26657]++
								return io.ErrUnexpectedEOF
//line /usr/local/go/src/compress/gzip/gunzip.go:42
		// _ = "end of CoverTab[26657]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:43
		_go_fuzz_dep_.CoverTab[26658]++
//line /usr/local/go/src/compress/gzip/gunzip.go:43
		// _ = "end of CoverTab[26658]"
//line /usr/local/go/src/compress/gzip/gunzip.go:43
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:43
	// _ = "end of CoverTab[26655]"
//line /usr/local/go/src/compress/gzip/gunzip.go:43
	_go_fuzz_dep_.CoverTab[26656]++
							return err
//line /usr/local/go/src/compress/gzip/gunzip.go:44
	// _ = "end of CoverTab[26656]"
}

// The gzip file stores a header giving metadata about the compressed file.
//line /usr/local/go/src/compress/gzip/gunzip.go:47
// That header is exposed as the fields of the Writer and Reader structs.
//line /usr/local/go/src/compress/gzip/gunzip.go:47
//
//line /usr/local/go/src/compress/gzip/gunzip.go:47
// Strings must be UTF-8 encoded and may only contain Unicode code points
//line /usr/local/go/src/compress/gzip/gunzip.go:47
// U+0001 through U+00FF, due to limitations of the GZIP file format.
//line /usr/local/go/src/compress/gzip/gunzip.go:52
type Header struct {
	Comment	string		// comment
	Extra	[]byte		// "extra data"
	ModTime	time.Time	// modification time
	Name	string		// file name
	OS	byte		// operating system type
}

// A Reader is an io.Reader that can be read to retrieve
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// uncompressed data from a gzip-format compressed file.
//line /usr/local/go/src/compress/gzip/gunzip.go:60
//
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// In general, a gzip file can be a concatenation of gzip files,
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// each with its own header. Reads from the Reader
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// return the concatenation of the uncompressed data of each.
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// Only the first header is recorded in the Reader fields.
//line /usr/local/go/src/compress/gzip/gunzip.go:60
//
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// Gzip files store a length and checksum of the uncompressed data.
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// The Reader will return an ErrChecksum when Read
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// reaches the end of the uncompressed data if it does not
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// have the expected length or checksum. Clients should treat data
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// returned by Read as tentative until they receive the io.EOF
//line /usr/local/go/src/compress/gzip/gunzip.go:60
// marking the end of the data.
//line /usr/local/go/src/compress/gzip/gunzip.go:74
type Reader struct {
	Header		// valid after NewReader or Reader.Reset
	r		flate.Reader
	decompressor	io.ReadCloser
	digest		uint32	// CRC-32, IEEE polynomial (section 8)
	size		uint32	// Uncompressed size (section 2.3.1)
	buf		[512]byte
	err		error
	multistream	bool
}

// NewReader creates a new Reader reading the given reader.
//line /usr/local/go/src/compress/gzip/gunzip.go:85
// If r does not also implement io.ByteReader,
//line /usr/local/go/src/compress/gzip/gunzip.go:85
// the decompressor may read more data than necessary from r.
//line /usr/local/go/src/compress/gzip/gunzip.go:85
//
//line /usr/local/go/src/compress/gzip/gunzip.go:85
// It is the caller's responsibility to call Close on the Reader when done.
//line /usr/local/go/src/compress/gzip/gunzip.go:85
//
//line /usr/local/go/src/compress/gzip/gunzip.go:85
// The Reader.Header fields will be valid in the Reader returned.
//line /usr/local/go/src/compress/gzip/gunzip.go:92
func NewReader(r io.Reader) (*Reader, error) {
//line /usr/local/go/src/compress/gzip/gunzip.go:92
	_go_fuzz_dep_.CoverTab[26659]++
							z := new(Reader)
							if err := z.Reset(r); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:94
		_go_fuzz_dep_.CoverTab[26661]++
								return nil, err
//line /usr/local/go/src/compress/gzip/gunzip.go:95
		// _ = "end of CoverTab[26661]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:96
		_go_fuzz_dep_.CoverTab[26662]++
//line /usr/local/go/src/compress/gzip/gunzip.go:96
		// _ = "end of CoverTab[26662]"
//line /usr/local/go/src/compress/gzip/gunzip.go:96
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:96
	// _ = "end of CoverTab[26659]"
//line /usr/local/go/src/compress/gzip/gunzip.go:96
	_go_fuzz_dep_.CoverTab[26660]++
							return z, nil
//line /usr/local/go/src/compress/gzip/gunzip.go:97
	// _ = "end of CoverTab[26660]"
}

// Reset discards the Reader z's state and makes it equivalent to the
//line /usr/local/go/src/compress/gzip/gunzip.go:100
// result of its original state from NewReader, but reading from r instead.
//line /usr/local/go/src/compress/gzip/gunzip.go:100
// This permits reusing a Reader rather than allocating a new one.
//line /usr/local/go/src/compress/gzip/gunzip.go:103
func (z *Reader) Reset(r io.Reader) error {
//line /usr/local/go/src/compress/gzip/gunzip.go:103
	_go_fuzz_dep_.CoverTab[26663]++
							*z = Reader{
		decompressor:	z.decompressor,
		multistream:	true,
	}
	if rr, ok := r.(flate.Reader); ok {
//line /usr/local/go/src/compress/gzip/gunzip.go:108
		_go_fuzz_dep_.CoverTab[26665]++
								z.r = rr
//line /usr/local/go/src/compress/gzip/gunzip.go:109
		// _ = "end of CoverTab[26665]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:110
		_go_fuzz_dep_.CoverTab[26666]++
								z.r = bufio.NewReader(r)
//line /usr/local/go/src/compress/gzip/gunzip.go:111
		// _ = "end of CoverTab[26666]"
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:112
	// _ = "end of CoverTab[26663]"
//line /usr/local/go/src/compress/gzip/gunzip.go:112
	_go_fuzz_dep_.CoverTab[26664]++
							z.Header, z.err = z.readHeader()
							return z.err
//line /usr/local/go/src/compress/gzip/gunzip.go:114
	// _ = "end of CoverTab[26664]"
}

// Multistream controls whether the reader supports multistream files.
//line /usr/local/go/src/compress/gzip/gunzip.go:117
//
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// If enabled (the default), the Reader expects the input to be a sequence
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// of individually gzipped data streams, each with its own header and
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// trailer, ending at EOF. The effect is that the concatenation of a sequence
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// of gzipped files is treated as equivalent to the gzip of the concatenation
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// of the sequence. This is standard behavior for gzip readers.
//line /usr/local/go/src/compress/gzip/gunzip.go:117
//
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// Calling Multistream(false) disables this behavior; disabling the behavior
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// can be useful when reading file formats that distinguish individual gzip
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// data streams or mix gzip data streams with other data streams.
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// In this mode, when the Reader reaches the end of the data stream,
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// Read returns io.EOF. The underlying reader must implement io.ByteReader
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// in order to be left positioned just after the gzip stream.
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// To start the next stream, call z.Reset(r) followed by z.Multistream(false).
//line /usr/local/go/src/compress/gzip/gunzip.go:117
// If there is no next stream, z.Reset(r) will return io.EOF.
//line /usr/local/go/src/compress/gzip/gunzip.go:133
func (z *Reader) Multistream(ok bool) {
//line /usr/local/go/src/compress/gzip/gunzip.go:133
	_go_fuzz_dep_.CoverTab[26667]++
							z.multistream = ok
//line /usr/local/go/src/compress/gzip/gunzip.go:134
	// _ = "end of CoverTab[26667]"
}

// readString reads a NUL-terminated string from z.r.
//line /usr/local/go/src/compress/gzip/gunzip.go:137
// It treats the bytes read as being encoded as ISO 8859-1 (Latin-1) and
//line /usr/local/go/src/compress/gzip/gunzip.go:137
// will output a string encoded using UTF-8.
//line /usr/local/go/src/compress/gzip/gunzip.go:137
// This method always updates z.digest with the data read.
//line /usr/local/go/src/compress/gzip/gunzip.go:141
func (z *Reader) readString() (string, error) {
//line /usr/local/go/src/compress/gzip/gunzip.go:141
	_go_fuzz_dep_.CoverTab[26668]++
							var err error
							needConv := false
							for i := 0; ; i++ {
//line /usr/local/go/src/compress/gzip/gunzip.go:144
		_go_fuzz_dep_.CoverTab[26669]++
								if i >= len(z.buf) {
//line /usr/local/go/src/compress/gzip/gunzip.go:145
			_go_fuzz_dep_.CoverTab[26673]++
									return "", ErrHeader
//line /usr/local/go/src/compress/gzip/gunzip.go:146
			// _ = "end of CoverTab[26673]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:147
			_go_fuzz_dep_.CoverTab[26674]++
//line /usr/local/go/src/compress/gzip/gunzip.go:147
			// _ = "end of CoverTab[26674]"
//line /usr/local/go/src/compress/gzip/gunzip.go:147
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:147
		// _ = "end of CoverTab[26669]"
//line /usr/local/go/src/compress/gzip/gunzip.go:147
		_go_fuzz_dep_.CoverTab[26670]++
								z.buf[i], err = z.r.ReadByte()
								if err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:149
			_go_fuzz_dep_.CoverTab[26675]++
									return "", err
//line /usr/local/go/src/compress/gzip/gunzip.go:150
			// _ = "end of CoverTab[26675]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:151
			_go_fuzz_dep_.CoverTab[26676]++
//line /usr/local/go/src/compress/gzip/gunzip.go:151
			// _ = "end of CoverTab[26676]"
//line /usr/local/go/src/compress/gzip/gunzip.go:151
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:151
		// _ = "end of CoverTab[26670]"
//line /usr/local/go/src/compress/gzip/gunzip.go:151
		_go_fuzz_dep_.CoverTab[26671]++
								if z.buf[i] > 0x7f {
//line /usr/local/go/src/compress/gzip/gunzip.go:152
			_go_fuzz_dep_.CoverTab[26677]++
									needConv = true
//line /usr/local/go/src/compress/gzip/gunzip.go:153
			// _ = "end of CoverTab[26677]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:154
			_go_fuzz_dep_.CoverTab[26678]++
//line /usr/local/go/src/compress/gzip/gunzip.go:154
			// _ = "end of CoverTab[26678]"
//line /usr/local/go/src/compress/gzip/gunzip.go:154
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:154
		// _ = "end of CoverTab[26671]"
//line /usr/local/go/src/compress/gzip/gunzip.go:154
		_go_fuzz_dep_.CoverTab[26672]++
								if z.buf[i] == 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:155
			_go_fuzz_dep_.CoverTab[26679]++

									z.digest = crc32.Update(z.digest, crc32.IEEETable, z.buf[:i+1])

//line /usr/local/go/src/compress/gzip/gunzip.go:160
			if needConv {
//line /usr/local/go/src/compress/gzip/gunzip.go:160
				_go_fuzz_dep_.CoverTab[26681]++
										s := make([]rune, 0, i)
										for _, v := range z.buf[:i] {
//line /usr/local/go/src/compress/gzip/gunzip.go:162
					_go_fuzz_dep_.CoverTab[26683]++
											s = append(s, rune(v))
//line /usr/local/go/src/compress/gzip/gunzip.go:163
					// _ = "end of CoverTab[26683]"
				}
//line /usr/local/go/src/compress/gzip/gunzip.go:164
				// _ = "end of CoverTab[26681]"
//line /usr/local/go/src/compress/gzip/gunzip.go:164
				_go_fuzz_dep_.CoverTab[26682]++
										return string(s), nil
//line /usr/local/go/src/compress/gzip/gunzip.go:165
				// _ = "end of CoverTab[26682]"
			} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:166
				_go_fuzz_dep_.CoverTab[26684]++
//line /usr/local/go/src/compress/gzip/gunzip.go:166
				// _ = "end of CoverTab[26684]"
//line /usr/local/go/src/compress/gzip/gunzip.go:166
			}
//line /usr/local/go/src/compress/gzip/gunzip.go:166
			// _ = "end of CoverTab[26679]"
//line /usr/local/go/src/compress/gzip/gunzip.go:166
			_go_fuzz_dep_.CoverTab[26680]++
									return string(z.buf[:i]), nil
//line /usr/local/go/src/compress/gzip/gunzip.go:167
			// _ = "end of CoverTab[26680]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:168
			_go_fuzz_dep_.CoverTab[26685]++
//line /usr/local/go/src/compress/gzip/gunzip.go:168
			// _ = "end of CoverTab[26685]"
//line /usr/local/go/src/compress/gzip/gunzip.go:168
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:168
		// _ = "end of CoverTab[26672]"
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:169
	// _ = "end of CoverTab[26668]"
}

// readHeader reads the GZIP header according to section 2.3.1.
//line /usr/local/go/src/compress/gzip/gunzip.go:172
// This method does not set z.err.
//line /usr/local/go/src/compress/gzip/gunzip.go:174
func (z *Reader) readHeader() (hdr Header, err error) {
//line /usr/local/go/src/compress/gzip/gunzip.go:174
	_go_fuzz_dep_.CoverTab[26686]++
							if _, err = io.ReadFull(z.r, z.buf[:10]); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:175
		_go_fuzz_dep_.CoverTab[26695]++

//line /usr/local/go/src/compress/gzip/gunzip.go:183
		return hdr, err
//line /usr/local/go/src/compress/gzip/gunzip.go:183
		// _ = "end of CoverTab[26695]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:184
		_go_fuzz_dep_.CoverTab[26696]++
//line /usr/local/go/src/compress/gzip/gunzip.go:184
		// _ = "end of CoverTab[26696]"
//line /usr/local/go/src/compress/gzip/gunzip.go:184
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:184
	// _ = "end of CoverTab[26686]"
//line /usr/local/go/src/compress/gzip/gunzip.go:184
	_go_fuzz_dep_.CoverTab[26687]++
							if z.buf[0] != gzipID1 || func() bool {
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		_go_fuzz_dep_.CoverTab[26697]++
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		return z.buf[1] != gzipID2
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		// _ = "end of CoverTab[26697]"
//line /usr/local/go/src/compress/gzip/gunzip.go:185
	}() || func() bool {
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		_go_fuzz_dep_.CoverTab[26698]++
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		return z.buf[2] != gzipDeflate
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		// _ = "end of CoverTab[26698]"
//line /usr/local/go/src/compress/gzip/gunzip.go:185
	}() {
//line /usr/local/go/src/compress/gzip/gunzip.go:185
		_go_fuzz_dep_.CoverTab[26699]++
								return hdr, ErrHeader
//line /usr/local/go/src/compress/gzip/gunzip.go:186
		// _ = "end of CoverTab[26699]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:187
		_go_fuzz_dep_.CoverTab[26700]++
//line /usr/local/go/src/compress/gzip/gunzip.go:187
		// _ = "end of CoverTab[26700]"
//line /usr/local/go/src/compress/gzip/gunzip.go:187
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:187
	// _ = "end of CoverTab[26687]"
//line /usr/local/go/src/compress/gzip/gunzip.go:187
	_go_fuzz_dep_.CoverTab[26688]++
							flg := z.buf[3]
							if t := int64(le.Uint32(z.buf[4:8])); t > 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:189
		_go_fuzz_dep_.CoverTab[26701]++

//line /usr/local/go/src/compress/gzip/gunzip.go:192
		hdr.ModTime = time.Unix(t, 0)
//line /usr/local/go/src/compress/gzip/gunzip.go:192
		// _ = "end of CoverTab[26701]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:193
		_go_fuzz_dep_.CoverTab[26702]++
//line /usr/local/go/src/compress/gzip/gunzip.go:193
		// _ = "end of CoverTab[26702]"
//line /usr/local/go/src/compress/gzip/gunzip.go:193
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:193
	// _ = "end of CoverTab[26688]"
//line /usr/local/go/src/compress/gzip/gunzip.go:193
	_go_fuzz_dep_.CoverTab[26689]++

							hdr.OS = z.buf[9]
							z.digest = crc32.ChecksumIEEE(z.buf[:10])

							if flg&flagExtra != 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:198
		_go_fuzz_dep_.CoverTab[26703]++
								if _, err = io.ReadFull(z.r, z.buf[:2]); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:199
			_go_fuzz_dep_.CoverTab[26706]++
									return hdr, noEOF(err)
//line /usr/local/go/src/compress/gzip/gunzip.go:200
			// _ = "end of CoverTab[26706]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:201
			_go_fuzz_dep_.CoverTab[26707]++
//line /usr/local/go/src/compress/gzip/gunzip.go:201
			// _ = "end of CoverTab[26707]"
//line /usr/local/go/src/compress/gzip/gunzip.go:201
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:201
		// _ = "end of CoverTab[26703]"
//line /usr/local/go/src/compress/gzip/gunzip.go:201
		_go_fuzz_dep_.CoverTab[26704]++
								z.digest = crc32.Update(z.digest, crc32.IEEETable, z.buf[:2])
								data := make([]byte, le.Uint16(z.buf[:2]))
								if _, err = io.ReadFull(z.r, data); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:204
			_go_fuzz_dep_.CoverTab[26708]++
									return hdr, noEOF(err)
//line /usr/local/go/src/compress/gzip/gunzip.go:205
			// _ = "end of CoverTab[26708]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:206
			_go_fuzz_dep_.CoverTab[26709]++
//line /usr/local/go/src/compress/gzip/gunzip.go:206
			// _ = "end of CoverTab[26709]"
//line /usr/local/go/src/compress/gzip/gunzip.go:206
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:206
		// _ = "end of CoverTab[26704]"
//line /usr/local/go/src/compress/gzip/gunzip.go:206
		_go_fuzz_dep_.CoverTab[26705]++
								z.digest = crc32.Update(z.digest, crc32.IEEETable, data)
								hdr.Extra = data
//line /usr/local/go/src/compress/gzip/gunzip.go:208
		// _ = "end of CoverTab[26705]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:209
		_go_fuzz_dep_.CoverTab[26710]++
//line /usr/local/go/src/compress/gzip/gunzip.go:209
		// _ = "end of CoverTab[26710]"
//line /usr/local/go/src/compress/gzip/gunzip.go:209
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:209
	// _ = "end of CoverTab[26689]"
//line /usr/local/go/src/compress/gzip/gunzip.go:209
	_go_fuzz_dep_.CoverTab[26690]++

							var s string
							if flg&flagName != 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:212
		_go_fuzz_dep_.CoverTab[26711]++
								if s, err = z.readString(); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:213
			_go_fuzz_dep_.CoverTab[26713]++
									return hdr, noEOF(err)
//line /usr/local/go/src/compress/gzip/gunzip.go:214
			// _ = "end of CoverTab[26713]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:215
			_go_fuzz_dep_.CoverTab[26714]++
//line /usr/local/go/src/compress/gzip/gunzip.go:215
			// _ = "end of CoverTab[26714]"
//line /usr/local/go/src/compress/gzip/gunzip.go:215
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:215
		// _ = "end of CoverTab[26711]"
//line /usr/local/go/src/compress/gzip/gunzip.go:215
		_go_fuzz_dep_.CoverTab[26712]++
								hdr.Name = s
//line /usr/local/go/src/compress/gzip/gunzip.go:216
		// _ = "end of CoverTab[26712]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:217
		_go_fuzz_dep_.CoverTab[26715]++
//line /usr/local/go/src/compress/gzip/gunzip.go:217
		// _ = "end of CoverTab[26715]"
//line /usr/local/go/src/compress/gzip/gunzip.go:217
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:217
	// _ = "end of CoverTab[26690]"
//line /usr/local/go/src/compress/gzip/gunzip.go:217
	_go_fuzz_dep_.CoverTab[26691]++

							if flg&flagComment != 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:219
		_go_fuzz_dep_.CoverTab[26716]++
								if s, err = z.readString(); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:220
			_go_fuzz_dep_.CoverTab[26718]++
									return hdr, noEOF(err)
//line /usr/local/go/src/compress/gzip/gunzip.go:221
			// _ = "end of CoverTab[26718]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:222
			_go_fuzz_dep_.CoverTab[26719]++
//line /usr/local/go/src/compress/gzip/gunzip.go:222
			// _ = "end of CoverTab[26719]"
//line /usr/local/go/src/compress/gzip/gunzip.go:222
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:222
		// _ = "end of CoverTab[26716]"
//line /usr/local/go/src/compress/gzip/gunzip.go:222
		_go_fuzz_dep_.CoverTab[26717]++
								hdr.Comment = s
//line /usr/local/go/src/compress/gzip/gunzip.go:223
		// _ = "end of CoverTab[26717]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:224
		_go_fuzz_dep_.CoverTab[26720]++
//line /usr/local/go/src/compress/gzip/gunzip.go:224
		// _ = "end of CoverTab[26720]"
//line /usr/local/go/src/compress/gzip/gunzip.go:224
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:224
	// _ = "end of CoverTab[26691]"
//line /usr/local/go/src/compress/gzip/gunzip.go:224
	_go_fuzz_dep_.CoverTab[26692]++

							if flg&flagHdrCrc != 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:226
		_go_fuzz_dep_.CoverTab[26721]++
								if _, err = io.ReadFull(z.r, z.buf[:2]); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:227
			_go_fuzz_dep_.CoverTab[26723]++
									return hdr, noEOF(err)
//line /usr/local/go/src/compress/gzip/gunzip.go:228
			// _ = "end of CoverTab[26723]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:229
			_go_fuzz_dep_.CoverTab[26724]++
//line /usr/local/go/src/compress/gzip/gunzip.go:229
			// _ = "end of CoverTab[26724]"
//line /usr/local/go/src/compress/gzip/gunzip.go:229
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:229
		// _ = "end of CoverTab[26721]"
//line /usr/local/go/src/compress/gzip/gunzip.go:229
		_go_fuzz_dep_.CoverTab[26722]++
								digest := le.Uint16(z.buf[:2])
								if digest != uint16(z.digest) {
//line /usr/local/go/src/compress/gzip/gunzip.go:231
			_go_fuzz_dep_.CoverTab[26725]++
									return hdr, ErrHeader
//line /usr/local/go/src/compress/gzip/gunzip.go:232
			// _ = "end of CoverTab[26725]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:233
			_go_fuzz_dep_.CoverTab[26726]++
//line /usr/local/go/src/compress/gzip/gunzip.go:233
			// _ = "end of CoverTab[26726]"
//line /usr/local/go/src/compress/gzip/gunzip.go:233
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:233
		// _ = "end of CoverTab[26722]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:234
		_go_fuzz_dep_.CoverTab[26727]++
//line /usr/local/go/src/compress/gzip/gunzip.go:234
		// _ = "end of CoverTab[26727]"
//line /usr/local/go/src/compress/gzip/gunzip.go:234
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:234
	// _ = "end of CoverTab[26692]"
//line /usr/local/go/src/compress/gzip/gunzip.go:234
	_go_fuzz_dep_.CoverTab[26693]++

							z.digest = 0
							if z.decompressor == nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:237
		_go_fuzz_dep_.CoverTab[26728]++
								z.decompressor = flate.NewReader(z.r)
//line /usr/local/go/src/compress/gzip/gunzip.go:238
		// _ = "end of CoverTab[26728]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:239
		_go_fuzz_dep_.CoverTab[26729]++
								z.decompressor.(flate.Resetter).Reset(z.r, nil)
//line /usr/local/go/src/compress/gzip/gunzip.go:240
		// _ = "end of CoverTab[26729]"
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:241
	// _ = "end of CoverTab[26693]"
//line /usr/local/go/src/compress/gzip/gunzip.go:241
	_go_fuzz_dep_.CoverTab[26694]++
							return hdr, nil
//line /usr/local/go/src/compress/gzip/gunzip.go:242
	// _ = "end of CoverTab[26694]"
}

// Read implements io.Reader, reading uncompressed bytes from its underlying Reader.
func (z *Reader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/compress/gzip/gunzip.go:246
	_go_fuzz_dep_.CoverTab[26730]++
							if z.err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:247
		_go_fuzz_dep_.CoverTab[26733]++
								return 0, z.err
//line /usr/local/go/src/compress/gzip/gunzip.go:248
		// _ = "end of CoverTab[26733]"
	} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:249
		_go_fuzz_dep_.CoverTab[26734]++
//line /usr/local/go/src/compress/gzip/gunzip.go:249
		// _ = "end of CoverTab[26734]"
//line /usr/local/go/src/compress/gzip/gunzip.go:249
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:249
	// _ = "end of CoverTab[26730]"
//line /usr/local/go/src/compress/gzip/gunzip.go:249
	_go_fuzz_dep_.CoverTab[26731]++

							for n == 0 {
//line /usr/local/go/src/compress/gzip/gunzip.go:251
		_go_fuzz_dep_.CoverTab[26735]++
								n, z.err = z.decompressor.Read(p)
								z.digest = crc32.Update(z.digest, crc32.IEEETable, p[:n])
								z.size += uint32(n)
								if z.err != io.EOF {
//line /usr/local/go/src/compress/gzip/gunzip.go:255
			_go_fuzz_dep_.CoverTab[26740]++

									return n, z.err
//line /usr/local/go/src/compress/gzip/gunzip.go:257
			// _ = "end of CoverTab[26740]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:258
			_go_fuzz_dep_.CoverTab[26741]++
//line /usr/local/go/src/compress/gzip/gunzip.go:258
			// _ = "end of CoverTab[26741]"
//line /usr/local/go/src/compress/gzip/gunzip.go:258
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:258
		// _ = "end of CoverTab[26735]"
//line /usr/local/go/src/compress/gzip/gunzip.go:258
		_go_fuzz_dep_.CoverTab[26736]++

//line /usr/local/go/src/compress/gzip/gunzip.go:261
		if _, err := io.ReadFull(z.r, z.buf[:8]); err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:261
			_go_fuzz_dep_.CoverTab[26742]++
									z.err = noEOF(err)
									return n, z.err
//line /usr/local/go/src/compress/gzip/gunzip.go:263
			// _ = "end of CoverTab[26742]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:264
			_go_fuzz_dep_.CoverTab[26743]++
//line /usr/local/go/src/compress/gzip/gunzip.go:264
			// _ = "end of CoverTab[26743]"
//line /usr/local/go/src/compress/gzip/gunzip.go:264
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:264
		// _ = "end of CoverTab[26736]"
//line /usr/local/go/src/compress/gzip/gunzip.go:264
		_go_fuzz_dep_.CoverTab[26737]++
								digest := le.Uint32(z.buf[:4])
								size := le.Uint32(z.buf[4:8])
								if digest != z.digest || func() bool {
//line /usr/local/go/src/compress/gzip/gunzip.go:267
			_go_fuzz_dep_.CoverTab[26744]++
//line /usr/local/go/src/compress/gzip/gunzip.go:267
			return size != z.size
//line /usr/local/go/src/compress/gzip/gunzip.go:267
			// _ = "end of CoverTab[26744]"
//line /usr/local/go/src/compress/gzip/gunzip.go:267
		}() {
//line /usr/local/go/src/compress/gzip/gunzip.go:267
			_go_fuzz_dep_.CoverTab[26745]++
									z.err = ErrChecksum
									return n, z.err
//line /usr/local/go/src/compress/gzip/gunzip.go:269
			// _ = "end of CoverTab[26745]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:270
			_go_fuzz_dep_.CoverTab[26746]++
//line /usr/local/go/src/compress/gzip/gunzip.go:270
			// _ = "end of CoverTab[26746]"
//line /usr/local/go/src/compress/gzip/gunzip.go:270
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:270
		// _ = "end of CoverTab[26737]"
//line /usr/local/go/src/compress/gzip/gunzip.go:270
		_go_fuzz_dep_.CoverTab[26738]++
								z.digest, z.size = 0, 0

//line /usr/local/go/src/compress/gzip/gunzip.go:274
		if !z.multistream {
//line /usr/local/go/src/compress/gzip/gunzip.go:274
			_go_fuzz_dep_.CoverTab[26747]++
									return n, io.EOF
//line /usr/local/go/src/compress/gzip/gunzip.go:275
			// _ = "end of CoverTab[26747]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:276
			_go_fuzz_dep_.CoverTab[26748]++
//line /usr/local/go/src/compress/gzip/gunzip.go:276
			// _ = "end of CoverTab[26748]"
//line /usr/local/go/src/compress/gzip/gunzip.go:276
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:276
		// _ = "end of CoverTab[26738]"
//line /usr/local/go/src/compress/gzip/gunzip.go:276
		_go_fuzz_dep_.CoverTab[26739]++
								z.err = nil

								if _, z.err = z.readHeader(); z.err != nil {
//line /usr/local/go/src/compress/gzip/gunzip.go:279
			_go_fuzz_dep_.CoverTab[26749]++
									return n, z.err
//line /usr/local/go/src/compress/gzip/gunzip.go:280
			// _ = "end of CoverTab[26749]"
		} else {
//line /usr/local/go/src/compress/gzip/gunzip.go:281
			_go_fuzz_dep_.CoverTab[26750]++
//line /usr/local/go/src/compress/gzip/gunzip.go:281
			// _ = "end of CoverTab[26750]"
//line /usr/local/go/src/compress/gzip/gunzip.go:281
		}
//line /usr/local/go/src/compress/gzip/gunzip.go:281
		// _ = "end of CoverTab[26739]"
	}
//line /usr/local/go/src/compress/gzip/gunzip.go:282
	// _ = "end of CoverTab[26731]"
//line /usr/local/go/src/compress/gzip/gunzip.go:282
	_go_fuzz_dep_.CoverTab[26732]++

							return n, nil
//line /usr/local/go/src/compress/gzip/gunzip.go:284
	// _ = "end of CoverTab[26732]"
}

// Close closes the Reader. It does not close the underlying io.Reader.
//line /usr/local/go/src/compress/gzip/gunzip.go:287
// In order for the GZIP checksum to be verified, the reader must be
//line /usr/local/go/src/compress/gzip/gunzip.go:287
// fully consumed until the io.EOF.
//line /usr/local/go/src/compress/gzip/gunzip.go:290
func (z *Reader) Close() error {
//line /usr/local/go/src/compress/gzip/gunzip.go:290
	_go_fuzz_dep_.CoverTab[26751]++
//line /usr/local/go/src/compress/gzip/gunzip.go:290
	return z.decompressor.Close()
//line /usr/local/go/src/compress/gzip/gunzip.go:290
	// _ = "end of CoverTab[26751]"
//line /usr/local/go/src/compress/gzip/gunzip.go:290
}

//line /usr/local/go/src/compress/gzip/gunzip.go:290
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/compress/gzip/gunzip.go:290
var _ = _go_fuzz_dep_.CoverTab
