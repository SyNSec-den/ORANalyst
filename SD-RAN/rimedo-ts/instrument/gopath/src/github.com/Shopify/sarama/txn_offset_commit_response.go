//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:1
)

import (
	"time"
)

type TxnOffsetCommitResponse struct {
	ThrottleTime	time.Duration
	Topics		map[string][]*PartitionError
}

func (t *TxnOffsetCommitResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:12
	_go_fuzz_dep_.CoverTab[106974]++
													pe.putInt32(int32(t.ThrottleTime / time.Millisecond))
													if err := pe.putArrayLength(len(t.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:14
		_go_fuzz_dep_.CoverTab[106977]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:15
		// _ = "end of CoverTab[106977]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:16
		_go_fuzz_dep_.CoverTab[106978]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:16
		// _ = "end of CoverTab[106978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:16
	// _ = "end of CoverTab[106974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:16
	_go_fuzz_dep_.CoverTab[106975]++

													for topic, e := range t.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:18
		_go_fuzz_dep_.CoverTab[106979]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:19
			_go_fuzz_dep_.CoverTab[106982]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:20
			// _ = "end of CoverTab[106982]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:21
			_go_fuzz_dep_.CoverTab[106983]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:21
			// _ = "end of CoverTab[106983]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:21
		// _ = "end of CoverTab[106979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:21
		_go_fuzz_dep_.CoverTab[106980]++
														if err := pe.putArrayLength(len(e)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:22
			_go_fuzz_dep_.CoverTab[106984]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:23
			// _ = "end of CoverTab[106984]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:24
			_go_fuzz_dep_.CoverTab[106985]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:24
			// _ = "end of CoverTab[106985]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:24
		// _ = "end of CoverTab[106980]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:24
		_go_fuzz_dep_.CoverTab[106981]++
														for _, partitionError := range e {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:25
			_go_fuzz_dep_.CoverTab[106986]++
															if err := partitionError.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:26
				_go_fuzz_dep_.CoverTab[106987]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:27
				// _ = "end of CoverTab[106987]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:28
				_go_fuzz_dep_.CoverTab[106988]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:28
				// _ = "end of CoverTab[106988]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:28
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:28
			// _ = "end of CoverTab[106986]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:29
		// _ = "end of CoverTab[106981]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:30
	// _ = "end of CoverTab[106975]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:30
	_go_fuzz_dep_.CoverTab[106976]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:32
	// _ = "end of CoverTab[106976]"
}

func (t *TxnOffsetCommitResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:35
	_go_fuzz_dep_.CoverTab[106989]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:37
		_go_fuzz_dep_.CoverTab[106993]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:38
		// _ = "end of CoverTab[106993]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:39
		_go_fuzz_dep_.CoverTab[106994]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:39
		// _ = "end of CoverTab[106994]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:39
	// _ = "end of CoverTab[106989]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:39
	_go_fuzz_dep_.CoverTab[106990]++
													t.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:43
		_go_fuzz_dep_.CoverTab[106995]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:44
		// _ = "end of CoverTab[106995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:45
		_go_fuzz_dep_.CoverTab[106996]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:45
		// _ = "end of CoverTab[106996]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:45
	// _ = "end of CoverTab[106990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:45
	_go_fuzz_dep_.CoverTab[106991]++

													t.Topics = make(map[string][]*PartitionError)

													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:49
		_go_fuzz_dep_.CoverTab[106997]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:51
			_go_fuzz_dep_.CoverTab[107000]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:52
			// _ = "end of CoverTab[107000]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:53
			_go_fuzz_dep_.CoverTab[107001]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:53
			// _ = "end of CoverTab[107001]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:53
		// _ = "end of CoverTab[106997]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:53
		_go_fuzz_dep_.CoverTab[106998]++

														m, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:56
			_go_fuzz_dep_.CoverTab[107002]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:57
			// _ = "end of CoverTab[107002]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:58
			_go_fuzz_dep_.CoverTab[107003]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:58
			// _ = "end of CoverTab[107003]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:58
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:58
		// _ = "end of CoverTab[106998]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:58
		_go_fuzz_dep_.CoverTab[106999]++

														t.Topics[topic] = make([]*PartitionError, m)

														for j := 0; j < m; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:62
			_go_fuzz_dep_.CoverTab[107004]++
															t.Topics[topic][j] = new(PartitionError)
															if err := t.Topics[topic][j].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:64
				_go_fuzz_dep_.CoverTab[107005]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:65
				// _ = "end of CoverTab[107005]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:66
				_go_fuzz_dep_.CoverTab[107006]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:66
				// _ = "end of CoverTab[107006]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:66
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:66
			// _ = "end of CoverTab[107004]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:67
		// _ = "end of CoverTab[106999]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:68
	// _ = "end of CoverTab[106991]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:68
	_go_fuzz_dep_.CoverTab[106992]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:70
	// _ = "end of CoverTab[106992]"
}

func (a *TxnOffsetCommitResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:73
	_go_fuzz_dep_.CoverTab[107007]++
													return 28
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:74
	// _ = "end of CoverTab[107007]"
}

func (a *TxnOffsetCommitResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:77
	_go_fuzz_dep_.CoverTab[107008]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:78
	// _ = "end of CoverTab[107008]"
}

func (a *TxnOffsetCommitResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:81
	_go_fuzz_dep_.CoverTab[107009]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:82
	// _ = "end of CoverTab[107009]"
}

func (a *TxnOffsetCommitResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:85
	_go_fuzz_dep_.CoverTab[107010]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:86
	// _ = "end of CoverTab[107010]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_response.go:87
var _ = _go_fuzz_dep_.CoverTab
