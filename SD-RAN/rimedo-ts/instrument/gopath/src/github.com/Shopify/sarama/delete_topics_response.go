//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:1
)

import "time"

type DeleteTopicsResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	TopicErrorCodes	map[string]KError
}

func (d *DeleteTopicsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:11
	_go_fuzz_dep_.CoverTab[101972]++
													if d.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:12
		_go_fuzz_dep_.CoverTab[101976]++
														pe.putInt32(int32(d.ThrottleTime / time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:13
		// _ = "end of CoverTab[101976]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:14
		_go_fuzz_dep_.CoverTab[101977]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:14
		// _ = "end of CoverTab[101977]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:14
	// _ = "end of CoverTab[101972]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:14
	_go_fuzz_dep_.CoverTab[101973]++

													if err := pe.putArrayLength(len(d.TopicErrorCodes)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:16
		_go_fuzz_dep_.CoverTab[101978]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:17
		// _ = "end of CoverTab[101978]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:18
		_go_fuzz_dep_.CoverTab[101979]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:18
		// _ = "end of CoverTab[101979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:18
	// _ = "end of CoverTab[101973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:18
	_go_fuzz_dep_.CoverTab[101974]++
													for topic, errorCode := range d.TopicErrorCodes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:19
		_go_fuzz_dep_.CoverTab[101980]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:20
			_go_fuzz_dep_.CoverTab[101982]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:21
			// _ = "end of CoverTab[101982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:22
			_go_fuzz_dep_.CoverTab[101983]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:22
			// _ = "end of CoverTab[101983]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:22
		// _ = "end of CoverTab[101980]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:22
		_go_fuzz_dep_.CoverTab[101981]++
														pe.putInt16(int16(errorCode))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:23
		// _ = "end of CoverTab[101981]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:24
	// _ = "end of CoverTab[101974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:24
	_go_fuzz_dep_.CoverTab[101975]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:26
	// _ = "end of CoverTab[101975]"
}

func (d *DeleteTopicsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:29
	_go_fuzz_dep_.CoverTab[101984]++
													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:30
		_go_fuzz_dep_.CoverTab[101988]++
														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:32
			_go_fuzz_dep_.CoverTab[101990]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:33
			// _ = "end of CoverTab[101990]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:34
			_go_fuzz_dep_.CoverTab[101991]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:34
			// _ = "end of CoverTab[101991]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:34
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:34
		// _ = "end of CoverTab[101988]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:34
		_go_fuzz_dep_.CoverTab[101989]++
														d.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

														d.Version = version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:37
		// _ = "end of CoverTab[101989]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:38
		_go_fuzz_dep_.CoverTab[101992]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:38
		// _ = "end of CoverTab[101992]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:38
	// _ = "end of CoverTab[101984]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:38
	_go_fuzz_dep_.CoverTab[101985]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:41
		_go_fuzz_dep_.CoverTab[101993]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:42
		// _ = "end of CoverTab[101993]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:43
		_go_fuzz_dep_.CoverTab[101994]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:43
		// _ = "end of CoverTab[101994]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:43
	// _ = "end of CoverTab[101985]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:43
	_go_fuzz_dep_.CoverTab[101986]++

													d.TopicErrorCodes = make(map[string]KError, n)

													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:47
		_go_fuzz_dep_.CoverTab[101995]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:49
			_go_fuzz_dep_.CoverTab[101998]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:50
			// _ = "end of CoverTab[101998]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:51
			_go_fuzz_dep_.CoverTab[101999]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:51
			// _ = "end of CoverTab[101999]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:51
		// _ = "end of CoverTab[101995]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:51
		_go_fuzz_dep_.CoverTab[101996]++
														errorCode, err := pd.getInt16()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:53
			_go_fuzz_dep_.CoverTab[102000]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:54
			// _ = "end of CoverTab[102000]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:55
			_go_fuzz_dep_.CoverTab[102001]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:55
			// _ = "end of CoverTab[102001]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:55
		// _ = "end of CoverTab[101996]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:55
		_go_fuzz_dep_.CoverTab[101997]++

														d.TopicErrorCodes[topic] = KError(errorCode)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:57
		// _ = "end of CoverTab[101997]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:58
	// _ = "end of CoverTab[101986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:58
	_go_fuzz_dep_.CoverTab[101987]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:60
	// _ = "end of CoverTab[101987]"
}

func (d *DeleteTopicsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:63
	_go_fuzz_dep_.CoverTab[102002]++
													return 20
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:64
	// _ = "end of CoverTab[102002]"
}

func (d *DeleteTopicsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:67
	_go_fuzz_dep_.CoverTab[102003]++
													return d.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:68
	// _ = "end of CoverTab[102003]"
}

func (d *DeleteTopicsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:71
	_go_fuzz_dep_.CoverTab[102004]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:72
	// _ = "end of CoverTab[102004]"
}

func (d *DeleteTopicsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:75
	_go_fuzz_dep_.CoverTab[102005]++
													switch d.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:77
		_go_fuzz_dep_.CoverTab[102006]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:78
		// _ = "end of CoverTab[102006]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:79
		_go_fuzz_dep_.CoverTab[102007]++
														return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:80
		// _ = "end of CoverTab[102007]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:81
	// _ = "end of CoverTab[102005]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:82
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_response.go:82
var _ = _go_fuzz_dep_.CoverTab
