//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:1
)

import (
	"sort"
	"time"
)

const invalidPreferredReplicaID = -1

type AbortedTransaction struct {
	ProducerID	int64
	FirstOffset	int64
}

func (t *AbortedTransaction) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:15
	_go_fuzz_dep_.CoverTab[103017]++
												if t.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:16
		_go_fuzz_dep_.CoverTab[103020]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:17
		// _ = "end of CoverTab[103020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:18
		_go_fuzz_dep_.CoverTab[103021]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:18
		// _ = "end of CoverTab[103021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:18
	// _ = "end of CoverTab[103017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:18
	_go_fuzz_dep_.CoverTab[103018]++

												if t.FirstOffset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:20
		_go_fuzz_dep_.CoverTab[103022]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:21
		// _ = "end of CoverTab[103022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:22
		_go_fuzz_dep_.CoverTab[103023]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:22
		// _ = "end of CoverTab[103023]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:22
	// _ = "end of CoverTab[103018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:22
	_go_fuzz_dep_.CoverTab[103019]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:24
	// _ = "end of CoverTab[103019]"
}

func (t *AbortedTransaction) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:27
	_go_fuzz_dep_.CoverTab[103024]++
												pe.putInt64(t.ProducerID)
												pe.putInt64(t.FirstOffset)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:31
	// _ = "end of CoverTab[103024]"
}

type FetchResponseBlock struct {
	Err			KError
	HighWaterMarkOffset	int64
	LastStableOffset	int64
	LastRecordsBatchOffset	*int64
	LogStartOffset		int64
	AbortedTransactions	[]*AbortedTransaction
	PreferredReadReplica	int32
	Records			*Records	// deprecated: use FetchResponseBlock.RecordsSet
	RecordsSet		[]*Records
	Partial			bool
}

