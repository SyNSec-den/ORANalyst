//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:1
)

// ConsumerMetadataRequest is used for metadata requests
type ConsumerMetadataRequest struct {
	ConsumerGroup string
}

func (r *ConsumerMetadataRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:8
	_go_fuzz_dep_.CoverTab[101352]++
													tmp := new(FindCoordinatorRequest)
													tmp.CoordinatorKey = r.ConsumerGroup
													tmp.CoordinatorType = CoordinatorGroup
													return tmp.encode(pe)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:12
	// _ = "end of CoverTab[101352]"
}

func (r *ConsumerMetadataRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:15
	_go_fuzz_dep_.CoverTab[101353]++
													tmp := new(FindCoordinatorRequest)
													if err := tmp.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:17
		_go_fuzz_dep_.CoverTab[101355]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:18
		// _ = "end of CoverTab[101355]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:19
		_go_fuzz_dep_.CoverTab[101356]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:19
		// _ = "end of CoverTab[101356]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:19
	// _ = "end of CoverTab[101353]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:19
	_go_fuzz_dep_.CoverTab[101354]++
													r.ConsumerGroup = tmp.CoordinatorKey
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:21
	// _ = "end of CoverTab[101354]"
}

func (r *ConsumerMetadataRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:24
	_go_fuzz_dep_.CoverTab[101357]++
													return 10
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:25
	// _ = "end of CoverTab[101357]"
}

func (r *ConsumerMetadataRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:28
	_go_fuzz_dep_.CoverTab[101358]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:29
	// _ = "end of CoverTab[101358]"
}

func (r *ConsumerMetadataRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:32
	_go_fuzz_dep_.CoverTab[101359]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:33
	// _ = "end of CoverTab[101359]"
}

func (r *ConsumerMetadataRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:36
	_go_fuzz_dep_.CoverTab[101360]++
													return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:37
	// _ = "end of CoverTab[101360]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:38
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_metadata_request.go:38
var _ = _go_fuzz_dep_.CoverTab
