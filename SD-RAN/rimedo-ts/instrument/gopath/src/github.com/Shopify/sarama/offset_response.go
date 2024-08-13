//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:1
)

type OffsetResponseBlock struct {
	Err		KError
	Offsets		[]int64	// Version 0
	Offset		int64	// Version 1
	Timestamp	int64	// Version 1
}

func (b *OffsetResponseBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:10
	_go_fuzz_dep_.CoverTab[105468]++
												tmp, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:12
		_go_fuzz_dep_.CoverTab[105473]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:13
		// _ = "end of CoverTab[105473]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:14
		_go_fuzz_dep_.CoverTab[105474]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:14
		// _ = "end of CoverTab[105474]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:14
	// _ = "end of CoverTab[105468]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:14
	_go_fuzz_dep_.CoverTab[105469]++
												b.Err = KError(tmp)

												if version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:17
		_go_fuzz_dep_.CoverTab[105475]++
													b.Offsets, err = pd.getInt64Array()

													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:20
		// _ = "end of CoverTab[105475]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:21
		_go_fuzz_dep_.CoverTab[105476]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:21
		// _ = "end of CoverTab[105476]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:21
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:21
	// _ = "end of CoverTab[105469]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:21
	_go_fuzz_dep_.CoverTab[105470]++

												b.Timestamp, err = pd.getInt64()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:24
		_go_fuzz_dep_.CoverTab[105477]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:25
		// _ = "end of CoverTab[105477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:26
		_go_fuzz_dep_.CoverTab[105478]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:26
		// _ = "end of CoverTab[105478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:26
	// _ = "end of CoverTab[105470]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:26
	_go_fuzz_dep_.CoverTab[105471]++

												b.Offset, err = pd.getInt64()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:29
		_go_fuzz_dep_.CoverTab[105479]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:30
		// _ = "end of CoverTab[105479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:31
		_go_fuzz_dep_.CoverTab[105480]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:31
		// _ = "end of CoverTab[105480]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:31
	// _ = "end of CoverTab[105471]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:31
	_go_fuzz_dep_.CoverTab[105472]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:34
	b.Offsets = []int64{b.Offset}

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:36
	// _ = "end of CoverTab[105472]"
}

func (b *OffsetResponseBlock) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:39
	_go_fuzz_dep_.CoverTab[105481]++
												pe.putInt16(int16(b.Err))

												if version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:42
		_go_fuzz_dep_.CoverTab[105483]++
													return pe.putInt64Array(b.Offsets)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:43
		// _ = "end of CoverTab[105483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:44
		_go_fuzz_dep_.CoverTab[105484]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:44
		// _ = "end of CoverTab[105484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:44
	// _ = "end of CoverTab[105481]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:44
	_go_fuzz_dep_.CoverTab[105482]++

												pe.putInt64(b.Timestamp)
												pe.putInt64(b.Offset)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:49
	// _ = "end of CoverTab[105482]"
}

type OffsetResponse struct {
	Version		int16
	ThrottleTimeMs	int32
	Blocks		map[string]map[int32]*OffsetResponseBlock
}

