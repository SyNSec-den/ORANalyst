//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:1
)

type DeleteOffsetsRequest struct {
	Group		string
	partitions	map[string][]int32
}

func (r *DeleteOffsetsRequest) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:8
	_go_fuzz_dep_.CoverTab[101735]++
													err = pe.putString(r.Group)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:10
		_go_fuzz_dep_.CoverTab[101739]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:11
		// _ = "end of CoverTab[101739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:12
		_go_fuzz_dep_.CoverTab[101740]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:12
		// _ = "end of CoverTab[101740]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:12
	// _ = "end of CoverTab[101735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:12
	_go_fuzz_dep_.CoverTab[101736]++

													if r.partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:14
		_go_fuzz_dep_.CoverTab[101741]++
														pe.putInt32(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:15
		// _ = "end of CoverTab[101741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:16
		_go_fuzz_dep_.CoverTab[101742]++
														if err = pe.putArrayLength(len(r.partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:17
			_go_fuzz_dep_.CoverTab[101743]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:18
			// _ = "end of CoverTab[101743]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:19
			_go_fuzz_dep_.CoverTab[101744]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:19
			// _ = "end of CoverTab[101744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:19
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:19
		// _ = "end of CoverTab[101742]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:20
	// _ = "end of CoverTab[101736]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:20
	_go_fuzz_dep_.CoverTab[101737]++
													for topic, partitions := range r.partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:21
		_go_fuzz_dep_.CoverTab[101745]++
														err = pe.putString(topic)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:23
			_go_fuzz_dep_.CoverTab[101747]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:24
			// _ = "end of CoverTab[101747]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:25
			_go_fuzz_dep_.CoverTab[101748]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:25
			// _ = "end of CoverTab[101748]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:25
		// _ = "end of CoverTab[101745]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:25
		_go_fuzz_dep_.CoverTab[101746]++
														err = pe.putInt32Array(partitions)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:27
			_go_fuzz_dep_.CoverTab[101749]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:28
			// _ = "end of CoverTab[101749]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:29
			_go_fuzz_dep_.CoverTab[101750]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:29
			// _ = "end of CoverTab[101750]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:29
		// _ = "end of CoverTab[101746]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:30
	// _ = "end of CoverTab[101737]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:30
	_go_fuzz_dep_.CoverTab[101738]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:31
	// _ = "end of CoverTab[101738]"
}

func (r *DeleteOffsetsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:34
	_go_fuzz_dep_.CoverTab[101751]++
													r.Group, err = pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:36
		_go_fuzz_dep_.CoverTab[101756]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:37
		// _ = "end of CoverTab[101756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:38
		_go_fuzz_dep_.CoverTab[101757]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:38
		// _ = "end of CoverTab[101757]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:38
	// _ = "end of CoverTab[101751]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:38
	_go_fuzz_dep_.CoverTab[101752]++
													var partitionCount int

													partitionCount, err = pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:42
		_go_fuzz_dep_.CoverTab[101758]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:43
		// _ = "end of CoverTab[101758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:44
		_go_fuzz_dep_.CoverTab[101759]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:44
		// _ = "end of CoverTab[101759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:44
	// _ = "end of CoverTab[101752]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:44
	_go_fuzz_dep_.CoverTab[101753]++

													if (partitionCount == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		_go_fuzz_dep_.CoverTab[101760]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		return version < 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		// _ = "end of CoverTab[101760]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		_go_fuzz_dep_.CoverTab[101761]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		return partitionCount < 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		// _ = "end of CoverTab[101761]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:46
		_go_fuzz_dep_.CoverTab[101762]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:47
		// _ = "end of CoverTab[101762]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:48
		_go_fuzz_dep_.CoverTab[101763]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:48
		// _ = "end of CoverTab[101763]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:48
	// _ = "end of CoverTab[101753]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:48
	_go_fuzz_dep_.CoverTab[101754]++

													r.partitions = make(map[string][]int32, partitionCount)
													for i := 0; i < partitionCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:51
		_go_fuzz_dep_.CoverTab[101764]++
														var topic string
														topic, err = pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:54
			_go_fuzz_dep_.CoverTab[101767]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:55
			// _ = "end of CoverTab[101767]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:56
			_go_fuzz_dep_.CoverTab[101768]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:56
			// _ = "end of CoverTab[101768]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:56
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:56
		// _ = "end of CoverTab[101764]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:56
		_go_fuzz_dep_.CoverTab[101765]++

														var partitions []int32
														partitions, err = pd.getInt32Array()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:60
			_go_fuzz_dep_.CoverTab[101769]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:61
			// _ = "end of CoverTab[101769]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:62
			_go_fuzz_dep_.CoverTab[101770]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:62
			// _ = "end of CoverTab[101770]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:62
		// _ = "end of CoverTab[101765]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:62
		_go_fuzz_dep_.CoverTab[101766]++

														r.partitions[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:64
		// _ = "end of CoverTab[101766]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:65
	// _ = "end of CoverTab[101754]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:65
	_go_fuzz_dep_.CoverTab[101755]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:67
	// _ = "end of CoverTab[101755]"
}

func (r *DeleteOffsetsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:70
	_go_fuzz_dep_.CoverTab[101771]++
													return 47
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:71
	// _ = "end of CoverTab[101771]"
}

func (r *DeleteOffsetsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:74
	_go_fuzz_dep_.CoverTab[101772]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:75
	// _ = "end of CoverTab[101772]"
}

func (r *DeleteOffsetsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:78
	_go_fuzz_dep_.CoverTab[101773]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:79
	// _ = "end of CoverTab[101773]"
}

func (r *DeleteOffsetsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:82
	_go_fuzz_dep_.CoverTab[101774]++
													return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:83
	// _ = "end of CoverTab[101774]"
}

func (r *DeleteOffsetsRequest) AddPartition(topic string, partitionID int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:86
	_go_fuzz_dep_.CoverTab[101775]++
													if r.partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:87
		_go_fuzz_dep_.CoverTab[101777]++
														r.partitions = make(map[string][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:88
		// _ = "end of CoverTab[101777]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:89
		_go_fuzz_dep_.CoverTab[101778]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:89
		// _ = "end of CoverTab[101778]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:89
	// _ = "end of CoverTab[101775]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:89
	_go_fuzz_dep_.CoverTab[101776]++

													r.partitions[topic] = append(r.partitions[topic], partitionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:91
	// _ = "end of CoverTab[101776]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:92
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_offsets_request.go:92
var _ = _go_fuzz_dep_.CoverTab
