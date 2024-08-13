//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:1
)

import (
	"encoding/binary"
	"time"
)

const (
	isTransactionalMask	= 0x10
	controlMask		= 0x20
	maximumRecordOverhead	= 5*binary.MaxVarintLen32 + binary.MaxVarintLen64 + 1
)

// RecordHeader stores key and value for a record header
type RecordHeader struct {
	Key	[]byte
	Value	[]byte
}

func (h *RecordHeader) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:20
	_go_fuzz_dep_.CoverTab[106252]++
											if err := pe.putVarintBytes(h.Key); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:21
		_go_fuzz_dep_.CoverTab[106254]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:22
		// _ = "end of CoverTab[106254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:23
		_go_fuzz_dep_.CoverTab[106255]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:23
		// _ = "end of CoverTab[106255]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:23
	// _ = "end of CoverTab[106252]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:23
	_go_fuzz_dep_.CoverTab[106253]++
											return pe.putVarintBytes(h.Value)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:24
	// _ = "end of CoverTab[106253]"
}

func (h *RecordHeader) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:27
	_go_fuzz_dep_.CoverTab[106256]++
											if h.Key, err = pd.getVarintBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:28
		_go_fuzz_dep_.CoverTab[106259]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:29
		// _ = "end of CoverTab[106259]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:30
		_go_fuzz_dep_.CoverTab[106260]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:30
		// _ = "end of CoverTab[106260]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:30
	// _ = "end of CoverTab[106256]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:30
	_go_fuzz_dep_.CoverTab[106257]++

											if h.Value, err = pd.getVarintBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:32
		_go_fuzz_dep_.CoverTab[106261]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:33
		// _ = "end of CoverTab[106261]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:34
		_go_fuzz_dep_.CoverTab[106262]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:34
		// _ = "end of CoverTab[106262]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:34
	// _ = "end of CoverTab[106257]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:34
	_go_fuzz_dep_.CoverTab[106258]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:35
	// _ = "end of CoverTab[106258]"
}

// Record is kafka record type
type Record struct {
	Headers	[]*RecordHeader

	Attributes	int8
	TimestampDelta	time.Duration
	OffsetDelta	int64
	Key		[]byte
	Value		[]byte
	length		varintLengthField
}

func (r *Record) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:50
	_go_fuzz_dep_.CoverTab[106263]++
											pe.push(&r.length)
											pe.putInt8(r.Attributes)
											pe.putVarint(int64(r.TimestampDelta / time.Millisecond))
											pe.putVarint(r.OffsetDelta)
											if err := pe.putVarintBytes(r.Key); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:55
		_go_fuzz_dep_.CoverTab[106267]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:56
		// _ = "end of CoverTab[106267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:57
		_go_fuzz_dep_.CoverTab[106268]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:57
		// _ = "end of CoverTab[106268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:57
	// _ = "end of CoverTab[106263]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:57
	_go_fuzz_dep_.CoverTab[106264]++
											if err := pe.putVarintBytes(r.Value); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:58
		_go_fuzz_dep_.CoverTab[106269]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:59
		// _ = "end of CoverTab[106269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:60
		_go_fuzz_dep_.CoverTab[106270]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:60
		// _ = "end of CoverTab[106270]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:60
	// _ = "end of CoverTab[106264]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:60
	_go_fuzz_dep_.CoverTab[106265]++
											pe.putVarint(int64(len(r.Headers)))

											for _, h := range r.Headers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:63
		_go_fuzz_dep_.CoverTab[106271]++
												if err := h.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:64
			_go_fuzz_dep_.CoverTab[106272]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:65
			// _ = "end of CoverTab[106272]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:66
			_go_fuzz_dep_.CoverTab[106273]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:66
			// _ = "end of CoverTab[106273]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:66
		// _ = "end of CoverTab[106271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:67
	// _ = "end of CoverTab[106265]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:67
	_go_fuzz_dep_.CoverTab[106266]++

											return pe.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:69
	// _ = "end of CoverTab[106266]"
}

