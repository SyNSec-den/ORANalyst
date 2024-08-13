//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:1
)

// ApiVersionsResponseKey contains the APIs supported by the broker.
type ApiVersionsResponseKey struct {
	// Version defines the protocol version to use for encode and decode
	Version	int16
	// ApiKey contains the API index.
	ApiKey	int16
	// MinVersion contains the minimum supported version, inclusive.
	MinVersion	int16
	// MaxVersion contains the maximum supported version, inclusive.
	MaxVersion	int16
}

func (a *ApiVersionsResponseKey) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:15
	_go_fuzz_dep_.CoverTab[98472]++
													a.Version = version
													pe.putInt16(a.ApiKey)

													pe.putInt16(a.MinVersion)

													pe.putInt16(a.MaxVersion)

													if version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:23
		_go_fuzz_dep_.CoverTab[98474]++
														pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:24
		// _ = "end of CoverTab[98474]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:25
		_go_fuzz_dep_.CoverTab[98475]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:25
		// _ = "end of CoverTab[98475]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:25
	// _ = "end of CoverTab[98472]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:25
	_go_fuzz_dep_.CoverTab[98473]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:27
	// _ = "end of CoverTab[98473]"
}

func (a *ApiVersionsResponseKey) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:30
	_go_fuzz_dep_.CoverTab[98476]++
													a.Version = version
													if a.ApiKey, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:32
		_go_fuzz_dep_.CoverTab[98481]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:33
		// _ = "end of CoverTab[98481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:34
		_go_fuzz_dep_.CoverTab[98482]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:34
		// _ = "end of CoverTab[98482]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:34
	// _ = "end of CoverTab[98476]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:34
	_go_fuzz_dep_.CoverTab[98477]++

													if a.MinVersion, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:36
		_go_fuzz_dep_.CoverTab[98483]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:37
		// _ = "end of CoverTab[98483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:38
		_go_fuzz_dep_.CoverTab[98484]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:38
		// _ = "end of CoverTab[98484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:38
	// _ = "end of CoverTab[98477]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:38
	_go_fuzz_dep_.CoverTab[98478]++

													if a.MaxVersion, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:40
		_go_fuzz_dep_.CoverTab[98485]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:41
		// _ = "end of CoverTab[98485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:42
		_go_fuzz_dep_.CoverTab[98486]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:42
		// _ = "end of CoverTab[98486]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:42
	// _ = "end of CoverTab[98478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:42
	_go_fuzz_dep_.CoverTab[98479]++

													if version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:44
		_go_fuzz_dep_.CoverTab[98487]++
														if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:45
			_go_fuzz_dep_.CoverTab[98488]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:46
			// _ = "end of CoverTab[98488]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:47
			_go_fuzz_dep_.CoverTab[98489]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:47
			// _ = "end of CoverTab[98489]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:47
		// _ = "end of CoverTab[98487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:48
		_go_fuzz_dep_.CoverTab[98490]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:48
		// _ = "end of CoverTab[98490]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:48
	// _ = "end of CoverTab[98479]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:48
	_go_fuzz_dep_.CoverTab[98480]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:50
	// _ = "end of CoverTab[98480]"
}

type ApiVersionsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version	int16
	// ErrorCode contains the top-level error code.
	ErrorCode	int16
	// ApiKeys contains the APIs supported by the broker.
	ApiKeys	[]ApiVersionsResponseKey
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs	int32
}

