//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:1
)

import "time"

// IncrementalAlterConfigsResponse is a response type for incremental alter config
type IncrementalAlterConfigsResponse struct {
	ThrottleTime	time.Duration
	Resources	[]*AlterConfigsResourceResponse
}

func (a *IncrementalAlterConfigsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:11
	_go_fuzz_dep_.CoverTab[103493]++
														pe.putInt32(int32(a.ThrottleTime / time.Millisecond))

														if err := pe.putArrayLength(len(a.Resources)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:14
		_go_fuzz_dep_.CoverTab[103496]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:15
		// _ = "end of CoverTab[103496]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:16
		_go_fuzz_dep_.CoverTab[103497]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:16
		// _ = "end of CoverTab[103497]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:16
	// _ = "end of CoverTab[103493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:16
	_go_fuzz_dep_.CoverTab[103494]++

														for _, v := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:18
		_go_fuzz_dep_.CoverTab[103498]++
															if err := v.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:19
			_go_fuzz_dep_.CoverTab[103499]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:20
			// _ = "end of CoverTab[103499]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:21
			_go_fuzz_dep_.CoverTab[103500]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:21
			// _ = "end of CoverTab[103500]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:21
		// _ = "end of CoverTab[103498]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:22
	// _ = "end of CoverTab[103494]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:22
	_go_fuzz_dep_.CoverTab[103495]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:24
	// _ = "end of CoverTab[103495]"
}

func (a *IncrementalAlterConfigsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:27
	_go_fuzz_dep_.CoverTab[103501]++
														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:29
		_go_fuzz_dep_.CoverTab[103505]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:30
		// _ = "end of CoverTab[103505]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:31
		_go_fuzz_dep_.CoverTab[103506]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:31
		// _ = "end of CoverTab[103506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:31
	// _ = "end of CoverTab[103501]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:31
	_go_fuzz_dep_.CoverTab[103502]++
														a.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

														responseCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:35
		_go_fuzz_dep_.CoverTab[103507]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:36
		// _ = "end of CoverTab[103507]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:37
		_go_fuzz_dep_.CoverTab[103508]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:37
		// _ = "end of CoverTab[103508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:37
	// _ = "end of CoverTab[103502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:37
	_go_fuzz_dep_.CoverTab[103503]++

														a.Resources = make([]*AlterConfigsResourceResponse, responseCount)

														for i := range a.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:41
		_go_fuzz_dep_.CoverTab[103509]++
															a.Resources[i] = new(AlterConfigsResourceResponse)

															if err := a.Resources[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:44
			_go_fuzz_dep_.CoverTab[103510]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:45
			// _ = "end of CoverTab[103510]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:46
			_go_fuzz_dep_.CoverTab[103511]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:46
			// _ = "end of CoverTab[103511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:46
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:46
		// _ = "end of CoverTab[103509]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:47
	// _ = "end of CoverTab[103503]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:47
	_go_fuzz_dep_.CoverTab[103504]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:49
	// _ = "end of CoverTab[103504]"
}

func (a *IncrementalAlterConfigsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:52
	_go_fuzz_dep_.CoverTab[103512]++
														return 44
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:53
	// _ = "end of CoverTab[103512]"
}

func (a *IncrementalAlterConfigsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:56
	_go_fuzz_dep_.CoverTab[103513]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:57
	// _ = "end of CoverTab[103513]"
}

func (a *IncrementalAlterConfigsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:60
	_go_fuzz_dep_.CoverTab[103514]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:61
	// _ = "end of CoverTab[103514]"
}

func (a *IncrementalAlterConfigsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:64
	_go_fuzz_dep_.CoverTab[103515]++
														return V2_3_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:65
	// _ = "end of CoverTab[103515]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:66
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/incremental_alter_configs_response.go:66
var _ = _go_fuzz_dep_.CoverTab
