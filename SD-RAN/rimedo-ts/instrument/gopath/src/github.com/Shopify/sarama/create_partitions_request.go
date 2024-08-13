//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:1
)

import "time"

type CreatePartitionsRequest struct {
	TopicPartitions	map[string]*TopicPartition
	Timeout		time.Duration
	ValidateOnly	bool
}

func (c *CreatePartitionsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:11
	_go_fuzz_dep_.CoverTab[101432]++
													if err := pe.putArrayLength(len(c.TopicPartitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:12
		_go_fuzz_dep_.CoverTab[101435]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:13
		// _ = "end of CoverTab[101435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:14
		_go_fuzz_dep_.CoverTab[101436]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:14
		// _ = "end of CoverTab[101436]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:14
	// _ = "end of CoverTab[101432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:14
	_go_fuzz_dep_.CoverTab[101433]++

													for topic, partition := range c.TopicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:16
		_go_fuzz_dep_.CoverTab[101437]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:17
			_go_fuzz_dep_.CoverTab[101439]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:18
			// _ = "end of CoverTab[101439]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:19
			_go_fuzz_dep_.CoverTab[101440]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:19
			// _ = "end of CoverTab[101440]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:19
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:19
		// _ = "end of CoverTab[101437]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:19
		_go_fuzz_dep_.CoverTab[101438]++
														if err := partition.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:20
			_go_fuzz_dep_.CoverTab[101441]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:21
			// _ = "end of CoverTab[101441]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:22
			_go_fuzz_dep_.CoverTab[101442]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:22
			// _ = "end of CoverTab[101442]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:22
		// _ = "end of CoverTab[101438]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:23
	// _ = "end of CoverTab[101433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:23
	_go_fuzz_dep_.CoverTab[101434]++

													pe.putInt32(int32(c.Timeout / time.Millisecond))

													pe.putBool(c.ValidateOnly)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:29
	// _ = "end of CoverTab[101434]"
}

func (c *CreatePartitionsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:32
	_go_fuzz_dep_.CoverTab[101443]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:34
		_go_fuzz_dep_.CoverTab[101448]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:35
		// _ = "end of CoverTab[101448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:36
		_go_fuzz_dep_.CoverTab[101449]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:36
		// _ = "end of CoverTab[101449]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:36
	// _ = "end of CoverTab[101443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:36
	_go_fuzz_dep_.CoverTab[101444]++
													c.TopicPartitions = make(map[string]*TopicPartition, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:38
		_go_fuzz_dep_.CoverTab[101450]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:40
			_go_fuzz_dep_.CoverTab[101452]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:41
			// _ = "end of CoverTab[101452]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:42
			_go_fuzz_dep_.CoverTab[101453]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:42
			// _ = "end of CoverTab[101453]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:42
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:42
		// _ = "end of CoverTab[101450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:42
		_go_fuzz_dep_.CoverTab[101451]++
														c.TopicPartitions[topic] = new(TopicPartition)
														if err := c.TopicPartitions[topic].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:44
			_go_fuzz_dep_.CoverTab[101454]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:45
			// _ = "end of CoverTab[101454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:46
			_go_fuzz_dep_.CoverTab[101455]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:46
			// _ = "end of CoverTab[101455]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:46
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:46
		// _ = "end of CoverTab[101451]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:47
	// _ = "end of CoverTab[101444]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:47
	_go_fuzz_dep_.CoverTab[101445]++

													timeout, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:50
		_go_fuzz_dep_.CoverTab[101456]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:51
		// _ = "end of CoverTab[101456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:52
		_go_fuzz_dep_.CoverTab[101457]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:52
		// _ = "end of CoverTab[101457]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:52
	// _ = "end of CoverTab[101445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:52
	_go_fuzz_dep_.CoverTab[101446]++
													c.Timeout = time.Duration(timeout) * time.Millisecond

													if c.ValidateOnly, err = pd.getBool(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:55
		_go_fuzz_dep_.CoverTab[101458]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:56
		// _ = "end of CoverTab[101458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:57
		_go_fuzz_dep_.CoverTab[101459]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:57
		// _ = "end of CoverTab[101459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:57
	// _ = "end of CoverTab[101446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:57
	_go_fuzz_dep_.CoverTab[101447]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:59
	// _ = "end of CoverTab[101447]"
}

func (r *CreatePartitionsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:62
	_go_fuzz_dep_.CoverTab[101460]++
													return 37
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:63
	// _ = "end of CoverTab[101460]"
}

func (r *CreatePartitionsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:66
	_go_fuzz_dep_.CoverTab[101461]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:67
	// _ = "end of CoverTab[101461]"
}

func (r *CreatePartitionsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:70
	_go_fuzz_dep_.CoverTab[101462]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:71
	// _ = "end of CoverTab[101462]"
}

func (r *CreatePartitionsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:74
	_go_fuzz_dep_.CoverTab[101463]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:75
	// _ = "end of CoverTab[101463]"
}

type TopicPartition struct {
	Count		int32
	Assignment	[][]int32
}

func (t *TopicPartition) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:83
	_go_fuzz_dep_.CoverTab[101464]++
													pe.putInt32(t.Count)

													if len(t.Assignment) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:86
		_go_fuzz_dep_.CoverTab[101468]++
														pe.putInt32(-1)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:88
		// _ = "end of CoverTab[101468]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:89
		_go_fuzz_dep_.CoverTab[101469]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:89
		// _ = "end of CoverTab[101469]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:89
	// _ = "end of CoverTab[101464]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:89
	_go_fuzz_dep_.CoverTab[101465]++

													if err := pe.putArrayLength(len(t.Assignment)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:91
		_go_fuzz_dep_.CoverTab[101470]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:92
		// _ = "end of CoverTab[101470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:93
		_go_fuzz_dep_.CoverTab[101471]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:93
		// _ = "end of CoverTab[101471]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:93
	// _ = "end of CoverTab[101465]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:93
	_go_fuzz_dep_.CoverTab[101466]++

													for _, assign := range t.Assignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:95
		_go_fuzz_dep_.CoverTab[101472]++
														if err := pe.putInt32Array(assign); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:96
			_go_fuzz_dep_.CoverTab[101473]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:97
			// _ = "end of CoverTab[101473]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:98
			_go_fuzz_dep_.CoverTab[101474]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:98
			// _ = "end of CoverTab[101474]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:98
		// _ = "end of CoverTab[101472]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:99
	// _ = "end of CoverTab[101466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:99
	_go_fuzz_dep_.CoverTab[101467]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:101
	// _ = "end of CoverTab[101467]"
}

func (t *TopicPartition) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:104
	_go_fuzz_dep_.CoverTab[101475]++
													if t.Count, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:105
		_go_fuzz_dep_.CoverTab[101480]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:106
		// _ = "end of CoverTab[101480]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:107
		_go_fuzz_dep_.CoverTab[101481]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:107
		// _ = "end of CoverTab[101481]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:107
	// _ = "end of CoverTab[101475]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:107
	_go_fuzz_dep_.CoverTab[101476]++

													n, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:110
		_go_fuzz_dep_.CoverTab[101482]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:111
		// _ = "end of CoverTab[101482]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:112
		_go_fuzz_dep_.CoverTab[101483]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:112
		// _ = "end of CoverTab[101483]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:112
	// _ = "end of CoverTab[101476]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:112
	_go_fuzz_dep_.CoverTab[101477]++
													if n <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:113
		_go_fuzz_dep_.CoverTab[101484]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:114
		// _ = "end of CoverTab[101484]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:115
		_go_fuzz_dep_.CoverTab[101485]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:115
		// _ = "end of CoverTab[101485]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:115
	// _ = "end of CoverTab[101477]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:115
	_go_fuzz_dep_.CoverTab[101478]++
													t.Assignment = make([][]int32, n)

													for i := 0; i < int(n); i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:118
		_go_fuzz_dep_.CoverTab[101486]++
														if t.Assignment[i], err = pd.getInt32Array(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:119
			_go_fuzz_dep_.CoverTab[101487]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:120
			// _ = "end of CoverTab[101487]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:121
			_go_fuzz_dep_.CoverTab[101488]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:121
			// _ = "end of CoverTab[101488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:121
		// _ = "end of CoverTab[101486]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:122
	// _ = "end of CoverTab[101478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:122
	_go_fuzz_dep_.CoverTab[101479]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:124
	// _ = "end of CoverTab[101479]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_request.go:125
var _ = _go_fuzz_dep_.CoverTab
