//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:1
)

type PartitionReplicaReassignmentsStatus struct {
	Replicas		[]int32
	AddingReplicas		[]int32
	RemovingReplicas	[]int32
}

func (b *PartitionReplicaReassignmentsStatus) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:9
	_go_fuzz_dep_.CoverTab[103869]++
															if err := pe.putCompactInt32Array(b.Replicas); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:10
		_go_fuzz_dep_.CoverTab[103873]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:11
		// _ = "end of CoverTab[103873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:12
		_go_fuzz_dep_.CoverTab[103874]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:12
		// _ = "end of CoverTab[103874]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:12
	// _ = "end of CoverTab[103869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:12
	_go_fuzz_dep_.CoverTab[103870]++
															if err := pe.putCompactInt32Array(b.AddingReplicas); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:13
		_go_fuzz_dep_.CoverTab[103875]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:14
		// _ = "end of CoverTab[103875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:15
		_go_fuzz_dep_.CoverTab[103876]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:15
		// _ = "end of CoverTab[103876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:15
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:15
	// _ = "end of CoverTab[103870]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:15
	_go_fuzz_dep_.CoverTab[103871]++
															if err := pe.putCompactInt32Array(b.RemovingReplicas); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:16
		_go_fuzz_dep_.CoverTab[103877]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:17
		// _ = "end of CoverTab[103877]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:18
		_go_fuzz_dep_.CoverTab[103878]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:18
		// _ = "end of CoverTab[103878]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:18
	// _ = "end of CoverTab[103871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:18
	_go_fuzz_dep_.CoverTab[103872]++

															pe.putEmptyTaggedFieldArray()

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:22
	// _ = "end of CoverTab[103872]"
}

func (b *PartitionReplicaReassignmentsStatus) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:25
	_go_fuzz_dep_.CoverTab[103879]++
															if b.Replicas, err = pd.getCompactInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:26
		_go_fuzz_dep_.CoverTab[103884]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:27
		// _ = "end of CoverTab[103884]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:28
		_go_fuzz_dep_.CoverTab[103885]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:28
		// _ = "end of CoverTab[103885]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:28
	// _ = "end of CoverTab[103879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:28
	_go_fuzz_dep_.CoverTab[103880]++

															if b.AddingReplicas, err = pd.getCompactInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:30
		_go_fuzz_dep_.CoverTab[103886]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:31
		// _ = "end of CoverTab[103886]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:32
		_go_fuzz_dep_.CoverTab[103887]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:32
		// _ = "end of CoverTab[103887]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:32
	// _ = "end of CoverTab[103880]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:32
	_go_fuzz_dep_.CoverTab[103881]++

															if b.RemovingReplicas, err = pd.getCompactInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:34
		_go_fuzz_dep_.CoverTab[103888]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:35
		// _ = "end of CoverTab[103888]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:36
		_go_fuzz_dep_.CoverTab[103889]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:36
		// _ = "end of CoverTab[103889]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:36
	// _ = "end of CoverTab[103881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:36
	_go_fuzz_dep_.CoverTab[103882]++

															if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:38
		_go_fuzz_dep_.CoverTab[103890]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:39
		// _ = "end of CoverTab[103890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:40
		_go_fuzz_dep_.CoverTab[103891]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:40
		// _ = "end of CoverTab[103891]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:40
	// _ = "end of CoverTab[103882]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:40
	_go_fuzz_dep_.CoverTab[103883]++

															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:42
	// _ = "end of CoverTab[103883]"
}

type ListPartitionReassignmentsResponse struct {
	Version		int16
	ThrottleTimeMs	int32
	ErrorCode	KError
	ErrorMessage	*string
	TopicStatus	map[string]map[int32]*PartitionReplicaReassignmentsStatus
}

