//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:6
import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/flags"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/iana/patype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

type marshalKDCRep struct {
	PVNO	int			`asn1:"explicit,tag:0"`
	MsgType	int			`asn1:"explicit,tag:1"`
	PAData	types.PADataSequence	`asn1:"explicit,optional,tag:2"`
	CRealm	string			`asn1:"generalstring,explicit,tag:3"`
	CName	types.PrincipalName	`asn1:"explicit,tag:4"`
	// Ticket needs to be a raw value as it is wrapped in an APPLICATION tag
	Ticket	asn1.RawValue		`asn1:"explicit,tag:5"`
	EncPart	types.EncryptedData	`asn1:"explicit,tag:6"`
}

// KDCRepFields represents the KRB_KDC_REP fields.
type KDCRepFields struct {
	PVNO			int
	MsgType			int
	PAData			[]types.PAData
	CRealm			string
	CName			types.PrincipalName
	Ticket			Ticket
	EncPart			types.EncryptedData
	DecryptedEncPart	EncKDCRepPart
}

// ASRep implements RFC 4120 KRB_AS_REP: https://tools.ietf.org/html/rfc4120#section-5.4.2.
type ASRep struct {
	KDCRepFields
}

// TGSRep implements RFC 4120 KRB_TGS_REP: https://tools.ietf.org/html/rfc4120#section-5.4.2.
type TGSRep struct {
	KDCRepFields
}

// EncKDCRepPart is the encrypted part of KRB_KDC_REP.
type EncKDCRepPart struct {
	Key		types.EncryptionKey	`asn1:"explicit,tag:0"`
	LastReqs	[]LastReq		`asn1:"explicit,tag:1"`
	Nonce		int			`asn1:"explicit,tag:2"`
	KeyExpiration	time.Time		`asn1:"generalized,explicit,optional,tag:3"`
	Flags		asn1.BitString		`asn1:"explicit,tag:4"`
	AuthTime	time.Time		`asn1:"generalized,explicit,tag:5"`
	StartTime	time.Time		`asn1:"generalized,explicit,optional,tag:6"`
	EndTime		time.Time		`asn1:"generalized,explicit,tag:7"`
	RenewTill	time.Time		`asn1:"generalized,explicit,optional,tag:8"`
	SRealm		string			`asn1:"generalstring,explicit,tag:9"`
	SName		types.PrincipalName	`asn1:"explicit,tag:10"`
	CAddr		[]types.HostAddress	`asn1:"explicit,optional,tag:11"`
	EncPAData	types.PADataSequence	`asn1:"explicit,optional,tag:12"`
}

// LastReq part of KRB_KDC_REP.
type LastReq struct {
	LRType	int32		`asn1:"explicit,tag:0"`
	LRValue	time.Time	`asn1:"generalized,explicit,tag:1"`
}

