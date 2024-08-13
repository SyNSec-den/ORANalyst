//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:1
)

type alterPartitionReassignmentsErrorBlock struct {
	errorCode	KError
	errorMessage	*string
}

func (b *alterPartitionReassignmentsErrorBlock) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:8
	_go_fuzz_dep_.CoverTab[98268]++
															pe.putInt16(int16(b.errorCode))
															if err := pe.putNullableCompactString(b.errorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:10
		_go_fuzz_dep_.CoverTab[98270]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:11
		// _ = "end of CoverTab[98270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:12
		_go_fuzz_dep_.CoverTab[98271]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:12
		// _ = "end of CoverTab[98271]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:12
	// _ = "end of CoverTab[98268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:12
	_go_fuzz_dep_.CoverTab[98269]++
															pe.putEmptyTaggedFieldArray()

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:15
	// _ = "end of CoverTab[98269]"
}

func (b *alterPartitionReassignmentsErrorBlock) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:18
	_go_fuzz_dep_.CoverTab[98272]++
															errorCode, err := pd.getInt16()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:20
		_go_fuzz_dep_.CoverTab[98275]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:21
		// _ = "end of CoverTab[98275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:22
		_go_fuzz_dep_.CoverTab[98276]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:22
		// _ = "end of CoverTab[98276]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:22
	// _ = "end of CoverTab[98272]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:22
	_go_fuzz_dep_.CoverTab[98273]++
															b.errorCode = KError(errorCode)
															b.errorMessage, err = pd.getCompactNullableString()

															if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:26
		_go_fuzz_dep_.CoverTab[98277]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:27
		// _ = "end of CoverTab[98277]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:28
		_go_fuzz_dep_.CoverTab[98278]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:28
		// _ = "end of CoverTab[98278]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:28
	// _ = "end of CoverTab[98273]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:28
	_go_fuzz_dep_.CoverTab[98274]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:29
	// _ = "end of CoverTab[98274]"
}

type AlterPartitionReassignmentsResponse struct {
	Version		int16
	ThrottleTimeMs	int32
	ErrorCode	KError
	ErrorMessage	*string
	Errors		map[string]map[int32]*alterPartitionReassignmentsErrorBlock
}

