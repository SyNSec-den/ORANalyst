// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:38
import (
	"errors"
	"fmt"
	"io"
)

// errOverflow is returned when an integer is too large to be represented.
var errOverflow = errors.New("proto: integer overflow")

// ErrInternalBadWireType is returned by generated code when an incorrect
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:47
// wire type is encountered. It does not get returned to user code.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:49
var ErrInternalBadWireType = errors.New("proto: internal error: bad wiretype for oneof")

// DecodeVarint reads a varint-encoded integer from the slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:51
// It returns the integer and the number of bytes consumed, or
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:51
// zero if there is not enough.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:51
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:51
// int32, int64, uint32, uint64, bool, and enum
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:51
// protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:57
func DecodeVarint(buf []byte) (x uint64, n int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:57
	_go_fuzz_dep_.CoverTab[107681]++
											for shift := uint(0); shift < 64; shift += 7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:58
		_go_fuzz_dep_.CoverTab[107683]++
												if n >= len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:59
			_go_fuzz_dep_.CoverTab[107685]++
													return 0, 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:60
			// _ = "end of CoverTab[107685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:61
			_go_fuzz_dep_.CoverTab[107686]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:61
			// _ = "end of CoverTab[107686]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:61
		// _ = "end of CoverTab[107683]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:61
		_go_fuzz_dep_.CoverTab[107684]++
												b := uint64(buf[n])
												n++
												x |= (b & 0x7F) << shift
												if (b & 0x80) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:65
			_go_fuzz_dep_.CoverTab[107687]++
													return x, n
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:66
			// _ = "end of CoverTab[107687]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:67
			_go_fuzz_dep_.CoverTab[107688]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:67
			// _ = "end of CoverTab[107688]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:67
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:67
		// _ = "end of CoverTab[107684]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:68
	// _ = "end of CoverTab[107681]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:68
	_go_fuzz_dep_.CoverTab[107682]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:71
	return 0, 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:71
	// _ = "end of CoverTab[107682]"
}

