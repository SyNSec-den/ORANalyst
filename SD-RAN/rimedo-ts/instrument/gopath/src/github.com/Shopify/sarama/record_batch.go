//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:1
)

import (
	"fmt"
	"time"
)

const recordBatchOverhead = 49

type recordsArray []*Record

func (e recordsArray) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:12
	_go_fuzz_dep_.CoverTab[106304]++
												for _, r := range e {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:13
		_go_fuzz_dep_.CoverTab[106306]++
													if err := r.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:14
			_go_fuzz_dep_.CoverTab[106307]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:15
			// _ = "end of CoverTab[106307]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:16
			_go_fuzz_dep_.CoverTab[106308]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:16
			// _ = "end of CoverTab[106308]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:16
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:16
		// _ = "end of CoverTab[106306]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:17
	// _ = "end of CoverTab[106304]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:17
	_go_fuzz_dep_.CoverTab[106305]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:18
	// _ = "end of CoverTab[106305]"
}

func (e recordsArray) decode(pd packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:21
	_go_fuzz_dep_.CoverTab[106309]++
												for i := range e {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:22
		_go_fuzz_dep_.CoverTab[106311]++
													rec := &Record{}
													if err := rec.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:24
			_go_fuzz_dep_.CoverTab[106313]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:25
			// _ = "end of CoverTab[106313]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:26
			_go_fuzz_dep_.CoverTab[106314]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:26
			// _ = "end of CoverTab[106314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:26
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:26
		// _ = "end of CoverTab[106311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:26
		_go_fuzz_dep_.CoverTab[106312]++
													e[i] = rec
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:27
		// _ = "end of CoverTab[106312]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:28
	// _ = "end of CoverTab[106309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:28
	_go_fuzz_dep_.CoverTab[106310]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:29
	// _ = "end of CoverTab[106310]"
}

type RecordBatch struct {
	FirstOffset		int64
	PartitionLeaderEpoch	int32
	Version			int8
	Codec			CompressionCodec
	CompressionLevel	int
	Control			bool
	LogAppendTime		bool
	LastOffsetDelta		int32
	FirstTimestamp		time.Time
	MaxTimestamp		time.Time
	ProducerID		int64
	ProducerEpoch		int16
	FirstSequence		int32
	Records			[]*Record
	PartialTrailingRecord	bool
	IsTransactional		bool

	compressedRecords	[]byte
	recordsLen		int	// uncompressed records size
}

func (b *RecordBatch) LastOffset() int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:54
	_go_fuzz_dep_.CoverTab[106315]++
												return b.FirstOffset + int64(b.LastOffsetDelta)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:55
	// _ = "end of CoverTab[106315]"
}

