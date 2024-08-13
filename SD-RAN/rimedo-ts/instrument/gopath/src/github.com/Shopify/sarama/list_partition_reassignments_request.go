//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:1
)

type ListPartitionReassignmentsRequest struct {
	TimeoutMs	int32
	blocks		map[string][]int32
	Version		int16
}

func (r *ListPartitionReassignmentsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:9
	_go_fuzz_dep_.CoverTab[103823]++
														pe.putInt32(r.TimeoutMs)

														pe.putCompactArrayLength(len(r.blocks))

														for topic, partitions := range r.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:14
		_go_fuzz_dep_.CoverTab[103825]++
																if err := pe.putCompactString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:15
			_go_fuzz_dep_.CoverTab[103828]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:16
			// _ = "end of CoverTab[103828]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:17
			_go_fuzz_dep_.CoverTab[103829]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:17
			// _ = "end of CoverTab[103829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:17
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:17
		// _ = "end of CoverTab[103825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:17
		_go_fuzz_dep_.CoverTab[103826]++

																if err := pe.putCompactInt32Array(partitions); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:19
			_go_fuzz_dep_.CoverTab[103830]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:20
			// _ = "end of CoverTab[103830]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:21
			_go_fuzz_dep_.CoverTab[103831]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:21
			// _ = "end of CoverTab[103831]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:21
		// _ = "end of CoverTab[103826]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:21
		_go_fuzz_dep_.CoverTab[103827]++

																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:23
		// _ = "end of CoverTab[103827]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:24
	// _ = "end of CoverTab[103823]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:24
	_go_fuzz_dep_.CoverTab[103824]++

															pe.putEmptyTaggedFieldArray()

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:28
	// _ = "end of CoverTab[103824]"
}

func (r *ListPartitionReassignmentsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:31
	_go_fuzz_dep_.CoverTab[103832]++
															r.Version = version

															if r.TimeoutMs, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:34
		_go_fuzz_dep_.CoverTab[103837]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:35
		// _ = "end of CoverTab[103837]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:36
		_go_fuzz_dep_.CoverTab[103838]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:36
		// _ = "end of CoverTab[103838]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:36
	// _ = "end of CoverTab[103832]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:36
	_go_fuzz_dep_.CoverTab[103833]++

															topicCount, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:39
		_go_fuzz_dep_.CoverTab[103839]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:40
		// _ = "end of CoverTab[103839]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:41
		_go_fuzz_dep_.CoverTab[103840]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:41
		// _ = "end of CoverTab[103840]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:41
	// _ = "end of CoverTab[103833]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:41
	_go_fuzz_dep_.CoverTab[103834]++
															if topicCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:42
		_go_fuzz_dep_.CoverTab[103841]++
																r.blocks = make(map[string][]int32)
																for i := 0; i < topicCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:44
			_go_fuzz_dep_.CoverTab[103842]++
																	topic, err := pd.getCompactString()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:46
				_go_fuzz_dep_.CoverTab[103846]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:47
				// _ = "end of CoverTab[103846]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:48
				_go_fuzz_dep_.CoverTab[103847]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:48
				// _ = "end of CoverTab[103847]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:48
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:48
			// _ = "end of CoverTab[103842]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:48
			_go_fuzz_dep_.CoverTab[103843]++
																	partitionCount, err := pd.getCompactArrayLength()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:50
				_go_fuzz_dep_.CoverTab[103848]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:51
				// _ = "end of CoverTab[103848]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:52
				_go_fuzz_dep_.CoverTab[103849]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:52
				// _ = "end of CoverTab[103849]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:52
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:52
			// _ = "end of CoverTab[103843]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:52
			_go_fuzz_dep_.CoverTab[103844]++
																	r.blocks[topic] = make([]int32, partitionCount)
																	for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:54
				_go_fuzz_dep_.CoverTab[103850]++
																		partition, err := pd.getInt32()
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:56
					_go_fuzz_dep_.CoverTab[103852]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:57
					// _ = "end of CoverTab[103852]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:58
					_go_fuzz_dep_.CoverTab[103853]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:58
					// _ = "end of CoverTab[103853]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:58
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:58
				// _ = "end of CoverTab[103850]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:58
				_go_fuzz_dep_.CoverTab[103851]++
																		r.blocks[topic][j] = partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:59
				// _ = "end of CoverTab[103851]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:60
			// _ = "end of CoverTab[103844]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:60
			_go_fuzz_dep_.CoverTab[103845]++
																	if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:61
				_go_fuzz_dep_.CoverTab[103854]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:62
				// _ = "end of CoverTab[103854]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:63
				_go_fuzz_dep_.CoverTab[103855]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:63
				// _ = "end of CoverTab[103855]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:63
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:63
			// _ = "end of CoverTab[103845]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:64
		// _ = "end of CoverTab[103841]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:65
		_go_fuzz_dep_.CoverTab[103856]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:65
		// _ = "end of CoverTab[103856]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:65
	// _ = "end of CoverTab[103834]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:65
	_go_fuzz_dep_.CoverTab[103835]++

															if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:67
		_go_fuzz_dep_.CoverTab[103857]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:68
		// _ = "end of CoverTab[103857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:69
		_go_fuzz_dep_.CoverTab[103858]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:69
		// _ = "end of CoverTab[103858]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:69
	// _ = "end of CoverTab[103835]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:69
	_go_fuzz_dep_.CoverTab[103836]++

															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:71
	// _ = "end of CoverTab[103836]"
}

func (r *ListPartitionReassignmentsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:74
	_go_fuzz_dep_.CoverTab[103859]++
															return 46
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:75
	// _ = "end of CoverTab[103859]"
}

func (r *ListPartitionReassignmentsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:78
	_go_fuzz_dep_.CoverTab[103860]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:79
	// _ = "end of CoverTab[103860]"
}

func (r *ListPartitionReassignmentsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:82
	_go_fuzz_dep_.CoverTab[103861]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:83
	// _ = "end of CoverTab[103861]"
}

func (r *ListPartitionReassignmentsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:86
	_go_fuzz_dep_.CoverTab[103862]++
															return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:87
	// _ = "end of CoverTab[103862]"
}

func (r *ListPartitionReassignmentsRequest) AddBlock(topic string, partitionIDs []int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:90
	_go_fuzz_dep_.CoverTab[103863]++
															if r.blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:91
		_go_fuzz_dep_.CoverTab[103865]++
																r.blocks = make(map[string][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:92
		// _ = "end of CoverTab[103865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:93
		_go_fuzz_dep_.CoverTab[103866]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:93
		// _ = "end of CoverTab[103866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:93
	// _ = "end of CoverTab[103863]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:93
	_go_fuzz_dep_.CoverTab[103864]++

															if r.blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:95
		_go_fuzz_dep_.CoverTab[103867]++
																r.blocks[topic] = partitionIDs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:96
		// _ = "end of CoverTab[103867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:97
		_go_fuzz_dep_.CoverTab[103868]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:97
		// _ = "end of CoverTab[103868]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:97
	// _ = "end of CoverTab[103864]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_request.go:98
var _ = _go_fuzz_dep_.CoverTab
