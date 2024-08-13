//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:1
)

type OffsetFetchResponseBlock struct {
	Offset		int64
	LeaderEpoch	int32
	Metadata	string
	Err		KError
}

func (b *OffsetFetchResponseBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:10
	_go_fuzz_dep_.CoverTab[105039]++
													isFlexible := version >= 6

													b.Offset, err = pd.getInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:14
		_go_fuzz_dep_.CoverTab[105046]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:15
		// _ = "end of CoverTab[105046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:16
		_go_fuzz_dep_.CoverTab[105047]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:16
		// _ = "end of CoverTab[105047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:16
	// _ = "end of CoverTab[105039]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:16
	_go_fuzz_dep_.CoverTab[105040]++

													if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:18
		_go_fuzz_dep_.CoverTab[105048]++
														b.LeaderEpoch, err = pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:20
			_go_fuzz_dep_.CoverTab[105049]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:21
			// _ = "end of CoverTab[105049]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:22
			_go_fuzz_dep_.CoverTab[105050]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:22
			// _ = "end of CoverTab[105050]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:22
		// _ = "end of CoverTab[105048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:23
		_go_fuzz_dep_.CoverTab[105051]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:23
		// _ = "end of CoverTab[105051]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:23
	// _ = "end of CoverTab[105040]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:23
	_go_fuzz_dep_.CoverTab[105041]++

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:25
		_go_fuzz_dep_.CoverTab[105052]++
														b.Metadata, err = pd.getCompactString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:26
		// _ = "end of CoverTab[105052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:27
		_go_fuzz_dep_.CoverTab[105053]++
														b.Metadata, err = pd.getString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:28
		// _ = "end of CoverTab[105053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:29
	// _ = "end of CoverTab[105041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:29
	_go_fuzz_dep_.CoverTab[105042]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:30
		_go_fuzz_dep_.CoverTab[105054]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:31
		// _ = "end of CoverTab[105054]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:32
		_go_fuzz_dep_.CoverTab[105055]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:32
		// _ = "end of CoverTab[105055]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:32
	// _ = "end of CoverTab[105042]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:32
	_go_fuzz_dep_.CoverTab[105043]++

													tmp, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:35
		_go_fuzz_dep_.CoverTab[105056]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:36
		// _ = "end of CoverTab[105056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:37
		_go_fuzz_dep_.CoverTab[105057]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:37
		// _ = "end of CoverTab[105057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:37
	// _ = "end of CoverTab[105043]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:37
	_go_fuzz_dep_.CoverTab[105044]++
													b.Err = KError(tmp)

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:40
		_go_fuzz_dep_.CoverTab[105058]++
														if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:41
			_go_fuzz_dep_.CoverTab[105059]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:42
			// _ = "end of CoverTab[105059]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:43
			_go_fuzz_dep_.CoverTab[105060]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:43
			// _ = "end of CoverTab[105060]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:43
		// _ = "end of CoverTab[105058]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:44
		_go_fuzz_dep_.CoverTab[105061]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:44
		// _ = "end of CoverTab[105061]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:44
	// _ = "end of CoverTab[105044]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:44
	_go_fuzz_dep_.CoverTab[105045]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:46
	// _ = "end of CoverTab[105045]"
}

