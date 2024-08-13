//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:1
)

import (
	"fmt"
	"log"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana"
	"github.com/jcmturner/gokrb5/v8/iana/adtype"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/errorcode"
	"github.com/jcmturner/gokrb5/v8/iana/flags"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/pac"
	"github.com/jcmturner/gokrb5/v8/types"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:26
// Ticket implements the Kerberos ticket.
type Ticket struct {
	TktVNO			int			`asn1:"explicit,tag:0"`
	Realm			string			`asn1:"generalstring,explicit,tag:1"`
	SName			types.PrincipalName	`asn1:"explicit,tag:2"`
	EncPart			types.EncryptedData	`asn1:"explicit,tag:3"`
	DecryptedEncPart	EncTicketPart		`asn1:"optional"`	// Not part of ASN1 bytes so marked as optional so unmarshalling works
}

// EncTicketPart is the encrypted part of the Ticket.
type EncTicketPart struct {
	Flags			asn1.BitString		`asn1:"explicit,tag:0"`
	Key			types.EncryptionKey	`asn1:"explicit,tag:1"`
	CRealm			string			`asn1:"generalstring,explicit,tag:2"`
	CName			types.PrincipalName	`asn1:"explicit,tag:3"`
	Transited		TransitedEncoding	`asn1:"explicit,tag:4"`
	AuthTime		time.Time		`asn1:"generalized,explicit,tag:5"`
	StartTime		time.Time		`asn1:"generalized,explicit,optional,tag:6"`
	EndTime			time.Time		`asn1:"generalized,explicit,tag:7"`
	RenewTill		time.Time		`asn1:"generalized,explicit,optional,tag:8"`
	CAddr			types.HostAddresses	`asn1:"explicit,optional,tag:9"`
	AuthorizationData	types.AuthorizationData	`asn1:"explicit,optional,tag:10"`
}

// TransitedEncoding part of the ticket's encrypted part.
type TransitedEncoding struct {
	TRType		int32	`asn1:"explicit,tag:0"`
	Contents	[]byte	`asn1:"explicit,tag:1"`
}

// NewTicket creates a new Ticket instance.
func NewTicket(cname types.PrincipalName, crealm string, sname types.PrincipalName, srealm string, flags asn1.BitString, sktab *keytab.Keytab, eTypeID int32, kvno int, authTime, startTime, endTime, renewTill time.Time) (Ticket, types.EncryptionKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:57
	_go_fuzz_dep_.CoverTab[88098]++
												etype, err := crypto.GetEtype(eTypeID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:59
		_go_fuzz_dep_.CoverTab[88104]++
													return Ticket{}, types.EncryptionKey{}, krberror.Errorf(err, krberror.EncryptingError, "error getting etype for new ticket")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:60
		// _ = "end of CoverTab[88104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:61
		_go_fuzz_dep_.CoverTab[88105]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:61
		// _ = "end of CoverTab[88105]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:61
	// _ = "end of CoverTab[88098]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:61
	_go_fuzz_dep_.CoverTab[88099]++
												sessionKey, err := types.GenerateEncryptionKey(etype)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:63
		_go_fuzz_dep_.CoverTab[88106]++
													return Ticket{}, types.EncryptionKey{}, krberror.Errorf(err, krberror.EncryptingError, "error generating session key")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:64
		// _ = "end of CoverTab[88106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:65
		_go_fuzz_dep_.CoverTab[88107]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:65
		// _ = "end of CoverTab[88107]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:65
	// _ = "end of CoverTab[88099]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:65
	_go_fuzz_dep_.CoverTab[88100]++

												etp := EncTicketPart{
		Flags:		flags,
		Key:		sessionKey,
		CRealm:		crealm,
		CName:		cname,
		Transited:	TransitedEncoding{},
		AuthTime:	authTime,
		StartTime:	startTime,
		EndTime:	endTime,
		RenewTill:	renewTill,
	}
	b, err := asn1.Marshal(etp)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:79
		_go_fuzz_dep_.CoverTab[88108]++
													return Ticket{}, types.EncryptionKey{}, krberror.Errorf(err, krberror.EncodingError, "error marshalling ticket encpart")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:80
		// _ = "end of CoverTab[88108]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:81
		_go_fuzz_dep_.CoverTab[88109]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:81
		// _ = "end of CoverTab[88109]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:81
	// _ = "end of CoverTab[88100]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:81
	_go_fuzz_dep_.CoverTab[88101]++
												b = asn1tools.AddASNAppTag(b, asnAppTag.EncTicketPart)
												skey, _, err := sktab.GetEncryptionKey(sname, srealm, kvno, eTypeID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:84
		_go_fuzz_dep_.CoverTab[88110]++
													return Ticket{}, types.EncryptionKey{}, krberror.Errorf(err, krberror.EncryptingError, "error getting encryption key for new ticket")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:85
		// _ = "end of CoverTab[88110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:86
		_go_fuzz_dep_.CoverTab[88111]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:86
		// _ = "end of CoverTab[88111]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:86
	// _ = "end of CoverTab[88101]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:86
	_go_fuzz_dep_.CoverTab[88102]++
												ed, err := crypto.GetEncryptedData(b, skey, keyusage.KDC_REP_TICKET, kvno)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:88
		_go_fuzz_dep_.CoverTab[88112]++
													return Ticket{}, types.EncryptionKey{}, krberror.Errorf(err, krberror.EncryptingError, "error encrypting ticket encpart")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:89
		// _ = "end of CoverTab[88112]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:90
		_go_fuzz_dep_.CoverTab[88113]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:90
		// _ = "end of CoverTab[88113]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:90
	// _ = "end of CoverTab[88102]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:90
	_go_fuzz_dep_.CoverTab[88103]++
												tkt := Ticket{
		TktVNO:		iana.PVNO,
		Realm:		srealm,
		SName:		sname,
		EncPart:	ed,
	}
												return tkt, sessionKey, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:97
	// _ = "end of CoverTab[88103]"
}

// Unmarshal bytes b into a Ticket struct.
func (t *Ticket) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:101
	_go_fuzz_dep_.CoverTab[88114]++
												_, err := asn1.UnmarshalWithParams(b, t, fmt.Sprintf("application,explicit,tag:%d", asnAppTag.Ticket))
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:103
	// _ = "end of CoverTab[88114]"
}

// Marshal the Ticket.
func (t *Ticket) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:107
	_go_fuzz_dep_.CoverTab[88115]++
												b, err := asn1.Marshal(*t)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:109
		_go_fuzz_dep_.CoverTab[88117]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:110
		// _ = "end of CoverTab[88117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:111
		_go_fuzz_dep_.CoverTab[88118]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:111
		// _ = "end of CoverTab[88118]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:111
	// _ = "end of CoverTab[88115]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:111
	_go_fuzz_dep_.CoverTab[88116]++
												b = asn1tools.AddASNAppTag(b, asnAppTag.Ticket)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:113
	// _ = "end of CoverTab[88116]"
}

// Unmarshal bytes b into the EncTicketPart struct.
func (t *EncTicketPart) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:117
	_go_fuzz_dep_.CoverTab[88119]++
												_, err := asn1.UnmarshalWithParams(b, t, fmt.Sprintf("application,explicit,tag:%d", asnAppTag.EncTicketPart))
												return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:119
	// _ = "end of CoverTab[88119]"
}

