//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:1
)

import (
	"fmt"
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:23
// partition_responses in protocol
type ProduceResponseBlock struct {
	Err		KError		// v0, error_code
	Offset		int64		// v0, base_offset
	Timestamp	time.Time	// v2, log_append_time, and the broker is configured with `LogAppendTime`
	StartOffset	int64		// v5, log_start_offset
}

func (b *ProduceResponseBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:31
	_go_fuzz_dep_.CoverTab[105775]++
												tmp, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:33
		_go_fuzz_dep_.CoverTab[105780]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:34
		// _ = "end of CoverTab[105780]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:35
		_go_fuzz_dep_.CoverTab[105781]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:35
		// _ = "end of CoverTab[105781]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:35
	// _ = "end of CoverTab[105775]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:35
	_go_fuzz_dep_.CoverTab[105776]++
												b.Err = KError(tmp)

												b.Offset, err = pd.getInt64()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:39
		_go_fuzz_dep_.CoverTab[105782]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:40
		// _ = "end of CoverTab[105782]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:41
		_go_fuzz_dep_.CoverTab[105783]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:41
		// _ = "end of CoverTab[105783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:41
	// _ = "end of CoverTab[105776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:41
	_go_fuzz_dep_.CoverTab[105777]++

												if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:43
		_go_fuzz_dep_.CoverTab[105784]++
													if millis, err := pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:44
			_go_fuzz_dep_.CoverTab[105785]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:45
			// _ = "end of CoverTab[105785]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:46
			_go_fuzz_dep_.CoverTab[105786]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:46
			if millis != -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:46
				_go_fuzz_dep_.CoverTab[105787]++
															b.Timestamp = time.Unix(millis/1000, (millis%1000)*int64(time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:47
				// _ = "end of CoverTab[105787]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:48
				_go_fuzz_dep_.CoverTab[105788]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:48
				// _ = "end of CoverTab[105788]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:48
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:48
			// _ = "end of CoverTab[105786]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:48
		// _ = "end of CoverTab[105784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:49
		_go_fuzz_dep_.CoverTab[105789]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:49
		// _ = "end of CoverTab[105789]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:49
	// _ = "end of CoverTab[105777]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:49
	_go_fuzz_dep_.CoverTab[105778]++

												if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:51
		_go_fuzz_dep_.CoverTab[105790]++
													b.StartOffset, err = pd.getInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:53
			_go_fuzz_dep_.CoverTab[105791]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:54
			// _ = "end of CoverTab[105791]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:55
			_go_fuzz_dep_.CoverTab[105792]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:55
			// _ = "end of CoverTab[105792]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:55
		// _ = "end of CoverTab[105790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:56
		_go_fuzz_dep_.CoverTab[105793]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:56
		// _ = "end of CoverTab[105793]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:56
	// _ = "end of CoverTab[105778]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:56
	_go_fuzz_dep_.CoverTab[105779]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:58
	// _ = "end of CoverTab[105779]"
}

func (b *ProduceResponseBlock) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:61
	_go_fuzz_dep_.CoverTab[105794]++
												pe.putInt16(int16(b.Err))
												pe.putInt64(b.Offset)

												if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:65
		_go_fuzz_dep_.CoverTab[105797]++
													timestamp := int64(-1)
													if !b.Timestamp.Before(time.Unix(0, 0)) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:67
			_go_fuzz_dep_.CoverTab[105799]++
														timestamp = b.Timestamp.UnixNano() / int64(time.Millisecond)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:68
			// _ = "end of CoverTab[105799]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:69
			_go_fuzz_dep_.CoverTab[105800]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:69
			if !b.Timestamp.IsZero() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:69
				_go_fuzz_dep_.CoverTab[105801]++
															return PacketEncodingError{fmt.Sprintf("invalid timestamp (%v)", b.Timestamp)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:70
				// _ = "end of CoverTab[105801]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
				_go_fuzz_dep_.CoverTab[105802]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
				// _ = "end of CoverTab[105802]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
			// _ = "end of CoverTab[105800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
		// _ = "end of CoverTab[105797]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:71
		_go_fuzz_dep_.CoverTab[105798]++
													pe.putInt64(timestamp)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:72
		// _ = "end of CoverTab[105798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:73
		_go_fuzz_dep_.CoverTab[105803]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:73
		// _ = "end of CoverTab[105803]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:73
	// _ = "end of CoverTab[105794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:73
	_go_fuzz_dep_.CoverTab[105795]++

												if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:75
		_go_fuzz_dep_.CoverTab[105804]++
													pe.putInt64(b.StartOffset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:76
		// _ = "end of CoverTab[105804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:77
		_go_fuzz_dep_.CoverTab[105805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:77
		// _ = "end of CoverTab[105805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:77
	// _ = "end of CoverTab[105795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:77
	_go_fuzz_dep_.CoverTab[105796]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:79
	// _ = "end of CoverTab[105796]"
}

type ProduceResponse struct {
	Blocks		map[string]map[int32]*ProduceResponseBlock	// v0, responses
	Version		int16
	ThrottleTime	time.Duration	// v1, throttle_time_ms
}

func (r *ProduceResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:88
	_go_fuzz_dep_.CoverTab[105806]++
												r.Version = version

												numTopics, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:92
		_go_fuzz_dep_.CoverTab[105810]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:93
		// _ = "end of CoverTab[105810]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:94
		_go_fuzz_dep_.CoverTab[105811]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:94
		// _ = "end of CoverTab[105811]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:94
	// _ = "end of CoverTab[105806]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:94
	_go_fuzz_dep_.CoverTab[105807]++

												r.Blocks = make(map[string]map[int32]*ProduceResponseBlock, numTopics)
												for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:97
		_go_fuzz_dep_.CoverTab[105812]++
													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:99
			_go_fuzz_dep_.CoverTab[105815]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:100
			// _ = "end of CoverTab[105815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:101
			_go_fuzz_dep_.CoverTab[105816]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:101
			// _ = "end of CoverTab[105816]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:101
		// _ = "end of CoverTab[105812]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:101
		_go_fuzz_dep_.CoverTab[105813]++

													numBlocks, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:104
			_go_fuzz_dep_.CoverTab[105817]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:105
			// _ = "end of CoverTab[105817]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:106
			_go_fuzz_dep_.CoverTab[105818]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:106
			// _ = "end of CoverTab[105818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:106
		// _ = "end of CoverTab[105813]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:106
		_go_fuzz_dep_.CoverTab[105814]++

													r.Blocks[name] = make(map[int32]*ProduceResponseBlock, numBlocks)

													for j := 0; j < numBlocks; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:110
			_go_fuzz_dep_.CoverTab[105819]++
														id, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:112
				_go_fuzz_dep_.CoverTab[105822]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:113
				// _ = "end of CoverTab[105822]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:114
				_go_fuzz_dep_.CoverTab[105823]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:114
				// _ = "end of CoverTab[105823]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:114
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:114
			// _ = "end of CoverTab[105819]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:114
			_go_fuzz_dep_.CoverTab[105820]++

														block := new(ProduceResponseBlock)
														err = block.decode(pd, version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:118
				_go_fuzz_dep_.CoverTab[105824]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:119
				// _ = "end of CoverTab[105824]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:120
				_go_fuzz_dep_.CoverTab[105825]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:120
				// _ = "end of CoverTab[105825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:120
			// _ = "end of CoverTab[105820]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:120
			_go_fuzz_dep_.CoverTab[105821]++
														r.Blocks[name][id] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:121
			// _ = "end of CoverTab[105821]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:122
		// _ = "end of CoverTab[105814]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:123
	// _ = "end of CoverTab[105807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:123
	_go_fuzz_dep_.CoverTab[105808]++

												if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:125
		_go_fuzz_dep_.CoverTab[105826]++
													millis, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:127
			_go_fuzz_dep_.CoverTab[105828]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:128
			// _ = "end of CoverTab[105828]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:129
			_go_fuzz_dep_.CoverTab[105829]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:129
			// _ = "end of CoverTab[105829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:129
		// _ = "end of CoverTab[105826]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:129
		_go_fuzz_dep_.CoverTab[105827]++

													r.ThrottleTime = time.Duration(millis) * time.Millisecond
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:131
		// _ = "end of CoverTab[105827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:132
		_go_fuzz_dep_.CoverTab[105830]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:132
		// _ = "end of CoverTab[105830]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:132
	// _ = "end of CoverTab[105808]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:132
	_go_fuzz_dep_.CoverTab[105809]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:134
	// _ = "end of CoverTab[105809]"
}

func (r *ProduceResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:137
	_go_fuzz_dep_.CoverTab[105831]++
												err := pe.putArrayLength(len(r.Blocks))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:139
		_go_fuzz_dep_.CoverTab[105835]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:140
		// _ = "end of CoverTab[105835]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:141
		_go_fuzz_dep_.CoverTab[105836]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:141
		// _ = "end of CoverTab[105836]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:141
	// _ = "end of CoverTab[105831]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:141
	_go_fuzz_dep_.CoverTab[105832]++
												for topic, partitions := range r.Blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:142
		_go_fuzz_dep_.CoverTab[105837]++
													err = pe.putString(topic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:144
			_go_fuzz_dep_.CoverTab[105840]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:145
			// _ = "end of CoverTab[105840]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:146
			_go_fuzz_dep_.CoverTab[105841]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:146
			// _ = "end of CoverTab[105841]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:146
		// _ = "end of CoverTab[105837]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:146
		_go_fuzz_dep_.CoverTab[105838]++
													err = pe.putArrayLength(len(partitions))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:148
			_go_fuzz_dep_.CoverTab[105842]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:149
			// _ = "end of CoverTab[105842]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:150
			_go_fuzz_dep_.CoverTab[105843]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:150
			// _ = "end of CoverTab[105843]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:150
		// _ = "end of CoverTab[105838]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:150
		_go_fuzz_dep_.CoverTab[105839]++
													for id, prb := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:151
			_go_fuzz_dep_.CoverTab[105844]++
														pe.putInt32(id)
														err = prb.encode(pe, r.Version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:154
				_go_fuzz_dep_.CoverTab[105845]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:155
				// _ = "end of CoverTab[105845]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:156
				_go_fuzz_dep_.CoverTab[105846]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:156
				// _ = "end of CoverTab[105846]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:156
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:156
			// _ = "end of CoverTab[105844]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:157
		// _ = "end of CoverTab[105839]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:158
	// _ = "end of CoverTab[105832]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:158
	_go_fuzz_dep_.CoverTab[105833]++

												if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:160
		_go_fuzz_dep_.CoverTab[105847]++
													pe.putInt32(int32(r.ThrottleTime / time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:161
		// _ = "end of CoverTab[105847]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:162
		_go_fuzz_dep_.CoverTab[105848]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:162
		// _ = "end of CoverTab[105848]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:162
	// _ = "end of CoverTab[105833]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:162
	_go_fuzz_dep_.CoverTab[105834]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:163
	// _ = "end of CoverTab[105834]"
}

func (r *ProduceResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:166
	_go_fuzz_dep_.CoverTab[105849]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:167
	// _ = "end of CoverTab[105849]"
}

func (r *ProduceResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:170
	_go_fuzz_dep_.CoverTab[105850]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:171
	// _ = "end of CoverTab[105850]"
}

func (r *ProduceResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:174
	_go_fuzz_dep_.CoverTab[105851]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:175
	// _ = "end of CoverTab[105851]"
}

func (r *ProduceResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:178
	_go_fuzz_dep_.CoverTab[105852]++
												return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:179
	// _ = "end of CoverTab[105852]"
}

func (r *ProduceResponse) GetBlock(topic string, partition int32) *ProduceResponseBlock {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:182
	_go_fuzz_dep_.CoverTab[105853]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:183
		_go_fuzz_dep_.CoverTab[105856]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:184
		// _ = "end of CoverTab[105856]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:185
		_go_fuzz_dep_.CoverTab[105857]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:185
		// _ = "end of CoverTab[105857]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:185
	// _ = "end of CoverTab[105853]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:185
	_go_fuzz_dep_.CoverTab[105854]++

												if r.Blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:187
		_go_fuzz_dep_.CoverTab[105858]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:188
		// _ = "end of CoverTab[105858]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:189
		_go_fuzz_dep_.CoverTab[105859]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:189
		// _ = "end of CoverTab[105859]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:189
	// _ = "end of CoverTab[105854]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:189
	_go_fuzz_dep_.CoverTab[105855]++

												return r.Blocks[topic][partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:191
	// _ = "end of CoverTab[105855]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:196
func (r *ProduceResponse) AddTopicPartition(topic string, partition int32, err KError) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:196
	_go_fuzz_dep_.CoverTab[105860]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:197
		_go_fuzz_dep_.CoverTab[105864]++
													r.Blocks = make(map[string]map[int32]*ProduceResponseBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:198
		// _ = "end of CoverTab[105864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:199
		_go_fuzz_dep_.CoverTab[105865]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:199
		// _ = "end of CoverTab[105865]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:199
	// _ = "end of CoverTab[105860]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:199
	_go_fuzz_dep_.CoverTab[105861]++
												byTopic, ok := r.Blocks[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:201
		_go_fuzz_dep_.CoverTab[105866]++
													byTopic = make(map[int32]*ProduceResponseBlock)
													r.Blocks[topic] = byTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:203
		// _ = "end of CoverTab[105866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:204
		_go_fuzz_dep_.CoverTab[105867]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:204
		// _ = "end of CoverTab[105867]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:204
	// _ = "end of CoverTab[105861]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:204
	_go_fuzz_dep_.CoverTab[105862]++
												block := &ProduceResponseBlock{
		Err: err,
	}
	if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:208
		_go_fuzz_dep_.CoverTab[105868]++
													block.Timestamp = time.Now()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:209
		// _ = "end of CoverTab[105868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:210
		_go_fuzz_dep_.CoverTab[105869]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:210
		// _ = "end of CoverTab[105869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:210
	// _ = "end of CoverTab[105862]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:210
	_go_fuzz_dep_.CoverTab[105863]++
												byTopic[partition] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:211
	// _ = "end of CoverTab[105863]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:212
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_response.go:212
var _ = _go_fuzz_dep_.CoverTab