func (b *RecordBatch) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:58
	_go_fuzz_dep_.CoverTab[106316]++
												if b.Version != 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:59
		_go_fuzz_dep_.CoverTab[106324]++
													return PacketEncodingError{fmt.Sprintf("unsupported compression codec (%d)", b.Codec)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:60
		// _ = "end of CoverTab[106324]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:61
		_go_fuzz_dep_.CoverTab[106325]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:61
		// _ = "end of CoverTab[106325]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:61
	// _ = "end of CoverTab[106316]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:61
	_go_fuzz_dep_.CoverTab[106317]++
												pe.putInt64(b.FirstOffset)
												pe.push(&lengthField{})
												pe.putInt32(b.PartitionLeaderEpoch)
												pe.putInt8(b.Version)
												pe.push(newCRC32Field(crcCastagnoli))
												pe.putInt16(b.computeAttributes())
												pe.putInt32(b.LastOffsetDelta)

												if err := (Timestamp{&b.FirstTimestamp}).encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:70
		_go_fuzz_dep_.CoverTab[106326]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:71
		// _ = "end of CoverTab[106326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:72
		_go_fuzz_dep_.CoverTab[106327]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:72
		// _ = "end of CoverTab[106327]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:72
	// _ = "end of CoverTab[106317]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:72
	_go_fuzz_dep_.CoverTab[106318]++

												if err := (Timestamp{&b.MaxTimestamp}).encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:74
		_go_fuzz_dep_.CoverTab[106328]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:75
		// _ = "end of CoverTab[106328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:76
		_go_fuzz_dep_.CoverTab[106329]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:76
		// _ = "end of CoverTab[106329]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:76
	// _ = "end of CoverTab[106318]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:76
	_go_fuzz_dep_.CoverTab[106319]++

												pe.putInt64(b.ProducerID)
												pe.putInt16(b.ProducerEpoch)
												pe.putInt32(b.FirstSequence)

												if err := pe.putArrayLength(len(b.Records)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:82
		_go_fuzz_dep_.CoverTab[106330]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:83
		// _ = "end of CoverTab[106330]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:84
		_go_fuzz_dep_.CoverTab[106331]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:84
		// _ = "end of CoverTab[106331]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:84
	// _ = "end of CoverTab[106319]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:84
	_go_fuzz_dep_.CoverTab[106320]++

												if b.compressedRecords == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:86
		_go_fuzz_dep_.CoverTab[106332]++
													if err := b.encodeRecords(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:87
			_go_fuzz_dep_.CoverTab[106333]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:88
			// _ = "end of CoverTab[106333]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:89
			_go_fuzz_dep_.CoverTab[106334]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:89
			// _ = "end of CoverTab[106334]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:89
		// _ = "end of CoverTab[106332]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:90
		_go_fuzz_dep_.CoverTab[106335]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:90
		// _ = "end of CoverTab[106335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:90
	// _ = "end of CoverTab[106320]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:90
	_go_fuzz_dep_.CoverTab[106321]++
												if err := pe.putRawBytes(b.compressedRecords); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:91
		_go_fuzz_dep_.CoverTab[106336]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:92
		// _ = "end of CoverTab[106336]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:93
		_go_fuzz_dep_.CoverTab[106337]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:93
		// _ = "end of CoverTab[106337]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:93
	// _ = "end of CoverTab[106321]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:93
	_go_fuzz_dep_.CoverTab[106322]++

												if err := pe.pop(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:95
		_go_fuzz_dep_.CoverTab[106338]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:96
		// _ = "end of CoverTab[106338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:97
		_go_fuzz_dep_.CoverTab[106339]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:97
		// _ = "end of CoverTab[106339]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:97
	// _ = "end of CoverTab[106322]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:97
	_go_fuzz_dep_.CoverTab[106323]++
												return pe.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:98
	// _ = "end of CoverTab[106323]"
}

func (b *RecordBatch) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:101
	_go_fuzz_dep_.CoverTab[106340]++
												if b.FirstOffset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:102
		_go_fuzz_dep_.CoverTab[106359]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:103
		// _ = "end of CoverTab[106359]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:104
		_go_fuzz_dep_.CoverTab[106360]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:104
		// _ = "end of CoverTab[106360]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:104
	// _ = "end of CoverTab[106340]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:104
	_go_fuzz_dep_.CoverTab[106341]++

												batchLen, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:107
		_go_fuzz_dep_.CoverTab[106361]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:108
		// _ = "end of CoverTab[106361]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:109
		_go_fuzz_dep_.CoverTab[106362]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:109
		// _ = "end of CoverTab[106362]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:109
	// _ = "end of CoverTab[106341]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:109
	_go_fuzz_dep_.CoverTab[106342]++

												if b.PartitionLeaderEpoch, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:111
		_go_fuzz_dep_.CoverTab[106363]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:112
		// _ = "end of CoverTab[106363]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:113
		_go_fuzz_dep_.CoverTab[106364]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:113
		// _ = "end of CoverTab[106364]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:113
	// _ = "end of CoverTab[106342]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:113
	_go_fuzz_dep_.CoverTab[106343]++

												if b.Version, err = pd.getInt8(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:115
		_go_fuzz_dep_.CoverTab[106365]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:116
		// _ = "end of CoverTab[106365]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:117
		_go_fuzz_dep_.CoverTab[106366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:117
		// _ = "end of CoverTab[106366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:117
	// _ = "end of CoverTab[106343]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:117
	_go_fuzz_dep_.CoverTab[106344]++

												crc32Decoder := acquireCrc32Field(crcCastagnoli)
												defer releaseCrc32Field(crc32Decoder)

												if err = pd.push(crc32Decoder); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:122
		_go_fuzz_dep_.CoverTab[106367]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:123
		// _ = "end of CoverTab[106367]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:124
		_go_fuzz_dep_.CoverTab[106368]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:124
		// _ = "end of CoverTab[106368]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:124
	// _ = "end of CoverTab[106344]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:124
	_go_fuzz_dep_.CoverTab[106345]++

												attributes, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:127
		_go_fuzz_dep_.CoverTab[106369]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:128
		// _ = "end of CoverTab[106369]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:129
		_go_fuzz_dep_.CoverTab[106370]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:129
		// _ = "end of CoverTab[106370]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:129
	// _ = "end of CoverTab[106345]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:129
	_go_fuzz_dep_.CoverTab[106346]++
												b.Codec = CompressionCodec(int8(attributes) & compressionCodecMask)
												b.Control = attributes&controlMask == controlMask
												b.LogAppendTime = attributes&timestampTypeMask == timestampTypeMask
												b.IsTransactional = attributes&isTransactionalMask == isTransactionalMask

												if b.LastOffsetDelta, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:135
		_go_fuzz_dep_.CoverTab[106371]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:136
		// _ = "end of CoverTab[106371]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:137
		_go_fuzz_dep_.CoverTab[106372]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:137
		// _ = "end of CoverTab[106372]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:137
	// _ = "end of CoverTab[106346]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:137
	_go_fuzz_dep_.CoverTab[106347]++

												if err = (Timestamp{&b.FirstTimestamp}).decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:139
		_go_fuzz_dep_.CoverTab[106373]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:140
		// _ = "end of CoverTab[106373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:141
		_go_fuzz_dep_.CoverTab[106374]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:141
		// _ = "end of CoverTab[106374]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:141
	// _ = "end of CoverTab[106347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:141
	_go_fuzz_dep_.CoverTab[106348]++

												if err = (Timestamp{&b.MaxTimestamp}).decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:143
		_go_fuzz_dep_.CoverTab[106375]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:144
		// _ = "end of CoverTab[106375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:145
		_go_fuzz_dep_.CoverTab[106376]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:145
		// _ = "end of CoverTab[106376]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:145
	// _ = "end of CoverTab[106348]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:145
	_go_fuzz_dep_.CoverTab[106349]++

												if b.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:147
		_go_fuzz_dep_.CoverTab[106377]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:148
		// _ = "end of CoverTab[106377]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:149
		_go_fuzz_dep_.CoverTab[106378]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:149
		// _ = "end of CoverTab[106378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:149
	// _ = "end of CoverTab[106349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:149
	_go_fuzz_dep_.CoverTab[106350]++

												if b.ProducerEpoch, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:151
		_go_fuzz_dep_.CoverTab[106379]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:152
		// _ = "end of CoverTab[106379]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:153
		_go_fuzz_dep_.CoverTab[106380]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:153
		// _ = "end of CoverTab[106380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:153
	// _ = "end of CoverTab[106350]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:153
	_go_fuzz_dep_.CoverTab[106351]++

												if b.FirstSequence, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:155
		_go_fuzz_dep_.CoverTab[106381]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:156
		// _ = "end of CoverTab[106381]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:157
		_go_fuzz_dep_.CoverTab[106382]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:157
		// _ = "end of CoverTab[106382]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:157
	// _ = "end of CoverTab[106351]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:157
	_go_fuzz_dep_.CoverTab[106352]++

												numRecs, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:160
		_go_fuzz_dep_.CoverTab[106383]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:161
		// _ = "end of CoverTab[106383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:162
		_go_fuzz_dep_.CoverTab[106384]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:162
		// _ = "end of CoverTab[106384]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:162
	// _ = "end of CoverTab[106352]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:162
	_go_fuzz_dep_.CoverTab[106353]++
												if numRecs >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:163
		_go_fuzz_dep_.CoverTab[106385]++
													b.Records = make([]*Record, numRecs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:164
		// _ = "end of CoverTab[106385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:165
		_go_fuzz_dep_.CoverTab[106386]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:165
		// _ = "end of CoverTab[106386]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:165
	// _ = "end of CoverTab[106353]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:165
	_go_fuzz_dep_.CoverTab[106354]++

												bufSize := int(batchLen) - recordBatchOverhead
												recBuffer, err := pd.getRawBytes(bufSize)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:169
		_go_fuzz_dep_.CoverTab[106387]++
													if err == ErrInsufficientData {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:170
			_go_fuzz_dep_.CoverTab[106389]++
														b.PartialTrailingRecord = true
														b.Records = nil
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:173
			// _ = "end of CoverTab[106389]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:174
			_go_fuzz_dep_.CoverTab[106390]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:174
			// _ = "end of CoverTab[106390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:174
		// _ = "end of CoverTab[106387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:174
		_go_fuzz_dep_.CoverTab[106388]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:175
		// _ = "end of CoverTab[106388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:176
		_go_fuzz_dep_.CoverTab[106391]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:176
		// _ = "end of CoverTab[106391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:176
	// _ = "end of CoverTab[106354]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:176
	_go_fuzz_dep_.CoverTab[106355]++

												if err = pd.pop(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:178
		_go_fuzz_dep_.CoverTab[106392]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:179
		// _ = "end of CoverTab[106392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:180
		_go_fuzz_dep_.CoverTab[106393]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:180
		// _ = "end of CoverTab[106393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:180
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:180
	// _ = "end of CoverTab[106355]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:180
	_go_fuzz_dep_.CoverTab[106356]++

												recBuffer, err = decompress(b.Codec, recBuffer)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:183
		_go_fuzz_dep_.CoverTab[106394]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:184
		// _ = "end of CoverTab[106394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:185
		_go_fuzz_dep_.CoverTab[106395]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:185
		// _ = "end of CoverTab[106395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:185
	// _ = "end of CoverTab[106356]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:185
	_go_fuzz_dep_.CoverTab[106357]++

												b.recordsLen = len(recBuffer)
												err = decode(recBuffer, recordsArray(b.Records))
												if err == ErrInsufficientData {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:189
		_go_fuzz_dep_.CoverTab[106396]++
													b.PartialTrailingRecord = true
													b.Records = nil
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:192
		// _ = "end of CoverTab[106396]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:193
		_go_fuzz_dep_.CoverTab[106397]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:193
		// _ = "end of CoverTab[106397]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:193
	// _ = "end of CoverTab[106357]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:193
	_go_fuzz_dep_.CoverTab[106358]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:194
	// _ = "end of CoverTab[106358]"
}

func (b *RecordBatch) encodeRecords(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:197
	_go_fuzz_dep_.CoverTab[106398]++
												var raw []byte
												var err error
												if raw, err = encode(recordsArray(b.Records), pe.metricRegistry()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:200
		_go_fuzz_dep_.CoverTab[106400]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:201
		// _ = "end of CoverTab[106400]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:202
		_go_fuzz_dep_.CoverTab[106401]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:202
		// _ = "end of CoverTab[106401]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:202
	// _ = "end of CoverTab[106398]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:202
	_go_fuzz_dep_.CoverTab[106399]++
												b.recordsLen = len(raw)

												b.compressedRecords, err = compress(b.Codec, b.CompressionLevel, raw)
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:206
	// _ = "end of CoverTab[106399]"
}

func (b *RecordBatch) computeAttributes() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:209
	_go_fuzz_dep_.CoverTab[106402]++
												attr := int16(b.Codec) & int16(compressionCodecMask)
												if b.Control {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:211
		_go_fuzz_dep_.CoverTab[106406]++
													attr |= controlMask
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:212
		// _ = "end of CoverTab[106406]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:213
		_go_fuzz_dep_.CoverTab[106407]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:213
		// _ = "end of CoverTab[106407]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:213
	// _ = "end of CoverTab[106402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:213
	_go_fuzz_dep_.CoverTab[106403]++
												if b.LogAppendTime {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:214
		_go_fuzz_dep_.CoverTab[106408]++
													attr |= timestampTypeMask
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:215
		// _ = "end of CoverTab[106408]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:216
		_go_fuzz_dep_.CoverTab[106409]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:216
		// _ = "end of CoverTab[106409]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:216
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:216
	// _ = "end of CoverTab[106403]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:216
	_go_fuzz_dep_.CoverTab[106404]++
												if b.IsTransactional {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:217
		_go_fuzz_dep_.CoverTab[106410]++
													attr |= isTransactionalMask
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:218
		// _ = "end of CoverTab[106410]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:219
		_go_fuzz_dep_.CoverTab[106411]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:219
		// _ = "end of CoverTab[106411]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:219
	// _ = "end of CoverTab[106404]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:219
	_go_fuzz_dep_.CoverTab[106405]++
												return attr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:220
	// _ = "end of CoverTab[106405]"
}

func (b *RecordBatch) addRecord(r *Record) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:223
	_go_fuzz_dep_.CoverTab[106412]++
												b.Records = append(b.Records, r)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:224
	// _ = "end of CoverTab[106412]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:225
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/record_batch.go:225
var _ = _go_fuzz_dep_.CoverTab
