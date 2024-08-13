//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:1
)

import (
	"time"
)

// AddPartitionsToTxnResponse is a partition errors to transaction type
type AddPartitionsToTxnResponse struct {
	ThrottleTime	time.Duration
	Errors		map[string][]*PartitionError
}

func (a *AddPartitionsToTxnResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:13
	_go_fuzz_dep_.CoverTab[97484]++
														pe.putInt32(int32(a.ThrottleTime / time.Millisecond))
														if err := pe.putArrayLength(len(a.Errors)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:15
		_go_fuzz_dep_.CoverTab[97487]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:16
		// _ = "end of CoverTab[97487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:17
		_go_fuzz_dep_.CoverTab[97488]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:17
		// _ = "end of CoverTab[97488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:17
	// _ = "end of CoverTab[97484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:17
	_go_fuzz_dep_.CoverTab[97485]++

														for topic, e := range a.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:19
		_go_fuzz_dep_.CoverTab[97489]++
															if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:20
			_go_fuzz_dep_.CoverTab[97492]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:21
			// _ = "end of CoverTab[97492]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:22
			_go_fuzz_dep_.CoverTab[97493]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:22
			// _ = "end of CoverTab[97493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:22
		// _ = "end of CoverTab[97489]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:22
		_go_fuzz_dep_.CoverTab[97490]++
															if err := pe.putArrayLength(len(e)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:23
			_go_fuzz_dep_.CoverTab[97494]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:24
			// _ = "end of CoverTab[97494]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:25
			_go_fuzz_dep_.CoverTab[97495]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:25
			// _ = "end of CoverTab[97495]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:25
		// _ = "end of CoverTab[97490]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:25
		_go_fuzz_dep_.CoverTab[97491]++
															for _, partitionError := range e {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:26
			_go_fuzz_dep_.CoverTab[97496]++
																if err := partitionError.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:27
				_go_fuzz_dep_.CoverTab[97497]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:28
				// _ = "end of CoverTab[97497]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:29
				_go_fuzz_dep_.CoverTab[97498]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:29
				// _ = "end of CoverTab[97498]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:29
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:29
			// _ = "end of CoverTab[97496]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:30
		// _ = "end of CoverTab[97491]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:31
	// _ = "end of CoverTab[97485]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:31
	_go_fuzz_dep_.CoverTab[97486]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:33
	// _ = "end of CoverTab[97486]"
}

func (a *AddPartitionsToTxnResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:36
	_go_fuzz_dep_.CoverTab[97499]++
														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:38
		_go_fuzz_dep_.CoverTab[97503]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:39
		// _ = "end of CoverTab[97503]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:40
		_go_fuzz_dep_.CoverTab[97504]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:40
		// _ = "end of CoverTab[97504]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:40
	// _ = "end of CoverTab[97499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:40
	_go_fuzz_dep_.CoverTab[97500]++
														a.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

														n, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:44
		_go_fuzz_dep_.CoverTab[97505]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:45
		// _ = "end of CoverTab[97505]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:46
		_go_fuzz_dep_.CoverTab[97506]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:46
		// _ = "end of CoverTab[97506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:46
	// _ = "end of CoverTab[97500]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:46
	_go_fuzz_dep_.CoverTab[97501]++

														a.Errors = make(map[string][]*PartitionError)

														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:50
		_go_fuzz_dep_.CoverTab[97507]++
															topic, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:52
			_go_fuzz_dep_.CoverTab[97510]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:53
			// _ = "end of CoverTab[97510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:54
			_go_fuzz_dep_.CoverTab[97511]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:54
			// _ = "end of CoverTab[97511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:54
		// _ = "end of CoverTab[97507]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:54
		_go_fuzz_dep_.CoverTab[97508]++

															m, err := pd.getArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:57
			_go_fuzz_dep_.CoverTab[97512]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:58
			// _ = "end of CoverTab[97512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:59
			_go_fuzz_dep_.CoverTab[97513]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:59
			// _ = "end of CoverTab[97513]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:59
		// _ = "end of CoverTab[97508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:59
		_go_fuzz_dep_.CoverTab[97509]++

															a.Errors[topic] = make([]*PartitionError, m)

															for j := 0; j < m; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:63
			_go_fuzz_dep_.CoverTab[97514]++
																a.Errors[topic][j] = new(PartitionError)
																if err := a.Errors[topic][j].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:65
				_go_fuzz_dep_.CoverTab[97515]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:66
				// _ = "end of CoverTab[97515]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:67
				_go_fuzz_dep_.CoverTab[97516]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:67
				// _ = "end of CoverTab[97516]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:67
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:67
			// _ = "end of CoverTab[97514]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:68
		// _ = "end of CoverTab[97509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:69
	// _ = "end of CoverTab[97501]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:69
	_go_fuzz_dep_.CoverTab[97502]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:71
	// _ = "end of CoverTab[97502]"
}

func (a *AddPartitionsToTxnResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:74
	_go_fuzz_dep_.CoverTab[97517]++
														return 24
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:75
	// _ = "end of CoverTab[97517]"
}

func (a *AddPartitionsToTxnResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:78
	_go_fuzz_dep_.CoverTab[97518]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:79
	// _ = "end of CoverTab[97518]"
}

func (a *AddPartitionsToTxnResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:82
	_go_fuzz_dep_.CoverTab[97519]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:83
	// _ = "end of CoverTab[97519]"
}

func (a *AddPartitionsToTxnResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:86
	_go_fuzz_dep_.CoverTab[97520]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:87
	// _ = "end of CoverTab[97520]"
}

// PartitionError is a partition error type
type PartitionError struct {
	Partition	int32
	Err		KError
}

func (p *PartitionError) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:96
	_go_fuzz_dep_.CoverTab[97521]++
														pe.putInt32(p.Partition)
														pe.putInt16(int16(p.Err))
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:99
	// _ = "end of CoverTab[97521]"
}

func (p *PartitionError) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:102
	_go_fuzz_dep_.CoverTab[97522]++
														if p.Partition, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:103
		_go_fuzz_dep_.CoverTab[97525]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:104
		// _ = "end of CoverTab[97525]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:105
		_go_fuzz_dep_.CoverTab[97526]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:105
		// _ = "end of CoverTab[97526]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:105
	// _ = "end of CoverTab[97522]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:105
	_go_fuzz_dep_.CoverTab[97523]++

														kerr, err := pd.getInt16()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:108
		_go_fuzz_dep_.CoverTab[97527]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:109
		// _ = "end of CoverTab[97527]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:110
		_go_fuzz_dep_.CoverTab[97528]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:110
		// _ = "end of CoverTab[97528]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:110
	// _ = "end of CoverTab[97523]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:110
	_go_fuzz_dep_.CoverTab[97524]++
														p.Err = KError(kerr)

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:113
	// _ = "end of CoverTab[97524]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_partitions_to_txn_response.go:114
var _ = _go_fuzz_dep_.CoverTab
