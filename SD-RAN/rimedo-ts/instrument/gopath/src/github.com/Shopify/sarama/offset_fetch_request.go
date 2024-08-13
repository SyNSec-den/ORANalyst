//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:1
)

type OffsetFetchRequest struct {
	Version		int16
	ConsumerGroup	string
	RequireStable	bool	// requires v7+
	partitions	map[string][]int32
}

func (r *OffsetFetchRequest) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:10
	_go_fuzz_dep_.CoverTab[104922]++
													if r.Version < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:11
		_go_fuzz_dep_.CoverTab[104931]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:11
		return r.Version > 7
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:11
		// _ = "end of CoverTab[104931]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:11
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:11
		_go_fuzz_dep_.CoverTab[104932]++
														return PacketEncodingError{"invalid or unsupported OffsetFetchRequest version field"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:12
		// _ = "end of CoverTab[104932]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:13
		_go_fuzz_dep_.CoverTab[104933]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:13
		// _ = "end of CoverTab[104933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:13
	// _ = "end of CoverTab[104922]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:13
	_go_fuzz_dep_.CoverTab[104923]++

													isFlexible := r.Version >= 6

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:17
		_go_fuzz_dep_.CoverTab[104934]++
														err = pe.putCompactString(r.ConsumerGroup)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:18
		// _ = "end of CoverTab[104934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:19
		_go_fuzz_dep_.CoverTab[104935]++
														err = pe.putString(r.ConsumerGroup)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:20
		// _ = "end of CoverTab[104935]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:21
	// _ = "end of CoverTab[104923]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:21
	_go_fuzz_dep_.CoverTab[104924]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:22
		_go_fuzz_dep_.CoverTab[104936]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:23
		// _ = "end of CoverTab[104936]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:24
		_go_fuzz_dep_.CoverTab[104937]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:24
		// _ = "end of CoverTab[104937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:24
	// _ = "end of CoverTab[104924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:24
	_go_fuzz_dep_.CoverTab[104925]++

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:26
		_go_fuzz_dep_.CoverTab[104938]++
														if r.partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:27
			_go_fuzz_dep_.CoverTab[104939]++
															pe.putUVarint(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:28
			// _ = "end of CoverTab[104939]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:29
			_go_fuzz_dep_.CoverTab[104940]++
															pe.putCompactArrayLength(len(r.partitions))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:30
			// _ = "end of CoverTab[104940]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:31
		// _ = "end of CoverTab[104938]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:32
		_go_fuzz_dep_.CoverTab[104941]++
														if r.partitions == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:33
			_go_fuzz_dep_.CoverTab[104942]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:33
			return r.Version >= 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:33
			// _ = "end of CoverTab[104942]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:33
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:33
			_go_fuzz_dep_.CoverTab[104943]++
															pe.putInt32(-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:34
			// _ = "end of CoverTab[104943]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:35
			_go_fuzz_dep_.CoverTab[104944]++
															if err = pe.putArrayLength(len(r.partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:36
				_go_fuzz_dep_.CoverTab[104945]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:37
				// _ = "end of CoverTab[104945]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:38
				_go_fuzz_dep_.CoverTab[104946]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:38
				// _ = "end of CoverTab[104946]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:38
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:38
			// _ = "end of CoverTab[104944]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:39
		// _ = "end of CoverTab[104941]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:40
	// _ = "end of CoverTab[104925]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:40
	_go_fuzz_dep_.CoverTab[104926]++

													for topic, partitions := range r.partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:42
		_go_fuzz_dep_.CoverTab[104947]++
														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:43
			_go_fuzz_dep_.CoverTab[104952]++
															err = pe.putCompactString(topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:44
			// _ = "end of CoverTab[104952]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:45
			_go_fuzz_dep_.CoverTab[104953]++
															err = pe.putString(topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:46
			// _ = "end of CoverTab[104953]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:47
		// _ = "end of CoverTab[104947]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:47
		_go_fuzz_dep_.CoverTab[104948]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:48
			_go_fuzz_dep_.CoverTab[104954]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:49
			// _ = "end of CoverTab[104954]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:50
			_go_fuzz_dep_.CoverTab[104955]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:50
			// _ = "end of CoverTab[104955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:50
		// _ = "end of CoverTab[104948]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:50
		_go_fuzz_dep_.CoverTab[104949]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:54
		if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:54
			_go_fuzz_dep_.CoverTab[104956]++
															err = pe.putCompactInt32Array(partitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:55
			// _ = "end of CoverTab[104956]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:56
			_go_fuzz_dep_.CoverTab[104957]++
															err = pe.putInt32Array(partitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:57
			// _ = "end of CoverTab[104957]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:58
		// _ = "end of CoverTab[104949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:58
		_go_fuzz_dep_.CoverTab[104950]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:59
			_go_fuzz_dep_.CoverTab[104958]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:60
			// _ = "end of CoverTab[104958]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:61
			_go_fuzz_dep_.CoverTab[104959]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:61
			// _ = "end of CoverTab[104959]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:61
		// _ = "end of CoverTab[104950]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:61
		_go_fuzz_dep_.CoverTab[104951]++

														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:63
			_go_fuzz_dep_.CoverTab[104960]++
															pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:64
			// _ = "end of CoverTab[104960]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:65
			_go_fuzz_dep_.CoverTab[104961]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:65
			// _ = "end of CoverTab[104961]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:65
		// _ = "end of CoverTab[104951]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:66
	// _ = "end of CoverTab[104926]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:66
	_go_fuzz_dep_.CoverTab[104927]++

													if r.RequireStable && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:68
		_go_fuzz_dep_.CoverTab[104962]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:68
		return r.Version < 7
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:68
		// _ = "end of CoverTab[104962]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:68
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:68
		_go_fuzz_dep_.CoverTab[104963]++
														return PacketEncodingError{"requireStable is not supported. use version 7 or later"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:69
		// _ = "end of CoverTab[104963]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:70
		_go_fuzz_dep_.CoverTab[104964]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:70
		// _ = "end of CoverTab[104964]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:70
	// _ = "end of CoverTab[104927]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:70
	_go_fuzz_dep_.CoverTab[104928]++

													if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:72
		_go_fuzz_dep_.CoverTab[104965]++
														pe.putBool(r.RequireStable)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:73
		// _ = "end of CoverTab[104965]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:74
		_go_fuzz_dep_.CoverTab[104966]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:74
		// _ = "end of CoverTab[104966]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:74
	// _ = "end of CoverTab[104928]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:74
	_go_fuzz_dep_.CoverTab[104929]++

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:76
		_go_fuzz_dep_.CoverTab[104967]++
														pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:77
		// _ = "end of CoverTab[104967]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:78
		_go_fuzz_dep_.CoverTab[104968]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:78
		// _ = "end of CoverTab[104968]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:78
	// _ = "end of CoverTab[104929]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:78
	_go_fuzz_dep_.CoverTab[104930]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:80
	// _ = "end of CoverTab[104930]"
}

func (r *OffsetFetchRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:83
	_go_fuzz_dep_.CoverTab[104969]++
													r.Version = version
													isFlexible := r.Version >= 6
													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:86
		_go_fuzz_dep_.CoverTab[104978]++
														r.ConsumerGroup, err = pd.getCompactString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:87
		// _ = "end of CoverTab[104978]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:88
		_go_fuzz_dep_.CoverTab[104979]++
														r.ConsumerGroup, err = pd.getString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:89
		// _ = "end of CoverTab[104979]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:90
	// _ = "end of CoverTab[104969]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:90
	_go_fuzz_dep_.CoverTab[104970]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:91
		_go_fuzz_dep_.CoverTab[104980]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:92
		// _ = "end of CoverTab[104980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:93
		_go_fuzz_dep_.CoverTab[104981]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:93
		// _ = "end of CoverTab[104981]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:93
	// _ = "end of CoverTab[104970]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:93
	_go_fuzz_dep_.CoverTab[104971]++

													var partitionCount int

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:97
		_go_fuzz_dep_.CoverTab[104982]++
														partitionCount, err = pd.getCompactArrayLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:98
		// _ = "end of CoverTab[104982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:99
		_go_fuzz_dep_.CoverTab[104983]++
														partitionCount, err = pd.getArrayLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:100
		// _ = "end of CoverTab[104983]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:101
	// _ = "end of CoverTab[104971]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:101
	_go_fuzz_dep_.CoverTab[104972]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:102
		_go_fuzz_dep_.CoverTab[104984]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:103
		// _ = "end of CoverTab[104984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:104
		_go_fuzz_dep_.CoverTab[104985]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:104
		// _ = "end of CoverTab[104985]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:104
	// _ = "end of CoverTab[104972]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:104
	_go_fuzz_dep_.CoverTab[104973]++

													if (partitionCount == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		_go_fuzz_dep_.CoverTab[104986]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		return version < 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		// _ = "end of CoverTab[104986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		_go_fuzz_dep_.CoverTab[104987]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		return partitionCount < 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		// _ = "end of CoverTab[104987]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:106
		_go_fuzz_dep_.CoverTab[104988]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:107
		// _ = "end of CoverTab[104988]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:108
		_go_fuzz_dep_.CoverTab[104989]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:108
		// _ = "end of CoverTab[104989]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:108
	// _ = "end of CoverTab[104973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:108
	_go_fuzz_dep_.CoverTab[104974]++

													r.partitions = make(map[string][]int32, partitionCount)
													for i := 0; i < partitionCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:111
		_go_fuzz_dep_.CoverTab[104990]++
														var topic string
														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:113
			_go_fuzz_dep_.CoverTab[104996]++
															topic, err = pd.getCompactString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:114
			// _ = "end of CoverTab[104996]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:115
			_go_fuzz_dep_.CoverTab[104997]++
															topic, err = pd.getString()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:116
			// _ = "end of CoverTab[104997]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:117
		// _ = "end of CoverTab[104990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:117
		_go_fuzz_dep_.CoverTab[104991]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:118
			_go_fuzz_dep_.CoverTab[104998]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:119
			// _ = "end of CoverTab[104998]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:120
			_go_fuzz_dep_.CoverTab[104999]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:120
			// _ = "end of CoverTab[104999]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:120
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:120
		// _ = "end of CoverTab[104991]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:120
		_go_fuzz_dep_.CoverTab[104992]++

														var partitions []int32
														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:123
			_go_fuzz_dep_.CoverTab[105000]++
															partitions, err = pd.getCompactInt32Array()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:124
			// _ = "end of CoverTab[105000]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:125
			_go_fuzz_dep_.CoverTab[105001]++
															partitions, err = pd.getInt32Array()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:126
			// _ = "end of CoverTab[105001]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:127
		// _ = "end of CoverTab[104992]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:127
		_go_fuzz_dep_.CoverTab[104993]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:128
			_go_fuzz_dep_.CoverTab[105002]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:129
			// _ = "end of CoverTab[105002]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:130
			_go_fuzz_dep_.CoverTab[105003]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:130
			// _ = "end of CoverTab[105003]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:130
		// _ = "end of CoverTab[104993]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:130
		_go_fuzz_dep_.CoverTab[104994]++
														if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:131
			_go_fuzz_dep_.CoverTab[105004]++
															_, err = pd.getEmptyTaggedFieldArray()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:133
				_go_fuzz_dep_.CoverTab[105005]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:134
				// _ = "end of CoverTab[105005]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:135
				_go_fuzz_dep_.CoverTab[105006]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:135
				// _ = "end of CoverTab[105006]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:135
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:135
			// _ = "end of CoverTab[105004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:136
			_go_fuzz_dep_.CoverTab[105007]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:136
			// _ = "end of CoverTab[105007]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:136
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:136
		// _ = "end of CoverTab[104994]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:136
		_go_fuzz_dep_.CoverTab[104995]++

														r.partitions[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:138
		// _ = "end of CoverTab[104995]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:139
	// _ = "end of CoverTab[104974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:139
	_go_fuzz_dep_.CoverTab[104975]++

													if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:141
		_go_fuzz_dep_.CoverTab[105008]++
														r.RequireStable, err = pd.getBool()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:143
			_go_fuzz_dep_.CoverTab[105009]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:144
			// _ = "end of CoverTab[105009]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:145
			_go_fuzz_dep_.CoverTab[105010]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:145
			// _ = "end of CoverTab[105010]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:145
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:145
		// _ = "end of CoverTab[105008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:146
		_go_fuzz_dep_.CoverTab[105011]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:146
		// _ = "end of CoverTab[105011]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:146
	// _ = "end of CoverTab[104975]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:146
	_go_fuzz_dep_.CoverTab[104976]++

													if isFlexible {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:148
		_go_fuzz_dep_.CoverTab[105012]++
														_, err = pd.getEmptyTaggedFieldArray()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:150
			_go_fuzz_dep_.CoverTab[105013]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:151
			// _ = "end of CoverTab[105013]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:152
			_go_fuzz_dep_.CoverTab[105014]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:152
			// _ = "end of CoverTab[105014]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:152
		// _ = "end of CoverTab[105012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:153
		_go_fuzz_dep_.CoverTab[105015]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:153
		// _ = "end of CoverTab[105015]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:153
	// _ = "end of CoverTab[104976]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:153
	_go_fuzz_dep_.CoverTab[104977]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:155
	// _ = "end of CoverTab[104977]"
}

func (r *OffsetFetchRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:158
	_go_fuzz_dep_.CoverTab[105016]++
													return 9
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:159
	// _ = "end of CoverTab[105016]"
}

func (r *OffsetFetchRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:162
	_go_fuzz_dep_.CoverTab[105017]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:163
	// _ = "end of CoverTab[105017]"
}

func (r *OffsetFetchRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:166
	_go_fuzz_dep_.CoverTab[105018]++
													if r.Version >= 6 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:167
		_go_fuzz_dep_.CoverTab[105020]++
														return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:168
		// _ = "end of CoverTab[105020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:169
		_go_fuzz_dep_.CoverTab[105021]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:169
		// _ = "end of CoverTab[105021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:169
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:169
	// _ = "end of CoverTab[105018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:169
	_go_fuzz_dep_.CoverTab[105019]++

													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:171
	// _ = "end of CoverTab[105019]"
}

func (r *OffsetFetchRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:174
	_go_fuzz_dep_.CoverTab[105022]++
													switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:176
		_go_fuzz_dep_.CoverTab[105023]++
														return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:177
		// _ = "end of CoverTab[105023]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:178
		_go_fuzz_dep_.CoverTab[105024]++
														return V0_10_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:179
		// _ = "end of CoverTab[105024]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:180
		_go_fuzz_dep_.CoverTab[105025]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:181
		// _ = "end of CoverTab[105025]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:182
		_go_fuzz_dep_.CoverTab[105026]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:183
		// _ = "end of CoverTab[105026]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:184
		_go_fuzz_dep_.CoverTab[105027]++
														return V2_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:185
		// _ = "end of CoverTab[105027]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:186
		_go_fuzz_dep_.CoverTab[105028]++
														return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:187
		// _ = "end of CoverTab[105028]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:188
		_go_fuzz_dep_.CoverTab[105029]++
														return V2_5_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:189
		// _ = "end of CoverTab[105029]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:190
		_go_fuzz_dep_.CoverTab[105030]++
														return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:191
		// _ = "end of CoverTab[105030]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:192
	// _ = "end of CoverTab[105022]"
}

func (r *OffsetFetchRequest) ZeroPartitions() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:195
	_go_fuzz_dep_.CoverTab[105031]++
													if r.partitions == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:196
		_go_fuzz_dep_.CoverTab[105032]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:196
		return r.Version >= 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:196
		// _ = "end of CoverTab[105032]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:196
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:196
		_go_fuzz_dep_.CoverTab[105033]++
														r.partitions = make(map[string][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:197
		// _ = "end of CoverTab[105033]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:198
		_go_fuzz_dep_.CoverTab[105034]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:198
		// _ = "end of CoverTab[105034]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:198
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:198
	// _ = "end of CoverTab[105031]"
}

func (r *OffsetFetchRequest) AddPartition(topic string, partitionID int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:201
	_go_fuzz_dep_.CoverTab[105035]++
													if r.partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:202
		_go_fuzz_dep_.CoverTab[105037]++
														r.partitions = make(map[string][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:203
		// _ = "end of CoverTab[105037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:204
		_go_fuzz_dep_.CoverTab[105038]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:204
		// _ = "end of CoverTab[105038]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:204
	// _ = "end of CoverTab[105035]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:204
	_go_fuzz_dep_.CoverTab[105036]++

													r.partitions[topic] = append(r.partitions[topic], partitionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:206
	// _ = "end of CoverTab[105036]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:207
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_fetch_request.go:207
var _ = _go_fuzz_dep_.CoverTab
