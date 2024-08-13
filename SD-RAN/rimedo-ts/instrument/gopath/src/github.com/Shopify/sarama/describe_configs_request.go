//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:1
)

type DescribeConfigsRequest struct {
	Version		int16
	Resources	[]*ConfigResource
	IncludeSynonyms	bool
}

type ConfigResource struct {
	Type		ConfigResourceType
	Name		string
	ConfigNames	[]string
}

func (r *DescribeConfigsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:15
	_go_fuzz_dep_.CoverTab[102158]++
													if err := pe.putArrayLength(len(r.Resources)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:16
		_go_fuzz_dep_.CoverTab[102162]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:17
		// _ = "end of CoverTab[102162]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:18
		_go_fuzz_dep_.CoverTab[102163]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:18
		// _ = "end of CoverTab[102163]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:18
	// _ = "end of CoverTab[102158]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:18
	_go_fuzz_dep_.CoverTab[102159]++

													for _, c := range r.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:20
		_go_fuzz_dep_.CoverTab[102164]++
														pe.putInt8(int8(c.Type))
														if err := pe.putString(c.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:22
			_go_fuzz_dep_.CoverTab[102167]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:23
			// _ = "end of CoverTab[102167]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:24
			_go_fuzz_dep_.CoverTab[102168]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:24
			// _ = "end of CoverTab[102168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:24
		// _ = "end of CoverTab[102164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:24
		_go_fuzz_dep_.CoverTab[102165]++

														if len(c.ConfigNames) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:26
			_go_fuzz_dep_.CoverTab[102169]++
															pe.putInt32(-1)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:28
			// _ = "end of CoverTab[102169]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:29
			_go_fuzz_dep_.CoverTab[102170]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:29
			// _ = "end of CoverTab[102170]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:29
		// _ = "end of CoverTab[102165]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:29
		_go_fuzz_dep_.CoverTab[102166]++
														if err := pe.putStringArray(c.ConfigNames); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:30
			_go_fuzz_dep_.CoverTab[102171]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:31
			// _ = "end of CoverTab[102171]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:32
			_go_fuzz_dep_.CoverTab[102172]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:32
			// _ = "end of CoverTab[102172]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:32
		// _ = "end of CoverTab[102166]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:33
	// _ = "end of CoverTab[102159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:33
	_go_fuzz_dep_.CoverTab[102160]++

													if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:35
		_go_fuzz_dep_.CoverTab[102173]++
														pe.putBool(r.IncludeSynonyms)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:36
		// _ = "end of CoverTab[102173]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:37
		_go_fuzz_dep_.CoverTab[102174]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:37
		// _ = "end of CoverTab[102174]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:37
	// _ = "end of CoverTab[102160]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:37
	_go_fuzz_dep_.CoverTab[102161]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:39
	// _ = "end of CoverTab[102161]"
}

func (r *DescribeConfigsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:42
	_go_fuzz_dep_.CoverTab[102175]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:44
		_go_fuzz_dep_.CoverTab[102179]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:45
		// _ = "end of CoverTab[102179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:46
		_go_fuzz_dep_.CoverTab[102180]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:46
		// _ = "end of CoverTab[102180]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:46
	// _ = "end of CoverTab[102175]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:46
	_go_fuzz_dep_.CoverTab[102176]++

													r.Resources = make([]*ConfigResource, n)

													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:50
		_go_fuzz_dep_.CoverTab[102181]++
														r.Resources[i] = &ConfigResource{}
														t, err := pd.getInt8()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:53
			_go_fuzz_dep_.CoverTab[102187]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:54
			// _ = "end of CoverTab[102187]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:55
			_go_fuzz_dep_.CoverTab[102188]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:55
			// _ = "end of CoverTab[102188]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:55
		// _ = "end of CoverTab[102181]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:55
		_go_fuzz_dep_.CoverTab[102182]++
														r.Resources[i].Type = ConfigResourceType(t)
														name, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:58
			_go_fuzz_dep_.CoverTab[102189]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:59
			// _ = "end of CoverTab[102189]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:60
			_go_fuzz_dep_.CoverTab[102190]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:60
			// _ = "end of CoverTab[102190]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:60
		// _ = "end of CoverTab[102182]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:60
		_go_fuzz_dep_.CoverTab[102183]++
														r.Resources[i].Name = name

														confLength, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:64
			_go_fuzz_dep_.CoverTab[102191]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:65
			// _ = "end of CoverTab[102191]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:66
			_go_fuzz_dep_.CoverTab[102192]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:66
			// _ = "end of CoverTab[102192]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:66
		// _ = "end of CoverTab[102183]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:66
		_go_fuzz_dep_.CoverTab[102184]++

														if confLength == -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:68
			_go_fuzz_dep_.CoverTab[102193]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:69
			// _ = "end of CoverTab[102193]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:70
			_go_fuzz_dep_.CoverTab[102194]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:70
			// _ = "end of CoverTab[102194]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:70
		// _ = "end of CoverTab[102184]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:70
		_go_fuzz_dep_.CoverTab[102185]++

														cfnames := make([]string, confLength)
														for i := 0; i < confLength; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:73
			_go_fuzz_dep_.CoverTab[102195]++
															s, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:75
				_go_fuzz_dep_.CoverTab[102197]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:76
				// _ = "end of CoverTab[102197]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:77
				_go_fuzz_dep_.CoverTab[102198]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:77
				// _ = "end of CoverTab[102198]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:77
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:77
			// _ = "end of CoverTab[102195]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:77
			_go_fuzz_dep_.CoverTab[102196]++
															cfnames[i] = s
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:78
			// _ = "end of CoverTab[102196]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:79
		// _ = "end of CoverTab[102185]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:79
		_go_fuzz_dep_.CoverTab[102186]++
														r.Resources[i].ConfigNames = cfnames
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:80
		// _ = "end of CoverTab[102186]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:81
	// _ = "end of CoverTab[102176]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:81
	_go_fuzz_dep_.CoverTab[102177]++
													r.Version = version
													if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:83
		_go_fuzz_dep_.CoverTab[102199]++
														b, err := pd.getBool()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:85
			_go_fuzz_dep_.CoverTab[102201]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:86
			// _ = "end of CoverTab[102201]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:87
			_go_fuzz_dep_.CoverTab[102202]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:87
			// _ = "end of CoverTab[102202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:87
		// _ = "end of CoverTab[102199]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:87
		_go_fuzz_dep_.CoverTab[102200]++
														r.IncludeSynonyms = b
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:88
		// _ = "end of CoverTab[102200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:89
		_go_fuzz_dep_.CoverTab[102203]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:89
		// _ = "end of CoverTab[102203]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:89
	// _ = "end of CoverTab[102177]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:89
	_go_fuzz_dep_.CoverTab[102178]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:91
	// _ = "end of CoverTab[102178]"
}

func (r *DescribeConfigsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:94
	_go_fuzz_dep_.CoverTab[102204]++
													return 32
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:95
	// _ = "end of CoverTab[102204]"
}

func (r *DescribeConfigsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:98
	_go_fuzz_dep_.CoverTab[102205]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:99
	// _ = "end of CoverTab[102205]"
}

func (r *DescribeConfigsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:102
	_go_fuzz_dep_.CoverTab[102206]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:103
	// _ = "end of CoverTab[102206]"
}

func (r *DescribeConfigsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:106
	_go_fuzz_dep_.CoverTab[102207]++
													switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:108
		_go_fuzz_dep_.CoverTab[102208]++
														return V1_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:109
		// _ = "end of CoverTab[102208]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:110
		_go_fuzz_dep_.CoverTab[102209]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:111
		// _ = "end of CoverTab[102209]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:112
		_go_fuzz_dep_.CoverTab[102210]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:113
		// _ = "end of CoverTab[102210]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:114
	// _ = "end of CoverTab[102207]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:115
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_configs_request.go:115
var _ = _go_fuzz_dep_.CoverTab
