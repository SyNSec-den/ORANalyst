//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:1
)

import "time"

type DeleteTopicsRequest struct {
	Version	int16
	Topics	[]string
	Timeout	time.Duration
}

func (d *DeleteTopicsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:11
	_go_fuzz_dep_.CoverTab[101955]++
													if err := pe.putStringArray(d.Topics); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:12
		_go_fuzz_dep_.CoverTab[101957]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:13
		// _ = "end of CoverTab[101957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:14
		_go_fuzz_dep_.CoverTab[101958]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:14
		// _ = "end of CoverTab[101958]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:14
	// _ = "end of CoverTab[101955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:14
	_go_fuzz_dep_.CoverTab[101956]++
													pe.putInt32(int32(d.Timeout / time.Millisecond))

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:17
	// _ = "end of CoverTab[101956]"
}

func (d *DeleteTopicsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:20
	_go_fuzz_dep_.CoverTab[101959]++
													if d.Topics, err = pd.getStringArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:21
		_go_fuzz_dep_.CoverTab[101962]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:22
		// _ = "end of CoverTab[101962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:23
		_go_fuzz_dep_.CoverTab[101963]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:23
		// _ = "end of CoverTab[101963]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:23
	// _ = "end of CoverTab[101959]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:23
	_go_fuzz_dep_.CoverTab[101960]++
													timeout, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:25
		_go_fuzz_dep_.CoverTab[101964]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:26
		// _ = "end of CoverTab[101964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:27
		_go_fuzz_dep_.CoverTab[101965]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:27
		// _ = "end of CoverTab[101965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:27
	// _ = "end of CoverTab[101960]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:27
	_go_fuzz_dep_.CoverTab[101961]++
													d.Timeout = time.Duration(timeout) * time.Millisecond
													d.Version = version
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:30
	// _ = "end of CoverTab[101961]"
}

func (d *DeleteTopicsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:33
	_go_fuzz_dep_.CoverTab[101966]++
													return 20
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:34
	// _ = "end of CoverTab[101966]"
}

func (d *DeleteTopicsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:37
	_go_fuzz_dep_.CoverTab[101967]++
													return d.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:38
	// _ = "end of CoverTab[101967]"
}

func (d *DeleteTopicsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:41
	_go_fuzz_dep_.CoverTab[101968]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:42
	// _ = "end of CoverTab[101968]"
}

func (d *DeleteTopicsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:45
	_go_fuzz_dep_.CoverTab[101969]++
													switch d.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:47
		_go_fuzz_dep_.CoverTab[101970]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:48
		// _ = "end of CoverTab[101970]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:49
		_go_fuzz_dep_.CoverTab[101971]++
														return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:50
		// _ = "end of CoverTab[101971]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:51
	// _ = "end of CoverTab[101969]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_topics_request.go:52
var _ = _go_fuzz_dep_.CoverTab
