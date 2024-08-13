//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:1
)

type topicPartitionAssignment struct {
	Topic		string
	Partition	int32
}

type StickyAssignorUserData interface {
	partitions() []topicPartitionAssignment
	hasGeneration() bool
	generation() int
}

// StickyAssignorUserDataV0 holds topic partition information for an assignment
type StickyAssignorUserDataV0 struct {
	Topics	map[string][]int32

	topicPartitions	[]topicPartitionAssignment
}

func (m *StickyAssignorUserDataV0) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:21
	_go_fuzz_dep_.CoverTab[106723]++
													if err := pe.putArrayLength(len(m.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:22
		_go_fuzz_dep_.CoverTab[106726]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:23
		// _ = "end of CoverTab[106726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:24
		_go_fuzz_dep_.CoverTab[106727]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:24
		// _ = "end of CoverTab[106727]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:24
	// _ = "end of CoverTab[106723]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:24
	_go_fuzz_dep_.CoverTab[106724]++

													for topic, partitions := range m.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:26
		_go_fuzz_dep_.CoverTab[106728]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:27
			_go_fuzz_dep_.CoverTab[106730]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:28
			// _ = "end of CoverTab[106730]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:29
			_go_fuzz_dep_.CoverTab[106731]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:29
			// _ = "end of CoverTab[106731]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:29
		// _ = "end of CoverTab[106728]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:29
		_go_fuzz_dep_.CoverTab[106729]++
														if err := pe.putInt32Array(partitions); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:30
			_go_fuzz_dep_.CoverTab[106732]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:31
			// _ = "end of CoverTab[106732]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:32
			_go_fuzz_dep_.CoverTab[106733]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:32
			// _ = "end of CoverTab[106733]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:32
		// _ = "end of CoverTab[106729]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:33
	// _ = "end of CoverTab[106724]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:33
	_go_fuzz_dep_.CoverTab[106725]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:34
	// _ = "end of CoverTab[106725]"
}

func (m *StickyAssignorUserDataV0) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:37
	_go_fuzz_dep_.CoverTab[106734]++
													var topicLen int
													if topicLen, err = pd.getArrayLength(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:39
		_go_fuzz_dep_.CoverTab[106737]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:40
		// _ = "end of CoverTab[106737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:41
		_go_fuzz_dep_.CoverTab[106738]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:41
		// _ = "end of CoverTab[106738]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:41
	// _ = "end of CoverTab[106734]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:41
	_go_fuzz_dep_.CoverTab[106735]++

													m.Topics = make(map[string][]int32, topicLen)
													for i := 0; i < topicLen; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:44
		_go_fuzz_dep_.CoverTab[106739]++
														var topic string
														if topic, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:46
			_go_fuzz_dep_.CoverTab[106741]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:47
			// _ = "end of CoverTab[106741]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:48
			_go_fuzz_dep_.CoverTab[106742]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:48
			// _ = "end of CoverTab[106742]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:48
		// _ = "end of CoverTab[106739]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:48
		_go_fuzz_dep_.CoverTab[106740]++
														if m.Topics[topic], err = pd.getInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:49
			_go_fuzz_dep_.CoverTab[106743]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:50
			// _ = "end of CoverTab[106743]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:51
			_go_fuzz_dep_.CoverTab[106744]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:51
			// _ = "end of CoverTab[106744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:51
		// _ = "end of CoverTab[106740]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:52
	// _ = "end of CoverTab[106735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:52
	_go_fuzz_dep_.CoverTab[106736]++
													m.topicPartitions = populateTopicPartitions(m.Topics)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:54
	// _ = "end of CoverTab[106736]"
}

func (m *StickyAssignorUserDataV0) partitions() []topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:57
	_go_fuzz_dep_.CoverTab[106745]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:57
	return m.topicPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:57
	// _ = "end of CoverTab[106745]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:57
}
func (m *StickyAssignorUserDataV0) hasGeneration() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:58
	_go_fuzz_dep_.CoverTab[106746]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:58
	return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:58
	// _ = "end of CoverTab[106746]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:58
}
func (m *StickyAssignorUserDataV0) generation() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:59
	_go_fuzz_dep_.CoverTab[106747]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:59
	return defaultGeneration
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:59
	// _ = "end of CoverTab[106747]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:59
}

// StickyAssignorUserDataV1 holds topic partition information for an assignment
type StickyAssignorUserDataV1 struct {
	Topics		map[string][]int32
	Generation	int32

	topicPartitions	[]topicPartitionAssignment
}

