//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:6
import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/flags"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/iana/patype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

type marshalKDCReq struct {
	PVNO	int			`asn1:"explicit,tag:1"`
	MsgType	int			`asn1:"explicit,tag:2"`
	PAData	types.PADataSequence	`asn1:"explicit,optional,tag:3"`
	ReqBody	asn1.RawValue		`asn1:"explicit,tag:4"`
}

// KDCReqFields represents the KRB_KDC_REQ fields.
type KDCReqFields struct {
	PVNO	int
	MsgType	int
	PAData	types.PADataSequence
	ReqBody	KDCReqBody
	Renewal	bool
}

// ASReq implements RFC 4120 KRB_AS_REQ: https://tools.ietf.org/html/rfc4120#section-5.4.1.
type ASReq struct {
	KDCReqFields
}

// TGSReq implements RFC 4120 KRB_TGS_REQ: https://tools.ietf.org/html/rfc4120#section-5.4.1.
type TGSReq struct {
	KDCReqFields
}

type marshalKDCReqBody struct {
	KDCOptions	asn1.BitString		`asn1:"explicit,tag:0"`
	CName		types.PrincipalName	`asn1:"explicit,optional,tag:1"`
	Realm		string			`asn1:"generalstring,explicit,tag:2"`
	SName		types.PrincipalName	`asn1:"explicit,optional,tag:3"`
	From		time.Time		`asn1:"generalized,explicit,optional,tag:4"`
	Till		time.Time		`asn1:"generalized,explicit,tag:5"`
	RTime		time.Time		`asn1:"generalized,explicit,optional,tag:6"`
	Nonce		int			`asn1:"explicit,tag:7"`
	EType		[]int32			`asn1:"explicit,tag:8"`
	Addresses	[]types.HostAddress	`asn1:"explicit,optional,tag:9"`
	EncAuthData	types.EncryptedData	`asn1:"explicit,optional,tag:10"`
	// Ticket needs to be a raw value as it is wrapped in an APPLICATION tag
	AdditionalTickets	asn1.RawValue	`asn1:"explicit,optional,tag:11"`
}

// KDCReqBody implements the KRB_KDC_REQ request body.
type KDCReqBody struct {
	KDCOptions		asn1.BitString		`asn1:"explicit,tag:0"`
	CName			types.PrincipalName	`asn1:"explicit,optional,tag:1"`
	Realm			string			`asn1:"generalstring,explicit,tag:2"`
	SName			types.PrincipalName	`asn1:"explicit,optional,tag:3"`
	From			time.Time		`asn1:"generalized,explicit,optional,tag:4"`
	Till			time.Time		`asn1:"generalized,explicit,tag:5"`
	RTime			time.Time		`asn1:"generalized,explicit,optional,tag:6"`
	Nonce			int			`asn1:"explicit,tag:7"`
	EType			[]int32			`asn1:"explicit,tag:8"`
	Addresses		[]types.HostAddress	`asn1:"explicit,optional,tag:9"`
	EncAuthData		types.EncryptedData	`asn1:"explicit,optional,tag:10"`
	AdditionalTickets	[]Ticket		`asn1:"explicit,optional,tag:11"`
}

// NewASReqForTGT generates a new KRB_AS_REQ struct for a TGT request.
func NewASReqForTGT(realm string, c *config.Config, cname types.PrincipalName) (ASReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:87
	_go_fuzz_dep_.CoverTab[87884]++
												sname := types.PrincipalName{
		NameType:	nametype.KRB_NT_SRV_INST,
		NameString:	[]string{"krbtgt", realm},
	}
												return NewASReq(realm, c, cname, sname)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:92
	// _ = "end of CoverTab[87884]"
}

// NewASReqForChgPasswd generates a new KRB_AS_REQ struct for a change password request.
func NewASReqForChgPasswd(realm string, c *config.Config, cname types.PrincipalName) (ASReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:96
	_go_fuzz_dep_.CoverTab[87885]++
												sname := types.PrincipalName{
		NameType:	nametype.KRB_NT_PRINCIPAL,
		NameString:	[]string{"kadmin", "changepw"},
	}
												return NewASReq(realm, c, cname, sname)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:101
	// _ = "end of CoverTab[87885]"
}