func (b *FetchResponseBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:47
	_go_fuzz_dep_.CoverTab[103025]++
												tmp, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:49
		_go_fuzz_dep_.CoverTab[103033]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:50
		// _ = "end of CoverTab[103033]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:51
		_go_fuzz_dep_.CoverTab[103034]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:51
		// _ = "end of CoverTab[103034]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:51
	// _ = "end of CoverTab[103025]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:51
	_go_fuzz_dep_.CoverTab[103026]++
												b.Err = KError(tmp)

												b.HighWaterMarkOffset, err = pd.getInt64()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:55
		_go_fuzz_dep_.CoverTab[103035]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:56
		// _ = "end of CoverTab[103035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:57
		_go_fuzz_dep_.CoverTab[103036]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:57
		// _ = "end of CoverTab[103036]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:57
	// _ = "end of CoverTab[103026]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:57
	_go_fuzz_dep_.CoverTab[103027]++

												if version >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:59
		_go_fuzz_dep_.CoverTab[103037]++
													b.LastStableOffset, err = pd.getInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:61
			_go_fuzz_dep_.CoverTab[103042]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:62
			// _ = "end of CoverTab[103042]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:63
			_go_fuzz_dep_.CoverTab[103043]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:63
			// _ = "end of CoverTab[103043]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:63
		// _ = "end of CoverTab[103037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:63
		_go_fuzz_dep_.CoverTab[103038]++

													if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:65
			_go_fuzz_dep_.CoverTab[103044]++
														b.LogStartOffset, err = pd.getInt64()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:67
				_go_fuzz_dep_.CoverTab[103045]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:68
				// _ = "end of CoverTab[103045]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:69
				_go_fuzz_dep_.CoverTab[103046]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:69
				// _ = "end of CoverTab[103046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:69
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:69
			// _ = "end of CoverTab[103044]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:70
			_go_fuzz_dep_.CoverTab[103047]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:70
			// _ = "end of CoverTab[103047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:70
		// _ = "end of CoverTab[103038]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:70
		_go_fuzz_dep_.CoverTab[103039]++

													numTransact, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:73
			_go_fuzz_dep_.CoverTab[103048]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:74
			// _ = "end of CoverTab[103048]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:75
			_go_fuzz_dep_.CoverTab[103049]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:75
			// _ = "end of CoverTab[103049]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:75
		// _ = "end of CoverTab[103039]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:75
		_go_fuzz_dep_.CoverTab[103040]++

													if numTransact >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:77
			_go_fuzz_dep_.CoverTab[103050]++
														b.AbortedTransactions = make([]*AbortedTransaction, numTransact)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:78
			// _ = "end of CoverTab[103050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:79
			_go_fuzz_dep_.CoverTab[103051]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:79
			// _ = "end of CoverTab[103051]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:79
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:79
		// _ = "end of CoverTab[103040]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:79
		_go_fuzz_dep_.CoverTab[103041]++

													for i := 0; i < numTransact; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:81
			_go_fuzz_dep_.CoverTab[103052]++
														transact := new(AbortedTransaction)
														if err = transact.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:83
				_go_fuzz_dep_.CoverTab[103054]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:84
				// _ = "end of CoverTab[103054]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:85
				_go_fuzz_dep_.CoverTab[103055]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:85
				// _ = "end of CoverTab[103055]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:85
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:85
			// _ = "end of CoverTab[103052]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:85
			_go_fuzz_dep_.CoverTab[103053]++
														b.AbortedTransactions[i] = transact
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:86
			// _ = "end of CoverTab[103053]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:87
		// _ = "end of CoverTab[103041]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:88
		_go_fuzz_dep_.CoverTab[103056]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:88
		// _ = "end of CoverTab[103056]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:88
	// _ = "end of CoverTab[103027]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:88
	_go_fuzz_dep_.CoverTab[103028]++

												if version >= 11 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:90
		_go_fuzz_dep_.CoverTab[103057]++
													b.PreferredReadReplica, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:92
			_go_fuzz_dep_.CoverTab[103058]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:93
			// _ = "end of CoverTab[103058]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:94
			_go_fuzz_dep_.CoverTab[103059]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:94
			// _ = "end of CoverTab[103059]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:94
		// _ = "end of CoverTab[103057]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:95
		_go_fuzz_dep_.CoverTab[103060]++
													b.PreferredReadReplica = -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:96
		// _ = "end of CoverTab[103060]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:97
	// _ = "end of CoverTab[103028]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:97
	_go_fuzz_dep_.CoverTab[103029]++

												recordsSize, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:100
		_go_fuzz_dep_.CoverTab[103061]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:101
		// _ = "end of CoverTab[103061]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:102
		_go_fuzz_dep_.CoverTab[103062]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:102
		// _ = "end of CoverTab[103062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:102
	// _ = "end of CoverTab[103029]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:102
	_go_fuzz_dep_.CoverTab[103030]++

												recordsDecoder, err := pd.getSubset(int(recordsSize))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:105
		_go_fuzz_dep_.CoverTab[103063]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:106
		// _ = "end of CoverTab[103063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:107
		_go_fuzz_dep_.CoverTab[103064]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:107
		// _ = "end of CoverTab[103064]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:107
	// _ = "end of CoverTab[103030]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:107
	_go_fuzz_dep_.CoverTab[103031]++

												b.RecordsSet = []*Records{}

												for recordsDecoder.remaining() > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:111
		_go_fuzz_dep_.CoverTab[103065]++
													records := &Records{}
													if err := records.decode(recordsDecoder); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:113
			_go_fuzz_dep_.CoverTab[103072]++

														if err == ErrInsufficientData {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:115
				_go_fuzz_dep_.CoverTab[103074]++
															if len(b.RecordsSet) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:116
					_go_fuzz_dep_.CoverTab[103076]++
																b.Partial = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:117
					// _ = "end of CoverTab[103076]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:118
					_go_fuzz_dep_.CoverTab[103077]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:118
					// _ = "end of CoverTab[103077]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:118
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:118
				// _ = "end of CoverTab[103074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:118
				_go_fuzz_dep_.CoverTab[103075]++
															break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:119
				// _ = "end of CoverTab[103075]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:120
				_go_fuzz_dep_.CoverTab[103078]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:120
				// _ = "end of CoverTab[103078]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:120
			// _ = "end of CoverTab[103072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:120
			_go_fuzz_dep_.CoverTab[103073]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:121
			// _ = "end of CoverTab[103073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:122
			_go_fuzz_dep_.CoverTab[103079]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:122
			// _ = "end of CoverTab[103079]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:122
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:122
		// _ = "end of CoverTab[103065]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:122
		_go_fuzz_dep_.CoverTab[103066]++

													b.LastRecordsBatchOffset, err = records.recordsOffset()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:125
			_go_fuzz_dep_.CoverTab[103080]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:126
			// _ = "end of CoverTab[103080]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:127
			_go_fuzz_dep_.CoverTab[103081]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:127
			// _ = "end of CoverTab[103081]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:127
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:127
		// _ = "end of CoverTab[103066]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:127
		_go_fuzz_dep_.CoverTab[103067]++

													partial, err := records.isPartial()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:130
			_go_fuzz_dep_.CoverTab[103082]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:131
			// _ = "end of CoverTab[103082]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:132
			_go_fuzz_dep_.CoverTab[103083]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:132
			// _ = "end of CoverTab[103083]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:132
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:132
		// _ = "end of CoverTab[103067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:132
		_go_fuzz_dep_.CoverTab[103068]++

													n, err := records.numRecords()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:135
			_go_fuzz_dep_.CoverTab[103084]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:136
			// _ = "end of CoverTab[103084]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:137
			_go_fuzz_dep_.CoverTab[103085]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:137
			// _ = "end of CoverTab[103085]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:137
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:137
		// _ = "end of CoverTab[103068]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:137
		_go_fuzz_dep_.CoverTab[103069]++

													if n > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
			_go_fuzz_dep_.CoverTab[103086]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
			return (partial && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
				_go_fuzz_dep_.CoverTab[103087]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
				return len(b.RecordsSet) == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
				// _ = "end of CoverTab[103087]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
			}())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
			// _ = "end of CoverTab[103086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:139
			_go_fuzz_dep_.CoverTab[103088]++
														b.RecordsSet = append(b.RecordsSet, records)

														if b.Records == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:142
				_go_fuzz_dep_.CoverTab[103089]++
															b.Records = records
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:143
				// _ = "end of CoverTab[103089]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:144
				_go_fuzz_dep_.CoverTab[103090]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:144
				// _ = "end of CoverTab[103090]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:144
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:144
			// _ = "end of CoverTab[103088]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:145
			_go_fuzz_dep_.CoverTab[103091]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:145
			// _ = "end of CoverTab[103091]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:145
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:145
		// _ = "end of CoverTab[103069]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:145
		_go_fuzz_dep_.CoverTab[103070]++

													overflow, err := records.isOverflow()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:148
			_go_fuzz_dep_.CoverTab[103092]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:149
			// _ = "end of CoverTab[103092]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:150
			_go_fuzz_dep_.CoverTab[103093]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:150
			// _ = "end of CoverTab[103093]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:150
		// _ = "end of CoverTab[103070]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:150
		_go_fuzz_dep_.CoverTab[103071]++

													if partial || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:152
			_go_fuzz_dep_.CoverTab[103094]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:152
			return overflow
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:152
			// _ = "end of CoverTab[103094]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:152
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:152
			_go_fuzz_dep_.CoverTab[103095]++
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:153
			// _ = "end of CoverTab[103095]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:154
			_go_fuzz_dep_.CoverTab[103096]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:154
			// _ = "end of CoverTab[103096]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:154
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:154
		// _ = "end of CoverTab[103071]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:155
	// _ = "end of CoverTab[103031]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:155
	_go_fuzz_dep_.CoverTab[103032]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:157
	// _ = "end of CoverTab[103032]"
}

