// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The wire protocol for HTTP's "chunked" Transfer-Encoding.

//line /usr/local/go/src/net/http/internal/chunked.go:7
// Package internal contains HTTP internals shared by net/http and
//line /usr/local/go/src/net/http/internal/chunked.go:7
// net/http/httputil.
//line /usr/local/go/src/net/http/internal/chunked.go:9
package internal

//line /usr/local/go/src/net/http/internal/chunked.go:9
import (
//line /usr/local/go/src/net/http/internal/chunked.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/internal/chunked.go:9
)
//line /usr/local/go/src/net/http/internal/chunked.go:9
import (
//line /usr/local/go/src/net/http/internal/chunked.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/internal/chunked.go:9
)

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

const maxLineLength = 4096	// assumed <= bufio.defaultBufSize

var ErrLineTooLong = errors.New("header line too long")

// NewChunkedReader returns a new chunkedReader that translates the data read from r
//line /usr/local/go/src/net/http/internal/chunked.go:23
// out of HTTP "chunked" format before returning it.
//line /usr/local/go/src/net/http/internal/chunked.go:23
// The chunkedReader returns io.EOF when the final 0-length chunk is read.
//line /usr/local/go/src/net/http/internal/chunked.go:23
//
//line /usr/local/go/src/net/http/internal/chunked.go:23
// NewChunkedReader is not needed by normal applications. The http package
//line /usr/local/go/src/net/http/internal/chunked.go:23
// automatically decodes chunking when reading response bodies.
//line /usr/local/go/src/net/http/internal/chunked.go:29
func NewChunkedReader(r io.Reader) io.Reader {
//line /usr/local/go/src/net/http/internal/chunked.go:29
	_go_fuzz_dep_.CoverTab[36459]++
								br, ok := r.(*bufio.Reader)
								if !ok {
//line /usr/local/go/src/net/http/internal/chunked.go:31
		_go_fuzz_dep_.CoverTab[36461]++
									br = bufio.NewReader(r)
//line /usr/local/go/src/net/http/internal/chunked.go:32
		// _ = "end of CoverTab[36461]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:33
		_go_fuzz_dep_.CoverTab[36462]++
//line /usr/local/go/src/net/http/internal/chunked.go:33
		// _ = "end of CoverTab[36462]"
//line /usr/local/go/src/net/http/internal/chunked.go:33
	}
//line /usr/local/go/src/net/http/internal/chunked.go:33
	// _ = "end of CoverTab[36459]"
//line /usr/local/go/src/net/http/internal/chunked.go:33
	_go_fuzz_dep_.CoverTab[36460]++
								return &chunkedReader{r: br}
//line /usr/local/go/src/net/http/internal/chunked.go:34
	// _ = "end of CoverTab[36460]"
}

type chunkedReader struct {
	r		*bufio.Reader
	n		uint64	// unread bytes in chunk
	err		error
	buf		[2]byte
	checkEnd	bool	// whether need to check for \r\n chunk footer
}

func (cr *chunkedReader) beginChunk() {
//line /usr/local/go/src/net/http/internal/chunked.go:45
	_go_fuzz_dep_.CoverTab[36463]++
	// chunk-size CRLF
	var line []byte
	line, cr.err = readChunkLine(cr.r)
	if cr.err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:49
		_go_fuzz_dep_.CoverTab[36466]++
									return
//line /usr/local/go/src/net/http/internal/chunked.go:50
		// _ = "end of CoverTab[36466]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:51
		_go_fuzz_dep_.CoverTab[36467]++
//line /usr/local/go/src/net/http/internal/chunked.go:51
		// _ = "end of CoverTab[36467]"
//line /usr/local/go/src/net/http/internal/chunked.go:51
	}
//line /usr/local/go/src/net/http/internal/chunked.go:51
	// _ = "end of CoverTab[36463]"
//line /usr/local/go/src/net/http/internal/chunked.go:51
	_go_fuzz_dep_.CoverTab[36464]++
								cr.n, cr.err = parseHexUint(line)
								if cr.err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:53
		_go_fuzz_dep_.CoverTab[36468]++
									return
//line /usr/local/go/src/net/http/internal/chunked.go:54
		// _ = "end of CoverTab[36468]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:55
		_go_fuzz_dep_.CoverTab[36469]++
//line /usr/local/go/src/net/http/internal/chunked.go:55
		// _ = "end of CoverTab[36469]"
//line /usr/local/go/src/net/http/internal/chunked.go:55
	}
//line /usr/local/go/src/net/http/internal/chunked.go:55
	// _ = "end of CoverTab[36464]"
