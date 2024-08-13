//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:1
)

import "time"

type InitProducerIDRequest struct {
	TransactionalID		*string
	TransactionTimeout	time.Duration
}

func (i *InitProducerIDRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:10
	_go_fuzz_dep_.CoverTab[103516]++
													if err := pe.putNullableString(i.TransactionalID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:11
		_go_fuzz_dep_.CoverTab[103518]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:12
		// _ = "end of CoverTab[103518]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:13
		_go_fuzz_dep_.CoverTab[103519]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:13
		// _ = "end of CoverTab[103519]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:13
	// _ = "end of CoverTab[103516]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:13
	_go_fuzz_dep_.CoverTab[103517]++
													pe.putInt32(int32(i.TransactionTimeout / time.Millisecond))

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:16
	// _ = "end of CoverTab[103517]"
}

func (i *InitProducerIDRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:19
	_go_fuzz_dep_.CoverTab[103520]++
													if i.TransactionalID, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:20
		_go_fuzz_dep_.CoverTab[103523]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:21
		// _ = "end of CoverTab[103523]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:22
		_go_fuzz_dep_.CoverTab[103524]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:22
		// _ = "end of CoverTab[103524]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:22
	// _ = "end of CoverTab[103520]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:22
	_go_fuzz_dep_.CoverTab[103521]++

													timeout, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:25
		_go_fuzz_dep_.CoverTab[103525]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:26
		// _ = "end of CoverTab[103525]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:27
		_go_fuzz_dep_.CoverTab[103526]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:27
		// _ = "end of CoverTab[103526]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:27
	// _ = "end of CoverTab[103521]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:27
	_go_fuzz_dep_.CoverTab[103522]++
													i.TransactionTimeout = time.Duration(timeout) * time.Millisecond

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:30
	// _ = "end of CoverTab[103522]"
}

func (i *InitProducerIDRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:33
	_go_fuzz_dep_.CoverTab[103527]++
													return 22
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:34
	// _ = "end of CoverTab[103527]"
}

func (i *InitProducerIDRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:37
	_go_fuzz_dep_.CoverTab[103528]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:38
	// _ = "end of CoverTab[103528]"
}

func (i *InitProducerIDRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:41
	_go_fuzz_dep_.CoverTab[103529]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:42
	// _ = "end of CoverTab[103529]"
}

func (i *InitProducerIDRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:45
	_go_fuzz_dep_.CoverTab[103530]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:46
	// _ = "end of CoverTab[103530]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:47
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/init_producer_id_request.go:47
var _ = _go_fuzz_dep_.CoverTab
