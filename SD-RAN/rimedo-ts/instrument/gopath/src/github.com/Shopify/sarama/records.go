//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:1
)

import "fmt"

const (
	unknownRecords	= iota
	legacyRecords
	defaultRecords

	magicOffset	= 16
)

// Records implements a union type containing either a RecordBatch or a legacy MessageSet.
type Records struct {
	recordsType	int
	MsgSet		*MessageSet
	RecordBatch	*RecordBatch
}

func newLegacyRecords(msgSet *MessageSet) Records {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:20
	_go_fuzz_dep_.CoverTab[106413]++
											return Records{recordsType: legacyRecords, MsgSet: msgSet}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:21
	// _ = "end of CoverTab[106413]"
}

func newDefaultRecords(batch *RecordBatch) Records {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:24
	_go_fuzz_dep_.CoverTab[106414]++
											return Records{recordsType: defaultRecords, RecordBatch: batch}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:25
	// _ = "end of CoverTab[106414]"
}

// setTypeFromFields sets type of Records depending on which of MsgSet or RecordBatch is not nil.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:28
// The first return value indicates whether both fields are nil (and the type is not set).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:28
// If both fields are not nil, it returns an error.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:31
func (r *Records) setTypeFromFields() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:31
	_go_fuzz_dep_.CoverTab[106415]++
											if r.MsgSet == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:32
		_go_fuzz_dep_.CoverTab[106419]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:32
		return r.RecordBatch == nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:32
		// _ = "end of CoverTab[106419]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:32
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:32
		_go_fuzz_dep_.CoverTab[106420]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:33
		// _ = "end of CoverTab[106420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:34
		_go_fuzz_dep_.CoverTab[106421]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:34
		// _ = "end of CoverTab[106421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:34
	// _ = "end of CoverTab[106415]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:34
	_go_fuzz_dep_.CoverTab[106416]++
											if r.MsgSet != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:35
		_go_fuzz_dep_.CoverTab[106422]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:35
		return r.RecordBatch != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:35
		// _ = "end of CoverTab[106422]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:35
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:35
		_go_fuzz_dep_.CoverTab[106423]++
												return false, fmt.Errorf("both MsgSet and RecordBatch are set, but record type is unknown")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:36
		// _ = "end of CoverTab[106423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:37
		_go_fuzz_dep_.CoverTab[106424]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:37
		// _ = "end of CoverTab[106424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:37
	// _ = "end of CoverTab[106416]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:37
	_go_fuzz_dep_.CoverTab[106417]++
											r.recordsType = defaultRecords
											if r.MsgSet != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:39
		_go_fuzz_dep_.CoverTab[106425]++
												r.recordsType = legacyRecords
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:40
		// _ = "end of CoverTab[106425]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:41
		_go_fuzz_dep_.CoverTab[106426]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:41
		// _ = "end of CoverTab[106426]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:41
	// _ = "end of CoverTab[106417]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:41
	_go_fuzz_dep_.CoverTab[106418]++
											return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:42
	// _ = "end of CoverTab[106418]"
}