//line /usr/local/go/src/net/http/internal/chunked.go:55
	_go_fuzz_dep_.CoverTab[36465]++
								if cr.n == 0 {
//line /usr/local/go/src/net/http/internal/chunked.go:56
		_go_fuzz_dep_.CoverTab[36470]++
									cr.err = io.EOF
//line /usr/local/go/src/net/http/internal/chunked.go:57
		// _ = "end of CoverTab[36470]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:58
		_go_fuzz_dep_.CoverTab[36471]++
//line /usr/local/go/src/net/http/internal/chunked.go:58
		// _ = "end of CoverTab[36471]"
//line /usr/local/go/src/net/http/internal/chunked.go:58
	}
//line /usr/local/go/src/net/http/internal/chunked.go:58
	// _ = "end of CoverTab[36465]"
}

func (cr *chunkedReader) chunkHeaderAvailable() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:61
	_go_fuzz_dep_.CoverTab[36472]++
								n := cr.r.Buffered()
								if n > 0 {
//line /usr/local/go/src/net/http/internal/chunked.go:63
		_go_fuzz_dep_.CoverTab[36474]++
									peek, _ := cr.r.Peek(n)
									return bytes.IndexByte(peek, '\n') >= 0
//line /usr/local/go/src/net/http/internal/chunked.go:65
		// _ = "end of CoverTab[36474]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:66
		_go_fuzz_dep_.CoverTab[36475]++
//line /usr/local/go/src/net/http/internal/chunked.go:66
		// _ = "end of CoverTab[36475]"
//line /usr/local/go/src/net/http/internal/chunked.go:66
	}
//line /usr/local/go/src/net/http/internal/chunked.go:66
	// _ = "end of CoverTab[36472]"
//line /usr/local/go/src/net/http/internal/chunked.go:66
	_go_fuzz_dep_.CoverTab[36473]++
								return false
//line /usr/local/go/src/net/http/internal/chunked.go:67
	// _ = "end of CoverTab[36473]"
}

