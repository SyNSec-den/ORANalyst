//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:1
)

import (
	"fmt"
	"time"
)

type ConfigSource int8

func (s ConfigSource) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:10
	_go_fuzz_dep_.CoverTab[102211]++
													switch s {
	case SourceUnknown:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:12
		_go_fuzz_dep_.CoverTab[102213]++
														return "Unknown"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:13
		// _ = "end of CoverTab[102213]"
	case SourceTopic:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:14
		_go_fuzz_dep_.CoverTab[102214]++
														return "Topic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:15
		// _ = "end of CoverTab[102214]"
	case SourceDynamicBroker:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:16
		_go_fuzz_dep_.CoverTab[102215]++
														return "DynamicBroker"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:17
		// _ = "end of CoverTab[102215]"
	case SourceDynamicDefaultBroker:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:18
		_go_fuzz_dep_.CoverTab[102216]++
														return "DynamicDefaultBroker"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:19
		// _ = "end of CoverTab[102216]"
	case SourceStaticBroker:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:20
		_go_fuzz_dep_.CoverTab[102217]++
														return "StaticBroker"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:21
		// _ = "end of CoverTab[102217]"
	case SourceDefault:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:22
		_go_fuzz_dep_.CoverTab[102218]++
														return "Default"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:23
		// _ = "end of CoverTab[102218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:23
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:23
		_go_fuzz_dep_.CoverTab[102219]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:23
		// _ = "end of CoverTab[102219]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:24
	// _ = "end of CoverTab[102211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:24
	_go_fuzz_dep_.CoverTab[102212]++
													return fmt.Sprintf("Source Invalid: %d", int(s))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:25
	// _ = "end of CoverTab[102212]"
}

const (
	SourceUnknown	ConfigSource	= iota
	SourceTopic
	SourceDynamicBroker
	SourceDynamicDefaultBroker
	SourceStaticBroker
	SourceDefault
)

type DescribeConfigsResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	Resources	[]*ResourceResponse
}

type ResourceResponse struct {
	ErrorCode	int16
	ErrorMsg	string
	Type		ConfigResourceType
	Name		string
	Configs		[]*ConfigEntry
}

type ConfigEntry struct {
	Name		string
	Value		string
	ReadOnly	bool
	Default		bool
	Source		ConfigSource
	Sensitive	bool
	Synonyms	[]*ConfigSynonym
}

type ConfigSynonym struct {
	ConfigName	string
	ConfigValue	string
	Source		ConfigSource
}

