//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:1
)

import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

type marshalKRBCred struct {
	PVNO	int			`asn1:"explicit,tag:0"`
	MsgType	int			`asn1:"explicit,tag:1"`
	Tickets	asn1.RawValue		`asn1:"explicit,tag:2"`
	EncPart	types.EncryptedData	`asn1:"explicit,tag:3"`
}

// KRBCred implements RFC 4120 KRB_CRED: https://tools.ietf.org/html/rfc4120#section-5.8.1.
type KRBCred struct {
	PVNO			int
	MsgType			int
	Tickets			[]Ticket
	EncPart			types.EncryptedData
	DecryptedEncPart	EncKrbCredPart
}

// EncKrbCredPart is the encrypted part of KRB_CRED.
type EncKrbCredPart struct {
	TicketInfo	[]KrbCredInfo		`asn1:"explicit,tag:0"`
	Nouce		int			`asn1:"optional,explicit,tag:1"`
	Timestamp	time.Time		`asn1:"generalized,optional,explicit,tag:2"`
	Usec		int			`asn1:"optional,explicit,tag:3"`
	SAddress	types.HostAddress	`asn1:"optional,explicit,tag:4"`
	RAddress	types.HostAddress	`asn1:"optional,explicit,tag:5"`
}

// KrbCredInfo is the KRB_CRED_INFO part of KRB_CRED.
type KrbCredInfo struct {
	Key		types.EncryptionKey	`asn1:"explicit,tag:0"`
	PRealm		string			`asn1:"generalstring,optional,explicit,tag:1"`
	PName		types.PrincipalName	`asn1:"optional,explicit,tag:2"`
	Flags		asn1.BitString		`asn1:"optional,explicit,tag:3"`
	AuthTime	time.Time		`asn1:"generalized,optional,explicit,tag:4"`
	StartTime	time.Time		`asn1:"generalized,optional,explicit,tag:5"`
	EndTime		time.Time		`asn1:"generalized,optional,explicit,tag:6"`
	RenewTill	time.Time		`asn1:"generalized,optional,explicit,tag:7"`
	SRealm		string			`asn1:"optional,explicit,ia5,tag:8"`
	SName		types.PrincipalName	`asn1:"optional,explicit,tag:9"`
	CAddr		types.HostAddresses	`asn1:"optional,explicit,tag:10"`
}

// Unmarshal bytes b into the KRBCred struct.
func (k *KRBCred) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:58
	_go_fuzz_dep_.CoverTab[88016]++
												var m marshalKRBCred
												_, err := asn1.UnmarshalWithParams(b, &m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.KRBCred))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:61
		_go_fuzz_dep_.CoverTab[88020]++
													return processUnmarshalReplyError(b, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:62
		// _ = "end of CoverTab[88020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:63
		_go_fuzz_dep_.CoverTab[88021]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:63
		// _ = "end of CoverTab[88021]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:63
	// _ = "end of CoverTab[88016]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:63
	_go_fuzz_dep_.CoverTab[88017]++
												expectedMsgType := msgtype.KRB_CRED
												if m.MsgType != expectedMsgType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:65
		_go_fuzz_dep_.CoverTab[88022]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate a KRB_CRED. Expected: %v; Actual: %v", expectedMsgType, m.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:66
		// _ = "end of CoverTab[88022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:67
		_go_fuzz_dep_.CoverTab[88023]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:67
		// _ = "end of CoverTab[88023]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:67
	// _ = "end of CoverTab[88017]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:67
	_go_fuzz_dep_.CoverTab[88018]++
												k.PVNO = m.PVNO
												k.MsgType = m.MsgType
												k.EncPart = m.EncPart
												if len(m.Tickets.Bytes) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:71
		_go_fuzz_dep_.CoverTab[88024]++
													k.Tickets, err = unmarshalTicketsSequence(m.Tickets)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:73
			_go_fuzz_dep_.CoverTab[88025]++
														return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling tickets within KRB_CRED")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:74
			// _ = "end of CoverTab[88025]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:75
			_go_fuzz_dep_.CoverTab[88026]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:75
			// _ = "end of CoverTab[88026]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:75
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:75
		// _ = "end of CoverTab[88024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:76
		_go_fuzz_dep_.CoverTab[88027]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:76
		// _ = "end of CoverTab[88027]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:76
	// _ = "end of CoverTab[88018]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:76
	_go_fuzz_dep_.CoverTab[88019]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:77
	// _ = "end of CoverTab[88019]"
}

// DecryptEncPart decrypts the encrypted part of a KRB_CRED.
func (k *KRBCred) DecryptEncPart(key types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:81
	_go_fuzz_dep_.CoverTab[88028]++
												b, err := crypto.DecryptEncPart(k.EncPart, key, keyusage.KRB_CRED_ENCPART)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:83
		_go_fuzz_dep_.CoverTab[88031]++
													return krberror.Errorf(err, krberror.DecryptingError, "error decrypting KRB_CRED EncPart")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:84
		// _ = "end of CoverTab[88031]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:85
		_go_fuzz_dep_.CoverTab[88032]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:85
		// _ = "end of CoverTab[88032]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:85
	// _ = "end of CoverTab[88028]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:85
	_go_fuzz_dep_.CoverTab[88029]++
												var denc EncKrbCredPart
												err = denc.Unmarshal(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:88
		_go_fuzz_dep_.CoverTab[88033]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling encrypted part of KRB_CRED")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:89
		// _ = "end of CoverTab[88033]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:90
		_go_fuzz_dep_.CoverTab[88034]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:90
		// _ = "end of CoverTab[88034]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:90
	// _ = "end of CoverTab[88029]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:90
	_go_fuzz_dep_.CoverTab[88030]++
												k.DecryptedEncPart = denc
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:92
	// _ = "end of CoverTab[88030]"
}

// Unmarshal bytes b into the encrypted part of KRB_CRED.
func (k *EncKrbCredPart) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:96
	_go_fuzz_dep_.CoverTab[88035]++
												_, err := asn1.UnmarshalWithParams(b, k, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.EncKrbCredPart))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:98
		_go_fuzz_dep_.CoverTab[88037]++
													return krberror.Errorf(err, krberror.EncodingError, "error unmarshaling EncKrbCredPart")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:99
		// _ = "end of CoverTab[88037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:100
		_go_fuzz_dep_.CoverTab[88038]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:100
		// _ = "end of CoverTab[88038]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:100
	// _ = "end of CoverTab[88035]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:100
	_go_fuzz_dep_.CoverTab[88036]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:101
	// _ = "end of CoverTab[88036]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:102
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBCred.go:102
var _ = _go_fuzz_dep_.CoverTab