func (cr *chunkedReader) Read(b []uint8) (n int, err error) {
//line /usr/local/go/src/net/http/internal/chunked.go:70
	_go_fuzz_dep_.CoverTab[36476]++
								for cr.err == nil {
//line /usr/local/go/src/net/http/internal/chunked.go:71
		_go_fuzz_dep_.CoverTab[36478]++
									if cr.checkEnd {
//line /usr/local/go/src/net/http/internal/chunked.go:72
			_go_fuzz_dep_.CoverTab[36483]++
										if n > 0 && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:73
				_go_fuzz_dep_.CoverTab[36486]++
//line /usr/local/go/src/net/http/internal/chunked.go:73
				return cr.r.Buffered() < 2
//line /usr/local/go/src/net/http/internal/chunked.go:73
				// _ = "end of CoverTab[36486]"
//line /usr/local/go/src/net/http/internal/chunked.go:73
			}() {
//line /usr/local/go/src/net/http/internal/chunked.go:73
				_go_fuzz_dep_.CoverTab[36487]++

//line /usr/local/go/src/net/http/internal/chunked.go:77
				break
//line /usr/local/go/src/net/http/internal/chunked.go:77
				// _ = "end of CoverTab[36487]"
			} else {
//line /usr/local/go/src/net/http/internal/chunked.go:78
				_go_fuzz_dep_.CoverTab[36488]++
//line /usr/local/go/src/net/http/internal/chunked.go:78
				// _ = "end of CoverTab[36488]"
//line /usr/local/go/src/net/http/internal/chunked.go:78
			}
//line /usr/local/go/src/net/http/internal/chunked.go:78
			// _ = "end of CoverTab[36483]"
//line /usr/local/go/src/net/http/internal/chunked.go:78
			_go_fuzz_dep_.CoverTab[36484]++
										if _, cr.err = io.ReadFull(cr.r, cr.buf[:2]); cr.err == nil {
//line /usr/local/go/src/net/http/internal/chunked.go:79
				_go_fuzz_dep_.CoverTab[36489]++
											if string(cr.buf[:]) != "\r\n" {
//line /usr/local/go/src/net/http/internal/chunked.go:80
					_go_fuzz_dep_.CoverTab[36490]++
												cr.err = errors.New("malformed chunked encoding")
												break
//line /usr/local/go/src/net/http/internal/chunked.go:82
					// _ = "end of CoverTab[36490]"
				} else {
//line /usr/local/go/src/net/http/internal/chunked.go:83
					_go_fuzz_dep_.CoverTab[36491]++
//line /usr/local/go/src/net/http/internal/chunked.go:83
					// _ = "end of CoverTab[36491]"
//line /usr/local/go/src/net/http/internal/chunked.go:83
				}
//line /usr/local/go/src/net/http/internal/chunked.go:83
				// _ = "end of CoverTab[36489]"
			} else {
//line /usr/local/go/src/net/http/internal/chunked.go:84
				_go_fuzz_dep_.CoverTab[36492]++
											if cr.err == io.EOF {
//line /usr/local/go/src/net/http/internal/chunked.go:85
					_go_fuzz_dep_.CoverTab[36494]++
												cr.err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/internal/chunked.go:86
					// _ = "end of CoverTab[36494]"
				} else {
//line /usr/local/go/src/net/http/internal/chunked.go:87
					_go_fuzz_dep_.CoverTab[36495]++
//line /usr/local/go/src/net/http/internal/chunked.go:87
					// _ = "end of CoverTab[36495]"
//line /usr/local/go/src/net/http/internal/chunked.go:87
				}
//line /usr/local/go/src/net/http/internal/chunked.go:87
				// _ = "end of CoverTab[36492]"
//line /usr/local/go/src/net/http/internal/chunked.go:87
				_go_fuzz_dep_.CoverTab[36493]++
											break
//line /usr/local/go/src/net/http/internal/chunked.go:88
				// _ = "end of CoverTab[36493]"
			}
//line /usr/local/go/src/net/http/internal/chunked.go:89
			// _ = "end of CoverTab[36484]"
//line /usr/local/go/src/net/http/internal/chunked.go:89
			_go_fuzz_dep_.CoverTab[36485]++
										cr.checkEnd = false
//line /usr/local/go/src/net/http/internal/chunked.go:90
			// _ = "end of CoverTab[36485]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:91
			_go_fuzz_dep_.CoverTab[36496]++
//line /usr/local/go/src/net/http/internal/chunked.go:91
			// _ = "end of CoverTab[36496]"
//line /usr/local/go/src/net/http/internal/chunked.go:91
		}
//line /usr/local/go/src/net/http/internal/chunked.go:91
		// _ = "end of CoverTab[36478]"
//line /usr/local/go/src/net/http/internal/chunked.go:91
		_go_fuzz_dep_.CoverTab[36479]++
									if cr.n == 0 {
//line /usr/local/go/src/net/http/internal/chunked.go:92
			_go_fuzz_dep_.CoverTab[36497]++
										if n > 0 && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:93
				_go_fuzz_dep_.CoverTab[36499]++
//line /usr/local/go/src/net/http/internal/chunked.go:93
				return !cr.chunkHeaderAvailable()
//line /usr/local/go/src/net/http/internal/chunked.go:93
				// _ = "end of CoverTab[36499]"
//line /usr/local/go/src/net/http/internal/chunked.go:93
			}() {
//line /usr/local/go/src/net/http/internal/chunked.go:93
				_go_fuzz_dep_.CoverTab[36500]++

//line /usr/local/go/src/net/http/internal/chunked.go:96
				break
//line /usr/local/go/src/net/http/internal/chunked.go:96
				// _ = "end of CoverTab[36500]"
			} else {
//line /usr/local/go/src/net/http/internal/chunked.go:97
				_go_fuzz_dep_.CoverTab[36501]++
//line /usr/local/go/src/net/http/internal/chunked.go:97
				// _ = "end of CoverTab[36501]"
//line /usr/local/go/src/net/http/internal/chunked.go:97
			}
//line /usr/local/go/src/net/http/internal/chunked.go:97
			// _ = "end of CoverTab[36497]"
//line /usr/local/go/src/net/http/internal/chunked.go:97
			_go_fuzz_dep_.CoverTab[36498]++
										cr.beginChunk()
										continue
//line /usr/local/go/src/net/http/internal/chunked.go:99
			// _ = "end of CoverTab[36498]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:100
			_go_fuzz_dep_.CoverTab[36502]++
//line /usr/local/go/src/net/http/internal/chunked.go:100
			// _ = "end of CoverTab[36502]"
//line /usr/local/go/src/net/http/internal/chunked.go:100
		}
//line /usr/local/go/src/net/http/internal/chunked.go:100
		// _ = "end of CoverTab[36479]"
//line /usr/local/go/src/net/http/internal/chunked.go:100
		_go_fuzz_dep_.CoverTab[36480]++
									if len(b) == 0 {
//line /usr/local/go/src/net/http/internal/chunked.go:101
			_go_fuzz_dep_.CoverTab[36503]++
										break
//line /usr/local/go/src/net/http/internal/chunked.go:102
			// _ = "end of CoverTab[36503]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:103
			_go_fuzz_dep_.CoverTab[36504]++
//line /usr/local/go/src/net/http/internal/chunked.go:103
			// _ = "end of CoverTab[36504]"
//line /usr/local/go/src/net/http/internal/chunked.go:103
		}
//line /usr/local/go/src/net/http/internal/chunked.go:103
		// _ = "end of CoverTab[36480]"
//line /usr/local/go/src/net/http/internal/chunked.go:103
		_go_fuzz_dep_.CoverTab[36481]++
									rbuf := b
									if uint64(len(rbuf)) > cr.n {
//line /usr/local/go/src/net/http/internal/chunked.go:105
			_go_fuzz_dep_.CoverTab[36505]++
										rbuf = rbuf[:cr.n]
//line /usr/local/go/src/net/http/internal/chunked.go:106
			// _ = "end of CoverTab[36505]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:107
			_go_fuzz_dep_.CoverTab[36506]++
//line /usr/local/go/src/net/http/internal/chunked.go:107
			// _ = "end of CoverTab[36506]"
//line /usr/local/go/src/net/http/internal/chunked.go:107
		}
