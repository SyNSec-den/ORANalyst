//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:1
)

import (
	"encoding/binary"
	"math"
)

var (
	errInvalidArrayLength		= PacketDecodingError{"invalid array length"}
	errInvalidByteSliceLength	= PacketDecodingError{"invalid byteslice length"}
	errInvalidStringLength		= PacketDecodingError{"invalid string length"}
	errVarintOverflow		= PacketDecodingError{"varint overflow"}
	errUVarintOverflow		= PacketDecodingError{"uvarint overflow"}
	errInvalidBool			= PacketDecodingError{"invalid bool"}
)

type realDecoder struct {
	raw	[]byte
	off	int
	stack	[]pushDecoder
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:25
func (rd *realDecoder) getInt8() (int8, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:25
	_go_fuzz_dep_.CoverTab[105978]++
												if rd.remaining() < 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:26
		_go_fuzz_dep_.CoverTab[105980]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:28
		// _ = "end of CoverTab[105980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:29
		_go_fuzz_dep_.CoverTab[105981]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:29
		// _ = "end of CoverTab[105981]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:29
	// _ = "end of CoverTab[105978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:29
	_go_fuzz_dep_.CoverTab[105979]++
												tmp := int8(rd.raw[rd.off])
												rd.off++
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:32
	// _ = "end of CoverTab[105979]"
}

func (rd *realDecoder) getInt16() (int16, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:35
	_go_fuzz_dep_.CoverTab[105982]++
												if rd.remaining() < 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:36
		_go_fuzz_dep_.CoverTab[105984]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:38
		// _ = "end of CoverTab[105984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:39
		_go_fuzz_dep_.CoverTab[105985]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:39
		// _ = "end of CoverTab[105985]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:39
	// _ = "end of CoverTab[105982]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:39
	_go_fuzz_dep_.CoverTab[105983]++
												tmp := int16(binary.BigEndian.Uint16(rd.raw[rd.off:]))
												rd.off += 2
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:42
	// _ = "end of CoverTab[105983]"
}

func (rd *realDecoder) getInt32() (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:45
	_go_fuzz_dep_.CoverTab[105986]++
												if rd.remaining() < 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:46
		_go_fuzz_dep_.CoverTab[105988]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:48
		// _ = "end of CoverTab[105988]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:49
		_go_fuzz_dep_.CoverTab[105989]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:49
		// _ = "end of CoverTab[105989]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:49
	// _ = "end of CoverTab[105986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:49
	_go_fuzz_dep_.CoverTab[105987]++
												tmp := int32(binary.BigEndian.Uint32(rd.raw[rd.off:]))
												rd.off += 4
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:52
	// _ = "end of CoverTab[105987]"
}

func (rd *realDecoder) getInt64() (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:55
	_go_fuzz_dep_.CoverTab[105990]++
												if rd.remaining() < 8 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:56
		_go_fuzz_dep_.CoverTab[105992]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:58
		// _ = "end of CoverTab[105992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:59
		_go_fuzz_dep_.CoverTab[105993]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:59
		// _ = "end of CoverTab[105993]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:59
	// _ = "end of CoverTab[105990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:59
	_go_fuzz_dep_.CoverTab[105991]++
												tmp := int64(binary.BigEndian.Uint64(rd.raw[rd.off:]))
												rd.off += 8
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:62
	// _ = "end of CoverTab[105991]"
}

func (rd *realDecoder) getVarint() (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:65
	_go_fuzz_dep_.CoverTab[105994]++
												tmp, n := binary.Varint(rd.raw[rd.off:])
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:67
		_go_fuzz_dep_.CoverTab[105997]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:69
		// _ = "end of CoverTab[105997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:70
		_go_fuzz_dep_.CoverTab[105998]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:70
		// _ = "end of CoverTab[105998]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:70
	// _ = "end of CoverTab[105994]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:70
	_go_fuzz_dep_.CoverTab[105995]++
												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:71
		_go_fuzz_dep_.CoverTab[105999]++
													rd.off -= n
													return -1, errVarintOverflow
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:73
		// _ = "end of CoverTab[105999]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:74
		_go_fuzz_dep_.CoverTab[106000]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:74
		// _ = "end of CoverTab[106000]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:74
	// _ = "end of CoverTab[105995]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:74
	_go_fuzz_dep_.CoverTab[105996]++
												rd.off += n
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:76
	// _ = "end of CoverTab[105996]"
}

func (rd *realDecoder) getUVarint() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:79
	_go_fuzz_dep_.CoverTab[106001]++
												tmp, n := binary.Uvarint(rd.raw[rd.off:])
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:81
		_go_fuzz_dep_.CoverTab[106004]++
													rd.off = len(rd.raw)
													return 0, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:83
		// _ = "end of CoverTab[106004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:84
		_go_fuzz_dep_.CoverTab[106005]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:84
		// _ = "end of CoverTab[106005]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:84
	// _ = "end of CoverTab[106001]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:84
	_go_fuzz_dep_.CoverTab[106002]++

												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:86
		_go_fuzz_dep_.CoverTab[106006]++
													rd.off -= n
													return 0, errUVarintOverflow
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:88
		// _ = "end of CoverTab[106006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:89
		_go_fuzz_dep_.CoverTab[106007]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:89
		// _ = "end of CoverTab[106007]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:89
	// _ = "end of CoverTab[106002]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:89
	_go_fuzz_dep_.CoverTab[106003]++

												rd.off += n
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:92
	// _ = "end of CoverTab[106003]"
}

