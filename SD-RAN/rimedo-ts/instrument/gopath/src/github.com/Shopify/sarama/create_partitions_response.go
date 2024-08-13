//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:1
)

import (
	"fmt"
	"time"
)

type CreatePartitionsResponse struct {
	ThrottleTime		time.Duration
	TopicPartitionErrors	map[string]*TopicPartitionError
}

func (c *CreatePartitionsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:13
	_go_fuzz_dep_.CoverTab[101489]++
													pe.putInt32(int32(c.ThrottleTime / time.Millisecond))
													if err := pe.putArrayLength(len(c.TopicPartitionErrors)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:15
		_go_fuzz_dep_.CoverTab[101492]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:16
		// _ = "end of CoverTab[101492]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:17
		_go_fuzz_dep_.CoverTab[101493]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:17
		// _ = "end of CoverTab[101493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:17
	// _ = "end of CoverTab[101489]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:17
	_go_fuzz_dep_.CoverTab[101490]++

													for topic, partitionError := range c.TopicPartitionErrors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:19
		_go_fuzz_dep_.CoverTab[101494]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:20
			_go_fuzz_dep_.CoverTab[101496]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:21
			// _ = "end of CoverTab[101496]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:22
			_go_fuzz_dep_.CoverTab[101497]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:22
			// _ = "end of CoverTab[101497]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:22
		// _ = "end of CoverTab[101494]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:22
		_go_fuzz_dep_.CoverTab[101495]++
														if err := partitionError.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:23
			_go_fuzz_dep_.CoverTab[101498]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:24
			// _ = "end of CoverTab[101498]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:25
			_go_fuzz_dep_.CoverTab[101499]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:25
			// _ = "end of CoverTab[101499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:25
		// _ = "end of CoverTab[101495]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:26
	// _ = "end of CoverTab[101490]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:26
	_go_fuzz_dep_.CoverTab[101491]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:28
	// _ = "end of CoverTab[101491]"
}

func (c *CreatePartitionsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:31
	_go_fuzz_dep_.CoverTab[101500]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:33
		_go_fuzz_dep_.CoverTab[101504]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:34
		// _ = "end of CoverTab[101504]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:35
		_go_fuzz_dep_.CoverTab[101505]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:35
		// _ = "end of CoverTab[101505]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:35
	// _ = "end of CoverTab[101500]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:35
	_go_fuzz_dep_.CoverTab[101501]++
													c.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:39
		_go_fuzz_dep_.CoverTab[101506]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:40
		// _ = "end of CoverTab[101506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:41
		_go_fuzz_dep_.CoverTab[101507]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:41
		// _ = "end of CoverTab[101507]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:41
	// _ = "end of CoverTab[101501]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:41
	_go_fuzz_dep_.CoverTab[101502]++

													c.TopicPartitionErrors = make(map[string]*TopicPartitionError, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:44
		_go_fuzz_dep_.CoverTab[101508]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:46
			_go_fuzz_dep_.CoverTab[101510]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:47
			// _ = "end of CoverTab[101510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:48
			_go_fuzz_dep_.CoverTab[101511]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:48
			// _ = "end of CoverTab[101511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:48
		// _ = "end of CoverTab[101508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:48
		_go_fuzz_dep_.CoverTab[101509]++
														c.TopicPartitionErrors[topic] = new(TopicPartitionError)
														if err := c.TopicPartitionErrors[topic].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:50
			_go_fuzz_dep_.CoverTab[101512]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:51
			// _ = "end of CoverTab[101512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:52
			_go_fuzz_dep_.CoverTab[101513]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:52
			// _ = "end of CoverTab[101513]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:52
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:52
		// _ = "end of CoverTab[101509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:53
	// _ = "end of CoverTab[101502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:53
	_go_fuzz_dep_.CoverTab[101503]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:55
	// _ = "end of CoverTab[101503]"
}

func (r *CreatePartitionsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:58
	_go_fuzz_dep_.CoverTab[101514]++
													return 37
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:59
	// _ = "end of CoverTab[101514]"
}

func (r *CreatePartitionsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:62
	_go_fuzz_dep_.CoverTab[101515]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:63
	// _ = "end of CoverTab[101515]"
}

func (r *CreatePartitionsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:66
	_go_fuzz_dep_.CoverTab[101516]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:67
	// _ = "end of CoverTab[101516]"
}

func (r *CreatePartitionsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:70
	_go_fuzz_dep_.CoverTab[101517]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:71
	// _ = "end of CoverTab[101517]"
}

type TopicPartitionError struct {
	Err	KError
	ErrMsg	*string
}

func (t *TopicPartitionError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:79
	_go_fuzz_dep_.CoverTab[101518]++
													text := t.Err.Error()
													if t.ErrMsg != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:81
		_go_fuzz_dep_.CoverTab[101520]++
														text = fmt.Sprintf("%s - %s", text, *t.ErrMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:82
		// _ = "end of CoverTab[101520]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:83
		_go_fuzz_dep_.CoverTab[101521]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:83
		// _ = "end of CoverTab[101521]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:83
	// _ = "end of CoverTab[101518]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:83
	_go_fuzz_dep_.CoverTab[101519]++
													return text
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:84
	// _ = "end of CoverTab[101519]"
}

func (t *TopicPartitionError) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:87
	_go_fuzz_dep_.CoverTab[101522]++
													pe.putInt16(int16(t.Err))

													if err := pe.putNullableString(t.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:90
		_go_fuzz_dep_.CoverTab[101524]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:91
		// _ = "end of CoverTab[101524]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:92
		_go_fuzz_dep_.CoverTab[101525]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:92
		// _ = "end of CoverTab[101525]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:92
	// _ = "end of CoverTab[101522]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:92
	_go_fuzz_dep_.CoverTab[101523]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:94
	// _ = "end of CoverTab[101523]"
}

func (t *TopicPartitionError) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:97
	_go_fuzz_dep_.CoverTab[101526]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:99
		_go_fuzz_dep_.CoverTab[101529]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:100
		// _ = "end of CoverTab[101529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:101
		_go_fuzz_dep_.CoverTab[101530]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:101
		// _ = "end of CoverTab[101530]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:101
	// _ = "end of CoverTab[101526]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:101
	_go_fuzz_dep_.CoverTab[101527]++
													t.Err = KError(kerr)

													if t.ErrMsg, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:104
		_go_fuzz_dep_.CoverTab[101531]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:105
		// _ = "end of CoverTab[101531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:106
		_go_fuzz_dep_.CoverTab[101532]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:106
		// _ = "end of CoverTab[101532]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:106
	// _ = "end of CoverTab[101527]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:106
	_go_fuzz_dep_.CoverTab[101528]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:108
	// _ = "end of CoverTab[101528]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_partitions_response.go:109
var _ = _go_fuzz_dep_.CoverTab
