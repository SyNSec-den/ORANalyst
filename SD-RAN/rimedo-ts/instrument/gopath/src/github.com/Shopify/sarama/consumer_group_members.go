//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:1
)

// ConsumerGroupMemberMetadata holds the metadata for consumer group
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:3
// https://github.com/apache/kafka/blob/trunk/clients/src/main/resources/common/message/ConsumerProtocolSubscription.json
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:5
type ConsumerGroupMemberMetadata struct {
	Version		int16
	Topics		[]string
	UserData	[]byte
	OwnedPartitions	[]*OwnedPartition
}

func (m *ConsumerGroupMemberMetadata) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:12
	_go_fuzz_dep_.CoverTab[101282]++
													pe.putInt16(m.Version)

													if err := pe.putStringArray(m.Topics); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:15
		_go_fuzz_dep_.CoverTab[101285]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:16
		// _ = "end of CoverTab[101285]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:17
		_go_fuzz_dep_.CoverTab[101286]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:17
		// _ = "end of CoverTab[101286]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:17
	// _ = "end of CoverTab[101282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:17
	_go_fuzz_dep_.CoverTab[101283]++

													if err := pe.putBytes(m.UserData); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:19
		_go_fuzz_dep_.CoverTab[101287]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:20
		// _ = "end of CoverTab[101287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:21
		_go_fuzz_dep_.CoverTab[101288]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:21
		// _ = "end of CoverTab[101288]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:21
	// _ = "end of CoverTab[101283]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:21
	_go_fuzz_dep_.CoverTab[101284]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:23
	// _ = "end of CoverTab[101284]"
}

func (m *ConsumerGroupMemberMetadata) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:26
	_go_fuzz_dep_.CoverTab[101289]++
													if m.Version, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:27
		_go_fuzz_dep_.CoverTab[101294]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:28
		// _ = "end of CoverTab[101294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:29
		_go_fuzz_dep_.CoverTab[101295]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:29
		// _ = "end of CoverTab[101295]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:29
	// _ = "end of CoverTab[101289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:29
	_go_fuzz_dep_.CoverTab[101290]++

													if m.Topics, err = pd.getStringArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:31
		_go_fuzz_dep_.CoverTab[101296]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:32
		// _ = "end of CoverTab[101296]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:33
		_go_fuzz_dep_.CoverTab[101297]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:33
		// _ = "end of CoverTab[101297]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:33
	// _ = "end of CoverTab[101290]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:33
	_go_fuzz_dep_.CoverTab[101291]++

													if m.UserData, err = pd.getBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:35
		_go_fuzz_dep_.CoverTab[101298]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:36
		// _ = "end of CoverTab[101298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:37
		_go_fuzz_dep_.CoverTab[101299]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:37
		// _ = "end of CoverTab[101299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:37
	// _ = "end of CoverTab[101291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:37
	_go_fuzz_dep_.CoverTab[101292]++
													if m.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:38
		_go_fuzz_dep_.CoverTab[101300]++
														n, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:40
			_go_fuzz_dep_.CoverTab[101303]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:44
			if err == ErrInsufficientData {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:44
				_go_fuzz_dep_.CoverTab[101305]++
																return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:45
				// _ = "end of CoverTab[101305]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:46
				_go_fuzz_dep_.CoverTab[101306]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:46
				// _ = "end of CoverTab[101306]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:46
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:46
			// _ = "end of CoverTab[101303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:46
			_go_fuzz_dep_.CoverTab[101304]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:47
			// _ = "end of CoverTab[101304]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:48
			_go_fuzz_dep_.CoverTab[101307]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:48
			// _ = "end of CoverTab[101307]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:48
		// _ = "end of CoverTab[101300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:48
		_go_fuzz_dep_.CoverTab[101301]++
														if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:49
			_go_fuzz_dep_.CoverTab[101308]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:50
			// _ = "end of CoverTab[101308]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:51
			_go_fuzz_dep_.CoverTab[101309]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:51
			// _ = "end of CoverTab[101309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:51
		// _ = "end of CoverTab[101301]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:51
		_go_fuzz_dep_.CoverTab[101302]++
														m.OwnedPartitions = make([]*OwnedPartition, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:53
			_go_fuzz_dep_.CoverTab[101310]++
															m.OwnedPartitions[i] = &OwnedPartition{}
															if err := m.OwnedPartitions[i].decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:55
				_go_fuzz_dep_.CoverTab[101311]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:56
				// _ = "end of CoverTab[101311]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:57
				_go_fuzz_dep_.CoverTab[101312]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:57
				// _ = "end of CoverTab[101312]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:57
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:57
			// _ = "end of CoverTab[101310]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:58
		// _ = "end of CoverTab[101302]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:59
		_go_fuzz_dep_.CoverTab[101313]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:59
		// _ = "end of CoverTab[101313]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:59
	// _ = "end of CoverTab[101292]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:59
	_go_fuzz_dep_.CoverTab[101293]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:61
	// _ = "end of CoverTab[101293]"
}

type OwnedPartition struct {
	Topic		string
	Partitions	[]int32
}

func (m *OwnedPartition) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:69
	_go_fuzz_dep_.CoverTab[101314]++
													if m.Topic, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:70
		_go_fuzz_dep_.CoverTab[101317]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:71
		// _ = "end of CoverTab[101317]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:72
		_go_fuzz_dep_.CoverTab[101318]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:72
		// _ = "end of CoverTab[101318]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:72
	// _ = "end of CoverTab[101314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:72
	_go_fuzz_dep_.CoverTab[101315]++
													if m.Partitions, err = pd.getInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:73
		_go_fuzz_dep_.CoverTab[101319]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:74
		// _ = "end of CoverTab[101319]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:75
		_go_fuzz_dep_.CoverTab[101320]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:75
		// _ = "end of CoverTab[101320]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:75
	// _ = "end of CoverTab[101315]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:75
	_go_fuzz_dep_.CoverTab[101316]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:77
	// _ = "end of CoverTab[101316]"
}

