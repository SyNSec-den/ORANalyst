//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:1
)

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/rcrowley/go-metrics"
)

type realEncoder struct {
	raw		[]byte
	off		int
	stack		[]pushEncoder
	registry	metrics.Registry
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:20
func (re *realEncoder) putInt8(in int8) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:20
	_go_fuzz_dep_.CoverTab[106182]++
												re.raw[re.off] = byte(in)
												re.off++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:22
	// _ = "end of CoverTab[106182]"
}

func (re *realEncoder) putInt16(in int16) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:25
	_go_fuzz_dep_.CoverTab[106183]++
												binary.BigEndian.PutUint16(re.raw[re.off:], uint16(in))
												re.off += 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:27
	// _ = "end of CoverTab[106183]"
}

func (re *realEncoder) putInt32(in int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:30
	_go_fuzz_dep_.CoverTab[106184]++
												binary.BigEndian.PutUint32(re.raw[re.off:], uint32(in))
												re.off += 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:32
	// _ = "end of CoverTab[106184]"
}

func (re *realEncoder) putInt64(in int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:35
	_go_fuzz_dep_.CoverTab[106185]++
												binary.BigEndian.PutUint64(re.raw[re.off:], uint64(in))
												re.off += 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:37
	// _ = "end of CoverTab[106185]"
}

