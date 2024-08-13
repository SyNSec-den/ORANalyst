//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:1
)

import (
	"sort"
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:15
type DeleteRecordsResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	Topics		map[string]*DeleteRecordsResponseTopic
}

func (d *DeleteRecordsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:21
	_go_fuzz_dep_.CoverTab[101887]++
													pe.putInt32(int32(d.ThrottleTime / time.Millisecond))

													if err := pe.putArrayLength(len(d.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:24
		_go_fuzz_dep_.CoverTab[101891]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:25
		// _ = "end of CoverTab[101891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:26
		_go_fuzz_dep_.CoverTab[101892]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:26
		// _ = "end of CoverTab[101892]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:26
	// _ = "end of CoverTab[101887]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:26
	_go_fuzz_dep_.CoverTab[101888]++
													keys := make([]string, 0, len(d.Topics))
													for topic := range d.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:28
		_go_fuzz_dep_.CoverTab[101893]++
														keys = append(keys, topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:29
		// _ = "end of CoverTab[101893]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:30
	// _ = "end of CoverTab[101888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:30
	_go_fuzz_dep_.CoverTab[101889]++
													sort.Strings(keys)
													for _, topic := range keys {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:32
		_go_fuzz_dep_.CoverTab[101894]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:33
			_go_fuzz_dep_.CoverTab[101896]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:34
			// _ = "end of CoverTab[101896]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:35
			_go_fuzz_dep_.CoverTab[101897]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:35
			// _ = "end of CoverTab[101897]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:35
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:35
		// _ = "end of CoverTab[101894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:35
		_go_fuzz_dep_.CoverTab[101895]++
														if err := d.Topics[topic].encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:36
			_go_fuzz_dep_.CoverTab[101898]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:37
			// _ = "end of CoverTab[101898]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:38
			_go_fuzz_dep_.CoverTab[101899]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:38
			// _ = "end of CoverTab[101899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:38
		// _ = "end of CoverTab[101895]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:39
	// _ = "end of CoverTab[101889]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:39
	_go_fuzz_dep_.CoverTab[101890]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:40
	// _ = "end of CoverTab[101890]"
}

func (d *DeleteRecordsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:43
	_go_fuzz_dep_.CoverTab[101900]++
													d.Version = version

													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:47
		_go_fuzz_dep_.CoverTab[101904]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:48
		// _ = "end of CoverTab[101904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:49
		_go_fuzz_dep_.CoverTab[101905]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:49
		// _ = "end of CoverTab[101905]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:49
	// _ = "end of CoverTab[101900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:49
	_go_fuzz_dep_.CoverTab[101901]++
													d.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:53
		_go_fuzz_dep_.CoverTab[101906]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:54
		// _ = "end of CoverTab[101906]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:55
		_go_fuzz_dep_.CoverTab[101907]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:55
		// _ = "end of CoverTab[101907]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:55
	// _ = "end of CoverTab[101901]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:55
	_go_fuzz_dep_.CoverTab[101902]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:57
		_go_fuzz_dep_.CoverTab[101908]++
														d.Topics = make(map[string]*DeleteRecordsResponseTopic, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:59
			_go_fuzz_dep_.CoverTab[101909]++
															topic, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:61
				_go_fuzz_dep_.CoverTab[101912]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:62
				// _ = "end of CoverTab[101912]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:63
				_go_fuzz_dep_.CoverTab[101913]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:63
				// _ = "end of CoverTab[101913]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:63
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:63
			// _ = "end of CoverTab[101909]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:63
			_go_fuzz_dep_.CoverTab[101910]++
															details := new(DeleteRecordsResponseTopic)
															if err = details.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:65
				_go_fuzz_dep_.CoverTab[101914]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:66
				// _ = "end of CoverTab[101914]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:67
				_go_fuzz_dep_.CoverTab[101915]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:67
				// _ = "end of CoverTab[101915]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:67
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:67
			// _ = "end of CoverTab[101910]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:67
			_go_fuzz_dep_.CoverTab[101911]++
															d.Topics[topic] = details
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:68
			// _ = "end of CoverTab[101911]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:69
		// _ = "end of CoverTab[101908]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:70
		_go_fuzz_dep_.CoverTab[101916]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:70
		// _ = "end of CoverTab[101916]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:70
	// _ = "end of CoverTab[101902]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:70
	_go_fuzz_dep_.CoverTab[101903]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:72
	// _ = "end of CoverTab[101903]"
}

func (d *DeleteRecordsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:75
	_go_fuzz_dep_.CoverTab[101917]++
													return 21
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:76
	// _ = "end of CoverTab[101917]"
}

func (d *DeleteRecordsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:79
	_go_fuzz_dep_.CoverTab[101918]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:80
	// _ = "end of CoverTab[101918]"
}

func (d *DeleteRecordsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:83
	_go_fuzz_dep_.CoverTab[101919]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:84
	// _ = "end of CoverTab[101919]"
}

func (d *DeleteRecordsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:87
	_go_fuzz_dep_.CoverTab[101920]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:88
	// _ = "end of CoverTab[101920]"
}

type DeleteRecordsResponseTopic struct {
	Partitions map[int32]*DeleteRecordsResponsePartition
}

func (t *DeleteRecordsResponseTopic) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:95
	_go_fuzz_dep_.CoverTab[101921]++
													if err := pe.putArrayLength(len(t.Partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:96
		_go_fuzz_dep_.CoverTab[101926]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:97
		// _ = "end of CoverTab[101926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:98
		_go_fuzz_dep_.CoverTab[101927]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:98
		// _ = "end of CoverTab[101927]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:98
	// _ = "end of CoverTab[101921]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:98
	_go_fuzz_dep_.CoverTab[101922]++
													keys := make([]int32, 0, len(t.Partitions))
													for partition := range t.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:100
		_go_fuzz_dep_.CoverTab[101928]++
														keys = append(keys, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:101
		// _ = "end of CoverTab[101928]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:102
	// _ = "end of CoverTab[101922]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:102
	_go_fuzz_dep_.CoverTab[101923]++
													sort.Slice(keys, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:103
		_go_fuzz_dep_.CoverTab[101929]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:103
		return keys[i] < keys[j]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:103
		// _ = "end of CoverTab[101929]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:103
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:103
	// _ = "end of CoverTab[101923]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:103
	_go_fuzz_dep_.CoverTab[101924]++
													for _, partition := range keys {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:104
		_go_fuzz_dep_.CoverTab[101930]++
														pe.putInt32(partition)
														if err := t.Partitions[partition].encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:106
			_go_fuzz_dep_.CoverTab[101931]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:107
			// _ = "end of CoverTab[101931]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:108
			_go_fuzz_dep_.CoverTab[101932]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:108
			// _ = "end of CoverTab[101932]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:108
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:108
		// _ = "end of CoverTab[101930]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:109
	// _ = "end of CoverTab[101924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:109
	_go_fuzz_dep_.CoverTab[101925]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:110
	// _ = "end of CoverTab[101925]"
}

func (t *DeleteRecordsResponseTopic) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:113
	_go_fuzz_dep_.CoverTab[101933]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:115
		_go_fuzz_dep_.CoverTab[101936]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:116
		// _ = "end of CoverTab[101936]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:117
		_go_fuzz_dep_.CoverTab[101937]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:117
		// _ = "end of CoverTab[101937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:117
	// _ = "end of CoverTab[101933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:117
	_go_fuzz_dep_.CoverTab[101934]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:119
		_go_fuzz_dep_.CoverTab[101938]++
														t.Partitions = make(map[int32]*DeleteRecordsResponsePartition, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:121
			_go_fuzz_dep_.CoverTab[101939]++
															partition, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:123
				_go_fuzz_dep_.CoverTab[101942]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:124
				// _ = "end of CoverTab[101942]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:125
				_go_fuzz_dep_.CoverTab[101943]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:125
				// _ = "end of CoverTab[101943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:125
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:125
			// _ = "end of CoverTab[101939]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:125
			_go_fuzz_dep_.CoverTab[101940]++
															details := new(DeleteRecordsResponsePartition)
															if err = details.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:127
				_go_fuzz_dep_.CoverTab[101944]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:128
				// _ = "end of CoverTab[101944]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:129
				_go_fuzz_dep_.CoverTab[101945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:129
				// _ = "end of CoverTab[101945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:129
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:129
			// _ = "end of CoverTab[101940]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:129
			_go_fuzz_dep_.CoverTab[101941]++
															t.Partitions[partition] = details
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:130
			// _ = "end of CoverTab[101941]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:131
		// _ = "end of CoverTab[101938]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:132
		_go_fuzz_dep_.CoverTab[101946]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:132
		// _ = "end of CoverTab[101946]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:132
	// _ = "end of CoverTab[101934]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:132
	_go_fuzz_dep_.CoverTab[101935]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:134
	// _ = "end of CoverTab[101935]"
}

type DeleteRecordsResponsePartition struct {
	LowWatermark	int64
	Err		KError
}

func (t *DeleteRecordsResponsePartition) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:142
	_go_fuzz_dep_.CoverTab[101947]++
													pe.putInt64(t.LowWatermark)
													pe.putInt16(int16(t.Err))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:145
	// _ = "end of CoverTab[101947]"
}

func (t *DeleteRecordsResponsePartition) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:148
	_go_fuzz_dep_.CoverTab[101948]++
													lowWatermark, err := pd.getInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:150
		_go_fuzz_dep_.CoverTab[101951]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:151
		// _ = "end of CoverTab[101951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:152
		_go_fuzz_dep_.CoverTab[101952]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:152
		// _ = "end of CoverTab[101952]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:152
	// _ = "end of CoverTab[101948]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:152
	_go_fuzz_dep_.CoverTab[101949]++
													t.LowWatermark = lowWatermark

													kErr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:156
		_go_fuzz_dep_.CoverTab[101953]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:157
		// _ = "end of CoverTab[101953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:158
		_go_fuzz_dep_.CoverTab[101954]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:158
		// _ = "end of CoverTab[101954]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:158
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:158
	// _ = "end of CoverTab[101949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:158
	_go_fuzz_dep_.CoverTab[101950]++
													t.Err = KError(kErr)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:161
	// _ = "end of CoverTab[101950]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:162
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_response.go:162
var _ = _go_fuzz_dep_.CoverTab
