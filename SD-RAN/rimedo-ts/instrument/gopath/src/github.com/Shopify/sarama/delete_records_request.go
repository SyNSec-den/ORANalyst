//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:1
)

import (
	"sort"
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:15
type DeleteRecordsRequest struct {
	Topics	map[string]*DeleteRecordsRequestTopic
	Timeout	time.Duration
}

func (d *DeleteRecordsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:20
	_go_fuzz_dep_.CoverTab[101829]++
													if err := pe.putArrayLength(len(d.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:21
		_go_fuzz_dep_.CoverTab[101833]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:22
		// _ = "end of CoverTab[101833]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:23
		_go_fuzz_dep_.CoverTab[101834]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:23
		// _ = "end of CoverTab[101834]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:23
	// _ = "end of CoverTab[101829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:23
	_go_fuzz_dep_.CoverTab[101830]++
													keys := make([]string, 0, len(d.Topics))
													for topic := range d.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:25
		_go_fuzz_dep_.CoverTab[101835]++
														keys = append(keys, topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:26
		// _ = "end of CoverTab[101835]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:27
	// _ = "end of CoverTab[101830]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:27
	_go_fuzz_dep_.CoverTab[101831]++
													sort.Strings(keys)
													for _, topic := range keys {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:29
		_go_fuzz_dep_.CoverTab[101836]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:30
			_go_fuzz_dep_.CoverTab[101838]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:31
			// _ = "end of CoverTab[101838]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:32
			_go_fuzz_dep_.CoverTab[101839]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:32
			// _ = "end of CoverTab[101839]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:32
		// _ = "end of CoverTab[101836]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:32
		_go_fuzz_dep_.CoverTab[101837]++
														if err := d.Topics[topic].encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:33
			_go_fuzz_dep_.CoverTab[101840]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:34
			// _ = "end of CoverTab[101840]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:35
			_go_fuzz_dep_.CoverTab[101841]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:35
			// _ = "end of CoverTab[101841]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:35
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:35
		// _ = "end of CoverTab[101837]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:36
	// _ = "end of CoverTab[101831]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:36
	_go_fuzz_dep_.CoverTab[101832]++
													pe.putInt32(int32(d.Timeout / time.Millisecond))

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:39
	// _ = "end of CoverTab[101832]"
}

func (d *DeleteRecordsRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:42
	_go_fuzz_dep_.CoverTab[101842]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:44
		_go_fuzz_dep_.CoverTab[101846]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:45
		// _ = "end of CoverTab[101846]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:46
		_go_fuzz_dep_.CoverTab[101847]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:46
		// _ = "end of CoverTab[101847]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:46
	// _ = "end of CoverTab[101842]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:46
	_go_fuzz_dep_.CoverTab[101843]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:48
		_go_fuzz_dep_.CoverTab[101848]++
														d.Topics = make(map[string]*DeleteRecordsRequestTopic, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:50
			_go_fuzz_dep_.CoverTab[101849]++
															topic, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:52
				_go_fuzz_dep_.CoverTab[101852]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:53
				// _ = "end of CoverTab[101852]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:54
				_go_fuzz_dep_.CoverTab[101853]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:54
				// _ = "end of CoverTab[101853]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:54
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:54
			// _ = "end of CoverTab[101849]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:54
			_go_fuzz_dep_.CoverTab[101850]++
															details := new(DeleteRecordsRequestTopic)
															if err = details.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:56
				_go_fuzz_dep_.CoverTab[101854]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:57
				// _ = "end of CoverTab[101854]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:58
				_go_fuzz_dep_.CoverTab[101855]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:58
				// _ = "end of CoverTab[101855]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:58
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:58
			// _ = "end of CoverTab[101850]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:58
			_go_fuzz_dep_.CoverTab[101851]++
															d.Topics[topic] = details
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:59
			// _ = "end of CoverTab[101851]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:60
		// _ = "end of CoverTab[101848]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:61
		_go_fuzz_dep_.CoverTab[101856]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:61
		// _ = "end of CoverTab[101856]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:61
	// _ = "end of CoverTab[101843]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:61
	_go_fuzz_dep_.CoverTab[101844]++

													timeout, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:64
		_go_fuzz_dep_.CoverTab[101857]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:65
		// _ = "end of CoverTab[101857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:66
		_go_fuzz_dep_.CoverTab[101858]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:66
		// _ = "end of CoverTab[101858]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:66
	// _ = "end of CoverTab[101844]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:66
	_go_fuzz_dep_.CoverTab[101845]++
													d.Timeout = time.Duration(timeout) * time.Millisecond

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:69
	// _ = "end of CoverTab[101845]"
}

func (d *DeleteRecordsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:72
	_go_fuzz_dep_.CoverTab[101859]++
													return 21
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:73
	// _ = "end of CoverTab[101859]"
}

func (d *DeleteRecordsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:76
	_go_fuzz_dep_.CoverTab[101860]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:77
	// _ = "end of CoverTab[101860]"
}

func (d *DeleteRecordsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:80
	_go_fuzz_dep_.CoverTab[101861]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:81
	// _ = "end of CoverTab[101861]"
}

func (d *DeleteRecordsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:84
	_go_fuzz_dep_.CoverTab[101862]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:85
	// _ = "end of CoverTab[101862]"
}

type DeleteRecordsRequestTopic struct {
	PartitionOffsets map[int32]int64	// partition => offset
}

func (t *DeleteRecordsRequestTopic) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:92
	_go_fuzz_dep_.CoverTab[101863]++
													if err := pe.putArrayLength(len(t.PartitionOffsets)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:93
		_go_fuzz_dep_.CoverTab[101868]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:94
		// _ = "end of CoverTab[101868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:95
		_go_fuzz_dep_.CoverTab[101869]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:95
		// _ = "end of CoverTab[101869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:95
	// _ = "end of CoverTab[101863]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:95
	_go_fuzz_dep_.CoverTab[101864]++
													keys := make([]int32, 0, len(t.PartitionOffsets))
													for partition := range t.PartitionOffsets {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:97
		_go_fuzz_dep_.CoverTab[101870]++
														keys = append(keys, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:98
		// _ = "end of CoverTab[101870]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:99
	// _ = "end of CoverTab[101864]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:99
	_go_fuzz_dep_.CoverTab[101865]++
													sort.Slice(keys, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:100
		_go_fuzz_dep_.CoverTab[101871]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:100
		return keys[i] < keys[j]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:100
		// _ = "end of CoverTab[101871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:100
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:100
	// _ = "end of CoverTab[101865]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:100
	_go_fuzz_dep_.CoverTab[101866]++
													for _, partition := range keys {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:101
		_go_fuzz_dep_.CoverTab[101872]++
														pe.putInt32(partition)
														pe.putInt64(t.PartitionOffsets[partition])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:103
		// _ = "end of CoverTab[101872]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:104
	// _ = "end of CoverTab[101866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:104
	_go_fuzz_dep_.CoverTab[101867]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:105
	// _ = "end of CoverTab[101867]"
}

func (t *DeleteRecordsRequestTopic) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:108
	_go_fuzz_dep_.CoverTab[101873]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:110
		_go_fuzz_dep_.CoverTab[101876]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:111
		// _ = "end of CoverTab[101876]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:112
		_go_fuzz_dep_.CoverTab[101877]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:112
		// _ = "end of CoverTab[101877]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:112
	// _ = "end of CoverTab[101873]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:112
	_go_fuzz_dep_.CoverTab[101874]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:114
		_go_fuzz_dep_.CoverTab[101878]++
														t.PartitionOffsets = make(map[int32]int64, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:116
			_go_fuzz_dep_.CoverTab[101879]++
															partition, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:118
				_go_fuzz_dep_.CoverTab[101882]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:119
				// _ = "end of CoverTab[101882]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:120
				_go_fuzz_dep_.CoverTab[101883]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:120
				// _ = "end of CoverTab[101883]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:120
			// _ = "end of CoverTab[101879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:120
			_go_fuzz_dep_.CoverTab[101880]++
															offset, err := pd.getInt64()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:122
				_go_fuzz_dep_.CoverTab[101884]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:123
				// _ = "end of CoverTab[101884]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:124
				_go_fuzz_dep_.CoverTab[101885]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:124
				// _ = "end of CoverTab[101885]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:124
			// _ = "end of CoverTab[101880]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:124
			_go_fuzz_dep_.CoverTab[101881]++
															t.PartitionOffsets[partition] = offset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:125
			// _ = "end of CoverTab[101881]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:126
		// _ = "end of CoverTab[101878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:127
		_go_fuzz_dep_.CoverTab[101886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:127
		// _ = "end of CoverTab[101886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:127
	// _ = "end of CoverTab[101874]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:127
	_go_fuzz_dep_.CoverTab[101875]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:129
	// _ = "end of CoverTab[101875]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_records_request.go:130
var _ = _go_fuzz_dep_.CoverTab
