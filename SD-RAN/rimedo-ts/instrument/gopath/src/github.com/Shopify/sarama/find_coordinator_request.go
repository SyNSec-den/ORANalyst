//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:1
)

type CoordinatorType int8

const (
	CoordinatorGroup	CoordinatorType	= iota
	CoordinatorTransaction
)

type FindCoordinatorRequest struct {
	Version		int16
	CoordinatorKey	string
	CoordinatorType	CoordinatorType
}

func (f *FindCoordinatorRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:16
	_go_fuzz_dep_.CoverTab[103258]++
													if err := pe.putString(f.CoordinatorKey); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:17
		_go_fuzz_dep_.CoverTab[103261]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:18
		// _ = "end of CoverTab[103261]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:19
		_go_fuzz_dep_.CoverTab[103262]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:19
		// _ = "end of CoverTab[103262]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:19
	// _ = "end of CoverTab[103258]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:19
	_go_fuzz_dep_.CoverTab[103259]++

													if f.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:21
		_go_fuzz_dep_.CoverTab[103263]++
														pe.putInt8(int8(f.CoordinatorType))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:22
		// _ = "end of CoverTab[103263]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:23
		_go_fuzz_dep_.CoverTab[103264]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:23
		// _ = "end of CoverTab[103264]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:23
	// _ = "end of CoverTab[103259]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:23
	_go_fuzz_dep_.CoverTab[103260]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:25
	// _ = "end of CoverTab[103260]"
}

func (f *FindCoordinatorRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:28
	_go_fuzz_dep_.CoverTab[103265]++
													if f.CoordinatorKey, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:29
		_go_fuzz_dep_.CoverTab[103268]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:30
		// _ = "end of CoverTab[103268]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:31
		_go_fuzz_dep_.CoverTab[103269]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:31
		// _ = "end of CoverTab[103269]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:31
	// _ = "end of CoverTab[103265]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:31
	_go_fuzz_dep_.CoverTab[103266]++

													if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:33
		_go_fuzz_dep_.CoverTab[103270]++
														f.Version = version
														coordinatorType, err := pd.getInt8()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:36
			_go_fuzz_dep_.CoverTab[103272]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:37
			// _ = "end of CoverTab[103272]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:38
			_go_fuzz_dep_.CoverTab[103273]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:38
			// _ = "end of CoverTab[103273]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:38
		// _ = "end of CoverTab[103270]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:38
		_go_fuzz_dep_.CoverTab[103271]++

														f.CoordinatorType = CoordinatorType(coordinatorType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:40
		// _ = "end of CoverTab[103271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:41
		_go_fuzz_dep_.CoverTab[103274]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:41
		// _ = "end of CoverTab[103274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:41
	// _ = "end of CoverTab[103266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:41
	_go_fuzz_dep_.CoverTab[103267]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:43
	// _ = "end of CoverTab[103267]"
}

func (f *FindCoordinatorRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:46
	_go_fuzz_dep_.CoverTab[103275]++
													return 10
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:47
	// _ = "end of CoverTab[103275]"
}

func (f *FindCoordinatorRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:50
	_go_fuzz_dep_.CoverTab[103276]++
													return f.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:51
	// _ = "end of CoverTab[103276]"
}

func (r *FindCoordinatorRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:54
	_go_fuzz_dep_.CoverTab[103277]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:55
	// _ = "end of CoverTab[103277]"
}

func (f *FindCoordinatorRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:58
	_go_fuzz_dep_.CoverTab[103278]++
													switch f.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:60
		_go_fuzz_dep_.CoverTab[103279]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:61
		// _ = "end of CoverTab[103279]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:62
		_go_fuzz_dep_.CoverTab[103280]++
														return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:63
		// _ = "end of CoverTab[103280]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:64
	// _ = "end of CoverTab[103278]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:65
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/find_coordinator_request.go:65
var _ = _go_fuzz_dep_.CoverTab