func (re *realEncoder) putVarint(in int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:40
	_go_fuzz_dep_.CoverTab[106186]++
												re.off += binary.PutVarint(re.raw[re.off:], in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:41
	// _ = "end of CoverTab[106186]"
}

func (re *realEncoder) putUVarint(in uint64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:44
	_go_fuzz_dep_.CoverTab[106187]++
												re.off += binary.PutUvarint(re.raw[re.off:], in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:45
	// _ = "end of CoverTab[106187]"
}

func (re *realEncoder) putFloat64(in float64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:48
	_go_fuzz_dep_.CoverTab[106188]++
												binary.BigEndian.PutUint64(re.raw[re.off:], math.Float64bits(in))
												re.off += 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:50
	// _ = "end of CoverTab[106188]"
}

func (re *realEncoder) putArrayLength(in int) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:53
	_go_fuzz_dep_.CoverTab[106189]++
												re.putInt32(int32(in))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:55
	// _ = "end of CoverTab[106189]"
}

func (re *realEncoder) putCompactArrayLength(in int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:58
	_go_fuzz_dep_.CoverTab[106190]++

												re.putUVarint(uint64(in + 1))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:60
	// _ = "end of CoverTab[106190]"
}

func (re *realEncoder) putBool(in bool) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:63
	_go_fuzz_dep_.CoverTab[106191]++
												if in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:64
		_go_fuzz_dep_.CoverTab[106193]++
													re.putInt8(1)
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:66
		// _ = "end of CoverTab[106193]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:67
		_go_fuzz_dep_.CoverTab[106194]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:67
		// _ = "end of CoverTab[106194]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:67
	// _ = "end of CoverTab[106191]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:67
	_go_fuzz_dep_.CoverTab[106192]++
												re.putInt8(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:68
	// _ = "end of CoverTab[106192]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:73
func (re *realEncoder) putRawBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:73
	_go_fuzz_dep_.CoverTab[106195]++
												copy(re.raw[re.off:], in)
												re.off += len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:76
	// _ = "end of CoverTab[106195]"
}

func (re *realEncoder) putBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:79
	_go_fuzz_dep_.CoverTab[106196]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:80
		_go_fuzz_dep_.CoverTab[106198]++
													re.putInt32(-1)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:82
		// _ = "end of CoverTab[106198]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:83
		_go_fuzz_dep_.CoverTab[106199]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:83
		// _ = "end of CoverTab[106199]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:83
	// _ = "end of CoverTab[106196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:83
	_go_fuzz_dep_.CoverTab[106197]++
												re.putInt32(int32(len(in)))
												return re.putRawBytes(in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:85
	// _ = "end of CoverTab[106197]"
}

func (re *realEncoder) putVarintBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:88
	_go_fuzz_dep_.CoverTab[106200]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:89
		_go_fuzz_dep_.CoverTab[106202]++
													re.putVarint(-1)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:91
		// _ = "end of CoverTab[106202]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:92
		_go_fuzz_dep_.CoverTab[106203]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:92
		// _ = "end of CoverTab[106203]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:92
	// _ = "end of CoverTab[106200]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:92
	_go_fuzz_dep_.CoverTab[106201]++
												re.putVarint(int64(len(in)))
												return re.putRawBytes(in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:94
	// _ = "end of CoverTab[106201]"
}

func (re *realEncoder) putCompactBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:97
	_go_fuzz_dep_.CoverTab[106204]++
												re.putUVarint(uint64(len(in) + 1))
												return re.putRawBytes(in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:99
	// _ = "end of CoverTab[106204]"
}

func (re *realEncoder) putCompactString(in string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:102
	_go_fuzz_dep_.CoverTab[106205]++
												re.putCompactArrayLength(len(in))
												return re.putRawBytes([]byte(in))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:104
	// _ = "end of CoverTab[106205]"
}

func (re *realEncoder) putNullableCompactString(in *string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:107
	_go_fuzz_dep_.CoverTab[106206]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:108
		_go_fuzz_dep_.CoverTab[106208]++
													re.putInt8(0)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:110
		// _ = "end of CoverTab[106208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:111
		_go_fuzz_dep_.CoverTab[106209]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:111
		// _ = "end of CoverTab[106209]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:111
	// _ = "end of CoverTab[106206]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:111
	_go_fuzz_dep_.CoverTab[106207]++
												return re.putCompactString(*in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:112
	// _ = "end of CoverTab[106207]"
}

func (re *realEncoder) putString(in string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:115
	_go_fuzz_dep_.CoverTab[106210]++
												re.putInt16(int16(len(in)))
												copy(re.raw[re.off:], in)
												re.off += len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:119
	// _ = "end of CoverTab[106210]"
}

func (re *realEncoder) putNullableString(in *string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:122
	_go_fuzz_dep_.CoverTab[106211]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:123
		_go_fuzz_dep_.CoverTab[106213]++
													re.putInt16(-1)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:125
		// _ = "end of CoverTab[106213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:126
		_go_fuzz_dep_.CoverTab[106214]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:126
		// _ = "end of CoverTab[106214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:126
	// _ = "end of CoverTab[106211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:126
	_go_fuzz_dep_.CoverTab[106212]++
												return re.putString(*in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:127
	// _ = "end of CoverTab[106212]"
}

func (re *realEncoder) putStringArray(in []string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:130
	_go_fuzz_dep_.CoverTab[106215]++
												err := re.putArrayLength(len(in))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:132
		_go_fuzz_dep_.CoverTab[106218]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:133
		// _ = "end of CoverTab[106218]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:134
		_go_fuzz_dep_.CoverTab[106219]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:134
		// _ = "end of CoverTab[106219]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:134
	// _ = "end of CoverTab[106215]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:134
	_go_fuzz_dep_.CoverTab[106216]++

												for _, val := range in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:136
		_go_fuzz_dep_.CoverTab[106220]++
													if err := re.putString(val); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:137
			_go_fuzz_dep_.CoverTab[106221]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:138
			// _ = "end of CoverTab[106221]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:139
			_go_fuzz_dep_.CoverTab[106222]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:139
			// _ = "end of CoverTab[106222]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:139
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:139
		// _ = "end of CoverTab[106220]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:140
	// _ = "end of CoverTab[106216]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:140
	_go_fuzz_dep_.CoverTab[106217]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:142
	// _ = "end of CoverTab[106217]"
}

func (re *realEncoder) putCompactInt32Array(in []int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:145
	_go_fuzz_dep_.CoverTab[106223]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:146
		_go_fuzz_dep_.CoverTab[106226]++
													return errors.New("expected int32 array to be non null")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:147
		// _ = "end of CoverTab[106226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:148
		_go_fuzz_dep_.CoverTab[106227]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:148
		// _ = "end of CoverTab[106227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:148
	// _ = "end of CoverTab[106223]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:148
	_go_fuzz_dep_.CoverTab[106224]++

												re.putUVarint(uint64(len(in)) + 1)
												for _, val := range in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:151
		_go_fuzz_dep_.CoverTab[106228]++
													re.putInt32(val)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:152
		// _ = "end of CoverTab[106228]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:153
	// _ = "end of CoverTab[106224]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:153
	_go_fuzz_dep_.CoverTab[106225]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:154
	// _ = "end of CoverTab[106225]"
}

func (re *realEncoder) putNullableCompactInt32Array(in []int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:157
	_go_fuzz_dep_.CoverTab[106229]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:158
		_go_fuzz_dep_.CoverTab[106232]++
													re.putUVarint(0)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:160
		// _ = "end of CoverTab[106232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:161
		_go_fuzz_dep_.CoverTab[106233]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:161
		// _ = "end of CoverTab[106233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:161
	// _ = "end of CoverTab[106229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:161
	_go_fuzz_dep_.CoverTab[106230]++

												re.putUVarint(uint64(len(in)) + 1)
												for _, val := range in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:164
		_go_fuzz_dep_.CoverTab[106234]++
													re.putInt32(val)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:165
		// _ = "end of CoverTab[106234]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:166
	// _ = "end of CoverTab[106230]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:166
	_go_fuzz_dep_.CoverTab[106231]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:167
	// _ = "end of CoverTab[106231]"
}

func (re *realEncoder) putInt32Array(in []int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:170
	_go_fuzz_dep_.CoverTab[106235]++
												err := re.putArrayLength(len(in))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:172
		_go_fuzz_dep_.CoverTab[106238]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:173
		// _ = "end of CoverTab[106238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:174
		_go_fuzz_dep_.CoverTab[106239]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:174
		// _ = "end of CoverTab[106239]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:174
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:174
	// _ = "end of CoverTab[106235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:174
	_go_fuzz_dep_.CoverTab[106236]++
												for _, val := range in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:175
		_go_fuzz_dep_.CoverTab[106240]++
													re.putInt32(val)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:176
		// _ = "end of CoverTab[106240]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:177
	// _ = "end of CoverTab[106236]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:177
	_go_fuzz_dep_.CoverTab[106237]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:178
	// _ = "end of CoverTab[106237]"
}

func (re *realEncoder) putInt64Array(in []int64) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:181
	_go_fuzz_dep_.CoverTab[106241]++
												err := re.putArrayLength(len(in))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:183
		_go_fuzz_dep_.CoverTab[106244]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:184
		// _ = "end of CoverTab[106244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:185
		_go_fuzz_dep_.CoverTab[106245]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:185
		// _ = "end of CoverTab[106245]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:185
	// _ = "end of CoverTab[106241]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:185
	_go_fuzz_dep_.CoverTab[106242]++
												for _, val := range in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:186
		_go_fuzz_dep_.CoverTab[106246]++
													re.putInt64(val)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:187
		// _ = "end of CoverTab[106246]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:188
	// _ = "end of CoverTab[106242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:188
	_go_fuzz_dep_.CoverTab[106243]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:189
	// _ = "end of CoverTab[106243]"
}

func (re *realEncoder) putEmptyTaggedFieldArray() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:192
	_go_fuzz_dep_.CoverTab[106247]++
												re.putUVarint(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:193
	// _ = "end of CoverTab[106247]"
}

func (re *realEncoder) offset() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:196
	_go_fuzz_dep_.CoverTab[106248]++
												return re.off
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:197
	// _ = "end of CoverTab[106248]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:202
func (re *realEncoder) push(in pushEncoder) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:202
	_go_fuzz_dep_.CoverTab[106249]++
												in.saveOffset(re.off)
												re.off += in.reserveLength()
												re.stack = append(re.stack, in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:205
	// _ = "end of CoverTab[106249]"
}

func (re *realEncoder) pop() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:208
	_go_fuzz_dep_.CoverTab[106250]++

												in := re.stack[len(re.stack)-1]
												re.stack = re.stack[:len(re.stack)-1]

												return in.run(re.off, re.raw)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:213
	// _ = "end of CoverTab[106250]"
}

// we do record metrics during the real encoder pass
func (re *realEncoder) metricRegistry() metrics.Registry {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:217
	_go_fuzz_dep_.CoverTab[106251]++
												return re.registry
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:218
	// _ = "end of CoverTab[106251]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:219
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/real_encoder.go:219
var _ = _go_fuzz_dep_.CoverTab
