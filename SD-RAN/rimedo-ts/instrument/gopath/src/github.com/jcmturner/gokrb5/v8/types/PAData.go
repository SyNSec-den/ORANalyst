//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
package types

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:1
)

// Reference: https://www.ietf.org/rfc/rfc4120.txt
// Section: 5.2.7
import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/iana/patype"
)

// PAData implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7
type PAData struct {
	PADataType	int32	`asn1:"explicit,tag:1"`
	PADataValue	[]byte	`asn1:"explicit,tag:2"`
}

// PADataSequence implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7
type PADataSequence []PAData

// MethodData implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.9.1
type MethodData []PAData

// PAEncTimestamp implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7.2
type PAEncTimestamp EncryptedData

// PAEncTSEnc implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7.2
type PAEncTSEnc struct {
	PATimestamp	time.Time	`asn1:"generalized,explicit,tag:0"`
	PAUSec		int		`asn1:"explicit,optional,tag:1"`
}

// Contains tests if a PADataSequence contains PA Data of a certain type.
func (pas *PADataSequence) Contains(patype int32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:35
	_go_fuzz_dep_.CoverTab[86059]++
												for _, pa := range *pas {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:36
		_go_fuzz_dep_.CoverTab[86061]++
													if pa.PADataType == patype {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:37
			_go_fuzz_dep_.CoverTab[86062]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:38
			// _ = "end of CoverTab[86062]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:39
			_go_fuzz_dep_.CoverTab[86063]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:39
			// _ = "end of CoverTab[86063]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:39
		// _ = "end of CoverTab[86061]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:40
	// _ = "end of CoverTab[86059]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:40
	_go_fuzz_dep_.CoverTab[86060]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:41
	// _ = "end of CoverTab[86060]"
}

// GetPAEncTSEncAsnMarshalled returns the bytes of a PAEncTSEnc.
func GetPAEncTSEncAsnMarshalled() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:45
	_go_fuzz_dep_.CoverTab[86064]++
												t := time.Now().UTC()
												p := PAEncTSEnc{
		PATimestamp:	t,
		PAUSec:		int((t.UnixNano() / int64(time.Microsecond)) - (t.Unix() * 1e6)),
	}
	b, err := asn1.Marshal(p)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:52
		_go_fuzz_dep_.CoverTab[86066]++
													return b, fmt.Errorf("error mashaling PAEncTSEnc: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:53
		// _ = "end of CoverTab[86066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:54
		_go_fuzz_dep_.CoverTab[86067]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:54
		// _ = "end of CoverTab[86067]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:54
	// _ = "end of CoverTab[86064]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:54
	_go_fuzz_dep_.CoverTab[86065]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:55
	// _ = "end of CoverTab[86065]"
}

// ETypeInfoEntry implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7.4
type ETypeInfoEntry struct {
	EType	int32	`asn1:"explicit,tag:0"`
	Salt	[]byte	`asn1:"explicit,optional,tag:1"`
}

// ETypeInfo implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7.4
type ETypeInfo []ETypeInfoEntry

// ETypeInfo2Entry implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7.5
type ETypeInfo2Entry struct {
	EType		int32	`asn1:"explicit,tag:0"`
	Salt		string	`asn1:"explicit,optional,generalstring,tag:1"`
	S2KParams	[]byte	`asn1:"explicit,optional,tag:2"`
}

// ETypeInfo2 implements RFC 4120 types: https://tools.ietf.org/html/rfc4120#section-5.2.7.5
type ETypeInfo2 []ETypeInfo2Entry

// PAReqEncPARep PA Data Type
type PAReqEncPARep struct {
	ChksumType	int32	`asn1:"explicit,tag:0"`
	Chksum		[]byte	`asn1:"explicit,tag:1"`
}

// Unmarshal bytes into the PAData
func (pa *PAData) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:84
	_go_fuzz_dep_.CoverTab[86068]++
												_, err := asn1.Unmarshal(b, pa)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:86
	// _ = "end of CoverTab[86068]"
}

// Unmarshal bytes into the PADataSequence
func (pas *PADataSequence) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:90
	_go_fuzz_dep_.CoverTab[86069]++
												_, err := asn1.Unmarshal(b, pas)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:92
	// _ = "end of CoverTab[86069]"
}

