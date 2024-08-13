//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:1
)

type HeartbeatRequest struct {
	GroupId		string
	GenerationId	int32
	MemberId	string
}

func (r *HeartbeatRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:9
	_go_fuzz_dep_.CoverTab[103394]++
												if err := pe.putString(r.GroupId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:10
		_go_fuzz_dep_.CoverTab[103397]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:11
		// _ = "end of CoverTab[103397]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:12
		_go_fuzz_dep_.CoverTab[103398]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:12
		// _ = "end of CoverTab[103398]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:12
	// _ = "end of CoverTab[103394]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:12
	_go_fuzz_dep_.CoverTab[103395]++

												pe.putInt32(r.GenerationId)

												if err := pe.putString(r.MemberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:16
		_go_fuzz_dep_.CoverTab[103399]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:17
		// _ = "end of CoverTab[103399]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:18
		_go_fuzz_dep_.CoverTab[103400]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:18
		// _ = "end of CoverTab[103400]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:18
	// _ = "end of CoverTab[103395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:18
	_go_fuzz_dep_.CoverTab[103396]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:20
	// _ = "end of CoverTab[103396]"
}

func (r *HeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:23
	_go_fuzz_dep_.CoverTab[103401]++
												if r.GroupId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:24
		_go_fuzz_dep_.CoverTab[103405]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:25
		// _ = "end of CoverTab[103405]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:26
		_go_fuzz_dep_.CoverTab[103406]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:26
		// _ = "end of CoverTab[103406]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:26
	// _ = "end of CoverTab[103401]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:26
	_go_fuzz_dep_.CoverTab[103402]++
												if r.GenerationId, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:27
		_go_fuzz_dep_.CoverTab[103407]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:28
		// _ = "end of CoverTab[103407]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:29
		_go_fuzz_dep_.CoverTab[103408]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:29
		// _ = "end of CoverTab[103408]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:29
	// _ = "end of CoverTab[103402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:29
	_go_fuzz_dep_.CoverTab[103403]++
												if r.MemberId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:30
		_go_fuzz_dep_.CoverTab[103409]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:31
		// _ = "end of CoverTab[103409]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:32
		_go_fuzz_dep_.CoverTab[103410]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:32
		// _ = "end of CoverTab[103410]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:32
	// _ = "end of CoverTab[103403]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:32
	_go_fuzz_dep_.CoverTab[103404]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:34
	// _ = "end of CoverTab[103404]"
}

func (r *HeartbeatRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:37
	_go_fuzz_dep_.CoverTab[103411]++
												return 12
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:38
	// _ = "end of CoverTab[103411]"
}

func (r *HeartbeatRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:41
	_go_fuzz_dep_.CoverTab[103412]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:42
	// _ = "end of CoverTab[103412]"
}

func (r *HeartbeatRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:45
	_go_fuzz_dep_.CoverTab[103413]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:46
	// _ = "end of CoverTab[103413]"
}

func (r *HeartbeatRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:49
	_go_fuzz_dep_.CoverTab[103414]++
												return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:50
	// _ = "end of CoverTab[103414]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/heartbeat_request.go:51
var _ = _go_fuzz_dep_.CoverTab