func (rd *realDecoder) getFloat64() (float64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:95
	_go_fuzz_dep_.CoverTab[106008]++
												if rd.remaining() < 8 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:96
		_go_fuzz_dep_.CoverTab[106010]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:98
		// _ = "end of CoverTab[106010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:99
		_go_fuzz_dep_.CoverTab[106011]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:99
		// _ = "end of CoverTab[106011]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:99
	// _ = "end of CoverTab[106008]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:99
	_go_fuzz_dep_.CoverTab[106009]++
												tmp := math.Float64frombits(binary.BigEndian.Uint64(rd.raw[rd.off:]))
												rd.off += 8
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:102
	// _ = "end of CoverTab[106009]"
}

func (rd *realDecoder) getArrayLength() (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:105
	_go_fuzz_dep_.CoverTab[106012]++
												if rd.remaining() < 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:106
		_go_fuzz_dep_.CoverTab[106015]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:108
		// _ = "end of CoverTab[106015]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:109
		_go_fuzz_dep_.CoverTab[106016]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:109
		// _ = "end of CoverTab[106016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:109
	// _ = "end of CoverTab[106012]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:109
	_go_fuzz_dep_.CoverTab[106013]++
												tmp := int(int32(binary.BigEndian.Uint32(rd.raw[rd.off:])))
												rd.off += 4
												if tmp > rd.remaining() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:112
		_go_fuzz_dep_.CoverTab[106017]++
													rd.off = len(rd.raw)
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:114
		// _ = "end of CoverTab[106017]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:115
		_go_fuzz_dep_.CoverTab[106018]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:115
		if tmp > 2*math.MaxUint16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:115
			_go_fuzz_dep_.CoverTab[106019]++
														return -1, errInvalidArrayLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:116
			// _ = "end of CoverTab[106019]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
			_go_fuzz_dep_.CoverTab[106020]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
			// _ = "end of CoverTab[106020]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
		// _ = "end of CoverTab[106018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
	// _ = "end of CoverTab[106013]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:117
	_go_fuzz_dep_.CoverTab[106014]++
												return tmp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:118
	// _ = "end of CoverTab[106014]"
}

func (rd *realDecoder) getCompactArrayLength() (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:121
	_go_fuzz_dep_.CoverTab[106021]++
												n, err := rd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:123
		_go_fuzz_dep_.CoverTab[106024]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:124
		// _ = "end of CoverTab[106024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:125
		_go_fuzz_dep_.CoverTab[106025]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:125
		// _ = "end of CoverTab[106025]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:125
	// _ = "end of CoverTab[106021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:125
	_go_fuzz_dep_.CoverTab[106022]++

												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:127
		_go_fuzz_dep_.CoverTab[106026]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:128
		// _ = "end of CoverTab[106026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:129
		_go_fuzz_dep_.CoverTab[106027]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:129
		// _ = "end of CoverTab[106027]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:129
	// _ = "end of CoverTab[106022]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:129
	_go_fuzz_dep_.CoverTab[106023]++

												return int(n) - 1, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:131
	// _ = "end of CoverTab[106023]"
}

