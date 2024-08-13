//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:1
)

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/rcrowley/go-metrics"
)

type prepEncoder struct {
	stack	[]pushEncoder
	length	int
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:19
func (pe *prepEncoder) putInt8(in int8) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:19
	_go_fuzz_dep_.CoverTab[105591]++
												pe.length++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:20
	// _ = "end of CoverTab[105591]"
}

func (pe *prepEncoder) putInt16(in int16) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:23
	_go_fuzz_dep_.CoverTab[105592]++
												pe.length += 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:24
	// _ = "end of CoverTab[105592]"
}

func (pe *prepEncoder) putInt32(in int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:27
	_go_fuzz_dep_.CoverTab[105593]++
												pe.length += 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:28
	// _ = "end of CoverTab[105593]"
}

func (pe *prepEncoder) putInt64(in int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:31
	_go_fuzz_dep_.CoverTab[105594]++
												pe.length += 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:32
	// _ = "end of CoverTab[105594]"
}

func (pe *prepEncoder) putVarint(in int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:35
	_go_fuzz_dep_.CoverTab[105595]++
												var buf [binary.MaxVarintLen64]byte
												pe.length += binary.PutVarint(buf[:], in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:37
	// _ = "end of CoverTab[105595]"
}

func (pe *prepEncoder) putUVarint(in uint64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:40
	_go_fuzz_dep_.CoverTab[105596]++
												var buf [binary.MaxVarintLen64]byte
												pe.length += binary.PutUvarint(buf[:], in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:42
	// _ = "end of CoverTab[105596]"
}

func (pe *prepEncoder) putFloat64(in float64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:45
	_go_fuzz_dep_.CoverTab[105597]++
												pe.length += 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:46
	// _ = "end of CoverTab[105597]"
}

func (pe *prepEncoder) putArrayLength(in int) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:49
	_go_fuzz_dep_.CoverTab[105598]++
												if in > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:50
		_go_fuzz_dep_.CoverTab[105600]++
													return PacketEncodingError{fmt.Sprintf("array too long (%d)", in)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:51
		// _ = "end of CoverTab[105600]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:52
		_go_fuzz_dep_.CoverTab[105601]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:52
		// _ = "end of CoverTab[105601]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:52
	// _ = "end of CoverTab[105598]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:52
	_go_fuzz_dep_.CoverTab[105599]++
												pe.length += 4
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:54
	// _ = "end of CoverTab[105599]"
}

func (pe *prepEncoder) putCompactArrayLength(in int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:57
	_go_fuzz_dep_.CoverTab[105602]++
												pe.putUVarint(uint64(in + 1))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:58
	// _ = "end of CoverTab[105602]"
}

func (pe *prepEncoder) putBool(in bool) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:61
	_go_fuzz_dep_.CoverTab[105603]++
												pe.length++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:62
	// _ = "end of CoverTab[105603]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:67
func (pe *prepEncoder) putBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:67
	_go_fuzz_dep_.CoverTab[105604]++
												pe.length += 4
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:69
		_go_fuzz_dep_.CoverTab[105606]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:70
		// _ = "end of CoverTab[105606]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:71
		_go_fuzz_dep_.CoverTab[105607]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:71
		// _ = "end of CoverTab[105607]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:71
	// _ = "end of CoverTab[105604]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:71
	_go_fuzz_dep_.CoverTab[105605]++
												return pe.putRawBytes(in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:72
	// _ = "end of CoverTab[105605]"
}

func (pe *prepEncoder) putVarintBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:75
	_go_fuzz_dep_.CoverTab[105608]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:76
		_go_fuzz_dep_.CoverTab[105610]++
													pe.putVarint(-1)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:78
		// _ = "end of CoverTab[105610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:79
		_go_fuzz_dep_.CoverTab[105611]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:79
		// _ = "end of CoverTab[105611]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:79
	// _ = "end of CoverTab[105608]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:79
	_go_fuzz_dep_.CoverTab[105609]++
												pe.putVarint(int64(len(in)))
												return pe.putRawBytes(in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:81
	// _ = "end of CoverTab[105609]"
}

func (pe *prepEncoder) putCompactBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:84
	_go_fuzz_dep_.CoverTab[105612]++
												pe.putUVarint(uint64(len(in) + 1))
												return pe.putRawBytes(in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:86
	// _ = "end of CoverTab[105612]"
}

func (pe *prepEncoder) putCompactString(in string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:89
	_go_fuzz_dep_.CoverTab[105613]++
												pe.putCompactArrayLength(len(in))
												return pe.putRawBytes([]byte(in))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:91
	// _ = "end of CoverTab[105613]"
}

func (pe *prepEncoder) putNullableCompactString(in *string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:94
	_go_fuzz_dep_.CoverTab[105614]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:95
		_go_fuzz_dep_.CoverTab[105615]++
													pe.putUVarint(0)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:97
		// _ = "end of CoverTab[105615]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:98
		_go_fuzz_dep_.CoverTab[105616]++
													return pe.putCompactString(*in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:99
		// _ = "end of CoverTab[105616]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:100
	// _ = "end of CoverTab[105614]"
}

func (pe *prepEncoder) putRawBytes(in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:103
	_go_fuzz_dep_.CoverTab[105617]++
												if len(in) > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:104
		_go_fuzz_dep_.CoverTab[105619]++
													return PacketEncodingError{fmt.Sprintf("byteslice too long (%d)", len(in))}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:105
		// _ = "end of CoverTab[105619]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:106
		_go_fuzz_dep_.CoverTab[105620]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:106
		// _ = "end of CoverTab[105620]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:106
	// _ = "end of CoverTab[105617]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:106
	_go_fuzz_dep_.CoverTab[105618]++
												pe.length += len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:108
	// _ = "end of CoverTab[105618]"
}

func (pe *prepEncoder) putNullableString(in *string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:111
	_go_fuzz_dep_.CoverTab[105621]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:112
		_go_fuzz_dep_.CoverTab[105623]++
													pe.length += 2
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:114
		// _ = "end of CoverTab[105623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:115
		_go_fuzz_dep_.CoverTab[105624]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:115
		// _ = "end of CoverTab[105624]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:115
	// _ = "end of CoverTab[105621]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:115
	_go_fuzz_dep_.CoverTab[105622]++
												return pe.putString(*in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:116
	// _ = "end of CoverTab[105622]"
}

func (pe *prepEncoder) putString(in string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:119
	_go_fuzz_dep_.CoverTab[105625]++
												pe.length += 2
												if len(in) > math.MaxInt16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:121
		_go_fuzz_dep_.CoverTab[105627]++
													return PacketEncodingError{fmt.Sprintf("string too long (%d)", len(in))}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:122
		// _ = "end of CoverTab[105627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:123
		_go_fuzz_dep_.CoverTab[105628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:123
		// _ = "end of CoverTab[105628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:123
	// _ = "end of CoverTab[105625]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:123
	_go_fuzz_dep_.CoverTab[105626]++
												pe.length += len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:125
	// _ = "end of CoverTab[105626]"
}

func (pe *prepEncoder) putStringArray(in []string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:128
	_go_fuzz_dep_.CoverTab[105629]++
												err := pe.putArrayLength(len(in))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:130
		_go_fuzz_dep_.CoverTab[105632]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:131
		// _ = "end of CoverTab[105632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:132
		_go_fuzz_dep_.CoverTab[105633]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:132
		// _ = "end of CoverTab[105633]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:132
	// _ = "end of CoverTab[105629]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:132
	_go_fuzz_dep_.CoverTab[105630]++

												for _, str := range in {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:134
		_go_fuzz_dep_.CoverTab[105634]++
													if err := pe.putString(str); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:135
			_go_fuzz_dep_.CoverTab[105635]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:136
			// _ = "end of CoverTab[105635]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:137
			_go_fuzz_dep_.CoverTab[105636]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:137
			// _ = "end of CoverTab[105636]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:137
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:137
		// _ = "end of CoverTab[105634]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:138
	// _ = "end of CoverTab[105630]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:138
	_go_fuzz_dep_.CoverTab[105631]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:140
	// _ = "end of CoverTab[105631]"
}

func (pe *prepEncoder) putCompactInt32Array(in []int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:143
	_go_fuzz_dep_.CoverTab[105637]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:144
		_go_fuzz_dep_.CoverTab[105639]++
													return errors.New("expected int32 array to be non null")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:145
		// _ = "end of CoverTab[105639]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:146
		_go_fuzz_dep_.CoverTab[105640]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:146
		// _ = "end of CoverTab[105640]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:146
	// _ = "end of CoverTab[105637]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:146
	_go_fuzz_dep_.CoverTab[105638]++

												pe.putUVarint(uint64(len(in)) + 1)
												pe.length += 4 * len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:150
	// _ = "end of CoverTab[105638]"
}

func (pe *prepEncoder) putNullableCompactInt32Array(in []int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:153
	_go_fuzz_dep_.CoverTab[105641]++
												if in == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:154
		_go_fuzz_dep_.CoverTab[105643]++
													pe.putUVarint(0)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:156
		// _ = "end of CoverTab[105643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:157
		_go_fuzz_dep_.CoverTab[105644]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:157
		// _ = "end of CoverTab[105644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:157
	// _ = "end of CoverTab[105641]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:157
	_go_fuzz_dep_.CoverTab[105642]++

												pe.putUVarint(uint64(len(in)) + 1)
												pe.length += 4 * len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:161
	// _ = "end of CoverTab[105642]"
}

func (pe *prepEncoder) putInt32Array(in []int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:164
	_go_fuzz_dep_.CoverTab[105645]++
												err := pe.putArrayLength(len(in))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:166
		_go_fuzz_dep_.CoverTab[105647]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:167
		// _ = "end of CoverTab[105647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:168
		_go_fuzz_dep_.CoverTab[105648]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:168
		// _ = "end of CoverTab[105648]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:168
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:168
	// _ = "end of CoverTab[105645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:168
	_go_fuzz_dep_.CoverTab[105646]++
												pe.length += 4 * len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:170
	// _ = "end of CoverTab[105646]"
}

func (pe *prepEncoder) putInt64Array(in []int64) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:173
	_go_fuzz_dep_.CoverTab[105649]++
												err := pe.putArrayLength(len(in))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:175
		_go_fuzz_dep_.CoverTab[105651]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:176
		// _ = "end of CoverTab[105651]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:177
		_go_fuzz_dep_.CoverTab[105652]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:177
		// _ = "end of CoverTab[105652]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:177
	// _ = "end of CoverTab[105649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:177
	_go_fuzz_dep_.CoverTab[105650]++
												pe.length += 8 * len(in)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:179
	// _ = "end of CoverTab[105650]"
}

func (pe *prepEncoder) putEmptyTaggedFieldArray() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:182
	_go_fuzz_dep_.CoverTab[105653]++
												pe.putUVarint(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:183
	// _ = "end of CoverTab[105653]"
}

func (pe *prepEncoder) offset() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:186
	_go_fuzz_dep_.CoverTab[105654]++
												return pe.length
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:187
	// _ = "end of CoverTab[105654]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:192
func (pe *prepEncoder) push(in pushEncoder) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:192
	_go_fuzz_dep_.CoverTab[105655]++
												in.saveOffset(pe.length)
												pe.length += in.reserveLength()
												pe.stack = append(pe.stack, in)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:195
	// _ = "end of CoverTab[105655]"
}

func (pe *prepEncoder) pop() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:198
	_go_fuzz_dep_.CoverTab[105656]++
												in := pe.stack[len(pe.stack)-1]
												pe.stack = pe.stack[:len(pe.stack)-1]
												if dpe, ok := in.(dynamicPushEncoder); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:201
		_go_fuzz_dep_.CoverTab[105658]++
													pe.length += dpe.adjustLength(pe.length)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:202
		// _ = "end of CoverTab[105658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:203
		_go_fuzz_dep_.CoverTab[105659]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:203
		// _ = "end of CoverTab[105659]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:203
	// _ = "end of CoverTab[105656]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:203
	_go_fuzz_dep_.CoverTab[105657]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:205
	// _ = "end of CoverTab[105657]"
}

// we do not record metrics during the prep encoder pass
func (pe *prepEncoder) metricRegistry() metrics.Registry {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:209
	_go_fuzz_dep_.CoverTab[105660]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:210
	// _ = "end of CoverTab[105660]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:211
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/prep_encoder.go:211
var _ = _go_fuzz_dep_.CoverTab
