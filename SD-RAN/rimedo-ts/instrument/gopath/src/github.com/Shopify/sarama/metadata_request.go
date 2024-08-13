//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:1
)

type MetadataRequest struct {
	Version			int16
	Topics			[]string
	AllowAutoTopicCreation	bool
}

func (r *MetadataRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:9
	_go_fuzz_dep_.CoverTab[104053]++
												if r.Version < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:10
		_go_fuzz_dep_.CoverTab[104057]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:10
		return r.Version > 5
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:10
		// _ = "end of CoverTab[104057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:10
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:10
		_go_fuzz_dep_.CoverTab[104058]++
													return PacketEncodingError{"invalid or unsupported MetadataRequest version field"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:11
		// _ = "end of CoverTab[104058]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:12
		_go_fuzz_dep_.CoverTab[104059]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:12
		// _ = "end of CoverTab[104059]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:12
	// _ = "end of CoverTab[104053]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:12
	_go_fuzz_dep_.CoverTab[104054]++
												if r.Version == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:13
		_go_fuzz_dep_.CoverTab[104060]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:13
		return len(r.Topics) > 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:13
		// _ = "end of CoverTab[104060]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:13
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:13
		_go_fuzz_dep_.CoverTab[104061]++
													err := pe.putArrayLength(len(r.Topics))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:15
			_go_fuzz_dep_.CoverTab[104063]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:16
			// _ = "end of CoverTab[104063]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:17
			_go_fuzz_dep_.CoverTab[104064]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:17
			// _ = "end of CoverTab[104064]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:17
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:17
		// _ = "end of CoverTab[104061]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:17
		_go_fuzz_dep_.CoverTab[104062]++

													for i := range r.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:19
			_go_fuzz_dep_.CoverTab[104065]++
														err = pe.putString(r.Topics[i])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:21
				_go_fuzz_dep_.CoverTab[104066]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:22
				// _ = "end of CoverTab[104066]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:23
				_go_fuzz_dep_.CoverTab[104067]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:23
				// _ = "end of CoverTab[104067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:23
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:23
			// _ = "end of CoverTab[104065]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:24
		// _ = "end of CoverTab[104062]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:25
		_go_fuzz_dep_.CoverTab[104068]++
													pe.putInt32(-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:26
		// _ = "end of CoverTab[104068]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:27
	// _ = "end of CoverTab[104054]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:27
	_go_fuzz_dep_.CoverTab[104055]++
												if r.Version > 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:28
		_go_fuzz_dep_.CoverTab[104069]++
													pe.putBool(r.AllowAutoTopicCreation)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:29
		// _ = "end of CoverTab[104069]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:30
		_go_fuzz_dep_.CoverTab[104070]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:30
		// _ = "end of CoverTab[104070]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:30
	// _ = "end of CoverTab[104055]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:30
	_go_fuzz_dep_.CoverTab[104056]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:31
	// _ = "end of CoverTab[104056]"
}

func (r *MetadataRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:34
	_go_fuzz_dep_.CoverTab[104071]++
												r.Version = version
												size, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:37
		_go_fuzz_dep_.CoverTab[104075]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:38
		// _ = "end of CoverTab[104075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:39
		_go_fuzz_dep_.CoverTab[104076]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:39
		// _ = "end of CoverTab[104076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:39
	// _ = "end of CoverTab[104071]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:39
	_go_fuzz_dep_.CoverTab[104072]++
												if size > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:40
		_go_fuzz_dep_.CoverTab[104077]++
													r.Topics = make([]string, size)
													for i := range r.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:42
			_go_fuzz_dep_.CoverTab[104078]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:44
				_go_fuzz_dep_.CoverTab[104080]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:45
				// _ = "end of CoverTab[104080]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:46
				_go_fuzz_dep_.CoverTab[104081]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:46
				// _ = "end of CoverTab[104081]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:46
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:46
			// _ = "end of CoverTab[104078]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:46
			_go_fuzz_dep_.CoverTab[104079]++
														r.Topics[i] = topic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:47
			// _ = "end of CoverTab[104079]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:48
		// _ = "end of CoverTab[104077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:49
		_go_fuzz_dep_.CoverTab[104082]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:49
		// _ = "end of CoverTab[104082]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:49
	// _ = "end of CoverTab[104072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:49
	_go_fuzz_dep_.CoverTab[104073]++
												if r.Version > 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:50
		_go_fuzz_dep_.CoverTab[104083]++
													autoCreation, err := pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:52
			_go_fuzz_dep_.CoverTab[104085]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:53
			// _ = "end of CoverTab[104085]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:54
			_go_fuzz_dep_.CoverTab[104086]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:54
			// _ = "end of CoverTab[104086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:54
		// _ = "end of CoverTab[104083]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:54
		_go_fuzz_dep_.CoverTab[104084]++
													r.AllowAutoTopicCreation = autoCreation
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:55
		// _ = "end of CoverTab[104084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:56
		_go_fuzz_dep_.CoverTab[104087]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:56
		// _ = "end of CoverTab[104087]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:56
	// _ = "end of CoverTab[104073]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:56
	_go_fuzz_dep_.CoverTab[104074]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:57
	// _ = "end of CoverTab[104074]"
}

func (r *MetadataRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:60
	_go_fuzz_dep_.CoverTab[104088]++
												return 3
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:61
	// _ = "end of CoverTab[104088]"
}

func (r *MetadataRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:64
	_go_fuzz_dep_.CoverTab[104089]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:65
	// _ = "end of CoverTab[104089]"
}

func (r *MetadataRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:68
	_go_fuzz_dep_.CoverTab[104090]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:69
	// _ = "end of CoverTab[104090]"
}

func (r *MetadataRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:72
	_go_fuzz_dep_.CoverTab[104091]++
												switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:74
		_go_fuzz_dep_.CoverTab[104092]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:75
		// _ = "end of CoverTab[104092]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:76
		_go_fuzz_dep_.CoverTab[104093]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:77
		// _ = "end of CoverTab[104093]"
	case 3, 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:78
		_go_fuzz_dep_.CoverTab[104094]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:79
		// _ = "end of CoverTab[104094]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:80
		_go_fuzz_dep_.CoverTab[104095]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:81
		// _ = "end of CoverTab[104095]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:82
		_go_fuzz_dep_.CoverTab[104096]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:83
		// _ = "end of CoverTab[104096]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:84
	// _ = "end of CoverTab[104091]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:85
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_request.go:85
var _ = _go_fuzz_dep_.CoverTab