func (r *ListPartitionReassignmentsResponse) AddBlock(topic string, partition int32, replicas, addingReplicas, removingReplicas []int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:53
	_go_fuzz_dep_.CoverTab[103892]++
															if r.TopicStatus == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:54
		_go_fuzz_dep_.CoverTab[103895]++
																r.TopicStatus = make(map[string]map[int32]*PartitionReplicaReassignmentsStatus)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:55
		// _ = "end of CoverTab[103895]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:56
		_go_fuzz_dep_.CoverTab[103896]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:56
		// _ = "end of CoverTab[103896]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:56
	// _ = "end of CoverTab[103892]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:56
	_go_fuzz_dep_.CoverTab[103893]++
															partitions := r.TopicStatus[topic]
															if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:58
		_go_fuzz_dep_.CoverTab[103897]++
																partitions = make(map[int32]*PartitionReplicaReassignmentsStatus)
																r.TopicStatus[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:60
		// _ = "end of CoverTab[103897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:61
		_go_fuzz_dep_.CoverTab[103898]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:61
		// _ = "end of CoverTab[103898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:61
	// _ = "end of CoverTab[103893]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:61
	_go_fuzz_dep_.CoverTab[103894]++

															partitions[partition] = &PartitionReplicaReassignmentsStatus{Replicas: replicas, AddingReplicas: addingReplicas, RemovingReplicas: removingReplicas}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:63
	// _ = "end of CoverTab[103894]"
}

func (r *ListPartitionReassignmentsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:66
	_go_fuzz_dep_.CoverTab[103899]++
															pe.putInt32(r.ThrottleTimeMs)
															pe.putInt16(int16(r.ErrorCode))
															if err := pe.putNullableCompactString(r.ErrorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:69
		_go_fuzz_dep_.CoverTab[103902]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:70
		// _ = "end of CoverTab[103902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:71
		_go_fuzz_dep_.CoverTab[103903]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:71
		// _ = "end of CoverTab[103903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:71
	// _ = "end of CoverTab[103899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:71
	_go_fuzz_dep_.CoverTab[103900]++

															pe.putCompactArrayLength(len(r.TopicStatus))
															for topic, partitions := range r.TopicStatus {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:74
		_go_fuzz_dep_.CoverTab[103904]++
																if err := pe.putCompactString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:75
			_go_fuzz_dep_.CoverTab[103907]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:76
			// _ = "end of CoverTab[103907]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:77
			_go_fuzz_dep_.CoverTab[103908]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:77
			// _ = "end of CoverTab[103908]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:77
		// _ = "end of CoverTab[103904]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:77
		_go_fuzz_dep_.CoverTab[103905]++
																pe.putCompactArrayLength(len(partitions))
																for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:79
			_go_fuzz_dep_.CoverTab[103909]++
																	pe.putInt32(partition)

																	if err := block.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:82
				_go_fuzz_dep_.CoverTab[103910]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:83
				// _ = "end of CoverTab[103910]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:84
				_go_fuzz_dep_.CoverTab[103911]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:84
				// _ = "end of CoverTab[103911]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:84
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:84
			// _ = "end of CoverTab[103909]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:85
		// _ = "end of CoverTab[103905]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:85
		_go_fuzz_dep_.CoverTab[103906]++
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:86
		// _ = "end of CoverTab[103906]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:87
	// _ = "end of CoverTab[103900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:87
	_go_fuzz_dep_.CoverTab[103901]++

															pe.putEmptyTaggedFieldArray()

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:91
	// _ = "end of CoverTab[103901]"
}

func (r *ListPartitionReassignmentsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:94
	_go_fuzz_dep_.CoverTab[103912]++
															r.Version = version

															if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:97
		_go_fuzz_dep_.CoverTab[103919]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:98
		// _ = "end of CoverTab[103919]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:99
		_go_fuzz_dep_.CoverTab[103920]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:99
		// _ = "end of CoverTab[103920]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:99
	// _ = "end of CoverTab[103912]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:99
	_go_fuzz_dep_.CoverTab[103913]++

															kerr, err := pd.getInt16()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:102
		_go_fuzz_dep_.CoverTab[103921]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:103
		// _ = "end of CoverTab[103921]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:104
		_go_fuzz_dep_.CoverTab[103922]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:104
		// _ = "end of CoverTab[103922]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:104
	// _ = "end of CoverTab[103913]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:104
	_go_fuzz_dep_.CoverTab[103914]++

															r.ErrorCode = KError(kerr)

															if r.ErrorMessage, err = pd.getCompactNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:108
		_go_fuzz_dep_.CoverTab[103923]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:109
		// _ = "end of CoverTab[103923]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:110
		_go_fuzz_dep_.CoverTab[103924]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:110
		// _ = "end of CoverTab[103924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:110
	// _ = "end of CoverTab[103914]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:110
	_go_fuzz_dep_.CoverTab[103915]++

															numTopics, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:113
		_go_fuzz_dep_.CoverTab[103925]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:114
		// _ = "end of CoverTab[103925]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:115
		_go_fuzz_dep_.CoverTab[103926]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:115
		// _ = "end of CoverTab[103926]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:115
	// _ = "end of CoverTab[103915]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:115
	_go_fuzz_dep_.CoverTab[103916]++

															r.TopicStatus = make(map[string]map[int32]*PartitionReplicaReassignmentsStatus, numTopics)
															for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:118
		_go_fuzz_dep_.CoverTab[103927]++
																topic, err := pd.getCompactString()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:120
			_go_fuzz_dep_.CoverTab[103931]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:121
			// _ = "end of CoverTab[103931]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:122
			_go_fuzz_dep_.CoverTab[103932]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:122
			// _ = "end of CoverTab[103932]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:122
		// _ = "end of CoverTab[103927]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:122
		_go_fuzz_dep_.CoverTab[103928]++

																ongoingPartitionReassignments, err := pd.getCompactArrayLength()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:125
			_go_fuzz_dep_.CoverTab[103933]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:126
			// _ = "end of CoverTab[103933]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:127
			_go_fuzz_dep_.CoverTab[103934]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:127
			// _ = "end of CoverTab[103934]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:127
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:127
		// _ = "end of CoverTab[103928]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:127
		_go_fuzz_dep_.CoverTab[103929]++

																r.TopicStatus[topic] = make(map[int32]*PartitionReplicaReassignmentsStatus, ongoingPartitionReassignments)

																for j := 0; j < ongoingPartitionReassignments; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:131
			_go_fuzz_dep_.CoverTab[103935]++
																	partition, err := pd.getInt32()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:133
				_go_fuzz_dep_.CoverTab[103938]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:134
				// _ = "end of CoverTab[103938]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:135
				_go_fuzz_dep_.CoverTab[103939]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:135
				// _ = "end of CoverTab[103939]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:135
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:135
			// _ = "end of CoverTab[103935]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:135
			_go_fuzz_dep_.CoverTab[103936]++

																	block := &PartitionReplicaReassignmentsStatus{}
																	if err := block.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:138
				_go_fuzz_dep_.CoverTab[103940]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:139
				// _ = "end of CoverTab[103940]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:140
				_go_fuzz_dep_.CoverTab[103941]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:140
				// _ = "end of CoverTab[103941]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:140
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:140
			// _ = "end of CoverTab[103936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:140
			_go_fuzz_dep_.CoverTab[103937]++
																	r.TopicStatus[topic][partition] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:141
			// _ = "end of CoverTab[103937]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:142
		// _ = "end of CoverTab[103929]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:142
		_go_fuzz_dep_.CoverTab[103930]++

																if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:144
			_go_fuzz_dep_.CoverTab[103942]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:145
			// _ = "end of CoverTab[103942]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:146
			_go_fuzz_dep_.CoverTab[103943]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:146
			// _ = "end of CoverTab[103943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:146
		// _ = "end of CoverTab[103930]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:147
	// _ = "end of CoverTab[103916]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:147
	_go_fuzz_dep_.CoverTab[103917]++
															if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:148
		_go_fuzz_dep_.CoverTab[103944]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:149
		// _ = "end of CoverTab[103944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:150
		_go_fuzz_dep_.CoverTab[103945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:150
		// _ = "end of CoverTab[103945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:150
	// _ = "end of CoverTab[103917]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:150
	_go_fuzz_dep_.CoverTab[103918]++

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:152
	// _ = "end of CoverTab[103918]"
}

func (r *ListPartitionReassignmentsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:155
	_go_fuzz_dep_.CoverTab[103946]++
															return 46
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:156
	// _ = "end of CoverTab[103946]"
}

func (r *ListPartitionReassignmentsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:159
	_go_fuzz_dep_.CoverTab[103947]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:160
	// _ = "end of CoverTab[103947]"
}

func (r *ListPartitionReassignmentsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:163
	_go_fuzz_dep_.CoverTab[103948]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:164
	// _ = "end of CoverTab[103948]"
}

func (r *ListPartitionReassignmentsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:167
	_go_fuzz_dep_.CoverTab[103949]++
															return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:168
	// _ = "end of CoverTab[103949]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:169
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_partition_reassignments_response.go:169
var _ = _go_fuzz_dep_.CoverTab