func (b *OffsetFetchResponseBlock) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:49
	_go_fuzz_dep_.CoverTab[105062]++
													isFlexible := version >= 6
													pe.putInt64(b.Offset)

													if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:53
		_go_fuzz_dep_.CoverTab[105067]++
														pe.putInt32(b.LeaderEpoch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:54
		// _ = "end of CoverTab[105067]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:55
		_go_fuzz_dep_.CoverTab[105068]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:55
		// _ = "end of CoverTab[105068]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:55
	// _ = "end of CoverTab[105062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:55
	_go_fuzz_dep_.CoverTab[105063]++
													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:56
		_go_fuzz_dep_.CoverTab[105069]++
														err = pe.putCompactString(b.Metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:57
		// _ = "end of CoverTab[105069]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:58
		_go_fuzz_dep_.CoverTab[105070]++
														err = pe.putString(b.Metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:59
		// _ = "end of CoverTab[105070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:60
	// _ = "end of CoverTab[105063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:60
	_go_fuzz_dep_.CoverTab[105064]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:61
		_go_fuzz_dep_.CoverTab[105071]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:62
		// _ = "end of CoverTab[105071]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:63
		_go_fuzz_dep_.CoverTab[105072]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:63
		// _ = "end of CoverTab[105072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:63
	// _ = "end of CoverTab[105064]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:63
	_go_fuzz_dep_.CoverTab[105065]++

													pe.putInt16(int16(b.Err))

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:67
		_go_fuzz_dep_.CoverTab[105073]++
														pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:68
		// _ = "end of CoverTab[105073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:69
		_go_fuzz_dep_.CoverTab[105074]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:69
		// _ = "end of CoverTab[105074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:69
	// _ = "end of CoverTab[105065]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:69
	_go_fuzz_dep_.CoverTab[105066]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:71
	// _ = "end of CoverTab[105066]"
}

type OffsetFetchResponse struct {
	Version		int16
	ThrottleTimeMs	int32
	Blocks		map[string]map[int32]*OffsetFetchResponseBlock
	Err		KError
}

func (r *OffsetFetchResponse) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:81
	_go_fuzz_dep_.CoverTab[105075]++
													isFlexible := r.Version >= 6

													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:84
		_go_fuzz_dep_.CoverTab[105082]++
														pe.putInt32(r.ThrottleTimeMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:85
		// _ = "end of CoverTab[105082]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:86
		_go_fuzz_dep_.CoverTab[105083]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:86
		// _ = "end of CoverTab[105083]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:86
	// _ = "end of CoverTab[105075]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:86
	_go_fuzz_dep_.CoverTab[105076]++
													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:87
		_go_fuzz_dep_.CoverTab[105084]++
														pe.putCompactArrayLength(len(r.Blocks))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:88
		// _ = "end of CoverTab[105084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:89
		_go_fuzz_dep_.CoverTab[105085]++
														err = pe.putArrayLength(len(r.Blocks))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:90
		// _ = "end of CoverTab[105085]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:91
	// _ = "end of CoverTab[105076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:91
	_go_fuzz_dep_.CoverTab[105077]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:92
		_go_fuzz_dep_.CoverTab[105086]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:93
		// _ = "end of CoverTab[105086]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:94
		_go_fuzz_dep_.CoverTab[105087]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:94
		// _ = "end of CoverTab[105087]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:94
	// _ = "end of CoverTab[105077]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:94
	_go_fuzz_dep_.CoverTab[105078]++

													for topic, partitions := range r.Blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:96
		_go_fuzz_dep_.CoverTab[105088]++
														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:97
			_go_fuzz_dep_.CoverTab[105094]++
															err = pe.putCompactString(topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:98
			// _ = "end of CoverTab[105094]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:99
			_go_fuzz_dep_.CoverTab[105095]++
															err = pe.putString(topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:100
			// _ = "end of CoverTab[105095]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:101
		// _ = "end of CoverTab[105088]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:101
		_go_fuzz_dep_.CoverTab[105089]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:102
			_go_fuzz_dep_.CoverTab[105096]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:103
			// _ = "end of CoverTab[105096]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:104
			_go_fuzz_dep_.CoverTab[105097]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:104
			// _ = "end of CoverTab[105097]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:104
		// _ = "end of CoverTab[105089]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:104
		_go_fuzz_dep_.CoverTab[105090]++

														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:106
			_go_fuzz_dep_.CoverTab[105098]++
															pe.putCompactArrayLength(len(partitions))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:107
			// _ = "end of CoverTab[105098]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:108
			_go_fuzz_dep_.CoverTab[105099]++
															err = pe.putArrayLength(len(partitions))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:109
			// _ = "end of CoverTab[105099]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:110
		// _ = "end of CoverTab[105090]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:110
		_go_fuzz_dep_.CoverTab[105091]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:111
			_go_fuzz_dep_.CoverTab[105100]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:112
			// _ = "end of CoverTab[105100]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:113
			_go_fuzz_dep_.CoverTab[105101]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:113
			// _ = "end of CoverTab[105101]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:113
		// _ = "end of CoverTab[105091]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:113
		_go_fuzz_dep_.CoverTab[105092]++
														for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:114
			_go_fuzz_dep_.CoverTab[105102]++
															pe.putInt32(partition)
															if err := block.encode(pe, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:116
				_go_fuzz_dep_.CoverTab[105103]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:117
				// _ = "end of CoverTab[105103]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:118
				_go_fuzz_dep_.CoverTab[105104]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:118
				// _ = "end of CoverTab[105104]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:118
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:118
			// _ = "end of CoverTab[105102]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:119
		// _ = "end of CoverTab[105092]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:119
		_go_fuzz_dep_.CoverTab[105093]++
														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:120
			_go_fuzz_dep_.CoverTab[105105]++
															pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:121
			// _ = "end of CoverTab[105105]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:122
			_go_fuzz_dep_.CoverTab[105106]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:122
			// _ = "end of CoverTab[105106]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:122
		// _ = "end of CoverTab[105093]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:123
	// _ = "end of CoverTab[105078]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:123
	_go_fuzz_dep_.CoverTab[105079]++
													if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:124
		_go_fuzz_dep_.CoverTab[105107]++
														pe.putInt16(int16(r.Err))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:125
		// _ = "end of CoverTab[105107]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:126
		_go_fuzz_dep_.CoverTab[105108]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:126
		// _ = "end of CoverTab[105108]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:126
	// _ = "end of CoverTab[105079]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:126
	_go_fuzz_dep_.CoverTab[105080]++
													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:127
		_go_fuzz_dep_.CoverTab[105109]++
														pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:128
		// _ = "end of CoverTab[105109]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:129
		_go_fuzz_dep_.CoverTab[105110]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:129
		// _ = "end of CoverTab[105110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:129
	// _ = "end of CoverTab[105080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:129
	_go_fuzz_dep_.CoverTab[105081]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:130
	// _ = "end of CoverTab[105081]"
}

func (r *OffsetFetchResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:133
	_go_fuzz_dep_.CoverTab[105111]++
													r.Version = version
													isFlexible := version >= 6

													if version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:137
		_go_fuzz_dep_.CoverTab[105118]++
														r.ThrottleTimeMs, err = pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:139
			_go_fuzz_dep_.CoverTab[105119]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:140
			// _ = "end of CoverTab[105119]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:141
			_go_fuzz_dep_.CoverTab[105120]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:141
			// _ = "end of CoverTab[105120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:141
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:141
		// _ = "end of CoverTab[105118]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:142
		_go_fuzz_dep_.CoverTab[105121]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:142
		// _ = "end of CoverTab[105121]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:142
	// _ = "end of CoverTab[105111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:142
	_go_fuzz_dep_.CoverTab[105112]++

													var numTopics int
													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:145
		_go_fuzz_dep_.CoverTab[105122]++
														numTopics, err = pd.getCompactArrayLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:146
		// _ = "end of CoverTab[105122]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:147
		_go_fuzz_dep_.CoverTab[105123]++
														numTopics, err = pd.getArrayLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:148
		// _ = "end of CoverTab[105123]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:149
	// _ = "end of CoverTab[105112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:149
	_go_fuzz_dep_.CoverTab[105113]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:150
		_go_fuzz_dep_.CoverTab[105124]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:151
		// _ = "end of CoverTab[105124]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:152
		_go_fuzz_dep_.CoverTab[105125]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:152
		// _ = "end of CoverTab[105125]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:152
	// _ = "end of CoverTab[105113]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:152
	_go_fuzz_dep_.CoverTab[105114]++

													if numTopics > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:154
		_go_fuzz_dep_.CoverTab[105126]++
														r.Blocks = make(map[string]map[int32]*OffsetFetchResponseBlock, numTopics)
														for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:156
			_go_fuzz_dep_.CoverTab[105127]++
															var name string
															if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:158
				_go_fuzz_dep_.CoverTab[105134]++
																name, err = pd.getCompactString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:159
				// _ = "end of CoverTab[105134]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:160
				_go_fuzz_dep_.CoverTab[105135]++
																name, err = pd.getString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:161
				// _ = "end of CoverTab[105135]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:162
			// _ = "end of CoverTab[105127]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:162
			_go_fuzz_dep_.CoverTab[105128]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:163
				_go_fuzz_dep_.CoverTab[105136]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:164
				// _ = "end of CoverTab[105136]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:165
				_go_fuzz_dep_.CoverTab[105137]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:165
				// _ = "end of CoverTab[105137]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:165
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:165
			// _ = "end of CoverTab[105128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:165
			_go_fuzz_dep_.CoverTab[105129]++

															var numBlocks int
															if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:168
				_go_fuzz_dep_.CoverTab[105138]++
																numBlocks, err = pd.getCompactArrayLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:169
				// _ = "end of CoverTab[105138]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:170
				_go_fuzz_dep_.CoverTab[105139]++
																numBlocks, err = pd.getArrayLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:171
				// _ = "end of CoverTab[105139]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:172
			// _ = "end of CoverTab[105129]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:172
			_go_fuzz_dep_.CoverTab[105130]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:173
				_go_fuzz_dep_.CoverTab[105140]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:174
				// _ = "end of CoverTab[105140]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:175
				_go_fuzz_dep_.CoverTab[105141]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:175
				// _ = "end of CoverTab[105141]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:175
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:175
			// _ = "end of CoverTab[105130]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:175
			_go_fuzz_dep_.CoverTab[105131]++

															r.Blocks[name] = nil
															if numBlocks > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:178
				_go_fuzz_dep_.CoverTab[105142]++
																r.Blocks[name] = make(map[int32]*OffsetFetchResponseBlock, numBlocks)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:179
				// _ = "end of CoverTab[105142]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:180
				_go_fuzz_dep_.CoverTab[105143]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:180
				// _ = "end of CoverTab[105143]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:180
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:180
			// _ = "end of CoverTab[105131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:180
			_go_fuzz_dep_.CoverTab[105132]++
															for j := 0; j < numBlocks; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:181
				_go_fuzz_dep_.CoverTab[105144]++
																id, err := pd.getInt32()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:183
					_go_fuzz_dep_.CoverTab[105147]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:184
					// _ = "end of CoverTab[105147]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:185
					_go_fuzz_dep_.CoverTab[105148]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:185
					// _ = "end of CoverTab[105148]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:185
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:185
				// _ = "end of CoverTab[105144]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:185
				_go_fuzz_dep_.CoverTab[105145]++

																block := new(OffsetFetchResponseBlock)
																err = block.decode(pd, version)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:189
					_go_fuzz_dep_.CoverTab[105149]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:190
					// _ = "end of CoverTab[105149]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:191
					_go_fuzz_dep_.CoverTab[105150]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:191
					// _ = "end of CoverTab[105150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:191
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:191
				// _ = "end of CoverTab[105145]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:191
				_go_fuzz_dep_.CoverTab[105146]++

																r.Blocks[name][id] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:193
				// _ = "end of CoverTab[105146]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:194
			// _ = "end of CoverTab[105132]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:194
			_go_fuzz_dep_.CoverTab[105133]++

															if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:196
				_go_fuzz_dep_.CoverTab[105151]++
																if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:197
					_go_fuzz_dep_.CoverTab[105152]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:198
					// _ = "end of CoverTab[105152]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:199
					_go_fuzz_dep_.CoverTab[105153]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:199
					// _ = "end of CoverTab[105153]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:199
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:199
				// _ = "end of CoverTab[105151]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:200
				_go_fuzz_dep_.CoverTab[105154]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:200
				// _ = "end of CoverTab[105154]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:200
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:200
			// _ = "end of CoverTab[105133]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:201
		// _ = "end of CoverTab[105126]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:202
		_go_fuzz_dep_.CoverTab[105155]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:202
		// _ = "end of CoverTab[105155]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:202
	// _ = "end of CoverTab[105114]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:202
	_go_fuzz_dep_.CoverTab[105115]++

													if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:204
		_go_fuzz_dep_.CoverTab[105156]++
														kerr, err := pd.getInt16()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:206
			_go_fuzz_dep_.CoverTab[105158]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:207
			// _ = "end of CoverTab[105158]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:208
			_go_fuzz_dep_.CoverTab[105159]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:208
			// _ = "end of CoverTab[105159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:208
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:208
		// _ = "end of CoverTab[105156]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:208
		_go_fuzz_dep_.CoverTab[105157]++
														r.Err = KError(kerr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:209
		// _ = "end of CoverTab[105157]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:210
		_go_fuzz_dep_.CoverTab[105160]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:210
		// _ = "end of CoverTab[105160]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:210
	// _ = "end of CoverTab[105115]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:210
	_go_fuzz_dep_.CoverTab[105116]++

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:212
		_go_fuzz_dep_.CoverTab[105161]++
														if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:213
			_go_fuzz_dep_.CoverTab[105162]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:214
			// _ = "end of CoverTab[105162]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:215
			_go_fuzz_dep_.CoverTab[105163]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:215
			// _ = "end of CoverTab[105163]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:215
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:215
		// _ = "end of CoverTab[105161]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:216
		_go_fuzz_dep_.CoverTab[105164]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:216
		// _ = "end of CoverTab[105164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:216
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:216
	// _ = "end of CoverTab[105116]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:216
	_go_fuzz_dep_.CoverTab[105117]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:218
	// _ = "end of CoverTab[105117]"
}

func (r *OffsetFetchResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:221
	_go_fuzz_dep_.CoverTab[105165]++
													return 9
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:222
	// _ = "end of CoverTab[105165]"
}

func (r *OffsetFetchResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:225
	_go_fuzz_dep_.CoverTab[105166]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:226
	// _ = "end of CoverTab[105166]"
}

func (r *OffsetFetchResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:229
	_go_fuzz_dep_.CoverTab[105167]++
													if r.Version >= 6 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:230
		_go_fuzz_dep_.CoverTab[105169]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:231
		// _ = "end of CoverTab[105169]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:232
		_go_fuzz_dep_.CoverTab[105170]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:232
		// _ = "end of CoverTab[105170]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:232
	// _ = "end of CoverTab[105167]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:232
	_go_fuzz_dep_.CoverTab[105168]++

													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:234
	// _ = "end of CoverTab[105168]"
}

func (r *OffsetFetchResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:237
	_go_fuzz_dep_.CoverTab[105171]++
													switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:239
		_go_fuzz_dep_.CoverTab[105172]++
														return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:240
		// _ = "end of CoverTab[105172]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:241
		_go_fuzz_dep_.CoverTab[105173]++
														return V0_10_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:242
		// _ = "end of CoverTab[105173]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:243
		_go_fuzz_dep_.CoverTab[105174]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:244
		// _ = "end of CoverTab[105174]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:245
		_go_fuzz_dep_.CoverTab[105175]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:246
		// _ = "end of CoverTab[105175]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:247
		_go_fuzz_dep_.CoverTab[105176]++
														return V2_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:248
		// _ = "end of CoverTab[105176]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:249
		_go_fuzz_dep_.CoverTab[105177]++
														return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:250
		// _ = "end of CoverTab[105177]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:251
		_go_fuzz_dep_.CoverTab[105178]++
														return V2_5_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:252
		// _ = "end of CoverTab[105178]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:253
		_go_fuzz_dep_.CoverTab[105179]++
														return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:254
		// _ = "end of CoverTab[105179]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:255
	// _ = "end of CoverTab[105171]"
}

func (r *OffsetFetchResponse) GetBlock(topic string, partition int32) *OffsetFetchResponseBlock {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:258
	_go_fuzz_dep_.CoverTab[105180]++
													if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:259
		_go_fuzz_dep_.CoverTab[105183]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:260
		// _ = "end of CoverTab[105183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:261
		_go_fuzz_dep_.CoverTab[105184]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:261
		// _ = "end of CoverTab[105184]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:261
	// _ = "end of CoverTab[105180]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:261
	_go_fuzz_dep_.CoverTab[105181]++

													if r.Blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:263
		_go_fuzz_dep_.CoverTab[105185]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:264
		// _ = "end of CoverTab[105185]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:265
		_go_fuzz_dep_.CoverTab[105186]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:265
		// _ = "end of CoverTab[105186]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:265
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:265
	// _ = "end of CoverTab[105181]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:265
	_go_fuzz_dep_.CoverTab[105182]++

													return r.Blocks[topic][partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:267
	// _ = "end of CoverTab[105182]"
}

func (r *OffsetFetchResponse) AddBlock(topic string, partition int32, block *OffsetFetchResponseBlock) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:270
	_go_fuzz_dep_.CoverTab[105187]++
													if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:271
		_go_fuzz_dep_.CoverTab[105190]++
														r.Blocks = make(map[string]map[int32]*OffsetFetchResponseBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:272
		// _ = "end of CoverTab[105190]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:273
		_go_fuzz_dep_.CoverTab[105191]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:273
		// _ = "end of CoverTab[105191]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:273
	// _ = "end of CoverTab[105187]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:273
	_go_fuzz_dep_.CoverTab[105188]++
													partitions := r.Blocks[topic]
													if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:275
		_go_fuzz_dep_.CoverTab[105192]++
														partitions = make(map[int32]*OffsetFetchResponseBlock)
														r.Blocks[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:277
		// _ = "end of CoverTab[105192]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:278
		_go_fuzz_dep_.CoverTab[105193]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:278
		// _ = "end of CoverTab[105193]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:278
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:278
	// _ = "end of CoverTab[105188]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:278
	_go_fuzz_dep_.CoverTab[105189]++
													partitions[partition] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:279
	// _ = "end of CoverTab[105189]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:280
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_response.go:280
var _ = _go_fuzz_dep_.CoverTab