func (r *OffsetResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:58
	_go_fuzz_dep_.CoverTab[105485]++
												if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:59
		_go_fuzz_dep_.CoverTab[105489]++
													r.ThrottleTimeMs, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:61
			_go_fuzz_dep_.CoverTab[105490]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:62
			// _ = "end of CoverTab[105490]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:63
			_go_fuzz_dep_.CoverTab[105491]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:63
			// _ = "end of CoverTab[105491]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:63
		// _ = "end of CoverTab[105489]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:64
		_go_fuzz_dep_.CoverTab[105492]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:64
		// _ = "end of CoverTab[105492]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:64
	// _ = "end of CoverTab[105485]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:64
	_go_fuzz_dep_.CoverTab[105486]++

												numTopics, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:67
		_go_fuzz_dep_.CoverTab[105493]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:68
		// _ = "end of CoverTab[105493]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:69
		_go_fuzz_dep_.CoverTab[105494]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:69
		// _ = "end of CoverTab[105494]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:69
	// _ = "end of CoverTab[105486]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:69
	_go_fuzz_dep_.CoverTab[105487]++

												r.Blocks = make(map[string]map[int32]*OffsetResponseBlock, numTopics)
												for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:72
		_go_fuzz_dep_.CoverTab[105495]++
													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:74
			_go_fuzz_dep_.CoverTab[105498]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:75
			// _ = "end of CoverTab[105498]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:76
			_go_fuzz_dep_.CoverTab[105499]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:76
			// _ = "end of CoverTab[105499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:76
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:76
		// _ = "end of CoverTab[105495]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:76
		_go_fuzz_dep_.CoverTab[105496]++

													numBlocks, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:79
			_go_fuzz_dep_.CoverTab[105500]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:80
			// _ = "end of CoverTab[105500]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:81
			_go_fuzz_dep_.CoverTab[105501]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:81
			// _ = "end of CoverTab[105501]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:81
		// _ = "end of CoverTab[105496]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:81
		_go_fuzz_dep_.CoverTab[105497]++

													r.Blocks[name] = make(map[int32]*OffsetResponseBlock, numBlocks)

													for j := 0; j < numBlocks; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:85
			_go_fuzz_dep_.CoverTab[105502]++
														id, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:87
				_go_fuzz_dep_.CoverTab[105505]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:88
				// _ = "end of CoverTab[105505]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:89
				_go_fuzz_dep_.CoverTab[105506]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:89
				// _ = "end of CoverTab[105506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:89
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:89
			// _ = "end of CoverTab[105502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:89
			_go_fuzz_dep_.CoverTab[105503]++

														block := new(OffsetResponseBlock)
														err = block.decode(pd, version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:93
				_go_fuzz_dep_.CoverTab[105507]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:94
				// _ = "end of CoverTab[105507]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:95
				_go_fuzz_dep_.CoverTab[105508]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:95
				// _ = "end of CoverTab[105508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:95
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:95
			// _ = "end of CoverTab[105503]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:95
			_go_fuzz_dep_.CoverTab[105504]++
														r.Blocks[name][id] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:96
			// _ = "end of CoverTab[105504]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:97
		// _ = "end of CoverTab[105497]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:98
	// _ = "end of CoverTab[105487]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:98
	_go_fuzz_dep_.CoverTab[105488]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:100
	// _ = "end of CoverTab[105488]"
}

func (r *OffsetResponse) GetBlock(topic string, partition int32) *OffsetResponseBlock {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:103
	_go_fuzz_dep_.CoverTab[105509]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:104
		_go_fuzz_dep_.CoverTab[105512]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:105
		// _ = "end of CoverTab[105512]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:106
		_go_fuzz_dep_.CoverTab[105513]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:106
		// _ = "end of CoverTab[105513]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:106
	// _ = "end of CoverTab[105509]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:106
	_go_fuzz_dep_.CoverTab[105510]++

												if r.Blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:108
		_go_fuzz_dep_.CoverTab[105514]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:109
		// _ = "end of CoverTab[105514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:110
		_go_fuzz_dep_.CoverTab[105515]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:110
		// _ = "end of CoverTab[105515]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:110
	// _ = "end of CoverTab[105510]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:110
	_go_fuzz_dep_.CoverTab[105511]++

												return r.Blocks[topic][partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:112
	// _ = "end of CoverTab[105511]"
}

/*
// [0 0 0 1 ntopics
0 8 109 121 95 116 111 112 105 99 topic
0 0 0 1 npartitions
0 0 0 0 id
0 0

0 0 0 1 0 0 0 0
0 1 1 1 0 0 0 1
0 8 109 121 95 116 111 112
105 99 0 0 0 1 0 0
0 0 0 0 0 0 0 1
0 0 0 0 0 1 1 1] <nil>
*/
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:130
func (r *OffsetResponse) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:130
	_go_fuzz_dep_.CoverTab[105516]++
												if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:131
		_go_fuzz_dep_.CoverTab[105520]++
													pe.putInt32(r.ThrottleTimeMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:132
		// _ = "end of CoverTab[105520]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:133
		_go_fuzz_dep_.CoverTab[105521]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:133
		// _ = "end of CoverTab[105521]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:133
	// _ = "end of CoverTab[105516]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:133
	_go_fuzz_dep_.CoverTab[105517]++

												if err = pe.putArrayLength(len(r.Blocks)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:135
		_go_fuzz_dep_.CoverTab[105522]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:136
		// _ = "end of CoverTab[105522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:137
		_go_fuzz_dep_.CoverTab[105523]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:137
		// _ = "end of CoverTab[105523]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:137
	// _ = "end of CoverTab[105517]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:137
	_go_fuzz_dep_.CoverTab[105518]++

												for topic, partitions := range r.Blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:139
		_go_fuzz_dep_.CoverTab[105524]++
													if err = pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:140
			_go_fuzz_dep_.CoverTab[105527]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:141
			// _ = "end of CoverTab[105527]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:142
			_go_fuzz_dep_.CoverTab[105528]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:142
			// _ = "end of CoverTab[105528]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:142
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:142
		// _ = "end of CoverTab[105524]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:142
		_go_fuzz_dep_.CoverTab[105525]++
													if err = pe.putArrayLength(len(partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:143
			_go_fuzz_dep_.CoverTab[105529]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:144
			// _ = "end of CoverTab[105529]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:145
			_go_fuzz_dep_.CoverTab[105530]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:145
			// _ = "end of CoverTab[105530]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:145
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:145
		// _ = "end of CoverTab[105525]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:145
		_go_fuzz_dep_.CoverTab[105526]++
													for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:146
			_go_fuzz_dep_.CoverTab[105531]++
														pe.putInt32(partition)
														if err = block.encode(pe, r.version()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:148
				_go_fuzz_dep_.CoverTab[105532]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:149
				// _ = "end of CoverTab[105532]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:150
				_go_fuzz_dep_.CoverTab[105533]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:150
				// _ = "end of CoverTab[105533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:150
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:150
			// _ = "end of CoverTab[105531]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:151
		// _ = "end of CoverTab[105526]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:152
	// _ = "end of CoverTab[105518]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:152
	_go_fuzz_dep_.CoverTab[105519]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:154
	// _ = "end of CoverTab[105519]"
}

func (r *OffsetResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:157
	_go_fuzz_dep_.CoverTab[105534]++
												return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:158
	// _ = "end of CoverTab[105534]"
}

func (r *OffsetResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:161
	_go_fuzz_dep_.CoverTab[105535]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:162
	// _ = "end of CoverTab[105535]"
}

func (r *OffsetResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:165
	_go_fuzz_dep_.CoverTab[105536]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:166
	// _ = "end of CoverTab[105536]"
}

func (r *OffsetResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:169
	_go_fuzz_dep_.CoverTab[105537]++
												switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:171
		_go_fuzz_dep_.CoverTab[105538]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:172
		// _ = "end of CoverTab[105538]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:173
		_go_fuzz_dep_.CoverTab[105539]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:174
		// _ = "end of CoverTab[105539]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:175
		_go_fuzz_dep_.CoverTab[105540]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:176
		// _ = "end of CoverTab[105540]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:177
	// _ = "end of CoverTab[105537]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:182
func (r *OffsetResponse) AddTopicPartition(topic string, partition int32, offset int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:182
	_go_fuzz_dep_.CoverTab[105541]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:183
		_go_fuzz_dep_.CoverTab[105544]++
													r.Blocks = make(map[string]map[int32]*OffsetResponseBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:184
		// _ = "end of CoverTab[105544]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:185
		_go_fuzz_dep_.CoverTab[105545]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:185
		// _ = "end of CoverTab[105545]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:185
	// _ = "end of CoverTab[105541]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:185
	_go_fuzz_dep_.CoverTab[105542]++
												byTopic, ok := r.Blocks[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:187
		_go_fuzz_dep_.CoverTab[105546]++
													byTopic = make(map[int32]*OffsetResponseBlock)
													r.Blocks[topic] = byTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:189
		// _ = "end of CoverTab[105546]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:190
		_go_fuzz_dep_.CoverTab[105547]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:190
		// _ = "end of CoverTab[105547]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:190
	// _ = "end of CoverTab[105542]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:190
	_go_fuzz_dep_.CoverTab[105543]++
												byTopic[partition] = &OffsetResponseBlock{Offsets: []int64{offset}, Offset: offset}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:191
	// _ = "end of CoverTab[105543]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:192
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_response.go:192
var _ = _go_fuzz_dep_.CoverTab
