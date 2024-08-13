//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
package mstypes

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:1
)

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// Byte sizes of primitive types
const (
	SizeBool	= 1
	SizeChar	= 1
	SizeUint8	= 1
	SizeUint16	= 2
	SizeUint32	= 4
	SizeUint64	= 8
	SizeEnum	= 2
	SizeSingle	= 4
	SizeDouble	= 8
	SizePtr		= 4
)

// Reader reads simple byte stream data into a Go representations
type Reader struct {
	r *bufio.Reader	// source of the data
}

// NewReader creates a new instance of a simple Reader.
func NewReader(r io.Reader) *Reader {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:30
	_go_fuzz_dep_.CoverTab[87360]++
												reader := new(Reader)
												reader.r = bufio.NewReader(r)
												return reader
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:33
	// _ = "end of CoverTab[87360]"
}

func (r *Reader) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:36
	_go_fuzz_dep_.CoverTab[87361]++
												return r.r.Read(p)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:37
	// _ = "end of CoverTab[87361]"
}

func (r *Reader) Uint8() (uint8, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:40
	_go_fuzz_dep_.CoverTab[87362]++
												b, err := r.r.ReadByte()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:42
		_go_fuzz_dep_.CoverTab[87364]++
													return uint8(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:43
		// _ = "end of CoverTab[87364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:44
		_go_fuzz_dep_.CoverTab[87365]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:44
		// _ = "end of CoverTab[87365]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:44
	// _ = "end of CoverTab[87362]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:44
	_go_fuzz_dep_.CoverTab[87363]++
												return uint8(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:45
	// _ = "end of CoverTab[87363]"
}

func (r *Reader) Uint16() (uint16, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:48
	_go_fuzz_dep_.CoverTab[87366]++
												b, err := r.ReadBytes(SizeUint16)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:50
		_go_fuzz_dep_.CoverTab[87368]++
													return uint16(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:51
		// _ = "end of CoverTab[87368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:52
		_go_fuzz_dep_.CoverTab[87369]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:52
		// _ = "end of CoverTab[87369]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:52
	// _ = "end of CoverTab[87366]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:52
	_go_fuzz_dep_.CoverTab[87367]++
												return binary.LittleEndian.Uint16(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:53
	// _ = "end of CoverTab[87367]"
}

func (r *Reader) Uint32() (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:56
	_go_fuzz_dep_.CoverTab[87370]++
												b, err := r.ReadBytes(SizeUint32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:58
		_go_fuzz_dep_.CoverTab[87372]++
													return uint32(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:59
		// _ = "end of CoverTab[87372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:60
		_go_fuzz_dep_.CoverTab[87373]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:60
		// _ = "end of CoverTab[87373]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:60
	// _ = "end of CoverTab[87370]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:60
	_go_fuzz_dep_.CoverTab[87371]++
												return binary.LittleEndian.Uint32(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:61
	// _ = "end of CoverTab[87371]"
}

func (r *Reader) Uint64() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:64
	_go_fuzz_dep_.CoverTab[87374]++
												b, err := r.ReadBytes(SizeUint64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:66
		_go_fuzz_dep_.CoverTab[87376]++
													return uint64(0), err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:67
		// _ = "end of CoverTab[87376]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:68
		_go_fuzz_dep_.CoverTab[87377]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:68
		// _ = "end of CoverTab[87377]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:68
	// _ = "end of CoverTab[87374]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:68
	_go_fuzz_dep_.CoverTab[87375]++
												return binary.LittleEndian.Uint64(b), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:69
	// _ = "end of CoverTab[87375]"
}

func (r *Reader) FileTime() (f FileTime, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:72
	_go_fuzz_dep_.CoverTab[87378]++
												f.LowDateTime, err = r.Uint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:74
		_go_fuzz_dep_.CoverTab[87381]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:75
		// _ = "end of CoverTab[87381]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:76
		_go_fuzz_dep_.CoverTab[87382]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:76
		// _ = "end of CoverTab[87382]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:76
	// _ = "end of CoverTab[87378]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:76
	_go_fuzz_dep_.CoverTab[87379]++
												f.HighDateTime, err = r.Uint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:78
		_go_fuzz_dep_.CoverTab[87383]++
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:79
		// _ = "end of CoverTab[87383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:80
		_go_fuzz_dep_.CoverTab[87384]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:80
		// _ = "end of CoverTab[87384]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:80
	// _ = "end of CoverTab[87379]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:80
	_go_fuzz_dep_.CoverTab[87380]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:81
	// _ = "end of CoverTab[87380]"
}

// UTF16String returns a string that is UTF16 encoded in a byte slice. n is the number of bytes representing the string
func (r *Reader) UTF16String(n int) (str string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:85
	_go_fuzz_dep_.CoverTab[87385]++

												s := make([]rune, n/2, n/2)
												for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:88
		_go_fuzz_dep_.CoverTab[87387]++
													var u uint16
													u, err = r.Uint16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:91
			_go_fuzz_dep_.CoverTab[87389]++
														return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:92
			// _ = "end of CoverTab[87389]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:93
			_go_fuzz_dep_.CoverTab[87390]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:93
			// _ = "end of CoverTab[87390]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:93
		// _ = "end of CoverTab[87387]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:93
		_go_fuzz_dep_.CoverTab[87388]++
													s[i] = rune(u)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:94
		// _ = "end of CoverTab[87388]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:95
	// _ = "end of CoverTab[87385]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:95
	_go_fuzz_dep_.CoverTab[87386]++
												str = string(s)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:97
	// _ = "end of CoverTab[87386]"
}

// readBytes returns a number of bytes from the NDR byte stream.
func (r *Reader) ReadBytes(n int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:101
	_go_fuzz_dep_.CoverTab[87391]++

												b := make([]byte, n, n)
												m, err := r.r.Read(b)
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:105
		_go_fuzz_dep_.CoverTab[87393]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:105
		return m != n
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:105
		// _ = "end of CoverTab[87393]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:105
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:105
		_go_fuzz_dep_.CoverTab[87394]++
													return b, fmt.Errorf("error reading bytes from stream: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:106
		// _ = "end of CoverTab[87394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:107
		_go_fuzz_dep_.CoverTab[87395]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:107
		// _ = "end of CoverTab[87395]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:107
	// _ = "end of CoverTab[87391]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:107
	_go_fuzz_dep_.CoverTab[87392]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:108
	// _ = "end of CoverTab[87392]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/mstypes/reader.go:109
var _ = _go_fuzz_dep_.CoverTab