//line /usr/local/go/src/net/http/internal/chunked.go:107
		// _ = "end of CoverTab[36481]"
//line /usr/local/go/src/net/http/internal/chunked.go:107
		_go_fuzz_dep_.CoverTab[36482]++
									var n0 int
									n0, cr.err = cr.r.Read(rbuf)
									n += n0
									b = b[n0:]
									cr.n -= uint64(n0)

//line /usr/local/go/src/net/http/internal/chunked.go:115
		if cr.n == 0 && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:115
			_go_fuzz_dep_.CoverTab[36507]++
//line /usr/local/go/src/net/http/internal/chunked.go:115
			return cr.err == nil
//line /usr/local/go/src/net/http/internal/chunked.go:115
			// _ = "end of CoverTab[36507]"
//line /usr/local/go/src/net/http/internal/chunked.go:115
		}() {
//line /usr/local/go/src/net/http/internal/chunked.go:115
			_go_fuzz_dep_.CoverTab[36508]++
										cr.checkEnd = true
//line /usr/local/go/src/net/http/internal/chunked.go:116
			// _ = "end of CoverTab[36508]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:117
			_go_fuzz_dep_.CoverTab[36509]++
//line /usr/local/go/src/net/http/internal/chunked.go:117
			if cr.err == io.EOF {
//line /usr/local/go/src/net/http/internal/chunked.go:117
				_go_fuzz_dep_.CoverTab[36510]++
											cr.err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/internal/chunked.go:118
				// _ = "end of CoverTab[36510]"
			} else {
//line /usr/local/go/src/net/http/internal/chunked.go:119
				_go_fuzz_dep_.CoverTab[36511]++
//line /usr/local/go/src/net/http/internal/chunked.go:119
				// _ = "end of CoverTab[36511]"
//line /usr/local/go/src/net/http/internal/chunked.go:119
			}
//line /usr/local/go/src/net/http/internal/chunked.go:119
			// _ = "end of CoverTab[36509]"
//line /usr/local/go/src/net/http/internal/chunked.go:119
		}
//line /usr/local/go/src/net/http/internal/chunked.go:119
		// _ = "end of CoverTab[36482]"
	}
//line /usr/local/go/src/net/http/internal/chunked.go:120
	// _ = "end of CoverTab[36476]"
//line /usr/local/go/src/net/http/internal/chunked.go:120
	_go_fuzz_dep_.CoverTab[36477]++
								return n, cr.err
//line /usr/local/go/src/net/http/internal/chunked.go:121
	// _ = "end of CoverTab[36477]"
}

// Read a line of bytes (up to \n) from b.
//line /usr/local/go/src/net/http/internal/chunked.go:124
// Give up if the line exceeds maxLineLength.
//line /usr/local/go/src/net/http/internal/chunked.go:124
// The returned bytes are owned by the bufio.Reader
//line /usr/local/go/src/net/http/internal/chunked.go:124
// so they are only valid until the next bufio read.
//line /usr/local/go/src/net/http/internal/chunked.go:128
func readChunkLine(b *bufio.Reader) ([]byte, error) {
//line /usr/local/go/src/net/http/internal/chunked.go:128
	_go_fuzz_dep_.CoverTab[36512]++
								p, err := b.ReadSlice('\n')
								if err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:130
		_go_fuzz_dep_.CoverTab[36516]++

//line /usr/local/go/src/net/http/internal/chunked.go:133
		if err == io.EOF {
//line /usr/local/go/src/net/http/internal/chunked.go:133
			_go_fuzz_dep_.CoverTab[36518]++
										err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/internal/chunked.go:134
			// _ = "end of CoverTab[36518]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:135
			_go_fuzz_dep_.CoverTab[36519]++
//line /usr/local/go/src/net/http/internal/chunked.go:135
			if err == bufio.ErrBufferFull {
//line /usr/local/go/src/net/http/internal/chunked.go:135
				_go_fuzz_dep_.CoverTab[36520]++
											err = ErrLineTooLong
//line /usr/local/go/src/net/http/internal/chunked.go:136
				// _ = "end of CoverTab[36520]"
			} else {
//line /usr/local/go/src/net/http/internal/chunked.go:137
				_go_fuzz_dep_.CoverTab[36521]++
//line /usr/local/go/src/net/http/internal/chunked.go:137
				// _ = "end of CoverTab[36521]"
//line /usr/local/go/src/net/http/internal/chunked.go:137
			}
//line /usr/local/go/src/net/http/internal/chunked.go:137
			// _ = "end of CoverTab[36519]"
//line /usr/local/go/src/net/http/internal/chunked.go:137
		}
//line /usr/local/go/src/net/http/internal/chunked.go:137
		// _ = "end of CoverTab[36516]"
//line /usr/local/go/src/net/http/internal/chunked.go:137
		_go_fuzz_dep_.CoverTab[36517]++
									return nil, err
//line /usr/local/go/src/net/http/internal/chunked.go:138
		// _ = "end of CoverTab[36517]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:139
		_go_fuzz_dep_.CoverTab[36522]++
//line /usr/local/go/src/net/http/internal/chunked.go:139
		// _ = "end of CoverTab[36522]"
//line /usr/local/go/src/net/http/internal/chunked.go:139
	}