func (rd *realDecoder) getBool() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:134
	_go_fuzz_dep_.CoverTab[106028]++
												b, err := rd.getInt8()
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:136
		_go_fuzz_dep_.CoverTab[106031]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:136
		return b == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:136
		// _ = "end of CoverTab[106031]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:136
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:136
		_go_fuzz_dep_.CoverTab[106032]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:137
		// _ = "end of CoverTab[106032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:138
		_go_fuzz_dep_.CoverTab[106033]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:138
		// _ = "end of CoverTab[106033]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:138
	// _ = "end of CoverTab[106028]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:138
	_go_fuzz_dep_.CoverTab[106029]++
												if b != 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:139
		_go_fuzz_dep_.CoverTab[106034]++
													return false, errInvalidBool
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:140
		// _ = "end of CoverTab[106034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:141
		_go_fuzz_dep_.CoverTab[106035]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:141
		// _ = "end of CoverTab[106035]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:141
	// _ = "end of CoverTab[106029]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:141
	_go_fuzz_dep_.CoverTab[106030]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:142
	// _ = "end of CoverTab[106030]"
}

func (rd *realDecoder) getEmptyTaggedFieldArray() (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:145
	_go_fuzz_dep_.CoverTab[106036]++
												tagCount, err := rd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:147
		_go_fuzz_dep_.CoverTab[106039]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:148
		// _ = "end of CoverTab[106039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:149
		_go_fuzz_dep_.CoverTab[106040]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:149
		// _ = "end of CoverTab[106040]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:149
	// _ = "end of CoverTab[106036]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:149
	_go_fuzz_dep_.CoverTab[106037]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:153
	for i := uint64(0); i < tagCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:153
		_go_fuzz_dep_.CoverTab[106041]++

													_, err := rd.getUVarint()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:156
			_go_fuzz_dep_.CoverTab[106044]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:157
			// _ = "end of CoverTab[106044]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:158
			_go_fuzz_dep_.CoverTab[106045]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:158
			// _ = "end of CoverTab[106045]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:158
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:158
		// _ = "end of CoverTab[106041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:158
		_go_fuzz_dep_.CoverTab[106042]++
													length, err := rd.getUVarint()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:160
			_go_fuzz_dep_.CoverTab[106046]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:161
			// _ = "end of CoverTab[106046]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:162
			_go_fuzz_dep_.CoverTab[106047]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:162
			// _ = "end of CoverTab[106047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:162
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:162
		// _ = "end of CoverTab[106042]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:162
		_go_fuzz_dep_.CoverTab[106043]++
													if _, err := rd.getRawBytes(int(length)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:163
			_go_fuzz_dep_.CoverTab[106048]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:164
			// _ = "end of CoverTab[106048]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:165
			_go_fuzz_dep_.CoverTab[106049]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:165
			// _ = "end of CoverTab[106049]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:165
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:165
		// _ = "end of CoverTab[106043]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:166
	// _ = "end of CoverTab[106037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:166
	_go_fuzz_dep_.CoverTab[106038]++

												return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:168
	// _ = "end of CoverTab[106038]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:173
func (rd *realDecoder) getBytes() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:173
	_go_fuzz_dep_.CoverTab[106050]++
												tmp, err := rd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:175
		_go_fuzz_dep_.CoverTab[106053]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:176
		// _ = "end of CoverTab[106053]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:177
		_go_fuzz_dep_.CoverTab[106054]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:177
		// _ = "end of CoverTab[106054]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:177
	// _ = "end of CoverTab[106050]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:177
	_go_fuzz_dep_.CoverTab[106051]++
												if tmp == -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:178
		_go_fuzz_dep_.CoverTab[106055]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:179
		// _ = "end of CoverTab[106055]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:180
		_go_fuzz_dep_.CoverTab[106056]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:180
		// _ = "end of CoverTab[106056]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:180
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:180
	// _ = "end of CoverTab[106051]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:180
	_go_fuzz_dep_.CoverTab[106052]++

												return rd.getRawBytes(int(tmp))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:182
	// _ = "end of CoverTab[106052]"
}

func (rd *realDecoder) getVarintBytes() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:185
	_go_fuzz_dep_.CoverTab[106057]++
												tmp, err := rd.getVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:187
		_go_fuzz_dep_.CoverTab[106060]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:188
		// _ = "end of CoverTab[106060]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:189
		_go_fuzz_dep_.CoverTab[106061]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:189
		// _ = "end of CoverTab[106061]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:189
	// _ = "end of CoverTab[106057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:189
	_go_fuzz_dep_.CoverTab[106058]++
												if tmp == -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:190
		_go_fuzz_dep_.CoverTab[106062]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:191
		// _ = "end of CoverTab[106062]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:192
		_go_fuzz_dep_.CoverTab[106063]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:192
		// _ = "end of CoverTab[106063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:192
	// _ = "end of CoverTab[106058]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:192
	_go_fuzz_dep_.CoverTab[106059]++

												return rd.getRawBytes(int(tmp))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:194
	// _ = "end of CoverTab[106059]"
}

