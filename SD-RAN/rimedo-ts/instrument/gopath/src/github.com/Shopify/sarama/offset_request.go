//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:1
)

type offsetRequestBlock struct {
	time		int64
	maxOffsets	int32	// Only used in version 0
}

func (b *offsetRequestBlock) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:8
	_go_fuzz_dep_.CoverTab[105376]++
												pe.putInt64(b.time)
												if version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:10
		_go_fuzz_dep_.CoverTab[105378]++
													pe.putInt32(b.maxOffsets)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:11
		// _ = "end of CoverTab[105378]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:12
		_go_fuzz_dep_.CoverTab[105379]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:12
		// _ = "end of CoverTab[105379]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:12
	// _ = "end of CoverTab[105376]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:12
	_go_fuzz_dep_.CoverTab[105377]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:14
	// _ = "end of CoverTab[105377]"
}

func (b *offsetRequestBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:17
	_go_fuzz_dep_.CoverTab[105380]++
												if b.time, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:18
		_go_fuzz_dep_.CoverTab[105383]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:19
		// _ = "end of CoverTab[105383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:20
		_go_fuzz_dep_.CoverTab[105384]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:20
		// _ = "end of CoverTab[105384]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:20
	// _ = "end of CoverTab[105380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:20
	_go_fuzz_dep_.CoverTab[105381]++
												if version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:21
		_go_fuzz_dep_.CoverTab[105385]++
													if b.maxOffsets, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:22
			_go_fuzz_dep_.CoverTab[105386]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:23
			// _ = "end of CoverTab[105386]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:24
			_go_fuzz_dep_.CoverTab[105387]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:24
			// _ = "end of CoverTab[105387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:24
		// _ = "end of CoverTab[105385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:25
		_go_fuzz_dep_.CoverTab[105388]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:25
		// _ = "end of CoverTab[105388]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:25
	// _ = "end of CoverTab[105381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:25
	_go_fuzz_dep_.CoverTab[105382]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:26
	// _ = "end of CoverTab[105382]"
}

type OffsetRequest struct {
	Version		int16
	IsolationLevel	IsolationLevel
	replicaID	int32
	isReplicaIDSet	bool
	blocks		map[string]map[int32]*offsetRequestBlock
}

