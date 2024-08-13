//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:1
)

import "time"

type InitProducerIDResponse struct {
	ThrottleTime	time.Duration
	Err		KError
	ProducerID	int64
	ProducerEpoch	int16
}

func (i *InitProducerIDResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:12
	_go_fuzz_dep_.CoverTab[103531]++
													pe.putInt32(int32(i.ThrottleTime / time.Millisecond))
													pe.putInt16(int16(i.Err))
													pe.putInt64(i.ProducerID)
													pe.putInt16(i.ProducerEpoch)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:18
	// _ = "end of CoverTab[103531]"
}

func (i *InitProducerIDResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:21
	_go_fuzz_dep_.CoverTab[103532]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:23
		_go_fuzz_dep_.CoverTab[103537]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:24
		// _ = "end of CoverTab[103537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:25
		_go_fuzz_dep_.CoverTab[103538]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:25
		// _ = "end of CoverTab[103538]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:25
	// _ = "end of CoverTab[103532]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:25
	_go_fuzz_dep_.CoverTab[103533]++
													i.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:29
		_go_fuzz_dep_.CoverTab[103539]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:30
		// _ = "end of CoverTab[103539]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:31
		_go_fuzz_dep_.CoverTab[103540]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:31
		// _ = "end of CoverTab[103540]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:31
	// _ = "end of CoverTab[103533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:31
	_go_fuzz_dep_.CoverTab[103534]++
													i.Err = KError(kerr)

													if i.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:34
		_go_fuzz_dep_.CoverTab[103541]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:35
		// _ = "end of CoverTab[103541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:36
		_go_fuzz_dep_.CoverTab[103542]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:36
		// _ = "end of CoverTab[103542]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:36
	// _ = "end of CoverTab[103534]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:36
	_go_fuzz_dep_.CoverTab[103535]++

													if i.ProducerEpoch, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:38
		_go_fuzz_dep_.CoverTab[103543]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:39
		// _ = "end of CoverTab[103543]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:40
		_go_fuzz_dep_.CoverTab[103544]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:40
		// _ = "end of CoverTab[103544]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:40
	// _ = "end of CoverTab[103535]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:40
	_go_fuzz_dep_.CoverTab[103536]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:42
	// _ = "end of CoverTab[103536]"
}

func (i *InitProducerIDResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:45
	_go_fuzz_dep_.CoverTab[103545]++
													return 22
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:46
	// _ = "end of CoverTab[103545]"
}

func (i *InitProducerIDResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:49
	_go_fuzz_dep_.CoverTab[103546]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:50
	// _ = "end of CoverTab[103546]"
}

func (i *InitProducerIDResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:53
	_go_fuzz_dep_.CoverTab[103547]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:54
	// _ = "end of CoverTab[103547]"
}

func (i *InitProducerIDResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:57
	_go_fuzz_dep_.CoverTab[103548]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:58
	// _ = "end of CoverTab[103548]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:59
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_response.go:59
var _ = _go_fuzz_dep_.CoverTab
