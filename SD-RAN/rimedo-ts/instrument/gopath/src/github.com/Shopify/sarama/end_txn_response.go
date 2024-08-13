//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:1
)

import (
	"time"
)

type EndTxnResponse struct {
	ThrottleTime	time.Duration
	Err		KError
}

func (e *EndTxnResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:12
	_go_fuzz_dep_.CoverTab[102733]++
												pe.putInt32(int32(e.ThrottleTime / time.Millisecond))
												pe.putInt16(int16(e.Err))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:15
	// _ = "end of CoverTab[102733]"
}

func (e *EndTxnResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:18
	_go_fuzz_dep_.CoverTab[102734]++
												throttleTime, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:20
		_go_fuzz_dep_.CoverTab[102737]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:21
		// _ = "end of CoverTab[102737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:22
		_go_fuzz_dep_.CoverTab[102738]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:22
		// _ = "end of CoverTab[102738]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:22
	// _ = "end of CoverTab[102734]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:22
	_go_fuzz_dep_.CoverTab[102735]++
												e.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

												kerr, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:26
		_go_fuzz_dep_.CoverTab[102739]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:27
		// _ = "end of CoverTab[102739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:28
		_go_fuzz_dep_.CoverTab[102740]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:28
		// _ = "end of CoverTab[102740]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:28
	// _ = "end of CoverTab[102735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:28
	_go_fuzz_dep_.CoverTab[102736]++
												e.Err = KError(kerr)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:31
	// _ = "end of CoverTab[102736]"
}

func (e *EndTxnResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:34
	_go_fuzz_dep_.CoverTab[102741]++
												return 25
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:35
	// _ = "end of CoverTab[102741]"
}

func (e *EndTxnResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:38
	_go_fuzz_dep_.CoverTab[102742]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:39
	// _ = "end of CoverTab[102742]"
}

func (r *EndTxnResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:42
	_go_fuzz_dep_.CoverTab[102743]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:43
	// _ = "end of CoverTab[102743]"
}

func (e *EndTxnResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:46
	_go_fuzz_dep_.CoverTab[102744]++
												return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:47
	// _ = "end of CoverTab[102744]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_response.go:48
var _ = _go_fuzz_dep_.CoverTab
