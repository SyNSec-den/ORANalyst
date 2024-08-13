//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:1
)

type TxnOffsetCommitRequest struct {
	TransactionalID	string
	GroupID		string
	ProducerID	int64
	ProducerEpoch	int16
	Topics		map[string][]*PartitionOffsetMetadata
}

func (t *TxnOffsetCommitRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:11
	_go_fuzz_dep_.CoverTab[106907]++
													if err := pe.putString(t.TransactionalID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:12
		_go_fuzz_dep_.CoverTab[106912]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:13
		// _ = "end of CoverTab[106912]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:14
		_go_fuzz_dep_.CoverTab[106913]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:14
		// _ = "end of CoverTab[106913]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:14
	// _ = "end of CoverTab[106907]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:14
	_go_fuzz_dep_.CoverTab[106908]++
													if err := pe.putString(t.GroupID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:15
		_go_fuzz_dep_.CoverTab[106914]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:16
		// _ = "end of CoverTab[106914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:17
		_go_fuzz_dep_.CoverTab[106915]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:17
		// _ = "end of CoverTab[106915]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:17
	// _ = "end of CoverTab[106908]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:17
	_go_fuzz_dep_.CoverTab[106909]++
													pe.putInt64(t.ProducerID)
													pe.putInt16(t.ProducerEpoch)

													if err := pe.putArrayLength(len(t.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:21
		_go_fuzz_dep_.CoverTab[106916]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:22
		// _ = "end of CoverTab[106916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:23
		_go_fuzz_dep_.CoverTab[106917]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:23
		// _ = "end of CoverTab[106917]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:23
	// _ = "end of CoverTab[106909]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:23
	_go_fuzz_dep_.CoverTab[106910]++
													for topic, partitions := range t.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:24
		_go_fuzz_dep_.CoverTab[106918]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:25
			_go_fuzz_dep_.CoverTab[106921]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:26
			// _ = "end of CoverTab[106921]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:27
			_go_fuzz_dep_.CoverTab[106922]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:27
			// _ = "end of CoverTab[106922]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:27
		// _ = "end of CoverTab[106918]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:27
		_go_fuzz_dep_.CoverTab[106919]++
														if err := pe.putArrayLength(len(partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:28
			_go_fuzz_dep_.CoverTab[106923]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:29
			// _ = "end of CoverTab[106923]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:30
			_go_fuzz_dep_.CoverTab[106924]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:30
			// _ = "end of CoverTab[106924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:30
		// _ = "end of CoverTab[106919]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:30
		_go_fuzz_dep_.CoverTab[106920]++
														for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:31
			_go_fuzz_dep_.CoverTab[106925]++
															if err := partition.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:32
				_go_fuzz_dep_.CoverTab[106926]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:33
				// _ = "end of CoverTab[106926]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:34
				_go_fuzz_dep_.CoverTab[106927]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:34
				// _ = "end of CoverTab[106927]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:34
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:34
			// _ = "end of CoverTab[106925]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:35
		// _ = "end of CoverTab[106920]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:36
	// _ = "end of CoverTab[106910]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:36
	_go_fuzz_dep_.CoverTab[106911]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:38
	// _ = "end of CoverTab[106911]"
}

func (t *TxnOffsetCommitRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:41
	_go_fuzz_dep_.CoverTab[106928]++
													if t.TransactionalID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:42
		_go_fuzz_dep_.CoverTab[106935]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:43
		// _ = "end of CoverTab[106935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:44
		_go_fuzz_dep_.CoverTab[106936]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:44
		// _ = "end of CoverTab[106936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:44
	// _ = "end of CoverTab[106928]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:44
	_go_fuzz_dep_.CoverTab[106929]++
													if t.GroupID, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:45
		_go_fuzz_dep_.CoverTab[106937]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:46
		// _ = "end of CoverTab[106937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:47
		_go_fuzz_dep_.CoverTab[106938]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:47
		// _ = "end of CoverTab[106938]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:47
	// _ = "end of CoverTab[106929]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:47
	_go_fuzz_dep_.CoverTab[106930]++
													if t.ProducerID, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:48
		_go_fuzz_dep_.CoverTab[106939]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:49
		// _ = "end of CoverTab[106939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:50
		_go_fuzz_dep_.CoverTab[106940]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:50
		// _ = "end of CoverTab[106940]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:50
	// _ = "end of CoverTab[106930]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:50
	_go_fuzz_dep_.CoverTab[106931]++
													if t.ProducerEpoch, err = pd.getInt16(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:51
		_go_fuzz_dep_.CoverTab[106941]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:52
		// _ = "end of CoverTab[106941]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:53
		_go_fuzz_dep_.CoverTab[106942]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:53
		// _ = "end of CoverTab[106942]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:53
	// _ = "end of CoverTab[106931]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:53
	_go_fuzz_dep_.CoverTab[106932]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:56
		_go_fuzz_dep_.CoverTab[106943]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:57
		// _ = "end of CoverTab[106943]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:58
		_go_fuzz_dep_.CoverTab[106944]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:58
		// _ = "end of CoverTab[106944]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:58
	// _ = "end of CoverTab[106932]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:58
	_go_fuzz_dep_.CoverTab[106933]++

													t.Topics = make(map[string][]*PartitionOffsetMetadata)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:61
		_go_fuzz_dep_.CoverTab[106945]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:63
			_go_fuzz_dep_.CoverTab[106948]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:64
			// _ = "end of CoverTab[106948]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:65
			_go_fuzz_dep_.CoverTab[106949]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:65
			// _ = "end of CoverTab[106949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:65
		// _ = "end of CoverTab[106945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:65
		_go_fuzz_dep_.CoverTab[106946]++

														m, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:68
			_go_fuzz_dep_.CoverTab[106950]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:69
			// _ = "end of CoverTab[106950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:70
			_go_fuzz_dep_.CoverTab[106951]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:70
			// _ = "end of CoverTab[106951]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:70
		// _ = "end of CoverTab[106946]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:70
		_go_fuzz_dep_.CoverTab[106947]++

														t.Topics[topic] = make([]*PartitionOffsetMetadata, m)

														for j := 0; j < m; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:74
			_go_fuzz_dep_.CoverTab[106952]++
															partitionOffsetMetadata := new(PartitionOffsetMetadata)
															if err := partitionOffsetMetadata.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:76
				_go_fuzz_dep_.CoverTab[106954]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:77
				// _ = "end of CoverTab[106954]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:78
				_go_fuzz_dep_.CoverTab[106955]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:78
				// _ = "end of CoverTab[106955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:78
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:78
			// _ = "end of CoverTab[106952]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:78
			_go_fuzz_dep_.CoverTab[106953]++
															t.Topics[topic][j] = partitionOffsetMetadata
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:79
			// _ = "end of CoverTab[106953]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:80
		// _ = "end of CoverTab[106947]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:81
	// _ = "end of CoverTab[106933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:81
	_go_fuzz_dep_.CoverTab[106934]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:83
	// _ = "end of CoverTab[106934]"
}

func (a *TxnOffsetCommitRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:86
	_go_fuzz_dep_.CoverTab[106956]++
													return 28
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:87
	// _ = "end of CoverTab[106956]"
}

func (a *TxnOffsetCommitRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:90
	_go_fuzz_dep_.CoverTab[106957]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:91
	// _ = "end of CoverTab[106957]"
}

func (a *TxnOffsetCommitRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:94
	_go_fuzz_dep_.CoverTab[106958]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:95
	// _ = "end of CoverTab[106958]"
}

func (a *TxnOffsetCommitRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:98
	_go_fuzz_dep_.CoverTab[106959]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:99
	// _ = "end of CoverTab[106959]"
}

type PartitionOffsetMetadata struct {
	Partition	int32
	Offset		int64
	Metadata	*string
}

func (p *PartitionOffsetMetadata) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:108
	_go_fuzz_dep_.CoverTab[106960]++
													pe.putInt32(p.Partition)
													pe.putInt64(p.Offset)
													if err := pe.putNullableString(p.Metadata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:111
		_go_fuzz_dep_.CoverTab[106962]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:112
		// _ = "end of CoverTab[106962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:113
		_go_fuzz_dep_.CoverTab[106963]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:113
		// _ = "end of CoverTab[106963]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:113
	// _ = "end of CoverTab[106960]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:113
	_go_fuzz_dep_.CoverTab[106961]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:115
	// _ = "end of CoverTab[106961]"
}

func (p *PartitionOffsetMetadata) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:118
	_go_fuzz_dep_.CoverTab[106964]++
													if p.Partition, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:119
		_go_fuzz_dep_.CoverTab[106968]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:120
		// _ = "end of CoverTab[106968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:121
		_go_fuzz_dep_.CoverTab[106969]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:121
		// _ = "end of CoverTab[106969]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:121
	// _ = "end of CoverTab[106964]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:121
	_go_fuzz_dep_.CoverTab[106965]++
													if p.Offset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:122
		_go_fuzz_dep_.CoverTab[106970]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:123
		// _ = "end of CoverTab[106970]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:124
		_go_fuzz_dep_.CoverTab[106971]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:124
		// _ = "end of CoverTab[106971]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:124
	// _ = "end of CoverTab[106965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:124
	_go_fuzz_dep_.CoverTab[106966]++
													if p.Metadata, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:125
		_go_fuzz_dep_.CoverTab[106972]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:126
		// _ = "end of CoverTab[106972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:127
		_go_fuzz_dep_.CoverTab[106973]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:127
		// _ = "end of CoverTab[106973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:127
	// _ = "end of CoverTab[106966]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:127
	_go_fuzz_dep_.CoverTab[106967]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:129
	// _ = "end of CoverTab[106967]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/txn_offset_commit_request.go:130
var _ = _go_fuzz_dep_.CoverTab