func (b *FetchResponseBlock) numRecords() (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:160
	_go_fuzz_dep_.CoverTab[103097]++
												sum := 0

												for _, records := range b.RecordsSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:163
		_go_fuzz_dep_.CoverTab[103099]++
													count, err := records.numRecords()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:165
			_go_fuzz_dep_.CoverTab[103101]++
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:166
			// _ = "end of CoverTab[103101]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:167
			_go_fuzz_dep_.CoverTab[103102]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:167
			// _ = "end of CoverTab[103102]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:167
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:167
		// _ = "end of CoverTab[103099]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:167
		_go_fuzz_dep_.CoverTab[103100]++

													sum += count
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:169
		// _ = "end of CoverTab[103100]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:170
	// _ = "end of CoverTab[103097]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:170
	_go_fuzz_dep_.CoverTab[103098]++

												return sum, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:172
	// _ = "end of CoverTab[103098]"
}

func (b *FetchResponseBlock) isPartial() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:175
	_go_fuzz_dep_.CoverTab[103103]++
												if b.Partial {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:176
		_go_fuzz_dep_.CoverTab[103106]++
													return true, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:177
		// _ = "end of CoverTab[103106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:178
		_go_fuzz_dep_.CoverTab[103107]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:178
		// _ = "end of CoverTab[103107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:178
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:178
	// _ = "end of CoverTab[103103]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:178
	_go_fuzz_dep_.CoverTab[103104]++

												if len(b.RecordsSet) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:180
		_go_fuzz_dep_.CoverTab[103108]++
													return b.RecordsSet[0].isPartial()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:181
		// _ = "end of CoverTab[103108]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:182
		_go_fuzz_dep_.CoverTab[103109]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:182
		// _ = "end of CoverTab[103109]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:182
	// _ = "end of CoverTab[103104]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:182
	_go_fuzz_dep_.CoverTab[103105]++

												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:184
	// _ = "end of CoverTab[103105]"
}

