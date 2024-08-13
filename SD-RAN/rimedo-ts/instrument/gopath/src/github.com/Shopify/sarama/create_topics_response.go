//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:1
)

import (
	"fmt"
	"time"
)

type CreateTopicsResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	TopicErrors	map[string]*TopicError
}

func (c *CreateTopicsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:14
	_go_fuzz_dep_.CoverTab[101623]++
													if c.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:15
		_go_fuzz_dep_.CoverTab[101627]++
														pe.putInt32(int32(c.ThrottleTime / time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:16
		// _ = "end of CoverTab[101627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:17
		_go_fuzz_dep_.CoverTab[101628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:17
		// _ = "end of CoverTab[101628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:17
	// _ = "end of CoverTab[101623]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:17
	_go_fuzz_dep_.CoverTab[101624]++

													if err := pe.putArrayLength(len(c.TopicErrors)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:19
		_go_fuzz_dep_.CoverTab[101629]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:20
		// _ = "end of CoverTab[101629]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:21
		_go_fuzz_dep_.CoverTab[101630]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:21
		// _ = "end of CoverTab[101630]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:21
	// _ = "end of CoverTab[101624]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:21
	_go_fuzz_dep_.CoverTab[101625]++
													for topic, topicError := range c.TopicErrors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:22
		_go_fuzz_dep_.CoverTab[101631]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:23
			_go_fuzz_dep_.CoverTab[101633]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:24
			// _ = "end of CoverTab[101633]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:25
			_go_fuzz_dep_.CoverTab[101634]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:25
			// _ = "end of CoverTab[101634]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:25
		// _ = "end of CoverTab[101631]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:25
		_go_fuzz_dep_.CoverTab[101632]++
														if err := topicError.encode(pe, c.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:26
			_go_fuzz_dep_.CoverTab[101635]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:27
			// _ = "end of CoverTab[101635]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:28
			_go_fuzz_dep_.CoverTab[101636]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:28
			// _ = "end of CoverTab[101636]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:28
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:28
		// _ = "end of CoverTab[101632]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:29
	// _ = "end of CoverTab[101625]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:29
	_go_fuzz_dep_.CoverTab[101626]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:31
	// _ = "end of CoverTab[101626]"
}

func (c *CreateTopicsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:34
	_go_fuzz_dep_.CoverTab[101637]++
													c.Version = version

													if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:37
		_go_fuzz_dep_.CoverTab[101641]++
														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:39
			_go_fuzz_dep_.CoverTab[101643]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:40
			// _ = "end of CoverTab[101643]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:41
			_go_fuzz_dep_.CoverTab[101644]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:41
			// _ = "end of CoverTab[101644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:41
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:41
		// _ = "end of CoverTab[101641]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:41
		_go_fuzz_dep_.CoverTab[101642]++
														c.ThrottleTime = time.Duration(throttleTime) * time.Millisecond
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:42
		// _ = "end of CoverTab[101642]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:43
		_go_fuzz_dep_.CoverTab[101645]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:43
		// _ = "end of CoverTab[101645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:43
	// _ = "end of CoverTab[101637]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:43
	_go_fuzz_dep_.CoverTab[101638]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:46
		_go_fuzz_dep_.CoverTab[101646]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:47
		// _ = "end of CoverTab[101646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:48
		_go_fuzz_dep_.CoverTab[101647]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:48
		// _ = "end of CoverTab[101647]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:48
	// _ = "end of CoverTab[101638]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:48
	_go_fuzz_dep_.CoverTab[101639]++

													c.TopicErrors = make(map[string]*TopicError, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:51
		_go_fuzz_dep_.CoverTab[101648]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:53
			_go_fuzz_dep_.CoverTab[101650]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:54
			// _ = "end of CoverTab[101650]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:55
			_go_fuzz_dep_.CoverTab[101651]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:55
			// _ = "end of CoverTab[101651]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:55
		// _ = "end of CoverTab[101648]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:55
		_go_fuzz_dep_.CoverTab[101649]++
														c.TopicErrors[topic] = new(TopicError)
														if err := c.TopicErrors[topic].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:57
			_go_fuzz_dep_.CoverTab[101652]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:58
			// _ = "end of CoverTab[101652]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:59
			_go_fuzz_dep_.CoverTab[101653]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:59
			// _ = "end of CoverTab[101653]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:59
		// _ = "end of CoverTab[101649]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:60
	// _ = "end of CoverTab[101639]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:60
	_go_fuzz_dep_.CoverTab[101640]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:62
	// _ = "end of CoverTab[101640]"
}

func (c *CreateTopicsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:65
	_go_fuzz_dep_.CoverTab[101654]++
													return 19
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:66
	// _ = "end of CoverTab[101654]"
}

func (c *CreateTopicsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:69
	_go_fuzz_dep_.CoverTab[101655]++
													return c.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:70
	// _ = "end of CoverTab[101655]"
}

func (c *CreateTopicsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:73
	_go_fuzz_dep_.CoverTab[101656]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:74
	// _ = "end of CoverTab[101656]"
}

func (c *CreateTopicsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:77
	_go_fuzz_dep_.CoverTab[101657]++
													switch c.Version {
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:79
		_go_fuzz_dep_.CoverTab[101658]++
														return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:80
		// _ = "end of CoverTab[101658]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:81
		_go_fuzz_dep_.CoverTab[101659]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:82
		// _ = "end of CoverTab[101659]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:83
		_go_fuzz_dep_.CoverTab[101660]++
														return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:84
		// _ = "end of CoverTab[101660]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:85
	// _ = "end of CoverTab[101657]"
}

type TopicError struct {
	Err	KError
	ErrMsg	*string
}

func (t *TopicError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:93
	_go_fuzz_dep_.CoverTab[101661]++
													text := t.Err.Error()
													if t.ErrMsg != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:95
		_go_fuzz_dep_.CoverTab[101663]++
														text = fmt.Sprintf("%s - %s", text, *t.ErrMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:96
		// _ = "end of CoverTab[101663]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:97
		_go_fuzz_dep_.CoverTab[101664]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:97
		// _ = "end of CoverTab[101664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:97
	// _ = "end of CoverTab[101661]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:97
	_go_fuzz_dep_.CoverTab[101662]++
													return text
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:98
	// _ = "end of CoverTab[101662]"
}

func (t *TopicError) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:101
	_go_fuzz_dep_.CoverTab[101665]++
													pe.putInt16(int16(t.Err))

													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:104
		_go_fuzz_dep_.CoverTab[101667]++
														if err := pe.putNullableString(t.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:105
			_go_fuzz_dep_.CoverTab[101668]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:106
			// _ = "end of CoverTab[101668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:107
			_go_fuzz_dep_.CoverTab[101669]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:107
			// _ = "end of CoverTab[101669]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:107
		// _ = "end of CoverTab[101667]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:108
		_go_fuzz_dep_.CoverTab[101670]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:108
		// _ = "end of CoverTab[101670]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:108
	// _ = "end of CoverTab[101665]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:108
	_go_fuzz_dep_.CoverTab[101666]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:110
	// _ = "end of CoverTab[101666]"
}

func (t *TopicError) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:113
	_go_fuzz_dep_.CoverTab[101671]++
													kErr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:115
		_go_fuzz_dep_.CoverTab[101674]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:116
		// _ = "end of CoverTab[101674]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:117
		_go_fuzz_dep_.CoverTab[101675]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:117
		// _ = "end of CoverTab[101675]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:117
	// _ = "end of CoverTab[101671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:117
	_go_fuzz_dep_.CoverTab[101672]++
													t.Err = KError(kErr)

													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:120
		_go_fuzz_dep_.CoverTab[101676]++
														if t.ErrMsg, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:121
			_go_fuzz_dep_.CoverTab[101677]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:122
			// _ = "end of CoverTab[101677]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:123
			_go_fuzz_dep_.CoverTab[101678]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:123
			// _ = "end of CoverTab[101678]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:123
		// _ = "end of CoverTab[101676]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:124
		_go_fuzz_dep_.CoverTab[101679]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:124
		// _ = "end of CoverTab[101679]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:124
	// _ = "end of CoverTab[101672]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:124
	_go_fuzz_dep_.CoverTab[101673]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:126
	// _ = "end of CoverTab[101673]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:127
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/create_topics_response.go:127
var _ = _go_fuzz_dep_.CoverTab
