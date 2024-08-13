//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:1
)

// AddOffsetsToTxnRequest adds offsets to a transaction request
type AddOffsetsToTxnRequest struct {
	TransactionalID	string
	ProducerID	int64
	ProducerEpoch	int16
	GroupID		string
}

func (a *AddOffsetsToTxnRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:11
	_go_fuzz_dep_.CoverTab[97409]++
													if err := pe.putString(a.TransactionalID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:12
		_go_fuzz_dep_.CoverTab[97412]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:13
		// _ = "end of CoverTab[97412]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:14
		_go_fuzz_dep_.CoverTab[97413]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:14
		// _ = "end of CoverTab[97413]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:14
	// _ = "end of CoverTab[97409]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:14
	_go_fuzz_dep_.CoverTab[97410]++

													pe.putInt64(a.ProducerID)

													pe.putInt16(a.ProducerEpoch)

													if err := pe.putString(a.GroupID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:20
		_go_fuzz_dep_.CoverTab[97414]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:21
		// _ = "end of CoverTab[97414]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:22
		_go_fuzz_dep_.CoverTab[97415]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:22
		// _ = "end of CoverTab[97415]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:22
	// _ = "end of CoverTab[97410]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:22
	_go_fuzz_dep_.CoverTab[97411]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:24
	// _ = "end of CoverTab[97411]"
}

func (a *AddOffsetsToTxnRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:27
	_go_fuzz_dep_.CoverTab[97416]++
													if a.TransactionalID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:28
		_go_fuzz_dep_.CoverTab[97421]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:29
		// _ = "end of CoverTab[97421]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:30
		_go_fuzz_dep_.CoverTab[97422]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:30
		// _ = "end of CoverTab[97422]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:30
	// _ = "end of CoverTab[97416]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:30
	_go_fuzz_dep_.CoverTab[97417]++
													if a.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:31
		_go_fuzz_dep_.CoverTab[97423]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:32
		// _ = "end of CoverTab[97423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:33
		_go_fuzz_dep_.CoverTab[97424]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:33
		// _ = "end of CoverTab[97424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:33
	// _ = "end of CoverTab[97417]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:33
	_go_fuzz_dep_.CoverTab[97418]++
													if a.ProducerEpoch, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:34
		_go_fuzz_dep_.CoverTab[97425]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:35
		// _ = "end of CoverTab[97425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:36
		_go_fuzz_dep_.CoverTab[97426]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:36
		// _ = "end of CoverTab[97426]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:36
	// _ = "end of CoverTab[97418]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:36
	_go_fuzz_dep_.CoverTab[97419]++
													if a.GroupID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:37
		_go_fuzz_dep_.CoverTab[97427]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:38
		// _ = "end of CoverTab[97427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:39
		_go_fuzz_dep_.CoverTab[97428]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:39
		// _ = "end of CoverTab[97428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:39
	// _ = "end of CoverTab[97419]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:39
	_go_fuzz_dep_.CoverTab[97420]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:40
	// _ = "end of CoverTab[97420]"
}

func (a *AddOffsetsToTxnRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:43
	_go_fuzz_dep_.CoverTab[97429]++
													return 25
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:44
	// _ = "end of CoverTab[97429]"
}

func (a *AddOffsetsToTxnRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:47
	_go_fuzz_dep_.CoverTab[97430]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:48
	// _ = "end of CoverTab[97430]"
}

func (a *AddOffsetsToTxnRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:51
	_go_fuzz_dep_.CoverTab[97431]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:52
	// _ = "end of CoverTab[97431]"
}

func (a *AddOffsetsToTxnRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:55
	_go_fuzz_dep_.CoverTab[97432]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:56
	// _ = "end of CoverTab[97432]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/add_offsets_to_txn_request.go:57
var _ = _go_fuzz_dep_.CoverTab
