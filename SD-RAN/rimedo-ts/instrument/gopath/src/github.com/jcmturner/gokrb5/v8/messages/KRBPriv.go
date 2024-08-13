//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
package messages

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:1
)

import (
	"fmt"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/iana/msgtype"
	"github.com/jcmturner/gokrb5/v8/krberror"
	"github.com/jcmturner/gokrb5/v8/types"
)

// KRBPriv implements RFC 4120 type: https://tools.ietf.org/html/rfc4120#section-5.7.1.
type KRBPriv struct {
	PVNO			int			`asn1:"explicit,tag:0"`
	MsgType			int			`asn1:"explicit,tag:1"`
	EncPart			types.EncryptedData	`asn1:"explicit,tag:3"`
	DecryptedEncPart	EncKrbPrivPart		`asn1:"optional,omitempty"`	// Not part of ASN1 bytes so marked as optional so unmarshalling works
}

// EncKrbPrivPart is the encrypted part of KRB_PRIV.
type EncKrbPrivPart struct {
	UserData	[]byte			`asn1:"explicit,tag:0"`
	Timestamp	time.Time		`asn1:"generalized,optional,explicit,tag:1"`
	Usec		int			`asn1:"optional,explicit,tag:2"`
	SequenceNumber	int64			`asn1:"optional,explicit,tag:3"`
	SAddress	types.HostAddress	`asn1:"explicit,tag:4"`
	RAddress	types.HostAddress	`asn1:"optional,explicit,tag:5"`
}

// NewKRBPriv returns a new KRBPriv type.
func NewKRBPriv(part EncKrbPrivPart) KRBPriv {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:37
	_go_fuzz_dep_.CoverTab[88061]++
												return KRBPriv{
		PVNO:			iana.PVNO,
		MsgType:		msgtype.KRB_PRIV,
		DecryptedEncPart:	part,
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:42
	// _ = "end of CoverTab[88061]"
}

// Unmarshal bytes b into the KRBPriv struct.
func (k *KRBPriv) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:46
	_go_fuzz_dep_.CoverTab[88062]++
												_, err := asn1.UnmarshalWithParams(b, k, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.KRBPriv))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:48
		_go_fuzz_dep_.CoverTab[88065]++
													return processUnmarshalReplyError(b, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:49
		// _ = "end of CoverTab[88065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:50
		_go_fuzz_dep_.CoverTab[88066]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:50
		// _ = "end of CoverTab[88066]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:50
	// _ = "end of CoverTab[88062]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:50
	_go_fuzz_dep_.CoverTab[88063]++
												expectedMsgType := msgtype.KRB_PRIV
												if k.MsgType != expectedMsgType {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:52
		_go_fuzz_dep_.CoverTab[88067]++
													return krberror.NewErrorf(krberror.KRBMsgError, "message ID does not indicate a KRB_PRIV. Expected: %v; Actual: %v", expectedMsgType, k.MsgType)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:53
		// _ = "end of CoverTab[88067]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:54
		_go_fuzz_dep_.CoverTab[88068]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:54
		// _ = "end of CoverTab[88068]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:54
	// _ = "end of CoverTab[88063]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:54
	_go_fuzz_dep_.CoverTab[88064]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:55
	// _ = "end of CoverTab[88064]"
}

// Unmarshal bytes b into the EncKrbPrivPart struct.
func (k *EncKrbPrivPart) Unmarshal(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:59
	_go_fuzz_dep_.CoverTab[88069]++
												_, err := asn1.UnmarshalWithParams(b, k, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.EncKrbPrivPart))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:61
		_go_fuzz_dep_.CoverTab[88071]++
													return krberror.Errorf(err, krberror.EncodingError, "KRB_PRIV unmarshal error")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:62
		// _ = "end of CoverTab[88071]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:63
		_go_fuzz_dep_.CoverTab[88072]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:63
		// _ = "end of CoverTab[88072]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:63
	// _ = "end of CoverTab[88069]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:63
	_go_fuzz_dep_.CoverTab[88070]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:64
	// _ = "end of CoverTab[88070]"
}

// Marshal the KRBPriv.
func (k *KRBPriv) Marshal() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:68
	_go_fuzz_dep_.CoverTab[88073]++
												tk := KRBPriv{
		PVNO:		k.PVNO,
		MsgType:	k.MsgType,
		EncPart:	k.EncPart,
	}
	b, err := asn1.Marshal(tk)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:75
		_go_fuzz_dep_.CoverTab[88075]++
													return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:76
		// _ = "end of CoverTab[88075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:77
		_go_fuzz_dep_.CoverTab[88076]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:77
		// _ = "end of CoverTab[88076]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:77
	// _ = "end of CoverTab[88073]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:77
	_go_fuzz_dep_.CoverTab[88074]++
												b = asn1tools.AddASNAppTag(b, asnAppTag.KRBPriv)
												return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:79
	// _ = "end of CoverTab[88074]"
}

// EncryptEncPart encrypts the DecryptedEncPart within the KRBPriv.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:82
// Use to prepare for marshaling.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:84
func (k *KRBPriv) EncryptEncPart(key types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:84
	_go_fuzz_dep_.CoverTab[88077]++
												b, err := asn1.Marshal(k.DecryptedEncPart)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:86
		_go_fuzz_dep_.CoverTab[88080]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:87
		// _ = "end of CoverTab[88080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:88
		_go_fuzz_dep_.CoverTab[88081]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:88
		// _ = "end of CoverTab[88081]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:88
	// _ = "end of CoverTab[88077]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:88
	_go_fuzz_dep_.CoverTab[88078]++
												b = asn1tools.AddASNAppTag(b, asnAppTag.EncKrbPrivPart)
												k.EncPart, err = crypto.GetEncryptedData(b, key, keyusage.KRB_PRIV_ENCPART, 1)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:91
		_go_fuzz_dep_.CoverTab[88082]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:92
		// _ = "end of CoverTab[88082]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:93
		_go_fuzz_dep_.CoverTab[88083]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:93
		// _ = "end of CoverTab[88083]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:93
	// _ = "end of CoverTab[88078]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:93
	_go_fuzz_dep_.CoverTab[88079]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:94
	// _ = "end of CoverTab[88079]"
}

// DecryptEncPart decrypts the encrypted part of the KRBPriv message.
func (k *KRBPriv) DecryptEncPart(key types.EncryptionKey) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:98
	_go_fuzz_dep_.CoverTab[88084]++
												b, err := crypto.DecryptEncPart(k.EncPart, key, keyusage.KRB_PRIV_ENCPART)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:100
		_go_fuzz_dep_.CoverTab[88087]++
														return fmt.Errorf("error decrypting KRBPriv EncPart: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:101
		// _ = "end of CoverTab[88087]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:102
		_go_fuzz_dep_.CoverTab[88088]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:102
		// _ = "end of CoverTab[88088]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:102
	// _ = "end of CoverTab[88084]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:102
	_go_fuzz_dep_.CoverTab[88085]++
													err = k.DecryptedEncPart.Unmarshal(b)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:104
		_go_fuzz_dep_.CoverTab[88089]++
														return fmt.Errorf("error unmarshaling encrypted part: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:105
		// _ = "end of CoverTab[88089]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:106
		_go_fuzz_dep_.CoverTab[88090]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:106
		// _ = "end of CoverTab[88090]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:106
	// _ = "end of CoverTab[88085]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:106
	_go_fuzz_dep_.CoverTab[88086]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:107
	// _ = "end of CoverTab[88086]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:108
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/messages/KRBPriv.go:108
var _ = _go_fuzz_dep_.CoverTab
