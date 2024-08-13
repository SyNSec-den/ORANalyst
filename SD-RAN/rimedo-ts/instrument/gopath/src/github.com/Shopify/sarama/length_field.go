//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:1
)

import (
	"encoding/binary"
	"sync"
)

// LengthField implements the PushEncoder and PushDecoder interfaces for calculating 4-byte lengths.
type lengthField struct {
	startOffset	int
	length		int32
}

var lengthFieldPool = sync.Pool{}

func acquireLengthField() *lengthField {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:16
	_go_fuzz_dep_.CoverTab[103756]++
												val := lengthFieldPool.Get()
												if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:18
		_go_fuzz_dep_.CoverTab[103758]++
													return val.(*lengthField)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:19
		// _ = "end of CoverTab[103758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:20
		_go_fuzz_dep_.CoverTab[103759]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:20
		// _ = "end of CoverTab[103759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:20
	// _ = "end of CoverTab[103756]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:20
	_go_fuzz_dep_.CoverTab[103757]++
												return &lengthField{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:21
	// _ = "end of CoverTab[103757]"
}

func releaseLengthField(m *lengthField) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:24
	_go_fuzz_dep_.CoverTab[103760]++
												lengthFieldPool.Put(m)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:25
	// _ = "end of CoverTab[103760]"
}

func (l *lengthField) decode(pd packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:28
	_go_fuzz_dep_.CoverTab[103761]++
												var err error
												l.length, err = pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:31
		_go_fuzz_dep_.CoverTab[103764]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:32
		// _ = "end of CoverTab[103764]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:33
		_go_fuzz_dep_.CoverTab[103765]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:33
		// _ = "end of CoverTab[103765]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:33
	// _ = "end of CoverTab[103761]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:33
	_go_fuzz_dep_.CoverTab[103762]++
												if l.length > int32(pd.remaining()) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:34
		_go_fuzz_dep_.CoverTab[103766]++
													return ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:35
		// _ = "end of CoverTab[103766]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:36
		_go_fuzz_dep_.CoverTab[103767]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:36
		// _ = "end of CoverTab[103767]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:36
	// _ = "end of CoverTab[103762]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:36
	_go_fuzz_dep_.CoverTab[103763]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:37
	// _ = "end of CoverTab[103763]"
}

func (l *lengthField) saveOffset(in int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:40
	_go_fuzz_dep_.CoverTab[103768]++
												l.startOffset = in
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:41
	// _ = "end of CoverTab[103768]"
}

func (l *lengthField) reserveLength() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:44
	_go_fuzz_dep_.CoverTab[103769]++
												return 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:45
	// _ = "end of CoverTab[103769]"
}

func (l *lengthField) run(curOffset int, buf []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:48
	_go_fuzz_dep_.CoverTab[103770]++
												binary.BigEndian.PutUint32(buf[l.startOffset:], uint32(curOffset-l.startOffset-4))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:50
	// _ = "end of CoverTab[103770]"
}

func (l *lengthField) check(curOffset int, buf []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:53
	_go_fuzz_dep_.CoverTab[103771]++
												if int32(curOffset-l.startOffset-4) != l.length {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:54
		_go_fuzz_dep_.CoverTab[103773]++
													return PacketDecodingError{"length field invalid"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:55
		// _ = "end of CoverTab[103773]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:56
		_go_fuzz_dep_.CoverTab[103774]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:56
		// _ = "end of CoverTab[103774]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:56
	// _ = "end of CoverTab[103771]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:56
	_go_fuzz_dep_.CoverTab[103772]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:58
	// _ = "end of CoverTab[103772]"
}

type varintLengthField struct {
	startOffset	int
	length		int64
}

func (l *varintLengthField) decode(pd packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:66
	_go_fuzz_dep_.CoverTab[103775]++
												var err error
												l.length, err = pd.getVarint()
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:69
	// _ = "end of CoverTab[103775]"
}

func (l *varintLengthField) saveOffset(in int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:72
	_go_fuzz_dep_.CoverTab[103776]++
												l.startOffset = in
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:73
	// _ = "end of CoverTab[103776]"
}

func (l *varintLengthField) adjustLength(currOffset int) int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:76
	_go_fuzz_dep_.CoverTab[103777]++
												oldFieldSize := l.reserveLength()
												l.length = int64(currOffset - l.startOffset - oldFieldSize)

												return l.reserveLength() - oldFieldSize
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:80
	// _ = "end of CoverTab[103777]"
}

func (l *varintLengthField) reserveLength() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:83
	_go_fuzz_dep_.CoverTab[103778]++
												var tmp [binary.MaxVarintLen64]byte
												return binary.PutVarint(tmp[:], l.length)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:85
	// _ = "end of CoverTab[103778]"
}

func (l *varintLengthField) run(curOffset int, buf []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:88
	_go_fuzz_dep_.CoverTab[103779]++
												binary.PutVarint(buf[l.startOffset:], l.length)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:90
	// _ = "end of CoverTab[103779]"
}

func (l *varintLengthField) check(curOffset int, buf []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:93
	_go_fuzz_dep_.CoverTab[103780]++
												if int64(curOffset-l.startOffset-l.reserveLength()) != l.length {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:94
		_go_fuzz_dep_.CoverTab[103782]++
													return PacketDecodingError{"length field invalid"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:95
		// _ = "end of CoverTab[103782]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:96
		_go_fuzz_dep_.CoverTab[103783]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:96
		// _ = "end of CoverTab[103783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:96
	// _ = "end of CoverTab[103780]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:96
	_go_fuzz_dep_.CoverTab[103781]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:98
	// _ = "end of CoverTab[103781]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/length_field.go:99
var _ = _go_fuzz_dep_.CoverTab