// unmarshalTicket returns a ticket from the bytes provided.
func unmarshalTicket(b []byte) (t Ticket, err error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:123
	_go_fuzz_dep_.CoverTab[88120]++
												err = t.Unmarshal(b)
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:125
	// _ = "end of CoverTab[88120]"
}

// UnmarshalTicketsSequence returns a slice of Tickets from a raw ASN1 value.
func unmarshalTicketsSequence(in asn1.RawValue) ([]Ticket, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:129
	_go_fuzz_dep_.CoverTab[88121]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:132
	b := in.Bytes

	p := 1 + asn1tools.GetNumberBytesInLengthHeader(in.Bytes)
	var tkts []Ticket
	var raw asn1.RawValue
	for p < (len(b)) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:137
		_go_fuzz_dep_.CoverTab[88123]++
													_, err := asn1.UnmarshalWithParams(b[p:], &raw, fmt.Sprintf("application,tag:%d", asnAppTag.Ticket))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:139
			_go_fuzz_dep_.CoverTab[88126]++
														return nil, fmt.Errorf("unmarshaling sequence of tickets failed getting length of ticket: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:140
			// _ = "end of CoverTab[88126]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:141
			_go_fuzz_dep_.CoverTab[88127]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:141
			// _ = "end of CoverTab[88127]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:141
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:141
		// _ = "end of CoverTab[88123]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:141
		_go_fuzz_dep_.CoverTab[88124]++
													t, err := unmarshalTicket(b[p:])
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:143
			_go_fuzz_dep_.CoverTab[88128]++
														return nil, fmt.Errorf("unmarshaling sequence of tickets failed: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:144
			// _ = "end of CoverTab[88128]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:145
			_go_fuzz_dep_.CoverTab[88129]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:145
			// _ = "end of CoverTab[88129]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:145
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:145
		// _ = "end of CoverTab[88124]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:145
		_go_fuzz_dep_.CoverTab[88125]++
													p += len(raw.FullBytes)
													tkts = append(tkts, t)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:147
		// _ = "end of CoverTab[88125]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:148
	// _ = "end of CoverTab[88121]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:148
	_go_fuzz_dep_.CoverTab[88122]++
												MarshalTicketSequence(tkts)
												return tkts, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:150
	// _ = "end of CoverTab[88122]"
}

// MarshalTicketSequence marshals a slice of Tickets returning an ASN1 raw value containing the ticket sequence.
func MarshalTicketSequence(tkts []Ticket) (asn1.RawValue, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:154
	_go_fuzz_dep_.CoverTab[88130]++
												raw := asn1.RawValue{
		Class:		2,
		IsCompound:	true,
	}
	if len(tkts) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:159
		_go_fuzz_dep_.CoverTab[88133]++

													return raw, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:161
		// _ = "end of CoverTab[88133]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:162
		_go_fuzz_dep_.CoverTab[88134]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:162
		// _ = "end of CoverTab[88134]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:162
	// _ = "end of CoverTab[88130]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:162
	_go_fuzz_dep_.CoverTab[88131]++
												var btkts []byte
												for i, t := range tkts {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:164
		_go_fuzz_dep_.CoverTab[88135]++
													b, err := t.Marshal()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:166
			_go_fuzz_dep_.CoverTab[88137]++
														return raw, fmt.Errorf("error marshaling ticket number %d in sequence of tickets", i+1)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:167
			// _ = "end of CoverTab[88137]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:168
			_go_fuzz_dep_.CoverTab[88138]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:168
			// _ = "end of CoverTab[88138]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:168
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:168
		// _ = "end of CoverTab[88135]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:168
		_go_fuzz_dep_.CoverTab[88136]++
													btkts = append(btkts, b...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:169
		// _ = "end of CoverTab[88136]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:170
	// _ = "end of CoverTab[88131]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:170
	_go_fuzz_dep_.CoverTab[88132]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:178
	btkts = append(asn1tools.MarshalLengthBytes(len(btkts)), btkts...)
												btkts = append([]byte{byte(32 + asn1.TagSequence)}, btkts...)
												raw.Bytes = btkts

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:183
	return raw, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:183
	// _ = "end of CoverTab[88132]"
}

// DecryptEncPart decrypts the encrypted part of the ticket.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:186
// The sname argument can be used to specify which service principal's key should be used to decrypt the ticket.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:186
// If nil is passed as the sname then the service principal specified within the ticket it used.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:189
func (t *Ticket) DecryptEncPart(keytab *keytab.Keytab, sname *types.PrincipalName) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:189
	_go_fuzz_dep_.CoverTab[88139]++
												if sname == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:190
		_go_fuzz_dep_.CoverTab[88142]++
													sname = &t.SName
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:191
		// _ = "end of CoverTab[88142]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:192
		_go_fuzz_dep_.CoverTab[88143]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:192
		// _ = "end of CoverTab[88143]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:192
	// _ = "end of CoverTab[88139]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:192
	_go_fuzz_dep_.CoverTab[88140]++
												key, _, err := keytab.GetEncryptionKey(*sname, t.Realm, t.EncPart.KVNO, t.EncPart.EType)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:194
		_go_fuzz_dep_.CoverTab[88144]++
													return NewKRBError(t.SName, t.Realm, errorcode.KRB_AP_ERR_NOKEY, fmt.Sprintf("Could not get key from keytab: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:195
		// _ = "end of CoverTab[88144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:196
		_go_fuzz_dep_.CoverTab[88145]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:196
		// _ = "end of CoverTab[88145]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:196
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:196
	// _ = "end of CoverTab[88140]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:196
	_go_fuzz_dep_.CoverTab[88141]++
												return t.Decrypt(key)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:197
	// _ = "end of CoverTab[88141]"
}

// Decrypt decrypts the encrypted part of the ticket using the key provided.
func (t *Ticket) Decrypt(key types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:201
	_go_fuzz_dep_.CoverTab[88146]++
												b, err := crypto.DecryptEncPart(t.EncPart, key, keyusage.KDC_REP_TICKET)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:203
		_go_fuzz_dep_.CoverTab[88149]++
													return fmt.Errorf("error decrypting Ticket EncPart: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:204
		// _ = "end of CoverTab[88149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:205
		_go_fuzz_dep_.CoverTab[88150]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:205
		// _ = "end of CoverTab[88150]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:205
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:205
	// _ = "end of CoverTab[88146]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:205
	_go_fuzz_dep_.CoverTab[88147]++
												var denc EncTicketPart
												err = denc.Unmarshal(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:208
		_go_fuzz_dep_.CoverTab[88151]++
													return fmt.Errorf("error unmarshaling encrypted part: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:209
		// _ = "end of CoverTab[88151]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:210
		_go_fuzz_dep_.CoverTab[88152]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:210
		// _ = "end of CoverTab[88152]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:210
	// _ = "end of CoverTab[88147]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:210
	_go_fuzz_dep_.CoverTab[88148]++
												t.DecryptedEncPart = denc
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:212
	// _ = "end of CoverTab[88148]"
}

// GetPACType returns a Microsoft PAC that has been extracted from the ticket and processed.
func (t *Ticket) GetPACType(keytab *keytab.Keytab, sname *types.PrincipalName, l *log.Logger) (bool, pac.PACType, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:216
	_go_fuzz_dep_.CoverTab[88153]++
												var isPAC bool
												for _, ad := range t.DecryptedEncPart.AuthorizationData {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:218
		_go_fuzz_dep_.CoverTab[88155]++
													if ad.ADType == adtype.ADIfRelevant {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:219
			_go_fuzz_dep_.CoverTab[88156]++
														var ad2 types.AuthorizationData
														err := ad2.Unmarshal(ad.ADData)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:222
				_go_fuzz_dep_.CoverTab[88158]++
															l.Printf("PAC authorization data could not be unmarshaled: %v", err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:224
				// _ = "end of CoverTab[88158]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:225
				_go_fuzz_dep_.CoverTab[88159]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:225
				// _ = "end of CoverTab[88159]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:225
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:225
			// _ = "end of CoverTab[88156]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:225
			_go_fuzz_dep_.CoverTab[88157]++
														if ad2[0].ADType == adtype.ADWin2KPAC {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:226
				_go_fuzz_dep_.CoverTab[88160]++
															isPAC = true
															var p pac.PACType
															err = p.Unmarshal(ad2[0].ADData)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:230
					_go_fuzz_dep_.CoverTab[88164]++
																return isPAC, p, fmt.Errorf("error unmarshaling PAC: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:231
					// _ = "end of CoverTab[88164]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:232
					_go_fuzz_dep_.CoverTab[88165]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:232
					// _ = "end of CoverTab[88165]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:232
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:232
				// _ = "end of CoverTab[88160]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:232
				_go_fuzz_dep_.CoverTab[88161]++
															if sname == nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:233
					_go_fuzz_dep_.CoverTab[88166]++
																sname = &t.SName
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:234
					// _ = "end of CoverTab[88166]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:235
					_go_fuzz_dep_.CoverTab[88167]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:235
					// _ = "end of CoverTab[88167]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:235
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:235
				// _ = "end of CoverTab[88161]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:235
				_go_fuzz_dep_.CoverTab[88162]++
															key, _, err := keytab.GetEncryptionKey(*sname, t.Realm, t.EncPart.KVNO, t.EncPart.EType)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:237
					_go_fuzz_dep_.CoverTab[88168]++
																return isPAC, p, NewKRBError(t.SName, t.Realm, errorcode.KRB_AP_ERR_NOKEY, fmt.Sprintf("Could not get key from keytab: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:238
					// _ = "end of CoverTab[88168]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:239
					_go_fuzz_dep_.CoverTab[88169]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:239
					// _ = "end of CoverTab[88169]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:239
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:239
				// _ = "end of CoverTab[88162]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:239
				_go_fuzz_dep_.CoverTab[88163]++
															err = p.ProcessPACInfoBuffers(key, l)
															return isPAC, p, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:241
				// _ = "end of CoverTab[88163]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:242
				_go_fuzz_dep_.CoverTab[88170]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:242
				// _ = "end of CoverTab[88170]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:242
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:242
			// _ = "end of CoverTab[88157]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:243
			_go_fuzz_dep_.CoverTab[88171]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:243
			// _ = "end of CoverTab[88171]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:243
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:243
		// _ = "end of CoverTab[88155]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:244
	// _ = "end of CoverTab[88153]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:244
	_go_fuzz_dep_.CoverTab[88154]++
												return isPAC, pac.PACType{}, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:245
	// _ = "end of CoverTab[88154]"
}

// Valid checks it the ticket is currently valid. Max duration passed endtime passed in as argument.
func (t *Ticket) Valid(d time.Duration) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:249
	_go_fuzz_dep_.CoverTab[88172]++

												time := time.Now().UTC()
												if t.DecryptedEncPart.StartTime.Sub(time) > d || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:252
		_go_fuzz_dep_.CoverTab[88175]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:252
		return types.IsFlagSet(&t.DecryptedEncPart.Flags, flags.Invalid)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:252
		// _ = "end of CoverTab[88175]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:252
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:252
		_go_fuzz_dep_.CoverTab[88176]++
													return false, NewKRBError(t.SName, t.Realm, errorcode.KRB_AP_ERR_TKT_NYV, "service ticket provided is not yet valid")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:253
		// _ = "end of CoverTab[88176]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:254
		_go_fuzz_dep_.CoverTab[88177]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:254
		// _ = "end of CoverTab[88177]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:254
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:254
	// _ = "end of CoverTab[88172]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:254
	_go_fuzz_dep_.CoverTab[88173]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:257
	if time.Sub(t.DecryptedEncPart.EndTime) > d {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:257
		_go_fuzz_dep_.CoverTab[88178]++
													return false, NewKRBError(t.SName, t.Realm, errorcode.KRB_AP_ERR_TKT_EXPIRED, "service ticket provided has expired")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:258
		// _ = "end of CoverTab[88178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:259
		_go_fuzz_dep_.CoverTab[88179]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:259
		// _ = "end of CoverTab[88179]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:259
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:259
	// _ = "end of CoverTab[88173]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:259
	_go_fuzz_dep_.CoverTab[88174]++

												return true, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:261
	// _ = "end of CoverTab[88174]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:262
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/Ticket.go:262
var _ = _go_fuzz_dep_.CoverTab
