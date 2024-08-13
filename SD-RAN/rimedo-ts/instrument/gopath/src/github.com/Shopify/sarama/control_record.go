//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:1
)

// ControlRecordType ...
type ControlRecordType int

const (
	// ControlRecordAbort is a control record for abort
	ControlRecordAbort	ControlRecordType	= iota
	// ControlRecordCommit is a control record for commit
	ControlRecordCommit
	// ControlRecordUnknown is a control record of unknown type
	ControlRecordUnknown
)

// Control records are returned as a record by fetchRequest
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:15
// However unlike "normal" records, they mean nothing application wise.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:15
// They only serve internal logic for supporting transactions.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:18
type ControlRecord struct {
	Version			int16
	CoordinatorEpoch	int32
	Type			ControlRecordType
}

func (cr *ControlRecord) decode(key, value packetDecoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:24
	_go_fuzz_dep_.CoverTab[101385]++
												var err error

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:28
	cr.Version, err = key.getInt16()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:29
		_go_fuzz_dep_.CoverTab[101390]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:30
		// _ = "end of CoverTab[101390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:31
		_go_fuzz_dep_.CoverTab[101391]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:31
		// _ = "end of CoverTab[101391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:31
	// _ = "end of CoverTab[101385]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:31
	_go_fuzz_dep_.CoverTab[101386]++

												recordType, err := key.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:34
		_go_fuzz_dep_.CoverTab[101392]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:35
		// _ = "end of CoverTab[101392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:36
		_go_fuzz_dep_.CoverTab[101393]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:36
		// _ = "end of CoverTab[101393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:36
	// _ = "end of CoverTab[101386]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:36
	_go_fuzz_dep_.CoverTab[101387]++

												switch recordType {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:39
		_go_fuzz_dep_.CoverTab[101394]++
													cr.Type = ControlRecordAbort
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:40
		// _ = "end of CoverTab[101394]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:41
		_go_fuzz_dep_.CoverTab[101395]++
													cr.Type = ControlRecordCommit
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:42
		// _ = "end of CoverTab[101395]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:43
		_go_fuzz_dep_.CoverTab[101396]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:46
		cr.Type = ControlRecordUnknown
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:46
		// _ = "end of CoverTab[101396]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:47
	// _ = "end of CoverTab[101387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:47
	_go_fuzz_dep_.CoverTab[101388]++

												if cr.Type != ControlRecordUnknown {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:49
		_go_fuzz_dep_.CoverTab[101397]++
													cr.Version, err = value.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:51
			_go_fuzz_dep_.CoverTab[101399]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:52
			// _ = "end of CoverTab[101399]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:53
			_go_fuzz_dep_.CoverTab[101400]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:53
			// _ = "end of CoverTab[101400]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:53
		// _ = "end of CoverTab[101397]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:53
		_go_fuzz_dep_.CoverTab[101398]++

													cr.CoordinatorEpoch, err = value.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:56
			_go_fuzz_dep_.CoverTab[101401]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:57
			// _ = "end of CoverTab[101401]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:58
			_go_fuzz_dep_.CoverTab[101402]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:58
			// _ = "end of CoverTab[101402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:58
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:58
		// _ = "end of CoverTab[101398]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:59
		_go_fuzz_dep_.CoverTab[101403]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:59
		// _ = "end of CoverTab[101403]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:59
	// _ = "end of CoverTab[101388]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:59
	_go_fuzz_dep_.CoverTab[101389]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:60
	// _ = "end of CoverTab[101389]"
}

func (cr *ControlRecord) encode(key, value packetEncoder) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:63
	_go_fuzz_dep_.CoverTab[101404]++
												value.putInt16(cr.Version)
												value.putInt32(cr.CoordinatorEpoch)
												key.putInt16(cr.Version)

												switch cr.Type {
	case ControlRecordAbort:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:69
		_go_fuzz_dep_.CoverTab[101405]++
													key.putInt16(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:70
		// _ = "end of CoverTab[101405]"
	case ControlRecordCommit:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:71
		_go_fuzz_dep_.CoverTab[101406]++
													key.putInt16(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:72
		// _ = "end of CoverTab[101406]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:72
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:72
		_go_fuzz_dep_.CoverTab[101407]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:72
		// _ = "end of CoverTab[101407]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:73
	// _ = "end of CoverTab[101404]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/control_record.go:74
var _ = _go_fuzz_dep_.CoverTab
