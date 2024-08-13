//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:1
)

import (
	"time"
)

type DeleteOffsetsResponse struct {
	//The top-level error code, or 0 if there was no error.
	ErrorCode	KError
	ThrottleTime	time.Duration
	//The responses for each partition of the topics.
	Errors	map[string]map[int32]KError
}

func (r *DeleteOffsetsResponse) AddError(topic string, partition int32, errorCode KError) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:15
	_go_fuzz_dep_.CoverTab[101779]++
													if r.Errors == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:16
		_go_fuzz_dep_.CoverTab[101782]++
														r.Errors = make(map[string]map[int32]KError)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:17
		// _ = "end of CoverTab[101782]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:18
		_go_fuzz_dep_.CoverTab[101783]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:18
		// _ = "end of CoverTab[101783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:18
	// _ = "end of CoverTab[101779]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:18
	_go_fuzz_dep_.CoverTab[101780]++
													partitions := r.Errors[topic]
													if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:20
		_go_fuzz_dep_.CoverTab[101784]++
														partitions = make(map[int32]KError)
														r.Errors[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:22
		// _ = "end of CoverTab[101784]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:23
		_go_fuzz_dep_.CoverTab[101785]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:23
		// _ = "end of CoverTab[101785]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:23
	// _ = "end of CoverTab[101780]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:23
	_go_fuzz_dep_.CoverTab[101781]++
													partitions[partition] = errorCode
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:24
	// _ = "end of CoverTab[101781]"
}