func (r *DescribeConfigsResponse) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:67
	_go_fuzz_dep_.CoverTab[102220]++
													pe.putInt32(int32(r.ThrottleTime / time.Millisecond))
													if err = pe.putArrayLength(len(r.Resources)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:69
		_go_fuzz_dep_.CoverTab[102223]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:70
		// _ = "end of CoverTab[102223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:71
		_go_fuzz_dep_.CoverTab[102224]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:71
		// _ = "end of CoverTab[102224]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:71
	// _ = "end of CoverTab[102220]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:71
	_go_fuzz_dep_.CoverTab[102221]++

													for _, c := range r.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:73
		_go_fuzz_dep_.CoverTab[102225]++
														if err = c.encode(pe, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:74
			_go_fuzz_dep_.CoverTab[102226]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:75
			// _ = "end of CoverTab[102226]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:76
			_go_fuzz_dep_.CoverTab[102227]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:76
			// _ = "end of CoverTab[102227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:76
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:76
		// _ = "end of CoverTab[102225]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:77
	// _ = "end of CoverTab[102221]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:77
	_go_fuzz_dep_.CoverTab[102222]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:79
	// _ = "end of CoverTab[102222]"
}

func (r *DescribeConfigsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:82
	_go_fuzz_dep_.CoverTab[102228]++
													r.Version = version
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:85
		_go_fuzz_dep_.CoverTab[102232]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:86
		// _ = "end of CoverTab[102232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:87
		_go_fuzz_dep_.CoverTab[102233]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:87
		// _ = "end of CoverTab[102233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:87
	// _ = "end of CoverTab[102228]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:87
	_go_fuzz_dep_.CoverTab[102229]++
													r.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:91
		_go_fuzz_dep_.CoverTab[102234]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:92
		// _ = "end of CoverTab[102234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:93
		_go_fuzz_dep_.CoverTab[102235]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:93
		// _ = "end of CoverTab[102235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:93
	// _ = "end of CoverTab[102229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:93
	_go_fuzz_dep_.CoverTab[102230]++

													r.Resources = make([]*ResourceResponse, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:96
		_go_fuzz_dep_.CoverTab[102236]++
														rr := &ResourceResponse{}
														if err := rr.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:98
			_go_fuzz_dep_.CoverTab[102238]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:99
			// _ = "end of CoverTab[102238]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:100
			_go_fuzz_dep_.CoverTab[102239]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:100
			// _ = "end of CoverTab[102239]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:100
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:100
		// _ = "end of CoverTab[102236]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:100
		_go_fuzz_dep_.CoverTab[102237]++
														r.Resources[i] = rr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:101
		// _ = "end of CoverTab[102237]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:102
	// _ = "end of CoverTab[102230]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:102
	_go_fuzz_dep_.CoverTab[102231]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:104
	// _ = "end of CoverTab[102231]"
}

func (r *DescribeConfigsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:107
	_go_fuzz_dep_.CoverTab[102240]++
													return 32
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:108
	// _ = "end of CoverTab[102240]"
}

func (r *DescribeConfigsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:111
	_go_fuzz_dep_.CoverTab[102241]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:112
	// _ = "end of CoverTab[102241]"
}

func (r *DescribeConfigsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:115
	_go_fuzz_dep_.CoverTab[102242]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:116
	// _ = "end of CoverTab[102242]"
}

func (r *DescribeConfigsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:119
	_go_fuzz_dep_.CoverTab[102243]++
													switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:121
		_go_fuzz_dep_.CoverTab[102244]++
														return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:122
		// _ = "end of CoverTab[102244]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:123
		_go_fuzz_dep_.CoverTab[102245]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:124
		// _ = "end of CoverTab[102245]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:125
		_go_fuzz_dep_.CoverTab[102246]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:126
		// _ = "end of CoverTab[102246]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:127
	// _ = "end of CoverTab[102243]"
}

func (r *ResourceResponse) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:130
	_go_fuzz_dep_.CoverTab[102247]++
													pe.putInt16(r.ErrorCode)

													if err = pe.putString(r.ErrorMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:133
		_go_fuzz_dep_.CoverTab[102252]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:134
		// _ = "end of CoverTab[102252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:135
		_go_fuzz_dep_.CoverTab[102253]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:135
		// _ = "end of CoverTab[102253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:135
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:135
	// _ = "end of CoverTab[102247]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:135
	_go_fuzz_dep_.CoverTab[102248]++

													pe.putInt8(int8(r.Type))

													if err = pe.putString(r.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:139
		_go_fuzz_dep_.CoverTab[102254]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:140
		// _ = "end of CoverTab[102254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:141
		_go_fuzz_dep_.CoverTab[102255]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:141
		// _ = "end of CoverTab[102255]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:141
	// _ = "end of CoverTab[102248]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:141
	_go_fuzz_dep_.CoverTab[102249]++

													if err = pe.putArrayLength(len(r.Configs)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:143
		_go_fuzz_dep_.CoverTab[102256]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:144
		// _ = "end of CoverTab[102256]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:145
		_go_fuzz_dep_.CoverTab[102257]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:145
		// _ = "end of CoverTab[102257]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:145
	// _ = "end of CoverTab[102249]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:145
	_go_fuzz_dep_.CoverTab[102250]++

													for _, c := range r.Configs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:147
		_go_fuzz_dep_.CoverTab[102258]++
														if err = c.encode(pe, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:148
			_go_fuzz_dep_.CoverTab[102259]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:149
			// _ = "end of CoverTab[102259]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:150
			_go_fuzz_dep_.CoverTab[102260]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:150
			// _ = "end of CoverTab[102260]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:150
		// _ = "end of CoverTab[102258]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:151
	// _ = "end of CoverTab[102250]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:151
	_go_fuzz_dep_.CoverTab[102251]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:152
	// _ = "end of CoverTab[102251]"
}

func (r *ResourceResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:155
	_go_fuzz_dep_.CoverTab[102261]++
													ec, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:157
		_go_fuzz_dep_.CoverTab[102268]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:158
		// _ = "end of CoverTab[102268]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:159
		_go_fuzz_dep_.CoverTab[102269]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:159
		// _ = "end of CoverTab[102269]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:159
	// _ = "end of CoverTab[102261]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:159
	_go_fuzz_dep_.CoverTab[102262]++
													r.ErrorCode = ec

													em, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:163
		_go_fuzz_dep_.CoverTab[102270]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:164
		// _ = "end of CoverTab[102270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:165
		_go_fuzz_dep_.CoverTab[102271]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:165
		// _ = "end of CoverTab[102271]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:165
	// _ = "end of CoverTab[102262]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:165
	_go_fuzz_dep_.CoverTab[102263]++
													r.ErrorMsg = em

													t, err := pd.getInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:169
		_go_fuzz_dep_.CoverTab[102272]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:170
		// _ = "end of CoverTab[102272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:171
		_go_fuzz_dep_.CoverTab[102273]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:171
		// _ = "end of CoverTab[102273]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:171
	// _ = "end of CoverTab[102263]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:171
	_go_fuzz_dep_.CoverTab[102264]++
													r.Type = ConfigResourceType(t)

													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:175
		_go_fuzz_dep_.CoverTab[102274]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:176
		// _ = "end of CoverTab[102274]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:177
		_go_fuzz_dep_.CoverTab[102275]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:177
		// _ = "end of CoverTab[102275]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:177
	// _ = "end of CoverTab[102264]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:177
	_go_fuzz_dep_.CoverTab[102265]++
													r.Name = name

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:181
		_go_fuzz_dep_.CoverTab[102276]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:182
		// _ = "end of CoverTab[102276]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:183
		_go_fuzz_dep_.CoverTab[102277]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:183
		// _ = "end of CoverTab[102277]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:183
	// _ = "end of CoverTab[102265]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:183
	_go_fuzz_dep_.CoverTab[102266]++

													r.Configs = make([]*ConfigEntry, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:186
		_go_fuzz_dep_.CoverTab[102278]++
														c := &ConfigEntry{}
														if err := c.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:188
			_go_fuzz_dep_.CoverTab[102280]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:189
			// _ = "end of CoverTab[102280]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:190
			_go_fuzz_dep_.CoverTab[102281]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:190
			// _ = "end of CoverTab[102281]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:190
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:190
		// _ = "end of CoverTab[102278]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:190
		_go_fuzz_dep_.CoverTab[102279]++
														r.Configs[i] = c
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:191
		// _ = "end of CoverTab[102279]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:192
	// _ = "end of CoverTab[102266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:192
	_go_fuzz_dep_.CoverTab[102267]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:193
	// _ = "end of CoverTab[102267]"
}

func (r *ConfigEntry) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:196
	_go_fuzz_dep_.CoverTab[102282]++
													if err = pe.putString(r.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:197
		_go_fuzz_dep_.CoverTab[102286]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:198
		// _ = "end of CoverTab[102286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:199
		_go_fuzz_dep_.CoverTab[102287]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:199
		// _ = "end of CoverTab[102287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:199
	// _ = "end of CoverTab[102282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:199
	_go_fuzz_dep_.CoverTab[102283]++

													if err = pe.putString(r.Value); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:201
		_go_fuzz_dep_.CoverTab[102288]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:202
		// _ = "end of CoverTab[102288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:203
		_go_fuzz_dep_.CoverTab[102289]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:203
		// _ = "end of CoverTab[102289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:203
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:203
	// _ = "end of CoverTab[102283]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:203
	_go_fuzz_dep_.CoverTab[102284]++

													pe.putBool(r.ReadOnly)

													if version <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:207
		_go_fuzz_dep_.CoverTab[102290]++
														pe.putBool(r.Default)
														pe.putBool(r.Sensitive)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:209
		// _ = "end of CoverTab[102290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:210
		_go_fuzz_dep_.CoverTab[102291]++
														pe.putInt8(int8(r.Source))
														pe.putBool(r.Sensitive)

														if err := pe.putArrayLength(len(r.Synonyms)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:214
			_go_fuzz_dep_.CoverTab[102293]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:215
			// _ = "end of CoverTab[102293]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:216
			_go_fuzz_dep_.CoverTab[102294]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:216
			// _ = "end of CoverTab[102294]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:216
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:216
		// _ = "end of CoverTab[102291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:216
		_go_fuzz_dep_.CoverTab[102292]++
														for _, c := range r.Synonyms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:217
			_go_fuzz_dep_.CoverTab[102295]++
															if err = c.encode(pe, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:218
				_go_fuzz_dep_.CoverTab[102296]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:219
				// _ = "end of CoverTab[102296]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:220
				_go_fuzz_dep_.CoverTab[102297]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:220
				// _ = "end of CoverTab[102297]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:220
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:220
			// _ = "end of CoverTab[102295]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:221
		// _ = "end of CoverTab[102292]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:222
	// _ = "end of CoverTab[102284]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:222
	_go_fuzz_dep_.CoverTab[102285]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:224
	// _ = "end of CoverTab[102285]"
}

// https://cwiki.apache.org/confluence/display/KAFKA/KIP-226+-+Dynamic+Broker+Configuration
func (r *ConfigEntry) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:228
	_go_fuzz_dep_.CoverTab[102298]++
													if version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:229
		_go_fuzz_dep_.CoverTab[102306]++
														r.Source = SourceUnknown
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:230
		// _ = "end of CoverTab[102306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:231
		_go_fuzz_dep_.CoverTab[102307]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:231
		// _ = "end of CoverTab[102307]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:231
	// _ = "end of CoverTab[102298]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:231
	_go_fuzz_dep_.CoverTab[102299]++
													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:233
		_go_fuzz_dep_.CoverTab[102308]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:234
		// _ = "end of CoverTab[102308]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:235
		_go_fuzz_dep_.CoverTab[102309]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:235
		// _ = "end of CoverTab[102309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:235
	// _ = "end of CoverTab[102299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:235
	_go_fuzz_dep_.CoverTab[102300]++
													r.Name = name

													value, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:239
		_go_fuzz_dep_.CoverTab[102310]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:240
		// _ = "end of CoverTab[102310]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:241
		_go_fuzz_dep_.CoverTab[102311]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:241
		// _ = "end of CoverTab[102311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:241
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:241
	// _ = "end of CoverTab[102300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:241
	_go_fuzz_dep_.CoverTab[102301]++
													r.Value = value

													read, err := pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:245
		_go_fuzz_dep_.CoverTab[102312]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:246
		// _ = "end of CoverTab[102312]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:247
		_go_fuzz_dep_.CoverTab[102313]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:247
		// _ = "end of CoverTab[102313]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:247
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:247
	// _ = "end of CoverTab[102301]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:247
	_go_fuzz_dep_.CoverTab[102302]++
													r.ReadOnly = read

													if version == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:250
		_go_fuzz_dep_.CoverTab[102314]++
														defaultB, err := pd.getBool()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:252
			_go_fuzz_dep_.CoverTab[102316]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:253
			// _ = "end of CoverTab[102316]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:254
			_go_fuzz_dep_.CoverTab[102317]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:254
			// _ = "end of CoverTab[102317]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:254
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:254
		// _ = "end of CoverTab[102314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:254
		_go_fuzz_dep_.CoverTab[102315]++
														r.Default = defaultB
														if defaultB {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:256
			_go_fuzz_dep_.CoverTab[102318]++
															r.Source = SourceDefault
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:257
			// _ = "end of CoverTab[102318]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:258
			_go_fuzz_dep_.CoverTab[102319]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:258
			// _ = "end of CoverTab[102319]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:258
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:258
		// _ = "end of CoverTab[102315]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:259
		_go_fuzz_dep_.CoverTab[102320]++
														source, err := pd.getInt8()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:261
			_go_fuzz_dep_.CoverTab[102322]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:262
			// _ = "end of CoverTab[102322]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:263
			_go_fuzz_dep_.CoverTab[102323]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:263
			// _ = "end of CoverTab[102323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:263
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:263
		// _ = "end of CoverTab[102320]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:263
		_go_fuzz_dep_.CoverTab[102321]++
														r.Source = ConfigSource(source)
														r.Default = r.Source == SourceDefault
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:265
		// _ = "end of CoverTab[102321]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:266
	// _ = "end of CoverTab[102302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:266
	_go_fuzz_dep_.CoverTab[102303]++

													sensitive, err := pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:269
		_go_fuzz_dep_.CoverTab[102324]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:270
		// _ = "end of CoverTab[102324]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:271
		_go_fuzz_dep_.CoverTab[102325]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:271
		// _ = "end of CoverTab[102325]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:271
	// _ = "end of CoverTab[102303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:271
	_go_fuzz_dep_.CoverTab[102304]++
													r.Sensitive = sensitive

													if version > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:274
		_go_fuzz_dep_.CoverTab[102326]++
														n, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:276
			_go_fuzz_dep_.CoverTab[102328]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:277
			// _ = "end of CoverTab[102328]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:278
			_go_fuzz_dep_.CoverTab[102329]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:278
			// _ = "end of CoverTab[102329]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:278
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:278
		// _ = "end of CoverTab[102326]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:278
		_go_fuzz_dep_.CoverTab[102327]++
														r.Synonyms = make([]*ConfigSynonym, n)

														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:281
			_go_fuzz_dep_.CoverTab[102330]++
															s := &ConfigSynonym{}
															if err := s.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:283
				_go_fuzz_dep_.CoverTab[102332]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:284
				// _ = "end of CoverTab[102332]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:285
				_go_fuzz_dep_.CoverTab[102333]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:285
				// _ = "end of CoverTab[102333]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:285
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:285
			// _ = "end of CoverTab[102330]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:285
			_go_fuzz_dep_.CoverTab[102331]++
															r.Synonyms[i] = s
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:286
			// _ = "end of CoverTab[102331]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:287
		// _ = "end of CoverTab[102327]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:288
		_go_fuzz_dep_.CoverTab[102334]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:288
		// _ = "end of CoverTab[102334]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:288
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:288
	// _ = "end of CoverTab[102304]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:288
	_go_fuzz_dep_.CoverTab[102305]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:289
	// _ = "end of CoverTab[102305]"
}

func (c *ConfigSynonym) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:292
	_go_fuzz_dep_.CoverTab[102335]++
													err = pe.putString(c.ConfigName)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:294
		_go_fuzz_dep_.CoverTab[102338]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:295
		// _ = "end of CoverTab[102338]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:296
		_go_fuzz_dep_.CoverTab[102339]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:296
		// _ = "end of CoverTab[102339]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:296
	// _ = "end of CoverTab[102335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:296
	_go_fuzz_dep_.CoverTab[102336]++

													err = pe.putString(c.ConfigValue)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:299
		_go_fuzz_dep_.CoverTab[102340]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:300
		// _ = "end of CoverTab[102340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:301
		_go_fuzz_dep_.CoverTab[102341]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:301
		// _ = "end of CoverTab[102341]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:301
	// _ = "end of CoverTab[102336]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:301
	_go_fuzz_dep_.CoverTab[102337]++

													pe.putInt8(int8(c.Source))

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:305
	// _ = "end of CoverTab[102337]"
}

func (c *ConfigSynonym) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:308
	_go_fuzz_dep_.CoverTab[102342]++
													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:310
		_go_fuzz_dep_.CoverTab[102346]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:311
		// _ = "end of CoverTab[102346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:312
		_go_fuzz_dep_.CoverTab[102347]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:312
		// _ = "end of CoverTab[102347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:312
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:312
	// _ = "end of CoverTab[102342]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:312
	_go_fuzz_dep_.CoverTab[102343]++
													c.ConfigName = name

													value, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:316
		_go_fuzz_dep_.CoverTab[102348]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:317
		// _ = "end of CoverTab[102348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:318
		_go_fuzz_dep_.CoverTab[102349]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:318
		// _ = "end of CoverTab[102349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:318
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:318
	// _ = "end of CoverTab[102343]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:318
	_go_fuzz_dep_.CoverTab[102344]++
													c.ConfigValue = value

													source, err := pd.getInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:322
		_go_fuzz_dep_.CoverTab[102350]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:323
		// _ = "end of CoverTab[102350]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:324
		_go_fuzz_dep_.CoverTab[102351]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:324
		// _ = "end of CoverTab[102351]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:324
	// _ = "end of CoverTab[102344]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:324
	_go_fuzz_dep_.CoverTab[102345]++
													c.Source = ConfigSource(source)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:326
	// _ = "end of CoverTab[102345]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:327
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_response.go:327
var _ = _go_fuzz_dep_.CoverTab