// ConsumerGroupMemberAssignment holds the member assignment for a consume group
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:80
// https://github.com/apache/kafka/blob/trunk/clients/src/main/resources/common/message/ConsumerProtocolAssignment.json
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:82
type ConsumerGroupMemberAssignment struct {
	Version		int16
	Topics		map[string][]int32
	UserData	[]byte
}

func (m *ConsumerGroupMemberAssignment) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:88
	_go_fuzz_dep_.CoverTab[101321]++
													pe.putInt16(m.Version)

													if err := pe.putArrayLength(len(m.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:91
		_go_fuzz_dep_.CoverTab[101325]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:92
		// _ = "end of CoverTab[101325]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:93
		_go_fuzz_dep_.CoverTab[101326]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:93
		// _ = "end of CoverTab[101326]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:93
	// _ = "end of CoverTab[101321]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:93
	_go_fuzz_dep_.CoverTab[101322]++

													for topic, partitions := range m.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:95
		_go_fuzz_dep_.CoverTab[101327]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:96
			_go_fuzz_dep_.CoverTab[101329]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:97
			// _ = "end of CoverTab[101329]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:98
			_go_fuzz_dep_.CoverTab[101330]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:98
			// _ = "end of CoverTab[101330]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:98
		// _ = "end of CoverTab[101327]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:98
		_go_fuzz_dep_.CoverTab[101328]++
														if err := pe.putInt32Array(partitions); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:99
			_go_fuzz_dep_.CoverTab[101331]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:100
			// _ = "end of CoverTab[101331]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:101
			_go_fuzz_dep_.CoverTab[101332]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:101
			// _ = "end of CoverTab[101332]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:101
		// _ = "end of CoverTab[101328]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:102
	// _ = "end of CoverTab[101322]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:102
	_go_fuzz_dep_.CoverTab[101323]++

													if err := pe.putBytes(m.UserData); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:104
		_go_fuzz_dep_.CoverTab[101333]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:105
		// _ = "end of CoverTab[101333]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:106
		_go_fuzz_dep_.CoverTab[101334]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:106
		// _ = "end of CoverTab[101334]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:106
	// _ = "end of CoverTab[101323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:106
	_go_fuzz_dep_.CoverTab[101324]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:108
	// _ = "end of CoverTab[101324]"
}

func (m *ConsumerGroupMemberAssignment) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:111
	_go_fuzz_dep_.CoverTab[101335]++
													if m.Version, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:112
		_go_fuzz_dep_.CoverTab[101340]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:113
		// _ = "end of CoverTab[101340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:114
		_go_fuzz_dep_.CoverTab[101341]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:114
		// _ = "end of CoverTab[101341]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:114
	// _ = "end of CoverTab[101335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:114
	_go_fuzz_dep_.CoverTab[101336]++

													var topicLen int
													if topicLen, err = pd.getArrayLength(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:117
		_go_fuzz_dep_.CoverTab[101342]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:118
		// _ = "end of CoverTab[101342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:119
		_go_fuzz_dep_.CoverTab[101343]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:119
		// _ = "end of CoverTab[101343]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:119
	// _ = "end of CoverTab[101336]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:119
	_go_fuzz_dep_.CoverTab[101337]++

													m.Topics = make(map[string][]int32, topicLen)
													for i := 0; i < topicLen; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:122
		_go_fuzz_dep_.CoverTab[101344]++
														var topic string
														if topic, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:124
			_go_fuzz_dep_.CoverTab[101346]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:125
			// _ = "end of CoverTab[101346]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:126
			_go_fuzz_dep_.CoverTab[101347]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:126
			// _ = "end of CoverTab[101347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:126
		// _ = "end of CoverTab[101344]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:126
		_go_fuzz_dep_.CoverTab[101345]++
														if m.Topics[topic], err = pd.getInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:127
			_go_fuzz_dep_.CoverTab[101348]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:128
			// _ = "end of CoverTab[101348]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:129
			_go_fuzz_dep_.CoverTab[101349]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:129
			// _ = "end of CoverTab[101349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:129
		// _ = "end of CoverTab[101345]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:130
	// _ = "end of CoverTab[101337]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:130
	_go_fuzz_dep_.CoverTab[101338]++

													if m.UserData, err = pd.getBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:132
		_go_fuzz_dep_.CoverTab[101350]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:133
		// _ = "end of CoverTab[101350]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:134
		_go_fuzz_dep_.CoverTab[101351]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:134
		// _ = "end of CoverTab[101351]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:134
	// _ = "end of CoverTab[101338]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:134
	_go_fuzz_dep_.CoverTab[101339]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:136
	// _ = "end of CoverTab[101339]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:137
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group_members.go:137
var _ = _go_fuzz_dep_.CoverTab
