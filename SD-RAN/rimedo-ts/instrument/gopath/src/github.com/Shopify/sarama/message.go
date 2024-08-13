//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:1
)

import (
	"fmt"
	"time"
)

const (
	// CompressionNone no compression
	CompressionNone	CompressionCodec	= iota
	// CompressionGZIP compression using GZIP
	CompressionGZIP
	// CompressionSnappy compression using snappy
	CompressionSnappy
	// CompressionLZ4 compression using LZ4
	CompressionLZ4
	// CompressionZSTD compression using ZSTD
	CompressionZSTD

	// The lowest 3 bits contain the compression codec used for the message
	compressionCodecMask	int8	= 0x07

	// Bit 3 set for "LogAppend" timestamps
	timestampTypeMask	= 0x08

	// CompressionLevelDefault is the constant to use in CompressionLevel
	// to have the default compression level for any codec. The value is picked
	// that we don't use any existing compression levels.
	CompressionLevelDefault	= -1000
)

// CompressionCodec represents the various compression codecs recognized by Kafka in messages.
type CompressionCodec int8

func (cc CompressionCodec) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:35
	_go_fuzz_dep_.CoverTab[103950]++
											return []string{
		"none",
		"gzip",
		"snappy",
		"lz4",
		"zstd",
	}[int(cc)]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:42
	// _ = "end of CoverTab[103950]"
}

// Message is a kafka message type
type Message struct {
	Codec			CompressionCodec	// codec used to compress the message contents
	CompressionLevel	int			// compression level
	LogAppendTime		bool			// the used timestamp is LogAppendTime
	Key			[]byte			// the message key, may be nil
	Value			[]byte			// the message contents
	Set			*MessageSet		// the message set a message might wrap
	Version			int8			// v1 requires Kafka 0.10
	Timestamp		time.Time		// the timestamp of the message (version 1+ only)

	compressedCache	[]byte
	compressedSize	int	// used for computing the compression ratio metrics
}