// NewASReq generates a new KRB_AS_REQ struct for a given SNAME.
func NewASReq(realm string, c *config.Config, cname, sname types.PrincipalName) (ASReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:105
	_go_fuzz_dep_.CoverTab[87886]++
												nonce, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:107
		_go_fuzz_dep_.CoverTab[87893]++
													return ASReq{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:108
		// _ = "end of CoverTab[87893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:109
		_go_fuzz_dep_.CoverTab[87894]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:109
		// _ = "end of CoverTab[87894]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:109
	// _ = "end of CoverTab[87886]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:109
	_go_fuzz_dep_.CoverTab[87887]++
												t := time.Now().UTC()

												kopts := types.NewKrbFlags()
												copy(kopts.Bytes, c.LibDefaults.KDCDefaultOptions.Bytes)
												kopts.BitLength = c.LibDefaults.KDCDefaultOptions.BitLength
												a := ASReq{
		KDCReqFields{
			PVNO:		iana.PVNO,
			MsgType:	msgtype.KRB_AS_REQ,
			PAData:		types.PADataSequence{},
			ReqBody: KDCReqBody{
				KDCOptions:	kopts,
				Realm:		realm,
				CName:		cname,
				SName:		sname,
				Till:		t.Add(c.LibDefaults.TicketLifetime),
				Nonce:		int(nonce.Int64()),
				EType:		c.LibDefaults.DefaultTktEnctypeIDs,
			},
		},
	}
	if c.LibDefaults.Forwardable {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:131
		_go_fuzz_dep_.CoverTab[87895]++
													types.SetFlag(&a.ReqBody.KDCOptions, flags.Forwardable)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:132
		// _ = "end of CoverTab[87895]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:133
		_go_fuzz_dep_.CoverTab[87896]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:133
		// _ = "end of CoverTab[87896]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:133
	// _ = "end of CoverTab[87887]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:133
	_go_fuzz_dep_.CoverTab[87888]++
												if c.LibDefaults.Canonicalize {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:134
		_go_fuzz_dep_.CoverTab[87897]++
													types.SetFlag(&a.ReqBody.KDCOptions, flags.Canonicalize)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:135
		// _ = "end of CoverTab[87897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:136
		_go_fuzz_dep_.CoverTab[87898]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:136
		// _ = "end of CoverTab[87898]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:136
	// _ = "end of CoverTab[87888]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:136
	_go_fuzz_dep_.CoverTab[87889]++
												if c.LibDefaults.Proxiable {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:137
		_go_fuzz_dep_.CoverTab[87899]++
													types.SetFlag(&a.ReqBody.KDCOptions, flags.Proxiable)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:138
		// _ = "end of CoverTab[87899]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:139
		_go_fuzz_dep_.CoverTab[87900]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:139
		// _ = "end of CoverTab[87900]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:139
	// _ = "end of CoverTab[87889]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:139
	_go_fuzz_dep_.CoverTab[87890]++
												if c.LibDefaults.RenewLifetime != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:140
		_go_fuzz_dep_.CoverTab[87901]++
													types.SetFlag(&a.ReqBody.KDCOptions, flags.Renewable)
													a.ReqBody.RTime = t.Add(c.LibDefaults.RenewLifetime)
													a.ReqBody.RTime = t.Add(time.Duration(48) * time.Hour)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:143
		// _ = "end of CoverTab[87901]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:144
		_go_fuzz_dep_.CoverTab[87902]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:144
		// _ = "end of CoverTab[87902]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:144
	// _ = "end of CoverTab[87890]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:144
	_go_fuzz_dep_.CoverTab[87891]++
												if !c.LibDefaults.NoAddresses {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:145
		_go_fuzz_dep_.CoverTab[87903]++
													ha, err := types.LocalHostAddresses()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:147
			_go_fuzz_dep_.CoverTab[87905]++
														return a, fmt.Errorf("could not get local addresses: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:148
			// _ = "end of CoverTab[87905]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:149
			_go_fuzz_dep_.CoverTab[87906]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:149
			// _ = "end of CoverTab[87906]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:149
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:149
		// _ = "end of CoverTab[87903]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:149
		_go_fuzz_dep_.CoverTab[87904]++
													ha = append(ha, types.HostAddressesFromNetIPs(c.LibDefaults.ExtraAddresses)...)
													a.ReqBody.Addresses = ha
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:151
		// _ = "end of CoverTab[87904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:152
		_go_fuzz_dep_.CoverTab[87907]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:152
		// _ = "end of CoverTab[87907]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:152
	// _ = "end of CoverTab[87891]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:152
	_go_fuzz_dep_.CoverTab[87892]++
												return a, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:153
	// _ = "end of CoverTab[87892]"
}

// NewTGSReq generates a new KRB_TGS_REQ struct.
func NewTGSReq(cname types.PrincipalName, kdcRealm string, c *config.Config, tgt Ticket, sessionKey types.EncryptionKey, sname types.PrincipalName, renewal bool) (TGSReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:157
	_go_fuzz_dep_.CoverTab[87908]++
												a, err := tgsReq(cname, sname, kdcRealm, renewal, c)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:159
		_go_fuzz_dep_.CoverTab[87910]++
													return a, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:160
		// _ = "end of CoverTab[87910]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:161
		_go_fuzz_dep_.CoverTab[87911]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:161
		// _ = "end of CoverTab[87911]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:161
	// _ = "end of CoverTab[87908]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:161
	_go_fuzz_dep_.CoverTab[87909]++
												err = a.setPAData(tgt, sessionKey)
												return a, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:163
	// _ = "end of CoverTab[87909]"
}

// NewUser2UserTGSReq returns a TGS-REQ suitable for user-to-user authentication (https://tools.ietf.org/html/rfc4120#section-3.7)
func NewUser2UserTGSReq(cname types.PrincipalName, kdcRealm string, c *config.Config, clientTGT Ticket, sessionKey types.EncryptionKey, sname types.PrincipalName, renewal bool, verifyingTGT Ticket) (TGSReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:167
	_go_fuzz_dep_.CoverTab[87912]++
												a, err := tgsReq(cname, sname, kdcRealm, renewal, c)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:169
		_go_fuzz_dep_.CoverTab[87914]++
													return a, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:170
		// _ = "end of CoverTab[87914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:171
		_go_fuzz_dep_.CoverTab[87915]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:171
		// _ = "end of CoverTab[87915]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:171
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:171
	// _ = "end of CoverTab[87912]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:171
	_go_fuzz_dep_.CoverTab[87913]++
												a.ReqBody.AdditionalTickets = []Ticket{verifyingTGT}
												types.SetFlag(&a.ReqBody.KDCOptions, flags.EncTktInSkey)
												err = a.setPAData(clientTGT, sessionKey)
												return a, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:175
	// _ = "end of CoverTab[87913]"
}

// tgsReq populates the fields for a TGS_REQ
func tgsReq(cname, sname types.PrincipalName, kdcRealm string, renewal bool, c *config.Config) (TGSReq, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:179
	_go_fuzz_dep_.CoverTab[87916]++
												nonce, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:181
		_go_fuzz_dep_.CoverTab[87924]++
													return TGSReq{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:182
		// _ = "end of CoverTab[87924]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:183
		_go_fuzz_dep_.CoverTab[87925]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:183
		// _ = "end of CoverTab[87925]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:183
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:183
	// _ = "end of CoverTab[87916]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:183
	_go_fuzz_dep_.CoverTab[87917]++
												t := time.Now().UTC()
												k := KDCReqFields{
		PVNO:		iana.PVNO,
		MsgType:	msgtype.KRB_TGS_REQ,
		ReqBody: KDCReqBody{
			KDCOptions:	types.NewKrbFlags(),
			Realm:		kdcRealm,
			CName:		cname,
			SName:		sname,
			Till:		t.Add(c.LibDefaults.TicketLifetime),
			Nonce:		int(nonce.Int64()),
			EType:		c.LibDefaults.DefaultTGSEnctypeIDs,
		},
		Renewal:	renewal,
	}
	if c.LibDefaults.Forwardable {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:199
		_go_fuzz_dep_.CoverTab[87926]++
													types.SetFlag(&k.ReqBody.KDCOptions, flags.Forwardable)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:200
		// _ = "end of CoverTab[87926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:201
		_go_fuzz_dep_.CoverTab[87927]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:201
		// _ = "end of CoverTab[87927]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:201
	// _ = "end of CoverTab[87917]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:201
	_go_fuzz_dep_.CoverTab[87918]++
												if c.LibDefaults.Canonicalize {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:202
		_go_fuzz_dep_.CoverTab[87928]++
													types.SetFlag(&k.ReqBody.KDCOptions, flags.Canonicalize)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:203
		// _ = "end of CoverTab[87928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:204
		_go_fuzz_dep_.CoverTab[87929]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:204
		// _ = "end of CoverTab[87929]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:204
	// _ = "end of CoverTab[87918]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:204
	_go_fuzz_dep_.CoverTab[87919]++
												if c.LibDefaults.Proxiable {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:205
		_go_fuzz_dep_.CoverTab[87930]++
													types.SetFlag(&k.ReqBody.KDCOptions, flags.Proxiable)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:206
		// _ = "end of CoverTab[87930]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:207
		_go_fuzz_dep_.CoverTab[87931]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:207
		// _ = "end of CoverTab[87931]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:207
	// _ = "end of CoverTab[87919]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:207
	_go_fuzz_dep_.CoverTab[87920]++
												if c.LibDefaults.RenewLifetime > time.Duration(0) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:208
		_go_fuzz_dep_.CoverTab[87932]++
													types.SetFlag(&k.ReqBody.KDCOptions, flags.Renewable)
													k.ReqBody.RTime = t.Add(c.LibDefaults.RenewLifetime)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:210
		// _ = "end of CoverTab[87932]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:211
		_go_fuzz_dep_.CoverTab[87933]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:211
		// _ = "end of CoverTab[87933]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:211
	// _ = "end of CoverTab[87920]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:211
	_go_fuzz_dep_.CoverTab[87921]++
												if !c.LibDefaults.NoAddresses {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:212
		_go_fuzz_dep_.CoverTab[87934]++
													ha, err := types.LocalHostAddresses()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:214
			_go_fuzz_dep_.CoverTab[87936]++
														return TGSReq{}, fmt.Errorf("could not get local addresses: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:215
			// _ = "end of CoverTab[87936]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:216
			_go_fuzz_dep_.CoverTab[87937]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:216
			// _ = "end of CoverTab[87937]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:216
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:216
		// _ = "end of CoverTab[87934]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:216
		_go_fuzz_dep_.CoverTab[87935]++
													ha = append(ha, types.HostAddressesFromNetIPs(c.LibDefaults.ExtraAddresses)...)
													k.ReqBody.Addresses = ha
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:218
		// _ = "end of CoverTab[87935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:219
		_go_fuzz_dep_.CoverTab[87938]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:219
		// _ = "end of CoverTab[87938]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:219
	// _ = "end of CoverTab[87921]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:219
	_go_fuzz_dep_.CoverTab[87922]++
												if renewal {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:220
		_go_fuzz_dep_.CoverTab[87939]++
													types.SetFlag(&k.ReqBody.KDCOptions, flags.Renew)
													types.SetFlag(&k.ReqBody.KDCOptions, flags.Renewable)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:222
		// _ = "end of CoverTab[87939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:223
		_go_fuzz_dep_.CoverTab[87940]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:223
		// _ = "end of CoverTab[87940]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:223
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:223
	// _ = "end of CoverTab[87922]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:223
	_go_fuzz_dep_.CoverTab[87923]++
												return TGSReq{
		k,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:226
	// _ = "end of CoverTab[87923]"
}

func (k *TGSReq) setPAData(tgt Ticket, sessionKey types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:229
	_go_fuzz_dep_.CoverTab[87941]++

												b, err := k.ReqBody.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:232
		_go_fuzz_dep_.CoverTab[87948]++
													return krberror.Errorf(err, krberror.EncodingError, "error marshaling TGS_REQ body")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:233
		// _ = "end of CoverTab[87948]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:234
		_go_fuzz_dep_.CoverTab[87949]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:234
		// _ = "end of CoverTab[87949]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:234
	// _ = "end of CoverTab[87941]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:234
	_go_fuzz_dep_.CoverTab[87942]++
												etype, err := crypto.GetEtype(sessionKey.KeyType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:236
		_go_fuzz_dep_.CoverTab[87950]++
													return krberror.Errorf(err, krberror.EncryptingError, "error getting etype to encrypt authenticator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:237
		// _ = "end of CoverTab[87950]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:238
		_go_fuzz_dep_.CoverTab[87951]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:238
		// _ = "end of CoverTab[87951]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:238
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:238
	// _ = "end of CoverTab[87942]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:238
	_go_fuzz_dep_.CoverTab[87943]++
												cb, err := etype.GetChecksumHash(sessionKey.KeyValue, b, keyusage.TGS_REQ_PA_TGS_REQ_AP_REQ_AUTHENTICATOR_CHKSUM)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:240
		_go_fuzz_dep_.CoverTab[87952]++
													return krberror.Errorf(err, krberror.ChksumError, "error getting etype checksum hash")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:241
		// _ = "end of CoverTab[87952]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:242
		_go_fuzz_dep_.CoverTab[87953]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:242
		// _ = "end of CoverTab[87953]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:242
	// _ = "end of CoverTab[87943]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:242
	_go_fuzz_dep_.CoverTab[87944]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:246
	auth, err := types.NewAuthenticator(tgt.Realm, k.ReqBody.CName)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:247
		_go_fuzz_dep_.CoverTab[87954]++
													return krberror.Errorf(err, krberror.KRBMsgError, "error generating new authenticator")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:248
		// _ = "end of CoverTab[87954]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:249
		_go_fuzz_dep_.CoverTab[87955]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:249
		// _ = "end of CoverTab[87955]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:249
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:249
	// _ = "end of CoverTab[87944]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:249
	_go_fuzz_dep_.CoverTab[87945]++
												auth.Cksum = types.Checksum{
		CksumType:	etype.GetHashID(),
		Checksum:	cb,
	}

	apReq, err := NewAPReq(tgt, sessionKey, auth)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:256
		_go_fuzz_dep_.CoverTab[87956]++
													return krberror.Errorf(err, krberror.KRBMsgError, "error generating new AP_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:257
		// _ = "end of CoverTab[87956]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:258
		_go_fuzz_dep_.CoverTab[87957]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:258
		// _ = "end of CoverTab[87957]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:258
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:258
	// _ = "end of CoverTab[87945]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:258
	_go_fuzz_dep_.CoverTab[87946]++
												apb, err := apReq.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:260
		_go_fuzz_dep_.CoverTab[87958]++
													return krberror.Errorf(err, krberror.EncodingError, "error marshaling AP_REQ for pre-authentication data")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:261
		// _ = "end of CoverTab[87958]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:262
		_go_fuzz_dep_.CoverTab[87959]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:262
		// _ = "end of CoverTab[87959]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:262
	// _ = "end of CoverTab[87946]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:262
	_go_fuzz_dep_.CoverTab[87947]++
												k.PAData = types.PADataSequence{
		types.PAData{
			PADataType:	patype.PA_TGS_REQ,
			PADataValue:	apb,
		},
	}
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:269
	// _ = "end of CoverTab[87947]"
}

// Unmarshal bytes b into the ASReq struct.
func (k *ASReq) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:273
	_go_fuzz_dep_.CoverTab[87960]++
												var m marshalKDCReq
												_, err := asn1.UnmarshalWithParams(b, &m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.ASREQ))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:276
		_go_fuzz_dep_.CoverTab[87964]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:277
		// _ = "end of CoverTab[87964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:278
		_go_fuzz_dep_.CoverTab[87965]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:278
		// _ = "end of CoverTab[87965]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:278
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:278
	// _ = "end of CoverTab[87960]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:278
	_go_fuzz_dep_.CoverTab[87961]++
												expectedMsgType := msgtype.KRB_AS_REQ
												if m.MsgType != expectedMsgType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:280
		_go_fuzz_dep_.CoverTab[87966]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate a AS_REQ. Expected: %v; Actual: %v", expectedMsgType, m.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:281
		// _ = "end of CoverTab[87966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:282
		_go_fuzz_dep_.CoverTab[87967]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:282
		// _ = "end of CoverTab[87967]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:282
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:282
	// _ = "end of CoverTab[87961]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:282
	_go_fuzz_dep_.CoverTab[87962]++
												var reqb KDCReqBody
												err = reqb.Unmarshal(m.ReqBody.Bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:285
		_go_fuzz_dep_.CoverTab[87968]++
													return krberror.Errorf(err, krberror.EncodingError, "error processing AS_REQ body")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:286
		// _ = "end of CoverTab[87968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:287
		_go_fuzz_dep_.CoverTab[87969]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:287
		// _ = "end of CoverTab[87969]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:287
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:287
	// _ = "end of CoverTab[87962]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:287
	_go_fuzz_dep_.CoverTab[87963]++
												k.MsgType = m.MsgType
												k.PAData = m.PAData
												k.PVNO = m.PVNO
												k.ReqBody = reqb
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:292
	// _ = "end of CoverTab[87963]"
}

// Unmarshal bytes b into the TGSReq struct.
func (k *TGSReq) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:296
	_go_fuzz_dep_.CoverTab[87970]++
												var m marshalKDCReq
												_, err := asn1.UnmarshalWithParams(b, &m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.TGSREQ))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:299
		_go_fuzz_dep_.CoverTab[87974]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling TGS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:300
		// _ = "end of CoverTab[87974]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:301
		_go_fuzz_dep_.CoverTab[87975]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:301
		// _ = "end of CoverTab[87975]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:301
	// _ = "end of CoverTab[87970]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:301
	_go_fuzz_dep_.CoverTab[87971]++
												expectedMsgType := msgtype.KRB_TGS_REQ
												if m.MsgType != expectedMsgType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:303
		_go_fuzz_dep_.CoverTab[87976]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate a TGS_REQ. Expected: %v; Actual: %v", expectedMsgType, m.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:304
		// _ = "end of CoverTab[87976]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:305
		_go_fuzz_dep_.CoverTab[87977]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:305
		// _ = "end of CoverTab[87977]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:305
	// _ = "end of CoverTab[87971]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:305
	_go_fuzz_dep_.CoverTab[87972]++
												var reqb KDCReqBody
												err = reqb.Unmarshal(m.ReqBody.Bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:308
		_go_fuzz_dep_.CoverTab[87978]++
													return krberror.Errorf(err, krberror.EncodingError, "error processing TGS_REQ body")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:309
		// _ = "end of CoverTab[87978]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:310
		_go_fuzz_dep_.CoverTab[87979]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:310
		// _ = "end of CoverTab[87979]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:310
	// _ = "end of CoverTab[87972]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:310
	_go_fuzz_dep_.CoverTab[87973]++
												k.MsgType = m.MsgType
												k.PAData = m.PAData
												k.PVNO = m.PVNO
												k.ReqBody = reqb
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:315
	// _ = "end of CoverTab[87973]"
}

// Unmarshal bytes b into the KRB_KDC_REQ body struct.
func (k *KDCReqBody) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:319
	_go_fuzz_dep_.CoverTab[87980]++
												var m marshalKDCReqBody
												_, err := asn1.Unmarshal(b, &m)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:322
		_go_fuzz_dep_.CoverTab[87984]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling KDC_REQ body")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:323
		// _ = "end of CoverTab[87984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:324
		_go_fuzz_dep_.CoverTab[87985]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:324
		// _ = "end of CoverTab[87985]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:324
	// _ = "end of CoverTab[87980]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:324
	_go_fuzz_dep_.CoverTab[87981]++
												k.KDCOptions = m.KDCOptions
												if len(k.KDCOptions.Bytes) < 4 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:326
		_go_fuzz_dep_.CoverTab[87986]++
													tb := make([]byte, 4-len(k.KDCOptions.Bytes))
													k.KDCOptions.Bytes = append(tb, k.KDCOptions.Bytes...)
													k.KDCOptions.BitLength = len(k.KDCOptions.Bytes) * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:329
		// _ = "end of CoverTab[87986]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:330
		_go_fuzz_dep_.CoverTab[87987]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:330
		// _ = "end of CoverTab[87987]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:330
	// _ = "end of CoverTab[87981]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:330
	_go_fuzz_dep_.CoverTab[87982]++
												k.CName = m.CName
												k.Realm = m.Realm
												k.SName = m.SName
												k.From = m.From
												k.Till = m.Till
												k.RTime = m.RTime
												k.Nonce = m.Nonce
												k.EType = m.EType
												k.Addresses = m.Addresses
												k.EncAuthData = m.EncAuthData
												if len(m.AdditionalTickets.Bytes) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:341
		_go_fuzz_dep_.CoverTab[87988]++
													k.AdditionalTickets, err = unmarshalTicketsSequence(m.AdditionalTickets)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:343
			_go_fuzz_dep_.CoverTab[87989]++
														return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling additional tickets")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:344
			// _ = "end of CoverTab[87989]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:345
			_go_fuzz_dep_.CoverTab[87990]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:345
			// _ = "end of CoverTab[87990]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:345
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:345
		// _ = "end of CoverTab[87988]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:346
		_go_fuzz_dep_.CoverTab[87991]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:346
		// _ = "end of CoverTab[87991]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:346
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:346
	// _ = "end of CoverTab[87982]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:346
	_go_fuzz_dep_.CoverTab[87983]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:347
	// _ = "end of CoverTab[87983]"
}

// Marshal ASReq struct.
func (k *ASReq) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:351
	_go_fuzz_dep_.CoverTab[87992]++
												m := marshalKDCReq{
		PVNO:		k.PVNO,
		MsgType:	k.MsgType,
		PAData:		k.PAData,
	}
	b, err := k.ReqBody.Marshal()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:358
		_go_fuzz_dep_.CoverTab[87995]++
													var mk []byte
													return mk, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:360
		// _ = "end of CoverTab[87995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:361
		_go_fuzz_dep_.CoverTab[87996]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:361
		// _ = "end of CoverTab[87996]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:361
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:361
	// _ = "end of CoverTab[87992]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:361
	_go_fuzz_dep_.CoverTab[87993]++
												m.ReqBody = asn1.RawValue{
		Class:		asn1.ClassContextSpecific,
		IsCompound:	true,
		Tag:		4,
		Bytes:		b,
	}
	mk, err := asn1.Marshal(m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:369
		_go_fuzz_dep_.CoverTab[87997]++
													return mk, krberror.Errorf(err, krberror.EncodingError, "error marshaling AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:370
		// _ = "end of CoverTab[87997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:371
		_go_fuzz_dep_.CoverTab[87998]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:371
		// _ = "end of CoverTab[87998]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:371
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:371
	// _ = "end of CoverTab[87993]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:371
	_go_fuzz_dep_.CoverTab[87994]++
												mk = asn1tools.AddASNAppTag(mk, asnAppTag.ASREQ)
												return mk, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:373
	// _ = "end of CoverTab[87994]"
}

// Marshal TGSReq struct.
func (k *TGSReq) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:377
	_go_fuzz_dep_.CoverTab[87999]++
												m := marshalKDCReq{
		PVNO:		k.PVNO,
		MsgType:	k.MsgType,
		PAData:		k.PAData,
	}
	b, err := k.ReqBody.Marshal()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:384
		_go_fuzz_dep_.CoverTab[88002]++
													var mk []byte
													return mk, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:386
		// _ = "end of CoverTab[88002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:387
		_go_fuzz_dep_.CoverTab[88003]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:387
		// _ = "end of CoverTab[88003]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:387
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:387
	// _ = "end of CoverTab[87999]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:387
	_go_fuzz_dep_.CoverTab[88000]++
												m.ReqBody = asn1.RawValue{
		Class:		asn1.ClassContextSpecific,
		IsCompound:	true,
		Tag:		4,
		Bytes:		b,
	}
	mk, err := asn1.Marshal(m)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:395
		_go_fuzz_dep_.CoverTab[88004]++
													return mk, krberror.Errorf(err, krberror.EncodingError, "error marshaling AS_REQ")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:396
		// _ = "end of CoverTab[88004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:397
		_go_fuzz_dep_.CoverTab[88005]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:397
		// _ = "end of CoverTab[88005]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:397
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:397
	// _ = "end of CoverTab[88000]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:397
	_go_fuzz_dep_.CoverTab[88001]++
												mk = asn1tools.AddASNAppTag(mk, asnAppTag.TGSREQ)
												return mk, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:399
	// _ = "end of CoverTab[88001]"
}

// Marshal KRB_KDC_REQ body struct.
func (k *KDCReqBody) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:403
	_go_fuzz_dep_.CoverTab[88006]++
												var b []byte
												m := marshalKDCReqBody{
		KDCOptions:	k.KDCOptions,
		CName:		k.CName,
		Realm:		k.Realm,
		SName:		k.SName,
		From:		k.From,
		Till:		k.Till,
		RTime:		k.RTime,
		Nonce:		k.Nonce,
		EType:		k.EType,
		Addresses:	k.Addresses,
		EncAuthData:	k.EncAuthData,
	}
	rawtkts, err := MarshalTicketSequence(k.AdditionalTickets)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:419
		_go_fuzz_dep_.CoverTab[88010]++
													return b, krberror.Errorf(err, krberror.EncodingError, "error in marshaling KDC request body additional tickets")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:420
		// _ = "end of CoverTab[88010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:421
		_go_fuzz_dep_.CoverTab[88011]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:421
		// _ = "end of CoverTab[88011]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:421
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:421
	// _ = "end of CoverTab[88006]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:421
	_go_fuzz_dep_.CoverTab[88007]++

												rawtkts.Tag = 11
												if len(rawtkts.Bytes) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:424
		_go_fuzz_dep_.CoverTab[88012]++
													m.AdditionalTickets = rawtkts
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:425
		// _ = "end of CoverTab[88012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:426
		_go_fuzz_dep_.CoverTab[88013]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:426
		// _ = "end of CoverTab[88013]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:426
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:426
	// _ = "end of CoverTab[88007]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:426
	_go_fuzz_dep_.CoverTab[88008]++
												b, err = asn1.Marshal(m)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:428
		_go_fuzz_dep_.CoverTab[88014]++
													return b, krberror.Errorf(err, krberror.EncodingError, "error in marshaling KDC request body")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:429
		// _ = "end of CoverTab[88014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:430
		_go_fuzz_dep_.CoverTab[88015]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:430
		// _ = "end of CoverTab[88015]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:430
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:430
	// _ = "end of CoverTab[88008]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:430
	_go_fuzz_dep_.CoverTab[88009]++
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:431
	// _ = "end of CoverTab[88009]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:432
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KDCReq.go:432
var _ = _go_fuzz_dep_.CoverTab
