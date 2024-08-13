//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:1
)

import "errors"

// ReceiveTime is a special value for the timestamp field of Offset Commit Requests which
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:5
// tells the broker to set the timestamp to the time at which the request was received.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:5
// The timestamp is only used if message version 1 is used, which requires kafka 0.8.2.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:8
const ReceiveTime int64 = -1

// GroupGenerationUndefined is a special value for the group generation field of
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:10
// Offset Commit Requests that should be used when a consumer group does not rely
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:10
// on Kafka for partition management.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:13
const GroupGenerationUndefined = -1

type offsetCommitRequestBlock struct {
	offset		int64
	timestamp	int64
	metadata	string
}

func (b *offsetCommitRequestBlock) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:21
	_go_fuzz_dep_.CoverTab[104752]++
													pe.putInt64(b.offset)
													if version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:23
		_go_fuzz_dep_.CoverTab[104754]++
														pe.putInt64(b.timestamp)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:24
		// _ = "end of CoverTab[104754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:25
		_go_fuzz_dep_.CoverTab[104755]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:25
		if b.timestamp != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:25
			_go_fuzz_dep_.CoverTab[104756]++
															Logger.Println("Non-zero timestamp specified for OffsetCommitRequest not v1, it will be ignored")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:26
			// _ = "end of CoverTab[104756]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
			_go_fuzz_dep_.CoverTab[104757]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
			// _ = "end of CoverTab[104757]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
		// _ = "end of CoverTab[104755]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
	// _ = "end of CoverTab[104752]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:27
	_go_fuzz_dep_.CoverTab[104753]++

													return pe.putString(b.metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:29
	// _ = "end of CoverTab[104753]"
}

func (b *offsetCommitRequestBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:32
	_go_fuzz_dep_.CoverTab[104758]++
													if b.offset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:33
		_go_fuzz_dep_.CoverTab[104761]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:34
		// _ = "end of CoverTab[104761]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:35
		_go_fuzz_dep_.CoverTab[104762]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:35
		// _ = "end of CoverTab[104762]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:35
	// _ = "end of CoverTab[104758]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:35
	_go_fuzz_dep_.CoverTab[104759]++
													if version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:36
		_go_fuzz_dep_.CoverTab[104763]++
														if b.timestamp, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:37
			_go_fuzz_dep_.CoverTab[104764]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:38
			// _ = "end of CoverTab[104764]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:39
			_go_fuzz_dep_.CoverTab[104765]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:39
			// _ = "end of CoverTab[104765]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:39
		// _ = "end of CoverTab[104763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:40
		_go_fuzz_dep_.CoverTab[104766]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:40
		// _ = "end of CoverTab[104766]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:40
	// _ = "end of CoverTab[104759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:40
	_go_fuzz_dep_.CoverTab[104760]++
													b.metadata, err = pd.getString()
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:42
	// _ = "end of CoverTab[104760]"
}

type OffsetCommitRequest struct {
	ConsumerGroup		string
	ConsumerGroupGeneration	int32	// v1 or later
	ConsumerID		string	// v1 or later
	RetentionTime		int64	// v2 or later

	// Version can be:
	// - 0 (kafka 0.8.1 and later)
	// - 1 (kafka 0.8.2 and later)
	// - 2 (kafka 0.9.0 and later)
	// - 3 (kafka 0.11.0 and later)
	// - 4 (kafka 2.0.0 and later)
	Version	int16
	blocks	map[string]map[int32]*offsetCommitRequestBlock
}