func (r *OffsetRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:37
	_go_fuzz_dep_.CoverTab[105389]++
												if r.isReplicaIDSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:38
		_go_fuzz_dep_.CoverTab[105394]++
													pe.putInt32(r.replicaID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:39
		// _ = "end of CoverTab[105394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:40
		_go_fuzz_dep_.CoverTab[105395]++

													pe.putInt32(-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:42
		// _ = "end of CoverTab[105395]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:43
	// _ = "end of CoverTab[105389]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:43
	_go_fuzz_dep_.CoverTab[105390]++

												if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:45
		_go_fuzz_dep_.CoverTab[105396]++
													pe.putBool(r.IsolationLevel == ReadCommitted)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:46
		// _ = "end of CoverTab[105396]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:47
		_go_fuzz_dep_.CoverTab[105397]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:47
		// _ = "end of CoverTab[105397]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:47
	// _ = "end of CoverTab[105390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:47
	_go_fuzz_dep_.CoverTab[105391]++

												err := pe.putArrayLength(len(r.blocks))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:50
		_go_fuzz_dep_.CoverTab[105398]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:51
		// _ = "end of CoverTab[105398]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:52
		_go_fuzz_dep_.CoverTab[105399]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:52
		// _ = "end of CoverTab[105399]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:52
	// _ = "end of CoverTab[105391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:52
	_go_fuzz_dep_.CoverTab[105392]++
												for topic, partitions := range r.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:53
		_go_fuzz_dep_.CoverTab[105400]++
													err = pe.putString(topic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:55
			_go_fuzz_dep_.CoverTab[105403]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:56
			// _ = "end of CoverTab[105403]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:57
			_go_fuzz_dep_.CoverTab[105404]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:57
			// _ = "end of CoverTab[105404]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:57
		// _ = "end of CoverTab[105400]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:57
		_go_fuzz_dep_.CoverTab[105401]++
													err = pe.putArrayLength(len(partitions))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:59
			_go_fuzz_dep_.CoverTab[105405]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:60
			// _ = "end of CoverTab[105405]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:61
			_go_fuzz_dep_.CoverTab[105406]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:61
			// _ = "end of CoverTab[105406]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:61
		// _ = "end of CoverTab[105401]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:61
		_go_fuzz_dep_.CoverTab[105402]++
													for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:62
			_go_fuzz_dep_.CoverTab[105407]++
														pe.putInt32(partition)
														if err = block.encode(pe, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:64
				_go_fuzz_dep_.CoverTab[105408]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:65
				// _ = "end of CoverTab[105408]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:66
				_go_fuzz_dep_.CoverTab[105409]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:66
				// _ = "end of CoverTab[105409]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:66
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:66
			// _ = "end of CoverTab[105407]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:67
		// _ = "end of CoverTab[105402]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:68
	// _ = "end of CoverTab[105392]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:68
	_go_fuzz_dep_.CoverTab[105393]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:69
	// _ = "end of CoverTab[105393]"
}

func (r *OffsetRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:72
	_go_fuzz_dep_.CoverTab[105410]++
												r.Version = version

												replicaID, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:76
		_go_fuzz_dep_.CoverTab[105417]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:77
		// _ = "end of CoverTab[105417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:78
		_go_fuzz_dep_.CoverTab[105418]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:78
		// _ = "end of CoverTab[105418]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:78
	// _ = "end of CoverTab[105410]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:78
	_go_fuzz_dep_.CoverTab[105411]++
												if replicaID >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:79
		_go_fuzz_dep_.CoverTab[105419]++
													r.SetReplicaID(replicaID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:80
		// _ = "end of CoverTab[105419]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:81
		_go_fuzz_dep_.CoverTab[105420]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:81
		// _ = "end of CoverTab[105420]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:81
	// _ = "end of CoverTab[105411]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:81
	_go_fuzz_dep_.CoverTab[105412]++

												if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:83
		_go_fuzz_dep_.CoverTab[105421]++
													tmp, err := pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:85
			_go_fuzz_dep_.CoverTab[105423]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:86
			// _ = "end of CoverTab[105423]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:87
			_go_fuzz_dep_.CoverTab[105424]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:87
			// _ = "end of CoverTab[105424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:87
		// _ = "end of CoverTab[105421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:87
		_go_fuzz_dep_.CoverTab[105422]++

													r.IsolationLevel = ReadUncommitted
													if tmp {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:90
			_go_fuzz_dep_.CoverTab[105425]++
														r.IsolationLevel = ReadCommitted
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:91
			// _ = "end of CoverTab[105425]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:92
			_go_fuzz_dep_.CoverTab[105426]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:92
			// _ = "end of CoverTab[105426]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:92
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:92
		// _ = "end of CoverTab[105422]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:93
		_go_fuzz_dep_.CoverTab[105427]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:93
		// _ = "end of CoverTab[105427]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:93
	// _ = "end of CoverTab[105412]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:93
	_go_fuzz_dep_.CoverTab[105413]++

												blockCount, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:96
		_go_fuzz_dep_.CoverTab[105428]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:97
		// _ = "end of CoverTab[105428]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:98
		_go_fuzz_dep_.CoverTab[105429]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:98
		// _ = "end of CoverTab[105429]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:98
	// _ = "end of CoverTab[105413]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:98
	_go_fuzz_dep_.CoverTab[105414]++
												if blockCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:99
		_go_fuzz_dep_.CoverTab[105430]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:100
		// _ = "end of CoverTab[105430]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:101
		_go_fuzz_dep_.CoverTab[105431]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:101
		// _ = "end of CoverTab[105431]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:101
	// _ = "end of CoverTab[105414]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:101
	_go_fuzz_dep_.CoverTab[105415]++
												r.blocks = make(map[string]map[int32]*offsetRequestBlock)
												for i := 0; i < blockCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:103
		_go_fuzz_dep_.CoverTab[105432]++
													topic, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:105
			_go_fuzz_dep_.CoverTab[105435]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:106
			// _ = "end of CoverTab[105435]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:107
			_go_fuzz_dep_.CoverTab[105436]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:107
			// _ = "end of CoverTab[105436]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:107
		// _ = "end of CoverTab[105432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:107
		_go_fuzz_dep_.CoverTab[105433]++
													partitionCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:109
			_go_fuzz_dep_.CoverTab[105437]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:110
			// _ = "end of CoverTab[105437]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:111
			_go_fuzz_dep_.CoverTab[105438]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:111
			// _ = "end of CoverTab[105438]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:111
		// _ = "end of CoverTab[105433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:111
		_go_fuzz_dep_.CoverTab[105434]++
													r.blocks[topic] = make(map[int32]*offsetRequestBlock)
													for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:113
			_go_fuzz_dep_.CoverTab[105439]++
														partition, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:115
				_go_fuzz_dep_.CoverTab[105442]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:116
				// _ = "end of CoverTab[105442]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:117
				_go_fuzz_dep_.CoverTab[105443]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:117
				// _ = "end of CoverTab[105443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:117
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:117
			// _ = "end of CoverTab[105439]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:117
			_go_fuzz_dep_.CoverTab[105440]++
														block := &offsetRequestBlock{}
														if err := block.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:119
				_go_fuzz_dep_.CoverTab[105444]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:120
				// _ = "end of CoverTab[105444]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:121
				_go_fuzz_dep_.CoverTab[105445]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:121
				// _ = "end of CoverTab[105445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:121
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:121
			// _ = "end of CoverTab[105440]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:121
			_go_fuzz_dep_.CoverTab[105441]++
														r.blocks[topic][partition] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:122
			// _ = "end of CoverTab[105441]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:123
		// _ = "end of CoverTab[105434]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:124
	// _ = "end of CoverTab[105415]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:124
	_go_fuzz_dep_.CoverTab[105416]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:125
	// _ = "end of CoverTab[105416]"
}

func (r *OffsetRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:128
	_go_fuzz_dep_.CoverTab[105446]++
												return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:129
	// _ = "end of CoverTab[105446]"
}

func (r *OffsetRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:132
	_go_fuzz_dep_.CoverTab[105447]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:133
	// _ = "end of CoverTab[105447]"
}

func (r *OffsetRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:136
	_go_fuzz_dep_.CoverTab[105448]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:137
	// _ = "end of CoverTab[105448]"
}

func (r *OffsetRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:140
	_go_fuzz_dep_.CoverTab[105449]++
												switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:142
		_go_fuzz_dep_.CoverTab[105450]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:143
		// _ = "end of CoverTab[105450]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:144
		_go_fuzz_dep_.CoverTab[105451]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:145
		// _ = "end of CoverTab[105451]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:146
		_go_fuzz_dep_.CoverTab[105452]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:147
		// _ = "end of CoverTab[105452]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:148
	// _ = "end of CoverTab[105449]"
}

func (r *OffsetRequest) SetReplicaID(id int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:151
	_go_fuzz_dep_.CoverTab[105453]++
												r.replicaID = id
												r.isReplicaIDSet = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:153
	// _ = "end of CoverTab[105453]"
}

func (r *OffsetRequest) ReplicaID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:156
	_go_fuzz_dep_.CoverTab[105454]++
												if r.isReplicaIDSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:157
		_go_fuzz_dep_.CoverTab[105456]++
													return r.replicaID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:158
		// _ = "end of CoverTab[105456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:159
		_go_fuzz_dep_.CoverTab[105457]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:159
		// _ = "end of CoverTab[105457]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:159
	// _ = "end of CoverTab[105454]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:159
	_go_fuzz_dep_.CoverTab[105455]++
												return -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:160
	// _ = "end of CoverTab[105455]"
}

func (r *OffsetRequest) AddBlock(topic string, partitionID int32, time int64, maxOffsets int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:163
	_go_fuzz_dep_.CoverTab[105458]++
												if r.blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:164
		_go_fuzz_dep_.CoverTab[105462]++
													r.blocks = make(map[string]map[int32]*offsetRequestBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:165
		// _ = "end of CoverTab[105462]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:166
		_go_fuzz_dep_.CoverTab[105463]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:166
		// _ = "end of CoverTab[105463]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:166
	// _ = "end of CoverTab[105458]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:166
	_go_fuzz_dep_.CoverTab[105459]++

												if r.blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:168
		_go_fuzz_dep_.CoverTab[105464]++
													r.blocks[topic] = make(map[int32]*offsetRequestBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:169
		// _ = "end of CoverTab[105464]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:170
		_go_fuzz_dep_.CoverTab[105465]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:170
		// _ = "end of CoverTab[105465]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:170
	// _ = "end of CoverTab[105459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:170
	_go_fuzz_dep_.CoverTab[105460]++

												tmp := new(offsetRequestBlock)
												tmp.time = time
												if r.Version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:174
		_go_fuzz_dep_.CoverTab[105466]++
													tmp.maxOffsets = maxOffsets
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:175
		// _ = "end of CoverTab[105466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:176
		_go_fuzz_dep_.CoverTab[105467]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:176
		// _ = "end of CoverTab[105467]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:176
	// _ = "end of CoverTab[105460]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:176
	_go_fuzz_dep_.CoverTab[105461]++

												r.blocks[topic][partitionID] = tmp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:178
	// _ = "end of CoverTab[105461]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:179
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_request.go:179
var _ = _go_fuzz_dep_.CoverTab