func (b *FetchResponseBlock) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:187
	_go_fuzz_dep_.CoverTab[103110]++
												pe.putInt16(int16(b.Err))

												pe.putInt64(b.HighWaterMarkOffset)

												if version >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:192
		_go_fuzz_dep_.CoverTab[103114]++
													pe.putInt64(b.LastStableOffset)

													if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:195
			_go_fuzz_dep_.CoverTab[103117]++
														pe.putInt64(b.LogStartOffset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:196
			// _ = "end of CoverTab[103117]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:197
			_go_fuzz_dep_.CoverTab[103118]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:197
			// _ = "end of CoverTab[103118]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:197
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:197
		// _ = "end of CoverTab[103114]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:197
		_go_fuzz_dep_.CoverTab[103115]++

													if err = pe.putArrayLength(len(b.AbortedTransactions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:199
			_go_fuzz_dep_.CoverTab[103119]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:200
			// _ = "end of CoverTab[103119]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:201
			_go_fuzz_dep_.CoverTab[103120]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:201
			// _ = "end of CoverTab[103120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:201
		// _ = "end of CoverTab[103115]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:201
		_go_fuzz_dep_.CoverTab[103116]++
													for _, transact := range b.AbortedTransactions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:202
			_go_fuzz_dep_.CoverTab[103121]++
														if err = transact.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:203
				_go_fuzz_dep_.CoverTab[103122]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:204
				// _ = "end of CoverTab[103122]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:205
				_go_fuzz_dep_.CoverTab[103123]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:205
				// _ = "end of CoverTab[103123]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:205
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:205
			// _ = "end of CoverTab[103121]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:206
		// _ = "end of CoverTab[103116]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:207
		_go_fuzz_dep_.CoverTab[103124]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:207
		// _ = "end of CoverTab[103124]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:207
	// _ = "end of CoverTab[103110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:207
	_go_fuzz_dep_.CoverTab[103111]++

												if version >= 11 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:209
		_go_fuzz_dep_.CoverTab[103125]++
													pe.putInt32(b.PreferredReadReplica)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:210
		// _ = "end of CoverTab[103125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:211
		_go_fuzz_dep_.CoverTab[103126]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:211
		// _ = "end of CoverTab[103126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:211
	// _ = "end of CoverTab[103111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:211
	_go_fuzz_dep_.CoverTab[103112]++

												pe.push(&lengthField{})
												for _, records := range b.RecordsSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:214
		_go_fuzz_dep_.CoverTab[103127]++
													err = records.encode(pe)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:216
			_go_fuzz_dep_.CoverTab[103128]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:217
			// _ = "end of CoverTab[103128]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:218
			_go_fuzz_dep_.CoverTab[103129]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:218
			// _ = "end of CoverTab[103129]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:218
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:218
		// _ = "end of CoverTab[103127]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:219
	// _ = "end of CoverTab[103112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:219
	_go_fuzz_dep_.CoverTab[103113]++
												return pe.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:220
	// _ = "end of CoverTab[103113]"
}

func (b *FetchResponseBlock) getAbortedTransactions() []*AbortedTransaction {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:223
	_go_fuzz_dep_.CoverTab[103130]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:226
	at := b.AbortedTransactions
	sort.Slice(
		at,
		func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:229
			_go_fuzz_dep_.CoverTab[103132]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:229
			return at[i].FirstOffset < at[j].FirstOffset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:229
			// _ = "end of CoverTab[103132]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:229
		},
	)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:230
	// _ = "end of CoverTab[103130]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:230
	_go_fuzz_dep_.CoverTab[103131]++
												return at
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:231
	// _ = "end of CoverTab[103131]"
}

type FetchResponse struct {
	Blocks		map[string]map[int32]*FetchResponseBlock
	ThrottleTime	time.Duration
	ErrorCode	int16
	SessionID	int32
	Version		int16
	LogAppendTime	bool
	Timestamp	time.Time
}

func (r *FetchResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:244
	_go_fuzz_dep_.CoverTab[103133]++
												r.Version = version

												if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:247
		_go_fuzz_dep_.CoverTab[103138]++
													throttle, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:249
			_go_fuzz_dep_.CoverTab[103140]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:250
			// _ = "end of CoverTab[103140]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:251
			_go_fuzz_dep_.CoverTab[103141]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:251
			// _ = "end of CoverTab[103141]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:251
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:251
		// _ = "end of CoverTab[103138]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:251
		_go_fuzz_dep_.CoverTab[103139]++
													r.ThrottleTime = time.Duration(throttle) * time.Millisecond
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:252
		// _ = "end of CoverTab[103139]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:253
		_go_fuzz_dep_.CoverTab[103142]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:253
		// _ = "end of CoverTab[103142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:253
	// _ = "end of CoverTab[103133]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:253
	_go_fuzz_dep_.CoverTab[103134]++

												if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:255
		_go_fuzz_dep_.CoverTab[103143]++
													r.ErrorCode, err = pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:257
			_go_fuzz_dep_.CoverTab[103145]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:258
			// _ = "end of CoverTab[103145]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:259
			_go_fuzz_dep_.CoverTab[103146]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:259
			// _ = "end of CoverTab[103146]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:259
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:259
		// _ = "end of CoverTab[103143]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:259
		_go_fuzz_dep_.CoverTab[103144]++
													r.SessionID, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:261
			_go_fuzz_dep_.CoverTab[103147]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:262
			// _ = "end of CoverTab[103147]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:263
			_go_fuzz_dep_.CoverTab[103148]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:263
			// _ = "end of CoverTab[103148]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:263
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:263
		// _ = "end of CoverTab[103144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:264
		_go_fuzz_dep_.CoverTab[103149]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:264
		// _ = "end of CoverTab[103149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:264
	// _ = "end of CoverTab[103134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:264
	_go_fuzz_dep_.CoverTab[103135]++

												numTopics, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:267
		_go_fuzz_dep_.CoverTab[103150]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:268
		// _ = "end of CoverTab[103150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:269
		_go_fuzz_dep_.CoverTab[103151]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:269
		// _ = "end of CoverTab[103151]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:269
	// _ = "end of CoverTab[103135]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:269
	_go_fuzz_dep_.CoverTab[103136]++

												r.Blocks = make(map[string]map[int32]*FetchResponseBlock, numTopics)
												for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:272
		_go_fuzz_dep_.CoverTab[103152]++
													name, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:274
			_go_fuzz_dep_.CoverTab[103155]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:275
			// _ = "end of CoverTab[103155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:276
			_go_fuzz_dep_.CoverTab[103156]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:276
			// _ = "end of CoverTab[103156]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:276
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:276
		// _ = "end of CoverTab[103152]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:276
		_go_fuzz_dep_.CoverTab[103153]++

													numBlocks, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:279
			_go_fuzz_dep_.CoverTab[103157]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:280
			// _ = "end of CoverTab[103157]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:281
			_go_fuzz_dep_.CoverTab[103158]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:281
			// _ = "end of CoverTab[103158]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:281
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:281
		// _ = "end of CoverTab[103153]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:281
		_go_fuzz_dep_.CoverTab[103154]++

													r.Blocks[name] = make(map[int32]*FetchResponseBlock, numBlocks)

													for j := 0; j < numBlocks; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:285
			_go_fuzz_dep_.CoverTab[103159]++
														id, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:287
				_go_fuzz_dep_.CoverTab[103162]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:288
				// _ = "end of CoverTab[103162]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:289
				_go_fuzz_dep_.CoverTab[103163]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:289
				// _ = "end of CoverTab[103163]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:289
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:289
			// _ = "end of CoverTab[103159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:289
			_go_fuzz_dep_.CoverTab[103160]++

														block := new(FetchResponseBlock)
														err = block.decode(pd, version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:293
				_go_fuzz_dep_.CoverTab[103164]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:294
				// _ = "end of CoverTab[103164]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:295
				_go_fuzz_dep_.CoverTab[103165]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:295
				// _ = "end of CoverTab[103165]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:295
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:295
			// _ = "end of CoverTab[103160]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:295
			_go_fuzz_dep_.CoverTab[103161]++
														r.Blocks[name][id] = block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:296
			// _ = "end of CoverTab[103161]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:297
		// _ = "end of CoverTab[103154]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:298
	// _ = "end of CoverTab[103136]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:298
	_go_fuzz_dep_.CoverTab[103137]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:300
	// _ = "end of CoverTab[103137]"
}

func (r *FetchResponse) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:303
	_go_fuzz_dep_.CoverTab[103166]++
												if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:304
		_go_fuzz_dep_.CoverTab[103171]++
													pe.putInt32(int32(r.ThrottleTime / time.Millisecond))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:305
		// _ = "end of CoverTab[103171]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:306
		_go_fuzz_dep_.CoverTab[103172]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:306
		// _ = "end of CoverTab[103172]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:306
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:306
	// _ = "end of CoverTab[103166]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:306
	_go_fuzz_dep_.CoverTab[103167]++

												if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:308
		_go_fuzz_dep_.CoverTab[103173]++
													pe.putInt16(r.ErrorCode)
													pe.putInt32(r.SessionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:310
		// _ = "end of CoverTab[103173]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:311
		_go_fuzz_dep_.CoverTab[103174]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:311
		// _ = "end of CoverTab[103174]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:311
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:311
	// _ = "end of CoverTab[103167]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:311
	_go_fuzz_dep_.CoverTab[103168]++

												err = pe.putArrayLength(len(r.Blocks))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:314
		_go_fuzz_dep_.CoverTab[103175]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:315
		// _ = "end of CoverTab[103175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:316
		_go_fuzz_dep_.CoverTab[103176]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:316
		// _ = "end of CoverTab[103176]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:316
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:316
	// _ = "end of CoverTab[103168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:316
	_go_fuzz_dep_.CoverTab[103169]++

												for topic, partitions := range r.Blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:318
		_go_fuzz_dep_.CoverTab[103177]++
													err = pe.putString(topic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:320
			_go_fuzz_dep_.CoverTab[103180]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:321
			// _ = "end of CoverTab[103180]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:322
			_go_fuzz_dep_.CoverTab[103181]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:322
			// _ = "end of CoverTab[103181]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:322
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:322
		// _ = "end of CoverTab[103177]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:322
		_go_fuzz_dep_.CoverTab[103178]++

													err = pe.putArrayLength(len(partitions))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:325
			_go_fuzz_dep_.CoverTab[103182]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:326
			// _ = "end of CoverTab[103182]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:327
			_go_fuzz_dep_.CoverTab[103183]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:327
			// _ = "end of CoverTab[103183]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:327
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:327
		// _ = "end of CoverTab[103178]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:327
		_go_fuzz_dep_.CoverTab[103179]++

													for id, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:329
			_go_fuzz_dep_.CoverTab[103184]++
														pe.putInt32(id)
														err = block.encode(pe, r.Version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:332
				_go_fuzz_dep_.CoverTab[103185]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:333
				// _ = "end of CoverTab[103185]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:334
				_go_fuzz_dep_.CoverTab[103186]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:334
				// _ = "end of CoverTab[103186]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:334
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:334
			// _ = "end of CoverTab[103184]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:335
		// _ = "end of CoverTab[103179]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:336
	// _ = "end of CoverTab[103169]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:336
	_go_fuzz_dep_.CoverTab[103170]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:337
	// _ = "end of CoverTab[103170]"
}

func (r *FetchResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:340
	_go_fuzz_dep_.CoverTab[103187]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:341
	// _ = "end of CoverTab[103187]"
}

func (r *FetchResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:344
	_go_fuzz_dep_.CoverTab[103188]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:345
	// _ = "end of CoverTab[103188]"
}

func (r *FetchResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:348
	_go_fuzz_dep_.CoverTab[103189]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:349
	// _ = "end of CoverTab[103189]"
}

func (r *FetchResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:352
	_go_fuzz_dep_.CoverTab[103190]++
												switch r.Version {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:354
		_go_fuzz_dep_.CoverTab[103191]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:355
		// _ = "end of CoverTab[103191]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:356
		_go_fuzz_dep_.CoverTab[103192]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:357
		// _ = "end of CoverTab[103192]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:358
		_go_fuzz_dep_.CoverTab[103193]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:359
		// _ = "end of CoverTab[103193]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:360
		_go_fuzz_dep_.CoverTab[103194]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:361
		// _ = "end of CoverTab[103194]"
	case 4, 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:362
		_go_fuzz_dep_.CoverTab[103195]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:363
		// _ = "end of CoverTab[103195]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:364
		_go_fuzz_dep_.CoverTab[103196]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:365
		// _ = "end of CoverTab[103196]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:366
		_go_fuzz_dep_.CoverTab[103197]++
													return V1_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:367
		// _ = "end of CoverTab[103197]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:368
		_go_fuzz_dep_.CoverTab[103198]++
													return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:369
		// _ = "end of CoverTab[103198]"
	case 9, 10:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:370
		_go_fuzz_dep_.CoverTab[103199]++
													return V2_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:371
		// _ = "end of CoverTab[103199]"
	case 11:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:372
		_go_fuzz_dep_.CoverTab[103200]++
													return V2_3_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:373
		// _ = "end of CoverTab[103200]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:374
		_go_fuzz_dep_.CoverTab[103201]++
													return MaxVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:375
		// _ = "end of CoverTab[103201]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:376
	// _ = "end of CoverTab[103190]"
}

func (r *FetchResponse) GetBlock(topic string, partition int32) *FetchResponseBlock {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:379
	_go_fuzz_dep_.CoverTab[103202]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:380
		_go_fuzz_dep_.CoverTab[103205]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:381
		// _ = "end of CoverTab[103205]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:382
		_go_fuzz_dep_.CoverTab[103206]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:382
		// _ = "end of CoverTab[103206]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:382
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:382
	// _ = "end of CoverTab[103202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:382
	_go_fuzz_dep_.CoverTab[103203]++

												if r.Blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:384
		_go_fuzz_dep_.CoverTab[103207]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:385
		// _ = "end of CoverTab[103207]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:386
		_go_fuzz_dep_.CoverTab[103208]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:386
		// _ = "end of CoverTab[103208]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:386
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:386
	// _ = "end of CoverTab[103203]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:386
	_go_fuzz_dep_.CoverTab[103204]++

												return r.Blocks[topic][partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:388
	// _ = "end of CoverTab[103204]"
}

func (r *FetchResponse) AddError(topic string, partition int32, err KError) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:391
	_go_fuzz_dep_.CoverTab[103209]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:392
		_go_fuzz_dep_.CoverTab[103213]++
													r.Blocks = make(map[string]map[int32]*FetchResponseBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:393
		// _ = "end of CoverTab[103213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:394
		_go_fuzz_dep_.CoverTab[103214]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:394
		// _ = "end of CoverTab[103214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:394
	// _ = "end of CoverTab[103209]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:394
	_go_fuzz_dep_.CoverTab[103210]++
												partitions, ok := r.Blocks[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:396
		_go_fuzz_dep_.CoverTab[103215]++
													partitions = make(map[int32]*FetchResponseBlock)
													r.Blocks[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:398
		// _ = "end of CoverTab[103215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:399
		_go_fuzz_dep_.CoverTab[103216]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:399
		// _ = "end of CoverTab[103216]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:399
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:399
	// _ = "end of CoverTab[103210]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:399
	_go_fuzz_dep_.CoverTab[103211]++
												frb, ok := partitions[partition]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:401
		_go_fuzz_dep_.CoverTab[103217]++
													frb = new(FetchResponseBlock)
													partitions[partition] = frb
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:403
		// _ = "end of CoverTab[103217]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:404
		_go_fuzz_dep_.CoverTab[103218]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:404
		// _ = "end of CoverTab[103218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:404
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:404
	// _ = "end of CoverTab[103211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:404
	_go_fuzz_dep_.CoverTab[103212]++
												frb.Err = err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:405
	// _ = "end of CoverTab[103212]"
}

func (r *FetchResponse) getOrCreateBlock(topic string, partition int32) *FetchResponseBlock {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:408
	_go_fuzz_dep_.CoverTab[103219]++
												if r.Blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:409
		_go_fuzz_dep_.CoverTab[103223]++
													r.Blocks = make(map[string]map[int32]*FetchResponseBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:410
		// _ = "end of CoverTab[103223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:411
		_go_fuzz_dep_.CoverTab[103224]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:411
		// _ = "end of CoverTab[103224]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:411
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:411
	// _ = "end of CoverTab[103219]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:411
	_go_fuzz_dep_.CoverTab[103220]++
												partitions, ok := r.Blocks[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:413
		_go_fuzz_dep_.CoverTab[103225]++
													partitions = make(map[int32]*FetchResponseBlock)
													r.Blocks[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:415
		// _ = "end of CoverTab[103225]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:416
		_go_fuzz_dep_.CoverTab[103226]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:416
		// _ = "end of CoverTab[103226]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:416
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:416
	// _ = "end of CoverTab[103220]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:416
	_go_fuzz_dep_.CoverTab[103221]++
												frb, ok := partitions[partition]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:418
		_go_fuzz_dep_.CoverTab[103227]++
													frb = new(FetchResponseBlock)
													partitions[partition] = frb
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:420
		// _ = "end of CoverTab[103227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:421
		_go_fuzz_dep_.CoverTab[103228]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:421
		// _ = "end of CoverTab[103228]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:421
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:421
	// _ = "end of CoverTab[103221]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:421
	_go_fuzz_dep_.CoverTab[103222]++

												return frb
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:423
	// _ = "end of CoverTab[103222]"
}

func encodeKV(key, value Encoder) ([]byte, []byte) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:426
	_go_fuzz_dep_.CoverTab[103229]++
												var kb []byte
												var vb []byte
												if key != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:429
		_go_fuzz_dep_.CoverTab[103232]++
													kb, _ = key.Encode()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:430
		// _ = "end of CoverTab[103232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:431
		_go_fuzz_dep_.CoverTab[103233]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:431
		// _ = "end of CoverTab[103233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:431
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:431
	// _ = "end of CoverTab[103229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:431
	_go_fuzz_dep_.CoverTab[103230]++
												if value != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:432
		_go_fuzz_dep_.CoverTab[103234]++
													vb, _ = value.Encode()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:433
		// _ = "end of CoverTab[103234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:434
		_go_fuzz_dep_.CoverTab[103235]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:434
		// _ = "end of CoverTab[103235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:434
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:434
	// _ = "end of CoverTab[103230]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:434
	_go_fuzz_dep_.CoverTab[103231]++

												return kb, vb
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:436
	// _ = "end of CoverTab[103231]"
}

func (r *FetchResponse) AddMessageWithTimestamp(topic string, partition int32, key, value Encoder, offset int64, timestamp time.Time, version int8) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:439
	_go_fuzz_dep_.CoverTab[103236]++
												frb := r.getOrCreateBlock(topic, partition)
												kb, vb := encodeKV(key, value)
												if r.LogAppendTime {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:442
		_go_fuzz_dep_.CoverTab[103239]++
													timestamp = r.Timestamp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:443
		// _ = "end of CoverTab[103239]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:444
		_go_fuzz_dep_.CoverTab[103240]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:444
		// _ = "end of CoverTab[103240]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:444
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:444
	// _ = "end of CoverTab[103236]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:444
	_go_fuzz_dep_.CoverTab[103237]++
												msg := &Message{Key: kb, Value: vb, LogAppendTime: r.LogAppendTime, Timestamp: timestamp, Version: version}
												msgBlock := &MessageBlock{Msg: msg, Offset: offset}
												if len(frb.RecordsSet) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:447
		_go_fuzz_dep_.CoverTab[103241]++
													records := newLegacyRecords(&MessageSet{})
													frb.RecordsSet = []*Records{&records}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:449
		// _ = "end of CoverTab[103241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:450
		_go_fuzz_dep_.CoverTab[103242]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:450
		// _ = "end of CoverTab[103242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:450
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:450
	// _ = "end of CoverTab[103237]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:450
	_go_fuzz_dep_.CoverTab[103238]++
												set := frb.RecordsSet[0].MsgSet
												set.Messages = append(set.Messages, msgBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:452
	// _ = "end of CoverTab[103238]"
}

func (r *FetchResponse) AddRecordWithTimestamp(topic string, partition int32, key, value Encoder, offset int64, timestamp time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:455
	_go_fuzz_dep_.CoverTab[103243]++
												frb := r.getOrCreateBlock(topic, partition)
												kb, vb := encodeKV(key, value)
												if len(frb.RecordsSet) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:458
		_go_fuzz_dep_.CoverTab[103245]++
													records := newDefaultRecords(&RecordBatch{Version: 2, LogAppendTime: r.LogAppendTime, FirstTimestamp: timestamp, MaxTimestamp: r.Timestamp})
													frb.RecordsSet = []*Records{&records}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:460
		// _ = "end of CoverTab[103245]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:461
		_go_fuzz_dep_.CoverTab[103246]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:461
		// _ = "end of CoverTab[103246]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:461
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:461
	// _ = "end of CoverTab[103243]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:461
	_go_fuzz_dep_.CoverTab[103244]++
												batch := frb.RecordsSet[0].RecordBatch
												rec := &Record{Key: kb, Value: vb, OffsetDelta: offset, TimestampDelta: timestamp.Sub(batch.FirstTimestamp)}
												batch.addRecord(rec)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:464
	// _ = "end of CoverTab[103244]"
}

// AddRecordBatchWithTimestamp is similar to AddRecordWithTimestamp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:467
// But instead of appending 1 record to a batch, it append a new batch containing 1 record to the fetchResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:467
// Since transaction are handled on batch level (the whole batch is either committed or aborted), use this to test transactions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:470
func (r *FetchResponse) AddRecordBatchWithTimestamp(topic string, partition int32, key, value Encoder, offset int64, producerID int64, isTransactional bool, timestamp time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:470
	_go_fuzz_dep_.CoverTab[103247]++
												frb := r.getOrCreateBlock(topic, partition)
												kb, vb := encodeKV(key, value)

												records := newDefaultRecords(&RecordBatch{Version: 2, LogAppendTime: r.LogAppendTime, FirstTimestamp: timestamp, MaxTimestamp: r.Timestamp})
												batch := &RecordBatch{
		Version:		2,
		LogAppendTime:		r.LogAppendTime,
		FirstTimestamp:		timestamp,
		MaxTimestamp:		r.Timestamp,
		FirstOffset:		offset,
		LastOffsetDelta:	0,
		ProducerID:		producerID,
		IsTransactional:	isTransactional,
	}
												rec := &Record{Key: kb, Value: vb, OffsetDelta: 0, TimestampDelta: timestamp.Sub(batch.FirstTimestamp)}
												batch.addRecord(rec)
												records.RecordBatch = batch

												frb.RecordsSet = append(frb.RecordsSet, &records)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:489
	// _ = "end of CoverTab[103247]"
}

func (r *FetchResponse) AddControlRecordWithTimestamp(topic string, partition int32, offset int64, producerID int64, recordType ControlRecordType, timestamp time.Time) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:492
	_go_fuzz_dep_.CoverTab[103248]++
												frb := r.getOrCreateBlock(topic, partition)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:496
	batch := &RecordBatch{
		Version:		2,
		LogAppendTime:		r.LogAppendTime,
		FirstTimestamp:		timestamp,
		MaxTimestamp:		r.Timestamp,
		FirstOffset:		offset,
		LastOffsetDelta:	0,
		ProducerID:		producerID,
		IsTransactional:	true,
		Control:		true,
	}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:509
	records := newDefaultRecords(nil)
												records.RecordBatch = batch

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:513
	crAbort := ControlRecord{
		Version:	0,
		Type:		recordType,
	}
												crKey := &realEncoder{raw: make([]byte, 4)}
												crValue := &realEncoder{raw: make([]byte, 6)}
												crAbort.encode(crKey, crValue)
												rec := &Record{Key: ByteEncoder(crKey.raw), Value: ByteEncoder(crValue.raw), OffsetDelta: 0, TimestampDelta: timestamp.Sub(batch.FirstTimestamp)}
												batch.addRecord(rec)

												frb.RecordsSet = append(frb.RecordsSet, &records)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:523
	// _ = "end of CoverTab[103248]"
}

func (r *FetchResponse) AddMessage(topic string, partition int32, key, value Encoder, offset int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:526
	_go_fuzz_dep_.CoverTab[103249]++
												r.AddMessageWithTimestamp(topic, partition, key, value, offset, time.Time{}, 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:527
	// _ = "end of CoverTab[103249]"
}

func (r *FetchResponse) AddRecord(topic string, partition int32, key, value Encoder, offset int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:530
	_go_fuzz_dep_.CoverTab[103250]++
												r.AddRecordWithTimestamp(topic, partition, key, value, offset, time.Time{})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:531
	// _ = "end of CoverTab[103250]"
}

func (r *FetchResponse) AddRecordBatch(topic string, partition int32, key, value Encoder, offset int64, producerID int64, isTransactional bool) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:534
	_go_fuzz_dep_.CoverTab[103251]++
												r.AddRecordBatchWithTimestamp(topic, partition, key, value, offset, producerID, isTransactional, time.Time{})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:535
	// _ = "end of CoverTab[103251]"
}

func (r *FetchResponse) AddControlRecord(topic string, partition int32, offset int64, producerID int64, recordType ControlRecordType) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:538
	_go_fuzz_dep_.CoverTab[103252]++

												r.AddControlRecordWithTimestamp(topic, partition, offset, producerID, recordType, time.Time{})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:540
	// _ = "end of CoverTab[103252]"
}

func (r *FetchResponse) SetLastOffsetDelta(topic string, partition int32, offset int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:543
	_go_fuzz_dep_.CoverTab[103253]++
												frb := r.getOrCreateBlock(topic, partition)
												if len(frb.RecordsSet) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:545
		_go_fuzz_dep_.CoverTab[103255]++
													records := newDefaultRecords(&RecordBatch{Version: 2})
													frb.RecordsSet = []*Records{&records}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:547
		// _ = "end of CoverTab[103255]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:548
		_go_fuzz_dep_.CoverTab[103256]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:548
		// _ = "end of CoverTab[103256]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:548
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:548
	// _ = "end of CoverTab[103253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:548
	_go_fuzz_dep_.CoverTab[103254]++
												batch := frb.RecordsSet[0].RecordBatch
												batch.LastOffsetDelta = offset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:550
	// _ = "end of CoverTab[103254]"
}

func (r *FetchResponse) SetLastStableOffset(topic string, partition int32, offset int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:553
	_go_fuzz_dep_.CoverTab[103257]++
												frb := r.getOrCreateBlock(topic, partition)
												frb.LastStableOffset = offset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:555
	// _ = "end of CoverTab[103257]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:556
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_response.go:556
var _ = _go_fuzz_dep_.CoverTab
