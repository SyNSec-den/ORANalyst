//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:1
)

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"sync"
)

type crcPolynomial int8

const (
	crcIEEE	crcPolynomial	= iota
	crcCastagnoli
)

var crc32FieldPool = sync.Pool{}

func acquireCrc32Field(polynomial crcPolynomial) *crc32Field {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:19
	_go_fuzz_dep_.CoverTab[101408]++
											val := crc32FieldPool.Get()
											if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:21
		_go_fuzz_dep_.CoverTab[101410]++
												c := val.(*crc32Field)
												c.polynomial = polynomial
												return c
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:24
		// _ = "end of CoverTab[101410]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:25
		_go_fuzz_dep_.CoverTab[101411]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:25
		// _ = "end of CoverTab[101411]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:25
	// _ = "end of CoverTab[101408]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:25
	_go_fuzz_dep_.CoverTab[101409]++
											return newCRC32Field(polynomial)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:26
	// _ = "end of CoverTab[101409]"
}

func releaseCrc32Field(c *crc32Field) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:29
	_go_fuzz_dep_.CoverTab[101412]++
											crc32FieldPool.Put(c)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:30
	// _ = "end of CoverTab[101412]"
}

var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

// crc32Field implements the pushEncoder and pushDecoder interfaces for calculating CRC32s.
type crc32Field struct {
	startOffset	int
	polynomial	crcPolynomial
}

func (c *crc32Field) saveOffset(in int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:41
	_go_fuzz_dep_.CoverTab[101413]++
											c.startOffset = in
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:42
	// _ = "end of CoverTab[101413]"
}

func (c *crc32Field) reserveLength() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:45
	_go_fuzz_dep_.CoverTab[101414]++
											return 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:46
	// _ = "end of CoverTab[101414]"
}

func newCRC32Field(polynomial crcPolynomial) *crc32Field {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:49
	_go_fuzz_dep_.CoverTab[101415]++
											return &crc32Field{polynomial: polynomial}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:50
	// _ = "end of CoverTab[101415]"
}

func (c *crc32Field) run(curOffset int, buf []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:53
	_go_fuzz_dep_.CoverTab[101416]++
											crc, err := c.crc(curOffset, buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:55
		_go_fuzz_dep_.CoverTab[101418]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:56
		// _ = "end of CoverTab[101418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:57
		_go_fuzz_dep_.CoverTab[101419]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:57
		// _ = "end of CoverTab[101419]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:57
	// _ = "end of CoverTab[101416]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:57
	_go_fuzz_dep_.CoverTab[101417]++
											binary.BigEndian.PutUint32(buf[c.startOffset:], crc)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:59
	// _ = "end of CoverTab[101417]"
}

func (c *crc32Field) check(curOffset int, buf []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:62
	_go_fuzz_dep_.CoverTab[101420]++
											crc, err := c.crc(curOffset, buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:64
		_go_fuzz_dep_.CoverTab[101423]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:65
		// _ = "end of CoverTab[101423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:66
		_go_fuzz_dep_.CoverTab[101424]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:66
		// _ = "end of CoverTab[101424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:66
	// _ = "end of CoverTab[101420]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:66
	_go_fuzz_dep_.CoverTab[101421]++

											expected := binary.BigEndian.Uint32(buf[c.startOffset:])
											if crc != expected {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:69
		_go_fuzz_dep_.CoverTab[101425]++
												return PacketDecodingError{fmt.Sprintf("CRC didn't match expected %#x got %#x", expected, crc)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:70
		// _ = "end of CoverTab[101425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:71
		_go_fuzz_dep_.CoverTab[101426]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:71
		// _ = "end of CoverTab[101426]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:71
	// _ = "end of CoverTab[101421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:71
	_go_fuzz_dep_.CoverTab[101422]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:73
	// _ = "end of CoverTab[101422]"
}

func (c *crc32Field) crc(curOffset int, buf []byte) (uint32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:76
	_go_fuzz_dep_.CoverTab[101427]++
											var tab *crc32.Table
											switch c.polynomial {
	case crcIEEE:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:79
		_go_fuzz_dep_.CoverTab[101429]++
												tab = crc32.IEEETable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:80
		// _ = "end of CoverTab[101429]"
	case crcCastagnoli:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:81
		_go_fuzz_dep_.CoverTab[101430]++
												tab = castagnoliTable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:82
		// _ = "end of CoverTab[101430]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:83
		_go_fuzz_dep_.CoverTab[101431]++
												return 0, PacketDecodingError{"invalid CRC type"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:84
		// _ = "end of CoverTab[101431]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:85
	// _ = "end of CoverTab[101427]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:85
	_go_fuzz_dep_.CoverTab[101428]++
											return crc32.Checksum(buf[c.startOffset+4:curOffset], tab), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:86
	// _ = "end of CoverTab[101428]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/crc32_field.go:87
var _ = _go_fuzz_dep_.CoverTab
