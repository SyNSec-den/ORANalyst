//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:1
)

type alterPartitionReassignmentsBlock struct {
	replicas []int32
}

func (b *alterPartitionReassignmentsBlock) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:7
	_go_fuzz_dep_.CoverTab[98207]++
															if err := pe.putNullableCompactInt32Array(b.replicas); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:8
		_go_fuzz_dep_.CoverTab[98209]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:9
		// _ = "end of CoverTab[98209]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:10
		_go_fuzz_dep_.CoverTab[98210]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:10
		// _ = "end of CoverTab[98210]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:10
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:10
	// _ = "end of CoverTab[98207]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:10
	_go_fuzz_dep_.CoverTab[98208]++

															pe.putEmptyTaggedFieldArray()
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:13
	// _ = "end of CoverTab[98208]"
}

func (b *alterPartitionReassignmentsBlock) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:16
	_go_fuzz_dep_.CoverTab[98211]++
															if b.replicas, err = pd.getCompactInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:17
		_go_fuzz_dep_.CoverTab[98213]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:18
		// _ = "end of CoverTab[98213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:19
		_go_fuzz_dep_.CoverTab[98214]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:19
		// _ = "end of CoverTab[98214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:19
	// _ = "end of CoverTab[98211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:19
	_go_fuzz_dep_.CoverTab[98212]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:20
	// _ = "end of CoverTab[98212]"
}

type AlterPartitionReassignmentsRequest struct {
	TimeoutMs	int32
	blocks		map[string]map[int32]*alterPartitionReassignmentsBlock
	Version		int16
}

func (r *AlterPartitionReassignmentsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:29
	_go_fuzz_dep_.CoverTab[98215]++
															pe.putInt32(r.TimeoutMs)

															pe.putCompactArrayLength(len(r.blocks))

															for topic, partitions := range r.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:34
		_go_fuzz_dep_.CoverTab[98217]++
																if err := pe.putCompactString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:35
			_go_fuzz_dep_.CoverTab[98220]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:36
			// _ = "end of CoverTab[98220]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:37
			_go_fuzz_dep_.CoverTab[98221]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:37
			// _ = "end of CoverTab[98221]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:37
		// _ = "end of CoverTab[98217]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:37
		_go_fuzz_dep_.CoverTab[98218]++
																pe.putCompactArrayLength(len(partitions))
																for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:39
			_go_fuzz_dep_.CoverTab[98222]++
																	pe.putInt32(partition)
																	if err := block.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:41
				_go_fuzz_dep_.CoverTab[98223]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:42
				// _ = "end of CoverTab[98223]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:43
				_go_fuzz_dep_.CoverTab[98224]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:43
				// _ = "end of CoverTab[98224]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:43
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:43
			// _ = "end of CoverTab[98222]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:44
		// _ = "end of CoverTab[98218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:44
		_go_fuzz_dep_.CoverTab[98219]++
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:45
		// _ = "end of CoverTab[98219]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:46
	// _ = "end of CoverTab[98215]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:46
	_go_fuzz_dep_.CoverTab[98216]++

															pe.putEmptyTaggedFieldArray()

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:50
	// _ = "end of CoverTab[98216]"
}

func (r *AlterPartitionReassignmentsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:53
	_go_fuzz_dep_.CoverTab[98225]++
															r.Version = version

															if r.TimeoutMs, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:56
		_go_fuzz_dep_.CoverTab[98230]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:57
		// _ = "end of CoverTab[98230]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:58
		_go_fuzz_dep_.CoverTab[98231]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:58
		// _ = "end of CoverTab[98231]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:58
	// _ = "end of CoverTab[98225]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:58
	_go_fuzz_dep_.CoverTab[98226]++

															topicCount, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:61
		_go_fuzz_dep_.CoverTab[98232]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:62
		// _ = "end of CoverTab[98232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:63
		_go_fuzz_dep_.CoverTab[98233]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:63
		// _ = "end of CoverTab[98233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:63
	// _ = "end of CoverTab[98226]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:63
	_go_fuzz_dep_.CoverTab[98227]++
															if topicCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:64
		_go_fuzz_dep_.CoverTab[98234]++
																r.blocks = make(map[string]map[int32]*alterPartitionReassignmentsBlock)
																for i := 0; i < topicCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:66
			_go_fuzz_dep_.CoverTab[98235]++
																	topic, err := pd.getCompactString()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:68
				_go_fuzz_dep_.CoverTab[98239]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:69
				// _ = "end of CoverTab[98239]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:70
				_go_fuzz_dep_.CoverTab[98240]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:70
				// _ = "end of CoverTab[98240]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:70
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:70
			// _ = "end of CoverTab[98235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:70
			_go_fuzz_dep_.CoverTab[98236]++
																	partitionCount, err := pd.getCompactArrayLength()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:72
				_go_fuzz_dep_.CoverTab[98241]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:73
				// _ = "end of CoverTab[98241]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:74
				_go_fuzz_dep_.CoverTab[98242]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:74
				// _ = "end of CoverTab[98242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:74
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:74
			// _ = "end of CoverTab[98236]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:74
			_go_fuzz_dep_.CoverTab[98237]++
																	r.blocks[topic] = make(map[int32]*alterPartitionReassignmentsBlock)
																	for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:76
				_go_fuzz_dep_.CoverTab[98243]++
																		partition, err := pd.getInt32()
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:78
					_go_fuzz_dep_.CoverTab[98246]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:79
					// _ = "end of CoverTab[98246]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:80
					_go_fuzz_dep_.CoverTab[98247]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:80
					// _ = "end of CoverTab[98247]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:80
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:80
				// _ = "end of CoverTab[98243]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:80
				_go_fuzz_dep_.CoverTab[98244]++
																		block := &alterPartitionReassignmentsBlock{}
																		if err := block.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:82
					_go_fuzz_dep_.CoverTab[98248]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:83
					// _ = "end of CoverTab[98248]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:84
					_go_fuzz_dep_.CoverTab[98249]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:84
					// _ = "end of CoverTab[98249]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:84
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:84
				// _ = "end of CoverTab[98244]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:84
				_go_fuzz_dep_.CoverTab[98245]++
																		r.blocks[topic][partition] = block

																		if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:87
					_go_fuzz_dep_.CoverTab[98250]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:88
					// _ = "end of CoverTab[98250]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:89
					_go_fuzz_dep_.CoverTab[98251]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:89
					// _ = "end of CoverTab[98251]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:89
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:89
				// _ = "end of CoverTab[98245]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:90
			// _ = "end of CoverTab[98237]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:90
			_go_fuzz_dep_.CoverTab[98238]++
																	if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:91
				_go_fuzz_dep_.CoverTab[98252]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:92
				// _ = "end of CoverTab[98252]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:93
				_go_fuzz_dep_.CoverTab[98253]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:93
				// _ = "end of CoverTab[98253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:93
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:93
			// _ = "end of CoverTab[98238]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:94
		// _ = "end of CoverTab[98234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:95
		_go_fuzz_dep_.CoverTab[98254]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:95
		// _ = "end of CoverTab[98254]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:95
	// _ = "end of CoverTab[98227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:95
	_go_fuzz_dep_.CoverTab[98228]++

															if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:97
		_go_fuzz_dep_.CoverTab[98255]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:98
		// _ = "end of CoverTab[98255]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:99
		_go_fuzz_dep_.CoverTab[98256]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:99
		// _ = "end of CoverTab[98256]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:99
	// _ = "end of CoverTab[98228]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:99
	_go_fuzz_dep_.CoverTab[98229]++

															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:101
	// _ = "end of CoverTab[98229]"
}

func (r *AlterPartitionReassignmentsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:104
	_go_fuzz_dep_.CoverTab[98257]++
															return 45
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:105
	// _ = "end of CoverTab[98257]"
}

func (r *AlterPartitionReassignmentsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:108
	_go_fuzz_dep_.CoverTab[98258]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:109
	// _ = "end of CoverTab[98258]"
}

func (r *AlterPartitionReassignmentsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:112
	_go_fuzz_dep_.CoverTab[98259]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:113
	// _ = "end of CoverTab[98259]"
}

func (r *AlterPartitionReassignmentsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:116
	_go_fuzz_dep_.CoverTab[98260]++
															return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:117
	// _ = "end of CoverTab[98260]"
}

func (r *AlterPartitionReassignmentsRequest) AddBlock(topic string, partitionID int32, replicas []int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:120
	_go_fuzz_dep_.CoverTab[98261]++
															if r.blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:121
		_go_fuzz_dep_.CoverTab[98264]++
																r.blocks = make(map[string]map[int32]*alterPartitionReassignmentsBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:122
		// _ = "end of CoverTab[98264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:123
		_go_fuzz_dep_.CoverTab[98265]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:123
		// _ = "end of CoverTab[98265]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:123
	// _ = "end of CoverTab[98261]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:123
	_go_fuzz_dep_.CoverTab[98262]++

															if r.blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:125
		_go_fuzz_dep_.CoverTab[98266]++
																r.blocks[topic] = make(map[int32]*alterPartitionReassignmentsBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:126
		// _ = "end of CoverTab[98266]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:127
		_go_fuzz_dep_.CoverTab[98267]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:127
		// _ = "end of CoverTab[98267]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:127
	// _ = "end of CoverTab[98262]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:127
	_go_fuzz_dep_.CoverTab[98263]++

															r.blocks[topic][partitionID] = &alterPartitionReassignmentsBlock{replicas}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:129
	// _ = "end of CoverTab[98263]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_request.go:130
var _ = _go_fuzz_dep_.CoverTab