//line /usr/local/go/src/net/http/internal/chunked.go:139
	// _ = "end of CoverTab[36512]"
//line /usr/local/go/src/net/http/internal/chunked.go:139
	_go_fuzz_dep_.CoverTab[36513]++
								if len(p) >= maxLineLength {
//line /usr/local/go/src/net/http/internal/chunked.go:140
		_go_fuzz_dep_.CoverTab[36523]++
									return nil, ErrLineTooLong
//line /usr/local/go/src/net/http/internal/chunked.go:141
		// _ = "end of CoverTab[36523]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:142
		_go_fuzz_dep_.CoverTab[36524]++
//line /usr/local/go/src/net/http/internal/chunked.go:142
		// _ = "end of CoverTab[36524]"
//line /usr/local/go/src/net/http/internal/chunked.go:142
	}
//line /usr/local/go/src/net/http/internal/chunked.go:142
	// _ = "end of CoverTab[36513]"
//line /usr/local/go/src/net/http/internal/chunked.go:142
	_go_fuzz_dep_.CoverTab[36514]++
								p = trimTrailingWhitespace(p)
								p, err = removeChunkExtension(p)
								if err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:145
		_go_fuzz_dep_.CoverTab[36525]++
									return nil, err
//line /usr/local/go/src/net/http/internal/chunked.go:146
		// _ = "end of CoverTab[36525]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:147
		_go_fuzz_dep_.CoverTab[36526]++
//line /usr/local/go/src/net/http/internal/chunked.go:147
		// _ = "end of CoverTab[36526]"
//line /usr/local/go/src/net/http/internal/chunked.go:147
	}
//line /usr/local/go/src/net/http/internal/chunked.go:147
	// _ = "end of CoverTab[36514]"
//line /usr/local/go/src/net/http/internal/chunked.go:147
	_go_fuzz_dep_.CoverTab[36515]++
								return p, nil
//line /usr/local/go/src/net/http/internal/chunked.go:148
	// _ = "end of CoverTab[36515]"
}

func trimTrailingWhitespace(b []byte) []byte {
//line /usr/local/go/src/net/http/internal/chunked.go:151
	_go_fuzz_dep_.CoverTab[36527]++
								for len(b) > 0 && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:152
		_go_fuzz_dep_.CoverTab[36529]++
//line /usr/local/go/src/net/http/internal/chunked.go:152
		return isASCIISpace(b[len(b)-1])
//line /usr/local/go/src/net/http/internal/chunked.go:152
		// _ = "end of CoverTab[36529]"
//line /usr/local/go/src/net/http/internal/chunked.go:152
	}() {
//line /usr/local/go/src/net/http/internal/chunked.go:152
		_go_fuzz_dep_.CoverTab[36530]++
									b = b[:len(b)-1]
//line /usr/local/go/src/net/http/internal/chunked.go:153
		// _ = "end of CoverTab[36530]"
	}
//line /usr/local/go/src/net/http/internal/chunked.go:154
	// _ = "end of CoverTab[36527]"
//line /usr/local/go/src/net/http/internal/chunked.go:154
	_go_fuzz_dep_.CoverTab[36528]++
								return b
//line /usr/local/go/src/net/http/internal/chunked.go:155
	// _ = "end of CoverTab[36528]"
}

func isASCIISpace(b byte) bool {
//line /usr/local/go/src/net/http/internal/chunked.go:158
	_go_fuzz_dep_.CoverTab[36531]++
								return b == ' ' || func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:159
		_go_fuzz_dep_.CoverTab[36532]++
//line /usr/local/go/src/net/http/internal/chunked.go:159
		return b == '\t'
//line /usr/local/go/src/net/http/internal/chunked.go:159
		// _ = "end of CoverTab[36532]"
//line /usr/local/go/src/net/http/internal/chunked.go:159
	}() || func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:159
		_go_fuzz_dep_.CoverTab[36533]++
//line /usr/local/go/src/net/http/internal/chunked.go:159
		return b == '\n'
//line /usr/local/go/src/net/http/internal/chunked.go:159
		// _ = "end of CoverTab[36533]"
//line /usr/local/go/src/net/http/internal/chunked.go:159
	}() || func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:159
		_go_fuzz_dep_.CoverTab[36534]++
//line /usr/local/go/src/net/http/internal/chunked.go:159
		return b == '\r'
