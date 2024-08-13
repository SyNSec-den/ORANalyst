//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:1
)

import "time"

// AlterConfigsResponse is a response type for alter config
type AlterConfigsResponse struct {
	ThrottleTime	time.Duration
	Resources	[]*AlterConfigsResourceResponse
}

// AlterConfigsResourceResponse is a response type for alter config resource
type AlterConfigsResourceResponse struct {
	ErrorCode	int16
	ErrorMsg	string
	Type		ConfigResourceType
	Name		string
}

func (a *AlterConfigsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:19
	_go_fuzz_dep_.CoverTab[98164]++
													pe.putInt32(int32(a.ThrottleTime / time.Millisecond))

													if err := pe.putArrayLength(len(a.Resources)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:22
		_go_fuzz_dep_.CoverTab[98167]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:23
		// _ = "end of CoverTab[98167]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:24
		_go_fuzz_dep_.CoverTab[98168]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:24
		// _ = "end of CoverTab[98168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:24
	// _ = "end of CoverTab[98164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:24
	_go_fuzz_dep_.CoverTab[98165]++

													for _, v := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:26
		_go_fuzz_dep_.CoverTab[98169]++
														if err := v.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:27
			_go_fuzz_dep_.CoverTab[98170]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:28
			// _ = "end of CoverTab[98170]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:29
			_go_fuzz_dep_.CoverTab[98171]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:29
			// _ = "end of CoverTab[98171]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:29
		// _ = "end of CoverTab[98169]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:30
	// _ = "end of CoverTab[98165]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:30
	_go_fuzz_dep_.CoverTab[98166]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:32
	// _ = "end of CoverTab[98166]"
}

func (a *AlterConfigsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:35
	_go_fuzz_dep_.CoverTab[98172]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:37
		_go_fuzz_dep_.CoverTab[98176]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:38
		// _ = "end of CoverTab[98176]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:39
		_go_fuzz_dep_.CoverTab[98177]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:39
		// _ = "end of CoverTab[98177]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:39
	// _ = "end of CoverTab[98172]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:39
	_go_fuzz_dep_.CoverTab[98173]++
													a.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													responseCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:43
		_go_fuzz_dep_.CoverTab[98178]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:44
		// _ = "end of CoverTab[98178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:45
		_go_fuzz_dep_.CoverTab[98179]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:45
		// _ = "end of CoverTab[98179]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:45
	// _ = "end of CoverTab[98173]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:45
	_go_fuzz_dep_.CoverTab[98174]++

													a.Resources = make([]*AlterConfigsResourceResponse, responseCount)

													for i := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:49
		_go_fuzz_dep_.CoverTab[98180]++
														a.Resources[i] = new(AlterConfigsResourceResponse)

														if err := a.Resources[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:52
			_go_fuzz_dep_.CoverTab[98181]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:53
			// _ = "end of CoverTab[98181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:54
			_go_fuzz_dep_.CoverTab[98182]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:54
			// _ = "end of CoverTab[98182]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:54
		// _ = "end of CoverTab[98180]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:55
	// _ = "end of CoverTab[98174]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:55
	_go_fuzz_dep_.CoverTab[98175]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:57
	// _ = "end of CoverTab[98175]"
}

func (a *AlterConfigsResourceResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:60
	_go_fuzz_dep_.CoverTab[98183]++
													pe.putInt16(a.ErrorCode)
													err := pe.putString(a.ErrorMsg)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:63
		_go_fuzz_dep_.CoverTab[98186]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:64
		// _ = "end of CoverTab[98186]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:65
		_go_fuzz_dep_.CoverTab[98187]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:65
		// _ = "end of CoverTab[98187]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:65
	// _ = "end of CoverTab[98183]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:65
	_go_fuzz_dep_.CoverTab[98184]++
													pe.putInt8(int8(a.Type))
													err = pe.putString(a.Name)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:68
		_go_fuzz_dep_.CoverTab[98188]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:69
		// _ = "end of CoverTab[98188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:70
		_go_fuzz_dep_.CoverTab[98189]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:70
		// _ = "end of CoverTab[98189]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:70
	// _ = "end of CoverTab[98184]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:70
	_go_fuzz_dep_.CoverTab[98185]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:71
	// _ = "end of CoverTab[98185]"
}

func (a *AlterConfigsResourceResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:74
	_go_fuzz_dep_.CoverTab[98190]++
													errCode, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:76
		_go_fuzz_dep_.CoverTab[98195]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:77
		// _ = "end of CoverTab[98195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:78
		_go_fuzz_dep_.CoverTab[98196]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:78
		// _ = "end of CoverTab[98196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:78
	// _ = "end of CoverTab[98190]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:78
	_go_fuzz_dep_.CoverTab[98191]++
													a.ErrorCode = errCode

													e, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:82
		_go_fuzz_dep_.CoverTab[98197]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:83
		// _ = "end of CoverTab[98197]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:84
		_go_fuzz_dep_.CoverTab[98198]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:84
		// _ = "end of CoverTab[98198]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:84
	// _ = "end of CoverTab[98191]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:84
	_go_fuzz_dep_.CoverTab[98192]++
													a.ErrorMsg = e

													t, err := pd.getInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:88
		_go_fuzz_dep_.CoverTab[98199]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:89
		// _ = "end of CoverTab[98199]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:90
		_go_fuzz_dep_.CoverTab[98200]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:90
		// _ = "end of CoverTab[98200]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:90
	// _ = "end of CoverTab[98192]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:90
	_go_fuzz_dep_.CoverTab[98193]++
													a.Type = ConfigResourceType(t)

													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:94
		_go_fuzz_dep_.CoverTab[98201]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:95
		// _ = "end of CoverTab[98201]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:96
		_go_fuzz_dep_.CoverTab[98202]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:96
		// _ = "end of CoverTab[98202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:96
	// _ = "end of CoverTab[98193]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:96
	_go_fuzz_dep_.CoverTab[98194]++
													a.Name = name

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:99
	// _ = "end of CoverTab[98194]"
}

func (a *AlterConfigsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:102
	_go_fuzz_dep_.CoverTab[98203]++
													return 32
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:103
	// _ = "end of CoverTab[98203]"
}

func (a *AlterConfigsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:106
	_go_fuzz_dep_.CoverTab[98204]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:107
	// _ = "end of CoverTab[98204]"
}

func (a *AlterConfigsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:110
	_go_fuzz_dep_.CoverTab[98205]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:111
	// _ = "end of CoverTab[98205]"
}

func (a *AlterConfigsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:114
	_go_fuzz_dep_.CoverTab[98206]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:115
	// _ = "end of CoverTab[98206]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:116
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_configs_response.go:116
var _ = _go_fuzz_dep_.CoverTab