// Unmarshal bytes into the PAReqEncPARep
func (pa *PAReqEncPARep) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:96
	_go_fuzz_dep_.CoverTab[86070]++
												_, err := asn1.Unmarshal(b, pa)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:98
	// _ = "end of CoverTab[86070]"
}

// Unmarshal bytes into the PAEncTimestamp
func (pa *PAEncTimestamp) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:102
	_go_fuzz_dep_.CoverTab[86071]++
												_, err := asn1.Unmarshal(b, pa)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:104
	// _ = "end of CoverTab[86071]"
}

// Unmarshal bytes into the PAEncTSEnc
func (pa *PAEncTSEnc) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:108
	_go_fuzz_dep_.CoverTab[86072]++
												_, err := asn1.Unmarshal(b, pa)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:110
	// _ = "end of CoverTab[86072]"
}

// Unmarshal bytes into the ETypeInfo
func (a *ETypeInfo) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:114
	_go_fuzz_dep_.CoverTab[86073]++
												_, err := asn1.Unmarshal(b, a)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:116
	// _ = "end of CoverTab[86073]"
}

// Unmarshal bytes into the ETypeInfoEntry
func (a *ETypeInfoEntry) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:120
	_go_fuzz_dep_.CoverTab[86074]++
												_, err := asn1.Unmarshal(b, a)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:122
	// _ = "end of CoverTab[86074]"
}

// Unmarshal bytes into the ETypeInfo2
func (a *ETypeInfo2) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:126
	_go_fuzz_dep_.CoverTab[86075]++
												_, err := asn1.Unmarshal(b, a)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:128
	// _ = "end of CoverTab[86075]"
}

// Unmarshal bytes into the ETypeInfo2Entry
func (a *ETypeInfo2Entry) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:132
	_go_fuzz_dep_.CoverTab[86076]++
												_, err := asn1.Unmarshal(b, a)
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:134
	// _ = "end of CoverTab[86076]"
}

// GetETypeInfo returns an ETypeInfo from the PAData.
func (pa *PAData) GetETypeInfo() (d ETypeInfo, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:138
	_go_fuzz_dep_.CoverTab[86077]++
												if pa.PADataType != patype.PA_ETYPE_INFO {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:139
		_go_fuzz_dep_.CoverTab[86079]++
													err = fmt.Errorf("PAData does not contain PA EType Info data. TypeID Expected: %v; Actual: %v", patype.PA_ETYPE_INFO, pa.PADataType)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:141
		// _ = "end of CoverTab[86079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:142
		_go_fuzz_dep_.CoverTab[86080]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:142
		// _ = "end of CoverTab[86080]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:142
	// _ = "end of CoverTab[86077]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:142
	_go_fuzz_dep_.CoverTab[86078]++
												_, err = asn1.Unmarshal(pa.PADataValue, &d)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:144
	// _ = "end of CoverTab[86078]"
}

// GetETypeInfo2 returns an ETypeInfo2 from the PAData.
func (pa *PAData) GetETypeInfo2() (d ETypeInfo2, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:148
	_go_fuzz_dep_.CoverTab[86081]++
												if pa.PADataType != patype.PA_ETYPE_INFO2 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:149
		_go_fuzz_dep_.CoverTab[86083]++
													err = fmt.Errorf("PAData does not contain PA EType Info 2 data. TypeID Expected: %v; Actual: %v", patype.PA_ETYPE_INFO2, pa.PADataType)
													return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:151
		// _ = "end of CoverTab[86083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:152
		_go_fuzz_dep_.CoverTab[86084]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:152
		// _ = "end of CoverTab[86084]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:152
	// _ = "end of CoverTab[86081]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:152
	_go_fuzz_dep_.CoverTab[86082]++
												_, err = asn1.Unmarshal(pa.PADataValue, &d)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:154
	// _ = "end of CoverTab[86082]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:155
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/types/PAData.go:155
var _ = _go_fuzz_dep_.CoverTab
