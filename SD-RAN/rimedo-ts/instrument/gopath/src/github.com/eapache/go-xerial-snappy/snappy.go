//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
package snappy

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:1
)

import (
	"bytes"
	"encoding/binary"
	"errors"

	master "github.com/golang/snappy"
)

const (
	sizeOffset	= 16
	sizeBytes	= 4
)

var (
	xerialHeader	= []byte{130, 83, 78, 65, 80, 80, 89, 0}

	// This is xerial version 1 and minimally compatible with version 1
	xerialVersionInfo	= []byte{0, 0, 0, 1, 0, 0, 0, 1}

	// ErrMalformed is returned by the decoder when the xerial framing
	// is malformed
	ErrMalformed	= errors.New("malformed xerial framing")
)

func min(x, y int) int {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:27
	_go_fuzz_dep_.CoverTab[82229]++
															if x < y {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:28
		_go_fuzz_dep_.CoverTab[82231]++
																return x
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:29
		// _ = "end of CoverTab[82231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:30
		_go_fuzz_dep_.CoverTab[82232]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:30
		// _ = "end of CoverTab[82232]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:30
	// _ = "end of CoverTab[82229]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:30
	_go_fuzz_dep_.CoverTab[82230]++
															return y
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:31
	// _ = "end of CoverTab[82230]"
}

// Encode encodes data as snappy with no framing header.
func Encode(src []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:35
	_go_fuzz_dep_.CoverTab[82233]++
															return master.Encode(nil, src)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:36
	// _ = "end of CoverTab[82233]"
}

// EncodeStream *appends* to the specified 'dst' the compressed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:39
// 'src' in xerial framing format. If 'dst' does not have enough
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:39
// capacity, then a new slice will be allocated. If 'dst' has
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:39
// non-zero length, then if *must* have been built using this function.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:43
func EncodeStream(dst, src []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:43
	_go_fuzz_dep_.CoverTab[82234]++
															if len(dst) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:44
		_go_fuzz_dep_.CoverTab[82237]++
																dst = append(dst, xerialHeader...)
																dst = append(dst, xerialVersionInfo...)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:46
		// _ = "end of CoverTab[82237]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:47
		_go_fuzz_dep_.CoverTab[82238]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:47
		// _ = "end of CoverTab[82238]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:47
	// _ = "end of CoverTab[82234]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:47
	_go_fuzz_dep_.CoverTab[82235]++

	// Snappy encode in blocks of maximum 32KB
	var (
		max		= len(src)
		blockSize	= 32 * 1024
		pos		= 0
		chunk		[]byte
	)

	for pos < max {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:57
		_go_fuzz_dep_.CoverTab[82239]++
																newPos := min(pos+blockSize, max)
																chunk = master.Encode(chunk[:cap(chunk)], src[pos:newPos])

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:63
		origLen := len(dst)
																dst = append(dst, dst[0:4]...)
																binary.BigEndian.PutUint32(dst[origLen:], uint32(len(chunk)))

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:68
		dst = append(dst, chunk...)
																pos = newPos
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:69
		// _ = "end of CoverTab[82239]"
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:70
	// _ = "end of CoverTab[82235]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:70
	_go_fuzz_dep_.CoverTab[82236]++
															return dst
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:71
	// _ = "end of CoverTab[82236]"
}

// Decode decodes snappy data whether it is traditional unframed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:74
// or includes the xerial framing format.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:76
func Decode(src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:76
	_go_fuzz_dep_.CoverTab[82240]++
															return DecodeInto(nil, src)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:77
	// _ = "end of CoverTab[82240]"
}

// DecodeInto decodes snappy data whether it is traditional unframed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:80
// or includes the xerial framing format into the specified `dst`.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:80
// It is assumed that the entirety of `dst` including all capacity is available
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:80
// for use by this function. If `dst` is nil *or* insufficiently large to hold
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:80
// the decoded `src`, new space will be allocated.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:85
func DecodeInto(dst, src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:85
	_go_fuzz_dep_.CoverTab[82241]++
															var max = len(src)
															if max < len(xerialHeader) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:87
		_go_fuzz_dep_.CoverTab[82247]++
																return nil, ErrMalformed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:88
		// _ = "end of CoverTab[82247]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:89
		_go_fuzz_dep_.CoverTab[82248]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:89
		// _ = "end of CoverTab[82248]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:89
	// _ = "end of CoverTab[82241]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:89
	_go_fuzz_dep_.CoverTab[82242]++

															if !bytes.Equal(src[:8], xerialHeader) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:91
		_go_fuzz_dep_.CoverTab[82249]++
																return master.Decode(dst[:cap(dst)], src)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:92
		// _ = "end of CoverTab[82249]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:93
		_go_fuzz_dep_.CoverTab[82250]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:93
		// _ = "end of CoverTab[82250]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:93
	// _ = "end of CoverTab[82242]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:93
	_go_fuzz_dep_.CoverTab[82243]++

															if max < sizeOffset+sizeBytes {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:95
		_go_fuzz_dep_.CoverTab[82251]++
																return nil, ErrMalformed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:96
		// _ = "end of CoverTab[82251]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:97
		_go_fuzz_dep_.CoverTab[82252]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:97
		// _ = "end of CoverTab[82252]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:97
	// _ = "end of CoverTab[82243]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:97
	_go_fuzz_dep_.CoverTab[82244]++

															if dst == nil {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:99
		_go_fuzz_dep_.CoverTab[82253]++
																dst = make([]byte, 0, len(src))
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:100
		// _ = "end of CoverTab[82253]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:101
		_go_fuzz_dep_.CoverTab[82254]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:101
		// _ = "end of CoverTab[82254]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:101
	// _ = "end of CoverTab[82244]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:101
	_go_fuzz_dep_.CoverTab[82245]++

															dst = dst[:0]
															var (
		pos	= sizeOffset
		chunk	[]byte
		err	error
	)

	for pos+sizeBytes <= max {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:110
		_go_fuzz_dep_.CoverTab[82255]++
																size := int(binary.BigEndian.Uint32(src[pos : pos+sizeBytes]))
																pos += sizeBytes

																nextPos := pos + size

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:118
		if nextPos < pos || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:118
			_go_fuzz_dep_.CoverTab[82258]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:118
			return nextPos > max
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:118
			// _ = "end of CoverTab[82258]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:118
		}() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:118
			_go_fuzz_dep_.CoverTab[82259]++
																	return nil, ErrMalformed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:119
			// _ = "end of CoverTab[82259]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:120
			_go_fuzz_dep_.CoverTab[82260]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:120
			// _ = "end of CoverTab[82260]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:120
		}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:120
		// _ = "end of CoverTab[82255]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:120
		_go_fuzz_dep_.CoverTab[82256]++

																chunk, err = master.Decode(chunk[:cap(chunk)], src[pos:nextPos])

																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:124
			_go_fuzz_dep_.CoverTab[82261]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:125
			// _ = "end of CoverTab[82261]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:126
			_go_fuzz_dep_.CoverTab[82262]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:126
			// _ = "end of CoverTab[82262]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:126
		// _ = "end of CoverTab[82256]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:126
		_go_fuzz_dep_.CoverTab[82257]++
																pos = nextPos
																dst = append(dst, chunk...)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:128
		// _ = "end of CoverTab[82257]"
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:129
	// _ = "end of CoverTab[82245]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:129
	_go_fuzz_dep_.CoverTab[82246]++
															return dst, nil
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:130
	// _ = "end of CoverTab[82246]"
}

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:131
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-xerial-snappy@v0.0.0-20180814174437-776d5712da21/snappy.go:131
var _ = _go_fuzz_dep_.CoverTab
