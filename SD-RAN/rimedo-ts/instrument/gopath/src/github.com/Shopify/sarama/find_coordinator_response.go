//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:1
)

import (
	"time"
)

var NoNode = &Broker{id: -1, addr: ":-1"}

type FindCoordinatorResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	Err		KError
	ErrMsg		*string
	Coordinator	*Broker
}

func (f *FindCoordinatorResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:17
	_go_fuzz_dep_.CoverTab[103281]++
													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:18
		_go_fuzz_dep_.CoverTab[103287]++
														f.Version = version

														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:22
			_go_fuzz_dep_.CoverTab[103289]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:23
			// _ = "end of CoverTab[103289]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:24
			_go_fuzz_dep_.CoverTab[103290]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:24
			// _ = "end of CoverTab[103290]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:24
		// _ = "end of CoverTab[103287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:24
		_go_fuzz_dep_.CoverTab[103288]++
														f.ThrottleTime = time.Duration(throttleTime) * time.Millisecond
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:25
		// _ = "end of CoverTab[103288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:26
		_go_fuzz_dep_.CoverTab[103291]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:26
		// _ = "end of CoverTab[103291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:26
	// _ = "end of CoverTab[103281]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:26
	_go_fuzz_dep_.CoverTab[103282]++

													tmp, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:29
		_go_fuzz_dep_.CoverTab[103292]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:30
		// _ = "end of CoverTab[103292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:31
		_go_fuzz_dep_.CoverTab[103293]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:31
		// _ = "end of CoverTab[103293]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:31
	// _ = "end of CoverTab[103282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:31
	_go_fuzz_dep_.CoverTab[103283]++
													f.Err = KError(tmp)

													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:34
		_go_fuzz_dep_.CoverTab[103294]++
														if f.ErrMsg, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:35
			_go_fuzz_dep_.CoverTab[103295]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:36
			// _ = "end of CoverTab[103295]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:37
			_go_fuzz_dep_.CoverTab[103296]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:37
			// _ = "end of CoverTab[103296]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:37
		// _ = "end of CoverTab[103294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:38
		_go_fuzz_dep_.CoverTab[103297]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:38
		// _ = "end of CoverTab[103297]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:38
	// _ = "end of CoverTab[103283]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:38
	_go_fuzz_dep_.CoverTab[103284]++

													coordinator := new(Broker)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:43
	if err := coordinator.decode(pd, 0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:43
		_go_fuzz_dep_.CoverTab[103298]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:44
		// _ = "end of CoverTab[103298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:45
		_go_fuzz_dep_.CoverTab[103299]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:45
		// _ = "end of CoverTab[103299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:45
	// _ = "end of CoverTab[103284]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:45
	_go_fuzz_dep_.CoverTab[103285]++
													if coordinator.addr == ":0" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:46
		_go_fuzz_dep_.CoverTab[103300]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:47
		// _ = "end of CoverTab[103300]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:48
		_go_fuzz_dep_.CoverTab[103301]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:48
		// _ = "end of CoverTab[103301]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:48
	// _ = "end of CoverTab[103285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:48
	_go_fuzz_dep_.CoverTab[103286]++
													f.Coordinator = coordinator

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:51
	// _ = "end of CoverTab[103286]"
}

func (f *FindCoordinatorResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:54
	_go_fuzz_dep_.CoverTab[103302]++
													if f.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:55
		_go_fuzz_dep_.CoverTab[103307]++
														pe.putInt32(int32(f.ThrottleTime / time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:56
		// _ = "end of CoverTab[103307]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:57
		_go_fuzz_dep_.CoverTab[103308]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:57
		// _ = "end of CoverTab[103308]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:57
	// _ = "end of CoverTab[103302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:57
	_go_fuzz_dep_.CoverTab[103303]++

													pe.putInt16(int16(f.Err))

													if f.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:61
		_go_fuzz_dep_.CoverTab[103309]++
														if err := pe.putNullableString(f.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:62
			_go_fuzz_dep_.CoverTab[103310]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:63
			// _ = "end of CoverTab[103310]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:64
			_go_fuzz_dep_.CoverTab[103311]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:64
			// _ = "end of CoverTab[103311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:64
		// _ = "end of CoverTab[103309]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:65
		_go_fuzz_dep_.CoverTab[103312]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:65
		// _ = "end of CoverTab[103312]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:65
	// _ = "end of CoverTab[103303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:65
	_go_fuzz_dep_.CoverTab[103304]++

													coordinator := f.Coordinator
													if coordinator == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:68
		_go_fuzz_dep_.CoverTab[103313]++
														coordinator = NoNode
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:69
		// _ = "end of CoverTab[103313]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:70
		_go_fuzz_dep_.CoverTab[103314]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:70
		// _ = "end of CoverTab[103314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:70
	// _ = "end of CoverTab[103304]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:70
	_go_fuzz_dep_.CoverTab[103305]++
													if err := coordinator.encode(pe, 0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:71
		_go_fuzz_dep_.CoverTab[103315]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:72
		// _ = "end of CoverTab[103315]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:73
		_go_fuzz_dep_.CoverTab[103316]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:73
		// _ = "end of CoverTab[103316]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:73
	// _ = "end of CoverTab[103305]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:73
	_go_fuzz_dep_.CoverTab[103306]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:74
	// _ = "end of CoverTab[103306]"
}

func (f *FindCoordinatorResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:77
	_go_fuzz_dep_.CoverTab[103317]++
													return 10
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:78
	// _ = "end of CoverTab[103317]"
}

func (f *FindCoordinatorResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:81
	_go_fuzz_dep_.CoverTab[103318]++
													return f.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:82
	// _ = "end of CoverTab[103318]"
}

func (r *FindCoordinatorResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:85
	_go_fuzz_dep_.CoverTab[103319]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:86
	// _ = "end of CoverTab[103319]"
}

func (f *FindCoordinatorResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:89
	_go_fuzz_dep_.CoverTab[103320]++
													switch f.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:91
		_go_fuzz_dep_.CoverTab[103321]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:92
		// _ = "end of CoverTab[103321]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:93
		_go_fuzz_dep_.CoverTab[103322]++
														return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:94
		// _ = "end of CoverTab[103322]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:95
	// _ = "end of CoverTab[103320]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:96
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_response.go:96
var _ = _go_fuzz_dep_.CoverTab
