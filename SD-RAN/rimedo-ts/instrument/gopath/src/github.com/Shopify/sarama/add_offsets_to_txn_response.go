//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:1
)

import (
	"time"
)

// AddOffsetsToTxnResponse is a response type for adding offsets to txns
type AddOffsetsToTxnResponse struct {
	ThrottleTime	time.Duration
	Err		KError
}

func (a *AddOffsetsToTxnResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:13
	_go_fuzz_dep_.CoverTab[97433]++
													pe.putInt32(int32(a.ThrottleTime / time.Millisecond))
													pe.putInt16(int16(a.Err))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:16
	// _ = "end of CoverTab[97433]"
}

func (a *AddOffsetsToTxnResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:19
	_go_fuzz_dep_.CoverTab[97434]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:21
		_go_fuzz_dep_.CoverTab[97437]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:22
		// _ = "end of CoverTab[97437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:23
		_go_fuzz_dep_.CoverTab[97438]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:23
		// _ = "end of CoverTab[97438]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:23
	// _ = "end of CoverTab[97434]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:23
	_go_fuzz_dep_.CoverTab[97435]++
													a.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:27
		_go_fuzz_dep_.CoverTab[97439]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:28
		// _ = "end of CoverTab[97439]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:29
		_go_fuzz_dep_.CoverTab[97440]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:29
		// _ = "end of CoverTab[97440]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:29
	// _ = "end of CoverTab[97435]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:29
	_go_fuzz_dep_.CoverTab[97436]++
													a.Err = KError(kerr)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:32
	// _ = "end of CoverTab[97436]"
}

func (a *AddOffsetsToTxnResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:35
	_go_fuzz_dep_.CoverTab[97441]++
													return 25
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:36
	// _ = "end of CoverTab[97441]"
}

func (a *AddOffsetsToTxnResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:39
	_go_fuzz_dep_.CoverTab[97442]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:40
	// _ = "end of CoverTab[97442]"
}

func (a *AddOffsetsToTxnResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:43
	_go_fuzz_dep_.CoverTab[97443]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:44
	// _ = "end of CoverTab[97443]"
}

func (a *AddOffsetsToTxnResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:47
	_go_fuzz_dep_.CoverTab[97444]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:48
	// _ = "end of CoverTab[97444]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:49
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_response.go:49
var _ = _go_fuzz_dep_.CoverTab
