//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:1
)

// AddPartitionsToTxnRequest is a add paartition request
type AddPartitionsToTxnRequest struct {
	TransactionalID	string
	ProducerID	int64
	ProducerEpoch	int16
	TopicPartitions	map[string][]int32
}

func (a *AddPartitionsToTxnRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:11
	_go_fuzz_dep_.CoverTab[97445]++
														if err := pe.putString(a.TransactionalID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:12
		_go_fuzz_dep_.CoverTab[97449]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:13
		// _ = "end of CoverTab[97449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:14
		_go_fuzz_dep_.CoverTab[97450]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:14
		// _ = "end of CoverTab[97450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:14
	// _ = "end of CoverTab[97445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:14
	_go_fuzz_dep_.CoverTab[97446]++
														pe.putInt64(a.ProducerID)
														pe.putInt16(a.ProducerEpoch)

														if err := pe.putArrayLength(len(a.TopicPartitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:18
		_go_fuzz_dep_.CoverTab[97451]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:19
		// _ = "end of CoverTab[97451]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:20
		_go_fuzz_dep_.CoverTab[97452]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:20
		// _ = "end of CoverTab[97452]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:20
	// _ = "end of CoverTab[97446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:20
	_go_fuzz_dep_.CoverTab[97447]++
														for topic, partitions := range a.TopicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:21
		_go_fuzz_dep_.CoverTab[97453]++
															if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:22
			_go_fuzz_dep_.CoverTab[97455]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:23
			// _ = "end of CoverTab[97455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:24
			_go_fuzz_dep_.CoverTab[97456]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:24
			// _ = "end of CoverTab[97456]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:24
		// _ = "end of CoverTab[97453]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:24
		_go_fuzz_dep_.CoverTab[97454]++
															if err := pe.putInt32Array(partitions); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:25
			_go_fuzz_dep_.CoverTab[97457]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:26
			// _ = "end of CoverTab[97457]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:27
			_go_fuzz_dep_.CoverTab[97458]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:27
			// _ = "end of CoverTab[97458]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:27
		// _ = "end of CoverTab[97454]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:28
	// _ = "end of CoverTab[97447]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:28
	_go_fuzz_dep_.CoverTab[97448]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:30
	// _ = "end of CoverTab[97448]"
}

func (a *AddPartitionsToTxnRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:33
	_go_fuzz_dep_.CoverTab[97459]++
														if a.TransactionalID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:34
		_go_fuzz_dep_.CoverTab[97465]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:35
		// _ = "end of CoverTab[97465]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:36
		_go_fuzz_dep_.CoverTab[97466]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:36
		// _ = "end of CoverTab[97466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:36
	// _ = "end of CoverTab[97459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:36
	_go_fuzz_dep_.CoverTab[97460]++
														if a.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:37
		_go_fuzz_dep_.CoverTab[97467]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:38
		// _ = "end of CoverTab[97467]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:39
		_go_fuzz_dep_.CoverTab[97468]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:39
		// _ = "end of CoverTab[97468]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:39
	// _ = "end of CoverTab[97460]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:39
	_go_fuzz_dep_.CoverTab[97461]++
														if a.ProducerEpoch, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:40
		_go_fuzz_dep_.CoverTab[97469]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:41
		// _ = "end of CoverTab[97469]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:42
		_go_fuzz_dep_.CoverTab[97470]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:42
		// _ = "end of CoverTab[97470]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:42
	// _ = "end of CoverTab[97461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:42
	_go_fuzz_dep_.CoverTab[97462]++

														n, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:45
		_go_fuzz_dep_.CoverTab[97471]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:46
		// _ = "end of CoverTab[97471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:47
		_go_fuzz_dep_.CoverTab[97472]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:47
		// _ = "end of CoverTab[97472]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:47
	// _ = "end of CoverTab[97462]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:47
	_go_fuzz_dep_.CoverTab[97463]++

														a.TopicPartitions = make(map[string][]int32)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:50
		_go_fuzz_dep_.CoverTab[97473]++
															topic, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:52
			_go_fuzz_dep_.CoverTab[97476]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:53
			// _ = "end of CoverTab[97476]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:54
			_go_fuzz_dep_.CoverTab[97477]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:54
			// _ = "end of CoverTab[97477]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:54
		// _ = "end of CoverTab[97473]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:54
		_go_fuzz_dep_.CoverTab[97474]++

															partitions, err := pd.getInt32Array()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:57
			_go_fuzz_dep_.CoverTab[97478]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:58
			// _ = "end of CoverTab[97478]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:59
			_go_fuzz_dep_.CoverTab[97479]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:59
			// _ = "end of CoverTab[97479]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:59
		// _ = "end of CoverTab[97474]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:59
		_go_fuzz_dep_.CoverTab[97475]++

															a.TopicPartitions[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:61
		// _ = "end of CoverTab[97475]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:62
	// _ = "end of CoverTab[97463]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:62
	_go_fuzz_dep_.CoverTab[97464]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:64
	// _ = "end of CoverTab[97464]"
}

func (a *AddPartitionsToTxnRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:67
	_go_fuzz_dep_.CoverTab[97480]++
														return 24
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:68
	// _ = "end of CoverTab[97480]"
}

func (a *AddPartitionsToTxnRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:71
	_go_fuzz_dep_.CoverTab[97481]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:72
	// _ = "end of CoverTab[97481]"
}

func (a *AddPartitionsToTxnRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:75
	_go_fuzz_dep_.CoverTab[97482]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:76
	// _ = "end of CoverTab[97482]"
}

func (a *AddPartitionsToTxnRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:79
	_go_fuzz_dep_.CoverTab[97483]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:80
	// _ = "end of CoverTab[97483]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_request.go:81
var _ = _go_fuzz_dep_.CoverTab