func (m *StickyAssignorUserDataV1) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:69
	_go_fuzz_dep_.CoverTab[106748]++
													if err := pe.putArrayLength(len(m.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:70
		_go_fuzz_dep_.CoverTab[106751]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:71
		// _ = "end of CoverTab[106751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:72
		_go_fuzz_dep_.CoverTab[106752]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:72
		// _ = "end of CoverTab[106752]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:72
	// _ = "end of CoverTab[106748]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:72
	_go_fuzz_dep_.CoverTab[106749]++

													for topic, partitions := range m.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:74
		_go_fuzz_dep_.CoverTab[106753]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:75
			_go_fuzz_dep_.CoverTab[106755]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:76
			// _ = "end of CoverTab[106755]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:77
			_go_fuzz_dep_.CoverTab[106756]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:77
			// _ = "end of CoverTab[106756]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:77
		// _ = "end of CoverTab[106753]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:77
		_go_fuzz_dep_.CoverTab[106754]++
														if err := pe.putInt32Array(partitions); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:78
			_go_fuzz_dep_.CoverTab[106757]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:79
			// _ = "end of CoverTab[106757]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:80
			_go_fuzz_dep_.CoverTab[106758]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:80
			// _ = "end of CoverTab[106758]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:80
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:80
		// _ = "end of CoverTab[106754]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:81
	// _ = "end of CoverTab[106749]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:81
	_go_fuzz_dep_.CoverTab[106750]++

													pe.putInt32(m.Generation)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:84
	// _ = "end of CoverTab[106750]"
}

func (m *StickyAssignorUserDataV1) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:87
	_go_fuzz_dep_.CoverTab[106759]++
													var topicLen int
													if topicLen, err = pd.getArrayLength(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:89
		_go_fuzz_dep_.CoverTab[106763]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:90
		// _ = "end of CoverTab[106763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:91
		_go_fuzz_dep_.CoverTab[106764]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:91
		// _ = "end of CoverTab[106764]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:91
	// _ = "end of CoverTab[106759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:91
	_go_fuzz_dep_.CoverTab[106760]++

													m.Topics = make(map[string][]int32, topicLen)
													for i := 0; i < topicLen; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:94
		_go_fuzz_dep_.CoverTab[106765]++
														var topic string
														if topic, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:96
			_go_fuzz_dep_.CoverTab[106767]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:97
			// _ = "end of CoverTab[106767]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:98
			_go_fuzz_dep_.CoverTab[106768]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:98
			// _ = "end of CoverTab[106768]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:98
		// _ = "end of CoverTab[106765]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:98
		_go_fuzz_dep_.CoverTab[106766]++
														if m.Topics[topic], err = pd.getInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:99
			_go_fuzz_dep_.CoverTab[106769]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:100
			// _ = "end of CoverTab[106769]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:101
			_go_fuzz_dep_.CoverTab[106770]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:101
			// _ = "end of CoverTab[106770]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:101
		// _ = "end of CoverTab[106766]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:102
	// _ = "end of CoverTab[106760]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:102
	_go_fuzz_dep_.CoverTab[106761]++

													m.Generation, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:105
		_go_fuzz_dep_.CoverTab[106771]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:106
		// _ = "end of CoverTab[106771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:107
		_go_fuzz_dep_.CoverTab[106772]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:107
		// _ = "end of CoverTab[106772]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:107
	// _ = "end of CoverTab[106761]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:107
	_go_fuzz_dep_.CoverTab[106762]++
													m.topicPartitions = populateTopicPartitions(m.Topics)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:109
	// _ = "end of CoverTab[106762]"
}

func (m *StickyAssignorUserDataV1) partitions() []topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:112
	_go_fuzz_dep_.CoverTab[106773]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:112
	return m.topicPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:112
	// _ = "end of CoverTab[106773]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:112
}
func (m *StickyAssignorUserDataV1) hasGeneration() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:113
	_go_fuzz_dep_.CoverTab[106774]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:113
	return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:113
	// _ = "end of CoverTab[106774]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:113
}
func (m *StickyAssignorUserDataV1) generation() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:114
	_go_fuzz_dep_.CoverTab[106775]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:114
	return int(m.Generation)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:114
	// _ = "end of CoverTab[106775]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:114
}

func populateTopicPartitions(topics map[string][]int32) []topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:116
	_go_fuzz_dep_.CoverTab[106776]++
													topicPartitions := make([]topicPartitionAssignment, 0)
													for topic, partitions := range topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:118
		_go_fuzz_dep_.CoverTab[106778]++
														for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:119
			_go_fuzz_dep_.CoverTab[106779]++
															topicPartitions = append(topicPartitions, topicPartitionAssignment{Topic: topic, Partition: partition})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:120
			// _ = "end of CoverTab[106779]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:121
		// _ = "end of CoverTab[106778]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:122
	// _ = "end of CoverTab[106776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:122
	_go_fuzz_dep_.CoverTab[106777]++
													return topicPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:123
	// _ = "end of CoverTab[106777]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:124
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sticky_assignor_user_data.go:124
var _ = _go_fuzz_dep_.CoverTab