//line /usr/local/go/src/net/http/internal/chunked.go:159
		// _ = "end of CoverTab[36534]"
//line /usr/local/go/src/net/http/internal/chunked.go:159
	}()
//line /usr/local/go/src/net/http/internal/chunked.go:159
	// _ = "end of CoverTab[36531]"
}

var semi = []byte(";")

// removeChunkExtension removes any chunk-extension from p.
//line /usr/local/go/src/net/http/internal/chunked.go:164
// For example,
//line /usr/local/go/src/net/http/internal/chunked.go:164
//
//line /usr/local/go/src/net/http/internal/chunked.go:164
//	"0" => "0"
//line /usr/local/go/src/net/http/internal/chunked.go:164
//	"0;token" => "0"
//line /usr/local/go/src/net/http/internal/chunked.go:164
//	"0;token=val" => "0"
//line /usr/local/go/src/net/http/internal/chunked.go:164
//	`0;token="quoted string"` => "0"
//line /usr/local/go/src/net/http/internal/chunked.go:171
func removeChunkExtension(p []byte) ([]byte, error) {
//line /usr/local/go/src/net/http/internal/chunked.go:171
	_go_fuzz_dep_.CoverTab[36535]++
								p, _, _ = bytes.Cut(p, semi)

//line /usr/local/go/src/net/http/internal/chunked.go:176
	return p, nil
//line /usr/local/go/src/net/http/internal/chunked.go:176
	// _ = "end of CoverTab[36535]"
}

// NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
//line /usr/local/go/src/net/http/internal/chunked.go:179
// "chunked" format before writing them to w. Closing the returned chunkedWriter
//line /usr/local/go/src/net/http/internal/chunked.go:179
// sends the final 0-length chunk that marks the end of the stream but does
//line /usr/local/go/src/net/http/internal/chunked.go:179
// not send the final CRLF that appears after trailers; trailers and the last
//line /usr/local/go/src/net/http/internal/chunked.go:179
// CRLF must be written separately.
//line /usr/local/go/src/net/http/internal/chunked.go:179
//
//line /usr/local/go/src/net/http/internal/chunked.go:179
// NewChunkedWriter is not needed by normal applications. The http
//line /usr/local/go/src/net/http/internal/chunked.go:179
// package adds chunking automatically if handlers don't set a
//line /usr/local/go/src/net/http/internal/chunked.go:179
// Content-Length header. Using newChunkedWriter inside a handler
//line /usr/local/go/src/net/http/internal/chunked.go:179
// would result in double chunking or chunking with a Content-Length
//line /usr/local/go/src/net/http/internal/chunked.go:179
// length, both of which are wrong.
//line /usr/local/go/src/net/http/internal/chunked.go:190
func NewChunkedWriter(w io.Writer) io.WriteCloser {
//line /usr/local/go/src/net/http/internal/chunked.go:190
	_go_fuzz_dep_.CoverTab[36536]++
								return &chunkedWriter{w}
//line /usr/local/go/src/net/http/internal/chunked.go:191
	// _ = "end of CoverTab[36536]"
}

// Writing to chunkedWriter translates to writing in HTTP chunked Transfer
//line /usr/local/go/src/net/http/internal/chunked.go:194
// Encoding wire format to the underlying Wire chunkedWriter.
//line /usr/local/go/src/net/http/internal/chunked.go:196
type chunkedWriter struct {
	Wire io.Writer
}

// Write the contents of data as one chunk to Wire.
//line /usr/local/go/src/net/http/internal/chunked.go:200
// NOTE: Note that the corresponding chunk-writing procedure in Conn.Write has
//line /usr/local/go/src/net/http/internal/chunked.go:200
// a bug since it does not check for success of io.WriteString
//line /usr/local/go/src/net/http/internal/chunked.go:203
func (cw *chunkedWriter) Write(data []byte) (n int, err error) {
//line /usr/local/go/src/net/http/internal/chunked.go:203
	_go_fuzz_dep_.CoverTab[36537]++

//line /usr/local/go/src/net/http/internal/chunked.go:206
	if len(data) == 0 {
//line /usr/local/go/src/net/http/internal/chunked.go:206
		_go_fuzz_dep_.CoverTab[36544]++
									return 0, nil
//line /usr/local/go/src/net/http/internal/chunked.go:207
		// _ = "end of CoverTab[36544]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:208
		_go_fuzz_dep_.CoverTab[36545]++
//line /usr/local/go/src/net/http/internal/chunked.go:208
		// _ = "end of CoverTab[36545]"
//line /usr/local/go/src/net/http/internal/chunked.go:208
	}
//line /usr/local/go/src/net/http/internal/chunked.go:208
	// _ = "end of CoverTab[36537]"
//line /usr/local/go/src/net/http/internal/chunked.go:208
	_go_fuzz_dep_.CoverTab[36538]++

								if _, err = fmt.Fprintf(cw.Wire, "%x\r\n", len(data)); err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:210
		_go_fuzz_dep_.CoverTab[36546]++
									return 0, err
//line /usr/local/go/src/net/http/internal/chunked.go:211
		// _ = "end of CoverTab[36546]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:212
		_go_fuzz_dep_.CoverTab[36547]++
//line /usr/local/go/src/net/http/internal/chunked.go:212
		// _ = "end of CoverTab[36547]"
//line /usr/local/go/src/net/http/internal/chunked.go:212
	}
