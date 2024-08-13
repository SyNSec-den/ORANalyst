//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:1
)

import (
	"time"
)

type CreateTopicsRequest struct {
	Version	int16

	TopicDetails	map[string]*TopicDetail
	Timeout		time.Duration
	ValidateOnly	bool
}

func (c *CreateTopicsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:15
	_go_fuzz_dep_.CoverTab[101533]++
													if err := pe.putArrayLength(len(c.TopicDetails)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:16
		_go_fuzz_dep_.CoverTab[101537]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:17
		// _ = "end of CoverTab[101537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:18
		_go_fuzz_dep_.CoverTab[101538]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:18
		// _ = "end of CoverTab[101538]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:18
	// _ = "end of CoverTab[101533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:18
	_go_fuzz_dep_.CoverTab[101534]++
													for topic, detail := range c.TopicDetails {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:19
		_go_fuzz_dep_.CoverTab[101539]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:20
			_go_fuzz_dep_.CoverTab[101541]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:21
			// _ = "end of CoverTab[101541]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:22
			_go_fuzz_dep_.CoverTab[101542]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:22
			// _ = "end of CoverTab[101542]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:22
		// _ = "end of CoverTab[101539]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:22
		_go_fuzz_dep_.CoverTab[101540]++
														if err := detail.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:23
			_go_fuzz_dep_.CoverTab[101543]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:24
			// _ = "end of CoverTab[101543]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:25
			_go_fuzz_dep_.CoverTab[101544]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:25
			// _ = "end of CoverTab[101544]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:25
		// _ = "end of CoverTab[101540]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:26
	// _ = "end of CoverTab[101534]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:26
	_go_fuzz_dep_.CoverTab[101535]++

													pe.putInt32(int32(c.Timeout / time.Millisecond))

													if c.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:30
		_go_fuzz_dep_.CoverTab[101545]++
														pe.putBool(c.ValidateOnly)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:31
		// _ = "end of CoverTab[101545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:32
		_go_fuzz_dep_.CoverTab[101546]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:32
		// _ = "end of CoverTab[101546]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:32
	// _ = "end of CoverTab[101535]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:32
	_go_fuzz_dep_.CoverTab[101536]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:34
	// _ = "end of CoverTab[101536]"
}

func (c *CreateTopicsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:37
	_go_fuzz_dep_.CoverTab[101547]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:39
		_go_fuzz_dep_.CoverTab[101552]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:40
		// _ = "end of CoverTab[101552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:41
		_go_fuzz_dep_.CoverTab[101553]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:41
		// _ = "end of CoverTab[101553]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:41
	// _ = "end of CoverTab[101547]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:41
	_go_fuzz_dep_.CoverTab[101548]++

													c.TopicDetails = make(map[string]*TopicDetail, n)

													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:45
		_go_fuzz_dep_.CoverTab[101554]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:47
			_go_fuzz_dep_.CoverTab[101556]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:48
			// _ = "end of CoverTab[101556]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:49
			_go_fuzz_dep_.CoverTab[101557]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:49
			// _ = "end of CoverTab[101557]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:49
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:49
		// _ = "end of CoverTab[101554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:49
		_go_fuzz_dep_.CoverTab[101555]++
														c.TopicDetails[topic] = new(TopicDetail)
														if err = c.TopicDetails[topic].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:51
			_go_fuzz_dep_.CoverTab[101558]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:52
			// _ = "end of CoverTab[101558]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:53
			_go_fuzz_dep_.CoverTab[101559]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:53
			// _ = "end of CoverTab[101559]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:53
		// _ = "end of CoverTab[101555]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:54
	// _ = "end of CoverTab[101548]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:54
	_go_fuzz_dep_.CoverTab[101549]++

													timeout, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:57
		_go_fuzz_dep_.CoverTab[101560]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:58
		// _ = "end of CoverTab[101560]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:59
		_go_fuzz_dep_.CoverTab[101561]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:59
		// _ = "end of CoverTab[101561]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:59
	// _ = "end of CoverTab[101549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:59
	_go_fuzz_dep_.CoverTab[101550]++
													c.Timeout = time.Duration(timeout) * time.Millisecond

													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:62
		_go_fuzz_dep_.CoverTab[101562]++
														c.ValidateOnly, err = pd.getBool()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:64
			_go_fuzz_dep_.CoverTab[101564]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:65
			// _ = "end of CoverTab[101564]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:66
			_go_fuzz_dep_.CoverTab[101565]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:66
			// _ = "end of CoverTab[101565]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:66
		// _ = "end of CoverTab[101562]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:66
		_go_fuzz_dep_.CoverTab[101563]++

														c.Version = version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:68
		// _ = "end of CoverTab[101563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:69
		_go_fuzz_dep_.CoverTab[101566]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:69
		// _ = "end of CoverTab[101566]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:69
	// _ = "end of CoverTab[101550]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:69
	_go_fuzz_dep_.CoverTab[101551]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:71
	// _ = "end of CoverTab[101551]"
}

func (c *CreateTopicsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:74
	_go_fuzz_dep_.CoverTab[101567]++
													return 19
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:75
	// _ = "end of CoverTab[101567]"
}

func (c *CreateTopicsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:78
	_go_fuzz_dep_.CoverTab[101568]++
													return c.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:79
	// _ = "end of CoverTab[101568]"
}

func (r *CreateTopicsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:82
	_go_fuzz_dep_.CoverTab[101569]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:83
	// _ = "end of CoverTab[101569]"
}

func (c *CreateTopicsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:86
	_go_fuzz_dep_.CoverTab[101570]++
													switch c.Version {
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:88
		_go_fuzz_dep_.CoverTab[101571]++
														return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:89
		// _ = "end of CoverTab[101571]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:90
		_go_fuzz_dep_.CoverTab[101572]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:91
		// _ = "end of CoverTab[101572]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:92
		_go_fuzz_dep_.CoverTab[101573]++
														return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:93
		// _ = "end of CoverTab[101573]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:94
	// _ = "end of CoverTab[101570]"
}

type TopicDetail struct {
	NumPartitions		int32
	ReplicationFactor	int16
	ReplicaAssignment	map[int32][]int32
	ConfigEntries		map[string]*string
}

func (t *TopicDetail) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:104
	_go_fuzz_dep_.CoverTab[101574]++
													pe.putInt32(t.NumPartitions)
													pe.putInt16(t.ReplicationFactor)

													if err := pe.putArrayLength(len(t.ReplicaAssignment)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:108
		_go_fuzz_dep_.CoverTab[101579]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:109
		// _ = "end of CoverTab[101579]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:110
		_go_fuzz_dep_.CoverTab[101580]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:110
		// _ = "end of CoverTab[101580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:110
	// _ = "end of CoverTab[101574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:110
	_go_fuzz_dep_.CoverTab[101575]++
													for partition, assignment := range t.ReplicaAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:111
		_go_fuzz_dep_.CoverTab[101581]++
														pe.putInt32(partition)
														if err := pe.putInt32Array(assignment); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:113
			_go_fuzz_dep_.CoverTab[101582]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:114
			// _ = "end of CoverTab[101582]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:115
			_go_fuzz_dep_.CoverTab[101583]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:115
			// _ = "end of CoverTab[101583]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:115
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:115
		// _ = "end of CoverTab[101581]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:116
	// _ = "end of CoverTab[101575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:116
	_go_fuzz_dep_.CoverTab[101576]++

													if err := pe.putArrayLength(len(t.ConfigEntries)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:118
		_go_fuzz_dep_.CoverTab[101584]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:119
		// _ = "end of CoverTab[101584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:120
		_go_fuzz_dep_.CoverTab[101585]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:120
		// _ = "end of CoverTab[101585]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:120
	// _ = "end of CoverTab[101576]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:120
	_go_fuzz_dep_.CoverTab[101577]++
													for configKey, configValue := range t.ConfigEntries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:121
		_go_fuzz_dep_.CoverTab[101586]++
														if err := pe.putString(configKey); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:122
			_go_fuzz_dep_.CoverTab[101588]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:123
			// _ = "end of CoverTab[101588]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:124
			_go_fuzz_dep_.CoverTab[101589]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:124
			// _ = "end of CoverTab[101589]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:124
		// _ = "end of CoverTab[101586]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:124
		_go_fuzz_dep_.CoverTab[101587]++
														if err := pe.putNullableString(configValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:125
			_go_fuzz_dep_.CoverTab[101590]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:126
			// _ = "end of CoverTab[101590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:127
			_go_fuzz_dep_.CoverTab[101591]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:127
			// _ = "end of CoverTab[101591]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:127
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:127
		// _ = "end of CoverTab[101587]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:128
	// _ = "end of CoverTab[101577]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:128
	_go_fuzz_dep_.CoverTab[101578]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:130
	// _ = "end of CoverTab[101578]"
}

func (t *TopicDetail) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:133
	_go_fuzz_dep_.CoverTab[101592]++
													if t.NumPartitions, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:134
		_go_fuzz_dep_.CoverTab[101599]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:135
		// _ = "end of CoverTab[101599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:136
		_go_fuzz_dep_.CoverTab[101600]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:136
		// _ = "end of CoverTab[101600]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:136
	// _ = "end of CoverTab[101592]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:136
	_go_fuzz_dep_.CoverTab[101593]++
													if t.ReplicationFactor, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:137
		_go_fuzz_dep_.CoverTab[101601]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:138
		// _ = "end of CoverTab[101601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:139
		_go_fuzz_dep_.CoverTab[101602]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:139
		// _ = "end of CoverTab[101602]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:139
	// _ = "end of CoverTab[101593]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:139
	_go_fuzz_dep_.CoverTab[101594]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:142
		_go_fuzz_dep_.CoverTab[101603]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:143
		// _ = "end of CoverTab[101603]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:144
		_go_fuzz_dep_.CoverTab[101604]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:144
		// _ = "end of CoverTab[101604]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:144
	// _ = "end of CoverTab[101594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:144
	_go_fuzz_dep_.CoverTab[101595]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:146
		_go_fuzz_dep_.CoverTab[101605]++
														t.ReplicaAssignment = make(map[int32][]int32, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:148
			_go_fuzz_dep_.CoverTab[101606]++
															replica, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:150
				_go_fuzz_dep_.CoverTab[101608]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:151
				// _ = "end of CoverTab[101608]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:152
				_go_fuzz_dep_.CoverTab[101609]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:152
				// _ = "end of CoverTab[101609]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:152
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:152
			// _ = "end of CoverTab[101606]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:152
			_go_fuzz_dep_.CoverTab[101607]++
															if t.ReplicaAssignment[replica], err = pd.getInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:153
				_go_fuzz_dep_.CoverTab[101610]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:154
				// _ = "end of CoverTab[101610]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:155
				_go_fuzz_dep_.CoverTab[101611]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:155
				// _ = "end of CoverTab[101611]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:155
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:155
			// _ = "end of CoverTab[101607]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:156
		// _ = "end of CoverTab[101605]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:157
		_go_fuzz_dep_.CoverTab[101612]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:157
		// _ = "end of CoverTab[101612]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:157
	// _ = "end of CoverTab[101595]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:157
	_go_fuzz_dep_.CoverTab[101596]++

													n, err = pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:160
		_go_fuzz_dep_.CoverTab[101613]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:161
		// _ = "end of CoverTab[101613]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:162
		_go_fuzz_dep_.CoverTab[101614]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:162
		// _ = "end of CoverTab[101614]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:162
	// _ = "end of CoverTab[101596]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:162
	_go_fuzz_dep_.CoverTab[101597]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:164
		_go_fuzz_dep_.CoverTab[101615]++
														t.ConfigEntries = make(map[string]*string, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:166
			_go_fuzz_dep_.CoverTab[101616]++
															configKey, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:168
				_go_fuzz_dep_.CoverTab[101618]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:169
				// _ = "end of CoverTab[101618]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:170
				_go_fuzz_dep_.CoverTab[101619]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:170
				// _ = "end of CoverTab[101619]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:170
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:170
			// _ = "end of CoverTab[101616]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:170
			_go_fuzz_dep_.CoverTab[101617]++
															if t.ConfigEntries[configKey], err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:171
				_go_fuzz_dep_.CoverTab[101620]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:172
				// _ = "end of CoverTab[101620]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:173
				_go_fuzz_dep_.CoverTab[101621]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:173
				// _ = "end of CoverTab[101621]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:173
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:173
			// _ = "end of CoverTab[101617]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:174
		// _ = "end of CoverTab[101615]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:175
		_go_fuzz_dep_.CoverTab[101622]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:175
		// _ = "end of CoverTab[101622]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:175
	// _ = "end of CoverTab[101597]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:175
	_go_fuzz_dep_.CoverTab[101598]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:177
	// _ = "end of CoverTab[101598]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:178
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_request.go:178
var _ = _go_fuzz_dep_.CoverTab