func (r *OffsetCommitRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:61
	_go_fuzz_dep_.CoverTab[104767]++
													if r.Version < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:62
		_go_fuzz_dep_.CoverTab[104774]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:62
		return r.Version > 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:62
		// _ = "end of CoverTab[104774]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:62
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:62
		_go_fuzz_dep_.CoverTab[104775]++
														return PacketEncodingError{"invalid or unsupported OffsetCommitRequest version field"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:63
		// _ = "end of CoverTab[104775]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:64
		_go_fuzz_dep_.CoverTab[104776]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:64
		// _ = "end of CoverTab[104776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:64
	// _ = "end of CoverTab[104767]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:64
	_go_fuzz_dep_.CoverTab[104768]++

													if err := pe.putString(r.ConsumerGroup); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:66
		_go_fuzz_dep_.CoverTab[104777]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:67
		// _ = "end of CoverTab[104777]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:68
		_go_fuzz_dep_.CoverTab[104778]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:68
		// _ = "end of CoverTab[104778]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:68
	// _ = "end of CoverTab[104768]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:68
	_go_fuzz_dep_.CoverTab[104769]++

													if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:70
		_go_fuzz_dep_.CoverTab[104779]++
														pe.putInt32(r.ConsumerGroupGeneration)
														if err := pe.putString(r.ConsumerID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:72
			_go_fuzz_dep_.CoverTab[104780]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:73
			// _ = "end of CoverTab[104780]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:74
			_go_fuzz_dep_.CoverTab[104781]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:74
			// _ = "end of CoverTab[104781]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:74
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:74
		// _ = "end of CoverTab[104779]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:75
		_go_fuzz_dep_.CoverTab[104782]++
														if r.ConsumerGroupGeneration != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:76
			_go_fuzz_dep_.CoverTab[104784]++
															Logger.Println("Non-zero ConsumerGroupGeneration specified for OffsetCommitRequest v0, it will be ignored")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:77
			// _ = "end of CoverTab[104784]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:78
			_go_fuzz_dep_.CoverTab[104785]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:78
			// _ = "end of CoverTab[104785]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:78
		// _ = "end of CoverTab[104782]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:78
		_go_fuzz_dep_.CoverTab[104783]++
														if r.ConsumerID != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:79
			_go_fuzz_dep_.CoverTab[104786]++
															Logger.Println("Non-empty ConsumerID specified for OffsetCommitRequest v0, it will be ignored")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:80
			// _ = "end of CoverTab[104786]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:81
			_go_fuzz_dep_.CoverTab[104787]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:81
			// _ = "end of CoverTab[104787]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:81
		// _ = "end of CoverTab[104783]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:82
	// _ = "end of CoverTab[104769]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:82
	_go_fuzz_dep_.CoverTab[104770]++

													if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:84
		_go_fuzz_dep_.CoverTab[104788]++
														pe.putInt64(r.RetentionTime)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:85
		// _ = "end of CoverTab[104788]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:86
		_go_fuzz_dep_.CoverTab[104789]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:86
		if r.RetentionTime != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:86
			_go_fuzz_dep_.CoverTab[104790]++
															Logger.Println("Non-zero RetentionTime specified for OffsetCommitRequest version <2, it will be ignored")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:87
			// _ = "end of CoverTab[104790]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
			_go_fuzz_dep_.CoverTab[104791]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
			// _ = "end of CoverTab[104791]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
		// _ = "end of CoverTab[104789]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
	// _ = "end of CoverTab[104770]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:88
	_go_fuzz_dep_.CoverTab[104771]++

													if err := pe.putArrayLength(len(r.blocks)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:90
		_go_fuzz_dep_.CoverTab[104792]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:91
		// _ = "end of CoverTab[104792]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:92
		_go_fuzz_dep_.CoverTab[104793]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:92
		// _ = "end of CoverTab[104793]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:92
	// _ = "end of CoverTab[104771]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:92
	_go_fuzz_dep_.CoverTab[104772]++
													for topic, partitions := range r.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:93
		_go_fuzz_dep_.CoverTab[104794]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:94
			_go_fuzz_dep_.CoverTab[104797]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:95
			// _ = "end of CoverTab[104797]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:96
			_go_fuzz_dep_.CoverTab[104798]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:96
			// _ = "end of CoverTab[104798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:96
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:96
		// _ = "end of CoverTab[104794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:96
		_go_fuzz_dep_.CoverTab[104795]++
														if err := pe.putArrayLength(len(partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:97
			_go_fuzz_dep_.CoverTab[104799]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:98
			// _ = "end of CoverTab[104799]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:99
			_go_fuzz_dep_.CoverTab[104800]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:99
			// _ = "end of CoverTab[104800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:99
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:99
		// _ = "end of CoverTab[104795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:99
		_go_fuzz_dep_.CoverTab[104796]++
														for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:100
			_go_fuzz_dep_.CoverTab[104801]++
															pe.putInt32(partition)
															if err := block.encode(pe, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:102
				_go_fuzz_dep_.CoverTab[104802]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:103
				// _ = "end of CoverTab[104802]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:104
				_go_fuzz_dep_.CoverTab[104803]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:104
				// _ = "end of CoverTab[104803]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:104
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:104
			// _ = "end of CoverTab[104801]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:105
		// _ = "end of CoverTab[104796]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:106
	// _ = "end of CoverTab[104772]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:106
	_go_fuzz_dep_.CoverTab[104773]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:107
	// _ = "end of CoverTab[104773]"
}

func (r *OffsetCommitRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:110
	_go_fuzz_dep_.CoverTab[104804]++
													r.Version = version

													if r.ConsumerGroup, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:113
		_go_fuzz_dep_.CoverTab[104811]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:114
		// _ = "end of CoverTab[104811]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:115
		_go_fuzz_dep_.CoverTab[104812]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:115
		// _ = "end of CoverTab[104812]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:115
	// _ = "end of CoverTab[104804]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:115
	_go_fuzz_dep_.CoverTab[104805]++

													if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:117
		_go_fuzz_dep_.CoverTab[104813]++
														if r.ConsumerGroupGeneration, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:118
			_go_fuzz_dep_.CoverTab[104815]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:119
			// _ = "end of CoverTab[104815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:120
			_go_fuzz_dep_.CoverTab[104816]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:120
			// _ = "end of CoverTab[104816]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:120
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:120
		// _ = "end of CoverTab[104813]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:120
		_go_fuzz_dep_.CoverTab[104814]++
														if r.ConsumerID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:121
			_go_fuzz_dep_.CoverTab[104817]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:122
			// _ = "end of CoverTab[104817]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:123
			_go_fuzz_dep_.CoverTab[104818]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:123
			// _ = "end of CoverTab[104818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:123
		// _ = "end of CoverTab[104814]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:124
		_go_fuzz_dep_.CoverTab[104819]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:124
		// _ = "end of CoverTab[104819]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:124
	// _ = "end of CoverTab[104805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:124
	_go_fuzz_dep_.CoverTab[104806]++

													if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:126
		_go_fuzz_dep_.CoverTab[104820]++
														if r.RetentionTime, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:127
			_go_fuzz_dep_.CoverTab[104821]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:128
			// _ = "end of CoverTab[104821]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:129
			_go_fuzz_dep_.CoverTab[104822]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:129
			// _ = "end of CoverTab[104822]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:129
		// _ = "end of CoverTab[104820]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:130
		_go_fuzz_dep_.CoverTab[104823]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:130
		// _ = "end of CoverTab[104823]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:130
	// _ = "end of CoverTab[104806]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:130
	_go_fuzz_dep_.CoverTab[104807]++

													topicCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:133
		_go_fuzz_dep_.CoverTab[104824]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:134
		// _ = "end of CoverTab[104824]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:135
		_go_fuzz_dep_.CoverTab[104825]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:135
		// _ = "end of CoverTab[104825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:135
	// _ = "end of CoverTab[104807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:135
	_go_fuzz_dep_.CoverTab[104808]++
													if topicCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:136
		_go_fuzz_dep_.CoverTab[104826]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:137
		// _ = "end of CoverTab[104826]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:138
		_go_fuzz_dep_.CoverTab[104827]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:138
		// _ = "end of CoverTab[104827]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:138
	// _ = "end of CoverTab[104808]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:138
	_go_fuzz_dep_.CoverTab[104809]++
													r.blocks = make(map[string]map[int32]*offsetCommitRequestBlock)
													for i := 0; i < topicCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:140
		_go_fuzz_dep_.CoverTab[104828]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:142
			_go_fuzz_dep_.CoverTab[104831]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:143
			// _ = "end of CoverTab[104831]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:144
			_go_fuzz_dep_.CoverTab[104832]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:144
			// _ = "end of CoverTab[104832]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:144
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:144
		// _ = "end of CoverTab[104828]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:144
		_go_fuzz_dep_.CoverTab[104829]++
														partitionCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:146
			_go_fuzz_dep_.CoverTab[104833]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:147
			// _ = "end of CoverTab[104833]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:148
			_go_fuzz_dep_.CoverTab[104834]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:148
			// _ = "end of CoverTab[104834]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:148
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:148
		// _ = "end of CoverTab[104829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:148
		_go_fuzz_dep_.CoverTab[104830]++
														r.blocks[topic] = make(map[int32]*offsetCommitRequestBlock)
														for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:150
			_go_fuzz_dep_.CoverTab[104835]++
															partition, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:152
				_go_fuzz_dep_.CoverTab[104838]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:153
				// _ = "end of CoverTab[104838]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:154
				_go_fuzz_dep_.CoverTab[104839]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:154
				// _ = "end of CoverTab[104839]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:154
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:154
			// _ = "end of CoverTab[104835]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:154
			_go_fuzz_dep_.CoverTab[104836]++
															block := &offsetCommitRequestBlock{}
															if err := block.decode(pd, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:156
				_go_fuzz_dep_.CoverTab[104840]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:157
				// _ = "end of CoverTab[104840]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:158
				_go_fuzz_dep_.CoverTab[104841]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:158
				// _ = "end of CoverTab[104841]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:158
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:158
			// _ = "end of CoverTab[104836]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:158
			_go_fuzz_dep_.CoverTab[104837]++
															r.blocks[topic][partition] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:159
			// _ = "end of CoverTab[104837]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:160
		// _ = "end of CoverTab[104830]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:161
	// _ = "end of CoverTab[104809]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:161
	_go_fuzz_dep_.CoverTab[104810]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:162
	// _ = "end of CoverTab[104810]"
}

func (r *OffsetCommitRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:165
	_go_fuzz_dep_.CoverTab[104842]++
													return 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:166
	// _ = "end of CoverTab[104842]"
}

func (r *OffsetCommitRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:169
	_go_fuzz_dep_.CoverTab[104843]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:170
	// _ = "end of CoverTab[104843]"
}

func (r *OffsetCommitRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:173
	_go_fuzz_dep_.CoverTab[104844]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:174
	// _ = "end of CoverTab[104844]"
}

func (r *OffsetCommitRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:177
	_go_fuzz_dep_.CoverTab[104845]++
													switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:179
		_go_fuzz_dep_.CoverTab[104846]++
														return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:180
		// _ = "end of CoverTab[104846]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:181
		_go_fuzz_dep_.CoverTab[104847]++
														return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:182
		// _ = "end of CoverTab[104847]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:183
		_go_fuzz_dep_.CoverTab[104848]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:184
		// _ = "end of CoverTab[104848]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:185
		_go_fuzz_dep_.CoverTab[104849]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:186
		// _ = "end of CoverTab[104849]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:187
		_go_fuzz_dep_.CoverTab[104850]++
														return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:188
		// _ = "end of CoverTab[104850]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:189
	// _ = "end of CoverTab[104845]"
}

func (r *OffsetCommitRequest) AddBlock(topic string, partitionID int32, offset int64, timestamp int64, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:192
	_go_fuzz_dep_.CoverTab[104851]++
													if r.blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:193
		_go_fuzz_dep_.CoverTab[104854]++
														r.blocks = make(map[string]map[int32]*offsetCommitRequestBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:194
		// _ = "end of CoverTab[104854]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:195
		_go_fuzz_dep_.CoverTab[104855]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:195
		// _ = "end of CoverTab[104855]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:195
	// _ = "end of CoverTab[104851]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:195
	_go_fuzz_dep_.CoverTab[104852]++

													if r.blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:197
		_go_fuzz_dep_.CoverTab[104856]++
														r.blocks[topic] = make(map[int32]*offsetCommitRequestBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:198
		// _ = "end of CoverTab[104856]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:199
		_go_fuzz_dep_.CoverTab[104857]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:199
		// _ = "end of CoverTab[104857]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:199
	// _ = "end of CoverTab[104852]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:199
	_go_fuzz_dep_.CoverTab[104853]++

													r.blocks[topic][partitionID] = &offsetCommitRequestBlock{offset, timestamp, metadata}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:201
	// _ = "end of CoverTab[104853]"
}

func (r *OffsetCommitRequest) Offset(topic string, partitionID int32) (int64, string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:204
	_go_fuzz_dep_.CoverTab[104858]++
													partitions := r.blocks[topic]
													if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:206
		_go_fuzz_dep_.CoverTab[104861]++
														return 0, "", errors.New("no such offset")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:207
		// _ = "end of CoverTab[104861]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:208
		_go_fuzz_dep_.CoverTab[104862]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:208
		// _ = "end of CoverTab[104862]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:208
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:208
	// _ = "end of CoverTab[104858]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:208
	_go_fuzz_dep_.CoverTab[104859]++
													block := partitions[partitionID]
													if block == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:210
		_go_fuzz_dep_.CoverTab[104863]++
														return 0, "", errors.New("no such offset")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:211
		// _ = "end of CoverTab[104863]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:212
		_go_fuzz_dep_.CoverTab[104864]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:212
		// _ = "end of CoverTab[104864]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:212
	// _ = "end of CoverTab[104859]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:212
	_go_fuzz_dep_.CoverTab[104860]++
													return block.offset, block.metadata, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:213
	// _ = "end of CoverTab[104860]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:214
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_request.go:214
var _ = _go_fuzz_dep_.CoverTab
