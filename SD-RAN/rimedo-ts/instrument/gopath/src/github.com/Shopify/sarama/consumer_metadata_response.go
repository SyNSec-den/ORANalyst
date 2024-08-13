//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:1
)

import (
	"net"
	"strconv"
)

// ConsumerMetadataResponse holds the response for a consumer group meta data requests
type ConsumerMetadataResponse struct {
	Err		KError
	Coordinator	*Broker
	CoordinatorID	int32	// deprecated: use Coordinator.ID()
	CoordinatorHost	string	// deprecated: use Coordinator.Addr()
	CoordinatorPort	int32	// deprecated: use Coordinator.Addr()
}

func (r *ConsumerMetadataResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:17
	_go_fuzz_dep_.CoverTab[101361]++
													tmp := new(FindCoordinatorResponse)

													if err := tmp.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:20
		_go_fuzz_dep_.CoverTab[101366]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:21
		// _ = "end of CoverTab[101366]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:22
		_go_fuzz_dep_.CoverTab[101367]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:22
		// _ = "end of CoverTab[101367]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:22
	// _ = "end of CoverTab[101361]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:22
	_go_fuzz_dep_.CoverTab[101362]++

													r.Err = tmp.Err

													r.Coordinator = tmp.Coordinator
													if tmp.Coordinator == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:27
		_go_fuzz_dep_.CoverTab[101368]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:28
		// _ = "end of CoverTab[101368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:29
		_go_fuzz_dep_.CoverTab[101369]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:29
		// _ = "end of CoverTab[101369]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:29
	// _ = "end of CoverTab[101362]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:29
	_go_fuzz_dep_.CoverTab[101363]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:33
	host, portstr, err := net.SplitHostPort(r.Coordinator.Addr())
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:34
		_go_fuzz_dep_.CoverTab[101370]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:35
		// _ = "end of CoverTab[101370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:36
		_go_fuzz_dep_.CoverTab[101371]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:36
		// _ = "end of CoverTab[101371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:36
	// _ = "end of CoverTab[101363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:36
	_go_fuzz_dep_.CoverTab[101364]++
													port, err := strconv.ParseInt(portstr, 10, 32)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:38
		_go_fuzz_dep_.CoverTab[101372]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:39
		// _ = "end of CoverTab[101372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:40
		_go_fuzz_dep_.CoverTab[101373]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:40
		// _ = "end of CoverTab[101373]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:40
	// _ = "end of CoverTab[101364]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:40
	_go_fuzz_dep_.CoverTab[101365]++
													r.CoordinatorID = r.Coordinator.ID()
													r.CoordinatorHost = host
													r.CoordinatorPort = int32(port)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:45
	// _ = "end of CoverTab[101365]"
}

func (r *ConsumerMetadataResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:48
	_go_fuzz_dep_.CoverTab[101374]++
													if r.Coordinator == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:49
		_go_fuzz_dep_.CoverTab[101377]++
														r.Coordinator = new(Broker)
														r.Coordinator.id = r.CoordinatorID
														r.Coordinator.addr = net.JoinHostPort(r.CoordinatorHost, strconv.Itoa(int(r.CoordinatorPort)))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:52
		// _ = "end of CoverTab[101377]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:53
		_go_fuzz_dep_.CoverTab[101378]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:53
		// _ = "end of CoverTab[101378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:53
	// _ = "end of CoverTab[101374]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:53
	_go_fuzz_dep_.CoverTab[101375]++

													tmp := &FindCoordinatorResponse{
		Version:	0,
		Err:		r.Err,
		Coordinator:	r.Coordinator,
	}

	if err := tmp.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:61
		_go_fuzz_dep_.CoverTab[101379]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:62
		// _ = "end of CoverTab[101379]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:63
		_go_fuzz_dep_.CoverTab[101380]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:63
		// _ = "end of CoverTab[101380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:63
	// _ = "end of CoverTab[101375]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:63
	_go_fuzz_dep_.CoverTab[101376]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:65
	// _ = "end of CoverTab[101376]"
}

func (r *ConsumerMetadataResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:68
	_go_fuzz_dep_.CoverTab[101381]++
													return 10
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:69
	// _ = "end of CoverTab[101381]"
}

func (r *ConsumerMetadataResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:72
	_go_fuzz_dep_.CoverTab[101382]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:73
	// _ = "end of CoverTab[101382]"
}

func (r *ConsumerMetadataResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:76
	_go_fuzz_dep_.CoverTab[101383]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:77
	// _ = "end of CoverTab[101383]"
}

func (r *ConsumerMetadataResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:80
	_go_fuzz_dep_.CoverTab[101384]++
													return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:81
	// _ = "end of CoverTab[101384]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:82
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_response.go:82
var _ = _go_fuzz_dep_.CoverTab