func (m *Message) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:60
	_go_fuzz_dep_.CoverTab[103951]++
											pe.push(newCRC32Field(crcIEEE))

											pe.putInt8(m.Version)

											attributes := int8(m.Codec) & compressionCodecMask
											if m.LogAppendTime {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:66
		_go_fuzz_dep_.CoverTab[103957]++
												attributes |= timestampTypeMask
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:67
		// _ = "end of CoverTab[103957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:68
		_go_fuzz_dep_.CoverTab[103958]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:68
		// _ = "end of CoverTab[103958]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:68
	// _ = "end of CoverTab[103951]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:68
	_go_fuzz_dep_.CoverTab[103952]++
											pe.putInt8(attributes)

											if m.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:71
		_go_fuzz_dep_.CoverTab[103959]++
												if err := (Timestamp{&m.Timestamp}).encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:72
			_go_fuzz_dep_.CoverTab[103960]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:73
			// _ = "end of CoverTab[103960]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:74
			_go_fuzz_dep_.CoverTab[103961]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:74
			// _ = "end of CoverTab[103961]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:74
		// _ = "end of CoverTab[103959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:75
		_go_fuzz_dep_.CoverTab[103962]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:75
		// _ = "end of CoverTab[103962]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:75
	// _ = "end of CoverTab[103952]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:75
	_go_fuzz_dep_.CoverTab[103953]++

											err := pe.putBytes(m.Key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:78
		_go_fuzz_dep_.CoverTab[103963]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:79
		// _ = "end of CoverTab[103963]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:80
		_go_fuzz_dep_.CoverTab[103964]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:80
		// _ = "end of CoverTab[103964]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:80
	// _ = "end of CoverTab[103953]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:80
	_go_fuzz_dep_.CoverTab[103954]++

											var payload []byte

											if m.compressedCache != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:84
		_go_fuzz_dep_.CoverTab[103965]++
												payload = m.compressedCache
												m.compressedCache = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:86
		// _ = "end of CoverTab[103965]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:87
		_go_fuzz_dep_.CoverTab[103966]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:87
		if m.Value != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:87
			_go_fuzz_dep_.CoverTab[103967]++
													payload, err = compress(m.Codec, m.CompressionLevel, m.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:89
				_go_fuzz_dep_.CoverTab[103969]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:90
				// _ = "end of CoverTab[103969]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:91
				_go_fuzz_dep_.CoverTab[103970]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:91
				// _ = "end of CoverTab[103970]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:91
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:91
			// _ = "end of CoverTab[103967]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:91
			_go_fuzz_dep_.CoverTab[103968]++
													m.compressedCache = payload

													m.compressedSize = len(payload)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:94
			// _ = "end of CoverTab[103968]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
			_go_fuzz_dep_.CoverTab[103971]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
			// _ = "end of CoverTab[103971]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
		// _ = "end of CoverTab[103966]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
	// _ = "end of CoverTab[103954]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:95
	_go_fuzz_dep_.CoverTab[103955]++

											if err = pe.putBytes(payload); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:97
		_go_fuzz_dep_.CoverTab[103972]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:98
		// _ = "end of CoverTab[103972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:99
		_go_fuzz_dep_.CoverTab[103973]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:99
		// _ = "end of CoverTab[103973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:99
	// _ = "end of CoverTab[103955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:99
	_go_fuzz_dep_.CoverTab[103956]++

											return pe.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:101
	// _ = "end of CoverTab[103956]"
}

func (m *Message) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:104
	_go_fuzz_dep_.CoverTab[103974]++
											crc32Decoder := acquireCrc32Field(crcIEEE)
											defer releaseCrc32Field(crc32Decoder)

											err = pd.push(crc32Decoder)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:109
		_go_fuzz_dep_.CoverTab[103983]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:110
		// _ = "end of CoverTab[103983]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:111
		_go_fuzz_dep_.CoverTab[103984]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:111
		// _ = "end of CoverTab[103984]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:111
	// _ = "end of CoverTab[103974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:111
	_go_fuzz_dep_.CoverTab[103975]++

											m.Version, err = pd.getInt8()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:114
		_go_fuzz_dep_.CoverTab[103985]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:115
		// _ = "end of CoverTab[103985]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:116
		_go_fuzz_dep_.CoverTab[103986]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:116
		// _ = "end of CoverTab[103986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:116
	// _ = "end of CoverTab[103975]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:116
	_go_fuzz_dep_.CoverTab[103976]++

											if m.Version > 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:118
		_go_fuzz_dep_.CoverTab[103987]++
												return PacketDecodingError{fmt.Sprintf("unknown magic byte (%v)", m.Version)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:119
		// _ = "end of CoverTab[103987]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:120
		_go_fuzz_dep_.CoverTab[103988]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:120
		// _ = "end of CoverTab[103988]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:120
	// _ = "end of CoverTab[103976]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:120
	_go_fuzz_dep_.CoverTab[103977]++

											attribute, err := pd.getInt8()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:123
		_go_fuzz_dep_.CoverTab[103989]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:124
		// _ = "end of CoverTab[103989]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:125
		_go_fuzz_dep_.CoverTab[103990]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:125
		// _ = "end of CoverTab[103990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:125
	// _ = "end of CoverTab[103977]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:125
	_go_fuzz_dep_.CoverTab[103978]++
											m.Codec = CompressionCodec(attribute & compressionCodecMask)
											m.LogAppendTime = attribute&timestampTypeMask == timestampTypeMask

											if m.Version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:129
		_go_fuzz_dep_.CoverTab[103991]++
												if err := (Timestamp{&m.Timestamp}).decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:130
			_go_fuzz_dep_.CoverTab[103992]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:131
			// _ = "end of CoverTab[103992]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:132
			_go_fuzz_dep_.CoverTab[103993]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:132
			// _ = "end of CoverTab[103993]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:132
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:132
		// _ = "end of CoverTab[103991]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:133
		_go_fuzz_dep_.CoverTab[103994]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:133
		// _ = "end of CoverTab[103994]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:133
	// _ = "end of CoverTab[103978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:133
	_go_fuzz_dep_.CoverTab[103979]++

											m.Key, err = pd.getBytes()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:136
		_go_fuzz_dep_.CoverTab[103995]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:137
		// _ = "end of CoverTab[103995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:138
		_go_fuzz_dep_.CoverTab[103996]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:138
		// _ = "end of CoverTab[103996]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:138
	// _ = "end of CoverTab[103979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:138
	_go_fuzz_dep_.CoverTab[103980]++

											m.Value, err = pd.getBytes()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:141
		_go_fuzz_dep_.CoverTab[103997]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:142
		// _ = "end of CoverTab[103997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:143
		_go_fuzz_dep_.CoverTab[103998]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:143
		// _ = "end of CoverTab[103998]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:143
	// _ = "end of CoverTab[103980]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:143
	_go_fuzz_dep_.CoverTab[103981]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:147
	m.compressedSize = len(m.Value)

	if m.Value != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:149
		_go_fuzz_dep_.CoverTab[103999]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:149
		return m.Codec != CompressionNone
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:149
		// _ = "end of CoverTab[103999]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:149
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:149
		_go_fuzz_dep_.CoverTab[104000]++
												m.Value, err = decompress(m.Codec, m.Value)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:151
			_go_fuzz_dep_.CoverTab[104002]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:152
			// _ = "end of CoverTab[104002]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:153
			_go_fuzz_dep_.CoverTab[104003]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:153
			// _ = "end of CoverTab[104003]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:153
		// _ = "end of CoverTab[104000]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:153
		_go_fuzz_dep_.CoverTab[104001]++

												if err := m.decodeSet(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:155
			_go_fuzz_dep_.CoverTab[104004]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:156
			// _ = "end of CoverTab[104004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:157
			_go_fuzz_dep_.CoverTab[104005]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:157
			// _ = "end of CoverTab[104005]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:157
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:157
		// _ = "end of CoverTab[104001]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:158
		_go_fuzz_dep_.CoverTab[104006]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:158
		// _ = "end of CoverTab[104006]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:158
	// _ = "end of CoverTab[103981]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:158
	_go_fuzz_dep_.CoverTab[103982]++

											return pd.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:160
	// _ = "end of CoverTab[103982]"
}

// decodes a message set from a previously encoded bulk-message
func (m *Message) decodeSet() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:164
	_go_fuzz_dep_.CoverTab[104007]++
											pd := realDecoder{raw: m.Value}
											m.Set = &MessageSet{}
											return m.Set.decode(&pd)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:167
	// _ = "end of CoverTab[104007]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message.go:168
var _ = _go_fuzz_dep_.CoverTab