func (r *Record) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:72
	_go_fuzz_dep_.CoverTab[106274]++
											if err = pd.push(&r.length); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:73
		_go_fuzz_dep_.CoverTab[106284]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:74
		// _ = "end of CoverTab[106284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:75
		_go_fuzz_dep_.CoverTab[106285]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:75
		// _ = "end of CoverTab[106285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:75
	// _ = "end of CoverTab[106274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:75
	_go_fuzz_dep_.CoverTab[106275]++

											if r.Attributes, err = pd.getInt8(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:77
		_go_fuzz_dep_.CoverTab[106286]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:78
		// _ = "end of CoverTab[106286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:79
		_go_fuzz_dep_.CoverTab[106287]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:79
		// _ = "end of CoverTab[106287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:79
	// _ = "end of CoverTab[106275]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:79
	_go_fuzz_dep_.CoverTab[106276]++

											timestamp, err := pd.getVarint()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:82
		_go_fuzz_dep_.CoverTab[106288]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:83
		// _ = "end of CoverTab[106288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:84
		_go_fuzz_dep_.CoverTab[106289]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:84
		// _ = "end of CoverTab[106289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:84
	// _ = "end of CoverTab[106276]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:84
	_go_fuzz_dep_.CoverTab[106277]++
											r.TimestampDelta = time.Duration(timestamp) * time.Millisecond

											if r.OffsetDelta, err = pd.getVarint(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:87
		_go_fuzz_dep_.CoverTab[106290]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:88
		// _ = "end of CoverTab[106290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:89
		_go_fuzz_dep_.CoverTab[106291]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:89
		// _ = "end of CoverTab[106291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:89
	// _ = "end of CoverTab[106277]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:89
	_go_fuzz_dep_.CoverTab[106278]++

											if r.Key, err = pd.getVarintBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:91
		_go_fuzz_dep_.CoverTab[106292]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:92
		// _ = "end of CoverTab[106292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:93
		_go_fuzz_dep_.CoverTab[106293]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:93
		// _ = "end of CoverTab[106293]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:93
	// _ = "end of CoverTab[106278]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:93
	_go_fuzz_dep_.CoverTab[106279]++

											if r.Value, err = pd.getVarintBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:95
		_go_fuzz_dep_.CoverTab[106294]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:96
		// _ = "end of CoverTab[106294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:97
		_go_fuzz_dep_.CoverTab[106295]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:97
		// _ = "end of CoverTab[106295]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:97
	// _ = "end of CoverTab[106279]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:97
	_go_fuzz_dep_.CoverTab[106280]++

											numHeaders, err := pd.getVarint()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:100
		_go_fuzz_dep_.CoverTab[106296]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:101
		// _ = "end of CoverTab[106296]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:102
		_go_fuzz_dep_.CoverTab[106297]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:102
		// _ = "end of CoverTab[106297]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:102
	// _ = "end of CoverTab[106280]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:102
	_go_fuzz_dep_.CoverTab[106281]++

											if numHeaders >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:104
		_go_fuzz_dep_.CoverTab[106298]++
												r.Headers = make([]*RecordHeader, numHeaders)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:105
		// _ = "end of CoverTab[106298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:106
		_go_fuzz_dep_.CoverTab[106299]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:106
		// _ = "end of CoverTab[106299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:106
	// _ = "end of CoverTab[106281]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:106
	_go_fuzz_dep_.CoverTab[106282]++
											for i := int64(0); i < numHeaders; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:107
		_go_fuzz_dep_.CoverTab[106300]++
												hdr := new(RecordHeader)
												if err := hdr.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:109
			_go_fuzz_dep_.CoverTab[106302]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:110
			// _ = "end of CoverTab[106302]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:111
			_go_fuzz_dep_.CoverTab[106303]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:111
			// _ = "end of CoverTab[106303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:111
		// _ = "end of CoverTab[106300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:111
		_go_fuzz_dep_.CoverTab[106301]++
												r.Headers[i] = hdr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:112
		// _ = "end of CoverTab[106301]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:113
	// _ = "end of CoverTab[106282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:113
	_go_fuzz_dep_.CoverTab[106283]++

											return pd.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:115
	// _ = "end of CoverTab[106283]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:116
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record.go:116
var _ = _go_fuzz_dep_.CoverTab