func (r *ApiVersionsResponse) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:64
	_go_fuzz_dep_.CoverTab[98491]++
													pe.putInt16(r.ErrorCode)

													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:67
		_go_fuzz_dep_.CoverTab[98496]++
														pe.putCompactArrayLength(len(r.ApiKeys))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:68
		// _ = "end of CoverTab[98496]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:69
		_go_fuzz_dep_.CoverTab[98497]++
														if err := pe.putArrayLength(len(r.ApiKeys)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:70
			_go_fuzz_dep_.CoverTab[98498]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:71
			// _ = "end of CoverTab[98498]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:72
			_go_fuzz_dep_.CoverTab[98499]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:72
			// _ = "end of CoverTab[98499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:72
		// _ = "end of CoverTab[98497]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:73
	// _ = "end of CoverTab[98491]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:73
	_go_fuzz_dep_.CoverTab[98492]++
													for _, block := range r.ApiKeys {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:74
		_go_fuzz_dep_.CoverTab[98500]++
														if err := block.encode(pe, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:75
			_go_fuzz_dep_.CoverTab[98501]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:76
			// _ = "end of CoverTab[98501]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:77
			_go_fuzz_dep_.CoverTab[98502]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:77
			// _ = "end of CoverTab[98502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:77
		// _ = "end of CoverTab[98500]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:78
	// _ = "end of CoverTab[98492]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:78
	_go_fuzz_dep_.CoverTab[98493]++

													if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:80
		_go_fuzz_dep_.CoverTab[98503]++
														pe.putInt32(r.ThrottleTimeMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:81
		// _ = "end of CoverTab[98503]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:82
		_go_fuzz_dep_.CoverTab[98504]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:82
		// _ = "end of CoverTab[98504]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:82
	// _ = "end of CoverTab[98493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:82
	_go_fuzz_dep_.CoverTab[98494]++

													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:84
		_go_fuzz_dep_.CoverTab[98505]++
														pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:85
		// _ = "end of CoverTab[98505]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:86
		_go_fuzz_dep_.CoverTab[98506]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:86
		// _ = "end of CoverTab[98506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:86
	// _ = "end of CoverTab[98494]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:86
	_go_fuzz_dep_.CoverTab[98495]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:88
	// _ = "end of CoverTab[98495]"
}

func (r *ApiVersionsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:91
	_go_fuzz_dep_.CoverTab[98507]++
													r.Version = version
													if r.ErrorCode, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:93
		_go_fuzz_dep_.CoverTab[98513]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:94
		// _ = "end of CoverTab[98513]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:95
		_go_fuzz_dep_.CoverTab[98514]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:95
		// _ = "end of CoverTab[98514]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:95
	// _ = "end of CoverTab[98507]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:95
	_go_fuzz_dep_.CoverTab[98508]++

													var numApiKeys int
													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:98
		_go_fuzz_dep_.CoverTab[98515]++
														numApiKeys, err = pd.getCompactArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:100
			_go_fuzz_dep_.CoverTab[98516]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:101
			// _ = "end of CoverTab[98516]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:102
			_go_fuzz_dep_.CoverTab[98517]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:102
			// _ = "end of CoverTab[98517]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:102
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:102
		// _ = "end of CoverTab[98515]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:103
		_go_fuzz_dep_.CoverTab[98518]++
														numApiKeys, err = pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:105
			_go_fuzz_dep_.CoverTab[98519]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:106
			// _ = "end of CoverTab[98519]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:107
			_go_fuzz_dep_.CoverTab[98520]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:107
			// _ = "end of CoverTab[98520]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:107
		// _ = "end of CoverTab[98518]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:108
	// _ = "end of CoverTab[98508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:108
	_go_fuzz_dep_.CoverTab[98509]++
													r.ApiKeys = make([]ApiVersionsResponseKey, numApiKeys)
													for i := 0; i < numApiKeys; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:110
		_go_fuzz_dep_.CoverTab[98521]++
														var block ApiVersionsResponseKey
														if err = block.decode(pd, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:112
			_go_fuzz_dep_.CoverTab[98523]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:113
			// _ = "end of CoverTab[98523]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:114
			_go_fuzz_dep_.CoverTab[98524]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:114
			// _ = "end of CoverTab[98524]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:114
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:114
		// _ = "end of CoverTab[98521]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:114
		_go_fuzz_dep_.CoverTab[98522]++
														r.ApiKeys[i] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:115
		// _ = "end of CoverTab[98522]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:116
	// _ = "end of CoverTab[98509]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:116
	_go_fuzz_dep_.CoverTab[98510]++

													if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:118
		_go_fuzz_dep_.CoverTab[98525]++
														if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:119
			_go_fuzz_dep_.CoverTab[98526]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:120
			// _ = "end of CoverTab[98526]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:121
			_go_fuzz_dep_.CoverTab[98527]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:121
			// _ = "end of CoverTab[98527]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:121
		// _ = "end of CoverTab[98525]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:122
		_go_fuzz_dep_.CoverTab[98528]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:122
		// _ = "end of CoverTab[98528]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:122
	// _ = "end of CoverTab[98510]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:122
	_go_fuzz_dep_.CoverTab[98511]++

													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:124
		_go_fuzz_dep_.CoverTab[98529]++
														if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:125
			_go_fuzz_dep_.CoverTab[98530]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:126
			// _ = "end of CoverTab[98530]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:127
			_go_fuzz_dep_.CoverTab[98531]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:127
			// _ = "end of CoverTab[98531]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:127
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:127
		// _ = "end of CoverTab[98529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:128
		_go_fuzz_dep_.CoverTab[98532]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:128
		// _ = "end of CoverTab[98532]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:128
	// _ = "end of CoverTab[98511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:128
	_go_fuzz_dep_.CoverTab[98512]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:130
	// _ = "end of CoverTab[98512]"
}

func (r *ApiVersionsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:133
	_go_fuzz_dep_.CoverTab[98533]++
													return 18
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:134
	// _ = "end of CoverTab[98533]"
}

func (r *ApiVersionsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:137
	_go_fuzz_dep_.CoverTab[98534]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:138
	// _ = "end of CoverTab[98534]"
}

func (r *ApiVersionsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:141
	_go_fuzz_dep_.CoverTab[98535]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:144
	return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:144
	// _ = "end of CoverTab[98535]"
}

func (r *ApiVersionsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:147
	_go_fuzz_dep_.CoverTab[98536]++
													switch r.Version {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:149
		_go_fuzz_dep_.CoverTab[98537]++
														return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:150
		// _ = "end of CoverTab[98537]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:151
		_go_fuzz_dep_.CoverTab[98538]++
														return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:152
		// _ = "end of CoverTab[98538]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:153
		_go_fuzz_dep_.CoverTab[98539]++
														return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:154
		// _ = "end of CoverTab[98539]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:155
	// _ = "end of CoverTab[98536]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:156
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_response.go:156
var _ = _go_fuzz_dep_.CoverTab