// Unmarshal bytes b into the ASRep struct.
func (k *ASRep) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:81
	_go_fuzz_dep_.CoverTab[87740]++
												var m marshalKDCRep
												_, err := asn1.UnmarshalWithParams(b, &m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.ASREP))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:84
		_go_fuzz_dep_.CoverTab[87744]++
													return processUnmarshalReplyError(b, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:85
		// _ = "end of CoverTab[87744]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:86
		_go_fuzz_dep_.CoverTab[87745]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:86
		// _ = "end of CoverTab[87745]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:86
	// _ = "end of CoverTab[87740]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:86
	_go_fuzz_dep_.CoverTab[87741]++
												if m.MsgType != msgtype.KRB_AS_REP {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:87
		_go_fuzz_dep_.CoverTab[87746]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate an AS_REP. Expected: %v; Actual: %v", msgtype.KRB_AS_REP, m.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:88
		// _ = "end of CoverTab[87746]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:89
		_go_fuzz_dep_.CoverTab[87747]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:89
		// _ = "end of CoverTab[87747]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:89
	// _ = "end of CoverTab[87741]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:89
	_go_fuzz_dep_.CoverTab[87742]++

												tkt, err := unmarshalTicket(m.Ticket.Bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:92
		_go_fuzz_dep_.CoverTab[87748]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling Ticket within AS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:93
		// _ = "end of CoverTab[87748]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:94
		_go_fuzz_dep_.CoverTab[87749]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:94
		// _ = "end of CoverTab[87749]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:94
	// _ = "end of CoverTab[87742]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:94
	_go_fuzz_dep_.CoverTab[87743]++
												k.KDCRepFields = KDCRepFields{
		PVNO:		m.PVNO,
		MsgType:	m.MsgType,
		PAData:		m.PAData,
		CRealm:		m.CRealm,
		CName:		m.CName,
		Ticket:		tkt,
		EncPart:	m.EncPart,
	}
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:104
	// _ = "end of CoverTab[87743]"
}

// Marshal ASRep struct.
func (k *ASRep) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:108
	_go_fuzz_dep_.CoverTab[87750]++
												m := marshalKDCRep{
		PVNO:		k.PVNO,
		MsgType:	k.MsgType,
		PAData:		k.PAData,
		CRealm:		k.CRealm,
		CName:		k.CName,
		EncPart:	k.EncPart,
	}
	b, err := k.Ticket.Marshal()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:118
		_go_fuzz_dep_.CoverTab[87753]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:119
		// _ = "end of CoverTab[87753]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:120
		_go_fuzz_dep_.CoverTab[87754]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:120
		// _ = "end of CoverTab[87754]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:120
	// _ = "end of CoverTab[87750]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:120
	_go_fuzz_dep_.CoverTab[87751]++
												m.Ticket = asn1.RawValue{
		Class:		asn1.ClassContextSpecific,
		IsCompound:	true,
		Tag:		5,
		Bytes:		b,
	}
	mk, err := asn1.Marshal(m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:128
		_go_fuzz_dep_.CoverTab[87755]++
													return mk, krberror.Errorf(err, krberror.EncodingError, "error marshaling AS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:129
		// _ = "end of CoverTab[87755]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:130
		_go_fuzz_dep_.CoverTab[87756]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:130
		// _ = "end of CoverTab[87756]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:130
	// _ = "end of CoverTab[87751]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:130
	_go_fuzz_dep_.CoverTab[87752]++
												mk = asn1tools.AddASNAppTag(mk, asnAppTag.ASREP)
												return mk, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:132
	// _ = "end of CoverTab[87752]"
}

// Unmarshal bytes b into the TGSRep struct.
func (k *TGSRep) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:136
	_go_fuzz_dep_.CoverTab[87757]++
												var m marshalKDCRep
												_, err := asn1.UnmarshalWithParams(b, &m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.TGSREP))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:139
		_go_fuzz_dep_.CoverTab[87761]++
													return processUnmarshalReplyError(b, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:140
		// _ = "end of CoverTab[87761]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:141
		_go_fuzz_dep_.CoverTab[87762]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:141
		// _ = "end of CoverTab[87762]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:141
	// _ = "end of CoverTab[87757]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:141
	_go_fuzz_dep_.CoverTab[87758]++
												if m.MsgType != msgtype.KRB_TGS_REP {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:142
		_go_fuzz_dep_.CoverTab[87763]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate an TGS_REP. Expected: %v; Actual: %v", msgtype.KRB_TGS_REP, m.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:143
		// _ = "end of CoverTab[87763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:144
		_go_fuzz_dep_.CoverTab[87764]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:144
		// _ = "end of CoverTab[87764]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:144
	// _ = "end of CoverTab[87758]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:144
	_go_fuzz_dep_.CoverTab[87759]++

												tkt, err := unmarshalTicket(m.Ticket.Bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:147
		_go_fuzz_dep_.CoverTab[87765]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling Ticket within TGS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:148
		// _ = "end of CoverTab[87765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:149
		_go_fuzz_dep_.CoverTab[87766]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:149
		// _ = "end of CoverTab[87766]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:149
	// _ = "end of CoverTab[87759]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:149
	_go_fuzz_dep_.CoverTab[87760]++
												k.KDCRepFields = KDCRepFields{
		PVNO:		m.PVNO,
		MsgType:	m.MsgType,
		PAData:		m.PAData,
		CRealm:		m.CRealm,
		CName:		m.CName,
		Ticket:		tkt,
		EncPart:	m.EncPart,
	}
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:159
	// _ = "end of CoverTab[87760]"
}

// Marshal TGSRep struct.
func (k *TGSRep) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:163
	_go_fuzz_dep_.CoverTab[87767]++
												m := marshalKDCRep{
		PVNO:		k.PVNO,
		MsgType:	k.MsgType,
		PAData:		k.PAData,
		CRealm:		k.CRealm,
		CName:		k.CName,
		EncPart:	k.EncPart,
	}
	b, err := k.Ticket.Marshal()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:173
		_go_fuzz_dep_.CoverTab[87770]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:174
		// _ = "end of CoverTab[87770]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:175
		_go_fuzz_dep_.CoverTab[87771]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:175
		// _ = "end of CoverTab[87771]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:175
	// _ = "end of CoverTab[87767]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:175
	_go_fuzz_dep_.CoverTab[87768]++
												m.Ticket = asn1.RawValue{
		Class:		asn1.ClassContextSpecific,
		IsCompound:	true,
		Tag:		5,
		Bytes:		b,
	}
	mk, err := asn1.Marshal(m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:183
		_go_fuzz_dep_.CoverTab[87772]++
													return mk, krberror.Errorf(err, krberror.EncodingError, "error marshaling TGS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:184
		// _ = "end of CoverTab[87772]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:185
		_go_fuzz_dep_.CoverTab[87773]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:185
		// _ = "end of CoverTab[87773]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:185
	// _ = "end of CoverTab[87768]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:185
	_go_fuzz_dep_.CoverTab[87769]++
												mk = asn1tools.AddASNAppTag(mk, asnAppTag.TGSREP)
												return mk, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:187
	// _ = "end of CoverTab[87769]"
}

// Unmarshal bytes b into encrypted part of KRB_KDC_REP.
func (e *EncKDCRepPart) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:191
	_go_fuzz_dep_.CoverTab[87774]++
												_, err := asn1.UnmarshalWithParams(b, e, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.EncASRepPart))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:193
		_go_fuzz_dep_.CoverTab[87776]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:197
		_, err = asn1.UnmarshalWithParams(b, e, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.EncTGSRepPart))
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:198
			_go_fuzz_dep_.CoverTab[87777]++
														return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling encrypted part within KDC_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:199
			// _ = "end of CoverTab[87777]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:200
			_go_fuzz_dep_.CoverTab[87778]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:200
			// _ = "end of CoverTab[87778]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:200
		// _ = "end of CoverTab[87776]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:201
		_go_fuzz_dep_.CoverTab[87779]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:201
		// _ = "end of CoverTab[87779]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:201
	// _ = "end of CoverTab[87774]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:201
	_go_fuzz_dep_.CoverTab[87775]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:202
	// _ = "end of CoverTab[87775]"
}

// Marshal encrypted part of KRB_KDC_REP.
func (e *EncKDCRepPart) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:206
	_go_fuzz_dep_.CoverTab[87780]++
												b, err := asn1.Marshal(*e)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:208
		_go_fuzz_dep_.CoverTab[87782]++
													return b, krberror.Errorf(err, krberror.EncodingError, "marshaling error of AS_REP encpart")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:209
		// _ = "end of CoverTab[87782]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:210
		_go_fuzz_dep_.CoverTab[87783]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:210
		// _ = "end of CoverTab[87783]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:210
	// _ = "end of CoverTab[87780]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:210
	_go_fuzz_dep_.CoverTab[87781]++
												b = asn1tools.AddASNAppTag(b, asnAppTag.EncASRepPart)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:212
	// _ = "end of CoverTab[87781]"
}

// DecryptEncPart decrypts the encrypted part of an AS_REP.
func (k *ASRep) DecryptEncPart(c *credentials.Credentials) (types.EncryptionKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:216
	_go_fuzz_dep_.CoverTab[87784]++
												var key types.EncryptionKey
												var err error
												if c.HasKeytab() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:219
		_go_fuzz_dep_.CoverTab[87790]++
													key, _, err = c.Keytab().GetEncryptionKey(k.CName, k.CRealm, k.EncPart.KVNO, k.EncPart.EType)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:221
			_go_fuzz_dep_.CoverTab[87791]++
														return key, krberror.Errorf(err, krberror.DecryptingError, "error decrypting AS_REP encrypted part")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:222
			// _ = "end of CoverTab[87791]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:223
			_go_fuzz_dep_.CoverTab[87792]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:223
			// _ = "end of CoverTab[87792]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:223
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:223
		// _ = "end of CoverTab[87790]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:224
		_go_fuzz_dep_.CoverTab[87793]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:224
		// _ = "end of CoverTab[87793]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:224
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:224
	// _ = "end of CoverTab[87784]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:224
	_go_fuzz_dep_.CoverTab[87785]++
												if c.HasPassword() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:225
		_go_fuzz_dep_.CoverTab[87794]++
													key, _, err = crypto.GetKeyFromPassword(c.Password(), k.CName, k.CRealm, k.EncPart.EType, k.PAData)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:227
			_go_fuzz_dep_.CoverTab[87795]++
														return key, krberror.Errorf(err, krberror.DecryptingError, "error decrypting AS_REP encrypted part")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:228
			// _ = "end of CoverTab[87795]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:229
			_go_fuzz_dep_.CoverTab[87796]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:229
			// _ = "end of CoverTab[87796]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:229
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:229
		// _ = "end of CoverTab[87794]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:230
		_go_fuzz_dep_.CoverTab[87797]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:230
		// _ = "end of CoverTab[87797]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:230
	// _ = "end of CoverTab[87785]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:230
	_go_fuzz_dep_.CoverTab[87786]++
												if !c.HasKeytab() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:231
		_go_fuzz_dep_.CoverTab[87798]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:231
		return !c.HasPassword()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:231
		// _ = "end of CoverTab[87798]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:231
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:231
		_go_fuzz_dep_.CoverTab[87799]++
													return key, krberror.NewErrorf(krberror.DecryptingError, "no secret available in credentials to perform decryption of AS_REP encrypted part")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:232
		// _ = "end of CoverTab[87799]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:233
		_go_fuzz_dep_.CoverTab[87800]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:233
		// _ = "end of CoverTab[87800]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:233
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:233
	// _ = "end of CoverTab[87786]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:233
	_go_fuzz_dep_.CoverTab[87787]++
												b, err := crypto.DecryptEncPart(k.EncPart, key, keyusage.AS_REP_ENCPART)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:235
		_go_fuzz_dep_.CoverTab[87801]++
													return key, krberror.Errorf(err, krberror.DecryptingError, "error decrypting AS_REP encrypted part")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:236
		// _ = "end of CoverTab[87801]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:237
		_go_fuzz_dep_.CoverTab[87802]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:237
		// _ = "end of CoverTab[87802]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:237
	// _ = "end of CoverTab[87787]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:237
	_go_fuzz_dep_.CoverTab[87788]++
												var denc EncKDCRepPart
												err = denc.Unmarshal(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:240
		_go_fuzz_dep_.CoverTab[87803]++
													return key, krberror.Errorf(err, krberror.EncodingError, "error unmarshaling decrypted encpart of AS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:241
		// _ = "end of CoverTab[87803]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:242
		_go_fuzz_dep_.CoverTab[87804]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:242
		// _ = "end of CoverTab[87804]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:242
	// _ = "end of CoverTab[87788]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:242
	_go_fuzz_dep_.CoverTab[87789]++
												k.DecryptedEncPart = denc
												return key, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:244
	// _ = "end of CoverTab[87789]"
}

// Verify checks the validity of AS_REP message.
func (k *ASRep) Verify(cfg *config.Config, creds *credentials.Credentials, asReq ASReq) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:248
	_go_fuzz_dep_.CoverTab[87805]++

												if !k.CName.Equal(asReq.ReqBody.CName) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:250
		_go_fuzz_dep_.CoverTab[87815]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "CName in response does not match what was requested. Requested: %+v; Reply: %+v", asReq.ReqBody.CName, k.CName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:251
		// _ = "end of CoverTab[87815]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:252
		_go_fuzz_dep_.CoverTab[87816]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:252
		// _ = "end of CoverTab[87816]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:252
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:252
	// _ = "end of CoverTab[87805]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:252
	_go_fuzz_dep_.CoverTab[87806]++
												if k.CRealm != asReq.ReqBody.Realm {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:253
		_go_fuzz_dep_.CoverTab[87817]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "CRealm in response does not match what was requested. Requested: %s; Reply: %s", asReq.ReqBody.Realm, k.CRealm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:254
		// _ = "end of CoverTab[87817]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:255
		_go_fuzz_dep_.CoverTab[87818]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:255
		// _ = "end of CoverTab[87818]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:255
	// _ = "end of CoverTab[87806]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:255
	_go_fuzz_dep_.CoverTab[87807]++
												key, err := k.DecryptEncPart(creds)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:257
		_go_fuzz_dep_.CoverTab[87819]++
													return false, krberror.Errorf(err, krberror.DecryptingError, "error decrypting EncPart of AS_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:258
		// _ = "end of CoverTab[87819]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:259
		_go_fuzz_dep_.CoverTab[87820]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:259
		// _ = "end of CoverTab[87820]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:259
	// _ = "end of CoverTab[87807]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:259
	_go_fuzz_dep_.CoverTab[87808]++
												if k.DecryptedEncPart.Nonce != asReq.ReqBody.Nonce {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:260
		_go_fuzz_dep_.CoverTab[87821]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "possible replay attack, nonce in response does not match that in request")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:261
		// _ = "end of CoverTab[87821]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:262
		_go_fuzz_dep_.CoverTab[87822]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:262
		// _ = "end of CoverTab[87822]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:262
	// _ = "end of CoverTab[87808]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:262
	_go_fuzz_dep_.CoverTab[87809]++
												if !k.DecryptedEncPart.SName.Equal(asReq.ReqBody.SName) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:263
		_go_fuzz_dep_.CoverTab[87823]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "SName in response does not match what was requested. Requested: %v; Reply: %v", asReq.ReqBody.SName, k.DecryptedEncPart.SName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:264
		// _ = "end of CoverTab[87823]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:265
		_go_fuzz_dep_.CoverTab[87824]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:265
		// _ = "end of CoverTab[87824]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:265
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:265
	// _ = "end of CoverTab[87809]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:265
	_go_fuzz_dep_.CoverTab[87810]++
												if k.DecryptedEncPart.SRealm != asReq.ReqBody.Realm {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:266
		_go_fuzz_dep_.CoverTab[87825]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "SRealm in response does not match what was requested. Requested: %s; Reply: %s", asReq.ReqBody.Realm, k.DecryptedEncPart.SRealm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:267
		// _ = "end of CoverTab[87825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:268
		_go_fuzz_dep_.CoverTab[87826]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:268
		// _ = "end of CoverTab[87826]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:268
	// _ = "end of CoverTab[87810]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:268
	_go_fuzz_dep_.CoverTab[87811]++
												if len(asReq.ReqBody.Addresses) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:269
		_go_fuzz_dep_.CoverTab[87827]++
													if !types.HostAddressesEqual(k.DecryptedEncPart.CAddr, asReq.ReqBody.Addresses) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:270
			_go_fuzz_dep_.CoverTab[87828]++
														return false, krberror.NewErrorf(krberror.KRBMsgError, "addresses listed in the AS_REP does not match those listed in the AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:271
			// _ = "end of CoverTab[87828]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:272
			_go_fuzz_dep_.CoverTab[87829]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:272
			// _ = "end of CoverTab[87829]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:272
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:272
		// _ = "end of CoverTab[87827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:273
		_go_fuzz_dep_.CoverTab[87830]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:273
		// _ = "end of CoverTab[87830]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:273
	// _ = "end of CoverTab[87811]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:273
	_go_fuzz_dep_.CoverTab[87812]++
												t := time.Now().UTC()
												if t.Sub(k.DecryptedEncPart.AuthTime) > cfg.LibDefaults.Clockskew || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:275
		_go_fuzz_dep_.CoverTab[87831]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:275
		return k.DecryptedEncPart.AuthTime.Sub(t) > cfg.LibDefaults.Clockskew
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:275
		// _ = "end of CoverTab[87831]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:275
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:275
		_go_fuzz_dep_.CoverTab[87832]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "clock skew with KDC too large. Greater than %v seconds", cfg.LibDefaults.Clockskew.Seconds())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:276
		// _ = "end of CoverTab[87832]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:277
		_go_fuzz_dep_.CoverTab[87833]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:277
		// _ = "end of CoverTab[87833]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:277
	// _ = "end of CoverTab[87812]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:277
	_go_fuzz_dep_.CoverTab[87813]++

												if asReq.PAData.Contains(patype.PA_REQ_ENC_PA_REP) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:279
		_go_fuzz_dep_.CoverTab[87834]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:279
		return types.IsFlagSet(&k.DecryptedEncPart.Flags, flags.EncPARep)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:279
		// _ = "end of CoverTab[87834]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:279
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:279
		_go_fuzz_dep_.CoverTab[87835]++
													if len(k.DecryptedEncPart.EncPAData) < 2 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:280
			_go_fuzz_dep_.CoverTab[87837]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:280
			return !k.DecryptedEncPart.EncPAData.Contains(patype.PA_FX_FAST)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:280
			// _ = "end of CoverTab[87837]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:280
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:280
			_go_fuzz_dep_.CoverTab[87838]++
														return false, krberror.NewErrorf(krberror.KRBMsgError, "KDC did not respond appropriately to FAST negotiation")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:281
			// _ = "end of CoverTab[87838]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:282
			_go_fuzz_dep_.CoverTab[87839]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:282
			// _ = "end of CoverTab[87839]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:282
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:282
		// _ = "end of CoverTab[87835]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:282
		_go_fuzz_dep_.CoverTab[87836]++
													for _, pa := range k.DecryptedEncPart.EncPAData {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:283
			_go_fuzz_dep_.CoverTab[87840]++
														if pa.PADataType == patype.PA_REQ_ENC_PA_REP {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:284
				_go_fuzz_dep_.CoverTab[87841]++
															var pafast types.PAReqEncPARep
															err := pafast.Unmarshal(pa.PADataValue)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:287
					_go_fuzz_dep_.CoverTab[87844]++
																return false, krberror.Errorf(err, krberror.EncodingError, "KDC FAST negotiation response error, could not unmarshal PA_REQ_ENC_PA_REP")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:288
					// _ = "end of CoverTab[87844]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:289
					_go_fuzz_dep_.CoverTab[87845]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:289
					// _ = "end of CoverTab[87845]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:289
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:289
				// _ = "end of CoverTab[87841]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:289
				_go_fuzz_dep_.CoverTab[87842]++
															etype, err := crypto.GetChksumEtype(pafast.ChksumType)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:291
					_go_fuzz_dep_.CoverTab[87846]++
																return false, krberror.Errorf(err, krberror.ChksumError, "KDC FAST negotiation response error")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:292
					// _ = "end of CoverTab[87846]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:293
					_go_fuzz_dep_.CoverTab[87847]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:293
					// _ = "end of CoverTab[87847]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:293
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:293
				// _ = "end of CoverTab[87842]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:293
				_go_fuzz_dep_.CoverTab[87843]++
															ab, _ := asReq.Marshal()
															if !etype.VerifyChecksum(key.KeyValue, ab, pafast.Chksum, keyusage.KEY_USAGE_AS_REQ) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:295
					_go_fuzz_dep_.CoverTab[87848]++
																return false, krberror.Errorf(err, krberror.ChksumError, "KDC FAST negotiation response checksum invalid")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:296
					// _ = "end of CoverTab[87848]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:297
					_go_fuzz_dep_.CoverTab[87849]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:297
					// _ = "end of CoverTab[87849]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:297
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:297
				// _ = "end of CoverTab[87843]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:298
				_go_fuzz_dep_.CoverTab[87850]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:298
				// _ = "end of CoverTab[87850]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:298
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:298
			// _ = "end of CoverTab[87840]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:299
		// _ = "end of CoverTab[87836]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:300
		_go_fuzz_dep_.CoverTab[87851]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:300
		// _ = "end of CoverTab[87851]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:300
	// _ = "end of CoverTab[87813]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:300
	_go_fuzz_dep_.CoverTab[87814]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:301
	// _ = "end of CoverTab[87814]"
}

// DecryptEncPart decrypts the encrypted part of an TGS_REP.
func (k *TGSRep) DecryptEncPart(key types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:305
	_go_fuzz_dep_.CoverTab[87852]++
												b, err := crypto.DecryptEncPart(k.EncPart, key, keyusage.TGS_REP_ENCPART_SESSION_KEY)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:307
		_go_fuzz_dep_.CoverTab[87855]++
													return krberror.Errorf(err, krberror.DecryptingError, "error decrypting TGS_REP EncPart")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:308
		// _ = "end of CoverTab[87855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:309
		_go_fuzz_dep_.CoverTab[87856]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:309
		// _ = "end of CoverTab[87856]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:309
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:309
	// _ = "end of CoverTab[87852]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:309
	_go_fuzz_dep_.CoverTab[87853]++
												var denc EncKDCRepPart
												err = denc.Unmarshal(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:312
		_go_fuzz_dep_.CoverTab[87857]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling encrypted part")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:313
		// _ = "end of CoverTab[87857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:314
		_go_fuzz_dep_.CoverTab[87858]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:314
		// _ = "end of CoverTab[87858]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:314
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:314
	// _ = "end of CoverTab[87853]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:314
	_go_fuzz_dep_.CoverTab[87854]++
												k.DecryptedEncPart = denc
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:316
	// _ = "end of CoverTab[87854]"
}

// Verify checks the validity of the TGS_REP message.
func (k *TGSRep) Verify(cfg *config.Config, tgsReq TGSReq) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:320
	_go_fuzz_dep_.CoverTab[87859]++
												if !k.CName.Equal(tgsReq.ReqBody.CName) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:321
		_go_fuzz_dep_.CoverTab[87866]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "CName in response does not match what was requested. Requested: %+v; Reply: %+v", tgsReq.ReqBody.CName, k.CName)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:322
		// _ = "end of CoverTab[87866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:323
		_go_fuzz_dep_.CoverTab[87867]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:323
		// _ = "end of CoverTab[87867]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:323
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:323
	// _ = "end of CoverTab[87859]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:323
	_go_fuzz_dep_.CoverTab[87860]++
												if k.Ticket.Realm != tgsReq.ReqBody.Realm {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:324
		_go_fuzz_dep_.CoverTab[87868]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "realm in response ticket does not match what was requested. Requested: %s; Reply: %s", tgsReq.ReqBody.Realm, k.Ticket.Realm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:325
		// _ = "end of CoverTab[87868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:326
		_go_fuzz_dep_.CoverTab[87869]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:326
		// _ = "end of CoverTab[87869]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:326
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:326
	// _ = "end of CoverTab[87860]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:326
	_go_fuzz_dep_.CoverTab[87861]++
												if k.DecryptedEncPart.Nonce != tgsReq.ReqBody.Nonce {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:327
		_go_fuzz_dep_.CoverTab[87870]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "possible replay attack, nonce in response does not match that in request")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:328
		// _ = "end of CoverTab[87870]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:329
		_go_fuzz_dep_.CoverTab[87871]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:329
		// _ = "end of CoverTab[87871]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:329
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:329
	// _ = "end of CoverTab[87861]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:329
	_go_fuzz_dep_.CoverTab[87862]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:346
	if k.DecryptedEncPart.SRealm != tgsReq.ReqBody.Realm {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:346
		_go_fuzz_dep_.CoverTab[87872]++
													return false, krberror.NewErrorf(krberror.KRBMsgError, "SRealm in response does not match what was requested. Requested: %s; Reply: %s", tgsReq.ReqBody.Realm, k.DecryptedEncPart.SRealm)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:347
		// _ = "end of CoverTab[87872]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:348
		_go_fuzz_dep_.CoverTab[87873]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:348
		// _ = "end of CoverTab[87873]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:348
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:348
	// _ = "end of CoverTab[87862]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:348
	_go_fuzz_dep_.CoverTab[87863]++
												if len(k.DecryptedEncPart.CAddr) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:349
		_go_fuzz_dep_.CoverTab[87874]++
													if !types.HostAddressesEqual(k.DecryptedEncPart.CAddr, tgsReq.ReqBody.Addresses) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:350
			_go_fuzz_dep_.CoverTab[87875]++
														return false, krberror.NewErrorf(krberror.KRBMsgError, "addresses listed in the TGS_REP does not match those listed in the TGS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:351
			// _ = "end of CoverTab[87875]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:352
			_go_fuzz_dep_.CoverTab[87876]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:352
			// _ = "end of CoverTab[87876]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:352
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:352
		// _ = "end of CoverTab[87874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:353
		_go_fuzz_dep_.CoverTab[87877]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:353
		// _ = "end of CoverTab[87877]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:353
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:353
	// _ = "end of CoverTab[87863]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:353
	_go_fuzz_dep_.CoverTab[87864]++
												if time.Since(k.DecryptedEncPart.StartTime) > cfg.LibDefaults.Clockskew || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:354
		_go_fuzz_dep_.CoverTab[87878]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:354
		return k.DecryptedEncPart.StartTime.Sub(time.Now().UTC()) > cfg.LibDefaults.Clockskew
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:354
		// _ = "end of CoverTab[87878]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:354
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:354
		_go_fuzz_dep_.CoverTab[87879]++
													if time.Since(k.DecryptedEncPart.AuthTime) > cfg.LibDefaults.Clockskew || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:355
			_go_fuzz_dep_.CoverTab[87880]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:355
			return k.DecryptedEncPart.AuthTime.Sub(time.Now().UTC()) > cfg.LibDefaults.Clockskew
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:355
			// _ = "end of CoverTab[87880]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:355
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:355
			_go_fuzz_dep_.CoverTab[87881]++
														return false, krberror.NewErrorf(krberror.KRBMsgError, "clock skew with KDC too large. Greater than %v seconds.", cfg.LibDefaults.Clockskew.Seconds())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:356
			// _ = "end of CoverTab[87881]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:357
			_go_fuzz_dep_.CoverTab[87882]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:357
			// _ = "end of CoverTab[87882]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:357
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:357
		// _ = "end of CoverTab[87879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:358
		_go_fuzz_dep_.CoverTab[87883]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:358
		// _ = "end of CoverTab[87883]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:358
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:358
	// _ = "end of CoverTab[87864]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:358
	_go_fuzz_dep_.CoverTab[87865]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:359
	// _ = "end of CoverTab[87865]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:360
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCRep.go:360
var _ = _go_fuzz_dep_.CoverTab