func (rd *realDecoder) getCompactBytes() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:197
	_go_fuzz_dep_.CoverTab[106064]++
												n, err := rd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:199
		_go_fuzz_dep_.CoverTab[106066]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:200
		// _ = "end of CoverTab[106066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:201
		_go_fuzz_dep_.CoverTab[106067]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:201
		// _ = "end of CoverTab[106067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:201
	// _ = "end of CoverTab[106064]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:201
	_go_fuzz_dep_.CoverTab[106065]++

												length := int(n - 1)
												return rd.getRawBytes(length)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:204
	// _ = "end of CoverTab[106065]"
}

func (rd *realDecoder) getStringLength() (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:207
	_go_fuzz_dep_.CoverTab[106068]++
												length, err := rd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:209
		_go_fuzz_dep_.CoverTab[106071]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:210
		// _ = "end of CoverTab[106071]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:211
		_go_fuzz_dep_.CoverTab[106072]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:211
		// _ = "end of CoverTab[106072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:211
	// _ = "end of CoverTab[106068]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:211
	_go_fuzz_dep_.CoverTab[106069]++

												n := int(length)

												switch {
	case n < -1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:216
		_go_fuzz_dep_.CoverTab[106073]++
													return 0, errInvalidStringLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:217
		// _ = "end of CoverTab[106073]"
	case n > rd.remaining():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:218
		_go_fuzz_dep_.CoverTab[106074]++
													rd.off = len(rd.raw)
													return 0, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:220
		// _ = "end of CoverTab[106074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:220
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:220
		_go_fuzz_dep_.CoverTab[106075]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:220
		// _ = "end of CoverTab[106075]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:221
	// _ = "end of CoverTab[106069]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:221
	_go_fuzz_dep_.CoverTab[106070]++

												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:223
	// _ = "end of CoverTab[106070]"
}

func (rd *realDecoder) getString() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:226
	_go_fuzz_dep_.CoverTab[106076]++
												n, err := rd.getStringLength()
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:228
		_go_fuzz_dep_.CoverTab[106078]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:228
		return n == -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:228
		// _ = "end of CoverTab[106078]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:228
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:228
		_go_fuzz_dep_.CoverTab[106079]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:229
		// _ = "end of CoverTab[106079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:230
		_go_fuzz_dep_.CoverTab[106080]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:230
		// _ = "end of CoverTab[106080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:230
	// _ = "end of CoverTab[106076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:230
	_go_fuzz_dep_.CoverTab[106077]++

												tmpStr := string(rd.raw[rd.off : rd.off+n])
												rd.off += n
												return tmpStr, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:234
	// _ = "end of CoverTab[106077]"
}

