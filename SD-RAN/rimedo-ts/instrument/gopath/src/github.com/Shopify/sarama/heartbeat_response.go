//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:1
)

type HeartbeatResponse struct {
	Err KError
}

func (r *HeartbeatResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:7
	_go_fuzz_dep_.CoverTab[103415]++
												pe.putInt16(int16(r.Err))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:9
	// _ = "end of CoverTab[103415]"
}

func (r *HeartbeatResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:12
	_go_fuzz_dep_.CoverTab[103416]++
												kerr, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:14
		_go_fuzz_dep_.CoverTab[103418]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:15
		// _ = "end of CoverTab[103418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:16
		_go_fuzz_dep_.CoverTab[103419]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:16
		// _ = "end of CoverTab[103419]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:16
	// _ = "end of CoverTab[103416]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:16
	_go_fuzz_dep_.CoverTab[103417]++
												r.Err = KError(kerr)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:19
	// _ = "end of CoverTab[103417]"
}

func (r *HeartbeatResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:22
	_go_fuzz_dep_.CoverTab[103420]++
												return 12
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:23
	// _ = "end of CoverTab[103420]"
}

func (r *HeartbeatResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:26
	_go_fuzz_dep_.CoverTab[103421]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:27
	// _ = "end of CoverTab[103421]"
}

func (r *HeartbeatResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:30
	_go_fuzz_dep_.CoverTab[103422]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:31
	// _ = "end of CoverTab[103422]"
}

func (r *HeartbeatResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:34
	_go_fuzz_dep_.CoverTab[103423]++
												return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:35
	// _ = "end of CoverTab[103423]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_response.go:36
var _ = _go_fuzz_dep_.CoverTab