func (r *DeleteOffsetsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:27
	_go_fuzz_dep_.CoverTab[101786]++
													pe.putInt16(int16(r.ErrorCode))
													pe.putInt32(int32(r.ThrottleTime / time.Millisecond))

													if err := pe.putArrayLength(len(r.Errors)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:31
		_go_fuzz_dep_.CoverTab[101789]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:32
		// _ = "end of CoverTab[101789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:33
		_go_fuzz_dep_.CoverTab[101790]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:33
		// _ = "end of CoverTab[101790]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:33
	// _ = "end of CoverTab[101786]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:33
	_go_fuzz_dep_.CoverTab[101787]++
													for topic, partitions := range r.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:34
		_go_fuzz_dep_.CoverTab[101791]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:35
			_go_fuzz_dep_.CoverTab[101794]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:36
			// _ = "end of CoverTab[101794]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:37
			_go_fuzz_dep_.CoverTab[101795]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:37
			// _ = "end of CoverTab[101795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:37
		// _ = "end of CoverTab[101791]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:37
		_go_fuzz_dep_.CoverTab[101792]++
														if err := pe.putArrayLength(len(partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:38
			_go_fuzz_dep_.CoverTab[101796]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:39
			// _ = "end of CoverTab[101796]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:40
			_go_fuzz_dep_.CoverTab[101797]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:40
			// _ = "end of CoverTab[101797]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:40
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:40
		// _ = "end of CoverTab[101792]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:40
		_go_fuzz_dep_.CoverTab[101793]++
														for partition, errorCode := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:41
			_go_fuzz_dep_.CoverTab[101798]++
															pe.putInt32(partition)
															pe.putInt16(int16(errorCode))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:43
			// _ = "end of CoverTab[101798]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:44
		// _ = "end of CoverTab[101793]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:45
	// _ = "end of CoverTab[101787]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:45
	_go_fuzz_dep_.CoverTab[101788]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:46
	// _ = "end of CoverTab[101788]"
}

func (r *DeleteOffsetsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:49
	_go_fuzz_dep_.CoverTab[101799]++
													tmpErr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:51
		_go_fuzz_dep_.CoverTab[101804]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:52
		// _ = "end of CoverTab[101804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:53
		_go_fuzz_dep_.CoverTab[101805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:53
		// _ = "end of CoverTab[101805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:53
	// _ = "end of CoverTab[101799]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:53
	_go_fuzz_dep_.CoverTab[101800]++
													r.ErrorCode = KError(tmpErr)

													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:57
		_go_fuzz_dep_.CoverTab[101806]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:58
		// _ = "end of CoverTab[101806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:59
		_go_fuzz_dep_.CoverTab[101807]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:59
		// _ = "end of CoverTab[101807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:59
	// _ = "end of CoverTab[101800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:59
	_go_fuzz_dep_.CoverTab[101801]++
													r.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													numTopics, err := pd.getArrayLength()
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:63
		_go_fuzz_dep_.CoverTab[101808]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:63
		return numTopics == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:63
		// _ = "end of CoverTab[101808]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:63
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:63
		_go_fuzz_dep_.CoverTab[101809]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:64
		// _ = "end of CoverTab[101809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:65
		_go_fuzz_dep_.CoverTab[101810]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:65
		// _ = "end of CoverTab[101810]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:65
	// _ = "end of CoverTab[101801]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:65
	_go_fuzz_dep_.CoverTab[101802]++

													r.Errors = make(map[string]map[int32]KError, numTopics)
													for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:68
		_go_fuzz_dep_.CoverTab[101811]++
														name, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:70
			_go_fuzz_dep_.CoverTab[101814]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:71
			// _ = "end of CoverTab[101814]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:72
			_go_fuzz_dep_.CoverTab[101815]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:72
			// _ = "end of CoverTab[101815]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:72
		// _ = "end of CoverTab[101811]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:72
		_go_fuzz_dep_.CoverTab[101812]++

														numErrors, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:75
			_go_fuzz_dep_.CoverTab[101816]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:76
			// _ = "end of CoverTab[101816]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:77
			_go_fuzz_dep_.CoverTab[101817]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:77
			// _ = "end of CoverTab[101817]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:77
		// _ = "end of CoverTab[101812]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:77
		_go_fuzz_dep_.CoverTab[101813]++

														r.Errors[name] = make(map[int32]KError, numErrors)

														for j := 0; j < numErrors; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:81
			_go_fuzz_dep_.CoverTab[101818]++
															id, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:83
				_go_fuzz_dep_.CoverTab[101821]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:84
				// _ = "end of CoverTab[101821]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:85
				_go_fuzz_dep_.CoverTab[101822]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:85
				// _ = "end of CoverTab[101822]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:85
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:85
			// _ = "end of CoverTab[101818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:85
			_go_fuzz_dep_.CoverTab[101819]++

															tmp, err := pd.getInt16()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:88
				_go_fuzz_dep_.CoverTab[101823]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:89
				// _ = "end of CoverTab[101823]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:90
				_go_fuzz_dep_.CoverTab[101824]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:90
				// _ = "end of CoverTab[101824]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:90
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:90
			// _ = "end of CoverTab[101819]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:90
			_go_fuzz_dep_.CoverTab[101820]++
															r.Errors[name][id] = KError(tmp)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:91
			// _ = "end of CoverTab[101820]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:92
		// _ = "end of CoverTab[101813]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:93
	// _ = "end of CoverTab[101802]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:93
	_go_fuzz_dep_.CoverTab[101803]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:95
	// _ = "end of CoverTab[101803]"
}

func (r *DeleteOffsetsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:98
	_go_fuzz_dep_.CoverTab[101825]++
													return 47
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:99
	// _ = "end of CoverTab[101825]"
}

func (r *DeleteOffsetsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:102
	_go_fuzz_dep_.CoverTab[101826]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:103
	// _ = "end of CoverTab[101826]"
}

func (r *DeleteOffsetsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:106
	_go_fuzz_dep_.CoverTab[101827]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:107
	// _ = "end of CoverTab[101827]"
}

func (r *DeleteOffsetsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:110
	_go_fuzz_dep_.CoverTab[101828]++
													return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:111
	// _ = "end of CoverTab[101828]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:112
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_response.go:112
var _ = _go_fuzz_dep_.CoverTab