func (r *Records) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:45
	_go_fuzz_dep_.CoverTab[106427]++
											if r.recordsType == unknownRecords {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:46
		_go_fuzz_dep_.CoverTab[106430]++
												if empty, err := r.setTypeFromFields(); err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:47
			_go_fuzz_dep_.CoverTab[106431]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:47
			return empty
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:47
			// _ = "end of CoverTab[106431]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:47
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:47
			_go_fuzz_dep_.CoverTab[106432]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:48
			// _ = "end of CoverTab[106432]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:49
			_go_fuzz_dep_.CoverTab[106433]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:49
			// _ = "end of CoverTab[106433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:49
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:49
		// _ = "end of CoverTab[106430]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:50
		_go_fuzz_dep_.CoverTab[106434]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:50
		// _ = "end of CoverTab[106434]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:50
	// _ = "end of CoverTab[106427]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:50
	_go_fuzz_dep_.CoverTab[106428]++

											switch r.recordsType {
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:53
		_go_fuzz_dep_.CoverTab[106435]++
												if r.MsgSet == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:54
			_go_fuzz_dep_.CoverTab[106440]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:55
			// _ = "end of CoverTab[106440]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:56
			_go_fuzz_dep_.CoverTab[106441]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:56
			// _ = "end of CoverTab[106441]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:56
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:56
		// _ = "end of CoverTab[106435]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:56
		_go_fuzz_dep_.CoverTab[106436]++
												return r.MsgSet.encode(pe)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:57
		// _ = "end of CoverTab[106436]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:58
		_go_fuzz_dep_.CoverTab[106437]++
												if r.RecordBatch == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:59
			_go_fuzz_dep_.CoverTab[106442]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:60
			// _ = "end of CoverTab[106442]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:61
			_go_fuzz_dep_.CoverTab[106443]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:61
			// _ = "end of CoverTab[106443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:61
		// _ = "end of CoverTab[106437]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:61
		_go_fuzz_dep_.CoverTab[106438]++
												return r.RecordBatch.encode(pe)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:62
		// _ = "end of CoverTab[106438]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:62
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:62
		_go_fuzz_dep_.CoverTab[106439]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:62
		// _ = "end of CoverTab[106439]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:63
	// _ = "end of CoverTab[106428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:63
	_go_fuzz_dep_.CoverTab[106429]++

											return fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:65
	// _ = "end of CoverTab[106429]"
}

func (r *Records) setTypeFromMagic(pd packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:68
	_go_fuzz_dep_.CoverTab[106444]++
											magic, err := magicValue(pd)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:70
		_go_fuzz_dep_.CoverTab[106447]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:71
		// _ = "end of CoverTab[106447]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:72
		_go_fuzz_dep_.CoverTab[106448]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:72
		// _ = "end of CoverTab[106448]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:72
	// _ = "end of CoverTab[106444]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:72
	_go_fuzz_dep_.CoverTab[106445]++

											r.recordsType = defaultRecords
											if magic < 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:75
		_go_fuzz_dep_.CoverTab[106449]++
												r.recordsType = legacyRecords
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:76
		// _ = "end of CoverTab[106449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:77
		_go_fuzz_dep_.CoverTab[106450]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:77
		// _ = "end of CoverTab[106450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:77
	// _ = "end of CoverTab[106445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:77
	_go_fuzz_dep_.CoverTab[106446]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:79
	// _ = "end of CoverTab[106446]"
}

func (r *Records) decode(pd packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:82
	_go_fuzz_dep_.CoverTab[106451]++
											if r.recordsType == unknownRecords {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:83
		_go_fuzz_dep_.CoverTab[106454]++
												if err := r.setTypeFromMagic(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:84
			_go_fuzz_dep_.CoverTab[106455]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:85
			// _ = "end of CoverTab[106455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:86
			_go_fuzz_dep_.CoverTab[106456]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:86
			// _ = "end of CoverTab[106456]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:86
		// _ = "end of CoverTab[106454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:87
		_go_fuzz_dep_.CoverTab[106457]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:87
		// _ = "end of CoverTab[106457]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:87
	// _ = "end of CoverTab[106451]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:87
	_go_fuzz_dep_.CoverTab[106452]++

											switch r.recordsType {
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:90
		_go_fuzz_dep_.CoverTab[106458]++
												r.MsgSet = &MessageSet{}
												return r.MsgSet.decode(pd)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:92
		// _ = "end of CoverTab[106458]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:93
		_go_fuzz_dep_.CoverTab[106459]++
												r.RecordBatch = &RecordBatch{}
												return r.RecordBatch.decode(pd)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:95
		// _ = "end of CoverTab[106459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:95
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:95
		_go_fuzz_dep_.CoverTab[106460]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:95
		// _ = "end of CoverTab[106460]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:96
	// _ = "end of CoverTab[106452]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:96
	_go_fuzz_dep_.CoverTab[106453]++
											return fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:97
	// _ = "end of CoverTab[106453]"
}

func (r *Records) numRecords() (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:100
	_go_fuzz_dep_.CoverTab[106461]++
											if r.recordsType == unknownRecords {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:101
		_go_fuzz_dep_.CoverTab[106464]++
												if empty, err := r.setTypeFromFields(); err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:102
			_go_fuzz_dep_.CoverTab[106465]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:102
			return empty
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:102
			// _ = "end of CoverTab[106465]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:102
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:102
			_go_fuzz_dep_.CoverTab[106466]++
													return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:103
			// _ = "end of CoverTab[106466]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:104
			_go_fuzz_dep_.CoverTab[106467]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:104
			// _ = "end of CoverTab[106467]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:104
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:104
		// _ = "end of CoverTab[106464]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:105
		_go_fuzz_dep_.CoverTab[106468]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:105
		// _ = "end of CoverTab[106468]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:105
	// _ = "end of CoverTab[106461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:105
	_go_fuzz_dep_.CoverTab[106462]++

											switch r.recordsType {
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:108
		_go_fuzz_dep_.CoverTab[106469]++
												if r.MsgSet == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:109
			_go_fuzz_dep_.CoverTab[106474]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:110
			// _ = "end of CoverTab[106474]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:111
			_go_fuzz_dep_.CoverTab[106475]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:111
			// _ = "end of CoverTab[106475]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:111
		// _ = "end of CoverTab[106469]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:111
		_go_fuzz_dep_.CoverTab[106470]++
												return len(r.MsgSet.Messages), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:112
		// _ = "end of CoverTab[106470]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:113
		_go_fuzz_dep_.CoverTab[106471]++
												if r.RecordBatch == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:114
			_go_fuzz_dep_.CoverTab[106476]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:115
			// _ = "end of CoverTab[106476]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:116
			_go_fuzz_dep_.CoverTab[106477]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:116
			// _ = "end of CoverTab[106477]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:116
		// _ = "end of CoverTab[106471]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:116
		_go_fuzz_dep_.CoverTab[106472]++
												return len(r.RecordBatch.Records), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:117
		// _ = "end of CoverTab[106472]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:117
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:117
		_go_fuzz_dep_.CoverTab[106473]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:117
		// _ = "end of CoverTab[106473]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:118
	// _ = "end of CoverTab[106462]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:118
	_go_fuzz_dep_.CoverTab[106463]++
											return 0, fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:119
	// _ = "end of CoverTab[106463]"
}

func (r *Records) isPartial() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:122
	_go_fuzz_dep_.CoverTab[106478]++
											if r.recordsType == unknownRecords {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:123
		_go_fuzz_dep_.CoverTab[106481]++
												if empty, err := r.setTypeFromFields(); err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:124
			_go_fuzz_dep_.CoverTab[106482]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:124
			return empty
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:124
			// _ = "end of CoverTab[106482]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:124
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:124
			_go_fuzz_dep_.CoverTab[106483]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:125
			// _ = "end of CoverTab[106483]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:126
			_go_fuzz_dep_.CoverTab[106484]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:126
			// _ = "end of CoverTab[106484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:126
		// _ = "end of CoverTab[106481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:127
		_go_fuzz_dep_.CoverTab[106485]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:127
		// _ = "end of CoverTab[106485]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:127
	// _ = "end of CoverTab[106478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:127
	_go_fuzz_dep_.CoverTab[106479]++

											switch r.recordsType {
	case unknownRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:130
		_go_fuzz_dep_.CoverTab[106486]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:131
		// _ = "end of CoverTab[106486]"
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:132
		_go_fuzz_dep_.CoverTab[106487]++
												if r.MsgSet == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:133
			_go_fuzz_dep_.CoverTab[106492]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:134
			// _ = "end of CoverTab[106492]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:135
			_go_fuzz_dep_.CoverTab[106493]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:135
			// _ = "end of CoverTab[106493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:135
		// _ = "end of CoverTab[106487]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:135
		_go_fuzz_dep_.CoverTab[106488]++
												return r.MsgSet.PartialTrailingMessage, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:136
		// _ = "end of CoverTab[106488]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:137
		_go_fuzz_dep_.CoverTab[106489]++
												if r.RecordBatch == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:138
			_go_fuzz_dep_.CoverTab[106494]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:139
			// _ = "end of CoverTab[106494]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:140
			_go_fuzz_dep_.CoverTab[106495]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:140
			// _ = "end of CoverTab[106495]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:140
		// _ = "end of CoverTab[106489]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:140
		_go_fuzz_dep_.CoverTab[106490]++
												return r.RecordBatch.PartialTrailingRecord, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:141
		// _ = "end of CoverTab[106490]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:141
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:141
		_go_fuzz_dep_.CoverTab[106491]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:141
		// _ = "end of CoverTab[106491]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:142
	// _ = "end of CoverTab[106479]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:142
	_go_fuzz_dep_.CoverTab[106480]++
											return false, fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:143
	// _ = "end of CoverTab[106480]"
}

func (r *Records) isControl() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:146
	_go_fuzz_dep_.CoverTab[106496]++
											if r.recordsType == unknownRecords {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:147
		_go_fuzz_dep_.CoverTab[106499]++
												if empty, err := r.setTypeFromFields(); err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:148
			_go_fuzz_dep_.CoverTab[106500]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:148
			return empty
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:148
			// _ = "end of CoverTab[106500]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:148
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:148
			_go_fuzz_dep_.CoverTab[106501]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:149
			// _ = "end of CoverTab[106501]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:150
			_go_fuzz_dep_.CoverTab[106502]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:150
			// _ = "end of CoverTab[106502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:150
		// _ = "end of CoverTab[106499]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:151
		_go_fuzz_dep_.CoverTab[106503]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:151
		// _ = "end of CoverTab[106503]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:151
	// _ = "end of CoverTab[106496]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:151
	_go_fuzz_dep_.CoverTab[106497]++

											switch r.recordsType {
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:154
		_go_fuzz_dep_.CoverTab[106504]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:155
		// _ = "end of CoverTab[106504]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:156
		_go_fuzz_dep_.CoverTab[106505]++
												if r.RecordBatch == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:157
			_go_fuzz_dep_.CoverTab[106508]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:158
			// _ = "end of CoverTab[106508]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:159
			_go_fuzz_dep_.CoverTab[106509]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:159
			// _ = "end of CoverTab[106509]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:159
		// _ = "end of CoverTab[106505]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:159
		_go_fuzz_dep_.CoverTab[106506]++
												return r.RecordBatch.Control, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:160
		// _ = "end of CoverTab[106506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:160
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:160
		_go_fuzz_dep_.CoverTab[106507]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:160
		// _ = "end of CoverTab[106507]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:161
	// _ = "end of CoverTab[106497]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:161
	_go_fuzz_dep_.CoverTab[106498]++
											return false, fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:162
	// _ = "end of CoverTab[106498]"
}

func (r *Records) isOverflow() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:165
	_go_fuzz_dep_.CoverTab[106510]++
											if r.recordsType == unknownRecords {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:166
		_go_fuzz_dep_.CoverTab[106513]++
												if empty, err := r.setTypeFromFields(); err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:167
			_go_fuzz_dep_.CoverTab[106514]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:167
			return empty
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:167
			// _ = "end of CoverTab[106514]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:167
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:167
			_go_fuzz_dep_.CoverTab[106515]++
													return false, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:168
			// _ = "end of CoverTab[106515]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:169
			_go_fuzz_dep_.CoverTab[106516]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:169
			// _ = "end of CoverTab[106516]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:169
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:169
		// _ = "end of CoverTab[106513]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:170
		_go_fuzz_dep_.CoverTab[106517]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:170
		// _ = "end of CoverTab[106517]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:170
	// _ = "end of CoverTab[106510]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:170
	_go_fuzz_dep_.CoverTab[106511]++

											switch r.recordsType {
	case unknownRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:173
		_go_fuzz_dep_.CoverTab[106518]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:174
		// _ = "end of CoverTab[106518]"
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:175
		_go_fuzz_dep_.CoverTab[106519]++
												if r.MsgSet == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:176
			_go_fuzz_dep_.CoverTab[106523]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:177
			// _ = "end of CoverTab[106523]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:178
			_go_fuzz_dep_.CoverTab[106524]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:178
			// _ = "end of CoverTab[106524]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:178
		// _ = "end of CoverTab[106519]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:178
		_go_fuzz_dep_.CoverTab[106520]++
												return r.MsgSet.OverflowMessage, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:179
		// _ = "end of CoverTab[106520]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:180
		_go_fuzz_dep_.CoverTab[106521]++
												return false, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:181
		// _ = "end of CoverTab[106521]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:181
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:181
		_go_fuzz_dep_.CoverTab[106522]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:181
		// _ = "end of CoverTab[106522]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:182
	// _ = "end of CoverTab[106511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:182
	_go_fuzz_dep_.CoverTab[106512]++
											return false, fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:183
	// _ = "end of CoverTab[106512]"
}

func (r *Records) recordsOffset() (*int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:186
	_go_fuzz_dep_.CoverTab[106525]++
											switch r.recordsType {
	case unknownRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:188
		_go_fuzz_dep_.CoverTab[106527]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:189
		// _ = "end of CoverTab[106527]"
	case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:190
		_go_fuzz_dep_.CoverTab[106528]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:191
		// _ = "end of CoverTab[106528]"
	case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:192
		_go_fuzz_dep_.CoverTab[106529]++
												if r.RecordBatch == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:193
			_go_fuzz_dep_.CoverTab[106532]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:194
			// _ = "end of CoverTab[106532]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:195
			_go_fuzz_dep_.CoverTab[106533]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:195
			// _ = "end of CoverTab[106533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:195
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:195
		// _ = "end of CoverTab[106529]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:195
		_go_fuzz_dep_.CoverTab[106530]++
												return &r.RecordBatch.FirstOffset, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:196
		// _ = "end of CoverTab[106530]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:196
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:196
		_go_fuzz_dep_.CoverTab[106531]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:196
		// _ = "end of CoverTab[106531]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:197
	// _ = "end of CoverTab[106525]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:197
	_go_fuzz_dep_.CoverTab[106526]++
											return nil, fmt.Errorf("unknown records type: %v", r.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:198
	// _ = "end of CoverTab[106526]"
}

func magicValue(pd packetDecoder) (int8, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:201
	_go_fuzz_dep_.CoverTab[106534]++
											return pd.peekInt8(magicOffset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:202
	// _ = "end of CoverTab[106534]"
}

func (r *Records) getControlRecord() (ControlRecord, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:205
	_go_fuzz_dep_.CoverTab[106535]++
											if r.RecordBatch == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:206
		_go_fuzz_dep_.CoverTab[106538]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:206
		return len(r.RecordBatch.Records) <= 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:206
		// _ = "end of CoverTab[106538]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:206
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:206
		_go_fuzz_dep_.CoverTab[106539]++
												return ControlRecord{}, fmt.Errorf("cannot get control record, record batch is empty")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:207
		// _ = "end of CoverTab[106539]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:208
		_go_fuzz_dep_.CoverTab[106540]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:208
		// _ = "end of CoverTab[106540]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:208
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:208
	// _ = "end of CoverTab[106535]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:208
	_go_fuzz_dep_.CoverTab[106536]++

											firstRecord := r.RecordBatch.Records[0]
											controlRecord := ControlRecord{}
											err := controlRecord.decode(&realDecoder{raw: firstRecord.Key}, &realDecoder{raw: firstRecord.Value})
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:213
		_go_fuzz_dep_.CoverTab[106541]++
												return ControlRecord{}, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:214
		// _ = "end of CoverTab[106541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:215
		_go_fuzz_dep_.CoverTab[106542]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:215
		// _ = "end of CoverTab[106542]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:215
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:215
	// _ = "end of CoverTab[106536]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:215
	_go_fuzz_dep_.CoverTab[106537]++

											return controlRecord, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:217
	// _ = "end of CoverTab[106537]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:218
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/records.go:218
var _ = _go_fuzz_dep_.CoverTab
