//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:1
)

type EndTxnRequest struct {
	TransactionalID		string
	ProducerID		int64
	ProducerEpoch		int16
	TransactionResult	bool
}

func (a *EndTxnRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:10
	_go_fuzz_dep_.CoverTab[102712]++
												if err := pe.putString(a.TransactionalID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:11
		_go_fuzz_dep_.CoverTab[102714]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:12
		// _ = "end of CoverTab[102714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:13
		_go_fuzz_dep_.CoverTab[102715]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:13
		// _ = "end of CoverTab[102715]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:13
	// _ = "end of CoverTab[102712]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:13
	_go_fuzz_dep_.CoverTab[102713]++

												pe.putInt64(a.ProducerID)

												pe.putInt16(a.ProducerEpoch)

												pe.putBool(a.TransactionResult)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:21
	// _ = "end of CoverTab[102713]"
}

func (a *EndTxnRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:24
	_go_fuzz_dep_.CoverTab[102716]++
												if a.TransactionalID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:25
		_go_fuzz_dep_.CoverTab[102721]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:26
		// _ = "end of CoverTab[102721]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:27
		_go_fuzz_dep_.CoverTab[102722]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:27
		// _ = "end of CoverTab[102722]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:27
	// _ = "end of CoverTab[102716]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:27
	_go_fuzz_dep_.CoverTab[102717]++
												if a.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:28
		_go_fuzz_dep_.CoverTab[102723]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:29
		// _ = "end of CoverTab[102723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:30
		_go_fuzz_dep_.CoverTab[102724]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:30
		// _ = "end of CoverTab[102724]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:30
	// _ = "end of CoverTab[102717]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:30
	_go_fuzz_dep_.CoverTab[102718]++
												if a.ProducerEpoch, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:31
		_go_fuzz_dep_.CoverTab[102725]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:32
		// _ = "end of CoverTab[102725]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:33
		_go_fuzz_dep_.CoverTab[102726]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:33
		// _ = "end of CoverTab[102726]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:33
	// _ = "end of CoverTab[102718]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:33
	_go_fuzz_dep_.CoverTab[102719]++
												if a.TransactionResult, err = pd.getBool(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:34
		_go_fuzz_dep_.CoverTab[102727]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:35
		// _ = "end of CoverTab[102727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:36
		_go_fuzz_dep_.CoverTab[102728]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:36
		// _ = "end of CoverTab[102728]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:36
	// _ = "end of CoverTab[102719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:36
	_go_fuzz_dep_.CoverTab[102720]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:37
	// _ = "end of CoverTab[102720]"
}

func (a *EndTxnRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:40
	_go_fuzz_dep_.CoverTab[102729]++
												return 26
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:41
	// _ = "end of CoverTab[102729]"
}

func (a *EndTxnRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:44
	_go_fuzz_dep_.CoverTab[102730]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:45
	// _ = "end of CoverTab[102730]"
}

func (r *EndTxnRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:48
	_go_fuzz_dep_.CoverTab[102731]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:49
	// _ = "end of CoverTab[102731]"
}

func (a *EndTxnRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:52
	_go_fuzz_dep_.CoverTab[102732]++
												return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:53
	// _ = "end of CoverTab[102732]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:54
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/end_txn_request.go:54
var _ = _go_fuzz_dep_.CoverTab