func (p *Buffer) decodeVarintSlow() (x uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:74
	_go_fuzz_dep_.CoverTab[107689]++
											i := p.index
											l := len(p.buf)

											for shift := uint(0); shift < 64; shift += 7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:78
		_go_fuzz_dep_.CoverTab[107691]++
												if i >= l {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:79
			_go_fuzz_dep_.CoverTab[107693]++
													err = io.ErrUnexpectedEOF
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:81
			// _ = "end of CoverTab[107693]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:82
			_go_fuzz_dep_.CoverTab[107694]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:82
			// _ = "end of CoverTab[107694]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:82
		// _ = "end of CoverTab[107691]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:82
		_go_fuzz_dep_.CoverTab[107692]++
												b := p.buf[i]
												i++
												x |= (uint64(b) & 0x7F) << shift
												if b < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:86
			_go_fuzz_dep_.CoverTab[107695]++
													p.index = i
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:88
			// _ = "end of CoverTab[107695]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:89
			_go_fuzz_dep_.CoverTab[107696]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:89
			// _ = "end of CoverTab[107696]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:89
		// _ = "end of CoverTab[107692]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:90
	// _ = "end of CoverTab[107689]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:90
	_go_fuzz_dep_.CoverTab[107690]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:93
	err = errOverflow
											return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:94
	// _ = "end of CoverTab[107690]"
}

// DecodeVarint reads a varint-encoded integer from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:97
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:97
// int32, int64, uint32, uint64, bool, and enum
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:97
// protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:101
func (p *Buffer) DecodeVarint() (x uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:101
	_go_fuzz_dep_.CoverTab[107697]++
											i := p.index
											buf := p.buf

											if i >= len(buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:105
		_go_fuzz_dep_.CoverTab[107708]++
												return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:106
		// _ = "end of CoverTab[107708]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:107
		_go_fuzz_dep_.CoverTab[107709]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:107
		if buf[i] < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:107
			_go_fuzz_dep_.CoverTab[107710]++
													p.index++
													return uint64(buf[i]), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:109
			// _ = "end of CoverTab[107710]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:110
			_go_fuzz_dep_.CoverTab[107711]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:110
			if len(buf)-i < 10 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:110
				_go_fuzz_dep_.CoverTab[107712]++
														return p.decodeVarintSlow()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:111
				// _ = "end of CoverTab[107712]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
				_go_fuzz_dep_.CoverTab[107713]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
				// _ = "end of CoverTab[107713]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
			// _ = "end of CoverTab[107711]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
		// _ = "end of CoverTab[107709]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
	// _ = "end of CoverTab[107697]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:112
	_go_fuzz_dep_.CoverTab[107698]++

											var b uint64

											x = uint64(buf[i]) - 0x80
											i++

											b = uint64(buf[i])
											i++
											x += b << 7
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:122
		_go_fuzz_dep_.CoverTab[107714]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:123
		// _ = "end of CoverTab[107714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:124
		_go_fuzz_dep_.CoverTab[107715]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:124
		// _ = "end of CoverTab[107715]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:124
	// _ = "end of CoverTab[107698]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:124
	_go_fuzz_dep_.CoverTab[107699]++
											x -= 0x80 << 7

											b = uint64(buf[i])
											i++
											x += b << 14
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:130
		_go_fuzz_dep_.CoverTab[107716]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:131
		// _ = "end of CoverTab[107716]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:132
		_go_fuzz_dep_.CoverTab[107717]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:132
		// _ = "end of CoverTab[107717]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:132
	// _ = "end of CoverTab[107699]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:132
	_go_fuzz_dep_.CoverTab[107700]++
											x -= 0x80 << 14

											b = uint64(buf[i])
											i++
											x += b << 21
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:138
		_go_fuzz_dep_.CoverTab[107718]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:139
		// _ = "end of CoverTab[107718]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:140
		_go_fuzz_dep_.CoverTab[107719]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:140
		// _ = "end of CoverTab[107719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:140
	// _ = "end of CoverTab[107700]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:140
	_go_fuzz_dep_.CoverTab[107701]++
											x -= 0x80 << 21

											b = uint64(buf[i])
											i++
											x += b << 28
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:146
		_go_fuzz_dep_.CoverTab[107720]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:147
		// _ = "end of CoverTab[107720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:148
		_go_fuzz_dep_.CoverTab[107721]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:148
		// _ = "end of CoverTab[107721]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:148
	// _ = "end of CoverTab[107701]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:148
	_go_fuzz_dep_.CoverTab[107702]++
											x -= 0x80 << 28

											b = uint64(buf[i])
											i++
											x += b << 35
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:154
		_go_fuzz_dep_.CoverTab[107722]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:155
		// _ = "end of CoverTab[107722]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:156
		_go_fuzz_dep_.CoverTab[107723]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:156
		// _ = "end of CoverTab[107723]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:156
	// _ = "end of CoverTab[107702]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:156
	_go_fuzz_dep_.CoverTab[107703]++
											x -= 0x80 << 35

											b = uint64(buf[i])
											i++
											x += b << 42
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:162
		_go_fuzz_dep_.CoverTab[107724]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:163
		// _ = "end of CoverTab[107724]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:164
		_go_fuzz_dep_.CoverTab[107725]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:164
		// _ = "end of CoverTab[107725]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:164
	// _ = "end of CoverTab[107703]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:164
	_go_fuzz_dep_.CoverTab[107704]++
											x -= 0x80 << 42

											b = uint64(buf[i])
											i++
											x += b << 49
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:170
		_go_fuzz_dep_.CoverTab[107726]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:171
		// _ = "end of CoverTab[107726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:172
		_go_fuzz_dep_.CoverTab[107727]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:172
		// _ = "end of CoverTab[107727]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:172
	// _ = "end of CoverTab[107704]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:172
	_go_fuzz_dep_.CoverTab[107705]++
											x -= 0x80 << 49

											b = uint64(buf[i])
											i++
											x += b << 56
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:178
		_go_fuzz_dep_.CoverTab[107728]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:179
		// _ = "end of CoverTab[107728]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:180
		_go_fuzz_dep_.CoverTab[107729]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:180
		// _ = "end of CoverTab[107729]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:180
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:180
	// _ = "end of CoverTab[107705]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:180
	_go_fuzz_dep_.CoverTab[107706]++
											x -= 0x80 << 56

											b = uint64(buf[i])
											i++
											x += b << 63
											if b&0x80 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:186
		_go_fuzz_dep_.CoverTab[107730]++
												goto done
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:187
		// _ = "end of CoverTab[107730]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:188
		_go_fuzz_dep_.CoverTab[107731]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:188
		// _ = "end of CoverTab[107731]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:188
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:188
	// _ = "end of CoverTab[107706]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:188
	_go_fuzz_dep_.CoverTab[107707]++

											return 0, errOverflow

done:
											p.index = i
											return x, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:194
	// _ = "end of CoverTab[107707]"
}

// DecodeFixed64 reads a 64-bit integer from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:197
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:197
// fixed64, sfixed64, and double protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:200
func (p *Buffer) DecodeFixed64() (x uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:200
	_go_fuzz_dep_.CoverTab[107732]++

											i := p.index + 8
											if i < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:203
		_go_fuzz_dep_.CoverTab[107734]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:203
		return i > len(p.buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:203
		// _ = "end of CoverTab[107734]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:203
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:203
		_go_fuzz_dep_.CoverTab[107735]++
												err = io.ErrUnexpectedEOF
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:205
		// _ = "end of CoverTab[107735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:206
		_go_fuzz_dep_.CoverTab[107736]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:206
		// _ = "end of CoverTab[107736]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:206
	// _ = "end of CoverTab[107732]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:206
	_go_fuzz_dep_.CoverTab[107733]++
											p.index = i

											x = uint64(p.buf[i-8])
											x |= uint64(p.buf[i-7]) << 8
											x |= uint64(p.buf[i-6]) << 16
											x |= uint64(p.buf[i-5]) << 24
											x |= uint64(p.buf[i-4]) << 32
											x |= uint64(p.buf[i-3]) << 40
											x |= uint64(p.buf[i-2]) << 48
											x |= uint64(p.buf[i-1]) << 56
											return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:217
	// _ = "end of CoverTab[107733]"
}

// DecodeFixed32 reads a 32-bit integer from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:220
// This is the format for the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:220
// fixed32, sfixed32, and float protocol buffer types.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:223
func (p *Buffer) DecodeFixed32() (x uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:223
	_go_fuzz_dep_.CoverTab[107737]++

											i := p.index + 4
											if i < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:226
		_go_fuzz_dep_.CoverTab[107739]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:226
		return i > len(p.buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:226
		// _ = "end of CoverTab[107739]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:226
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:226
		_go_fuzz_dep_.CoverTab[107740]++
												err = io.ErrUnexpectedEOF
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:228
		// _ = "end of CoverTab[107740]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:229
		_go_fuzz_dep_.CoverTab[107741]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:229
		// _ = "end of CoverTab[107741]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:229
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:229
	// _ = "end of CoverTab[107737]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:229
	_go_fuzz_dep_.CoverTab[107738]++
											p.index = i

											x = uint64(p.buf[i-4])
											x |= uint64(p.buf[i-3]) << 8
											x |= uint64(p.buf[i-2]) << 16
											x |= uint64(p.buf[i-1]) << 24
											return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:236
	// _ = "end of CoverTab[107738]"
}

// DecodeZigzag64 reads a zigzag-encoded 64-bit integer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:239
// from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:239
// This is the format used for the sint64 protocol buffer type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:242
func (p *Buffer) DecodeZigzag64() (x uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:242
	_go_fuzz_dep_.CoverTab[107742]++
											x, err = p.DecodeVarint()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:244
		_go_fuzz_dep_.CoverTab[107744]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:245
		// _ = "end of CoverTab[107744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:246
		_go_fuzz_dep_.CoverTab[107745]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:246
		// _ = "end of CoverTab[107745]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:246
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:246
	// _ = "end of CoverTab[107742]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:246
	_go_fuzz_dep_.CoverTab[107743]++
											x = (x >> 1) ^ uint64((int64(x&1)<<63)>>63)
											return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:248
	// _ = "end of CoverTab[107743]"
}

// DecodeZigzag32 reads a zigzag-encoded 32-bit integer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:251
// from  the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:251
// This is the format used for the sint32 protocol buffer type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:254
func (p *Buffer) DecodeZigzag32() (x uint64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:254
	_go_fuzz_dep_.CoverTab[107746]++
											x, err = p.DecodeVarint()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:256
		_go_fuzz_dep_.CoverTab[107748]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:257
		// _ = "end of CoverTab[107748]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:258
		_go_fuzz_dep_.CoverTab[107749]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:258
		// _ = "end of CoverTab[107749]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:258
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:258
	// _ = "end of CoverTab[107746]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:258
	_go_fuzz_dep_.CoverTab[107747]++
											x = uint64((uint32(x) >> 1) ^ uint32((int32(x&1)<<31)>>31))
											return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:260
	// _ = "end of CoverTab[107747]"
}

// DecodeRawBytes reads a count-delimited byte buffer from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:263
// This is the format used for the bytes protocol buffer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:263
// type and for embedded messages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:266
func (p *Buffer) DecodeRawBytes(alloc bool) (buf []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:266
	_go_fuzz_dep_.CoverTab[107750]++
											n, err := p.DecodeVarint()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:268
		_go_fuzz_dep_.CoverTab[107755]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:269
		// _ = "end of CoverTab[107755]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:270
		_go_fuzz_dep_.CoverTab[107756]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:270
		// _ = "end of CoverTab[107756]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:270
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:270
	// _ = "end of CoverTab[107750]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:270
	_go_fuzz_dep_.CoverTab[107751]++

											nb := int(n)
											if nb < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:273
		_go_fuzz_dep_.CoverTab[107757]++
												return nil, fmt.Errorf("proto: bad byte length %d", nb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:274
		// _ = "end of CoverTab[107757]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:275
		_go_fuzz_dep_.CoverTab[107758]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:275
		// _ = "end of CoverTab[107758]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:275
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:275
	// _ = "end of CoverTab[107751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:275
	_go_fuzz_dep_.CoverTab[107752]++
											end := p.index + nb
											if end < p.index || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:277
		_go_fuzz_dep_.CoverTab[107759]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:277
		return end > len(p.buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:277
		// _ = "end of CoverTab[107759]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:277
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:277
		_go_fuzz_dep_.CoverTab[107760]++
												return nil, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:278
		// _ = "end of CoverTab[107760]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:279
		_go_fuzz_dep_.CoverTab[107761]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:279
		// _ = "end of CoverTab[107761]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:279
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:279
	// _ = "end of CoverTab[107752]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:279
	_go_fuzz_dep_.CoverTab[107753]++

											if !alloc {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:281
		_go_fuzz_dep_.CoverTab[107762]++

												buf = p.buf[p.index:end]
												p.index += nb
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:285
		// _ = "end of CoverTab[107762]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:286
		_go_fuzz_dep_.CoverTab[107763]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:286
		// _ = "end of CoverTab[107763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:286
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:286
	// _ = "end of CoverTab[107753]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:286
	_go_fuzz_dep_.CoverTab[107754]++

											buf = make([]byte, nb)
											copy(buf, p.buf[p.index:])
											p.index += nb
											return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:291
	// _ = "end of CoverTab[107754]"
}

// DecodeStringBytes reads an encoded string from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:294
// This is the format used for the proto2 string type.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:296
func (p *Buffer) DecodeStringBytes() (s string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:296
	_go_fuzz_dep_.CoverTab[107764]++
											buf, err := p.DecodeRawBytes(false)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:298
		_go_fuzz_dep_.CoverTab[107766]++
												return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:299
		// _ = "end of CoverTab[107766]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:300
		_go_fuzz_dep_.CoverTab[107767]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:300
		// _ = "end of CoverTab[107767]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:300
	// _ = "end of CoverTab[107764]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:300
	_go_fuzz_dep_.CoverTab[107765]++
											return string(buf), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:301
	// _ = "end of CoverTab[107765]"
}

// Unmarshaler is the interface representing objects that can
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// unmarshal themselves.  The argument points to data that may be
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// overwritten, so implementations should not keep references to the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// Unmarshal implementations should not clear the receiver.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// Any unmarshaled data should be merged into the receiver.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// Callers of Unmarshal that do not want to retain existing data
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:304
// should Reset the receiver before calling Unmarshal.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:312
type Unmarshaler interface {
	Unmarshal([]byte) error
}

// newUnmarshaler is the interface representing objects that can
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:316
// unmarshal themselves. The semantics are identical to Unmarshaler.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:316
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:316
// This exists to support protoc-gen-go generated messages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:316
// The proto package will stop type-asserting to this interface in the future.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:316
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:316
// DO NOT DEPEND ON THIS.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:323
type newUnmarshaler interface {
	XXX_Unmarshal([]byte) error
}

// Unmarshal parses the protocol buffer representation in buf and places the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:327
// decoded result in pb.  If the struct underlying pb does not match
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:327
// the data in buf, the results can be unpredictable.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:327
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:327
// Unmarshal resets pb before starting to unmarshal, so any
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:327
// existing data in pb is always removed. Use UnmarshalMerge
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:327
// to preserve and append to existing data.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:334
func Unmarshal(buf []byte, pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:334
	_go_fuzz_dep_.CoverTab[107768]++
											pb.Reset()
											if u, ok := pb.(newUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:336
		_go_fuzz_dep_.CoverTab[107771]++
												return u.XXX_Unmarshal(buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:337
		// _ = "end of CoverTab[107771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:338
		_go_fuzz_dep_.CoverTab[107772]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:338
		// _ = "end of CoverTab[107772]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:338
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:338
	// _ = "end of CoverTab[107768]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:338
	_go_fuzz_dep_.CoverTab[107769]++
											if u, ok := pb.(Unmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:339
		_go_fuzz_dep_.CoverTab[107773]++
												return u.Unmarshal(buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:340
		// _ = "end of CoverTab[107773]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:341
		_go_fuzz_dep_.CoverTab[107774]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:341
		// _ = "end of CoverTab[107774]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:341
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:341
	// _ = "end of CoverTab[107769]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:341
	_go_fuzz_dep_.CoverTab[107770]++
											return NewBuffer(buf).Unmarshal(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:342
	// _ = "end of CoverTab[107770]"
}

// UnmarshalMerge parses the protocol buffer representation in buf and
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:345
// writes the decoded result to pb.  If the struct underlying pb does not match
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:345
// the data in buf, the results can be unpredictable.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:345
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:345
// UnmarshalMerge merges into existing data in pb.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:345
// Most code should use Unmarshal instead.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:351
func UnmarshalMerge(buf []byte, pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:351
	_go_fuzz_dep_.CoverTab[107775]++
											if u, ok := pb.(newUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:352
		_go_fuzz_dep_.CoverTab[107778]++
												return u.XXX_Unmarshal(buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:353
		// _ = "end of CoverTab[107778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:354
		_go_fuzz_dep_.CoverTab[107779]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:354
		// _ = "end of CoverTab[107779]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:354
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:354
	// _ = "end of CoverTab[107775]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:354
	_go_fuzz_dep_.CoverTab[107776]++
											if u, ok := pb.(Unmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:355
		_go_fuzz_dep_.CoverTab[107780]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:362
		return u.Unmarshal(buf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:362
		// _ = "end of CoverTab[107780]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:363
		_go_fuzz_dep_.CoverTab[107781]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:363
		// _ = "end of CoverTab[107781]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:363
	// _ = "end of CoverTab[107776]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:363
	_go_fuzz_dep_.CoverTab[107777]++
											return NewBuffer(buf).Unmarshal(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:364
	// _ = "end of CoverTab[107777]"
}

// DecodeMessage reads a count-delimited message from the Buffer.
func (p *Buffer) DecodeMessage(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:368
	_go_fuzz_dep_.CoverTab[107782]++
											enc, err := p.DecodeRawBytes(false)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:370
		_go_fuzz_dep_.CoverTab[107784]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:371
		// _ = "end of CoverTab[107784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:372
		_go_fuzz_dep_.CoverTab[107785]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:372
		// _ = "end of CoverTab[107785]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:372
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:372
	// _ = "end of CoverTab[107782]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:372
	_go_fuzz_dep_.CoverTab[107783]++
											return NewBuffer(enc).Unmarshal(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:373
	// _ = "end of CoverTab[107783]"
}

// DecodeGroup reads a tag-delimited group from the Buffer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:376
// StartGroup tag is already consumed. This function consumes
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:376
// EndGroup tag.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:379
func (p *Buffer) DecodeGroup(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:379
	_go_fuzz_dep_.CoverTab[107786]++
											b := p.buf[p.index:]
											x, y := findEndGroup(b)
											if x < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:382
		_go_fuzz_dep_.CoverTab[107788]++
												return io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:383
		// _ = "end of CoverTab[107788]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:384
		_go_fuzz_dep_.CoverTab[107789]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:384
		// _ = "end of CoverTab[107789]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:384
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:384
	// _ = "end of CoverTab[107786]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:384
	_go_fuzz_dep_.CoverTab[107787]++
											err := Unmarshal(b[:x], pb)
											p.index += y
											return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:387
	// _ = "end of CoverTab[107787]"
}

// Unmarshal parses the protocol buffer representation in the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:390
// Buffer and places the decoded result in pb.  If the struct
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:390
// underlying pb does not match the data in the buffer, the results can be
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:390
// unpredictable.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:390
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:390
// Unlike proto.Unmarshal, this does not reset pb before starting to unmarshal.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:396
func (p *Buffer) Unmarshal(pb Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:396
	_go_fuzz_dep_.CoverTab[107790]++

											if u, ok := pb.(newUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:398
		_go_fuzz_dep_.CoverTab[107793]++
												err := u.XXX_Unmarshal(p.buf[p.index:])
												p.index = len(p.buf)
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:401
		// _ = "end of CoverTab[107793]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:402
		_go_fuzz_dep_.CoverTab[107794]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:402
		// _ = "end of CoverTab[107794]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:402
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:402
	// _ = "end of CoverTab[107790]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:402
	_go_fuzz_dep_.CoverTab[107791]++
											if u, ok := pb.(Unmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:403
		_go_fuzz_dep_.CoverTab[107795]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:410
		err := u.Unmarshal(p.buf[p.index:])
												p.index = len(p.buf)
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:412
		// _ = "end of CoverTab[107795]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:413
		_go_fuzz_dep_.CoverTab[107796]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:413
		// _ = "end of CoverTab[107796]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:413
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:413
	// _ = "end of CoverTab[107791]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:413
	_go_fuzz_dep_.CoverTab[107792]++

	// Slow workaround for messages that aren't Unmarshalers.
	// This includes some hand-coded .pb.go files and
	// bootstrap protos.
	// TODO: fix all of those and then add Unmarshal to
	// the Message interface. Then:
	// The cast above and code below can be deleted.
	// The old unmarshaler can be deleted.
											// Clients can call Unmarshal directly (can already do that, actually).
											var info InternalMessageInfo
											err := info.Unmarshal(pb, p.buf[p.index:])
											p.index = len(p.buf)
											return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:426
	// _ = "end of CoverTab[107792]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:427
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/decode.go:427
var _ = _go_fuzz_dep_.CoverTab
