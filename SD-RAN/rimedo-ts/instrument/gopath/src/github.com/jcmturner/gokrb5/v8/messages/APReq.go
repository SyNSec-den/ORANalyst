//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:1
)

import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

type marshalAPReq struct {
	PVNO		int		`asn1:"explicit,tag:0"`
	MsgType		int		`asn1:"explicit,tag:1"`
	APOptions	asn1.BitString	`asn1:"explicit,tag:2"`
	// Ticket needs to be a raw value as it is wrapped in an APPLICATION tag
	Ticket			asn1.RawValue		`asn1:"explicit,tag:3"`
	EncryptedAuthenticator	types.EncryptedData	`asn1:"explicit,tag:4"`
}

// APReq implements RFC 4120 KRB_AP_REQ: https://tools.ietf.org/html/rfc4120#section-5.5.1.
type APReq struct {
	PVNO			int			`asn1:"explicit,tag:0"`
	MsgType			int			`asn1:"explicit,tag:1"`
	APOptions		asn1.BitString		`asn1:"explicit,tag:2"`
	Ticket			Ticket			`asn1:"explicit,tag:3"`
	EncryptedAuthenticator	types.EncryptedData	`asn1:"explicit,tag:4"`
	Authenticator		types.Authenticator	`asn1:"optional"`
}