//line /usr/local/go/src/net/http/internal/chunked.go:212
	// _ = "end of CoverTab[36538]"
//line /usr/local/go/src/net/http/internal/chunked.go:212
	_go_fuzz_dep_.CoverTab[36539]++
								if n, err = cw.Wire.Write(data); err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:213
		_go_fuzz_dep_.CoverTab[36548]++
									return
//line /usr/local/go/src/net/http/internal/chunked.go:214
		// _ = "end of CoverTab[36548]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:215
		_go_fuzz_dep_.CoverTab[36549]++
//line /usr/local/go/src/net/http/internal/chunked.go:215
		// _ = "end of CoverTab[36549]"
//line /usr/local/go/src/net/http/internal/chunked.go:215
	}
//line /usr/local/go/src/net/http/internal/chunked.go:215
	// _ = "end of CoverTab[36539]"
//line /usr/local/go/src/net/http/internal/chunked.go:215
	_go_fuzz_dep_.CoverTab[36540]++
								if n != len(data) {
//line /usr/local/go/src/net/http/internal/chunked.go:216
		_go_fuzz_dep_.CoverTab[36550]++
									err = io.ErrShortWrite
									return
//line /usr/local/go/src/net/http/internal/chunked.go:218
		// _ = "end of CoverTab[36550]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:219
		_go_fuzz_dep_.CoverTab[36551]++
//line /usr/local/go/src/net/http/internal/chunked.go:219
		// _ = "end of CoverTab[36551]"
//line /usr/local/go/src/net/http/internal/chunked.go:219
	}
//line /usr/local/go/src/net/http/internal/chunked.go:219
	// _ = "end of CoverTab[36540]"
//line /usr/local/go/src/net/http/internal/chunked.go:219
	_go_fuzz_dep_.CoverTab[36541]++
								if _, err = io.WriteString(cw.Wire, "\r\n"); err != nil {
//line /usr/local/go/src/net/http/internal/chunked.go:220
		_go_fuzz_dep_.CoverTab[36552]++
									return
//line /usr/local/go/src/net/http/internal/chunked.go:221
		// _ = "end of CoverTab[36552]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:222
		_go_fuzz_dep_.CoverTab[36553]++
//line /usr/local/go/src/net/http/internal/chunked.go:222
		// _ = "end of CoverTab[36553]"
//line /usr/local/go/src/net/http/internal/chunked.go:222
	}
//line /usr/local/go/src/net/http/internal/chunked.go:222
	// _ = "end of CoverTab[36541]"
//line /usr/local/go/src/net/http/internal/chunked.go:222
	_go_fuzz_dep_.CoverTab[36542]++
								if bw, ok := cw.Wire.(*FlushAfterChunkWriter); ok {
//line /usr/local/go/src/net/http/internal/chunked.go:223
		_go_fuzz_dep_.CoverTab[36554]++
									err = bw.Flush()
//line /usr/local/go/src/net/http/internal/chunked.go:224
		// _ = "end of CoverTab[36554]"
	} else {
//line /usr/local/go/src/net/http/internal/chunked.go:225
		_go_fuzz_dep_.CoverTab[36555]++
//line /usr/local/go/src/net/http/internal/chunked.go:225
		// _ = "end of CoverTab[36555]"
//line /usr/local/go/src/net/http/internal/chunked.go:225
	}
//line /usr/local/go/src/net/http/internal/chunked.go:225
	// _ = "end of CoverTab[36542]"
//line /usr/local/go/src/net/http/internal/chunked.go:225
	_go_fuzz_dep_.CoverTab[36543]++
								return
//line /usr/local/go/src/net/http/internal/chunked.go:226
	// _ = "end of CoverTab[36543]"
}

func (cw *chunkedWriter) Close() error {
//line /usr/local/go/src/net/http/internal/chunked.go:229
	_go_fuzz_dep_.CoverTab[36556]++
								_, err := io.WriteString(cw.Wire, "0\r\n")
								return err
//line /usr/local/go/src/net/http/internal/chunked.go:231
	// _ = "end of CoverTab[36556]"
}

