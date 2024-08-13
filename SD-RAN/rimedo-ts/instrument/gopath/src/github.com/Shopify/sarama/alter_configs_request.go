//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:1
)

// AlterConfigsRequest is an alter config request type
type AlterConfigsRequest struct {
	Resources	[]*AlterConfigsResource
	ValidateOnly	bool
}

// AlterConfigsResource is an alter config resource type
type AlterConfigsResource struct {
	Type		ConfigResourceType
	Name		string
	ConfigEntries	map[string]*string
}

func (a *AlterConfigsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:16
	_go_fuzz_dep_.CoverTab[98107]++
													if err := pe.putArrayLength(len(a.Resources)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:17
		_go_fuzz_dep_.CoverTab[98110]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:18
		// _ = "end of CoverTab[98110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:19
		_go_fuzz_dep_.CoverTab[98111]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:19
		// _ = "end of CoverTab[98111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:19
	// _ = "end of CoverTab[98107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:19
	_go_fuzz_dep_.CoverTab[98108]++

													for _, r := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:21
		_go_fuzz_dep_.CoverTab[98112]++
														if err := r.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:22
			_go_fuzz_dep_.CoverTab[98113]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:23
			// _ = "end of CoverTab[98113]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:24
			_go_fuzz_dep_.CoverTab[98114]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:24
			// _ = "end of CoverTab[98114]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:24
		// _ = "end of CoverTab[98112]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:25
	// _ = "end of CoverTab[98108]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:25
	_go_fuzz_dep_.CoverTab[98109]++

													pe.putBool(a.ValidateOnly)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:28
	// _ = "end of CoverTab[98109]"
}

func (a *AlterConfigsRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:31
	_go_fuzz_dep_.CoverTab[98115]++
													resourceCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:33
		_go_fuzz_dep_.CoverTab[98119]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:34
		// _ = "end of CoverTab[98119]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:35
		_go_fuzz_dep_.CoverTab[98120]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:35
		// _ = "end of CoverTab[98120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:35
	// _ = "end of CoverTab[98115]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:35
	_go_fuzz_dep_.CoverTab[98116]++

													a.Resources = make([]*AlterConfigsResource, resourceCount)
													for i := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:38
		_go_fuzz_dep_.CoverTab[98121]++
														r := &AlterConfigsResource{}
														err = r.decode(pd, version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:41
			_go_fuzz_dep_.CoverTab[98123]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:42
			// _ = "end of CoverTab[98123]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:43
			_go_fuzz_dep_.CoverTab[98124]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:43
			// _ = "end of CoverTab[98124]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:43
		// _ = "end of CoverTab[98121]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:43
		_go_fuzz_dep_.CoverTab[98122]++
														a.Resources[i] = r
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:44
		// _ = "end of CoverTab[98122]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:45
	// _ = "end of CoverTab[98116]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:45
	_go_fuzz_dep_.CoverTab[98117]++

													validateOnly, err := pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:48
		_go_fuzz_dep_.CoverTab[98125]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:49
		// _ = "end of CoverTab[98125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:50
		_go_fuzz_dep_.CoverTab[98126]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:50
		// _ = "end of CoverTab[98126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:50
	// _ = "end of CoverTab[98117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:50
	_go_fuzz_dep_.CoverTab[98118]++

													a.ValidateOnly = validateOnly

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:54
	// _ = "end of CoverTab[98118]"
}

func (a *AlterConfigsResource) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:57
	_go_fuzz_dep_.CoverTab[98127]++
													pe.putInt8(int8(a.Type))

													if err := pe.putString(a.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:60
		_go_fuzz_dep_.CoverTab[98131]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:61
		// _ = "end of CoverTab[98131]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:62
		_go_fuzz_dep_.CoverTab[98132]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:62
		// _ = "end of CoverTab[98132]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:62
	// _ = "end of CoverTab[98127]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:62
	_go_fuzz_dep_.CoverTab[98128]++

													if err := pe.putArrayLength(len(a.ConfigEntries)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:64
		_go_fuzz_dep_.CoverTab[98133]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:65
		// _ = "end of CoverTab[98133]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:66
		_go_fuzz_dep_.CoverTab[98134]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:66
		// _ = "end of CoverTab[98134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:66
	// _ = "end of CoverTab[98128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:66
	_go_fuzz_dep_.CoverTab[98129]++
													for configKey, configValue := range a.ConfigEntries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:67
		_go_fuzz_dep_.CoverTab[98135]++
														if err := pe.putString(configKey); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:68
			_go_fuzz_dep_.CoverTab[98137]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:69
			// _ = "end of CoverTab[98137]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:70
			_go_fuzz_dep_.CoverTab[98138]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:70
			// _ = "end of CoverTab[98138]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:70
		// _ = "end of CoverTab[98135]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:70
		_go_fuzz_dep_.CoverTab[98136]++
														if err := pe.putNullableString(configValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:71
			_go_fuzz_dep_.CoverTab[98139]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:72
			// _ = "end of CoverTab[98139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:73
			_go_fuzz_dep_.CoverTab[98140]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:73
			// _ = "end of CoverTab[98140]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:73
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:73
		// _ = "end of CoverTab[98136]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:74
	// _ = "end of CoverTab[98129]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:74
	_go_fuzz_dep_.CoverTab[98130]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:76
	// _ = "end of CoverTab[98130]"
}

func (a *AlterConfigsResource) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:79
	_go_fuzz_dep_.CoverTab[98141]++
													t, err := pd.getInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:81
		_go_fuzz_dep_.CoverTab[98146]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:82
		// _ = "end of CoverTab[98146]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:83
		_go_fuzz_dep_.CoverTab[98147]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:83
		// _ = "end of CoverTab[98147]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:83
	// _ = "end of CoverTab[98141]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:83
	_go_fuzz_dep_.CoverTab[98142]++
													a.Type = ConfigResourceType(t)

													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:87
		_go_fuzz_dep_.CoverTab[98148]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:88
		// _ = "end of CoverTab[98148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:89
		_go_fuzz_dep_.CoverTab[98149]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:89
		// _ = "end of CoverTab[98149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:89
	// _ = "end of CoverTab[98142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:89
	_go_fuzz_dep_.CoverTab[98143]++
													a.Name = name

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:93
		_go_fuzz_dep_.CoverTab[98150]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:94
		// _ = "end of CoverTab[98150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:95
		_go_fuzz_dep_.CoverTab[98151]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:95
		// _ = "end of CoverTab[98151]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:95
	// _ = "end of CoverTab[98143]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:95
	_go_fuzz_dep_.CoverTab[98144]++

													if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:97
		_go_fuzz_dep_.CoverTab[98152]++
														a.ConfigEntries = make(map[string]*string, n)
														for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:99
			_go_fuzz_dep_.CoverTab[98153]++
															configKey, err := pd.getString()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:101
				_go_fuzz_dep_.CoverTab[98155]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:102
				// _ = "end of CoverTab[98155]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:103
				_go_fuzz_dep_.CoverTab[98156]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:103
				// _ = "end of CoverTab[98156]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:103
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:103
			// _ = "end of CoverTab[98153]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:103
			_go_fuzz_dep_.CoverTab[98154]++
															if a.ConfigEntries[configKey], err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:104
				_go_fuzz_dep_.CoverTab[98157]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:105
				// _ = "end of CoverTab[98157]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:106
				_go_fuzz_dep_.CoverTab[98158]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:106
				// _ = "end of CoverTab[98158]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:106
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:106
			// _ = "end of CoverTab[98154]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:107
		// _ = "end of CoverTab[98152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:108
		_go_fuzz_dep_.CoverTab[98159]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:108
		// _ = "end of CoverTab[98159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:108
	// _ = "end of CoverTab[98144]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:108
	_go_fuzz_dep_.CoverTab[98145]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:109
	// _ = "end of CoverTab[98145]"
}

func (a *AlterConfigsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:112
	_go_fuzz_dep_.CoverTab[98160]++
													return 33
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:113
	// _ = "end of CoverTab[98160]"
}

func (a *AlterConfigsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:116
	_go_fuzz_dep_.CoverTab[98161]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:117
	// _ = "end of CoverTab[98161]"
}

func (a *AlterConfigsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:120
	_go_fuzz_dep_.CoverTab[98162]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:121
	// _ = "end of CoverTab[98162]"
}

func (a *AlterConfigsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:124
	_go_fuzz_dep_.CoverTab[98163]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:125
	// _ = "end of CoverTab[98163]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:126
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_request.go:126
var _ = _go_fuzz_dep_.CoverTab
