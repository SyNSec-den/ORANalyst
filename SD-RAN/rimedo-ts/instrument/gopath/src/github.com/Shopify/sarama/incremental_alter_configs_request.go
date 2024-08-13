//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:1
)

type IncrementalAlterConfigsOperation int8

const (
	IncrementalAlterConfigsOperationSet	IncrementalAlterConfigsOperation	= iota
	IncrementalAlterConfigsOperationDelete
	IncrementalAlterConfigsOperationAppend
	IncrementalAlterConfigsOperationSubtract
)

// IncrementalAlterConfigsRequest is an incremental alter config request type
type IncrementalAlterConfigsRequest struct {
	Resources	[]*IncrementalAlterConfigsResource
	ValidateOnly	bool
}

type IncrementalAlterConfigsResource struct {
	Type		ConfigResourceType
	Name		string
	ConfigEntries	map[string]IncrementalAlterConfigsEntry
}

type IncrementalAlterConfigsEntry struct {
	Operation	IncrementalAlterConfigsOperation
	Value		*string
}

func (a *IncrementalAlterConfigsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:29
	_go_fuzz_dep_.CoverTab[103424]++
														if err := pe.putArrayLength(len(a.Resources)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:30
		_go_fuzz_dep_.CoverTab[103427]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:31
		// _ = "end of CoverTab[103427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:32
		_go_fuzz_dep_.CoverTab[103428]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:32
		// _ = "end of CoverTab[103428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:32
	// _ = "end of CoverTab[103424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:32
	_go_fuzz_dep_.CoverTab[103425]++

														for _, r := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:34
		_go_fuzz_dep_.CoverTab[103429]++
															if err := r.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:35
			_go_fuzz_dep_.CoverTab[103430]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:36
			// _ = "end of CoverTab[103430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:37
			_go_fuzz_dep_.CoverTab[103431]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:37
			// _ = "end of CoverTab[103431]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:37
		// _ = "end of CoverTab[103429]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:38
	// _ = "end of CoverTab[103425]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:38
	_go_fuzz_dep_.CoverTab[103426]++

														pe.putBool(a.ValidateOnly)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:41
	// _ = "end of CoverTab[103426]"
}

func (a *IncrementalAlterConfigsRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:44
	_go_fuzz_dep_.CoverTab[103432]++
														resourceCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:46
		_go_fuzz_dep_.CoverTab[103436]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:47
		// _ = "end of CoverTab[103436]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:48
		_go_fuzz_dep_.CoverTab[103437]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:48
		// _ = "end of CoverTab[103437]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:48
	// _ = "end of CoverTab[103432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:48
	_go_fuzz_dep_.CoverTab[103433]++

														a.Resources = make([]*IncrementalAlterConfigsResource, resourceCount)
														for i := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:51
		_go_fuzz_dep_.CoverTab[103438]++
															r := &IncrementalAlterConfigsResource{}
															err = r.decode(pd, version)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:54
			_go_fuzz_dep_.CoverTab[103440]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:55
			// _ = "end of CoverTab[103440]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:56
			_go_fuzz_dep_.CoverTab[103441]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:56
			// _ = "end of CoverTab[103441]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:56
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:56
		// _ = "end of CoverTab[103438]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:56
		_go_fuzz_dep_.CoverTab[103439]++
															a.Resources[i] = r
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:57
		// _ = "end of CoverTab[103439]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:58
	// _ = "end of CoverTab[103433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:58
	_go_fuzz_dep_.CoverTab[103434]++

														validateOnly, err := pd.getBool()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:61
		_go_fuzz_dep_.CoverTab[103442]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:62
		// _ = "end of CoverTab[103442]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:63
		_go_fuzz_dep_.CoverTab[103443]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:63
		// _ = "end of CoverTab[103443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:63
	// _ = "end of CoverTab[103434]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:63
	_go_fuzz_dep_.CoverTab[103435]++

														a.ValidateOnly = validateOnly

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:67
	// _ = "end of CoverTab[103435]"
}

func (a *IncrementalAlterConfigsResource) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:70
	_go_fuzz_dep_.CoverTab[103444]++
														pe.putInt8(int8(a.Type))

														if err := pe.putString(a.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:73
		_go_fuzz_dep_.CoverTab[103448]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:74
		// _ = "end of CoverTab[103448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:75
		_go_fuzz_dep_.CoverTab[103449]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:75
		// _ = "end of CoverTab[103449]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:75
	// _ = "end of CoverTab[103444]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:75
	_go_fuzz_dep_.CoverTab[103445]++

														if err := pe.putArrayLength(len(a.ConfigEntries)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:77
		_go_fuzz_dep_.CoverTab[103450]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:78
		// _ = "end of CoverTab[103450]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:79
		_go_fuzz_dep_.CoverTab[103451]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:79
		// _ = "end of CoverTab[103451]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:79
	// _ = "end of CoverTab[103445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:79
	_go_fuzz_dep_.CoverTab[103446]++

														for name, e := range a.ConfigEntries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:81
		_go_fuzz_dep_.CoverTab[103452]++
															if err := pe.putString(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:82
			_go_fuzz_dep_.CoverTab[103454]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:83
			// _ = "end of CoverTab[103454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:84
			_go_fuzz_dep_.CoverTab[103455]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:84
			// _ = "end of CoverTab[103455]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:84
		// _ = "end of CoverTab[103452]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:84
		_go_fuzz_dep_.CoverTab[103453]++

															if err := e.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:86
			_go_fuzz_dep_.CoverTab[103456]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:87
			// _ = "end of CoverTab[103456]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:88
			_go_fuzz_dep_.CoverTab[103457]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:88
			// _ = "end of CoverTab[103457]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:88
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:88
		// _ = "end of CoverTab[103453]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:89
	// _ = "end of CoverTab[103446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:89
	_go_fuzz_dep_.CoverTab[103447]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:91
	// _ = "end of CoverTab[103447]"
}

func (a *IncrementalAlterConfigsResource) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:94
	_go_fuzz_dep_.CoverTab[103458]++
														t, err := pd.getInt8()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:96
		_go_fuzz_dep_.CoverTab[103463]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:97
		// _ = "end of CoverTab[103463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:98
		_go_fuzz_dep_.CoverTab[103464]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:98
		// _ = "end of CoverTab[103464]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:98
	// _ = "end of CoverTab[103458]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:98
	_go_fuzz_dep_.CoverTab[103459]++
														a.Type = ConfigResourceType(t)

														name, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:102
		_go_fuzz_dep_.CoverTab[103465]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:103
		// _ = "end of CoverTab[103465]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:104
		_go_fuzz_dep_.CoverTab[103466]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:104
		// _ = "end of CoverTab[103466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:104
	// _ = "end of CoverTab[103459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:104
	_go_fuzz_dep_.CoverTab[103460]++
														a.Name = name

														n, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:108
		_go_fuzz_dep_.CoverTab[103467]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:109
		// _ = "end of CoverTab[103467]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:110
		_go_fuzz_dep_.CoverTab[103468]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:110
		// _ = "end of CoverTab[103468]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:110
	// _ = "end of CoverTab[103460]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:110
	_go_fuzz_dep_.CoverTab[103461]++

														if n > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:112
		_go_fuzz_dep_.CoverTab[103469]++
															a.ConfigEntries = make(map[string]IncrementalAlterConfigsEntry, n)
															for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:114
			_go_fuzz_dep_.CoverTab[103470]++
																name, err := pd.getString()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:116
				_go_fuzz_dep_.CoverTab[103473]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:117
				// _ = "end of CoverTab[103473]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:118
				_go_fuzz_dep_.CoverTab[103474]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:118
				// _ = "end of CoverTab[103474]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:118
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:118
			// _ = "end of CoverTab[103470]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:118
			_go_fuzz_dep_.CoverTab[103471]++

																var v IncrementalAlterConfigsEntry

																if err := v.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:122
				_go_fuzz_dep_.CoverTab[103475]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:123
				// _ = "end of CoverTab[103475]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:124
				_go_fuzz_dep_.CoverTab[103476]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:124
				// _ = "end of CoverTab[103476]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:124
			// _ = "end of CoverTab[103471]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:124
			_go_fuzz_dep_.CoverTab[103472]++

																a.ConfigEntries[name] = v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:126
			// _ = "end of CoverTab[103472]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:127
		// _ = "end of CoverTab[103469]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:128
		_go_fuzz_dep_.CoverTab[103477]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:128
		// _ = "end of CoverTab[103477]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:128
	// _ = "end of CoverTab[103461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:128
	_go_fuzz_dep_.CoverTab[103462]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:129
	// _ = "end of CoverTab[103462]"
}

func (a *IncrementalAlterConfigsEntry) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:132
	_go_fuzz_dep_.CoverTab[103478]++
														pe.putInt8(int8(a.Operation))

														if err := pe.putNullableString(a.Value); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:135
		_go_fuzz_dep_.CoverTab[103480]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:136
		// _ = "end of CoverTab[103480]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:137
		_go_fuzz_dep_.CoverTab[103481]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:137
		// _ = "end of CoverTab[103481]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:137
	// _ = "end of CoverTab[103478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:137
	_go_fuzz_dep_.CoverTab[103479]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:139
	// _ = "end of CoverTab[103479]"
}

func (a *IncrementalAlterConfigsEntry) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:142
	_go_fuzz_dep_.CoverTab[103482]++
														t, err := pd.getInt8()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:144
		_go_fuzz_dep_.CoverTab[103485]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:145
		// _ = "end of CoverTab[103485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:146
		_go_fuzz_dep_.CoverTab[103486]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:146
		// _ = "end of CoverTab[103486]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:146
	// _ = "end of CoverTab[103482]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:146
	_go_fuzz_dep_.CoverTab[103483]++
														a.Operation = IncrementalAlterConfigsOperation(t)

														s, err := pd.getNullableString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:150
		_go_fuzz_dep_.CoverTab[103487]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:151
		// _ = "end of CoverTab[103487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:152
		_go_fuzz_dep_.CoverTab[103488]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:152
		// _ = "end of CoverTab[103488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:152
	// _ = "end of CoverTab[103483]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:152
	_go_fuzz_dep_.CoverTab[103484]++

														a.Value = s

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:156
	// _ = "end of CoverTab[103484]"
}

func (a *IncrementalAlterConfigsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:159
	_go_fuzz_dep_.CoverTab[103489]++
														return 44
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:160
	// _ = "end of CoverTab[103489]"
}

func (a *IncrementalAlterConfigsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:163
	_go_fuzz_dep_.CoverTab[103490]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:164
	// _ = "end of CoverTab[103490]"
}

func (a *IncrementalAlterConfigsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:167
	_go_fuzz_dep_.CoverTab[103491]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:168
	// _ = "end of CoverTab[103491]"
}

func (a *IncrementalAlterConfigsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:171
	_go_fuzz_dep_.CoverTab[103492]++
														return V2_3_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:172
	// _ = "end of CoverTab[103492]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:173
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_request.go:173
var _ = _go_fuzz_dep_.CoverTab
