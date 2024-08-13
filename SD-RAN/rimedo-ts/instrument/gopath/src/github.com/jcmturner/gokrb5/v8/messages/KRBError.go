//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:1
// Package messages implements Kerberos 5 message types and methods.
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:2
)

import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/iana"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

// KRBError implements RFC 4120 KRB_ERROR: https://tools.ietf.org/html/rfc4120#section-5.9.1.
type KRBError struct {
	PVNO		int			`asn1:"explicit,tag:0"`
	MsgType		int			`asn1:"explicit,tag:1"`
	CTime		time.Time		`asn1:"generalized,optional,explicit,tag:2"`
	Cusec		int			`asn1:"optional,explicit,tag:3"`
	STime		time.Time		`asn1:"generalized,explicit,tag:4"`
	Susec		int			`asn1:"explicit,tag:5"`
	ErrorCode	int32			`asn1:"explicit,tag:6"`
	CRealm		string			`asn1:"generalstring,optional,explicit,tag:7"`
	CName		types.PrincipalName	`asn1:"optional,explicit,tag:8"`
	Realm		string			`asn1:"generalstring,explicit,tag:9"`
	SName		types.PrincipalName	`asn1:"explicit,tag:10"`
	EText		string			`asn1:"generalstring,optional,explicit,tag:11"`
	EData		[]byte			`asn1:"optional,explicit,tag:12"`
}

// NewKRBError creates a new KRBError.
func NewKRBError(sname types.PrincipalName, realm string, code int32, etext string) KRBError {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:36
	_go_fuzz_dep_.CoverTab[88039]++
													t := time.Now().UTC()
													return KRBError{
		PVNO:		iana.PVNO,
		MsgType:	msgtype.KRB_ERROR,
		STime:		t,
		Susec:		int((t.UnixNano() / int64(time.Microsecond)) - (t.Unix() * 1e6)),
		ErrorCode:	code,
		SName:		sname,
		Realm:		realm,
		EText:		etext,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:47
	// _ = "end of CoverTab[88039]"
}

// Unmarshal bytes b into the KRBError struct.
func (k *KRBError) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:51
	_go_fuzz_dep_.CoverTab[88040]++
													_, err := asn1.UnmarshalWithParams(b, k, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.KRBError))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:53
		_go_fuzz_dep_.CoverTab[88043]++
														return krberror.Errorf(err, krberror.EncodingError, "KRB_ERROR unmarshal error")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:54
		// _ = "end of CoverTab[88043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:55
		_go_fuzz_dep_.CoverTab[88044]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:55
		// _ = "end of CoverTab[88044]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:55
	// _ = "end of CoverTab[88040]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:55
	_go_fuzz_dep_.CoverTab[88041]++
													expectedMsgType := msgtype.KRB_ERROR
													if k.MsgType != expectedMsgType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:57
		_go_fuzz_dep_.CoverTab[88045]++
														return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate a KRB_ERROR. Expected: %v; Actual: %v", expectedMsgType, k.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:58
		// _ = "end of CoverTab[88045]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:59
		_go_fuzz_dep_.CoverTab[88046]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:59
		// _ = "end of CoverTab[88046]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:59
	// _ = "end of CoverTab[88041]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:59
	_go_fuzz_dep_.CoverTab[88042]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:60
	// _ = "end of CoverTab[88042]"
}

// Marshal a KRBError into bytes.
func (k *KRBError) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:64
	_go_fuzz_dep_.CoverTab[88047]++
													b, err := asn1.Marshal(*k)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:66
		_go_fuzz_dep_.CoverTab[88049]++
														return b, krberror.Errorf(err, krberror.EncodingError, "error marshaling KRBError")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:67
		// _ = "end of CoverTab[88049]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:68
		_go_fuzz_dep_.CoverTab[88050]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:68
		// _ = "end of CoverTab[88050]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:68
	// _ = "end of CoverTab[88047]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:68
	_go_fuzz_dep_.CoverTab[88048]++
													b = asn1tools.AddASNAppTag(b, asnAppTag.KRBError)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:70
	// _ = "end of CoverTab[88048]"
}

// Error method implementing error interface on KRBError struct.
func (k KRBError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:74
	_go_fuzz_dep_.CoverTab[88051]++
													etxt := fmt.Sprintf("KRB Error: %s", errorcode.Lookup(k.ErrorCode))
													if k.EText != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:76
		_go_fuzz_dep_.CoverTab[88053]++
														etxt = fmt.Sprintf("%s - %s", etxt, k.EText)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:77
		// _ = "end of CoverTab[88053]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:78
		_go_fuzz_dep_.CoverTab[88054]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:78
		// _ = "end of CoverTab[88054]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:78
	// _ = "end of CoverTab[88051]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:78
	_go_fuzz_dep_.CoverTab[88052]++
													return etxt
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:79
	// _ = "end of CoverTab[88052]"
}

func processUnmarshalReplyError(b []byte, err error) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:82
	_go_fuzz_dep_.CoverTab[88055]++
													switch err.(type) {
	case asn1.StructuralError:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:84
		_go_fuzz_dep_.CoverTab[88056]++
														var krberr KRBError
														tmperr := krberr.Unmarshal(b)
														if tmperr != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:87
			_go_fuzz_dep_.CoverTab[88059]++
															return krberror.Errorf(err, krberror.EncodingError, "failed to unmarshal KDC's reply")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:88
			// _ = "end of CoverTab[88059]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:89
			_go_fuzz_dep_.CoverTab[88060]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:89
			// _ = "end of CoverTab[88060]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:89
		// _ = "end of CoverTab[88056]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:89
		_go_fuzz_dep_.CoverTab[88057]++
														return krberr
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:90
		// _ = "end of CoverTab[88057]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:91
		_go_fuzz_dep_.CoverTab[88058]++
														return krberror.Errorf(err, krberror.EncodingError, "failed to unmarshal KDC's reply")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:92
		// _ = "end of CoverTab[88058]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:93
	// _ = "end of CoverTab[88055]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBError.go:94
var _ = _go_fuzz_dep_.CoverTab