func (r *AlterPartitionReassignmentsResponse) AddError(topic string, partition int32, kerror KError, message *string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:40
	_go_fuzz_dep_.CoverTab[98279]++
															if r.Errors == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:41
		_go_fuzz_dep_.CoverTab[98282]++
																r.Errors = make(map[string]map[int32]*alterPartitionReassignmentsErrorBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:42
		// _ = "end of CoverTab[98282]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:43
		_go_fuzz_dep_.CoverTab[98283]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:43
		// _ = "end of CoverTab[98283]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:43
	// _ = "end of CoverTab[98279]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:43
	_go_fuzz_dep_.CoverTab[98280]++
															partitions := r.Errors[topic]
															if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:45
		_go_fuzz_dep_.CoverTab[98284]++
																partitions = make(map[int32]*alterPartitionReassignmentsErrorBlock)
																r.Errors[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:47
		// _ = "end of CoverTab[98284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:48
		_go_fuzz_dep_.CoverTab[98285]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:48
		// _ = "end of CoverTab[98285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:48
	// _ = "end of CoverTab[98280]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:48
	_go_fuzz_dep_.CoverTab[98281]++

															partitions[partition] = &alterPartitionReassignmentsErrorBlock{errorCode: kerror, errorMessage: message}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:50
	// _ = "end of CoverTab[98281]"
}

func (r *AlterPartitionReassignmentsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:53
	_go_fuzz_dep_.CoverTab[98286]++
															pe.putInt32(r.ThrottleTimeMs)
															pe.putInt16(int16(r.ErrorCode))
															if err := pe.putNullableCompactString(r.ErrorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:56
		_go_fuzz_dep_.CoverTab[98289]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:57
		// _ = "end of CoverTab[98289]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:58
		_go_fuzz_dep_.CoverTab[98290]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:58
		// _ = "end of CoverTab[98290]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:58
	// _ = "end of CoverTab[98286]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:58
	_go_fuzz_dep_.CoverTab[98287]++

															pe.putCompactArrayLength(len(r.Errors))
															for topic, partitions := range r.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:61
		_go_fuzz_dep_.CoverTab[98291]++
																if err := pe.putCompactString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:62
			_go_fuzz_dep_.CoverTab[98294]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:63
			// _ = "end of CoverTab[98294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:64
			_go_fuzz_dep_.CoverTab[98295]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:64
			// _ = "end of CoverTab[98295]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:64
		// _ = "end of CoverTab[98291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:64
		_go_fuzz_dep_.CoverTab[98292]++
																pe.putCompactArrayLength(len(partitions))
																for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:66
			_go_fuzz_dep_.CoverTab[98296]++
																	pe.putInt32(partition)

																	if err := block.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:69
				_go_fuzz_dep_.CoverTab[98297]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:70
				// _ = "end of CoverTab[98297]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:71
				_go_fuzz_dep_.CoverTab[98298]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:71
				// _ = "end of CoverTab[98298]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:71
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:71
			// _ = "end of CoverTab[98296]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:72
		// _ = "end of CoverTab[98292]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:72
		_go_fuzz_dep_.CoverTab[98293]++
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:73
		// _ = "end of CoverTab[98293]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:74
	// _ = "end of CoverTab[98287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:74
	_go_fuzz_dep_.CoverTab[98288]++

															pe.putEmptyTaggedFieldArray()
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:77
	// _ = "end of CoverTab[98288]"
}

func (r *AlterPartitionReassignmentsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:80
	_go_fuzz_dep_.CoverTab[98299]++
															r.Version = version

															if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:83
		_go_fuzz_dep_.CoverTab[98306]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:84
		// _ = "end of CoverTab[98306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:85
		_go_fuzz_dep_.CoverTab[98307]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:85
		// _ = "end of CoverTab[98307]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:85
	// _ = "end of CoverTab[98299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:85
	_go_fuzz_dep_.CoverTab[98300]++

															kerr, err := pd.getInt16()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:88
		_go_fuzz_dep_.CoverTab[98308]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:89
		// _ = "end of CoverTab[98308]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:90
		_go_fuzz_dep_.CoverTab[98309]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:90
		// _ = "end of CoverTab[98309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:90
	// _ = "end of CoverTab[98300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:90
	_go_fuzz_dep_.CoverTab[98301]++

															r.ErrorCode = KError(kerr)

															if r.ErrorMessage, err = pd.getCompactNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:94
		_go_fuzz_dep_.CoverTab[98310]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:95
		// _ = "end of CoverTab[98310]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:96
		_go_fuzz_dep_.CoverTab[98311]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:96
		// _ = "end of CoverTab[98311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:96
	// _ = "end of CoverTab[98301]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:96
	_go_fuzz_dep_.CoverTab[98302]++

															numTopics, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:99
		_go_fuzz_dep_.CoverTab[98312]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:100
		// _ = "end of CoverTab[98312]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:101
		_go_fuzz_dep_.CoverTab[98313]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:101
		// _ = "end of CoverTab[98313]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:101
	// _ = "end of CoverTab[98302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:101
	_go_fuzz_dep_.CoverTab[98303]++

															if numTopics > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:103
		_go_fuzz_dep_.CoverTab[98314]++
																r.Errors = make(map[string]map[int32]*alterPartitionReassignmentsErrorBlock, numTopics)
																for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:105
			_go_fuzz_dep_.CoverTab[98315]++
																	topic, err := pd.getCompactString()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:107
				_go_fuzz_dep_.CoverTab[98319]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:108
				// _ = "end of CoverTab[98319]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:109
				_go_fuzz_dep_.CoverTab[98320]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:109
				// _ = "end of CoverTab[98320]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:109
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:109
			// _ = "end of CoverTab[98315]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:109
			_go_fuzz_dep_.CoverTab[98316]++

																	ongoingPartitionReassignments, err := pd.getCompactArrayLength()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:112
				_go_fuzz_dep_.CoverTab[98321]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:113
				// _ = "end of CoverTab[98321]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:114
				_go_fuzz_dep_.CoverTab[98322]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:114
				// _ = "end of CoverTab[98322]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:114
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:114
			// _ = "end of CoverTab[98316]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:114
			_go_fuzz_dep_.CoverTab[98317]++

																	r.Errors[topic] = make(map[int32]*alterPartitionReassignmentsErrorBlock, ongoingPartitionReassignments)

																	for j := 0; j < ongoingPartitionReassignments; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:118
				_go_fuzz_dep_.CoverTab[98323]++
																		partition, err := pd.getInt32()
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:120
					_go_fuzz_dep_.CoverTab[98326]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:121
					// _ = "end of CoverTab[98326]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:122
					_go_fuzz_dep_.CoverTab[98327]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:122
					// _ = "end of CoverTab[98327]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:122
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:122
				// _ = "end of CoverTab[98323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:122
				_go_fuzz_dep_.CoverTab[98324]++
																		block := &alterPartitionReassignmentsErrorBlock{}
																		if err := block.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:124
					_go_fuzz_dep_.CoverTab[98328]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:125
					// _ = "end of CoverTab[98328]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:126
					_go_fuzz_dep_.CoverTab[98329]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:126
					// _ = "end of CoverTab[98329]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:126
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:126
				// _ = "end of CoverTab[98324]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:126
				_go_fuzz_dep_.CoverTab[98325]++

																		r.Errors[topic][partition] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:128
				// _ = "end of CoverTab[98325]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:129
			// _ = "end of CoverTab[98317]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:129
			_go_fuzz_dep_.CoverTab[98318]++
																	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:130
				_go_fuzz_dep_.CoverTab[98330]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:131
				// _ = "end of CoverTab[98330]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:132
				_go_fuzz_dep_.CoverTab[98331]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:132
				// _ = "end of CoverTab[98331]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:132
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:132
			// _ = "end of CoverTab[98318]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:133
		// _ = "end of CoverTab[98314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:134
		_go_fuzz_dep_.CoverTab[98332]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:134
		// _ = "end of CoverTab[98332]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:134
	// _ = "end of CoverTab[98303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:134
	_go_fuzz_dep_.CoverTab[98304]++

															if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:136
		_go_fuzz_dep_.CoverTab[98333]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:137
		// _ = "end of CoverTab[98333]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:138
		_go_fuzz_dep_.CoverTab[98334]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:138
		// _ = "end of CoverTab[98334]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:138
	// _ = "end of CoverTab[98304]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:138
	_go_fuzz_dep_.CoverTab[98305]++

															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:140
	// _ = "end of CoverTab[98305]"
}

func (r *AlterPartitionReassignmentsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:143
	_go_fuzz_dep_.CoverTab[98335]++
															return 45
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:144
	// _ = "end of CoverTab[98335]"
}

func (r *AlterPartitionReassignmentsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:147
	_go_fuzz_dep_.CoverTab[98336]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:148
	// _ = "end of CoverTab[98336]"
}

func (r *AlterPartitionReassignmentsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:151
	_go_fuzz_dep_.CoverTab[98337]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:152
	// _ = "end of CoverTab[98337]"
}

func (r *AlterPartitionReassignmentsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:155
	_go_fuzz_dep_.CoverTab[98338]++
															return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:156
	// _ = "end of CoverTab[98338]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:157
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_partition_reassignments_response.go:157
var _ = _go_fuzz_dep_.CoverTab