// FlushAfterChunkWriter signals from the caller of NewChunkedWriter
//line /usr/local/go/src/net/http/internal/chunked.go:234
// that each chunk should be followed by a flush. It is used by the
//line /usr/local/go/src/net/http/internal/chunked.go:234
// http.Transport code to keep the buffering behavior for headers and
//line /usr/local/go/src/net/http/internal/chunked.go:234
// trailers, but flush out chunks aggressively in the middle for
//line /usr/local/go/src/net/http/internal/chunked.go:234
// request bodies which may be generated slowly. See Issue 6574.
//line /usr/local/go/src/net/http/internal/chunked.go:239
type FlushAfterChunkWriter struct {
	*bufio.Writer
}

func parseHexUint(v []byte) (n uint64, err error) {
//line /usr/local/go/src/net/http/internal/chunked.go:243
	_go_fuzz_dep_.CoverTab[36557]++
								for i, b := range v {
//line /usr/local/go/src/net/http/internal/chunked.go:244
		_go_fuzz_dep_.CoverTab[36559]++
									switch {
		case '0' <= b && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:246
			_go_fuzz_dep_.CoverTab[36566]++
//line /usr/local/go/src/net/http/internal/chunked.go:246
			return b <= '9'
//line /usr/local/go/src/net/http/internal/chunked.go:246
			// _ = "end of CoverTab[36566]"
//line /usr/local/go/src/net/http/internal/chunked.go:246
		}():
//line /usr/local/go/src/net/http/internal/chunked.go:246
			_go_fuzz_dep_.CoverTab[36562]++
										b = b - '0'
//line /usr/local/go/src/net/http/internal/chunked.go:247
			// _ = "end of CoverTab[36562]"
		case 'a' <= b && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:248
			_go_fuzz_dep_.CoverTab[36567]++
//line /usr/local/go/src/net/http/internal/chunked.go:248
			return b <= 'f'
//line /usr/local/go/src/net/http/internal/chunked.go:248
			// _ = "end of CoverTab[36567]"
//line /usr/local/go/src/net/http/internal/chunked.go:248
		}():
//line /usr/local/go/src/net/http/internal/chunked.go:248
			_go_fuzz_dep_.CoverTab[36563]++
										b = b - 'a' + 10
//line /usr/local/go/src/net/http/internal/chunked.go:249
			// _ = "end of CoverTab[36563]"
		case 'A' <= b && func() bool {
//line /usr/local/go/src/net/http/internal/chunked.go:250
			_go_fuzz_dep_.CoverTab[36568]++
//line /usr/local/go/src/net/http/internal/chunked.go:250
			return b <= 'F'
//line /usr/local/go/src/net/http/internal/chunked.go:250
			// _ = "end of CoverTab[36568]"
//line /usr/local/go/src/net/http/internal/chunked.go:250
		}():
//line /usr/local/go/src/net/http/internal/chunked.go:250
			_go_fuzz_dep_.CoverTab[36564]++
										b = b - 'A' + 10
//line /usr/local/go/src/net/http/internal/chunked.go:251
			// _ = "end of CoverTab[36564]"
		default:
//line /usr/local/go/src/net/http/internal/chunked.go:252
			_go_fuzz_dep_.CoverTab[36565]++
										return 0, errors.New("invalid byte in chunk length")
//line /usr/local/go/src/net/http/internal/chunked.go:253
			// _ = "end of CoverTab[36565]"
		}
//line /usr/local/go/src/net/http/internal/chunked.go:254
		// _ = "end of CoverTab[36559]"
//line /usr/local/go/src/net/http/internal/chunked.go:254
		_go_fuzz_dep_.CoverTab[36560]++
									if i == 16 {
//line /usr/local/go/src/net/http/internal/chunked.go:255
			_go_fuzz_dep_.CoverTab[36569]++
										return 0, errors.New("http chunk length too large")
//line /usr/local/go/src/net/http/internal/chunked.go:256
			// _ = "end of CoverTab[36569]"
		} else {
//line /usr/local/go/src/net/http/internal/chunked.go:257
			_go_fuzz_dep_.CoverTab[36570]++
//line /usr/local/go/src/net/http/internal/chunked.go:257
			// _ = "end of CoverTab[36570]"
//line /usr/local/go/src/net/http/internal/chunked.go:257
		}
//line /usr/local/go/src/net/http/internal/chunked.go:257
		// _ = "end of CoverTab[36560]"
//line /usr/local/go/src/net/http/internal/chunked.go:257
		_go_fuzz_dep_.CoverTab[36561]++
									n <<= 4
									n |= uint64(b)
//line /usr/local/go/src/net/http/internal/chunked.go:259
		// _ = "end of CoverTab[36561]"
	}
//line /usr/local/go/src/net/http/internal/chunked.go:260
	// _ = "end of CoverTab[36557]"
//line /usr/local/go/src/net/http/internal/chunked.go:260
	_go_fuzz_dep_.CoverTab[36558]++
								return
//line /usr/local/go/src/net/http/internal/chunked.go:261
	// _ = "end of CoverTab[36558]"
}

//line /usr/local/go/src/net/http/internal/chunked.go:262
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/internal/chunked.go:262
var _ = _go_fuzz_dep_.CoverTab