func (rd *realDecoder) getNullableString() (*string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:237
	_go_fuzz_dep_.CoverTab[106081]++
												n, err := rd.getStringLength()
												if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:239
		_go_fuzz_dep_.CoverTab[106083]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:239
		return n == -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:239
		// _ = "end of CoverTab[106083]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:239
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:239
		_go_fuzz_dep_.CoverTab[106084]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:240
		// _ = "end of CoverTab[106084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:241
		_go_fuzz_dep_.CoverTab[106085]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:241
		// _ = "end of CoverTab[106085]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:241
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:241
	// _ = "end of CoverTab[106081]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:241
	_go_fuzz_dep_.CoverTab[106082]++

												tmpStr := string(rd.raw[rd.off : rd.off+n])
												rd.off += n
												return &tmpStr, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:245
	// _ = "end of CoverTab[106082]"
}

func (rd *realDecoder) getCompactString() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:248
	_go_fuzz_dep_.CoverTab[106086]++
												n, err := rd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:250
		_go_fuzz_dep_.CoverTab[106089]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:251
		// _ = "end of CoverTab[106089]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:252
		_go_fuzz_dep_.CoverTab[106090]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:252
		// _ = "end of CoverTab[106090]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:252
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:252
	// _ = "end of CoverTab[106086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:252
	_go_fuzz_dep_.CoverTab[106087]++

												length := int(n - 1)
												if length < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:255
		_go_fuzz_dep_.CoverTab[106091]++
													return "", errInvalidByteSliceLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:256
		// _ = "end of CoverTab[106091]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:257
		_go_fuzz_dep_.CoverTab[106092]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:257
		// _ = "end of CoverTab[106092]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:257
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:257
	// _ = "end of CoverTab[106087]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:257
	_go_fuzz_dep_.CoverTab[106088]++
												tmpStr := string(rd.raw[rd.off : rd.off+length])
												rd.off += length
												return tmpStr, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:260
	// _ = "end of CoverTab[106088]"
}

func (rd *realDecoder) getCompactNullableString() (*string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:263
	_go_fuzz_dep_.CoverTab[106093]++
												n, err := rd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:265
		_go_fuzz_dep_.CoverTab[106096]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:266
		// _ = "end of CoverTab[106096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:267
		_go_fuzz_dep_.CoverTab[106097]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:267
		// _ = "end of CoverTab[106097]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:267
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:267
	// _ = "end of CoverTab[106093]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:267
	_go_fuzz_dep_.CoverTab[106094]++

												length := int(n - 1)

												if length < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:271
		_go_fuzz_dep_.CoverTab[106098]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:272
		// _ = "end of CoverTab[106098]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:273
		_go_fuzz_dep_.CoverTab[106099]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:273
		// _ = "end of CoverTab[106099]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:273
	// _ = "end of CoverTab[106094]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:273
	_go_fuzz_dep_.CoverTab[106095]++

												tmpStr := string(rd.raw[rd.off : rd.off+length])
												rd.off += length
												return &tmpStr, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:277
	// _ = "end of CoverTab[106095]"
}

func (rd *realDecoder) getCompactInt32Array() ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:280
	_go_fuzz_dep_.CoverTab[106100]++
												n, err := rd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:282
		_go_fuzz_dep_.CoverTab[106104]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:283
		// _ = "end of CoverTab[106104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:284
		_go_fuzz_dep_.CoverTab[106105]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:284
		// _ = "end of CoverTab[106105]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:284
	// _ = "end of CoverTab[106100]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:284
	_go_fuzz_dep_.CoverTab[106101]++

												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:286
		_go_fuzz_dep_.CoverTab[106106]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:287
		// _ = "end of CoverTab[106106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:288
		_go_fuzz_dep_.CoverTab[106107]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:288
		// _ = "end of CoverTab[106107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:288
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:288
	// _ = "end of CoverTab[106101]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:288
	_go_fuzz_dep_.CoverTab[106102]++

												arrayLength := int(n) - 1

												ret := make([]int32, arrayLength)

												for i := range ret {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:294
		_go_fuzz_dep_.CoverTab[106108]++
													ret[i] = int32(binary.BigEndian.Uint32(rd.raw[rd.off:]))
													rd.off += 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:296
		// _ = "end of CoverTab[106108]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:297
	// _ = "end of CoverTab[106102]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:297
	_go_fuzz_dep_.CoverTab[106103]++
												return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:298
	// _ = "end of CoverTab[106103]"
}