// NewAPReq generates a new KRB_AP_REQ struct.
func NewAPReq(tkt Ticket, sessionKey types.EncryptionKey, auth types.Authenticator) (APReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:40
	_go_fuzz_dep_.CoverTab[87675]++
												var a APReq
												ed, err := encryptAuthenticator(auth, sessionKey, tkt)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:43
		_go_fuzz_dep_.CoverTab[87677]++
													return a, krberror.Errorf(err, krberror.KRBMsgError, "error creating Authenticator for AP_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:44
		// _ = "end of CoverTab[87677]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:45
		_go_fuzz_dep_.CoverTab[87678]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:45
		// _ = "end of CoverTab[87678]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:45
	// _ = "end of CoverTab[87675]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:45
	_go_fuzz_dep_.CoverTab[87676]++
												a = APReq{
		PVNO:			iana.PVNO,
		MsgType:		msgtype.KRB_AP_REQ,
		APOptions:		types.NewKrbFlags(),
		Ticket:			tkt,
		EncryptedAuthenticator:	ed,
	}
												return a, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:53
	// _ = "end of CoverTab[87676]"
}

// Encrypt Authenticator
func encryptAuthenticator(a types.Authenticator, sessionKey types.EncryptionKey, tkt Ticket) (types.EncryptedData, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:57
	_go_fuzz_dep_.CoverTab[87679]++
												var ed types.EncryptedData
												m, err := a.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:60
		_go_fuzz_dep_.CoverTab[87682]++
													return ed, krberror.Errorf(err, krberror.EncodingError, "marshaling error of EncryptedData form of Authenticator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:61
		// _ = "end of CoverTab[87682]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:62
		_go_fuzz_dep_.CoverTab[87683]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:62
		// _ = "end of CoverTab[87683]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:62
	// _ = "end of CoverTab[87679]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:62
	_go_fuzz_dep_.CoverTab[87680]++
												usage := authenticatorKeyUsage(tkt.SName)
												ed, err = crypto.GetEncryptedData(m, sessionKey, uint32(usage), tkt.EncPart.KVNO)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:65
		_go_fuzz_dep_.CoverTab[87684]++
													return ed, krberror.Errorf(err, krberror.EncryptingError, "error encrypting Authenticator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:66
		// _ = "end of CoverTab[87684]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:67
		_go_fuzz_dep_.CoverTab[87685]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:67
		// _ = "end of CoverTab[87685]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:67
	// _ = "end of CoverTab[87680]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:67
	_go_fuzz_dep_.CoverTab[87681]++
												return ed, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:68
	// _ = "end of CoverTab[87681]"
}

// DecryptAuthenticator decrypts the Authenticator within the AP_REQ.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:71
// sessionKey may simply be the key within the decrypted EncPart of the ticket within the AP_REQ.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:73
func (a *APReq) DecryptAuthenticator(sessionKey types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:73
	_go_fuzz_dep_.CoverTab[87686]++
												usage := authenticatorKeyUsage(a.Ticket.SName)
												ab, e := crypto.DecryptEncPart(a.EncryptedAuthenticator, sessionKey, uint32(usage))
												if e != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:76
		_go_fuzz_dep_.CoverTab[87689]++
													return fmt.Errorf("error decrypting authenticator: %v", e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:77
		// _ = "end of CoverTab[87689]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:78
		_go_fuzz_dep_.CoverTab[87690]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:78
		// _ = "end of CoverTab[87690]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:78
	// _ = "end of CoverTab[87686]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:78
	_go_fuzz_dep_.CoverTab[87687]++
												err := a.Authenticator.Unmarshal(ab)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:80
		_go_fuzz_dep_.CoverTab[87691]++
													return fmt.Errorf("error unmarshaling authenticator: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:81
		// _ = "end of CoverTab[87691]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:82
		_go_fuzz_dep_.CoverTab[87692]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:82
		// _ = "end of CoverTab[87692]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:82
	// _ = "end of CoverTab[87687]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:82
	_go_fuzz_dep_.CoverTab[87688]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:83
	// _ = "end of CoverTab[87688]"
}

func authenticatorKeyUsage(pn types.PrincipalName) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:86
	_go_fuzz_dep_.CoverTab[87693]++
												if pn.NameString[0] == "krbtgt" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:87
		_go_fuzz_dep_.CoverTab[87695]++
													return keyusage.TGS_REQ_PA_TGS_REQ_AP_REQ_AUTHENTICATOR
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:88
		// _ = "end of CoverTab[87695]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:89
		_go_fuzz_dep_.CoverTab[87696]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:89
		// _ = "end of CoverTab[87696]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:89
	// _ = "end of CoverTab[87693]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:89
	_go_fuzz_dep_.CoverTab[87694]++
												return keyusage.AP_REQ_AUTHENTICATOR
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:90
	// _ = "end of CoverTab[87694]"
}

// Unmarshal bytes b into the APReq struct.
func (a *APReq) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:94
	_go_fuzz_dep_.CoverTab[87697]++
												var m marshalAPReq
												_, err := asn1.UnmarshalWithParams(b, &m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.APREQ))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:97
		_go_fuzz_dep_.CoverTab[87701]++
													return krberror.Errorf(err, krberror.EncodingError, "unmarshal error of AP_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:98
		// _ = "end of CoverTab[87701]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:99
		_go_fuzz_dep_.CoverTab[87702]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:99
		// _ = "end of CoverTab[87702]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:99
	// _ = "end of CoverTab[87697]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:99
	_go_fuzz_dep_.CoverTab[87698]++
												if m.MsgType != msgtype.KRB_AP_REQ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:100
		_go_fuzz_dep_.CoverTab[87703]++
													return NewKRBError(types.PrincipalName{}, "", errorcode.KRB_AP_ERR_MSG_TYPE, errorcode.Lookup(errorcode.KRB_AP_ERR_MSG_TYPE))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:101
		// _ = "end of CoverTab[87703]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:102
		_go_fuzz_dep_.CoverTab[87704]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:102
		// _ = "end of CoverTab[87704]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:102
	// _ = "end of CoverTab[87698]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:102
	_go_fuzz_dep_.CoverTab[87699]++
												a.PVNO = m.PVNO
												a.MsgType = m.MsgType
												a.APOptions = m.APOptions
												a.EncryptedAuthenticator = m.EncryptedAuthenticator
												a.Ticket, err = unmarshalTicket(m.Ticket.Bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:108
		_go_fuzz_dep_.CoverTab[87705]++
													return krberror.Errorf(err, krberror.EncodingError, "unmarshaling error of Ticket within AP_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:109
		// _ = "end of CoverTab[87705]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:110
		_go_fuzz_dep_.CoverTab[87706]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:110
		// _ = "end of CoverTab[87706]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:110
	// _ = "end of CoverTab[87699]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:110
	_go_fuzz_dep_.CoverTab[87700]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:111
	// _ = "end of CoverTab[87700]"
}

// Marshal APReq struct.
func (a *APReq) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:115
	_go_fuzz_dep_.CoverTab[87707]++
												m := marshalAPReq{
		PVNO:			a.PVNO,
		MsgType:		a.MsgType,
		APOptions:		a.APOptions,
		EncryptedAuthenticator:	a.EncryptedAuthenticator,
	}
	var b []byte
	b, err := a.Ticket.Marshal()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:124
		_go_fuzz_dep_.CoverTab[87710]++
													return b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:125
		// _ = "end of CoverTab[87710]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:126
		_go_fuzz_dep_.CoverTab[87711]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:126
		// _ = "end of CoverTab[87711]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:126
	// _ = "end of CoverTab[87707]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:126
	_go_fuzz_dep_.CoverTab[87708]++
												m.Ticket = asn1.RawValue{
		Class:		asn1.ClassContextSpecific,
		IsCompound:	true,
		Tag:		3,
		Bytes:		b,
	}
	mk, err := asn1.Marshal(m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:134
		_go_fuzz_dep_.CoverTab[87712]++
													return mk, krberror.Errorf(err, krberror.EncodingError, "marshaling error of AP_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:135
		// _ = "end of CoverTab[87712]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:136
		_go_fuzz_dep_.CoverTab[87713]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:136
		// _ = "end of CoverTab[87713]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:136
	// _ = "end of CoverTab[87708]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:136
	_go_fuzz_dep_.CoverTab[87709]++
												mk = asn1tools.AddASNAppTag(mk, asnAppTag.APREQ)
												return mk, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:138
	// _ = "end of CoverTab[87709]"
}

// Verify an AP_REQ using service's keytab, spn and max acceptable clock skew duration.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:141
// The service ticket encrypted part and authenticator will be decrypted as part of this operation.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:143
func (a *APReq) Verify(kt *keytab.Keytab, d time.Duration, cAddr types.HostAddress, snameOverride *types.PrincipalName) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:143
	_go_fuzz_dep_.CoverTab[87714]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:157
	sname := &a.Ticket.SName
	if snameOverride != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:158
		_go_fuzz_dep_.CoverTab[87722]++
													sname = snameOverride
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:159
		// _ = "end of CoverTab[87722]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:160
		_go_fuzz_dep_.CoverTab[87723]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:160
		// _ = "end of CoverTab[87723]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:160
	// _ = "end of CoverTab[87714]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:160
	_go_fuzz_dep_.CoverTab[87715]++
												err := a.Ticket.DecryptEncPart(kt, sname)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:162
		_go_fuzz_dep_.CoverTab[87724]++
													return false, krberror.Errorf(err, krberror.DecryptingError, "error decrypting encpart of service ticket provided")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:163
		// _ = "end of CoverTab[87724]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:164
		_go_fuzz_dep_.CoverTab[87725]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:164
		// _ = "end of CoverTab[87725]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:164
	// _ = "end of CoverTab[87715]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:164
	_go_fuzz_dep_.CoverTab[87716]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:167
	ok, err := a.Ticket.Valid(d)
	if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:168
		_go_fuzz_dep_.CoverTab[87726]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:168
		return !ok
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:168
		// _ = "end of CoverTab[87726]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:168
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:168
		_go_fuzz_dep_.CoverTab[87727]++
													return ok, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:169
		// _ = "end of CoverTab[87727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:170
		_go_fuzz_dep_.CoverTab[87728]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:170
		// _ = "end of CoverTab[87728]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:170
	// _ = "end of CoverTab[87716]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:170
	_go_fuzz_dep_.CoverTab[87717]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:173
	if len(a.Ticket.DecryptedEncPart.CAddr) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:173
		_go_fuzz_dep_.CoverTab[87729]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:176
		if !types.HostAddressesContains(a.Ticket.DecryptedEncPart.CAddr, cAddr) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:176
			_go_fuzz_dep_.CoverTab[87730]++
														return false, NewKRBError(a.Ticket.SName, a.Ticket.Realm, errorcode.KRB_AP_ERR_BADADDR, "client address not within the list contained in the service ticket")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:177
			// _ = "end of CoverTab[87730]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:178
			_go_fuzz_dep_.CoverTab[87731]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:178
			// _ = "end of CoverTab[87731]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:178
		// _ = "end of CoverTab[87729]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:179
		_go_fuzz_dep_.CoverTab[87732]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:179
		// _ = "end of CoverTab[87732]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:179
	// _ = "end of CoverTab[87717]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:179
	_go_fuzz_dep_.CoverTab[87718]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:182
	err = a.DecryptAuthenticator(a.Ticket.DecryptedEncPart.Key)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:183
		_go_fuzz_dep_.CoverTab[87733]++
													return false, NewKRBError(a.Ticket.SName, a.Ticket.Realm, errorcode.KRB_AP_ERR_BAD_INTEGRITY, "could not decrypt authenticator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:184
		// _ = "end of CoverTab[87733]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:185
		_go_fuzz_dep_.CoverTab[87734]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:185
		// _ = "end of CoverTab[87734]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:185
	// _ = "end of CoverTab[87718]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:185
	_go_fuzz_dep_.CoverTab[87719]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:188
	if !a.Authenticator.CName.Equal(a.Ticket.DecryptedEncPart.CName) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:188
		_go_fuzz_dep_.CoverTab[87735]++
													return false, NewKRBError(a.Ticket.SName, a.Ticket.Realm, errorcode.KRB_AP_ERR_BADMATCH, "CName in Authenticator does not match that in service ticket")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:189
		// _ = "end of CoverTab[87735]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:190
		_go_fuzz_dep_.CoverTab[87736]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:190
		// _ = "end of CoverTab[87736]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:190
	// _ = "end of CoverTab[87719]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:190
	_go_fuzz_dep_.CoverTab[87720]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:193
	ct := a.Authenticator.CTime.Add(time.Duration(a.Authenticator.Cusec) * time.Microsecond)
	t := time.Now().UTC()
	if t.Sub(ct) > d || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:195
		_go_fuzz_dep_.CoverTab[87737]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:195
		return ct.Sub(t) > d
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:195
		// _ = "end of CoverTab[87737]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:195
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:195
		_go_fuzz_dep_.CoverTab[87738]++
													return false, NewKRBError(a.Ticket.SName, a.Ticket.Realm, errorcode.KRB_AP_ERR_SKEW, fmt.Sprintf("clock skew with client too large. greater than %v seconds", d))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:196
		// _ = "end of CoverTab[87738]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:197
		_go_fuzz_dep_.CoverTab[87739]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:197
		// _ = "end of CoverTab[87739]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:197
	// _ = "end of CoverTab[87720]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:197
	_go_fuzz_dep_.CoverTab[87721]++
												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:198
	// _ = "end of CoverTab[87721]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:199
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/APReq.go:199
var _ = _go_fuzz_dep_.CoverTab
