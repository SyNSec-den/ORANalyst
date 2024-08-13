//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:1
)

import (
	"encoding/binary"
	"errors"
	"time"
)

type partitionSet struct {
	msgs		[]*ProducerMessage
	recordsToSend	Records
	bufferBytes	int
}

type produceSet struct {
	parent		*asyncProducer
	msgs		map[string]map[int32]*partitionSet
	producerID	int64
	producerEpoch	int16

	bufferBytes	int
	bufferCount	int
}

func newProduceSet(parent *asyncProducer) *produceSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:25
	_go_fuzz_dep_.CoverTab[105870]++
											pid, epoch := parent.txnmgr.getProducerID()
											return &produceSet{
		msgs:		make(map[string]map[int32]*partitionSet),
		parent:		parent,
		producerID:	pid,
		producerEpoch:	epoch,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:32
	// _ = "end of CoverTab[105870]"
}

func (ps *produceSet) add(msg *ProducerMessage) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:35
	_go_fuzz_dep_.CoverTab[105871]++
											var err error
											var key, val []byte

											if msg.Key != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:39
		_go_fuzz_dep_.CoverTab[105879]++
												if key, err = msg.Key.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:40
			_go_fuzz_dep_.CoverTab[105880]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:41
			// _ = "end of CoverTab[105880]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:42
			_go_fuzz_dep_.CoverTab[105881]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:42
			// _ = "end of CoverTab[105881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:42
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:42
		// _ = "end of CoverTab[105879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:43
		_go_fuzz_dep_.CoverTab[105882]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:43
		// _ = "end of CoverTab[105882]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:43
	// _ = "end of CoverTab[105871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:43
	_go_fuzz_dep_.CoverTab[105872]++

											if msg.Value != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:45
		_go_fuzz_dep_.CoverTab[105883]++
												if val, err = msg.Value.Encode(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:46
			_go_fuzz_dep_.CoverTab[105884]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:47
			// _ = "end of CoverTab[105884]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:48
			_go_fuzz_dep_.CoverTab[105885]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:48
			// _ = "end of CoverTab[105885]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:48
		// _ = "end of CoverTab[105883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:49
		_go_fuzz_dep_.CoverTab[105886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:49
		// _ = "end of CoverTab[105886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:49
	// _ = "end of CoverTab[105872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:49
	_go_fuzz_dep_.CoverTab[105873]++

											timestamp := msg.Timestamp
											if timestamp.IsZero() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:52
		_go_fuzz_dep_.CoverTab[105887]++
												timestamp = time.Now()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:53
		// _ = "end of CoverTab[105887]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:54
		_go_fuzz_dep_.CoverTab[105888]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:54
		// _ = "end of CoverTab[105888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:54
	// _ = "end of CoverTab[105873]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:54
	_go_fuzz_dep_.CoverTab[105874]++
											timestamp = timestamp.Truncate(time.Millisecond)

											partitions := ps.msgs[msg.Topic]
											if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:58
		_go_fuzz_dep_.CoverTab[105889]++
												partitions = make(map[int32]*partitionSet)
												ps.msgs[msg.Topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:60
		// _ = "end of CoverTab[105889]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:61
		_go_fuzz_dep_.CoverTab[105890]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:61
		// _ = "end of CoverTab[105890]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:61
	// _ = "end of CoverTab[105874]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:61
	_go_fuzz_dep_.CoverTab[105875]++

											var size int

											set := partitions[msg.Partition]
											if set == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:66
		_go_fuzz_dep_.CoverTab[105891]++
												if ps.parent.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:67
			_go_fuzz_dep_.CoverTab[105893]++
													batch := &RecordBatch{
				FirstTimestamp:		timestamp,
				Version:		2,
				Codec:			ps.parent.conf.Producer.Compression,
				CompressionLevel:	ps.parent.conf.Producer.CompressionLevel,
				ProducerID:		ps.producerID,
				ProducerEpoch:		ps.producerEpoch,
			}
			if ps.parent.conf.Producer.Idempotent {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:76
				_go_fuzz_dep_.CoverTab[105895]++
														batch.FirstSequence = msg.sequenceNumber
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:77
				// _ = "end of CoverTab[105895]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:78
				_go_fuzz_dep_.CoverTab[105896]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:78
				// _ = "end of CoverTab[105896]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:78
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:78
			// _ = "end of CoverTab[105893]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:78
			_go_fuzz_dep_.CoverTab[105894]++
													set = &partitionSet{recordsToSend: newDefaultRecords(batch)}
													size = recordBatchOverhead
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:80
			// _ = "end of CoverTab[105894]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:81
			_go_fuzz_dep_.CoverTab[105897]++
													set = &partitionSet{recordsToSend: newLegacyRecords(new(MessageSet))}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:82
			// _ = "end of CoverTab[105897]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:83
		// _ = "end of CoverTab[105891]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:83
		_go_fuzz_dep_.CoverTab[105892]++
												partitions[msg.Partition] = set
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:84
		// _ = "end of CoverTab[105892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:85
		_go_fuzz_dep_.CoverTab[105898]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:85
		// _ = "end of CoverTab[105898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:85
	// _ = "end of CoverTab[105875]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:85
	_go_fuzz_dep_.CoverTab[105876]++

											if ps.parent.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:87
		_go_fuzz_dep_.CoverTab[105899]++
												if ps.parent.conf.Producer.Idempotent && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:88
			_go_fuzz_dep_.CoverTab[105900]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:88
			return msg.sequenceNumber < set.recordsToSend.RecordBatch.FirstSequence
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:88
			// _ = "end of CoverTab[105900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:88
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:88
			_go_fuzz_dep_.CoverTab[105901]++
													return errors.New("assertion failed: message out of sequence added to a batch")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:89
			// _ = "end of CoverTab[105901]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:90
			_go_fuzz_dep_.CoverTab[105902]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:90
			// _ = "end of CoverTab[105902]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:90
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:90
		// _ = "end of CoverTab[105899]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:91
		_go_fuzz_dep_.CoverTab[105903]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:91
		// _ = "end of CoverTab[105903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:91
	// _ = "end of CoverTab[105876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:91
	_go_fuzz_dep_.CoverTab[105877]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:94
	set.msgs = append(set.msgs, msg)

	if ps.parent.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:96
		_go_fuzz_dep_.CoverTab[105904]++

												size += maximumRecordOverhead
												rec := &Record{
			Key:		key,
			Value:		val,
			TimestampDelta:	timestamp.Sub(set.recordsToSend.RecordBatch.FirstTimestamp),
		}
		size += len(key) + len(val)
		if len(msg.Headers) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:105
			_go_fuzz_dep_.CoverTab[105906]++
														rec.Headers = make([]*RecordHeader, len(msg.Headers))
														for i := range msg.Headers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:107
				_go_fuzz_dep_.CoverTab[105907]++
															rec.Headers[i] = &msg.Headers[i]
															size += len(rec.Headers[i].Key) + len(rec.Headers[i].Value) + 2*binary.MaxVarintLen32
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:109
				// _ = "end of CoverTab[105907]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:110
			// _ = "end of CoverTab[105906]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:111
			_go_fuzz_dep_.CoverTab[105908]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:111
			// _ = "end of CoverTab[105908]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:111
		// _ = "end of CoverTab[105904]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:111
		_go_fuzz_dep_.CoverTab[105905]++
													set.recordsToSend.RecordBatch.addRecord(rec)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:112
		// _ = "end of CoverTab[105905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:113
		_go_fuzz_dep_.CoverTab[105909]++
													msgToSend := &Message{Codec: CompressionNone, Key: key, Value: val}
													if ps.parent.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:115
			_go_fuzz_dep_.CoverTab[105911]++
														msgToSend.Timestamp = timestamp
														msgToSend.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:117
			// _ = "end of CoverTab[105911]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:118
			_go_fuzz_dep_.CoverTab[105912]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:118
			// _ = "end of CoverTab[105912]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:118
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:118
		// _ = "end of CoverTab[105909]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:118
		_go_fuzz_dep_.CoverTab[105910]++
													set.recordsToSend.MsgSet.addMessage(msgToSend)
													size = producerMessageOverhead + len(key) + len(val)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:120
		// _ = "end of CoverTab[105910]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:121
	// _ = "end of CoverTab[105877]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:121
	_go_fuzz_dep_.CoverTab[105878]++

												set.bufferBytes += size
												ps.bufferBytes += size
												ps.bufferCount++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:127
	// _ = "end of CoverTab[105878]"
}

func (ps *produceSet) buildRequest() *ProduceRequest {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:130
	_go_fuzz_dep_.CoverTab[105913]++
												req := &ProduceRequest{
		RequiredAcks:	ps.parent.conf.Producer.RequiredAcks,
		Timeout:	int32(ps.parent.conf.Producer.Timeout / time.Millisecond),
	}
	if ps.parent.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:135
		_go_fuzz_dep_.CoverTab[105918]++
													req.Version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:136
		// _ = "end of CoverTab[105918]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:137
		_go_fuzz_dep_.CoverTab[105919]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:137
		// _ = "end of CoverTab[105919]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:137
	// _ = "end of CoverTab[105913]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:137
	_go_fuzz_dep_.CoverTab[105914]++
												if ps.parent.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:138
		_go_fuzz_dep_.CoverTab[105920]++
													req.Version = 3
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:139
		// _ = "end of CoverTab[105920]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:140
		_go_fuzz_dep_.CoverTab[105921]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:140
		// _ = "end of CoverTab[105921]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:140
	// _ = "end of CoverTab[105914]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:140
	_go_fuzz_dep_.CoverTab[105915]++

												if ps.parent.conf.Producer.Compression == CompressionZSTD && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:142
		_go_fuzz_dep_.CoverTab[105922]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:142
		return ps.parent.conf.Version.IsAtLeast(V2_1_0_0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:142
		// _ = "end of CoverTab[105922]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:142
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:142
		_go_fuzz_dep_.CoverTab[105923]++
													req.Version = 7
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:143
		// _ = "end of CoverTab[105923]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:144
		_go_fuzz_dep_.CoverTab[105924]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:144
		// _ = "end of CoverTab[105924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:144
	// _ = "end of CoverTab[105915]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:144
	_go_fuzz_dep_.CoverTab[105916]++

												for topic, partitionSets := range ps.msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:146
		_go_fuzz_dep_.CoverTab[105925]++
													for partition, set := range partitionSets {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:147
			_go_fuzz_dep_.CoverTab[105926]++
														if req.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:148
				_go_fuzz_dep_.CoverTab[105928]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:156
				rb := set.recordsToSend.RecordBatch
				if len(rb.Records) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:157
					_go_fuzz_dep_.CoverTab[105930]++
																rb.LastOffsetDelta = int32(len(rb.Records) - 1)
																for i, record := range rb.Records {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:159
						_go_fuzz_dep_.CoverTab[105931]++
																	record.OffsetDelta = int64(i)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:160
						// _ = "end of CoverTab[105931]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:161
					// _ = "end of CoverTab[105930]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:162
					_go_fuzz_dep_.CoverTab[105932]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:162
					// _ = "end of CoverTab[105932]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:162
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:162
				// _ = "end of CoverTab[105928]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:162
				_go_fuzz_dep_.CoverTab[105929]++
															req.AddBatch(topic, partition, rb)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:164
				// _ = "end of CoverTab[105929]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:165
				_go_fuzz_dep_.CoverTab[105933]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:165
				// _ = "end of CoverTab[105933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:165
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:165
			// _ = "end of CoverTab[105926]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:165
			_go_fuzz_dep_.CoverTab[105927]++
														if ps.parent.conf.Producer.Compression == CompressionNone {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:166
				_go_fuzz_dep_.CoverTab[105934]++
															req.AddSet(topic, partition, set.recordsToSend.MsgSet)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:167
				// _ = "end of CoverTab[105934]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:168
				_go_fuzz_dep_.CoverTab[105935]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:174
				if ps.parent.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:174
					_go_fuzz_dep_.CoverTab[105939]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:180
					for i, msg := range set.recordsToSend.MsgSet.Messages {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:180
						_go_fuzz_dep_.CoverTab[105940]++
																	msg.Offset = int64(i)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:181
						// _ = "end of CoverTab[105940]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:182
					// _ = "end of CoverTab[105939]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:183
					_go_fuzz_dep_.CoverTab[105941]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:183
					// _ = "end of CoverTab[105941]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:183
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:183
				// _ = "end of CoverTab[105935]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:183
				_go_fuzz_dep_.CoverTab[105936]++
															payload, err := encode(set.recordsToSend.MsgSet, ps.parent.conf.MetricRegistry)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:185
					_go_fuzz_dep_.CoverTab[105942]++
																Logger.Println(err)
																panic(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:187
					// _ = "end of CoverTab[105942]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:188
					_go_fuzz_dep_.CoverTab[105943]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:188
					// _ = "end of CoverTab[105943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:188
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:188
				// _ = "end of CoverTab[105936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:188
				_go_fuzz_dep_.CoverTab[105937]++
															compMsg := &Message{
					Codec:			ps.parent.conf.Producer.Compression,
					CompressionLevel:	ps.parent.conf.Producer.CompressionLevel,
					Key:			nil,
					Value:			payload,
					Set:			set.recordsToSend.MsgSet,
				}
				if ps.parent.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:196
					_go_fuzz_dep_.CoverTab[105944]++
																compMsg.Version = 1
																compMsg.Timestamp = set.recordsToSend.MsgSet.Messages[0].Msg.Timestamp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:198
					// _ = "end of CoverTab[105944]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:199
					_go_fuzz_dep_.CoverTab[105945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:199
					// _ = "end of CoverTab[105945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:199
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:199
				// _ = "end of CoverTab[105937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:199
				_go_fuzz_dep_.CoverTab[105938]++
															req.AddMessage(topic, partition, compMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:200
				// _ = "end of CoverTab[105938]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:201
			// _ = "end of CoverTab[105927]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:202
		// _ = "end of CoverTab[105925]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:203
	// _ = "end of CoverTab[105916]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:203
	_go_fuzz_dep_.CoverTab[105917]++

												return req
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:205
	// _ = "end of CoverTab[105917]"
}

func (ps *produceSet) eachPartition(cb func(topic string, partition int32, pSet *partitionSet)) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:208
	_go_fuzz_dep_.CoverTab[105946]++
												for topic, partitionSet := range ps.msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:209
		_go_fuzz_dep_.CoverTab[105947]++
													for partition, set := range partitionSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:210
			_go_fuzz_dep_.CoverTab[105948]++
														cb(topic, partition, set)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:211
			// _ = "end of CoverTab[105948]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:212
		// _ = "end of CoverTab[105947]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:213
	// _ = "end of CoverTab[105946]"
}

func (ps *produceSet) dropPartition(topic string, partition int32) []*ProducerMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:216
	_go_fuzz_dep_.CoverTab[105949]++
												if ps.msgs[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:217
		_go_fuzz_dep_.CoverTab[105952]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:218
		// _ = "end of CoverTab[105952]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:219
		_go_fuzz_dep_.CoverTab[105953]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:219
		// _ = "end of CoverTab[105953]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:219
	// _ = "end of CoverTab[105949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:219
	_go_fuzz_dep_.CoverTab[105950]++
												set := ps.msgs[topic][partition]
												if set == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:221
		_go_fuzz_dep_.CoverTab[105954]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:222
		// _ = "end of CoverTab[105954]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:223
		_go_fuzz_dep_.CoverTab[105955]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:223
		// _ = "end of CoverTab[105955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:223
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:223
	// _ = "end of CoverTab[105950]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:223
	_go_fuzz_dep_.CoverTab[105951]++
												ps.bufferBytes -= set.bufferBytes
												ps.bufferCount -= len(set.msgs)
												delete(ps.msgs[topic], partition)
												return set.msgs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:227
	// _ = "end of CoverTab[105951]"
}

func (ps *produceSet) wouldOverflow(msg *ProducerMessage) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:230
	_go_fuzz_dep_.CoverTab[105956]++
												version := 1
												if ps.parent.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:232
		_go_fuzz_dep_.CoverTab[105958]++
													version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:233
		// _ = "end of CoverTab[105958]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:234
		_go_fuzz_dep_.CoverTab[105959]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:234
		// _ = "end of CoverTab[105959]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:234
	// _ = "end of CoverTab[105956]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:234
	_go_fuzz_dep_.CoverTab[105957]++

												switch {

	case ps.bufferBytes+msg.byteSize(version) >= int(MaxRequestSize-(10*1024)):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:238
		_go_fuzz_dep_.CoverTab[105960]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:239
		// _ = "end of CoverTab[105960]"

	case ps.msgs[msg.Topic] != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:241
		_go_fuzz_dep_.CoverTab[105964]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:241
		return ps.msgs[msg.Topic][msg.Partition] != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:241
		// _ = "end of CoverTab[105964]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:241
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:241
		_go_fuzz_dep_.CoverTab[105965]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:241
		return ps.msgs[msg.Topic][msg.Partition].bufferBytes+msg.byteSize(version) >= ps.parent.conf.Producer.MaxMessageBytes
													// _ = "end of CoverTab[105965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:242
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:242
		_go_fuzz_dep_.CoverTab[105961]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:243
		// _ = "end of CoverTab[105961]"

	case ps.parent.conf.Producer.Flush.MaxMessages > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:245
		_go_fuzz_dep_.CoverTab[105966]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:245
		return ps.bufferCount >= ps.parent.conf.Producer.Flush.MaxMessages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:245
		// _ = "end of CoverTab[105966]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:245
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:245
		_go_fuzz_dep_.CoverTab[105962]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:246
		// _ = "end of CoverTab[105962]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:247
		_go_fuzz_dep_.CoverTab[105963]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:248
		// _ = "end of CoverTab[105963]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:249
	// _ = "end of CoverTab[105957]"
}

func (ps *produceSet) readyToFlush() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:252
	_go_fuzz_dep_.CoverTab[105967]++
												switch {

	case ps.empty():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:255
		_go_fuzz_dep_.CoverTab[105968]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:256
		// _ = "end of CoverTab[105968]"

	case ps.parent.conf.Producer.Flush.Frequency == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		_go_fuzz_dep_.CoverTab[105973]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		return ps.parent.conf.Producer.Flush.Bytes == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		// _ = "end of CoverTab[105973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		_go_fuzz_dep_.CoverTab[105974]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		return ps.parent.conf.Producer.Flush.Messages == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		// _ = "end of CoverTab[105974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:258
		_go_fuzz_dep_.CoverTab[105969]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:259
		// _ = "end of CoverTab[105969]"

	case ps.parent.conf.Producer.Flush.Messages > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:261
		_go_fuzz_dep_.CoverTab[105975]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:261
		return ps.bufferCount >= ps.parent.conf.Producer.Flush.Messages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:261
		// _ = "end of CoverTab[105975]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:261
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:261
		_go_fuzz_dep_.CoverTab[105970]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:262
		// _ = "end of CoverTab[105970]"

	case ps.parent.conf.Producer.Flush.Bytes > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:264
		_go_fuzz_dep_.CoverTab[105976]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:264
		return ps.bufferBytes >= ps.parent.conf.Producer.Flush.Bytes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:264
		// _ = "end of CoverTab[105976]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:264
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:264
		_go_fuzz_dep_.CoverTab[105971]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:265
		// _ = "end of CoverTab[105971]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:266
		_go_fuzz_dep_.CoverTab[105972]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:267
		// _ = "end of CoverTab[105972]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:268
	// _ = "end of CoverTab[105967]"
}

func (ps *produceSet) empty() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:271
	_go_fuzz_dep_.CoverTab[105977]++
												return ps.bufferCount == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:272
	// _ = "end of CoverTab[105977]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:273
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_set.go:273
var _ = _go_fuzz_dep_.CoverTab