func (rd *realDecoder) getInt32Array() ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:301
	_go_fuzz_dep_.CoverTab[106109]++
												if rd.remaining() < 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:302
		_go_fuzz_dep_.CoverTab[106115]++
													rd.off = len(rd.raw)
													return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:304
		// _ = "end of CoverTab[106115]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:305
		_go_fuzz_dep_.CoverTab[106116]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:305
		// _ = "end of CoverTab[106116]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:305
	// _ = "end of CoverTab[106109]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:305
	_go_fuzz_dep_.CoverTab[106110]++
												n := int(binary.BigEndian.Uint32(rd.raw[rd.off:]))
												rd.off += 4

												if rd.remaining() < 4*n {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:309
		_go_fuzz_dep_.CoverTab[106117]++
													rd.off = len(rd.raw)
													return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:311
		// _ = "end of CoverTab[106117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:312
		_go_fuzz_dep_.CoverTab[106118]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:312
		// _ = "end of CoverTab[106118]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:312
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:312
	// _ = "end of CoverTab[106110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:312
	_go_fuzz_dep_.CoverTab[106111]++

												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:314
		_go_fuzz_dep_.CoverTab[106119]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:315
		// _ = "end of CoverTab[106119]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:316
		_go_fuzz_dep_.CoverTab[106120]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:316
		// _ = "end of CoverTab[106120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:316
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:316
	// _ = "end of CoverTab[106111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:316
	_go_fuzz_dep_.CoverTab[106112]++

												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:318
		_go_fuzz_dep_.CoverTab[106121]++
													return nil, errInvalidArrayLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:319
		// _ = "end of CoverTab[106121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:320
		_go_fuzz_dep_.CoverTab[106122]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:320
		// _ = "end of CoverTab[106122]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:320
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:320
	// _ = "end of CoverTab[106112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:320
	_go_fuzz_dep_.CoverTab[106113]++

												ret := make([]int32, n)
												for i := range ret {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:323
		_go_fuzz_dep_.CoverTab[106123]++
													ret[i] = int32(binary.BigEndian.Uint32(rd.raw[rd.off:]))
													rd.off += 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:325
		// _ = "end of CoverTab[106123]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:326
	// _ = "end of CoverTab[106113]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:326
	_go_fuzz_dep_.CoverTab[106114]++
												return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:327
	// _ = "end of CoverTab[106114]"
}

func (rd *realDecoder) getInt64Array() ([]int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:330
	_go_fuzz_dep_.CoverTab[106124]++
												if rd.remaining() < 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:331
		_go_fuzz_dep_.CoverTab[106130]++
													rd.off = len(rd.raw)
													return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:333
		// _ = "end of CoverTab[106130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:334
		_go_fuzz_dep_.CoverTab[106131]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:334
		// _ = "end of CoverTab[106131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:334
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:334
	// _ = "end of CoverTab[106124]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:334
	_go_fuzz_dep_.CoverTab[106125]++
												n := int(binary.BigEndian.Uint32(rd.raw[rd.off:]))
												rd.off += 4

												if rd.remaining() < 8*n {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:338
		_go_fuzz_dep_.CoverTab[106132]++
													rd.off = len(rd.raw)
													return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:340
		// _ = "end of CoverTab[106132]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:341
		_go_fuzz_dep_.CoverTab[106133]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:341
		// _ = "end of CoverTab[106133]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:341
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:341
	// _ = "end of CoverTab[106125]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:341
	_go_fuzz_dep_.CoverTab[106126]++

												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:343
		_go_fuzz_dep_.CoverTab[106134]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:344
		// _ = "end of CoverTab[106134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:345
		_go_fuzz_dep_.CoverTab[106135]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:345
		// _ = "end of CoverTab[106135]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:345
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:345
	// _ = "end of CoverTab[106126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:345
	_go_fuzz_dep_.CoverTab[106127]++

												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:347
		_go_fuzz_dep_.CoverTab[106136]++
													return nil, errInvalidArrayLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:348
		// _ = "end of CoverTab[106136]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:349
		_go_fuzz_dep_.CoverTab[106137]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:349
		// _ = "end of CoverTab[106137]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:349
	// _ = "end of CoverTab[106127]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:349
	_go_fuzz_dep_.CoverTab[106128]++

												ret := make([]int64, n)
												for i := range ret {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:352
		_go_fuzz_dep_.CoverTab[106138]++
													ret[i] = int64(binary.BigEndian.Uint64(rd.raw[rd.off:]))
													rd.off += 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:354
		// _ = "end of CoverTab[106138]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:355
	// _ = "end of CoverTab[106128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:355
	_go_fuzz_dep_.CoverTab[106129]++
												return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:356
	// _ = "end of CoverTab[106129]"
}

func (rd *realDecoder) getStringArray() ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:359
	_go_fuzz_dep_.CoverTab[106139]++
												if rd.remaining() < 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:360
		_go_fuzz_dep_.CoverTab[106144]++
													rd.off = len(rd.raw)
													return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:362
		// _ = "end of CoverTab[106144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:363
		_go_fuzz_dep_.CoverTab[106145]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:363
		// _ = "end of CoverTab[106145]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:363
	// _ = "end of CoverTab[106139]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:363
	_go_fuzz_dep_.CoverTab[106140]++
												n := int(binary.BigEndian.Uint32(rd.raw[rd.off:]))
												rd.off += 4

												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:367
		_go_fuzz_dep_.CoverTab[106146]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:368
		// _ = "end of CoverTab[106146]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:369
		_go_fuzz_dep_.CoverTab[106147]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:369
		// _ = "end of CoverTab[106147]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:369
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:369
	// _ = "end of CoverTab[106140]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:369
	_go_fuzz_dep_.CoverTab[106141]++

												if n < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:371
		_go_fuzz_dep_.CoverTab[106148]++
													return nil, errInvalidArrayLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:372
		// _ = "end of CoverTab[106148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:373
		_go_fuzz_dep_.CoverTab[106149]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:373
		// _ = "end of CoverTab[106149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:373
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:373
	// _ = "end of CoverTab[106141]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:373
	_go_fuzz_dep_.CoverTab[106142]++

												ret := make([]string, n)
												for i := range ret {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:376
		_go_fuzz_dep_.CoverTab[106150]++
													str, err := rd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:378
			_go_fuzz_dep_.CoverTab[106152]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:379
			// _ = "end of CoverTab[106152]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:380
			_go_fuzz_dep_.CoverTab[106153]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:380
			// _ = "end of CoverTab[106153]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:380
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:380
		// _ = "end of CoverTab[106150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:380
		_go_fuzz_dep_.CoverTab[106151]++

													ret[i] = str
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:382
		// _ = "end of CoverTab[106151]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:383
	// _ = "end of CoverTab[106142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:383
	_go_fuzz_dep_.CoverTab[106143]++
												return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:384
	// _ = "end of CoverTab[106143]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:389
func (rd *realDecoder) remaining() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:389
	_go_fuzz_dep_.CoverTab[106154]++
												return len(rd.raw) - rd.off
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:390
	// _ = "end of CoverTab[106154]"
}

func (rd *realDecoder) getSubset(length int) (packetDecoder, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:393
	_go_fuzz_dep_.CoverTab[106155]++
												buf, err := rd.getRawBytes(length)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:395
		_go_fuzz_dep_.CoverTab[106157]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:396
		// _ = "end of CoverTab[106157]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:397
		_go_fuzz_dep_.CoverTab[106158]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:397
		// _ = "end of CoverTab[106158]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:397
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:397
	// _ = "end of CoverTab[106155]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:397
	_go_fuzz_dep_.CoverTab[106156]++
												return &realDecoder{raw: buf}, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:398
	// _ = "end of CoverTab[106156]"
}

func (rd *realDecoder) getRawBytes(length int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:401
	_go_fuzz_dep_.CoverTab[106159]++
												if length < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:402
		_go_fuzz_dep_.CoverTab[106161]++
													return nil, errInvalidByteSliceLength
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:403
		// _ = "end of CoverTab[106161]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:404
		_go_fuzz_dep_.CoverTab[106162]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:404
		if length > rd.remaining() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:404
			_go_fuzz_dep_.CoverTab[106163]++
														rd.off = len(rd.raw)
														return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:406
			// _ = "end of CoverTab[106163]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
			_go_fuzz_dep_.CoverTab[106164]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
			// _ = "end of CoverTab[106164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
		// _ = "end of CoverTab[106162]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
	// _ = "end of CoverTab[106159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:407
	_go_fuzz_dep_.CoverTab[106160]++

												start := rd.off
												rd.off += length
												return rd.raw[start:rd.off], nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:411
	// _ = "end of CoverTab[106160]"
}

func (rd *realDecoder) peek(offset, length int) (packetDecoder, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:414
	_go_fuzz_dep_.CoverTab[106165]++
												if rd.remaining() < offset+length {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:415
		_go_fuzz_dep_.CoverTab[106167]++
													return nil, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:416
		// _ = "end of CoverTab[106167]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:417
		_go_fuzz_dep_.CoverTab[106168]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:417
		// _ = "end of CoverTab[106168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:417
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:417
	// _ = "end of CoverTab[106165]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:417
	_go_fuzz_dep_.CoverTab[106166]++
												off := rd.off + offset
												return &realDecoder{raw: rd.raw[off : off+length]}, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:419
	// _ = "end of CoverTab[106166]"
}

func (rd *realDecoder) peekInt8(offset int) (int8, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:422
	_go_fuzz_dep_.CoverTab[106169]++
												const byteLen = 1
												if rd.remaining() < offset+byteLen {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:424
		_go_fuzz_dep_.CoverTab[106171]++
													return -1, ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:425
		// _ = "end of CoverTab[106171]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:426
		_go_fuzz_dep_.CoverTab[106172]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:426
		// _ = "end of CoverTab[106172]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:426
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:426
	// _ = "end of CoverTab[106169]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:426
	_go_fuzz_dep_.CoverTab[106170]++
												return int8(rd.raw[rd.off+offset]), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:427
	// _ = "end of CoverTab[106170]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:432
func (rd *realDecoder) push(in pushDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:432
	_go_fuzz_dep_.CoverTab[106173]++
												in.saveOffset(rd.off)

												var reserve int
												if dpd, ok := in.(dynamicPushDecoder); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:436
		_go_fuzz_dep_.CoverTab[106175]++
													if err := dpd.decode(rd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:437
			_go_fuzz_dep_.CoverTab[106176]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:438
			// _ = "end of CoverTab[106176]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:439
			_go_fuzz_dep_.CoverTab[106177]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:439
			// _ = "end of CoverTab[106177]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:439
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:439
		// _ = "end of CoverTab[106175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:440
		_go_fuzz_dep_.CoverTab[106178]++
													reserve = in.reserveLength()
													if rd.remaining() < reserve {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:442
			_go_fuzz_dep_.CoverTab[106179]++
														rd.off = len(rd.raw)
														return ErrInsufficientData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:444
			// _ = "end of CoverTab[106179]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:445
			_go_fuzz_dep_.CoverTab[106180]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:445
			// _ = "end of CoverTab[106180]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:445
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:445
		// _ = "end of CoverTab[106178]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:446
	// _ = "end of CoverTab[106173]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:446
	_go_fuzz_dep_.CoverTab[106174]++

												rd.stack = append(rd.stack, in)

												rd.off += reserve

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:452
	// _ = "end of CoverTab[106174]"
}

func (rd *realDecoder) pop() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:455
	_go_fuzz_dep_.CoverTab[106181]++

												in := rd.stack[len(rd.stack)-1]
												rd.stack = rd.stack[:len(rd.stack)-1]

												return in.check(rd.off, rd.raw)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:460
	// _ = "end of CoverTab[106181]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:461
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_decoder.go:461
var _ = _go_fuzz_dep_.CoverTab
